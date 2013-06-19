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
	specialProcedures["glGetActiveAttrib"] = func(ctx, prefix string, p *Procedure) (name, args, body, returns string) {
		body = fmt.Sprintf(`
	var (
		cname C.GLchar
	)

	C.%sGetActiveAttrib(%s, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
	name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
	return
`, prefix, ctx)

		return "GetActiveAttrib", "program, index uint32, bufSize int32", body, "(length int32, size int32, Type uint32, name string)"
	}
}