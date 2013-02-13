package chippy

import(
    "code.google.com/p/azul3d/chippy/wrappers/win32"
    "errors"
    "fmt"
)

var windowsByHwnd = make(map[win32.HWND]*Window)

type backend_Window struct {
    w32Hwnd win32.HWND
    w32WindowClass string
}

// WindowsHWND returns the handle to this Window (HWND), this is of course, Windows specific, and
// is only useful in less ways than others.
func (w *Window) WindowsHWND() win32.HWND {
    return w.w32Hwnd
}
 
func mainWindowProc(hwnd win32.HWND, msg win32.UINT, wParam win32.WPARAM, lParam win32.LPARAM) (ret win32.LRESULT) {
    // Warning! Never place win32.Sync***** calls inside here, or you'll end in deadlock!

    w, ok := windowsByHwnd[hwnd]
    if ok {
        switch(msg) {
            case win32.WM_GETMINMAXINFO:
                w.access.RLock()
                minWidth, minHeight := w.minWidth, w.minHeight
                maxWidth, maxHeight := w.maxWidth, w.maxHeight
                extentLeft, extentRight, extentBottom, extentTop := w.extentLeft, w.extentRight, w.extentBottom, w.extentTop
                w.access.RUnlock()

                // Set maximum and minimum window sizes, 0 means unlimited
                minMaxInfo := lParam.MINMAXINFO()

                if minWidth > 0 {
                    minMaxInfo.PtMinTrackSize().SetX(int32(minWidth + (extentLeft + extentRight)))
                }
                if minHeight > 0 {
                    minMaxInfo.PtMinTrackSize().SetY(int32(minHeight + (extentBottom + extentTop)))
                }

                if maxWidth > 0 {
                    minMaxInfo.PtMaxTrackSize().SetX(int32(maxWidth + (extentLeft + extentRight)))
                }
                if maxHeight > 0 {
                    minMaxInfo.PtMaxTrackSize().SetY(int32(maxHeight + (extentBottom + extentTop)))
                }
        }

        //logger.Println("mainWindowProc()", msg, win32.WM_SIZE, win32.WM_SIZING)
    }

    ret = win32.DefWindowProc(hwnd, msg, wParam, lParam)
    return
}

func (w *Window) open() (err error) {
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
        w.w32WindowClass = fmt.Sprintf("ChippyWindow%d", nextCounter())
        windowClass := win32.NewWNDCLASSEX()
        windowClass.SetLpfnWndProc()
        //windowClass.SetHIcon(win32.LoadIcon(hInstance, szAppName))
        //windowClass.SetHCursor(win32.LoadCursor(nil, win32.IDC_ARROW))
        //windowClass.SetHbrBackground(win32.IntToHBRUSH(win32.COLOR_WINDOW+2)) // Black background
        //windowClass.SetLpszMenuName(szAppName)

        windowClass.SetHInstance(hInstance)
        windowClass.SetLpszClassName(w.w32WindowClass)

        classAtom := win32.RegisterClassEx(windowClass)
        if classAtom == 0{
            err = errors.New(fmt.Sprintf("Unable to open window; RegisterClassEx(): %s", win32.GetLastErrorString()))
            return
        }


        // Handle min/max sizes, also CreateWindowEx() takes window region coordinates, not client region
        width := w.width
        if width > w.maxWidth {
            width = w.maxWidth
        }
        if width < w.minWidth {
            width = w.minWidth
        }
        width = width + (w.extentLeft + w.extentRight)

        height := w.height
        if height > w.maxHeight {
            height = w.maxHeight
        }
        if height < w.minHeight {
            height = w.minHeight
        }
        height = height + (w.extentBottom + w.extentTop)


        style := win32.DWORD(0)
        if w.maximized {
            style |= win32.WS_MAXIMIZE
        }
        if w.minimized {
            style |= win32.WS_MINIMIZE
        }
        style |= win32.WS_OVERLAPPEDWINDOW

        w.w32Hwnd = win32.CreateWindowEx(0, w.w32WindowClass, w.title, style, win32.Int(w.x), win32.Int(w.y), win32.Int(width), win32.Int(height), nil, nil, hInstance, nil)
        if w.w32Hwnd == nil {
            err = errors.New(fmt.Sprintf("Unable to open window; CreateWindowEx(): %s", win32.GetLastErrorString()))
            return
        }
        windowsByHwnd[w.w32Hwnd] = w
        win32.RegisterWndProc(w.w32Hwnd, mainWindowProc)

        if w.visible {
            win32.ShowWindowAsync(w.w32Hwnd, win32.SW_SHOW)
        }

        if !w.decorated && w.visible {
            style := win32.GetWindowLongPtr(w.w32Hwnd, win32.GWL_STYLE)
            style &^= win32.WS_OVERLAPPEDWINDOW
            win32.SetWindowLongPtr(w.w32Hwnd, win32.GWL_STYLE, style)

            if !win32.SetWindowPos(w.w32Hwnd, nil, 0, 0, 0, 0, win32.SWP_ASYNCWINDOWPOS|win32.SWP_FRAMECHANGED|win32.SWP_NOMOVE|win32.SWP_NOSIZE|win32.SWP_NOZORDER|win32.SWP_NOOWNERZORDER) {
                logger.Println("SetWindowPos():", win32.GetLastErrorString())
            }
        }
    })
    return
}

