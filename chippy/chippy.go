// Copyright 2012 Lightpoke. All rights reserved.
// Use of this source code is governed by an BSD
// license found in the License.txt file

// Package chippy implements cross platform window managing.
// Windows Implementation Notes
// 1. Windows 2000 Professional/Server or higher is required.
//    Roughly ever windows API we use requires Windows 2000 Professional/Server or higher.
//
//    I imagine that most people will be at least at this version, and Microsoft probably no longer
//    even supports versions older than that, so I imagine we will not any time in the future
//    either unless someone would like to add and maintain such support as an intirely different
//    backend.
//
// Linux Implementation Notes
//
// 1. Xlib is required.
//    We use this to interface with the X server, at this time there are no plans to use xcb or any
//    other X libraries (xgb, etc), but we're open to suggestions, so if you have valid reasoning,
//    then talk to us about it.
//
// 2. (optional) The X extension Xrandr is highly reccomended
//  Xrandr provides an accurate way to determine an screen's physical size (Xlib is only accurate
//  if screen resolution is at max), also, Xrandr provides an way to change the screen resolution
//  and refresh rate (making the xf86vm extension, below, useless).
//
//  Disable at build time with the build tag: 'no_xrandr'
//
// 3. (optional) The X extension Xxf86vm is used as an fallback if the Xrandr extension (above) is
// unavailable.
//  xf86vm provides an way to change the screen resolution, and refresh rate (but not determine
//  the physical screen size, as in the case of Xrandr above)
//
//  Note: xf86vm appears to sometimes report incorrect display modes, and sometimes refuses to
//  switch to actually available display modes, specifically in cases where Xrandr reports correct
//  display modes, and has no problem switching to the same display mode that xf86vm ignores.
//
//  Disable at build time with the build tag: 'no_xf86vm'
//
// 4. (optional) The X extension Xinput2 is reccomended.
//  Xinput2 provides access to sub-pixel mouse movement, very important for First Person Shooter
//  games, for instance.
//
//  Disable at build time with the build tag: 'no_xinput2'
package chippy

import (
	"io"
	"io/ioutil"
	"log"
	"sync"
)

// Destroy callbacks here, these callbacks are called when the user calls chippy.Destroy()
type callback struct {
	callback func()
}

var destroyCallbacks []*callback

func addDestroyCallback(c *callback) {
	removeDestroyCallback(c) // In case it's already in
	destroyCallbacks = append(destroyCallbacks, c)
}

func removeDestroyCallback(c *callback) {
	for i := 0; i < len(destroyCallbacks); i++ {
		if destroyCallbacks[i] == c {
			// Remove it
			destroyCallbacks = append(destroyCallbacks[:i], destroyCallbacks[i+1:]...)
			break
		}
	}
}

var (
	// Any function calls that have to go back further into C, need to use this global lock
	// basically our reasoning for this is that, majority of the underlying C api's are
	// specifically non-thread safe. So apply this global lock to *most* of the C api we use
	globalLock sync.RWMutex

	// Tells weather chippy has been previously Init()
	isInit bool

	// Tells weather a previous call to Init() failed
	initError error
)

// IsInit returns weather Chippy has been initialized via a previous call to Init().
//
// IsInit() returns false if Destroy() was previously called.
func IsInit() bool {
	globalLock.RLock()
	defer globalLock.RUnlock()

	return isInit
}

// Helper to panic unless previously initialized, use globalLock.Rlock with this!
func panicUnlessInit() {
	if !IsInit() {
		panic("Chippy must be initialized before calling this; Use Init() properly!")
	}
}

var logger *log.Logger

// SetDebugOutput specifies the io.Writer that debug output will be written to (ioutil.Discard by
// default).
func SetDebugOutput(w io.Writer) {
	logger = log.New(w, "chippy: ", log.Ltime|log.Lshortfile)
}

func init() {
	SetDebugOutput(ioutil.Discard)
}

// Init initializes Chippy, returning an error if there is a problem initializing some
// lower level part of Chippy, if an error was returned, it is disallowed to call any
// other Chippy functions. (And any attempt to do so will cause the program to panic.)
func Init() error {
	globalLock.Lock()
	defer globalLock.Unlock()

	if isInit == false {
		// Now we try and initialize the backend, which may fail due to user configurations
		// or something of the sort (dumb user tries to run application on Linux box without
		// any working X11 server or something silly)
		err := backend_Init()
		if err != nil {
			initError = err
			return initError
		}

		// If we made it this far, Chippy should be loaded and ready, and everything is up to
		// the backend to handle things properly now
		isInit = true
		return nil
	}
	return nil
}

// Destroy will destroy Chippy, closing all windows previously opened using NewWindow(), etc.
// Only you know when you're done using Chippy's API, so you should know the appropriate time
// to call this as well. After calling this you are no longer allowed to call any Chippy
// functions.
//
// Typical usage is something like the following:
//
//  err := chippy.Init()
//  if err != nil {
//      handleError(err)
//  }
//  defer chippy.Destroy()
//
// You may call Init() again after calling Destroy() should you want to re-gain access to the API.
func Destroy() {
	globalLock.Lock()
	defer globalLock.Unlock()

	if isInit == true {
		// Firstly, we call each destroy callback, chippyAccess is explicitly unlocked here
		globalLock.Unlock()
		for i := 0; i < len(destroyCallbacks); i++ {
			destroyCallbacks[i].callback()
		}
		globalLock.Lock()
		backend_Destroy()
		isInit = false
		initError = nil
		destroyCallbacks = []*callback{}
	}
}
