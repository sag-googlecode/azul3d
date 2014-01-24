// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

import (
	"azul3d.org/scene/geom"
)

type backToFrontSorter struct{}

// implements Sorter interface. Sorts node's based on their integer sort value
// specified by the SetSort() function.
func (s *backToFrontSorter) Less(cam, i, j Node) bool {
	iBounds, ok := geom.LocalBounds(i.Node())
	if !ok {
		geom.CalculateBounds(i.Node())
		iBounds, _ = geom.LocalBounds(i.Node())
	}

	jBounds, ok := geom.LocalBounds(j.Node())
	if !ok {
		geom.CalculateBounds(j.Node())
		jBounds, _ = geom.LocalBounds(j.Node())
	}

	// Find the center of each bounding box in camera space.
	camParent := cam.Node().Parent()

	iRelCam := camParent.RelativeTransform(i.Node()).Mat4()
	iCenter := iBounds.Center().TransformMat4(iRelCam)

	jRelCam := camParent.RelativeTransform(j.Node()).Mat4()
	jCenter := jBounds.Center().TransformMat4(jRelCam)

	camPos := cam.Node().PosVec3()
	return camPos.Sub(iCenter).LengthSquared() > camPos.Sub(jCenter).LengthSquared()
}

// NewBackToFrontSorter returns a new back-to-front sorter which sorts nodes
// based on the distance between the node's local bounding box center to the
// camera node.
func NewBackToFrontSorter() Sorter {
	return new(backToFrontSorter)
}
