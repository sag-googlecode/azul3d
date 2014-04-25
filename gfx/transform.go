// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"azul3d.org/v1/math"
	"sync"
)

// Transformable represents a generic interface to any object that can return a
// transformation matrix.
type Transformable interface {
	Mat4() math.Mat4
}

type TransformSpace uint8

const (
	// World space is the top-most (world/global) space. A transform whose
	// parent is nil implies that the parent is actually the world itself. All
	// object's transform's are converted to world space for display.
	WorldSpace TransformSpace = iota

	// Local space is the space that the transform itself defines. It is
	// positioned, rotated, scaled, etc according to the components of the
	// transform.
	//
	// For example, the vertices of an object are in local space (i.e.
	// converting a vertex from local to world space means that the transform's
	// position, rotation, etc are included).
	LocalSpace
)

// Transform represents a generic 3D transformation. It can be safely used from
// multiple goroutines concurrently. It is built from various components such
// as position, scale, and shear values and may use euler or quaternion
// rotation. It supports a hierarchial tree system of transforms to create
// complex transformations.
//
// Many 3D modeling packages and game engines use the terms 'world space' and
// 'local space' but not consistently. We define them explicitly as follows:
//  World Space:
//      The top-most 'world' space. A transform whose parent is nil explicitly
//      means the parent is the world. e.g. Each vertex of an object is
//      converted to world space for display.
//
//  Local space:
//      The space that the local transform defines. The space is positioned,
//      rotated, scaled, etc relative to the parent. e.g. Each vertex of an
//      object is in local space and is then converted to world space for
//      display.
//
// World space serves as a common factor across all transforms. Positions,
// rotations, etc can all be converted from one transform's local space to
// world space and back (thus allowing for relative movement, rotation, etc).
type Transform struct {
	access sync.RWMutex

	// The parent transform, or nil if there is none.
	parent     *Transform
	lastParent *Transform

	// A pointer to the built (i.e. cached) transformation matrix or nil if a
	// rebuild is required.
	built *math.Mat4

	// Pointers to the matrices describing local-to-world and world-to-local
	// space conversions.
	localToWorld, worldToLocal *math.Mat4

	// A pointer to a quaternion rotation, or nil if euler rotation is in use.
	quat *math.Quat

	// The position, rotation, scaling, and shearing components.
	pos, rot, scale, shear math.Vec3
}

// Equals tells if the two transforms are equal.
func (t *Transform) Equals(other *Transform) bool {
	t.access.RLock()
	other.access.RLock()

	// Compare parent pointers.
	if t.parent != other.parent {
		goto fail
	}

	// Two-step quaternion comparison.
	if (t.quat != nil) != (other.quat != nil) {
		goto fail
	}
	if t.quat != nil && !(*t.quat).Equals(*other.quat) {
		goto fail
	}

	// Compare position, rotation, scale, and shear.
	if !t.pos.Equals(other.pos) {
		goto fail
	}
	if !t.rot.Equals(other.rot) {
		goto fail
	}
	if !t.scale.Equals(other.scale) {
		goto fail
	}
	if !t.shear.Equals(other.shear) {
		goto fail
	}

	t.access.RUnlock()
	other.access.RUnlock()
	return true

fail:
	t.access.RUnlock()
	other.access.RUnlock()
	return false
}

// build builds and stores the transformation matrix from the components of
// this transform.
func (t *Transform) build() {
	if t.built != nil && (t.lastParent != nil && t.parent != nil && t.lastParent.Equals(t.parent)) {
		// No update is required.
		return
	}
	if t.parent != nil {
		t.lastParent = t.parent.Copy()
	}

	// Apply rotation
	var hpr math.Vec3
	if t.quat != nil {
		// Use quaternion rotation.
		hpr = (*t.quat).Hpr(math.CoordSysZUpRight)
	} else {
		// Use euler rotation.
		hpr = t.rot.XyzToHpr().Radians()
	}

	// Compose upper 3x3 matrics using scale, shear, and HPR components.
	scaleShearHpr := math.Mat3Compose(t.scale, t.shear, hpr, math.CoordSysZUpRight)

	// Build this space's transformation matrix.
	built := math.Mat4Identity.SetUpperMat3(scaleShearHpr)
	built = built.SetTranslation(t.pos)
	t.built = &built

	// Build the local-to-world transformation matrix.
	ltw := built
	if t.parent != nil {
		ltw = ltw.Mul(t.parent.underLocalToWorld())
	}
	t.localToWorld = &ltw

	// Build the world-to-local transformation matrix.
	wtl, _ := built.Inverse()
	if t.parent != nil {
		parent := t.parent.worldToUnderLocal()
		wtl = wtl.Mul(parent)
	}
	t.worldToLocal = &wtl
}

