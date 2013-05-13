// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import(
	"image"
)

type BlitRenderable interface {
	PixelBlit(x, y uint, pixels image.Image)
	PixelClear(x, y, width, height uint)
}

