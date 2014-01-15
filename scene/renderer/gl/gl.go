// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/native/gl"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/shader"
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
	dcMakeCurrent, lcMakeCurrent func(current bool)
	lcAccess                     sync.Mutex
	gl, lcgl                     *opengl.Context
	width, height                int
	glArbMultisample             bool

	texturesToFreeAccess     sync.RWMutex
	texturesToFree           []uint32
	compressedTextureFormats []int32

	scissorStack       [][]uint
	meshesToFreeAccess sync.RWMutex
	meshesToFree       []*GLBufferedMesh
	lastSortedGeoms    sortedGeoms

	lastRegion                                                          *camera.Region
	lastColorClear                                                      *math.Vec4
	lastDepthClear                                                      math.Real
	lastStencilClear                                                    uint
	lastScissorX, lastScissorY, lastScissorWidth, lastScissorHeight     uint
	lastViewportX, lastViewportY, lastViewportWidth, lastViewportHeight uint

	maxTextureSize, glslMaxVaryingFloats, glslMaxVertexShaderInputs,
	glslMaxFragmentShaderInputs int

	gpuName, gpuVendorName string
}

func (r *Renderer) useRegion(region *camera.Region) {
	if r.lastRegion != nil && r.lastRegion.Id() == region.Id() {
		return
	}
	r.lastRegion = region

	x, y, width, height := region.Region()
	r.viewport(x, y, int(width), int(height))

	var clearFlags int32
	if region.ClearColorActive() {
		clearFlags = clearFlags | opengl.COLOR_BUFFER_BIT

		// Set clear color to color of the (root) display.Node
		r.clearColor(region.ClearColor())
	}

	if region.ClearDepthActive() {
		clearFlags = clearFlags | opengl.DEPTH_BUFFER_BIT
		r.clearDepth(region.ClearDepth())
	}

	if region.ClearStencilActive() {
		clearFlags = clearFlags | opengl.STENCIL_BUFFER_BIT
		r.clearStencil(region.ClearStencil())
	}

	if clearFlags != 0 {
		// Restrict clear to region (viewport) area
		r.pushScissor(x, y, width, height)
		r.gl.Clear(uint32(clearFlags))
		r.popScissor()
	}
}

