// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// Mat3 represents an 3x3 matrix, indices are in m[row][column] order.
type Mat3 [3][3]Real

// Copy creates an new *Mat4, which is an 1:1 clone of this matrix.
func (a *Mat3) Copy() *Mat3 {
	r := new(Mat3)

	r[0][0] = a[0][0]
	r[0][1] = a[0][1]
	r[0][2] = a[0][2]

	r[1][0] = a[1][0]
	r[1][1] = a[1][1]
	r[1][2] = a[1][2]

	r[2][0] = a[2][0]
	r[2][1] = a[2][1]
	r[2][2] = a[2][2]

	return r
}

// String returns an string representation of this matrix.
func (a *Mat3) String() string {
	return fmt.Sprintf("Mat3(\n    %v, %v, %v\n    %v, %v, %v\n    %v, %v, %v\n)", a[0][0], a[0][1], a[0][2], a[1][0], a[1][1], a[1][2], a[2][0], a[2][1], a[2][2])
}

// Assign assigns the specified values to this matrix
func (a *Mat3) Assign(m00, m01, m02, m10, m11, m12, m20, m21, m22 Real) {
	a[0][0] = m00
	a[0][1] = m01
	a[0][2] = m02

	a[1][0] = m10
	a[1][1] = m11
	a[1][2] = m12

	a[2][0] = m20
	a[2][1] = m21
	a[2][2] = m22
}

// Fill fills this matrix with the value specified. Useful sometimes for initializing to zero, etc.
func (a *Mat3) Fill(v Real) {
	a.Assign(
		v, v, v,
		v, v, v,
		v, v, v,
	)
}

// SetRow sets the values in the specified matrix row to the values in the specified
// three-component vector.
//
// The row parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) SetRow(row int, values *Vec3) {
	a[row][0] = values.X
	a[row][1] = values.Y
	a[row][2] = values.Z
}

// Row returns the values in the specified matrix row as an three-component vector.
//
// The row parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) Row(row int) *Vec3 {
	return &Vec3{a[row][0], a[row][1], a[row][2]}
}

// SetRowVec2 sets the values in the specified matrix row to the values in the specified
// three-component vector, leaving the third element in the row untouched.
//
// The row parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) SetRowVec2(row int, values *Vec2) {
	a[row][0] = values.X
	a[row][1] = values.Y
}

// RowVec2 returns the values in the specified matrix row as an two-component vector, the third
// element of the row is ignored.
//
// The row parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) RowVec2(row int) *Vec2 {
	return &Vec2{a[row][0], a[row][1]}
}

// SetCol sets the values in the specified matrix column to the values in the specified
// three-component vector.
//
// The column parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) SetCol(column int, values *Vec3) {
	a[0][column] = values.X
	a[1][column] = values.Y
	a[2][column] = values.Z
}

// Col returns the values in the specified matrix column as an three-component vector.
//
// The column parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) Col(column int) *Vec3 {
	return &Vec3{a[0][column], a[1][column], a[2][column]}
}

// SetColVec2 sets the values in the specified matrix column to the values in the specified
// two-component vector, leaving the third element in the column untouched.
//
// The column parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) SetColVec2(column int, values *Vec2) {
	a[0][column] = values.X
	a[1][column] = values.Y
}

// ColVec2 returns the values in the specified matrix column as an two-component vector, ignoring
// the third element of the column.
//
// The column parameter must be 0, 1, or 2 or else an panic will occur.
func (a *Mat3) ColVec2(column int) *Vec2 {
	return &Vec2{a[0][column], a[1][column]}
}

// EqualsTolerence tells weather a is memberwise equal to b within the given tolerence for Real
// equality.
func (a *Mat3) EqualsTolerence(b *Mat3, tolerence Real) bool {
	return a[0][0].EqualsTolerence(b[0][0], tolerence) &&
		a[0][1].EqualsTolerence(b[0][1], tolerence) &&
		a[0][2].EqualsTolerence(b[0][2], tolerence) &&

		a[1][0].EqualsTolerence(b[1][0], tolerence) &&
		a[1][1].EqualsTolerence(b[1][1], tolerence) &&
		a[1][2].EqualsTolerence(b[1][2], tolerence) &&

		a[2][0].EqualsTolerence(b[2][0], tolerence) &&
		a[2][1].EqualsTolerence(b[2][1], tolerence) &&
		a[2][2].EqualsTolerence(b[2][2], tolerence)
}

