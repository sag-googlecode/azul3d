// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package color

import (
	"testing"
)

func TestHexHashRGB(t *testing.T) {
	hex := Hex("39C")
	if !hex.Equals(Hex("39C")) {
		t.Log("Stripping hash failed.")
		t.Fail()
	}
	if hex.R != 0.2 || hex.G != 0.6 || hex.B != 0.8 || hex.A != 1 {
		t.Log(hex)
		t.Fail()
	}
}

func TestHexHashRGBA(t *testing.T) {
	hex := Hex("#39C3")
	if !hex.Equals(Hex("39C3")) {
		t.Log("Stripping hash failed.")
		t.Fail()
	}
	if hex.R != 0.2 || hex.G != 0.6 || hex.B != 0.8 || hex.A != 0.2 {
		t.Log(hex)
		t.Fail()
	}
}

func TestHexHashRRGGBB(t *testing.T) {
	hex := Hex("#3399CC")
	if !hex.Equals(Hex("3399CC")) {
		t.Log("Stripping hash failed.")
		t.Fail()
	}
	if hex.R != 0.2 || hex.G != 0.6 || hex.B != 0.8 || hex.A != 1 {
		t.Log(hex)
		t.Fail()
	}
}

func TestHexHashRRGGBBAA(t *testing.T) {
	hex := Hex("#3399CC33")
	if !hex.Equals(Hex("3399CC33")) {
		t.Log("Stripping hash failed.")
		t.Fail()
	}
	if hex.R != 0.2 || hex.G != 0.6 || hex.B != 0.8 || hex.A != 0.2 {
		t.Log(hex)
		t.Fail()
	}
}

func TestRGBA(t *testing.T) {
	rgba := RGBA(51, 153, 204, 51)
	if rgba.R != 0.2 || rgba.G != 0.6 || rgba.B != 0.8 || rgba.A != 0.2 {
		t.Log(rgba)
		t.Fail()
	}
}
