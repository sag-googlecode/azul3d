// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/x11"
	"errors"
	"fmt"
	"image"
	"sync"
	"unsafe"
)

var (
	xDisplay                 *x11.Display
	xim                      *x11.XIM
	xConnection              *x11.Connection
	xDisplayName             string
	xrandrMajor, xrandrMinor int
	xinputMajor, xinputMinor int
	xDefaultScreenNumber     int
	clearCursor              *Cursor

	ErrInvalidGLXVersion = errors.New("GLX version 1.4 is required but not available.")
)

const (
	// We at least need 1.2 for multiple monitor support, etc, etc..
	xrandrMinMajor = 1
	xrandrMinMinor = 2

	// We need Xinput2 for raw mouse input
	xinputMinMajor = 2
	xinputMinMinor = 0

	// We need GLX 1.4 for multisampling
	glxMinMajor = 1
	glxMinMinor = 4
)

var (
	// EWMH atoms
	aNetRequestFrameExtents, aNetFrameExtents, aNetWmName, aNetWmState,
	aNetWmStateFullscreen, aNetWmStateAbove, aNetWmStateMaximizedHorz,
	aNetWmStateMaximizedVert, aNetWmStateDemandsAttention, aNetWmIcon,
	aUtf8String x11.Atom

	// MOTIF atoms
	aMotifWmHints x11.Atom

	// WM atoms
	aWmProtocols, aWmDeleteWindow, aWmChangeState x11.Atom
)

func initAtoms() {
	aNetRequestFrameExtents = xConnection.InternAtom(false, "_NET_REQUEST_FRAME_EXTENTS")
	aNetFrameExtents = xConnection.InternAtom(false, "_NET_FRAME_EXTENTS")
	aNetWmName = xConnection.InternAtom(false, "_NET_WM_NAME")
	aNetWmState = xConnection.InternAtom(false, "_NET_WM_STATE")
	aNetWmStateFullscreen = xConnection.InternAtom(false, "_NET_WM_STATE_FULLSCREEN")
	aNetWmStateAbove = xConnection.InternAtom(false, "_NET_WM_STATE_ABOVE")
	aNetWmStateMaximizedHorz = xConnection.InternAtom(false, "_NET_WM_STATE_MAXIMIZED_HORZ")
	aNetWmStateMaximizedVert = xConnection.InternAtom(false, "_NET_WM_STATE_MAXIMIZED_VERT")
	aNetWmStateDemandsAttention = xConnection.InternAtom(false, "_NET_WM_STATE_DEMANDS_ATTENTION")
	aNetWmIcon = xConnection.InternAtom(false, "_NET_WM_ICON")
	aUtf8String = xConnection.InternAtom(false, "UTF8_STRING")

	aMotifWmHints = xConnection.InternAtom(false, "_MOTIF_WM_HINTS")

	aWmProtocols = xConnection.InternAtom(false, "WM_PROTOCOLS")
	aWmDeleteWindow = xConnection.InternAtom(false, "WM_DELETE_WINDOW")
	aWmChangeState = xConnection.InternAtom(false, "WM_CHANGE_STATE")
}

// SetDisplayName sets the string that will be passed into XOpenDisplay; equivalent to the DISPLAY
// environment variable on posix complaint systems.
//
// If set, this is used in place of the default DISPLAY environment variable.
//
// This function is only available on Linux.
func SetDisplayName(displayName string) {
	globalLock.Lock()
	defer globalLock.Unlock()
	xDisplayName = displayName
}

// DisplayName returns the display_name string, as it was passed into SetDisplayName.
//
// This function is only available on Linux.
func DisplayName() string {
	globalLock.RLock()
	defer globalLock.RUnlock()
	return xDisplayName
}

/*
func xGenericErrorHandler(display *x11.Display, event *x11.XErrorEvent) {
	logger().Println("Xlib Error:", x11.XGetErrorText(display, event.Code()))
}
*/

var (
	xWindowLookupAccess sync.RWMutex
	xWindowLookup       = make(map[x11.Window]*NativeWindow, 1)
)

func findWindow(w x11.Window) (*NativeWindow, bool) {
	xWindowLookupAccess.RLock()
	defer xWindowLookupAccess.RUnlock()

	nw, ok := xWindowLookup[w]
	return nw, ok
}

var shutdownEventLoop = make(chan bool, 1)
var eventLoopReady = make(chan bool, 1)

