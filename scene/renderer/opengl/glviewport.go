// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

func (r *Renderer) stateClearViewport() {
	r.lastViewportX = 0xFFFFFF
	r.lastViewportY = 0xFFFFFF
	r.lastViewportWidth = 0xFFFFFF
	r.lastViewportHeight = 0xFFFFFF
}

func (r *Renderer) viewport(x, y uint, width, height int) {
	if x == 0 && y == 0 && width == 0 && height == 0 {
		width = r.width
		height = r.height
	}
	if r.lastViewportX == x && r.lastViewportY == y && int(r.lastViewportWidth) == width && int(r.lastViewportHeight) == height {
		// Identical Viewport region as last one -- we don't need to do it again.
		return
	}
	r.lastViewportX = x
	r.lastViewportY = y
	r.lastViewportWidth = uint(width)
	r.lastViewportHeight = uint(height)

	// glViewport expects that coordinates are in bottom-left to top-right, but we expect the to use
	// top-left to bottom-right.
	vx, vy, vWidth, vHeight := regionToGL(uint(r.width), uint(r.height), x, y, uint(width), uint(height))
	r.gl.Viewport(vx, vy, vWidth, vHeight)
}
