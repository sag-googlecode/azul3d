// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package shader

import (
	"log"
	"sync"
)

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
	Vertex Type = iota
	Fragment
)

type Shader struct {
	sync.RWMutex

	name                   string
	nativeIdentity         interface{}
	clearOnLoad            bool
	vertSource, fragSource []byte
	error                  []byte
	inputs                 map[string]interface{}
	changed                []string

	loaded    bool
	notifiers []chan bool
}

func (s *Shader) SetInput(name string, value interface{}) {
	switch value.(type) {
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
		panic("Invalid shader input type!")
	}

	current, ok := s.Input(name)
	if !ok || current != value {
		s.Lock()
		defer s.Unlock()
		s.inputs[name] = value
		s.changed = append(s.changed, name)
	}
}

func (s *Shader) Input(name string) (value interface{}, ok bool) {
	s.RLock()
	defer s.RUnlock()

	value, ok = s.inputs[name]
	return
}

func (s *Shader) Changed() []string {
	s.RLock()
	defer s.RUnlock()

	changed := s.changed
	s.changed = nil
	return changed
}

func (s *Shader) Inputs() map[string]interface{} {
	s.RLock()
	defer s.RUnlock()

	cpy := make(map[string]interface{}, len(s.inputs))
	for k, v := range s.inputs {
		cpy[k] = v
	}
	return cpy
}

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
// Default: true
func (s *Shader) SetClearOnLoad(clearOnLoad bool) {
	s.Lock()
	defer s.Unlock()

	s.clearOnLoad = clearOnLoad
}

// ClearOnLoad tells weather or not the source of this shader will be cleared
// when it is loaded.
//
// Default: true
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
	s.clearOnLoad = true
	s.inputs = make(map[string]interface{})
	return s
}
