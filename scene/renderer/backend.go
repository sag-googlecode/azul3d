// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package renderer

const (
	// OpenGL backend (automatically chooses best version to use)
	OpenGL BackendType = iota

	// OpenGL Fixed Function backend (AKA OpenGL 1.0 to 3.0)
	OpenGLFixedFunction
)

// BackendType represents an single renderer backend type which can be used.
// For an backend type to be considered valid, it must be one of the predefined
// constants.
type BackendType uint

// String returns an string representation of this backend type.
func (b BackendType) String() string {
	switch b {
	case OpenGL:
		return "OpenGL"

	case OpenGLFixedFunction:
		return "OpenGLFixedFunction"
	}
	return "Invalid"
}

// Name returns an name string for this backend type.
func (b BackendType) Name() string {
	switch b {
	case OpenGL:
		return "OpenGL"

	case OpenGLFixedFunction:
		return "OpenGL 1.0 to 3.0 Renderer"
	}
	return "Invalid"
}

// Valid returns weather this backend type is an valid one (I.e. one of the predefined constants).
func (b BackendType) Valid() bool {
	switch b {
	case OpenGL:
		return true

	case OpenGLFixedFunction:
		return true
	}
	return false
}
