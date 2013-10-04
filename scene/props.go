// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

type Prop string

func (u *Prop) String() string {
	return string(*u)
}

func NewProp(name string) *Prop {
	p := Prop(name)
	return &p
}

func (n *Node) saveActiveProp(prop, value interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.activePropCache == nil {
		n.activePropCache = make(map[interface{}]interface{})
	}
	n.activePropCache[prop] = value
}

func (n *Node) loadActiveProp(prop interface{}) (value interface{}, ok bool) {
	return nil, false

	n.access.Lock()
	defer n.access.Unlock()

	if n.activePropCache == nil {
		return nil, false
	}
	value, ok = n.activePropCache[prop]
	return
}

func (n *Node) doClearActiveProp(prop interface{}) {
	delete(n.activePropCache, prop)
}

func (n *Node) doClearActiveProps() {
	n.activePropCache = nil
	n.parents = nil
}

func (n *Node) clearActiveProps() {
	n.access.Lock()
	defer n.access.Unlock()

	n.doClearActiveProps()
}

func (n *Node) doRecursiveClearActiveProps() {
	n.doClearActiveProps()
	for _, child := range n.children {
		child.doRecursiveClearActiveProps()
	}
}

func (n *Node) ActiveProp(prop interface{}) (value interface{}, ok bool) {
	value, ok = n.loadActiveProp(prop)
	if ok {
		return
	}

	// If we are forcing the property to an value, then return that now.
	if n.PropForced(prop) {
		value, ok = n.Prop(prop)
		return
	}

	parents := n.Parents()

	// Search in child-to-parent order through all parents nodes and see if any
	// have the property forced.
	nearestForcedProp := len(parents)
	for i, parent := range parents {
		if parent.PropForced(prop) {
			nearestForcedProp = i
			break
		}
	}

	// Search in parent-to-child order through all parent nodes starting at the
	// nearest forced parent to see if any have the given property.
	for i := nearestForcedProp; i > 0; i-- {
		parent := parents[i-1]
		value, ok = parent.Prop(prop)
		if ok {
			n.saveActiveProp(prop, value)
			return
		}
	}

	// It seems that no parents have the given property at this point. We can
	// see if this node itself has the property, though.
	value, ok = n.Prop(prop)
	if ok {
		n.saveActiveProp(prop, value)
	}

	return
}

func (n *Node) SetProp(prop, value interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.props == nil {
		n.props = make(map[interface{}]interface{})
	}
	n.props[prop] = value

	n.doClearActiveProp(prop)
}

func (n *Node) Prop(prop interface{}) (value interface{}, ok bool) {
	n.access.RLock()
	defer n.access.RUnlock()

	if n.props == nil {
		return nil, false
	}
	value, ok = n.props[prop]
	return
}

func (n *Node) ClearProp(prop interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	delete(n.props, prop)
	n.doClearActiveProp(prop)
}

func (n *Node) Props() map[interface{}]interface{} {
	n.access.RLock()
	defer n.access.RUnlock()

	c := make(map[interface{}]interface{})
	for prop, value := range n.tags {
		c[prop] = value
	}
	return c
}

func (n *Node) ClearProps() {
	n.access.Lock()
	defer n.access.Unlock()

	n.props = nil
	n.doClearActiveProps()
}

// PropForced specifies the property should be explicitly forced to be actively
// used regardless of parent node property wishes.
func (n *Node) SetPropForced(prop interface{}, forced bool) {
	n.access.Lock()
	defer n.access.Unlock()

	if forced == false {
		delete(n.forcedProps, prop)
	} else {
		if n.forcedProps == nil {
			n.forcedProps = make(map[interface{}]bool)
		}
		n.forcedProps[prop] = forced
	}
}

// PropForced tells if the specified property is explicitly forced to be
// actively used regardless of parent node property wishes.
func (n *Node) PropForced(prop interface{}) bool {
	n.access.RLock()
	defer n.access.RUnlock()

	if n.forcedProps == nil {
		return false
	}
	forced, ok := n.forcedProps[prop]
	if !ok {
		return false
	}
	return forced
}
