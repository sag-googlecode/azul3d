// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package transparency

import (
	"azul3d.org/scene"
)

var (
	// The property used for storing the transparency mode of a node.
	PMode = scene.NewProp("Mode")
)

// Set sets the transparency mode of this node.
//
// The transparency mode of an node controls how it's transparent or partially
// transparent parts are drawn.
//
// The transparency mode must be valid (I.e. one of the predefined constants)
// or else a panic will occur.
//
// If the alpha mode to be set is AlphaBlend you might also want to place the
// node inside a back-to-front sorted bucket, such as bucket.Transparent.
func Set(n *scene.Node, m ModeType) {
	if !m.Valid() {
		panic("Set(): Invalid transparency mode specified!")
	}

	n.SetProp(PMode, m)
}

// Mode returns the transparency mode of this node or None if there is no
// transparency on this node.
//
// The transparency mode of an node controls how it's transparent or partially
// transparent parts are drawn.
func Mode(n *scene.Node) ModeType {
	i, ok := n.Prop(PMode)
	if !ok {
		return None
	}
	return i.(ModeType)
}

// ActiveMode returns the active transparency of this node or None if there is
// no transparency on this node.
func ActiveMode(n *scene.Node) ModeType {
	i, ok := n.ActiveProp(PMode)
	if !ok {
		return None
	}
	return i.(ModeType)
}

// ClearMode clears the transparency mode of this node.
func ClearMode(n *scene.Node) {
	n.ClearProp(PMode)
}

// SetModeForced specifies if the transparency of this node should be forced to
// be actively used instead of obeying the parent node wishes.
func SetModeForced(n *scene.Node, forced bool) {
	n.SetPropForced(PMode, forced)
}

// ModeForced tells if the transparency of this node is forced as being
// actively used instead of obeying the parent node wishes.
func ModeForced(n *scene.Node) bool {
	return n.PropForced(PMode)
}
