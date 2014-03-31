// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"azul3d.org/v1/scene"
)

var (
	// The property for storing the local bounds of a node.
	PLocalBounds = scene.NewProp("LocalBounds")

	// The property for storing the bounds of a node.
	PBounds = scene.NewProp("Bounds")
)

func testBounds(b1, b2 *BoundingBox) (r *BoundingBox) {
	r = new(BoundingBox)

	if b1.Min.X < b2.Min.X {
		r.Min.X = b1.Min.X
	} else if b1.Min.X > b2.Min.X {
		r.Min.X = b2.Min.X
	}

	if b1.Min.Y < b2.Min.Y {
		r.Min.Y = b1.Min.Y
	} else if b1.Min.Y > b2.Min.Y {
		r.Min.Y = b2.Min.Y
	}

	if b1.Min.Z < b2.Min.Z {
		r.Min.Z = b1.Min.Z
	} else if b1.Min.Z > b2.Min.Z {
		r.Min.Z = b2.Min.Z
	}

	if b1.Max.X > b2.Max.X {
		r.Max.X = b1.Max.X
	} else if b1.Max.X > b2.Max.X {
		r.Max.X = b2.Max.X
	}

	if b1.Max.Y > b2.Max.Y {
		r.Max.Y = b1.Max.Y
	} else if b1.Max.Y > b2.Max.Y {
		r.Max.Y = b2.Max.Y
	}

	if b1.Max.Z > b2.Max.Z {
		r.Max.Z = b1.Max.Z
	} else if b1.Max.Z > b2.Max.Z {
		r.Max.Z = b2.Max.Z
	}

	return
}

func calculateBounds(n *scene.Node, done chan bool) {
	// Calculate child bounding boxes
	children := n.Children()
	childBoundsCompleted := make([]chan bool, len(children))
	for i, child := range children {
		ch := make(chan bool, 1)
		childBoundsCompleted[i] = ch
		go calculateBounds(child, ch)
	}

	// Calculate our local bounding box while children bounding boxes are being
	// calculated.
	var lbb *BoundingBox
	for _, m := range Meshes(n) {
		bb := m.BoundingBox()
		if bb == nil {
			m.CalculateBounds()
		}
		bb = m.BoundingBox()
		if bb != nil {
			if lbb == nil {
				lbb = bb.Copy()
			} else {
				lbb = testBounds(bb, lbb)
			}
		}
	}

	// Assign local bounding box
	obb, ok := LocalBounds(n)
	n.SetProp(PLocalBounds, lbb)

	// Keep track of any changes to our local or children bounding boxes
	anyBoundsChanged := false
	if !ok || obb == nil || !obb.Equals(lbb) {
		anyBoundsChanged = true
	}

	// Wait for child bounding boxes to finish calculating
	for _, ch := range childBoundsCompleted {
		changed := <-ch
		if changed {
			anyBoundsChanged = true
		}
	}

	// If our local bounding box or children bounding boxes have changed, then
	// we need to update our "bounding box" which encapsulates this node's
	// meshes and all nodes below it's meshes.
	if anyBoundsChanged {
		bb, _ := LocalBounds(n)

		// Extend bounding box by all children bounding boxes
		for _, child := range children {
			cbb, _ := LocalBounds(child)
			if cbb != nil {
				if bb == nil {
					bb = new(BoundingBox)
				}
				bb = testBounds(cbb, bb)
			}
		}

		n.SetProp(PBounds, bb)
	}

	done <- true
}

// CalculateBounds calculates the bounding box of the specified root node and
// all child nodes below it.
func CalculateBounds(root *scene.Node) {
	ch := make(chan bool)
	go calculateBounds(root, ch)
	<-ch
}

// LocalBounds returns a axis-aligned bounding box which encapsulates this
// node's meshes.
//
// If this function returns ok=false, then there is no local bounding box yet
// calculated (and CalculateBounds() should be invoked).
//
// If the returned bounding box is nil, then there is no local bounding box
// (I.e. no meshes) at this node.
func LocalBounds(n *scene.Node) (bb *BoundingBox, ok bool) {
	i, ok := n.Prop(PLocalBounds)
	if ok && i != nil {
		bb = i.(*BoundingBox)
	}
	return
}

// Bounds returns a axis-aligned bounding box which encapsulates this
// node's meshes and all children's (and distant children's) meshes.
//
// If this function returns ok=false, then there is no bounding box yet
// calculated (and CalculateBounds() should be invoked).
//
// If the returned bounding box is nil, then there is no bounding box (I.e. no
// meshes) at or below this node.
func Bounds(n *scene.Node) (bb *BoundingBox, ok bool) {
	i, ok := n.Prop(PBounds)
	if ok && i != nil {
		bb = i.(*BoundingBox)
	}
	return
}
