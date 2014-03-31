// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"azul3d.org/v1/math"
)

// Transform represents a single 3D transformation.
type Transform struct {
	// The transformation matrix that will be applied to each mesh's vertices
	// by the renderer.
	//
	// Clients that do not intend to directly specify transformation matrix
	// will use the Build method to create one from this transform's rotation,
	// scale, shear, and position values.
	math.Mat4

	// UseQuat specifies if quaternion rotation should be used to build this
	// object's transformation matrix instead of the euler rotation stored in
	// the Rot field.
	UseQuat bool

	// Quat is the quaternion rotation of this object, it is only used if
	// UseQuat is set to true.
	Quat math.Quat

	// Euler rotation (in degrees) about the X, Y, and Z axis. Only used if
	// UseQuat is set to false (it is by default).
	Rot math.Vec3

	// Scale of the object about the X, Y, and Z axis (e.g. math.Vec3{2.0, 2.0,
	// 2.0} would make the object appear twice as large).
	Scale math.Vec3

	// Shear of the object about the X, Y, and Z axis (visually it looks like
	// the object is skewed).
	Shear math.Vec3

	// Position of the object about the X, Y, and Z axis. Since the gfx package
	// standardly uses a right-handed Z-up coordinate system, X is
	// left-to-right, Y is front-to-back (or 'in-and-out' of your monitor), and
	// Z is up-and-down.
	Pos math.Vec3
}

// Build builds a transformation matrix usign the Quat/Rot, Scale, Shear, and
// Pos values in the transform t. A new transform with t.Mat4 set to the built
// matrix is returned.
func (t Transform) Build() Transform {
	built := math.Mat4Identity

	// Apply rotation
	hpr := math.Vec3Zero
	if !t.UseQuat {
		// Use euler rotation.
		hpr = t.Rot.Radians().XyzToHpr()
	} else {
		// Use quaternion rotation.
		hpr = t.Quat.Hpr(math.CoordSysZUpRight)
	}

	// Compose upper 3x3 matrics using scale, shear, and HPR components.
	m := math.Mat3Compose(t.Scale, t.Shear, hpr, math.CoordSysZUpRight)
	built = built.SetUpperMat3(m)

	// Add in translation.
	trans := math.Mat4FromTranslation(t.Pos)
	built = built.Mul(trans)

	t.Mat4 = built
	return t
}

var (
	DefaultTransform = Transform{
		Mat4:  math.Mat4Identity,
		Scale: math.Vec3One,
	}
)
