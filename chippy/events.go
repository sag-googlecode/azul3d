package chippy

import(
    "code.google.com/p/azul3d/chippy/keyboard"
    "code.google.com/p/azul3d/chippy/mouse"
)

// KeyEvent represents an event where the user pushed a keyboard button, or released one, while
// the Window had focus
type KeyEvent struct {
    // The Window of which this event occured
	Window  *Window

    // The actual button that the event occured because of
    Button *keyboard.Button

    // Weather this event is an repeated event, according to Window's SetKeyRepeat(), SetKeyRepeatDelay(), and
    // SetKeyRepeatInterval() functions. (FIXME: add key repeat support)
    Repeat bool
}

func addKeyEvent(w *Window, b *keyboard.Button) {
	if w.KeyEvents() && w.Focus() {
        repeated := false
        w.access.RLock()
        defer w.access.RUnlock()
		w.keyEventsChan <- &KeyEvent{w, b, repeated}
	}
}

// KeyEvent returns a channel on which all KeyEvent's, which occur on this Window, will be sent to
func (w *Window) KeyEvent() <-chan *KeyEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.keyEventsChan
}

// SetKeyEvents enables KeyEvent's on this Window, meaning all KeyEvent's will be sent to the channel returned
// by KeyEvent(), this also implies that you will read from the channel returned by KeyEvent() if you set this
// to true, and ignore reading from the channel returned by KeyEvent() the buffer may overflow and cause a deadlock,
// causing your program to panic
func (w *Window) SetKeyEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.keyEvents = allowed
}

// KeyEvents tells weather KeyEvent's are enabled and will be sent to the channel returned by KeyEvent()
func (w *Window) KeyEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.keyEvents
}

// MouseButtonEvent represents an event where the user pushed a mouse button, released a mouse button,
// or scrolled the mouse wheel in a direction, while the window had focus.
type MouseButtonEvent struct {
    // The Window of which this event occured
	Window  *Window

    // The actual button that the event occured because of
    Button *mouse.Button
}

func addMouseButtonEvent(w *Window, b *mouse.Button) {
	if w.MouseButtonEvents() && w.Focus() {
        w.access.RLock()
        defer w.access.RUnlock()
		w.mouseButtonEventsChan <- &MouseButtonEvent{w, b}
	}
}

// MouseButtonEvent returns a channel on which all MouseButtonEvent's, which occur on this Window, will be sent to
func (w *Window) MouseButtonEvent() <-chan *MouseButtonEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.mouseButtonEventsChan
}

// SetMouseButtonEvents enables MouseButtonEvent's on this Window, meaning all MouseButtonEvent's will be sent to
// the channel returned by MouseButtonEvent(), this also implies that you will read from the channel returned by
// MouseButtonEvent() if you set this to true, and ignore reading from the channel returned by MouseButtonEvent()
// the buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetMouseButtonEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.mouseButtonEvents = allowed
}

// MouseButtonEvents tells weather MouseButtonEvent's are enabled and will be sent to the channel returned by MouseButtonEvent()
func (w *Window) MouseButtonEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.mouseButtonEvents
}



// CursorWithinEvent represents an event where the user moved the mouse cursor either inside, or outside of the
// Window's drawable region.
type CursorWithinEvent struct {
    // The Window of which this event occured
	Window *Window

    // Either true (inside) or false (outside)
	Within bool
}

func addCursorWithinEvent(w *Window) {
	if w.CursorWithinEvents() {
        w.access.RLock()
        defer w.access.RUnlock()
		w.cursorWithinEventsChan <- &CursorWithinEvent{w, w.CursorWithin()}
	}
}

// CursorWithinEvent returns a channel on which all CursorWithinEvent's, which occur on this Window, will be sent to
func (w *Window) CursorWithinEvent() <-chan *CursorWithinEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.cursorWithinEventsChan
}

// SetCursorWithinEvents enables CursorWithinEvent's on this Window, meaning all CursorWithinEvent's
// will be sent to the channel returned by CursorWithinEvent(), this also implies that you will read from
// the channel returned by CursorWithinEvent() if you set this to true, and ignore reading from the channel
// returned by CursorWithinEvent() the buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetCursorWithinEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.cursorWithinEvents = allowed
}

// CursorWithinEvents tells weather this Window has CursorWithinEvents's enabled
func (w *Window) CursorWithinEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorWithinEvents
}

// MouseMoveEvent represents an event where the user moved their mouse (within the Window's drawable
// region) in a certain direction. This is a sub-pixel accurate event, using raw mouse input on
// systems where this is available, therefor, for applications that would prefer to use sub-pixel
// mouse movement, use this over something such as CursorMoveEvent.
//
// For instance in the case of a First Person Shooter's camera, whose angle is based off mouse movement,
// use this over CursorMoveEvent due to the sub-pixel accuracy you will gain.
//
// On systems where sub-pixel mouse movement is unavailable, this event will still be sent, but only
// it will have no sub-pixel accurate number, it will be based off the change in pixel-based cursor
// movement.
type MouseMoveEvent struct {
    // The Window of which this event occured
	Window *Window

    // Sub-pixel accurate floating point number representing how many pixels the mouse moved and
    // in which direction.
	X, Y float64
}

