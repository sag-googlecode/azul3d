// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an single window, changes it's fullscreen property
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

	window.SetFullscreen(true)

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what they currently has property-wise
	log.Println(window)

	log.Println("Waiting 10 seconds...")
	<-time.After(10 * time.Second)
	window.SetFullscreen(false)

	log.Println("Waiting 10 seconds...")
	<-time.After(10 * time.Second)
	window.SetFullscreen(true)

	log.Println("Waiting 5 seconds...")
	<-time.After(5 * time.Second)
	window.SetFullscreen(false)

	// Just wait an while so they can enjoy the window
	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)
}
