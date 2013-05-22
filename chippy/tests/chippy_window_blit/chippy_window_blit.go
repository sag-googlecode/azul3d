// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an windows, uses blitting to copy pixel graphics on to it
package main

import (
	"code.google.com/p/azul3d/chippy"
	_ "image/png"
	"image"
	"time"
	"log"
	"os"
)

func main() {
	var err error

	log.SetFlags(0)

	// Load the image that we'll use for the window icon
	file, err := os.Open("data/chippy_720x320.png")
	if err != nil {
		log.Fatal(err)
	}

	image, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	err = chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer chippy.Destroy()

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

	paintEvents := window.PaintEvents()

	for {
		// Clear the rectangle on the screen
		bounds := image.Bounds()
		window.PixelClear(0, 0, uint(bounds.Max.X), uint(bounds.Max.Y))

		// Blit the image to the window, at x=0, y=0, blitting the entire image
		window.PixelBlit(0, 0, image)

		// Wait for an paint event
		<- paintEvents.Read
	}

	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)
}
