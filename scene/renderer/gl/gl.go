// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	opengl "azul3d.org/v1/native/gl"
	"azul3d.org/v1/scene"
	"azul3d.org/v1/scene/camera"
	"azul3d.org/v1/scene/color"
	"azul3d.org/v1/scene/geom"
	"azul3d.org/v1/scene/shader"
	"azul3d.org/v1/scene/transparency"
	"strconv"
	"sync"
)

// Used as the native identity of geom.Mesh.
type GLBufferedMesh struct {
	Indices       uint32
	Vertices      uint32
	Normals       uint32
	Tangents      uint32
	Bitangents    uint32
	TextureCoords []uint32
	Colors        uint32
	BoneWeights   uint32
}

type GLShader struct {
	Program, Vertex, Fragment uint32
}

type Renderer struct {
	dcMakeCurrent, lcMakeCurrent         func(current bool)
	lcAccess                             sync.Mutex
	gl, lcgl                             *opengl.Context
	width, height                        uint
	samples, sampleBuffers               int32
	glArbMultisample, haveMSTransparency bool

	texturesToFreeAccess     sync.RWMutex
	texturesToFree           []uint32
	compressedTextureFormats []int32

	scissorStack       [][]float64
	meshesToFreeAccess sync.RWMutex
	meshesToFree       []*GLBufferedMesh

	regionBuffers        chan sortedRegions
	nodeBuffers          chan *sortedNodes
	inputTexturesBuffers chan []int32

	lastRegion                            *camera.Region
	lastColorClear                        color.Color
	lastDepthClear                        float64
	lastStencilClear                      uint
	lastScissorX, lastScissorY            int32
	lastScissorWidth, lastScissorHeight   uint32
	lastViewportX, lastViewportY          int32
	lastViewportWidth, lastViewportHeight uint32

	maxTextureSize, glslMaxVaryingFloats, glslMaxVertexShaderInputs,
	glslMaxFragmentShaderInputs int

	gpuName, gpuVendorName string
}

func (r *Renderer) clearRegion(rr renderRegion) {
	x, y, width, height := rr.Region.Region()
	r.viewport(x, y, width, height)

	var clearFlags int32
	if rr.Region.ClearColorActive() {
		clearFlags = clearFlags | opengl.COLOR_BUFFER_BIT

		// Set clear color to color of the (root) display.Node
		r.clearColor(rr.Region.ClearColor())
	}

	if rr.Region.ClearDepthActive() {
		clearFlags = clearFlags | opengl.DEPTH_BUFFER_BIT
		r.clearDepth(rr.Region.ClearDepth())
	}

	if rr.Region.ClearStencilActive() {
		clearFlags = clearFlags | opengl.STENCIL_BUFFER_BIT
		r.clearStencil(rr.Region.ClearStencil())
	}

	if clearFlags != 0 {
		// Restrict clear to region (viewport) area
		r.pushScissor(x, y, width, height)
		r.gl.Clear(uint32(clearFlags))
		r.popScissor()
	}
}

func (r *Renderer) useRegion(rr renderRegion) {
	x, y, width, height := rr.Region.Region()
	r.viewport(x, y, width, height)
}

