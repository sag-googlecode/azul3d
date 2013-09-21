// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"fmt"
)

const (
	debug = false
)

func logf(args ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(args[0].(string), args[1:]...))
}
