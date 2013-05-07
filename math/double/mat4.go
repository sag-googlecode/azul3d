// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package double

import (
	"bytes"
	"fmt"
	"math"
)

func eq(a, b float64) bool {
	diff := math.Abs(a - b)
	if diff > math.Nextafter(0, math.MaxFloat64) {
		return false
	}
	return true
}

// Mat4 represents an 4x4 matrix
type Mat4 [4][4]float64

// Swap swaps the two specified elements in this matrix
//
// Equivilent to:
//
//  r1c1 := m[r1][c1]
//  m[r1][c1] = m[r2][c2]
//  m[r2][c2] = r1c1
//
func (m *Mat4) Swap(r1, c1, r2, c2 int) {
	r1c1 := m[r1][c1]
	m[r1][c1] = m[r2][c2]
	m[r2][c2] = r1c1
}

// Equals tells weather this matrix is equal element wise within the default tolerance.
func (m *Mat4) Equals(m2 *Mat4) bool {
	return eq(m[0][0], m2[0][0]) && eq(m[0][1], m2[0][1]) && eq(m[0][2], m[0][2]) && eq(m[0][3], m[0][3]) && eq(m[1][0], m2[1][0]) && eq(m[1][1], m2[1][1]) && eq(m[1][2], m[1][2]) && eq(m[1][3], m[1][3]) && eq(m[2][0], m2[2][0]) && eq(m[2][1], m2[2][1]) && eq(m[2][2], m[2][2]) && eq(m[2][3], m[2][3]) && eq(m[3][0], m2[3][0]) && eq(m[3][1], m2[3][1]) && eq(m[3][2], m[3][2]) && eq(m[3][3], m[3][3])
}

// Multiply multiplies this matrix by the other matrix, m2.
//
// (the result is stored in this matrix)
//
//  m *= m2
//
func (m *Mat4) Multiply(m2 *Mat4) {
	m00 := m[0][0]
	m01 := m[0][1]
	m02 := m[0][2]
	m03 := m[0][3]

	m10 := m[1][0]
	m11 := m[1][1]
	m12 := m[1][2]
	m13 := m[1][3]

	m20 := m[2][0]
	m21 := m[2][1]
	m22 := m[2][2]
	m23 := m[2][3]

	m30 := m[3][0]
	m31 := m[3][1]
	m32 := m[3][2]
	m33 := m[3][3]

	m[0][0] = m00*m2[0][0] + m01*m2[1][0] + m02*m2[2][0] + m03*m2[3][0]
	m[0][1] = m00*m2[0][1] + m01*m2[1][1] + m02*m2[2][1] + m03*m2[3][1]
	m[0][2] = m00*m2[0][2] + m01*m2[1][2] + m02*m2[2][2] + m03*m2[3][2]
	m[0][3] = m00*m2[0][3] + m01*m2[1][3] + m02*m2[2][3] + m03*m2[3][3]

	m[1][0] = m10*m2[0][0] + m11*m2[1][0] + m12*m2[2][0] + m13*m2[3][0]
	m[1][1] = m10*m2[0][1] + m11*m2[1][1] + m12*m2[2][1] + m13*m2[3][1]
	m[1][2] = m10*m2[0][2] + m11*m2[1][2] + m12*m2[2][2] + m13*m2[3][2]
	m[1][3] = m10*m2[0][3] + m11*m2[1][3] + m12*m2[2][3] + m13*m2[3][3]

	m[2][0] = m20*m2[0][0] + m21*m2[1][0] + m22*m2[2][0] + m23*m2[3][0]
	m[2][1] = m20*m2[0][1] + m21*m2[1][1] + m22*m2[2][1] + m23*m2[3][1]
	m[2][2] = m20*m2[0][2] + m21*m2[1][2] + m22*m2[2][2] + m23*m2[3][2]
	m[2][3] = m20*m2[0][3] + m21*m2[1][3] + m22*m2[2][3] + m23*m2[3][3]

	m[3][0] = m30*m2[0][0] + m31*m2[1][0] + m32*m2[2][0] + m33*m2[3][0]
	m[3][1] = m30*m2[0][1] + m31*m2[1][1] + m32*m2[2][1] + m33*m2[3][1]
	m[3][2] = m30*m2[0][2] + m31*m2[1][2] + m32*m2[2][2] + m33*m2[3][2]
	m[3][3] = m30*m2[0][3] + m31*m2[1][3] + m32*m2[2][3] + m33*m2[3][3]
}

