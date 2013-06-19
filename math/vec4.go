// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// Vec4 represents an vector of four components, X, Y, Z, and W.
type Vec4 struct {
	X, Y, Z, W Real
}

// Copy returns an new 1:1 copy of this Vec4
func (a *Vec4) Copy() *Vec4 {
	return &Vec4{a.X, a.Y, a.Z, a.W}
}

// String returns an string representation of this vector.
func (a *Vec4) String() string {
	return fmt.Sprintf("Vec4(X=%v, Y=%v, Z=%v, W=%v)", a.X, a.Y, a.Z, a.W)
}

// Assign assigns the X, Y, and Z values in this vector to the specified values.
func (a *Vec4) Assign(x, y, z, w Real) {
	a.X = x
	a.Y = y
	a.Z = z
	a.W = w
}

// Fill sets the X, Y, Z and W components of this *Vec4 to the Real, n, parameter.
func (a *Vec4) Fill(n Real) {
	a.X = n
	a.Y = n
	a.Z = n
	a.W = n
}

// Equals tells if this vector is equal to the other vector, by determining if it is within the
// default tolerence for Real equality.
func (a *Vec4) Equals(b *Vec4) bool {
	return a.X.Equals(b.X) && a.Y.Equals(b.Y) && a.Z.Equals(b.Z) && a.W.Equals(b.W)
}

// EqualsTolerence tells if this vector is equal to the other vector, by determining if it is
// within the specified tolerence for Real equality.
func (a *Vec4) EqualsTolerence(b *Vec4, tolerence Real) bool {
	return a.X.EqualsTolerence(b.X, tolerence) && a.Y.EqualsTolerence(b.Y, tolerence) && a.Z.EqualsTolerence(b.Z, tolerence) && a.W.EqualsTolerence(b.W, tolerence)
}

// Negate negates all components of this vector.
func (a *Vec4) Negate() {
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
	a.W = -a.W
}

// Add returns the result of a + b
func (a *Vec4) Add(b *Vec4) *Vec4 {
	return &Vec4{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
		a.W + b.W,
	}
}

// AddScalar returns the result of a + b
func (a *Vec4) AddScalar(b Real) *Vec4 {
	return &Vec4{
		a.X + b,
		a.Y + b,
		a.Z + b,
		a.W + b,
	}
}

// Sub returns the result of a - b
func (a *Vec4) Sub(b *Vec4) *Vec4 {
	return &Vec4{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
		a.W - b.W,
	}
}

// SubScalar returns the result of a - b
func (a *Vec4) SubScalar(b Real) *Vec4 {
	return &Vec4{
		a.X - b,
		a.Y - b,
		a.Z - b,
		a.W - b,
	}
}

// Mul returns the result of a + b
func (a *Vec4) Mul(b *Vec4) *Vec4 {
	return &Vec4{
		a.X * b.X,
		a.Y * b.Y,
		a.Z * b.Z,
		a.W * b.W,
	}
}

// MulScalar returns the result of a * b
func (a *Vec4) MulScalar(b Real) *Vec4 {
	return &Vec4{
		a.X * b,
		a.Y * b,
		a.Z * b,
		a.W * b,
	}
}

// Div returns the result of a + b
func (a *Vec4) Div(b *Vec4) *Vec4 {
	return &Vec4{
		a.X / b.X,
		a.Y / b.Y,
		a.Z / b.Z,
		a.W / b.W,
	}
}

// DivScalar returns the result of a / b
func (a *Vec4) DivScalar(b Real) *Vec4 {
	return &Vec4{
		a.X / b,
		a.Y / b,
		a.Z / b,
		a.W / b,
	}
}

