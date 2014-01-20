// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/shader"
	"code.google.com/p/azul3d/scene/texture"
	"code.google.com/p/azul3d/scene/transparency"
	"fmt"
	"sort"
)

type sortedGeom struct {
	region        *camera.Region
	sorter        *scene.Sorter
	traversalSort uint
	geom          *geom.Mesh
	node          *scene.Node
	camera        *scene.Node
	transparency  transparency.ModeType
	textures      map[*texture.Layer]texture.Type
	shader        *shader.Shader
}

func (s sortedGeom) String() string {
	return fmt.Sprintf("\n\tNode=%v\n\tRegion=%v\n\tSorter=%v\n\tTraversalSort=%v", s.node, s.region, s.sorter, s.traversalSort)
}

func (s sortedGeom) calculateChangeWeight(o *sortedGeom) int {
	weight := 0
	if s.transparency != o.transparency {
		weight++
	}
	return weight
}

type sortedGeoms []*sortedGeom

func (s sortedGeoms) Len() int {
	return len(s)
}

func (s sortedGeoms) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedGeoms) Less(i, j int) bool {
	a := s[i]
	b := s[j]

	// If inside different regions, then sort by region sort first.
	if !a.region.Equals(b.region) {
		return a.region.Sort() < b.region.Sort()
	}

	// If they're under different sorters, then we sort them based off their sorters sort values.
	if a.sorter.SortType() != b.sorter.SortType() {
		return a.sorter.Sort() < b.sorter.Sort()
	}

	// They're the same sorters -- meaning these two geoms are inside the same sorter, so we can
	// freely sort based off the sorter's sort type.
	switch a.sorter.SortType() {
	case scene.ValueSort:
		// We want to sort based off node sort values
		aSort, _ := a.node.Sort()
		bSort, _ := b.node.Sort()
		return aSort < bSort

	case scene.StateSort:
		// We want to sort based off minimal renderer state change, so see who will cause the least
		// amount of change from the previous (i-1) node.
		if i == 0 {
			// We can just say that this one is best, then.
			return true
		}
		z := s[i-1]

		aChangeWeight := a.calculateChangeWeight(z)
		bChangeWeight := b.calculateChangeWeight(z)
		return aChangeWeight < bChangeWeight

	case scene.TraversalSort:
		// We want to sort based off traversal order (see sortGeoms for where this happens)
		return a.traversalSort < b.traversalSort
	}

	if a.node == nil {
		return true
	}
	return false
}

