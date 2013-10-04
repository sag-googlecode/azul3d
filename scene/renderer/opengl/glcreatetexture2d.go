// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/native/opengl/1.0"
	"code.google.com/p/azul3d/scene/texture"
	"image"
	"runtime"
	"unsafe"
)

func (r *Renderer) createTexture2D(t *texture.Texture2D) {
	img := t.Image()

	// Create texture identity
	var ident uint32
	r.gl.GenTextures(1, &ident)
	t.SetNativeIdentity(ident)

	// Select the texture
	r.gl.BindTexture(opengl.TEXTURE_2D, ident)
	defer r.gl.BindTexture(opengl.TEXTURE_2D, 0)

	// Setup wrap modes
	var glWrapU, glWrapV int32

	// Find U wrap mode
	uWrapMode := t.WrapModeU()
	switch uWrapMode {
	case texture.Repeat:
		glWrapU = opengl.REPEAT
	case texture.Clamp:
		glWrapU = opengl.CLAMP_TO_EDGE
	case texture.BorderColor:
		glWrapU = opengl.CLAMP_TO_BORDER
	case texture.Mirror:
		glWrapU = opengl.MIRRORED_REPEAT
	}

	// Find V wrap mode
	vWrapMode := t.WrapModeV()
	switch vWrapMode {
	case texture.Repeat:
		glWrapV = opengl.REPEAT
	case texture.Clamp:
		glWrapV = opengl.CLAMP_TO_EDGE
	case texture.BorderColor:
		glWrapV = opengl.CLAMP_TO_BORDER
	case texture.Mirror:
		glWrapV = opengl.MIRRORED_REPEAT
	}

	if uWrapMode == texture.BorderColor || vWrapMode == texture.BorderColor {
		// If either wrap mode is BorderColor, we need to specify the actual
		// border color.
		borderColor := t.BorderColorFloat32()
		r.gl.TexParameterfv(opengl.TEXTURE_2D, opengl.TEXTURE_BORDER_COLOR, &borderColor[0])
	}

	r.gl.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_WRAP_S, glWrapU)
	r.gl.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_WRAP_T, glWrapV)

	// Setup min/mag filters
	var glMinFilter, glMagFilter int32

	minFilter := t.MinFilter()
	switch minFilter {
	case texture.Nearest:
		glMinFilter = opengl.NEAREST
	case texture.Linear:
		glMinFilter = opengl.LINEAR
	case texture.NearestMipmapNearest:
		glMinFilter = opengl.NEAREST_MIPMAP_NEAREST
	case texture.LinearMipmapNearest:
		glMinFilter = opengl.LINEAR_MIPMAP_NEAREST
	case texture.NearestMipmapLinear:
		glMinFilter = opengl.NEAREST_MIPMAP_LINEAR
	case texture.LinearMipmapLinear:
		glMinFilter = opengl.LINEAR_MIPMAP_LINEAR
	}

	magFilter := t.MagFilter()
	switch magFilter {
	case texture.Nearest:
		glMagFilter = opengl.NEAREST
	case texture.Linear:
		glMagFilter = opengl.LINEAR
	case texture.NearestMipmapNearest:
		glMagFilter = opengl.NEAREST_MIPMAP_NEAREST
	case texture.LinearMipmapNearest:
		glMagFilter = opengl.LINEAR_MIPMAP_NEAREST
	case texture.NearestMipmapLinear:
		glMagFilter = opengl.NEAREST_MIPMAP_LINEAR
	case texture.LinearMipmapLinear:
		glMagFilter = opengl.LINEAR_MIPMAP_LINEAR
	}

	r.gl.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_MIN_FILTER, glMinFilter)
	r.gl.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_MAG_FILTER, glMagFilter)

	// Enable mipmap generation if either filter is mipmapped.
	if minFilter.Mipmapped() || magFilter.Mipmapped() {
		r.gl.TexParameteri(opengl.TEXTURE_2D, opengl.GENERATE_MIPMAP, opengl.TRUE)
	}

	// Determine formatting, load texture into OpenGL
	internalFormat := int32(opengl.RGBA)
	if t.Compressed() {
		internalFormat = opengl.COMPRESSED_RGBA
	}

	// Load texture
	sz := t.Bounds().Size()

	t.RLock()
	img = t.Source.(*image.RGBA)
	r.gl.TexImage2D(opengl.TEXTURE_2D, 0, internalFormat, int32(sz.X), int32(sz.Y), 0, opengl.RGBA, opengl.UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	t.RUnlock()

	runtime.SetFinalizer(t, func(t *texture.Texture2D) {
		r.texturesToFreeAccess.Lock()
		defer r.texturesToFreeAccess.Unlock()

		r.texturesToFree = append(r.texturesToFree, t.NativeIdentity().(uint32))
	})
}
