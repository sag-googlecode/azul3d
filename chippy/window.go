package chippy

import(
    "code.google.com/p/azul3d/chippy/resize"
    "errors"
    "image"
    "sync"
    "fmt"
)

// Use this to avoid an OS call to MakeCurrent in case someone calls it multiple times while being silly
var currentContext *Window

// CurrentContext returns whichever Window last had MakeCurrent() called on it, or nil if no Window
// has yet to have MakeCurrent() called on it.
func CurrentContext() (*Window, error) {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	err := getInitError()
	if err != nil {
		return nil, err
	}

	return currentContext, nil
}

// Window represents a visual window that the user will see and that you will display graphics on
type Window struct {
    backend_window
	screen                           *Screen
	frameBufferConfig *FrameBufferConfig

	destroyed, vsync, decorated, visible, minimized, fullscreen, maximized, cursorWithin,
	cursorHidden, focus, alwaysOnTop, mouseGrabbed, resizable,
	keyEvents, mouseButtonEvents, cursorWithinEvents, cursorMoveEvents, mouseMoveEvents,
    resizeEvents, moveEvents, focusEvents, minimizedEvents, maximizedEvents, redrawEvents,
	alwaysOnTopEvents, closeEvents bool

	width, height, cursorX, cursorY, maxWidth, maxHeight, minWidth, minHeight uint16

    extentLeft, extentRight, extentBottom, extentTop uint8

	title                                  string
	x, y                                   int16
    cursorIcon, icon                       image.Image
    cursorIconPointerX, cursorIconPointerY uint8
	destroyCallback                        *callback
	access                                 sync.RWMutex

    keyEventsChan               chan *KeyEvent
    mouseButtonEventsChan               chan *MouseButtonEvent
    cursorWithinEventsChan      chan *CursorWithinEvent
    mouseMoveEventsChan         chan *MouseMoveEvent
    cursorMoveEventsChan        chan *CursorMoveEvent
    windowResizeEventsChan      chan *WindowResizeEvent
    windowMoveEventsChan        chan *WindowMoveEvent
    windowFocusEventsChan       chan *WindowFocusEvent
    windowMinimizedEventsChan   chan *WindowMinimizedEvent
    windowMaximizedEventsChan   chan *WindowMaximizedEvent
    windowRedrawEventsChan      chan *WindowRedrawEvent
    windowAlwaysOnTopEventsChan chan *WindowAlwaysOnTopEvent
    windowCloseEventsChan       chan *WindowCloseEvent
}

