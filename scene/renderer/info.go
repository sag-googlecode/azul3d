// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package renderer

import (
	"code.google.com/p/azul3d/scene"
)

// MaxTextureSize returns the maximum dimension (width/height) that a texture
// may be for use with the renderer.
func MaxTextureSize(n *scene.Node) int {
	o := mustGetObj(n)
	return o.maxTextureSize
}

// GPUName returns the name of the GPU in use by this renderer.
func GPUName(n *scene.Node) string {
	o := mustGetObj(n)
	return o.gpuName
}

// GPUVendor returns the vendor name of the GPU in use by this renderer.
func GPUVendor(n *scene.Node) string {
	o := mustGetObj(n)
	return o.gpuVendor
}

// GPUDriverVersion returns the driver version of the GPU in use by this
// renderer, or zero if it is not known.
func GPUDriverVersion(n *scene.Node) int {
	o := mustGetObj(n)
	return o.gpuDriverVersion
}

// GLExtensions returns a slice of all the OpenGL extensions supported by the
// OpenGL implementation.
func GLExtensions(n *scene.Node) []string {
	o := mustGetObj(n)
	return o.glExtensions
}

// GLVersion returns the OpenGL version supported and in use by this renderer.
func GLVersion(n *scene.Node) (major, minor int) {
	o := mustGetObj(n)
	return o.glMajorVersion, o.glMinorVersion
}

// GLSLVersion returns the GLSL version supported and in use by this renderer.
func GLSLVersion(n *scene.Node) (major, minor int) {
	o := mustGetObj(n)
	return o.glslMajorVersion, o.glslMinorVersion
}
