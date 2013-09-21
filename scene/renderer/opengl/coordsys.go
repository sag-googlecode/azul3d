// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package opengl

import (
	"code.google.com/p/azul3d/math"
)

var (
	// Get an matrix which will translate our matrix from ZUpRight to YUpRight
	coordSysConversion = math.CoordSysZUpRight.ConvertMat4(math.CoordSysYUpRight)
)

func regionToGL(totalWidth, totalHeight, x, y, width, height uint) (glX, glY, glWidth, glHeight int32) {
	var rx, ry, rWidth, rHeight uint
	if x == 0 && y == 0 && width == 0 && height == 0 {
		rWidth = totalWidth
		rHeight = totalHeight
	} else {
		rx = x
		ry = totalHeight - (y + height)

		rWidth = width
		rHeight = height

		if rx < 0 {
			rx = 0
		}
		if ry < 0 {
			ry = 0
		}
		if rWidth > totalWidth {
			rWidth = totalWidth
		}
		if rHeight > totalHeight {
			rHeight = totalHeight
		}
	}

	return int32(rx), int32(ry), int32(rWidth), int32(rHeight)
}
