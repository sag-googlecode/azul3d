package math

import gmath "math"

func RadiansToDegrees(radians float64) float64 {
    return radians * (180.0 / gmath.Pi)
}

func DegreesToRadians(degrees float64) float64 {
    return degrees * (gmath.Pi / 180.0)
}

func FindPointOnCircle(cx, cy, radius, angle float64) (float64, float64) {
    nx := cx + radius * gmath.Cos(DegreesToRadians(angle))
    ny := cy + radius * gmath.Sin(DegreesToRadians(angle))
    return nx, ny
}

func DistanceBetweenPoints(x1, y1, x2, y2 float64) (float64) {
    d := gmath.Pow(x2 - x1, 2) + gmath.Pow(y2 - y1, 2)
    return gmath.Sqrt(d)
}
