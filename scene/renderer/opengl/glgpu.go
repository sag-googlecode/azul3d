// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/native/gl"
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
	r.maxTextureSize = int(maxTextureSize)
	return r.maxTextureSize
}
