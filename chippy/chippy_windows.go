// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"time"
	"errors"
	"fmt"
	"sync"
)

func eventLoop() {
	for{
		// Small sleep just to stop hogging CPU
		time.Sleep(2 * time.Millisecond)
		dispatch(func() {
			var msg *win32.MSG
			hasMessage := true

			for hasMessage {
				hasMessage, msg = win32.PeekMessage(nil, 0, 0, win32.PM_REMOVE|win32.PM_NOYIELD)
				if hasMessage {
					win32.TranslateMessage(msg)
					win32.DispatchMessage(msg)
				}
			}
		})
	}
}

var classNameCounter = 0
var classNameCounterAccess sync.Mutex

func nextCounter() int {
	classNameCounterAccess.Lock()
	defer classNameCounterAccess.Unlock()
	classNameCounter++
	return classNameCounter
}

//var classAtom win32.ATOM
//var windowClass *win32.WNDCLASSEX
var hInstance win32.HINSTANCE

var w32VersionMajor, w32VersionMinor win32.DWORD

func backend_Init() error {
	go dispatchRequests()

	var err error
	dispatch(func() {
		hInstance = win32.HINSTANCE(win32.GetModuleHandle(""))
		if hInstance == nil {
			err = errors.New(fmt.Sprintf("Unable to determine hInstance; GetModuleHandle():", win32.GetLastErrorString()))
			return
		}

		// Get OS version, we use this to do some hack-ish fixes for different windows versions
		ret, vi := win32.GetVersionEx()
		if ret {
			w32VersionMajor = vi.DwMajorVersion()
			w32VersionMinor = vi.DwMinorVersion()
		} else {
			err = errors.New(fmt.Sprintf("Unable to determine windows version information; GetVersionEx():", win32.GetLastErrorString()))
			return
		}

	})

	if err != nil {
		return err
	}

	go eventLoop()

	return nil
}

func backend_Destroy() {
	stopDispatching()

	classNameCounter = 0
}