// Equals tells weather a is memberwise equal to b within the default tolerence for Real equality.
func (a *Mat3) Equals(b *Mat3) bool {
	return a.EqualsTolerence(b, RealNearZero)
}

// CompareTolerence sorts the two matrices lexicographically, componentwise.
//
// Returns -1 if matrix a sorts before matrix b.
//
// Returns 1 if matrix a sorts after b.
//
// Returns 0 if matrix a and b are equal (within the given tolerence).
func (a *Mat3) CompareTolerence(b *Mat3, tolerence Real) int {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if !a[r][c].EqualsCompeq(b[r][c], tolerence) {
				if a[r][c] < b[r][c] {
					return -1
				}
				return 1
			}
		}
	}
	return 0
}

// Compare works just like CompareTolerence() except it uses the default tolerence for Real equality.
func (a *Mat3) Compare(b *Mat3) int {
	return a.CompareTolerence(b, RealNearZero)
}

// LessThan returns a < b
//
// Also see the Compare() and CompareTolerence() functions.
func (a *Mat3) LessThan(b *Mat3) bool {
	return a.Compare(b) < 0
}

// AddScalar performs in-place memberwise addition (a += s).
//
// The matrix itself, a, is returned such that you can perform non in-place addition using:
//
//  result := a.Copy().AddScalar()
//
func (a *Mat3) AddScalar(s Real) *Mat3 {
	a.Assign(
		a[0][0]+s, a[0][1]+s, a[0][2]+s,
		a[1][0]+s, a[1][1]+s, a[1][2]+s,
		a[2][0]+s, a[2][1]+s, a[2][2]+s,
	)
	return a
}

// SubScalar performs in-place memberwise subtraction (a -= s).
//
// The matrix itself, a, is returned such that you can perform non in-place subtraction using:
//
//  result := a.Copy().SubScalar()
//
func (a *Mat3) SubScalar(s Real) *Mat3 {
	a.Assign(
		a[0][0]-s, a[0][1]-s, a[0][2]-s,
		a[1][0]-s, a[1][1]-s, a[1][2]-s,
		a[2][0]-s, a[2][1]-s, a[2][2]-s,
	)
	return a
}

// MulScalar performs in-place memberwise multiplication (a *= s).
//
// The matrix itself, a, is returned such that you can perform non in-place muliplication using:
//
//  result := a.Copy().MulScalar()
//
func (a *Mat3) MulScalar(s Real) *Mat3 {
	a.Assign(
		a[0][0]*s, a[0][1]*s, a[0][2]*s,
		a[1][0]*s, a[1][1]*s, a[1][2]*s,
		a[2][0]*s, a[2][1]*s, a[2][2]*s,
	)
	return a
}

// DivScalar performs in-place memberwise division (a /= s).
//
// The matrix itself, a, is returned such that you can perform non in-place division using:
//
//  result := a.Copy().DivScalar()
//
func (a *Mat3) DivScalar(s Real) *Mat3 {
	a.Assign(
		a[0][0]/s, a[0][1]/s, a[0][2]/s,
		a[1][0]/s, a[1][1]/s, a[1][2]/s,
		a[2][0]/s, a[2][1]/s, a[2][2]/s,
	)
	return a
}

// Add performs in-place memberwise addition (a += b).
//
// The matrix itself, a, is returned such that you can perform non in-place addition using:
//
//  result := a.Copy().Add()
//
func (a *Mat3) Add(b *Mat3) *Mat3 {
	a.Assign(
		a[0][0]+b[0][0], a[0][1]+b[0][1], a[0][2]+b[0][2],
		a[1][0]+b[1][0], a[1][1]+b[1][1], a[1][2]+b[1][2],
		a[2][0]+b[2][0], a[2][1]+b[2][1], a[2][2]+b[2][2],
	)
	return a
}