func addMouseMoveEvent(w *Window, x, y float64) {
	if w.MouseMoveEvents() && w.CursorWithin() {
        // On certain platforms we can only ensure that we get mouse events while the
        // mouse is inside the window, so we need to ensure that is the case here.
        w.access.RLock()
        defer w.access.RUnlock()
		w.mouseMoveEventsChan <- &MouseMoveEvent{w, x, y}
	}
}

// MouseMoveEvent returns a channel on which all MouseMoveEvent's, which occur on this Window, will be sent to
func (w *Window) MouseMoveEvent() <-chan *MouseMoveEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.mouseMoveEventsChan
}

// SetMouseMoveEvents enables MouseMoveEvent's on this Window, meaning all MouseMoveEvent's will be sent to the
// channel returned by MouseMoveEvent(), this also implies that you will read from the channel returned by
// MouseMoveEvent() if you set this to true, and ignore reading from the channel returned by MouseMoveEvent()
// the buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetMouseMoveEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.mouseMoveEvents = allowed
}

// MouseMoveEvents tells weather this Window has MouseMoveEvent's enabled
func (w *Window) MouseMoveEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.mouseMoveEvents
}

// CursorMoveEvent represents an event where the user moved the mouse cursor while inside of the Window's
// drawable region.
type CursorMoveEvent struct {
    // The Window of which this event occured
	Window *Window

    // The X and Y position of the mouse cursor, and in which direction it moved.
	X, Y uint16
}

func addCursorMoveEvent(w *Window) {
	if w.CursorMoveEvents() {
		x, y := w.CursorPosition()
        w.access.RLock()
        defer w.access.RUnlock()
		w.cursorMoveEventsChan <- &CursorMoveEvent{w, x, y}
	}
}

// CursorMoveEvent returns a channel on which all CursorMoveEvent's, which occur on this Window, will be sent to
func (w *Window) CursorMoveEvent() <-chan *CursorMoveEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.cursorMoveEventsChan
}

// SetCursorMoveEvents enables CursorMoveEvent's on this Window, meaning all CursorMoveEvent's will be sent to the
// channel returned by CursorMoveEvent(), this also implies that you will read from the channel returned by
// CursorMoveEvent() if you set this to true, and ignore reading from the channel returned by CursorMoveEvent()
// the buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetCursorMoveEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.cursorMoveEvents = allowed
}

// CursorMoveEvents tells weather this Window has CursorMoveEvent's enabled
func (w *Window) CursorMoveEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.cursorMoveEvents
}

// WindowResizeEvent represents an event where the Window was resized by the user, this means that both the region and
// drawable region of the Window was resized.
type WindowResizeEvent struct {
    // The Window of which this event occured
	Window        *Window

    // The width and height of the window's drawable region, in pixels
	DrawableWidth, DrawableHeight uint16

    // The width and height of the window's region, in pixels
    Width, Height uint16
}

func addWindowResizeEvent(w *Window) {
	if w.ResizeEvents() {
		width, height := w.Size()
        w.access.RLock()
        defer w.access.RUnlock()
        ev := &WindowResizeEvent{}
        ev.Window = w
        ev.DrawableWidth = width
        ev.DrawableHeight = height
        ev.Width = width + uint16(w.extentLeft + w.extentRight)
        ev.Height = height + uint16(w.extentBottom + w.extentTop)
        w.windowResizeEventsChan <- ev
	}
}

// ResizeEvent returns a channel on which all WindowResizeEvent's, which occur on this Window, will be sent to
func (w *Window) ResizeEvent() <-chan *WindowResizeEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowResizeEventsChan
}

// SetResizeEvents enables WindowResizeEvent's on this Window, meaning all WindowResizeEvent's will be sent to
// the channel returned by ResizeEvent(), this also implies that you will read from the channel returned by
// ResizeEvent() if you set this to true, and ignore reading from the channel returned by ResizeEvent() the
// buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetResizeEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.resizeEvents = allowed
}

// ResizeEvents tells weather this Window has WindowResizeEvent's enabled
func (w *Window) ResizeEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.resizeEvents
}

// WindowMoveEvent represents an event where the user moved the window to a new location on the screen
type WindowMoveEvent struct {
    // The Window of which this event occured
	Window *Window

    // How many pixels the Window's drawable region is away from the top left corner of the Screen (0, 0).
    DrawableX, DrawableY int16

    // How many pixels the Window's region is away from the top left corner of the Screen (0, 0).
	X, Y int16
}

