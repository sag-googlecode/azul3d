// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"image"
	"sync"
)

// Downloadable represents a image that can be downloaded from the graphics
// hardware into system memory (e.g. for taking a screen-shot).
type Downloadable interface {
	// Download should download the given intersecting rectangle of this
	// downloadable image from the graphics hardware into system memory and
	// send it to the complete channel when done.
	//
	// If downloading this texture is impossible (i.e. hardware does not
	// support this) then nil will be sent over the channel and all future
	// attempts to download this texture will fail as well.
	//
	// It should be noted that the downloaded image may not be pixel-identical
	// to the previously uploaded source image of a texture, for instance if
	// texture compression was used it may suffer from compression artifacts,
	// etc.
	Download(r image.Rectangle, complete chan image.Image)
}

// NativeTexture represents the native object of a *Texture, the renderer is
// responsible for creating these and fulfilling the interface.
type NativeTexture interface {
	Downloadable
}

// Texture represents a single 2D texture that may be applied to a mesh for
// drawing.
//
// Clients are responsible for utilizing the RWMutex of the texture when using
// it or invoking methods.
type Texture struct {
	sync.RWMutex

	// The native object of this texture. The renderer using this texture will
	// assign this at load time.
	NativeTexture

	// Weather or not this texture is currently loaded or not.
	Loaded bool

	// If true then when this texture is loaded the data image source of it
	// will be kept instead of being set to nil (which allows it to be garbage
	// collected).
	KeepDataOnLoad bool

	// The bounds of the texture, in the case of a texture loaded from a image
	// this should be set to the image's bounds. In the case of rendering to a
	// texture this should be set to the desired canvas resolution.
	Bounds image.Rectangle

	// The source image of the texture, may be nil (i.e. in the case of render
	// to texture, unless downloaded).
	Source image.Image

	// If set to true then the renderer using this texture should try it's best
	// to avoid texture compression related artifacts at the cost of increased
	// memory consumption.
	Uncompressed bool

	// The U and V wrap modes of this texture.
	WrapU, WrapV TexWrap

	// The color of the border when a wrap mode is set to WM_BORDER_COLOR.
	BorderColor Color

	// The texture filtering used for minification and magnification of the
	// texture.
	MinFilter, MagFilter TexFilter
}

// Copy returns a new copy of this Texture. Explicitly not copied over is the
// native texture, the OnLoad slice, the Loaded status, and the source image
// (because the image type is not strictly known). Because the texture's source
// image is not copied over, you may want to copy it directly over yourself.
//
// The texture's read lock must be held for this method to operate safely.
func (t *Texture) Copy() *Texture {
	return &Texture{
		sync.RWMutex{},
		nil,   // Native texture -- not copied.
		false, // Loaded status -- not copied.
		t.KeepDataOnLoad,
		t.Bounds,
		nil, // Source image -- not copied.
		t.Uncompressed,
		t.WrapU,
		t.WrapV,
		t.BorderColor,
		t.MinFilter,
		t.MagFilter,
	}
}

// ClearData sets the data source image, t.Source, of this texture to nil if
// t.KeepDataOnLoad is set to false.
//
// The texture's write lock must be held for this method to operate safely.
func (t *Texture) ClearData() {
	if !t.KeepDataOnLoad {
		t.Source = nil
	}
}
