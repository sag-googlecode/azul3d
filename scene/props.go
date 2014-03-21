// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

// Prop represents a unique named property. Properties themselves are just
// interface{} values, but since Prop is a pointer type, it is guaranteed to be
// runtime unique, as such it cannot collide with other named properties (e.g.
// ones whose type are just "a string").
type Prop string

// String returns the dereferenced string. Aka. the name of the property.
func (u *Prop) String() string {
	return string(*u)
}

// NewProp returns a new unique named property. The returned property may be
// set on nodes while knowing that the property cannot collide with other ones
// because it is a pointer type (instead of just a string type).
func NewProp(name string) *Prop {
	p := Prop(name)
	return &p
}

type mapLookupPair struct {
	value interface{}
	ok    bool
}

func (n *Node) saveActiveProp(prop interface{}, pair mapLookupPair) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.activePropCache == nil {
		n.activePropCache = make(map[interface{}]mapLookupPair)
	}
	n.activePropCache[prop] = pair
}

func (n *Node) loadActiveProp(prop interface{}) (pair mapLookupPair, ok bool) {
	n.access.RLock()
	defer n.access.RUnlock()

	if n.activePropCache == nil {
		return mapLookupPair{}, false
	}
	pair, ok = n.activePropCache[prop]
	return
}

func (n *Node) recursiveClearActiveProps() {
	n.access.Lock()
	defer n.access.Unlock()

	n.doRecursiveClearActiveProps()
}

func (n *Node) doRecursiveClearActiveProps() {
	n.activePropCache = nil
	n.parents = nil
	for _, child := range n.children {
		child.recursiveClearActiveProps()
	}
}

// ActiveProp returns the active property value for this node.
//
// This means the value that this node would use, as affected by whether or not
// parent node's have the property specified (essentially overriding this child
// nodes wishes), or whether or not SetPropForced(prop, true) was used on this
// node, etc.
func (n *Node) ActiveProp(prop interface{}) (value interface{}, ok bool) {
	pair, havePair := n.loadActiveProp(prop)
	if havePair {
		return pair.value, pair.ok
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
			n.saveActiveProp(prop, mapLookupPair{value, ok})
			return
		}
	}

	// It seems that no parents have the given property at this point. We can
	// see if this node itself has the property, though.
	value, ok = n.Prop(prop)
	n.saveActiveProp(prop, mapLookupPair{value, ok})

	return
}

// SetProp specifies the value of the given property on this node. The 'prop'
// parameter is the property name, which may be any interface{} type which
// defines equality (as properties are stored internally in a map).
//
// The value may be later retrieved using Prop(), or ActiveProp() with the same
// property name.
func (n *Node) SetProp(prop, value interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	if n.props == nil {
		n.props = make(map[interface{}]interface{})
	}
	n.props[prop] = value

	n.doRecursiveClearActiveProps()
}

// Prop returns the property value of the named property for this node.
//
// Unlike ActiveProp(), this method is not affected by parent nodes or anything
// else. As such, it operates completely locally on this node only.
func (n *Node) Prop(prop interface{}) (value interface{}, ok bool) {
	n.access.RLock()
	defer n.access.RUnlock()

	if n.props == nil {
		return nil, false
	}
	value, ok = n.props[prop]
	return
}

// ClearProp clears the specified property value on this node. If this node
// does not have a value for the given property, then this method is no-op.
func (n *Node) ClearProp(prop interface{}) {
	n.access.Lock()
	defer n.access.Unlock()

	delete(n.props, prop)
	n.doRecursiveClearActiveProps()
}

// Props returns a map of property names and values for this node.
func (n *Node) Props() map[interface{}]interface{} {
	n.access.RLock()
	defer n.access.RUnlock()

	c := make(map[interface{}]interface{})
	for prop, value := range n.props {
		c[prop] = value
	}
	return c
}

// ClearProps clears all the values of all properties on this node.
func (n *Node) ClearProps() {
	n.access.Lock()
	defer n.access.Unlock()

	n.props = nil
	n.doRecursiveClearActiveProps()
}

// SetPropForced specifies the property should be explicitly forced to be
// actively used regardless of parent node property wishes.
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
	n.doRecursiveClearActiveProps()
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
