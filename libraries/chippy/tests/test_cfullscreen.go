//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: test_cfullscreen.go
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
package main

import "code.google.com/p/azul3d/libraries/chippy"
import "time"
import "fmt"

func main() {
    err := chippy.Init()
    if err != nil {
        panic(err.Error())
    }
    defer chippy.Destroy()

    minAttribs := chippy.FBConfig{
        RedBits: 1,
        BlueBits: 1,
        GreenBits: 1,
        AlphaBits: 1,
        DepthBits: 0,
        StencilBits: 0,
        Samples: 0,
        SampleBuffers: 0,
        DoubleBuffered: false,
    }

    defaultScreen, err := chippy.DefaultScreen()
    if err != nil {
        panic(err.Error())
    }

    win, err := chippy.NewWindow(defaultScreen, &minAttribs, chippy.BestFBConfig)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Opened a window with these Frame Buffer configurations:")
    contextVersion, err := win.ContextVersionString()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("The window is capable of OpenGL", contextVersion)

    t := 5 * time.Second
    time.Sleep(t)

    fmt.Println("Going fullscreen")
    win.SetFullscreen(true)
    time.Sleep(t)

    fmt.Println("Going windowed")
    win.SetFullscreen(false)
    time.Sleep(t)
}

