// +build examples

package main

import (
	"code.google.com/p/azul3d"
	"code.google.com/p/azul3d/scene/ui"
	"code.google.com/p/azul3d/math"
	_ "image/png"
	"log"
	"os"
)

var (
	// Create the engine.
	engine = azul3d.NewEngine()
)

func program() {
	panel := ui.New("myPanel")
	ui.Apply(panel, ui.StateOptions{
		ui.Default: ui.Options{
			ui.Parent: engine.Scene2d,
			ui.Width:  256,
			ui.Height: 256,
			ui.Layout: ui.Vertical,
			ui.Color: math.Vector4(1, 0, 0, 1),
			ui.ColorScale: math.Vector4(0.5, 0.5, 0.5, 1),
		},
		ui.Hover: ui.Options{
			ui.Width: 100,
		},
		ui.Click: ui.Options{
			ui.Height: 100,
			ui.Layout: ui.Horizontal,
		},
		ui.Active: ui.Options{
			ui.Width:  150,
			ui.Height: 150,
		},
	})

	button := ui.New("myButton")
	ui.Apply(button, ui.StateOptions{
		ui.Default: ui.Options{
			ui.Parent: panel,
			ui.Width: 100,
			ui.Height: 100,
			ui.MarginLeft: 2,
			ui.MarginRight: 2,
			ui.MarginBottom: 2,
			ui.MarginTop: 2,
			ui.Layout: ui.Horizontal,
			ui.Color: math.Vector4(0, 0, 1, 1),
		},
		ui.Hover: ui.Options{
			ui.ColorScale: math.Vector4(0, 0, 0.8, 1),
		},
		ui.Click: ui.Options{
			ui.ColorScale: math.Vector4(0, 0, 0.6, 1),
		},
		ui.Active: ui.Options{
			ui.ColorScale: math.Vector4(0, 0, 0.5, 1),
		},
	})
	button.SetPos(128, 0, 0)
}

func main() {
	// For debugging anything
	azul3d.SetDebugOutput(os.Stdout)

	// Initialize azul3d
	err := azul3d.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Launch program
	go program()

	// Enter main loop
	azul3d.MainLoop()
}
