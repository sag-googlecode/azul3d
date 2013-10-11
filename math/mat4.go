// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// Mat4 represents an 4x4 matrix, indices are in m[row][column] order.
type Mat4 [4][4]Real

// Copy creates an new *Mat4, which is an 1:1 clone of this matrix.
func (a *Mat4) Copy() *Mat4 {
	r := new(Mat4)

	r[0][0] = a[0][0]
	r[0][1] = a[0][1]
	r[0][2] = a[0][2]
	r[0][3] = a[0][3]

	r[1][0] = a[1][0]
	r[1][1] = a[1][1]
	r[1][2] = a[1][2]
	r[1][3] = a[1][3]

	r[2][0] = a[2][0]
	r[2][1] = a[2][1]
	r[2][2] = a[2][2]
	r[2][3] = a[2][3]

	r[3][0] = a[3][0]
	r[3][1] = a[3][1]
	r[3][2] = a[3][2]
	r[3][3] = a[3][3]

	return r
}

// String returns an string representation of this matrix.
func (a *Mat4) String() string {
	return fmt.Sprintf("Mat4(\n    %v, %v, %v, %v\n    %v, %v, %v, %v\n    %v, %v, %v, %v\n    %v, %v, %v, %v\n)", a[0][0], a[0][1], a[0][2], a[0][3], a[1][0], a[1][1], a[1][2], a[1][3], a[2][0], a[2][1], a[2][2], a[2][3], a[3][0], a[3][1], a[3][2], a[3][3])
}

// Assign assigns the specified values to this matrix
func (a *Mat4) Assign(m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33 Real) {
	a[0][0] = m00
	a[0][1] = m01
	a[0][2] = m02
	a[0][3] = m03

	a[1][0] = m10
	a[1][1] = m11
	a[1][2] = m12
	a[1][3] = m13

	a[2][0] = m20
	a[2][1] = m21
	a[2][2] = m22
	a[2][3] = m23

	a[3][0] = m30
	a[3][1] = m31
	a[3][2] = m32
	a[3][3] = m33
}

// Fill fills this matrix with the value specified. Useful sometimes for initializing to zero, etc.
func (a *Mat4) Fill(v Real) {
	a.Assign(
		v, v, v, v,
		v, v, v, v,
		v, v, v, v,
		v, v, v, v,
	)
}

// SetRow sets the values in the specified matrix row to the values in the specified
// fourth-component vector.
//
// The row parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) SetRow(row int, values *Vec4) {
	a[row][0] = values.X
	a[row][1] = values.Y
	a[row][2] = values.Z
	a[row][3] = values.W
}

// Row returns the values in the specified matrix row as an three-component vector.
//
// The row parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) Row(row int) *Vec4 {
	return &Vec4{a[row][0], a[row][1], a[row][2], a[row][3]}
}

// SetRowVec3 sets the values in the specified matrix row to the values in the specified
// three-component vector, leaving the fourth element in the row untouched.
//
// The row parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) SetRowVec3(row int, values *Vec3) {
	a[row][0] = values.X
	a[row][1] = values.Y
	a[row][2] = values.Z
}

// RowVec3 returns the values in the specified matrix row as an three-component vector, the fourth
// element of the row is ignored.
//
// The row parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) RowVec3(row int) *Vec3 {
	return &Vec3{a[row][0], a[row][1], a[row][2]}
}

// SetCol sets the values in the specified matrix column to the values in the specified
// four-component vector.
//
// The column parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) SetCol(column int, values *Vec4) {
	a[0][column] = values.X
	a[1][column] = values.Y
	a[2][column] = values.Z
	a[3][column] = values.W
}

// Col returns the values in the specified matrix column as an four-component vector.
//
// The column parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) Col(column int) *Vec4 {
	return &Vec4{a[0][column], a[1][column], a[2][column], a[3][column]}
}

// SetColVec3 sets the values in the specified matrix column to the values in the specified
// three-component vector, leaving the fourth element in the column untouched.
//
// The column parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) SetColVec3(column int, values *Vec3) {
	a[0][column] = values.X
	a[1][column] = values.Y
	a[2][column] = values.Z
}

// ColVec3 returns the values in the specified matrix column as an three-component vector, ignoring
// the fourth element of the column.
//
// The column parameter must be 0, 1, 2, or 3 or else an panic will occur.
func (a *Mat4) ColVec3(column int) *Vec3 {
	return &Vec3{a[0][column], a[1][column], a[2][column]}
}

