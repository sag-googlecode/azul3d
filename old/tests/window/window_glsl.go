package main

import "code.google.com/p/azul3d"
import "code.google.com/p/azul3d/window"
import "code.google.com/p/azul3d/wrappers/glu"

import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "code.google.com/p/azul3d/render/glutil"

import "runtime"
import "fmt"


func main() {
    runtime.GOMAXPROCS(8)

    defer azul.Destroy()

    w, err := window.New()
    if err != nil {
        panic(err.Error())
    }

    var vbo *glutil.VertexBuffer

    w.NewFrame(func() {
        fmt.Println("GL VERSION", gl.GoStringUb(gl.GetString(gl.VERSION)) )

        var err error

        prog := glutil.NewProgram()

        vert := glutil.NewShader()
        err = vert.LoadFile("shader.vert", glutil.VERTEX_SHADER)
        if err != nil {
            // IO errors
            panic(err.Error())
        }

        if vert.HasError() {
            fmt.Println(glutil.GetErrors())
            panic(fmt.Sprintf("Vertex shader error:\n%s", vert.ErrorString()))
        }
        prog.AddShader(vert)

        frag := glutil.NewShader()
        err = frag.LoadFile("shader.frag", glutil.FRAGMENT_SHADER)
        if err != nil {
            // IO errors
            panic(err.Error())
        }

        if frag.HasError() {
            fmt.Println(glutil.GetErrors())
            panic(fmt.Sprintf("Frament shader error:\n%s", frag.ErrorString()))
        }
        prog.AddShader(frag)


        prog.Compile()
        if prog.HasError() {
            fmt.Println(glutil.GetErrors())
            panic(fmt.Sprintf("Error:\n%s", prog.ErrorString()))
        }
        prog.Use() //glUseProgram()


        vbo = glutil.NewVertexBuffer()
        //Vertices of a triangle (counter-clockwise winding)
        vbo.SetData([]float32{1.0, 0.0, 1.0, 0.0, 0.0, -1.0, -1.0, 0.0, 1.0}, glutil.STATIC_DRAW)
    })


    for !w.Destroyed(){
        runtime.Gosched()

        fmt.Println("fps", w.Clock().Fps())

        // Get window size (may be different than the requested size)
        width, height := w.Size()

        // Special case: avoid division by zero below
        if height < 0 {
            height = 1
        }

        w.NewFrame(func() {
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

            gl.Begin(gl.TRIANGLES)
            gl.Color3f(1.0, 0.0, 0.0)
            gl.Vertex3f(-5.0, 0.0, -4.0)
            gl.Color3f(0.0, 1.0, 0.0)
            gl.Vertex3f(5.0, 0.0, -4.0)
            gl.Color3f(0.0, 0.0, 1.0)
            gl.Vertex3f(0.0, 0.0, 6.0)
            gl.End()
        })

        if w.CloseRequested() {
            w.Destroy()
        }
    }
}
