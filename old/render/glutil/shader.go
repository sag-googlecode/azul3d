package glutil

import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "unsafe"
import "bytes"
import "fmt"
import "os"

var VERTEX_SHADER uint8   = 1
var FRAGMENT_SHADER uint8 = 2

type Shader struct {
    Shader gl.Uint
    ShaderSource string
    ShaderType uint8
    hasError bool
    errorString string
}

func NewShader() *Shader {
    return &Shader{}
}

func (s *Shader) LoadString(shaderSource string, shaderType uint8) {
    s.ShaderSource = shaderSource
    s.ShaderType = shaderType

    if shaderType == VERTEX_SHADER {
        s.Shader = gl.CreateShader(gl.VERTEX_SHADER)
    } else if shaderType == FRAGMENT_SHADER {
        s.Shader = gl.CreateShader(gl.FRAGMENT_SHADER)
    } else {
        s.hasError = true
        s.errorString = "Invalid shaderType specified"
        return // Invalid shader type
    }

    if gl.IsShader(s.Shader) == gl.FALSE {
        s.hasError = true
        s.errorString = "Shader returned by glCreateShader() is not a valid shader, according to glIsShader()\nPerhaps you are missing a valid GL context?"
        return
    }

    src := gl.GLString(shaderSource)
    gl.ShaderSource(s.Shader, 1, (**gl.Char)(unsafe.Pointer(&src)), nil)
    gl.GLStringFree(src)
    
    //sa := gl.GLStringArray(shaderSource)
    //gl.ShaderSource(s.Shader, 1, (**gl.Char)(unsafe.Pointer(&sa)), nil)
    //gl.GLStringArrayFree(sa)

    gl.CompileShader(s.Shader)

    var sourceLength gl.Int
    gl.GetShaderiv(s.Shader,gl.SHADER_SOURCE_LENGTH, &sourceLength)
    if int(sourceLength) != len(s.ShaderSource)+1 {
        fmt.Println("GL_SHADER_SOURCE_LENGTH:", sourceLength)
        fmt.Println("Actual length:", len(s.ShaderSource))
        panic("Copied the shader to GL improperly, they are of different lengths!")
    }

    var shaderCompiled gl.Int
    gl.GetShaderiv(s.Shader, gl.COMPILE_STATUS, &shaderCompiled)
    if shaderCompiled == gl.FALSE {
        s.hasError = true
        defer s.Delete() // Delete the shader, but only later, we need it for glGetShaderiv()
    } else {
        s.hasError = false
    }

    var logLength gl.Int
    var charsWritten gl.Sizei

    gl.GetShaderiv(s.Shader, gl.INFO_LOG_LENGTH, &logLength)

    if logLength > 0 {
        infoLog := gl.GLStringAlloc(gl.Sizei(logLength))
        defer gl.GLStringFree(infoLog)

        gl.GetShaderInfoLog(s.Shader, gl.Sizei(logLength), &charsWritten, infoLog)
        s.errorString = gl.GoStringN(infoLog, charsWritten)
    }
}

func (s *Shader) LoadFile(filePath string, shaderType uint8) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }

    data := bytes.Buffer{}
    _, err = data.ReadFrom(file)
    if err != nil {
        return err
    }
    s.LoadString(data.String(), shaderType)
    return nil
}

func (s *Shader) Delete() {
    gl.DeleteShader(s.Shader)
}

func (s *Shader) HasError() bool {
    return s.hasError
}

func (s *Shader) ErrorString() string {
    return s.errorString
}

