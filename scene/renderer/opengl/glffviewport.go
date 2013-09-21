// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

func (r *GLFFRenderer) viewport(x, y uint, width, height int) {
	if debug {
		logf("\nviewport(%v, %v, %v, %v)", x, y, width, height)
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
	if debug {
		logf("    -> glViewport(%v, %v, %v, %v)", vx, vy, vWidth, vHeight)
	}
	r.gl.Viewport(vx, vy, vWidth, vHeight)
}
