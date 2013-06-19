// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// Code generated by this program is also under
// the above license.

package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const goHelperCode = `
// Package 'opengl' implements OpenGL version <VERSION>
package opengl

// #cgo LDFLAGS: -lopengl32
// #include "gl<VERSION_WITHOUT_DOTS>.h"
import "C"

import(
	"strconv"
	"strings"
	"unsafe"
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
			if major == wantedMajor && minor >= wantedMinor {
				return true
			}
		} else {
			if major == wantedMajor && minor == wantedMinor && rev >= wantedRev {
				return true
			}
		}
	}
	return false
}
`

func generateGo(packageDir, prefix, version, versionWithoutDots string, versionProcs, possibleProcs []*Procedure, constants map[string]string) {
	// Build list of all procedures (version required procedures + possible procedures)
	var allProcs []*Procedure

	allProcs = append(allProcs, versionProcs...)
	allProcs = append(allProcs, possibleProcs...)

	// Create Go file
	code, err := os.Create(filepath.Join(packageDir, "gl"+versionWithoutDots+".go"))
	if err != nil {
		log.Fatal(err)
	}
	defer code.Close()

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
	fmt.Fprintf(code, "\tcontext *C.%sContext\n", prefix)
	for _, p := range allProcs {
		var name, args, returns string

		if fn, ok := specialProcedures[p.Name]; ok {
			name, args, _, returns = fn("glc.context", prefix, p)
		} else {
			name, args, _, returns = autoProcedure("glc.context", prefix, p)
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
		// For debugging
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
		var name, args, body, returns string

		if fn, ok := specialProcedures[p.Name]; ok {
			name, args, body, returns = fn("glc.context", prefix, p)
		} else {
			name, args, body, returns = autoProcedure("glc.context", prefix, p)
		}
		if len(returns) > 0 {
			returns = " " + returns
		}

		fmt.Fprintf(code, "\tglc.%s = func(%s)%s {\n", name, args, returns)
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

	fmt.Fprintf(code, "\treturn glc\n")
	fmt.Fprintf(code, "}\n")
}