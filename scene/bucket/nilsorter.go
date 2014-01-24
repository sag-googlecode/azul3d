// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

type nilSorter struct{}

// implements Sorter interface. Does no sorting at all.
func (s *nilSorter) Less(cam, i, j Node) bool {
	return false
}

// NewNilSorter returns a new nil sorter which does not sort nodes at all, as
// such they are drawn in the order in which they appear in the scene graph.
func NewNilSorter() Sorter {
	return new(nilSorter)
}
