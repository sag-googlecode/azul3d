// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package ui

import (
	"azul3d.org/math"
	"azul3d.org/scene"
)

type OptionType int

const (
	Parent OptionType = iota
	Width
	Height
	MarginLeft
	MarginRight
	MarginBottom
	MarginTop
	Text
	Layout
	Origin
	Color
	ColorScale
)

func (o OptionType) Valid() bool {
	switch {
	case o == Parent:
		return true
	case o == Width || o == Height:
		return true
	case o == MarginLeft || o == MarginRight || o == MarginBottom || o == MarginTop:
		return true
	case o == Text:
		return true
	case o == Layout || o == Origin:
		return true
	case o == Color:
		return true
	case o == ColorScale:
		return true
	case o == Layout:
		return true
	}
	return false
}

func (o OptionType) ValidValue(v interface{}) (ok bool) {
	switch {
	case o == Parent:
		_, ok = v.(*scene.Node)
	case o == Width || o == Height || o == MarginLeft || o == MarginRight || o == MarginBottom || o == MarginTop:
		var vv int
		vv, ok = v.(int)
		if ok {
			ok = vv >= 0
		}
	case o == Text:
		_, ok = v.(string)
	case o == Layout || o == Origin:
		var vv LayoutType
		vv, ok = v.(LayoutType)
		if ok {
			ok = vv.Valid()
		}
	case o == Color || o == ColorScale:
		_, ok = v.(*math.Vec4)
	}
	return
}
