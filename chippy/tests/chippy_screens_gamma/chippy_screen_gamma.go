// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Changes screen gamma ramp (brightness, contrast, gamma), restores old one
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
	}

	log.Printf("\nEnter screen: #")
	var screen int
	_, err = fmt.Scanln(&screen)
	if err != nil {
		log.Fatal(err)
	}

	if screen < 0 || screen > len(screens)-1 {
		log.Fatal("Incorrect screen number.")
	}
	chosenScreen := screens[screen]

	if chosenScreen.GammaRampSize() == 0 {
		log.Fatal("Chosen screen has no support for gamma ramps (color correction lookup table / LUT)")
	}

	// Ensure that we restore the screen to it's original state. (restores the gamma ramp)
	defer chosenScreen.Restore()

	do := func(red, green, blue float32) {
		log.Printf("r = %.2f, g = %.2f, b = %.2f\n", red, green, blue)

		// First create an gamma ramp
		myGammaRamp := new(chippy.GammaRamp)

		// Now intialize the ramp based off ramp size, intensity
		gammaRampSize := chosenScreen.GammaRampSize()

		// We could just as easilly use our own LUT calculations here -- see gammaramp.go for more
		// information.
		myGammaRamp.SetAsLinearIntensity(gammaRampSize, red, green, blue)

		// Now actually tell the screen to use it
		err := chosenScreen.SetGammaRamp(myGammaRamp)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("\nNormal")
	do(1.0, 1.0, 1.0)
	<-time.After(3 * time.Second)

	// Stuck on the wrong settings? uncomment this line and you'll get reset to normal here
	// return

	log.Println("All")
	for v := float32(0); v < 1.0; v += (1.0 / 30.0) {
		do(v, v, v)
		<-time.After(60 * time.Millisecond)
	}

	log.Println("Red")
	for r := float32(0); r < 1.0; r += (1.0 / 30.0) {
		do(r, 1.0, 1.0)
		<-time.After(60 * time.Millisecond)
	}

	log.Println("Green")
	for g := float32(0); g < 1.0; g += (1.0 / 30.0) {
		do(1.0, g, 1.0)
		<-time.After(60 * time.Millisecond)
	}

	log.Println("Blue")
	for b := float32(0); b < 1.0; b += (1.0 / 30.0) {
		do(1.0, 1.0, b)
		<-time.After(60 * time.Millisecond)
	}

	do = func(brightness, contrast, gamma float32) {
		log.Printf("brightness = %.2f, contrast = %.2f, gamma = %.2f\n", brightness, contrast, gamma)

		// First create an gamma ramp
		myGammaRamp := new(chippy.GammaRamp)

		// Now intialize the ramp based off ramp size, intensity
		gammaRampSize := chosenScreen.GammaRampSize()

		// We could just as easilly use our own LUT calculations here -- see gammaramp.go for more
		// information.
		myGammaRamp.SetAsBrightnessContrastGamma(gammaRampSize, brightness, contrast, gamma)

		// Now actually tell the screen to use it
		err := chosenScreen.SetGammaRamp(myGammaRamp)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("\nNormal")
	do(1.0, 1.0, 2.2)
	<-time.After(3 * time.Second)

	// Stuck on the wrong settings? uncomment this line and you'll get reset to normal here
	// return

	log.Println("Brightness")
	for v := float32(-1); v < 1.0; v += (1.0 / 30.0) {
		do(v, 1.0, 2.2)
		<-time.After(60 * time.Millisecond)
	}

	log.Println("Contrast")
	for v := float32(0); v < 5.0; v += (1.0 / 30.0) {
		do(0, v, 2.2)
		<-time.After(60 * time.Millisecond)
	}

	log.Println("Gamma")
	for v := float32(0); v < 3.0; v += (1.0 / 30.0) {
		do(0, 1.0, v)
		<-time.After(60 * time.Millisecond)
	}

	// Uncomment this line and the gamma will stay active on the screen after the program exits
	//chippy.SetAutoRestoreOriginalGammaRamp(false)
}
