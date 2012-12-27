// +build tests

package main

import(
    "code.google.com/p/azul3d/chippy"
    _ "image/png"
    "image"
    "time"
    "fmt"
    "os"
)

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



    fmt.Println("Cursor at 16x16")
	// Open the file.
	file, err := os.Open("data/cursor_16x16_2x2.png")
	if err != nil {
        panic(err)
	}
	defer file.Close()

	// Decode the image.
	m, _, err := image.Decode(file)
	if err != nil {
        panic(err)
	}
    win.SetCursorIcon(m, 2, 2)
    time.Sleep(t)



    fmt.Println("Cursor show/hide")
    win.SetCursorHidden(true)
    win.SetCursorHidden(false)
    time.Sleep(t)



    fmt.Println("Cursor at 32x32")
	// Open the file.
	file, err = os.Open("data/cursor_32x32_4x3.png")
	if err != nil {
        panic(err)
	}
	defer file.Close()

	// Decode the image.
	m, _, err = image.Decode(file)
	if err != nil {
        panic(err)
	}
    win.SetCursorIcon(m, 4, 3)
    time.Sleep(t)



    fmt.Println("Cursor show/hide")
    win.SetCursorHidden(true)
    win.SetCursorHidden(false)
    time.Sleep(t)



    fmt.Println("Cursor at 48x48")
	// Open the file.
	file, err = os.Open("data/cursor_48x48_6x5.png")
	if err != nil {
        panic(err)
	}
	defer file.Close()

	// Decode the image.
	m, _, err = image.Decode(file)
	if err != nil {
        panic(err)
	}
    win.SetCursorIcon(m, 6, 5)
    time.Sleep(t)



    fmt.Println("Cursor show/hide")
    win.SetCursorHidden(true)
    win.SetCursorHidden(false)
    time.Sleep(t)



    fmt.Println("Making cursor the original one")
    win.SetCursorIcon(nil, 0, 0)
    time.Sleep(t)



    fmt.Println("Color cursor 48x48")
	// Open the file.
	file, err = os.Open("data/cursor_color_48x48.png")
	if err != nil {
        panic(err)
	}
	defer file.Close()

	// Decode the image.
	m, _, err = image.Decode(file)
	if err != nil {
        panic(err)
	}
    win.SetCursorIcon(m, 0, 0)
    time.Sleep(t)
}

