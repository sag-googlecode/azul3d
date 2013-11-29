// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build linux

// Do not use - Wrappers around very specific x11 APIs.
package x11

// Yes, I know it's ugly to use both Xlib and XCB; but suprisingly it is not
// really possible to just use XCB for what we are doing. For one thing GLX
// does not work with XCB (without mixing in Xlib, of course). For another,
// XCB's keyboard support is rudimentary at best, e.g. Only Xlib: XLookupKeysym
// supports UCS encoding, etc..
//
// It looks like this is *the way to do it*, and it looks like other
// open source projects do it this way too; so, there is that as well.

/*
#include <stdlib.h>
#include <X11/Xlib-xcb.h>
#include <X11/XKBlib.h>
#include <X11/Xutil.h>

#define XK_3270
#include <X11/keysym.h>
#include <X11/Xlibint.h>

#cgo LDFLAGS: -lX11-xcb -lX11 -lxcb

int chippy_xlib_error(Display* d, XErrorEvent* e);
XIC chippy_CreateIC(XIM xim, Display* d, Window w);
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

const (
	NoSymbol            = C.NoSymbol
	XK_minus            = C.XK_minus
	XK_BackSpace        = C.XK_BackSpace
	XK_Tab              = C.XK_Tab
	XK_Linefeed         = C.XK_Linefeed
	XK_Clear            = C.XK_Clear
	XK_Return           = C.XK_Return
	XK_Pause            = C.XK_Pause
	XK_Scroll_Lock      = C.XK_Scroll_Lock
	XK_Num_Lock         = C.XK_Num_Lock
	XK_Shift_Lock       = C.XK_Shift_Lock
	XK_Sys_Req          = C.XK_Sys_Req
	XK_Escape           = C.XK_Escape
	XK_Delete           = C.XK_Delete
	XK_underscore       = C.XK_underscore
	XK_asciitilde       = C.XK_asciitilde
	XK_grave            = C.XK_grave
	XK_plus             = C.XK_plus
	XK_equal            = C.XK_equal
	XK_colon            = C.XK_colon
	XK_semicolon        = C.XK_semicolon
	XK_quotedbl         = C.XK_quotedbl
	XK_apostrophe       = C.XK_apostrophe
	XK_comma            = C.XK_comma
	XK_less             = C.XK_less
	XK_period           = C.XK_period
	XK_greater          = C.XK_greater
	XK_slash            = C.XK_slash
	XK_question         = C.XK_question
	XK_backslash        = C.XK_backslash
	XK_bar              = C.XK_bar
	XK_space            = C.XK_space
	XK_parenright       = C.XK_parenright
	XK_exclam           = C.XK_exclam
	XK_at               = C.XK_at
	XK_numbersign       = C.XK_numbersign
	XK_dollar           = C.XK_dollar
	XK_percent          = C.XK_percent
	XK_asciicircum      = C.XK_asciicircum
	XK_ampersand        = C.XK_ampersand
	XK_asterisk         = C.XK_asterisk
	XK_parenleft        = C.XK_parenleft
	XK_braceleft        = C.XK_braceleft
	XK_bracketleft      = C.XK_bracketleft
	XK_braceright       = C.XK_braceright
	XK_bracketright     = C.XK_bracketright
	XK_Print            = C.XK_Print
	XK_Insert           = C.XK_Insert
	XK_ISO_Left_Tab     = C.XK_ISO_Left_Tab
	XK_3270_PrintScreen = C.XK_3270_PrintScreen

	XK_Select            = C.XK_Select
	XK_Execute           = C.XK_Execute
	XK_Help              = C.XK_Help
	XK_3270_Play         = C.XK_3270_Play
	XK_3270_ExSelect     = C.XK_3270_ExSelect
	XK_3270_CursorSelect = C.XK_3270_CursorSelect
	XK_Kanji             = C.XK_Kanji
	XK_3270_Attn         = C.XK_3270_Attn
	XK_3270_EraseEOF     = C.XK_3270_EraseEOF

	XK_0 = C.XK_0
	XK_1 = C.XK_1
	XK_2 = C.XK_2
	XK_3 = C.XK_3
	XK_4 = C.XK_4
	XK_5 = C.XK_5
	XK_6 = C.XK_6
	XK_7 = C.XK_7
	XK_8 = C.XK_8
	XK_9 = C.XK_9

	XK_a = C.XK_a
	XK_b = C.XK_b
	XK_c = C.XK_c
	XK_d = C.XK_d
	XK_e = C.XK_e
	XK_f = C.XK_f
	XK_g = C.XK_g
	XK_h = C.XK_h
	XK_i = C.XK_i
	XK_j = C.XK_j
	XK_k = C.XK_k
	XK_l = C.XK_l
	XK_m = C.XK_m
	XK_n = C.XK_n
	XK_o = C.XK_o
	XK_p = C.XK_p
	XK_q = C.XK_q
	XK_r = C.XK_r
	XK_s = C.XK_s
	XK_t = C.XK_t
	XK_u = C.XK_u
	XK_v = C.XK_v
	XK_w = C.XK_w
	XK_x = C.XK_x
	XK_y = C.XK_y
	XK_z = C.XK_z

	XK_A = C.XK_A
	XK_B = C.XK_B
	XK_C = C.XK_C
	XK_D = C.XK_D
	XK_E = C.XK_E
	XK_F = C.XK_F
	XK_G = C.XK_G
	XK_H = C.XK_H
	XK_I = C.XK_I
	XK_J = C.XK_J
	XK_K = C.XK_K
	XK_L = C.XK_L
	XK_M = C.XK_M
	XK_N = C.XK_N
	XK_O = C.XK_O
	XK_P = C.XK_P
	XK_Q = C.XK_Q
	XK_R = C.XK_R
	XK_S = C.XK_S
	XK_T = C.XK_T
	XK_U = C.XK_U
	XK_V = C.XK_V
	XK_W = C.XK_W
	XK_X = C.XK_X
	XK_Y = C.XK_Y
	XK_Z = C.XK_Z

	XK_F1  = C.XK_F1
	XK_F2  = C.XK_F2
	XK_F3  = C.XK_F3
	XK_F4  = C.XK_F4
	XK_F5  = C.XK_F5
	XK_F6  = C.XK_F6
	XK_F7  = C.XK_F7
	XK_F8  = C.XK_F8
	XK_F9  = C.XK_F9
	XK_F10 = C.XK_F10
	XK_F11 = C.XK_F11
	XK_F12 = C.XK_F12
	XK_F13 = C.XK_F13
	XK_F14 = C.XK_F14
	XK_F15 = C.XK_F15
	XK_F16 = C.XK_F16
	XK_F17 = C.XK_F17
	XK_F18 = C.XK_F18
	XK_F19 = C.XK_F19
	XK_F20 = C.XK_F20
	XK_F21 = C.XK_F21
	XK_F22 = C.XK_F22
	XK_F23 = C.XK_F23
	XK_F24 = C.XK_F24

	XK_Shift_L   = C.XK_Shift_L
	XK_Shift_R   = C.XK_Shift_R
	XK_Control_L = C.XK_Control_L
	XK_Control_R = C.XK_Control_R
	XK_Caps_Lock = C.XK_Caps_Lock
	XK_Alt_L     = C.XK_Alt_L
	XK_Alt_R     = C.XK_Alt_R
	XK_Super_L   = C.XK_Super_L
	XK_Super_R   = C.XK_Super_R

	XK_Home      = C.XK_Home
	XK_Left      = C.XK_Left
	XK_Up        = C.XK_Up
	XK_Right     = C.XK_Right
	XK_Down      = C.XK_Down
	XK_Prior     = C.XK_Prior
	XK_Page_Up   = C.XK_Page_Up
	XK_Next      = C.XK_Next
	XK_Page_Down = C.XK_Page_Down
	XK_End       = C.XK_End
	XK_Begin     = C.XK_Begin

	XK_KP_Space     = C.XK_KP_Space
	XK_KP_Tab       = C.XK_KP_Tab
	XK_KP_Enter     = C.XK_KP_Enter
	XK_KP_F1        = C.XK_KP_F1
	XK_KP_F2        = C.XK_KP_F2
	XK_KP_F3        = C.XK_KP_F3
	XK_KP_F4        = C.XK_KP_F4
	XK_KP_Home      = C.XK_KP_Home
	XK_KP_Left      = C.XK_KP_Left
	XK_KP_Up        = C.XK_KP_Up
	XK_KP_Right     = C.XK_KP_Right
	XK_KP_Down      = C.XK_KP_Down
	XK_KP_Prior     = C.XK_KP_Prior
	XK_KP_Page_Up   = C.XK_KP_Page_Up
	XK_KP_Next      = C.XK_KP_Next
	XK_KP_Page_Down = C.XK_KP_Page_Down

	XK_KP_End       = C.XK_KP_End
	XK_KP_Begin     = C.XK_KP_Begin
	XK_KP_Insert    = C.XK_KP_Insert
	XK_KP_Delete    = C.XK_KP_Delete
	XK_KP_Equal     = C.XK_KP_Equal
	XK_KP_Multiply  = C.XK_KP_Multiply
	XK_KP_Separator = C.XK_KP_Separator
	XK_KP_Add       = C.XK_KP_Add
	XK_KP_Subtract  = C.XK_KP_Subtract
	XK_KP_Decimal   = C.XK_KP_Decimal
	XK_KP_Divide    = C.XK_KP_Divide

	XK_KP_0 = C.XK_KP_0
	XK_KP_1 = C.XK_KP_1
	XK_KP_2 = C.XK_KP_2
	XK_KP_3 = C.XK_KP_3
	XK_KP_4 = C.XK_KP_4
	XK_KP_5 = C.XK_KP_5
	XK_KP_6 = C.XK_KP_6
	XK_KP_7 = C.XK_KP_7
	XK_KP_8 = C.XK_KP_8
	XK_KP_9 = C.XK_KP_9

	Success       = C.Success
	XkbUseCoreKbd = C.XkbUseCoreKbd

	BadRequest = C.BadRequest
)

type (
	Display     C.Display
	XVisualInfo C.XVisualInfo
)

func (v *XVisualInfo) Depth() int {
	return int(v.depth)
}

func (v *XVisualInfo) RedMask() int {
	return int(v.red_mask)
}

func (v *XVisualInfo) GreenMask() int {
	return int(v.green_mask)
}

func (v *XVisualInfo) BlueMask() int {
	return int(v.blue_mask)
}

func (v *XVisualInfo) Visualid() int {
	return int(v.visualid)
}

func (c *Display) c() *C.Display {
	return (*C.Display)(unsafe.Pointer(c))
}

var exportedXlibErrorCallback func(err string)

//export chippy_xlib_error_callback
func chippy_xlib_error_callback(d *C.Display, e *C.XErrorEvent) {
	msg := make([]byte, 80)
	C.XGetErrorText(d, C.int(e.error_code), (*C.char)(unsafe.Pointer(&msg[0])), C.int(len(msg)))

	err := fmt.Sprintf("Error %d (%s): request %d.%d", e.error_code, msg, e.request_code, e.minor_code)
	exportedXlibErrorCallback(err)
}

func XSetErrorHandler(c func(err string)) {
	exportedXlibErrorCallback = c
	C.XSetErrorHandler((*[0]byte)(unsafe.Pointer(C.chippy_xlib_error)))
}

func XOpenDisplay(name string) *Display {
	var cstr *C.char
	if len(name) > 0 {
		cstr = C.CString(name)
		defer C.free(unsafe.Pointer(cstr))
	}
	dpy := C.XOpenDisplay(cstr)
	return (*Display)(unsafe.Pointer(dpy))
}

func XGetXCBConnection(d *Display) *Connection {
	c := C.XGetXCBConnection(d.c())
	return (*Connection)(unsafe.Pointer(c))
}

var (
	XlibOwnsEventQueue uint32 = C.XlibOwnsEventQueue
	XCBOwnsEventQueue  uint32 = C.XCBOwnsEventQueue
)

func (d *Display) XSetEventQueueOwner(owner uint32) {
	C.XSetEventQueueOwner(d.c(), owner)
}

func (d *Display) XDefaultScreen() int {
	return int(C.XDefaultScreen(d.c()))
}

type XKeyEvent C.XKeyEvent

func (ev *KeyPressEvent) XKeyEvent(display *Display) *XKeyEvent {
	// Convert between XCB KeyPressEvent and Xlib XKeyEvent; this won't work in
	// all cases but does for the most important one: XLookupKeysym; it is kind
	// of a hacky approach admittedly.
	x := new(XKeyEvent)
	x._type = C.KeyPress
	x.serial = C.ulong(ev.Sequence)
	x.send_event = 0
	x.display = (*C.Display)(unsafe.Pointer(display))
	x.window = (C.Window)(ev.Event)
	x.root = (C.Window)(ev.Root)
	x.subwindow = (C.Window)(ev.Child)
	x.time = C.Time(ev.Time)
	x.x = C.int(ev.EventX)
	x.y = C.int(ev.EventY)
	x.x_root = C.int(ev.RootX)
	x.y_root = C.int(ev.RootY)
	x.state = C.uint(ev.State)
	x.keycode = C.uint(ev.Detail)
	x.same_screen = C.Bool(ev.SameScreen)
	return x
}

// Copy+Pasta
func (ev *KeyReleaseEvent) XKeyEvent(display *Display) *XKeyEvent {
	// Convert between XCB KeyPressEvent and Xlib XKeyEvent; this won't work in
	// all cases but does for the most important one: XLookupKeysym; it is kind
	// of a hacky approach admittidly.
	x := new(XKeyEvent)
	x._type = C.KeyPress
	x.serial = C.ulong(ev.Sequence)
	x.send_event = 0
	x.display = (*C.Display)(unsafe.Pointer(display))
	x.window = (C.Window)(ev.Event)
	x.root = (C.Window)(ev.Root)
	x.subwindow = (C.Window)(ev.Child)
	x.time = C.Time(ev.Time)
	x.x = C.int(ev.EventX)
	x.y = C.int(ev.EventY)
	x.x_root = C.int(ev.RootX)
	x.y_root = C.int(ev.RootY)
	x.state = C.uint(ev.State)
	x.keycode = C.uint(ev.Detail)
	x.same_screen = C.Bool(ev.SameScreen)
	return x
}
func (d *Display) XLookupKeysym(ev *XKeyEvent, index int) Keysym {
	ret := C.XLookupKeysym(
		(*C.XKeyEvent)(unsafe.Pointer(ev)),
		C.int(index),
	)
	return Keysym(ret)
}

type XMappingEvent C.XMappingEvent

func (d *Display) XRefreshKeyboardMapping(ev *XMappingEvent) {
	C.XRefreshKeyboardMapping((*C.XMappingEvent)(unsafe.Pointer(ev)))
}

func XConvertCase(keysym Keysym, lower, upper *Keysym) {
	C.XConvertCase(
		C.KeySym(keysym),
		(*C.KeySym)(unsafe.Pointer(lower)),
		(*C.KeySym)(unsafe.Pointer(upper)),
	)
}

func (d *Display) XKeysymToKeycode(keysym Keysym) Keycode {
	return Keycode(C.XKeysymToKeycode(
		d.c(),
		C.KeySym(keysym),
	))
}

type EXIM C.struct_XIM
type XIM struct {
	*EXIM
}

func (d *Display) XOpenIM(rdb *C.struct__XrmHashBucketRec, resName, resClass *C.char) *XIM {
	c := C.XOpenIM(d.c(), rdb, resName, resClass)
	if c == nil {
		return nil
	}
	xim := new(XIM)
	xim.EXIM = (*EXIM)(unsafe.Pointer(c))
	runtime.SetFinalizer(xim, func(f *XIM) {
		C.XCloseIM((C.XIM)(unsafe.Pointer(f.EXIM)))
	})
	return xim
}

type EXIC C.struct_XIC
type XIC struct {
	*EXIC
}

func (d *Display) CreateIC(im *XIM, w Window) *XIC {
	c := C.chippy_CreateIC((C.XIM)(unsafe.Pointer(im.EXIM)), d.c(), C.Window(w))
	if c == nil {
		return nil
	}
	xic := new(XIC)
	xic.EXIC = (*EXIC)(unsafe.Pointer(c))
	runtime.SetFinalizer(xic, func(f *XIC) {
		C.XDestroyIC((C.XIC)(unsafe.Pointer(f.EXIC)))
	})
	return xic
}

type Status C.int

func (d *Display) XkbGetIndicatorState(deviceSpec uint32) (state uint32, status Status) {
	status = Status(C.XkbGetIndicatorState(
		d.c(),
		C.uint(deviceSpec),
		(*C.uint)(unsafe.Pointer(&state)),
	))
	return
}

func (d *Display) XLookupString(kev *XKeyEvent) (s string, keysym Keysym) {
	data := make([]byte, 32)
	bytesLen := int(C.XLookupString(
		(*C.XKeyEvent)(unsafe.Pointer(kev)),
		(*C.char)(unsafe.Pointer(&data[0])),
		C.int(len(data)),
		(*C.KeySym)(unsafe.Pointer(&keysym)),
		nil,
	))
	if bytesLen > 0 {
		s = string(data[:bytesLen])
	}
	return
}

func (d *Display) XFlush() {
	C.XFlush(d.c())
}

func (d *Display) XSync(x bool) {
	cx := C.Bool(C.False)
	if x {
		cx = C.True
	}
	C.XSync(d.c(), cx)
}

func (d *Display) Xutf8LookupString(ic *XIC, kev *XKeyEvent) (s string, keysym Keysym, hasKeysym bool) {
	var cstat Status

	data := make([]byte, 1)
	bytesLen := int(C.Xutf8LookupString(
		(C.XIC)(unsafe.Pointer(ic.EXIC)),
		(*C.XKeyPressedEvent)(unsafe.Pointer(kev)),
		(*C.char)(unsafe.Pointer(&data[0])),
		C.int(len(data)),
		(*C.KeySym)(unsafe.Pointer(&keysym)),
		(*C.Status)(unsafe.Pointer(&cstat)),
	))
	if cstat == C.XBufferOverflow {
		cstat = 0
		data = make([]byte, bytesLen)
		bytesLen = int(C.Xutf8LookupString(
			(C.XIC)(unsafe.Pointer(ic.EXIC)),
			(*C.XKeyPressedEvent)(unsafe.Pointer(kev)),
			(*C.char)(unsafe.Pointer(&data[0])),
			C.int(len(data)),
			(*C.KeySym)(unsafe.Pointer(&keysym)),
			(*C.Status)(unsafe.Pointer(&cstat)),
		))
	}

	if cstat == C.XLookupKeySym || cstat == C.XLookupBoth {
		hasKeysym = true
	}
	if cstat == C.XLookupChars || cstat == C.XLookupBoth {
		if bytesLen > 0 {
			s = string(data[:bytesLen])
		}
	}
	return
}

func XInitThreads() {
	C.XInitThreads()
}
