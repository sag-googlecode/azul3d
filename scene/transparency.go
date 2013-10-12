// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"fmt"
)

// TransparencyMode describes a single type of transparency that can be used
// during rendering.
type TransparencyMode int

// String returns a string representation of this transparency mode.
//
// E.g. NoTransparency -> "NoTransparency"
func (t TransparencyMode) String() string {
	switch t {
	case NoTransparency:
		return "NoTransparency"
	case Transparency:
		return "Transparency"
	case Binary:
		return "Binary"
	case Multisample:
		return "Multisample"
	}
	return fmt.Sprintf("TransparencyMode(%d)", t)
}

// Valid tells if this transparency mode is a valid one (I.e one of the
// pre-defined constants in this package).
func (t TransparencyMode) Valid() bool {
	switch t {
	case NoTransparency:
		return true
	case Transparency:
		return true
	case Binary:
		return true
	case Multisample:
		return true
	}
	return false
}

const (
	// NoTransparency means the node should be drawn with no transparency at
	// all. If transparent pixels would exist on this node, then they will be
	// shown only as black pixels.
	NoTransparency TransparencyMode = iota

	// Transparency means the node should be drawn with default transparency.
	// This works well for many objects, but in some cases (such as self
	// occluding transparent objects) causes ordering issues.
	Transparency

	// Binary means the node should be drawn with binary transparency, where
	// an alpha of less than
	// 0.5 means 'solid', and an alpha value greater than or equal to 0.5 means
	// 'clear'.
	//
	// Binary transparency suffers from jagged edges (as there is only solid or
	// fully opaque pixels that can exist), but suffers from no ordering
	// issues.
	Binary

	// Multisample transparency means the node should be drawn with multisample
	// transparency, this form of transparency has no ordering issues at all
	// and is considered 'true' or 'per-pixel' transparency.
	//
	// As this only works on high-end cards who have support for multisample
	// transparency, if multisample transparency is not supported, the
	// fallback will be Binary transparency as it also suffers from no ordering
	// issues (although it causes jagged edges).
	Multisample
)
