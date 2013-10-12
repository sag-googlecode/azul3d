// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

var (
	// The property used for storing the sorter of a node.
	PSorter = NewProp("Sorter")
)

// SetSorter specifies the sorter which describes how this node should be
// sorted for rendering.
func (n *Node) SetSorter(sorter *Sorter) {
	n.SetProp(PSorter, sorter)
}

// Sorter returns the sorter which describes how this node should be sorted for
// rendering and true, or nil and false if this node has no sorter.
func (n *Node) Sorter() (sorter *Sorter, ok bool) {
	i, ok := n.Prop(PSorter)
	if ok {
		sorter = i.(*Sorter)
	}
	return
}

// ActiveSorter returns the active sorter of this node, as affected by parent
// node sorter values.
func (n *Node) ActiveSorter() (sorter *Sorter, ok bool) {
	i, ok := n.ActiveProp(PSorter)
	if ok {
		sorter = i.(*Sorter)
	}
	return
}

// ClearSorter clears the sorter value of this node.
func (n *Node) ClearSorter() {
	n.ClearProp(PSorter)
}

// SetSorterForced specifies if the sorter value of this node should be forced
// to be actively used instead of obeying the parent node wishes.
func (n *Node) SetSorterForced(forced bool) {
	n.SetPropForced(PSorter, forced)
}

// SorterForced tells if the sorter value of this node is forced as being
// actively used instead of obeying the parent node wishes.
func (n *Node) SorterForced() bool {
	return n.PropForced(PSorter)
}