func (r *Renderer) drawGeom(current *sortedGeom) {
	g := current.geom

	if g.IsHidden() {
		return
	}

	// Load the mesh into OpenGL
	r.loadMesh(g, true)

	// Load the shader into OpenGL
	r.loadShader(current.shader, true)

	gls := current.shader.NativeIdentity().(*GLShader)
	if gls.Program == 0 {
		return
	}
	r.gl.UseProgram(gls.Program)

	// Add texture shader inputs, and drop textures that are not loaded
	inputTextures := make([]int32, len(current.textures))
	count := int32(0)
	for layer, tex := range current.textures {
		// If the texture is not loaded yet, then we simply never render with
		// it.
		//
		// Additionally, we start loading it if auto loading is enabled for it.
		i := tex.NativeIdentity()
		if i == nil {
			if tex.AutoLoad() {
				// Maybe we can force the texture to load right now
				r.loadTexture(tex, true)
			}

			// It's possible the texture could not load (no source, for instance)
			if tex.NativeIdentity() == nil {
				// Not loaded yet; remove from map.
				delete(current.textures, layer)
				continue
			}
		}

		// Add texture input
		inputTextures[count] = count
		count++
	}

	shader.SetInput(current.node, "Textures", inputTextures)

	switch current.transparency {
	case scene.Multisample:
		if r.glArbMultisample {
			r.gl.Enable(opengl.SAMPLE_ALPHA_TO_COVERAGE)
			defer r.gl.Disable(opengl.SAMPLE_ALPHA_TO_COVERAGE)
		} else {
			shader.SetInput(current.node, "BinaryTransparency", int32(1))
		}
	case scene.Transparency:
		r.gl.Enable(opengl.BLEND)
		defer r.gl.Disable(opengl.BLEND)
		r.gl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
	}

	r.updateShaderInputs(current.node, current.shader, gls)

	bm := g.NativeIdentity().(*GLBufferedMesh)

	g.RLock()
	defer g.RUnlock()

	if len(g.Vertices) > 0 {
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

		if len(g.Normals) > 0 {
			// Send normal VBO
			idx, ok = getLocation("Normal")
			if ok {
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Normals)
				r.gl.EnableVertexAttribArray(idx)
				defer r.gl.DisableVertexAttribArray(idx)
				r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
			}
		}

		if len(g.Tangents) > 0 {
			// Send tangent VBO
			idx, ok = getLocation("Tangent")
			if ok {
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Tangents)
				r.gl.EnableVertexAttribArray(idx)
				defer r.gl.DisableVertexAttribArray(idx)
				r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
			}
		}

		if len(g.Bitangents) > 0 {
			// Send bitangent VBO
			idx, ok = getLocation("Bitangent")
			if ok {
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Bitangents)
				r.gl.EnableVertexAttribArray(idx)
				defer r.gl.DisableVertexAttribArray(idx)
				r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
			}
		}

		if len(g.Colors) > 0 {
			// Send color VBO
			idx, ok = getLocation("Color")
			if ok {
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Colors)
				r.gl.EnableVertexAttribArray(idx)
				defer r.gl.DisableVertexAttribArray(idx)
				r.gl.VertexAttribPointer(idx, 4, opengl.FLOAT, opengl.GLBool(false), 0, nil)
			}
		}

		if len(g.BoneWeights) > 0 {
			// Send BoneWeight VBO
			idx, ok = getLocation("BoneWeight")
			if ok {
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.BoneWeights)
				r.gl.EnableVertexAttribArray(idx)
				defer r.gl.DisableVertexAttribArray(idx)
				r.gl.VertexAttribPointer(idx, 3, opengl.FLOAT, opengl.GLBool(false), 0, nil)
			}
		}

		for index, _ := range g.TextureCoords {
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

		count := 0
		for _, tex := range current.textures {
			ident := tex.NativeIdentity().(uint32)
			r.gl.ActiveTexture(opengl.TEXTURE0 + int32(count))
			r.gl.BindTexture(opengl.TEXTURE_2D, ident)
			count++
		}

		if len(g.Indices) > 0 {
			r.gl.BindBuffer(opengl.ELEMENT_ARRAY_BUFFER, bm.Indices)
			r.gl.DrawElements(opengl.TRIANGLES, uint32(len(g.Indices)), opengl.UNSIGNED_INT, nil)
		} else {
			r.gl.DrawArrays(opengl.TRIANGLES, 0, uint32(len(g.Vertices)))
		}

		for count := 0; count < len(current.textures); count++ {
			r.gl.ActiveTexture(opengl.TEXTURE0 + int32(count))
			r.gl.BindTexture(opengl.TEXTURE_2D, 0)
		}
	}
}

func (r *Renderer) Render(rootNode *scene.Node) func() {
	// Locate cameras in the scene who are active, have an scene specified, and have at least one
	// camera region.
	var cameras []*scene.Node

	rootNode.Traverse(func(i int, n *scene.Node) bool {
		if (!n.Hidden() || n.ShownThrough()) && camera.Is(n) {
			cameras = append(cameras, n)
		}

		// Please continue traversal
		return true
	})

	// Sort all visible geom nodes by their active sorters and other properties.
	geoms := r.sortGeoms(rootNode, cameras)

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

		for _, g := range geoms {
			r.useRegion(g.region)

			if g.node != nil {
				r.drawGeom(g)
			}
		}

		r.gl.Flush()
		r.gl.Execute()
	}
}

func (r *Renderer) Resize(width, height int) {
	r.width = width
	r.height = height

	// Reset viewport now
	r.viewport(0, 0, width, height)
	r.gl.Execute()
}

func NewRenderer(dcMakeCurrent, lcMakeCurrent func(current bool)) (*Renderer, error) {
	r := new(Renderer)

	r.dcMakeCurrent = dcMakeCurrent
	r.lcMakeCurrent = lcMakeCurrent

	r.dcMakeCurrent(true)

	r.gl = opengl.New()
	r.gl.SetBatching(true)

	r.lcgl = opengl.New()
	r.lcgl.SetBatching(true)

	r.glArbMultisample = r.gl.Extension("GL_ARB_multisample")
	if r.glArbMultisample {
		r.gl.Enable(opengl.MULTISAMPLE)
	}

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