// EqualsTolerence tells weather a is memberwise equal to b within the given tolerence for Real
// equality.
//
// Equality against nil is always false.
func (a *Mat4) EqualsTolerence(b *Mat4, tolerence Real) bool {
	if b == nil {
		return false
	}
	return a[0][0].EqualsTolerence(b[0][0], tolerence) &&
		a[0][1].EqualsTolerence(b[0][1], tolerence) &&
		a[0][2].EqualsTolerence(b[0][2], tolerence) &&
		a[0][3].EqualsTolerence(b[0][3], tolerence) &&

		a[1][0].EqualsTolerence(b[1][0], tolerence) &&
		a[1][1].EqualsTolerence(b[1][1], tolerence) &&
		a[1][2].EqualsTolerence(b[1][2], tolerence) &&
		a[1][3].EqualsTolerence(b[1][3], tolerence) &&

		a[2][0].EqualsTolerence(b[2][0], tolerence) &&
		a[2][1].EqualsTolerence(b[2][1], tolerence) &&
		a[2][2].EqualsTolerence(b[2][2], tolerence) &&
		a[2][3].EqualsTolerence(b[2][3], tolerence) &&

		a[3][0].EqualsTolerence(b[3][0], tolerence) &&
		a[3][1].EqualsTolerence(b[3][1], tolerence) &&
		a[3][2].EqualsTolerence(b[3][2], tolerence) &&
		a[3][3].EqualsTolerence(b[3][3], tolerence)
}

// Equals tells weather a is memberwise equal to b within the default tolerence for Real equality.
//
// Equality against nil is always false.
func (a *Mat4) Equals(b *Mat4) bool {
	if b == nil {
		return false
	}
	return a.EqualsTolerence(b, RealNearZero)
}

