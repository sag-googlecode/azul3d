// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Loads and renders a tmx map file.
package main

import (
	"code.google.com/p/azul3d/engine"
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/tmx/tmxmesh"
	_ "image/png"
	"flag"
	"log"
)

// Event handler to manipulate the camera when cursor is moved
func onCursorPosition(ev *event.Event) {
	pos := ev.Data.(*chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform the camera.
	if !engine.Window.CursorGrabbed() {
		return
	}

	// Apply relative mouse movement to the 2D camera
	x := math.Real(pos.X) * 0.1
	z := math.Real(-pos.Y) * 0.1
	engine.Camera2d.SetRelativePos(engine.Camera2d, x.Rounded(), 0, z.Rounded())
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")
	engine.Camera2d.ResetTransform()
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed(ev *event.Event) {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}

var(
	defaultMapFile = "src/code.google.com/p/azul3d/tmx/testdata/test_csv_tsx.tmx"
	mapFile = flag.String("file", defaultMapFile, "tmx map file to load")
)

func init() {
	flag.Parse()
}

func program() {
	tmxMap, err := tmxmesh.LoadFile(*mapFile, nil)
	if err != nil {
		log.Fatal(err)
	}

	tmxMap.SetParent(engine.Scene2d)

	// Grab the cursor
	engine.Window.SetCursorGrabbed(true)

	var stop func()
	stop = event.Define(event.Handlers{
		// Listen for alt keys to toggle cursor grabbed
		"RightAlt": toggleCursorGrabbed,
		"LeftAlt":  toggleCursorGrabbed,

		// Listen for R key to reset transformations
		"R": resetTransforms,

		// Listen for cursor position to move the camera
		"cursor-position": onCursorPosition,

		"window-destroyed": func(ev *event.Event) {
			stop()
		},
	})
}

func main() {
	// Run our program, enter main loop.
	engine.Run(program)
}
