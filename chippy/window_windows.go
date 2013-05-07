// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	//"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy/mouse"
	"runtime"
	"errors"
	"fmt"
	"image"
	"math"
	"sync"
	"time"
)

var windowsByHwnd = make(map[win32.HWND]*W32Window)

type W32Window struct {
	eventDispatcher

	access sync.RWMutex

	icon, cursor image.Image

	opened, isDestroyed, focused, visible, decorated, minimized, maximized, fullscreen,
	alwaysOnTop, cursorGrabbed, cursorWithin, transparent bool

	extentLeft, extentRight, extentBottom, extentTop, width, height, minWidth, minHeight,
	maxWidth, maxHeight uint

	x, y, cursorX, cursorY int

	aspectRatio float32

	screen Screen

	title string

	dc, dcRender                                                             win32.HDC
	hwnd, hwndRender                                                         win32.HWND
	windowClass                                                              string
	styleFlags                                                               win32.DWORD
	lastWmSizingLeft, lastWmSizingRight, lastWmSizingBottom, lastWmSizingTop win32.LONG

	// OpenGL things here
	glConfig *GLConfig
}

func (w *W32Window) Open(screen Screen) (err error) {
	// If Opened returns true, this function is no-op.
	// if Destroyed returns true, this function is no-op.
	if w.Opened() || w.Destroyed() {
		return nil
	}

	w.access.Lock()
	defer w.access.Unlock()

	w.focused = true
	w.addFocusedEvent(w.focused)
	w.screen = screen

	dispatch(func() {
		// Get window extents
		titleHeight := win32.GetSystemMetrics(win32.SM_CYCAPTION)
		borderHeight := win32.GetSystemMetrics(win32.SM_CYSIZEFRAME)
		borderWidth := win32.GetSystemMetrics(win32.SM_CXSIZEFRAME)

		w.extentLeft = uint(borderWidth)
		w.extentRight = uint(borderWidth)
		w.extentBottom = uint(borderHeight)
		w.extentTop = uint(borderHeight + titleHeight)

		// Make our window class
		w.windowClass = fmt.Sprintf("ChippyWindow%d", nextCounter())
		windowClass := win32.NewWNDCLASSEX()
		windowClass.SetLpfnWndProc()
		windowClass.SetHbrBackground(nil)
		//windowClass.SetHIcon(win32.LoadIcon(hInstance, szAppName))
		//windowClass.SetHCursor(win32.LoadCursor(nil, win32.IDC_ARROW))
		//windowClass.SetHbrBackground(win32.IntToHBRUSH(win32.COLOR_WINDOW+2)) // Black background
		//windowClass.SetLpszMenuName(szAppName)

		windowClass.SetHInstance(hInstance)
		windowClass.SetLpszClassName(w.windowClass)

		classAtom := win32.RegisterClassEx(windowClass)
		if classAtom == 0 {
			err = errors.New(fmt.Sprintf("Unable to open window; RegisterClassEx(): %s", win32.GetLastErrorString()))
			return
		}

		// w.styleFlags will be updated to reflect current settings, that are passed into
		// CreateWindowEx to avoid some flicker
		w.doUpdateStyle()

		// SetPixelFormat() may only be called once -- so if we want to change any pixel format
		// values then our only option is to destroy the window and create it again.
		//
		// Since that would provide an largely noticable flicker to the user, we instead have an
		// 'rendering' window parented to our 'user managed' window, and we create the 'rendering'
		// window whenever we want to, thus bypassing the SetPixelFormat() issue noted above.
		//
		w.hwnd = win32.CreateWindowEx(0, w.windowClass, w.title, w.styleFlags, 0, 0, 0, 0, nil, nil, hInstance, nil)
		if w.hwnd == nil {
			err = errors.New(fmt.Sprintf("Unable to open window; CreateWindowEx(): %s", win32.GetLastErrorString()))
			return
		}
		w.dc = win32.GetDC(w.hwnd)
		if w.dc == nil {
			err = errors.New(fmt.Sprintf("Unable to get window DC; GetDC(): %s", win32.GetLastErrorString()))
			return
		}

		/*
			w.hwndRender = win32.CreateWindowEx(0, w.windowClass, "", win32.WS_CHILD, 0, 0, 100, 100, w.hwnd, nil, hInstance, nil)
			if w.hwndRender == nil {
				err = errors.New(fmt.Sprintf("Unable to open render window; CreateWindowEx(): %s", win32.GetLastErrorString()))
				return
			}
			w.dcRender = win32.GetDC(w.hwndRender)
			if w.dcRender == nil {
				err = errors.New(fmt.Sprintf("Unable to get render window DC; GetDC(): %s", win32.GetLastErrorString()))
				return
			}
		*/

		w.doSetWindowPos()

		// Make sure to enable opened now so that doUpdateStyle sets the new style properly
		w.opened = true
		w.doUpdateStyle()

		if w.visible {
			win32.ShowWindow(w.hwnd, win32.SW_SHOWDEFAULT)
			if w.minimized {
				win32.ShowWindow(w.hwnd, win32.SW_MINIMIZE)
			} else if w.maximized {
				win32.ShowWindow(w.hwnd, win32.SW_MAXIMIZE)
			}
		}
	})

	windowsByHwnd[w.hwnd] = w
	win32.RegisterWndProc(w.hwnd, mainWindowProc)
	return
}