// Dot returns the dot product of the two vectors a and b, respectively.
func (a *Vec4) Dot(b *Vec4) Real {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// LengthSquared returns the squared length of this vector.
func (a *Vec4) LengthSquared() Real {
	return a.Dot(a)
}

// Length returns the length of this vector.
func (a *Vec4) Length() Real {
	return Sqrt(a.LengthSquared())
}

// Normalize normalizes this vector, returns true if it was normalized or false if it was an zero
// length vector.
func (a *Vec4) Normalize() bool {
	lengthSquared := a.LengthSquared()

	if lengthSquared.Equals(0) {
		a.X = 0
		a.Y = 0
		a.Z = 0
		a.W = 0
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
func (a *Vec4) Project(b *Vec4) *Vec4 {
	return b.MulScalar(a.Dot(b) / b.LengthSquared())
}

// Min returns an new vector representing the smaller components of the two vectors.
func (a *Vec4) Min(b *Vec4) *Vec4 {
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

	if a.Z < b.Z {
		cpy.Z = a.Z
	} else {
		cpy.Z = b.Z
	}

	if a.W < b.W {
		cpy.W = a.W
	} else {
		cpy.W = b.W
	}
	return &cpy
}

// Max returns an new vector representing the larger components of the two vectors.
func (a *Vec4) Max(b *Vec4) *Vec4 {
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

	if a.Z > b.Z {
		cpy.Z = a.Z
	} else {
		cpy.Z = b.Z
	}

	if a.W > b.W {
		cpy.W = a.W
	} else {
		cpy.W = b.W
	}
	return &cpy
}

// CompareTolerence sorts the two vectors lexicographically, componentwise.
//
// Returns -1 if this vector sorts before the other one, and returns +1 if it sorts after.
//
// Returns exactly zero if they are equal within the specified tolerence for Real equality.
func (a *Vec4) CompareTolerence(b *Vec4, tolerence Real) int {
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

	if a.Z.EqualsCompeq(b.Z, tolerence) {
		if a.Z < b.Z {
			return -1
		}
		return 1
	}

	if a.W.EqualsCompeq(b.W, tolerence) {
		if a.W < b.W {
			return -1
		}
		return 1
	}

	return 0
}

// Compare is just like CompareTolerence except it uses the default tolerence.
func (a *Vec4) Compare(b *Vec4) int {
	return a.CompareTolerence(b, RealNearZero)
}

// LessThan returns a < b
//
// Also see the Compare() and CompareTolerence() functions.
func (a *Vec4) LessThan(b *Vec4) bool {
	return a.Compare(b) < 0
}

// Lerp returns an new vector representing an linear interpolation between the a and b vectors.
//
// The parameter t is interpolation amount (0.0 - 1.0) between the two vectors.
//
// Short hand for:
//  a.Mul(b.MulScalar(t))
//
func (a *Vec4) Lerp(b *Vec4, t Real) *Vec4 {
	return a.Mul(b.MulScalar(t))
}

// IsNan tells if any components of this vector are not an number.
func (a *Vec4) IsNan() bool {
	return IsNaN(a.X) || IsNaN(a.Y) || IsNaN(a.Z) || IsNaN(a.W)
}

// Quat converts the vector into an quaternion. (No magic -- just an 'type cast')
func (a *Vec4) Quat() *Quat {
	return &Quat{a.X, a.Y, a.Z, a.W}
}

// Transform transforms this vector by the matrix (vector * matrix), and returns the new vector
// result.
//
// This is an fully general operation.
func (v *Vec4) Transform(m *Mat4) *Vec4 {
	return Vector4(
		v.X*m[0][0]+v.Y*m[1][0]+v.Z*m[2][0]+v.W*m[3][0],
		v.X*m[0][1]+v.Y*m[1][1]+v.Z*m[2][1]+v.W*m[3][1],
		v.X*m[0][2]+v.Y*m[1][2]+v.Z*m[2][2]+v.W*m[3][2],
		v.X*m[0][3]+v.Y*m[1][3]+v.Z*m[2][3]+v.W*m[3][3],
	)
}

// Vector4 returns an new *Vec4 with the specified values.
func Vector4(x, y, z, w Real) *Vec4 {
	return &Vec4{x, y, z, w}
}

var (
	Vec4Zero  = Vector4(0, 0, 0, 0)
	Vec4One   = Vector4(1, 1, 1, 1)
	Vec4UnitX = Vector4(1, 0, 0, 0)
	Vec4UnitY = Vector4(0, 1, 0, 0)
	Vec4UnitZ = Vector4(0, 0, 1, 0)
	Vec4UnitW = Vector4(0, 0, 0, 1)
)
