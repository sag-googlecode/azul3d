// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"errors"
)

var (
	// Panic that will occur when there is an attempt made to create a circular
	// reference within the scene graph.
	CircularErr = errors.New("Attempt to create circular reference in scene graph.")
)
