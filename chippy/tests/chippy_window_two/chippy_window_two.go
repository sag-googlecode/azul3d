// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an two windows, each on the specified screens
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

    // Actually open the windows
    err = window1.Open()
    if err != nil {
        log.Fatal(err)
    }

    err = window2.Open()
    if err != nil {
        log.Fatal(err)
    }

    // Print out what they currently has property-wise
    log.Println(window1)
    log.Println(window2)

    // Just wait an while so they can enjoy the window
    log.Println("Waiting 15 seconds...")
    <- time.After(15 * time.Second)
}

