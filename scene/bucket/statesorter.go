// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

import (
	"azul3d.org/v1/scene/shader"
	"azul3d.org/v1/scene/texture"
	"azul3d.org/v1/scene/transparency"
)

type stateSorter struct{}

// implements Sorter interface. Sorts based off render state changes (i.e.
// similar shaders, textures, etc).
func (s *stateSorter) Less(cam, a, b Node) bool {
	i := a.Node()
	j := b.Node()

	// Sort by shader.
	iShader := shader.Active(i)
	jShader := shader.Active(j)
	if iShader != jShader {
		return false
	}

	// Sort by transparency mode.
	iTransparency := transparency.ActiveMode(i)
	jTransparency := transparency.ActiveMode(j)
	if iTransparency != jTransparency {
		return false
	}

	// Sort by textures.
	iTextures := texture.Sorted(i)
	jTextures := texture.Sorted(j)
	if len(iTextures) != len(jTextures) {
		return false
	}

	// No textures to sort, last texture comparison would fail below.
	if len(iTextures) == 0 {
		return false
	}

	// Try all but the last texture comparison.
	var k int
	for k = 0; k < len(iTextures)-1; k++ {
		if jTextures[k] != iTextures[k] {
			return false
		}
	}

	// All textures are equal, so just return the last texture comparison.
	return jTextures[k] != iTextures[k]
}

// NewStateSorter returns a new state sorter which sorts nodes based off render
// state changes, such as changes in shaders, transparency modes, textures, and
// other render states. This can be an important part of rendering many objects
// quickly when they have diverse states, as these state changes can be costly
// performance-wise on graphics cards.
func NewStateSorter() Sorter {
	return new(stateSorter)
}
