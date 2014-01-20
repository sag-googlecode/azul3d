// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"code.google.com/p/azul3d/math"
)

var (
	// Get an matrix which will translate our matrix from ZUpRight to YUpRight
	coordSysConversion = math.CoordSysZUpRight.ConvertMat4(math.CoordSysYUpRight)
)

func regionToGL(totalWidth, totalHeight, x, y, width, height uint) (glX, glY int32, glWidth, glHeight uint32) {
	if x == 0 && y == 0 && width == 0 && height == 0 {
		// A region of all zeros means the entire area
		return 0, 0, uint32(totalWidth), uint32(totalHeight)
	}

	// Flip Y axis and leave X axis untouched
	glX = int32(x)
	glY = int32(totalHeight - (y + height))
	if glY < 0 {
		glY = 0
	}

	// Verify max sizes
	if width > totalWidth {
		width = totalWidth
	}
	if height > totalHeight {
		height = totalHeight
	}
	glWidth = uint32(width)
	glHeight = uint32(height)
	return
}
