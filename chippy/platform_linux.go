package chippy

// In here we will define all C calls, and make them into nicer Go functions which we will use later
// We prefix all XFunctionCall's to a c_XFunctionCall


/*
#include <stdlib.h>
#include <string.h>
#include <X11/Xlib.h>
#include <X11/Xatom.h>

// For getting video modes etc
#include <X11/extensions/xf86vmode.h>

// For reading raw mouse input and 32 bit keycodes
#include <X11/extensions/XInput2.h>

// For X11 OpenGL support
#include <GL/glx.h>

#cgo LDFLAGS: -lX11 -lXxf86vm -lGL -lXi

GLXFBConfig fbConfigAtIndex(GLXFBConfig* configs, int index) {
    return configs[index];
}

XF86VidModeModeInfo* vidModeAtIndex(XF86VidModeModeInfo** modes, int index) {
    return modes[index];
}

void setAtomState(Display* display, Window win, const char* atomName, int state) {
    Atom netWmState = XInternAtom(display, "_NET_WM_STATE", True);
    Atom atomToSet = XInternAtom(display, atomName, True);

    XClientMessageEvent xMessage;
    xMessage.type = ClientMessage;
    xMessage.serial = 0;
    xMessage.send_event = True;
    xMessage.window = win;
    xMessage.format = 32;
    xMessage.message_type = netWmState;
    xMessage.data.l[0] = state;
    xMessage.data.l[1] = atomToSet;
    xMessage.data.l[2] = 0;
	XSendEvent(display, DefaultRootWindow(display), False, SubstructureRedirectMask | SubstructureNotifyMask, (XEvent*)&xMessage);
}

int isDeleteWindowAtom(Display* display, XClientMessageEvent e) {
    Atom wmDeleteWindow = XInternAtom(display, "WM_DELETE_WINDOW", False);
    return e.data.l[0] == wmDeleteWindow;
}

int waitForEvent(Display* display, int timeout_sec, int timeout_usec) {
    int fd = ConnectionNumber(display);

    struct timeval timeout;
	timeout.tv_sec = timeout_sec;
	timeout.tv_usec = timeout_usec;

    fd_set fds;
    FD_ZERO(&fds);
    FD_SET(fd, &fds);

    return select(fd + 1, &fds, NULL, NULL, &timeout) > 0;
}


// Some helpers to get XEvent union members..
int XEvent_type(XEvent event) {
    return event.type;
}

XCrossingEvent XEvent_xcrossing(XEvent event) {
    return event.xcrossing;
}

XMotionEvent XEvent_xmotion(XEvent event) {
    return event.xmotion;
}

XConfigureEvent XEvent_xconfigure(XEvent event) {
    return event.xconfigure;
}

XPropertyEvent XEvent_xproperty(XEvent event) {
    return event.xproperty;
}

XFocusChangeEvent XEvent_xfocus(XEvent event) {
    return event.xfocus;
}

XMapEvent XEvent_xmap(XEvent event) {
    return event.xmap;
}

XUnmapEvent XEvent_xunmap(XEvent event) {
    return event.xunmap;
}

XExposeEvent XEvent_xexpose(XEvent event) {
    return event.xexpose;
}

XClientMessageEvent XEvent_xclient(XEvent event) {
    return event.xclient;
}

XGenericEventCookie* XGenericEventCookie_xcookie(XEvent event) {
    XGenericEventCookie* ptr = &event.xcookie;
    return ptr;
}

XIDeviceEvent* XGenericEventCookie_XIDeviceEventPtr(XGenericEventCookie* cookie) {
    return (XIDeviceEvent*)(cookie->data);
}


long propertyAtIndex(unsigned char *properties, int index) {
    return ((long*)(properties))[index];
}

void setXIEventMask(Display* display) {
    XIEventMask mask;
    mask.mask_len = XIMaskLen(XI_RawMotion);
    mask.mask = calloc(mask.mask_len, sizeof(char));
    mask.deviceid = XIAllMasterDevices;
    memset(mask.mask, 0, mask.mask_len);

    //XISetMask(mask.mask, XI_KeyPress);
    //XISetMask(mask.mask, XI_KeyRelease);
    //XISetMask(mask.mask, XI_ButtonPress);
    //XISetMask(mask.mask, XI_ButtonRelease);
    //XISetMask(mask.mask, XI_Motion);
    //XISetMask(mask.mask, XI_Enter);
    //XISetMask(mask.mask, XI_Leave);
    //XISetMask(mask.mask, XI_FocusIn);
    //XISetMask(mask.mask, XI_FocusOut);
    XISetMask(mask.mask, XI_RawMotion);

    XISelectEvents(display, DefaultRootWindow(display), &mask, 1);
    free(mask.mask);
}

*/
import "C"

