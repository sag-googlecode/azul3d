// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/native/gl"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene/shader"
	"code.google.com/p/azul3d/scene/texture"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unsafe"
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
	gl                           *opengl.Context
	width, height                int

	texturesToFreeAccess     sync.RWMutex
	texturesToFree           []uint32
	compressedTextureFormats []int32

	scissorStack       [][]uint
	meshesToFreeAccess sync.RWMutex
	meshesToFree       []*GLBufferedMesh

	lastRegion                                                          *camera.Region
	lastColorClear                                                      *math.Vec4
	lastDepthClear                                                      math.Real
	lastStencilClear                                                    uint
	lastScissorX, lastScissorY, lastScissorWidth, lastScissorHeight     uint
	lastViewportX, lastViewportY, lastViewportWidth, lastViewportHeight uint

	maxTextureCoords, maxTextureLayers, maxTextureSize, maxMemoryBytes int
	gpuName, gpuVendorName                                             string
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

	r.loadMesh(g, true)

	// Load the shader into OpenGL
	s, ok := shader.Active(current.node)
	if !ok {
		// No custom shader, use default one.
		s = defaultShader
	}
	r.loadShader(s, true)
	gls := s.NativeIdentity().(*GLShader)
	if gls.Program == 0 {
		return
	}
	r.gl.UseProgram(gls.Program)

	// Add texture shader inputs, and drop textures that are not loaded
	layeredTextures := texture.Textures(current.node)
	count := 0
	for layer, tex := range layeredTextures {
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
				delete(layeredTextures, layer)
				continue
			}
		}

		// Add texture input
		shader.SetInput(current.node, "Texture"+strconv.Itoa(count), int32(count))
		count++
	}

	shader.SetInput(current.node, "Projection", current.projection)
	shader.SetInput(current.node, "ModelView", current.modelView)
	shader.SetInput(current.node, "ModelViewProjection", current.modelViewProjection)

	r.updateShaderInputs(current.node, s, gls)

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
		for _, tex := range layeredTextures {
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

		for count := 0; count < len(layeredTextures); count++ {
			r.gl.ActiveTexture(opengl.TEXTURE0 + int32(count))
			r.gl.BindTexture(opengl.TEXTURE_2D, 0)
		}
	}
}

func (r *Renderer) loadTexture(t texture.Type, now bool) {
	if !texture.IsValid(t) {
		panic("LoadTexture(): Invalid texture type!")
	}

	if t.Loaded() || t.Loading() {
		return
	}

	// Create the texture
	switch tex := t.(type) {
	case *texture.Texture2D:
		img := tex.Image()
		if img == nil || len(img.Pix) == 0 {
			// Cannot upload texture without image specified
			return
		}

	default:
		panic("LoadTexture(): Renderer does not support texture type!")
	}

	t.MarkLoading()

	doLoadTexture := func() {
		if now {
			// Release our display context
			r.dcMakeCurrent(false)

			// Later on we will use it
			defer r.dcMakeCurrent(true)
		} else {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
		}

		// Lock the loading context to thread
		r.lcAccess.Lock()
		defer r.lcAccess.Unlock()

		// Make the loading context active for this OS thread.
		r.lcMakeCurrent(true)

		// Later on release the loading context.
		defer r.lcMakeCurrent(false)

		// Create the texture
		switch tex := t.(type) {
		case *texture.Texture2D:
			r.createTexture2D(tex)
		}

		// Wait for texture to be uploaded
		r.gl.Finish()

		// Notify of completion
		t.MarkLoaded()
	}
	if now {
		doLoadTexture()
	} else {
		go doLoadTexture()
	}
}

// As with other renderer calls, this is made inside an single OS thread only.
//
// But we may push it to an different thread if we wish to (we want to, of
// course).
func (r *Renderer) LoadTexture(t texture.Type) {
	r.loadTexture(t, false)
}

