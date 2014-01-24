// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

import (
	"azul3d.org/scene"
	"sort"
	"sync"
)

var (
	// A slice of texture layers that the node uses. Each texture is stored
	// uniquely as it's own property where the key is the *Layer and the value
	// is the *Texture, but to get all of the active textures on a node we must
	// maintain this list.
	PLayers = scene.NewProp("Layers")

	// Stores the *sync.RWMutex used for synchronization of the slice stored in
	// the PLayers property above.
	PLayersLock = scene.NewProp("LayersLock")
)

func getLayersLock(n *scene.Node) *sync.RWMutex {
	l, ok := n.Prop(PLayersLock)
	if !ok {
		newLock := new(sync.RWMutex)
		n.SetProp(PLayersLock, newLock)
		return newLock
	}
	return l.(*sync.RWMutex)
}

func getLayers(n *scene.Node) []*Layer {
	i, ok := n.Prop(PLayers)
	if !ok {
		return nil
	}
	return i.([]*Layer)
}

func getActiveLayers(n *scene.Node) []*Layer {
	i, ok := n.ActiveProp(PLayers)
	if !ok {
		return nil
	}
	return i.([]*Layer)
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
	}
	if !IsValid(t) {
		panic("Set(): Invalid texture type.")
	}

	access := getLayersLock(n)
	access.Lock()
	defer access.Unlock()

	// Store the layer.
	updatedLayers := append(getLayers(n), l)
	n.SetProp(PLayers, updatedLayers)

	// Pair the layer to the texture value.
	n.SetProp(l, t)
}

// Get returns the texture currently stored in the given texture layer.
//
// Returns t=nil, ok=false in the event that there is no texture for the given
// layer on the specified node.
func Get(n *scene.Node, l *Layer) (t Type, ok bool) {
	i, ok := n.Prop(l)
	if !ok {
		return nil, false
	}
	return i.(Type), true
}

// Active returns the active texture for the specified node currently stored in
// the given texture layer.
//
// Returns t=nil, ok=false in the event that there is no active texture for the
// given layer on the specified node.
func Active(n *scene.Node, l *Layer) (t Type, ok bool) {
	i, ok := n.ActiveProp(l)
	if !ok {
		return nil, false
	}
	return i.(Type), true
}

// Remove removes the texture currently stored in the given texture layer from
// the specified node.
func Remove(n *scene.Node, l *Layer) {
	// Remove the pairing between the layer and texture value.
	n.ClearProp(l)

	access := getLayersLock(n)
	access.Lock()
	defer access.Unlock()

	// Search through the layers slice for the given layer's index.
	layers := getLayers(n)
	found := -1
	for index, existingLayer := range layers {
		if existingLayer == l {
			found = index
			break
		}
	}
	if found != -1 {
		// We found the layer we want to remove, so remove it from the slice
		// and update the property.
		layers = append(layers[:found], layers[found:]...)
		n.SetProp(PLayers, layers)
	}
}

// Textures returns a map of texture layer's to their respective textures.
func Textures(n *scene.Node) map[*Layer]Type {
	access := getLayersLock(n)
	access.RLock()
	defer access.RUnlock()

	layers := getLayers(n)
	m := make(map[*Layer]Type, len(layers))
	for _, layer := range layers {
		tex, ok := Get(n, layer)
		if ok {
			m[layer] = tex
		}
	}
	return m
}

// ActiveTextures returns a map of texture layer's to their respective active
// textures on the given node.
func ActiveTextures(n *scene.Node) map[*Layer]Type {
	access := getLayersLock(n)
	access.RLock()
	defer access.RUnlock()

	layers := getActiveLayers(n)
	m := make(map[*Layer]Type, len(layers))
	for _, layer := range layers {
		tex, ok := Active(n, layer)
		if ok {
			m[layer] = tex
		}
	}
	return m
}

type Pair struct {
	*Layer
	Type
}

type sortedPairs []Pair

func (s sortedPairs) Len() int      { return len(s) }
func (s sortedPairs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortedPairs) Less(i, j int) bool {
	return s[i].Layer.Sort() < s[j].Layer.Sort()
}

// Sorted returns a sorted slice of pairs (a texture layer and it's associated
// active texture) for the given node.
//
// It is like the map returned by ActiveTextures(n) except sorted per each
// layer's sort value.
func Sorted(n *scene.Node) []Pair {
	// Build the slice we will sort.
	access := getLayersLock(n)
	access.RLock()

	layers := getActiveLayers(n)
	pairs := make(sortedPairs, len(layers))
	pairs = pairs[:0]
	for _, layer := range layers {
		tex, ok := Active(n, layer)
		if ok {
			pairs = append(pairs, Pair{layer, tex})
		}
	}

	// Unlock now since sort may take some time.
	access.RUnlock()

	// Sort the pairs.
	sort.Sort(pairs)
	return pairs
}
