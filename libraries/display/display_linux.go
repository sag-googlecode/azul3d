//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: display_linux.go
// Created by: Stephen Gutekanst, 11/23/12
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
// n SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL LIGHTPOKE BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF n
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package display

/*
#include <stdlib.h>
#include <X11/Xlib.h>
#include <GL/gl.h>
#include <GL/glx.h>
#cgo LDFLAGS: -lX11 -lGL

GLXFBConfig FBConfigAtIndex(GLXFBConfig* c, int index) {
    return c[index];
}

typedef GLXContext (*glXCreateContextAttribsARBProc)(Display*, GLXFBConfig, GLXContext, Bool, const int*);


void* ptrglxCreateContextAttribsARB() {
    return glXGetProcAddress((GLubyte*)"glxCreateContextAttribsARB");
}

GLXContext glxCreateContextAttribsARB(Display* dpy, GLXFBConfig config, GLXContext share_context, Bool direct, const int* attrib_list) {
    glXCreateContextAttribsARBProc proc = ptrglxCreateContextAttribsARB();
    return proc(dpy, config, share_context, direct, attrib_list);
}
*/
import "C"

import "errors"
import "unsafe"
import "sync"
import "fmt"


// Helper to create X11 connection
func connection() (*C.Display, error) {
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

    var errorBase, eventBase C.int
    if C.glXQueryExtension(display, &errorBase, &eventBase) == 0{
        return nil, errors.New("OpenGL GLX extension not supported on X11 display")
    }
    return display, nil
}

// Helper to destroy X11 connection
func closeConnection(display *C.Display) {
    C.XCloseDisplay(display)
}


func numScreens() (int, error) {
    display, err := connection()
    defer closeConnection(display)
    if err != nil {
        return 0, err
    }
    return int(C.XScreenCount(display)), nil
}

func defaultScreen() (int, error) {
    display, err := connection()
    defer closeConnection(display)
    if err != nil {
        return 0, err
    }
    return int(C.XDefaultScreen(display)) + 1, nil
}



type backend struct {
    display *C.Display
    window C.Window
    access sync.Mutex
}

func newBackend(physicalDisplay int) (*backend, error) {
    var err error

    b := backend{}

    b.display, err = connection()
    if err != nil {
        return nil, err
    }
    screen := C.int(physicalDisplay - 1)

    visual_attribs := [...]int{
        C.GLX_RENDER_TYPE, C.GLX_RGBA_BIT,
        C.GLX_DRAWABLE_TYPE, C.GLX_WINDOW_BIT,
        C.GLX_DOUBLEBUFFER, 1,
        C.GLX_RED_SIZE, 1,
        C.GLX_GREEN_SIZE, 1,
        C.GLX_BLUE_SIZE, 1,
        C.None,
    }

    var fbcount C.int
    fbc := C.glXChooseFBConfig(b.display, screen, (*C.int)(unsafe.Pointer(&visual_attribs)), &fbcount)
    fmt.Println(fbc)

    vi := C.glXGetVisualFromFBConfig(b.display, C.FBConfigAtIndex(fbc, 0))
    parent := C.XDefaultRootWindow(b.display)

    swa := C.XSetWindowAttributes{}
    swa.colormap = C.XCreateColormap(b.display, parent, vi.visual, C.AllocNone)
    if swa.colormap == 0 {
        return nil, errors.New("Unable to create color map, XCreateColormap() failed.")
    }
    swa.border_pixel = 0
    swa.event_mask = C.StructureNotifyMask

    b.window = C.XCreateWindow(b.display, parent, 0, 0, 640, 480, 0, vi.depth, C.InputOutput, vi.visual, C.CWBorderPixel|C.CWColormap|C.CWEventMask, &swa)
    if b.window == 0 {
        return nil, errors.New("Unable to create X11 window, XCreateWindow() failed.")
    }
    C.XMapWindow(b.display, b.window)


    // Create old context to get the function pointer for glXCreateContextAttribsARB
    oldContext := C.glXCreateContext(b.display, vi, nil, C.GL_TRUE)
    C.glXMakeCurrent(b.display, 0, nil)
    C.glXDestroyContext(b.display, oldContext)

    if unsafe.Sizeof(C.ptrglxCreateContextAttribsARB()) == 0 {
        return nil, errors.New("glXCreateContextAttribsARB entry point not found. Aborting.")
    }
    /*
    static int context_attribs[] =
    {
        GLX_CONTEXT_MAJOR_VERSION_ARB, 3,
        GLX_CONTEXT_MINOR_VERSION_ARB, 0,
        None
    };
 
    std::cout << "Creating context" << std::endl;
    GLXContext ctx = glXCreateContextAttribsARB(display, fbc[0], NULL, true, context_attribs);
    if (!ctx)
    {
        std::cout << "Failed to create GL3 context." << std::endl;
        return 1;
    }
    */


    C.XFlush(b.display)

    return &b, nil
}