func (w *Window) close() {
    win32.UnregisterWndProc(w.w32Hwnd)
    delete(windowsByHwnd, w.w32Hwnd)

    dispatch(func() {
        if !win32.DestroyWindow(w.w32Hwnd) {
            logger.Println("Unable to destroy window; DestroyWindow():", win32.GetLastErrorString())
        }

        if !win32.UnregisterClass(w.w32WindowClass, hInstance) {
            logger.Println("UnregisterClass() failed!")
        }
    })
}

func (w *Window) handleSetGetEvent(ev uint) {
    var style win32.LONG_PTR
    dispatch(func() {
        style = win32.GetWindowLongPtr(w.w32Hwnd, win32.GWL_STYLE)
    })
    original := style

    if ev == opSetDecorated {
        if w.decorated && w.visible {
            style |= win32.WS_OVERLAPPEDWINDOW

        } else if !w.decorated && w.visible {
            style &^= win32.WS_OVERLAPPEDWINDOW
        }
    }

    if ev == opSetMinimized || ev == opSetMaximized || ev == opSetVisible {
        flag := win32.Int(win32.SW_HIDE)

        if w.visible {
            if w.minimized {
                flag = win32.SW_MINIMIZE
            } else if w.maximized {
                flag = win32.SW_MAXIMIZE
            } else if (!w.minimized || !w.maximized) && (ev == opSetMinimized || ev == opSetMaximized) {
                flag = win32.SW_RESTORE
            } else {
                flag = win32.SW_SHOW
            }
        }
        dispatch(func() {
            win32.ShowWindowAsync(w.w32Hwnd, flag)
        })
    }

    width := w.width
    if width > w.maxWidth {
        width = w.maxWidth
    }
    if width < w.minWidth {
        width = w.minWidth
    }

    height := w.height
    if height > w.maxHeight {
        height = w.maxHeight
    }
    if height < w.minHeight {
        height = w.minHeight
    }

    dispatch(func() {
        if style != original {
            win32.SetWindowLongPtr(w.w32Hwnd, win32.GWL_STYLE, style)
            if win32.GetLastError() != 0 {
                logger.Println("SetWindowLongPtr():", win32.GetLastErrorString())
            }

            win32.EnableWindow(w.w32Hwnd, true)

            if !win32.SetWindowPos(w.w32Hwnd, nil, 0, 0, 0, 0, win32.SWP_ASYNCWINDOWPOS|win32.SWP_FRAMECHANGED|win32.SWP_NOMOVE|win32.SWP_NOSIZE|win32.SWP_NOZORDER|win32.SWP_NOOWNERZORDER) {
                logger.Println("SetWindowPos():", win32.GetLastErrorString())
            }
        }

        if !win32.SetWindowPos(w.w32Hwnd, nil, win32.Int(w.x), win32.Int(w.y), win32.Int(w.x + int(w.width)), win32.Int(w.y + int(w.height)), win32.SWP_ASYNCWINDOWPOS|win32.SWP_FRAMECHANGED|win32.SWP_NOZORDER|win32.SWP_NOOWNERZORDER) {
            logger.Println("SetWindowPos():", win32.GetLastErrorString())
        }
    })
    fmt.Println(opString(ev))
}

