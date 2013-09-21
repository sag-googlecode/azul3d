// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

type Primitive uint8

func (p Primitive) Valid() bool {
	switch p {
	case Triangles:
		return true
	case TriangleStrips:
		return true
	case Points:
		return true
	case Lines:
		return true
	case LineStrips:
		return true
	}
	return false
}

const (
	Triangles Primitive = iota
	TriangleStrips
	Points
	Lines
	LineStrips
)
