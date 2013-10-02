// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package chippy implements cross platform window management, and window
// rendering access.
//
// OpenGL Support
//
// We support creating both OpenGL new and old style contexts (any OpenGL
// version), also we can abstract away many platform specific OpenGL functions
// for you (WGL, GLX, AGL extensions), we support shared OpenGL contexts, etc.
//
// Chippy works with any OpenGL wrappers, it does not try to solve the problem
// of wrapping OpenGL or OpenGL extensions inside Go. (Although azul3d
// provides an nice OpenGL wrapper in azul3d/native/opengl).
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
// Thread Safety
//
// Chippy is thread safe, and can be fully used from within multiple
// goroutines without any worry about operating system threads, or locks, etc.
//
// It should be explicitly noted that while Chippy and it's API's are thread
// safe, anything to do with OpenGL needs special care regarding thread
// safety.
//
// Microsoft Windows Requirements
//
// Chippy requires Windows XP or higher.
//
// It *should* also work on Windows 2000 Professional and Server editions, but
// support for these version is not tested actively.
//
// FAQ
//
// Q: On Microsoft Windows, how do I add an application icon to my program?
//  You can place .syso files with the source of your main package, and the
//  6l/8l linker will link that file with your program.
//
//  Take an look at the "app.rc" file inside the chippy/tests/data folder for
//  more information. Also look at the single window test located in the
//  chippy/tests/chippy_window_single directory for an example of this.
//
//  This is more of an Go programming question, but since it's closely related
//  to the work you're probably using Chippy for, we thought we'd provide an
//  answer here.
//
// Q: On Microsoft Windows, how do I stop the command prompt from appearing
// when my application starts?
//  You can stop the terminal from appearing by using the 8l/6l linker flag
//  "-H windowsgui" on your 'go install' command, like so:
//
//  go install -ldflags "-H windowsgui" path/to/pkg
//
//  This is more of an Go programming question, but since it's closely related
//  to the work you're probably using Chippy for, we thought we'd provide an
//  answer here.
//
// Q: What are those sub-packages wrappers/... ? Can I use those?
//  They provide support code that Chippy uses internally.
//
//  You should not use them.
//
package chippy
