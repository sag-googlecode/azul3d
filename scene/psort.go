// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

var (
	PSort = NewProp("Sort")
)

// SetSort specifies the sort value of this node; an node with an smaller sort
// value is drawn before nodes with larger sort values.
//
// An node's sort value is only used given that it's sorter is the predefined
// sorter, Fixed.
//
// The default sort value is 0.
func (n *Node) SetSort(sort int) {
	n.SetProp(PSort, sort)
}

// Sort returns the sort value of this node; an node with an smaller sort value
// is drawn before nodes with larger sort values.
//
// An node's sort value is only used given that it's sorter is the predefined
// sorter, Fixed.
//
// The default sort value is 0.
func (n *Node) Sort() (sort int, ok bool) {
	i, ok := n.Prop(PSort)
	if ok {
		sort = i.(int)
	}
	return
}

// ActiveSort returns the active sort value of this node, as affected by parent
// node sort values.
func (n *Node) ActiveSort() (sort int, ok bool) {
	i, ok := n.ActiveProp(PSort)
	if ok {
		sort = i.(int)
	}
	return
}

// ClearSort clears the sort value of this node.
func (n *Node) ClearSort() {
	n.ClearProp(PSort)
}

// SetSortForced specifies if the sort value of this node should be forced
// to be actively used instead of obeying the parent node wishes.
func (n *Node) SetSortForced(forced bool) {
	n.SetPropForced(PSort, forced)
}

// SortForced tells if the sort value of this node is forced as being actively
// used instead of obeying the parent node wishes.
func (n *Node) SortForced() bool {
	return n.PropForced(PSort)
}