import "unsafe"
import "errors"
import "sync"

type c_Screen C.Screen
type c_Display C.Display

func (d *c_Display) C() *C.Display {
    return (*C.Display)(d)
}

type c_Atom C.Atom
type c_Window C.Window
type c_Visual C.Visual
type c_Colormap C.Colormap
type c_Pixmap C.Pixmap
type c_XWindowAttributes C.XWindowAttributes
type c_XSetWindowAttributes C.XSetWindowAttributes
type c_XVisualInfo C.XVisualInfo
type c_XWindowChanges C.XWindowChanges
type c_XClientMessageEvent C.XClientMessageEvent
type c_XEvent C.XEvent
type c_XCrossingEvent C.XCrossingEvent
type c_XMotionEvent C.XMotionEvent
type c_XConfigureEvent C.XConfigureEvent
type c_XPropertyEvent C.XPropertyEvent
type c_XFocusChangeEvent C.XFocusChangeEvent
type c_XMapEvent C.XMapEvent
type c_XUnmapEvent C.XUnmapEvent
type c_XExposeEvent C.XExposeEvent
type c_XGenericEventCookie C.XGenericEventCookie


type c_XIDeviceEvent C.XIDeviceEvent

/*
type c_XAnyEvent C.XAnyEvent
type c_XKeyEvent C.XKeyEvent
type c_XButtonEvent C.XButtonEvent
type c_XMotionEvent C.XMotionEvent
type c_XCrossingEvent C.XCrossingEvent
type c_XFocusChangeEvent C.XFocusChangeEvent
type c_XExposeEvent C.XExposeEvent
type c_XGraphicsExposeEvent C.XGraphicsExposeEvent
type c_XNoExposeEvent C.XNoExposeEvent
type c_XVisibilityEvent C.XVisibilityEvent
type c_XCreateWindowEvent C.XCreateWindowEvent
type c_XDestroyWindowEvent C.XDestroyWindowEvent
type c_XUnmapEvent C.XUnmapEvent
type c_XMapEvent C.XMapEvent
type c_XMapRequestEvent C.XMapRequestEvent
type c_XReparentEvent C.XReparentEvent
type c_XConfigureEvent C.XConfigureEvent
type c_XGravityEvent C.XGravityEvent
type c_XResizeRequestEvent C.XResizeRequestEvent
type c_XConfigureRequestEvent C.XConfigureRequestEvent
type c_XCirculateEvent C.XCirculateEvent
type c_XCirculateRequestEvent C.XCirculateRequestEvent
type c_XPropertyEvent C.XPropertyEvent
type c_XSelectionClearEvent C.XSelectionClearEvent
type c_XSelectionRequestEvent C.XSelectionRequestEvent
type c_XSelectionEvent C.XSelectionEvent
type c_XColormapEvent C.XColormapEvent
type c_XClientMessageEvent C.XClientMessageEvent
type c_XMappingEvent C.XMappingEvent
type c_XErrorEvent C.XErrorEvent
type c_XKeymapEvent C.XKeymapEvent
type c_XGenericEvent C.XGenericEvent
type c_XGenericEventCookie C.XGenericEventCookie
*/

type c_XSizeHints C.XSizeHints
type c_XWMHints C.XWMHints
type c_XClassHint C.XClassHint

type c_GLXFBConfig C.GLXFBConfig
type c_XF86VidModeModeInfo C.XF86VidModeModeInfo
type c_Pointer unsafe.Pointer

type c_Hints struct {
    flags, functions, decorations, status uint64
    inputMode int64
}

func (c *c_XGenericEventCookie) C() *C.XGenericEventCookie {
    return (*C.XGenericEventCookie)(unsafe.Pointer(c))
}

func (c *c_XGenericEventCookie) XIDeviceEvent() *c_XIDeviceEvent {
    return (*c_XIDeviceEvent)(C.XGenericEventCookie_XIDeviceEventPtr(c.C()))
}


func (c c_Colormap) C() C.Colormap {
    return C.Colormap(c)
}

func (c c_Pixmap) C() C.Pixmap {
    return C.Pixmap(c)
}

