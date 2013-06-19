// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an window and grabs the cursor
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

	cursorPositionEvents := window.CursorPositionEvents()
	closeEvents := window.CloseEvents()
	for {
		select {
		case v := <-cursorPositionEvents.Read:
			log.Printf("Grabbed? %v | Position: %v", window.CursorGrabbed(), v)

		case <-closeEvents.Read:
			return
		}
	}
}
