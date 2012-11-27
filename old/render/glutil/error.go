package glutil

import gl "code.google.com/p/azul3d/wrappers/gl/gl21"
import "errors"
import "bytes"

func GetError() error {
    err := gl.GetError()
    switch(err) {
        case gl.NO_ERROR:
            return nil
        case gl.INVALID_ENUM:
            return errors.New("GL_INVALID_ENUM: An unacceptable value is specified for an enumerated argument.")
        case gl.INVALID_VALUE:
            return errors.New("GL_INVALID_VALUE: A numeric argument is out of range.")
        case gl.INVALID_OPERATION:
            return errors.New("GL_INVALID_OPERATION: The specified operation is not allowed in the current state.")
        // 2.1 has no INVALID_FRAMEBUFFER_OPERATION
        //
        //case gl.INVALID_FRAMEBUFFER_OPERATION:
        //    return errors.New("GL_INVALID_FRAMEBUFFER_OPERATION: The framebuffer object is not complete.")
        case gl.OUT_OF_MEMORY:
            return errors.New("GL_OUT_OF_MEMORY: There is not enough memory left to execute the command.")
        case gl.STACK_UNDERFLOW:
            return errors.New("GL_STACK_UNDERFLOW: An attempt has been made to perform an operation that would cause an internal stack to underflow.")
        case gl.STACK_OVERFLOW:
            return errors.New("GL_STACK_OVERFLOW: An attempt has been made to perform an operation that would cause an internal stack to overflow.")
    }
    return nil
}


func GetErrors() string {
    b := bytes.Buffer{}
    err := GetError()
    for err != nil {
        if err == nil {
            break
        }
        b.WriteString(err.Error())
        b.WriteString("\n")
        err = GetError()
    }
    return b.String()
}

