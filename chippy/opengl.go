// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

package chippy

type GLContext interface {
	// Like wglShareLists(thisContext, c)
	Share(c GLContext)
}

type GLRenderable interface {
	// GLConfigs returns all possible OpenGL configurations, these are valid configurations that
	// may be used in an call to GLSetConfig.
	GLConfigs() []*GLConfig

	// GLSetConfig sets the OpenGL configuration, this action may only be performed once, if you
	// attempt to call this function twice, it will become no-op.
	//
	// The config parameter must be an *GLConfig that originally came from the GLConfigs() function
	// due to the fact that it must be initialized internally.
	GLSetConfig(config *GLConfig)

	// GLConfig returns the *GLConfig or nil, as it was previously set via GLSetConfig
	GLConfig() *GLConfig

	// GLCreateContext creates an OpenGL context for the specified OpenGL version, or returns
	// an error in the event that we cannot create
	GLCreateContext(glVersionMinor, glVersionMajor uint) (GLContext, error)

	// GLDestroyContext destroys the specified OpenGL context.
	GLDestroyContext(c GLContext)

	// GLMakeCurrent makes the specified context the current rendering context for this Renderable
	// in the current OS thread.
	GLMakeCurrent(c GLContext)

	// GLSwapBuffers swaps the front and back buffers of this Renderable, if the GLConfig set
	// previously via GLSetConfig is not DoubleBuffered, then this function is no-op and is safe to
	// call.
	GLSwapBuffers()
}