func (c c_Window) C() C.Window {
    return C.Window(c)
}

func (c c_Atom) C() C.Atom {
    return C.Atom(c)
}

// Some helpers to get XEvent union members..
func (c c_XEvent) C() C.XEvent {
    return C.XEvent(c)
}

func (e c_XEvent) _type() int32 {
    return int32(C.XEvent_type(e.C()))
}

func (e c_XEvent) xcrossing() c_XCrossingEvent {
    return (c_XCrossingEvent)(C.XEvent_xcrossing(e.C()))
}

func (e c_XEvent) xmotion() c_XMotionEvent {
    return (c_XMotionEvent)(C.XEvent_xmotion(e.C()))
}

func (e c_XEvent) xconfigure() c_XConfigureEvent {
    return (c_XConfigureEvent)(C.XEvent_xconfigure(e.C()))
}

func (e c_XEvent) xproperty() c_XPropertyEvent {
    return (c_XPropertyEvent)(C.XEvent_xproperty(e.C()))
}

func (e c_XEvent) xfocus() c_XFocusChangeEvent {
    return (c_XFocusChangeEvent)(C.XEvent_xfocus(e.C()))
}

func (e c_XEvent) xmap() c_XMapEvent {
    return (c_XMapEvent)(C.XEvent_xmap(e.C()))
}

func (e c_XEvent) xunmap() c_XUnmapEvent {
    return (c_XUnmapEvent)(C.XEvent_xunmap(e.C()))
}

func (e c_XEvent) xexpose() c_XExposeEvent {
    return (c_XExposeEvent)(C.XEvent_xexpose(e.C()))
}

func (e c_XEvent) xclient() c_XClientMessageEvent {
    return (c_XClientMessageEvent)(C.XEvent_xclient(e.C()))
}

func (e c_XEvent) xcookie() *c_XGenericEventCookie {
    return (*c_XGenericEventCookie)(C.XGenericEventCookie_xcookie(e.C()))
}


func (c c_XClientMessageEvent) C() C.XClientMessageEvent {
    return C.XClientMessageEvent(c)
}

func (c c_XClientMessageEvent) asXEvent() *c_XEvent {
    return (*c_XEvent)(unsafe.Pointer(&c))
}

var c_AllocNone int32 = C.AllocNone
//var c_None C.Pixmap = C.None
var c_None = C.None

var c_NoEventMask int64 = C.NoEventMask
var c_KeyPressMask int64 = C.KeyPressMask
var c_KeyReleaseMask int64 = C.KeyReleaseMask
var c_ButtonPressMask int64 = C.ButtonPressMask
var c_ButtonReleaseMask int64 = C.ButtonReleaseMask
var c_EnterWindowMask int64 = C.EnterWindowMask
var c_LeaveWindowMask int64 = C.LeaveWindowMask
var c_PointerMotionMask int64 = C.PointerMotionMask
var c_PointerMotionHintMask int64 = C.PointerMotionHintMask
var c_Button1MotionMask int64 = C.Button1MotionMask
var c_Button2MotionMask int64 = C.Button2MotionMask
var c_Button3MotionMask int64 = C.Button3MotionMask
var c_Button4MotionMask int64 = C.Button4MotionMask
var c_Button5MotionMask int64 = C.Button5MotionMask
var c_ButtonMotionMask int64 = C.ButtonMotionMask
var c_KeymapStateMask int64 = C.KeymapStateMask
var c_ExposureMask int64 = C.ExposureMask
var c_VisibilityChangeMask int64 = C.VisibilityChangeMask
var c_StructureNotifyMask int64 = C.StructureNotifyMask
var c_ResizeRedirectMask int64 = C.ResizeRedirectMask
var c_SubstructureNotifyMask int64 = C.SubstructureNotifyMask
var c_SubstructureRedirectMask int64 = C.SubstructureRedirectMask
var c_FocusChangeMask int64 = C.FocusChangeMask
var c_PropertyChangeMask int64 = C.PropertyChangeMask
var c_ColormapChangeMask int64 = C.ColormapChangeMask
var c_OwnerGrabButtonMask int64 = C.OwnerGrabButtonMask


