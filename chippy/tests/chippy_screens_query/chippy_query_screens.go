// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Queries for all available screens, lists and changes screen modes
package main

import (
	"code.google.com/p/azul3d/chippy"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(0)

	// Enable debug messages
	chippy.SetDebugOutput(os.Stdout)

	err := chippy.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer chippy.Destroy()

	screens := chippy.Screens()
	log.Printf("There are %d screens.\n", len(screens))
	log.Println("Default screen:", chippy.DefaultScreen())

	for i, screen := range screens {
		log.Printf("\nScreen %d - %s", i, screen)

		for i, mode := range screen.ScreenModes() {
			log.Printf("    -> ScreenMode %d - %s", i, mode)
		}
	}

	fmt.Printf("Change Screen: #")
	var screen int
	_, err = fmt.Scanln(&screen)
	if err != nil {
		log.Fatal(err)
	}

	if screen < 0 || screen > len(screens)-1 {
		log.Fatal("Incorrect screen number.")
	}
	chosenScreen := screens[screen]

	fmt.Printf("Change Screen #%d to mode: #", screen)
	var mode int
	_, err = fmt.Scanln(&mode)
	if err != nil {
		log.Fatal(err)
	}

	if mode < 0 || mode > len(screens[screen].ScreenModes())-1 {
		log.Fatal("Incorrect screen number.")
	}
	chosenMode := chosenScreen.ScreenModes()[mode]

	// Ensure that we restore the screen to it's original state before exiting
	defer chosenScreen.Restore()

	// Change screen mode
	chosenScreen.SetScreenMode(chosenMode)

	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)

	// Disable this with chippy.SetAutoRestoreOriginalScreenMode(false)
	log.Println("Original resolution should restore now")
}