func (r *Renderer) useNode(rn renderNode) {
	// Load the shader.
	r.loadShader(rn.shader, true)

	// Verify that the shader compiled.
	gls := rn.shader.NativeIdentity().(*GLShader)
	if gls.Program == 0 {
		return
	}

	// Use the shader.
	r.gl.UseProgram(gls.Program)

	// Drop textures from the renderNode that are not currently loaded.
	count := int32(0)
	for pairIndex, texPair := range rn.textures {
		if texPair.Type.NativeIdentity() == nil {
			// The texture is not loaded yet.
			if texPair.Type.AutoLoad() {
				// They want us to wait before rendering untill the texture is
				// fully loaded, do so now.
				r.loadTexture(texPair.Type, true)
			}

			// It's also possible that even if they wanted the texture to load
			// before rendering that it couldn't do so for some reason. (if no
			// source image was specified, for instance).
			if texPair.Type.NativeIdentity() == nil {
				// We couldn't load the texture, drop it from the textures used
				// for rendering.
				rn.textures = append(rn.textures[:pairIndex], rn.textures[pairIndex+1:]...)
				continue
			}
		}

		// We have a loaded texture for use with rendering.
		rn.inputTextures = append(rn.inputTextures, count)
		count++
	}

	// Add textures input.
	shader.SetInput(rn.node, "Textures", rn.inputTextures)

	// Set transparency state.
	switch rn.transparency {
	case transparency.Multisample:
		if r.haveMSTransparency {
			r.gl.Enable(opengl.SAMPLE_ALPHA_TO_COVERAGE)
		}

	case transparency.AlphaBlend:
		r.gl.Enable(opengl.BLEND)
		r.gl.BlendFunc(opengl.ONE, opengl.ONE_MINUS_SRC_ALPHA)
	}

	// Update the shader program's inputs.
	r.updateShaderInputs(rn.node, rn.shader, gls)

	// Bind textures.
	for i, texPair := range rn.textures {
		ident := texPair.Type.NativeIdentity().(uint32)
		r.gl.ActiveTexture(opengl.TEXTURE0 + int32(i))
		r.gl.BindTexture(opengl.TEXTURE_2D, ident)
	}
}

func (r *Renderer) finishNode(rn renderNode) {
	// Unbind textures (Note: we use rn.inputTextures because it is the only
	// texture slice of the proper size).
	for i := 0; i < len(rn.inputTextures); i++ {
		r.gl.ActiveTexture(opengl.TEXTURE0 + int32(i))
		r.gl.BindTexture(opengl.TEXTURE_2D, 0)
	}

	// Undo transparency state.
	switch rn.transparency {
	case transparency.Multisample:
		if r.haveMSTransparency {
			r.gl.Disable(opengl.SAMPLE_ALPHA_TO_COVERAGE)
		}

	case transparency.AlphaBlend:
		r.gl.Disable(opengl.BLEND)
	}
}

