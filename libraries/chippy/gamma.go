//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: gamma.go
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

import "math"

// Re-writes of the SDL functions for these next two helpers
func calculateGammaRamp(gamma float32) [256]uint16 {
    ramp := [256]uint16{}
    if gamma <= 0.0 {
        // 0.0 gamma is all black
        for i := 0; i < len(ramp); i++ {
            ramp[i] = 0
        }
        return ramp
    } else if gamma == 1.0 {
        // 1.0 is identity
        for i := 0; i < len(ramp); i++ {
            ramp[i] = uint16((i << 8) | i)
        }
        return ramp
    }

    // Calculate a real gamma ramp
    gamma = 1.0 / gamma
    for i := 0; i < len(ramp); i++ {
        value := int32(math.Pow(float64(i) / 256.0, float64(gamma)) * 65535.0 + 0.5)
        if value > 65535 {
            value = 65535
        }
        ramp[i] = uint16(value)
    }
    return ramp
}

func calculateGammaFromRamp(ramp [256]uint16) float32 {
    gamma := 1.0
    sum := 0.0
    count := 0

    for i := 1; i < len(ramp); i++ {
        if (ramp[i] != 0) && (ramp[i] != 65535) {
            B := float64(i) / 256.0
            A := float64(ramp[i]) / 65535.0
            sum += math.Log(A) / math.Log(B)
            count += 1
        }
    }

    if count != 0 && sum > 0.0 {
        gamma = 1.0 / (sum / float64(count))
    }
    return float32(gamma)
}




// Ramp represents a ramp for the gamma of the screen
type Ramp struct {
    Red [256]uint16
    Green [256]uint16
    Blue [256]uint16
}

// SetGammaRamp sets the currently in use Ramp for the gamma
// This call is thread safe
func (s *Screen) SetGammaRamp(ramp *Ramp) error {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return err
    }

    return setGammaRamp(s, ramp)
}

// GammaRamp returns the currently in use Ramp for the gamma
// This call is thread safe
func (s *Screen) GammaRamp() (*Ramp, error) {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    return getGammaRamp(s)
}

// SetGammaRgb sets the gamma of the screen as rgb float32
// This call is thread safe
func (s *Screen) SetGammaRgb(r, g, b float32) error {
    err := getInitError()
    if err != nil {
        return err
    }

    ramp := Ramp{}

    ramp.Red = calculateGammaRamp(r)
    ramp.Green = calculateGammaRamp(g)
    ramp.Blue = calculateGammaRamp(b)

    return s.SetGammaRamp(&ramp)
}

// SetGamma sets the rgb gamma of the screen as an float32
// This call is thread safe
func (s *Screen) SetGamma(gamma float32) error {
    return s.SetGammaRgb(gamma, gamma, gamma)
}

// GammaRgb returns the gamma of the screen as rgb float32
func (s *Screen) GammaRgb() (float32, float32, float32, error) {
    err := getInitError()
    if err != nil {
        return 0, 0, 0, err
    }

    ramp, err := s.GammaRamp()
    if err != nil {
        return 0, 0, 0, err
    }

    r := calculateGammaFromRamp(ramp.Red)
    g := calculateGammaFromRamp(ramp.Green)
    b := calculateGammaFromRamp(ramp.Blue)
    return r, g, b, nil
}

// Gamma returns the rgb gamma of the screen as an float32
// This call is thread safe
func (s *Screen) Gamma() (float32, error) {
    r, g, b, err := s.GammaRgb()
    if err != nil {
        return 0.0, err
    }
    return (r + g + b) / 3.0, nil
}

