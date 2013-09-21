// +build examples

package main

import(
	"code.google.com/p/azul3d"
	"code.google.com/p/azul3d/scene/renderer"
	"log"
	"os"
)

func init() {
	azul3d.SetDebugOutput(os.Stdout)
}

var(
	// Create the engine.
	engine = azul3d.NewEngine()
)

func program() {
	// Hide window
	engine.Window.SetVisible(false)

	gpuName := renderer.GPUName(engine.Renderer)
	maxTextureLayers := renderer.MaxTextureLayers(engine.Renderer)
	maxTextureCoords := renderer.MaxTextureCoords(engine.Renderer)
	maxTextureSize := renderer.MaxTextureSize(engine.Renderer)

	log.Println("GPU:", gpuName)
	log.Println("Max texture layers:", maxTextureLayers)
	log.Println("Max texture coordinates:", maxTextureCoords)
	log.Println("Max texture size:", maxTextureSize)

	// Quit right away instead of waiting for window to be closed
	azul3d.Exit()
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

