// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build linux

package x11

/*
#include <stdlib.h>
#include <X11/Xlib.h>

// So we get functions defined instead of just the macros
#define XUTIL_DEFINE_FUNCTIONS
#include <X11/Xutil.h>

#cgo LDFLAGS: -lX11

#ifndef X_HAVE_UTF8_STRING
    #error Version of X11/Xlib.h installed must have X_HAVE_UTF8_STRING defined, please update Xlib version!
#endif
*/
import "C"

import (
	"unsafe"
)

const (
	X_HAVE_UTF8_STRING = C.X_HAVE_UTF8_STRING
)

type (
	XWMHints   C.XWMHints
	XClassHint C.XClassHint
	XSizeHints C.XSizeHints
	XIconSize  C.XIconSize
)

func XSetWMHints(display *Display, w Window, wmhints *XWMHints) int {
	return int(C.XSetWMHints((*C.Display)(display), C.Window(w), (*C.XWMHints)(wmhints)))
}

func XGetWMHints(display *Display, w Window) *XWMHints {
	return (*XWMHints)(C.XGetWMHints((*C.Display)(display), C.Window(w)))
}

func XGetIconSizes(display *Display, w Window) (int, []*XIconSize) {
	var count C.int
	var list *C.XIconSize

	slice := []*XIconSize{}

	ret := int(C.XGetIconSizes((*C.Display)(display), C.Window(w), &list, &count))
	if ret == 0 {
		return ret, slice
	}
	defer C.XFree(unsafe.Pointer(list))

	for i := 0; i < int(count); i++ {
		size := (*C.XIconSize)(elementAtIndex(unsafe.Pointer(list), uintptr(i), unsafe.Sizeof(C.XIconSize{})))
		slice = append(slice, (*XIconSize)(size))
	}
	return ret, slice
}

func Xutf8SetWMProperties(display *Display, w Window, window_name, icon_name, argv string, normal_hints *XSizeHints, wm_hints *XWMHints, class_hints *XClassHint) {
	cwindow_name := C.CString(window_name)
	defer C.free(unsafe.Pointer(cwindow_name))

	cicon_name := C.CString(icon_name)
	defer C.free(unsafe.Pointer(cicon_name))

	cargv := C.CString(argv)
	defer C.free(unsafe.Pointer(cargv))

	C.Xutf8SetWMProperties((*C.Display)(display), C.Window(w), cwindow_name, cicon_name, &cargv, C.int(len(argv)), (*C.XSizeHints)(normal_hints), (*C.XWMHints)(wm_hints), (*C.XClassHint)(class_hints))
}
