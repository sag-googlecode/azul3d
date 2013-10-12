// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

// New creates and returns a new node with the specified name, and attaches it
// as a child of this node. Short hand for:
//
//  child := New(name)
//  child.SetParent(n)
//
func (n *Node) New(name string) *Node {
	child := New(name)
	child.SetParent(n)
	return child
}
