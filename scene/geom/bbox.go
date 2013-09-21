// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"code.google.com/p/azul3d/math"
)

type BoundingBox struct {
	Min, Max *math.Vec3
}

func (a *BoundingBox) Copy() *BoundingBox {
	cpy := *a
	return &cpy
}

func (a *BoundingBox) Equals(b *BoundingBox) bool {
	if b == nil {
		return false
	}

	if a.Min == nil && b.Min != nil {
		return false
	}
	if a.Min != nil && b.Min == nil {
		return false
	}

	if a.Max == nil && b.Max != nil {
		return false
	}
	if a.Max != nil && b.Max == nil {
		return false
	}

	if a == nil && b == nil {
		return true
	}
	return a.Min.Equals(b.Min) && a.Max.Equals(b.Max)
}
