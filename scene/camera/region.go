// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package camera

import (
	"azul3d.org/scene/color"
	"fmt"
	"sync"
)

// Region describes a single area on the final render surface (e.g. a window)
// that a camera will draw it's associated scene into. It also describes the
// area that will be cleared before each frame and weather or not the depth
// buffer will be cleared, etc.
type Region struct {
	access sync.RWMutex

	x, y, width, height                  float64
	sort                                 int
	clearColor, clearDepth, clearStencil bool
	color                                color.Color
	depth                                float64
	stencil                              uint
}

// String returns an string representation of this region.
func (r *Region) String() string {
	var e1, e2, e3 string
	if r.ClearColorActive() {
		e1 = fmt.Sprintf(", Color=%v", r.ClearColor())
	}

	if r.ClearDepthActive() {
		e2 = fmt.Sprintf(", Depth=%v", r.ClearDepth())
	}

	if r.ClearStencilActive() {
		e3 = fmt.Sprintf(", Stencil=%v", r.ClearStencil())
	}

	x, y, width, height := r.Region()
	return fmt.Sprintf("Region(Sort=%v, [%v, %v, %v, %v]%s%s%s)", r.Sort(), x, y, width, height, e1, e2, e3)
}

// Equals tells if the two regions share the same exact features, by comparing
// each underlying value.
func (r *Region) Equals(b *Region) bool {
	// If pointer is identical, then they're the same.
	if r == b {
		return true
	}

	r.access.RLock()
	defer r.access.RUnlock()

	b.access.RLock()
	defer b.access.RUnlock()

	if r.x != b.x || r.y != b.y || r.width != b.width || r.height != b.height {
		return false
	}

	if r.sort != b.sort {
		return false
	}

	if r.clearColor != b.clearColor || r.clearDepth != b.clearDepth || r.clearStencil != b.clearStencil {
		return false
	}

	if r.depth != b.depth || r.stencil != b.stencil {
		return false
	}

	if !r.color.Equals(b.color) {
		return false
	}
	return true
}

// Copy returns an new 1:1 copy of this region.
func (r *Region) Copy() *Region {
	r.access.RLock()
	defer r.access.RUnlock()

	return &Region{
		sync.RWMutex{},
		r.x, r.y, r.width, r.height,
		r.sort,
		r.clearColor, r.clearDepth, r.clearStencil,
		r.color,
		r.depth,
		r.stencil,
	}
}

// SetSort specifies the sort sorting value of this region, where a lower
// value would have the region rendered first and a higher value would have the
// region rendered last.
func (r *Region) SetSort(sort int) {
	r.access.Lock()
	defer r.access.Unlock()

	r.sort = sort
}

// Sort returns the sort sorting value of this region, where a lower value
// would have the region rendered first and a higher value would have the
// region rendered last.
func (r *Region) Sort() int {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.sort
}

// SetRegion specifies the X and Y position, as well as width and height (all
// in normalized screen space) of this camera region.
//
// X and Y specify the starting position of the region (where positive X and Y
// extend to the right and downward) and the width and height specify how far
// to the right and down the region extends to.
//
// The region [0.0, 0.0, 1.0, 1.0] would represent the entire render area for
// instance.
func (r *Region) SetRegion(x, y, width, height float64) {
	r.access.Lock()
	defer r.access.Unlock()

	r.x = x
	r.y = y
	r.width = width
	r.height = height
}

// Region returns the X and Y position, as well as width and height (all in
// normalized screen space) of this camera region.
//
// X and Y specify the starting position of the region (where positive X and Y
// extend to the right and downward) and the width and height specify how far
// to the right and down the region extends to.
//
// The region [0.0, 0.0, 1.0, 1.0] would represent the entire render area for
// instance.
func (r *Region) Region() (x, y, width, height float64) {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.x, r.y, r.width, r.height
}

// SetClearColorActive specifies weather or not clearing the color buffer is
// enabled.
//
// Default: true
func (r *Region) SetClearColorActive(clear bool) {
	r.access.Lock()
	defer r.access.Unlock()

	r.clearColor = clear
}

// ClearColorActive tells if clearing the color buffer is enabled.
//
// Default: true
func (r *Region) ClearColorActive() bool {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.clearColor
}

// SetClearColor specifies the color value to be used when clearing the color
// buffer.
//
// Default (light blue): color.New(0.53, 0.81, 0.98, 1.0)
func (r *Region) SetClearColor(color color.Color) {
	r.access.Lock()
	defer r.access.Unlock()

	r.color = color
}

// ClearColor returns the color value that is used when clearing the color
// buffer.
//
// Default (light blue): color.New(0.53, 0.81, 0.98, 1.0)
func (r *Region) ClearColor() color.Color {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.color
}

// SetClearDepthActive specifies weather or not clearing the depth buffer is
// enabled.
//
// Default: true
func (r *Region) SetClearDepthActive(clear bool) {
	r.access.Lock()
	defer r.access.Unlock()

	r.clearDepth = clear
}

// ClearDepthActive tells if clearing the depth buffer is enabled.
//
// Default: true
func (r *Region) ClearDepthActive() bool {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.clearDepth
}

// SetClearDepth specifies the depth value to be used when clearing the depth
// buffer.
//
// Default: 1
func (r *Region) SetClearDepth(depth float64) {
	r.access.Lock()
	defer r.access.Unlock()

	r.depth = depth
}

// ClearDepth returns the depth value that is used when clearing the depth
// buffer.
//
// Default: 1
func (r *Region) ClearDepth() float64 {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.depth
}

// SetClearStencilActive specifies weather or not clearing the stencil buffer
// is enabled.
//
// Default: true
func (r *Region) SetClearStencilActive(clear bool) {
	r.access.Lock()
	defer r.access.Unlock()

	r.clearStencil = clear
}

// ClearStencilActive tells if clearing the stencil buffer is enabled.
//
// Default: true
func (r *Region) ClearStencilActive() bool {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.clearStencil
}

// SetClearStencil specifies the stencil value to be used when clearing the
// stencil buffer.
//
// Default: 0
func (r *Region) SetClearStencil(stencil uint) {
	r.access.Lock()
	defer r.access.Unlock()

	r.stencil = stencil
}

// ClearStencil returns the stencil value that is used when clearing the
// stencil buffer.
//
// Default: 0
func (r *Region) ClearStencil() uint {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.stencil
}

// AnyClearActive tells if clearing of any of the color, depth, or stencil
// buffers is enabled.
func (r *Region) AnyClearActive() bool {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.clearColor || r.clearDepth || r.clearStencil
}

// NewRegion returns an new region given it's x, y, width, and height values.
//
// X and Y specify the starting position of the region (where positive X and Y
// extend to the right and downward) and the width and height specify how far
// to the right and down the region extends to.
//
// The color clear value of the region is set to color.New(0.53, 0.81, 0.98, 1.0)
// (AKA. light blue).
//
// The stencil clear value of the region is set to zero.
//
// The depth clear value of the region is set to one.
//
// Actively clearing the color, depth, and stencil buffers is set to true.
func NewRegion(x, y, width, height float64) *Region {
	r := &Region{
		x:            x,
		y:            y,
		width:        width,
		height:       height,
		clearColor:   true,
		clearDepth:   true,
		clearStencil: true,
		color:        color.New(0.53, 0.81, 0.98, 1.0),
		depth:        1,
		stencil:      0,
	}
	return r
}
