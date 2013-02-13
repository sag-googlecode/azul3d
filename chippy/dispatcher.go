package chippy

// dispatch() and dispatchRequests() are used by certain backends to implement thread-safe calls,
// in cases where all window manager specific calls, must be made on the same OS thread.
//
// Whichever OS thread dispatchRequests() is called on will be designated to execute all functions
// passed into dispatch().
//
// dispatchRequests() will never return and will hold the OS thread (using runtime.LockOSThread)
// until stopDispatching() is called, at which point the OS thread is released, and the function
// returns.
//
// (obviously) dispatch() will only ever execute one function at an time.
// dispatch() will only return once the function is done executing.

import(
    "runtime"
)

type request struct{
    funcChan chan func()
    completedChan chan bool
}

var(
    requestChan = make(chan *request, 5)
    shutdownChan = make(chan bool)
)

func dispatchRequests() {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()
    for{
        select{
            case r := <-requestChan:
                action := <-r.funcChan
                action()
                r.completedChan <- true

            case <-shutdownChan:
                return
        }
    }
}

func stopDispatching() {
    shutdownChan <- true
}

// Dispatches the function on the dispatcher thread, and waits for the operation to complete before
// returning.
func dispatch(f func()) {
    r := &request{
        make(chan func(), 1),
        make(chan bool, 1),
    }
    requestChan <- r
    r.funcChan <- f
    <- r.completedChan
}

// Dispatches the function on the dispatcher thread, and returns immedietly without waiting for the
// operation to complete.
func dispatchNoWait(f func()) {
    r := &request{
        make(chan func(), 1),
        make(chan bool, 1),
    }
    requestChan <- r
    r.funcChan <- f
    //<- r.completedChan
}

