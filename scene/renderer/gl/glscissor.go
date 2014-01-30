// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

func (r *Renderer) scissor(x, y, width, height float64) {
	// glScissor expects that coordinates are in bottom-left to top-right, but we expect the to use
	// top-left to bottom-right.
	sx, sy, sWidth, sHeight := regionToGL(r.width, r.height, x, y, width, height)

	if r.lastScissorX == sx && r.lastScissorY == sy && r.lastScissorWidth == sWidth && r.lastScissorHeight == sHeight {
		// Identical scissor region as last one -- we don't need to do it again.
		return
	}
	r.lastScissorX = sx
	r.lastScissorY = sy
	r.lastScissorWidth = sWidth
	r.lastScissorHeight = sHeight
	r.gl.Scissor(sx, sy, sWidth, sHeight)
}

func (r *Renderer) pushScissor(x, y, width, height float64) {
	r.scissorStack = append(r.scissorStack, []float64{x, y, width, height})
	r.scissor(x, y, width, height)
}

func (r *Renderer) popScissor() {
	if len(r.scissorStack) > 0 {
		r.scissorStack = r.scissorStack[:len(r.scissorStack)-1]

		if len(r.scissorStack) > 0 {
			last := r.scissorStack[len(r.scissorStack)-1]
			r.scissor(last[0], last[1], last[2], last[3])
		} else {
			r.scissor(0, 0, 1, 1)
		}
	}
	return
}
