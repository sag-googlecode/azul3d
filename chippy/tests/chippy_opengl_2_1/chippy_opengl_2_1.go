// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

// Test application - Opens multiple windows, uses OpenGL 2.1 rendering in each of them
package main

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/native/opengl/2.1"
	"log"
	"os"
	"runtime"
)

func main() {
	log.SetFlags(0)

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	err := chippy.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer chippy.Destroy()

	window := chippy.NewWindow()

	// Actually open the windows
	screen := chippy.DefaultScreen()
	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what they currently has property-wise
	log.Println(window)

	// All OpenGL related calls must occur in the same OS thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Choose an buffer format, these include things like double buffering, bytes per pixel, number of depth bits, etc.
	configs := window.GLConfigs()

	// See documentation for this function and vars to see how it determines the 'best' format
	bestConfig := chippy.GLChooseConfig(configs, chippy.GLWorstConfig, chippy.GLBestConfig)
	window.GLSetConfig(bestConfig)

	// Print out all the formats, and which one we determined to be the 'best'.
	log.Println("\nChosen configuration:")
	log.Println(bestConfig)

	// Create an OpenGL context with the OpenGL version we wish
	context, err := window.GLCreateContext(2, 1)
	if err != nil {
		log.Fatal(err)
	}

	// Make the context current in this OS thread
	window.GLMakeCurrent(context)

	// We get an ContextAccess, which references (and caches) per-opengl context functions which
	// will be gotten through fooGetProcAddress, barGetProcAddress, etc.
	//
	// An good idea is that for every opengl context, you use an single one of these. (but you can
	// use more than one without an large performance penalty, of course. Just don't use one for
	// every opengl call that you make.)
	gl := opengl.New()
	if gl == nil {
		log.Fatal("You have no support for OpenGL 1.1!")
	}
	log.Println(gl.GetError())

	//pa := chippy.GLGetProcAddress("glAddSwapHintRectWIN")
	//log.Println("glAddSwapHintRectWIN:", pa)

	//pa = chippy.GLGetProcAddress("wglGetProcAddress")
	//log.Println("wglGetProcAddress", pa)

	//pa = chippy.GLGetProcAddress("glGetString")
	//log.Println("glGetString:", pa)

	/*
		glGetString := *(*func(name uint32)(*uint8))(unsafe.Pointer(&pa))
		r := glGetString(0x1F01)
		log.Println("glGetString()", r)
		log.Println(C.GoString((*C.char)(unsafe.Pointer(r))))
	*/

	/*
		s := []byte{}
		i := 0
		for {
			e := uintptr(unsafe.Pointer(&r)) + uintptr(i)
			character := *((*byte)(unsafe.Pointer(&e)))
			if character == 0 {
				break
			}
			s = append(s, character)
			i++
		}
		log.Println(string(s))
	*/

	/*
		// Begin our rendering loop
		for !window.IsDestroyed() {
			runtime.Gosched()

			// Our OpenGL rendering code goes here
			window.GLSwapBuffers()
		}
	*/
}
