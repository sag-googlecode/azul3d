// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"image"
	"sync"
)

type nilRenderer struct {
	// The MSAA state.
	msaa struct {
		sync.RWMutex
		enabled bool
	}
}

func (n *nilRenderer) Bounds() image.Rectangle {
	return image.Rect(0, 0, 640, 480)
}

func (n *nilRenderer) Precision() Precision {
	return Precision{
		RedBits:     255,
		GreenBits:   255,
		BlueBits:    255,
		AlphaBits:   255,
		DepthBits:   255,
		StencilBits: 255,
	}
}

func (n *nilRenderer) GPUInfo() GPUInfo {
	return GPUInfo{
		MaxTextureSize:  8096,
		AlphaToCoverage: true,
	}
}
func (n *nilRenderer) Download(r image.Rectangle, complete chan image.Image) {
	complete <- image.NewRGBA(image.Rect(0, 0, 1, 1))
}
func (n *nilRenderer) SetMSAA(msaa bool) {
	n.msaa.Lock()
	n.msaa.enabled = msaa
	n.msaa.Unlock()
}
func (n *nilRenderer) MSAA() (msaa bool) {
	n.msaa.RLock()
	msaa = n.msaa.enabled
	n.msaa.RUnlock()
	return
}
func (n *nilRenderer) Clear(r image.Rectangle, bg Color)           {}
func (n *nilRenderer) ClearDepth(r image.Rectangle, depth float64) {}
func (n *nilRenderer) ClearStencil(r image.Rectangle, stencil int) {}
func (n *nilRenderer) Draw(r image.Rectangle, o *Object, c *Camera) {
}
func (n *nilRenderer) Render() {}

func (n *nilRenderer) LoadMesh(m *Mesh, done chan *Mesh) {
	m.Lock()
	m.Loaded = true
	m.ClearData()
	m.Unlock()
	select {
	case done <- m:
	default:
	}
}
func (n *nilRenderer) LoadTexture(t *Texture, done chan *Texture) {
	t.Lock()
	t.Loaded = true
	t.ClearData()
	t.Unlock()
	select {
	case done <- t:
	default:
	}
}
func (n *nilRenderer) LoadShader(s *Shader, done chan *Shader) {
	s.Lock()
	s.Loaded = true
	s.ClearData()
	s.Unlock()
	select {
	case done <- s:
	default:
	}
}

func (n *nilRenderer) RenderToTexture(t *Texture) Canvas {
	return new(nilRenderer)
}

// Nil returns a renderer that does not actually render anything.
func Nil() Renderer {
	r := new(nilRenderer)
	r.msaa.enabled = true
	return r
}