var c_KeyPress int32 = C.KeyPress
var c_KeyRelease int32 = C.KeyRelease
var c_ButtonPress int32 = C.ButtonPress
var c_ButtonRelease int32 = C.ButtonRelease
var c_MotionNotify int32 = C.MotionNotify
var c_EnterNotify int32 = C.EnterNotify
var c_LeaveNotify int32 = C.LeaveNotify
var c_FocusIn int32 = C.FocusIn
var c_FocusOut int32 = C.FocusOut
var c_KeymapNotify int32 = C.KeymapNotify
var c_Expose int32 = C.Expose
var c_GraphicsExpose int32 = C.GraphicsExpose
var c_NoExpose int32 = C.NoExpose
var c_VisibilityNotify int32 = C.VisibilityNotify
var c_CreateNotify int32 = C.CreateNotify
var c_DestroyNotify int32 = C.DestroyNotify
var c_UnmapNotify int32 = C.UnmapNotify
var c_MapNotify int32 = C.MapNotify
var c_MapRequest int32 = C.MapRequest
var c_ReparentNotify int32 = C.ReparentNotify
var c_ConfigureNotify int32 = C.ConfigureNotify
var c_ConfigureRequest int32 = C.ConfigureRequest
var c_GravityNotify int32 = C.GravityNotify
var c_ResizeRequest int32 = C.ResizeRequest
var c_CirculateNotify int32 = C.CirculateNotify
var c_CirculateRequest int32 = C.CirculateRequest
var c_PropertyNotify int32 = C.PropertyNotify
var c_SelectionClear int32 = C.SelectionClear
var c_SelectionRequest int32 = C.SelectionRequest
var c_SelectionNotify int32 = C.SelectionNotify
var c_ColormapNotify int32 = C.ColormapNotify
var c_ClientMessage int32 = C.ClientMessage
var c_MappingNotify int32 = C.MappingNotify
var c_GenericEvent int32 = C.GenericEvent


var c_XI_RawMotion int32 = C.XI_RawMotion



var c_InputOutput uint32 = C.InputOutput
var c_CWBorderPixel uint64 = C.CWBorderPixel
var c_CWColormap uint64 = C.CWColormap
var c_CWEventMask uint64 = C.CWEventMask
var c_PropModeReplace int32 = C.PropModeReplace

func c_int(i int32) C.int {
    return C.int(i)
}

func c_bool(b bool) C.int {
    if b{
        return C.True
    }
    return C.False
}

func c_long(c int64) C.long {
    return C.long(c)
}


// Helper to call XInitThreads
func c_XInitThreads() {
    C.XInitThreads()
}

// Helper to open an X display
func c_XOpenDisplay() (*c_Display, error) {
    var display *C.Display

    if len(linuxDisplayName) > 0 {
        cstr := C.CString(linuxDisplayName)
        defer C.free(unsafe.Pointer(cstr))
        display = C.XOpenDisplay(cstr)
    } else {
        display = C.XOpenDisplay(nil)
    }

    if display == nil {
        return nil, errors.New("Unable to open an X11 connection! XOpenDisplay() failed.")
    }
    return (*c_Display)(display), nil
}

// Helper to close an X display
func c_XCloseDisplay(display *c_Display) error {
    ret := C.XCloseDisplay(display.C())
    if ret == C.False {
        return errors.New("Call failed XCloseDisplay()")
    }
    return nil
}

// Helper to call XFlush
func c_XFlush(display *c_Display) error {
    ret := C.XFlush(display.C())
    if ret == C.False {
        return errors.New("Call failed XFlush()")
    }
    return nil
}

// Helper to call XSync
func c_XSync(display *c_Display, b bool) error {
    var ret C.int
    if b == true {
        ret = C.XSync(display.C(), C.True)
    } else {
        ret = C.XSync(display.C(), C.False)
    }
    if ret == C.False {
        return errors.New("Call failed XSync()")
    }
    return nil
}

func c_XScreenNumberOfScreen(screen *c_Screen) int32 {
    return int32(C.XScreenNumberOfScreen((*C.Screen)(screen)))
}

func c_XWidthOfScreen(screen *c_Screen) int32 {
    return int32(C.XWidthOfScreen((*C.Screen)(screen)))
}

func c_XHeightOfScreen(screen *c_Screen) int32 {
    return int32(C.XHeightOfScreen((*C.Screen)(screen)))
}

func c_XScreenCount(display *c_Display) int32 {
    return int32(C.XScreenCount(display.C()))
}

func c_XScreenOfDisplay(display *c_Display, screen_number int32) *c_Screen {
    return (*c_Screen)(C.XScreenOfDisplay((*C.Display)(xDisplay), C.int(screen_number)))
}

