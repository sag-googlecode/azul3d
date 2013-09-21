// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

func (n *Node) SetTag(tag, value interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.tags == nil {
		n.tags = make(map[interface{}]interface{})
	}
	n.tags[tag] = value
}

func (n *Node) Tag(tag interface{}) (value interface{}, ok bool) {
	n.access.RLock()
	defer n.access.RUnlock()

	if n.tags == nil {
		return nil, false
	}
	value, ok = n.tags[tag]
	return
}

func (n *Node) ClearTag(tag interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	delete(n.tags, tag)
}

func (n *Node) Tags() map[interface{}]interface{} {
	n.access.RLock()
	defer n.access.RUnlock()

	c := make(map[interface{}]interface{})
	for tag, value := range n.tags {
		c[tag] = value
	}
	return c
}

func (n *Node) ClearTags() {
	n.access.Lock()
	defer n.access.Unlock()

	n.tags = nil
}
