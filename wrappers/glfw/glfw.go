package glfw

/*
#cgo CFLAGS: -DGLFW_INCLUDE_GLU
#cgo LDFLAGS: -lglfw
#include <GL/glfw3.h>
#include <stdlib.h>

GLFWvidmode* vidModeAtIndex(GLFWvidmode* vm, int index);
void error_initCallback(void);
void size_initCallback(void);
void close_initCallback(void);
void refresh_initCallback(void);
void focus_initCallback(void);
void iconify_initCallback(void);
void key_initCallback(void);
void char_initCallback(void);
void mouseButton_initCallback(void);
void cursorPos_initCallback(void);
void cursorEnter_initCallback(void);
void scroll_initCallback(void);
*/
import "C"

import "runtime"
import "unsafe"
import "errors"
import "sync"

//import "code.google.com/p/azul3d"
import "code.google.com/p/azul3d/logger"

// Typedefs

// OpenGL function pointer type

// The video mode structure used by glfwGetVideoModes
type Vidmode struct {
	Width, Height, RedBits, BlueBits, GreenBits int
}

func (this *Vidmode) Equals(v *Vidmode) bool {
	return this.Width == v.Width && this.Height == v.Height && this.RedBits == v.RedBits && this.BlueBits == v.BlueBits && this.GreenBits == v.GreenBits
}

// Gamma ramp
type Ramp struct {
	Red   [256]uint16
	Green [256]uint16
	Blue  [256]uint16
}

func gammarampToGLFW(r *Ramp) *C.GLFWgammaramp {
	var gr *C.GLFWgammaramp
	gr = (*C.GLFWgammaramp)(C.malloc(C.size_t(unsafe.Sizeof(gr))))
	// Copy over red
	for i := 0; i < len(r.Red); i++ {
		gr.red[i] = C.ushort(r.Red[i])
	}
	// Copy over green
	for i := 0; i < len(r.Green); i++ {
		gr.green[i] = C.ushort(r.Green[i])
	}
	// Copy over blue
	for i := 0; i < len(r.Blue); i++ {
		gr.blue[i] = C.ushort(r.Blue[i])
	}
	return gr
}

func glfwToGammaramp(gr *C.GLFWgammaramp) *Ramp {
	r := Ramp{}
	// Copy over red
	for i := 0; i < len(r.Red); i++ {
		r.Red[i] = uint16(gr.red[i])
	}
	// Copy over green
	for i := 0; i < len(r.Green); i++ {
		r.Green[i] = uint16(gr.green[i])
	}
	// Copy over blue
	for i := 0; i < len(r.Blue); i++ {
		r.Blue[i] = uint16(gr.blue[i])
	}
	return &r
}

//    unsigned short red[256] #GLFW_GAMMA_RAMP_SIZE
//    unsigned short green[256] #GLFW_GAMMA_RAMP_SIZE
//    unsigned short blue[256] #GLFW_GAMMA_RAMP_SIZE

var log = logger.New("glfw")

// Any call that will put something into functionChan and expect a response
// via returnChan must use this lock, otherwise requests/responses could be
// mixed up, causing a deadlock or stall
var glfwLock sync.Mutex

// glfw calls are ran through a channel and values returned on a channel
// glfw calls may only be made on the main thread, so this channel
// handles this.
var functionChan = make(chan []interface{})
var returnChan = make(chan []interface{})

func Terminate() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwTerminate"}
	<-returnChan
}

