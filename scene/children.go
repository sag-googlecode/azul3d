// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

func (n *Node) addChild(child *Node) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.children == nil {
		n.children = make(map[*Node]bool)
	}
	n.children[child] = true
}

func (n *Node) AddChild(child *Node) {
	n.addChild(child)
	child.setParent(n)
	n.checkForCircular()
}

func (n *Node) RemoveChild(child *Node) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.children == nil {
		return
	}
	delete(n.children, child)
	child.setParent(nil)
}

func (n *Node) doRemoveChildren() {
	// Get rid of all child nodes
	for child, _ := range n.children {
		child.SetParent(nil)
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

	if n.children == nil {
		return false
	}
	_, ok := n.children[child]
	return ok
}

func (n *Node) Children() []*Node {
	n.access.RLock()
	defer n.access.RUnlock()

	if n.children == nil {
		return []*Node{}
	}

	children := make([]*Node, len(n.children))
	i := 0
	for child, _ := range n.children {
		children[i] = child
		i++
	}
	return children
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

	for child, _ := range n.children {
		children = append(children, child)
		children = append(children, child.RecursiveChildren()...)
	}

	return children
}
