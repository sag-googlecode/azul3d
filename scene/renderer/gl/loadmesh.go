// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"azul3d.org/native/gl"
	"azul3d.org/scene/geom"
	"runtime"
	"unsafe"
)

func (r *Renderer) createVBO(ctx *opengl.Context) (vboId uint32) {
	// Generate new VBO.
	ctx.GenBuffers(1, &vboId)
	ctx.Execute()
	return
}

func (r *Renderer) updateVBO(ctx *opengl.Context, usageHint int32, dataSize uintptr, dataLength int, data unsafe.Pointer, vboId uint32) {
	// Bind the VBO now.
	ctx.BindBuffer(opengl.ARRAY_BUFFER, vboId)

	// Fill the VBO with the data.
	ctx.BufferData(
		opengl.ARRAY_BUFFER,
		dataSize*uintptr(dataLength),
		data,
		usageHint,
	)
	ctx.Execute()
}

func (r *Renderer) deleteVBO(ctx *opengl.Context, vboId uint32) {
	// Delete the VBO.
	ctx.DeleteBuffers(1, &vboId)
	ctx.Execute()
}

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
				// Delete indices VBO.
				r.deleteVBO(ctx, bm.Indices)
				bm.Indices = 0
			} else {
				// Update indices VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.Indices[0]),
					len(g.Indices),
					unsafe.Pointer(&g.Indices[0]),
					bm.Indices,
				)
			}
			g.IndicesChanged = false
		}

		// Update Vertices VBO
		if verticesChanged {
			if g.Vertices == nil || len(g.Vertices) == 0 {
				// Delete vertices VBO.
				r.deleteVBO(ctx, bm.Vertices)
				bm.Vertices = 0
			} else {
				// Update vertices VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.Vertices[0]),
					len(g.Vertices),
					unsafe.Pointer(&g.Vertices[0]),
					bm.Vertices,
				)
			}
			g.VerticesChanged = false
		}

		// Update Normals VBO
		if normalsChanged {
			if g.Normals == nil || len(g.Normals) == 0 {
				// Delete normals VBO.
				r.deleteVBO(ctx, bm.Normals)
				bm.Normals = 0
			} else {
				// Update normals VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.Normals[0]),
					len(g.Normals),
					unsafe.Pointer(&g.Normals[0]),
					bm.Normals,
				)
			}
			g.NormalsChanged = false
		}

		// Update Tangents VBO
		if tangentsChanged {
			if g.Tangents == nil || len(g.Tangents) == 0 {
				// Delete tangents VBO.
				r.deleteVBO(ctx, bm.Tangents)
				bm.Tangents = 0
			} else {
				// Update tangents VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.Tangents[0]),
					len(g.Tangents),
					unsafe.Pointer(&g.Tangents[0]),
					bm.Tangents,
				)
			}
			g.TangentsChanged = false
		}

		// Update Bitangents VBO
		if bitangentsChanged {
			if g.Bitangents == nil || len(g.Bitangents) == 0 {
				// Delete bitangents VBO.
				r.deleteVBO(ctx, bm.Bitangents)
				bm.Bitangents = 0
			} else {
				// Update bitangents VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.Bitangents[0]),
					len(g.Bitangents),
					unsafe.Pointer(&g.Bitangents[0]),
					bm.Bitangents,
				)
			}
			g.BitangentsChanged = false
		}

		// Update Colors VBO
		if colorsChanged {
			if g.Colors == nil || len(g.Colors) == 0 {
				// Delete colors VBO.
				r.deleteVBO(ctx, bm.Colors)
				bm.Colors = 0
			} else {
				// Update colors VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.Colors[0]),
					len(g.Colors),
					unsafe.Pointer(&g.Colors[0]),
					bm.Colors,
				)
			}
			g.ColorsChanged = false
		}

		// Update BoneWeights VBO
		if boneWeightsChanged {
			if g.BoneWeights == nil || len(g.BoneWeights) == 0 {
				// Delete bone weights VBO.
				r.deleteVBO(ctx, bm.BoneWeights)
				bm.BoneWeights = 0
			} else {
				// Update bone weights VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(g.BoneWeights[0]),
					len(g.BoneWeights),
					unsafe.Pointer(&g.BoneWeights[0]),
					bm.BoneWeights,
				)
			}
			g.BoneWeightsChanged = false
		}

		// Update TextureCoords VBO's
		for index, texCoords := range g.TextureCoords {
			changed, ok := g.TextureCoordsChanged[index]
			if !ok || !changed {
				continue
			}

			if texCoords == nil || len(texCoords) == 0 {
				// Delete texture coord VBO.
				r.deleteVBO(ctx, bm.TextureCoords[index])
				bm.TextureCoords[index] = 0
			} else {
				// Update texture coord VBO.
				r.updateVBO(
					ctx, usageHint,
					unsafe.Sizeof(texCoords[0]),
					len(texCoords),
					unsafe.Pointer(&texCoords[0]),
					bm.TextureCoords[index],
				)
			}
			delete(g.TextureCoordsChanged, index)
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
		// Create vertices VBO.
		bm.Vertices = r.createVBO(ctx)
		r.updateVBO(
			ctx, usageHint,
			unsafe.Sizeof(g.Vertices[0]),
			len(g.Vertices),
			unsafe.Pointer(&g.Vertices[0]),
			bm.Vertices,
		)

		if len(g.Indices) > 0 {
			// Create indices VBO.
			bm.Indices = r.createVBO(ctx)
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(g.Indices[0]),
				len(g.Indices),
				unsafe.Pointer(&g.Indices[0]),
				bm.Indices,
			)
		}

		if len(g.Normals) > 0 {
			// Create normals VBO.
			bm.Normals = r.createVBO(ctx)
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(g.Normals[0]),
				len(g.Normals),
				unsafe.Pointer(&g.Normals[0]),
				bm.Normals,
			)
		}

		if len(g.Tangents) > 0 {
			// Create tangents VBO.
			bm.Tangents = r.createVBO(ctx)
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(g.Tangents[0]),
				len(g.Tangents),
				unsafe.Pointer(&g.Tangents[0]),
				bm.Tangents,
			)
		}

		if len(g.Bitangents) > 0 {
			// Create bitangents VBO.
			bm.Bitangents = r.createVBO(ctx)
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(g.Bitangents[0]),
				len(g.Bitangents),
				unsafe.Pointer(&g.Bitangents[0]),
				bm.Bitangents,
			)
		}

		if len(g.Colors) > 0 {
			// Create colors VBO.
			bm.Colors = r.createVBO(ctx)
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(g.Colors[0]),
				len(g.Colors),
				unsafe.Pointer(&g.Colors[0]),
				bm.Colors,
			)
		}

		if len(g.BoneWeights) > 0 {
			// Create bone weights VBO.
			bm.BoneWeights = r.createVBO(ctx)
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(g.BoneWeights[0]),
				len(g.BoneWeights),
				unsafe.Pointer(&g.BoneWeights[0]),
				bm.BoneWeights,
			)
		}

		bm.TextureCoords = make([]uint32, len(g.TextureCoords))
		for index, texCoords := range g.TextureCoords {
			// Create texture coordinates VBO.
			vboId := r.createVBO(ctx)
			bm.TextureCoords[index] = vboId
			r.updateVBO(
				ctx, usageHint,
				unsafe.Sizeof(texCoords[0]),
				len(texCoords),
				unsafe.Pointer(&texCoords[0]),
				vboId,
			)
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
