// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// Code generated by this program is also under
// the above license.

package generator

import (
	"fmt"
)

func init() {
	specialProcedures["glBegin"] = func(ctx, prefix string, p *Procedure) (name, args, body, returns string) {
		body = fmt.Sprintf(`

	glc.inBeginEnd = true
	C.%sBegin(%s, C.GLenum(mode))
	return
`, prefix, ctx)

		return "Begin", "mode uint32", body, ""
	}
}