func (r *Renderer) loadMesh(g *geom.Mesh, now bool) {
	doUpdateMesh := func() {
		bm := g.NativeIdentity().(*GLBufferedMesh)

		g.Lock()
		defer g.Unlock()

		indicesChanged := g.IndicesChanged
		verticesChanged := g.VerticesChanged
		normalsChanged := g.NormalsChanged
		tangentsChanged := g.TangentsChanged
		bitangentsChanged := g.BitangentsChanged
		colorsChanged := g.ColorsChanged
		boneWeightsChanged := g.BoneWeightsChanged
		anyTextureCoordsChanged := len(g.TextureCoordsChanged) > 0

		// Check to avoid a (possibly expensive) context switch.
		if indicesChanged || verticesChanged || normalsChanged || tangentsChanged || bitangentsChanged || colorsChanged || boneWeightsChanged || anyTextureCoordsChanged {

			if now {
				// Release our display context
				r.dcMakeCurrent(false)

				// Later on we will use it
				defer r.dcMakeCurrent(true)
			} else {
				runtime.LockOSThread()
				defer runtime.UnlockOSThread()
			}

			// Lock the loading context to thread
			r.lcAccess.Lock()
			defer r.lcAccess.Unlock()

			// Make the loading context active for this OS thread.
			r.lcMakeCurrent(true)

			// Later on release the loading context.
			defer r.lcMakeCurrent(false)

			usageHint := opengl.STATIC_DRAW
			if g.Hint == geom.Dynamic {
				usageHint = opengl.DYNAMIC_DRAW
			}

			// Update Indices VBO
			if indicesChanged {
				if g.Indices == nil || len(g.Indices) == 0 {
					// Remove Indices VBO
					r.gl.DeleteBuffers(1, &bm.Indices)
					bm.Indices = 0
				} else {
					// Update Indices VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Indices)

					sz := int(unsafe.Sizeof(g.Indices[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.Indices)),
						unsafe.Pointer(&g.Indices[0]),
						usageHint,
					)
				}
				g.IndicesChanged = false
			}

			// Update Vertices VBO
			if verticesChanged {
				if g.Vertices == nil || len(g.Vertices) == 0 {
					// Remove Vertices VBO
					r.gl.DeleteBuffers(1, &bm.Vertices)
					bm.Vertices = 0
				} else {
					// Update Vertices VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Vertices)

					sz := int(unsafe.Sizeof(g.Vertices[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.Vertices)),
						unsafe.Pointer(&g.Vertices[0]),
						usageHint,
					)
				}
				g.VerticesChanged = false
			}

			// Update Normals VBO
			if normalsChanged {
				if g.Normals == nil || len(g.Normals) == 0 {
					// Remove Normals VBO
					r.gl.DeleteBuffers(1, &bm.Normals)
					bm.Normals = 0
				} else {
					// Update Normals VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Normals)

					sz := int(unsafe.Sizeof(g.Normals[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.Normals)),
						unsafe.Pointer(&g.Normals[0]),
						usageHint,
					)
				}
				g.NormalsChanged = false
			}

			// Update Tangents VBO
			if tangentsChanged {
				if g.Tangents == nil || len(g.Tangents) == 0 {
					// Remove Tangents VBO
					r.gl.DeleteBuffers(1, &bm.Tangents)
					bm.Tangents = 0
				} else {
					// Update Tangents VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Tangents)

					sz := int(unsafe.Sizeof(g.Tangents[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.Tangents)),
						unsafe.Pointer(&g.Tangents[0]),
						usageHint,
					)
				}
				g.TangentsChanged = false
			}

			// Update Bitangents VBO
			if bitangentsChanged {
				if g.Bitangents == nil || len(g.Bitangents) == 0 {
					// Remove Bitangents VBO
					r.gl.DeleteBuffers(1, &bm.Bitangents)
					bm.Bitangents = 0
				} else {
					// Update Bitangents VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Bitangents)

					sz := int(unsafe.Sizeof(g.Bitangents[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.Bitangents)),
						unsafe.Pointer(&g.Bitangents[0]),
						usageHint,
					)
				}
				g.BitangentsChanged = false
			}

			// Update Colors VBO
			if colorsChanged {
				if g.Colors == nil || len(g.Colors) == 0 {
					// Remove Colors VBO
					r.gl.DeleteBuffers(1, &bm.Colors)
					bm.Colors = 0
				} else {
					// Update Colors VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Colors)

					sz := int(unsafe.Sizeof(g.Colors[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.Colors)),
						unsafe.Pointer(&g.Colors[0]),
						usageHint,
					)
				}
				g.ColorsChanged = false
			}

			// Update BoneWeights VBO
			if boneWeightsChanged {
				if g.BoneWeights == nil || len(g.BoneWeights) == 0 {
					// Remove BoneWeights VBO
					r.gl.DeleteBuffers(1, &bm.BoneWeights)
					bm.BoneWeights = 0
				} else {
					// Update BoneWeights VBO
					r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.BoneWeights)

					sz := int(unsafe.Sizeof(g.BoneWeights[0]))
					r.gl.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(g.BoneWeights)),
						unsafe.Pointer(&g.BoneWeights[0]),
						usageHint,
					)
				}
				g.BoneWeightsChanged = false
			}

			// Update TextureCoords VBO's
			for index, texCoords := range g.TextureCoords {
				changed, ok := g.TextureCoordsChanged[index]
				if ok && changed {
					if texCoords == nil || len(texCoords) == 0 {
						// Remove TextureCoord VBO
						r.gl.DeleteBuffers(1, &bm.TextureCoords[index])
						bm.TextureCoords[index] = 0
					} else {
						// Update TextureCoord VBO
						r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.TextureCoords[index])

						sz := int(unsafe.Sizeof(texCoords[0]))
						r.gl.BufferData(
							opengl.ARRAY_BUFFER,
							uintptr(sz*len(texCoords)),
							unsafe.Pointer(&texCoords[0]),
							usageHint,
						)
					}
					delete(g.TextureCoordsChanged, index)
				}
			}

			// Bind buffer 0 -- make no-buffer active
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, 0)

			//r.gl.Flush()

			// Wait for geom to be uploaded
			r.gl.Finish()

		}
	}

	doLoadMesh := func() {
		if now {
			// Release our display context
			r.dcMakeCurrent(false)

			// Later on we will use it
			defer r.dcMakeCurrent(true)
		} else {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
		}

		// Lock the loading context to thread
		r.lcAccess.Lock()
		defer r.lcAccess.Unlock()

		// Make the loading context active for this OS thread.
		r.lcMakeCurrent(true)

		// Later on release the loading context.
		defer r.lcMakeCurrent(false)

		usageHint := opengl.STATIC_DRAW
		if g.Hint == geom.Dynamic {
			usageHint = opengl.DYNAMIC_DRAW
		}

		// Lock the mesh
		g.RLock()

		// Create the object
		bm := new(GLBufferedMesh)

		if len(g.Vertices) > 0 {
			// Create vertices buffer
			r.gl.GenBuffers(1, &bm.Vertices)
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Vertices)

			sz := int(unsafe.Sizeof(g.Vertices[0]))
			r.gl.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.Vertices)),
				unsafe.Pointer(&g.Vertices[0]),
				usageHint,
			)

			if len(g.Indices) > 0 {
				// Create indices buffer
				r.gl.GenBuffers(1, &bm.Indices)
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Indices)

				sz := int(unsafe.Sizeof(g.Indices[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Indices)),
					unsafe.Pointer(&g.Indices[0]),
					usageHint,
				)
			}

			if len(g.Normals) > 0 {
				// Create normals buffer
				r.gl.GenBuffers(1, &bm.Normals)
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Normals)

				sz := int(unsafe.Sizeof(g.Normals[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Normals)),
					unsafe.Pointer(&g.Normals[0]),
					usageHint,
				)
			}

			if len(g.Tangents) > 0 {
				// Create tangents buffer
				r.gl.GenBuffers(1, &bm.Tangents)
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Tangents)

				sz := int(unsafe.Sizeof(g.Tangents[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Tangents)),
					unsafe.Pointer(&g.Tangents[0]),
					usageHint,
				)
			}

			if len(g.Bitangents) > 0 {
				// Create bitangent buffer
				r.gl.GenBuffers(1, &bm.Bitangents)
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Bitangents)

				sz := int(unsafe.Sizeof(g.Bitangents[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Bitangents)),
					unsafe.Pointer(&g.Bitangents[0]),
					usageHint,
				)
			}

			if len(g.Colors) > 0 {
				// Create colors buffer
				r.gl.GenBuffers(1, &bm.Colors)
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.Colors)

				sz := int(unsafe.Sizeof(g.Colors[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Colors)),
					unsafe.Pointer(&g.Colors[0]),
					usageHint,
				)
			}

			if len(g.BoneWeights) > 0 {
				// Create bone weights buffer
				r.gl.GenBuffers(1, &bm.BoneWeights)
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.BoneWeights)

				sz := int(unsafe.Sizeof(g.BoneWeights[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.BoneWeights)),
					unsafe.Pointer(&g.BoneWeights[0]),
					usageHint,
				)
			}

			bm.TextureCoords = make([]uint32, len(g.TextureCoords))
			for index, texCoords := range g.TextureCoords {
				// Create texture coordinates buffer
				r.gl.GenBuffers(1, &bm.TextureCoords[index])
				r.gl.BindBuffer(opengl.ARRAY_BUFFER, bm.TextureCoords[index])

				sz := int(unsafe.Sizeof(texCoords[0]))
				r.gl.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(texCoords)),
					unsafe.Pointer(&texCoords[0]),
					usageHint,
				)
			}

			// Bind buffer 0 -- make no-buffer active
			r.gl.BindBuffer(opengl.ARRAY_BUFFER, 0)
		}

		// Unlock the mesh
		g.RUnlock()

		// Store the native identity
		g.SetNativeIdentity(bm)

		// Attach finalizer
		runtime.SetFinalizer(g, func(g *geom.Mesh) {
			r.meshesToFreeAccess.Lock()
			defer r.meshesToFreeAccess.Unlock()

			bm := g.NativeIdentity().(*GLBufferedMesh)
			r.meshesToFree = append(r.meshesToFree, bm)
		})

		// Wait for geom to be uploaded
		r.gl.Finish()

		// Notify of completion
		g.MarkLoaded()
	}

	if g.IsLoaded() {
		if now {
			doUpdateMesh()
		} else {
			go doUpdateMesh()
		}
	} else {
		if now {
			doLoadMesh()
		} else {
			go doLoadMesh()
		}
	}
}

