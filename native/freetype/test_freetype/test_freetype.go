// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build tests

// Test - Load a font, renders a glyph, displays some metrics.
package main

import (
	"azul3d.org/v1/native/freetype"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	ctx, err := freetype.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("FreeType context initialized")

	fontFile, err := os.Open("src/azul3d.org/v1/assets/fonts/vera/Vera.ttf")
	if err != nil {
		log.Fatal(err)
	}

	fontFileData, err := ioutil.ReadAll(fontFile)
	if err != nil {
		log.Fatal(err)
	}

	font, err := ctx.Load(fontFileData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded the font.")

	size := 500 * 64 // 50pt (expressed in 26.6 font units)
	dpi := 100       // 72 pixels per inch
	font.SetSize(size, 0, dpi, 0)
	//font.SetSizePixels(512, 512)

	glyphIndex := font.Index('A')
	log.Println("Glyph index for 'A':", glyphIndex)

	glyph, err := font.Load(glyphIndex)
	if err != nil {
		log.Fatal("Failed to load glyph.")
	}
	log.Println("Loaded the glyph.")

	m := glyph.HMetrics
	m = glyph.VMetrics
	xKern, yKern, err := font.Kerning('A', 'V')
	if err != nil {
		xKern = -1
		yKern = -1
	}

	log.Println("")
	log.Println("BBox:", font.BBox)
	log.Println("UnitsPerEm:", font.UnitsPerEm)
	log.Println("Ascender:", font.Ascender)
	log.Println("Descender:", font.Descender)
	log.Println("LineHeight:", font.LineHeight)
	log.Println("MaxAdvanceWidth:", font.MaxAdvanceWidth)
	log.Println("MaxAdvanceHeight:", font.MaxAdvanceHeight)
	log.Println("UnderlinePosition:", font.UnderlinePosition)
	log.Println("UnderlineThickness:", font.UnderlineThickness)
	log.Println("")
	log.Println("Glyph:")
	log.Println("")
	log.Println("    Width:", glyph.Width)
	log.Println("    Height:", glyph.Height)
	log.Println("")
	log.Println("    HMetrics:")
	log.Println("        BearingX:", m.BearingX)
	log.Println("        BearingY:", m.BearingY)
	log.Println("        Advance:", m.Advance)
	log.Println("        UnhintedAdvance:", m.UnhintedAdvance)
	log.Println("")
	log.Println("    VMetrics:")
	log.Println("        BearingX:", m.BearingX)
	log.Println("        BearingY:", m.BearingY)
	log.Println("        Advance:", m.Advance)
	log.Println("        UnhintedAdvance:", m.UnhintedAdvance)
	log.Println("")
	log.Printf("Kerning('A', 'V'): [%v, %v]\n", xKern, yKern)

	outFile, err := os.Create("test_freetype_out.png")
	if err != nil {
		log.Fatal(err)
	}

	glyphImage, err := glyph.Image()
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outFile, glyphImage)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Wrote test_freetype_out.png file.")
}