func (r *Renderer) drawMesh(rn renderNode, m *geom.Mesh) {
	// Verify that the shader compiled.
	gls := rn.shader.NativeIdentity().(*GLShader)
	if gls.Program == 0 {
		return
	}

	// Load the mesh into OpenGL
	r.loadMesh(m, true)

	bm := m.NativeIdentity().(*GLBufferedMesh)

	m.RLock()
	defer m.RUnlock()

	// VAA Index | Description
	//         0 - Vertices
	//         1 - Normals
	//         2 - Tangents
	//         3 - Bitangent
	//         4 - Colors
	//         5 - Bone Weights
	//         6 - Texture Coord 0
	//         7 - Texture Coord 1
	//       ... - Texture Coord ...
	//     23535 - Texture Coord 23535
	defer r.gl.BindBuffer(opengl.ARRAY_BUFFER, 0)

	getLocation := func(name string) (idx uint32, ok bool) {
		b := []byte(name)
		b = append(b, 0)
		i := r.gl.GetAttribLocation(gls.Program, &b[0])
		r.gl.Execute()
		if i >= 0 {
			return uint32(i), true
		}
		return 0, false
	}

	// Send vertices VBO
	idx, ok := getLocation("Vertex")
	if ok {
		r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Vertices)
		r.gl.EnableVertexAttribArray(idx)
		defer r.gl.DisableVertexAttribArray(idx)
		r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
	}

	if len(m.Normals) > 0 {
		// Send normal VBO
		idx, ok = getLocation("Normal")
		if ok {
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Normals)
			r.gl.EnableVertexAttribArray(idx)
			defer r.gl.DisableVertexAttribArray(idx)
			r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
		}
	}

	if len(m.Tangents) > 0 {
		// Send tangent VBO
		idx, ok = getLocation("Tangent")
		if ok {
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Tangents)
			r.gl.EnableVertexAttribArray(idx)
			defer r.gl.DisableVertexAttribArray(idx)
			r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
		}
	}

	if len(m.Bitangents) > 0 {
		// Send bitangent VBO
		idx, ok = getLocation("Bitangent")
		if ok {
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Bitangents)
			r.gl.EnableVertexAttribArray(idx)
			defer r.gl.DisableVertexAttribArray(idx)
			r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
		}
	}

	if len(m.Colors) > 0 {
		// Send color VBO
		idx, ok = getLocation("Color")
		if ok {
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Colors)
			r.gl.EnableVertexAttribArray(idx)
			defer r.gl.DisableVertexAttribArray(idx)
			r.gl.VertexAttribPointer(idx, 4, opengl.FLOAT, opengl.GLBool(false), 0, nil)
		}
	}

	if len(m.BoneWeights) > 0 {
		// Send BoneWeight VBO
		idx, ok = getLocation("BoneWeight")
		if ok {
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.BoneWeights)
			r.gl.EnableVertexAttribArray(idx)
			defer r.gl.DisableVertexAttribArray(idx)
			r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
		}
	}

	for index, _ := range m.TextureCoords {
		// Send TextureCoord VBO's
		stringIndex := strconv.Itoa(index)
		idx, ok = getLocation("TextureCoord" + stringIndex)
		if ok {
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.TextureCoords[index])
			r.gl.EnableVertexAttribArray(idx)
			defer r.gl.DisableVertexAttribArray(idx)
			r.gl.VertexAttribPointer(idx, 4, opengl.FLOAT, opengl.GLBool(false), 0, nil)
		}
	}

	if len(m.Indices) > 0 {
		r.gl.BindBuffer(opengl.ELEMENT_ARRAY_BUFFER, bm.Indices)
		r.gl.DrawElements(opengl.TRIANGLES, uint32(len(m.Indices)), opengl.UNSIGNED_INT, nil)
	} else {
		r.gl.DrawArrays(opengl.TRIANGLES, 0, uint32(len(m.Vertices)))
	}
}