func Init() (error) {
	//azul.RegisterDestroy(destroy) // Register destroy callback

	// We need to make a glfw call, and those can only be made from
	// the main thread, so we grab hold of the main thread and then
	// release it later.
	runtime.LockOSThread() // Grab main thread

	if C.glfwInit() == 0 {
        return errors.New("Failed to initialize GLFW!")
	}

	C.size_initCallback()
	C.close_initCallback()
	C.refresh_initCallback()
	C.focus_initCallback()
	C.iconify_initCallback()
	C.key_initCallback()
	C.char_initCallback()
	C.mouseButton_initCallback()
	C.cursorPos_initCallback()
	C.cursorEnter_initCallback()
	C.scroll_initCallback()

    runtime.UnlockOSThread() // release main thread

	go func() {
		for {
			runtime.Gosched()
			fc := <-functionChan
			//log.Log("Call", fc)
			runtime.LockOSThread() // Grab main thread

			if fc[0] == "glfwCreateWindow" {
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

			} else if fc[0] == "glfwSetWindowPos" {
				C.glfwSetWindowPos(fc[1].(C.GLFWwindow), fc[2].(C.int), fc[3].(C.int))
				returnChan <- nil

			} else if fc[0] == "glfwGetWindowPos" {
				C.glfwGetWindowPos(fc[1].(C.GLFWwindow), fc[2].(*C.int), fc[3].(*C.int))
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

			} else if fc[0] == "glfwSwapBuffers" {
				C.glfwSwapBuffers(fc[1].(C.GLFWwindow))
				returnChan <- nil

			} else if fc[0] == "glfwCopyContext" {
				C.glfwCopyContext(fc[1].(C.GLFWwindow), fc[2].(C.GLFWwindow), fc[3].(C.ulong))
				returnChan <- nil

			} else if fc[0] == "glfwGetWindowUserPointer" {
				r := C.glfwGetWindowUserPointer(fc[1].(C.GLFWwindow))
				returnChan <- []interface{}{r}

			} else if fc[0] == "glfwPollEvents" {
				C.glfwPollEvents()
				returnChan <- nil

			} else if fc[0] == "glfwWaitEvents" {
				C.glfwWaitEvents()
				returnChan <- nil

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

				// JUMP1
			} else if fc[0] == "glfwTerminate" {
				C.glfwTerminate()
				returnChan <- nil
				break // Kill the main loop since they terminated
			}

			runtime.UnlockOSThread() // Release main thread
		}
	}()

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


type Window struct {
	actual               C.GLFWwindow
	destroyed, closeable bool
	size_callback        func(int, int)
	close_callback       func()
	refresh_callback     func()
	focus_callback       func(bool)
	iconify_callback     func(bool)
	key_callback         func(int, int)
	char_callback        func(int)
	mouseButton_callback func(int, int)
	cursorPos_callback   func(int, int)
	cursorEnter_callback func(bool)
	scroll_callback      func(float32, float32)
    access               sync.RWMutex
}

func CreateWindow(width, height, mode int, title string, share *Window) (*Window, error) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	if share != nil {
		functionChan <- []interface{}{"glfwCreateWindow", C.int(width), C.int(height), C.int(mode), C.CString(title), share.actual}
	} else {
		functionChan <- []interface{}{"glfwCreateWindow", C.int(width), C.int(height), C.int(mode), C.CString(title), nil}
	}

	rv := <-returnChan
	actual := rv[0].(C.GLFWwindow)

    if actual == nil {
        return nil, errors.New("Unable to open an GLFW window")
    }

	w := Window{}
	w.actual = actual
	w.destroyed = false
	w.closeable = true
	functionChan <- []interface{}{"glfwSetWindowUserPointer", w.actual, unsafe.Pointer(&w)}
	<-returnChan
	return &w, nil
}

func (w *Window) Destroyed() bool {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.destroyed
}

func (w *Window) Destroy() {
    w.access.Lock()
    defer w.access.Unlock()

	if w.destroyed == false {
        glfwLock.Lock()
        defer glfwLock.Unlock()

		functionChan <- []interface{}{"glfwDestroyWindow", w.actual}
		<-returnChan
		w.destroyed = true
	}
}

func (w *Window) SetTitle(title string) {
    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSetWindowTitle", w.actual, C.CString(title)}
	<-returnChan
}

func (w *Window) SetSize(width, height int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSetWindowSize", w.actual, C.int(width), C.int(height)}
	<-returnChan
}

func (w *Window) Size() (int, int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	var width, height C.int
	functionChan <- []interface{}{"glfwGetWindowSize", w.actual, &width, &height}
	<-returnChan
	return int(width), int(height)
}

func (w *Window) SetPos(xpos, ypos int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSetWindowPos", w.actual, C.int(xpos), C.int(ypos)}
	<-returnChan
}

func (w *Window) Pos() (int, int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	var xpos, ypos C.int
	functionChan <- []interface{}{"glfwGetWindowPos", w.actual, &xpos, &ypos}
	<-returnChan
	return int(xpos), int(ypos)
}

func (w *Window) Iconify() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwIconifyWindow", w.actual}
	<-returnChan
}

func (w *Window) Restore() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwRestoreWindow", w.actual}
	<-returnChan
}

func (w *Window) Show() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwShowWindow", w.actual}
	<-returnChan
}

func (w *Window) Hide() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwHideWindow", w.actual}
	<-returnChan
}

func (w *Window) Param(param int) int {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwGetWindowParam", w.actual, C.int(param)}
	rv := <-returnChan
	return int(rv[0].(C.int))
}

func (w *Window) InputMode(mode int) int {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwGetInputMode", w.actual, C.int(mode)}
	rv := <-returnChan
	return int(rv[0].(C.int))
}

func (w *Window) SetInputMode(mode, value int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSetInputMode", w.actual, C.int(mode), C.int(value)}
	<-returnChan
}

func (w *Window) Key(key int) bool {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwGetKey", w.actual, C.int(key)}
	rv := <-returnChan
	return rv[0].(C.int) != 0
}

func (w *Window) MouseButton(button int) bool {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwGetMouseButton", w.actual, C.int(button)}
	rv := <-returnChan
	return rv[0].(C.int) != 0
}

