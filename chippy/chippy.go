/*
Chippy is an high performance OpenGL windowing library supporting Windows, Mac, and Linux.


Features
  | Feature Description          | Windows Support | Mac Support | Linux Support |
  |                              |                 |             |               |
  | Multiple Monitors            | No              | No          | Yes           |
  | Multiple Windows             | No              | No          | Yes           |
  | Raw Mouse Input              | No              | No          | Yes           |
  | UTF-8 Clipboard Strings      | No              | No          | No            |
  | Create Window Notification   | No              | No          | Yes           |
  | Raise Window Above Others    | No              | No          | Yes           |
  | Change Window Title (UTF-8)  | No              | No          | Yes           |
  | Change Window Size           | No              | No          | Yes           |
  | Change Window Position       | No              | No          | Yes           |
  | Change Window Visibility     | No              | No          | Yes           |
  | Change Window Decorations    | No              | No          | Yes           |
  | Change Window Minimized      | No              | No          | Yes           |
  | Change Window Maximized      | No              | No          | Yes           |
  | Change Window Fullscreen     | No              | No          | Yes           |
  | Change Window Always On Top  | No              | No          | Yes           |
  | Change Window Maximum Size   | No              | No          | Yes           |
  | CHange Window Minimum Size   | No              | No          | Yes           |
  | Change Window Resizable      | No              | No          | Yes           |
  | Change Window Cursor         | No              | No          | Yes           |
  | Change Cursor Visibility     | No              | No          | Yes           |
  | Change Window Vertical Sync  | No              | No          | No            |
  | Change and get Window Pixel  | No              | No          | No            |
  | Grabbing Mouse Cursor        | No              | No          | Yes           |
  | Specify Window Icon          | No              | No          | Yes           |
  | UTF-8 Keyboard Input         | No              | No          | Yes           |
  | OpenGL 1.1 - 2.1 Support     | No              | No          | No            |
  | OpenGL 3.0 - Current Support | No              | No          | No            |
  | OpenGL Render To Texture     | No              | No          | No            |
  | Direct3D Support             | Future          | N/A         | N/A           |
  |
  | Non-Unicode representable key identification (Caps Lock, Shift, etc)
  |
  | United States                | No              | No          | Yes           |

Terminology
 There are a few terms that will help you to understand how certain functions operate or what they do.


 Region
     An Window's region is the Window's entire region on the user's screen, it *includes* all
     window decorations (title bar, borders, etc).

     An Window's region begins in the top-left corner (0, 0), and extends all the way to the bottom-right
     corner (positive width, positive height).

 Drawable Region
     An Window's drawable region is the Window's inner, drawable region, the area that you are
     able to draw graphics on to, it *excludes* all window decorations (title bar, borders, etc).

     An Window's drawable region, begins in the top-left corner (0, 0) and extends all the way to
     the bottom right corner (positive width, positive height).

Concurrency
 There are a few very important things to note about Chippy and concurrency (goroutines / threads)

 Functions may be called from another goroutine/thread, without using any special synchronization or
 communication at all between the goroutines, chippy is goroutine/thread safe.

 Just because Chippy is thread safe, doesn't mean OpenGL is thread safe, whatever OpenGL wrapper you
 are using, is still non-thread safe, unless it mentions otherwise. (Unlikely!)

 There is one specific portion of Chippy that requires special attention regarding seperate goroutines:

 Calling MakeCurrent on an Window makes that Window the current OpenGL context, in *that* OS thread. It
 is required because of such, that you use runtime.LockOSThread() and runtime.UnlockOSThread() before
 you call MakeCurrent() and only call runtime.UnlockOSThread() once you're done calling OpenGL api's
 that you wish to occur inside of that Window.

 Buffered channels are used to store Window events; this decouples Window events such as mouse input,
 keyboard input, and other Window events such as dragging the Window etc, from the rendering thread.
 It's important to take advantage of this feature, because it will greatly improve the responsiveness
 of your real time application. Because of this feature, there is one small suggestion:

 Always have GOMAXPROCS at at least two, even if the computer doesn't have two CPUs. The reason for
 this is because Chippy in some cases will block one OS thread while waiting for events, the timeout
 is small (a fraction of a second), but can cause other portions of your real time application to behave
 badly or sluggish.

 You can achieve setting GOMAXPROCS two at least two at all times using the following snippet:

    procs := runtime.NumCPU()
    if procs < 2 {
        procs = 2
    }
    runtime.GOMAXPROCS(procs)

 The above code uses the number of CPU's on the user's computer (with an minimum of two CPUs at all times)

Notes
 1. Nvidia drivers have an option that allows users with multiple monitors on Linux/X11 to report multiple monitors
 to X11 as an single monitor device. In this case Chippy will only see an single Screen, in addition, per screen
 operations, such as the gamma functions, will operate across all monitors.

 2. "Why does Chippy require that an Window's icon be passed in at Window creation time? Why is there no SetIcon() function
 or other way of changing the Window icon after the Window has been created?", Put simply, many Linux/X11 window managers
 either have no support for changing an Window's icon after the Window has been opened, or it's a very buggy implimentation.

 3. "I'm running Linux/X11, and calling SetMinimized(false), never restores the Window to the non-minimized state?!" Several
 X11 window managers have full rights according to the X documentation to ignore all requests to restore an previously minimized
 Window.
    Happens with Unity window manager on Ubuntu.
    Works fine on Gnome 3 window manager (and probably Gnome) on Ubuntu
    

// Note: Ubuntu (at least on 12.10, probably on all version prior as well) Unity window manager will
// ignore all requests to restore (SetMinimized(false)) an Window. This is a bug (or 'feature') of
// Unity.


*/
package chippy

