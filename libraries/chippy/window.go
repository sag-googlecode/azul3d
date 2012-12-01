//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: window.go
// Created by: Stephen Gutekanst, 11/27/12
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
//       documentation and/or other materials provided with the distribution.
//     * Neither the name of Lightpoke nor the
//       names of its contributors may be used to endorse or promote products
//       derived from n software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL LIGHTPOKE BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF n
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package chippy

import "fmt"

var BestFBConfig = &FBConfig{
    RedBits: 255, GreenBits: 255, BlueBits: 255, AlphaBits: 255,
    AccumRedBits: 255, AccumGreenBits: 255, AccumBlueBits: 255, AccumAlphaBits: 255,
    DepthBits: 255, StencilBits: 255, Samples: 255, SampleBuffers: 255, AuxBuffers: 255,
    DoubleBuffered: true, StereoScopic: true,
}

func (f *FBConfig) String() string {
    return fmt.Sprintf("FBConfig(RedBits=%d, greenBits=%d, BlueBits=%d, AlphaBits=%d, AccumRedBits=%d, AccumGreenBits=%d, AccumBlueBits=%d, AccumAlphaBits=%d, DepthBits=%d, StencilBits=%d, Samples=%d, SampleBuffers=%d, AuxBuffers=%d, DoubleBuffered=%t, StereoScopic=%t)", f.RedBits, f.GreenBits, f.BlueBits, f.AlphaBits, f.AccumRedBits, f.AccumGreenBits, f.AccumBlueBits, f.AccumAlphaBits, f.DepthBits, f.StencilBits, f.Samples, f.SampleBuffers, f.AuxBuffers, f.DoubleBuffered, f.StereoScopic)
}

// Use this to avoid an OS call to MakeCurrent in case someone calls it multiple times
var currentContext *Window

func CurrentContext() *Window {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    return currentContext
}

func freeContext(w *Window) {
    if currentContext == w {
        currentContext = nil
    }
}

// NewWindow returns a new, open window that the user will see visually
// You will recieve an window capable of at least minAttribs, and as close
// as possible to maxAttribs.
// This function is thread safe
func NewWindow(screen *Screen, minAttribs, maxAttribs *FBConfig) (*Window, error) {
    w := Window{}
    w.destroyCallback = &Callback{func(){
            w.Destroy()
        }}
    addDestroyCallback(w.destroyCallback)

    w.screen = screen
    w.minAttribs = minAttribs
    w.maxAttribs = maxAttribs
    w.vsync = true
    w.visible = true
    w.decorated = true
    w.title = "Chippy Window"
    w.width = 640
    w.height = 480
    w.x = int16((screen.Width() / 2) - (w.width / 2))
    w.y = int16((screen.Height() / 2) - (w.height / 2))

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    err := w.create()
    if err != nil {
        return nil, err
    }
    return &w, nil
}

// String returns a string representation of this Window
// This function is thread safe
func (w *Window) String() string {
    w.access.RLock()
    defer w.access.RUnlock()
    return fmt.Sprintf("Window(\"%s\", size=%dx%d, pos=%dx%d, vsync=%t, visible=%t, minimized=%t, decorated=%t, fullscreen=%t, destroyed=%t)", w.title, w.width, w.height, w.x, w.y, w.vsync, w.visible, w.minimized, w.decorated, w.fullscreen, w.destroyed)
}

// FBConfig returns the FBConfig that this Window is using, it is at least minAttribs,
// and no greater than maxAttribs, which you specified in your call to NewWindow()
// This function is thread safe
func (w *Window) FBConfig() *FBConfig {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return nil // Window is already destroyed
    }

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    return w.fbConfig
}

// Screen returns the Screen that this Window was created using
// This function is thread safe
func (w *Window) Screen() *Screen {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.screen
}

// MakeCurrent makes this window's OpenGL context current
// This function is thread safe, but do note that OpenGL is state based
// calling this function makes this OpenGL context the current context within
// the current thread only.
func (w *Window) MakeCurrent() {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    if currentContext != w {
        currentContext = w
        w.makeCurrent()
    }
}

// SwapBuffers swaps the background and foreground buffers of this Window
// You should always call this on Windows with DoubleBuffered FBConfigs,
// It is okay to call this function on a Window that has no DoubleBuffered FBConfig
// This function is thread safe
func (w *Window) SwapBuffers() {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if w.fbConfig.DoubleBuffered {
        
    }
}

// ContextVersion returns the major, minor, and revision versions of this OpenGL context. For example: [1, 2, 1], or [2, 0, 0]
// This function is thread safe
func (w *Window) ContextVersion() (uint8, uint8, uint8) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return 0, 0, 0 // Window is already destroyed
    }

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    return w.contextVersion()
}

// ContextVersionString returns the ContextVersion but as a string, for example: "1.2.1" or "2.0"
// Note that the last ('revision') version is only ommited (in the case of "2.0" vs "1.2.1") if
// the revision version is zero. The minor version is never ommited, in the case of "2.0" for instance
// This function is thread safe
func (w *Window) ContextVersionString() string {
    major, minor, revision := w.ContextVersion()
    if revision == 0 {
        return fmt.Sprintf("%d.%d", major, minor)
    }
    return fmt.Sprintf("%d.%d.%d", major, minor, revision)
}

