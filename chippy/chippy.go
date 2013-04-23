// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package chippy implements cross platform window management, and window rendering access.
//
// Key Features
//
// 1. Cross platform
//  We do our best to provide you with an easy to use cross platform API for manipulating multiple
//  windows and screens.
//
//  Of course, you'll always have to deal with some level of platform specific behavior, and some
//  things our just plain out of our hands. Most behavior on Linux depends on what X extensions are
//  available on the server, and there are of course inconsitencies between what you're able to do
//  through specific operating system API's (Windows places limits on GammaRamp's, for instance).
//
//  Go support for mobile devices is currently lacking, and it's hard to say if or when it will be
//  available to developers, but in the future we would like to include support for Android, IOS,
//  and Windows RT (Windows 8 for ARM) mobile operating systems (This of course, also includes each
//  of their respective graphics API's, OpenGL ES, etc).
//
//  It's also possible that there could be support for some sort of remote procedure call HTML5
//  based Javascript framework that would allow for using Chippy through the web using WebGL and
//  canvas, another idea is having Chippy work inside Google Chrome using Native Client, this of
//  course depends on what Go support for Native Client (or vice-versa) is in the future. Point is
//  that we're open to idea's.
//
// 2. Multiple monitor and window support.
//  We fully support multiple screens, changing between available screen resolutions and refresh
//  rates, as well as creating multiple windows across multiple monitors.
//
//  Of course, this can only be done on computers with multiple monitors, under specific software
//  conditions (drivers, Linux X extensions, etc).
//
// 3. Support for OpenGL
//  We support creating both OpenGL new and old style contexts, and using them inside Window's, we
//  also offer pbuffer support, shared contexts, etc, basically anything you would want to do with
//  wgl, glx, or agl.
//
//  It should be noted we only provide access to the window management portions of OpenGL at this
//  point in time, so you'll need to use any OpenGL wrapper of your choice to perform any actual
//  rendering.
//
//  It should further be noted, that OpenGL is state-based, and is inheritely very thread-specific,
//  and there is no way for us to change without changing OpenGL's nature, so you'll have to be
//  at least somewhat concious about thread-safety when using OpenGL, sorry.
//
// 4. (Future) Support for Direct3D
//  In the future, we hope to fully support using Direct3D from within Chippy window's. It should
//  be fairly easy to rig this together currently as well, if you understand alot about the window
//  management specific direct3D API's.
//
// 5. (Future) Support for OpenGL ES
//  In the future, we hope to support OpenGL ES for Android, IOS, Raspberry Pi, and finally Windows
//  RT (through Google's Angle project - http://code.google.com/p/angleproject/)
//
// 6. Thread safety
//  Chippy is thread safe, and can be fully used from within multiple goroutines without any worry
//  about operating system threads, or locks, etc.
//
//  It should be explicitly noted that while Chippy and it's API's are thread safe, anything to do
//  with OpenGL or Direct3D needs to take special care with thread-safety, due to those API's state
//  based nature.
//
// Windows Implementation Notes
//
// 1. Chippy should run on any version of Windows that is officially supported by Microsoft.
//    This currently means Windows XP or higher is required.
//
//    Most of the API's we use only require Windows 2000 Professional/Server or higher, so you may
//    have luck running on those versions of Windows as well, though it should be explicitly noted
//    that we will provide no effort to maintain support for any such version of Windows whose
//    Microsoft Extended Support date has passed.
//
//    What this means is that at any point in time, there may be new bugs, compile-time errors, or
//    any other issues with running Chippy on those platforms, and it is likely nobody will be
//    willing to fix it.
//
//    However, we will accept patches to make Chippy work on such versions of Windows, in the event
//    that anyone is interested in creating such patches, and provided that the code for these very
//    outdated versions of Window never hinder future development of Chippy.
//
// Linux Implementation Notes
//
// 1. Xlib is required.
//  We use this to interface with the X server, at this time there are no plans to use xcb or any
//  other X libraries (xgb, etc), but we're open to suggestions, so if you have valid reasoning,
//  feel free to mention it.
//
// 2. (optional) Xrandr X extension is highly reccomended
//  Xrandr provides an accurate way to determine an screen's physical size (Xlib is only accurate
//  if screen resolution is at max), also, Xrandr provides an way to change the screen resolution
//  and refresh rate (making the xf86vm extension, below, useless).
//
//  Disable at build time with the build tag: 'no_xrandr'
//
// 3. (optional) Xxf86vm X extension is used as an fallback if the Xrandr extension (above) is
// unavailable.
//  xf86vm provides an way to change the screen resolution, and refresh rate (but not determine
//  the physical screen size, as in the case of Xrandr above)
//
//  Note: xf86vm appears to sometimes report incorrect display modes, and sometimes refuses to
//  switch to actually available display modes, specifically in cases where Xrandr reports correct
//  display modes, and has no problem switching to the same display mode that xf86vm ignores.
//
//  Disable at build time with the build tag: 'no_xf86vm'
//
// 4. (optional) The X extension Xinput2 is reccomended.
//  Xinput2 provides access to sub-pixel mouse movement, very important for First Person Shooter
//  games, for instance.
//
//  Disable at build time with the build tag: 'no_xinput2'
//
// Mac OS X Implementation notes
//
// 1. There is no implementation yet.
//
// Feature Support Table
//
// The following table shows which tests are working on each respective platform:
//
//  |=========================================================================================================================|
//  | Screen Related Tests                                                                            | Windows | Linux | Mac |
//  |_________________________________________________________________________________________________|_________|_______|_____|
//  | chippy_screens_query (Queries for all available screens, lists and changes screen modes)        | Yes     | Yes   | No  |
//  | chippy_screens_gamma (Changes screen gamma ramp (brightness, contrast, gamma), restores old one)| Yes(1)  | Yes   | No  |
//  |=========================================================================================================================|
//  | Window Related Tests                                                                            | Windows | Linux | Mac |
//  |_________________________________________________________________________________________________|_________|_______|_____|
//  | chippy_window_extents (Tells the extents of an window's region and client region)               | Yes     | No    | No  |
//  | chippy_window_single (Opens an single window on the specified screen)                           | Yes     | No    | No  |
//  | chippy_window_two (Opens an two windows, each on the specified screens)                         | Yes     | No    | No  |
//  | chippy_window_visible (Opens two windows, changes each of their visibile properties)            | Yes     | No    | No  |
//  | chippy_window_decorated (Opens two windows, changes each of their decorated properties)         | Yes     | No    | No  |
//  | chippy_window_title (Opens two windows, changes each of their title properties)                 | Yes     | No    | No  |
//  | chippy_window_position (Opens two windows, changes each of their position properties)           | Yes     | No    | No  |
//  | chippy_window_size (Opens two windows, changes each of their size properties)                   | Yes     | No    | No  |
//  | chippy_window_minSize (Opens two windows, changes each of their minimum size properties)        | Yes     | No    | No  |
//  | chippy_window_maxSize (Opens two windows, changes each of their maximum size properties)        | Yes     | No    | No  |
//  | chippy_window_aspectRatio (Opens two windows, changes each of their aspect ratio properties)    | Yes     | No    | No  |
//  | chippy_window_minimized (Opens two windows, changes each of their minimized properties)         | Yes     | No    | No  |
//  | chippy_window_maximized (Opens two windows, changes each of their maximized properties)         | Yes     | No    | No  |
//  | chippy_window_fullscreen (Opens an single window, changes it's fullscreen property)             | No      | No    | No  |
//  | chippy_window_alwaysOnTop (Opens two windows, changes each of their alwaysOnTop properties)     | Yes     | No    | No  |
//  | chippy_window_notify (Opens two windows, requests each one notify the user of an event)         | Yes     | No    | No  |
//  | chippy_window_icon (Opens two windows, changes each of their icon properties)                   | No      | No    | No  |
//  | chippy_window_cursor (Opens two windows, changes each of their cursor properties)               | No      | No    | No  |
//  | chippy_window_blit (Opens two windows, uses blitting to copy pixel graphics onto each of them)  | No      | No    | No  |
//  | chippy_window_splash (Demonstrates how one might create an typical splash screen)               | No      | No    | No  |
//  |=========================================================================================================================|
//  | OpenGL Related Tests                                                                            | Windows | Linux | Mac |
//  |_________________________________________________________________________________________________|_________|_______|_____|
//  | chippy_opengl_2_1 (Opens multiple windows, uses OpenGL 2.1 rendering in each of them)           | No      | No    | No  |
//  | chippy_opengl_3_0 (Opens multiple windows, uses OpenGL 3.0 rendering in each of them)           | No      | No    | No  |
//  | chippy_opengl_shared (Opens multiple windows, uses OpenGL 2.1 with an shared context)           | No      | No    | No  |
//  | chippy_opengl_pbuffer (Opens multiple windows, uses OpenGL 2.1 and an pbuffer)                  | No      | No    | No  |
//  |=========================================================================================================================|
//  | Direct3D Related Tests                                                                          | Windows | Linux | Mac |
//  |_________________________________________________________________________________________________|_________|_______|_____|
//  | chippy_direct3d_8 (Opens multiple windows, uses Direct3D 8 rendering in each of them)           | No      | N/A   | N/A |
//  | chippy_direct3d_9 (Opens multiple windows, uses Direct3D 9 rendering in each of them)           | No      | N/A   | N/A |
//  | chippy_direct3d_10 (Opens multiple windows, uses Direct3D 10 rendering in each of them)         | No      | N/A   | N/A |
//  | chippy_direct3d_11 (Opens multiple windows, uses Direct3D 11 rendering in each of them)         | No      | N/A   | N/A |
//  |=========================================================================================================================|
//  | N/A = This feature is not available on this operating system, and will likely never become available.
//  |
//  | Support marked with (#) where # is any number, means there is an defect of some sort, as described below:
//  |
//  | (1) Microsoft Windows limits how much you can change the gamma ramp, it's hard and very hacky to
//  | remove this limitation, so you will notice only subtle changes during this test.
//
// FAQ
//
// Q: "Can I use my favorite OpenGL wrapper with Chippy? Do you have any suggestions for wrappers?"
//  Yes, in fact the very idea is that Chippy is OpenGL wrapper independant. We'll handle the OS
//  specific window handling portions for you, and you do the OpenGL rendering call portions
//  yourself using any wrapper library you like.
//
// Q: "Are there event Direct3D/DirectX wrapper libraries out there for Go? If not, what's the point of have Direct3D support?"
//  AFAIK, there are no DirectX or Direct3D wrapper libraries that anyone has released yet.
//
//  I would suggest creating your own, as it appears that is what you'll need to do in order to do anything useful with Direct3D.
//
//  The point of having Direct3D support is to provide access to it, so that others can create Direct3D wrapper libraries, and
//  hopfully, as with the general idea of Chippy, there is an more unified push towards using Go to make games and other
//  interactive 2D and 3D hardware accelerated applications.
//
// Q: "In the future will Chippy provide it's own wrappers for OpenGL or Direct3D API's?"
//  Yes, in the future we intend to provide an nice, friendly way to access each specific version of OpenGL and Direct3D.
//
//  We want to make sure, though, that Chippy never forces an specific graphics API wrapper upon developers, we'll do our best
//  to ensure that it works with all wrapper libraries, if possible.
//
// Q: "Chippy doesn't seem to provide an way to do X on operating system Y, why is this?"
//  We probably never thought of it, submit an bug report and we'll look into adding it or putting
//  it on the TODO.
//
//  Anything you can do through the window manager API should be available through Chippy -- even
//  in the event it's very platform specific.
//
// Q: "What are those sub-packages, wrappers/x11, wrappers/win32, etc? Can I use those?"
//  Of course you can, but bear in mind these are just the lowest level support code that Chippy
//  needs to get itself working on those operating systems.
//
//  It's also highly likely that they are very unfriendly to work with, you'll probably need to
//  refer yourself to several online documents regarding what they do, it also should be explicitly
//  noted that these are hardly full wrappers to those respective API's, they only do and only ever
//  will do what Chippy needs them to do, which probably is exposed through Chippy's actual API.
//
//  In other words, you can use them of course, but we really reccomend you never do.
//
// Q: "On Microsoft Windows, can I add an application icon to my program?"
//  This is more of an Go question -- but since it's closely related to the work you're probably
//  using Chippy for, we thought we'd provide an answer here.
//
//  You can place .syso files with the source of your main package, and the 6l/8l linker will
//  link that file with your program.
//
//  Take an look at the "app.rc" file inside the chippy/tests/data folder, also look at the test
//  chippy_window_single in the chippy/tests directory for more information.
//
// Q: "On Microsoft Windows, can I stop the terminal from appearing when my application launches?"
//  This is more of an Go question -- but since it's closely related to the work you're probably
//  using Chippy for, we thought we'd provide an answer here.
//
//  You can stop the terminal from appearing by using the 8l/6l linker flag "-H windowsgui" on your
//  install command, like so:
//
//  go install -ldflags "-H windowsgui" path/to/pkg
//
package chippy

