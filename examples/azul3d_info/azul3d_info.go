package main

import (
	"code.google.com/p/azul3d"
	"code.google.com/p/azul3d/scene/renderer"
	"log"
)

func program() {
	// Hide window
	azul3d.Window.SetVisible(false)

	log.Println("GPU Name:", renderer.GPUName(azul3d.Renderer))
	log.Println("GPU Vendor:", renderer.GPUVendor(azul3d.Renderer))
	log.Println("GPU Driver Version:", renderer.GPUDriverVersion(azul3d.Renderer))

	log.Println("Max Texture Size:", renderer.MaxTextureSize(azul3d.Renderer))

	glMajor, glMinor := renderer.GLVersion(azul3d.Renderer)
	log.Printf("OpenGL Version %d.%d\n", glMajor, glMinor)

	glslMajor, glslMinor := renderer.GLSLVersion(azul3d.Renderer)
	log.Printf("GLSL Version %d.%d\n", glslMajor, glslMinor)

	log.Println("OpenGL Extensions:", renderer.GLExtensions(azul3d.Renderer))

	// Quit right away instead of waiting for window to be closed
	azul3d.Exit()
}

func main() {
	// Run our program, enter main loop.
	azul3d.Run(program)
}