// As with other renderer calls, this is made inside an single OS thread only.
//
// But we may push it to an different thread if we wish to (we want to, of
// course).
func (r *Renderer) LoadMesh(m *geom.Mesh) {
	r.loadMesh(m, false)
}

func (r *Renderer) updateShaderInput(gls *GLShader, name string, value interface{}) {
	bts := []byte(name)
	bts = append(bts, 0)
	location := r.gl.GetUniformLocation(gls.Program, &bts[0])
	if location < 0 {
		return
	}

	switch v := value.(type) {
	case float32:
		r.gl.Uniform1fv(location, 1, &v)
	case []float32:
		r.gl.Uniform1fv(location, uint32(len(v)), &v[0])

	case shader.Vec2:
		r.gl.Uniform2fv(location, 1, &v[0])
	case []shader.Vec2:
		r.gl.Uniform2fv(location, uint32(len(v)), &v[0][0])

	case shader.Vec3:
		r.gl.Uniform3fv(location, 1, &v[0])
	case []shader.Vec3:
		r.gl.Uniform3fv(location, uint32(len(v)), &v[0][0])

	case shader.Vec4:
		r.gl.Uniform4fv(location, 1, &v[0])
	case []shader.Vec4:
		r.gl.Uniform4fv(location, uint32(len(v)), &v[0][0])

	case int32:
		r.gl.Uniform1iv(location, 1, &v)
	case []int32:
		r.gl.Uniform1iv(location, uint32(len(v)), &v[0])

	case shader.Vec2i:
		r.gl.Uniform2iv(location, 1, &v[0])
	case []shader.Vec2i:
		r.gl.Uniform2iv(location, uint32(len(v)), &v[0][0])

	case shader.Vec3i:
		r.gl.Uniform3iv(location, 1, &v[0])
	case []shader.Vec3i:
		r.gl.Uniform3iv(location, uint32(len(v)), &v[0][0])

	case shader.Vec4i:
		r.gl.Uniform4iv(location, 1, &v[0])
	case []shader.Vec4i:
		r.gl.Uniform4iv(location, uint32(len(v)), &v[0][0])

	case uint32:
		r.gl.Uniform1uiv(location, 1, &v)
	case []uint32:
		r.gl.Uniform1uiv(location, uint32(len(v)), &v[0])

	case shader.Vec2ui:
		r.gl.Uniform2uiv(location, 1, &v[0])
	case []shader.Vec2ui:
		r.gl.Uniform2uiv(location, uint32(len(v)), &v[0][0])

	case shader.Vec3ui:
		r.gl.Uniform3uiv(location, 1, &v[0])
	case []shader.Vec3ui:
		r.gl.Uniform3uiv(location, uint32(len(v)), &v[0][0])

	case shader.Vec4ui:
		r.gl.Uniform4uiv(location, 1, &v[0])
	case []shader.Vec4ui:
		r.gl.Uniform4uiv(location, uint32(len(v)), &v[0][0])

	case shader.Mat2:
		r.gl.UniformMatrix2fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat2:
		r.gl.UniformMatrix2fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat3:
		r.gl.UniformMatrix3fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat3:
		r.gl.UniformMatrix3fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat4:
		r.gl.UniformMatrix4fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat4:
		r.gl.UniformMatrix4fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat2x3:
		r.gl.UniformMatrix2x3fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat2x3:
		r.gl.UniformMatrix2x3fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat3x2:
		r.gl.UniformMatrix3x2fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat3x2:
		r.gl.UniformMatrix3x2fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat2x4:
		r.gl.UniformMatrix2x4fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat2x4:
		r.gl.UniformMatrix2x4fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat4x2:
		r.gl.UniformMatrix4x2fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat4x2:
		r.gl.UniformMatrix4x2fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat3x4:
		r.gl.UniformMatrix3x4fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat3x4:
		r.gl.UniformMatrix3x4fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat4x3:
		r.gl.UniformMatrix4x3fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat4x3:
		r.gl.UniformMatrix4x3fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	default:
		panic("Invalid shader input type!")
	}
}