import (
	"io"
	"io/ioutil"
	"log"
	"sync"
)

// Destroy callbacks here, these callbacks are called when the user calls chippy.Destroy()
type callback struct {
	callback func()
}

var destroyCallbacks []*callback

func addDestroyCallback(c *callback) {
	removeDestroyCallback(c) // In case it's already in
	destroyCallbacks = append(destroyCallbacks, c)
}

func removeDestroyCallback(c *callback) {
	for i := 0; i < len(destroyCallbacks); i++ {
		if destroyCallbacks[i] == c {
			// Remove it
			destroyCallbacks = append(destroyCallbacks[:i], destroyCallbacks[i+1:]...)
			break
		}
	}
}

var (
	globalLock sync.RWMutex

	autoRestoreOriginalScreenMode, autoRestoreOriginalGammaRamp bool

	// Tells weather chippy has been previously Init()
	isInit bool

	// Tells weather a previous call to Init() failed
	initError error
)

// SetAutoRestoreOriginalScreenMode specifies weather each Screen should automatically restore it's
// original ScreenMode.
//
// Turning this off leaves ScreenMode changes on each Screen perminantly active after an call to
// chippy.Destroy.
//
// It's advised against to turn this off (on by default) due to the fact that you will most likely
// anger users if you alter their screen resolution perminantly, forcing them to change it back by
// themselves.
func SetAutoRestoreOriginalScreenMode(restore bool) {
	globalLock.Lock()
	defer globalLock.Unlock()
	autoRestoreOriginalScreenMode = restore
}

