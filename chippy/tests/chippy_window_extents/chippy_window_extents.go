// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Test application - Tells the extents of an window's region and client region
package main

import(
    "code.google.com/p/azul3d/chippy"
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

    // Actually open the windows
    err = window.Open()
    if err != nil {
        log.Fatal(err)
    }

    // Print out what they currently has property-wise
    log.Println(window.Extents())
}

