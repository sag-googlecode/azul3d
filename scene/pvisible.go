// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

var (
	// The property used for storing the visibility status of a node.
	PVisible = NewProp("Visible")
)

// Show makes this node visible.
//
// Default: visible
func (n *Node) Show() {
	n.SetProp(PVisible, true)
}

// Hide makes this node hidden.
//
// Default: visible
func (n *Node) Hide() {
	n.SetProp(PVisible, false)
}

// Hidden determines the active visibility status of this node; as affected by
// parent node's visiblities, etc.
//
// Default: visible
func (n *Node) Hidden() bool {
	visible, ok := n.ActiveProp(PVisible)
	if !ok {
		return false
	}
	return !visible.(bool)
}

// SetShownThrough specifies weather or not this visibility of this node should
// be forced regardless of parent node wishes.
//
// Default: false
func (n *Node) SetShownThrough(shownThrough bool) {
	n.SetPropForced(PVisible, shownThrough)
}

// ShownThrough tells weather or not the visibility of this node is forced
// regardless of parent node wishes.
//
// Default: false
func (n *Node) ShownThrough() bool {
	return n.PropForced(PVisible)
}
