//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: stats.go
// Created by: Stephen Gutekanst, 11/23/12
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
package main

import "code.google.com/p/azul3d/libraries/display"
import "fmt"

func main() {
    numScreens, err := display.NumScreens()
    if err != nil {
        panic(err.Error())
    }

    defaultScreen, err := display.DefaultScreen()
    if err != nil {
        panic(err.Error())
    }

    configs, err := display.FBConfigs(defaultScreen)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("There are", len(configs), "possible frame buffer configurations.")
    for i := 0; i < len(configs); i++ {
        fmt.Println("Config", i, configs[i])
    }

    fmt.Println("There are", numScreens, "screens (monitors)")
    fmt.Println("The OS says we should by default choose to use screen", defaultScreen)
}

