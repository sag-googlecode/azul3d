package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	specFilePath      string
	outputDir         string
	specialProcedures map[string]func(ctx, prefix string, p *Procedure) (name, args, body, returns string)
)

func init() {
	log.SetFlags(0)
	flag.StringVar(&specFilePath, "spec", "", "JSON specification file")
	flag.StringVar(&outputDir, "o", "", "output directory")

	specialProcedures = make(map[string]func(ctx, prefix string, p *Procedure) (name, args, body, returns string))

	///////////////////////////////////////////////////////////////////////////////////////////////////
	specialProcedures["glAreTexturesResident"] = func(ctx, prefix string, p *Procedure) (name, args, body, returns string) {
		body = fmt.Sprintf(`
	var cRes *C.GLboolean
	status = C.%sAreTexturesResident(%s, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0

	residencies = make([]bool, len(textures))
	for i := 0; i < len(textures); i++ {
		residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
	}
	return
`, prefix, ctx)

		return "AreTexturesResident", "textures []uint32", body, "(status bool, residencies []bool)"
	}
	///////////////////////////////////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////////////////////////////////
	specialProcedures["glGetString"] = func(ctx, prefix string, p *Procedure) (name, args, body, returns string) {
		body = fmt.Sprintf(`
	cstr := C.%sGetString(%s, C.GLenum(name))
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
`, prefix, ctx)

		return "GetString", "name uint32", body, "string"
	}
	///////////////////////////////////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////////////////////////////////
	specialProcedures["glBindAttribLocation"] = func(ctx, prefix string, p *Procedure) (name, args, body, returns string) {
		body = fmt.Sprintf(`
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(&cstr))

	C.%sBindAttribLocation(%s, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
	return
`, prefix, ctx)

		return "BindAttribLocation", "program, index uint32, name string", body, ""
	}
	///////////////////////////////////////////////////////////////////////////////////////////////////
}

type Procedure struct {
	Constants map[string]string
	Name      string
	Returns   string
	Takes     []string
	Versions  []string
	Extension bool
	Reference string
}

func cToGoName(n string) string {
	switch n {
	case "func":
		return "Func"
	case "type":
		return "Type"
	case "range":
		return "Range"
	case "map":
		return "Map"
	}
	return n
}

func cToGoConstName(n string) string {
	for _, number := range "0123456789" {
		if strings.HasPrefix(n, string(number)) {
			return "GL_" + n
		}
	}
	return n
}

func cToGoType(c string) string {
	switch c {
	case "GLboolean":
		return "bool"
	case "GLbyte":
		return "int8"
	case "GLubyte":
		return "uint8"
	case "GLshort":
		return "int16"
	case "GLushort":
		return "uint16"
	case "GLint":
		return "int32"
	case "GLuint":
		return "uint32"
	case "GLsizei":
		return "int32"
	case "GLenum":
		return "uint32"
	case "GLbitfield":
		return "uint32"
	case "GLfloat":
		return "float32"
	case "GLclampf":
		return "float32"
	case "GLclampd":
		return "float64"
	case "GLdouble":
		return "float64"
	case "GLsizeiptr":
		return "int32"

	case "GLvoid*":
		return "unsafe.Pointer"
	case "GLshort*":
		return "[]int16"
	case "GLushort*":
		return "[]uint16"
	case "GLint*":
		return "[]int32"
	case "GLuint*":
		return "[]uint32"
	case "GLboolean*":
		return "[]bool"
	case "GLbyte*":
		return "[]int8"
	case "GLubyte*":
		return "[]uint8"
	case "GLfloat*":
		return "[]float32"
	case "GLclampf*":
		return "[]float32"
	case "GLdouble*":
		return "[]float64"
	case "GLclampd*":
		return "[]float64"
	}
	return c
}

func cToGoConversion(c string) (pre, post string) {
	if c == "GLboolean" {
		post = " != 0"
		return
	}

	if strings.HasSuffix(c, "*") {
		pre = cToGoType(c) + "("
		post = ")"
		return
	}

	pre = cToGoType(c) + "("
	post = ")"
	return
}

