// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

import (
	"code.google.com/p/azul3d/scene"
	"sync"
)

var (
	textureMap     = scene.NewProp("textureMap")
	textureMapLock = scene.NewProp("textureMapLock")
)

func getLock(n *scene.Node) *sync.RWMutex {
	l, ok := n.Tag(textureMapLock)
	if !ok {
		newLock := new(sync.RWMutex)
		n.SetTag(textureMapLock, newLock)
		return newLock
	}
	return l.(*sync.RWMutex)
}

func getTextures(n *scene.Node) map[*Layer]Type {
	i, ok := n.Tag(textureMap)
	if !ok {
		newMap := make(map[*Layer]Type)
		n.SetTag(textureMap, newMap)
		return newMap
	}
	return i.(map[*Layer]Type)
}

// Set stores the given texture into the given texture layer of this node.
//
// The texture must be a valid texture type or else a panic will occur.
//
// If the layer or texture type is nil, a panic will occur.
func Set(n *scene.Node, l *Layer, t Type) {
	if l == nil {
		panic("Set(): Layer is nil!")
	}
	if t == nil {
		panic("Set(): Texture is nil!")
		return
	}
	if !IsValid(t) {
		panic("Set(): Invalid texture type.")
	}

	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	textures := getTextures(n)
	textures[l] = t
}

// Get returns the texture currently stored in the given texture layer.
func Get(n *scene.Node, l *Layer) (t Type, ok bool) {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	textures := getTextures(n)
	t, ok = textures[l]
	return
}

// Remove removes the texture currently stored in the given texture layer.
func Remove(n *scene.Node, l *Layer) {
	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	textures := getTextures(n)
	delete(textures, l)
}

// Textures returns a copy of the internal map used to hold texture layers and
// their associated textures.
func Textures(n *scene.Node) map[*Layer]Type {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	textures := getTextures(n)
	c := make(map[*Layer]Type, len(textures))
	for l, t := range textures {
		c[l] = t
	}
	return c
}