func (w *Window) CursorPos() (int, int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	var xpos, ypos C.int
	functionChan <- []interface{}{"glfwGetCursorPos", w.actual, &xpos, &ypos}
	<-returnChan
	return int(xpos), int(ypos)
}

func (w *Window) SetCursorPos(xpos, ypos int) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSetCursorPos", w.actual, C.int(xpos), C.int(ypos)}
	<-returnChan
}

func (w *Window) ScrollOffset() (float32, float32) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	var xoffset, yoffset C.double
	functionChan <- []interface{}{"glfwGetScrollOffset", w.actual, &xoffset, &yoffset}
	<-returnChan
	return float32(xoffset), float32(yoffset)
}

func (w *Window) SetClipboardString(s string) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSetClipboardString", w.actual, C.CString(s)}
	<-returnChan
}

func (w *Window) ClipboardString() string {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwGetClipboardString", w.actual}
	r := <-returnChan
	return C.GoString(r[0].(*C.char))
}

func (w *Window) MakeContextCurrent() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwMakeContextCurrent", w.actual}
	<-returnChan
}

func (w *Window) SwapBuffers() {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

	functionChan <- []interface{}{"glfwSwapBuffers", w.actual}
	<-returnChan
}

func (w *Window) CopyContext(dst *Window, mask int64) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.RLock()
    defer w.access.RUnlock()

    dst.access.RLock()
    defer dst.access.RUnlock()

	functionChan <- []interface{}{"glfwCopyContext", w.actual, dst.actual, C.ulong(mask)}
	<-returnChan
}

// ------------------------------------------------------------------------- //
// Window-Object size callback
// ------------------------------------------------------------------------- //
func (w *Window) SetSizeCallback(f func(int, int)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.size_callback = f
}

func (w *Window) GetSizeCallback() func(int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.size_callback
}

// ------------------------------------------------------------------------- //
// Window-Object close callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCloseCallback(f func()) {
    w.access.Lock()
    defer w.access.Unlock()

	w.close_callback = f
}

func (w *Window) GetCloseCallback() func() {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.close_callback
}

func (w *Window) SetCloseable(closeable bool) {
    w.access.Lock()
    defer w.access.Unlock()

	w.closeable = closeable
}

func (w *Window) Closeable() bool {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.closeable
}

// ------------------------------------------------------------------------- //
// Window-Object refresh callback
// ------------------------------------------------------------------------- //
func (w *Window) SetRefreshCallback(f func()) {
    w.access.Lock()
    defer w.access.Unlock()

	w.refresh_callback = f
}

func (w *Window) GetRefreshCallback() func() {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.refresh_callback
}

// ------------------------------------------------------------------------- //
// Window-Object focus callback
// ------------------------------------------------------------------------- //
func (w *Window) SetFocusCallback(f func(bool)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.focus_callback = f
}

func (w *Window) GetFocusCallback() func(bool) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.focus_callback
}

// ------------------------------------------------------------------------- //
// Window-Object iconify callback
// ------------------------------------------------------------------------- //
func (w *Window) SetIconifyCallback(f func(bool)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.iconify_callback = f
}

func (w *Window) GetIconifyCallback() func(bool) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.iconify_callback
}

// ------------------------------------------------------------------------- //
// Window-Object key callback
// ------------------------------------------------------------------------- //
func (w *Window) SetKeyCallback(f func(int, int)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.key_callback = f
}

func (w *Window) GetKeyCallback() func(int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.key_callback
}

// ------------------------------------------------------------------------- //
// Window-Object char callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCharCallback(f func(int)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.char_callback = f
}

func (w *Window) GetCharCallback() func(int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.char_callback
}

// ------------------------------------------------------------------------- //
// Window-Object mouseButton callback
// ------------------------------------------------------------------------- //
func (w *Window) SetMouseButtonCallback(f func(int, int)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.mouseButton_callback = f
}

func (w *Window) GetMouseButtonCallback() func(int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.mouseButton_callback
}

// ------------------------------------------------------------------------- //
// Window-Object cursorPos callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCursorPosCallback(f func(int, int)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.cursorPos_callback = f
}

func (w *Window) GetCursorPosCallback() func(int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.cursorPos_callback
}

// ------------------------------------------------------------------------- //
// Window-Object cursorEnter callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCursorEnterCallback(f func(bool)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.cursorEnter_callback = f
}

func (w *Window) GetCursorEnterCallback() func(bool) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.cursorEnter_callback
}

// ------------------------------------------------------------------------- //
// Window-Object scroll callback
// ------------------------------------------------------------------------- //
func (w *Window) SetScrollCallback(f func(float32, float32)) {
    w.access.Lock()
    defer w.access.Unlock()

	w.scroll_callback = f
}

func (w *Window) GetScrollCallback() func(float32, float32) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.scroll_callback
}

