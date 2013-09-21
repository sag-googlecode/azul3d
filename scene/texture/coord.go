// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

type Coord struct {
	U, V, W, Z float32
}

func UV(u, v float32) Coord {
	return Coord{u, v, 1, 1}
}

func UVW(u, v, w float32) Coord {
	return Coord{u, v, w, 1}
}
