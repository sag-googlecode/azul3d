// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

func (n *Node) SetName(name string) {
	n.access.Lock()
	defer n.access.Unlock()

	n.name = name
}

func (n *Node) Name() string {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.name
}
