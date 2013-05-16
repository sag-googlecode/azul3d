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
	"fmt"
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
	screens := chippy.Screens()
	log.Printf("There are %d screens.\n", len(screens))
	log.Println("Default screen:", chippy.DefaultScreen())

	for i, screen := range screens {
		log.Printf("\nScreen %d - %s", i, screen)
	}

	fmt.Printf("Open window on screen: #")
	var screen int
	_, err = fmt.Scanln(&screen)
	if err != nil {
		log.Fatal(err)
	}

	if screen < 0 || screen > len(screens)-1 {
		log.Fatal("Incorrect screen number.")
	}
	chosenScreen := screens[screen]

	err = window.Open(chosenScreen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what it currently has property-wise
	log.Println(window)

	paintEvents := window.PaintEvents()
	closeEvents := window.CloseEvents()
	cursorPositionEvents := window.CursorPositionEvents()
	keyboardStateEvents := window.KeyboardStateEvents()
	keyboardTypedEvents := window.KeyboardTypedEvents()
	maximizedEvents := window.MaximizedEvents()
	minimizedEvents := window.MinimizedEvents()
	mouseEvents := window.MouseEvents()
	focusedEvents := window.FocusedEvents()
	positionEvents := window.PositionEvents()
	sizeEvents := window.SizeEvents()

	for {
		select {
		case <-paintEvents.Read:
			log.Println("paint")

		case <-closeEvents.Read:
			log.Println("close")

		case v := <-cursorPositionEvents.Read:
			log.Println("cursorPosition", v)

		case v := <-keyboardStateEvents.Read:
			log.Println("keyboard", v)

		case v := <-keyboardTypedEvents.Read:
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
		//log.Println(window)
	}
}
