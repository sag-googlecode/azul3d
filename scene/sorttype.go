// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

// SortType represents an specific sort type.
//
// For an sort type to be considered valid, it must be one of the predefined
// constants in this package.
type SortType uint8

// String returns an string representation of this sort type.
func (s SortType) String() string {
	switch s {
	case InvalidSort:
		return "InvalidSort"
	case ValueSort:
		return "ValueSort"
	case StateSort:
		return "StateSort"
	case BackToFrontSort:
		return "BackToFrontSort"
	case TraversalSort:
		return "TraversalSort"
	}
	return "InvalidSort"
}

// Valid tells weather this sort type is valid.
//
// For an sort type to be considered valid, it must be one of the predefined
// constants in this package.
func (s SortType) Valid() bool {
	switch s {
	case InvalidSort:
		return false
	case ValueSort:
		return true
	case StateSort:
		return true
	case BackToFrontSort:
		return true
	case TraversalSort:
		return true
	}
	return false
}

const (
	// InvalidSort is defined only for purposes of catching zero-value errors.
	// Do not use.
	InvalidSort SortType = iota

	// Nodes are drawn based on the Sort value specified on the node (via the
	// node.SetSort() function).
	ValueSort

	// Nodes are drawn based on their state transitions, such that nodes with
	// similar state are drawn after eachother in an attempt to minimize state
	// switches.
	StateSort

	// Nodes are drawn according to the center of the nodes' bounding volume,
	// such that the nodes who are further away from the camera are drawn
	// first.
	BackToFrontSort

	// Nodes are drawn in the order in which they appear in the scene graph; in
	// node traversal order (See the node.Traverse() function for more
	// information).
	TraversalSort
)
