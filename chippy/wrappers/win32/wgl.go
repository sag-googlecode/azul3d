// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

package win32

/*
#define UNICODE
#include <windows.h>

#cgo LDFLAGS: -lopengl32
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// OpenGL functions
type HGLRC C.HGLRC

func WglCreateContext(hdc HDC) HGLRC {
	return HGLRC(C.wglCreateContext(C.HDC(hdc)))
}

func WglDeleteContext(hglrc HGLRC) bool {
	return C.wglDeleteContext(C.HGLRC(hglrc)) != 0
}

var opengl32 C.HMODULE

func init() {
	cstr := StringToLPTSTR("opengl32.dll")
	defer C.free(unsafe.Pointer(cstr))
	opengl32 = C.LoadLibrary(cstr)
	if opengl32 == nil {
		panic(fmt.Sprintf("Unable to load opengl32.dll; LoadLibrary():", GetLastErrorString()))
	}
}

func OpenGL32GetProcAddress(proc string) uintptr {
	cstr := (*C.CHAR)(unsafe.Pointer(C.CString(proc)))
	defer C.free(unsafe.Pointer(cstr))
	return uintptr(unsafe.Pointer(C.GetProcAddress(opengl32, cstr)))
}

func WglGetProcAddress(lpszProc string) uintptr {
	cstr := (*C.CHAR)(unsafe.Pointer(C.CString(lpszProc)))
	defer C.free(unsafe.Pointer(cstr))
	return uintptr(unsafe.Pointer(C.wglGetProcAddress(cstr)))
}

func WglMakeCurrent(hdc HDC, hglrc HGLRC) bool {
	return C.wglMakeCurrent(C.HDC(hdc), C.HGLRC(hglrc)) != 0
}
