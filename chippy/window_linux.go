package chippy

import(
    "errors"
    "image"
)

type backend_frameBufferConfig struct {
	actual c_GLXFBConfig
}

func backend_frameBufferConfigs(screen *Screen) ([]*FrameBufferConfig, error) {
    return c_getFrameBufferConfigs(xDisplay, screen.xScreenNumber)
}

type backend_window struct {
	colormap c_Colormap
	window   c_Window
    origCursorX, origCursorY uint16
    xGrabPointerWorked bool
    xCursor c_Cursor
}

var cWindowMapping = map[c_Window]*Window{}

func findWindow(w c_Window) *Window {
	return cWindowMapping[w]
}

func (w *Window) create() error {
	vi := c_glXGetVisualFromFBConfig(xDisplay, w.frameBufferConfig.actual)
	if vi == nil {
		return errors.New("Unexpected FrameBufferConfig is invalid! glXGetVisualFromFBConfig() failed!") // Should never happen
	}

	parent := c_XRootWindow(xDisplay, int32(vi.screen))

	cmap := c_XCreateColormap(xDisplay, parent, (*c_Visual)(vi.visual), c_AllocNone)
	defer c_XFreeColormap(xDisplay, cmap)

	swa := c_XSetWindowAttributes{}
	swa.colormap = cmap.C()
	swa.background_pixmap = c_Pixmap(c_None).C()
	swa.border_pixel = 0
	swa.event_mask = c_long(c_StructureNotifyMask)

	w.window = c_XCreateWindow(xDisplay, parent, int32(w.x), int32(w.y), uint32(w.width), uint32(w.height), 0, int32(vi.depth), c_InputOutput, (*c_Visual)(vi.visual), c_CWBorderPixel|c_CWColormap|c_CWEventMask, &swa)
	if w.window == 0 {
		return errors.New("Failed to create window; XCreateWindow() failed!")
	}
	cWindowMapping[w.window] = w

	// Accept delete events
	wmDelete := c_XInternAtom(xDisplay, "WM_DELETE_WINDOW", true)
	c_XSetWMProtocols(xDisplay, w.window, &wmDelete, 1)

	// Set up our event mask here
	c_setXIEventMask(xDisplay)

	c_XSelectInput(xDisplay, w.window,
		c_KeyPressMask|c_KeyReleaseMask|
		c_ButtonPressMask|c_ButtonReleaseMask|
		c_EnterWindowMask|c_LeaveWindowMask|
		c_PointerMotionMask|c_ExposureMask|
		c_FocusChangeMask|c_StructureNotifyMask|
		c_PropertyChangeMask)


	//c_ExposureMask | c_FocusChangeMask | c_StructureNotifyMask | c_PropertyChangeMask)

	//|c_Button1MotionMask|c_Button2MotionMask|c_Button3MotionMask|c_Button4MotionMask|c_Button5MotionMask|c_ButtonMotionMask|c_KeymapStateMask|c_ExposureMask|c_VisibilityChangeMask|c_StructureNotifyMask|c_ResizeRedirectMask|c_SubstructureNotifyMask|c_SubstructureRedirectMask|c_FocusChangeMask|c_PropertyChangeMask|c_ColormapChangeMask|c_OwnerGrabButtonMask)

    // Important, update our cursorX and cursorY to the actual cursor position here, if it's inside the window
    found, x, y := c_getPointerPosition(xDisplay, w.window)

    // Clip it to the window's region
    if x < 0 {
        x = 0
    }
    if y < 0 {
        y = 0
    }

    if x > int(w.width) {
        x = int(w.width)
    }

    if y > int(w.height) {
        y = int(w.height)
    }

    if found {
        w.cursorX = uint16(x)
        w.cursorY = uint16(y)
    } else {
        // Probably the best assumption we can make in this case..
        w.cursorX = w.width / 2
        w.cursorY = w.height / 2
    }



	w.setTitle()
	w.setSize()
	w.setMaxSize()
	w.setMinSize()
	w.setPosition()
	w.setVerticalSync()
	w.setVisible()
	w.setDecorated()
	w.setFullscreen()
	return nil
}

func (w *Window) contextVersion() (uint8, uint8, uint8, error) {
	return 0, 0, 0, nil
}

func (w *Window) swapBuffers() error {
	c_glXSwapBuffers(xDisplay, w.window)
	return nil
}

func (w *Window) makeCurrent() error {
	return nil
}

