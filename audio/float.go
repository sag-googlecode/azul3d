// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package audio

type (
	// F32 represents an 32-bit floating-point linear audio sample in the range
	// of -1 to +1.
	F32 float32

	// F32Samples represents an slice of F32 encoded audio samples.
	F32Samples []F32
)

// Implements Samples interface.
func (p F32Samples) Len() int {
	return len(p)
}

// Implements Samples interface.
func (p F32Samples) At(i int) F64 {
	return F64(p[i])
}

type (
	// F64 represents an 64-bit floating-point linear audio sample in the range
	// of -1 to +1.
	F64 float64

	// F32Samples represents an slice of F32 encoded audio samples.
	F64Samples []F64
)

// Implements Samples interface.
func (p F64Samples) Len() int {
	return len(p)
}

// Implements Samples interface.
func (p F64Samples) At(i int) F64 {
	return p[i]
}
