// +build tests

package main

import "code.google.com/p/azul3d/chippy"
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

    win, err := chippy.NewWindow(defaultScreen, &minAttribs, chippy.BestFBConfig)
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

    // It's very important to note, that any events you enable for a window
    // you are responsible for at least recieving them on the channel
    // all event channels are buffered, if you never receive items out of
    // the channel, the buffer will fill up, and you will encounter an deadlock.
    //
    // We will receieve events when:
    win.SetKeyEvents(true)          // keyboard buttons are pushed down or released
    win.SetCursorWithinEvents(true) // the mouse enters or exits the window
    win.SetCursorMoveEvents(true)   // the mouse cursor is moved (this event is only sent with window.SetMouseGrabbed(false))
    win.SetMouseMoveEvents(true)    // the mouse physically is moved (typically sub-pixel movement) (this event is ONLY sent with window.SetMouseGrabbed(true))
    win.SetResizeEvents(true)       // the window is resized
    win.SetMoveEvents(true)         // the window is moved 
    win.SetFocusEvents(true)        // the window loses or gains focus
    win.SetMinimizedEvents(true)    // the window becomes minimized or non-minimized
    win.SetMaximizedEvents(true)    // the window becomes maximized or non-maximized
    win.SetRedrawEvents(true)       // the window needs to be re-drawn
    win.SetAlwaysOnTopEvents(true)  // the user changed the always on top state of the window
    // Note, unless you specify that you will handle close events (the line below)
    // then Chippy will close the window by calling Destroy() for you when such an
    // event is received.
    win.SetCloseEvents(true)        // the user requests that we close the window

    closeRequested := false
    for !closeRequested{
        select{
            case ev := <-win.KeyEvent():
                //fmt.Println(ev.Window)
                if ev.Pressed {
                    fmt.Println("A key was pressed")
                } else {
                    fmt.Println("A key was released")
                }

            case ev := <-win.CursorWithinEvent():
                //fmt.Println(ev.Window)
                if ev.Within {
                    fmt.Println("Cursor enter window region")
                } else {
                    fmt.Println("Cursor exit window region")
                }

            case ev := <-win.CursorMoveEvent():
                //fmt.Println(ev.Window)
                fmt.Printf("Cursor moved to %dx%d\n", ev.X, ev.Y)

            case ev := <-win.MouseMoveEvent():
                //fmt.Println(ev.Window)
                fmt.Printf("Mouse move by %fx%f\n", ev.X, ev.Y)

            case ev := <-win.ResizeEvent():
                //fmt.Println(ev.Window)
                fmt.Printf("Window resized to %dx%d\n", ev.Width, ev.Height)

            case ev := <-win.MoveEvent():
                //fmt.Println(ev.Window)
                fmt.Printf("Window moved to %dx%d\n", ev.X, ev.Y)

            case ev := <-win.FocusEvent():
                //fmt.Println(ev.Window)
                if ev.Focus {
                    fmt.Println("Window gained focus")
                } else {
                    fmt.Println("Window lost focus")
                }

            case ev := <-win.MinimizedEvent():
                //fmt.Println(ev.Window)
                if ev.Minimized {
                    fmt.Println("Window minimized")
                } else {
                    fmt.Println("Window no longer minimized")
                }

            case ev := <-win.MaximizedEvent():
                //fmt.Println(ev.Window)
                if ev.Maximized {
                    fmt.Println("Window maximized")
                } else {
                    fmt.Println("Window no longer maximized")
                }

            case <-win.RedrawEvent():
                //fmt.Println(ev.Window)
                fmt.Println("Window redraw")

            case ev := <-win.AlwaysOnTopEvent():
                //fmt.Println(ev.Window)
                if ev.AlwaysOnTop {
                    fmt.Println("always on top? yes")
                    fmt.Println("GRABBING MOUSE")
                    win.SetMouseGrabbed(true)
                } else {
                    fmt.Println("always on top? no")
                    fmt.Println("RELEASING MOUSE")
                    win.SetMouseGrabbed(false)
                }

            case <-win.CloseEvent():
                //fmt.Println(ev.Window)
                fmt.Println("Window close!")
                closeRequested = true
        }
    }
}

