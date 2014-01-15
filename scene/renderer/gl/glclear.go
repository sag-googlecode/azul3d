// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"code.google.com/p/azul3d/math"
)

func (r *Renderer) stateClearClearColor() {
	r.lastColorClear = nil
}

func (r *Renderer) clearColor(color *math.Vec4) {
	if r.lastColorClear == nil || !r.lastColorClear.Equals(color) {
		r.lastColorClear = color
		r.gl.ClearColor(float32(color.X), float32(color.Y), float32(color.Z), float32(color.W))
	}
}

func (r *Renderer) stateClearClearDepth() {
	r.lastDepthClear = 0xFFFFFF
}

func (r *Renderer) clearDepth(depth math.Real) {
	if !r.lastDepthClear.Equals(depth) {
		r.lastDepthClear = depth
		r.gl.ClearDepth(float64(depth))
	}
}

func (r *Renderer) stateClearClearStencil() {
	r.lastStencilClear = 0xFFFFFF
}

func (r *Renderer) clearStencil(stencil uint) {
	if r.lastStencilClear != stencil {
		r.lastStencilClear = stencil
		r.gl.ClearStencil(int32(stencil))
	}
}
