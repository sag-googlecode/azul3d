// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build tests

// Test - Opens two windows, requests each one notify the user of an event.
package main

import (
	"azul3d.org/v1/chippy"
	"log"
	"os"
	"time"
)

func program() {
	defer chippy.Exit()

	window1 := chippy.NewWindow()
	window2 := chippy.NewWindow()

	window1.SetTitle("Window 1")
	window2.SetTitle("Window 2")

	window1.SetMinimized(true)
	window2.SetMinimized(true)

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err := window1.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	err = window2.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// In case their window manager ignores minimized requests (not common).
	log.Println("Please manually minimize the windows if they are not already..")
	log.Println("Waiting 5 seconds...")
	<-time.After(5 * time.Second)

	window1.Notify()
	window2.Notify()

	// Inform user of their notifications.
	log.Println("Both windows have been notified. This may appear as blinking or restoring the window.")

	// Just wait an while so they can enjoy the window
	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)
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
