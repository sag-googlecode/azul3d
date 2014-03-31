// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"azul3d.org/v1/math"
	"sync"
)

// Transform represents a three-dimensional transformation.
//
// A transform may be made up of components such as position,
// rotation/quaternion, scale, shear or may be an arbitrary transformation
// matrix.
type Transform struct {
	access sync.RWMutex

	arbitrary                        bool
	built                            math.Mat4
	position, rotation, scale, shear math.Vec3
	quat                             math.Quat
}

func (t *Transform) hasComponents() bool {
	if t.arbitrary {
		return false
	}

	positionZero := t.position.Equals(math.Vec3Zero)
	rotationZero := t.rotation.Equals(math.Vec3Zero)
	scaleZero := t.scale.Equals(math.Vec3Zero)
	shearZero := t.shear.Equals(math.Vec3Zero)
	quatZero := t.quat.Equals(math.QuatZero)
	return !positionZero || !rotationZero || !scaleZero || !shearZero || !quatZero
}

// HasComponents tells if this transformation has Pos(), Rot(), Scale(),
// Shear(), or Quat() components specified.
//
// Additionally, if this transformation has an arbitrary matrix set on it,
// false is returned.
func (t *Transform) HasComponents() bool {
	t.access.RLock()
	defer t.access.RUnlock()

	return t.hasComponents()
}

// Compose returns the composition of t * other, I.e. a composed transformation
// that describes the two transformations combined.
func (t *Transform) Compose(other *Transform) *Transform {
	composed := new(Transform)
	c := other.Mat4().Mul(t.Mat4())
	composed.SetMat4(c)
	return composed
}

func (t *Transform) build() {
	if t.arbitrary {
		// An arbitrary transformation matrix has been set, so we do not build
		// an built matrix on our own, instead we simply use that one (which
		// is already set as the built matrix).
		return
	}

	// Set to identity matrix to clear old transformation
	t.built = math.Mat4Identity

	// If we don't have components to build with, we can just leave it as the
	// identity matrix above.
	if !t.hasComponents() {
		return
	}

	scale := math.Vec3One
	if !t.scale.Equals(math.Vec3Zero) {
		scale = t.scale
	}

	shear := math.Vec3Zero
	if !t.shear.Equals(math.Vec3Zero) {
		// We expect shear in an different format from the math package. So
		// handle this as we want here.
		shear = t.shear
	}

	// Apply translation
	pos := math.Vec3Zero
	if !t.position.Equals(math.Vec3Zero) {
		pos = t.position
	}

	// Apply rotation
	hpr := math.Vec3Zero
	if !t.rotation.Equals(math.Vec3Zero) {
		// We have euler rotation
		hpr = t.rotation.Radians().XyzToHpr()

	} else if !t.quat.Equals(math.QuatZero) {
		// We have quaternion rotation
		hpr = t.quat.Hpr(math.CoordSysZUpRight)
	}

	m := math.Mat3Compose(scale, shear, hpr, math.CoordSysZUpRight)
	t.built = t.built.SetUpperMat3(m)

	trans := math.Mat4FromTranslation(pos)
	t.built = t.built.Mul(trans)
}

// SetArbitraryMat4 specifies an arbitrary transformation matrix to represent
// this transformation.
//
// Once an arbitrary transformation matrix is set, this transform's components
// (such as position, rotation/quaternion, shear, etc) are not used.
func (t *Transform) SetArbitraryMat4(m math.Mat4) {
	t.access.Lock()
	defer t.access.Unlock()

	if m.Equals(math.Mat4Zeros) {
		t.arbitrary = false
		t.built = math.Mat4Zeros
	} else {
		t.arbitrary = true
		t.built = m
	}
}

// ArbitraryMat4 returns the arbitrary transformation matrix of this
// transformation,
//
// math.Mat4Zeros is returned if there is no arbitrary transformation matrix
// set.
func (t *Transform) ArbitraryMat4() math.Mat4 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.arbitrary {
		return t.built
	}
	return math.Mat4Zeros
}

