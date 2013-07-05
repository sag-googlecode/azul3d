// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// Vec3 represents an vector of three components, X, Y, and Z.
type Vec3 struct {
	X, Y, Z Real
}

// Copy returns an new 1:1 copy of this Vec3
func (a *Vec3) Copy() *Vec3 {
	return &Vec3{a.X, a.Y, a.Z}
}

// String returns an string representation of this vector.
func (a *Vec3) String() string {
	return fmt.Sprintf("Vec3(X=%v, Y=%v, Z=%v)", a.X, a.Y, a.Z)
}

// Assign assigns the X, Y, and Z values in this vector to the specified values.
func (a *Vec3) Assign(x, y, z Real) {
	a.X = x
	a.Y = y
	a.Z = z
}

// Fill sets the X, Y and Z components of this *Vec3 to the Real, n, parameter.
func (a *Vec3) Fill(n Real) {
	a.X = n
	a.Y = n
	a.Z = n
}

// Equals tells if this vector is equal to the other vector, by determining if it is within the
// default tolerence for Real equality.
func (a *Vec3) Equals(b *Vec3) bool {
	return a.X.Equals(b.X) && a.Y.Equals(b.Y) && a.Z.Equals(b.Z)
}

// EqualsTolerence tells if this vector is equal to the other vector, by determining if it is
// within the specified tolerence for Real equality.
func (a *Vec3) EqualsTolerence(b *Vec3, tolerence Real) bool {
	return a.X.EqualsTolerence(b.X, tolerence) && a.Y.EqualsTolerence(b.Y, tolerence) && a.Z.EqualsTolerence(b.Z, tolerence)
}

// Negate negates all components of this vector in place.
func (a *Vec3) Negate() {
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
}

