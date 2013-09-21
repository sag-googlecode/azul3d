// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package text

import (
	"code.google.com/p/azul3d/native/freetype"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
)

type Freetype struct {
	context *freetype.Context
	font    *freetype.Font
}

func (f *Freetype) Kerning(a, b rune) (x, y float64) {
	xi, yi, err := f.font.Kerning(a, b)
	if err != nil {
		log.Println("Freetype.Kerning():", err)
		return 0, 0
	}
	return float64(xi / 64), float64(yi / 64)
}

// Rasterize must rasterize the single glyph, g, into a RGBA image.
//
// If the glyph is not known to the font source, the font source should
// return ErrGlyphNotFound.
func (f *Freetype) Rasterize(r rune, o *GlyphOptions) (*GlyphRaster, error) {
	// Locate the glyph index in the front
	glyphIndex := f.font.Index(r)
	if glyphIndex == 0 {
		// We couldn't locate it.
		return nil, ErrGlyphNotFound
	}

	// Load size/DPI
	iSize := int(o.Size+0.5) * 64
	iDPI := int(o.DPI + 0.5)
	f.font.SetSize(iSize, iSize, iDPI, iDPI)

	// Load the glyph
	glyph, err := f.font.Load(glyphIndex)
	if err != nil {
		// We couldn't load it.
		return nil, err
	}

	// Rasterize the glyph
	mask, err := glyph.Image()
	if err != nil {
		// Couldn't rasterize it
		return nil, err
	}

	// Create RGBA image
	img := image.NewRGBA(mask.Bounds())

	// Draw background color
	if o.Background != nil {
		bgImage := image.NewUniform(o.Background)
		draw.Draw(img, img.Bounds(), bgImage, image.ZP, draw.Over)
	}

	// Draw foreground color
	fgColor := color.Color(color.Black)
	if o.Foreground != nil {
		fgColor = o.Foreground
	}

	fgImage := image.NewUniform(fgColor)
	draw.DrawMask(img, img.Bounds(), fgImage, image.ZP, mask, image.ZP, draw.Over)

	return &GlyphRaster{
		Image: img,
		HMetrics: GlyphMetrics{
			Advance:  float64(glyph.HMetrics.Advance) / 64.0,
			BearingX: float64(glyph.HMetrics.BearingX) / 64.0,
			BearingY: float64(glyph.HMetrics.BearingY) / 64.0,
		},
		VMetrics: GlyphMetrics{
			Advance:  float64(glyph.VMetrics.Advance) / 64.0,
			BearingX: float64(glyph.VMetrics.BearingX) / 64.0,
			BearingY: float64(glyph.VMetrics.BearingY) / 64.0,
		},
	}, nil
}

// LoadFont loads a freetype font file and returns a new, initialized freetype
// font source.
func LoadFont(data []byte) (*Freetype, error) {
	var (
		err error
		ft  = new(Freetype)
	)

	// Initialize freetype context
	ft.context, err = freetype.Init()
	if err != nil {
		return nil, err
	}

	// Load font
	ft.font, err = ft.context.Load(data)
	if err != nil {
		return nil, err
	}

	return ft, nil
}

// LoadFontFile is a small helper function to LoadFont. It simply opens the
// specified file path, and loads the freetype font.
//
// Any errors which occur are returned.
func LoadFontFile(filePath string) (*Freetype, error) {
	// Open font file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Read all data
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Load the font
	ft, err := LoadFont(fileData)
	if err != nil {
		return nil, err
	}

	return ft, nil
}

