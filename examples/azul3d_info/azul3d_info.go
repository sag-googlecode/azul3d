package main

import (
	"azul3d.org/engine"
	"azul3d.org/scene/renderer"
	"log"
)

func program() {
	// Hide window
	engine.Window.SetVisible(false)

	log.Println("GPU Name:", renderer.GPUName(engine.Renderer))
	log.Println("GPU Vendor:", renderer.GPUVendor(engine.Renderer))
	log.Println("GPU Driver Version:", renderer.GPUDriverVersion(engine.Renderer))

	log.Println("Max Texture Size:", renderer.MaxTextureSize(engine.Renderer))

	glMajor, glMinor := renderer.GLVersion(engine.Renderer)
	log.Printf("OpenGL Version %d.%d\n", glMajor, glMinor)

	glslMajor, glslMinor := renderer.GLSLVersion(engine.Renderer)
	log.Printf("GLSL Version %d.%d\n", glslMajor, glslMinor)

	log.Println("OpenGL Extensions:", renderer.GLExtensions(engine.Renderer))

	// Quit right away instead of waiting for window to be closed
	engine.Exit()
}

func main() {
	// Run our program, enter main loop.
	engine.Run(program)
}
