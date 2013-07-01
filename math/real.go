// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !float64

package math

import "math"

// Real represents an 64 bit or 32 bit floating point number, based on weather the build tag
// "float64" is present.
//
// By default, an Real's type is float32. If you specify the build tag "float64" then it's data
// type will be float64.
type Real float32

const (
	RealIsFloat64      = false
	RealSizeBits       = 32
	RealSizeBytes      = 32 / 8
	RealNearZero  Real = math.SmallestNonzeroFloat32
	MaxReal = math.MaxFloat32
)
