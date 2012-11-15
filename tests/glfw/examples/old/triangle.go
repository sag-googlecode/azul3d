package main

import "code.google.com/p/azul3d"
import "code.google.com/p/azul3d/wrappers/glfw"
import "code.google.com/p/azul3d/wrappers/gl"
import "code.google.com/p/azul3d/wrappers/glu"

func main() {
    defer azul.Destroy()

    w, err := glfw.CreateWindow(640, 480, glfw.WINDOWED, "Spinning Triangle", nil)
    if err != nil {
        panic(err.Error())
    }
    defer w.Destroy()

    var x, width, height int

    w.MakeContextCurrent()
    glfw.SwapInterval(1)

    for{
        t := glfw.Time()
        x, _ = w.CursorPos()

        // Get window size (may be different than the requested size)
        width, height = w.Size()

        // Special case: avoid division by zero below
        if height < 0 {
            height = 1
        }

        gl.Viewport(0, 0, width, height)

        // Clear color buffer to black
        gl.ClearColor(0.0, 0.0, 0.0, 0.0)
        gl.Clear(gl.COLOR_BUFFER_BIT)

        // Select and setup the projection matrix
        gl.MatrixMode(gl.PROJECTION)
        gl.LoadIdentity()
        glu.Perspective(65.0, float32(width) / float32(height), 1.0, 100.0)

        // Select and setup the modelview matrix
        gl.MatrixMode(gl.MODELVIEW)
        gl.LoadIdentity()
        glu.LookAt(0.0, 1.0, 0.0,  // Eye-position
                  0.0, 20.0, 0.0,  // View-point
                  0.0, 0.0, 1.0)   // Up-vector

        // Draw a rotating colorful triangle
        gl.Translatef(0.0, 14.0, 0.0)
        gl.Rotatef(0.3 * float32(x) + float32(t) * 100.0, 0.0, 0.0, 1.0);

        gl.Begin(gl.TRIANGLES)
        gl.Color3f(1.0, 0.0, 0.0)
        gl.Vertex3f(-5.0, 0.0, -4.0)
        gl.Color3f(0.0, 1.0, 0.0)
        gl.Vertex3f(5.0, 0.0, -4.0)
        gl.Color3f(0.0, 0.0, 1.0)
        gl.Vertex3f(0.0, 0.0, 6.0)
        gl.End()

        // Swap buffers
        w.SwapBuffers()
        glfw.PollEvents()

        // Check if the ESC key was pressed or the window should be closed
        if w.Key(glfw.KEY_ESCAPE) || w.Param(glfw.CLOSE_REQUESTED) {
            break
        }
    }
}

