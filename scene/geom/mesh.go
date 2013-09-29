// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package geom implements various geometrical data types.
package geom

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene/color"
	"code.google.com/p/azul3d/scene/texture"
	"fmt"
	"sync"
)

type Mesh struct {
	sync.RWMutex

	Hint      Hint
	Hidden    bool
	Primitive Primitive

	Indices       []uint32
	Vertices      []Vertex
	Normals       []Normal
	Tangents      []Tangent
	Bitangents    []Bitangent
	Colors        []color.Color
	BoneWeights   []BoneWeight
	TextureCoords [][]texture.Coord

	IndicesChanged, VerticesChanged, NormalsChanged, TangentsChanged, BitangentsChanged, ColorsChanged, BoneWeightsChanged bool

	TextureCoordsChanged map[int]bool

	Loaded bool
	aabb   *BoundingBox

	nativeIdentity interface{}
	notifiers      []chan bool
}

// MarkLoaded marks this geom as loaded. Only the renderer should call this,
// and as such you should (normally) never call this function.
func (m *Mesh) MarkLoaded() {
	m.Lock()
	defer m.Unlock()

	m.Loaded = true
	for _, notify := range m.notifiers {
		notify <- true
	}
}

// Loaded tells if this geom is currently loaded or not.
func (m *Mesh) IsLoaded() bool {
	m.RLock()
	defer m.RUnlock()

	return m.Loaded
}

// LoadNotify returns an channel on which true is sent once this geom is marked
// as loaded (normally by the renderer).
func (m *Mesh) LoadNotify() chan bool {
	m.Lock()
	defer m.Unlock()

	notify := make(chan bool, 1)
	if m.Loaded {
		notify <- true
		return notify
	}

	if m.notifiers == nil {
		m.notifiers = make([]chan bool, 0)
	}
	m.notifiers = append(m.notifiers, notify)
	return notify
}

// SetNativeIdentity specifies the native identity of this mesh.
//
// This should mostly not be used (except in very rare, advanced cases).
func (m *Mesh) SetNativeIdentity(identity interface{}) {
	m.Lock()
	defer m.Unlock()

	m.nativeIdentity = identity
}

// NativeIdentity returns the native identity of this mesh.
//
// This should mostly not be used (except in very rare, advanced cases).
func (m *Mesh) NativeIdentity() interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.nativeIdentity
}

// String returns a string representation of this mesh.
func (m *Mesh) String() string {
	m.RLock()
	defer m.RUnlock()

	return fmt.Sprintf("Mesh(%v, Hidden=%v, %v Vertices)", m.Hint, m.Hidden, len(m.Vertices))
}

// SetUsageHint sets the usage hint of this mesh.
func (m *Mesh) SetUsageHint(hint Hint) {
	m.Lock()
	defer m.Unlock()

	m.Hint = hint
}

// UsageHint returns the usage hint of this mesh.
func (m *Mesh) UsageHint() Hint {
	m.RLock()
	defer m.RUnlock()

	return m.Hint
}

// SetHidden specifies weather this mesh is considered hidden or not.
func (m *Mesh) SetHidden(hidden bool) {
	m.Lock()
	defer m.Unlock()

	m.Hidden = hidden
}

// IsHidden tells if this mesh is considered hidden.
func (m *Mesh) IsHidden() bool {
	m.RLock()
	defer m.RUnlock()

	return m.Hidden
}

// Valid checks this mesh for validity. For an mesh to be valid, it's vertex
// property arrays must all be of the same length.
func (m *Mesh) Valid() bool {
	m.RLock()
	defer m.RUnlock()

	var verts int
	if m.Vertices != nil {
		verts = len(m.Vertices)
	}

	// Check normals validity
	if m.Normals != nil {
		if len(m.Normals) != verts {
			return false
		}
	}

	// Check tangents validity
	if m.Tangents != nil {
		if len(m.Tangents) != verts {
			return false
		}
	}

	// Check texture coordinates key map validity
	if m.TextureCoords != nil {
		for _, coords := range m.TextureCoords {
			if len(coords) != verts {
				return false
			}
		}
	}

	// Check colors validity
	if m.Colors != nil {
		if len(m.Colors) != verts {
			return false
		}
	}

	// And finally check bone weights validity
	if m.BoneWeights != nil {
		if len(m.BoneWeights) != verts {
			return false
		}
	}

	// We're valid! Yay!
	return true
}

func (m *Mesh) Copy() *Mesh {
	m.RLock()
	defer m.RUnlock()

	c := &Mesh{
		Vertices:      make([]Vertex, len(m.Vertices)),
		Normals:       make([]Normal, len(m.Normals)),
		Tangents:      make([]Tangent, len(m.Tangents)),
		TextureCoords: make([][]texture.Coord, len(m.TextureCoords)),
		Colors:        make([]color.Color, len(m.Colors)),
		BoneWeights:   make([]BoneWeight, len(m.BoneWeights)),

		IndicesChanged:       m.IndicesChanged,
		VerticesChanged:      m.VerticesChanged,
		NormalsChanged:       m.NormalsChanged,
		TangentsChanged:      m.TangentsChanged,
		BitangentsChanged:    m.BitangentsChanged,
		ColorsChanged:        m.ColorsChanged,
		BoneWeightsChanged:   m.BoneWeightsChanged,
		TextureCoordsChanged: make(map[int]bool, len(m.TextureCoordsChanged)),
	}
	copy(c.Vertices, m.Vertices)
	copy(c.Normals, m.Normals)
	copy(c.Tangents, m.Tangents)
	for i, uvs := range m.TextureCoords {
		uvsCopy := make([]texture.Coord, len(uvs))
		copy(uvsCopy, uvs)
		c.TextureCoords[i] = uvsCopy
	}
	copy(c.Colors, m.Colors)
	copy(c.BoneWeights, m.BoneWeights)
	for index, changed := range m.TextureCoordsChanged {
		c.TextureCoordsChanged[index] = changed
	}
	c.aabb = m.aabb.Copy()
	return c
}

