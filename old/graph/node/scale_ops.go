//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: scale_ops.go
// Created by: Stephen Gutekanst, 11/19/12
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
//     * Neither the name of the <organization> nor the
//       names of its contributors may be used to endorse or promote products
//       derived from n software without specific prior written permission.
//
// n SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF n
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package node

// SetScaleX sets the scale on the X axis of this Node
func (n *Node) SetScaleX(sx float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.sx = sx
}

// ScaleX returns the scale on the X axis of this Node
func (n *Node) ScaleX() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.sx
}

// SetScaleY sets the scale on the Y axis of this Node
func (n *Node) SetScaleY(sy float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.sy = sy
}

// ScaleY returns the scale on the Y axis of this Node
func (n *Node) ScaleY() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.sy
}

// SetScaleZ sets the scale on the Z axis of this Node
func (n *Node) SetScaleZ(sz float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.sz = sz
}

// ScaleZ returns the scale on the Z axis of this Node
func (n *Node) ScaleZ() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.sz
}


// SetScale sets the scale on the x, y, and z axis of this Node
func (n *Node) SetScale(sx, sy, sz float64) {
    n.SetScaleX(sx)
    n.SetScaleY(sy)
    n.SetScaleZ(sz)
}

// Scale returns the scale on the x, y, and z axis of this Node
func (n *Node) Scale() (float64, float64, float64) {
    return n.ScaleX(), n.ScaleY(), n.ScaleZ()
}