func (w *W32Window) Destroy() {
	// If Opened returns false, this function is no-op.
	// If Destroyed returns true, this function is no-op.
	if !w.Opened() || w.Destroyed() {
		return
	}

	w.access.Lock()
	defer w.access.Unlock()
	w.isDestroyed = true

	win32.UnregisterWndProc(w.hwnd)
	delete(windowsByHwnd, w.hwnd)

	dispatch(func() {
		if !win32.DestroyWindow(w.hwnd) {
			logger.Println("Unable to destroy window; DestroyWindow():", win32.GetLastErrorString())
		}

		if !win32.UnregisterClass(w.windowClass, hInstance) {
			logger.Println("Failed to unregister window class; UnregisterClass():", win32.GetLastErrorString())
		}
	})
}

func (w *W32Window) Notify() {
	w.panicIfDestroyed()

	w.access.RLock()
	defer w.access.RUnlock()

	go func() {
		blinkDelay := 3 * time.Second
		maxBlinks := 3

		timesBlinked := 0
		for {
			if w.Destroyed() {
				return
			}
			if timesBlinked >= maxBlinks {
				return
			}
			timesBlinked += 1
			dispatch(func() {
				win32.FlashWindow(w.hwnd, true)
			})
			<-time.After(blinkDelay)
		}
	}()
}

func (w *W32Window) SetTransparent(transparent bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.transparent != transparent {
		w.transparent = transparent
		if w.opened {
			unlock()
			dispatch(func() {
				bb := win32.DWM_BLURBEHIND{}
				bb.DwFlags = win32.DWM_BB_ENABLE
				bb.FEnable = 1
				bb.HRgbBlur = 0
				bb.FTransitionOnMaximized = 1
				win32.DwmEnableBlurBehindWindow(w.hwnd, &bb)
			})
		}
	}
}

func (w *W32Window) SetTitle(title string) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.title != title {
		w.title = title
		if w.opened {
			dispatch(func() {
				if !win32.SetWindowText(w.hwnd, w.title) {
					logger.Println("Unable to set window title; SetWindowText():", win32.GetLastErrorString())
				}
			})
		}
	}
}

func (w *W32Window) SetVisible(visible bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	w.visible = visible
	if w.visible {
		unlock()
		dispatch(func() {
			win32.ShowWindow(w.hwnd, win32.SW_SHOW)
			win32.EnableWindow(w.hwnd, true)
		})
	} else {
		unlock()
		dispatch(func() {
			win32.ShowWindow(w.hwnd, win32.SW_HIDE)
		})
	}
}