// ------------------------------------------------------------------------- //
// Window from GLFWwindow
// ------------------------------------------------------------------------- //
func windowFromGLFW(w C.GLFWwindow) *Window {
	return (*Window)(C.glfwGetWindowUserPointer(w))
}

func windowFromGLFWChan(w C.GLFWwindow) *Window {
    glfwLock.Lock()
    defer glfwLock.Unlock()

	functionChan <- []interface{}{"glfwGetWindowUserPointer", w}
	r := <-returnChan
	return (*Window)(r[0].(unsafe.Pointer))
}

// ------------------------------------------------------------------------- //
// Global error callback
// ------------------------------------------------------------------------- //
var error_callback func(int, string)
var error_initCallback bool

//export error_doCallback
func error_doCallback(error C.int, description *C.char) {
	if error_callback != nil {
		go error_callback(int(error), C.GoString(description))
	}
}

func SetErrorCallback(f func(int, string)) {
	if !error_initCallback {
		C.error_initCallback()
		error_initCallback = true
	}
	error_callback = f
}

func GetErrorCallback() func(int, string) {
	return error_callback
}

// ------------------------------------------------------------------------- //
// Global size callback
// ------------------------------------------------------------------------- //
//export size_doCallback
func size_doCallback(win C.GLFWwindow, width, height C.int) {
	window := windowFromGLFW(win)
	if window.size_callback != nil {
		go window.size_callback(int(width), int(height))
	}
}

// ------------------------------------------------------------------------- //
// Global close callback
// ------------------------------------------------------------------------- //
//export close_doCallback
func close_doCallback(win C.GLFWwindow) bool {
	window := windowFromGLFW(win)
	if window.close_callback != nil {
		go window.close_callback()
		return window.Closeable()
	}
	return true
}

// ------------------------------------------------------------------------- //
// Global refresh callback
// ------------------------------------------------------------------------- //
//export refresh_doCallback
func refresh_doCallback(win C.GLFWwindow) {
	window := windowFromGLFW(win)
	if window.refresh_callback != nil {
		go window.refresh_callback()
	}
}

// ------------------------------------------------------------------------- //
// Global focus callback
// ------------------------------------------------------------------------- //
//export focus_doCallback
func focus_doCallback(win C.GLFWwindow, activated bool) {
	window := windowFromGLFW(win)
	if window.focus_callback != nil {
		go window.focus_callback(activated)
	}
}

// ------------------------------------------------------------------------- //
// Global iconify callback
// ------------------------------------------------------------------------- //
//export iconify_doCallback
func iconify_doCallback(win C.GLFWwindow, iconified bool) {
	window := windowFromGLFW(win)
	if window.iconify_callback != nil {
		go window.iconify_callback(iconified)
	}
}

// ------------------------------------------------------------------------- //
// Global key callback
// ------------------------------------------------------------------------- //
//export key_doCallback
func key_doCallback(win C.GLFWwindow, key, action C.int) {
	window := windowFromGLFW(win)
	if window.key_callback != nil {
		go window.key_callback(int(key), int(action))
	}
}

// ------------------------------------------------------------------------- //
// Global char callback
// ------------------------------------------------------------------------- //
//export char_doCallback
func char_doCallback(win C.GLFWwindow, character C.int) {
	window := windowFromGLFW(win)
	if window.char_callback != nil {
		go window.char_callback(int(character))
	}
}

// ------------------------------------------------------------------------- //
// Global mouseButton callback
// ------------------------------------------------------------------------- //
//export mouseButton_doCallback
func mouseButton_doCallback(win C.GLFWwindow, button, action int) {
	window := windowFromGLFW(win)
	if window.mouseButton_callback != nil {
		go window.mouseButton_callback(button, action)
	}
}

// ------------------------------------------------------------------------- //
// Global cursorPos callback
// ------------------------------------------------------------------------- //
//export cursorPos_doCallback
func cursorPos_doCallback(win C.GLFWwindow, x, y int) {
	window := windowFromGLFW(win)
	if window.cursorPos_callback != nil {
		go window.cursorPos_callback(x, y)
	}
}

// ------------------------------------------------------------------------- //
// Global cursorEnter callback
// ------------------------------------------------------------------------- //
//export cursorEnter_doCallback
func cursorEnter_doCallback(win C.GLFWwindow, entered bool) {
	window := windowFromGLFW(win)
	if window.cursorEnter_callback != nil {
		go window.cursorEnter_callback(entered)
	}
}

// ------------------------------------------------------------------------- //
// Global scroll callback
// ------------------------------------------------------------------------- //
//export scroll_doCallback
func scroll_doCallback(win C.GLFWwindow, x, y float32) {
	window := windowFromGLFW(win)
	if window.scroll_callback != nil {
		go window.scroll_callback(x, y)
	}
}

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
