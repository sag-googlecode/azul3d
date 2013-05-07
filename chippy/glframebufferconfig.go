// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

package chippy

import (
	"fmt"
)

// GLConfig represents an single opengl (frame buffer / pixel) configuration.
type GLConfig struct {
	backend_GLConfig

	// An GLConfig is 'valid' when the user did not create it themself.
	valid bool

	// Tells weather this configuration is hardware accelerated or uses some software
	// implementation version of OpenGL.
	//
	// Note: Most software implementations are very low OpenGL versions. (I.e. GL 1.1)
	Accelerated bool

	// Number of anti-aliasing samples this configuration supports
	Samples uint8

	// The number of bits that represent an color per pixel in the frame buffer.
	RedBits, GreenBits, BlueBits, AlphaBits uint8

	// The number of bits that represent an color per pixel in the accumulation buffer.
	//
	// Note: GLSL shaders can perform an much better job of anything you would be trying to do with
	// the accumulation buffer.
	AccumRedBits, AccumGreenBits, AccumBlueBits, AccumAlphaBits uint8

	// The number of auxiliary buffers available.
	//
	// Note: Auxiliary buffers are very rarely supported on most OpenGL implementations, For more
	// information about this see the following forum URL:
	//     http://www.opengl.org/discussion_boards/showthread.php/171060-auxiliary-buffers
	AuxBuffers uint8

	// The number of bits that represent an pixel in the depth buffer.
	DepthBits uint8

	// The number of bits that represent an pixel in the stencil buffer.
	StencilBits uint8

	// Weather this frame buffer configuration is double buffered.
	DoubleBuffered bool

	// Weather this frame buffer configuration is stereoscopic capable.
	StereoScopic bool
}

func (c *GLConfig) panicUnlessValid() {
	if !c.valid {
		panic("Invalid GLBufferFormat; did you attempt to create it yourself?")
	}
}

func (c *GLConfig) String() string {
	return fmt.Sprintf("GLConfig(Accelerated=%t, %dbpp[%d,%d,%d,%d], AccumBits=[%d,%d,%d,%d], AuxBuffers=%d, DepthBits=%d, StencilBits=%d, DoubleBuffered=%t, StereoScopic=%t)", c.Accelerated, c.RedBits+c.GreenBits+c.BlueBits+c.AlphaBits, c.RedBits, c.GreenBits, c.BlueBits, c.AlphaBits, c.AccumRedBits, c.AccumGreenBits, c.AccumBlueBits, c.AccumAlphaBits, c.AuxBuffers, c.DepthBits, c.StencilBits, c.DoubleBuffered, c.StereoScopic)
}

// Equals tells weather this GLConfig equals the other GLFrameBufferConfig, by comparing
// each attribute.
func (c *GLConfig) Equals(other *GLConfig) bool {
	o := other

	if c.Accelerated != o.Accelerated || c.RedBits != o.RedBits || c.GreenBits != o.GreenBits || c.BlueBits != o.BlueBits || c.AlphaBits != o.AlphaBits || c.AccumRedBits != o.AccumRedBits || c.AccumGreenBits != o.AccumGreenBits || c.AccumBlueBits != o.AccumBlueBits || c.AccumAlphaBits != o.AccumAlphaBits || c.AuxBuffers != o.AuxBuffers || c.DepthBits != o.DepthBits || c.StencilBits != o.StencilBits || c.DoubleBuffered != o.DoubleBuffered || c.StereoScopic != o.StereoScopic {
		return false
	}
	return true
}

var (
	GLWorstConfig = &GLConfig{
		Accelerated:    false,
		RedBits:        0,
		GreenBits:      0,
		BlueBits:       0,
		AlphaBits:      0,
		AccumRedBits:   0,
		AccumGreenBits: 0,
		AccumBlueBits:  0,
		AccumAlphaBits: 0,
		AuxBuffers:     0,
		DepthBits:      0,
		StencilBits:    0,
		DoubleBuffered: false,
		StereoScopic:   false,
	}

	GLWorstHWConfig = &GLConfig{
		Accelerated:    true,
		RedBits:        0,
		GreenBits:      0,
		BlueBits:       0,
		AlphaBits:      0,
		AccumRedBits:   0,
		AccumGreenBits: 0,
		AccumBlueBits:  0,
		AccumAlphaBits: 0,
		AuxBuffers:     0,
		DepthBits:      0,
		StencilBits:    0,
		DoubleBuffered: false,
		StereoScopic:   false,
	}

	GLBestConfig = &GLConfig{
		Accelerated:    true,
		RedBits:        255,
		GreenBits:      255,
		BlueBits:       255,
		AlphaBits:      255,
		AccumRedBits:   255,
		AccumGreenBits: 255,
		AccumBlueBits:  255,
		AccumAlphaBits: 255,
		AuxBuffers:     255,
		DepthBits:      255,
		StencilBits:    255,
		DoubleBuffered: true,
		StereoScopic:   false,
	}
)

