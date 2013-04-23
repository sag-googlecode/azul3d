// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Do not use - Wrappers around very specific X11 APIs.
package x11

/*
#include <stdlib.h>
#include <X11/Xlib.h>

#cgo LDFLAGS: -lX11

void init();
*/
import "C"

import (
	"unsafe"
)

type (
	XAnyEvent              C.XAnyEvent
	XKeyEvent              C.XKeyEvent
	XButtonEvent           C.XButtonEvent
	XMotionEvent           C.XMotionEvent
	XCrossingEvent         C.XCrossingEvent
	XFocusChangeEvent      C.XFocusChangeEvent
	XExposeEvent           C.XExposeEvent
	XGraphicsExposeEvent   C.XGraphicsExposeEvent
	XNoExposeEvent         C.XNoExposeEvent
	XVisibilityEvent       C.XVisibilityEvent
	XCreateWindowEvent     C.XCreateWindowEvent
	XDestroyWindowEvent    C.XDestroyWindowEvent
	XUnmapEvent            C.XUnmapEvent
	XMapEvent              C.XMapEvent
	XMapRequestEvent       C.XMapRequestEvent
	XReparentEvent         C.XReparentEvent
	XConfigureEvent        C.XConfigureEvent
	XGravityEvent          C.XGravityEvent
	XResizeRequestEvent    C.XResizeRequestEvent
	XConfigureRequestEvent C.XConfigureRequestEvent
	XCirculateEvent        C.XCirculateEvent
	XCirculateRequestEvent C.XCirculateRequestEvent
	XPropertyEvent         C.XPropertyEvent
	XSelectionClearEvent   C.XSelectionClearEvent
	XSelectionRequestEvent C.XSelectionRequestEvent
	XSelectionEvent        C.XSelectionEvent
	XColormapEvent         C.XColormapEvent
	XClientMessageEvent    C.XClientMessageEvent
	XMappingEvent          C.XMappingEvent
	XErrorEvent            C.XErrorEvent
	XKeymapEvent           C.XKeymapEvent
	XGenericEvent          C.XGenericEvent
	XGenericEventCookie    C.XGenericEventCookie
	XSetWindowAttributes   C.XSetWindowAttributes
	XWindowAttributes      C.XWindowAttributes
	XWindowChanges         C.XWindowChanges

	Display  C.Display
	Screen   C.Screen
	Window   C.Window
	Drawable C.Drawable
	Colormap C.Colormap
	Visual   C.Visual
	Atom     C.Atom
	Cursor   C.Cursor
	Pixmap   C.Pixmap
	GC       C.GC
	Time     C.Time
	KeySym   C.KeySym
	KeyCode  C.KeyCode
)

const (
	CurrentTime = C.CurrentTime
	None        = C.None
)

// Wrapper for XEvent union
type XEvent struct {
	_type             C.int
	xany              XAnyEvent
	xkey              XKeyEvent
	xbutton           XButtonEvent
	xmotion           XMotionEvent
	xcrossing         XCrossingEvent
	xfocus            XFocusChangeEvent
	xexpose           XExposeEvent
	xgraphicsexpose   XGraphicsExposeEvent
	xnoexpose         XNoExposeEvent
	xvisibility       XVisibilityEvent
	xcreatewindow     XCreateWindowEvent
	xdestroywindow    XDestroyWindowEvent
	xunmap            XUnmapEvent
	xmap              XMapEvent
	xmaprequest       XMapRequestEvent
	xreparent         XReparentEvent
	xconfigure        XConfigureEvent
	xgravity          XGravityEvent
	xresizerequest    XResizeRequestEvent
	xconfigurerequest XConfigureRequestEvent
	xcirculate        XCirculateEvent
	xcirculaterequest XCirculateRequestEvent
	xproperty         XPropertyEvent
	xselectionclear   XSelectionClearEvent
	xselectionrequest XSelectionRequestEvent
	xselection        XSelectionEvent
	xcolormap         XColormapEvent
	xclient           XClientMessageEvent
	xmapping          XMappingEvent
	xerror            XErrorEvent
	xkeymap           XKeymapEvent
	xgeneric          XGenericEvent
	xcookie           XGenericEventCookie
	pad               [24]C.long
}

func cBool(b bool) C.Bool {
	if b == true {
		return C.True
	}
	return C.False
}

var xErrorHandler func(*Display, *XErrorEvent)

//export fakeXErrorHandler
func fakeXErrorHandler(d, e unsafe.Pointer) {
	if xErrorHandler != nil {
		xErrorHandler((*Display)(d), (*XErrorEvent)(e))
	}
}

func XSetErrorHandler(handler func(*Display, *XErrorEvent)) {
	xErrorHandler = handler
}

func init() {
	C.init()
}

func elementAtIndex(ptr unsafe.Pointer, index uintptr, T uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + (uintptr(index) * T))
}

func XInitThreads() {
	C.XInitThreads()
}

// display_name defaults to DISPLAY environment variable.
// returns nil upon failure
func XOpenDisplay(display_name string) *Display {
	if len(display_name) > 0 {
		cstr := C.CString(display_name)
		defer C.free(unsafe.Pointer(cstr))
		return (*Display)(C.XOpenDisplay(cstr))
	}
	return (*Display)(C.XOpenDisplay(nil))
}

