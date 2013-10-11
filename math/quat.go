// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// Quat represents an vector of four components, W, X, Y, and Z.
type Quat struct {
	W, X, Y, Z Real
}

// Copy returns an new 1:1 copy of this Quat
func (a *Quat) Copy() *Quat {
	return &Quat{a.W, a.X, a.Y, a.Z}
}

// String returns an string representation of this quaternion.
func (a *Quat) String() string {
	return fmt.Sprintf("Quat(W=%v, X=%v, Y=%v, Z=%v)", a.W, a.X, a.Y, a.Z)
}

// Assign assigns the W, X, Y, and Z values in this quaternion to the specified values.
func (a *Quat) Assign(w, x, y, z Real) {
	a.W = w
	a.X = x
	a.Y = y
	a.Z = z
}

// Fill sets the W, X, Y, and Z components of this *Quat to the Real, n, parameter.
func (a *Quat) Fill(n Real) {
	a.W = n
	a.X = n
	a.Y = n
	a.Z = n
}

// Equals tells if this quaternion is equal to the other quaternion, by determining if it is within
// the default tolerence for Real equality.
//
// Equality against nil is always false.
func (a *Quat) Equals(b *Quat) bool {
	if b == nil {
		return false
	}
	return a.W.Equals(b.W) && a.X.Equals(b.X) && a.Y.Equals(b.Y) && a.Z.Equals(b.Z)
}

// EqualsTolerence tells if this quaternion is equal to the other quaternion, by determining if it
// is within the specified tolerence for Real equality.
//
// Equality against nil is always false.
func (a *Quat) EqualsTolerence(b *Quat, tolerence Real) bool {
	if b == nil {
		return false
	}
	return a.W.EqualsTolerence(b.W, tolerence) && a.X.EqualsTolerence(b.X, tolerence) && a.Y.EqualsTolerence(b.Y, tolerence) && a.Z.EqualsTolerence(b.Z, tolerence)
}

// Negate negates all components of this quaternion.
//
// The quaternion a itself is returned, for chaining.
func (a *Quat) Negate() *Quat {
	a.W = -a.W
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
	return a
}

