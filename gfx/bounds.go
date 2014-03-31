// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"azul3d.org/v1/math"
)

// Boundable is any type which can return it's axis aligned bounding box.
type Boundable interface {
	AABB() AABB
}

// AABB represents a single axis aligned bounding box in 3D space.
type AABB struct {
	// The minimum (e.g. most negative) and maximum (e.g. most positive)
	// points.
	Min, Max math.Vec3
}

// Equals tells if a == b using the default math.EPSILON value for comparison.
func (a AABB) Equals(b AABB) bool {
	return a.Min.Equals(b.Min) && a.Max.Equals(b.Max)
}

// Empty reports whether the AABB contains no points.
func (a AABB) Empty() bool {
	return a.Min.Equals(math.Vec3Zero) && a.Max.Equals(math.Vec3Zero)
}

// Contains tells if the AABB a contains the AABB b (i.e. if b is inside a).
func (a AABB) Contains(b AABB) bool {
	if a.Min.Equals(b.Min) && a.Max.Equals(b.Max) {
		return true
	}

	if b.Min.X < a.Min.X {
		return false
	}
	if b.Min.Y < a.Min.Y {
		return false
	}
	if b.Min.Z < a.Min.Z {
		return false
	}

	if b.Max.X > a.Max.X {
		return false
	}
	if b.Max.Y > a.Max.Y {
		return false
	}
	if b.Max.Z > a.Max.Z {
		return false
	}
	return true
}

// Fit returns a AABB that contains both a and b.
func (a AABB) Fit(b AABB) AABB {
	if b.Min.X < a.Min.X {
		a.Min.X = b.Min.X
	}
	if b.Min.Y < a.Min.Y {
		a.Min.Y = b.Min.Y
	}
	if b.Min.Z < a.Min.Z {
		a.Min.Z = b.Min.Z
	}

	if b.Max.X > a.Max.X {
		a.Max.X = b.Max.X
	}
	if b.Max.Y > a.Max.Y {
		a.Max.Y = b.Max.Y
	}
	if b.Max.Z > a.Max.Z {
		a.Max.Z = b.Max.Z
	}
	return a
}

// Size returns the size of this AABB along the X, Y, and Z axis.
func (a AABB) Size() math.Vec3 {
	return a.Max.Sub(a.Min)
}

// Center returns the center point of this AABB.
func (a AABB) Center() math.Vec3 {
	return a.Min.Add(a.Size().DivScalar(2.0))
}