func c_XDefaultScreenOfDisplay(display *c_Display) *c_Screen {
    return (*c_Screen)(C.XDefaultScreenOfDisplay(display.C()))
}

func c_XDefaultRootWindow(display *c_Display) c_Window {
    return (c_Window)(C.XDefaultRootWindow(display.C()))
}

func c_XCreateColormap(display *c_Display, w c_Window, visual *c_Visual, alloc int32) c_Colormap {
    return (c_Colormap)(C.XCreateColormap(display.C(), (C.Window)(w), (*C.Visual)(visual), C.int(alloc)))
}

func c_XCreateWindow(display *c_Display, parent c_Window, x, y int32, width, height, border_width uint32, depth int32, class uint32, visual *c_Visual, valuemask uint64, attributes *c_XSetWindowAttributes) c_Window {
    return (c_Window)(C.XCreateWindow(display.C(), (C.Window)(parent), C.int(x), C.int(y), C.uint(width), C.uint(height), C.uint(border_width), C.int(depth), C.uint(class), (*C.Visual)(visual), C.ulong(valuemask), (*C.XSetWindowAttributes)(attributes)))
}

func c_XRootWindow(display *c_Display, screen int32) c_Window {
    return (c_Window)(C.XRootWindow(display.C(), C.int(screen)))
}

func c_XMapWindow(display *c_Display, window c_Window) error {
    if C.XMapWindow(display.C(), (C.Window)(window)) == 0 {
        return errors.New("Unable to map X window; XMapWindow() failed!")
    }
    return nil
}

func c_XStoreName(display *c_Display, window c_Window, title string) error {
    cstr := C.CString(title)
    //C.free(unsafe.Pointer(&cstr))
    if C.XStoreName(display.C(), (C.Window)(window), cstr) == 0 {
        return errors.New("Unable to store name; XStoreName() failed!")
    }
    return nil
}

func c_XIconifyWindow(display *c_Display, window c_Window, screen_number int32) error {
    if C.XIconifyWindow(display.C(), C.Window(window), C.int(screen_number)) == 0 {
        return errors.New("Unable to iconify window; XIconifyWindow() failed!")
    }
    return nil
}

func c_XRaiseWindow(display *c_Display, window c_Window) error {
    if C.XRaiseWindow(display.C(), C.Window(window)) == 0 {
        return errors.New("Unable to raise window; XRaiseWindow() failed!")
    }
    return nil
}

func c_XWithdrawWindow(display *c_Display, window c_Window, screen_number int32) error {
    if C.XWithdrawWindow(display.C(), C.Window(window), C.int(screen_number)) == 0 {
        return errors.New("Unable to iconify window; XIconifyWindow() failed!")
    }
    return nil
}

func c_XFreeColormap(display *c_Display, cmap c_Colormap) {
    C.XFreeColormap(display.C(), C.Colormap(cmap))
}

func c_XMoveWindow(display *c_Display, w c_Window, x, y int32) error {
    if C.XMoveWindow(display.C(), C.Window(w), C.int(x), C.int(y)) == 0 {
        return errors.New("Unable to move window, XMoveWindow() failed!")
    }
    return nil
}

func c_XResizeWindow(display *c_Display, w c_Window, x, y uint32) error {
    if C.XResizeWindow(display.C(), C.Window(w), C.uint(x), C.uint(y)) == 0 {
        return errors.New("Unable to resize window, XResizeWindow() failed!")
    }
    return nil
}

func c_XInternAtom(display *c_Display, atom string, only_if_exists bool) c_Atom {
    cstr := C.CString(atom)
    defer C.free(unsafe.Pointer(cstr))

    v := C.int(C.False)
    if only_if_exists {
        v = C.True
    }
    return (c_Atom)(C.XInternAtom(display.C(), cstr, v))
}

func c_XChangeProperty(display *c_Display, window c_Window, property, _type c_Atom, format, mode int32, data c_Pointer, nelements int32) error {
    rdata := unsafe.Pointer(data)
    if C.XChangeProperty(display.C(), C.Window(window), C.Atom(property), C.Atom(_type), C.int(format), C.int(mode), (*C.uchar)(*&rdata), C.int(nelements)) == 0 {
        return errors.New("Unable to change property; XChangeProperty() failed!")
    }
    return nil
}

