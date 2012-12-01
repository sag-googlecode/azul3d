//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: clipboard_linux.go
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

//import "fmt"
//import "unsafe"

func setClipboard(contents string) error {
    return nil
    /*
    display, err := xOpenDisplay()
    if err != nil {
        return err
    }
    defer xCloseDisplay(display)
    window := C.XDefaultRootWindow(display)

    clipboard := atom(display, "CLIPBOARD", true)

    cstr := C.CString(contents)
    defer C.free(unsafe.Pointer(cstr))
    C.XChangeProperty(display, window, clipboard, C.XA_ATOM, 8, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&cstr)), C.int(len(contents)))

    C.XSetSelectionOwner(display, clipboard, window, C.CurrentTime)
    */
}

func clipboard() (string, error) {
    return "", nil
    /*
    display, err := xOpenDisplay()
    if err != nil {
        return "", err
    }
    defer xCloseDisplay(display)
    //window := C.XDefaultRootWindow(display)

    clipboard := atom(display, "CLIPBOARD", true)
    fmt.Println(clipboard)
    fmt.Println(C.GoString(C.XGetAtomName(display, clipboard)))
    return "a", nil
    */

    /*
	XGetWindowProperty(dpy,
			   win,
			   pty,
			   0,
			   0,
			   False,
			   AnyPropertyType, type, &pty_format, &pty_items, &pty_size, &buffer);
	XFree(buffer);



	     send data all at once (not using INCR) 
	    XChangeProperty(dpy,
			    *win,
			    *pty,
			    XA_ATOM,
			    32, PropModeReplace, (unsigned char *) types,
			    (int) (sizeof(types) / sizeof(Atom))
		);
	}
    */
}

