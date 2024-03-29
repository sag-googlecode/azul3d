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
//
// The Less() method properly read-locks the objects when required.
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
	k := b.Objects[i]
	v := b.Objects[j]

	// Lock both objects for reading.
	k.RLock()
	v.RLock()

	// Grab their transforms.
	iTransform := b.Objects[i].Transform
	jTransform := b.Objects[j].Transform

	// Unlock the objects.
	v.RUnlock()
	k.RUnlock()

	// Calculate the distance from each object to the target position.
	iDist := iTransform.Pos().Sub(b.Target).Length()
	jDist := jTransform.Pos().Sub(b.Target).Length()

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
//
// The Less() method properly read-locks the objects when required.
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

	// Lock both objects for reading.
	k.RLock()
	v.RLock()

	// Grab state comparison.
	less := k.Compare(v)

	// Unlock the objects.
	v.RUnlock()
	k.RUnlock()
	return less
}
