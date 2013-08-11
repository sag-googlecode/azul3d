// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package chippy

import (
	"image"
)

type Blitable interface {
	// PixelBlit blits the specified RGBA image onto the window, at the given X
	// and Y coordinates.
	PixelBlit(x, y uint, pixels *image.RGBA)

	// PixelClear clears the given rectangle on the window's client region.
	PixelClear(rect image.Rectangle)
}