func c_XWarpPointer(display *c_Display, src_w c_Window, dest_w c_Window, src_x, src_y int32, src_width, src_height uint32, dest_x, dest_y int32) {
    C.XWarpPointer(display.C(), C.Window(src_w), C.Window(dest_w), C.int(src_x), C.int(src_y), C.uint(src_width), C.uint(src_height), C.int(dest_x), C.int(dest_y))
}

func c_XReconfigureWMWindow(display *c_Display, w c_Window, screen_number int32, value_mask uint32, values *c_XWindowChanges) error {
    ret := C.XReconfigureWMWindow(display.C(), C.Window(w), C.int(screen_number), C.uint(screen_number), (*C.XWindowChanges)(values))
    if ret == 0 {
        return errors.New("call to XReconfigureWMWindow failed!")
    }
    return nil
}

func c_Xutf8SetWMProperties(display *c_Display, w c_Window, window_name, icon_name, argv string, argc int32, normal_hints *c_XSizeHints, wm_hints *c_XWMHints, class_hints *c_XClassHint) {
    c_window_name := C.CString(window_name)
    c_icon_name := C.CString(icon_name)
    c_argv := C.CString(argv)
    C.Xutf8SetWMProperties(display.C(), C.Window(w), c_window_name, c_icon_name, &c_argv, C.int(argc), (*C.XSizeHints)(normal_hints), (*C.XWMHints)(wm_hints), (*C.XClassHint)(class_hints))
}


func c_setAtomState(display *c_Display, w c_Window, atom string, state bool) {
    cstr := C.CString(atom)
    C.setAtomState(display.C(), C.Window(w), cstr, c_bool(state))
}

func c_XSendEvent(display *c_Display, w c_Window, propagate bool, event_mask int64, event_send *c_XEvent) error {
    if C.XSendEvent(display.C(), C.Window(w), c_bool(propagate), C.long(event_mask), (*C.XEvent)(event_send)) == 0 {
        return errors.New("Failed to send X event; XSendEvent() failed!")
    }
    return nil
}

func c_XConnectionNumber(display *c_Display) int32 {
    return int32(C.XConnectionNumber(display.C()))
}

func c_XSelectInput(display *c_Display, w c_Window, mask int64) {
    C.XSelectInput(display.C(), C.Window(w), C.long(mask))
}

func c_XDestroyWindow(display *c_Display, w c_Window) {
    C.XDestroyWindow(display.C(), C.Window(w))
}

func c_XNextEvent(display *c_Display) *c_XEvent {
    var ev C.XEvent
    C.XNextEvent(display.C(), &ev)
    return (*c_XEvent)(&ev)
}

func c_XPending(display *c_Display) int32 {
    return int32(C.XPending(display.C()))
}

func c_XGetWindowAttributes(display *c_Display, window c_Window) (*c_XWindowAttributes, error) {
    var attribs C.XWindowAttributes
    if C.XGetWindowAttributes(display.C(), window.C(), &attribs) == 0 {
        return nil, errors.New("Unable to query window attributes; XGetWindowAttributes() failed!")
    }
    return (*c_XWindowAttributes)(&attribs), nil
}

func c_XSetWMProtocols(display *c_Display, window c_Window, protocols *c_Atom, count int32) error {
    if C.XSetWMProtocols(display.C(), window.C(), (*C.Atom)(protocols), C.int(count)) == 0 {
        return errors.New("Unable to set WM protocols; XSetWMProtocols() failed!")
    }
    return nil
}

func c_XGetAtomName(display *c_Display, atom c_Atom) string {
    return C.GoString(C.XGetAtomName(display.C(), (C.Atom)(atom)))
}

func c_XQueryExtension(display *c_Display, name string) (int32, int32, int32, error) {
    cstr := C.CString(name)
    var major_opcode_return, first_event_return, first_error_return C.int
    if C.XQueryExtension(display.C(), cstr, &major_opcode_return, &first_event_return, &first_error_return) == 0 {
        return int32(major_opcode_return), int32(first_event_return), int32(first_error_return), errors.New("Unable to query X extension, XQueryExtension() failed!")
    }
    return int32(major_opcode_return), int32(first_event_return), int32(first_error_return), nil
}

func c_XGetEventData(display *c_Display, cookie *c_XGenericEventCookie) bool {
    if C.XGetEventData(display.C(), cookie.C()) == 0 {
        return false
    }
    return true
}

func c_XFreeEventData(display *c_Display, cookie *c_XGenericEventCookie) {
    C.XFreeEventData(display.C(), cookie.C())
}