func (w *W32Window) SetDecorated(decorated bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.decorated != decorated {
		w.decorated = decorated
		if w.opened && w.visible {
			unlock()
			dispatch(func() {
				w.doUpdateStyle()
			})
		}
	}
}

func (w *W32Window) SetPosition(x, y int) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.x != x || w.y != y {
		w.x = x
		w.y = y
		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetSize(width, height uint) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.width != width || w.height != height {
		w.width = width
		w.height = height
		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetMinimumSize(width, height uint) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.minWidth != width || w.minHeight != height {
		w.minWidth = width
		w.minHeight = height
		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetMaximumSize(width, height uint) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.maxWidth != width || w.maxHeight != height {
		w.maxWidth = width
		w.maxHeight = height
		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetAspectRatio(ratio float32) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.aspectRatio != ratio {
		w.aspectRatio = ratio

		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetMinimized(minimized bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	w.minimized = minimized
	if w.minimized {
		if w.visible {
			if w.opened {
				unlock()
				dispatch(func() {
					win32.ShowWindow(w.hwnd, win32.SW_MINIMIZE)
				})
			}
		}
	} else {
		if w.visible {
			if w.opened {
				unlock()
				dispatch(func() {
					win32.ShowWindow(w.hwnd, win32.SW_RESTORE)
				})
			}
		}
	}

	if w.opened {
		unlock()
		dispatch(func() {
			win32.EnableWindow(w.hwnd, true)
			w.doSetWindowPos()
		})
	}
}

func (w *W32Window) SetMaximized(maximized bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	w.maximized = maximized
	if w.maximized {
		if w.visible {
			if w.opened {
				unlock()
				dispatch(func() {
					win32.ShowWindow(w.hwnd, win32.SW_MAXIMIZE)
				})
			}
		}
	} else {
		if w.opened {
			unlock()
			dispatch(func() {
				win32.ShowWindow(w.hwnd, win32.SW_RESTORE)
			})
		}
	}

	if w.opened {
		unlock()
		dispatch(func() {
			win32.EnableWindow(w.hwnd, true)
		})
	}
}

func (w *W32Window) SetFullscreen(fullscreen bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.fullscreen != fullscreen {
		w.fullscreen = fullscreen
		if w.opened {
			unlock()
			dispatch(func() {
				w.doUpdateStyle()
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetAlwaysOnTop(alwaysOnTop bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.alwaysOnTop != alwaysOnTop {
		w.alwaysOnTop = alwaysOnTop
		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
}

func (w *W32Window) SetIcon(icon image.Image) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	w.icon = icon
}

func (w *W32Window) SetCursor(cursor image.Image) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	w.cursor = cursor
}

func (w *W32Window) SetCursorPosition(x, y int) {
	w.panicUnlessOpen()
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.cursorX != x || w.cursorY != y {
		w.cursorX = x
		w.cursorY = y

		if w.opened {
			unlock()
			dispatch(func() {
				if !win32.SetCursorPos(int32(w.x+w.cursorX), int32(w.y+w.cursorY)) {
					logger.Println("Unable to set cursor position: SetCursorPos():", win32.GetLastErrorString())
				}
			})
		}
	}
}

func (w *W32Window) SetCursorGrabbed(grabbed bool) {
	w.panicIfDestroyed()

	// FIXME
	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.cursorGrabbed != grabbed {
		w.cursorGrabbed = grabbed
	}
}

func (w *W32Window) String() string {
	w.access.RLock()
	defer w.access.RUnlock()
	return fmt.Sprintf("Window(title=\"%s\", focused=%t, visible=%t, decorated=%t, transparent=%t, minimized=%t, maximized=%t, fullscreen=%t, alwaysOnTop=%t, cursorGrabbed=%t, extents=[%d, %d, %d, %d], size=%dx%dpx, minimumSize=%dx%dpx, maximumSize=%dx%dpx, position=%dx%d, cursorPosition=%dx%d)", w.title, w.focused, w.visible, w.decorated, w.transparent, w.minimized, w.maximized, w.fullscreen, w.alwaysOnTop, w.cursorGrabbed, w.extentLeft, w.extentRight, w.extentBottom, w.extentTop, w.width, w.height, w.minWidth, w.minHeight, w.maxWidth, w.maxHeight, w.x, w.y, w.cursorX, w.cursorY)
}

func (w *W32Window) Opened() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.opened
}

func (w *W32Window) Destroyed() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.isDestroyed
}

func (w *W32Window) Transparent() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.transparent
}

func (w *W32Window) Screen() Screen {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.screen
}

func (w *W32Window) Extents() (left, right, bottom, top uint) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.extentLeft, w.extentRight, w.extentBottom, w.extentTop
}

func (w *W32Window) Focused() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.focused
}

func (w *W32Window) Title() string {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.title
}

func (w *W32Window) Visible() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.visible
}

func (w *W32Window) Decorated() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.decorated
}

func (w *W32Window) Position() (x, y int) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.x, w.y
}

func (w *W32Window) MaximumSize() (width, height uint) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.maxWidth, w.maxHeight
}

