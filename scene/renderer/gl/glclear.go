// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"code.google.com/p/azul3d/scene/color"
)

func (r *Renderer) stateClearClearColor() {
	r.lastColorClear = color.None
}

func (r *Renderer) clearColor(c color.Color) {
	if r.lastColorClear.Equals(color.None) || !r.lastColorClear.Equals(c) {
		r.lastColorClear = c
		r.gl.ClearColor(c.R, c.G, c.B, c.A)
	}
}

func (r *Renderer) stateClearClearDepth() {
	r.lastDepthClear = 0xFFFFFF
}

func (r *Renderer) clearDepth(depth float64) {
	if r.lastDepthClear != depth {
		r.lastDepthClear = depth
		r.gl.ClearDepth(depth)
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
