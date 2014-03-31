// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"azul3d.org/v1/scene"
	"azul3d.org/v1/scene/color"
	"sync"
)

var (
	// The property for storing the mesh slice of a node.
	PMeshSlice = scene.NewProp("MeshSlice")

	// The property for storing the mesh slice lock of a node.
	PMeshSliceLock = scene.NewProp("MeshSliceLock")
)

func getLock(n *scene.Node) *sync.RWMutex {
	l, ok := n.Prop(PMeshSliceLock)
	if !ok {
		newLock := new(sync.RWMutex)
		n.SetProp(PMeshSliceLock, newLock)
		return newLock
	}
	return l.(*sync.RWMutex)
}

func getMeshes(n *scene.Node) []*Mesh {
	i, ok := n.Prop(PMeshSlice)
	if !ok {
		newSlice := make([]*Mesh, 0)
		n.SetProp(PMeshSlice, newSlice)
		return newSlice
	}
	return i.([]*Mesh)
}

func appendMesh(n *scene.Node, m *Mesh) {
	meshes := getMeshes(n)
	meshes = append(meshes, m)
	n.SetProp(PMeshSlice, meshes)
}

func removeMesh(n *scene.Node, index int) {
	meshes := getMeshes(n)
	meshes = append(meshes[:index], meshes[index+1:]...)
	n.SetProp(PMeshSlice, meshes)
}

func makeCopy(n *scene.Node, deep bool) {
	_, ok := n.Prop(PMeshSlice)
	if !ok {
		// Not an mesh
		return
	}

	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	meshes := getMeshes(n)
	cpy := make([]*Mesh, len(meshes))
	for index, mesh := range meshes {
		if deep {
			cpy[index] = mesh.Copy()
		} else {
			cpy[index] = mesh
		}
	}
	n.SetProp(PMeshSlice, cpy)

	return
}

// Copy copies in-place the meshes used by this node. This does not make copy
// the actual meshes themselves, just the internal "mesh" object which stores
// meshes.
func Copy(n *scene.Node) {
	makeCopy(n, false)
}

// DeepCopy works just like Copy(), except it makes copies of all individual
// meshes themselves, which may be a very expensive operation!
func DeepCopy(n *scene.Node) {
	makeCopy(n, true)
}

// Add adds the specified mesh to the specified node.
//
// If the node already has the specified mesh, this function is no-op.
func Add(n *scene.Node, m *Mesh) {
	if m == nil || !m.Valid() {
		panic("Add(): Mesh is invalid or nil!")
	}

	access := getLock(n)

	// See if we already have the mesh.
	access.RLock()
	alreadyHaveMesh := false
	for _, mesh := range getMeshes(n) {
		if mesh == m {
			alreadyHaveMesh = true
			break
		}
	}
	access.RUnlock()

	if alreadyHaveMesh {
		return
	}

	// We don't have the mesh yet, append it
	access.Lock()
	defer access.Unlock()

	appendMesh(n, m)
}

// Remove removes the specified mesh from the specified node.
//
// If the node does not have the specified mesh, this function is no-op.
func Remove(n *scene.Node, m *Mesh) {
	if m == nil {
		return
	}

	access := getLock(n)

	// See if we actually have the mesh.
	access.RLock()
	meshIndex := -1
	for index, mesh := range getMeshes(n) {
		if mesh == m {
			meshIndex = index
			break
		}
	}
	access.RUnlock()

	if meshIndex == -1 {
		// We don't have the mesh.
		return
	}

	// We have the mesh, remove it.
	access.Lock()
	defer access.Unlock()

	removeMesh(n, meshIndex)
}

// Has tells if the node has the specified mesh.
func Has(n *scene.Node, m *Mesh) bool {
	if m == nil {
		return false
	}

	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	// See if we actually have the mesh
	for _, mesh := range getMeshes(n) {
		if mesh == m {
			return true
		}
	}
	return false
}

// Meshes returns a slice of all the meshes that the specified node has.
func Meshes(n *scene.Node) []*Mesh {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	meshes := getMeshes(n)
	cpy := make([]*Mesh, len(meshes))
	copy(cpy, meshes)
	return cpy
}

// BakeColors takes the active color and color scales of the specified node and
// all children below it, and 'bakes' them into the vertice colors of the
// meshes in these nodes.
func BakeColors(n *scene.Node, recursive bool) {
	n.Traverse(func(i int, n *scene.Node) bool {
		c, ok := color.Active(n)
		if !ok {
			c = color.None
		}
		cs, ok := color.ActiveScale(n)
		if !ok {
			cs = color.None
		}

		for _, m := range Meshes(n) {
			if m.IsHidden() {
				continue
			}

			m.BakeColors(c, cs)
		}

		return recursive
	})
}
