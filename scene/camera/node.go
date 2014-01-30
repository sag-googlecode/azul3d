// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package camera implements a scene graph camera node.
package camera

import (
	"azul3d.org/math"
	"azul3d.org/scene"
	"sync"
)

var (
	// The property for storing the lens projection of a node.
	PLensProjection = scene.NewProp("LensProjection")

	// The property for storing the target scene of a node.
	PScene = scene.NewProp("Scene")

	// The property for storing the camera regions map of a node.
	PRegionMap = scene.NewProp("RegionMap")

	// The property for storing the camera regions map lock of a node.
	PRegionMapLock = scene.NewProp("RegionMapLock")
)

func getLock(n *scene.Node) *sync.RWMutex {
	l, ok := n.Prop(PRegionMapLock)
	if !ok {
		newLock := new(sync.RWMutex)
		n.SetProp(PRegionMapLock, newLock)
		return newLock
	}
	return l.(*sync.RWMutex)
}

func getRegions(n *scene.Node) map[*Region]bool {
	regions, ok := n.Prop(PRegionMap)
	if !ok {
		newPRegionMap := make(map[*Region]bool)
		n.SetProp(PRegionMap, newPRegionMap)
		return newPRegionMap
	}
	return regions.(map[*Region]bool)
}

// ClearRegions removes all of the regions associated with the specified camera
// node.
func ClearRegions(n *scene.Node) {
	n.ClearProp(PRegionMap)
}

// Is tells if the specified scene node is an camera node. An camera node is
// any node which specifies an camera lens, as such, this is short hand for:
//
//  Lens(n) != nil
//
func Is(n *scene.Node) bool {
	return Lens(n) != nil
}

// AddRegion adds the specified display region to the camera's list of display
// regions.
func AddRegion(n *scene.Node, region *Region) {
	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	regions := getRegions(n)
	regions[region] = true
}

// RemoveRegion removes the specified display region from the camera's list of
// display regions.
func RemoveRegion(n *scene.Node, region *Region) {
	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	regions := getRegions(n)
	delete(regions, region)
}

// HasRegion tells if the specified display region exists inside this camera's
// list of display regions.
func HasRegion(n *scene.Node, region *Region) bool {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	regions := getRegions(n)
	_, ok := regions[region]
	return ok
}

// Regions returns an list of the regions used by this camera node.
func Regions(n *scene.Node) []*Region {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	regions := getRegions(n)

	c := make([]*Region, len(regions))
	i := 0
	for region, _ := range regions {
		c[i] = region
		i++
	}
	return c
}

// SetScene specifies the scene that this camera node renders.
func SetScene(n, scene *scene.Node) {
	n.SetProp(PScene, scene)
}

// Scene returns the scene that this camera node renders.
func Scene(n *scene.Node) *scene.Node {
	s, ok := n.Prop(PScene)
	if !ok {
		return nil
	}
	return s.(*scene.Node)
}

// SetLens changes the lens of this camera to the specified one.
func SetLens(n *scene.Node, lens *LensProjection) {
	n.SetProp(PLensProjection, lens)
}

// Lens returns the current lens attatched to this camera.
func Lens(n *scene.Node) *LensProjection {
	l, ok := n.Prop(PLensProjection)
	if !ok {
		return nil
	}
	return l.(*LensProjection)
}

// PointToFilm converts a single point in the the specified node's coordinate
// space, to the specified camera 'n' node's film space.
//
// E.g. Converts a 3D point to a 2D screen point.
func PointToFilm(n *scene.Node, point *math.Vec3, space *scene.Node) *math.Vec2 {
	// Find top node
	world := n.Top()

	// Create transform from space to world space
	spaceToWorld := space.RelativeTransform(world).Mat4()

	// Create translation matrix
	mat := new(math.Mat4)
	mat.SetTranslation(point)

	// Convert point from space to world space
	mat = spaceToWorld.Mul(mat)

	// Apply camera transform
	camWorldTransform := n.RelativeTransform(world)
	camInverse, _ := camWorldTransform.Mat4().Invert()
	mat = mat.Mul(camInverse)

	// Retreive point from translation matrix (now in camera space).
	point = mat.Translation()

	// Convert the 3D camera space point to a camera 2D space point.
	p2, ok := Lens(n).Project(point)
	if !ok {
		return nil
	}

	return math.Vector2(p2.X, p2.Z)
}

// InView tells if the specified point in the specified node space is in the
// view frustum of the camera node, n.
func InView(n *scene.Node, point *math.Vec3, space *scene.Node) bool {
	filmSpace := PointToFilm(n, point, space)
	if filmSpace == nil {
		return false
	}
	return true
}
