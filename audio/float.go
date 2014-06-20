// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

type (
	// F32 represents an 32-bit floating-point linear audio sample in the range
	// of -1 to +1.
	F32 float32

	// F32Samples represents an slice of F32 encoded audio samples.
	F32Samples []F32
)

// Implements Buffer interface.
func (p F32Samples) Len() int {
	return len(p)
}

// Implements Buffer interface.
func (p F32Samples) At(i int) F64 {
	return F64(p[i])
}

// Implements Buffer interface.
func (p F32Samples) Set(i int, s F64) {
	p[i] = F32(s)
}

// Implements Buffer interface.
func (p F32Samples) Slice(low, high int) Buffer {
	return p[low:high]
}

// Implements Buffer interface.
func (p F32Samples) Make(length, capacity int) Buffer {
	return make(F32Samples, length, capacity)
}

type (
	// F64 represents an 64-bit floating-point linear audio sample in the range
	// of -1 to +1.
	F64 float64

	// F32Samples represents an slice of F32 encoded audio samples.
	F64Samples []F64
)

// Implements Buffer interface.
func (p F64Samples) Len() int {
	return len(p)
}

// Implements Buffer interface.
func (p F64Samples) At(i int) F64 {
	return p[i]
}

// Implements Buffer interface.
func (p F64Samples) Set(i int, s F64) {
	p[i] = s
}

// Implements Buffer interface.
func (p F64Samples) Slice(low, high int) Buffer {
	return p[low:high]
}

// Implements Buffer interface.
func (p F64Samples) Make(length, capacity int) Buffer {
	return make(F64Samples, length, capacity)
}
