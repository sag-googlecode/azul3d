//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: screens_linux.go
// Created by: Stephen Gutekanst, 11/24/12
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
package chippy

// #cgo LDFLAGS: -lX11 -lXxf86dga
// #include "includes_linux.h"
import "C"

// Screen represents a physical screen, you can query which screen
// it is (number), and the size (width and height) in pixels
type Screen struct {
    // Required
    number uint16
    width, height uint16
    origGammaRamp *Ramp

    // Platform
    xScreen *C.Screen
    xScreenNumber C.int
}

// Helper to create a new screen struct
func newScreen(xScreen *C.Screen) *Screen {
    s := Screen{}
    s.xScreen = xScreen
    s.xScreenNumber = C.XScreenNumberOfScreen(s.xScreen)
    s.number = uint16(s.xScreenNumber + 1)
    s.width = uint16(C.XWidthOfScreen(s.xScreen))
    s.height = uint16(C.XHeightOfScreen(s.xScreen))
    s.origGammaRamp, _ = getGammaRamp(&s)
    addDestroyCallback(func(){
        s.RestoreOriginalGamma()
    })
    return &s
}

func screens() []*Screen {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    screens := []*Screen{}
    screenCount := C.XScreenCount(xDisplay)
    for i := C.int(0); i < screenCount; i++ {
        screen := C.XScreenOfDisplay(xDisplay, i)
        screens = append(screens, newScreen(screen))
    }
    return screens
}

func defaultScreen() *Screen {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    return newScreen(C.XDefaultScreenOfDisplay(xDisplay))
}

