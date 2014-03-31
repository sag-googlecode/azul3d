// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

import "azul3d.org/v1/scene"

// Sorter works on this node interface, so that the underlying type can
// store other node information (e.g. a renderer could store information
// without touching the node's data directly).
type Node interface {
	// Should return the actual scene node that this interface represents.
	Node() *scene.Node
}

// Sorter represents a sorter for a bucket of nodes. The Less() method must be
// safe to call from multiple goroutines at the same time.
//
// The interface does not directly handle the slice of nodes in the bucket for
// efficiency reasons.
type Sorter interface {
	// Less should return a boolean value specifying if the node i in the
	// bucket should render (sort) before the node j.
	//
	// The 'cam' parameter is the camera whose scene the two nodes, i and j,
	// reside under. This parameter is useful for back-to-front sorting by
	// comparing the distance to the camera node, etc.
	Less(cam, i, j Node) bool
}