func eventLoop() {
	readySent := false
	for {
		select {
		case <-shutdownEventLoop:
			return
		default:
			break
		}

		if !readySent {
			readySent = true
			eventLoopReady <- true
		}
		e := xConnection.WaitForEvent()
		if e == nil {
			// connection is closed
			return

		} else {
			switch e.ResponseType &^ 0x80 {
			case x11.KEY_PRESS:
				ev := (*x11.KeyPressEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.KEY_RELEASE:
				ev := (*x11.KeyReleaseEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.BUTTON_PRESS:
				ev := (*x11.ButtonPressEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.BUTTON_RELEASE:
				ev := (*x11.ButtonReleaseEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.MOTION_NOTIFY:
				ev := (*x11.MotionNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.ENTER_NOTIFY:
				ev := (*x11.EnterNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.LEAVE_NOTIFY:
				ev := (*x11.LeaveNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.FOCUS_IN:
				ev := (*x11.FocusInEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.FOCUS_OUT:
				ev := (*x11.FocusOutEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Event)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.EXPOSE:
				ev := (*x11.ExposeEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.VISIBILITY_NOTIFY:
				ev := (*x11.VisibilityNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.CREATE_NOTIFY:
				ev := (*x11.CreateNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.DESTROY_NOTIFY:
				ev := (*x11.DestroyNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.MAPPING_NOTIFY:
				//ev := (*x11.MappingNotifyEvent)(unsafe.Pointer(e.EGenericEvent))

			case x11.CLIENT_MESSAGE:
				ev := (*x11.ClientMessageEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.PROPERTY_NOTIFY:
				ev := (*x11.PropertyNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.CONFIGURE_NOTIFY:
				ev := (*x11.ConfigureNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.REPARENT_NOTIFY:
				ev := (*x11.ReparentNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.MAP_NOTIFY:
				ev := (*x11.MapNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case x11.UNMAP_NOTIFY:
				ev := (*x11.UnmapNotifyEvent)(unsafe.Pointer(e.EGenericEvent))
				w, ok := findWindow(ev.Window)
				if ok {
					go w.handleEvent(e, ev)
				}

			case 0:
				ev := (*x11.RequestError)(unsafe.Pointer(e.EGenericEvent))
				str := fmt.Sprintf("ErrorCode(%d)", ev.ErrorCode)
				switch ev.ErrorCode {
				case 1:
					str = "BadRequest"
				case 2:
					str = "BadValue"
				case 3:
					str = "BadWindow"
				case 4:
					str = "BadPixmap"
				case 5:
					str = "BadAtom"
				case 6:
					str = "BadCursor"
				case 7:
					str = "BadFont"
				case 8:
					str = "BadMatch"
				case 9:
					str = "BadDrawable"
				case 10:
					str = "BadAccess"
				case 11:
					str = "BadAlloc"
				case 12:
					str = "BadColor"
				case 13:
					str = "BadGC"
				case 14:
					str = "BadIDChoice"
				case 15:
					str = "BadName"
				case 16:
					str = "BadLength"
				case 17:
					str = "BadImplementation"
				}
				logger().Println("X Request Error:", str)
				logger().Printf("%+v\n", ev)

			default:
				if !xDisplay.HandleGLXSecretEvent(e) {
					logger().Printf("Unhandled X event: %+v\n", e.EGenericEvent)
				}
			}
		}

		select {
		case <-shutdownEventLoop:
			return
		default:
			break
		}
	}
}

func atLeastVersion(realMajor, realMinor, wantedMajor, wantedMinor int) bool {
	if realMajor != wantedMajor {
		return realMajor > wantedMajor
	}
	if realMinor != wantedMinor {
		return realMinor > wantedMinor
	}
	return true
}

func init() {
	x11.XInitThreads()
}

func backend_Init() (err error) {
	// It's not really safe to clear these at backend_Destroy() time.
	xDisplayName = ""
	xrandrMajor = 0
	xrandrMinor = 0

	xDisplay = x11.XOpenDisplay(xDisplayName)
	if xDisplay == nil {
		return errors.New("Unable to open X11 display; Is the X server running?")
	}

	if clearCursor == nil {
		clearCursor = &Cursor{
			Image: image.NewRGBA(image.Rect(0, 0, 16, 16)),
			X:     0,
			Y:     0,
		}
	}

	// We want XCB to own the event queue, not Xlib which does a poor job.
	xDisplay.XSetEventQueueOwner(x11.XCBOwnsEventQueue)

	x11.XSetErrorHandler(func(err string) {
		// Erm, if something below caused an error, we might hit a deadlock
		// here because we already have the lock used by logger().. so.. this
		// is kind of a cheap and easy fix (spawning a goroutine).
		go func() {
			logger().Println("X11 Error:", err)
		}()
	})

	xConnection = x11.XGetXCBConnection(xDisplay)
	xDefaultScreenNumber = xDisplay.XDefaultScreen()

	// Initialize atoms used
	initAtoms()

	xim = xDisplay.XOpenIM(nil, nil, nil)

	go eventLoop()
	<-eventLoopReady

	// See if we have xrandr support
	xrandrMajor = -1
	xrandrMinor = -1
	if xConnection.QueryExtension("RANDR") {
		reply, err := xConnection.RandrQueryVersionReply(xConnection.RandrQueryVersion(xrandrMinMajor, xrandrMinMinor))
		if err == nil {
			xrandrMajor = int(reply.MajorVersion)
			xrandrMinor = int(reply.MinorVersion)
		} else {
			theLogger.Println(err)
		}
	}

	// Tell what we're going to use
	if !atLeastVersion(xrandrMajor, xrandrMinor, xrandrMinMajor, xrandrMinMinor) {
		if xrandrMajor > 0 || xrandrMinor > 0 {
			theLogger.Printf("xrandr version %d.%d exists, we require at least %d.%d\n", xrandrMajor, xrandrMinor, xrandrMinMajor, xrandrMinMinor)
		} else {
			theLogger.Printf("xrandr extension is missing on X display.\n")
		}
		theLogger.Println("Falling back to pure X11; screen mode switching is impossible.")
	}

	var glxMajor, glxMinor x11.Int
	if !xDisplay.GLXQueryVersion(&glxMajor, &glxMinor) || !atLeastVersion(int(glxMajor), int(glxMinor), glxMinMajor, glxMinMinor) {
		return ErrInvalidGLXVersion
	}

	return
}

func backend_Destroy() {
	shutdownEventLoop <- true
	xConnection.Disconnect()
}
