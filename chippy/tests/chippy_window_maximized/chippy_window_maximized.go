// Test application - Opens two windows, changes each of their minimized properties
package main

import(
    "code.google.com/p/azul3d/chippy"
    "time"
    "log"
    "os"
)

func main() {
    log.SetFlags(0)

    // Enable debug output
    chippy.SetDebugOutput(os.Stdout)

    err := chippy.Init()
    if err != nil {
        log.Fatal(err)
    }
    defer chippy.Destroy()


    window1 := chippy.NewWindow()
    window2 := chippy.NewWindow()

    window1.SetMaximized(true)

    // Actually open the windows
    err = window1.Open()
    if err != nil {
        log.Fatal(err)
    }

    err = window2.Open()
    if err != nil {
        log.Fatal(err)
    }

    window2.SetMaximized(true)

    // Print out what they currently has property-wise
    log.Println(window1)
    log.Println(window2)

    log.Println("Waiting 5 seconds...")
    <- time.After(5 * time.Second)

    window1.SetMaximized(false)
    window2.SetMaximized(false)

    // Just wait an while so they can enjoy the window
    log.Println("Waiting 15 seconds...")
    <- time.After(15 * time.Second)
}

