// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package renderer

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy/mouse"
	"code.google.com/p/azul3d/clock"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/renderer/opengl"
	"code.google.com/p/azul3d/scene/shader"
	"code.google.com/p/azul3d/scene/texture"
	"code.google.com/p/azul3d/scene/util"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Renderer is an generic interface which must convert an tree of nodes into an
// render function which will, when called, effectively renders the tree of
// nodes.
type backend interface {
	Render(root *scene.Node, region *util.Region) func()
	Resize(x, y int)
	LoadTexture(t texture.Type)
	LoadMesh(m *geom.Mesh)
	LoadShader(s *shader.Shader)
	MaxTextureCoords() int
	MaxTextureLayers() int
	MaxTextureSize() int
	GPUName() string
}

var (
	objTag = scene.NewProp("renderer.obj")
)

type obj struct {
	sync.RWMutex

	node *scene.Node

	Clock  *clock.Clock
	Window *chippy.Window
	Region *util.Region

	renderer backend
	backend  BackendType

	maxBufferedFrames uint
	playing           bool

	rendererCreateExecute  chan func()
	rendererCreateComplete chan bool
	preparedFrames         chan func()
	wantChangeBufferSize   chan uint

	vsyncChanged bool
	vsync        chippy.VSyncMode

	maxTextureCoords, maxTextureLayers, maxTextureSize int
	gpuName                                            string

	currentlyChangingBuffer sync.Mutex

	displayContext, loaderContext chippy.GLContext
}

func (n *obj) renderLoop() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	paused, stop := event.Notify("renderer-paused")
	defer stop()

	destroyed, stop := event.Notify("window-destroyed")
	defer stop()

	for {
		select {
		case max := <-n.wantChangeBufferSize:
			// Change the buffer size to max, to do that we need to first empty the buffer.

			// To empty the buffer, we need to simply execute each frame that is currently
			// prepared and waiting in the buffer, also we need to ensure that no new frames
			// are added to the buffer (we block this by using this lock).
			n.currentlyChangingBuffer.Lock()

			// Empty the buffer by executing each prepared frame in the buffer.
		loop:
			for {
				select {
				case frame := <-n.preparedFrames:
					// Execute the prepared frame that we received from the buffered frame
					// channel.
					frame()

				default:
					// Window was destroyed while trying to change buffer size,
					// so we give up.
					if n.Window.Destroyed() {
						// Pay attention to the lock here!
						n.currentlyChangingBuffer.Unlock()
						return
					}

					// No more prepared frames, we can continue to replace the buffered
					// frame channel.
					break loop
				}
			}

			// All good now, there are no prepared frames that need to be executed (we executed
			// them all above), we can now go forth with changing the buffer size on the frame
			// channel.
			n.preparedFrames = make(chan func(), max)

			// Last but certainly not least, allow frames to be added to the buffer again.
			n.currentlyChangingBuffer.Unlock()

		case f := <-n.rendererCreateExecute:
			// Renderer creation render thread -- execute function
			f()
			// Signal it's completion
			n.rendererCreateComplete <- true

		case frame := <-n.preparedFrames:
			// An new frame has been prepared for rendering, simply execute it.
			frame()
			runtime.Gosched()

		case e := <-paused:
			which := e.Data.(*scene.Node)
			if which == n.node {
				// We want to pause now, so wait until another signal is sent over the channel
				// to begin rendering again.
				which = nil
				for which != n.node {
					e = <-paused
					which = e.Data.(*scene.Node)
				}

				// Now we begin rendering again.
			}

		case <-destroyed:
			return
		}
	}
}

func (n *obj) buildFramesLoop() {
	resized, stop := event.Notify("window-resized")
	defer stop()

	destroyed, stop := event.Notify("window-destroyed")
	defer stop()

	paused, stop := event.Notify("renderer-paused")
	defer stop()

	for {
		select {
		case e := <-paused:
			which := e.Data.(*scene.Node)
			if which == n.node {
				// We want to pause now, so wait until another signal is sent over the channel
				// to begin rendering again.
				which = nil
				for which != n.node {
					e = <-paused
					which = e.Data.(*scene.Node)
				}

				// Now we begin rendering again.
			}

		case e := <-resized:
			buffered := len(resized)
			for i := 0; i < buffered; i++ {
				e = <-resized
			}
			ev := e.Data.(*chippy.ResizedEvent)

			n.renderFrame(func() {
				n.renderer.Resize(ev.Width, ev.Height)
			})

		case <-destroyed:
			return

		default:
			n.renderFrame(nil)
			runtime.Gosched()
		}
	}
}