func (w *Window) setTitle() error {
	//err := c_XStoreName(xDisplay, w.window, w.title)
	//if err != nil {
	//    return err
	//}
	c_Xutf8SetWMProperties(xDisplay, w.window, w.title, w.title, "", 0, nil, nil, nil)
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setSize() error {
	width := w.width
	height := w.height

    // Check for max
    if w.maxWidth > 0 && width > w.maxWidth {
        width = w.maxWidth
    }
    if w.maxHeight > 0 && height > w.maxHeight {
        height = w.maxHeight
    }

    // Check for min
    if w.minHeight > 0 && width < w.minWidth {
        width = w.minWidth
    }
    if w.minHeight > 0 && height < w.minHeight {
        height = w.minHeight
    }

	err := c_XResizeWindow(xDisplay, w.window, uint32(width), uint32(height))
	if err != nil {
		return err
	}
	c_XSync(xDisplay, false)
	return nil
}


// X11 specific!
func (w *Window) updateSizeHints() error {
    var err error

    if w.resizable == false {
        err = c_setSizeHints(xDisplay, w.window, w.width, w.height, w.width, w.height)
    } else {
        err = c_setSizeHints(xDisplay, w.window, w.minWidth, w.minHeight, w.maxWidth, w.maxHeight)
    }
    if err != nil {
        return err
    }

    // The window may already be larger than the max size, or smaller than the min size
    err = w.setSize()
    if err != nil {
        return err
    }

	c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setMaxSize() error {
    return w.updateSizeHints()
}

func (w *Window) setMinSize() error {
    return w.updateSizeHints()
}

func (w *Window) setPosition() error {
	x := int32(w.x)
	y := int32(w.y)

	err := c_XMoveWindow(xDisplay, w.window, x, y)
	if err != nil {
		return err
	}
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setVerticalSync() error {
	return nil
}

func (w *Window) setVisible() error {
	if w.visible {
        w.setIcon()

		err := c_XMapWindow(xDisplay, w.window)
		if err != nil {
			return err
		}
        w.raise()
	} else {
        c_XUnmapWindow(xDisplay, w.window)
		//err := c_XWithdrawWindow(xDisplay, w.window, w.screen.xScreenNumber)
		//if err != nil {
		//	return err
		//}
	}

	c_XSync(xDisplay, false)

    // At this point we can query the window extents
    extents := c_getWindowExtents(xDisplay, w.window)
    w.extentLeft, w.extentRight, w.extentTop, w.extentBottom = uint8(extents[0]), uint8(extents[1]), uint8(extents[2]), uint8(extents[3])

	// XMapWindow removes the size/position of the window, so
    // we set them back here (and hence no XSync here)
	//err := w.setSize()
	err := w.setPosition()
    //err = w.setMaxSize()
    //err = w.setMinSize()
	if err != nil {
		return err
	}
	return nil
}

func (w *Window) setDecorated() error {
	hints := c_Hints{}
	hints.flags = 2 // Specify that we're changing the window decorations.
	if w.decorated {
		hints.decorations = 1 // Decorations on
	} else {
		hints.decorations = 0 // Decorations off
	}
	property := c_XInternAtom(xDisplay, "_MOTIF_WM_HINTS", true)
	if property == 0 {
		return errors.New("Unable to remove window decorations; _MOTIF_WM_HINTS not supported!")
	}
	err := c_XChangeProperty(xDisplay, w.window, property, property, 32, c_PropModeReplace, c_Pointer(&hints), 5)
	if err != nil {
		return err
	}

	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setMaximized() error {
	c_setAtomState(xDisplay, w.window, "_NET_WM_STATE_MAXIMIZED_HORZ", w.maximized)
	c_setAtomState(xDisplay, w.window, "_NET_WM_STATE_MAXIMIZED_VERT", w.maximized)
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setMinimized() error {
	if w.minimized {
		err := c_XIconifyWindow(xDisplay, w.window, w.screen.xScreenNumber)
		if err != nil {
			return err
		}
	} else {
		// It appears most X window managers will refuse to restore an previously minimized
		// application window, unless the user requests it themself.
		// We do this here because some window managers may choose to allow it, and this
		// appears to be the only way to request that the window be restored from an iconified
		// state
		err := c_XMapWindow(xDisplay, w.window)
		if err != nil {
			return err
		}
	}
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setAlwaysOnTop() error {
	c_setAtomState(xDisplay, w.window, "_NET_WM_STATE_ABOVE", w.alwaysOnTop)
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setMouseGrabbed() error {
    // We're responsible for restoring the cursor position from before we forced it into the center of the window

    if w.mouseGrabbed {
        // reset it first
        w.xGrabPointerWorked = false

        // Store the original cursor position, *before* we moved it, w.cursorX and w.cursorY
        // are clipped to the Window's region, and we'd *prefer* to avoid that clipping, so
        // we will use the non-clipped version here.
        found, x, y := c_getPointerPosition(xDisplay, w.window)

        if found {
            w.origCursorX = uint16(x)
            w.origCursorY = uint16(y)
        } else {
            // Otherwise use the last clipped cursor position
            w.origCursorX = w.cursorX
            w.origCursorY = w.cursorY
        }

        // Now try and grab the cursor, as long as the cursor is currently in the window
        if w.cursorWithin {
            if c_XGrabPointer(xDisplay, w.window, true,
                    uint(c_ButtonPressMask|c_ButtonReleaseMask|c_PointerMotionMask), c_GrabModeAsync,
                    c_GrabModeAsync, w.window, 0, c_CurrentTime) == c_GrabSuccess {

                w.xGrabPointerWorked = true
            }
        }
    } else {
        c_XUngrabPointer(xDisplay, c_CurrentTime)

        // Restore the original cursor position, the one we stored *before* we moved it
        // *only* do this if the window has focus, otherwise we may be moving the cursor
        // against the user's will
        if w.focus && !w.minimized {
            w.cursorX = w.origCursorX
            w.cursorY = w.origCursorY
            err := w.setCursorPosition()
            if err != nil {
                return err
            }
        }
    }

	// Then in the event loop we query Window.mouseGrabbed when the mouse is moved
	return nil
}

func (w *Window) notify() error {
	c_setAtomState(xDisplay, w.window, "_NET_WM_STATE_DEMANDS_ATTENTION", true)
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) raise() error {
	c_XRaiseWindow(xDisplay, w.window)
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setResizable() error {
    return w.updateSizeHints()
}

func (w *Window) setCursorPosition() error {
	c_XWarpPointer(xDisplay, 0, w.window, 0, 0, 0, 0, int32(w.cursorX), int32(w.cursorY))
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setCursorHidden() error {
    if w.cursorHidden {
        // We use 32x32 here since it's most likely to be supported (I wonder heh?)
        image := image.NewRGBA(image.Rect(0, 0, 32, 32))
        cursor := c_xCursorFromImage(xDisplay, image, uint(w.cursorIconPointerX), uint(w.cursorIconPointerY))

        c_XDefineCursor(xDisplay, w.window, cursor)
    } else {
        if w.xCursor != 0 {
            // There could be a cursor currently set, so use that one
            c_XDefineCursor(xDisplay, w.window, w.xCursor)
        } else {
            // Otherwise we restore the original WM one
            c_XDefineCursor(xDisplay, w.window, 0)
        }
    }
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setCursorIcon() error {
    if w.cursorIcon == nil {
        // Set the cursor to the original WM one
        c_XDefineCursor(xDisplay, w.window, 0)
        w.xCursor = 0
    } else {
        w.xCursor = c_xCursorFromImage(xDisplay, w.cursorIcon, uint(w.cursorIconPointerX), uint(w.cursorIconPointerY))

        // If the cursor is hidden, we strictly want to avoid setting it
        if !w.cursorHidden {
            c_XDefineCursor(xDisplay, w.window, w.xCursor)
        }
    }
    c_XSync(xDisplay, false)
	return nil
}

func (w *Window) setIcon() error {
    if w.icon != nil {
        err := c_setWindowIcon(xDisplay, w.window, w.icon)
        if err != nil {
            return err
        }
    }

    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setFullscreen() error {
	err := c_XRaiseWindow(xDisplay, w.window) // This raises the window above all others, or tries to anyway
	if err != nil {
		return err
	}
	c_setAtomState(xDisplay, w.window, "_NET_WM_STATE_FULLSCREEN", w.fullscreen)
	c_XSync(xDisplay, false)
	return nil
}

func (w *Window) destroy() {
	delete(cWindowMapping, w.window)

    // important! Turn back on keyboard auto repeat!
    c_XAutoRepeatOn(xDisplay)

	c_XDestroyWindow(xDisplay, w.window)
}
