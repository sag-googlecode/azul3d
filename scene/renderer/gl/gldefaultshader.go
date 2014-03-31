// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"azul3d.org/v1/scene/shader"
)

var defaultShader = shader.New("DefaultShader")

func init() {
	/*
				=======================================================================
				============================ Vertex shader ============================
		 		=======================================================================
	*/
	defaultShader.SetSource([]byte(`
#version 120

attribute vec4 Vertex;
attribute vec4 Normal;
attribute vec4 Tangent;
attribute vec4 Bitangent;
attribute vec4 Color;
attribute vec4 BoneWeight;
attribute vec4 TextureCoord0;

varying vec2 tc0;

uniform mat4 Projection;
uniform mat4 ModelView;
uniform mat4 ModelViewProjection;

void main()
{
	tc0 = TextureCoord0.xy;
	//gl_FrontColor = Color;
	gl_Position = ModelViewProjection * Vertex;
}
`), shader.Vertex)

	/*
				=======================================================================
				=========================== Fragment shader ===========================
		 		=======================================================================
	*/
	defaultShader.SetSource([]byte(`
#version 120

varying vec2 tc0;

const int NumTextures = 2;
uniform sampler2D[NumTextures] Textures;
uniform bool BinaryTransparency;

void main()
{
	vec4 final = vec4(0, 0, 0, 0);
	for(int t = 0; t < NumTextures; t++) {
		vec4 tc = texture2D(Textures[t], tc0);
		final = mix(final, tc, tc.a);
	}

	if (BinaryTransparency && final.a < 0.5) {
		discard;
	}
	gl_FragColor = final;
}
`), shader.Fragment)
}