// NewWindow returns a new, open window that the user will see visually on their screen.
//
//
// The config parameter represents the new Window's frame buffer configuration, it should be
// an valid FrameBufferConfig returned by FrameBufferConfigs(), specifying one that you've made
// yourself is strictly dissallowed, it *must* have came from FrameBufferConfigs() at some point.
//
// If you simply want the *best* possible frame buffer configuration available, you may use the
// ChooseConfig() function, like so:
//
//  configs := chippy.FrameBufferConfigs()
//  chosenConfig := chippy.ChooseConfig(configs, chippy.WorstConfig, chippy.BestConfig)
//  window := NewWindow(..., chosenConfig, ...)
//
// And the Window will be capable of at least chippy.WorstConfig and at max be capable of chippy.BestConfig
//
//
// The icon parameter should be either nil (no Window icon), or an 128x128 image with 8-bit transparency.
//
// The icon parameter specifies the Image that will be this Window's icon, typically this is
// shown on the system's 'task bar', 'window switcher' or equivilent.
//
// This icon should be an 128x128 image (with 8 bit transparency) and it will be resized as it
// is appropriate for the platform. If the image specified for the icon parameter is anything
// other than 128x128, it will be automatically resized to 128x128, which may cause the image
// to become blurry (if the image is smaller than 128x128) or become stretched (if non-square).
func NewWindow(screen *Screen, config *FrameBufferConfig, icon image.Image) (*Window, error) {
    if config == nil {
        return nil, errors.New("Invalid config parameter; config = nil")
    }

	// Calling into C -- Get the lock
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	err := getInitError()
	if err != nil {
		return nil, err
	}

	w := Window{}
	w.destroyCallback = &callback{func() {
		w.Destroy()
	}}
	addDestroyCallback(w.destroyCallback)

	w.screen = screen
    w.frameBufferConfig = config
	w.vsync = true
	w.visible = true
	w.decorated = true
    w.resizable = true
	w.title = "Chippy Window"
	w.width = 640
	w.height = 480
    w.maxWidth = 0 // unlimited
    w.maxHeight = 0 // unlimited
    w.minWidth = 100
    w.minHeight = 100

	w.x = int16((screen.Width() / 2) - (w.width / 2))
	w.y = int16((screen.Height() / 2) - (w.height / 2))

    w.keyEventsChan               = make(chan *KeyEvent, eventBufferLarge)
    w.mouseButtonEventsChan       = make(chan *MouseButtonEvent, eventBufferLarge)
    w.cursorWithinEventsChan      = make(chan *CursorWithinEvent, eventBufferSmall)
    w.mouseMoveEventsChan         = make(chan *MouseMoveEvent, eventBufferLarge)
    w.cursorMoveEventsChan        = make(chan *CursorMoveEvent, eventBufferLarge)
    w.windowResizeEventsChan      = make(chan *WindowResizeEvent, eventBufferLarge)
    w.windowMoveEventsChan        = make(chan *WindowMoveEvent, eventBufferLarge)
    w.windowFocusEventsChan       = make(chan *WindowFocusEvent, eventBufferTiny)
    w.windowMinimizedEventsChan   = make(chan *WindowMinimizedEvent, eventBufferTiny)
    w.windowMaximizedEventsChan   = make(chan *WindowMaximizedEvent, eventBufferTiny)
    w.windowRedrawEventsChan      = make(chan *WindowRedrawEvent, eventBufferSmall)
    w.windowAlwaysOnTopEventsChan = make(chan *WindowAlwaysOnTopEvent, eventBufferTiny)
    w.windowCloseEventsChan       = make(chan *WindowCloseEvent, eventBufferTiny)


    // Set the icon
    if icon != nil {
        if icon.Bounds().Max.X != 128 || icon.Bounds().Max.Y != 128 {
            icon = resize.Resize(icon, icon.Bounds(), 128, 128)
        }

	    if icon != w.icon {
		    w.icon = icon
	    }
    }

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

// FrameBufferConfig returns the FrameBufferConfig that is in use by this Window,
// the one specified in the original call to NewWindow()
func (w *Window) FrameBufferConfig() (*FrameBufferConfig, error) {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return nil, errors.New("Window already had Destroy() called!")
	}
	return w.frameBufferConfig, nil
}

// Screen returns the Screen that this Window was created using.
//
// Note: There is no guarantee that this Window is still located on this Screen.
func (w *Window) Screen() (*Screen, error) {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return nil, errors.New("Window already had Destroy() called!")
	}
	return w.screen, nil
}

