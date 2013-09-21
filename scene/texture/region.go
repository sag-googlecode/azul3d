// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

type Region struct {
	U, V, S, T float32
}

func RegionFromPixels(width, height, u, v, s, t int) Region {
	return Region{
		U: float32(u) / float32(width),
		V: float32(v) / float32(height),
		S: float32(s) / float32(width),
		T: float32(t) / float32(height),
	}
}
