//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: node.go
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

import "bytes"
import "sync"
import "fmt"

// Node represents a single object in the node tree
type Node struct{
    name string
    x, y, z float64       // Position
    rx, ry, rz float64    // Rotation
    sx, sy, sz float64    // Scale
    shx, shy, shz float64 // Shear
    hidden, depthTest bool
    access sync.RWMutex
}

// New returns a new initialized Node
func New(name string) *Node {
    n := new(Node)
    n.name = name
    return n
}

// String returns a string representation of n Node
func (n *Node) String() string {
    return fmt.Sprintf("Node(\"%s\")", n.name)
}

// Tree returns a string representation of n Node tree
func (n *Node) Tree() string {
    return n.treeIndent(0)
}

func (n *Node) treeIndent(indent int) string {
    spacing := bytes.Buffer{}
    for i := 0; i < indent; i++ {
        spacing.WriteString(" ")
    }
    spacing.WriteString(n.String())
    return spacing.String()
}