// MakeCurrent makes this Window's OpenGL context, the current OpenGL context.
//
// Do note that OpenGL is state based and calling this function makes this OpenGL context the current context within
// the current OS thread only, so use runtime.LockOSThread and runtime.UnlockOSThread appropriately as needed when
// interfacing with this.
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
// You should always call this on Windows with DoubleBuffered FrameBufferConfig's,
// specifically after you are done executing your render code for the frame.
//
// This function is no-op if this Window has no DoubleBuffered FrameBufferConfig.
func (w *Window) SwapBuffers() error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if w.frameBufferConfig.DoubleBuffered {
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
		return errors.New("Window already had Destroy() called!")
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

// Title returns the title string of this Window, this is an valid UTF-8 Go string
func (w *Window) Title() string {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.title
}

// SetWidth sets the width (in pixels) of this Window's drawable region, and also inheritely sets the width
// of this Window's region as well.
func (w *Window) SetWidth(width uint16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
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

// Width returns the width (in pixels) of this Window's drawable region
func (w *Window) Width() uint16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.width
}

// SetHeight sets the height (in pixels) of this Window's drawable region, and also inheritely sets the height
// of this Window's region as well.
func (w *Window) SetHeight(height uint16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
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

// Height returns the height (in pixels) of this Window's drawable region
func (w *Window) Height() uint16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.height
}

// SetSize sets the width and height of this Window's drawable region (in pixels)
func (w *Window) SetSize(width, height uint16) error {
	err := w.SetWidth(width)
	err = w.SetHeight(height)
	if err != nil {
		return err
	}
	return nil
}

// Size returns the width and height of this Window's drawable region (in pixels)
func (w *Window) Size() (uint16, uint16) {
	return w.Width(), w.Height()
}

// Extents returns the left, right, bottom, and top extents of this Window. That is, how many pixels
// out this Window's region extends from this Window's drawable region. (The size of this Window's
// decorations/borders)
func (w *Window) Extents() (uint8, uint8, uint8, uint8) {
	w.access.RLock()
	defer w.access.RUnlock()
    return w.extentLeft, w.extentRight, w.extentBottom, w.extentTop
}

// SetMaxWidth sets the maximum width (in pixels) that this Window's drawable region may be
// or use 0 for 'unlimited' width
func (w *Window) SetMaxWidth(maxWidth uint16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if maxWidth != w.maxWidth {
		w.maxWidth = maxWidth

		// Calling into C -- Get the lock
		chippyAccess.Lock()
		defer chippyAccess.Unlock()

		err := getInitError()
		if err != nil {
			return err
		}

		err = w.setMaxSize()
		if err != nil {
			return err
		}
	}
	return nil
}

// MaxWidth returns the maximum width (in pixels) that this Window's drawable region may be
func (w *Window) MaxWidth() uint16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.maxWidth
}

// SetMaxHeight sets the maximum height (in pixels) that this Window's drawable region may be,
// or use 0 for 'unlimited' height.
func (w *Window) SetMaxHeight(maxHeight uint16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if maxHeight != w.maxHeight {
		w.maxHeight = maxHeight

		// Calling into C -- Get the lock
		chippyAccess.Lock()
		defer chippyAccess.Unlock()

		err := getInitError()
		if err != nil {
			return err
		}

		err = w.setMaxSize()
		if err != nil {
			return err
		}
	}
	return nil
}

// MaxHeight returns the maximum height (in pixels) that this Window's drawable region may be
func (w *Window) MaxHeight() uint16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.maxHeight
}

// SetMaxSize sets the maximum width and height of this Window's drawable region (in pixels),
// or use 0, 0 for 'unlimited' width and height
func (w *Window) SetMaxSize(maxWidth, maxHeight uint16) error {
	err := w.SetMaxWidth(maxWidth)
	err = w.SetMaxHeight(maxHeight)
	if err != nil {
		return err
	}
	return nil
}

// MaxSize returns the maximum width and height of this Window's drawable region (in pixels)
func (w *Window) MaxSize() (uint16, uint16) {
	return w.MaxWidth(), w.MaxHeight()
}

// SetMinWidth sets the minimum width (in pixels) that this Window's drawable region may be
func (w *Window) SetMinWidth(minWidth uint16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if minWidth != w.minWidth {
		w.minWidth = minWidth

		// Calling into C -- Get the lock
		chippyAccess.Lock()
		defer chippyAccess.Unlock()

		err := getInitError()
		if err != nil {
			return err
		}

		err = w.setMinSize()
		if err != nil {
			return err
		}
	}
	return nil
}

// MinWidth returns the minimum width (in pixels) that this Window's drawable region may be
func (w *Window) MinWidth() uint16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.minWidth
}

// SetMinHeight sets the minimum height (in pixels) that this Window's drawable region may be
func (w *Window) SetMinHeight(minHeight uint16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if minHeight != w.minHeight {
		w.minHeight = minHeight

		// Calling into C -- Get the lock
		chippyAccess.Lock()
		defer chippyAccess.Unlock()

		err := getInitError()
		if err != nil {
			return err
		}

		err = w.setMinSize()
		if err != nil {
			return err
		}
	}
	return nil
}

// MinHeight returns the minimum height (in pixels) that this Window's drawable region may be
func (w *Window) MinHeight() uint16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.minHeight
}