// Add returns the result of a + b
func (a *Quat) Add(b *Quat) *Quat {
	return &Quat{
		a.W + b.W,
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

// AddScalar returns the result of a + b
func (a *Quat) AddScalar(b Real) *Quat {
	return &Quat{
		a.W + b,
		a.X + b,
		a.Y + b,
		a.Z + b,
	}
}

// Sub returns the result of a - b
func (a *Quat) Sub(b *Quat) *Quat {
	return &Quat{
		a.W - b.W,
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

// SubScalar returns the result of a - b
func (a *Quat) SubScalar(b Real) *Quat {
	return &Quat{
		a.W - b,
		a.X - b,
		a.Y - b,
		a.Z - b,
	}
}

// Mul returns the result of the two quaternions multiplied (a * b), implimented as follows:
//
//  return &Quat{
//      (b.W * a.W) - (b.X * a.X) - (b.Y * a.Y) - (b.Z * a.Z),
//      (b.X * a.W) + (b.W * a.X) - (b.Z * a.Y) + (b.Y * a.Z),
//      (b.Y * a.W) + (b.Z * a.X) + (b.W * a.Y) - (b.X * a.Z),
//      (b.Z * a.W) - (b.Y * a.X) + (b.X * a.Y) + (b.W * a.Z),
//  }
//
func (a *Quat) Mul(b *Quat) *Quat {
	return &Quat{
		(b.W * a.W) - (b.X * a.X) - (b.Y * a.Y) - (b.Z * a.Z),
		(b.X * a.W) + (b.W * a.X) - (b.Z * a.Y) + (b.Y * a.Z),
		(b.Y * a.W) + (b.Z * a.X) + (b.W * a.Y) - (b.X * a.Z),
		(b.Z * a.W) - (b.Y * a.X) + (b.X * a.Y) + (b.W * a.Z),
	}
}

// MulScalar returns the result of a * b
func (a *Quat) MulScalar(b Real) *Quat {
	return &Quat{
		a.W * b,
		a.X * b,
		a.Y * b,
		a.Z * b,
	}
}

// Div returns the result of a + b
func (a *Quat) Div(b *Quat) *Quat {
	return &Quat{
		a.W / b.W,
		a.X / b.X,
		a.Y / b.Y,
		a.Z / b.Z,
	}
}

// DivScalar returns the result of a / b
func (a *Quat) DivScalar(b Real) *Quat {
	return &Quat{
		a.W / b,
		a.X / b,
		a.Y / b,
		a.Z / b,
	}
}

// Dot returns the dot product of the two quaternions a and b, respectively.
func (a *Quat) Dot(b *Quat) Real {
	return a.W*a.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// LengthSquared returns the squared length of this quaternion.
func (a *Quat) LengthSquared() Real {
	return a.Dot(a)
}

// Length returns the length of this quaternion.
func (a *Quat) Length() Real {
	return Sqrt(a.LengthSquared())
}

// Normalize normalizes this quaternion, returns true if it was normalized or false if it was an
// zero length quaternion.
func (a *Quat) Normalize() bool {
	lengthSquared := a.LengthSquared()

	if lengthSquared.Equals(0) {
		a.W = 0
		a.X = 0
		a.Y = 0
		a.Z = 0
		return false

	} else if lengthSquared.EqualsTolerence(1.0, RealNearZero) {
		length := Sqrt(lengthSquared)
		*a = *a.DivScalar(length)
		return true
	}

	return true
}

// Project returns an new quaternion representing the projection of this quaternion onto the other
// one. The resulting quaternion will be a scalar multiple of onto.
func (a *Quat) Project(b *Quat) *Quat {
	return b.MulScalar(a.Dot(b) / b.LengthSquared())
}

// Min returns an new quaternion representing the smaller components of the two quaternions.
func (a *Quat) Min(b *Quat) *Quat {
	cpy := *a

	if a.W < b.W {
		cpy.W = a.W
	} else {
		cpy.W = b.W
	}

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

// Max returns an new quaternion representing the larger components of the two quaternions.
func (a *Quat) Max(b *Quat) *Quat {
	cpy := *a

	if a.W > b.W {
		cpy.W = a.W
	} else {
		cpy.W = b.W
	}

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

// CompareTolerence sorts the two quaternions lexicographically, componentwise.
//
// Returns -1 if this quaternion sorts before the other one, and returns +1 if it sorts after.
//
// Returns exactly zero if they are equal within the specified tolerence for Real equality.
func (a *Quat) CompareTolerence(b *Quat, tolerence Real) int {
	if a.W.EqualsCompeq(b.W, tolerence) {
		if a.W < b.W {
			return -1
		}
		return 1
	}

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
func (a *Quat) Compare(b *Quat) int {
	return a.CompareTolerence(b, RealNearZero)
}

// LessThan returns a < b
//
// Also see the Compare() and CompareTolerence() functions.
func (a *Quat) LessThan(b *Quat) bool {
	return a.Compare(b) < 0
}

// Lerp returns an new quaternion representing an linear interpolation between the a and b
// quaternions.
//
// The parameter t is interpolation amount (0.0 - 1.0) between the two quaternions.
//
// Short hand for:
//  a.Mul(b.MulScalar(t))
//
func (a *Quat) Lerp(b *Quat, t Real) *Quat {
	return a.Mul(b.MulScalar(t))
}

// IsNan tells if any components of this quaternion are not an number.
func (a *Quat) IsNan() bool {
	return IsNaN(a.W) || IsNaN(a.X) || IsNaN(a.Y) || IsNaN(a.Z)
}

// Conjugate calculates and stores the conjugate of the quaternion a inside this *Quat.
//
// You can use Copy() before calling this if you do not wish for this function to operate in-place.
func (a *Quat) Conjugate() {
	//a.W = a.W
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
}

// TransformVec3 transforms the 3-component vector by the specified quaternion rotation and returns
// the result.
func (a *Quat) TransformVec3(v *Vec3) *Vec3 {
	vecQuat := &Quat{0.0, v.X, v.Y, v.Z}

	conjugate := a.Copy()
	conjugate.Conjugate()

	vecQuat = conjugate.Mul(vecQuat).Mul(a)

	return &Vec3{vecQuat.X, vecQuat.Y, vecQuat.Z}
}

// TransformVec4 transforms the 4-component vector by the specified quaternion rotation and returns
// the result.
func (a *Quat) TransformVec4(v *Vec4) *Vec4 {
	vecQuat := &Quat{v.W, v.X, v.Y, v.Z}

	conjugate := a.Copy()
	conjugate.Conjugate()

	vecQuat = conjugate.Mul(vecQuat).Mul(a)

	return vecQuat.Vec4()
}

// Invert inverts the quaternion and stores the result inside this *Quat.
//
// You can use Copy() before calling this if you do not wish for this function to operate in-place.
func (a *Quat) Invert() {
	a.W = -a.W
	//a.X = a.X
	//a.Y = a.y
	//a.Z = a.Z
}

// EqualsIdentityTolerence tells weather this Quat represents the identity transformation within
// the given tolerence for Real equality.
func (a *Quat) EqualsIdentityTolerence(tolerence Real) bool {
	return a.W.EqualsTolerence(-1.0, tolerence) || a.W.EqualsTolerence(1.0, tolerence)
}

// EqualsIdentity tells weather this Quat represents the identity transformation within the default
// tolerence for Real equality.
func (a *Quat) EqualsIdentity() bool {
	return a.EqualsIdentityTolerence(RealNearZero)
}

// SameDirectionTolerence tells weather this quaternion represents the same rotation as the other
// quaternion as told by the specified tolerence for Real equality.
func (a *Quat) SameDirectionTolerence(b *Quat, tolerence Real) bool {
	bInverted := b.Copy()
	bInverted.Invert()
	return a.Mul(bInverted).EqualsIdentityTolerence(tolerence)
}

// DirectionTolerence tells weather this quaternion represents the same rotation as the other
// quaternion as told by the default tolerence for Real equality.
func (a *Quat) SameDirection(b *Quat) bool {
	return a.SameDirectionTolerence(b, RealNearZero)
}

// Up returns the orientation represented by this quaternion, expressed as an up vector.
func (a *Quat) Up(cs CoordSys) *Vec3 {
	return a.TransformVec3(cs.Up())
}

// Right returns the orientation represented by this quaternion, expressed as an right vector.
func (a *Quat) Right(cs CoordSys) *Vec3 {
	return a.TransformVec3(cs.Right())
}

// Forward returns the orientation represented by this quaternion, expressed as an forward vector.
func (a *Quat) Forward(cs CoordSys) *Vec3 {
	return a.TransformVec3(cs.Forward())
}

// Angle returns the angle between the orientation represented by this quaternion and the
// orientation represent by the quaternion b, in radians.
func (a *Quat) AngleQuat(cs CoordSys, b *Quat) Real {
	return a.Forward(cs).Angle(b.Forward(cs))
}

// Axis returns the axis of the rotation represented by the quaternion.
//
// The returned vector is not normalized, you can normalize it yourself using the Normalize()
// function.
func (a *Quat) Axis() *Vec3 {
	return &Vec3{a.X, a.Y, a.Z}
}

// Angle returns the rotation represented by the quaternion as an angle about an arbitrary axis
// (returned by the Axis() function).
//
// The return value is expressed in radian units counterclockwise about the axis.
//
// You must normalize the quaternion before calling this function (See the Normalize() function).
func (a *Quat) Angle() Real {
	return Acos(a.W) * 2
}

// ExtractToMat3 Extracts the quaternion into the specified matrix.
//
// If the parameter to is nil, an panic occurs.
func (a *Quat) ExtractToMat3(to *Mat3) {
	if to == nil {
		panic("ExtractToMat3(): parameter 'to' is nil!")
	}

	n := a.Dot(a)

	s := Real(0)
	if !n.Equals(0) {
		s = 2.0 / n
	}

	xs := a.X * s
	ys := a.Y * s
	zs := a.Z * s
	wx := a.W * xs
	wy := a.W * ys
	wz := a.W * zs
	xx := a.X * xs
	xy := a.X * ys
	xz := a.X * zs
	yy := a.Y * ys
	yz := a.Y * zs
	zz := a.Z * zs

	to.Assign(
		1.0-(yy+zz), xy+wz, xz-wy,
		xy-wz, 1.0-(xx+zz), yz+wx,
		xz+wy, yz-wx, 1.0-(xx+yy),
	)
}

// ExtractToMat4 Extracts the quaternion into the specified matrix.
//
// If the parameter to is nil, an panic occurs.
func (a *Quat) ExtractToMat4(to *Mat4) {
	if to == nil {
		panic("ExtractToMat4(): parameter 'to' is nil!")
	}

	n := a.Dot(a)

	s := Real(0)
	if !n.Equals(0) {
		s = 2.0 / n
	}

	xs := a.X * s
	ys := a.Y * s
	zs := a.Z * s
	wx := a.W * xs
	wy := a.W * ys
	wz := a.W * zs
	xx := a.X * xs
	xy := a.X * ys
	xz := a.X * zs
	yy := a.Y * ys
	yz := a.Y * zs
	zz := a.Z * zs

	to.Assign(
		1.0-(yy+zz), xy+wz, xz-wy, 0,
		xy-wz, 1.0-(xx+zz), yz+wx, 0,
		xz+wy, yz-wx, 1.0-(xx+yy), 0,
		0, 0, 0, 1,
	)
}

// MulMat3 multiplies the extracted Mat3 from this quaternion against the specified matrix m and
// returns the result.
//
// Quat * Mat3 = Mat3
func (a *Quat) MulMat3(m *Mat3) *Mat3 {
	result := new(Mat3)
	a.ExtractToMat3(result)
	return result.Mul(m)
}

// MulMat4 multiplies the extracted Mat4 from this quaternion against the specified matrix m and
// returns the result.
//
// Quat * Mat4 = Mat4
func (a *Quat) MulMat4(m *Mat4) *Mat4 {
	upper3 := m.UpperMat3()

	quatMat := new(Mat3)
	a.ExtractToMat3(quatMat)

	result := new(Mat4)
	result.SetUpperMat3(quatMat.Mul(upper3))
	result.SetRow(3, m.Row(3))
	result.SetCol(3, m.Col(3))
	return result
}

// SetFromMat3 sets the quaternion according to the rotation represented by the matrix.
func (a *Quat) SetFromMat3(m *Mat3) {
	m00 := m[0][0]
	m10 := m[1][0]
	m20 := m[2][0]
	m01 := m[0][1]
	m11 := m[1][1]
	m21 := m[2][1]
	m02 := m[0][2]
	m12 := m[1][2]
	m22 := m[2][2]

	trace := m00 + m11 + m22

	if trace > 0 {
		// The easy case.
		s := Sqrt(trace + 1.0)
		a.W = s * 0.5
		s = 0.5 / s
		a.X = (m12 - m21) * s
		a.Y = (m20 - m02) * s
		a.Z = (m01 - m10) * s
	} else {
		// The harder case.  First, figure out which column to take as
		// root.  This will be the column with the largest value.
		//
		// It is tempting to try to compare the absolute values of the
		// diagonal values in the code below, instead of their normal,
		// signed values.  Don't do it.  We are actually maximizing the
		// value of S, which must always be positive, and is therefore
		// based on the diagonal whose actual value--not absolute
		// value--is greater than those of the other two.
		//
		// We already know that m00 + m11 + m22 <= 0 (because we are here
		// in the harder case).
		if m00 > m11 && m00 > m22 {
			// m00 is larger than m11 and m22.
			s := 1.0 + m00 - (m11 + m22)
			s = Sqrt(s)
			a.X = s * 0.5
			s = 0.5 / s
			a.Y = (m01 + m10) * s
			a.Z = (m02 + m20) * s
			a.W = (m12 - m21) * s

		} else if m11 > m22 {
			// m11 is larger than m00 and m22.
			s := 1.0 + m11 - (m22 + m00)
			s = Sqrt(s)
			a.Y = s * 0.5
			s = 0.5 / s
			a.Z = (m12 + m21) * s
			a.X = (m10 + m01) * s
			a.W = (m20 - m02) * s

		} else {
			// m22 is larger than m00 and m11.
			s := 1.0 + m22 - (m00 + m11)
			s = Sqrt(s)
			a.Z = s * 0.5
			s = 0.5 / s
			a.X = (m20 + m02) * s
			a.Y = (m21 + m12) * s
			a.W = (m01 - m10) * s
		}
	}
}

// SetFromMat4 sets the quaternion according to the rotation represented by the matrix.
func (a *Quat) SetFromMat4(m *Mat4) {
	a.SetFromMat3(m.UpperMat3())
}

// SetFromAngleAxis sets the quaternion to the angle (in radians) about the axis.
//
// The axis vector should be normalized before calling this function.
func (a *Quat) SetFromAngleAxis(angle Real, axis *Vec3) {
	if !axis.Length().EqualsTolerence(1.0, 0.001) {
		panic("SetFromAngleAxis(): Axis vector non-normalized!")
	}

	sinHalfAngle := Sin(angle * 0.5)
	a.W = Cos(angle * 0.5)
	a.X = axis.X * sinHalfAngle
	a.Y = axis.Y * sinHalfAngle
	a.Z = axis.Z * sinHalfAngle
}

// SetHpr sets the quaternion as the unit quaternion that is equivalent to these heading, pitch,
// and roll Euler angles for the given coordinate system.
//
// The Euler angles are expected to be in Radians.
func (a *Quat) SetHpr(hpr *Vec3, cs CoordSys) {
	v := cs.Up()
	n := hpr.X * 0.5
	s := Sin(n)
	c := Cos(n)
	quatHeading := &Quat{c, v.X * s, v.Y * s, v.Z * s}

	v = cs.Right()
	n = hpr.Y * 0.5
	s = Sin(n)
	c = Cos(n)
	quatPitch := &Quat{c, v.X * s, v.Y * s, v.Z * s}

	v = cs.Forward()
	n = hpr.Z * 0.5
	s = Sin(n)
	c = Cos(n)
	quatRoll := &Quat{c, v.X * s, v.Y * s, v.Z * s}

	if cs.RightHanded() {
		*a = *quatRoll.Mul(quatPitch).Mul(quatHeading)
	} else {
		result := quatHeading.Mul(quatPitch).Mul(quatRoll)
		result.Invert()
		*a = *result
	}
}

// Hpr extracts the equivalent heading, pitch, and roll Euler angles from the unit quaternion for
// the given coordinate system.
//
// The Euler angles are expressed in Radians.
func (a *Quat) Hpr(cs CoordSys) (hpr *Vec3) {
	hpr = new(Vec3)

	if cs == CoordSysZUpRight && false {
		n := a.Dot(a)

		s := Real(0)
		if !n.Equals(0) {
			s = 2.0 / n
		}

		xs := a.X * s
		ys := a.Y * s
		zs := a.Z * s
		wx := a.W * xs
		wy := a.W * ys
		wz := a.W * zs
		xx := a.X * xs
		xy := a.X * ys
		xz := a.X * zs
		yy := a.Y * ys
		yz := a.Y * zs
		zz := a.Z * zs
		c1 := xz - wy
		c2 := 1.0 - (xx + yy)
		c3 := 1.0 - (yy + zz)
		c4 := xy + wz

		var ch, sh, cp, cr, sr, sp Real

		if c1.Equals(0) {
			// (roll = 0 or 180) or (pitch = +/- 90)
			if c2 >= 0 {
				hpr.Z = 0
				ch = c3
				sh = c4
				cp = c2
			} else {
				hpr.Z = Real(180).Radians()
				ch = -c3
				sh = -c4
				cp = -c2
			}
		} else {
			// this should work all the time, but the above saves some trig operations
			hpr.Z = Atan2(-c1, c2)
			sr = Sin(hpr.Z)
			cr = Cos(hpr.Z)
			ch = (cr * c3) + (sr * (xz + wy))
			sh = (cr * c4) + (sr * (yz - wx))
			cp = (cr * c2) - (sr * c1)
		}
		sp = yz + wx
		hpr.X = Atan2(sh, ch)
		hpr.Y = Atan2(sp, cp)

	} else {
		// The code above implements quat-to-hpr for CoordSysZUpRight only.
		//
		// For other coordinate systems, someone is welcome to extend the
		// implementation; I'm going to choose the slower, lazy path till then.
		mat := new(Mat3)
		a.ExtractToMat3(mat)
		_, _, hpr = mat.Decompose(cs)
	}

	return
}

// Vec4 converts the quaternion into an vector. (No magic -- just an 'type cast')
func (a *Quat) Vec4() *Vec4 {
	return &Vec4{a.W, a.X, a.Y, a.Z}
}

// Quaternion returns an new *Quat with the specified values.
func Quaternion(w, x, y, z Real) *Quat {
	return &Quat{w, x, y, z}
}

var (
	QuatIdentity = Quaternion(1, 0, 0, 0)
)
