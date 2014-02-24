// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	opengl "azul3d.org/native/gl"
)

func (r *Renderer) GPUName() string {
	if len(r.gpuName) > 0 {
		return r.gpuName
	}
	r.gpuName = opengl.String(r.gl.GetString(opengl.RENDERER))
	return r.gpuName
}

func (r *Renderer) GPUVendor() string {
	if len(r.gpuVendorName) > 0 {
		return r.gpuVendorName
	}
	r.gpuVendorName = opengl.String(r.gl.GetString(opengl.VENDOR))
	return r.gpuVendorName
}

func (r *Renderer) GPUDriverVersion() int {
	_, _, release := r.gl.Version()
	return release
}

func (r *Renderer) GLExtensions() []string {
	return r.gl.Extensions()
}

func (r *Renderer) GLVersion() (major, minor int) {
	major, minor, _ = r.gl.Version()
	return
}

func (r *Renderer) GLSLVersion() (major, minor int) {
	major, minor, _ = r.gl.ShaderVersion()
	return
}

func (r *Renderer) MaxTextureSize() int {
	if r.maxTextureSize != 0 {
		return r.maxTextureSize
	}

	var maxTextureSize int32
	r.gl.GetIntegerv(opengl.MAX_TEXTURE_SIZE, &maxTextureSize)
	r.gl.Execute()
	r.maxTextureSize = int(maxTextureSize)
	return r.maxTextureSize
}

func (r *Renderer) GLSLMaxVaryingFloats() int {
	if r.glslMaxVaryingFloats != 0 {
		return r.glslMaxVaryingFloats
	}

	var glslMaxVaryingFloats int32
	r.gl.GetIntegerv(opengl.MAX_VARYING_FLOATS, &glslMaxVaryingFloats)
	r.gl.Execute()
	r.glslMaxVaryingFloats = int(glslMaxVaryingFloats)
	return r.glslMaxVaryingFloats
}

func (r *Renderer) GLSLMaxVertexShaderInputs() int {
	if r.glslMaxVertexShaderInputs != 0 {
		return r.glslMaxVertexShaderInputs
	}

	var glslMaxVertexShaderInputs int32
	r.gl.GetIntegerv(opengl.MAX_VERTEX_UNIFORM_COMPONENTS, &glslMaxVertexShaderInputs)
	r.gl.Execute()
	r.glslMaxVertexShaderInputs = int(glslMaxVertexShaderInputs)
	return r.glslMaxVertexShaderInputs
}

func (r *Renderer) GLSLMaxFragmentShaderInputs() int {
	if r.glslMaxFragmentShaderInputs != 0 {
		return r.glslMaxFragmentShaderInputs
	}

	var glslMaxFragmentShaderInputs int32
	r.gl.GetIntegerv(opengl.MAX_FRAGMENT_UNIFORM_COMPONENTS, &glslMaxFragmentShaderInputs)
	r.glslMaxFragmentShaderInputs = int(glslMaxFragmentShaderInputs)
	return r.glslMaxFragmentShaderInputs
}