// SetMinSize sets the minimum width and height of this Window's drawable region (in pixels)
func (w *Window) SetMinSize(minWidth, minHeight uint16) error {
	err := w.SetMinWidth(minWidth)
	err = w.SetMinHeight(minHeight)
	if err != nil {
		return err
	}
	return nil
}

// MinSize returns the minimum width and height of this Window's drawable region (in pixels)
func (w *Window) MinSize() (uint16, uint16) {
	return w.MinWidth(), w.MinHeight()
}

// SetX sets the x position of this Window's region (in pixels), relative to the top left corner of the Screen
func (w *Window) SetX(x int16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
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

// X returns the x position of this Window's region (in pixels), relative to the top left corner of the Screen
func (w *Window) X() int16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.x
}

// SetY sets the y position of this Window's region (in pixels), relative to the top left corner of the Screen
func (w *Window) SetY(y int16) error {
	w.access.RLock()
	defer w.access.RUnlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
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

// Y returns the y position of this Window's region (in pixels), relative to the top left corner of the screen
func (w *Window) Y() int16 {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.y
}

// SetPosition sets the x and y position of this Window's region (in pixels), relative to the top left corner of the screen
func (w *Window) SetPosition(x, y int16) error {
	err := w.SetX(x)
	err = w.SetY(y)
	if err != nil {
		return err
	}
	return nil
}

// Position returns the x and y position of this Window's drawable region (in pixels), relative to the top left corner of the screen (0, 0)
func (w *Window) Position() (int16, int16) {
	return w.X(), w.Y()
}

// SetPositionCenter sets this window to the center of Window.Screen()
func (w *Window) SetPositionCenter() error {
	w.access.RLock()
	defer w.access.RUnlock()
	//return w.SetPosition(int16((w.screen.Width()/2)-(w.Width()/2)), int16((w.screen.Height()/2)-(w.Height()/2)))

    width := w.screen.Width() / 2
    width += uint16(w.extentLeft + w.extentRight)
    halfWidth := width / 2

    height := w.screen.Height()
    height += uint16(w.extentBottom + w.extentTop)
    halfHeight := height / 2
    return w.SetPosition(int16(halfWidth), int16(halfHeight))
}

// SetVerticalSync sets weather vertical sync (vsync) will be on or off on this Window
//
// If true, this Window will try to match the vertical refresh rate of Window.Screen
func (w *Window) SetVerticalSync(vsync bool) error {
	w.access.Lock()
	defer w.access.Unlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
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

// SetVisible sets weather the window is visible (true) or hidden (false), a hidden window will appear to
// the user as simply *gone*, as if the application was closed.
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

// Visible returns weather the window is visible or hidden, see SetVisible()
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

// Minimize minimizes this window, hiding it from the user while still allow the user to retrieve it
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

// Minimized returns weather the Window is currently minimized
func (w *Window) Minimized() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.minimized
}

// SetDecoration sets weather the Window will have Window decorations
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

// SetMouseGrabbed If mouseGrabbed is true, the mouse cursor will be hidden, this Window will stop
// sending CursorMoveEvent's (but will continue sending MouseMoveEvent's), and once the mouse cursor
// is inside of this Window's drawable region, the user will be unable to move the mouse outside of
// this Window's drawable region.
//
// If mouseGrabbed is false, the mouse cursor will be shown, this Window will begin sending
// CursorMoveEvent's again, and the user will be able to move the cursor outside of this Window's
// drawable region again.
func (w *Window) SetMouseGrabbed(mouseGrabbed bool) error {
	w.access.Lock()
	defer w.access.Unlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if mouseGrabbed != w.mouseGrabbed {
		w.mouseGrabbed = mouseGrabbed

		w.access.Unlock()
        // If they grab the mouse, hide the cursor, if they release the mouse, show the cursor
		err := w.SetCursorHidden(mouseGrabbed)
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

// SetResizable specifies that the window is either resizable (true), or non-resizable (false).
func (w *Window) SetResizable(resizable bool) error {
	w.access.Lock()
	defer w.access.Unlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if resizable != w.resizable {
		w.resizable = resizable

		// Calling into C -- Get the lock
		chippyAccess.Lock()
		defer chippyAccess.Unlock()

		err := getInitError()
		if err != nil {
			return err
		}

		err = w.setResizable()
		if err != nil {
			return err
		}
	}
	return nil
}

// Resizable tells weather this Window is currently resizable
func (w *Window) Resizable() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.resizable
}

// SetCursorPosition sets the position of the cursor relative to the top left corner of this Window's
// drawable region, in pixel coordinates, for example, X=50, Y=50 places the cursor 50 pixels right of
// the top left corner of the Window's drawable region, and 50 pixels down from the top left corner of
// the Window's drawable region.
//
// It is impossible to move the cursor further than the Window's Width() or Height() using this function.
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

// CursorPosition returns the position of the cursor relative to the top left corner of this Window's
// drawable region, in pixel coordinates, for example, X=50, Y=50 means the cursor is 50 pixels right
// of the top left corner of the Window's drawable region, and 50 pixels down from the top left corner
// of the Window's drawable region.
func (w *Window) CursorPosition() (uint16, uint16) {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorX, w.cursorY
}

// CursorWithin tells weather the cursor cursor is within this Window's drawable region or outside of it
func (w *Window) CursorWithin() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorWithin
}

// SetCursorHidden sets weather the mouse cursor will be hidden visually from the user when it enters this
// Window's drawable region.
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

// CursorHidden returns weather the mouse cursor is set to currently be hidden when it enters this
// Window's drawable region.
func (w *Window) CursorHidden() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorHidden
}

