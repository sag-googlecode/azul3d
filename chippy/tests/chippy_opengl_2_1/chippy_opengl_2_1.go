// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Opens an window and uses OpenGL 2.1 rendering.
package main

import (
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/clock"
	"code.google.com/p/azul3d/native/gl"
	"code.google.com/p/azul3d/math"
	"unsafe"
	"log"
	"os"
	"runtime"
	"time"
)

var vertexShaderSource = []byte(`
#version 110

attribute vec4 Vertex;
attribute vec4 Color;

uniform mat4 Projection;
uniform mat4 Model;

varying vec4 colorVarying;

void main()
{
	colorVarying = Color;
	gl_Position = Projection * Model * Vertex;
}
`)

var fragmentShaderSource = []byte(`
#version 110

varying vec4 colorVarying;

void main()
{
	gl_FragColor = vec4(1, 0, 0, 1);
}
`)

var(
	// OpenGL context API access
	gl *opengl.Context

	// rotation of the triangle in degrees
	triangle *object

	// Shader variables
	shaderProgram, vertexShader, fragmentShader uint32

	// Perspective camera matrix
	perspective = new(math.Mat4)

	// Our OpenGL window
	window *chippy.Window

	// The clock used for measuring frame rate
	glClock *clock.Clock
)

type object struct {
	vertices []float32
	VertexAttribIndex, ColorAttribIndex uint32
	vboVertices, vboColors uint32
	rotation math.Real
	position *math.Vec3
	matrix *math.Mat4
}

// updateMatrix updates the matrix of the object to reflect the position of the
// object and the rotation of the object about the Y (up and down) axis.
func (o *object) updateMatrix() {
	// Set matrix to identity matrix
	o.matrix = math.Mat4Identity.Copy()

	// Set the translation of the matrix
	o.matrix.SetTranslation(o.position)

	// Convert our rotation from degrees to radians
	rads := o.rotation.Radians()
	o.matrix.SetRotation(
		rads,                  // Rotation in radians
		math.Vector3(0, 1, 0), // Rotate around the Y axis
		math.CoordSysYUpRight, // Y-Up Right-Handed coordinate system (the one OpenGL uses)
	)
}

// initScene is responsible for initializing the OpenGL scene.
func initScene() {
	// Background color white
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	// Enable depth testing
	gl.ClearDepth(1.0)
	gl.DepthFunc(opengl.LESS)
	gl.Enable(opengl.DEPTH_TEST)

	// Setup our perspective camera matrix
	width, height := window.Size()
	aspectRatio := float64(width) / float64(height)
	perspective.SetPerspective(45.0, math.Real(aspectRatio), 0.1, 100.0)

	// Initialize our triangle object
	triangle = new(object)

	// Triangle is 6 units into the screen
	triangle.position = math.Vector3(0, 0, -6.0)

	// Update position/rotation matrix of triangle object
	triangle.updateMatrix()


	// Later on, unbind the active VBO
	defer gl.BindBuffer(opengl.ARRAY_BUFFER, 0)


	// Create a Vertex Buffer Object to store vertices of the triangle
	triangle.vertices = []float32{
		 0.0,  1.0, 0.0, // Top
		 1.0, -1.0, 0.0, // Bottom Right
		-1.0, -1.0, 0.0, // Bottom Left
	}

	gl.GenBuffers(1, &triangle.vboVertices)
	gl.BindBuffer(opengl.ARRAY_BUFFER, triangle.vboVertices)
	gl.BufferData(opengl.ARRAY_BUFFER, uintptr(len(triangle.vertices)), unsafe.Pointer(&triangle.vertices[0]), opengl.STATIC_DRAW)
	gl.Execute()


	// Create a Vertex Buffer Object to store vertex colors of the triangle
	colors := []float32{
		1.0, 0.0, 0.0, 1.0, // Red
		0.0, 1.0, 0.0, 1.0, // Green
		0.0, 0.0, 1.0, 1.0, // Blue
	}

	gl.GenBuffers(1, &triangle.vboColors)
	gl.BindBuffer(opengl.ARRAY_BUFFER, triangle.vboColors)
	gl.BufferData(opengl.ARRAY_BUFFER, uintptr(len(colors)), unsafe.Pointer(&colors[0]), opengl.STATIC_DRAW)
	gl.Execute()

	handleShaderErrors := func(shader uint32, initialMessage string) {
		var ok int32
		gl.GetShaderiv(shader, opengl.COMPILE_STATUS, &ok)
		gl.Execute()
		if ok == opengl.FALSE {
			// Shader compiler error
			var logSize int32
			gl.GetShaderiv(shader, opengl.INFO_LOG_LENGTH, &logSize)
			gl.Execute()

			shaderLog := make([]byte, logSize)
			gl.GetShaderInfoLog(shader, uint32(logSize), nil, &shaderLog[0])
			gl.Execute()
			log.Println(initialMessage)
			log.Fatal(string(shaderLog))
		}
	}


	// Create vertex shader
	vertexShader = gl.CreateShader(opengl.VERTEX_SHADER)
	lengths := int32(len(vertexShaderSource))
	sources := &vertexShaderSource[0]
	gl.ShaderSource(vertexShader, 1, &sources, &lengths)
	gl.CompileShader(vertexShader)

	// Because gl.CompileShader above must first be executed for there to be
	// shader compilation errors.
	gl.Execute()
	handleShaderErrors(vertexShader, "Vertex shader has errors:")


	// Create fragment shader
	fragmentShader = gl.CreateShader(opengl.FRAGMENT_SHADER)
	lengths = int32(len(fragmentShaderSource))
	sources = &fragmentShaderSource[0]
	gl.ShaderSource(fragmentShader, 1, &sources, &lengths)
	gl.CompileShader(fragmentShader)

	// Because gl.CompileShader above must first be executed for there to be
	// shader compilation errors.
	gl.Execute()
	handleShaderErrors(fragmentShader, "Fragment shader has errors:")


	// Link the shader program all together
	shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)


	// Check for shader program link errors
	var ok int32
	gl.GetProgramiv(shaderProgram, opengl.LINK_STATUS, &ok)
	gl.Execute()
	if ok == opengl.FALSE {
		// Program link error
		var logSize int32
		gl.GetProgramiv(shaderProgram, opengl.INFO_LOG_LENGTH, &logSize)
		gl.Execute()

		shaderLog := make([]byte, logSize)
		gl.GetProgramInfoLog(shaderProgram, uint32(logSize), nil, &shaderLog[0])
		gl.Execute()
		log.Println("Program link errors:")
		log.Fatal(string(shaderLog))
	}
}

