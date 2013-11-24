// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// Code generated by this program is also under
// the above license.

// glwrap is a tool for generating Go OpenGL wrappers.
package main

import (
	"io"
	"text/template"
)

const cCodeTemplate = `// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// This source file was automatically generated using glwrap.
// +build{{.BuildTags}}

#include "gl.h"

typedef void (*__glwrap_func_ptr)(void);

#if defined(__WIN32) || defined(__WIN64)
	// Todo: Windows support.

#elif defined(__linux) || defined(__unix) || defined(__posix)
	// See http://dri.freedesktop.org/wiki/glXGetProcAddressNeverReturnsNULL
	//
	// glXGetProcAddressARB is *required* to be statically defined in libGL,
	// but glXGetProcAddress is not, and will fail to be found in nvidia's
	// libGL
	extern __glwrap_func_ptr glXGetProcAddressARB(const GLubyte * procName);

#elif defined(__APPLE__)
	// Todo: OS X support.
#endif

inline __glwrap_func_ptr gl_wrap_get_pointer(const char* name) {
	#if defined(__WIN32) || defined(__WIN64)
		// Todo: Windows support.

	#elif defined(__linux) || defined(__unix) || defined(__posix)
		return glXGetProcAddressARB(name);

	#elif defined(__APPLE__)
		// Todo: OS X support.
	#endif
}

// Function definition for each appropriate OpenGL function.
//
// If the pointer in the context for the function is null; it is loaded
// immedietly (as such this is effectively lazy-loading).
{{range $function := .Functions}}{{.CReturns}} gl_wrap_context_{{.CName}}(gl_wrap_context* ctx{{range $e := .CArgsPairs}}, {{.Type}} {{.Name}}{{end}}) {
	if(ctx->{{.CName}}Proc == NULL) {
		ctx->{{.CName}}Proc = (PFN{{.CCapsName}}PROC)gl_wrap_get_pointer("{{.CName}}");
	}
	{{if .GoReturns}}return {{end}}ctx->{{.CName}}Proc({{range $index, $element := .CArgsPairs}}{{.Name}}{{if NeedComma $index $function.CArgsPairs}}{{", "}}{{end}}{{end}});
};
{{end}}


// Handler functions are defined for each OpenGL call; each handler function
// takes the OpenGL context struct and a pointer to the same OpenGL function's
// arguments stored in a struct.
//
// Each handler function is responsible for invoking the OpenGL function with
// the proper parameters.
//
// All function handlers are placed with respect to order in the defined jump
// table (see below), which allows batched OpenGL calls to be made without
// using a large (and costly) switch statement.
//
// Handler functions are not defined for OpenGL functions which return any
// value, as these function calls cannot be batched (see the Go documentation
// for this package, which explains this in more detail).
{{range $element := .Functions}}{{if not .GoReturns}}inline void gl_wrap_handler_{{.CName}}(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_{{.CName}}_args args = *(gl_wrap_handler_{{.CName}}_args*)argsPtr;
	gl_wrap_context_{{.CName}}(ctx{{range $e := .CArgsPairs}}, args.{{.Name}}{{end}});
}

{{end}}{{end}}// This is the jump table used for executing each batched OpenGL function
// without doing a large (approx. 1k cases) and costly switch statement.
gl_wrap_jump_handler gl_wrap_jump_table[] = {
{{range $element := .Functions}}{{if not .GoReturns}}	gl_wrap_handler_{{.CName}},
{{end}}{{end}}{{"};"}}

// Executes the functions in a batch for a given context.
void gl_wrap_batch_exec(gl_wrap_context* ctx, gl_wrap_batch_func* funcs, int numFuncs) {
	int i;
	for(i = 0; i < numFuncs; i++) {
		// Grab the function from the array
		gl_wrap_batch_func func = funcs[i];

		// Locate the handler function in the jump table at jump_index, execute
		// it using the context and function arguments.
		gl_wrap_jump_table[func.jump_index](ctx, func.args);
	}
}
`

func writeCCode(out io.Writer, tmplFunctions []TmplFunction, tmplTypes []TmplType) error {
	funcMap := map[string]interface{}{
		"NeedComma": func(index int, s []TmplCArgPair) bool {
			return index != len(s)-1
		},
	}
	tmpl := template.Must(template.New("cCodeTemplate").Funcs(funcMap).Parse(cCodeTemplate))
	err = tmpl.Execute(out, map[string]interface{}{
		"Types":     tmplTypes,
		"Functions": tmplFunctions,
	})
	if err != nil {
		return err
	}
	return nil
}
