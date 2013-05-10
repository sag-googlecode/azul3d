// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens two windows, changes each of their cursor properties
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
	file, err := os.Open("data/cursor_32x32_3x4.png")
	if err != nil {
		log.Fatal(err)
	}

	cursorImage, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// The last two numbers are the (X, Y) hotspot
	cursor := &chippy.Cursor{
		Image: cursorImage,
		X: 3,
		Y: 4,
	}

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	err = chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer chippy.Destroy()

	window1 := chippy.NewWindow()
	window2 := chippy.NewWindow()
	window1.SetTransparent(true)
	window2.SetTransparent(true)

	window1.SetCursor(cursor)

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window1.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	err = window2.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	window2.SetCursor(cursor)

	// Print out what they currently has property-wise
	log.Println(window1)
	log.Println(window2)

	log.Println("Waiting 5 seconds...")
	<-time.After(5 * time.Second)

	// We don't need those other cursors, get rid of them!
	window1.FreeCursor(cursor)
	window2.FreeCursor(cursor)

	// Use the default cursor
	window1.SetCursor(nil)
	window2.SetCursor(nil)

	// Just wait an while so they can enjoy the window
	log.Println("Waiting 5 seconds...")
	<-time.After(5 * time.Second)
}
