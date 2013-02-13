package chippy

import(
    "image"
    "sync"
    "fmt"
)

type Window struct {
    backend_Window

    parent *Window
    isOpen, visible, decorated, minimized, maximized, fullscreen, alwaysOnTop bool
    title string
    x, y int
    width, height, minWidth, minHeight, maxWidth, maxHeight uint
    extentLeft, extentRight, extentBottom, extentTop uint
    icon, cursor image.Image

    destroyCallback *callback

    valid bool
    access sync.RWMutex
}

func (w *Window) String() string {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if w.isOpen {
        w.handleSetGetEvent(opGetAll)
    }

    return fmt.Sprintf("Window(open=%t, visible=%t, decorated=%t, minimized=%t, maximized=%t, fullscreen=%t, alwaysOnTop=%t, title=\"%s\", position=%dx%d, size=%dx%d, minSize=%dx%d, maxSize=%dx%d, extents=[%d,%d,%d,%d])", w.isOpen, w.visible, w.decorated, w.minimized, w.maximized, w.fullscreen, w.alwaysOnTop, w.title, w.x, w.y, w.width, w.height, w.minWidth, w.minHeight, w.maxWidth, w.maxHeight, w.extentLeft, w.extentRight, w.extentBottom, w.extentTop)
}

func NewWindow() *Window {
    w := &Window{}
    w.valid = true

    // Window defaults
    w.visible = true
    w.decorated = true
    w.title = "Chippy Window 世界"
    w.x = 50
    w.y = 50
    w.width = 640
    w.height = 480
    w.minWidth = 100
    w.minHeight = 100
    w.maxWidth = 0
    w.maxHeight = 0

    w.destroyCallback = &callback{func() {
        w.Close()
    }}
    addDestroyCallback(w.destroyCallback)

    return w
}

func (w *Window) panicUnlessValid() {
    if w.valid != true {
        panic("Window is invalid; must create window using chippy.NewWindow()")
    }
}

// Open opens the window, with whatever it's current properties are.
//
// An error is returned in the event that we are unable to open an window, for some reason.
func (w *Window) Open() error {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()
    if !w.isOpen {
        err := w.open()
        if err != nil {
            return err
        }
        w.isOpen = true
        return nil
    }
    return nil
}

func (w *Window) IsOpen() bool {
    w.access.RLock()
    defer w.access.RUnlock()
    w.handleSetGetEvent(opIsOpen)
    return w.isOpen
}

func (w *Window) Close() {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()
    if w.isOpen {
        w.close()
        w.isOpen = false
    }
}

