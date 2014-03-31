// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"sync"
)

// Object represents a single graphics object for rendering, it has a
// transformation matrix which is applied to each vertex of each mesh, it
// has a shader program, meshes, and textures used for rendering the object.
//
// Clients are responsible for utilizing the RWMutex of the object when using
// it or invoking methods.
type Object struct {
	sync.RWMutex

	// The native object of this graphics object. The renderer using this
	// graphics object may assign anything to this interface. Typically clients
	// will not use this field at all.
	Native interface{}

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

// NewObject creates and returns a new object with:
//  o.State == DefaultState
//  o.Transform == DefaultTransform
func NewObject() *Object {
	return &Object{
		State:     DefaultState,
		Transform: DefaultTransform,
	}
}
