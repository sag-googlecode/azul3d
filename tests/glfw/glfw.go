package main

import "code.google.com/p/azul3d"
import "code.google.com/p/azul3d/wrappers/glfw3"
import "runtime"
import "fmt"

func sizeCallback(w *glfw.Window, x, y int) {
    fmt.Println(w, x, y)
    w.Destroy() // Seriously, never do this because your users will hate you..
}

func closeCallback(w *glfw.Window) (bool) {
    fmt.Println("close", w)
    return true // Window is closeable
}

func refreshCallback(w *glfw.Window) {
    fmt.Println("refresh", w)
}

func focusCallback(w *glfw.Window, activated bool) {
    fmt.Println("focus", w, activated)
}

func iconifyCallback(w *glfw.Window, iconified bool) {
    fmt.Println("iconify", w, iconified)
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    defer azul.Destroy()
    fmt.Println(glfw.VERSION_MAJOR)
    fmt.Println(glfw.Version())
    fmt.Println(glfw.VersionString())

    glfw.SetErrorCallback(func(error int, description string) {
        fmt.Println(error, description)
    })

    w, err := glfw.CreateWindow(640, 480, glfw.WINDOWED, "Window title goes here", nil)
    if err != nil {
        panic(err.Error())
    }
    defer w.Destroy()

    w.SetSizeCallback(func(width, height int) {
        fmt.Println("size", width, height)
    })

    w.SetCloseCallback(func() {
        fmt.Println("close")
    })

    w.SetRefreshCallback(func() {
        fmt.Println("refresh")
    })

    w.SetFocusCallback(func(focused bool) {
        fmt.Println("focused:", focused)
    })

    w.SetIconifyCallback(func(iconified bool) {
        fmt.Println("iconified:", iconified)
    })

    w.SetKeyCallback(func(key, action int) {
        fmt.Println("key,action:", key, action)
    })

    w.SetCharCallback(func(character string) {
        fmt.Println("char:", character)
    })

    w.SetMouseButtonCallback(func(button, action int) {
        fmt.Println("button, action:", button, action)
    })

    w.SetCursorPosCallback(func(x, y int) {
        fmt.Println("cursorPos:", x, y)
    })

    w.SetCursorEnterCallback(func(entered bool) {
        fmt.Println("cursorEnter:", entered)
    })

    w.SetScrollCallback(func(x, y float32) {
        fmt.Println("scroll:", x, y)
    })

    fmt.Println("Querying DesktopMode()")

    fmt.Println(glfw.DesktopMode().Width, glfw.DesktopMode().Height)

    vidModes := glfw.VideoModes()
    for i := 0; i < len(vidModes); i++ {
        if vidModes[i].Equals(glfw.DesktopMode()) {
            fmt.Println(vidModes[i].Width, vidModes[i].Height, "(Desktop mode)")
        } else {
            fmt.Println(vidModes[i].Width, vidModes[i].Height)
        }
    }

    for ; !w.Destroyed(); {
        glfw.WaitEvents()
        if w.Key(glfw.KEY_A) {
            w.SetTitle(fmt.Sprintf("GetTime() = %f", glfw.Time()))
            fmt.Println("go routines", runtime.NumGoroutine())
        }
        if w.Key(glfw.KEY_Q) || w.Param(glfw.CLOSE_REQUESTED){
            w.Destroy()
        }
    }
}
