// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// This source file was automatically generated using glwrap.
// +build arm android

// Package opengl implements Go bindings to OpenGL ES
package opengl

/*
#include "gl.h"
*/
import "C"
import "unsafe"

type(
	Sync C.GLsync
	DEBUGPROC C.GLDEBUGPROC
	DEBUGPROCARB C.GLDEBUGPROCARB
	DEBUGPROCKHR C.GLDEBUGPROCKHR
	DEBUGPROCAMD C.GLDEBUGPROCAMD

)

// String converts from a OpenGL *uint8 string to a Go string.
func String(v *uint8) string {
	return C.GoString((*C.char)(unsafe.Pointer(v)))
}

// ByteString converts from a OpenGL *byte string to a Go string.
func ByteString(v *byte) string {
	return C.GoString((*C.char)(unsafe.Pointer(v)))
}

// Bool converts from a OpenGL uint8 boolean to a Go bool.
func Bool(v uint8) bool {
	if v == 1 {
		return true
	}
	return false
}

// GLBool converts from a Go bool to a OpenGL uint8 boolean.
func GLBool(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}