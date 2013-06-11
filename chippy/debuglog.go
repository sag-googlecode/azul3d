// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build chippy_debug

package chippy

import (
	"fmt"
)

func debug(params ...interface{}) {
	params = append([]interface{}{"Chippy:"}, params...)
	fmt.Println(params...)
}