func (r *Renderer) updateShaderInputs(n *scene.Node, s *shader.Shader, gls *GLShader) {
	r.gl.UseProgram(gls.Program)

	for name, value := range shader.Inputs(n) {
		r.updateShaderInput(gls, name, value)
	}
}

func (r *Renderer) loadShader(s *shader.Shader, now bool) {
	if s.Loaded() {
		return
	}

	doLoadShader := func() {
		if now {
			// Release our display context
			r.dcMakeCurrent(false)

			// Later on we will use it
			defer r.dcMakeCurrent(true)
		} else {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
		}

		// Lock the loading context to thread
		r.lcAccess.Lock()
		defer r.lcAccess.Unlock()

		// Make the loading context active for this OS thread.
		r.lcMakeCurrent(true)

		// Later on release the loading context.
		defer r.lcMakeCurrent(false)

		// Create the shader
		gls := new(GLShader)

		shaderCompilerLog := func(s uint32) []byte {
			var ok int32
			r.gl.GetShaderiv(s, opengl.COMPILE_STATUS, &ok)
			if ok == 0 {
				// Shader compiler error
				var logSize int32
				r.gl.GetShaderiv(s, opengl.INFO_LOG_LENGTH, &logSize)

				log := make([]byte, logSize)
				r.gl.GetShaderInfoLog(s, uint32(logSize), nil, &log[0])
				return log
			}
			return nil
		}

		appendError := func(err []byte) {
			s.SetError(append(s.Error(), err...))
		}

		vertSource := s.Source(shader.Vertex)
		if vertSource != nil {
			sVertSource := string(vertSource)
			sVertSource = strings.Replace(sVertSource, " ", "", -1)
			sVertSource = strings.Replace(sVertSource, "\t", "", -1)
			sVertSource = strings.Replace(sVertSource, "\n", "", -1)
			sVertSource = strings.Replace(sVertSource, "\r", "", -1)
			sVertSource = strings.Replace(sVertSource, "\r\n", "", -1)
			if len(sVertSource) == 0 {
				// Behavior is undefined (normally driver crashes).
				appendError([]byte(s.Name() + " | Vertex shader with no source code.\n"))

			} else {
				// Build vertex shader
				gls.Vertex = r.gl.CreateShader(opengl.VERTEX_SHADER)
				lengths := int32(len(vertSource))
				sources := &vertSource[0]
				r.gl.ShaderSource(gls.Vertex, 1, &sources, &lengths)
				r.gl.CompileShader(gls.Vertex)

				log := shaderCompilerLog(gls.Vertex)
				if log != nil {
					// Sanity
					gls.Vertex = 0

					appendError([]byte(s.Name() + " | Vertex shader errors:\n"))
					appendError(log)
				}
			}
		}

		fragSource := s.Source(shader.Fragment)
		if fragSource != nil {
			sFragSource := string(fragSource)
			sFragSource = strings.Replace(sFragSource, " ", "", -1)
			sFragSource = strings.Replace(sFragSource, "\t", "", -1)
			sFragSource = strings.Replace(sFragSource, "\n", "", -1)
			sFragSource = strings.Replace(sFragSource, "\r", "", -1)
			sFragSource = strings.Replace(sFragSource, "\r\n", "", -1)
			if len(sFragSource) == 0 {
				// Behavior is undefined (normally driver crashes).
				appendError([]byte(s.Name() + " | Fragment shader with no source code.\n"))

			} else {
				// Build fragment shader
				gls.Fragment = r.gl.CreateShader(opengl.FRAGMENT_SHADER)
				lengths := int32(len(fragSource))
				sources := &fragSource[0]
				r.gl.ShaderSource(gls.Fragment, 1, &sources, &lengths)
				r.gl.CompileShader(gls.Fragment)

				log := shaderCompilerLog(gls.Fragment)
				if log != nil {
					// Sanity
					gls.Fragment = 0

					appendError([]byte(s.Name() + " | Fragment shader errors:\n"))
					appendError(log)
				}
			}
		}

		if gls.Vertex != 0 && gls.Fragment != 0 {
			gls.Program = r.gl.CreateProgram()

			r.gl.AttachShader(gls.Program, gls.Vertex)
			r.gl.AttachShader(gls.Program, gls.Fragment)
			r.gl.LinkProgram(gls.Program)

			// Link shader program
			var ok int32
			r.gl.GetProgramiv(gls.Program, opengl.LINK_STATUS, &ok)
			if ok == 0 {
				// Program linker error
				var logSize int32
				r.gl.GetProgramiv(gls.Program, opengl.INFO_LOG_LENGTH, &logSize)

				log := make([]byte, logSize)
				r.gl.GetProgramInfoLog(gls.Program, uint32(logSize), nil, &log[0])

				// Sanity
				gls.Program = 0

				appendError([]byte(s.Name() + " | Linker errors:\n"))
				appendError(log)
			}
		}

		// Store the native identity
		s.SetNativeIdentity(gls)

		// Wait for shader to be compiled
		r.gl.Finish()

		// Notify of completion
		s.MarkLoaded()
	}
	if now {
		doLoadShader()
	} else {
		go doLoadShader()
	}
}

