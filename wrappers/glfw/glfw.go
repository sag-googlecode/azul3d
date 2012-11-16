package glfw

/*
#cgo CFLAGS: -DGLFW_INCLUDE_GLU
#cgo LDFLAGS: -lglfw
#include <GL/glfw3.h>
#include <stdlib.h>

GLFWvidmode* vidModeAtIndex(GLFWvidmode* vm, int index) {
    return vm + index;
}

void error_initCallback(void);
void size_initCallback(GLFWwindow);
void close_initCallback(GLFWwindow);
void refresh_initCallback(GLFWwindow);
void focus_initCallback(GLFWwindow);
void iconify_initCallback(GLFWwindow);
void key_initCallback(GLFWwindow);
void char_initCallback(GLFWwindow);
void mouseButton_initCallback(GLFWwindow);
void cursorPos_initCallback(GLFWwindow);
void cursorEnter_initCallback(GLFWwindow);
void scroll_initCallback(GLFWwindow);
*/
import "C"

import "runtime"
import "unsafe"
import "errors"
import "sync"

// Any call that will put something into functionChan and expect a response
// via returnChan must use this lock, otherwise requests/responses could be
// mixed up, causing a deadlock or stall
var glfwLock sync.Mutex

// glfw calls are ran through a channel and values returned on a channel
// glfw calls may only be made on the main thread, so these channels
// handle this.
var functionChan = make(chan []interface{})
var returnChan = make(chan []interface{})

func Terminate() {
    glfwLock.Lock()
    defer glfwLock.Unlock()


    terminateEvents()
	functionChan <- []interface{}{"glfwTerminate"}
	<-returnChan
}