// Sub performs in-place memberwise subtraction (a -= b).
//
// The matrix itself, a, is returned such that you can perform non in-place addition using:
//
//  result := a.Copy().Sub()
//
func (a *Mat3) Sub(b *Mat3) *Mat3 {
	a.Assign(
		a[0][0]-b[0][0], a[0][1]-b[0][1], a[0][2]-b[0][2],
		a[1][0]-b[1][0], a[1][1]-b[1][1], a[1][2]-b[1][2],
		a[2][0]-b[2][0], a[2][1]-b[2][1], a[2][2]-b[2][2],
	)
	return a
}

// StoreMul performs matrix multiplication and stores the result inside of out (out = a * b).
//
// The a and b parameters cannot be the same as the out parameter or an panic will occur. ( (create an copy instead).
func (out *Mat3) StoreMul(a, b *Mat3) {
	if a == out || b == out {
		panic("StoreMul(): a == out || b == out; a or b parameter cannot be out!")
	}

	out[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0]
	out[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1]
	out[0][2] = a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2]
	out[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0]
	out[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1]
	out[1][2] = a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2]
	out[2][0] = a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0]
	out[2][1] = a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1]
	out[2][2] = a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2]
}

// Mul is short hand for:
//
//  result := new(Mat3)
//  result.StoreMul(a, b)
//
func (a *Mat3) Mul(b *Mat3) *Mat3 {
	result := new(Mat3)
	result.StoreMul(a, b)
	return result
}

// Sets the values in the matrix such that it will rotate by the given angle in radians
// counterclockwise about the indicated vector.
//
// The axis is expected to be normalized already, via an call to axis.Normalize() for instance.
func (a *Mat3) SetRotation(angle Real, axis *Vec3, cs CoordSys) {
	// Avoid touching their axis and modifying it directly
	axis = axis.Copy()

	// In a left-handed coordinate system, counterclockwise is the
	// other direction.
	if cs.LeftHanded() {
		angle = -angle
	}

	s := Sin(angle)
	c := Cos(angle)
	t := 1.0 - c

	t0 := t * axis.X
	t1 := t * axis.Y
	t2 := t * axis.Z
	s0 := s * axis.X
	s1 := s * axis.Y
	s2 := s * axis.Z

	a[0][0] = t0*axis.X + c
	a[0][1] = t0*axis.Y + s2
	a[0][2] = t0*axis.Z - s1

	a[1][0] = t1*axis.X - s2
	a[1][1] = t1*axis.Y + c
	a[1][2] = t1*axis.Z + s0

	a[2][0] = t2*axis.X + s1
	a[2][1] = t2*axis.Y - s0
	a[2][2] = t2*axis.Z + c
}

// SetScaleShear sets the values in the matrix such that it will apply the given scale and shear
// along their respective axis in the coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (a *Mat3) SetScaleShear(scale, shear *Vec3, cs CoordSys) {
	// We have to match the placement of the shear components in the
	// matrix to the way we extract out the rotation in
	// decompose_matrix().  Therefore, the shear is sensitive to the
	// coordinate system.
	switch cs {
	case CoordSysZUpRight:
		a.Assign(
			scale.X, shear.X*scale.X, 0,
			0, scale.Y, 0,
			shear.Y*scale.Z, shear.Z*scale.Z, scale.Z,
		)

	case CoordSysZUpLeft:
		a.Assign(
			scale.X, shear.X*scale.X, 0,
			0, scale.Y, 0,
			-shear.Y*scale.Z, -shear.Z*scale.Z, scale.Z,
		)

	case CoordSysYUpRight:
		a.Assign(
			scale.X, 0, shear.Y*scale.X,
			shear.X*scale.Y, scale.Y, shear.Z*scale.Y,
			0, 0, scale.Z,
		)

	case CoordSysYUpLeft:
		a.Assign(
			scale.X, 0, -shear.Y*scale.X,
			shear.X*scale.Y, scale.Y, -shear.Z*scale.Y,
			0, 0, scale.Z,
		)

	default:
		panic(fmt.Sprintf("SetScaleShear(): Invalid coordinate system %d", cs))
	}
}

// SetTranslation sets the values in the matrix such that it will apply the given scale along the
// respective axis in the coordinate system.
func (a *Mat3) SetTranslation(trans *Vec2) {
	a.Assign(
		1, 0, 0,
		0, 1, 0,
		trans.X, trans.Y, 1,
	)
}