func (m *Mat4) ScalarMultiply(s float64) {
	for rowI, row := range m {
		for colI, col := range row {
			m[rowI][colI] = col * s
		}
	}
}

// Translation sets the matrix to be an translation matrix which translates with the given vector.
//
// (visualized)
//
//  [1, 0, 0, v.X]
//  [0, 1, 0, v.Y]
//  [0, 0, 1, v.Z]
//  [0, 0, 0, 1]
//
/*
func (m *Mat4) Translation(v *Vec3) {
	m[0][0] = 1; m[0][1] = 0; m[0][2] = 0; m[0][3] = v.X
	m[1][0] = 0; m[1][1] = 1; m[1][2] = 0; m[1][3] = v.Y
	m[2][0] = 0; m[2][1] = 0; m[2][2] = 1; m[2][3] = v.Z
	m[3][0] = 0; m[3][1] = 0; m[3][2] = 0; m[3][3] = 1
}

// Translate translates the current matrix by the given vector.
//
// (visualized)
//
//  [1, 0, 0, x * v.X]
//  [0, 1, 0, y * v.Y]
//  [0, 0, 1, z * v.Z]
//  [0, 0, 0,    1   ]
//
func (m *Mat4) Translate(v *Vec3) {
	//m[0][3] = m[0][0] * v.X + m[0][3] * v.Y + m[1][

	//out[3][0] = a[0] * x + a[4] * y + a[8] * z + a[12];
	//out[3][1] = a[1] * x + a[5] * y + a[9] * z + a[13];
	//out[3][2] = a[2] * x + a[6] * y + a[10] * z + a[14];
	//out[3][3] = a[3] * x + a[7] * y + a[11] * z + a[15];
}
*/

// Transpose transposes the elements of this matrix, such that the matrix:
//
//  [00, 01, 02, 03]
//  [10, 11, 12, 13]
//  [20, 11, 12, 23]
//  [30, 31, 32, 33]
//
// Would become:
//
//  [00, 10, 20, 30]
//  [01, 11, 21, 31]
//  [02, 12, 22, 32]
//  [03, 13, 23, 33]
//
func (m *Mat4) Transpose() {
	// These stay the same:
	//
	// [--, 01, 02, 03]
	// [10, --, 12, 13]
	// [20, 21, --, 23]
	// [30, 31, 32, --]

	//m.Swap(0, 1, 1, 0)
	m01 := m[0][1]
	m[0][1] = m[1][0]
	m[1][0] = m01

	//m.Swap(0, 2, 2, 0)
	m02 := m[0][2]
	m[0][2] = m[2][0]
	m[2][0] = m02

	//m.Swap(0, 3, 3, 0)
	m03 := m[0][3]
	m[0][3] = m[3][0]
	m[3][0] = m03

	//m.Swap(1, 2, 2, 1)
	m12 := m[1][2]
	m[1][2] = m[2][1]
	m[2][1] = m12

	//m.Swap(1, 3, 3, 1)
	m13 := m[1][3]
	m[1][3] = m[3][1]
	m[3][1] = m13

	//m.Swap(2, 3, 3, 2)
	m23 := m[2][3]
	m[2][3] = m[3][2]
	m[3][2] = m23
}

