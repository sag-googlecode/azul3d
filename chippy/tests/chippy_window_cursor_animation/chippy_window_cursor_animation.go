// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an windows, animates the cursor property
package main

import (
	"code.google.com/p/azul3d/chippy"
	_ "image/png"
	"image"
	"time"
	"log"
	"fmt"
	"os"
)

func main() {
	var err error

	log.SetFlags(0)

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	err = chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer chippy.Destroy()

	window := chippy.NewWindow()
	window.SetTransparent(true)

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	var cursors []*chippy.Cursor
	for i := 1; i < 25; i++ {
		// Load the image frame that we'll use for the animated cursor
		file, err := os.Open(fmt.Sprintf("data/loading/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}

		cursorImage, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		cursor := &chippy.Cursor{
			Image: cursorImage,
			X: 16,
			Y: 16,
		}
		window.PrepareCursor(cursor)
		cursors = append(cursors, cursor)
	}

	// Those cursors use up memory (and are cached for speed), if we wanted to stop using the
	// animated cursor, we could do the following:
	//for _, cursor := range cursors {
	//    window.FreeCursor(cursor)
	//}

	// Print out what they currently has property-wise
	log.Println(window)

	frame := 0
	for !window.Destroyed() {
		// Play back at 24 FPS
		time.Sleep((1000 / 24) * time.Millisecond)

		cursor := cursors[frame]
		window.SetCursor(cursor)
		frame += 1
		if frame >= 24 {
			frame = 0
		}
	}
}
