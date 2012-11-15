package main

import "code.google.com/p/azul3d/wrappers/glfw"
import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "runtime"
import "fmt"

var window *glfw.Window
var closed bool = false

func control_is_down(window *glfw.Window) (bool) {
    return window.Key(glfw.KEY_LEFT_CONTROL) || window.Key(glfw.KEY_RIGHT_CONTROL)
}

func window_close_callback() {
    closed = true
    window.SetCloseable(false)
}

func key_callback(key, action int) {
    if action != glfw.PRESS {
        return
    }

    switch key {
        case glfw.KEY_ESCAPE:
            closed = true
            break

        case glfw.KEY_V:
            if control_is_down(window) {
                s := window.ClipboardString()
                if len(s) != 0{
                    fmt.Printf("Clipboard contains \"%s\"\n", s)
                } else {
                    fmt.Printf("Clipboard does not contain a string")
                }
            }
            break

        case glfw.KEY_C:
            if control_is_down(window) {
                s := "Hello GLFW world!"
                window.SetClipboardString(s)
                fmt.Printf("Setting clipboard to \"%s\"\n", s)
            }
            break
    }
}

func window_size_callback(width, height int) {
    runtime.LockOSThread()
    gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
    runtime.UnlockOSThread()
}

func error_callback(error int, description string) {
    fmt.Printf("Error in %s", description)
}

func main() {
    var err error

    err = glfw.Init()
    if err != nil {
        panic(err.Error())
    }

    runtime.LockOSThread()
    err = gl.Init()
    if err != nil {
        panic(err.Error())
    }
    runtime.UnlockOSThread()

    glfw.SetErrorCallback(error_callback)

    window, err = glfw.CreateWindow(0, 0, glfw.WINDOWED, "Clipboard test", nil)
    if err != nil {
        panic(err.Error())
    }

    window.MakeContextCurrent()
    glfw.SwapInterval(1)

    window.SetKeyCallback(key_callback)
    window.SetSizeCallback(window_size_callback)
    window.SetCloseCallback(window_close_callback)

    runtime.LockOSThread()
    gl.MatrixMode(gl.PROJECTION)
    gl.Ortho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)
    gl.MatrixMode(gl.MODELVIEW)

    gl.ClearColor(0.5, 0.5, 0.5, 0)
    runtime.UnlockOSThread()

    for ; !closed; {
        runtime.LockOSThread()
        gl.Clear(gl.COLOR_BUFFER_BIT)

        gl.Color3f(0.8, 0.2, 0.4);
        gl.Rectf(-0.5, -0.5, 0.5, 0.5);
        runtime.UnlockOSThread()

        window.SwapBuffers()
        glfw.WaitEvents()
    }
}

