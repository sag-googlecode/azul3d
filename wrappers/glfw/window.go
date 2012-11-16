package glfw

/*
#cgo CFLAGS: -DGLFW_INCLUDE_GLU
#cgo LDFLAGS: -lglfw
#include <GL/glfw3.h>
#include <stdlib.h>
*/
import "C"

import "unsafe"
import "errors"
import "sync"


type Window struct {
	actual               C.GLFWwindow
	destroyed, closeable bool
	size_callback        func(*Window, int, int)
	close_callback       func(*Window, )
	refresh_callback     func(*Window, )
	focus_callback       func(*Window, bool)
	iconify_callback     func(*Window, bool)
	key_callback         func(*Window, int, int)
	char_callback        func(*Window, int)
	mouseButton_callback func(*Window, int, int)
	cursorPos_callback   func(*Window, int, int)
	cursorEnter_callback func(*Window, bool)
	scroll_callback      func(*Window, float32, float32)
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
    // This is the same as glfwSwapBuffers, it specifies to use the thread-local context
    // so this call is intentionally non thread-safe. Only use it from the thread you
    // make OpenGL calls from.
    C.glfwMakeContextCurrent(w.actual)
}

func (w *Window) SwapBuffers() {
    // glfwSwapBuffers uses thread-local data, specifically local to OpenGL calls
    // so this call needs to be ran in the same thread as your OpenGL calls.
    // meaning that this call is about as thread safe as OpenGL: it never is.
    C.glfwSwapBuffers(w.actual)
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
func (w *Window) SetSizeCallback(f func(*Window, int, int)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.size_callback = f

	functionChan <- []interface{}{"size_callback", w.actual}
	<-returnChan
}

func (w *Window) GetSizeCallback() func(*Window, int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.size_callback
}

// ------------------------------------------------------------------------- //
// Window-Object close callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCloseCallback(f func(*Window)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.close_callback = f

	functionChan <- []interface{}{"close_callback", w.actual}
	<-returnChan
}

func (w *Window) GetCloseCallback() func(*Window) {
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
func (w *Window) SetRefreshCallback(f func(*Window)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.refresh_callback = f

	functionChan <- []interface{}{"refresh_callback", w.actual}
	<-returnChan
}

func (w *Window) GetRefreshCallback() func(*Window) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.refresh_callback
}

// ------------------------------------------------------------------------- //
// Window-Object focus callback
// ------------------------------------------------------------------------- //
func (w *Window) SetFocusCallback(f func(*Window, bool)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.focus_callback = f

	functionChan <- []interface{}{"focus_callback", w.actual}
	<-returnChan
}

func (w *Window) GetFocusCallback() func(*Window, bool) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.focus_callback
}

// ------------------------------------------------------------------------- //
// Window-Object iconify callback
// ------------------------------------------------------------------------- //
func (w *Window) SetIconifyCallback(f func(*Window, bool)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.iconify_callback = f

	functionChan <- []interface{}{"iconify_callback", w.actual}
	<-returnChan
}

func (w *Window) GetIconifyCallback() func(*Window, bool) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.iconify_callback
}

// ------------------------------------------------------------------------- //
// Window-Object key callback
// ------------------------------------------------------------------------- //
func (w *Window) SetKeyCallback(f func(*Window, int, int)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.key_callback = f

	functionChan <- []interface{}{"key_callback", w.actual}
	<-returnChan
}

func (w *Window) GetKeyCallback() func(*Window, int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.key_callback
}

// ------------------------------------------------------------------------- //
// Window-Object char callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCharCallback(f func(*Window, int)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.char_callback = f

	functionChan <- []interface{}{"char_callback", w.actual}
	<-returnChan
}

func (w *Window) GetCharCallback() func(*Window, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.char_callback
}

// ------------------------------------------------------------------------- //
// Window-Object mouseButton callback
// ------------------------------------------------------------------------- //
func (w *Window) SetMouseButtonCallback(f func(*Window, int, int)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.mouseButton_callback = f

	functionChan <- []interface{}{"mouseButton_callback", w.actual}
	<-returnChan
}

func (w *Window) GetMouseButtonCallback() func(*Window, int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.mouseButton_callback
}

// ------------------------------------------------------------------------- //
// Window-Object cursorPos callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCursorPosCallback(f func(*Window, int, int)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.cursorPos_callback = f

	functionChan <- []interface{}{"cursorPos_callback", w.actual}
	<-returnChan
}

func (w *Window) GetCursorPosCallback() func(*Window, int, int) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.cursorPos_callback
}

// ------------------------------------------------------------------------- //
// Window-Object cursorEnter callback
// ------------------------------------------------------------------------- //
func (w *Window) SetCursorEnterCallback(f func(*Window, bool)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.cursorEnter_callback = f

	functionChan <- []interface{}{"cursorEnter_callback", w.actual}
	<-returnChan
}

func (w *Window) GetCursorEnterCallback() func(*Window, bool) {
    w.access.RLock()
    defer w.access.RUnlock()

	return w.cursorEnter_callback
}

// ------------------------------------------------------------------------- //
// Window-Object scroll callback
// ------------------------------------------------------------------------- //
func (w *Window) SetScrollCallback(f func(*Window, float32, float32)) {
    glfwLock.Lock()
    defer glfwLock.Unlock()

    w.access.Lock()
    defer w.access.Unlock()

	w.scroll_callback = f

	functionChan <- []interface{}{"scroll_callback", w.actual}
	<-returnChan
}

func (w *Window) GetScrollCallback() func(*Window, float32, float32) {
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

