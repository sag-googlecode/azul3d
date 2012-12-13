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

    clipboard, err := chippy.Clipboard()
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("Clipboard is: \"%s\"\n", clipboard)


    fmt.Println("Setting clipboard to \"Chippy is awesome\"")
    chippy.SetClipboard("Chippy is awesome")


    clipboard, err = chippy.Clipboard()
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("Clipboard is: \"%s\"\n", clipboard)
}