func c_waitForEvent(display *c_Display, sec, usec int32) bool {
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
    */

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
*/



func c_XF86VidModeSetViewPort(display *c_Display, screen, x, y int32) {
    C.XF86VidModeSetViewPort(display.C(), C.int(screen), C.int(x), C.int(y))
}

func c_XF86VidModeGetAllModeLines(display *c_Display, screen int32) []*c_XF86VidModeModeInfo {
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

func c_XF86VidModeSwitchToMode(display *c_Display, screen int32, mi *c_XF86VidModeModeInfo) error {
    if C.XF86VidModeSwitchToMode(display.C(), C.int(screen), (*C.XF86VidModeModeInfo)(mi)) == 0 {
        return errors.New("Unable to switch to video mode; XF86VidModeSwitchToMode() failed!")
    }
    return nil
}

// Helper to call XF86VidModeGetGammaRampSize
func c_XF86VidModeGetGammaRampSize(display *c_Display, screen int32) (int32, error) {
    var size C.int
    ret := C.XF86VidModeGetGammaRampSize(display.C(), C.int(screen), &size)
    if ret == C.False {
        return 0, errors.New("Call failed XF86VidModeGetGammaRampSize()")
    }
    return int32(size), nil
}

func c_XF86VidModeSetGammaRamp(display *c_Display, screen int32, red, green, blue [256]uint16) error {
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

func c_XF86VidModeGetGammaRamp(display *c_Display, screen int32) ([256]uint16, [256]uint16, [256]uint16, error) {
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
/*
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
*/

func c_glXSwapBuffers(display *c_Display, window c_Window) {
    C.glXSwapBuffers(display.C(), C.GLXDrawable(window))
}

func c_glXGetVisualFromFBConfig(display *c_Display, config c_GLXFBConfig) *c_XVisualInfo {
    return (*c_XVisualInfo)(C.glXGetVisualFromFBConfig(display.C(), C.GLXFBConfig(config)))
}


func c_glXQueryExtension(dpy *c_Display) (int32, int32, error) {
    var errorb, event C.int
    ret := C.glXQueryExtension((*C.Display)(dpy), &errorb, &event)
    if ret == 1 {
        return int32(errorb), int32(event), nil
    }
    return int32(errorb), int32(event), errors.New("GLX support non existant. glXQueryExtension() failed!")
}

func c_glXQueryVersion(dpy *c_Display) (int32, int32, error) {
    var major, minor C.int
    if C.glXQueryVersion((*C.Display)(dpy), &major, &minor) == 1 {
        return int32(major), int32(minor), nil
    }
    return int32(major), int32(minor), errors.New("Failed ot retrieve GLX version. glXQueryVersion() failed!")
}

func glXGetFBConfigAttrib(dpy *c_Display, config C.GLXFBConfig, attribute C.int) int32 {
    var value C.int
    C.glXGetFBConfigAttrib((*C.Display)(dpy), config, attribute, &value)
    return int32(value)
}

// Helper to choose the best frame buffer configuration
func c_chooseFBConfig(display *c_Display, screen int32, minAttribs, maxAttribs *FBConfig) *FBConfig {
    var nElements C.int
    configs := C.glXGetFBConfigs(display.C(), C.int(screen), &nElements)

    // First we get the configs
    fbconfigs := []*FBConfig{}
    for i := 0; i < int(nElements); i++ {
        config := C.fbConfigAtIndex(configs, C.int(i))

        fbconfig := FBConfig{}
        fbconfig.RedBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_RED_SIZE))
        fbconfig.GreenBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_GREEN_SIZE))
        fbconfig.BlueBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_BLUE_SIZE))
        fbconfig.AlphaBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_ALPHA_SIZE))

        fbconfig.AccumRedBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_RED_SIZE))
        fbconfig.AccumGreenBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_GREEN_SIZE))
        fbconfig.AccumBlueBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_BLUE_SIZE))
        fbconfig.AccumAlphaBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_ACCUM_ALPHA_SIZE))

        fbconfig.DepthBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_DEPTH_SIZE))
        fbconfig.StencilBits = uint8(glXGetFBConfigAttrib(display, config, C.GLX_STENCIL_SIZE))
        fbconfig.Samples = uint8(glXGetFBConfigAttrib(display, config, C.GLX_SAMPLES))
        fbconfig.SampleBuffers = uint8(glXGetFBConfigAttrib(display, config, C.GLX_SAMPLE_BUFFERS))
        fbconfig.AuxBuffers = uint8(glXGetFBConfigAttrib(display, config, C.GLX_AUX_BUFFERS))

        if glXGetFBConfigAttrib(display, config, C.GLX_DOUBLEBUFFER) == 1 {
            fbconfig.DoubleBuffered = true
        } else {
            fbconfig.DoubleBuffered = false
        }

        if glXGetFBConfigAttrib(display, config, C.GLX_STEREO) == 1 {
            fbconfig.StereoScopic = true
        } else {
            fbconfig.StereoScopic = false
        }

        fbconfig.actual = c_GLXFBConfig(config)

        valid := c_glXGetVisualFromFBConfig(xDisplay, c_GLXFBConfig(config))
        if valid == nil {
            continue // This is an invalid GLX fb config, probably without GL rendering!
        }

        fbconfigs = append(fbconfigs, &fbconfig)
    }

    // Now remove any configs that are over maxAttribs
    nfbconfigs := []*FBConfig{}
    for i := 0; i < len(fbconfigs); i++ {
        config := fbconfigs[i]
        if config.RedBits > maxAttribs.RedBits { continue }
        if config.GreenBits > maxAttribs.GreenBits { continue }
        if config.BlueBits > maxAttribs.BlueBits { continue }
        if config.AlphaBits > maxAttribs.AlphaBits { continue }

        if config.AccumRedBits > maxAttribs.AccumRedBits { continue }
        if config.AccumGreenBits > maxAttribs.AccumGreenBits { continue }
        if config.AccumBlueBits > maxAttribs.AccumBlueBits { continue }
        if config.AccumAlphaBits > maxAttribs.AccumAlphaBits { continue }

        if config.DepthBits > maxAttribs.DepthBits { continue }
        if config.StencilBits > maxAttribs.StencilBits { continue }
        if config.Samples > maxAttribs.Samples { continue }
        if config.SampleBuffers > maxAttribs.SampleBuffers { continue }
        if config.AuxBuffers > maxAttribs.AuxBuffers { continue }

        if config.DoubleBuffered && !maxAttribs.DoubleBuffered { continue }
        if config.StereoScopic && !maxAttribs.StereoScopic { continue }

        nfbconfigs = append(nfbconfigs, config)
    }
    fbconfigs = nfbconfigs

    // Now accumulate the frame buffers
    accumulated := make(map[*FBConfig]int32)
    for i := 0; i < len(fbconfigs); i++ {
        config := fbconfigs[i]
        a := int32(0)
        a += int32(config.RedBits + config.GreenBits + config.BlueBits + config.AlphaBits)
        a += int32(config.AccumRedBits + config.AccumGreenBits + config.AccumBlueBits + config.AccumAlphaBits)
        a += int32(config.DepthBits + config.StencilBits + config.Samples + config.SampleBuffers + config.AuxBuffers)
        if config.DoubleBuffered { a += 1 }
        if config.StereoScopic { a += 1 }
        accumulated[config] = a
    }

    // Now grab the largest (best) one
    var best *FBConfig
    var bestValue int32
    for k, v := range accumulated {
        if v > bestValue {
            bestValue = v
            best = k
        }
    }

    return best
}


