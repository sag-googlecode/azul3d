// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"fmt"
	"sync"
)

var (
	idCounter       uint
	idCounterAccess sync.Mutex
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

	id              uint
	name            string
	parent          *Node
	parents         []*Node
	children        []*Node
	tags            map[interface{}]interface{}
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

// Copy returns an exact copy of this node.
func (n *Node) Copy() *Node {
	n.access.RLock()
	defer n.access.RUnlock()

	children := make([]*Node, len(n.children))
	copy(children, n.children)

	var tags map[interface{}]interface{}
	if n.tags != nil {
		tags = make(map[interface{}]interface{}, len(n.tags))
		for tag, value := range n.tags {
			tags[tag] = value
		}
	}

	var props map[interface{}]interface{}
	if n.props != nil {
		props = make(map[interface{}]interface{}, len(n.props))
		for tag, value := range n.props {
			props[tag] = value
		}
	}

	copy := &Node{
		id:        n.id,
		name:      n.name,
		children:  children,
		tags:      tags,
		props:     props,
		transform: n.transform.Copy(),
	}

	// In an Copy() we don't preserve parents, but we do for all our children (
	// and distant ones, too).
	for _, child := range children {
		child.SetParent(copy)
	}

	return copy
}

func (n *Node) doDetatch() {
	n.doRemoveChildren()

	n.parent = nil

	// Inform of parents changing
	n.doRecursiveClearActiveProps()
}

// Detatch detatches this node from the scene graph; removing both it's parent
// node and removing all of it's child nodes.
//
// A detatched node may be later re-added to a scene graph using the SetParent,
// or AddChild methods, for instance.
func (n *Node) Detatch() {
	n.access.Lock()
	defer n.access.Unlock()

	n.doDetatch()
}

// Destroy destroys this node, and all of it's children. Destroy should be
// called on a node once you are completely done using it and all of it's
// children nodes.
//
// This function plays a crucial role in the memory management of programs and
// it is dire that you understand it properly.
//
// At the core, scene graphs are just a bunch of nodes who hold parent and
// child relationships. Consider a single parent, with a single child node,
// like so:
//
//  *Parent
//      *Child
//
// Now consider that the parent node holds a reference to the child node, and
// likewise the child node holds a reference to the parent node. If you are
// experienced with programming in garbage collected languages, like Go, then
// you will immedietly see the problem. You have a memory leak.
//
// The solution to this is to remove the reference to each other, and to do
// this you could simply use the Detatch() method, which does exactly this.
//
// However yet still in more complex scene graphs, there remains a problem:
//
//  *Node1
//      *Node2
//          *Node3
//
// If we were to simply Detatch() Node1 in the above example, Node2 and Node3
// still hold a circular reference to each other and cannot be garbage
// collected.
//
// Because of this, you would need to recursively Detatch() every node in the
// above graph.
//
// In addition, since the Node type is a generic container type, which might
// hold references to other things, like shaders, textures, etc, simply
// detatching a node is not enough.
//
// To remedy all of this easilly, there is a very simply function, Destroy().
//
// When a root node is destroyed, all child nodes are also destroyed. The
// references between nodes are broken, allowing them to be garbage collected,
// and the properties of the node are wiped clean, meaning shaders, textures,
// and other resources tied to a node may also be garbage collected.
func (n *Node) Destroy() {
	parent := n.Parent()
	if parent != nil {
		parent.RemoveChild(n)
	}

	n.access.Lock()
	defer n.access.Unlock()

	// Once we detatch this node, it will have no children, so defer destroying
	// the child node.
	for _, child := range n.children {
		defer child.Destroy()
	}

	// Detatch this node (making it have no parent and no children).
	n.doDetatch()

	// Clear tags
	n.tags = nil

	// Clear props
	n.props = nil

	// Clear active property cache
	n.activePropCache = nil

	// Clear forced props
	n.forcedProps = nil
}

// New returns a new node with the given name.
//
// A node's name is used for reference when printing nodes, etc. A good node
// name is one which represents the node itself, e.g. "Beer" would be a good
// node name for a node which will have a beer model attached to it.
func New(name string) *Node {
	idCounterAccess.Lock()
	id := idCounter
	idCounter++
	idCounterAccess.Unlock()

	return &Node{
		id:        id,
		name:      name,
		transform: new(Transform),
	}
}
