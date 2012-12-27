package chippy

import(
    "code.google.com/p/azul3d/chippy/keyboard"
    //"code.google.com/p/azul3d/chippy/mouse"

    "runtime"
)

var eventsDispatching bool

func setWindowMaximized(w *Window, maximized bool) {
	w.access.Lock()
	if w.maximized != maximized {
		w.maximized = maximized
		w.access.Unlock()
		addWindowMaximizedEvent(w)
	} else {
		w.access.Unlock()
	}
}

func setWindowMinimized(w *Window, minimized bool) {
    if w.Destroyed() {
        panic("destroyed")
    }

	w.access.Lock()
	if w.minimized != minimized {
		w.minimized = minimized
		w.access.Unlock()
		addWindowMinimizedEvent(w)
	} else {
		w.access.Unlock()
	}
}

func setWindowAlwaysOnTop(w *Window, alwaysOnTop bool) {
	w.access.Lock()
	if w.alwaysOnTop != alwaysOnTop {
		w.alwaysOnTop = alwaysOnTop
		w.access.Unlock()
		addWindowAlwaysOnTopEvent(w)
	} else {
		w.access.Unlock()
	}
}

func propertiesContainsName(props []c_Atom, name string) bool {
	for i := 0; i < len(props); i++ {
		if c_XGetAtomName(xDisplay, props[i]) == name {
			return true
		}
	}
	return false
}

