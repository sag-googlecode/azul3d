// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package shader

import (
	"log"
	"sync"
)

// Type represents the type of which the shader source is.
type Type uint8

// Valid tells if the shader type is a valid one (I.e. one of the predefined
// shader type constants).
func (t Type) Valid() bool {
	switch t {
	case Vertex:
		return true
	case Fragment:
		return true
	}
	return false
}

const (
	// Vertex shader source type
	Vertex Type = iota

	// Fragment shader source type
	Fragment
)

// Shader represents a single shader object.
type Shader struct {
	sync.RWMutex

	name                   string
	nativeIdentity         interface{}
	clearOnLoad            bool
	vertSource, fragSource []byte
	error                  []byte

	loaded    bool
	notifiers []chan bool
}

// Copy returns a new 1:1 copy of this shader. The underlying source is not
// copied but is instead referenced.
func (s *Shader) Copy() *Shader {
	cpy := new(Shader)
	cpy.name = s.name
	cpy.clearOnLoad = s.clearOnLoad
	cpy.vertSource = s.vertSource
	cpy.fragSource = s.fragSource
	cpy.error = s.error
	return cpy
}

// IsValidType tells if the specified value is of a valid type to be a shader
// input value.
func IsValidType(value interface{}) bool {
	switch value.(type) {
	case bool:
		break
	case float32:
		break
	case int32:
		break
	case uint32:
		break
	case Vec2:
		break
	case Vec3:
		break
	case Vec4:
		break
	case Vec2i:
		break
	case Vec3i:
		break
	case Vec4i:
		break
	case Vec2ui:
		break
	case Vec3ui:
		break
	case Vec4ui:
		break
	case Mat2:
		break
	case Mat3:
		break
	case Mat4:
		break
	case Mat2x3:
		break
	case Mat3x2:
		break
	case Mat2x4:
		break
	case Mat4x2:
		break
	case Mat3x4:
		break
	case Mat4x3:
		break

	case []float32:
		break
	case []int32:
		break
	case []uint32:
		break
	case []Vec2:
		break
	case []Vec3:
		break
	case []Vec4:
		break
	case []Vec2i:
		break
	case []Vec3i:
		break
	case []Vec4i:
		break
	case []Vec2ui:
		break
	case []Vec3ui:
		break
	case []Vec4ui:
		break
	case []Mat2:
		break
	case []Mat3:
		break
	case []Mat4:
		break
	case []Mat2x3:
		break
	case []Mat3x2:
		break
	case []Mat2x4:
		break
	case []Mat4x2:
		break
	case []Mat3x4:
		break
	case []Mat4x3:
		break

	default:
		return false
	}
	return true
}

// Name returns the name of this shader, as it was passed in at creation.
func (s *Shader) Name() string {
	s.RLock()
	defer s.RUnlock()

	return s.name
}

// MarkLoaded marks this shader as loaded. Only the renderer should call this,
// and as such you should (normally) never call this function.
func (s *Shader) MarkLoaded() {
	s.Lock()

	s.loaded = true
	notifiers := s.notifiers
	s.notifiers = nil

	if s.clearOnLoad {
		s.vertSource = nil
		s.fragSource = nil
	}

	s.Unlock()

	errLog := s.Error()
	if errLog != nil {
		log.Println(string(errLog))
	}

	for _, notify := range notifiers {
		notify <- true
	}
}

// Loaded tells if this shader is currently loaded or not.
func (s *Shader) Loaded() bool {
	s.RLock()
	defer s.RUnlock()

	return s.loaded
}

// LoadNotify returns an channel on which true is sent once this shader is
// marked as loaded (normally by the renderer).
func (s *Shader) LoadNotify() chan bool {
	s.Lock()
	defer s.Unlock()

	notify := make(chan bool, 1)
	if s.loaded {
		notify <- true
		return notify
	}

	if s.notifiers == nil {
		s.notifiers = make([]chan bool, 0)
	}
	s.notifiers = append(s.notifiers, notify)
	return notify
}

// SetNativeIdentity specifies the native identity of this shader.
//
// This should mostly not be used (except in very rare, advanced cases).
func (s *Shader) SetNativeIdentity(identity interface{}) {
	s.Lock()
	defer s.Unlock()

	s.nativeIdentity = identity
}

// NativeIdentity returns the native identity of this shader.
//
// This should mostly not be used (except in very rare, advanced cases).
func (s *Shader) NativeIdentity() interface{} {
	s.RLock()
	defer s.RUnlock()

	return s.nativeIdentity
}

// SetClearOnLoad specifies weather or not to clear the source of this shader
// when it is loaded.
//
// Default: false
func (s *Shader) SetClearOnLoad(clearOnLoad bool) {
	s.Lock()
	defer s.Unlock()

	s.clearOnLoad = clearOnLoad
}

// ClearOnLoad tells weather or not the source of this shader will be cleared
// when it is loaded.
//
// Default: false
func (s *Shader) ClearOnLoad() bool {
	s.RLock()
	defer s.RUnlock()

	return s.clearOnLoad
}

// SetSource specifies the source of the shader, and it's associated source type.
//
// The shader source type must be valid or else a panic will occur.
func (s *Shader) SetSource(source []byte, t Type) {
	if !t.Valid() {
		panic("SetSource(): Invalid shader source type!")
	}

	s.Lock()
	defer s.Unlock()

	switch t {
	case Vertex:
		s.vertSource = source
	case Fragment:
		s.fragSource = source
	}
}

// Source returns the source of the shader for the given shader source type.
//
// The shader source type must be valid or else a panic will occur.
func (s *Shader) Source(t Type) []byte {
	if !t.Valid() {
		panic("SetSource(): Invalid shader source type!")
	}

	s.RLock()
	defer s.RUnlock()

	switch t {
	case Vertex:
		return s.vertSource
	case Fragment:
		return s.fragSource
	}

	// Never happens
	return nil
}

// SetError sets the compiler error log of this shader object.
func (s *Shader) SetError(error []byte) {
	s.RLock()
	defer s.RUnlock()

	s.error = error
}

// Error returns the (read-only) compiler error log of this shader.
func (s *Shader) Error() []byte {
	s.RLock()
	defer s.RUnlock()

	return s.error
}

// New returns a new shader object.
func New(name string) *Shader {
	s := new(Shader)
	s.name = name
	return s
}
