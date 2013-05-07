// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an window and sets the cursor position
package main

import (
	"code.google.com/p/azul3d/chippy"
	"log"
	"os"
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

	// Actually open the window
	screen := chippy.DefaultScreen()
	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	window.SetCursorPosition(0, 0)

	// Print out what they currently has property-wise
	log.Println(window)

	log.Println("Screen top left")
	x, y := window.Position()
	window.SetCursorPosition(-x, -y)
	<-time.After(5 * time.Second)

	log.Println("Screen bottom right")
	screenWidth, screenHeight := window.Screen().ScreenMode().Resolution()
	window.SetCursorPosition(int(screenWidth), int(screenHeight))
	<-time.After(5 * time.Second)

	log.Println("Window top left")
	window.SetCursorPosition(0, 0)
	<-time.After(5 * time.Second)

	log.Println("Window bottom right")
	width, height := window.Size()
	window.SetCursorPosition(int(width), int(height))
	<-time.After(5 * time.Second)
}
