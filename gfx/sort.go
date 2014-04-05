// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"azul3d.org/v1/math"
	"sort"
)

// ByDist sorts a list of graphics objects based on their distance away from
// a target position (typically the camera). As such if the sorted objects are
// drawn in order then they are drawn back-to-front (which is useful for
// rendering alpha-blended objects such that transparency appears correct).
//
// Using sort.Reverse this doubles as front-to-back sorting (which is useful
// for drawing opaque objects efficiently due to depth testing).
type ByDist struct {
	// The list of objects to sort.
	Objects []*Object

	// The target position to compare against. The list is sorted based off
	// each object's distance away from this position (typically this is the
	// camera's position).
	Target math.Vec3
}

// Implements sort.Interface.
func (b ByDist) Len() int {
	return len(b.Objects)
}

// Implements sort.Interface.
func (b ByDist) Swap(i, j int) {
	b.Objects[i], b.Objects[j] = b.Objects[j], b.Objects[i]
}

// Implements sort.Interface.
func (b ByDist) Less(i, j int) bool {
	// Calculate the distance from each object to the target position.
	iTransform := b.Objects[i].Transform
	jTransform := b.Objects[j].Transform
	iDist := iTransform.Pos.Sub(b.Target).Length()
	jDist := jTransform.Pos.Sub(b.Target).Length()

	// If i is further away from j (greater value) then it should sort first.
	return iDist > jDist
}

// InsertionSort performs a simple insertion sort on the sort interface. In the
// case of ByDist it performs generally as fast as sort.Sort() except that it
// can exploit temporal coherence improving performance dramatically when the
// objects have not moved much.
func InsertionSort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// ByState sorts a list of graphics objects based on the change of their
// graphics state in order to reduce graphics state changes and increase the
// overall throughput when rendering several objects whose graphics state
// differ.
type ByState []*Object

// Implements sort.Interface.
func (b ByState) Len() int {
	return len(b)
}

// Implements sort.Interface.
func (b ByState) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Implements sort.Interface.
func (b ByState) Less(i, j int) bool {
	k := b[i]
	v := b[j]
	return k.Compare(v)

	// Compare shaders.
	if k.Shader != v.Shader {
		return false
	}

	// Compare textures.
	for tIndex, t := range k.Textures {
		if v.Textures[tIndex] != t {
			return false
		}
	}

	// Compare state in order of most-commonly-changed.
	if k.AlphaMode != v.AlphaMode {
		return k.AlphaMode == DefaultState.AlphaMode
	}
	if k.Blend != v.Blend {
		return k.Blend.Compare(v.Blend)
	}
	if k.DepthTest != v.DepthTest {
		return k.DepthTest == DefaultState.DepthTest
	}
	if k.StencilTest != v.StencilTest {
		return k.StencilTest == DefaultState.StencilTest
	}
	if k.StencilFront != v.StencilFront {
		return k.StencilFront.Compare(v.StencilFront)
	}
	if k.StencilBack != v.StencilBack {
		return k.StencilBack.Compare(v.StencilBack)
	}
	if k.DepthWrite != v.DepthWrite {
		return k.DepthWrite == DefaultState.DepthWrite
	}
	if k.DepthCmp != v.DepthCmp {
		return k.DepthCmp == DefaultState.DepthCmp
	}
	if k.FaceCulling != v.FaceCulling {
		return k.FaceCulling == DefaultState.FaceCulling
	}
	if k.WriteRed != v.WriteRed {
		return k.WriteRed == DefaultState.WriteRed
	}
	if k.WriteGreen != v.WriteGreen {
		return k.WriteGreen == DefaultState.WriteGreen
	}
	if k.WriteBlue != v.WriteBlue {
		return k.WriteBlue == DefaultState.WriteBlue
	}
	if k.WriteAlpha != v.WriteAlpha {
		return k.WriteAlpha == DefaultState.WriteAlpha
	}
	if k.Dithering != v.Dithering {
		return k.Dithering == DefaultState.Dithering
	}
	return true
}
