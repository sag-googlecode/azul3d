// Azul3D, cross platform 3D game engine.
// See included "License.txt".
/*
	Native wrappers for Linux, used to open an OpenGL window
*/

package native

/*
    #cgo LDFLAGS: -lX11 -lGL
    #include <X11/Xutil.h> //Needed for XVisualInfo
    #include <X11/extensions/Xrandr.h> // X11 extension for resizing etc
    #include <X11/keysym.h>
    #include <X11/Xlib.h>
    #include <GL/glx.h>

    GLXFBConfig* FBConfigAtIndex(GLXFBConfig *fbc, int index){
        return &fbc[index];
    }
*/
import "C"
import "unsafe"
import "errors"
import "reflect"
//import "fmt"


var GLX_X_RENDERABLE int = C.GLX_X_RENDERABLE
var GLX_DRAWABLE_TYPE int = C.GLX_DRAWABLE_TYPE
var GLX_RENDER_TYPE int = C.GLX_RENDER_TYPE
var GLX_X_VISUAL_TYPE int = C.GLX_X_VISUAL_TYPE
var GLX_RED_SIZE int = C.GLX_RED_SIZE
var GLX_GREEN_SIZE int = C.GLX_GREEN_SIZE
var GLX_BLUE_SIZE int = C.GLX_BLUE_SIZE
var GLX_ALPHA_SIZE int = C.GLX_ALPHA_SIZE
var GLX_DEPTH_SIZE int = C.GLX_DEPTH_SIZE
var GLX_STENCIL_SIZE int = C.GLX_STENCIL_SIZE
var GLX_DOUBLEBUFFER int = C.GLX_DOUBLEBUFFER
var GLX_SAMPLE_BUFFERS int = C.GLX_SAMPLE_BUFFERS
var GLX_SAMPLES int = C.GLX_SAMPLES
var GLX_WINDOW_BIT int = C.GLX_WINDOW_BIT
var GLX_RGBA_BIT int = C.GLX_RGBA_BIT
var GLX_TRUE_COLOR int = C.GLX_TRUE_COLOR




type Display struct{
	This *C.Display
}

type XVisualInfo struct{
	This *C.XVisualInfo
}

func (vi *XVisualInfo) Free(){
    C.XFree(unsafe.Pointer(vi.This))
}

func XDefaultScreen(display *Display) (int){
    return int(C.XDefaultScreen(display.This))
}

func XOpenDisplay(displayName string) (*Display, error){
    d := Display{}
    d.This = C.XOpenDisplay(C.CString(displayName))
    if d.This == nil{
        return nil, errors.New("Unable to open display")
    }
    return &d, nil
}



type FBConfig struct{
    This *C.GLXFBConfig
}

func (parent *FBConfig) Index(i int) (*FBConfig){
    fbc := FBConfig{}
    fbc.This = C.FBConfigAtIndex(parent.This, C.int(i))
    return &fbc
}

func ChooseFBConfig(dpy *Display, screen int, attrib_list []int) (*FBConfig, int, error) {
    var visualAttributes []C.int

    for i := 0; i < len(attrib_list); i++{
        visualAttributes = append(visualAttributes, C.int(attrib_list[i]))
    }
    visualAttributes = append(visualAttributes, 0) // Should be NULL ended

    header := (*reflect.SliceHeader)(unsafe.Pointer(&visualAttributes)).Data
    visualAttributesPtr := (*C.int)(unsafe.Pointer(header))

    var fbCount C.int
    fbc := FBConfig{}

    fbc.This = C.glXChooseFBConfig(dpy.This, C.int(screen), visualAttributesPtr, &fbCount)
    if fbc.This == nil{
        return nil, 0, errors.New("glXChooseFBConfig() call Failed")
    }

    return &fbc, int(fbCount), nil
}

func GetVisualFromFBConfig(dpy *Display, config *FBConfig)  (*XVisualInfo, error){
    vi := XVisualInfo{}

    vi.This = C.glXGetVisualFromFBConfig(dpy.This, config.This)

	//fmt.Println("visual is", vi.This, config.This)

	if vi.This == nil{
		vi.Free()
		return nil, errors.New("glXGetVisualFromFBConfig() call failed")
	}

    //visualInformation != nil
    return &vi, nil
}

func QueryVersion(dpy *Display) (int, int, error){
    var major, minor C.int

    if C.glXQueryVersion(dpy.This, &major, &minor) == 0{
        return -1, -1, errors.New("Unable to query version")
    }

    return int(major), int(minor), nil
}

func GetFBConfigAttrib(dpy *Display, config *FBConfig, bits int) (int){
	var value C.int

	C.glXGetFBConfigAttrib(dpy.This, (*[0]byte)(unsafe.Pointer(config.This)), C.int(bits), &value)

	return int(value)
}



/*

//XVisualInfo * glXGetVisualFromFBConfig(Display *  dpy,  GLXFBConfig  config);
//d.display, C.FBConfigAtIndex(fbc, i)
*/
