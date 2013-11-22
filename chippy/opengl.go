// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

package chippy

import (
	"fmt"
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
			logger().Printf("OpenGL: *** Driver reported version parsing failed for %q ***\n", ver)
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

type GLContextFlags uint8

const (
	GLDebug GLContextFlags = iota
	GLForwardCompatible
	GLCoreProfile
	GLCompatibilityProfile
)

// VSyncMode represents an single vertical reresh rate sync mode.
type VSyncMode uint8

const (
	VerticalSync VSyncMode = iota
	NoVerticalSync
	AdaptiveVerticalSync
)

// Valid tells if the vertical sync mode is one of the predefined constants
// defined in this package or not.
func (mode VSyncMode) Valid() bool {
	switch mode {
	case VerticalSync:
		return true

	case NoVerticalSync:
		return true

	case AdaptiveVerticalSync:
		return true
	}
	return false
}

// String returns a string representation of this vertical sync mode.
func (mode VSyncMode) String() string {
	switch mode {
	case VerticalSync:
		return "VerticalSync"

	case NoVerticalSync:
		return "NoVerticalSync"

	case AdaptiveVerticalSync:
		return "AdaptiveVerticalSync"
	}
	return fmt.Sprintf("VSyncMode(%d)", mode)
}

// GLContext represents an OpenGL contect; although it represents any value it
// represents an important idea of what it's data actually is.
type GLContext interface {
}

type GLRenderable interface {
	// GLConfigs returns all possible OpenGL configurations, these are valid
	// configurations that may be used in an call to GLSetConfig.
	GLConfigs() []*GLConfig

	// GLSetConfig sets the OpenGL framebuffer configuration, unlike other
	// window management libraries, this action may be performed multiple
	// times.
	//
	// The config parameter must be an *GLConfig that originally came from the
	// GLConfigs() function mainly do to the fact that it must be initialized
	// internally.
	GLSetConfig(config *GLConfig)

	// GLConfig returns the currently in use *GLConfig or nil, as it was
	// previously set via an call to GLSetConfig()
	GLConfig() *GLConfig

	// GLCreateContext creates an OpenGL context for the specified OpenGL
	// version, or returns an error in the event that we cannot create an
	// context for that version.
	//
	// The flags parameter may be any combination of the predifined flags, as
	// follows:
	//
	// GLDebug, you will receive an OpenGL debug context. *
	//
	// GLForwardCompatible, you will receive an OpenGL forward compatible
	// context. *
	//
	// GLCoreProfile, you will receive an OpenGL core context.
	//
	// GLCompatibilityProfile, you will receive an OpenGL compatibility
	// context.
	//
	// Only one of GLCoreProfile or GLCompatibilityProfile should be present.
	//
	// GLCompatabilityProfile will be used if neither GLCoreProfile or
	// GLCompatibilityProfile are present, or if both are present.
	//
	// * = It is not advised to use this flag in production.
	//
	// You must call GLSetConfig() before calling this function.
	GLCreateContext(major, minor uint, flags GLContextFlags, share GLContext) (GLContext, error)

	// GLDestroyContext destroys the specified OpenGL context.
	//
	// The context to destroy must not be active in any thread, period.
	GLDestroyContext(c GLContext)

	// GLMakeCurrent makes the specified context the current, active OpenGL
	// context in the current operating system thread.
	//
	// To make the OpenGL context inactive, you may call this function using
	// nil as the context, which will release the context.
	//
	// This function may be called from any thread, but an OpenGL context may
	// only be active inside one thread at an time.
	GLMakeCurrent(c GLContext)

	// GLSwapBuffers swaps the front and back buffers of this Renderable.
	//
	// This function may only be called in the presence of an active OpenGL
	// context.
	//
	// If the GLConfig set previously via GLSetConfig() is not DoubleBuffered,
	// then this function is no-op.
	GLSwapBuffers()

	// GLSetVerticalSync sets the vertical refresh rate sync mode (vsync).
	//
	// This function should only be called in the presence of an active OpenGL
	// context or else the call may fail due to drivers or platforms that
	// require an active context (e.g. Mesa).
	GLSetVerticalSync(mode VSyncMode)
}
