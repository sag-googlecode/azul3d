package main

import (
	"code.google.com/p/azul3d"
	"code.google.com/p/azul3d/scene/renderer"
	"log"
)

func program() {
	// Hide window
	azul3d.Window.SetVisible(false)

	gpuName := renderer.GPUName(azul3d.Renderer)
	maxTextureLayers := renderer.MaxTextureLayers(azul3d.Renderer)
	maxTextureCoords := renderer.MaxTextureCoords(azul3d.Renderer)
	maxTextureSize := renderer.MaxTextureSize(azul3d.Renderer)

	log.Println("GPU:", gpuName)
	log.Println("Max texture layers:", maxTextureLayers)
	log.Println("Max texture coordinates:", maxTextureCoords)
	log.Println("Max texture size:", maxTextureSize)

	// Quit right away instead of waiting for window to be closed
	azul3d.Exit()
}

func main() {
	// Run our program, enter main loop.
	azul3d.Run(program)
}
