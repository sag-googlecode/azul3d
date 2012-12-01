//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: platform_linux.go
// Created by: Stephen Gutekanst, 11/24/12
//===========================================================================//
//===========================================================================//
// Copyright (c) 2012, Lightpoke
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//     * Redistributions of source code must retain the above copyright
//       notice, n list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, n list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//     * Neither the name of Lightpoke nor the
//       names of its contributors may be used to endorse or promote products
//       derived from n software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL LIGHTPOKE BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF n
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package chippy

// In here we will define all C calls, and make them into nicer Go functions which we will use later
// We prefix all XFunctionCalls to a lower case 'x', XBlahBlah -> xBlahBlah

/*
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/Xatom.h>
#include <X11/extensions/xf86vmode.h>
#include <GL/glx.h>
#cgo LDFLAGS: -lX11 -lXxf86vm -lGL

GLXFBConfig fbConfigAtIndex(GLXFBConfig* configs, int index) {
    return configs[index];
}

XF86VidModeModeInfo* vidModeAtIndex(XF86VidModeModeInfo** modes, int index) {
    return modes[index];
}
*/
import "C"

import "unsafe"
import "errors"
import "sync"

type c_Screen C.Screen
type c_Display C.Display
type c_Atom C.Atom
type c_Window C.Window
type c_Visual C.Visual
type c_Colormap C.Colormap
type c_XSetWindowAttributes C.XSetWindowAttributes
type c_GLXFBConfig C.GLXFBConfig
type c_XVisualInfo C.XVisualInfo
type c_XWindowChanges C.XWindowChanges
type c_XF86VidModeModeInfo C.XF86VidModeModeInfo
type c_Pointer unsafe.Pointer

type c_Hints struct {
    flags, functions, decorations, status uint64
    inputMode int64
}

func (c c_Colormap) C() C.Colormap {
    return C.Colormap(c)
}

var c_AllocNone int32 = C.AllocNone
var c_None C.Pixmap = C.None
var c_StructureNotifyMask C.long = C.StructureNotifyMask
var c_InputOutput uint32 = C.InputOutput
var c_CWBorderPixel uint64 = C.CWBorderPixel
var c_CWColormap uint64 = C.CWColormap
var c_CWEventMask uint64 = C.CWEventMask
var c_PropModeReplace int32 = C.PropModeReplace



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
    ret := C.XCloseDisplay((*C.Display)(display))
    if ret == C.False {
        return errors.New("Call failed XCloseDisplay()")
    }
    return nil
}

// Helper to call XFlush
func c_XFlush(display *c_Display) error {
    ret := C.XFlush((*C.Display)(display))
    if ret == C.False {
        return errors.New("Call failed XFlush()")
    }
    return nil
}

