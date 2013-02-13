// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import(
    "image"
)

// All Set/Get methods are in here, to stop window.go from becoming too bloated

const(
    opGetAll uint = iota

    opIsOpen

    opExtents

    opParent
    opSetParent

    opVisible
    opSetVisible

    opDecorated
    opSetDecorated

    opTitle
    opSetTitle

    opPosition
    opSetPosition

    opSize
    opSetSize

    opMinimumSize
    opSetMinimumSize

    opMaximumSize
    opSetMaximumSize

    opMinimized
    opSetMinimized

    opMaximized
    opSetMaximized

    opFullscreen
    opSetFullscreen

    opAlwaysOnTop
    opSetAlwaysOnTop

    opIcon
    opSetIcon

    opCursor
    opSetCursor
)

func opString(op uint) string {
    switch(op) {
        case opGetAll: return "opGetAll"
        case opIsOpen: return "opIsOpen"
        case opExtents: return "opExtents"
        case opParent: return "opParent"
        case opSetParent: return "opSetParent"
        case opVisible: return "opVisible"
        case opSetVisible: return "opSetVisible"
        case opDecorated: return "opDecorated"
        case opSetDecorated: return "opSetDecorated"
        case opTitle: return "opTitle"
        case opSetTitle: return "opSetTitle"
        case opPosition: return "opPosition"
        case opSetPosition: return "opSetPosition"
        case opSize: return "opSize"
        case opSetSize: return "opSetSize"
        case opMinimumSize: return "opMinimumSize"
        case opSetMinimumSize: return "opSetMinimumSize"
        case opMaximumSize: return "opMaximumSize"
        case opSetMaximumSize: return "opSetMaximumSize"
        case opMinimized: return "opMinimized"
        case opSetMinimized: return "opSetMinimized"
        case opMaximized: return "opMaximized"
        case opSetMaximized: return "opSetMaximized"
        case opFullscreen: return "opFullscreen"
        case opSetFullscreen: return "opSetFullscreen"
        case opAlwaysOnTop: return "opAlwaysOnTop"
        case opSetAlwaysOnTop: return "opSetAlwaysOnTop"
        case opIcon: return "opIcon"
        case opSetIcon: return "opSetIcon"
        case opCursor: return "opCursor"
        case opSetCursor: return "opSetCursor"
    }
    return "Unknown operator identifier"
}

// Extents returns how far the window's decorations extend outwards from the client region of this
// Window (in pixels).
func (w *Window) Extents() (left, right, bottom, top uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opExtents)
    return w.extentLeft, w.extentRight, w.extentBottom, w.extentTop
}

// Parent returns the Window that is this Window's parent, or nil if this Window has no parent.
func (w *Window) Parent() *Window {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opParent)
    return w.parent
}

// SetParent sets the Window that will be this Window's parent.
func (w *Window) SetParent(parent *Window) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if parent != w.parent {
        w.parent = parent
        if w.isOpen {
            w.handleSetGetEvent(opSetParent)
        }
    }
}

// Visible tells weather this Window can be visibly seen by the user currently
func (w *Window) Visible() bool {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opVisible)
    return w.visible
}

// SetVisible specifies weather this Window can be visibily seen by the user (true) or should be
// hidden (false)
func (w *Window) SetVisible(visible bool) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if visible != w.visible {
        w.visible = visible
        if w.isOpen {
            w.handleSetGetEvent(opSetVisible)
        }
    }
}

// Decorated tells weather this Window has decorations enabled, decorations typically include the
// title bar, title bar buttons, and border/edge around the window.
func (w *Window) Decorated() bool {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opDecorated)
    return w.decorated
}

// SetDecorated specifies weather this Window will have decorations enabled, decorations typically
// include the title bar, title bar buttons, and border/edge around the window.
func (w *Window) SetDecorated(decorated bool) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if decorated != w.decorated {
        w.decorated = decorated
        if w.isOpen {
            w.handleSetGetEvent(opSetDecorated)
        }
    }
}

// Title returns the title string of this Window, as it is shown to the user
func (w *Window) Title() string {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opTitle)
    return w.title
}

// SetTitle sets the title string of this Window, this is any valid unicode string
func (w *Window) SetTitle(title string) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if title != w.title {
        w.title = title
        if w.isOpen {
            w.handleSetGetEvent(opSetTitle)
        }
    }
}

// Position returns the x, and y positions of this Window's client area, in pixel units 'away from'
// the top left corner of the screen, [0, 0]
func (w *Window) Position() (int, int) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opPosition)
    return w.x, w.y
}

// SetPosition sets the x, and y positions of this Window's client area, specified in pixel units
// 'away from' the top left corner of the screen, [0, 0]
func (w *Window) SetPosition(x, y int) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if x != w.x || y != w.y {
        w.x = x
        w.y = y
        if w.isOpen {
            w.handleSetGetEvent(opSetPosition)
        }
    }
}

// Size returns the width, and height of this Window's client area, in pixel units
func (w *Window) Size() (uint, uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opSize)
    return w.width, w.height
}

// SetSize sets the width, and height of this Window's client area, in pixel units
func (w *Window) SetSize(width, height uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if w.width != width || w.height != height {
        w.width = width
        w.height = height
        if w.isOpen {
            w.handleSetGetEvent(opSetSize)
        }
    }
}