// CompareTolerence sorts the two matrices lexicographically, componentwise.
//
// Returns -1 if matrix a sorts before matrix b.
//
// Returns 1 if matrix a sorts after b.
//
// Returns 0 if matrix a and b are equal (within the given tolerence).
func (a *Mat4) CompareTolerence(b *Mat4, tolerence Real) int {
	// We compare values in reverse order, since the last row of the matrix is most likely to be
	// different between multiple matrices.
	for r := 3; r >= 0; r-- {
		for c := 0; c < 4; c++ {
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
func (a *Mat4) Compare(b *Mat4) int {
	return a.CompareTolerence(b, RealNearZero)
}

// LessThan returns a < b
//
// Also see the Compare() and CompareTolerence() functions.
func (a *Mat4) LessThan(b *Mat4) bool {
	return a.Compare(b) < 0
}

// AddScalar performs in-place memberwise addition (a += s).
//
// The matrix itself, a, is returned such that you can perform non in-place addition using:
//
//  result := a.Copy().AddScalar()
//
func (a *Mat4) AddScalar(s Real) *Mat4 {
	a.Assign(
		a[0][0]+s, a[0][1]+s, a[0][2]+s, a[0][3]+s,
		a[1][0]+s, a[1][1]+s, a[1][2]+s, a[1][3]+s,
		a[2][0]+s, a[2][1]+s, a[2][2]+s, a[2][3]+s,
		a[3][0]+s, a[3][1]+s, a[3][2]+s, a[3][3]+s,
	)
	return a
}

// SubScalar performs in-place memberwise subtraction (a -= s).
//
// The matrix itself, a, is returned such that you can perform non in-place subtraction using:
//
//  result := a.Copy().SubScalar()
//
func (a *Mat4) SubScalar(s Real) *Mat4 {
	a.Assign(
		a[0][0]-s, a[0][1]-s, a[0][2]-s, a[0][3]-s,
		a[1][0]-s, a[1][1]-s, a[1][2]-s, a[1][3]-s,
		a[2][0]-s, a[2][1]-s, a[2][2]-s, a[2][3]-s,
		a[3][0]-s, a[3][1]-s, a[3][2]-s, a[3][3]-s,
	)
	return a
}

// MulScalar performs in-place memberwise multiplication (a *= s).
//
// The matrix itself, a, is returned such that you can perform non in-place multiplication using:
//
//  result := a.Copy().MulScalar()
//
func (a *Mat4) MulScalar(s Real) *Mat4 {
	a.Assign(
		a[0][0]*s, a[0][1]*s, a[0][2]*s, a[0][3]*s,
		a[1][0]*s, a[1][1]*s, a[1][2]*s, a[1][3]*s,
		a[2][0]*s, a[2][1]*s, a[2][2]*s, a[2][3]*s,
		a[3][0]*s, a[3][1]*s, a[3][2]*s, a[3][3]*s,
	)
	return a
}

// DivScalar performs in-place memberwise division (a /= s).
//
// The matrix itself, a, is returned such that you can perform non in-place division using:
//
//  result := a.Copy().DivScalar()
//
func (a *Mat4) DivScalar(s Real) *Mat4 {
	a.Assign(
		a[0][0]/s, a[0][1]/s, a[0][2]/s, a[0][3]/s,
		a[1][0]/s, a[1][1]/s, a[1][2]/s, a[1][3]/s,
		a[2][0]/s, a[2][1]/s, a[2][2]/s, a[2][3]/s,
		a[3][0]/s, a[3][1]/s, a[3][2]/s, a[3][3]/s,
	)
	return a
}

// Add performs in-place memberwise addition (a += b).
//
// The matrix itself, a, is returned such that you can perform non in-place addition using:
//
//  result := a.Copy().Add()
//
func (a *Mat4) Add(b *Mat4) *Mat4 {
	a.Assign(
		a[0][0]+b[0][0], a[0][1]+b[0][1], a[0][2]+b[0][2], a[0][3]+b[0][3],
		a[1][0]+b[1][0], a[1][1]+b[1][1], a[1][2]+b[1][2], a[1][3]+b[1][3],
		a[2][0]+b[2][0], a[2][1]+b[2][1], a[2][2]+b[2][2], a[2][3]+b[2][3],
		a[3][0]+b[3][0], a[3][1]+b[3][1], a[3][2]+b[3][2], a[3][3]+b[3][3],
	)
	return a
}

// Sub performs in-place memberwise subtraction (a -= b).
//
// The matrix itself, a, is returned such that you can perform non in-place addition using:
//
//  result := a.Copy().Sub()
//
func (a *Mat4) Sub(b *Mat4) *Mat4 {
	a.Assign(
		a[0][0]-b[0][0], a[0][1]-b[0][1], a[0][2]-b[0][2], a[0][3]-b[0][3],
		a[1][0]-b[1][0], a[1][1]-b[1][1], a[1][2]-b[1][2], a[1][3]-b[1][3],
		a[2][0]-b[2][0], a[2][1]-b[2][1], a[2][2]-b[2][2], a[2][3]-b[2][3],
		a[3][0]-b[3][0], a[3][1]-b[3][1], a[3][2]-b[3][2], a[3][3]-b[3][3],
	)
	return a
}

// StoreMul performs matrix multiplication and stores the result inside of out (out = a * b).
//
// The a and b parameters cannot be the same as the out parameter or an panic will occur. (create
// an copy instead).
func (out *Mat4) StoreMul(a, b *Mat4) {
	if a == out || b == out {
		panic("StoreMul(): a == out || b == out; a or b parameter cannot be out!")
	}
	out[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0]
	out[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1]
	out[0][2] = a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*b[3][2]
	out[0][3] = a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3] + a[0][3]*b[3][3]

	out[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0]
	out[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1]
	out[1][2] = a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*b[3][2]
	out[1][3] = a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3] + a[1][3]*b[3][3]

	out[2][0] = a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0]
	out[2][1] = a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1]
	out[2][2] = a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*b[3][2]
	out[2][3] = a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3] + a[2][3]*b[3][3]

	out[3][0] = a[3][0]*b[0][0] + a[3][1]*b[1][0] + a[3][2]*b[2][0] + a[3][3]*b[3][0]
	out[3][1] = a[3][0]*b[0][1] + a[3][1]*b[1][1] + a[3][2]*b[2][1] + a[3][3]*b[3][1]
	out[3][2] = a[3][0]*b[0][2] + a[3][1]*b[1][2] + a[3][2]*b[2][2] + a[3][3]*b[3][2]
	out[3][3] = a[3][0]*b[0][3] + a[3][1]*b[1][3] + a[3][2]*b[2][3] + a[3][3]*b[3][3]
}

// Mul is short hand for:
//
//  result := new(Mat4)
//  result.StoreMul(a, b)
//
func (a *Mat4) Mul(b *Mat4) *Mat4 {
	result := new(Mat4)
	result.StoreMul(a, b)
	return result
}

