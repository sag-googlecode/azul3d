// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

// Coord represents a texture coordinate, with U, V, W, and Z components.
type Coord struct {
	U, V, W, Z float32
}

// UV returns a two-dimensional texture coordinate with U, and V components.
func UV(u, v float32) Coord {
	return Coord{u, v, 1, 1}
}

// UV returns a three-dimensional texture coordinate with U, V, and W components.
func UVW(u, v, w float32) Coord {
	return Coord{u, v, w, 1}
}
