// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package procedural

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/texture"
)

// Cube builds and returns an new 3D cube geom
func Cube(scale math.Real, hint geom.Hint) *geom.Mesh {
	s := scale

	vertices := []geom.Vertex{
		// Bottom
		geom.Vertex{-s, s, -s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{-s, s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, s, -s},

		// Front
		geom.Vertex{-s, -s, s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{-s, -s, s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, -s, s},

		// Left
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, s, -s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{-s, -s, s},

		// Back
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, s, -s},
		geom.Vertex{s, s, -s},
		geom.Vertex{-s, s, s},
		geom.Vertex{s, s, -s},
		geom.Vertex{s, s, s},

		// Right
		geom.Vertex{s, s, s},
		geom.Vertex{s, s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, s, s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, -s, s},

		// Top
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, -s, s},
		geom.Vertex{s, -s, s},
		geom.Vertex{-s, s, s},
		geom.Vertex{s, -s, s},
		geom.Vertex{s, s, s},
	}

	textureCoords := []texture.Coord{
		// Bottom
		texture.UV(0, 1),
		texture.UV(0, 0),
		texture.UV(1, 0),
		texture.UV(0, 1),
		texture.UV(1, 0),
		texture.UV(1, 1),

		// Front
		texture.UV(0, 0),
		texture.UV(0, 1),
		texture.UV(1, 1),
		texture.UV(0, 0),
		texture.UV(1, 1),
		texture.UV(1, 0),

		// Left
		texture.UV(0, 0),
		texture.UV(0, 1),
		texture.UV(1, 1),
		texture.UV(0, 0),
		texture.UV(1, 1),
		texture.UV(1, 0),

		// Back
		texture.UV(1, 0),
		texture.UV(1, 1),
		texture.UV(0, 1),
		texture.UV(1, 0),
		texture.UV(0, 1),
		texture.UV(0, 0),

		// Right
		texture.UV(1, 0),
		texture.UV(1, 1),
		texture.UV(0, 1),
		texture.UV(1, 0),
		texture.UV(0, 1),
		texture.UV(0, 0),

		// Top
		texture.UV(0, 0),
		texture.UV(0, 1),
		texture.UV(1, 1),
		texture.UV(0, 0),
		texture.UV(1, 1),
		texture.UV(1, 0),
	}

	return &geom.Mesh{
		Hint:      hint,
		Vertices:  vertices,
		TextureCoords: [][]texture.Coord{
			textureCoords,
		},
	}
}
