package main

import "code.google.com/p/azul3d/math"
import "fmt"

func main() {
    fmt.Println(math.DegreesToRadians(90))
    fmt.Println(math.RadiansToDegrees(math.DegreesToRadians(90)))
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")
    var x, y, radius, angle float64

    x, y = 0, 0

    radius = 3
    angle = 90
    fmt.Printf("(%f, %f) angle=%f, radius=%f\n", x, y, angle, radius)
    nx, ny := math.FindPointOnCircle(x, y, radius, angle)
    fmt.Println(int(nx), int(ny))
    fmt.Println("")

    radius = 3
    angle = -90
    fmt.Printf("(%f, %f) angle=%f, radius=%f\n", x, y, angle, radius)
    fmt.Println(math.FindPointOnCircle(x, y, radius, angle))
    fmt.Println("")

    radius = 3
    angle = 45
    fmt.Printf("(%f, %f) angle=%f, radius=%f\n", x, y, angle, radius)
    fmt.Println(math.FindPointOnCircle(x, y, radius, angle))
    fmt.Println("")

    radius = 3
    angle = 180
    fmt.Printf("(%f, %f) angle=%f, radius=%f\n", x, y, angle, radius)
    fmt.Println(math.FindPointOnCircle(x, y, radius, angle))
    fmt.Println("")

    radius = 3
    angle = -180
    fmt.Printf("(%f, %f) angle=%f, radius=%f\n", x, y, angle, radius)
    fmt.Println(math.FindPointOnCircle(x, y, radius, angle))
    fmt.Println("")

    fmt.Println(math.DistanceBetweenPoints(0, 0, -2, -2))
}
