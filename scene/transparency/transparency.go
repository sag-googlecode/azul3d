// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package transparency

import (
	"fmt"
)

// ModeType describes a single transparency mode that can be used on a node for
// drawing.
type ModeType int

// String returns a string representation of this transparency mode.
//
// E.g. AlphaBlend -> "AlphaBlend"
func (t ModeType) String() string {
	switch t {
	case None:
		return "None"
	case AlphaBlend:
		return "AlphaBlend"
	case Binary:
		return "Binary"
	case Multisample:
		return "Multisample"
	}
	return fmt.Sprintf("Mode(%d)", t)
}

// Valid tells if this transparency mode is a valid one (I.e one of the
// pre-defined constants in this package).
func (t ModeType) Valid() bool {
	switch t {
	case None:
		return true
	case AlphaBlend:
		return true
	case Binary:
		return true
	case Multisample:
		return true
	}
	return false
}

const (
	// None means the node should be drawn with no transparency at all. This is
	// the default mode. If transparent pixels would exist on the drawn node,
	// then they will show up as black pixels only.
	None ModeType = iota

	// AlphaBlend means the node's transparent parts should be drawn using
	// alpha-blending. This type of transparency works well for some (but not
	// all) objects.
	//
	// Pros:
	//     Pixels can be semi-transparent.
	//
	// Cons:
	//     Render order dependant. Nodes must be sorted such that opaque nodes
	//     are drawn before transparent ones (because alpha blending literally
	//     blends pixels with the ones already written to the framebuffer).
	//
	//     Does not work well with self-occluding transparent objects (e.g. a
	//     cube where all faces are semi-transparent) because individual faces
	//     would have to be sorted for correct order -- not feasible in
	//     realtime applications.
	//
	AlphaBlend

	// Binary means the node should be drawn with binary transparency, this
	// causes transparency to be thought of as a 'binary' decision, where each
	// pixel is either fully transparent or opaque.
	//
	// Pixels with an alpha value of less than 0.5 are considered fully
	// transparent, likewise pixels with an alpha value of greater than or
	// equal to 0.5 are considered fully-opaque.
	//
	// Pros:
	//     Render order independent. Nodes do not have to be sorted in any way
	//     (unlike AlphaBlend transparency for instance).
	//
	// Cons:
	//     Jagged looking edges. Pixels may not be semi-transparent (they are
	//     either opaque or fully transparent).
	//
	Binary

	// Multisample transparency means the node should be drawn with multisample
	// transparency.
	//
	// In OpenGL this is implemented using 'SAMPLE_ALPHA_TO_COVERAGE' from the
	// "GL_ARB_multisample" extension.
	//
	// Pros:
	//     Render order independent. Nodes do not have to be sorted in any way
	//     (unlike AlphaBlend transparency for instance).
	//
	//     No jagged edges. Pixels may be semi-transparent (unlike Binary
	//     transparency for instance).
	//
	// Cons:
	//     Only some hardware supports it (in the event that hardware does not
	//     support it, Binary transparency is used because it also does not
	//     suffer from ordering issues, although it causes jagged edges).
	Multisample
)
