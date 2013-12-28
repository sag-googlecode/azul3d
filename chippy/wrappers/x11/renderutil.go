// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build linux

// Do not use - Wrappers around very specific x11 APIs.
package x11

/*
#include <stdlib.h>
#include <xcb/xcb.h>
#include <xcb/xcb_renderutil.h>

#cgo LDFLAGS: -lxcb -lxcb-render-util
*/
import "C"

import (
	"unsafe"
)

type PictStandard C.xcb_pict_standard_t

const (
	PICT_STANDARD_ARGB_32 PictStandard = C.XCB_PICT_STANDARD_ARGB_32
	PICT_STANDARD_RGB_24               = C.XCB_PICT_STANDARD_RGB_24
)

type RenderPictforminfo struct {
	Id       RenderPictformat
	Type     C.uint8_t
	Depth    C.uint8_t
	Pad0     [2]C.uint8_t
	Direct   C.xcb_render_directformat_t
	Colormap Colormap
}

func RenderUtilFindStandardFormat(formats *RenderQueryPictFormatsReply, format PictStandard) *RenderPictforminfo {
	ptr := C.xcb_render_util_find_standard_format(
		formats.c(),
		C.xcb_pict_standard_t(format),
	)
	return (*RenderPictforminfo)(unsafe.Pointer(ptr))
}