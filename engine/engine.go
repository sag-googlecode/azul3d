// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package engine is a thin wrapper around the various azul3d packages.
//
// This package is a very thin wrapper around the various other sub-packages
// found inside this package directory.
//
// Serious applications will not make use of this package, but instead use it
// as a reference.
package engine

import (
	"azul3d.org/v0/chippy"
	"azul3d.org/v0/clock"
	"azul3d.org/v0/event"
	"azul3d.org/v0/scene"
	"azul3d.org/v0/scene/camera"
	"azul3d.org/v0/scene/renderer"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Renderer *scene.Node
	Window   *chippy.Window

	Scene3d *scene.Node
	Scene2d *scene.Node

	Camera3d *scene.Node
	Camera2d *scene.Node

	Clock *clock.Clock

	// The default clock stats.
	Stats = clock.NewStats()
)

func init() {
	Stats.SetEnabled(false)
	chippy.SetDebugOutput(os.Stdout)

	err := chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func setup() {
	Window = chippy.NewWindow()
	Window.SetPositionCenter(chippy.DefaultScreen())

	err := Window.Open(chippy.DefaultScreen())
	if err != nil {
		log.Fatal(err)
	}

	Renderer = scene.New("renderer")
	err = renderer.Create(Renderer, Window, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	Clock = renderer.Clock(Renderer)

	// Start an goroutine to display FPS
	go func() {
		for {
			// Sleep half a second
			time.Sleep(500 * time.Millisecond)

			if Window.Destroyed() {
				break
			}

			// Display FPS
			Window.SetTitle(fmt.Sprintf("Azul3D - %v FPS", Clock.FrameRate()))
		}
	}()

	Scene3d = Renderer.New("scene3d")
	Scene2d = Renderer.New("scene2d")

	// Our 3D camera, which will display the 3D scene
	Camera3d = Scene3d.New("camera3d")
	camera.SetScene(Camera3d, Scene3d)

	// The camera must have a region (a square area describing where the camera
	// will render to on the window).
	//
	// In our case, we will use the region [0, 0, 1, 1] which covers the entire
	// window.
	underlay := camera.NewRegion(0, 0, 1, 1)
	camera.AddRegion(Camera3d, underlay)

	// Our 2D camera, which will display the 2D scene (UI elements.. etc).
	Camera2d = Scene2d.New("camera2d")
	camera.SetScene(Camera2d, Scene2d)

	// For the 2D camera, it is 'overlay'd atop the 3D camera's region, so we
	// want to make sure that this region does not clear the color, depth, or
	// stencil buffers below it (the already-rendered 3D scene).
	overlay := camera.NewRegion(0, 0, 1, 1)
	overlay.SetClearColorActive(false)
	overlay.SetClearDepthActive(false)
	overlay.SetClearStencilActive(false)

	// The sort value is any number, since the sort value for the 'overlay'
	// region is 100, and the sort value for the 'underlay' region is 0
	// (value by default), the 'overlay' region will be drawn last.
	overlay.SetSort(100)
	camera.AddRegion(Camera2d, overlay)

	handleResized := func(w, h int) {
		width := float64(w)
		height := float64(h)

		camera.SetLens(Camera3d, camera.PerspectiveLens(75, width/height, 0.001, 1000))
		camera.SetLens(Camera2d, camera.OrthoLens(0, width, -height, 0, -1000000, 1000000))
	}

	w, h := Window.Size()
	handleResized(w, h)

	var stop func()
	stop = event.Define(event.Handlers{
		"window-resized": func(e *event.Event) {
			ev := e.Data.(chippy.ResizedEvent)
			handleResized(ev.Width, ev.Height)
		},

		"window-minimized": func(e *event.Event) {
			ev := e.Data.(chippy.MinimizedEvent)
			if ev.Minimized {
				renderer.Pause(Renderer)
			} else {
				renderer.Play(Renderer)
			}
		},

		"window-close": func(e *event.Event) {
			renderer.Destroy(Renderer)

			// Exit main loop
			Exit()

			// Stop event handlers
			stop()
		},
	})
}

// Exit exits the main loop.
func Exit() {
	chippy.Exit()
}

// Run initializes the public variables in this package, invokes the
// afterInit() function, and then enters the main loop.
//
// Because this function enters the main loop, it *must* be called on the main
// thread (due to some restrictions some platforms place on us).
//
// It is best to call this function inside your main() or init() function.
func Run(afterInit func()) {
	go func() {
		setup()
		afterInit()
	}()
	chippy.MainLoop()
}
