// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package geom implements various geometrical data types.
package geom

import (
	"azul3d.org/v1/math"
	"azul3d.org/v1/scene/color"
	"azul3d.org/v1/scene/texture"
	"fmt"
	"sync"
)

// Mesh represents a single mesh made up of several components. A mesh may or
// may not be made up of indexed triangles, normals, etc, depending on whether
// or not len(m.Indices) == 0 holds true.
//
// In the event that a mesh is indexed, m.Indices holds the indices and it can
// be expected that Vertices, Normals, Tangents, etc, will hold at least enough
// elements (or zero elements) such that the each index will not be out of
// bounds.
//
// Most of the members that make up a mesh are public, as such synchronization
// through the use of a mesh's RWMutex is necessary (and was done so for
// performance reasons).
type Mesh struct {
	sync.RWMutex

	Hint   Hint
	Hidden bool

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
//
// This function is thread-safe.
func (m *Mesh) MarkLoaded() {
	m.Lock()
	defer m.Unlock()

	m.Loaded = true
	for _, notify := range m.notifiers {
		notify <- true
	}
}

// Loaded tells if this geom is currently loaded or not.
//
// This function is thread-safe.
func (m *Mesh) IsLoaded() bool {
	m.RLock()
	defer m.RUnlock()

	return m.Loaded
}

// LoadNotify returns an channel on which true is sent once this geom is marked
// as loaded (normally by the renderer).
//
// This function is thread-safe.
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
//
// This function is thread-safe.
func (m *Mesh) SetNativeIdentity(identity interface{}) {
	m.Lock()
	defer m.Unlock()

	m.nativeIdentity = identity
}

// NativeIdentity returns the native identity of this mesh.
//
// This should mostly not be used (except in very rare, advanced cases).
//
// This function is thread-safe.
func (m *Mesh) NativeIdentity() interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.nativeIdentity
}

// String returns a string representation of this mesh.
//
// This function is thread-safe.
func (m *Mesh) String() string {
	m.RLock()
	defer m.RUnlock()

	return fmt.Sprintf("Mesh(%v, Hidden=%v, %v Vertices)", m.Hint, m.Hidden, len(m.Vertices))
}

// SetUsageHint sets the usage hint of this mesh.
//
// This function is thread-safe.
func (m *Mesh) SetUsageHint(hint Hint) {
	m.Lock()
	defer m.Unlock()

	m.Hint = hint
}

// UsageHint returns the usage hint of this mesh.
//
// This function is thread-safe.
func (m *Mesh) UsageHint() Hint {
	m.RLock()
	defer m.RUnlock()

	return m.Hint
}

// SetHidden specifies whether this mesh is considered hidden or not.
//
// This function is thread-safe.
func (m *Mesh) SetHidden(hidden bool) {
	m.Lock()
	defer m.Unlock()

	m.Hidden = hidden
}

// IsHidden tells if this mesh is considered hidden.
//
// This function is thread-safe.
func (m *Mesh) IsHidden() bool {
	m.RLock()
	defer m.RUnlock()

	return m.Hidden
}

// Valid checks this mesh for validity. For an mesh to be valid, it's vertex
// property arrays must all be of the same length.
//
// This function is thread-safe.
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

// Copy returns a new 1:1 copy of this Mesh. This is a potentially expensive
// operation depending on how many vertices, etc this mesh contains.
//
// This function is thread-safe.
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
//
// This function is thread-safe.
func (m *Mesh) MakePixelPerfect() {
	m.Lock()
	defer m.Unlock()

	for vi, v := range m.Vertices {
		rx := float32(math.Rounded(float64(v.X)))
		ry := float32(math.Rounded(float64(v.Y)))
		rz := float32(math.Rounded(float64(v.Z)))
		m.Vertices[vi] = Vertex{rx + +0.5, ry + 0.5, rz + 0.5}
	}
}

// Transform transforms each vertex in this mesh by the specified affine
// transformation matrix.
//
// This function is thread-safe.
func (m *Mesh) Transform(mat math.Mat4) {
	m.Lock()
	defer m.Unlock()

	for vi, v := range m.Vertices {
		vect := math.Vec3{float64(v.X), float64(v.Y), float64(v.Z)}
		vect = vect.TransformMat4(mat)
		m.Vertices[vi] = Vertex{float32(vect.X), float32(vect.Y), float32(vect.Z)}
	}
}

// BoundingBox returns the axis aligned bounding box of this mesh, or nil if it
// does not have one.
//
// Also see CalculateBounds() which will calculate an axis aligned bounding box
// for this mesh automatically.
//
// This function is thread-safe.
func (m *Mesh) BoundingBox() *BoundingBox {
	m.RLock()
	defer m.RUnlock()

	return m.aabb
}

// SetBoundingBox sets the axis aligned bounding box of this mesh to the
// specified one.
//
// You may pass in nil to imply that the mesh has no bounding box.
//
// This function is thread-safe.
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
//
// This function is thread-safe.
func (m *Mesh) CalculateBounds() {
	// We only need read access to calculate a bounding box, which can relieve
	// some contengency while calculating the AABB.
	m.RLock()
	var bb *BoundingBox

	if len(m.Vertices) > 0 {
		bb = new(BoundingBox)
		var min, max math.Vec3
		for _, vert := range m.Vertices {
			vx := float64(vert.X)
			vy := float64(vert.Y)
			vz := float64(vert.Z)

			if vx < min.X {
				min.X = vx
			} else if vx > max.X {
				max.X = vx
			}

			if vy < min.Y {
				min.Y = vy
			} else if vy > max.Y {
				max.Y = vy
			}

			if vz < min.Z {
				min.Z = vz
			} else if vz > max.Z {
				max.Z = vz
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
//
// This function is thread-safe.
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
