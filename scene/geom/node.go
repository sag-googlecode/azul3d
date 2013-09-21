// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/color"
	"sync"
)

var (
	meshSlice     = scene.NewProp("meshSlice")
	meshSliceLock = scene.NewProp("meshSliceLock")
)

func getLock(n *scene.Node) *sync.RWMutex {
	l, ok := n.Tag(meshSliceLock)
	if !ok {
		newLock := new(sync.RWMutex)
		n.SetTag(meshSliceLock, newLock)
		return newLock
	}
	return l.(*sync.RWMutex)
}

func getMeshes(n *scene.Node) []*Mesh {
	i, ok := n.Tag(meshSlice)
	if !ok {
		newSlice := make([]*Mesh, 0)
		n.SetTag(meshSlice, newSlice)
		return newSlice
	}
	return i.([]*Mesh)
}

func appendMesh(n *scene.Node, m *Mesh) {
	meshes := getMeshes(n)
	meshes = append(meshes, m)
	n.SetTag(meshSlice, meshes)
}

func removeMesh(n *scene.Node, index int) {
	meshes := getMeshes(n)
	meshes = append(meshes[:index], meshes[index+1:]...)
	n.SetTag(meshSlice, meshes)
}

func makeCopy(n *scene.Node, deep bool) {
	_, ok := n.Tag(meshSlice)
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
	n.SetTag(meshSlice, cpy)

	return
}

func Copy(n *scene.Node) {
	makeCopy(n, false)
}

func DeepCopy(n *scene.Node) {
	makeCopy(n, true)
}

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

func Meshes(n *scene.Node) []*Mesh {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	meshes := getMeshes(n)
	cpy := make([]*Mesh, len(meshes))
	copy(cpy, meshes)
	return cpy
}

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
