// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// Code generated by this program is also under
// the above license.

package generator

import (
	"strings"
)

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
	case "GLintptr":
		return "int32"

	case "GLvoid*":
		return "unsafe.Pointer"

	case "GLchar*":
		return "*byte"
	case "GLchar**":
		return "**byte"
	case "GLshort*":
		return "*int16"
	case "GLushort*":
		return "*uint16"
	case "GLint*":
		return "*int32"
	case "GLsizei*":
		return "*int32"
	case "GLuint*":
		return "*uint32"
	case "GLenum*":
		return "*uint32"
	case "GLboolean*":
		return "*bool"
	case "GLbyte*":
		return "*int8"
	case "GLubyte*":
		return "*uint8"
	case "GLfloat*":
		return "*float32"
	case "GLclampf*":
		return "*float32"
	case "GLdouble*":
		return "*float64"
	case "GLclampd*":
		return "*float64"
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