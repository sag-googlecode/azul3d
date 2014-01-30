// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

func (r *Renderer) stateClearViewport() {
	r.lastViewportX = 0xFFFFFF
	r.lastViewportY = 0xFFFFFF
	r.lastViewportWidth = 0xFFFFFF
	r.lastViewportHeight = 0xFFFFFF
}

func (r *Renderer) viewport(x, y, width, height float64) {
	// glViewport expects that coordinates are in bottom-left to top-right, but we expect the to use
	// top-left to bottom-right.
	vx, vy, vWidth, vHeight := regionToGL(r.width, r.height, x, y, width, height)

	if r.lastViewportX == vx && r.lastViewportY == vy && r.lastViewportWidth == vWidth && r.lastViewportHeight == vHeight {
		// Identical Viewport region as last one -- we don't need to do it again.
		return
	}
	r.lastViewportX = vx
	r.lastViewportY = vy
	r.lastViewportWidth = vWidth
	r.lastViewportHeight = vHeight
	r.gl.Viewport(vx, vy, vWidth, vHeight)
}
