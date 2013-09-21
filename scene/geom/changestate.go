// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

func (m *Mesh) SetIndicesChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.IndicesChanged = changed
}

func (m *Mesh) HaveIndicesChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.IndicesChanged
}

func (m *Mesh) SetVerticesChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.VerticesChanged = changed
}

func (m *Mesh) HaveVerticesChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.VerticesChanged
}

func (m *Mesh) SetNormalsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.NormalsChanged = changed
}

func (m *Mesh) HaveNormalsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.NormalsChanged
}

func (m *Mesh) SetTangentsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.TangentsChanged = changed
}

func (m *Mesh) HaveTangentsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.TangentsChanged
}

func (m *Mesh) SetBitangentsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.BitangentsChanged = changed
}

func (m *Mesh) HaveBitangentsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.BitangentsChanged
}

func (m *Mesh) SetColorsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.ColorsChanged = changed
}

func (m *Mesh) HaveColorsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.ColorsChanged
}

func (m *Mesh) SetBoneWeightsChanged(changed bool) {
	m.Lock()
	defer m.Unlock()

	m.BoneWeightsChanged = changed
}

func (m *Mesh) HaveBoneWeightsChanged() bool {
	m.RLock()
	defer m.RUnlock()

	return m.BoneWeightsChanged
}

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
