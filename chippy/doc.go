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
// 2. Multiple monitor and window support.
//  We fully support multiple screens, changing between available screen resolutions and refresh
//  rates, as well as creating multiple windows across multiple monitors.
//
//  Of course, this can only be done on computers with multiple monitors, under specific software
//  conditions (drivers, Linux X extensions, etc).
//
// 3. Support for OpenGL
//  We support creating both OpenGL new and old style contexts (any OpenGL version), also we can
//  abstract away many platform specific OpenGL functions for you (WGL, GLX, AGL extensions), we
//  support shared OpenGL contexts, etc.
//
//  Chippy works with any OpenGL wrappers, it does not try to solve the problem of wrapping OpenGL
//  or OpenGL extensions inside Go. (Although azul3d provides an nice OpenGL wrapper in
//  azul3d/native/opengl).
//
//  Chippy fixes up all of the platform specific parts of OpenGL for you -- but it can't do any
//  magic. OpenGL is still designed to be thread-local, Chippy will never fix this issue with
//  OpenGL, and OpenGL will never fix this because of the inherit single-threaded nature of the
//  graphics pipeline.
//
//  As such these words of advice are important: You still need to use runtime.LockOSThread(), and
//  runtime.UnlockOSThread() with OpenGL specific things, and be aware of threads in your OpenGL
//  applications, sorry.
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
//  with OpenGL, Direct3D needs to take special care with thread-safety, due to those API's use of
//  thread local data.
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
// Linux Implementation Notes
//
// 1. There is no implementation yet.
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
//  | N/A = This feature is not available on this operating system, and will likely never become
//  | available.
//  |
//  | Support marked with Yes(#) where # is any number, means there is an defect of some sort, as
//  | described below:
//  |
//  | (1) Microsoft Windows limits how much you can change the gamma ramp, it's hard and very hacky
//  | to remove this limitation, so you will notice only subtle changes during this test as opposed
//  | to other operating systems. It is not really advised in modern applications to use gamma
//  | ramps, anyhow.
//
// FAQ
//
// Q: "Are there event Direct3D/DirectX wrapper libraries out there for Go? If not, what's the
// point of have Direct3D support?"
//  To my knowledge there are no available Direct3D wrappers in Go.
//
//  Chippy in the future will provide Direct3D support -- but just the window management portions.
//
//  Once support is added for Direct3D to Chippy, there will likely be an solid wrapper also added
//  to azul3d inside azul3d/native/direct3d or something of the sort.
//
// Q: "Chippy doesn't seem to provide an way to do X on operating system Y, why is this?"
//  We probably never thought of it or had an need for it, submit an bug report.
//
// Q: "What are those sub-packages, wrappers/win32, etc? Can I use those?"
//  You should not use the sub-package wrapper libraries -- they are specific support code for
//  Chippy that can change at any point in time.
//
// Q: "On Microsoft Windows, how do I add an application icon to my program?"
//  You can place .syso files with the source of your main package, and the 6l/8l linker will
//  link that file with your program.
//
//  Take an look at the "app.rc" file inside the chippy/tests/data folder for more information.
//  Also look at the single window test located in the chippy/tests/chippy_window_single directory
//  for an example of this.
//
//  This is more of an Go programming question, but since it's closely related to the work you're
//  probably using Chippy for, we thought we'd provide an answer here.
//
// Q: "On Microsoft Windows, how do I stop the command prompt from appearing when my application
//  starts?"
//  You can stop the terminal from appearing by using the 8l/6l linker flag "-H windowsgui" on your
//  go install command, like so:
//
//  go install -ldflags "-H windowsgui" path/to/pkg
//
//  This is more of an Go programming question, but since it's closely related to the work you're
//  probably using Chippy for, we thought we'd provide an answer here.
//
package chippy

