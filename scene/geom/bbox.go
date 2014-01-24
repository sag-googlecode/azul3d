// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"azul3d.org/math"
)

// BoundingBox represents a single axis aligned bounding box.
type BoundingBox struct {
	Min, Max *math.Vec3
}

// Copy() returns a new 1:1 copy of this bounding box.
func (a *BoundingBox) Copy() *BoundingBox {
	cpy := *a
	return &cpy
}

// Equals tells if this bounding box is directly equal to b.
//
// Equality against nil is always false.
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

// Center returns a point representing the center of this bounding box.
func (a *BoundingBox) Center() *math.Vec3 {
	return a.Min.Lerp(a.Max, 0.5)
}
