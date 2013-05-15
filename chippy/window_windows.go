// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/chippy/mouse"
	"code.google.com/p/azul3d/chippy/thirdparty/resize"
	"errors"
	"unsafe"
	"image"
	"math"
	"sync"
	"time"
	"fmt"
)

var windowsByHwnd = make(map[win32.HWND]*W32Window)

type loadedCursor struct {
	hCursor                             win32.HICON
	cursorColorBitmap, cursorMaskBitmap win32.HBITMAP
	cursorColorBits, cursorMaskBits     []uint32
}

type W32Window struct {
	eventDispatcher

	access sync.RWMutex

	icon image.Image
	cursor *Cursor

	opened, isDestroyed, focused, visible, decorated, minimized, maximized, fullscreen,
	alwaysOnTop, cursorGrabbed, cursorWithin, transparent bool

	extentLeft, extentRight, extentBottom, extentTop, width, height, preFullscreenWidth,
	preFullscreenHeight, minWidth, minHeight, maxWidth, maxHeight uint

	x, y, preFullscreenX, preFullscreenY, cursorX, cursorY, lastCursorX, lastCursorY,
	preGrabCursorX, preGrabCursorY int

	aspectRatio float32

	originalScreen, screen Screen

	title string

	cursors                                                                  map[*Cursor]*loadedCursor
	loadedCursor                                                             *loadedCursor
	hIcon, hSmIcon                                                           win32.HICON
	iconColorBitmap, iconMaskBitmap, smIconColorBitmap, smIconMaskBitmap     win32.HBITMAP
	iconColorBits, iconMaskBits, smIconColorBits, smIconMaskBits             []uint32
	dc, dcRender                                                             win32.HDC
	hwnd, hwndRender                                                         win32.HWND
	windowClass                                                              string
	styleFlags                                                               win32.DWORD
	lastWmSizingLeft, lastWmSizingRight, lastWmSizingBottom, lastWmSizingTop win32.LONG
	leftShiftDown, leftAltDown, leftCtrlDown bool
	leftWindowsState, rightWindowsState keyboard.State

	// Blit things here
	blitBitmap win32.HBITMAP
	blitBitmapDc win32.HDC
	blitBits []uint32

	// OpenGL things here
	glConfig *GLConfig
	glPixelFormatSet bool
}

func (w *W32Window) Open(screen Screen) (err error) {
	// If Opened returns true, this function is no-op.
	// if Destroyed returns true, this function is no-op.
	if w.Opened() || w.Destroyed() {
		return nil
	}

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.icon == nil {
		w.icon = defaultIcon
	}

	w.cursors = make(map[*Cursor]*loadedCursor)
	w.focused = true
	w.addFocusedEvent(w.focused)
	w.originalScreen = screen
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

		err = w.doRebuildWindow()
		if err != nil {
			return
		}
		w.doMakeIcon()

		supportRawInput := w32VersionMajor >= 5 && w32VersionMinor >= 1
		if supportRawInput {
			rid := win32.RAWINPUTDEVICE{}
			rid.UsagePage = win32.HID_USAGE_PAGE_GENERIC
			rid.Usage = win32.HID_USAGE_GENERIC_MOUSE
			rid.Flags = win32.RIDEV_INPUTSINK
			rid.Target = w.hwnd
			win32.RegisterRawInputDevices(&rid, 1, win32.UINT(unsafe.Sizeof(rid)))
		}

		// To let the application know of current toggle key states, we need to send them right now
		// otherwise they might already be on or off -- something the application might want to
		// know.
		var state keyboard.State

		if (win32.GetKeyState(win32.VK_CAPITAL) & 0x0001) != 0 {
			state = keyboard.On
		} else {
			state = keyboard.Off
		}
		w.addKeyboardEvent(&keyboard.Event{
			Key: keyboard.CapsLock,
			OSKey: keyboard.OSKey(win32.VK_CAPITAL),
			Unicode: "",
			State: state,
		})

		if (win32.GetKeyState(win32.VK_NUMLOCK) & 0x0001) != 0 {
			state = keyboard.On
		} else {
			state = keyboard.Off
		}
		w.addKeyboardEvent(&keyboard.Event{
			Key: keyboard.NumLock,
			OSKey: keyboard.OSKey(win32.VK_NUMLOCK),
			Unicode: "",
			State: state,
		})

		if (win32.GetKeyState(win32.VK_SCROLL) & 0x0001) != 0 {
			state = keyboard.On
		} else {
			state = keyboard.Off
		}
		w.addKeyboardEvent(&keyboard.Event{
			Key: keyboard.ScrollLock,
			OSKey: keyboard.OSKey(win32.VK_SCROLL),
			Unicode: "",
			State: state,
		})

		// Store the Shift, Alt, and Ctrl key states
		w.leftShiftDown = (uint16(win32.GetAsyncKeyState(win32.VK_LSHIFT)) & 0x8000) != 0
		w.leftAltDown = (uint16(win32.GetAsyncKeyState(win32.VK_LMENU)) & 0x8000) != 0
		w.leftCtrlDown = (uint16(win32.GetAsyncKeyState(win32.VK_LCONTROL)) & 0x8000) != 0


	})

	unlock()

	if err == nil {
		w.SetCursor(nil)
	}
	return
}

