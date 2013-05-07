// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

package chippy

import (
	"strconv"
	"strings"
)

// Do not use for multiple extensions as it splits the string and searches it slowly..
//
// I.e. do not expose to end users..
func extSupported(str, ext string) bool {
	for _, s := range strings.Split(str, " ") {
		if s == ext {
			return true
		}
	}
	return false
}

func versionSupported(ver string, wantedMajor, wantedMinor int) bool {
	if len(ver) > 0 {
		var (
			major, minor int
		)

		versions := strings.Split(ver, ".")
		versions = versions[0:2]

		if len(versions) == 2 {
			major, _ = strconv.Atoi(versions[0])
			minor, _ = strconv.Atoi(versions[1])
		} else {
			logger.Printf("OpenGL: *** Driver reported version parsing failed for %q ***\n", ver)
			return false
		}

		if major > wantedMajor {
			return true
		}
		if major == wantedMajor && minor >= wantedMinor {
			return true
		}
	}
	return false
}

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
	// an error in the event that we cannot create an context for that version.
	GLCreateContext(major, minor uint) (GLContext, error)

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
