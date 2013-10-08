// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"code.google.com/p/azul3d/math"
	"sync"
)

type Transform struct {
	access sync.RWMutex

	arbitrary                        bool
	built                            *math.Mat4
	position, rotation, scale, shear *math.Vec3
	quat                             *math.Quat
}

func (t *Transform) hasComponents() bool {
	if t.arbitrary {
		return false
	}

	return t.position != nil || t.rotation != nil || t.scale != nil || t.shear != nil || t.quat != nil
}

func (t *Transform) HasComponents() bool {
	t.access.RLock()
	defer t.access.RUnlock()

	return t.hasComponents()
}

func (t *Transform) Compose(other *Transform) *Transform {
	composed := new(Transform)

	if true { //t.ArbitraryMat4() != nil || other.ArbitraryMat4() != nil {
		// Combine using matrices
		newMat := other.Mat4().Mul(t.Mat4())
		composed.SetMat4(newMat)

	} else {
		// Compose using components
		pos := t.Pos()
		quat := t.Quat()
		scale := t.Scale()

		pos = pos.Add(quat.TransformVec3(other.Pos()).Mul(scale))
		quat = other.Quat().Mul(quat)
		newScale := other.Scale().Mul(scale)

		composed.SetPos(pos)
		composed.SetQuat(quat)
		composed.SetScale(newScale)
	}

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
	t.built = math.Mat4Identity.Copy()

	// If we don't have components to build with, we can just leave it as the
	// identity matrix above.
	if !t.hasComponents() {
		return
	}

	scale := math.Vec3One
	if t.scale != nil {
		scale = t.scale
	}

	shear := math.Vec3Zero
	if t.shear != nil {
		// We expect shear in an different format from the math package. So
		// handle this as we want here.
		shear = t.shear
	}

	// Apply translation
	pos := math.Vec3Zero
	if t.position != nil {
		pos = t.position
	}

	// Apply rotation
	hpr := math.Vec3Zero
	if t.rotation != nil {
		// We have euler rotation
		hpr = t.rotation.Radians().XyzToHpr()

	} else if t.quat != nil {
		// We have quaternion rotation
		hpr = t.quat.Hpr(math.CoordSysZUpRight)
	}

	m := math.Mat3Identity.Copy()
	m.Compose(scale, shear, hpr, math.CoordSysZUpRight)
	t.built.SetUpperMat3(m)

	trans := new(math.Mat4)
	trans.SetTranslation(pos)
	t.built = t.built.Mul(trans)
}

// Note: this function does not make an copy of the matrix m, you should pass in one.
func (t *Transform) SetArbitraryMat4(m *math.Mat4) {
	t.access.Lock()
	defer t.access.Unlock()

	if m == nil {
		t.arbitrary = false
		t.built = nil
	} else {
		t.arbitrary = true
		t.built = m
	}
}

func (t *Transform) ArbitraryMat4() *math.Mat4 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.arbitrary {
		return t.built
	}
	return nil
}

func (t *Transform) SetMat4(m *math.Mat4) {
	t.access.Lock()
	defer t.access.Unlock()

	t.position = m.Translation()
	t.scale, t.shear, t.rotation = m.UpperMat3().Decompose(math.CoordSysZUpRight)

	// Decompose returns HPR, we expect RX/RY/RZ, and degrees.
	t.rotation = t.rotation.Degrees().HprToXyz()
}

// Note: this function does not return an copy of the matrix m, you should copy it instead.
func (t *Transform) Mat4() *math.Mat4 {
	t.access.Lock()
	defer t.access.Unlock()

	if t.built == nil {
		t.build()
	}
	return t.built
}

func (t *Transform) SetPos(p *math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	t.position = p.Copy()
	t.built = nil
}

func (t *Transform) Pos() *math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.position == nil {
		return math.Vec3Zero.Copy()
	}
	return t.position
}

func (t *Transform) SetRot(r *math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	t.rotation = r.Copy()
	t.built = nil
}

func (t *Transform) Rot() *math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.rotation == nil {
		if t.quat != nil {
			hpr := t.quat.Hpr(math.CoordSysZUpRight)
			return hpr.Degrees().HprToXyz()
		} else {
			return math.Vec3Zero.Copy()
		}
	}
	return t.rotation
}

func (t *Transform) SetQuat(q *math.Quat) {
	t.access.Lock()
	defer t.access.Unlock()

	t.quat = q.Copy()
	t.built = nil
}

func (t *Transform) Quat() *math.Quat {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.quat == nil {
		if t.rotation != nil {
			// We have euler rotation values we can convert
			hpr := t.rotation.Radians().XyzToHpr()

			quat := new(math.Quat)
			quat.SetHpr(hpr, math.CoordSysZUpRight)
			return quat

		} else {
			return math.QuatIdentity.Copy()
		}
	}
	return t.quat.Copy()
}

func (t *Transform) SetScale(s *math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	t.scale = s.Copy()
	t.built = nil
}

func (t *Transform) Scale() *math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.scale == nil {
		return math.Vec3One.Copy()
	}
	return t.scale.Copy()
}

func (t *Transform) SetShear(s *math.Vec3) {
	t.access.Lock()
	defer t.access.Unlock()

	t.shear = s.Copy()
	t.built = nil
}

func (t *Transform) Shear() *math.Vec3 {
	t.access.RLock()
	defer t.access.RUnlock()

	if t.shear == nil {
		return math.Vec3Zero.Copy()
	}
	return t.shear.Copy()
}

func (t *Transform) Reset() {
	t.access.RLock()
	defer t.access.RUnlock()

	t.arbitrary = false
	t.built = nil
	t.position = nil
	t.rotation = nil
	t.scale = nil
	t.shear = nil
	t.quat = nil
}

func (t *Transform) Copy() *Transform {
	t.access.RLock()
	defer t.access.RUnlock()

	c := new(Transform)

	if t.arbitrary {
		c.arbitrary = true
	}

	if t.built != nil {
		c.built = t.built.Copy()
	}

	if t.position != nil {
		c.position = t.position.Copy()
	}

	if t.rotation != nil {
		c.rotation = t.rotation.Copy()
	}

	if t.scale != nil {
		c.scale = t.scale.Copy()
	}

	if t.shear != nil {
		c.shear = t.shear.Copy()
	}

	if t.quat != nil {
		c.quat = t.quat.Copy()
	}

	return c
}
