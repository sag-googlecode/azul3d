// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"code.google.com/p/azul3d/native/gl"
	"code.google.com/p/azul3d/scene/geom"
	"runtime"
	"unsafe"
)

func (r *Renderer) doUpdateMesh(ctx *opengl.Context, g *geom.Mesh, now bool) {
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
				ctx.DeleteBuffers(1, &bm.Indices)
				ctx.Execute()
				bm.Indices = 0
			} else {
				// Update Indices VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Indices)

				sz := int(unsafe.Sizeof(g.Indices[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Indices)),
					unsafe.Pointer(&g.Indices[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.IndicesChanged = false
		}

		// Update Vertices VBO
		if verticesChanged {
			if g.Vertices == nil || len(g.Vertices) == 0 {
				// Remove Vertices VBO
				ctx.DeleteBuffers(1, &bm.Vertices)
				ctx.Execute()
				bm.Vertices = 0
			} else {
				// Update Vertices VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Vertices)

				sz := int(unsafe.Sizeof(g.Vertices[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Vertices)),
					unsafe.Pointer(&g.Vertices[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.VerticesChanged = false
		}

		// Update Normals VBO
		if normalsChanged {
			if g.Normals == nil || len(g.Normals) == 0 {
				// Remove Normals VBO
				ctx.DeleteBuffers(1, &bm.Normals)
				ctx.Execute()
				bm.Normals = 0
			} else {
				// Update Normals VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Normals)

				sz := int(unsafe.Sizeof(g.Normals[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Normals)),
					unsafe.Pointer(&g.Normals[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.NormalsChanged = false
		}

		// Update Tangents VBO
		if tangentsChanged {
			if g.Tangents == nil || len(g.Tangents) == 0 {
				// Remove Tangents VBO
				ctx.DeleteBuffers(1, &bm.Tangents)
				ctx.Execute()
				bm.Tangents = 0
			} else {
				// Update Tangents VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Tangents)

				sz := int(unsafe.Sizeof(g.Tangents[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Tangents)),
					unsafe.Pointer(&g.Tangents[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.TangentsChanged = false
		}

		// Update Bitangents VBO
		if bitangentsChanged {
			if g.Bitangents == nil || len(g.Bitangents) == 0 {
				// Remove Bitangents VBO
				ctx.DeleteBuffers(1, &bm.Bitangents)
				ctx.Execute()
				bm.Bitangents = 0
			} else {
				// Update Bitangents VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Bitangents)

				sz := int(unsafe.Sizeof(g.Bitangents[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Bitangents)),
					unsafe.Pointer(&g.Bitangents[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.BitangentsChanged = false
		}

		// Update Colors VBO
		if colorsChanged {
			if g.Colors == nil || len(g.Colors) == 0 {
				// Remove Colors VBO
				ctx.DeleteBuffers(1, &bm.Colors)
				ctx.Execute()
				bm.Colors = 0
			} else {
				// Update Colors VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Colors)

				sz := int(unsafe.Sizeof(g.Colors[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.Colors)),
					unsafe.Pointer(&g.Colors[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.ColorsChanged = false
		}

		// Update BoneWeights VBO
		if boneWeightsChanged {
			if g.BoneWeights == nil || len(g.BoneWeights) == 0 {
				// Remove BoneWeights VBO
				ctx.DeleteBuffers(1, &bm.BoneWeights)
				ctx.Execute()
				bm.BoneWeights = 0
			} else {
				// Update BoneWeights VBO
				ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.BoneWeights)

				sz := int(unsafe.Sizeof(g.BoneWeights[0]))
				ctx.BufferData(
					opengl.ARRAY_BUFFER,
					uintptr(sz*len(g.BoneWeights)),
					unsafe.Pointer(&g.BoneWeights[0]),
					usageHint,
				)
				ctx.Execute()
			}
			g.BoneWeightsChanged = false
		}

		// Update TextureCoords VBO's
		for index, texCoords := range g.TextureCoords {
			changed, ok := g.TextureCoordsChanged[index]
			if ok && changed {
				if texCoords == nil || len(texCoords) == 0 {
					// Remove TextureCoord VBO
					ctx.DeleteBuffers(1, &bm.TextureCoords[index])
					ctx.Execute()
					bm.TextureCoords[index] = 0
				} else {
					// Update TextureCoord VBO
					ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.TextureCoords[index])

					sz := int(unsafe.Sizeof(texCoords[0]))
					ctx.BufferData(
						opengl.ARRAY_BUFFER,
						uintptr(sz*len(texCoords)),
						unsafe.Pointer(&texCoords[0]),
						usageHint,
					)
					ctx.Execute()
				}
				delete(g.TextureCoordsChanged, index)
			}
		}

		// Bind buffer 0 -- make no-buffer active
		ctx.BindBuffer(opengl.ARRAY_BUFFER, 0)

		// Wait for geom to be uploaded
		ctx.Finish()
		ctx.Execute()
	}
}

func (r *Renderer) doLoadMesh(ctx *opengl.Context, g *geom.Mesh, now bool) {
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
		ctx.GenBuffers(1, &bm.Vertices)
		ctx.Execute()
		ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Vertices)

		sz := int(unsafe.Sizeof(g.Vertices[0]))
		ctx.BufferData(
			opengl.ARRAY_BUFFER,
			uintptr(sz*len(g.Vertices)),
			unsafe.Pointer(&g.Vertices[0]),
			usageHint,
		)
		ctx.Execute()

		if len(g.Indices) > 0 {
			// Create indices buffer
			ctx.GenBuffers(1, &bm.Indices)
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Indices)

			sz := int(unsafe.Sizeof(g.Indices[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.Indices)),
				unsafe.Pointer(&g.Indices[0]),
				usageHint,
			)
			ctx.Execute()
		}

		if len(g.Normals) > 0 {
			// Create normals buffer
			ctx.GenBuffers(1, &bm.Normals)
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Normals)

			sz := int(unsafe.Sizeof(g.Normals[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.Normals)),
				unsafe.Pointer(&g.Normals[0]),
				usageHint,
			)
			ctx.Execute()
		}

		if len(g.Tangents) > 0 {
			// Create tangents buffer
			ctx.GenBuffers(1, &bm.Tangents)
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Tangents)

			sz := int(unsafe.Sizeof(g.Tangents[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.Tangents)),
				unsafe.Pointer(&g.Tangents[0]),
				usageHint,
			)
			ctx.Execute()
		}

		if len(g.Bitangents) > 0 {
			// Create bitangent buffer
			ctx.GenBuffers(1, &bm.Bitangents)
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Bitangents)

			sz := int(unsafe.Sizeof(g.Bitangents[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.Bitangents)),
				unsafe.Pointer(&g.Bitangents[0]),
				usageHint,
			)
			ctx.Execute()
		}

		if len(g.Colors) > 0 {
			// Create colors buffer
			ctx.GenBuffers(1, &bm.Colors)
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.Colors)

			sz := int(unsafe.Sizeof(g.Colors[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.Colors)),
				unsafe.Pointer(&g.Colors[0]),
				usageHint,
			)
			ctx.Execute()
		}

		if len(g.BoneWeights) > 0 {
			// Create bone weights buffer
			ctx.GenBuffers(1, &bm.BoneWeights)
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.BoneWeights)

			sz := int(unsafe.Sizeof(g.BoneWeights[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(g.BoneWeights)),
				unsafe.Pointer(&g.BoneWeights[0]),
				usageHint,
			)
			ctx.Execute()
		}

		bm.TextureCoords = make([]uint32, len(g.TextureCoords))
		for index, texCoords := range g.TextureCoords {
			// Create texture coordinates buffer
			ctx.GenBuffers(1, &bm.TextureCoords[index])
			ctx.Execute()
			ctx.BindBuffer(opengl.ARRAY_BUFFER, bm.TextureCoords[index])

			sz := int(unsafe.Sizeof(texCoords[0]))
			ctx.BufferData(
				opengl.ARRAY_BUFFER,
				uintptr(sz*len(texCoords)),
				unsafe.Pointer(&texCoords[0]),
				usageHint,
			)
			ctx.Execute()
		}

		// Bind buffer 0 -- make no-buffer active
		ctx.BindBuffer(opengl.ARRAY_BUFFER, 0)
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
	ctx.Finish()
	ctx.Execute()

	// Notify of completion
	g.MarkLoaded()
}

func (r *Renderer) loadMesh(g *geom.Mesh, now bool) {
	if g.IsLoaded() {
		if now {
			r.doUpdateMesh(r.gl, g, now)
		} else {
			go r.doUpdateMesh(r.lcgl, g, now)
		}
	} else {
		if now {
			r.doLoadMesh(r.gl, g, now)
		} else {
			go r.doLoadMesh(r.lcgl, g, now)
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
