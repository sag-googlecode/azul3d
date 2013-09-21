// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/util"
	"fmt"
	"sort"
)

type sortedGeom struct {
	region        *util.Region
	sorter        *scene.Sorter
	sortType      scene.SortType
	sorterSort    uint
	traversalSort uint
	geom          *geom.Mesh
	node          *scene.Node
	camera        *scene.Node
	transparency  scene.TransparencyMode
	mat           *math.Mat4
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
	if a.region.Id() != b.region.Id() {
		// If they share an identical sort value -- incorrect ordering and flickering could occur,
		// so to fix this we sort by their id's in that case.
		aSort := a.region.Sort()
		bSort := b.region.Sort()
		if aSort != bSort {
			return aSort < bSort
		}
		return a.region.Id() < b.region.Id()
	}

	// If they're under different sorters, then we sort them based off their sorters sort values.
	if a.sortType != b.sortType {
		return a.sorterSort < b.sorterSort
	}

	// They're the same sorters -- meaning these two geoms are inside the same sorter, so we can
	// freely sort based off the sorter's sort type.
	switch a.sortType {
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
	} else if b.node == nil {
		return false
	}

	// Never gets here.
	panic("Unable to sort geoms; reached end of function.")
}

func (r *GLFFRenderer) sortGeoms(root *scene.Node, cameras []*scene.Node, defaultRegion *util.Region) sortedGeoms {
	var geoms sortedGeoms

	anyCameraHasRegion := false
	for _, camNode := range cameras {
		if len(camera.Regions(camNode)) > 0 {
			anyCameraHasRegion = true
			break
		}
	}

	// Search for geom nodes inside each region's camera's scene.
	for _, camNode := range cameras {
		var camInverse *math.Mat4

		var camRegions []*util.Region
		if anyCameraHasRegion {
			camRegions = camera.Regions(camNode)
		} else {
			camRegions = []*util.Region{
				defaultRegion,
			}
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
					activeTransparency := n.ActiveTransparency()

					// Calculate an world-transformed matrix.
					nWorldTransform := n.RelativeTransform(root)
					worldMat := nWorldTransform.Mat4()

					// Apply camera transform
					mat := worldMat.Mul(camInverse)

					// Apply the ZUpRight -> YUpRight coordinate system conversion
					mat = mat.Mul(coordSysConversion)

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
							mat:           mat,
						}

						if hasActiveSorter {
							g.sorter = activeSorter
						} else {
							if activeTransparency == scene.NoTransparency {
								g.sorter = scene.Unsorted
							} else {
								g.sorter = scene.BackToFront
							}
						}
						g.sortType = g.sorter.SortType()
						g.sorterSort = g.sorter.Sort()

						// Add the sorted geom to the list of scene-sorted geoms
						geoms = append(geoms, g)
					}
				}
				return true
			})
		}
	}

	sort.Sort(geoms)

	//if debug {
	//	fmt.Println("\nAfter sorting geoms")
	//	for _, geom := range geoms {
	//		fmt.Println(geom)
	//	}
	//}

	//sort.Sort(geoms)

	return geoms
}
