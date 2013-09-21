// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package color implements various color types.
package color

import (
	"fmt"
	"math"
	"strconv"
)

func float32Equals(a, b, tolerence float32) bool {
	if a == b {
		return true
	}
	aa := float64(a)
	bb := float64(b)
	if float32(math.Abs(aa-bb)) < tolerence {
		return true
	}
	return false
}

// Color represents an single 32-bit floating point RGBA color.
//
// each color component is in the range of zero to one, where zero acts as
// 'no color' or transparent, and 1 acts as 'full color' or opaque.
//
// This is an appropriate color storage model for vertices as most graphics
// drivers internally convert to 32 bit floating point storage anyway.
//
// However, this is not an appropriate storage model for say, textures images,
// as the memory consumption would be too large (given that each color
// component is 32 bits).
type Color struct {
	R, G, B, A float32
}

// String returns an string representation of this color.
func (c Color) String() string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("Color([%.2f, %.2f, %.2f, %.2f], [%d, %d, %d, %d], #%s)", c.R, c.G, c.B, c.A, r, g, b, a, c.Hex())
}

// EqualsTolerence tells if this color is equal to the other color, by
// determining if it is within the specified tolerence for equality.
func (a Color) EqualsTolerence(b Color, tol float32) bool {
	rEqual := float32Equals(a.R, b.R, tol)
	gEqual := float32Equals(a.G, b.G, tol)
	bEqual := float32Equals(a.B, b.B, tol)
	aEqual := float32Equals(a.A, b.A, tol)
	return rEqual && gEqual && bEqual && aEqual
}

// Equals tells if this color is equal to the other color, by determining if it
// is within math.SmallestNonzeroFloat32 tolerence for equality.
func (a Color) Equals(b Color) bool {
	return a.EqualsTolerence(b, math.SmallestNonzeroFloat32)
}

// Blend returns the result of the two colors equally blended together, E.g. :
//
//  (a + b) / 2
//
func (a Color) Blend(b Color) Color {
	return Color{
		(a.R + b.R) / 2.0,
		(a.G + b.G) / 2.0,
		(a.B + b.B) / 2.0,
		(a.A + b.A) / 2.0,
	}
}

// Add returns the result of a + b
func (a Color) Add(b Color) Color {
	return Color{
		a.R + b.R,
		a.G + b.G,
		a.B + b.B,
		a.A + b.A,
	}
}

// AddScalar returns the result of a + b
func (a Color) AddScalar(b float32) Color {
	return Color{
		a.R + b,
		a.G + b,
		a.B + b,
		a.A + b,
	}
}

// Sub returns the result of a - b
func (a Color) Sub(b Color) Color {
	return Color{
		a.R - b.R,
		a.G - b.G,
		a.B - b.B,
		a.A - b.A,
	}
}

// SubScalar returns the result of a - b
func (a Color) SubScalar(b float32) Color {
	return Color{
		a.R - b,
		a.G - b,
		a.B - b,
		a.A - b,
	}
}

// Mul returns the result of a + b
func (a Color) Mul(b Color) Color {
	return Color{
		a.R * b.R,
		a.G * b.G,
		a.B * b.B,
		a.A * b.A,
	}
}

// MulScalar returns the result of a * b
func (a Color) MulScalar(b float32) Color {
	return Color{
		a.R * b,
		a.G * b,
		a.B * b,
		a.A * b,
	}
}

// Div returns the result of a + b
func (a Color) Div(b Color) Color {
	return Color{
		a.R / b.R,
		a.G / b.G,
		a.B / b.B,
		a.A / b.A,
	}
}

// DivScalar returns the result of a / b
func (a Color) DivScalar(b float32) Color {
	return Color{
		a.R / b,
		a.G / b,
		a.B / b,
		a.A / b,
	}
}

// RGBA returns the values of this color as 8-bit unsigned integers in the range of 0-255.
func (c Color) RGBA() (r, g, b, a uint8) {
	r = uint8(c.R * 255)
	g = uint8(c.G * 255)
	b = uint8(c.B * 255)
	a = uint8(c.A * 255)
	return
}

// Hex returns an color code in the format of:
//
// RRGGBBAA
//
// You may strip the alpha component if you wish, by slicing the string.
//
// The color code will not have an prefixed hash sign (#).
func (c Color) Hex() string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("%02X%02X%02X%02X", r, g, b, a)
}

// New returns a new color given red, green, blue, and alpha values in the
// range of 0.0 to 1.0.
//
// Values outside the range of 0.0 to 1.0 are clamped.
func New(r, g, b, a float32) Color {
	if r < 0 {
		r = 0
	} else if r > 1.0 {
		r = 1.0
	}

	if g < 0 {
		g = 0
	} else if g > 1.0 {
		g = 1.0
	}

	if b < 0 {
		b = 0
	} else if b > 1.0 {
		b = 1.0
	}

	if a < 0 {
		a = 0
	} else if a > 1.0 {
		a = 1.0
	}
	return Color{r, g, b, a}
}

// RGBA returns a new color given red, green, blue, and alpha values in the
// range of 0 to 255 (unsigned 8 bit values).
func RGBA(r, g, b, a uint8) Color {
	return Color{
		float32(r) / 255.0,
		float32(g) / 255.0,
		float32(b) / 255.0,
		float32(a) / 255.0,
	}
}

// Hex returns a new color given a hexadecimal color code in one of the
// following formats:
//
//  #RGB
//  #RGBA
//  #RRGGBB
//  #RRGGBBAA
//
// The code can optionally start with a hash character (#), as it will
// automatically be stripped out.
//
// The code (excluding hash) must be of length three (RGB), four (RGBA), six
// (RRGGBB), or eight (RRGGBBAA).
//
// If the color code is invalid, the color [0, 0, 0, 1] will be returned.
func Hex(hex string) Color {
	if len(hex) > 0 {
		// Strip prefixed hash
		if hex[0] == '#' {
			hex = hex[1:]
		}

		l := len(hex)
		if l == 3 {
			// RGB, rebuild such that it becomes RRGGBB
			r := hex[:1]
			g := hex[1:2]
			b := hex[2:]
			hex = r + r + g + g + b + b

		} else if l == 4 {
			// RGBA, rebuild such that it becomes RRGGBBAA
			r := hex[:1]
			g := hex[1:2]
			b := hex[2:3]
			a := hex[3:]
			hex = r + r + g + g + b + b + a + a
		}

		if len(hex) == 6 {
			// RRGGBB
			rgb, err := strconv.ParseUint(hex, 16, 32)
			if err != nil {
				return Default
			}
			return RGBA(
				uint8(rgb>>16),
				uint8((rgb<<48)>>56),
				uint8((rgb<<48)>>48),
				255,
			)

		} else if len(hex) == 8 {
			// RRGGBBAA
			rgba, err := strconv.ParseUint(hex, 16, 32)
			if err != nil {
				return Default
			}
			return RGBA(
				uint8(rgba>>24),
				uint8((rgba<<40)>>56),
				uint8((rgba<<48)>>56),
				uint8((rgba<<56)>>56),
			)
		}
	}
	return Default
}
