package render

import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "code.google.com/p/azul3d/wrappers/glfw"
//import "code.google.com/p/azul3d/wrappers/glu"
import "code.google.com/p/azul3d/clock"

import "runtime"
import "sync"

func Init() error {
    return gl.Init()
}

type Dispatcher struct {
    window *glfw.Window
    dispatching sync.Mutex
    clock *clock.Clock

    // Triple buffering
    frameChan chan func()
}

func NewDispatcher(w *glfw.Window, clock *clock.Clock) *Dispatcher {
    d := Dispatcher{}
    d.window = w
    d.clock = clock
    d.frameChan = make(chan func(), 1)
    go d.dispatch()
    return &d
}

func (this *Dispatcher) NewFrame(f func()) {
    this.frameChan <- f
}

func (this *Dispatcher) Start() {
    this.dispatching.Unlock()
}

func (this *Dispatcher) Stop() {
    this.dispatching.Lock()
}

func (this *Dispatcher) dispatch() {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    this.window.MakeContextCurrent()
    glfw.SwapInterval(1)

    for{
        runtime.Gosched()
        if this.window.Destroyed() {
            break
        }
        // We actually just want to block here in the case that
        // Stop() was called
        //this.dispatching.Lock()
        //this.dispatching.Unlock()

        // Draw the frame
        frameCallback := <-this.frameChan
        frameCallback()

        // Swap buffers
        this.window.SwapBuffers()
        this.clock.Tick()
    }
}