// StoreTranspose stores the transposed matrix m inside of the out matrix parameter.
//
// The m parameter cannot be the same as the out parameter or an panic will occur (create an copy
// instead).
func (out *Mat3) StoreTranspose(m *Mat3) {
	if m == out {
		panic("StoreTranspose(): m == out; m parameter cannot be out!")
	}

	out[0][0] = m[0][0]
	out[0][1] = m[1][0]
	out[0][2] = m[2][0]

	out[1][0] = m[0][1]
	out[1][1] = m[1][1]
	out[1][2] = m[2][1]

	out[2][0] = m[0][2]
	out[2][1] = m[1][2]
	out[2][2] = m[2][2]
}

// Transpose transposes the matrix m in-place. The matrix m is also returned, such that you can do
// the following to perform an non in-place transpose:
//
//  m.Copy().Transpose()
//
func (m *Mat3) Transpose() *Mat3 {
	m01 := m[0][1]
	m[0][1] = m[1][0]
	m[1][0] = m01

	m02 := m[0][2]
	m[0][2] = m[2][0]
	m[2][0] = m02

	m12 := m[1][2]
	m[1][2] = m[2][1]
	m[2][1] = m12
	return m
}

// StoreInvert stores the inverse matrix m inside of the out matrix parameter.
//
// This is an fully general operation and has no requirements about the type of transform
// represented by this matrix.
//
// The m parameter cannot be the same as the out parameter or an panic will occur (create an copy
// instead).
//
// The return value is true if the matrix was successfully inverted, or false if the matrix was
// singular.
func (out *Mat3) StoreInvert(m *Mat3) bool {
	if m == out {
		panic("StoreInvert(): m == out; m parameter cannot be out!")
	}

	det := m.Determinant()

	// We throw the value out only if it's smaller than our "small"
	// tolerence squared.  This helps reduce overly-sensitive
	// rejections.
	if det.Equals(RealNearZero * RealNearZero) {
		*out = *Mat3Identity
		return false
	}

	det = 1.0 / det

	det2 := func(e00, e01, e10, e11 Real) Real {
		return e00*e11 - e10*e01
	}

	out[0][0] = det * det2(m[1][1], m[1][2], m[2][1], m[2][2])
	out[1][0] = -det * det2(m[1][0], m[1][2], m[2][0], m[2][2])
	out[2][0] = det * det2(m[1][0], m[1][1], m[2][0], m[2][1])

	out[0][1] = -det * det2(m[0][1], m[0][2], m[2][1], m[2][2])
	out[1][1] = det * det2(m[0][0], m[0][2], m[2][0], m[2][2])
	out[2][1] = -det * det2(m[0][0], m[0][1], m[2][0], m[2][1])

	out[0][2] = det * det2(m[0][1], m[0][2], m[1][1], m[1][2])
	out[1][2] = -det * det2(m[0][0], m[0][2], m[1][0], m[1][2])
	out[2][2] = det * det2(m[0][0], m[0][1], m[1][0], m[1][1])

	return true
}

// Invert is short hand for:
//
//  result := new(Mat3)
//  success := result.StoreInvert(a)
//
func (a *Mat3) Invert() (success bool, result *Mat3) {
	result = new(Mat3)
	success = result.StoreInvert(a)
	return
}

// StoreInvertTransposeMat3 simultaneously computes and stores the inverse of the matrix m,
// and then the transpose of that inverse inside of the out matrix parameter.
//
// The m parameter cannot be the same as the out parameter or an panic will occur (create an copy
// instead).
//
// The return value is true if the matrix was successfully inverted, or false if the matrix was
// singular.
func (out *Mat3) StoreInvertTransposeMat3(m *Mat3) bool {
	if m == out {
		panic("StoreInvertTransposeMat3(): m == out; m parameter cannot be out!")
	}

	det := m.Determinant()

	if det.Equals(RealNearZero * RealNearZero) {
		*out = *Mat3Identity
		return false
	}

	det = 1.0 / det

	det2 := func(e00, e01, e10, e11 Real) Real {
		return e00*e11 - e10*e01
	}

	out[0][0] = det * det2(m[1][1], m[1][2], m[2][1], m[2][2])
	out[0][1] = -det * det2(m[1][0], m[1][2], m[2][0], m[2][2])
	out[0][2] = det * det2(m[1][0], m[1][1], m[2][0], m[2][1])

	out[1][0] = -det * det2(m[0][1], m[0][2], m[2][1], m[2][2])
	out[1][1] = det * det2(m[0][0], m[0][2], m[2][0], m[2][2])
	out[1][2] = -det * det2(m[0][0], m[0][1], m[2][0], m[2][1])

	out[2][0] = det * det2(m[0][1], m[0][2], m[1][1], m[1][2])
	out[2][1] = -det * det2(m[0][0], m[0][2], m[1][0], m[1][2])
	out[2][2] = det * det2(m[0][0], m[0][1], m[1][0], m[1][1])
	return true
}

