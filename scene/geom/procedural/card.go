// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package procedural

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/texture"
)

// Card builds and returns an new textured card geom
func Card(left, right, bottom, top math.Real, tex texture.Region, hint geom.Hint) *geom.Mesh {
	l, r, b, t := left, right, bottom, top

	indices := []uint32{
		0, // Top Left
		2, // Bottom Left
		3, // Bottom Right
		0, // Top Left
		3, // Bottom Right
		1, // Top Right
	}

	vertices := []geom.Vertex{
		// Top Left
		geom.Vertex{l, 0, t},

		// Top Right
		geom.Vertex{r, 0, t},

		// Bottom Left
		geom.Vertex{l, 0, b},

		// Bottom Right
		geom.Vertex{r, 0, b},
	}

	textureCoords := []texture.Coord{
		// Top Left
		texture.UV(tex.U, tex.V),

		// Top Right
		texture.UV(tex.S, tex.V),

		// Bottom Left
		texture.UV(tex.U, tex.T),

		// Bottom Right
		texture.UV(tex.S, tex.T),
	}

	return &geom.Mesh{
		Hint:     hint,
		Indices:  indices,
		Vertices: vertices,
		TextureCoords: [][]texture.Coord{
			textureCoords,
		},
	}
}
