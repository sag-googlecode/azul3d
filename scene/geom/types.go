// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

// Vertex represents a single vertex.
type Vertex struct {
	X, Y, Z float32
}

// Normal represents a single surface normal.
type Normal struct {
	X, Y, Z float32
}

// Tangent represents a single tangent vector.
type Tangent struct {
	X, Y, Z float32
}

// Bitangent represents a single bitangent (sometimes known as binormal) vector.
type Bitangent struct {
	X, Y, Z float32
}

// BoneWeight represents a single boneweight.
type BoneWeight struct {
	X, Y, Z float32
}