// SetMat4 specifies the position, scale, shear, and rotation components of
// this transformation to the ones decomposed from the specified matrix.
//
// This attempts to decompose the matrix into position, scale, shear, and
// rotation values which are then set on this transform. If the transformation
// matrix would not decompose nicely in this way, you may want to set an
// arbitrary transformation matrix instead using SetArbitraryMat4() which does
// not attempt to decompose the matrix.
func (t *Transform) SetMat4(m math.Mat4) {
	t.access.Lock()
	defer t.access.Unlock()

	t.position = m.Translation()
	t.scale, t.shear, t.rotation = m.UpperMat3().Decompose(math.CoordSysZUpRight)

	// Decompose returns HPR, we expect RX/RY/RZ, and degrees.
	t.rotation = t.rotation.Degrees().HprToXyz()
}

// Mat4 returns the matrix which perfectly defines this transformation.
//
// This is the final transformation matrix. Weather this transform is
// represented by an arbitrary transformation matrix or by components (i.e.
// position, rotation/quaternion, scale, shear, etc) a matrix is always
// returned.
func (t *Transform) Mat4() math.Mat4 {
	t.access.Lock()
	defer t.access.Unlock()

	if t.built.Equals(math.Mat4Zeros) {
		t.build()
	}
	return t.built
}

// SetPos specifies the position component of this transformation.
func (t *Transform) SetPos(p math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	if !p.Equals(t.position) {
		t.position = p
		t.built = math.Mat4Zeros
	}
}

// Pos returns the position component of this transformation.
func (t *Transform) Pos() math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()
	return t.position
}

// SetRot specifies the rotation component of this transformation.
func (t *Transform) SetRot(r math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	if !r.Equals(t.rotation) {
		t.rotation = r
		t.built = math.Mat4Zeros
	}
}

// Rot returns the rotation component of this transformation.
func (t *Transform) Rot() math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.rotation.Equals(math.Vec3Zero) {
		if !t.quat.Equals(math.QuatZero) {
			hpr := t.quat.Hpr(math.CoordSysZUpRight)
			return hpr.Degrees().HprToXyz()
		} else {
			return math.Vec3Zero
		}
	}
	return t.rotation
}

// SetQuat specifies the quaternion rotation component of this transformation.
func (t *Transform) SetQuat(q math.Quat) {
	t.access.Lock()
	defer t.access.Unlock()

	if !q.Equals(t.quat) {
		t.quat = q
		t.built = math.Mat4Zeros
	}
}

// Quat returns the quaternion component of this transformation.
func (t *Transform) Quat() math.Quat {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.quat.Equals(math.QuatZero) {
		if !t.rotation.Equals(math.Vec3Zero) {
			// We have euler rotation values we can convert
			hpr := t.rotation.Radians().XyzToHpr()
			return math.QuatFromHpr(hpr, math.CoordSysZUpRight)
		} else {
			return math.QuatIdentity
		}
	}
	return t.quat
}

// SetScale specifies the scale component of this transformation.
func (t *Transform) SetScale(s math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	if !s.Equals(t.scale) {
		t.scale = s
		t.built = math.Mat4Zeros
	}
}

// Scale returns the scale component of this transformation.
func (t *Transform) Scale() math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.scale.Equals(math.Vec3Zero) {
		return math.Vec3One
	}
	return t.scale
}

// SetShear specifies the shear component of this transformation.
func (t *Transform) SetShear(s math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	if !s.Equals(t.shear) {
		t.shear = s
		t.built = math.Mat4Zeros
	}
}

// Shear returns the shear component of this transformation.
func (t *Transform) Shear() math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()
	return t.shear
}

// Reset resets all the components of this transformation such that it is in
// it's original (i.e. identity matrix) state.
//
// Invoking this also removes an arbitrary transformation matrix from this
// transform, if it has one.
func (t *Transform) Reset() {
	t.access.RLock()
	defer t.access.RUnlock()

	t.arbitrary = false
	t.built = math.Mat4Zeros
	t.position = math.Vec3Zero
	t.rotation = math.Vec3Zero
	t.scale = math.Vec3Zero
	t.shear = math.Vec3Zero
	t.quat = math.QuatZero
}

// Copy returns a new 1:1 copy of this transformation.
func (t *Transform) Copy() *Transform {
	t.access.RLock()
	defer t.access.RUnlock()

	return &Transform{
		arbitrary: t.arbitrary,
		built:     t.built,
		position:  t.position,
		rotation:  t.rotation,
		scale:     t.scale,
		shear:     t.shear,
		quat:      t.quat,
	}
}
