// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package text

import (
	"image"
	"image/color"
)

type GlyphOptions struct {
	// Source to use for rasterizing the glyph
	Source Source

	// Size of the glyph in points ("12 point font")
	Size float64

	// DPI to use for rasterizing the glyph.
	DPI float64

	// Foreground and background color of glyph.
	Foreground, Background color.Color
}

type GlyphRaster struct {
	// Rasterized glyph image.
	//
	// This will be nil if this glyph is not yet rasterized.
	Image *image.RGBA

	// Area describes the area of the rasterized glyph on a larger font atlas.
	//
	// This will be nil if this glyph is not yet on a larger font atlas.
	Area *image.Rectangle

	// HMetrics describes the horizontal metrics of this glyph.
	HMetrics GlyphMetrics

	// VMetrics describes the vertical metrics of this glyph.
	VMetrics GlyphMetrics
}

// Reference for these values can be found at:
//
//  http://www.freetype.org/freetype2/docs/glyphs/glyphs-3.html
//
type GlyphMetrics struct {
	Advance  float64
	BearingX float64
	BearingY float64
}

var DefaultOptions = &GlyphOptions{
	Source:     nil,
	Size:       50.0,
	DPI:        72,
	Foreground: color.Black,
	Background: color.Transparent,
}
