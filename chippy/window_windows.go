// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"errors"
	"fmt"
	"image"
	"math"
	"sync"
	"time"
)

var windowsByHwnd = make(map[win32.HWND]*W32Window)

type W32Window struct {
	access sync.RWMutex

	icon, cursor image.Image

	wasOpened, isDestroyed, visible, decorated, minimized, maximized, fullscreen, alwaysOnTop,
	cursorGrabbed bool

	extentLeft, extentRight, extentBottom, extentTop, width, height, minWidth, minHeight,
	maxWidth, maxHeight uint

	x, y, cursorX, cursorY int

	aspectRatio float32

	screen Screen

	title string

	dc                                                                       win32.HDC
	hwnd                                                                     win32.HWND
	windowClass                                                              string
	styleFlags                                                               win32.DWORD
	lastWmSizingLeft, lastWmSizingRight, lastWmSizingBottom, lastWmSizingTop win32.LONG

	// OpenGL things here
	glConfig *GLConfig
}

func (w *W32Window) Open(screen Screen) (err error) {
	// If WasOpened returns true, this function is no-op.
	// if IsDestroyed returns true, this function is no-op.
	if w.WasOpened() || w.IsDestroyed() {
		return nil
	}

	w.access.Lock()
	defer w.access.Unlock()

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

		w.doSetWindowPos()

		// Make sure to enable wasOpened now so that doUpdateStyle sets the new style properly
		w.wasOpened = true
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
	// If WasOpened returns false, this function is no-op.
	// If IsDestroyed returns true, this function is no-op.
	if !w.WasOpened() || w.IsDestroyed() {
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
			if w.IsDestroyed() {
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

func (w *W32Window) SetTitle(title string) {
	w.panicIfDestroyed()

	w.access.Lock()
	defer w.access.Unlock()

	if w.title != title {
		w.title = title
		if w.wasOpened {
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

	if w.decorated != decorated {
		w.decorated = decorated
		if w.wasOpened && w.visible {
			unlock()
			dispatch(func() {
				w.doUpdateStyle()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetPosition(x, y int) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.x != x || w.y != y {
		w.x = x
		w.y = y
		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetSize(width, height uint) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.width != width || w.height != height {
		w.width = width
		w.height = height
		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetMinimumSize(width, height uint) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.minWidth != width || w.minHeight != height {
		w.minWidth = width
		w.minHeight = height
		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetMaximumSize(width, height uint) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.maxWidth != width || w.maxHeight != height {
		w.maxWidth = width
		w.maxHeight = height
		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetAspectRatio(ratio float32) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.aspectRatio != ratio {
		w.aspectRatio = ratio

		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetMinimized(minimized bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	w.minimized = minimized
	if w.minimized {
		if w.visible {
			if w.wasOpened {
				unlock()
				dispatch(func() {
					win32.ShowWindow(w.hwnd, win32.SW_MINIMIZE)
				})
			}
		}
	} else {
		if w.visible {
			if w.wasOpened {
				unlock()
				dispatch(func() {
					win32.ShowWindow(w.hwnd, win32.SW_RESTORE)
				})
			}
		}
	}
	unlock()
	if w.wasOpened {
		dispatch(func() {
			win32.EnableWindow(w.hwnd, true)
			w.doSetWindowPos()
		})
	}
}

func (w *W32Window) SetMaximized(maximized bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	w.maximized = maximized
	if w.maximized {
		if w.visible {
			if w.wasOpened {
				unlock()
				dispatch(func() {
					win32.ShowWindow(w.hwnd, win32.SW_MAXIMIZE)
				})
			}
		}
	} else {
		if w.wasOpened {
			unlock()
			dispatch(func() {
				win32.ShowWindow(w.hwnd, win32.SW_RESTORE)
			})
		}
	}
	unlock()
	if w.wasOpened {
		dispatch(func() {
			win32.EnableWindow(w.hwnd, true)
		})
	}
}

func (w *W32Window) SetFullscreen(fullscreen bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.fullscreen != fullscreen {
		w.fullscreen = fullscreen
		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doUpdateStyle()
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetAlwaysOnTop(alwaysOnTop bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()

	if w.alwaysOnTop != alwaysOnTop {
		w.alwaysOnTop = alwaysOnTop
		if w.wasOpened {
			unlock()
			dispatch(func() {
				w.doSetWindowPos()
			})
		}
	}
	unlock()
}

func (w *W32Window) SetIcon(icon image.Image) {
	w.panicIfDestroyed()

	// FIXME
	w.access.Lock()
	defer w.access.Unlock()
}

func (w *W32Window) SetCursor(Cursor image.Image) {
	w.panicIfDestroyed()

	// FIXME
	w.access.Lock()
	defer w.access.Unlock()
}

func (w *W32Window) SetCursorPosition(x, y int) {
	w.panicIfDestroyed()

	// FIXME
	w.access.Lock()
	defer w.access.Unlock()

	if w.cursorX != x || w.cursorY != y {
		w.cursorX = x
		w.cursorY = y
	}
}

func (w *W32Window) SetCursorGrabbed(grabbed bool) {
	w.panicIfDestroyed()

	// FIXME
	w.access.Lock()
	defer w.access.Unlock()

	if w.cursorGrabbed != grabbed {
		w.cursorGrabbed = grabbed
	}
}

func (w *W32Window) String() string {
	w.access.RLock()
	defer w.access.RUnlock()
	return fmt.Sprintf("Window(title=\"%s\", visible=%t, decorated=%t, minimized=%t, maximized=%t, fullscreen=%t, alwaysOnTop=%t, cursorGrabbed=%t, extents=[%d, %d, %d, %d], size=%dx%dpx, minimumSize=%dx%dpx, maximumSize=%dx%dpx, position=%dx%d, cursorPosition=%dx%d)", w.title, w.visible, w.decorated, w.minimized, w.maximized, w.fullscreen, w.alwaysOnTop, w.cursorGrabbed, w.extentLeft, w.extentRight, w.extentBottom, w.extentTop, w.width, w.height, w.minWidth, w.minHeight, w.maxWidth, w.maxHeight, w.x, w.y, w.cursorX, w.cursorY)
}

func (w *W32Window) WasOpened() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.wasOpened
}

func (w *W32Window) IsDestroyed() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.isDestroyed
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

func (w *W32Window) CursorGrabbed() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorGrabbed
}

// HWND returns the handle to this Window (HWND), this is of course, Windows specific, and
// is only useful when doing an small, select amount of things.
func (w *W32Window) HWND() win32.HWND {
	return w.hwnd
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

	if w.wasOpened {
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

		case win32.WM_SIZING:
			ratio := w.AspectRatio()
			if ratio != 0 {
				r := lParam.RECT()

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

		}
	}

	ret = win32.DefWindowProc(hwnd, msg, wParam, lParam)
	return
}

func (w *W32Window) panicIfDestroyed() {
	if w.IsDestroyed() {
		panic("Window has already been destroyed.")
	}
}

func backend_NewWindow() Window {
	return new(W32Window)
}
