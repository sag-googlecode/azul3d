// Test application - Opens two windows, changes each of their minimum size properties
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

    window1.SetMinimumSize(100, 100)

    // Actually open the windows
    err = window1.Open()
    if err != nil {
        log.Fatal(err)
    }

    err = window2.Open()
    if err != nil {
        log.Fatal(err)
    }

    window2.SetMinimumSize(100, 100)

    // Print out what they currently has property-wise
    log.Println(window1)
    log.Println(window2)

    // Just wait an while so they can enjoy the window
    log.Println("Waiting 30 seconds...")
    <- time.After(30 * time.Second)
}

