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

    screens, err := chippy.Screens()
    if err != nil {
        panic(err.Error())
    }

    defaultScreen, err := chippy.DefaultScreen()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("There are", len(screens), "screens")
    fmt.Println("Default screen:", defaultScreen)
    fmt.Println("Screens", screens)
}

