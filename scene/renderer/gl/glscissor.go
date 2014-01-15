// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

func (r *Renderer) scissor(x, y, width, height uint) {
	if x == 0 && y == 0 && width == 0 && height == 0 {
		width = uint(r.width)
		height = uint(r.height)
	}
	if r.lastScissorX == x && r.lastScissorY == y && r.lastScissorWidth == width && r.lastScissorHeight == height {
		// Identical scissor region as last one -- we don't need to do it again.
		return
	}
	r.lastScissorX = x
	r.lastScissorY = y
	r.lastScissorWidth = width
	r.lastScissorHeight = height

	// glScissor expects that coordinates are in bottom-left to top-right, but we expect the to use
	// top-left to bottom-right.
	sx, sy, sWidth, sHeight := regionToGL(uint(r.width), uint(r.height), x, y, width, height)
	r.gl.Scissor(sx, sy, uint32(sWidth), uint32(sHeight))
}

func (r *Renderer) pushScissor(x, y, width, height uint) {
	r.scissorStack = append(r.scissorStack, []uint{x, y, width, height})
	r.scissor(x, y, width, height)
}

func (r *Renderer) popScissor() {
	if len(r.scissorStack) > 0 {
		//last := r.scissorStack[len(r.scissorStack)-1]
		r.scissorStack = r.scissorStack[:len(r.scissorStack)-1]

		if len(r.scissorStack) > 0 {
			last := r.scissorStack[len(r.scissorStack)-1]
			r.scissor(last[0], last[1], last[2], last[3])
		} else {
			r.scissor(0, 0, 0, 0)
		}
	}
	return
}
