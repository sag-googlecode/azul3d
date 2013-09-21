// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

func (r *GLFFRenderer) scissor(x, y, width, height uint) {
	if debug {
		logf("\nscissor(%v, %v, %v, %v)", x, y, width, height)
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
	if debug {
		logf("    -> glScissor(%v, %v, %v, %v)", sx, sy, sWidth, sHeight)
	}
	r.gl.Scissor(sx, sy, sWidth, sHeight)
}

func (r *GLFFRenderer) pushScissor(x, y, width, height uint, restricted bool) {
	// Ensure region does not exceed any region who is already on the stack
	if restricted {
		for _, region := range r.scissorStack {
			if x < region[0] {
				x = region[0]
			}
			if y < region[1] {
				y = region[1]
			}
			if width > region[2] {
				width = region[2]
			}
			if height > region[3] {
				height = region[3]
			}
		}
	}

	r.scissorStack = append(r.scissorStack, []uint{x, y, width, height})
	r.scissor(x, y, width, height)
}

func (r *GLFFRenderer) popScissor() {
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
