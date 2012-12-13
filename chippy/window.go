package chippy

import "errors"
import "fmt"

// BestFBConfig represents the best possible Frame Buffer Configuration, you will
// never actually get an FBConfig with these properties. Pass this into NewWindow()
// as the maxAttribs parameter and you will always be given the *best* possible 
// FBConfig that is available.
var BestFBConfig = &FBConfig{
    RedBits: 255, GreenBits: 255, BlueBits: 255, AlphaBits: 255,
    AccumRedBits: 255, AccumGreenBits: 255, AccumBlueBits: 255, AccumAlphaBits: 255,
    DepthBits: 255, StencilBits: 255, Samples: 255, SampleBuffers: 255, AuxBuffers: 255,
    DoubleBuffered: true, StereoScopic: true,
}

// String returns a string representation of this FBConfig
func (f *FBConfig) String() string {
    return fmt.Sprintf("FBConfig(RedBits=%d, greenBits=%d, BlueBits=%d, AlphaBits=%d, AccumRedBits=%d, AccumGreenBits=%d, AccumBlueBits=%d, AccumAlphaBits=%d, DepthBits=%d, StencilBits=%d, Samples=%d, SampleBuffers=%d, AuxBuffers=%d, DoubleBuffered=%t, StereoScopic=%t)", f.RedBits, f.GreenBits, f.BlueBits, f.AlphaBits, f.AccumRedBits, f.AccumGreenBits, f.AccumBlueBits, f.AccumAlphaBits, f.DepthBits, f.StencilBits, f.Samples, f.SampleBuffers, f.AuxBuffers, f.DoubleBuffered, f.StereoScopic)
}

// Use this to avoid an OS call to MakeCurrent in case someone calls it multiple times while being silly
var currentContext *Window

// CurrentContext returns whichever Window last has MakeCurrent() called on it
func CurrentContext() (*Window, error) {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    return currentContext, nil
}

// NewWindow returns a new, open window that the user will see visually on their screen.
//
// This window will be capable of at least minAttribs, and as close as possible to maxAttribs.
//
// If you simply want the *best* possible frame buffer configuration available, pass in
// chippy.BestFBConfig as maxAttribs here, and you will always recieve the *best* possible
// frame buffer configuration available.
func NewWindow(screen *Screen, minAttribs, maxAttribs *FBConfig) (*Window, error) {
    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return nil, err
    }

    w := Window{}
    w.destroyCallback = &callback{func(){
            w.Destroy()
        }}
    addDestroyCallback(w.destroyCallback)

    w.screen = screen
    w.minAttribs = minAttribs
    w.maxAttribs = maxAttribs
    w.vsync = true
    w.visible = true
    w.decorated = true
    w.title = "Chippy Window"
    w.width = 640
    w.height = 480
    w.x = int16((screen.Width() / 2) - (w.width / 2))
    w.y = int16((screen.Height() / 2) - (w.height / 2))

    err = w.create()
    if err != nil {
        return nil, err
    }
    return &w, nil
}

// String returns a string representation of this Window
func (w *Window) String() string {
    w.access.RLock()
    defer w.access.RUnlock()
    return fmt.Sprintf("Window(\"%s\", size=%dx%d, pos=%dx%d, cursorPos=%dx%d, cursorWithin=%t, vsync=%t, focus=%t, visible=%t, minimized=%t, maximized=%t, decorated=%t, fullscreen=%t, destroyed=%t)", w.title, w.width, w.height, w.x, w.y, w.cursorX, w.cursorY, w.cursorWithin, w.vsync, w.focus, w.visible, w.minimized, w.maximized, w.decorated, w.fullscreen, w.destroyed)
}

// FBConfig returns the frame buffer configuration that this Window is using, this is
// the frame buffer configuration that was found to be matching according to when you
// called NewWindow() specifiying the minAttribs and maxAttribs parameters
func (w *Window) FBConfig() (*FBConfig, error) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return nil, errors.New("Window already had Destroy() called!")
    }
    return w.fbConfig, nil
}

// Screen returns the Screen that this Window was created on
func (w *Window) Screen() (*Screen, error) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return nil, errors.New("Window already had Destroy() called!")
    }
    return w.screen, nil
}

