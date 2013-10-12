// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

// SetIndicesChanged marks the indices of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetIndicesChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.IndicesChanged = changed
}

// HaveIndicesChanged tells if the indices of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveIndicesChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.IndicesChanged
}

// SetVerticesChanged marks the vertices of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetVerticesChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.VerticesChanged = changed
}

// HaveVerticesChanged tells if the vertices of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveVerticesChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.VerticesChanged
}

// SetNormalsChanged marks the normals of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetNormalsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.NormalsChanged = changed
}

// HaveNormalsChanged tells if the normals of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveNormalsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.NormalsChanged
}

// SetTangentsChanged marks the tangents of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetTangentsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.TangentsChanged = changed
}

// HaveTangentsChanged tells if the tangents of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveTangentsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.TangentsChanged
}

// SetBitangentsChanged marks the bitangents of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetBitangentsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.BitangentsChanged = changed
}

// HaveBitangentsChanged tells if the bitangents of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveBitangentsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.BitangentsChanged
}

// SetColorsChanged marks the colors of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetColorsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.ColorsChanged = changed
}

// HaveColorsChanged tells if the colors of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveColorsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.ColorsChanged
}

// SetBoneWeightsChanged marks the bone-weights of this mesh as changed.
//
// This function is thread-safe.
func (m *Mesh) SetBoneWeightsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.BoneWeightsChanged = changed
}

// HaveBoneWeightsChanged tells if the bone-weights of this mesh have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveBoneWeightsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.BoneWeightsChanged
}

// SetTextureCoordsChanged marks the texture coordinates of the specified index
// texture as changed.
//
// This function is thread-safe.
func (m *Mesh) SetTextureCoordsChanged(index int, changed bool) {
	m.Lock()
	defer m.Unlock()

	if index < 0 || index > len(m.TextureCoords) {
		panic("SetTextureCoordsChanged(): Index is out of range!")
	}

	if !changed {
		delete(m.TextureCoordsChanged, index)
	} else {
		if m.TextureCoordsChanged == nil {
			m.TextureCoordsChanged = make(map[int]bool, 1)
		}
		m.TextureCoordsChanged[index] = changed
	}
}

// HaveTextureCoordsChanged tells if the texture coordinates of the specified
// index have changed.
//
// This function is thread-safe.
func (m *Mesh) HaveTextureCoordsChanged(index int) bool {
	m.RLock()
	defer m.RUnlock()

	if index < 0 || index > len(m.TextureCoords) {
		panic("TextureCoordsChanged(): Index is out of range!")
	}

	if m.TextureCoordsChanged == nil {
		return false
	}
	changed, ok := m.TextureCoordsChanged[index]
	if !ok {
		return false
	}
	return changed
}
