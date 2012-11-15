package main

import "code.google.com/p/azul3d/wrappers/glfw"
import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "code.google.com/p/azul3d/wrappers/gl/arb"
import "fmt"

type ParamGL struct {
    param int
    ext, name string
}

type ParamGLFW struct {
    param int
    name string
}


var gl_params = [...]ParamGL {
    ParamGL{gl.RED_BITS, "", "red bits"},
    ParamGL{gl.GREEN_BITS, "", "green bits"},
    ParamGL{gl.BLUE_BITS, "", "blue bits"},
    ParamGL{gl.ALPHA_BITS, "", "alpha bits"},
    ParamGL{gl.DEPTH_BITS, "", "depth bits"},
    ParamGL{gl.STENCIL_BITS, "", "stencil bits"},
    ParamGL{gl.STEREO, "", "stereo"},
    ParamGL{arb.SAMPLES_ARB, "GL_ARB_multisample", "FSAA samples"},
}

var glfw_params = [...]ParamGLFW {
    ParamGLFW{glfw.REFRESH_RATE, "refresh rate"},
    ParamGLFW{glfw.OPENGL_VERSION_MAJOR, "OpenGL major"},
    ParamGLFW{glfw.OPENGL_VERSION_MINOR, "OpenGL minor"},
    ParamGLFW{glfw.OPENGL_FORWARD_COMPAT, "OpenGL forward compatible"},
    ParamGLFW{glfw.OPENGL_DEBUG_CONTEXT, "OpenGL debug context"},
    ParamGLFW{glfw.OPENGL_PROFILE, "OpenGL profile"},
}

func main() {
    var err error

    err = glfw.Init()
    if err != nil {
        panic(err.Error())
    }
    defer glfw.Terminate()

    err = gl.Init()
    if err != nil {
        panic(err.Error())
    }

    glfw.WindowHint(glfw.VISIBLE, 0)

    window, err := glfw.CreateWindow(0, 0, glfw.WINDOWED, "Defaults", nil)
    if err != nil {
        panic(err.Error())
    }

    window.MakeContextCurrent()
    width, height := window.Size()

    fmt.Printf("window size: %dx%d\n", width, height)

    var i int
    for i = 0; i < len(glfw_params); i++ {
        fmt.Printf("%s: %d\n", glfw_params[i].name, window.Param(glfw_params[i].param))
    }

    for i := 0; i < len(gl_params); i++ {
        if len(gl_params[i].ext) != 0 {
            if !glfw.ExtensionSupported(gl_params[i].ext) {
                continue
            }
        }

        var value gl.Int
        gl.GetIntegerv(gl.Enum(gl_params[i].param), &value)

        fmt.Printf("%s: %d\n", gl_params[i].name, int(value))
    }
}

