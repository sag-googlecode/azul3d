// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

import (
	"sync"
)

// Layer represents a single texture layer.
//
// When textures are applied to meshes, they are applied in order of their
// associated layer's sort values.
//
// Consider the following:
//
//  Texture | Layer Sort | Layer Coordinate Index
//  Grass   | 7          | 1
//  Dirt    | 0          | 2
//
// Since the dirt texture has the lowest sort value, it will be the first drawn
// texture. The coordinate set index directly specifies which texture
// coordinate set the texture should be applied to on the mesh, where indices
// can range anywhere from zero to the number of texture coordinate sets the
// mesh actually has.
type Layer struct {
	access           sync.RWMutex
	sort, coordIndex int
}

// SetSort specifies the sort value of this layer. Layers who have lower sort
// values will be drawn first.
func (l *Layer) SetSort(sort int) {
	l.access.Lock()
	defer l.access.Unlock()

	l.sort = sort
}

// Sort returns the sort value of this layer. Layers who have lower sort values
// will be drawn first.
func (l *Layer) Sort() int {
	l.access.RLock()
	defer l.access.RUnlock()

	return l.sort
}

// SetCoordIndex specifies the texture coordinate index of this layer.
//
// This number directly relates to the index of the texture coordinate that
// this texture layer will be mapped to, according to the mesh's texture
// coordinate slice.
//
// If index is < 0; a panic will occur.
func (l *Layer) SetCoordIndex(index int) {
	if index < 0 {
		panic("SetCoordIndex(): Texture coordinate index cannot be < 0!")
	}

	l.access.Lock()
	defer l.access.Unlock()
	l.coordIndex = index
}

// CoordIndex returns the texture coordinate index of this layer.
//
// This number directly relates to the index of the texture coordinate that
// this texture layer will be mapped to, according to the mesh's texture
// coordinate slice.
//
// The number returned will always be >= 0.
func (l *Layer) CoordIndex() int {
	l.access.RLock()
	defer l.access.RUnlock()

	return l.coordIndex
}

// Copy returns a new 1:1 copy of this layer.
func (l *Layer) Copy() *Layer {
	l.access.RLock()
	defer l.access.RUnlock()

	cpy := new(Layer)
	cpy.sort = l.sort
	cpy.coordIndex = l.coordIndex
	return cpy
}

// NewLayer returns a new initialized texture layer, with the default sort and
// texture coordinate index values of zero.
func NewLayer() *Layer {
	l := new(Layer)
	return l
}

var (
	// The default texture layer used by many non-multitextured meshes.
	DefaultLayer = NewLayer()
)
