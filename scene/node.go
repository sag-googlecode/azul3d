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

type Node struct {
	access sync.RWMutex

	id                     uint
	name                   string
	parent                 *Node
	parents                []*Node
	children               map[*Node]bool
	tags                   map[interface{}]interface{}
	props, activePropCache map[interface{}]interface{}
	forcedProps            map[interface{}]bool
	transform              *Transform
}

func (n *Node) String() string {
	x, y, z := n.Pos()
	rx, ry, rz := n.Rot()
	sx, sy, sz := n.Scale()
	return fmt.Sprintf("Node(%q, Pos=[%v, %v, %v], Rot=[%v, %v, %v], Scale=[%v, %v, %v])", n.Name(), x, y, z, rx, ry, rz, sx, sy, sz)
}

func (n *Node) Copy() *Node {
	n.access.RLock()
	defer n.access.RUnlock()

	var children map[*Node]bool
	if n.children != nil {
		children = make(map[*Node]bool, len(n.children))
		for child, _ := range n.children {
			childCopy := child.Copy()
			children[childCopy] = true
		}
	}

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
	for child, _ := range children {
		child.SetParent(copy)
	}

	return copy
}

func (n *Node) doDetatch() {
	n.doRemoveChildren()

	// Inform of parents changing
	n.doRecursiveClearActiveProps()
}

func (n *Node) Detatch() {
	n.access.Lock()
	defer n.access.Unlock()

	n.doDetatch()
}

func (n *Node) Destroy() {
	parent := n.Parent()
	if parent != nil {
		parent.RemoveChild(n)
	}

	n.access.Lock()
	defer n.access.Unlock()

	// Once we detatch this node, it will have no children, so defer destroying
	// the child node.
	for child, _ := range n.children {
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
