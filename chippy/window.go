// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/keyboard"
	"image"
)

type Cursor struct {
	Image image.Image
	X     uint
	Y     uint
}

// This is quite an large interface... I know.

// Window represents an single window, it will be non-visible untill the Open function is called.
type Window interface {
	keyboard.StateWatcherInterface
	BlitRenderable
	GLRenderable

	// String returns an string representation of this Window
	String() string

	// Open opens the window using the current settings, on the specified screen, or returns an
	// error in the event that we are unable to open the window for some reason (the error will be
	// descriptive).
	//
	// If WasOpened returns true, this function is no-op.
	//
	// if Destroyed returns true, this function is no-op.
	Open(screen Screen) error

	// Opened tells weather there was an previous call to the Open function or not.
	Opened() bool

	// OriginalScreen returns the screen that this window was created on, via Open()
	OriginalScreen() Screen

	// Screen returns the current screen that this window is on
	Screen() Screen

	// Destroy closes the window perminantly, after calling this function it is disallowed to call
	// any of the Set* functions, and attempting to do so will cause an panic as documented per
	// function.
	//
	// If WasOpened returns false, this function is no-op.
	//
	// If Destroyed returns true, this function is no-op.
	Destroy()

	// DestroyedEvent is an special event that is sent after Destroy() is called, despite the name
	// of this function, only one destroyed event will ever be sent over DestroyEventBuffer.Read,
	// after which the channel will be closed.
	DestroyedEvents() *DestroyedEventBuffer

	// Destroyed tells weather there was an previous call to the Destroy function.
	Destroyed() bool

	// Notify causes the window to notify the user that an event has happened with the application,
	// and they should look at the application.
	//
	// Typically this is an small flashing animation, etc.
	//
	// if Destroyed returns true, this function will panic.
	Notify()

	// Extents returns how far the window region extends outward from the client region of this
	// window, in pixels.
	Extents() (left, right, bottom, top uint)

	// Focused tells weather this window currently has focus, and is therefor the current window
	// that is being interacted with by the user.
	Focused() bool

	// FocusedEvents returns an new *FocusEventBuffer on which this Window's focus events will be
	// sent.
	FocusedEvents() *FocusedEventBuffer

	// SetTransparent specifies weather this window should be transparent, used in things like splash screens, etc.
	//
	// Default: false
	SetTransparent(transparent bool)

	// Transparent tells weather this window is transparent, via an previous call to SetTransparent()
	Transparent() bool

	// SetTitle sets the title of the window, this is shown anywhere where there needs to be an string
	// representation, typical places include the windows Title Bar decoration, and in the
	// icon tray (which displays minimized windows, etc).
	//
	// If Destroyed returns true, this function will panic.
	SetTitle(title string)

	// Title returns the title of the window, as it was set by SetTitle, or the default title:
	// "Chippy Window"
	Title() string

	// SetVisible specifies weather this window should be visibly seen by the user, if false the
	// window will appear simply gone (even though it actually exists, and you may render to it,
	// and at an later time show the window again).
	//
	// If Destroyed returns true, this function will panic.
	SetVisible(visible bool)

	// Visible tells weather this window is currently visible to the user, as previously set by the
	// SetVisible function, or the default value of true (visible)
	Visible() bool

	// SetDecorated specifies weather this window should have window decorations, this includes the
	// title bar, exit buttons, borders, system menu buttons, icons, etc.
	//
	// If Destroyed returns true, this function will panic.
	SetDecorated(decorated bool)

	// Decorations tells weather this window has window decorations on, as previously set by the
	// SetDecorations function, or the default value true (on)
	Decorated() bool

	// SetPositionCenter sets the window position such that it is perfectly in the center of the
	// specified screen.
	SetPositionCenter(screen Screen)

	// SetPosition specifies the new x and y position of this window's client region, relative to
	// the top-left corner of the screen, in pixels.
	//
	// If Destroyed returns true, this function will panic.
	SetPosition(x, y int)

	// PositionEvents returns an new *PositionEventBuffer on which this Window's position events
	// will be sent.
	PositionEvents() *PositionEventBuffer

	// Position tells what the current x and y position of this window's client region.
	Position() (x, y int)

	// SetSize specifies the new width and height of this window's client region, in pixels.
	//
	// If Destroyed returns true, this function will panic.
	SetSize(width, height uint)

	// SizeEvents returns an new *SizeEventBuffer on which this Window's size events will be sent.
	SizeEvents() *SizeEventBuffer

	// Size tells the current width and height of this window, as set previously by an call to the
	// SetSize function, or due to the user resizing the window through the window manager itself.
	Size() (width, height uint)

	// SetMinimumSize specifies the minimum width and height that this windows client region is
	// allowed to have, the user will be disallowed to resize the window any smaller than this
	// specified size.
	//
	// If either width or height are zero, then there will be no maximum size restriction placed.
	//
	// If the size passed into both SetMinimumSize and SetMaximumSize are the same, then the window
	// will be non-resizable.
	//
	// If Destroyed returns true, this function will panic.
	SetMinimumSize(width, height uint)

	// MinimumSize tells the current minimum width and height of this windows client region, as set
	// previously via the SetMinimumSize function, or the default values of width=150, height=150.
	MinimumSize() (width, height uint)

	// SetMaximumSize specifies the maximum width and height that this windows client region is
	// allowed to have, the user will be disallowed to resize the window any larger than this
	// specified size.
	//
	// If the size passed into both SetMaximumSize and SetMinimumSize are the same, then the window
	// will be non-resizable.
	//
	// If either width or height are zero, then there will be no maximum size restriction placed.
	//
	// If Destroyed returns true, this function will panic.
	SetMaximumSize(width, height uint)

	// MaximumSize tells the current maximum width and height of this windows client region, as set
	// previously via the SetMaximumSize function, or the default values of width=0, height=0
	MaximumSize() (width, height uint)

	// SetAspectRatio specifies the aspect ratio that the window should try to keep when the user
	// resizes the window.
	//
	// If the ratio is zero, then the window will be allowed to resize freely, without being
	// restricted to an aspect ratio.
	//
	// If Destroyed returns true, this function will panic.
	SetAspectRatio(ratio float32)

	// AspectRatio tells the aspect ratio that the window should try and keep when the user resizes
	// the window, as previously set via SetAspectRatio, or the default, 0.
	//
	// Note: If you want to determine the aspect ratio of the window, you should instead calculate
	// it from the Size() function, by dividing width by height.
	//
	// (Because if there was no previous call to SetAspectRatio, this function will return 0, which
	// is not the actual window aspect ratio.)
	AspectRatio() float32

	// SetMinimized specifies weather the window should currently be minimized.
	//
	// If Destroyed returns true, this function will panic.
	SetMinimized(minimized bool)

	// MinimizedEvents returns an new *MinimizedEventBuffer on which this Window's minimized events
	// will be sent.
	MinimizedEvents() *MinimizedEventBuffer

	// Minimized tells weather the window is currently minimized, as previously set via an call to
	// the SetMinimized function, or due to the user changing the minimized status of the window
	// directly through the window manager, or the default value of false.
	Minimized() bool

	// SetMaximized specifies weather the window should currently be maximized.
	//
	// If Destroyed returns true, this function will panic.
	SetMaximized(maximized bool)

	// MaximizedEvents returns an new *MaximizedEventBuffer on which this Window's maximized events
	// will be sent.
	MaximizedEvents() *MaximizedEventBuffer

	// Maximized tells weather the window is currently maximized, as previously set via an call to
	// the SetMaximized function, or due to the user changing the maximized status of the window
	// directly through the window manager, or the default value of false.
	Maximized() bool

	// SetFullscreen specifies weather the window should be full screen, consuming the entire
	// screen's size, and being the only thing displayed on the screen.
	//
	// If Destroyed returns true, this function will panic.
	SetFullscreen(fullscreen bool)

	// Fullscreen tells weather the window is currently full screen, as previously set by an call
	// to the SetFullscreen function.
	Fullscreen() bool

	// SetAlwaysOnTop specifies weather the window should be always on top of other windows.
	//
	// If Destroyed returns true, this function will panic.
	SetAlwaysOnTop(alwaysOnTop bool)

	// AlwaysOnTop tells weather the window is currently always on top of other windows, due to an
	// previous call to the SetAlwaysOnTop function, or due to the user changing the always on top
	// state directly through the window manager itself.
	AlwaysOnTop() bool

	// SetIcon specifies the window icon which should be displayed anywhere that an window icon is
	// needed, this typically includes in the title bar decoration, or in the icon tray.
	//
	// If Destroyed returns true, this function will panic.
	SetIcon(icon image.Image)

	// Icon returns the currently in use icon image, as previously set via an call to SetIcon.
	//
	// Changes made to this Image *after* an initial call to SetIcon will not be reflected by the
	// window unless you call SetIcon again.
	Icon() image.Image

	// PrepareCursor prepares the specified cursor to be displayed, but does not display it. This
	// is useful when you wish to load each frame for an cursor animation, but not cause the cursor
	// to flicker while loading them into the cache.
	PrepareCursor(cursor *Cursor)

	// FreeCursor removes the specified cursor from the internal cache. Cursors are cached because
	// it allows for SetCursor() operations to perform more quickly without performing multiple
	// copy operations underneath.
	//
	// This allows you to simply use SetCursor() as often as you wish, for instance creating cursor
	// animations and such.
	FreeCursor(cursor *Cursor)

	// SetCursor specifies the cursor to become the default cursor while the mouse is inside this
	// window's client region.
	//
	// If the cursor does not exist within the internal cache already -- it will be cached using
	// PrepareCursor() automatically. Once you are done using the cursor, you should use the
	// FreeCursor() function to remove the cursor from the internal cache (the cursor can still be
	// displayed again using FreeCursor(), it just will have to be loaded into the cache again).
	//
	// If Destroyed returns true, this function will panic.
	SetCursor(cursor *Cursor)

	// Cursor returns the currently in use cursor image, as previously set via an call to SetCursor.
	//
	// Changes made to this Image *after* an initial call to SetCursor will not be reflected by the
	// window unless you call SetCursor again.
	Cursor() *Cursor

	// SetCursorPosition sets the mouse cursor to the new position x and y, specified in pixels
	// relative to the client region of this window.
	//
	// It is possible to move the cursor outside both the client region and window region, either
	// by specifying an negative number, or an positive number larger than the window region.
	//
	// An panic will occur if this window has yet to be opened, or if it has been destroyed.
	SetCursorPosition(x, y int)

	// CursorPositionEvents returns an new *CursorPositionEventBuffer on which this Window's cursor
	// position events will be sent.
	CursorPositionEvents() *CursorPositionEventBuffer

	// CursorPosition tells the current mouse cursor position, both x and y, relative to the client
	// region of this window (specified in pixels)
	CursorPosition() (x, y int)

	// CursorWithin tells weather the mouse cursor is inside the window's client region or not.
	CursorWithin() bool

	// CursorWithinEvents returns an new *CursorWithinEventBuffer on which this Window's cursor
	// within events will be sent.
	CursorWithinEvents() *CursorWithinEventBuffer

	// SetCursorGrabbed specifies weather the mouse cursor should be grabbed, this means the cursor
	// will be invisible, and will be forced to stay within the client region of the window. This
	// behavior is the same as you would typically see in first person shooter games.
	//
	// If the cursor is being released (false), then the original cursor position will be restored
	// to where it was originally at the time of the last call to SetCursorGrabbed(true).
	//
	// If Destroyed returns true, this function will panic.
	SetCursorGrabbed(grabbed bool)

	// CursorGrabbed tells weather the mouse cursor is currently grabbed, as previously set via an
	// call to the SetCursorGrabbed function.
	CursorGrabbed() bool

	// PaintEvents returns an new *PaintEventBuffer on which this Window's paint events will be
	// sent.
	//
	// An paint event is sent when some portion of this windows client region has been damaged and
	// needs to be rendered again or else it will not be displayed to the user.
	//
	// For example, when the window is resized, or minimized and then restored, and paint event is
	// sent.
	PaintEvents() *PaintEventBuffer

	// CloseEvents returns an new *CloseEventBuffer on which this Window's close events will be
	// sent.
	CloseEvents() *CloseEventBuffer

	// MouseEvents returns an new *MouseEventBuffer on which this Window's mouse events will be
	// sent.
	MouseEvents() *MouseEventBuffer

	// KeyboardStateEvents returns an new *KeyboardStateEventBuffer on which this Window's keyboard
	// state events will be sent.
	KeyboardStateEvents() *KeyboardStateEventBuffer

	// KeyboardTypedEvents returns an new *KeyboardTypedEventBuffer on which this Window's keyboard
	// typed events will be sent.
	KeyboardTypedEvents() *KeyboardTypedEventBuffer

	// ScreenChangedEvents returns an new *ScreenChangedEventBuffer on which this Window's screen
	// change events will be sent.
	ScreenChangedEvents() *ScreenChangedEventBuffer
}

func genericSetPositionCenter(window Window, screen Screen) {
	screenWidth, screenHeight := screen.ScreenMode().Resolution()
	windowWidth, windowHeight := window.Size()
	halfScreenWidth := int(screenWidth / 2)
	halfScreenHeight := int(screenHeight / 2)
	halfWindowWidth := int(windowWidth / 2)
	halfWindowHeight := int(windowHeight / 2)
	window.SetPosition(halfScreenWidth-halfWindowWidth, halfScreenHeight-halfWindowHeight)
}

func NewWindow() Window {
	w := backend_NewWindow()
	w.SetTitle("Chippy Window")
	w.SetVisible(true)
	w.SetDecorated(true)
	w.SetPosition(100, 100)
	w.SetSize(640, 480)
	//w.SetMaximumSize(0, 0)
	w.SetMinimumSize(150, 150)
	w.SetAspectRatio(0.0)
	//w.SetMinimized(false)
	//w.SetMaximized(false)
	//w.SetFullscreen(false)
	//w.SetAlwaysOnTop(false)
	return w
}
