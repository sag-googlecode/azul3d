// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package renderer

import (
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/shader"
	"code.google.com/p/azul3d/scene/texture"
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
