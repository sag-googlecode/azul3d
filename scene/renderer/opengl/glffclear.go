// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/math"
)

func (r *GLFFRenderer) clearColor(color *math.Vec4) {
	if r.lastColorClear == nil || !r.lastColorClear.Equals(color) {
		if debug {
			logf("glClearColor(%v, %v, %v, %v)", color.X, color.Y, color.Z, color.W)
		}

		r.lastColorClear = color
		r.gl.ClearColor(float32(color.X), float32(color.Y), float32(color.Z), float32(color.W))
	}
}

func (r *GLFFRenderer) clearDepth(depth math.Real) {
	if !r.lastDepthClear.Equals(depth) {
		if debug {
			logf("glClearDepth(%v)", depth)
		}

		r.lastDepthClear = depth
		r.gl.ClearDepth(float64(depth))
	}
}

func (r *GLFFRenderer) clearStencil(stencil uint) {
	if r.lastStencilClear != stencil {
		if debug {
			logf("glClearStencil(%v)", stencil)
		}

		r.lastStencilClear = stencil
		r.gl.ClearStencil(int32(stencil))
	}
}