// We will use these xDisplay as our global connection object
// and xDisplayAccess is our lock for that global connection
var xDisplayAccess sync.RWMutex
var xDisplay *c_Display
var XI_opcode C.int

// platformInit and platformDestroy are called from chippy.go
// These two calls are already locked for us, by chippy.go
// So we can safely assume they won't be called at the same time.
func initPlatform() error {
    c_XInitThreads()

    var err error
    xDisplay, err = c_XOpenDisplay()
    if err != nil {
        return err
    }

    // Verify we have a working XInput2 extensions
    opcode, _, _, err := c_XQueryExtension(xDisplay, "XInputExtension")
    XI_opcode = C.int(opcode)
    if err != nil {
        return errors.New("Chippy needs XInput2 extension to support raw mouse input!")
    }

    // Verify there is working GLX on xDisplay
    _, _, err = c_glXQueryExtension(xDisplay)
    if err != nil {
        return err
    }

    major, minor, _ := c_glXQueryVersion(xDisplay)
    if major < 1 || minor < 3 {
        return errors.New("Chippy requires GLX version 1.3 or greater")
    }

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

func destroyPlatform() {
    destroyScreens()
    destroyEvents()

    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    c_XFlush(xDisplay)
    c_XSync(xDisplay, false)
    //c_XCloseDisplay(xDisplay)
}

