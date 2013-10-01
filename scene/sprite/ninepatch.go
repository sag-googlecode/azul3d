// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package sprite

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/geom/procedural"
	"code.google.com/p/azul3d/scene/texture"
)

// Patches describes the size of the patches in an nine-patch mesh.
type Patches struct {
	// Width and height of center patch.
	Width, Height math.Real

	// Width of left and right patches.
	Left, Right math.Real

	// Height of bottom and top patches.
	Bottom, Top math.Real
}

// Regions describes the texture regions of all nine patches of an nine-patch
// mesh.
type Regions struct {
	// Region of center patch.
	Center texture.Region

	// Region of left, right, bottom, and top patches.
	Left, Right, Bottom, Top texture.Region

	// Region of top-left, top-right, bottom-left, and bottom-right patches.
	TopLeft, TopRight, BottomLeft, BottomRight texture.Region
}

// Ninepatch builds and returns an new nine patch geom
//
// If width and height are the only non-zero demensions specified in the
// patches, then an procedural.Card() is returned instead of an nine-patch
// mesh.
func Ninepatch(p *Patches, r *Regions, hint geom.Hint) *geom.Mesh {
	if p.Left.Equals(0) && p.Right.Equals(0) && p.Bottom.Equals(0) && p.Top.Equals(0) {
		w2 := p.Width / 2
		h2 := p.Height / 2
		return procedural.Card(-w2, w2, -h2, h2, r.Center, hint)
	}

	// ------------------------------------
	// -x0      x1-     y0     -x2      x3-
	// -          -            -          -
	// -          -            -          -
	// -          -            -          -
	// -          -     y1     -          -
	// ------------------------------------
	// -                                  -
	// -                                  -
	// -                                  -
	// -                                  -
	// -                                  -
	// -                                  -
	// -                y2                -
	// ------------------------------------
	// -          -            -          -
	// -          -            -          -
	// -          -            -          -
	// -          -            -          -
	// -          -     y3     -          -
	// ------------------------------------
	halfWidth := p.Width / 2
	halfHeight := p.Height / 2

	x0 := -p.Left - halfWidth
	x1 := -halfWidth
	x2 := halfWidth
	x3 := halfWidth + p.Right

	y0 := p.Top + halfHeight
	y1 := halfHeight
	y2 := -halfHeight
	y3 := -halfHeight - p.Bottom

	vertices := []geom.Vertex{
		// Top-left patch
		geom.Vertex{x0, 0, y0},
		geom.Vertex{x0, 0, y1},
		geom.Vertex{x1, 0, y1},

		geom.Vertex{x0, 0, y0},
		geom.Vertex{x1, 0, y1},
		geom.Vertex{x1, 0, y0},

		// Top patch
		geom.Vertex{x1, 0, y0},
		geom.Vertex{x1, 0, y1},
		geom.Vertex{x2, 0, y1},

		geom.Vertex{x1, 0, y0},
		geom.Vertex{x2, 0, y1},
		geom.Vertex{x2, 0, y0},

		// Top-right patch
		geom.Vertex{x2, 0, y0},
		geom.Vertex{x2, 0, y1},
		geom.Vertex{x3, 0, y1},

		geom.Vertex{x2, 0, y0},
		geom.Vertex{x3, 0, y1},
		geom.Vertex{x3, 0, y0},

		// Left patch
		geom.Vertex{x0, 0, y1},
		geom.Vertex{x0, 0, y2},
		geom.Vertex{x1, 0, y2},

		geom.Vertex{x0, 0, y1},
		geom.Vertex{x1, 0, y2},
		geom.Vertex{x1, 0, y1},

		// Center patch
		geom.Vertex{x1, 0, y1},
		geom.Vertex{x1, 0, y2},
		geom.Vertex{x2, 0, y2},

		geom.Vertex{x1, 0, y1},
		geom.Vertex{x2, 0, y2},
		geom.Vertex{x2, 0, y1},

		// Right patch
		geom.Vertex{x2, 0, y1},
		geom.Vertex{x2, 0, y2},
		geom.Vertex{x3, 0, y2},

		geom.Vertex{x2, 0, y1},
		geom.Vertex{x3, 0, y2},
		geom.Vertex{x3, 0, y1},

		// Bottom-left patch
		geom.Vertex{x0, 0, y2},
		geom.Vertex{x0, 0, y3},
		geom.Vertex{x1, 0, y3},

		geom.Vertex{x0, 0, y2},
		geom.Vertex{x1, 0, y3},
		geom.Vertex{x1, 0, y2},

		// Bottom patch
		geom.Vertex{x1, 0, y2},
		geom.Vertex{x1, 0, y3},
		geom.Vertex{x2, 0, y3},

		geom.Vertex{x1, 0, y2},
		geom.Vertex{x2, 0, y3},
		geom.Vertex{x2, 0, y2},

		// Bottom-right patch
		geom.Vertex{x2, 0, y2},
		geom.Vertex{x2, 0, y3},
		geom.Vertex{x3, 0, y3},

		geom.Vertex{x2, 0, y2},
		geom.Vertex{x3, 0, y3},
		geom.Vertex{x3, 0, y2},
	}

	textureCoords := []texture.Coord{
		// Top-left patch
		texture.UV(r.TopLeft.U, r.TopLeft.V),
		texture.UV(r.TopLeft.U, r.TopLeft.T),
		texture.UV(r.TopLeft.S, r.TopLeft.T),

		texture.UV(r.TopLeft.U, r.TopLeft.V),
		texture.UV(r.TopLeft.S, r.TopLeft.T),
		texture.UV(r.TopLeft.S, r.TopLeft.V),

		// Top patch
		texture.UV(r.Top.U, r.Top.V),
		texture.UV(r.Top.U, r.Top.T),
		texture.UV(r.Top.S, r.Top.T),

		texture.UV(r.Top.U, r.Top.V),
		texture.UV(r.Top.S, r.Top.T),
		texture.UV(r.Top.S, r.Top.V),

		// Top-right patch
		texture.UV(r.TopRight.U, r.TopRight.V),
		texture.UV(r.TopRight.U, r.TopRight.T),
		texture.UV(r.TopRight.S, r.TopRight.T),

		texture.UV(r.TopRight.U, r.TopRight.V),
		texture.UV(r.TopRight.S, r.TopRight.T),
		texture.UV(r.TopRight.S, r.TopRight.V),

		// Left patch
		texture.UV(r.Left.U, r.Left.V),
		texture.UV(r.Left.U, r.Left.T),
		texture.UV(r.Left.S, r.Left.T),

		texture.UV(r.Left.U, r.Left.V),
		texture.UV(r.Left.S, r.Left.T),
		texture.UV(r.Left.S, r.Left.V),

		// Center patch
		texture.UV(r.Center.U, r.Center.V),
		texture.UV(r.Center.U, r.Center.T),
		texture.UV(r.Center.S, r.Center.T),

		texture.UV(r.Center.U, r.Center.V),
		texture.UV(r.Center.S, r.Center.T),
		texture.UV(r.Center.S, r.Center.V),

		// Right patch
		texture.UV(r.Right.U, r.Right.V),
		texture.UV(r.Right.U, r.Right.T),
		texture.UV(r.Right.S, r.Right.T),

		texture.UV(r.Right.U, r.Right.V),
		texture.UV(r.Right.S, r.Right.T),
		texture.UV(r.Right.S, r.Right.V),

		// Bottom-left patch
		texture.UV(r.BottomLeft.U, r.BottomLeft.V),
		texture.UV(r.BottomLeft.U, r.BottomLeft.T),
		texture.UV(r.BottomLeft.S, r.BottomLeft.T),

		texture.UV(r.BottomLeft.U, r.BottomLeft.V),
		texture.UV(r.BottomLeft.S, r.BottomLeft.T),
		texture.UV(r.BottomLeft.S, r.BottomLeft.V),

		// Bottom patch
		texture.UV(r.Bottom.U, r.Bottom.V),
		texture.UV(r.Bottom.U, r.Bottom.T),
		texture.UV(r.Bottom.S, r.Bottom.T),

		texture.UV(r.Bottom.U, r.Bottom.V),
		texture.UV(r.Bottom.S, r.Bottom.T),
		texture.UV(r.Bottom.S, r.Bottom.V),

		// Bottom-right patch
		texture.UV(r.BottomRight.U, r.BottomRight.V),
		texture.UV(r.BottomRight.U, r.BottomRight.T),
		texture.UV(r.BottomRight.S, r.BottomRight.T),

		texture.UV(r.BottomRight.U, r.BottomRight.V),
		texture.UV(r.BottomRight.S, r.BottomRight.T),
		texture.UV(r.BottomRight.S, r.BottomRight.V),
	}

	return &geom.Mesh{
		Hint:      hint,
		Vertices:  vertices,
		TextureCoords: [][]texture.Coord{
			textureCoords,
		},
	}
}
