package glutil

import gl "code.google.com/p/azul3d/wrappers/gl/gl21"

type Program struct {
    Program gl.Uint
    shaders []*Shader
    hasError bool
    errorString string
}

func NewProgram() *Program {
    p := Program{}
    p.Program = gl.CreateProgram()
    if gl.IsProgram(p.Program) == gl.FALSE {
        p.hasError = true
        p.errorString = "Program returned by glCreateProgram() is not a valid program, according to glIsProgram()\nPerhaps you are missing a valid GL context?"
    }
    return &p
}

func (p *Program) AddShader(s *Shader) {
    p.shaders = append(p.shaders, s)
}

func (p *Program) Shaders() []*Shader {
    shaders := []*Shader{}
    copy(p.shaders, shaders)
    return shaders
}

func (p *Program) Compile() {
    // Firstly we verify that all the shaders are compiled without errors
    for i := 0; i < len(p.shaders); i++ {
        if p.shaders[i].HasError() {
            p.hasError = true
            p.errorString = p.shaders[i].ErrorString() // Suppose this is sufficient enough here
            return
        }
    }

    // Attach all shaders
    for i := 0; i < len(p.shaders); i++ {
        gl.AttachShader(p.Program, p.shaders[i].Shader)
    }

    gl.LinkProgram(p.Program)

    var programCompiled gl.Int
    gl.GetProgramiv(p.Program, gl.LINK_STATUS, &programCompiled)

    if programCompiled == gl.FALSE {
        p.hasError = true
        defer p.Delete() // We still need it for glGetProgramiv
    } else {
        p.hasError = false
    }

    var infoLogLength gl.Int
    var charsWritten gl.Sizei

    gl.GetProgramiv(p.Program, gl.INFO_LOG_LENGTH, &infoLogLength)


    if infoLogLength > 0 {
        infoLog := gl.GLStringAlloc(gl.Sizei(infoLogLength))
        gl.GetProgramInfoLog(p.Program, gl.Sizei(infoLogLength), &charsWritten, infoLog)
        p.errorString = gl.GoStringN(infoLog, charsWritten)
        gl.GLStringFree(infoLog)
    }
}

func (p *Program) Use() {
    if p.HasError() == false {
        gl.UseProgram(p.Program)
    }
}

func (p *Program) Delete() {
    gl.DeleteProgram(p.Program)
}

func (p *Program) HasError() bool {
    return p.hasError
}

func (p *Program) ErrorString() string {
    return p.errorString
}