// SetRotation sets the values in the matrix such that it will rotate by the given angle in radians
// counterclockwise about the indicated vector.
//
// The axis is expected to be normalized already, via an call to axis.Normalize() for instance.
func (a *Mat4) SetRotation(angle Real, axis *Vec3, cs CoordSys) {
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

	a[0][3] = 0
	a[1][3] = 0
	a[2][3] = 0

	a[3][0] = 0
	a[3][1] = 0
	a[3][2] = 0
	a[3][3] = 1
}

// SetScaleShear sets the values in the matrix such that it will apply the given scale and shear
// along their respective axis in the coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (a *Mat4) SetScaleShear(scale, shear *Vec3, cs CoordSys) {
	m3 := new(Mat3)
	m3.SetScaleShear(scale, shear, cs)

	a.SetUpperMat3(m3)

	a[0][3] = 0
	a[1][3] = 0
	a[2][3] = 0
	a[3][3] = 1

	a[3][0] = 0.0
	a[3][1] = 0.0
	a[3][2] = 0.0
}

// SetScale sets the values in the matrix such that it will apply the given scale along the
// respective axis in the coordinate system.
func (a *Mat4) SetScale(scale *Vec3) {
	a.Assign(
		scale.X, 0, 0, 0,
		0, scale.Y, 0, 0,
		0, 0, scale.Z, 0,
		0, 0, 0, 1,
	)
}

// SetTranslation sets the values in the matrix such that it will apply the given translation along
// the respective axis in the coordinate system.
func (a *Mat4) SetTranslation(trans *Vec3) {
	a.Assign(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		trans.X, trans.Y, trans.Z, 1,
	)
}

// Translation returns the translation of the given matrix.
func (a *Mat4) Translation() *Vec3 {
	return Vector3(a[3][0], a[3][1], a[3][2])
}

// SetFrustum sets this matrix to be an frustum matrix, using the left, right, bottom, top, near
// and far bounds of the frustum.
func (a *Mat4) SetFrustum(left, right, bottom, top, near, far Real) {
	normalWidth := 1.0 / (right - left)
	normalHeight := 1.0 / (top - bottom)
	normalDepth := 1.0 / (near - far)

	a.Assign(
		(near*2.0)*normalWidth,
		0,
		0,
		0,

		0,
		(near*2.0)*normalHeight,
		0,
		0,

		(right+left)*normalWidth,
		(top+bottom)*normalHeight,
		(far+near)*normalDepth,
		-1,

		0,
		0,
		(far*near*2.0)*normalDepth,
		0,
	)
}

// SetPerspective turns this matrix into an frustum matrix representing the specified field of view,
// aspect ratio, and zNear/zFar values.
func (a *Mat4) SetPerspective(fovY, aspect, zNear, zFar Real) {
	fH := Tan(fovY/360*Pi) * zNear
	fW := fH * aspect
	a.SetFrustum(-fW, fW, -fH, fH, zNear, zFar)
}

// SetOrtho assigns this matrix to be an orthographic projection matrix using the
// specified bounds for the frustum.
//
// See: http://en.wikipedia.org/wiki/Orthographic_projection_(geometry)
func (a *Mat4) SetOrtho(left, right, bottom, top, near, far Real) {
	a.Assign(
		2.0/(right-left),
		0,
		0,
		0,

		0,
		2.0/(top-bottom),
		0,
		0,

		0,
		0,
		-2.0/(far-near),
		0,

		-(right+left)/(right-left),
		-(top+bottom)/(top-bottom),
		-(far+near)/(far-near),
		1,
	)
}

/// SetUnOrtho assigns this matrix to be an orthographic unprojection matrix using the
// specified bounds for the frustum.
//
// See: http://en.wikipedia.org/wiki/Orthographic_projection_(geometry)
func (a *Mat4) SetUnOrtho(left, right, bottom, top, near, far Real) {
	a.Assign(
		(right-left)/2.0,
		0,
		0,
		0,

		0,
		(top-bottom)/2.0,
		0,
		0,

		0,
		0,
		(far-near)/-2.0,
		0,

		(left+right)/2,
		(top+bottom)/2,
		(far+near)/-2,
		1,
	)
}