// Helper to call XSync
func c_XSync(display *c_Display, b bool) error {
    var ret C.int
    if b == true {
        ret = C.XSync((*C.Display)(display), C.True)
    } else {
        ret = C.XSync((*C.Display)(display), C.False)
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
    return int32(C.XScreenCount((*C.Display)(display)))
}

func c_XScreenOfDisplay(display *c_Display, screen_number int32) *c_Screen {
    return (*c_Screen)(C.XScreenOfDisplay((*C.Display)(xDisplay), C.int(screen_number)))
}

func c_XDefaultScreenOfDisplay(display *c_Display) *c_Screen {
    return (*c_Screen)(C.XDefaultScreenOfDisplay((*C.Display)(display)))
}

func c_XDefaultRootWindow(display *c_Display) c_Window {
    return (c_Window)(C.XDefaultRootWindow((*C.Display)(display)))
}

func c_XCreateColormap(display *c_Display, w c_Window, visual *c_Visual, alloc int32) c_Colormap {
    return (c_Colormap)(C.XCreateColormap((*C.Display)(display), (C.Window)(w), (*C.Visual)(visual), C.int(alloc)))
}

func c_XCreateWindow(display *c_Display, parent c_Window, x, y int32, width, height, border_width uint32, depth int32, class uint32, visual *c_Visual, valuemask uint64, attributes *c_XSetWindowAttributes) c_Window {
    return (c_Window)(C.XCreateWindow((*C.Display)(display), (C.Window)(parent), C.int(x), C.int(y), C.uint(width), C.uint(height), C.uint(border_width), C.int(depth), C.uint(class), (*C.Visual)(visual), C.ulong(valuemask), (*C.XSetWindowAttributes)(attributes)))
}

func c_XRootWindow(display *c_Display, screen int32) c_Window {
    return (c_Window)(C.XRootWindow((*C.Display)(display), C.int(screen)))
}

func c_XMapWindow(display *c_Display, window c_Window) {
    C.XMapWindow((*C.Display)(display), (C.Window)(window))
}

func c_XStoreName(display *c_Display, window c_Window, title string) {
    cstr := C.CString(title)
    C.XStoreName((*C.Display)(display), (C.Window)(window), cstr)
    //C.free(unsafe.Pointer(&cstr))
}

func c_XIconifyWindow(display *c_Display, window c_Window, screen_number int32) error {
    if C.XIconifyWindow((*C.Display)(display), C.Window(window), C.int(screen_number)) == 0 {
        return errors.New("Unable to iconify window; XIconifyWindow() failed!")
    }
    return nil
}

func c_XRaiseWindow(display *c_Display, window c_Window) error {
    if C.XRaiseWindow((*C.Display)(display), C.Window(window)) == 0 {
        return errors.New("Unable to raise window; XRaiseWindow() failed!")
    }
    return nil
}

func c_XWithdrawWindow(display *c_Display, window c_Window, screen_number int32) error {
    if C.XWithdrawWindow((*C.Display)(display), C.Window(window), C.int(screen_number)) == 0 {
        return errors.New("Unable to iconify window; XIconifyWindow() failed!")
    }
    return nil
}

func c_XFreeColormap(display *c_Display, cmap c_Colormap) {
    C.XFreeColormap((*C.Display)(display), C.Colormap(cmap))
}

func c_XMoveWindow(display *c_Display, w c_Window, x, y int32) {
    C.XMoveWindow((*C.Display)(display), C.Window(w), C.int(x), C.int(y))
}

func c_XResizeWindow(display *c_Display, w c_Window, x, y uint32) {
    C.XResizeWindow((*C.Display)(display), C.Window(w), C.uint(x), C.uint(y))
}

func c_XInternAtom(display *c_Display, atom string, only_if_exists bool) c_Atom {
    cstr := C.CString(atom)
    defer C.free(unsafe.Pointer(cstr))

    v := C.int(C.False)
    if only_if_exists {
        v = C.True
    }
    return (c_Atom)(C.XInternAtom((*C.Display)(display), cstr, v))
}

func c_XChangeProperty(display *c_Display, window c_Window, property, _type c_Atom, format, mode int32, data c_Pointer, nelements int32) {
    rdata := unsafe.Pointer(data)
    C.XChangeProperty((*C.Display)(display), C.Window(window), C.Atom(property), C.Atom(_type), C.int(format), C.int(mode), (*C.uchar)(*&rdata), C.int(nelements))
}


func c_glXSwapBuffers(display *c_Display, window c_Window) {
    C.glXSwapBuffers((*C.Display)(display), C.GLXDrawable(window))
}

func c_glXGetVisualFromFBConfig(display *c_Display, config c_GLXFBConfig) *c_XVisualInfo {
    return (*c_XVisualInfo)(C.glXGetVisualFromFBConfig((*C.Display)(display), C.GLXFBConfig(config)))
}


func c_XF86VidModeGetAllModeLines(display *c_Display, screen int32) []*c_XF86VidModeModeInfo {
    var modecount C.int
    var modelines **C.XF86VidModeModeInfo
    C.XF86VidModeGetAllModeLines((*C.Display)(display), C.int(screen), &modecount, &modelines)
    modes := []*c_XF86VidModeModeInfo{}
    for i := 0; i < int(modecount); i++ {
        mode := C.vidModeAtIndex(modelines, C.int(i))
        modes = append(modes, (*c_XF86VidModeModeInfo)(mode))
    }
    return modes
}

// Helper to call XF86VidModeGetGammaRampSize
func c_XF86VidModeGetGammaRampSize(display *c_Display, screen int32) (int32, error) {
    var size C.int
    ret := C.XF86VidModeGetGammaRampSize((*C.Display)(display), C.int(screen), &size)
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
    worked := C.XF86VidModeSetGammaRamp((*C.Display)(display), C.int(screen), C.int(size), r, g, b)
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

func c_XReconfigureWMWindow(display *c_Display, w c_Window, screen_number int32, value_mask uint32, values *c_XWindowChanges) error {
    ret := C.XReconfigureWMWindow((*C.Display)(display), C.Window(w), C.int(screen_number), C.uint(screen_number), (*C.XWindowChanges)(values))
    if ret == 0 {
        return errors.New("call to XReconfigureWMWindow failed!")
    }
    return nil
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
        C.XChangeProperty((*C.Display)(display), C.Window(window), netWmState, atomType, 32, C.PropModePrepend, (*C.uchar)(unsafe.Pointer(&atoms)), C.int(len(atoms)))
    } else {
        C.XDeleteProperty((*C.Display)(display), C.Window(window), netWmState)
    }
}
*/

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
    configs := C.glXGetFBConfigs((*C.Display)(display), C.int(screen), &nElements)

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

// platformInit and platformDestroy are called from chippy.go
// These two calls are already locked for us, by chippy.go
// So we can safely assume they won't be called at the same time.
func platformInit() error {
    c_XInitThreads()

    var err error
    xDisplay, err = c_XOpenDisplay()
    if err != nil {
        return err
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
    return nil
}

func platformDestroy() {
    xDisplayAccess.RLock()
    defer xDisplayAccess.RUnlock()

    c_XFlush(xDisplay)
    c_XSync(xDisplay, false)
    c_XCloseDisplay(xDisplay)
}

