// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// Vec2 represents an vector of two components, X and Y.
type Vec2 struct {
	X, Y Real
}

// Copy returns an new 1:1 copy of this Vec2
func (a *Vec2) Copy() *Vec2 {
	return &Vec2{a.X, a.Y}
}

// String returns an string representation of this vector.
func (a *Vec2) String() string {
	return fmt.Sprintf("Vec2(X=%v, Y=%v)", a.X, a.Y)
}

// Assign assigns the X, and Y values in this vector to the specified values.
func (a *Vec2) Assign(x, y Real) {
	a.X = x
	a.Y = y
}

// Fill sets the X and Y components of this *Vec2 to the Real, n, parameter.
func (a *Vec2) Fill(n Real) {
	a.X = n
	a.Y = n
}

// Equals tells if this vector is equal to the other vector, by determining if it is within the
// default tolerence for Real equality.
//
// Equality against nil is always false.
func (a *Vec2) Equals(b *Vec2) bool {
	if b == nil {
		return false
	}
	return a.X.Equals(b.X) && a.Y.Equals(b.Y)
}

// EqualsTolerence tells if this vector is equal to the other vector, by determining if it is
// within the specified tolerence for Real equality.
//
// Equality against nil is always false.
func (a *Vec2) EqualsTolerence(b *Vec2, tolerence Real) bool {
	if b == nil {
		return false
	}
	return a.X.EqualsTolerence(b.X, tolerence) && a.Y.EqualsTolerence(b.Y, tolerence)
}

// Negate negates all components of this vector.
func (a *Vec2) Negate() {
	a.X = -a.X
	a.Y = -a.Y
}

// Add returns the result of a + b
func (a *Vec2) Add(b *Vec2) *Vec2 {
	return &Vec2{
		a.X + b.X,
		a.Y + b.Y,
	}
}

// AddScalar returns the result of a + b
func (a *Vec2) AddScalar(b Real) *Vec2 {
	return &Vec2{
		a.X + b,
		a.Y + b,
	}
}

// Sub returns the result of a - b
func (a *Vec2) Sub(b *Vec2) *Vec2 {
	return &Vec2{
		a.X - b.X,
		a.Y - b.Y,
	}
}

// SubScalar returns the result of a - b
func (a *Vec2) SubScalar(b Real) *Vec2 {
	return &Vec2{
		a.X - b,
		a.Y - b,
	}
}

// Mul returns the result of a * b
func (a *Vec2) Mul(b *Vec2) *Vec2 {
	return &Vec2{
		a.X * b.X,
		a.Y * b.Y,
	}
}

// MulScalar returns the result of a * b
func (a *Vec2) MulScalar(b Real) *Vec2 {
	return &Vec2{
		a.X * b,
		a.Y * b,
	}
}

// Div returns the result of a / b
func (a *Vec2) Div(b *Vec2) *Vec2 {
	return &Vec2{
		a.X / b.X,
		a.Y / b.Y,
	}
}

// DivScalar returns the result of a / b
func (a *Vec2) DivScalar(b Real) *Vec2 {
	return &Vec2{
		a.X / b,
		a.Y / b,
	}
}

// Dot returns the dot product of the two vectors a and b, respectively.
func (a *Vec2) Dot(b *Vec2) Real {
	return a.X*b.X + a.Y*b.Y
}

// LengthSquared returns the squared length of this vector.
func (a *Vec2) LengthSquared() Real {
	return a.Dot(a)
}

// Length returns the length of this vector.
func (a *Vec2) Length() Real {
	return Sqrt(a.LengthSquared())
}

// Normalize normalizes this vector, returns true if it was normalized or false if it was an zero
// length vector.
func (a *Vec2) Normalize() bool {
	lengthSquared := a.LengthSquared()

	if lengthSquared.Equals(0) {
		a.X = 0
		a.Y = 0
		return false

	} else if lengthSquared.EqualsTolerence(1.0, RealNearZero*RealNearZero) {
		length := Sqrt(lengthSquared)
		*a = *a.DivScalar(length)
		return true
	}

	return true
}

