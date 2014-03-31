// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"azul3d.org/v1/math"
)

// SetPosVec3 sets the position component of this node's transformation to the
// specified position point vector.
func (n *Node) SetPosVec3(pos math.Vec3) {
	n.access.RLock()
	defer n.access.RUnlock()

	n.transform.SetPos(pos)
}

// PosVec3 returns the position component of this node's transformation as an
// point vector.
func (n *Node) PosVec3() math.Vec3 {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.transform.Pos()
}

// SetRotVec3 sets the rotation component of this node's transformation to the
// specified rotation vector (specified in degrees of rotation about each
// respective axis).
func (n *Node) SetRotVec3(rot math.Vec3) {
	n.access.RLock()
	defer n.access.RUnlock()

	n.transform.SetRot(rot)
}

// RotVec3 returns the rotation component of this node's transformation as an
// rotation vector (specified in degrees of rotation about each respective
// axis).
func (n *Node) RotVec3() math.Vec3 {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.transform.Rot()
}

// SetScaleVec3 sets the scale component of this node's transformation to the
// specified scale vector (specified as an scaling multiplier about each
// respective axis).
func (n *Node) SetScaleVec3(scale math.Vec3) {
	n.access.RLock()
	defer n.access.RUnlock()

	n.transform.SetScale(scale)
}

// ScaleVec3 returns the scale component of this node's transformation as an
// scale vector (specified as an scaling multiplier about each respective
// axis).
func (n *Node) ScaleVec3() math.Vec3 {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.transform.Scale()
}

// SetShearVec3 sets the shear component of this node's transformation to the
// specified shear vector (specified as an shearing value along each respective
// axis).
func (n *Node) SetShearVec3(shear math.Vec3) {
	n.access.RLock()
	defer n.access.RUnlock()

	n.transform.SetShear(shear)
}

// ShearVec3 returns the shear component of this node's transformation as an
// shear vector (specified as an shearing value along each respective axis).
func (n *Node) ShearVec3() math.Vec3 {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.transform.Shear()
}

// SetPos is short hand for:
//
//  n.SetPosVec3(math.Vec3{x, y, z))
//
func (n *Node) SetPos(x, y, z float64) {
	n.SetPosVec3(math.Vec3{x, y, z})
}

// Pos is short hand for:
//
//  p := n.PosVec3()
//  return p.X, p.Y, p.Z
//
func (n *Node) Pos() (x, y, z float64) {
	p := n.PosVec3()
	return p.X, p.Y, p.Z
}

// SetRot is short hand for:
//
//  n.SetRotVec3(math.Vec3{rx, ry, rz))
//
func (n *Node) SetRot(rx, ry, rz float64) {
	n.SetRotVec3(math.Vec3{rx, ry, rz})
}

// Rot is short hand for:
//
//  r := n.RotVec3()
//  return r.X, r.Y, r.Z
//
func (n *Node) Rot() (rx, ry, rz float64) {
	r := n.RotVec3()
	return r.X, r.Y, r.Z
}

// SetScale is short hand for:
//
//  n.SetScaleVec3(math.Vec3{sx, sy, sz))
//
func (n *Node) SetScale(sx, sy, sz float64) {
	n.SetScaleVec3(math.Vec3{sx, sy, sz})
}

// Scale is short hand for:
//
//  s := n.ScaleVec3()
//  return s.X, s.Y, s.Z
//
func (n *Node) Scale() (sx, sy, sz float64) {
	s := n.ScaleVec3()
	return s.X, s.Y, s.Z
}

// SetShear is short hand for:
//
//  n.SetShearVec3(math.Vec3{shx, shy, shz))
//
func (n *Node) SetShear(shx, shy, shz float64) {
	n.SetShearVec3(math.Vec3{shx, shy, shz})
}

// Shear is short hand for:
//
//  sh := n.ShearVec3()
//  return sh.X, sh.Y, sh.Z
//
func (n *Node) Shear() (shx, shy, shz float64) {
	sh := n.ShearVec3()
	return sh.X, sh.Y, sh.Z
}