func autoProcedure(ctx, prefix string, p *Procedure) (name, args, body, returns string) {
	glStripped := strings.TrimLeft(p.Name, "gl")

	name = glStripped

	argsBuf := new(bytes.Buffer)
	for i, arg := range p.Takes {
		if arg != "void" {
			split := strings.Split(arg, " ")

			if i != len(p.Takes)-1 {
				var futureSplit = strings.Split(p.Takes[i+1], " ")

				if split[0] == futureSplit[0] {
					fmt.Fprintf(argsBuf, "%s", cToGoName(split[1]))
				} else {
					fmt.Fprintf(argsBuf, "%s %s", cToGoName(split[1]), cToGoType(split[0]))
				}

				fmt.Fprintf(argsBuf, ", ")
			} else {
				fmt.Fprintf(argsBuf, "%s %s", cToGoName(split[1]), cToGoType(split[0]))
			}
		}
	}
	args = string(argsBuf.Bytes())

	if p.Returns != "void" {
		returns = cToGoType(p.Returns)
	}

	bodyBuf := new(bytes.Buffer)

	procArgNames := new(bytes.Buffer)
	fmt.Fprintf(procArgNames, "glc.context, ")

	for i, arg := range p.Takes {
		if arg != "void" {

			split := strings.Split(arg, " ")
			argType := split[0]
			argName := cToGoName(split[1])

			if strings.HasPrefix(cToGoType(argType), "[]") {
				// Do unsafe conversion

				// Move * to other side
				argType = strings.TrimRight(argType, "*")
				fmt.Fprintf(procArgNames, "(*C.%s)(unsafe.Pointer(&%s[0]))", argType, argName)
			} else if argType == "GLvoid*" {
				fmt.Fprintf(procArgNames, "%s", argName)
			} else if argType == "GLboolean" {
				fmt.Fprintf(procArgNames, "boolToGL(%s)", argName)
			} else {
				fmt.Fprintf(procArgNames, "C.%s(%s)", argType, argName)
			}

			if i != len(p.Takes)-1 {
				fmt.Fprintf(procArgNames, ", ")
			}
		}
	}

	fmt.Fprintf(bodyBuf, "    ")

	var preConvert, postConvert string
	if p.Returns != "void" {
		preConvert, postConvert = cToGoConversion(p.Returns)
		fmt.Fprintf(bodyBuf, "return %s", preConvert)
	}

	fmt.Fprintf(bodyBuf, "C.%s%s(%s)", prefix, glStripped, procArgNames.Bytes())

	if p.Returns != "void" {
		fmt.Fprintf(bodyBuf, postConvert)
	}
	fmt.Fprintf(bodyBuf, "\n")

	body = string(bodyBuf.Bytes())
	return
}

func printSyntaxError(js string, err error) {
	syntax, ok := err.(*json.SyntaxError)
	if !ok {
		fmt.Println(err)
		return
	}

	start, end := strings.LastIndex(js[:syntax.Offset], "\n")+1, len(js)
	if idx := strings.Index(js[start:], "\n"); idx >= 0 {
		end = start + idx
	}

	line, pos := strings.Count(js[:start], "\n"), int(syntax.Offset)-start-1

	fmt.Printf("Error in line %d: %s \n", line, err)
	fmt.Printf("%s\n%s^", js[start:end], strings.Repeat(" ", pos))
}

