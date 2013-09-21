// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

// WrapMode specifies an single mode to be used for either the U or V axis of
// an texture's wrapping.
type WrapMode uint8

// Valid returns true if this WrapMode is one of the predefined constants in
// this package.
func (w WrapMode) Valid() bool {
	switch w {
	case Repeat:
		return true
	case Clamp:
		return true
	case BorderColor:
		return true
	case Mirror:
		return true
	}
	return false
}

// String returns an string representation of this WrapMode.
func (w WrapMode) String() string {
	switch w {
	case Repeat:
		return "Repeat"
	case Clamp:
		return "Clamp"
	case BorderColor:
		return "BorderColor"
	case Mirror:
		return "Mirror"
	}
	return "Invalid"
}

const (
	// The extra area of the texture is repeated into infinity.
	Repeat WrapMode = iota

	// The extra area of the texture is represented by stretching the edge
	// pixels out into infinity.
	Clamp

	// The extra area of the texture is represented by the border color
	// specified on the texture.
	BorderColor

	// The extra area of the texture is represented by itself mirrored into
	// infinity.
	Mirror
)

// SetWrapModeU specifies the wrap mode for the U axis of this texture.
func (t *Texture) SetWrapModeU(u WrapMode) {
	t.Lock()
	defer t.Unlock()

	t.wrapModeU = u
}

// WrapModeU returns the wrap mode for the U axis of this texture.
func (t *Texture) WrapModeU() WrapMode {
	t.RLock()
	defer t.RUnlock()

	return t.wrapModeU
}

// SetWrapModeV specifies the wrap mode for the V axis of this texture.
func (t *Texture) SetWrapModeV(v WrapMode) {
	t.Lock()
	defer t.Unlock()

	t.wrapModeV = v
}

// WrapModeV returns the wrap mode for the V axis of this texture.
func (t *Texture) WrapModeV() WrapMode {
	t.RLock()
	defer t.RUnlock()

	return t.wrapModeV
}

// SetWrapModeW specifies the wrap mode for the W axis of this texture.
func (t *Texture) SetWrapModeW(w WrapMode) {
	t.Lock()
	defer t.Unlock()

	t.wrapModeW = w
}

// WrapModeW returns the wrap mode for the W axis of this texture.
func (t *Texture) WrapModeW() WrapMode {
	t.RLock()
	defer t.RUnlock()

	return t.wrapModeW
}
