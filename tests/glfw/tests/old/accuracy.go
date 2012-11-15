package main

import "code.google.com/p/azul3d"
import "code.google.com/p/azul3d/wrappers/glfw"
import "code.google.com/p/azul3d/wrappers/gl"
import "code.google.com/p/azul3d/wrappers/glu"
import "fmt"


var cursor_x, cursor_y int
var window_width, window_height int
var swap_interval int = 1
var window *glfw.Window

func set_swap_interval(interval int) {
    swap_interval = interval
    glfw.SwapInterval(swap_interval)
    fmt.Println(swap_interval)
    window.SetTitle(fmt.Sprintf("Cursor Inaccuracy Detector (interval %d)", swap_interval))
}

func window_size_callback(width, height int) {
    window_width = width
    window_height = height

    gl.Viewport(0, 0, window_width, window_height)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    glu.Ortho2D(0.0, float32(window_width), 0.0, float32(window_height))
}

func cursor_position_callback(x, y int) {
    cursor_x = x
    cursor_y = y
}

func key_callback(key, action int) {
    if key == glfw.KEY_SPACE && action == glfw.PRESS {
        set_swap_interval(1 - swap_interval)
    }
}

func main() {
    defer azul.Destroy()
    var err error

    window, err = glfw.CreateWindow(window_width, window_height, glfw.WINDOWED, "", nil)
    if err != nil {
        panic(err.Error())
    }
    window.SetCursorPosCallback(cursor_position_callback)
    window.SetSizeCallback(window_size_callback)
    window.SetKeyCallback(key_callback)

    window.MakeContextCurrent()

    width, height := window.Size()
    window_size_callback(width, height)

    set_swap_interval(swap_interval)


    for ; !window.Param(glfw.CLOSE_REQUESTED); {
        gl.Clear(gl.COLOR_BUFFER_BIT)

        gl.Begin(gl.LINES)
        gl.Vertex2f(0.0, float32(window_height - cursor_y))
        gl.Vertex2f(float32(window_width), float32(window_height - cursor_y))
        gl.Vertex2f(float32(cursor_x), 0.0)
        gl.Vertex2f(float32(cursor_x), float32(window_height))
        gl.End()

        window.SwapBuffers()
        glfw.PollEvents()
    }
}

