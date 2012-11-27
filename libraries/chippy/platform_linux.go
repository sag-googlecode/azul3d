//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: platform_linux.go
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

import "unsafe"
import "errors"
import "sync"

// Helper to open an X display
func xOpenDisplay() (*C.Display, error) {
    var display *C.Display

    if len(linuxDisplayName) > 0 {
        cstr := C.CString(linuxDisplayName)
        defer C.free(unsafe.Pointer(cstr))
        display = C.XOpenDisplay(cstr)
    } else {
        display = C.XOpenDisplay(nil)
    }

    if display == nil {
        return nil, errors.New("Unable to open an X11 connection! XOpenDisplay() failed.")
    }
    return display, nil
}

// Helper to close an X display
func xCloseDisplay(display *C.Display) {
    C.XCloseDisplay(display)
}

// Helper to get an atom
func atom(display *C.Display, atom string, only_if_exists bool) C.Atom {
    cstr := C.CString(atom)
    defer C.free(unsafe.Pointer(cstr))

    v := C.int(C.False)
    if only_if_exists {
        v = C.True
    }
    return C.XInternAtom(display, cstr, v)
}


var xDisplayAccess sync.RWMutex
var xDisplay *C.Display
func platformInit() error {
    C.XInitThreads()

    var err error
    xDisplay, err = xOpenDisplay()
    if err != nil {
        return err
    }

    go startEventLoop()
    return nil
}

func platformDestroy() {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()
    stopEventLoop()
    C.XFlush(xDisplay)
    C.XSync(xDisplay, C.False)
    xCloseDisplay(xDisplay)
}
