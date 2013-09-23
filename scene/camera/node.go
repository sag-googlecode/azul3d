// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package camera implements a scene graph camera node.
package camera

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/util"
	"sync"
)

var (
	lensTag       = scene.NewProp("cameraLens")
	sceneTag      = scene.NewProp("cameraScene")
	regionMap     = scene.NewProp("cameraRegionMap")
	regionMapLock = scene.NewProp("cameraRegionMapLock")
)

func getLock(n *scene.Node) *sync.RWMutex {
	l, ok := n.Tag(regionMapLock)
	if !ok {
		newLock := new(sync.RWMutex)
		n.SetTag(regionMapLock, newLock)
		return newLock
	}
	return l.(*sync.RWMutex)
}

func getRegions(n *scene.Node) map[*util.Region]bool {
	regions, ok := n.Tag(regionMap)
	if !ok {
		newRegionMap := make(map[*util.Region]bool)
		n.SetTag(regionMap, newRegionMap)
		return newRegionMap
	}
	return regions.(map[*util.Region]bool)
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
func AddRegion(n *scene.Node, region *util.Region) {
	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	regions := getRegions(n)
	regions[region] = true
}

// RemoveRegion removes the specified display region from the camera's list of
// display regions.
func RemoveRegion(n *scene.Node, region *util.Region) {
	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	regions := getRegions(n)
	delete(regions, region)
}

// HasRegion tells if the specified display region exists inside this camera's
// list of display regions.
func HasRegion(n *scene.Node, region *util.Region) bool {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	regions := getRegions(n)
	_, ok := regions[region]
	return ok
}

// Regions returns an list of the regions used by this camera node.
func Regions(n *scene.Node) []*util.Region {
	access := getLock(n)
	access.RLock()
	defer access.RUnlock()

	regions := getRegions(n)

	c := make([]*util.Region, len(regions))
	i := 0
	for region, _ := range regions {
		c[i] = region
		i++
	}
	return c
}

// SetScene specifies the scene that this camera node renders.
func SetScene(n, scene *scene.Node) {
	n.SetTag(sceneTag, scene)
}

// Scene returns the scene that this camera node renders.
func Scene(n *scene.Node) *scene.Node {
	s, ok := n.Tag(sceneTag)
	if !ok {
		return nil
	}
	return s.(*scene.Node)
}

// SetLens changes the lens of this camera to the specified one.
func SetLens(n *scene.Node, lens *util.Lens) {
	n.SetTag(lensTag, lens)
}

// Lens returns the current lens attatched to this camera.
func Lens(n *scene.Node) *util.Lens {
	l, ok := n.Tag(lensTag)
	if !ok {
		return nil
	}
	return l.(*util.Lens)
}

func makeCopy(n *scene.Node, deep bool) {
	// Copy lens
	l := Lens(n)
	if l != nil {
		n.SetTag(lensTag, l)
	}

	// Copy scene
	s := Scene(n)
	if s != nil {
		if deep {
			n.SetTag(sceneTag, s.Copy())
		} else {
			n.SetTag(sceneTag, s)
		}
	}

	access := getLock(n)
	access.Lock()
	defer access.Unlock()

	// Copy regions
	regionsCopy := make(map[*util.Region]bool)
	for region, _ := range getRegions(n) {
		regionsCopy[region.Copy()] = true
	}
	n.SetTag(regionMap, regionsCopy)

	return
}

// Copy copies in-place the camera tag associated with this node.
//
// If an copy of a node is made, the underlying camera object is still
// identical. You must then use:
//
//  camera.Copy(copiedNode)
//
func Copy(n *scene.Node) {
	makeCopy(n, false)
}

// DeepCopy is just like Copy(), except it also copies it's underlying scene
// node, this is therefor expensive.
func DeepCopy(n *scene.Node) {
	makeCopy(n, true)
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