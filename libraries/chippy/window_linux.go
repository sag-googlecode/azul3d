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
    swa.background_pixmap = c_Pixmap(c_None).C()
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

func (w *Window) contextVersion() (uint8, uint8, uint8, error) {
    return 0, 0, 0, nil
}

func (w *Window) swapBuffers() error {
    c_glXSwapBuffers(xDisplay, w.window)
    return nil
}

func (w *Window) makeCurrent() error {
    return nil
}

func (w *Window) setTitle() error {
    err := c_XStoreName(xDisplay, w.window, w.title)
    if err != nil {
        return err
    }
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setSize() error {
    width := uint32(w.width)
    height := uint32(w.height)

    err := c_XResizeWindow(xDisplay, w.window, width, height)
    if err != nil {
        return err
    }
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setPos() error {
    x := int32(w.x)
    y := int32(w.y)

    err := c_XMoveWindow(xDisplay, w.window, x, y)
    if err != nil {
        return err
    }
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setVerticalSync() error {
    return nil
}

func (w *Window) setVisible() error {
    if w.visible {
        err := c_XMapWindow(xDisplay, w.window)
        if err != nil {
            return err
        }
    } else {
        err := c_XWithdrawWindow(xDisplay, w.window, w.screen.xScreenNumber)
        if err != nil {
            return err
        }
    }
    // It appears either XMapWindow or XWithdrawWindow (or both)
    // remove the size and/or position of the window so we set
    // it back right here (hence no XSync)
    err := w.setSize()
    err = w.setPos()
    if err != nil {
        return err
    }
    return nil
}

func (w *Window) setDecorated() error {
    hints := c_Hints{}
    hints.flags = 2 // Specify that we're changing the window decorations.
    if w.decorated && !w.fullscreen {
        hints.decorations = 1 // Decorations on
    } else {
        hints.decorations = 0 // Decorations off
    }
    property := c_XInternAtom(xDisplay, "_MOTIF_WM_HINTS", true)
    if property == 0 {
        return errors.New("Unable to remove window decorations; _MOTIF_WM_HINTS not supported!")
    }
    err := c_XChangeProperty(xDisplay, w.window, property, property, 32, c_PropModeReplace, c_Pointer(&hints), 5);
    if err != nil {
        return err
    }
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setMinimized() error {
    if w.minimized {
        err := c_XIconifyWindow(xDisplay, w.window, w.screen.xScreenNumber)
        if err != nil {
            return err
        }
    } else {
        // I hope this brings the window back
        // I've only been able to test on Ubuntu 12.10 with Unity,
        // but as it appears most WM's ignore restore window
        // requests anyway:
        // "Focus Prevention, only the user has the right to give a window focus"
        // We will try anyway here, heh.
        err := c_XRaiseWindow(xDisplay, w.window)
        if err != nil {
            return err
        }
    }
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setFullscreen() error {
    err := c_XRaiseWindow(xDisplay, w.window) // This raises the window above all others, or tries to anyway
    if err != nil {
        return err
    }
    c_setWindowFullscreen(xDisplay, w.window, w.fullscreen)
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) destroy() {
}

