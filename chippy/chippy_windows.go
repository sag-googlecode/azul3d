// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"code.google.com/p/azul3d/chippy/keyboard"
	"unsafe"
	"errors"
	"time"
	"sync"
	"fmt"
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

var windowsKeyDisabled bool
func SetWindowsKeyDisabled(disabled bool) {
	globalLock.Lock()
	defer globalLock.Unlock()

	windowsKeyDisabled = disabled
}

func WindowsKeyDisabled() bool {
	globalLock.RLock()
	defer globalLock.RUnlock()

	return windowsKeyDisabled
}

var hKeyboardHook win32.HHOOK

func keyboardHook(nCode win32.Int, wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT {
	if nCode < 0 || nCode != win32.HC_ACTION {
		return win32.CallNextHookEx(hKeyboardHook, nCode, wParam, lParam)
	}

	eatKeystroke := false
	if wParam == win32.WM_KEYDOWN || wParam == win32.WM_KEYUP {
		if WindowsKeyDisabled() {
			p := (*win32.KBDLLHOOKSTRUCT)(unsafe.Pointer(uintptr(lParam)))

			keysToEat := []win32.DWORD{
				win32.VK_LWIN,
				win32.VK_RWIN,
			}


			anyKeysToEat := false
			for _, k := range keysToEat {
				if k == p.VkCode {
					anyKeysToEat = true
					break
				}
			}

			if anyKeysToEat {
				for _, window := range windowsByHwnd {
					if window.Focused() {
						eatKeystroke = true

						// Send the event to the window
						state := keyboard.Down
						if wParam == win32.WM_KEYUP {
							state = keyboard.Up
						}

						switch(p.VkCode) {
							case win32.VK_LWIN:
								if state != window.leftWindowsState {
									window.leftWindowsState = state

									window.addKeyboardEvent(&keyboard.Event{
										State: state,
										Key: keyboard.LeftSuper,
										OSKey: keyboard.OSKey(p.ScanCode),
									})
								}

							case win32.VK_RWIN:
								if state != window.rightWindowsState {
									window.rightWindowsState = state

									window.addKeyboardEvent(&keyboard.Event{
										State: state,
										Key: keyboard.RightSuper,
										OSKey: keyboard.OSKey(p.ScanCode),
									})
								}
						}
					}
				}
			}
		}
	}

	if eatKeystroke {
		return 1
	}
	return win32.CallNextHookEx(hKeyboardHook, nCode, wParam, lParam)
}

//var classAtom win32.ATOM
//var windowClass *win32.WNDCLASSEX
var hInstance win32.HINSTANCE

var w32VersionMajor, w32VersionMinor win32.DWORD

func backend_Init() error {
	go dispatchRequests()

	windowsKeyDisabled = true

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

		hKeyboardHook = win32.SetLowLevelKeyboardHook(keyboardHook, hInstance, 0)
		if hKeyboardHook == nil {
			logger.Println("Failed to disable keyboard shortcuts; SetWindowsHookEx():", win32.GetLastErrorString())
		}
	})

	if err != nil {
		return err
	}

	go eventLoop()

	return nil
}

func backend_Destroy() {
	dispatch(func() {
		if hKeyboardHook != nil {
			if !win32.UnhookWindowsHookEx(hKeyboardHook) {
				logger.Println("Failed to unhook keyboard hook; UnhookWindowsHookEx():", win32.GetLastErrorString())
			}
		}
	})

	stopDispatching()

	classNameCounter = 0
}
