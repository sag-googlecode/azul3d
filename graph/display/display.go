//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: display.go
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
//       cumentation and/or other materials provided with the distribution.
//     * Neither the name of Lightpoke nor the
//       names of its contributors may be used to enrse or promote products
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

import "code.google.com/p/azul3d/graph/node"
import "sync"
import "fmt"

// Display represents an visual Display node
// Depending on the backend you are using, it may choose
// to use or ignore certain options that you specify
type Display struct {
    node.Node
    plat *platform
    title string
    width, height uint16
    x, y int16
    fullscreen, minimized, paused bool
    access sync.RWMutex
}

// New returns a new initialized Display node
func New(name string) *Display {
    d := Display{node.Node: *node.New(name)}
    d.plat = newPlatform()
    d.SetTitle("Display Title")
    d.SetSize(640, 480)
    d.SetFullscreen(false)
    d.SetMinimized(false)
    d.SetPaused(false)
    d.SetPos(-1, -1)

    // Note: there is an implied trust that you will never call any Node
    // functions in here otherwise you will hit an deadlock
    d.Node.StringCallback = func() string {
        d.access.RLock()
        defer d.access.RUnlock()
        return fmt.Sprintf("Display(\"%s\", title=\"%s\", size=%dx%d, pos=(%d,%d) fullscreen=%t, minimized=%t, paused=%t)", d.Name(), d.title, d.width, d.height, d.x, d.y, d.fullscreen, d.minimized, d.paused)
    }

    d.Node.Removers = append(d.Node.Removers, func() {
        d.access.RLock()
        defer d.access.RUnlock()
        d.plat.close()
    })
    return &d
}

