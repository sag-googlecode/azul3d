// Azul3D, cross platform 3D game engine.
// See included "License.txt".
/*
	Display for linux, uses native/native_linux to open OpenGL windows
*/

package display

import "code.google.com/p/azul3d/native"
import "errors"
import "fmt"

type Display struct{
	display *native.Display
}

func NewDisplay() (*Display, error){
	var e error
	var glxMajor, glxMinor int

	d := Display{}

    d.display, e = native.XOpenDisplay("")
	if e != nil{
		return nil, errors.New(e.Error())
	}

	glxMajor, glxMinor, e = native.QueryVersion(d.display)
	if e != nil{
		return nil, errors.New(e.Error())
	}


	// FBConfigs where added in GLX version 1.3
	if (glxMajor == 1) && (glxMinor < 3) || (glxMajor < 1){
        return nil, errors.New("Display: GLX version must be 1.3 or greater!")
    }

    visualAttributes := []int{
        native.GLX_X_RENDERABLE,   1, // bool
        native.GLX_DRAWABLE_TYPE,  native.GLX_WINDOW_BIT,
        native.GLX_RENDER_TYPE,    native.GLX_RGBA_BIT,
        native.GLX_X_VISUAL_TYPE,  native.GLX_TRUE_COLOR,
        native.GLX_RED_SIZE,       8,
        native.GLX_GREEN_SIZE,     8,
        native.GLX_BLUE_SIZE,      8,
        native.GLX_ALPHA_SIZE,     8,
        native.GLX_DEPTH_SIZE,     24,
        native.GLX_STENCIL_SIZE,   8,
        native.GLX_DOUBLEBUFFER,   1, // bool
        //native.GLX_SAMPLE_BUFFERS, 1,
        //native.GLX_SAMPLES,        4,
    }

    fmt.Println("Display: Getting matching frame buffer configurations")

    var fbc *native.FBConfig
    var fbCount int
    fbc, fbCount, e = native.ChooseFBConfig(d.display, native.XDefaultScreen(d.display), visualAttributes)
    if e != nil{
        return nil, errors.New(e.Error())
    }

    fmt.Println("Display:", fbCount, "matching configurations")


    // Pick the visual with the most samples per pixel
    fmt.Println("Getting visual information via glXGetVisualFromFBConfig()")

    var bestFbcIndex int = -1
    var worstFbc int = -1
    var bestNumSamples int = -1
    var worstNumSamples int = 999

	//fmt.Println("FBC is", fbc.This, fbc.Index(1000).This)

    for i := 0; i < fbCount; i++{
		var visualInformation *native.XVisualInfo
        visualInformation, e = native.GetVisualFromFBConfig(d.display, fbc.Index(i))
		fmt.Println(visualInformation)
        if e == nil{
			native.GetFBConfigAttrib(d.display, fbc.Index(i), native.GLX_SAMPLE_BUFFERS)
		}

		/*
        if e == nil{
			var sampBuf, samples int
			sampBuf = native.GetFBConfigAttrib(d.display, fbc.Index(i), native.GLX_SAMPLE_BUFFERS)

            samples = native.GetFBConfigAttrib(d.display, fbc.Index(i), native.GLX_SAMPLES)

            fmt.Printf("Display: (%d) visual 0x%2x, SAMPLE_BUFFERS = %d, SAMPLES = %d\n", i+1, visualInformation, sampBuf, samples)

			visualInformation.Free()
        }
		*/
    }




    fmt.Println(fbc, bestFbcIndex, worstFbc, bestNumSamples, worstNumSamples)



    /*
    var i int
    for i = 0; i < fbCount; i++{

        if visualInformation != nil{
            var sampBuf, samples int

            C.glXGetFBConfigAttrib(d.display, C.FBConfigAtIndex(fbc, i), C.GLX_SAMPLE_BUFFERS, (*C.int)(unsafe.Pointer(&sampBuf)))

            C.glXGetFBConfigAttrib(d.display, C.FBConfigAtIndex(fbc, i), C.GLX_SAMPLES, (*C.int)(unsafe.Pointer(&samples)))

            fmt.Printf("Display: (%d) visual 0x%2x, SAMPLE_BUFFERS = %d, SAMPLES = %d\n", i+1, visualInformation.visualid, sampBuf, samples)





            if bestFbcIndex < 0 || sampBuf != 0 && samples > bestNumSamples{
                bestFbcIndex = int(i)
                bestNumSamples = samples
            }
            if worstFbc < 0 || sampBuf == 0 || samples < worstNumSamples{
                worstFbc = int(i)
                worstNumSamples = samples
            }
        }

        C.XFree(unsafe.Pointer(visualInformation))
    }
    */

	return &d, nil
}