func handleEvent(ev *c_XEvent, cookie *c_XGenericEventCookie) {
	switch ev._type() {
	case c_PropertyNotify:
		e := ev.xproperty()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

		name := c_XGetAtomName(xDisplay, c_Atom(e.atom))
		if name == "_NET_WM_STATE" {
			properties := c_GetNetWmStates(xDisplay, w.window)
			if propertiesContainsName(properties, "_NET_WM_STATE_MAXIMIZED_VERT") && propertiesContainsName(properties, "_NET_WM_STATE_MAXIMIZED_HORZ") {
				// Maximized
				setWindowMaximized(w, true)
			} else {
				// Non-Maximized
				setWindowMaximized(w, false)
			}

			if propertiesContainsName(properties, "_NET_WM_STATE_HIDDEN") {
				// Minimized
				setWindowMinimized(w, true)
			} else {
				// Non-Minimized
				setWindowMinimized(w, false)
			}

			if propertiesContainsName(properties, "_NET_WM_STATE_ABOVE") {
				setWindowAlwaysOnTop(w, true)
			} else {
				setWindowAlwaysOnTop(w, false)
			}

			//for i := 0; i < len(properties); i++ {
			//    logger.Println("PROPERTY", c_XGetAtomName(xDisplay, properties[i]))
			//}
		}

    case c_ReparentNotify:
        // No need to do anything with this right now

	// Window Close
	case c_ClientMessage, c_DestroyNotify:
		e := ev.xclient()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

		if c_isDeleteWindowAtom(xDisplay, e) {
			addWindowCloseEvent(w)
		}

	// Redraw window
	case c_Expose:
		e := ev.xexpose()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }
		addWindowRedrawEvent(w)

	// Window restored/non-minimized
	case c_MapNotify:
		e := ev.xmap()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }
		setWindowMinimized(w, false)

	// Window Minimized
	case c_UnmapNotify:
		e := ev.xunmap()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }
		setWindowMinimized(w, true)

	// Window gained focus
	case c_FocusIn:
		e := ev.xfocus()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

        // important! Disable auto repeat of keys
        c_XAutoRepeatOff(xDisplay)

		w.access.Lock()
		if w.focus != true {
			w.focus = true
			w.access.Unlock()
			addWindowFocusEvent(w)
		} else {
			w.access.Unlock()
		}

	// Window lost focus
	case c_FocusOut:
		e := ev.xfocus()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

        // important! Turn back on keyboard auto repeat!
        c_XAutoRepeatOn(xDisplay)

		w.access.Lock()
		if w.focus != false {
			w.focus = false
			w.access.Unlock()
			addWindowFocusEvent(w)
		} else {
			w.access.Unlock()
		}

	// Window configuration changed
	case c_ConfigureNotify:
		e := ev.xconfigure()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

		// Check if the size of the window changed
		w.access.Lock()
		if uint16(e.width) != w.width || uint16(e.height) != w.height {
			w.width = uint16(e.width)
			w.height = uint16(e.height)
			w.access.Unlock()
			addWindowResizeEvent(w)
		} else {
			w.access.Unlock()
		}

		// Check if the position of the window changed
		w.access.Lock()

        x := int16(e.x)
        y := int16(e.y)

        if w.decorated {
            // Take into account the window decorations, since the move events *do not* include window decorations
            x -= int16(w.extentLeft)
            y -= int16(w.extentTop)
        }

        if x != w.x || y != w.y {
	        w.x = x
	        w.y = y
		    w.access.Unlock()
	        addWindowMoveEvent(w)
        } else {
	        w.access.Unlock()
        }

	// Cursor movement
	case c_MotionNotify:
		e := ev.xmotion()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

		w.access.Lock()
	    if uint16(e.x) != w.cursorX || uint16(e.y) != w.cursorY {
            if w.mouseGrabbed {
                // It's possible that they requested mouse grab but we where denied it, so here we
                // will request it again since it's likely we can have it if the cursor is inside
                // of this Window's region.
                if w.xGrabPointerWorked == false {
                    if c_XGrabPointer(xDisplay, w.window, true,
                            uint(c_ButtonPressMask|c_ButtonReleaseMask|c_PointerMotionMask), c_GrabModeAsync,
                            c_GrabModeAsync, w.window, 0, c_CurrentTime) == c_GrabSuccess {

                        w.xGrabPointerWorked = true

                        // Since we *actually* grabbed the mouse here, store the cursor position as the original one
                        w.origCursorX = uint16(e.x)
                        w.origCursorY = uint16(e.y)

                        // It can be annoying to place the cursor on the exact edge of the window, because most
                        // WM's make this the resize window area, causing for an user easilly resizing the window
                        // when they never wanted to in the first place
                        if w.origCursorX == 0 {
                            w.origCursorX = 1
                        }

                        if w.origCursorY == 0 {
                            w.origCursorY = 1
                        }

                        if w.origCursorX == w.width {
                            w.origCursorX = w.width - 1
                        }

                        if w.origCursorY == w.width {
                            w.origCursorY = w.width - 1
                        }
                    }
                }
                w.access.Unlock()
            } else {
                // Send movement event
		        w.cursorX = uint16(e.x)
		        w.cursorY = uint16(e.y)
                w.access.Unlock()
		        addCursorMoveEvent(w)
            }
	    } else {
		    w.access.Unlock()
	    }

	// Cursor enter/exit
	case c_EnterNotify:
		e := ev.xcrossing()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

		w.access.Lock()
		if w.cursorWithin != true {
			w.cursorWithin = true
			w.access.Unlock()
			addCursorWithinEvent(w)
		} else {
			w.access.Unlock()
		}
	case c_LeaveNotify:
		e := ev.xcrossing()
		w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

		w.access.Lock()
		if w.cursorWithin != false {
			w.cursorWithin = false
			w.access.Unlock()
			addCursorWithinEvent(w)
		} else {
			w.access.Unlock()
		}

    case c_KeyPress:
        e := ev.xkey()
        w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

        keySym := c_getKeySym(e)
        keyCode := c_keysym2ucs(keySym)

        code := c_findKeyboardCode(keySym)
        r := rune(0)
        if keyCode != -1 {
            r = rune(keyCode)
        }

        button := &keyboard.Button{r, code, keyboard.Pressed, int(keySym)}
        addKeyEvent(w, button)

        logger.Println(button)

    case c_KeyRelease:
        e := ev.xkey()
        w := findWindow(c_Window(e.window))
        if w == nil {
            break
        }

        keySym := c_getKeySym(e)
        keyCode := c_keysym2ucs(keySym)

        code := c_findKeyboardCode(keySym)
        r := rune(0)
        if keyCode != -1 {
            r = rune(keyCode)
        }

        button := &keyboard.Button{r, code, keyboard.Released, int(keySym)}
        addKeyEvent(w, button)

        logger.Println(button)

	case c_GenericEvent:
		if cookie.extension == xI_opcode {
			xiEvent := cookie.XIDeviceEvent()

			switch int(xiEvent.evtype) {
			    case c_XI_RawMotion:
				    //fmt.Println(xiEvent.detail, xiEvent.flags)
				    x, y := float64(xiEvent.event_x), float64(xiEvent.event_y)
				    // Sanity check, just in case
				    if x != 0.0 || y != 0.0 {
					    // We send this to all windows, as there is no window-specific event here at all
					    for _, w := range cWindowMapping {
						    addMouseMoveEvent(w, x, y)
					    }
				    }

			    default:
				    logger.Println("Unknown XIDeviceEvent:", xiEvent.evtype)
			}
		}

	default:
		logger.Println("Unknown XEvent:", ev._type())
	}
}


func eventDispatcher() {
	runtime.GOMAXPROCS(4)

	for eventsDispatching {
		runtime.Gosched()

		//runtime.GC()
		//runtime.LockOSThread()

		// Ensure that at this point, all X11 calls where flushed and executed
		c_XSync(xDisplay, false)

		// Wait for an event, or timeout (allowing other goroutines to run via runtime.Gosched)
		event := c_waitForEvent(xDisplay, 0, 100000)

		// If there was an event, then we never timed out waiting for one
        // It's possible to wait for an event and get one even though we're not dispatching
        // so check here also
		if event && eventsDispatching {

			// We need to handle every event available
			for c_XPending(xDisplay) > 0 {
				ev := c_XNextEvent(xDisplay)
				cookie := ev.xcookie()

				c_XGetEventData(xDisplay, cookie) // Make sure this is called *before* the next call to XNextEvent
				handleEvent(ev, cookie)
				c_XFreeEventData(xDisplay, cookie)
			}
		}
		//runtime.UnlockOSThread()
	}
}

func initEvents() error {
	eventsDispatching = true
	go eventDispatcher()
	return nil
}

func destroyEvents() {
	eventsDispatching = false
}
