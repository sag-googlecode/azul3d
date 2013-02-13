package chippy

import(
    "code.google.com/p/azul3d/chippy/wrappers/win32"
    "errors"
    "sync"
    "fmt"
)

func eventLoop() {
    timerId := 1

    // Be careful! calling GetMessageSync is ran on the same dispatching thread as all other win32 api calls!
    // so other API calls in our case will never complete untill after this times out!
    //
    // This is here merely to throttle down hogging all the CPU when we really shouldn't be.
    //
    // Wish windows let you poll this as an file descriptor.. sigh..
    timeoutMs := 5

    ret := true
    for ret {
        var(
            timer win32.UINT_PTR
            msg *win32.MSG
        )

        dispatch(func() {
            //ret, msg = win32.PeekMessage(nil, 0, 0, win32.PM_REMOVE|win32.PM_NOYIELD)

            timer = win32.SetTimer(nil, win32.UINT_PTR(timerId), win32.UINT(timeoutMs), nil)
            ret, msg = win32.GetMessage(nil, 0, 0)

            if !win32.KillTimer(nil, timer) {
                panic(fmt.Sprintf("Unable to KillTimer():", win32.GetLastErrorString()))
            }
        })


        if ret {
            if msg.Message() == win32.WM_TIMER && msg.Hwnd() == nil {
                continue
            }

            dispatch(func() {
                win32.TranslateMessage(msg)
                win32.DispatchMessage(msg)
            })
        }
    }
}


var classNameCounter = 0
var classNameCounterAccess sync.Mutex
func nextCounter() int {
    classNameCounterAccess.Lock()
    defer classNameCounterAccess.Unlock()
    classNameCounter++
    return classNameCounter
}


//var classAtom win32.ATOM
//var windowClass *win32.WNDCLASSEX
var hInstance win32.HINSTANCE

var w32VersionMajor, w32VersionMinor win32.DWORD


func backend_Init() error {
    go dispatchRequests()

    var err error
    dispatch(func() {
        hInstance = win32.HINSTANCE(win32.GetModuleHandle(""))
        if hInstance == nil {
            err = errors.New(fmt.Sprintf("Unable to determine hInstance; GetModuleHandle():", win32.GetLastErrorString()))
			return
        }

		// Get OS version, we use this to do some hack-ish fixes for different windows versions
		ret, vi := win32.GetVersionEx()
		if ret {
			w32VersionMajor = vi.DwMajorVersion()
			w32VersionMinor = vi.DwMinorVersion()
		} else {
			err = errors.New(fmt.Sprintf("Unable to determine windows version information; GetVersionEx():", win32.GetLastErrorString()))
			return
		}

    })

    if err != nil {
        return err
    }

    go eventLoop()

	return nil
}

func backend_Destroy() {
    stopDispatching()

    classNameCounter = 0
}

