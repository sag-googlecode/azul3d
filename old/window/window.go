package window

import "code.google.com/p/azul3d/wrappers/glfw"
import "code.google.com/p/azul3d/render"
import "code.google.com/p/azul3d/clock"

import "runtime"
import "sync"

var hasInitGlfw bool
var hasInitRender bool

func waitEvents() {
    for{
        runtime.Gosched()
        glfw.WaitEvents()
    }
}

func init() {
    go waitEvents()
}

type Window struct {
    rendering, destroyed bool
    title string
    actual *glfw.Window
    dispatcher *render.Dispatcher
    clock *clock.Clock
    access sync.RWMutex
}

func New() (*Window, error) {
    if hasInitGlfw != true {
        err := glfw.Init()
        if err != nil {
            return nil, err
        }
        hasInitGlfw = true
    }
    if hasInitRender != true {
        err := render.Init()
        if err != nil {
            return nil, err
        }
        hasInitRender = true
    }

    width, height := 640, 480
    dx, dy := glfw.DesktopMode().Width, glfw.DesktopMode().Height
    rx := (dx/2) - (width/2)
    ry := (dy/2) - (height/2)

    glfw.WindowHint(glfw.POSITION_X, rx)
    glfw.WindowHint(glfw.POSITION_Y, ry)

    actual, err := glfw.CreateWindow(640, 480, glfw.WINDOWED, "Window", nil)
    if err != nil {
        return nil, err
    }

    window := new(Window)
    window.actual = actual
    window.clock = clock.New()
    window.dispatcher = render.NewDispatcher(window.actual, window.clock)
    return window, nil
}

func (this *Window) Dispatcher() *render.Dispatcher {
    this.access.RLock()
    defer this.access.RUnlock()
    return this.dispatcher
}

func (this *Window) Clock() *clock.Clock {
    this.access.RLock()
    defer this.access.RUnlock()
    return this.clock
}

func (this *Window) NewFrame(f func()) {
    this.access.RLock()
    defer this.access.RUnlock()
    this.dispatcher.NewFrame(f)
}

func (this *Window) SetRendering(rendering bool) {
    this.access.Lock()
    defer this.access.Unlock()
    if rendering != this.rendering {
        this.rendering = rendering
        if rendering {
            this.dispatcher.Start()
        } else {
            this.dispatcher.Stop()
        }
    }
}

func (this *Window) Rendering() bool {
    this.access.RLock()
    defer this.access.RUnlock()
    return this.rendering
}

func (this *Window) Destroy() {
    this.access.Lock()
    defer this.access.Unlock()
    this.destroyed = true
}

func (this *Window) Destroyed() bool {
    this.access.RLock()
    defer this.access.RUnlock()
    return this.destroyed
}

func (this *Window) CloseRequested() bool {
    this.access.RLock()
    defer this.access.RUnlock()
    if this.destroyed == false {
        return this.actual.Param(glfw.CLOSE_REQUESTED) != 0
    }
    return false
}

func (this *Window) SetTitle(title string) {
    this.actual.SetTitle(title)
    this.title = title
}

func (this *Window) Title() string {
    return this.title
}


func (this *Window) SetSize(width, height int) {
    this.actual.SetSize(width, height)
}

func (this *Window) Size() (int, int) {
    return this.actual.Size()
}

func (this *Window) SetPos(x, y int) {
    // :( glfw provides no SetPos API
}

func (this *Window) Pos() (int, int) {
    return this.actual.Param(glfw.POSITION_X), this.actual.Param(glfw.POSITION_Y)
}