// MakeCurrent makes this Window's OpenGL context current, the current OpenGL context.
//
// This function is thread safe, but do note that OpenGL is state based
// calling this function makes this OpenGL context the current context within
// the current OS thread only, so use runtime.LockOSThread and runtime.UnlockOSThread
// appropriately as needed when interfacing with this.
func (w *Window) MakeCurrent() error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if currentContext != w {
        currentContext = w

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.makeCurrent()
        if err != nil {
            return err
        }
    }
    return nil
}

// SwapBuffers swaps the background and foreground buffers of this Window.
//
// You should always call this on Windows with DoubleBuffered FBConfigs,
// specifically after you are done executing your render code for the frame.
//
// This function is no-op if this Window has no DoubleBuffered FBConfig.
func (w *Window) SwapBuffers() error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if w.fbConfig.DoubleBuffered {
        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.swapBuffers()
        if err != nil {
            return err
        }
    }
    return nil
}

// ContextVersion returns the major, minor, and revision versions of this OpenGL context. For example: [1, 2, 1], or [2, 0, 0]
//
// See also ContextVersionString() for another useful alternative to this function
func (w *Window) ContextVersion() (uint8, uint8, uint8, error) {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        return 0, 0, 0, errors.New("Window already had Destroy() called!")
    }

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return 0, 0, 0, err
    }

    return w.contextVersion()
}

// ContextVersionString returns the ContextVersion but as a string, for example: "1.2.1" or "2.0"
//
// Note that the last ('revision') version is only ommited (in the case of "2.0" vs "1.2.1") if
// the revision version is zero. The minor version is never ommited, in the case of "2.0" for instance
func (w *Window) ContextVersionString() (string, error) {
    major, minor, revision, err := w.ContextVersion()
    if err != nil {
        return "", err
    }
    if revision == 0 {
        return fmt.Sprintf("%d.%d", major, minor), nil
    }
    return fmt.Sprintf("%d.%d.%d", major, minor, revision), nil
}

// SetTitle sets the title string of this Window, this can be any valid UTF-8 Go string
func (w *Window) SetTitle(title string) error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        errors.New("Window already had Destroy() called!")
    }
    if title != w.title {
        w.title = title

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setTitle()
        if err != nil {
            return err
        }
    }
    return nil
}

// Title returns the title string of this Window, this is any valid UTF-8 Go string
func (w *Window) Title() string {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.title
}

// SetWidth sets the width (in pixels) of this Window
func (w *Window) SetWidth(width uint16) error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        errors.New("Window already had Destroy() called!")
    }

    if width != w.width {
        w.width = width

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setSize()
        if err != nil {
            return err
        }
    }
    return nil
}

// Width returns the width (in pixels) of this Window
func (w *Window) Width() uint16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.width
}

// SetHeight sets the height (in pixels) of this Window
func (w *Window) SetHeight(height uint16) error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        errors.New("Window already had Destroy() called!")
    }

    if height != w.height {
        w.height = height

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setSize()
        if err != nil {
            return err
        }
    }
    return nil
}

// Height returns the height (in pixels) of this Window
func (w *Window) Height() uint16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.height
}

// SetSize sets the width and height of this Window (in pixels)
func (w *Window) SetSize(width, height uint16) error {
    err := w.SetWidth(width)
    err = w.SetHeight(height)
    if err != nil {
        return err
    }
    return nil
}

// Size returns the width and height of this Window (in pixels)
func (w *Window) Size() (uint16, uint16) {
    return w.Width(), w.Height()
}

// SetX sets the x position of this Window (in pixels), relative to the top left corner of the screen
func (w *Window) SetX(x int16) error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        errors.New("Window already had Destroy() called!")
    }

    if x != w.x {
        w.x = x

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setPosition()
        if err != nil {
            return err
        }
    }
    return nil
}

// X returns the x position of this Window (in pixels), relative to the top left corner of the screen
func (w *Window) X() int16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.x
}

