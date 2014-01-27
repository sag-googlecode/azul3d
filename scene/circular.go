// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build !fast

package scene

const (
	whileFindingTop = false
	whileTraversing = true
)

func (n *Node) checkForCircular() {
	// A map of visited nodes, where the boolean value is either
	// whileFindingTop or whileTraversing (depending on where the node was
	// visited).
	visited := make(map[*Node]bool)

	// Note: we cannot use n.Top() because it would mean we have to make the
	// Top() implementation aware of circular graphs which would be more
	// expensive when circular testing is off.
	var findTop func(n *Node) *Node
	findTop = func(n *Node) *Node {
		// Check if we have visited this node already before.
		v, ok := visited[n]
		if ok && v == whileFindingTop {
			panic(CircularErr)
		}
		visited[n] = whileFindingTop

		// If the node has a parent, then we should return findTop(parent) this
		// enters the recursive search.
		parent := n.Parent()
		if parent != nil {
			return findTop(parent)
		}

		// Node does not have a parent, this must be the top node then.
		return n
	}

	// Find the top node.
	top := findTop(n)

	// Traverse the whole scene starting at the top node to see if there are
	// any nodes that we visit twice, and if we do then panic.
	top.Traverse(func(i int, n *Node) bool {
		// Check if we have visited this node already before.
		v, ok := visited[n]
		if ok && v == whileTraversing {
			panic(CircularErr)
		}
		visited[n] = whileTraversing

		// Continue the traversal.
		return true
	})
}
