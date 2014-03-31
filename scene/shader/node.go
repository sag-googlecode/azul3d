// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package shader

import (
	"azul3d.org/v1/scene"
	"sync"
)

var (
	// The property for storing the shader object of a node.
	PShader = scene.NewProp("Shader")

	// The property for storing the shader inputs map RWMutex of a node.
	PShaderInputsAccess = scene.NewProp("ShaderInputsAccess")

	// The property for storing the shader inputs map of a node.
	PShaderInputs = scene.NewProp("ShaderInputs")
)

// Set specifies the shader this node should have. The shader of the node will
// be used to shade this node's (and it's children node's) geoms.
func Set(n *scene.Node, s *Shader) {
	n.SetProp(PShader, s)
}

// Get returns the shader of this node, if it has one, otherwise this function
// returns nil.
func Get(n *scene.Node) *Shader {
	i, ok := n.Prop(PShader)
	if ok {
		return i.(*Shader)
	}
	return nil
}

// Active returns the active shader of this node, or returns nil if there is no
// active shader.
func Active(n *scene.Node) *Shader {
	i, ok := n.ActiveProp(PShader)
	if ok {
		return i.(*Shader)
	}
	return nil
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

func getLock(n *scene.Node) *sync.RWMutex {
	i, ok := n.Prop(PShaderInputsAccess)
	if !ok {
		i = new(sync.RWMutex)
	}
	return i.(*sync.RWMutex)
}

// SetInput sets the named shader input on the node to the given value.
//
// If the value is not of a valid type (see the IsValidType() function), then
// a panic will occur.
func SetInput(n *scene.Node, name string, value interface{}) {
	if !IsValidType(value) {
		panic("SetInput(): Input value is of an invalid type!")
	}

	lock := getLock(n)
	lock.Lock()
	defer lock.Unlock()

	i, ok := n.Prop(PShaderInputs)
	if !ok {
		i = make(map[string]interface{}, 1)
	}
	inputs := i.(map[string]interface{})

	inputs[name] = value
	n.SetProp(PShaderInputs, inputs)
}

// Input returns the value of the named shader input and ok=true, or returns
// nil and ok=false if there is no input with the given name.
func Input(n *scene.Node, name string) (value interface{}, ok bool) {
	lock := getLock(n)
	lock.RLock()
	defer lock.RUnlock()

	i, hasInputs := n.Prop(PShaderInputs)
	if !hasInputs {
		return
	}
	inputs := i.(map[string]interface{})

	value, ok = inputs[name]
	return
}

// Inputs returns a map of shader input names and their corrisponding values
// for the given node.
func Inputs(n *scene.Node) map[string]interface{} {
	lock := getLock(n)
	lock.RLock()
	defer lock.RUnlock()

	i, hasInputs := n.Prop(PShaderInputs)
	if !hasInputs {
		return make(map[string]interface{}, 0)
	}
	inputs := i.(map[string]interface{})

	cpy := make(map[string]interface{}, len(inputs))
	for k, v := range inputs {
		cpy[k] = v
	}
	return cpy
}
