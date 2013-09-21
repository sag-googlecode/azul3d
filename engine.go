// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package azul3d

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/clock"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/renderer"
	"code.google.com/p/azul3d/scene/util"
	"fmt"
	"log"
	"time"
)

// Engine is a base type to store a generic single-window Azul3D application.
type Engine struct {
	Renderer *scene.Node
	Window   *chippy.Window

	Scene3d *scene.Node
	Scene2d *scene.Node

	Camera3d *scene.Node
	Camera2d *scene.Node

	Clock *clock.Clock
}

// NewEngine returns a new intialized engine. Should any errors occur during
// initialization, log.Fatal() will be executed.
func NewEngine() *Engine {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	window := chippy.NewWindow()
	window.SetPositionCenter(chippy.DefaultScreen())

	err = window.Open(chippy.DefaultScreen())
	if err != nil {
		log.Fatal(err)
	}

	renderNode := scene.New("renderer")
	renderer.Create(renderNode, window)
	clock := renderer.Clock(renderNode)

	// We'll be using the OpenGL rendering backend
	err = renderer.SetBackend(renderNode, renderer.OpenGL)
	if err != nil {
		log.Fatal(err)
	}

	// Start an goroutine to display FPS
	go func() {
		for {
			// Sleep half a second
			time.Sleep(500 * time.Millisecond)

			if window.Destroyed() {
				break
			}

			// Display FPS
			window.SetTitle(fmt.Sprintf("Azul3D - %v FPS", clock.FrameRate()))
		}
	}()

	scene3d := renderNode.New("scene3d")
	scene2d := renderNode.New("scene2d")

	cam3d := scene3d.New("camera3d")
	camera.SetScene(cam3d, scene3d)

	cam2d := scene2d.New("camera2d")
	camera.SetScene(cam2d, scene2d)

	handleResized := func(w, h int) {
		width := math.Real(w)
		height := math.Real(h)

		camera.SetLens(cam3d, util.PerspectiveLens(75, width/height, 0.001, 1000))
		camera.SetLens(cam2d, util.OrthoLens(0, width, -height, 0, -1000000, 1000000))
	}

	w, h := window.Size()
	handleResized(w, h)

	var stop func()
	stop = event.Define(event.Handlers{
		"window-resized": func(e *event.Event) {
			ev := e.Data.(*chippy.ResizedEvent)
			handleResized(ev.Width, ev.Height)
		},

		"window-minimized": func(e *event.Event) {
			ev := e.Data.(*chippy.MinimizedEvent)
			if ev.Minimized {
				renderer.Pause(renderNode)
			} else {
				renderer.Play(renderNode)
			}
		},

		"window-close": func(e *event.Event) {
			renderer.Destroy(renderNode)

			// Exit main loop
			Exit()

			// Stop event handlers
			stop()
		},
	})

	return &Engine{
		Window:   window,
		Renderer: renderNode,
		Scene3d:  scene3d,
		Scene2d:  scene2d,
		Camera3d: cam3d,
		Camera2d: cam2d,
		Clock:    clock,
	}
}