// Project returns an new vector representing the projection of this vector onto the other one. The
// resulting vector will be a scalar multiple of onto.
func (a *Vec2) Project(b *Vec2) *Vec2 {
	return b.MulScalar(a.Dot(b) / b.LengthSquared())
}

// Min returns an new vector representing the smaller components of the two vectors.
func (a *Vec2) Min(b *Vec2) *Vec2 {
	cpy := *a

	if a.X < b.X {
		cpy.X = a.X
	} else {
		cpy.X = b.X
	}

	if a.Y < b.Y {
		cpy.Y = a.Y
	} else {
		cpy.Y = b.Y
	}

	return &cpy
}

// Max returns an new vector representing the larger components of the two vectors.
func (a *Vec2) Max(b *Vec2) *Vec2 {
	cpy := *a

	if a.X > b.X {
		cpy.X = a.X
	} else {
		cpy.X = b.X
	}

	if a.Y > b.Y {
		cpy.Y = a.Y
	} else {
		cpy.Y = b.Y
	}

	return &cpy
}

// CompareTolerence sorts the two vectors lexicographically, componentwise.
//
// Returns -1 if this vector sorts before the other one, and returns +1 if it sorts after.
//
// Returns exactly zero if they are equal within the specified tolerence for Real equality.
func (a *Vec2) CompareTolerence(b *Vec2, tolerence Real) int {
	if a.X.EqualsCompeq(b.X, tolerence) {
		if a.X < b.X {
			return -1
		}
		return 1
	}

	if a.Y.EqualsCompeq(b.Y, tolerence) {
		if a.Y < b.Y {
			return -1
		}
		return 1
	}

	return 0
}

// Compare is just like CompareTolerence except it uses the default tolerence.
func (a *Vec2) Compare(b *Vec2) int {
	return a.CompareTolerence(b, RealNearZero)
}

// LessThan returns a < b
//
// Also see the Compare() and CompareTolerence() functions.
func (a *Vec2) LessThan(b *Vec2) bool {
	return a.Compare(b) < 0
}

// Lerp returns an new vector representing an linear interpolation between the a and b vectors.
//
// The parameter t is interpolation amount (0.0 - 1.0) between the two vectors.
//
// Short hand for:
//  a.Mul(b.MulScalar(t))
//
func (a *Vec2) Lerp(b *Vec2, t Real) *Vec2 {
	return a.Mul(b.MulScalar(t))
}

// IsNan tells if any components of this vector are not an number.
func (a *Vec2) IsNan() bool {
	return IsNaN(a.X) || IsNaN(a.Y)
}

// TransformVec2 transforms a 2-component point vector (without translation component) and returns
// the new vector result.
//
// This function assumes that the matrix is an affine transformation.
func (v *Vec2) TransformVec2(m *Mat3) *Vec2 {
	return Vector2(
		v.X*m[0][0]+v.Y*m[1][0],
		v.X*m[0][1]+v.Y*m[1][1],
	)
}

// TransformPointVec2 transforms a 2-component point vector (with translation component) and
// returns the new vector result.
//
// This function assumes that the matrix is an affine transformation.
func (v *Vec2) TransformPointVec2(m *Mat3) *Vec2 {
	return Vector2(
		v.X*m[0][0]+v.Y*m[1][0]+m[2][0],
		v.X*m[0][1]+v.Y*m[1][1]+m[2][1],
	)
}

// Vector2 returns an new *Vec2 with the specified values.
func Vector2(x, y Real) *Vec2 {
	return &Vec2{x, y}
}

var (
	Vec2Zero  = Vector2(0, 0)
	Vec2One   = Vector2(1, 1)
	Vec2UnitX = Vector2(1, 0)
	Vec2UnitY = Vector2(0, 1)
)
