// +build tests

package main

import "code.google.com/p/azul3d/chippy"
import "time"
import "fmt"

func main() {
    err := chippy.Init()
    if err != nil {
        panic(err.Error())
    }
    defer chippy.Destroy()

    screen, err := chippy.DefaultScreen()
    if err != nil {
        panic(err)
    }

    do := func(x float64) {
        fmt.Println(x)
        err := screen.SetGamma(float32(x))
        if err != nil {
            // Unable to set gamma
            panic(err.Error())
        }
        time.Sleep(1 * time.Millisecond)
        gamma, err := screen.Gamma()
        if err != nil {
            // Unable to get gamma
            panic(err.Error())
        }
        fmt.Println("Gamma is", gamma)
    }

    for x := 0; x < 3; x ++ {
        for i := 1.0; i <= 2.0; i += 0.01 {
            do(i)
        }
        for i := 2.0; i >= 0.0; i -= 0.01 {
            do(i)
        }
        for i := 0.0; i <= 1.0; i += 0.01 {
            do(i)
        }
        do(1.0)
    }

    do(0.3) // ensure restore works
    time.Sleep(1 * time.Second)
    //screen.SetAutoRestoreOriginalGamma(false)
    //do(1)
}