// StoreTranspose stores the transposed matrix m inside of the out matrix parameter.
//
// The m parameter cannot be the same as the out parameter or an panic will occur (create an copy
// instead).
func (out *Mat4) StoreTranspose(m *Mat4) {
	if m == out {
		panic("StoreTranspose(): m == out; m parameter cannot be out!")
	}

	out[0][0] = m[0][0]
	out[0][1] = m[1][0]
	out[0][2] = m[2][0]
	out[0][3] = m[3][0]

	out[1][0] = m[0][1]
	out[1][1] = m[1][1]
	out[1][2] = m[2][1]
	out[1][3] = m[3][1]

	out[2][0] = m[0][2]
	out[2][1] = m[1][2]
	out[2][2] = m[2][2]
	out[2][3] = m[3][2]

	out[3][0] = m[0][3]
	out[3][1] = m[1][3]
	out[3][2] = m[2][3]
	out[3][3] = m[3][3]
}

// Transpose transposes the matrix m in-place. The matrix m is also returned, such that you can do
// the following to perform an non in-place transpose:
//
//  m.Copy().Transpose()
//
func (m *Mat4) Transpose() *Mat4 {
	m01 := m[0][1]
	m[0][1] = m[1][0]
	m[1][0] = m01

	m02 := m[0][2]
	m[0][2] = m[2][0]
	m[2][0] = m02

	m03 := m[0][3]
	m[0][3] = m[3][0]
	m[3][0] = m03

	m12 := m[1][2]
	m[1][2] = m[2][1]
	m[2][1] = m12

	m13 := m[1][3]
	m[1][3] = m[3][1]
	m[3][1] = m13

	m23 := m[2][3]
	m[2][3] = m[3][2]
	m[3][2] = m23
	return m
}

// StoreInvertAffine performs an invert of the matrix m, storing the result in the out matrix. The
// calculation is only correct if the matrix m represents an affine transformation.
//
// The m parameter cannot be the same as the out parameter or an panic will occur (create an copy
// instead).
//
// The return value is true if the matrix was successfully inverted, or false if the matrix was
// singular.
func (out *Mat4) StoreInvertAffine(m *Mat4) bool {
	if m == out {
		panic("StoreInvertAffine(): m == out; m parameter cannot be out!")
	}

	rot := new(Mat3)
	if !rot.StoreInvert(m.UpperMat3()) {
		return false
	}

	out.SetUpperMat3(rot)

	out[0][3] = 0
	out[1][3] = 0
	out[2][3] = 0
	out[3][3] = 1

	out[3][0] = -(m[3][0]*out[0][0] + m[3][1]*out[1][0] + m[3][2]*out[2][0])
	out[3][1] = -(m[3][0]*out[0][1] + m[3][1]*out[1][1] + m[3][2]*out[2][1])
	out[3][2] = -(m[3][0]*out[0][2] + m[3][1]*out[1][2] + m[3][2]*out[2][2])
	return true
}

func (m *Mat4) decomposeMat(index *[4]int) bool {
	var vv [4]Real

	for i := 0; i < 4; i++ {
		big := Real(0.0)
		for j := 0; j < 4; j++ {
			temp := Abs(m[i][j])
			if temp > big {
				big = temp
			}
		}

		// We throw the value out only if it's smaller than our "small"
		// threshold squared.  This helps reduce overly-sensitive
		// rejections.
		if big.Equals(RealNearZero * RealNearZero) {
			return false
		}
		vv[i] = 1.0 / big
	}

	for j := 0; j < 4; j++ {
		for i := 0; i < j; i++ {
			sum := m[i][j]
			for k := 0; k < i; k++ {
				sum -= m[i][k] * m[k][j]
			}
			m[i][j] = sum
		}

		big := Real(0)
		imax := -1

		for i := j; i < 4; i++ {
			sum := m[i][j]
			for k := 0; k < j; k++ {
				sum -= m[i][k] * m[k][j]
			}

			m[i][j] = sum

			dum := vv[i] * Abs(sum)

			if dum >= big {
				big = dum
				imax = i
			}
		}
		if j != imax {
			for k := 0; k < 4; k++ {
				dum := m[imax][k]
				m[imax][k] = m[j][k]
				m[j][k] = dum
			}
			vv[imax] = vv[j]
		}
		index[j] = imax

		if m[j][j].Equals(0) {
			m[j][j] = RealNearZero
		}

		if j != 4-1 {
			dum := 1.0 / m[j][j]

			for i := j + 1; i < 4; i++ {
				m[i][j] *= dum
			}
		}
	}
	return true
}