// Add returns the result of a + b
func (a *Vec3) Add(b *Vec3) *Vec3 {
	return &Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

// AddScalar returns the result of a + b
func (a *Vec3) AddScalar(b Real) *Vec3 {
	return &Vec3{
		a.X + b,
		a.Y + b,
		a.Z + b,
	}
}

// Sub returns the result of a - b
func (a *Vec3) Sub(b *Vec3) *Vec3 {
	return &Vec3{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

// SubScalar returns the result of a - b
func (a *Vec3) SubScalar(b Real) *Vec3 {
	return &Vec3{
		a.X - b,
		a.Y - b,
		a.Z - b,
	}
}

// Mul returns the result of a * b
func (a *Vec3) Mul(b *Vec3) *Vec3 {
	return &Vec3{
		a.X * b.X,
		a.Y * b.Y,
		a.Z * b.Z,
	}
}

// MulScalar returns the result of a * b
func (a *Vec3) MulScalar(b Real) *Vec3 {
	return &Vec3{
		a.X * b,
		a.Y * b,
		a.Z * b,
	}
}

// Div returns the result of a / b
func (a *Vec3) Div(b *Vec3) *Vec3 {
	return &Vec3{
		a.X / b.X,
		a.Y / b.Y,
		a.Z / b.Z,
	}
}

// DivScalar returns the result of a / b
func (a *Vec3) DivScalar(b Real) *Vec3 {
	return &Vec3{
		a.X / b,
		a.Y / b,
		a.Z / b,
	}
}

// Dot returns the dot product of the two vectors a and b, respectively.
func (a *Vec3) Dot(b *Vec3) Real {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// LengthSquared returns the squared length of this vector.
func (a *Vec3) LengthSquared() Real {
	return a.Dot(a)
}

// Length returns the length of this vector.
func (a *Vec3) Length() Real {
	return Sqrt(a.LengthSquared())
}

// Normalize normalizes this vector, returns true if it was normalized or false if it was an zero
// length vector.
func (a *Vec3) Normalize() bool {
	lengthSquared := a.LengthSquared()

	if lengthSquared.Equals(0) {
		a.X = 0
		a.Y = 0
		a.Z = 0
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
func (a *Vec3) Project(b *Vec3) *Vec3 {
	return b.MulScalar(a.Dot(b) / b.LengthSquared())
}

// Min returns an new vector representing the smaller components of the two vectors.
func (a *Vec3) Min(b *Vec3) *Vec3 {
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

	return &cpy
}

// Max returns an new vector representing the larger components of the two vectors.
func (a *Vec3) Max(b *Vec3) *Vec3 {
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

	return &cpy
}

// CompareTolerence sorts the two vectors lexicographically, componentwise.
//
// Returns -1 if this vector sorts before the other one, and returns +1 if it sorts after.
//
// Returns exactly zero if they are equal within the specified tolerence for Real equality.
func (a *Vec3) CompareTolerence(b *Vec3, tolerence Real) int {
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

	return 0
}

// Compare is just like CompareTolerence except it uses the default tolerence.
func (a *Vec3) Compare(b *Vec3) int {
	return a.CompareTolerence(b, RealNearZero)
}

// LessThan returns a < b
//
// Also see the Compare() and CompareTolerence() functions.
func (a *Vec3) LessThan(b *Vec3) bool {
	return a.Compare(b) < 0
}

// Lerp returns an new vector representing an linear interpolation between the a and b vectors.
//
// The parameter t is interpolation amount (0.0 - 1.0) between the two vectors.
//
// Short hand for:
//  a.Mul(b.MulScalar(t))
//
func (a *Vec3) Lerp(b *Vec3, t Real) *Vec3 {
	return a.Mul(b.MulScalar(t))
}

// IsNan tells if any components of this vector are not an number.
func (a *Vec3) IsNan() bool {
	return IsNaN(a.X) || IsNaN(a.Y) || IsNaN(a.Z)
}

// Cross returns the cross product of the two vectors.
func (a *Vec3) Cross(b *Vec3) *Vec3 {
	return &Vec3{
		a.Y*b.Z - b.Y*a.Z,
		b.X*a.Z - a.X*b.Z,
		a.X*b.Y - b.X*a.Y,
	}
}

func standardizedRotation(angleDegrees Real) Real {
	if angleDegrees < 0 {
		angleDegrees = 360 - Mod(angleDegrees*-1.0, 360)
	} else {
		angleDegrees = Mod(angleDegrees, 360)
	}

	// We now have value in range of 0.0 to 359.99999

	if angleDegrees < 180 {
		return angleDegrees
	}
	return angleDegrees - 360.0
}

// StandardizedHpr() tries to un-spin the hpr to a standard form. This function assumes that 0 and
// 360 are the same, as is 720 and -360. Also 180 and -180 are the same. Another example is -90 and
// 270.
//
// Each element will be in the range -180 to 179.99999. The use of this function should be strictly
// used for human readable output, not for comparison.
//
// See also Quat.SameDirection()
func (a *Vec3) StandardizedHpr() *Vec3 {
	return &Vec3{
		standardizedRotation(a.X),
		standardizedRotation(a.Y),
		standardizedRotation(a.Z),
	}
}

// Angle returns the unsigned angle between vectors a and b, expressed in radians.
//
// Both vectors are expected to be normalized already before calling this function.
func (a *Vec3) Angle(b *Vec3) Real {
	var n Real
	if a.Dot(b) < 0 {
		n = a.Add(b).Length() / 2
		return Pi - 2.0*Asin(Min(n, 1))
	}
	n = a.Sub(b).Length() / 2
	return 2.0 * Asin(Min(n, 1))
}

// SignedAngle returns the signed angle between vectors a and b, expressed in radians.
//
// The angle is positive if the rotation from a to b is clockwise when looking in the direction of
// the ref vector.
//
// Both vectors are expected to be normalized already (but not the ref vector) before calling this
// function.
func (a *Vec3) SignedAngle(b, ref *Vec3) Real {
	angle := a.Angle(b)
	if a.Cross(b).Dot(ref) < 0.0 {
		angle = -angle
	}
	return angle
}

// TransformMat3 transforms this point vector by the matrix (vector * matrix), and returns the new
// vector result.
//
// Can operate on orthonormal transformation matrices.
func (v *Vec3) TransformMat3(m *Mat3) *Vec3 {
	return Vector3(
		v.X*m[0][0]+v.Y*m[1][0]+v.Z*m[2][0],
		v.X*m[0][1]+v.Y*m[1][1]+v.Z*m[2][1],
		v.X*m[0][2]+v.Y*m[1][2]+v.Z*m[2][2],
	)
}

// TransformGeneralMat3 transforms this vector by the matrix (vector * matrix) without translation
// component, and returns the new vector result, as an fully general operation.
func (v *Vec3) TransformGeneralMat3(m *Mat3) *Vec3 {
	i := new(Mat3)
	i.StoreInvert(m)
	return v.TransformMat3(i)
}

// TransformMat4 transforms this point vector by the affine transformation matrix (vector * matrix)
// and returns the new vector result.
//
// The m parameter must be an affine transformation matrix.
func (v *Vec3) TransformMat4(m *Mat4) *Vec3 {
	return Vector3(
		v.X*m[0][0]+v.Y*m[1][0]+v.Z*m[2][0]+m[3][0],
		v.X*m[0][1]+v.Y*m[1][1]+v.Z*m[2][1]+m[3][1],
		v.X*m[0][2]+v.Y*m[1][2]+v.Z*m[2][2]+m[3][2],
	)
}

// TransformVecMat4 transforms this vector (without translation component) by the orthonormal
// matrix and returns the result.
func (v *Vec3) TransformVecMat4(m *Mat4) *Vec3 {
	return Vector3(
		v.X*m[0][0]+v.Y*m[1][0]+v.Z*m[2][0],
		v.X*m[0][1]+v.Y*m[1][1]+v.Z*m[2][1],
		v.X*m[0][2]+v.Y*m[1][2]+v.Z*m[2][2],
	)
}

// TransformGeneralMat4 transforms this vector by the matrix (vector * matrix) without translation
// component, and returns the new vector result, as an fully general operation.
func (v *Vec3) TransformGeneralMat4(m *Mat4) *Vec3 {
	i := new(Mat3)
	i.StoreInvertTransposeMat4(m)
	return v.TransformMat3(i)
}

// PureImaginaryQuat returns an new *Quat using the vectors X, Y, and Z components for the
// quaternion's respective X, Y, and Z components (leaving W=0 in the quaternion).
//
// Short hand for:
//
//  Quaternion(0, v.X, v.Y, v.Z)
func (v *Vec3) PureImaginaryQuat() *Quat {
	return &Quat{0, v.X, v.Y, v.Z}
}

// HprToXyz returns this:
//
//  return math.Vector3(v.Y, v.Z, v.X)
//
func (v *Vec3) HprToXyz() *Vec3 {
	return &Vec3{v.Y, v.Z, v.X}
}

// XyzToHpr returns this:
//
//  return math.Vector3(v.Z, v.X, v.Y)
//
func (v *Vec3) XyzToHpr() *Vec3 {
	return &Vec3{v.Z, v.X, v.Y}
}

// Radians returns an new vector representing each element of the vector v, in degrees, converted
// to radians.
func (v *Vec3) Radians() *Vec3 {
	return &Vec3{v.X.Radians(), v.Y.Radians(), v.Z.Radians()}
}

// Degrees returns an new vector representing each element of the vector v, in radians, converted
// to degrees.
func (v *Vec3) Degrees() *Vec3 {
	return &Vec3{v.X.Degrees(), v.Y.Degrees(), v.Z.Degrees()}
}

// Clamp clamps each component of this vector to the specified min and max
// values.
func (v *Vec3) Clamp(min, max Real) *Vec3 {
	return Vector3(
		v.X.Clamp(min, max),
		v.Y.Clamp(min, max),
		v.Z.Clamp(min, max),
	)
}

// Vector3 returns an new *Vec3 with the specified values.
func Vector3(x, y, z Real) *Vec3 {
	return &Vec3{x, y, z}
}

var (
	Vec3Zero  = Vector3(0, 0, 0)
	Vec3One   = Vector3(1, 1, 1)
	Vec3UnitX = Vector3(1, 0, 0)
	Vec3UnitY = Vector3(0, 1, 0)
	Vec3UnitZ = Vector3(0, 0, 1)
)
