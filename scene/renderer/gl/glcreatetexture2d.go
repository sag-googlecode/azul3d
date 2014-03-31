// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	opengl "azul3d.org/v1/native/gl"
	"azul3d.org/v1/scene/texture"
	"image"
	"runtime"
	"unsafe"
)

func (r *Renderer) createTexture2D(ctx *opengl.Context, t *texture.Texture2D) {
	img := t.Image()

	// Create texture identity
	var ident uint32
	ctx.GenTextures(1, &ident)
	ctx.Execute()
	t.SetNativeIdentity(ident)

	// Select the texture
	ctx.BindTexture(opengl.TEXTURE_2D, ident)
	defer ctx.BindTexture(opengl.TEXTURE_2D, 0)

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
		ctx.TexParameterfv(opengl.TEXTURE_2D, opengl.TEXTURE_BORDER_COLOR, &borderColor[0])
		ctx.Execute()
	}

	ctx.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_WRAP_S, int32(glWrapU))
	ctx.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_WRAP_T, int32(glWrapV))

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

	ctx.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_MIN_FILTER, int32(glMinFilter))
	ctx.TexParameteri(opengl.TEXTURE_2D, opengl.TEXTURE_MAG_FILTER, int32(glMagFilter))

	// Enable mipmap generation if either filter is mipmapped.
	if minFilter.Mipmapped() || magFilter.Mipmapped() {
		ctx.TexParameteri(opengl.TEXTURE_2D, opengl.GENERATE_MIPMAP, int32(opengl.TRUE))
	}

	// Determine formatting, load texture into OpenGL
	internalFormat := opengl.RGBA
	if t.Compressed() {
		// We can only ask the driver to convert to a compressed format if it
		// actually supports it.
		for _, format := range r.compressedTextureFormats {
			switch format {
			case opengl.COMPRESSED_RGBA:
				internalFormat = format
				break
			}
		}
	}

	// Load texture
	sz := t.Bounds().Size()

	t.RLock()
	img = t.Source.(*image.RGBA)
	ctx.TexImage2D(opengl.TEXTURE_2D, 0, int32(internalFormat), uint32(sz.X), uint32(sz.Y), 0, opengl.RGBA, opengl.UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	ctx.Execute()
	t.RUnlock()

	runtime.SetFinalizer(t, func(t *texture.Texture2D) {
		r.texturesToFreeAccess.Lock()
		defer r.texturesToFreeAccess.Unlock()

		r.texturesToFree = append(r.texturesToFree, t.NativeIdentity().(uint32))
	})
}