// As with other renderer calls, this is made inside an single OS thread only.
//
// But we may push it to an different thread if we wish to (we want to, of
// course).
func (r *Renderer) LoadShader(s *shader.Shader) {
	r.loadShader(s, false)
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
				for _, vboId := range bm.TextureCoords {
					if vboId != 0 {
						r.gl.DeleteBuffers(1, &vboId)
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
	}
}

func (r *Renderer) Resize(width, height int) {
	r.width = width
	r.height = height

	// Reset viewport now
	r.viewport(0, 0, width, height)
}

func NewRenderer(dcMakeCurrent, lcMakeCurrent func(current bool)) (*Renderer, error) {
	r := new(Renderer)

	r.dcMakeCurrent = dcMakeCurrent
	r.lcMakeCurrent = lcMakeCurrent

	r.dcMakeCurrent(true)

	r.gl = opengl.New()
	if r.gl == nil {
		return nil, fmt.Errorf("No support for OpenGL 2.0 found.")
	}

	r.gl.Enable(opengl.TEXTURE_2D)
	r.gl.Enable(opengl.SCISSOR_TEST)
	r.gl.Enable(opengl.PROGRAM_POINT_SIZE)

	r.gl.ClearDepth(1.0)
	r.gl.DepthFunc(opengl.LESS)
	r.gl.Enable(opengl.DEPTH_TEST)

	r.gl.Enable(opengl.BLEND)
	r.gl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)

	var numFormats int32
	r.gl.GetIntegerv(opengl.NUM_COMPRESSED_TEXTURE_FORMATS, &numFormats)

	r.compressedTextureFormats = make([]int32, numFormats)
	r.gl.GetIntegerv(opengl.COMPRESSED_TEXTURE_FORMATS, &r.compressedTextureFormats[0])

	return r, nil
}