func (r *Renderer) Render(rootNode *scene.Node) func() {
	// Collect and sort every camera's scene, etc.
	regionBuf := r.collect(rootNode)

	// Enter read lock
	r.texturesToFreeAccess.RLock()

	var texturesToFree []uint32
	if len(r.texturesToFree) > 0 {
		// Exit read lock
		r.texturesToFreeAccess.RUnlock()

		// Enter write lock
		r.texturesToFreeAccess.Lock()

		texturesToFree = r.texturesToFree
		r.texturesToFree = nil

		// Exit write lock
		r.texturesToFreeAccess.Unlock()
	} else {
		// Exit read lock
		r.texturesToFreeAccess.RUnlock()
	}

	// Enter read lock
	r.meshesToFreeAccess.RLock()

	var meshesToFree []*GLBufferedMesh
	if len(r.meshesToFree) > 0 {
		// Exit read lock
		r.meshesToFreeAccess.RUnlock()

		// Enter write lock
		r.meshesToFreeAccess.Lock()

		meshesToFree = r.meshesToFree
		r.meshesToFree = nil

		// Exit write lock
		r.meshesToFreeAccess.Unlock()
	} else {
		// Exit read lock
		r.meshesToFreeAccess.RUnlock()
	}

	// Build an render function to return
	return func() {
		defer r.releaseRegionBuf(regionBuf)

		if len(texturesToFree) > 0 {
			r.gl.DeleteTextures(uint32(len(texturesToFree)), &texturesToFree[0])
			r.gl.Execute()
		}

		if len(meshesToFree) > 0 {
			for _, bm := range meshesToFree {
				if bm.Indices != 0 {
					r.gl.DeleteBuffers(1, &bm.Indices)
				}
				if bm.Vertices != 0 {
					r.gl.DeleteBuffers(1, &bm.Vertices)
				}
				if bm.Normals != 0 {
					r.gl.DeleteBuffers(1, &bm.Normals)
				}
				if bm.Tangents != 0 {
					r.gl.DeleteBuffers(1, &bm.Tangents)
				}
				if bm.Bitangents != 0 {
					r.gl.DeleteBuffers(1, &bm.Bitangents)
				}
				if bm.Colors != 0 {
					r.gl.DeleteBuffers(1, &bm.Colors)
				}
				if bm.BoneWeights != 0 {
					r.gl.DeleteBuffers(1, &bm.BoneWeights)
				}
				r.gl.Execute()
				for _, vboId := range bm.TextureCoords {
					if vboId != 0 {
						r.gl.DeleteBuffers(1, &vboId)
						r.gl.Execute()
					}
				}
			}
		}

		// We shouldn't carry OpenGL state into the next frame -- other threads might be using GL
		// to render.
		r.lastRegion = nil

		r.stateClearClearColor()
		r.stateClearClearDepth()
		r.stateClearClearStencil()
		r.stateClearViewport()

		// First individually clear each region. They might overlap so we must
		// render each region after clearing them all.
		for _, rRegion := range regionBuf {
			r.clearRegion(rRegion)
		}

		// Render each region's slice of nodes.
		for _, rRegion := range regionBuf {
			r.useRegion(rRegion)
			for _, rNode := range rRegion.nodes.slice {
				r.useNode(rNode)
				for _, m := range rNode.meshes {
					r.drawMesh(rNode, m)
				}
				r.finishNode(rNode)
			}
		}

		// Flush and execute opengl commands.
		r.gl.Flush()
		r.gl.Execute()
	}
}

func (r *Renderer) Resize(width, height int) {
	r.width = uint(width)
	r.height = uint(height)
}

func NewRenderer(dcMakeCurrent, lcMakeCurrent func(current bool)) (*Renderer, error) {
	r := new(Renderer)

	r.regionBuffers = make(chan sortedRegions, 16)
	r.nodeBuffers = make(chan *sortedNodes, 128)
	r.inputTexturesBuffers = make(chan []int32, 256)

	r.dcMakeCurrent = dcMakeCurrent
	r.lcMakeCurrent = lcMakeCurrent

	r.dcMakeCurrent(true)

	r.gl = opengl.New()
	r.gl.SetBatching(true)

	r.lcgl = opengl.New()
	r.lcgl.SetBatching(true)

	r.glArbMultisample = r.gl.Extension("GL_ARB_multisample")
	if r.glArbMultisample {
		r.gl.GetIntegerv(opengl.SAMPLES, &r.samples)
		r.gl.GetIntegerv(opengl.SAMPLE_BUFFERS, &r.sampleBuffers)
		r.gl.Execute()

		r.gl.Enable(opengl.MULTISAMPLE)
	}

	r.haveMSTransparency = r.glArbMultisample && r.samples > 0 && r.sampleBuffers > 0

	r.gl.Enable(opengl.TEXTURE_2D)
	r.gl.Enable(opengl.SCISSOR_TEST)
	r.gl.Enable(opengl.PROGRAM_POINT_SIZE)

	r.gl.ClearDepth(1.0)
	r.gl.DepthFunc(opengl.LESS)
	r.gl.Enable(opengl.DEPTH_TEST)

	var numFormats int32
	r.gl.GetIntegerv(opengl.NUM_COMPRESSED_TEXTURE_FORMATS, &numFormats)
	r.gl.Execute()

	r.compressedTextureFormats = make([]int32, numFormats)
	r.gl.GetIntegerv(opengl.COMPRESSED_TEXTURE_FORMATS, &r.compressedTextureFormats[0])
	r.gl.Execute()

	return r, nil
}
