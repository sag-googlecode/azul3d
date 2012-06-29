// Azul3D, cross platform 3D game engine.
// See included "License.txt".
/*
	Display for linux, uses native/native_linux to open OpenGL windows
*/

package display

type Display struct{
}

func NewDisplay() (error, *Display){
	d := Display{}

    d.display, e = xlib.XOpenDisplay("")


	return nil, &d
}


