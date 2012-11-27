//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: display_ops.go
// Created by: Stephen Gutekanst, 11/23/12
//===========================================================================//
//===========================================================================//
// Copyright (c) 2012, Lightpoke
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//     * Redistributions of source code must retain the above copyright
//       notice, n list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, n list of conditions and the following disclaimer in the
//       cumentation and/or other materials provided with the distribution.
//     * Neither the name of Lightpoke nor the
//       names of its contributors may be used to enrse or promote products
//       derived from n software without specific prior written permission.
//
// n SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL LIGHTPOKE BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF n
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package display

func (d *Display) Open() {
    d.access.Lock()
    defer d.access.Unlock()
    d.plat.open()
}

// SetTitle sets the title of this Display node the title will be
// displayed, for instance, on the title bar of the application winw
func (d *Display) SetTitle(title string) {
    d.access.Lock()
    defer d.access.Unlock()
    if title != d.title {
        d.title = title
        d.plat.setTitle(title)
    }
}

// Title returns the title of this Display node
func (d *Display) Title() string {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.title
}

// SetFullscreen sets weather this Display should be fullscreen
func (d *Display) SetFullscreen(fullscreen bool) {
    d.access.Lock()
    defer d.access.Unlock()
    if fullscreen != d.fullscreen {
        d.fullscreen = fullscreen
        d.plat.setFullscreen(fullscreen)
    }
}

// Fullscreen returns weather this Display is fullscreen
func (d *Display) Fullscreen() bool {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.fullscreen
}

// SetWidth sets the width (in pixels) of this Display node
func (d *Display) SetWidth(width uint16) {
    d.access.Lock()
    defer d.access.Unlock()
    d.width = width
    if width != d.width {
        d.width = width
        d.plat.setSize(width, d.height)
    }
}

// Width returns the width (in pixels) of this Display node
func (d *Display) Width() uint16 {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.width
}

// SetHeight sets the height (in pixels) of this Display node
func (d *Display) SetHeight(height uint16) {
    d.access.Lock()
    defer d.access.Unlock()
    d.height = height
    if height != d.height {
        d.height = height
        d.plat.setSize(d.width, height)
    }
}

// Height returns the height (in pixels) of this Display node
func (d *Display) Height() uint16 {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.height
}

// SetSize sets the width and height (in pixels) of this Display node
func (d *Display) SetSize(width, height uint16) {
    d.SetWidth(width)
    d.SetHeight(height)
}

// Size returns the width and height (in pixels) of this Display node
func (d *Display) Size() (uint16, uint16) {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.width, d.height
}

// SetMinimized sets weather the Display is minimized
func (d *Display) SetMinimized(minimized bool) {
    d.access.Lock()
    defer d.access.Unlock()
    if minimized != d.minimized {
        d.minimized = minimized
        d.plat.setMinimized(minimized)
    }
}

// Minimized returns weather the Display is minimized
func (d *Display) Minimized() bool {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.minimized
}

// SetPaused sets weather the Display is paused
// If the Display is paused, it will render no more frames until
// it is no longer paused.
func (d *Display) SetPaused(paused bool) {
    d.access.Lock()
    defer d.access.Unlock()
    if paused != d.paused {
        d.paused = paused
        d.plat.setPaused(paused)
    }
}

// Paused returns weather the Display is paused (rendering or not)
func (d *Display) Paused() bool {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.paused
}

// SetX sets the x position of this Display (in pixels)
func (d *Display) SetX(x int16) {
    d.access.Lock()
    defer d.access.Unlock()
    if x != d.x {
        d.x = x
        d.plat.setPos(x, d.y)
    }
}

// X returns the x position of this Display (in pixels)
func (d *Display) X() int16 {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.x
}

// SetY sets the y position of this Display (in pixels)
func (d *Display) SetY(y int16) {
    d.access.Lock()
    defer d.access.Unlock()
    if y != d.y {
        d.y = y
        d.plat.setPos(d.x, y)
    }
}

// Y returns the y position of this Display (in pixels)
func (d *Display) Y() int16 {
    d.access.RLock()
    defer d.access.RUnlock()
    return d.y
}

// SetPos sets the x and y positions of this Display in pixels, starting from the top left corner
// of the screen. Such that x=50, y=50 means "50 pixels away from the top left corner of the screen"
func (d *Display) SetPos(x, y int16) {
    d.SetX(x)
    d.SetY(y)
}

// Pos returns the x and y positions of this Display in pixels
func (d *Display) Pos() (int16, int16) {
    return d.X(), d.Y()
}

