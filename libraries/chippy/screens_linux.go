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
import "errors"

// Resolution represents a Screen's Resolution
type Resolution struct {
    // Required
    width, height uint16
    screen *Screen
    access sync.RWMutex

    // Platform
    videoMode *c_XF86VidModeModeInfo
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
func newScreen(xScreen *c_Screen) (*Screen, error) {
    s := Screen{}
    s.xScreen = xScreen
    s.xScreenNumber = c_XScreenNumberOfScreen(s.xScreen)
    s.number = uint16(s.xScreenNumber + 1)

    // The largest video mode available is the screen
    modes, err := s.resolutions()
    if err != nil {
        return nil, err
    }
    var working *Resolution
    for i := 0; i < len(modes); i++ {
        current := modes[i]
        if int32(current.Width()) == c_XWidthOfScreen(s.xScreen) && int32(current.Height()) == c_XHeightOfScreen(s.xScreen) {
            working = current
            break
        }
    }
    if working == nil {
        return nil, errors.New("Unable to locate current screen resolution; does glxinfo report a valid fbconfig?")
    }

    s.resolution = working
    s.origResolution = working
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
    return &s, nil
}

func (s *Screen) resolutions() ([]*Resolution, error) {
    modes := c_XF86VidModeGetAllModeLines(xDisplay, s.xScreenNumber)
    resolutions := []*Resolution{}
    for i := 0; i < len(modes); i++ {
        mode := modes[i]

        resolution := Resolution{}
        resolution.screen = s
        resolution.width = uint16(mode.hdisplay)
        resolution.height = uint16(mode.vdisplay)
        resolution.videoMode = mode
        resolutions = append(resolutions, &resolution)
    }
    return resolutions, nil
}

func (s *Screen) setResolution() error {
    // Warp the pointer to the upper left corner, this is
    // necessary as the: XF86VidModeSetViewPort() call
    // below does not seem to do anything on newer Xorg
    // servers, instead the viewport appears to be always
    // centered at the last mouse position
    c_XWarpPointer(xDisplay, c_Window(c_None), c_XDefaultRootWindow(xDisplay), 0, 0, 0, 0, 0, 0)

    err := c_XF86VidModeSwitchToMode(xDisplay, s.xScreenNumber, s.resolution.videoMode)
    if err != nil {
        return err
    }
    c_XF86VidModeSetViewPort(xDisplay, s.xScreenNumber, 0, 0)
    c_XSync(xDisplay, false)
    return nil
}

func screens() ([]*Screen, error) {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    screens := []*Screen{}
    screenCount := c_XScreenCount(xDisplay)
    for i := int32(0); i < screenCount; i++ {
        screen := c_XScreenOfDisplay(xDisplay, i)
        s, err := newScreen(screen)
        if err != nil {
            return nil, err
        }
        screens = append(screens, s)
    }
    return screens, nil
}

func defaultScreen() (*Screen, error) {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    return newScreen(c_XDefaultScreenOfDisplay(xDisplay))
}

func screensInit() error {
    var err error
    platformDefaultScreen, err = defaultScreen()
    if err != nil {
        return err
    }
    platformScreens, err = screens()
    if err != nil {
        return err
    }
    return nil
}

func screensDestroy() {
    platformDefaultScreen = nil
    platformScreens = nil
}