func main() {
	var (
		file       *os.File
		err        error
		procedures []*Procedure
	)

	flag.Parse()

	if len(specFilePath) == 0 || len(outputDir) == 0 {
		flag.Usage()
	}

	file, err = os.Open(specFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &procedures)
	if err != nil {
		printSyntaxError(string(data), err)
		return
	}

	err = os.MkdirAll(outputDir, os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}

	// Make list of all OpenGL versions determined from spec
	versions := make(map[string]bool)
	for _, p := range procedures {
		for _, v := range p.Versions {
			versions[v] = true
		}
	}

	for version, _ := range versions {
		var (
			versionProcs []*Procedure
			constants    map[string]string
		)

		for _, p := range procedures {
			if len(p.Constants) > 0 {
				constants = p.Constants
			}
			for _, v := range p.Versions {
				if v == version {
					versionProcs = append(versionProcs, p)
					break
				}
			}
		}

		versionWithoutDots := strings.Replace(version, ".", "", -1)
		packageDir := filepath.Join(outputDir, version)

		err = os.MkdirAll(packageDir, os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		prefix := "gl" + versionWithoutDots

		///////////////////////
		// H code generation //
		///////////////////////
		header, err := os.Create(filepath.Join(packageDir, "gl"+versionWithoutDots+".h"))
		if err != nil {
			log.Fatal(err)
		}
		defer header.Close()

		fmt.Fprintf(header, "#include <stdbool.h>\n")
		fmt.Fprintf(header, "#include <stdlib.h>\n")
		fmt.Fprintf(header, `
typedef char GLchar;
typedef unsigned int GLenum;
typedef unsigned char GLboolean;
typedef unsigned int GLbitfield;
typedef signed char GLbyte;
typedef short GLshort;
typedef int GLint;
typedef int GLsizei;
typedef unsigned char GLubyte;
typedef unsigned short GLushort;
typedef unsigned int GLuint;
typedef float GLfloat;
typedef float GLclampf;
typedef double GLdouble;
typedef double GLclampd;
typedef void GLvoid;
typedef ptrdiff_t GLsizeiptr;
`)
		fmt.Fprintf(header, "\n")

		for _, p := range versionProcs {
			procArgs := new(bytes.Buffer)
			for i, arg := range p.Takes {
				if arg != "void" {
					split := strings.Split(arg, " ")
					argType := split[0]

					fmt.Fprintf(procArgs, argType)

					if i != len(p.Takes)-1 {
						fmt.Fprintf(procArgs, ", ")
					}
				}
			}

			glStripped := strings.TrimLeft(p.Name, "gl")
			fmt.Fprintf(header, "typedef %s(*%sP%s)(%s);\n", p.Returns, prefix, glStripped, procArgs.Bytes())
		}
		fmt.Fprintf(header, "\n")

		fmt.Fprintf(header, "typedef struct {\n")
		for _, p := range versionProcs {
			glStripped := strings.TrimLeft(p.Name, "gl")
			fmt.Fprintf(header, "    %sP%s fn%s;\n", prefix, glStripped, glStripped)
		}
		fmt.Fprintf(header, "} %sContext;\n\n", prefix)

		fmt.Fprintf(header, "extern %sContext* %sNewContext();\n", prefix, prefix)

		// Invokers
		for _, p := range versionProcs {
			glStripped := strings.TrimLeft(p.Name, "gl")

			fmt.Fprintf(header, "%s %s%s(", p.Returns, prefix, glStripped)
			fmt.Fprintf(header, "%sContext* glc", prefix)
			for i, arg := range p.Takes {
				if arg != "void" {
					if i == 0 {
						fmt.Fprintf(header, ", ")
					}

					fmt.Fprintf(header, arg)

					if i != len(p.Takes)-1 {
						fmt.Fprintf(header, ", ")
					}
				}
			}
			fmt.Fprintf(header, ");\n")
		}

		///////////////////////
		// C code generation //
		///////////////////////
		api, err := os.Create(filepath.Join(packageDir, "gl"+versionWithoutDots+".c"))
		if err != nil {
			log.Fatal(err)
		}
		defer api.Close()

		fmt.Fprintf(api, "#include <stdbool.h>\n")
		fmt.Fprintf(api, "#include <stdlib.h>\n")
		fmt.Fprintf(api, "#ifdef _WIN32\n")
		fmt.Fprintf(api, "    #include <windows.h>\n")
		fmt.Fprintf(api, "#endif\n")
		fmt.Fprintf(api, "#include \"gl%s.h\"\n", versionWithoutDots)

		//fmt.Fprintf(api, "#include \"_cgo_export.h\"\n")
		fmt.Fprintf(api, "\n")

		fmt.Fprintf(api, "#ifdef _WIN32\n")
		fmt.Fprintf(api, "    HMODULE gl%sOpenGL32;\n", versionWithoutDots)
		fmt.Fprintf(api, "    void* doGetProcAddress(char* name) {\n")
		fmt.Fprintf(api, "        if(gl%sOpenGL32 == NULL) {\n", versionWithoutDots)
		fmt.Fprintf(api, "            gl%sOpenGL32 = LoadLibrary(TEXT(\"opengl32.dll\"));\n", versionWithoutDots)
		fmt.Fprintf(api, "        }\n")
		fmt.Fprintf(api, "        return GetProcAddress(gl%sOpenGL32, TEXT(name));\n", versionWithoutDots)
		fmt.Fprintf(api, "    }\n")
		fmt.Fprintf(api, "#endif\n")
		fmt.Fprintf(api, "\n")

		/*
		   HMODULE WINAPI LoadLibrary(
		     _In_  LPCTSTR lpFileName
		   );
		   FARPROC WINAPI GetProcAddress(
		     _In_  HMODULE hModule,
		     _In_  LPCSTR lpProcName
		   );
		*/

		for _, p := range versionProcs {
			glStripped := strings.TrimLeft(p.Name, "gl")

			procArgs := new(bytes.Buffer)
			procArgNames := new(bytes.Buffer)
			for i, arg := range p.Takes {
				if arg != "void" {
					split := strings.Split(arg, " ")
					argType := split[0]
					argName := split[1]

					fmt.Fprintf(procArgs, argType)
					fmt.Fprintf(procArgNames, argName)

					if i != len(p.Takes)-1 {
						fmt.Fprintf(procArgs, ", ")
						fmt.Fprintf(procArgNames, ", ")
					}
				}
			}

			// The invoker, which simply calls the proc *assuming* it exists
			fmt.Fprintf(api, "%s %s%s(", p.Returns, prefix, glStripped)
			fmt.Fprintf(api, "%sContext* glc", prefix)
			for i, arg := range p.Takes {
				if arg != "void" {
					if i == 0 {
						fmt.Fprintf(api, ", ")
					}

					fmt.Fprintf(api, arg)

					if i != len(p.Takes)-1 {
						fmt.Fprintf(api, ", ")
					}
				}
			}
			fmt.Fprintf(api, ") {\n")
			fmt.Fprintf(api, "    return glc->fn%s(%s);\n", glStripped, procArgNames.Bytes())
			fmt.Fprintf(api, "}\n\n")
		}

		fmt.Fprintf(api, "%sContext* %sNewContext() {\n", prefix, prefix)
		fmt.Fprintf(api, "    %sContext* glc = calloc(1, sizeof(%sContext));\n", prefix, prefix)
		fmt.Fprintf(api, "\n")
		fmt.Fprintf(api, "    // Preload all procedures\n")
		for _, p := range versionProcs {
			glStripped := strings.TrimLeft(p.Name, "gl")

			fmt.Fprintf(api, "    glc->fn%s = (%sP%s)", glStripped, prefix, glStripped)
			if !p.Extension {
				fmt.Fprintf(api, "doGetProcAddress(\"%s\");\n", p.Name)
			} else {
				fmt.Fprintf(api, "wglGetProcAddress(\"%s\");\n", p.Name)
			}
		}
		fmt.Fprintf(api, "    return glc;\n")
		fmt.Fprintf(api, "}\n\n")

		////////////////////////
		// Go code generation //
		////////////////////////
		code, err := os.Create(filepath.Join(packageDir, "gl"+versionWithoutDots+".go"))
		if err != nil {
			log.Fatal(err)
		}
		defer code.Close()

		fmt.Fprintf(code, `package opengl

// #cgo LDFLAGS: -lopengl32
// #include "%s"
import "C"

import "unsafe"

func boolToGL(b bool) C.GLboolean {
	if b {
		return C.GLboolean(1)
	}
	return C.GLboolean(0)
}

`, "gl"+versionWithoutDots+".h")

		fmt.Fprintf(code, "const(\n")
		for constName, constValue := range constants {
			fmt.Fprintf(code, "    %s = %s\n", cToGoConstName(constName), constValue)
		}
		fmt.Fprintf(code, ")\n")
		fmt.Fprintf(code, "\n")

		fmt.Fprintf(code, "type Context struct {\n")
		fmt.Fprintf(code, "    context *C.%sContext\n", prefix)
		for _, p := range versionProcs {
			var name, args, returns string

			if fn, ok := specialProcedures[p.Name]; ok {
				name, args, _, returns = fn("glc.context", prefix, p)
			} else {
				name, args, _, returns = autoProcedure("glc.context", prefix, p)
			}
			if len(returns) > 0 {
				returns = " " + returns
			}

			fmt.Fprintf(code, "    %s func(%s)%s\n", name, args, returns)
		}
		fmt.Fprintf(code, "}\n\n")

		fmt.Fprintf(code, "func New() *Context {\n")
		fmt.Fprintf(code, "    glc := new(Context)\n")
		fmt.Fprintf(code, "    glc.context = C.%sNewContext()\n", prefix)
		fmt.Fprintf(code, "\n")

		// Verify each mandatory one exists
		for _, p := range versionProcs {
			if !p.Extension {
				glStripped := strings.TrimLeft(p.Name, "gl")
				fmt.Fprintf(code, "    if glc.context.fn%s == nil { ", glStripped)

				// For debugging
				//fmt.Fprintf(code, "panic(\"%s\")", p.Name)
				fmt.Fprintf(code, "return nil")

				fmt.Fprintf(code, " }\n")
			}
		}
		fmt.Fprintf(code, "\n")

		// Wrap each function with an closure tied to the object
		for _, p := range versionProcs {
			var name, args, body, returns string

			if fn, ok := specialProcedures[p.Name]; ok {
				name, args, body, returns = fn("glc.context", prefix, p)
			} else {
				name, args, body, returns = autoProcedure("glc.context", prefix, p)
			}
			if len(returns) > 0 {
				returns = " " + returns
			}

			fmt.Fprintf(code, "    glc.%s = func(%s)%s {\n", name, args, returns)
			for _, line := range strings.Split(body, "\n") {
				if len(line) > 0 {
					fmt.Fprintf(code, "    %s\n", line)
				}
			}
			fmt.Fprintf(code, "    }\n\n")
		}
		fmt.Fprintf(code, "    return glc\n")
		fmt.Fprintf(code, "}\n\n")
	}
}
