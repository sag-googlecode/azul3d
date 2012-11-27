//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: node.go
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
package node

import "bytes"
import "sync"
import "fmt"

type Node struct {
    name string
    parent *Node
    children []*Node
    StringCallback func() string
    access sync.RWMutex
    Removers []func()
}

func New(name string) *Node {
    n := Node{}
    n.name = name
    return &n
}

func (n *Node) String() string {
    n.access.RLock()
    defer n.access.RUnlock()
    if n.StringCallback != nil {
        return n.StringCallback()
    }
    return fmt.Sprintf("Node(\"%s\")", n.name)
}

func (n *Node) Tree() string {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.treeIndent(0)
}

func (n *Node) treeIndent(indent uint8) string {
    bb := bytes.Buffer{}
    for i := 0; i < int(indent); i++ {
        bb.WriteString("  ")
    }
 
   bb.WriteString(n.String())
    for i := 0; i < len(n.children); i++ {
        bb.WriteString("\n")
        bb.WriteString(n.children[i].treeIndent(indent + 1))
    }
    return bb.String()
}

func (n *Node) Remove() {
    n.access.RLock()
    defer n.access.RUnlock()
    for i := 0; i < len(n.Removers); i++ {
        n.Removers[i]()
    }

}
