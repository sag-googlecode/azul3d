// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

import (
	"code.google.com/p/azul3d/math"
	"sync"
)

// Type represents one of the texture types defined in this package.
//
// An texture type is only valid if IsValid() would return true.
type Type interface {
	AnisotropicDegree() uint
	AutoLoad() bool
	BorderColor() (r, g, b, a math.Real)
	BorderColorFloat32() []float32
	BorderColorVec4() *math.Vec4
	ClearOnLoad() bool
	Compressed() bool
	LoadNotify() chan bool
	Loaded() bool
	Loading() bool
	MagFilter() Filter
	MarkLoaded()
	MarkLoading()
	MinFilter() Filter
	NativeIdentity() interface{}
	SetAnisotropicDegree(degree uint)
	SetAutoLoad(autoLoad bool)
	SetBorderColor(r, g, b, a math.Real)
	SetBorderColorVec4(color *math.Vec4)
	SetClearOnLoad(clearOnLoad bool)
	SetCompressed(compressed bool)
	SetNativeIdentity(identity interface{})
	SetMagFilter(filter Filter)
	SetMinFilter(filter Filter)
	SetWrapModeU(u WrapMode)
	SetWrapModeV(v WrapMode)
	SetWrapModeW(w WrapMode)
	WrapModeU() WrapMode
	WrapModeV() WrapMode
	WrapModeW() WrapMode
}

// IsValid tells if this texture type interface value is one of the following:
//
//  *Texture2D
//
func IsValid(t Type) bool {
	switch t.(type) {
	case *Texture2D:
		return true
	}
	return false
}

// Texture is a base texture type which fills the Type interface. All texture
// types will use this type as an embedded struct.
type Texture struct {
	sync.RWMutex

	Source interface{}

	nativeIdentity                  interface{}
	clearOnLoad, autoLoad           bool
	anisotropicDegree               uint
	wrapModeU, wrapModeV, wrapModeW WrapMode
	borderColor                     *math.Vec4
	minFilter, magFilter            Filter
	compressed                      bool

	loading, loaded bool
	notifiers       []chan bool
}

// MarkLoading marks this texture as currently loading. Only the renderer
// should call this, and as such you should (normally) never call this
// function.
func (t *Texture) MarkLoading() {
	t.Lock()
	defer t.Unlock()

	t.loading = true
}

// Loading tells if this texture is currently loading or not.
func (t *Texture) Loading() bool {
	t.RLock()
	defer t.RUnlock()

	return t.loading
}

// MarkLoaded marks this texture as loaded. Only the renderer should call this,
// and as such you should (normally) never call this function.
func (t *Texture) MarkLoaded() {
	t.Lock()

	t.loaded = true
	notifiers := t.notifiers
	t.notifiers = nil

	if t.clearOnLoad {
		t.Source = nil
	}

	t.Unlock()

	for _, notify := range notifiers {
		notify <- true
	}
}

// Loaded tells if this texture is currently loaded or not.
func (t *Texture) Loaded() bool {
	t.RLock()
	defer t.RUnlock()

	return t.loaded
}

// LoadNotify returns an channel on which true is sent once this texture is
// marked as loaded (normally by the renderer).
func (t *Texture) LoadNotify() chan bool {
	t.Lock()
	defer t.Unlock()

	notify := make(chan bool, 1)
	if t.loaded {
		notify <- true
		return notify
	}

	if t.notifiers == nil {
		t.notifiers = make([]chan bool, 0)
	}
	t.notifiers = append(t.notifiers, notify)
	return notify
}

// SetNativeIdentity specifies the native identity of this texture.
//
// This should mostly not be used (except in very rare, advanced cases).
func (t *Texture) SetNativeIdentity(identity interface{}) {
	t.Lock()
	defer t.Unlock()

	t.nativeIdentity = identity
}

// NativeIdentity returns the native identity of this texture.
//
// This should mostly not be used (except in very rare, advanced cases).
func (t *Texture) NativeIdentity() interface{} {
	t.RLock()
	defer t.RUnlock()

	return t.nativeIdentity
}