func (w *W32Window) doRebuildWindow() (err error) {
	if w.opened {
		win32.UnregisterWndProc(w.hwnd)
		delete(windowsByHwnd, w.hwnd)

		if !win32.DestroyWindow(w.hwnd) {
			logger.Println("Unable to destroy window; DestroyWindow():", win32.GetLastErrorString())
		}

		if !win32.UnregisterClass(w.windowClass, hInstance) {
			logger.Println("Failed to unregister window class; UnregisterClass():", win32.GetLastErrorString())
		}
		w.hwnd = nil
	}

	w.glPixelFormatSet = false

	// Make our window class
	w.windowClass = fmt.Sprintf("ChippyWindow%d", nextCounter())
	windowClass := win32.NewWNDCLASSEX()
	windowClass.SetLpfnWndProc()
	windowClass.SetHbrBackground(win32.CreateSolidBrush(0x00000000))
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

	w.doUpdateTransparency()
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

	dispatch(func() {
		win32.UnregisterWndProc(w.hwnd)
		delete(windowsByHwnd, w.hwnd)

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

func (w *W32Window) PixelClear(x, y, width, height uint) {
	if !w.opened {
		return
	}

	dispatch(func() {
		r := &win32.RECT{}
		r.SetLeft(win32.LONG(x))
		r.SetTop(win32.LONG(y))
		r.SetBottom(win32.LONG(y + height))
		r.SetRight(win32.LONG(x + width))
		win32.FillRect(w.dc, r, win32.HBRUSH(win32.GetStockObject(win32.BLACK_BRUSH)))
	})
}

func (w *W32Window) PixelBlit(x, y uint, image image.Image) {
	if !w.opened {
		return
	}

	dispatch(func() {
		if w.blitBitmap != nil{
			if !win32.DeleteDC(w.blitBitmapDc) {
				logger.Println("Unable to delete blit bitmap; DeleteDC():", win32.GetLastErrorString())
			}

			if !win32.DeleteObject(win32.HGDIOBJ(w.blitBitmap)) {
				logger.Println("Unable to delete blit bitmap; DeleteObject():", win32.GetLastErrorString())
			}
		}

		bounds := image.Bounds()
		width := uint(bounds.Max.X)
		height := uint(bounds.Max.Y)

		w.blitBitmapDc = win32.CreateCompatibleDC(w.dc)
		w.blitBitmap = win32.CreateCompatibleBitmap(w.dc, win32.Int(width), win32.Int(height))
		if win32.SelectObject(w.blitBitmapDc, win32.HGDIOBJ(w.blitBitmap)) == nil {
			logger.Println("Unable to blit; SelectObject():", win32.GetLastErrorString())
			return
		}

		w.blitBits = make([]uint32, width * height)

		for y := 0; y < int(height); y++ {
			for x := 0; x < int(width); x++ {
				r, g, b, a := image.At(x, y).RGBA()
				c := (uint32(a >> 8) << 24) | (uint32(r >> 8) << 16) | (uint32(g >> 8) << 8) | (uint32(b >> 8) << 0)

				index := (int(height)-1 - y) * int(width)
				index += x
				w.blitBits[index] = c
			}
		}

		bitmapInfo := win32.BITMAPINFO{
			BmiHeader: win32.BITMAPINFOHEADER{
				Size: win32.DWORD(unsafe.Sizeof(win32.BITMAPINFOHEADER{})),
				Width: win32.LONG(width),
				Height: win32.LONG(height),
				Planes: 1,
				BitCount: 32,
				Compression: win32.BI_RGB,
				SizeImage: 0,
				XPelsPerMeter: 0,
				YPelsPerMeter: 0,
				ClrUsed: 0,
				ClrImportant: 0,
			},
		}
		if win32.SetDIBits(w.dc, w.blitBitmap, 0, win32.UINT(height), unsafe.Pointer(&w.blitBits[0]), &bitmapInfo, win32.DIB_RGB_COLORS) == 0 {
			logger.Println("Unable to blit; SetDiBits():", win32.GetLastErrorString())
			return
		}

		blend := win32.BLENDFUNCTION{
			BlendOp: win32.AC_SRC_OVER,
			BlendFlags: 0,
			SourceConstantAlpha: 255,
			AlphaFormat: win32.AC_SRC_ALPHA,
		}

		if !win32.AlphaBlend(w.dc, win32.Int(x), win32.Int(y), win32.Int(width), win32.Int(height), w.blitBitmapDc, 0, 0, win32.Int(width), win32.Int(height), &blend) {
			logger.Println("Unable to blit: TransparentBlt():", win32.GetLastErrorString())
			return
		}
	})
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
				w.doUpdateTransparency()
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

		if w.fullscreen {
			// Entering fullscreen: Save window size
			w.preFullscreenX = w.x
			w.preFullscreenY = w.y
			w.preFullscreenWidth = w.width
			w.preFullscreenHeight = w.height
		} else {
			// Leaving fullscreen: Restore original size
			w.x = w.preFullscreenX
			w.y = w.preFullscreenY
			w.width = w.preFullscreenWidth
			w.height = w.preFullscreenHeight
		}

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

	if w.opened {
		unlock()
		dispatch(func() {
			w.doMakeIcon()
		})
	}

	w.icon = icon
}

func (w *W32Window) doPrepareCursor(cursor *Cursor) {
	lc := new(loadedCursor)

	cursorWidth := win32.GetSystemMetrics(win32.SM_CXCURSOR)
	cursorHeight := win32.GetSystemMetrics(win32.SM_CYCURSOR)

	cursorImage := resize.Resize(cursor.Image, cursor.Image.Bounds(), int(cursorWidth), int(cursorHeight))

	cursorBitmapInfo := win32.BITMAPINFO{
		BmiHeader: win32.BITMAPINFOHEADER{
			Size: win32.DWORD(unsafe.Sizeof(win32.BITMAPINFOHEADER{})),
			Width: win32.LONG(cursorWidth),
			Height: win32.LONG(cursorHeight),
			Planes: 1,
			BitCount: 32,
			Compression: win32.BI_RGB,
			SizeImage: 0,
			XPelsPerMeter: 0,
			YPelsPerMeter: 0,
			ClrUsed: 0,
			ClrImportant: 0,
		},
	}

	lc.cursorColorBitmap = win32.CreateCompatibleBitmap(w.dc, cursorWidth, cursorHeight)
	lc.cursorColorBits = make([]uint32, cursorWidth * cursorHeight)
	for y := 0; y < int(cursorHeight); y++ {
		for x := 0; x < int(cursorWidth); x++ {
			r, g, b, _ := cursorImage.At(x, y).RGBA()
			c := (uint32(r >> 8) << 16) | (uint32(g >> 8) << 8) | (uint32(b >> 8) << 0)

			index := (int(cursorHeight)-1 - y) * int(cursorWidth)
			index += x
			lc.cursorColorBits[index] = c //0xFF0000
		}
	}
	if win32.SetDIBits(w.dc, lc.cursorColorBitmap, 0, win32.UINT(cursorHeight), unsafe.Pointer(&lc.cursorColorBits[0]), &cursorBitmapInfo, win32.DIB_RGB_COLORS) == 0 {
		logger.Println("Unable to set cursor; SetDiBits():", win32.GetLastErrorString())
		return
	}

	lc.cursorMaskBitmap = win32.CreateCompatibleBitmap(w.dc, cursorWidth, cursorHeight)
	lc.cursorMaskBits = make([]uint32, cursorWidth * cursorHeight)
	for y := 0; y < int(cursorHeight); y++ {
		for x := 0; x < int(cursorWidth); x++ {
			_, _, _, a := cursorImage.At(x, y).RGBA()
			c := uint32(0xFFFFFF)
			if a > 0 {
				c = 0
			}

			index := (int(cursorHeight)-1 - y) * int(cursorWidth)
			index += x
			lc.cursorMaskBits[index] = c
		}
	}
	if win32.SetDIBits(w.dc, lc.cursorMaskBitmap, 0, win32.UINT(cursorHeight), unsafe.Pointer(&lc.cursorMaskBits[0]), &cursorBitmapInfo, win32.DIB_RGB_COLORS) == 0 {
		logger.Println("Unable to set cursor; SetDiBits():", win32.GetLastErrorString())
		return
	}

	cursorInfo := win32.ICONINFO{
		FIcon: 0,
		XHotspot: win32.DWORD(cursor.X),
		YHotspot: win32.DWORD(cursor.Y),
		HbmMask: lc.cursorMaskBitmap,
		HbmColor: lc.cursorColorBitmap,
	}

	lc.hCursor = win32.CreateIconIndirect(&cursorInfo)
	if lc.hCursor == nil {
		logger.Println("Unable to set cursor; CreateIconIndirect():", win32.GetLastErrorString())
		return
	}

	w.cursors[cursor] = lc
}

func (w *W32Window) PrepareCursor(cursor *Cursor) {
	w.panicIfDestroyed()

	if cursor == nil {
		panic("PrepareCursor(): Cannot prepare nil cursor!")
	}

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.opened {
		_, ok := w.cursors[cursor]
		if ok {
			// It's already loaded!
			return
		}

		unlock()
		dispatch(func() {
			w.doPrepareCursor(cursor)
		})
	}
}

func (w *W32Window) FreeCursor(cursor *Cursor) {
	w.panicIfDestroyed()

	if cursor == nil {
		panic("FreeCursor(): Cannot free nil cursor!")
	}

	// Special case: We're using this cursor right now
	if cursor == w.Cursor() {
		// Restore the default cursor first
		w.SetCursor(nil)
	}

	unlock := w.newAttemptUnlocker()
	defer unlock()

	lc, ok := w.cursors[cursor]
	if ok {
		delete(w.cursors, cursor)

		unlock()
		dispatch(func() {
			if !win32.DestroyCursor(win32.HCURSOR(lc.hCursor)) {
				logger.Println("Failed to free cursor; DestroyCursor():", win32.GetLastErrorString())
			}

			if !win32.DeleteObject(win32.HGDIOBJ(lc.cursorColorBitmap)) {
				logger.Println("Failed to free cursor; DeleteObject(cursorColorBitmap) failed!")
			}

			if !win32.DeleteObject(win32.HGDIOBJ(lc.cursorMaskBitmap)) {
				logger.Println("Failed to free cursor; DeleteObject(cursorMaskBitmap) failed!")
			}
		})
	}
}

func (w *W32Window) SetCursor(cursor *Cursor) {
	w.panicIfDestroyed()

	if cursor != nil {
		w.PrepareCursor(cursor)
	}

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.cursor != cursor {
		w.cursor = cursor
		w.loadedCursor = nil

		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetCursor()
			})
		}
	}
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
				w.doSetCursorPos()
			})
		}
	}
}

