// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import "sync"

// NativeObject represents a native graphics object, they are normally only
// created by renderers.
type NativeObject interface{
	// If the GPU supports occlusion queries (see GPUInfo.OcclusionQuery) and
	// OcclusionTest is set to true on the graphics object, then this method
	// will return the number of samples that passed the depth and stencil
	// testing phases the last time the object was drawn. If occlusion queries
	// are not supported then -1 will be returned.
	SampleCount() int
}

// Object represents a single graphics object for rendering, it has a
// transformation matrix which is applied to each vertex of each mesh, it
// has a shader program, meshes, and textures used for rendering the object.
//
// Clients are responsible for utilizing the RWMutex of the object when using
// it or invoking methods.
type Object struct {
	sync.RWMutex

	// The native object of this graphics object. The renderer using this
	// graphics object must assign a value to this field after a call to
	// Draw() has finished before unlocking the object.
	NativeObject

	// Whether or not this object should be occlusion tested. See also the
	// SampleCount() method of NativeObject.
	OcclusionTest bool

	// The render state of this object.
	State

	// The transformation of the object.
	Transform

	// The shader program to be used during rendering the object.
	*Shader

	// A slice of meshes which make up the object. The order in which the
	// meshes appear in this slice also affects the order in which they are
	// sent to the graphics card.
	//
	// This is a slice specifically to allow renderer implementations to
	// optimize the number of draw calls that must occur to render
	// consecutively listed meshes here (this allows for 'hardware' instancing
	// support).
	Meshes []*Mesh

	// A slice of textures which are used to texture the meshes of this object.
	// The order in which the textures appear in this slice is also the order
	// in which they are sent to the graphics card.
	Textures []*Texture
}

// Compare compares this object's state (including shader and textures) against
// the other one and determines if it should sort before the other one for
// state sorting purposes.
func (o *Object) Compare(other *Object) bool {
	if o == other {
		return true
	}

	// Compare shaders.
	if o.Shader != other.Shader {
		return false
	}

	// Compare textures.
	for i, tex := range o.Textures {
		if other.Textures[i] != tex {
			return false
		}
	}

	// Compare state then.
	return o.State.Compare(other.State)
}

// NewObject creates and returns a new object with:
//  o.State == DefaultState
//  o.Transform == DefaultTransform
func NewObject() *Object {
	return &Object{
		State:     DefaultState,
		Transform: DefaultTransform,
	}
}