// SetTitle sets the title string of this Window
// This function is thread safe
func (w *Window) SetTitle(title string) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }
    if title != w.title {
        w.title = title

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setTitle()
    }
}

// Title returns the title string of this Window
// This function is thread safe
func (w *Window) Title() string {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.title
}

// SetWidth sets the width (in pixels) of this Window
// This function is thread safe
func (w *Window) SetWidth(width uint16) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if width != w.width {
        w.width = width

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setSize()
    }
}

// Width returns the width (in pixels) of this Window
// This function is thread safe
func (w *Window) Width() uint16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.width
}

// SetHeight sets the height (in pixels) of this Window
// This function is thread safe
func (w *Window) SetHeight(height uint16) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if height != w.height {
        w.height = height

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setSize()
    }
}

// SetSize sets the width and height of this Window (in pixels)
// This function is thread safe
func (w *Window) SetSize(width, height uint16) {
    w.SetWidth(width)
    w.SetHeight(height)
}

// Size returns the width and height of this Window (in pixels)
// This function is thread safe
func (w *Window) Size() (uint16, uint16) {
    return w.Width(), w.Height()
}

// Height returns the height (in pixels) of this Window
// This function is thread safe
func (w *Window) Height() uint16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.height
}

// SetX sets the x position of this Window (in pixels)
// This function is thread safe
func (w *Window) SetX(x int16) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if x != w.x {
        w.x = x

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setPos()
    }
}

// X returns the x position of this Window (in pixels)
// This function is thread safe
func (w *Window) X() int16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.x
}

// SetY sets the y position of this Window (in pixels)
// This function is thread safe
func (w *Window) SetY(y int16) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if y != w.y {
        w.y = y

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setPos()
    }
}

// Y returns the y position of this Window (in pixels)
// This function is thread safe
func (w *Window) Y() int16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.y
}

// SetPos sets the x and y position of this Window (in pixels)
// This function is thread safe
func (w *Window) SetPos(x, y int16) {
    w.SetX(x)
    w.SetY(y)
}

// SetPosCenter sets this window to the center of Window.Screen()
// This just calls Window.SetPos(int16((Window.Screen().Width() / 2) - (Window.Width() / 2)), int16((Window.Screen().Height() / 2) - (Window.Height() / 2)))
// This function is thread safe
func (w *Window) SetPosCenter() {
    w.SetPos(int16((w.Screen().Width() / 2) - (w.Width() / 2)), int16((w.Screen().Height() / 2) - (w.Height() / 2)))
}

// Pos returns the x and y position of this Window (in pixels)
// This function is thread safe
func (w *Window) Pos() (int16, int16) {
    return w.X(), w.Y()
}

// SetVerticalSync sets weather vertical sync (vsync) will be on or off on this Window
// This function is thread safe
func (w *Window) SetVerticalSync(vsync bool) {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if vsync != w.vsync {
        w.vsync = vsync

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setVerticalSync()
    }
}

// VerticalSync returns weather this Window has vertical sync (vsync) enabled
// This function is thread safe
func (w *Window) VerticalSync() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.vsync
}

// SetHidden sets weather the window is visible or hidden
// This function is thread safe
func (w *Window) SetVisible(visible bool) {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if visible != w.visible {
        w.visible = visible

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setVisible()
    }
}

// Visible returns weather the window is visible or hidden
// This function is thread safe
func (w *Window) Visible() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.visible
}

// SetMinimized sets weather the window is minimized
// This function is thread safe
func (w *Window) SetMinimized(minimized bool) {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if minimized != w.minimized {
        w.minimized = minimized

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setMinimized()
    }
}

// Minimized returns weather the window is minimized
// This function is thread safe
func (w *Window) Minimized() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.minimized
}

// SetDecoration sets weather the window will have decorations
// This function is thread safe
func (w *Window) SetDecorated(decorated bool) {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if decorated != w.decorated {
        w.decorated = decorated

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setDecorated()
    }
}

// Decorated returns weather the window is decorated
// This function is thread safe
func (w *Window) Decorated() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.decorated
}

// SetFullscreen sets weather the window will be fullscreen
// This function is thread safe
func (w *Window) SetFullscreen(fullscreen bool) {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    if fullscreen != w.fullscreen {
        w.fullscreen = fullscreen

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()
        w.setFullscreen()
    }
}

// Fullscreen returns weather the window is fullscreen
// This function is thread safe
func (w *Window) Fullscreen() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.fullscreen
}

// Destroyed returns weather the window has been Destroyed, via a call to Destroy()
// This function is thread safe
func (w *Window) Destroyed() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.destroyed
}

// Destroy destroys the window, closing it visually
// It's only needed to call this if you want to close the window right away,
// otherwise the final call to chippy.Destroy() will call this for you.
// This function is thread safe
func (w *Window) Destroy() {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    w.destroyed = true
    removeDestroyCallback(w.destroyCallback)

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    freeContext(w)
    w.destroy()
}

