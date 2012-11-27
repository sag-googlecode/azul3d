//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: pos_ops.go
// Created by: Stephen Gutekanst, 11/18/12
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

// SetX sets the X position of this Node
func (n *Node) SetX(x float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.x = x
}

// X returns the X position of this Node
func (n *Node) X() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.x
}

// SetY sets the Y position of this Node
func (n *Node) SetY(y float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.y = y
}

// Y returns the Y position of this Node
func (n *Node) Y() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.y
}

// SetZ sets the Z position of this Node
func (n *Node) SetZ(z float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.z = z
}

// Z returns the Z position of this Node
func (n *Node) Z() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.z
}

// SetPos sets the x, y, and z position of this Node
func (n *Node) SetPos(x, y, z float64) {
    n.SetX(x)
    n.SetY(y)
    n.SetZ(z)
}

// Pos returns the x, y, and z position of this Node
func (n *Node) Pos() (float64, float64, float64) {
    return n.X(), n.Y(), n.Z()
}
