// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package tmxmesh implements routines for rendering tmx maps, it loads a 2D
// tmx map into a scene node with all required meshes and textures applied so
// it would render correctly.
//
// At present the package only supports orthogonal tile map rendering, and has
// some issues with proper ordering of perspective (I.e. non-uniformly sized)
// tiles (see for instance tiled-qt/examples/perspective_walls.tmx).
package tmxmesh

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/geom/procedural"
	"code.google.com/p/azul3d/scene/texture"
	"code.google.com/p/azul3d/tmx"
	"image"
	"image/draw"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	cw90, cwn90, horizFlip, vertFlip *math.Mat4
)

func init() {
	// Setup rotations
	cw90 = new(math.Mat4)
	cw90.SetRotation(
		math.Real(90).Radians(),
		math.Vector3(0, 1, 0),
		math.CoordSysZUpRight,
	)

	cwn90 = new(math.Mat4)
	cwn90.SetRotation(
		math.Real(-90).Radians(),
		math.Vector3(0, 1, 0),
		math.CoordSysZUpRight,
	)

	// Setup horizontal flip
	horizFlip = new(math.Mat4)
	horizFlip.SetRotation(
		math.Real(180).Radians(),
		math.Vector3(0, 0, 1),
		math.CoordSysZUpRight,
	)

	// Setup vertical flip
	vertFlip = new(math.Mat4)
	vertFlip.SetRotation(
		math.Real(180).Radians(),
		math.Vector3(1, 0, 0),
		math.CoordSysZUpRight,
	)
}

// Config represents a tmx mesh configuration
type Config struct {
	// The offset in node-local units along the Y axis to offset each layer for
	// a proper layer-effect.
	LayerOffset float64
}

// Load loads the given tmx map, m, and returns a scene.Node with the proper
// meshes and textures attatched to it to render the map properly.
//
// If the configuration, c, is non-nil then it is used in place of the default
// configuration.
//
// The tsImages map should be a map of tileset image filenames and their
// associated loaded RGBA images. Tiles who reference tilesets who are not
// found in the map will be omited (not rendered) in the returned node.
func Load(m *tmx.Map, c *Config, tsImages map[string]*image.RGBA) *scene.Node {
	if c == nil {
		c = &Config{
			LayerOffset: 1,
		}
	}
	mapNode := scene.New("tmx")

	// The map node will use multisample transparency (that is, order
	// independent transparency)
	mapNode.SetTransparency(scene.Multisample)

	textures := make(map[string]*texture.Texture2D)
	layerOffset := 0.0
	for _, layer := range m.Layers {
		baseNode := mapNode.New(layer.Name)
		texturedNodes := make(map[*texture.Texture2D]*scene.Node)

		for x := 0; x < m.Width; x++ {
			for y := 0; y < m.Height; y++ {
				gid, hasTile := layer.Tiles[tmx.Coord{x, y}]
				if hasTile {
					tileset := m.FindTileset(gid)
					var region texture.Region

					// Load the tileset texture if needed
					tsImage := filepath.Base(tileset.Image.Source)
					rgba, haveTilesetImage := tsImages[tsImage]
					if !haveTilesetImage {
						// We weren't given a RGBA image for the tileset, so we
						// will just omit this tile.
					}

					tex, ok := textures[tsImage]
					if !ok {
						tex = texture.New()
						tex.SetImage(rgba)
						textures[tsImage] = tex
					}
					r := m.TilesetRect(tileset, rgba.Bounds().Dx(), rgba.Bounds().Dy(), true, gid)
					region = tex.Region(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)

					// Determine which node to attach the tile to, we use a
					// single node per tileset image that way we can avoid
					// making people write multi-texture compatible shaders.
					tsImageNode, ok := texturedNodes[tex]
					if !ok {
						tsImageNode = baseNode.New(tsImage)
						texturedNodes[tex] = tsImageNode
					}

					halfWidth := math.Real(tileset.Width) / 2
					halfHeight := math.Real(tileset.Height) / 2
					card := procedural.Card(-halfWidth, halfWidth, -halfHeight, halfHeight, region, geom.Static)

					// apply necessary flips
					trans := math.Mat4Identity.Copy()
					diagFlipped := (gid & tmx.FLIPPED_DIAGONALLY_FLAG) > 0
					horizFlipped := (gid & tmx.FLIPPED_HORIZONTALLY_FLAG) > 0
					vertFlipped := (gid & tmx.FLIPPED_VERTICALLY_FLAG) > 0
					if diagFlipped {
						if horizFlipped && vertFlipped {
							trans = cw90.Mul(trans)
							trans = horizFlip.Mul(trans)
						} else if horizFlipped {
							trans = cw90.Mul(trans)
						} else if vertFlipped {
							trans = cwn90.Mul(trans)
						} else {
							trans = horizFlip.Mul(trans)
							trans = cw90.Mul(trans)
						}
					} else {
						if horizFlipped {
							trans = horizFlip.Mul(trans)
						}
						if vertFlipped {
							trans = vertFlip.Mul(trans)
						}
					}
					card.Transform(trans)

					// Move the card,
					trans = math.Mat4Identity.Copy()
					tileX := math.Real(x*m.TileWidth) + halfWidth
					tileY := -math.Real(y*m.TileHeight) - halfHeight
					tileLayerOffset := -0.25 * float64(x) // * float64(x * y)
					trans.SetTranslation(math.Vector3(tileX, math.Real(tileLayerOffset), tileY))
					card.Transform(trans)

					geom.Add(tsImageNode, card)
				}
			}
		}

		// Collect all the little tile meshes
		for tex, tsImageNode := range texturedNodes {
			_, collected := geom.Collect(tsImageNode)
			collected.SetParent(baseNode)
			tsImageNode.Destroy()
			//collected := tsImageNode

			// Add the tileset texture to the collected set of tile meshes
			tex.SetWrapModeU(texture.Clamp)
			tex.SetWrapModeV(texture.Clamp)
			texture.Set(collected, texture.DefaultLayer, tex)

			// Decrease Y by layer offset
			x, _, z := collected.Pos()
			collected.SetPos(x, math.Real(layerOffset), z)
		}

		// Decrease layer offset
		layerOffset -= c.LayerOffset
	}
	return mapNode
}

