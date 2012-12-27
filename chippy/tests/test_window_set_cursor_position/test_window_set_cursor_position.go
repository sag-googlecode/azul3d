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

    fmt.Println("Cursor moved to 0x0")
    win.SetCursorPosition(0, 0)
    time.Sleep(t)

    fmt.Println("Cursor moved to 50x50")
    win.SetCursorPosition(50, 50)
    time.Sleep(t)

    fmt.Println("Cursor moved to 65535x65535")
    win.SetCursorPosition(65535, 65535)
    time.Sleep(t)
}