// SetAutoLoad specifies if this texture should automatically be loaded when
// a mesh is rendered using this texture.
func (t *Texture) SetAutoLoad(autoLoad bool) {
	t.Lock()
	defer t.Unlock()

	t.autoLoad = autoLoad
}

// AutoLoad tells if this texture will be automatically be loaded when a mesh
// is rendered using this texture.
//
// The mustWait option tells if meshes using this texture must wait for
// the texture to be loaded before rendering, or if incomplete renders can be
// performed (I.e. without the texture).
func (t *Texture) AutoLoad() bool {
	t.RLock()
	defer t.RUnlock()

	return t.autoLoad
}

// SetClearOnLoad specifies weather or not to clear the source of this texture
// when it is loaded.
//
// Default: true
func (t *Texture) SetClearOnLoad(clearOnLoad bool) {
	t.Lock()
	defer t.Unlock()

	t.clearOnLoad = clearOnLoad
}

// ClearOnLoad tells weather or not the source of this texture will be cleared
// when it is loaded.
//
// Default: true
func (t *Texture) ClearOnLoad() bool {
	t.RLock()
	defer t.RUnlock()

	return t.clearOnLoad
}

// SetAnisotropicDegree sets the anisotropic degree of this texture.
func (t *Texture) SetAnisotropicDegree(degree uint) {
	if degree <= 0 || degree > 9 {
		panic("SetAnisoDegree(): Anisotropic filtering degree must be 1-9")
	}
	t.Lock()
	defer t.Unlock()

	t.anisotropicDegree = degree
}

// AnisotropicDegree returns the anisotropic degree of this texture.
func (t *Texture) AnisotropicDegree() uint {
	t.RLock()
	defer t.RUnlock()

	return t.anisotropicDegree
}

// SetBorderColorVec4 sets the border color that will be used by this texture.
func (t *Texture) SetBorderColorVec4(color *math.Vec4) {
	t.Lock()
	defer t.Unlock()

	t.borderColor = color.Clamp(0, 1)
}

// BorderColorVec4 returns the border color that is used by this texture.
func (t *Texture) BorderColorVec4() *math.Vec4 {
	t.RLock()
	defer t.RUnlock()

	return t.borderColor.Copy()
}

// BorderColorFloat32 returns a float32 array of the border color used by this
// texture.
func (t *Texture) BorderColorFloat32() []float32 {
	t.RLock()
	defer t.RUnlock()

	return []float32{
		float32(t.borderColor.X),
		float32(t.borderColor.Y),
		float32(t.borderColor.Z),
		float32(t.borderColor.W),
	}
}

// SetBorderColor sets the border color of this texture as math.Real types.
func (t *Texture) SetBorderColor(r, g, b, a math.Real) {
	t.Lock()
	defer t.Unlock()

	r = r.Clamp(0, 1)
	g = r.Clamp(0, 1)
	b = r.Clamp(0, 1)
	a = r.Clamp(0, 1)

	t.borderColor = math.Vector4(r, g, b, a)
}

// BorderColor returns the border color of this texture as math.Real types.
func (t *Texture) BorderColor() (r, g, b, a math.Real) {
	t.RLock()
	defer t.RUnlock()

	return t.borderColor.X, t.borderColor.Y, t.borderColor.Z, t.borderColor.W
}

// SetCompressed sets weather or not this texture should be compressed in
// memory.
func (t *Texture) SetCompressed(compressed bool) {
	t.Lock()
	defer t.Unlock()

	t.compressed = compressed
}

// Compressed returns weather or not this texture is set to be compressed in
// memory.
func (t *Texture) Compressed() bool {
	t.RLock()
	defer t.RUnlock()

	return t.compressed
}

// New returns an new texture object, with these default values:
//
// An anisotropic filtering degree of 2.
//
// U, V, and W wrap modes of Repeat.
//
// An minification and magnification texture filter of type Linear.
//
// Texture compression enabled.
//
// Auto loading enabled.
func NewBase() *Texture {
	t := new(Texture)

	t.anisotropicDegree = 2
	t.wrapModeU = Repeat
	t.wrapModeV = Repeat
	t.wrapModeW = Repeat
	t.minFilter = Linear
	t.magFilter = Linear
	t.compressed = true
	t.autoLoad = true
	return t
}
