// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package audio

import (
	"fmt"
)

// Type represents an single audio sample encoding type.
type Type uint8

// NewBuffer() returns a new buffer of the predefined audio type and the given
// size.
//
// E.g. TYPE_PCM8 -> make(PCM8Samples, size)
func (t Type) NewBuffer(size int) Buffer {
	switch t {
	case TYPE_PCM8:
		return make(PCM8Samples, size)
	case TYPE_PCM16:
		return make(PCM16Samples, size)
	case TYPE_PCM32:
		return make(PCM32Samples, size)
	case TYPE_F32:
		return make(F32Samples, size)
	case TYPE_F64:
		return make(F64Samples, size)
	case TYPE_MULAW:
		return make(MuLawSamples, size)
	case TYPE_ALAW:
		return make(ALawSamples, size)
	}
	return nil
}

// String returns a string representation of this audio sample encoding type.
//
// E.g. TYPE_PCM8 -> "TYPE_PCM8"
func (t Type) String() string {
	switch t {
	case TYPE_PCM8:
		return "TYPE_PCM8"
	case TYPE_PCM16:
		return "TYPE_PCM16"
	case TYPE_PCM32:
		return "TYPE_PCM32"
	case TYPE_F32:
		return "TYPE_F32"
	case TYPE_F64:
		return "TYPE_F64"
	case TYPE_MULAW:
		return "TYPE_MULAW"
	case TYPE_ALAW:
		return "TYPE_ALAW"
	}
	return fmt.Sprintf("Type(%v)", t)
}

const (
	// Type for PCM8 encoded audio samples.
	TYPE_PCM8 Type = iota

	// Type for PCM16 encoded audio samples.
	TYPE_PCM16

	// Type for PCM32 encoded audio samples.
	TYPE_PCM32

	// Type for F32 encoded audio samples.
	TYPE_F32

	// Type for F64 encoded audio samples.
	TYPE_F64

	// Type for MuLaw encoded audio samples.
	TYPE_MULAW

	// Type for ALaw encoded audio samples.
	TYPE_ALAW
)

// Convert converts the specified buffer to one with the specified target audio
// sample type.
//
// If the source buffer is already the specified target type, the source buffer
// is simply returned.
func Convert(src Buffer, target Type) Buffer {
	switch target {

	case TYPE_PCM16:
		_, ok := src.(PCM16Samples)
		if ok {
			return src
		}
		converted := target.NewBuffer(src.Len())
		Copy(converted, src)
		return converted

	case TYPE_PCM32:
		_, ok := src.(PCM32Samples)
		if ok {
			return src
		}
		converted := target.NewBuffer(src.Len())
		Copy(converted, src)
		return converted

	case TYPE_F32:
		_, ok := src.(F32Samples)
		if ok {
			return src
		}
		converted := target.NewBuffer(src.Len())
		Copy(converted, src)
		return converted

	case TYPE_F64:
		_, ok := src.(F64Samples)
		if ok {
			return src
		}
		converted := target.NewBuffer(src.Len())
		Copy(converted, src)
		return converted

	case TYPE_MULAW:
		_, ok := src.(MuLawSamples)
		if ok {
			return src
		}
		converted := target.NewBuffer(src.Len())
		Copy(converted, src)
		return converted

	case TYPE_ALAW:
		_, ok := src.(ALawSamples)
		if ok {
			return src
		}
		converted := target.NewBuffer(src.Len())
		Copy(converted, src)
		return converted
	}

	panic("Convert(): Unknown target type!")
}

// Buffer is a generic audio buffer, it can conceptually be thought of as a
// slice of some audio encoding type.
type Buffer interface {
	// Len returns the number of elements in the buffer.
	//
	// Equivilent slice syntax:
	//
	//  len(b)
	Len() int

	// Set sets the specified index in the buffer to the specified F64 encoded
	// audio sample, s.
	//
	// If the buffer's audio samples are not stored in F64 encoding, then the
	// sample should be converted to the buffer's internal format and then
	// stored.
	//
	// Just like slices, buffer indices must be non-negative; and no greater
	// than (Len() - 1), or else a panic may occur.
	//
	// Equivilent slice syntax:
	//
	//  b[index] = s
	//   -> b.Set(index, s)
	//
	Set(index int, s F64)

	// At returns the F64 encoded audio sample at the specified index in the
	// buffer.
	//
	// If the buffer's audio samples are not stored in F64 encoding, then the
	// sample should be converted to F64 encoding, and subsequently returned.
	//
	// Just like slices, buffer indices must be non-negative; and no greater
	// than (Len() - 1), or else a panic may occur.
	//
	// Equivilent slice syntax:
	//
	//  b[index]
	//   -> b.At(index)
	//
	At(index int) F64

	// Slice returns a new slice of the buffer, using the low and high
	// parameters.
	//
	// Equivilent slice syntax:
	//
	//  b[low:high]
	//   -> b.Slice(low, high)
	//
	//  b[2:]
	//   -> b.Slice(2, a.Len())
	//
	//  b[:3]
	//   -> b.Slice(0, 3)
	//
	//  b[:]
	//   -> b.Slice(0, a.Len())
	//
	Slice(low, high int) Buffer
}

// Copy copies copies audio samples from the source buffer to the destination
// buffer. Returns the number of elements copied, which is the minimum of
// the dst.Len() and src.Len() values.
func Copy(dst Buffer, src Buffer) int {
	var i int
	for i = 0; i < src.Len() && i < dst.Len(); i++ {
		dst.Set(i, src.At(i))
	}
	return i
}
