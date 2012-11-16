package glfw

/*
#cgo CFLAGS: -DGLFW_INCLUDE_GLU
#cgo LDFLAGS: -lglfw
#include <GL/glfw3.h>
#include <stdlib.h>
*/
import "C"

import "runtime"
import "sync"

var(
    eFunctionChan = make(chan string)
    eReturnChan = make(chan bool)
    eventAccess sync.RWMutex

    error_callback func(int, string)
    size_callback func(*Window, int, int)
    close_callback func(*Window)
    refresh_callback func(*Window)
    focus_callback func(*Window, bool)
    iconify_callback func(*Window, bool)
    key_callback func(*Window, int, int)
    char_callback func(*Window, int)
    mouseButton_callback func(*Window, int, int)
    cursorPos_callback func(*Window, int, int)
    cursorEnter_callback func(*Window, bool)
    scroll_callback func(*Window, float32, float32)
)

func terminateEvents() {
    eFunctionChan <- "terminateEvents"
    <- eReturnChan
}

func glfwEventDispatcher() {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    dispatching := true
	for dispatching{
		//runtime.Gosched()
		fc := <-eFunctionChan

        if fc == "glfwPollEvents" {
            C.glfwPollEvents()
            eReturnChan <- true

        } else if fc == "glfwWaitEvents" {
            C.glfwWaitEvents()
            eReturnChan <- true

        } else if fc == "terminateEvents" {
            eReturnChan <- true
            dispatching = false // Kill the dispatcher
        }
	}
}

func initEvents() {
    go glfwEventDispatcher()
}

func PollEvents() {
    eventAccess.Lock()
    defer eventAccess.Unlock()
    eFunctionChan <- "glfwPollEvents"
    <- eReturnChan
}

func WaitEvents() {
    eventAccess.Lock()
    defer eventAccess.Unlock()
    eFunctionChan <- "glfwWaitEvents"
    <- eReturnChan
}

// ------------------------------------------------------------------------- //
// Global error callback
// ------------------------------------------------------------------------- //
//export error_doCallback
func error_doCallback(error C.int, description *C.char) {
	if error_callback != nil {
        error_callback(int(error), C.GoString(description))
	}
}

func SetErrorCallback(f func(int, string)) {
    eventAccess.Lock()
    defer eventAccess.Unlock()
	error_callback = f
}

func GetErrorCallback() func(int, string) {
    eventAccess.RLock()
    defer eventAccess.RUnlock()
	return error_callback
}

// ------------------------------------------------------------------------- //
// Global size callback
// ------------------------------------------------------------------------- //
//export size_doCallback
func size_doCallback(win C.GLFWwindow, width, height C.int) {
	window := windowFromGLFW(win)
	if window.size_callback != nil {
        window.size_callback(window, int(width), int(height))
	}
}

// ------------------------------------------------------------------------- //
// Global close callback
// ------------------------------------------------------------------------- //
//export close_doCallback
func close_doCallback(win C.GLFWwindow) bool {
	window := windowFromGLFW(win)
	if window.close_callback != nil {
        window.close_callback(window)
	}
	return window.Closeable()
}

// ------------------------------------------------------------------------- //
// Global refresh callback
// ------------------------------------------------------------------------- //
//export refresh_doCallback
func refresh_doCallback(win C.GLFWwindow) {
	window := windowFromGLFW(win)
	if window.refresh_callback != nil {
        window.refresh_callback(window)
	}
}

// ------------------------------------------------------------------------- //
// Global focus callback
// ------------------------------------------------------------------------- //
//export focus_doCallback
func focus_doCallback(win C.GLFWwindow, activated bool) {
	window := windowFromGLFW(win)
	if window.focus_callback != nil {
        window.focus_callback(window, activated)
	}
}

// ------------------------------------------------------------------------- //
// Global iconify callback
// ------------------------------------------------------------------------- //
//export iconify_doCallback
func iconify_doCallback(win C.GLFWwindow, iconified bool) {
	window := windowFromGLFW(win)
	if window.iconify_callback != nil {
        window.iconify_callback(window, iconified)
	}
}

// ------------------------------------------------------------------------- //
// Global key callback
// ------------------------------------------------------------------------- //
//export key_doCallback
func key_doCallback(win C.GLFWwindow, key, action C.int) {
	window := windowFromGLFW(win)
	if window.key_callback != nil {
        window.key_callback(window, int(key), int(action))
	}
}

// ------------------------------------------------------------------------- //
// Global char callback
// ------------------------------------------------------------------------- //
//export char_doCallback
func char_doCallback(win C.GLFWwindow, character C.int) {
	window := windowFromGLFW(win)
	if window.char_callback != nil {
        window.char_callback(window, int(character))
	}
}

// ------------------------------------------------------------------------- //
// Global mouseButton callback
// ------------------------------------------------------------------------- //
//export mouseButton_doCallback
func mouseButton_doCallback(win C.GLFWwindow, button, action C.int) {
	window := windowFromGLFW(win)
	if window.mouseButton_callback != nil {
        window.mouseButton_callback(window, int(button), int(action))
	}
}

// ------------------------------------------------------------------------- //
// Global cursorPos callback
// ------------------------------------------------------------------------- //
//export cursorPos_doCallback
func cursorPos_doCallback(win C.GLFWwindow, x, y C.int) {
	window := windowFromGLFW(win)
	if window.cursorPos_callback != nil {
        window.cursorPos_callback(window, int(x), int(y))
	}
}

// ------------------------------------------------------------------------- //
// Global cursorEnter callback
// ------------------------------------------------------------------------- //
//export cursorEnter_doCallback
func cursorEnter_doCallback(win C.GLFWwindow, entered bool) {
	window := windowFromGLFW(win)
	if window.cursorEnter_callback != nil {
        window.cursorEnter_callback(window, entered)
	}
}

// ------------------------------------------------------------------------- //
// Global scroll callback
// ------------------------------------------------------------------------- //
//export scroll_doCallback
func scroll_doCallback(win C.GLFWwindow, x, y float32) {
	window := windowFromGLFW(win)
	if window.scroll_callback != nil {
        window.scroll_callback(window, x, y)
	}
}

