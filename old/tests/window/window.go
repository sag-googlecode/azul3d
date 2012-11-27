package main

import "code.google.com/p/azul3d"
import "code.google.com/p/azul3d/window"
import "code.google.com/p/azul3d/wrappers/glu"
import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "runtime"
import "fmt"

func main() {
    runtime.GOMAXPROCS(8)

    defer azul.Destroy()

    w, err := window.New()
    if err != nil {
        panic(err.Error())
    }

    var rotation float64

    for !w.Destroyed(){
        runtime.Gosched()

        fmt.Println("fps", w.Clock().Fps())
        frame := w.NewFrame()

        rotation += 60.0 * w.Clock().Delta() // degrees/sec

        // Get window size (may be different than the requested size)
        width, height := w.Size()

        // Special case: avoid division by zero below
        if height < 0 {
            height = 1
        }

        frame.Add(func() {
            gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))

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
            gl.Rotatef(gl.Float(rotation), 0.0, 0.0, 1.0);

            gl.Begin(gl.TRIANGLES)
            gl.Color3f(1.0, 0.0, 0.0)
            gl.Vertex3f(-5.0, 0.0, -4.0)
            gl.Color3f(0.0, 1.0, 0.0)
            gl.Vertex3f(5.0, 0.0, -4.0)
            gl.Color3f(0.0, 0.0, 1.0)
            gl.Vertex3f(0.0, 0.0, 6.0)
            gl.End()
        })

        frame.Finish()

        if w.CloseRequested() {
            w.Destroy()
        }
    }
}
