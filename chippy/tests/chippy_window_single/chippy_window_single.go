// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Opens an single window on the specified screen
package main

// Note: On Windows build with:
//   go install -ldflags "-H windowsgui" path/to/pkg
// to hide the command prompt

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


    window := chippy.NewWindow()

    // Actually open the window
    err = window.Open()
    if err != nil {
        log.Fatal(err)
    }

    // Print out what it currently has property-wise
    log.Println(window)

    // Just wait an while so they can enjoy the window
    log.Println("Waiting 15 seconds...")
    <- time.After(15 * time.Second)
}

