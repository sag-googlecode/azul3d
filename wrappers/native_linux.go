// Azul3D, cross platform 3D game engine.
// See included "License.txt".
/*
	Native wrappers for Linux, used to open an OpenGL window
*/

package native

/*
    #cgo LDFLAGS: -lGL -lX11
    #include <GL/glx.h>
    #include <X11/Xutil.h> //Needed for XVisualInfo
    #include <X11/Xlib.h>

    GLXFBConfig* FBConfigAtIndex(GLXFBConfig *fbc, int index){
        return &fbc[index];
    }
*/








/*
type Display struct{
    RealDisplay *C.Display
}

type XVisualInfo struct{
    RealXVisualInfo *C.XVisualInfo
}

func (vi *XVisualInfo) Free(){
    C.XFree(unsafe.Pointer(vi.RealXVisualInfo))
}


func XDefaultScreen(display *Display) (int){
    return int(C.XDefaultScreen(display.RealDisplay))
}

func XOpenDisplay(displayName string) (*Display, error){
    d := Display{}
    d.RealDisplay = C.XOpenDisplay(C.CString(displayName))
    if d.RealDisplay == nil{
        return nil, errors.New("Unable to open display")
    }
    return &d, nil
}




type FBConfig struct{
    RealFBConfig *C.GLXFBConfig
}

func (parent *FBConfig) Index(i int) (*FBConfig){
    fbc := FBConfig{}
    fbc.RealFBConfig = C.FBConfigAtIndex(parent.RealFBConfig, C.int(i))
    return &fbc
}


func ChooseFBConfig(dpy *xlib.Display, screen int, attrib_list []int) (*FBConfig, int, error) {
    var visualAttributes []C.int

    for i := 0; i < len(attrib_list); i++{
        visualAttributes = append(visualAttributes, C.int(attrib_list[i]))
    }
    visualAttributes = append(visualAttributes, 0) // Should be NULL ended

    header := (*reflect.SliceHeader)(unsafe.Pointer(&visualAttributes)).Data
    visualAttributesPtr := (*C.int)(unsafe.Pointer(header))

    var fbCount C.int
    fbc := FBConfig{}

    fbc.RealFBConfig = C.glXChooseFBConfig(dpy.RealDisplay, C.int(screen), visualAttributesPtr, &fbCount)
    if fbc.RealFBConfig == nil{
        return nil, 0, errors.New("glXChooseFBConfig() call Failed")
    }

    return &fbc, int(fbCount), nil
}


func GetVisualFromFBConfig(dpy *xlib.Display, config *FBConfig)  (*xlib.XVisualInfo, error){
    vi := xlib.XVisualInfo{}

    vi.RealXVisualInfo = C.glXGetVisualFromFBConfig(dpy.RealDisplay, (*[0]byte)(unsafe.Pointer(config.RealFBConfig)))

    //visualInformation != nil
    return &vi, nil
}


//XVisualInfo * glXGetVisualFromFBConfig(Display *  dpy,  GLXFBConfig  config);
//d.display, C.FBConfigAtIndex(fbc, i)

func QueryVersion(dpy *xlib.Display) (int, int, error){
    var major, minor C.int

    if C.glXQueryVersion(dpy.RealDisplay, &major, &minor) == 0{
        return -1, -1, errors.New("Unable to query version")
    }

    return int(major), int(minor), nil
}
*/
