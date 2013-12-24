// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"bytes"
	"io"
	"os"
	"sync"
)

func (n *Node) writeTree(out io.Writer, indention int) {
	var buf bytes.Buffer
	for i := 0; i < indention; i++ {
		buf.WriteByte(' ')
	}
	buf.WriteString(n.String())
	buf.WriteString("\n")
	buf.WriteTo(out)
	for _, child := range n.Children() {
		child.writeTree(out, indention+4)
	}
}

// WriteTree writes to the specified io.Writer, out, an string-representation of the scene graph
// tree represented by this node, and all children nodes of this node.
func (n *Node) WriteTree(out io.Writer) {
	n.writeTree(out, 0)
}

var (
	printLock sync.RWMutex
)

// PrintTree prints the tree (from WriteTree) into the standard output (os.Stdout).
func (n *Node) PrintTree() {
	printLock.Lock()
	defer printLock.Unlock()

	n.WriteTree(os.Stdout)
}

// Parents returns all possible parents, in order.
//
// Consider this graph (we are the node C here):
//
//  A
//      B
//          C
//
// The returned slice will look like:
//
//  [B A]
func (n *Node) Parents() []*Node {
	n.access.RLock()
	if n.parents != nil {
		defer n.access.RUnlock()
		return n.parents
	}

	// Exit read lock
	n.access.RUnlock()

	// We need to verify that no circular references exit while building this
	// slice of parents.
	var parents []*Node

	// We keep track of each parent we visit. If we visit an parent twice, then
	// it's an circular reference and we can safely panic.
	visited := make(map[*Node]bool)

	var locateParents func(of *Node)
	locateParents = func(of *Node) {
		_, visitedAlready := visited[of]
		if visitedAlready {
			panic(CircularErr)
		}
		visited[of] = true

		parent := of.Parent()

		if parent != nil {
			parents = append(parents, parent)
			locateParents(parent)
		}
	}
	locateParents(n)

	// Enter write lock
	n.access.Lock()
	defer n.access.Unlock()

	n.parents = parents
	return n.parents
}

// Top returns the top node of this scene graph.
//
// Consider the following graph:
//
//  A
//      B
//          C
//
// C.Top() would return A.
//
// B.Top() would return A.
//
// A.Top() would return nil.
//
func (n *Node) Top() *Node {
	parents := n.Parents()
	if len(parents) > 0 {
		return parents[len(parents)-1]
	}
	return n
}

func (n *Node) setParent(parent *Node) {
	n.access.RLock()
	if n.parent == parent {
		n.access.RUnlock()
		return
	}
	n.access.RUnlock()

	n.access.Lock()
	defer n.access.Unlock()

	// Assign new parent
	n.parent = parent
}

// SetParent specifies the parent node of this node, making this node a child
// of the parent node.
func (n *Node) SetParent(parent *Node) {
	if parent != nil {
		if !parent.addChild(n) {
			// We're already a child of the parent? So we're done.
			return
		}
	}
	n.setParent(parent)
	if parent != nil {
		n.checkForCircular()
	}

	// Since parent is changing, we need to recursively clear the active props
	// of this node and all children nodes, as they can rely on the previous
	// parent.
	n.recursiveClearActiveProps()
}

// Parent returns the parent node of this node.
func (n *Node) Parent() *Node {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.parent
}

// FindCommonParent finds a parent that both this node and the other node have
// in common and returns it, or returns nil if there is not common node.
//
// If the other node is nil; a panic will occur.
func (n *Node) FindCommonParent(other *Node) *Node {
	if other == nil {
		panic("FindCommonParent(): other node is nil!")
	}

	// What is the common parent between our node and our node? Obviously, our parent itself.
	if n == other {
		return n.Parent()
	}

	firstParents := n.Parents()
	secondParents := other.Parents()

	// Walk up graph
	for _, parent := range firstParents {
		// Compare this parent node with each of the other node's parents.
		for _, oParent := range secondParents {
			if parent == oParent {
				return parent
			}
		}
	}

	return nil
}
