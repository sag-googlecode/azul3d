// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Opens two windows, changes their aspect ratio properties.
package main

import (
	"code.google.com/p/azul3d/chippy"
	"log"
	"os"
	"time"
)

func program() {
	defer chippy.Exit()

	window1 := chippy.NewWindow()
	window2 := chippy.NewWindow()

	window1.SetTitle("Window 1")
	window2.SetTitle("Window 2")

	window1.SetAspectRatio(480.0 / 640.0) // 0.75 aspect ratio

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err := window1.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	err = window2.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	window2.SetAspectRatio(640.0 / 480.0) // 1.333 aspect ratio

	// Print out what they currently has property-wise
	log.Println(window1)
	log.Println(window2)

	// Just wait an while so they can enjoy the window
	log.Println("Waiting 30 seconds...")
	<-time.After(30 * time.Second)
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
