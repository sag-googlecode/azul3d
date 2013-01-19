// Copyright 2012 Lightpoke. All rights reserved.
// Use of this source code is governed by an BSD
// license found in the License.txt file

package chippy

import(
    "code.google.com/p/azul3d/chippy/wrappers/x11"
    "errors"
)

var(
    xDisplay *x11.Display
    xDisplayName string
    xf86vm, xrandr float32
)

const(
    // We at least need 1.2 for multiple monitor support, etc, etc..
    xrandrMinimum = 1.2

    // xf86vm has almost no docs or version information that I can find...
    xf86vmMinimum = 0.0
)

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

func xGenericErrorHandler(display *x11.Display, event *x11.XErrorEvent) {
    logger.Println("ERROR", event)
}

func backend_Init() error {
    // Must be the first X11 call
    x11.XInitThreads()

    x11.XSetErrorHandler(xGenericErrorHandler)

    xDisplay = x11.XOpenDisplay(xDisplayName)
    if xDisplay == nil {
        return errors.New("Unable to connect to X11 display! Do you have an working X server installed?")
    }

    // See if we have xrandr support
    xrandr = -1
    status, _, _ := x11.XRRQueryExtension(xDisplay)
    if status != 0 {
        status, major, minor := x11.XRRQueryVersion(xDisplay)
        if status != 0 {
            xrandr = float32(major) + float32(minor) * 0.1
        }
    }

    // We get xf86vm version regardless of xrandr support
    xf86vm = -1
    status, _, _ = x11.XF86VidModeQueryExtension(xDisplay)
    if status != 0 {
        status, major, minor := x11.XF86VidModeQueryVersion(xDisplay)
        if status != 0 {
            xf86vm = float32(major) + float32(minor) * 0.1
        }
    }


    // Tell what we're going to use
    if xrandr < xrandrMinimum {
        if xrandr > 0 {
            logger.Printf("xrandr version %.1f exists, we require at least %.1f\n", xrandr, xrandrMinimum)
        } else {
            logger.Printf("xrandr extension is missing on X display.\n")
        }

        if xf86vm >= xf86vmMinimum {
            logger.Printf("Falling back to xf86vm extension...\n")
        } else {
            if xf86vm > 0 {
                logger.Printf("xf86vm version %.1f exists, we require at least %.1f\n", xf86vm, xf86vmMinimum)
            } else {
                logger.Printf("xf86vm extension is missing on X display.\n")
            }
            logger.Printf("Falling back to pure Xlib...\n")
        }
    }

    return nil
}

func backend_Destroy() {
    x11.XCloseDisplay(xDisplay)

    xDisplay = nil
    xDisplayName = ""
    xrandr = 0.0
    xf86vm = 0.0
}

