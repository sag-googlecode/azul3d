// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"math"
)

// Some wrappers around math.* functions that simply do the type conversion for you, just to make
// code look prettier.
//
// Feel free to add anything here you need.

var Pi Real = math.Pi

func Mod(x, y Real) Real {
	return Real(math.Mod(float64(x), float64(y)))
}

func Floor(x Real) Real {
	return Real(math.Floor(float64(x)))
}

func IsNaN(x Real) bool {
	return math.IsNaN(float64(x))
}

func Pow(x, y Real) Real {
	return Real(math.Pow(float64(x), float64(y)))
}

func Abs(x Real) Real {
	return Real(math.Abs(float64(x)))
}

func Acos(x Real) Real {
	return Real(math.Acos(float64(x)))
}

func Atan2(x, y Real) Real {
	return Real(math.Atan2(float64(x), float64(y)))
}

func Asin(x Real) Real {
	return Real(math.Asin(float64(x)))
}

func Cos(x Real) Real {
	return Real(math.Cos(float64(x)))
}

func Sin(x Real) Real {
	return Real(math.Sin(float64(x)))
}

func Sqrt(x Real) Real {
	return Real(math.Sqrt(float64(x)))
}

func Tan(x Real) Real {
	return Real(math.Tan(float64(x)))
}

func Min(a, b Real) Real {
	return Real(math.Min(float64(a), float64(b)))
}

func Max(a, b Real) Real {
	return Real(math.Max(float64(a), float64(b)))
}