func glfwFunctionDispatcher() {
    runtime.LockOSThread() // Lock to thread
    defer runtime.UnlockOSThread()

    dispatching := true
    for dispatching{
	    //runtime.Gosched()
	    fc := <-functionChan
        //fmt.Println("call", fc[0], C.glfwGetTime())

        if fc[0] == "glfwInit" {
            C.error_initCallback()

            r := C.glfwInit()
            returnChan <- []interface{}{r}
            if r != 0 {
                initEvents()
            }

	    } else if fc[0] == "glfwCreateWindow" {
		    if fc[5] != nil {
			    returnChan <- []interface{}{C.glfwCreateWindow(fc[1].(C.int), fc[2].(C.int), fc[3].(C.int), fc[4].(*C.char), fc[5].(C.GLFWwindow))}
		    } else {
			    returnChan <- []interface{}{C.glfwCreateWindow(fc[1].(C.int), fc[2].(C.int), fc[3].(C.int), fc[4].(*C.char), nil)}
		    }

	    } else if fc[0] == "glfwSetWindowUserPointer" {
		    C.glfwSetWindowUserPointer(fc[1].(C.GLFWwindow), fc[2].(unsafe.Pointer))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetVersion" {
		    C.glfwGetVersion(fc[1].(*C.int), fc[2].(*C.int), fc[3].(*C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetVersionString" {
		    returnChan <- []interface{}{C.glfwGetVersionString()}

	    } else if fc[0] == "glfwGetError" {
		    returnChan <- []interface{}{C.glfwGetError()}

	    } else if fc[0] == "glfwErrorString" {
		    C.glfwErrorString(fc[1].(C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetDesktopMode" {
		    C.glfwGetDesktopMode(fc[1].(*C.GLFWvidmode))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetVideoModes" {
		    returnChan <- []interface{}{C.glfwGetVideoModes(fc[1].(*C.int))}

	    } else if fc[0] == "glfwSetGamma" {
		    C.glfwSetGamma(fc[1].(C.float))
		    returnChan <- nil

	    } else if fc[0] == "glfwWindowHint" {
		    C.glfwWindowHint(fc[1].(C.int), fc[2].(C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwDestroyWindow" {
		    C.glfwDestroyWindow(fc[1].(C.GLFWwindow))
		    returnChan <- nil

	    } else if fc[0] == "glfwSetWindowTitle" {
		    C.glfwSetWindowTitle(fc[1].(C.GLFWwindow), fc[2].(*C.char))
		    returnChan <- nil

	    } else if fc[0] == "glfwSetWindowSize" {
		    C.glfwSetWindowSize(fc[1].(C.GLFWwindow), fc[2].(C.int), fc[3].(C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetWindowSize" {
		    C.glfwGetWindowSize(fc[1].(C.GLFWwindow), fc[2].(*C.int), fc[3].(*C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwIconifyWindow" {
		    C.glfwIconifyWindow(fc[1].(C.GLFWwindow))
		    returnChan <- nil

	    } else if fc[0] == "glfwRestoreWindow" {
		    C.glfwRestoreWindow(fc[1].(C.GLFWwindow))
		    returnChan <- nil

	    } else if fc[0] == "glfwShowWindow" {
		    C.glfwShowWindow(fc[1].(C.GLFWwindow))
		    returnChan <- nil

	    } else if fc[0] == "glfwHideWindow" {
		    C.glfwHideWindow(fc[1].(C.GLFWwindow))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetWindowParam" {
		    r := C.glfwGetWindowParam(fc[1].(C.GLFWwindow), fc[2].(C.int))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwGetWindowParam" {
		    r := C.glfwGetInputMode(fc[1].(C.GLFWwindow), fc[2].(C.int))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwSetInputMode" {
		    C.glfwSetInputMode(fc[1].(C.GLFWwindow), fc[2].(C.int), fc[3].(C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetKey" {
		    r := C.glfwGetKey(fc[1].(C.GLFWwindow), fc[2].(C.int))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwGetMouseButton" {
		    r := C.glfwGetMouseButton(fc[1].(C.GLFWwindow), fc[2].(C.int))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwGetCursorPos" {
		    C.glfwGetCursorPos(fc[1].(C.GLFWwindow), fc[2].(*C.int), fc[3].(*C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwSetCursorPos" {
		    C.glfwSetCursorPos(fc[1].(C.GLFWwindow), fc[2].(C.int), fc[3].(C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetScrollOffset" {
		    C.glfwGetScrollOffset(fc[1].(C.GLFWwindow), fc[2].(*C.double), fc[3].(*C.double))
		    returnChan <- nil

	    } else if fc[0] == "glfwSetClipboardString" {
		    C.glfwSetClipboardString(fc[1].(C.GLFWwindow), fc[2].(*C.char))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetClipboardString" {
		    r := C.glfwGetClipboardString(fc[1].(C.GLFWwindow))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwMakeContextCurrent" {
		    C.glfwMakeContextCurrent(fc[1].(C.GLFWwindow))
		    returnChan <- nil

	    } else if fc[0] == "glfwCopyContext" {
		    C.glfwCopyContext(fc[1].(C.GLFWwindow), fc[2].(C.GLFWwindow), fc[3].(C.ulong))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetWindowUserPointer" {
		    r := C.glfwGetWindowUserPointer(fc[1].(C.GLFWwindow))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwGetJoystickParam" {
		    r := C.glfwGetJoystickParam(fc[1].(C.int), fc[2].(C.int))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwSetTime" {
		    C.glfwSetTime(fc[1].(C.double))
		    returnChan <- nil

	    } else if fc[0] == "glfwGetTime" {
		    r := C.glfwGetTime()
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwGetCurrentContext" {
		    r := C.glfwGetCurrentContext()
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwSwapInterval" {
		    C.glfwSwapInterval(fc[1].(C.int))
		    returnChan <- nil

	    } else if fc[0] == "glfwExtensionSupported" {
		    r := C.glfwExtensionSupported(fc[1].(*C.char))
		    returnChan <- []interface{}{r}

	    } else if fc[0] == "glfwGetProcAddress" {
		    r := C.glfwGetProcAddress(fc[1].(*C.char))
		    returnChan <- []interface{}{r}


        } else if fc[0] == "size_callback" {
            C.size_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "close_callback" {
            C.close_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "refresh_callback" {
            C.refresh_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "focus_callback" {
            C.focus_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "iconify_callback" {
            C.iconify_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "key_callback" {
            C.key_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "char_callback" {
            C.char_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "mouseButton_callback" {
            C.mouseButton_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "cursorPos_callback" {
            C.cursorPos_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "cursorEnter_callback" {
            C.cursorEnter_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil

        } else if fc[0] == "scroll_callback" {
            C.scroll_initCallback(fc[1].(C.GLFWwindow))
            returnChan <- nil


		    // JUMP1
	    } else if fc[0] == "glfwTerminate" {
		    C.glfwTerminate()
		    returnChan <- nil
            dispatching = false // Kill the dispatcher
	    }
    }
}

func Init() (error) {
    // glfw calls can only be made from the main thread, so we will
    // only execute glfw calls using the glfwFunctionDispatcher()
    go glfwFunctionDispatcher()

	functionChan <- []interface{}{"glfwInit"}
	rv := <-returnChan
    if rv[0].(C.int) == 0 {
        return errors.New("Failed to initialize GLFW!")
    }
    return nil
}

func Version() (int, int, int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	var major, minor, rev C.int
	functionChan <- []interface{}{"glfwGetVersion", &major, &minor, &rev}
	<-returnChan
	return int(major), int(minor), int(rev)
}

func VersionString() string {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetVersionString"}
	rv := <-returnChan
	return C.GoString(rv[0].(*C.char))
}

// Error handling
func GetError() int {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetError"}
	rv := <-returnChan
	return int(rv[0].(C.int))
}

func ErrorString(err int) string {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwErrorString", C.int(err)}
	rv := <-returnChan
	return C.GoString(rv[0].(*C.char))
}

// Video mode functions
func vidmodeToGLFW(v *Vidmode) *C.GLFWvidmode {
	var vm *C.GLFWvidmode
	vm = (*C.GLFWvidmode)(C.malloc(C.size_t(unsafe.Sizeof(vm))))
	vm.width = C.int(v.Width)
	vm.height = C.int(v.Height)
	vm.redBits = C.int(v.RedBits)
	vm.greenBits = C.int(v.GreenBits)
	vm.blueBits = C.int(v.BlueBits)
	return vm
}

func glfwToVidmode(v *C.GLFWvidmode) *Vidmode {
	vm := Vidmode{}
	vm.Width = int(v.width)
	vm.Height = int(v.height)
	vm.RedBits = int(v.redBits)
	vm.GreenBits = int(v.greenBits)
	vm.BlueBits = int(v.blueBits)
	return &vm
}

//GLFWvidmode* glfwGetVideoModes(int* count)
func VideoModes() []*Vidmode {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	var count C.int
	var ar []*Vidmode
	functionChan <- []interface{}{"glfwGetVideoModes", &count}
	rv := <-returnChan
	modes := rv[0].(*C.GLFWvidmode)

	for i := 0; i < int(count); i++ {
		current := C.vidModeAtIndex(modes, C.int(i))
		ar = append(ar, glfwToVidmode(current))
	}
	return ar
}

func DesktopMode() *Vidmode {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	var dm C.GLFWvidmode
	functionChan <- []interface{}{"glfwGetDesktopMode", &dm}
	<-returnChan
	return glfwToVidmode(&dm)
}

// Gamma ramp functions
func SetGamma(gamma float32) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwSetGamma", C.float(gamma)}
	<-returnChan
}

//void glfwSetGammaRamp(GLFWgammaramp* ramp)
func SetGammaRamp(ramp *Ramp) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwSetGammaRamp", gammarampToGLFW(ramp)}
	<-returnChan
}

//void glfwGetGammaRamp(GLFWgammaramp* ramp)
func GammaRamp() *Ramp {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	var ramp C.GLFWgammaramp
	functionChan <- []interface{}{"glfwGetGammaRamp", &ramp}
	<-returnChan
	return glfwToGammaramp(&ramp)
}

// Window handling
func WindowHint(target, hint int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwWindowHint", C.int(target), C.int(hint)}
	<-returnChan
}

/*
// Event handling
func PollEvents() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwPollEvents"}
	<-returnChan
}

func WaitEvents() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwWaitEvents"}
	<-returnChan
}
*/

// Input handling

// Joystick input
func JoystickParam(joy, param int) bool {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetJoystickParam", C.int(joy), C.int(param)}
	rv := <-returnChan
	return rv[0].(C.int) != 0
}

/*
int glfwGetJoystickAxes(int joy, float* axes, int numaxes)
int glfwGetJoystickButtons(int joy, unsigned char* buttons, int numbuttons)
*/

// Clipboard

// Time
func SetTime(time float32) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwSetTime", C.double(time)}
	<-returnChan
}

func Time() float32 {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetTime"}
	rv := <-returnChan
	return float32(rv[0].(C.double))
}

// OpenGL support
func CurrentContext() *Window {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetCurrentContext"}
	rv := <-returnChan
	return windowFromGLFWChan(rv[0].(C.GLFWwindow))
}

func SwapInterval(interval int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwSwapInterval", C.int(interval)}
	<-returnChan
}

func ExtensionSupported(extension string) bool {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwExtensionSupported", C.CString(extension)}
	rv := <-returnChan
	return rv[0].(C.int) != 0
}

//GLFWglproc glfwGetProcAddress(char* procname)
func ProcAddress(procname string) uintptr {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetProcAddress", C.CString(procname)}
	rv := <-returnChan
	return uintptr((unsafe.Pointer)(rv[0].(unsafe.Pointer)))
}
