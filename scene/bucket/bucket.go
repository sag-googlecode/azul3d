// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

import (
	"azul3d.org/v0/scene"
	"fmt"
	"sync"
)

// Type represents a single bucket type. It has a integer sort value that
// represents the render order of this bucket (see the SetSort() method for
// more information).
type Type struct {
	Sorter
	access sync.RWMutex
	name   string
	sort   int
}

// SetSort sets the integer sort value of this bucket type. A lower value would
// have all the nodes in this bucket rendered before all the nodes in other
// buckets, likewise a higher value would have all the nodes in this bucket
// rendered after all the nodes in other buckets.
func (t *Type) SetSort(sortValue int) {
	t.access.Lock()
	defer t.access.Unlock()

	t.sort = sortValue
}

// Sort returns the integer sort value of this bucket type. See the SetSort()
// method for more information.
func (t *Type) Sort() int {
	t.access.RLock()
	defer t.access.RUnlock()

	return t.sort
}

// Name returns the name of the bucket type, as was specified at creation time.
func (t *Type) Name() string {
	t.access.RLock()
	defer t.access.RUnlock()

	return t.name
}

// String returns a string representation of this bucket type, including name
// and sort value.
func (t *Type) String() string {
	return fmt.Sprintf("bucket.Type(%q, Sort=%v)", t.Name(), t.Sort())
}

// New creates a new bucket type with a given name, integer sort value and a
// sorter to use for sorting the bucket's associated nodes.
func New(name string, sort int, sorter Sorter) *Type {
	return &Type{
		sorter,
		sync.RWMutex{},
		name,
		sort,
	}
}

// PBucket is the property for storing a node's bucket type.
var PBucket = scene.NewProp("Bucket")

// Set sets the bucket of the given node to the specified one.
func Set(n *scene.Node, bucket *Type) {
	n.SetProp(PBucket, bucket)
}

// Get returns the bucket of the given node or Opaque if it has none.
func Get(n *scene.Node) *Type {
	i, ok := n.Prop(PBucket)
	if !ok {
		return Opaque
	}
	return i.(*Type)
}

// Active returns the active bucket of the given node or Opaque if it has none.
func Active(n *scene.Node) *Type {
	i, ok := n.ActiveProp(PBucket)
	if !ok {
		return Opaque
	}
	return i.(*Type)
}

// SetForced sets if the bucket of this node should be forced as the active
// bucket instead of obeying parent node wishes.
func SetForced(n *scene.Node, forced bool) {
	n.SetPropForced(PBucket, forced)
}

// Forced tells if the bucket of this node is forced as being actively used
// instead of obeying parent node wishes.
func Forced(n *scene.Node) bool {
	return n.PropForced(PBucket)
}

// Clear clears the bucket of the given node.
func Clear(n *scene.Node) {
	n.ClearProp(PBucket)
}

var (
	// Nodes in the background (only *visually* in the background if depth
	// testing is turned off for the nodes). These nodes are rendered first and
	// are sorted based on their integer sort values since this is a value
	// sorter.
	Background = New("Background", 0, NewValueSorter())

	// Opaque nodes (i.e. nodes without transparent parts), these nodes are
	// drawn directly after nodes in the Background bucket and are state sorted
	// to improve overall rendering time.
	Opaque = New("Opaque", 1000, NewStateSorter())

	// Transparent nodes (i.e. nodes with at least some transparent parts),
	// these nodes are rendered directly after nodes in the Opaque bucket and
	// are sorted back-to-front based off distance to the camera node so that
	// alpha blending works as expected.
	Transparent = New("Transparent", 2000, NewBackToFrontSorter())

	// Sorted nodes (i.e. nodes with a specific integer sort value set via the
	// SetSort() function), these nodes are rendered directly after nodes in
	// the Transparent bucket and are sorted based off each node's active
	// integer sort value.
	Sorted = New("Sorted", 3000, NewValueSorter())

	// Unsorted nodes (i.e. nodes that are specifically not sorted in any
	// special way at all), these nodes are rendered directly after nodes in
	// the Sorted bucket and are not sorted at all, as such they are drawn in
	// the order in which they appear in the scene.
	Unsorted = New("Unsorted", 4000, NewNilSorter())
)
