// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"code.google.com/p/azul3d/math"
)

type Vertex struct {
	X, Y, Z math.Real
}

type Normal struct {
	X, Y, Z float32
}

type Tangent struct {
	X, Y, Z float32
}

type Bitangent struct {
	X, Y, Z float32
}

type BoneWeight struct {
	X, Y, Z float32
}
