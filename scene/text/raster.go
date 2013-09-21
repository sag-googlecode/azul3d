// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package text

import (
	"errors"
	"log"
	"sync"
)

var (
	ErrGlyphNotFound = errors.New("Glyph not found in font source.")
)

// Source represents a generic font source.
type Source interface {
	// Rasterize must rasterize the single glyph, g, into a RGBA image.
	//
	// If the glyph is not known to the font source, the font source should
	// return ErrGlyphNotFound.
	Rasterize(r rune, o *GlyphOptions) (*GlyphRaster, error)

	// Kerning should return the amount of kerning between the pair of glyphs.
	//
	// The kerning values are _relative_ to where the glyphs would be placed
	// normally.
	//
	// If there is no known kerning amount for the pair, then [0, 0] should be
	// returned.
	Kerning(a, b rune) (x, y float64)
}

var (
	missingAccess  sync.RWMutex
	missingRune    rune
	missingOptions *GlyphOptions
)

func SetMissingGlyph(r rune, o *GlyphOptions) {
	missingAccess.Lock()
	defer missingAccess.Unlock()

	missingRune = r
	missingOptions = o
}

func MissingGlyph() (rune, *GlyphOptions) {
	missingAccess.RLock()
	defer missingAccess.RUnlock()

	return missingRune, missingOptions
}

func rasterize(r rune, o *GlyphOptions, fallback bool) *GlyphRaster {
	raster, err := o.Source.Rasterize(r, o)
	if err != nil {
		if err == ErrGlyphNotFound && fallback {
			missingRune, missingOptions := MissingGlyph()
			if missingOptions != nil {
				return rasterize(missingRune, missingOptions, false)
			} else {
				log.Printf("text.Rasterize(%q): %s\n", r, err.Error())
				return nil
			}
		} else {
			log.Printf("text.Rasterize(%q): %s\n", r, err.Error())
			return nil
		}
	}

	return raster
}

func Rasterize(r rune, o *GlyphOptions) *GlyphRaster {
	if o == nil {
		panic("Rasterize(): GlyphOptions parameter is nil!")
	}
	if o.Source == nil {
		panic("Rasterize(): GlyphOptions parameter has no source!")
	}
	return rasterize(r, o, true)
}

func Kerning(a, b rune, o *GlyphOptions) (x, y float64) {
	if o == nil {
		panic("Kerning(): GlyphOptions parameter is nil!")
	}
	if o.Source == nil {
		panic("Kerning(): GlyphOptions parameter has no source!")
	}
	return o.Source.Kerning(a, b)
}
