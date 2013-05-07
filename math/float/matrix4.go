// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

// Matrix4 represents an four by four matrix.
//
// The most immediate array specifies the row, and the second most immediate array specifies the
// column.
//
//  [0x0, 0x1, 0x2, 0x3]
//  [1x0, 1x1, 1x2, 1x3]
//  [2x0, 2x1, 2x2, 2x3]
//  [3x0, 3x1, 3x2, 3x3]
//
// where (row)x(column) notation was used, you could access each element visualized above using:
//
//  m[row][column]
//
type Matrix4 [4][4]Real

func (m Matrix4) MultiplyVector3(v Vector3) Vector3 {
	var (
		r     Vector3
		fInvW Real = 1.0 / (m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3])
	)

	r.X = (m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]) * fInvW
	r.Y = (m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]) * fInvW
	r.Z = (m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]) * fInvW
	return r
}

func (m Matrix4) MultiplyVector4(v Vector4) Vector4 {
	var r Vector4
	r.X = m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W
	r.Y = m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W
	r.Z = m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W
	r.W = m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]*v.W
	return r
}

/*

func (m Matrix4) MultiplyPlane(p Plane) Plane {
	var r Plane
	invTrans := m.Inverse().Transpose()
	v4 := Vector4f(p.Normal.X, p.Normal.Y, p.Normal.Z, p.D)
	v4 = invTrans.Multiply(v4)
	r.Normal.X = v4.X
	r.Normal.Y = v4.Y
	r.Normal.Z = v4.Z
	r.D = v4.W / r.Normal.Normalize()
	return r

//237         inline Plane operator * (const Plane& p) const
//238         {
//239             Plane ret;
//240             Matrix4 invTrans = inverse().transpose();
//241             Vector4 v4( p.normal.x, p.normal.y, p.normal.z, p.d );
//242             v4 = invTrans * v4;
//243             ret.normal.x = v4.x;
//244             ret.normal.y = v4.y;
//245             ret.normal.z = v4.z;
//246             ret.d = v4.w / ret.normal.normalise();
//247
//248             return ret;
}

func (m Matrix4) SetMatrix3f(m2 Matrix3f) {
 inline void operator = ( const Matrix3& mat3 )
  338         {
  339             m[0][0] = mat3.m[0][0]; m[0][1] = mat3.m[0][1]; m[0][2] = mat3.m[0][2]
  340             m[1][0] = mat3.m[1][0]; m[1][1] = mat3.m[1][1]; m[1][2] = mat3.m[1][2]
  341             m[2][0] = mat3.m[2][0]; m[2][1] = mat3.m[2][1]; m[2][2] = mat3.m[2][2]
  342         }
}
*/

