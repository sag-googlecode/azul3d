// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build tests

// Test - Opens two windows, changes their title properties.
package main

import (
	"azul3d.org/v1/chippy"
	"log"
	"os"
	"time"
)

func program() {
	defer chippy.Exit()

	window1 := chippy.NewWindow()
	window2 := chippy.NewWindow()

	window1.SetTitle("Window 1 - Hello 世界")

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

	window2.SetTitle("Window 2 - Hello 世界")

	// Print out what they currently has property-wise
	log.Println(window1)
	log.Println(window2)

	log.Println("Waiting 5 seconds...")
	<-time.After(5 * time.Second)

	window1.SetTitle("Hello after 5 seconds.. what an nice 世界")
	window2.SetTitle("Hello after 5 seconds.. what an nice 世界")

	// Just wait an while so they can enjoy the window
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