// Can generate BadGC error
func XCloseDisplay(display *Display) int {
	return int(C.XCloseDisplay((*C.Display)(display)))
}

func XFlush(display *Display) int {
	return int(C.XFlush((*C.Display)(display)))
}

func XSync(display *Display, discard bool) int {
	return int(C.XSync((*C.Display)(display), cBool(discard)))
}

func XScreenNumberOfScreen(screen *Screen) int {
	return int(C.XScreenNumberOfScreen((*C.Screen)(screen)))
}

func XScreenCount(display *Display) int {
	return int(C.XScreenCount((*C.Display)(display)))
}

func XWidthOfScreen(screen *Screen) int {
	return int(C.XWidthOfScreen((*C.Screen)(screen)))
}

func XWidthMMOfScreen(screen *Screen) int {
	return int(C.XWidthMMOfScreen((*C.Screen)(screen)))
}

func XHeightOfScreen(screen *Screen) int {
	return int(C.XHeightOfScreen((*C.Screen)(screen)))
}

func XHeightMMOfScreen(screen *Screen) int {
	return int(C.XHeightMMOfScreen((*C.Screen)(screen)))
}

func XScreenOfDisplay(display *Display, screen_number int) *Screen {
	return (*Screen)(C.XScreenOfDisplay((*C.Display)(display), C.int(screen_number)))
}

func XDefaultScreenOfDisplay(display *Display) *Screen {
	return (*Screen)(C.XDefaultScreenOfDisplay((*C.Display)(display)))
}

func XDefaultRootWindow(display *Display) Window {
	return Window(C.XDefaultRootWindow((*C.Display)(display)))
}

func XCreateColormap(display *Display, w Window, visual *Visual, alloc int) Colormap {
	return Colormap(C.XCreateColormap((*C.Display)(display), C.Window(w), (*C.Visual)(visual), C.int(alloc)))
}

func XCreateWindow(display *Display, parent Window, x, y int, width, height, border_width uint, depth int, class uint, visual *Visual, valuemask uint64, attributes *XSetWindowAttributes) Window {
	return Window(C.XCreateWindow((*C.Display)(display), C.Window(parent), C.int(x), C.int(y), C.uint(width), C.uint(height), C.uint(border_width), C.int(depth), C.uint(class), (*C.Visual)(visual), C.ulong(valuemask), (*C.XSetWindowAttributes)(attributes)))
}

func XRootWindow(display *Display, screen int) Window {
	return Window(C.XRootWindow((*C.Display)(display), C.int(screen)))
}

func XMapWindow(display *Display, window Window) int {
	return int(C.XMapWindow((*C.Display)(display), C.Window(window)))
}

func XUnmapWindow(display *Display, window Window) int {
	return int(C.XUnmapWindow((*C.Display)(display), C.Window(window)))
}

func XStoreName(display *Display, window Window, title string) int {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.XStoreName((*C.Display)(display), C.Window(window), cstr))
}

func XIconifyWindow(display *Display, window Window, screen_number int) int {
	return int(C.XIconifyWindow((*C.Display)(display), C.Window(window), C.int(screen_number)))
}

func XRaiseWindow(display *Display, window Window) int {
	return int(C.XRaiseWindow((*C.Display)(display), C.Window(window)))
}

func XWithdrawWindow(display *Display, window Window, screen_number int) int {
	return int(C.XWithdrawWindow((*C.Display)(display), C.Window(window), C.int(screen_number)))
}

func XFreeColormap(display *Display, cmap Colormap) {
	C.XFreeColormap((*C.Display)(display), C.Colormap(cmap))
}

func XMoveWindow(display *Display, w Window, x, y int) int {
	return int(C.XMoveWindow((*C.Display)(display), C.Window(w), C.int(x), C.int(y)))
}

func XResizeWindow(display *Display, w Window, x, y uint) int {
	return int(C.XResizeWindow((*C.Display)(display), C.Window(w), C.uint(x), C.uint(y)))
}

func XInternAtom(display *Display, atom string, only_if_exists bool) Atom {
	cstr := C.CString(atom)
	defer C.free(unsafe.Pointer(cstr))
	return Atom(C.XInternAtom((*C.Display)(display), cstr, cBool(only_if_exists)))
}

func XChangeProperty(display *Display, window Window, property, _type Atom, format, mode int, data unsafe.Pointer, nelements int) int {
	return int(C.XChangeProperty((*C.Display)(display), C.Window(window), C.Atom(property), C.Atom(_type), C.int(format), C.int(mode), (*C.uchar)(data), C.int(nelements)))
}

func XWarpPointer(display *Display, src_w, dest_w Window, src_x, src_y int, src_width, src_height uint, dest_x, dest_y int) int {
	return int(C.XWarpPointer((*C.Display)(display), C.Window(src_w), C.Window(dest_w), C.int(src_x), C.int(src_y), C.uint(src_width), C.uint(src_height), C.int(dest_x), C.int(dest_y)))
}

func XReconfigureWMWindow(display *Display, w Window, screen_number int, value_mask uint, values *XWindowChanges) int {
	return int(C.XReconfigureWMWindow((*C.Display)(display), C.Window(w), C.int(screen_number), C.uint(value_mask), (*C.XWindowChanges)(values)))
}

