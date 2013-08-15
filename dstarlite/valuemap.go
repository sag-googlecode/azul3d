// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package dstarlite

import (
	"math"
)

// Just a small helper type.
type valueMap map[State]float64

// Get returns the specified key in the map, or if the specified key does not
// exist returns +Inf
func (v valueMap) get(s State) float64 {
	val, ok := v[s]
	if !ok {
		return math.Inf(1)
	}
	return val
}
