// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"fmt"
	"sync"
)

// Node represents a single node within a 3D scene graph. A node may have a
// parent and/or several children nodes related to it.
//
// A node is a generic container type which may be used to hold several
// properties related to the node as well.
//
// A node holds a three dimensional transformation, although it is not
// typically manipulated directly.
type Node struct {
	access sync.RWMutex

	name            string
	parent          *Node
	parents         []*Node
	children        []*Node
	props           map[interface{}]interface{}
	activePropCache map[interface{}]mapLookupPair
	forcedProps     map[interface{}]bool
	transform       *Transform
}

// String returns a string representation of this node.
func (n *Node) String() string {
	x, y, z := n.Pos()
	rx, ry, rz := n.Rot()
	sx, sy, sz := n.Scale()
	return fmt.Sprintf("Node(%q, Pos=[%v, %v, %v], Rot=[%v, %v, %v], Scale=[%v, %v, %v])", n.Name(), x, y, z, rx, ry, rz, sx, sy, sz)
}

// Detatch detatches this node from the scene graph; removing both it's parent
// node and removing all of it's child nodes.
//
// A detatched node may be later re-added to a scene graph using the SetParent,
// or AddChild methods, for instance.
func (n *Node) Detatch() {
	n.Parent().RemoveChild(n)

	n.access.Lock()
	defer n.access.Unlock()

	// Inform children nodes of parent node changing.
	n.doRecursiveClearActiveProps()

	// Remove children nodes.
	n.doRemoveChildren()
}

// New returns a new node with the given name.
//
// A node's name is used for reference when printing nodes, etc. A good node
// name is one which represents the node itself, e.g. "Beer" would be a good
// node name for a node which will have a beer model attached to it.
func New(name string) *Node {
	return &Node{
		name:      name,
		transform: new(Transform),
	}
}
