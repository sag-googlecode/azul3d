package main

import "code.google.com/p/azul3d/clock"

import "math/rand"
import "math"
import "time"
import "fmt"

func main() {
    c := clock.New()
    for{
        randomFps := 240 * rand.Float64()
        randomFps = math.Max(randomFps, 30)

        <-time.After(time.Second / time.Duration(randomFps))
        c.Tick()
        fmt.Println(c)
    }
}

