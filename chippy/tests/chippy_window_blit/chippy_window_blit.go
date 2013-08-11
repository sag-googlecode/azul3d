// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an windows, uses blitting to copy pixel graphics on to it
package main

import (
	"code.google.com/p/azul3d/chippy"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
	"time"
)

func program() {
	defer chippy.Exit()

	var err error

	// Load the image that we'll use for the window icon
	file, err := os.Open("src/code.google.com/p/azul3d/chippy/tests/data/chippy_720x320.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	rgba, ok := img.(*image.RGBA)
	if !ok {
		// Need to convert to RGBA image
		b := img.Bounds()
		rgba = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	}

	window := chippy.NewWindow()
	window.SetSize(720, 320)
	window.SetTransparent(true)
	window.SetDecorated(false)
	window.SetAlwaysOnTop(true)

	// Actually open the windows
	screen := chippy.DefaultScreen()

	// Center the window on the screen
	window.SetPositionCenter(screen)

	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what they currently has property-wise
	log.Println(window)

	events := window.Events()
	defer window.CloseEvents(events)

	for {
		// Clear the rectangle on the window
		window.PixelClear(rgba.Bounds())

		// Blit the image to the window, at x=0, y=0, blitting the entire image
		window.PixelBlit(0, 0, rgba)

		// Wait for an paint event
		gotPaintEvent := false
		for !gotPaintEvent {
			e := <-events
			switch e.(type) {
			case *chippy.PaintEvent:
				log.Println(e)
				gotPaintEvent = true

			case *chippy.CloseEvent:
				return
			}
		}
	}

	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)
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
