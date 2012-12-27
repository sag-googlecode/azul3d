// +build tests

package main

import "code.google.com/p/azul3d/chippy"
import "time"
import "fmt"

func main() {
    err := chippy.Init()
    if err != nil {
        panic(err.Error())
    }
    defer chippy.Destroy()

    minAttribs := chippy.FBConfig{
        RedBits: 1,
        BlueBits: 1,
        GreenBits: 1,
        AlphaBits: 1,
        DepthBits: 0,
        StencilBits: 0,
        Samples: 0,
        SampleBuffers: 0,
        DoubleBuffered: false,
    }

    defaultScreen, err := chippy.DefaultScreen()
    if err != nil {
        panic(err.Error())
    }

    win, err := chippy.NewWindow(defaultScreen, &minAttribs, chippy.BestFBConfig, nil)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Opened a window with these Frame Buffer configurations:")
    fmt.Println(win.FBConfig())
    contextVersion, err := win.ContextVersionString()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("The window is capable of OpenGL", contextVersion)

    t := 5 * time.Second
    time.Sleep(t)

    fmt.Println("Max size 640x480")
    win.SetMaxSize(640, 480)
    time.Sleep(t)

    fmt.Println("Max size 300x300")
    win.SetMaxSize(300, 300)
    time.Sleep(t)

    fmt.Println("Max size 150x150")
    win.SetMaxSize(150, 150)
    time.Sleep(t)

    fmt.Println("Max size 800x600")
    win.SetMaxSize(800, 600)
    time.Sleep(t)
}

