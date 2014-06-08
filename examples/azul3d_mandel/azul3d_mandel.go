// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example - Generates a mandelbrot on the CPU and displays it with the GPU.
package main

import (
	"azul3d.org/v1/chippy"
	"azul3d.org/v1/gfx"
	"azul3d.org/v1/gfx/gl2"
	"azul3d.org/v1/gfx/window"
	"azul3d.org/v1/keyboard"
	"azul3d.org/v1/mouse"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	gmath "math"
	"os"
)

var glslVert = []byte(`
#version 120

attribute vec3 Vertex;
attribute vec2 TexCoord0;

varying vec2 tc0;

void main()
{
	tc0 = TexCoord0;
	gl_Position = vec4(Vertex, 1.0);
}
`)

var glslFrag = []byte(`
#version 120

varying vec2 tc0;

uniform sampler2D Texture0;

void main()
{
	gl_FragColor = texture2D(Texture0, tc0);
}
`)

// gfxLoop is responsible for drawing things to the window. This loop must be
// independent of the Chippy main loop.
func gfxLoop(w *chippy.Window, r gfx.Renderer) {
	w.SetSize(720, 480)
	w.SetPositionCenter(chippy.DefaultScreen())
	glr := r.(*gl2.Renderer)
	glr.UpdateBounds(image.Rect(0, 0, 720, 480))

	// Create a simple shader.
	shader := gfx.NewShader("SimpleShader")
	shader.GLSLVert = glslVert
	shader.GLSLFrag = glslFrag

	// Create a card object.
	card := gfx.NewObject()
	card.Shader = shader
	card.Textures = []*gfx.Texture{nil}
	card.Meshes = []*gfx.Mesh{
		&gfx.Mesh{
			Vertices: []gfx.Vec3{
				// Left triangle.
				{-1, 1, 0},  // Left-Top
				{-1, -1, 0}, // Left-Bottom
				{1, -1, 0},  // Right-Bottom

				// Right triangle.
				{-1, 1, 0}, // Left-Top
				{1, -1, 0}, // Right-Bottom
				{1, 1, 0},  // Right-Top
			},
			TexCoords: []gfx.TexCoordSet{
				gfx.TexCoordSet{
					Slice: []gfx.TexCoord{
						// Left triangle.
						{0, 0},
						{0, 1},
						{1, 1},

						// Right triangle.
						{0, 0},
						{1, 1},
						{1, 0},
					},
				},
			},
		},
	}

	// Create a texture.
	zoom := 1.0
	x := -0.5
	y := 0.0
	res := 1
	maxIter := 1000
	updateTex := func() {
		width, height := w.Size()
		mbrot := Mandelbrot(width/res, height/res, maxIter, zoom, x, y)

		// Insert a small red square in the top-left of the image for ensuring
		// proper orientation exists in textures (this is just for testing).
		for x := 0; x < 20; x++ {
			for y := 0; y < 20; y++ {
				mbrot.Set(x, y, color.RGBA{255, 0, 0, 255})
			}
		}

		// Create new texture and ask the renderer to load it. We don't use DXT
		// compression because those textures cannot be downloaded.
		tex := &gfx.Texture{
			Source:    mbrot,
			MinFilter: gfx.Nearest,
			MagFilter: gfx.Nearest,
		}
		onLoad := make(chan *gfx.Texture, 1)
		r.LoadTexture(tex, onLoad)
		<-onLoad

		// Swap the texture with the old one on the card.
		card.Lock()
		card.Textures[0] = tex
		card.Unlock()
	}
	updateTex()

	go func() {
		handle := func(e chippy.Event) (needUpdate bool) {
			ev, ok := e.(mouse.Event)
			if ok {
				if ev.Button == mouse.Left && ev.State == mouse.Down {
					w.SetCursorGrabbed(!w.CursorGrabbed())
				}

				if ev.Button == mouse.Right && ev.State == mouse.Down {
					res += 2
					if res > 8 {
						res = 1
					}
					return true
				}

				if ev.Button == mouse.Wheel {
					if ev.State == mouse.ScrollForward {
						zoom += 0.06 * gmath.Abs(zoom)
					} else {
						zoom -= 0.06 * gmath.Abs(1.0-zoom)
					}
					return true
				}
			}

			kev, ok := e.(keyboard.TypedEvent)
			if ok {
				if kev.Rune == 's' || kev.Rune == 'S' {
					fmt.Println("Writing texture to file...")
					// Download the image from the graphics hardware and save
					// it to disk.
					complete := make(chan image.Image, 1)

					// Lock the card/texture.
					card.RLock()
					card.Textures[0].Lock()

					// Begin downloading it's texture.
					card.Textures[0].Download(image.Rect(0, 0, 640, 480), complete)

					// Unlock the texture/card.
					card.Textures[0].Unlock()
					card.RUnlock()

					img := <-complete // Wait for download to complete.
					if img == nil {
						fmt.Println("Failed to download texture.")
					} else {
						// Save to png.
						f, err := os.Create("mandel.png")
						if err != nil {
							log.Fatal(err)
						}
						err = png.Encode(f, img)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Wrote texture to mandel.png")
					}
				}
			}

			cev, ok := e.(chippy.CursorPositionEvent)
			if ok && w.CursorGrabbed() {
				x += (cev.X / 900.0) / gmath.Abs(zoom)
				y += (cev.Y / 900.0) / gmath.Abs(zoom)
				return true
			}
			return false
		}

		events := w.Events()
		for {
			e := <-events
			needUpdate := handle(e)
			l := len(events)
			for i := 0; i < l; i++ {
				if handle(<-events) {
					needUpdate = true
				}
			}
			if needUpdate {
				// Generate new mandel texture.
				updateTex()
			}
		}
	}()

	for {
		// Clear the entire area (empty rectangle means "the whole area").
		r.Clear(image.Rect(0, 0, 0, 0), gfx.Color{1, 1, 1, 1})
		r.ClearDepth(image.Rect(0, 0, 0, 0), 1.0)

		// Draw the card to the screen.
		r.Draw(image.Rect(0, 0, 0, 0), card, nil)

		// Render the whole frame.
		r.Render()
	}
}

func main() {
	window.Run(gfxLoop)
}
