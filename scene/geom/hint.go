// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package geom

import (
	"fmt"
)

// Hint represents a single mesh usage hint, used to 'hint' at the intended
// usage or dynamic-ness of a mesh.
//
// The chosen hint does not restrict how a mesh might be used, but can make it
// less optimal performance-wise.
type Hint uint8

const (
	// An Geom who is expected to never change.
	Static Hint = iota

	// An Geom who is expected to change often.
	Dynamic

	// An Geom who is expected to change every single frame.
	Stream
)

// String returns an string representation of this UsageHint type.
func (hint Hint) String() string {
	switch hint {
	case Static:
		return "Static"
	case Dynamic:
		return "Dynamic"
	case Stream:
		return "Stream"
	}
	return fmt.Sprintf("%d", hint)
}