func addWindowMoveEvent(w *Window) {
	if w.MoveEvents() {
		x, y := w.Position()
        w.access.RLock()
        defer w.access.RUnlock()
        ev := &WindowMoveEvent{}
        ev.Window = w
        ev.X = x
        ev.Y = y
        ev.DrawableX = x - int16(w.extentLeft)
        ev.DrawableY = y - int16(w.extentTop)
		w.windowMoveEventsChan <- ev
	}
}

// MoveEvent returns a channel on which all WindowMoveEvent's, which occur on this Window, will be sent to
func (w *Window) MoveEvent() <-chan *WindowMoveEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowMoveEventsChan
}

// SetMoveEvents enables WindowMoveEvent's on this Window, meaning all WindowMoveEvent's will be sent to the
// channel returned by MoveEvent(), this also implies that you will read from the channel returned by MoveEvent()
// if you set this to true, and ignore reading from the channel returned by MoveEvent() the buffer may overflow
// and cause a deadlock, causing your program to panic
func (w *Window) SetMoveEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.moveEvents = allowed
}

// MoveEvents tells weather this Window has WindowMoveEvent's enabled
func (w *Window) MoveEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.moveEvents
}

// WindowFocusEvent represents an event where the user changed the focus of the Window
// either by clicking a different Window/area on the screen (no focus), or by clicking
// on this Window, or it's decorations (focused)
type WindowFocusEvent struct {
    // The Window of which this event occured
	Window *Window

    // Either true (the Window has focus) or false (the Window has no focus, some other Window probably does)
	Focus  bool
}

func addWindowFocusEvent(w *Window) {
	if w.FocusEvents() {
		focus := w.Focus()
        w.access.RLock()
        defer w.access.RUnlock()
		w.windowFocusEventsChan <- &WindowFocusEvent{w, focus}
	}
}

// FocusEvent returns a channel on which all WindowFocusEvent's, which occur on this Window, will be sent to
func (w *Window) FocusEvent() <-chan *WindowFocusEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowFocusEventsChan
}

// SetFocusEvents enables WindowFocusEvent's on this Window, meaning all WindowFocusEvent's will be sent to the
// channel returned by FocusEvent(), this also implies that you will read from the channel returned by FocusEvent()
// if you set this to true, and ignore reading from the channel returned by FocusEvent() the buffer may overflow
// and cause a deadlock, causing your program to panic
func (w *Window) SetFocusEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.focusEvents = allowed
}

// FocusEvents tells weather this Window has WindowFocusEvent's enabled
func (w *Window) FocusEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.focusEvents
}

// WindowMinimizedEvent represents an event where the user minimized (or restored/non-minimized) the window
type WindowMinimizedEvent struct {
    // The Window of which this event occured
	Window    *Window

    // Either true (minimized) or false (restored/non-minimized)
	Minimized bool
}

func addWindowMinimizedEvent(w *Window) {
	if w.MinimizedEvents() {
		minimized := w.Minimized()
        w.access.RLock()
        defer w.access.RUnlock()
		w.windowMinimizedEventsChan <- &WindowMinimizedEvent{w, minimized}
	}
}

// MinimizedEvent returns a channel on which all WindowMinimizedEvent's, which occur on this Window, will be sent to
func (w *Window) MinimizedEvent() <-chan *WindowMinimizedEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowMinimizedEventsChan
}

// SetMinimizedEvent enables WindowMinimizedEvent's on this Window, meaning all WindowMinimizedEvent's will be sent
// to the channel returned by MinimizedEvent(), this also implies that you will read from the channel returned by
// MinimizedEvent() if you set this to true, and ignore reading from the channel returned by MinimizedEvent() the
// buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetMinimizedEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.minimizedEvents = allowed
}

// MinimizedEvents tells weather this Window has WindowMinimizedEvent's enabled
func (w *Window) MinimizedEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.minimizedEvents
}

// WindowMaximizedEvent represents an event where the user maximized (or windowed/non-maximized) the window
type WindowMaximizedEvent struct {
    // The Window of which this event occured
	Window    *Window

    // Either true (maximized) or false (windowed/non-maximized)
	Maximized bool
}

func addWindowMaximizedEvent(w *Window) {
	if w.MaximizedEvents() {
		maximized := w.Maximized()
        w.access.RLock()
        defer w.access.RUnlock()
		w.windowMaximizedEventsChan <- &WindowMaximizedEvent{w, maximized}
	}
}

// MaximizedEvent returns a channel on which all WindowMaximizedEvent's, which occur on this Window, will be sent to
func (w *Window) MaximizedEvent() <-chan *WindowMaximizedEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowMaximizedEventsChan
}

