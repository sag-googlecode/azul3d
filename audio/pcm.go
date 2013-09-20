// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package audio

import (
	"math"
)

type (
	// PCM8 represents an 8-bit linear PCM audio sample.
	PCM8 uint8

	// PCM6Samples represents an slice of PCM8 encoded audio samples.
	PCM8Samples []PCM8
)

// Implements Samples interface.
func (p PCM8Samples) Len() int {
	return len(p)
}

// Implements Samples interface.
func (p PCM8Samples) At(i int) F64 {
	return F64(p[i]) / F64(math.MaxUint8)
}

type (
	// PCM16 represents an signed 16-bit linear PCM audio sample.
	PCM16 int16

	// PCM16Samples represents an slice of PCM16 encoded audio samples.
	PCM16Samples []PCM16
)

// Implements Samples interface.
func (p PCM16Samples) Len() int {
	return len(p)
}

// Implements Samples interface.
func (p PCM16Samples) At(i int) F64 {
	return F64(p[i]) / F64(math.MaxInt16)
}

type (
	// PCM32 represents an signed 32-bit linear PCM audio sample.
	PCM32 int32

	// PCM32Samples represents an slice of PCM32 encoded audio samples.
	PCM32Samples []PCM32
)

// Implements Samples interface.
func (p PCM32Samples) Len() int {
	return len(p)
}

// Implements Samples interface.
func (p PCM32Samples) At(i int) F64 {
	return F64(p[i]) / F64(math.MaxInt32)
}