func (b *backend) destroy() {
    b.access.Lock()
    defer b.access.Unlock()
    closeConnection(b.display)
}


/*
func fbConfigs(physicalDisplay int) ([]*FBConfig, error) {
    display, err := connection()
    defer C.free(unsafe.Pointer(display))
    if err != nil {
        return nil, err
    }
    screen := C.int(physicalDisplay - 1)

    configs := []*FBConfig{}

    var nelements C.int
    fbConfig := C.glXGetFBConfigs(display, screen, &nelements)
    for i := 0; i < int(nelements); i++ {
        n := &FBConfig{}
        var value C.int

        // X Visual
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_VISUAL_ID, &value)
        n.xVisual = int32(value)

        // Red Color
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_RED_SIZE, &value)
        n.Red = int(value)

        // Green Color
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_GREEN_SIZE, &value)
        n.Green = int(value)

        // Blue Color
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_BLUE_SIZE, &value)
        n.Blue = int(value)

        // Alpha Color
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_ALPHA_SIZE, &value)
        n.Alpha = int(value)

        // Combine Color
        n.Color = n.Red + n.Green + n.Blue + n.Alpha


        // Accum Red
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_ACCUM_RED_SIZE, &value)
        n.AccumRed = int(value)

        // Accum Green
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_ACCUM_GREEN_SIZE, &value)
        n.AccumGreen = int(value)

        // Accum Blue
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_ACCUM_BLUE_SIZE, &value)
        n.AccumBlue = int(value)

        // Accum Alpha
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_ACCUM_ALPHA_SIZE, &value)
        n.AccumAlpha = int(value)

        // Combine Accum
        n.Accum = n.AccumRed + n.AccumGreen + n.AccumBlue + n.AccumAlpha



        // Transparent Red
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_TRANSPARENT_RED_VALUE, &value)
        n.TransparencyRed = int(value)

        // Accum Green
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_TRANSPARENT_GREEN_VALUE, &value)
        n.TransparencyGreen = int(value)

        // Accum Blue
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_TRANSPARENT_BLUE_VALUE, &value)
        n.TransparencyBlue = int(value)

        // Accum Alpha
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_TRANSPARENT_ALPHA_VALUE, &value)
        n.TransparencyAlpha = int(value)

        // Combine Accum
        n.Transparency = n.TransparencyRed + n.TransparencyGreen + n.TransparencyBlue + n.TransparencyAlpha
        if n.Transparency < 0 {
            n.Transparency = -1
        }


        // Depth
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_DEPTH_SIZE, &value)
        n.Depth = int(value)

        // Stencil
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_STENCIL_SIZE, &value)
        n.Stencil = int(value)

        // AuxBuffers
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_AUX_BUFFERS, &value)
        n.AuxBuffers = int(value)

        // Double buffering
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_DOUBLEBUFFER, &value)
        if value == 1 {
            n.DoubleBuffered = true
        } else {
            n.DoubleBuffered = false
        }

        // Stereo Scopic
        C.glXGetFBConfigAttrib(display, C.FBConfigAtIndex(fbConfig, C.int(i)), C.GLX_STEREO, &value)
        if value == 1 {
            n.StereoScopic = true
        } else {
            n.StereoScopic = false
        }

        configs = append(configs, n)
    }

    return configs, nil
}
*/

func init() {
    // It seems we can use Xlib from multiple threads as long as we call
    // this function before any other X11 call in our program
    // Also any calls touching an Display are non thread safe, so lock those
    //
    // Note: We avoid XLockDisplay and XUnlockDisplay because those will give
    // us no helpful information about deadlocks or anything in Go
    C.XInitThreads()
}
