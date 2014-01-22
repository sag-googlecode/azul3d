// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Opens an window and grabs the cursor.
package main

import (
	"azul3d.org/chippy"
	"log"
	"os"
	"time"
)

func program() {
	defer chippy.Exit()

	window := chippy.NewWindow()

	// Actually open the window
	screen := chippy.DefaultScreen()
	err := window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	window.SetCursorGrabbed(true)

	// Print out what they currently has property-wise
	log.Println(window)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			isGrabbed := window.CursorGrabbed()
			window.SetCursorGrabbed(!isGrabbed)
			log.Printf("Grabbed? %v\n", window.CursorGrabbed())
		}
	}()

	events := window.Events()
	defer window.CloseEvents(events)

	for {
		e := <-events

		switch e.(type) {
		case *chippy.CursorPositionEvent:
			log.Printf("Grabbed? %v | %v", window.CursorGrabbed(), e)

		case *chippy.CloseEvent:
			return

		default:
			// We don't care about whatever event this is.
			break
		}
	}
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
