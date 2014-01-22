package main

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/engine"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/color"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/geom/procedural"
	"code.google.com/p/azul3d/scene/shader"
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

	pos := ev.Data.(*chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform cubes.
	if !engine.Window.CursorGrabbed() {
		return
	}

	kb := engine.Window.Keyboard
	if kb.Down(keyboard.LeftCtrl) {
		// If left ctrl key is currently down, we apply scaling to current
		// cube.
		sx, sy, sz := currentCube.Scale()

		sx += math.Real(pos.X / 220)
		sz += math.Real(-pos.Y / 220)

		currentCube.SetScale(sx, sy, sz)

	} else if kb.Down(keyboard.RightCtrl) {
		// If right ctrl key is currently down, we apply shearing to current
		// cube.
		shx, shy, shz := currentCube.Shear()

		shy += math.Real(pos.X / 220)
		shz += math.Real(-pos.Y / 220)

		currentCube.SetShear(shx, shy, shz)

	} else if kb.Down(keyboard.LeftShift) || kb.Down(keyboard.RightShift) {
		// If an shift key is currently down, we apply relative rotation to the
		// current cube.
		pos.X /= 10
		pos.Y /= 10

		currentCube.SetRelativeRot(relativeCube, math.Real(pos.Y), 0, math.Real(-pos.X))

	} else {
		// Otherwise we apply relative movement to the current cube.
		pos.X /= 100
		pos.Y /= 100
		currentCube.SetRelativePos(relativeCube, math.Real(pos.X), 0, math.Real(-pos.Y))
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
	// Add four camera regions, which will give us four views of the same scene
	// on the window.

	// For the first region, we *know* that the 'azul3d' package adds a default
	// region to the camera. We'll just modify that one.
	topLeft := camera.Regions(engine.Camera3d)[0]
	topLeft.SetRegion(0, 0, 640/2, 480/2)
	camera.AddRegion(engine.Camera3d, topLeft)

	// For the other Camera3d regions, we'll create new ones.
	topRight := camera.NewRegion(640/2, 0, 640/2, 480/2)
	camera.AddRegion(engine.Camera3d, topRight)

	bottomLeft := camera.NewRegion(0, 480/2, 640/2, 480/2)
	camera.AddRegion(engine.Camera3d, bottomLeft)

	bottomRight := camera.NewRegion(640/2, 480/2, 640/2, 480/2)
	camera.AddRegion(engine.Camera3d, bottomRight)

	// We don't need to touch the Camera2d's region, because by default it is
	// set to draw without clearing the color, depth, or stencil buffers, and
	// with a sort value of 100 (higher than the above regions default sort
	// values of zero).

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
