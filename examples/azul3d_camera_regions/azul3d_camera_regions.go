package main

import (
	"azul3d.org/chippy"
	"azul3d.org/chippy/keyboard"
	"azul3d.org/engine"
	"azul3d.org/event"
	"azul3d.org/scene"
	"azul3d.org/scene/camera"
	"azul3d.org/scene/color"
	"azul3d.org/scene/geom"
	"azul3d.org/scene/geom/procedural"
	"azul3d.org/scene/shader"
	"log"
	"sync"
)

var (
	// These are the red, green, and blue cubes.
	red, green, blue *scene.Node

	// Lock for when we are changing any of the below variables
	globalLock sync.RWMutex

	// The current cube that we're manipulating, and the cube we're
	// manipulating relative to.
	currentCube, relativeCube *scene.Node

	// Color shader
	colorShader = shader.New("colorShader")
)

func init() {
	// Vertex shader
	colorShader.SetSource([]byte(`
#version 110

attribute vec4 Vertex;
attribute vec4 Color;

uniform mat4 Projection;
uniform mat4 ModelViewProjection;

void main()
{
	gl_FrontColor = Color;
	gl_Position = ModelViewProjection * Vertex;
}
`), shader.Vertex)

	// Fragment shader
	colorShader.SetSource([]byte(`
#version 110

void main()
{
	gl_FragColor = gl_Color;
}
`), shader.Fragment)
}

func createCube(name string, c color.Color) *scene.Node {
	// Create an geom node
	geomNode := scene.New(name)

	// Build the cube geom
	cube := procedural.Cube(1.0, geom.Static)
	cube.BakeColors(c, color.Default)

	// Attach it to the geom.Node
	geom.Add(geomNode, cube)

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
		sx, sy, sz := currentCube.Scale()

		sx += float64(pos.X / 220)
		sz += float64(-pos.Y / 220)

		currentCube.SetScale(sx, sy, sz)

	} else if kb.Down(keyboard.RightCtrl) {
		// If right ctrl key is currently down, we apply shearing to current
		// cube.
		shx, shy, shz := currentCube.Shear()

		shy += float64(pos.X / 220)
		shz += float64(-pos.Y / 220)

		currentCube.SetShear(shx, shy, shz)

	} else if kb.Down(keyboard.LeftShift) || kb.Down(keyboard.RightShift) {
		// If an shift key is currently down, we apply relative rotation to the
		// current cube.
		pos.X /= 10
		pos.Y /= 10

		currentCube.SetRelativeRot(relativeCube, float64(pos.Y), 0, float64(-pos.X))

	} else {
		// Otherwise we apply relative movement to the current cube.
		pos.X /= 100
		pos.Y /= 100
		currentCube.SetRelativePos(relativeCube, float64(pos.X), 0, float64(-pos.Y))
	}
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")

	red.ResetTransform()
	green.ResetTransform()
	blue.ResetTransform()

	green.SetPos(2, 0, 0)
	blue.SetScale(50, 50, 50)
	blue.SetPos(50, 0, -50)
}

func printViewed(ev *event.Event) {
	redInView := camera.InView(engine.Camera3d, red.PosVec3(), red.Parent())
	greenInView := camera.InView(engine.Camera3d, green.PosVec3(), green.Parent())
	blueInView := camera.InView(engine.Camera2d, blue.PosVec3(), blue.Parent())

	log.Printf("Visibility: %s=%t | %s=%t | %s=%t\n", red.Name(), redInView, green.Name(), greenInView, blue.Name(), blueInView)
}

// Event handler which sets the relative cube
func setRelativeCube(cube *scene.Node) {
	globalLock.Lock()
	defer globalLock.Unlock()

	relativeCube = cube

	// Print out relation used for relative movement/rotation
	relativeCubeName := "nil"
	if relativeCube != nil {
		relativeCubeName = relativeCube.Name()
	}
	log.Println(currentCube.Name(), "RELATIVE TO", relativeCubeName)
}

// Event handler which sets the current cube
func setCurrentCube(cube *scene.Node) {
	globalLock.Lock()
	defer globalLock.Unlock()

	currentCube = cube

	// Print out relation used for relative movement/rotation
	relativeCubeName := "nil"
	if relativeCube != nil {
		relativeCubeName = relativeCube.Name()
	}
	log.Println(currentCube.Name(), "RELATIVE TO", relativeCubeName)
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed(ev *event.Event) {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}

// Our scene graph will look like this:
//
// > Display
//     > Scene
//         > Camera
//         > Red
//             > Green
//     > Scene2d
//         > Camera2d
//         > Blue
//
func program() {
	// The 'engine' package sets up a few camera regions for us already -- but
	// we don't want to use those so we remove them now.
	camera.ClearRegions(engine.Camera3d)
	camera.ClearRegions(engine.Camera2d)

	// We will have four camera regions which displays the 3D scene.
	topLeft := camera.NewRegion(0, 0, 0.5, 0.5)
	topRight := camera.NewRegion(0.5, 0, 0.5, 0.5)
	bottomLeft := camera.NewRegion(0, 0.5, 0.5, 0.5)
	bottomRight := camera.NewRegion(0.5, 0.5, 0.5, 0.5)

	// Add the regions to the 3D camera.
	camera.AddRegion(engine.Camera3d, topLeft)
	camera.AddRegion(engine.Camera3d, topRight)
	camera.AddRegion(engine.Camera3d, bottomLeft)
	camera.AddRegion(engine.Camera3d, bottomRight)

	// We will have one single camera region which displays the 2D scene.
	whole := camera.NewRegion(0, 0, 1, 1)

	// We want to make sure that the Camera2d's region renders *after* the 3D
	// ones above so that we have the effect of an overlay. To make the 'whole'
	// region render last we can make the sort value larger that the others
	// (zero by default).
	whole.SetSort(100)

	// It's important that the 'whole' region does not try to clear the color
	// buffer, or else our 3D scene below it would not show up!
	whole.SetClearColorActive(false)

	// Add the region to the 2D camera.
	camera.AddRegion(engine.Camera2d, whole)

	// Set camera 12 units back (Y is depth) so that we see the cubes placed in
	// the center of the scene
	engine.Camera3d.SetPos(0, -12, 0)

	// Color shader will affect all nodes below the 2D and 3D scene nodes
	shader.Set(engine.Scene2d, colorShader)
	shader.Set(engine.Scene3d, colorShader)

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

	// The current cube is the cube we are currently moving, and the relative
	// cube is the cube we are moving relative to.
	currentCube = red
	relativeCube = red

	var stop func()
	stop = event.Define(event.Handlers{
		// Listen for alt keys to toggle cursor grabbed
		"RightAlt": toggleCursorGrabbed,
		"LeftAlt":  toggleCursorGrabbed,

		// Listen for R key to reset transformations
		"R": resetTransforms,

		// Listen for V key to print view status
		"V": printViewed,

		// Listen for 1-4 keys to change relative cube.
		"One": func(ev *event.Event) {
			setRelativeCube(red)
		},
		"Two": func(ev *event.Event) {
			setRelativeCube(green)
		},
		"Three": func(ev *event.Event) {
			setRelativeCube(blue)
		},
		"Four": func(ev *event.Event) {
			setRelativeCube(nil)
		},

		// Listen for F1-F4 keys to change current cube.
		"F1": func(ev *event.Event) {
			setCurrentCube(red)
		},
		"F2": func(ev *event.Event) {
			setCurrentCube(green)
		},
		"F3": func(ev *event.Event) {
			setCurrentCube(blue)
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