// SetY sets the y position of this Window (in pixels), relative to the top left corner of the screen
func (w *Window) SetY(y int16) error {
    w.access.RLock()
    defer w.access.RUnlock()
    if w.destroyed {
        errors.New("Window already had Destroy() called!")
    }

    if y != w.y {
        w.y = y

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setPosition()
        if err != nil {
            return err
        }
    }
    return nil
}

// Y returns the y position of this Window (in pixels), relative to the top left corner of the screen
func (w *Window) Y() int16 {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.y
}

// SetPosition sets the x and y position of this Window (in pixels), relative to the top left corner of the screen
func (w *Window) SetPosition(x, y int16) error {
    err := w.SetX(x)
    err = w.SetY(y)
    if err != nil {
        return err
    }
    return nil
}

// SetPositionCenter sets this window to the center of Window.Screen()
//
// This call is equivilent to
//  Window.SetPosition(int16((Window.Screen().Width() / 2) - (Window.Width() / 2)), int16((Window.Screen().Height() / 2) - (Window.Height() / 2)))
func (w *Window) SetPositionCenter() error {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.SetPosition(int16((w.screen.Width() / 2) - (w.Width() / 2)), int16((w.screen.Height() / 2) - (w.Height() / 2)))
}

// Pos returns the x and y position of this Window (in pixels), relative to the top left corner of the screen
func (w *Window) Position() (int16, int16) {
    return w.X(), w.Y()
}

// SetVerticalSync sets weather vertical sync (vsync) will be on or off on this Window
//
// If true, this Window will try to match the vertical refresh rate of Window.Screen
func (w *Window) SetVerticalSync(vsync bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        errors.New("Window already had Destroy() called!")
    }

    if vsync != w.vsync {
        w.vsync = vsync

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setVerticalSync()
        if err != nil {
            return err
        }
    }
    return nil
}

// VerticalSync returns weather this Window has vertical sync (vsync) enabled
func (w *Window) VerticalSync() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.vsync
}

// SetHidden sets weather the window is visible or hidden, a hidden window
// will appear to the user as simply gone.
func (w *Window) SetVisible(visible bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if visible != w.visible {
        w.visible = visible

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setVisible()
        if err != nil {
            return err
        }
    }
    return nil
}

// Visible returns weather the window is visible or hidden
func (w *Window) Visible() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.visible
}

// SetMaximized sets weather this window is maximized (true) or un-maximized/windowed (false)
func (w *Window) SetMaximized(maximized bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if maximized != w.maximized {
        w.maximized = maximized

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setMaximized()
        if err != nil {
            return err
        }
    }
    return nil
}

// Maximized returns weather the window is currently maximized
func (w *Window) Maximized() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.maximized
}


// Minimize minimizes this window, hiding it from the user while still allow the user to retreive it
func (w *Window) SetMinimized(minimized bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if minimized != w.minimized {
        w.minimized = minimized

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setMinimized()
        if err != nil {
            return err
        }
    }
    return nil
}

// Minimized returns weather the window is currently minimized
func (w *Window) Minimized() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.minimized
}

// SetDecoration sets weather the window will have window decorations
func (w *Window) SetDecorated(decorated bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if decorated != w.decorated {
        w.decorated = decorated

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setDecorated()
        if err != nil {
            return err
        }
    }
    return nil
}

// Decorated returns weather the window is decorated
func (w *Window) Decorated() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.decorated
}

// SetAlwaysOnTop specifies that this Window should always be shown to the user above all others
func (w *Window) SetAlwaysOnTop(alwaysOnTop bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if alwaysOnTop != w.alwaysOnTop {
        w.alwaysOnTop = alwaysOnTop

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setAlwaysOnTop()
        if err != nil {
            return err
        }
    }
    return nil
}

// AlwaysOnTop tells weather this Window is set to be always shown on top of other windows to the user
func (w *Window) AlwaysOnTop() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.alwaysOnTop
}

