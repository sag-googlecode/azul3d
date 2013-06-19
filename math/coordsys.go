// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package math

import (
	"fmt"
)

// CoordSys represents an specific coordinate system.
type CoordSys uint8

// RightHanded tells weather this coordinate system is right-handed.
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) RightHanded() bool {
	switch cs {
	case CoordSysZUpRight:
		return true
	case CoordSysYUpRight:
		return true

	case CoordSysZUpLeft:
		return false
	case CoordSysYUpLeft:
		return false
	}
	panic(fmt.Sprintf("RightHanded(): Invalid coordinate system %d", cs))
}

// LeftHanded is short for:
//
//  !cs.RightHanded()
func (cs CoordSys) LeftHanded() bool {
	return !cs.RightHanded()
}

const (
	// Invalid coordinate system
	CoordSysInvalid CoordSys = iota

	// Z up axis, right handed coordinate system
	CoordSysZUpRight

	// Y up axis, right handed coordinate system
	CoordSysYUpRight

	// Z up axis, left handed coordinate system
	CoordSysZUpLeft

	// Y up axis, left handed coordinate system
	CoordSysYUpLeft
)

// Up returns the up vector for the given coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) Up() *Vec3 {
	switch cs {
	case CoordSysZUpRight:
		return &Vec3{0, 0, 1}
	case CoordSysZUpLeft:
		return &Vec3{0, 0, 1}

	case CoordSysYUpRight:
		return &Vec3{0, 1, 0}
	case CoordSysYUpLeft:
		return &Vec3{0, 1, 0}
	}
	panic(fmt.Sprintf("Up(): Invalid coordinate system %d", cs))
}

// Right returns the right vector for the given coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) Right() *Vec3 {
	return &Vec3{1, 0, 0}
}

// Forward returns the forward vector for the given coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) Forward() *Vec3 {
	switch cs {
	case CoordSysZUpRight:
		return &Vec3{0, 1, 0}

	case CoordSysZUpLeft:
		return &Vec3{0, -1, 0}

	case CoordSysYUpRight:
		return &Vec3{0, 0, -1}

	case CoordSysYUpLeft:
		return &Vec3{0, 0, 1}
	}
	panic(fmt.Sprintf("Forward(): Invalid coordinate system %d", cs))
}

// Down returns the down vector for the given coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) Down() *Vec3 {
	v := cs.Up()
	v.Negate()
	return v
}

// Left returns the left vector for the given coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) Left() *Vec3 {
	v := cs.Right()
	v.Negate()
	return v
}

// Back returns the back vector for the given coordinate system
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (cs CoordSys) Back() *Vec3 {
	v := cs.Forward()
	v.Negate()
	return v
}

// RightFrontUp returns an vector that is described by it's right, forward, and up components in
// whatever way the specified coordinate system represents that vector.
func (cs CoordSys) RightFrontUp(right, forward, up Real) *Vec3 {
	var vz, vy Real

	switch cs {
	case CoordSysYUpRight:
		vz = -forward
		vy = up

	case CoordSysYUpLeft:
		vz = forward
		vy = up

	case CoordSysZUpRight:
		vy = forward
		vz = up

	case CoordSysZUpLeft:
		vy = -forward
		vz = up

	default:
		panic(fmt.Sprintf("Forward(): Invalid coordinate system %d", cs))
	}
	return &Vec3{right, vy, vz}
}

// ConvertMat3 returns a matrix that transforms from the indicated coordinate system, to the
// coordinate system specified.
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (from CoordSys) ConvertMat3(to CoordSys) *Mat3 {
	switch from {
	case CoordSysZUpLeft:
		switch to {
		case CoordSysZUpLeft:
			return Mat3Identity
		case CoordSysYUpLeft:
			return Mat3ZToYUp
		case CoordSysZUpRight:
			return Mat3FlipY
		case CoordSysYUpRight:
			return Mat3LZToRY
		}

	case CoordSysYUpLeft:
		switch to {
		case CoordSysZUpLeft:
			return Mat3YToZUp
		case CoordSysYUpLeft:
			return Mat3Identity
		case CoordSysZUpRight:
			return Mat3LYToRZ
		case CoordSysYUpRight:
			return Mat3FlipZ
		}

	case CoordSysZUpRight:
		switch to {
		case CoordSysZUpLeft:
			return Mat3FlipY
		case CoordSysYUpLeft:
			return Mat3LZToRY
		case CoordSysZUpRight:
			return Mat3Identity
		case CoordSysYUpRight:
			return Mat3ZToYUp
		}

	case CoordSysYUpRight:
		switch to {
		case CoordSysZUpLeft:
			return Mat3LYToRZ
		case CoordSysYUpLeft:
			return Mat3FlipZ
		case CoordSysZUpRight:
			return Mat3YToZUp
		case CoordSysYUpRight:
			return Mat3Identity
		}
	}
	panic(fmt.Sprintf("ConvertMat3(): Invalid coordinate system %d / %d", from, to))
}

// ConvertMat4 returns a matrix that transforms from the indicated coordinate system, to the
// coordinate system specified.
//
// An panic occurs if the coordinate system is not one of the valid coordinate system constants
// defined in this package.
func (from CoordSys) ConvertMat4(to CoordSys) *Mat4 {
	switch from {
	case CoordSysZUpLeft:
		switch to {
		case CoordSysZUpLeft:
			return Mat4Identity
		case CoordSysYUpLeft:
			return Mat4ZToYUp
		case CoordSysZUpRight:
			return Mat4FlipY
		case CoordSysYUpRight:
			return Mat4LZToRY
		}

	case CoordSysYUpLeft:
		switch to {
		case CoordSysZUpLeft:
			return Mat4YToZUp
		case CoordSysYUpLeft:
			return Mat4Identity
		case CoordSysZUpRight:
			return Mat4LYToRZ
		case CoordSysYUpRight:
			return Mat4FlipZ
		}

	case CoordSysZUpRight:
		switch to {
		case CoordSysZUpLeft:
			return Mat4FlipY
		case CoordSysYUpLeft:
			return Mat4LZToRY
		case CoordSysZUpRight:
			return Mat4Identity
		case CoordSysYUpRight:
			return Mat4ZToYUp
		}

	case CoordSysYUpRight:
		switch to {
		case CoordSysZUpLeft:
			return Mat4LYToRZ
		case CoordSysYUpLeft:
			return Mat4FlipZ
		case CoordSysZUpRight:
			return Mat4YToZUp
		case CoordSysYUpRight:
			return Mat4Identity
		}
	}
	panic(fmt.Sprintf("ConvertMat4(): Invalid coordinate system %d / %d", from, to))
}