func (w *W32Window) SetCursorGrabbed(grabbed bool) {
	w.panicIfDestroyed()

	unlock := w.newAttemptUnlocker()
	defer unlock()

	if w.cursorGrabbed != grabbed {
		w.cursorGrabbed = grabbed
		if w.cursorGrabbed {
			w.preGrabCursorX = w.cursorX
			w.preGrabCursorY = w.cursorY
		} else {
			w.cursorX = w.preGrabCursorX
			w.cursorY = w.preGrabCursorY
		}
		if w.opened {
			unlock()
			dispatch(func() {
				w.doSetCursor()
				w.doSetCursorPos()
			})
		}
	}
}

func (w *W32Window) String() string {
	w.access.RLock()
	defer w.access.RUnlock()
	return fmt.Sprintf("Window(title=%q, focused=%t, visible=%t, decorated=%t, transparent=%t, minimized=%t, maximized=%t, fullscreen=%t, alwaysOnTop=%t, cursorGrabbed=%t, extents=[%d, %d, %d, %d], size=%dx%dpx, minimumSize=%dx%dpx, maximumSize=%dx%dpx, position=%dx%d, cursorPosition=%dx%d)", w.title, w.focused, w.visible, w.decorated, w.transparent, w.minimized, w.maximized, w.fullscreen, w.alwaysOnTop, w.cursorGrabbed, w.extentLeft, w.extentRight, w.extentBottom, w.extentTop, w.width, w.height, w.minWidth, w.minHeight, w.maxWidth, w.maxHeight, w.x, w.y, w.cursorX, w.cursorY)
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

func (w *W32Window) OriginalScreen() Screen {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.originalScreen
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

func (w *W32Window) Cursor() *Cursor {
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

func (w *W32Window) doMakeIcon() {
	///////////////////
	// Standard icon //
	///////////////////
	if w.hIcon != nil {
		if !win32.DestroyIcon(w.hIcon) {
			logger.Println("Failed to destroy icon; DestroyIcon():", win32.GetLastErrorString())
		}

		if !win32.DeleteObject(win32.HGDIOBJ(w.iconColorBitmap)) {
			logger.Println("Failed to destroy icon; DeleteObject(iconColorBitmap) failed!")
		}

		if !win32.DeleteObject(win32.HGDIOBJ(w.iconMaskBitmap)) {
			logger.Println("Failed to destroy icon; DeleteObject(iconMaskBitmap) failed!")
		}
	}

	iconWidth := win32.GetSystemMetrics(win32.SM_CXICON)
	iconHeight := win32.GetSystemMetrics(win32.SM_CYICON)

	iconImage := resize.Resize(w.icon, w.icon.Bounds(), int(iconWidth), int(iconHeight))

	iconBitmapInfo := win32.BITMAPINFO{
		BmiHeader: win32.BITMAPINFOHEADER{
			Size: win32.DWORD(unsafe.Sizeof(win32.BITMAPINFOHEADER{})),
			Width: win32.LONG(iconWidth),
			Height: win32.LONG(iconHeight),
			Planes: 1,
			BitCount: 32,
			Compression: win32.BI_RGB,
			SizeImage: 0,
			XPelsPerMeter: 0,
			YPelsPerMeter: 0,
			ClrUsed: 0,
			ClrImportant: 0,
		},
	}

	w.iconColorBitmap = win32.CreateCompatibleBitmap(w.dc, iconWidth, iconHeight)
	w.iconColorBits = make([]uint32, iconWidth * iconHeight)
	for y := 0; y < int(iconHeight); y++ {
		for x := 0; x < int(iconWidth); x++ {
			r, g, b, _ := iconImage.At(x, y).RGBA()
			c := (uint32(r >> 8) << 16) | (uint32(g >> 8) << 8) | (uint32(b >> 8) << 0)

			index := (int(iconHeight)-1 - y) * int(iconWidth)
			index += x
			w.iconColorBits[index] = c //0xFF0000
		}
	}
	if win32.SetDIBits(w.dc, w.iconColorBitmap, 0, win32.UINT(iconHeight), unsafe.Pointer(&w.iconColorBits[0]), &iconBitmapInfo, win32.DIB_RGB_COLORS) == 0 {
		logger.Println("Unable to set icon; SetDiBits():", win32.GetLastErrorString())
		return
	}

	w.iconMaskBitmap = win32.CreateCompatibleBitmap(w.dc, iconWidth, iconHeight)
	w.iconMaskBits = make([]uint32, iconWidth * iconHeight)
	for y := 0; y < int(iconHeight); y++ {
		for x := 0; x < int(iconWidth); x++ {
			_, _, _, a := iconImage.At(x, y).RGBA()
			c := uint32(0xFFFFFF)
			if a > 0 {
				c = 0
			}

			index := (int(iconHeight)-1 - y) * int(iconWidth)
			index += x
			w.iconMaskBits[index] = c
		}
	}
	if win32.SetDIBits(w.dc, w.iconMaskBitmap, 0, win32.UINT(iconHeight), unsafe.Pointer(&w.iconMaskBits[0]), &iconBitmapInfo, win32.DIB_RGB_COLORS) == 0 {
		logger.Println("Unable to set icon; SetDiBits():", win32.GetLastErrorString())
		return
	}

	iconInfo := win32.ICONINFO{
		FIcon: 1,
		XHotspot: 0,
		YHotspot: 0,
		HbmMask: w.iconMaskBitmap,
		HbmColor: w.iconColorBitmap,
	}

	w.hIcon = win32.CreateIconIndirect(&iconInfo)
	if w.hIcon == nil {
		logger.Println("Unable to set icon; CreateIconIndirect():", win32.GetLastErrorString())
		return
	}

	////////////////
	// Small icon //
	////////////////
	if w.hSmIcon != nil {
		if !win32.DestroyIcon(w.hSmIcon) {
			logger.Println("Failed to destroy icon; DestroyIcon():", win32.GetLastErrorString())
		}

		if !win32.DeleteObject(win32.HGDIOBJ(w.smIconColorBitmap)) {
			logger.Println("Failed to destroy icon; DeleteObject(smIconColorBitmap) failed!")
		}

		if !win32.DeleteObject(win32.HGDIOBJ(w.smIconMaskBitmap)) {
			logger.Println("Failed to destroy icon; DeleteObject(smIconMaskBitmap) failed!")
		}
	}

	iconWidth = win32.GetSystemMetrics(win32.SM_CXSMICON)
	iconHeight = win32.GetSystemMetrics(win32.SM_CYSMICON)

	iconImage = resize.Resize(w.icon, w.icon.Bounds(), int(iconWidth), int(iconHeight))

	iconBitmapInfo = win32.BITMAPINFO{
		BmiHeader: win32.BITMAPINFOHEADER{
			Size: win32.DWORD(unsafe.Sizeof(win32.BITMAPINFOHEADER{})),
			Width: win32.LONG(iconWidth),
			Height: win32.LONG(iconHeight),
			Planes: 1,
			BitCount: 32,
			Compression: win32.BI_RGB,
			SizeImage: 0,
			XPelsPerMeter: 0,
			YPelsPerMeter: 0,
			ClrUsed: 0,
			ClrImportant: 0,
		},
	}

	w.smIconColorBitmap = win32.CreateCompatibleBitmap(w.dc, iconWidth, iconHeight)
	w.smIconColorBits = make([]uint32, iconWidth * iconHeight)
	for y := 0; y < int(iconHeight); y++ {
		for x := 0; x < int(iconWidth); x++ {
			r, g, b, _ := iconImage.At(x, y).RGBA()
			c := (uint32(r >> 8) << 16) | (uint32(g >> 8) << 8) | (uint32(b >> 8) << 0)

			index := (int(iconHeight)-1 - y) * int(iconWidth)
			index += x
			w.smIconColorBits[index] = c //0xFF0000
		}
	}
	if win32.SetDIBits(w.dc, w.smIconColorBitmap, 0, win32.UINT(iconHeight), unsafe.Pointer(&w.smIconColorBits[0]), &iconBitmapInfo, win32.DIB_RGB_COLORS) == 0 {
		logger.Println("Unable to set icon; SetDiBits():", win32.GetLastErrorString())
		return
	}

	w.smIconMaskBitmap = win32.CreateCompatibleBitmap(w.dc, iconWidth, iconHeight)
	w.smIconMaskBits = make([]uint32, iconWidth * iconHeight)
	for y := 0; y < int(iconHeight); y++ {
		for x := 0; x < int(iconWidth); x++ {
			_, _, _, a := iconImage.At(x, y).RGBA()
			c := uint32(0xFFFFFF)
			if a > 0 {
				c = 0
			}

			index := (int(iconHeight)-1 - y) * int(iconWidth)
			index += x
			w.smIconMaskBits[index] = c
		}
	}
	if win32.SetDIBits(w.dc, w.smIconMaskBitmap, 0, win32.UINT(iconHeight), unsafe.Pointer(&w.smIconMaskBits[0]), &iconBitmapInfo, win32.DIB_RGB_COLORS) == 0 {
		logger.Println("Unable to set icon; SetDiBits():", win32.GetLastErrorString())
		return
	}

	iconInfo = win32.ICONINFO{
		FIcon: 1,
		XHotspot: 0,
		YHotspot: 0,
		HbmMask: w.smIconMaskBitmap,
		HbmColor: w.smIconColorBitmap,
	}

	w.hSmIcon = win32.CreateIconIndirect(&iconInfo)
	if w.hSmIcon == nil {
		logger.Println("Unable to set icon; CreateIconIndirect():", win32.GetLastErrorString())
		return
	}
	w.doSetIcon()
}

func (w *W32Window) doSetIcon() {
	win32.SendMessage(w.hwnd, win32.WM_SETICON, win32.ICON_BIG, win32.LPARAM(uintptr(unsafe.Pointer(w.hIcon))))
	win32.SendMessage(w.hwnd, win32.WM_SETICON, win32.ICON_SMALL, win32.LPARAM(uintptr(unsafe.Pointer(w.hSmIcon))))
}

func (w *W32Window) doSetCursor() {
	if w.cursorWithin {
		if w.cursorGrabbed {
			win32.SetCursor(nil)
			return
		}

		if w.loadedCursor == nil {
			if w.cursor != nil {
				// Cursor is just not prepared yet.
				w.doPrepareCursor(w.cursor)

				lc, ok := w.cursors[w.cursor]
				if ok {
					w.loadedCursor = lc
				}

			} else {
				// There is no cursor.
				lc := new(loadedCursor)
				lc.hCursor = win32.HICON(win32.LoadCursor(nil, win32.IDC_ARROW))
				if lc.hCursor == nil {
					logger.Println("Unable to load default (IDC_ARROW) cursor! LoadCursor():", win32.GetLastErrorString())
				} else {
					w.loadedCursor = lc
				}
			}
		}

		if w.loadedCursor != nil {
			win32.SetCursor(win32.HCURSOR(w.loadedCursor.hCursor))
		}
	}
}

func (w *W32Window) doSetCursorPos() {
	if w.cursorGrabbed {
		if !w.cursorWithin {
			return
		}
		w.cursorX = int(w.width / 2)
		w.cursorY = int(w.height / 2)
	}

	if !win32.SetCursorPos(int32(w.x+w.cursorX), int32(w.y+w.cursorY)) {
		logger.Println("Unable to set cursor position: SetCursorPos():", win32.GetLastErrorString())
	}
}

func (w *W32Window) doUpdateTransparency() {
	bb := win32.DWM_BLURBEHIND{}
	bb.DwFlags = win32.DWM_BB_ENABLE|win32.DWM_BB_BLURREGION
	if w.transparent {
		bb.FEnable = 1
	} else {
		bb.FEnable = 0
	}
	rgn := win32.CreateRectRgn(0, 0, -1, -1)
	bb.HRgbBlur = rgn
	err := win32.DwmEnableBlurBehindWindow(w.hwnd, &bb)
	if err != nil {
		logger.Println(err)
	}
	err = win32.DwmEnableBlurBehindWindow(w.hwndRender, &bb)
	if err != nil {
		logger.Println(err)
	}
}

func (w *W32Window) doUpdateStyle() {
	originalStyle := win32.GetWindowLongPtr(w.hwnd, win32.GWL_STYLE)

	if w.decorated && !w.fullscreen {
		w.styleFlags = win32.WS_OVERLAPPEDWINDOW
	} else {
		w.styleFlags = win32.WS_SYSMENU | win32.WS_POPUP | win32.WS_CLIPCHILDREN | win32.WS_CLIPSIBLINGS
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

	// Need to make the position relative to the original screen
	r := w.originalScreen.(*w32Screen).w32Position
	x = win32.Int(r.Left()) + x
	y = win32.Int(r.Top()) + y

	if !w.decorated {
		x += win32.Int(w.extentLeft)
		y += win32.Int(w.extentTop)
	}

	// Clip to maxWidth/maxHeight and minWidth/minHeight, append extents so that width/height is
	// the client region specifically.
	width := float64(w.width)
	if w.maxWidth > 0 {
		width = math.Min(float64(w.width), float64(w.maxWidth))
	}
	width = math.Max(width, float64(w.minWidth))

	height := float64(w.height)
	if w.maxHeight > 0 {
		height = math.Min(height, float64(w.maxHeight))
	}
	height = math.Max(height, float64(w.minHeight))

	if w.decorated {
		width += float64(w.extentLeft)
		width += float64(w.extentRight)
		height += float64(w.extentBottom)
		height += float64(w.extentTop)
	}

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
		r := w.screen.(*w32Screen).w32Position
		x = win32.Int(r.Left())
		y = win32.Int(r.Top())
		sm := w.screen.ScreenMode()
		screenWidth, screenHeight := sm.Resolution()
		width, height = float64(screenWidth), float64(screenHeight)

		if w.width != uint(screenWidth) || w.height != uint(screenHeight) {
			w.width = uint(screenWidth)
			w.height = uint(screenHeight)

			w.addSizeEvent([]uint{w.width, w.height})
		}
	}

	// |win32.SWP_NOZORDER|win32.SWP_NOOWNERZORDER
	if !win32.SetWindowPos(w.hwnd, insertAfter, x, y, win32.Int(width), win32.Int(height), win32.SWP_FRAMECHANGED) {
		logger.Println("Unable to set window position; SetWindowPos():", win32.GetLastErrorString())
	}
}


func (w *W32Window) translateKey(wParam win32.WPARAM) (key keyboard.Key) {
	switch(wParam) {
		// ? "Control-break processing" doesn't seem useful
		//case win32.VK_CANCEL:

		case win32.VK_BACK: key = keyboard.Backspace
		case win32.VK_TAB: key = keyboard.Tab
		case win32.VK_CLEAR: key = keyboard.Clear
		case win32.VK_RETURN: key = keyboard.Enter
		case win32.VK_PAUSE: key = keyboard.Pause
		case win32.VK_KANA: key = keyboard.Kana
		case win32.VK_JUNJA: key = keyboard.Junja
		case win32.VK_KANJI: key = keyboard.Kanji
		case win32.VK_ESCAPE: key = keyboard.Escape
		case win32.VK_SPACE: key = keyboard.Space
		case win32.VK_PRIOR: key = keyboard.PageUp
		case win32.VK_NEXT: key = keyboard.PageDown
		case win32.VK_END: key = keyboard.End
		case win32.VK_HOME: key = keyboard.Home
		case win32.VK_LEFT: key = keyboard.ArrowLeft
		case win32.VK_UP: key = keyboard.ArrowUp
		case win32.VK_RIGHT: key = keyboard.ArrowRight
		case win32.VK_DOWN: key = keyboard.ArrowDown
		case win32.VK_SELECT: key = keyboard.Select
		case win32.VK_PRINT: key = keyboard.Print
		case win32.VK_EXECUTE: key = keyboard.Execute
		case win32.VK_SNAPSHOT: key = keyboard.PrintScreen
		case win32.VK_INSERT: key = keyboard.Insert
		case win32.VK_DELETE: key = keyboard.Delete
		case win32.VK_HELP: key = keyboard.Help
		case win32.VK_UNDEF_0: key = keyboard.Zero
		case win32.VK_UNDEF_1: key = keyboard.One
		case win32.VK_UNDEF_2: key = keyboard.Two
		case win32.VK_UNDEF_3: key = keyboard.Three
		case win32.VK_UNDEF_4: key = keyboard.Four
		case win32.VK_UNDEF_5: key = keyboard.Five
		case win32.VK_UNDEF_6: key = keyboard.Six
		case win32.VK_UNDEF_7: key = keyboard.Seven
		case win32.VK_UNDEF_8: key = keyboard.Eight
		case win32.VK_UNDEF_9: key = keyboard.Nine
		case win32.VK_UNDEF_A: key = keyboard.A
		case win32.VK_UNDEF_B: key = keyboard.B
		case win32.VK_UNDEF_C: key = keyboard.C
		case win32.VK_UNDEF_D: key = keyboard.D
		case win32.VK_UNDEF_E: key = keyboard.E
		case win32.VK_UNDEF_F: key = keyboard.F
		case win32.VK_UNDEF_G: key = keyboard.G
		case win32.VK_UNDEF_H: key = keyboard.H
		case win32.VK_UNDEF_I: key = keyboard.I
		case win32.VK_UNDEF_J: key = keyboard.J
		case win32.VK_UNDEF_K: key = keyboard.K
		case win32.VK_UNDEF_L: key = keyboard.L
		case win32.VK_UNDEF_M: key = keyboard.M
		case win32.VK_UNDEF_N: key = keyboard.N
		case win32.VK_UNDEF_O: key = keyboard.O
		case win32.VK_UNDEF_P: key = keyboard.P
		case win32.VK_UNDEF_Q: key = keyboard.Q
		case win32.VK_UNDEF_R: key = keyboard.R
		case win32.VK_UNDEF_S: key = keyboard.S
		case win32.VK_UNDEF_T: key = keyboard.T
		case win32.VK_UNDEF_U: key = keyboard.U
		case win32.VK_UNDEF_V: key = keyboard.V
		case win32.VK_UNDEF_W: key = keyboard.W
		case win32.VK_UNDEF_X: key = keyboard.X
		case win32.VK_UNDEF_Y: key = keyboard.Y
		case win32.VK_UNDEF_Z: key = keyboard.Z

		//case win32.VK_CONVERT: key = keyboard.IMEConvert
		//case win32.VK_NONCONVERT: key = keyboard.IMENonConvert
		//case win32.VK_ACCEPT: key = keyboard.IMEAccept
		//case win32.VK_MODECHANGE: key = keyboard.IMEModeChange
		//case win32.VK_PROCESSKEY: key = keyboard.IMEProcess

		case win32.VK_LWIN: key = keyboard.LeftSuper
		case win32.VK_RWIN: key = keyboard.RightSuper
		case win32.VK_APPS: key = keyboard.Applications
		case win32.VK_SLEEP: key = keyboard.Sleep

		case win32.VK_NUMPAD0: key = keyboard.NumZero
		case win32.VK_NUMPAD1: key = keyboard.NumOne
		case win32.VK_NUMPAD2: key = keyboard.NumTwo
		case win32.VK_NUMPAD3: key = keyboard.NumThree
		case win32.VK_NUMPAD4: key = keyboard.NumFour
		case win32.VK_NUMPAD5: key = keyboard.NumFive
		case win32.VK_NUMPAD6: key = keyboard.NumSix
		case win32.VK_NUMPAD7: key = keyboard.NumSeven
		case win32.VK_NUMPAD8: key = keyboard.NumEight
		case win32.VK_NUMPAD9: key = keyboard.NumNine
		case win32.VK_MULTIPLY: key = keyboard.NumMultiply
		case win32.VK_ADD: key = keyboard.NumAdd
		case win32.VK_SEPARATOR: key = keyboard.NumComma
		case win32.VK_SUBTRACT: key = keyboard.NumSubtract
		case win32.VK_DECIMAL: key = keyboard.NumDecimal
		case win32.VK_DIVIDE: key = keyboard.NumDivide

		case win32.VK_F1: key = keyboard.F1
		case win32.VK_F2: key = keyboard.F2
		case win32.VK_F3: key = keyboard.F3
		case win32.VK_F4: key = keyboard.F4
		case win32.VK_F5: key = keyboard.F5
		case win32.VK_F6: key = keyboard.F6
		case win32.VK_F7: key = keyboard.F7
		case win32.VK_F8: key = keyboard.F8
		case win32.VK_F9: key = keyboard.F9
		case win32.VK_F10: key = keyboard.F10
		case win32.VK_F11: key = keyboard.F11
		case win32.VK_F12: key = keyboard.F12
		case win32.VK_F13: key = keyboard.F13
		case win32.VK_F14: key = keyboard.F14
		case win32.VK_F15: key = keyboard.F15
		case win32.VK_F16: key = keyboard.F16
		case win32.VK_F17: key = keyboard.F17
		case win32.VK_F18: key = keyboard.F18
		case win32.VK_F19: key = keyboard.F19
		case win32.VK_F20: key = keyboard.F20
		case win32.VK_F21: key = keyboard.F21
		case win32.VK_F22: key = keyboard.F22
		case win32.VK_F23: key = keyboard.F23
		case win32.VK_F24: key = keyboard.F24

		case win32.VK_BROWSER_BACK: key = keyboard.BrowserBack
		case win32.VK_BROWSER_FORWARD: key = keyboard.BrowserForward
		case win32.VK_BROWSER_REFRESH: key = keyboard.BrowserRefresh
		case win32.VK_BROWSER_STOP: key = keyboard.BrowserStop
		case win32.VK_BROWSER_SEARCH: key = keyboard.BrowserSearch
		case win32.VK_BROWSER_FAVORITES: key = keyboard.BrowserFavorites
		case win32.VK_BROWSER_HOME: key = keyboard.BrowserHome

		// User expects these to control windows volume -- I don't thing we should allow
		// intercepting these..
		//VK_VOLUME_MUTE
		//VK_VOLUME_DOWN
		//VK_VOLUME_UP

		case win32.VK_MEDIA_NEXT_TRACK: key = keyboard.MediaNext
		case win32.VK_MEDIA_PREV_TRACK: key = keyboard.MediaPrevious
		case win32.VK_MEDIA_STOP: key = keyboard.MediaStop
		case win32.VK_MEDIA_PLAY_PAUSE: key = keyboard.MediaPlayPause

		case win32.VK_LAUNCH_MAIL: key = keyboard.LaunchMail
		case win32.VK_LAUNCH_MEDIA_SELECT: key = keyboard.LaunchMedia
		case win32.VK_LAUNCH_APP1: key = keyboard.LaunchAppOne
		case win32.VK_LAUNCH_APP2: key = keyboard.LaunchAppTwo

		case win32.VK_OEM_PLUS: key = keyboard.Equals
		case win32.VK_OEM_COMMA: key = keyboard.Comma
		case win32.VK_OEM_MINUS: key = keyboard.Dash
		case win32.VK_OEM_PERIOD: key = keyboard.Period
		case win32.VK_OEM_1: key = keyboard.Semicolon
		case win32.VK_OEM_2: key = keyboard.ForwardSlash
		case win32.VK_OEM_3: key = keyboard.Tilde
		case win32.VK_OEM_4: key = keyboard.LeftBracket
		case win32.VK_OEM_5: key = keyboard.BackSlash
		case win32.VK_OEM_6: key = keyboard.RightBracket
		case win32.VK_OEM_7: key = keyboard.Apostrophe
		//case win32.VK_OEM_8:
		case win32.VK_OEM_102: key = keyboard.RightBracket

		case win32.VK_ATTN: key = keyboard.Attn
		case win32.VK_CRSEL: key = keyboard.CrSel
		case win32.VK_EXSEL: key = keyboard.ExSel
		case win32.VK_EREOF: key = keyboard.EraseEOF
		case win32.VK_PLAY: key = keyboard.Play
		case win32.VK_ZOOM: key = keyboard.Zoom
		//case win32.VK_PA1: key = keyboard.PA1
		case win32.VK_OEM_CLEAR: key = keyboard.Clear

		case win32.VK_SHIFT:
			leftShiftDown := (uint16(win32.GetAsyncKeyState(win32.VK_LSHIFT)) & 0x8000) != 0

			if leftShiftDown != w.leftShiftDown {
				key = keyboard.LeftShift
			} else {
				key = keyboard.RightShift
			}

			w.leftShiftDown = leftShiftDown

		case win32.VK_MENU:
			leftAltDown := (uint16(win32.GetAsyncKeyState(win32.VK_LMENU)) & 0x8000) != 0

			if leftAltDown != w.leftAltDown {
				key = keyboard.LeftAlt
			} else {
				key = keyboard.RightAlt
			}

			w.leftAltDown = leftAltDown

		case win32.VK_CONTROL:
			leftCtrlDown := (uint16(win32.GetAsyncKeyState(win32.VK_LCONTROL)) & 0x8000) != 0

			if leftCtrlDown != w.leftCtrlDown {
				key = keyboard.LeftCtrl
			} else {
				key = keyboard.RightCtrl
			}

			w.leftCtrlDown = leftCtrlDown

		case win32.VK_CAPITAL: key = keyboard.CapsLock
		case win32.VK_NUMLOCK: key = keyboard.NumLock
		case win32.VK_SCROLL: key = keyboard.ScrollLock
	}
	return key
}

// Our MS windows event handler
func mainWindowProc(hwnd win32.HWND, msg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) (ret win32.LRESULT) {
	w, ok := windowsByHwnd[hwnd]
	if ok {
		switch {
		case msg == win32.WM_PAINT:
			if win32.GetUpdateRect(w.hwnd, nil, false) {
				win32.ValidateRect(w.hwnd, nil)
				w.addPaintEvent()
			}
			return 0

		case msg == win32.WM_GETMINMAXINFO:
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
				minMaxInfo.PtMinTrackSize().SetX(win32.LONG(newMinWidth))
			}
			if minHeight > 0 {
				minMaxInfo.PtMinTrackSize().SetY(win32.LONG(newMinHeight))
			}

			if maxWidth > 0 {
				minMaxInfo.PtMaxTrackSize().SetX(win32.LONG(newMaxWidth))
			}
			if maxHeight > 0 {
				minMaxInfo.PtMaxTrackSize().SetY(win32.LONG(newMaxHeight))
			}
			return 0

		case msg == win32.WM_SIZING:
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

		case msg == win32.WM_SIZE:
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

		case msg == win32.WM_MOVE:
			xPos := int(int16(lParam))
			yPos := int(int16((uint32(lParam) >> 16) & 0xFFFF))

			if !win32.IsIconic(w.hwnd) {
				// Clamp when it goes onto an monitor to the left.. very unsure how to handle
				// window/screen interaction -- it's never very cross platform..

				if w.x != xPos || w.y != yPos {
					w.x = xPos
					w.y = yPos
					w.addPositionEvent([]int{w.x, w.y})
				}
			}

			return 0

		case msg == win32.WM_EXITSIZEMOVE:
			hMonitor := win32.MonitorFromWindow(w.hwnd, win32.MONITOR_DEFAULTTONEAREST)

			mi := new(win32.MONITORINFOEX)
			mi.SetSize()
			if !win32.GetMonitorInfo(hMonitor, mi) {
				logger.Println("Unable to detect monitor position; GetMonitorInfo():", win32.GetLastErrorString())
			} else {
				screens := backend_doScreens()
				for _, screen := range screens {
					if screen.(*w32Screen).w32GraphicsDeviceName == mi.Device() {
						screen.(*w32Screen).w32Position = mi.RcMonitor

						w.screen = screen
						w.addScreenChangedEvent(w.screen)
						return 0
					}
				}
			}

			return 0

		case msg == win32.WM_ACTIVATE:
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

		case msg == win32.WM_GETICON:
			switch(wParam) {
				case win32.ICON_BIG:
					if w.hIcon != nil {
						return win32.LRESULT(uintptr(unsafe.Pointer(w.hIcon)))
					}

				case win32.ICON_SMALL:
					if w.hSmIcon != nil {
						return win32.LRESULT(uintptr(unsafe.Pointer(w.hSmIcon)))
					}

				case win32.ICON_SMALL2:
					if w.hSmIcon != nil {
						return win32.LRESULT(uintptr(unsafe.Pointer(w.hSmIcon)))
					}
			}

		case msg == win32.WM_KEYDOWN || msg == win32.WM_SYSKEYDOWN || msg == win32.WM_KEYUP || msg == win32.WM_SYSKEYUP:
			if msg == win32.WM_KEYDOWN || msg == win32.WM_SYSKEYDOWN {
				keyRepeat := (lParam & 0x40000000) > 0
				if keyRepeat {
					return 0
				}
			}

			if (msg == win32.WM_SYSKEYDOWN || msg == win32.WM_SYSKEYUP) && wParam == win32.VK_F4 {
				altDown := (uint16(win32.GetAsyncKeyState(win32.VK_MENU)) & 0x8000) != 0
				if altDown {
					if msg == win32.WM_SYSKEYDOWN {
						// Trick: Consider this to be WM_CLOSE
						if !w.addCloseEvent() {
							go w.Destroy()
						}
					}
					return 0
				}
			}

			k := w.translateKey(wParam)
			if k == keyboard.Invalid {
				logger.Printf("Unknown key wParam=%d hex=0x%X)\n", wParam, wParam)
			}

			var state keyboard.State
			if msg == win32.WM_KEYDOWN || msg == win32.WM_SYSKEYDOWN {
				state = keyboard.Down
			} else {
				state = keyboard.Up
			}

			switch(k) {
				case keyboard.CapsLock:
					if (win32.GetKeyState(win32.VK_CAPITAL) & 0x0001) != 0 {
						state = keyboard.On
					} else {
						state = keyboard.Off
					}

				case keyboard.NumLock:
					if (win32.GetKeyState(win32.VK_NUMLOCK) & 0x0001) != 0 {
						state = keyboard.On
					} else {
						state = keyboard.Off
					}

				case keyboard.ScrollLock:
					if (win32.GetKeyState(win32.VK_SCROLL) & 0x0001) != 0 {
						state = keyboard.On
					} else {
						state = keyboard.Off
					}
			}

			w.addKeyboardEvent(&keyboard.Event{
				Key: k,
				OSKey: keyboard.OSKey(lParam),
				Unicode: "",
				State: state,
			})
			return 0

		case msg == win32.WM_MOUSEMOVE:
			xPos := float64(int16(lParam))
			yPos := float64(int16((uint32(lParam) >> 16) & 0xFFFF))

			if float64(w.cursorX) != xPos || float64(w.cursorY) != yPos {
				if xPos < 0 {
					xPos = 0
				}
				if yPos < 0 {
					yPos = 0
				}
				w.cursorX = int(xPos)
				w.cursorY = int(yPos)

				if !w.cursorGrabbed {
					w.addCursorPositionEvent([]float64{float64(w.cursorX), float64(w.cursorY)})
				}
			}

			if w.cursorX >= int(w.width) || w.cursorY >= int(w.height) || w.cursorX <= 0 || w.cursorY <= 0 || !w.focused {
				// Better than WM_MOUSELEAVE
				if w.cursorWithin {
					if !w.cursorGrabbed {
						win32.ReleaseCapture()
					}

					w.cursorWithin = false
					w.addCursorWithinEvent(w.cursorWithin)
				}
			} else {
				// Closest we'll get to WM_MOUSEENTER
				if !w.cursorWithin {
					win32.SetCapture(w.hwnd)

					w.cursorWithin = true
					w.addCursorWithinEvent(w.cursorWithin)
				}
			}

			if w.cursorGrabbed {
				supportRawInput := w32VersionMajor >= 5 && w32VersionMinor >= 1
				halfWidth := int(w.width / 2)
				halfHeight := int(w.height / 2)

				if w.cursorX != halfWidth || w.cursorY != halfHeight {
					if !supportRawInput {
						// If we have no support for raw mouse input, then we need to fall back to finding
						// the mouse movement on our own.
						diffX := w.cursorX - halfWidth
						diffY := w.cursorY - halfHeight
						if diffX != 0 || diffY != 0 {
							w.addCursorPositionEvent([]float64{float64(diffX), float64(diffY)})
						}
					}
					w.doSetCursorPos()
				}
			}
			w.doSetCursor()
			return 0

		case msg == win32.WM_INPUT:
			if w.cursorWithin && w.cursorGrabbed {
				var raw win32.RAWINPUT
				cbSize := win32.UINT(unsafe.Sizeof(raw))

				win32.GetRawInputData((win32.HRAWINPUT)(unsafe.Pointer(uintptr(lParam))), win32.RID_INPUT, unsafe.Pointer(&raw), &cbSize, win32.UINT(unsafe.Sizeof(win32.RAWINPUTHEADER{})))

				if raw.Header().Type == win32.RIM_TYPEMOUSE {
					diffX := raw.Mouse().LastX()
					diffY := raw.Mouse().LastY()
					if diffX != 0 || diffY != 0 {
						w.addCursorPositionEvent([]float64{float64(diffX), float64(diffY)})
					}
				}
			}
			return 0

		// Mouse Buttons
		case msg == win32.WM_LBUTTONDOWN:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Left,
				State:  mouse.Down,
			})
			return 0

		case msg == win32.WM_LBUTTONUP:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Left,
				State:  mouse.Up,
			})
			return 0

		case msg == win32.WM_RBUTTONDOWN:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Right,
				State:  mouse.Down,
			})
			return 0

		case msg == win32.WM_RBUTTONUP:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Right,
				State:  mouse.Up,
			})
			return 0

		case msg == win32.WM_MBUTTONDOWN:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Wheel,
				State:  mouse.Down,
			})
			return 0

		case msg == win32.WM_MBUTTONUP:
			w.addMouseEvent(&mouse.Event{
				Button: mouse.Wheel,
				State:  mouse.Up,
			})
			return 0

		case msg == win32.WM_XBUTTONDOWN:
			var button mouse.Button

			switch(int16(wParam)) {
				case win32.MK_XBUTTON1:
					button = mouse.Four

				case win32.MK_XBUTTON2:
					button = mouse.Five
			}

			w.addMouseEvent(&mouse.Event{
				Button: button,
				State:  mouse.Down,
			})
			return 0

		case msg == win32.WM_XBUTTONUP:
			var button mouse.Button

			switch(int16(wParam)) {
				case win32.MK_XBUTTON1:
					button = mouse.Four

				case win32.MK_XBUTTON2:
					button = mouse.Five
			}

			w.addMouseEvent(&mouse.Event{
				Button: button,
				State:  mouse.Up,
			})
			return 0

		case msg == win32.WM_MOUSEWHEEL:
			delta := float64(int16((uint32(wParam) >> 16) & 0xFFFF))
			ticks := int(math.Abs(delta / 120))

			if delta > 0 {
				for i := 0; i < ticks; i++ {
					w.addMouseEvent(&mouse.Event{
						Button: mouse.Wheel,
						State:  mouse.ScrollForward,
					})
				}
			} else {
				for i := 0; i < ticks; i++ {
					w.addMouseEvent(&mouse.Event{
						Button: mouse.Wheel,
						State:  mouse.ScrollBack,
					})
				}
			}
			return 0

		//default:
		//	fmt.Printf("0x%x\n", msg)

		case msg == win32.WM_CLOSE:
			if !w.addCloseEvent() {
				go w.Destroy()
			}
			return 0
		}
	}

	return win32.DefWindowProc(hwnd, msg, wParam, lParam)
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