// AutoRestoreOriginalScreenMode tells weather each Screen will currently automatically restore
// it's original ScreenMode.
//
// See: SetAutoRestoreOriginalScreenMode
func AutoRestoreOriginalScreenMode() bool {
	globalLock.RLock()
	defer globalLock.RUnlock()
	return autoRestoreOriginalScreenMode
}

// SetAutoRestoreOriginalGammaRamp specifies weather each Screen should automatically restore it's
// original GammaRamp.
//
// Turning this off leaves GammaRamp changes on each Screen perminantly active after an call to
// chippy.Destroy.
//
// It's advised against to turn this off (on by default) due to the fact that you will most likely
// anger users if you alter their screen's gamma ramp, forcing them to restart their computer in
// order to restore their original one (most likely).
func SetAutoRestoreOriginalGammaRamp(restore bool) {
	globalLock.Lock()
	defer globalLock.Unlock()
	autoRestoreOriginalGammaRamp = restore
}

// AutoRestoreOriginalGammaRamp tells weather each Screen will currently automatically restore
// it's original GammaRamp.
//
// See: SetAutoRestoreOriginalGammaRamp
func AutoRestoreOriginalGammaRamp() bool {
	globalLock.RLock()
	defer globalLock.RUnlock()
	return autoRestoreOriginalGammaRamp
}

