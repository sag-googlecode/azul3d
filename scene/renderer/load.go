// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package renderer

import (
	"azul3d.org/v0/scene"
	"azul3d.org/v0/scene/geom"
	"azul3d.org/v0/scene/shader"
	"azul3d.org/v0/scene/texture"
)

func LoadTexture(n *scene.Node, t texture.Type) {
	if !texture.IsValid(t) {
		panic("LoadTexture(): Invalid texture type.")
	}
	o := mustGetObj(n)
	o.renderFrame(func() {
		o.renderer.LoadTexture(t)
	})
}

func LoadMesh(n *scene.Node, m *geom.Mesh) {
	o := mustGetObj(n)
	o.renderFrame(func() {
		o.renderer.LoadMesh(m)
	})
}

func LoadShader(n *scene.Node, s *shader.Shader) {
	o := mustGetObj(n)
	o.renderFrame(func() {
		o.renderer.LoadShader(s)
	})
}
