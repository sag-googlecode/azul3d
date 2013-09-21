// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package util

import (
	"fmt"
	"sync"
)

// An camera region is an area on which the camera will draw it's associated scene (it can draw
// the scene to multiple display regions, too).
type Region struct {
	*Clearable

	access sync.RWMutex

	id                  uint
	x, y, width, height uint
	sort                uint
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
	return fmt.Sprintf("Region(Id=%v, Sort=%v, [%v, %v, %v, %v]%s%s%s)", r.Id(), r.Sort(), x, y, width, height, e1, e2, e3)
}

// Id returns the unique identifier number of this region -- which is transfered when this region
// is copied.
func (r *Region) Id() uint {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.id
}

// Equals tells if the two regions share the same exact features, by comparing each underlying
// value.
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

	return r.Clearable.Equals(b.Clearable)
}

// Copy returns an new 1:1 copy of this region.
//
// Note: If you intend to use this new copied region as an individual region itself, you should
// invoke ResetId() right after calling this function.
func (r *Region) Copy() *Region {
	r.access.RLock()
	defer r.access.RUnlock()

	c := new(Region)
	c.id = r.id
	c.x = r.x
	c.y = r.y
	c.width = r.width
	c.height = r.height
	c.sort = r.sort

	c.Clearable = new(Clearable)
	*c.Clearable = *r.Clearable
	return c
}

// SetSort specifies the sort value of this region, an lower value means that this region will be
// drawn first, an higher sort value means this region will be drawn last.
func (r *Region) SetSort(sort uint) {
	r.access.Lock()
	defer r.access.Unlock()

	r.sort = sort
}

// Sort returns the sort value of this region, an lower value means that this region will be drawn
// first, an higher sort value means this region will be drawn last.
func (r *Region) Sort() uint {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.sort
}

// SetRegion specifies the X and Y position, as well as width and height (all in pixels) of this
// camera region, where X and Y specify the starting position of the region (where +X and +Y extend
// to the right and downward) and the width and height specify how far to the right and down the
// region extends to.
//
// The special region of [x=0, y=0, width=0, height=0] implies that the region should cover the
// entire area.
func (r *Region) SetRegion(x, y, width, height uint) {
	r.access.Lock()
	defer r.access.Unlock()

	r.x = x
	r.y = y
	r.width = width
	r.height = height
}

// Region returns the X and Y position, as well as width and height (all in pixels) of this camera
// region, where X and Y specify the starting position of the region (where +X and +Y extend to the
// right and downward) and the width and height specify how far to the right and down the region
// extends to.
//
// The special region of [x=0, y=0, width=0, height=0] implies that the region should cover the
// entire area.
func (r *Region) Region() (x, y, width, height uint) {
	r.access.RLock()
	defer r.access.RUnlock()

	return r.x, r.y, r.width, r.height
}

var (
	regionIdCounterAccess sync.Mutex
	regionIdCounter       uint
)

// ResetId resets the unique identifier number of this region. When an region is copied, the id
// number will also be copied. This allows comparison of region's id numbers to determine if they
// are equal.
//
// However, if you make an copy of an region and intend to use it as an seperate region (and do not
// wish for it to be confused with the region you originally copied) then invoking this function
// right after an Copy() operation is needed.
func (r *Region) ResetId() {
	regionIdCounterAccess.Lock()
	regionIdCounter++
	ourId := regionIdCounter
	regionIdCounterAccess.Unlock()

	r.access.Lock()
	defer r.access.Unlock()

	r.id = ourId
}

// NewRegion returns an new region given it's x, y, width, and height values.
//
// The color clear value is by default set to math.Vector4(0.53, 0.81, 0.98, 0) (AKA. light blue).
//
// The stencil clear value defaults to 0.
//
// The depth clear value defaults to 1.
//
func NewRegion(x, y, width, height uint) *Region {
	r := &Region{
		Clearable: NewClearable(),
		x:         x,
		y:         y,
		width:     width,
		height:    height,
	}
	r.ResetId()
	return r
}
