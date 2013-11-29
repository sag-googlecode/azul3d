// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build linux

// Do not use - Wrappers around very specific x11 APIs.
package x11

/*
#include <stdlib.h>
#include <xcb/xcb.h>
#include <xcb/render.h>

#cgo LDFLAGS: -lxcb -lxcb-render
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

type (
	RenderPictformat C.xcb_render_pictformat_t
	RenderPicture    C.xcb_render_picture_t
)

func (c *Connection) RenderCreatePictureChecked(pid RenderPicture, d Drawable, format RenderPictformat, mask uint32, list *uint32) VoidCookie {
	return VoidCookie(C.xcb_render_create_picture_checked(
		c.c(),
		C.xcb_render_picture_t(pid),
		C.xcb_drawable_t(d),
		C.xcb_render_pictformat_t(format),
		C.uint32_t(mask),
		(*C.uint32_t)(unsafe.Pointer(list)),
	))
}

func (c *Connection) RenderFreePicture(pid RenderPicture) VoidCookie {
	return VoidCookie(C.xcb_render_free_picture(
		c.c(),
		C.xcb_render_picture_t(pid),
	))
}

func (c *Connection) RenderCreateCursorChecked(cid Cursor, source RenderPicture, x, y uint16) VoidCookie {
	return VoidCookie(C.xcb_render_create_cursor_checked(
		c.c(),
		C.xcb_cursor_t(cid),
		C.xcb_render_picture_t(source),
		C.uint16_t(x),
		C.uint16_t(y),
	))
}

type ERenderQueryPictFormatsReply struct {
	ResponseType C.uint8_t
	Pad0         C.uint8_t
	Sequence     C.uint16_t
	Length       C.uint32_t
	NumFormats   C.uint32_t
	NumScreens   C.uint32_t
	NumDepths    C.uint32_t
	NumVisuals   C.uint32_t
	NumSubpixel  C.uint32_t
	Pad1         [4]C.uint8_t
}
type RenderQueryPictFormatsReply struct {
	*ERenderQueryPictFormatsReply
}

func (c *RenderQueryPictFormatsReply) c() *C.xcb_render_query_pict_formats_reply_t {
	ptr := c.ERenderQueryPictFormatsReply
	return (*C.xcb_render_query_pict_formats_reply_t)(unsafe.Pointer(ptr))
}

type RenderQueryPictFormatsCookie C.xcb_render_query_pict_formats_cookie_t

func (c RenderQueryPictFormatsCookie) c() C.xcb_render_query_pict_formats_cookie_t {
	return C.xcb_render_query_pict_formats_cookie_t(c)
}

func (c *Connection) RenderQueryPictFormats() RenderQueryPictFormatsCookie {
	cookie := C.xcb_render_query_pict_formats(c.c())
	return RenderQueryPictFormatsCookie(cookie)
}

func (c *Connection) RenderQueryPictFormatsReply(cookie RenderQueryPictFormatsCookie) (reply *RenderQueryPictFormatsReply, err error) {
	var e *C.xcb_generic_error_t
	cReply := C.xcb_render_query_pict_formats_reply(c.c(), cookie.c(), &e)
	if e == nil {
		reply = new(RenderQueryPictFormatsReply)
		reply.ERenderQueryPictFormatsReply = (*ERenderQueryPictFormatsReply)(unsafe.Pointer(cReply))
		runtime.SetFinalizer(reply, func(r *RenderQueryPictFormatsReply) {
			C.free(unsafe.Pointer(r.ERenderQueryPictFormatsReply))
		})
	}
	if e != nil {
		err = errors.New("RenderQueryPictFormatsReply(): " + xcbError(e))
		C.free(unsafe.Pointer(e))
	}
	return
}