/*

// Freetype is a freetype font source.
type Freetype struct {
	access sync.Mutex

	font *truetype.Font
	context *freetype.Context
	buffer *truetype.GlyphBuf
}

// Implements font Source interface.
//
// Note: Must be goroutine safe.
func (t *Freetype) Rasterize(r rune, o *GlyphOptions) (*GlyphRaster, error) {
	// Since this function needs to be safe to invoke from multiple goroutines
	// we use a mutex here (because freetype context is not goroutine safe).
	t.access.Lock()
	defer t.access.Unlock()

	// Locate glyph in font
	glyphIndex := t.font.Index(r)
	if glyphIndex == 0 {
		// Glyph doesn't exist in font
		return nil, ErrGlyphNotFound
	}

	// Determine the scale
	scale := int32(o.Size * o.DPI * (64.0 / 72.0))

	hMetric := t.font.HMetric(int32(o.Size), glyphIndex)

	// Find the bounds of the glyph
	err := t.buffer.Load(t.font, scale, glyphIndex, &truetype.Hinter{})
	if err != nil {
		return nil, err
	}
	bounds := t.buffer.B

	// Calculate pixel width and height of the glyph's bounding box
	pixelWidth := int(math.Ceil(float64(bounds.XMax - bounds.XMin) / 64.0)) + 2
	pixelHeight := int(math.Ceil(float64(bounds.YMax - bounds.YMin) / 64.0)) + 2


	// Create an RGBA image to hold the rasterized glyph
	img := image.NewRGBA(image.Rect(0, 0, pixelWidth, pixelHeight))

	// Paint the background color
	if o.Background != color.Transparent {
		bgUniform := image.NewUniform(o.Background)
		draw.Draw(img, img.Bounds(), bgUniform, image.ZP, draw.Src)
	}

	// Update FreeType context
	t.context.SetDPI(o.DPI)
	t.context.SetFontSize(o.Size)

	fgUniform := image.NewUniform(o.Foreground)
	t.context.SetSrc(fgUniform)

	// Inform the FreeType context to draw to the image
	t.context.SetDst(img)
	t.context.SetClip(img.Bounds())

	// Offset by the minimum bounds, because the bounds of the image are not always relative
	// to an zero origin. (e.g. lowercase "j" character goes to left and below (0, 0) origin)
	xPos := img.Bounds().Min.X
	yPos := img.Bounds().Max.Y

	xMin := float64(bounds.XMin)
	yMin := float64(bounds.YMin)
	xMax := float64(bounds.XMax)
	yMax := float64(bounds.YMax)

	if xMin < 0 {
		xPos += int((-xMin / 64.0) + 0.5)
	} else if xMin > 0 {
		xPos -= int((xMin / 64.0) + 0.5)
	}

	if yMin < 0 {
		yPos -= int((-yMin / 64) + 0.5)
	} else if yMin > 0 {
		yPos += int((yMin / 64) + 0.5)
	}

	log.Printf("\n\n\nBEGIN %q\n", string(r))
	// Now paint the glyph to the image
	pt := freetype.Pt(xPos, yPos)
	_, err = t.context.DrawString(string(r), pt)
	if err != nil {
		return nil, err
	}

	// Find [X, Y] bearing
	xPos = img.Bounds().Min.X
	log.Println("xPos", xPos)
	log.Println("xMin", xMin)


	var bearingX, bearingY int
	if xMin < 0 {
		bearingX = int(hMetric.AdvanceWidth) + (int((-xMin) + 0.5) >> 2)
	} else {
		bearingX = int(hMetric.AdvanceWidth) + (int((xMin) + 0.5) >> 2)
	}
	bearingX = int(hMetric.AdvanceWidth)

	_ = yMax
	_ = xMax

	log.Println("bearingX", bearingX)
	log.Println("width", img.Bounds().Size().X)
	log.Println("actual", int(hMetric.AdvanceWidth))
	if bearingX != int(hMetric.AdvanceWidth) {
		log.Println("******************WARNING************************")
	}

	return &GlyphRaster{
		Image: img,
		HMetrics: GlyphMetrics {
			Advance: int(hMetric.AdvanceWidth),
			BearingX: bearingX,
			BearingY: bearingY,
		},
	}, nil
}

// LoadFont loads a freetype font and returns a new, initialized freetype font
// source.
func LoadFont(data []byte) (*Freetype, error) {
	var(
		err error
		ft = new(Freetype)
	)

	// Load font
	ft.font, err = freetype.ParseFont(data)
	if err != nil {
		return nil, err
	}

	// Create freetype context
	ft.context = freetype.NewContext()

	// Set font now (since we tie context to font).
	ft.context.SetFont(ft.font)

	ft.buffer = truetype.NewGlyphBuf()

	return ft, nil
}

// LoadFontFile is a small helper function to LoadFont. It simply opens the
// specified file path, and loads the freetype font.
//
// Any errors which occur are returned.
func LoadFontFile(filePath string) (*Freetype, error) {
	// Open font file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Read all data
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Load the font
	ft, err := LoadFont(fileData)
	if err != nil {
		return nil, err
	}

	return ft, nil
}
*/