// LoadFile works just like Load except it loads all associated dependencies
// (external tsx tileset files, tileset texture images) for you.
//
// Advanced clients who wish to have more control over file IO will use Load()
// directly instead of using this function.
func LoadFile(path string, c *Config) (*scene.Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	m, err := tmx.Load(data)
	if err != nil {
		return nil, err
	}

	relativeDir := filepath.Dir(path)

	// External tilesets in the map must be loaded seperately
	for _, ts := range m.Tilesets {
		if len(ts.Source) > 0 {
			// Open tsx file
			f, err := os.Open(filepath.Join(relativeDir, filepath.Base(ts.Source)))
			if err != nil {
				return nil, err
			}

			// Read file data
			data, err := ioutil.ReadAll(f)
			if err != nil {
				return nil, err
			}

			// Load the tileset
			err = ts.Load(data)
			if err != nil {
				return nil, err
			}
		}
	}

	// We must also load the images of the tileset
	tsImages := make(map[string]*image.RGBA)
	for _, ts := range m.Tilesets {
		// Name of the tileset image file
		tsImage := filepath.Base(ts.Image.Source)

		// Open tileset image
		f, err := os.Open(filepath.Join(relativeDir, tsImage))
		if err != nil {
			return nil, err
		}

		// Decode the image
		src, _, err := image.Decode(f)
		if err != nil {
			return nil, err
		}

		// If need be, convert to RGBA
		rgba, ok := src.(*image.RGBA)
		if !ok {
			b := src.Bounds()
			rgba = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
			draw.Draw(rgba, rgba.Bounds(), src, b.Min, draw.Src)
		}

		// Put into the tileset images map
		tsImages[tsImage] = rgba
	}

	return Load(m, c, tsImages), nil
}
