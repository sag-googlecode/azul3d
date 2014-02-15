// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

//import "log"

// SetTransform sets the transform that is in use by this node.
func (n *Node) SetTransform(t *Transform) {
	n.access.Lock()
	defer n.access.Unlock()

	n.transform = t
}

// Transform returns the transform that is currently in use by this node.
func (n *Node) Transform() *Transform {
	n.access.RLock()
	defer n.access.RUnlock()

	return n.transform
}

// ResetTransform is short hand for:
//
//  n.Transform().Reset()
//
func (n *Node) ResetTransform() {
	n.Transform().Reset()
}

// SetRelativeTransform sets the transform of this node to the specified one in
// the coordinate space of the other node.
func (n *Node) SetRelativeTransform(other *Node, t *Transform) {
	// Convert the transform to our parent's space.
	//rel := other.RelativeTransform(n)

	var rel *Transform

	parent := n.Parent()
	if parent != nil {
		rel = other.RelativeTransform(parent)
	} else {
		rel = other.RelativeTransform(n)
	}

	// Set our transform to the relative one.
	n.SetTransform(rel.Compose(t))
}

// RelativeTransform returns the transform of this node in the coordinate space
// of the other node.
//
// If the other node is nil; a panic will occur.
func (n *Node) RelativeTransform(other *Node) *Transform {
	if other == nil {
		panic("RelativeTransform(): other node is nil!")
	}

	logn := func(args ...interface{}) {
		//if other.Name() != "renderer" {
		//	log.Println(args...)
		//}
	}

	logn("\n\nRelativeTransform", n, other)

	// Determine the closest parent node that both our node and the other node
	// have in common.
	common := n.FindCommonParent(other)
	logn("common", common)

	logn("MMUL", n)
	m := n.Transform().Mat4()

	// Take our local transformation and multiply all of our parent's matrices
	// up to the common node.
	for _, parent := range n.Parents() {
		if common != nil && common == parent {
			logn("MUL STOP", parent)
			break
		}

		logn("MUL", parent)
		m = m.Mul(parent.Transform().Mat4())
	}

	// The matrix m is now in common space, as such we can multiply it by
	// each parent's transformation matrix (starting at the common parent)
	// descending downwards to our node itself.
	started := false
	oParents := other.Parents()
	for i := len(oParents) - 1; i >= 0; i-- {
		oParent := oParents[i]

		if started {
			logn("INV", oParent)
			inv, _ := oParent.Transform().Mat4().Inverse()
			m = m.Mul(inv)
		} else {
			logn("INV IGNORE", oParent)
			if common == nil {
				started = true
			} else if oParent == common {
				started = true
			}
		}
	}

	logn("INV", other)
	inv, _ := other.Transform().Mat4().Inverse()
	m = m.Mul(inv)

	// The matrix m is now in our node's local space, success! Simply build an
	// transform object and we're good to go.
	localTransform := new(Transform)
	localTransform.SetMat4(m)

	return localTransform
}
