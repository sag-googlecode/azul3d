package clock

import "sync"
import "math"
import "time"
import "fmt"

type Clock struct {
    numFrames uint64
    startTime time.Time
    lastFrameTime time.Time
    delta float64
    maxDelta float64

    framesElapsed int
    lastFpsTime time.Time
    fps float64

    access sync.RWMutex
}

func New() *Clock {
    c := new(Clock)
    c.startTime = time.Now()
    c.maxDelta = 0.2
    return c
}

// String returns a string representation of this Clock
func (this *Clock) String() string {
    return fmt.Sprintf("Clock(FPS=%f, delta=%f, frames=%d, seconds=%f)", this.Fps(), this.Delta(), this.Frames(), this.Seconds())
}

// Tick informs this clock that one single frame has rendered
func (this *Clock) Tick() {
    this.access.Lock()
    defer this.access.Unlock()

    currentTime := time.Now()

    this.numFrames += 1

    this.delta = currentTime.Sub(this.lastFrameTime).Seconds()
    this.lastFrameTime = currentTime

    this.framesElapsed += 1
    elapsed := currentTime.Sub(this.lastFpsTime).Seconds()
    if elapsed > 1.0 {
        this.fps = float64(this.framesElapsed) / elapsed
        this.lastFpsTime = currentTime
        this.framesElapsed = 0
    }
}

// Fps returns the number of frames rendered in a single second
func (this *Clock) Fps() float64 {
    this.access.RLock()
    defer this.access.RUnlock()

    return this.fps
}

// Delta returns the time in seconds that the last frame took to render
func (this *Clock) Delta() float64 {
    this.access.RLock()
    defer this.access.RUnlock()

    return math.Min(this.delta, this.maxDelta)
}

// MaxDelta returns the maximum time in seconds that the Delta can be (0.2 by default)
func (this *Clock) MaxDelta() float64 {
    this.access.RLock()
    defer this.access.RUnlock()

    return this.maxDelta
}

// SetMaxDelta sets the maximum time in seconds that the Delta can be (0.2 by default)
func (this *Clock) SetMaxDelta(max float64) {
    this.access.Lock()
    defer this.access.Unlock()

    this.maxDelta = max
}

// Frames returns the number of frames rendered since clock startup
func (this *Clock) Frames() uint64 {
    this.access.RLock()
    defer this.access.RUnlock()

    return this.numFrames
}

// Seconds returns the number of seconds passed since clock startup
func (this *Clock) Seconds() float64 {
    this.access.RLock()
    defer this.access.RUnlock()

    duration := time.Now().Sub(this.startTime)
    return duration.Seconds()
}