func (w *W32Window) Size() (width, height uint) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.width, w.height
}

func (w *W32Window) MinimumSize() (width, height uint) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.minWidth, w.minHeight
}

func (w *W32Window) AspectRatio() float32 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.aspectRatio
}

func (w *W32Window) Minimized() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.minimized
}

func (w *W32Window) Maximized() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.maximized
}

func (w *W32Window) Fullscreen() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.fullscreen
}

func (w *W32Window) AlwaysOnTop() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.alwaysOnTop
}

func (w *W32Window) Icon() image.Image {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.icon
}

func (w *W32Window) Cursor() image.Image {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursor
}

func (w *W32Window) CursorPosition() (x, y int) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorX, w.cursorY
}

func (w *W32Window) CursorWithin() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorWithin
}

func (w *W32Window) CursorGrabbed() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorGrabbed
}

// HWND returns the win32 handle to this Window, and it's child render window HWND.
//
// This is only useful when doing an few very select hack-ish things.
func (w *W32Window) HWND() (managed, render win32.HWND) {
	return w.hwnd, w.hwndRender
}

// Class returns the window class string of this Window (lpClassName), this is of course, Windows
// specific, and is only useful doing an small, select amount of things.
func (w *W32Window) Class() string {
	return w.windowClass
}

func (w *W32Window) newAttemptUnlocker() (unlock func()) {
	w.access.Lock()
	unlocked := false
	return func() {
		if !unlocked {
			unlocked = true
			w.access.Unlock()
		}
	}
}

func (w *W32Window) doUpdateVisibility() {
	/*
		    flag := win32.Int(win32.SW_HIDE)
			if w.visible {
		        if w.minimized {
		            flag = win32.SW_MINIMIZE
		        } else if w.maximized {
		            flag = win32.SW_MAXIMIZE
		        } else if !w.minimized || !w.maximized {
		            flag = win32.SW_RESTORE
		        } else {
		            flag = win32.SW_SHOW
		        }
			}
			win32.ShowWindow(w.hwnd, flag)
			win32.EnableWindow(w.hwnd, true)
	*/

}

func (w *W32Window) doUpdateStyle() {
	originalStyle := win32.GetWindowLongPtr(w.hwnd, win32.GWL_STYLE)

	if w.decorated && !w.fullscreen {
		w.styleFlags |= win32.WS_OVERLAPPEDWINDOW
		w.styleFlags |= win32.WS_BORDER
	} else {
		w.styleFlags ^= win32.WS_OVERLAPPEDWINDOW
	}

	if w.visible {
		w.styleFlags |= win32.WS_VISIBLE
	}

	if w.opened {
		if win32.DWORD(originalStyle) != w.styleFlags {
			win32.SetWindowLongPtr(w.hwnd, win32.GWL_STYLE, win32.LONG_PTR(w.styleFlags))

			//win32.EnableWindow(w.hwnd, true)
			//if w.visible {
			//    win32.ShowWindow(w.hwnd, win32.SW_SHOWNA)
			//}
			w.doSetWindowPos()
		}
	}
}

