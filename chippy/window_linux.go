package chippy

import "errors"
import "sync"

// FBConfig represents options to configure the Frame Buffer at window creation time
type FBConfig struct {
    RedBits, GreenBits, BlueBits, AlphaBits,
    AccumRedBits, AccumGreenBits, AccumBlueBits, AccumAlphaBits,
    DepthBits, StencilBits, Samples, SampleBuffers, AuxBuffers uint8
    DoubleBuffered, StereoScopic bool

    // Platform
    actual c_GLXFBConfig
}

var cWindowMapping = map[c_Window]*Window{}

func findWindow(w c_Window) *Window {
    return cWindowMapping[w]
}

// Window represents a visual window that the user will see and that you will render to
type Window struct {
    screen *Screen
    minAttribs, maxAttribs, fbConfig *FBConfig

    destroyed, vsync, decorated, visible, minimized, fullscreen, maximized, cursorWithin,
    cursorHidden, focus, alwaysOnTop, mouseGrabbed,
    keyEvents, cursorWithinEvents, cursorMoveEvents, mouseMoveEvents, resizeEvents,
    moveEvents, focusEvents, minimizedEvents, maximizedEvents, redrawEvents,
    alwaysOnTopEvents, closeEvents bool

    title string
    width, height, cursorX, cursorY uint16
    x, y int16
    destroyCallback *callback
    access sync.RWMutex

    // Platform
    colormap c_Colormap
    window c_Window
}

func (w *Window) create() error {
    w.fbConfig = c_chooseFBConfig(xDisplay, w.screen.xScreenNumber, w.minAttribs, w.maxAttribs)
    if w.fbConfig == nil {
        return errors.New("Unable to retrieve a matching FBConfig")
    }
    vi := c_glXGetVisualFromFBConfig(xDisplay, w.fbConfig.actual)
    if vi == nil {
        return errors.New("Unexpected FBConfig is invalid! glXGetVisualFromFBConfig() failed!")
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

    c_XSelectInput(xDisplay, w.window, 
        c_KeyPressMask | c_KeyReleaseMask |
        c_ButtonPressMask | c_ButtonReleaseMask |
        c_EnterWindowMask | c_LeaveWindowMask |
        c_PointerMotionMask | c_ExposureMask |
        c_FocusChangeMask | c_StructureNotifyMask |
        c_PropertyChangeMask)

    //c_ExposureMask | c_FocusChangeMask | c_StructureNotifyMask | c_PropertyChangeMask)

        //|c_Button1MotionMask|c_Button2MotionMask|c_Button3MotionMask|c_Button4MotionMask|c_Button5MotionMask|c_ButtonMotionMask|c_KeymapStateMask|c_ExposureMask|c_VisibilityChangeMask|c_StructureNotifyMask|c_ResizeRedirectMask|c_SubstructureNotifyMask|c_SubstructureRedirectMask|c_FocusChangeMask|c_PropertyChangeMask|c_ColormapChangeMask|c_OwnerGrabButtonMask)

    c_setXIEventMask(xDisplay)

    w.setTitle()
    w.setSize()
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
    width := uint32(w.width)
    height := uint32(w.height)

    err := c_XResizeWindow(xDisplay, w.window, width, height)
    if err != nil {
        return err
    }
    c_XSync(xDisplay, false)
    return nil
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
        err := c_XMapWindow(xDisplay, w.window)
        if err != nil {
            return err
        }
    } else {
        err := c_XWithdrawWindow(xDisplay, w.window, w.screen.xScreenNumber)
        if err != nil {
            return err
        }
    }
    // It appears either XMapWindow or XWithdrawWindow (or both)
    // remove the size and/or position of the window so we set
    // it back right here (hence no XSync)
    err := w.setSize()
    err = w.setPosition()
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
    err := c_XChangeProperty(xDisplay, w.window, property, property, 32, c_PropModeReplace, c_Pointer(&hints), 5);
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
        err := c_XRaiseWindow(xDisplay, w.window)
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
    // We just query Window.mouseGrabbed when we want to, so we will ignore this event
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

func (w *Window) setCursorPosition() error {
    c_XWarpPointer(xDisplay, 0, w.window, 0, 0, 0, 0, int32(w.cursorX), int32(w.cursorY))
    c_XSync(xDisplay, false)
    return nil
}

func (w *Window) setCursorHidden() error {
    // We will query Window.cursorHidden when we want to, so we will ignore this event
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
    c_XDestroyWindow(xDisplay, w.window)
}

