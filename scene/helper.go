// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

func (n *Node) New(name string) *Node {
	child := New(name)
	child.SetParent(n)
	return child
}