func (w *W32Window) doSetWindowPos() {
	// win32.SWP_ASYNCWINDOWPOS|win32.SWP_FRAMECHANGED|win32.SWP_NOMOVE|win32.SWP_NOSIZE|win32.SWP_NOZORDER|win32.SWP_NOOWNERZORDER

	x := win32.Int(w.x - int(w.extentLeft))
	y := win32.Int(w.y - int(w.extentTop))

	// Clip to maxWidth/maxHeight and minWidth/minHeight, append extents so that width/height is
	// the client region specifically.
	width := float64(w.width)
	if w.maxWidth > 0 {
		width = math.Min(float64(w.width), float64(w.maxWidth))
	}
	width = math.Max(width, float64(w.minWidth))
	width += float64(w.extentLeft)
	width += float64(w.extentRight)

	height := float64(w.height)
	if w.maxHeight > 0 {
		height = math.Min(height, float64(w.maxHeight))
	}
	height = math.Max(height, float64(w.minHeight))
	height += float64(w.extentBottom)
	height += float64(w.extentTop)

	ratio := w.aspectRatio
	if ratio != 0.0 {
		if ratio > 1.0 {
			// Wider instead of taller
			width = float64(ratio * float32(height))
			width = float64(ratio * float32(height))
		} else {
			// Taller instead of wider
			height = float64((1.0 / ratio) * float32(width))
			height = float64((1.0 / ratio) * float32(width))
		}
	}

	insertAfter := win32.HWND_NOTOPMOST
	if w.alwaysOnTop {
		insertAfter = win32.HWND_TOPMOST
	}
	//		win32.SetWindowPos(w.hwnd, flag, 0, 0, 0, 0, win32.SWP_NOMOVE|win32.SWP_NOSIZE)

	if w.fullscreen {
		x = 0
		y = 0
		sm := w.screen.ScreenMode()
		w, h := sm.Resolution()
		width, height = float64(w), float64(h)
	}

	// |win32.SWP_NOZORDER|win32.SWP_NOOWNERZORDER
	if !win32.SetWindowPos(w.hwnd, insertAfter, x, y, win32.Int(width), win32.Int(height), win32.SWP_FRAMECHANGED) {
		logger.Println("Unable to set window position; SetWindowPos():", win32.GetLastErrorString())
	}
}

