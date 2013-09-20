package freetype

/*
#cgo windows amd64 LDFLAGS: freetype_windows_amd64.a
#cgo CFLAGS: -I freetype-2.5.0.1/include/

#include <ft2build.h>
#include FT_FREETYPE_H
#include FT_SIZES_H
#include FT_GLYPH_H
*/
import "C"

import (
	"fmt"
	"image"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

type GlyphMetrics struct {
	// Left side bearing and top side bearing
	// X values extend to the right, and positive Y values downward.
	// Expressed in font units.
	BearingX, BearingY int

	// Advance (and unhinted advance) amount of glyph.
	// For horizontal metrics, a positive value means advancing to the right.
	// For vertical metrics, a positive value means advancing downward.
	// Expressed in font units.
	Advance, UnhintedAdvance int
}

type Glyph struct {
	img         *image.Alpha
	renderImage func() (*image.Alpha, error)

	// Width and height of glyph.
	// Expressed in font units.
	Width, Height int

	// Horizontal and vertical glyph metrics.
	HMetrics, VMetrics GlyphMetrics
}

// Renders and returns a 8-bit grayscale image of this glyph.
func (g *Glyph) Image() (*image.Alpha, error) {
	if g.img == nil {
		var err error
		g.img, err = g.renderImage()
		if err != nil {
			return nil, err
		}
	}
	return g.img, nil
}

type Font struct {
	ctx *Context

	data []uint8
	c    C.FT_Face

	// Bounding box that is large enough to contain any glyph in the font face.
	// Expressed in font units.
	BBox image.Rectangle

	// The number of font units per EM square for this font face.
	// Expressed in font units.
	UnitsPerEm int

	// The typographic ascender of the face
	// Expressed in font units.
	Ascender int

	// The typographic descender of the face.
	// Expressed in font units.
	Descender int

	// The vertical distance between two consecutive baselines.
	// Expressed in font units.
	LineHeight int

	// The maximum advance width for all glyphs in this face.
	// This can be used to make word wrapping computations faster.
	// Expressed in font units.
	MaxAdvanceWidth int

	// The maximum advance height, for all glyphs in this face.
	// This is only relevant for vertical layouts, and is set to ‘height’
	// for fonts that do not provide vertical metrics.
	// Expressed in font units.
	MaxAdvanceHeight int

	// The position of the underline for this font face.
	// Expressed in font units.
	UnderlinePosition int

	// The thickness for the underline of this font face.
	// Expressed in font units.
	UnderlineThickness int
}

func (f *Font) init() {
	f.SetSize(24*64, 24*64, 72, 72)

	f.ctx.access.Lock()
	defer f.ctx.access.Unlock()

	err := C.FT_Select_Charmap(f.c, C.FT_ENCODING_UNICODE)
	if err != 0 {
		fmt.Println("Font.init(): FT_Select_Charmap() failed!")
	}

	b := f.c.bbox
	f.BBox = image.Rect(
		int(b.xMin),
		int(b.yMin),
		int(b.xMax),
		int(b.yMax),
	)

	f.UnitsPerEm = int(f.c.units_per_EM)
	f.Ascender = int(f.c.ascender)
	f.Descender = int(f.c.descender)
	f.LineHeight = int(f.c.height)
	f.MaxAdvanceWidth = int(f.c.max_advance_width)
	f.MaxAdvanceHeight = int(f.c.max_advance_height)
	f.UnderlinePosition = int(f.c.underline_position)
	f.UnderlineThickness = int(f.c.underline_thickness)
}

func (f *Font) SetSize(width, height, xResolution, yResolution int) error {
	f.ctx.access.Lock()
	defer f.ctx.access.Unlock()

	if width < 0 || height < 0 {
		panic("SetSize(): width < 0 || height < 0")
	}

	if xResolution < 0 || xResolution < 0 {
		panic("SetSize(): width < 0 || height < 0")
	}

	err := C.FT_Set_Char_Size(
		f.c,
		C.FT_F26Dot6(width),
		C.FT_F26Dot6(height),
		C.FT_UInt(xResolution),
		C.FT_UInt(yResolution),
	)
	if err != 0 {
		return lookupErr[int(err)]
	}
	return nil
}

func (f *Font) SetSizePixels(width, height int) error {
	f.ctx.access.Lock()
	defer f.ctx.access.Unlock()

	if width < 0 || height < 0 {
		panic("SetSizePixels(): width < 0 || height < 0")
	}

	err := C.FT_Set_Pixel_Sizes(
		f.c,
		C.FT_UInt(width),
		C.FT_UInt(height),
	)
	if err != 0 {
		return lookupErr[int(err)]
	}
	return nil
}

func (f *Font) Index(r rune) (glyphIndex uint) {
	f.ctx.access.Lock()
	defer f.ctx.access.Unlock()

	return uint(C.FT_Get_Char_Index(f.c, C.FT_ULong(r)))
}

func (f *Font) Kerning(leftGlyph, rightGlyph rune) (x, y int, e error) {
	f.ctx.access.Lock()
	defer f.ctx.access.Unlock()

	left := C.FT_Get_Char_Index(f.c, C.FT_ULong(leftGlyph))
	right := C.FT_Get_Char_Index(f.c, C.FT_ULong(rightGlyph))
	if left == 0 || right == 0 {
		return 0, 0, nil
	}

	var vec C.FT_Vector
	err := C.FT_Get_Kerning(
		f.c,
		left,
		right,
		C.FT_KERNING_DEFAULT,
		&vec,
	)
	if err != 0 {
		return 0, 0, lookupErr[int(err)]
	}
	return int(vec.x), int(vec.y), nil
}

func (f *Font) Load(glyphIndex uint) (*Glyph, error) {
	f.ctx.access.Lock()
	defer f.ctx.access.Unlock()

	err := C.FT_Load_Glyph(
		f.c,
		C.FT_UInt(glyphIndex),
		C.FT_LOAD_DEFAULT|C.FT_LOAD_LINEAR_DESIGN,
	)
	if err != 0 {
		return nil, lookupErr[int(err)]
	}

	g := f.c.glyph

	renderImage := func() (*image.Alpha, error) {
		err = C.FT_Render_Glyph(g, C.FT_RENDER_MODE_NORMAL)
		if err != 0 {
			return nil, lookupErr[int(err)]
		}

		// The face's glyph slot will change, so we need to copy the bitmap.
		width := int(g.bitmap.width)
		height := int(g.bitmap.rows)
		length := width * height

		var data []uint8
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
		sliceHeader.Cap = length
		sliceHeader.Len = length
		sliceHeader.Data = uintptr(unsafe.Pointer(g.bitmap.buffer))

		//cpy := make([]uint8, len(data))
		//copy(cpy, data)

		img := image.NewAlpha(image.Rect(0, 0, width, height))
		img.Pix = data
		img.Stride = width
		return img, nil
	}

	m := g.metrics
	return &Glyph{
		renderImage: renderImage,
		Width:       int(m.width),
		Height:      int(m.height),
		HMetrics: GlyphMetrics{
			BearingX:        int(m.horiBearingX),
			BearingY:        int(m.horiBearingY),
			Advance:         int(m.horiAdvance),
			UnhintedAdvance: int(g.linearHoriAdvance),
		},
		VMetrics: GlyphMetrics{
			BearingX:        int(m.vertBearingX),
			BearingY:        int(m.vertBearingY),
			Advance:         int(m.vertAdvance),
			UnhintedAdvance: int(g.linearVertAdvance),
		},
	}, nil
}

type Context struct {
	access sync.Mutex
	c      C.FT_Library
}

func (c *Context) Load(fontFileData []byte) (*Font, error) {
	c.access.Lock()

	f := new(Font)
	f.ctx = c
	f.data = fontFileData
	err := C.FT_New_Memory_Face(
		c.c,
		(*C.FT_Byte)(unsafe.Pointer(&f.data[0])),
		C.FT_Long(len(f.data)),
		0,
		&f.c,
	)
	if err != 0 {
		return nil, lookupErr[int(err)]
	}

	c.access.Unlock()

	f.init()

	runtime.SetFinalizer(f, func(f *Font) {
		c.access.Lock()
		defer c.access.Unlock()

		C.FT_Done_Face(f.c)
	})
	return f, nil
}

func Init() (*Context, error) {
	c := new(Context)
	err := C.FT_Init_FreeType(&c.c)
	if err != 0 {
		return nil, lookupErr[int(err)]
	}

	runtime.SetFinalizer(c, func(c *Context) {
		c.access.Lock()
		defer c.access.Unlock()

		C.FT_Done_FreeType(c.c)
	})
	return c, nil
}
