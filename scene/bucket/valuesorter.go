// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

type valueSorter struct{}

// implements Sorter interface. Sorts node's based on their integer sort value
// specified by the SetSort() function.
func (s *valueSorter) Less(cam, i, j Node) bool {
	return ActiveSort(i.Node()) < ActiveSort(j.Node())
}

// NewValueSorter returns a new value sorter which sorts nodes based on their
// active integer sort values set by the SetSort() function.
func NewValueSorter() Sorter {
	return new(valueSorter)
}
