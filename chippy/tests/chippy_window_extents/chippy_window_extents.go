// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Prints the extents of an window's region.
package main

import (
	"code.google.com/p/azul3d/chippy"
	"log"
	"os"
)

func program() {
	defer chippy.Exit()

	window := chippy.NewWindow()
	err := window.Open(chippy.DefaultScreen())
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	// Print what the window extents are
	log.Println(window.Extents())
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
