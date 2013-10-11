// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// Code generated by this program is also under
// the above license.

package generator

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const goNormalHeader = `//
// +build !opengl_debug

// Package 'opengl' implements OpenGL version <VERSION>
package opengl
`

const goDebugHeader = `//
// +build opengl_debug

// Package 'opengl' implements OpenGL version <VERSION>
package opengl
`

const goHelperCode = `
// #cgo windows LDFLAGS: -lopengl32
// #cgo linux LDFLAGS: -lGL -ldl
// #include "gl<VERSION_WITHOUT_DOTS>.h"
import "C"

import(
	"strconv"
	"strings"
	"unsafe"
	"sync"
	"fmt"
)

func boolToGL(b bool) C.GLboolean {
	if b {
		return C.GLboolean(1)
	}
	return C.GLboolean(0)
}

func parseVersions(s string) (n, major, minor, rev int) {
	var err error

	versions := strings.Split(s, ".")
	if len(versions) > 2 {
		versions = versions[0:3]
	}

	if len(versions) > 0 {
		major, err = strconv.Atoi(versions[0])
		if err != nil {
			return 0, 0, 0, 0
		}
	}

	if len(versions) > 1 {
		minor, err = strconv.Atoi(versions[1])
		if err != nil {
			return 0, 0, 0, 0
		}
	}
	n = len(versions)

	if len(versions) > 2 {
		_, err = fmt.Sscanf(versions[2] + "\n", "%d", &rev)
		if err != nil {
			n = 2
		}
	}

	return
}

func versionSupported(glc *Context) bool {
	ver := glc.GetString(VERSION)
	if len(ver) > 0 {
		n, wantedMajor, wantedMinor, wantedRev := parseVersions("<VERSION>")
		if n < 2 {
			fmt.Printf("OpenGL: *** JSON version parsing failed for %q ***\n", "<VERSION>")
			return false
		}

		n, major, minor, rev := parseVersions(ver)
		if n < 2 {
			fmt.Printf("OpenGL: *** Driver reported version parsing failed for %q ***\n", ver)
			return false
		}


		if major > wantedMajor {
			return true
		}
		if n == 2 {
			fmt.Printf("OpenGL: *** Driver reported version has no revision! %q ***\n", ver)
			if major >= wantedMajor && minor >= wantedMinor {
				return true
			}
		} else {
			if major >= wantedMajor && minor >= wantedMinor && rev >= wantedRev {
				return true
			}
		}
	}
	return false
}

func (glc *Context) queryExtensions() {
	// Initialize extensions map
	glc.extensions = make(map[string]bool)

	// Query extensions string
	extString := glc.GetString(EXTENSIONS)

	if len(extString) > 0 {
		for _, ext := range strings.Split(extString, " ") {
			if len(ext) > 0 {
				glc.extensions[ext] = true
			}
		}
	}
}

func (glc *Context) Panic(err string) {
	glc.access.Lock()
	defer glc.access.Unlock()

	fmt.Println("OpenGL call stack (last 500 - most recent first).")

	// Print stack now
	count := 0
	for i := len(glc.traceback); i > 0; i-- {
		count++
		fmt.Printf("%3.d. %s\n", count, glc.traceback[i-1])
	}

	panic(err)
}

func (glc *Context) trace(name string) {
	glc.access.Lock()

	glc.traceback = append(glc.traceback, name)
	l := len(glc.traceback)
	if l > 500 {
		glc.traceback = glc.traceback[l-500:l]
	}

	if glc.inBeginEnd {
		glc.access.Unlock()
		return
	}
	err := glc.GetError()
	if err != NO_ERROR {
		glc.access.Unlock()

		switch err {
		case INVALID_ENUM:
			glc.Panic("GL_INVALID_ENUM: An unacceptable value was specified for an enumerated argument.")
		case INVALID_VALUE:
			glc.Panic("GL_INVALID_VALUE: A numeric argument is out of range.")
		case INVALID_OPERATION:
			glc.Panic("GL_INVALID_OPERATION: The specified operation is not allowed in the current state.")
		case INVALID_FRAMEBUFFER_OPERATION:
			glc.Panic("GL_INVALID_FRAMEBUFFER_OPERATION: The framebuffer object is not complete.")
		case OUT_OF_MEMORY:
			glc.Panic("GL_OUT_OF_MEMORY: There is not enough memory left to execute the command.")
		case STACK_UNDERFLOW:
			glc.Panic("GL_STACK_UNDERFLOW: An attempt has been made to perform an operation that would cause an internal stack to underflow.")
		case STACK_OVERFLOW:
			glc.Panic("GL_STACK_OVERFLOW: An attempt has been made to perform an operation that would cause an internal stack to overflow.")
		}
	} else {
		glc.access.Unlock()
	}
}

// Extension tells if the specified extension is supported by the OpenGL
// context.
//
// Extensions are stored internally as an map for performance, so lookups are
// very quick and require no OpenGL calls.
//
// Like other OpenGL functions, this is not thread safe.
//
// If this function always returns false, ensure that New() has been called in
// an active OpenGL context (as the extensions are queried then).
func (glc *Context) Extension(ext string) bool {
	_, ok := glc.extensions[ext]
	return ok
}

`

