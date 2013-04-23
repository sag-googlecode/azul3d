// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

// Test application - Opens multiple windows, uses OpenGL 1.5 rendering in each of them
package main

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/clock"
	"code.google.com/p/azul3d/native/opengl/1.5"
	"log"
	"math"
	"os"
	"runtime"
	"time"
)

func main() {
	log.SetFlags(0)

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	err := chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer chippy.Destroy()

	window := chippy.NewWindow()

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what the window currently has property-wise
	log.Println(window)

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
	context, err := window.GLCreateContext(1, 5)
	if err != nil {
		log.Fatal(err)
	}

	// Make the context current in this OS thread
	window.GLMakeCurrent(context)

	// Create an opengl.Context (which provides API access to an existing OpenGL context), for each
	// OpenGL context you wish to interace.
	//
	// We only make one here (as we are only using one context).
	gl := opengl.New()
	if gl == nil {
		log.Fatal("You have no support for OpenGL 1.1!")
	}
	log.Println(gl.GetError())

	// Initialize some things
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1.0)
	gl.DepthFunc(opengl.LESS)
	gl.Enable(opengl.DEPTH_TEST)
	gl.ShadeModel(opengl.SMOOTH)

	gl.MatrixMode(opengl.PROJECTION)
	gl.LoadIdentity()

	width, height := window.Size()

	// Alternative for gluPerspective.
	gluPerspective := func(fovY, aspect, zNear, zFar float64) {
		fH := math.Tan(fovY/360*math.Pi) * zNear
		fW := fH * aspect
		gl.Frustum(-fW, fW, -fH, fH, zNear, zFar)
	}

	gluPerspective(45.0, float64(width)/float64(height), 0.1, 100.0)

	gl.MatrixMode(opengl.MODELVIEW)

	// We'll use this clock for timing things
	clock := clock.New()

	// Start an goroutine to display statistics
	go func() {
		delay := 0 * time.Second
		for {
			<-time.After(delay)
			delay = 1 * time.Second

			// Print our FPS and average FPS
			log.Printf("FPS: %4.3f\tAverage: %4.3f\tDeviation: %f\n", clock.FrameRate(), clock.AverageFrameRate(), clock.FrameRateDeviation())
			//log.Println(1.0 / clock.Delta().Seconds())
		}
	}()

	var rot float64
	// Begin our rendering loop
	for !window.IsDestroyed() {
		runtime.Gosched()

		// Inform the clock that an new frame has begun
		clock.Tick()

		// Clear The Screen And The Depth Buffer
		gl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)
		gl.LoadIdentity() // Reset The View

		// Move into the screen 6.0 units.
		gl.Translatef(0.0, 0.0, -6.0)

		// We have smooth color mode on, this will blend across the vertices.
		// Draw a triangle rotated on the Y axis.
		gl.Rotatef(float32(rot), 0.0, 1.0, 0.0) // Rotate
		gl.Begin(opengl.POLYGON)                // Start drawing a polygon
		gl.Color3f(1.0, 0.0, 0.0)               // Red
		gl.Vertex3f(0.0, 1.0, 0.0)              // Top
		gl.Color3f(0.0, 1.0, 0.0)               // Green
		gl.Vertex3f(1.0, -1.0, 0.0)             // Bottom Right
		gl.Color3f(0.0, 0.0, 1.0)               // Blue
		gl.Vertex3f(-1.0, -1.0, 0.0)            // Bottom Left
		gl.End()                                // We are done with the polygon

		// Determine time since frame began
		delta := clock.Delta()

		// Increase the rotation by 90 degrees each second
		rot += 90.0 * delta.Seconds()

		// Clamp the result to 360 degrees
		if rot >= 360 {
			rot = 0
		}
		if rot < 0 {
			rot = 360
		}

		// Swap the display buffers
		window.GLSwapBuffers()
	}
}
