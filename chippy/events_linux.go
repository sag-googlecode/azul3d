package chippy

import "runtime"
import "fmt"

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

func handleEvent(ev *c_XEvent) {
    switch ev._type() {
        case c_PropertyNotify:
            e := ev.xproperty()
            w := findWindow(c_Window(e.window))
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
                //    fmt.Println("PROPERTY", c_XGetAtomName(xDisplay, properties[i]))
                //}
            }

        // Window Close
        case c_ClientMessage:
            e := ev.xclient()
            w := findWindow(c_Window(e.window))

            if c_isDeleteWindowAtom(xDisplay, e) {
                addWindowCloseEvent(w)
            }

        // Redraw window
        case c_Expose:
            e := ev.xexpose()
            w := findWindow(c_Window(e.window))
            addWindowRedrawEvent(w)

        // Window restored/non-minimized
        case c_MapNotify:
            e := ev.xmap()
            w := findWindow(c_Window(e.window))
            setWindowMinimized(w, false)

        // Window Minimized
        case c_UnmapNotify:
            e := ev.xunmap()
            w := findWindow(c_Window(e.window))
            setWindowMinimized(w, true)

        // Window gained focus
        case c_FocusIn:
            e := ev.xfocus()
            w := findWindow(c_Window(e.window))

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
            if int16(e.x) != w.x || int16(e.y) != w.y {
                w.x = int16(e.x)
                w.y = int16(e.y)
                w.access.Unlock()
                addWindowMoveEvent(w)
            } else {
                w.access.Unlock()
            }

        // Cursor movement
        case c_MotionNotify:
            e := ev.xmotion()
            w := findWindow(c_Window(e.window))
            w.access.Lock()
            if uint16(e.x) != w.cursorX || uint16(e.y) != w.cursorY {
                w.cursorX = uint16(e.x)
                w.cursorY = uint16(e.y)
                w.access.Unlock()
                addCursorMoveEvent(w)
            } else {
                w.access.Unlock()
            }

        // Cursor enter/exit
        case c_EnterNotify:
            e := ev.xcrossing()
            w := findWindow(c_Window(e.window))
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
            w.access.Lock()
            if w.cursorWithin != false {
                w.cursorWithin = false
                w.access.Unlock()
                addCursorWithinEvent(w)
            } else {
                w.access.Unlock()
            }

        case c_GenericEvent:
            cookie := ev.xcookie()
            if c_XGetEventData(xDisplay, cookie) && cookie.extension == XI_opcode {
                defer c_XFreeEventData(xDisplay, cookie)
                xiEvent := cookie.XIDeviceEvent()

                switch int32(xiEvent.evtype) {
                    case c_XI_RawMotion:
                        x, y := float64(xiEvent.event_x), float64(xiEvent.event_y)
                        // Sanity check, just in case
                        if x != 0.0 || y != 0.0 {
                            // We send this to all windows, as there is no window-specific event here at all
                            for _, w := range cWindowMapping {
                                addMouseMoveEvent(w, x, y)
                            }
                        }

                    default:
                        fmt.Println("chippy: Unknown XIDeviceEvent:", xiEvent.evtype)
                }

            } else {
                fmt.Println("chippy: Unknown GenericEvent", cookie.extension)
            }

        default:
            fmt.Println("chippy: Unknown XEvent:", ev._type())
    }
}

func eventDispatcher() {
    for eventsDispatching {
        runtime.Gosched()

        // Ensure that at this point, all X11 calls where flushed and executed
        c_XSync(xDisplay, false)

        // Wait for an event, or timeout (allowing other goroutines to run via runtime.Gosched)
        event := c_waitForEvent(xDisplay, 0, 100000)
        // If there was an event, then we never timed out waiting for one
        if event {
            // We need to handle every event available
            for c_XPending(xDisplay) > 0 {
                // Grab the next event, and handle it
                ev := c_XNextEvent(xDisplay)
                handleEvent(ev)
            }
        }
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