// Determinant calculates and returns the determinant of this matrix
func (m *Mat4) Determinant() float64 {
	b00 := m[0][0]*m[1][1] - m[0][1]*m[1][0]
	b01 := m[0][0]*m[1][2] - m[0][2]*m[1][0]
	b02 := m[0][0]*m[1][3] - m[0][3]*m[1][0]
	b03 := m[0][1]*m[1][2] - m[0][2]*m[1][1]
	b04 := m[0][1]*m[1][3] - m[0][3]*m[1][1]
	b05 := m[0][2]*m[1][3] - m[0][3]*m[1][2]
	b06 := m[2][0]*m[3][1] - m[2][1]*m[3][0]
	b07 := m[2][0]*m[3][2] - m[2][2]*m[3][0]
	b08 := m[2][0]*m[3][3] - m[2][3]*m[3][0]
	b09 := m[2][1]*m[3][2] - m[2][2]*m[3][1]
	b10 := m[2][1]*m[3][3] - m[2][3]*m[3][1]
	b11 := m[2][2]*m[3][3] - m[2][3]*m[3][2]

	// Calculate the determinant
	return b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06
}

// Invert inverts this matrix, and returns true if the matrix was inverted, or returns false if the
// determinant is zero (and the matrix was not inverted).
func (m *Mat4) Invert() bool {
	m00 := m[0][0]
	m01 := m[0][1]
	m02 := m[0][2]
	m03 := m[0][3]

	m10 := m[1][0]
	m11 := m[1][1]
	m12 := m[1][2]
	m13 := m[1][3]

	m20 := m[2][0]
	m21 := m[2][1]
	m22 := m[2][2]
	m23 := m[2][3]

	m30 := m[3][0]
	m31 := m[3][1]
	m32 := m[3][2]
	m33 := m[3][3]

	b00 := m00*m11 - m01*m10
	b01 := m00*m12 - m02*m10
	b02 := m00*m13 - m03*m10
	b03 := m01*m12 - m02*m11
	b04 := m01*m13 - m03*m11
	b05 := m02*m13 - m03*m12
	b06 := m20*m31 - m21*m30
	b07 := m20*m32 - m22*m30
	b08 := m20*m33 - m23*m30
	b09 := m21*m32 - m22*m31
	b10 := m21*m33 - m23*m31
	b11 := m22*m33 - m23*m32

	// Calculate the determinant
	det := b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06

	if det == 0 {
		return false
	}
	det = 1.0 / det

	m[0][0] = (m11*b11 - m12*b10 + m13*b09) * det
	m[0][1] = (m02*b10 - m01*b11 - m03*b09) * det
	m[0][2] = (m31*b05 - m32*b04 + m33*b03) * det
	m[0][3] = (m22*b04 - m21*b05 - m23*b03) * det
	m[1][0] = (m12*b08 - m10*b11 - m13*b07) * det
	m[1][1] = (m00*b11 - m02*b08 + m03*b07) * det
	m[1][2] = (m32*b02 - m30*b05 - m33*b01) * det
	m[1][3] = (m20*b05 - m22*b02 + m23*b01) * det
	m[2][0] = (m10*b10 - m11*b08 + m13*b06) * det
	m[2][1] = (m01*b08 - m00*b10 - m03*b06) * det
	m[2][2] = (m30*b04 - m31*b02 + m33*b00) * det
	m[2][3] = (m21*b02 - m20*b04 - m23*b00) * det
	m[3][0] = (m11*b07 - m10*b09 - m12*b06) * det
	m[3][1] = (m00*b09 - m01*b07 + m02*b06) * det
	m[3][2] = (m31*b01 - m30*b03 - m32*b00) * det
	m[3][3] = (m20*b03 - m21*b01 + m22*b00) * det
	return true
}

// Fill fills all elements in the matrix using the number, n
func (m *Mat4) Fill(n float64) {
	for rowI, row := range m {
		for col, _ := range row {
			m[rowI][col] = n
		}
	}
}