// SetCursorIcon Sets the cursor icon image that will be the mouse cursor visually when the cursor is
// inside of this Window's drawable region. the pointerX and pointerY parameters specify where the
// pointer will be located within the image (in image/pixel coordinates).
//
// This should be either an 48x48, 32x32, or 16x16 icon image (with 8-bit or 1-bit transparency).
// Other sizes _may_ work on certain platforms, but are strictly non-tested and may provide buggy
// results.
//
// Passing nil into parameter cursorIcon will restore the original cursor, as it was before any custom
// image cursors where set.
//
// Windows XP and greater have support for 48x48 icons.
//
// OS-X 8.5 and greater have support for 48x48 icons.
//
// Expect similar time frame for Linux (X11) machines.
func (w *Window) SetCursorIcon(cursorIcon image.Image, pointerX, pointerY uint8) error {
	w.access.Lock()
	defer w.access.Unlock()
	if w.destroyed {
		return errors.New("Window already had Destroy() called!")
	}

	if cursorIcon != w.cursorIcon || pointerX != w.cursorIconPointerX || pointerY != w.cursorIconPointerY {
		w.cursorIcon = cursorIcon
        w.cursorIconPointerX = pointerX
        w.cursorIconPointerY = pointerY

		// Calling into C -- Get the lock
		chippyAccess.Lock()
		defer chippyAccess.Unlock()

		err := getInitError()
		if err != nil {
			return err
		}

		err = w.setCursorIcon()
		if err != nil {
			return err
		}
	}
	return nil
}

// CursorIcon returns the cursor icon image, see SetCursorIcon()
func (w *Window) CursorIcon() image.Image {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorIcon
}

// Icon returns the Image that represents this Window's icon
//
// See NewWindow() for more information about what the icon is.
func (w *Window) Icon() image.Image {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.icon
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

// Focus tells weather this window (true), or some other window (false) has focus
func (w *Window) Focus() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.focus
}

// Notify tells the user that something important happened, typically
// the window manager displays this as a flashing icon in the tray, or
// something.
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

// Raise raises this window above all other windows, showing it visually on top of all other Window's
// (as per what the window manager decides).
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

// Destroy destroys the window, closing it visually.
//
// It's only needed to call this if you want to close the window right away, or otherwise before a call to chippy.Destroy()
//
// Otherwise a call to chippy.Destroy() will call this for you on all Windows.
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

// Destroyed returns weather the window has been Destroyed, via a call to Destroy() on the Window.
//
// See Window.Destroy()
func (w *Window) Destroyed() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.destroyed
}
