// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens two windows, changes each of their icon properties
package main

import (
	"code.google.com/p/azul3d/chippy"
	_ "image/png"
	"image"
	"time"
	"log"
	"os"
)

func main() {
	var err error

	log.SetFlags(0)

	// Load the image that we'll use for the window icon
	file, err := os.Open("data/icon_128x128.png")
	if err != nil {
		log.Fatal(err)
	}

	icon, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	err = chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer chippy.Destroy()

	window1 := chippy.NewWindow()
	window2 := chippy.NewWindow()

	window1.SetIcon(icon)

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window1.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	err = window2.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	window1.SetIcon(icon)
	window2.SetIcon(icon)

	// Print out what they currently has property-wise
	log.Println(window1)
	log.Println(window2)

	log.Println("Waiting 5 seconds...")
	<-time.After(5 * time.Second)

	window1.SetIcon(nil)
	window2.SetIcon(nil)

	// Just wait an while so they can enjoy the window
	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)
}
