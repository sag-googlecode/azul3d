// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	opengl "azul3d.org/native/gl"
	"azul3d.org/scene/texture"
	"runtime"
)

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

	doLoadTexture := func(ctx *opengl.Context) {
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
			r.createTexture2D(ctx, tex)
		}

		// Wait for texture to be uploaded
		ctx.Finish()
		ctx.Execute()

		// Notify of completion
		t.MarkLoaded()
	}
	if now {
		doLoadTexture(r.gl)
	} else {
		go doLoadTexture(r.lcgl)
	}
}

// As with other renderer calls, this is made inside an single OS thread only.
//
// But we may push it to an different thread if we wish to (we want to, of
// course).
func (r *Renderer) LoadTexture(t texture.Type) {
	r.loadTexture(t, false)
}
