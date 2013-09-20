// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package audio

import (
	"math"
)

// Samples is a the generic interface between audio samples of different types
// and expressions, e.g. any type who has an Len() and At() methods can be used
// as audio samples.
//
// This is an important concept as audio samples can therefor be converted
// between one another's types (although sometimes lossily).
type Samples interface {
	// Len returns the length of the internal slice.
	Len() int

	// At returns the 64-bit floating point audio sample at index i.
	//
	// If samples are not internally stored in F64 encoding, they should be
	// converted on-the-fly (even if lossily).
	At(i int) F64
}

// Type represents an single audio sample encoding type.
type Type uint8

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

// Convert converts the specified audio samples to the target sample encoding
// type, and returns the newly encoded audio samples.
//
// Convert will simply return the src parameter if it's type matches the target
// type (as such, the returned audio samples may simply be src).
func Convert(src Samples, target Type) Samples {
	switch target {
	case TYPE_PCM8:
		_, ok := src.(PCM8Samples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(PCM8Samples, ln)
		for i := 0; i < ln; i++ {
			f := src.At(i)
			f += 1.0
			f /= 2.0
			converted[i] = PCM8(f * F64(math.MaxUint8))
		}
		return converted

	case TYPE_PCM16:
		_, ok := src.(PCM16Samples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(PCM16Samples, ln)
		for i := 0; i < ln; i++ {
			converted[i] = PCM16(src.At(i) * F64(math.MaxInt16))
		}
		return converted

	case TYPE_PCM32:
		_, ok := src.(PCM32Samples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(PCM32Samples, ln)
		for i := 0; i < ln; i++ {
			converted[i] = PCM32(src.At(i) * F64(math.MaxInt32))
		}
		return converted

	case TYPE_F32:
		_, ok := src.(F32Samples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(F32Samples, ln)
		for i := 0; i < ln; i++ {
			converted[i] = F32(src.At(i))
		}
		return converted

	case TYPE_F64:
		_, ok := src.(F64Samples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(F64Samples, ln)
		for i := 0; i < ln; i++ {
			converted[i] = src.At(i)
		}
		return converted

	case TYPE_MULAW:
		_, ok := src.(MuLawSamples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(MuLawSamples, ln)
		for i := 0; i < ln; i++ {
			p16 := PCM16(src.At(i) * F64(math.MaxInt16))
			converted[i] = PCM16ToMuLaw(p16)
		}
		return converted

	case TYPE_ALAW:
		_, ok := src.(ALawSamples)
		if ok {
			return src
		}

		ln := src.Len()
		converted := make(ALawSamples, ln)
		for i := 0; i < ln; i++ {
			p16 := PCM16(src.At(i) * F64(math.MaxInt16))
			converted[i] = PCM16ToALaw(p16)
		}
		return converted
	}

	panic("Convert(): Unknown target type!")
}