func XSendEvent(display *Display, w Window, propagate bool, event_mask int64, event_send *XEvent) int {
	return int(C.XSendEvent((*C.Display)(display), C.Window(w), cBool(propagate), C.long(event_mask), (*C.XEvent)(unsafe.Pointer(event_send))))
}

func XConnectionNumber(display *Display) int {
	return int(C.XConnectionNumber((*C.Display)(display)))
}

func XSelectInput(display *Display, w Window, mask int64) int {
	return int(C.XSelectInput((*C.Display)(display), C.Window(w), C.long(mask)))
}

func XDestroyWindow(display *Display, w Window) int {
	return int(C.XDestroyWindow((*C.Display)(display), C.Window(w)))
}

func XNextEvent(display *Display) (int, *XEvent) {
	var ev *C.XEvent
	ret := C.XNextEvent((*C.Display)(display), ev)
	return int(ret), (*XEvent)(unsafe.Pointer(ev))
}

func XPending(display *Display) int {
	return int(C.XPending((*C.Display)(display)))
}

func XGetWindowAttributes(display *Display, window Window) (int, *XWindowAttributes) {
	var att *C.XWindowAttributes
	ret := C.XGetWindowAttributes((*C.Display)(display), C.Window(window), att)
	return int(ret), (*XWindowAttributes)(att)
}

func XSetWMProtocols(display *Display, window Window, protocols *Atom, count int) int {
	return int(C.XSetWMProtocols((*C.Display)(display), C.Window(window), (*C.Atom)(protocols), C.int(count)))
}

func XGetAtomName(display *Display, atom Atom) string {
	return C.GoString(C.XGetAtomName((*C.Display)(display), C.Atom(atom)))
}

// returns status, major_opcode_return, first_event_return, first_error_return
func XQueryExtension(display *Display, name string) (int, int, int, int) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var major, event, error C.int
	ret := C.XQueryExtension((*C.Display)(display), cstr, &major, &event, &error)
	return int(ret), int(major), int(event), int(error)
}

func XGetEventData(display *Display, cookie *XGenericEventCookie) int {
	return int(C.XGetEventData((*C.Display)(display), (*C.XGenericEventCookie)(cookie)))
}

func XFreeEventData(display *Display, cookie *XGenericEventCookie) {
	C.XFreeEventData((*C.Display)(display), (*C.XGenericEventCookie)(cookie))
}

func XDefineCursor(display *Display, window Window, cursor Cursor) int {
	return int(C.XDefineCursor((*C.Display)(display), C.Window(window), C.Cursor(cursor)))
}

func XFreePixmap(display *Display, pixmap Pixmap) {
	C.XFreePixmap((*C.Display)(display), C.Pixmap(pixmap))
}

func XFreeGC(display *Display, gc GC) {
	C.XFreeGC((*C.Display)(display), C.GC(gc))
}

func XGrabPointer(display *Display, grab_window Window, owner_events bool, event_mask uint, pointer_mode, keyboard_mode int, confine_to Window, cursor Cursor, time Time) int {
	return int(C.XGrabPointer((*C.Display)(display), C.Window(grab_window), cBool(owner_events), C.uint(event_mask), C.int(pointer_mode), C.int(keyboard_mode), C.Window(confine_to), C.Cursor(cursor), C.Time(time)))
}

func XUngrabPointer(display *Display, time Time) int {
	return int(C.XUngrabPointer((*C.Display)(display), C.Time(time)))
}

func XGetKeyboardMapping(display *Display, first_keycode KeyCode, keycode_count int) (int, []KeySym) {
	var keysyms_per_keycode C.int
	var sym *C.KeySym = C.XGetKeyboardMapping((*C.Display)(display), C.KeyCode(first_keycode), C.int(keycode_count), &keysyms_per_keycode)
	defer C.XFree(unsafe.Pointer(sym))

	slice := make([]KeySym, keycode_count-1)
	for i := 0; i < keycode_count; i++ {
		element := elementAtIndex(unsafe.Pointer(sym), uintptr(i), unsafe.Sizeof(C.KeySym(0)))
		slice = append(slice, *(*KeySym)(element))
	}
	return int(keysyms_per_keycode), slice
}

func XAutoRepeatOn(display *Display) {
	C.XAutoRepeatOn((*C.Display)(display))
}

func XAutoRepeatOff(display *Display) {
	C.XAutoRepeatOff((*C.Display)(display))
}

