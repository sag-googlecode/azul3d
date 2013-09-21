// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

import (
	"code.google.com/p/azul3d/binpack"
	"image"
	"image/draw"
	"sort"
)

type sortedImages []image.Image

func (s sortedImages) Len() int {
	return len(s)
}

func (s sortedImages) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedImages) Less(i, j int) bool {
	iBoundsSize := s[i].Bounds().Size()
	jBoundsSize := s[j].Bounds().Size()

	iSize := iBoundsSize.Y
	if iBoundsSize.X > iSize {
		iSize = iBoundsSize.X
	}

	jSize := jBoundsSize.Y
	if jBoundsSize.X > jSize {
		jSize = jBoundsSize.X
	}

	// We sort the images such that the image with the larger area will sort
	// before images with smaller areas.
	//
	// So we simply use Less() in reverse order (I.e. i should sort before j,
	// given that iSize is larger than jSize).
	return iSize > jSize
}

type packableRects []image.Rectangle

func (p packableRects) Len() int {
	return len(p)
}

func (p packableRects) Size(n int) (width, height int) {
	sz := p[n].Size()
	return sz.X, sz.Y
}

func (p packableRects) Place(n, x, y int) {
	sz := p[n].Size()
	p[n] = image.Rect(x, y, x+sz.X, y+sz.Y)
}

// Atlas creates a texture atlas by combining several images into one single
// larger one.
//
// The pot argument specifies if the returned image should be a power of two
// size (to obtain the power of two size, the resulting image will have extra
// 'wasted' pixels added).
//
// The returned map simply maps each input image to it's new area on the
// resulting image.
//
// If len(images) == 0; then [nil, nil] is returned.
func Atlas(images []image.Image, padding int, pot bool) (*image.RGBA, map[image.Image]image.Rectangle) {
	if len(images) == 0 {
		return nil, nil
	}

	// Sort the images by largest first
	si := sortedImages(images)
	sort.Sort(si)

	// Now build rectangle slice
	p := make(packableRects, len(si))
	for i, img := range si {
		bds := img.Bounds()
		bds.Min.X -= padding
		bds.Min.Y -= padding
		bds.Max.X += padding
		bds.Max.Y += padding
		p[i] = bds
	}

	// binpack the rectangles
	width, height := binpack.Pack(p)
	if width == -1 || height == -1 {
		panic("Atlas(): binpack failed due to incorrect sorting; this should never happen!")
	}

	// Create a image of [width, height]
	result := image.NewRGBA(image.Rect(0, 0, width, height))

	// Create rectangle mapping
	rects := make(map[image.Image]image.Rectangle)

	// Draw each sub-rectangle onto the new image, and map each rectangle now.
	for i, rect := range p {
		srcImage := si[i]
		rect.Min.X += padding
		rect.Min.Y += padding
		rect.Max.X -= padding
		rect.Max.Y -= padding
		rects[srcImage] = rect
		draw.Draw(result, rect, srcImage, image.ZP, draw.Over)
	}

	return result, rects
}