// Our MS windows event handler
func mainWindowProc(hwnd win32.HWND, msg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) (ret win32.LRESULT) {
	w, ok := windowsByHwnd[hwnd]
	if ok {
		switch msg {
		//case win32.WM_PAINT:
		//	logger.Println("WM_PAINT")

		case win32.WM_GETMINMAXINFO:
			minWidth, minHeight := w.MinimumSize()
			maxWidth, maxHeight := w.MaximumSize()
			extentLeft, extentRight, extentBottom, extentTop := w.Extents()
			ratio := w.AspectRatio()

			// Add extents, so we operate on client region space only
			newMinWidth := minWidth + extentLeft + extentRight
			newMaxWidth := maxWidth + extentLeft + extentRight
			newMinHeight := minHeight + extentBottom + extentTop
			newMaxHeight := maxHeight + extentBottom + extentTop

			if ratio != 0.0 {
				if ratio > 1.0 {
					// Wider instead of taller
					newMinWidth = uint(ratio * float32(newMinHeight))
					newMaxWidth = uint(ratio * float32(newMaxHeight))
				} else {
					// Taller instead of wider
					newMinHeight = uint((1.0 / ratio) * float32(newMinWidth))
					newMaxHeight = uint((1.0 / ratio) * float32(newMaxWidth))
				}
			}

			// Set maximum and minimum window sizes, 0 means unlimited
			minMaxInfo := lParam.MINMAXINFO()

			if minWidth > 0 {
				minMaxInfo.PtMinTrackSize().SetX(int32(newMinWidth))
			}
			if minHeight > 0 {
				minMaxInfo.PtMinTrackSize().SetY(int32(newMinHeight))
			}

			if maxWidth > 0 {
				minMaxInfo.PtMaxTrackSize().SetX(int32(newMaxWidth))
			}
			if maxHeight > 0 {
				minMaxInfo.PtMaxTrackSize().SetY(int32(newMaxHeight))
			}
			return 0

		case win32.WM_SIZING:
			ratio := w.AspectRatio()
			r := lParam.RECT()

			if ratio != 0 {
				width := r.Right() - r.Left()
				height := r.Bottom() - r.Top()

				newHeight := (1.0 / ratio) * float32(width)
				newWidth := ratio * float32(height)

				newRight := r.Left() + win32.LONG(newWidth)
				//newLeft := r.Right() - win32.LONG(newWidth)
				newBottom := r.Top() + win32.LONG(newHeight)
				newTop := r.Bottom() - win32.LONG(newHeight)

				if wParam == win32.WMSZ_RIGHT || wParam == win32.WMSZ_LEFT {
					r.SetBottom(newBottom)
				} else if wParam == win32.WMSZ_BOTTOM || wParam == win32.WMSZ_TOP {
					r.SetRight(newRight)

				} else if wParam == win32.WMSZ_TOPLEFT || wParam == win32.WMSZ_TOPRIGHT {
					r.SetTop(newTop)
				} else if wParam == win32.WMSZ_BOTTOMLEFT || wParam == win32.WMSZ_BOTTOMRIGHT {
					r.SetBottom(newBottom)
				}

				w.lastWmSizingLeft = r.Left()
				w.lastWmSizingRight = r.Right()
				w.lastWmSizingBottom = r.Bottom()
				w.lastWmSizingTop = r.Top()
			}

			newWidth := uint(w.lastWmSizingRight - w.lastWmSizingLeft)
			newHeight := uint(w.lastWmSizingBottom - w.lastWmSizingTop)

			if newWidth != 0 && newHeight != 0 {
				if w.width != newWidth || w.height != newHeight {
					w.width = newWidth
					w.height = newHeight

					w.addSizeEvent([]uint{w.width, w.height})
				}
			}
			return 0

		case win32.WM_SIZE:
			if wParam == win32.SIZE_MAXIMIZED {
				if w.minimized != false {
					w.minimized = false
					w.addMinimizedEvent(w.minimized)
				}
				if w.maximized != true {
					w.maximized = true
					w.addMaximizedEvent(w.maximized)
				}
			} else if wParam == win32.SIZE_MINIMIZED {
				if w.minimized != true {
					w.minimized = true
					w.addMinimizedEvent(w.minimized)
				}
				if w.maximized != false {
					w.maximized = false
					w.addMaximizedEvent(w.maximized)
				}
			} else {
				if w.minimized != false {
					w.minimized = false
					w.addMinimizedEvent(w.minimized)
				}
				if w.maximized != false {
					w.maximized = false
					w.addMaximizedEvent(w.maximized)
				}
			}

			if wParam != win32.SIZE_MINIMIZED {
				newWidth := uint(lParam.LOWORD())
				newHeight := uint(lParam.HIWORD())

				if w.width != newWidth || w.height != newHeight {
					w.width = newWidth
					w.height = newHeight

					w.addSizeEvent([]uint{w.width, w.height})
				}
			}
			return 0

		case win32.WM_MOVE:
			xPos := int(lParam.LOWORD())
			yPos := int(lParam.HIWORD())

			if !win32.IsIconic(w.hwnd) {
				if w.x != xPos || w.y != yPos {
					w.x = xPos
					w.y = yPos
					w.addPositionEvent([]int{w.x, w.y})
				}
			}
			return 0

		case win32.WM_ACTIVATE:
			if wParam.LOWORD() == win32.WA_INACTIVE || wParam.HIWORD() != 0 {
				if w.focused {
					w.focused = false
					w.addFocusedEvent(w.focused)
				}
			} else {
				if !w.focused {
					w.focused = true
					w.addFocusedEvent(w.focused)
				}
			}
			return 0

		case win32.WM_GETICON:
			logger.Println("WM_GETICON")
			return 0

		case win32.WM_KEYDOWN:
			//fmt.Println(lParam, lParam.LOWORD(), lParam.HIWORD(), string(lParam.HIWORD()))
			//fmt.Println(0x09)
			return 0

		case win32.WM_MOUSEMOVE:
			xPos := int(lParam.LOWORD())
			yPos := int(lParam.HIWORD())

			if w.cursorX != xPos || w.cursorY != yPos {
				w.cursorX = xPos
				w.cursorY = yPos

				w.addCursorPositionEvent([]int{w.cursorX, w.cursorY})
			}
			return 0

		// Mouse Buttons
		case win32.WM_LBUTTONDOWN:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Left,
				State:  mouse.Down,
			})
			return 0

		case win32.WM_LBUTTONUP:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Left,
				State:  mouse.Up,
			})
			return 0

		case win32.WM_RBUTTONDOWN:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Right,
				State:  mouse.Down,
			})
			return 0

		case win32.WM_RBUTTONUP:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Right,
				State:  mouse.Up,
			})
			return 0

		case win32.WM_MBUTTONDOWN:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Wheel,
				State:  mouse.Down,
			})
			return 0

		case win32.WM_MBUTTONUP:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Wheel,
				State:  mouse.Up,
			})
			return 0

		case win32.WM_XBUTTONDOWN:
			fmt.Println(lParam, lParam.LOWORD(), lParam.HIWORD())
			return 0

		case win32.WM_XBUTTONUP:
			fmt.Println(lParam, lParam.LOWORD(), lParam.HIWORD())
			return 0

		case win32.WM_MOUSEWHEEL:
			fmt.Println(lParam, lParam.LOWORD(), lParam.HIWORD())
			return 0

		case win32.WM_MOUSEHWHEEL:
			fmt.Println(lParam, lParam.LOWORD(), lParam.HIWORD())
			return 0

		//default:
		//	fmt.Printf("0x%x\n", msg)

		case win32.WM_CLOSE:
			if !w.addCloseEvent() {
				go w.Destroy()
			}
			return 0
		}
	}

	// Note: This line is EXTREAMELY important when GOMAXPROCS=1!
	//
	// Many win32 API calls simply perform an callback into mainWindowProc and provide little to
	// no information as to which ones do and when they do.
	//
	// Input devices (mice) that generate an lot of WM_MOUSEMOVE events will cause an temporary
	// block when these input devices send too many messages, the temporary block stops once the
	// devices stop sending the messages.
	//
	// The effect this has is, an call to SwapBuffers() for instance, will block as long as the
	// user is moving their mouse inside the window constantly:
	//
	// /_Causing the render loop to pause while the user moves their mouse_/
	//
	// This will pass feedback to the render loop without blocking when the messages are spammed.
	runtime.Gosched()

	ret = win32.DefWindowProc(hwnd, msg, wParam, lParam)
	return
}

func (w *W32Window) panicIfDestroyed() {
	if w.Destroyed() {
		panic("Window has already been destroyed.")
	}
}

func (w *W32Window) panicUnlessOpen() {
	if !w.Opened() {
		panic("Window is not open.")
	}
}

func backend_NewWindow() Window {
	return new(W32Window)
}