// MakePixelPerfect rounds each vertex position to it's nearest whole value, and
// adds 0.5 to it.
//
// This is needed for pixel-perfect sprites for example, due to the way OpenGL
// workss.
func (m *Mesh) MakePixelPerfect() {
	m.Lock()
	defer m.Unlock()

	for vi, v := range m.Vertices {
		m.Vertices[vi] = Vertex{v.X.Rounded() + 0.5, v.Y.Rounded() + 0.5, v.Z.Rounded() + 0.5}
	}
}

// Transform transforms each vertex in this mesh by the specified affine
// transformation matrix.
func (m *Mesh) Transform(mat *math.Mat4) {
	m.Lock()
	defer m.Unlock()

	for vi, v := range m.Vertices {
		vect := math.Vector3(v.X, v.Y, v.Z)
		vect.TransformMat4(mat)
		m.Vertices[vi] = Vertex{vect.X, vect.Y, vect.Z}
	}
}

// BoundingBox returns the axis aligned bounding box of this mesh, or nil if it
// does not have one.
//
// Also see CalculateBounds() which will calculate an axis aligned bounding box
// for this mesh automatically.
func (m *Mesh) BoundingBox() *BoundingBox {
	m.RLock()
	defer m.RUnlock()

	return m.aabb
}

// SetBoundingBox sets the axis aligned bounding box of this mesh to the
// specified one.
//
// You may pass in nil to imply that the mesh has no bounding box.
func (m *Mesh) SetBoundingBox(bb *BoundingBox) {
	m.Lock()
	defer m.Unlock()

	m.aabb = bb
}

// CalculateBounds calculates a new axis aligned bounding box for this mesh.
//
// Note: The mesh must have vertices or else the bounding box of this mesh will
// be set to nil.
//
// Note: Even if this mesh already has a bounding box, it will be calculated
// again. You should check by yourself first.
//
// Note: It may be benificial depending on your use case to run this function
// in a seperate goroutine.
func (m *Mesh) CalculateBounds() {
	// We only need read access to calculate a bounding box, which can relieve
	// some contengency while calculating the AABB.
	m.RLock()
	var bb *BoundingBox

	if len(m.Vertices) > 0 {
		bb = new(BoundingBox)
		min := new(math.Vec3)
		max := new(math.Vec3)
		for _, vert := range m.Vertices {
			if vert.X < min.X {
				min.X = vert.X
			} else if vert.X > max.X {
				max.X = vert.X
			}

			if vert.Y < min.Y {
				min.Y = vert.Y
			} else if vert.Y > max.Y {
				max.Y = vert.Y
			}

			if vert.Z < min.Z {
				min.Z = vert.Z
			} else if vert.Z > max.Z {
				max.Z = vert.Z
			}
		}
		bb.Min = min
		bb.Max = max
	}

	// Now we need to assign the AABB, so we exit read lock and assign,
	m.RUnlock()

	m.SetBoundingBox(bb)
}

// BakeColors bakes the specified rgba color and rgba color scale into this
// meshes' vertex colors.
//
// If the c parameter is not color.None, then vertex colors are replaced with
// the specified parameter.
//
// If the color scale parameter is not color.None, then resulting vertex colors
// are multiplied against each respective RGBA component of the specified color
// scale.
func (m *Mesh) BakeColors(c, colorScale color.Color) {
	if c.Equals(color.None) && colorScale.Equals(color.None) {
		return
	}

	// Enter read lock
	m.RLock()

	if len(m.Vertices) > 0 {
		// Exit read lock
		m.RUnlock()

		// Enter write lock
		m.Lock()
		defer m.Unlock()

		if c.Equals(color.None) {
			// No color specified, so we want to scale per-vertex colors.
			if len(m.Colors) == 0 {
				// If vertex colors don't exist yet, they're considered to be
				// [1, 1, 1, 1], we need to account for this here.
				m.Colors = make([]color.Color, len(m.Vertices))
				for i, _ := range m.Colors {
					// We can skip multiplication since the color is known to
					// be [1, 1, 1, 1], and just assign the scale itself here.
					m.Colors[i] = colorScale
				}
			} else {
				// Vertex colors exist, scale them appropriately.
				if colorScale.Equals(color.New(1, 1, 1, 1)) {
					// We won't gain anything by doing multiplication against
					// one, since thats the color itself. We don't need to do
					// anything here.
					return
				}

				for i, vertColor := range m.Colors {
					m.Colors[i] = colorScale.Mul(vertColor)
				}
			}

		} else {
			// Okay, we have an color to replace all existing vertex colors.
			if !colorScale.Equals(color.None) {
				// Since we're replacing existing vertex colors, we can simply
				// multiply the color scale here instead of per-vertex.
				c = c.Mul(colorScale)
			}

			if len(m.Colors) == 0 {
				// Vertex colors don't exist yet in the mesh, let's add them.
				m.Colors = make([]color.Color, len(m.Vertices))
			}

			// Assign each vertex color here now.
			for i, _ := range m.Colors {
				m.Colors[i] = c
			}
		}

		return
	}

	// Exit read lock
	m.RUnlock()
}
