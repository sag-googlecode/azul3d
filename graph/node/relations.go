//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: relations.go
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

func (n *Node) AddChild(child *Node) {
    n.access.Lock()
    defer n.access.Unlock()
    n.children = append(n.children, child)
}

// RemoveChild removes the child node from this node's children,
// or if child is not a child, simply returns
func (n *Node) RemoveChild(child *Node) {
    n.access.Lock()
    defer n.access.Unlock()
    index := n.indexChild(child)
    if index == -1 {
        return
    }
    n.children = append(n.children[:index], n.children[index+1:]...)
}

// IsChild returns weather child is a child of this node
func (n *Node) IsChild(child *Node) bool {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.indexChild(child) != -1
}

// indexChild returns the index of the child node or -1 if it is not a child
func (n *Node) indexChild(child *Node) int {
    index := -1
    for i := 0; i < len(n.children); i++ {
        if n.children[i] == child {
            index = i
            break
        }
    }
    return index
}

// Children returns an slice of *Node that is the children of this node
func (n *Node) Children() []*Node {
    n.access.RLock()
    defer n.access.RUnlock()
    children := make([]*Node, len(n.children))
    copy(children, n.children)
    return children
}

func (n *Node) SetParent(parent *Node) {
    n.access.Lock()
    defer n.access.Unlock()
    n.parent = parent
}

// Parent returns the parent node of this node
func (n *Node) Parent() *Node {
    n.access.RLock()
    defer n.access.RUnlock()
    return n.parent
}