// Identity sets the matrix to be the identity matrix
//
// (visualized)
//
//  [1, 0, 0, 0]
//  [0, 1, 0, 0]
//  [0, 0, 1, 0]
//  [0, 0, 0, 1]
//
func (m *Mat4) Identity() {
	m[0][0] = 1
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0
	m[1][0] = 0
	m[1][1] = 1
	m[1][2] = 0
	m[1][3] = 0
	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 1
	m[2][3] = 0
	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[3][3] = 1
}

// IsIdentity tells weather or not this matrix is the identity matrix
func (m *Mat4) IsIdentity() bool {
	return m[0][0] == 1 && m[0][1] == 0 && m[0][2] == 0 && m[0][3] == 0 && m[1][0] == 0 && m[1][1] == 1 && m[1][2] == 0 && m[1][3] == 0 && m[2][0] == 0 && m[2][1] == 0 && m[2][2] == 1 && m[2][3] == 0 && m[3][0] == 0 && m[3][1] == 0 && m[3][2] == 0 && m[3][3] == 1
}

// Copy returns an new 1:1 copy of this matrix with each element being copied identically.
func (m *Mat4) Clone() *Mat4 {
	return &Mat4{
		{m[0][0], m[0][1], m[0][2], m[0][3]},
		{m[1][0], m[1][1], m[1][2], m[1][3]},
		{m[2][0], m[2][1], m[2][2], m[2][3]},
		{m[3][0], m[3][1], m[3][2], m[3][3]},
	}
}

// Assign assigns all elements in this matrix to the elements in the matrix m2.
func (m *Mat4) Assign(m2 *Mat4) {
	m2[0][0] = m[0][0]
	m2[0][1] = m[0][1]
	m2[0][2] = m[0][2]
	m2[0][3] = m[0][3]
	m2[1][0] = m[1][0]
	m2[1][1] = m[1][1]
	m2[1][2] = m[1][2]
	m2[1][3] = m[1][3]
	m2[2][0] = m[2][0]
	m2[2][1] = m[2][1]
	m2[2][2] = m[2][2]
	m2[2][3] = m[2][3]
	m2[3][0] = m[3][0]
	m2[3][1] = m[3][1]
	m2[3][2] = m[3][2]
	m2[3][3] = m[3][3]
}

// Mat3 returns the upper 3 elements of this matrix as an *Mat3.
//
// (visualized)
//
//  [-, -, -,  ]
//  [-, -, -,  ]
//  [-, -, -,  ]
//  [          ]
//
//func (m *Mat4) Mat3() *Mat3 {
//}

func (m *Mat4) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "Mat4(\n")
	fmt.Fprintf(buf, "    %g, %g, %g, %g,\n", m[0][0], m[0][1], m[0][2], m[0][3])
	fmt.Fprintf(buf, "    %g, %g, %g, %g,\n", m[1][0], m[1][1], m[1][2], m[1][3])
	fmt.Fprintf(buf, "    %g, %g, %g, %g,\n", m[2][0], m[2][1], m[2][2], m[2][3])
	fmt.Fprintf(buf, "    %g, %g, %g, %g,\n", m[3][0], m[3][1], m[3][2], m[3][3])
	fmt.Fprintf(buf, ")\n")
	return buf.String()
}

// Matrix4 returns an new *Mat4 using the specified elements like so:
//
//  [e00, e01, e02, e03]
//  [e10, e11, e12, e13]
//  [e20, e21, e22, e23]
//  [e30, e31, e32, e33]
//
func Matrix4(e00, e01, e02, e03, e10, e11, e12, e13, e20, e21, e22, e23, e30, e31, e32, e33 float64) *Mat4 {
	return &Mat4{
		{e00, e01, e02, e03},
		{e10, e11, e12, e13},
		{e20, e21, e22, e23},
		{e30, e31, e32, e33},
	}
}
