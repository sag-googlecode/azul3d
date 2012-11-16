package glfw

/*
#cgo CFLAGS: -DGLFW_INCLUDE_GLU
#cgo LDFLAGS: -lglfw
#include <GL/glfw3.h>
#include <stdlib.h>
*/
import "C"

import "unsafe"

// The video mode structure used by glfwGetVideoModes
type Vidmode struct {
	Width, Height, RedBits, BlueBits, GreenBits int
}

func (this *Vidmode) Equals(v *Vidmode) bool {
	return this.Width == v.Width && this.Height == v.Height && this.RedBits == v.RedBits && this.BlueBits == v.BlueBits && this.GreenBits == v.GreenBits
}

// Gamma ramp
type Ramp struct {
	Red   [256]uint16
	Green [256]uint16
	Blue  [256]uint16
}

func gammarampToGLFW(r *Ramp) *C.GLFWgammaramp {
	var gr *C.GLFWgammaramp
	gr = (*C.GLFWgammaramp)(C.malloc(C.size_t(unsafe.Sizeof(gr))))
	// Copy over red
	for i := 0; i < len(r.Red); i++ {
		gr.red[i] = C.ushort(r.Red[i])
	}
	// Copy over green
	for i := 0; i < len(r.Green); i++ {
		gr.green[i] = C.ushort(r.Green[i])
	}
	// Copy over blue
	for i := 0; i < len(r.Blue); i++ {
		gr.blue[i] = C.ushort(r.Blue[i])
	}
	return gr
}

func glfwToGammaramp(gr *C.GLFWgammaramp) *Ramp {
	r := Ramp{}
	// Copy over red
	for i := 0; i < len(r.Red); i++ {
		r.Red[i] = uint16(gr.red[i])
	}
	// Copy over green
	for i := 0; i < len(r.Green); i++ {
		r.Green[i] = uint16(gr.green[i])
	}
	// Copy over blue
	for i := 0; i < len(r.Blue); i++ {
		r.Blue[i] = uint16(gr.blue[i])
	}
	return &r
}

//    unsigned short red[256] #GLFW_GAMMA_RAMP_SIZE
//    unsigned short green[256] #GLFW_GAMMA_RAMP_SIZE
//    unsigned short blue[256] #GLFW_GAMMA_RAMP_SIZE