// InvertTransposeMat3 is short hand for:
//
//  result := new(Mat3)
//  success := result.StoreInvertTransposeMat3(a)
//
func (a *Mat3) InvertTransposeMat3(b *Mat3) (success bool, result *Mat3) {
	result = new(Mat3)
	success = result.StoreInvertTransposeMat3(a)
	return
}

// StoreInvertTransposeMat4 simultaneously computes and stores the inverse of the matrix m,
// and then the transpose of that inverse inside of the out matrix parameter.
//
// The return value is true if the matrix was successfully inverted, or false if the matrix was
// singular.
func (out *Mat3) StoreInvertTransposeMat4(m *Mat4) bool {
	det := m.UpperMat3().Determinant()

	if det.Equals(RealNearZero * RealNearZero) {
		*out = *Mat3Identity
		return false
	}

	det = 1.0 / det

	det2 := func(e00, e01, e10, e11 Real) Real {
		return e00*e11 - e10*e01
	}

	out[0][0] = det * det2(m[1][1], m[1][2], m[2][1], m[2][2])
	out[0][1] = -det * det2(m[1][0], m[1][2], m[2][0], m[2][2])
	out[0][2] = det * det2(m[1][0], m[1][1], m[2][0], m[2][1])

	out[1][0] = -det * det2(m[0][1], m[0][2], m[2][1], m[2][2])
	out[1][1] = det * det2(m[0][0], m[0][2], m[2][0], m[2][2])
	out[1][2] = -det * det2(m[0][0], m[0][1], m[2][0], m[2][1])

	out[2][0] = det * det2(m[0][1], m[0][2], m[1][1], m[1][2])
	out[2][1] = -det * det2(m[0][0], m[0][2], m[1][0], m[1][2])
	out[2][2] = det * det2(m[0][0], m[0][1], m[1][0], m[1][1])
	return true
}

// IsNan tells if any components of this matrix are not an number.
func (a *Mat3) IsNaN() bool {
	return IsNaN(a[0][0]) || IsNaN(a[0][1]) || IsNaN(a[0][2]) ||
		IsNaN(a[1][0]) || IsNaN(a[1][1]) || IsNaN(a[1][2]) ||
		IsNaN(a[2][0]) || IsNaN(a[2][1]) || IsNaN(a[2][2])
}

// Matrix3 returns an new *Mat3 given the specified matrix components.
func Matrix3(m00, m01, m02, m10, m11, m12, m20, m21, m22 Real) *Mat3 {
	return &Mat3{
		{m00, m01, m02},
		{m10, m11, m12},
		{m20, m21, m22},
	}
}

var (
	Mat3Identity = Matrix3(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)

	Mat3YToZUp = Matrix3(
		1, 0, 0,
		0, 0, 1,
		0, -1, 0,
	)

	Mat3ZToYUp = Matrix3(
		1, 0, 0,
		0, 0, -1,
		0, 1, 0,
	)

	Mat3FlipY = Matrix3(
		1, 0, 0,
		0, -1, 0,
		0, 0, 1,
	)

	Mat3FlipZ = Matrix3(
		1, 0, 0,
		0, 1, 0,
		0, 0, -1,
	)

	Mat3LZToRY = Mat3FlipY.Mul(Mat3ZToYUp)
	Mat3LYToRZ = Mat3FlipZ.Mul(Mat3YToZUp)
)

