// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	"azul3d.org/v0/math"
)

var (
	// Get an matrix which will translate our matrix from ZUpRight to YUpRight
	coordSysConversion = math.CoordSysZUpRight.ConvertMat4(math.CoordSysYUpRight)
)

func regionToGL(totalWidth, totalHeight uint, nx, ny, nWidth, nHeight float64) (glX, glY int32, glWidth, glHeight uint32) {
	// Convert normalized coordinates to pixel ones.
	x := uint((nx * float64(totalWidth)) + 0.5)
	y := uint((ny * float64(totalHeight)) + 0.5)
	width := uint((nWidth * float64(totalWidth)) + 0.5)
	height := uint((nHeight * float64(totalHeight)) + 0.5)

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
