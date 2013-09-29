// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/scene/shader"
)

var defaultShader = shader.New("DefaultShader")

func init() {
	/*
				=======================================================================
				============================ Vertex shader ============================
		 		=======================================================================
	*/
	defaultShader.SetSource([]byte(`
#version 110

attribute vec4 Vertex;
attribute vec4 Normal;
attribute vec4 Tangent;
attribute vec4 Bitangent;
attribute vec4 Color;
attribute vec4 BoneWeight;
attribute vec4 TextureCoord0;

uniform int NumTextureCoords;

varying vec2 tc0;

void main()
{
	gl_Position = gl_ModelViewProjectionMatrix * Vertex;
	tc0 = TextureCoord0.xy;
}
`), shader.Vertex)

	/*
				=======================================================================
				=========================== Fragment shader ===========================
		 		=======================================================================
	*/
	defaultShader.SetSource([]byte(`
#version 110

varying vec2 tc0;

uniform sampler2D Texture0;
uniform sampler2D Texture1;

uniform int NumTextures;

void main()
{
	vec4 Color0 = texture2D(Texture0, tc0);
	vec4 Color1 = texture2D(Texture1, tc0);
	gl_FragColor = mix(Color0, Color1, Color1.a);
}
`), shader.Fragment)
}
