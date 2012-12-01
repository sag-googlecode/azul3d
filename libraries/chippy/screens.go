//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: screens.go
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

import "fmt"

// String returns a string representation of this Resolution
// This function is thread safe
func (r *Resolution) String() string {
    r.access.RLock()
    defer r.access.RUnlock()
    return fmt.Sprintf("Resolution(width=%d, height=%d)", r.width, r.height)
}

// Use calls Resolution.Screen().SetResolution for you, setting this resolution
// to be active on the Screen
// This function is thread safe
func (r *Resolution) Use() {
    r.Screen().SetResolution(r)
}

// Screen returns the Screen that this resolution is from
// This function is thread safe
func (r *Resolution) Screen() *Screen {
    r.access.RLock()
    defer r.access.RUnlock()
    return r.screen
}

// Width returns the width in pixels of this Resolution
// This function is thread safe
func (r *Resolution) Width() uint16 {
    r.access.RLock()
    defer r.access.RUnlock()
    return r.width
}

// Height returns the height in pixels of this Resolution
// This function is thread safe
func (r *Resolution) Height() uint16 {
    r.access.RLock()
    defer r.access.RUnlock()
    return r.height
}




// String returns a formatted string representation of this Screen
// This function is thread safe
func (s *Screen) String() string {
    s.access.RLock()
    defer s.access.RUnlock()
    return fmt.Sprintf("Screen(%d, %dx%d)", s.number, s.resolution.Width(), s.resolution.Height())
}

// Number returns the number (id) of this screen
// This function is thread safe
func (s *Screen) Number() uint16 {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.number
}

// Width returns the width in pixels of this Screen
// This function is thread safe
func (s *Screen) Width() uint16 {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.resolution.Width()
}

// Height returns the height in pixels of this Screen
// This function is thread safe
func (s *Screen) Height() uint16 {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.resolution.Height()
}

// SetResolution sets the current Resolution of this Screen
func (s *Screen) SetResolution(resolution *Resolution) {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling C things, so get lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    s.setResolution()
}

// RestoreOriginalResolution restores the original resolution of this screen
// just like it was when the application started
// Chippy will automatically call this for you when you call chippy.Destroy()
// You can turn off this default behavior with screen.SetAutoRestoreOriginalResolution(bool)
// This function is thread safe
func (s *Screen) RestoreOriginalResolution() {
    s.access.RLock()
    oldResolution := s.origResolution
    s.access.RUnlock()// Important to release before call to SetResolution
    s.SetResolution(oldResolution)
}


// SetAutoRestoreOriginalResolution if restore is true, chippy will automatically restore the resolution of this screen
// as it was when the application started. If restore is false, then chippy will instead leave the resolution as-is
// when the application finally calls chippy.Destroy()
// This function is thread safe
func (s *Screen) SetAutoRestoreOriginalResolution(restore bool) {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling add/remove DestroyCallback
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    if restore != s.autoRestoreOriginalResolution {
        s.autoRestoreOriginalResolution = restore
        if restore == true {
            addDestroyCallback(s.autoRestoreOriginalResolutionCallback)
        } else {
            removeDestroyCallback(s.autoRestoreOriginalResolutionCallback)
        }
    }
}

// Resolution returns the current Resolution of this Screen
// This function is thread safe
func (s *Screen) Resolution() *Resolution {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.resolution
}

// OriginalResolution returns the current Resolution of this Screen
// This function is thread safe
func (s *Screen) OriginalResolution() *Resolution {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.origResolution
}

// Resolutions returns all available resolutions for this Screen
// This function is thread safe
func (s *Screen) Resolutions() []*Resolution {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling C things, so get lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return s.resolutions()
}



// RestoreOriginalGamma restores the original gamma ramp of this screen
// just like it was when the application started
// Chippy will automatically call this for you when you call chippy.Destroy()
// You can turn off this default behavior with screen.SetAutoRestoreOriginalGamma(bool)
// This function is thread safe
func (s *Screen) RestoreOriginalGamma() {
    s.access.RLock()
    oldRamp := s.origGammaRamp
    s.access.RUnlock()
    s.SetGammaRamp(oldRamp)
}

// SetAutoRestoreOriginalGamma if restore is true, chippy will automatically restore the gamma of this screen
// as it was when the application started. If restore is false, then chippy will instead leave the gamma as-is
// when the application finally calls chippy.Destroy()
// This function is thread safe
func (s *Screen) SetAutoRestoreOriginalGamma(restore bool) {
    s.access.Lock()
    defer s.access.Unlock()

    // Calling add/remove DestroyCallback
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    if restore != s.autoRestoreOriginalGamma {
        s.autoRestoreOriginalGamma = restore
        if restore == true {
            addDestroyCallback(s.autoRestoreOriginalGammaCallback)
        } else {
            removeDestroyCallback(s.autoRestoreOriginalGammaCallback)
        }
    }
}

// OriginalGamma returns the gamma Ramp that was in use when this application started
func (s *Screen) OriginalGamma() *Ramp {
    s.access.RLock()
    defer s.access.RUnlock()
    return s.origGammaRamp
}


// Screens returns all available screens
// This function is thread safe
func Screens() []*Screen {
    // Calling C things, so get lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return screens()
}

// DefaultScreen returns the default screen
// This function is thread safe
func DefaultScreen() *Screen {
    // Calling C things, so get lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return defaultScreen()
}