// Multiply multiplies this matrix and the matrix m2 and returns the result.
func (m Matrix4) Multiply(m2 Matrix4) Matrix4 {
	var r Matrix4
	r[0][0] = m[0][0]*m2[0][0] + m[0][1]*m2[1][0] + m[0][2]*m2[2][0] + m[0][3]*m2[3][0]
	r[0][1] = m[0][0]*m2[0][1] + m[0][1]*m2[1][1] + m[0][2]*m2[2][1] + m[0][3]*m2[3][1]
	r[0][2] = m[0][0]*m2[0][2] + m[0][1]*m2[1][2] + m[0][2]*m2[2][2] + m[0][3]*m2[3][2]
	r[0][3] = m[0][0]*m2[0][3] + m[0][1]*m2[1][3] + m[0][2]*m2[2][3] + m[0][3]*m2[3][3]
	r[1][0] = m[1][0]*m2[0][0] + m[1][1]*m2[1][0] + m[1][2]*m2[2][0] + m[1][3]*m2[3][0]
	r[1][1] = m[1][0]*m2[0][1] + m[1][1]*m2[1][1] + m[1][2]*m2[2][1] + m[1][3]*m2[3][1]
	r[1][2] = m[1][0]*m2[0][2] + m[1][1]*m2[1][2] + m[1][2]*m2[2][2] + m[1][3]*m2[3][2]
	r[1][3] = m[1][0]*m2[0][3] + m[1][1]*m2[1][3] + m[1][2]*m2[2][3] + m[1][3]*m2[3][3]
	r[2][0] = m[2][0]*m2[0][0] + m[2][1]*m2[1][0] + m[2][2]*m2[2][0] + m[2][3]*m2[3][0]
	r[2][1] = m[2][0]*m2[0][1] + m[2][1]*m2[1][1] + m[2][2]*m2[2][1] + m[2][3]*m2[3][1]
	r[2][2] = m[2][0]*m2[0][2] + m[2][1]*m2[1][2] + m[2][2]*m2[2][2] + m[2][3]*m2[3][2]
	r[2][3] = m[2][0]*m2[0][3] + m[2][1]*m2[1][3] + m[2][2]*m2[2][3] + m[2][3]*m2[3][3]
	r[3][0] = m[3][0]*m2[0][0] + m[3][1]*m2[1][0] + m[3][2]*m2[2][0] + m[3][3]*m2[3][0]
	r[3][1] = m[3][0]*m2[0][1] + m[3][1]*m2[1][1] + m[3][2]*m2[2][1] + m[3][3]*m2[3][1]
	r[3][2] = m[3][0]*m2[0][2] + m[3][1]*m2[1][2] + m[3][2]*m2[2][2] + m[3][3]*m2[3][2]
	r[3][3] = m[3][0]*m2[0][3] + m[3][1]*m2[1][3] + m[3][2]*m2[2][3] + m[3][3]*m2[3][3]
	return r
}

// Add adds the matrix m2 to this matrix and returns the result.
func (m Matrix4) Add(m2 Matrix4) Matrix4 {
	var r Matrix4
	r[0][0] = m[0][0] + m2[0][0]
	r[0][1] = m[0][1] + m2[0][1]
	r[0][2] = m[0][2] + m2[0][2]
	r[0][3] = m[0][3] + m2[0][3]

	r[1][0] = m[1][0] + m2[1][0]
	r[1][1] = m[1][1] + m2[1][1]
	r[1][2] = m[1][2] + m2[1][2]
	r[1][3] = m[1][3] + m2[1][3]

	r[2][0] = m[2][0] + m2[2][0]
	r[2][1] = m[2][1] + m2[2][1]
	r[2][2] = m[2][2] + m2[2][2]
	r[2][3] = m[2][3] + m2[2][3]

	r[3][0] = m[3][0] + m2[3][0]
	r[3][1] = m[3][1] + m2[3][1]
	r[3][2] = m[3][2] + m2[3][2]
	r[3][3] = m[3][3] + m2[3][3]
	return r
}

// Subtract subtracts the matrix m2 from this matrix and returns the result.
func (m Matrix4) Subtract(m2 Matrix4) Matrix4 {
	var r Matrix4
	r[0][0] = m[0][0] - m2[0][0]
	r[0][1] = m[0][1] - m2[0][1]
	r[0][2] = m[0][2] - m2[0][2]
	r[0][3] = m[0][3] - m2[0][3]

	r[1][0] = m[1][0] - m2[1][0]
	r[1][1] = m[1][1] - m2[1][1]
	r[1][2] = m[1][2] - m2[1][2]
	r[1][3] = m[1][3] - m2[1][3]

	r[2][0] = m[2][0] - m2[2][0]
	r[2][1] = m[2][1] - m2[2][1]
	r[2][2] = m[2][2] - m2[2][2]
	r[2][3] = m[2][3] - m2[2][3]

	r[3][0] = m[3][0] - m2[3][0]
	r[3][1] = m[3][1] - m2[3][1]
	r[3][2] = m[3][2] - m2[3][2]
	r[3][3] = m[3][3] - m2[3][3]
	return r
}