// destroyScene is responsible for cleaning up after OpenGL
func destroyScene() {
	gl.DeleteBuffers(1, &triangle.vboVertices)
	gl.DeleteBuffers(1, &triangle.vboColors)
}

// resizeScene is responsible for resizing the OpenGL scene when our window
// is resized.
func resizeScene(width, height int) {
	gl.Viewport(0, 0, uint32(width), uint32(height))

	aspectRatio := float64(width) / float64(height)
	perspective.SetPerspective(45.0, math.Real(aspectRatio), 0.1, 100.0)
}

// renderScene is responsible for rendering a single frame of the OpenGL scene.
func renderScene() {
	// Clear The Screen And The Depth Buffer
	gl.Clear(uint32(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT))

	// Determine time since frame began
	delta := glClock.Delta()

	// Increase the rotation of the triangle object by 90 degrees each second
	triangle.rotation += math.Real(90.0 * delta.Seconds())

	// Clamp the result to 360 degrees
	triangle.rotation = triangle.rotation.Clamp(0, 360)


	// Make the shader program active
	gl.UseProgram(shaderProgram)

	// After rendering the triangle, we can make no shader program active.
	defer gl.UseProgram(0)


	// We do an unsafe conversion to float32, as the matrix is made of up
	// [4][4]math.Real, and it needs a *float32 not *math.Real.
	//
	// math.Real is either float32 OR float64 depending on math.RealIsFloat64.
	matrixData := (*float32)(unsafe.Pointer(&perspective[0][0]))

	// Update perspective camera matrix input
	byteName := []byte("Projection")
	byteName = append(byteName, 0)
	location := gl.GetUniformLocation(shaderProgram, &byteName[0])
	if location > 0 {
		gl.UniformMatrix4fv(
			location,             // Index -- order in vertex shader
			1,                    // Just one 4x4 matrix
			opengl.GLBool(false), // transpose
			matrixData,           // A pointer to the actual matrix data
		)
	}

	// See above comments
	matrixData = (*float32)(unsafe.Pointer(&triangle.matrix[0][0]))

	// Update triangle matrix input
	byteName = []byte("Model")
	byteName = append(byteName, 0)
	location = gl.GetUniformLocation(shaderProgram, &byteName[0])
	if location > 0 {
		gl.UniformMatrix4fv(
			location,             // Index -- order in vertex shader
			1,                    // Just one 4x4 matrix
			opengl.GLBool(false), // transpose
			matrixData,           // A pointer to the actual matrix data
		)
	}
	gl.Execute()


	// After rendering the triangle, we can unbind the active buffer
	defer gl.BindBuffer(opengl.ARRAY_BUFFER, 0)

	// Bind the vertex VBO
	gl.BindBuffer(opengl.ARRAY_BUFFER, triangle.vboVertices)

	// Enable vertex attrib array and after rendering triangle, disable it.
	gl.EnableVertexAttribArray(0)
	defer gl.DisableVertexAttribArray(0)

	// Load vertices from the bound triangle VBO
	gl.VertexAttribPointer(0, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)


	// Bind the color VBO
	gl.BindBuffer(opengl.ARRAY_BUFFER, triangle.vboColors)

	// Enable vertex attrib array and after rendering triangle, disable it.
	gl.EnableVertexAttribArray(1)
	defer gl.DisableVertexAttribArray(1)

	// Load vertex attributes for color attribute index from the bound color VBO
	gl.VertexAttribPointer(1, 4, opengl.FLOAT, opengl.GLBool(false), 0, nil)


	// Draw the triangle
	gl.DrawArrays(opengl.TRIANGLES, 0, uint32(len(triangle.vertices)))

	gl.Flush()
}

