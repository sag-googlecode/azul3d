// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example - Displays a few colored triangles.
package main

import (
	"azul3d.org/v1/chippy"
	"azul3d.org/v1/gfx"
	"azul3d.org/v1/gfx/window"
	"azul3d.org/v1/keyboard"
	"azul3d.org/v1/math"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

var glslVert = []byte(`
#version 120

attribute vec3 Vertex;
attribute vec4 Color;

uniform mat4 MVP;

varying vec4 frontColor;

void main()
{
	frontColor = Color;
	gl_Position = MVP * vec4(Vertex, 1.0);
}
`)

var glslFrag = []byte(`
#version 120

varying vec4 frontColor;

void main()
{
	gl_FragColor = frontColor;
}
`)

// gfxLoop is responsible for drawing things to the window. This loop must be
// independent of the Chippy main loop.
func gfxLoop(w *chippy.Window, r gfx.Renderer) {
	// Create a camera.
	camera := &gfx.Camera{
		Object: gfx.NewObject(),
	}

	// Setup the camera to use a perspective projection.
	camFOV := 75.0
	camNear := 0.0001
	camFar := 1000.0
	camera.SetPersp(r.Bounds(), camFOV, camNear, camFar)

	// Move the camera -2 on the Y axis (back two units away from the triangle
	// object).
	camera.Object.Pos.Y = -2

	// Update the camera's transformation matrix.
	camera.Transform = camera.Transform.Build()

	// Create a simple shader.
	shader := &gfx.Shader{
		Name:     "SimpleShader",
		GLSLVert: glslVert,
		GLSLFrag: glslFrag,
	}

	// Preload the shader (useful for seeing shader errors, if any).
	onLoad := make(chan *gfx.Shader, 1)
	r.LoadShader(shader, onLoad)
	go func() {
		<-onLoad
		shader.RLock()
		if shader.Loaded {
			fmt.Println("Shader loaded")
		} else {
			fmt.Println(string(shader.Error))
		}
		shader.RUnlock()
	}()

	// Create a triangle object.
	triangle := gfx.NewObject()
	triangle.Shader = shader
	triangle.OcclusionTest = true
	triangle.State.FaceCulling = gfx.NoFaceCulling
	triangle.Meshes = []*gfx.Mesh{
		&gfx.Mesh{
			Vertices: []gfx.Vec3{
				// Top
				{0, 0, 1},
				{-.5, 0, 0},
				{.5, 0, 0},

				// Bottom-Left
				{-.5, 0, 0},
				{-1, 0, -1},
				{0, 0, -1},

				// Bottom-Right
				{.5, 0, 0},
				{0, 0, -1},
				{1, 0, -1},
			},
			Colors: []gfx.Color{
				// Top
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},

				// Bottom-Left
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},

				// Bottom-Right
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},
			},
		},
	}

	go func() {
		event := w.Events()
		for e := range event {
			switch ev := e.(type) {
			case chippy.ResizedEvent:
				// Update the camera's projection matrix for the new width and
				// height.
				camera.Lock()
				camera.SetPersp(r.Bounds(), camFOV, camNear, camFar)
				camera.Unlock()

			case keyboard.TypedEvent:
				if ev.Rune == 'm' {
					// Toggle MSAA now.
					msaa := !r.MSAA()
					r.SetMSAA(msaa)
					fmt.Println("MSAA Enabled?", msaa)

				} else if ev.Rune == '1' {
					// Take a screenshot.
					fmt.Println("Writing screenshot to file...")
					// Download the image from the graphics hardware and save
					// it to disk.
					complete := make(chan image.Image, 1)
					r.Download(image.Rect(256, 256, 512, 512), complete)
					img := <-complete // Wait for download to complete.

					// Save to png.
					f, err := os.Create("screenshot.png")
					if err != nil {
						log.Fatal(err)
					}
					err = png.Encode(f, img)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Wrote texture to screenshot.png")
				}
			}
		}
	}()

	for {
		var v math.Vec2
		// Depending on keyboard state, transform the triangle.
		kb := w.Keyboard
		if kb.Down(keyboard.ArrowLeft) {
			v.X -= 1
		}
		if kb.Down(keyboard.ArrowRight) {
			v.X += 1
		}
		if kb.Down(keyboard.ArrowDown) {
			v.Y -= 1
		}
		if kb.Down(keyboard.ArrowUp) {
			v.Y += 1
		}

		// Update the triangle's transformation matrix.
		triangle.Lock()
		if kb.Down(keyboard.R) {
			// Reset transformation.
			triangle.Transform = gfx.DefaultTransform

		} else if kb.Down(keyboard.RightAlt) {
			// Apply shearing on X/Y axis.
			s := math.Vec3{v.Y, v.X, 0}
			if kb.Down(keyboard.RightShift) {
				// Apply shearing on X/Z axis.
				s = math.Vec3{v.Y, 0, v.X}
			}
			triangle.Shear = triangle.Shear.Add(s.MulScalar(0.05))

		} else if kb.Down(keyboard.LeftAlt) {
			// Apply scaling on X/Z axis.
			s := math.Vec3{v.X, 0, v.Y}
			if kb.Down(keyboard.LeftShift) {
				// Apply scaling on X/Y axis.
				s = math.Vec3{v.X, v.Y, 0}
			}
			triangle.Scale = triangle.Scale.Add(s.MulScalar(0.05))

		} else if kb.Down(keyboard.LeftCtrl) {
			// Apply rotation on X/Z axis.
			r := math.Vec3{v.Y, 0, v.X}
			if kb.Down(keyboard.LeftShift) {
				// Apply rotation on X/Y axis.
				r = math.Vec3{v.Y, v.X, 0}
			}
			triangle.Rot = triangle.Rot.Add(r.MulScalar(3))

		} else {
			// Apply movement on X/Z axis.
			p := math.Vec3{v.X, 0, v.Y}
			if kb.Down(keyboard.LeftShift) {
				// Apply movement on X/Y axis.
				p = math.Vec3{v.X, v.Y, 0}
			}
			triangle.Pos = triangle.Pos.Add(p.MulScalar(0.05))
		}

		// Update the triangle transformation matrix.
		triangle.Transform = triangle.Transform.Build()
		triangle.Unlock()

		// Clear the entire area (empty rectangle means "the whole area").
		r.Clear(image.Rect(0, 0, 0, 0), gfx.Color{1, 1, 1, 1})
		r.ClearDepth(image.Rect(0, 0, 0, 0), 1.0)

		// Clear a few rectangles on the window using different background
		// colors.
		r.Clear(image.Rect(0, 100, 640, 380), gfx.Color{0, 1, 0, 1})
		r.Clear(image.Rect(100, 100, 540, 380), gfx.Color{1, 0, 0, 1})
		r.Clear(image.Rect(100, 200, 540, 280), gfx.Color{0, 0.5, 0.5, 1})
		r.Clear(image.Rect(200, 200, 440, 280), gfx.Color{1, 1, 0, 1})

		// Draw the triangle to the screen.
		bounds := r.Bounds()
		r.Draw(bounds.Inset(50), triangle, camera)

		// Render the whole frame.
		r.Render()

		// Print the number of samples the triangle drew (only if the GPU
		// supports occlusion queries).
		if r.GPUInfo().OcclusionQuery {
			// The number of samples the triangle drew:
			samples := triangle.SampleCount()

			// The number of pixels the triangle drew:
			pixels := samples / r.Precision().Samples

			// The percent of the window that the triangle drew to:
			bounds := r.Bounds()
			percentage := float64(pixels) / float64(bounds.Dx()*bounds.Dy())

			fmt.Printf("Drew %v samples (%vpx, %f%% of window)\n", samples, pixels, percentage)
		}
	}
}

func main() {
	window.Run(gfxLoop)
}