func generateGo(packageDir, prefix, version, versionWithoutDots string, versionProcs, possibleProcs []*Procedure, constants map[string]string, trace bool) {
	// Build list of all procedures (version required procedures + possible procedures)
	var allProcs []*Procedure

	allProcs = append(allProcs, versionProcs...)
	allProcs = append(allProcs, possibleProcs...)

	// Create Go file
	var name string
	if trace {
		name = "gl" + versionWithoutDots + "debug.go"
	} else {
		name = "gl" + versionWithoutDots + ".go"
	}
	code, err := os.Create(filepath.Join(packageDir, name))
	if err != nil {
		log.Fatal(err)
	}
	defer code.Close()

	// Write license to file
	code.Write([]byte(licenseHeader))

	// Write trace / !trace header
	if trace {
		debugHeader := strings.Replace(goDebugHeader, "<VERSION>", version, -1)
		code.Write([]byte(debugHeader))
	} else {
		normalHeader := strings.Replace(goNormalHeader, "<VERSION>", version, -1)
		code.Write([]byte(normalHeader))
	}

	// Write helper code to file
	helperCode := strings.Replace(goHelperCode, "<VERSION>", version, -1)
	helperCode = strings.Replace(helperCode, "<VERSION_WITHOUT_DOTS>", versionWithoutDots, -1)
	code.Write([]byte(helperCode))

	fmt.Fprintf(code, "const(\n")
	for constName, constValue := range constants {
		fmt.Fprintf(code, "\t%s = %s\n", cToGoConstName(constName), constValue)
	}
	fmt.Fprintf(code, ")\n")
	fmt.Fprintf(code, "\n")

	// Write out the Context type
	fmt.Fprintf(code, "type Context struct {\n")
	fmt.Fprintf(code, "\taccess sync.Mutex\n")
	fmt.Fprintf(code, "\tcontext *C.%sContext\n", prefix)
	fmt.Fprintf(code, "\textensions map[string]bool\n")
	fmt.Fprintf(code, "\tinBeginEnd bool\n")
	fmt.Fprintf(code, "\ttraceback []string\n")
	for _, p := range allProcs {
		var name, args, returns string

		if fn, ok := specialProcedures[p.Name]; ok {
			name, args, _, _, returns = fn("glc.context", prefix, p)
		} else {
			name, args, _, _, returns = autoProcedure("glc.context", prefix, p)
		}
		if len(returns) > 0 {
			returns = " " + returns
		}

		fmt.Fprintf(code, "\t%s func(%s)%s\n", name, args, returns)
	}
	fmt.Fprintf(code, "}\n\n")

	fmt.Fprintf(code, "func New() *Context {\n")
	fmt.Fprintf(code, "\tglc := new(Context)\n")
	fmt.Fprintf(code, "\tglc.context = C.%sNewContext()\n", prefix)
	fmt.Fprintf(code, "\n")

	/*
		// Can be used for debugging to assert that all function pointers where found properly.
		for _, p := range versionProcs {
			if !p.Extension {
				glStripped := strings.TrimLeft(p.Name, "gl")
				fmt.Fprintf(code, "\tif glc.context.fn%s == nil { ", glStripped)

				fmt.Fprintf(code, "panic(\"%s missing\")", p.Name)

				fmt.Fprintf(code, " }\n")
			}
		}
		fmt.Fprintf(code, "\n")
	*/

	// Wrap each function with an closure tied to the object
	for _, p := range allProcs {
		var name, untypedArgs, args, body, returns string

		if fn, ok := specialProcedures[p.Name]; ok {
			name, args, untypedArgs, body, returns = fn("glc.context", prefix, p)
		} else {
			name, args, untypedArgs, body, returns = autoProcedure("glc.context", prefix, p)
		}
		if len(returns) > 0 {
			returns = " " + returns
		}

		fmt.Fprintf(code, "\tglc.%s = func(%s)%s {\n", name, args, returns)
		if trace && name != "GetError" {

			// Make a proper "%v,%v,%v,%v" formatting string
			percentVeesBuf := new(bytes.Buffer)
			splitUntypedArgs := strings.Split(untypedArgs, ",")
			for i, arg := range splitUntypedArgs {
				if len(arg) > 0 {
					if i != len(splitUntypedArgs)-1 {
						fmt.Fprintf(percentVeesBuf, "%%v, ")
					} else {
						fmt.Fprintf(percentVeesBuf, "%%v")
					}
				}
			}
			percentVees := percentVeesBuf.String()

			if percentVeesBuf.Len() > 0 {
				fmt.Fprintf(code, "\t\tfmtCall := fmt.Sprintf(\"%s(%s)\", %s)\n", name, percentVees, untypedArgs)
			} else {
				fmt.Fprintf(code, "\t\tfmtCall := \"%s()\"\n", name)
			}
			fmt.Fprintf(code, "\t\tdefer glc.trace(fmtCall)\n")
		}
		for _, line := range strings.Split(body, "\n") {
			if len(line) > 0 {
				fmt.Fprintf(code, "\t%s\n", line)
			}
		}
		fmt.Fprintf(code, "\t}\n\n")
	}

	fmt.Fprintf(code, "\tif !versionSupported(glc) {\n")
	fmt.Fprintf(code, "\t\treturn nil\n")
	fmt.Fprintf(code, "\t}\n")

	fmt.Fprintf(code, "\tglc.queryExtensions()\n")

	fmt.Fprintf(code, "\treturn glc\n")
	fmt.Fprintf(code, "}\n")
}