// Equals tells weather or not this matrix is equal to the matrix m2.
func (m Matrix4) Equals(m2 Matrix4) bool {
	if m[0][0] != m2[0][0] || m[0][1] != m2[0][1] || m[0][2] != m2[0][2] || m[0][3] != m[0][3] || m[1][0] != m2[1][0] || m[1][1] != m2[1][1] || m[1][2] != m2[1][2] || m[1][3] != m[1][3] || m[2][0] != m2[2][0] || m[2][1] != m2[2][1] || m[2][2] != m2[2][2] || m[2][3] != m[2][3] || m[3][0] != m2[3][0] || m[3][1] != m2[3][1] || m[3][2] != m2[3][2] || m[3][3] != m[3][3] {
		return false
	}
	return true
}

// Transpose returns the matrix m transposed.
//
// For instance the following matrix:
//
//  [01, 02, 03, 04]
//  [05, 06, 07, 08]
//  [09, 10, 11, 12]
//  [13, 14, 15, 16]
//
// Would become:
//
//  [01, 05, 09, 13]
//  [02, 06, 10, 14]
//  [03, 07, 11, 15]
//  [04, 08, 12, 16]
//
func (m Matrix4) Transpose() Matrix4 {
	return NewMatrix4(
		m[0][0], m[1][0], m[2][0], m[3][0],
		m[0][1], m[1][1], m[2][1], m[3][1],
		m[0][2], m[1][2], m[2][2], m[3][2],
		m[0][3], m[1][3], m[2][3], m[3][3],
	)
}

// SetTranslation sets the translation (X, Y, and Z) elements of this matrix to the specified
// vector.
//
//  [-, -, -, X]
//  [-, -, -, Y]
//  [-, -, -, Z]
//  [-, -, -, -]
func (m Matrix4) SetTranslation(v Vector3) {
	m[0][3] = v.X
	m[1][3] = v.Y
	m[2][3] = v.Z
}

// Translation returns the translation (X, Y, and Z) elements of this matrix as an vector.
//
//  [-, -, -, X]
//  [-, -, -, Y]
//  [-, -, -, Z]
//  [-, -, -, -]
func (m Matrix4) Translation() Vector3 {
	return Vector3{m[0][3], m[1][3], m[2][3]}
}

// MakeTranslation sets the elements of this matrix such that it is an translation matrix using
// the vector, v.
//
//  [1, 0, 0, X]
//  [0, 1, 0, Y]
//  [0, 0, 1, Z]
//  [0, 0, 0, 1]
func (m Matrix4) MakeTranslation(v Vector3) {
	m[0][0] = 1.0
	m[0][1] = 0.0
	m[0][2] = 0.0
	m[0][3] = v.X
	m[1][0] = 0.0
	m[1][1] = 1.0
	m[1][2] = 0.0
	m[1][3] = v.Y
	m[2][0] = 0.0
	m[2][1] = 0.0
	m[2][2] = 1.0
	m[2][3] = v.Z
	m[3][0] = 0.0
	m[3][1] = 0.0
	m[3][2] = 0.0
	m[3][3] = 1.0
}

// SetScale sets the scale (X, Y, and Z) elements of this matrix to the specified vector.
//
//  [X, -, -, -]
//  [-, Y, -, -]
//  [-, -, Z, -]
//  [-, -, -, -]
func (m Matrix4) SetScale(v Vector3) {
	m[0][0] = v.X
	m[1][1] = v.Y
	m[2][2] = v.Z
}

// Scale returns the scale (X, Y, and Z) elements of this matrix as an vector.
//
//  [X, -, -, -]
//  [-, Y, -, -]
//  [-, -, Z, -]
//  [-, -, -, -]
func (m Matrix4) Scale() Vector3 {
	return Vector3{m[0][0], m[1][1], m[2][2]}
}