/*
func c_setAtomState(display *c_Display, w c_Window, atom string, state bool) {
	cstr := C.CString(atom)
	C.setAtomState(display.C(), C.Window(w), cstr, c_bool(state))
}


// Note: On allocation, the hotspot and the pixels are left uninitialized. The size is set to the maximum of width and height.
func c_XcursorImageCreate(width, height int) *c_XcursorImage {
    return (*c_XcursorImage)(C.XcursorImageCreate(C.int(width), C.int(height)))
}

func c_XcursorImageDestroy(image *c_XcursorImage) {
    C.XcursorImageDestroy((*C.XcursorImage)(image))
}

func c_XcursorImageLoadCursor(display *c_Display, image *c_XcursorImage) c_Cursor {
    return c_Cursor(C.XcursorImageLoadCursor(display.C(), (*C.XcursorImage)(image)))
}


func c_keysym2ucs(keysym c_KeySym) int {
    return int(C.keysym2ucs(keysym.C()))
}

func c_getKeySym(event c_XKeyEvent) c_KeySym {
    var keysym C.KeySym
    ev := event.C()
    C.XLookupString(&ev, nil, 0, &keysym, nil)
    return c_KeySym(keysym)
}

// translate keysym to keyboard.Code
func c_findKeyboardCode(keysym c_KeySym) keyboard.Code {
    switch(keysym) {
        case C.XK_F1:
            return keyboard.F1
        case C.XK_F2:
            return keyboard.F2
        case C.XK_F3:
            return keyboard.F3
        case C.XK_F4:
            return keyboard.F4
        case C.XK_F5:
            return keyboard.F5
        case C.XK_F6:
            return keyboard.F6
        case C.XK_F7:
            return keyboard.F7
        case C.XK_F8:
            return keyboard.F8
        case C.XK_F9:
            return keyboard.F9
        case C.XK_F10:
            return keyboard.F10
        case C.XK_F11:
            return keyboard.F11
        case C.XK_F12:
            return keyboard.F12
        case C.XK_F13:
            return keyboard.F13
        case C.XK_F14:
            return keyboard.F14
        case C.XK_F15:
            return keyboard.F15
        case C.XK_F16:
            return keyboard.F16
        case C.XK_F17:
            return keyboard.F17
        case C.XK_F18:
            return keyboard.F18
        case C.XK_F19:
            return keyboard.F19
        case C.XK_F20:
            return keyboard.F20
        case C.XK_F21:
            return keyboard.F21
        case C.XK_F22:
            return keyboard.F22
        case C.XK_F23:
            return keyboard.F23
        case C.XK_F24:
            return keyboard.F24
        case C.XK_F25:
            return keyboard.F25
        case C.XK_F26:
            return keyboard.F26
        case C.XK_F27:
            return keyboard.F27
        case C.XK_F28:
            return keyboard.F28
        case C.XK_F29:
            return keyboard.F29
        case C.XK_F30:
            return keyboard.F30
        case C.XK_F31:
            return keyboard.F31
        case C.XK_F32:
            return keyboard.F32
        case C.XK_F33:
            return keyboard.F33
        case C.XK_F34:
            return keyboard.F34
        case C.XK_F35:
            return keyboard.F35


        case C.XK_Escape:
            return keyboard.Escape
        case C.XK_Tab:
            return keyboard.Tab
        case C.XK_Caps_Lock:
            return keyboard.CapsLock
        case C.XK_Scroll_Lock:
            return keyboard.ScrollLock
        case C.XK_Print:
            return keyboard.PrintScreen
        case C.XK_BackSpace:
            return keyboard.Backspace
        case C.XK_Menu:
            return keyboard.Menu
        case C.XK_Pause:
            return keyboard.Pause
        case C.XK_Shift_L:
            return keyboard.LeftShift
        case C.XK_Shift_R:
            return keyboard.RightShift
        case C.XK_Control_L:
            return keyboard.LeftControl
        case C.XK_Control_R:
            return keyboard.RightControl
        case C.XK_Alt_L:
            return keyboard.LeftAlt
        case C.XK_Alt_R:
            return keyboard.RightAlt
        case C.XK_Super_L:
            return keyboard.LeftSuper
        case C.XK_Super_R:
            return keyboard.RightSuper

        case C.XK_0:
            return keyboard.Number0
        case C.XK_1:
            return keyboard.Number1
        case C.XK_2:
            return keyboard.Number2
        case C.XK_3:
            return keyboard.Number3
        case C.XK_4:
            return keyboard.Number4
        case C.XK_5:
            return keyboard.Number5
        case C.XK_6:
            return keyboard.Number6
        case C.XK_7:
            return keyboard.Number7
        case C.XK_8:
            return keyboard.Number8
        case C.XK_9:
            return keyboard.Number9
        case C.XK_slash:
            return keyboard.ForwardSlash
        case C.XK_asterisk:
            return keyboard.Asterisk
        case C.XK_plus:
            return keyboard.Plus
        case C.XK_minus:
            return keyboard.HyphenMinus
        case C.XK_Delete:
            return keyboard.Delete
        case C.XK_Insert:
            return keyboard.Insert
        case C.XK_Num_Lock:
            return keyboard.NumLock
        case C.XK_Home:
            return keyboard.Home
        case C.XK_Page_Down:
            return keyboard.PageDown
        case C.XK_Page_Up:
            return keyboard.PageUp
        case C.XK_End:
            return keyboard.End
        case C.XK_Return:
            return keyboard.Enter
        case C.XK_Up:
            return keyboard.Up
        case C.XK_Down:
            return keyboard.Down
        case C.XK_Left:
            return keyboard.Left
        case C.XK_Right:
            return keyboard.Right
        case C.XK_period:
            return keyboard.Period

        case C.XK_KP_0:
            return keyboard.NumPadNumber0
        case C.XK_KP_1:
            return keyboard.NumPadNumber1
        case C.XK_KP_2:
            return keyboard.NumPadNumber2
        case C.XK_KP_3:
            return keyboard.NumPadNumber3
        case C.XK_KP_4:
            return keyboard.NumPadNumber4
        case C.XK_KP_5:
            return keyboard.NumPadNumber5
        case C.XK_KP_6:
            return keyboard.NumPadNumber6
        case C.XK_KP_7:
            return keyboard.NumPadNumber7
        case C.XK_KP_8:
            return keyboard.NumPadNumber8
        case C.XK_KP_9:
            return keyboard.NumPadNumber9
        case C.XK_KP_Divide:
            return keyboard.NumPadForwardSlash
        case C.XK_KP_Multiply:
            return keyboard.NumPadAsterisk
        case C.XK_KP_Add:
            return keyboard.NumPadPlus
        case C.XK_KP_Subtract:
            return keyboard.NumPadHyphenMinus
        case C.XK_KP_Delete:
            return keyboard.NumPadDelete
        case C.XK_KP_Insert:
            return keyboard.NumPadInsert
        case C.XK_KP_Home:
            return keyboard.NumPadHome
        case C.XK_KP_Page_Up:
            return keyboard.NumPadPageUp
        case C.XK_KP_Page_Down:
            return keyboard.NumPadPageDown
        case C.XK_KP_End:
            return keyboard.NumPadEnd
        case C.XK_KP_Enter:
            return keyboard.NumPadEnter
        case C.XK_KP_Up:
            return keyboard.NumPadUp
        case C.XK_KP_Down:
            return keyboard.NumPadDown
        case C.XK_KP_Left:
            return keyboard.NumPadLeft
        case C.XK_KP_Right:
            return keyboard.NumPadRight
        case C.XK_KP_Decimal:
            return keyboard.NumPadPeriod
    }
    return keyboard.Unknown
}

func c_getWindowExtents(display *c_Display, window c_Window) [4]int16 {
	var _type C.Atom
	var format C.int
	var nItem, bytesAfter C.ulong
	var properties *C.uchar

    netFrameExtents := c_XInternAtom(display, "_NET_FRAME_EXTENTS", true)
    C.XGetWindowProperty(display.C(), window.C(), netFrameExtents.C(), 0, ^C.long(0), C.False, C.AnyPropertyType, &_type, &format, &nItem, &bytesAfter, &properties)

    extents := [4]int16{}
    for i := 0; i < int(nItem); i++ {
        property := C.propertyAtIndex(properties, C.int(i))
        extents[i] = int16(property)
    }
    return extents


    /*
_NET_FRAME_EXTENTS, left, right, top, bottom, CARDINAL[4]/32
	wmState := c_XInternAtom(display, "_NET_WM_STATE", true)
	C.XGetWindowProperty(display.C(), window.C(), wmState.C(), 0, ^C.long(0), C.False, C.AnyPropertyType, &_type, &format, &nItem, &bytesAfter, &properties)
	atoms := []c_Atom{}
	for i := 0; i < int(nItem); i++ {
		property := C.propertyAtIndex(properties, C.int(i))
		atoms = append(atoms, c_Atom(property))
	}
	return atoms
}

func c_setWindowIcon(display *c_Display, window c_Window, img image.Image) error {
    // So it looks like the best thing we can do is give X11 a few different icon choices (32x32 & 128x128)
    // and just hope the WM chooses the right one or resizes it. From everything I've seen alot of WM's do
    // a really poor job at handling window icons.. so.. results may vary...

    bufferSize := 2+32*32 + 2+128*128
    buffer := [2+32*32 + 2+128*128]uint{}
    i := 0

    // Firstly, we copy the 32x32 image into the buffer
    img32 := resize.Resize(img, img.Bounds(), 32, 32)
    buffer[i] = uint(32)
    i++
    buffer[i] = uint(32)
    i++
    for x := 0; x < 32; x++ {
        for y := 0; y < 32; y++ {
            ri, gi, bi, ai := img32.At(x, y).RGBA()
            r := uint8(ri)
            g := uint8(gi)
            b := uint8(bi)
            a := uint8(ai)
            pixel := uint(b) | uint(g) << 8 | uint(r) << 16 | uint(a) << 24

            buffer[i] = pixel
            i++
        }
    }

    // And now a 128x128 copy, incase it supports higher resolution ones
    buffer[i] = uint(128)
    i++
    buffer[i] = uint(128)
    i++
    for x := 0; x < 128; x++ {
        for y := 0; y < 128; y++ {
            ri, gi, bi, ai := img.At(x, y).RGBA()
            r := uint8(ri)
            g := uint8(gi)
            b := uint8(bi)
            a := uint8(ai)
            pixel := uint(b) | uint(g) << 8 | uint(r) << 16 | uint(a) << 24

            buffer[i] = pixel
            i++
        }
    }



    netWmIcon := c_XInternAtom(xDisplay, "_NET_WM_ICON", false)
    //C.XDeleteProperty(display.C(), window.C(), netWmIcon.C())
    C.XChangeProperty(display.C(), window.C(), netWmIcon.C(), C.XA_CARDINAL, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&buffer)), C.int(bufferSize))
    return nil
}

func c_setCursorParams(image *c_XcursorImage, xhot, yhot uint, width, height uint, pixels []uint) {
    // Sanity check here
    if uint(len(pixels)) != width * height {
        panic("len(pixels) != width * height")
    }

    image.xhot = C.XcursorDim(xhot)
    image.yhot = C.XcursorDim(yhot)

    image.width = C.XcursorDim(width)
    image.height = C.XcursorDim(height)

    // Allocate pixels pointer
    image.pixels = (*C.XcursorPixel)(C.malloc(C.size_t(uint(unsafe.Sizeof(C.XcursorPixel(0))) * width * height)))

    for i := range pixels {
        C.setCursorImagePixel((*C.XcursorImage)(image), C.int(i), C.uint(pixels[C.int(i)]))
    }
}

func c_xCursorFromImage(display *c_Display, image image.Image, pointerX, pointerY uint) c_Cursor {
    width, height := image.Bounds().Max.X, image.Bounds().Max.Y

    pixels := make([]uint, width * height)
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            ri, gi, bi, ai := image.At(x, y).RGBA()
            r := uint8(ri)
            g := uint8(gi)
            b := uint8(bi)
            a := uint8(ai)
            pixel := uint(b) | uint(g) << 8 | uint(r) << 16 | uint(a) << 24
            pixels[width*y + x] = pixel
        }
    }

    // Note: On allocation, the hotspot and the pixels are left uninitialized. The size is set to the maximum of width and height.
    cursorImage := c_XcursorImageCreate(width, height)
    defer c_XcursorImageDestroy(cursorImage)
    c_setCursorParams(cursorImage, pointerX, pointerY, uint(width), uint(height), pixels)

    return c_XcursorImageLoadCursor(display, cursorImage)
}

func c_waitForEvent(display *c_Display, sec, usec int) bool {
	if C.waitForEvent(display.C(), C.int(sec), C.int(usec)) == 0 {
		return false
	}
	return true
}

func c_isDeleteWindowAtom(display *c_Display, e c_XClientMessageEvent) bool {
	if C.isDeleteWindowAtom(display.C(), e.C()) == 0 {
		return false
	}
	return true
}

func c_GetNetWmStates(display *c_Display, window c_Window) []c_Atom {
	var _type C.Atom
	var format C.int
	var nItem, bytesAfter C.ulong
	var properties *C.uchar

	/*
	   long_offset	 Specifies the offset in the specified property (in 32-bit quantities) where the data is to be retrieved.
	   long_length	 Specifies the length in 32-bit multiples of the data to be retrieved.
	   delete	 Specifies a Boolean value that determines whether the property is deleted.
	   req_type	 Specifies the atom identifier associated with the property type or AnyPropertyType.
	   actual_type_return	 Returns the atom identifier that defines the actual type of the property.
	   actual_format_return	 Returns the actual format of the property.
	   nitems_return	 Returns the actual number of 8-bit, 16-bit, or 32-bit items stored in the prop_return data.
	   bytes_after_return	 Returns the number of bytes remaining to be read in the property if a partial read was performed.
	   prop_return	 Returns the data in the specified format.

	wmState := c_XInternAtom(display, "_NET_WM_STATE", true)
	C.XGetWindowProperty(display.C(), window.C(), wmState.C(), 0, ^C.long(0), C.False, C.AnyPropertyType, &_type, &format, &nItem, &bytesAfter, &properties)
	atoms := []c_Atom{}
	for i := 0; i < int(nItem); i++ {
		property := C.propertyAtIndex(properties, C.int(i))
		atoms = append(atoms, c_Atom(property))
	}
	return atoms
}

func c_setXIEventMask(display *c_Display) {
	C.setXIEventMask(display.C())
}

func c_setSizeHints(display *c_Display, window c_Window, min_width, min_height, max_width, max_height uint16) error {
    var hints *C.XSizeHints = C.XAllocSizeHints()
    hints.min_width = C.int(min_width)
    hints.min_height = C.int(min_height)
    hints.flags |= C.PMinSize
    if max_width != 0 || max_height != 0 {
        hints.max_width = C.int(max_width)
        hints.max_height = C.int(max_height)
        hints.flags |= C.PMaxSize
    }
    C.XSetWMNormalHints(display.C(), window.C(), hints)
    return nil
}

func c_getPointerPosition(display *c_Display, window c_Window) (bool, int, int) {
    var root, child C.Window
    var root_x, root_y, win_x, win_y C.int
    var mask C.uint
    if C.XQueryPointer(display.C(), window.C(), &root, &child, &root_x, &root_y, &win_x, &win_y, &mask) == 0 {
        return false, 0, 0
    }
    return true, int(win_x), int(win_y)
}

/*
Atom wmState = XInternAtom(display, "_NET_WM_STATE", True);
Atom type;
int format;
unsigned long nItem, bytesAfter;
unsigned char *properties = NULL;
(long*)(properties)[nItem]

XGetWindowProperty(display, window, wmState, 0, (~0L), False, AnyPropertyType, &type, &format, &nItem, &bytesAfter, &properties);
int iItem;
usleep(1000000);
//printf("itemmmmmm=%d\n",nItem);
for (iItem = 0; iItem < nItem; ++iItem)
	printf("property=%ld\n",(long*)(properties)[nItem]);

c_XGetNetWmStates

func c_XF86VidModeSetViewPort(display *c_Display, screen, x, y int) {
	C.XF86VidModeSetViewPort(display.C(), C.int(screen), C.int(x), C.int(y))
}

func c_XF86VidModeGetAllModeLines(display *c_Display, screen int) []*c_XF86VidModeModeInfo {
	var modecount C.int
	var modelines **C.XF86VidModeModeInfo
	C.XF86VidModeGetAllModeLines(display.C(), C.int(screen), &modecount, &modelines)
	modes := []*c_XF86VidModeModeInfo{}
	for i := 0; i < int(modecount); i++ {
		mode := C.vidModeAtIndex(modelines, C.int(i))
		modes = append(modes, (*c_XF86VidModeModeInfo)(mode))
	}
	return modes
}

func c_XF86VidModeSwitchToMode(display *c_Display, screen int, mi *c_XF86VidModeModeInfo) error {
	if C.XF86VidModeSwitchToMode(display.C(), C.int(screen), (*C.XF86VidModeModeInfo)(mi)) == 0 {
		return errors.New("Unable to switch to video mode; XF86VidModeSwitchToMode() failed!")
	}
	return nil
}

// Helper to call XF86VidModeGetGammaRampSize
func c_XF86VidModeGetGammaRampSize(display *c_Display, screen int) (int, error) {
	var size C.int
	ret := C.XF86VidModeGetGammaRampSize(display.C(), C.int(screen), &size)
	if ret == C.False {
		return 0, errors.New("Call failed XF86VidModeGetGammaRampSize()")
	}
	return int(size), nil
}

func c_XF86VidModeSetGammaRamp(display *c_Display, screen int, red, green, blue [256]uint16) error {
	// We make assumption that gamma ramp is 256
	size, err := c_XF86VidModeGetGammaRampSize(xDisplay, screen)
	if err != nil {
		return errors.New("Unable to set gamma ramp; XF86VidModeGetGammaRampSize() failed")
	}
	if size != 256 {
		return errors.New("Unable to set gamma ramp; Gamma ramp size > 256")
	}

	r := (*C.ushort)(unsafe.Pointer(&red))
	g := (*C.ushort)(unsafe.Pointer(&green))
	b := (*C.ushort)(unsafe.Pointer(&blue))
	worked := C.XF86VidModeSetGammaRamp(display.C(), C.int(screen), C.int(size), r, g, b)
	if worked == C.False {
		return errors.New("Call failed XF86VidModeSetGammaRamp()")
	}
	return nil
}

func c_XF86VidModeGetGammaRamp(display *c_Display, screen int) ([256]uint16, [256]uint16, [256]uint16, error) {
	red, green, blue := [256]uint16{}, [256]uint16{}, [256]uint16{}

	// We make assumption that gamma ramp is 256
	size, err := c_XF86VidModeGetGammaRampSize(xDisplay, screen)
	if err != nil {
		return red, green, blue, errors.New("Unable to set gamma ramp; XF86VidModeGetGammaRampSize() failed")
	}
	if size != 256 {
		return red, green, blue, errors.New("Unable to set gamma ramp; Gamma ramp size > 256")
	}

	var r, g, b [256]C.ushort
	ret := C.XF86VidModeGetGammaRamp((*C.Display)(xDisplay), C.int(screen), 256, (*C.ushort)(unsafe.Pointer(&r)), (*C.ushort)(unsafe.Pointer(&g)), (*C.ushort)(unsafe.Pointer(&b)))
	if ret == C.False {
		return red, green, blue, errors.New("Unable to get gamma ramp; XF86VidModeGetGammaRamp() failed")
	}

	// Red
	for i := 0; i < 256; i++ {
		red[i] = uint16(r[i])
	}
	// Green
	for i := 0; i < 256; i++ {
		green[i] = uint16(g[i])
	}
	// Blue
	for i := 0; i < 256; i++ {
		blue[i] = uint16(b[i])
	}
	return red, green, blue, nil
}

// Helper to set window states
func c_SetWindowStates(display *c_Display, window c_Window, states []string) {
    netWmState := C.Atom(c_atom(display, "_NET_WM_STATE", false))
    atoms := []c_Atom{}
    for i := 0; i < len(states); i++ {
        state := states[i]
        atoms = append(atoms, c_atom(display, state, false))
    }
    atomType := C.Atom(c_atom(display, "ATOM", false))
    if len(atoms) > 0 {
        C.XChangeProperty(display.C(), C.Window(window), netWmState, atomType, 32, C.PropModePrepend, (*C.uchar)(unsafe.Pointer(&atoms)), C.int(len(atoms)))
    } else {
        C.XDeleteProperty(display.C(), C.Window(window), netWmState)
    }
}

func c_glXSwapBuffers(display *c_Display, window c_Window) {
	C.glXSwapBuffers(display.C(), C.GLXDrawable(window))
}

func c_glXGetVisualFromFBConfig(display *c_Display, config c_GLXFBConfig) *c_XVisualInfo {
	return (*c_XVisualInfo)(C.glXGetVisualFromFBConfig(display.C(), C.GLXFBConfig(config)))
}

func c_glXQueryExtension(dpy *c_Display) (int, int, error) {
	var errorb, event C.int
	ret := C.glXQueryExtension((*C.Display)(dpy), &errorb, &event)
	if ret == 1 {
		return int(errorb), int(event), nil
	}
	return int(errorb), int(event), errors.New("GLX support non existant. glXQueryExtension() failed!")
}

func c_glXQueryVersion(dpy *c_Display) (int, int, error) {
	var major, minor C.int
	if C.glXQueryVersion((*C.Display)(dpy), &major, &minor) == 1 {
		return int(major), int(minor), nil
	}
	return int(major), int(minor), errors.New("Failed to retrieve GLX version. glXQueryVersion() failed!")
}

func glXGetFBConfigAttrib(dpy *c_Display, config C.GLXFBConfig, attribute C.int) int {
	var value C.int
	C.glXGetFBConfigAttrib((*C.Display)(dpy), config, attribute, &value)
	return int(value)
}

func c_getGLXConfigurations(display *c_Display, screen int) ([]*Configuration, error) {
	var nElements C.int
	configs := C.glXGetFBConfigs(display.C(), C.int(screen), &nElements)
    defer C.XFree(unsafe.Pointer(configs))

	// First we get the configs
	fbconfigs := []*Configuration{}
	for i := 0; i < int(nElements); i++ {
		config := C.fbConfigAtIndex(configs, C.int(i))

		fbconfig := Configuration{}
        fbconfig.Renderer = OpenGL
        fbconfig.ColorDepth = uint8(glXGetFBConfigAttrib(display, config, C.GLX_RED_SIZE) +
            glXGetFBConfigAttrib(display, config, C.GLX_GREEN_SIZE) +
            glXGetFBConfigAttrib(display, config, C.GLX_BLUE_SIZE) +
            glXGetFBConfigAttrib(display, config, C.GLX_ALPHA_SIZE))

        fbconfig.AccumDepth = uint8(glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_RED_SIZE) +
            glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_GREEN_SIZE) +
            glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_BLUE_SIZE) +
            glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_ALPHA_SIZE))

		fbconfig.DepthBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_DEPTH_SIZE))
		fbconfig.StencilBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_STENCIL_SIZE))
		fbconfig.Samples = uint8(glXGetFBConfigAttrib(display, config, C.GLX_SAMPLES))
		fbconfig.AuxBuffers = uint8(glXGetFBConfigAttrib(display, config, C.GLX_AUX_BUFFERS))

		if glXGetFBConfigAttrib(display, config, C.GLX_DOUBLEBUFFER) == 1 {
			fbconfig.DoubleBuffered = true
		} else {
			fbconfig.DoubleBuffered = false
		}

		if glXGetFBConfigAttrib(display, config, C.GLX_STEREO) == 1 {
			fbconfig.Stereoscopic = true
		} else {
			fbconfig.Stereoscopic = false
		}

		visual := c_glXGetVisualFromFBConfig(xDisplay, c_GLXFBConfig(config))
		if visual == nil {
			continue // This is an invalid GLX fb config, probably without GL rendering!
		}
        fbconfig.visual = visual

		fbconfigs = append(fbconfigs, &fbconfig)
	}
    return fbconfigs, nil
}

func c_getPixelCopyConfigurations(display *c_Display, screen int) ([]*Configuration, error) {
	configs := []*Configuration{}
    var n C.int
    var template C.XVisualInfo
    vi := C.XGetVisualInfo(display.C(), C.VisualNoMask, &template, &n)
    if vi == nil {
        return configs, errors.New("Unable to get an proper XVisualInfo; XGetVisualInfo() failed!")
    }
    defer C.XFree(unsafe.Pointer(vi))

    for i := 0; i < int(n); i++ {
        element := (*C.XVisualInfo)(c_elementAtIndex(unsafe.Pointer(vi), uintptr(i), unsafe.Sizeof(C.XVisualInfo{})))

        if int(element.screen) != screen {
            continue
        }

        config := &Configuration{}
        config.Renderer = PixelCopy
        //config.ColorDepth = uint8(element.bits_per_rgb) // actually includes alpha, name is just a lier
        config.ColorDepth = uint8(element.depth)
        config.visual = (*c_XVisualInfo)(element)
        //element.visual/element.visualid

        configs = append(configs, config)
    }

    return configs, nil
}

// We will use these xDisplay as our global connection object
// and xDisplayAccess is our lock for that global connection
var xDisplayAccess sync.RWMutex
var xDisplay *c_Display
var xI_opcode C.int

// backend_init and backend_destroy are called from chippy.go
// These two calls are already locked for us, by chippy.go
// So we can safely assume they won't be called at the same time.
func backend_init() error {
	c_XInitThreads()

	var err error
	xDisplay, err = c_XOpenDisplay()
	if err != nil {
		return err
	}


	// Verify we have a working XInput2 extensions
	opcode, _, _, err := c_XQueryExtension(xDisplay, "XInputExtension")
	if err != nil {
		return errors.New("Chippy needs XInput2 extension!") // For raw mouse input, multiple language keyboard codes.. etc
	}
	xI_opcode = C.int(opcode)



	err = initScreens()
	if err != nil {
		return err
	}

	err = initEvents()
	if err != nil {
		return err
	}
	return nil
}

func backend_destroy() {
	destroyScreens()
	destroyEvents()

	xDisplayAccess.RLock()
	defer xDisplayAccess.RUnlock()

	c_XFlush(xDisplay)
	c_XSync(xDisplay, false)
	c_XCloseDisplay(xDisplay)
}*/