// IsInit returns weather Chippy has been initialized via a previous call to Init().
//
// IsInit() returns false if Destroy() was previously called.
func IsInit() bool {
	globalLock.RLock()
	defer globalLock.RUnlock()

	return isInit
}

// Helper to panic unless previously initialized
func panicUnlessInit() {
	globalLock.RLock()
	defer globalLock.RUnlock()

	if !IsInit() {
		panic("Chippy must be initialized before calling this; Use Init() properly!")
	}
}

var logger *log.Logger

// SetDebugOutput specifies the io.Writer that debug output will be written to (ioutil.Discard by
// default).
func SetDebugOutput(w io.Writer) {
	logger = log.New(w, "chippy: ", log.Ltime|log.Lshortfile)
}

func init() {
	SetDebugOutput(ioutil.Discard)
}

// Init initializes Chippy, returning an error if there is a problem initializing some
// lower level part of Chippy, if an error was returned, it is disallowed to call any
// other Chippy functions. (And any attempt to do so will cause the program to panic.)
func Init() error {
	globalLock.Lock()
	defer globalLock.Unlock()

	if isInit == false {
		autoRestoreOriginalScreenMode = true
		autoRestoreOriginalGammaRamp = true

		// Now we try and initialize the backend, which may fail due to user configurations
		// or something of the sort (dumb user tries to run application on Linux box without
		// any working X11 server or something silly)
		err := backend_Init()
		if err != nil {
			initError = err
			return initError
		}

		// If we made it this far, Chippy should be loaded and ready, and everything is up to
		// the backend to handle things properly now
		isInit = true
		return nil
	}
	return nil
}

// Destroy will destroy Chippy, closing all windows previously opened using NewWindow(), etc.
// Only you know when you're done using Chippy's API, so you should know the appropriate time
// to call this as well. After calling this you are no longer allowed to call any Chippy
// functions.
//
// Typical usage is something like the following:
//
//  err := chippy.Init()
//  if err != nil {
//      handleError(err)
//  }
//  defer chippy.Destroy()
//
// You may call Init() again after calling Destroy() should you want to re-gain access to the API.
func Destroy() {
	globalLock.Lock()
	defer globalLock.Unlock()

	if isInit == true {
		// Firstly, we call each destroy callback, chippyAccess is explicitly unlocked here
		globalLock.Unlock()
		for _, callback := range destroyCallbacks {
			callback.callback()
		}
		globalLock.Lock()
		backend_Destroy()
		isInit = false
		initError = nil
		destroyCallbacks = []*callback{}
	}
}