// renderFrame builds and pushes the next frame of the scene described below
// this node onto the always-running render loop.
//
// If the rendering loop is currently busy (rendering an frame already, or is
// paused) then this frame will be buffered up to the value returned by
// BufferedFrames().
//
// If the frame cannot be buffered, then this call will block untill the
// rendering loop has created room for it, or untill Play() has been called
// again.
func (n *obj) renderFrame(pre func()) {
	vsyncChanged := n.vsyncChanged
	vsync := n.vsync
	if n.vsyncChanged {
		n.vsyncChanged = false
	}

	// Wrap the frame with an function which informs the clock of an new frame,
	// and executes the pre function, in addition to swapping the window's
	// buffers.
	wrappedFrame := func() {
		// Inform the clock that an new frame has just begun.
		n.Clock.Tick()

		// Send pre-frame event
		event.Send("pre-frame", n.node)

		if pre != nil {
			pre()
		}

		if vsyncChanged {
			n.Window.GLSetVerticalSync(vsync)
		}

		// Prepare the frame at this moment.
		frame := n.renderer.Render(n.node, n.Region)
		frame()

		if n.backend == OpenGL {
			n.Window.GLSwapBuffers()
		}

		// Send frame event.
		event.Send("frame", n.node)
	}

	// Try to acquire the changing lock, the scenario in which we cannot acquire it is when the
	// rendering loop is changing the size of the buffered frame channel.
	//
	// After we have acquired it, we can simply release it immedietly as it does not synchronize
	// anything except changing the buffered frame channel.
	n.currentlyChangingBuffer.Lock()
	n.currentlyChangingBuffer.Unlock()

	for !n.Window.Destroyed() {
		select {
		case n.preparedFrames <- wrappedFrame:
			// ^ Push an new prepared frame onto the channel

			// Break loop
			return

		case <-time.After(5 * time.Second):
			// Continue since window might be destroyed now.
			continue
		}
	}
}

func (n *obj) eventLoop() {
	events := n.Window.Events()
	defer n.Window.CloseEvents(events)

	for {
		e := <-events
		switch ev := e.(type) {
		case *chippy.CloseEvent:
			go event.Send("window-close", ev)

		case *chippy.CursorPositionEvent:
			go event.Send("cursor-position", ev)

		case *chippy.CursorWithinEvent:
			go event.Send("cursor-within", ev)

		case *chippy.FocusedEvent:
			go event.Send("window-focused", ev)

		case *chippy.MaximizedEvent:
			go event.Send("window-maximized", ev)

		case *chippy.MinimizedEvent:
			go event.Send("window-minimized", ev)

		case *chippy.PaintEvent:
			go event.Send("window-paint", ev)

		case *chippy.PositionEvent:
			go event.Send("window-position", ev)

		case *chippy.ResizedEvent:
			go event.Send("window-resized", ev)

		case *chippy.ScreenChangedEvent:
			go event.Send("window-screen-changed", ev)

		case *mouse.Event:
			eventName := fmt.Sprintf("mouse-%s-%s", ev.Button, ev.State)
			go event.Send(eventName, ev)
			go event.Send("mouse-state", ev)

		case *keyboard.StateEvent:
			if ev.State == keyboard.Down {
				go event.Send(ev.Key.String(), ev)
			} else {
				eventName := fmt.Sprintf("%s-%s", ev.Key, ev.State)
				go event.Send(eventName, ev)
			}
			go event.Send("keyboard-state", ev)

		case *keyboard.TypedEvent:
			go event.Send("keyboard-typed", ev)

		case *chippy.DestroyedEvent:
			go event.Send("window-destroyed", ev)
			return
		}
	}
}

/*
func (o *obj) calculateBounds() {
	preFrame, stop := event.Notify("pre-frame")
	defer stop()

	destroyed, stop := event.Notify("window-destroyed")
	defer stop()

	for {
		select{
		case <-preFrame:
			geom.CalculateBounds(o.node)

		case <-destroyed:
			return
		}
	}
}
*/

func getObj(n *scene.Node) *obj {
	o, ok := n.Tag(objTag)
	if !ok {
		return nil
	}
	return o.(*obj)
}

func mustGetObj(n *scene.Node) *obj {
	o := getObj(n)
	if o != nil {
		return o
	}
	panic("Specified node is not an renderer node.")
}

func GPUName(n *scene.Node) string {
	o := mustGetObj(n)
	return o.gpuName
}

func MaxTextureCoords(n *scene.Node) int {
	o := mustGetObj(n)
	return o.maxTextureCoords
}

func MaxTextureLayers(n *scene.Node) int {
	o := mustGetObj(n)
	return o.maxTextureLayers
}

func MaxTextureSize(n *scene.Node) int {
	o := mustGetObj(n)
	return o.maxTextureSize
}

// Create creates an renderer node on the specified scene node. An renderer
// node is responsible for taking all nodes parented below it and rendering
// them appropriately into an window or offscreen buffer.
//
// The window parameter specifies which window will be used for rendering.
func Create(n *scene.Node, window *chippy.Window) {
	o := getObj(n)
	if o == nil {
		o = new(obj)
		o.Clock = clock.New()
		o.Window = window
		o.Region = util.NewRegion(0, 0, 0, 0)

		// Limit clock to 75FPS, for drivers that do not support vsync
		o.Clock.SetMaxFrameRate(75)

		// Use adaptive vsync -- most should have this and it's better.
		o.vsyncChanged = true
		o.vsync = chippy.AdaptiveVerticalSync

		o.rendererCreateExecute = make(chan func())
		o.rendererCreateComplete = make(chan bool)
		o.preparedFrames = make(chan func(), 0)
		o.wantChangeBufferSize = make(chan uint)
		o.playing = true

		go o.eventLoop()
		go o.renderLoop()
		//go o.calculateBounds()

		o.node = n
		n.SetTag(objTag, o)
	}
}

