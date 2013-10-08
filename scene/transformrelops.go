// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"code.google.com/p/azul3d/math"
)

// SetRelativePosVec3 sets the position component of this node's transformation
// to the specified position, as seen by (relative to) the other node.
//
// If the other node is nil, this function simply invokes n.SetPosVec3(pos) and
// returns.
func (n *Node) SetRelativePosVec3(other *Node, pos *math.Vec3) {
	if other == nil {
		n.SetPosVec3(pos)
		return
	}

	t := n.Transform()
	relative := n.RelativeTransform(other)
	relative.SetPos(pos)

	if t.HasComponents() {
		// Get the old rotation, scale, and shear values (fixes floating point
		// errors).
		origRot := n.Transform().Rot()
		origScale := n.Transform().Scale()
		origShear := n.Transform().Shear()

		// Apply relative transformation
		n.SetRelativeTransform(other, relative)

		// Get updated transformation
		t = n.Transform()

		// Restore old rotation, scale, and shear values (fixes floating point
		// errors).
		t.SetRot(origRot)
		t.SetScale(origScale)
		t.SetShear(origShear)

	} else {
		// No transformation components specified, we don't need to save them.
		n.SetRelativeTransform(other, relative)
	}
}

// RelativePosVec3 returns the position component of this node's transformation
// as seen by (relative to) the other node.
//
// If the other node is nil, this function simply returns n.PosVec3().
func (n *Node) RelativePosVec3(other *Node) *math.Vec3 {
	if other == nil {
		return n.PosVec3()
	}

	relative := n.RelativeTransform(other)
	return relative.Pos()
}

// SetRelativeRotVec3 sets the rotation component of this node's transformation
// to the specified rotation, as seen by (relative to) the other node.
//
// If the other node is nil, this function simply invokes n.SetRotVec3(rot) and
// returns.
func (n *Node) SetRelativeRotVec3(other *Node, rot *math.Vec3) {
	if other == nil {
		n.SetRotVec3(rot)
		return
	}

	t := n.Transform()
	relative := n.RelativeTransform(other)
	relative.SetRot(rot)

	if t.HasComponents() {
		// Get the old position, scale, and shear values (fixes floating point
		// errors).
		origPos := n.Transform().Pos()
		origScale := n.Transform().Scale()
		origShear := n.Transform().Shear()

		// Apply relative transformation
		n.SetRelativeTransform(other, relative)

		// Get updated transformation
		t = n.Transform()

		// Restore old position, scale, and shear values (fixes floating point
		// errors).
		t.SetPos(origPos)
		t.SetScale(origScale)
		t.SetShear(origShear)

	} else {
		// No transformation components specified, we don't need to save them.
		n.SetRelativeTransform(other, relative)
	}
}

// RelativeRotVec3 returns the rotation component of this node's transformation
// as seen by (relative to) the other node.
//
// If the other node is nil, this function simply returns n.RotVec3().
func (n *Node) RelativeRotVec3(other *Node) *math.Vec3 {
	if other == nil {
		return n.RotVec3()
	}

	relative := n.RelativeTransform(other)
	return relative.Rot()
}

// SetRelativeScaleVec3 sets the scale component of this node's transformation
// to the specified scale, as seen by (relative to) the other node.
//
// If the other node is nil, this function simply invokes n.SetScaleVec3(pos) and
// returns.
func (n *Node) SetRelativeScaleVec3(other *Node, scale *math.Vec3) {
	if other == nil {
		n.SetScaleVec3(scale)
		return
	}

	t := n.Transform()
	relative := n.RelativeTransform(other)
	relative.SetScale(scale)

	if t.HasComponents() {
		// Get the old position, rotation, and shear values (fixes floating
		// point errors).
		origPos := n.Transform().Pos()
		origRot := n.Transform().Rot()
		origShear := n.Transform().Shear()

		// Apply relative transformation
		n.SetRelativeTransform(other, relative)

		// Get updated transformation
		t = n.Transform()

		// Restore old position, rotation, and shear values (fixes floating
		// point errors).
		t.SetPos(origPos)
		t.SetRot(origRot)
		t.SetShear(origShear)

	} else {
		// No transformation components specified, we don't need to save them.
		n.SetRelativeTransform(other, relative)
	}
}