func (m *Mat4) backSubMat(index *[4]int, inv *Mat4, row int) bool {
	var i int

	ii := -1

	for i = 0; i < 4; i++ {
		ip := index[i]
		sum := inv[row][ip]
		inv[row][ip] = inv[row][i]
		if ii >= 0 {
			for j := ii; j <= i-1; j++ {
				sum -= m[i][j] * inv[row][j]
			}
		} else if sum != 0 {
			ii = i
		}

		inv[row][i] = sum
	}

	for i = 4 - 1; i >= 0; i-- {
		sum := inv[row][i]
		for j := i + 1; j < 4; j++ {
			sum -= m[i][j] * inv[row][j]
		}
		inv[row][i] = sum / m[i][i]
	}

	return true
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
func (out *Mat4) StoreInvert(m *Mat4) bool {
	if m == out {
		panic("StoreInvert(): m == out; m parameter cannot be out!")
	}

	if m[0][3].Equals(0) && m[1][3].Equals(0) && m[2][3].Equals(0) && m[3][3].Equals(0) {
		return out.StoreInvertAffine(m)
	}

	*out = *m

	index := [4]int{}

	// Other order?
	if !out.decomposeMat(&index) {
		*out = *Mat4Identity
		return false
	}

	inv := new(Mat4)
	*inv = *Mat4Identity

	for row := 0; row < 4; row++ {
		out.backSubMat(&index, inv, row)
	}

	out.StoreTranspose(inv)
	return true
}

// Invert is short hand for:
//
//  result := new(Mat4)
//  success := result.StoreInvert(a)
//
func (a *Mat4) Invert() (result *Mat4, success bool) {
	result = new(Mat4)
	success = result.StoreInvert(a)
	return
}

// IsNan tells if any components of this matrix are not an number.
func (a *Mat4) IsNaN() bool {
	return IsNaN(a[0][0]) || IsNaN(a[0][1]) || IsNaN(a[0][2]) || IsNaN(a[0][3]) ||
		IsNaN(a[1][0]) || IsNaN(a[1][1]) || IsNaN(a[1][2]) || IsNaN(a[1][3]) ||
		IsNaN(a[2][0]) || IsNaN(a[2][1]) || IsNaN(a[2][2]) || IsNaN(a[2][3]) ||
		IsNaN(a[3][0]) || IsNaN(a[3][1]) || IsNaN(a[3][2]) || IsNaN(a[3][3])
}

// Matrix4 returns an new *Mat4 given the specified matrix components.
func Matrix4(m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33 Real) *Mat4 {
	return &Mat4{
		{m00, m01, m02, m03},
		{m10, m11, m12, m13},
		{m20, m21, m22, m23},
		{m30, m31, m32, m33},
	}
}

var (
	Mat4Identity = Matrix4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	Mat4Zeros = Matrix4(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)

	Mat4Ones = Matrix4(
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	)

	Mat4YToZUp = Matrix4(
		1, 0, 0, 0,
		0, 0, 1, 0,
		0, -1, 0, 0,
		0, 0, 0, 1,
	)

	Mat4ZToYUp = Matrix4(
		1, 0, 0, 0,
		0, 0, -1, 0,
		0, 1, 0, 0,
		0, 0, 0, 1,
	)

	Mat4FlipY = Matrix4(
		1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	Mat4FlipZ = Matrix4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	)

	Mat4LZToRY = Mat4FlipY.Mul(Mat4ZToYUp)
	Mat4LYToRZ = Mat4FlipZ.Mul(Mat4YToZUp)
)

// SetUpperMat3 sets the upper-left 3x3 matrix to the specified one.
func (a *Mat4) SetUpperMat3(b *Mat3) {
	a[0][0] = b[0][0]
	a[0][1] = b[0][1]
	a[0][2] = b[0][2]

	a[1][0] = b[1][0]
	a[1][1] = b[1][1]
	a[1][2] = b[1][2]

	a[2][0] = b[2][0]
	a[2][1] = b[2][1]
	a[2][2] = b[2][2]
}

// UpperMat3 returns the upper-left 3x3 matrix as an new *Mat3.
func (a *Mat4) UpperMat3() *Mat3 {
	return Matrix3(
		a[0][0], a[0][1], a[0][2],
		a[1][0], a[1][1], a[1][2],
		a[2][0], a[2][1], a[2][2],
	)
}

func (a *Mat4) MulQuat(b *Quat) *Mat4 {
	quatMat := new(Mat4)
	b.ExtractToMat4(quatMat)

	// Preserve the homogeneous coords and the translate
	row3 := a.Row(3)
	col3 := a.Col(3)

	quatMat = a.Mul(quatMat)
	quatMat.SetRow(3, row3)
	quatMat.SetCol(3, col3)
	return quatMat
}