func Destroy(n *scene.Node) {
	o := getObj(n)
	if o != nil {
		o.Window.Destroy()
		n.ClearTag(objTag)
	}
}

func Is(n *scene.Node) bool {
	return getObj(n) != nil
}

func Clock(n *scene.Node) *clock.Clock {
	o := mustGetObj(n)
	return o.Clock
}

func Window(n *scene.Node) *chippy.Window {
	o := mustGetObj(n)
	return o.Window
}

func SetBackend(n *scene.Node, b BackendType) error {
	if !b.Valid() {
		panic("SetBackend(): Invalid backend type specified!")
	}

	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	var (
		err    error
		render backend
	)

	//FIXME: expose
	glContextFlags := chippy.GLCoreProfile

	switch b {
	case OpenGL:
		//FIXME: expose
		o.Window.GLSetConfig(chippy.GLChooseConfig(o.Window.GLConfigs(), chippy.GLWorstConfig, chippy.GLBestConfig))

		o.loaderContext, err = o.Window.GLCreateContext(2, 0, glContextFlags, nil)
		if err != nil {
			return err
		}

		o.displayContext, err = o.Window.GLCreateContext(2, 0, glContextFlags, o.loaderContext)
		if err != nil {
			return err
		}

		dcMakeCurrent := func(current bool) {
			if current {
				o.Window.GLMakeCurrent(o.displayContext)
			} else {
				o.Window.GLMakeCurrent(nil)
			}
		}

		lcMakeCurrent := func(current bool) {
			if current {
				o.Window.GLMakeCurrent(o.loaderContext)
			} else {
				o.Window.GLMakeCurrent(nil)
			}
		}

		o.rendererCreateExecute <- func() {
			render, err = opengl.NewGLFFRenderer(dcMakeCurrent, lcMakeCurrent)
			if err == nil {
				width, height := o.Window.Size()
				render.Resize(width, height)

				o.maxTextureCoords = render.MaxTextureCoords()
				o.maxTextureLayers = render.MaxTextureLayers()
				o.maxTextureSize = render.MaxTextureSize()
				o.gpuName = render.GPUName()
			}
		}
		<-o.rendererCreateComplete
	}

	if err != nil {
		return err
	}

	o.renderer = render
	if o.renderer == nil {
		panic("Renderer is nil!")
	}

	o.backend = b

	if o.backend == 0 {
		go o.buildFramesLoop()
	}

	return nil
}

func Backend(n *scene.Node) BackendType {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return o.backend
}

// Pause causes the rendering loop to stop rendering.
func Pause(n *scene.Node) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	if o.playing {
		o.playing = false
		event.Send("renderer-paused", n)
	}
}

// Play causes the rendering loop to begin rendering again.
func Play(n *scene.Node) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	if !o.playing {
		o.playing = true
		event.Send("renderer-paused", n)
	}
}

// Playing tells wether or not the rendering loop is currently rendering.
func Playing(n *scene.Node) bool {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return o.playing
}

// SetMaxBufferedFrames specifies the maximum number of frames that can be
// buffered before an call to RenderFrame() will block.
//
// Default: 3
func SetMaxBufferedFrames(n *scene.Node, max uint) {
	o := mustGetObj(n)

	if o.Window.Destroyed() {
		return
	}

	//destroyedEvents := o.Window.DestroyedEvents()
	//defer destroyedEvents.Close()

	select {
	case o.wantChangeBufferSize <- max:
		return

		//case <-destroyedEvents.Read:
		//return
	}
}

// MaxBufferedFrames returns the maximimum number of frames that can be
// buffered before an call to RenderFrame() will block.
//
// Default: 3
func MaxBufferedFrames(n *scene.Node) uint {
	o := mustGetObj(n)

	return uint(cap(o.preparedFrames))
}

// BufferedFrames returns the current number of buffered frames.
func BufferedFrames(n *scene.Node) uint {
	o := mustGetObj(n)

	return uint(len(o.preparedFrames))
}

// SetVerticalSync sets the vertical sync value of this renderer node.
//
// The vertical sync value must be valid or else an panic will occur.
func SetVerticalSync(n *scene.Node, vsync chippy.VSyncMode) {
	if !vsync.Valid() {
		panic("SetVerticalSync(): Invalid vertical sync value specified!")
	}

	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.vsyncChanged = true
	o.vsync = vsync
}

// VerticalSync returns the vertical sync value of this renderer node.
//
// The vertical sync value must be valid or else an panic will occur.
func VerticalSync(n *scene.Node) chippy.VSyncMode {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return o.vsync
}