// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package text

import (
	"azul3d.org/v1/scene/texture"
	"image"
)

func Atlas(glyphs map[rune]*GlyphOptions, padding int, pot bool) (*image.RGBA, map[rune]map[*GlyphOptions]*GlyphRaster) {
	var images []image.Image

	rasterMap := make(map[rune]map[*GlyphOptions]*GlyphRaster)

	// Rasterize each glyph, create slice of rasterized images.
	for r, options := range glyphs {
		optionRasterMap, ok := rasterMap[r]
		if !ok {
			optionRasterMap = make(map[*GlyphOptions]*GlyphRaster, 1)
			rasterMap[r] = optionRasterMap
		}

		raster, ok := optionRasterMap[options]
		if !ok {
			raster = Rasterize(r, options)
			optionRasterMap[options] = raster
		}

		if raster == nil {
			// Unable to rasterize
			delete(glyphs, r)
		} else {
			// We rasterized it
			images = append(images, raster.Image)
		}
	}

	// Pack raster images.
	packed, rects := texture.Atlas(images, padding, pot)

	// Assign raster areas.
	for r, options := range glyphs {
		optionRasterMap, ok := rasterMap[r]
		if ok {
			raster, ok := optionRasterMap[options]
			if ok {
				rect := rects[raster.Image]
				raster.Area = &rect
			}
		}
	}

	return packed, rasterMap
}
