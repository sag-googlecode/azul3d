// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

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

// Helper to panic unless previously initialized
func panicUnlessInit() {
	globalLock.RLock()
	defer globalLock.RUnlock()

	if !IsInit() {
		panic("Chippy must be initialized before calling this; Use Init() properly!")
	}
}

var theLogger *log.Logger

func logger() *log.Logger {
	globalLock.RLock()
	defer globalLock.RUnlock()

	return theLogger
}

// SetDebugOutput specifies the io.Writer that debug output will be written to (ioutil.Discard by
// default).
func SetDebugOutput(w io.Writer) {
	globalLock.Lock()
	defer globalLock.Unlock()

	theLogger = log.New(w, "chippy: ", log.Ltime|log.Lshortfile)
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
		for _, callback := range destroyCallbacks {
			callback.callback()
		}
		globalLock.Lock()
		backend_Destroy()
		isInit = false
		initError = nil
		destroyCallbacks = []*callback{}
	}
}