// Implements Transformable interface by simply returning the local-to-world
// matrix.
func (t *Transform) Mat4() math.Mat4 {
	return t.LocalToWorld()
}

// LocalMat4 returns a matrix describing the local transformation. It is the
// matrix that is built out of the components of this transform.
func (t *Transform) LocalMat4() math.Mat4 {
	t.access.Lock()
	t.build()
	l := *t.built
	t.access.Unlock()
	return l
}

func (t *Transform) underLocalToWorld() math.Mat4 {
	t.access.Lock()
	t.build()
	ltw := *t.localToWorld
	t.access.Unlock()
	return ltw
}

// LocalToWorld returns a matrix which converts from this transform's local
// space into world space.
func (t *Transform) LocalToWorld() math.Mat4 {
	t.access.Lock()
	t.build()
	ltw := *t.localToWorld
	if t.parent != nil {
		// Undo local transform.
		localInv, _ := (*t.built).Inverse()
		ltw = localInv.Mul(ltw)
	}
	t.access.Unlock()
	return ltw
}

func (t *Transform) worldToUnderLocal() math.Mat4 {
	t.access.Lock()
	t.build()
	wtl := *t.worldToLocal
	t.access.Unlock()
	return wtl
}

// WorldToLocal returns a matrix which converts from world space to this
// transform's local space.
func (t *Transform) WorldToLocal() math.Mat4 {
	t.access.Lock()
	t.build()
	wtl := *t.worldToLocal
	if t.parent != nil {
		// Undo local transform.
		wtl = wtl.Mul(*t.built)
	}
	t.access.Unlock()
	return wtl
}

// SetParent sets a parent transform for this transform to effectively inherit
// from. This allows creating complex hierarchies of transformations.
//
// e.g. setting the parent of a camera's transform to the player's transform
// makes it such that the camera follows the player.
func (t *Transform) SetParent(p *Transform) {
	t.access.Lock()
	if t.parent != p {
		t.built = nil
		t.parent = p
	}
	t.access.Unlock()
}

// Parent returns the parent of this transform, as previously set.
func (t *Transform) Parent() *Transform {
	t.access.RLock()
	p := t.parent
	t.access.RUnlock()
	return p
}

// SetQuat sets the quaternion rotation of this transform.
//
// The last call to either SetQuat or SetRot is what effictively determines
// whether quaternion or euler rotation will be used by this transform.
func (t *Transform) SetQuat(q math.Quat) {
	t.access.Lock()
	if (*t.quat) != q {
		t.built = nil
		t.quat = &q
	}
	t.access.Unlock()
}

// Quat returns the quaternion rotation of this transform. If this transform is
// instead using euler rotation (see IsQuat) then a quaternion is created from
// the euler rotation of this transform and returned.
//
// The last call to either SetQuat or SetRot is what effictively determines
// whether quaternion or euler rotation will be used by this transform.
func (t *Transform) Quat() math.Quat {
	var q math.Quat
	t.access.RLock()
	if t.quat != nil {
		q = *t.quat
	} else {
		// Convert euler rotation to quaternion.
		q = math.QuatFromHpr(t.rot.XyzToHpr().Radians(), math.CoordSysZUpRight)
	}
	t.access.RUnlock()
	return q
}

// IsQuat tells if this transform is currently utilizing quaternion or euler
// rotation.
//
// The last call to either SetQuat or SetRot is what effictively determines
// whether quaternion or euler rotation will be used by this transform.
func (t *Transform) IsQuat() bool {
	t.access.RLock()
	isQuat := t.quat != nil
	t.access.RUnlock()
	return isQuat
}

// SetRot sets the euler rotation of this transform in degrees about their
// respective axis (e.g. if r.X == 45 then it is 45 degrees around the X
// axis).
//
// The last call to either SetQuat or SetRot is what effictively determines
// whether quaternion or euler rotation will be used by this transform.
func (t *Transform) SetRot(r math.Vec3) {
	t.access.Lock()
	if t.rot != r {
		t.built = nil
		t.quat = nil
		t.rot = r
	}
	t.access.Unlock()
}

