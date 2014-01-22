// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package procedural

import (
	"azul3d.org/scene/geom"
	"azul3d.org/scene/texture"
)

// AppendCard builds and appends a texture card mesh to the specified mesh 'm'.
func AppendCard(m *geom.Mesh, left, right, bottom, top float32, tex texture.Region) {
	l, r, b, t := left, right, bottom, top

	vertices := []geom.Vertex{
		// Top Left
		geom.Vertex{l, 0, t},

		// Bottom Left
		geom.Vertex{l, 0, b},

		// Bottom Right
		geom.Vertex{r, 0, b},

		// Top Left
		geom.Vertex{l, 0, t},

		// Bottom Right
		geom.Vertex{r, 0, b},

		// Top Right
		geom.Vertex{r, 0, t},
	}

	textureCoords := []texture.Coord{
		// Top Left
		texture.UV(tex.U, tex.V),

		// Bottom Left
		texture.UV(tex.U, tex.T),

		// Bottom Right
		texture.UV(tex.S, tex.T),

		// Top Left
		texture.UV(tex.U, tex.V),

		// Bottom Right
		texture.UV(tex.S, tex.T),

		// Top Right
		texture.UV(tex.S, tex.V),
	}

	m.Lock()
	defer m.Unlock()
	m.Vertices = append(m.Vertices, vertices...)

	if len(m.TextureCoords) > 0 {
		m.TextureCoords[0] = append(m.TextureCoords[0], textureCoords...)
		return
	}
	m.TextureCoords = append(m.TextureCoords, textureCoords)
}

// Card builds and returns an new textured card mesh.
func Card(left, right, bottom, top float32, tex texture.Region, hint geom.Hint) *geom.Mesh {
	m := new(geom.Mesh)
	m.Hint = hint
	AppendCard(m, left, right, bottom, top, tex)
	return m
}
