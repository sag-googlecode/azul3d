// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/native/opengl/1.0"
)

func (r *GLFFRenderer) GPUName() string {
	if len(r.gpuName) > 0 {
		return r.gpuName
	}
	r.gpuName = r.gl.GetString(opengl.RENDERER)
	return r.gpuName
}

func (r *GLFFRenderer) MaxTextureCoords() int {
	if r.maxTextureCoords != 0 {
		return r.maxTextureCoords
	}

	var maxTextureCoords int32
	r.gl.GetIntegerv(opengl.MAX_TEXTURE_COORDS, &maxTextureCoords)
	r.maxTextureCoords = int(maxTextureCoords)
	return r.maxTextureCoords
}

func (r *GLFFRenderer) MaxTextureLayers() int {
	if r.maxTextureLayers != 0 {
		return r.maxTextureLayers
	}

	var maxTextureLayers int32
	r.gl.GetIntegerv(opengl.MAX_TEXTURE_UNITS, &maxTextureLayers)
	r.maxTextureLayers = int(maxTextureLayers)
	return r.maxTextureLayers
}

func (r *GLFFRenderer) MaxTextureSize() int {
	if r.maxTextureSize != 0 {
		return r.maxTextureSize
	}

	var maxTextureSize int32
	r.gl.GetIntegerv(opengl.MAX_TEXTURE_SIZE, &maxTextureSize)
	r.maxTextureSize = int(maxTextureSize)
	return r.maxTextureSize
}
