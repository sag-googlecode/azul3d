// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

func (n *Node) traverse(callback func(index int, current *Node) bool) bool {
	cont := callback(0, n)
	if !cont {
		return false
	}

	for _, child := range n.Children() {
		cont = child.traverse(callback)
		if !cont {
			return false
		}
	}
	return cont
}

// Traverse allows you to traverse the scene graph in an generic easy-to-use way; it works by
// calling the specified function at each node who makes up the scene graph.
//
// The callback function should return true to continue the traversal (onto the next node) or false
// which means the traversal should be stopped (no longer invoke the callback).
//
//  Root
//      A
//          B
//      C
//          D
//      E
//          F
//          G
//          H
//              I
//              J
//
// For instance; the above scene would invoke the callback function, as follows, given that it was
// called using 'Root.Traverse(...)':
//
//  each(0, Root)
//  each(0, A)
//  each(0, B)
//  each(0, C)
//  each(0, D)
//  each(0, E)
//  each(0, F)
//  each(1, G)
//  each(2, H)
//  each(0, I)
//  each(1, J)
//
func (n *Node) Traverse(callback func(index int, current *Node) bool) {
	n.traverse(callback)
}