func (r *Renderer) sortGeoms(root *scene.Node, cameras []*scene.Node) sortedGeoms {
	r.lastSortedGeoms = r.lastSortedGeoms[:0]
	geoms := r.lastSortedGeoms

	// Search for geom nodes inside each region's camera's scene.
	for _, camNode := range cameras {
		var camInverse *math.Mat4

		camRegions := camera.Regions(camNode)
		if len(camRegions) == 0 {
			continue
		}

		lp := camera.Lens(camNode).Projection()
		lensProjection := shader.Mat4{
			[4]float32{float32(lp[0][0]), float32(lp[0][1]), float32(lp[0][2]), float32(lp[0][3])},
			[4]float32{float32(lp[1][0]), float32(lp[1][1]), float32(lp[1][2]), float32(lp[1][3])},
			[4]float32{float32(lp[2][0]), float32(lp[2][1]), float32(lp[2][2]), float32(lp[2][3])},
			[4]float32{float32(lp[3][0]), float32(lp[3][1]), float32(lp[3][2]), float32(lp[3][3])},
		}

		for _, region := range camRegions {
			if camInverse == nil {
				//camInverse = camNode.LocalToWorld(camNode.Mat4())
				camWorldTransform := camNode.RelativeTransform(root)
				camInverse, _ = camWorldTransform.Mat4().Invert()
			}

			var traversalSort uint

			// Camera regions must clear regardless of if there are existing geoms.
			geoms = append(geoms, &sortedGeom{
				region: region,
				camera: camNode,
				sorter: scene.Unsorted,
			})

			// Scene must exist and be non-hidden
			camScene := camera.Scene(camNode)
			if camScene == nil || camScene.Hidden() {
				continue
			}

			camScene.Traverse(func(index int, n *scene.Node) bool {
				if !n.Hidden() || n.ShownThrough() {
					nGeoms := geom.Meshes(n)
					if len(nGeoms) == 0 {
						return true
					}

					// In traversal sort mode, nodes with an lower index are drawn first. We expect
					// that nodes who are deeper in the graph are drawn last.
					traversalSort += uint(index) + 1

					activeSorter, hasActiveSorter := n.ActiveSorter()
					activeTransparency := transparency.ActiveMode(n)
					if activeTransparency == transparency.Binary {
						shader.SetInput(n, "BinaryTransparency", int32(1))
					} else {
						shader.SetInput(n, "BinaryTransparency", int32(0))
					}

					activeShader, ok := shader.Active(n)
					if !ok {
						// No custom shader, use default one.
						activeShader = defaultShader
					}

					textures := texture.Textures(n)

					// Calculate an world-transformed matrix.
					nWorldTransform := n.RelativeTransform(root)
					worldMat := nWorldTransform.Mat4()

					// Apply camera transform
					mv := worldMat.Mul(camInverse)

					// Apply the ZUpRight -> YUpRight coordinate system conversion
					mv = mv.Mul(coordSysConversion)
					modelView := shader.Mat4{
						[4]float32{float32(mv[0][0]), float32(mv[0][1]), float32(mv[0][2]), float32(mv[0][3])},
						[4]float32{float32(mv[1][0]), float32(mv[1][1]), float32(mv[1][2]), float32(mv[1][3])},
						[4]float32{float32(mv[2][0]), float32(mv[2][1]), float32(mv[2][2]), float32(mv[2][3])},
						[4]float32{float32(mv[3][0]), float32(mv[3][1]), float32(mv[3][2]), float32(mv[3][3])},
					}

					// modelView * lensProjection
					mvp := mv.Mul(lp)
					modelViewProjection := shader.Mat4{
						[4]float32{float32(mvp[0][0]), float32(mvp[0][1]), float32(mvp[0][2]), float32(mvp[0][3])},
						[4]float32{float32(mvp[1][0]), float32(mvp[1][1]), float32(mvp[1][2]), float32(mvp[1][3])},
						[4]float32{float32(mvp[2][0]), float32(mvp[2][1]), float32(mvp[2][2]), float32(mvp[2][3])},
						[4]float32{float32(mvp[3][0]), float32(mvp[3][1]), float32(mvp[3][2]), float32(mvp[3][3])},
					}

					shader.SetInput(n, "Projection", lensProjection)
					shader.SetInput(n, "ModelView", modelView)
					shader.SetInput(n, "ModelViewProjection", modelViewProjection)

					for _, theGeom := range nGeoms {
						theGeom.RLock()
						defer theGeom.RUnlock()

						if theGeom.Hidden || len(theGeom.Vertices) == 0 {
							// We're not interested in hidden geoms, or ones that contain no vertices.
							continue
						}

						g := &sortedGeom{
							region:        region,
							geom:          theGeom,
							node:          n,
							traversalSort: traversalSort,
							camera:        camNode,
							transparency:  activeTransparency,
							textures:      textures,
							shader:        activeShader,
						}

						if hasActiveSorter {
							g.sorter = activeSorter
						} else {
							if activeTransparency == transparency.None {
								g.sorter = scene.Unsorted
							} else {
								g.sorter = scene.BackToFront
							}
						}

						// Add the sorted geom to the list of scene-sorted geoms
						geoms = append(geoms, g)
					}
				}
				return true
			})
		}
	}

	sort.Sort(geoms)
	r.lastSortedGeoms = geoms
	return geoms
}
