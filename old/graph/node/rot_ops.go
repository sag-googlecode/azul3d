//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: rot_ops.go
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

// SetRotationX sets the rotation around the X axis of this Node
func (n *Node) SetRotationX(rx float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.rx = rx
}

// RotationX returns the rotation around the X axis of this Node
func (n *Node) RotationX() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.rx
}

// SetRotationY sets the rotation around the Y axis of this Node
func (n *Node) SetRotationY(ry float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.ry = ry
}

// RotationY returns the rotation around the Y axis of this Node
func (n *Node) RotationY() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.ry
}

// SetRotationZ sets the rotation around the Z axis of this Node
func (n *Node) SetRotationZ(rz float64) {
    n.access.Lock()
    defer n.access.Unlock()
    n.rz = rz
}

// RotationZ returns the rotation around the Z axis of this Node
func (n *Node) RotationZ() float64 {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.rz
}

// SetRotation sets the rotation around the x, y, and z axis of this Node
func (n *Node) SetRotation(rx, ry, rz float64) {
    n.SetRotationX(rx)
    n.SetRotationY(ry)
    n.SetRotationZ(rz)
}
