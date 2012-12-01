//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: window_linux.go
// Created by: Stephen Gutekanst, 11/25/12
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

import "errors"
import "sync"

// FBConfig represents options to configure the Frame Buffer
type FBConfig struct {
    // Required
    RedBits, GreenBits, BlueBits, AlphaBits uint8
    AccumRedBits, AccumGreenBits, AccumBlueBits, AccumAlphaBits uint8
    DepthBits, StencilBits, Samples, SampleBuffers, AuxBuffers uint8
    DoubleBuffered, StereoScopic bool

    // Platform
    actual c_GLXFBConfig
}

// Window represents a visual window that the user will see and that you will render to
type Window struct {
    // Required
    screen *Screen
    minAttribs, maxAttribs, fbConfig *FBConfig
    destroyed, vsync, decorated, visible, minimized, fullscreen bool
    title string
    width, height uint16
    x, y int16
    destroyCallback *Callback
    access sync.RWMutex

    // Platform
    colormap c_Colormap
    window c_Window
}

func (w *Window) create() error {
    w.fbConfig = c_chooseFBConfig(xDisplay, w.screen.xScreenNumber, w.minAttribs, w.maxAttribs)
    if w.fbConfig == nil {
        return errors.New("Unable to retrieve a matching FBConfig")
    }
    vi := c_glXGetVisualFromFBConfig(xDisplay, w.fbConfig.actual)
    if vi == nil {
        return errors.New("Unexpected FBConfig is invalid! glXGetVisualFromFBConfig() failed!")
    }

    parent := c_XRootWindow(xDisplay, int32(vi.screen))

    cmap := c_XCreateColormap(xDisplay, parent, (*c_Visual)(vi.visual), c_AllocNone)
    defer c_XFreeColormap(xDisplay, cmap)

    swa := c_XSetWindowAttributes{}
    swa.colormap = cmap.C()
    swa.background_pixmap = c_None
    swa.border_pixel = 0
    swa.event_mask = c_StructureNotifyMask

    w.window = c_XCreateWindow(xDisplay, parent, int32(w.x), int32(w.y), uint32(w.width), uint32(w.height), 0, int32(vi.depth), c_InputOutput, (*c_Visual)(vi.visual), c_CWBorderPixel|c_CWColormap|c_CWEventMask, &swa)
    if w.window == 0 {
        return errors.New("Failed to create window; XCreateWindow() failed!")
    }

    w.setTitle()
    w.setSize()
    w.setPos()
    w.setVerticalSync()
    w.setVisible()
    w.setDecorated()
    w.setMinimized()
    w.setFullscreen()
    return nil
}

func (w *Window) contextVersion() (uint8, uint8, uint8) {
    return 0, 0, 0
}

func (w *Window) swapBuffers() {
    c_glXSwapBuffers(xDisplay, w.window)
}

func (w *Window) makeCurrent() {
}

func (w *Window) setTitle() {
    c_XStoreName(xDisplay, w.window, w.title)
    c_XSync(xDisplay, false)
}

func (w *Window) setSize() {
    c_XResizeWindow(xDisplay, w.window, uint32(w.width), uint32(w.height))
    c_XSync(xDisplay, false)
}

func (w *Window) setPos() {
    c_XMoveWindow(xDisplay, w.window, int32(w.x), int32(w.y))
    c_XSync(xDisplay, false)
}

func (w *Window) setVerticalSync() {
}

func (w *Window) setVisible() {
    if w.visible {
        c_XMapWindow(xDisplay, w.window)
    } else {
        c_XWithdrawWindow(xDisplay, w.window, w.screen.xScreenNumber)
    }
    // It appears either XMapWindow or XWithdrawWindow (or both)
    // remove the size and/or position of the window so we set
    // it back right here (hence no XSync)
    w.setSize()
    w.setPos()
}

func (w *Window) setDecorated() {
    hints := c_Hints{}
    hints.flags = 2 // Specify that we're changing the window decorations.
    if w.decorated {
        hints.decorations = 1 // Decorations on
    } else {
        hints.decorations = 0 // Decorations off
    }
    property := c_XInternAtom(xDisplay, "_MOTIF_WM_HINTS", true)
    if property == 0 {
        return
    }
    c_XChangeProperty(xDisplay, w.window, property, property, 32, c_PropModeReplace, c_Pointer(&hints), 5);
    c_XSync(xDisplay, false)
}

func (w *Window) setMinimized() {
    if w.minimized {
        c_XIconifyWindow(xDisplay, w.window, w.screen.xScreenNumber)
    } else {
        // I hope this brings the window back
        // I've only been able to test on Ubuntu 12.10 with Unity,
        // but as it appears most WM's ignore restore window
        // requests anyway:
        // "Focus Prevention, only the user has the right to give a window focus"
        // We will try anyway here, heh.
        c_XRaiseWindow(xDisplay, w.window)
    }
    c_XSync(xDisplay, false)
}

func (w *Window) setFullscreen() {
}

func (w *Window) destroy() {
}

