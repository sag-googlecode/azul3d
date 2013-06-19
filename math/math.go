// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package math implements general 3D linear math data types and operations.
//
// Many algorithms found inside this package are based off ones found in Panda3D's linmath library,
// whose code is licensed under an Modified BSD License described in the License.txt file.
package math

var realTolerence Real = RealNearZero

// SetTolerence sets the default tolerence for Real equality in Equals() operations.
func SetTolerence(tolerence Real) {
	realTolerence = tolerence
}

// Tolerence returns the default tolerence for Real equality in Equals() operations.
func Tolerence() Real {
	return realTolerence
}

// Clamp returns an real clamped to the range of min.. to max.
func (a Real) Clamp(min, max Real) Real {
	return Max(Min(a, max), min)
}

// Equals tells if a is equal to b within the default tolerence for Real equality.
func (a Real) Equals(b Real) bool {
	if Abs(a-b) < realTolerence {
		return true
	}
	return false
}

// EqualsTolerence tells if a is equal to b within the specified tolerence for Real equality.
func (a Real) EqualsTolerence(b, tolerence Real) bool {
	if Abs(a-b) < tolerence {
		return true
	}
	return false
}

// EqualsCompeq tells if a is equal to b within the specified tolerence for Real equality,
// unlike EqualsTolerence the transitive principle is guaranteed:
//
// a.EqualsCompeq(b, t) && b.EqualsCompeq(c, t) implies a.EqualsCompeq(c, t).
func (a Real) EqualsCompeq(b, tolerence Real) bool {
	return Floor(a/tolerence+0.5) == Floor(b/tolerence+0.5)
}

// Converts from Degrees to Radians
func (degrees Real) Radians() Real {
	return Pi * degrees / 180.0
}

// Converts from Radians to Degrees
func (radians Real) Degrees() Real {
	return radians * (180.0 / Pi)
}
