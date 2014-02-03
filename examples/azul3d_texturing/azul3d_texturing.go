package main

import (
	"azul3d.org/chippy"
	"azul3d.org/chippy/keyboard"
	"azul3d.org/engine"
	"azul3d.org/event"
	"azul3d.org/math"
	"azul3d.org/scene"
	"azul3d.org/scene/color"
	"azul3d.org/scene/geom"
	"azul3d.org/scene/geom/procedural"
	"azul3d.org/scene/renderer"
	"azul3d.org/scene/texture"
	"azul3d.org/scene/transparency"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"runtime"
	"runtime/debug"
	"sync"
)

var (
	// blue cube.
	blue *scene.Node

	// Lock for when we are changing any of the below variables
	globalLock sync.RWMutex
)

func createCube(name string, c color.Color) *scene.Node {
	// Create an geom node
	geomNode := scene.New(name)

	// Build the cube geom
	cube := procedural.Cube(1.0, geom.Static)
	cube.BakeColors(c, color.Default)

	// Attach it to the geom.Node
	geom.Add(geomNode, cube)

	// Enable multisample transparency on the node
	transparency.Set(geomNode, transparency.Multisample)

	// Load texture file
	//
	// You might try one of the following textures instead to increase load
	// time and show asynchronous file loading better.
	//
	// Note that the larger the texture the more memory it will consume, and it
	// could eat all your system's memory.
	prefix := "src/azul3d.org/assets/textures"
	//texFilePath := prefix + "/clouds_8192x8192.jpg"
	//texFilePath := prefix+"/clouds_4096x4096.jpg"
	//texFilePath := prefix+"/clouds_2048x2048.jpg"
	//texFilePath := prefix+"/clouds_1024x1024.jpg"
	texFilePath := prefix + "/texture_coords_1024x1024.png"

	tex, err := renderer.LoadTextureFile(engine.Renderer, texFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Assign texture to the cube
	texture.Set(geomNode, texture.DefaultLayer, tex)

	// Start an goroutine to print when the texture is loaded
	go func() {
		log.Println("Loading the texture...")
		<-tex.LoadNotify()
		log.Println("Loaded the texture.")
	}()

	return geomNode
}

// Event handler to manipulate cubes when cursor is moved
func onCursorPosition(ev *event.Event) {
	globalLock.RLock()
	defer globalLock.RUnlock()

	pos := ev.Data.(chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform cubes.
	if !engine.Window.CursorGrabbed() {
		return
	}

	kb := engine.Window.Keyboard
	if kb.Down(keyboard.LeftCtrl) {
		// If left ctrl key is currently down, we apply scaling to current
		// cube.
		sx, sy, sz := blue.Scale()

		sx += math.Real(pos.X / 10)
		sy += math.Real(-pos.Y / 10)
		sz += math.Real(-pos.Y / 10)

		blue.SetScale(sx, sy, sz)

	} else if kb.Down(keyboard.RightCtrl) {
		// If right ctrl key is currently down, we apply shearing to current
		// cube.
		shx, shy, shz := blue.Shear()

		shy += math.Real(pos.X / 10)
		shz += math.Real(-pos.Y / 10)

		blue.SetShear(shx, shy, shz)

	} else if kb.Down(keyboard.LeftShift) || kb.Down(keyboard.RightShift) {
		// If an shift key is currently down, we apply relative rotation to the
		// current cube.
		pos.X /= 10
		pos.Y /= 10

		blue.SetRelativeRot(blue, math.Real(pos.Y), 0, math.Real(-pos.X))

	} else {
		// Otherwise we apply relative movement to the current cube.
		pos.X /= 100
		pos.Y /= 100
		blue.SetRelativePos(blue, math.Real(pos.X), 0, math.Real(-pos.Y))
	}
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")

	blue.Detatch()
	blue = createCube("blue-cube", color.New(0, 0, 1, 1))
	blue.SetParent(engine.Scene2d)

	//blue.ResetTransform()
	blue.SetScale(50, 50, 50)
	blue.SetPos(50, 0, -50)
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed(ev *event.Event) {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}

// Event handler which toggles transparency
func toggleTransparency(ev *event.Event) {
	n := engine.Scene2d
	switch transparency.Mode(n) {
	case transparency.None:
		transparency.Set(n, transparency.AlphaBlend)
	case transparency.AlphaBlend:
		transparency.Set(n, transparency.Binary)
	case transparency.Binary:
		transparency.Set(n, transparency.Multisample)
	case transparency.Multisample:
		transparency.Set(n, transparency.None)
	}
	log.Println(transparency.Mode(n))
}

// Our scene graph will look like this:
//
// > Display
//     > Scene2d
//         > Camera2d
//         > Blue
//
func program() {
	// Blue cube will be an child of the 2D scene, it will look flat and have
	// no depth (Orthogonic camera lens is used to acheive this effect).
	blue = createCube("blue-cube", color.New(0, 0, 1, 1))
	blue.SetParent(engine.Scene2d)

	// Since it's in the 2D scene -- it's units are in pixels. The cube from
	// createCube() is two units wide, making it two pixels wide. We will make
	// it 200 pixels wide by changing it's scale ot 50 on the X, Y, and Z axis.
	blue.SetScale(50, 50, 50)

	// Also, since 2D scene starts at top-left of screen, we need to move the
	// blue cube 50 pixels to the right and down so that we can actually see
	// it.
	blue.SetPos(50, 0, -50)

	// Print scene graph
	engine.Renderer.PrintTree()

	// Grab the cursor
	engine.Window.SetCursorGrabbed(true)

	var stop func()
	stop = event.Define(event.Handlers{
		// Listen for alt keys to toggle cursor grabbed
		"RightAlt": toggleCursorGrabbed,
		"LeftAlt":  toggleCursorGrabbed,

		// Listen for R key to reset transformations
		"R": resetTransforms,

		// Listen for T key to change transparency modes
		"T": toggleTransparency,

		// Listen for F key to try and free texture
		"F": func(ev *event.Event) {
			log.Println("Free OS Memory")
			debug.FreeOSMemory()
		},

		"G": func(ev *event.Event) {
			log.Println("Garbage Collect")
			runtime.GC()
		},

		// Listen for cursor position to move the sprite
		"cursor-position": onCursorPosition,

		"window-destroyed": func(ev *event.Event) {
			stop()
		},
	})
}

func main() {
	// Run our program, enter main loop.
	engine.Run(program)
}