// toggleVerticalSync is responsible for switching the vertical sync mode to
// the next one.
func toggleVerticalSync() {
	vsync := window.GLVerticalSync()

	switch(vsync) {
		case chippy.NoVerticalSync:
			vsync = chippy.VerticalSync

		case chippy.VerticalSync:
			vsync = chippy.AdaptiveVerticalSync

		case chippy.AdaptiveVerticalSync:
			vsync = chippy.NoVerticalSync
	}

	log.Println(vsync)
	window.GLSetVerticalSync(vsync)
}

// toggleMSAA is responsible for turning on/off OpenGL Multi Sample Anti
// Aliasing (MSAA)
var MSAA = true
func toggleMSAA() {
	if MSAA {
		MSAA = false
		gl.Disable(opengl.MULTISAMPLE)
	} else {
		MSAA = true
		gl.Enable(opengl.MULTISAMPLE)
	}
	log.Println("MSAA enabled?", MSAA)
}

func program() {
	defer chippy.Exit()

	window = chippy.NewWindow()

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err := window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print some instructions for the user
	log.Println("Instructions:")
	log.Println("v key - Toggle Vertical Sync")
	log.Println("m key - Toggle Multi Sample Anti Aliasing")
	log.Println("b key - Toggle OpenGL call batching")

	// Choose an buffer format, these include things like double buffering, bytes per pixel, number of depth bits, etc.
	configs := window.GLConfigs()

	// See documentation for this function and vars to see how it determines the 'best' format
	bestConfig := chippy.GLChooseConfig(configs, chippy.GLWorstConfig, chippy.GLBestConfig)
	window.GLSetConfig(bestConfig)

	// Print out all the formats, and which one we determined to be the 'best'.
	log.Println("\nChosen configuration:")
	log.Println(bestConfig)

	// All OpenGL related calls must occur in the same OS thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Create an OpenGL context with the OpenGL version we wish
	context, err := window.GLCreateContext(2, 1, chippy.GLCoreProfile, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make the context current in this OS thread
	window.GLMakeCurrent(context)

	// Create an opengl.Context (which provides API access to an existing OpenGL context), for each
	// OpenGL context you wish to interace.
	//
	// We only make one here (as we are only using one context).
	gl = opengl.New()

	// Ensure that we have at least GLSL version 110
	if !gl.AtLeastShaderVersion(1, 10) {
		log.Fatal("No support for GLSL 1.1 found!")
	}

	// Turn on batching of OpenGL commands
	gl.SetBatching(true)

	// Initialize some things
	initScene()

	// Cleanup later
	defer destroyScene()

	// We'll use this glClock for timing things
	glClock = clock.New()

	// Start an goroutine to display statistics
	go func() {
		for {
			<-time.After(1 * time.Second)

			// Print our FPS and average FPS
			log.Printf("FPS: %4.3f\tAverage: %4.3f\tDeviation: %f\n", glClock.FrameRate(), glClock.AverageFrameRate(), glClock.FrameRateDeviation())
		}
	}()

	events := window.Events()
	defer window.CloseEvents(events)

	// Begin our rendering loop
	for !window.Destroyed() {
		// Inform the clock that an new frame has begun
		glClock.Tick()

		for i := 0; i < len(events); i++ {
			e := <-events

			switch ev := e.(type) {
				case *chippy.ResizedEvent:
					resizeScene(ev.Width, ev.Height)

				case *keyboard.StateEvent:
					if ev.State == keyboard.Down {
						switch ev.Key {
						case keyboard.V:
							toggleVerticalSync()
						case keyboard.M:
							toggleMSAA()
						case keyboard.B:
							gl.SetBatching(!gl.Batching())
							log.Println("Batching?", gl.Batching())
						}
					}

				case *chippy.CloseEvent:
					return
			}
		}

		// Render the scene
		renderScene()

		// Execute all the pending OpenGL commands
		gl.Execute()

		// Swap the display buffers
		window.GLSwapBuffers()
	}
}

func main() {
	log.SetFlags(0)

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	// Initialize Chippy
	err := chippy.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Start program
	go program()

	// Enter main loop
	chippy.MainLoop()
}

