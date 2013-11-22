// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package util

import (
	"code.google.com/p/azul3d/math"
	"sync"
)

// Clearable implements an generic pixel-buffer clearable object.
type Clearable struct {
	access sync.RWMutex

	clearColor, clearDepth, clearStencil bool
	color                                *math.Vec4
	depth                                math.Real
	stencil                              uint
}

// Equals performs an deep equals comparison between the two clearable interfaces.
func (b *Clearable) Equals(other *Clearable) bool {
	if b.ClearColorActive() != other.ClearColorActive() {
		return false
	}

	if b.ClearDepthActive() != other.ClearDepthActive() {
		return false
	}

	if b.ClearStencilActive() != other.ClearStencilActive() {
		return false
	}

	if !b.ClearColor().Equals(other.ClearColor()) {
		return false
	}

	if !b.ClearDepth().Equals(other.ClearDepth()) {
		return false
	}

	if b.ClearStencil() != other.ClearStencil() {
		return false
	}
	return true
}

// SetClearColorActive specifies weather or not clearing the color buffer is enabled.
//
// Default: true
func (c *Clearable) SetClearColorActive(clear bool) {
	c.access.Lock()
	defer c.access.Unlock()

	c.clearColor = clear
}

// ClearColorActive tells if clearing the color buffer is enabled.
//
// Default: true
func (c *Clearable) ClearColorActive() bool {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.clearColor
}

// SetClearColor specifies the color value to be used when clearing the color buffer.
//
// An color value of nil, means to restore the default color value of light blue:
//
//  Vector4(0.53, 0.81, 0.98, 0)
//
func (c *Clearable) SetClearColor(color *math.Vec4) {
	c.access.Lock()
	defer c.access.Unlock()

	if color == nil {
		c.color = math.Vector4(0.53, 0.81, 0.98, 0)
	} else {
		c.color = color.Copy()
	}
}

// ClearColor returns the color value that is used when clearing the color buffer.
//
// The default color value is light blue:
//
//  Vector4(0.53, 0.81, 0.98, 0)
//
func (c *Clearable) ClearColor() *math.Vec4 {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.color.Copy()
}

// SetClearDepthActive specifies weather or not clearing the depth buffer is enabled.
//
// Default: true
func (c *Clearable) SetClearDepthActive(clear bool) {
	c.access.Lock()
	defer c.access.Unlock()

	c.clearDepth = clear
}

// ClearDepthActive tells if clearing the depth buffer is enabled.
//
// Default: true
func (c *Clearable) ClearDepthActive() bool {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.clearDepth
}

// SetClearDepth specifies the depth value to be used when clearing the depth buffer.
//
// The default depth value is 1.
func (c *Clearable) SetClearDepth(depth math.Real) {
	c.access.Lock()
	defer c.access.Unlock()

	c.depth = depth
}

// ClearDepth returns the depth value that is used when clearing the depth buffer.
//
// The default depth value is 1.
func (c *Clearable) ClearDepth() math.Real {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.depth
}

// SetClearStencilActive specifies weather or not clearing the stencil buffer is enabled.
//
// Default: true
func (c *Clearable) SetClearStencilActive(clear bool) {
	c.access.Lock()
	defer c.access.Unlock()

	c.clearStencil = clear
}

// ClearStencilActive tells if clearing the stencil buffer is enabled.
//
// Default: true
func (c *Clearable) ClearStencilActive() bool {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.clearStencil
}

// SetClearStencil specifies the stencil value to be used when clearing the stencil buffer.
//
// The default stencil value is 0.
func (c *Clearable) SetClearStencil(stencil uint) {
	c.access.Lock()
	defer c.access.Unlock()

	c.stencil = stencil
}

// ClearStencil returns the stencil value that is used when clearing the stencil buffer.
//
// The default stencil value is 0.
func (c *Clearable) ClearStencil() uint {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.stencil
}

// AnyClearActive tells if clearing of any of the color, depth, or stencil buffers is enabled.
func (c *Clearable) AnyClearActive() bool {
	c.access.RLock()
	defer c.access.RUnlock()

	return c.clearColor || c.clearDepth || c.clearStencil
}

// Clearable returns an new initialized *Clearable
func NewClearable() *Clearable {
	return &Clearable{
		clearColor:   true,
		clearDepth:   true,
		clearStencil: true,
		color:        math.Vector4(0.53, 0.81, 0.98, 1),
		depth:        1,
		stencil:      0,
	}
}