// Determinant returns the calculated determinant of the matrix.
func (a *Mat3) Determinant() Real {
	det2 := func(e00, e01, e10, e11 Real) Real {
		return e00*e11 - e10*e01
	}

	return a[0][0]*det2(a[1][1], a[1][2], a[2][1], a[2][2]) - a[0][1]*det2(a[1][0], a[1][2], a[2][0], a[2][2]) + a[0][2]*det2(a[1][0], a[1][1], a[2][0], a[2][1])
}

// MulQuat multiplies this matrix by the specified quaternion and returns the result.
func (a *Mat3) MulQuat(b *Quat) *Mat3 {
	quatMat := new(Mat3)
	b.ExtractToMat3(quatMat)
	return a.Mul(quatMat)
}

// Compose sets the 3x3 matrix using scale, shear, and heading/pitch/roll rotation values for the
// given coordinate system.
func (a *Mat3) Compose(scale, shear, hpr *Vec3, cs CoordSys) {
	a.SetScaleShear(scale, shear, cs)

	if !hpr.Z.Equals(0) {
		r := new(Mat3)
		r.SetRotation(hpr.Z, cs.Forward(), cs)
		a.Mul(r)
	}
	if !hpr.Y.Equals(0) {
		r := new(Mat3)
		r.SetRotation(hpr.Y, cs.Right(), cs)
		a.Mul(r)
	}
	if !hpr.X.Equals(0) {
		r := new(Mat3)
		r.SetRotation(hpr.X, cs.Up(), cs)
		a.Mul(r)
	}
}

// UnwindYUpRotation extracts the rotation about the x, y, and z axes from the given hpr & scale
// matrix. Adjusts the matrix to eliminate the rotation.
//
// This function assumes the matrix is stored in a right-handed Y-up coordinate system.
func (a *Mat3) UnwindYUpRotation(cs CoordSys) (hpr *Vec3) {
	// Extract the axes from the matrix.
	x := a.Row(0)
	y := a.Row(1)
	z := a.Row(2)

	// Project Z into the XZ plane.
	xz := &Vec2{z.X, z.Z}
	xz.Normalize()

	// Compute the rotation about the +Y (up) axis.  This is yaw / heading.
	heading := Atan2(xz.X, xz.Y)

	// Unwind the heading, and continue.
	rotY := new(Mat3)
	rotY.SetRotation(-heading, &Vec3{0, 1, 0}, CoordSysYUpRight)

	x = x.TransformMat3(rotY)
	y = y.TransformMat3(rotY)
	z = z.TransformMat3(rotY)

	// Project the rotated Z into the YZ plane.
	yz := &Vec2{z.Y, z.Z}
	yz.Normalize()

	// Compute the rotation about the +X (right) axis.  This is pitch.
	pitch := -Atan2(yz.X, yz.Y)

	// Unwind the pitch.
	rotX := new(Mat3)
	rotX.SetRotation(-pitch, &Vec3{1, 0, 0}, CoordSysYUpRight)

	x = x.TransformMat3(rotX)
	y = y.TransformMat3(rotX)
	z = z.TransformMat3(rotX)

	// Project the rotated X onto the XY plane.
	xy := &Vec2{x.X, x.Y}
	xy.Normalize()

	// Compute the rotation about the +Z (back) axis.  This is roll.
	roll := -Atan2(xy.Y, xy.X)

	// Unwind the roll from the axes, and continue.
	rotZ := new(Mat3)
	rotZ.SetRotation(roll, &Vec3{0, 0, 1}, CoordSysYUpRight)

	x = x.TransformMat3(rotZ)
	y = y.TransformMat3(rotZ)
	z = z.TransformMat3(rotZ)

	// Reset the matrix to reflect the unwinding.
	a.SetRow(0, x)
	a.SetRow(1, y)
	a.SetRow(2, z)

	// Return the three rotation components.
	hpr = &Vec3{heading, pitch, roll}
	return
}

