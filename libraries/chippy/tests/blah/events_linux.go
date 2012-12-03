//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: events_linux.go
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

/*
#cgo LDFLAGS: -lX11
#include "includes_linux.h"

int do_select(int fd) {
    fd_set fds;

    FD_ZERO(&fds);
    FD_SET(fd, &fds);

    return select(fd + 1, &fds, NULL, NULL, NULL);
}
import "C"

import "fmt"

var looping bool

func eventLoop() {
    for looping {
        xDisplayAccess.RLock()

        fd := C.XConnectionNumber(xDisplay)
        C.XFlush(xDisplay);

        if C.do_select(fd) > 0 {
            var event C.XEvent

            for C.XCheckMaskEvent(xDisplay, 0, &event) == 1 || C.XCheckTypedEvent(xDisplay, C.ClientMessage, &event) == 1 {
                //processEvent(&event);
                fmt.Println(event)
            }


            /*
            // Check whether the cursor has moved inside an active window that has
            // captured the cursor (because then it needs to be re-centered)

            _GLFWwindow* window;
            window = _glfwLibrary.activeWindow;
            if (window)
            {
                if (window->cursorMode == GLFW_CURSOR_CAPTURED &&
                    !window->X11.cursorCentered)
                {
                    _glfwPlatformSetCursorPos(window,
                                              window->width / 2,
                                              window->height / 2);
                    window->X11.cursorCentered = GL_TRUE;

                    // NOTE: This is a temporary fix.  It works as long as you use
                    //       offsets accumulated over the course of a frame, instead of
                    //       performing the necessary actions per callback call.
                    XFlush( _glfwLibrary.X11.display );
                }
            }
            


        xDisplayAccess.RUnlock()
        }
    }
}

func startEventLoop() {
    looping = true
    go eventLoop()
}

func stopEventLoop() {
    looping = false
}

*/