// Rot returns the euler rotation of this transform. If this transform is
// instead using quaternion (see IsQuat) rotation then it is converted to euler
// rotation and returned.
//
// The last call to either SetQuat or SetRot is what effictively determines
// whether quaternion or euler rotation will be used by this transform.
func (t *Transform) Rot() math.Vec3 {
	var r math.Vec3
	t.access.RLock()
	if t.quat == nil {
		r = t.rot
	} else {
		// Convert quaternion rotation to euler rotation.
		r = (*t.quat).Hpr(math.CoordSysZUpRight).HprToXyz().Degrees()
	}
	t.access.RUnlock()
	return r
}

// SetPos sets the local position of this transform.
func (t *Transform) SetPos(p math.Vec3) {
	t.access.Lock()
	if t.pos != p {
		t.built = nil
		t.pos = p
	}
	t.access.Unlock()
}

// Pos returns the local position of this transform.
func (t *Transform) Pos() math.Vec3 {
	t.access.RLock()
	p := t.pos
	t.access.RUnlock()
	return p
}

// SetScale sets the local scale of this transform (e.g. a scale of
// math.Vec3{2, 1.5, 1} would make an object appear twice as large on the local
// X axis, one and a half times larger on the local Y axis, and would not scale
// on the local Z axis at all).
func (t *Transform) SetScale(s math.Vec3) {
	t.access.Lock()
	if t.scale != s {
		t.built = nil
		t.scale = s
	}
	t.access.Unlock()
}

// Scale returns the local scacle of this transform.
func (t *Transform) Scale() math.Vec3 {
	t.access.RLock()
	s := t.scale
	t.access.RUnlock()
	return s
}

// SetShear sets the local shear of this transform.
func (t *Transform) SetShear(s math.Vec3) {
	t.access.Lock()
	if t.shear != s {
		t.built = nil
		t.shear = s
	}
	t.access.Unlock()
}

// Shear returns the local shear of this transform.
func (t *Transform) Shear() math.Vec3 {
	t.access.RLock()
	s := t.shear
	t.access.RUnlock()
	return s
}

// Reset sets all of the values of this transform to the default ones.
func (t *Transform) Reset() {
	t.access.Lock()
	t.parent = nil
	t.built = nil
	t.localToWorld = nil
	t.worldToLocal = nil
	t.quat = nil
	t.pos = math.Vec3Zero
	t.rot = math.Vec3Zero
	t.scale = math.Vec3One
	t.shear = math.Vec3Zero
	t.access.Unlock()
}

// Copy returns a new transform with all of it's values set equal to t (i.e. a
// copy of this transform).
func (t *Transform) Copy() *Transform {
	t.access.RLock()
	cpy := &Transform{
		parent: t.parent,
		pos:    t.pos,
		rot:    t.rot,
		scale:  t.scale,
		shear:  t.shear,
	}
	if t.built != nil {
		builtCpy := *t.built
		cpy.built = &builtCpy
	}
	if t.localToWorld != nil {
		ltwCpy := *t.localToWorld
		cpy.localToWorld = &ltwCpy
	}
	if t.worldToLocal != nil {
		wtlCpy := *t.worldToLocal
		cpy.worldToLocal = &wtlCpy
	}
	if t.quat != nil {
		quatCpy := *t.quat
		cpy.quat = &quatCpy
	}
	t.access.RUnlock()
	return cpy
}

// NewTransform returns a new *Transform with the default values (a uniform
// scale of one).
func NewTransform() *Transform {
	return &Transform{
		scale: math.Vec3One,
	}
}

// PosToWorld converts the given position in this transform's local space to
// world space.
func (t *Transform) PosToWorld(p math.Vec3) math.Vec3 {
	return p.TransformMat4(t.LocalToWorld())
}

// PosToLocal converts the given position in world space to this transform's
// local space.
func (t *Transform) PosToLocal(p math.Vec3) math.Vec3 {
	return p.TransformMat4(t.WorldToLocal())
}

func (t *Transform) rotConv(r math.Vec3, m math.Mat4) math.Vec3 {
	q := math.QuatFromHpr(r.XyzToHpr().Radians(), math.CoordSysZUpRight)
	m = q.ExtractToMat4().Mul(m)
	q = math.QuatFromMat3(m.UpperMat3())
	return q.Hpr(math.CoordSysZUpRight).HprToXyz().Degrees()
}

// RotToWorld converts the given euler rotation in this transform's local space
// to world space.
func (t *Transform) RotToWorld(r math.Vec3) math.Vec3 {
	return t.rotConv(r, t.LocalToWorld())
}

// RotToLocal converts the given euler rotation in world space to this
// transform's local space.
func (t *Transform) RotToLocal(r math.Vec3) math.Vec3 {
	return t.rotConv(r, t.WorldToLocal())
}
