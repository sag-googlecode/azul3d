//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: shear_ops.go
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

// SetShearX sets the shear along the X axis of this Node
func (n *Node) SetShearX(shx float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.shx = shx
}

// ShearX returns the shear along the X axis of this Node
func (n *Node) ShearX() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.shx
}

// SetShearY sets the shear along the Y axis of this Node
func (n *Node) SetShearY(shy float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.shy = shy
}

// ShearY returns the shear along the Y axis of this Node
func (n *Node) ShearY() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.shy
}

// SetShearZ sets the shear along the Z axis of this Node
func (n *Node) SetShearZ(shz float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.shz = shz
}

// ShearZ returns the shear along the Z axis of this Node
func (n *Node) ShearZ() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.shz
}


// SetShear sets the shear along the x, y, and z axis of this Node
func (n *Node) SetShear(shx, shy, shz float64) {
    n.SetShearX(shx)
    n.SetShearY(shy)
    n.SetShearZ(shz)
}

// Shear returns the shear along the x, y, and z axis of this Node
func (n *Node) Shear() (float64, float64, float64) {
    return n.ShearX(), n.ShearY(), n.ShearZ()
}
