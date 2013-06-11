// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"image"
)

type BlitRenderable interface {
	// PixelBlit blits the specified 32bpp RGBA image onto the window, at the specified X and Y
	// coordinates.
	PixelBlit(x, y uint, pixels image.Image)

	// PixelClear clears an select portion of the window's client region, it begins clearing at X
	// and Y, and extends until (x + width, y + height).
	PixelClear(x, y, width, height uint)
}
