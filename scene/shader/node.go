// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package shader

import (
	"code.google.com/p/azul3d/scene"
)

var (
	PShader = scene.NewProp("Shader")
)

// Set specifies the shader this node should have. The shader of the node will
// be used to shade this node's (and it's children node's) geoms.
func Set(n *scene.Node, s *Shader) {
	n.SetProp(PShader, s)
}

// Get returns the shader of this node, if it has one.
//
// If this node does not have a shader; ok will be false.
func Get(n *scene.Node) (s *Shader, ok bool) {
	i, ok := n.Prop(PShader)
	if ok {
		s = i.(*Shader)
	}
	return
}

// Active returns the active shader of this node and true, or returns nil and
// false if there is no active shader.
func Active(n *scene.Node) (s *Shader, ok bool) {
	i, ok := n.ActiveProp(PShader)
	if ok {
		s = i.(*Shader)
	}
	return
}

// Clear clears the shader of this node such that it has no shader.
func Clear(n *scene.Node) {
	n.ClearProp(PShader)
}

// SetForced specifies if the shader of this node should be forced to be
// actively used instead of obeying the parent node wishes.
func SetForced(n *scene.Node, forced bool) {
	n.SetPropForced(PShader, forced)
}

// Forced tells if the shader of this node is forced as being actively used
// instead of obeying the parent node wishes.
func Forced(n *scene.Node) bool {
	return n.PropForced(PShader)
}