// GLChooseConfig chooses an appropriate configuration from the slice of possible configurations.
//
// The returned configuration will have at least minConfig's attributes, or nil will be returned if
// there is no configuration that has at least minConfig's attributes.
//
// The returned configuration will have no greater than maxConfig's attributes, or nil will be
// returned if there is no configuration that is below maxConfig's attributes.
//
// You may use the predefined GLWorstConfig and GLBestConfig variables if they suite your case.
func GLChooseConfig(possible []*GLConfig, minConfig, maxConfig *GLConfig) *GLConfig {
	min := minConfig
	max := maxConfig

	// Remove any which are below minConfig
	var removed []*GLConfig
	for _, c := range possible {
		if c.RedBits < min.RedBits {
			continue
		}
		if c.GreenBits < min.GreenBits {
			continue
		}
		if c.BlueBits < min.BlueBits {
			continue
		}
		if c.AlphaBits < min.AlphaBits {
			continue
		}

		if c.AccumRedBits < min.AccumRedBits {
			continue
		}
		if c.AccumGreenBits < min.AccumGreenBits {
			continue
		}
		if c.AccumBlueBits < min.AccumBlueBits {
			continue
		}
		if c.AccumAlphaBits < min.AccumAlphaBits {
			continue
		}

		if c.AuxBuffers < min.AuxBuffers {
			continue
		}
		if c.DepthBits < min.DepthBits {
			continue
		}
		if c.StencilBits < min.StencilBits {
			continue
		}

		if min.Accelerated && !c.Accelerated {
			continue
		}
		if min.DoubleBuffered && !c.DoubleBuffered {
			continue
		}
		if min.StereoScopic && !c.StereoScopic {
			continue
		}
		removed = append(removed, c)
	}
	possible = removed

	// Remove any which are above maxConfig
	removed = make([]*GLConfig, 0)
	for _, c := range possible {
		if c.RedBits > max.RedBits {
			continue
		}
		if c.GreenBits > max.GreenBits {
			continue
		}
		if c.BlueBits > max.BlueBits {
			continue
		}
		if c.AlphaBits > max.AlphaBits {
			continue
		}

		if c.AccumRedBits > max.AccumRedBits {
			continue
		}
		if c.AccumGreenBits > max.AccumGreenBits {
			continue
		}
		if c.AccumBlueBits > max.AccumBlueBits {
			continue
		}
		if c.AccumAlphaBits > max.AccumAlphaBits {
			continue
		}

		if c.AuxBuffers > max.AuxBuffers {
			continue
		}
		if c.DepthBits > max.DepthBits {
			continue
		}
		if c.StencilBits > max.StencilBits {
			continue
		}

		if c.Accelerated && !max.Accelerated {
			continue
		}
		if c.DoubleBuffered && !max.DoubleBuffered {
			continue
		}
		if c.StereoScopic && !max.StereoScopic {
			continue
		}
		removed = append(removed, c)
	}
	possible = removed

	// Sort by order of closest-to maxConfig in order of
	//
	// Accelerated
	// RedBits, GreenBits, BlueBits, AlphaBits
	// DoubleBuffered
	// DepthBits
	// StencilBits
	// StereoScopic
	// AccumRedBits, AccumGreenBits, AccumBlueBits, AccumAlphaBits
	// AuxBuffers

	if len(possible) == 0 {
		return nil
	}

	bc := possible[0]
	for _, t := range possible {
		if !t.Accelerated && bc.Accelerated {
			continue
		}

		if t.RedBits < bc.RedBits || t.GreenBits < bc.GreenBits || t.BlueBits < bc.BlueBits || t.AlphaBits < bc.AlphaBits {
			continue
		}

		if !t.DoubleBuffered && bc.DoubleBuffered {
			continue
		}

		if t.DepthBits < bc.DepthBits {
			continue
		}

		if t.StencilBits < bc.StencilBits {
			continue
		}

		if !t.StereoScopic && bc.StereoScopic {
			continue
		}

		if t.AccumRedBits < bc.AccumRedBits || t.AccumGreenBits < bc.AccumGreenBits {
			continue
		}

		if t.AccumBlueBits < bc.AccumBlueBits || t.AccumAlphaBits < bc.AccumAlphaBits {
			continue
		}

		if t.AuxBuffers < bc.AuxBuffers {
			continue
		}

		// Whew, so t is greater than bc (best config), store that one instead!
		bc = t
	}
	return bc
}
