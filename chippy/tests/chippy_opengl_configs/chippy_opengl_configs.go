// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build tests,!no_opengl

// Test - Opens an window displays all possible configurations
package main

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

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what the window currently has property-wise
	log.Println(window)

	// Choose an buffer format, these include things like double buffering, bytes per pixel, number of depth bits, etc.
	configs := window.GLConfigs()

	for _, config := range configs {
		log.Println(config)
	}
}
