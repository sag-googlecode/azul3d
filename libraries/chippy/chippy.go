//===========================================================================//
//============ Chippy, a cross platform windowing library in Go! ============//
//===========================================================================//
// File: chippy.go
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

//  ______________________________________________
// |==============================================|
// |                FEATURE SUPPORT               |
// |======================|=========|=====|=======|
// | Core Features        | Windows | Mac | Linux |
// |======================|=========|=====|=======|
// | OpenGL Context       | No      | No  | No    |
// |                      |         |     |       |
// |======================|=========|=====|=======|
// | Monitor Features     | Windows | Mac | Linux |
// |======================|=========|=====|=======|
// | Multiple Monitors    | No      | No  | Yes   |
// | Gamma Corrections    | No      | No  | Yes   |
// | Multiple Windows     | No      | No  | Yes   |
// |                      |         |     |       |
// |======================|=========|=====|=======|
// | Window Features      | Windows | Mac | Linux |
// |======================|=========|=====|=======|
// | Specify Screen       | No      | No  | Yes   |
// | Change Vertical Sync | No      | No  | No    |
// | Change Decorated     | No      | No  | No    |
// | Change Visibility    | No      | No  | Yes   |
// | Change Minimized     | No      | No  | Yes   |
// | Change Fullscreen    | No      | No  | No    |
// | Change Title         | No      | No  | Yes   |
// | Change Width/Height  | No      | No  | No    |
// | Change Position      | No      | No  | No    |
// |______________________|_________|_____|_______|
//

import "errors"
import "sync"
import "fmt"

type Callback struct{
    Callback func()
}

var chippyAccess sync.Mutex
var destroyCallbacks []*Callback
var linuxDisplayName string
var isInit bool
var initError error

// Helper to add a destroy callback
func addDestroyCallback(c *Callback) {
    destroyCallbacks = append(destroyCallbacks, c)
}

// Helper to ensure a destroy callback is non-registered. (removed)
func removeDestroyCallback(c *Callback) {
    found := false
    var i int
    for i = 0; i < len(destroyCallbacks); i++ {
        if destroyCallbacks[i] == c {
            found = true
            break
        }
    }
    if found == false {
        return // It was never in the destroyCallbacks array, so just continue happily
    }
    destroyCallbacks = append(destroyCallbacks[:i], destroyCallbacks[i+1:]...)
}

// Helper to get intialization
func getInitError() error {
    if isInit == false {
        return errors.New("Chippy is not initialized yet!")
    }
    return initError
}

// IsInit returns weather chippy has been initialized via a previous call
// to Init()
// Since calling Destroy() de-initializes chippy, after calling Destroy()
// IsInit() will return false
// This call is thread safe
func IsInit() bool {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return isInit
}

// Initialize chippy, return an error should we be unable to initialize
// This call is thread safe
func Init() error {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    if isInit == false {
        err := platformInit()
        if err != nil {
            initError = err
            return err
        }
        isInit = true
        return nil
    }
    return nil
}

// Destroy chippy, cleaning up anything previously created. Call this
// with defer, only once you're absolutely done with using chippy.
// After calling Destroy() you may call Init() again should you want.
// This call is thread safe
func Destroy() {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    if isInit == true {
        chippyAccess.Unlock()
        for i := 0; i < len(destroyCallbacks); i++ {
            fmt.Println("Calling", destroyCallbacks[i])
            destroyCallbacks[i].Callback()
        }
        chippyAccess.Lock()
        platformDestroy()

        linuxDisplayName = ""
        isInit = false
        initError = nil
        destroyCallbacks = []*Callback{}
    }
}


// SetLinuxDisplayName sets the display_name string that will be passed into
// XOpenDisplay (See http://tronche.com/gui/x/xlib/display/opening.html)
//
// "Specifies the hardware display name, which determines the display and
// communications domain to be used. On a POSIX-conformant system, if the
// display_name is NULL, it defaults to the value of the DISPLAY environment
// variable."
// This call is thread safe
func SetLinuxDisplayName(display_name string) error {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return err
    }

    linuxDisplayName = display_name
    return nil
}

// LinuxDisplayName returns the string previously set by SetLinuxDisplayName
// This call is thread safe
func LinuxDisplayName() (string, error) {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return "", err
    }

    return linuxDisplayName, nil
}

