// +build examples

package main

import(
	"code.google.com/p/azul3d/scene/geom/procedural"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/color"
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d"
	"sync"
	"log"
	"os"
)

var(
	// Create the engine.
	engine = azul3d.NewEngine()

	// These are the red, green, and blue cubes.
	red, green, blue *scene.Node

	// Lock for when we are changing any of the below variables
	globalLock sync.RWMutex

	// The current cube that we're manipulating, and the cube we're
	// manipulating relative to.
	currentCube, relativeCube *scene.Node
)

func createCube(name string, c color.Color) *scene.Node {
	// Create an geom node
	geomNode := scene.New(name)

	// Build the cube geom
	cube := procedural.Cube(1.0, geom.Static)
	cube.BakeColors(c, color.None)

	// Attach it to the geom.Node
	geom.Add(geomNode, cube)

	return geomNode
}

// Event handler to manipulate cubes when cursor is moved
func onCursorPosition(ev event.Event) {
	globalLock.RLock()
	defer globalLock.RUnlock()

	pos := ev.Data.(*chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform cubes.
	if !engine.Window.CursorGrabbed() {
		return
	}

	kb := engine.Window.Keyboard
	if kb.Down(keyboard.LeftCtrl) {
		cs, ok := color.Scale(currentCube)
		if !ok {
			cs = color.Default
		}

		x := float32(pos.X / 220)
		y := float32(-pos.Y / 220)

		color.SetScale(currentCube, color.New(
			cs.R + x,
			cs.G + (x + y) / 2,
			cs.B + y,
			cs.A,
		))
		log.Println(cs)

	} else {
		c, ok := color.Get(currentCube)
		if !ok {
			c = color.Default
		}

		x := float32(pos.X / 220)
		y := float32(-pos.Y / 220)

		color.Set(currentCube, color.New(
			c.R + x,
			c.G + (x + y) / 2,
			c.B + y,
			c.A,
		))
		log.Println(c)
	}
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetColors() {
	log.Println("Reset")

	color.Clear(red)
	color.ClearScale(red)

	color.Clear(green)
	color.ClearScale(green)

	color.Clear(blue)
	color.ClearScale(blue)
}

// Event handler which sets the current cube
func setCurrentCube(cube *scene.Node) {
	globalLock.Lock()
	defer globalLock.Unlock()

	currentCube = cube

	// Print cube changing
	log.Println(currentCube.Name())
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed() {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}


// Our scene graph will look like this:
//
// > Renderer
//     > Scene3d
//         > Camera3d
//         > Red
//             > Green
//                 > Blue
//
func program() {
	// Set camera 12 units back (Y is depth) so that we see the cubes placed in
	// the center of the scene
	engine.Camera3d.SetPos(0, -12, 0)

	// Red cube will be an child of the scene, that way it's seen by the
	// camera.
	red = createCube("red-cube", color.New(1, 0, 0, 1))
	red.SetParent(engine.Scene3d)

	// Green cube will be an child of the red cube.
	green = createCube("green-cube", color.New(0, 1, 0, 1))
	green.SetPos(2, 0, 0)
	green.SetParent(red)

	// Blue cube will be an child of the 2D scene, it will look flat and have
	// no depth (Orthogonic camera lens is used to acheive this effect).
	blue = createCube("blue-cube", color.New(0, 0, 1, 1))
	blue.SetPos(2, 0, 0)
	blue.SetParent(green)

	// Print scene graph
	engine.Renderer.PrintTree()

	// Grab the cursor
	engine.Window.SetCursorGrabbed(true)

	// The current cube is the cube we are currently moving, and the relative
	// cube is the cube we are moving relative to.
	currentCube = red


	// Listen for alt keys to toggle cursor grabbed
	keyRightAlt := event.Notify("RightAlt")
	defer event.Close(keyRightAlt)

	keyLeftAlt := event.Notify("LeftAlt")
	defer event.Close(keyLeftAlt)

	// Listen for R key to reset colors
	keyR := event.Notify("R")
	defer event.Close(keyR)

	// Add event listeners for 1-4 keys to change relative cube.
	keyOne := event.Notify("One")
	keyTwo := event.Notify("Two")
	keyThree := event.Notify("Three")
	keyFour := event.Notify("Four")

	// Listen for cursor position to move the sprite
	cursorPosition := event.Notify("cursor-position")
	defer event.Close(cursorPosition)

	destroyed := event.Notify("window-destroyed")
	defer event.Close(destroyed)

	for{
		select{
		case <-keyRightAlt:
			toggleCursorGrabbed()

		case <-keyLeftAlt:
			toggleCursorGrabbed()

		case <-keyR:
			resetColors()

		case <-keyOne:
			setCurrentCube(red)
		case <-keyTwo:
			setCurrentCube(green)
		case <-keyThree:
			setCurrentCube(blue)
		case <-keyFour:
			setCurrentCube(nil)

		case ev := <-cursorPosition:
			onCursorPosition(ev)

		case <-destroyed:
			return
		}
	}
}

func main() {
	// For debugging anything
	azul3d.SetDebugOutput(os.Stdout)

	// Initialize azul3d
	err := azul3d.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Launch program
	go program()

	// Enter main loop
	azul3d.MainLoop()
}
