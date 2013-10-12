// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

// SetName sets the string name of this node.
func (n *Node) SetName(name string) {
	n.access.Lock()
	defer n.access.Unlock()

	n.name = name
}

// Name returns the string name of this node.
func (n *Node) Name() string {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.name
}
