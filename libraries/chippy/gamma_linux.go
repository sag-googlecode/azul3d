//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: gamma_linux.go
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

// #cgo LDFLAGS: -lX11 -lXxf86vm
// #include "includes_linux.h"
import "C"

import "unsafe"
import "errors"

func setGammaRamp(screen *Screen, ramp *Ramp) error {
    // We make assumption that gamma ramp is 256
    var size C.int
    C.XF86VidModeGetGammaRampSize(xDisplay, screen.xScreenNumber, &size)
    if size != 256 {
        return errors.New("Gamma ramp size > 256")
    }

    // Could have been non-256 before, so no setting gama
    if ramp == nil {
        return errors.New("Gamma ramp is nil")
    }

    r := (*C.ushort)(unsafe.Pointer(&ramp.Red))
    g := (*C.ushort)(unsafe.Pointer(&ramp.Green))
    b := (*C.ushort)(unsafe.Pointer(&ramp.Blue))
    worked := C.XF86VidModeSetGammaRamp(xDisplay, screen.xScreenNumber, 256, r, g, b)
    if worked == C.False {
        return errors.New("Unable to set gamma ramp; XF86VidModeSetGammaRamp()")
    }
    return nil
}

func getGammaRamp(screen *Screen) (*Ramp, error) {
    // We make assumption that gamma ramp is 256
    var size C.int
    C.XF86VidModeGetGammaRampSize(xDisplay, screen.xScreenNumber, &size)
    if size != 256 {
        return nil, errors.New("Gamma ramp size > 256")
    }

    var r, g, b [256]C.ushort
    worked := C.XF86VidModeGetGammaRamp(xDisplay, screen.xScreenNumber, 256, (*C.ushort)(unsafe.Pointer(&r)), (*C.ushort)(unsafe.Pointer(&g)), (*C.ushort)(unsafe.Pointer(&b)))
    if worked == C.False {
        return nil, errors.New("Unable to get gamma ramp; XF86VidModeGetGammaRamp() failed")
    }

    ramp := Ramp{}
    // Red
    for i := 0; i < 256; i++ {
        ramp.Red[i] = uint16(r[i])
    }
    // Green
    for i := 0; i < 256; i++ {
        ramp.Green[i] = uint16(g[i])
    }
    // Blue
    for i := 0; i < 256; i++ {
        ramp.Blue[i] = uint16(b[i])
    }
    return &ramp, nil
}

