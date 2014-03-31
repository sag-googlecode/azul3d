// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

package main

import (
	"azul3d.org/v0/scene/text"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// Open font file
	file, err := os.Open("src/azul3d.org/v0/assets/fonts/vera/Vera.ttf")
	if err != nil {
		log.Fatal(err)
	}

	// Read all data
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Load the font
	ft, err := text.LoadFont(fileData)
	if err != nil {
		log.Fatal(err)
	}

	s := "QWERTYUIOPASDFGHJKLZXCVBNM{}|:\"<>?~!@#$%^&*()_+qwertyuiopasdfghjklzxcvbnm[]\\;',./`1234567890-="
	s += s
	s += s
	s += s
	s += s
	s += s
	s += s
	s += s
	s += s
	s += s
	s += s
	s += s

	// Must specify source for default options (it has none).
	o := text.DefaultOptions
	o.Source = ft
	o.Size = 32
	log.Printf("Rendering 64 glyph atlases with %v %v-point characters @ %v DPI\n", len(s), o.Size, o.DPI)

	start := time.Now()

	glyphs := make(map[rune]*text.GlyphOptions, len(s))
	for _, r := range s {
		glyphs[r] = o
	}

	var atlasImage image.Image
	var avg time.Duration
	for i := 0; i < 64; i++ {
		t0 := time.Now()

		// 1 px padding
		atlasImage, _ = text.Atlas(glyphs, 1, true)

		t1 := time.Since(t0)
		avg += t1
		log.Println("Atlas", i, t1)
	}
	log.Println("Average", avg/64)

	end := time.Since(start)
	log.Println("Time taken:", end)

	// Write glyph to file
	out, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(out, atlasImage)
	if err != nil {
		log.Fatal(err)
	}
}
