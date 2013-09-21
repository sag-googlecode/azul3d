// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package texture

type Filter uint8

// ValidMag tells if this filter is valid to be used for as a magnification
// filter, that is it must be either Nearest or Linear.
func (f Filter) ValidMag() bool {
	switch f {
	case Nearest:
		return true
	case Linear:
		return true
	}
	return false
}

// ValidMin tells if this filter is valid to be used as a minification filter,
// that is if it i one of the predefined filter constants in this package.
func (f Filter) ValidMin() bool {
	switch f {
	case Nearest:
		return true
	case Linear:
		return true
	case NearestMipmapNearest:
		return true
	case LinearMipmapNearest:
		return true
	case NearestMipmapLinear:
		return true
	case LinearMipmapLinear:
		return true
	}
	return false
}

// Mipmapped tells if the filter is one of the following or not:
//
//  NearestMipmapNearest
//  LinearMipmapNearest
//  NearestMipmapLinear
//  LinearMipmapLinear
//
func (f Filter) Mipmapped() bool {
	switch f {
	case NearestMipmapNearest:
		return true
	case LinearMipmapNearest:
		return true
	case NearestMipmapLinear:
		return true
	case LinearMipmapLinear:
		return true
	}
	return false
}

const (
	// Samples the nearest pixel.
	Nearest Filter = iota

	// Samples the four closest pixels, and linearly interpolates them.
	Linear

	// Samples point from the closest mipmap.
	NearestMipmapNearest

	// Bilinear filter the pixel from the closest mipmap.
	LinearMipmapNearest

	// Samples the pixel from two closest mipmaps, and linearly blends.
	NearestMipmapLinear

	// (Trilinear filtering) Bilinearly filters the pixel from two mipmaps, and
	// linearly blends the result.
	LinearMipmapLinear
)

// SetMinFilter specifies the minification filter of this texture. The filter
// must be an valid minification filter, or an panic will occur.
func (t *Texture) SetMinFilter(filter Filter) {
	if !filter.ValidMin() {
		panic("SetMinFilter(): filter must be an valid minification filter!")
	}

	t.Lock()
	defer t.Unlock()

	t.minFilter = filter
}

// MinFilter returns the minification filter of this texture.
func (t *Texture) MinFilter() Filter {
	t.RLock()
	defer t.RUnlock()

	return t.minFilter
}

// SetMagFilter specifies the magnification filter of this texture. The filter
// must be an valid magnification filter, or an panic will occur.
func (t *Texture) SetMagFilter(filter Filter) {
	if !filter.ValidMag() {
		panic("SetMagFilter(): filter must be an valid magnification filter!")
	}

	t.Lock()
	defer t.Unlock()

	t.magFilter = filter
}

// MagFilter returns the magnification filter of this texture.
func (t *Texture) MagFilter() Filter {
	t.RLock()
	defer t.RUnlock()

	return t.magFilter
}
