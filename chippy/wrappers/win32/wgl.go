// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build windows,!no_opengl

package win32

/*
#define UNICODE
#include <windows.h>
#include <GL/gl.h>

#cgo LDFLAGS: -lopengl32

typedef HGLRC (*chippy_p_wglCreateContextAttribsARB) (HDC, HGLRC, const int*);
typedef char* (*chippy_p_wglGetExtensionsStringARB) (HDC);
typedef BOOL (*chippy_p_wglSwapIntervalEXT) (int);

HGLRC chippy_wglCreateContextAttribsARB(void* p, HDC hDC, HGLRC hshareContext, const int* attribList) {
	chippy_p_wglCreateContextAttribsARB fn = (chippy_p_wglCreateContextAttribsARB)p;
	return fn(hDC, hshareContext, attribList);
}

char* chippy_wglGetExtensionsStringARB(void* p, HDC hdc) {
	chippy_p_wglGetExtensionsStringARB fn = (chippy_p_wglGetExtensionsStringARB)p;
	return fn(hdc);
}

BOOL chippy_wglSwapIntervalEXT(void* p, int interval) {
	chippy_p_wglSwapIntervalEXT fn = (chippy_p_wglSwapIntervalEXT)p;
	return fn(interval);
}

*/
import "C"

import (
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

func WglGetProcAddress(lpszProc string) uintptr {
	cstr := (*C.CHAR)(unsafe.Pointer(C.CString(lpszProc)))
	defer C.free(unsafe.Pointer(cstr))

	r := uintptr(unsafe.Pointer(C.wglGetProcAddress(cstr)))
	negativeOne := -1
	if r == 0 || r == 1 || r == 2 || r == 3 || r == uintptr(negativeOne) {
		return 0
	}
	return uintptr(r)
}

func WglMakeCurrent(hdc HDC, hglrc HGLRC) bool {
	return C.wglMakeCurrent(C.HDC(hdc), C.HGLRC(hglrc)) != 0
}

func WglShareLists(hglrc1, hglrc2 HGLRC) bool {
	return C.wglShareLists(C.HGLRC(hglrc1), C.HGLRC(hglrc2)) != 0
}

const (
	WGL_CONTEXT_MAJOR_VERSION_ARB = 0x2091
	WGL_CONTEXT_MINOR_VERSION_ARB = 0x2092
	WGL_CONTEXT_LAYER_PLANE_ARB   = 0x2093
	WGL_CONTEXT_FLAGS_ARB         = 0x2094
	WGL_CONTEXT_PROFILE_MASK_ARB  = 0x9126

	//Accepted as bits in the attribute value for WGL_CONTEXT_FLAGS in <*attribList>:

	WGL_CONTEXT_DEBUG_BIT_ARB              = 0x0001
	WGL_CONTEXT_FORWARD_COMPATIBLE_BIT_ARB = 0x0002

	// Accepted as bits in the attribute value for WGL_CONTEXT_PROFILE_MASK_ARB in <*attribList>:

	WGL_CONTEXT_CORE_PROFILE_BIT_ARB          = 0x00000001
	WGL_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB = 0x00000002

	// New errors returned by GetLastError:

	ERROR_INVALID_VERSION_ARB = 0x2095
	ERROR_INVALID_PROFILE_ARB = 0x2096
)

func WglCreateContextAttribsARB(hdc HDC, hglrc HGLRC, attribs []Int) (HGLRC, bool) {
	ptr := WglGetProcAddress("wglCreateContextAttribsARB")
	if ptr == 0 {
		return nil, false
	}
	ret := C.chippy_wglCreateContextAttribsARB(unsafe.Pointer(ptr), C.HDC(hdc), C.HGLRC(hglrc), (*C.int)(unsafe.Pointer(&attribs[0])))
	return HGLRC(ret), true
}

func WglGetExtensionsStringARB(hdc HDC) (string, bool) {
	ptr := WglGetProcAddress("wglGetExtensionsStringARB")
	if ptr == 0 {
		return "", false
	}
	ret := C.chippy_wglGetExtensionsStringARB(unsafe.Pointer(ptr), C.HDC(hdc))
	return C.GoString(ret), true
}

func WglSwapIntervalEXT(interval int) bool {
	ptr := WglGetProcAddress("wglSwapIntervalEXT")
	if ptr == 0 {
		return false
	}
	return C.chippy_wglSwapIntervalEXT(unsafe.Pointer(ptr), C.int(interval)) != 0
}

const (
	GL_VERSION = 0x1F02
)

func GlGetString(name uint32) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name)))))
}
