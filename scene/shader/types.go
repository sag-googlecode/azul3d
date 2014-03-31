// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package shader

import "azul3d.org/v1/math"

func ConvertMat4(m math.Mat4) Mat4 {
	return Mat4{
		[4]float32{float32(m[0][0]), float32(m[0][1]), float32(m[0][2]), float32(m[0][3])},
		[4]float32{float32(m[1][0]), float32(m[1][1]), float32(m[1][2]), float32(m[1][3])},
		[4]float32{float32(m[2][0]), float32(m[2][1]), float32(m[2][2]), float32(m[2][3])},
		[4]float32{float32(m[3][0]), float32(m[3][1]), float32(m[3][2]), float32(m[3][3])},
	}
}

// Vec2 is a two-component 32-bit float vector shader input type.
type Vec2 [2]float32

// Vec3 is a three-component 32-bit float vector shader input type.
type Vec3 [3]float32

// Vec4 is a four-component 32-bit float vector shader input type.
type Vec4 [4]float32

// Vec2i is a two-component 32-bit int vector shader input type.
type Vec2i [2]int32

// Vec3i is a three-component 32-bit int vector shader input type.
type Vec3i [3]int32

// Vec4i is a four-component 32-bit int vector shader input type.
type Vec4i [4]int32

// Vec2ui is a two-component 32-bit unsigned int vector shader input type.
type Vec2ui [2]uint32

// Vec3ui is a three-component 32-bit unsigned int vector shader input type.
type Vec3ui [3]uint32

// Vec4ui is a four-component 32-bit unsigned int vector shader input type.
type Vec4ui [4]uint32

// Mat2 is a 32-bit floating 2x2 matrix shader input type.
type Mat2 [2][2]float32

// Mat3 is a 32-bit floating 3x3 matrix shader input type.
type Mat3 [3][3]float32

// Mat4 is a 32-bit floating 4x4 matrix shader input type.
type Mat4 [4][4]float32

// Mat2x3 is a 32-bit floating 2x3 matrix shader input type.
type Mat2x3 [2][3]float32

// Mat3x2 is a 32-bit floating 3x2 matrix shader input type.
type Mat3x2 [3][2]float32

// Mat2x4 is a 32-bit floating 2x4 matrix shader input type.
type Mat2x4 [2][4]float32

// Mat4x2 is a 32-bit floating 4x2 matrix shader input type.
type Mat4x2 [4][2]float32

// Mat3x4 is a 32-bit floating 3x4 matrix shader input type.
type Mat3x4 [3][4]float32

// Mat4x3 is a 32-bit floating 4x3 matrix shader input type.
type Mat4x3 [4][3]float32
