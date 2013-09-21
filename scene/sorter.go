// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"fmt"
	"sync"
)

// Sorter represents an single sorter who describes how and in what order to
// sort an tree of nodes in order to prepare them for rendering.
type Sorter struct {
	access sync.RWMutex

	name         string
	sort         uint
	sortSortType SortType
}

// String returns an string representation of this sorter.
func (s *Sorter) String() string {
	return fmt.Sprintf("Sorter(%q, Sort=%v, SortType=%v)", s.Name(), s.Sort(), s.SortType())
}

// SetName sets the name of this sorter.
func (s *Sorter) SetName(name string) {
	s.access.Lock()
	defer s.access.Unlock()

	s.name = name
}

// Name returns the name of this sorter.
func (s *Sorter) Name() string {
	s.access.RLock()
	defer s.access.RUnlock()

	return s.name
}

// SetSort sets the sort value of this sorter.
//
// An low sort value causes this sorter's affected nodes to be drawn first, and
// an higher value causes them to be drawn last (of course, with relation to
// other sorters and their sort values).
func (s *Sorter) SetSort(sort uint) {
	s.access.Lock()
	defer s.access.Unlock()

	s.sort = sort
}

// Sort returns the sort value of this sorter.
//
// An low sort value causes this sorter's affected nodes to be drawn first, and
// an higher value causes them to be drawn last (of course, with relation to
// other sorters and their sort values).
func (s *Sorter) Sort() uint {
	s.access.RLock()
	defer s.access.RUnlock()

	return s.sort
}

// SetSortType sets the sort SortType to be used by this sorter.
//
// This must be an valid sort SortType (I.e. one of the predefined constants)
// or an panic will occur.
func (s *Sorter) SetSortType(sortSortType SortType) {
	if !sortSortType.Valid() {
		panic("SetSortType(): Invalid sort SortType specified!")
	}

	s.access.Lock()
	defer s.access.Unlock()

	s.sortSortType = sortSortType
}

// SortType returns the sort SortType that is in use by this sorter.
func (s *Sorter) SortType() SortType {
	s.access.RLock()
	defer s.access.RUnlock()

	return s.sortSortType
}

// NewSorter returns an new initialized *Sorter given an name, sort value, and
// sort SortType.
//
// Normal applications will not need to create their own sorters, and will
// typically use the ones predefined in this package.
func NewSorter(name string, sort uint, sortSortType SortType) *Sorter {
	if !sortSortType.Valid() {
		panic("NewSorter(): Invalid sort SortType specified.")
	}

	return &Sorter{
		name:         name,
		sort:         sort,
		sortSortType: sortSortType,
	}
}

var (
	// Nodes are drawn based on the Sort value specified on the node (via the
	// SetSort() function), because of this sorter's low sort value of 1000,
	// these objects will always be drawn in the background of other sorters
	// (unless new ones are made with lower values outside azul3d).
	Background = NewSorter("Background", 1000, ValueSort)

	// Nodes are drawn based on their state transitions, such that nodes with
	// similar state are drawn after eachother in an attempt to minimize state
	// switches.
	//
	// State transitions include anything which would incur an OpenGL or
	// Direct3D state transition, meaning coloring, etc.
	State = NewSorter("State", 2000, StateSort)

	// Nodes are drawn according to the center of the nodes' bounding volume,
	// such that the nodes who are further away from the camera are drawn
	// first.
	//
	// This is the proper sorter for objects who use transparency and occlude
	// themselves.
	BackToFront = NewSorter("BackToFront", 3000, BackToFrontSort)

	// Nodes are drawn based on the Sort value specified on the node (via the
	// SetSort() function).
	Fixed = NewSorter("Fixed", 4000, ValueSort)

	// Nodes are drawn in the order in which they appear in the scene graph;
	// in node traversal order (See the Traverse() function for more
	// information).
	Unsorted = NewSorter("Unsorted", 5000, TraversalSort)
)
