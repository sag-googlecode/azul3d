// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package bucket

import "azul3d.org/scene"

var PSort = scene.NewProp("Sort")

// SetSort sets an integer sort value for the given node. The value is only
// taken into account if the node's bucket sorter is a value sorter returned by
// NewValueSorter() -- that is other sorters (state sorters, distance sorters,
// etc) do not take this value into account.
//
// A low value would make the node render first and a high value would make it
// render last.
func SetSort(n *scene.Node, sortValue int) {
	n.SetProp(PSort, sortValue)
}

// Sort returns the integer sort value for the given node. See the SetSort()
// function for more information.
func Sort(n *scene.Node) int {
	i, ok := n.Prop(PSort)
	if !ok {
		return 0
	}
	return i.(int)
}

// ActiveSort returns the active integer sort value for the given node. See the
// SetSort() function for more information.
func ActiveSort(n *scene.Node) int {
	i, ok := n.ActiveProp(PSort)
	if !ok {
		return 0
	}
	return i.(int)
}

// SetSortForced sets if the sort value of this node should be forced as the
// active sort value instead of obeying parent node wishes.
func SetSortForced(n *scene.Node, forced bool) {
	n.SetPropForced(PSort, forced)
}

// SortForced tells if the sort value of this node is forced as being actively
// used instead of obeying parent node wishes.
func SortForced(n *scene.Node) bool {
	return n.PropForced(PSort)
}

// ClearSort clears the sort value of the given node (equivilent to zero).
func ClearSort(n *scene.Node) {
	n.ClearProp(PSort)
}