// Matrix3 returns the first three columns/rows from this matrix and returns them as an new Matrix3
//
//  [-, -, -,  ]
//  [-, -, -,  ]
//  [-, -, -,  ]
//  [ ,  ,  ,  ]
func (m Matrix4) Matrix3() Matrix3 {
	var n Matrix3
	n[0][0] = m[0][0]
	n[0][1] = m[0][1]
	n[0][2] = m[0][2]

	n[1][0] = m[1][0]
	n[1][1] = m[1][1]
	n[1][2] = m[1][2]

	n[2][0] = m[2][0]
	n[2][1] = m[2][1]
	n[2][2] = m[2][2]
	return n
}

// HasScale tells weather or not this matrix has an scaling
func (m Matrix4) HasScale() bool {
	var t Real

	// check magnitude of column vectors (==local axes)
	t = m[0][0]*m[0][0] + m[1][0]*m[1][0] + m[2][0]*m[2][0]
	if t == 1.0 {
		return true
	}
	t = m[0][1]*m[0][1] + m[1][1]*m[1][1] + m[2][1]*m[2][1]
	if t == 1.0 {
		return true
	}
	t = m[0][2]*m[0][2] + m[1][2]*m[1][2] + m[2][2]*m[2][2]
	if t == 1.0 {
		return true
	}
	return false
}

// HasNegativeScale tells tells weather or not this matrix has an negative scaling
func (m Matrix4) HasNegativeScale() bool {
	return m.Determinant() < 0
}

/*
func (m Matrix4) Quaternion() Quaternion {
	quaternion.FromMatrix3f(m.Matrix3())
}
*/

/*
  509     static const Matrix4 ZERO;
  510     static const Matrix4 ZEROAFFINE;
  511     static const Matrix4 IDENTITY;
  514         static const Matrix4 CLIPSPACE2DTOIMAGESPACE;
*/

// ScalarMultiply performs an scalar multiplication against the scalar s and returns the result.
func (m Matrix4) ScalarMultiply(s Real) {
	return NewMatrix4(
		s*m[0][0], s*m[0][1], s*m[0][2], s*m[0][3],
		s*m[1][0], s*m[1][1], s*m[1][2], s*m[1][3],
		s*m[2][0], s*m[2][1], s*m[2][2], s*m[2][3],
		s*m[3][0], s*m[3][1], s*m[3][2], s*m[3][3],
	)
}

// String returns an string representation of this matrix.
func (m Matrix4) String() string {
	var b bytes.Buffer
	fmt.Fprintf(b, "Matrix4(")
	for i, row := range m {
		fmt.Fprintf(b, " Row%d{", i)
		for j, col := range m[row] {
			fmt.Fprintf(b, "%d ", col)
		}
		fmt.Fprintf(b, "}")
	}
	fmt.Fprintf(b, ")")
}

/*
  543
  544         Matrix4 adjoint() const;
  545         Real determinant() const;
  546         Matrix4 inverse() const;
  547
  554         void makeTransform(const Vector3& position, const Vector3& scale, const Quaternion& orientation);
  555
  561         void makeInverseTransform(const Vector3& position, const Vector3& scale, const Quaternion& orientation);
  562
  565         void decomposition(Vector3& position, Vector3& scale, Quaternion& orientation) const;
*/

func (m Matrix4) Affine() bool {
	return m[3][0] == 0 && m[3][1] == 0 && m[3][2] == 0 && m[3][3] == 1
}

/*
  581         Matrix4 inverseAffine(void) const;
*/