// SetMaximizedEvents enables WindowMaximizedEvent's on this Window, meaning all WindowMaximizedEvent's will be sent
// to the channel returned by MaximizedEvent(), this also implies that you will read from the channel returned by
// MaximizedEvent() if you set this to true, and ignore reading from the channel returned by MaximizedEvent() the
// buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetMaximizedEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.maximizedEvents = allowed
}

// MaximizedEvents tells weather this Window has WindowMaximizedEvent's enabled
func (w *Window) MaximizedEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.maximizedEvents
}

// WindowRedrawEvent represents an event where the window's contents need to be re-drawn according to what the
// window manager believes.
type WindowRedrawEvent struct {
    // The Window of which this event occured
	Window *Window
}

func addWindowRedrawEvent(w *Window) {
	if w.RedrawEvents() {
        w.access.RLock()
        defer w.access.RUnlock()
		w.windowRedrawEventsChan <- &WindowRedrawEvent{w}
	}
}

// RedrawEvent returns a channel on which all WindowRedrawEvent's, which occur on this Window, will be sent to
func (w *Window) RedrawEvent() <-chan *WindowRedrawEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowRedrawEventsChan
}

// SetRedrawEvents enables WindowRedrawEvent's on this Window, meaning all WindowRedrawEvent's will be sent
// to the channel returned by RedrawEvent(), this also implies that you will read from the channel returned by
// RedrawEvent() if you set this to true, and ignore reading from the channel returned by RedrawEvent() the
// buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetRedrawEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.redrawEvents = allowed
}

// RedrawEvents tells weather this Window has WindowRedrawEvent's enabled
func (w *Window) RedrawEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.redrawEvents
}

// WindowAlwaysOnTopEvent represents an event where the user changed weather the window is visually always on top
// of other windows.
type WindowAlwaysOnTopEvent struct {
    // The Window of which this event occured
	Window      *Window


    // Either true (always on top) or false (default)
	AlwaysOnTop bool
}

func addWindowAlwaysOnTopEvent(w *Window) {
	if w.AlwaysOnTopEvents() {
        w.access.RLock()
        defer w.access.RUnlock()
		w.windowAlwaysOnTopEventsChan <- &WindowAlwaysOnTopEvent{w, w.AlwaysOnTop()}
	}
}

// AlwaysOnTopEvent returns a channel on which all WindowAlwaysOnTopEvent's, which occur on this Window, will be sent to
func (w *Window) AlwaysOnTopEvent() <-chan *WindowAlwaysOnTopEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowAlwaysOnTopEventsChan
}

// SetAlwaysOnTopEvents enables WindowAlwaysOnTopEvent's on this Window, meaning all
// WindowAlwaysOnTopEvent's will be sent to the channel returned by AlwaysOnTopEvent(),
// this also implies that you will read from the channel returned by AlwaysOnTopEvent()
// if you set this to true, and ignore reading from the channel returned by AlwaysOnTopEvent()
// the buffer may overflow and cause a deadlock, causing your program to panic
func (w *Window) SetAlwaysOnTopEvents(always bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.alwaysOnTopEvents = always
}

// AlwaysOnTopEvents tells weather this Window has WindowAlwaysOnTopEvent's enabled
func (w *Window) AlwaysOnTopEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.alwaysOnTopEvents
}

// WindowCloseEvent represents an event where the user tried to close the window, either by pressing
// the close decoration on the Window (the exit button), or by using some quick key that closes the
// Window (for instance, Alt + F4 on most window managers)
type WindowCloseEvent struct {
    // The Window of which this event occured
	Window *Window
}

func addWindowCloseEvent(w *Window) {
	if w.CloseEvents() {
        w.access.RLock()
        defer w.access.RUnlock()
		w.windowCloseEventsChan <- &WindowCloseEvent{w}
	} else {
		// By default, we destroy the window unless someone else handles it, considering the user wanted this
		w.Destroy()
	}
}

// CloseEvent returns a channel on which all WindowCloseEvent's, which occur on this Window, will be sent to
func (w *Window) CloseEvent() <-chan *WindowCloseEvent {
    w.access.RLock()
    defer w.access.RUnlock()
	return w.windowCloseEventsChan
}

// SetCloseEvents enables WindowCloseEvent's on this Window, meaning all WindowCloseEvent's will
// be sent to the channel returned by WindowCloseEvent(), this also implies that you will read
// from the channel returned by WindowCloseEvent() if you set this to true, and ignore reading
// from the channel returned by WindowCloseEvent() the buffer may overflow and cause a deadlock,
// causing your program to panic
func (w *Window) SetCloseEvents(allowed bool) {
	w.access.Lock()
	defer w.access.Unlock()
	w.closeEvents = allowed
}

// CloseEvents tells weather this Window has WindowCloseEvent's enabled
func (w *Window) CloseEvents() bool {
	w.access.RLock()
	defer w.access.RUnlock()
	return w.closeEvents
}

