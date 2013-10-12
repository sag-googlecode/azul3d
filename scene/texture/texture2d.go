// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

import (
	"image"
)

// Texture2D represents a two-dimensional texture.
type Texture2D struct {
	*Texture
	bounds image.Rectangle
}

// Region returns a texture region from the given pixel coordinates.
func (tx *Texture2D) Region(u, v, s, t int) Region {
	if tx.Image == nil {
		panic("Region(): Texture image is nil; cannot determine image size!")
	}
	sz := tx.Bounds().Size()
	return RegionFromPixels(sz.X, sz.Y, u, v, s, t)
}

// Image returns the image source set on this texture.
func (t *Texture2D) Image() *image.RGBA {
	t.RLock()
	defer t.RUnlock()

	if t.Source == nil {
		return nil
	}
	return t.Source.(*image.RGBA)
}

// SetImage sets the image source of this texture.
func (t *Texture2D) SetImage(i *image.RGBA) {
	if i == nil {
		panic("SetImage(): Image is nil!")
	}

	t.Lock()
	defer t.Unlock()

	t.Source = i
	t.bounds = i.Bounds()
}

// Bounds returns the bounding rectangle of this two-dimensional texture.
func (t *Texture2D) Bounds() image.Rectangle {
	t.RLock()
	defer t.RUnlock()

	return t.bounds
}

// New returns a new two-dimensional texture.
func New() *Texture2D {
	t := new(Texture2D)
	t.Texture = NewBase()
	return t
}