// X11/GLX: We ensure that all calls are thread safe by using a global
// lock, anything touching X11 or GLX requires the global lock. Since
// X11/GLX have no requirements about running on "specifically the main
// thread", this makes the underlying C library thread-safe. Also we use
// XInitThreads to ensure no thread local data is used by X11.

import(
    "errors"
    "sync"
    "log"
    "os"
)

// Destroy callbacks here, these callbacks are called when the user calls chippy.Destroy()
type callback struct {
	callback func()
}

var destroyCallbacks []*callback

func addDestroyCallback(c *callback) {
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
	// Any function calls that have to go back further into C, need to use this global lock
	// basically our reasoning for this is that, majority of the underlying C api's are
	// specifically non-thread safe. So apply this global lock to *most* of the C api we use
	chippyAccess sync.Mutex

	// Tells weather chippy has been previously Init()
	isInit bool

	// Tells weather a previous call to Init() failed
	initError error
)

// Helper to get intialization errors
// You should use chippyAccess with this!
func getInitError() error {
	if isInit == false {
		return errors.New("Chippy is not initialized yet!")
	}
	return initError
}

// IsInit returns weather Chippy has been initialized via a previous call to Init().
//
// IsInit() returns false if Destroy() was previously called.
func IsInit() bool {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	return isInit
}

var logger *log.Logger
func init() {
    logger = log.New(os.Stderr, "chippy: ", log.Ltime | log.Lshortfile)
}

// Init initializes Chippy, returning an error if there is a problem initializing some
// lower level part of Chippy, if an error was returned, it is disallowed to call any
// other Chippy functions. (And any attempt to do so will return the error that this
// returned.)
func Init() error {
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	if isInit == false {
		// Intialize the config system before the backend
		initConfig()

		// Now we try and initialize the backend, which may fail due to user configurations
		// or something of the sort (dumb user tries to run application on Linux box without
		// any working X11 server or something silly)
		err := backend_init()
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
	chippyAccess.Lock()
	defer chippyAccess.Unlock()

	if isInit == true {
		// Firstly, we call each destroy callback, chippyAccess is explicitly unlocked here
        chippyAccess.Unlock()
		for i := 0; i < len(destroyCallbacks); i++ {
			destroyCallbacks[i].callback()
		}
        chippyAccess.Lock()
		backend_destroy()
		isInit = false
		initError = nil
		destroyCallbacks = []*callback{}
	}
}
