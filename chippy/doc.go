// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package chippy implements cross platform window management, and window
// rendering access.
//
// Thread Safety
//
// Chippy is thread safe, and can be fully used from within multiple
// goroutines without any worry about operating system threads, or locks, etc.
//
// It should be explicitly noted that while Chippy and it's API's are thread
// safe, anything to do with OpenGL needs special care regarding thread
// safety.
//
// OpenGL Support
//
// We support creating both OpenGL new and old style contexts (any OpenGL
// version), also we can abstract away many platform specific OpenGL functions
// for you (WGL, GLX, AGL extensions), we support shared OpenGL contexts, etc.
//
// Chippy works with any OpenGL wrappers, it does not try to solve the problem
// of wrapping OpenGL or OpenGL extensions inside Go.
//
// Chippy fixes up all of the platform specific parts of OpenGL for you -- but
// it can't do any magic. OpenGL is still designed to be thread-local, Chippy
// will never fix this issue with OpenGL, and OpenGL will never fix this
// because of the inherit single-threaded nature of the graphics pipeline.
//
// As such you'll need to use runtime.LockOSThread(), and
// runtime.UnlockOSThread() with OpenGL specific things, and be aware of
// threads in your OpenGL related code, sorry!
//
// Microsoft Windows FAQ
//
// What versions of Windows are supported?
//  Chippy requires Windows XP or higher.
//
//  It might also work on Windows 2000 Professional and Server editions, but
//  support for these version is not tested actively.
//
// How do I add an application icon to my program?
//  You can place .syso files with the source of your main package, and the
//  6l/8l linker will link that file with your program.
//
//  Take an look at the "app.rc" file inside the chippy/tests/data folder for
//  more information. Also look at the single window test located in the
//  chippy/tests/chippy_window_single directory for an example of this.
//
// How do I stop the command prompt from appearing when my application starts?
//  You can stop the terminal from appearing by using the 8l/6l linker flag
//  "-H windowsgui" on your 'go install' command, like so:
//
//  go install -ldflags "-H windowsgui" path/to/pkg
//
// Linux-X11 FAQ
//
// What are the Linux/X11 requirements?
//  Only the Xfree86 X server is supported (I.e. the 'normal' Linux one; not
//  Apple's X server, Cygwin ones, etc).
//
//  The X-Extension 'xkb' is required and is used for various keyboard related
//  tasks.
//
//  The X-Extension 'randr' is reccomended, but not required. Without the randr
//  extension it becomes impossible to perform various screen related tasks,
//  such as switching screen modes, etc.
//
//  The X-Extension 'GLX' version 1.4 is required, it is needed for OpenGL
//  support (but the specific version 1.4 is needed for multisampling).
//
// Why is Xinput2 not supported?
//  We don't need Xinput2 for unicode keyboard input, as well we don't need it
//  for proper mouse button input.
//
//  We would like to use Xinput2's raw mouse events for while cursor grabs are
//  active, but unfortunetly due to the way the X server works raw mouse
//  movement events are only sent while no mouse buttons are pressed.
//
//  We *could* only use raw mouse movement events while no buttons are held
//  down, and then automatically switch back to pixel-based ones, but due to
//  the fact that pixel-based mouse movement is accelerated by the user's
//  configured mouse sensitivity, the movement events would be different speeds
//  depending on weather or not they held mouse buttons down or not -- which is
//  obviously very bad. For this reason we decided to not use Xinput2 at all.
//
//  Xinput2 may still be used in the future for 32-bit keycodes should there be
//  an actual use for them, multi-touch input events, or possibly multiple
//  mouse inputs (although other systems do not support this).
//
// What about Wayland support?
//  Wayland support would be interesting for the future, namely because it
//  operates with OpenGL ES which would make developing mobile games easy on a
//  desktop operating system.
//
//  But because Wayland doesn't support hardware accelerated OpenGL at this
//  time of writing, it's not currently a priority. We are however open to
//  patches for adding Wayland support.
//
// Other FAQ
//
// What are those sub-packages chippy/wrappers/... ? Can I use those?
//  You should not use them, they provide support code that Chippy uses
//  internally and may change without notice.
//
package chippy