// SetMouseGrabbed specifies that the mouse cursor should be hidden (via SetCursorHidden(true)),
// and stuck inside of this Window, that this Window should no longer send CursorMoveEvent's and
// now begin sending MouseMoveEvent's
func (w *Window) SetMouseGrabbed(mouseGrabbed bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if mouseGrabbed != w.mouseGrabbed {
        w.mouseGrabbed = mouseGrabbed

        w.access.Unlock()
        err := w.SetCursorHidden(mouseGrabbed) // If they grab the mouse, hide the cursor, if they show the mouse, show the cursor
        if err != nil {
            return err
        }
        w.access.Lock()

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err = getInitError()
        if err != nil {
            return err
        }

        err = w.setMouseGrabbed()
        if err != nil {
            return err
        }
    }
    return nil
}

// MouseGrabbed tells weather this Window currently has the mouse grabbed
func (w *Window) MouseGrabbed() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.mouseGrabbed
}

// SetCursorPosition sets the position of the cursor relative to the top left corner of this Window
// in pixel coordinates, for example, X=50, Y=50 places the cursor 50 pixels right of the top left
// corner of the Window, and 50 pixels down from the top left corner of the Window
// It is impossible to move the cursor further than the Window's Width() or Height(), in the case
// that you try to do this, it will be moved to Width() and Height()
func (w *Window) SetCursorPosition(x, y uint16) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    // We ensure that the placement position is no larger than this Window's Width() or Height()
    // here to provide consistent cursor placement across all platforms (for instance this might
    // work on platform X but be unavailable on platform Y)
    if x > w.width {
        x = w.width
    }
    if y > w.height {
        y = w.height
    }

    if x != w.cursorX && y != w.cursorY {
        w.cursorX = x
        w.cursorY = y

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setCursorPosition()
        if err != nil {
            return err
        }
    }
    return nil
}

// CursorPosition returns the position of the cursor relative to the top left corner of this Window
// in pixel coordinates, for example, X=50, Y=50 means the cursor is 50 pixels right of the top left
// corner of the Window, and 50 pixels down from the top left corner of the Window
func (w *Window) CursorPosition() (uint16, uint16) {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.cursorX, w.cursorY
}

// CursorWithin tells weather the cursor cursor is within this Window's region or outside of it
func (w *Window) CursorWithin() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.cursorWithin
}

// SetCursorHidden sets weather the Window has the cursor hidden while within this Window's region
func (w *Window) SetCursorHidden(cursorHidden bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if cursorHidden != w.cursorHidden {
        w.cursorHidden = cursorHidden

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setCursorHidden()
        if err != nil {
            return err
        }
    }
    return nil
}

// CursorHidden returns weather the mouse cursor is currently hidden while within this Window's region
func (w *Window) CursorHidden() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.cursorHidden
}

// SetFullscreen sets weather the window will be fullscreen
func (w *Window) SetFullscreen(fullscreen bool) error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    if fullscreen != w.fullscreen {
        w.fullscreen = fullscreen

        // Calling into C -- Get the lock
        chippyAccess.Lock()
        defer chippyAccess.Unlock()

        err := getInitError()
        if err != nil {
            return err
        }

        err = w.setFullscreen()
        if err != nil {
            return err
        }
    }
    return nil
}

// Fullscreen returns weather the window is fullscreen
func (w *Window) Fullscreen() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.fullscreen
}

// Focus tells weather this window, or some other window has focus
func (w *Window) Focus() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.focus
}

// Notify tells the user that something important happened, typically
// the window manager displays this as a flashing icon in the tray or
// something like that
func (w *Window) Notify() error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return w.notify()
}

// Raise raises this window above all other windows
func (w *Window) Raise() error {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return errors.New("Window already had Destroy() called!")
    }

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return w.raise()
}

// Destroy destroys the window, closing it visually
// It's only needed to call this if you want to close the window right away,
// otherwise the final call to chippy.Destroy() will call this for you.
func (w *Window) Destroy() {
    w.access.Lock()
    defer w.access.Unlock()
    if w.destroyed {
        return // Window is already destroyed
    }

    removeDestroyCallback(w.destroyCallback)

    // Calling into C -- Get the lock
    chippyAccess.Lock()
    defer chippyAccess.Unlock()
    if w == currentContext {
        currentContext = nil
    }
    w.destroy()
    w.destroyed = true
}

// Destroyed returns weather the window has been Destroyed, via a call to Destroy() on the Window
func (w *Window) Destroyed() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    return w.destroyed
}

