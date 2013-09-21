// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package util

import (
	"code.google.com/p/azul3d/math"
	"sync"
)

// Lens represents an camera node's lens. There are perspective and orthographic lenses which suit
// most purposes.
//
// For more advanced lenses, you can craft them yourself.
type Lens struct {
	sync.RWMutex

	projection *math.Mat4
	near, far  math.Real

	perspectiveFovY, perspectiveAspect                   math.Real
	orthoLeft, orthoRight, orthoBottom, orthoTop         math.Real
	frustumLeft, frustumRight, frustumBottom, frustumTop math.Real
}

// Returns an 2D point on this lens' frustum given an 3D point in the view frustum space.
//
// The Y (depth) coordinate will be set to an value in the range of -1 to 1. Where an value of 1
// means near, and -1 means far.
//
// Returns ok=true if the 3D point is within the viewing frustum, or ok=false if it is outside the
// viewing frustum, in which case the returned point may be unmeaningful.
func (l *Lens) Project(p3 *math.Vec3) (p2 *math.Vec3, ok bool) {
	p4 := math.Vector4(p3.X, p3.Z, -p3.Y, 1.0)

	p4 = p4.Transform(l.Projection())
	if p4.W == 0 {
		p2 = math.Vec3Zero.Copy()
		ok = false
		return
	}

	recipW := 1.0 / p4.W
	p2 = math.Vector3(p4.X*recipW, p4.Z*recipW, p4.Y*recipW)

	xValid := (p2.X >= -1) && (p2.X <= 1)
	zValid := (p2.Z >= -1) && (p2.Z <= 1)
	ok = (p4.W > 0) && xValid && zValid
	return
}

// SetNearFar sets the near and far values of this lens to the specified values.
func (l *Lens) SetNearFar(near, far math.Real) {
	l.Lock()
	defer l.Unlock()

	l.near = near
	l.far = far
}

// NearFar returns the near and far values of this lens.
func (l *Lens) NearFar() (near, far math.Real) {
	l.RLock()
	defer l.RUnlock()

	return l.near, l.far
}

// SetProjection sets the projection matrix of this lens.
func (l *Lens) SetProjection(projection *math.Mat4) {
	l.Lock()
	defer l.Unlock()

	l.projection = projection.Copy()
}

// Projection returns the projection matrix of this lens.
func (l *Lens) Projection() *math.Mat4 {
	l.RLock()
	defer l.RUnlock()

	return l.projection.Copy()
}

// Copy returns an new 1:1 copy of this lens.
func (l *Lens) Copy() *Lens {
	l.RLock()
	defer l.RUnlock()

	c := new(Lens)

	c.projection = l.projection.Copy()
	c.near = l.near
	c.far = l.far

	c.perspectiveFovY = l.perspectiveFovY
	c.perspectiveAspect = l.perspectiveAspect

	c.orthoLeft = l.orthoLeft
	c.orthoRight = l.orthoRight
	c.orthoBottom = l.orthoBottom
	c.orthoTop = l.orthoTop

	c.frustumLeft = l.frustumLeft
	c.frustumRight = l.frustumRight
	c.frustumBottom = l.frustumBottom
	c.frustumTop = l.frustumTop

	return c
}

// SetPerspective sets the projection of this lens using the specified perspective values.
func (l *Lens) SetPerspective(fovY, aspectRatio math.Real) {
	l.Lock()
	defer l.Unlock()

	l.perspectiveFovY = fovY
	l.perspectiveAspect = aspectRatio
	l.projection.SetPerspective(fovY, aspectRatio, l.near, l.far)
}

// Perspective returns the frustum values of this lens, as they where passed into SetFrustum().
func (l *Lens) Perspective() (fovY, aspect math.Real) {
	l.RLock()
	defer l.RUnlock()

	return l.perspectiveFovY, l.perspectiveAspect
}

// SetOrtho sets the projection of this lens using the specified orthographic values.
func (l *Lens) SetOrtho(left, right, bottom, top math.Real) {
	l.Lock()
	defer l.Unlock()

	l.orthoLeft = left
	l.orthoRight = right
	l.orthoBottom = bottom
	l.orthoTop = top

	l.projection.SetOrtho(left, right, bottom, top, l.near, l.far)
}

// Ortho returns the orthographic values of this lens, as they where passed into SetOrtho().
func (l *Lens) Ortho() (left, right, bottom, top math.Real) {
	l.RLock()
	defer l.RUnlock()

	return l.orthoLeft, l.orthoRight, l.orthoBottom, l.orthoTop
}

// SetFrustum sets the projection of this lens using the specified frustum values.
func (l *Lens) SetFrustum(left, right, bottom, top math.Real) {
	l.Lock()
	defer l.Unlock()

	l.frustumLeft = left
	l.frustumRight = right
	l.frustumBottom = bottom
	l.frustumTop = top

	l.projection.SetFrustum(left, right, bottom, top, l.near, l.far)
}

// Frustum returns the frstum values of this lens, as they where passed into SetFrustum().
func (l *Lens) Frustum() (left, right, bottom, top math.Real) {
	l.RLock()
	defer l.RUnlock()

	return l.frustumLeft, l.frustumRight, l.frustumBottom, l.frustumTop
}

// PerspectiveLens returns an new initialized lens given the perspective inputs
func PerspectiveLens(fovY, aspectRatio, near, far math.Real) *Lens {
	l := new(Lens)
	l.projection = new(math.Mat4)
	l.SetNearFar(near, far)
	l.SetPerspective(fovY, aspectRatio)
	return l
}

// OrthoLens returns an new initialized lens given the orthographic inputs
func OrthoLens(left, right, bottom, top, near, far math.Real) *Lens {
	l := new(Lens)
	l.projection = new(math.Mat4)
	l.SetNearFar(near, far)
	l.SetOrtho(left, right, bottom, top)
	return l
}

// FrustumLens returns an new initialized lens given the frustum inputs
func FrustumLens(left, right, bottom, top, near, far math.Real) *Lens {
	l := new(Lens)
	l.projection = new(math.Mat4)
	l.SetNearFar(near, far)
	l.SetFrustum(left, right, bottom, top)
	return l
}
