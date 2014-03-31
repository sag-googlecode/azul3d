// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package color

import (
	"azul3d.org/v1/scene"
)

var (
	// Color node property
	PColor = scene.NewProp("Color")

	// Color Scale node property
	PScale = scene.NewProp("ColorScale")

	// Default solid-white color
	Default = New(1, 1, 1, 1)

	// None represents a None / nil / Invalid color.
	None = Color{-1, -1, -1, -1}
)

// Set specifies the color the node should have. The color of the node will
// replace existing per-vertex colors with the specified color.
func Set(n *scene.Node, c Color) {
	n.SetProp(PColor, c)
}

// Get returns the color of this node, if it has one.
//
// If this node does not have a color; ok will be false.
func Get(n *scene.Node) (c Color, ok bool) {
	i, ok := n.Prop(PColor)
	if ok {
		c = i.(Color)
	}
	return
}

// Active returns the active color of this node and true, or returns ok=false
// if there is no active color.
func Active(n *scene.Node) (c Color, ok bool) {
	i, ok := n.ActiveProp(PColor)
	if ok {
		c = i.(Color)
	}
	return
}

// Clear clears the color of this node such that it has 'no color', as such it
// will use per-vertex colors and child nodes below this one will user their
// own respective vertex colors.
//
// This is very different from setting the default color on this node, as with
// the default color (instead of 'no color') the vertex colors of this node and
// nodes below this one will be *replaced* with the default color.
func Clear(n *scene.Node) {
	n.ClearProp(PColor)
}

// SetForced specifies if the color of this node should be forced to be
// actively used instead of obeying the parent node wishes.
func SetForced(n *scene.Node, forced bool) {
	n.SetPropForced(PColor, forced)
}

// Forced tells if the color of this node is forced as being actively used
// instead of obeying the parent node wishes.
func Forced(n *scene.Node) bool {
	return n.PropForced(PColor)
}

// SetScale specifies the color scale the node should have. The vertex colors
// of this node will be multiplied against this scale.
func SetScale(n *scene.Node, c Color) {
	n.SetProp(PScale, c)
}

// Scale returns the color scale of the specified node.
//
// If this node does not have a color scale; ok will be false.
func Scale(n *scene.Node) (c Color, ok bool) {
	i, ok := n.Prop(PScale)
	if ok {
		c = i.(Color)
	}
	return
}

// ActiveScale returns the active color scale of this node and true, or returns
// ok=false if there is no active color scale.
func ActiveScale(n *scene.Node) (c Color, ok bool) {
	i, ok := n.ActiveProp(PScale)
	if ok {
		c = i.(Color)
	}
	return
}

// ClearScale clears the color scale of the specified node.
func ClearScale(n *scene.Node) {
	n.ClearProp(PScale)
}

// SetScaleForced specifies if the color scale of this node should be forced
// to be actively used instead of obeying the parent node wishes.
func SetScaleForced(n *scene.Node, forced bool) {
	n.SetPropForced(PScale, forced)
}

// ScaleForced tells if the color scale of this node is forced as being
// actively used instead of obeying the parent node wishes.
func ScaleForced(n *scene.Node) bool {
	return n.PropForced(PScale)
}
