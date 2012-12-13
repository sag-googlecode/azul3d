package chippy

var eventBufferLarge = 32 // These events happen alot, higher buffer is lower latency at a small memory cost
var eventBufferSmall = 8  // These events happen rarely, so we can use smaller buffers
var eventBufferTiny  = 2  // These events are mostly toggle's, so we can use tiny buffers


// KeyEvent represents an event where the user pushed a keyboard button, or released one
type KeyEvent struct {
    Window *Window // The Window of which this event occured
    Pressed bool // Tells weather the key is pushed down (true) or released (false)
}

var keyEvents = make(chan *KeyEvent, eventBufferLarge)

func addKeyEvent(w *Window, pressed bool) {
    if w.KeyEvents() {
        keyEvents <- &KeyEvent{w, pressed}
    }
}

// KeyEvent returns a channel on which all KeyEvent's, which occur on this Window, will be sent to
func (w *Window) KeyEvent() <-chan *KeyEvent {
    return keyEvents
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



// CursorWithinEvent represents an event where the user moved the mouse cursor inside,
// or outside of the window's region on the screen.
type CursorWithinEvent struct {
    Window *Window // The Window of which this event occured
    Within bool // Either true (inside the window) or false (outside the window)
}

var cursorWithinEvents = make(chan *CursorWithinEvent, eventBufferSmall)

func addCursorWithinEvent(w *Window) {
    if w.CursorWithinEvents() {
        cursorWithinEvents <- &CursorWithinEvent{w, w.CursorWithin()}
    }
}

// CursorWithinEvent returns a channel on which all CursorWithinEvent's, which occur on this Window, will be sent to
func (w *Window) CursorWithinEvent() <-chan *CursorWithinEvent {
    return cursorWithinEvents
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



// MouseMoveEvent represents an event where the user moved there mouse in a certain direction
// this is a sub-pixel accurate event, using raw mouse input on systems where this is available
// therefor; for applications that would prefer to use sub-pixel mouse movement, use this over
// something such as CursorMoveEvent (for instance in the case of a First Person Shooter's camera
// angle based off mouse movement, use this over CursorMoveEvent due to sub-pixel accuracy)
// On systems where sub-pixel mouse movement is non-available, pixel based mouse movement will be
// used and sent as this event automatically.
type MouseMoveEvent struct {
    Window *Window // The Window of which this event occured
    X, Y float64 // Sub-pixel accurate number representing the direction in which the mouse moved
}

var mouseMoveEvents = make(chan *MouseMoveEvent, eventBufferLarge)

func addMouseMoveEvent(w *Window, x, y float64) {
    if w.MouseMoveEvents() {
        mouseMoveEvents <- &MouseMoveEvent{w, x, y}
    }
}

// MouseMoveEvent returns a channel on which all MouseMoveEvent's, which occur on this Window, will be sent to
func (w *Window) MouseMoveEvent() <-chan *MouseMoveEvent {
    return mouseMoveEvents
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



// CursorMoveEvent represents an event where the user moved the cursor while inside the window
type CursorMoveEvent struct {
    Window *Window // The Window of which this event occured
    X, Y uint16 // In pixel coordinates, a positive number specifiying how far the cursor is
    // away from the top left corner of the window
}

var cursorMoveEvents = make(chan *CursorMoveEvent, eventBufferLarge)

func addCursorMoveEvent(w *Window) {
    if w.CursorMoveEvents() {
        x, y := w.CursorPosition()
        cursorMoveEvents <- &CursorMoveEvent{w, x, y}
    }
}

// CursorMoveEvent returns a channel on which all CursorMoveEvent's, which occur on this Window, will be sent to
func (w *Window) CursorMoveEvent() <-chan *CursorMoveEvent {
    return cursorMoveEvents
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



// WindowResizeEvent represents an event where the window was resized by the user
type WindowResizeEvent struct {
    Window *Window // The Window of which this event occured
    Width, Height uint16 // The width and height of the window, in pixels
}

var windowResizeEvents = make(chan *WindowResizeEvent, eventBufferLarge)

func addWindowResizeEvent(w *Window) {
    if w.ResizeEvents() {
        width, height := w.Size()
        windowResizeEvents <- &WindowResizeEvent{w, width, height}
    }
}

// ResizeEvent returns a channel on which all WindowResizeEvent's, which occur on this Window, will be sent to
func (w *Window) ResizeEvent() <-chan *WindowResizeEvent {
    return windowResizeEvents
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



// WindowMoveEvent represents an event where the user moved the window to a new
// location on the screen
type WindowMoveEvent struct {
    Window *Window // The Window of which this event occured
    X, Y int16 // In pixel coordinates, a positive number specifiying how far the window is
    // away from the top left corner of the screen
}

var windowMoveEvents = make(chan *WindowMoveEvent, eventBufferLarge)

func addWindowMoveEvent(w *Window) {
    if w.MoveEvents() {
        x, y := w.Position()
        windowMoveEvents <- &WindowMoveEvent{w, x, y}
    }
}

// MoveEvent returns a channel on which all WindowMoveEvent's, which occur on this Window, will be sent to
func (w *Window) MoveEvent() <-chan *WindowMoveEvent {
    return windowMoveEvents
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



// WindowFocusEvent represents an event where the user changed the focus of the window
// either by clicking a different window/area on the screen (no focus), or by clicking
// this window (focused)
type WindowFocusEvent struct {
    Window *Window // The Window of which this event occured
    Focus bool // Either true (the window has focus) or false (some other window has focus)
}

var windowFocusEvents = make(chan *WindowFocusEvent, eventBufferTiny)

func addWindowFocusEvent(w *Window) {
    if w.FocusEvents() {
        focus := w.Focus()
        windowFocusEvents <- &WindowFocusEvent{w, focus}
    }
}

// FocusEvent returns a channel on which all WindowFocusEvent's, which occur on this Window, will be sent to
func (w *Window) FocusEvent() <-chan *WindowFocusEvent {
    return windowFocusEvents
}

// SetFocusEvents enables WindowFocusEvent's on this Window, meaning all WindowFocusEvent's will be sent to the
// channel returned by FocusEvent(), this also implies that you will read from the channel returned by MoveEvent()
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



// WindowMinimizedEvent represents an event where the user minimized (or restored) the window
type WindowMinimizedEvent struct {
    Window *Window // The Window of which this event occured
    Minimized bool // Either true (minimized) or false (restored/visible/non-minimized)
}

var windowMinimizedEvents = make(chan *WindowMinimizedEvent, eventBufferTiny)

func addWindowMinimizedEvent(w *Window) {
    if w.MinimizedEvents() {
        minimized := w.Minimized()
        windowMinimizedEvents <- &WindowMinimizedEvent{w, minimized}
    }
}

// MinimizedEvent returns a channel on which all WindowMinimizedEvent's, which occur on this Window, will be sent to
func (w *Window) MinimizedEvent() <-chan *WindowMinimizedEvent {
    return windowMinimizedEvents
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



// WindowMaximizedEvent represents an event where the user minimized (or restored) the window
type WindowMaximizedEvent struct {
    Window *Window
    Maximized bool // Either true (maximized) or false (windowed/non-maximized)
}

var windowMaximizedEvents = make(chan *WindowMaximizedEvent, eventBufferTiny)

func addWindowMaximizedEvent(w *Window) {
    if w.MaximizedEvents() {
        maximized := w.Maximized()
        windowMaximizedEvents <- &WindowMaximizedEvent{w, maximized}
    }
}

// MaximizedEvent returns a channel on which all WindowMaximizedEvent's, which occur on this Window, will be sent to
func (w *Window) MaximizedEvent() <-chan *WindowMaximizedEvent {
    return windowMaximizedEvents
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



// WindowRedrawEvent represents an event where the window's contents need to be re-drawn according
// to the window manager
type WindowRedrawEvent struct {
    Window *Window // The Window of which this event occured
}

var windowRedrawEvents = make(chan *WindowRedrawEvent, eventBufferSmall)

func addWindowRedrawEvent(w *Window) {
    if w.RedrawEvents() {
        windowRedrawEvents <- &WindowRedrawEvent{w}
    }
}

// RedrawEvent returns a channel on which all WindowRedrawEvent's, which occur on this Window, will be sent to
func (w *Window) RedrawEvent() <-chan *WindowRedrawEvent {
    return windowRedrawEvents
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



// WindowAlwaysOnTopEvent represents an event where the user changed weather the window is
// visually always on top of other windows
type WindowAlwaysOnTopEvent struct {
    Window *Window // The Window of which this event occured
    AlwaysOnTop bool // Tells weather the window is specified to be always on top of other windows
}

var windowAlwaysOnTopEvents = make(chan *WindowAlwaysOnTopEvent, eventBufferTiny)

func addWindowAlwaysOnTopEvent(w *Window) {
    if w.AlwaysOnTopEvents() {
        windowAlwaysOnTopEvents <- &WindowAlwaysOnTopEvent{w, w.AlwaysOnTop()}
    }
}

// AlwaysOnTopEvent returns a channel on which all WindowAlwaysOnTopEvent's, which occur on this Window, will be sent to
func (w *Window) AlwaysOnTopEvent() <-chan *WindowAlwaysOnTopEvent {
    return windowAlwaysOnTopEvents
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



// WindowCloseEvent represents an event where the user tried to close the window
type WindowCloseEvent struct {
    Window *Window // The Window of which this event occured
}

var windowCloseEvents = make(chan *WindowCloseEvent, eventBufferTiny)

func addWindowCloseEvent(w *Window) {
    if w.CloseEvents() {
        windowCloseEvents <- &WindowCloseEvent{w}
    } else {
        // By default, we destroy the window unless someone else handles it, considering the user wanted this
        w.Destroy()
    }
}

// CloseEvent returns a channel on which all WindowCloseEvent's, which occur on this Window, will be sent to
func (w *Window) CloseEvent() <-chan *WindowCloseEvent {
    return windowCloseEvents
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