// MinimumSize returns the minimum width, and height of this Window's client area, in pixel units.
func (w *Window) MinimumSize() (uint, uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opMinimumSize)
    return w.minWidth, w.minHeight
}

// SetMinimumSize sets the minimum width, and height of this Window's client area, in pixel units.
//
// Our best attempt will be made at stopping the user from making the Window any smaller than the
// specified size.
func (w *Window) SetMinimumSize(width, height uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if w.minWidth != width || w.minHeight != height {
        w.minWidth = width
        w.minHeight = height
        if w.isOpen {
            w.handleSetGetEvent(opSetMinimumSize)
        }
    }
}

// MaximumSize returns the maximum width, and height of this Window's client area, in pixel units.
func (w *Window) MaximumSize() (uint, uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opMaximumSize)
    return w.maxWidth, w.maxHeight
}

// SetMaximumSize sets the maximum width, and height of this Window's client area, in pixel units.
//
// Our best attempt will be made at stopping the user from making the Window any smaller than the
// specified size.
func (w *Window) SetMaximumSize(width, height uint) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if w.maxWidth != width || w.maxHeight != height {
        w.maxWidth = width
        w.maxHeight = height
        if w.isOpen {
            w.handleSetGetEvent(opSetMaximumSize)
        }
    }
}

// Minimized tells weather this Window is currently minimized or non-minimized
func (w *Window) Minimized() bool {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opMinimized)
    return w.minimized
}

// SetMinimized specifies weather this Window should be minimized or non-minimized, minimizing an
// Window implicitly disables maximized via SetMaximized(false)
func (w *Window) SetMinimized(minimized bool) {
    w.panicUnlessValid()
    panicUnlessInit()

    maximized := w.Maximized()
    if minimized && maximized {
        w.SetMaximized(false)
    }

    w.access.Lock()
    defer w.access.Unlock()

    // If they changed minimized state, or if maximized and want minimized
    if w.minimized != minimized || maximized && minimized {
        w.minimized = minimized
        if w.isOpen {
            w.handleSetGetEvent(opSetMinimized)
        }
    }
}

// Maximized tells weather this Window is currently maximized or non-maximized (windowed)
func (w *Window) Maximized() bool {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opMaximized)
    return w.maximized
}

// SetMaximized specifies weather this Window should be maximized or non-maximized, maximizing an
// Window implicitly disables minimized via SetMinimized(false)
func (w *Window) SetMaximized(maximized bool) {
    w.panicUnlessValid()
    panicUnlessInit()

    minimized := w.Minimized()
    if maximized && minimized {
        w.SetMinimized(false)
    }

    w.access.Lock()
    defer w.access.Unlock()

    // If they changed maximized state, or if minimized and want maximized
    if w.maximized != maximized || minimized && maximized {
        w.maximized = maximized
        if w.isOpen {
            w.handleSetGetEvent(opSetMaximized)
        }
    }
}

// Fullscreen tells weather this Window currently is fullscreen or windowed, consuming the entire
// space of the Screen.
func (w *Window) Fullscreen() bool {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opFullscreen)
    return w.fullscreen
}

// SetFullscreen makes this Window enter fullscreen mode, consuming the entire space of the Screen
func (w *Window) SetFullscreen(fullscreen bool) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if w.fullscreen != fullscreen {
        w.fullscreen = fullscreen
        if w.isOpen {
            w.handleSetGetEvent(opSetFullscreen)
        }
    }
}

// AlwaysOnTop tells weather this Window is set to be always on top of others, stacking-wise.
func (w *Window) AlwaysOnTop() bool {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opAlwaysOnTop)
    return w.alwaysOnTop
}

// SetAlwaysOnTop specifies weather this Window should always appear on top of other Windows.
func (w *Window) SetAlwaysOnTop(alwaysOnTop bool) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    if w.alwaysOnTop != alwaysOnTop {
        w.alwaysOnTop = alwaysOnTop
        if w.isOpen {
            w.handleSetGetEvent(opSetAlwaysOnTop)
        }
    }
}

// Icon returns the icon image that was previously set via an call to SetIcon, or nil if no such
// call has happened yet.
func (w *Window) Icon() image.Image {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opIcon)
    return w.icon
}

// SetIcon specifies the icon image that this Window will use in several places that the window
// manager likes, in most cases this includes the 'task bar', 'application switcher', as part of
// the Window's decorations, etc.
func (w *Window) SetIcon(icon image.Image) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.icon = icon
    if w.isOpen {
        w.handleSetGetEvent(opSetIcon)
    }
}

// Cursor returns the cursor image that was previously set via an call to SetCursor, or nil if no
// such call has happened yet.
func (w *Window) Cursor() image.Image {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.handleSetGetEvent(opCursor)
    return w.cursor
}

// SetCursor specifies the cursor image that will be displayed in place of the original cursor icon
// while the mouse is inside of this Window's region or client region.
func (w *Window) SetCursor(cursor image.Image) {
    w.panicUnlessValid()
    panicUnlessInit()
    w.access.Lock()
    defer w.access.Unlock()

    w.cursor = cursor
    if w.isOpen {
        w.handleSetGetEvent(opSetCursor)
    }
}


