// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

func (n *Node) addChild(child *Node) {
	n.access.Lock()
	defer n.access.Unlock()

	n.children = append(n.children, child)
}

func (n *Node) AddChild(child *Node) {
	n.addChild(child)
	child.setParent(n)
	n.checkForCircular()

	// Since parent is changing, we need to recursively clear the active props
	// of this node and all children nodes, as they can rely on the previous
	// parent.
	n.doRecursiveClearActiveProps()
}

func (n *Node) RemoveChild(child *Node) {
	n.access.Lock()
	defer n.access.Unlock()

	found := -1
	for i, c := range n.children {
		if c == child {
			found = i
			break
		}
	}
	if found != -1 {
		n.children[found] = nil
		n.children = append(n.children[:found], n.children[found+1:]...)

		child.setParent(nil)

		// Since parent is changing, we need to recursively clear the active props
		// of this node and all children nodes, as they can rely on the previous
		// parent.
		n.doRecursiveClearActiveProps()
	}
}

func (n *Node) doRemoveChildren() {
	// Get rid of all child nodes
	for _, child := range n.children {
		child.SetParent(nil)
	}

	for i, _ := range n.children {
		n.children[i] = nil
	}
	n.children = nil
}

func (n *Node) RemoveChildren() {
	n.access.Lock()
	defer n.access.Unlock()

	n.doRemoveChildren()
}

func (n *Node) HasChild(child *Node) bool {
	n.access.RLock()
	defer n.access.RUnlock()

	found := -1
	for i, c := range n.children {
		if c == child {
			found = i
			break
		}
	}
	return found != -1
}

func (n *Node) Children() []*Node {
	n.access.RLock()
	defer n.access.RUnlock()

	cpy := make([]*Node, len(n.children))
	copy(cpy, n.children)
	return cpy
}

// RecursiveChildren returns all children (and distant children) of this node,
// consider the following graph:
//
//  - Root
//    - A
//      - B
//        - C
//    - D
//      - E
//        - F
//    - G
//      - H
//        - I
//
// The resulting slice of an call to Root.RecursiveChildren() will look like:
//
//  [A, B, C, D, E, F, G, H, I]
//
func (n *Node) RecursiveChildren() []*Node {
	var children []*Node

	n.access.RLock()
	defer n.access.RUnlock()

	for _, child := range n.children {
		children = append(children, child)
		children = append(children, child.RecursiveChildren()...)
	}

	return children
}