func (m Matrix4) AffineMultiply(m2 Matrix4) {
	if !m.Affine() || !m2.Affine() {
		panic("Matrix4.AffineMultiply(): Matrix is not affine!")
	}

	return NewMatrix4(
		m[0][0]*m2[0][0]+m[0][1]*m2[1][0]+m[0][2]*m2[2][0],
		m[0][0]*m2[0][1]+m[0][1]*m2[1][1]+m[0][2]*m2[2][1],
		m[0][0]*m2[0][2]+m[0][1]*m2[1][2]+m[0][2]*m2[2][2],
		m[0][0]*m2[0][3]+m[0][1]*m2[1][3]+m[0][2]*m2[2][3]+m[0][3],

		m[1][0]*m2[0][0]+m[1][1]*m2[1][0]+m[1][2]*m2[2][0],
		m[1][0]*m2[0][1]+m[1][1]*m2[1][1]+m[1][2]*m2[2][1],
		m[1][0]*m2[0][2]+m[1][1]*m2[1][2]+m[1][2]*m2[2][2],
		m[1][0]*m2[0][3]+m[1][1]*m2[1][3]+m[1][2]*m2[2][3]+m[1][3],

		m[2][0]*m2[0][0]+m[2][1]*m2[1][0]+m[2][2]*m2[2][0],
		m[2][0]*m2[0][1]+m[2][1]*m2[1][1]+m[2][2]*m2[2][1],
		m[2][0]*m2[0][2]+m[2][1]*m2[1][2]+m[2][2]*m2[2][2],
		m[2][0]*m2[0][3]+m[2][1]*m2[1][3]+m[2][2]*m2[2][3]+m[2][3],

		0, 0, 0, 1,
	)
}

func (m Matrix4) AffineTransformVector3(v Vector3) Vector4 {
	if !m.Affine() {
		panic("Matrix4.AffineTransformVector3(): Matrix is not affine!")
	}

	return Vector3{
		m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3],
		m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3],
		m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3],
	}
}

func (m Matrix4) AffineTransformVector4(v Vector4) Vector4 {
	if !m.Affine() {
		panic("Matrix4.AffineTransformVector4(): Matrix is not affine!")
	}

	return Vector4{
		m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W,
		m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W,
		m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W,
		v.W,
	}
}

// NewMatrix4 returns an new matrix whose elements are set to the given row by column parameters.
//
//  [v00, v01, v02, v03]
//  [v10, v11, v12, v13]
//  [v20, v21, v22, v23]
//  [v30, v31, v32, v33]
func NewMatrix4(v00, v01, v02, v03, v10, v11, v12, v13, v20, v21, v22, v23, v30, v31, v32, v33 Real) Matrix4 {
	return Matrix4{
		{v00, v01, v02, v03},
		{v10, v11, v12, v13},
		{v20, v21, v22, v23},
		{v30, v31, v32, v33},
	}
}

// TranslationMatrix4 returns an new matrix whose translation (X, Y, and Z) elements are set to
// the given vector.
//
//  [1, 0, 0, X]
//  [0, 1, 0, Y]
//  [0, 0, 1, Z]
//  [0, 0, 0, 1]
func TranslationMatrix(v Vector3) Matrix4 {
	var r Matrix4
	r[0][0] = 1.0
	r[0][1] = 0.0
	r[0][2] = 0.0
	r[0][3] = v.X
	r[1][0] = 0.0
	r[1][1] = 1.0
	r[1][2] = 0.0
	r[1][3] = v.Y
	r[2][0] = 0.0
	r[2][1] = 0.0
	r[2][2] = 1.0
	r[2][3] = v.Z
	r[3][0] = 0.0
	r[3][1] = 0.0
	r[3][2] = 0.0
	r[3][3] = 1.0
	return r
}

// ScaleMatrix4 returns an new matrix whose scale (X, Y, and Z) elements are set to the given
// vector.
//
//  [X, 0, 0, 0]
//  [0, Y, 0, 0]
//  [0, 0, Z, 0]
//  [0, 0, 0, 0]
func (m Matrix4) ScaleMatrix(v Vector3) Matrix4 {
	var r Matrix4
	r[0][0] = v.X
	r[0][1] = 0.0
	r[0][2] = 0.0
	r[0][3] = 0.0
	r[1][0] = 0.0
	r[1][1] = v.Y
	r[1][2] = 0.0
	r[1][3] = 0.0
	r[2][0] = 0.0
	r[2][1] = 0.0
	r[2][2] = v.Z
	r[2][3] = 0.0
	r[3][0] = 0.0
	r[3][1] = 0.0
	r[3][2] = 0.0
	r[3][3] = 1.0
	return r
}
