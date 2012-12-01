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

import "sync"

// Resolution represents a Screen's Resolution
type Resolution struct {
    // Required
    width, height uint16
    screen *Screen
    access sync.RWMutex
}

// Screen represents a physical screen, you can query which screen
// it is (number), and the size (width and height) in pixels
type Screen struct {
    // Required
    number uint16
    resolution *Resolution
    origResolution *Resolution
    origGammaRamp *Ramp
    autoRestoreOriginalGamma bool
    autoRestoreOriginalGammaCallback *Callback
    autoRestoreOriginalResolution bool
    autoRestoreOriginalResolutionCallback *Callback
    access sync.RWMutex

    // Platform
    xScreen *c_Screen
    xScreenNumber int32
}

// Helper to create a new screen struct
func newScreen(xScreen *c_Screen) *Screen {
    s := Screen{}
    s.xScreen = xScreen
    s.xScreenNumber = c_XScreenNumberOfScreen(s.xScreen)
    s.number = uint16(s.xScreenNumber + 1)

    resolution := Resolution{}
    resolution.width = uint16(c_XWidthOfScreen(s.xScreen))
    resolution.height = uint16(c_XHeightOfScreen(s.xScreen))
    resolution.screen = &s
    s.resolution = &resolution
    s.origResolution = &resolution
    s.autoRestoreOriginalResolutionCallback = &Callback{func(){
        s.RestoreOriginalResolution()
    }}
    s.autoRestoreOriginalResolution = true
    addDestroyCallback(s.autoRestoreOriginalResolutionCallback)

    s.origGammaRamp, _ = getGammaRamp(&s)
    s.autoRestoreOriginalGammaCallback = &Callback{func(){
        s.RestoreOriginalGamma()
    }}
    s.autoRestoreOriginalGamma = true
    addDestroyCallback(s.autoRestoreOriginalGammaCallback)
    return &s
}

func (s *Screen) resolutions() []*Resolution {
    modes := c_XF86VidModeGetAllModeLines(xDisplay, s.xScreenNumber)
    /*
    unsigned int	dotclock;
    unsigned short	hdisplay;
    unsigned short	hsyncstart;
    unsigned short	hsyncend;
    unsigned short	htotal;
    unsigned short	hskew;
    unsigned short	vdisplay;
    unsigned short	vsyncstart;
    unsigned short	vsyncend;
    unsigned short	vtotal;
    unsigned int	flags;
    int			privsize;
    */
    resolutions := []*Resolution{}
    for i := 0; i < len(modes); i++ {
        resolution := Resolution{}
        resolution.width = uint16(modes[i].hdisplay)
        resolution.height = uint16(modes[i].vdisplay)
        resolutions = append(resolutions, &resolution)
    }
    return resolutions
}

func (s *Screen) setResolution() {
}

func screens() []*Screen {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    screens := []*Screen{}
    screenCount := c_XScreenCount(xDisplay)
    for i := int32(0); i < screenCount; i++ {
        screen := c_XScreenOfDisplay(xDisplay, i)
        screens = append(screens, newScreen(screen))
    }
    return screens
}

func defaultScreen() *Screen {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    return newScreen(c_XDefaultScreenOfDisplay(xDisplay))
}

