// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

var (
	PTransparency = NewProp("Transparency")
)

// SetTransparency sets the transparency mode of this node.
//
// The transparency mode of an node controls how it's transparent or partially
// transparent parts are drawn.
//
// The transparency mode must be valid (I.e. one of the predefined constants)
// or else a panic will occur.
func (n *Node) SetTransparency(mode TransparencyMode) {
	if !mode.Valid() {
		panic("SetRenderMode(): Invalid render mode specified!")
	}

	n.SetProp(PTransparency, mode)
}

// Transparency returns the transparency mode of this node or NoTransparency if
// there is no transparency on this node.
//
// The transparency mode of an node controls how it's transparent or partially
// transparent parts are drawn.
func (n *Node) Transparency() TransparencyMode {
	i, ok := n.Prop(PTransparency)
	if !ok {
		return NoTransparency
	}
	return i.(TransparencyMode)
}

// ActiveTransparency returns the active transparency of this node or
// NoTransparency if there is no transparency on this node.
func (n *Node) ActiveTransparency() TransparencyMode {
	i, ok := n.ActiveProp(PTransparency)
	if !ok {
		return NoTransparency
	}
	return i.(TransparencyMode)
}

// ClearTransparency clears the transparency mode of this node.
func (n *Node) ClearTransparency() {
	n.ClearProp(PTransparency)
}

// SetTransparencyForced specifies if the transparency of this node should be
// forced to be actively used instead of obeying the parent node wishes.
func (n *Node) SetTransparencyForced(forced bool) {
	n.SetPropForced(PTransparency, forced)
}

// TransparencyForced tells if the transparency of this node is forced as being
// actively used instead of obeying the parent node wishes.
func (n *Node) TransparencyForced() bool {
	return n.PropForced(PTransparency)
}
