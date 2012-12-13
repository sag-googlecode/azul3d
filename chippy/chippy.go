/*
Chippy is an high performance OpenGL windowing library supporting Windows, Mac, and Linux.

Features support may be seen the in table below:
  | Feature Description          | Windows Support | Mac Support | Linux Support |
  |                              |                 |             |               |
  | Multiple Monitors            | No              | No          | Yes           |
  | Multiple Windows             | No              | No          | Yes           |
  | Raw Mouse Input              | No              | No          | No            |
  | UTF-8 Clipboard Access       | No              | No          | No            |
  | Create Window Notification   | No              | No          | Yes           |
  | Raise Window Above Others    | No              | No          | Yes           |
  | Chang Window Title (utf8)    | No              | No          | Yes           |
  | Change Window Size           | No              | No          | Yes           |
  | Change Window Position       | No              | No          | Yes           |
  | Change Window Visibility     | No              | No          | Yes           |
  | Change Window Decorations    | No              | No          | Yes           |
  | Change Window Minimized      | No              | No          | Yes           |
  | Change Window Maximized      | No              | No          | Yes           |
  | Change Window Fullscreen     | No              | No          | Yes           |
  | Change Window Always On Top  | No              | No          | Yes           |
  | Change Window Maximum Size   | No              | No          | No            |
  | CHange Window Minimum Size   | No              | No          | No            |
  | Change Window Resizable      | No              | No          | No            |
  | Change Window Vertical Sync  | No              | No          | No            |
  | Change Window Icon           | No              | No          | No            |
  | Change Window Cursor         | No              | No          | No            |
  | OpenGL 1.x - 2.x Context     | No              | No          | No            |
  | OpenGL > 3.0 Context (ARB)   | No              | No          | No            |
  | Shared OpenGL Context        | No              | No          | No            |

There are a few very important things to note about Chippy and concurrency (gorountines / threads)

Chippy functions may be called from another goroutine / thread, without any synchronization or
communication at all between the goroutines / threads! Chippy is 100% goroutine / thread safe.

While Chippy is goroutine / thread safe, OpenGL is state based, and specifically operates on data
within the thread of execution. Anything touching OpenGL, is still non goroutine/thread safe!

The only portion of Chippy that requires special attention as far as goroutines/threads is:

Calling MakeCurrent() on an Window makes that OpenGL context the current OpenGL context, in the
current OS thread of execution, so calling MakeCurrent on an Window; is required before using that
Window's OpenGL context. This means you should use runtime.LockOSThread(), then call MakeCurrent(),
execute all of your OpenGL function calls, and then finally call runtime.LockOSThread(). This is
one of the single most important things to understand, otherwise you will receive weird issues.

Chippy uses buffered channels to store Window events; this decouples the Window events from the
likely main rendering thread; making the user's input and rendering seperate. This is one of the
reasons that Chippy performs so well. Since Chippy uses a seperate goroutine to handle a Window's
events, you should in nearly all situations ensure that Go provides you with two OS threads.
You can acheive this by ensuring that GOMAXPROCS is at least two. For example use the following:
  procs := runtime.NumCPU()
  if procs < 2 {
      procs = 2
  }
  runtime.GOMAXPROCS(procs)
The above code uses the number of CPU's on the user's computer (with an minimum of two CPU's)
*/
package chippy


// X11/GLX: We ensure that all calls are thread safe by using a global
// lock, anything touching X11 or GLX requires the global lock. Since 
// X11/GLX have no requirements about running on "specifically the main
// thread", this makes the underlying C library thread-safe. Also we use
// XInitThreads to ensure no thread local data is used by X11.


import "errors"
import "sync"

type callback struct{
    callback func()
}

var chippyAccess sync.Mutex
var destroyCallbacks []*callback
var linuxDisplayName string
var isInit bool
var initError error

// Helper to add a destroy callback
func addDestroyCallback(c *callback) {
    destroyCallbacks = append(destroyCallbacks, c)
}

// Helper to ensure a destroy callback is non-registered. (removed)
func removeDestroyCallback(c *callback) {
    found := false
    var i int
    for i = 0; i < len(destroyCallbacks); i++ {
        if destroyCallbacks[i] == c {
            found = true
            break
        }
    }
    if found == false {
        return // It was never in the destroyCallbacks array, so just continue happily
    }
    destroyCallbacks = append(destroyCallbacks[:i], destroyCallbacks[i+1:]...)
}

// Helper to get intialization
func getInitError() error {
    if isInit == false {
        return errors.New("Chippy is not initialized yet!")
    }
    return initError
}

// IsInit returns weather Chippy has been initialized via a previous call to Init().
//
// IsInit() returns false if Init() and Destroy() where both called in order, as well.
func IsInit() bool {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    return isInit
}

// Init Initializes Chippy, returning an error if there is a problem initializing the
// some lower level C part of Chippy, if an error was returned, it is disallowed to call
// any other Chippy functions.
func Init() error {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    if isInit == false {
        err := initPlatform()
        if err != nil {
            initError = err
            return err
        }
        isInit = true
        return nil
    }
    return nil
}

// Destroy will destroy Chippy, closing windows previously created, call this with a defer,
// but only once you're absolutely done using Chippy, since this will de-initialize Chippy.
//
// After calling this, IsInit() will return false.
//
// You may call Init() again in such a case, should you want.
func Destroy() {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    if isInit == true {
        chippyAccess.Unlock()
        for i := 0; i < len(destroyCallbacks); i++ {
            destroyCallbacks[i].callback()
        }
        chippyAccess.Lock()
        destroyPlatform()

        linuxDisplayName = ""
        isInit = false
        initError = nil
        destroyCallbacks = []*callback{}
    }
}


// SetLinuxDisplayName sets the display_name string that will be passed into
// XOpenDisplay (See http://tronche.com/gui/x/xlib/display/opening.html) on
// Linux operating systems.
//
// The string is similar to the DISPLAY environment variable of X11.
func SetLinuxDisplayName(display_name string) error {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return err
    }

    linuxDisplayName = display_name
    return nil
}

// LinuxDisplayName returns the string previously set by SetLinuxDisplayName
func LinuxDisplayName() (string, error) {
    chippyAccess.Lock()
    defer chippyAccess.Unlock()

    err := getInitError()
    if err != nil {
        return "", err
    }

    return linuxDisplayName, nil
}

