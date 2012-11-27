//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: display.go
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
package display

//import "errors"
import "sync"
//import "fmt"

var gaccess sync.RWMutex
var linuxDisplayName string

// SetLinuxDisplayName sets the display_name string that will be passed into
// XOpenDisplay (See http://tronche.com/gui/x/xlib/display/opening.html)
//
// "Specifies the hardware display name, which determines the display and
// communications domain to be used. On a POSIX-conformant system, if the
// display_name is NULL, it defaults to the value of the DISPLAY environment
// variable."
func SetLinuxDisplayName(display_name string) {
    gaccess.Lock()
    defer gaccess.Unlock()
    linuxDisplayName = display_name
}

// GetLinuxDisplayName returns the string previously set by SetLinuxDisplayName
func GetLinuxDisplayName() string {
    gaccess.RLock()
    defer gaccess.RUnlock()
    return linuxDisplayName
}



// Display represents a window (etc) that you can render into using OpenGL, etc
type Display struct {
    ops *backend
    destroyed bool
}

// New returns a new initialized Display, or an error in the event that
// something goes wrong and you are unable to open a Display
func New(physicalDisplay int) (*Display, error) {
    var err error
    d := Display{}
    d.ops, err = newBackend(physicalDisplay)
    if err != nil {
        return nil, err
    }
    return &d, nil
}

// Destroy will cleanup all memory bound by this Display
// it's very very important that you call Destroy once you are
// completely finished using this struct!
// Typical usage is using defer
func (d *Display) Destroy() {
    if d.destroyed == false {
        d.ops.destroy()
        d.destroyed = true
    }
}

// NumPhysicalDisplays returns the number of physical displays attatched to
// the computer. In most cases, this is the number of physical monitors.
// or an error should there be a problem
func NumScreens() (int, error) {
    n, err := numScreens()
    return n, err
}

// DefaultPhysicalDisplay returns the default physical display that will be
// used, as recommended by the OS, or an error should there be a problem
func DefaultScreen() (int, error) {
    d, err := defaultScreen()
    return d, err
}



/*
// FrameBufferConfig represents configuration options for the Frame Buffer
// You must specify an FrameBufferConfig in your call to New() when creating
// an new Display, you can choose the technically 'best' configuration via
// GetBestFrameBufferConfig()
type FBConfig struct {
    xVisual int32
    Color, Red, Green, Blue, Alpha int
    Accum, AccumRed, AccumGreen, AccumBlue, AccumAlpha int
    Transparency, TransparencyRed, TransparencyGreen, TransparencyBlue, TransparencyAlpha int
    Depth, Stencil, AuxBuffers int
    DoubleBuffered, StereoScopic bool
}

func (f *FBConfig) String() string {
    return fmt.Sprintf("FBConfig(Color=%dbpp, Accum=%dbpp, Transparency=%d Depth=%d, Stencil=%d, AuxBuffers=%d, DoubleBuffered=%t, StereoScopic=%t)", f.Color, f.Accum, f.Transparency, f.Depth, f.Stencil, f.AuxBuffers, f.DoubleBuffered, f.StereoScopic)
}

// FBConfigs returns all available frame buffer configurations
func FBConfigs(physicalDisplay int) ([]*FBConfig, error) {
    return fbConfigs(physicalDisplay)
}

// BestFBConfig returns the technically 'best' frame buffer configuration
// that is available
func BestFBConfig(physicalDisplay int) (*FBConfig, error) {
    configs, err := FBConfigs(physicalDisplay)
    if err != nil {
        return nil, err
    }

    var bestConfig *FBConfig

    for i := 0; i < len(configs); i++ {
        config := configs[i]
        if bestConfig == nil {
            bestConfig = config
            continue
        }
        if (config.Color > bestConfig.Color) {
            bestConfig = config
        }
    }

    if bestConfig == nil {
        return nil, errors.New("No frame buffer configurations are available")
    }
    return bestConfig, nil
}
*/