// UnwindZUpRotation extracts the rotation about the x, y, and z axes from the given hpr & scale
// matrix. Adjusts the matrix to eliminate the rotation.
//
// This function assumes the matrix is stored in a right-handed Z-up coordinate system.
func (a *Mat3) UnwindZUpRotation(cs CoordSys) (hpr *Vec3) {
	// Extract the axes from the matrix.
	x := a.Row(0)
	y := a.Row(1)
	z := a.Row(2)

	// Project Y into the XY plane.
	xy := &Vec2{y.X, y.Y}
	xy.Normalize()

	// Compute the rotation about the +Z (up) axis.  This is yaw / heading.
	heading := -Atan2(xy.X, xy.Y)

	// Unwind the heading, and continue.
	rotZ := new(Mat3)
	rotZ.SetRotation(-heading, &Vec3{0, 0, 1}, CoordSysZUpRight)

	x = x.TransformMat3(rotZ)
	y = y.TransformMat3(rotZ)
	z = z.TransformMat3(rotZ)

	// Project the rotated Y into the YZ plane.
	yz := &Vec2{y.Y, y.Z}
	yz.Normalize()

	// Compute the rotation about the +X (right) axis.  This is pitch.
	pitch := Atan2(yz.Y, yz.X)

	// Unwind the pitch.
	rotX := new(Mat3)
	rotX.SetRotation(-pitch, &Vec3{1, 0, 0}, CoordSysZUpRight)

	x = x.TransformMat3(rotX)
	y = y.TransformMat3(rotX)
	z = z.TransformMat3(rotX)

	// Project X into the XZ plane.
	xz := &Vec2{x.X, x.Z}
	xz.Normalize()

	// Compute the rotation about the -Y (back) axis.  This is roll.
	roll := -Atan2(xz.Y, xz.X)

	// Unwind the roll from the axes, and continue.
	rotY := new(Mat3)
	rotY.SetRotation(-roll, &Vec3{0, 1, 0}, CoordSysZUpRight)

	x = x.TransformMat3(rotY)
	y = y.TransformMat3(rotY)
	z = z.TransformMat3(rotY)

	// Reset the matrix to reflect the unwinding.
	a.SetRow(0, x)
	a.SetRow(1, y)
	a.SetRow(2, z)

	// Return the three rotation components.
	hpr = &Vec3{heading, pitch, roll}
	return
}

// Decompose extracts out the components of the *Mat3 rotation matrix.
func (a *Mat3) Decompose(cs CoordSys) (scale, shear, hpr *Vec3) {
	// Extract the rotation and scale, according to the coordinate
	// system of choice.
	mat := new(Mat3)
	*mat = *a

	switch cs {
	case CoordSysZUpRight:
		hpr = mat.UnwindZUpRotation(cs)

	case CoordSysZUpLeft:
		mat[0][2] = -mat[0][2]
		mat[1][2] = -mat[1][2]
		mat[2][0] = -mat[2][0]
		mat[2][1] = -mat[2][1]
		/*
			FLOATNAME(LMatrix3) lm(mat(0, 0), mat(0, 1), -mat(0, 2),
			mat(1, 0), mat(1, 1), -mat(1, 2),
			-mat(2, 0), -mat(2, 1), mat(2, 2));
		*/
		hpr = mat.UnwindZUpRotation(cs)
		hpr.X = -hpr.X
		hpr.Z = -hpr.Z

	case CoordSysYUpLeft:
		mat[0][2] = -mat[0][2]
		mat[1][2] = -mat[1][2]
		mat[2][0] = -mat[2][0]
		mat[2][1] = -mat[2][1]
		/*
			FLOATNAME(LMatrix3) lm(mat(0, 0), mat(0, 1), -mat(0, 2),
			mat(1, 0), mat(1, 1), -mat(1, 2),
			-mat(2, 0), -mat(2, 1), mat(2, 2));
		*/
		hpr = mat.UnwindYUpRotation(cs)

	default:
		panic(fmt.Sprintf("Decompose(): Unexpected coordinate system %d", cs))
	}

	scale = &Vec3{mat[0][0], mat[1][1], mat[2][2]}

	// Normalize the scale out of the shear components, and return the
	// shear.
	if !scale.X.Equals(0) {
		mat[0][1] /= scale.X
		mat[0][2] /= scale.X
	}
	if !scale.Y.Equals(0) {
		mat[1][0] /= scale.Y
		mat[1][2] /= scale.Y
	}
	if !scale.Z.Equals(0) {
		mat[2][0] /= scale.Z
		mat[2][1] /= scale.Z
	}

	shear = &Vec3{
		mat[0][1] + mat[1][0],
		mat[2][0] + mat[0][2],
		mat[2][1] + mat[1][2],
	}
	return
}