// RelativeScaleVec3 returns the scale component of this node's transformation
// as seen by (relative to) the other node.
//
// If the other node is nil, this function simply returns n.ScaleVec3().
func (n *Node) RelativeScaleVec3(other *Node) *math.Vec3 {
	if other == nil {
		return n.ScaleVec3()
	}

	relative := n.RelativeTransform(other)
	return relative.Scale()
}

// SetRelativeShearVec3 sets the shear component of this node's transformation
// to the specified shear, as seen by (relative to) the other node.
//
// If the other node is nil, this function simply invokes n.SetShearVec3(shear) and
// returns.
func (n *Node) SetRelativeShearVec3(other *Node, shear *math.Vec3) {
	if other == nil {
		n.SetShearVec3(shear)
		return
	}

	t := n.Transform()
	relative := n.RelativeTransform(other)
	relative.SetShear(shear)

	if t.HasComponents() {
		// Get the old position, rotation, and scale values (fixes floating
		// point errors).
		origPos := n.Transform().Pos()
		origRot := n.Transform().Rot()
		origScale := n.Transform().Scale()

		// Apply relative transformation
		n.SetRelativeTransform(other, relative)

		// Get updated transformation
		t = n.Transform()

		// Restore old position, rotation, and scale values (fixes floating
		// point errors).
		t.SetPos(origPos)
		t.SetRot(origRot)
		t.SetScale(origScale)

	} else {
		// No transformation components specified, we don't need to save them.
		n.SetRelativeTransform(other, relative)
	}
}

// RelativeShearVec3 returns the shear component of this node's transformation
// as seen by (relative to) the other node.
//
// If the other node is nil, this function simply returns n.ShearVec3().
func (n *Node) RelativeShearVec3(other *Node) *math.Vec3 {
	if other == nil {
		return n.ShearVec3()
	}

	relative := n.RelativeTransform(other)
	return relative.Shear()
}

// SetRelativePos is short hand for:
//
//  n.SetRelativePosVec3(other, math.Vector3(x, y, z))
//
func (n *Node) SetRelativePos(other *Node, x, y, z math.Real) {
	n.SetRelativePosVec3(other, math.Vector3(x, y, z))
}

// RelativePos is short hand for:
//
//  p := n.RelativePosVec3(other)
//  return p.X, p.Y, p.Z
//
func (n *Node) RelativePos(other *Node) (x, y, z math.Real) {
	p := n.RelativePosVec3(other)
	return p.X, p.Y, p.Z
}

// SetRelativeRot is short hand for:
//
//  n.SetRelativeRotVec3(other, math.Vector3(rx, ry, rz))
//
func (n *Node) SetRelativeRot(other *Node, rx, ry, rz math.Real) {
	n.SetRelativeRotVec3(other, math.Vector3(rx, ry, rz))
}

// RelativeRot is short hand for:
//
//  r := n.RelativeRotVec3(other)
//  return r.X, r.Y, r.Z
//
func (n *Node) RelativeRot(other *Node) (rx, ry, rz math.Real) {
	r := n.RelativeRotVec3(other)
	return r.X, r.Y, r.Z
}

// SetRelativeScale is short hand for:
//
//  n.SetRelativeScaleVec3(other, math.Vector3(sx, sy, sz))
//
func (n *Node) SetRelativeScale(other *Node, sx, sy, sz math.Real) {
	n.SetRelativeScaleVec3(other, math.Vector3(sx, sy, sz))
}

// RelativeScale is short hand for:
//
//  s := n.RelativeScaleVec3(other)
//  return s.X, s.Y, s.Z
//
func (n *Node) RelativeScale(other *Node) (sx, sy, sz math.Real) {
	s := n.RelativeScaleVec3(other)
	return s.X, s.Y, s.Z
}

// SetRelativeShear is short hand for:
//
//  n.SetRelativeShearVec3(other, math.Vector3(shx, shy, shz))
//
func (n *Node) SetRelativeShear(other *Node, shx, shy, shz math.Real) {
	n.SetRelativeShearVec3(other, math.Vector3(shx, shy, shz))
}

// RelativeShear is short hand for:
//
//  sh := n.RelativeShearVec3(other)
//  return sh.X, sh.Y, sh.Z
//
func (n *Node) RelativeShear(other *Node) (shx, shy, shz math.Real) {
	sh := n.RelativeShearVec3(other)
	return sh.X, sh.Y, sh.Z
}
