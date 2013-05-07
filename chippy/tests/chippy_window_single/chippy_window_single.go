// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an single window on the specified screen
package main

// Note: On Windows build with:
//   go install -ldflags "-H windowsgui" path/to/pkg
// to hide the command prompt

import (
	"code.google.com/p/azul3d/chippy"
	"log"
	"os"
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

	// Print out what it currently has property-wise
	log.Println(window)

	closeEvents := window.CloseEvents()
	cursorPositionEvents := window.CursorPositionEvents()
	keyboardEvents := window.KeyboardEvents()
	maximizedEvents := window.MaximizedEvents()
	minimizedEvents := window.MinimizedEvents()
	mouseEvents := window.MouseEvents()
	focusedEvents := window.FocusedEvents()
	positionEvents := window.PositionEvents()
	sizeEvents := window.SizeEvents()

	for {
		select {
		case v := <-closeEvents.Read:
			log.Println("close", v)

		case v := <-cursorPositionEvents.Read:
			log.Println("cursorPosition", v)

		case v := <-keyboardEvents.Read:
			log.Println("keyboard", v)

		case v := <-maximizedEvents.Read:
			log.Println("maximized", v)

		case v := <-minimizedEvents.Read:
			log.Println("minimized", v)

		case v := <-mouseEvents.Read:
			log.Println("mouse", v)

		case v := <-focusedEvents.Read:
			log.Println("focused", v)

		case v := <-positionEvents.Read:
			log.Println("position", v)

		case v := <-sizeEvents.Read:
			log.Println("size", v)
		}
		log.Println(window)
	}
}
