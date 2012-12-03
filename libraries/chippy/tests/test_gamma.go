//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: test_gamma.go
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

    screen, err := chippy.DefaultScreen()
    if err != nil {
        panic(err)
    }

    do := func(x float64) {
        fmt.Println(x)
        err := screen.SetGamma(float32(x))
        if err != nil {
            // Unable to set gamma
            panic(err.Error())
        }
        time.Sleep(1 * time.Millisecond)
        gamma, err := screen.Gamma()
        if err != nil {
            // Unable to get gamma
            panic(err.Error())
        }
        fmt.Println("Gamma is", gamma)
    }

    for x := 0; x < 3; x ++ {
        for i := 1.0; i <= 2.0; i += 0.01 {
            do(i)
        }
        for i := 2.0; i >= 0.0; i -= 0.01 {
            do(i)
        }
        for i := 0.0; i <= 1.0; i += 0.01 {
            do(i)
        }
        do(1.0)
    }

    do(0.3) // ensure restore works
    time.Sleep(1 * time.Second)
    //screen.SetAutoRestoreOriginalGamma(false)
    //do(1)
}

