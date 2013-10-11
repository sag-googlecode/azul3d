// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// +build opengl_debug

// Package 'opengl' implements OpenGL version 1.2
package opengl

// #cgo windows LDFLAGS: -lopengl32
// #cgo linux LDFLAGS: -lGL -ldl
// #include "gl12.h"
import "C"

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

func boolToGL(b bool) C.GLboolean {
	if b {
		return C.GLboolean(1)
	}
	return C.GLboolean(0)
}

func parseVersions(s string) (n, major, minor, rev int) {
	var err error

	versions := strings.Split(s, ".")
	if len(versions) > 2 {
		versions = versions[0:3]
	}

	if len(versions) > 0 {
		major, err = strconv.Atoi(versions[0])
		if err != nil {
			return 0, 0, 0, 0
		}
	}

	if len(versions) > 1 {
		minor, err = strconv.Atoi(versions[1])
		if err != nil {
			return 0, 0, 0, 0
		}
	}
	n = len(versions)

	if len(versions) > 2 {
		_, err = fmt.Sscanf(versions[2]+"\n", "%d", &rev)
		if err != nil {
			n = 2
		}
	}

	return
}

func versionSupported(glc *Context) bool {
	ver := glc.GetString(VERSION)
	if len(ver) > 0 {
		n, wantedMajor, wantedMinor, wantedRev := parseVersions("1.2")
		if n < 2 {
			fmt.Printf("OpenGL: *** JSON version parsing failed for %q ***\n", "1.2")
			return false
		}

		n, major, minor, rev := parseVersions(ver)
		if n < 2 {
			fmt.Printf("OpenGL: *** Driver reported version parsing failed for %q ***\n", ver)
			return false
		}

		if major > wantedMajor {
			return true
		}
		if n == 2 {
			fmt.Printf("OpenGL: *** Driver reported version has no revision! %q ***\n", ver)
			if major >= wantedMajor && minor >= wantedMinor {
				return true
			}
		} else {
			if major >= wantedMajor && minor >= wantedMinor && rev >= wantedRev {
				return true
			}
		}
	}
	return false
}

func (glc *Context) queryExtensions() {
	// Initialize extensions map
	glc.extensions = make(map[string]bool)

	// Query extensions string
	extString := glc.GetString(EXTENSIONS)

	if len(extString) > 0 {
		for _, ext := range strings.Split(extString, " ") {
			if len(ext) > 0 {
				glc.extensions[ext] = true
			}
		}
	}
}

func (glc *Context) Panic(err string) {
	glc.access.Lock()
	defer glc.access.Unlock()

	fmt.Println("OpenGL call stack (last 500 - most recent first).")

	// Print stack now
	count := 0
	for i := len(glc.traceback); i > 0; i-- {
		count++
		fmt.Printf("%3.d. %s\n", count, glc.traceback[i-1])
	}

	panic(err)
}

func (glc *Context) trace(name string) {
	glc.access.Lock()

	glc.traceback = append(glc.traceback, name)
	l := len(glc.traceback)
	if l > 500 {
		glc.traceback = glc.traceback[l-500 : l]
	}

	if glc.inBeginEnd {
		glc.access.Unlock()
		return
	}
	err := glc.GetError()
	if err != NO_ERROR {
		glc.access.Unlock()

		switch err {
		case INVALID_ENUM:
			glc.Panic("GL_INVALID_ENUM: An unacceptable value was specified for an enumerated argument.")
		case INVALID_VALUE:
			glc.Panic("GL_INVALID_VALUE: A numeric argument is out of range.")
		case INVALID_OPERATION:
			glc.Panic("GL_INVALID_OPERATION: The specified operation is not allowed in the current state.")
		case INVALID_FRAMEBUFFER_OPERATION:
			glc.Panic("GL_INVALID_FRAMEBUFFER_OPERATION: The framebuffer object is not complete.")
		case OUT_OF_MEMORY:
			glc.Panic("GL_OUT_OF_MEMORY: There is not enough memory left to execute the command.")
		case STACK_UNDERFLOW:
			glc.Panic("GL_STACK_UNDERFLOW: An attempt has been made to perform an operation that would cause an internal stack to underflow.")
		case STACK_OVERFLOW:
			glc.Panic("GL_STACK_OVERFLOW: An attempt has been made to perform an operation that would cause an internal stack to overflow.")
		}
	} else {
		glc.access.Unlock()
	}
}

// Extension tells if the specified extension is supported by the OpenGL
// context.
//
// Extensions are stored internally as an map for performance, so lookups are
// very quick and require no OpenGL calls.
//
// Like other OpenGL functions, this is not thread safe.
//
// If this function always returns false, ensure that New() has been called in
// an active OpenGL context (as the extensions are queried then).
func (glc *Context) Extension(ext string) bool {
	_, ok := glc.extensions[ext]
	return ok
}

type Enum int

const (
	TEXTURE24_ARB                                              Enum = 0x84D8
	INVERSE_TRANSPOSE_NV                                       Enum = 0x862D
	STENCIL_TAG_BITS_EXT                                       Enum = 0x88F2
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              Enum = 0x8C80
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             Enum = 0x8354
	MATRIX2_NV                                                 Enum = 0x8632
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      Enum = 0x8973
	STENCIL_INDEX1                                             Enum = 0x8D46
	RGB16_SNORM                                                Enum = 0x8F9A
	INT64_VEC3_NV                                              Enum = 0x8FEA
	MAT_SPECULAR_BIT_PGI                                       Enum = 0x04000000
	VERSION_1_4                                                Enum = 1
	SGI_color_matrix                                           Enum = 1
	STENCIL_VALUE_MASK                                         Enum = 0x0B93
	SCISSOR_BOX                                                Enum = 0x0C10
	UNPACK_SKIP_IMAGES_EXT                                     Enum = 0x806D
	MAX_ACTIVE_LIGHTS_SGIX                                     Enum = 0x8405
	PROXY_TEXTURE_CUBE_MAP                                     Enum = 0x851B
	MAX_DRAW_BUFFERS                                           Enum = 0x8824
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       Enum = 0x93D5
	PACK_IMAGE_DEPTH_SGIS                                      Enum = 0x8131
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          Enum = 0x8DDE
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               Enum = 0x809E
	MAP2_VERTEX_ATTRIB5_4_NV                                   Enum = 0x8675
	INTENSITY_SNORM                                            Enum = 0x9013
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        Enum = 0x92CB
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      Enum = 0x93A0
	LINE_BIT                                                   Enum = 0x00000004
	COLOR_ATTACHMENT4_NV                                       Enum = 0x8CE4
	TRANSPOSE_AFFINE_3D_NV                                     Enum = 0x9098
	COMPRESSED_R11_EAC                                         Enum = 0x9270
	SOURCE1_RGB_ARB                                            Enum = 0x8581
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        Enum = 0x885D
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         Enum = 0x886A
	NUM_INSTRUCTIONS_PER_PASS_ATI                              Enum = 0x8971
	RENDERBUFFER_ALPHA_SIZE_EXT                                Enum = 0x8D53
	DUAL_TEXTURE_SELECT_SGIS                                   Enum = 0x8124
	MODELVIEW21_ARB                                            Enum = 0x8735
	HI_BIAS_NV                                                 Enum = 0x8714
	REG_11_ATI                                                 Enum = 0x892C
	RGB565                                                     Enum = 0x8D62
	SAMPLE_ALPHA_TO_MASK_SGIS                                  Enum = 0x809E
	QUAD_ALPHA8_SGIS                                           Enum = 0x811F
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             Enum = 0x812E
	PREVIOUS_EXT                                               Enum = 0x8578
	BLEND_EQUATION_ALPHA_EXT                                   Enum = 0x883D
	STENCIL_BACK_REF                                           Enum = 0x8CA3
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     Enum = 0x8CD1
	ISOLINES                                                   Enum = 0x8E7A
	DECR                                                       Enum = 0x1E03
	SCALEBIAS_HINT_SGIX                                        Enum = 0x8322
	MAP2_VERTEX_ATTRIB1_4_NV                                   Enum = 0x8671
	INTERLACE_READ_OML                                         Enum = 0x8981
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   Enum = 0x8CD2
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     Enum = 0x8DA9
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               Enum = 0x900B
	TRANSLATE_3D_NV                                            Enum = 0x9091
	EXT_blend_subtract                                         Enum = 1
	EMBOSS_MAP_NV                                              Enum = 0x855F
	DSDT8_NV                                                   Enum = 0x8709
	BUFFER_SIZE                                                Enum = 0x8764
	BUFFER_MAP_POINTER_OES                                     Enum = 0x88BD
	LIGHTING                                                   Enum = 0x0B50
	RGBA2_EXT                                                  Enum = 0x8055
	DISCARD_NV                                                 Enum = 0x8530
	FLOAT_VEC4                                                 Enum = 0x8B52
	INVALID_FRAMEBUFFER_OPERATION_EXT                          Enum = 0x0506
	LIGHT_MODEL_TWO_SIDE                                       Enum = 0x0B52
	LIGHT_ENV_MODE_SGIX                                        Enum = 0x8407
	UNSIGNED_SHORT_1_15_REV_MESA                               Enum = 0x8754
	OUTPUT_TEXTURE_COORD13_EXT                                 Enum = 0x87AA
	RELATIVE_MOVE_TO_NV                                        Enum = 0x03
	BLEND_DST_RGB                                              Enum = 0x80C8
	DEBUG_SOURCE_THIRD_PARTY_ARB                               Enum = 0x8249
	FRONT_LEFT                                                 Enum = 0x0400
	REPLICATE_BORDER_HP                                        Enum = 0x8153
	UNIFORM_BARRIER_BIT_EXT                                    Enum = 0x00000004
	PROXY_TEXTURE_3D                                           Enum = 0x8070
	MAX_SPOT_EXPONENT_NV                                       Enum = 0x8505
	PROGRAM_NATIVE_PARAMETERS_ARB                              Enum = 0x88AA
	TIMESTAMP                                                  Enum = 0x8E28
	CONSTANT_COLOR_EXT                                         Enum = 0x8001
	COMBINE                                                    Enum = 0x8570
	TRANSPOSE_CURRENT_MATRIX_ARB                               Enum = 0x88B7
	FLOAT_VEC4_ARB                                             Enum = 0x8B52
	SHADER_SOURCE_LENGTH                                       Enum = 0x8B88
	BLEND_DST_ALPHA                                            Enum = 0x80CA
	PROXY_TEXTURE_CUBE_MAP_EXT                                 Enum = 0x851B
	COLOR4_BIT_PGI                                             Enum = 0x00020000
	NORMAL_ARRAY_TYPE                                          Enum = 0x807E
	FILTER4_SGIS                                               Enum = 0x8146
	PIXEL_PACK_BUFFER_ARB                                      Enum = 0x88EB
	DEBUG_LOGGED_MESSAGES_ARB                                  Enum = 0x9145
	PASS_THROUGH_TOKEN                                         Enum = 0x0700
	QUERY_COUNTER_BITS_ARB                                     Enum = 0x8864
	FIXED_ONLY_ARB                                             Enum = 0x891D
	QUERY_BUFFER_AMD                                           Enum = 0x9192
	DETAIL_TEXTURE_2D_BINDING_SGIS                             Enum = 0x8096
	SOURCE3_RGB_NV                                             Enum = 0x8583
	MAP2_VERTEX_ATTRIB8_4_NV                                   Enum = 0x8678
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              Enum = 0x8F45
	IMAGE_3D                                                   Enum = 0x904E
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     Enum = 0x93DC
	MULTISAMPLE_BIT_ARB                                        Enum = 0x20000000
	FEEDBACK_BUFFER_POINTER                                    Enum = 0x0DF0
	MAX_SHININESS_NV                                           Enum = 0x8504
	MODELVIEW18_ARB                                            Enum = 0x8732
	SHADER_TYPE                                                Enum = 0x8B4F
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         Enum = 0x8C7F
	COLOR_ATTACHMENT13_NV                                      Enum = 0x8CED
	TEXTURE_FETCH_BARRIER_BIT                                  Enum = 0x00000008
	ALPHA                                                      Enum = 0x1906
	COMBINE_RGB_ARB                                            Enum = 0x8571
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               Enum = 0x87CA
	DEPENDENT_HILO_TEXTURE_2D_NV                               Enum = 0x8858
	REG_0_ATI                                                  Enum = 0x8921
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            Enum = 0x8B49
	RGBA16I                                                    Enum = 0x8D88
	RGB_SNORM                                                  Enum = 0x8F92
	TESSELLATION_MODE_AMD                                      Enum = 0x9004
	DRAW_BUFFER0                                               Enum = 0x8825
	VERSION_3_0                                                Enum = 1
	TRACE_OPERATIONS_BIT_MESA                                  Enum = 0x0001
	TEXTURE_INTERNAL_FORMAT_QCOM                               Enum = 0x8BD5
	ETC1_RGB8_OES                                              Enum = 0x8D64
	BINNING_CONTROL_HINT_QCOM                                  Enum = 0x8FB0
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         Enum = 0x9103
	SHADER_STORAGE_BLOCK                                       Enum = 0x92E6
	MAX_NUM_COMPATIBLE_SUBROUTINES                             Enum = 0x92F8
	PIXEL_MAP_A_TO_A                                           Enum = 0x0C79
	CAVEAT_SUPPORT                                             Enum = 0x82B8
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           Enum = 0x8909
	R11F_G11F_B10F_EXT                                         Enum = 0x8C3A
	SLUMINANCE_ALPHA_NV                                        Enum = 0x8C44
	RENDERBUFFER_SAMPLES                                       Enum = 0x8CAB
	MAX_VARYING_VECTORS                                        Enum = 0x8DFC
	ONE_MINUS_CONSTANT_ALPHA_EXT                               Enum = 0x8004
	INTERNALFORMAT_BLUE_TYPE                                   Enum = 0x827A
	COLOR_ATTACHMENT15_EXT                                     Enum = 0x8CEF
	UTF8_NV                                                    Enum = 0x909A
	SGIS_fog_function                                          Enum = 1
	COLOR_TABLE_ALPHA_SIZE                                     Enum = 0x80DD
	MATRIX_INDEX_ARRAY_OES                                     Enum = 0x8844
	STENCIL_WRITEMASK                                          Enum = 0x0B98
	TEXTURE5                                                   Enum = 0x84C5
	MAP2_VERTEX_ATTRIB4_4_NV                                   Enum = 0x8674
	GPU_ADDRESS_NV                                             Enum = 0x8F34
	POST_CONVOLUTION_RED_BIAS_EXT                              Enum = 0x8020
	DISTANCE_ATTENUATION_EXT                                   Enum = 0x8129
	SLIM12S_SGIX                                               Enum = 0x831F
	SOURCE2_RGB                                                Enum = 0x8582
	COMPRESSED_RED_RGTC1                                       Enum = 0x8DBB
	DOUBLE_MAT2_EXT                                            Enum = 0x8F46
	BLEND_EQUATION_OES                                         Enum = 0x8009
	SAMPLE_COVERAGE_VALUE                                      Enum = 0x80AA
	BOOL                                                       Enum = 0x8B56
	STENCIL_ATTACHMENT                                         Enum = 0x8D20
	MIN_LOD_WARNING_AMD                                        Enum = 0x919C
	UNSIGNED_SHORT_8_8_APPLE                                   Enum = 0x85BA
	COMPRESSED_RED_GREEN_RGTC2_EXT                             Enum = 0x8DBD
	VERTEX_ARRAY_OBJECT_AMD                                    Enum = 0x9154
	COLOR_TABLE_FORMAT_SGI                                     Enum = 0x80D8
	REPLACE_OLDEST_SUN                                         Enum = 0x0003
	PREVIOUS                                                   Enum = 0x8578
	REG_4_ATI                                                  Enum = 0x8925
	COMPRESSED_SIGNED_RG_RGTC2                                 Enum = 0x8DBE
	BUFFER_GPU_ADDRESS_NV                                      Enum = 0x8F1D
	LUMINANCE4_ALPHA4_EXT                                      Enum = 0x8043
	SIGNED_LUMINANCE8_ALPHA8_NV                                Enum = 0x8704
	UNIFORM_OFFSET                                             Enum = 0x8A3B
	GLYPH_WIDTH_BIT_NV                                         Enum = 0x01
	MATRIX_STRIDE                                              Enum = 0x92FF
	NORMAL_ARRAY                                               Enum = 0x8075
	MAX_ASYNC_READ_PIXELS_SGIX                                 Enum = 0x8361
	PROGRAM_RESIDENT_NV                                        Enum = 0x8647
	INDEX_ARRAY_LIST_IBM                                       Enum = 103073
	INDEX_SHIFT                                                Enum = 0x0D12
	ALPHA_BITS                                                 Enum = 0x0D55
	MAX_ELEMENTS_INDICES_EXT                                   Enum = 0x80E9
	R1UI_V3F_SUN                                               Enum = 0x85C4
	DEPTH_CLEAR_VALUE                                          Enum = 0x0B73
	CONVOLUTION_BORDER_MODE_EXT                                Enum = 0x8013
	DEPTH_STENCIL_ATTACHMENT                                   Enum = 0x821A
	TEXTURE_2D_STACK_MESAX                                     Enum = 0x875A
	COLOR_ATTACHMENT0_NV                                       Enum = 0x8CE0
	COPY_PIXEL_TOKEN                                           Enum = 0x0706
	COLOR_ARRAY_EXT                                            Enum = 0x8076
	OCCLUSION_TEST_HP                                          Enum = 0x8165
	MAP2_GRID_SEGMENTS                                         Enum = 0x0DD3
	DEBUG_SOURCE_THIRD_PARTY                                   Enum = 0x8249
	STENCIL_BACK_WRITEMASK                                     Enum = 0x8CA5
	GEOMETRY_OUTPUT_TYPE_EXT                                   Enum = 0x8DDC
	PHONG_HINT_WIN                                             Enum = 0x80EB
	BUFFER_OBJECT_APPLE                                        Enum = 0x85B3
	SHADING_LANGUAGE_VERSION_ARB                               Enum = 0x8B8C
	UNSIGNED_INT64_AMD                                         Enum = 0x8BC2
	TEXTURE_CLIPMAP_OFFSET_SGIX                                Enum = 0x8173
	COMPRESSED_RGBA                                            Enum = 0x84EE
	RGB_SCALE_EXT                                              Enum = 0x8573
	MAGNITUDE_SCALE_NV                                         Enum = 0x8712
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   Enum = 0x8C88
	MULT                                                       Enum = 0x0103
	MODELVIEW0_ARB                                             Enum = 0x1700
	INT_2_10_10_10_REV                                         Enum = 0x8D9F
	UNSIGNED_INT_ATOMIC_COUNTER                                Enum = 0x92DB
	VENDOR                                                     Enum = 0x1F00
	POST_COLOR_MATRIX_RED_SCALE                                Enum = 0x80B4
	BOOL_VEC2_ARB                                              Enum = 0x8B57
	DRAW_PIXEL_TOKEN                                           Enum = 0x0705
	COLOR_TABLE_INTENSITY_SIZE                                 Enum = 0x80DF
	CURRENT_RASTER_SECONDARY_COLOR                             Enum = 0x845F
	NORMAL_MAP_NV                                              Enum = 0x8511
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            Enum = 0x8515
	SOURCE2_RGB_ARB                                            Enum = 0x8582
	OPERAND0_ALPHA_EXT                                         Enum = 0x8598
	RGBA_FLOAT_MODE_ATI                                        Enum = 0x8820
	CON_30_ATI                                                 Enum = 0x895F
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              Enum = 0x9039
	STRICT_DEPTHFUNC_HINT_PGI                                  Enum = 0x1A216
	COLOR_ARRAY_TYPE                                           Enum = 0x8082
	MAX_DRAW_BUFFERS_ARB                                       Enum = 0x8824
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             Enum = 0x8B4C
	MALI_SHADER_BINARY_ARM                                     Enum = 0x8F60
	PIXEL_MAP_S_TO_S                                           Enum = 0x0C71
	ALL_COMPLETED_NV                                           Enum = 0x84F2
	MAP1_VERTEX_ATTRIB7_4_NV                                   Enum = 0x8667
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         Enum = 0x87F5
	INT_SAMPLER_1D_ARRAY_EXT                                   Enum = 0x8DCE
	COLOR_MATERIAL_FACE                                        Enum = 0x0B55
	COLOR_TABLE                                                Enum = 0x80D0
	TEXTURE16_ARB                                              Enum = 0x84D0
	FLOAT_RG32_NV                                              Enum = 0x8887
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    Enum = 0x898E
	COMPRESSED_RG_RGTC2                                        Enum = 0x8DBD
	RED_BIAS                                                   Enum = 0x0D15
	MAX_PROJECTION_STACK_DEPTH                                 Enum = 0x0D38
	DUAL_ALPHA4_SGIS                                           Enum = 0x8110
	TEXTURE3                                                   Enum = 0x84C3
	REFLECTION_MAP_EXT                                         Enum = 0x8512
	PROGRAM_FORMAT_ASCII_ARB                                   Enum = 0x8875
	LINEAR_MIPMAP_LINEAR                                       Enum = 0x2703
	INTERNALFORMAT_ALPHA_SIZE                                  Enum = 0x8274
	ATOMIC_COUNTER_BUFFER_SIZE                                 Enum = 0x92C3
	EQUIV                                                      Enum = 0x1509
	DRAW_BUFFER1_ATI                                           Enum = 0x8826
	CLIP_PLANE1                                                Enum = 0x3001
	REDUCE_EXT                                                 Enum = 0x8016
	LUMINANCE16_ALPHA16_EXT                                    Enum = 0x8048
	PROGRAM_PIPELINE                                           Enum = 0x82E4
	MODELVIEW1_EXT                                             Enum = 0x850A
	POST_CONVOLUTION_GREEN_SCALE_EXT                           Enum = 0x801D
	TEXTURE_GREEN_SIZE_EXT                                     Enum = 0x805D
	BLEND_SRC_ALPHA_OES                                        Enum = 0x80CB
	LIGHT_MODEL_COLOR_CONTROL                                  Enum = 0x81F8
	CND0_ATI                                                   Enum = 0x896B
	RED_INTEGER                                                Enum = 0x8D94
	DOUBLE                                                     Enum = 0x140A
	HISTOGRAM_GREEN_SIZE_EXT                                   Enum = 0x8029
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              Enum = 0x83F3
	OBJECT_DELETE_STATUS_ARB                                   Enum = 0x8B80
	TRANSFORM_FEEDBACK_VARYING                                 Enum = 0x92F4
	LINES_ADJACENCY_ARB                                        Enum = 0x000A
	STACK_UNDERFLOW                                            Enum = 0x0504
	DETAIL_TEXTURE_MODE_SGIS                                   Enum = 0x809B
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           Enum = 0x8E23
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           Enum = 0x8F27
	DEBUG_LOGGED_MESSAGES                                      Enum = 0x9145
	SAMPLE_ALPHA_TO_COVERAGE                                   Enum = 0x809E
	STATIC_VERTEX_ARRAY_IBM                                    Enum = 103061
	CONVOLUTION_BORDER_COLOR                                   Enum = 0x8154
	REPLACEMENT_CODE_SUN                                       Enum = 0x81D8
	PROGRAM_INSTRUCTIONS_ARB                                   Enum = 0x88A0
	AUX2                                                       Enum = 0x040B
	SOURCE1_ALPHA_ARB                                          Enum = 0x8589
	MAP2_VERTEX_ATTRIB6_4_NV                                   Enum = 0x8676
	SAMPLER_2D_ARRAY_SHADOW                                    Enum = 0x8DC4
	GEOMETRY_INPUT_TYPE_ARB                                    Enum = 0x8DDB
	SHADER_BINARY_VIV                                          Enum = 0x8FC4
	SHADER_IMAGE_LOAD                                          Enum = 0x82A4
	VERTEX_BINDING_OFFSET                                      Enum = 0x82D7
	OP_MULTIPLY_MATRIX_EXT                                     Enum = 0x8798
	WRITE_ONLY_ARB                                             Enum = 0x88B9
	BOLD_BIT_NV                                                Enum = 0x01
	GLOBAL_ALPHA_SUN                                           Enum = 0x81D9
	UNSIGNED_INT_VEC2                                          Enum = 0x8DC6
	INT_SAMPLER_CUBE                                           Enum = 0x8DCC
	TEXTURE_3D                                                 Enum = 0x806F
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           Enum = 0x884E
	READ_PIXEL_DATA_RANGE_POINTER_NV                           Enum = 0x887D
	RENDERBUFFER_RED_SIZE_EXT                                  Enum = 0x8D50
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               Enum = 0x8DBA
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       Enum = 0x8CD4
	POINT_SIZE_MAX_ARB                                         Enum = 0x8127
	TEXTURE31_ARB                                              Enum = 0x84DF
	MAX_TESS_GEN_LEVEL                                         Enum = 0x8E7E
	VIEW_CLASS_BPTC_FLOAT                                      Enum = 0x82D3
	IMAGE_BINDING_NAME_EXT                                     Enum = 0x8F3A
	TEXTURE_ENV                                                Enum = 0x2300
	SAMPLES                                                    Enum = 0x80A9
	OUTPUT_TEXTURE_COORD27_EXT                                 Enum = 0x87B8
	TEXTURE_GREEN_TYPE_ARB                                     Enum = 0x8C11
	SAMPLES_SGIS                                               Enum = 0x80A9
	GLOBAL_ALPHA_FACTOR_SUN                                    Enum = 0x81DA
	INTERNALFORMAT_ALPHA_TYPE                                  Enum = 0x827B
	ALPHA_TEST_REF                                             Enum = 0x0BC2
	MAP2_VERTEX_ATTRIB3_4_NV                                   Enum = 0x8673
	MODELVIEW1_ARB                                             Enum = 0x850A
	MODELVIEW27_ARB                                            Enum = 0x873B
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            Enum = 0x8A09
	BOOL_VEC2                                                  Enum = 0x8B57
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            Enum = 0x90C7
	RIGHT                                                      Enum = 0x0407
	PACK_ALIGNMENT                                             Enum = 0x0D05
	VERTEX_ARRAY                                               Enum = 0x8074
	TEXTURE_MIN_FILTER                                         Enum = 0x2801
	IMAGE_SCALE_X_HP                                           Enum = 0x8155
	FRAMEBUFFER_DEFAULT                                        Enum = 0x8218
	SLIM10U_SGIX                                               Enum = 0x831E
	RGB_SCALE_ARB                                              Enum = 0x8573
	AUX_DEPTH_STENCIL_APPLE                                    Enum = 0x8A14
	MEDIUM_FLOAT                                               Enum = 0x8DF1
	POST_COLOR_MATRIX_BLUE_SCALE                               Enum = 0x80B6
	MAX_COMPUTE_ATOMIC_COUNTERS                                Enum = 0x8265
	OP_MUL_EXT                                                 Enum = 0x8786
	LUMINANCE                                                  Enum = 0x1909
	FRAGMENT_COLOR_EXT                                         Enum = 0x834C
	VERTEX_ATTRIB_ARRAY8_NV                                    Enum = 0x8658
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       Enum = 0x9105
	VERTEX_ARRAY_TYPE_EXT                                      Enum = 0x807B
	PIXEL_GROUP_COLOR_SGIS                                     Enum = 0x8356
	MATRIX30_ARB                                               Enum = 0x88DE
	ACCUM                                                      Enum = 0x0100
	MATRIX1_ARB                                                Enum = 0x88C1
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  Enum = 0x8C76
	KEEP                                                       Enum = 0x1E00
	SEPARABLE_2D                                               Enum = 0x8012
	TEXTURE_1D_STACK_MESAX                                     Enum = 0x8759
	INT_VEC4_ARB                                               Enum = 0x8B55
	LUMINANCE_SNORM                                            Enum = 0x9011
	UNSIGNED_BYTE_2_3_3_REV_EXT                                Enum = 0x8362
	TEXTURE_RESIDENT_EXT                                       Enum = 0x8067
	RELEASED_APPLE                                             Enum = 0x8A19
	MAX_SHADER_BUFFER_ADDRESS_NV                               Enum = 0x8F35
	WAIT_FAILED_APPLE                                          Enum = 0x911D
	COLOR_ARRAY_LIST_IBM                                       Enum = 103072
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   Enum = 0x8DA0
	IMAGE_BINDING_NAME                                         Enum = 0x8F3A
	SGIX_fog_offset                                            Enum = 1
	DEPTH_FUNC                                                 Enum = 0x0B74
	STENCIL_FUNC                                               Enum = 0x0B92
	MAX_EXT                                                    Enum = 0x8008
	CLIP_VOLUME_CLIPPING_HINT_EXT                              Enum = 0x80F0
	RENDERBUFFER_BINDING_EXT                                   Enum = 0x8CA7
	SGIS_point_line_texgen                                     Enum = 1
	GL_4_BYTES                                                 Enum = 0x1409
	FRAGMENT_DEPTH                                             Enum = 0x8452
	MATRIX0_NV                                                 Enum = 0x8630
	SIGNED_HILO8_NV                                            Enum = 0x885F
	STENCIL_CLEAR_TAG_VALUE_EXT                                Enum = 0x88F3
	RGB10_A2UI                                                 Enum = 0x906F
	GLYPH_HAS_KERNING_BIT_NV                                   Enum = 0x100
	DEPTH_COMPONENT32                                          Enum = 0x81A7
	DRAW_BUFFER4_ARB                                           Enum = 0x8829
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     Enum = 0x8852
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             Enum = 0x8A04
	COLOR_ATTACHMENT6_EXT                                      Enum = 0x8CE6
	PATCH_DEFAULT_INNER_LEVEL                                  Enum = 0x8E73
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        Enum = 0x00000001
	SIGNED_ALPHA_NV                                            Enum = 0x8705
	PATH_STROKE_BOUNDING_BOX_NV                                Enum = 0x90A2
	PROXY_TEXTURE_1D                                           Enum = 0x8063
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      Enum = 0x8E86
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              Enum = 0x9144
	NORMAL_ARRAY_STRIDE                                        Enum = 0x807F
	RESAMPLE_DECIMATE_SGIX                                     Enum = 0x8430
	TEXTURE12                                                  Enum = 0x84CC
	OPERAND0_RGB_EXT                                           Enum = 0x8590
	RED_BIT_ATI                                                Enum = 0x00000001
	BOOL_VEC3                                                  Enum = 0x8B58
	COMPRESSED_SRGB_ALPHA                                      Enum = 0x8C49
	POLYGON                                                    Enum = 0x0009
	LUMINANCE_FLOAT16_ATI                                      Enum = 0x881E
	CURRENT_PALETTE_MATRIX_ARB                                 Enum = 0x8843
	TEXCOORD2_BIT_PGI                                          Enum = 0x20000000
	CMYK_EXT                                                   Enum = 0x800C
	VARIABLE_B_NV                                              Enum = 0x8524
	CON_17_ATI                                                 Enum = 0x8952
	SRGB8_NV                                                   Enum = 0x8C41
	RASTERIZER_DISCARD_EXT                                     Enum = 0x8C89
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              Enum = 0x8F39
	PATH_FORMAT_SVG_NV                                         Enum = 0x9070
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              Enum = 0x00000001
	VERTEX_DATA_HINT_PGI                                       Enum = 0x1A22A
	V3F                                                        Enum = 0x2A21
	REFERENCE_PLANE_SGIX                                       Enum = 0x817D
	EYE_PLANE_ABSOLUTE_NV                                      Enum = 0x855C
	SGIS_multisample                                           Enum = 1
	OUTPUT_TEXTURE_COORD4_EXT                                  Enum = 0x87A1
	NOTEQUAL                                                   Enum = 0x0205
	MAX_GEOMETRY_UNIFORM_BLOCKS                                Enum = 0x8A2C
	SLUMINANCE8_ALPHA8                                         Enum = 0x8C45
	TEXTURE_SWIZZLE_B_EXT                                      Enum = 0x8E44
	PATH_DASH_ARRAY_COUNT_NV                                   Enum = 0x909F
	TEXTURE_WIDTH                                              Enum = 0x1000
	TEXTURE_WRAP_T                                             Enum = 0x2803
	TEXTURE5_ARB                                               Enum = 0x84C5
	LUMINANCE_ALPHA32I_EXT                                     Enum = 0x8D87
	STENCIL_EXT                                                Enum = 0x1802
	SYNC_X11_FENCE_EXT                                         Enum = 0x90E1
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               Enum = 0x92C5
	TESS_EVALUATION_SUBROUTINE                                 Enum = 0x92EA
	FUNC_REVERSE_SUBTRACT_EXT                                  Enum = 0x800B
	CONVOLUTION_1D_EXT                                         Enum = 0x8010
	PROGRAM_PIPELINE_BINDING_EXT                               Enum = 0x825A
	SLUMINANCE_EXT                                             Enum = 0x8C46
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    Enum = 0x8C8B
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         Enum = 0x8DA2
	INT_IMAGE_CUBE_EXT                                         Enum = 0x905B
	VERTEX4_BIT_PGI                                            Enum = 0x00000008
	SGIX_vertex_preclip                                        Enum = 1
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        Enum = 0x8217
	FOG_COORDINATE_ARRAY_POINTER                               Enum = 0x8456
	UNSIGNED_INT_24_8_NV                                       Enum = 0x84FA
	COMBINER6_NV                                               Enum = 0x8556
	FRAMEBUFFER                                                Enum = 0x8D40
	MAX_SAMPLE_MASK_WORDS_NV                                   Enum = 0x8E59
	PATH_JOIN_STYLE_NV                                         Enum = 0x9079
	CURRENT_NORMAL                                             Enum = 0x0B02
	CLIENT_ATTRIB_STACK_DEPTH                                  Enum = 0x0BB1
	DEPTH_SCALE                                                Enum = 0x0D1E
	PROXY_TEXTURE_CUBE_MAP_ARB                                 Enum = 0x851B
	SWIZZLE_STQ_DQ_ATI                                         Enum = 0x8979
	HORIZONTAL_LINE_TO_NV                                      Enum = 0x06
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         Enum = 0x85C3
	QUERY_BY_REGION_NO_WAIT_NV                                 Enum = 0x8E16
	FRAME_NV                                                   Enum = 0x8E26
	UNSIGNED_BYTE                                              Enum = 0x1401
	CLAMP                                                      Enum = 0x2900
	MAX_VERTEX_ATTRIB_BINDINGS                                 Enum = 0x82DA
	GENERIC_ATTRIB_NV                                          Enum = 0x8C7D
	MOVE_TO_CONTINUES_NV                                       Enum = 0x90B6
	FRAGMENT_SUBROUTINE_UNIFORM                                Enum = 0x92F2
	SOURCE2_ALPHA_EXT                                          Enum = 0x858A
	PALETTE8_RGBA4_OES                                         Enum = 0x8B98
	OR                                                         Enum = 0x1507
	REFLECTION_MAP_NV                                          Enum = 0x8512
	WEIGHT_ARRAY_POINTER_ARB                                   Enum = 0x86AC
	PERFMON_GLOBAL_MODE_QCOM                                   Enum = 0x8FA0
	BLEND_DST                                                  Enum = 0x0BE0
	MAX_CLIPMAP_DEPTH_SGIX                                     Enum = 0x8177
	MODELVIEW15_ARB                                            Enum = 0x872F
	SAMPLE_BUFFERS_3DFX                                        Enum = 0x86B3
	SIGNED_RGBA8_NV                                            Enum = 0x86FC
	IMAGE_3D_EXT                                               Enum = 0x904E
	C3F_V3F                                                    Enum = 0x2A24
	CLIP_PLANE0                                                Enum = 0x3000
	SAMPLE_COVERAGE_VALUE_ARB                                  Enum = 0x80AA
	R16F_EXT                                                   Enum = 0x822D
	MAX_LABEL_LENGTH                                           Enum = 0x82E8
	SIGNED_ALPHA8_NV                                           Enum = 0x8706
	ALPHA_FLOAT32_ATI                                          Enum = 0x8816
	ACCUM_ADJACENT_PAIRS_NV                                    Enum = 0x90AD
	POLYGON_OFFSET_UNITS                                       Enum = 0x2A00
	LUMINANCE16_EXT                                            Enum = 0x8042
	CONVOLUTION_BORDER_COLOR_HP                                Enum = 0x8154
	CON_29_ATI                                                 Enum = 0x895E
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   Enum = 0x8E4C
	MAP_READ_BIT                                               Enum = 0x0001
	STEREO                                                     Enum = 0x0C33
	IMAGE_ROTATE_ORIGIN_X_HP                                   Enum = 0x815A
	ADD_SIGNED_ARB                                             Enum = 0x8574
	NUM_INSTRUCTIONS_TOTAL_ATI                                 Enum = 0x8972
	FRAGMENT_PROGRAM_POSITION_MESA                             Enum = 0x8BB0
	COLOR_ATTACHMENT2_NV                                       Enum = 0x8CE2
	LAST_VIDEO_CAPTURE_STATUS_NV                               Enum = 0x9027
	VERTEX_ARRAY_LIST_IBM                                      Enum = 103070
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         Enum = 0x00000800
	MULTISAMPLE_SGIS                                           Enum = 0x809D
	MAX_VERTEX_SHADER_LOCALS_EXT                               Enum = 0x87C9
	W_EXT                                                      Enum = 0x87D8
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           Enum = 0x8C03
	INT_SAMPLER_2D_RECT                                        Enum = 0x8DCD
	AMBIENT                                                    Enum = 0x1200
	INTERNALFORMAT_SUPPORTED                                   Enum = 0x826F
	VERTEX_ATTRIB_BINDING                                      Enum = 0x82D4
	DRAW_BUFFER3_ATI                                           Enum = 0x8828
	TEXTURE_COLOR_TABLE_SGI                                    Enum = 0x80BC
	RGBA32F_ARB                                                Enum = 0x8814
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                Enum = 0x8857
	RENDERBUFFER_DEPTH_SIZE                                    Enum = 0x8D54
	PIXEL_MODE_BIT                                             Enum = 0x00000020
	PROXY_TEXTURE_2D_EXT                                       Enum = 0x8064
	DECR_WRAP                                                  Enum = 0x8508
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            Enum = 0x8517
	HALF_BIAS_NORMAL_NV                                        Enum = 0x853A
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            Enum = 0x88EF
	COLOR                                                      Enum = 0x1800
	COLOR_TABLE_GREEN_SIZE_SGI                                 Enum = 0x80DB
	BGRA_EXT                                                   Enum = 0x80E1
	GENERATE_MIPMAP                                            Enum = 0x8191
	OP_MIN_EXT                                                 Enum = 0x878B
	MAP1_VERTEX_ATTRIB4_4_NV                                   Enum = 0x8664
	EVAL_TRIANGULAR_2D_NV                                      Enum = 0x86C1
	ANY_SAMPLES_PASSED_EXT                                     Enum = 0x8C2F
	SAMPLER_CUBE_SHADOW                                        Enum = 0x8DC5
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         Enum = 0x8E8A
	ROUND_NV                                                   Enum = 0x90A4
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             Enum = 0x9128
	CUBIC_HP                                                   Enum = 0x815F
	MAX_VERTEX_UNITS_OES                                       Enum = 0x86A4
	STENCIL_BACK_PASS_DEPTH_PASS                               Enum = 0x8803
	ANY_SAMPLES_PASSED                                         Enum = 0x8C2F
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      Enum = 0x8C71
	PROXY_HISTOGRAM                                            Enum = 0x8025
	POINT_SIZE_MAX                                             Enum = 0x8127
	DOT3_RGB_ARB                                               Enum = 0x86AE
	MAX_MAP_TESSELLATION_NV                                    Enum = 0x86D6
	BLEND_DST_RGB_OES                                          Enum = 0x80C8
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          Enum = 0x840A
	ZERO_EXT                                                   Enum = 0x87DD
	RGB16F_EXT                                                 Enum = 0x881B
	POST_CONVOLUTION_COLOR_TABLE_SGI                           Enum = 0x80D1
	MODELVIEW28_ARB                                            Enum = 0x873C
	TRANSFORM_FEEDBACK_BUFFER_START                            Enum = 0x8C84
	INTERLEAVED_ATTRIBS_NV                                     Enum = 0x8C8C
	VERTEX_SUBROUTINE                                          Enum = 0x92E8
	TEXTURE_IMAGE_FORMAT                                       Enum = 0x828F
	SLUMINANCE                                                 Enum = 0x8C46
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            Enum = 0x00000020
	COLOR_ATTACHMENT11                                         Enum = 0x8CEB
	LINEAR_DETAIL_SGIS                                         Enum = 0x8097
	FOG_COORD_ARRAY_STRIDE                                     Enum = 0x8455
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        Enum = 0x8E7F
	SPRITE_AXIS_SGIX                                           Enum = 0x814A
	VERTEX_ATTRIB_ARRAY_LONG                                   Enum = 0x874E
	GEOMETRY_VERTICES_OUT                                      Enum = 0x8916
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              Enum = 0x8CD4
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              Enum = 0x9111
	CONSTANT_EXT                                               Enum = 0x8576
	WEIGHT_ARRAY_STRIDE_OES                                    Enum = 0x86AA
	QUERY_RESULT_AVAILABLE_ARB                                 Enum = 0x8867
	INDEX_ARRAY_BUFFER_BINDING                                 Enum = 0x8899
	TEXTURE_SWIZZLE_R                                          Enum = 0x8E42
	DIFFUSE                                                    Enum = 0x1201
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             Enum = 0x845A
	EDGEFLAG_BIT_PGI                                           Enum = 0x00040000
	PIXEL_TILE_GRID_WIDTH_SGIX                                 Enum = 0x8142
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               Enum = 0x81A8
	TEXTURE_RANGE_POINTER_APPLE                                Enum = 0x85B8
	LUMINANCE_ALPHA_FLOAT32_ATI                                Enum = 0x8819
	POINT_SPRITE_R_MODE_NV                                     Enum = 0x8863
	PIXEL_PACK_BUFFER_BINDING_ARB                              Enum = 0x88ED
	REG_28_ATI                                                 Enum = 0x893D
	LUMINANCE32I_EXT                                           Enum = 0x8D86
	DATA_BUFFER_AMD                                            Enum = 0x9151
	PIXEL_MAP_G_TO_G_SIZE                                      Enum = 0x0CB7
	PIXEL_TILE_HEIGHT_SGIX                                     Enum = 0x8141
	TRACK_MATRIX_TRANSFORM_NV                                  Enum = 0x8649
	UNSIGNED_INT_24_8_MESA                                     Enum = 0x8751
	OUTPUT_TEXTURE_COORD6_EXT                                  Enum = 0x87A3
	INVARIANT_VALUE_EXT                                        Enum = 0x87EA
	REG_7_ATI                                                  Enum = 0x8928
	GEOMETRY_SHADER                                            Enum = 0x8DD9
	ACCUM_BUFFER_BIT                                           Enum = 0x00000200
	UNSIGNED_INT                                               Enum = 0x1405
	INTENSITY_FLOAT16_APPLE                                    Enum = 0x881D
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       Enum = 0x93D4
	LIGHT_MODEL_LOCAL_VIEWER                                   Enum = 0x0B51
	DUAL_LUMINANCE8_SGIS                                       Enum = 0x8115
	CURRENT_BINORMAL_EXT                                       Enum = 0x843C
	RED_MIN_CLAMP_INGR                                         Enum = 0x8560
	WEIGHT_ARRAY_SIZE_OES                                      Enum = 0x86AB
	MAX_PALETTE_MATRICES_OES                                   Enum = 0x8842
	VERTEX_PROGRAM_CALLBACK_MESA                               Enum = 0x8BB4
	MAX_SUBROUTINES                                            Enum = 0x8DE7
	QUERY_NO_WAIT_NV                                           Enum = 0x8E14
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           Enum = 0x9137
	UNPACK_SKIP_IMAGES                                         Enum = 0x806D
	TEXTURE_MAX_LEVEL_SGIS                                     Enum = 0x813D
	OP_FRAC_EXT                                                Enum = 0x8789
	OBJECT_SUBTYPE_ARB                                         Enum = 0x8B4F
	FRAMEBUFFER_OES                                            Enum = 0x8D40
	LINEAR_SHARPEN_COLOR_SGIS                                  Enum = 0x80AF
	TEXTURE14_ARB                                              Enum = 0x84CE
	VERTEX_ATTRIB_ARRAY1_NV                                    Enum = 0x8651
	CULL_FRAGMENT_NV                                           Enum = 0x86E7
	FRAGMENT_COLOR_MATERIAL_SGIX                               Enum = 0x8401
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             Enum = 0x9065
	RG8                                                        Enum = 0x822B
	PIXEL_COUNT_NV                                             Enum = 0x8866
	NEGATE_BIT_ATI                                             Enum = 0x00000004
	COMPRESSED_SRGB_S3TC_DXT1_NV                               Enum = 0x8C4C
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  Enum = 0x9277
	NUM_SAMPLE_COUNTS                                          Enum = 0x9380
	SECONDARY_COLOR_ARRAY_TYPE                                 Enum = 0x845B
	INT_VEC2_ARB                                               Enum = 0x8B53
	COLOR_ATTACHMENT8                                          Enum = 0x8CE8
	STENCIL_INDEX1_OES                                         Enum = 0x8D46
	IMAGE_BINDING_LAYERED                                      Enum = 0x8F3C
	BUFFER_ACCESS_FLAGS                                        Enum = 0x911F
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 Enum = 0x9314
	CURRENT_RASTER_DISTANCE                                    Enum = 0x0B09
	INDEX_BITS                                                 Enum = 0x0D51
	OUTPUT_TEXTURE_COORD22_EXT                                 Enum = 0x87B3
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   Enum = 0x8C74
	IMAGE_BINDING_FORMAT                                       Enum = 0x906E
	ADJACENT_PAIRS_NV                                          Enum = 0x90AE
	COMBINER_INPUT_NV                                          Enum = 0x8542
	SOURCE0_RGB_EXT                                            Enum = 0x8580
	RENDERBUFFER_WIDTH                                         Enum = 0x8D42
	IMAGE_2D_ARRAY_EXT                                         Enum = 0x9053
	MAP_UNSYNCHRONIZED_BIT                                     Enum = 0x0020
	C4F_N3F_V3F                                                Enum = 0x2A26
	MIRRORED_REPEAT_ARB                                        Enum = 0x8370
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       Enum = 0x83F5
	FOG_COORDINATE_SOURCE                                      Enum = 0x8450
	RGBA_FLOAT_MODE_ARB                                        Enum = 0x8820
	READ_WRITE                                                 Enum = 0x88BA
	PERFMON_RESULT_AMD                                         Enum = 0x8BC6
	COLOR_ATTACHMENT5                                          Enum = 0x8CE5
	PROXY_TEXTURE_2D_MULTISAMPLE                               Enum = 0x9101
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           Enum = 0x914A
	TEXTURE_BINDING_3D                                         Enum = 0x806A
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 Enum = 0x808C
	TEXTURE_MAX_CLAMP_T_SGIX                                   Enum = 0x836A
	VERTEX_ATTRIB_ARRAY2_NV                                    Enum = 0x8652
	LUMINANCE_ALPHA8I_EXT                                      Enum = 0x8D93
	ALPHA8_EXT                                                 Enum = 0x803C
	UNKNOWN_CONTEXT_RESET_ARB                                  Enum = 0x8255
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              Enum = 0x8848
	REG_12_ATI                                                 Enum = 0x892D
	COLOR_MATERIAL_PARAMETER                                   Enum = 0x0B56
	MAX_SHADER_STORAGE_BLOCK_SIZE                              Enum = 0x90DE
	TESS_CONTROL_SHADER_BIT                                    Enum = 0x00000008
	TEXTURE_TOO_LARGE_EXT                                      Enum = 0x8065
	MODELVIEW22_ARB                                            Enum = 0x8736
	PN_TRIANGLES_NORMAL_MODE_ATI                               Enum = 0x87F3
	COLOR_ATTACHMENT_EXT                                       Enum = 0x90F0
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       Enum = 0x93D7
	TEXTURE10                                                  Enum = 0x84CA
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 Enum = 0x8C8B
	PATH_TERMINAL_DASH_CAP_NV                                  Enum = 0x907D
	VERTEX_ATTRIB_ARRAY7_NV                                    Enum = 0x8657
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     Enum = 0x87F8
	STATE_RESTORE                                              Enum = 0x8BDC
	PROXY_TEXTURE_COLOR_TABLE_SGI                              Enum = 0x80BD
	OUTPUT_TEXTURE_COORD5_EXT                                  Enum = 0x87A2
	POST_CONVOLUTION_RED_BIAS                                  Enum = 0x8020
	INDEX_TEST_REF_EXT                                         Enum = 0x81B7
	CURRENT_FOG_COORD                                          Enum = 0x8453
	OUTPUT_TEXTURE_COORD15_EXT                                 Enum = 0x87AC
	ELEMENT_ARRAY_BUFFER_ARB                                   Enum = 0x8893
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              Enum = 0x8C4C
	STENCIL_ATTACHMENT_OES                                     Enum = 0x8D20
	COMPATIBLE_SUBROUTINES                                     Enum = 0x8E4B
	RGBA8_OES                                                  Enum = 0x8058
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          Enum = 0x88FE
	READ_FRAMEBUFFER_ANGLE                                     Enum = 0x8CA8
	FIELDS_NV                                                  Enum = 0x8E27
	PATH_FILL_MASK_NV                                          Enum = 0x9081
	NATIVE_GRAPHICS_HANDLE_PGI                                 Enum = 0x1A202
	POLYGON_STIPPLE                                            Enum = 0x0B42
	MAX_PIXEL_MAP_TABLE                                        Enum = 0x0D34
	SAMPLE_COVERAGE_INVERT                                     Enum = 0x80AB
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      Enum = 0x8210
	MAP2_VERTEX_ATTRIB14_4_NV                                  Enum = 0x867E
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     Enum = 0x87CE
	COMPRESSED_RGBA_ASTC_10x6_KHR                              Enum = 0x93B9
	ALL_SHADER_BITS_EXT                                        Enum = 0xFFFFFFFF
	DRAW_BUFFER14_NV                                           Enum = 0x8833
	DOUBLE_VEC3_EXT                                            Enum = 0x8FFD
	RG32I                                                      Enum = 0x823B
	FRAGMENT_LIGHT3_SGIX                                       Enum = 0x840F
	COLOR_SUM                                                  Enum = 0x8458
	DISCARD_ATI                                                Enum = 0x8763
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         Enum = 0x8CD1
	PATH_STROKE_COVER_MODE_NV                                  Enum = 0x9083
	SRC_ALPHA                                                  Enum = 0x0302
	SPECULAR                                                   Enum = 0x1202
	R32UI                                                      Enum = 0x8236
	FRAGMENT_LIGHT2_SGIX                                       Enum = 0x840E
	NORMAL_MAP_OES                                             Enum = 0x8511
	PACK_INVERT_MESA                                           Enum = 0x8758
	VIDEO_BUFFER_NV                                            Enum = 0x9020
	GEOMETRY_SUBROUTINE                                        Enum = 0x92EB
	EXT_rescale_normal                                         Enum = 1
	BLEND_EQUATION_EXT                                         Enum = 0x8009
	IMAGE_1D                                                   Enum = 0x904C
	TEXTURE_BINDING_RECTANGLE                                  Enum = 0x84F6
	SCALAR_EXT                                                 Enum = 0x87BE
	TEXTURE_COMPARE_FUNC_EXT                                   Enum = 0x884D
	FRAMEBUFFER_COMPLETE_OES                                   Enum = 0x8CD5
	DRAW_BUFFER5                                               Enum = 0x882A
	RGBA32I                                                    Enum = 0x8D82
	INT8_NV                                                    Enum = 0x8FE0
	CONDITION_SATISFIED                                        Enum = 0x911C
	LIST_BIT                                                   Enum = 0x00020000
	VERTEX_SHADER_BIT                                          Enum = 0x00000001
	READ_PIXELS_TYPE                                           Enum = 0x828E
	DISPATCH_INDIRECT_BUFFER_BINDING                           Enum = 0x90EF
	COMMAND_BARRIER_BIT_EXT                                    Enum = 0x00000040
	BLUE_BITS                                                  Enum = 0x0D54
	MODELVIEW24_ARB                                            Enum = 0x8738
	DEPENDENT_GB_TEXTURE_2D_NV                                 Enum = 0x86EA
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    Enum = 0x880F
	COLOR_ATTACHMENT7_NV                                       Enum = 0x8CE7
	COLOR_MATRIX_STACK_DEPTH                                   Enum = 0x80B2
	SEPARATE_SPECULAR_COLOR_EXT                                Enum = 0x81FA
	TRANSPOSE_PROJECTION_MATRIX_ARB                            Enum = 0x84E4
	TEXTURE_FILTER_CONTROL_EXT                                 Enum = 0x8500
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            Enum = 0x8517
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              Enum = 0x8CD4
	TESS_GEN_MODE                                              Enum = 0x8E76
	ELEMENT_ARRAY_LENGTH_NV                                    Enum = 0x8F33
	PATH_FOG_GEN_MODE_NV                                       Enum = 0x90AC
	QUAD_INTENSITY8_SGIS                                       Enum = 0x8123
	VERTEX_PRECLIP_SGIX                                        Enum = 0x83EE
	OBJECT_COMPILE_STATUS_ARB                                  Enum = 0x8B81
	FRAMEBUFFER_BINDING_EXT                                    Enum = 0x8CA6
	STRICT_SCISSOR_HINT_PGI                                    Enum = 0x1A218
	ONE_MINUS_SRC_COLOR                                        Enum = 0x0301
	EVAL_VERTEX_ATTRIB7_NV                                     Enum = 0x86CD
	VERTEX_ARRAY_BUFFER_BINDING                                Enum = 0x8896
	REG_20_ATI                                                 Enum = 0x8935
	TRANSFORM_FEEDBACK_BUFFER                                  Enum = 0x8C8E
	INDEX_ARRAY_ADDRESS_NV                                     Enum = 0x8F24
	FONT_UNDERLINE_POSITION_BIT_NV                             Enum = 0x04000000
	COMPUTE_SUBROUTINE_UNIFORM                                 Enum = 0x92F3
	QUERY_WAIT                                                 Enum = 0x8E13
	LAST_VERTEX_CONVENTION_EXT                                 Enum = 0x8E4E
	TEXTURE_SAMPLES_IMG                                        Enum = 0x9136
	CLIP_DISTANCE1                                             Enum = 0x3001
	NUM_PROGRAM_BINARY_FORMATS                                 Enum = 0x87FE
	DEPTH_TEXTURE_MODE_ARB                                     Enum = 0x884B
	MAP1_INDEX                                                 Enum = 0x0D91
	TEXTURE2                                                   Enum = 0x84C2
	SIGNED_LUMINANCE_ALPHA_NV                                  Enum = 0x8703
	INDEX_ARRAY_POINTER_EXT                                    Enum = 0x8091
	RED_MAX_CLAMP_INGR                                         Enum = 0x8564
	ACTIVE_VERTEX_UNITS_ARB                                    Enum = 0x86A5
	OUTPUT_TEXTURE_COORD18_EXT                                 Enum = 0x87AF
	COMPRESSED_LUMINANCE_ARB                                   Enum = 0x84EA
	FLOAT_MAT3x2                                               Enum = 0x8B67
	VERSION_1_1                                                Enum = 1
	VERTEX_PROGRAM_BINDING_NV                                  Enum = 0x864A
	MAX_GEOMETRY_OUTPUT_VERTICES                               Enum = 0x8DE0
	ATTRIB_ARRAY_TYPE_NV                                       Enum = 0x8625
	RELATIVE_QUADRATIC_CURVE_TO_NV                             Enum = 0x0B
	DEBUG_SEVERITY_LOW_ARB                                     Enum = 0x9148
	DEPTH_COMPONENT24_ARB                                      Enum = 0x81A6
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            Enum = 0x8517
	MAP1_VERTEX_ATTRIB9_4_NV                                   Enum = 0x8669
	TRACE_TEXTURES_BIT_MESA                                    Enum = 0x0008
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           Enum = 0x8855
	SAMPLER_2D_RECT_ARB                                        Enum = 0x8B63
	RASTER_POSITION_UNCLIPPED_IBM                              Enum = 0x19262
	LIGHT_MODEL_AMBIENT                                        Enum = 0x0B53
	STENCIL_INDEX8_EXT                                         Enum = 0x8D48
	MAX_COMPUTE_WORK_GROUP_SIZE                                Enum = 0x91BF
	COMPRESSED_SRGB8_ETC2                                      Enum = 0x9275
	MAP_INVALIDATE_RANGE_BIT_EXT                               Enum = 0x0004
	READ_BUFFER                                                Enum = 0x0C02
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              Enum = 0x0D3B
	R                                                          Enum = 0x2002
	HISTOGRAM_WIDTH_EXT                                        Enum = 0x8026
	OPERAND1_RGB_EXT                                           Enum = 0x8591
	DRAW_BUFFER4                                               Enum = 0x8829
	QUARTER_BIT_ATI                                            Enum = 0x00000010
	MIN_MAP_BUFFER_ALIGNMENT                                   Enum = 0x90BC
	COMPRESSED_RGB8_ETC2                                       Enum = 0x9274
	COMPILE_AND_EXECUTE                                        Enum = 0x1301
	UNSIGNED_IDENTITY_NV                                       Enum = 0x8536
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          Enum = 0x85B0
	CLAMP_FRAGMENT_COLOR                                       Enum = 0x891B
	MAX_SAMPLE_MASK_WORDS                                      Enum = 0x8E59
	FUNC_SUBTRACT_EXT                                          Enum = 0x800A
	FOG_COORDINATE_ARRAY_TYPE_EXT                              Enum = 0x8454
	OPERAND3_RGB_NV                                            Enum = 0x8593
	RGB32F_ARB                                                 Enum = 0x8815
	CON_15_ATI                                                 Enum = 0x8950
	MAX_TESS_PATCH_COMPONENTS                                  Enum = 0x8E84
	REFERENCED_BY_FRAGMENT_SHADER                              Enum = 0x930A
	STENCIL_COMPONENTS                                         Enum = 0x8285
	VERTEX_BINDING_STRIDE                                      Enum = 0x82D8
	COLOR_ATTACHMENT8_EXT                                      Enum = 0x8CE8
	COMPRESSED_RGBA_ASTC_10x10_KHR                             Enum = 0x93BB
	EVAL_VERTEX_ATTRIB9_NV                                     Enum = 0x86CF
	CONVEX_HULL_NV                                             Enum = 0x908B
	SHADER_STORAGE_BUFFER_START                                Enum = 0x90D4
	DEBUG_TYPE_ERROR_ARB                                       Enum = 0x824C
	ALPHA_MIN_CLAMP_INGR                                       Enum = 0x8563
	MATRIX4_ARB                                                Enum = 0x88C4
	HALF_FLOAT_ARB                                             Enum = 0x140B
	MINMAX_SINK_EXT                                            Enum = 0x8030
	R8_EXT                                                     Enum = 0x8229
	MANUAL_GENERATE_MIPMAP                                     Enum = 0x8294
	ALPHA_FLOAT16_APPLE                                        Enum = 0x881C
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         Enum = 0x889F
	CON_4_ATI                                                  Enum = 0x8945
	SGIX_resample                                              Enum = 1
	MAX_DRAW_BUFFERS_NV                                        Enum = 0x8824
	COMPRESSED_SLUMINANCE                                      Enum = 0x8C4A
	MAP2_TEXTURE_COORD_4                                       Enum = 0x0DB6
	POSITION                                                   Enum = 0x1203
	FRAMEBUFFER_BLEND                                          Enum = 0x828B
	FOG_COORD_SRC                                              Enum = 0x8450
	TEXTURE_RECTANGLE_NV                                       Enum = 0x84F5
	MAX_VERTEX_STREAMS_ATI                                     Enum = 0x876B
	NORMAL_ARRAY_BUFFER_BINDING                                Enum = 0x8897
	BIAS_BIT_ATI                                               Enum = 0x00000008
	DRAW_INDIRECT_LENGTH_NV                                    Enum = 0x8F42
	SRGB_DECODE_ARB                                            Enum = 0x8299
	VERTEX_PROGRAM_TWO_SIDE                                    Enum = 0x8643
	OFFSET_TEXTURE_BIAS_NV                                     Enum = 0x86E3
	OUTPUT_TEXTURE_COORD12_EXT                                 Enum = 0x87A9
	RENDERBUFFER_BLUE_SIZE_EXT                                 Enum = 0x8D52
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            Enum = 0x40
	TIMEOUT_EXPIRED_APPLE                                      Enum = 0x911B
	COMPUTE_SHADER                                             Enum = 0x91B9
	EXT_point_parameters                                       Enum = 1
	TRACK_MATRIX_NV                                            Enum = 0x8648
	VERTEX_ATTRIB_ARRAY_TYPE                                   Enum = 0x8625
	DRAW_BUFFER15_ATI                                          Enum = 0x8834
	DEPTH24_STENCIL8_EXT                                       Enum = 0x88F0
	BITMAP_TOKEN                                               Enum = 0x0704
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          Enum = 0x8C8A
	POINT_TOKEN                                                Enum = 0x0701
	LUMINANCE16                                                Enum = 0x8042
	MINMAX_SINK                                                Enum = 0x8030
	COMBINE_EXT                                                Enum = 0x8570
	MATRIX14_ARB                                               Enum = 0x88CE
	SRC0_RGB                                                   Enum = 0x8580
	INT_SAMPLER_1D_ARRAY                                       Enum = 0x8DCE
	TEXTURE_COORD_ARRAY_SIZE_EXT                               Enum = 0x8088
	INT_SAMPLER_CUBE_EXT                                       Enum = 0x8DCC
	MAX_NAME_LENGTH                                            Enum = 0x92F6
	MATRIX22_ARB                                               Enum = 0x88D6
	RENDERBUFFER_GREEN_SIZE_EXT                                Enum = 0x8D51
	RGBA_MODE                                                  Enum = 0x0C31
	CONSTANT_ATTENUATION                                       Enum = 0x1207
	MAX                                                        Enum = 0x8008
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            Enum = 0x8519
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      Enum = 0x0F
	VERSION_1_2                                                Enum = 1
	HISTOGRAM_RED_SIZE                                         Enum = 0x8028
	MATRIX4_NV                                                 Enum = 0x8634
	MAP1_VERTEX_ATTRIB0_4_NV                                   Enum = 0x8660
	SWIZZLE_STRQ_DQ_ATI                                        Enum = 0x897B
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        Enum = 0x8C84
	ALPHA16_SNORM                                              Enum = 0x9018
	TEXCOORD1_BIT_PGI                                          Enum = 0x10000000
	FRAMEBUFFER_BARRIER_BIT                                    Enum = 0x00000400
	COLOR_CLEAR_VALUE                                          Enum = 0x0C22
	RENDER_MODE                                                Enum = 0x0C40
	RGB4_EXT                                                   Enum = 0x804F
	SECONDARY_COLOR_ARRAY_EXT                                  Enum = 0x845E
	MAP1_VERTEX_ATTRIB6_4_NV                                   Enum = 0x8666
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            Enum = 0x8896
	SKIP_MISSING_GLYPH_NV                                      Enum = 0x90A9
	TEXTURE_BUFFER_SIZE                                        Enum = 0x919E
	SHADER_BINARY_DMP                                          Enum = 0x9250
	GL_422_AVERAGE_EXT                                         Enum = 0x80CE
	VERTEX_SHADER_VARIANTS_EXT                                 Enum = 0x87D0
	WRITE_PIXEL_DATA_RANGE_NV                                  Enum = 0x8878
	ALREADY_SIGNALED_APPLE                                     Enum = 0x911A
	NAME_LENGTH                                                Enum = 0x92F9
	MAX_VERTEX_HINT_PGI                                        Enum = 0x1A22D
	DRAW_BUFFER11                                              Enum = 0x8830
	LUMINANCE_ALPHA_INTEGER_EXT                                Enum = 0x8D9D
	UNPACK_SKIP_ROWS                                           Enum = 0x0CF3
	CONVOLUTION_2D                                             Enum = 0x8011
	OUTPUT_TEXTURE_COORD30_EXT                                 Enum = 0x87BB
	VERTEX_SHADER_ARB                                          Enum = 0x8B31
	TEXTURE_SWIZZLE_A                                          Enum = 0x8E45
	INT_IMAGE_1D_ARRAY_EXT                                     Enum = 0x905D
	SGIS_point_parameters                                      Enum = 1
	STENCIL_FAIL                                               Enum = 0x0B94
	RED                                                        Enum = 0x1903
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            Enum = 0x8023
	DOT3_RGBA_ARB                                              Enum = 0x86AF
	FRAGMENT_PROGRAM_CALLBACK_MESA                             Enum = 0x8BB1
	VERTEX_ARRAY_TYPE                                          Enum = 0x807B
	MAX_RECTANGLE_TEXTURE_SIZE                                 Enum = 0x84F8
	OFFSET_TEXTURE_2D_SCALE_NV                                 Enum = 0x86E2
	RGBA2                                                      Enum = 0x8055
	UNSIGNED_INT_8_8_8_8_REV                                   Enum = 0x8367
	SRGB_ALPHA                                                 Enum = 0x8C42
	COLOR_LOGIC_OP                                             Enum = 0x0BF2
	FLOAT_MAT2x4                                               Enum = 0x8B66
	SAMPLER_EXTERNAL_OES                                       Enum = 0x8D66
	GREEN_INTEGER                                              Enum = 0x8D95
	VERTEX_SHADER_BIT_EXT                                      Enum = 0x00000001
	AVERAGE_HP                                                 Enum = 0x8160
	OPERAND0_ALPHA                                             Enum = 0x8598
	TEXTURE_BORDER_VALUES_NV                                   Enum = 0x871A
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             Enum = 0x88FE
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          Enum = 0x8C7F
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   Enum = 0x8D56
	VIEWPORT_BIT                                               Enum = 0x00000800
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         Enum = 0x84F1
	VERTEX_ARRAY_RANGE_APPLE                                   Enum = 0x851D
	RGBA_SNORM                                                 Enum = 0x8F93
	COLOR_ATTACHMENT14_EXT                                     Enum = 0x8CEE
	TEXTURE_COVERAGE_SAMPLES_NV                                Enum = 0x9045
	TEXTURE_BINDING_2D_MULTISAMPLE                             Enum = 0x9104
	COLOR_TABLE_BLUE_SIZE_SGI                                  Enum = 0x80DC
	VIEWPORT_INDEX_PROVOKING_VERTEX                            Enum = 0x825F
	INTENSITY_FLOAT16_ATI                                      Enum = 0x881D
	TEXTURE_CUBE_MAP_SEAMLESS                                  Enum = 0x884F
	DOUBLE_MAT3x4                                              Enum = 0x8F4C
	UNSIGNED_INT_IMAGE_1D_ARRAY                                Enum = 0x9068
	SGIS_texture_select                                        Enum = 1
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        Enum = 0x87F7
	MAX_TEXTURE_IMAGE_UNITS_NV                                 Enum = 0x8872
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           Enum = 0x887A
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           Enum = 0x8D68
	GREEN_INTEGER_EXT                                          Enum = 0x8D95
	MAX_COMPUTE_UNIFORM_BLOCKS                                 Enum = 0x91BB
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   Enum = 0x9276
	ARRAY_BUFFER_BINDING_ARB                                   Enum = 0x8894
	MAX_VARYING_COMPONENTS_EXT                                 Enum = 0x8B4B
	LUMINANCE8_SNORM                                           Enum = 0x9015
	TEXTURE_ENV_MODE                                           Enum = 0x2200
	READ_BUFFER_EXT                                            Enum = 0x0C02
	MAX_TEXTURE_COORDS                                         Enum = 0x8871
	DUDV_ATI                                                   Enum = 0x8779
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              Enum = 0x8A42
	CLIP_NEAR_HINT_PGI                                         Enum = 0x1A220
	UNPACK_LSB_FIRST                                           Enum = 0x0CF1
	SELECTION_BUFFER_SIZE                                      Enum = 0x0DF4
	LIGHT7                                                     Enum = 0x4007
	RG16I                                                      Enum = 0x8239
	TRANSFORM_FEEDBACK_VARYINGS                                Enum = 0x8C83
	RENDERBUFFER_BLUE_SIZE_OES                                 Enum = 0x8D52
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             Enum = 0x8F25
	TRANSFORM_BIT                                              Enum = 0x00001000
	TEXTURE_2D                                                 Enum = 0x0DE1
	WRAP_BORDER_SUN                                            Enum = 0x81D4
	TEXTURE28                                                  Enum = 0x84DC
	TEXTURE_BINDING_CUBE_MAP_EXT                               Enum = 0x8514
	VARIANT_ARRAY_POINTER_EXT                                  Enum = 0x87E9
	UNSIGNED_INT_VEC3_EXT                                      Enum = 0x8DC7
	UNSIGNED_INT_IMAGE_CUBE_EXT                                Enum = 0x9066
	ALLOW_DRAW_MEM_HINT_PGI                                    Enum = 0x1A211
	FEEDBACK_BUFFER_TYPE                                       Enum = 0x0DF2
	VIEW_CLASS_16_BITS                                         Enum = 0x82CA
	TEXTURE_ALPHA_MODULATE_IMG                                 Enum = 0x8C06
	SRGB8_ALPHA8                                               Enum = 0x8C43
	NAMED_STRING_TYPE_ARB                                      Enum = 0x8DEA
	TRANSFORM_FEEDBACK_PAUSED                                  Enum = 0x8E23
	RGB10_A2                                                   Enum = 0x8059
	UNSIGNED_SHORT_4_4_4_4                                     Enum = 0x8033
	TEXTURE4                                                   Enum = 0x84C4
	UNPACK_ROW_BYTES_APPLE                                     Enum = 0x8A16
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   Enum = 0x8DA8
	TEXTURE_COORD_ARRAY_LENGTH_NV                              Enum = 0x8F2F
	INT_IMAGE_3D_EXT                                           Enum = 0x9059
	NORMAL_MAP_ARB                                             Enum = 0x8511
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          Enum = 0x886C
	MATRIX28_ARB                                               Enum = 0x88DC
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         Enum = 0x9160
	TEXTURE_LOD_BIAS_S_SGIX                                    Enum = 0x818E
	VIEWPORT_SUBPIXEL_BITS                                     Enum = 0x825C
	PRIMARY_COLOR                                              Enum = 0x8577
	MATRIX7_NV                                                 Enum = 0x8637
	SIGNED_RGB8_NV                                             Enum = 0x86FF
	MODULATE_ADD_ATI                                           Enum = 0x8744
	MAX_PROGRAM_LOOP_COUNT_NV                                  Enum = 0x88F8
	POINT_SIZE_GRANULARITY                                     Enum = 0x0B13
	RESAMPLE_ZERO_FILL_OML                                     Enum = 0x8987
	SEPARATE_ATTRIBS_NV                                        Enum = 0x8C8D
	FONT_X_MAX_BOUNDS_BIT_NV                                   Enum = 0x00040000
	LUMINANCE_FLOAT32_ATI                                      Enum = 0x8818
	MAX_DRAW_BUFFERS_ATI                                       Enum = 0x8824
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   Enum = 0x8868
	POST_COLOR_MATRIX_RED_SCALE_SGI                            Enum = 0x80B4
	MAX_ELEMENTS_VERTICES_EXT                                  Enum = 0x80E8
	MAP2_VERTEX_ATTRIB12_4_NV                                  Enum = 0x867C
	DRAW_BUFFER12_ATI                                          Enum = 0x8831
	INT_IMAGE_2D                                               Enum = 0x9058
	ADD                                                        Enum = 0x0104
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             Enum = 0x82B1
	REGISTER_COMBINERS_NV                                      Enum = 0x8522
	RGB5_A1_OES                                                Enum = 0x8057
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           Enum = 0x880C
	TEXTURE_ALPHA_TYPE_ARB                                     Enum = 0x8C13
	MAX_VERTEX_UNIFORM_VECTORS                                 Enum = 0x8DFB
	PIXEL_MAP_R_TO_R_SIZE                                      Enum = 0x0CB6
	DEPTH_COMPONENT24_OES                                      Enum = 0x81A6
	VARIANT_ARRAY_EXT                                          Enum = 0x87E8
	CURRENT_QUERY_EXT                                          Enum = 0x8865
	VERTICAL_LINE_TO_NV                                        Enum = 0x08
	POLYGON_TOKEN                                              Enum = 0x0703
	SECONDARY_COLOR_ARRAY_SIZE                                 Enum = 0x845A
	COMBINER_AB_OUTPUT_NV                                      Enum = 0x854A
	SRC2_ALPHA                                                 Enum = 0x858A
	MVP_MATRIX_EXT                                             Enum = 0x87E3
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            Enum = 0x8B88
	COUNTER_RANGE_AMD                                          Enum = 0x8BC1
	TEXTURE_FORMAT_QCOM                                        Enum = 0x8BD6
	PROXY_TEXTURE_2D_ARRAY_EXT                                 Enum = 0x8C1B
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             Enum = 0x8E2F
	RG8_SNORM                                                  Enum = 0x8F95
	VERTEX_ARRAY_LIST_STRIDE_IBM                               Enum = 103080
	MAX_NAME_STACK_DEPTH                                       Enum = 0x0D37
	OBJECT_DISTANCE_TO_POINT_SGIS                              Enum = 0x81F1
	MATRIX13_ARB                                               Enum = 0x88CD
	UNIFORM_MATRIX_STRIDE                                      Enum = 0x8A3D
	HISTOGRAM_GREEN_SIZE                                       Enum = 0x8029
	VARIABLE_A_NV                                              Enum = 0x8523
	SHADER_OPERATION_NV                                        Enum = 0x86DF
	RGB9_E5_EXT                                                Enum = 0x8C3D
	MAP_INVALIDATE_RANGE_BIT                                   Enum = 0x0004
	TEXTURE_ENV_COLOR                                          Enum = 0x2201
	OP_SUB_EXT                                                 Enum = 0x8796
	SECONDARY_INTERPOLATOR_ATI                                 Enum = 0x896D
	COMPRESSED_SLUMINANCE_EXT                                  Enum = 0x8C4A
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         Enum = 0x8CDB
	CURRENT_TEXTURE_COORDS                                     Enum = 0x0B03
	FRAMEBUFFER_RENDERABLE_LAYERED                             Enum = 0x828A
	TRANSPOSE_COLOR_MATRIX_ARB                                 Enum = 0x84E6
	MAP2_VERTEX_ATTRIB9_4_NV                                   Enum = 0x8679
	VARIANT_VALUE_EXT                                          Enum = 0x87E4
	FLOAT_VEC3                                                 Enum = 0x8B51
	OFFSET                                                     Enum = 0x92FC
	CLIP_FAR_HINT_PGI                                          Enum = 0x1A221
	LINEAR_SHARPEN_ALPHA_SGIS                                  Enum = 0x80AE
	IMAGE_TEXEL_SIZE                                           Enum = 0x82A7
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          Enum = 0x88AB
	REG_2_ATI                                                  Enum = 0x8923
	MAX_VARYING_FLOATS                                         Enum = 0x8B4B
	FRAMEBUFFER_BINDING_OES                                    Enum = 0x8CA6
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             Enum = 0x9124
	COEFF                                                      Enum = 0x0A00
	AND                                                        Enum = 0x1501
	COLOR_ARRAY_TYPE_EXT                                       Enum = 0x8082
	MATRIX5_NV                                                 Enum = 0x8635
	MATRIX_INDEX_ARRAY_POINTER_OES                             Enum = 0x8849
	ITALIC_BIT_NV                                              Enum = 0x02
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        Enum = 0x92CF
	STATIC_DRAW                                                Enum = 0x88E4
	PRESENT_TIME_NV                                            Enum = 0x8E2A
	RED_SNORM                                                  Enum = 0x8F90
	UNSIGNED_INT_2_10_10_10_REV_EXT                            Enum = 0x8368
	RGB16_EXT                                                  Enum = 0x8054
	MAP2_TANGENT_EXT                                           Enum = 0x8445
	SIGNED_INTENSITY_NV                                        Enum = 0x8707
	STREAM_DRAW_ARB                                            Enum = 0x88E0
	COLOR_ATTACHMENT12_EXT                                     Enum = 0x8CEC
	COLOR_TABLE_LUMINANCE_SIZE                                 Enum = 0x80DE
	TEXTURE_STORAGE_HINT_APPLE                                 Enum = 0x85BC
	WRITEONLY_RENDERING_QCOM                                   Enum = 0x8823
	ACTIVE_STENCIL_FACE_EXT                                    Enum = 0x8911
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       Enum = 0x8DA3
	FOG_COORD_ARRAY_LENGTH_NV                                  Enum = 0x8F32
	QUERY_OBJECT_EXT                                           Enum = 0x9153
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 Enum = 0x8176
	OFFSET_TEXTURE_RECTANGLE_NV                                Enum = 0x864C
	FULL_RANGE_EXT                                             Enum = 0x87E1
	PALETTE4_R5_G6_B5_OES                                      Enum = 0x8B92
	LUMINANCE12_EXT                                            Enum = 0x8041
	NORMAL_ARRAY_EXT                                           Enum = 0x8075
	TEXTURE8                                                   Enum = 0x84C8
	VIBRANCE_BIAS_NV                                           Enum = 0x8719
	RGB_422_APPLE                                              Enum = 0x8A1F
	INT_VEC2                                                   Enum = 0x8B53
	STENCIL_INDEX4_EXT                                         Enum = 0x8D47
	COLOR_ATTACHMENT6_NV                                       Enum = 0x8CE6
	UNSIGNED_INT_IMAGE_2D_EXT                                  Enum = 0x9063
	SHADER_STORAGE_BUFFER_BINDING                              Enum = 0x90D3
	PIXEL_MAP_I_TO_B_SIZE                                      Enum = 0x0CB4
	TEXTURE18                                                  Enum = 0x84D2
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       Enum = 0x8DA4
	CIRCULAR_TANGENT_ARC_TO_NV                                 Enum = 0xFC
	PIXEL_BUFFER_BARRIER_BIT_EXT                               Enum = 0x00000080
	VERTEX_ARRAY_SIZE                                          Enum = 0x807A
	MIRROR_CLAMP_TO_BORDER_EXT                                 Enum = 0x8912
	OBJECT_ATTACHED_OBJECTS_ARB                                Enum = 0x8B85
	INT16_VEC2_NV                                              Enum = 0x8FE5
	CONVOLUTION_HEIGHT                                         Enum = 0x8019
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                Enum = 0x83F8
	RESAMPLE_REPLICATE_SGIX                                    Enum = 0x842E
	MAX_VERTEX_ATTRIBS_ARB                                     Enum = 0x8869
	FLOAT_RGB32_NV                                             Enum = 0x8889
	DEBUG_SOURCE_APPLICATION_ARB                               Enum = 0x824A
	RGB2_EXT                                                   Enum = 0x804E
	MAP2_VERTEX_ATTRIB2_4_NV                                   Enum = 0x8672
	BLEND_EQUATION_ALPHA_OES                                   Enum = 0x883D
	MAX_VARYING_FLOATS_ARB                                     Enum = 0x8B4B
	UNSIGNED_INT16_VEC4_NV                                     Enum = 0x8FF3
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       Enum = 0x8243
	TEXTURE18_ARB                                              Enum = 0x84D2
	TEXTURE_COMPRESSED_ARB                                     Enum = 0x86A1
	OP_DOT4_EXT                                                Enum = 0x8785
	FLOAT_MAT2                                                 Enum = 0x8B5A
	SHADING_LANGUAGE_VERSION                                   Enum = 0x8B8C
	TEXTURE_PRIORITY_EXT                                       Enum = 0x8066
	SAMPLES_EXT                                                Enum = 0x80A9
	WEIGHT_ARRAY_POINTER_OES                                   Enum = 0x86AC
	SAMPLE_POSITION                                            Enum = 0x8E50
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     Enum = 0x8E8F
	RG8_EXT                                                    Enum = 0x822B
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              Enum = 0x84F8
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          Enum = 0x8F39
	EMISSION                                                   Enum = 0x1600
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           Enum = 0x80B3
	PROGRAM_LENGTH_ARB                                         Enum = 0x8627
	WEIGHT_ARRAY_TYPE_ARB                                      Enum = 0x86A9
	INT_IMAGE_2D_RECT_EXT                                      Enum = 0x905A
	PATH_GEN_COMPONENTS_NV                                     Enum = 0x90B3
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         Enum = 0x90BE
	DEBUG_OUTPUT                                               Enum = 0x92E0
	DUAL_ALPHA8_SGIS                                           Enum = 0x8111
	FENCE_APPLE                                                Enum = 0x8A0B
	UNIFORM_BLOCK_DATA_SIZE                                    Enum = 0x8A40
	IMAGE_BINDING_ACCESS_EXT                                   Enum = 0x8F3E
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             Enum = 0x00001000
	HISTOGRAM_SINK_EXT                                         Enum = 0x802D
	TABLE_TOO_LARGE_EXT                                        Enum = 0x8031
	PROGRAM_SEPARABLE                                          Enum = 0x8258
	INT_VEC3                                                   Enum = 0x8B54
	PACK_COMPRESSED_BLOCK_WIDTH                                Enum = 0x912B
	OPERAND0_RGB_ARB                                           Enum = 0x8590
	MAX_TRACK_MATRICES_NV                                      Enum = 0x862F
	DRAW_BUFFER15                                              Enum = 0x8834
	UNSIGNED_INT_5_9_9_9_REV                                   Enum = 0x8C3E
	MAX_COMPUTE_LOCAL_INVOCATIONS                              Enum = 0x90EB
	LAYOUT_LINEAR_INTEL                                        Enum = 1
	PIXEL_MAP_I_TO_I                                           Enum = 0x0C70
	MAP2_NORMAL                                                Enum = 0x0DB2
	BINORMAL_ARRAY_TYPE_EXT                                    Enum = 0x8440
	DOT_PRODUCT_TEXTURE_3D_NV                                  Enum = 0x86EF
	COVERAGE_ATTACHMENT_NV                                     Enum = 0x8ED2
	XOR                                                        Enum = 0x1506
	SAMPLE_ALPHA_TO_ONE_ARB                                    Enum = 0x809F
	MAX_SAMPLES_EXT                                            Enum = 0x8D57
	FLOAT_RGBA32_NV                                            Enum = 0x888B
	PATH_STROKE_MASK_NV                                        Enum = 0x9084
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          Enum = 0x08
	TEXTURE_PRE_SPECULAR_HP                                    Enum = 0x8169
	DECODE_EXT                                                 Enum = 0x8A49
	RGB8_SNORM                                                 Enum = 0x8F96
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        Enum = 0x92D4
	SGIS_detail_texture                                        Enum = 1
	R16                                                        Enum = 0x822A
	BUMP_NUM_TEX_UNITS_ATI                                     Enum = 0x8777
	UNSIGNED_INT_SAMPLER_3D                                    Enum = 0x8DD3
	UNPACK_ROW_LENGTH                                          Enum = 0x0CF2
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             Enum = 0x850F
	RASTERIZER_DISCARD_NV                                      Enum = 0x8C89
	SAMPLER_1D_ARRAY                                           Enum = 0x8DC0
	MAX_SPARSE_TEXTURE_SIZE_AMD                                Enum = 0x9198
	HISTOGRAM_BLUE_SIZE                                        Enum = 0x802A
	TEXTURE17_ARB                                              Enum = 0x84D1
	PASS_THROUGH_NV                                            Enum = 0x86E6
	RGB_FLOAT32_ATI                                            Enum = 0x8815
	MAX_VERTEX_UNIFORM_COMPONENTS                              Enum = 0x8B4A
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                Enum = 0x8B8A
	INT_SAMPLER_1D_EXT                                         Enum = 0x8DC9
	ALPHA_BIAS                                                 Enum = 0x0D1D
	INDEX_MATERIAL_PARAMETER_EXT                               Enum = 0x81B9
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               Enum = 0x850D
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            Enum = 0x8516
	IDENTITY_NV                                                Enum = 0x862A
	CLAMP_READ_COLOR                                           Enum = 0x891C
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        Enum = 0x8162
	VERTEX_PRECLIP_HINT_SGIX                                   Enum = 0x83EF
	OFFSET_HILO_TEXTURE_2D_NV                                  Enum = 0x8854
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      Enum = 0x8E5E
	MAP2_COLOR_4                                               Enum = 0x0DB0
	PROXY_TEXTURE_4D_SGIS                                      Enum = 0x8135
	PROGRAM_POINT_SIZE_ARB                                     Enum = 0x8642
	DEPENDENT_AR_TEXTURE_2D_NV                                 Enum = 0x86E9
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            Enum = 0x00000001
	C4UB_V3F                                                   Enum = 0x2A23
	POST_CONVOLUTION_BLUE_BIAS                                 Enum = 0x8022
	CONVOLUTION_HINT_SGIX                                      Enum = 0x8316
	RENDERBUFFER_HEIGHT_OES                                    Enum = 0x8D43
	RGB16UI                                                    Enum = 0x8D77
	COMPRESSED_RGBA_ASTC_5x5_KHR                               Enum = 0x93B2
	MAX_ARRAY_TEXTURE_LAYERS                                   Enum = 0x88FF
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              Enum = 0x8F2A
	LIGHT3                                                     Enum = 0x4003
	MATRIX_EXT                                                 Enum = 0x87C0
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              Enum = 0x8A02
	TEXTURE_SHARED_SIZE                                        Enum = 0x8C3F
	FRAMEBUFFER_UNSUPPORTED_EXT                                Enum = 0x8CDD
	LIST_BASE                                                  Enum = 0x0B32
	RGB16                                                      Enum = 0x8054
	CLAMP_TO_BORDER_ARB                                        Enum = 0x812D
	MAP1_VERTEX_ATTRIB12_4_NV                                  Enum = 0x866C
	SATURATE_BIT_ATI                                           Enum = 0x00000040
	MAX_TEXTURE_BUFFER_SIZE_EXT                                Enum = 0x8C2B
	INTENSITY8_EXT                                             Enum = 0x804B
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          Enum = 0x8216
	PROGRAM                                                    Enum = 0x82E2
	HALF_BIT_ATI                                               Enum = 0x00000008
	DSDT_MAG_INTENSITY_NV                                      Enum = 0x86DC
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        Enum = 0x8809
	ARRAY_BUFFER_BINDING                                       Enum = 0x8894
	FRAMEBUFFER_COMPLETE_EXT                                   Enum = 0x8CD5
	LUMINANCE8UI_EXT                                           Enum = 0x8D80
	CLIP_PLANE2                                                Enum = 0x3002
	BLEND_DST_RGB_EXT                                          Enum = 0x80C8
	TRANSPOSE_MODELVIEW_MATRIX                                 Enum = 0x84E3
	COMBINER_MAPPING_NV                                        Enum = 0x8543
	OP_ROUND_EXT                                               Enum = 0x8790
	PALETTE4_RGBA8_OES                                         Enum = 0x8B91
	QUERY_NO_WAIT                                              Enum = 0x8E14
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             Enum = 0x9061
	DEBUG_CATEGORY_API_ERROR_AMD                               Enum = 0x9149
	PACK_LSB_FIRST                                             Enum = 0x0D01
	MULTISAMPLE_ARB                                            Enum = 0x809D
	LOSE_CONTEXT_ON_RESET_ARB                                  Enum = 0x8252
	VBO_FREE_MEMORY_ATI                                        Enum = 0x87FB
	DRAW_BUFFER8_NV                                            Enum = 0x882D
	BOOL_ARB                                                   Enum = 0x8B56
	TEXTURE_BINDING_1D_ARRAY                                   Enum = 0x8C1C
	SGIX_framezoom                                             Enum = 1
	DEPTH_COMPONENT16                                          Enum = 0x81A5
	DEPTH_ATTACHMENT_EXT                                       Enum = 0x8D00
	GEOMETRY_SHADER_INVOCATIONS                                Enum = 0x887F
	TEXTURE_SHADER_NV                                          Enum = 0x86DE
	REG_14_ATI                                                 Enum = 0x892F
	MAP2_GRID_DOMAIN                                           Enum = 0x0DD2
	GEOMETRY_DEFORMATION_SGIX                                  Enum = 0x8194
	SECONDARY_COLOR_ARRAY                                      Enum = 0x845E
	RECIP_ADD_SIGNED_ALPHA_IMG                                 Enum = 0x8C05
	TEXTURE1                                                   Enum = 0x84C1
	SAMPLER_2D_SHADOW_EXT                                      Enum = 0x8B62
	INFO_LOG_LENGTH                                            Enum = 0x8B84
	BINORMAL_ARRAY_STRIDE_EXT                                  Enum = 0x8441
	PROGRAM_NATIVE_TEMPORARIES_ARB                             Enum = 0x88A6
	MATRIX8_ARB                                                Enum = 0x88C8
	UNSIGNALED                                                 Enum = 0x9118
	MAX_CONVOLUTION_WIDTH                                      Enum = 0x801A
	MULTISAMPLE                                                Enum = 0x809D
	TEXTURE_POST_SPECULAR_HP                                   Enum = 0x8168
	SRGB_ALPHA_EXT                                             Enum = 0x8C42
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    Enum = 0x92CD
	MAX_VERTEX_ATOMIC_COUNTERS                                 Enum = 0x92D2
	BACK_PRIMARY_COLOR_NV                                      Enum = 0x8C77
	BUMP_ROT_MATRIX_ATI                                        Enum = 0x8775
	TEXTURE_SHARED_SIZE_EXT                                    Enum = 0x8C3F
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        Enum = 0x92D1
	MATRIX5_ARB                                                Enum = 0x88C5
	MAX_DEBUG_LOGGED_MESSAGES                                  Enum = 0x9144
	MAX_FRAMEBUFFER_LAYERS                                     Enum = 0x9317
	ALL_ATTRIB_BITS                                            Enum = 0xFFFFFFFF
	CONSTANT_ARB                                               Enum = 0x8576
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           Enum = 0x8802
	TEXTURE_DEPTH_QCOM                                         Enum = 0x8BD4
	PATH_STROKE_WIDTH_NV                                       Enum = 0x9075
	POST_COLOR_MATRIX_RED_BIAS                                 Enum = 0x80B8
	STENCIL_BACK_FUNC_ATI                                      Enum = 0x8800
	HILO8_NV                                                   Enum = 0x885E
	TESS_CONTROL_OUTPUT_VERTICES                               Enum = 0x8E75
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        Enum = 0x900F
	GL_4PASS_3_SGIS                                            Enum = 0x80A7
	ACTIVE_SUBROUTINE_UNIFORMS                                 Enum = 0x8DE6
	LINE_RESET_TOKEN                                           Enum = 0x0707
	DRAW_BUFFER_EXT                                            Enum = 0x0C01
	PACK_SKIP_VOLUMES_SGIS                                     Enum = 0x8130
	IUI_V3F_EXT                                                Enum = 0x81AE
	FRAMEBUFFER_RENDERABLE                                     Enum = 0x8289
	BUMP_ROT_MATRIX_SIZE_ATI                                   Enum = 0x8776
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              Enum = 0x851C
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          Enum = 0x887C
	RGBA_INTEGER_MODE_EXT                                      Enum = 0x8D9E
	NORMAL_ARRAY_LENGTH_NV                                     Enum = 0x8F2C
	INT16_VEC3_NV                                              Enum = 0x8FE6
	CONVOLUTION_HEIGHT_EXT                                     Enum = 0x8019
	STATIC_DRAW_ARB                                            Enum = 0x88E4
	FRAMEBUFFER_UNSUPPORTED                                    Enum = 0x8CDD
	TEXTURE_EXTERNAL_OES                                       Enum = 0x8D65
	MAP_STENCIL                                                Enum = 0x0D11
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            Enum = 0x809C
	COLOR_TABLE_BLUE_SIZE                                      Enum = 0x80DC
	CURRENT_QUERY                                              Enum = 0x8865
	LUMINANCE_ALPHA_FLOAT16_APPLE                              Enum = 0x881F
	CON_31_ATI                                                 Enum = 0x8960
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     Enum = 0x8CD1
	UNSIGNED_INT_IMAGE_CUBE                                    Enum = 0x9066
	SYSTEM_FONT_NAME_NV                                        Enum = 0x9073
	PERSPECTIVE_CORRECTION_HINT                                Enum = 0x0C50
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             Enum = 0x808A
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           Enum = 0x80B6
	MAX_UNIFORM_LOCATIONS                                      Enum = 0x826E
	SOURCE2_ALPHA_ARB                                          Enum = 0x858A
	COLOR_BUFFER_BIT                                           Enum = 0x00004000
	REPLACE_EXT                                                Enum = 0x8062
	REFLECTION_MAP_OES                                         Enum = 0x8512
	SOURCE1_ALPHA_EXT                                          Enum = 0x8589
	READ_FRAMEBUFFER_NV                                        Enum = 0x8CA8
	DEPTH_WRITEMASK                                            Enum = 0x0B72
	INDEX_ARRAY_STRIDE                                         Enum = 0x8086
	TEXTURE_GATHER                                             Enum = 0x82A2
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            Enum = 0x8518
	OUTPUT_TEXTURE_COORD31_EXT                                 Enum = 0x87BC
	STENCIL_ATTACHMENT_EXT                                     Enum = 0x8D20
	DEBUG_OUTPUT_SYNCHRONOUS                                   Enum = 0x8242
	NUM_PROGRAM_BINARY_FORMATS_OES                             Enum = 0x87FE
	CON_2_ATI                                                  Enum = 0x8943
	TEXTURE_SWIZZLE_RGBA                                       Enum = 0x8E46
	STANDARD_FONT_NAME_NV                                      Enum = 0x9072
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           Enum = 0x880D
	PATH_DASH_OFFSET_RESET_NV                                  Enum = 0x90B4
	INTENSITY16UI_EXT                                          Enum = 0x8D79
	HISTOGRAM_BLUE_SIZE_EXT                                    Enum = 0x802A
	FIRST_VERTEX_CONVENTION                                    Enum = 0x8E4D
	GL_3D_COLOR                                                Enum = 0x0602
	TEXTURE_COMPONENTS                                         Enum = 0x1003
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            Enum = 0x8515
	MAX_SAMPLES_ANGLE                                          Enum = 0x8D57
	VERTEX_ARRAY_POINTER_EXT                                   Enum = 0x808E
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     Enum = 0x8337
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       Enum = 0x8C29
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      Enum = 0x8211
	WEIGHT_ARRAY_TYPE_OES                                      Enum = 0x86A9
	OUTPUT_COLOR0_EXT                                          Enum = 0x879B
	LUMINANCE16F_ARB                                           Enum = 0x881E
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   Enum = 0x889C
	POLYGON_STIPPLE_BIT                                        Enum = 0x00000010
	BINORMAL_ARRAY_EXT                                         Enum = 0x843A
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          Enum = 0x92C7
	INDEX_WRITEMASK                                            Enum = 0x0C21
	COPY                                                       Enum = 0x1503
	RGB12                                                      Enum = 0x8053
	DRAW_BUFFER8_ATI                                           Enum = 0x882D
	MAX_FRAGMENT_UNIFORM_BLOCKS                                Enum = 0x8A2D
	PROGRAM_OBJECT_EXT                                         Enum = 0x8B40
	PARTIAL_SUCCESS_NV                                         Enum = 0x902E
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         Enum = 0x00000004
	BUFFER_MAP_POINTER_ARB                                     Enum = 0x88BD
	SGIX_texture_coordinate_clamp                              Enum = 1
	R8                                                         Enum = 0x8229
	TEXTURE6                                                   Enum = 0x84C6
	COMBINER_AB_DOT_PRODUCT_NV                                 Enum = 0x8545
	EDGE_FLAG_ARRAY_EXT                                        Enum = 0x8079
	CULL_VERTEX_OBJECT_POSITION_EXT                            Enum = 0x81AC
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         Enum = 0x86A2
	QUAD_MESH_SUN                                              Enum = 0x8614
	DRAW_BUFFER5_ATI                                           Enum = 0x882A
	STREAM_COPY                                                Enum = 0x88E2
	PIXEL_BUFFER_BARRIER_BIT                                   Enum = 0x00000080
	SRC0_ALPHA                                                 Enum = 0x8588
	TEXTURE_TARGET_QCOM                                        Enum = 0x8BDA
	RENDERBUFFER_STENCIL_SIZE                                  Enum = 0x8D55
	MULTISAMPLE_COVERAGE_MODES_NV                              Enum = 0x8E12
	COMPRESSED_RGBA_ASTC_12x12_KHR                             Enum = 0x93BD
	MAP_INVALIDATE_BUFFER_BIT_EXT                              Enum = 0x0008
	MAX_ATTRIB_STACK_DEPTH                                     Enum = 0x0D35
	MAP2_VERTEX_4                                              Enum = 0x0DB8
	UNSIGNED_INT_10_10_10_2                                    Enum = 0x8036
	FRAGMENT_MATERIAL_EXT                                      Enum = 0x8349
	TEXTURE11_ARB                                              Enum = 0x84CB
	HILO16_NV                                                  Enum = 0x86F8
	COLOR_ATTACHMENT11_NV                                      Enum = 0x8CEB
	MAX_IMAGE_SAMPLES                                          Enum = 0x906D
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              Enum = 0x8338
	DOUBLE_MAT4_EXT                                            Enum = 0x8F48
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               Enum = 0x905F
	QUERY_RESULT_NO_WAIT_AMD                                   Enum = 0x9194
	HINT_BIT                                                   Enum = 0x00008000
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         Enum = 0x824D
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          Enum = 0x8402
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            Enum = 0x8515
	COLOR_MATERIAL                                             Enum = 0x0B57
	FRAGMENT_LIGHT1_SGIX                                       Enum = 0x840D
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         Enum = 0x92DC
	FOG_SPECULAR_TEXTURE_WIN                                   Enum = 0x80EC
	TEXTURE7_ARB                                               Enum = 0x84C7
	IMAGE_BINDING_ACCESS                                       Enum = 0x8F3E
	RGBA8_SNORM                                                Enum = 0x8F97
	CONSTANT_ALPHA_EXT                                         Enum = 0x8003
	COLOR_INDEX4_EXT                                           Enum = 0x80E4
	TEXTURE_DEFORMATION_SGIX                                   Enum = 0x8195
	PALETTE4_RGB8_OES                                          Enum = 0x8B90
	COLOR_ATTACHMENT13_EXT                                     Enum = 0x8CED
	STENCIL_INDEX8                                             Enum = 0x8D48
	TEXTURE_2D_MULTISAMPLE_ARRAY                               Enum = 0x9102
	PROGRAM_BINARY_ANGLE                                       Enum = 0x93A6
	TEXTURE_RED_SIZE_EXT                                       Enum = 0x805C
	NORMAL_ARRAY_TYPE_EXT                                      Enum = 0x807E
	COLOR_MATRIX_STACK_DEPTH_SGI                               Enum = 0x80B2
	TEXT_FRAGMENT_SHADER_ATI                                   Enum = 0x8200
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              Enum = 0x86DB
	NORMAL_BIT_PGI                                             Enum = 0x08000000
	UNPACK_SWAP_BYTES                                          Enum = 0x0CF0
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         Enum = 0x8174
	MAX_TEXTURE_COORDS_ARB                                     Enum = 0x8871
	CLIP_DISTANCE_NV                                           Enum = 0x8C7A
	MAX_DEPTH_TEXTURE_SAMPLES                                  Enum = 0x910F
	VERTEX_SUBROUTINE_UNIFORM                                  Enum = 0x92EE
	ALPHA16                                                    Enum = 0x803E
	FENCE_CONDITION_NV                                         Enum = 0x84F4
	REG_1_ATI                                                  Enum = 0x8922
	INTERLACE_OML                                              Enum = 0x8980
	COLOR_ATTACHMENT1_NV                                       Enum = 0x8CE1
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         Enum = 0x9061
	MAT_AMBIENT_BIT_PGI                                        Enum = 0x00100000
	SGIX_flush_raster                                          Enum = 1
	LINES_ADJACENCY_EXT                                        Enum = 0x000A
	SHADE_MODEL                                                Enum = 0x0B54
	COLOR_INDEX1_EXT                                           Enum = 0x80E2
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        Enum = 0x8311
	OPERAND2_ALPHA_ARB                                         Enum = 0x859A
	FLOAT_RGBA_MODE_NV                                         Enum = 0x888E
	DYNAMIC_COPY_ARB                                           Enum = 0x88EA
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     Enum = 0x8B8A
	MAX_MULTIVIEW_BUFFERS_EXT                                  Enum = 0x90F2
	SYNC_FLAGS                                                 Enum = 0x9115
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           Enum = 0x92D3
	CLAMP_TO_BORDER_NV                                         Enum = 0x812D
	COLOR_ATTACHMENT13                                         Enum = 0x8CED
	MAX_FRAGMENT_UNIFORM_VECTORS                               Enum = 0x8DFD
	ELEMENT_ARRAY_UNIFIED_NV                                   Enum = 0x8F1F
	PATH_DASH_OFFSET_NV                                        Enum = 0x907E
	TEXTURE_IMMUTABLE_FORMAT                                   Enum = 0x912F
	VERTEX23_BIT_PGI                                           Enum = 0x00000004
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          Enum = 0x88B1
	TEXTURE_BLUE_TYPE_ARB                                      Enum = 0x8C12
	GREEN_SCALE                                                Enum = 0x0D18
	UNSIGNED_SHORT_5_6_5                                       Enum = 0x8363
	PREVIOUS_TEXTURE_INPUT_NV                                  Enum = 0x86E4
	TEXTURE_HI_SIZE_NV                                         Enum = 0x871B
	ONE_EXT                                                    Enum = 0x87DE
	BUFFER_SERIALIZED_MODIFY_APPLE                             Enum = 0x8A12
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          Enum = 0x8CD6
	FIXED_OES                                                  Enum = 0x140C
	DEBUG_CALLBACK_USER_PARAM_ARB                              Enum = 0x8245
	COMPRESSED_RG11_EAC                                        Enum = 0x9272
	RGBA12                                                     Enum = 0x805A
	SRGB_EXT                                                   Enum = 0x8C40
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        Enum = 0x8DDD
	TEXTURE_SWIZZLE_B                                          Enum = 0x8E44
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       Enum = 0x817C
	NUM_SHADING_LANGUAGE_VERSIONS                              Enum = 0x82E9
	MAP_ATTRIB_U_ORDER_NV                                      Enum = 0x86C3
	TEXTURE_SWIZZLE_R_EXT                                      Enum = 0x8E42
	BLEND_COLOR                                                Enum = 0x8005
	PN_TRIANGLES_ATI                                           Enum = 0x87F0
	TEXTURE_INTENSITY_TYPE                                     Enum = 0x8C15
	INT_SAMPLER_3D_EXT                                         Enum = 0x8DCB
	BUFFER_OBJECT_EXT                                          Enum = 0x9151
	UNSIGNED_SHORT_5_6_5_REV_EXT                               Enum = 0x8364
	INT_SAMPLER_RENDERBUFFER_NV                                Enum = 0x8E57
	RGBA16_EXT                                                 Enum = 0x805B
	COMBINER_COMPONENT_USAGE_NV                                Enum = 0x8544
	SAMPLES_3DFX                                               Enum = 0x86B4
	DRAW_BUFFER7_NV                                            Enum = 0x882C
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   Enum = 0x8CD2
	NO_ERROR                                                   Enum = 0
	INTENSITY32I_EXT                                           Enum = 0x8D85
	TRIANGLES_ADJACENCY_EXT                                    Enum = 0x000C
	CLAMP_TO_EDGE                                              Enum = 0x812F
	TEXTURE13_ARB                                              Enum = 0x84CD
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       Enum = 0x8E5F
	COLOR_ATTACHMENT3_NV                                       Enum = 0x8CE3
	LINE_TOKEN                                                 Enum = 0x0702
	MAX_FRAMEZOOM_FACTOR_SGIX                                  Enum = 0x818D
	PER_STAGE_CONSTANTS_NV                                     Enum = 0x8535
	SAMPLER_2D_ARB                                             Enum = 0x8B5E
	SLUMINANCE_ALPHA_EXT                                       Enum = 0x8C44
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              Enum = 0x8CD7
	DEBUG_SOURCE_SHADER_COMPILER                               Enum = 0x8248
	FRAGMENT_DEPTH_EXT                                         Enum = 0x8452
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       Enum = 0x8A43
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             Enum = 0x8C80
	RGBA16UI_EXT                                               Enum = 0x8D76
	COMPUTE_SHADER_BIT                                         Enum = 0x00000020
	PIXEL_MAP_R_TO_R                                           Enum = 0x0C76
	UNPACK_SKIP_PIXELS                                         Enum = 0x0CF4
	DONT_CARE                                                  Enum = 0x1100
	LUMINANCE12_ALPHA12_EXT                                    Enum = 0x8047
	IMAGE_CUBIC_WEIGHT_HP                                      Enum = 0x815E
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     Enum = 0x8408
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              Enum = 0x851C
	MAP1_VERTEX_ATTRIB11_4_NV                                  Enum = 0x866B
	DRAW_BUFFER1_NV                                            Enum = 0x8826
	NORMAL_ARRAY_ADDRESS_NV                                    Enum = 0x8F22
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            Enum = 0x20
	MAX_COLOR_TEXTURE_SAMPLES                                  Enum = 0x910E
	SGIX_texture_lod_bias                                      Enum = 1
	UNSIGNED_INT_8_8_8_8                                       Enum = 0x8035
	COLOR_TABLE_INTENSITY_SIZE_SGI                             Enum = 0x80DF
	FRAGMENT_NORMAL_EXT                                        Enum = 0x834A
	CLIENT_ACTIVE_TEXTURE                                      Enum = 0x84E1
	INCR_WRAP                                                  Enum = 0x8507
	VERTEX_ATTRIB_ARRAY15_NV                                   Enum = 0x865F
	VERTEX_ARRAY_ADDRESS_NV                                    Enum = 0x8F21
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          Enum = 0x11
	ACTIVE_RESOURCES                                           Enum = 0x92F5
	NORMALIZE                                                  Enum = 0x0BA1
	OR_INVERTED                                                Enum = 0x150D
	CONVOLUTION_WIDTH                                          Enum = 0x8018
	TEXTURE20                                                  Enum = 0x84D4
	POINT_SPRITE_NV                                            Enum = 0x8861
	TEXTURE_FLOAT_COMPONENTS_NV                                Enum = 0x888C
	MAP1_COLOR_4                                               Enum = 0x0D90
	TEXTURE_2D_BINDING_EXT                                     Enum = 0x8069
	ASYNC_READ_PIXELS_SGIX                                     Enum = 0x835E
	COMPRESSED_RGB_S3TC_DXT1_EXT                               Enum = 0x83F0
	RGB16F_ARB                                                 Enum = 0x881B
	REG_23_ATI                                                 Enum = 0x8938
	ALPHA16F_ARB                                               Enum = 0x881C
	DOT3_ATI                                                   Enum = 0x8966
	PURGEABLE_APPLE                                            Enum = 0x8A1D
	INDEX_TEST_FUNC_EXT                                        Enum = 0x81B6
	TEXTURE_COMPRESSED_BLOCK_SIZE                              Enum = 0x82B3
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         Enum = 0x8C4F
	ATOMIC_COUNTER_BUFFER_START                                Enum = 0x92C2
	POLYGON_SMOOTH                                             Enum = 0x0B41
	VERTEX_ATTRIB_ARRAY_SIZE                                   Enum = 0x8623
	DEPTH_BOUNDS_TEST_EXT                                      Enum = 0x8890
	CONTEXT_PROFILE_MASK                                       Enum = 0x9126
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               Enum = 0x8541
	VERTEX_ATTRIB_ARRAY_ENABLED                                Enum = 0x8622
	ALPHA_FLOAT16_ATI                                          Enum = 0x881C
	DRAW_BUFFER15_ARB                                          Enum = 0x8834
	MAX_DEPTH                                                  Enum = 0x8280
	PRIMITIVES_GENERATED_NV                                    Enum = 0x8C87
	TEXTURE_VIEW_MIN_LAYER                                     Enum = 0x82DD
	REFLECTION_MAP                                             Enum = 0x8512
	RGBA_FLOAT16_APPLE                                         Enum = 0x881A
	DRAW_BUFFER10                                              Enum = 0x882F
	SLIM8U_SGIX                                                Enum = 0x831D
	MATRIX21_ARB                                               Enum = 0x88D5
	COLOR_INDEXES                                              Enum = 0x1603
	RGBA4_EXT                                                  Enum = 0x8056
	POST_COLOR_MATRIX_RED_BIAS_SGI                             Enum = 0x80B8
	PIXEL_TILE_WIDTH_SGIX                                      Enum = 0x8140
	SRC1_COLOR                                                 Enum = 0x88F9
	DEPTH32F_STENCIL8                                          Enum = 0x8CAD
	INT_SAMPLER_2D_ARRAY_EXT                                   Enum = 0x8DCF
	EXT_histogram                                              Enum = 1
	TEXTURE_INTENSITY_SIZE                                     Enum = 0x8061
	COMPILE                                                    Enum = 0x1300
	DEBUG_SOURCE_WINDOW_SYSTEM                                 Enum = 0x8247
	DEBUG_TYPE_MARKER                                          Enum = 0x8268
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          Enum = 0x86A0
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           Enum = 0x8DE8
	SIGNED_HILO_NV                                             Enum = 0x86F9
	SWIZZLE_STQ_ATI                                            Enum = 0x8977
	COLOR_ATTACHMENT14                                         Enum = 0x8CEE
	TRIANGLE_STRIP_ADJACENCY                                   Enum = 0x000D
	TEXTURE_DEFORMATION_BIT_SGIX                               Enum = 0x00000001
	QUAD_ALPHA4_SGIS                                           Enum = 0x811E
	SCALE_BY_TWO_NV                                            Enum = 0x853E
	MATRIX0_ARB                                                Enum = 0x88C0
	UNPACK_COMPRESSED_BLOCK_DEPTH                              Enum = 0x9129
	PROJECTION_STACK_DEPTH                                     Enum = 0x0BA4
	PACK_SKIP_ROWS                                             Enum = 0x0D03
	TEXTURE_COORD_ARRAY_TYPE_EXT                               Enum = 0x8089
	NO_RESET_NOTIFICATION_ARB                                  Enum = 0x8261
	DITHER                                                     Enum = 0x0BD0
	VERTEX_STREAM4_ATI                                         Enum = 0x8770
	VERTEX_SHADER                                              Enum = 0x8B31
	FACTOR_MIN_AMD                                             Enum = 0x901C
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          Enum = 0x80B7
	IGNORE_BORDER_HP                                           Enum = 0x8150
	ATOMIC_COUNTER_BUFFER_INDEX                                Enum = 0x9301
	CCW                                                        Enum = 0x0901
	VIEWPORT_BOUNDS_RANGE                                      Enum = 0x825D
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       Enum = 0x82AE
	HI_SCALE_NV                                                Enum = 0x870E
	CON_23_ATI                                                 Enum = 0x8958
	EVAL_BIT                                                   Enum = 0x00010000
	OP_INDEX_EXT                                               Enum = 0x8782
	VERTEX_SHADER_INVARIANTS_EXT                               Enum = 0x87D1
	DOUBLE_VEC2_EXT                                            Enum = 0x8FFC
	MULTISAMPLE_BIT_EXT                                        Enum = 0x20000000
	TEXTURE_MIN_LOD                                            Enum = 0x813A
	DEPTH_TEXTURE_MODE                                         Enum = 0x884B
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                Enum = 0x8C93
	LIST_MODE                                                  Enum = 0x0B30
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              Enum = 0x851C
	RGBA16F                                                    Enum = 0x881A
	SET                                                        Enum = 0x150F
	CLIP_DISTANCE4                                             Enum = 0x3004
	ACTIVE_PROGRAM                                             Enum = 0x8259
	TEXTURE19_ARB                                              Enum = 0x84D3
	STENCIL_INDEX8_OES                                         Enum = 0x8D48
	PRIMITIVE_RESTART_FIXED_INDEX                              Enum = 0x8D69
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         Enum = 0x8CDA
	PATH_COORD_COUNT_NV                                        Enum = 0x909E
	MAX_CLIP_PLANES                                            Enum = 0x0D32
	UNSIGNED_SHORT_5_5_5_1                                     Enum = 0x8034
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           Enum = 0x8248
	VERTEX_STATE_PROGRAM_NV                                    Enum = 0x8621
	TEXTURE_LO_SIZE_NV                                         Enum = 0x871C
	MAX_PROGRAM_PARAMETERS_ARB                                 Enum = 0x88A9
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              Enum = 0x9067
	TEXTURE_MATERIAL_PARAMETER_EXT                             Enum = 0x8352
	DRAW_BUFFER13_ATI                                          Enum = 0x8832
	RENDERBUFFER_WIDTH_OES                                     Enum = 0x8D42
	LUMINANCE_ALPHA_SNORM                                      Enum = 0x9012
	STENCIL_BACK_OP_VALUE_AMD                                  Enum = 0x874D
	FACTOR_ALPHA_MODULATE_IMG                                  Enum = 0x8C07
	LUMINANCE16_ALPHA16                                        Enum = 0x8048
	RG16F_EXT                                                  Enum = 0x822F
	ENABLE_BIT                                                 Enum = 0x00002000
	ALPHA_TEST_FUNC                                            Enum = 0x0BC1
	INTERNALFORMAT_SHARED_SIZE                                 Enum = 0x8277
	DYNAMIC_DRAW                                               Enum = 0x88E8
	RGB16I                                                     Enum = 0x8D89
	TEXTURE_SWIZZLE_RGBA_EXT                                   Enum = 0x8E46
	SGIS_texture_filter4                                       Enum = 1
	TEXTURE_COORD_ARRAY_COUNT_EXT                              Enum = 0x808B
	OUTPUT_TEXTURE_COORD8_EXT                                  Enum = 0x87A5
	ONE_MINUS_SRC1_COLOR                                       Enum = 0x88FA
	TEXTURE_BINDING_BUFFER_ARB                                 Enum = 0x8C2C
	INTENSITY32UI_EXT                                          Enum = 0x8D73
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         Enum = 0x8DE3
	INVALID_FRAMEBUFFER_OPERATION                              Enum = 0x0506
	POST_CONVOLUTION_ALPHA_BIAS                                Enum = 0x8023
	TEXTURE24                                                  Enum = 0x84D8
	INVARIANT_DATATYPE_EXT                                     Enum = 0x87EB
	RED_INTEGER_EXT                                            Enum = 0x8D94
	PIXEL_MAP_I_TO_A_SIZE                                      Enum = 0x0CB5
	VERTEX_SOURCE_ATI                                          Enum = 0x8774
	MAX_PROGRAM_INSTRUCTIONS_ARB                               Enum = 0x88A1
	EDGE_FLAG_ARRAY_POINTER_EXT                                Enum = 0x8093
	TEXTURE_4D_BINDING_SGIS                                    Enum = 0x814F
	VERTEX_ATTRIB_ARRAY4_NV                                    Enum = 0x8654
	RENDERBUFFER_RED_SIZE_OES                                  Enum = 0x8D50
	RELATIVE_SMALL_CCW_ARC_TO_NV                               Enum = 0x13
	ALREADY_SIGNALED                                           Enum = 0x911A
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             Enum = 0x00000001
	T2F_IUI_V3F_EXT                                            Enum = 0x81B2
	UNSIGNED_INT_24_8                                          Enum = 0x84FA
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          Enum = 0x8C8F
	DOUBLE_MAT3_EXT                                            Enum = 0x8F47
	SMOOTH_QUADRATIC_CURVE_TO_NV                               Enum = 0x0E
	SYNC_CONDITION                                             Enum = 0x9113
	MAX_SAMPLES_IMG                                            Enum = 0x9135
	MAT_COLOR_INDEXES_BIT_PGI                                  Enum = 0x01000000
	LINE_SMOOTH_HINT                                           Enum = 0x0C52
	STENCIL_INDEX                                              Enum = 0x1901
	INDEX_MATERIAL_FACE_EXT                                    Enum = 0x81BA
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     Enum = 0x8E80
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           Enum = 0x902C
	MIN_SPARSE_LEVEL_AMD                                       Enum = 0x919B
	REFERENCED_BY_GEOMETRY_SHADER                              Enum = 0x9309
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        Enum = 0x8D6A
	TRUE                                                       Enum = 1
	MAP1_TEXTURE_COORD_3                                       Enum = 0x0D95
	AND_INVERTED                                               Enum = 0x1504
	MIN_EXT                                                    Enum = 0x8007
	COMPRESSED_RG                                              Enum = 0x8226
	VERTEX_STREAM1_ATI                                         Enum = 0x876D
	COLOR_ATTACHMENT8_NV                                       Enum = 0x8CE8
	PRESENT_DURATION_NV                                        Enum = 0x8E2B
	TEXTURE_SAMPLES                                            Enum = 0x9106
	TEXTURE_4DSIZE_SGIS                                        Enum = 0x8136
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          Enum = 0x845D
	ATTRIB_ARRAY_STRIDE_NV                                     Enum = 0x8624
	OP_EXP_BASE_2_EXT                                          Enum = 0x8791
	CON_11_ATI                                                 Enum = 0x894C
	SKIP_COMPONENTS2_NV                                        Enum = -5
	CIRCULAR_CCW_ARC_TO_NV                                     Enum = 0xF8
	SYNC_FENCE_APPLE                                           Enum = 0x9116
	FOG_BIT                                                    Enum = 0x00000080
	POLYGON_OFFSET_BIAS_EXT                                    Enum = 0x8039
	MATRIX3_NV                                                 Enum = 0x8633
	EVAL_VERTEX_ATTRIB14_NV                                    Enum = 0x86D4
	DEPTH_BOUNDS_EXT                                           Enum = 0x8891
	SHADER_OBJECT_ARB                                          Enum = 0x8B48
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              Enum = 0x8DD6
	IMAGE_CUBE_MAP_ARRAY_EXT                                   Enum = 0x9054
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              Enum = 0x92D9
	EDGE_FLAG                                                  Enum = 0x0B43
	PIXEL_MAP_A_TO_A_SIZE                                      Enum = 0x0CB9
	LINEAR_DETAIL_COLOR_SGIS                                   Enum = 0x8099
	OUTPUT_TEXTURE_COORD29_EXT                                 Enum = 0x87BA
	TEXTURE_FREE_MEMORY_ATI                                    Enum = 0x87FC
	COLOR_ATTACHMENT9_NV                                       Enum = 0x8CE9
	MAX_GEOMETRY_SHADER_INVOCATIONS                            Enum = 0x8E5A
	MULTISAMPLE_BIT_3DFX                                       Enum = 0x20000000
	CONVOLUTION_FILTER_BIAS_EXT                                Enum = 0x8015
	DUAL_LUMINANCE_ALPHA4_SGIS                                 Enum = 0x811C
	TEXTURE26                                                  Enum = 0x84DA
	T2F_N3F_V3F                                                Enum = 0x2A2B
	DOT_PRODUCT_PASS_THROUGH_NV                                Enum = 0x885B
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       Enum = 103086
	QUAD_INTENSITY4_SGIS                                       Enum = 0x8122
	TEXTURE7                                                   Enum = 0x84C7
	MATRIX3_ARB                                                Enum = 0x88C3
	VERTEX_ID_NV                                               Enum = 0x8C7B
	VIDEO_COLOR_CONVERSION_MAX_NV                              Enum = 0x902A
	DEBUG_LOGGED_MESSAGES_AMD                                  Enum = 0x9145
	VERSION_1_3                                                Enum = 1
	POLYGON_OFFSET_FACTOR                                      Enum = 0x8038
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     Enum = 0x889A
	UNIFORM_IS_ROW_MAJOR                                       Enum = 0x8A3E
	FLOAT_32_UNSIGNED_INT_24_8_REV                             Enum = 0x8DAD
	TEXTURE_SWIZZLE_A_EXT                                      Enum = 0x8E45
	TEXTURE_LIGHTING_MODE_HP                                   Enum = 0x8167
	MATRIX2_ARB                                                Enum = 0x88C2
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           Enum = 0x8D44
	TEXTURE25                                                  Enum = 0x84D9
	SOURCE2_ALPHA                                              Enum = 0x858A
	DS_BIAS_NV                                                 Enum = 0x8716
	MAGNITUDE_BIAS_NV                                          Enum = 0x8718
	BUFFER_SIZE_ARB                                            Enum = 0x8764
	OP_LOG_BASE_2_EXT                                          Enum = 0x8792
	MIN_PROGRAM_TEXEL_OFFSET                                   Enum = 0x8904
	NUM_FRAGMENT_CONSTANTS_ATI                                 Enum = 0x896F
	FRAGMENT_SHADER                                            Enum = 0x8B30
	SINGLE_COLOR                                               Enum = 0x81F9
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           Enum = 0x845C
	TEXTURE_RECTANGLE                                          Enum = 0x84F5
	EDGE_FLAG_ARRAY                                            Enum = 0x8079
	FUNC_SUBTRACT                                              Enum = 0x800A
	COVERAGE_AUTOMATIC_NV                                      Enum = 0x8ED7
	UNSIGNED_INT_IMAGE_2D                                      Enum = 0x9063
	ACTIVE_VARIABLES                                           Enum = 0x9305
	BLEND_SRC                                                  Enum = 0x0BE1
	MAX_EVAL_ORDER                                             Enum = 0x0D30
	INDEX_ARRAY_BUFFER_BINDING_ARB                             Enum = 0x8899
	REG_8_ATI                                                  Enum = 0x8929
	FONT_DESCENDER_BIT_NV                                      Enum = 0x00400000
	POINT_SIZE_RANGE                                           Enum = 0x0B12
	DETAIL_TEXTURE_LEVEL_SGIS                                  Enum = 0x809A
	MIRRORED_REPEAT_OES                                        Enum = 0x8370
	SGIS_texture_border_clamp                                  Enum = 1
	FILL                                                       Enum = 0x1B02
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            Enum = 0x9069
	BUFFER_UPDATE_BARRIER_BIT_EXT                              Enum = 0x00000200
	LUMINANCE_FLOAT32_APPLE                                    Enum = 0x8818
	FIXED                                                      Enum = 0x140C
	POINT_SIZE_MIN_SGIS                                        Enum = 0x8126
	EVAL_VERTEX_ATTRIB2_NV                                     Enum = 0x86C8
	EVAL_VERTEX_ATTRIB8_NV                                     Enum = 0x86CE
	Y_EXT                                                      Enum = 0x87D6
	CON_22_ATI                                                 Enum = 0x8957
	COUNTER_TYPE_AMD                                           Enum = 0x8BC0
	PATH_COVER_DEPTH_FUNC_NV                                   Enum = 0x90BF
	MAP1_VERTEX_4                                              Enum = 0x0D98
	RGBA_FLOAT32_APPLE                                         Enum = 0x8814
	SAMPLER_2D_RECT_SHADOW                                     Enum = 0x8B64
	TEXTURE_COORD_NV                                           Enum = 0x8C79
	RGB8I                                                      Enum = 0x8D8F
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         Enum = 0x914E
	VERTEX_ATTRIB_ARRAY12_NV                                   Enum = 0x865C
	NEGATIVE_X_EXT                                             Enum = 0x87D9
	MAX_PROGRAM_ATTRIBS_ARB                                    Enum = 0x88AD
	TEXTURE_RENDERBUFFER_NV                                    Enum = 0x8E55
	COMBINER_MUX_SUM_NV                                        Enum = 0x8547
	TEXTURE_BUFFER_OFFSET                                      Enum = 0x919D
	CON_25_ATI                                                 Enum = 0x895A
	PERFMON_RESULT_AVAILABLE_AMD                               Enum = 0x8BC4
	UNIFORM_BLOCK                                              Enum = 0x92E2
	UNPACK_IMAGE_DEPTH_SGIS                                    Enum = 0x8133
	SRGB                                                       Enum = 0x8C40
	PATH_MITER_LIMIT_NV                                        Enum = 0x907A
	VERSION                                                    Enum = 0x1F02
	ELEMENT_ARRAY_POINTER_APPLE                                Enum = 0x8A0E
	TEXTURE_LUMINANCE_TYPE                                     Enum = 0x8C14
	MEDIUM_INT                                                 Enum = 0x8DF4
	RECT_NV                                                    Enum = 0xF6
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     Enum = 0x82AF
	TEXTURE30                                                  Enum = 0x84DE
	COMBINER_SCALE_NV                                          Enum = 0x8548
	MAX_PROGRAM_TEXEL_OFFSET_NV                                Enum = 0x8905
	VIDEO_COLOR_CONVERSION_MIN_NV                              Enum = 0x902B
	COMPRESSED_RGBA_ASTC_10x8_KHR                              Enum = 0x93BA
	AMBIENT_AND_DIFFUSE                                        Enum = 0x1602
	READ_PIXELS                                                Enum = 0x828C
	CURRENT_FOG_COORDINATE_EXT                                 Enum = 0x8453
	EXPAND_NORMAL_NV                                           Enum = 0x8538
	VERTEX_ATTRIB_ARRAY5_NV                                    Enum = 0x8655
	TEXTURE_STENCIL_SIZE_EXT                                   Enum = 0x88F1
	REG_27_ATI                                                 Enum = 0x893C
	LOWER_LEFT                                                 Enum = 0x8CA1
	DOUBLE_MAT3x2_EXT                                          Enum = 0x8F4B
	PIXEL_MAP_I_TO_G_SIZE                                      Enum = 0x0CB3
	GENERATE_MIPMAP_SGIS                                       Enum = 0x8191
	OPERAND2_RGB_ARB                                           Enum = 0x8592
	OUTPUT_TEXTURE_COORD26_EXT                                 Enum = 0x87B7
	TEXCOORD4_BIT_PGI                                          Enum = 0x80000000
	CLIP_DISTANCE2                                             Enum = 0x3002
	SAMPLE_ALPHA_TO_MASK_EXT                                   Enum = 0x809E
	COLOR_MATRIX_SGI                                           Enum = 0x80B1
	DEBUG_CALLBACK_FUNCTION_ARB                                Enum = 0x8244
	MAX_LAYERS                                                 Enum = 0x8281
	LUMINANCE8                                                 Enum = 0x8040
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           Enum = 0x880B
	QUERY_RESULT_AVAILABLE                                     Enum = 0x8867
	POINT_SIZE_MIN_ARB                                         Enum = 0x8126
	TEXTURE_MIN_LOD_SGIS                                       Enum = 0x813A
	VERTEX_SHADER_INSTRUCTIONS_EXT                             Enum = 0x87CF
	DYNAMIC_DRAW_ARB                                           Enum = 0x88E8
	SAMPLER_3D_OES                                             Enum = 0x8B5F
	DEBUG_TYPE_OTHER                                           Enum = 0x8251
	SUBTRACT                                                   Enum = 0x84E7
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            Enum = 0x851E
	NUM_GENERAL_COMBINERS_NV                                   Enum = 0x854E
	PIXEL_UNPACK_BUFFER_EXT                                    Enum = 0x88EC
	RGBA16                                                     Enum = 0x805B
	GL_1PASS_SGIS                                              Enum = 0x80A1
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               Enum = 0x88FC
	MIN_SAMPLE_SHADING_VALUE_ARB                               Enum = 0x8C37
	RENDERBUFFER_BINDING_ANGLE                                 Enum = 0x8CA7
	DECAL                                                      Enum = 0x2101
	CLAMP_TO_EDGE_SGIS                                         Enum = 0x812F
	MIRRORED_REPEAT                                            Enum = 0x8370
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            Enum = 0x85C1
	VARIANT_EXT                                                Enum = 0x87C1
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        Enum = 0x8B8B
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                Enum = 0x8C75
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER Enum = 0x92C9
	UNIFORM                                                    Enum = 0x92E1
	SGI_color_table                                            Enum = 1
	GREEN_BITS                                                 Enum = 0x0D53
	BLUE                                                       Enum = 0x1905
	R1UI_C3F_V3F_SUN                                           Enum = 0x85C6
	DRAW_BUFFER6_ARB                                           Enum = 0x882B
	PIXEL_MAP_I_TO_A                                           Enum = 0x0C75
	UNPACK_CMYK_HINT_EXT                                       Enum = 0x800F
	DEPTH_COMPONENT16_SGIX                                     Enum = 0x81A5
	X_EXT                                                      Enum = 0x87D5
	COLOR_ATTACHMENT1_EXT                                      Enum = 0x8CE1
	SAMPLER_RENDERBUFFER_NV                                    Enum = 0x8E56
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              Enum = 0x9033
	SECONDARY_COLOR_ARRAY_LIST_IBM                             Enum = 103077
	MAX_DEBUG_GROUP_STACK_DEPTH                                Enum = 0x826C
	INTERPOLATE                                                Enum = 0x8575
	OUTPUT_TEXTURE_COORD9_EXT                                  Enum = 0x87A6
	TIME_ELAPSED                                               Enum = 0x88BF
	TIMEOUT_EXPIRED                                            Enum = 0x911B
	ALL_BARRIER_BITS                                           Enum = 0xFFFFFFFF
	DRAW_INDIRECT_BUFFER                                       Enum = 0x8F3F
	INT_IMAGE_1D_ARRAY                                         Enum = 0x905D
	STENCIL_BITS                                               Enum = 0x0D57
	CLIP_DISTANCE5                                             Enum = 0x3005
	DEBUG_TYPE_POP_GROUP                                       Enum = 0x826A
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            Enum = 0x8510
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         Enum = 0x862E
	FRAGMENT_SHADER_DERIVATIVE_HINT                            Enum = 0x8B8B
	CULL_VERTEX_IBM                                            Enum = 103050
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          Enum = 0x80D2
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           Enum = 0x8214
	PREVIOUS_ARB                                               Enum = 0x8578
	DRAW_BUFFER7_ATI                                           Enum = 0x882C
	INSTRUMENT_BUFFER_POINTER_SGIX                             Enum = 0x8180
	SOURCE0_ALPHA                                              Enum = 0x8588
	VECTOR_EXT                                                 Enum = 0x87BF
	MAX_COMBINED_UNIFORM_BLOCKS                                Enum = 0x8A2E
	TEXTURE_BINDING_BUFFER                                     Enum = 0x8C2C
	DSDT8_MAG8_NV                                              Enum = 0x870A
	RGB5                                                       Enum = 0x8050
	UNPACK_COMPRESSED_SIZE_SGIX                                Enum = 0x831A
	RGBA16F_ARB                                                Enum = 0x881A
	COUNT_UP_NV                                                Enum = 0x9088
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       Enum = 0x93D1
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          Enum = 0x824E
	TEXTURE30_ARB                                              Enum = 0x84DE
	ARRAY_BUFFER_ARB                                           Enum = 0x8892
	GL_8X_BIT_ATI                                              Enum = 0x00000004
	TEXTURE_BUFFER_FORMAT_EXT                                  Enum = 0x8C2E
	COVERAGE_EDGE_FRAGMENTS_NV                                 Enum = 0x8ED6
	PATH_TERMINAL_END_CAP_NV                                   Enum = 0x9078
	RELATIVE_LARGE_CW_ARC_TO_NV                                Enum = 0x19
	SPOT_DIRECTION                                             Enum = 0x1204
	MIN                                                        Enum = 0x8007
	IUI_N3F_V3F_EXT                                            Enum = 0x81B0
	VERTEX_STREAM6_ATI                                         Enum = 0x8772
	FLOAT_R16_NV                                               Enum = 0x8884
	TRANSFORM_FEEDBACK_NV                                      Enum = 0x8E22
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          Enum = 0x8E5E
	PACK_SKIP_IMAGES                                           Enum = 0x806B
	INVALID_OPERATION                                          Enum = 0x0502
	MODELVIEW0_STACK_DEPTH_EXT                                 Enum = 0x0BA3
	EDGE_FLAG_ARRAY_COUNT_EXT                                  Enum = 0x808D
	OUTPUT_TEXTURE_COORD14_EXT                                 Enum = 0x87AB
	TYPE                                                       Enum = 0x92FA
	RGBA_FLOAT16_ATI                                           Enum = 0x881A
	LINK_STATUS                                                Enum = 0x8B82
	SGIS_texture_edge_clamp                                    Enum = 1
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          Enum = 0x82D9
	SLICE_ACCUM_SUN                                            Enum = 0x85CC
	DEBUG_OBJECT_MESA                                          Enum = 0x8759
	PALETTE8_RGB8_OES                                          Enum = 0x8B95
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        Enum = 0x8DDD
	DEBUG_CATEGORY_DEPRECATION_AMD                             Enum = 0x914B
	TRIANGLES_ADJACENCY                                        Enum = 0x000C
	SPRITE_EYE_ALIGNED_SGIX                                    Enum = 0x814E
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         Enum = 0x8409
	MODELVIEW_PROJECTION_NV                                    Enum = 0x8629
	MAP1_VERTEX_ATTRIB1_4_NV                                   Enum = 0x8661
	MAP1_VERTEX_ATTRIB10_4_NV                                  Enum = 0x866A
	PROGRAM_TEX_INSTRUCTIONS_ARB                               Enum = 0x8806
	ACTIVE_UNIFORMS                                            Enum = 0x8B86
	WEIGHT_SUM_UNITY_ARB                                       Enum = 0x86A6
	PATH_GEN_MODE_NV                                           Enum = 0x90B0
	LIGHT4                                                     Enum = 0x4004
	MAD_ATI                                                    Enum = 0x8968
	FRAGMENT_SHADER_ARB                                        Enum = 0x8B30
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          Enum = 0x8FB3
	TEXTURE15                                                  Enum = 0x84CF
	RENDERBUFFER_DEPTH_SIZE_OES                                Enum = 0x8D54
	MAP1_TEXTURE_COORD_1                                       Enum = 0x0D93
	RED_EXT                                                    Enum = 0x1903
	SPRITE_SGIX                                                Enum = 0x8148
	COMPUTE_TEXTURE                                            Enum = 0x82A0
	TEXTURE_WRAP_R_EXT                                         Enum = 0x8072
	FOG_COORD_ARRAY_POINTER                                    Enum = 0x8456
	FENCE_STATUS_NV                                            Enum = 0x84F3
	DEPTH_STENCIL_OES                                          Enum = 0x84F9
	CON_6_ATI                                                  Enum = 0x8947
	SAMPLER_2D_RECT                                            Enum = 0x8B63
	IMAGE_1D_ARRAY                                             Enum = 0x9052
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               Enum = 0x8242
	TEXTURE_COMPRESSION_HINT_ARB                               Enum = 0x84EF
	EVAL_VERTEX_ATTRIB10_NV                                    Enum = 0x86D0
	FLOAT_RGB_NV                                               Enum = 0x8882
	MAX_PROGRAM_LOOP_DEPTH_NV                                  Enum = 0x88F7
	SWIZZLE_STRQ_ATI                                           Enum = 0x897A
	MAX_COLOR_ATTACHMENTS                                      Enum = 0x8CDF
	SGIX_icc_texture                                           Enum = 1
	STENCIL_BUFFER_BIT                                         Enum = 0x00000400
	FOG_INDEX                                                  Enum = 0x0B61
	REFERENCE_PLANE_EQUATION_SGIX                              Enum = 0x817E
	R32I                                                       Enum = 0x8235
	IMAGE_CLASS_2_X_16                                         Enum = 0x82BD
	EYE_RADIAL_NV                                              Enum = 0x855B
	VERTEX_ATTRIB_ARRAY11_NV                                   Enum = 0x865B
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            Enum = 0x88EF
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         Enum = 0x90C8
	LUMINANCE12                                                Enum = 0x8041
	RESET_NOTIFICATION_STRATEGY_ARB                            Enum = 0x8256
	CURRENT_RASTER_COLOR                                       Enum = 0x0B04
	DEPTH_BITS                                                 Enum = 0x0D56
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               Enum = 0x8184
	TRIANGLE_MESH_SUN                                          Enum = 0x8615
	BUFFER_MAP_OFFSET                                          Enum = 0x9121
	UNSIGNED_BYTE_3_3_2                                        Enum = 0x8032
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             Enum = 0x8178
	VIEW_CLASS_RGTC1_RED                                       Enum = 0x82D0
	RENDERBUFFER_BINDING_OES                                   Enum = 0x8CA7
	RENDERBUFFER                                               Enum = 0x8D41
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         Enum = 0x8DA7
	MODELVIEW29_ARB                                            Enum = 0x873D
	TEXTURE_RED_TYPE                                           Enum = 0x8C10
	DEPTH_CLAMP_NEAR_AMD                                       Enum = 0x901E
	EXT_blend_color                                            Enum = 1
	SAMPLER_2D                                                 Enum = 0x8B5E
	READ_FRAMEBUFFER                                           Enum = 0x8CA8
	READ_FRAMEBUFFER_BINDING                                   Enum = 0x8CAA
	NUM_VIDEO_CAPTURE_STREAMS_NV                               Enum = 0x9024
	TEXTURE_2D_MULTISAMPLE                                     Enum = 0x9100
	TEXTURE_1D                                                 Enum = 0x0DE0
	N3F_V3F                                                    Enum = 0x2A25
	DRAW_BUFFER11_ATI                                          Enum = 0x8830
	BYTE                                                       Enum = 0x1400
	VERTEX_ATTRIB_ARRAY10_NV                                   Enum = 0x865A
	ACTIVE_SUBROUTINE_MAX_LENGTH                               Enum = 0x8E48
	PATH_DASH_CAPS_NV                                          Enum = 0x907B
	CURRENT_RASTER_TEXTURE_COORDS                              Enum = 0x0B06
	TEXTURE_LUMINANCE_SIZE                                     Enum = 0x8060
	OUTPUT_TEXTURE_COORD19_EXT                                 Enum = 0x87B0
	TEXTURE_BUFFER_FORMAT_ARB                                  Enum = 0x8C2E
	INT_IMAGE_CUBE_MAP_ARRAY                                   Enum = 0x905F
	TEXTURE_1D_BINDING_EXT                                     Enum = 0x8068
	CONST_EYE_NV                                               Enum = 0x86E5
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             Enum = 0x9199
	COMPRESSED_RGBA_ASTC_8x8_KHR                               Enum = 0x93B7
	EQUAL                                                      Enum = 0x0202
	COLOR_ARRAY                                                Enum = 0x8076
	MODELVIEW0_EXT                                             Enum = 0x1700
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            Enum = 0x889E
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 Enum = 0x8D6C
	FRAGMENT_SHADER_BIT_EXT                                    Enum = 0x00000002
	TABLE_TOO_LARGE                                            Enum = 0x8031
	SAMPLE_MASK_INVERT_EXT                                     Enum = 0x80AB
	MODELVIEW7_ARB                                             Enum = 0x8727
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       Enum = 0x8CD2
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        Enum = 0x00000010
	DEBUG_SOURCE_OTHER                                         Enum = 0x824B
	READ_PIXELS_FORMAT                                         Enum = 0x828D
	VIEW_CLASS_64_BITS                                         Enum = 0x82C6
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              Enum = 0x8DD4
	GPU_OPTIMIZED_QCOM                                         Enum = 0x8FB2
	BLEND_DST_ALPHA_EXT                                        Enum = 0x80CA
	DEPTH24_STENCIL8                                           Enum = 0x88F0
	MUL_ATI                                                    Enum = 0x8964
	COVERAGE_BUFFERS_NV                                        Enum = 0x8ED3
	INT_IMAGE_2D_MULTISAMPLE                                   Enum = 0x9060
	INDEX_ARRAY_TYPE                                           Enum = 0x8085
	GET_TEXTURE_IMAGE_FORMAT                                   Enum = 0x8291
	MAP2_BINORMAL_EXT                                          Enum = 0x8447
	VERTEX_ATTRIB_MAP1_APPLE                                   Enum = 0x8A00
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     Enum = 0x8DA9
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    Enum = 0x10
	FOG_START                                                  Enum = 0x0B63
	FOG_MODE                                                   Enum = 0x0B65
	ACCUM_RED_BITS                                             Enum = 0x0D58
	NICEST                                                     Enum = 0x1102
	MAX_RATIONAL_EVAL_ORDER_NV                                 Enum = 0x86D7
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                Enum = 0x8C73
	ATOMIC_COUNTER_BUFFER_BINDING                              Enum = 0x92C1
	LINE_SMOOTH                                                Enum = 0x0B20
	RENDER                                                     Enum = 0x1C00
	UNSIGNED_INT_10_10_10_2_EXT                                Enum = 0x8036
	COLOR_MATRIX                                               Enum = 0x80B1
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            Enum = 0x8C4B
	RENDERBUFFER_SAMPLES_IMG                                   Enum = 0x9133
	SGIX_impact_pixel_texture                                  Enum = 1
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        Enum = 0x83F7
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         Enum = 0x8DA7
	GL_1PASS_EXT                                               Enum = 0x80A1
	MAX_ELEMENTS_VERTICES                                      Enum = 0x80E8
	CURRENT_MATRIX_NV                                          Enum = 0x8641
	BUFFER_MAPPED_OES                                          Enum = 0x88BC
	SYNC_OBJECT_APPLE                                          Enum = 0x8A53
	MAX_TEXTURE_BUFFER_SIZE_ARB                                Enum = 0x8C2B
	SRGB8_EXT                                                  Enum = 0x8C41
	OPERAND3_ALPHA_NV                                          Enum = 0x859B
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           Enum = 0x8803
	INVALID_INDEX                                              Enum = 0xFFFFFFFF
	SAMPLER_CUBE_SHADOW_NV                                     Enum = 0x8DC5
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        Enum = 0x903A
	PROGRAM_BINARY_LENGTH                                      Enum = 0x8741
	DRAW_BUFFER9_ATI                                           Enum = 0x882E
	TESS_EVALUATION_PROGRAM_NV                                 Enum = 0x891F
	DRAW_FRAMEBUFFER_BINDING                                   Enum = 0x8CA6
	TESS_CONTROL_SHADER                                        Enum = 0x8E88
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          Enum = 0x885A
	RGBA_INTEGER                                               Enum = 0x8D99
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      Enum = 0x93D8
	VIEW_CLASS_32_BITS                                         Enum = 0x82C8
	FOG_COORDINATE_ARRAY                                       Enum = 0x8457
	PN_TRIANGLES_POINT_MODE_ATI                                Enum = 0x87F2
	SAMPLER_CUBE_ARB                                           Enum = 0x8B60
	ELEMENT_ARRAY_BUFFER                                       Enum = 0x8893
	T2F_C3F_V3F                                                Enum = 0x2A2A
	REDUCE                                                     Enum = 0x8016
	PROGRAM_BINARY_LENGTH_OES                                  Enum = 0x8741
	DRAW_BUFFER8_ARB                                           Enum = 0x882D
	SAMPLER_2D_SHADOW_ARB                                      Enum = 0x8B62
	ACTIVE_VARYING_MAX_LENGTH_NV                               Enum = 0x8C82
	SHADER_STORAGE_BARRIER_BIT                                 Enum = 0x00002000
	NEAREST_CLIPMAP_NEAREST_SGIX                               Enum = 0x844D
	REG_16_ATI                                                 Enum = 0x8931
	POINT_SPRITE_COORD_ORIGIN                                  Enum = 0x8CA0
	PROGRAM_MATRIX_EXT                                         Enum = 0x8E2D
	TEXTURE_DEPTH_EXT                                          Enum = 0x8071
	INTERNALFORMAT_BLUE_SIZE                                   Enum = 0x8273
	SHADER                                                     Enum = 0x82E1
	SOURCE0_RGB                                                Enum = 0x8580
	SURFACE_STATE_NV                                           Enum = 0x86EB
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       Enum = 0x8E49
	FRAMEBUFFER_DEFAULT_SAMPLES                                Enum = 0x9313
	LIGHTING_BIT                                               Enum = 0x00000040
	PROXY_HISTOGRAM_EXT                                        Enum = 0x8025
	FRAMEBUFFER_UNDEFINED_OES                                  Enum = 0x8219
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         Enum = 0x8314
	NORMAL_MAP                                                 Enum = 0x8511
	COMBINE_ARB                                                Enum = 0x8570
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       Enum = 0x8B87
	SRC_ALPHA_SATURATE                                         Enum = 0x0308
	SAMPLER                                                    Enum = 0x82E6
	MAX_CUBE_MAP_TEXTURE_SIZE                                  Enum = 0x851C
	MODELVIEW23_ARB                                            Enum = 0x8737
	UNSIGNED_INT_8_24_REV_MESA                                 Enum = 0x8752
	DOUBLE_MAT3x4_EXT                                          Enum = 0x8F4C
	FOG_COLOR                                                  Enum = 0x0B66
	QUAD_TEXTURE_SELECT_SGIS                                   Enum = 0x8125
	RG8I                                                       Enum = 0x8237
	SGX_BINARY_IMG                                             Enum = 0x8C0A
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              Enum = 0x8CD7
	TRANSFORM_FEEDBACK_BINDING                                 Enum = 0x8E25
	DEBUG_SEVERITY_LOW                                         Enum = 0x9148
	POLYGON_BIT                                                Enum = 0x00000008
	PACK_SKIP_IMAGES_EXT                                       Enum = 0x806B
	OCCLUSION_TEST_RESULT_HP                                   Enum = 0x8166
	DEFORMATIONS_MASK_SGIX                                     Enum = 0x8196
	SIGNED_RGB_NV                                              Enum = 0x86FE
	LINE_QUALITY_HINT_SGIX                                     Enum = 0x835B
	TESS_GEN_SPACING                                           Enum = 0x8E77
	IMAGE_TRANSLATE_Y_HP                                       Enum = 0x8158
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              Enum = 0x83F2
	IMAGE_2D_RECT                                              Enum = 0x904F
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           Enum = 0x9279
	TRANSFORM_FEEDBACK_BARRIER_BIT                             Enum = 0x00000800
	TEXTURE_BINDING_RECTANGLE_ARB                              Enum = 0x84F6
	STENCIL_OP_VALUE_AMD                                       Enum = 0x874C
	DRAW_BUFFER8                                               Enum = 0x882D
	TEXTURE_VIEW                                               Enum = 0x82B5
	RENDERBUFFER_GREEN_SIZE_OES                                Enum = 0x8D51
	MODELVIEW3_ARB                                             Enum = 0x8723
	DEPTH_COMPONENT32F                                         Enum = 0x8CAC
	TRIANGULAR_NV                                              Enum = 0x90A5
	ONE_MINUS_CONSTANT_ALPHA                                   Enum = 0x8004
	SIGNED_LUMINANCE8_NV                                       Enum = 0x8702
	QUERY_RESULT_ARB                                           Enum = 0x8866
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           Enum = 0x88F4
	COMPRESSED_RGBA_ARB                                        Enum = 0x84EE
	COMBINER4_NV                                               Enum = 0x8554
	CURRENT_VERTEX_EXT                                         Enum = 0x87E2
	FRAMEZOOM_SGIX                                             Enum = 0x818B
	MAX_COMBINED_DIMENSIONS                                    Enum = 0x8282
	DRAW_BUFFER15_NV                                           Enum = 0x8834
	DOUBLE_VEC4                                                Enum = 0x8FFE
	TEXTURE23                                                  Enum = 0x84D7
	PROGRAM_POINT_SIZE                                         Enum = 0x8642
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       Enum = 0x8C29
	RENDERBUFFER_HEIGHT                                        Enum = 0x8D43
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   Enum = 0x8DE1
	TRANSLATE_X_NV                                             Enum = 0x908E
	FRAMEBUFFER_BARRIER_BIT_EXT                                Enum = 0x00000400
	MAJOR_VERSION                                              Enum = 0x821B
	IMAGE_CLASS_4_X_16                                         Enum = 0x82BC
	TEXTURE_DS_SIZE_NV                                         Enum = 0x871D
	FRAMEBUFFER_BINDING_ANGLE                                  Enum = 0x8CA6
	UNSIGNED_INT_SAMPLER_2D_EXT                                Enum = 0x8DD2
	CLIENT_VERTEX_ARRAY_BIT                                    Enum = 0x00000002
	TEXTURE_LUMINANCE_SIZE_EXT                                 Enum = 0x8060
	R16I                                                       Enum = 0x8233
	VERTEX_PROGRAM_TWO_SIDE_ARB                                Enum = 0x8643
	PIXEL_UNPACK_BUFFER                                        Enum = 0x88EC
	MAX_PROGRAM_CALL_DEPTH_NV                                  Enum = 0x88F5
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             Enum = 0x8F20
	ALPHA4                                                     Enum = 0x803B
	TEXTURE0                                                   Enum = 0x84C0
	TEXTURE19                                                  Enum = 0x84D3
	TRANSPOSE_TEXTURE_MATRIX                                   Enum = 0x84E5
	REPLACE_VALUE_AMD                                          Enum = 0x874B
	DRAW_BUFFER11_ARB                                          Enum = 0x8830
	CURRENT_QUERY_ARB                                          Enum = 0x8865
	TEXTURE_BUFFER                                             Enum = 0x8C2A
	MAX_COLOR_ATTACHMENTS_NV                                   Enum = 0x8CDF
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           Enum = 0x8DE0
	LINEAR                                                     Enum = 0x2601
	FILTER                                                     Enum = 0x829A
	CURRENT_TIME_NV                                            Enum = 0x8E28
	BLUE_SCALE                                                 Enum = 0x0D1A
	RGBA8                                                      Enum = 0x8058
	T2F_C4F_N3F_V3F                                            Enum = 0x2A2C
	CONTINUOUS_AMD                                             Enum = 0x9007
	PATH_ERROR_POSITION_NV                                     Enum = 0x90AB
	BACK                                                       Enum = 0x0405
	TEXTURE_3D_BINDING_OES                                     Enum = 0x806A
	UNSIGNED_INT_S8_S8_8_8_NV                                  Enum = 0x86DA
	MAX_IMAGE_SAMPLES_EXT                                      Enum = 0x906D
	RGB32F                                                     Enum = 0x8815
	QUERY_RESULT_EXT                                           Enum = 0x8866
	COLOR_ATTACHMENT4_EXT                                      Enum = 0x8CE4
	SGIS_pixel_texture                                         Enum = 1
	SAMPLE_BUFFERS                                             Enum = 0x80A8
	LAYER_PROVOKING_VERTEX                                     Enum = 0x825E
	R1UI_T2F_V3F_SUN                                           Enum = 0x85C9
	MAP1_VERTEX_ATTRIB5_4_NV                                   Enum = 0x8665
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       Enum = 0x93D0
	CURRENT_RASTER_POSITION_VALID                              Enum = 0x0B08
	IMAGE_SCALE_Y_HP                                           Enum = 0x8156
	COMPRESSED_RED_RGTC1_EXT                                   Enum = 0x8DBB
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            Enum = 0x8DD8
	COVERAGE_COMPONENT_NV                                      Enum = 0x8ED0
	INT_IMAGE_CUBE                                             Enum = 0x905B
	RGB5_EXT                                                   Enum = 0x8050
	POINT_SIZE_ARRAY_STRIDE_OES                                Enum = 0x898B
	SHORT                                                      Enum = 0x1402
	TEXTURE_BLUE_SIZE_EXT                                      Enum = 0x805E
	DEPTH_COMPONENT16_OES                                      Enum = 0x81A5
	MODULATE_SUBTRACT_ATI                                      Enum = 0x8746
	UNSIGNED_INT_IMAGE_2D_RECT                                 Enum = 0x9065
	IS_PER_PATCH                                               Enum = 0x92E7
	DOMAIN                                                     Enum = 0x0A02
	Q                                                          Enum = 0x2003
	TEXTURE21                                                  Enum = 0x84D5
	DRAW_BUFFER13                                              Enum = 0x8832
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      Enum = 0x8C76
	DRAW_FRAMEBUFFER_ANGLE                                     Enum = 0x8CA9
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        Enum = 0x90D0
	TRIANGLE_LIST_SUN                                          Enum = 0x81D7
	READ_FRAMEBUFFER_BINDING_NV                                Enum = 0x8CAA
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        Enum = 0x8DDF
	RELATIVE_ARC_TO_NV                                         Enum = 0xFF
	EXT_texture                                                Enum = 1
	INVERT                                                     Enum = 0x150A
	SAMPLE_MASK_SGIS                                           Enum = 0x80A0
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             Enum = 0x824D
	FOG_COORDINATE_ARRAY_EXT                                   Enum = 0x8457
	TEXTURE25_ARB                                              Enum = 0x84D9
	DOUBLE_MAT2                                                Enum = 0x8F46
	CLOSE_PATH_NV                                              Enum = 0x00
	OBJECT_LINEAR                                              Enum = 0x2401
	INTENSITY16                                                Enum = 0x804D
	PIXEL_SUBSAMPLE_4444_SGIX                                  Enum = 0x85A2
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      Enum = 0x8C2D
	ALPHA16UI_EXT                                              Enum = 0x8D78
	PATH_END_CAPS_NV                                           Enum = 0x9076
	WAIT_FAILED                                                Enum = 0x911D
	CULL_FACE_MODE                                             Enum = 0x0B45
	DRAW_BUFFER12_NV                                           Enum = 0x8831
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       Enum = 0x8E5E
	SPARE0_PLUS_SECONDARY_COLOR_NV                             Enum = 0x8532
	INVARIANT_EXT                                              Enum = 0x87C2
	RELATIVE_CUBIC_CURVE_TO_NV                                 Enum = 0x0D
	COLOR3_BIT_PGI                                             Enum = 0x00010000
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            Enum = 0x84F0
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       Enum = 0x889C
	REG_6_ATI                                                  Enum = 0x8927
	NEXT_BUFFER_NV                                             Enum = -2
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        Enum = 0x8E24
	INTENSITY16_SNORM                                          Enum = 0x901B
	LARGE_CCW_ARC_TO_NV                                        Enum = 0x16
	GEQUAL                                                     Enum = 0x0206
	EDGE_FLAG_ARRAY_POINTER                                    Enum = 0x8093
	TEXTURE_FILTER4_SIZE_SGIS                                  Enum = 0x8147
	TRANSPOSE_TEXTURE_MATRIX_ARB                               Enum = 0x84E5
	MAX_UNIFORM_BLOCK_SIZE                                     Enum = 0x8A30
	TEXTURE_SRGB_DECODE_EXT                                    Enum = 0x8A48
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          Enum = 0x80B5
	SGIX_ycrcb                                                 Enum = 1
	UNSIGNED_INT64_NV                                          Enum = 0x140F
	IMAGE_CLASS_4_X_32                                         Enum = 0x82B9
	TEXTURE_BINDING_CUBE_MAP                                   Enum = 0x8514
	OUTPUT_TEXTURE_COORD3_EXT                                  Enum = 0x87A0
	OUTPUT_TEXTURE_COORD24_EXT                                 Enum = 0x87B5
	IMAGE_TRANSLATE_X_HP                                       Enum = 0x8157
	COLOR_RENDERABLE                                           Enum = 0x8286
	SAMPLER_1D                                                 Enum = 0x8B5D
	COMPRESSED_RGBA8_ETC2_EAC                                  Enum = 0x9278
	VERTEX_CONSISTENT_HINT_PGI                                 Enum = 0x1A22B
	IMAGE_MIN_FILTER_HP                                        Enum = 0x815D
	MAX_RENDERBUFFER_SIZE_OES                                  Enum = 0x84E8
	INT16_VEC4_NV                                              Enum = 0x8FE7
	RELATIVE_SMALL_CW_ARC_TO_NV                                Enum = 0x15
	SYNC_FENCE                                                 Enum = 0x9116
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            Enum = 0x919A
	TEXTURE12_ARB                                              Enum = 0x84CC
	VERTEX_STREAM0_ATI                                         Enum = 0x876C
	FRAGMENT_ALPHA_MODULATE_IMG                                Enum = 0x8C08
	RENDERBUFFER_SAMPLES_NV                                    Enum = 0x8CAB
	ALPHA_SNORM                                                Enum = 0x9010
	INTENSITY_EXT                                              Enum = 0x8049
	POINT_FADE_THRESHOLD_SIZE                                  Enum = 0x8128
	CONSTANT_BORDER_HP                                         Enum = 0x8151
	TEXTURE29                                                  Enum = 0x84DD
	MODELVIEW6_ARB                                             Enum = 0x8726
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       Enum = 0x8E81
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         Enum = 0x90ED
	FRAGMENT_LIGHT4_SGIX                                       Enum = 0x8410
	DRAW_BUFFER5_NV                                            Enum = 0x882A
	REG_21_ATI                                                 Enum = 0x8936
	INDEX_ARRAY_EXT                                            Enum = 0x8077
	TEXTURE_COORD_ARRAY_EXT                                    Enum = 0x8078
	MINOR_VERSION                                              Enum = 0x821C
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              Enum = 0x8315
	TEXTURE_COMPRESSED_IMAGE_SIZE                              Enum = 0x86A0
	UNPACK_RESAMPLE_OML                                        Enum = 0x8985
	IMAGE_MAG_FILTER_HP                                        Enum = 0x815C
	SOURCE1_RGB                                                Enum = 0x8581
	DRAW_FRAMEBUFFER_BINDING_EXT                               Enum = 0x8CA6
	COLOR_ATTACHMENT10                                         Enum = 0x8CEA
	SAMPLE_MASK_VALUE                                          Enum = 0x8E52
	LINE_STRIP_ADJACENCY                                       Enum = 0x000B
	COLOR_ARRAY_STRIDE_EXT                                     Enum = 0x8083
	TEXTURE_SHADOW                                             Enum = 0x82A1
	OBJECT_BUFFER_SIZE_ATI                                     Enum = 0x8764
	RGBA_SIGNED_COMPONENTS_EXT                                 Enum = 0x8C3C
	COLOR_ATTACHMENT5_NV                                       Enum = 0x8CE5
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         Enum = 0x90DA
	HALF_APPLE                                                 Enum = 0x140B
	UNPACK_IMAGE_HEIGHT_EXT                                    Enum = 0x806E
	INTERNALFORMAT_DEPTH_SIZE                                  Enum = 0x8275
	REG_18_ATI                                                 Enum = 0x8933
	SIGNALED_APPLE                                             Enum = 0x9119
	REPEAT                                                     Enum = 0x2901
	RGB10_A2_EXT                                               Enum = 0x8059
	POINT_DISTANCE_ATTENUATION_ARB                             Enum = 0x8129
	INTERNALFORMAT_RED_TYPE                                    Enum = 0x8278
	SAMPLER_2D_RECT_SHADOW_ARB                                 Enum = 0x8B64
	SCALED_RESOLVE_NICEST_EXT                                  Enum = 0x90BB
	PACK_SKIP_PIXELS                                           Enum = 0x0D04
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             Enum = 0x813E
	MAX_TEXTURE_UNITS                                          Enum = 0x84E2
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       Enum = 0x86D9
	PROGRAM_ATTRIB_COMPONENTS_NV                               Enum = 0x8906
	TEXTURE_HEIGHT_QCOM                                        Enum = 0x8BD3
	VIEW_CLASS_128_BITS                                        Enum = 0x82C4
	SAMPLES_PASSED_ARB                                         Enum = 0x8914
	COMPRESSED_SIGNED_RED_RGTC1                                Enum = 0x8DBC
	AUTO_NORMAL                                                Enum = 0x0D80
	SMOOTH                                                     Enum = 0x1D01
	BUFFER                                                     Enum = 0x82E0
	MODELVIEW1_STACK_DEPTH_EXT                                 Enum = 0x8502
	UNSIGNED_SHORT_8_8_REV_APPLE                               Enum = 0x85BB
	OFFSET_TEXTURE_SCALE_NV                                    Enum = 0x86E2
	OBJECT_VALIDATE_STATUS_ARB                                 Enum = 0x8B83
	MAP_READ_BIT_EXT                                           Enum = 0x0001
	TEXTURE_VIEW_NUM_LEVELS                                    Enum = 0x82DC
	COMBINER5_NV                                               Enum = 0x8555
	SURFACE_MAPPED_NV                                          Enum = 0x8700
	MAX_SAMPLES_NV                                             Enum = 0x8D57
	UNSIGNED_INT_SAMPLER_1D_EXT                                Enum = 0x8DD1
	MAP1_BINORMAL_EXT                                          Enum = 0x8446
	SAMPLE_POSITION_NV                                         Enum = 0x8E50
	BLOCK_INDEX                                                Enum = 0x92FD
	SRC_COLOR                                                  Enum = 0x0300
	INVALID_FRAMEBUFFER_OPERATION_OES                          Enum = 0x0506
	TEXTURE31                                                  Enum = 0x84DF
	UNPACK_SUBSAMPLE_RATE_SGIX                                 Enum = 0x85A1
	INTENSITY_FLOAT32_APPLE                                    Enum = 0x8817
	SKIP_DECODE_EXT                                            Enum = 0x8A4A
	TEXTURE_NUM_LEVELS_QCOM                                    Enum = 0x8BD9
	RENDERBUFFER_STENCIL_SIZE_EXT                              Enum = 0x8D55
	INTENSITY16I_EXT                                           Enum = 0x8D8B
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            Enum = 0x8519
	SHADER_CONSISTENT_NV                                       Enum = 0x86DD
	PROGRAM_ALU_INSTRUCTIONS_ARB                               Enum = 0x8805
	SAMPLER_1D_SHADOW                                          Enum = 0x8B61
	TEXTURE_ALPHA_SIZE_EXT                                     Enum = 0x805F
	TEXTURE_3D_EXT                                             Enum = 0x806F
	OPERAND1_ALPHA_ARB                                         Enum = 0x8599
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               Enum = 0x870C
	FONT_HAS_KERNING_BIT_NV                                    Enum = 0x10000000
	ONE_MINUS_DST_ALPHA                                        Enum = 0x0305
	RG16F                                                      Enum = 0x822F
	STENCIL_INDEX4_OES                                         Enum = 0x8D47
	ANY_SAMPLES_PASSED_CONSERVATIVE                            Enum = 0x8D6A
	SAMPLER_CUBE_SHADOW_EXT                                    Enum = 0x8DC5
	RGB10_EXT                                                  Enum = 0x8052
	SLUMINANCE8_ALPHA8_NV                                      Enum = 0x8C45
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            Enum = 0x902D
	FONT_ASCENDER_BIT_NV                                       Enum = 0x00200000
	LOGIC_OP                                                   Enum = 0x0BF1
	SAMPLE_BUFFERS_ARB                                         Enum = 0x80A8
	PIXEL_TILE_CACHE_SIZE_SGIX                                 Enum = 0x8145
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 Enum = 0x8185
	DEPTH_COMPONENT24_SGIX                                     Enum = 0x81A6
	DEBUG_SEVERITY_NOTIFICATION                                Enum = 0x826B
	MAP1_TANGENT_EXT                                           Enum = 0x8444
	DEPTH_STENCIL_EXT                                          Enum = 0x84F9
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            Enum = 0x8518
	SAMPLER_1D_ARB                                             Enum = 0x8B5D
	POST_CONVOLUTION_GREEN_SCALE                               Enum = 0x801D
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    Enum = 0x80D5
	TANGENT_ARRAY_TYPE_EXT                                     Enum = 0x843E
	INT_SAMPLER_2D_RECT_EXT                                    Enum = 0x8DCD
	FRAMEBUFFER_DEFAULT_LAYERS                                 Enum = 0x9312
	OBJECT_POINT_SGIS                                          Enum = 0x81F5
	DEPTH_STENCIL                                              Enum = 0x84F9
	MATRIX18_ARB                                               Enum = 0x88D2
	RGB_INTEGER_EXT                                            Enum = 0x8D98
	SOURCE0_RGB_ARB                                            Enum = 0x8580
	MAX_VERTEX_ATTRIBS                                         Enum = 0x8869
	REG_5_ATI                                                  Enum = 0x8926
	GL_422_EXT                                                 Enum = 0x80CC
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            Enum = 0x87EE
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   Enum = 0x8A33
	RENDERBUFFER_INTERNAL_FORMAT                               Enum = 0x8D44
	RGB32I                                                     Enum = 0x8D83
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            Enum = 0x8DDF
	SAMPLER_BUFFER_AMD                                         Enum = 0x9001
	LEFT                                                       Enum = 0x0406
	DEBUG_TYPE_PORTABILITY_ARB                                 Enum = 0x824F
	VIEW_CLASS_S3TC_DXT1_RGB                                   Enum = 0x82CC
	INTERPOLATE_EXT                                            Enum = 0x8575
	SAMPLES_PASSED                                             Enum = 0x8914
	SHADER_INCLUDE_ARB                                         Enum = 0x8DAE
	RECLAIM_MEMORY_HINT_PGI                                    Enum = 0x1A1FE
	TESS_EVALUATION_SHADER_BIT                                 Enum = 0x00000010
	SAMPLE_COVERAGE_INVERT_ARB                                 Enum = 0x80AB
	DEPTH_COMPONENT16_ARB                                      Enum = 0x81A5
	VERTEX_PROGRAM_TWO_SIDE_NV                                 Enum = 0x8643
	MODELVIEW20_ARB                                            Enum = 0x8734
	MATRIX_INDEX_ARRAY_TYPE_OES                                Enum = 0x8847
	ELEMENT_ARRAY_BUFFER_BINDING                               Enum = 0x8895
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            Enum = 0x00200000
	SGIX_scalebias_hint                                        Enum = 1
	CLIP_DISTANCE6                                             Enum = 0x3006
	TEXTURE_LOD_BIAS_T_SGIX                                    Enum = 0x818F
	BOOL_VEC3_ARB                                              Enum = 0x8B58
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              Enum = 0x01000000
	UNSIGNED_INT_24_8_OES                                      Enum = 0x84FA
	DYNAMIC_ATI                                                Enum = 0x8761
	LUMINANCE_ALPHA8UI_EXT                                     Enum = 0x8D81
	POST_CONVOLUTION_ALPHA_SCALE                               Enum = 0x801F
	INTENSITY4_EXT                                             Enum = 0x804A
	TEXTURE_CUBE_MAP_NEGATIVE_X                                Enum = 0x8516
	COLOR_ARRAY_BUFFER_BINDING                                 Enum = 0x8898
	DOT4_ATI                                                   Enum = 0x8967
	INTERLEAVED_ATTRIBS_EXT                                    Enum = 0x8C8C
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      Enum = 0x8C8F
	STENCIL_INDEX1_EXT                                         Enum = 0x8D46
	PACK_ROW_LENGTH                                            Enum = 0x0D02
	PIXEL_MAG_FILTER_EXT                                       Enum = 0x8331
	UNPACK_RESAMPLE_SGIX                                       Enum = 0x842D
	OPERAND0_ALPHA_ARB                                         Enum = 0x8598
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               Enum = 0x9038
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             Enum = 0x9107
	BLEND_COLOR_EXT                                            Enum = 0x8005
	TEXTURE22                                                  Enum = 0x84D6
	FLOAT_MAT3x4                                               Enum = 0x8B68
	PROXY_TEXTURE_1D_ARRAY                                     Enum = 0x8C19
	EXT_blend_logic_op                                         Enum = 1
	PIXEL_MAP_B_TO_B_SIZE                                      Enum = 0x0CB8
	VARIABLE_D_NV                                              Enum = 0x8526
	CURRENT_OCCLUSION_QUERY_ID_NV                              Enum = 0x8865
	FORMAT_SUBSAMPLE_244_244_OML                               Enum = 0x8983
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        Enum = 0x8C4D
	COLOR_SAMPLES_NV                                           Enum = 0x8E20
	HISTOGRAM_FORMAT_EXT                                       Enum = 0x8027
	MAX_ELEMENTS_INDICES                                       Enum = 0x80E9
	VERTEX_ATTRIB_ARRAY6_NV                                    Enum = 0x8656
	OFFSET_TEXTURE_2D_NV                                       Enum = 0x86E8
	BLUE_INTEGER_EXT                                           Enum = 0x8D96
	UNIFORM_BUFFER_BINDING_EXT                                 Enum = 0x8DEF
	MAX_FRAMEBUFFER_SAMPLES                                    Enum = 0x9318
	DEBUG_CALLBACK_USER_PARAM                                  Enum = 0x8245
	NEVER                                                      Enum = 0x0200
	AUX_BUFFERS                                                Enum = 0x0C00
	UNSIGNED_INT8_VEC4_NV                                      Enum = 0x8FEF
	TESS_CONTROL_SUBROUTINE                                    Enum = 0x92E9
	LINE_LOOP                                                  Enum = 0x0002
	ONE_MINUS_CONSTANT_COLOR                                   Enum = 0x8002
	SIGNED_INTENSITY8_NV                                       Enum = 0x8708
	UNSIGNED_INT_IMAGE_3D_EXT                                  Enum = 0x9064
	MAX_FRAGMENT_INPUT_COMPONENTS                              Enum = 0x9125
	SEPARATE_SPECULAR_COLOR                                    Enum = 0x81FA
	CLIENT_ACTIVE_TEXTURE_ARB                                  Enum = 0x84E1
	ALPHA_INTEGER_EXT                                          Enum = 0x8D97
	UNPACK_COMPRESSED_BLOCK_WIDTH                              Enum = 0x9127
	YCRCB_444_SGIX                                             Enum = 0x81BC
	COLOR_ATTACHMENT0_OES                                      Enum = 0x8CE0
	TRIANGLE_STRIP_ADJACENCY_EXT                               Enum = 0x000D
	RGB                                                        Enum = 0x1907
	POST_CONVOLUTION_RED_SCALE_EXT                             Enum = 0x801C
	HISTOGRAM_ALPHA_SIZE_EXT                                   Enum = 0x802B
	GENERATE_MIPMAP_HINT_SGIS                                  Enum = 0x8192
	SHADER_OBJECT_EXT                                          Enum = 0x8B48
	MAX_PATCH_VERTICES                                         Enum = 0x8E7D
	INTERLACE_SGIX                                             Enum = 0x8094
	DUAL_LUMINANCE12_SGIS                                      Enum = 0x8116
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           Enum = 0x8243
	IMAGE_CLASS_11_11_10                                       Enum = 0x82C2
	TEXTURE0_ARB                                               Enum = 0x84C0
	ACTIVE_TEXTURE_ARB                                         Enum = 0x84E0
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              Enum = 0x824E
	SET_AMD                                                    Enum = 0x874A
	ELEMENT_ARRAY_APPLE                                        Enum = 0x8A0C
	BLUE_INTEGER                                               Enum = 0x8D96
	SGIX_polynomial_ffd                                        Enum = 1
	T4F_V4F                                                    Enum = 0x2A28
	EVAL_VERTEX_ATTRIB12_NV                                    Enum = 0x86D2
	RENDERBUFFER_SAMPLES_EXT                                   Enum = 0x8CAB
	SGIX_shadow                                                Enum = 1
	ALPHA8                                                     Enum = 0x803C
	COMBINER1_NV                                               Enum = 0x8551
	UNSIGNED_INT_SAMPLER_CUBE                                  Enum = 0x8DD4
	INT_SAMPLER_CUBE_MAP_ARRAY                                 Enum = 0x900E
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         Enum = 0x90CC
	IMAGE_PIXEL_TYPE                                           Enum = 0x82AA
	MODELVIEW5_ARB                                             Enum = 0x8725
	DEBUG_PRINT_MESA                                           Enum = 0x875A
	SYNC_STATUS                                                Enum = 0x9114
	NUM_ACTIVE_VARIABLES                                       Enum = 0x9304
	DRAW_BUFFER0_ARB                                           Enum = 0x8825
	MATRIX17_ARB                                               Enum = 0x88D1
	ACTIVE_SUBROUTINES                                         Enum = 0x8DE5
	VARIABLE_E_NV                                              Enum = 0x8527
	FLOAT_VEC2                                                 Enum = 0x8B50
	FLOAT_VEC2_ARB                                             Enum = 0x8B50
	PROJECTION_MATRIX                                          Enum = 0x0BA7
	CONSTANT_ALPHA                                             Enum = 0x8003
	RGB5_A1_EXT                                                Enum = 0x8057
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               Enum = 0x81A9
	MAX_ASYNC_TEX_IMAGE_SGIX                                   Enum = 0x835F
	RELATIVE_VERTICAL_LINE_TO_NV                               Enum = 0x09
	RGBA4_OES                                                  Enum = 0x8056
	TEXTURE_MAX_LOD                                            Enum = 0x813B
	INDEX_MATERIAL_EXT                                         Enum = 0x81B8
	COMPRESSED_RGB_ARB                                         Enum = 0x84ED
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            Enum = 0x8516
	PACK_SUBSAMPLE_RATE_SGIX                                   Enum = 0x85A0
	TEXTURE_DEPTH_SIZE_ARB                                     Enum = 0x884A
	COPY_WRITE_BUFFER_BINDING                                  Enum = 0x8F37
	UNSIGNED_INT64_VEC4_NV                                     Enum = 0x8FF7
	PROXY_TEXTURE_1D_EXT                                       Enum = 0x8063
	DUAL_LUMINANCE_ALPHA8_SGIS                                 Enum = 0x811D
	SAMPLE_MASK_EXT                                            Enum = 0x80A0
	INDEX_TEST_EXT                                             Enum = 0x81B5
	IMAGE_CLASS_4_X_8                                          Enum = 0x82BF
	BUMP_TARGET_ATI                                            Enum = 0x877C
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  Enum = 0x910D
	UNIFORM_BUFFER                                             Enum = 0x8A11
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        Enum = 0x92D0
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         Enum = 0x8C4D
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               Enum = 0x8E4C
	QUADRATIC_ATTENUATION                                      Enum = 0x1209
	INT64_NV                                                   Enum = 0x140E
	TEXTURE_CLIPMAP_CENTER_SGIX                                Enum = 0x8171
	FOG_COORDINATE_ARRAY_STRIDE                                Enum = 0x8455
	UNIFORM_SIZE                                               Enum = 0x8A38
	FRAMEBUFFER_COMPLETE                                       Enum = 0x8CD5
	SAMPLER_2D_MULTISAMPLE                                     Enum = 0x9108
	COLOR_INDEX                                                Enum = 0x1900
	MAX_VARYING_COMPONENTS                                     Enum = 0x8B4B
	MULTIVIEW_EXT                                              Enum = 0x90F1
	RG32F                                                      Enum = 0x8230
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        Enum = 0x8B8B
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      Enum = 0x8CD9
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        Enum = 0x8E23
	POST_COLOR_MATRIX_ALPHA_SCALE                              Enum = 0x80B7
	COLOR_TABLE_BIAS_SGI                                       Enum = 0x80D7
	MAX_TEXTURE_LOD_BIAS                                       Enum = 0x84FD
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            Enum = 0x8835
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           Enum = 0x8B4D
	DEPTH_CLAMP_FAR_AMD                                        Enum = 0x901F
	SGIX_instruments                                           Enum = 1
	POST_CONVOLUTION_GREEN_BIAS_EXT                            Enum = 0x8021
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            Enum = 0x83F3
	DU8DV8_ATI                                                 Enum = 0x877A
	COMPARE_REF_TO_TEXTURE_EXT                                 Enum = 0x884E
	REG_10_ATI                                                 Enum = 0x892B
	INTENSITY8UI_EXT                                           Enum = 0x8D7F
	DOUBLE_MAT4x3_EXT                                          Enum = 0x8F4E
	ATTRIB_STACK_DEPTH                                         Enum = 0x0BB0
	CLIP_DISTANCE3                                             Enum = 0x3003
	DEPTH_COMPONENT32_OES                                      Enum = 0x81A7
	R16UI                                                      Enum = 0x8234
	MATRIX10_ARB                                               Enum = 0x88CA
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      Enum = 0x914C
	GL_2D                                                      Enum = 0x0600
	FRAGMENT_PROGRAM_NV                                        Enum = 0x8870
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         Enum = 0x889B
	UNSIGNED_NORMALIZED                                        Enum = 0x8C17
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                Enum = 0x8C8B
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     Enum = 0x8CDC
	RESTART_PATH_NV                                            Enum = 0xF0
	DISPATCH_INDIRECT_BUFFER                                   Enum = 0x90EE
	FOG_COORDINATE_ARRAY_TYPE                                  Enum = 0x8454
	E_TIMES_F_NV                                               Enum = 0x8531
	DRAW_BUFFER13_NV                                           Enum = 0x8832
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      Enum = 0x906B
	INDEX_LOGIC_OP                                             Enum = 0x0BF1
	LUMINANCE8I_EXT                                            Enum = 0x8D92
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             Enum = 0x8DA5
	QUADRATIC_CURVE_TO_NV                                      Enum = 0x0A
	OBJECT_TYPE_ARB                                            Enum = 0x8B4E
	UNSIGNED_INT_IMAGE_3D                                      Enum = 0x9064
	DUAL_INTENSITY8_SGIS                                       Enum = 0x8119
	DEPTH_STENCIL_TEXTURE_MODE                                 Enum = 0x90EA
	SGIX_shadow_ambient                                        Enum = 1
	LINEAR_ATTENUATION                                         Enum = 0x1208
	REPLACE                                                    Enum = 0x1E01
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            Enum = 0x8212
	PROXY_TEXTURE_RECTANGLE_ARB                                Enum = 0x84F7
	TEXTURE_DEPTH_SIZE                                         Enum = 0x884A
	TEXTURE14                                                  Enum = 0x84CE
	VARIABLE_F_NV                                              Enum = 0x8528
	FUNC_ADD_OES                                               Enum = 0x8006
	TEXTURE_IMMUTABLE_LEVELS                                   Enum = 0x82DF
	VERTEX_PROGRAM_POINT_SIZE_ARB                              Enum = 0x8642
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         Enum = 0x90DC
	POINT_DISTANCE_ATTENUATION                                 Enum = 0x8129
	COMPRESSED_RGB                                             Enum = 0x84ED
	DISCRETE_AMD                                               Enum = 0x9006
	AFFINE_3D_NV                                               Enum = 0x9094
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             Enum = 0x00000100
	INDEX_ARRAY_POINTER                                        Enum = 0x8091
	TEXTURE_VIEW_MIN_LEVEL                                     Enum = 0x82DB
	FOG_COORDINATE                                             Enum = 0x8451
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               Enum = 0x850E
	DEPTH_STENCIL_MESA                                         Enum = 0x8750
	OP_RECIP_EXT                                               Enum = 0x8794
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           Enum = 0x9029
	RG16                                                       Enum = 0x822C
	SAMPLER_2D_ARRAY_SHADOW_EXT                                Enum = 0x8DC4
	PREFER_DOUBLEBUFFER_HINT_PGI                               Enum = 0x1A1F8
	INVERTED_SCREEN_W_REND                                     Enum = 0x8491
	DRAW_BUFFER10_NV                                           Enum = 0x882F
	MODULATE_COLOR_IMG                                         Enum = 0x8C04
	SIGNALED                                                   Enum = 0x9119
	BUFFER_BINDING                                             Enum = 0x9302
	SPARE0_NV                                                  Enum = 0x852E
	MATRIX9_ARB                                                Enum = 0x88C9
	MAX_ELEMENT_INDEX                                          Enum = 0x8D6B
	PROGRAM_ADDRESS_REGISTERS_ARB                              Enum = 0x88B0
	READ_ONLY_ARB                                              Enum = 0x88B8
	QUERY_BY_REGION_NO_WAIT                                    Enum = 0x8E16
	DS_SCALE_NV                                                Enum = 0x8710
	TRANSFORM_FEEDBACK_VARYINGS_NV                             Enum = 0x8C83
	LINES                                                      Enum = 0x0001
	DEPTH_RANGE                                                Enum = 0x0B70
	SPOT_CUTOFF                                                Enum = 0x1206
	MAX_FRAGMENT_LIGHTS_SGIX                                   Enum = 0x8404
	LOGIC_OP_MODE                                              Enum = 0x0BF0
	ACCUM_BLUE_BITS                                            Enum = 0x0D5A
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           Enum = 0x801F
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            Enum = 0x8A34
	INDEX_ARRAY_COUNT_EXT                                      Enum = 0x8087
	MODELVIEW14_ARB                                            Enum = 0x872E
	VERTEX_STREAM7_ATI                                         Enum = 0x8773
	INT8_VEC3_NV                                               Enum = 0x8FE2
	ALPHA16_EXT                                                Enum = 0x803E
	DEPENDENT_RGB_TEXTURE_3D_NV                                Enum = 0x8859
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             Enum = 0x88FD
	SAMPLER_BUFFER                                             Enum = 0x8DC2
	UNSIGNED_INT_IMAGE_2D_ARRAY                                Enum = 0x9069
	FONT_Y_MAX_BOUNDS_BIT_NV                                   Enum = 0x00080000
	VIEW_CLASS_S3TC_DXT3_RGBA                                  Enum = 0x82CE
	VERTEX_SHADER_BINDING_EXT                                  Enum = 0x8781
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 Enum = 0x8C80
	INT_IMAGE_2D_ARRAY_EXT                                     Enum = 0x905E
	REPLACE_MIDDLE_SUN                                         Enum = 0x0002
	PATH_FILL_COVER_MODE_NV                                    Enum = 0x9082
	SYNC_FLUSH_COMMANDS_BIT                                    Enum = 0x00000001
	CLIP_PLANE5                                                Enum = 0x3005
	COLOR_TABLE_BIAS                                           Enum = 0x80D7
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            Enum = 0x92C4
	COLOR_ARRAY_LIST_STRIDE_IBM                                Enum = 103082
	ONE_MINUS_DST_COLOR                                        Enum = 0x0307
	LUMINANCE6_ALPHA2                                          Enum = 0x8044
	CULL_MODES_NV                                              Enum = 0x86E0
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          Enum = 0x8C85
	IMAGE_1D_EXT                                               Enum = 0x904C
	POST_CONVOLUTION_COLOR_TABLE                               Enum = 0x80D1
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          Enum = 0x85C2
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          Enum = 0x87F6
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        Enum = 0x8BB2
	TEXTURE_DEPTH_TYPE_ARB                                     Enum = 0x8C16
	IMAGE_BINDING_LAYER                                        Enum = 0x8F3D
	LUMINANCE8_ALPHA8_EXT                                      Enum = 0x8045
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    Enum = 0x8266
	DRAW_BUFFER1_ARB                                           Enum = 0x8826
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     Enum = 0x8CDC
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             Enum = 0x9056
	ALPHA_SCALE                                                Enum = 0x0D1C
	INTERNALFORMAT_GREEN_SIZE                                  Enum = 0x8272
	PIXEL_MIN_FILTER_EXT                                       Enum = 0x8332
	INTERPOLATE_ARB                                            Enum = 0x8575
	GL_4PASS_1_EXT                                             Enum = 0x80A5
	BLEND_EQUATION_ALPHA                                       Enum = 0x883D
	FLOAT_MAT2_ARB                                             Enum = 0x8B5A
	MAX_GEOMETRY_ATOMIC_COUNTERS                               Enum = 0x92D5
	SGIX_subsample                                             Enum = 1
	MAX_3D_TEXTURE_SIZE_OES                                    Enum = 0x8073
	FLOAT_RGBA16_NV                                            Enum = 0x888A
	TEXTURE_USAGE_ANGLE                                        Enum = 0x93A2
	TEXTURE_1D_STACK_BINDING_MESAX                             Enum = 0x875D
	ALPHA8UI_EXT                                               Enum = 0x8D7E
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           Enum = 0x8E24
	EXT_packed_pixels                                          Enum = 1
	SOURCE3_ALPHA_NV                                           Enum = 0x858B
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             Enum = 0x870D
	ACCUM_GREEN_BITS                                           Enum = 0x0D59
	NAME_STACK_DEPTH                                           Enum = 0x0D70
	NOR                                                        Enum = 0x1508
	UNSIGNED_SHORT_5_6_5_EXT                                   Enum = 0x8363
	COMBINER3_NV                                               Enum = 0x8553
	BOOL_VEC4                                                  Enum = 0x8B59
	INNOCENT_CONTEXT_RESET_ARB                                 Enum = 0x8254
	SIGNED_NEGATE_NV                                           Enum = 0x853D
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           Enum = 0x8521
	DOT3_RGBA_EXT                                              Enum = 0x8741
	MATRIX25_ARB                                               Enum = 0x88D9
	TEXTURE_GEN_T                                              Enum = 0x0C61
	HISTOGRAM_RED_SIZE_EXT                                     Enum = 0x8028
	FRAGMENT_LIGHT7_SGIX                                       Enum = 0x8413
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            Enum = 0x8A05
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           Enum = 0x8CD3
	POLYGON_MODE                                               Enum = 0x0B40
	INTENSITY32F_ARB                                           Enum = 0x8817
	DRAW_BUFFER0_ATI                                           Enum = 0x8825
	CON_14_ATI                                                 Enum = 0x894F
	FRONT                                                      Enum = 0x0404
	GL_4D_COLOR_TEXTURE                                        Enum = 0x0604
	VIEW_CLASS_S3TC_DXT5_RGBA                                  Enum = 0x82CF
	COLOR_ATTACHMENT11_EXT                                     Enum = 0x8CEB
	SGIX_texture_scale_bias                                    Enum = 1
	PIXEL_MAP_I_TO_R                                           Enum = 0x0C72
	STENCIL_REF                                                Enum = 0x0B97
	COLOR_ARRAY_STRIDE                                         Enum = 0x8083
	TEXTURE_LOD_BIAS                                           Enum = 0x8501
	VERTEX_SHADER_LOCALS_EXT                                   Enum = 0x87D3
	INT_IMAGE_2D_MULTISAMPLE_EXT                               Enum = 0x9060
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        Enum = 0x80D5
	CURRENT_RASTER_NORMAL_SGIX                                 Enum = 0x8406
	SMALL_CW_ARC_TO_NV                                         Enum = 0x14
	REFERENCED_BY_VERTEX_SHADER                                Enum = 0x9306
	TEXTURE27_ARB                                              Enum = 0x84DB
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         Enum = 0x88A7
	QUERY_BY_REGION_WAIT_NV                                    Enum = 0x8E15
	POLYGON_OFFSET_POINT                                       Enum = 0x2A01
	CURRENT_SECONDARY_COLOR                                    Enum = 0x8459
	VERTEX_ATTRIB_ARRAY3_NV                                    Enum = 0x8653
	COMPRESSED_RGB_FXT1_3DFX                                   Enum = 0x86B0
	DRAW_BUFFER1                                               Enum = 0x8826
	TEXTURE_COMPARE_FUNC                                       Enum = 0x884D
	MAX_GEOMETRY_IMAGE_UNIFORMS                                Enum = 0x90CD
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       Enum = 0x93D3
	DEBUG_CALLBACK_FUNCTION                                    Enum = 0x8244
	TRACE_PIXELS_BIT_MESA                                      Enum = 0x0010
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             Enum = 0x8A03
	DOUBLE_MAT2x4_EXT                                          Enum = 0x8F4A
	COLOR_EXT                                                  Enum = 0x1800
	BLEND_EQUATION_RGB                                         Enum = 0x8009
	POST_CONVOLUTION_BLUE_SCALE                                Enum = 0x801E
	UNSIGNED_SHORT_4_4_4_4_EXT                                 Enum = 0x8033
	GL_4PASS_1_SGIS                                            Enum = 0x80A5
	TEXTURE_BASE_LEVEL                                         Enum = 0x813C
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            Enum = 0x82B2
	UNSIGNED_INVERT_NV                                         Enum = 0x8537
	DSDT_MAG_NV                                                Enum = 0x86F6
	COLOR_ATTACHMENT1                                          Enum = 0x8CE1
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          Enum = 0x04
	MODELVIEW11_ARB                                            Enum = 0x872B
	FLOAT_MAT4                                                 Enum = 0x8B5C
	VERTEX_ARRAY_LENGTH_NV                                     Enum = 0x8F2B
	DEBUG_SOURCE_OTHER_ARB                                     Enum = 0x824B
	PROGRAM_PARAMETERS_ARB                                     Enum = 0x88A8
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      Enum = 0x80
	COLOR_TABLE_SCALE                                          Enum = 0x80D6
	EVAL_2D_NV                                                 Enum = 0x86C0
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     Enum = 0x889F
	STREAM_READ                                                Enum = 0x88E1
	AFFINE_2D_NV                                               Enum = 0x9092
	ARRAY_SIZE                                                 Enum = 0x92FB
	DRAW_BUFFER                                                Enum = 0x0C01
	TEXTURE_ALPHA_SIZE                                         Enum = 0x805F
	HISTOGRAM_FORMAT                                           Enum = 0x8027
	COLOR_SUM_ARB                                              Enum = 0x8458
	SAMPLER_2D_ARRAY_SHADOW_NV                                 Enum = 0x8DC4
	TEXTURE_CONSTANT_DATA_SUNX                                 Enum = 0x81D6
	TANGENT_ARRAY_EXT                                          Enum = 0x8439
	TEXTURE_GREEN_TYPE                                         Enum = 0x8C11
	T2F_IUI_N3F_V2F_EXT                                        Enum = 0x81B3
	SRC2_RGB                                                   Enum = 0x8582
	MAP_TESSELLATION_NV                                        Enum = 0x86C2
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               Enum = 0x88FF
	FLOAT_MAT4x3                                               Enum = 0x8B6A
	RENDERBUFFER_COLOR_SAMPLES_NV                              Enum = 0x8E10
	FONT_UNDERLINE_THICKNESS_BIT_NV                            Enum = 0x08000000
	FRONT_RIGHT                                                Enum = 0x0401
	EMBOSS_LIGHT_NV                                            Enum = 0x855D
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            Enum = 0x862E
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           Enum = 0x8C8A
	RENDERBUFFER_ALPHA_SIZE                                    Enum = 0x8D53
	DEBUG_SEVERITY_MEDIUM                                      Enum = 0x9147
	TEXTURE22_ARB                                              Enum = 0x84D6
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    Enum = 0x906C
	PATH_STENCIL_FUNC_NV                                       Enum = 0x90B7
	COMBINER_SUM_OUTPUT_NV                                     Enum = 0x854C
	WEIGHT_ARRAY_BUFFER_BINDING                                Enum = 0x889E
	CON_12_ATI                                                 Enum = 0x894D
	PERFORMANCE_MONITOR_AMD                                    Enum = 0x9152
	SGIX_tag_sample_buffer                                     Enum = 1
	VERTEX_PROGRAM_NV                                          Enum = 0x8620
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            Enum = 0x889E
	DRAW_FRAMEBUFFER_NV                                        Enum = 0x8CA9
	COMPRESSED_RGBA_ASTC_6x5_KHR                               Enum = 0x93B3
	FUNC_SUBTRACT_OES                                          Enum = 0x800A
	MAX_4D_TEXTURE_SIZE_SGIS                                   Enum = 0x8138
	FOG_OFFSET_VALUE_SGIX                                      Enum = 0x8199
	OP_RECIP_SQRT_EXT                                          Enum = 0x8795
	MAX_PROGRAM_IF_DEPTH_NV                                    Enum = 0x88F6
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      Enum = 0x8CD9
	LUMINANCE12_ALPHA12                                        Enum = 0x8047
	PROGRAM_STRING_NV                                          Enum = 0x8628
	DRAW_BUFFER6                                               Enum = 0x882B
	SAMPLE_MASK_VALUE_NV                                       Enum = 0x8E52
	HISTOGRAM_WIDTH                                            Enum = 0x8026
	CON_8_ATI                                                  Enum = 0x8949
	DRAW_FRAMEBUFFER_BINDING_NV                                Enum = 0x8CA6
	PIXEL_TEXTURE_SGIS                                         Enum = 0x8353
	HILO_NV                                                    Enum = 0x86F4
	DEPTH_STENCIL_TO_BGRA_NV                                   Enum = 0x886F
	LUMINANCE_ALPHA32UI_EXT                                    Enum = 0x8D75
	ARC_TO_NV                                                  Enum = 0xFE
	LINE_STRIP_ADJACENCY_EXT                                   Enum = 0x000B
	CLIP_DISTANCE7                                             Enum = 0x3007
	NUM_COMPRESSED_TEXTURE_FORMATS                             Enum = 0x86A2
	TEXTURE_CUBE_MAP_EXT                                       Enum = 0x8513
	SOURCE0_ALPHA_EXT                                          Enum = 0x8588
	FLOAT_CLEAR_COLOR_VALUE_NV                                 Enum = 0x888D
	CONDITION_SATISFIED_APPLE                                  Enum = 0x911C
	DEBUG_SEVERITY_HIGH_AMD                                    Enum = 0x9146
	DRAW_BUFFER9_NV                                            Enum = 0x882E
	PIXEL_TEX_GEN_MODE_SGIX                                    Enum = 0x832B
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   Enum = 0x87CB
	PIXEL_PACK_BUFFER_EXT                                      Enum = 0x88EB
	DUAL_INTENSITY12_SGIS                                      Enum = 0x811A
	IMAGE_CLASS_2_X_8                                          Enum = 0x82C0
	TEXTURE16                                                  Enum = 0x84D0
	STORAGE_PRIVATE_APPLE                                      Enum = 0x85BD
	GL_3DC_XY_AMD                                              Enum = 0x87FA
	DEPTH_BUFFER_FLOAT_MODE_NV                                 Enum = 0x8DAF
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        Enum = 103084
	ALPHA12_EXT                                                Enum = 0x803D
	EVAL_VERTEX_ATTRIB3_NV                                     Enum = 0x86C9
	MIRROR_CLAMP_TO_EDGE_EXT                                   Enum = 0x8743
	OUTPUT_TEXTURE_COORD7_EXT                                  Enum = 0x87A4
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           Enum = 0x80BB
	DUAL_INTENSITY16_SGIS                                      Enum = 0x811B
	COMBINER7_NV                                               Enum = 0x8557
	QUERY_RESULT_AVAILABLE_EXT                                 Enum = 0x8867
	IMAGE_CUBE                                                 Enum = 0x9050
	LIGHT2                                                     Enum = 0x4002
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             Enum = 0x889B
	PERCENTAGE_AMD                                             Enum = 0x8BC3
	PIXEL_TILE_GRID_HEIGHT_SGIX                                Enum = 0x8143
	PALETTE8_RGB5_A1_OES                                       Enum = 0x8B99
	TEXTURE_BINDING_EXTERNAL_OES                               Enum = 0x8D67
	COLOR_INDEX2_EXT                                           Enum = 0x80E3
	DOUBLEBUFFER                                               Enum = 0x0C32
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       Enum = 0x8DA8
	LOW_INT                                                    Enum = 0x8DF3
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 Enum = 0x8186
	TEXTURE_MATERIAL_FACE_EXT                                  Enum = 0x8351
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             Enum = 0x88B5
	RENDERBUFFER_INTERNAL_FORMAT_OES                           Enum = 0x8D44
	LUMINANCE_ALPHA16I_EXT                                     Enum = 0x8D8D
	INDEX_ARRAY                                                Enum = 0x8077
	EXTENSIONS                                                 Enum = 0x1F03
	POST_TEXTURE_FILTER_SCALE_SGIX                             Enum = 0x817A
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   Enum = 0x8CD4
	GEOMETRY_OUTPUT_TYPE_ARB                                   Enum = 0x8DDC
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            Enum = 0x919F
	QUERY                                                      Enum = 0x82E3
	TRACE_NAME_MESA                                            Enum = 0x8756
	SRGB8                                                      Enum = 0x8C41
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           Enum = 0x910C
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       Enum = 0x93D6
	MAX_CONVOLUTION_HEIGHT_EXT                                 Enum = 0x801B
	VIEW_CLASS_BPTC_UNORM                                      Enum = 0x82D2
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            Enum = 0x851A
	COMPRESSED_SRGB_EXT                                        Enum = 0x8C48
	POST_CONVOLUTION_RED_SCALE                                 Enum = 0x801C
	GL_4PASS_0_SGIS                                            Enum = 0x80A4
	TEXTURE3_ARB                                               Enum = 0x84C3
	STORAGE_CLIENT_APPLE                                       Enum = 0x85B4
	MATRIX_PALETTE_ARB                                         Enum = 0x8840
	RENDERBUFFER_BINDING                                       Enum = 0x8CA7
	RGB_INTEGER                                                Enum = 0x8D98
	TEXTURE_BORDER_COLOR                                       Enum = 0x1004
	COPY_INVERTED                                              Enum = 0x150C
	ABGR_EXT                                                   Enum = 0x8000
	COLOR_COMPONENTS                                           Enum = 0x8283
	TEXTURE_VIEW_NUM_LAYERS                                    Enum = 0x82DE
	SIGNED_IDENTITY_NV                                         Enum = 0x853C
	VERTEX_BLEND_ARB                                           Enum = 0x86A7
	CON_3_ATI                                                  Enum = 0x8944
	STENCIL_INDEX16                                            Enum = 0x8D49
	COLOR_ATTACHMENT15_NV                                      Enum = 0x8CEF
	TEXTURE_COLOR_SAMPLES_NV                                   Enum = 0x9046
	INTENSITY12                                                Enum = 0x804C
	VERTEX_ARRAY_RANGE_VALID_NV                                Enum = 0x851F
	REG_22_ATI                                                 Enum = 0x8937
	REG_26_ATI                                                 Enum = 0x893B
	CURRENT_VERTEX_WEIGHT_EXT                                  Enum = 0x850B
	PATH_GEN_COLOR_FORMAT_NV                                   Enum = 0x90B2
	FOG_FUNC_SGIS                                              Enum = 0x812A
	OUTPUT_COLOR1_EXT                                          Enum = 0x879C
	MOVE_TO_NV                                                 Enum = 0x02
	POLYGON_OFFSET_EXT                                         Enum = 0x8037
	MODELVIEW31_ARB                                            Enum = 0x873F
	GL_4X_BIT_ATI                                              Enum = 0x00000002
	POINT_SIZE_MIN                                             Enum = 0x8126
	PROJECTION                                                 Enum = 0x1701
	TEXTURE_APPLICATION_MODE_EXT                               Enum = 0x834F
	DEBUG_SEVERITY_HIGH_ARB                                    Enum = 0x9146
	POLYGON_OFFSET_FACTOR_EXT                                  Enum = 0x8038
	TEXTURE26_ARB                                              Enum = 0x84DA
	VERTEX_ATTRIB_ARRAY_INTEGER                                Enum = 0x88FD
	TRANSFORM_FEEDBACK                                         Enum = 0x8E22
	INDEX_ARRAY_LENGTH_NV                                      Enum = 0x8F2E
	DEPTH_BUFFER_BIT                                           Enum = 0x00000100
	FRAMEZOOM_FACTOR_SGIX                                      Enum = 0x818C
	COMPRESSED_ALPHA_ARB                                       Enum = 0x84E9
	SAMPLER_1D_ARRAY_SHADOW_EXT                                Enum = 0x8DC3
	MODELVIEW8_ARB                                             Enum = 0x8728
	MAX_VERTEX_STREAMS                                         Enum = 0x8E71
	R8_SNORM                                                   Enum = 0x8F94
	PRIMITIVE_RESTART                                          Enum = 0x8F9D
	LUMINANCE16UI_EXT                                          Enum = 0x8D7A
	TIMEOUT_IGNORED_APPLE                                      Enum = 0xFFFFFFFF
	ALL_STATIC_DATA_IBM                                        Enum = 103060
	SAMPLER_BINDING                                            Enum = 0x8919
	LUMINANCE_INTEGER_EXT                                      Enum = 0x8D9C
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        Enum = 0x92CA
	ELEMENT_ARRAY_BARRIER_BIT                                  Enum = 0x00000002
	RED_SCALE                                                  Enum = 0x0D14
	DEBUG_TYPE_PUSH_GROUP                                      Enum = 0x8269
	OPERAND1_RGB                                               Enum = 0x8591
	IMPLEMENTATION_COLOR_READ_TYPE                             Enum = 0x8B9A
	TESS_EVALUATION_SHADER                                     Enum = 0x8E87
	EXT_blend_minmax                                           Enum = 1
	TEXTURE_GREEN_SIZE                                         Enum = 0x805D
	COMPRESSED_TEXTURE_FORMATS                                 Enum = 0x86A3
	LO_SCALE_NV                                                Enum = 0x870F
	MODELVIEW0_MATRIX_EXT                                      Enum = 0x0BA6
	STENCIL                                                    Enum = 0x1802
	SKIP_COMPONENTS4_NV                                        Enum = -3
	UNSIGNED_INT_VEC3                                          Enum = 0x8DC7
	UNSIGNED_INT8_VEC2_NV                                      Enum = 0x8FED
	IS_ROW_MAJOR                                               Enum = 0x9300
	SAMPLE_PATTERN_EXT                                         Enum = 0x80AC
	INTENSITY16F_ARB                                           Enum = 0x881D
	DRAW_BUFFER12_ARB                                          Enum = 0x8831
	STENCIL_BACK_VALUE_MASK                                    Enum = 0x8CA4
	MAX_COMBINED_ATOMIC_COUNTERS                               Enum = 0x92D7
	DOT_PRODUCT_TEXTURE_2D_NV                                  Enum = 0x86EE
	OBJECT_INFO_LOG_LENGTH_ARB                                 Enum = 0x8B84
	PALETTE8_RGBA8_OES                                         Enum = 0x8B96
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            Enum = 0x8C83
	DRAW_FRAMEBUFFER                                           Enum = 0x8CA9
	SHADER_STORAGE_BUFFER_SIZE                                 Enum = 0x90D5
	SUBPIXEL_BITS                                              Enum = 0x0D50
	OFFSET_TEXTURE_2D_MATRIX_NV                                Enum = 0x86E1
	ELEMENT_ARRAY_POINTER_ATI                                  Enum = 0x876A
	PRIMITIVES_GENERATED                                       Enum = 0x8C87
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           Enum = 0x8CD3
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        Enum = 0x90C9
	INCR                                                       Enum = 0x1E02
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             Enum = 0x84FF
	UNSIGNED_SHORT_8_8_MESA                                    Enum = 0x85BA
	OP_SET_GE_EXT                                              Enum = 0x878C
	FRAGMENT_SHADER_ATI                                        Enum = 0x8920
	RENDERBUFFER_DEPTH_SIZE_EXT                                Enum = 0x8D54
	LAYOUT_DEFAULT_INTEL                                       Enum = 0
	ORDER                                                      Enum = 0x0A01
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        Enum = 0x8BB3
	LOCATION_INDEX                                             Enum = 0x930F
	TEXTURE_DEPTH                                              Enum = 0x8071
	POINT_SIZE_MAX_EXT                                         Enum = 0x8127
	TEXTURE_CUBE_MAP                                           Enum = 0x8513
	R1UI_C4F_N3F_V3F_SUN                                       Enum = 0x85C8
	MAP2_VERTEX_ATTRIB10_4_NV                                  Enum = 0x867A
	DRAW_BUFFER7                                               Enum = 0x882C
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       Enum = 0x8A35
	RGB8UI_EXT                                                 Enum = 0x8D7D
	RGBA8I                                                     Enum = 0x8D8E
	SGI_texture_color_table                                    Enum = 1
	LINE_WIDTH                                                 Enum = 0x0B21
	VIEW_CLASS_48_BITS                                         Enum = 0x82C7
	ACTIVE_TEXTURE                                             Enum = 0x84E0
	PROGRAM_TARGET_NV                                          Enum = 0x8646
	WEIGHT_ARRAY_SIZE_ARB                                      Enum = 0x86AB
	DOT3_RGB_EXT                                               Enum = 0x8740
	MATRIX_INDEX_ARRAY_TYPE_ARB                                Enum = 0x8847
	COMPARE_R_TO_TEXTURE_ARB                                   Enum = 0x884E
	RGBA32UI                                                   Enum = 0x8D70
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           Enum = 0x9117
	FALSE                                                      Enum = 0
	VERTEX_ARRAY_BINDING                                       Enum = 0x85B5
	MAP_ATTRIB_V_ORDER_NV                                      Enum = 0x86C4
	MAX_PROGRAM_TEMPORARIES_ARB                                Enum = 0x88A5
	COPY_READ_BUFFER_BINDING                                   Enum = 0x8F36
	ALL_BARRIER_BITS_EXT                                       Enum = 0xFFFFFFFF
	CONVOLUTION_FILTER_BIAS                                    Enum = 0x8015
	DEBUG_TYPE_ERROR                                           Enum = 0x824C
	OUTPUT_TEXTURE_COORD17_EXT                                 Enum = 0x87AE
	RENDERBUFFER_EXT                                           Enum = 0x8D41
	DEBUG_CATEGORY_PERFORMANCE_AMD                             Enum = 0x914D
	SGIX_depth_texture                                         Enum = 1
	ADD_SIGNED_EXT                                             Enum = 0x8574
	BUFFER_MAPPED_ARB                                          Enum = 0x88BC
	PACK_RESAMPLE_OML                                          Enum = 0x8984
	COMPILE_STATUS                                             Enum = 0x8B81
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      Enum = 103087
	TESS_EVALUATION_TEXTURE                                    Enum = 0x829D
	RGB16UI_EXT                                                Enum = 0x8D77
	DSDT8_MAG8_INTENSITY8_NV                                   Enum = 0x870B
	CON_5_ATI                                                  Enum = 0x8946
	UNSIGNED_INT_IMAGE_1D_EXT                                  Enum = 0x9062
	IMAGE_2D_MULTISAMPLE                                       Enum = 0x9055
	FOG_COORDINATE_SOURCE_EXT                                  Enum = 0x8450
	SLUMINANCE8                                                Enum = 0x8C47
	UNSIGNED_INT16_NV                                          Enum = 0x8FF0
	PATH_STENCIL_REF_NV                                        Enum = 0x90B8
	PALETTE4_RGBA4_OES                                         Enum = 0x8B93
	TEXTURE_IMAGE_VALID_QCOM                                   Enum = 0x8BD8
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         Enum = 0x8C85
	MAX_FRAGMENT_ATOMIC_COUNTERS                               Enum = 0x92D6
	FOG_DENSITY                                                Enum = 0x0B62
	ALPHA_TEST_QCOM                                            Enum = 0x0BC0
	INTENSITY                                                  Enum = 0x8049
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           Enum = 0x87C7
	STENCIL_BACK_FAIL_ATI                                      Enum = 0x8801
	UNSIGNED_INT_IMAGE_1D                                      Enum = 0x9062
	MITER_REVERT_NV                                            Enum = 0x90A7
	VERTEX_WEIGHT_ARRAY_EXT                                    Enum = 0x850C
	DOT_PRODUCT_TEXTURE_1D_NV                                  Enum = 0x885C
	CON_24_ATI                                                 Enum = 0x8959
	TEXTURE_CUBE_MAP_ARRAY                                     Enum = 0x9009
	DEBUG_SEVERITY_MEDIUM_AMD                                  Enum = 0x9147
	BUFFER_VARIABLE                                            Enum = 0x92E5
	EYE_POINT_SGIS                                             Enum = 0x81F4
	VIEW_COMPATIBILITY_CLASS                                   Enum = 0x82B6
	DRAW_BUFFER3                                               Enum = 0x8828
	READ_WRITE_ARB                                             Enum = 0x88BA
	FLOAT_MAT3_ARB                                             Enum = 0x8B5B
	COLOR_ATTACHMENT2                                          Enum = 0x8CE2
	NORMAL_ARRAY_LIST_IBM                                      Enum = 103071
	SAMPLE_MASK_VALUE_SGIS                                     Enum = 0x80AA
	INDEX                                                      Enum = 0x8222
	DRAW_BUFFER11_NV                                           Enum = 0x8830
	SAMPLER_1D_SHADOW_ARB                                      Enum = 0x8B61
	FRAMEBUFFER_BINDING                                        Enum = 0x8CA6
	RGBA_FLOAT32_ATI                                           Enum = 0x8814
	PATH_FILL_BOUNDING_BOX_NV                                  Enum = 0x90A1
	ALPHA32I_EXT                                               Enum = 0x8D84
	FIRST_VERTEX_CONVENTION_EXT                                Enum = 0x8E4D
	DEBUG_TYPE_PORTABILITY                                     Enum = 0x824F
	RGBA32UI_EXT                                               Enum = 0x8D70
	MAP1_GRID_DOMAIN                                           Enum = 0x0DD0
	PACK_IMAGE_HEIGHT                                          Enum = 0x806C
	COLOR_TABLE_SGI                                            Enum = 0x80D0
	VERTEX_PROGRAM_POINT_SIZE_NV                               Enum = 0x8642
	COMP_BIT_ATI                                               Enum = 0x00000002
	DEPTH_ATTACHMENT_OES                                       Enum = 0x8D00
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         Enum = 0x8E5D
	GEOMETRY_SUBROUTINE_UNIFORM                                Enum = 0x92F1
	SHADER_IMAGE_ATOMIC                                        Enum = 0x82A6
	TEXTURE_MEMORY_LAYOUT_INTEL                                Enum = 0x83FF
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            Enum = 0x86F2
	POINT_SPRITE                                               Enum = 0x8861
	ACTIVE_UNIFORM_BLOCKS                                      Enum = 0x8A36
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 Enum = 0x90EC
	SAMPLE_ALPHA_TO_ONE_SGIS                                   Enum = 0x809F
	CLAMP_TO_BORDER_SGIS                                       Enum = 0x812D
	VALIDATE_STATUS                                            Enum = 0x8B83
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           Enum = 0x8DD5
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      Enum = 0x8C2D
	LINE_STIPPLE_REPEAT                                        Enum = 0x0B26
	SPRITE_OBJECT_ALIGNED_SGIX                                 Enum = 0x814D
	FIXED_ONLY                                                 Enum = 0x891D
	UNSIGNED_NORMALIZED_EXT                                    Enum = 0x8C17
	TRANSFORM_FEEDBACK_BUFFER_MODE                             Enum = 0x8C7F
	TEXTURE_UPDATE_BARRIER_BIT                                 Enum = 0x00000100
	PRIMARY_COLOR_ARB                                          Enum = 0x8577
	SAMPLER_1D_ARRAY_EXT                                       Enum = 0x8DC0
	FOG_COORDINATE_EXT                                         Enum = 0x8451
	COMPRESSED_LUMINANCE_ALPHA                                 Enum = 0x84EB
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        Enum = 0x8533
	LOCAL_CONSTANT_EXT                                         Enum = 0x87C3
	INT_VEC4                                                   Enum = 0x8B55
	POINT                                                      Enum = 0x1B00
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             Enum = 0x845B
	PRIMITIVE_RESTART_NV                                       Enum = 0x8558
	GREEN_MIN_CLAMP_INGR                                       Enum = 0x8561
	READ_FRAMEBUFFER_EXT                                       Enum = 0x8CA8
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          Enum = 0x8DDE
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   Enum = 0x8E85
	PROGRAM_INPUT                                              Enum = 0x92E3
	MAP2_TEXTURE_COORD_1                                       Enum = 0x0DB3
	NEAREST_MIPMAP_LINEAR                                      Enum = 0x2702
	SECONDARY_COLOR_ARRAY_STRIDE                               Enum = 0x845C
	COMPRESSED_INTENSITY_ARB                                   Enum = 0x84EC
	MATRIX_INDEX_ARRAY_POINTER_ARB                             Enum = 0x8849
	UNIFORM_ARRAY_STRIDE                                       Enum = 0x8A3C
	MAX_IMAGE_UNITS                                            Enum = 0x8F38
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            Enum = 0x9003
	COLOR_WRITEMASK                                            Enum = 0x0C23
	FEEDBACK_BUFFER_SIZE                                       Enum = 0x0DF1
	PIXEL_CUBIC_WEIGHT_EXT                                     Enum = 0x8333
	TEXTURE_BINDING_RECTANGLE_NV                               Enum = 0x84F6
	MODELVIEW13_ARB                                            Enum = 0x872D
	DEBUG_ASSERT_MESA                                          Enum = 0x875B
	BUFFER_USAGE_ARB                                           Enum = 0x8765
	STENCIL_BACK_PASS_DEPTH_FAIL                               Enum = 0x8802
	MAX_VIEWPORT_DIMS                                          Enum = 0x0D3A
	ASYNC_MARKER_SGIX                                          Enum = 0x8329
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            Enum = 0x851A
	COVERAGE_ALL_FRAGMENTS_NV                                  Enum = 0x8ED5
	INT_SAMPLER_2D_MULTISAMPLE                                 Enum = 0x9109
	NORMAL_ARRAY_STRIDE_EXT                                    Enum = 0x807F
	MAX_GENERAL_COMBINERS_NV                                   Enum = 0x854D
	DOT3_RGB                                                   Enum = 0x86AE
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     Enum = 0x898D
	UNSIGNED_INT_SAMPLER_2D                                    Enum = 0x8DD2
	INT_IMAGE_BUFFER                                           Enum = 0x905C
	FUNC_ADD                                                   Enum = 0x8006
	COLOR_TABLE_GREEN_SIZE                                     Enum = 0x80DB
	INTERNALFORMAT_STENCIL_TYPE                                Enum = 0x827D
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  Enum = 0x86F3
	NEGATIVE_Y_EXT                                             Enum = 0x87DA
	SKIP_COMPONENTS1_NV                                        Enum = -6
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        Enum = 0x90FC
	OUT_OF_MEMORY                                              Enum = 0x0505
	GREEN                                                      Enum = 0x1904
	CONVOLUTION_BORDER_MODE                                    Enum = 0x8013
	VERTEX_ARRAY_STRIDE_EXT                                    Enum = 0x807C
	T2F_IUI_V2F_EXT                                            Enum = 0x81B1
	TEXTURE_COMPRESSION_HINT                                   Enum = 0x84EF
	OP_SET_LT_EXT                                              Enum = 0x878D
	STATIC_COPY                                                Enum = 0x88E6
	PROXY_TEXTURE_1D_ARRAY_EXT                                 Enum = 0x8C19
	ALPHA12                                                    Enum = 0x803D
	DISTANCE_ATTENUATION_SGIS                                  Enum = 0x8129
	MODELVIEW16_ARB                                            Enum = 0x8730
	MODULATE_SIGNED_ADD_ATI                                    Enum = 0x8745
	PIXEL_PACK_BUFFER_BINDING                                  Enum = 0x88ED
	ELEMENT_ARRAY_TYPE_APPLE                                   Enum = 0x8A0D
	COLOR_ATTACHMENT3                                          Enum = 0x8CE3
	SAMPLER_2D_ARRAY_EXT                                       Enum = 0x8DC1
	IMAGE_2D_ARRAY                                             Enum = 0x9053
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          Enum = 0x906B
	VERTEX_ARRAY_POINTER                                       Enum = 0x808E
	MODELVIEW30_ARB                                            Enum = 0x873E
	EVAL_VERTEX_ATTRIB15_NV                                    Enum = 0x86D5
	YCRCB_422_SGIX                                             Enum = 0x81BB
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       Enum = 0x8B4D
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          Enum = 0x8BB6
	LUMINANCE16_ALPHA16_SNORM                                  Enum = 0x901A
	VERSION_3_1                                                Enum = 1
	MAP_WRITE_BIT_EXT                                          Enum = 0x0002
	SCALE_BY_ONE_HALF_NV                                       Enum = 0x8540
	CURRENT_MATRIX_ARB                                         Enum = 0x8641
	RGB565_OES                                                 Enum = 0x8D62
	UNIFORM_BUFFER_EXT                                         Enum = 0x8DEE
	SQUARE_NV                                                  Enum = 0x90A3
	ALLOW_DRAW_OBJ_HINT_PGI                                    Enum = 0x1A20E
	HISTOGRAM_EXT                                              Enum = 0x8024
	MAX_FOG_FUNC_POINTS_SGIS                                   Enum = 0x812C
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  Enum = 0x8211
	BGRA_INTEGER                                               Enum = 0x8D9B
	IMAGE_BINDING_LEVEL                                        Enum = 0x8F3B
	GL_2PASS_1_EXT                                             Enum = 0x80A3
	SAMPLE_BUFFERS_EXT                                         Enum = 0x80A8
	FOG_COORD                                                  Enum = 0x8451
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             Enum = 0x84E3
	WEIGHT_ARRAY_STRIDE_ARB                                    Enum = 0x86AA
	COORD_REPLACE                                              Enum = 0x8862
	POLYGON_SMOOTH_HINT                                        Enum = 0x0C53
	C4UB_V2F                                                   Enum = 0x2A22
	GL_4PASS_2_EXT                                             Enum = 0x80A6
	TESS_CONTROL_TEXTURE                                       Enum = 0x829C
	CND_ATI                                                    Enum = 0x896A
	EXT_polygon_offset                                         Enum = 1
	GREATER                                                    Enum = 0x0204
	MAX_TEXTURE_STACK_DEPTH                                    Enum = 0x0D39
	TEXTURE_WRAP_R                                             Enum = 0x8072
	PACK_ROW_BYTES_APPLE                                       Enum = 0x8A15
	RGB8UI                                                     Enum = 0x8D7D
	INT8_VEC4_NV                                               Enum = 0x8FE3
	BLUE_BIAS                                                  Enum = 0x0D1B
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     Enum = 0x80D4
	IUI_N3F_V2F_EXT                                            Enum = 0x81AF
	OP_MOV_EXT                                                 Enum = 0x8799
	SAMPLER_CUBE                                               Enum = 0x8B60
	LINEAR_MIPMAP_NEAREST                                      Enum = 0x2701
	COMPRESSED_TEXTURE_FORMATS_ARB                             Enum = 0x86A3
	SCALE_BY_FOUR_NV                                           Enum = 0x853F
	UNSIGNED_SHORT                                             Enum = 0x1403
	DUAL_ALPHA12_SGIS                                          Enum = 0x8112
	MODELVIEW25_ARB                                            Enum = 0x8739
	OUTPUT_FOG_EXT                                             Enum = 0x87BD
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         Enum = 0x8841
	SPRITE_MODE_SGIX                                           Enum = 0x8149
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       Enum = 0x886D
	UNSIGNED_INT_SAMPLER_1D                                    Enum = 0x8DD1
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              Enum = 0x8E5D
	CLIP_DISTANCE0                                             Enum = 0x3000
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           Enum = 0x8355
	NEGATIVE_ONE_EXT                                           Enum = 0x87DF
	INDEX_MODE                                                 Enum = 0x0C30
	POST_CONVOLUTION_BLUE_SCALE_EXT                            Enum = 0x801E
	TEXTURE_CLIPMAP_FRAME_SGIX                                 Enum = 0x8172
	AVERAGE_EXT                                                Enum = 0x8335
	TRACE_ARRAYS_BIT_MESA                                      Enum = 0x0004
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        Enum = 0x88A3
	TEXTURE_LUMINANCE_TYPE_ARB                                 Enum = 0x8C14
	COLOR_ATTACHMENT4                                          Enum = 0x8CE4
	LUMINANCE16I_EXT                                           Enum = 0x8D8C
	FOG                                                        Enum = 0x0B60
	LUMINANCE_ALPHA_FLOAT16_ATI                                Enum = 0x881F
	TEXTURE_STENCIL_SIZE                                       Enum = 0x88F1
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     Enum = 0x8C28
	TEXTURE_COORD_ARRAY_POINTER                                Enum = 0x8092
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             Enum = 0x80BF
	MIRRORED_REPEAT_IBM                                        Enum = 0x8370
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            Enum = 0x86F0
	RGB16F                                                     Enum = 0x881B
	TEXTURE_BLUE_TYPE                                          Enum = 0x8C12
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           Enum = 0x90D6
	RG16UI                                                     Enum = 0x823A
	PROXY_TEXTURE_RECTANGLE_NV                                 Enum = 0x84F7
	TRANSFORM_FEEDBACK_BUFFER_NV                               Enum = 0x8C8E
	IMAGE_2D_MULTISAMPLE_ARRAY                                 Enum = 0x9056
	TEXTURE_COMPARE_MODE_ARB                                   Enum = 0x884C
	CON_13_ATI                                                 Enum = 0x894E
	IMAGE_1D_ARRAY_EXT                                         Enum = 0x9052
	EXT_texture3D                                              Enum = 1
	DUAL_LUMINANCE4_SGIS                                       Enum = 0x8114
	VERTEX_ARRAY_BINDING_APPLE                                 Enum = 0x85B5
	TRANSPOSE_PROGRAM_MATRIX_EXT                               Enum = 0x8E2E
	DEBUG_CATEGORY_OTHER_AMD                                   Enum = 0x9150
	DRAW_BUFFER10_ATI                                          Enum = 0x882F
	DYNAMIC_READ                                               Enum = 0x88E9
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               Enum = 0x8E1E
	COMPRESSED_RGBA_ASTC_10x5_KHR                              Enum = 0x93B8
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     Enum = 0x93DD
	SGIX_pixel_tiles                                           Enum = 1
	MAP1_TEXTURE_COORD_2                                       Enum = 0x0D94
	DUAL_INTENSITY4_SGIS                                       Enum = 0x8118
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        Enum = 0x8188
	COMBINE_RGB_EXT                                            Enum = 0x8571
	PROGRAM_LENGTH_NV                                          Enum = 0x8627
	DRAW_BUFFER2_ARB                                           Enum = 0x8827
	FLOAT_RGBA_NV                                              Enum = 0x8883
	MODELVIEW_STACK_DEPTH                                      Enum = 0x0BA3
	QUAD_LUMINANCE4_SGIS                                       Enum = 0x8120
	TANGENT_ARRAY_POINTER_EXT                                  Enum = 0x8442
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        Enum = 0x8B49
	FRAGMENT_SUBROUTINE                                        Enum = 0x92EC
	COLOR_ENCODING                                             Enum = 0x8296
	RGB4_S3TC                                                  Enum = 0x83A1
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       Enum = 0x83F6
	LAYER_NV                                                   Enum = 0x8DAA
	UNSIGNED_INT_SAMPLER_2D_RECT                               Enum = 0x8DD5
	FRAMEBUFFER_DEFAULT_HEIGHT                                 Enum = 0x9311
	CALLIGRAPHIC_FRAGMENT_SGIX                                 Enum = 0x8183
	MAX_PROGRAM_MATRICES_ARB                                   Enum = 0x862F
	GL_2X_BIT_ATI                                              Enum = 0x00000001
	CURRENT_PROGRAM                                            Enum = 0x8B8D
	FACTOR_MAX_AMD                                             Enum = 0x901D
	INDEX_BIT_PGI                                              Enum = 0x00080000
	SGIX_sprite                                                Enum = 1
	STREAM_DRAW                                                Enum = 0x88E0
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        Enum = 0x903B
	COMPUTE_SUBROUTINE                                         Enum = 0x92ED
	RETURN                                                     Enum = 0x0102
	LEQUAL                                                     Enum = 0x0203
	LINE_STRIP                                                 Enum = 0x0003
	NONE                                                       Enum = 0
	RGBA4_DXT5_S3TC                                            Enum = 0x83A5
	FOG_COORD_ARRAY                                            Enum = 0x8457
	OUTPUT_TEXTURE_COORD11_EXT                                 Enum = 0x87A8
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         Enum = 0x87C5
	TEXTURE_BINDING_RENDERBUFFER_NV                            Enum = 0x8E53
	DRAW_INDIRECT_BUFFER_BINDING                               Enum = 0x8F43
	RENDERER                                                   Enum = 0x1F01
	LINEAR_CLIPMAP_NEAREST_SGIX                                Enum = 0x844F
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         Enum = 0x8C84
	FRAMEBUFFER_UNSUPPORTED_OES                                Enum = 0x8CDD
	COLOR_ARRAY_ADDRESS_NV                                     Enum = 0x8F23
	CURRENT_VERTEX_ATTRIB                                      Enum = 0x8626
	MODELVIEW26_ARB                                            Enum = 0x873A
	DOUBLE_VEC3                                                Enum = 0x8FFD
	FASTEST                                                    Enum = 0x1101
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            Enum = 0x8C01
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              Enum = 0x8C7E
	STENCIL_INDEX4                                             Enum = 0x8D47
	INT_SAMPLER_2D_EXT                                         Enum = 0x8DCA
	POST_COLOR_MATRIX_ALPHA_BIAS                               Enum = 0x80BB
	WEIGHT_ARRAY_ARB                                           Enum = 0x86AD
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             Enum = 0x8E8C
	EYE_PLANE                                                  Enum = 0x2502
	CURRENT_TANGENT_EXT                                        Enum = 0x843B
	DRAW_BUFFER3_NV                                            Enum = 0x8828
	WRITE_ONLY_OES                                             Enum = 0x88B9
	TEXTURE_SWIZZLE_G_EXT                                      Enum = 0x8E43
	COMPRESSED_SIGNED_R11_EAC                                  Enum = 0x9271
	PATCHES                                                    Enum = 0x000E
	ZOOM_Y                                                     Enum = 0x0D17
	MULTISAMPLE_FILTER_HINT_NV                                 Enum = 0x8534
	COMBINE_ALPHA                                              Enum = 0x8572
	UNSIGNED_SHORT_15_1_MESA                                   Enum = 0x8753
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        Enum = 0x880A
	FIELD_LOWER_NV                                             Enum = 0x9023
	COLOR_ATTACHMENT3_EXT                                      Enum = 0x8CE3
	TRIANGLE_STRIP                                             Enum = 0x0005
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            Enum = 0x8175
	TEXTURE_CUBE_MAP_POSITIVE_X                                Enum = 0x8515
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          Enum = 0x8B4A
	NUM_SHADER_BINARY_FORMATS                                  Enum = 0x8DF9
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            Enum = 0x8E89
	LINES_ADJACENCY                                            Enum = 0x000A
	OUTPUT_TEXTURE_COORD28_EXT                                 Enum = 0x87B9
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       Enum = 0x8856
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              Enum = 0x900D
	EXT_convolution                                            Enum = 1
	EMBOSS_CONSTANT_NV                                         Enum = 0x855E
	LOCAL_CONSTANT_DATATYPE_EXT                                Enum = 0x87ED
	RENDERBUFFER_HEIGHT_EXT                                    Enum = 0x8D43
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          Enum = 0x906A
	IMAGE_CLASS_1_X_32                                         Enum = 0x82BB
	FRAGMENT_LIGHT6_SGIX                                       Enum = 0x8412
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          Enum = 0x864D
	ALIASED_LINE_WIDTH_RANGE                                   Enum = 0x846E
	VERTEX_ATTRIB_ARRAY13_NV                                   Enum = 0x865D
	TRACE_ALL_BITS_MESA                                        Enum = 0xFFFF
	DEPTH_COMPONENT24                                          Enum = 0x81A6
	SRGB_READ                                                  Enum = 0x8297
	MODELVIEW17_ARB                                            Enum = 0x8731
	REG_13_ATI                                                 Enum = 0x892E
	UNSIGNED_SHORT_8_8_REV_MESA                                Enum = 0x85BB
	SIGNED_RGBA_NV                                             Enum = 0x86FB
	VERTEX_ATTRIB_MAP2_APPLE                                   Enum = 0x8A01
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     Enum = 0x8CDB
	TEXTURE                                                    Enum = 0x1702
	LUMINANCE8_ALPHA8                                          Enum = 0x8045
	DRAW_BUFFER14                                              Enum = 0x8833
	PROGRAM_ATTRIBS_ARB                                        Enum = 0x88AC
	NEGATIVE_W_EXT                                             Enum = 0x87DC
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         Enum = 0x8DE4
	PROVOKING_VERTEX                                           Enum = 0x8E4F
	SAMPLE_MASK                                                Enum = 0x8E51
	PATH_COMMAND_COUNT_NV                                      Enum = 0x909D
	PIXEL_TILE_GRID_DEPTH_SGIX                                 Enum = 0x8144
	OP_CLAMP_EXT                                               Enum = 0x878E
	MATRIX15_ARB                                               Enum = 0x88CF
	UNSIGNED_INT_VEC2_EXT                                      Enum = 0x8DC6
	POINT_BIT                                                  Enum = 0x00000002
	ALPHA_TEST_REF_QCOM                                        Enum = 0x0BC2
	LUMINANCE12_ALPHA4_EXT                                     Enum = 0x8046
	BLEND_SRC_ALPHA_EXT                                        Enum = 0x80CB
	TEXTURE2_ARB                                               Enum = 0x84C2
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  Enum = 0x8C88
	LINE_STRIP_ADJACENCY_ARB                                   Enum = 0x000B
	RESCALE_NORMAL                                             Enum = 0x803A
	BLEND_DST_ALPHA_OES                                        Enum = 0x80CA
	TEXTURE_COMPARE_SGIX                                       Enum = 0x819A
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                Enum = 0x8A46
	TEXTURE_SWIZZLE_G                                          Enum = 0x8E43
	PACK_COMPRESSED_BLOCK_SIZE                                 Enum = 0x912E
	SGIS_texture4D                                             Enum = 1
	VARIANT_ARRAY_TYPE_EXT                                     Enum = 0x87E7
	SCALED_RESOLVE_FASTEST_EXT                                 Enum = 0x90BA
	FRAGMENT_LIGHTING_SGIX                                     Enum = 0x8400
	GEOMETRY_OUTPUT_TYPE                                       Enum = 0x8918
	TEXTURE_COORD_ARRAY_TYPE                                   Enum = 0x8089
	GL_2_BYTES                                                 Enum = 0x1407
	S                                                          Enum = 0x2000
	FRAGMENT_TEXTURE                                           Enum = 0x829F
	MAX_UNIFORM_BUFFER_BINDINGS                                Enum = 0x8A2F
	UNSIGNED_INT_8_8_8_8_EXT                                   Enum = 0x8035
	PROGRAM_BINARY_RETRIEVABLE_HINT                            Enum = 0x8257
	OPERAND2_ALPHA_EXT                                         Enum = 0x859A
	CURRENT_ATTRIB_NV                                          Enum = 0x8626
	LO_BIAS_NV                                                 Enum = 0x8715
	DRAW_BUFFER3_ARB                                           Enum = 0x8828
	COLOR_ATTACHMENT10_EXT                                     Enum = 0x8CEA
	INT_SAMPLER_3D                                             Enum = 0x8DCB
	VERSION_1_5                                                Enum = 1
	DST_COLOR                                                  Enum = 0x0306
	HISTOGRAM_ALPHA_SIZE                                       Enum = 0x802B
	DEBUG_GROUP_STACK_DEPTH                                    Enum = 0x826D
	QUERY_COUNTER_BITS                                         Enum = 0x8864
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           Enum = 0x8CAB
	FULL_STIPPLE_HINT_PGI                                      Enum = 0x1A219
	EDGE_FLAG_ARRAY_STRIDE                                     Enum = 0x808C
	TEXTURE_INDEX_SIZE_EXT                                     Enum = 0x80ED
	CON_9_ATI                                                  Enum = 0x894A
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                Enum = 0x8A45
	TRANSFORM_FEEDBACK_BINDING_NV                              Enum = 0x8E25
	SAMPLE_COVERAGE_ARB                                        Enum = 0x80A0
	IMAGE_CLASS_1_X_8                                          Enum = 0x82C1
	STORAGE_CACHED_APPLE                                       Enum = 0x85BE
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      Enum = 0x8851
	CON_26_ATI                                                 Enum = 0x895B
	FRAMEBUFFER_SRGB_EXT                                       Enum = 0x8DB9
	SAMPLER_2D_ARRAY                                           Enum = 0x8DC1
	GL_422_REV_AVERAGE_EXT                                     Enum = 0x80CF
	REG_24_ATI                                                 Enum = 0x8939
	NORMAL_ARRAY_POINTER                                       Enum = 0x808F
	SMOOTH_LINE_WIDTH_GRANULARITY                              Enum = 0x0B23
	MATRIX12_ARB                                               Enum = 0x88CC
	TEXTURE_2D_ARRAY_EXT                                       Enum = 0x8C1A
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     Enum = 0x8CD0
	PATH_COMPUTED_LENGTH_NV                                    Enum = 0x90A0
	MAP2_VERTEX_3                                              Enum = 0x0DB7
	SECONDARY_COLOR_NV                                         Enum = 0x852D
	ADD_SIGNED                                                 Enum = 0x8574
	HALF_FLOAT                                                 Enum = 0x140B
	OBJECT_DISTANCE_TO_LINE_SGIS                               Enum = 0x81F3
	FLOAT_RGB16_NV                                             Enum = 0x8888
	SGIX_calligraphic_fragment                                 Enum = 1
	GL_4PASS_2_SGIS                                            Enum = 0x80A6
	MATRIX_INDEX_ARRAY_SIZE_ARB                                Enum = 0x8846
	INT_SAMPLER_BUFFER_AMD                                     Enum = 0x9002
	MAP_WRITE_BIT                                              Enum = 0x0002
	CUBIC_EXT                                                  Enum = 0x8334
	COMBINER0_NV                                               Enum = 0x8550
	OPERAND2_RGB_EXT                                           Enum = 0x8592
	COMPARE_R_TO_TEXTURE                                       Enum = 0x884E
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            Enum = 0x8897
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  Enum = 0x90D9
	MAX_CLIP_DISTANCES                                         Enum = 0x0D32
	DETAIL_TEXTURE_2D_SGIS                                     Enum = 0x8095
	CLEAR_BUFFER                                               Enum = 0x82B4
	IMAGE_CLASS_2_X_32                                         Enum = 0x82BA
	OPERAND2_ALPHA                                             Enum = 0x859A
	SIGNED_LUMINANCE_NV                                        Enum = 0x8701
	PROXY_TEXTURE_1D_STACK_MESAX                               Enum = 0x875B
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 Enum = 0x87CD
	MAP_FLUSH_EXPLICIT_BIT                                     Enum = 0x0010
	POST_COLOR_MATRIX_COLOR_TABLE                              Enum = 0x80D2
	SWIZZLE_STR_ATI                                            Enum = 0x8976
	TEXTURE_BUFFER_DATA_STORE_BINDING                          Enum = 0x8C2D
	REFERENCED_BY_COMPUTE_SHADER                               Enum = 0x930B
	SINGLE_COLOR_EXT                                           Enum = 0x81F9
	TEXTURE10_ARB                                              Enum = 0x84CA
	SUBSAMPLE_DISTANCE_AMD                                     Enum = 0x883F
	IMAGE_BINDING_LAYER_EXT                                    Enum = 0x8F3D
	FLOAT16_NV                                                 Enum = 0x8FF8
	VERTEX_ARRAY_EXT                                           Enum = 0x8074
	DRAW_BUFFER2_ATI                                           Enum = 0x8827
	PERFMON_RESULT_SIZE_AMD                                    Enum = 0x8BC5
	RENDERBUFFER_SAMPLES_ANGLE                                 Enum = 0x8CAB
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       Enum = 0x8E5C
	PATH_INITIAL_END_CAP_NV                                    Enum = 0x9077
	PIXEL_MAP_S_TO_S_SIZE                                      Enum = 0x0CB1
	GREEN_BIAS                                                 Enum = 0x0D19
	RGBA4                                                      Enum = 0x8056
	SPRITE_TRANSLATION_SGIX                                    Enum = 0x814B
	TEXTURE15_ARB                                              Enum = 0x84CF
	COMPRESSED_INTENSITY                                       Enum = 0x84EC
	OUTPUT_TEXTURE_COORD10_EXT                                 Enum = 0x87A7
	SAMPLER_1D_ARRAY_SHADOW                                    Enum = 0x8DC3
	IMAGE_BINDING_FORMAT_EXT                                   Enum = 0x906E
	TEXTURE_MATRIX                                             Enum = 0x0BA8
	EYE_LINEAR                                                 Enum = 0x2400
	NEAREST                                                    Enum = 0x2600
	MAP1_VERTEX_ATTRIB2_4_NV                                   Enum = 0x8662
	TEXTURE_2D_STACK_BINDING_MESAX                             Enum = 0x875E
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       Enum = 0x8E8D
	TEXTURE_INTERNAL_FORMAT                                    Enum = 0x1003
	V2F                                                        Enum = 0x2A20
	IUI_V2F_EXT                                                Enum = 0x81AD
	FRAGMENT_LIGHT0_SGIX                                       Enum = 0x840C
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    Enum = 0x8E82
	DEPTH_TEST                                                 Enum = 0x0B71
	SHADOW_AMBIENT_SGIX                                        Enum = 0x80BF
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             Enum = 0x8A07
	IMAGE_BUFFER                                               Enum = 0x9051
	STACK_OVERFLOW                                             Enum = 0x0503
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        Enum = 0x8808
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         Enum = 0x8B4C
	VERSION_3_2                                                Enum = 1
	MAP1_TEXTURE_COORD_4                                       Enum = 0x0D96
	DEBUG_SOURCE_API                                           Enum = 0x8246
	TEXTURE23_ARB                                              Enum = 0x84D7
	NEGATIVE_Z_EXT                                             Enum = 0x87DB
	LUMINANCE_ALPHA32F_ARB                                     Enum = 0x8819
	GREEN_BIT_ATI                                              Enum = 0x00000002
	VERTEX_PROGRAM_POSITION_MESA                               Enum = 0x8BB4
	UNSIGNED_INT64_VEC2_NV                                     Enum = 0x8FF5
	ACCUM_ALPHA_BITS                                           Enum = 0x0D5B
	TEXTURE_PRIORITY                                           Enum = 0x8066
	DEPTH                                                      Enum = 0x1801
	MAP2_VERTEX_ATTRIB15_4_NV                                  Enum = 0x867F
	GL_3D_COLOR_TEXTURE                                        Enum = 0x0603
	IMAGE_ROTATE_ORIGIN_Y_HP                                   Enum = 0x815B
	COMBINE4_NV                                                Enum = 0x8503
	VERTEX_PROGRAM_ARB                                         Enum = 0x8620
	DRAW_BUFFER2                                               Enum = 0x8827
	DRAW_BUFFER4_NV                                            Enum = 0x8829
	DEPTH_ATTACHMENT                                           Enum = 0x8D00
	PATH_STENCIL_VALUE_MASK_NV                                 Enum = 0x90B9
	MAX_DEFORMATION_ORDER_SGIX                                 Enum = 0x8197
	TEXTURE6_ARB                                               Enum = 0x84C6
	TRANSFORM_HINT_APPLE                                       Enum = 0x85B1
	SHADER_BINARY_FORMATS                                      Enum = 0x8DF8
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            Enum = 0x903C
	MAP2_TEXTURE_COORD_2                                       Enum = 0x0DB4
	TEXTURE_INTENSITY_SIZE_EXT                                 Enum = 0x8061
	REFLECTION_MAP_ARB                                         Enum = 0x8512
	ALPHA_MAX_CLAMP_INGR                                       Enum = 0x8567
	DRAW_BUFFER2_NV                                            Enum = 0x8827
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           Enum = 0x8C3B
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            Enum = 0x91BC
	TEXTURE_WRAP_Q_SGIS                                        Enum = 0x8137
	DISPLAY_LIST                                               Enum = 0x82E7
	COLOR_ATTACHMENT12_NV                                      Enum = 0x8CEC
	RGB32UI_EXT                                                Enum = 0x8D71
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             Enum = 0x1A203
	VERSION_2_0                                                Enum = 1
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             Enum = 0x8624
	SIGNED_HILO16_NV                                           Enum = 0x86FA
	BUMP_TEX_UNITS_ATI                                         Enum = 0x8778
	COMPRESSED_SRGB                                            Enum = 0x8C48
	ALPHA_TEST                                                 Enum = 0x0BC0
	SMOOTH_POINT_SIZE_GRANULARITY                              Enum = 0x0B13
	CON_16_ATI                                                 Enum = 0x8951
	TEXTURE_BINDING_2D_ARRAY                                   Enum = 0x8C1D
	PIXEL_MAP_B_TO_B                                           Enum = 0x0C78
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             Enum = 0x8C27
	PROXY_TEXTURE_RECTANGLE                                    Enum = 0x84F7
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          Enum = 0x8520
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        Enum = 0x889D
	ONE_MINUS_SRC1_ALPHA                                       Enum = 0x88FB
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       Enum = 0x898F
	OBJECT_ACTIVE_UNIFORMS_ARB                                 Enum = 0x8B86
	COMPRESSED_RGBA_ASTC_6x6_KHR                               Enum = 0x93B4
	SELECTION_BUFFER_POINTER                                   Enum = 0x0DF3
	SRGB8_ALPHA8_EXT                                           Enum = 0x8C43
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         Enum = 0x90D7
	SAMPLER_OBJECT_AMD                                         Enum = 0x9155
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   Enum = 0x8A32
	FRACTIONAL_ODD                                             Enum = 0x8E7B
	TEXTURE_GEN_S                                              Enum = 0x0C60
	MATRIX11_ARB                                               Enum = 0x88CB
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     Enum = 0x8CD0
	TRANSFORM_FEEDBACK_ACTIVE                                  Enum = 0x8E24
	CONSTANT                                                   Enum = 0x8576
	PRIMITIVES_GENERATED_EXT                                   Enum = 0x8C87
	CONSTANT_BORDER                                            Enum = 0x8151
	YCBCR_MESA                                                 Enum = 0x8757
	COLOR_ATTACHMENT7_EXT                                      Enum = 0x8CE7
	TESS_CONTROL_SUBROUTINE_UNIFORM                            Enum = 0x92EF
	LUMINANCE4_EXT                                             Enum = 0x803F
	MODELVIEW4_ARB                                             Enum = 0x8724
	DT_SCALE_NV                                                Enum = 0x8711
	SAMPLER_3D_ARB                                             Enum = 0x8B5F
	SAMPLE_MASK_NV                                             Enum = 0x8E51
	UNSIGNED_INT_2_10_10_10_REV                                Enum = 0x8368
	EYE_DISTANCE_TO_POINT_SGIS                                 Enum = 0x81F0
	EYE_LINE_SGIS                                              Enum = 0x81F6
	VERTEX_ATTRIB_RELATIVE_OFFSET                              Enum = 0x82D5
	TANGENT_ARRAY_STRIDE_EXT                                   Enum = 0x843F
	SOURCE1_RGB_EXT                                            Enum = 0x8581
	UNIFORM_BLOCK_NAME_LENGTH                                  Enum = 0x8A41
	PROGRAM_PIPELINE_OBJECT_EXT                                Enum = 0x8A4F
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 Enum = 0x92CE
	TEXTURE_COORD_ARRAY_STRIDE                                 Enum = 0x808A
	OBJECT_LINE_SGIS                                           Enum = 0x81F7
	SURFACE_REGISTERED_NV                                      Enum = 0x86FD
	PROGRAM_ERROR_STRING_NV                                    Enum = 0x8874
	ATTACHED_SHADERS                                           Enum = 0x8B85
	POINT_SIZE_ARRAY_OES                                       Enum = 0x8B9C
	IMAGE_CUBE_MAP_ARRAY                                       Enum = 0x9054
	UNSIGNED_INT_IMAGE_BUFFER                                  Enum = 0x9067
	MINMAX_FORMAT                                              Enum = 0x802F
	QUAD_LUMINANCE8_SGIS                                       Enum = 0x8121
	IMAGE_COMPATIBILITY_CLASS                                  Enum = 0x82A8
	COMPRESSED_RGBA_FXT1_3DFX                                  Enum = 0x86B1
	REG_19_ATI                                                 Enum = 0x8934
	RGBA16I_EXT                                                Enum = 0x8D88
	DUP_LAST_CUBIC_CURVE_TO_NV                                 Enum = 0xF4
	FLAT                                                       Enum = 0x1D00
	MODULATE                                                   Enum = 0x2100
	MAX_CONVOLUTION_HEIGHT                                     Enum = 0x801B
	DEPTH_COMPONENT32_ARB                                      Enum = 0x81A7
	PARALLEL_ARRAYS_INTEL                                      Enum = 0x83F4
	VERTEX_ATTRIB_ARRAY_POINTER                                Enum = 0x8645
	UNSIGNED_INT16_VEC2_NV                                     Enum = 0x8FF1
	VIBRANCE_SCALE_NV                                          Enum = 0x8713
	RGBA32I_EXT                                                Enum = 0x8D82
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              Enum = 0x8DD7
	UNSIGNED_INT8_NV                                           Enum = 0x8FEC
	INT_IMAGE_1D_EXT                                           Enum = 0x9057
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          Enum = 0x00000002
	EXT_vertex_array                                           Enum = 1
	SGIX_clipmap                                               Enum = 1
	GEOMETRY_INPUT_TYPE                                        Enum = 0x8917
	ALWAYS                                                     Enum = 0x0207
	CONVOLUTION_WIDTH_EXT                                      Enum = 0x8018
	FLOAT_MAT3                                                 Enum = 0x8B5B
	SUB_ATI                                                    Enum = 0x8965
	FIELD_UPPER_NV                                             Enum = 0x9022
	TEXTURE_DEPTH_TYPE                                         Enum = 0x8C16
	MAX_SAMPLES                                                Enum = 0x8D57
	R16_SNORM                                                  Enum = 0x8F98
	ALLOW_DRAW_FRG_HINT_PGI                                    Enum = 0x1A210
	ZOOM_X                                                     Enum = 0x0D16
	TEXTURE_NORMAL_EXT                                         Enum = 0x85AF
	CON_28_ATI                                                 Enum = 0x895D
	COMPRESSED_RGBA_ASTC_4x4_KHR                               Enum = 0x93B0
	MAX_ASYNC_HISTOGRAM_SGIX                                   Enum = 0x832D
	OUTPUT_TEXTURE_COORD1_EXT                                  Enum = 0x879E
	DRAW_BUFFER9_ARB                                           Enum = 0x882E
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       Enum = 0x88B2
	MATRIX20_ARB                                               Enum = 0x88D4
	CON_1_ATI                                                  Enum = 0x8942
	BGR_INTEGER                                                Enum = 0x8D9A
	LINEAR_CLIPMAP_LINEAR_SGIX                                 Enum = 0x8170
	ASYNC_TEX_IMAGE_SGIX                                       Enum = 0x835C
	TEXTURE_DT_SIZE_NV                                         Enum = 0x871E
	CONSERVE_MEMORY_HINT_PGI                                   Enum = 0x1A1FD
	TRANSPOSE_PROJECTION_MATRIX                                Enum = 0x84E4
	CLAMP_VERTEX_COLOR                                         Enum = 0x891A
	CIRCULAR_CW_ARC_TO_NV                                      Enum = 0xFA
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    Enum = 0x92C8
	AUX1                                                       Enum = 0x040A
	DEPTH_COMPONENT                                            Enum = 0x1902
	PROGRAM_PIPELINE_BINDING                                   Enum = 0x825A
	DRAW_INDIRECT_UNIFIED_NV                                   Enum = 0x8F40
	OPERAND2_RGB                                               Enum = 0x8592
	CURRENT_WEIGHT_ARB                                         Enum = 0x86A8
	EVAL_VERTEX_ATTRIB1_NV                                     Enum = 0x86C7
	MATRIX23_ARB                                               Enum = 0x88D7
	FRONT_AND_BACK                                             Enum = 0x0408
	POST_CONVOLUTION_GREEN_BIAS                                Enum = 0x8021
	ELEMENT_ARRAY_TYPE_ATI                                     Enum = 0x8769
	OP_POWER_EXT                                               Enum = 0x8793
	DOT2_ADD_ATI                                               Enum = 0x896C
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              Enum = 0x8DED
	MOVE_TO_RESETS_NV                                          Enum = 0x90B5
	SGIX_texture_multi_buffer                                  Enum = 1
	OUTPUT_TEXTURE_COORD21_EXT                                 Enum = 0x87B2
	INTENSITY_FLOAT32_ATI                                      Enum = 0x8817
	DRAW_BUFFER6_NV                                            Enum = 0x882B
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        Enum = 0x00000020
	POINT_SMOOTH                                               Enum = 0x0B10
	DEPTH_COMPONENTS                                           Enum = 0x8284
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        Enum = 0x82AC
	PROGRAM_PARAMETER_NV                                       Enum = 0x8644
	MAP1_VERTEX_ATTRIB3_4_NV                                   Enum = 0x8663
	STATIC_READ_ARB                                            Enum = 0x88E5
	SGIX_interlace                                             Enum = 1
	VIEWPORT                                                   Enum = 0x0BA2
	GL_3_BYTES                                                 Enum = 0x1408
	RGBA_S3TC                                                  Enum = 0x83A2
	UNDEFINED_APPLE                                            Enum = 0x8A1C
	INT_IMAGE_1D                                               Enum = 0x9057
	AUX3                                                       Enum = 0x040C
	UNPACK_ALIGNMENT                                           Enum = 0x0CF5
	MULTISAMPLE_3DFX                                           Enum = 0x86B2
	STATIC_ATI                                                 Enum = 0x8760
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       Enum = 0x8C8F
	COLOR_ATTACHMENT7                                          Enum = 0x8CE7
	COLOR_TABLE_RED_SIZE                                       Enum = 0x80DA
	SGX_PROGRAM_BINARY_IMG                                     Enum = 0x9130
	TEXTURE_IMAGE_TYPE                                         Enum = 0x8290
	SGIX_blend_alpha_minmax                                    Enum = 1
	POLYGON_OFFSET_FILL                                        Enum = 0x8037
	LIGHT0                                                     Enum = 0x4000
	LUMINANCE_ALPHA16F_ARB                                     Enum = 0x881F
	ALPHA4_EXT                                                 Enum = 0x803B
	TEXTURE_COLOR_WRITEMASK_SGIS                               Enum = 0x81EF
	PIXEL_UNPACK_BUFFER_ARB                                    Enum = 0x88EC
	RENDERBUFFER_GREEN_SIZE                                    Enum = 0x8D51
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      Enum = 0x93DA
	NATIVE_GRAPHICS_END_HINT_PGI                               Enum = 0x1A204
	LUMINANCE6_ALPHA2_EXT                                      Enum = 0x8044
	COLOR_INDEX12_EXT                                          Enum = 0x80E6
	VERTEX_TEXTURE                                             Enum = 0x829B
	MIRROR_CLAMP_EXT                                           Enum = 0x8742
	COMPRESSED_SLUMINANCE_ALPHA                                Enum = 0x8C4B
	UNIFORM_BARRIER_BIT                                        Enum = 0x00000004
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            Enum = 0x8455
	ALPHA_FLOAT32_APPLE                                        Enum = 0x8816
	FORCE_BLUE_TO_ONE_NV                                       Enum = 0x8860
	REG_25_ATI                                                 Enum = 0x893A
	UNSIGNED_INT_SAMPLER_3D_EXT                                Enum = 0x8DD3
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                Enum = 0x818A
	CULL_VERTEX_EYE_POSITION_EXT                               Enum = 0x81AB
	DEBUG_SOURCE_API_ARB                                       Enum = 0x8246
	SCREEN_COORDINATES_REND                                    Enum = 0x8490
	COMPRESSED_LUMINANCE                                       Enum = 0x84EA
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            Enum = 0x851A
	NOOP                                                       Enum = 0x1505
	UNSIGNED_SHORT_5_5_5_1_EXT                                 Enum = 0x8034
	COLOR_ARRAY_POINTER_EXT                                    Enum = 0x8090
	RG                                                         Enum = 0x8227
	DEPTH_PASS_INSTRUMENT_SGIX                                 Enum = 0x8310
	DSDT_NV                                                    Enum = 0x86F5
	ALPHA32F_ARB                                               Enum = 0x8816
	BEVEL_NV                                                   Enum = 0x90A6
	STATIC_READ                                                Enum = 0x88E5
	MAX_LIST_NESTING                                           Enum = 0x0B31
	GL_2PASS_1_SGIS                                            Enum = 0x80A3
	IMAGE_ROTATE_ANGLE_HP                                      Enum = 0x8159
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             Enum = 0x840B
	BACK_NORMALS_HINT_PGI                                      Enum = 0x1A223
	LINEAR_DETAIL_ALPHA_SGIS                                   Enum = 0x8098
	TEXTURE28_ARB                                              Enum = 0x84DC
	REG_15_ATI                                                 Enum = 0x8930
	PIXEL_MAP_I_TO_I_SIZE                                      Enum = 0x0CB0
	FLOAT_MAT4x2                                               Enum = 0x8B69
	TEXTURE_INTENSITY_TYPE_ARB                                 Enum = 0x8C15
	SLUMINANCE8_EXT                                            Enum = 0x8C47
	RELATIVE_HORIZONTAL_LINE_TO_NV                             Enum = 0x07
	INT                                                        Enum = 0x1404
	EVAL_VERTEX_ATTRIB11_NV                                    Enum = 0x86D1
	TEXTURE_BUFFER_ARB                                         Enum = 0x8C2A
	MAX_COMBINED_IMAGE_UNIFORMS                                Enum = 0x90CF
	FOG_FUNC_POINTS_SGIS                                       Enum = 0x812B
	TEXTURE29_ARB                                              Enum = 0x84DD
	WIDE_LINE_HINT_PGI                                         Enum = 0x1A222
	BLEND                                                      Enum = 0x0BE2
	COMBINER_CD_OUTPUT_NV                                      Enum = 0x854B
	VERTEX_SHADER_OPTIMIZED_EXT                                Enum = 0x87D4
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           Enum = 0x88B4
	VERTEX_ARRAY_OBJECT_EXT                                    Enum = 0x9154
	MAP_UNSYNCHRONIZED_BIT_EXT                                 Enum = 0x0020
	INCR_WRAP_EXT                                              Enum = 0x8507
	INVERSE_NV                                                 Enum = 0x862B
	CON_18_ATI                                                 Enum = 0x8953
	TEXTURE_TYPE_QCOM                                          Enum = 0x8BD7
	COMPRESSED_SRGB_ALPHA_EXT                                  Enum = 0x8C49
	UNSIGNED_INT_VEC4                                          Enum = 0x8DC8
	DOUBLE_MAT3x2                                              Enum = 0x8F4B
	UNSIGNED_INT16_VEC3_NV                                     Enum = 0x8FF2
	SGIX_pixel_texture                                         Enum = 1
	PIXEL_TEX_GEN_SGIX                                         Enum = 0x8139
	FLOAT_MAT2x3                                               Enum = 0x8B65
	VIDEO_BUFFER_BINDING_NV                                    Enum = 0x9021
	MATRIX_MODE                                                Enum = 0x0BA0
	OR_REVERSE                                                 Enum = 0x150B
	PROGRAM_BINARY_FORMATS_OES                                 Enum = 0x87FF
	MAX_DEBUG_MESSAGE_LENGTH                                   Enum = 0x9143
	ONE                                                        Enum = 1
	TEXTURE_RECTANGLE_ARB                                      Enum = 0x84F5
	OUTPUT_TEXTURE_COORD16_EXT                                 Enum = 0x87AD
	IMAGE_2D                                                   Enum = 0x904D
	MITER_TRUNCATE_NV                                          Enum = 0x90A8
	GEOMETRY_DEFORMATION_BIT_SGIX                              Enum = 0x00000002
	NAND                                                       Enum = 0x150E
	IMAGE_CLASS_10_10_10_2                                     Enum = 0x82C3
	LUMINANCE_FLOAT16_APPLE                                    Enum = 0x881E
	MAX_COMPUTE_WORK_GROUP_COUNT                               Enum = 0x91BE
	TEXTURE_MAX_LOD_SGIS                                       Enum = 0x813B
	MAP1_VERTEX_ATTRIB15_4_NV                                  Enum = 0x866F
	FLOAT16_VEC3_NV                                            Enum = 0x8FFA
	PROGRAM_OUTPUT                                             Enum = 0x92E4
	ALPHA_MIN_SGIX                                             Enum = 0x8320
	CURRENT_MATRIX_STACK_DEPTH_ARB                             Enum = 0x8640
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      Enum = 0x8D56
	UNPACK_COMPRESSED_BLOCK_SIZE                               Enum = 0x912A
	SCISSOR_TEST                                               Enum = 0x0C11
	INTENSITY8                                                 Enum = 0x804B
	NUM_FRAGMENT_REGISTERS_ATI                                 Enum = 0x896E
	POINT_SIZE_ARRAY_TYPE_OES                                  Enum = 0x898A
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     Enum = 0x93DB
	EXT_copy_texture                                           Enum = 1
	INT16_NV                                                   Enum = 0x8FE4
	TRIANGLES                                                  Enum = 0x0004
	MAX_3D_TEXTURE_SIZE                                        Enum = 0x8073
	COLOR_TABLE_WIDTH_SGI                                      Enum = 0x80D9
	RGB_FLOAT32_APPLE                                          Enum = 0x8815
	BUFFER_UPDATE_BARRIER_BIT                                  Enum = 0x00000200
	RGBA8_EXT                                                  Enum = 0x8058
	TEXTURE_LIGHT_EXT                                          Enum = 0x8350
	TEXTURE4_ARB                                               Enum = 0x84C4
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       Enum = 0x8E8E
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        Enum = 0x9025
	DST_ALPHA                                                  Enum = 0x0304
	GL_4PASS_0_EXT                                             Enum = 0x80A4
	BLEND_SRC_ALPHA                                            Enum = 0x80CB
	SPRITE_AXIAL_SGIX                                          Enum = 0x814C
	NEAREST_CLIPMAP_LINEAR_SGIX                                Enum = 0x844E
	INTERNALFORMAT_GREEN_TYPE                                  Enum = 0x8279
	STORAGE_SHARED_APPLE                                       Enum = 0x85BF
	STENCIL_BACK_FUNC                                          Enum = 0x8800
	SKIP_COMPONENTS3_NV                                        Enum = -4
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            Enum = 0x8F31
	LESS                                                       Enum = 0x0201
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            Enum = 0x8516
	MODELVIEW9_ARB                                             Enum = 0x8729
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            Enum = 0x88A2
	SGIX_fragment_lighting                                     Enum = 1
	TEXTURE_GEN_R                                              Enum = 0x0C62
	TEXTURE_GEN_MODE                                           Enum = 0x2500
	CONVOLUTION_2D_EXT                                         Enum = 0x8011
	ACTIVE_PROGRAM_EXT                                         Enum = 0x8B8D
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        Enum = 0x8DDF
	DOUBLE_MAT3                                                Enum = 0x8F47
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        Enum = 0x92DA
	RG8UI                                                      Enum = 0x8238
	TEXTURE_MAX_CLAMP_S_SGIX                                   Enum = 0x8369
	DOT_PRODUCT_NV                                             Enum = 0x86EC
	STENCIL_BACK_FAIL                                          Enum = 0x8801
	PATCH_VERTICES                                             Enum = 0x8E72
	FONT_X_MIN_BOUNDS_BIT_NV                                   Enum = 0x00010000
	TEXTURE_CUBE_MAP_ARB                                       Enum = 0x8513
	WRITE_ONLY                                                 Enum = 0x88B9
	GEOMETRY_VERTICES_OUT_ARB                                  Enum = 0x8DDA
	ATOMIC_COUNTER_BARRIER_BIT                                 Enum = 0x00001000
	SAMPLE_ALPHA_TO_ONE                                        Enum = 0x809F
	TESS_GEN_VERTEX_ORDER                                      Enum = 0x8E78
	FUNC_REVERSE_SUBTRACT                                      Enum = 0x800B
	LOCAL_CONSTANT_VALUE_EXT                                   Enum = 0x87EC
	DEPTH_COMPONENT32F_NV                                      Enum = 0x8DAB
	QUERY_BY_REGION_WAIT                                       Enum = 0x8E15
	RELATIVE_LARGE_CCW_ARC_TO_NV                               Enum = 0x17
	SHADER_STORAGE_BUFFER                                      Enum = 0x90D2
	UNSIGNALED_APPLE                                           Enum = 0x9118
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        Enum = 0x92C6
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               Enum = 0x8B89
	IMAGE_CUBE_EXT                                             Enum = 0x9050
	MODELVIEW                                                  Enum = 0x1700
	BGR_EXT                                                    Enum = 0x80E0
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             Enum = 0x84F8
	VIDEO_BUFFER_PITCH_NV                                      Enum = 0x9028
	MAX_FRAGMENT_IMAGE_UNIFORMS                                Enum = 0x90CE
	DEPTH_EXT                                                  Enum = 0x1801
	LIGHT6                                                     Enum = 0x4006
	FUNC_REVERSE_SUBTRACT_OES                                  Enum = 0x800B
	LUMINANCE32F_ARB                                           Enum = 0x8818
	CON_7_ATI                                                  Enum = 0x8948
	COLOR_FLOAT_APPLE                                          Enum = 0x8A0F
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   Enum = 0x9034
	INT_IMAGE_2D_RECT                                          Enum = 0x905A
	FRAGMENT_SHADER_BIT                                        Enum = 0x00000002
	DOT_PRODUCT_DEPTH_REPLACE_NV                               Enum = 0x86ED
	BUFFER_MAPPED                                              Enum = 0x88BC
	INVALID_VALUE                                              Enum = 0x0501
	MINMAX_FORMAT_EXT                                          Enum = 0x802F
	SAMPLES_ARB                                                Enum = 0x80A9
	TIME_ELAPSED_EXT                                           Enum = 0x88BF
	RESAMPLE_DECIMATE_OML                                      Enum = 0x8989
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     Enum = 0x8A31
	INT_SAMPLER_BUFFER_EXT                                     Enum = 0x8DD0
	DOUBLE_MAT4x3                                              Enum = 0x8F4E
	UTF16_NV                                                   Enum = 0x909B
	VIRTUAL_PAGE_SIZE_Y_AMD                                    Enum = 0x9196
	MOV_ATI                                                    Enum = 0x8961
	HIGH_INT                                                   Enum = 0x8DF5
	MAX_VERTEX_IMAGE_UNIFORMS                                  Enum = 0x90CA
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              Enum = 0x9144
	PROXY_TEXTURE_2D                                           Enum = 0x8064
	SAMPLE_MASK_VALUE_EXT                                      Enum = 0x80AA
	TEXTURE13                                                  Enum = 0x84CD
	MAX_VERTEX_UNITS_ARB                                       Enum = 0x86A4
	VERTEX_ATTRIB_ARRAY_DIVISOR                                Enum = 0x88FE
	FLOAT16_VEC2_NV                                            Enum = 0x8FF9
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        Enum = 0x910A
	VERTEX_ARRAY_SIZE_EXT                                      Enum = 0x807A
	INSTRUMENT_MEASUREMENTS_SGIX                               Enum = 0x8181
	RGB9_E5                                                    Enum = 0x8C3D
	INT64_VEC4_NV                                              Enum = 0x8FEB
	BLEND_EQUATION_RGB_OES                                     Enum = 0x8009
	UNDEFINED_VERTEX                                           Enum = 0x8260
	PROGRAM_STRING_ARB                                         Enum = 0x8628
	PROGRAM_NATIVE_ATTRIBS_ARB                                 Enum = 0x88AE
	LUMINANCE_ALPHA16UI_EXT                                    Enum = 0x8D7B
	UNSIGNED_INT64_VEC3_NV                                     Enum = 0x8FF6
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          Enum = 0x92CC
	MAP1_GRID_SEGMENTS                                         Enum = 0x0DD1
	VERTEX_PROGRAM_POINT_SIZE                                  Enum = 0x8642
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               Enum = 0x8853
	LINE_STIPPLE_PATTERN                                       Enum = 0x0B25
	ASYNC_DRAW_PIXELS_SGIX                                     Enum = 0x835D
	R11F_G11F_B10F                                             Enum = 0x8C3A
	FRAMEBUFFER_EXT                                            Enum = 0x8D40
	FOG_COORD_ARRAY_ADDRESS_NV                                 Enum = 0x8F28
	OBJECT_TYPE_APPLE                                          Enum = 0x9112
	MATRIX_INDEX_ARRAY_STRIDE_OES                              Enum = 0x8848
	PIXEL_UNPACK_BUFFER_BINDING                                Enum = 0x88EF
	TRANSLATE_2D_NV                                            Enum = 0x9090
	VERTEX_ATTRIB_ARRAY14_NV                                   Enum = 0x865E
	OUTPUT_VERTEX_EXT                                          Enum = 0x879A
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         Enum = 0x8CDC
	SOURCE0_ALPHA_ARB                                          Enum = 0x8588
	EVAL_VERTEX_ATTRIB6_NV                                     Enum = 0x86CC
	QUERY_RESULT                                               Enum = 0x8866
	RESAMPLE_REPLICATE_OML                                     Enum = 0x8986
	TEXTURE_WRAP_S                                             Enum = 0x2802
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         Enum = 0x889A
	TRANSPOSE_AFFINE_2D_NV                                     Enum = 0x9096
	STENCIL_PASS_DEPTH_PASS                                    Enum = 0x0B96
	PACK_IMAGE_HEIGHT_EXT                                      Enum = 0x806C
	COLOR_TABLE_SCALE_SGI                                      Enum = 0x80D6
	POINT_SIZE_MAX_SGIS                                        Enum = 0x8127
	UNPACK_CONSTANT_DATA_SUNX                                  Enum = 0x81D5
	MAX_RENDERBUFFER_SIZE                                      Enum = 0x84E8
	COMPRESSED_LUMINANCE_ALPHA_ARB                             Enum = 0x84EB
	COMBINE_ALPHA_EXT                                          Enum = 0x8572
	MAT_EMISSION_BIT_PGI                                       Enum = 0x00800000
	SGIX_async_histogram                                       Enum = 1
	TEXTURE_BIT                                                Enum = 0x00040000
	LINE_WIDTH_GRANULARITY                                     Enum = 0x0B23
	POINT_SPRITE_ARB                                           Enum = 0x8861
	ADD_BLEND_IMG                                              Enum = 0x8C09
	LUMINANCE32UI_EXT                                          Enum = 0x8D74
	MAX_NUM_ACTIVE_VARIABLES                                   Enum = 0x92F7
	T2F_C4UB_V3F                                               Enum = 0x2A29
	LUMINANCE8_EXT                                             Enum = 0x8040
	INTENSITY12_EXT                                            Enum = 0x804C
	Z_EXT                                                      Enum = 0x87D7
	TEXTURE_STACK_DEPTH                                        Enum = 0x0BA5
	TEXTURE_BINDING_1D                                         Enum = 0x8068
	CLEAR                                                      Enum = 0x1500
	MATRIX29_ARB                                               Enum = 0x88DD
	EXP2                                                       Enum = 0x0801
	PIXEL_MAP_I_TO_B                                           Enum = 0x0C74
	RESCALE_NORMAL_EXT                                         Enum = 0x803A
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            Enum = 0x8DBC
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       Enum = 0x8E58
	SMOOTH_POINT_SIZE_RANGE                                    Enum = 0x0B12
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                Enum = 0x851A
	SEPARATE_ATTRIBS                                           Enum = 0x8C8D
	HISTOGRAM_LUMINANCE_SIZE                                   Enum = 0x802C
	TEXTURE_ENV_BIAS_SGIX                                      Enum = 0x80BE
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            Enum = 0x813F
	RG32UI                                                     Enum = 0x823C
	CONSTANT_COLOR1_NV                                         Enum = 0x852B
	TEXTURE_BLUE_SIZE                                          Enum = 0x805E
	BITMAP                                                     Enum = 0x1A00
	VERTEX_ATTRIB_ARRAY9_NV                                    Enum = 0x8659
	OUTPUT_TEXTURE_COORD20_EXT                                 Enum = 0x87B1
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    Enum = 0x8810
	REG_3_ATI                                                  Enum = 0x8924
	CON_21_ATI                                                 Enum = 0x8956
	PALETTE8_R5_G6_B5_OES                                      Enum = 0x8B97
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          Enum = 0x8E5B
	INTERLACE_READ_INGR                                        Enum = 0x8568
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   Enum = 0x9036
	NORMAL_ARRAY_LIST_STRIDE_IBM                               Enum = 103081
	PACK_CMYK_HINT_EXT                                         Enum = 0x800E
	CLAMP_TO_BORDER                                            Enum = 0x812D
	RESAMPLE_ZERO_FILL_SGIX                                    Enum = 0x842F
	TEXTURE_MAG_SIZE_NV                                        Enum = 0x871F
	MATRIX16_ARB                                               Enum = 0x88D0
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      Enum = 0x8B9E
	CONTEXT_CORE_PROFILE_BIT                                   Enum = 0x00000001
	ARRAY_OBJECT_BUFFER_ATI                                    Enum = 0x8766
	BUFFER_ACCESS                                              Enum = 0x88BB
	MAX_VERTEX_UNIFORM_BLOCKS                                  Enum = 0x8A2B
	RENDERBUFFER_OES                                           Enum = 0x8D41
	ALPHA8_SNORM                                               Enum = 0x9014
	AUX0                                                       Enum = 0x0409
	LINE_STIPPLE                                               Enum = 0x0B24
	TEXTURE_MAX_LEVEL                                          Enum = 0x813D
	VIEW_CLASS_RGTC2_RG                                        Enum = 0x82D1
	RGB_FLOAT16_APPLE                                          Enum = 0x881B
	TEXTURE_RESIDENT                                           Enum = 0x8067
	LIGHT5                                                     Enum = 0x4005
	RG_EXT                                                     Enum = 0x8227
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            Enum = 0x851F
	ATTRIB_ARRAY_POINTER_NV                                    Enum = 0x8645
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         Enum = 0x8B9A
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            Enum = 0x9068
	TEXTURE_COORD_ARRAY_LIST_IBM                               Enum = 103074
	UNIFORM_BUFFER_START                                       Enum = 0x8A29
	LARGE_CW_ARC_TO_NV                                         Enum = 0x18
	TEXTURE_FILTER_CONTROL                                     Enum = 0x8500
	PROGRAM_POINT_SIZE_EXT                                     Enum = 0x8642
	PRESERVE_ATI                                               Enum = 0x8762
	MATRIX26_ARB                                               Enum = 0x88DA
	DOUBLE_MAT4                                                Enum = 0x8F48
	MAP_INVALIDATE_BUFFER_BIT                                  Enum = 0x0008
	ONE_MINUS_SRC_ALPHA                                        Enum = 0x0303
	RGB_S3TC                                                   Enum = 0x83A0
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           Enum = 0x9138
	UNSIGNED_BYTE_3_3_2_EXT                                    Enum = 0x8032
	LIST_PRIORITY_SGIX                                         Enum = 0x8182
	YCRCBA_SGIX                                                Enum = 0x8319
	RGBA16UI                                                   Enum = 0x8D76
	LUMINANCE8_ALPHA8_SNORM                                    Enum = 0x9016
	TOP_LEVEL_ARRAY_STRIDE                                     Enum = 0x930D
	SHININESS                                                  Enum = 0x1601
	ALPHA_MAX_SGIX                                             Enum = 0x8321
	SAMPLER_3D                                                 Enum = 0x8B5F
	FRACTIONAL_EVEN                                            Enum = 0x8E7C
	PIXEL_MAP_I_TO_G                                           Enum = 0x0C73
	FOG_COORDINATE_ARRAY_POINTER_EXT                           Enum = 0x8456
	CURRENT_VERTEX_ATTRIB_ARB                                  Enum = 0x8626
	LOCAL_EXT                                                  Enum = 0x87C4
	MAX_TEXTURE_IMAGE_UNITS_ARB                                Enum = 0x8872
	TEXTURE_BORDER                                             Enum = 0x1005
	BGRA                                                       Enum = 0x80E1
	VERTEX_ARRAY_RANGE_POINTER_NV                              Enum = 0x8521
	R1UI_T2F_C4F_N3F_V3F_SUN                                   Enum = 0x85CB
	DOT3_RGBA_IMG                                              Enum = 0x86AF
	DRAW_FRAMEBUFFER_EXT                                       Enum = 0x8CA9
	MAX_FRAMEBUFFER_WIDTH                                      Enum = 0x9315
	PRIMARY_COLOR_EXT                                          Enum = 0x8577
	REPLACEMENT_CODE_ARRAY_SUN                                 Enum = 0x85C0
	DRAW_PIXELS_APPLE                                          Enum = 0x8A0A
	VOLATILE_APPLE                                             Enum = 0x8A1A
	RG16_SNORM                                                 Enum = 0x8F99
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       Enum = 0x8F39
	FLOAT_RG16_NV                                              Enum = 0x8886
	RGB32I_EXT                                                 Enum = 0x8D83
	LUMINANCE16_SNORM                                          Enum = 0x9019
	VERTEX_ARRAY_RANGE_LENGTH_NV                               Enum = 0x851E
	REG_17_ATI                                                 Enum = 0x8932
	FOG_COORD_ARRAY_TYPE                                       Enum = 0x8454
	MAP2_VERTEX_ATTRIB13_4_NV                                  Enum = 0x867D
	DRAW_BUFFER12                                              Enum = 0x8831
	Z4Y12Z4CB12Z4CR12_444_NV                                   Enum = 0x9037
	FILE_NAME_NV                                               Enum = 0x9074
	SYNC_FLAGS_APPLE                                           Enum = 0x9115
	PROGRAM_BINDING_ARB                                        Enum = 0x8677
	DOUBLE_MAT4x2_EXT                                          Enum = 0x8F4D
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 Enum = 0x0010
	CLAMP_VERTEX_COLOR_ARB                                     Enum = 0x891A
	GEOMETRY_SHADER_ARB                                        Enum = 0x8DD9
	INT_VEC3_ARB                                               Enum = 0x8B54
	DEBUG_CATEGORY_APPLICATION_AMD                             Enum = 0x914F
	TEXTURE_FETCH_BARRIER_BIT_EXT                              Enum = 0x00000008
	MAX_COLOR_MATRIX_STACK_DEPTH                               Enum = 0x80B3
	POINT_SIZE                                                 Enum = 0x0B11
	TEXTURE_BINDING_CUBE_MAP_OES                               Enum = 0x8514
	MODELVIEW2_ARB                                             Enum = 0x8722
	ETC1_SRGB8_NV                                              Enum = 0x88EE
	RETAINED_APPLE                                             Enum = 0x8A1B
	COMBINE_RGB                                                Enum = 0x8571
	VERTEX_ATTRIB_ARRAY_STRIDE                                 Enum = 0x8624
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         Enum = 0x8D56
	TEXCOORD3_BIT_PGI                                          Enum = 0x40000000
	SGIX_ir_instrument1                                        Enum = 1
	EVAL_VERTEX_ATTRIB0_NV                                     Enum = 0x86C6
	TEXTURE_COMPARE_FUNC_ARB                                   Enum = 0x884D
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          Enum = 0x8E5C
	GL_4PASS_3_EXT                                             Enum = 0x80A7
	DOUBLE_VEC2                                                Enum = 0x8FFC
	SUCCESS_NV                                                 Enum = 0x902F
	DOUBLE_EXT                                                 Enum = 0x140A
	T4F_C4F_N3F_V4F                                            Enum = 0x2A2D
	MINMAX_EXT                                                 Enum = 0x802E
	TEXTURE_CUBE_MAP_POSITIVE_Y                                Enum = 0x8517
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            Enum = 0x8850
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      Enum = 0x8CD6
	INDEX_ARRAY_LIST_STRIDE_IBM                                Enum = 103083
	COLOR_SUM_CLAMP_NV                                         Enum = 0x854F
	TEXTURE_RED_TYPE_ARB                                       Enum = 0x8C10
	BACK_SECONDARY_COLOR_NV                                    Enum = 0x8C78
	DEPTH_CLAMP_NV                                             Enum = 0x864F
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             Enum = 0x8C85
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             Enum = 0x8F1E
	PACK_SWAP_BYTES                                            Enum = 0x0D00
	MAP2_INDEX                                                 Enum = 0x0DB1
	PACK_COMPRESSED_SIZE_SGIX                                  Enum = 0x831C
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       Enum = 0x8F44
	FRAMEBUFFER_ATTACHMENT_ANGLE                               Enum = 0x93A3
	SGIX_async_pixel                                           Enum = 1
	ALL_SHADER_BITS                                            Enum = 0xFFFFFFFF
	OBJECT_PLANE                                               Enum = 0x2501
	GENERATE_MIPMAP_HINT                                       Enum = 0x8192
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       Enum = 0x8A52
	COLOR_ATTACHMENT12                                         Enum = 0x8CEC
	GEOMETRY_SHADER_BIT                                        Enum = 0x00000004
	ALPHA32UI_EXT                                              Enum = 0x8D72
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     Enum = 0x90DF
	FEEDBACK                                                   Enum = 0x1C01
	OPERAND1_ALPHA_EXT                                         Enum = 0x8599
	UNSIGNED_INT_10_10_10_2_OES                                Enum = 0x8DF6
	COVERAGE_SAMPLES_NV                                        Enum = 0x8ED4
	COPY_READ_BUFFER                                           Enum = 0x8F36
	TOP_LEVEL_ARRAY_SIZE                                       Enum = 0x930C
	PHONG_WIN                                                  Enum = 0x80EA
	VERTEX_STREAM3_ATI                                         Enum = 0x876F
	ARRAY_BUFFER                                               Enum = 0x8892
	LERP_ATI                                                   Enum = 0x8969
	BLUE_BIT_ATI                                               Enum = 0x00000004
	IMPLEMENTATION_COLOR_READ_FORMAT                           Enum = 0x8B9B
	LUMINANCE_ALPHA                                            Enum = 0x190A
	MAX_COLOR_ATTACHMENTS_EXT                                  Enum = 0x8CDF
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             Enum = 0x8262
	BLUE_MIN_CLAMP_INGR                                        Enum = 0x8562
	BLEND_EQUATION                                             Enum = 0x8009
	COMPARE_REF_TO_TEXTURE                                     Enum = 0x884E
	COMPUTE_LOCAL_WORK_SIZE                                    Enum = 0x8267
	EXPAND_NEGATE_NV                                           Enum = 0x8539
	SAMPLER_2D_SHADOW                                          Enum = 0x8B62
	TEXTURE20_ARB                                              Enum = 0x84D4
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     Enum = 0x8403
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            Enum = 0x8622
	OP_MADD_EXT                                                Enum = 0x8788
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             Enum = 0x886A
	CON_27_ATI                                                 Enum = 0x895C
	RENDERBUFFER_WIDTH_EXT                                     Enum = 0x8D42
	ONE_MINUS_CONSTANT_COLOR_EXT                               Enum = 0x8002
	TEXTURE_GATHER_SHADOW                                      Enum = 0x82A3
	BUMP_ENVMAP_ATI                                            Enum = 0x877B
	TEXTURE_CROP_RECT_OES                                      Enum = 0x8B9D
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      Enum = 0x8C88
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          Enum = 0x90DB
	EXP                                                        Enum = 0x0800
	TRACE_MASK_MESA                                            Enum = 0x8755
	BUFFER_MAP_LENGTH                                          Enum = 0x9120
	MATRIX_INDEX_ARRAY_ARB                                     Enum = 0x8844
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          Enum = 0x909C
	BUFFER_DATA_SIZE                                           Enum = 0x9303
	RED_BITS                                                   Enum = 0x0D52
	BINORMAL_ARRAY_POINTER_EXT                                 Enum = 0x8443
	DRAW_BUFFER4_ATI                                           Enum = 0x8829
	READ_PIXEL_DATA_RANGE_NV                                   Enum = 0x8879
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     Enum = 0x8D56
	MAX_SERVER_WAIT_TIMEOUT                                    Enum = 0x9111
	VERTEX_ATTRIB_ARRAY0_NV                                    Enum = 0x8650
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            Enum = 0x88FE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               Enum = 0x8CD3
	OBJECT_TYPE                                                Enum = 0x9112
	MAP1_NORMAL                                                Enum = 0x0D92
	INTERNALFORMAT_STENCIL_SIZE                                Enum = 0x8276
	R1UI_T2F_N3F_V3F_SUN                                       Enum = 0x85CA
	BOOL_VEC4_ARB                                              Enum = 0x8B59
	SLUMINANCE8_NV                                             Enum = 0x8C47
	SGIX_list_priority                                         Enum = 1
	AND_REVERSE                                                Enum = 0x1502
	BLEND_SRC_RGB                                              Enum = 0x80C9
	TEXTURE8_ARB                                               Enum = 0x84C8
	OPERAND1_RGB_ARB                                           Enum = 0x8591
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            Enum = 0x8645
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          Enum = 0x8BB7
	PRIMITIVE_ID_NV                                            Enum = 0x8C7C
	IMAGE_2D_MULTISAMPLE_EXT                                   Enum = 0x9055
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               Enum = 0x9143
	FULL_SUPPORT                                               Enum = 0x82B7
	MAX_TEXTURE_UNITS_ARB                                      Enum = 0x84E2
	PIXEL_SUBSAMPLE_2424_SGIX                                  Enum = 0x85A3
	PATH_FILL_MODE_NV                                          Enum = 0x9080
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     Enum = 0x9134
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             Enum = 0x80DE
	BGR                                                        Enum = 0x80E0
	COMBINER2_NV                                               Enum = 0x8552
	EVAL_VERTEX_ATTRIB5_NV                                     Enum = 0x86CB
	PROGRAM_FORMAT_ARB                                         Enum = 0x8876
	SAMPLER_CUBE_MAP_ARRAY                                     Enum = 0x900C
	FRAGMENT_PROGRAM_ARB                                       Enum = 0x8804
	MATRIX24_ARB                                               Enum = 0x88D8
	ARRAY_OBJECT_OFFSET_ATI                                    Enum = 0x8767
	ALWAYS_FAST_HINT_PGI                                       Enum = 0x1A20C
	COLOR_ARRAY_POINTER                                        Enum = 0x8090
	COLOR_ARRAY_SIZE                                           Enum = 0x8081
	CLIP_PLANE4                                                Enum = 0x3004
	YCRCB_SGIX                                                 Enum = 0x8318
	SOURCE1_ALPHA                                              Enum = 0x8589
	RGBA16F_EXT                                                Enum = 0x881A
	STENCIL_INDEX16_EXT                                        Enum = 0x8D49
	DUP_FIRST_CUBIC_CURVE_TO_NV                                Enum = 0xF2
	SELECT                                                     Enum = 0x1C02
	GL_3DC_X_AMD                                               Enum = 0x87F9
	COLOR_ATTACHMENT0_EXT                                      Enum = 0x8CE0
	TEXTURE_GEN_STR_OES                                        Enum = 0x8D60
	DEPTH_COMPONENT16_NONLINEAR_NV                             Enum = 0x8E2C
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          Enum = 0x9026
	VIEW_CLASS_S3TC_DXT1_RGBA                                  Enum = 0x82CD
	TEXTURE9_ARB                                               Enum = 0x84C9
	DSDT_MAG_VIB_NV                                            Enum = 0x86F7
	CURRENT_MATRIX_INDEX_ARB                                   Enum = 0x8845
	MATRIX_PALETTE_OES                                         Enum = 0x8840
	ATC_RGB_AMD                                                Enum = 0x8C92
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          Enum = 0x8E5F
	PATH_GEN_COEFF_NV                                          Enum = 0x90B1
	RGBA12_EXT                                                 Enum = 0x805A
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  Enum = 0x8163
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             Enum = 0x8247
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      Enum = 0x87C8
	TEXTURE_BINDING_2D_ARRAY_EXT                               Enum = 0x8C1D
	TEXTURE_BUFFER_EXT                                         Enum = 0x8C2A
	LOCATION                                                   Enum = 0x930E
	MAX_TEXTURE_SIZE                                           Enum = 0x0D33
	STENCIL_TEST_TWO_SIDE_EXT                                  Enum = 0x8910
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              Enum = 0x8C8A
	UPPER_LEFT                                                 Enum = 0x8CA2
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      Enum = 0x8DBE
	FONT_HEIGHT_BIT_NV                                         Enum = 0x00800000
	SGIX_texture_add_env                                       Enum = 1
	NORMALIZED_RANGE_EXT                                       Enum = 0x87E0
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             Enum = 0x88AF
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     Enum = 0x8CDB
	MAX_FRAMEBUFFER_HEIGHT                                     Enum = 0x9316
	MATERIAL_SIDE_HINT_PGI                                     Enum = 0x1A22C
	TEXTURE_BINDING_2D                                         Enum = 0x8069
	TEXTURE_3D_OES                                             Enum = 0x806F
	MIPMAP                                                     Enum = 0x8293
	VIEW_CLASS_24_BITS                                         Enum = 0x82C9
	SYNC_GPU_COMMANDS_COMPLETE                                 Enum = 0x9117
	FOG_COORDINATE_ARRAY_LIST_IBM                              Enum = 103076
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             Enum = 2
	FRONT_FACE                                                 Enum = 0x0B46
	TEXTURE_LOD_BIAS_R_SGIX                                    Enum = 0x8190
	GUILTY_CONTEXT_RESET_ARB                                   Enum = 0x8253
	MAX_WIDTH                                                  Enum = 0x827E
	GET_TEXTURE_IMAGE_TYPE                                     Enum = 0x8292
	OUTPUT_TEXTURE_COORD0_EXT                                  Enum = 0x879D
	VIRTUAL_PAGE_SIZE_X_AMD                                    Enum = 0x9195
	ASYNC_HISTOGRAM_SGIX                                       Enum = 0x832C
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 Enum = 0x8360
	SECONDARY_COLOR_ARRAY_POINTER                              Enum = 0x845D
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         Enum = 0x8C4E
	INTENSITY8I_EXT                                            Enum = 0x8D91
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           Enum = 0x80B9
	BUFFER_ACCESS_OES                                          Enum = 0x88BB
	UNSIGNED_INT_5_9_9_9_REV_EXT                               Enum = 0x8C3E
	ALPHA_TEST_FUNC_QCOM                                       Enum = 0x0BC1
	LUMINANCE4_ALPHA4                                          Enum = 0x8043
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             Enum = 0x8365
	INDEX_ARRAY_STRIDE_EXT                                     Enum = 0x8086
	GL_422_REV_EXT                                             Enum = 0x80CD
	SUBTRACT_ARB                                               Enum = 0x84E7
	MIRROR_CLAMP_ATI                                           Enum = 0x8742
	DRAW_BUFFER10_ARB                                          Enum = 0x882F
	SAMPLER_BUFFER_EXT                                         Enum = 0x8DC2
	IMAGE_BINDING_LAYERED_EXT                                  Enum = 0x8F3C
	MAX_LIGHTS                                                 Enum = 0x0D31
	CONVOLUTION_FILTER_SCALE                                   Enum = 0x8014
	TEXTURE_COORD_ARRAY_POINTER_EXT                            Enum = 0x8092
	REG_31_ATI                                                 Enum = 0x8940
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      Enum = 0x82AD
	BLUE_MAX_CLAMP_INGR                                        Enum = 0x8566
	PROXY_TEXTURE_3D_EXT                                       Enum = 0x8070
	INTERNALFORMAT_DEPTH_TYPE                                  Enum = 0x827C
	MODELVIEW12_ARB                                            Enum = 0x872C
	PROGRAM_ERROR_STRING_ARB                                   Enum = 0x8874
	CULL_VERTEX_EXT                                            Enum = 0x81AA
	DECR_WRAP_OES                                              Enum = 0x8508
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  Enum = 0x8A44
	OBJECT_LINK_STATUS_ARB                                     Enum = 0x8B82
	DOUBLE_MAT2x4                                              Enum = 0x8F4A
	DEPTH_BIAS                                                 Enum = 0x0D1F
	POST_COLOR_MATRIX_GREEN_SCALE                              Enum = 0x80B5
	PROGRAM_ERROR_POSITION_ARB                                 Enum = 0x864B
	REG_29_ATI                                                 Enum = 0x893E
	COLOR_ALPHA_PAIRING_ATI                                    Enum = 0x8975
	READ_FRAMEBUFFER_BINDING_EXT                               Enum = 0x8CAA
	MAX_VERTEX_OUTPUT_COMPONENTS                               Enum = 0x9122
	CLIENT_ALL_ATTRIB_BITS                                     Enum = 0xFFFFFFFF
	R3_G3_B2                                                   Enum = 0x2A10
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            Enum = 0x83F2
	SAMPLE_SHADING_ARB                                         Enum = 0x8C36
	INTERLEAVED_ATTRIBS                                        Enum = 0x8C8C
	INT_SAMPLER_2D                                             Enum = 0x8DCA
	SMOOTH_CUBIC_CURVE_TO_NV                                   Enum = 0x10
	STENCIL_TEST                                               Enum = 0x0B90
	T2F_V3F                                                    Enum = 0x2A27
	RG_INTEGER                                                 Enum = 0x8228
	PROGRAM_TEMPORARIES_ARB                                    Enum = 0x88A4
	UNIFORM_BLOCK_INDEX                                        Enum = 0x8A3A
	HIGH_FLOAT                                                 Enum = 0x8DF2
	TRANSPOSE_COLOR_MATRIX                                     Enum = 0x84E6
	FONT_Y_MIN_BOUNDS_BIT_NV                                   Enum = 0x00020000
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         Enum = 0x92F0
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            Enum = 103085
	TEXTURE_COORD_ARRAY                                        Enum = 0x8078
	INTENSITY4                                                 Enum = 0x804A
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         Enum = 0x8264
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               Enum = 0x8625
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          Enum = 0x87D2
	FOG_COORD_ARRAY_BUFFER_BINDING                             Enum = 0x889D
	RESAMPLE_AVERAGE_OML                                       Enum = 0x8988
	GEOMETRY_VERTICES_OUT_EXT                                  Enum = 0x8DDA
	SMOOTH_LINE_WIDTH_RANGE                                    Enum = 0x0B22
	COMPRESSED_LUMINANCE_LATC1_EXT                             Enum = 0x8C70
	TEXTURE_3D_BINDING_EXT                                     Enum = 0x806A
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          Enum = 0x8215
	R32F                                                       Enum = 0x822E
	COORD_REPLACE_NV                                           Enum = 0x8862
	SGIS_sharpen_texture                                       Enum = 1
	ACCUM_CLEAR_VALUE                                          Enum = 0x0B80
	FORMAT_SUBSAMPLE_24_24_OML                                 Enum = 0x8982
	CW                                                         Enum = 0x0900
	STENCIL_PASS_DEPTH_FAIL                                    Enum = 0x0B95
	POLYGON_OFFSET_LINE                                        Enum = 0x2A02
	SHADOW_ATTENUATION_EXT                                     Enum = 0x834E
	COMBINER_BIAS_NV                                           Enum = 0x8549
	FLOAT_R32_NV                                               Enum = 0x8885
	FAILURE_NV                                                 Enum = 0x9030
	TEXTURE17                                                  Enum = 0x84D1
	OP_MAX_EXT                                                 Enum = 0x878A
	FLOAT_R_NV                                                 Enum = 0x8880
	MAX_TEXTURE_BUFFER_SIZE                                    Enum = 0x8C2B
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        Enum = 0x8E5A
	UNSIGNED_INT8_VEC3_NV                                      Enum = 0x8FEE
	COMPRESSED_RGBA_ASTC_8x6_KHR                               Enum = 0x93B6
	INCR_WRAP_OES                                              Enum = 0x8507
	TEXTURE_CUBE_MAP_POSITIVE_Z                                Enum = 0x8519
	ELEMENT_ARRAY_ATI                                          Enum = 0x8768
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        Enum = 0x8B9F
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        Enum = 0x90BD
	FOG_END                                                    Enum = 0x0B64
	PROXY_COLOR_TABLE                                          Enum = 0x80D3
	EYE_DISTANCE_TO_LINE_SGIS                                  Enum = 0x81F2
	MATRIX6_NV                                                 Enum = 0x8636
	TRANSFORM_FEEDBACK_BUFFER_EXT                              Enum = 0x8C8E
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       Enum = 0x8E5B
	SIGNED_NORMALIZED                                          Enum = 0x8F9C
	INT_IMAGE_BUFFER_EXT                                       Enum = 0x905C
	PIXEL_MAP_I_TO_R_SIZE                                      Enum = 0x0CB2
	INDEX_OFFSET                                               Enum = 0x0D13
	HALF_FLOAT_NV                                              Enum = 0x140B
	RGB8_EXT                                                   Enum = 0x8051
	INTERNALFORMAT_RED_SIZE                                    Enum = 0x8271
	MATRIX31_ARB                                               Enum = 0x88DF
	RGBA8UI                                                    Enum = 0x8D7C
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          Enum = 0x8E11
	VERTEX_SHADER_EXT                                          Enum = 0x8780
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            Enum = 0x87CC
	SLUMINANCE_ALPHA                                           Enum = 0x8C44
	RENDERBUFFER_RED_SIZE                                      Enum = 0x8D50
	FONT_UNITS_PER_EM_BIT_NV                                   Enum = 0x00100000
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              Enum = 0x00000002
	GLYPH_HEIGHT_BIT_NV                                        Enum = 0x02
	MAP1_VERTEX_3                                              Enum = 0x0D97
	SLUMINANCE8_ALPHA8_EXT                                     Enum = 0x8C45
	LIGHT1                                                     Enum = 0x4001
	DEBUG_TYPE_PERFORMANCE_ARB                                 Enum = 0x8250
	SPARE1_NV                                                  Enum = 0x852F
	OPERAND1_ALPHA                                             Enum = 0x8599
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         Enum = 0x87F4
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    Enum = 0x880E
	GEOMETRY_SHADER_EXT                                        Enum = 0x8DD9
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 Enum = 0x8F26
	SAMPLER_2D_MULTISAMPLE_ARRAY                               Enum = 0x910B
	BACK_LEFT                                                  Enum = 0x0402
	RASTERIZER_DISCARD                                         Enum = 0x8C89
	PATH_CLIENT_LENGTH_NV                                      Enum = 0x907F
	CURRENT_COLOR                                              Enum = 0x0B00
	CON_0_ATI                                                  Enum = 0x8941
	NUM_FILL_STREAMS_NV                                        Enum = 0x8E29
	MAX_IMAGE_UNITS_EXT                                        Enum = 0x8F38
	SMALL_CCW_ARC_TO_NV                                        Enum = 0x12
	SGIX_reference_plane                                       Enum = 1
	SAMPLE_ALPHA_TO_ONE_EXT                                    Enum = 0x809F
	SHADER_IMAGE_STORE                                         Enum = 0x82A5
	FRAMEBUFFER_DEFAULT_WIDTH                                  Enum = 0x9310
	NORMAL_ARRAY_COUNT_EXT                                     Enum = 0x8080
	RENDERBUFFER_FREE_MEMORY_ATI                               Enum = 0x87FD
	TEXTURE_COMPARE_MODE                                       Enum = 0x884C
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       Enum = 0x8B9B
	TEXTURE_ALPHA_TYPE                                         Enum = 0x8C13
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     Enum = 0x87F1
	NAMED_STRING_LENGTH_ARB                                    Enum = 0x8DE9
	PROVOKING_VERTEX_EXT                                       Enum = 0x8E4F
	QUERY_BUFFER_BINDING_AMD                                   Enum = 0x9193
	CURRENT_RASTER_INDEX                                       Enum = 0x0B05
	DEBUG_TYPE_PERFORMANCE                                     Enum = 0x8250
	MIN_PROGRAM_TEXEL_OFFSET_NV                                Enum = 0x8904
	DELETE_STATUS                                              Enum = 0x8B80
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            Enum = 0x8C00
	LOW_FLOAT                                                  Enum = 0x8DF0
	DEBUG_SEVERITY_LOW_AMD                                     Enum = 0x9148
	PIXEL_MAP_G_TO_G                                           Enum = 0x0C77
	CONVOLUTION_FILTER_SCALE_EXT                               Enum = 0x8014
	PACK_RESAMPLE_SGIX                                         Enum = 0x842C
	OPERAND0_RGB                                               Enum = 0x8590
	BUFFER_USAGE                                               Enum = 0x8765
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           Enum = 0x8C02
	COVERAGE_BUFFER_BIT_NV                                     Enum = 0x00008000
	CUBIC_CURVE_TO_NV                                          Enum = 0x0C
	POINTS                                                     Enum = 0x0000
	CURRENT_RASTER_POSITION                                    Enum = 0x0B07
	TEXTURE_RED_SIZE                                           Enum = 0x805C
	WEIGHT_ARRAY_OES                                           Enum = 0x86AD
	RENDERBUFFER_ALPHA_SIZE_OES                                Enum = 0x8D53
	VIRTUAL_PAGE_SIZE_Z_AMD                                    Enum = 0x9197
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        Enum = 0x817B
	OP_NEGATE_EXT                                              Enum = 0x8783
	UNSIGNED_SHORT_5_6_5_REV                                   Enum = 0x8364
	OP_ADD_EXT                                                 Enum = 0x8787
	ALPHA8I_EXT                                                Enum = 0x8D90
	ELEMENT_ARRAY_ADDRESS_NV                                   Enum = 0x8F29
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                Enum = 0x906C
	COLOR_ATTACHMENT2_EXT                                      Enum = 0x8CE2
	RGBA_INTEGER_EXT                                           Enum = 0x8D99
	COLOR_ARRAY_LENGTH_NV                                      Enum = 0x8F2D
	EDGE_FLAG_ARRAY_LENGTH_NV                                  Enum = 0x8F30
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             Enum = 0x02000000
	SAMPLE_COVERAGE                                            Enum = 0x80A0
	WRITE_DISCARD_NV                                           Enum = 0x88BE
	COLOR_TABLE_ALPHA_SIZE_SGI                                 Enum = 0x80DD
	POST_TEXTURE_FILTER_BIAS_SGIX                              Enum = 0x8179
	FRAGMENT_LIGHT5_SGIX                                       Enum = 0x8411
	UNSIGNED_INT_24_8_EXT                                      Enum = 0x84FA
	VERTEX_ARRAY_RANGE_NV                                      Enum = 0x851D
	DEPTH24_STENCIL8_OES                                       Enum = 0x88F0
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  Enum = 0x8CD7
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       Enum = 0x8DE1
	INT_10_10_10_2_OES                                         Enum = 0x8DF7
	DYNAMIC_READ_ARB                                           Enum = 0x88E9
	INT8_VEC2_NV                                               Enum = 0x8FE1
	BLEND_SRC_RGB_EXT                                          Enum = 0x80C9
	RGBA_DXT5_S3TC                                             Enum = 0x83A4
	TEXTURE27                                                  Enum = 0x84DB
	FRAGMENT_PROGRAM_BINDING_NV                                Enum = 0x8873
	NUM_PASSES_ATI                                             Enum = 0x8970
	INT_SAMPLER_2D_ARRAY                                       Enum = 0x8DCF
	COPY_WRITE_BUFFER                                          Enum = 0x8F37
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  Enum = 0x8F9F
	SRC1_RGB                                                   Enum = 0x8581
	DRAW_BUFFER6_ATI                                           Enum = 0x882B
	PATH_INITIAL_DASH_CAP_NV                                   Enum = 0x907C
	BOUNDING_BOX_NV                                            Enum = 0x908D
	DEBUG_SEVERITY_MEDIUM_ARB                                  Enum = 0x9147
	COMPRESSED_ALPHA                                           Enum = 0x84E9
	NORMAL_MAP_EXT                                             Enum = 0x8511
	PROXY_TEXTURE_2D_ARRAY                                     Enum = 0x8C1B
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               Enum = 0x8623
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   Enum = 0x88B3
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       Enum = 0x8DA1
	INT_IMAGE_2D_ARRAY                                         Enum = 0x905E
	SGIS_generate_mipmap                                       Enum = 1
	IR_INSTRUMENT1_SGIX                                        Enum = 0x817F
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                Enum = 0x8518
	ADD_ATI                                                    Enum = 0x8963
	FLOAT_VEC3_ARB                                             Enum = 0x8B51
	FRAMEBUFFER_ATTACHMENT_LAYERED                             Enum = 0x8DA7
	INT_SAMPLER_1D                                             Enum = 0x8DC9
	YCBYCR8_422_NV                                             Enum = 0x9031
	RGB8                                                       Enum = 0x8051
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                Enum = 0x8189
	COLOR_SUM_EXT                                              Enum = 0x8458
	MODELVIEW10_ARB                                            Enum = 0x872A
	VERTEX_STREAM2_ATI                                         Enum = 0x876E
	RENDERBUFFER_STENCIL_SIZE_OES                              Enum = 0x8D55
	NUM_COMPATIBLE_SUBROUTINES                                 Enum = 0x8E4A
	MAX_INTEGER_SAMPLES                                        Enum = 0x9110
	PACK_COMPRESSED_BLOCK_DEPTH                                Enum = 0x912D
	EDGE_FLAG_ARRAY_LIST_IBM                                   Enum = 103075
	MULTISAMPLE_EXT                                            Enum = 0x809D
	MODELVIEW1_MATRIX_EXT                                      Enum = 0x8506
	DRAW_BUFFER0_NV                                            Enum = 0x8825
	SLUMINANCE_NV                                              Enum = 0x8C46
	SAMPLE_PATTERN_SGIS                                        Enum = 0x80AC
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           Enum = 0x80B0
	MAP1_VERTEX_ATTRIB13_4_NV                                  Enum = 0x866D
	PROGRAM_BINARY_FORMATS                                     Enum = 0x87FF
	STATIC_COPY_ARB                                            Enum = 0x88E6
	TESSELLATION_FACTOR_AMD                                    Enum = 0x9005
	ALLOW_DRAW_WIN_HINT_PGI                                    Enum = 0x1A20F
	MAX_MODELVIEW_STACK_DEPTH                                  Enum = 0x0D36
	DOUBLE_VEC4_EXT                                            Enum = 0x8FFE
	TEXTURE_BORDER_COLOR_NV                                    Enum = 0x1004
	DEPTH_STENCIL_NV                                           Enum = 0x84F9
	DEPTH_CLAMP                                                Enum = 0x864F
	CON_10_ATI                                                 Enum = 0x894B
	COLOR_ATTACHMENT5_EXT                                      Enum = 0x8CE5
	COLOR_ATTACHMENT10_NV                                      Enum = 0x8CEA
	MULTISAMPLE_BIT                                            Enum = 0x20000000
	COLOR_TABLE_RED_SIZE_SGI                                   Enum = 0x80DA
	TEXTURE_LEQUAL_R_SGIX                                      Enum = 0x819C
	INTERNALFORMAT_PREFERRED                                   Enum = 0x8270
	RGBA8UI_EXT                                                Enum = 0x8D7C
	DRAW_INDIRECT_ADDRESS_NV                                   Enum = 0x8F41
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         Enum = 0x90DD
	HISTOGRAM_SINK                                             Enum = 0x802D
	CONTEXT_FLAGS                                              Enum = 0x821E
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            Enum = 0x8518
	TRACE_ERRORS_BIT_MESA                                      Enum = 0x0020
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            Enum = 0x8E1F
	UNPACK_SKIP_VOLUMES_SGIS                                   Enum = 0x8132
	TRANSPOSE_NV                                               Enum = 0x862C
	UNIFORM_TYPE                                               Enum = 0x8A37
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      Enum = 0x906A
	FLOAT                                                      Enum = 0x1406
	CURRENT_MATRIX_STACK_DEPTH_NV                              Enum = 0x8640
	PIXEL_COUNT_AVAILABLE_NV                                   Enum = 0x8867
	COLOR_ATTACHMENT15                                         Enum = 0x8CEF
	FIRST_TO_REST_NV                                           Enum = 0x90AF
	TRIANGLES_ADJACENCY_ARB                                    Enum = 0x000C
	PROXY_COLOR_TABLE_SGI                                      Enum = 0x80D3
	T2F_IUI_N3F_V3F_EXT                                        Enum = 0x81B4
	TEXTURE_BINDING_BUFFER_EXT                                 Enum = 0x8C2C
	RGB8I_EXT                                                  Enum = 0x8D8F
	POINT_SIZE_MIN_EXT                                         Enum = 0x8126
	R8UI                                                       Enum = 0x8232
	MAP1_VERTEX_ATTRIB8_4_NV                                   Enum = 0x8668
	BUFFER_ACCESS_ARB                                          Enum = 0x88BB
	ACTIVE_VARYINGS_NV                                         Enum = 0x8C81
	MAP2_VERTEX_ATTRIB7_4_NV                                   Enum = 0x8677
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            Enum = 0x887B
	BGR_INTEGER_EXT                                            Enum = 0x8D9A
	LUMINANCE12_ALPHA4                                         Enum = 0x8046
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             Enum = 0x8366
	MAX_3D_TEXTURE_SIZE_EXT                                    Enum = 0x8073
	POST_COLOR_MATRIX_GREEN_BIAS                               Enum = 0x80B9
	PRIMARY_COLOR_NV                                           Enum = 0x852C
	YCBCR_422_APPLE                                            Enum = 0x85B9
	VARIANT_DATATYPE_EXT                                       Enum = 0x87E5
	REFERENCED_BY_TESS_EVALUATION_SHADER                       Enum = 0x9308
	OUTPUT_TEXTURE_COORD25_EXT                                 Enum = 0x87B6
	COLOR_ATTACHMENT9_EXT                                      Enum = 0x8CE9
	SYNC_STATUS_APPLE                                          Enum = 0x9114
	TEXTURE_GEN_Q                                              Enum = 0x0C63
	FRAGMENTS_INSTRUMENT_SGIX                                  Enum = 0x8313
	UNPACK_CLIENT_STORAGE_APPLE                                Enum = 0x85B2
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       Enum = 0x93D2
	TEXTURE_BASE_LEVEL_SGIS                                    Enum = 0x813C
	HALF_BIAS_NEGATE_NV                                        Enum = 0x853B
	OBJECT_BUFFER_USAGE_ATI                                    Enum = 0x8765
	MATRIX19_ARB                                               Enum = 0x88D3
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           Enum = 0x8908
	DT_BIAS_NV                                                 Enum = 0x8717
	COLOR_TABLE_WIDTH                                          Enum = 0x80D9
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            Enum = 0x8519
	RGB_FLOAT16_ATI                                            Enum = 0x881B
	BGRA_INTEGER_EXT                                           Enum = 0x8D9B
	SPHERE_MAP                                                 Enum = 0x2402
	SEPARABLE_2D_EXT                                           Enum = 0x8012
	COLOR_INDEX8_EXT                                           Enum = 0x80E5
	FOG_OFFSET_SGIX                                            Enum = 0x8198
	MAP2_VERTEX_ATTRIB0_4_NV                                   Enum = 0x8670
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           Enum = 0x8895
	PIXEL_PACK_BUFFER_BINDING_EXT                              Enum = 0x88ED
	INT64_VEC2_NV                                              Enum = 0x8FE9
	COMPRESSED_RGBA_ASTC_8x5_KHR                               Enum = 0x93B5
	RGB5_A1                                                    Enum = 0x8057
	CONSTANT_COLOR                                             Enum = 0x8001
	SAMPLE_MASK_INVERT_SGIS                                    Enum = 0x80AB
	TEXTURE_GEQUAL_R_SGIX                                      Enum = 0x819D
	R8I                                                        Enum = 0x8231
	GEOMETRY_INPUT_TYPE_EXT                                    Enum = 0x8DDB
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        Enum = 0x8E47
	READ_BUFFER_NV                                             Enum = 0x0C02
	GL_2PASS_0_SGIS                                            Enum = 0x80A2
	TEXTURE1_ARB                                               Enum = 0x84C1
	PATCH_DEFAULT_OUTER_LEVEL                                  Enum = 0x8E74
	ATOMIC_COUNTER_BUFFER                                      Enum = 0x92C0
	INVALID_ENUM                                               Enum = 0x0500
	LUMINANCE4                                                 Enum = 0x803F
	CLIP_PLANE3                                                Enum = 0x3003
	BLEND_EQUATION_RGB_EXT                                     Enum = 0x8009
	COORD_REPLACE_ARB                                          Enum = 0x8862
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            Enum = 0x88B6
	DEPTH32F_STENCIL8_NV                                       Enum = 0x8DAC
	TEXTURE_HEIGHT                                             Enum = 0x1001
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         Enum = 0x8CD0
	MAT_DIFFUSE_BIT_PGI                                        Enum = 0x00400000
	MATRIX_INDEX_ARRAY_SIZE_OES                                Enum = 0x8846
	BUFFER_FLUSHING_UNMAP_APPLE                                Enum = 0x8A13
	GEOMETRY_PROGRAM_NV                                        Enum = 0x8C26
	RGB32UI                                                    Enum = 0x8D71
	QUAD_STRIP                                                 Enum = 0x0008
	MAX_VERTEX_SHADER_VARIANTS_EXT                             Enum = 0x87C6
	RGBA32F                                                    Enum = 0x8814
	DRAW_BUFFER13_ARB                                          Enum = 0x8832
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              Enum = 0x8A06
	ACTIVE_ATTRIBUTES                                          Enum = 0x8B89
	UNSIGNED_INT_VEC4_EXT                                      Enum = 0x8DC8
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     Enum = 0x90D8
	LIST_INDEX                                                 Enum = 0x0B33
	RGB4                                                       Enum = 0x804F
	SHARED_TEXTURE_PALETTE_EXT                                 Enum = 0x81FB
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             Enum = 0x888F
	DYNAMIC_COPY                                               Enum = 0x88EA
	PROGRAM_OBJECT_ARB                                         Enum = 0x8B40
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      Enum = 0x8E5F
	IMAGE_BINDING_LEVEL_EXT                                    Enum = 0x8F3B
	RGB10                                                      Enum = 0x8052
	TEXTURE_BINDING_CUBE_MAP_ARB                               Enum = 0x8514
	SOURCE2_RGB_EXT                                            Enum = 0x8582
	CON_20_ATI                                                 Enum = 0x8955
	PALETTE4_RGB5_A1_OES                                       Enum = 0x8B94
	TEXTURE_1D_ARRAY_EXT                                       Enum = 0x8C18
	MAX_COMPUTE_IMAGE_UNIFORMS                                 Enum = 0x91BD
	COLOR_ARRAY_SIZE_EXT                                       Enum = 0x8081
	NORMAL_ARRAY_POINTER_EXT                                   Enum = 0x808F
	TEXTURE_MAG_FILTER                                         Enum = 0x2800
	INTENSITY16_EXT                                            Enum = 0x804D
	COLOR_INDEX16_EXT                                          Enum = 0x80E7
	INT_SAMPLER_BUFFER                                         Enum = 0x8DD0
	TESS_GEN_POINT_MODE                                        Enum = 0x8E79
	ALWAYS_SOFT_HINT_PGI                                       Enum = 0x1A20D
	UNSIGNED_SHORT_1_5_5_5_REV                                 Enum = 0x8366
	ARRAY_STRIDE                                               Enum = 0x92FE
	CONVOLUTION_FORMAT                                         Enum = 0x8017
	POINT_FADE_THRESHOLD_SIZE_EXT                              Enum = 0x8128
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             Enum = 0x8312
	RGBA4_S3TC                                                 Enum = 0x83A3
	COMBINER_CD_DOT_PRODUCT_NV                                 Enum = 0x8546
	MATRIX6_ARB                                                Enum = 0x88C6
	SEPARATE_ATTRIBS_EXT                                       Enum = 0x8C8D
	LAST_VERTEX_CONVENTION                                     Enum = 0x8E4E
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        Enum = 0x00000001
	UNSIGNED_INT_8_8_8_8_REV_EXT                               Enum = 0x8367
	COLOR_TABLE_FORMAT                                         Enum = 0x80D8
	PACK_MAX_COMPRESSED_SIZE_SGIX                              Enum = 0x831B
	QUADS                                                      Enum = 0x0007
	MIRROR_CLAMP_TO_EDGE_ATI                                   Enum = 0x8743
	DRAW_BUFFER5_ARB                                           Enum = 0x882A
	UNSIGNED_INT_SAMPLER_BUFFER                                Enum = 0x8DD8
	CMYKA_EXT                                                  Enum = 0x800D
	ATTENUATION_EXT                                            Enum = 0x834D
	Z400_BINARY_AMD                                            Enum = 0x8740
	DRAW_BUFFER7_ARB                                           Enum = 0x882C
	FLOAT_RG_NV                                                Enum = 0x8881
	HALF_FLOAT_OES                                             Enum = 0x8D61
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 Enum = 0x8E54
	POST_COLOR_MATRIX_BLUE_BIAS                                Enum = 0x80BA
	COLOR_ATTACHMENT14_NV                                      Enum = 0x8CEE
	INT_IMAGE_3D                                               Enum = 0x9059
	DEBUG_SEVERITY_HIGH                                        Enum = 0x9146
	NUM_EXTENSIONS                                             Enum = 0x821D
	PROGRAM_ERROR_POSITION_NV                                  Enum = 0x864B
	TRACE_PRIMITIVES_BIT_MESA                                  Enum = 0x0002
	PROGRAM_RESULT_COMPONENTS_NV                               Enum = 0x8907
	CLAMP_FRAGMENT_COLOR_ARB                                   Enum = 0x891B
	TRANSFORM_FEEDBACK_RECORD_NV                               Enum = 0x8C86
	POST_CONVOLUTION_BLUE_BIAS_EXT                             Enum = 0x8022
	RGB12_EXT                                                  Enum = 0x8053
	REPLICATE_BORDER                                           Enum = 0x8153
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          Enum = 0x8213
	EVAL_VERTEX_ATTRIB4_NV                                     Enum = 0x86CA
	DRAW_BUFFER9                                               Enum = 0x882E
	CURRENT_PALETTE_MATRIX_OES                                 Enum = 0x8843
	MAX_TEXTURE_COORDS_NV                                      Enum = 0x8871
	READ_ONLY                                                  Enum = 0x88B8
	UNIFORM_NAME_LENGTH                                        Enum = 0x8A39
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           Enum = 0x8DE0
	INT_IMAGE_2D_EXT                                           Enum = 0x9058
	VARIABLE_C_NV                                              Enum = 0x8525
	CLAMP_READ_COLOR_ARB                                       Enum = 0x891C
	IMAGE_BUFFER_EXT                                           Enum = 0x9051
	DEPTH_COMPONENT32_SGIX                                     Enum = 0x81A7
	PRIMITIVE_RESTART_INDEX_NV                                 Enum = 0x8559
	ATTRIB_ARRAY_SIZE_NV                                       Enum = 0x8623
	STREAM_COPY_ARB                                            Enum = 0x88E2
	MAX_TEXTURE_LOD_BIAS_EXT                                   Enum = 0x84FD
	REG_9_ATI                                                  Enum = 0x892A
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              Enum = 0x9035
	EXT_shared_texture_palette                                 Enum = 1
	ZERO                                                       Enum = 0
	MAX_VIEWPORTS                                              Enum = 0x825B
	STENCIL_RENDERABLE                                         Enum = 0x8288
	RGBA8I_EXT                                                 Enum = 0x8D8E
	FRAMEBUFFER_SRGB                                           Enum = 0x8DB9
	PATH_OBJECT_BOUNDING_BOX_NV                                Enum = 0x908A
	GL_3D                                                      Enum = 0x0601
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               Enum = 0x86D8
	DRAW_BUFFER14_ATI                                          Enum = 0x8833
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           Enum = 0x8C29
	COLOR_ATTACHMENT0                                          Enum = 0x8CE0
	PRIMITIVE_RESTART_INDEX                                    Enum = 0x8F9E
	TRIANGLE_FAN                                               Enum = 0x0006
	PIXEL_TRANSFORM_2D_EXT                                     Enum = 0x8330
	DOT3_RGBA                                                  Enum = 0x86AF
	SGIS_texture_lod                                           Enum = 1
	UNSIGNED_SHORT_4_4_4_4_REV                                 Enum = 0x8365
	TEXTURE21_ARB                                              Enum = 0x84D5
	LUMINANCE_ALPHA_FLOAT32_APPLE                              Enum = 0x8819
	PIXEL_COUNTER_BITS_NV                                      Enum = 0x8864
	GCCSO_SHADER_BINARY_FJ                                     Enum = 0x9260
	FOG_HINT                                                   Enum = 0x0C54
	SYNC_CL_EVENT_COMPLETE_ARB                                 Enum = 0x8241
	MAX_HEIGHT                                                 Enum = 0x827F
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            Enum = 0x86F1
	UNSIGNED_NORMALIZED_ARB                                    Enum = 0x8C17
	RENDERBUFFER_BLUE_SIZE                                     Enum = 0x8D52
	BACK_RIGHT                                                 Enum = 0x0403
	LINEAR_SHARPEN_SGIS                                        Enum = 0x80AD
	COMPRESSED_RED                                             Enum = 0x8225
	STREAM_READ_ARB                                            Enum = 0x88E1
	PIXEL_PACK_BUFFER                                          Enum = 0x88EB
	EIGHTH_BIT_ATI                                             Enum = 0x00000020
	RGB16I_EXT                                                 Enum = 0x8D89
	CULL_FACE                                                  Enum = 0x0B44
	MAP2_TEXTURE_COORD_3                                       Enum = 0x0DB5
	MAX_CONVOLUTION_WIDTH_EXT                                  Enum = 0x801A
	SRGB_WRITE                                                 Enum = 0x8298
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              Enum = 0x83F1
	MATRIX27_ARB                                               Enum = 0x88DB
	SWIZZLE_STR_DR_ATI                                         Enum = 0x8978
	TEXTURE_BINDING_1D_ARRAY_EXT                               Enum = 0x8C1C
	COMPRESSED_RGBA_ASTC_5x4_KHR                               Enum = 0x93B1
	FRAMEBUFFER_UNDEFINED                                      Enum = 0x8219
	TEXTURE_COMPARE_MODE_EXT                                   Enum = 0x884C
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       Enum = 0x8C72
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               Enum = 0x90D1
	EVAL_VERTEX_ATTRIB13_NV                                    Enum = 0x86D3
	TEXTURE_1D_ARRAY                                           Enum = 0x8C18
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        Enum = 0x8C4F
	COLOR_ATTACHMENT9                                          Enum = 0x8CE9
	T                                                          Enum = 0x2001
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          Enum = 0x8DAD
	USE_MISSING_GLYPH_NV                                       Enum = 0x90AA
	ARB_imaging                                                Enum = 1
	LOAD                                                       Enum = 0x0101
	POINT_SMOOTH_HINT                                          Enum = 0x0C51
	MAP_COLOR                                                  Enum = 0x0D10
	CONVOLUTION_1D                                             Enum = 0x8010
	COLOR_ARRAY_COUNT_EXT                                      Enum = 0x8084
	DEBUG_TYPE_OTHER_ARB                                       Enum = 0x8251
	RGB_SCALE                                                  Enum = 0x8573
	MODELVIEW19_ARB                                            Enum = 0x8733
	VERTEX_STREAM5_ATI                                         Enum = 0x8771
	COMPRESSED_SIGNED_RG11_EAC                                 Enum = 0x9273
	VERSION_2_1                                                Enum = 1
	EXT_subtexture                                             Enum = 1
	IMAGE_CLASS_1_X_16                                         Enum = 0x82BE
	COMBINE_ALPHA_ARB                                          Enum = 0x8572
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        Enum = 0x8C4E
	FLOAT16_VEC4_NV                                            Enum = 0x8FFB
	PROXY_TEXTURE_2D_STACK_MESAX                               Enum = 0x875C
	EXT_cmyka                                                  Enum = 1
	TEXTURE_WRAP_R_OES                                         Enum = 0x8072
	R1UI_C4UB_V3F_SUN                                          Enum = 0x85C5
	OP_DOT3_EXT                                                Enum = 0x8784
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          Enum = 0x8DD6
	COMPRESSED_RGBA_ASTC_12x10_KHR                             Enum = 0x93BC
	STRICT_LIGHTING_HINT_PGI                                   Enum = 0x1A217
	SGIX_async                                                 Enum = 1
	DUAL_LUMINANCE16_SGIS                                      Enum = 0x8117
	UNIFORM_BUFFER_BINDING                                     Enum = 0x8A28
	INDEX_CLEAR_VALUE                                          Enum = 0x0C20
	FUNC_ADD_EXT                                               Enum = 0x8006
	CONVOLUTION_FORMAT_EXT                                     Enum = 0x8017
	CURRENT_SECONDARY_COLOR_EXT                                Enum = 0x8459
	MAP2_VERTEX_ATTRIB11_4_NV                                  Enum = 0x867B
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             Enum = 0x8A08
	ACTIVE_UNIFORM_MAX_LENGTH                                  Enum = 0x8B87
	TEXTURE_COORD_ARRAY_SIZE                                   Enum = 0x8088
	VERTEX_ARRAY_COUNT_EXT                                     Enum = 0x807D
	TEXTURE_CUBE_MAP_OES                                       Enum = 0x8513
	TEXTURE_WIDTH_QCOM                                         Enum = 0x8BD2
	TEXTURE_OBJECT_VALID_QCOM                                  Enum = 0x8BDB
	DEPTH_STENCIL_TO_RGBA_NV                                   Enum = 0x886E
	LINE_TO_NV                                                 Enum = 0x04
	EXT_abgr                                                   Enum = 1
	POINT_FADE_THRESHOLD_SIZE_SGIS                             Enum = 0x8128
	GREEN_MAX_CLAMP_INGR                                       Enum = 0x8565
	SRC1_ALPHA                                                 Enum = 0x8589
	VARIANT_ARRAY_STRIDE_EXT                                   Enum = 0x87E6
	COLOR_ARRAY_BUFFER_BINDING_ARB                             Enum = 0x8898
	MAX_PROGRAM_TEXEL_OFFSET                                   Enum = 0x8905
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   Enum = 0x8DE1
	SHADER_COMPILER                                            Enum = 0x8DFA
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         Enum = 0x8E83
	MALI_PROGRAM_BINARY_ARM                                    Enum = 0x8F61
	COMPUTE_PROGRAM_NV                                         Enum = 0x90FB
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             Enum = 0x92D8
	DEBUG_SOURCE_APPLICATION                                   Enum = 0x824A
	OP_CROSS_PRODUCT_EXT                                       Enum = 0x8797
	PROGRAM_TEX_INDIRECTIONS_ARB                               Enum = 0x8807
	REG_30_ATI                                                 Enum = 0x893F
	AUTO_GENERATE_MIPMAP                                       Enum = 0x8295
	MAX_RENDERBUFFER_SIZE_EXT                                  Enum = 0x84E8
	STENCIL_CLEAR_VALUE                                        Enum = 0x0B91
	VERTEX_BINDING_DIVISOR                                     Enum = 0x82D6
	UNIFORM_BUFFER_SIZE                                        Enum = 0x8A2A
	UNSIGNED_INT_10F_11F_11F_REV                               Enum = 0x8C3B
	COLOR_ATTACHMENT6                                          Enum = 0x8CE6
	IMAGE_2D_EXT                                               Enum = 0x904D
	TIMEOUT_IGNORED                                            Enum = 0xFFFFFFFF
	GL_2PASS_0_EXT                                             Enum = 0x80A2
	MAX_TEXTURE_IMAGE_UNITS                                    Enum = 0x8872
	MAX_PROGRAM_GENERIC_RESULTS_NV                             Enum = 0x8DA6
	COMMAND_BARRIER_BIT                                        Enum = 0x00000040
	VERTEX_WEIGHTING_EXT                                       Enum = 0x8509
	CONSTANT_COLOR0_NV                                         Enum = 0x852A
	FLOAT_MAT4_ARB                                             Enum = 0x8B5C
	QUERY_OBJECT_AMD                                           Enum = 0x9153
	PACK_REVERSE_ROW_ORDER_ANGLE                               Enum = 0x93A4
	CONTEXT_FLAG_DEBUG_BIT                                     Enum = 0x00000002
	RGBA                                                       Enum = 0x1908
	HISTOGRAM                                                  Enum = 0x8024
	IMAGE_TRANSFORM_2D_HP                                      Enum = 0x8161
	VIEW_CLASS_8_BITS                                          Enum = 0x82CB
	OP_FLOOR_EXT                                               Enum = 0x878F
	BUFFER_MAP_POINTER                                         Enum = 0x88BD
	UNIFORM_BLOCK_BINDING                                      Enum = 0x8A3F
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             Enum = 0x8E70
	DUAL_ALPHA16_SGIS                                          Enum = 0x8113
	MAP1_VERTEX_ATTRIB14_4_NV                                  Enum = 0x866E
	MAX_PALETTE_MATRICES_ARB                                   Enum = 0x8842
	CON_19_ATI                                                 Enum = 0x8954
	POINT_SIZE_ARRAY_POINTER_OES                               Enum = 0x898C
	QUERY_WAIT_NV                                              Enum = 0x8E13
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             Enum = 0x900A
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            Enum = 0x80BA
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         Enum = 0x80D4
	ALPHA16I_EXT                                               Enum = 0x8D8A
	RELATIVE_LINE_TO_NV                                        Enum = 0x05
	SCISSOR_BIT                                                Enum = 0x00080000
	VERTEX_ARRAY_STRIDE                                        Enum = 0x807C
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           Enum = 0x8187
	LIGHT_MODEL_COLOR_CONTROL_EXT                              Enum = 0x81F8
	CURRENT_FOG_COORDINATE                                     Enum = 0x8453
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               Enum = 0x9143
	REFERENCED_BY_TESS_CONTROL_SHADER                          Enum = 0x9307
	UNSIGNED_BYTE_2_3_3_REV                                    Enum = 0x8362
	R1UI_N3F_V3F_SUN                                           Enum = 0x85C7
	TESS_CONTROL_PROGRAM_NV                                    Enum = 0x891E
	ALPHA_INTEGER                                              Enum = 0x8D97
	TRANSLATE_Y_NV                                             Enum = 0x908F
	LINE                                                       Enum = 0x1B01
	FOG_DISTANCE_MODE_NV                                       Enum = 0x855A
	EVAL_FRACTIONAL_TESSELLATION_NV                            Enum = 0x86C5
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          Enum = 0x8DD7
	RGBA16_SNORM                                               Enum = 0x8F9B
	COUNT_DOWN_NV                                              Enum = 0x9089
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            Enum = 0x90CB
	MAT_SHININESS_BIT_PGI                                      Enum = 0x02000000
	TEXTURE_MAX_CLAMP_R_SGIX                                   Enum = 0x836B
	TEXTURE11                                                  Enum = 0x84CB
	TEXTURE_MAX_ANISOTROPY_EXT                                 Enum = 0x84FE
	INTENSITY8_SNORM                                           Enum = 0x9017
	LINE_WIDTH_RANGE                                           Enum = 0x0B22
	TEXTURE9                                                   Enum = 0x84C9
	TEXTURE_LOD_BIAS_EXT                                       Enum = 0x8501
	OUTPUT_TEXTURE_COORD23_EXT                                 Enum = 0x87B4
	SGIX_convolution_accuracy                                  Enum = 1
	IMAGE_PIXEL_FORMAT                                         Enum = 0x82A9
	TEXTURE_RANGE_LENGTH_APPLE                                 Enum = 0x85B7
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           Enum = 0x864E
	OFFSET_TEXTURE_2D_BIAS_NV                                  Enum = 0x86E3
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         Enum = 0x8CDA
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   Enum = 0x8DA8
	IMAGE_2D_RECT_EXT                                          Enum = 0x904F
	CLIENT_PIXEL_STORE_BIT                                     Enum = 0x00000001
	R16F                                                       Enum = 0x822D
	VIEW_CLASS_96_BITS                                         Enum = 0x82C5
	PIXEL_SUBSAMPLE_4242_SGIX                                  Enum = 0x85A4
	OFFSET_TEXTURE_MATRIX_NV                                   Enum = 0x86E1
	CPU_OPTIMIZED_QCOM                                         Enum = 0x8FB1
	MODELVIEW_MATRIX                                           Enum = 0x0BA6
	SPOT_EXPONENT                                              Enum = 0x1205
	MINMAX                                                     Enum = 0x802E
	EXT_texture_object                                         Enum = 1
	PROGRAM_SEPARABLE_EXT                                      Enum = 0x8258
	SYNC_CONDITION_APPLE                                       Enum = 0x9113
	SYNC_CL_EVENT_ARB                                          Enum = 0x8240
	DECR_WRAP_EXT                                              Enum = 0x8508
	DRAW_BUFFER14_ARB                                          Enum = 0x8833
	RG_SNORM                                                   Enum = 0x8F91
	MAX_GEOMETRY_INPUT_COMPONENTS                              Enum = 0x9123
	TRIANGLE_STRIP_ADJACENCY_ARB                               Enum = 0x000D
	SAMPLE_BUFFERS_SGIS                                        Enum = 0x80A8
	POINT_FADE_THRESHOLD_SIZE_ARB                              Enum = 0x8128
	DEPTH_RENDERABLE                                           Enum = 0x8287
	ALIASED_POINT_SIZE_RANGE                                   Enum = 0x846D
	MATRIX7_ARB                                                Enum = 0x88C7
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           Enum = 0x8DE2
	NEAREST_MIPMAP_NEAREST                                     Enum = 0x2700
	INDEX_ARRAY_TYPE_EXT                                       Enum = 0x8085
	DOUBLE_MAT2x3                                              Enum = 0x8F49
	HISTOGRAM_LUMINANCE_SIZE_EXT                               Enum = 0x802C
	GEOMETRY_TEXTURE                                           Enum = 0x829E
	TEXTURE_COMPRESSED                                         Enum = 0x86A1
	VARIABLE_G_NV                                              Enum = 0x8529
	CURRENT_BIT                                                Enum = 0x00000001
	TEXTURE_4D_SGIS                                            Enum = 0x8134
	RESTART_SUN                                                Enum = 0x0001
	PERTURB_EXT                                                Enum = 0x85AE
	TEXTURE_2D_ARRAY                                           Enum = 0x8C1A
	DOUBLE_MAT2x3_EXT                                          Enum = 0x8F49
	CURRENT_INDEX                                              Enum = 0x0B01
	TEXTURE_COMPARE_OPERATOR_SGIX                              Enum = 0x819B
	MAX_COMPUTE_UNIFORM_COMPONENTS                             Enum = 0x8263
	COVERAGE_COMPONENT4_NV                                     Enum = 0x8ED1
	PATH_FORMAT_PS_NV                                          Enum = 0x9071
	MATRIX1_NV                                                 Enum = 0x8631
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      Enum = 0x8CD6
	UNPACK_IMAGE_HEIGHT                                        Enum = 0x806E
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         Enum = 0x8336
	OUTPUT_TEXTURE_COORD2_EXT                                  Enum = 0x879F
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    Enum = 0x889D
	NUM_LOOPBACK_COMPONENTS_ATI                                Enum = 0x8974
	PACK_COMPRESSED_BLOCK_HEIGHT                               Enum = 0x912C
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      Enum = 0x93D9
	BLEND_SRC_RGB_OES                                          Enum = 0x80C9
	DOUBLE_MAT4x2                                              Enum = 0x8F4D
	YCBAYCR8A_4224_NV                                          Enum = 0x9032
)

var enumLookup = make(map[Enum]string, 4834)

func init() {
	enumLookup[WEIGHT_SUM_UNITY_ARB] = "WEIGHT_SUM_UNITY_ARB"
	enumLookup[PATH_GEN_MODE_NV] = "PATH_GEN_MODE_NV"
	enumLookup[LIGHT4] = "LIGHT4"
	enumLookup[MAD_ATI] = "MAD_ATI"
	enumLookup[FRAGMENT_SHADER_ARB] = "FRAGMENT_SHADER_ARB"
	enumLookup[RENDER_DIRECT_TO_FRAMEBUFFER_QCOM] = "RENDER_DIRECT_TO_FRAMEBUFFER_QCOM"
	enumLookup[TEXTURE15] = "TEXTURE15"
	enumLookup[RENDERBUFFER_DEPTH_SIZE_OES] = "RENDERBUFFER_DEPTH_SIZE_OES"
	enumLookup[MAP1_TEXTURE_COORD_1] = "MAP1_TEXTURE_COORD_1"
	enumLookup[RED_EXT] = "RED_EXT"
	enumLookup[SPRITE_SGIX] = "SPRITE_SGIX"
	enumLookup[COMPUTE_TEXTURE] = "COMPUTE_TEXTURE"
	enumLookup[TEXTURE_WRAP_R_EXT] = "TEXTURE_WRAP_R_EXT"
	enumLookup[FOG_COORD_ARRAY_POINTER] = "FOG_COORD_ARRAY_POINTER"
	enumLookup[FENCE_STATUS_NV] = "FENCE_STATUS_NV"
	enumLookup[DEPTH_STENCIL_OES] = "DEPTH_STENCIL_OES"
	enumLookup[CON_6_ATI] = "CON_6_ATI"
	enumLookup[SAMPLER_2D_RECT] = "SAMPLER_2D_RECT"
	enumLookup[IMAGE_1D_ARRAY] = "IMAGE_1D_ARRAY"
	enumLookup[DEBUG_OUTPUT_SYNCHRONOUS_ARB] = "DEBUG_OUTPUT_SYNCHRONOUS_ARB"
	enumLookup[TEXTURE_COMPRESSION_HINT_ARB] = "TEXTURE_COMPRESSION_HINT_ARB"
	enumLookup[EVAL_VERTEX_ATTRIB10_NV] = "EVAL_VERTEX_ATTRIB10_NV"
	enumLookup[FLOAT_RGB_NV] = "FLOAT_RGB_NV"
	enumLookup[MAX_PROGRAM_LOOP_DEPTH_NV] = "MAX_PROGRAM_LOOP_DEPTH_NV"
	enumLookup[SWIZZLE_STRQ_ATI] = "SWIZZLE_STRQ_ATI"
	enumLookup[MAX_COLOR_ATTACHMENTS] = "MAX_COLOR_ATTACHMENTS"
	enumLookup[SGIX_icc_texture] = "SGIX_icc_texture"
	enumLookup[STENCIL_BUFFER_BIT] = "STENCIL_BUFFER_BIT"
	enumLookup[FOG_INDEX] = "FOG_INDEX"
	enumLookup[REFERENCE_PLANE_EQUATION_SGIX] = "REFERENCE_PLANE_EQUATION_SGIX"
	enumLookup[R32I] = "R32I"
	enumLookup[IMAGE_CLASS_2_X_16] = "IMAGE_CLASS_2_X_16"
	enumLookup[EYE_RADIAL_NV] = "EYE_RADIAL_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY11_NV] = "VERTEX_ATTRIB_ARRAY11_NV"
	enumLookup[PIXEL_UNPACK_BUFFER_BINDING_EXT] = "PIXEL_UNPACK_BUFFER_BINDING_EXT"
	enumLookup[IMAGE_FORMAT_COMPATIBILITY_BY_SIZE] = "IMAGE_FORMAT_COMPATIBILITY_BY_SIZE"
	enumLookup[LUMINANCE12] = "LUMINANCE12"
	enumLookup[RESET_NOTIFICATION_STRATEGY_ARB] = "RESET_NOTIFICATION_STRATEGY_ARB"
	enumLookup[CURRENT_RASTER_COLOR] = "CURRENT_RASTER_COLOR"
	enumLookup[DEPTH_BITS] = "DEPTH_BITS"
	enumLookup[PIXEL_TEX_GEN_Q_CEILING_SGIX] = "PIXEL_TEX_GEN_Q_CEILING_SGIX"
	enumLookup[TRIANGLE_MESH_SUN] = "TRIANGLE_MESH_SUN"
	enumLookup[BUFFER_MAP_OFFSET] = "BUFFER_MAP_OFFSET"
	enumLookup[UNSIGNED_BYTE_3_3_2] = "UNSIGNED_BYTE_3_3_2"
	enumLookup[MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX] = "MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX"
	enumLookup[VIEW_CLASS_RGTC1_RED] = "VIEW_CLASS_RGTC1_RED"
	enumLookup[RENDERBUFFER_BINDING_OES] = "RENDERBUFFER_BINDING_OES"
	enumLookup[RENDERBUFFER] = "RENDERBUFFER"
	enumLookup[FRAMEBUFFER_ATTACHMENT_LAYERED_EXT] = "FRAMEBUFFER_ATTACHMENT_LAYERED_EXT"
	enumLookup[MODELVIEW29_ARB] = "MODELVIEW29_ARB"
	enumLookup[TEXTURE_RED_TYPE] = "TEXTURE_RED_TYPE"
	enumLookup[DEPTH_CLAMP_NEAR_AMD] = "DEPTH_CLAMP_NEAR_AMD"
	enumLookup[EXT_blend_color] = "EXT_blend_color"
	enumLookup[SAMPLER_2D] = "SAMPLER_2D"
	enumLookup[READ_FRAMEBUFFER] = "READ_FRAMEBUFFER"
	enumLookup[READ_FRAMEBUFFER_BINDING] = "READ_FRAMEBUFFER_BINDING"
	enumLookup[NUM_VIDEO_CAPTURE_STREAMS_NV] = "NUM_VIDEO_CAPTURE_STREAMS_NV"
	enumLookup[TEXTURE_2D_MULTISAMPLE] = "TEXTURE_2D_MULTISAMPLE"
	enumLookup[TEXTURE_1D] = "TEXTURE_1D"
	enumLookup[N3F_V3F] = "N3F_V3F"
	enumLookup[DRAW_BUFFER11_ATI] = "DRAW_BUFFER11_ATI"
	enumLookup[BYTE] = "BYTE"
	enumLookup[VERTEX_ATTRIB_ARRAY10_NV] = "VERTEX_ATTRIB_ARRAY10_NV"
	enumLookup[ACTIVE_SUBROUTINE_MAX_LENGTH] = "ACTIVE_SUBROUTINE_MAX_LENGTH"
	enumLookup[PATH_DASH_CAPS_NV] = "PATH_DASH_CAPS_NV"
	enumLookup[CURRENT_RASTER_TEXTURE_COORDS] = "CURRENT_RASTER_TEXTURE_COORDS"
	enumLookup[TEXTURE_LUMINANCE_SIZE] = "TEXTURE_LUMINANCE_SIZE"
	enumLookup[OUTPUT_TEXTURE_COORD19_EXT] = "OUTPUT_TEXTURE_COORD19_EXT"
	enumLookup[TEXTURE_BUFFER_FORMAT_ARB] = "TEXTURE_BUFFER_FORMAT_ARB"
	enumLookup[INT_IMAGE_CUBE_MAP_ARRAY] = "INT_IMAGE_CUBE_MAP_ARRAY"
	enumLookup[TEXTURE_1D_BINDING_EXT] = "TEXTURE_1D_BINDING_EXT"
	enumLookup[CONST_EYE_NV] = "CONST_EYE_NV"
	enumLookup[MAX_SPARSE_3D_TEXTURE_SIZE_AMD] = "MAX_SPARSE_3D_TEXTURE_SIZE_AMD"
	enumLookup[COMPRESSED_RGBA_ASTC_8x8_KHR] = "COMPRESSED_RGBA_ASTC_8x8_KHR"
	enumLookup[EQUAL] = "EQUAL"
	enumLookup[COLOR_ARRAY] = "COLOR_ARRAY"
	enumLookup[MODELVIEW0_EXT] = "MODELVIEW0_EXT"
	enumLookup[WEIGHT_ARRAY_BUFFER_BINDING_OES] = "WEIGHT_ARRAY_BUFFER_BINDING_OES"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT"
	enumLookup[FRAGMENT_SHADER_BIT_EXT] = "FRAGMENT_SHADER_BIT_EXT"
	enumLookup[TABLE_TOO_LARGE] = "TABLE_TOO_LARGE"
	enumLookup[SAMPLE_MASK_INVERT_EXT] = "SAMPLE_MASK_INVERT_EXT"
	enumLookup[MODELVIEW7_ARB] = "MODELVIEW7_ARB"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL"
	enumLookup[SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV] = "SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV"
	enumLookup[DEBUG_SOURCE_OTHER] = "DEBUG_SOURCE_OTHER"
	enumLookup[READ_PIXELS_FORMAT] = "READ_PIXELS_FORMAT"
	enumLookup[VIEW_CLASS_64_BITS] = "VIEW_CLASS_64_BITS"
	enumLookup[UNSIGNED_INT_SAMPLER_CUBE_EXT] = "UNSIGNED_INT_SAMPLER_CUBE_EXT"
	enumLookup[GPU_OPTIMIZED_QCOM] = "GPU_OPTIMIZED_QCOM"
	enumLookup[BLEND_DST_ALPHA_EXT] = "BLEND_DST_ALPHA_EXT"
	enumLookup[DEPTH24_STENCIL8] = "DEPTH24_STENCIL8"
	enumLookup[MUL_ATI] = "MUL_ATI"
	enumLookup[COVERAGE_BUFFERS_NV] = "COVERAGE_BUFFERS_NV"
	enumLookup[INT_IMAGE_2D_MULTISAMPLE] = "INT_IMAGE_2D_MULTISAMPLE"
	enumLookup[INDEX_ARRAY_TYPE] = "INDEX_ARRAY_TYPE"
	enumLookup[GET_TEXTURE_IMAGE_FORMAT] = "GET_TEXTURE_IMAGE_FORMAT"
	enumLookup[MAP2_BINORMAL_EXT] = "MAP2_BINORMAL_EXT"
	enumLookup[VERTEX_ATTRIB_MAP1_APPLE] = "VERTEX_ATTRIB_MAP1_APPLE"
	enumLookup[FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB] = "FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB"
	enumLookup[GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV] = "GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV"
	enumLookup[FOG_START] = "FOG_START"
	enumLookup[FOG_MODE] = "FOG_MODE"
	enumLookup[ACCUM_RED_BITS] = "ACCUM_RED_BITS"
	enumLookup[NICEST] = "NICEST"
	enumLookup[MAX_RATIONAL_EVAL_ORDER_NV] = "MAX_RATIONAL_EVAL_ORDER_NV"
	enumLookup[COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT] = "COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT"
	enumLookup[ATOMIC_COUNTER_BUFFER_BINDING] = "ATOMIC_COUNTER_BUFFER_BINDING"
	enumLookup[LINE_SMOOTH] = "LINE_SMOOTH"
	enumLookup[RENDER] = "RENDER"
	enumLookup[UNSIGNED_INT_10_10_10_2_EXT] = "UNSIGNED_INT_10_10_10_2_EXT"
	enumLookup[COLOR_MATRIX] = "COLOR_MATRIX"
	enumLookup[COMPRESSED_SLUMINANCE_ALPHA_EXT] = "COMPRESSED_SLUMINANCE_ALPHA_EXT"
	enumLookup[RENDERBUFFER_SAMPLES_IMG] = "RENDERBUFFER_SAMPLES_IMG"
	enumLookup[SGIX_impact_pixel_texture] = "SGIX_impact_pixel_texture"
	enumLookup[COLOR_ARRAY_PARALLEL_POINTERS_INTEL] = "COLOR_ARRAY_PARALLEL_POINTERS_INTEL"
	enumLookup[FRAMEBUFFER_ATTACHMENT_LAYERED_ARB] = "FRAMEBUFFER_ATTACHMENT_LAYERED_ARB"
	enumLookup[GL_1PASS_EXT] = "GL_1PASS_EXT"
	enumLookup[MAX_ELEMENTS_VERTICES] = "MAX_ELEMENTS_VERTICES"
	enumLookup[CURRENT_MATRIX_NV] = "CURRENT_MATRIX_NV"
	enumLookup[BUFFER_MAPPED_OES] = "BUFFER_MAPPED_OES"
	enumLookup[SYNC_OBJECT_APPLE] = "SYNC_OBJECT_APPLE"
	enumLookup[MAX_TEXTURE_BUFFER_SIZE_ARB] = "MAX_TEXTURE_BUFFER_SIZE_ARB"
	enumLookup[SRGB8_EXT] = "SRGB8_EXT"
	enumLookup[OPERAND3_ALPHA_NV] = "OPERAND3_ALPHA_NV"
	enumLookup[STENCIL_BACK_PASS_DEPTH_PASS_ATI] = "STENCIL_BACK_PASS_DEPTH_PASS_ATI"
	enumLookup[INVALID_INDEX] = "INVALID_INDEX"
	enumLookup[SAMPLER_CUBE_SHADOW_NV] = "SAMPLER_CUBE_SHADOW_NV"
	enumLookup[VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV] = "VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV"
	enumLookup[PROGRAM_BINARY_LENGTH] = "PROGRAM_BINARY_LENGTH"
	enumLookup[DRAW_BUFFER9_ATI] = "DRAW_BUFFER9_ATI"
	enumLookup[TESS_EVALUATION_PROGRAM_NV] = "TESS_EVALUATION_PROGRAM_NV"
	enumLookup[DRAW_FRAMEBUFFER_BINDING] = "DRAW_FRAMEBUFFER_BINDING"
	enumLookup[TESS_CONTROL_SHADER] = "TESS_CONTROL_SHADER"
	enumLookup[DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV] = "DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV"
	enumLookup[RGBA_INTEGER] = "RGBA_INTEGER"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR"
	enumLookup[VIEW_CLASS_32_BITS] = "VIEW_CLASS_32_BITS"
	enumLookup[FOG_COORDINATE_ARRAY] = "FOG_COORDINATE_ARRAY"
	enumLookup[PN_TRIANGLES_POINT_MODE_ATI] = "PN_TRIANGLES_POINT_MODE_ATI"
	enumLookup[SAMPLER_CUBE_ARB] = "SAMPLER_CUBE_ARB"
	enumLookup[ELEMENT_ARRAY_BUFFER] = "ELEMENT_ARRAY_BUFFER"
	enumLookup[T2F_C3F_V3F] = "T2F_C3F_V3F"
	enumLookup[REDUCE] = "REDUCE"
	enumLookup[PROGRAM_BINARY_LENGTH_OES] = "PROGRAM_BINARY_LENGTH_OES"
	enumLookup[DRAW_BUFFER8_ARB] = "DRAW_BUFFER8_ARB"
	enumLookup[SAMPLER_2D_SHADOW_ARB] = "SAMPLER_2D_SHADOW_ARB"
	enumLookup[ACTIVE_VARYING_MAX_LENGTH_NV] = "ACTIVE_VARYING_MAX_LENGTH_NV"
	enumLookup[SHADER_STORAGE_BARRIER_BIT] = "SHADER_STORAGE_BARRIER_BIT"
	enumLookup[NEAREST_CLIPMAP_NEAREST_SGIX] = "NEAREST_CLIPMAP_NEAREST_SGIX"
	enumLookup[REG_16_ATI] = "REG_16_ATI"
	enumLookup[POINT_SPRITE_COORD_ORIGIN] = "POINT_SPRITE_COORD_ORIGIN"
	enumLookup[PROGRAM_MATRIX_EXT] = "PROGRAM_MATRIX_EXT"
	enumLookup[TEXTURE_DEPTH_EXT] = "TEXTURE_DEPTH_EXT"
	enumLookup[INTERNALFORMAT_BLUE_SIZE] = "INTERNALFORMAT_BLUE_SIZE"
	enumLookup[SHADER] = "SHADER"
	enumLookup[SOURCE0_RGB] = "SOURCE0_RGB"
	enumLookup[SURFACE_STATE_NV] = "SURFACE_STATE_NV"
	enumLookup[ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH] = "ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH"
	enumLookup[FRAMEBUFFER_DEFAULT_SAMPLES] = "FRAMEBUFFER_DEFAULT_SAMPLES"
	enumLookup[LIGHTING_BIT] = "LIGHTING_BIT"
	enumLookup[PROXY_HISTOGRAM_EXT] = "PROXY_HISTOGRAM_EXT"
	enumLookup[FRAMEBUFFER_UNDEFINED_OES] = "FRAMEBUFFER_UNDEFINED_OES"
	enumLookup[FRAGMENTS_INSTRUMENT_COUNTERS_SGIX] = "FRAGMENTS_INSTRUMENT_COUNTERS_SGIX"
	enumLookup[NORMAL_MAP] = "NORMAL_MAP"
	enumLookup[COMBINE_ARB] = "COMBINE_ARB"
	enumLookup[OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB] = "OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB"
	enumLookup[SRC_ALPHA_SATURATE] = "SRC_ALPHA_SATURATE"
	enumLookup[SAMPLER] = "SAMPLER"
	enumLookup[MAX_CUBE_MAP_TEXTURE_SIZE] = "MAX_CUBE_MAP_TEXTURE_SIZE"
	enumLookup[MODELVIEW23_ARB] = "MODELVIEW23_ARB"
	enumLookup[UNSIGNED_INT_8_24_REV_MESA] = "UNSIGNED_INT_8_24_REV_MESA"
	enumLookup[DOUBLE_MAT3x4_EXT] = "DOUBLE_MAT3x4_EXT"
	enumLookup[FOG_COLOR] = "FOG_COLOR"
	enumLookup[QUAD_TEXTURE_SELECT_SGIS] = "QUAD_TEXTURE_SELECT_SGIS"
	enumLookup[RG8I] = "RG8I"
	enumLookup[SGX_BINARY_IMG] = "SGX_BINARY_IMG"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT] = "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT"
	enumLookup[TRANSFORM_FEEDBACK_BINDING] = "TRANSFORM_FEEDBACK_BINDING"
	enumLookup[DEBUG_SEVERITY_LOW] = "DEBUG_SEVERITY_LOW"
	enumLookup[POLYGON_BIT] = "POLYGON_BIT"
	enumLookup[PACK_SKIP_IMAGES_EXT] = "PACK_SKIP_IMAGES_EXT"
	enumLookup[OCCLUSION_TEST_RESULT_HP] = "OCCLUSION_TEST_RESULT_HP"
	enumLookup[DEFORMATIONS_MASK_SGIX] = "DEFORMATIONS_MASK_SGIX"
	enumLookup[SIGNED_RGB_NV] = "SIGNED_RGB_NV"
	enumLookup[LINE_QUALITY_HINT_SGIX] = "LINE_QUALITY_HINT_SGIX"
	enumLookup[TESS_GEN_SPACING] = "TESS_GEN_SPACING"
	enumLookup[IMAGE_TRANSLATE_Y_HP] = "IMAGE_TRANSLATE_Y_HP"
	enumLookup[COMPRESSED_RGBA_S3TC_DXT3_EXT] = "COMPRESSED_RGBA_S3TC_DXT3_EXT"
	enumLookup[IMAGE_2D_RECT] = "IMAGE_2D_RECT"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ETC2_EAC] = "COMPRESSED_SRGB8_ALPHA8_ETC2_EAC"
	enumLookup[TRANSFORM_FEEDBACK_BARRIER_BIT] = "TRANSFORM_FEEDBACK_BARRIER_BIT"
	enumLookup[TEXTURE_BINDING_RECTANGLE_ARB] = "TEXTURE_BINDING_RECTANGLE_ARB"
	enumLookup[STENCIL_OP_VALUE_AMD] = "STENCIL_OP_VALUE_AMD"
	enumLookup[DRAW_BUFFER8] = "DRAW_BUFFER8"
	enumLookup[TEXTURE_VIEW] = "TEXTURE_VIEW"
	enumLookup[RENDERBUFFER_GREEN_SIZE_OES] = "RENDERBUFFER_GREEN_SIZE_OES"
	enumLookup[MODELVIEW3_ARB] = "MODELVIEW3_ARB"
	enumLookup[DEPTH_COMPONENT32F] = "DEPTH_COMPONENT32F"
	enumLookup[TRIANGULAR_NV] = "TRIANGULAR_NV"
	enumLookup[ONE_MINUS_CONSTANT_ALPHA] = "ONE_MINUS_CONSTANT_ALPHA"
	enumLookup[SIGNED_LUMINANCE8_NV] = "SIGNED_LUMINANCE8_NV"
	enumLookup[QUERY_RESULT_ARB] = "QUERY_RESULT_ARB"
	enumLookup[MAX_PROGRAM_EXEC_INSTRUCTIONS_NV] = "MAX_PROGRAM_EXEC_INSTRUCTIONS_NV"
	enumLookup[COMPRESSED_RGBA_ARB] = "COMPRESSED_RGBA_ARB"
	enumLookup[COMBINER4_NV] = "COMBINER4_NV"
	enumLookup[CURRENT_VERTEX_EXT] = "CURRENT_VERTEX_EXT"
	enumLookup[FRAMEZOOM_SGIX] = "FRAMEZOOM_SGIX"
	enumLookup[MAX_COMBINED_DIMENSIONS] = "MAX_COMBINED_DIMENSIONS"
	enumLookup[DRAW_BUFFER15_NV] = "DRAW_BUFFER15_NV"
	enumLookup[DOUBLE_VEC4] = "DOUBLE_VEC4"
	enumLookup[TEXTURE23] = "TEXTURE23"
	enumLookup[PROGRAM_POINT_SIZE] = "PROGRAM_POINT_SIZE"
	enumLookup[MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB] = "MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB"
	enumLookup[RENDERBUFFER_HEIGHT] = "RENDERBUFFER_HEIGHT"
	enumLookup[MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT] = "MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT"
	enumLookup[TRANSLATE_X_NV] = "TRANSLATE_X_NV"
	enumLookup[FRAMEBUFFER_BARRIER_BIT_EXT] = "FRAMEBUFFER_BARRIER_BIT_EXT"
	enumLookup[MAJOR_VERSION] = "MAJOR_VERSION"
	enumLookup[IMAGE_CLASS_4_X_16] = "IMAGE_CLASS_4_X_16"
	enumLookup[TEXTURE_DS_SIZE_NV] = "TEXTURE_DS_SIZE_NV"
	enumLookup[FRAMEBUFFER_BINDING_ANGLE] = "FRAMEBUFFER_BINDING_ANGLE"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_EXT] = "UNSIGNED_INT_SAMPLER_2D_EXT"
	enumLookup[CLIENT_VERTEX_ARRAY_BIT] = "CLIENT_VERTEX_ARRAY_BIT"
	enumLookup[TEXTURE_LUMINANCE_SIZE_EXT] = "TEXTURE_LUMINANCE_SIZE_EXT"
	enumLookup[R16I] = "R16I"
	enumLookup[VERTEX_PROGRAM_TWO_SIDE_ARB] = "VERTEX_PROGRAM_TWO_SIDE_ARB"
	enumLookup[PIXEL_UNPACK_BUFFER] = "PIXEL_UNPACK_BUFFER"
	enumLookup[MAX_PROGRAM_CALL_DEPTH_NV] = "MAX_PROGRAM_CALL_DEPTH_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_ADDRESS_NV] = "VERTEX_ATTRIB_ARRAY_ADDRESS_NV"
	enumLookup[ALPHA4] = "ALPHA4"
	enumLookup[TEXTURE0] = "TEXTURE0"
	enumLookup[TEXTURE19] = "TEXTURE19"
	enumLookup[TRANSPOSE_TEXTURE_MATRIX] = "TRANSPOSE_TEXTURE_MATRIX"
	enumLookup[REPLACE_VALUE_AMD] = "REPLACE_VALUE_AMD"
	enumLookup[DRAW_BUFFER11_ARB] = "DRAW_BUFFER11_ARB"
	enumLookup[CURRENT_QUERY_ARB] = "CURRENT_QUERY_ARB"
	enumLookup[TEXTURE_BUFFER] = "TEXTURE_BUFFER"
	enumLookup[MAX_COLOR_ATTACHMENTS_NV] = "MAX_COLOR_ATTACHMENTS_NV"
	enumLookup[MAX_GEOMETRY_OUTPUT_VERTICES_ARB] = "MAX_GEOMETRY_OUTPUT_VERTICES_ARB"
	enumLookup[LINEAR] = "LINEAR"
	enumLookup[FILTER] = "FILTER"
	enumLookup[CURRENT_TIME_NV] = "CURRENT_TIME_NV"
	enumLookup[BLUE_SCALE] = "BLUE_SCALE"
	enumLookup[RGBA8] = "RGBA8"
	enumLookup[T2F_C4F_N3F_V3F] = "T2F_C4F_N3F_V3F"
	enumLookup[CONTINUOUS_AMD] = "CONTINUOUS_AMD"
	enumLookup[PATH_ERROR_POSITION_NV] = "PATH_ERROR_POSITION_NV"
	enumLookup[BACK] = "BACK"
	enumLookup[TEXTURE_3D_BINDING_OES] = "TEXTURE_3D_BINDING_OES"
	enumLookup[UNSIGNED_INT_S8_S8_8_8_NV] = "UNSIGNED_INT_S8_S8_8_8_NV"
	enumLookup[MAX_IMAGE_SAMPLES_EXT] = "MAX_IMAGE_SAMPLES_EXT"
	enumLookup[RGB32F] = "RGB32F"
	enumLookup[QUERY_RESULT_EXT] = "QUERY_RESULT_EXT"
	enumLookup[COLOR_ATTACHMENT4_EXT] = "COLOR_ATTACHMENT4_EXT"
	enumLookup[SGIS_pixel_texture] = "SGIS_pixel_texture"
	enumLookup[SAMPLE_BUFFERS] = "SAMPLE_BUFFERS"
	enumLookup[LAYER_PROVOKING_VERTEX] = "LAYER_PROVOKING_VERTEX"
	enumLookup[R1UI_T2F_V3F_SUN] = "R1UI_T2F_V3F_SUN"
	enumLookup[MAP1_VERTEX_ATTRIB5_4_NV] = "MAP1_VERTEX_ATTRIB5_4_NV"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR"
	enumLookup[CURRENT_RASTER_POSITION_VALID] = "CURRENT_RASTER_POSITION_VALID"
	enumLookup[IMAGE_SCALE_Y_HP] = "IMAGE_SCALE_Y_HP"
	enumLookup[COMPRESSED_RED_RGTC1_EXT] = "COMPRESSED_RED_RGTC1_EXT"
	enumLookup[UNSIGNED_INT_SAMPLER_BUFFER_EXT] = "UNSIGNED_INT_SAMPLER_BUFFER_EXT"
	enumLookup[COVERAGE_COMPONENT_NV] = "COVERAGE_COMPONENT_NV"
	enumLookup[INT_IMAGE_CUBE] = "INT_IMAGE_CUBE"
	enumLookup[RGB5_EXT] = "RGB5_EXT"
	enumLookup[POINT_SIZE_ARRAY_STRIDE_OES] = "POINT_SIZE_ARRAY_STRIDE_OES"
	enumLookup[SHORT] = "SHORT"
	enumLookup[TEXTURE_BLUE_SIZE_EXT] = "TEXTURE_BLUE_SIZE_EXT"
	enumLookup[DEPTH_COMPONENT16_OES] = "DEPTH_COMPONENT16_OES"
	enumLookup[MODULATE_SUBTRACT_ATI] = "MODULATE_SUBTRACT_ATI"
	enumLookup[UNSIGNED_INT_IMAGE_2D_RECT] = "UNSIGNED_INT_IMAGE_2D_RECT"
	enumLookup[IS_PER_PATCH] = "IS_PER_PATCH"
	enumLookup[DOMAIN] = "DOMAIN"
	enumLookup[Q] = "Q"
	enumLookup[TEXTURE21] = "TEXTURE21"
	enumLookup[DRAW_BUFFER13] = "DRAW_BUFFER13"
	enumLookup[TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH] = "TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH"
	enumLookup[DRAW_FRAMEBUFFER_ANGLE] = "DRAW_FRAMEBUFFER_ANGLE"
	enumLookup[MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV] = "MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV"
	enumLookup[TRIANGLE_LIST_SUN] = "TRIANGLE_LIST_SUN"
	enumLookup[READ_FRAMEBUFFER_BINDING_NV] = "READ_FRAMEBUFFER_BINDING_NV"
	enumLookup[MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB] = "MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB"
	enumLookup[RELATIVE_ARC_TO_NV] = "RELATIVE_ARC_TO_NV"
	enumLookup[EXT_texture] = "EXT_texture"
	enumLookup[INVERT] = "INVERT"
	enumLookup[SAMPLE_MASK_SGIS] = "SAMPLE_MASK_SGIS"
	enumLookup[DEBUG_TYPE_DEPRECATED_BEHAVIOR] = "DEBUG_TYPE_DEPRECATED_BEHAVIOR"
	enumLookup[FOG_COORDINATE_ARRAY_EXT] = "FOG_COORDINATE_ARRAY_EXT"
	enumLookup[TEXTURE25_ARB] = "TEXTURE25_ARB"
	enumLookup[DOUBLE_MAT2] = "DOUBLE_MAT2"
	enumLookup[CLOSE_PATH_NV] = "CLOSE_PATH_NV"
	enumLookup[OBJECT_LINEAR] = "OBJECT_LINEAR"
	enumLookup[INTENSITY16] = "INTENSITY16"
	enumLookup[PIXEL_SUBSAMPLE_4444_SGIX] = "PIXEL_SUBSAMPLE_4444_SGIX"
	enumLookup[TEXTURE_BUFFER_DATA_STORE_BINDING_ARB] = "TEXTURE_BUFFER_DATA_STORE_BINDING_ARB"
	enumLookup[ALPHA16UI_EXT] = "ALPHA16UI_EXT"
	enumLookup[PATH_END_CAPS_NV] = "PATH_END_CAPS_NV"
	enumLookup[WAIT_FAILED] = "WAIT_FAILED"
	enumLookup[CULL_FACE_MODE] = "CULL_FACE_MODE"
	enumLookup[DRAW_BUFFER12_NV] = "DRAW_BUFFER12_NV"
	enumLookup[MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV] = "MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV"
	enumLookup[SPARE0_PLUS_SECONDARY_COLOR_NV] = "SPARE0_PLUS_SECONDARY_COLOR_NV"
	enumLookup[INVARIANT_EXT] = "INVARIANT_EXT"
	enumLookup[RELATIVE_CUBIC_CURVE_TO_NV] = "RELATIVE_CUBIC_CURVE_TO_NV"
	enumLookup[COLOR3_BIT_PGI] = "COLOR3_BIT_PGI"
	enumLookup[UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER] = "UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER"
	enumLookup[SECONDARY_COLOR_ARRAY_BUFFER_BINDING] = "SECONDARY_COLOR_ARRAY_BUFFER_BINDING"
	enumLookup[REG_6_ATI] = "REG_6_ATI"
	enumLookup[NEXT_BUFFER_NV] = "NEXT_BUFFER_NV"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV] = "TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV"
	enumLookup[INTENSITY16_SNORM] = "INTENSITY16_SNORM"
	enumLookup[LARGE_CCW_ARC_TO_NV] = "LARGE_CCW_ARC_TO_NV"
	enumLookup[GEQUAL] = "GEQUAL"
	enumLookup[EDGE_FLAG_ARRAY_POINTER] = "EDGE_FLAG_ARRAY_POINTER"
	enumLookup[TEXTURE_FILTER4_SIZE_SGIS] = "TEXTURE_FILTER4_SIZE_SGIS"
	enumLookup[TRANSPOSE_TEXTURE_MATRIX_ARB] = "TRANSPOSE_TEXTURE_MATRIX_ARB"
	enumLookup[MAX_UNIFORM_BLOCK_SIZE] = "MAX_UNIFORM_BLOCK_SIZE"
	enumLookup[TEXTURE_SRGB_DECODE_EXT] = "TEXTURE_SRGB_DECODE_EXT"
	enumLookup[POST_COLOR_MATRIX_GREEN_SCALE_SGI] = "POST_COLOR_MATRIX_GREEN_SCALE_SGI"
	enumLookup[SGIX_ycrcb] = "SGIX_ycrcb"
	enumLookup[UNSIGNED_INT64_NV] = "UNSIGNED_INT64_NV"
	enumLookup[IMAGE_CLASS_4_X_32] = "IMAGE_CLASS_4_X_32"
	enumLookup[TEXTURE_BINDING_CUBE_MAP] = "TEXTURE_BINDING_CUBE_MAP"
	enumLookup[OUTPUT_TEXTURE_COORD3_EXT] = "OUTPUT_TEXTURE_COORD3_EXT"
	enumLookup[OUTPUT_TEXTURE_COORD24_EXT] = "OUTPUT_TEXTURE_COORD24_EXT"
	enumLookup[IMAGE_TRANSLATE_X_HP] = "IMAGE_TRANSLATE_X_HP"
	enumLookup[COLOR_RENDERABLE] = "COLOR_RENDERABLE"
	enumLookup[SAMPLER_1D] = "SAMPLER_1D"
	enumLookup[COMPRESSED_RGBA8_ETC2_EAC] = "COMPRESSED_RGBA8_ETC2_EAC"
	enumLookup[VERTEX_CONSISTENT_HINT_PGI] = "VERTEX_CONSISTENT_HINT_PGI"
	enumLookup[IMAGE_MIN_FILTER_HP] = "IMAGE_MIN_FILTER_HP"
	enumLookup[MAX_RENDERBUFFER_SIZE_OES] = "MAX_RENDERBUFFER_SIZE_OES"
	enumLookup[INT16_VEC4_NV] = "INT16_VEC4_NV"
	enumLookup[RELATIVE_SMALL_CW_ARC_TO_NV] = "RELATIVE_SMALL_CW_ARC_TO_NV"
	enumLookup[SYNC_FENCE] = "SYNC_FENCE"
	enumLookup[MAX_SPARSE_ARRAY_TEXTURE_LAYERS] = "MAX_SPARSE_ARRAY_TEXTURE_LAYERS"
	enumLookup[TEXTURE12_ARB] = "TEXTURE12_ARB"
	enumLookup[VERTEX_STREAM0_ATI] = "VERTEX_STREAM0_ATI"
	enumLookup[FRAGMENT_ALPHA_MODULATE_IMG] = "FRAGMENT_ALPHA_MODULATE_IMG"
	enumLookup[RENDERBUFFER_SAMPLES_NV] = "RENDERBUFFER_SAMPLES_NV"
	enumLookup[ALPHA_SNORM] = "ALPHA_SNORM"
	enumLookup[INTENSITY_EXT] = "INTENSITY_EXT"
	enumLookup[POINT_FADE_THRESHOLD_SIZE] = "POINT_FADE_THRESHOLD_SIZE"
	enumLookup[CONSTANT_BORDER_HP] = "CONSTANT_BORDER_HP"
	enumLookup[TEXTURE29] = "TEXTURE29"
	enumLookup[MODELVIEW6_ARB] = "MODELVIEW6_ARB"
	enumLookup[MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS] = "MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS"
	enumLookup[ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER] = "ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER"
	enumLookup[FRAGMENT_LIGHT4_SGIX] = "FRAGMENT_LIGHT4_SGIX"
	enumLookup[DRAW_BUFFER5_NV] = "DRAW_BUFFER5_NV"
	enumLookup[REG_21_ATI] = "REG_21_ATI"
	enumLookup[INDEX_ARRAY_EXT] = "INDEX_ARRAY_EXT"
	enumLookup[TEXTURE_COORD_ARRAY_EXT] = "TEXTURE_COORD_ARRAY_EXT"
	enumLookup[MINOR_VERSION] = "MINOR_VERSION"
	enumLookup[FRAGMENTS_INSTRUMENT_MAX_SGIX] = "FRAGMENTS_INSTRUMENT_MAX_SGIX"
	enumLookup[TEXTURE_COMPRESSED_IMAGE_SIZE] = "TEXTURE_COMPRESSED_IMAGE_SIZE"
	enumLookup[UNPACK_RESAMPLE_OML] = "UNPACK_RESAMPLE_OML"
	enumLookup[IMAGE_MAG_FILTER_HP] = "IMAGE_MAG_FILTER_HP"
	enumLookup[SOURCE1_RGB] = "SOURCE1_RGB"
	enumLookup[DRAW_FRAMEBUFFER_BINDING_EXT] = "DRAW_FRAMEBUFFER_BINDING_EXT"
	enumLookup[COLOR_ATTACHMENT10] = "COLOR_ATTACHMENT10"
	enumLookup[SAMPLE_MASK_VALUE] = "SAMPLE_MASK_VALUE"
	enumLookup[LINE_STRIP_ADJACENCY] = "LINE_STRIP_ADJACENCY"
	enumLookup[COLOR_ARRAY_STRIDE_EXT] = "COLOR_ARRAY_STRIDE_EXT"
	enumLookup[TEXTURE_SHADOW] = "TEXTURE_SHADOW"
	enumLookup[OBJECT_BUFFER_SIZE_ATI] = "OBJECT_BUFFER_SIZE_ATI"
	enumLookup[RGBA_SIGNED_COMPONENTS_EXT] = "RGBA_SIGNED_COMPONENTS_EXT"
	enumLookup[COLOR_ATTACHMENT5_NV] = "COLOR_ATTACHMENT5_NV"
	enumLookup[MAX_FRAGMENT_SHADER_STORAGE_BLOCKS] = "MAX_FRAGMENT_SHADER_STORAGE_BLOCKS"
	enumLookup[HALF_APPLE] = "HALF_APPLE"
	enumLookup[UNPACK_IMAGE_HEIGHT_EXT] = "UNPACK_IMAGE_HEIGHT_EXT"
	enumLookup[INTERNALFORMAT_DEPTH_SIZE] = "INTERNALFORMAT_DEPTH_SIZE"
	enumLookup[REG_18_ATI] = "REG_18_ATI"
	enumLookup[SIGNALED_APPLE] = "SIGNALED_APPLE"
	enumLookup[REPEAT] = "REPEAT"
	enumLookup[RGB10_A2_EXT] = "RGB10_A2_EXT"
	enumLookup[POINT_DISTANCE_ATTENUATION_ARB] = "POINT_DISTANCE_ATTENUATION_ARB"
	enumLookup[INTERNALFORMAT_RED_TYPE] = "INTERNALFORMAT_RED_TYPE"
	enumLookup[SAMPLER_2D_RECT_SHADOW_ARB] = "SAMPLER_2D_RECT_SHADOW_ARB"
	enumLookup[SCALED_RESOLVE_NICEST_EXT] = "SCALED_RESOLVE_NICEST_EXT"
	enumLookup[PACK_SKIP_PIXELS] = "PACK_SKIP_PIXELS"
	enumLookup[PIXEL_TILE_BEST_ALIGNMENT_SGIX] = "PIXEL_TILE_BEST_ALIGNMENT_SGIX"
	enumLookup[MAX_TEXTURE_UNITS] = "MAX_TEXTURE_UNITS"
	enumLookup[RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV] = "RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV"
	enumLookup[PROGRAM_ATTRIB_COMPONENTS_NV] = "PROGRAM_ATTRIB_COMPONENTS_NV"
	enumLookup[TEXTURE_HEIGHT_QCOM] = "TEXTURE_HEIGHT_QCOM"
	enumLookup[VIEW_CLASS_128_BITS] = "VIEW_CLASS_128_BITS"
	enumLookup[SAMPLES_PASSED_ARB] = "SAMPLES_PASSED_ARB"
	enumLookup[COMPRESSED_SIGNED_RED_RGTC1] = "COMPRESSED_SIGNED_RED_RGTC1"
	enumLookup[AUTO_NORMAL] = "AUTO_NORMAL"
	enumLookup[SMOOTH] = "SMOOTH"
	enumLookup[BUFFER] = "BUFFER"
	enumLookup[MODELVIEW1_STACK_DEPTH_EXT] = "MODELVIEW1_STACK_DEPTH_EXT"
	enumLookup[UNSIGNED_SHORT_8_8_REV_APPLE] = "UNSIGNED_SHORT_8_8_REV_APPLE"
	enumLookup[OFFSET_TEXTURE_SCALE_NV] = "OFFSET_TEXTURE_SCALE_NV"
	enumLookup[OBJECT_VALIDATE_STATUS_ARB] = "OBJECT_VALIDATE_STATUS_ARB"
	enumLookup[MAP_READ_BIT_EXT] = "MAP_READ_BIT_EXT"
	enumLookup[TEXTURE_VIEW_NUM_LEVELS] = "TEXTURE_VIEW_NUM_LEVELS"
	enumLookup[COMBINER5_NV] = "COMBINER5_NV"
	enumLookup[SURFACE_MAPPED_NV] = "SURFACE_MAPPED_NV"
	enumLookup[MAX_SAMPLES_NV] = "MAX_SAMPLES_NV"
	enumLookup[UNSIGNED_INT_SAMPLER_1D_EXT] = "UNSIGNED_INT_SAMPLER_1D_EXT"
	enumLookup[MAP1_BINORMAL_EXT] = "MAP1_BINORMAL_EXT"
	enumLookup[SAMPLE_POSITION_NV] = "SAMPLE_POSITION_NV"
	enumLookup[BLOCK_INDEX] = "BLOCK_INDEX"
	enumLookup[SRC_COLOR] = "SRC_COLOR"
	enumLookup[INVALID_FRAMEBUFFER_OPERATION_OES] = "INVALID_FRAMEBUFFER_OPERATION_OES"
	enumLookup[TEXTURE31] = "TEXTURE31"
	enumLookup[UNPACK_SUBSAMPLE_RATE_SGIX] = "UNPACK_SUBSAMPLE_RATE_SGIX"
	enumLookup[INTENSITY_FLOAT32_APPLE] = "INTENSITY_FLOAT32_APPLE"
	enumLookup[SKIP_DECODE_EXT] = "SKIP_DECODE_EXT"
	enumLookup[TEXTURE_NUM_LEVELS_QCOM] = "TEXTURE_NUM_LEVELS_QCOM"
	enumLookup[RENDERBUFFER_STENCIL_SIZE_EXT] = "RENDERBUFFER_STENCIL_SIZE_EXT"
	enumLookup[INTENSITY16I_EXT] = "INTENSITY16I_EXT"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Z_EXT] = "TEXTURE_CUBE_MAP_POSITIVE_Z_EXT"
	enumLookup[SHADER_CONSISTENT_NV] = "SHADER_CONSISTENT_NV"
	enumLookup[PROGRAM_ALU_INSTRUCTIONS_ARB] = "PROGRAM_ALU_INSTRUCTIONS_ARB"
	enumLookup[SAMPLER_1D_SHADOW] = "SAMPLER_1D_SHADOW"
	enumLookup[TEXTURE_ALPHA_SIZE_EXT] = "TEXTURE_ALPHA_SIZE_EXT"
	enumLookup[TEXTURE_3D_EXT] = "TEXTURE_3D_EXT"
	enumLookup[OPERAND1_ALPHA_ARB] = "OPERAND1_ALPHA_ARB"
	enumLookup[SIGNED_RGB_UNSIGNED_ALPHA_NV] = "SIGNED_RGB_UNSIGNED_ALPHA_NV"
	enumLookup[FONT_HAS_KERNING_BIT_NV] = "FONT_HAS_KERNING_BIT_NV"
	enumLookup[ONE_MINUS_DST_ALPHA] = "ONE_MINUS_DST_ALPHA"
	enumLookup[RG16F] = "RG16F"
	enumLookup[STENCIL_INDEX4_OES] = "STENCIL_INDEX4_OES"
	enumLookup[ANY_SAMPLES_PASSED_CONSERVATIVE] = "ANY_SAMPLES_PASSED_CONSERVATIVE"
	enumLookup[SAMPLER_CUBE_SHADOW_EXT] = "SAMPLER_CUBE_SHADOW_EXT"
	enumLookup[RGB10_EXT] = "RGB10_EXT"
	enumLookup[SLUMINANCE8_ALPHA8_NV] = "SLUMINANCE8_ALPHA8_NV"
	enumLookup[VIDEO_BUFFER_INTERNAL_FORMAT_NV] = "VIDEO_BUFFER_INTERNAL_FORMAT_NV"
	enumLookup[FONT_ASCENDER_BIT_NV] = "FONT_ASCENDER_BIT_NV"
	enumLookup[LOGIC_OP] = "LOGIC_OP"
	enumLookup[SAMPLE_BUFFERS_ARB] = "SAMPLE_BUFFERS_ARB"
	enumLookup[PIXEL_TILE_CACHE_SIZE_SGIX] = "PIXEL_TILE_CACHE_SIZE_SGIX"
	enumLookup[PIXEL_TEX_GEN_Q_ROUND_SGIX] = "PIXEL_TEX_GEN_Q_ROUND_SGIX"
	enumLookup[DEPTH_COMPONENT24_SGIX] = "DEPTH_COMPONENT24_SGIX"
	enumLookup[DEBUG_SEVERITY_NOTIFICATION] = "DEBUG_SEVERITY_NOTIFICATION"
	enumLookup[MAP1_TANGENT_EXT] = "MAP1_TANGENT_EXT"
	enumLookup[DEPTH_STENCIL_EXT] = "DEPTH_STENCIL_EXT"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Y_OES] = "TEXTURE_CUBE_MAP_NEGATIVE_Y_OES"
	enumLookup[SAMPLER_1D_ARB] = "SAMPLER_1D_ARB"
	enumLookup[POST_CONVOLUTION_GREEN_SCALE] = "POST_CONVOLUTION_GREEN_SCALE"
	enumLookup[PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI] = "PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI"
	enumLookup[TANGENT_ARRAY_TYPE_EXT] = "TANGENT_ARRAY_TYPE_EXT"
	enumLookup[INT_SAMPLER_2D_RECT_EXT] = "INT_SAMPLER_2D_RECT_EXT"
	enumLookup[FRAMEBUFFER_DEFAULT_LAYERS] = "FRAMEBUFFER_DEFAULT_LAYERS"
	enumLookup[OBJECT_POINT_SGIS] = "OBJECT_POINT_SGIS"
	enumLookup[DEPTH_STENCIL] = "DEPTH_STENCIL"
	enumLookup[MATRIX18_ARB] = "MATRIX18_ARB"
	enumLookup[RGB_INTEGER_EXT] = "RGB_INTEGER_EXT"
	enumLookup[SOURCE0_RGB_ARB] = "SOURCE0_RGB_ARB"
	enumLookup[MAX_VERTEX_ATTRIBS] = "MAX_VERTEX_ATTRIBS"
	enumLookup[REG_5_ATI] = "REG_5_ATI"
	enumLookup[GL_422_EXT] = "GL_422_EXT"
	enumLookup[ATC_RGBA_INTERPOLATED_ALPHA_AMD] = "ATC_RGBA_INTERPOLATED_ALPHA_AMD"
	enumLookup[MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS] = "MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS"
	enumLookup[RENDERBUFFER_INTERNAL_FORMAT] = "RENDERBUFFER_INTERNAL_FORMAT"
	enumLookup[RGB32I] = "RGB32I"
	enumLookup[MAX_GEOMETRY_UNIFORM_COMPONENTS] = "MAX_GEOMETRY_UNIFORM_COMPONENTS"
	enumLookup[SAMPLER_BUFFER_AMD] = "SAMPLER_BUFFER_AMD"
	enumLookup[LEFT] = "LEFT"
	enumLookup[DEBUG_TYPE_PORTABILITY_ARB] = "DEBUG_TYPE_PORTABILITY_ARB"
	enumLookup[VIEW_CLASS_S3TC_DXT1_RGB] = "VIEW_CLASS_S3TC_DXT1_RGB"
	enumLookup[INTERPOLATE_EXT] = "INTERPOLATE_EXT"
	enumLookup[SAMPLES_PASSED] = "SAMPLES_PASSED"
	enumLookup[SHADER_INCLUDE_ARB] = "SHADER_INCLUDE_ARB"
	enumLookup[RECLAIM_MEMORY_HINT_PGI] = "RECLAIM_MEMORY_HINT_PGI"
	enumLookup[TESS_EVALUATION_SHADER_BIT] = "TESS_EVALUATION_SHADER_BIT"
	enumLookup[SAMPLE_COVERAGE_INVERT_ARB] = "SAMPLE_COVERAGE_INVERT_ARB"
	enumLookup[DEPTH_COMPONENT16_ARB] = "DEPTH_COMPONENT16_ARB"
	enumLookup[VERTEX_PROGRAM_TWO_SIDE_NV] = "VERTEX_PROGRAM_TWO_SIDE_NV"
	enumLookup[MODELVIEW20_ARB] = "MODELVIEW20_ARB"
	enumLookup[MATRIX_INDEX_ARRAY_TYPE_OES] = "MATRIX_INDEX_ARRAY_TYPE_OES"
	enumLookup[ELEMENT_ARRAY_BUFFER_BINDING] = "ELEMENT_ARRAY_BUFFER_BINDING"
	enumLookup[MAT_AMBIENT_AND_DIFFUSE_BIT_PGI] = "MAT_AMBIENT_AND_DIFFUSE_BIT_PGI"
	enumLookup[SGIX_scalebias_hint] = "SGIX_scalebias_hint"
	enumLookup[CLIP_DISTANCE6] = "CLIP_DISTANCE6"
	enumLookup[TEXTURE_LOD_BIAS_T_SGIX] = "TEXTURE_LOD_BIAS_T_SGIX"
	enumLookup[BOOL_VEC3_ARB] = "BOOL_VEC3_ARB"
	enumLookup[FONT_MAX_ADVANCE_WIDTH_BIT_NV] = "FONT_MAX_ADVANCE_WIDTH_BIT_NV"
	enumLookup[UNSIGNED_INT_24_8_OES] = "UNSIGNED_INT_24_8_OES"
	enumLookup[DYNAMIC_ATI] = "DYNAMIC_ATI"
	enumLookup[LUMINANCE_ALPHA8UI_EXT] = "LUMINANCE_ALPHA8UI_EXT"
	enumLookup[POST_CONVOLUTION_ALPHA_SCALE] = "POST_CONVOLUTION_ALPHA_SCALE"
	enumLookup[INTENSITY4_EXT] = "INTENSITY4_EXT"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_X] = "TEXTURE_CUBE_MAP_NEGATIVE_X"
	enumLookup[COLOR_ARRAY_BUFFER_BINDING] = "COLOR_ARRAY_BUFFER_BINDING"
	enumLookup[DOT4_ATI] = "DOT4_ATI"
	enumLookup[INTERLEAVED_ATTRIBS_EXT] = "INTERLEAVED_ATTRIBS_EXT"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT] = "TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT"
	enumLookup[STENCIL_INDEX1_EXT] = "STENCIL_INDEX1_EXT"
	enumLookup[PACK_ROW_LENGTH] = "PACK_ROW_LENGTH"
	enumLookup[PIXEL_MAG_FILTER_EXT] = "PIXEL_MAG_FILTER_EXT"
	enumLookup[UNPACK_RESAMPLE_SGIX] = "UNPACK_RESAMPLE_SGIX"
	enumLookup[OPERAND0_ALPHA_ARB] = "OPERAND0_ALPHA_ARB"
	enumLookup[VIDEO_CAPTURE_FRAME_WIDTH_NV] = "VIDEO_CAPTURE_FRAME_WIDTH_NV"
	enumLookup[TEXTURE_FIXED_SAMPLE_LOCATIONS] = "TEXTURE_FIXED_SAMPLE_LOCATIONS"
	enumLookup[BLEND_COLOR_EXT] = "BLEND_COLOR_EXT"
	enumLookup[TEXTURE22] = "TEXTURE22"
	enumLookup[FLOAT_MAT3x4] = "FLOAT_MAT3x4"
	enumLookup[PROXY_TEXTURE_1D_ARRAY] = "PROXY_TEXTURE_1D_ARRAY"
	enumLookup[EXT_blend_logic_op] = "EXT_blend_logic_op"
	enumLookup[PIXEL_MAP_B_TO_B_SIZE] = "PIXEL_MAP_B_TO_B_SIZE"
	enumLookup[VARIABLE_D_NV] = "VARIABLE_D_NV"
	enumLookup[CURRENT_OCCLUSION_QUERY_ID_NV] = "CURRENT_OCCLUSION_QUERY_ID_NV"
	enumLookup[FORMAT_SUBSAMPLE_244_244_OML] = "FORMAT_SUBSAMPLE_244_244_OML"
	enumLookup[COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT] = "COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT"
	enumLookup[COLOR_SAMPLES_NV] = "COLOR_SAMPLES_NV"
	enumLookup[HISTOGRAM_FORMAT_EXT] = "HISTOGRAM_FORMAT_EXT"
	enumLookup[MAX_ELEMENTS_INDICES] = "MAX_ELEMENTS_INDICES"
	enumLookup[VERTEX_ATTRIB_ARRAY6_NV] = "VERTEX_ATTRIB_ARRAY6_NV"
	enumLookup[OFFSET_TEXTURE_2D_NV] = "OFFSET_TEXTURE_2D_NV"
	enumLookup[BLUE_INTEGER_EXT] = "BLUE_INTEGER_EXT"
	enumLookup[UNIFORM_BUFFER_BINDING_EXT] = "UNIFORM_BUFFER_BINDING_EXT"
	enumLookup[MAX_FRAMEBUFFER_SAMPLES] = "MAX_FRAMEBUFFER_SAMPLES"
	enumLookup[DEBUG_CALLBACK_USER_PARAM] = "DEBUG_CALLBACK_USER_PARAM"
	enumLookup[NEVER] = "NEVER"
	enumLookup[AUX_BUFFERS] = "AUX_BUFFERS"
	enumLookup[UNSIGNED_INT8_VEC4_NV] = "UNSIGNED_INT8_VEC4_NV"
	enumLookup[TESS_CONTROL_SUBROUTINE] = "TESS_CONTROL_SUBROUTINE"
	enumLookup[LINE_LOOP] = "LINE_LOOP"
	enumLookup[ONE_MINUS_CONSTANT_COLOR] = "ONE_MINUS_CONSTANT_COLOR"
	enumLookup[SIGNED_INTENSITY8_NV] = "SIGNED_INTENSITY8_NV"
	enumLookup[UNSIGNED_INT_IMAGE_3D_EXT] = "UNSIGNED_INT_IMAGE_3D_EXT"
	enumLookup[MAX_FRAGMENT_INPUT_COMPONENTS] = "MAX_FRAGMENT_INPUT_COMPONENTS"
	enumLookup[SEPARATE_SPECULAR_COLOR] = "SEPARATE_SPECULAR_COLOR"
	enumLookup[CLIENT_ACTIVE_TEXTURE_ARB] = "CLIENT_ACTIVE_TEXTURE_ARB"
	enumLookup[ALPHA_INTEGER_EXT] = "ALPHA_INTEGER_EXT"
	enumLookup[UNPACK_COMPRESSED_BLOCK_WIDTH] = "UNPACK_COMPRESSED_BLOCK_WIDTH"
	enumLookup[YCRCB_444_SGIX] = "YCRCB_444_SGIX"
	enumLookup[COLOR_ATTACHMENT0_OES] = "COLOR_ATTACHMENT0_OES"
	enumLookup[TRIANGLE_STRIP_ADJACENCY_EXT] = "TRIANGLE_STRIP_ADJACENCY_EXT"
	enumLookup[RGB] = "RGB"
	enumLookup[POST_CONVOLUTION_RED_SCALE_EXT] = "POST_CONVOLUTION_RED_SCALE_EXT"
	enumLookup[HISTOGRAM_ALPHA_SIZE_EXT] = "HISTOGRAM_ALPHA_SIZE_EXT"
	enumLookup[GENERATE_MIPMAP_HINT_SGIS] = "GENERATE_MIPMAP_HINT_SGIS"
	enumLookup[SHADER_OBJECT_EXT] = "SHADER_OBJECT_EXT"
	enumLookup[MAX_PATCH_VERTICES] = "MAX_PATCH_VERTICES"
	enumLookup[INTERLACE_SGIX] = "INTERLACE_SGIX"
	enumLookup[DUAL_LUMINANCE12_SGIS] = "DUAL_LUMINANCE12_SGIS"
	enumLookup[DEBUG_NEXT_LOGGED_MESSAGE_LENGTH] = "DEBUG_NEXT_LOGGED_MESSAGE_LENGTH"
	enumLookup[IMAGE_CLASS_11_11_10] = "IMAGE_CLASS_11_11_10"
	enumLookup[TEXTURE0_ARB] = "TEXTURE0_ARB"
	enumLookup[ACTIVE_TEXTURE_ARB] = "ACTIVE_TEXTURE_ARB"
	enumLookup[DEBUG_TYPE_UNDEFINED_BEHAVIOR] = "DEBUG_TYPE_UNDEFINED_BEHAVIOR"
	enumLookup[SET_AMD] = "SET_AMD"
	enumLookup[ELEMENT_ARRAY_APPLE] = "ELEMENT_ARRAY_APPLE"
	enumLookup[BLUE_INTEGER] = "BLUE_INTEGER"
	enumLookup[SGIX_polynomial_ffd] = "SGIX_polynomial_ffd"
	enumLookup[T4F_V4F] = "T4F_V4F"
	enumLookup[EVAL_VERTEX_ATTRIB12_NV] = "EVAL_VERTEX_ATTRIB12_NV"
	enumLookup[RENDERBUFFER_SAMPLES_EXT] = "RENDERBUFFER_SAMPLES_EXT"
	enumLookup[SGIX_shadow] = "SGIX_shadow"
	enumLookup[ALPHA8] = "ALPHA8"
	enumLookup[COMBINER1_NV] = "COMBINER1_NV"
	enumLookup[UNSIGNED_INT_SAMPLER_CUBE] = "UNSIGNED_INT_SAMPLER_CUBE"
	enumLookup[INT_SAMPLER_CUBE_MAP_ARRAY] = "INT_SAMPLER_CUBE_MAP_ARRAY"
	enumLookup[MAX_TESS_EVALUATION_IMAGE_UNIFORMS] = "MAX_TESS_EVALUATION_IMAGE_UNIFORMS"
	enumLookup[IMAGE_PIXEL_TYPE] = "IMAGE_PIXEL_TYPE"
	enumLookup[MODELVIEW5_ARB] = "MODELVIEW5_ARB"
	enumLookup[DEBUG_PRINT_MESA] = "DEBUG_PRINT_MESA"
	enumLookup[SYNC_STATUS] = "SYNC_STATUS"
	enumLookup[NUM_ACTIVE_VARIABLES] = "NUM_ACTIVE_VARIABLES"
	enumLookup[DRAW_BUFFER0_ARB] = "DRAW_BUFFER0_ARB"
	enumLookup[MATRIX17_ARB] = "MATRIX17_ARB"
	enumLookup[ACTIVE_SUBROUTINES] = "ACTIVE_SUBROUTINES"
	enumLookup[VARIABLE_E_NV] = "VARIABLE_E_NV"
	enumLookup[FLOAT_VEC2] = "FLOAT_VEC2"
	enumLookup[FLOAT_VEC2_ARB] = "FLOAT_VEC2_ARB"
	enumLookup[PROJECTION_MATRIX] = "PROJECTION_MATRIX"
	enumLookup[CONSTANT_ALPHA] = "CONSTANT_ALPHA"
	enumLookup[RGB5_A1_EXT] = "RGB5_A1_EXT"
	enumLookup[ARRAY_ELEMENT_LOCK_COUNT_EXT] = "ARRAY_ELEMENT_LOCK_COUNT_EXT"
	enumLookup[MAX_ASYNC_TEX_IMAGE_SGIX] = "MAX_ASYNC_TEX_IMAGE_SGIX"
	enumLookup[RELATIVE_VERTICAL_LINE_TO_NV] = "RELATIVE_VERTICAL_LINE_TO_NV"
	enumLookup[RGBA4_OES] = "RGBA4_OES"
	enumLookup[TEXTURE_MAX_LOD] = "TEXTURE_MAX_LOD"
	enumLookup[INDEX_MATERIAL_EXT] = "INDEX_MATERIAL_EXT"
	enumLookup[COMPRESSED_RGB_ARB] = "COMPRESSED_RGB_ARB"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_X_ARB] = "TEXTURE_CUBE_MAP_NEGATIVE_X_ARB"
	enumLookup[PACK_SUBSAMPLE_RATE_SGIX] = "PACK_SUBSAMPLE_RATE_SGIX"
	enumLookup[TEXTURE_DEPTH_SIZE_ARB] = "TEXTURE_DEPTH_SIZE_ARB"
	enumLookup[COPY_WRITE_BUFFER_BINDING] = "COPY_WRITE_BUFFER_BINDING"
	enumLookup[UNSIGNED_INT64_VEC4_NV] = "UNSIGNED_INT64_VEC4_NV"
	enumLookup[PROXY_TEXTURE_1D_EXT] = "PROXY_TEXTURE_1D_EXT"
	enumLookup[DUAL_LUMINANCE_ALPHA8_SGIS] = "DUAL_LUMINANCE_ALPHA8_SGIS"
	enumLookup[SAMPLE_MASK_EXT] = "SAMPLE_MASK_EXT"
	enumLookup[INDEX_TEST_EXT] = "INDEX_TEST_EXT"
	enumLookup[IMAGE_CLASS_4_X_8] = "IMAGE_CLASS_4_X_8"
	enumLookup[BUMP_TARGET_ATI] = "BUMP_TARGET_ATI"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY] = "UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY"
	enumLookup[UNIFORM_BUFFER] = "UNIFORM_BUFFER"
	enumLookup[MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS] = "MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS"
	enumLookup[COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV] = "COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV"
	enumLookup[QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT] = "QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT"
	enumLookup[QUADRATIC_ATTENUATION] = "QUADRATIC_ATTENUATION"
	enumLookup[INT64_NV] = "INT64_NV"
	enumLookup[TEXTURE_CLIPMAP_CENTER_SGIX] = "TEXTURE_CLIPMAP_CENTER_SGIX"
	enumLookup[FOG_COORDINATE_ARRAY_STRIDE] = "FOG_COORDINATE_ARRAY_STRIDE"
	enumLookup[UNIFORM_SIZE] = "UNIFORM_SIZE"
	enumLookup[FRAMEBUFFER_COMPLETE] = "FRAMEBUFFER_COMPLETE"
	enumLookup[SAMPLER_2D_MULTISAMPLE] = "SAMPLER_2D_MULTISAMPLE"
	enumLookup[COLOR_INDEX] = "COLOR_INDEX"
	enumLookup[MAX_VARYING_COMPONENTS] = "MAX_VARYING_COMPONENTS"
	enumLookup[MULTIVIEW_EXT] = "MULTIVIEW_EXT"
	enumLookup[RG32F] = "RG32F"
	enumLookup[FRAGMENT_SHADER_DERIVATIVE_HINT_OES] = "FRAGMENT_SHADER_DERIVATIVE_HINT_OES"
	enumLookup[FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES] = "FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV] = "TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV"
	enumLookup[POST_COLOR_MATRIX_ALPHA_SCALE] = "POST_COLOR_MATRIX_ALPHA_SCALE"
	enumLookup[COLOR_TABLE_BIAS_SGI] = "COLOR_TABLE_BIAS_SGI"
	enumLookup[MAX_TEXTURE_LOD_BIAS] = "MAX_TEXTURE_LOD_BIAS"
	enumLookup[COLOR_CLEAR_UNCLAMPED_VALUE_ATI] = "COLOR_CLEAR_UNCLAMPED_VALUE_ATI"
	enumLookup[MAX_COMBINED_TEXTURE_IMAGE_UNITS] = "MAX_COMBINED_TEXTURE_IMAGE_UNITS"
	enumLookup[DEPTH_CLAMP_FAR_AMD] = "DEPTH_CLAMP_FAR_AMD"
	enumLookup[SGIX_instruments] = "SGIX_instruments"
	enumLookup[POST_CONVOLUTION_GREEN_BIAS_EXT] = "POST_CONVOLUTION_GREEN_BIAS_EXT"
	enumLookup[COMPRESSED_RGBA_S3TC_DXT5_ANGLE] = "COMPRESSED_RGBA_S3TC_DXT5_ANGLE"
	enumLookup[DU8DV8_ATI] = "DU8DV8_ATI"
	enumLookup[COMPARE_REF_TO_TEXTURE_EXT] = "COMPARE_REF_TO_TEXTURE_EXT"
	enumLookup[REG_10_ATI] = "REG_10_ATI"
	enumLookup[INTENSITY8UI_EXT] = "INTENSITY8UI_EXT"
	enumLookup[DOUBLE_MAT4x3_EXT] = "DOUBLE_MAT4x3_EXT"
	enumLookup[ATTRIB_STACK_DEPTH] = "ATTRIB_STACK_DEPTH"
	enumLookup[CLIP_DISTANCE3] = "CLIP_DISTANCE3"
	enumLookup[DEPTH_COMPONENT32_OES] = "DEPTH_COMPONENT32_OES"
	enumLookup[R16UI] = "R16UI"
	enumLookup[MATRIX10_ARB] = "MATRIX10_ARB"
	enumLookup[DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD] = "DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD"
	enumLookup[GL_2D] = "GL_2D"
	enumLookup[FRAGMENT_PROGRAM_NV] = "FRAGMENT_PROGRAM_NV"
	enumLookup[EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB] = "EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[UNSIGNED_NORMALIZED] = "UNSIGNED_NORMALIZED"
	enumLookup[MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT] = "MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT"
	enumLookup[FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES] = "FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES"
	enumLookup[RESTART_PATH_NV] = "RESTART_PATH_NV"
	enumLookup[DISPATCH_INDIRECT_BUFFER] = "DISPATCH_INDIRECT_BUFFER"
	enumLookup[FOG_COORDINATE_ARRAY_TYPE] = "FOG_COORDINATE_ARRAY_TYPE"
	enumLookup[E_TIMES_F_NV] = "E_TIMES_F_NV"
	enumLookup[DRAW_BUFFER13_NV] = "DRAW_BUFFER13_NV"
	enumLookup[UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT] = "UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT"
	enumLookup[INDEX_LOGIC_OP] = "INDEX_LOGIC_OP"
	enumLookup[LUMINANCE8I_EXT] = "LUMINANCE8I_EXT"
	enumLookup[MAX_PROGRAM_GENERIC_ATTRIBS_NV] = "MAX_PROGRAM_GENERIC_ATTRIBS_NV"
	enumLookup[QUADRATIC_CURVE_TO_NV] = "QUADRATIC_CURVE_TO_NV"
	enumLookup[OBJECT_TYPE_ARB] = "OBJECT_TYPE_ARB"
	enumLookup[UNSIGNED_INT_IMAGE_3D] = "UNSIGNED_INT_IMAGE_3D"
	enumLookup[DUAL_INTENSITY8_SGIS] = "DUAL_INTENSITY8_SGIS"
	enumLookup[DEPTH_STENCIL_TEXTURE_MODE] = "DEPTH_STENCIL_TEXTURE_MODE"
	enumLookup[SGIX_shadow_ambient] = "SGIX_shadow_ambient"
	enumLookup[LINEAR_ATTENUATION] = "LINEAR_ATTENUATION"
	enumLookup[REPLACE] = "REPLACE"
	enumLookup[FRAMEBUFFER_ATTACHMENT_RED_SIZE] = "FRAMEBUFFER_ATTACHMENT_RED_SIZE"
	enumLookup[PROXY_TEXTURE_RECTANGLE_ARB] = "PROXY_TEXTURE_RECTANGLE_ARB"
	enumLookup[TEXTURE_DEPTH_SIZE] = "TEXTURE_DEPTH_SIZE"
	enumLookup[TEXTURE14] = "TEXTURE14"
	enumLookup[VARIABLE_F_NV] = "VARIABLE_F_NV"
	enumLookup[FUNC_ADD_OES] = "FUNC_ADD_OES"
	enumLookup[TEXTURE_IMMUTABLE_LEVELS] = "TEXTURE_IMMUTABLE_LEVELS"
	enumLookup[VERTEX_PROGRAM_POINT_SIZE_ARB] = "VERTEX_PROGRAM_POINT_SIZE_ARB"
	enumLookup[MAX_COMBINED_SHADER_STORAGE_BLOCKS] = "MAX_COMBINED_SHADER_STORAGE_BLOCKS"
	enumLookup[POINT_DISTANCE_ATTENUATION] = "POINT_DISTANCE_ATTENUATION"
	enumLookup[COMPRESSED_RGB] = "COMPRESSED_RGB"
	enumLookup[DISCRETE_AMD] = "DISCRETE_AMD"
	enumLookup[AFFINE_3D_NV] = "AFFINE_3D_NV"
	enumLookup[TEXTURE_UPDATE_BARRIER_BIT_EXT] = "TEXTURE_UPDATE_BARRIER_BIT_EXT"
	enumLookup[INDEX_ARRAY_POINTER] = "INDEX_ARRAY_POINTER"
	enumLookup[TEXTURE_VIEW_MIN_LEVEL] = "TEXTURE_VIEW_MIN_LEVEL"
	enumLookup[FOG_COORDINATE] = "FOG_COORDINATE"
	enumLookup[VERTEX_WEIGHT_ARRAY_TYPE_EXT] = "VERTEX_WEIGHT_ARRAY_TYPE_EXT"
	enumLookup[DEPTH_STENCIL_MESA] = "DEPTH_STENCIL_MESA"
	enumLookup[OP_RECIP_EXT] = "OP_RECIP_EXT"
	enumLookup[VIDEO_COLOR_CONVERSION_MATRIX_NV] = "VIDEO_COLOR_CONVERSION_MATRIX_NV"
	enumLookup[RG16] = "RG16"
	enumLookup[SAMPLER_2D_ARRAY_SHADOW_EXT] = "SAMPLER_2D_ARRAY_SHADOW_EXT"
	enumLookup[PREFER_DOUBLEBUFFER_HINT_PGI] = "PREFER_DOUBLEBUFFER_HINT_PGI"
	enumLookup[INVERTED_SCREEN_W_REND] = "INVERTED_SCREEN_W_REND"
	enumLookup[DRAW_BUFFER10_NV] = "DRAW_BUFFER10_NV"
	enumLookup[MODULATE_COLOR_IMG] = "MODULATE_COLOR_IMG"
	enumLookup[SIGNALED] = "SIGNALED"
	enumLookup[BUFFER_BINDING] = "BUFFER_BINDING"
	enumLookup[SPARE0_NV] = "SPARE0_NV"
	enumLookup[MATRIX9_ARB] = "MATRIX9_ARB"
	enumLookup[MAX_ELEMENT_INDEX] = "MAX_ELEMENT_INDEX"
	enumLookup[PROGRAM_ADDRESS_REGISTERS_ARB] = "PROGRAM_ADDRESS_REGISTERS_ARB"
	enumLookup[READ_ONLY_ARB] = "READ_ONLY_ARB"
	enumLookup[QUERY_BY_REGION_NO_WAIT] = "QUERY_BY_REGION_NO_WAIT"
	enumLookup[DS_SCALE_NV] = "DS_SCALE_NV"
	enumLookup[TRANSFORM_FEEDBACK_VARYINGS_NV] = "TRANSFORM_FEEDBACK_VARYINGS_NV"
	enumLookup[LINES] = "LINES"
	enumLookup[DEPTH_RANGE] = "DEPTH_RANGE"
	enumLookup[SPOT_CUTOFF] = "SPOT_CUTOFF"
	enumLookup[MAX_FRAGMENT_LIGHTS_SGIX] = "MAX_FRAGMENT_LIGHTS_SGIX"
	enumLookup[LOGIC_OP_MODE] = "LOGIC_OP_MODE"
	enumLookup[ACCUM_BLUE_BITS] = "ACCUM_BLUE_BITS"
	enumLookup[POST_CONVOLUTION_ALPHA_SCALE_EXT] = "POST_CONVOLUTION_ALPHA_SCALE_EXT"
	enumLookup[UNIFORM_BUFFER_OFFSET_ALIGNMENT] = "UNIFORM_BUFFER_OFFSET_ALIGNMENT"
	enumLookup[INDEX_ARRAY_COUNT_EXT] = "INDEX_ARRAY_COUNT_EXT"
	enumLookup[MODELVIEW14_ARB] = "MODELVIEW14_ARB"
	enumLookup[VERTEX_STREAM7_ATI] = "VERTEX_STREAM7_ATI"
	enumLookup[INT8_VEC3_NV] = "INT8_VEC3_NV"
	enumLookup[ALPHA16_EXT] = "ALPHA16_EXT"
	enumLookup[DEPENDENT_RGB_TEXTURE_3D_NV] = "DEPENDENT_RGB_TEXTURE_3D_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_INTEGER_NV] = "VERTEX_ATTRIB_ARRAY_INTEGER_NV"
	enumLookup[SAMPLER_BUFFER] = "SAMPLER_BUFFER"
	enumLookup[UNSIGNED_INT_IMAGE_2D_ARRAY] = "UNSIGNED_INT_IMAGE_2D_ARRAY"
	enumLookup[FONT_Y_MAX_BOUNDS_BIT_NV] = "FONT_Y_MAX_BOUNDS_BIT_NV"
	enumLookup[VIEW_CLASS_S3TC_DXT3_RGBA] = "VIEW_CLASS_S3TC_DXT3_RGBA"
	enumLookup[VERTEX_SHADER_BINDING_EXT] = "VERTEX_SHADER_BINDING_EXT"
	enumLookup[MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS] = "MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS"
	enumLookup[INT_IMAGE_2D_ARRAY_EXT] = "INT_IMAGE_2D_ARRAY_EXT"
	enumLookup[REPLACE_MIDDLE_SUN] = "REPLACE_MIDDLE_SUN"
	enumLookup[PATH_FILL_COVER_MODE_NV] = "PATH_FILL_COVER_MODE_NV"
	enumLookup[SYNC_FLUSH_COMMANDS_BIT] = "SYNC_FLUSH_COMMANDS_BIT"
	enumLookup[CLIP_PLANE5] = "CLIP_PLANE5"
	enumLookup[COLOR_TABLE_BIAS] = "COLOR_TABLE_BIAS"
	enumLookup[ATOMIC_COUNTER_BUFFER_DATA_SIZE] = "ATOMIC_COUNTER_BUFFER_DATA_SIZE"
	enumLookup[COLOR_ARRAY_LIST_STRIDE_IBM] = "COLOR_ARRAY_LIST_STRIDE_IBM"
	enumLookup[ONE_MINUS_DST_COLOR] = "ONE_MINUS_DST_COLOR"
	enumLookup[LUMINANCE6_ALPHA2] = "LUMINANCE6_ALPHA2"
	enumLookup[CULL_MODES_NV] = "CULL_MODES_NV"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_SIZE_NV] = "TRANSFORM_FEEDBACK_BUFFER_SIZE_NV"
	enumLookup[IMAGE_1D_EXT] = "IMAGE_1D_EXT"
	enumLookup[POST_CONVOLUTION_COLOR_TABLE] = "POST_CONVOLUTION_COLOR_TABLE"
	enumLookup[REPLACEMENT_CODE_ARRAY_STRIDE_SUN] = "REPLACEMENT_CODE_ARRAY_STRIDE_SUN"
	enumLookup[PN_TRIANGLES_POINT_MODE_CUBIC_ATI] = "PN_TRIANGLES_POINT_MODE_CUBIC_ATI"
	enumLookup[FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA] = "FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA"
	enumLookup[TEXTURE_DEPTH_TYPE_ARB] = "TEXTURE_DEPTH_TYPE_ARB"
	enumLookup[IMAGE_BINDING_LAYER] = "IMAGE_BINDING_LAYER"
	enumLookup[LUMINANCE8_ALPHA8_EXT] = "LUMINANCE8_ALPHA8_EXT"
	enumLookup[MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS] = "MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS"
	enumLookup[DRAW_BUFFER1_ARB] = "DRAW_BUFFER1_ARB"
	enumLookup[FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT] = "FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT"
	enumLookup[IMAGE_2D_MULTISAMPLE_ARRAY_EXT] = "IMAGE_2D_MULTISAMPLE_ARRAY_EXT"
	enumLookup[ALPHA_SCALE] = "ALPHA_SCALE"
	enumLookup[INTERNALFORMAT_GREEN_SIZE] = "INTERNALFORMAT_GREEN_SIZE"
	enumLookup[PIXEL_MIN_FILTER_EXT] = "PIXEL_MIN_FILTER_EXT"
	enumLookup[INTERPOLATE_ARB] = "INTERPOLATE_ARB"
	enumLookup[GL_4PASS_1_EXT] = "GL_4PASS_1_EXT"
	enumLookup[BLEND_EQUATION_ALPHA] = "BLEND_EQUATION_ALPHA"
	enumLookup[FLOAT_MAT2_ARB] = "FLOAT_MAT2_ARB"
	enumLookup[MAX_GEOMETRY_ATOMIC_COUNTERS] = "MAX_GEOMETRY_ATOMIC_COUNTERS"
	enumLookup[SGIX_subsample] = "SGIX_subsample"
	enumLookup[MAX_3D_TEXTURE_SIZE_OES] = "MAX_3D_TEXTURE_SIZE_OES"
	enumLookup[FLOAT_RGBA16_NV] = "FLOAT_RGBA16_NV"
	enumLookup[TEXTURE_USAGE_ANGLE] = "TEXTURE_USAGE_ANGLE"
	enumLookup[TEXTURE_1D_STACK_BINDING_MESAX] = "TEXTURE_1D_STACK_BINDING_MESAX"
	enumLookup[ALPHA8UI_EXT] = "ALPHA8UI_EXT"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_ACTIVE] = "TRANSFORM_FEEDBACK_BUFFER_ACTIVE"
	enumLookup[EXT_packed_pixels] = "EXT_packed_pixels"
	enumLookup[SOURCE3_ALPHA_NV] = "SOURCE3_ALPHA_NV"
	enumLookup[SIGNED_RGB8_UNSIGNED_ALPHA8_NV] = "SIGNED_RGB8_UNSIGNED_ALPHA8_NV"
	enumLookup[ACCUM_GREEN_BITS] = "ACCUM_GREEN_BITS"
	enumLookup[NAME_STACK_DEPTH] = "NAME_STACK_DEPTH"
	enumLookup[NOR] = "NOR"
	enumLookup[UNSIGNED_SHORT_5_6_5_EXT] = "UNSIGNED_SHORT_5_6_5_EXT"
	enumLookup[COMBINER3_NV] = "COMBINER3_NV"
	enumLookup[BOOL_VEC4] = "BOOL_VEC4"
	enumLookup[INNOCENT_CONTEXT_RESET_ARB] = "INNOCENT_CONTEXT_RESET_ARB"
	enumLookup[SIGNED_NEGATE_NV] = "SIGNED_NEGATE_NV"
	enumLookup[VERTEX_ARRAY_RANGE_POINTER_APPLE] = "VERTEX_ARRAY_RANGE_POINTER_APPLE"
	enumLookup[DOT3_RGBA_EXT] = "DOT3_RGBA_EXT"
	enumLookup[MATRIX25_ARB] = "MATRIX25_ARB"
	enumLookup[TEXTURE_GEN_T] = "TEXTURE_GEN_T"
	enumLookup[HISTOGRAM_RED_SIZE_EXT] = "HISTOGRAM_RED_SIZE_EXT"
	enumLookup[FRAGMENT_LIGHT7_SGIX] = "FRAGMENT_LIGHT7_SGIX"
	enumLookup[VERTEX_ATTRIB_MAP1_DOMAIN_APPLE] = "VERTEX_ATTRIB_MAP1_DOMAIN_APPLE"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES"
	enumLookup[POLYGON_MODE] = "POLYGON_MODE"
	enumLookup[INTENSITY32F_ARB] = "INTENSITY32F_ARB"
	enumLookup[DRAW_BUFFER0_ATI] = "DRAW_BUFFER0_ATI"
	enumLookup[CON_14_ATI] = "CON_14_ATI"
	enumLookup[FRONT] = "FRONT"
	enumLookup[GL_4D_COLOR_TEXTURE] = "GL_4D_COLOR_TEXTURE"
	enumLookup[VIEW_CLASS_S3TC_DXT5_RGBA] = "VIEW_CLASS_S3TC_DXT5_RGBA"
	enumLookup[COLOR_ATTACHMENT11_EXT] = "COLOR_ATTACHMENT11_EXT"
	enumLookup[SGIX_texture_scale_bias] = "SGIX_texture_scale_bias"
	enumLookup[PIXEL_MAP_I_TO_R] = "PIXEL_MAP_I_TO_R"
	enumLookup[STENCIL_REF] = "STENCIL_REF"
	enumLookup[COLOR_ARRAY_STRIDE] = "COLOR_ARRAY_STRIDE"
	enumLookup[TEXTURE_LOD_BIAS] = "TEXTURE_LOD_BIAS"
	enumLookup[VERTEX_SHADER_LOCALS_EXT] = "VERTEX_SHADER_LOCALS_EXT"
	enumLookup[INT_IMAGE_2D_MULTISAMPLE_EXT] = "INT_IMAGE_2D_MULTISAMPLE_EXT"
	enumLookup[PROXY_POST_COLOR_MATRIX_COLOR_TABLE] = "PROXY_POST_COLOR_MATRIX_COLOR_TABLE"
	enumLookup[CURRENT_RASTER_NORMAL_SGIX] = "CURRENT_RASTER_NORMAL_SGIX"
	enumLookup[SMALL_CW_ARC_TO_NV] = "SMALL_CW_ARC_TO_NV"
	enumLookup[REFERENCED_BY_VERTEX_SHADER] = "REFERENCED_BY_VERTEX_SHADER"
	enumLookup[TEXTURE27_ARB] = "TEXTURE27_ARB"
	enumLookup[MAX_PROGRAM_NATIVE_TEMPORARIES_ARB] = "MAX_PROGRAM_NATIVE_TEMPORARIES_ARB"
	enumLookup[QUERY_BY_REGION_WAIT_NV] = "QUERY_BY_REGION_WAIT_NV"
	enumLookup[POLYGON_OFFSET_POINT] = "POLYGON_OFFSET_POINT"
	enumLookup[CURRENT_SECONDARY_COLOR] = "CURRENT_SECONDARY_COLOR"
	enumLookup[VERTEX_ATTRIB_ARRAY3_NV] = "VERTEX_ATTRIB_ARRAY3_NV"
	enumLookup[COMPRESSED_RGB_FXT1_3DFX] = "COMPRESSED_RGB_FXT1_3DFX"
	enumLookup[DRAW_BUFFER1] = "DRAW_BUFFER1"
	enumLookup[TEXTURE_COMPARE_FUNC] = "TEXTURE_COMPARE_FUNC"
	enumLookup[MAX_GEOMETRY_IMAGE_UNIFORMS] = "MAX_GEOMETRY_IMAGE_UNIFORMS"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR"
	enumLookup[DEBUG_CALLBACK_FUNCTION] = "DEBUG_CALLBACK_FUNCTION"
	enumLookup[TRACE_PIXELS_BIT_MESA] = "TRACE_PIXELS_BIT_MESA"
	enumLookup[VERTEX_ATTRIB_MAP1_COEFF_APPLE] = "VERTEX_ATTRIB_MAP1_COEFF_APPLE"
	enumLookup[DOUBLE_MAT2x4_EXT] = "DOUBLE_MAT2x4_EXT"
	enumLookup[COLOR_EXT] = "COLOR_EXT"
	enumLookup[BLEND_EQUATION_RGB] = "BLEND_EQUATION_RGB"
	enumLookup[POST_CONVOLUTION_BLUE_SCALE] = "POST_CONVOLUTION_BLUE_SCALE"
	enumLookup[UNSIGNED_SHORT_4_4_4_4_EXT] = "UNSIGNED_SHORT_4_4_4_4_EXT"
	enumLookup[GL_4PASS_1_SGIS] = "GL_4PASS_1_SGIS"
	enumLookup[TEXTURE_BASE_LEVEL] = "TEXTURE_BASE_LEVEL"
	enumLookup[TEXTURE_COMPRESSED_BLOCK_HEIGHT] = "TEXTURE_COMPRESSED_BLOCK_HEIGHT"
	enumLookup[UNSIGNED_INVERT_NV] = "UNSIGNED_INVERT_NV"
	enumLookup[DSDT_MAG_NV] = "DSDT_MAG_NV"
	enumLookup[COLOR_ATTACHMENT1] = "COLOR_ATTACHMENT1"
	enumLookup[GLYPH_HORIZONTAL_BEARING_X_BIT_NV] = "GLYPH_HORIZONTAL_BEARING_X_BIT_NV"
	enumLookup[MODELVIEW11_ARB] = "MODELVIEW11_ARB"
	enumLookup[FLOAT_MAT4] = "FLOAT_MAT4"
	enumLookup[VERTEX_ARRAY_LENGTH_NV] = "VERTEX_ARRAY_LENGTH_NV"
	enumLookup[DEBUG_SOURCE_OTHER_ARB] = "DEBUG_SOURCE_OTHER_ARB"
	enumLookup[PROGRAM_PARAMETERS_ARB] = "PROGRAM_PARAMETERS_ARB"
	enumLookup[GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV] = "GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV"
	enumLookup[COLOR_TABLE_SCALE] = "COLOR_TABLE_SCALE"
	enumLookup[EVAL_2D_NV] = "EVAL_2D_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB] = "VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[STREAM_READ] = "STREAM_READ"
	enumLookup[AFFINE_2D_NV] = "AFFINE_2D_NV"
	enumLookup[ARRAY_SIZE] = "ARRAY_SIZE"
	enumLookup[DRAW_BUFFER] = "DRAW_BUFFER"
	enumLookup[TEXTURE_ALPHA_SIZE] = "TEXTURE_ALPHA_SIZE"
	enumLookup[HISTOGRAM_FORMAT] = "HISTOGRAM_FORMAT"
	enumLookup[COLOR_SUM_ARB] = "COLOR_SUM_ARB"
	enumLookup[SAMPLER_2D_ARRAY_SHADOW_NV] = "SAMPLER_2D_ARRAY_SHADOW_NV"
	enumLookup[TEXTURE_CONSTANT_DATA_SUNX] = "TEXTURE_CONSTANT_DATA_SUNX"
	enumLookup[TANGENT_ARRAY_EXT] = "TANGENT_ARRAY_EXT"
	enumLookup[TEXTURE_GREEN_TYPE] = "TEXTURE_GREEN_TYPE"
	enumLookup[T2F_IUI_N3F_V2F_EXT] = "T2F_IUI_N3F_V2F_EXT"
	enumLookup[SRC2_RGB] = "SRC2_RGB"
	enumLookup[MAP_TESSELLATION_NV] = "MAP_TESSELLATION_NV"
	enumLookup[MAX_ARRAY_TEXTURE_LAYERS_EXT] = "MAX_ARRAY_TEXTURE_LAYERS_EXT"
	enumLookup[FLOAT_MAT4x3] = "FLOAT_MAT4x3"
	enumLookup[RENDERBUFFER_COLOR_SAMPLES_NV] = "RENDERBUFFER_COLOR_SAMPLES_NV"
	enumLookup[FONT_UNDERLINE_THICKNESS_BIT_NV] = "FONT_UNDERLINE_THICKNESS_BIT_NV"
	enumLookup[FRONT_RIGHT] = "FRONT_RIGHT"
	enumLookup[EMBOSS_LIGHT_NV] = "EMBOSS_LIGHT_NV"
	enumLookup[MAX_TRACK_MATRIX_STACK_DEPTH_NV] = "MAX_TRACK_MATRIX_STACK_DEPTH_NV"
	enumLookup[MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV] = "MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV"
	enumLookup[RENDERBUFFER_ALPHA_SIZE] = "RENDERBUFFER_ALPHA_SIZE"
	enumLookup[DEBUG_SEVERITY_MEDIUM] = "DEBUG_SEVERITY_MEDIUM"
	enumLookup[TEXTURE22_ARB] = "TEXTURE22_ARB"
	enumLookup[UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY] = "UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY"
	enumLookup[PATH_STENCIL_FUNC_NV] = "PATH_STENCIL_FUNC_NV"
	enumLookup[COMBINER_SUM_OUTPUT_NV] = "COMBINER_SUM_OUTPUT_NV"
	enumLookup[WEIGHT_ARRAY_BUFFER_BINDING] = "WEIGHT_ARRAY_BUFFER_BINDING"
	enumLookup[CON_12_ATI] = "CON_12_ATI"
	enumLookup[PERFORMANCE_MONITOR_AMD] = "PERFORMANCE_MONITOR_AMD"
	enumLookup[SGIX_tag_sample_buffer] = "SGIX_tag_sample_buffer"
	enumLookup[VERTEX_PROGRAM_NV] = "VERTEX_PROGRAM_NV"
	enumLookup[WEIGHT_ARRAY_BUFFER_BINDING_ARB] = "WEIGHT_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[DRAW_FRAMEBUFFER_NV] = "DRAW_FRAMEBUFFER_NV"
	enumLookup[COMPRESSED_RGBA_ASTC_6x5_KHR] = "COMPRESSED_RGBA_ASTC_6x5_KHR"
	enumLookup[FUNC_SUBTRACT_OES] = "FUNC_SUBTRACT_OES"
	enumLookup[MAX_4D_TEXTURE_SIZE_SGIS] = "MAX_4D_TEXTURE_SIZE_SGIS"
	enumLookup[FOG_OFFSET_VALUE_SGIX] = "FOG_OFFSET_VALUE_SGIX"
	enumLookup[OP_RECIP_SQRT_EXT] = "OP_RECIP_SQRT_EXT"
	enumLookup[MAX_PROGRAM_IF_DEPTH_NV] = "MAX_PROGRAM_IF_DEPTH_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT] = "FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT"
	enumLookup[LUMINANCE12_ALPHA12] = "LUMINANCE12_ALPHA12"
	enumLookup[PROGRAM_STRING_NV] = "PROGRAM_STRING_NV"
	enumLookup[DRAW_BUFFER6] = "DRAW_BUFFER6"
	enumLookup[SAMPLE_MASK_VALUE_NV] = "SAMPLE_MASK_VALUE_NV"
	enumLookup[HISTOGRAM_WIDTH] = "HISTOGRAM_WIDTH"
	enumLookup[CON_8_ATI] = "CON_8_ATI"
	enumLookup[DRAW_FRAMEBUFFER_BINDING_NV] = "DRAW_FRAMEBUFFER_BINDING_NV"
	enumLookup[PIXEL_TEXTURE_SGIS] = "PIXEL_TEXTURE_SGIS"
	enumLookup[HILO_NV] = "HILO_NV"
	enumLookup[DEPTH_STENCIL_TO_BGRA_NV] = "DEPTH_STENCIL_TO_BGRA_NV"
	enumLookup[LUMINANCE_ALPHA32UI_EXT] = "LUMINANCE_ALPHA32UI_EXT"
	enumLookup[ARC_TO_NV] = "ARC_TO_NV"
	enumLookup[LINE_STRIP_ADJACENCY_EXT] = "LINE_STRIP_ADJACENCY_EXT"
	enumLookup[CLIP_DISTANCE7] = "CLIP_DISTANCE7"
	enumLookup[NUM_COMPRESSED_TEXTURE_FORMATS] = "NUM_COMPRESSED_TEXTURE_FORMATS"
	enumLookup[TEXTURE_CUBE_MAP_EXT] = "TEXTURE_CUBE_MAP_EXT"
	enumLookup[SOURCE0_ALPHA_EXT] = "SOURCE0_ALPHA_EXT"
	enumLookup[FLOAT_CLEAR_COLOR_VALUE_NV] = "FLOAT_CLEAR_COLOR_VALUE_NV"
	enumLookup[CONDITION_SATISFIED_APPLE] = "CONDITION_SATISFIED_APPLE"
	enumLookup[DEBUG_SEVERITY_HIGH_AMD] = "DEBUG_SEVERITY_HIGH_AMD"
	enumLookup[DRAW_BUFFER9_NV] = "DRAW_BUFFER9_NV"
	enumLookup[PIXEL_TEX_GEN_MODE_SGIX] = "PIXEL_TEX_GEN_MODE_SGIX"
	enumLookup[MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT] = "MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT"
	enumLookup[PIXEL_PACK_BUFFER_EXT] = "PIXEL_PACK_BUFFER_EXT"
	enumLookup[DUAL_INTENSITY12_SGIS] = "DUAL_INTENSITY12_SGIS"
	enumLookup[IMAGE_CLASS_2_X_8] = "IMAGE_CLASS_2_X_8"
	enumLookup[TEXTURE16] = "TEXTURE16"
	enumLookup[STORAGE_PRIVATE_APPLE] = "STORAGE_PRIVATE_APPLE"
	enumLookup[GL_3DC_XY_AMD] = "GL_3DC_XY_AMD"
	enumLookup[DEPTH_BUFFER_FLOAT_MODE_NV] = "DEPTH_BUFFER_FLOAT_MODE_NV"
	enumLookup[TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM] = "TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM"
	enumLookup[ALPHA12_EXT] = "ALPHA12_EXT"
	enumLookup[EVAL_VERTEX_ATTRIB3_NV] = "EVAL_VERTEX_ATTRIB3_NV"
	enumLookup[MIRROR_CLAMP_TO_EDGE_EXT] = "MIRROR_CLAMP_TO_EDGE_EXT"
	enumLookup[OUTPUT_TEXTURE_COORD7_EXT] = "OUTPUT_TEXTURE_COORD7_EXT"
	enumLookup[POST_COLOR_MATRIX_ALPHA_BIAS_SGI] = "POST_COLOR_MATRIX_ALPHA_BIAS_SGI"
	enumLookup[DUAL_INTENSITY16_SGIS] = "DUAL_INTENSITY16_SGIS"
	enumLookup[COMBINER7_NV] = "COMBINER7_NV"
	enumLookup[QUERY_RESULT_AVAILABLE_EXT] = "QUERY_RESULT_AVAILABLE_EXT"
	enumLookup[IMAGE_CUBE] = "IMAGE_CUBE"
	enumLookup[LIGHT2] = "LIGHT2"
	enumLookup[EDGE_FLAG_ARRAY_BUFFER_BINDING] = "EDGE_FLAG_ARRAY_BUFFER_BINDING"
	enumLookup[PERCENTAGE_AMD] = "PERCENTAGE_AMD"
	enumLookup[PIXEL_TILE_GRID_HEIGHT_SGIX] = "PIXEL_TILE_GRID_HEIGHT_SGIX"
	enumLookup[PALETTE8_RGB5_A1_OES] = "PALETTE8_RGB5_A1_OES"
	enumLookup[TEXTURE_BINDING_EXTERNAL_OES] = "TEXTURE_BINDING_EXTERNAL_OES"
	enumLookup[COLOR_INDEX2_EXT] = "COLOR_INDEX2_EXT"
	enumLookup[DOUBLEBUFFER] = "DOUBLEBUFFER"
	enumLookup[FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS] = "FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS"
	enumLookup[LOW_INT] = "LOW_INT"
	enumLookup[PIXEL_TEX_GEN_Q_FLOOR_SGIX] = "PIXEL_TEX_GEN_Q_FLOOR_SGIX"
	enumLookup[TEXTURE_MATERIAL_FACE_EXT] = "TEXTURE_MATERIAL_FACE_EXT"
	enumLookup[MAX_PROGRAM_ENV_PARAMETERS_ARB] = "MAX_PROGRAM_ENV_PARAMETERS_ARB"
	enumLookup[RENDERBUFFER_INTERNAL_FORMAT_OES] = "RENDERBUFFER_INTERNAL_FORMAT_OES"
	enumLookup[LUMINANCE_ALPHA16I_EXT] = "LUMINANCE_ALPHA16I_EXT"
	enumLookup[INDEX_ARRAY] = "INDEX_ARRAY"
	enumLookup[EXTENSIONS] = "EXTENSIONS"
	enumLookup[POST_TEXTURE_FILTER_SCALE_SGIX] = "POST_TEXTURE_FILTER_SCALE_SGIX"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT"
	enumLookup[GEOMETRY_OUTPUT_TYPE_ARB] = "GEOMETRY_OUTPUT_TYPE_ARB"
	enumLookup[TEXTURE_BUFFER_OFFSET_ALIGNMENT] = "TEXTURE_BUFFER_OFFSET_ALIGNMENT"
	enumLookup[QUERY] = "QUERY"
	enumLookup[TRACE_NAME_MESA] = "TRACE_NAME_MESA"
	enumLookup[SRGB8] = "SRGB8"
	enumLookup[INT_SAMPLER_2D_MULTISAMPLE_ARRAY] = "INT_SAMPLER_2D_MULTISAMPLE_ARRAY"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR"
	enumLookup[MAX_CONVOLUTION_HEIGHT_EXT] = "MAX_CONVOLUTION_HEIGHT_EXT"
	enumLookup[VIEW_CLASS_BPTC_UNORM] = "VIEW_CLASS_BPTC_UNORM"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT] = "TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT"
	enumLookup[COMPRESSED_SRGB_EXT] = "COMPRESSED_SRGB_EXT"
	enumLookup[POST_CONVOLUTION_RED_SCALE] = "POST_CONVOLUTION_RED_SCALE"
	enumLookup[GL_4PASS_0_SGIS] = "GL_4PASS_0_SGIS"
	enumLookup[TEXTURE3_ARB] = "TEXTURE3_ARB"
	enumLookup[STORAGE_CLIENT_APPLE] = "STORAGE_CLIENT_APPLE"
	enumLookup[MATRIX_PALETTE_ARB] = "MATRIX_PALETTE_ARB"
	enumLookup[RENDERBUFFER_BINDING] = "RENDERBUFFER_BINDING"
	enumLookup[RGB_INTEGER] = "RGB_INTEGER"
	enumLookup[TEXTURE_BORDER_COLOR] = "TEXTURE_BORDER_COLOR"
	enumLookup[COPY_INVERTED] = "COPY_INVERTED"
	enumLookup[ABGR_EXT] = "ABGR_EXT"
	enumLookup[COLOR_COMPONENTS] = "COLOR_COMPONENTS"
	enumLookup[TEXTURE_VIEW_NUM_LAYERS] = "TEXTURE_VIEW_NUM_LAYERS"
	enumLookup[SIGNED_IDENTITY_NV] = "SIGNED_IDENTITY_NV"
	enumLookup[VERTEX_BLEND_ARB] = "VERTEX_BLEND_ARB"
	enumLookup[CON_3_ATI] = "CON_3_ATI"
	enumLookup[STENCIL_INDEX16] = "STENCIL_INDEX16"
	enumLookup[COLOR_ATTACHMENT15_NV] = "COLOR_ATTACHMENT15_NV"
	enumLookup[TEXTURE_COLOR_SAMPLES_NV] = "TEXTURE_COLOR_SAMPLES_NV"
	enumLookup[INTENSITY12] = "INTENSITY12"
	enumLookup[VERTEX_ARRAY_RANGE_VALID_NV] = "VERTEX_ARRAY_RANGE_VALID_NV"
	enumLookup[REG_22_ATI] = "REG_22_ATI"
	enumLookup[REG_26_ATI] = "REG_26_ATI"
	enumLookup[CURRENT_VERTEX_WEIGHT_EXT] = "CURRENT_VERTEX_WEIGHT_EXT"
	enumLookup[PATH_GEN_COLOR_FORMAT_NV] = "PATH_GEN_COLOR_FORMAT_NV"
	enumLookup[FOG_FUNC_SGIS] = "FOG_FUNC_SGIS"
	enumLookup[OUTPUT_COLOR1_EXT] = "OUTPUT_COLOR1_EXT"
	enumLookup[MOVE_TO_NV] = "MOVE_TO_NV"
	enumLookup[POLYGON_OFFSET_EXT] = "POLYGON_OFFSET_EXT"
	enumLookup[MODELVIEW31_ARB] = "MODELVIEW31_ARB"
	enumLookup[GL_4X_BIT_ATI] = "GL_4X_BIT_ATI"
	enumLookup[POINT_SIZE_MIN] = "POINT_SIZE_MIN"
	enumLookup[PROJECTION] = "PROJECTION"
	enumLookup[TEXTURE_APPLICATION_MODE_EXT] = "TEXTURE_APPLICATION_MODE_EXT"
	enumLookup[DEBUG_SEVERITY_HIGH_ARB] = "DEBUG_SEVERITY_HIGH_ARB"
	enumLookup[POLYGON_OFFSET_FACTOR_EXT] = "POLYGON_OFFSET_FACTOR_EXT"
	enumLookup[TEXTURE26_ARB] = "TEXTURE26_ARB"
	enumLookup[VERTEX_ATTRIB_ARRAY_INTEGER] = "VERTEX_ATTRIB_ARRAY_INTEGER"
	enumLookup[TRANSFORM_FEEDBACK] = "TRANSFORM_FEEDBACK"
	enumLookup[INDEX_ARRAY_LENGTH_NV] = "INDEX_ARRAY_LENGTH_NV"
	enumLookup[DEPTH_BUFFER_BIT] = "DEPTH_BUFFER_BIT"
	enumLookup[FRAMEZOOM_FACTOR_SGIX] = "FRAMEZOOM_FACTOR_SGIX"
	enumLookup[COMPRESSED_ALPHA_ARB] = "COMPRESSED_ALPHA_ARB"
	enumLookup[SAMPLER_1D_ARRAY_SHADOW_EXT] = "SAMPLER_1D_ARRAY_SHADOW_EXT"
	enumLookup[MODELVIEW8_ARB] = "MODELVIEW8_ARB"
	enumLookup[MAX_VERTEX_STREAMS] = "MAX_VERTEX_STREAMS"
	enumLookup[R8_SNORM] = "R8_SNORM"
	enumLookup[PRIMITIVE_RESTART] = "PRIMITIVE_RESTART"
	enumLookup[LUMINANCE16UI_EXT] = "LUMINANCE16UI_EXT"
	enumLookup[TIMEOUT_IGNORED_APPLE] = "TIMEOUT_IGNORED_APPLE"
	enumLookup[ALL_STATIC_DATA_IBM] = "ALL_STATIC_DATA_IBM"
	enumLookup[SAMPLER_BINDING] = "SAMPLER_BINDING"
	enumLookup[LUMINANCE_INTEGER_EXT] = "LUMINANCE_INTEGER_EXT"
	enumLookup[ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER] = "ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER"
	enumLookup[ELEMENT_ARRAY_BARRIER_BIT] = "ELEMENT_ARRAY_BARRIER_BIT"
	enumLookup[RED_SCALE] = "RED_SCALE"
	enumLookup[DEBUG_TYPE_PUSH_GROUP] = "DEBUG_TYPE_PUSH_GROUP"
	enumLookup[OPERAND1_RGB] = "OPERAND1_RGB"
	enumLookup[IMPLEMENTATION_COLOR_READ_TYPE] = "IMPLEMENTATION_COLOR_READ_TYPE"
	enumLookup[TESS_EVALUATION_SHADER] = "TESS_EVALUATION_SHADER"
	enumLookup[EXT_blend_minmax] = "EXT_blend_minmax"
	enumLookup[TEXTURE_GREEN_SIZE] = "TEXTURE_GREEN_SIZE"
	enumLookup[COMPRESSED_TEXTURE_FORMATS] = "COMPRESSED_TEXTURE_FORMATS"
	enumLookup[LO_SCALE_NV] = "LO_SCALE_NV"
	enumLookup[MODELVIEW0_MATRIX_EXT] = "MODELVIEW0_MATRIX_EXT"
	enumLookup[STENCIL] = "STENCIL"
	enumLookup[SKIP_COMPONENTS4_NV] = "SKIP_COMPONENTS4_NV"
	enumLookup[UNSIGNED_INT_VEC3] = "UNSIGNED_INT_VEC3"
	enumLookup[UNSIGNED_INT8_VEC2_NV] = "UNSIGNED_INT8_VEC2_NV"
	enumLookup[IS_ROW_MAJOR] = "IS_ROW_MAJOR"
	enumLookup[SAMPLE_PATTERN_EXT] = "SAMPLE_PATTERN_EXT"
	enumLookup[INTENSITY16F_ARB] = "INTENSITY16F_ARB"
	enumLookup[DRAW_BUFFER12_ARB] = "DRAW_BUFFER12_ARB"
	enumLookup[STENCIL_BACK_VALUE_MASK] = "STENCIL_BACK_VALUE_MASK"
	enumLookup[MAX_COMBINED_ATOMIC_COUNTERS] = "MAX_COMBINED_ATOMIC_COUNTERS"
	enumLookup[DOT_PRODUCT_TEXTURE_2D_NV] = "DOT_PRODUCT_TEXTURE_2D_NV"
	enumLookup[OBJECT_INFO_LOG_LENGTH_ARB] = "OBJECT_INFO_LOG_LENGTH_ARB"
	enumLookup[PALETTE8_RGBA8_OES] = "PALETTE8_RGBA8_OES"
	enumLookup[TRANSFORM_FEEDBACK_VARYINGS_EXT] = "TRANSFORM_FEEDBACK_VARYINGS_EXT"
	enumLookup[DRAW_FRAMEBUFFER] = "DRAW_FRAMEBUFFER"
	enumLookup[SHADER_STORAGE_BUFFER_SIZE] = "SHADER_STORAGE_BUFFER_SIZE"
	enumLookup[SUBPIXEL_BITS] = "SUBPIXEL_BITS"
	enumLookup[OFFSET_TEXTURE_2D_MATRIX_NV] = "OFFSET_TEXTURE_2D_MATRIX_NV"
	enumLookup[ELEMENT_ARRAY_POINTER_ATI] = "ELEMENT_ARRAY_POINTER_ATI"
	enumLookup[PRIMITIVES_GENERATED] = "PRIMITIVES_GENERATED"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT"
	enumLookup[IMAGE_FORMAT_COMPATIBILITY_BY_CLASS] = "IMAGE_FORMAT_COMPATIBILITY_BY_CLASS"
	enumLookup[INCR] = "INCR"
	enumLookup[MAX_TEXTURE_MAX_ANISOTROPY_EXT] = "MAX_TEXTURE_MAX_ANISOTROPY_EXT"
	enumLookup[UNSIGNED_SHORT_8_8_MESA] = "UNSIGNED_SHORT_8_8_MESA"
	enumLookup[OP_SET_GE_EXT] = "OP_SET_GE_EXT"
	enumLookup[FRAGMENT_SHADER_ATI] = "FRAGMENT_SHADER_ATI"
	enumLookup[RENDERBUFFER_DEPTH_SIZE_EXT] = "RENDERBUFFER_DEPTH_SIZE_EXT"
	enumLookup[LAYOUT_DEFAULT_INTEL] = "LAYOUT_DEFAULT_INTEL"
	enumLookup[ORDER] = "ORDER"
	enumLookup[FRAGMENT_PROGRAM_CALLBACK_DATA_MESA] = "FRAGMENT_PROGRAM_CALLBACK_DATA_MESA"
	enumLookup[LOCATION_INDEX] = "LOCATION_INDEX"
	enumLookup[TEXTURE_DEPTH] = "TEXTURE_DEPTH"
	enumLookup[POINT_SIZE_MAX_EXT] = "POINT_SIZE_MAX_EXT"
	enumLookup[TEXTURE_CUBE_MAP] = "TEXTURE_CUBE_MAP"
	enumLookup[R1UI_C4F_N3F_V3F_SUN] = "R1UI_C4F_N3F_V3F_SUN"
	enumLookup[MAP2_VERTEX_ATTRIB10_4_NV] = "MAP2_VERTEX_ATTRIB10_4_NV"
	enumLookup[DRAW_BUFFER7] = "DRAW_BUFFER7"
	enumLookup[ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH] = "ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH"
	enumLookup[RGB8UI_EXT] = "RGB8UI_EXT"
	enumLookup[RGBA8I] = "RGBA8I"
	enumLookup[SGI_texture_color_table] = "SGI_texture_color_table"
	enumLookup[LINE_WIDTH] = "LINE_WIDTH"
	enumLookup[VIEW_CLASS_48_BITS] = "VIEW_CLASS_48_BITS"
	enumLookup[ACTIVE_TEXTURE] = "ACTIVE_TEXTURE"
	enumLookup[PROGRAM_TARGET_NV] = "PROGRAM_TARGET_NV"
	enumLookup[WEIGHT_ARRAY_SIZE_ARB] = "WEIGHT_ARRAY_SIZE_ARB"
	enumLookup[DOT3_RGB_EXT] = "DOT3_RGB_EXT"
	enumLookup[MATRIX_INDEX_ARRAY_TYPE_ARB] = "MATRIX_INDEX_ARRAY_TYPE_ARB"
	enumLookup[COMPARE_R_TO_TEXTURE_ARB] = "COMPARE_R_TO_TEXTURE_ARB"
	enumLookup[RGBA32UI] = "RGBA32UI"
	enumLookup[SYNC_GPU_COMMANDS_COMPLETE_APPLE] = "SYNC_GPU_COMMANDS_COMPLETE_APPLE"
	enumLookup[FALSE] = "FALSE"
	enumLookup[VERTEX_ARRAY_BINDING] = "VERTEX_ARRAY_BINDING"
	enumLookup[MAP_ATTRIB_V_ORDER_NV] = "MAP_ATTRIB_V_ORDER_NV"
	enumLookup[MAX_PROGRAM_TEMPORARIES_ARB] = "MAX_PROGRAM_TEMPORARIES_ARB"
	enumLookup[COPY_READ_BUFFER_BINDING] = "COPY_READ_BUFFER_BINDING"
	enumLookup[ALL_BARRIER_BITS_EXT] = "ALL_BARRIER_BITS_EXT"
	enumLookup[CONVOLUTION_FILTER_BIAS] = "CONVOLUTION_FILTER_BIAS"
	enumLookup[DEBUG_TYPE_ERROR] = "DEBUG_TYPE_ERROR"
	enumLookup[OUTPUT_TEXTURE_COORD17_EXT] = "OUTPUT_TEXTURE_COORD17_EXT"
	enumLookup[RENDERBUFFER_EXT] = "RENDERBUFFER_EXT"
	enumLookup[DEBUG_CATEGORY_PERFORMANCE_AMD] = "DEBUG_CATEGORY_PERFORMANCE_AMD"
	enumLookup[SGIX_depth_texture] = "SGIX_depth_texture"
	enumLookup[ADD_SIGNED_EXT] = "ADD_SIGNED_EXT"
	enumLookup[BUFFER_MAPPED_ARB] = "BUFFER_MAPPED_ARB"
	enumLookup[PACK_RESAMPLE_OML] = "PACK_RESAMPLE_OML"
	enumLookup[COMPILE_STATUS] = "COMPILE_STATUS"
	enumLookup[SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM] = "SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM"
	enumLookup[TESS_EVALUATION_TEXTURE] = "TESS_EVALUATION_TEXTURE"
	enumLookup[RGB16UI_EXT] = "RGB16UI_EXT"
	enumLookup[DSDT8_MAG8_INTENSITY8_NV] = "DSDT8_MAG8_INTENSITY8_NV"
	enumLookup[CON_5_ATI] = "CON_5_ATI"
	enumLookup[UNSIGNED_INT_IMAGE_1D_EXT] = "UNSIGNED_INT_IMAGE_1D_EXT"
	enumLookup[IMAGE_2D_MULTISAMPLE] = "IMAGE_2D_MULTISAMPLE"
	enumLookup[FOG_COORDINATE_SOURCE_EXT] = "FOG_COORDINATE_SOURCE_EXT"
	enumLookup[SLUMINANCE8] = "SLUMINANCE8"
	enumLookup[UNSIGNED_INT16_NV] = "UNSIGNED_INT16_NV"
	enumLookup[PATH_STENCIL_REF_NV] = "PATH_STENCIL_REF_NV"
	enumLookup[PALETTE4_RGBA4_OES] = "PALETTE4_RGBA4_OES"
	enumLookup[TEXTURE_IMAGE_VALID_QCOM] = "TEXTURE_IMAGE_VALID_QCOM"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT] = "TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT"
	enumLookup[MAX_FRAGMENT_ATOMIC_COUNTERS] = "MAX_FRAGMENT_ATOMIC_COUNTERS"
	enumLookup[FOG_DENSITY] = "FOG_DENSITY"
	enumLookup[ALPHA_TEST_QCOM] = "ALPHA_TEST_QCOM"
	enumLookup[INTENSITY] = "INTENSITY"
	enumLookup[MAX_VERTEX_SHADER_INVARIANTS_EXT] = "MAX_VERTEX_SHADER_INVARIANTS_EXT"
	enumLookup[STENCIL_BACK_FAIL_ATI] = "STENCIL_BACK_FAIL_ATI"
	enumLookup[UNSIGNED_INT_IMAGE_1D] = "UNSIGNED_INT_IMAGE_1D"
	enumLookup[MITER_REVERT_NV] = "MITER_REVERT_NV"
	enumLookup[VERTEX_WEIGHT_ARRAY_EXT] = "VERTEX_WEIGHT_ARRAY_EXT"
	enumLookup[DOT_PRODUCT_TEXTURE_1D_NV] = "DOT_PRODUCT_TEXTURE_1D_NV"
	enumLookup[CON_24_ATI] = "CON_24_ATI"
	enumLookup[TEXTURE_CUBE_MAP_ARRAY] = "TEXTURE_CUBE_MAP_ARRAY"
	enumLookup[DEBUG_SEVERITY_MEDIUM_AMD] = "DEBUG_SEVERITY_MEDIUM_AMD"
	enumLookup[BUFFER_VARIABLE] = "BUFFER_VARIABLE"
	enumLookup[EYE_POINT_SGIS] = "EYE_POINT_SGIS"
	enumLookup[VIEW_COMPATIBILITY_CLASS] = "VIEW_COMPATIBILITY_CLASS"
	enumLookup[DRAW_BUFFER3] = "DRAW_BUFFER3"
	enumLookup[READ_WRITE_ARB] = "READ_WRITE_ARB"
	enumLookup[FLOAT_MAT3_ARB] = "FLOAT_MAT3_ARB"
	enumLookup[COLOR_ATTACHMENT2] = "COLOR_ATTACHMENT2"
	enumLookup[NORMAL_ARRAY_LIST_IBM] = "NORMAL_ARRAY_LIST_IBM"
	enumLookup[SAMPLE_MASK_VALUE_SGIS] = "SAMPLE_MASK_VALUE_SGIS"
	enumLookup[INDEX] = "INDEX"
	enumLookup[DRAW_BUFFER11_NV] = "DRAW_BUFFER11_NV"
	enumLookup[SAMPLER_1D_SHADOW_ARB] = "SAMPLER_1D_SHADOW_ARB"
	enumLookup[FRAMEBUFFER_BINDING] = "FRAMEBUFFER_BINDING"
	enumLookup[RGBA_FLOAT32_ATI] = "RGBA_FLOAT32_ATI"
	enumLookup[PATH_FILL_BOUNDING_BOX_NV] = "PATH_FILL_BOUNDING_BOX_NV"
	enumLookup[ALPHA32I_EXT] = "ALPHA32I_EXT"
	enumLookup[FIRST_VERTEX_CONVENTION_EXT] = "FIRST_VERTEX_CONVENTION_EXT"
	enumLookup[DEBUG_TYPE_PORTABILITY] = "DEBUG_TYPE_PORTABILITY"
	enumLookup[RGBA32UI_EXT] = "RGBA32UI_EXT"
	enumLookup[MAP1_GRID_DOMAIN] = "MAP1_GRID_DOMAIN"
	enumLookup[PACK_IMAGE_HEIGHT] = "PACK_IMAGE_HEIGHT"
	enumLookup[COLOR_TABLE_SGI] = "COLOR_TABLE_SGI"
	enumLookup[VERTEX_PROGRAM_POINT_SIZE_NV] = "VERTEX_PROGRAM_POINT_SIZE_NV"
	enumLookup[COMP_BIT_ATI] = "COMP_BIT_ATI"
	enumLookup[DEPTH_ATTACHMENT_OES] = "DEPTH_ATTACHMENT_OES"
	enumLookup[FRAGMENT_INTERPOLATION_OFFSET_BITS] = "FRAGMENT_INTERPOLATION_OFFSET_BITS"
	enumLookup[GEOMETRY_SUBROUTINE_UNIFORM] = "GEOMETRY_SUBROUTINE_UNIFORM"
	enumLookup[SHADER_IMAGE_ATOMIC] = "SHADER_IMAGE_ATOMIC"
	enumLookup[TEXTURE_MEMORY_LAYOUT_INTEL] = "TEXTURE_MEMORY_LAYOUT_INTEL"
	enumLookup[DOT_PRODUCT_REFLECT_CUBE_MAP_NV] = "DOT_PRODUCT_REFLECT_CUBE_MAP_NV"
	enumLookup[POINT_SPRITE] = "POINT_SPRITE"
	enumLookup[ACTIVE_UNIFORM_BLOCKS] = "ACTIVE_UNIFORM_BLOCKS"
	enumLookup[UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER] = "UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER"
	enumLookup[SAMPLE_ALPHA_TO_ONE_SGIS] = "SAMPLE_ALPHA_TO_ONE_SGIS"
	enumLookup[CLAMP_TO_BORDER_SGIS] = "CLAMP_TO_BORDER_SGIS"
	enumLookup[VALIDATE_STATUS] = "VALIDATE_STATUS"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_RECT_EXT] = "UNSIGNED_INT_SAMPLER_2D_RECT_EXT"
	enumLookup[TEXTURE_BUFFER_DATA_STORE_BINDING_EXT] = "TEXTURE_BUFFER_DATA_STORE_BINDING_EXT"
	enumLookup[LINE_STIPPLE_REPEAT] = "LINE_STIPPLE_REPEAT"
	enumLookup[SPRITE_OBJECT_ALIGNED_SGIX] = "SPRITE_OBJECT_ALIGNED_SGIX"
	enumLookup[FIXED_ONLY] = "FIXED_ONLY"
	enumLookup[UNSIGNED_NORMALIZED_EXT] = "UNSIGNED_NORMALIZED_EXT"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_MODE] = "TRANSFORM_FEEDBACK_BUFFER_MODE"
	enumLookup[TEXTURE_UPDATE_BARRIER_BIT] = "TEXTURE_UPDATE_BARRIER_BIT"
	enumLookup[PRIMARY_COLOR_ARB] = "PRIMARY_COLOR_ARB"
	enumLookup[SAMPLER_1D_ARRAY_EXT] = "SAMPLER_1D_ARRAY_EXT"
	enumLookup[FOG_COORDINATE_EXT] = "FOG_COORDINATE_EXT"
	enumLookup[COMPRESSED_LUMINANCE_ALPHA] = "COMPRESSED_LUMINANCE_ALPHA"
	enumLookup[VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV] = "VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV"
	enumLookup[LOCAL_CONSTANT_EXT] = "LOCAL_CONSTANT_EXT"
	enumLookup[INT_VEC4] = "INT_VEC4"
	enumLookup[POINT] = "POINT"
	enumLookup[SECONDARY_COLOR_ARRAY_TYPE_EXT] = "SECONDARY_COLOR_ARRAY_TYPE_EXT"
	enumLookup[PRIMITIVE_RESTART_NV] = "PRIMITIVE_RESTART_NV"
	enumLookup[GREEN_MIN_CLAMP_INGR] = "GREEN_MIN_CLAMP_INGR"
	enumLookup[READ_FRAMEBUFFER_EXT] = "READ_FRAMEBUFFER_EXT"
	enumLookup[MAX_VERTEX_VARYING_COMPONENTS_ARB] = "MAX_VERTEX_VARYING_COMPONENTS_ARB"
	enumLookup[MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS] = "MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS"
	enumLookup[PROGRAM_INPUT] = "PROGRAM_INPUT"
	enumLookup[MAP2_TEXTURE_COORD_1] = "MAP2_TEXTURE_COORD_1"
	enumLookup[NEAREST_MIPMAP_LINEAR] = "NEAREST_MIPMAP_LINEAR"
	enumLookup[SECONDARY_COLOR_ARRAY_STRIDE] = "SECONDARY_COLOR_ARRAY_STRIDE"
	enumLookup[COMPRESSED_INTENSITY_ARB] = "COMPRESSED_INTENSITY_ARB"
	enumLookup[MATRIX_INDEX_ARRAY_POINTER_ARB] = "MATRIX_INDEX_ARRAY_POINTER_ARB"
	enumLookup[UNIFORM_ARRAY_STRIDE] = "UNIFORM_ARRAY_STRIDE"
	enumLookup[MAX_IMAGE_UNITS] = "MAX_IMAGE_UNITS"
	enumLookup[UNSIGNED_INT_SAMPLER_BUFFER_AMD] = "UNSIGNED_INT_SAMPLER_BUFFER_AMD"
	enumLookup[COLOR_WRITEMASK] = "COLOR_WRITEMASK"
	enumLookup[FEEDBACK_BUFFER_SIZE] = "FEEDBACK_BUFFER_SIZE"
	enumLookup[PIXEL_CUBIC_WEIGHT_EXT] = "PIXEL_CUBIC_WEIGHT_EXT"
	enumLookup[TEXTURE_BINDING_RECTANGLE_NV] = "TEXTURE_BINDING_RECTANGLE_NV"
	enumLookup[MODELVIEW13_ARB] = "MODELVIEW13_ARB"
	enumLookup[DEBUG_ASSERT_MESA] = "DEBUG_ASSERT_MESA"
	enumLookup[BUFFER_USAGE_ARB] = "BUFFER_USAGE_ARB"
	enumLookup[STENCIL_BACK_PASS_DEPTH_FAIL] = "STENCIL_BACK_PASS_DEPTH_FAIL"
	enumLookup[MAX_VIEWPORT_DIMS] = "MAX_VIEWPORT_DIMS"
	enumLookup[ASYNC_MARKER_SGIX] = "ASYNC_MARKER_SGIX"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Z_OES] = "TEXTURE_CUBE_MAP_NEGATIVE_Z_OES"
	enumLookup[COVERAGE_ALL_FRAGMENTS_NV] = "COVERAGE_ALL_FRAGMENTS_NV"
	enumLookup[INT_SAMPLER_2D_MULTISAMPLE] = "INT_SAMPLER_2D_MULTISAMPLE"
	enumLookup[NORMAL_ARRAY_STRIDE_EXT] = "NORMAL_ARRAY_STRIDE_EXT"
	enumLookup[MAX_GENERAL_COMBINERS_NV] = "MAX_GENERAL_COMBINERS_NV"
	enumLookup[DOT3_RGB] = "DOT3_RGB"
	enumLookup[MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES] = "MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES"
	enumLookup[UNSIGNED_INT_SAMPLER_2D] = "UNSIGNED_INT_SAMPLER_2D"
	enumLookup[INT_IMAGE_BUFFER] = "INT_IMAGE_BUFFER"
	enumLookup[FUNC_ADD] = "FUNC_ADD"
	enumLookup[COLOR_TABLE_GREEN_SIZE] = "COLOR_TABLE_GREEN_SIZE"
	enumLookup[INTERNALFORMAT_STENCIL_TYPE] = "INTERNALFORMAT_STENCIL_TYPE"
	enumLookup[DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV] = "DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV"
	enumLookup[NEGATIVE_Y_EXT] = "NEGATIVE_Y_EXT"
	enumLookup[SKIP_COMPONENTS1_NV] = "SKIP_COMPONENTS1_NV"
	enumLookup[COMPUTE_PROGRAM_PARAMETER_BUFFER_NV] = "COMPUTE_PROGRAM_PARAMETER_BUFFER_NV"
	enumLookup[OUT_OF_MEMORY] = "OUT_OF_MEMORY"
	enumLookup[GREEN] = "GREEN"
	enumLookup[CONVOLUTION_BORDER_MODE] = "CONVOLUTION_BORDER_MODE"
	enumLookup[VERTEX_ARRAY_STRIDE_EXT] = "VERTEX_ARRAY_STRIDE_EXT"
	enumLookup[T2F_IUI_V2F_EXT] = "T2F_IUI_V2F_EXT"
	enumLookup[TEXTURE_COMPRESSION_HINT] = "TEXTURE_COMPRESSION_HINT"
	enumLookup[OP_SET_LT_EXT] = "OP_SET_LT_EXT"
	enumLookup[STATIC_COPY] = "STATIC_COPY"
	enumLookup[PROXY_TEXTURE_1D_ARRAY_EXT] = "PROXY_TEXTURE_1D_ARRAY_EXT"
	enumLookup[ALPHA12] = "ALPHA12"
	enumLookup[DISTANCE_ATTENUATION_SGIS] = "DISTANCE_ATTENUATION_SGIS"
	enumLookup[MODELVIEW16_ARB] = "MODELVIEW16_ARB"
	enumLookup[MODULATE_SIGNED_ADD_ATI] = "MODULATE_SIGNED_ADD_ATI"
	enumLookup[PIXEL_PACK_BUFFER_BINDING] = "PIXEL_PACK_BUFFER_BINDING"
	enumLookup[ELEMENT_ARRAY_TYPE_APPLE] = "ELEMENT_ARRAY_TYPE_APPLE"
	enumLookup[COLOR_ATTACHMENT3] = "COLOR_ATTACHMENT3"
	enumLookup[SAMPLER_2D_ARRAY_EXT] = "SAMPLER_2D_ARRAY_EXT"
	enumLookup[IMAGE_2D_ARRAY] = "IMAGE_2D_ARRAY"
	enumLookup[UNSIGNED_INT_IMAGE_2D_MULTISAMPLE] = "UNSIGNED_INT_IMAGE_2D_MULTISAMPLE"
	enumLookup[VERTEX_ARRAY_POINTER] = "VERTEX_ARRAY_POINTER"
	enumLookup[MODELVIEW30_ARB] = "MODELVIEW30_ARB"
	enumLookup[EVAL_VERTEX_ATTRIB15_NV] = "EVAL_VERTEX_ATTRIB15_NV"
	enumLookup[YCRCB_422_SGIX] = "YCRCB_422_SGIX"
	enumLookup[MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB] = "MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB"
	enumLookup[VERTEX_PROGRAM_CALLBACK_FUNC_MESA] = "VERTEX_PROGRAM_CALLBACK_FUNC_MESA"
	enumLookup[LUMINANCE16_ALPHA16_SNORM] = "LUMINANCE16_ALPHA16_SNORM"
	enumLookup[VERSION_3_1] = "VERSION_3_1"
	enumLookup[MAP_WRITE_BIT_EXT] = "MAP_WRITE_BIT_EXT"
	enumLookup[SCALE_BY_ONE_HALF_NV] = "SCALE_BY_ONE_HALF_NV"
	enumLookup[CURRENT_MATRIX_ARB] = "CURRENT_MATRIX_ARB"
	enumLookup[RGB565_OES] = "RGB565_OES"
	enumLookup[UNIFORM_BUFFER_EXT] = "UNIFORM_BUFFER_EXT"
	enumLookup[SQUARE_NV] = "SQUARE_NV"
	enumLookup[ALLOW_DRAW_OBJ_HINT_PGI] = "ALLOW_DRAW_OBJ_HINT_PGI"
	enumLookup[HISTOGRAM_EXT] = "HISTOGRAM_EXT"
	enumLookup[MAX_FOG_FUNC_POINTS_SGIS] = "MAX_FOG_FUNC_POINTS_SGIS"
	enumLookup[FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT] = "FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT"
	enumLookup[BGRA_INTEGER] = "BGRA_INTEGER"
	enumLookup[IMAGE_BINDING_LEVEL] = "IMAGE_BINDING_LEVEL"
	enumLookup[GL_2PASS_1_EXT] = "GL_2PASS_1_EXT"
	enumLookup[SAMPLE_BUFFERS_EXT] = "SAMPLE_BUFFERS_EXT"
	enumLookup[FOG_COORD] = "FOG_COORD"
	enumLookup[TRANSPOSE_MODELVIEW_MATRIX_ARB] = "TRANSPOSE_MODELVIEW_MATRIX_ARB"
	enumLookup[WEIGHT_ARRAY_STRIDE_ARB] = "WEIGHT_ARRAY_STRIDE_ARB"
	enumLookup[COORD_REPLACE] = "COORD_REPLACE"
	enumLookup[POLYGON_SMOOTH_HINT] = "POLYGON_SMOOTH_HINT"
	enumLookup[C4UB_V2F] = "C4UB_V2F"
	enumLookup[GL_4PASS_2_EXT] = "GL_4PASS_2_EXT"
	enumLookup[TESS_CONTROL_TEXTURE] = "TESS_CONTROL_TEXTURE"
	enumLookup[CND_ATI] = "CND_ATI"
	enumLookup[EXT_polygon_offset] = "EXT_polygon_offset"
	enumLookup[GREATER] = "GREATER"
	enumLookup[MAX_TEXTURE_STACK_DEPTH] = "MAX_TEXTURE_STACK_DEPTH"
	enumLookup[TEXTURE_WRAP_R] = "TEXTURE_WRAP_R"
	enumLookup[PACK_ROW_BYTES_APPLE] = "PACK_ROW_BYTES_APPLE"
	enumLookup[RGB8UI] = "RGB8UI"
	enumLookup[INT8_VEC4_NV] = "INT8_VEC4_NV"
	enumLookup[BLUE_BIAS] = "BLUE_BIAS"
	enumLookup[PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI] = "PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI"
	enumLookup[IUI_N3F_V2F_EXT] = "IUI_N3F_V2F_EXT"
	enumLookup[OP_MOV_EXT] = "OP_MOV_EXT"
	enumLookup[SAMPLER_CUBE] = "SAMPLER_CUBE"
	enumLookup[LINEAR_MIPMAP_NEAREST] = "LINEAR_MIPMAP_NEAREST"
	enumLookup[COMPRESSED_TEXTURE_FORMATS_ARB] = "COMPRESSED_TEXTURE_FORMATS_ARB"
	enumLookup[SCALE_BY_FOUR_NV] = "SCALE_BY_FOUR_NV"
	enumLookup[UNSIGNED_SHORT] = "UNSIGNED_SHORT"
	enumLookup[DUAL_ALPHA12_SGIS] = "DUAL_ALPHA12_SGIS"
	enumLookup[MODELVIEW25_ARB] = "MODELVIEW25_ARB"
	enumLookup[OUTPUT_FOG_EXT] = "OUTPUT_FOG_EXT"
	enumLookup[MAX_MATRIX_PALETTE_STACK_DEPTH_ARB] = "MAX_MATRIX_PALETTE_STACK_DEPTH_ARB"
	enumLookup[SPRITE_MODE_SGIX] = "SPRITE_MODE_SGIX"
	enumLookup[MAX_TESS_EVALUATION_INPUT_COMPONENTS] = "MAX_TESS_EVALUATION_INPUT_COMPONENTS"
	enumLookup[UNSIGNED_INT_SAMPLER_1D] = "UNSIGNED_INT_SAMPLER_1D"
	enumLookup[FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV] = "FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV"
	enumLookup[CLIP_DISTANCE0] = "CLIP_DISTANCE0"
	enumLookup[PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS] = "PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS"
	enumLookup[NEGATIVE_ONE_EXT] = "NEGATIVE_ONE_EXT"
	enumLookup[INDEX_MODE] = "INDEX_MODE"
	enumLookup[POST_CONVOLUTION_BLUE_SCALE_EXT] = "POST_CONVOLUTION_BLUE_SCALE_EXT"
	enumLookup[TEXTURE_CLIPMAP_FRAME_SGIX] = "TEXTURE_CLIPMAP_FRAME_SGIX"
	enumLookup[AVERAGE_EXT] = "AVERAGE_EXT"
	enumLookup[TRACE_ARRAYS_BIT_MESA] = "TRACE_ARRAYS_BIT_MESA"
	enumLookup[MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB] = "MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB"
	enumLookup[TEXTURE_LUMINANCE_TYPE_ARB] = "TEXTURE_LUMINANCE_TYPE_ARB"
	enumLookup[COLOR_ATTACHMENT4] = "COLOR_ATTACHMENT4"
	enumLookup[LUMINANCE16I_EXT] = "LUMINANCE16I_EXT"
	enumLookup[FOG] = "FOG"
	enumLookup[LUMINANCE_ALPHA_FLOAT16_ATI] = "LUMINANCE_ALPHA_FLOAT16_ATI"
	enumLookup[TEXTURE_STENCIL_SIZE] = "TEXTURE_STENCIL_SIZE"
	enumLookup[MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV] = "MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV"
	enumLookup[TEXTURE_COORD_ARRAY_POINTER] = "TEXTURE_COORD_ARRAY_POINTER"
	enumLookup[TEXTURE_COMPARE_FAIL_VALUE_ARB] = "TEXTURE_COMPARE_FAIL_VALUE_ARB"
	enumLookup[MIRRORED_REPEAT_IBM] = "MIRRORED_REPEAT_IBM"
	enumLookup[DOT_PRODUCT_TEXTURE_CUBE_MAP_NV] = "DOT_PRODUCT_TEXTURE_CUBE_MAP_NV"
	enumLookup[RGB16F] = "RGB16F"
	enumLookup[TEXTURE_BLUE_TYPE] = "TEXTURE_BLUE_TYPE"
	enumLookup[MAX_VERTEX_SHADER_STORAGE_BLOCKS] = "MAX_VERTEX_SHADER_STORAGE_BLOCKS"
	enumLookup[RG16UI] = "RG16UI"
	enumLookup[PROXY_TEXTURE_RECTANGLE_NV] = "PROXY_TEXTURE_RECTANGLE_NV"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_NV] = "TRANSFORM_FEEDBACK_BUFFER_NV"
	enumLookup[IMAGE_2D_MULTISAMPLE_ARRAY] = "IMAGE_2D_MULTISAMPLE_ARRAY"
	enumLookup[TEXTURE_COMPARE_MODE_ARB] = "TEXTURE_COMPARE_MODE_ARB"
	enumLookup[CON_13_ATI] = "CON_13_ATI"
	enumLookup[IMAGE_1D_ARRAY_EXT] = "IMAGE_1D_ARRAY_EXT"
	enumLookup[EXT_texture3D] = "EXT_texture3D"
	enumLookup[DUAL_LUMINANCE4_SGIS] = "DUAL_LUMINANCE4_SGIS"
	enumLookup[VERTEX_ARRAY_BINDING_APPLE] = "VERTEX_ARRAY_BINDING_APPLE"
	enumLookup[TRANSPOSE_PROGRAM_MATRIX_EXT] = "TRANSPOSE_PROGRAM_MATRIX_EXT"
	enumLookup[DEBUG_CATEGORY_OTHER_AMD] = "DEBUG_CATEGORY_OTHER_AMD"
	enumLookup[DRAW_BUFFER10_ATI] = "DRAW_BUFFER10_ATI"
	enumLookup[DYNAMIC_READ] = "DYNAMIC_READ"
	enumLookup[MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS] = "MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS"
	enumLookup[COMPRESSED_RGBA_ASTC_10x5_KHR] = "COMPRESSED_RGBA_ASTC_10x5_KHR"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR"
	enumLookup[SGIX_pixel_tiles] = "SGIX_pixel_tiles"
	enumLookup[MAP1_TEXTURE_COORD_2] = "MAP1_TEXTURE_COORD_2"
	enumLookup[DUAL_INTENSITY4_SGIS] = "DUAL_INTENSITY4_SGIS"
	enumLookup[PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX] = "PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX"
	enumLookup[COMBINE_RGB_EXT] = "COMBINE_RGB_EXT"
	enumLookup[PROGRAM_LENGTH_NV] = "PROGRAM_LENGTH_NV"
	enumLookup[DRAW_BUFFER2_ARB] = "DRAW_BUFFER2_ARB"
	enumLookup[FLOAT_RGBA_NV] = "FLOAT_RGBA_NV"
	enumLookup[MODELVIEW_STACK_DEPTH] = "MODELVIEW_STACK_DEPTH"
	enumLookup[QUAD_LUMINANCE4_SGIS] = "QUAD_LUMINANCE4_SGIS"
	enumLookup[TANGENT_ARRAY_POINTER_EXT] = "TANGENT_ARRAY_POINTER_EXT"
	enumLookup[MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB] = "MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB"
	enumLookup[FRAGMENT_SUBROUTINE] = "FRAGMENT_SUBROUTINE"
	enumLookup[COLOR_ENCODING] = "COLOR_ENCODING"
	enumLookup[RGB4_S3TC] = "RGB4_S3TC"
	enumLookup[NORMAL_ARRAY_PARALLEL_POINTERS_INTEL] = "NORMAL_ARRAY_PARALLEL_POINTERS_INTEL"
	enumLookup[LAYER_NV] = "LAYER_NV"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_RECT] = "UNSIGNED_INT_SAMPLER_2D_RECT"
	enumLookup[FRAMEBUFFER_DEFAULT_HEIGHT] = "FRAMEBUFFER_DEFAULT_HEIGHT"
	enumLookup[CALLIGRAPHIC_FRAGMENT_SGIX] = "CALLIGRAPHIC_FRAGMENT_SGIX"
	enumLookup[MAX_PROGRAM_MATRICES_ARB] = "MAX_PROGRAM_MATRICES_ARB"
	enumLookup[GL_2X_BIT_ATI] = "GL_2X_BIT_ATI"
	enumLookup[CURRENT_PROGRAM] = "CURRENT_PROGRAM"
	enumLookup[FACTOR_MAX_AMD] = "FACTOR_MAX_AMD"
	enumLookup[INDEX_BIT_PGI] = "INDEX_BIT_PGI"
	enumLookup[SGIX_sprite] = "SGIX_sprite"
	enumLookup[STREAM_DRAW] = "STREAM_DRAW"
	enumLookup[VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV] = "VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV"
	enumLookup[COMPUTE_SUBROUTINE] = "COMPUTE_SUBROUTINE"
	enumLookup[RETURN] = "RETURN"
	enumLookup[LEQUAL] = "LEQUAL"
	enumLookup[LINE_STRIP] = "LINE_STRIP"
	enumLookup[NONE] = "NONE"
	enumLookup[RGBA4_DXT5_S3TC] = "RGBA4_DXT5_S3TC"
	enumLookup[FOG_COORD_ARRAY] = "FOG_COORD_ARRAY"
	enumLookup[OUTPUT_TEXTURE_COORD11_EXT] = "OUTPUT_TEXTURE_COORD11_EXT"
	enumLookup[MAX_VERTEX_SHADER_INSTRUCTIONS_EXT] = "MAX_VERTEX_SHADER_INSTRUCTIONS_EXT"
	enumLookup[TEXTURE_BINDING_RENDERBUFFER_NV] = "TEXTURE_BINDING_RENDERBUFFER_NV"
	enumLookup[DRAW_INDIRECT_BUFFER_BINDING] = "DRAW_INDIRECT_BUFFER_BINDING"
	enumLookup[RENDERER] = "RENDERER"
	enumLookup[LINEAR_CLIPMAP_NEAREST_SGIX] = "LINEAR_CLIPMAP_NEAREST_SGIX"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_START_NV] = "TRANSFORM_FEEDBACK_BUFFER_START_NV"
	enumLookup[FRAMEBUFFER_UNSUPPORTED_OES] = "FRAMEBUFFER_UNSUPPORTED_OES"
	enumLookup[COLOR_ARRAY_ADDRESS_NV] = "COLOR_ARRAY_ADDRESS_NV"
	enumLookup[CURRENT_VERTEX_ATTRIB] = "CURRENT_VERTEX_ATTRIB"
	enumLookup[MODELVIEW26_ARB] = "MODELVIEW26_ARB"
	enumLookup[DOUBLE_VEC3] = "DOUBLE_VEC3"
	enumLookup[FASTEST] = "FASTEST"
	enumLookup[COMPRESSED_RGB_PVRTC_2BPPV1_IMG] = "COMPRESSED_RGB_PVRTC_2BPPV1_IMG"
	enumLookup[TRANSFORM_FEEDBACK_ATTRIBS_NV] = "TRANSFORM_FEEDBACK_ATTRIBS_NV"
	enumLookup[STENCIL_INDEX4] = "STENCIL_INDEX4"
	enumLookup[INT_SAMPLER_2D_EXT] = "INT_SAMPLER_2D_EXT"
	enumLookup[POST_COLOR_MATRIX_ALPHA_BIAS] = "POST_COLOR_MATRIX_ALPHA_BIAS"
	enumLookup[WEIGHT_ARRAY_ARB] = "WEIGHT_ARRAY_ARB"
	enumLookup[COMPRESSED_RGBA_BPTC_UNORM_ARB] = "COMPRESSED_RGBA_BPTC_UNORM_ARB"
	enumLookup[EYE_PLANE] = "EYE_PLANE"
	enumLookup[CURRENT_TANGENT_EXT] = "CURRENT_TANGENT_EXT"
	enumLookup[DRAW_BUFFER3_NV] = "DRAW_BUFFER3_NV"
	enumLookup[WRITE_ONLY_OES] = "WRITE_ONLY_OES"
	enumLookup[TEXTURE_SWIZZLE_G_EXT] = "TEXTURE_SWIZZLE_G_EXT"
	enumLookup[COMPRESSED_SIGNED_R11_EAC] = "COMPRESSED_SIGNED_R11_EAC"
	enumLookup[PATCHES] = "PATCHES"
	enumLookup[ZOOM_Y] = "ZOOM_Y"
	enumLookup[MULTISAMPLE_FILTER_HINT_NV] = "MULTISAMPLE_FILTER_HINT_NV"
	enumLookup[COMBINE_ALPHA] = "COMBINE_ALPHA"
	enumLookup[UNSIGNED_SHORT_15_1_MESA] = "UNSIGNED_SHORT_15_1_MESA"
	enumLookup[PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB] = "PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB"
	enumLookup[FIELD_LOWER_NV] = "FIELD_LOWER_NV"
	enumLookup[COLOR_ATTACHMENT3_EXT] = "COLOR_ATTACHMENT3_EXT"
	enumLookup[TRIANGLE_STRIP] = "TRIANGLE_STRIP"
	enumLookup[TEXTURE_CLIPMAP_LOD_OFFSET_SGIX] = "TEXTURE_CLIPMAP_LOD_OFFSET_SGIX"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_X] = "TEXTURE_CUBE_MAP_POSITIVE_X"
	enumLookup[MAX_VERTEX_UNIFORM_COMPONENTS_ARB] = "MAX_VERTEX_UNIFORM_COMPONENTS_ARB"
	enumLookup[NUM_SHADER_BINARY_FORMATS] = "NUM_SHADER_BINARY_FORMATS"
	enumLookup[MAX_TESS_CONTROL_UNIFORM_BLOCKS] = "MAX_TESS_CONTROL_UNIFORM_BLOCKS"
	enumLookup[LINES_ADJACENCY] = "LINES_ADJACENCY"
	enumLookup[OUTPUT_TEXTURE_COORD28_EXT] = "OUTPUT_TEXTURE_COORD28_EXT"
	enumLookup[OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV] = "OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV"
	enumLookup[SAMPLER_CUBE_MAP_ARRAY_SHADOW] = "SAMPLER_CUBE_MAP_ARRAY_SHADOW"
	enumLookup[EXT_convolution] = "EXT_convolution"
	enumLookup[EMBOSS_CONSTANT_NV] = "EMBOSS_CONSTANT_NV"
	enumLookup[LOCAL_CONSTANT_DATATYPE_EXT] = "LOCAL_CONSTANT_DATATYPE_EXT"
	enumLookup[RENDERBUFFER_HEIGHT_EXT] = "RENDERBUFFER_HEIGHT_EXT"
	enumLookup[UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY] = "UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY"
	enumLookup[IMAGE_CLASS_1_X_32] = "IMAGE_CLASS_1_X_32"
	enumLookup[FRAGMENT_LIGHT6_SGIX] = "FRAGMENT_LIGHT6_SGIX"
	enumLookup[OFFSET_TEXTURE_RECTANGLE_SCALE_NV] = "OFFSET_TEXTURE_RECTANGLE_SCALE_NV"
	enumLookup[ALIASED_LINE_WIDTH_RANGE] = "ALIASED_LINE_WIDTH_RANGE"
	enumLookup[VERTEX_ATTRIB_ARRAY13_NV] = "VERTEX_ATTRIB_ARRAY13_NV"
	enumLookup[TRACE_ALL_BITS_MESA] = "TRACE_ALL_BITS_MESA"
	enumLookup[DEPTH_COMPONENT24] = "DEPTH_COMPONENT24"
	enumLookup[SRGB_READ] = "SRGB_READ"
	enumLookup[MODELVIEW17_ARB] = "MODELVIEW17_ARB"
	enumLookup[REG_13_ATI] = "REG_13_ATI"
	enumLookup[UNSIGNED_SHORT_8_8_REV_MESA] = "UNSIGNED_SHORT_8_8_REV_MESA"
	enumLookup[SIGNED_RGBA_NV] = "SIGNED_RGBA_NV"
	enumLookup[VERTEX_ATTRIB_MAP2_APPLE] = "VERTEX_ATTRIB_MAP2_APPLE"
	enumLookup[FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES] = "FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES"
	enumLookup[TEXTURE] = "TEXTURE"
	enumLookup[LUMINANCE8_ALPHA8] = "LUMINANCE8_ALPHA8"
	enumLookup[DRAW_BUFFER14] = "DRAW_BUFFER14"
	enumLookup[PROGRAM_ATTRIBS_ARB] = "PROGRAM_ATTRIBS_ARB"
	enumLookup[NEGATIVE_W_EXT] = "NEGATIVE_W_EXT"
	enumLookup[MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT] = "MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT"
	enumLookup[PROVOKING_VERTEX] = "PROVOKING_VERTEX"
	enumLookup[SAMPLE_MASK] = "SAMPLE_MASK"
	enumLookup[PATH_COMMAND_COUNT_NV] = "PATH_COMMAND_COUNT_NV"
	enumLookup[PIXEL_TILE_GRID_DEPTH_SGIX] = "PIXEL_TILE_GRID_DEPTH_SGIX"
	enumLookup[OP_CLAMP_EXT] = "OP_CLAMP_EXT"
	enumLookup[MATRIX15_ARB] = "MATRIX15_ARB"
	enumLookup[UNSIGNED_INT_VEC2_EXT] = "UNSIGNED_INT_VEC2_EXT"
	enumLookup[POINT_BIT] = "POINT_BIT"
	enumLookup[ALPHA_TEST_REF_QCOM] = "ALPHA_TEST_REF_QCOM"
	enumLookup[LUMINANCE12_ALPHA4_EXT] = "LUMINANCE12_ALPHA4_EXT"
	enumLookup[BLEND_SRC_ALPHA_EXT] = "BLEND_SRC_ALPHA_EXT"
	enumLookup[TEXTURE2_ARB] = "TEXTURE2_ARB"
	enumLookup[TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT] = "TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT"
	enumLookup[LINE_STRIP_ADJACENCY_ARB] = "LINE_STRIP_ADJACENCY_ARB"
	enumLookup[RESCALE_NORMAL] = "RESCALE_NORMAL"
	enumLookup[BLEND_DST_ALPHA_OES] = "BLEND_DST_ALPHA_OES"
	enumLookup[TEXTURE_COMPARE_SGIX] = "TEXTURE_COMPARE_SGIX"
	enumLookup[UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER] = "UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER"
	enumLookup[TEXTURE_SWIZZLE_G] = "TEXTURE_SWIZZLE_G"
	enumLookup[PACK_COMPRESSED_BLOCK_SIZE] = "PACK_COMPRESSED_BLOCK_SIZE"
	enumLookup[SGIS_texture4D] = "SGIS_texture4D"
	enumLookup[VARIANT_ARRAY_TYPE_EXT] = "VARIANT_ARRAY_TYPE_EXT"
	enumLookup[SCALED_RESOLVE_FASTEST_EXT] = "SCALED_RESOLVE_FASTEST_EXT"
	enumLookup[FRAGMENT_LIGHTING_SGIX] = "FRAGMENT_LIGHTING_SGIX"
	enumLookup[GEOMETRY_OUTPUT_TYPE] = "GEOMETRY_OUTPUT_TYPE"
	enumLookup[TEXTURE_COORD_ARRAY_TYPE] = "TEXTURE_COORD_ARRAY_TYPE"
	enumLookup[GL_2_BYTES] = "GL_2_BYTES"
	enumLookup[S] = "S"
	enumLookup[FRAGMENT_TEXTURE] = "FRAGMENT_TEXTURE"
	enumLookup[MAX_UNIFORM_BUFFER_BINDINGS] = "MAX_UNIFORM_BUFFER_BINDINGS"
	enumLookup[UNSIGNED_INT_8_8_8_8_EXT] = "UNSIGNED_INT_8_8_8_8_EXT"
	enumLookup[PROGRAM_BINARY_RETRIEVABLE_HINT] = "PROGRAM_BINARY_RETRIEVABLE_HINT"
	enumLookup[OPERAND2_ALPHA_EXT] = "OPERAND2_ALPHA_EXT"
	enumLookup[CURRENT_ATTRIB_NV] = "CURRENT_ATTRIB_NV"
	enumLookup[LO_BIAS_NV] = "LO_BIAS_NV"
	enumLookup[DRAW_BUFFER3_ARB] = "DRAW_BUFFER3_ARB"
	enumLookup[COLOR_ATTACHMENT10_EXT] = "COLOR_ATTACHMENT10_EXT"
	enumLookup[INT_SAMPLER_3D] = "INT_SAMPLER_3D"
	enumLookup[VERSION_1_5] = "VERSION_1_5"
	enumLookup[DST_COLOR] = "DST_COLOR"
	enumLookup[HISTOGRAM_ALPHA_SIZE] = "HISTOGRAM_ALPHA_SIZE"
	enumLookup[DEBUG_GROUP_STACK_DEPTH] = "DEBUG_GROUP_STACK_DEPTH"
	enumLookup[QUERY_COUNTER_BITS] = "QUERY_COUNTER_BITS"
	enumLookup[RENDERBUFFER_COVERAGE_SAMPLES_NV] = "RENDERBUFFER_COVERAGE_SAMPLES_NV"
	enumLookup[FULL_STIPPLE_HINT_PGI] = "FULL_STIPPLE_HINT_PGI"
	enumLookup[EDGE_FLAG_ARRAY_STRIDE] = "EDGE_FLAG_ARRAY_STRIDE"
	enumLookup[TEXTURE_INDEX_SIZE_EXT] = "TEXTURE_INDEX_SIZE_EXT"
	enumLookup[CON_9_ATI] = "CON_9_ATI"
	enumLookup[UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER] = "UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER"
	enumLookup[TRANSFORM_FEEDBACK_BINDING_NV] = "TRANSFORM_FEEDBACK_BINDING_NV"
	enumLookup[SAMPLE_COVERAGE_ARB] = "SAMPLE_COVERAGE_ARB"
	enumLookup[IMAGE_CLASS_1_X_8] = "IMAGE_CLASS_1_X_8"
	enumLookup[STORAGE_CACHED_APPLE] = "STORAGE_CACHED_APPLE"
	enumLookup[OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV] = "OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV"
	enumLookup[CON_26_ATI] = "CON_26_ATI"
	enumLookup[FRAMEBUFFER_SRGB_EXT] = "FRAMEBUFFER_SRGB_EXT"
	enumLookup[SAMPLER_2D_ARRAY] = "SAMPLER_2D_ARRAY"
	enumLookup[GL_422_REV_AVERAGE_EXT] = "GL_422_REV_AVERAGE_EXT"
	enumLookup[REG_24_ATI] = "REG_24_ATI"
	enumLookup[NORMAL_ARRAY_POINTER] = "NORMAL_ARRAY_POINTER"
	enumLookup[SMOOTH_LINE_WIDTH_GRANULARITY] = "SMOOTH_LINE_WIDTH_GRANULARITY"
	enumLookup[MATRIX12_ARB] = "MATRIX12_ARB"
	enumLookup[TEXTURE_2D_ARRAY_EXT] = "TEXTURE_2D_ARRAY_EXT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES] = "FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES"
	enumLookup[PATH_COMPUTED_LENGTH_NV] = "PATH_COMPUTED_LENGTH_NV"
	enumLookup[MAP2_VERTEX_3] = "MAP2_VERTEX_3"
	enumLookup[SECONDARY_COLOR_NV] = "SECONDARY_COLOR_NV"
	enumLookup[ADD_SIGNED] = "ADD_SIGNED"
	enumLookup[HALF_FLOAT] = "HALF_FLOAT"
	enumLookup[OBJECT_DISTANCE_TO_LINE_SGIS] = "OBJECT_DISTANCE_TO_LINE_SGIS"
	enumLookup[FLOAT_RGB16_NV] = "FLOAT_RGB16_NV"
	enumLookup[SGIX_calligraphic_fragment] = "SGIX_calligraphic_fragment"
	enumLookup[GL_4PASS_2_SGIS] = "GL_4PASS_2_SGIS"
	enumLookup[MATRIX_INDEX_ARRAY_SIZE_ARB] = "MATRIX_INDEX_ARRAY_SIZE_ARB"
	enumLookup[INT_SAMPLER_BUFFER_AMD] = "INT_SAMPLER_BUFFER_AMD"
	enumLookup[MAP_WRITE_BIT] = "MAP_WRITE_BIT"
	enumLookup[CUBIC_EXT] = "CUBIC_EXT"
	enumLookup[COMBINER0_NV] = "COMBINER0_NV"
	enumLookup[OPERAND2_RGB_EXT] = "OPERAND2_RGB_EXT"
	enumLookup[COMPARE_R_TO_TEXTURE] = "COMPARE_R_TO_TEXTURE"
	enumLookup[NORMAL_ARRAY_BUFFER_BINDING_ARB] = "NORMAL_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS] = "MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS"
	enumLookup[MAX_CLIP_DISTANCES] = "MAX_CLIP_DISTANCES"
	enumLookup[DETAIL_TEXTURE_2D_SGIS] = "DETAIL_TEXTURE_2D_SGIS"
	enumLookup[CLEAR_BUFFER] = "CLEAR_BUFFER"
	enumLookup[IMAGE_CLASS_2_X_32] = "IMAGE_CLASS_2_X_32"
	enumLookup[OPERAND2_ALPHA] = "OPERAND2_ALPHA"
	enumLookup[SIGNED_LUMINANCE_NV] = "SIGNED_LUMINANCE_NV"
	enumLookup[PROXY_TEXTURE_1D_STACK_MESAX] = "PROXY_TEXTURE_1D_STACK_MESAX"
	enumLookup[MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT] = "MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT"
	enumLookup[MAP_FLUSH_EXPLICIT_BIT] = "MAP_FLUSH_EXPLICIT_BIT"
	enumLookup[POST_COLOR_MATRIX_COLOR_TABLE] = "POST_COLOR_MATRIX_COLOR_TABLE"
	enumLookup[SWIZZLE_STR_ATI] = "SWIZZLE_STR_ATI"
	enumLookup[TEXTURE_BUFFER_DATA_STORE_BINDING] = "TEXTURE_BUFFER_DATA_STORE_BINDING"
	enumLookup[REFERENCED_BY_COMPUTE_SHADER] = "REFERENCED_BY_COMPUTE_SHADER"
	enumLookup[SINGLE_COLOR_EXT] = "SINGLE_COLOR_EXT"
	enumLookup[TEXTURE10_ARB] = "TEXTURE10_ARB"
	enumLookup[SUBSAMPLE_DISTANCE_AMD] = "SUBSAMPLE_DISTANCE_AMD"
	enumLookup[IMAGE_BINDING_LAYER_EXT] = "IMAGE_BINDING_LAYER_EXT"
	enumLookup[FLOAT16_NV] = "FLOAT16_NV"
	enumLookup[VERTEX_ARRAY_EXT] = "VERTEX_ARRAY_EXT"
	enumLookup[DRAW_BUFFER2_ATI] = "DRAW_BUFFER2_ATI"
	enumLookup[PERFMON_RESULT_SIZE_AMD] = "PERFMON_RESULT_SIZE_AMD"
	enumLookup[RENDERBUFFER_SAMPLES_ANGLE] = "RENDERBUFFER_SAMPLES_ANGLE"
	enumLookup[MAX_FRAGMENT_INTERPOLATION_OFFSET_NV] = "MAX_FRAGMENT_INTERPOLATION_OFFSET_NV"
	enumLookup[PATH_INITIAL_END_CAP_NV] = "PATH_INITIAL_END_CAP_NV"
	enumLookup[PIXEL_MAP_S_TO_S_SIZE] = "PIXEL_MAP_S_TO_S_SIZE"
	enumLookup[GREEN_BIAS] = "GREEN_BIAS"
	enumLookup[RGBA4] = "RGBA4"
	enumLookup[SPRITE_TRANSLATION_SGIX] = "SPRITE_TRANSLATION_SGIX"
	enumLookup[TEXTURE15_ARB] = "TEXTURE15_ARB"
	enumLookup[COMPRESSED_INTENSITY] = "COMPRESSED_INTENSITY"
	enumLookup[OUTPUT_TEXTURE_COORD10_EXT] = "OUTPUT_TEXTURE_COORD10_EXT"
	enumLookup[SAMPLER_1D_ARRAY_SHADOW] = "SAMPLER_1D_ARRAY_SHADOW"
	enumLookup[IMAGE_BINDING_FORMAT_EXT] = "IMAGE_BINDING_FORMAT_EXT"
	enumLookup[TEXTURE_MATRIX] = "TEXTURE_MATRIX"
	enumLookup[EYE_LINEAR] = "EYE_LINEAR"
	enumLookup[NEAREST] = "NEAREST"
	enumLookup[MAP1_VERTEX_ATTRIB2_4_NV] = "MAP1_VERTEX_ATTRIB2_4_NV"
	enumLookup[TEXTURE_2D_STACK_BINDING_MESAX] = "TEXTURE_2D_STACK_BINDING_MESAX"
	enumLookup[COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB] = "COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB"
	enumLookup[TEXTURE_INTERNAL_FORMAT] = "TEXTURE_INTERNAL_FORMAT"
	enumLookup[V2F] = "V2F"
	enumLookup[IUI_V2F_EXT] = "IUI_V2F_EXT"
	enumLookup[FRAGMENT_LIGHT0_SGIX] = "FRAGMENT_LIGHT0_SGIX"
	enumLookup[MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS] = "MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS"
	enumLookup[DEPTH_TEST] = "DEPTH_TEST"
	enumLookup[SHADOW_AMBIENT_SGIX] = "SHADOW_AMBIENT_SGIX"
	enumLookup[VERTEX_ATTRIB_MAP2_COEFF_APPLE] = "VERTEX_ATTRIB_MAP2_COEFF_APPLE"
	enumLookup[IMAGE_BUFFER] = "IMAGE_BUFFER"
	enumLookup[STACK_OVERFLOW] = "STACK_OVERFLOW"
	enumLookup[PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB] = "PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB"
	enumLookup[MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB] = "MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB"
	enumLookup[VERSION_3_2] = "VERSION_3_2"
	enumLookup[MAP1_TEXTURE_COORD_4] = "MAP1_TEXTURE_COORD_4"
	enumLookup[DEBUG_SOURCE_API] = "DEBUG_SOURCE_API"
	enumLookup[TEXTURE23_ARB] = "TEXTURE23_ARB"
	enumLookup[NEGATIVE_Z_EXT] = "NEGATIVE_Z_EXT"
	enumLookup[LUMINANCE_ALPHA32F_ARB] = "LUMINANCE_ALPHA32F_ARB"
	enumLookup[GREEN_BIT_ATI] = "GREEN_BIT_ATI"
	enumLookup[VERTEX_PROGRAM_POSITION_MESA] = "VERTEX_PROGRAM_POSITION_MESA"
	enumLookup[UNSIGNED_INT64_VEC2_NV] = "UNSIGNED_INT64_VEC2_NV"
	enumLookup[ACCUM_ALPHA_BITS] = "ACCUM_ALPHA_BITS"
	enumLookup[TEXTURE_PRIORITY] = "TEXTURE_PRIORITY"
	enumLookup[DEPTH] = "DEPTH"
	enumLookup[MAP2_VERTEX_ATTRIB15_4_NV] = "MAP2_VERTEX_ATTRIB15_4_NV"
	enumLookup[GL_3D_COLOR_TEXTURE] = "GL_3D_COLOR_TEXTURE"
	enumLookup[IMAGE_ROTATE_ORIGIN_Y_HP] = "IMAGE_ROTATE_ORIGIN_Y_HP"
	enumLookup[COMBINE4_NV] = "COMBINE4_NV"
	enumLookup[VERTEX_PROGRAM_ARB] = "VERTEX_PROGRAM_ARB"
	enumLookup[DRAW_BUFFER2] = "DRAW_BUFFER2"
	enumLookup[DRAW_BUFFER4_NV] = "DRAW_BUFFER4_NV"
	enumLookup[DEPTH_ATTACHMENT] = "DEPTH_ATTACHMENT"
	enumLookup[PATH_STENCIL_VALUE_MASK_NV] = "PATH_STENCIL_VALUE_MASK_NV"
	enumLookup[MAX_DEFORMATION_ORDER_SGIX] = "MAX_DEFORMATION_ORDER_SGIX"
	enumLookup[TEXTURE6_ARB] = "TEXTURE6_ARB"
	enumLookup[TRANSFORM_HINT_APPLE] = "TRANSFORM_HINT_APPLE"
	enumLookup[SHADER_BINARY_FORMATS] = "SHADER_BINARY_FORMATS"
	enumLookup[VIDEO_CAPTURE_SURFACE_ORIGIN_NV] = "VIDEO_CAPTURE_SURFACE_ORIGIN_NV"
	enumLookup[MAP2_TEXTURE_COORD_2] = "MAP2_TEXTURE_COORD_2"
	enumLookup[TEXTURE_INTENSITY_SIZE_EXT] = "TEXTURE_INTENSITY_SIZE_EXT"
	enumLookup[REFLECTION_MAP_ARB] = "REFLECTION_MAP_ARB"
	enumLookup[ALPHA_MAX_CLAMP_INGR] = "ALPHA_MAX_CLAMP_INGR"
	enumLookup[DRAW_BUFFER2_NV] = "DRAW_BUFFER2_NV"
	enumLookup[UNSIGNED_INT_10F_11F_11F_REV_EXT] = "UNSIGNED_INT_10F_11F_11F_REV_EXT"
	enumLookup[MAX_COMPUTE_TEXTURE_IMAGE_UNITS] = "MAX_COMPUTE_TEXTURE_IMAGE_UNITS"
	enumLookup[TEXTURE_WRAP_Q_SGIS] = "TEXTURE_WRAP_Q_SGIS"
	enumLookup[DISPLAY_LIST] = "DISPLAY_LIST"
	enumLookup[COLOR_ATTACHMENT12_NV] = "COLOR_ATTACHMENT12_NV"
	enumLookup[RGB32UI_EXT] = "RGB32UI_EXT"
	enumLookup[NATIVE_GRAPHICS_BEGIN_HINT_PGI] = "NATIVE_GRAPHICS_BEGIN_HINT_PGI"
	enumLookup[VERSION_2_0] = "VERSION_2_0"
	enumLookup[VERTEX_ATTRIB_ARRAY_STRIDE_ARB] = "VERTEX_ATTRIB_ARRAY_STRIDE_ARB"
	enumLookup[SIGNED_HILO16_NV] = "SIGNED_HILO16_NV"
	enumLookup[BUMP_TEX_UNITS_ATI] = "BUMP_TEX_UNITS_ATI"
	enumLookup[COMPRESSED_SRGB] = "COMPRESSED_SRGB"
	enumLookup[ALPHA_TEST] = "ALPHA_TEST"
	enumLookup[SMOOTH_POINT_SIZE_GRANULARITY] = "SMOOTH_POINT_SIZE_GRANULARITY"
	enumLookup[CON_16_ATI] = "CON_16_ATI"
	enumLookup[TEXTURE_BINDING_2D_ARRAY] = "TEXTURE_BINDING_2D_ARRAY"
	enumLookup[PIXEL_MAP_B_TO_B] = "PIXEL_MAP_B_TO_B"
	enumLookup[MAX_PROGRAM_OUTPUT_VERTICES_NV] = "MAX_PROGRAM_OUTPUT_VERTICES_NV"
	enumLookup[PROXY_TEXTURE_RECTANGLE] = "PROXY_TEXTURE_RECTANGLE"
	enumLookup[MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV] = "MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV"
	enumLookup[FOG_COORDINATE_ARRAY_BUFFER_BINDING] = "FOG_COORDINATE_ARRAY_BUFFER_BINDING"
	enumLookup[ONE_MINUS_SRC1_ALPHA] = "ONE_MINUS_SRC1_ALPHA"
	enumLookup[TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES] = "TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES"
	enumLookup[OBJECT_ACTIVE_UNIFORMS_ARB] = "OBJECT_ACTIVE_UNIFORMS_ARB"
	enumLookup[COMPRESSED_RGBA_ASTC_6x6_KHR] = "COMPRESSED_RGBA_ASTC_6x6_KHR"
	enumLookup[SELECTION_BUFFER_POINTER] = "SELECTION_BUFFER_POINTER"
	enumLookup[SRGB8_ALPHA8_EXT] = "SRGB8_ALPHA8_EXT"
	enumLookup[MAX_GEOMETRY_SHADER_STORAGE_BLOCKS] = "MAX_GEOMETRY_SHADER_STORAGE_BLOCKS"
	enumLookup[SAMPLER_OBJECT_AMD] = "SAMPLER_OBJECT_AMD"
	enumLookup[MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS] = "MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS"
	enumLookup[FRACTIONAL_ODD] = "FRACTIONAL_ODD"
	enumLookup[TEXTURE_GEN_S] = "TEXTURE_GEN_S"
	enumLookup[MATRIX11_ARB] = "MATRIX11_ARB"
	enumLookup[FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT] = "FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT"
	enumLookup[TRANSFORM_FEEDBACK_ACTIVE] = "TRANSFORM_FEEDBACK_ACTIVE"
	enumLookup[CONSTANT] = "CONSTANT"
	enumLookup[PRIMITIVES_GENERATED_EXT] = "PRIMITIVES_GENERATED_EXT"
	enumLookup[CONSTANT_BORDER] = "CONSTANT_BORDER"
	enumLookup[YCBCR_MESA] = "YCBCR_MESA"
	enumLookup[COLOR_ATTACHMENT7_EXT] = "COLOR_ATTACHMENT7_EXT"
	enumLookup[TESS_CONTROL_SUBROUTINE_UNIFORM] = "TESS_CONTROL_SUBROUTINE_UNIFORM"
	enumLookup[LUMINANCE4_EXT] = "LUMINANCE4_EXT"
	enumLookup[MODELVIEW4_ARB] = "MODELVIEW4_ARB"
	enumLookup[DT_SCALE_NV] = "DT_SCALE_NV"
	enumLookup[SAMPLER_3D_ARB] = "SAMPLER_3D_ARB"
	enumLookup[SAMPLE_MASK_NV] = "SAMPLE_MASK_NV"
	enumLookup[UNSIGNED_INT_2_10_10_10_REV] = "UNSIGNED_INT_2_10_10_10_REV"
	enumLookup[EYE_DISTANCE_TO_POINT_SGIS] = "EYE_DISTANCE_TO_POINT_SGIS"
	enumLookup[EYE_LINE_SGIS] = "EYE_LINE_SGIS"
	enumLookup[VERTEX_ATTRIB_RELATIVE_OFFSET] = "VERTEX_ATTRIB_RELATIVE_OFFSET"
	enumLookup[TANGENT_ARRAY_STRIDE_EXT] = "TANGENT_ARRAY_STRIDE_EXT"
	enumLookup[SOURCE1_RGB_EXT] = "SOURCE1_RGB_EXT"
	enumLookup[UNIFORM_BLOCK_NAME_LENGTH] = "UNIFORM_BLOCK_NAME_LENGTH"
	enumLookup[PROGRAM_PIPELINE_OBJECT_EXT] = "PROGRAM_PIPELINE_OBJECT_EXT"
	enumLookup[MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS] = "MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS"
	enumLookup[TEXTURE_COORD_ARRAY_STRIDE] = "TEXTURE_COORD_ARRAY_STRIDE"
	enumLookup[OBJECT_LINE_SGIS] = "OBJECT_LINE_SGIS"
	enumLookup[SURFACE_REGISTERED_NV] = "SURFACE_REGISTERED_NV"
	enumLookup[PROGRAM_ERROR_STRING_NV] = "PROGRAM_ERROR_STRING_NV"
	enumLookup[ATTACHED_SHADERS] = "ATTACHED_SHADERS"
	enumLookup[POINT_SIZE_ARRAY_OES] = "POINT_SIZE_ARRAY_OES"
	enumLookup[IMAGE_CUBE_MAP_ARRAY] = "IMAGE_CUBE_MAP_ARRAY"
	enumLookup[UNSIGNED_INT_IMAGE_BUFFER] = "UNSIGNED_INT_IMAGE_BUFFER"
	enumLookup[MINMAX_FORMAT] = "MINMAX_FORMAT"
	enumLookup[QUAD_LUMINANCE8_SGIS] = "QUAD_LUMINANCE8_SGIS"
	enumLookup[IMAGE_COMPATIBILITY_CLASS] = "IMAGE_COMPATIBILITY_CLASS"
	enumLookup[COMPRESSED_RGBA_FXT1_3DFX] = "COMPRESSED_RGBA_FXT1_3DFX"
	enumLookup[REG_19_ATI] = "REG_19_ATI"
	enumLookup[RGBA16I_EXT] = "RGBA16I_EXT"
	enumLookup[DUP_LAST_CUBIC_CURVE_TO_NV] = "DUP_LAST_CUBIC_CURVE_TO_NV"
	enumLookup[FLAT] = "FLAT"
	enumLookup[MODULATE] = "MODULATE"
	enumLookup[MAX_CONVOLUTION_HEIGHT] = "MAX_CONVOLUTION_HEIGHT"
	enumLookup[DEPTH_COMPONENT32_ARB] = "DEPTH_COMPONENT32_ARB"
	enumLookup[PARALLEL_ARRAYS_INTEL] = "PARALLEL_ARRAYS_INTEL"
	enumLookup[VERTEX_ATTRIB_ARRAY_POINTER] = "VERTEX_ATTRIB_ARRAY_POINTER"
	enumLookup[UNSIGNED_INT16_VEC2_NV] = "UNSIGNED_INT16_VEC2_NV"
	enumLookup[VIBRANCE_SCALE_NV] = "VIBRANCE_SCALE_NV"
	enumLookup[RGBA32I_EXT] = "RGBA32I_EXT"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_ARRAY] = "UNSIGNED_INT_SAMPLER_2D_ARRAY"
	enumLookup[UNSIGNED_INT8_NV] = "UNSIGNED_INT8_NV"
	enumLookup[INT_IMAGE_1D_EXT] = "INT_IMAGE_1D_EXT"
	enumLookup[CONTEXT_COMPATIBILITY_PROFILE_BIT] = "CONTEXT_COMPATIBILITY_PROFILE_BIT"
	enumLookup[EXT_vertex_array] = "EXT_vertex_array"
	enumLookup[SGIX_clipmap] = "SGIX_clipmap"
	enumLookup[GEOMETRY_INPUT_TYPE] = "GEOMETRY_INPUT_TYPE"
	enumLookup[ALWAYS] = "ALWAYS"
	enumLookup[CONVOLUTION_WIDTH_EXT] = "CONVOLUTION_WIDTH_EXT"
	enumLookup[FLOAT_MAT3] = "FLOAT_MAT3"
	enumLookup[SUB_ATI] = "SUB_ATI"
	enumLookup[FIELD_UPPER_NV] = "FIELD_UPPER_NV"
	enumLookup[TEXTURE_DEPTH_TYPE] = "TEXTURE_DEPTH_TYPE"
	enumLookup[MAX_SAMPLES] = "MAX_SAMPLES"
	enumLookup[R16_SNORM] = "R16_SNORM"
	enumLookup[ALLOW_DRAW_FRG_HINT_PGI] = "ALLOW_DRAW_FRG_HINT_PGI"
	enumLookup[ZOOM_X] = "ZOOM_X"
	enumLookup[TEXTURE_NORMAL_EXT] = "TEXTURE_NORMAL_EXT"
	enumLookup[CON_28_ATI] = "CON_28_ATI"
	enumLookup[COMPRESSED_RGBA_ASTC_4x4_KHR] = "COMPRESSED_RGBA_ASTC_4x4_KHR"
	enumLookup[MAX_ASYNC_HISTOGRAM_SGIX] = "MAX_ASYNC_HISTOGRAM_SGIX"
	enumLookup[OUTPUT_TEXTURE_COORD1_EXT] = "OUTPUT_TEXTURE_COORD1_EXT"
	enumLookup[DRAW_BUFFER9_ARB] = "DRAW_BUFFER9_ARB"
	enumLookup[PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB] = "PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB"
	enumLookup[MATRIX20_ARB] = "MATRIX20_ARB"
	enumLookup[CON_1_ATI] = "CON_1_ATI"
	enumLookup[BGR_INTEGER] = "BGR_INTEGER"
	enumLookup[LINEAR_CLIPMAP_LINEAR_SGIX] = "LINEAR_CLIPMAP_LINEAR_SGIX"
	enumLookup[ASYNC_TEX_IMAGE_SGIX] = "ASYNC_TEX_IMAGE_SGIX"
	enumLookup[TEXTURE_DT_SIZE_NV] = "TEXTURE_DT_SIZE_NV"
	enumLookup[CONSERVE_MEMORY_HINT_PGI] = "CONSERVE_MEMORY_HINT_PGI"
	enumLookup[TRANSPOSE_PROJECTION_MATRIX] = "TRANSPOSE_PROJECTION_MATRIX"
	enumLookup[CLAMP_VERTEX_COLOR] = "CLAMP_VERTEX_COLOR"
	enumLookup[CIRCULAR_CW_ARC_TO_NV] = "CIRCULAR_CW_ARC_TO_NV"
	enumLookup[ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER] = "ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER"
	enumLookup[AUX1] = "AUX1"
	enumLookup[DEPTH_COMPONENT] = "DEPTH_COMPONENT"
	enumLookup[PROGRAM_PIPELINE_BINDING] = "PROGRAM_PIPELINE_BINDING"
	enumLookup[DRAW_INDIRECT_UNIFIED_NV] = "DRAW_INDIRECT_UNIFIED_NV"
	enumLookup[OPERAND2_RGB] = "OPERAND2_RGB"
	enumLookup[CURRENT_WEIGHT_ARB] = "CURRENT_WEIGHT_ARB"
	enumLookup[EVAL_VERTEX_ATTRIB1_NV] = "EVAL_VERTEX_ATTRIB1_NV"
	enumLookup[MATRIX23_ARB] = "MATRIX23_ARB"
	enumLookup[FRONT_AND_BACK] = "FRONT_AND_BACK"
	enumLookup[POST_CONVOLUTION_GREEN_BIAS] = "POST_CONVOLUTION_GREEN_BIAS"
	enumLookup[ELEMENT_ARRAY_TYPE_ATI] = "ELEMENT_ARRAY_TYPE_ATI"
	enumLookup[OP_POWER_EXT] = "OP_POWER_EXT"
	enumLookup[DOT2_ADD_ATI] = "DOT2_ADD_ATI"
	enumLookup[MAX_BINDABLE_UNIFORM_SIZE_EXT] = "MAX_BINDABLE_UNIFORM_SIZE_EXT"
	enumLookup[MOVE_TO_RESETS_NV] = "MOVE_TO_RESETS_NV"
	enumLookup[SGIX_texture_multi_buffer] = "SGIX_texture_multi_buffer"
	enumLookup[OUTPUT_TEXTURE_COORD21_EXT] = "OUTPUT_TEXTURE_COORD21_EXT"
	enumLookup[INTENSITY_FLOAT32_ATI] = "INTENSITY_FLOAT32_ATI"
	enumLookup[DRAW_BUFFER6_NV] = "DRAW_BUFFER6_NV"
	enumLookup[SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT] = "SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT"
	enumLookup[POINT_SMOOTH] = "POINT_SMOOTH"
	enumLookup[DEPTH_COMPONENTS] = "DEPTH_COMPONENTS"
	enumLookup[SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST] = "SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST"
	enumLookup[PROGRAM_PARAMETER_NV] = "PROGRAM_PARAMETER_NV"
	enumLookup[MAP1_VERTEX_ATTRIB3_4_NV] = "MAP1_VERTEX_ATTRIB3_4_NV"
	enumLookup[STATIC_READ_ARB] = "STATIC_READ_ARB"
	enumLookup[SGIX_interlace] = "SGIX_interlace"
	enumLookup[VIEWPORT] = "VIEWPORT"
	enumLookup[GL_3_BYTES] = "GL_3_BYTES"
	enumLookup[RGBA_S3TC] = "RGBA_S3TC"
	enumLookup[UNDEFINED_APPLE] = "UNDEFINED_APPLE"
	enumLookup[INT_IMAGE_1D] = "INT_IMAGE_1D"
	enumLookup[AUX3] = "AUX3"
	enumLookup[UNPACK_ALIGNMENT] = "UNPACK_ALIGNMENT"
	enumLookup[MULTISAMPLE_3DFX] = "MULTISAMPLE_3DFX"
	enumLookup[STATIC_ATI] = "STATIC_ATI"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_BINDING_NV] = "TRANSFORM_FEEDBACK_BUFFER_BINDING_NV"
	enumLookup[COLOR_ATTACHMENT7] = "COLOR_ATTACHMENT7"
	enumLookup[COLOR_TABLE_RED_SIZE] = "COLOR_TABLE_RED_SIZE"
	enumLookup[SGX_PROGRAM_BINARY_IMG] = "SGX_PROGRAM_BINARY_IMG"
	enumLookup[TEXTURE_IMAGE_TYPE] = "TEXTURE_IMAGE_TYPE"
	enumLookup[SGIX_blend_alpha_minmax] = "SGIX_blend_alpha_minmax"
	enumLookup[POLYGON_OFFSET_FILL] = "POLYGON_OFFSET_FILL"
	enumLookup[LIGHT0] = "LIGHT0"
	enumLookup[LUMINANCE_ALPHA16F_ARB] = "LUMINANCE_ALPHA16F_ARB"
	enumLookup[ALPHA4_EXT] = "ALPHA4_EXT"
	enumLookup[TEXTURE_COLOR_WRITEMASK_SGIS] = "TEXTURE_COLOR_WRITEMASK_SGIS"
	enumLookup[PIXEL_UNPACK_BUFFER_ARB] = "PIXEL_UNPACK_BUFFER_ARB"
	enumLookup[RENDERBUFFER_GREEN_SIZE] = "RENDERBUFFER_GREEN_SIZE"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR"
	enumLookup[NATIVE_GRAPHICS_END_HINT_PGI] = "NATIVE_GRAPHICS_END_HINT_PGI"
	enumLookup[LUMINANCE6_ALPHA2_EXT] = "LUMINANCE6_ALPHA2_EXT"
	enumLookup[COLOR_INDEX12_EXT] = "COLOR_INDEX12_EXT"
	enumLookup[VERTEX_TEXTURE] = "VERTEX_TEXTURE"
	enumLookup[MIRROR_CLAMP_EXT] = "MIRROR_CLAMP_EXT"
	enumLookup[COMPRESSED_SLUMINANCE_ALPHA] = "COMPRESSED_SLUMINANCE_ALPHA"
	enumLookup[UNIFORM_BARRIER_BIT] = "UNIFORM_BARRIER_BIT"
	enumLookup[FOG_COORDINATE_ARRAY_STRIDE_EXT] = "FOG_COORDINATE_ARRAY_STRIDE_EXT"
	enumLookup[ALPHA_FLOAT32_APPLE] = "ALPHA_FLOAT32_APPLE"
	enumLookup[FORCE_BLUE_TO_ONE_NV] = "FORCE_BLUE_TO_ONE_NV"
	enumLookup[REG_25_ATI] = "REG_25_ATI"
	enumLookup[UNSIGNED_INT_SAMPLER_3D_EXT] = "UNSIGNED_INT_SAMPLER_3D_EXT"
	enumLookup[PIXEL_TEX_GEN_ALPHA_MS_SGIX] = "PIXEL_TEX_GEN_ALPHA_MS_SGIX"
	enumLookup[CULL_VERTEX_EYE_POSITION_EXT] = "CULL_VERTEX_EYE_POSITION_EXT"
	enumLookup[DEBUG_SOURCE_API_ARB] = "DEBUG_SOURCE_API_ARB"
	enumLookup[SCREEN_COORDINATES_REND] = "SCREEN_COORDINATES_REND"
	enumLookup[COMPRESSED_LUMINANCE] = "COMPRESSED_LUMINANCE"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB] = "TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB"
	enumLookup[NOOP] = "NOOP"
	enumLookup[UNSIGNED_SHORT_5_5_5_1_EXT] = "UNSIGNED_SHORT_5_5_5_1_EXT"
	enumLookup[COLOR_ARRAY_POINTER_EXT] = "COLOR_ARRAY_POINTER_EXT"
	enumLookup[RG] = "RG"
	enumLookup[DEPTH_PASS_INSTRUMENT_SGIX] = "DEPTH_PASS_INSTRUMENT_SGIX"
	enumLookup[DSDT_NV] = "DSDT_NV"
	enumLookup[ALPHA32F_ARB] = "ALPHA32F_ARB"
	enumLookup[BEVEL_NV] = "BEVEL_NV"
	enumLookup[STATIC_READ] = "STATIC_READ"
	enumLookup[MAX_LIST_NESTING] = "MAX_LIST_NESTING"
	enumLookup[GL_2PASS_1_SGIS] = "GL_2PASS_1_SGIS"
	enumLookup[IMAGE_ROTATE_ANGLE_HP] = "IMAGE_ROTATE_ANGLE_HP"
	enumLookup[FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX] = "FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX"
	enumLookup[BACK_NORMALS_HINT_PGI] = "BACK_NORMALS_HINT_PGI"
	enumLookup[LINEAR_DETAIL_ALPHA_SGIS] = "LINEAR_DETAIL_ALPHA_SGIS"
	enumLookup[TEXTURE28_ARB] = "TEXTURE28_ARB"
	enumLookup[REG_15_ATI] = "REG_15_ATI"
	enumLookup[PIXEL_MAP_I_TO_I_SIZE] = "PIXEL_MAP_I_TO_I_SIZE"
	enumLookup[FLOAT_MAT4x2] = "FLOAT_MAT4x2"
	enumLookup[TEXTURE_INTENSITY_TYPE_ARB] = "TEXTURE_INTENSITY_TYPE_ARB"
	enumLookup[SLUMINANCE8_EXT] = "SLUMINANCE8_EXT"
	enumLookup[RELATIVE_HORIZONTAL_LINE_TO_NV] = "RELATIVE_HORIZONTAL_LINE_TO_NV"
	enumLookup[INT] = "INT"
	enumLookup[EVAL_VERTEX_ATTRIB11_NV] = "EVAL_VERTEX_ATTRIB11_NV"
	enumLookup[TEXTURE_BUFFER_ARB] = "TEXTURE_BUFFER_ARB"
	enumLookup[MAX_COMBINED_IMAGE_UNIFORMS] = "MAX_COMBINED_IMAGE_UNIFORMS"
	enumLookup[FOG_FUNC_POINTS_SGIS] = "FOG_FUNC_POINTS_SGIS"
	enumLookup[TEXTURE29_ARB] = "TEXTURE29_ARB"
	enumLookup[WIDE_LINE_HINT_PGI] = "WIDE_LINE_HINT_PGI"
	enumLookup[BLEND] = "BLEND"
	enumLookup[COMBINER_CD_OUTPUT_NV] = "COMBINER_CD_OUTPUT_NV"
	enumLookup[VERTEX_SHADER_OPTIMIZED_EXT] = "VERTEX_SHADER_OPTIMIZED_EXT"
	enumLookup[MAX_PROGRAM_LOCAL_PARAMETERS_ARB] = "MAX_PROGRAM_LOCAL_PARAMETERS_ARB"
	enumLookup[VERTEX_ARRAY_OBJECT_EXT] = "VERTEX_ARRAY_OBJECT_EXT"
	enumLookup[MAP_UNSYNCHRONIZED_BIT_EXT] = "MAP_UNSYNCHRONIZED_BIT_EXT"
	enumLookup[INCR_WRAP_EXT] = "INCR_WRAP_EXT"
	enumLookup[INVERSE_NV] = "INVERSE_NV"
	enumLookup[CON_18_ATI] = "CON_18_ATI"
	enumLookup[TEXTURE_TYPE_QCOM] = "TEXTURE_TYPE_QCOM"
	enumLookup[COMPRESSED_SRGB_ALPHA_EXT] = "COMPRESSED_SRGB_ALPHA_EXT"
	enumLookup[UNSIGNED_INT_VEC4] = "UNSIGNED_INT_VEC4"
	enumLookup[DOUBLE_MAT3x2] = "DOUBLE_MAT3x2"
	enumLookup[UNSIGNED_INT16_VEC3_NV] = "UNSIGNED_INT16_VEC3_NV"
	enumLookup[SGIX_pixel_texture] = "SGIX_pixel_texture"
	enumLookup[PIXEL_TEX_GEN_SGIX] = "PIXEL_TEX_GEN_SGIX"
	enumLookup[FLOAT_MAT2x3] = "FLOAT_MAT2x3"
	enumLookup[VIDEO_BUFFER_BINDING_NV] = "VIDEO_BUFFER_BINDING_NV"
	enumLookup[MATRIX_MODE] = "MATRIX_MODE"
	enumLookup[OR_REVERSE] = "OR_REVERSE"
	enumLookup[PROGRAM_BINARY_FORMATS_OES] = "PROGRAM_BINARY_FORMATS_OES"
	enumLookup[MAX_DEBUG_MESSAGE_LENGTH] = "MAX_DEBUG_MESSAGE_LENGTH"
	enumLookup[ONE] = "ONE"
	enumLookup[TEXTURE_RECTANGLE_ARB] = "TEXTURE_RECTANGLE_ARB"
	enumLookup[OUTPUT_TEXTURE_COORD16_EXT] = "OUTPUT_TEXTURE_COORD16_EXT"
	enumLookup[IMAGE_2D] = "IMAGE_2D"
	enumLookup[MITER_TRUNCATE_NV] = "MITER_TRUNCATE_NV"
	enumLookup[GEOMETRY_DEFORMATION_BIT_SGIX] = "GEOMETRY_DEFORMATION_BIT_SGIX"
	enumLookup[NAND] = "NAND"
	enumLookup[IMAGE_CLASS_10_10_10_2] = "IMAGE_CLASS_10_10_10_2"
	enumLookup[LUMINANCE_FLOAT16_APPLE] = "LUMINANCE_FLOAT16_APPLE"
	enumLookup[MAX_COMPUTE_WORK_GROUP_COUNT] = "MAX_COMPUTE_WORK_GROUP_COUNT"
	enumLookup[TEXTURE_MAX_LOD_SGIS] = "TEXTURE_MAX_LOD_SGIS"
	enumLookup[MAP1_VERTEX_ATTRIB15_4_NV] = "MAP1_VERTEX_ATTRIB15_4_NV"
	enumLookup[FLOAT16_VEC3_NV] = "FLOAT16_VEC3_NV"
	enumLookup[PROGRAM_OUTPUT] = "PROGRAM_OUTPUT"
	enumLookup[ALPHA_MIN_SGIX] = "ALPHA_MIN_SGIX"
	enumLookup[CURRENT_MATRIX_STACK_DEPTH_ARB] = "CURRENT_MATRIX_STACK_DEPTH_ARB"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV] = "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV"
	enumLookup[UNPACK_COMPRESSED_BLOCK_SIZE] = "UNPACK_COMPRESSED_BLOCK_SIZE"
	enumLookup[SCISSOR_TEST] = "SCISSOR_TEST"
	enumLookup[INTENSITY8] = "INTENSITY8"
	enumLookup[NUM_FRAGMENT_REGISTERS_ATI] = "NUM_FRAGMENT_REGISTERS_ATI"
	enumLookup[POINT_SIZE_ARRAY_TYPE_OES] = "POINT_SIZE_ARRAY_TYPE_OES"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR"
	enumLookup[EXT_copy_texture] = "EXT_copy_texture"
	enumLookup[INT16_NV] = "INT16_NV"
	enumLookup[TRIANGLES] = "TRIANGLES"
	enumLookup[MAX_3D_TEXTURE_SIZE] = "MAX_3D_TEXTURE_SIZE"
	enumLookup[COLOR_TABLE_WIDTH_SGI] = "COLOR_TABLE_WIDTH_SGI"
	enumLookup[RGB_FLOAT32_APPLE] = "RGB_FLOAT32_APPLE"
	enumLookup[BUFFER_UPDATE_BARRIER_BIT] = "BUFFER_UPDATE_BARRIER_BIT"
	enumLookup[RGBA8_EXT] = "RGBA8_EXT"
	enumLookup[TEXTURE_LIGHT_EXT] = "TEXTURE_LIGHT_EXT"
	enumLookup[TEXTURE4_ARB] = "TEXTURE4_ARB"
	enumLookup[COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB] = "COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB"
	enumLookup[NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV] = "NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV"
	enumLookup[DST_ALPHA] = "DST_ALPHA"
	enumLookup[GL_4PASS_0_EXT] = "GL_4PASS_0_EXT"
	enumLookup[BLEND_SRC_ALPHA] = "BLEND_SRC_ALPHA"
	enumLookup[SPRITE_AXIAL_SGIX] = "SPRITE_AXIAL_SGIX"
	enumLookup[NEAREST_CLIPMAP_LINEAR_SGIX] = "NEAREST_CLIPMAP_LINEAR_SGIX"
	enumLookup[INTERNALFORMAT_GREEN_TYPE] = "INTERNALFORMAT_GREEN_TYPE"
	enumLookup[STORAGE_SHARED_APPLE] = "STORAGE_SHARED_APPLE"
	enumLookup[STENCIL_BACK_FUNC] = "STENCIL_BACK_FUNC"
	enumLookup[SKIP_COMPONENTS3_NV] = "SKIP_COMPONENTS3_NV"
	enumLookup[SECONDARY_COLOR_ARRAY_LENGTH_NV] = "SECONDARY_COLOR_ARRAY_LENGTH_NV"
	enumLookup[LESS] = "LESS"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_X_EXT] = "TEXTURE_CUBE_MAP_NEGATIVE_X_EXT"
	enumLookup[MODELVIEW9_ARB] = "MODELVIEW9_ARB"
	enumLookup[PROGRAM_NATIVE_INSTRUCTIONS_ARB] = "PROGRAM_NATIVE_INSTRUCTIONS_ARB"
	enumLookup[SGIX_fragment_lighting] = "SGIX_fragment_lighting"
	enumLookup[TEXTURE_GEN_R] = "TEXTURE_GEN_R"
	enumLookup[TEXTURE_GEN_MODE] = "TEXTURE_GEN_MODE"
	enumLookup[CONVOLUTION_2D_EXT] = "CONVOLUTION_2D_EXT"
	enumLookup[ACTIVE_PROGRAM_EXT] = "ACTIVE_PROGRAM_EXT"
	enumLookup[MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT] = "MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT"
	enumLookup[DOUBLE_MAT3] = "DOUBLE_MAT3"
	enumLookup[UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX] = "UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX"
	enumLookup[RG8UI] = "RG8UI"
	enumLookup[TEXTURE_MAX_CLAMP_S_SGIX] = "TEXTURE_MAX_CLAMP_S_SGIX"
	enumLookup[DOT_PRODUCT_NV] = "DOT_PRODUCT_NV"
	enumLookup[STENCIL_BACK_FAIL] = "STENCIL_BACK_FAIL"
	enumLookup[PATCH_VERTICES] = "PATCH_VERTICES"
	enumLookup[FONT_X_MIN_BOUNDS_BIT_NV] = "FONT_X_MIN_BOUNDS_BIT_NV"
	enumLookup[TEXTURE_CUBE_MAP_ARB] = "TEXTURE_CUBE_MAP_ARB"
	enumLookup[WRITE_ONLY] = "WRITE_ONLY"
	enumLookup[GEOMETRY_VERTICES_OUT_ARB] = "GEOMETRY_VERTICES_OUT_ARB"
	enumLookup[ATOMIC_COUNTER_BARRIER_BIT] = "ATOMIC_COUNTER_BARRIER_BIT"
	enumLookup[SAMPLE_ALPHA_TO_ONE] = "SAMPLE_ALPHA_TO_ONE"
	enumLookup[TESS_GEN_VERTEX_ORDER] = "TESS_GEN_VERTEX_ORDER"
	enumLookup[FUNC_REVERSE_SUBTRACT] = "FUNC_REVERSE_SUBTRACT"
	enumLookup[LOCAL_CONSTANT_VALUE_EXT] = "LOCAL_CONSTANT_VALUE_EXT"
	enumLookup[DEPTH_COMPONENT32F_NV] = "DEPTH_COMPONENT32F_NV"
	enumLookup[QUERY_BY_REGION_WAIT] = "QUERY_BY_REGION_WAIT"
	enumLookup[RELATIVE_LARGE_CCW_ARC_TO_NV] = "RELATIVE_LARGE_CCW_ARC_TO_NV"
	enumLookup[SHADER_STORAGE_BUFFER] = "SHADER_STORAGE_BUFFER"
	enumLookup[UNSIGNALED_APPLE] = "UNSIGNALED_APPLE"
	enumLookup[ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES] = "ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES"
	enumLookup[OBJECT_ACTIVE_ATTRIBUTES_ARB] = "OBJECT_ACTIVE_ATTRIBUTES_ARB"
	enumLookup[IMAGE_CUBE_EXT] = "IMAGE_CUBE_EXT"
	enumLookup[MODELVIEW] = "MODELVIEW"
	enumLookup[BGR_EXT] = "BGR_EXT"
	enumLookup[MAX_RECTANGLE_TEXTURE_SIZE_ARB] = "MAX_RECTANGLE_TEXTURE_SIZE_ARB"
	enumLookup[VIDEO_BUFFER_PITCH_NV] = "VIDEO_BUFFER_PITCH_NV"
	enumLookup[MAX_FRAGMENT_IMAGE_UNIFORMS] = "MAX_FRAGMENT_IMAGE_UNIFORMS"
	enumLookup[DEPTH_EXT] = "DEPTH_EXT"
	enumLookup[LIGHT6] = "LIGHT6"
	enumLookup[FUNC_REVERSE_SUBTRACT_OES] = "FUNC_REVERSE_SUBTRACT_OES"
	enumLookup[LUMINANCE32F_ARB] = "LUMINANCE32F_ARB"
	enumLookup[CON_7_ATI] = "CON_7_ATI"
	enumLookup[COLOR_FLOAT_APPLE] = "COLOR_FLOAT_APPLE"
	enumLookup[Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV] = "Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV"
	enumLookup[INT_IMAGE_2D_RECT] = "INT_IMAGE_2D_RECT"
	enumLookup[FRAGMENT_SHADER_BIT] = "FRAGMENT_SHADER_BIT"
	enumLookup[DOT_PRODUCT_DEPTH_REPLACE_NV] = "DOT_PRODUCT_DEPTH_REPLACE_NV"
	enumLookup[BUFFER_MAPPED] = "BUFFER_MAPPED"
	enumLookup[INVALID_VALUE] = "INVALID_VALUE"
	enumLookup[MINMAX_FORMAT_EXT] = "MINMAX_FORMAT_EXT"
	enumLookup[SAMPLES_ARB] = "SAMPLES_ARB"
	enumLookup[TIME_ELAPSED_EXT] = "TIME_ELAPSED_EXT"
	enumLookup[RESAMPLE_DECIMATE_OML] = "RESAMPLE_DECIMATE_OML"
	enumLookup[MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS] = "MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS"
	enumLookup[INT_SAMPLER_BUFFER_EXT] = "INT_SAMPLER_BUFFER_EXT"
	enumLookup[DOUBLE_MAT4x3] = "DOUBLE_MAT4x3"
	enumLookup[UTF16_NV] = "UTF16_NV"
	enumLookup[VIRTUAL_PAGE_SIZE_Y_AMD] = "VIRTUAL_PAGE_SIZE_Y_AMD"
	enumLookup[MOV_ATI] = "MOV_ATI"
	enumLookup[HIGH_INT] = "HIGH_INT"
	enumLookup[MAX_VERTEX_IMAGE_UNIFORMS] = "MAX_VERTEX_IMAGE_UNIFORMS"
	enumLookup[MAX_DEBUG_LOGGED_MESSAGES_AMD] = "MAX_DEBUG_LOGGED_MESSAGES_AMD"
	enumLookup[PROXY_TEXTURE_2D] = "PROXY_TEXTURE_2D"
	enumLookup[SAMPLE_MASK_VALUE_EXT] = "SAMPLE_MASK_VALUE_EXT"
	enumLookup[TEXTURE13] = "TEXTURE13"
	enumLookup[MAX_VERTEX_UNITS_ARB] = "MAX_VERTEX_UNITS_ARB"
	enumLookup[VERTEX_ATTRIB_ARRAY_DIVISOR] = "VERTEX_ATTRIB_ARRAY_DIVISOR"
	enumLookup[FLOAT16_VEC2_NV] = "FLOAT16_VEC2_NV"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE] = "UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE"
	enumLookup[VERTEX_ARRAY_SIZE_EXT] = "VERTEX_ARRAY_SIZE_EXT"
	enumLookup[INSTRUMENT_MEASUREMENTS_SGIX] = "INSTRUMENT_MEASUREMENTS_SGIX"
	enumLookup[RGB9_E5] = "RGB9_E5"
	enumLookup[INT64_VEC4_NV] = "INT64_VEC4_NV"
	enumLookup[BLEND_EQUATION_RGB_OES] = "BLEND_EQUATION_RGB_OES"
	enumLookup[UNDEFINED_VERTEX] = "UNDEFINED_VERTEX"
	enumLookup[PROGRAM_STRING_ARB] = "PROGRAM_STRING_ARB"
	enumLookup[PROGRAM_NATIVE_ATTRIBS_ARB] = "PROGRAM_NATIVE_ATTRIBS_ARB"
	enumLookup[LUMINANCE_ALPHA16UI_EXT] = "LUMINANCE_ALPHA16UI_EXT"
	enumLookup[UNSIGNED_INT64_VEC3_NV] = "UNSIGNED_INT64_VEC3_NV"
	enumLookup[MAX_VERTEX_ATOMIC_COUNTER_BUFFERS] = "MAX_VERTEX_ATOMIC_COUNTER_BUFFERS"
	enumLookup[MAP1_GRID_SEGMENTS] = "MAP1_GRID_SEGMENTS"
	enumLookup[VERTEX_PROGRAM_POINT_SIZE] = "VERTEX_PROGRAM_POINT_SIZE"
	enumLookup[OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV] = "OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV"
	enumLookup[LINE_STIPPLE_PATTERN] = "LINE_STIPPLE_PATTERN"
	enumLookup[ASYNC_DRAW_PIXELS_SGIX] = "ASYNC_DRAW_PIXELS_SGIX"
	enumLookup[R11F_G11F_B10F] = "R11F_G11F_B10F"
	enumLookup[FRAMEBUFFER_EXT] = "FRAMEBUFFER_EXT"
	enumLookup[FOG_COORD_ARRAY_ADDRESS_NV] = "FOG_COORD_ARRAY_ADDRESS_NV"
	enumLookup[OBJECT_TYPE_APPLE] = "OBJECT_TYPE_APPLE"
	enumLookup[MATRIX_INDEX_ARRAY_STRIDE_OES] = "MATRIX_INDEX_ARRAY_STRIDE_OES"
	enumLookup[PIXEL_UNPACK_BUFFER_BINDING] = "PIXEL_UNPACK_BUFFER_BINDING"
	enumLookup[TRANSLATE_2D_NV] = "TRANSLATE_2D_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY14_NV] = "VERTEX_ATTRIB_ARRAY14_NV"
	enumLookup[OUTPUT_VERTEX_EXT] = "OUTPUT_VERTEX_EXT"
	enumLookup[FRAMEBUFFER_INCOMPLETE_READ_BUFFER] = "FRAMEBUFFER_INCOMPLETE_READ_BUFFER"
	enumLookup[SOURCE0_ALPHA_ARB] = "SOURCE0_ALPHA_ARB"
	enumLookup[EVAL_VERTEX_ATTRIB6_NV] = "EVAL_VERTEX_ATTRIB6_NV"
	enumLookup[QUERY_RESULT] = "QUERY_RESULT"
	enumLookup[RESAMPLE_REPLICATE_OML] = "RESAMPLE_REPLICATE_OML"
	enumLookup[TEXTURE_WRAP_S] = "TEXTURE_WRAP_S"
	enumLookup[TEXTURE_COORD_ARRAY_BUFFER_BINDING] = "TEXTURE_COORD_ARRAY_BUFFER_BINDING"
	enumLookup[TRANSPOSE_AFFINE_2D_NV] = "TRANSPOSE_AFFINE_2D_NV"
	enumLookup[STENCIL_PASS_DEPTH_PASS] = "STENCIL_PASS_DEPTH_PASS"
	enumLookup[PACK_IMAGE_HEIGHT_EXT] = "PACK_IMAGE_HEIGHT_EXT"
	enumLookup[COLOR_TABLE_SCALE_SGI] = "COLOR_TABLE_SCALE_SGI"
	enumLookup[POINT_SIZE_MAX_SGIS] = "POINT_SIZE_MAX_SGIS"
	enumLookup[UNPACK_CONSTANT_DATA_SUNX] = "UNPACK_CONSTANT_DATA_SUNX"
	enumLookup[MAX_RENDERBUFFER_SIZE] = "MAX_RENDERBUFFER_SIZE"
	enumLookup[COMPRESSED_LUMINANCE_ALPHA_ARB] = "COMPRESSED_LUMINANCE_ALPHA_ARB"
	enumLookup[COMBINE_ALPHA_EXT] = "COMBINE_ALPHA_EXT"
	enumLookup[MAT_EMISSION_BIT_PGI] = "MAT_EMISSION_BIT_PGI"
	enumLookup[SGIX_async_histogram] = "SGIX_async_histogram"
	enumLookup[TEXTURE_BIT] = "TEXTURE_BIT"
	enumLookup[LINE_WIDTH_GRANULARITY] = "LINE_WIDTH_GRANULARITY"
	enumLookup[POINT_SPRITE_ARB] = "POINT_SPRITE_ARB"
	enumLookup[ADD_BLEND_IMG] = "ADD_BLEND_IMG"
	enumLookup[LUMINANCE32UI_EXT] = "LUMINANCE32UI_EXT"
	enumLookup[MAX_NUM_ACTIVE_VARIABLES] = "MAX_NUM_ACTIVE_VARIABLES"
	enumLookup[T2F_C4UB_V3F] = "T2F_C4UB_V3F"
	enumLookup[LUMINANCE8_EXT] = "LUMINANCE8_EXT"
	enumLookup[INTENSITY12_EXT] = "INTENSITY12_EXT"
	enumLookup[Z_EXT] = "Z_EXT"
	enumLookup[TEXTURE_STACK_DEPTH] = "TEXTURE_STACK_DEPTH"
	enumLookup[TEXTURE_BINDING_1D] = "TEXTURE_BINDING_1D"
	enumLookup[CLEAR] = "CLEAR"
	enumLookup[MATRIX29_ARB] = "MATRIX29_ARB"
	enumLookup[EXP2] = "EXP2"
	enumLookup[PIXEL_MAP_I_TO_B] = "PIXEL_MAP_I_TO_B"
	enumLookup[RESCALE_NORMAL_EXT] = "RESCALE_NORMAL_EXT"
	enumLookup[COMPRESSED_SIGNED_RED_RGTC1_EXT] = "COMPRESSED_SIGNED_RED_RGTC1_EXT"
	enumLookup[UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV] = "UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV"
	enumLookup[SMOOTH_POINT_SIZE_RANGE] = "SMOOTH_POINT_SIZE_RANGE"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Z] = "TEXTURE_CUBE_MAP_NEGATIVE_Z"
	enumLookup[SEPARATE_ATTRIBS] = "SEPARATE_ATTRIBS"
	enumLookup[HISTOGRAM_LUMINANCE_SIZE] = "HISTOGRAM_LUMINANCE_SIZE"
	enumLookup[TEXTURE_ENV_BIAS_SGIX] = "TEXTURE_ENV_BIAS_SGIX"
	enumLookup[PIXEL_TILE_CACHE_INCREMENT_SGIX] = "PIXEL_TILE_CACHE_INCREMENT_SGIX"
	enumLookup[RG32UI] = "RG32UI"
	enumLookup[CONSTANT_COLOR1_NV] = "CONSTANT_COLOR1_NV"
	enumLookup[TEXTURE_BLUE_SIZE] = "TEXTURE_BLUE_SIZE"
	enumLookup[BITMAP] = "BITMAP"
	enumLookup[VERTEX_ATTRIB_ARRAY9_NV] = "VERTEX_ATTRIB_ARRAY9_NV"
	enumLookup[OUTPUT_TEXTURE_COORD20_EXT] = "OUTPUT_TEXTURE_COORD20_EXT"
	enumLookup[MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB] = "MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB"
	enumLookup[REG_3_ATI] = "REG_3_ATI"
	enumLookup[CON_21_ATI] = "CON_21_ATI"
	enumLookup[PALETTE8_R5_G6_B5_OES] = "PALETTE8_R5_G6_B5_OES"
	enumLookup[MIN_FRAGMENT_INTERPOLATION_OFFSET] = "MIN_FRAGMENT_INTERPOLATION_OFFSET"
	enumLookup[INTERLACE_READ_INGR] = "INTERLACE_READ_INGR"
	enumLookup[Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV] = "Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV"
	enumLookup[NORMAL_ARRAY_LIST_STRIDE_IBM] = "NORMAL_ARRAY_LIST_STRIDE_IBM"
	enumLookup[PACK_CMYK_HINT_EXT] = "PACK_CMYK_HINT_EXT"
	enumLookup[CLAMP_TO_BORDER] = "CLAMP_TO_BORDER"
	enumLookup[RESAMPLE_ZERO_FILL_SGIX] = "RESAMPLE_ZERO_FILL_SGIX"
	enumLookup[TEXTURE_MAG_SIZE_NV] = "TEXTURE_MAG_SIZE_NV"
	enumLookup[MATRIX16_ARB] = "MATRIX16_ARB"
	enumLookup[MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES] = "MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES"
	enumLookup[CONTEXT_CORE_PROFILE_BIT] = "CONTEXT_CORE_PROFILE_BIT"
	enumLookup[ARRAY_OBJECT_BUFFER_ATI] = "ARRAY_OBJECT_BUFFER_ATI"
	enumLookup[BUFFER_ACCESS] = "BUFFER_ACCESS"
	enumLookup[MAX_VERTEX_UNIFORM_BLOCKS] = "MAX_VERTEX_UNIFORM_BLOCKS"
	enumLookup[RENDERBUFFER_OES] = "RENDERBUFFER_OES"
	enumLookup[ALPHA8_SNORM] = "ALPHA8_SNORM"
	enumLookup[AUX0] = "AUX0"
	enumLookup[LINE_STIPPLE] = "LINE_STIPPLE"
	enumLookup[TEXTURE_MAX_LEVEL] = "TEXTURE_MAX_LEVEL"
	enumLookup[VIEW_CLASS_RGTC2_RG] = "VIEW_CLASS_RGTC2_RG"
	enumLookup[RGB_FLOAT16_APPLE] = "RGB_FLOAT16_APPLE"
	enumLookup[TEXTURE_RESIDENT] = "TEXTURE_RESIDENT"
	enumLookup[LIGHT5] = "LIGHT5"
	enumLookup[RG_EXT] = "RG_EXT"
	enumLookup[VERTEX_ARRAY_STORAGE_HINT_APPLE] = "VERTEX_ARRAY_STORAGE_HINT_APPLE"
	enumLookup[ATTRIB_ARRAY_POINTER_NV] = "ATTRIB_ARRAY_POINTER_NV"
	enumLookup[IMPLEMENTATION_COLOR_READ_TYPE_OES] = "IMPLEMENTATION_COLOR_READ_TYPE_OES"
	enumLookup[UNSIGNED_INT_IMAGE_1D_ARRAY_EXT] = "UNSIGNED_INT_IMAGE_1D_ARRAY_EXT"
	enumLookup[TEXTURE_COORD_ARRAY_LIST_IBM] = "TEXTURE_COORD_ARRAY_LIST_IBM"
	enumLookup[UNIFORM_BUFFER_START] = "UNIFORM_BUFFER_START"
	enumLookup[LARGE_CW_ARC_TO_NV] = "LARGE_CW_ARC_TO_NV"
	enumLookup[TEXTURE_FILTER_CONTROL] = "TEXTURE_FILTER_CONTROL"
	enumLookup[PROGRAM_POINT_SIZE_EXT] = "PROGRAM_POINT_SIZE_EXT"
	enumLookup[PRESERVE_ATI] = "PRESERVE_ATI"
	enumLookup[MATRIX26_ARB] = "MATRIX26_ARB"
	enumLookup[DOUBLE_MAT4] = "DOUBLE_MAT4"
	enumLookup[MAP_INVALIDATE_BUFFER_BIT] = "MAP_INVALIDATE_BUFFER_BIT"
	enumLookup[ONE_MINUS_SRC_ALPHA] = "ONE_MINUS_SRC_ALPHA"
	enumLookup[RGB_S3TC] = "RGB_S3TC"
	enumLookup[COMPRESSED_RGBA_PVRTC_4BPPV2_IMG] = "COMPRESSED_RGBA_PVRTC_4BPPV2_IMG"
	enumLookup[UNSIGNED_BYTE_3_3_2_EXT] = "UNSIGNED_BYTE_3_3_2_EXT"
	enumLookup[LIST_PRIORITY_SGIX] = "LIST_PRIORITY_SGIX"
	enumLookup[YCRCBA_SGIX] = "YCRCBA_SGIX"
	enumLookup[RGBA16UI] = "RGBA16UI"
	enumLookup[LUMINANCE8_ALPHA8_SNORM] = "LUMINANCE8_ALPHA8_SNORM"
	enumLookup[TOP_LEVEL_ARRAY_STRIDE] = "TOP_LEVEL_ARRAY_STRIDE"
	enumLookup[SHININESS] = "SHININESS"
	enumLookup[ALPHA_MAX_SGIX] = "ALPHA_MAX_SGIX"
	enumLookup[SAMPLER_3D] = "SAMPLER_3D"
	enumLookup[FRACTIONAL_EVEN] = "FRACTIONAL_EVEN"
	enumLookup[PIXEL_MAP_I_TO_G] = "PIXEL_MAP_I_TO_G"
	enumLookup[FOG_COORDINATE_ARRAY_POINTER_EXT] = "FOG_COORDINATE_ARRAY_POINTER_EXT"
	enumLookup[CURRENT_VERTEX_ATTRIB_ARB] = "CURRENT_VERTEX_ATTRIB_ARB"
	enumLookup[LOCAL_EXT] = "LOCAL_EXT"
	enumLookup[MAX_TEXTURE_IMAGE_UNITS_ARB] = "MAX_TEXTURE_IMAGE_UNITS_ARB"
	enumLookup[TEXTURE_BORDER] = "TEXTURE_BORDER"
	enumLookup[BGRA] = "BGRA"
	enumLookup[VERTEX_ARRAY_RANGE_POINTER_NV] = "VERTEX_ARRAY_RANGE_POINTER_NV"
	enumLookup[R1UI_T2F_C4F_N3F_V3F_SUN] = "R1UI_T2F_C4F_N3F_V3F_SUN"
	enumLookup[DOT3_RGBA_IMG] = "DOT3_RGBA_IMG"
	enumLookup[DRAW_FRAMEBUFFER_EXT] = "DRAW_FRAMEBUFFER_EXT"
	enumLookup[MAX_FRAMEBUFFER_WIDTH] = "MAX_FRAMEBUFFER_WIDTH"
	enumLookup[PRIMARY_COLOR_EXT] = "PRIMARY_COLOR_EXT"
	enumLookup[REPLACEMENT_CODE_ARRAY_SUN] = "REPLACEMENT_CODE_ARRAY_SUN"
	enumLookup[DRAW_PIXELS_APPLE] = "DRAW_PIXELS_APPLE"
	enumLookup[VOLATILE_APPLE] = "VOLATILE_APPLE"
	enumLookup[RG16_SNORM] = "RG16_SNORM"
	enumLookup[MAX_COMBINED_SHADER_OUTPUT_RESOURCES] = "MAX_COMBINED_SHADER_OUTPUT_RESOURCES"
	enumLookup[FLOAT_RG16_NV] = "FLOAT_RG16_NV"
	enumLookup[RGB32I_EXT] = "RGB32I_EXT"
	enumLookup[LUMINANCE16_SNORM] = "LUMINANCE16_SNORM"
	enumLookup[VERTEX_ARRAY_RANGE_LENGTH_NV] = "VERTEX_ARRAY_RANGE_LENGTH_NV"
	enumLookup[REG_17_ATI] = "REG_17_ATI"
	enumLookup[FOG_COORD_ARRAY_TYPE] = "FOG_COORD_ARRAY_TYPE"
	enumLookup[MAP2_VERTEX_ATTRIB13_4_NV] = "MAP2_VERTEX_ATTRIB13_4_NV"
	enumLookup[DRAW_BUFFER12] = "DRAW_BUFFER12"
	enumLookup[Z4Y12Z4CB12Z4CR12_444_NV] = "Z4Y12Z4CB12Z4CR12_444_NV"
	enumLookup[FILE_NAME_NV] = "FILE_NAME_NV"
	enumLookup[SYNC_FLAGS_APPLE] = "SYNC_FLAGS_APPLE"
	enumLookup[PROGRAM_BINDING_ARB] = "PROGRAM_BINDING_ARB"
	enumLookup[DOUBLE_MAT4x2_EXT] = "DOUBLE_MAT4x2_EXT"
	enumLookup[MAP_FLUSH_EXPLICIT_BIT_EXT] = "MAP_FLUSH_EXPLICIT_BIT_EXT"
	enumLookup[CLAMP_VERTEX_COLOR_ARB] = "CLAMP_VERTEX_COLOR_ARB"
	enumLookup[GEOMETRY_SHADER_ARB] = "GEOMETRY_SHADER_ARB"
	enumLookup[INT_VEC3_ARB] = "INT_VEC3_ARB"
	enumLookup[DEBUG_CATEGORY_APPLICATION_AMD] = "DEBUG_CATEGORY_APPLICATION_AMD"
	enumLookup[TEXTURE_FETCH_BARRIER_BIT_EXT] = "TEXTURE_FETCH_BARRIER_BIT_EXT"
	enumLookup[MAX_COLOR_MATRIX_STACK_DEPTH] = "MAX_COLOR_MATRIX_STACK_DEPTH"
	enumLookup[POINT_SIZE] = "POINT_SIZE"
	enumLookup[TEXTURE_BINDING_CUBE_MAP_OES] = "TEXTURE_BINDING_CUBE_MAP_OES"
	enumLookup[MODELVIEW2_ARB] = "MODELVIEW2_ARB"
	enumLookup[ETC1_SRGB8_NV] = "ETC1_SRGB8_NV"
	enumLookup[RETAINED_APPLE] = "RETAINED_APPLE"
	enumLookup[COMBINE_RGB] = "COMBINE_RGB"
	enumLookup[VERTEX_ATTRIB_ARRAY_STRIDE] = "VERTEX_ATTRIB_ARRAY_STRIDE"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MULTISAMPLE] = "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"
	enumLookup[TEXCOORD3_BIT_PGI] = "TEXCOORD3_BIT_PGI"
	enumLookup[SGIX_ir_instrument1] = "SGIX_ir_instrument1"
	enumLookup[EVAL_VERTEX_ATTRIB0_NV] = "EVAL_VERTEX_ATTRIB0_NV"
	enumLookup[TEXTURE_COMPARE_FUNC_ARB] = "TEXTURE_COMPARE_FUNC_ARB"
	enumLookup[MAX_FRAGMENT_INTERPOLATION_OFFSET] = "MAX_FRAGMENT_INTERPOLATION_OFFSET"
	enumLookup[GL_4PASS_3_EXT] = "GL_4PASS_3_EXT"
	enumLookup[DOUBLE_VEC2] = "DOUBLE_VEC2"
	enumLookup[SUCCESS_NV] = "SUCCESS_NV"
	enumLookup[DOUBLE_EXT] = "DOUBLE_EXT"
	enumLookup[T4F_C4F_N3F_V4F] = "T4F_C4F_N3F_V4F"
	enumLookup[MINMAX_EXT] = "MINMAX_EXT"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Y] = "TEXTURE_CUBE_MAP_POSITIVE_Y"
	enumLookup[OFFSET_PROJECTIVE_TEXTURE_2D_NV] = "OFFSET_PROJECTIVE_TEXTURE_2D_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES] = "FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES"
	enumLookup[INDEX_ARRAY_LIST_STRIDE_IBM] = "INDEX_ARRAY_LIST_STRIDE_IBM"
	enumLookup[COLOR_SUM_CLAMP_NV] = "COLOR_SUM_CLAMP_NV"
	enumLookup[TEXTURE_RED_TYPE_ARB] = "TEXTURE_RED_TYPE_ARB"
	enumLookup[BACK_SECONDARY_COLOR_NV] = "BACK_SECONDARY_COLOR_NV"
	enumLookup[DEPTH_CLAMP_NV] = "DEPTH_CLAMP_NV"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_SIZE] = "TRANSFORM_FEEDBACK_BUFFER_SIZE"
	enumLookup[VERTEX_ATTRIB_ARRAY_UNIFIED_NV] = "VERTEX_ATTRIB_ARRAY_UNIFIED_NV"
	enumLookup[PACK_SWAP_BYTES] = "PACK_SWAP_BYTES"
	enumLookup[MAP2_INDEX] = "MAP2_INDEX"
	enumLookup[PACK_COMPRESSED_SIZE_SGIX] = "PACK_COMPRESSED_SIZE_SGIX"
	enumLookup[MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV] = "MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV"
	enumLookup[FRAMEBUFFER_ATTACHMENT_ANGLE] = "FRAMEBUFFER_ATTACHMENT_ANGLE"
	enumLookup[SGIX_async_pixel] = "SGIX_async_pixel"
	enumLookup[ALL_SHADER_BITS] = "ALL_SHADER_BITS"
	enumLookup[OBJECT_PLANE] = "OBJECT_PLANE"
	enumLookup[GENERATE_MIPMAP_HINT] = "GENERATE_MIPMAP_HINT"
	enumLookup[FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT] = "FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT"
	enumLookup[COLOR_ATTACHMENT12] = "COLOR_ATTACHMENT12"
	enumLookup[GEOMETRY_SHADER_BIT] = "GEOMETRY_SHADER_BIT"
	enumLookup[ALPHA32UI_EXT] = "ALPHA32UI_EXT"
	enumLookup[SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT] = "SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT"
	enumLookup[FEEDBACK] = "FEEDBACK"
	enumLookup[OPERAND1_ALPHA_EXT] = "OPERAND1_ALPHA_EXT"
	enumLookup[UNSIGNED_INT_10_10_10_2_OES] = "UNSIGNED_INT_10_10_10_2_OES"
	enumLookup[COVERAGE_SAMPLES_NV] = "COVERAGE_SAMPLES_NV"
	enumLookup[COPY_READ_BUFFER] = "COPY_READ_BUFFER"
	enumLookup[TOP_LEVEL_ARRAY_SIZE] = "TOP_LEVEL_ARRAY_SIZE"
	enumLookup[PHONG_WIN] = "PHONG_WIN"
	enumLookup[VERTEX_STREAM3_ATI] = "VERTEX_STREAM3_ATI"
	enumLookup[ARRAY_BUFFER] = "ARRAY_BUFFER"
	enumLookup[LERP_ATI] = "LERP_ATI"
	enumLookup[BLUE_BIT_ATI] = "BLUE_BIT_ATI"
	enumLookup[IMPLEMENTATION_COLOR_READ_FORMAT] = "IMPLEMENTATION_COLOR_READ_FORMAT"
	enumLookup[LUMINANCE_ALPHA] = "LUMINANCE_ALPHA"
	enumLookup[MAX_COLOR_ATTACHMENTS_EXT] = "MAX_COLOR_ATTACHMENTS_EXT"
	enumLookup[MAX_COMPUTE_SHARED_MEMORY_SIZE] = "MAX_COMPUTE_SHARED_MEMORY_SIZE"
	enumLookup[BLUE_MIN_CLAMP_INGR] = "BLUE_MIN_CLAMP_INGR"
	enumLookup[BLEND_EQUATION] = "BLEND_EQUATION"
	enumLookup[COMPARE_REF_TO_TEXTURE] = "COMPARE_REF_TO_TEXTURE"
	enumLookup[COMPUTE_LOCAL_WORK_SIZE] = "COMPUTE_LOCAL_WORK_SIZE"
	enumLookup[EXPAND_NEGATE_NV] = "EXPAND_NEGATE_NV"
	enumLookup[SAMPLER_2D_SHADOW] = "SAMPLER_2D_SHADOW"
	enumLookup[TEXTURE20_ARB] = "TEXTURE20_ARB"
	enumLookup[FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX] = "FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX"
	enumLookup[VERTEX_ATTRIB_ARRAY_ENABLED_ARB] = "VERTEX_ATTRIB_ARRAY_ENABLED_ARB"
	enumLookup[OP_MADD_EXT] = "OP_MADD_EXT"
	enumLookup[VERTEX_ATTRIB_ARRAY_NORMALIZED] = "VERTEX_ATTRIB_ARRAY_NORMALIZED"
	enumLookup[CON_27_ATI] = "CON_27_ATI"
	enumLookup[RENDERBUFFER_WIDTH_EXT] = "RENDERBUFFER_WIDTH_EXT"
	enumLookup[ONE_MINUS_CONSTANT_COLOR_EXT] = "ONE_MINUS_CONSTANT_COLOR_EXT"
	enumLookup[TEXTURE_GATHER_SHADOW] = "TEXTURE_GATHER_SHADOW"
	enumLookup[BUMP_ENVMAP_ATI] = "BUMP_ENVMAP_ATI"
	enumLookup[TEXTURE_CROP_RECT_OES] = "TEXTURE_CROP_RECT_OES"
	enumLookup[TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN] = "TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN"
	enumLookup[MAX_COMPUTE_SHADER_STORAGE_BLOCKS] = "MAX_COMPUTE_SHADER_STORAGE_BLOCKS"
	enumLookup[EXP] = "EXP"
	enumLookup[TRACE_MASK_MESA] = "TRACE_MASK_MESA"
	enumLookup[BUFFER_MAP_LENGTH] = "BUFFER_MAP_LENGTH"
	enumLookup[MATRIX_INDEX_ARRAY_ARB] = "MATRIX_INDEX_ARRAY_ARB"
	enumLookup[BOUNDING_BOX_OF_BOUNDING_BOXES_NV] = "BOUNDING_BOX_OF_BOUNDING_BOXES_NV"
	enumLookup[BUFFER_DATA_SIZE] = "BUFFER_DATA_SIZE"
	enumLookup[RED_BITS] = "RED_BITS"
	enumLookup[BINORMAL_ARRAY_POINTER_EXT] = "BINORMAL_ARRAY_POINTER_EXT"
	enumLookup[DRAW_BUFFER4_ATI] = "DRAW_BUFFER4_ATI"
	enumLookup[READ_PIXEL_DATA_RANGE_NV] = "READ_PIXEL_DATA_RANGE_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT] = "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT"
	enumLookup[MAX_SERVER_WAIT_TIMEOUT] = "MAX_SERVER_WAIT_TIMEOUT"
	enumLookup[VERTEX_ATTRIB_ARRAY0_NV] = "VERTEX_ATTRIB_ARRAY0_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_DIVISOR_ARB] = "VERTEX_ATTRIB_ARRAY_DIVISOR_ARB"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE"
	enumLookup[OBJECT_TYPE] = "OBJECT_TYPE"
	enumLookup[MAP1_NORMAL] = "MAP1_NORMAL"
	enumLookup[INTERNALFORMAT_STENCIL_SIZE] = "INTERNALFORMAT_STENCIL_SIZE"
	enumLookup[R1UI_T2F_N3F_V3F_SUN] = "R1UI_T2F_N3F_V3F_SUN"
	enumLookup[BOOL_VEC4_ARB] = "BOOL_VEC4_ARB"
	enumLookup[SLUMINANCE8_NV] = "SLUMINANCE8_NV"
	enumLookup[SGIX_list_priority] = "SGIX_list_priority"
	enumLookup[AND_REVERSE] = "AND_REVERSE"
	enumLookup[BLEND_SRC_RGB] = "BLEND_SRC_RGB"
	enumLookup[TEXTURE8_ARB] = "TEXTURE8_ARB"
	enumLookup[OPERAND1_RGB_ARB] = "OPERAND1_RGB_ARB"
	enumLookup[VERTEX_ATTRIB_ARRAY_POINTER_ARB] = "VERTEX_ATTRIB_ARRAY_POINTER_ARB"
	enumLookup[VERTEX_PROGRAM_CALLBACK_DATA_MESA] = "VERTEX_PROGRAM_CALLBACK_DATA_MESA"
	enumLookup[PRIMITIVE_ID_NV] = "PRIMITIVE_ID_NV"
	enumLookup[IMAGE_2D_MULTISAMPLE_EXT] = "IMAGE_2D_MULTISAMPLE_EXT"
	enumLookup[MAX_DEBUG_MESSAGE_LENGTH_ARB] = "MAX_DEBUG_MESSAGE_LENGTH_ARB"
	enumLookup[FULL_SUPPORT] = "FULL_SUPPORT"
	enumLookup[MAX_TEXTURE_UNITS_ARB] = "MAX_TEXTURE_UNITS_ARB"
	enumLookup[PIXEL_SUBSAMPLE_2424_SGIX] = "PIXEL_SUBSAMPLE_2424_SGIX"
	enumLookup[PATH_FILL_MODE_NV] = "PATH_FILL_MODE_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG] = "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG"
	enumLookup[COLOR_TABLE_LUMINANCE_SIZE_SGI] = "COLOR_TABLE_LUMINANCE_SIZE_SGI"
	enumLookup[BGR] = "BGR"
	enumLookup[COMBINER2_NV] = "COMBINER2_NV"
	enumLookup[EVAL_VERTEX_ATTRIB5_NV] = "EVAL_VERTEX_ATTRIB5_NV"
	enumLookup[PROGRAM_FORMAT_ARB] = "PROGRAM_FORMAT_ARB"
	enumLookup[SAMPLER_CUBE_MAP_ARRAY] = "SAMPLER_CUBE_MAP_ARRAY"
	enumLookup[FRAGMENT_PROGRAM_ARB] = "FRAGMENT_PROGRAM_ARB"
	enumLookup[MATRIX24_ARB] = "MATRIX24_ARB"
	enumLookup[ARRAY_OBJECT_OFFSET_ATI] = "ARRAY_OBJECT_OFFSET_ATI"
	enumLookup[ALWAYS_FAST_HINT_PGI] = "ALWAYS_FAST_HINT_PGI"
	enumLookup[COLOR_ARRAY_POINTER] = "COLOR_ARRAY_POINTER"
	enumLookup[COLOR_ARRAY_SIZE] = "COLOR_ARRAY_SIZE"
	enumLookup[CLIP_PLANE4] = "CLIP_PLANE4"
	enumLookup[YCRCB_SGIX] = "YCRCB_SGIX"
	enumLookup[SOURCE1_ALPHA] = "SOURCE1_ALPHA"
	enumLookup[RGBA16F_EXT] = "RGBA16F_EXT"
	enumLookup[STENCIL_INDEX16_EXT] = "STENCIL_INDEX16_EXT"
	enumLookup[DUP_FIRST_CUBIC_CURVE_TO_NV] = "DUP_FIRST_CUBIC_CURVE_TO_NV"
	enumLookup[SELECT] = "SELECT"
	enumLookup[GL_3DC_X_AMD] = "GL_3DC_X_AMD"
	enumLookup[COLOR_ATTACHMENT0_EXT] = "COLOR_ATTACHMENT0_EXT"
	enumLookup[TEXTURE_GEN_STR_OES] = "TEXTURE_GEN_STR_OES"
	enumLookup[DEPTH_COMPONENT16_NONLINEAR_NV] = "DEPTH_COMPONENT16_NONLINEAR_NV"
	enumLookup[VIDEO_CAPTURE_TO_422_SUPPORTED_NV] = "VIDEO_CAPTURE_TO_422_SUPPORTED_NV"
	enumLookup[VIEW_CLASS_S3TC_DXT1_RGBA] = "VIEW_CLASS_S3TC_DXT1_RGBA"
	enumLookup[TEXTURE9_ARB] = "TEXTURE9_ARB"
	enumLookup[DSDT_MAG_VIB_NV] = "DSDT_MAG_VIB_NV"
	enumLookup[CURRENT_MATRIX_INDEX_ARB] = "CURRENT_MATRIX_INDEX_ARB"
	enumLookup[MATRIX_PALETTE_OES] = "MATRIX_PALETTE_OES"
	enumLookup[ATC_RGB_AMD] = "ATC_RGB_AMD"
	enumLookup[MAX_PROGRAM_TEXTURE_GATHER_OFFSET] = "MAX_PROGRAM_TEXTURE_GATHER_OFFSET"
	enumLookup[PATH_GEN_COEFF_NV] = "PATH_GEN_COEFF_NV"
	enumLookup[RGBA12_EXT] = "RGBA12_EXT"
	enumLookup[PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP] = "PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP"
	enumLookup[DEBUG_SOURCE_WINDOW_SYSTEM_ARB] = "DEBUG_SOURCE_WINDOW_SYSTEM_ARB"
	enumLookup[MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT] = "MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT"
	enumLookup[TEXTURE_BINDING_2D_ARRAY_EXT] = "TEXTURE_BINDING_2D_ARRAY_EXT"
	enumLookup[TEXTURE_BUFFER_EXT] = "TEXTURE_BUFFER_EXT"
	enumLookup[LOCATION] = "LOCATION"
	enumLookup[MAX_TEXTURE_SIZE] = "MAX_TEXTURE_SIZE"
	enumLookup[STENCIL_TEST_TWO_SIDE_EXT] = "STENCIL_TEST_TWO_SIDE_EXT"
	enumLookup[MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS] = "MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS"
	enumLookup[UPPER_LEFT] = "UPPER_LEFT"
	enumLookup[COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT] = "COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT"
	enumLookup[FONT_HEIGHT_BIT_NV] = "FONT_HEIGHT_BIT_NV"
	enumLookup[SGIX_texture_add_env] = "SGIX_texture_add_env"
	enumLookup[NORMALIZED_RANGE_EXT] = "NORMALIZED_RANGE_EXT"
	enumLookup[MAX_PROGRAM_NATIVE_ATTRIBS_ARB] = "MAX_PROGRAM_NATIVE_ATTRIBS_ARB"
	enumLookup[FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT] = "FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT"
	enumLookup[MAX_FRAMEBUFFER_HEIGHT] = "MAX_FRAMEBUFFER_HEIGHT"
	enumLookup[MATERIAL_SIDE_HINT_PGI] = "MATERIAL_SIDE_HINT_PGI"
	enumLookup[TEXTURE_BINDING_2D] = "TEXTURE_BINDING_2D"
	enumLookup[TEXTURE_3D_OES] = "TEXTURE_3D_OES"
	enumLookup[MIPMAP] = "MIPMAP"
	enumLookup[VIEW_CLASS_24_BITS] = "VIEW_CLASS_24_BITS"
	enumLookup[SYNC_GPU_COMMANDS_COMPLETE] = "SYNC_GPU_COMMANDS_COMPLETE"
	enumLookup[FOG_COORDINATE_ARRAY_LIST_IBM] = "FOG_COORDINATE_ARRAY_LIST_IBM"
	enumLookup[LAYOUT_LINEAR_CPU_CACHED_INTEL] = "LAYOUT_LINEAR_CPU_CACHED_INTEL"
	enumLookup[FRONT_FACE] = "FRONT_FACE"
	enumLookup[TEXTURE_LOD_BIAS_R_SGIX] = "TEXTURE_LOD_BIAS_R_SGIX"
	enumLookup[GUILTY_CONTEXT_RESET_ARB] = "GUILTY_CONTEXT_RESET_ARB"
	enumLookup[MAX_WIDTH] = "MAX_WIDTH"
	enumLookup[GET_TEXTURE_IMAGE_TYPE] = "GET_TEXTURE_IMAGE_TYPE"
	enumLookup[OUTPUT_TEXTURE_COORD0_EXT] = "OUTPUT_TEXTURE_COORD0_EXT"
	enumLookup[VIRTUAL_PAGE_SIZE_X_AMD] = "VIRTUAL_PAGE_SIZE_X_AMD"
	enumLookup[ASYNC_HISTOGRAM_SGIX] = "ASYNC_HISTOGRAM_SGIX"
	enumLookup[MAX_ASYNC_DRAW_PIXELS_SGIX] = "MAX_ASYNC_DRAW_PIXELS_SGIX"
	enumLookup[SECONDARY_COLOR_ARRAY_POINTER] = "SECONDARY_COLOR_ARRAY_POINTER"
	enumLookup[COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV] = "COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV"
	enumLookup[INTENSITY8I_EXT] = "INTENSITY8I_EXT"
	enumLookup[POST_COLOR_MATRIX_GREEN_BIAS_SGI] = "POST_COLOR_MATRIX_GREEN_BIAS_SGI"
	enumLookup[BUFFER_ACCESS_OES] = "BUFFER_ACCESS_OES"
	enumLookup[UNSIGNED_INT_5_9_9_9_REV_EXT] = "UNSIGNED_INT_5_9_9_9_REV_EXT"
	enumLookup[ALPHA_TEST_FUNC_QCOM] = "ALPHA_TEST_FUNC_QCOM"
	enumLookup[LUMINANCE4_ALPHA4] = "LUMINANCE4_ALPHA4"
	enumLookup[UNSIGNED_SHORT_4_4_4_4_REV_EXT] = "UNSIGNED_SHORT_4_4_4_4_REV_EXT"
	enumLookup[INDEX_ARRAY_STRIDE_EXT] = "INDEX_ARRAY_STRIDE_EXT"
	enumLookup[GL_422_REV_EXT] = "GL_422_REV_EXT"
	enumLookup[SUBTRACT_ARB] = "SUBTRACT_ARB"
	enumLookup[MIRROR_CLAMP_ATI] = "MIRROR_CLAMP_ATI"
	enumLookup[DRAW_BUFFER10_ARB] = "DRAW_BUFFER10_ARB"
	enumLookup[SAMPLER_BUFFER_EXT] = "SAMPLER_BUFFER_EXT"
	enumLookup[IMAGE_BINDING_LAYERED_EXT] = "IMAGE_BINDING_LAYERED_EXT"
	enumLookup[MAX_LIGHTS] = "MAX_LIGHTS"
	enumLookup[CONVOLUTION_FILTER_SCALE] = "CONVOLUTION_FILTER_SCALE"
	enumLookup[TEXTURE_COORD_ARRAY_POINTER_EXT] = "TEXTURE_COORD_ARRAY_POINTER_EXT"
	enumLookup[REG_31_ATI] = "REG_31_ATI"
	enumLookup[SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST] = "SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST"
	enumLookup[BLUE_MAX_CLAMP_INGR] = "BLUE_MAX_CLAMP_INGR"
	enumLookup[PROXY_TEXTURE_3D_EXT] = "PROXY_TEXTURE_3D_EXT"
	enumLookup[INTERNALFORMAT_DEPTH_TYPE] = "INTERNALFORMAT_DEPTH_TYPE"
	enumLookup[MODELVIEW12_ARB] = "MODELVIEW12_ARB"
	enumLookup[PROGRAM_ERROR_STRING_ARB] = "PROGRAM_ERROR_STRING_ARB"
	enumLookup[CULL_VERTEX_EXT] = "CULL_VERTEX_EXT"
	enumLookup[DECR_WRAP_OES] = "DECR_WRAP_OES"
	enumLookup[UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER] = "UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER"
	enumLookup[OBJECT_LINK_STATUS_ARB] = "OBJECT_LINK_STATUS_ARB"
	enumLookup[DOUBLE_MAT2x4] = "DOUBLE_MAT2x4"
	enumLookup[DEPTH_BIAS] = "DEPTH_BIAS"
	enumLookup[POST_COLOR_MATRIX_GREEN_SCALE] = "POST_COLOR_MATRIX_GREEN_SCALE"
	enumLookup[PROGRAM_ERROR_POSITION_ARB] = "PROGRAM_ERROR_POSITION_ARB"
	enumLookup[REG_29_ATI] = "REG_29_ATI"
	enumLookup[COLOR_ALPHA_PAIRING_ATI] = "COLOR_ALPHA_PAIRING_ATI"
	enumLookup[READ_FRAMEBUFFER_BINDING_EXT] = "READ_FRAMEBUFFER_BINDING_EXT"
	enumLookup[MAX_VERTEX_OUTPUT_COMPONENTS] = "MAX_VERTEX_OUTPUT_COMPONENTS"
	enumLookup[CLIENT_ALL_ATTRIB_BITS] = "CLIENT_ALL_ATTRIB_BITS"
	enumLookup[R3_G3_B2] = "R3_G3_B2"
	enumLookup[COMPRESSED_RGBA_S3TC_DXT3_ANGLE] = "COMPRESSED_RGBA_S3TC_DXT3_ANGLE"
	enumLookup[SAMPLE_SHADING_ARB] = "SAMPLE_SHADING_ARB"
	enumLookup[INTERLEAVED_ATTRIBS] = "INTERLEAVED_ATTRIBS"
	enumLookup[INT_SAMPLER_2D] = "INT_SAMPLER_2D"
	enumLookup[SMOOTH_CUBIC_CURVE_TO_NV] = "SMOOTH_CUBIC_CURVE_TO_NV"
	enumLookup[STENCIL_TEST] = "STENCIL_TEST"
	enumLookup[T2F_V3F] = "T2F_V3F"
	enumLookup[RG_INTEGER] = "RG_INTEGER"
	enumLookup[PROGRAM_TEMPORARIES_ARB] = "PROGRAM_TEMPORARIES_ARB"
	enumLookup[UNIFORM_BLOCK_INDEX] = "UNIFORM_BLOCK_INDEX"
	enumLookup[HIGH_FLOAT] = "HIGH_FLOAT"
	enumLookup[TRANSPOSE_COLOR_MATRIX] = "TRANSPOSE_COLOR_MATRIX"
	enumLookup[FONT_Y_MIN_BOUNDS_BIT_NV] = "FONT_Y_MIN_BOUNDS_BIT_NV"
	enumLookup[TESS_EVALUATION_SUBROUTINE_UNIFORM] = "TESS_EVALUATION_SUBROUTINE_UNIFORM"
	enumLookup[EDGE_FLAG_ARRAY_LIST_STRIDE_IBM] = "EDGE_FLAG_ARRAY_LIST_STRIDE_IBM"
	enumLookup[TEXTURE_COORD_ARRAY] = "TEXTURE_COORD_ARRAY"
	enumLookup[INTENSITY4] = "INTENSITY4"
	enumLookup[MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS] = "MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS"
	enumLookup[VERTEX_ATTRIB_ARRAY_TYPE_ARB] = "VERTEX_ATTRIB_ARRAY_TYPE_ARB"
	enumLookup[VERTEX_SHADER_LOCAL_CONSTANTS_EXT] = "VERTEX_SHADER_LOCAL_CONSTANTS_EXT"
	enumLookup[FOG_COORD_ARRAY_BUFFER_BINDING] = "FOG_COORD_ARRAY_BUFFER_BINDING"
	enumLookup[RESAMPLE_AVERAGE_OML] = "RESAMPLE_AVERAGE_OML"
	enumLookup[GEOMETRY_VERTICES_OUT_EXT] = "GEOMETRY_VERTICES_OUT_EXT"
	enumLookup[SMOOTH_LINE_WIDTH_RANGE] = "SMOOTH_LINE_WIDTH_RANGE"
	enumLookup[COMPRESSED_LUMINANCE_LATC1_EXT] = "COMPRESSED_LUMINANCE_LATC1_EXT"
	enumLookup[TEXTURE_3D_BINDING_EXT] = "TEXTURE_3D_BINDING_EXT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE] = "FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE"
	enumLookup[R32F] = "R32F"
	enumLookup[COORD_REPLACE_NV] = "COORD_REPLACE_NV"
	enumLookup[SGIS_sharpen_texture] = "SGIS_sharpen_texture"
	enumLookup[ACCUM_CLEAR_VALUE] = "ACCUM_CLEAR_VALUE"
	enumLookup[FORMAT_SUBSAMPLE_24_24_OML] = "FORMAT_SUBSAMPLE_24_24_OML"
	enumLookup[CW] = "CW"
	enumLookup[STENCIL_PASS_DEPTH_FAIL] = "STENCIL_PASS_DEPTH_FAIL"
	enumLookup[POLYGON_OFFSET_LINE] = "POLYGON_OFFSET_LINE"
	enumLookup[SHADOW_ATTENUATION_EXT] = "SHADOW_ATTENUATION_EXT"
	enumLookup[COMBINER_BIAS_NV] = "COMBINER_BIAS_NV"
	enumLookup[FLOAT_R32_NV] = "FLOAT_R32_NV"
	enumLookup[FAILURE_NV] = "FAILURE_NV"
	enumLookup[TEXTURE17] = "TEXTURE17"
	enumLookup[OP_MAX_EXT] = "OP_MAX_EXT"
	enumLookup[FLOAT_R_NV] = "FLOAT_R_NV"
	enumLookup[MAX_TEXTURE_BUFFER_SIZE] = "MAX_TEXTURE_BUFFER_SIZE"
	enumLookup[MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV] = "MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV"
	enumLookup[UNSIGNED_INT8_VEC3_NV] = "UNSIGNED_INT8_VEC3_NV"
	enumLookup[COMPRESSED_RGBA_ASTC_8x6_KHR] = "COMPRESSED_RGBA_ASTC_8x6_KHR"
	enumLookup[INCR_WRAP_OES] = "INCR_WRAP_OES"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Z] = "TEXTURE_CUBE_MAP_POSITIVE_Z"
	enumLookup[ELEMENT_ARRAY_ATI] = "ELEMENT_ARRAY_ATI"
	enumLookup[POINT_SIZE_ARRAY_BUFFER_BINDING_OES] = "POINT_SIZE_ARRAY_BUFFER_BINDING_OES"
	enumLookup[PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV] = "PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV"
	enumLookup[FOG_END] = "FOG_END"
	enumLookup[PROXY_COLOR_TABLE] = "PROXY_COLOR_TABLE"
	enumLookup[EYE_DISTANCE_TO_LINE_SGIS] = "EYE_DISTANCE_TO_LINE_SGIS"
	enumLookup[MATRIX6_NV] = "MATRIX6_NV"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_EXT] = "TRANSFORM_FEEDBACK_BUFFER_EXT"
	enumLookup[MIN_FRAGMENT_INTERPOLATION_OFFSET_NV] = "MIN_FRAGMENT_INTERPOLATION_OFFSET_NV"
	enumLookup[SIGNED_NORMALIZED] = "SIGNED_NORMALIZED"
	enumLookup[INT_IMAGE_BUFFER_EXT] = "INT_IMAGE_BUFFER_EXT"
	enumLookup[PIXEL_MAP_I_TO_R_SIZE] = "PIXEL_MAP_I_TO_R_SIZE"
	enumLookup[INDEX_OFFSET] = "INDEX_OFFSET"
	enumLookup[HALF_FLOAT_NV] = "HALF_FLOAT_NV"
	enumLookup[RGB8_EXT] = "RGB8_EXT"
	enumLookup[INTERNALFORMAT_RED_SIZE] = "INTERNALFORMAT_RED_SIZE"
	enumLookup[MATRIX31_ARB] = "MATRIX31_ARB"
	enumLookup[RGBA8UI] = "RGBA8UI"
	enumLookup[MAX_MULTISAMPLE_COVERAGE_MODES_NV] = "MAX_MULTISAMPLE_COVERAGE_MODES_NV"
	enumLookup[VERTEX_SHADER_EXT] = "VERTEX_SHADER_EXT"
	enumLookup[MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT] = "MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT"
	enumLookup[SLUMINANCE_ALPHA] = "SLUMINANCE_ALPHA"
	enumLookup[RENDERBUFFER_RED_SIZE] = "RENDERBUFFER_RED_SIZE"
	enumLookup[FONT_UNITS_PER_EM_BIT_NV] = "FONT_UNITS_PER_EM_BIT_NV"
	enumLookup[ELEMENT_ARRAY_BARRIER_BIT_EXT] = "ELEMENT_ARRAY_BARRIER_BIT_EXT"
	enumLookup[GLYPH_HEIGHT_BIT_NV] = "GLYPH_HEIGHT_BIT_NV"
	enumLookup[MAP1_VERTEX_3] = "MAP1_VERTEX_3"
	enumLookup[SLUMINANCE8_ALPHA8_EXT] = "SLUMINANCE8_ALPHA8_EXT"
	enumLookup[LIGHT1] = "LIGHT1"
	enumLookup[DEBUG_TYPE_PERFORMANCE_ARB] = "DEBUG_TYPE_PERFORMANCE_ARB"
	enumLookup[SPARE1_NV] = "SPARE1_NV"
	enumLookup[OPERAND1_ALPHA] = "OPERAND1_ALPHA"
	enumLookup[PN_TRIANGLES_TESSELATION_LEVEL_ATI] = "PN_TRIANGLES_TESSELATION_LEVEL_ATI"
	enumLookup[MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB] = "MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB"
	enumLookup[GEOMETRY_SHADER_EXT] = "GEOMETRY_SHADER_EXT"
	enumLookup[EDGE_FLAG_ARRAY_ADDRESS_NV] = "EDGE_FLAG_ARRAY_ADDRESS_NV"
	enumLookup[SAMPLER_2D_MULTISAMPLE_ARRAY] = "SAMPLER_2D_MULTISAMPLE_ARRAY"
	enumLookup[BACK_LEFT] = "BACK_LEFT"
	enumLookup[RASTERIZER_DISCARD] = "RASTERIZER_DISCARD"
	enumLookup[PATH_CLIENT_LENGTH_NV] = "PATH_CLIENT_LENGTH_NV"
	enumLookup[CURRENT_COLOR] = "CURRENT_COLOR"
	enumLookup[CON_0_ATI] = "CON_0_ATI"
	enumLookup[NUM_FILL_STREAMS_NV] = "NUM_FILL_STREAMS_NV"
	enumLookup[MAX_IMAGE_UNITS_EXT] = "MAX_IMAGE_UNITS_EXT"
	enumLookup[SMALL_CCW_ARC_TO_NV] = "SMALL_CCW_ARC_TO_NV"
	enumLookup[SGIX_reference_plane] = "SGIX_reference_plane"
	enumLookup[SAMPLE_ALPHA_TO_ONE_EXT] = "SAMPLE_ALPHA_TO_ONE_EXT"
	enumLookup[SHADER_IMAGE_STORE] = "SHADER_IMAGE_STORE"
	enumLookup[FRAMEBUFFER_DEFAULT_WIDTH] = "FRAMEBUFFER_DEFAULT_WIDTH"
	enumLookup[NORMAL_ARRAY_COUNT_EXT] = "NORMAL_ARRAY_COUNT_EXT"
	enumLookup[RENDERBUFFER_FREE_MEMORY_ATI] = "RENDERBUFFER_FREE_MEMORY_ATI"
	enumLookup[TEXTURE_COMPARE_MODE] = "TEXTURE_COMPARE_MODE"
	enumLookup[IMPLEMENTATION_COLOR_READ_FORMAT_OES] = "IMPLEMENTATION_COLOR_READ_FORMAT_OES"
	enumLookup[TEXTURE_ALPHA_TYPE] = "TEXTURE_ALPHA_TYPE"
	enumLookup[MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI] = "MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI"
	enumLookup[NAMED_STRING_LENGTH_ARB] = "NAMED_STRING_LENGTH_ARB"
	enumLookup[PROVOKING_VERTEX_EXT] = "PROVOKING_VERTEX_EXT"
	enumLookup[QUERY_BUFFER_BINDING_AMD] = "QUERY_BUFFER_BINDING_AMD"
	enumLookup[CURRENT_RASTER_INDEX] = "CURRENT_RASTER_INDEX"
	enumLookup[DEBUG_TYPE_PERFORMANCE] = "DEBUG_TYPE_PERFORMANCE"
	enumLookup[MIN_PROGRAM_TEXEL_OFFSET_NV] = "MIN_PROGRAM_TEXEL_OFFSET_NV"
	enumLookup[DELETE_STATUS] = "DELETE_STATUS"
	enumLookup[COMPRESSED_RGB_PVRTC_4BPPV1_IMG] = "COMPRESSED_RGB_PVRTC_4BPPV1_IMG"
	enumLookup[LOW_FLOAT] = "LOW_FLOAT"
	enumLookup[DEBUG_SEVERITY_LOW_AMD] = "DEBUG_SEVERITY_LOW_AMD"
	enumLookup[PIXEL_MAP_G_TO_G] = "PIXEL_MAP_G_TO_G"
	enumLookup[CONVOLUTION_FILTER_SCALE_EXT] = "CONVOLUTION_FILTER_SCALE_EXT"
	enumLookup[PACK_RESAMPLE_SGIX] = "PACK_RESAMPLE_SGIX"
	enumLookup[OPERAND0_RGB] = "OPERAND0_RGB"
	enumLookup[BUFFER_USAGE] = "BUFFER_USAGE"
	enumLookup[COMPRESSED_RGBA_PVRTC_4BPPV1_IMG] = "COMPRESSED_RGBA_PVRTC_4BPPV1_IMG"
	enumLookup[COVERAGE_BUFFER_BIT_NV] = "COVERAGE_BUFFER_BIT_NV"
	enumLookup[CUBIC_CURVE_TO_NV] = "CUBIC_CURVE_TO_NV"
	enumLookup[POINTS] = "POINTS"
	enumLookup[CURRENT_RASTER_POSITION] = "CURRENT_RASTER_POSITION"
	enumLookup[TEXTURE_RED_SIZE] = "TEXTURE_RED_SIZE"
	enumLookup[WEIGHT_ARRAY_OES] = "WEIGHT_ARRAY_OES"
	enumLookup[RENDERBUFFER_ALPHA_SIZE_OES] = "RENDERBUFFER_ALPHA_SIZE_OES"
	enumLookup[VIRTUAL_PAGE_SIZE_Z_AMD] = "VIRTUAL_PAGE_SIZE_Z_AMD"
	enumLookup[POST_TEXTURE_FILTER_BIAS_RANGE_SGIX] = "POST_TEXTURE_FILTER_BIAS_RANGE_SGIX"
	enumLookup[OP_NEGATE_EXT] = "OP_NEGATE_EXT"
	enumLookup[UNSIGNED_SHORT_5_6_5_REV] = "UNSIGNED_SHORT_5_6_5_REV"
	enumLookup[OP_ADD_EXT] = "OP_ADD_EXT"
	enumLookup[ALPHA8I_EXT] = "ALPHA8I_EXT"
	enumLookup[ELEMENT_ARRAY_ADDRESS_NV] = "ELEMENT_ARRAY_ADDRESS_NV"
	enumLookup[UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT] = "UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT"
	enumLookup[COLOR_ATTACHMENT2_EXT] = "COLOR_ATTACHMENT2_EXT"
	enumLookup[RGBA_INTEGER_EXT] = "RGBA_INTEGER_EXT"
	enumLookup[COLOR_ARRAY_LENGTH_NV] = "COLOR_ARRAY_LENGTH_NV"
	enumLookup[EDGE_FLAG_ARRAY_LENGTH_NV] = "EDGE_FLAG_ARRAY_LENGTH_NV"
	enumLookup[FONT_MAX_ADVANCE_HEIGHT_BIT_NV] = "FONT_MAX_ADVANCE_HEIGHT_BIT_NV"
	enumLookup[SAMPLE_COVERAGE] = "SAMPLE_COVERAGE"
	enumLookup[WRITE_DISCARD_NV] = "WRITE_DISCARD_NV"
	enumLookup[COLOR_TABLE_ALPHA_SIZE_SGI] = "COLOR_TABLE_ALPHA_SIZE_SGI"
	enumLookup[POST_TEXTURE_FILTER_BIAS_SGIX] = "POST_TEXTURE_FILTER_BIAS_SGIX"
	enumLookup[FRAGMENT_LIGHT5_SGIX] = "FRAGMENT_LIGHT5_SGIX"
	enumLookup[UNSIGNED_INT_24_8_EXT] = "UNSIGNED_INT_24_8_EXT"
	enumLookup[VERTEX_ARRAY_RANGE_NV] = "VERTEX_ARRAY_RANGE_NV"
	enumLookup[DEPTH24_STENCIL8_OES] = "DEPTH24_STENCIL8_OES"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT] = "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	enumLookup[MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS] = "MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS"
	enumLookup[INT_10_10_10_2_OES] = "INT_10_10_10_2_OES"
	enumLookup[DYNAMIC_READ_ARB] = "DYNAMIC_READ_ARB"
	enumLookup[INT8_VEC2_NV] = "INT8_VEC2_NV"
	enumLookup[BLEND_SRC_RGB_EXT] = "BLEND_SRC_RGB_EXT"
	enumLookup[RGBA_DXT5_S3TC] = "RGBA_DXT5_S3TC"
	enumLookup[TEXTURE27] = "TEXTURE27"
	enumLookup[FRAGMENT_PROGRAM_BINDING_NV] = "FRAGMENT_PROGRAM_BINDING_NV"
	enumLookup[NUM_PASSES_ATI] = "NUM_PASSES_ATI"
	enumLookup[INT_SAMPLER_2D_ARRAY] = "INT_SAMPLER_2D_ARRAY"
	enumLookup[COPY_WRITE_BUFFER] = "COPY_WRITE_BUFFER"
	enumLookup[MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB] = "MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB"
	enumLookup[SRC1_RGB] = "SRC1_RGB"
	enumLookup[DRAW_BUFFER6_ATI] = "DRAW_BUFFER6_ATI"
	enumLookup[PATH_INITIAL_DASH_CAP_NV] = "PATH_INITIAL_DASH_CAP_NV"
	enumLookup[BOUNDING_BOX_NV] = "BOUNDING_BOX_NV"
	enumLookup[DEBUG_SEVERITY_MEDIUM_ARB] = "DEBUG_SEVERITY_MEDIUM_ARB"
	enumLookup[COMPRESSED_ALPHA] = "COMPRESSED_ALPHA"
	enumLookup[NORMAL_MAP_EXT] = "NORMAL_MAP_EXT"
	enumLookup[PROXY_TEXTURE_2D_ARRAY] = "PROXY_TEXTURE_2D_ARRAY"
	enumLookup[VERTEX_ATTRIB_ARRAY_SIZE_ARB] = "VERTEX_ATTRIB_ARRAY_SIZE_ARB"
	enumLookup[MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB] = "MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB"
	enumLookup[MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV] = "MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV"
	enumLookup[INT_IMAGE_2D_ARRAY] = "INT_IMAGE_2D_ARRAY"
	enumLookup[SGIS_generate_mipmap] = "SGIS_generate_mipmap"
	enumLookup[IR_INSTRUMENT1_SGIX] = "IR_INSTRUMENT1_SGIX"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Y] = "TEXTURE_CUBE_MAP_NEGATIVE_Y"
	enumLookup[ADD_ATI] = "ADD_ATI"
	enumLookup[FLOAT_VEC3_ARB] = "FLOAT_VEC3_ARB"
	enumLookup[FRAMEBUFFER_ATTACHMENT_LAYERED] = "FRAMEBUFFER_ATTACHMENT_LAYERED"
	enumLookup[INT_SAMPLER_1D] = "INT_SAMPLER_1D"
	enumLookup[YCBYCR8_422_NV] = "YCBYCR8_422_NV"
	enumLookup[RGB8] = "RGB8"
	enumLookup[PIXEL_TEX_GEN_ALPHA_LS_SGIX] = "PIXEL_TEX_GEN_ALPHA_LS_SGIX"
	enumLookup[COLOR_SUM_EXT] = "COLOR_SUM_EXT"
	enumLookup[MODELVIEW10_ARB] = "MODELVIEW10_ARB"
	enumLookup[VERTEX_STREAM2_ATI] = "VERTEX_STREAM2_ATI"
	enumLookup[RENDERBUFFER_STENCIL_SIZE_OES] = "RENDERBUFFER_STENCIL_SIZE_OES"
	enumLookup[NUM_COMPATIBLE_SUBROUTINES] = "NUM_COMPATIBLE_SUBROUTINES"
	enumLookup[MAX_INTEGER_SAMPLES] = "MAX_INTEGER_SAMPLES"
	enumLookup[PACK_COMPRESSED_BLOCK_DEPTH] = "PACK_COMPRESSED_BLOCK_DEPTH"
	enumLookup[EDGE_FLAG_ARRAY_LIST_IBM] = "EDGE_FLAG_ARRAY_LIST_IBM"
	enumLookup[MULTISAMPLE_EXT] = "MULTISAMPLE_EXT"
	enumLookup[MODELVIEW1_MATRIX_EXT] = "MODELVIEW1_MATRIX_EXT"
	enumLookup[DRAW_BUFFER0_NV] = "DRAW_BUFFER0_NV"
	enumLookup[SLUMINANCE_NV] = "SLUMINANCE_NV"
	enumLookup[SAMPLE_PATTERN_SGIS] = "SAMPLE_PATTERN_SGIS"
	enumLookup[SHARPEN_TEXTURE_FUNC_POINTS_SGIS] = "SHARPEN_TEXTURE_FUNC_POINTS_SGIS"
	enumLookup[MAP1_VERTEX_ATTRIB13_4_NV] = "MAP1_VERTEX_ATTRIB13_4_NV"
	enumLookup[PROGRAM_BINARY_FORMATS] = "PROGRAM_BINARY_FORMATS"
	enumLookup[STATIC_COPY_ARB] = "STATIC_COPY_ARB"
	enumLookup[TESSELLATION_FACTOR_AMD] = "TESSELLATION_FACTOR_AMD"
	enumLookup[ALLOW_DRAW_WIN_HINT_PGI] = "ALLOW_DRAW_WIN_HINT_PGI"
	enumLookup[MAX_MODELVIEW_STACK_DEPTH] = "MAX_MODELVIEW_STACK_DEPTH"
	enumLookup[DOUBLE_VEC4_EXT] = "DOUBLE_VEC4_EXT"
	enumLookup[TEXTURE_BORDER_COLOR_NV] = "TEXTURE_BORDER_COLOR_NV"
	enumLookup[DEPTH_STENCIL_NV] = "DEPTH_STENCIL_NV"
	enumLookup[DEPTH_CLAMP] = "DEPTH_CLAMP"
	enumLookup[CON_10_ATI] = "CON_10_ATI"
	enumLookup[COLOR_ATTACHMENT5_EXT] = "COLOR_ATTACHMENT5_EXT"
	enumLookup[COLOR_ATTACHMENT10_NV] = "COLOR_ATTACHMENT10_NV"
	enumLookup[MULTISAMPLE_BIT] = "MULTISAMPLE_BIT"
	enumLookup[COLOR_TABLE_RED_SIZE_SGI] = "COLOR_TABLE_RED_SIZE_SGI"
	enumLookup[TEXTURE_LEQUAL_R_SGIX] = "TEXTURE_LEQUAL_R_SGIX"
	enumLookup[INTERNALFORMAT_PREFERRED] = "INTERNALFORMAT_PREFERRED"
	enumLookup[RGBA8UI_EXT] = "RGBA8UI_EXT"
	enumLookup[DRAW_INDIRECT_ADDRESS_NV] = "DRAW_INDIRECT_ADDRESS_NV"
	enumLookup[MAX_SHADER_STORAGE_BUFFER_BINDINGS] = "MAX_SHADER_STORAGE_BUFFER_BINDINGS"
	enumLookup[HISTOGRAM_SINK] = "HISTOGRAM_SINK"
	enumLookup[CONTEXT_FLAGS] = "CONTEXT_FLAGS"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT] = "TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT"
	enumLookup[TRACE_ERRORS_BIT_MESA] = "TRACE_ERRORS_BIT_MESA"
	enumLookup[MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS] = "MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS"
	enumLookup[UNPACK_SKIP_VOLUMES_SGIS] = "UNPACK_SKIP_VOLUMES_SGIS"
	enumLookup[TRANSPOSE_NV] = "TRANSPOSE_NV"
	enumLookup[UNIFORM_TYPE] = "UNIFORM_TYPE"
	enumLookup[UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT] = "UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT"
	enumLookup[FLOAT] = "FLOAT"
	enumLookup[CURRENT_MATRIX_STACK_DEPTH_NV] = "CURRENT_MATRIX_STACK_DEPTH_NV"
	enumLookup[PIXEL_COUNT_AVAILABLE_NV] = "PIXEL_COUNT_AVAILABLE_NV"
	enumLookup[COLOR_ATTACHMENT15] = "COLOR_ATTACHMENT15"
	enumLookup[FIRST_TO_REST_NV] = "FIRST_TO_REST_NV"
	enumLookup[TRIANGLES_ADJACENCY_ARB] = "TRIANGLES_ADJACENCY_ARB"
	enumLookup[PROXY_COLOR_TABLE_SGI] = "PROXY_COLOR_TABLE_SGI"
	enumLookup[T2F_IUI_N3F_V3F_EXT] = "T2F_IUI_N3F_V3F_EXT"
	enumLookup[TEXTURE_BINDING_BUFFER_EXT] = "TEXTURE_BINDING_BUFFER_EXT"
	enumLookup[RGB8I_EXT] = "RGB8I_EXT"
	enumLookup[POINT_SIZE_MIN_EXT] = "POINT_SIZE_MIN_EXT"
	enumLookup[R8UI] = "R8UI"
	enumLookup[MAP1_VERTEX_ATTRIB8_4_NV] = "MAP1_VERTEX_ATTRIB8_4_NV"
	enumLookup[BUFFER_ACCESS_ARB] = "BUFFER_ACCESS_ARB"
	enumLookup[ACTIVE_VARYINGS_NV] = "ACTIVE_VARYINGS_NV"
	enumLookup[MAP2_VERTEX_ATTRIB7_4_NV] = "MAP2_VERTEX_ATTRIB7_4_NV"
	enumLookup[READ_PIXEL_DATA_RANGE_LENGTH_NV] = "READ_PIXEL_DATA_RANGE_LENGTH_NV"
	enumLookup[BGR_INTEGER_EXT] = "BGR_INTEGER_EXT"
	enumLookup[LUMINANCE12_ALPHA4] = "LUMINANCE12_ALPHA4"
	enumLookup[UNSIGNED_SHORT_1_5_5_5_REV_EXT] = "UNSIGNED_SHORT_1_5_5_5_REV_EXT"
	enumLookup[MAX_3D_TEXTURE_SIZE_EXT] = "MAX_3D_TEXTURE_SIZE_EXT"
	enumLookup[POST_COLOR_MATRIX_GREEN_BIAS] = "POST_COLOR_MATRIX_GREEN_BIAS"
	enumLookup[PRIMARY_COLOR_NV] = "PRIMARY_COLOR_NV"
	enumLookup[YCBCR_422_APPLE] = "YCBCR_422_APPLE"
	enumLookup[VARIANT_DATATYPE_EXT] = "VARIANT_DATATYPE_EXT"
	enumLookup[REFERENCED_BY_TESS_EVALUATION_SHADER] = "REFERENCED_BY_TESS_EVALUATION_SHADER"
	enumLookup[OUTPUT_TEXTURE_COORD25_EXT] = "OUTPUT_TEXTURE_COORD25_EXT"
	enumLookup[COLOR_ATTACHMENT9_EXT] = "COLOR_ATTACHMENT9_EXT"
	enumLookup[SYNC_STATUS_APPLE] = "SYNC_STATUS_APPLE"
	enumLookup[TEXTURE_GEN_Q] = "TEXTURE_GEN_Q"
	enumLookup[FRAGMENTS_INSTRUMENT_SGIX] = "FRAGMENTS_INSTRUMENT_SGIX"
	enumLookup[UNPACK_CLIENT_STORAGE_APPLE] = "UNPACK_CLIENT_STORAGE_APPLE"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR"
	enumLookup[TEXTURE_BASE_LEVEL_SGIS] = "TEXTURE_BASE_LEVEL_SGIS"
	enumLookup[HALF_BIAS_NEGATE_NV] = "HALF_BIAS_NEGATE_NV"
	enumLookup[OBJECT_BUFFER_USAGE_ATI] = "OBJECT_BUFFER_USAGE_ATI"
	enumLookup[MATRIX19_ARB] = "MATRIX19_ARB"
	enumLookup[MAX_PROGRAM_ATTRIB_COMPONENTS_NV] = "MAX_PROGRAM_ATTRIB_COMPONENTS_NV"
	enumLookup[DT_BIAS_NV] = "DT_BIAS_NV"
	enumLookup[COLOR_TABLE_WIDTH] = "COLOR_TABLE_WIDTH"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Z_OES] = "TEXTURE_CUBE_MAP_POSITIVE_Z_OES"
	enumLookup[RGB_FLOAT16_ATI] = "RGB_FLOAT16_ATI"
	enumLookup[BGRA_INTEGER_EXT] = "BGRA_INTEGER_EXT"
	enumLookup[SPHERE_MAP] = "SPHERE_MAP"
	enumLookup[SEPARABLE_2D_EXT] = "SEPARABLE_2D_EXT"
	enumLookup[COLOR_INDEX8_EXT] = "COLOR_INDEX8_EXT"
	enumLookup[FOG_OFFSET_SGIX] = "FOG_OFFSET_SGIX"
	enumLookup[MAP2_VERTEX_ATTRIB0_4_NV] = "MAP2_VERTEX_ATTRIB0_4_NV"
	enumLookup[ELEMENT_ARRAY_BUFFER_BINDING_ARB] = "ELEMENT_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[PIXEL_PACK_BUFFER_BINDING_EXT] = "PIXEL_PACK_BUFFER_BINDING_EXT"
	enumLookup[INT64_VEC2_NV] = "INT64_VEC2_NV"
	enumLookup[COMPRESSED_RGBA_ASTC_8x5_KHR] = "COMPRESSED_RGBA_ASTC_8x5_KHR"
	enumLookup[RGB5_A1] = "RGB5_A1"
	enumLookup[CONSTANT_COLOR] = "CONSTANT_COLOR"
	enumLookup[SAMPLE_MASK_INVERT_SGIS] = "SAMPLE_MASK_INVERT_SGIS"
	enumLookup[TEXTURE_GEQUAL_R_SGIX] = "TEXTURE_GEQUAL_R_SGIX"
	enumLookup[R8I] = "R8I"
	enumLookup[GEOMETRY_INPUT_TYPE_EXT] = "GEOMETRY_INPUT_TYPE_EXT"
	enumLookup[ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS] = "ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS"
	enumLookup[READ_BUFFER_NV] = "READ_BUFFER_NV"
	enumLookup[GL_2PASS_0_SGIS] = "GL_2PASS_0_SGIS"
	enumLookup[TEXTURE1_ARB] = "TEXTURE1_ARB"
	enumLookup[PATCH_DEFAULT_OUTER_LEVEL] = "PATCH_DEFAULT_OUTER_LEVEL"
	enumLookup[ATOMIC_COUNTER_BUFFER] = "ATOMIC_COUNTER_BUFFER"
	enumLookup[INVALID_ENUM] = "INVALID_ENUM"
	enumLookup[LUMINANCE4] = "LUMINANCE4"
	enumLookup[CLIP_PLANE3] = "CLIP_PLANE3"
	enumLookup[BLEND_EQUATION_RGB_EXT] = "BLEND_EQUATION_RGB_EXT"
	enumLookup[COORD_REPLACE_ARB] = "COORD_REPLACE_ARB"
	enumLookup[PROGRAM_UNDER_NATIVE_LIMITS_ARB] = "PROGRAM_UNDER_NATIVE_LIMITS_ARB"
	enumLookup[DEPTH32F_STENCIL8_NV] = "DEPTH32F_STENCIL8_NV"
	enumLookup[TEXTURE_HEIGHT] = "TEXTURE_HEIGHT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE] = "FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE"
	enumLookup[MAT_DIFFUSE_BIT_PGI] = "MAT_DIFFUSE_BIT_PGI"
	enumLookup[MATRIX_INDEX_ARRAY_SIZE_OES] = "MATRIX_INDEX_ARRAY_SIZE_OES"
	enumLookup[BUFFER_FLUSHING_UNMAP_APPLE] = "BUFFER_FLUSHING_UNMAP_APPLE"
	enumLookup[GEOMETRY_PROGRAM_NV] = "GEOMETRY_PROGRAM_NV"
	enumLookup[RGB32UI] = "RGB32UI"
	enumLookup[QUAD_STRIP] = "QUAD_STRIP"
	enumLookup[MAX_VERTEX_SHADER_VARIANTS_EXT] = "MAX_VERTEX_SHADER_VARIANTS_EXT"
	enumLookup[RGBA32F] = "RGBA32F"
	enumLookup[DRAW_BUFFER13_ARB] = "DRAW_BUFFER13_ARB"
	enumLookup[VERTEX_ATTRIB_MAP2_SIZE_APPLE] = "VERTEX_ATTRIB_MAP2_SIZE_APPLE"
	enumLookup[ACTIVE_ATTRIBUTES] = "ACTIVE_ATTRIBUTES"
	enumLookup[UNSIGNED_INT_VEC4_EXT] = "UNSIGNED_INT_VEC4_EXT"
	enumLookup[MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS] = "MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS"
	enumLookup[LIST_INDEX] = "LIST_INDEX"
	enumLookup[RGB4] = "RGB4"
	enumLookup[SHARED_TEXTURE_PALETTE_EXT] = "SHARED_TEXTURE_PALETTE_EXT"
	enumLookup[TEXTURE_UNSIGNED_REMAP_MODE_NV] = "TEXTURE_UNSIGNED_REMAP_MODE_NV"
	enumLookup[DYNAMIC_COPY] = "DYNAMIC_COPY"
	enumLookup[PROGRAM_OBJECT_ARB] = "PROGRAM_OBJECT_ARB"
	enumLookup[MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB] = "MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB"
	enumLookup[IMAGE_BINDING_LEVEL_EXT] = "IMAGE_BINDING_LEVEL_EXT"
	enumLookup[RGB10] = "RGB10"
	enumLookup[TEXTURE_BINDING_CUBE_MAP_ARB] = "TEXTURE_BINDING_CUBE_MAP_ARB"
	enumLookup[SOURCE2_RGB_EXT] = "SOURCE2_RGB_EXT"
	enumLookup[CON_20_ATI] = "CON_20_ATI"
	enumLookup[PALETTE4_RGB5_A1_OES] = "PALETTE4_RGB5_A1_OES"
	enumLookup[TEXTURE_1D_ARRAY_EXT] = "TEXTURE_1D_ARRAY_EXT"
	enumLookup[MAX_COMPUTE_IMAGE_UNIFORMS] = "MAX_COMPUTE_IMAGE_UNIFORMS"
	enumLookup[COLOR_ARRAY_SIZE_EXT] = "COLOR_ARRAY_SIZE_EXT"
	enumLookup[NORMAL_ARRAY_POINTER_EXT] = "NORMAL_ARRAY_POINTER_EXT"
	enumLookup[TEXTURE_MAG_FILTER] = "TEXTURE_MAG_FILTER"
	enumLookup[INTENSITY16_EXT] = "INTENSITY16_EXT"
	enumLookup[COLOR_INDEX16_EXT] = "COLOR_INDEX16_EXT"
	enumLookup[INT_SAMPLER_BUFFER] = "INT_SAMPLER_BUFFER"
	enumLookup[TESS_GEN_POINT_MODE] = "TESS_GEN_POINT_MODE"
	enumLookup[ALWAYS_SOFT_HINT_PGI] = "ALWAYS_SOFT_HINT_PGI"
	enumLookup[UNSIGNED_SHORT_1_5_5_5_REV] = "UNSIGNED_SHORT_1_5_5_5_REV"
	enumLookup[ARRAY_STRIDE] = "ARRAY_STRIDE"
	enumLookup[CONVOLUTION_FORMAT] = "CONVOLUTION_FORMAT"
	enumLookup[POINT_FADE_THRESHOLD_SIZE_EXT] = "POINT_FADE_THRESHOLD_SIZE_EXT"
	enumLookup[DEPTH_PASS_INSTRUMENT_MAX_SGIX] = "DEPTH_PASS_INSTRUMENT_MAX_SGIX"
	enumLookup[RGBA4_S3TC] = "RGBA4_S3TC"
	enumLookup[COMBINER_CD_DOT_PRODUCT_NV] = "COMBINER_CD_DOT_PRODUCT_NV"
	enumLookup[MATRIX6_ARB] = "MATRIX6_ARB"
	enumLookup[SEPARATE_ATTRIBS_EXT] = "SEPARATE_ATTRIBS_EXT"
	enumLookup[LAST_VERTEX_CONVENTION] = "LAST_VERTEX_CONVENTION"
	enumLookup[VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT] = "VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT"
	enumLookup[UNSIGNED_INT_8_8_8_8_REV_EXT] = "UNSIGNED_INT_8_8_8_8_REV_EXT"
	enumLookup[COLOR_TABLE_FORMAT] = "COLOR_TABLE_FORMAT"
	enumLookup[PACK_MAX_COMPRESSED_SIZE_SGIX] = "PACK_MAX_COMPRESSED_SIZE_SGIX"
	enumLookup[QUADS] = "QUADS"
	enumLookup[MIRROR_CLAMP_TO_EDGE_ATI] = "MIRROR_CLAMP_TO_EDGE_ATI"
	enumLookup[DRAW_BUFFER5_ARB] = "DRAW_BUFFER5_ARB"
	enumLookup[UNSIGNED_INT_SAMPLER_BUFFER] = "UNSIGNED_INT_SAMPLER_BUFFER"
	enumLookup[CMYKA_EXT] = "CMYKA_EXT"
	enumLookup[ATTENUATION_EXT] = "ATTENUATION_EXT"
	enumLookup[Z400_BINARY_AMD] = "Z400_BINARY_AMD"
	enumLookup[DRAW_BUFFER7_ARB] = "DRAW_BUFFER7_ARB"
	enumLookup[FLOAT_RG_NV] = "FLOAT_RG_NV"
	enumLookup[HALF_FLOAT_OES] = "HALF_FLOAT_OES"
	enumLookup[TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV] = "TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV"
	enumLookup[POST_COLOR_MATRIX_BLUE_BIAS] = "POST_COLOR_MATRIX_BLUE_BIAS"
	enumLookup[COLOR_ATTACHMENT14_NV] = "COLOR_ATTACHMENT14_NV"
	enumLookup[INT_IMAGE_3D] = "INT_IMAGE_3D"
	enumLookup[DEBUG_SEVERITY_HIGH] = "DEBUG_SEVERITY_HIGH"
	enumLookup[NUM_EXTENSIONS] = "NUM_EXTENSIONS"
	enumLookup[PROGRAM_ERROR_POSITION_NV] = "PROGRAM_ERROR_POSITION_NV"
	enumLookup[TRACE_PRIMITIVES_BIT_MESA] = "TRACE_PRIMITIVES_BIT_MESA"
	enumLookup[PROGRAM_RESULT_COMPONENTS_NV] = "PROGRAM_RESULT_COMPONENTS_NV"
	enumLookup[CLAMP_FRAGMENT_COLOR_ARB] = "CLAMP_FRAGMENT_COLOR_ARB"
	enumLookup[TRANSFORM_FEEDBACK_RECORD_NV] = "TRANSFORM_FEEDBACK_RECORD_NV"
	enumLookup[POST_CONVOLUTION_BLUE_BIAS_EXT] = "POST_CONVOLUTION_BLUE_BIAS_EXT"
	enumLookup[RGB12_EXT] = "RGB12_EXT"
	enumLookup[REPLICATE_BORDER] = "REPLICATE_BORDER"
	enumLookup[FRAMEBUFFER_ATTACHMENT_GREEN_SIZE] = "FRAMEBUFFER_ATTACHMENT_GREEN_SIZE"
	enumLookup[EVAL_VERTEX_ATTRIB4_NV] = "EVAL_VERTEX_ATTRIB4_NV"
	enumLookup[DRAW_BUFFER9] = "DRAW_BUFFER9"
	enumLookup[CURRENT_PALETTE_MATRIX_OES] = "CURRENT_PALETTE_MATRIX_OES"
	enumLookup[MAX_TEXTURE_COORDS_NV] = "MAX_TEXTURE_COORDS_NV"
	enumLookup[READ_ONLY] = "READ_ONLY"
	enumLookup[UNIFORM_NAME_LENGTH] = "UNIFORM_NAME_LENGTH"
	enumLookup[MAX_GEOMETRY_OUTPUT_VERTICES_EXT] = "MAX_GEOMETRY_OUTPUT_VERTICES_EXT"
	enumLookup[INT_IMAGE_2D_EXT] = "INT_IMAGE_2D_EXT"
	enumLookup[VARIABLE_C_NV] = "VARIABLE_C_NV"
	enumLookup[CLAMP_READ_COLOR_ARB] = "CLAMP_READ_COLOR_ARB"
	enumLookup[IMAGE_BUFFER_EXT] = "IMAGE_BUFFER_EXT"
	enumLookup[DEPTH_COMPONENT32_SGIX] = "DEPTH_COMPONENT32_SGIX"
	enumLookup[PRIMITIVE_RESTART_INDEX_NV] = "PRIMITIVE_RESTART_INDEX_NV"
	enumLookup[ATTRIB_ARRAY_SIZE_NV] = "ATTRIB_ARRAY_SIZE_NV"
	enumLookup[STREAM_COPY_ARB] = "STREAM_COPY_ARB"
	enumLookup[MAX_TEXTURE_LOD_BIAS_EXT] = "MAX_TEXTURE_LOD_BIAS_EXT"
	enumLookup[REG_9_ATI] = "REG_9_ATI"
	enumLookup[Z4Y12Z4CB12Z4Y12Z4CR12_422_NV] = "Z4Y12Z4CB12Z4Y12Z4CR12_422_NV"
	enumLookup[EXT_shared_texture_palette] = "EXT_shared_texture_palette"
	enumLookup[ZERO] = "ZERO"
	enumLookup[MAX_VIEWPORTS] = "MAX_VIEWPORTS"
	enumLookup[STENCIL_RENDERABLE] = "STENCIL_RENDERABLE"
	enumLookup[RGBA8I_EXT] = "RGBA8I_EXT"
	enumLookup[FRAMEBUFFER_SRGB] = "FRAMEBUFFER_SRGB"
	enumLookup[PATH_OBJECT_BOUNDING_BOX_NV] = "PATH_OBJECT_BOUNDING_BOX_NV"
	enumLookup[GL_3D] = "GL_3D"
	enumLookup[MAX_PROGRAM_PATCH_ATTRIBS_NV] = "MAX_PROGRAM_PATCH_ATTRIBS_NV"
	enumLookup[DRAW_BUFFER14_ATI] = "DRAW_BUFFER14_ATI"
	enumLookup[MAX_GEOMETRY_TEXTURE_IMAGE_UNITS] = "MAX_GEOMETRY_TEXTURE_IMAGE_UNITS"
	enumLookup[COLOR_ATTACHMENT0] = "COLOR_ATTACHMENT0"
	enumLookup[PRIMITIVE_RESTART_INDEX] = "PRIMITIVE_RESTART_INDEX"
	enumLookup[TRIANGLE_FAN] = "TRIANGLE_FAN"
	enumLookup[PIXEL_TRANSFORM_2D_EXT] = "PIXEL_TRANSFORM_2D_EXT"
	enumLookup[DOT3_RGBA] = "DOT3_RGBA"
	enumLookup[SGIS_texture_lod] = "SGIS_texture_lod"
	enumLookup[UNSIGNED_SHORT_4_4_4_4_REV] = "UNSIGNED_SHORT_4_4_4_4_REV"
	enumLookup[TEXTURE21_ARB] = "TEXTURE21_ARB"
	enumLookup[LUMINANCE_ALPHA_FLOAT32_APPLE] = "LUMINANCE_ALPHA_FLOAT32_APPLE"
	enumLookup[PIXEL_COUNTER_BITS_NV] = "PIXEL_COUNTER_BITS_NV"
	enumLookup[GCCSO_SHADER_BINARY_FJ] = "GCCSO_SHADER_BINARY_FJ"
	enumLookup[FOG_HINT] = "FOG_HINT"
	enumLookup[SYNC_CL_EVENT_COMPLETE_ARB] = "SYNC_CL_EVENT_COMPLETE_ARB"
	enumLookup[MAX_HEIGHT] = "MAX_HEIGHT"
	enumLookup[DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV] = "DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV"
	enumLookup[UNSIGNED_NORMALIZED_ARB] = "UNSIGNED_NORMALIZED_ARB"
	enumLookup[RENDERBUFFER_BLUE_SIZE] = "RENDERBUFFER_BLUE_SIZE"
	enumLookup[BACK_RIGHT] = "BACK_RIGHT"
	enumLookup[LINEAR_SHARPEN_SGIS] = "LINEAR_SHARPEN_SGIS"
	enumLookup[COMPRESSED_RED] = "COMPRESSED_RED"
	enumLookup[STREAM_READ_ARB] = "STREAM_READ_ARB"
	enumLookup[PIXEL_PACK_BUFFER] = "PIXEL_PACK_BUFFER"
	enumLookup[EIGHTH_BIT_ATI] = "EIGHTH_BIT_ATI"
	enumLookup[RGB16I_EXT] = "RGB16I_EXT"
	enumLookup[CULL_FACE] = "CULL_FACE"
	enumLookup[MAP2_TEXTURE_COORD_3] = "MAP2_TEXTURE_COORD_3"
	enumLookup[MAX_CONVOLUTION_WIDTH_EXT] = "MAX_CONVOLUTION_WIDTH_EXT"
	enumLookup[SRGB_WRITE] = "SRGB_WRITE"
	enumLookup[COMPRESSED_RGBA_S3TC_DXT1_EXT] = "COMPRESSED_RGBA_S3TC_DXT1_EXT"
	enumLookup[MATRIX27_ARB] = "MATRIX27_ARB"
	enumLookup[SWIZZLE_STR_DR_ATI] = "SWIZZLE_STR_DR_ATI"
	enumLookup[TEXTURE_BINDING_1D_ARRAY_EXT] = "TEXTURE_BINDING_1D_ARRAY_EXT"
	enumLookup[COMPRESSED_RGBA_ASTC_5x4_KHR] = "COMPRESSED_RGBA_ASTC_5x4_KHR"
	enumLookup[FRAMEBUFFER_UNDEFINED] = "FRAMEBUFFER_UNDEFINED"
	enumLookup[TEXTURE_COMPARE_MODE_EXT] = "TEXTURE_COMPARE_MODE_EXT"
	enumLookup[COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT] = "COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT"
	enumLookup[MAX_DEEP_3D_TEXTURE_DEPTH_NV] = "MAX_DEEP_3D_TEXTURE_DEPTH_NV"
	enumLookup[EVAL_VERTEX_ATTRIB13_NV] = "EVAL_VERTEX_ATTRIB13_NV"
	enumLookup[TEXTURE_1D_ARRAY] = "TEXTURE_1D_ARRAY"
	enumLookup[COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT] = "COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT"
	enumLookup[COLOR_ATTACHMENT9] = "COLOR_ATTACHMENT9"
	enumLookup[T] = "T"
	enumLookup[FLOAT_32_UNSIGNED_INT_24_8_REV_NV] = "FLOAT_32_UNSIGNED_INT_24_8_REV_NV"
	enumLookup[USE_MISSING_GLYPH_NV] = "USE_MISSING_GLYPH_NV"
	enumLookup[ARB_imaging] = "ARB_imaging"
	enumLookup[LOAD] = "LOAD"
	enumLookup[POINT_SMOOTH_HINT] = "POINT_SMOOTH_HINT"
	enumLookup[MAP_COLOR] = "MAP_COLOR"
	enumLookup[CONVOLUTION_1D] = "CONVOLUTION_1D"
	enumLookup[COLOR_ARRAY_COUNT_EXT] = "COLOR_ARRAY_COUNT_EXT"
	enumLookup[DEBUG_TYPE_OTHER_ARB] = "DEBUG_TYPE_OTHER_ARB"
	enumLookup[RGB_SCALE] = "RGB_SCALE"
	enumLookup[MODELVIEW19_ARB] = "MODELVIEW19_ARB"
	enumLookup[VERTEX_STREAM5_ATI] = "VERTEX_STREAM5_ATI"
	enumLookup[COMPRESSED_SIGNED_RG11_EAC] = "COMPRESSED_SIGNED_RG11_EAC"
	enumLookup[VERSION_2_1] = "VERSION_2_1"
	enumLookup[EXT_subtexture] = "EXT_subtexture"
	enumLookup[IMAGE_CLASS_1_X_16] = "IMAGE_CLASS_1_X_16"
	enumLookup[COMBINE_ALPHA_ARB] = "COMBINE_ALPHA_ARB"
	enumLookup[COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT] = "COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT"
	enumLookup[FLOAT16_VEC4_NV] = "FLOAT16_VEC4_NV"
	enumLookup[PROXY_TEXTURE_2D_STACK_MESAX] = "PROXY_TEXTURE_2D_STACK_MESAX"
	enumLookup[EXT_cmyka] = "EXT_cmyka"
	enumLookup[TEXTURE_WRAP_R_OES] = "TEXTURE_WRAP_R_OES"
	enumLookup[R1UI_C4UB_V3F_SUN] = "R1UI_C4UB_V3F_SUN"
	enumLookup[OP_DOT3_EXT] = "OP_DOT3_EXT"
	enumLookup[UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT] = "UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT"
	enumLookup[COMPRESSED_RGBA_ASTC_12x10_KHR] = "COMPRESSED_RGBA_ASTC_12x10_KHR"
	enumLookup[STRICT_LIGHTING_HINT_PGI] = "STRICT_LIGHTING_HINT_PGI"
	enumLookup[SGIX_async] = "SGIX_async"
	enumLookup[DUAL_LUMINANCE16_SGIS] = "DUAL_LUMINANCE16_SGIS"
	enumLookup[UNIFORM_BUFFER_BINDING] = "UNIFORM_BUFFER_BINDING"
	enumLookup[INDEX_CLEAR_VALUE] = "INDEX_CLEAR_VALUE"
	enumLookup[FUNC_ADD_EXT] = "FUNC_ADD_EXT"
	enumLookup[CONVOLUTION_FORMAT_EXT] = "CONVOLUTION_FORMAT_EXT"
	enumLookup[CURRENT_SECONDARY_COLOR_EXT] = "CURRENT_SECONDARY_COLOR_EXT"
	enumLookup[MAP2_VERTEX_ATTRIB11_4_NV] = "MAP2_VERTEX_ATTRIB11_4_NV"
	enumLookup[VERTEX_ATTRIB_MAP2_ORDER_APPLE] = "VERTEX_ATTRIB_MAP2_ORDER_APPLE"
	enumLookup[ACTIVE_UNIFORM_MAX_LENGTH] = "ACTIVE_UNIFORM_MAX_LENGTH"
	enumLookup[TEXTURE_COORD_ARRAY_SIZE] = "TEXTURE_COORD_ARRAY_SIZE"
	enumLookup[VERTEX_ARRAY_COUNT_EXT] = "VERTEX_ARRAY_COUNT_EXT"
	enumLookup[TEXTURE_CUBE_MAP_OES] = "TEXTURE_CUBE_MAP_OES"
	enumLookup[TEXTURE_WIDTH_QCOM] = "TEXTURE_WIDTH_QCOM"
	enumLookup[TEXTURE_OBJECT_VALID_QCOM] = "TEXTURE_OBJECT_VALID_QCOM"
	enumLookup[DEPTH_STENCIL_TO_RGBA_NV] = "DEPTH_STENCIL_TO_RGBA_NV"
	enumLookup[LINE_TO_NV] = "LINE_TO_NV"
	enumLookup[EXT_abgr] = "EXT_abgr"
	enumLookup[POINT_FADE_THRESHOLD_SIZE_SGIS] = "POINT_FADE_THRESHOLD_SIZE_SGIS"
	enumLookup[GREEN_MAX_CLAMP_INGR] = "GREEN_MAX_CLAMP_INGR"
	enumLookup[SRC1_ALPHA] = "SRC1_ALPHA"
	enumLookup[VARIANT_ARRAY_STRIDE_EXT] = "VARIANT_ARRAY_STRIDE_EXT"
	enumLookup[COLOR_ARRAY_BUFFER_BINDING_ARB] = "COLOR_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[MAX_PROGRAM_TEXEL_OFFSET] = "MAX_PROGRAM_TEXEL_OFFSET"
	enumLookup[MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB] = "MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB"
	enumLookup[SHADER_COMPILER] = "SHADER_COMPILER"
	enumLookup[MAX_TESS_CONTROL_OUTPUT_COMPONENTS] = "MAX_TESS_CONTROL_OUTPUT_COMPONENTS"
	enumLookup[MALI_PROGRAM_BINARY_ARM] = "MALI_PROGRAM_BINARY_ARM"
	enumLookup[COMPUTE_PROGRAM_NV] = "COMPUTE_PROGRAM_NV"
	enumLookup[MAX_ATOMIC_COUNTER_BUFFER_SIZE] = "MAX_ATOMIC_COUNTER_BUFFER_SIZE"
	enumLookup[DEBUG_SOURCE_APPLICATION] = "DEBUG_SOURCE_APPLICATION"
	enumLookup[OP_CROSS_PRODUCT_EXT] = "OP_CROSS_PRODUCT_EXT"
	enumLookup[PROGRAM_TEX_INDIRECTIONS_ARB] = "PROGRAM_TEX_INDIRECTIONS_ARB"
	enumLookup[REG_30_ATI] = "REG_30_ATI"
	enumLookup[AUTO_GENERATE_MIPMAP] = "AUTO_GENERATE_MIPMAP"
	enumLookup[MAX_RENDERBUFFER_SIZE_EXT] = "MAX_RENDERBUFFER_SIZE_EXT"
	enumLookup[STENCIL_CLEAR_VALUE] = "STENCIL_CLEAR_VALUE"
	enumLookup[VERTEX_BINDING_DIVISOR] = "VERTEX_BINDING_DIVISOR"
	enumLookup[UNIFORM_BUFFER_SIZE] = "UNIFORM_BUFFER_SIZE"
	enumLookup[UNSIGNED_INT_10F_11F_11F_REV] = "UNSIGNED_INT_10F_11F_11F_REV"
	enumLookup[COLOR_ATTACHMENT6] = "COLOR_ATTACHMENT6"
	enumLookup[IMAGE_2D_EXT] = "IMAGE_2D_EXT"
	enumLookup[TIMEOUT_IGNORED] = "TIMEOUT_IGNORED"
	enumLookup[GL_2PASS_0_EXT] = "GL_2PASS_0_EXT"
	enumLookup[MAX_TEXTURE_IMAGE_UNITS] = "MAX_TEXTURE_IMAGE_UNITS"
	enumLookup[MAX_PROGRAM_GENERIC_RESULTS_NV] = "MAX_PROGRAM_GENERIC_RESULTS_NV"
	enumLookup[COMMAND_BARRIER_BIT] = "COMMAND_BARRIER_BIT"
	enumLookup[VERTEX_WEIGHTING_EXT] = "VERTEX_WEIGHTING_EXT"
	enumLookup[CONSTANT_COLOR0_NV] = "CONSTANT_COLOR0_NV"
	enumLookup[FLOAT_MAT4_ARB] = "FLOAT_MAT4_ARB"
	enumLookup[QUERY_OBJECT_AMD] = "QUERY_OBJECT_AMD"
	enumLookup[PACK_REVERSE_ROW_ORDER_ANGLE] = "PACK_REVERSE_ROW_ORDER_ANGLE"
	enumLookup[CONTEXT_FLAG_DEBUG_BIT] = "CONTEXT_FLAG_DEBUG_BIT"
	enumLookup[RGBA] = "RGBA"
	enumLookup[HISTOGRAM] = "HISTOGRAM"
	enumLookup[IMAGE_TRANSFORM_2D_HP] = "IMAGE_TRANSFORM_2D_HP"
	enumLookup[VIEW_CLASS_8_BITS] = "VIEW_CLASS_8_BITS"
	enumLookup[OP_FLOOR_EXT] = "OP_FLOOR_EXT"
	enumLookup[BUFFER_MAP_POINTER] = "BUFFER_MAP_POINTER"
	enumLookup[UNIFORM_BLOCK_BINDING] = "UNIFORM_BLOCK_BINDING"
	enumLookup[MAX_TRANSFORM_FEEDBACK_BUFFERS] = "MAX_TRANSFORM_FEEDBACK_BUFFERS"
	enumLookup[DUAL_ALPHA16_SGIS] = "DUAL_ALPHA16_SGIS"
	enumLookup[MAP1_VERTEX_ATTRIB14_4_NV] = "MAP1_VERTEX_ATTRIB14_4_NV"
	enumLookup[MAX_PALETTE_MATRICES_ARB] = "MAX_PALETTE_MATRICES_ARB"
	enumLookup[CON_19_ATI] = "CON_19_ATI"
	enumLookup[POINT_SIZE_ARRAY_POINTER_OES] = "POINT_SIZE_ARRAY_POINTER_OES"
	enumLookup[QUERY_WAIT_NV] = "QUERY_WAIT_NV"
	enumLookup[TEXTURE_BINDING_CUBE_MAP_ARRAY] = "TEXTURE_BINDING_CUBE_MAP_ARRAY"
	enumLookup[POST_COLOR_MATRIX_BLUE_BIAS_SGI] = "POST_COLOR_MATRIX_BLUE_BIAS_SGI"
	enumLookup[PROXY_POST_CONVOLUTION_COLOR_TABLE] = "PROXY_POST_CONVOLUTION_COLOR_TABLE"
	enumLookup[ALPHA16I_EXT] = "ALPHA16I_EXT"
	enumLookup[RELATIVE_LINE_TO_NV] = "RELATIVE_LINE_TO_NV"
	enumLookup[SCISSOR_BIT] = "SCISSOR_BIT"
	enumLookup[VERTEX_ARRAY_STRIDE] = "VERTEX_ARRAY_STRIDE"
	enumLookup[PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX] = "PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX"
	enumLookup[LIGHT_MODEL_COLOR_CONTROL_EXT] = "LIGHT_MODEL_COLOR_CONTROL_EXT"
	enumLookup[CURRENT_FOG_COORDINATE] = "CURRENT_FOG_COORDINATE"
	enumLookup[MAX_DEBUG_MESSAGE_LENGTH_AMD] = "MAX_DEBUG_MESSAGE_LENGTH_AMD"
	enumLookup[REFERENCED_BY_TESS_CONTROL_SHADER] = "REFERENCED_BY_TESS_CONTROL_SHADER"
	enumLookup[UNSIGNED_BYTE_2_3_3_REV] = "UNSIGNED_BYTE_2_3_3_REV"
	enumLookup[R1UI_N3F_V3F_SUN] = "R1UI_N3F_V3F_SUN"
	enumLookup[TESS_CONTROL_PROGRAM_NV] = "TESS_CONTROL_PROGRAM_NV"
	enumLookup[ALPHA_INTEGER] = "ALPHA_INTEGER"
	enumLookup[TRANSLATE_Y_NV] = "TRANSLATE_Y_NV"
	enumLookup[LINE] = "LINE"
	enumLookup[FOG_DISTANCE_MODE_NV] = "FOG_DISTANCE_MODE_NV"
	enumLookup[EVAL_FRACTIONAL_TESSELLATION_NV] = "EVAL_FRACTIONAL_TESSELLATION_NV"
	enumLookup[UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT] = "UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT"
	enumLookup[RGBA16_SNORM] = "RGBA16_SNORM"
	enumLookup[COUNT_DOWN_NV] = "COUNT_DOWN_NV"
	enumLookup[MAX_TESS_CONTROL_IMAGE_UNIFORMS] = "MAX_TESS_CONTROL_IMAGE_UNIFORMS"
	enumLookup[MAT_SHININESS_BIT_PGI] = "MAT_SHININESS_BIT_PGI"
	enumLookup[TEXTURE_MAX_CLAMP_R_SGIX] = "TEXTURE_MAX_CLAMP_R_SGIX"
	enumLookup[TEXTURE11] = "TEXTURE11"
	enumLookup[TEXTURE_MAX_ANISOTROPY_EXT] = "TEXTURE_MAX_ANISOTROPY_EXT"
	enumLookup[INTENSITY8_SNORM] = "INTENSITY8_SNORM"
	enumLookup[LINE_WIDTH_RANGE] = "LINE_WIDTH_RANGE"
	enumLookup[TEXTURE9] = "TEXTURE9"
	enumLookup[TEXTURE_LOD_BIAS_EXT] = "TEXTURE_LOD_BIAS_EXT"
	enumLookup[OUTPUT_TEXTURE_COORD23_EXT] = "OUTPUT_TEXTURE_COORD23_EXT"
	enumLookup[SGIX_convolution_accuracy] = "SGIX_convolution_accuracy"
	enumLookup[IMAGE_PIXEL_FORMAT] = "IMAGE_PIXEL_FORMAT"
	enumLookup[TEXTURE_RANGE_LENGTH_APPLE] = "TEXTURE_RANGE_LENGTH_APPLE"
	enumLookup[DOT_PRODUCT_TEXTURE_RECTANGLE_NV] = "DOT_PRODUCT_TEXTURE_RECTANGLE_NV"
	enumLookup[OFFSET_TEXTURE_2D_BIAS_NV] = "OFFSET_TEXTURE_2D_BIAS_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_FORMATS_OES] = "FRAMEBUFFER_INCOMPLETE_FORMATS_OES"
	enumLookup[FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB] = "FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB"
	enumLookup[IMAGE_2D_RECT_EXT] = "IMAGE_2D_RECT_EXT"
	enumLookup[CLIENT_PIXEL_STORE_BIT] = "CLIENT_PIXEL_STORE_BIT"
	enumLookup[R16F] = "R16F"
	enumLookup[VIEW_CLASS_96_BITS] = "VIEW_CLASS_96_BITS"
	enumLookup[PIXEL_SUBSAMPLE_4242_SGIX] = "PIXEL_SUBSAMPLE_4242_SGIX"
	enumLookup[OFFSET_TEXTURE_MATRIX_NV] = "OFFSET_TEXTURE_MATRIX_NV"
	enumLookup[CPU_OPTIMIZED_QCOM] = "CPU_OPTIMIZED_QCOM"
	enumLookup[MODELVIEW_MATRIX] = "MODELVIEW_MATRIX"
	enumLookup[SPOT_EXPONENT] = "SPOT_EXPONENT"
	enumLookup[MINMAX] = "MINMAX"
	enumLookup[EXT_texture_object] = "EXT_texture_object"
	enumLookup[PROGRAM_SEPARABLE_EXT] = "PROGRAM_SEPARABLE_EXT"
	enumLookup[SYNC_CONDITION_APPLE] = "SYNC_CONDITION_APPLE"
	enumLookup[SYNC_CL_EVENT_ARB] = "SYNC_CL_EVENT_ARB"
	enumLookup[DECR_WRAP_EXT] = "DECR_WRAP_EXT"
	enumLookup[DRAW_BUFFER14_ARB] = "DRAW_BUFFER14_ARB"
	enumLookup[RG_SNORM] = "RG_SNORM"
	enumLookup[MAX_GEOMETRY_INPUT_COMPONENTS] = "MAX_GEOMETRY_INPUT_COMPONENTS"
	enumLookup[TRIANGLE_STRIP_ADJACENCY_ARB] = "TRIANGLE_STRIP_ADJACENCY_ARB"
	enumLookup[SAMPLE_BUFFERS_SGIS] = "SAMPLE_BUFFERS_SGIS"
	enumLookup[POINT_FADE_THRESHOLD_SIZE_ARB] = "POINT_FADE_THRESHOLD_SIZE_ARB"
	enumLookup[DEPTH_RENDERABLE] = "DEPTH_RENDERABLE"
	enumLookup[ALIASED_POINT_SIZE_RANGE] = "ALIASED_POINT_SIZE_RANGE"
	enumLookup[MATRIX7_ARB] = "MATRIX7_ARB"
	enumLookup[MAX_VERTEX_BINDABLE_UNIFORMS_EXT] = "MAX_VERTEX_BINDABLE_UNIFORMS_EXT"
	enumLookup[NEAREST_MIPMAP_NEAREST] = "NEAREST_MIPMAP_NEAREST"
	enumLookup[INDEX_ARRAY_TYPE_EXT] = "INDEX_ARRAY_TYPE_EXT"
	enumLookup[DOUBLE_MAT2x3] = "DOUBLE_MAT2x3"
	enumLookup[HISTOGRAM_LUMINANCE_SIZE_EXT] = "HISTOGRAM_LUMINANCE_SIZE_EXT"
	enumLookup[GEOMETRY_TEXTURE] = "GEOMETRY_TEXTURE"
	enumLookup[TEXTURE_COMPRESSED] = "TEXTURE_COMPRESSED"
	enumLookup[VARIABLE_G_NV] = "VARIABLE_G_NV"
	enumLookup[CURRENT_BIT] = "CURRENT_BIT"
	enumLookup[TEXTURE_4D_SGIS] = "TEXTURE_4D_SGIS"
	enumLookup[RESTART_SUN] = "RESTART_SUN"
	enumLookup[PERTURB_EXT] = "PERTURB_EXT"
	enumLookup[TEXTURE_2D_ARRAY] = "TEXTURE_2D_ARRAY"
	enumLookup[DOUBLE_MAT2x3_EXT] = "DOUBLE_MAT2x3_EXT"
	enumLookup[CURRENT_INDEX] = "CURRENT_INDEX"
	enumLookup[TEXTURE_COMPARE_OPERATOR_SGIX] = "TEXTURE_COMPARE_OPERATOR_SGIX"
	enumLookup[MAX_COMPUTE_UNIFORM_COMPONENTS] = "MAX_COMPUTE_UNIFORM_COMPONENTS"
	enumLookup[COVERAGE_COMPONENT4_NV] = "COVERAGE_COMPONENT4_NV"
	enumLookup[PATH_FORMAT_PS_NV] = "PATH_FORMAT_PS_NV"
	enumLookup[MATRIX1_NV] = "MATRIX1_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT] = "FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT"
	enumLookup[UNPACK_IMAGE_HEIGHT] = "UNPACK_IMAGE_HEIGHT"
	enumLookup[PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT] = "PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT"
	enumLookup[OUTPUT_TEXTURE_COORD2_EXT] = "OUTPUT_TEXTURE_COORD2_EXT"
	enumLookup[FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB] = "FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[NUM_LOOPBACK_COMPONENTS_ATI] = "NUM_LOOPBACK_COMPONENTS_ATI"
	enumLookup[PACK_COMPRESSED_BLOCK_HEIGHT] = "PACK_COMPRESSED_BLOCK_HEIGHT"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR"
	enumLookup[BLEND_SRC_RGB_OES] = "BLEND_SRC_RGB_OES"
	enumLookup[DOUBLE_MAT4x2] = "DOUBLE_MAT4x2"
	enumLookup[YCBAYCR8A_4224_NV] = "YCBAYCR8A_4224_NV"
	enumLookup[TEXTURE24_ARB] = "TEXTURE24_ARB"
	enumLookup[INVERSE_TRANSPOSE_NV] = "INVERSE_TRANSPOSE_NV"
	enumLookup[STENCIL_TAG_BITS_EXT] = "STENCIL_TAG_BITS_EXT"
	enumLookup[MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV] = "MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV"
	enumLookup[PIXEL_FRAGMENT_RGB_SOURCE_SGIS] = "PIXEL_FRAGMENT_RGB_SOURCE_SGIS"
	enumLookup[MATRIX2_NV] = "MATRIX2_NV"
	enumLookup[NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI] = "NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI"
	enumLookup[STENCIL_INDEX1] = "STENCIL_INDEX1"
	enumLookup[RGB16_SNORM] = "RGB16_SNORM"
	enumLookup[INT64_VEC3_NV] = "INT64_VEC3_NV"
	enumLookup[MAT_SPECULAR_BIT_PGI] = "MAT_SPECULAR_BIT_PGI"
	enumLookup[VERSION_1_4] = "VERSION_1_4"
	enumLookup[SGI_color_matrix] = "SGI_color_matrix"
	enumLookup[STENCIL_VALUE_MASK] = "STENCIL_VALUE_MASK"
	enumLookup[SCISSOR_BOX] = "SCISSOR_BOX"
	enumLookup[UNPACK_SKIP_IMAGES_EXT] = "UNPACK_SKIP_IMAGES_EXT"
	enumLookup[MAX_ACTIVE_LIGHTS_SGIX] = "MAX_ACTIVE_LIGHTS_SGIX"
	enumLookup[PROXY_TEXTURE_CUBE_MAP] = "PROXY_TEXTURE_CUBE_MAP"
	enumLookup[MAX_DRAW_BUFFERS] = "MAX_DRAW_BUFFERS"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR"
	enumLookup[PACK_IMAGE_DEPTH_SGIS] = "PACK_IMAGE_DEPTH_SGIS"
	enumLookup[MAX_VERTEX_VARYING_COMPONENTS_EXT] = "MAX_VERTEX_VARYING_COMPONENTS_EXT"
	enumLookup[SAMPLE_ALPHA_TO_COVERAGE_ARB] = "SAMPLE_ALPHA_TO_COVERAGE_ARB"
	enumLookup[MAP2_VERTEX_ATTRIB5_4_NV] = "MAP2_VERTEX_ATTRIB5_4_NV"
	enumLookup[INTENSITY_SNORM] = "INTENSITY_SNORM"
	enumLookup[ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER] = "ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER"
	enumLookup[TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE] = "TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE"
	enumLookup[LINE_BIT] = "LINE_BIT"
	enumLookup[COLOR_ATTACHMENT4_NV] = "COLOR_ATTACHMENT4_NV"
	enumLookup[TRANSPOSE_AFFINE_3D_NV] = "TRANSPOSE_AFFINE_3D_NV"
	enumLookup[COMPRESSED_R11_EAC] = "COMPRESSED_R11_EAC"
	enumLookup[SOURCE1_RGB_ARB] = "SOURCE1_RGB_ARB"
	enumLookup[DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV] = "DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB] = "VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB"
	enumLookup[NUM_INSTRUCTIONS_PER_PASS_ATI] = "NUM_INSTRUCTIONS_PER_PASS_ATI"
	enumLookup[RENDERBUFFER_ALPHA_SIZE_EXT] = "RENDERBUFFER_ALPHA_SIZE_EXT"
	enumLookup[DUAL_TEXTURE_SELECT_SGIS] = "DUAL_TEXTURE_SELECT_SGIS"
	enumLookup[MODELVIEW21_ARB] = "MODELVIEW21_ARB"
	enumLookup[HI_BIAS_NV] = "HI_BIAS_NV"
	enumLookup[REG_11_ATI] = "REG_11_ATI"
	enumLookup[RGB565] = "RGB565"
	enumLookup[SAMPLE_ALPHA_TO_MASK_SGIS] = "SAMPLE_ALPHA_TO_MASK_SGIS"
	enumLookup[QUAD_ALPHA8_SGIS] = "QUAD_ALPHA8_SGIS"
	enumLookup[TEXTURE_MULTI_BUFFER_HINT_SGIX] = "TEXTURE_MULTI_BUFFER_HINT_SGIX"
	enumLookup[PREVIOUS_EXT] = "PREVIOUS_EXT"
	enumLookup[BLEND_EQUATION_ALPHA_EXT] = "BLEND_EQUATION_ALPHA_EXT"
	enumLookup[STENCIL_BACK_REF] = "STENCIL_BACK_REF"
	enumLookup[FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT] = "FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT"
	enumLookup[ISOLINES] = "ISOLINES"
	enumLookup[DECR] = "DECR"
	enumLookup[SCALEBIAS_HINT_SGIX] = "SCALEBIAS_HINT_SGIX"
	enumLookup[MAP2_VERTEX_ATTRIB1_4_NV] = "MAP2_VERTEX_ATTRIB1_4_NV"
	enumLookup[INTERLACE_READ_OML] = "INTERLACE_READ_OML"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES"
	enumLookup[FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT] = "FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT"
	enumLookup[PROXY_TEXTURE_CUBE_MAP_ARRAY] = "PROXY_TEXTURE_CUBE_MAP_ARRAY"
	enumLookup[TRANSLATE_3D_NV] = "TRANSLATE_3D_NV"
	enumLookup[EXT_blend_subtract] = "EXT_blend_subtract"
	enumLookup[EMBOSS_MAP_NV] = "EMBOSS_MAP_NV"
	enumLookup[DSDT8_NV] = "DSDT8_NV"
	enumLookup[BUFFER_SIZE] = "BUFFER_SIZE"
	enumLookup[BUFFER_MAP_POINTER_OES] = "BUFFER_MAP_POINTER_OES"
	enumLookup[LIGHTING] = "LIGHTING"
	enumLookup[RGBA2_EXT] = "RGBA2_EXT"
	enumLookup[DISCARD_NV] = "DISCARD_NV"
	enumLookup[FLOAT_VEC4] = "FLOAT_VEC4"
	enumLookup[INVALID_FRAMEBUFFER_OPERATION_EXT] = "INVALID_FRAMEBUFFER_OPERATION_EXT"
	enumLookup[LIGHT_MODEL_TWO_SIDE] = "LIGHT_MODEL_TWO_SIDE"
	enumLookup[LIGHT_ENV_MODE_SGIX] = "LIGHT_ENV_MODE_SGIX"
	enumLookup[UNSIGNED_SHORT_1_15_REV_MESA] = "UNSIGNED_SHORT_1_15_REV_MESA"
	enumLookup[OUTPUT_TEXTURE_COORD13_EXT] = "OUTPUT_TEXTURE_COORD13_EXT"
	enumLookup[RELATIVE_MOVE_TO_NV] = "RELATIVE_MOVE_TO_NV"
	enumLookup[BLEND_DST_RGB] = "BLEND_DST_RGB"
	enumLookup[DEBUG_SOURCE_THIRD_PARTY_ARB] = "DEBUG_SOURCE_THIRD_PARTY_ARB"
	enumLookup[FRONT_LEFT] = "FRONT_LEFT"
	enumLookup[REPLICATE_BORDER_HP] = "REPLICATE_BORDER_HP"
	enumLookup[UNIFORM_BARRIER_BIT_EXT] = "UNIFORM_BARRIER_BIT_EXT"
	enumLookup[PROXY_TEXTURE_3D] = "PROXY_TEXTURE_3D"
	enumLookup[MAX_SPOT_EXPONENT_NV] = "MAX_SPOT_EXPONENT_NV"
	enumLookup[PROGRAM_NATIVE_PARAMETERS_ARB] = "PROGRAM_NATIVE_PARAMETERS_ARB"
	enumLookup[TIMESTAMP] = "TIMESTAMP"
	enumLookup[CONSTANT_COLOR_EXT] = "CONSTANT_COLOR_EXT"
	enumLookup[COMBINE] = "COMBINE"
	enumLookup[TRANSPOSE_CURRENT_MATRIX_ARB] = "TRANSPOSE_CURRENT_MATRIX_ARB"
	enumLookup[FLOAT_VEC4_ARB] = "FLOAT_VEC4_ARB"
	enumLookup[SHADER_SOURCE_LENGTH] = "SHADER_SOURCE_LENGTH"
	enumLookup[BLEND_DST_ALPHA] = "BLEND_DST_ALPHA"
	enumLookup[PROXY_TEXTURE_CUBE_MAP_EXT] = "PROXY_TEXTURE_CUBE_MAP_EXT"
	enumLookup[COLOR4_BIT_PGI] = "COLOR4_BIT_PGI"
	enumLookup[NORMAL_ARRAY_TYPE] = "NORMAL_ARRAY_TYPE"
	enumLookup[FILTER4_SGIS] = "FILTER4_SGIS"
	enumLookup[PIXEL_PACK_BUFFER_ARB] = "PIXEL_PACK_BUFFER_ARB"
	enumLookup[DEBUG_LOGGED_MESSAGES_ARB] = "DEBUG_LOGGED_MESSAGES_ARB"
	enumLookup[PASS_THROUGH_TOKEN] = "PASS_THROUGH_TOKEN"
	enumLookup[QUERY_COUNTER_BITS_ARB] = "QUERY_COUNTER_BITS_ARB"
	enumLookup[FIXED_ONLY_ARB] = "FIXED_ONLY_ARB"
	enumLookup[QUERY_BUFFER_AMD] = "QUERY_BUFFER_AMD"
	enumLookup[DETAIL_TEXTURE_2D_BINDING_SGIS] = "DETAIL_TEXTURE_2D_BINDING_SGIS"
	enumLookup[SOURCE3_RGB_NV] = "SOURCE3_RGB_NV"
	enumLookup[MAP2_VERTEX_ATTRIB8_4_NV] = "MAP2_VERTEX_ATTRIB8_4_NV"
	enumLookup[MAX_PROGRAM_SUBROUTINE_NUM_NV] = "MAX_PROGRAM_SUBROUTINE_NUM_NV"
	enumLookup[IMAGE_3D] = "IMAGE_3D"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR"
	enumLookup[MULTISAMPLE_BIT_ARB] = "MULTISAMPLE_BIT_ARB"
	enumLookup[FEEDBACK_BUFFER_POINTER] = "FEEDBACK_BUFFER_POINTER"
	enumLookup[MAX_SHININESS_NV] = "MAX_SHININESS_NV"
	enumLookup[MODELVIEW18_ARB] = "MODELVIEW18_ARB"
	enumLookup[SHADER_TYPE] = "SHADER_TYPE"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_MODE_EXT] = "TRANSFORM_FEEDBACK_BUFFER_MODE_EXT"
	enumLookup[COLOR_ATTACHMENT13_NV] = "COLOR_ATTACHMENT13_NV"
	enumLookup[TEXTURE_FETCH_BARRIER_BIT] = "TEXTURE_FETCH_BARRIER_BIT"
	enumLookup[ALPHA] = "ALPHA"
	enumLookup[COMBINE_RGB_ARB] = "COMBINE_RGB_ARB"
	enumLookup[MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT] = "MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT"
	enumLookup[DEPENDENT_HILO_TEXTURE_2D_NV] = "DEPENDENT_HILO_TEXTURE_2D_NV"
	enumLookup[REG_0_ATI] = "REG_0_ATI"
	enumLookup[MAX_FRAGMENT_UNIFORM_COMPONENTS] = "MAX_FRAGMENT_UNIFORM_COMPONENTS"
	enumLookup[RGBA16I] = "RGBA16I"
	enumLookup[RGB_SNORM] = "RGB_SNORM"
	enumLookup[TESSELLATION_MODE_AMD] = "TESSELLATION_MODE_AMD"
	enumLookup[DRAW_BUFFER0] = "DRAW_BUFFER0"
	enumLookup[VERSION_3_0] = "VERSION_3_0"
	enumLookup[TRACE_OPERATIONS_BIT_MESA] = "TRACE_OPERATIONS_BIT_MESA"
	enumLookup[TEXTURE_INTERNAL_FORMAT_QCOM] = "TEXTURE_INTERNAL_FORMAT_QCOM"
	enumLookup[ETC1_RGB8_OES] = "ETC1_RGB8_OES"
	enumLookup[BINNING_CONTROL_HINT_QCOM] = "BINNING_CONTROL_HINT_QCOM"
	enumLookup[PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY] = "PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY"
	enumLookup[SHADER_STORAGE_BLOCK] = "SHADER_STORAGE_BLOCK"
	enumLookup[MAX_NUM_COMPATIBLE_SUBROUTINES] = "MAX_NUM_COMPATIBLE_SUBROUTINES"
	enumLookup[PIXEL_MAP_A_TO_A] = "PIXEL_MAP_A_TO_A"
	enumLookup[CAVEAT_SUPPORT] = "CAVEAT_SUPPORT"
	enumLookup[MAX_PROGRAM_RESULT_COMPONENTS_NV] = "MAX_PROGRAM_RESULT_COMPONENTS_NV"
	enumLookup[R11F_G11F_B10F_EXT] = "R11F_G11F_B10F_EXT"
	enumLookup[SLUMINANCE_ALPHA_NV] = "SLUMINANCE_ALPHA_NV"
	enumLookup[RENDERBUFFER_SAMPLES] = "RENDERBUFFER_SAMPLES"
	enumLookup[MAX_VARYING_VECTORS] = "MAX_VARYING_VECTORS"
	enumLookup[ONE_MINUS_CONSTANT_ALPHA_EXT] = "ONE_MINUS_CONSTANT_ALPHA_EXT"
	enumLookup[INTERNALFORMAT_BLUE_TYPE] = "INTERNALFORMAT_BLUE_TYPE"
	enumLookup[COLOR_ATTACHMENT15_EXT] = "COLOR_ATTACHMENT15_EXT"
	enumLookup[UTF8_NV] = "UTF8_NV"
	enumLookup[SGIS_fog_function] = "SGIS_fog_function"
	enumLookup[COLOR_TABLE_ALPHA_SIZE] = "COLOR_TABLE_ALPHA_SIZE"
	enumLookup[MATRIX_INDEX_ARRAY_OES] = "MATRIX_INDEX_ARRAY_OES"
	enumLookup[STENCIL_WRITEMASK] = "STENCIL_WRITEMASK"
	enumLookup[TEXTURE5] = "TEXTURE5"
	enumLookup[MAP2_VERTEX_ATTRIB4_4_NV] = "MAP2_VERTEX_ATTRIB4_4_NV"
	enumLookup[GPU_ADDRESS_NV] = "GPU_ADDRESS_NV"
	enumLookup[POST_CONVOLUTION_RED_BIAS_EXT] = "POST_CONVOLUTION_RED_BIAS_EXT"
	enumLookup[DISTANCE_ATTENUATION_EXT] = "DISTANCE_ATTENUATION_EXT"
	enumLookup[SLIM12S_SGIX] = "SLIM12S_SGIX"
	enumLookup[SOURCE2_RGB] = "SOURCE2_RGB"
	enumLookup[COMPRESSED_RED_RGTC1] = "COMPRESSED_RED_RGTC1"
	enumLookup[DOUBLE_MAT2_EXT] = "DOUBLE_MAT2_EXT"
	enumLookup[BLEND_EQUATION_OES] = "BLEND_EQUATION_OES"
	enumLookup[SAMPLE_COVERAGE_VALUE] = "SAMPLE_COVERAGE_VALUE"
	enumLookup[BOOL] = "BOOL"
	enumLookup[STENCIL_ATTACHMENT] = "STENCIL_ATTACHMENT"
	enumLookup[MIN_LOD_WARNING_AMD] = "MIN_LOD_WARNING_AMD"
	enumLookup[UNSIGNED_SHORT_8_8_APPLE] = "UNSIGNED_SHORT_8_8_APPLE"
	enumLookup[COMPRESSED_RED_GREEN_RGTC2_EXT] = "COMPRESSED_RED_GREEN_RGTC2_EXT"
	enumLookup[VERTEX_ARRAY_OBJECT_AMD] = "VERTEX_ARRAY_OBJECT_AMD"
	enumLookup[COLOR_TABLE_FORMAT_SGI] = "COLOR_TABLE_FORMAT_SGI"
	enumLookup[REPLACE_OLDEST_SUN] = "REPLACE_OLDEST_SUN"
	enumLookup[PREVIOUS] = "PREVIOUS"
	enumLookup[REG_4_ATI] = "REG_4_ATI"
	enumLookup[COMPRESSED_SIGNED_RG_RGTC2] = "COMPRESSED_SIGNED_RG_RGTC2"
	enumLookup[BUFFER_GPU_ADDRESS_NV] = "BUFFER_GPU_ADDRESS_NV"
	enumLookup[LUMINANCE4_ALPHA4_EXT] = "LUMINANCE4_ALPHA4_EXT"
	enumLookup[SIGNED_LUMINANCE8_ALPHA8_NV] = "SIGNED_LUMINANCE8_ALPHA8_NV"
	enumLookup[UNIFORM_OFFSET] = "UNIFORM_OFFSET"
	enumLookup[GLYPH_WIDTH_BIT_NV] = "GLYPH_WIDTH_BIT_NV"
	enumLookup[MATRIX_STRIDE] = "MATRIX_STRIDE"
	enumLookup[NORMAL_ARRAY] = "NORMAL_ARRAY"
	enumLookup[MAX_ASYNC_READ_PIXELS_SGIX] = "MAX_ASYNC_READ_PIXELS_SGIX"
	enumLookup[PROGRAM_RESIDENT_NV] = "PROGRAM_RESIDENT_NV"
	enumLookup[INDEX_ARRAY_LIST_IBM] = "INDEX_ARRAY_LIST_IBM"
	enumLookup[INDEX_SHIFT] = "INDEX_SHIFT"
	enumLookup[ALPHA_BITS] = "ALPHA_BITS"
	enumLookup[MAX_ELEMENTS_INDICES_EXT] = "MAX_ELEMENTS_INDICES_EXT"
	enumLookup[R1UI_V3F_SUN] = "R1UI_V3F_SUN"
	enumLookup[DEPTH_CLEAR_VALUE] = "DEPTH_CLEAR_VALUE"
	enumLookup[CONVOLUTION_BORDER_MODE_EXT] = "CONVOLUTION_BORDER_MODE_EXT"
	enumLookup[DEPTH_STENCIL_ATTACHMENT] = "DEPTH_STENCIL_ATTACHMENT"
	enumLookup[TEXTURE_2D_STACK_MESAX] = "TEXTURE_2D_STACK_MESAX"
	enumLookup[COLOR_ATTACHMENT0_NV] = "COLOR_ATTACHMENT0_NV"
	enumLookup[COPY_PIXEL_TOKEN] = "COPY_PIXEL_TOKEN"
	enumLookup[COLOR_ARRAY_EXT] = "COLOR_ARRAY_EXT"
	enumLookup[OCCLUSION_TEST_HP] = "OCCLUSION_TEST_HP"
	enumLookup[MAP2_GRID_SEGMENTS] = "MAP2_GRID_SEGMENTS"
	enumLookup[DEBUG_SOURCE_THIRD_PARTY] = "DEBUG_SOURCE_THIRD_PARTY"
	enumLookup[STENCIL_BACK_WRITEMASK] = "STENCIL_BACK_WRITEMASK"
	enumLookup[GEOMETRY_OUTPUT_TYPE_EXT] = "GEOMETRY_OUTPUT_TYPE_EXT"
	enumLookup[PHONG_HINT_WIN] = "PHONG_HINT_WIN"
	enumLookup[BUFFER_OBJECT_APPLE] = "BUFFER_OBJECT_APPLE"
	enumLookup[SHADING_LANGUAGE_VERSION_ARB] = "SHADING_LANGUAGE_VERSION_ARB"
	enumLookup[UNSIGNED_INT64_AMD] = "UNSIGNED_INT64_AMD"
	enumLookup[TEXTURE_CLIPMAP_OFFSET_SGIX] = "TEXTURE_CLIPMAP_OFFSET_SGIX"
	enumLookup[COMPRESSED_RGBA] = "COMPRESSED_RGBA"
	enumLookup[RGB_SCALE_EXT] = "RGB_SCALE_EXT"
	enumLookup[MAGNITUDE_SCALE_NV] = "MAGNITUDE_SCALE_NV"
	enumLookup[TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV] = "TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV"
	enumLookup[MULT] = "MULT"
	enumLookup[MODELVIEW0_ARB] = "MODELVIEW0_ARB"
	enumLookup[INT_2_10_10_10_REV] = "INT_2_10_10_10_REV"
	enumLookup[UNSIGNED_INT_ATOMIC_COUNTER] = "UNSIGNED_INT_ATOMIC_COUNTER"
	enumLookup[VENDOR] = "VENDOR"
	enumLookup[POST_COLOR_MATRIX_RED_SCALE] = "POST_COLOR_MATRIX_RED_SCALE"
	enumLookup[BOOL_VEC2_ARB] = "BOOL_VEC2_ARB"
	enumLookup[DRAW_PIXEL_TOKEN] = "DRAW_PIXEL_TOKEN"
	enumLookup[COLOR_TABLE_INTENSITY_SIZE] = "COLOR_TABLE_INTENSITY_SIZE"
	enumLookup[CURRENT_RASTER_SECONDARY_COLOR] = "CURRENT_RASTER_SECONDARY_COLOR"
	enumLookup[NORMAL_MAP_NV] = "NORMAL_MAP_NV"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_X_OES] = "TEXTURE_CUBE_MAP_POSITIVE_X_OES"
	enumLookup[SOURCE2_RGB_ARB] = "SOURCE2_RGB_ARB"
	enumLookup[OPERAND0_ALPHA_EXT] = "OPERAND0_ALPHA_EXT"
	enumLookup[RGBA_FLOAT_MODE_ATI] = "RGBA_FLOAT_MODE_ATI"
	enumLookup[CON_30_ATI] = "CON_30_ATI"
	enumLookup[VIDEO_CAPTURE_FRAME_HEIGHT_NV] = "VIDEO_CAPTURE_FRAME_HEIGHT_NV"
	enumLookup[STRICT_DEPTHFUNC_HINT_PGI] = "STRICT_DEPTHFUNC_HINT_PGI"
	enumLookup[COLOR_ARRAY_TYPE] = "COLOR_ARRAY_TYPE"
	enumLookup[MAX_DRAW_BUFFERS_ARB] = "MAX_DRAW_BUFFERS_ARB"
	enumLookup[MAX_VERTEX_TEXTURE_IMAGE_UNITS] = "MAX_VERTEX_TEXTURE_IMAGE_UNITS"
	enumLookup[MALI_SHADER_BINARY_ARM] = "MALI_SHADER_BINARY_ARM"
	enumLookup[PIXEL_MAP_S_TO_S] = "PIXEL_MAP_S_TO_S"
	enumLookup[ALL_COMPLETED_NV] = "ALL_COMPLETED_NV"
	enumLookup[MAP1_VERTEX_ATTRIB7_4_NV] = "MAP1_VERTEX_ATTRIB7_4_NV"
	enumLookup[PN_TRIANGLES_POINT_MODE_LINEAR_ATI] = "PN_TRIANGLES_POINT_MODE_LINEAR_ATI"
	enumLookup[INT_SAMPLER_1D_ARRAY_EXT] = "INT_SAMPLER_1D_ARRAY_EXT"
	enumLookup[COLOR_MATERIAL_FACE] = "COLOR_MATERIAL_FACE"
	enumLookup[COLOR_TABLE] = "COLOR_TABLE"
	enumLookup[TEXTURE16_ARB] = "TEXTURE16_ARB"
	enumLookup[FLOAT_RG32_NV] = "FLOAT_RG32_NV"
	enumLookup[PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES] = "PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES"
	enumLookup[COMPRESSED_RG_RGTC2] = "COMPRESSED_RG_RGTC2"
	enumLookup[RED_BIAS] = "RED_BIAS"
	enumLookup[MAX_PROJECTION_STACK_DEPTH] = "MAX_PROJECTION_STACK_DEPTH"
	enumLookup[DUAL_ALPHA4_SGIS] = "DUAL_ALPHA4_SGIS"
	enumLookup[TEXTURE3] = "TEXTURE3"
	enumLookup[REFLECTION_MAP_EXT] = "REFLECTION_MAP_EXT"
	enumLookup[PROGRAM_FORMAT_ASCII_ARB] = "PROGRAM_FORMAT_ASCII_ARB"
	enumLookup[LINEAR_MIPMAP_LINEAR] = "LINEAR_MIPMAP_LINEAR"
	enumLookup[INTERNALFORMAT_ALPHA_SIZE] = "INTERNALFORMAT_ALPHA_SIZE"
	enumLookup[ATOMIC_COUNTER_BUFFER_SIZE] = "ATOMIC_COUNTER_BUFFER_SIZE"
	enumLookup[EQUIV] = "EQUIV"
	enumLookup[DRAW_BUFFER1_ATI] = "DRAW_BUFFER1_ATI"
	enumLookup[CLIP_PLANE1] = "CLIP_PLANE1"
	enumLookup[REDUCE_EXT] = "REDUCE_EXT"
	enumLookup[LUMINANCE16_ALPHA16_EXT] = "LUMINANCE16_ALPHA16_EXT"
	enumLookup[PROGRAM_PIPELINE] = "PROGRAM_PIPELINE"
	enumLookup[MODELVIEW1_EXT] = "MODELVIEW1_EXT"
	enumLookup[POST_CONVOLUTION_GREEN_SCALE_EXT] = "POST_CONVOLUTION_GREEN_SCALE_EXT"
	enumLookup[TEXTURE_GREEN_SIZE_EXT] = "TEXTURE_GREEN_SIZE_EXT"
	enumLookup[BLEND_SRC_ALPHA_OES] = "BLEND_SRC_ALPHA_OES"
	enumLookup[LIGHT_MODEL_COLOR_CONTROL] = "LIGHT_MODEL_COLOR_CONTROL"
	enumLookup[CND0_ATI] = "CND0_ATI"
	enumLookup[RED_INTEGER] = "RED_INTEGER"
	enumLookup[DOUBLE] = "DOUBLE"
	enumLookup[HISTOGRAM_GREEN_SIZE_EXT] = "HISTOGRAM_GREEN_SIZE_EXT"
	enumLookup[COMPRESSED_RGBA_S3TC_DXT5_EXT] = "COMPRESSED_RGBA_S3TC_DXT5_EXT"
	enumLookup[OBJECT_DELETE_STATUS_ARB] = "OBJECT_DELETE_STATUS_ARB"
	enumLookup[TRANSFORM_FEEDBACK_VARYING] = "TRANSFORM_FEEDBACK_VARYING"
	enumLookup[LINES_ADJACENCY_ARB] = "LINES_ADJACENCY_ARB"
	enumLookup[STACK_UNDERFLOW] = "STACK_UNDERFLOW"
	enumLookup[DETAIL_TEXTURE_MODE_SGIS] = "DETAIL_TEXTURE_MODE_SGIS"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_PAUSED] = "TRANSFORM_FEEDBACK_BUFFER_PAUSED"
	enumLookup[SECONDARY_COLOR_ARRAY_ADDRESS_NV] = "SECONDARY_COLOR_ARRAY_ADDRESS_NV"
	enumLookup[DEBUG_LOGGED_MESSAGES] = "DEBUG_LOGGED_MESSAGES"
	enumLookup[SAMPLE_ALPHA_TO_COVERAGE] = "SAMPLE_ALPHA_TO_COVERAGE"
	enumLookup[STATIC_VERTEX_ARRAY_IBM] = "STATIC_VERTEX_ARRAY_IBM"
	enumLookup[CONVOLUTION_BORDER_COLOR] = "CONVOLUTION_BORDER_COLOR"
	enumLookup[REPLACEMENT_CODE_SUN] = "REPLACEMENT_CODE_SUN"
	enumLookup[PROGRAM_INSTRUCTIONS_ARB] = "PROGRAM_INSTRUCTIONS_ARB"
	enumLookup[AUX2] = "AUX2"
	enumLookup[SOURCE1_ALPHA_ARB] = "SOURCE1_ALPHA_ARB"
	enumLookup[MAP2_VERTEX_ATTRIB6_4_NV] = "MAP2_VERTEX_ATTRIB6_4_NV"
	enumLookup[SAMPLER_2D_ARRAY_SHADOW] = "SAMPLER_2D_ARRAY_SHADOW"
	enumLookup[GEOMETRY_INPUT_TYPE_ARB] = "GEOMETRY_INPUT_TYPE_ARB"
	enumLookup[SHADER_BINARY_VIV] = "SHADER_BINARY_VIV"
	enumLookup[SHADER_IMAGE_LOAD] = "SHADER_IMAGE_LOAD"
	enumLookup[VERTEX_BINDING_OFFSET] = "VERTEX_BINDING_OFFSET"
	enumLookup[OP_MULTIPLY_MATRIX_EXT] = "OP_MULTIPLY_MATRIX_EXT"
	enumLookup[WRITE_ONLY_ARB] = "WRITE_ONLY_ARB"
	enumLookup[BOLD_BIT_NV] = "BOLD_BIT_NV"
	enumLookup[GLOBAL_ALPHA_SUN] = "GLOBAL_ALPHA_SUN"
	enumLookup[UNSIGNED_INT_VEC2] = "UNSIGNED_INT_VEC2"
	enumLookup[INT_SAMPLER_CUBE] = "INT_SAMPLER_CUBE"
	enumLookup[TEXTURE_3D] = "TEXTURE_3D"
	enumLookup[COMPARE_REF_DEPTH_TO_TEXTURE_EXT] = "COMPARE_REF_DEPTH_TO_TEXTURE_EXT"
	enumLookup[READ_PIXEL_DATA_RANGE_POINTER_NV] = "READ_PIXEL_DATA_RANGE_POINTER_NV"
	enumLookup[RENDERBUFFER_RED_SIZE_EXT] = "RENDERBUFFER_RED_SIZE_EXT"
	enumLookup[FRAMEBUFFER_SRGB_CAPABLE_EXT] = "FRAMEBUFFER_SRGB_CAPABLE_EXT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER"
	enumLookup[POINT_SIZE_MAX_ARB] = "POINT_SIZE_MAX_ARB"
	enumLookup[TEXTURE31_ARB] = "TEXTURE31_ARB"
	enumLookup[MAX_TESS_GEN_LEVEL] = "MAX_TESS_GEN_LEVEL"
	enumLookup[VIEW_CLASS_BPTC_FLOAT] = "VIEW_CLASS_BPTC_FLOAT"
	enumLookup[IMAGE_BINDING_NAME_EXT] = "IMAGE_BINDING_NAME_EXT"
	enumLookup[TEXTURE_ENV] = "TEXTURE_ENV"
	enumLookup[SAMPLES] = "SAMPLES"
	enumLookup[OUTPUT_TEXTURE_COORD27_EXT] = "OUTPUT_TEXTURE_COORD27_EXT"
	enumLookup[TEXTURE_GREEN_TYPE_ARB] = "TEXTURE_GREEN_TYPE_ARB"
	enumLookup[SAMPLES_SGIS] = "SAMPLES_SGIS"
	enumLookup[GLOBAL_ALPHA_FACTOR_SUN] = "GLOBAL_ALPHA_FACTOR_SUN"
	enumLookup[INTERNALFORMAT_ALPHA_TYPE] = "INTERNALFORMAT_ALPHA_TYPE"
	enumLookup[ALPHA_TEST_REF] = "ALPHA_TEST_REF"
	enumLookup[MAP2_VERTEX_ATTRIB3_4_NV] = "MAP2_VERTEX_ATTRIB3_4_NV"
	enumLookup[MODELVIEW1_ARB] = "MODELVIEW1_ARB"
	enumLookup[MODELVIEW27_ARB] = "MODELVIEW27_ARB"
	enumLookup[VERTEX_ATTRIB_MAP2_DOMAIN_APPLE] = "VERTEX_ATTRIB_MAP2_DOMAIN_APPLE"
	enumLookup[BOOL_VEC2] = "BOOL_VEC2"
	enumLookup[IMAGE_FORMAT_COMPATIBILITY_TYPE] = "IMAGE_FORMAT_COMPATIBILITY_TYPE"
	enumLookup[RIGHT] = "RIGHT"
	enumLookup[PACK_ALIGNMENT] = "PACK_ALIGNMENT"
	enumLookup[VERTEX_ARRAY] = "VERTEX_ARRAY"
	enumLookup[TEXTURE_MIN_FILTER] = "TEXTURE_MIN_FILTER"
	enumLookup[IMAGE_SCALE_X_HP] = "IMAGE_SCALE_X_HP"
	enumLookup[FRAMEBUFFER_DEFAULT] = "FRAMEBUFFER_DEFAULT"
	enumLookup[SLIM10U_SGIX] = "SLIM10U_SGIX"
	enumLookup[RGB_SCALE_ARB] = "RGB_SCALE_ARB"
	enumLookup[AUX_DEPTH_STENCIL_APPLE] = "AUX_DEPTH_STENCIL_APPLE"
	enumLookup[MEDIUM_FLOAT] = "MEDIUM_FLOAT"
	enumLookup[POST_COLOR_MATRIX_BLUE_SCALE] = "POST_COLOR_MATRIX_BLUE_SCALE"
	enumLookup[MAX_COMPUTE_ATOMIC_COUNTERS] = "MAX_COMPUTE_ATOMIC_COUNTERS"
	enumLookup[OP_MUL_EXT] = "OP_MUL_EXT"
	enumLookup[LUMINANCE] = "LUMINANCE"
	enumLookup[FRAGMENT_COLOR_EXT] = "FRAGMENT_COLOR_EXT"
	enumLookup[VERTEX_ATTRIB_ARRAY8_NV] = "VERTEX_ATTRIB_ARRAY8_NV"
	enumLookup[TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY] = "TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY"
	enumLookup[VERTEX_ARRAY_TYPE_EXT] = "VERTEX_ARRAY_TYPE_EXT"
	enumLookup[PIXEL_GROUP_COLOR_SGIS] = "PIXEL_GROUP_COLOR_SGIS"
	enumLookup[MATRIX30_ARB] = "MATRIX30_ARB"
	enumLookup[ACCUM] = "ACCUM"
	enumLookup[MATRIX1_ARB] = "MATRIX1_ARB"
	enumLookup[TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT] = "TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT"
	enumLookup[KEEP] = "KEEP"
	enumLookup[SEPARABLE_2D] = "SEPARABLE_2D"
	enumLookup[TEXTURE_1D_STACK_MESAX] = "TEXTURE_1D_STACK_MESAX"
	enumLookup[INT_VEC4_ARB] = "INT_VEC4_ARB"
	enumLookup[LUMINANCE_SNORM] = "LUMINANCE_SNORM"
	enumLookup[UNSIGNED_BYTE_2_3_3_REV_EXT] = "UNSIGNED_BYTE_2_3_3_REV_EXT"
	enumLookup[TEXTURE_RESIDENT_EXT] = "TEXTURE_RESIDENT_EXT"
	enumLookup[RELEASED_APPLE] = "RELEASED_APPLE"
	enumLookup[MAX_SHADER_BUFFER_ADDRESS_NV] = "MAX_SHADER_BUFFER_ADDRESS_NV"
	enumLookup[WAIT_FAILED_APPLE] = "WAIT_FAILED_APPLE"
	enumLookup[COLOR_ARRAY_LIST_IBM] = "COLOR_ARRAY_LIST_IBM"
	enumLookup[MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV] = "MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV"
	enumLookup[IMAGE_BINDING_NAME] = "IMAGE_BINDING_NAME"
	enumLookup[SGIX_fog_offset] = "SGIX_fog_offset"
	enumLookup[DEPTH_FUNC] = "DEPTH_FUNC"
	enumLookup[STENCIL_FUNC] = "STENCIL_FUNC"
	enumLookup[MAX_EXT] = "MAX_EXT"
	enumLookup[CLIP_VOLUME_CLIPPING_HINT_EXT] = "CLIP_VOLUME_CLIPPING_HINT_EXT"
	enumLookup[RENDERBUFFER_BINDING_EXT] = "RENDERBUFFER_BINDING_EXT"
	enumLookup[SGIS_point_line_texgen] = "SGIS_point_line_texgen"
	enumLookup[GL_4_BYTES] = "GL_4_BYTES"
	enumLookup[FRAGMENT_DEPTH] = "FRAGMENT_DEPTH"
	enumLookup[MATRIX0_NV] = "MATRIX0_NV"
	enumLookup[SIGNED_HILO8_NV] = "SIGNED_HILO8_NV"
	enumLookup[STENCIL_CLEAR_TAG_VALUE_EXT] = "STENCIL_CLEAR_TAG_VALUE_EXT"
	enumLookup[RGB10_A2UI] = "RGB10_A2UI"
	enumLookup[GLYPH_HAS_KERNING_BIT_NV] = "GLYPH_HAS_KERNING_BIT_NV"
	enumLookup[DEPTH_COMPONENT32] = "DEPTH_COMPONENT32"
	enumLookup[DRAW_BUFFER4_ARB] = "DRAW_BUFFER4_ARB"
	enumLookup[OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV] = "OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV"
	enumLookup[VERTEX_ATTRIB_MAP1_ORDER_APPLE] = "VERTEX_ATTRIB_MAP1_ORDER_APPLE"
	enumLookup[COLOR_ATTACHMENT6_EXT] = "COLOR_ATTACHMENT6_EXT"
	enumLookup[PATCH_DEFAULT_INNER_LEVEL] = "PATCH_DEFAULT_INNER_LEVEL"
	enumLookup[CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT] = "CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT"
	enumLookup[SIGNED_ALPHA_NV] = "SIGNED_ALPHA_NV"
	enumLookup[PATH_STROKE_BOUNDING_BOX_NV] = "PATH_STROKE_BOUNDING_BOX_NV"
	enumLookup[PROXY_TEXTURE_1D] = "PROXY_TEXTURE_1D"
	enumLookup[MAX_TESS_EVALUATION_OUTPUT_COMPONENTS] = "MAX_TESS_EVALUATION_OUTPUT_COMPONENTS"
	enumLookup[MAX_DEBUG_LOGGED_MESSAGES_ARB] = "MAX_DEBUG_LOGGED_MESSAGES_ARB"
	enumLookup[NORMAL_ARRAY_STRIDE] = "NORMAL_ARRAY_STRIDE"
	enumLookup[RESAMPLE_DECIMATE_SGIX] = "RESAMPLE_DECIMATE_SGIX"
	enumLookup[TEXTURE12] = "TEXTURE12"
	enumLookup[OPERAND0_RGB_EXT] = "OPERAND0_RGB_EXT"
	enumLookup[RED_BIT_ATI] = "RED_BIT_ATI"
	enumLookup[BOOL_VEC3] = "BOOL_VEC3"
	enumLookup[COMPRESSED_SRGB_ALPHA] = "COMPRESSED_SRGB_ALPHA"
	enumLookup[POLYGON] = "POLYGON"
	enumLookup[LUMINANCE_FLOAT16_ATI] = "LUMINANCE_FLOAT16_ATI"
	enumLookup[CURRENT_PALETTE_MATRIX_ARB] = "CURRENT_PALETTE_MATRIX_ARB"
	enumLookup[TEXCOORD2_BIT_PGI] = "TEXCOORD2_BIT_PGI"
	enumLookup[CMYK_EXT] = "CMYK_EXT"
	enumLookup[VARIABLE_B_NV] = "VARIABLE_B_NV"
	enumLookup[CON_17_ATI] = "CON_17_ATI"
	enumLookup[SRGB8_NV] = "SRGB8_NV"
	enumLookup[RASTERIZER_DISCARD_EXT] = "RASTERIZER_DISCARD_EXT"
	enumLookup[MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS] = "MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS"
	enumLookup[PATH_FORMAT_SVG_NV] = "PATH_FORMAT_SVG_NV"
	enumLookup[SYNC_FLUSH_COMMANDS_BIT_APPLE] = "SYNC_FLUSH_COMMANDS_BIT_APPLE"
	enumLookup[VERTEX_DATA_HINT_PGI] = "VERTEX_DATA_HINT_PGI"
	enumLookup[V3F] = "V3F"
	enumLookup[REFERENCE_PLANE_SGIX] = "REFERENCE_PLANE_SGIX"
	enumLookup[EYE_PLANE_ABSOLUTE_NV] = "EYE_PLANE_ABSOLUTE_NV"
	enumLookup[SGIS_multisample] = "SGIS_multisample"
	enumLookup[OUTPUT_TEXTURE_COORD4_EXT] = "OUTPUT_TEXTURE_COORD4_EXT"
	enumLookup[NOTEQUAL] = "NOTEQUAL"
	enumLookup[MAX_GEOMETRY_UNIFORM_BLOCKS] = "MAX_GEOMETRY_UNIFORM_BLOCKS"
	enumLookup[SLUMINANCE8_ALPHA8] = "SLUMINANCE8_ALPHA8"
	enumLookup[TEXTURE_SWIZZLE_B_EXT] = "TEXTURE_SWIZZLE_B_EXT"
	enumLookup[PATH_DASH_ARRAY_COUNT_NV] = "PATH_DASH_ARRAY_COUNT_NV"
	enumLookup[TEXTURE_WIDTH] = "TEXTURE_WIDTH"
	enumLookup[TEXTURE_WRAP_T] = "TEXTURE_WRAP_T"
	enumLookup[TEXTURE5_ARB] = "TEXTURE5_ARB"
	enumLookup[LUMINANCE_ALPHA32I_EXT] = "LUMINANCE_ALPHA32I_EXT"
	enumLookup[STENCIL_EXT] = "STENCIL_EXT"
	enumLookup[SYNC_X11_FENCE_EXT] = "SYNC_X11_FENCE_EXT"
	enumLookup[ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS] = "ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS"
	enumLookup[TESS_EVALUATION_SUBROUTINE] = "TESS_EVALUATION_SUBROUTINE"
	enumLookup[FUNC_REVERSE_SUBTRACT_EXT] = "FUNC_REVERSE_SUBTRACT_EXT"
	enumLookup[CONVOLUTION_1D_EXT] = "CONVOLUTION_1D_EXT"
	enumLookup[PROGRAM_PIPELINE_BINDING_EXT] = "PROGRAM_PIPELINE_BINDING_EXT"
	enumLookup[SLUMINANCE_EXT] = "SLUMINANCE_EXT"
	enumLookup[MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS] = "MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS"
	enumLookup[VERTEX_PROGRAM_PARAMETER_BUFFER_NV] = "VERTEX_PROGRAM_PARAMETER_BUFFER_NV"
	enumLookup[INT_IMAGE_CUBE_EXT] = "INT_IMAGE_CUBE_EXT"
	enumLookup[VERTEX4_BIT_PGI] = "VERTEX4_BIT_PGI"
	enumLookup[SGIX_vertex_preclip] = "SGIX_vertex_preclip"
	enumLookup[FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE] = "FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE"
	enumLookup[FOG_COORDINATE_ARRAY_POINTER] = "FOG_COORDINATE_ARRAY_POINTER"
	enumLookup[UNSIGNED_INT_24_8_NV] = "UNSIGNED_INT_24_8_NV"
	enumLookup[COMBINER6_NV] = "COMBINER6_NV"
	enumLookup[FRAMEBUFFER] = "FRAMEBUFFER"
	enumLookup[MAX_SAMPLE_MASK_WORDS_NV] = "MAX_SAMPLE_MASK_WORDS_NV"
	enumLookup[PATH_JOIN_STYLE_NV] = "PATH_JOIN_STYLE_NV"
	enumLookup[CURRENT_NORMAL] = "CURRENT_NORMAL"
	enumLookup[CLIENT_ATTRIB_STACK_DEPTH] = "CLIENT_ATTRIB_STACK_DEPTH"
	enumLookup[DEPTH_SCALE] = "DEPTH_SCALE"
	enumLookup[PROXY_TEXTURE_CUBE_MAP_ARB] = "PROXY_TEXTURE_CUBE_MAP_ARB"
	enumLookup[SWIZZLE_STQ_DQ_ATI] = "SWIZZLE_STQ_DQ_ATI"
	enumLookup[HORIZONTAL_LINE_TO_NV] = "HORIZONTAL_LINE_TO_NV"
	enumLookup[REPLACEMENT_CODE_ARRAY_POINTER_SUN] = "REPLACEMENT_CODE_ARRAY_POINTER_SUN"
	enumLookup[QUERY_BY_REGION_NO_WAIT_NV] = "QUERY_BY_REGION_NO_WAIT_NV"
	enumLookup[FRAME_NV] = "FRAME_NV"
	enumLookup[UNSIGNED_BYTE] = "UNSIGNED_BYTE"
	enumLookup[CLAMP] = "CLAMP"
	enumLookup[MAX_VERTEX_ATTRIB_BINDINGS] = "MAX_VERTEX_ATTRIB_BINDINGS"
	enumLookup[GENERIC_ATTRIB_NV] = "GENERIC_ATTRIB_NV"
	enumLookup[MOVE_TO_CONTINUES_NV] = "MOVE_TO_CONTINUES_NV"
	enumLookup[FRAGMENT_SUBROUTINE_UNIFORM] = "FRAGMENT_SUBROUTINE_UNIFORM"
	enumLookup[SOURCE2_ALPHA_EXT] = "SOURCE2_ALPHA_EXT"
	enumLookup[PALETTE8_RGBA4_OES] = "PALETTE8_RGBA4_OES"
	enumLookup[OR] = "OR"
	enumLookup[REFLECTION_MAP_NV] = "REFLECTION_MAP_NV"
	enumLookup[WEIGHT_ARRAY_POINTER_ARB] = "WEIGHT_ARRAY_POINTER_ARB"
	enumLookup[PERFMON_GLOBAL_MODE_QCOM] = "PERFMON_GLOBAL_MODE_QCOM"
	enumLookup[BLEND_DST] = "BLEND_DST"
	enumLookup[MAX_CLIPMAP_DEPTH_SGIX] = "MAX_CLIPMAP_DEPTH_SGIX"
	enumLookup[MODELVIEW15_ARB] = "MODELVIEW15_ARB"
	enumLookup[SAMPLE_BUFFERS_3DFX] = "SAMPLE_BUFFERS_3DFX"
	enumLookup[SIGNED_RGBA8_NV] = "SIGNED_RGBA8_NV"
	enumLookup[IMAGE_3D_EXT] = "IMAGE_3D_EXT"
	enumLookup[C3F_V3F] = "C3F_V3F"
	enumLookup[CLIP_PLANE0] = "CLIP_PLANE0"
	enumLookup[SAMPLE_COVERAGE_VALUE_ARB] = "SAMPLE_COVERAGE_VALUE_ARB"
	enumLookup[R16F_EXT] = "R16F_EXT"
	enumLookup[MAX_LABEL_LENGTH] = "MAX_LABEL_LENGTH"
	enumLookup[SIGNED_ALPHA8_NV] = "SIGNED_ALPHA8_NV"
	enumLookup[ALPHA_FLOAT32_ATI] = "ALPHA_FLOAT32_ATI"
	enumLookup[ACCUM_ADJACENT_PAIRS_NV] = "ACCUM_ADJACENT_PAIRS_NV"
	enumLookup[POLYGON_OFFSET_UNITS] = "POLYGON_OFFSET_UNITS"
	enumLookup[LUMINANCE16_EXT] = "LUMINANCE16_EXT"
	enumLookup[CONVOLUTION_BORDER_COLOR_HP] = "CONVOLUTION_BORDER_COLOR_HP"
	enumLookup[CON_29_ATI] = "CON_29_ATI"
	enumLookup[QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION] = "QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION"
	enumLookup[MAP_READ_BIT] = "MAP_READ_BIT"
	enumLookup[STEREO] = "STEREO"
	enumLookup[IMAGE_ROTATE_ORIGIN_X_HP] = "IMAGE_ROTATE_ORIGIN_X_HP"
	enumLookup[ADD_SIGNED_ARB] = "ADD_SIGNED_ARB"
	enumLookup[NUM_INSTRUCTIONS_TOTAL_ATI] = "NUM_INSTRUCTIONS_TOTAL_ATI"
	enumLookup[FRAGMENT_PROGRAM_POSITION_MESA] = "FRAGMENT_PROGRAM_POSITION_MESA"
	enumLookup[COLOR_ATTACHMENT2_NV] = "COLOR_ATTACHMENT2_NV"
	enumLookup[LAST_VIDEO_CAPTURE_STATUS_NV] = "LAST_VIDEO_CAPTURE_STATUS_NV"
	enumLookup[VERTEX_ARRAY_LIST_IBM] = "VERTEX_ARRAY_LIST_IBM"
	enumLookup[TRANSFORM_FEEDBACK_BARRIER_BIT_EXT] = "TRANSFORM_FEEDBACK_BARRIER_BIT_EXT"
	enumLookup[MULTISAMPLE_SGIS] = "MULTISAMPLE_SGIS"
	enumLookup[MAX_VERTEX_SHADER_LOCALS_EXT] = "MAX_VERTEX_SHADER_LOCALS_EXT"
	enumLookup[W_EXT] = "W_EXT"
	enumLookup[COMPRESSED_RGBA_PVRTC_2BPPV1_IMG] = "COMPRESSED_RGBA_PVRTC_2BPPV1_IMG"
	enumLookup[INT_SAMPLER_2D_RECT] = "INT_SAMPLER_2D_RECT"
	enumLookup[AMBIENT] = "AMBIENT"
	enumLookup[INTERNALFORMAT_SUPPORTED] = "INTERNALFORMAT_SUPPORTED"
	enumLookup[VERTEX_ATTRIB_BINDING] = "VERTEX_ATTRIB_BINDING"
	enumLookup[DRAW_BUFFER3_ATI] = "DRAW_BUFFER3_ATI"
	enumLookup[TEXTURE_COLOR_TABLE_SGI] = "TEXTURE_COLOR_TABLE_SGI"
	enumLookup[RGBA32F_ARB] = "RGBA32F_ARB"
	enumLookup[OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV] = "OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV"
	enumLookup[RENDERBUFFER_DEPTH_SIZE] = "RENDERBUFFER_DEPTH_SIZE"
	enumLookup[PIXEL_MODE_BIT] = "PIXEL_MODE_BIT"
	enumLookup[PROXY_TEXTURE_2D_EXT] = "PROXY_TEXTURE_2D_EXT"
	enumLookup[DECR_WRAP] = "DECR_WRAP"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Y_OES] = "TEXTURE_CUBE_MAP_POSITIVE_Y_OES"
	enumLookup[HALF_BIAS_NORMAL_NV] = "HALF_BIAS_NORMAL_NV"
	enumLookup[PIXEL_UNPACK_BUFFER_BINDING_ARB] = "PIXEL_UNPACK_BUFFER_BINDING_ARB"
	enumLookup[COLOR] = "COLOR"
	enumLookup[COLOR_TABLE_GREEN_SIZE_SGI] = "COLOR_TABLE_GREEN_SIZE_SGI"
	enumLookup[BGRA_EXT] = "BGRA_EXT"
	enumLookup[GENERATE_MIPMAP] = "GENERATE_MIPMAP"
	enumLookup[OP_MIN_EXT] = "OP_MIN_EXT"
	enumLookup[MAP1_VERTEX_ATTRIB4_4_NV] = "MAP1_VERTEX_ATTRIB4_4_NV"
	enumLookup[EVAL_TRIANGULAR_2D_NV] = "EVAL_TRIANGULAR_2D_NV"
	enumLookup[ANY_SAMPLES_PASSED_EXT] = "ANY_SAMPLES_PASSED_EXT"
	enumLookup[SAMPLER_CUBE_SHADOW] = "SAMPLER_CUBE_SHADOW"
	enumLookup[MAX_TESS_EVALUATION_UNIFORM_BLOCKS] = "MAX_TESS_EVALUATION_UNIFORM_BLOCKS"
	enumLookup[ROUND_NV] = "ROUND_NV"
	enumLookup[UNPACK_COMPRESSED_BLOCK_HEIGHT] = "UNPACK_COMPRESSED_BLOCK_HEIGHT"
	enumLookup[CUBIC_HP] = "CUBIC_HP"
	enumLookup[MAX_VERTEX_UNITS_OES] = "MAX_VERTEX_UNITS_OES"
	enumLookup[STENCIL_BACK_PASS_DEPTH_PASS] = "STENCIL_BACK_PASS_DEPTH_PASS"
	enumLookup[ANY_SAMPLES_PASSED] = "ANY_SAMPLES_PASSED"
	enumLookup[COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT] = "COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT"
	enumLookup[PROXY_HISTOGRAM] = "PROXY_HISTOGRAM"
	enumLookup[POINT_SIZE_MAX] = "POINT_SIZE_MAX"
	enumLookup[DOT3_RGB_ARB] = "DOT3_RGB_ARB"
	enumLookup[MAX_MAP_TESSELLATION_NV] = "MAX_MAP_TESSELLATION_NV"
	enumLookup[BLEND_DST_RGB_OES] = "BLEND_DST_RGB_OES"
	enumLookup[FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX] = "FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX"
	enumLookup[ZERO_EXT] = "ZERO_EXT"
	enumLookup[RGB16F_EXT] = "RGB16F_EXT"
	enumLookup[POST_CONVOLUTION_COLOR_TABLE_SGI] = "POST_CONVOLUTION_COLOR_TABLE_SGI"
	enumLookup[MODELVIEW28_ARB] = "MODELVIEW28_ARB"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_START] = "TRANSFORM_FEEDBACK_BUFFER_START"
	enumLookup[INTERLEAVED_ATTRIBS_NV] = "INTERLEAVED_ATTRIBS_NV"
	enumLookup[VERTEX_SUBROUTINE] = "VERTEX_SUBROUTINE"
	enumLookup[TEXTURE_IMAGE_FORMAT] = "TEXTURE_IMAGE_FORMAT"
	enumLookup[SLUMINANCE] = "SLUMINANCE"
	enumLookup[SHADER_IMAGE_ACCESS_BARRIER_BIT] = "SHADER_IMAGE_ACCESS_BARRIER_BIT"
	enumLookup[COLOR_ATTACHMENT11] = "COLOR_ATTACHMENT11"
	enumLookup[LINEAR_DETAIL_SGIS] = "LINEAR_DETAIL_SGIS"
	enumLookup[FOG_COORD_ARRAY_STRIDE] = "FOG_COORD_ARRAY_STRIDE"
	enumLookup[MAX_TESS_CONTROL_UNIFORM_COMPONENTS] = "MAX_TESS_CONTROL_UNIFORM_COMPONENTS"
	enumLookup[SPRITE_AXIS_SGIX] = "SPRITE_AXIS_SGIX"
	enumLookup[VERTEX_ATTRIB_ARRAY_LONG] = "VERTEX_ATTRIB_ARRAY_LONG"
	enumLookup[GEOMETRY_VERTICES_OUT] = "GEOMETRY_VERTICES_OUT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES"
	enumLookup[MAX_SERVER_WAIT_TIMEOUT_APPLE] = "MAX_SERVER_WAIT_TIMEOUT_APPLE"
	enumLookup[CONSTANT_EXT] = "CONSTANT_EXT"
	enumLookup[WEIGHT_ARRAY_STRIDE_OES] = "WEIGHT_ARRAY_STRIDE_OES"
	enumLookup[QUERY_RESULT_AVAILABLE_ARB] = "QUERY_RESULT_AVAILABLE_ARB"
	enumLookup[INDEX_ARRAY_BUFFER_BINDING] = "INDEX_ARRAY_BUFFER_BINDING"
	enumLookup[TEXTURE_SWIZZLE_R] = "TEXTURE_SWIZZLE_R"
	enumLookup[DIFFUSE] = "DIFFUSE"
	enumLookup[SECONDARY_COLOR_ARRAY_SIZE_EXT] = "SECONDARY_COLOR_ARRAY_SIZE_EXT"
	enumLookup[EDGEFLAG_BIT_PGI] = "EDGEFLAG_BIT_PGI"
	enumLookup[PIXEL_TILE_GRID_WIDTH_SGIX] = "PIXEL_TILE_GRID_WIDTH_SGIX"
	enumLookup[ARRAY_ELEMENT_LOCK_FIRST_EXT] = "ARRAY_ELEMENT_LOCK_FIRST_EXT"
	enumLookup[TEXTURE_RANGE_POINTER_APPLE] = "TEXTURE_RANGE_POINTER_APPLE"
	enumLookup[LUMINANCE_ALPHA_FLOAT32_ATI] = "LUMINANCE_ALPHA_FLOAT32_ATI"
	enumLookup[POINT_SPRITE_R_MODE_NV] = "POINT_SPRITE_R_MODE_NV"
	enumLookup[PIXEL_PACK_BUFFER_BINDING_ARB] = "PIXEL_PACK_BUFFER_BINDING_ARB"
	enumLookup[REG_28_ATI] = "REG_28_ATI"
	enumLookup[LUMINANCE32I_EXT] = "LUMINANCE32I_EXT"
	enumLookup[DATA_BUFFER_AMD] = "DATA_BUFFER_AMD"
	enumLookup[PIXEL_MAP_G_TO_G_SIZE] = "PIXEL_MAP_G_TO_G_SIZE"
	enumLookup[PIXEL_TILE_HEIGHT_SGIX] = "PIXEL_TILE_HEIGHT_SGIX"
	enumLookup[TRACK_MATRIX_TRANSFORM_NV] = "TRACK_MATRIX_TRANSFORM_NV"
	enumLookup[UNSIGNED_INT_24_8_MESA] = "UNSIGNED_INT_24_8_MESA"
	enumLookup[OUTPUT_TEXTURE_COORD6_EXT] = "OUTPUT_TEXTURE_COORD6_EXT"
	enumLookup[INVARIANT_VALUE_EXT] = "INVARIANT_VALUE_EXT"
	enumLookup[REG_7_ATI] = "REG_7_ATI"
	enumLookup[GEOMETRY_SHADER] = "GEOMETRY_SHADER"
	enumLookup[ACCUM_BUFFER_BIT] = "ACCUM_BUFFER_BIT"
	enumLookup[UNSIGNED_INT] = "UNSIGNED_INT"
	enumLookup[INTENSITY_FLOAT16_APPLE] = "INTENSITY_FLOAT16_APPLE"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR"
	enumLookup[LIGHT_MODEL_LOCAL_VIEWER] = "LIGHT_MODEL_LOCAL_VIEWER"
	enumLookup[DUAL_LUMINANCE8_SGIS] = "DUAL_LUMINANCE8_SGIS"
	enumLookup[CURRENT_BINORMAL_EXT] = "CURRENT_BINORMAL_EXT"
	enumLookup[RED_MIN_CLAMP_INGR] = "RED_MIN_CLAMP_INGR"
	enumLookup[WEIGHT_ARRAY_SIZE_OES] = "WEIGHT_ARRAY_SIZE_OES"
	enumLookup[MAX_PALETTE_MATRICES_OES] = "MAX_PALETTE_MATRICES_OES"
	enumLookup[VERTEX_PROGRAM_CALLBACK_MESA] = "VERTEX_PROGRAM_CALLBACK_MESA"
	enumLookup[MAX_SUBROUTINES] = "MAX_SUBROUTINES"
	enumLookup[QUERY_NO_WAIT_NV] = "QUERY_NO_WAIT_NV"
	enumLookup[COMPRESSED_RGBA_PVRTC_2BPPV2_IMG] = "COMPRESSED_RGBA_PVRTC_2BPPV2_IMG"
	enumLookup[UNPACK_SKIP_IMAGES] = "UNPACK_SKIP_IMAGES"
	enumLookup[TEXTURE_MAX_LEVEL_SGIS] = "TEXTURE_MAX_LEVEL_SGIS"
	enumLookup[OP_FRAC_EXT] = "OP_FRAC_EXT"
	enumLookup[OBJECT_SUBTYPE_ARB] = "OBJECT_SUBTYPE_ARB"
	enumLookup[FRAMEBUFFER_OES] = "FRAMEBUFFER_OES"
	enumLookup[LINEAR_SHARPEN_COLOR_SGIS] = "LINEAR_SHARPEN_COLOR_SGIS"
	enumLookup[TEXTURE14_ARB] = "TEXTURE14_ARB"
	enumLookup[VERTEX_ATTRIB_ARRAY1_NV] = "VERTEX_ATTRIB_ARRAY1_NV"
	enumLookup[CULL_FRAGMENT_NV] = "CULL_FRAGMENT_NV"
	enumLookup[FRAGMENT_COLOR_MATERIAL_SGIX] = "FRAGMENT_COLOR_MATERIAL_SGIX"
	enumLookup[UNSIGNED_INT_IMAGE_2D_RECT_EXT] = "UNSIGNED_INT_IMAGE_2D_RECT_EXT"
	enumLookup[RG8] = "RG8"
	enumLookup[PIXEL_COUNT_NV] = "PIXEL_COUNT_NV"
	enumLookup[NEGATE_BIT_ATI] = "NEGATE_BIT_ATI"
	enumLookup[COMPRESSED_SRGB_S3TC_DXT1_NV] = "COMPRESSED_SRGB_S3TC_DXT1_NV"
	enumLookup[COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2] = "COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	enumLookup[NUM_SAMPLE_COUNTS] = "NUM_SAMPLE_COUNTS"
	enumLookup[SECONDARY_COLOR_ARRAY_TYPE] = "SECONDARY_COLOR_ARRAY_TYPE"
	enumLookup[INT_VEC2_ARB] = "INT_VEC2_ARB"
	enumLookup[COLOR_ATTACHMENT8] = "COLOR_ATTACHMENT8"
	enumLookup[STENCIL_INDEX1_OES] = "STENCIL_INDEX1_OES"
	enumLookup[IMAGE_BINDING_LAYERED] = "IMAGE_BINDING_LAYERED"
	enumLookup[BUFFER_ACCESS_FLAGS] = "BUFFER_ACCESS_FLAGS"
	enumLookup[FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS] = "FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS"
	enumLookup[CURRENT_RASTER_DISTANCE] = "CURRENT_RASTER_DISTANCE"
	enumLookup[INDEX_BITS] = "INDEX_BITS"
	enumLookup[OUTPUT_TEXTURE_COORD22_EXT] = "OUTPUT_TEXTURE_COORD22_EXT"
	enumLookup[TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV] = "TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV"
	enumLookup[IMAGE_BINDING_FORMAT] = "IMAGE_BINDING_FORMAT"
	enumLookup[ADJACENT_PAIRS_NV] = "ADJACENT_PAIRS_NV"
	enumLookup[COMBINER_INPUT_NV] = "COMBINER_INPUT_NV"
	enumLookup[SOURCE0_RGB_EXT] = "SOURCE0_RGB_EXT"
	enumLookup[RENDERBUFFER_WIDTH] = "RENDERBUFFER_WIDTH"
	enumLookup[IMAGE_2D_ARRAY_EXT] = "IMAGE_2D_ARRAY_EXT"
	enumLookup[MAP_UNSYNCHRONIZED_BIT] = "MAP_UNSYNCHRONIZED_BIT"
	enumLookup[C4F_N3F_V3F] = "C4F_N3F_V3F"
	enumLookup[MIRRORED_REPEAT_ARB] = "MIRRORED_REPEAT_ARB"
	enumLookup[VERTEX_ARRAY_PARALLEL_POINTERS_INTEL] = "VERTEX_ARRAY_PARALLEL_POINTERS_INTEL"
	enumLookup[FOG_COORDINATE_SOURCE] = "FOG_COORDINATE_SOURCE"
	enumLookup[RGBA_FLOAT_MODE_ARB] = "RGBA_FLOAT_MODE_ARB"
	enumLookup[READ_WRITE] = "READ_WRITE"
	enumLookup[PERFMON_RESULT_AMD] = "PERFMON_RESULT_AMD"
	enumLookup[COLOR_ATTACHMENT5] = "COLOR_ATTACHMENT5"
	enumLookup[PROXY_TEXTURE_2D_MULTISAMPLE] = "PROXY_TEXTURE_2D_MULTISAMPLE"
	enumLookup[DEBUG_CATEGORY_WINDOW_SYSTEM_AMD] = "DEBUG_CATEGORY_WINDOW_SYSTEM_AMD"
	enumLookup[TEXTURE_BINDING_3D] = "TEXTURE_BINDING_3D"
	enumLookup[EDGE_FLAG_ARRAY_STRIDE_EXT] = "EDGE_FLAG_ARRAY_STRIDE_EXT"
	enumLookup[TEXTURE_MAX_CLAMP_T_SGIX] = "TEXTURE_MAX_CLAMP_T_SGIX"
	enumLookup[VERTEX_ATTRIB_ARRAY2_NV] = "VERTEX_ATTRIB_ARRAY2_NV"
	enumLookup[LUMINANCE_ALPHA8I_EXT] = "LUMINANCE_ALPHA8I_EXT"
	enumLookup[ALPHA8_EXT] = "ALPHA8_EXT"
	enumLookup[UNKNOWN_CONTEXT_RESET_ARB] = "UNKNOWN_CONTEXT_RESET_ARB"
	enumLookup[MATRIX_INDEX_ARRAY_STRIDE_ARB] = "MATRIX_INDEX_ARRAY_STRIDE_ARB"
	enumLookup[REG_12_ATI] = "REG_12_ATI"
	enumLookup[COLOR_MATERIAL_PARAMETER] = "COLOR_MATERIAL_PARAMETER"
	enumLookup[MAX_SHADER_STORAGE_BLOCK_SIZE] = "MAX_SHADER_STORAGE_BLOCK_SIZE"
	enumLookup[TESS_CONTROL_SHADER_BIT] = "TESS_CONTROL_SHADER_BIT"
	enumLookup[TEXTURE_TOO_LARGE_EXT] = "TEXTURE_TOO_LARGE_EXT"
	enumLookup[MODELVIEW22_ARB] = "MODELVIEW22_ARB"
	enumLookup[PN_TRIANGLES_NORMAL_MODE_ATI] = "PN_TRIANGLES_NORMAL_MODE_ATI"
	enumLookup[COLOR_ATTACHMENT_EXT] = "COLOR_ATTACHMENT_EXT"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR"
	enumLookup[TEXTURE10] = "TEXTURE10"
	enumLookup[MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV] = "MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV"
	enumLookup[PATH_TERMINAL_DASH_CAP_NV] = "PATH_TERMINAL_DASH_CAP_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY7_NV] = "VERTEX_ATTRIB_ARRAY7_NV"
	enumLookup[PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI] = "PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI"
	enumLookup[STATE_RESTORE] = "STATE_RESTORE"
	enumLookup[PROXY_TEXTURE_COLOR_TABLE_SGI] = "PROXY_TEXTURE_COLOR_TABLE_SGI"
	enumLookup[OUTPUT_TEXTURE_COORD5_EXT] = "OUTPUT_TEXTURE_COORD5_EXT"
	enumLookup[POST_CONVOLUTION_RED_BIAS] = "POST_CONVOLUTION_RED_BIAS"
	enumLookup[INDEX_TEST_REF_EXT] = "INDEX_TEST_REF_EXT"
	enumLookup[CURRENT_FOG_COORD] = "CURRENT_FOG_COORD"
	enumLookup[OUTPUT_TEXTURE_COORD15_EXT] = "OUTPUT_TEXTURE_COORD15_EXT"
	enumLookup[ELEMENT_ARRAY_BUFFER_ARB] = "ELEMENT_ARRAY_BUFFER_ARB"
	enumLookup[COMPRESSED_SRGB_S3TC_DXT1_EXT] = "COMPRESSED_SRGB_S3TC_DXT1_EXT"
	enumLookup[STENCIL_ATTACHMENT_OES] = "STENCIL_ATTACHMENT_OES"
	enumLookup[COMPATIBLE_SUBROUTINES] = "COMPATIBLE_SUBROUTINES"
	enumLookup[RGBA8_OES] = "RGBA8_OES"
	enumLookup[VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE] = "VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE"
	enumLookup[READ_FRAMEBUFFER_ANGLE] = "READ_FRAMEBUFFER_ANGLE"
	enumLookup[FIELDS_NV] = "FIELDS_NV"
	enumLookup[PATH_FILL_MASK_NV] = "PATH_FILL_MASK_NV"
	enumLookup[NATIVE_GRAPHICS_HANDLE_PGI] = "NATIVE_GRAPHICS_HANDLE_PGI"
	enumLookup[POLYGON_STIPPLE] = "POLYGON_STIPPLE"
	enumLookup[MAX_PIXEL_MAP_TABLE] = "MAX_PIXEL_MAP_TABLE"
	enumLookup[SAMPLE_COVERAGE_INVERT] = "SAMPLE_COVERAGE_INVERT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING] = "FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING"
	enumLookup[MAP2_VERTEX_ATTRIB14_4_NV] = "MAP2_VERTEX_ATTRIB14_4_NV"
	enumLookup[MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT] = "MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT"
	enumLookup[COMPRESSED_RGBA_ASTC_10x6_KHR] = "COMPRESSED_RGBA_ASTC_10x6_KHR"
	enumLookup[ALL_SHADER_BITS_EXT] = "ALL_SHADER_BITS_EXT"
	enumLookup[DRAW_BUFFER14_NV] = "DRAW_BUFFER14_NV"
	enumLookup[DOUBLE_VEC3_EXT] = "DOUBLE_VEC3_EXT"
	enumLookup[RG32I] = "RG32I"
	enumLookup[FRAGMENT_LIGHT3_SGIX] = "FRAGMENT_LIGHT3_SGIX"
	enumLookup[COLOR_SUM] = "COLOR_SUM"
	enumLookup[DISCARD_ATI] = "DISCARD_ATI"
	enumLookup[FRAMEBUFFER_ATTACHMENT_OBJECT_NAME] = "FRAMEBUFFER_ATTACHMENT_OBJECT_NAME"
	enumLookup[PATH_STROKE_COVER_MODE_NV] = "PATH_STROKE_COVER_MODE_NV"
	enumLookup[SRC_ALPHA] = "SRC_ALPHA"
	enumLookup[SPECULAR] = "SPECULAR"
	enumLookup[R32UI] = "R32UI"
	enumLookup[FRAGMENT_LIGHT2_SGIX] = "FRAGMENT_LIGHT2_SGIX"
	enumLookup[NORMAL_MAP_OES] = "NORMAL_MAP_OES"
	enumLookup[PACK_INVERT_MESA] = "PACK_INVERT_MESA"
	enumLookup[VIDEO_BUFFER_NV] = "VIDEO_BUFFER_NV"
	enumLookup[GEOMETRY_SUBROUTINE] = "GEOMETRY_SUBROUTINE"
	enumLookup[EXT_rescale_normal] = "EXT_rescale_normal"
	enumLookup[BLEND_EQUATION_EXT] = "BLEND_EQUATION_EXT"
	enumLookup[IMAGE_1D] = "IMAGE_1D"
	enumLookup[TEXTURE_BINDING_RECTANGLE] = "TEXTURE_BINDING_RECTANGLE"
	enumLookup[SCALAR_EXT] = "SCALAR_EXT"
	enumLookup[TEXTURE_COMPARE_FUNC_EXT] = "TEXTURE_COMPARE_FUNC_EXT"
	enumLookup[FRAMEBUFFER_COMPLETE_OES] = "FRAMEBUFFER_COMPLETE_OES"
	enumLookup[DRAW_BUFFER5] = "DRAW_BUFFER5"
	enumLookup[RGBA32I] = "RGBA32I"
	enumLookup[INT8_NV] = "INT8_NV"
	enumLookup[CONDITION_SATISFIED] = "CONDITION_SATISFIED"
	enumLookup[LIST_BIT] = "LIST_BIT"
	enumLookup[VERTEX_SHADER_BIT] = "VERTEX_SHADER_BIT"
	enumLookup[READ_PIXELS_TYPE] = "READ_PIXELS_TYPE"
	enumLookup[DISPATCH_INDIRECT_BUFFER_BINDING] = "DISPATCH_INDIRECT_BUFFER_BINDING"
	enumLookup[COMMAND_BARRIER_BIT_EXT] = "COMMAND_BARRIER_BIT_EXT"
	enumLookup[BLUE_BITS] = "BLUE_BITS"
	enumLookup[MODELVIEW24_ARB] = "MODELVIEW24_ARB"
	enumLookup[DEPENDENT_GB_TEXTURE_2D_NV] = "DEPENDENT_GB_TEXTURE_2D_NV"
	enumLookup[MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB] = "MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB"
	enumLookup[COLOR_ATTACHMENT7_NV] = "COLOR_ATTACHMENT7_NV"
	enumLookup[COLOR_MATRIX_STACK_DEPTH] = "COLOR_MATRIX_STACK_DEPTH"
	enumLookup[SEPARATE_SPECULAR_COLOR_EXT] = "SEPARATE_SPECULAR_COLOR_EXT"
	enumLookup[TRANSPOSE_PROJECTION_MATRIX_ARB] = "TRANSPOSE_PROJECTION_MATRIX_ARB"
	enumLookup[TEXTURE_FILTER_CONTROL_EXT] = "TEXTURE_FILTER_CONTROL_EXT"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Y_ARB] = "TEXTURE_CUBE_MAP_POSITIVE_Y_ARB"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT"
	enumLookup[TESS_GEN_MODE] = "TESS_GEN_MODE"
	enumLookup[ELEMENT_ARRAY_LENGTH_NV] = "ELEMENT_ARRAY_LENGTH_NV"
	enumLookup[PATH_FOG_GEN_MODE_NV] = "PATH_FOG_GEN_MODE_NV"
	enumLookup[QUAD_INTENSITY8_SGIS] = "QUAD_INTENSITY8_SGIS"
	enumLookup[VERTEX_PRECLIP_SGIX] = "VERTEX_PRECLIP_SGIX"
	enumLookup[OBJECT_COMPILE_STATUS_ARB] = "OBJECT_COMPILE_STATUS_ARB"
	enumLookup[FRAMEBUFFER_BINDING_EXT] = "FRAMEBUFFER_BINDING_EXT"
	enumLookup[STRICT_SCISSOR_HINT_PGI] = "STRICT_SCISSOR_HINT_PGI"
	enumLookup[ONE_MINUS_SRC_COLOR] = "ONE_MINUS_SRC_COLOR"
	enumLookup[EVAL_VERTEX_ATTRIB7_NV] = "EVAL_VERTEX_ATTRIB7_NV"
	enumLookup[VERTEX_ARRAY_BUFFER_BINDING] = "VERTEX_ARRAY_BUFFER_BINDING"
	enumLookup[REG_20_ATI] = "REG_20_ATI"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER] = "TRANSFORM_FEEDBACK_BUFFER"
	enumLookup[INDEX_ARRAY_ADDRESS_NV] = "INDEX_ARRAY_ADDRESS_NV"
	enumLookup[FONT_UNDERLINE_POSITION_BIT_NV] = "FONT_UNDERLINE_POSITION_BIT_NV"
	enumLookup[COMPUTE_SUBROUTINE_UNIFORM] = "COMPUTE_SUBROUTINE_UNIFORM"
	enumLookup[QUERY_WAIT] = "QUERY_WAIT"
	enumLookup[LAST_VERTEX_CONVENTION_EXT] = "LAST_VERTEX_CONVENTION_EXT"
	enumLookup[TEXTURE_SAMPLES_IMG] = "TEXTURE_SAMPLES_IMG"
	enumLookup[CLIP_DISTANCE1] = "CLIP_DISTANCE1"
	enumLookup[NUM_PROGRAM_BINARY_FORMATS] = "NUM_PROGRAM_BINARY_FORMATS"
	enumLookup[DEPTH_TEXTURE_MODE_ARB] = "DEPTH_TEXTURE_MODE_ARB"
	enumLookup[MAP1_INDEX] = "MAP1_INDEX"
	enumLookup[TEXTURE2] = "TEXTURE2"
	enumLookup[SIGNED_LUMINANCE_ALPHA_NV] = "SIGNED_LUMINANCE_ALPHA_NV"
	enumLookup[INDEX_ARRAY_POINTER_EXT] = "INDEX_ARRAY_POINTER_EXT"
	enumLookup[RED_MAX_CLAMP_INGR] = "RED_MAX_CLAMP_INGR"
	enumLookup[ACTIVE_VERTEX_UNITS_ARB] = "ACTIVE_VERTEX_UNITS_ARB"
	enumLookup[OUTPUT_TEXTURE_COORD18_EXT] = "OUTPUT_TEXTURE_COORD18_EXT"
	enumLookup[COMPRESSED_LUMINANCE_ARB] = "COMPRESSED_LUMINANCE_ARB"
	enumLookup[FLOAT_MAT3x2] = "FLOAT_MAT3x2"
	enumLookup[VERSION_1_1] = "VERSION_1_1"
	enumLookup[VERTEX_PROGRAM_BINDING_NV] = "VERTEX_PROGRAM_BINDING_NV"
	enumLookup[MAX_GEOMETRY_OUTPUT_VERTICES] = "MAX_GEOMETRY_OUTPUT_VERTICES"
	enumLookup[ATTRIB_ARRAY_TYPE_NV] = "ATTRIB_ARRAY_TYPE_NV"
	enumLookup[RELATIVE_QUADRATIC_CURVE_TO_NV] = "RELATIVE_QUADRATIC_CURVE_TO_NV"
	enumLookup[DEBUG_SEVERITY_LOW_ARB] = "DEBUG_SEVERITY_LOW_ARB"
	enumLookup[DEPTH_COMPONENT24_ARB] = "DEPTH_COMPONENT24_ARB"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Y_EXT] = "TEXTURE_CUBE_MAP_POSITIVE_Y_EXT"
	enumLookup[MAP1_VERTEX_ATTRIB9_4_NV] = "MAP1_VERTEX_ATTRIB9_4_NV"
	enumLookup[TRACE_TEXTURES_BIT_MESA] = "TRACE_TEXTURES_BIT_MESA"
	enumLookup[OFFSET_HILO_TEXTURE_RECTANGLE_NV] = "OFFSET_HILO_TEXTURE_RECTANGLE_NV"
	enumLookup[SAMPLER_2D_RECT_ARB] = "SAMPLER_2D_RECT_ARB"
	enumLookup[RASTER_POSITION_UNCLIPPED_IBM] = "RASTER_POSITION_UNCLIPPED_IBM"
	enumLookup[LIGHT_MODEL_AMBIENT] = "LIGHT_MODEL_AMBIENT"
	enumLookup[STENCIL_INDEX8_EXT] = "STENCIL_INDEX8_EXT"
	enumLookup[MAX_COMPUTE_WORK_GROUP_SIZE] = "MAX_COMPUTE_WORK_GROUP_SIZE"
	enumLookup[COMPRESSED_SRGB8_ETC2] = "COMPRESSED_SRGB8_ETC2"
	enumLookup[MAP_INVALIDATE_RANGE_BIT_EXT] = "MAP_INVALIDATE_RANGE_BIT_EXT"
	enumLookup[READ_BUFFER] = "READ_BUFFER"
	enumLookup[MAX_CLIENT_ATTRIB_STACK_DEPTH] = "MAX_CLIENT_ATTRIB_STACK_DEPTH"
	enumLookup[R] = "R"
	enumLookup[HISTOGRAM_WIDTH_EXT] = "HISTOGRAM_WIDTH_EXT"
	enumLookup[OPERAND1_RGB_EXT] = "OPERAND1_RGB_EXT"
	enumLookup[DRAW_BUFFER4] = "DRAW_BUFFER4"
	enumLookup[QUARTER_BIT_ATI] = "QUARTER_BIT_ATI"
	enumLookup[MIN_MAP_BUFFER_ALIGNMENT] = "MIN_MAP_BUFFER_ALIGNMENT"
	enumLookup[COMPRESSED_RGB8_ETC2] = "COMPRESSED_RGB8_ETC2"
	enumLookup[COMPILE_AND_EXECUTE] = "COMPILE_AND_EXECUTE"
	enumLookup[UNSIGNED_IDENTITY_NV] = "UNSIGNED_IDENTITY_NV"
	enumLookup[LIGHT_MODEL_SPECULAR_VECTOR_APPLE] = "LIGHT_MODEL_SPECULAR_VECTOR_APPLE"
	enumLookup[CLAMP_FRAGMENT_COLOR] = "CLAMP_FRAGMENT_COLOR"
	enumLookup[MAX_SAMPLE_MASK_WORDS] = "MAX_SAMPLE_MASK_WORDS"
	enumLookup[FUNC_SUBTRACT_EXT] = "FUNC_SUBTRACT_EXT"
	enumLookup[FOG_COORDINATE_ARRAY_TYPE_EXT] = "FOG_COORDINATE_ARRAY_TYPE_EXT"
	enumLookup[OPERAND3_RGB_NV] = "OPERAND3_RGB_NV"
	enumLookup[RGB32F_ARB] = "RGB32F_ARB"
	enumLookup[CON_15_ATI] = "CON_15_ATI"
	enumLookup[MAX_TESS_PATCH_COMPONENTS] = "MAX_TESS_PATCH_COMPONENTS"
	enumLookup[REFERENCED_BY_FRAGMENT_SHADER] = "REFERENCED_BY_FRAGMENT_SHADER"
	enumLookup[STENCIL_COMPONENTS] = "STENCIL_COMPONENTS"
	enumLookup[VERTEX_BINDING_STRIDE] = "VERTEX_BINDING_STRIDE"
	enumLookup[COLOR_ATTACHMENT8_EXT] = "COLOR_ATTACHMENT8_EXT"
	enumLookup[COMPRESSED_RGBA_ASTC_10x10_KHR] = "COMPRESSED_RGBA_ASTC_10x10_KHR"
	enumLookup[EVAL_VERTEX_ATTRIB9_NV] = "EVAL_VERTEX_ATTRIB9_NV"
	enumLookup[CONVEX_HULL_NV] = "CONVEX_HULL_NV"
	enumLookup[SHADER_STORAGE_BUFFER_START] = "SHADER_STORAGE_BUFFER_START"
	enumLookup[DEBUG_TYPE_ERROR_ARB] = "DEBUG_TYPE_ERROR_ARB"
	enumLookup[ALPHA_MIN_CLAMP_INGR] = "ALPHA_MIN_CLAMP_INGR"
	enumLookup[MATRIX4_ARB] = "MATRIX4_ARB"
	enumLookup[HALF_FLOAT_ARB] = "HALF_FLOAT_ARB"
	enumLookup[MINMAX_SINK_EXT] = "MINMAX_SINK_EXT"
	enumLookup[R8_EXT] = "R8_EXT"
	enumLookup[MANUAL_GENERATE_MIPMAP] = "MANUAL_GENERATE_MIPMAP"
	enumLookup[ALPHA_FLOAT16_APPLE] = "ALPHA_FLOAT16_APPLE"
	enumLookup[VERTEX_ATTRIB_ARRAY_BUFFER_BINDING] = "VERTEX_ATTRIB_ARRAY_BUFFER_BINDING"
	enumLookup[CON_4_ATI] = "CON_4_ATI"
	enumLookup[SGIX_resample] = "SGIX_resample"
	enumLookup[MAX_DRAW_BUFFERS_NV] = "MAX_DRAW_BUFFERS_NV"
	enumLookup[COMPRESSED_SLUMINANCE] = "COMPRESSED_SLUMINANCE"
	enumLookup[MAP2_TEXTURE_COORD_4] = "MAP2_TEXTURE_COORD_4"
	enumLookup[POSITION] = "POSITION"
	enumLookup[FRAMEBUFFER_BLEND] = "FRAMEBUFFER_BLEND"
	enumLookup[FOG_COORD_SRC] = "FOG_COORD_SRC"
	enumLookup[TEXTURE_RECTANGLE_NV] = "TEXTURE_RECTANGLE_NV"
	enumLookup[MAX_VERTEX_STREAMS_ATI] = "MAX_VERTEX_STREAMS_ATI"
	enumLookup[NORMAL_ARRAY_BUFFER_BINDING] = "NORMAL_ARRAY_BUFFER_BINDING"
	enumLookup[BIAS_BIT_ATI] = "BIAS_BIT_ATI"
	enumLookup[DRAW_INDIRECT_LENGTH_NV] = "DRAW_INDIRECT_LENGTH_NV"
	enumLookup[SRGB_DECODE_ARB] = "SRGB_DECODE_ARB"
	enumLookup[VERTEX_PROGRAM_TWO_SIDE] = "VERTEX_PROGRAM_TWO_SIDE"
	enumLookup[OFFSET_TEXTURE_BIAS_NV] = "OFFSET_TEXTURE_BIAS_NV"
	enumLookup[OUTPUT_TEXTURE_COORD12_EXT] = "OUTPUT_TEXTURE_COORD12_EXT"
	enumLookup[RENDERBUFFER_BLUE_SIZE_EXT] = "RENDERBUFFER_BLUE_SIZE_EXT"
	enumLookup[GLYPH_VERTICAL_BEARING_Y_BIT_NV] = "GLYPH_VERTICAL_BEARING_Y_BIT_NV"
	enumLookup[TIMEOUT_EXPIRED_APPLE] = "TIMEOUT_EXPIRED_APPLE"
	enumLookup[COMPUTE_SHADER] = "COMPUTE_SHADER"
	enumLookup[EXT_point_parameters] = "EXT_point_parameters"
	enumLookup[TRACK_MATRIX_NV] = "TRACK_MATRIX_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_TYPE] = "VERTEX_ATTRIB_ARRAY_TYPE"
	enumLookup[DRAW_BUFFER15_ATI] = "DRAW_BUFFER15_ATI"
	enumLookup[DEPTH24_STENCIL8_EXT] = "DEPTH24_STENCIL8_EXT"
	enumLookup[BITMAP_TOKEN] = "BITMAP_TOKEN"
	enumLookup[MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT] = "MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT"
	enumLookup[POINT_TOKEN] = "POINT_TOKEN"
	enumLookup[LUMINANCE16] = "LUMINANCE16"
	enumLookup[MINMAX_SINK] = "MINMAX_SINK"
	enumLookup[COMBINE_EXT] = "COMBINE_EXT"
	enumLookup[MATRIX14_ARB] = "MATRIX14_ARB"
	enumLookup[SRC0_RGB] = "SRC0_RGB"
	enumLookup[INT_SAMPLER_1D_ARRAY] = "INT_SAMPLER_1D_ARRAY"
	enumLookup[TEXTURE_COORD_ARRAY_SIZE_EXT] = "TEXTURE_COORD_ARRAY_SIZE_EXT"
	enumLookup[INT_SAMPLER_CUBE_EXT] = "INT_SAMPLER_CUBE_EXT"
	enumLookup[MAX_NAME_LENGTH] = "MAX_NAME_LENGTH"
	enumLookup[MATRIX22_ARB] = "MATRIX22_ARB"
	enumLookup[RENDERBUFFER_GREEN_SIZE_EXT] = "RENDERBUFFER_GREEN_SIZE_EXT"
	enumLookup[RGBA_MODE] = "RGBA_MODE"
	enumLookup[CONSTANT_ATTENUATION] = "CONSTANT_ATTENUATION"
	enumLookup[MAX] = "MAX"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_Z_ARB] = "TEXTURE_CUBE_MAP_POSITIVE_Z_ARB"
	enumLookup[RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV] = "RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV"
	enumLookup[VERSION_1_2] = "VERSION_1_2"
	enumLookup[HISTOGRAM_RED_SIZE] = "HISTOGRAM_RED_SIZE"
	enumLookup[MATRIX4_NV] = "MATRIX4_NV"
	enumLookup[MAP1_VERTEX_ATTRIB0_4_NV] = "MAP1_VERTEX_ATTRIB0_4_NV"
	enumLookup[SWIZZLE_STRQ_DQ_ATI] = "SWIZZLE_STRQ_DQ_ATI"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_START_EXT] = "TRANSFORM_FEEDBACK_BUFFER_START_EXT"
	enumLookup[ALPHA16_SNORM] = "ALPHA16_SNORM"
	enumLookup[TEXCOORD1_BIT_PGI] = "TEXCOORD1_BIT_PGI"
	enumLookup[FRAMEBUFFER_BARRIER_BIT] = "FRAMEBUFFER_BARRIER_BIT"
	enumLookup[COLOR_CLEAR_VALUE] = "COLOR_CLEAR_VALUE"
	enumLookup[RENDER_MODE] = "RENDER_MODE"
	enumLookup[RGB4_EXT] = "RGB4_EXT"
	enumLookup[SECONDARY_COLOR_ARRAY_EXT] = "SECONDARY_COLOR_ARRAY_EXT"
	enumLookup[MAP1_VERTEX_ATTRIB6_4_NV] = "MAP1_VERTEX_ATTRIB6_4_NV"
	enumLookup[VERTEX_ARRAY_BUFFER_BINDING_ARB] = "VERTEX_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[SKIP_MISSING_GLYPH_NV] = "SKIP_MISSING_GLYPH_NV"
	enumLookup[TEXTURE_BUFFER_SIZE] = "TEXTURE_BUFFER_SIZE"
	enumLookup[SHADER_BINARY_DMP] = "SHADER_BINARY_DMP"
	enumLookup[GL_422_AVERAGE_EXT] = "GL_422_AVERAGE_EXT"
	enumLookup[VERTEX_SHADER_VARIANTS_EXT] = "VERTEX_SHADER_VARIANTS_EXT"
	enumLookup[WRITE_PIXEL_DATA_RANGE_NV] = "WRITE_PIXEL_DATA_RANGE_NV"
	enumLookup[ALREADY_SIGNALED_APPLE] = "ALREADY_SIGNALED_APPLE"
	enumLookup[NAME_LENGTH] = "NAME_LENGTH"
	enumLookup[MAX_VERTEX_HINT_PGI] = "MAX_VERTEX_HINT_PGI"
	enumLookup[DRAW_BUFFER11] = "DRAW_BUFFER11"
	enumLookup[LUMINANCE_ALPHA_INTEGER_EXT] = "LUMINANCE_ALPHA_INTEGER_EXT"
	enumLookup[UNPACK_SKIP_ROWS] = "UNPACK_SKIP_ROWS"
	enumLookup[CONVOLUTION_2D] = "CONVOLUTION_2D"
	enumLookup[OUTPUT_TEXTURE_COORD30_EXT] = "OUTPUT_TEXTURE_COORD30_EXT"
	enumLookup[VERTEX_SHADER_ARB] = "VERTEX_SHADER_ARB"
	enumLookup[TEXTURE_SWIZZLE_A] = "TEXTURE_SWIZZLE_A"
	enumLookup[INT_IMAGE_1D_ARRAY_EXT] = "INT_IMAGE_1D_ARRAY_EXT"
	enumLookup[SGIS_point_parameters] = "SGIS_point_parameters"
	enumLookup[STENCIL_FAIL] = "STENCIL_FAIL"
	enumLookup[RED] = "RED"
	enumLookup[POST_CONVOLUTION_ALPHA_BIAS_EXT] = "POST_CONVOLUTION_ALPHA_BIAS_EXT"
	enumLookup[DOT3_RGBA_ARB] = "DOT3_RGBA_ARB"
	enumLookup[FRAGMENT_PROGRAM_CALLBACK_MESA] = "FRAGMENT_PROGRAM_CALLBACK_MESA"
	enumLookup[VERTEX_ARRAY_TYPE] = "VERTEX_ARRAY_TYPE"
	enumLookup[MAX_RECTANGLE_TEXTURE_SIZE] = "MAX_RECTANGLE_TEXTURE_SIZE"
	enumLookup[OFFSET_TEXTURE_2D_SCALE_NV] = "OFFSET_TEXTURE_2D_SCALE_NV"
	enumLookup[RGBA2] = "RGBA2"
	enumLookup[UNSIGNED_INT_8_8_8_8_REV] = "UNSIGNED_INT_8_8_8_8_REV"
	enumLookup[SRGB_ALPHA] = "SRGB_ALPHA"
	enumLookup[COLOR_LOGIC_OP] = "COLOR_LOGIC_OP"
	enumLookup[FLOAT_MAT2x4] = "FLOAT_MAT2x4"
	enumLookup[SAMPLER_EXTERNAL_OES] = "SAMPLER_EXTERNAL_OES"
	enumLookup[GREEN_INTEGER] = "GREEN_INTEGER"
	enumLookup[VERTEX_SHADER_BIT_EXT] = "VERTEX_SHADER_BIT_EXT"
	enumLookup[AVERAGE_HP] = "AVERAGE_HP"
	enumLookup[OPERAND0_ALPHA] = "OPERAND0_ALPHA"
	enumLookup[TEXTURE_BORDER_VALUES_NV] = "TEXTURE_BORDER_VALUES_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_DIVISOR_NV] = "VERTEX_ATTRIB_ARRAY_DIVISOR_NV"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_MODE_NV] = "TRANSFORM_FEEDBACK_BUFFER_MODE_NV"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE] = "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE"
	enumLookup[VIEWPORT_BIT] = "VIEWPORT_BIT"
	enumLookup[UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER] = "UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER"
	enumLookup[VERTEX_ARRAY_RANGE_APPLE] = "VERTEX_ARRAY_RANGE_APPLE"
	enumLookup[RGBA_SNORM] = "RGBA_SNORM"
	enumLookup[COLOR_ATTACHMENT14_EXT] = "COLOR_ATTACHMENT14_EXT"
	enumLookup[TEXTURE_COVERAGE_SAMPLES_NV] = "TEXTURE_COVERAGE_SAMPLES_NV"
	enumLookup[TEXTURE_BINDING_2D_MULTISAMPLE] = "TEXTURE_BINDING_2D_MULTISAMPLE"
	enumLookup[COLOR_TABLE_BLUE_SIZE_SGI] = "COLOR_TABLE_BLUE_SIZE_SGI"
	enumLookup[VIEWPORT_INDEX_PROVOKING_VERTEX] = "VIEWPORT_INDEX_PROVOKING_VERTEX"
	enumLookup[INTENSITY_FLOAT16_ATI] = "INTENSITY_FLOAT16_ATI"
	enumLookup[TEXTURE_CUBE_MAP_SEAMLESS] = "TEXTURE_CUBE_MAP_SEAMLESS"
	enumLookup[DOUBLE_MAT3x4] = "DOUBLE_MAT3x4"
	enumLookup[UNSIGNED_INT_IMAGE_1D_ARRAY] = "UNSIGNED_INT_IMAGE_1D_ARRAY"
	enumLookup[SGIS_texture_select] = "SGIS_texture_select"
	enumLookup[PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI] = "PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI"
	enumLookup[MAX_TEXTURE_IMAGE_UNITS_NV] = "MAX_TEXTURE_IMAGE_UNITS_NV"
	enumLookup[WRITE_PIXEL_DATA_RANGE_LENGTH_NV] = "WRITE_PIXEL_DATA_RANGE_LENGTH_NV"
	enumLookup[REQUIRED_TEXTURE_IMAGE_UNITS_OES] = "REQUIRED_TEXTURE_IMAGE_UNITS_OES"
	enumLookup[GREEN_INTEGER_EXT] = "GREEN_INTEGER_EXT"
	enumLookup[MAX_COMPUTE_UNIFORM_BLOCKS] = "MAX_COMPUTE_UNIFORM_BLOCKS"
	enumLookup[COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2] = "COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2"
	enumLookup[ARRAY_BUFFER_BINDING_ARB] = "ARRAY_BUFFER_BINDING_ARB"
	enumLookup[MAX_VARYING_COMPONENTS_EXT] = "MAX_VARYING_COMPONENTS_EXT"
	enumLookup[LUMINANCE8_SNORM] = "LUMINANCE8_SNORM"
	enumLookup[TEXTURE_ENV_MODE] = "TEXTURE_ENV_MODE"
	enumLookup[READ_BUFFER_EXT] = "READ_BUFFER_EXT"
	enumLookup[MAX_TEXTURE_COORDS] = "MAX_TEXTURE_COORDS"
	enumLookup[DUDV_ATI] = "DUDV_ATI"
	enumLookup[UNIFORM_BLOCK_ACTIVE_UNIFORMS] = "UNIFORM_BLOCK_ACTIVE_UNIFORMS"
	enumLookup[CLIP_NEAR_HINT_PGI] = "CLIP_NEAR_HINT_PGI"
	enumLookup[UNPACK_LSB_FIRST] = "UNPACK_LSB_FIRST"
	enumLookup[SELECTION_BUFFER_SIZE] = "SELECTION_BUFFER_SIZE"
	enumLookup[LIGHT7] = "LIGHT7"
	enumLookup[RG16I] = "RG16I"
	enumLookup[TRANSFORM_FEEDBACK_VARYINGS] = "TRANSFORM_FEEDBACK_VARYINGS"
	enumLookup[RENDERBUFFER_BLUE_SIZE_OES] = "RENDERBUFFER_BLUE_SIZE_OES"
	enumLookup[TEXTURE_COORD_ARRAY_ADDRESS_NV] = "TEXTURE_COORD_ARRAY_ADDRESS_NV"
	enumLookup[TRANSFORM_BIT] = "TRANSFORM_BIT"
	enumLookup[TEXTURE_2D] = "TEXTURE_2D"
	enumLookup[WRAP_BORDER_SUN] = "WRAP_BORDER_SUN"
	enumLookup[TEXTURE28] = "TEXTURE28"
	enumLookup[TEXTURE_BINDING_CUBE_MAP_EXT] = "TEXTURE_BINDING_CUBE_MAP_EXT"
	enumLookup[VARIANT_ARRAY_POINTER_EXT] = "VARIANT_ARRAY_POINTER_EXT"
	enumLookup[UNSIGNED_INT_VEC3_EXT] = "UNSIGNED_INT_VEC3_EXT"
	enumLookup[UNSIGNED_INT_IMAGE_CUBE_EXT] = "UNSIGNED_INT_IMAGE_CUBE_EXT"
	enumLookup[ALLOW_DRAW_MEM_HINT_PGI] = "ALLOW_DRAW_MEM_HINT_PGI"
	enumLookup[FEEDBACK_BUFFER_TYPE] = "FEEDBACK_BUFFER_TYPE"
	enumLookup[VIEW_CLASS_16_BITS] = "VIEW_CLASS_16_BITS"
	enumLookup[TEXTURE_ALPHA_MODULATE_IMG] = "TEXTURE_ALPHA_MODULATE_IMG"
	enumLookup[SRGB8_ALPHA8] = "SRGB8_ALPHA8"
	enumLookup[NAMED_STRING_TYPE_ARB] = "NAMED_STRING_TYPE_ARB"
	enumLookup[TRANSFORM_FEEDBACK_PAUSED] = "TRANSFORM_FEEDBACK_PAUSED"
	enumLookup[RGB10_A2] = "RGB10_A2"
	enumLookup[UNSIGNED_SHORT_4_4_4_4] = "UNSIGNED_SHORT_4_4_4_4"
	enumLookup[TEXTURE4] = "TEXTURE4"
	enumLookup[UNPACK_ROW_BYTES_APPLE] = "UNPACK_ROW_BYTES_APPLE"
	enumLookup[FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT] = "FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT"
	enumLookup[TEXTURE_COORD_ARRAY_LENGTH_NV] = "TEXTURE_COORD_ARRAY_LENGTH_NV"
	enumLookup[INT_IMAGE_3D_EXT] = "INT_IMAGE_3D_EXT"
	enumLookup[NORMAL_MAP_ARB] = "NORMAL_MAP_ARB"
	enumLookup[MAX_TESS_CONTROL_INPUT_COMPONENTS] = "MAX_TESS_CONTROL_INPUT_COMPONENTS"
	enumLookup[MATRIX28_ARB] = "MATRIX28_ARB"
	enumLookup[EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD] = "EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD"
	enumLookup[TEXTURE_LOD_BIAS_S_SGIX] = "TEXTURE_LOD_BIAS_S_SGIX"
	enumLookup[VIEWPORT_SUBPIXEL_BITS] = "VIEWPORT_SUBPIXEL_BITS"
	enumLookup[PRIMARY_COLOR] = "PRIMARY_COLOR"
	enumLookup[MATRIX7_NV] = "MATRIX7_NV"
	enumLookup[SIGNED_RGB8_NV] = "SIGNED_RGB8_NV"
	enumLookup[MODULATE_ADD_ATI] = "MODULATE_ADD_ATI"
	enumLookup[MAX_PROGRAM_LOOP_COUNT_NV] = "MAX_PROGRAM_LOOP_COUNT_NV"
	enumLookup[POINT_SIZE_GRANULARITY] = "POINT_SIZE_GRANULARITY"
	enumLookup[RESAMPLE_ZERO_FILL_OML] = "RESAMPLE_ZERO_FILL_OML"
	enumLookup[SEPARATE_ATTRIBS_NV] = "SEPARATE_ATTRIBS_NV"
	enumLookup[FONT_X_MAX_BOUNDS_BIT_NV] = "FONT_X_MAX_BOUNDS_BIT_NV"
	enumLookup[LUMINANCE_FLOAT32_ATI] = "LUMINANCE_FLOAT32_ATI"
	enumLookup[MAX_DRAW_BUFFERS_ATI] = "MAX_DRAW_BUFFERS_ATI"
	enumLookup[MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV] = "MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV"
	enumLookup[POST_COLOR_MATRIX_RED_SCALE_SGI] = "POST_COLOR_MATRIX_RED_SCALE_SGI"
	enumLookup[MAX_ELEMENTS_VERTICES_EXT] = "MAX_ELEMENTS_VERTICES_EXT"
	enumLookup[MAP2_VERTEX_ATTRIB12_4_NV] = "MAP2_VERTEX_ATTRIB12_4_NV"
	enumLookup[DRAW_BUFFER12_ATI] = "DRAW_BUFFER12_ATI"
	enumLookup[INT_IMAGE_2D] = "INT_IMAGE_2D"
	enumLookup[ADD] = "ADD"
	enumLookup[TEXTURE_COMPRESSED_BLOCK_WIDTH] = "TEXTURE_COMPRESSED_BLOCK_WIDTH"
	enumLookup[REGISTER_COMBINERS_NV] = "REGISTER_COMBINERS_NV"
	enumLookup[RGB5_A1_OES] = "RGB5_A1_OES"
	enumLookup[MAX_PROGRAM_TEX_INSTRUCTIONS_ARB] = "MAX_PROGRAM_TEX_INSTRUCTIONS_ARB"
	enumLookup[TEXTURE_ALPHA_TYPE_ARB] = "TEXTURE_ALPHA_TYPE_ARB"
	enumLookup[MAX_VERTEX_UNIFORM_VECTORS] = "MAX_VERTEX_UNIFORM_VECTORS"
	enumLookup[PIXEL_MAP_R_TO_R_SIZE] = "PIXEL_MAP_R_TO_R_SIZE"
	enumLookup[DEPTH_COMPONENT24_OES] = "DEPTH_COMPONENT24_OES"
	enumLookup[VARIANT_ARRAY_EXT] = "VARIANT_ARRAY_EXT"
	enumLookup[CURRENT_QUERY_EXT] = "CURRENT_QUERY_EXT"
	enumLookup[VERTICAL_LINE_TO_NV] = "VERTICAL_LINE_TO_NV"
	enumLookup[POLYGON_TOKEN] = "POLYGON_TOKEN"
	enumLookup[SECONDARY_COLOR_ARRAY_SIZE] = "SECONDARY_COLOR_ARRAY_SIZE"
	enumLookup[COMBINER_AB_OUTPUT_NV] = "COMBINER_AB_OUTPUT_NV"
	enumLookup[SRC2_ALPHA] = "SRC2_ALPHA"
	enumLookup[MVP_MATRIX_EXT] = "MVP_MATRIX_EXT"
	enumLookup[OBJECT_SHADER_SOURCE_LENGTH_ARB] = "OBJECT_SHADER_SOURCE_LENGTH_ARB"
	enumLookup[COUNTER_RANGE_AMD] = "COUNTER_RANGE_AMD"
	enumLookup[TEXTURE_FORMAT_QCOM] = "TEXTURE_FORMAT_QCOM"
	enumLookup[PROXY_TEXTURE_2D_ARRAY_EXT] = "PROXY_TEXTURE_2D_ARRAY_EXT"
	enumLookup[PROGRAM_MATRIX_STACK_DEPTH_EXT] = "PROGRAM_MATRIX_STACK_DEPTH_EXT"
	enumLookup[RG8_SNORM] = "RG8_SNORM"
	enumLookup[VERTEX_ARRAY_LIST_STRIDE_IBM] = "VERTEX_ARRAY_LIST_STRIDE_IBM"
	enumLookup[MAX_NAME_STACK_DEPTH] = "MAX_NAME_STACK_DEPTH"
	enumLookup[OBJECT_DISTANCE_TO_POINT_SGIS] = "OBJECT_DISTANCE_TO_POINT_SGIS"
	enumLookup[MATRIX13_ARB] = "MATRIX13_ARB"
	enumLookup[UNIFORM_MATRIX_STRIDE] = "UNIFORM_MATRIX_STRIDE"
	enumLookup[HISTOGRAM_GREEN_SIZE] = "HISTOGRAM_GREEN_SIZE"
	enumLookup[VARIABLE_A_NV] = "VARIABLE_A_NV"
	enumLookup[SHADER_OPERATION_NV] = "SHADER_OPERATION_NV"
	enumLookup[RGB9_E5_EXT] = "RGB9_E5_EXT"
	enumLookup[MAP_INVALIDATE_RANGE_BIT] = "MAP_INVALIDATE_RANGE_BIT"
	enumLookup[TEXTURE_ENV_COLOR] = "TEXTURE_ENV_COLOR"
	enumLookup[OP_SUB_EXT] = "OP_SUB_EXT"
	enumLookup[SECONDARY_INTERPOLATOR_ATI] = "SECONDARY_INTERPOLATOR_ATI"
	enumLookup[COMPRESSED_SLUMINANCE_EXT] = "COMPRESSED_SLUMINANCE_EXT"
	enumLookup[FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER] = "FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER"
	enumLookup[CURRENT_TEXTURE_COORDS] = "CURRENT_TEXTURE_COORDS"
	enumLookup[FRAMEBUFFER_RENDERABLE_LAYERED] = "FRAMEBUFFER_RENDERABLE_LAYERED"
	enumLookup[TRANSPOSE_COLOR_MATRIX_ARB] = "TRANSPOSE_COLOR_MATRIX_ARB"
	enumLookup[MAP2_VERTEX_ATTRIB9_4_NV] = "MAP2_VERTEX_ATTRIB9_4_NV"
	enumLookup[VARIANT_VALUE_EXT] = "VARIANT_VALUE_EXT"
	enumLookup[FLOAT_VEC3] = "FLOAT_VEC3"
	enumLookup[OFFSET] = "OFFSET"
	enumLookup[CLIP_FAR_HINT_PGI] = "CLIP_FAR_HINT_PGI"
	enumLookup[LINEAR_SHARPEN_ALPHA_SGIS] = "LINEAR_SHARPEN_ALPHA_SGIS"
	enumLookup[IMAGE_TEXEL_SIZE] = "IMAGE_TEXEL_SIZE"
	enumLookup[MAX_PROGRAM_NATIVE_PARAMETERS_ARB] = "MAX_PROGRAM_NATIVE_PARAMETERS_ARB"
	enumLookup[REG_2_ATI] = "REG_2_ATI"
	enumLookup[MAX_VARYING_FLOATS] = "MAX_VARYING_FLOATS"
	enumLookup[FRAMEBUFFER_BINDING_OES] = "FRAMEBUFFER_BINDING_OES"
	enumLookup[MAX_GEOMETRY_OUTPUT_COMPONENTS] = "MAX_GEOMETRY_OUTPUT_COMPONENTS"
	enumLookup[COEFF] = "COEFF"
	enumLookup[AND] = "AND"
	enumLookup[COLOR_ARRAY_TYPE_EXT] = "COLOR_ARRAY_TYPE_EXT"
	enumLookup[MATRIX5_NV] = "MATRIX5_NV"
	enumLookup[MATRIX_INDEX_ARRAY_POINTER_OES] = "MATRIX_INDEX_ARRAY_POINTER_OES"
	enumLookup[ITALIC_BIT_NV] = "ITALIC_BIT_NV"
	enumLookup[MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS] = "MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS"
	enumLookup[STATIC_DRAW] = "STATIC_DRAW"
	enumLookup[PRESENT_TIME_NV] = "PRESENT_TIME_NV"
	enumLookup[RED_SNORM] = "RED_SNORM"
	enumLookup[UNSIGNED_INT_2_10_10_10_REV_EXT] = "UNSIGNED_INT_2_10_10_10_REV_EXT"
	enumLookup[RGB16_EXT] = "RGB16_EXT"
	enumLookup[MAP2_TANGENT_EXT] = "MAP2_TANGENT_EXT"
	enumLookup[SIGNED_INTENSITY_NV] = "SIGNED_INTENSITY_NV"
	enumLookup[STREAM_DRAW_ARB] = "STREAM_DRAW_ARB"
	enumLookup[COLOR_ATTACHMENT12_EXT] = "COLOR_ATTACHMENT12_EXT"
	enumLookup[COLOR_TABLE_LUMINANCE_SIZE] = "COLOR_TABLE_LUMINANCE_SIZE"
	enumLookup[TEXTURE_STORAGE_HINT_APPLE] = "TEXTURE_STORAGE_HINT_APPLE"
	enumLookup[WRITEONLY_RENDERING_QCOM] = "WRITEONLY_RENDERING_QCOM"
	enumLookup[ACTIVE_STENCIL_FACE_EXT] = "ACTIVE_STENCIL_FACE_EXT"
	enumLookup[GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV] = "GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV"
	enumLookup[FOG_COORD_ARRAY_LENGTH_NV] = "FOG_COORD_ARRAY_LENGTH_NV"
	enumLookup[QUERY_OBJECT_EXT] = "QUERY_OBJECT_EXT"
	enumLookup[TEXTURE_CLIPMAP_DEPTH_SGIX] = "TEXTURE_CLIPMAP_DEPTH_SGIX"
	enumLookup[OFFSET_TEXTURE_RECTANGLE_NV] = "OFFSET_TEXTURE_RECTANGLE_NV"
	enumLookup[FULL_RANGE_EXT] = "FULL_RANGE_EXT"
	enumLookup[PALETTE4_R5_G6_B5_OES] = "PALETTE4_R5_G6_B5_OES"
	enumLookup[LUMINANCE12_EXT] = "LUMINANCE12_EXT"
	enumLookup[NORMAL_ARRAY_EXT] = "NORMAL_ARRAY_EXT"
	enumLookup[TEXTURE8] = "TEXTURE8"
	enumLookup[VIBRANCE_BIAS_NV] = "VIBRANCE_BIAS_NV"
	enumLookup[RGB_422_APPLE] = "RGB_422_APPLE"
	enumLookup[INT_VEC2] = "INT_VEC2"
	enumLookup[STENCIL_INDEX4_EXT] = "STENCIL_INDEX4_EXT"
	enumLookup[COLOR_ATTACHMENT6_NV] = "COLOR_ATTACHMENT6_NV"
	enumLookup[UNSIGNED_INT_IMAGE_2D_EXT] = "UNSIGNED_INT_IMAGE_2D_EXT"
	enumLookup[SHADER_STORAGE_BUFFER_BINDING] = "SHADER_STORAGE_BUFFER_BINDING"
	enumLookup[PIXEL_MAP_I_TO_B_SIZE] = "PIXEL_MAP_I_TO_B_SIZE"
	enumLookup[TEXTURE18] = "TEXTURE18"
	enumLookup[FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV] = "FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV"
	enumLookup[CIRCULAR_TANGENT_ARC_TO_NV] = "CIRCULAR_TANGENT_ARC_TO_NV"
	enumLookup[PIXEL_BUFFER_BARRIER_BIT_EXT] = "PIXEL_BUFFER_BARRIER_BIT_EXT"
	enumLookup[VERTEX_ARRAY_SIZE] = "VERTEX_ARRAY_SIZE"
	enumLookup[MIRROR_CLAMP_TO_BORDER_EXT] = "MIRROR_CLAMP_TO_BORDER_EXT"
	enumLookup[OBJECT_ATTACHED_OBJECTS_ARB] = "OBJECT_ATTACHED_OBJECTS_ARB"
	enumLookup[INT16_VEC2_NV] = "INT16_VEC2_NV"
	enumLookup[CONVOLUTION_HEIGHT] = "CONVOLUTION_HEIGHT"
	enumLookup[TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL] = "TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL"
	enumLookup[RESAMPLE_REPLICATE_SGIX] = "RESAMPLE_REPLICATE_SGIX"
	enumLookup[MAX_VERTEX_ATTRIBS_ARB] = "MAX_VERTEX_ATTRIBS_ARB"
	enumLookup[FLOAT_RGB32_NV] = "FLOAT_RGB32_NV"
	enumLookup[DEBUG_SOURCE_APPLICATION_ARB] = "DEBUG_SOURCE_APPLICATION_ARB"
	enumLookup[RGB2_EXT] = "RGB2_EXT"
	enumLookup[MAP2_VERTEX_ATTRIB2_4_NV] = "MAP2_VERTEX_ATTRIB2_4_NV"
	enumLookup[BLEND_EQUATION_ALPHA_OES] = "BLEND_EQUATION_ALPHA_OES"
	enumLookup[MAX_VARYING_FLOATS_ARB] = "MAX_VARYING_FLOATS_ARB"
	enumLookup[UNSIGNED_INT16_VEC4_NV] = "UNSIGNED_INT16_VEC4_NV"
	enumLookup[DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB] = "DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB"
	enumLookup[TEXTURE18_ARB] = "TEXTURE18_ARB"
	enumLookup[TEXTURE_COMPRESSED_ARB] = "TEXTURE_COMPRESSED_ARB"
	enumLookup[OP_DOT4_EXT] = "OP_DOT4_EXT"
	enumLookup[FLOAT_MAT2] = "FLOAT_MAT2"
	enumLookup[SHADING_LANGUAGE_VERSION] = "SHADING_LANGUAGE_VERSION"
	enumLookup[TEXTURE_PRIORITY_EXT] = "TEXTURE_PRIORITY_EXT"
	enumLookup[SAMPLES_EXT] = "SAMPLES_EXT"
	enumLookup[WEIGHT_ARRAY_POINTER_OES] = "WEIGHT_ARRAY_POINTER_OES"
	enumLookup[SAMPLE_POSITION] = "SAMPLE_POSITION"
	enumLookup[COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB] = "COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB"
	enumLookup[RG8_EXT] = "RG8_EXT"
	enumLookup[MAX_RECTANGLE_TEXTURE_SIZE_NV] = "MAX_RECTANGLE_TEXTURE_SIZE_NV"
	enumLookup[MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT] = "MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT"
	enumLookup[EMISSION] = "EMISSION"
	enumLookup[MAX_COLOR_MATRIX_STACK_DEPTH_SGI] = "MAX_COLOR_MATRIX_STACK_DEPTH_SGI"
	enumLookup[PROGRAM_LENGTH_ARB] = "PROGRAM_LENGTH_ARB"
	enumLookup[WEIGHT_ARRAY_TYPE_ARB] = "WEIGHT_ARRAY_TYPE_ARB"
	enumLookup[INT_IMAGE_2D_RECT_EXT] = "INT_IMAGE_2D_RECT_EXT"
	enumLookup[PATH_GEN_COMPONENTS_NV] = "PATH_GEN_COMPONENTS_NV"
	enumLookup[PATH_STENCIL_DEPTH_OFFSET_UNITS_NV] = "PATH_STENCIL_DEPTH_OFFSET_UNITS_NV"
	enumLookup[DEBUG_OUTPUT] = "DEBUG_OUTPUT"
	enumLookup[DUAL_ALPHA8_SGIS] = "DUAL_ALPHA8_SGIS"
	enumLookup[FENCE_APPLE] = "FENCE_APPLE"
	enumLookup[UNIFORM_BLOCK_DATA_SIZE] = "UNIFORM_BLOCK_DATA_SIZE"
	enumLookup[IMAGE_BINDING_ACCESS_EXT] = "IMAGE_BINDING_ACCESS_EXT"
	enumLookup[ATOMIC_COUNTER_BARRIER_BIT_EXT] = "ATOMIC_COUNTER_BARRIER_BIT_EXT"
	enumLookup[HISTOGRAM_SINK_EXT] = "HISTOGRAM_SINK_EXT"
	enumLookup[TABLE_TOO_LARGE_EXT] = "TABLE_TOO_LARGE_EXT"
	enumLookup[PROGRAM_SEPARABLE] = "PROGRAM_SEPARABLE"
	enumLookup[INT_VEC3] = "INT_VEC3"
	enumLookup[PACK_COMPRESSED_BLOCK_WIDTH] = "PACK_COMPRESSED_BLOCK_WIDTH"
	enumLookup[OPERAND0_RGB_ARB] = "OPERAND0_RGB_ARB"
	enumLookup[MAX_TRACK_MATRICES_NV] = "MAX_TRACK_MATRICES_NV"
	enumLookup[DRAW_BUFFER15] = "DRAW_BUFFER15"
	enumLookup[UNSIGNED_INT_5_9_9_9_REV] = "UNSIGNED_INT_5_9_9_9_REV"
	enumLookup[MAX_COMPUTE_LOCAL_INVOCATIONS] = "MAX_COMPUTE_LOCAL_INVOCATIONS"
	enumLookup[LAYOUT_LINEAR_INTEL] = "LAYOUT_LINEAR_INTEL"
	enumLookup[PIXEL_MAP_I_TO_I] = "PIXEL_MAP_I_TO_I"
	enumLookup[MAP2_NORMAL] = "MAP2_NORMAL"
	enumLookup[BINORMAL_ARRAY_TYPE_EXT] = "BINORMAL_ARRAY_TYPE_EXT"
	enumLookup[DOT_PRODUCT_TEXTURE_3D_NV] = "DOT_PRODUCT_TEXTURE_3D_NV"
	enumLookup[COVERAGE_ATTACHMENT_NV] = "COVERAGE_ATTACHMENT_NV"
	enumLookup[XOR] = "XOR"
	enumLookup[SAMPLE_ALPHA_TO_ONE_ARB] = "SAMPLE_ALPHA_TO_ONE_ARB"
	enumLookup[MAX_SAMPLES_EXT] = "MAX_SAMPLES_EXT"
	enumLookup[FLOAT_RGBA32_NV] = "FLOAT_RGBA32_NV"
	enumLookup[PATH_STROKE_MASK_NV] = "PATH_STROKE_MASK_NV"
	enumLookup[GLYPH_HORIZONTAL_BEARING_Y_BIT_NV] = "GLYPH_HORIZONTAL_BEARING_Y_BIT_NV"
	enumLookup[TEXTURE_PRE_SPECULAR_HP] = "TEXTURE_PRE_SPECULAR_HP"
	enumLookup[DECODE_EXT] = "DECODE_EXT"
	enumLookup[RGB8_SNORM] = "RGB8_SNORM"
	enumLookup[MAX_TESS_EVALUATION_ATOMIC_COUNTERS] = "MAX_TESS_EVALUATION_ATOMIC_COUNTERS"
	enumLookup[SGIS_detail_texture] = "SGIS_detail_texture"
	enumLookup[R16] = "R16"
	enumLookup[BUMP_NUM_TEX_UNITS_ATI] = "BUMP_NUM_TEX_UNITS_ATI"
	enumLookup[UNSIGNED_INT_SAMPLER_3D] = "UNSIGNED_INT_SAMPLER_3D"
	enumLookup[UNPACK_ROW_LENGTH] = "UNPACK_ROW_LENGTH"
	enumLookup[VERTEX_WEIGHT_ARRAY_STRIDE_EXT] = "VERTEX_WEIGHT_ARRAY_STRIDE_EXT"
	enumLookup[RASTERIZER_DISCARD_NV] = "RASTERIZER_DISCARD_NV"
	enumLookup[SAMPLER_1D_ARRAY] = "SAMPLER_1D_ARRAY"
	enumLookup[MAX_SPARSE_TEXTURE_SIZE_AMD] = "MAX_SPARSE_TEXTURE_SIZE_AMD"
	enumLookup[HISTOGRAM_BLUE_SIZE] = "HISTOGRAM_BLUE_SIZE"
	enumLookup[TEXTURE17_ARB] = "TEXTURE17_ARB"
	enumLookup[PASS_THROUGH_NV] = "PASS_THROUGH_NV"
	enumLookup[RGB_FLOAT32_ATI] = "RGB_FLOAT32_ATI"
	enumLookup[MAX_VERTEX_UNIFORM_COMPONENTS] = "MAX_VERTEX_UNIFORM_COMPONENTS"
	enumLookup[ACTIVE_ATTRIBUTE_MAX_LENGTH] = "ACTIVE_ATTRIBUTE_MAX_LENGTH"
	enumLookup[INT_SAMPLER_1D_EXT] = "INT_SAMPLER_1D_EXT"
	enumLookup[ALPHA_BIAS] = "ALPHA_BIAS"
	enumLookup[INDEX_MATERIAL_PARAMETER_EXT] = "INDEX_MATERIAL_PARAMETER_EXT"
	enumLookup[VERTEX_WEIGHT_ARRAY_SIZE_EXT] = "VERTEX_WEIGHT_ARRAY_SIZE_EXT"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_X_OES] = "TEXTURE_CUBE_MAP_NEGATIVE_X_OES"
	enumLookup[IDENTITY_NV] = "IDENTITY_NV"
	enumLookup[CLAMP_READ_COLOR] = "CLAMP_READ_COLOR"
	enumLookup[POST_IMAGE_TRANSFORM_COLOR_TABLE_HP] = "POST_IMAGE_TRANSFORM_COLOR_TABLE_HP"
	enumLookup[VERTEX_PRECLIP_HINT_SGIX] = "VERTEX_PRECLIP_HINT_SGIX"
	enumLookup[OFFSET_HILO_TEXTURE_2D_NV] = "OFFSET_HILO_TEXTURE_2D_NV"
	enumLookup[MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB] = "MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB"
	enumLookup[MAP2_COLOR_4] = "MAP2_COLOR_4"
	enumLookup[PROXY_TEXTURE_4D_SGIS] = "PROXY_TEXTURE_4D_SGIS"
	enumLookup[PROGRAM_POINT_SIZE_ARB] = "PROGRAM_POINT_SIZE_ARB"
	enumLookup[DEPENDENT_AR_TEXTURE_2D_NV] = "DEPENDENT_AR_TEXTURE_2D_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_BARRIER_BIT] = "VERTEX_ATTRIB_ARRAY_BARRIER_BIT"
	enumLookup[C4UB_V3F] = "C4UB_V3F"
	enumLookup[POST_CONVOLUTION_BLUE_BIAS] = "POST_CONVOLUTION_BLUE_BIAS"
	enumLookup[CONVOLUTION_HINT_SGIX] = "CONVOLUTION_HINT_SGIX"
	enumLookup[RENDERBUFFER_HEIGHT_OES] = "RENDERBUFFER_HEIGHT_OES"
	enumLookup[RGB16UI] = "RGB16UI"
	enumLookup[COMPRESSED_RGBA_ASTC_5x5_KHR] = "COMPRESSED_RGBA_ASTC_5x5_KHR"
	enumLookup[MAX_ARRAY_TEXTURE_LAYERS] = "MAX_ARRAY_TEXTURE_LAYERS"
	enumLookup[VERTEX_ATTRIB_ARRAY_LENGTH_NV] = "VERTEX_ATTRIB_ARRAY_LENGTH_NV"
	enumLookup[LIGHT3] = "LIGHT3"
	enumLookup[MATRIX_EXT] = "MATRIX_EXT"
	enumLookup[VERTEX_ATTRIB_MAP1_SIZE_APPLE] = "VERTEX_ATTRIB_MAP1_SIZE_APPLE"
	enumLookup[TEXTURE_SHARED_SIZE] = "TEXTURE_SHARED_SIZE"
	enumLookup[FRAMEBUFFER_UNSUPPORTED_EXT] = "FRAMEBUFFER_UNSUPPORTED_EXT"
	enumLookup[LIST_BASE] = "LIST_BASE"
	enumLookup[RGB16] = "RGB16"
	enumLookup[CLAMP_TO_BORDER_ARB] = "CLAMP_TO_BORDER_ARB"
	enumLookup[MAP1_VERTEX_ATTRIB12_4_NV] = "MAP1_VERTEX_ATTRIB12_4_NV"
	enumLookup[SATURATE_BIT_ATI] = "SATURATE_BIT_ATI"
	enumLookup[MAX_TEXTURE_BUFFER_SIZE_EXT] = "MAX_TEXTURE_BUFFER_SIZE_EXT"
	enumLookup[INTENSITY8_EXT] = "INTENSITY8_EXT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE] = "FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE"
	enumLookup[PROGRAM] = "PROGRAM"
	enumLookup[HALF_BIT_ATI] = "HALF_BIT_ATI"
	enumLookup[DSDT_MAG_INTENSITY_NV] = "DSDT_MAG_INTENSITY_NV"
	enumLookup[PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB] = "PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB"
	enumLookup[ARRAY_BUFFER_BINDING] = "ARRAY_BUFFER_BINDING"
	enumLookup[FRAMEBUFFER_COMPLETE_EXT] = "FRAMEBUFFER_COMPLETE_EXT"
	enumLookup[LUMINANCE8UI_EXT] = "LUMINANCE8UI_EXT"
	enumLookup[CLIP_PLANE2] = "CLIP_PLANE2"
	enumLookup[BLEND_DST_RGB_EXT] = "BLEND_DST_RGB_EXT"
	enumLookup[TRANSPOSE_MODELVIEW_MATRIX] = "TRANSPOSE_MODELVIEW_MATRIX"
	enumLookup[COMBINER_MAPPING_NV] = "COMBINER_MAPPING_NV"
	enumLookup[OP_ROUND_EXT] = "OP_ROUND_EXT"
	enumLookup[PALETTE4_RGBA8_OES] = "PALETTE4_RGBA8_OES"
	enumLookup[QUERY_NO_WAIT] = "QUERY_NO_WAIT"
	enumLookup[INT_IMAGE_2D_MULTISAMPLE_ARRAY] = "INT_IMAGE_2D_MULTISAMPLE_ARRAY"
	enumLookup[DEBUG_CATEGORY_API_ERROR_AMD] = "DEBUG_CATEGORY_API_ERROR_AMD"
	enumLookup[PACK_LSB_FIRST] = "PACK_LSB_FIRST"
	enumLookup[MULTISAMPLE_ARB] = "MULTISAMPLE_ARB"
	enumLookup[LOSE_CONTEXT_ON_RESET_ARB] = "LOSE_CONTEXT_ON_RESET_ARB"
	enumLookup[VBO_FREE_MEMORY_ATI] = "VBO_FREE_MEMORY_ATI"
	enumLookup[DRAW_BUFFER8_NV] = "DRAW_BUFFER8_NV"
	enumLookup[BOOL_ARB] = "BOOL_ARB"
	enumLookup[TEXTURE_BINDING_1D_ARRAY] = "TEXTURE_BINDING_1D_ARRAY"
	enumLookup[SGIX_framezoom] = "SGIX_framezoom"
	enumLookup[DEPTH_COMPONENT16] = "DEPTH_COMPONENT16"
	enumLookup[DEPTH_ATTACHMENT_EXT] = "DEPTH_ATTACHMENT_EXT"
	enumLookup[GEOMETRY_SHADER_INVOCATIONS] = "GEOMETRY_SHADER_INVOCATIONS"
	enumLookup[TEXTURE_SHADER_NV] = "TEXTURE_SHADER_NV"
	enumLookup[REG_14_ATI] = "REG_14_ATI"
	enumLookup[MAP2_GRID_DOMAIN] = "MAP2_GRID_DOMAIN"
	enumLookup[GEOMETRY_DEFORMATION_SGIX] = "GEOMETRY_DEFORMATION_SGIX"
	enumLookup[SECONDARY_COLOR_ARRAY] = "SECONDARY_COLOR_ARRAY"
	enumLookup[RECIP_ADD_SIGNED_ALPHA_IMG] = "RECIP_ADD_SIGNED_ALPHA_IMG"
	enumLookup[TEXTURE1] = "TEXTURE1"
	enumLookup[SAMPLER_2D_SHADOW_EXT] = "SAMPLER_2D_SHADOW_EXT"
	enumLookup[INFO_LOG_LENGTH] = "INFO_LOG_LENGTH"
	enumLookup[BINORMAL_ARRAY_STRIDE_EXT] = "BINORMAL_ARRAY_STRIDE_EXT"
	enumLookup[PROGRAM_NATIVE_TEMPORARIES_ARB] = "PROGRAM_NATIVE_TEMPORARIES_ARB"
	enumLookup[MATRIX8_ARB] = "MATRIX8_ARB"
	enumLookup[UNSIGNALED] = "UNSIGNALED"
	enumLookup[MAX_CONVOLUTION_WIDTH] = "MAX_CONVOLUTION_WIDTH"
	enumLookup[MULTISAMPLE] = "MULTISAMPLE"
	enumLookup[TEXTURE_POST_SPECULAR_HP] = "TEXTURE_POST_SPECULAR_HP"
	enumLookup[SRGB_ALPHA_EXT] = "SRGB_ALPHA_EXT"
	enumLookup[MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS] = "MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS"
	enumLookup[MAX_VERTEX_ATOMIC_COUNTERS] = "MAX_VERTEX_ATOMIC_COUNTERS"
	enumLookup[BACK_PRIMARY_COLOR_NV] = "BACK_PRIMARY_COLOR_NV"
	enumLookup[BUMP_ROT_MATRIX_ATI] = "BUMP_ROT_MATRIX_ATI"
	enumLookup[TEXTURE_SHARED_SIZE_EXT] = "TEXTURE_SHARED_SIZE_EXT"
	enumLookup[MAX_COMBINED_ATOMIC_COUNTER_BUFFERS] = "MAX_COMBINED_ATOMIC_COUNTER_BUFFERS"
	enumLookup[MATRIX5_ARB] = "MATRIX5_ARB"
	enumLookup[MAX_DEBUG_LOGGED_MESSAGES] = "MAX_DEBUG_LOGGED_MESSAGES"
	enumLookup[MAX_FRAMEBUFFER_LAYERS] = "MAX_FRAMEBUFFER_LAYERS"
	enumLookup[ALL_ATTRIB_BITS] = "ALL_ATTRIB_BITS"
	enumLookup[CONSTANT_ARB] = "CONSTANT_ARB"
	enumLookup[STENCIL_BACK_PASS_DEPTH_FAIL_ATI] = "STENCIL_BACK_PASS_DEPTH_FAIL_ATI"
	enumLookup[TEXTURE_DEPTH_QCOM] = "TEXTURE_DEPTH_QCOM"
	enumLookup[PATH_STROKE_WIDTH_NV] = "PATH_STROKE_WIDTH_NV"
	enumLookup[POST_COLOR_MATRIX_RED_BIAS] = "POST_COLOR_MATRIX_RED_BIAS"
	enumLookup[STENCIL_BACK_FUNC_ATI] = "STENCIL_BACK_FUNC_ATI"
	enumLookup[HILO8_NV] = "HILO8_NV"
	enumLookup[TESS_CONTROL_OUTPUT_VERTICES] = "TESS_CONTROL_OUTPUT_VERTICES"
	enumLookup[UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY] = "UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY"
	enumLookup[GL_4PASS_3_SGIS] = "GL_4PASS_3_SGIS"
	enumLookup[ACTIVE_SUBROUTINE_UNIFORMS] = "ACTIVE_SUBROUTINE_UNIFORMS"
	enumLookup[LINE_RESET_TOKEN] = "LINE_RESET_TOKEN"
	enumLookup[DRAW_BUFFER_EXT] = "DRAW_BUFFER_EXT"
	enumLookup[PACK_SKIP_VOLUMES_SGIS] = "PACK_SKIP_VOLUMES_SGIS"
	enumLookup[IUI_V3F_EXT] = "IUI_V3F_EXT"
	enumLookup[FRAMEBUFFER_RENDERABLE] = "FRAMEBUFFER_RENDERABLE"
	enumLookup[BUMP_ROT_MATRIX_SIZE_ATI] = "BUMP_ROT_MATRIX_SIZE_ATI"
	enumLookup[MAX_CUBE_MAP_TEXTURE_SIZE_EXT] = "MAX_CUBE_MAP_TEXTURE_SIZE_EXT"
	enumLookup[WRITE_PIXEL_DATA_RANGE_POINTER_NV] = "WRITE_PIXEL_DATA_RANGE_POINTER_NV"
	enumLookup[RGBA_INTEGER_MODE_EXT] = "RGBA_INTEGER_MODE_EXT"
	enumLookup[NORMAL_ARRAY_LENGTH_NV] = "NORMAL_ARRAY_LENGTH_NV"
	enumLookup[INT16_VEC3_NV] = "INT16_VEC3_NV"
	enumLookup[CONVOLUTION_HEIGHT_EXT] = "CONVOLUTION_HEIGHT_EXT"
	enumLookup[STATIC_DRAW_ARB] = "STATIC_DRAW_ARB"
	enumLookup[FRAMEBUFFER_UNSUPPORTED] = "FRAMEBUFFER_UNSUPPORTED"
	enumLookup[TEXTURE_EXTERNAL_OES] = "TEXTURE_EXTERNAL_OES"
	enumLookup[MAP_STENCIL] = "MAP_STENCIL"
	enumLookup[DETAIL_TEXTURE_FUNC_POINTS_SGIS] = "DETAIL_TEXTURE_FUNC_POINTS_SGIS"
	enumLookup[COLOR_TABLE_BLUE_SIZE] = "COLOR_TABLE_BLUE_SIZE"
	enumLookup[CURRENT_QUERY] = "CURRENT_QUERY"
	enumLookup[LUMINANCE_ALPHA_FLOAT16_APPLE] = "LUMINANCE_ALPHA_FLOAT16_APPLE"
	enumLookup[CON_31_ATI] = "CON_31_ATI"
	enumLookup[FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES] = "FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES"
	enumLookup[UNSIGNED_INT_IMAGE_CUBE] = "UNSIGNED_INT_IMAGE_CUBE"
	enumLookup[SYSTEM_FONT_NAME_NV] = "SYSTEM_FONT_NAME_NV"
	enumLookup[PERSPECTIVE_CORRECTION_HINT] = "PERSPECTIVE_CORRECTION_HINT"
	enumLookup[TEXTURE_COORD_ARRAY_STRIDE_EXT] = "TEXTURE_COORD_ARRAY_STRIDE_EXT"
	enumLookup[POST_COLOR_MATRIX_BLUE_SCALE_SGI] = "POST_COLOR_MATRIX_BLUE_SCALE_SGI"
	enumLookup[MAX_UNIFORM_LOCATIONS] = "MAX_UNIFORM_LOCATIONS"
	enumLookup[SOURCE2_ALPHA_ARB] = "SOURCE2_ALPHA_ARB"
	enumLookup[COLOR_BUFFER_BIT] = "COLOR_BUFFER_BIT"
	enumLookup[REPLACE_EXT] = "REPLACE_EXT"
	enumLookup[REFLECTION_MAP_OES] = "REFLECTION_MAP_OES"
	enumLookup[SOURCE1_ALPHA_EXT] = "SOURCE1_ALPHA_EXT"
	enumLookup[READ_FRAMEBUFFER_NV] = "READ_FRAMEBUFFER_NV"
	enumLookup[DEPTH_WRITEMASK] = "DEPTH_WRITEMASK"
	enumLookup[INDEX_ARRAY_STRIDE] = "INDEX_ARRAY_STRIDE"
	enumLookup[TEXTURE_GATHER] = "TEXTURE_GATHER"
	enumLookup[TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB] = "TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB"
	enumLookup[OUTPUT_TEXTURE_COORD31_EXT] = "OUTPUT_TEXTURE_COORD31_EXT"
	enumLookup[STENCIL_ATTACHMENT_EXT] = "STENCIL_ATTACHMENT_EXT"
	enumLookup[DEBUG_OUTPUT_SYNCHRONOUS] = "DEBUG_OUTPUT_SYNCHRONOUS"
	enumLookup[NUM_PROGRAM_BINARY_FORMATS_OES] = "NUM_PROGRAM_BINARY_FORMATS_OES"
	enumLookup[CON_2_ATI] = "CON_2_ATI"
	enumLookup[TEXTURE_SWIZZLE_RGBA] = "TEXTURE_SWIZZLE_RGBA"
	enumLookup[STANDARD_FONT_NAME_NV] = "STANDARD_FONT_NAME_NV"
	enumLookup[MAX_PROGRAM_TEX_INDIRECTIONS_ARB] = "MAX_PROGRAM_TEX_INDIRECTIONS_ARB"
	enumLookup[PATH_DASH_OFFSET_RESET_NV] = "PATH_DASH_OFFSET_RESET_NV"
	enumLookup[INTENSITY16UI_EXT] = "INTENSITY16UI_EXT"
	enumLookup[HISTOGRAM_BLUE_SIZE_EXT] = "HISTOGRAM_BLUE_SIZE_EXT"
	enumLookup[FIRST_VERTEX_CONVENTION] = "FIRST_VERTEX_CONVENTION"
	enumLookup[GL_3D_COLOR] = "GL_3D_COLOR"
	enumLookup[TEXTURE_COMPONENTS] = "TEXTURE_COMPONENTS"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_X_EXT] = "TEXTURE_CUBE_MAP_POSITIVE_X_EXT"
	enumLookup[MAX_SAMPLES_ANGLE] = "MAX_SAMPLES_ANGLE"
	enumLookup[VERTEX_ARRAY_POINTER_EXT] = "VERTEX_ARRAY_POINTER_EXT"
	enumLookup[MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT] = "MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT"
	enumLookup[MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT] = "MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT"
	enumLookup[FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE] = "FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE"
	enumLookup[WEIGHT_ARRAY_TYPE_OES] = "WEIGHT_ARRAY_TYPE_OES"
	enumLookup[OUTPUT_COLOR0_EXT] = "OUTPUT_COLOR0_EXT"
	enumLookup[LUMINANCE16F_ARB] = "LUMINANCE16F_ARB"
	enumLookup[SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB] = "SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[POLYGON_STIPPLE_BIT] = "POLYGON_STIPPLE_BIT"
	enumLookup[BINORMAL_ARRAY_EXT] = "BINORMAL_ARRAY_EXT"
	enumLookup[ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER] = "ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER"
	enumLookup[INDEX_WRITEMASK] = "INDEX_WRITEMASK"
	enumLookup[COPY] = "COPY"
	enumLookup[RGB12] = "RGB12"
	enumLookup[DRAW_BUFFER8_ATI] = "DRAW_BUFFER8_ATI"
	enumLookup[MAX_FRAGMENT_UNIFORM_BLOCKS] = "MAX_FRAGMENT_UNIFORM_BLOCKS"
	enumLookup[PROGRAM_OBJECT_EXT] = "PROGRAM_OBJECT_EXT"
	enumLookup[PARTIAL_SUCCESS_NV] = "PARTIAL_SUCCESS_NV"
	enumLookup[CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB] = "CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB"
	enumLookup[BUFFER_MAP_POINTER_ARB] = "BUFFER_MAP_POINTER_ARB"
	enumLookup[SGIX_texture_coordinate_clamp] = "SGIX_texture_coordinate_clamp"
	enumLookup[R8] = "R8"
	enumLookup[TEXTURE6] = "TEXTURE6"
	enumLookup[COMBINER_AB_DOT_PRODUCT_NV] = "COMBINER_AB_DOT_PRODUCT_NV"
	enumLookup[EDGE_FLAG_ARRAY_EXT] = "EDGE_FLAG_ARRAY_EXT"
	enumLookup[CULL_VERTEX_OBJECT_POSITION_EXT] = "CULL_VERTEX_OBJECT_POSITION_EXT"
	enumLookup[NUM_COMPRESSED_TEXTURE_FORMATS_ARB] = "NUM_COMPRESSED_TEXTURE_FORMATS_ARB"
	enumLookup[QUAD_MESH_SUN] = "QUAD_MESH_SUN"
	enumLookup[DRAW_BUFFER5_ATI] = "DRAW_BUFFER5_ATI"
	enumLookup[STREAM_COPY] = "STREAM_COPY"
	enumLookup[PIXEL_BUFFER_BARRIER_BIT] = "PIXEL_BUFFER_BARRIER_BIT"
	enumLookup[SRC0_ALPHA] = "SRC0_ALPHA"
	enumLookup[TEXTURE_TARGET_QCOM] = "TEXTURE_TARGET_QCOM"
	enumLookup[RENDERBUFFER_STENCIL_SIZE] = "RENDERBUFFER_STENCIL_SIZE"
	enumLookup[MULTISAMPLE_COVERAGE_MODES_NV] = "MULTISAMPLE_COVERAGE_MODES_NV"
	enumLookup[COMPRESSED_RGBA_ASTC_12x12_KHR] = "COMPRESSED_RGBA_ASTC_12x12_KHR"
	enumLookup[MAP_INVALIDATE_BUFFER_BIT_EXT] = "MAP_INVALIDATE_BUFFER_BIT_EXT"
	enumLookup[MAX_ATTRIB_STACK_DEPTH] = "MAX_ATTRIB_STACK_DEPTH"
	enumLookup[MAP2_VERTEX_4] = "MAP2_VERTEX_4"
	enumLookup[UNSIGNED_INT_10_10_10_2] = "UNSIGNED_INT_10_10_10_2"
	enumLookup[FRAGMENT_MATERIAL_EXT] = "FRAGMENT_MATERIAL_EXT"
	enumLookup[TEXTURE11_ARB] = "TEXTURE11_ARB"
	enumLookup[HILO16_NV] = "HILO16_NV"
	enumLookup[COLOR_ATTACHMENT11_NV] = "COLOR_ATTACHMENT11_NV"
	enumLookup[MAX_IMAGE_SAMPLES] = "MAX_IMAGE_SAMPLES"
	enumLookup[PIXEL_TRANSFORM_2D_MATRIX_EXT] = "PIXEL_TRANSFORM_2D_MATRIX_EXT"
	enumLookup[DOUBLE_MAT4_EXT] = "DOUBLE_MAT4_EXT"
	enumLookup[INT_IMAGE_CUBE_MAP_ARRAY_EXT] = "INT_IMAGE_CUBE_MAP_ARRAY_EXT"
	enumLookup[QUERY_RESULT_NO_WAIT_AMD] = "QUERY_RESULT_NO_WAIT_AMD"
	enumLookup[HINT_BIT] = "HINT_BIT"
	enumLookup[DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB] = "DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB"
	enumLookup[FRAGMENT_COLOR_MATERIAL_FACE_SGIX] = "FRAGMENT_COLOR_MATERIAL_FACE_SGIX"
	enumLookup[TEXTURE_CUBE_MAP_POSITIVE_X_ARB] = "TEXTURE_CUBE_MAP_POSITIVE_X_ARB"
	enumLookup[COLOR_MATERIAL] = "COLOR_MATERIAL"
	enumLookup[FRAGMENT_LIGHT1_SGIX] = "FRAGMENT_LIGHT1_SGIX"
	enumLookup[MAX_ATOMIC_COUNTER_BUFFER_BINDINGS] = "MAX_ATOMIC_COUNTER_BUFFER_BINDINGS"
	enumLookup[FOG_SPECULAR_TEXTURE_WIN] = "FOG_SPECULAR_TEXTURE_WIN"
	enumLookup[TEXTURE7_ARB] = "TEXTURE7_ARB"
	enumLookup[IMAGE_BINDING_ACCESS] = "IMAGE_BINDING_ACCESS"
	enumLookup[RGBA8_SNORM] = "RGBA8_SNORM"
	enumLookup[CONSTANT_ALPHA_EXT] = "CONSTANT_ALPHA_EXT"
	enumLookup[COLOR_INDEX4_EXT] = "COLOR_INDEX4_EXT"
	enumLookup[TEXTURE_DEFORMATION_SGIX] = "TEXTURE_DEFORMATION_SGIX"
	enumLookup[PALETTE4_RGB8_OES] = "PALETTE4_RGB8_OES"
	enumLookup[COLOR_ATTACHMENT13_EXT] = "COLOR_ATTACHMENT13_EXT"
	enumLookup[STENCIL_INDEX8] = "STENCIL_INDEX8"
	enumLookup[TEXTURE_2D_MULTISAMPLE_ARRAY] = "TEXTURE_2D_MULTISAMPLE_ARRAY"
	enumLookup[PROGRAM_BINARY_ANGLE] = "PROGRAM_BINARY_ANGLE"
	enumLookup[TEXTURE_RED_SIZE_EXT] = "TEXTURE_RED_SIZE_EXT"
	enumLookup[NORMAL_ARRAY_TYPE_EXT] = "NORMAL_ARRAY_TYPE_EXT"
	enumLookup[COLOR_MATRIX_STACK_DEPTH_SGI] = "COLOR_MATRIX_STACK_DEPTH_SGI"
	enumLookup[TEXT_FRAGMENT_SHADER_ATI] = "TEXT_FRAGMENT_SHADER_ATI"
	enumLookup[UNSIGNED_INT_8_8_S8_S8_REV_NV] = "UNSIGNED_INT_8_8_S8_S8_REV_NV"
	enumLookup[NORMAL_BIT_PGI] = "NORMAL_BIT_PGI"
	enumLookup[UNPACK_SWAP_BYTES] = "UNPACK_SWAP_BYTES"
	enumLookup[TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX] = "TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX"
	enumLookup[MAX_TEXTURE_COORDS_ARB] = "MAX_TEXTURE_COORDS_ARB"
	enumLookup[CLIP_DISTANCE_NV] = "CLIP_DISTANCE_NV"
	enumLookup[MAX_DEPTH_TEXTURE_SAMPLES] = "MAX_DEPTH_TEXTURE_SAMPLES"
	enumLookup[VERTEX_SUBROUTINE_UNIFORM] = "VERTEX_SUBROUTINE_UNIFORM"
	enumLookup[ALPHA16] = "ALPHA16"
	enumLookup[FENCE_CONDITION_NV] = "FENCE_CONDITION_NV"
	enumLookup[REG_1_ATI] = "REG_1_ATI"
	enumLookup[INTERLACE_OML] = "INTERLACE_OML"
	enumLookup[COLOR_ATTACHMENT1_NV] = "COLOR_ATTACHMENT1_NV"
	enumLookup[INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT] = "INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT"
	enumLookup[MAT_AMBIENT_BIT_PGI] = "MAT_AMBIENT_BIT_PGI"
	enumLookup[SGIX_flush_raster] = "SGIX_flush_raster"
	enumLookup[LINES_ADJACENCY_EXT] = "LINES_ADJACENCY_EXT"
	enumLookup[SHADE_MODEL] = "SHADE_MODEL"
	enumLookup[COLOR_INDEX1_EXT] = "COLOR_INDEX1_EXT"
	enumLookup[DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX] = "DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX"
	enumLookup[OPERAND2_ALPHA_ARB] = "OPERAND2_ALPHA_ARB"
	enumLookup[FLOAT_RGBA_MODE_NV] = "FLOAT_RGBA_MODE_NV"
	enumLookup[DYNAMIC_COPY_ARB] = "DYNAMIC_COPY_ARB"
	enumLookup[OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB] = "OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB"
	enumLookup[MAX_MULTIVIEW_BUFFERS_EXT] = "MAX_MULTIVIEW_BUFFERS_EXT"
	enumLookup[SYNC_FLAGS] = "SYNC_FLAGS"
	enumLookup[MAX_TESS_CONTROL_ATOMIC_COUNTERS] = "MAX_TESS_CONTROL_ATOMIC_COUNTERS"
	enumLookup[CLAMP_TO_BORDER_NV] = "CLAMP_TO_BORDER_NV"
	enumLookup[COLOR_ATTACHMENT13] = "COLOR_ATTACHMENT13"
	enumLookup[MAX_FRAGMENT_UNIFORM_VECTORS] = "MAX_FRAGMENT_UNIFORM_VECTORS"
	enumLookup[ELEMENT_ARRAY_UNIFIED_NV] = "ELEMENT_ARRAY_UNIFIED_NV"
	enumLookup[PATH_DASH_OFFSET_NV] = "PATH_DASH_OFFSET_NV"
	enumLookup[TEXTURE_IMMUTABLE_FORMAT] = "TEXTURE_IMMUTABLE_FORMAT"
	enumLookup[VERTEX23_BIT_PGI] = "VERTEX23_BIT_PGI"
	enumLookup[MAX_PROGRAM_ADDRESS_REGISTERS_ARB] = "MAX_PROGRAM_ADDRESS_REGISTERS_ARB"
	enumLookup[TEXTURE_BLUE_TYPE_ARB] = "TEXTURE_BLUE_TYPE_ARB"
	enumLookup[GREEN_SCALE] = "GREEN_SCALE"
	enumLookup[UNSIGNED_SHORT_5_6_5] = "UNSIGNED_SHORT_5_6_5"
	enumLookup[PREVIOUS_TEXTURE_INPUT_NV] = "PREVIOUS_TEXTURE_INPUT_NV"
	enumLookup[TEXTURE_HI_SIZE_NV] = "TEXTURE_HI_SIZE_NV"
	enumLookup[ONE_EXT] = "ONE_EXT"
	enumLookup[BUFFER_SERIALIZED_MODIFY_APPLE] = "BUFFER_SERIALIZED_MODIFY_APPLE"
	enumLookup[FRAMEBUFFER_INCOMPLETE_ATTACHMENT] = "FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	enumLookup[FIXED_OES] = "FIXED_OES"
	enumLookup[DEBUG_CALLBACK_USER_PARAM_ARB] = "DEBUG_CALLBACK_USER_PARAM_ARB"
	enumLookup[COMPRESSED_RG11_EAC] = "COMPRESSED_RG11_EAC"
	enumLookup[RGBA12] = "RGBA12"
	enumLookup[SRGB_EXT] = "SRGB_EXT"
	enumLookup[MAX_GEOMETRY_VARYING_COMPONENTS_ARB] = "MAX_GEOMETRY_VARYING_COMPONENTS_ARB"
	enumLookup[TEXTURE_SWIZZLE_B] = "TEXTURE_SWIZZLE_B"
	enumLookup[POST_TEXTURE_FILTER_SCALE_RANGE_SGIX] = "POST_TEXTURE_FILTER_SCALE_RANGE_SGIX"
	enumLookup[NUM_SHADING_LANGUAGE_VERSIONS] = "NUM_SHADING_LANGUAGE_VERSIONS"
	enumLookup[MAP_ATTRIB_U_ORDER_NV] = "MAP_ATTRIB_U_ORDER_NV"
	enumLookup[TEXTURE_SWIZZLE_R_EXT] = "TEXTURE_SWIZZLE_R_EXT"
	enumLookup[BLEND_COLOR] = "BLEND_COLOR"
	enumLookup[PN_TRIANGLES_ATI] = "PN_TRIANGLES_ATI"
	enumLookup[TEXTURE_INTENSITY_TYPE] = "TEXTURE_INTENSITY_TYPE"
	enumLookup[INT_SAMPLER_3D_EXT] = "INT_SAMPLER_3D_EXT"
	enumLookup[BUFFER_OBJECT_EXT] = "BUFFER_OBJECT_EXT"
	enumLookup[UNSIGNED_SHORT_5_6_5_REV_EXT] = "UNSIGNED_SHORT_5_6_5_REV_EXT"
	enumLookup[INT_SAMPLER_RENDERBUFFER_NV] = "INT_SAMPLER_RENDERBUFFER_NV"
	enumLookup[RGBA16_EXT] = "RGBA16_EXT"
	enumLookup[COMBINER_COMPONENT_USAGE_NV] = "COMBINER_COMPONENT_USAGE_NV"
	enumLookup[SAMPLES_3DFX] = "SAMPLES_3DFX"
	enumLookup[DRAW_BUFFER7_NV] = "DRAW_BUFFER7_NV"
	enumLookup[FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT] = "FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT"
	enumLookup[NO_ERROR] = "NO_ERROR"
	enumLookup[INTENSITY32I_EXT] = "INTENSITY32I_EXT"
	enumLookup[TRIANGLES_ADJACENCY_EXT] = "TRIANGLES_ADJACENCY_EXT"
	enumLookup[CLAMP_TO_EDGE] = "CLAMP_TO_EDGE"
	enumLookup[TEXTURE13_ARB] = "TEXTURE13_ARB"
	enumLookup[MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV] = "MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV"
	enumLookup[COLOR_ATTACHMENT3_NV] = "COLOR_ATTACHMENT3_NV"
	enumLookup[LINE_TOKEN] = "LINE_TOKEN"
	enumLookup[MAX_FRAMEZOOM_FACTOR_SGIX] = "MAX_FRAMEZOOM_FACTOR_SGIX"
	enumLookup[PER_STAGE_CONSTANTS_NV] = "PER_STAGE_CONSTANTS_NV"
	enumLookup[SAMPLER_2D_ARB] = "SAMPLER_2D_ARB"
	enumLookup[SLUMINANCE_ALPHA_EXT] = "SLUMINANCE_ALPHA_EXT"
	enumLookup[FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES] = "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES"
	enumLookup[DEBUG_SOURCE_SHADER_COMPILER] = "DEBUG_SOURCE_SHADER_COMPILER"
	enumLookup[FRAGMENT_DEPTH_EXT] = "FRAGMENT_DEPTH_EXT"
	enumLookup[UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES] = "UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES"
	enumLookup[MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT] = "MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT"
	enumLookup[RGBA16UI_EXT] = "RGBA16UI_EXT"
	enumLookup[COMPUTE_SHADER_BIT] = "COMPUTE_SHADER_BIT"
	enumLookup[PIXEL_MAP_R_TO_R] = "PIXEL_MAP_R_TO_R"
	enumLookup[UNPACK_SKIP_PIXELS] = "UNPACK_SKIP_PIXELS"
	enumLookup[DONT_CARE] = "DONT_CARE"
	enumLookup[LUMINANCE12_ALPHA12_EXT] = "LUMINANCE12_ALPHA12_EXT"
	enumLookup[IMAGE_CUBIC_WEIGHT_HP] = "IMAGE_CUBIC_WEIGHT_HP"
	enumLookup[FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX] = "FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX"
	enumLookup[MAX_CUBE_MAP_TEXTURE_SIZE_ARB] = "MAX_CUBE_MAP_TEXTURE_SIZE_ARB"
	enumLookup[MAP1_VERTEX_ATTRIB11_4_NV] = "MAP1_VERTEX_ATTRIB11_4_NV"
	enumLookup[DRAW_BUFFER1_NV] = "DRAW_BUFFER1_NV"
	enumLookup[NORMAL_ARRAY_ADDRESS_NV] = "NORMAL_ARRAY_ADDRESS_NV"
	enumLookup[GLYPH_VERTICAL_BEARING_X_BIT_NV] = "GLYPH_VERTICAL_BEARING_X_BIT_NV"
	enumLookup[MAX_COLOR_TEXTURE_SAMPLES] = "MAX_COLOR_TEXTURE_SAMPLES"
	enumLookup[SGIX_texture_lod_bias] = "SGIX_texture_lod_bias"
	enumLookup[UNSIGNED_INT_8_8_8_8] = "UNSIGNED_INT_8_8_8_8"
	enumLookup[COLOR_TABLE_INTENSITY_SIZE_SGI] = "COLOR_TABLE_INTENSITY_SIZE_SGI"
	enumLookup[FRAGMENT_NORMAL_EXT] = "FRAGMENT_NORMAL_EXT"
	enumLookup[CLIENT_ACTIVE_TEXTURE] = "CLIENT_ACTIVE_TEXTURE"
	enumLookup[INCR_WRAP] = "INCR_WRAP"
	enumLookup[VERTEX_ATTRIB_ARRAY15_NV] = "VERTEX_ATTRIB_ARRAY15_NV"
	enumLookup[VERTEX_ARRAY_ADDRESS_NV] = "VERTEX_ARRAY_ADDRESS_NV"
	enumLookup[RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV] = "RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV"
	enumLookup[ACTIVE_RESOURCES] = "ACTIVE_RESOURCES"
	enumLookup[NORMALIZE] = "NORMALIZE"
	enumLookup[OR_INVERTED] = "OR_INVERTED"
	enumLookup[CONVOLUTION_WIDTH] = "CONVOLUTION_WIDTH"
	enumLookup[TEXTURE20] = "TEXTURE20"
	enumLookup[POINT_SPRITE_NV] = "POINT_SPRITE_NV"
	enumLookup[TEXTURE_FLOAT_COMPONENTS_NV] = "TEXTURE_FLOAT_COMPONENTS_NV"
	enumLookup[MAP1_COLOR_4] = "MAP1_COLOR_4"
	enumLookup[TEXTURE_2D_BINDING_EXT] = "TEXTURE_2D_BINDING_EXT"
	enumLookup[ASYNC_READ_PIXELS_SGIX] = "ASYNC_READ_PIXELS_SGIX"
	enumLookup[COMPRESSED_RGB_S3TC_DXT1_EXT] = "COMPRESSED_RGB_S3TC_DXT1_EXT"
	enumLookup[RGB16F_ARB] = "RGB16F_ARB"
	enumLookup[REG_23_ATI] = "REG_23_ATI"
	enumLookup[ALPHA16F_ARB] = "ALPHA16F_ARB"
	enumLookup[DOT3_ATI] = "DOT3_ATI"
	enumLookup[PURGEABLE_APPLE] = "PURGEABLE_APPLE"
	enumLookup[INDEX_TEST_FUNC_EXT] = "INDEX_TEST_FUNC_EXT"
	enumLookup[TEXTURE_COMPRESSED_BLOCK_SIZE] = "TEXTURE_COMPRESSED_BLOCK_SIZE"
	enumLookup[COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV] = "COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV"
	enumLookup[ATOMIC_COUNTER_BUFFER_START] = "ATOMIC_COUNTER_BUFFER_START"
	enumLookup[POLYGON_SMOOTH] = "POLYGON_SMOOTH"
	enumLookup[VERTEX_ATTRIB_ARRAY_SIZE] = "VERTEX_ATTRIB_ARRAY_SIZE"
	enumLookup[DEPTH_BOUNDS_TEST_EXT] = "DEPTH_BOUNDS_TEST_EXT"
	enumLookup[CONTEXT_PROFILE_MASK] = "CONTEXT_PROFILE_MASK"
	enumLookup[BIAS_BY_NEGATIVE_ONE_HALF_NV] = "BIAS_BY_NEGATIVE_ONE_HALF_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY_ENABLED] = "VERTEX_ATTRIB_ARRAY_ENABLED"
	enumLookup[ALPHA_FLOAT16_ATI] = "ALPHA_FLOAT16_ATI"
	enumLookup[DRAW_BUFFER15_ARB] = "DRAW_BUFFER15_ARB"
	enumLookup[MAX_DEPTH] = "MAX_DEPTH"
	enumLookup[PRIMITIVES_GENERATED_NV] = "PRIMITIVES_GENERATED_NV"
	enumLookup[TEXTURE_VIEW_MIN_LAYER] = "TEXTURE_VIEW_MIN_LAYER"
	enumLookup[REFLECTION_MAP] = "REFLECTION_MAP"
	enumLookup[RGBA_FLOAT16_APPLE] = "RGBA_FLOAT16_APPLE"
	enumLookup[DRAW_BUFFER10] = "DRAW_BUFFER10"
	enumLookup[SLIM8U_SGIX] = "SLIM8U_SGIX"
	enumLookup[MATRIX21_ARB] = "MATRIX21_ARB"
	enumLookup[COLOR_INDEXES] = "COLOR_INDEXES"
	enumLookup[RGBA4_EXT] = "RGBA4_EXT"
	enumLookup[POST_COLOR_MATRIX_RED_BIAS_SGI] = "POST_COLOR_MATRIX_RED_BIAS_SGI"
	enumLookup[PIXEL_TILE_WIDTH_SGIX] = "PIXEL_TILE_WIDTH_SGIX"
	enumLookup[SRC1_COLOR] = "SRC1_COLOR"
	enumLookup[DEPTH32F_STENCIL8] = "DEPTH32F_STENCIL8"
	enumLookup[INT_SAMPLER_2D_ARRAY_EXT] = "INT_SAMPLER_2D_ARRAY_EXT"
	enumLookup[EXT_histogram] = "EXT_histogram"
	enumLookup[TEXTURE_INTENSITY_SIZE] = "TEXTURE_INTENSITY_SIZE"
	enumLookup[COMPILE] = "COMPILE"
	enumLookup[DEBUG_SOURCE_WINDOW_SYSTEM] = "DEBUG_SOURCE_WINDOW_SYSTEM"
	enumLookup[DEBUG_TYPE_MARKER] = "DEBUG_TYPE_MARKER"
	enumLookup[TEXTURE_COMPRESSED_IMAGE_SIZE_ARB] = "TEXTURE_COMPRESSED_IMAGE_SIZE_ARB"
	enumLookup[MAX_SUBROUTINE_UNIFORM_LOCATIONS] = "MAX_SUBROUTINE_UNIFORM_LOCATIONS"
	enumLookup[SIGNED_HILO_NV] = "SIGNED_HILO_NV"
	enumLookup[SWIZZLE_STQ_ATI] = "SWIZZLE_STQ_ATI"
	enumLookup[COLOR_ATTACHMENT14] = "COLOR_ATTACHMENT14"
	enumLookup[TRIANGLE_STRIP_ADJACENCY] = "TRIANGLE_STRIP_ADJACENCY"
	enumLookup[TEXTURE_DEFORMATION_BIT_SGIX] = "TEXTURE_DEFORMATION_BIT_SGIX"
	enumLookup[QUAD_ALPHA4_SGIS] = "QUAD_ALPHA4_SGIS"
	enumLookup[SCALE_BY_TWO_NV] = "SCALE_BY_TWO_NV"
	enumLookup[MATRIX0_ARB] = "MATRIX0_ARB"
	enumLookup[UNPACK_COMPRESSED_BLOCK_DEPTH] = "UNPACK_COMPRESSED_BLOCK_DEPTH"
	enumLookup[PROJECTION_STACK_DEPTH] = "PROJECTION_STACK_DEPTH"
	enumLookup[PACK_SKIP_ROWS] = "PACK_SKIP_ROWS"
	enumLookup[TEXTURE_COORD_ARRAY_TYPE_EXT] = "TEXTURE_COORD_ARRAY_TYPE_EXT"
	enumLookup[NO_RESET_NOTIFICATION_ARB] = "NO_RESET_NOTIFICATION_ARB"
	enumLookup[DITHER] = "DITHER"
	enumLookup[VERTEX_STREAM4_ATI] = "VERTEX_STREAM4_ATI"
	enumLookup[VERTEX_SHADER] = "VERTEX_SHADER"
	enumLookup[FACTOR_MIN_AMD] = "FACTOR_MIN_AMD"
	enumLookup[POST_COLOR_MATRIX_ALPHA_SCALE_SGI] = "POST_COLOR_MATRIX_ALPHA_SCALE_SGI"
	enumLookup[IGNORE_BORDER_HP] = "IGNORE_BORDER_HP"
	enumLookup[ATOMIC_COUNTER_BUFFER_INDEX] = "ATOMIC_COUNTER_BUFFER_INDEX"
	enumLookup[CCW] = "CCW"
	enumLookup[VIEWPORT_BOUNDS_RANGE] = "VIEWPORT_BOUNDS_RANGE"
	enumLookup[SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE] = "SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE"
	enumLookup[HI_SCALE_NV] = "HI_SCALE_NV"
	enumLookup[CON_23_ATI] = "CON_23_ATI"
	enumLookup[EVAL_BIT] = "EVAL_BIT"
	enumLookup[OP_INDEX_EXT] = "OP_INDEX_EXT"
	enumLookup[VERTEX_SHADER_INVARIANTS_EXT] = "VERTEX_SHADER_INVARIANTS_EXT"
	enumLookup[DOUBLE_VEC2_EXT] = "DOUBLE_VEC2_EXT"
	enumLookup[MULTISAMPLE_BIT_EXT] = "MULTISAMPLE_BIT_EXT"
	enumLookup[TEXTURE_MIN_LOD] = "TEXTURE_MIN_LOD"
	enumLookup[DEPTH_TEXTURE_MODE] = "DEPTH_TEXTURE_MODE"
	enumLookup[ATC_RGBA_EXPLICIT_ALPHA_AMD] = "ATC_RGBA_EXPLICIT_ALPHA_AMD"
	enumLookup[LIST_MODE] = "LIST_MODE"
	enumLookup[MAX_CUBE_MAP_TEXTURE_SIZE_OES] = "MAX_CUBE_MAP_TEXTURE_SIZE_OES"
	enumLookup[RGBA16F] = "RGBA16F"
	enumLookup[SET] = "SET"
	enumLookup[CLIP_DISTANCE4] = "CLIP_DISTANCE4"
	enumLookup[ACTIVE_PROGRAM] = "ACTIVE_PROGRAM"
	enumLookup[TEXTURE19_ARB] = "TEXTURE19_ARB"
	enumLookup[STENCIL_INDEX8_OES] = "STENCIL_INDEX8_OES"
	enumLookup[PRIMITIVE_RESTART_FIXED_INDEX] = "PRIMITIVE_RESTART_FIXED_INDEX"
	enumLookup[FRAMEBUFFER_INCOMPLETE_FORMATS_EXT] = "FRAMEBUFFER_INCOMPLETE_FORMATS_EXT"
	enumLookup[PATH_COORD_COUNT_NV] = "PATH_COORD_COUNT_NV"
	enumLookup[MAX_CLIP_PLANES] = "MAX_CLIP_PLANES"
	enumLookup[UNSIGNED_SHORT_5_5_5_1] = "UNSIGNED_SHORT_5_5_5_1"
	enumLookup[DEBUG_SOURCE_SHADER_COMPILER_ARB] = "DEBUG_SOURCE_SHADER_COMPILER_ARB"
	enumLookup[VERTEX_STATE_PROGRAM_NV] = "VERTEX_STATE_PROGRAM_NV"
	enumLookup[TEXTURE_LO_SIZE_NV] = "TEXTURE_LO_SIZE_NV"
	enumLookup[MAX_PROGRAM_PARAMETERS_ARB] = "MAX_PROGRAM_PARAMETERS_ARB"
	enumLookup[UNSIGNED_INT_IMAGE_BUFFER_EXT] = "UNSIGNED_INT_IMAGE_BUFFER_EXT"
	enumLookup[TEXTURE_MATERIAL_PARAMETER_EXT] = "TEXTURE_MATERIAL_PARAMETER_EXT"
	enumLookup[DRAW_BUFFER13_ATI] = "DRAW_BUFFER13_ATI"
	enumLookup[RENDERBUFFER_WIDTH_OES] = "RENDERBUFFER_WIDTH_OES"
	enumLookup[LUMINANCE_ALPHA_SNORM] = "LUMINANCE_ALPHA_SNORM"
	enumLookup[STENCIL_BACK_OP_VALUE_AMD] = "STENCIL_BACK_OP_VALUE_AMD"
	enumLookup[FACTOR_ALPHA_MODULATE_IMG] = "FACTOR_ALPHA_MODULATE_IMG"
	enumLookup[LUMINANCE16_ALPHA16] = "LUMINANCE16_ALPHA16"
	enumLookup[RG16F_EXT] = "RG16F_EXT"
	enumLookup[ENABLE_BIT] = "ENABLE_BIT"
	enumLookup[ALPHA_TEST_FUNC] = "ALPHA_TEST_FUNC"
	enumLookup[INTERNALFORMAT_SHARED_SIZE] = "INTERNALFORMAT_SHARED_SIZE"
	enumLookup[DYNAMIC_DRAW] = "DYNAMIC_DRAW"
	enumLookup[RGB16I] = "RGB16I"
	enumLookup[TEXTURE_SWIZZLE_RGBA_EXT] = "TEXTURE_SWIZZLE_RGBA_EXT"
	enumLookup[SGIS_texture_filter4] = "SGIS_texture_filter4"
	enumLookup[TEXTURE_COORD_ARRAY_COUNT_EXT] = "TEXTURE_COORD_ARRAY_COUNT_EXT"
	enumLookup[OUTPUT_TEXTURE_COORD8_EXT] = "OUTPUT_TEXTURE_COORD8_EXT"
	enumLookup[ONE_MINUS_SRC1_COLOR] = "ONE_MINUS_SRC1_COLOR"
	enumLookup[TEXTURE_BINDING_BUFFER_ARB] = "TEXTURE_BINDING_BUFFER_ARB"
	enumLookup[INTENSITY32UI_EXT] = "INTENSITY32UI_EXT"
	enumLookup[MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT] = "MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT"
	enumLookup[INVALID_FRAMEBUFFER_OPERATION] = "INVALID_FRAMEBUFFER_OPERATION"
	enumLookup[POST_CONVOLUTION_ALPHA_BIAS] = "POST_CONVOLUTION_ALPHA_BIAS"
	enumLookup[TEXTURE24] = "TEXTURE24"
	enumLookup[INVARIANT_DATATYPE_EXT] = "INVARIANT_DATATYPE_EXT"
	enumLookup[RED_INTEGER_EXT] = "RED_INTEGER_EXT"
	enumLookup[PIXEL_MAP_I_TO_A_SIZE] = "PIXEL_MAP_I_TO_A_SIZE"
	enumLookup[VERTEX_SOURCE_ATI] = "VERTEX_SOURCE_ATI"
	enumLookup[MAX_PROGRAM_INSTRUCTIONS_ARB] = "MAX_PROGRAM_INSTRUCTIONS_ARB"
	enumLookup[EDGE_FLAG_ARRAY_POINTER_EXT] = "EDGE_FLAG_ARRAY_POINTER_EXT"
	enumLookup[TEXTURE_4D_BINDING_SGIS] = "TEXTURE_4D_BINDING_SGIS"
	enumLookup[VERTEX_ATTRIB_ARRAY4_NV] = "VERTEX_ATTRIB_ARRAY4_NV"
	enumLookup[RENDERBUFFER_RED_SIZE_OES] = "RENDERBUFFER_RED_SIZE_OES"
	enumLookup[RELATIVE_SMALL_CCW_ARC_TO_NV] = "RELATIVE_SMALL_CCW_ARC_TO_NV"
	enumLookup[ALREADY_SIGNALED] = "ALREADY_SIGNALED"
	enumLookup[TEXTURE_STORAGE_SPARSE_BIT_AMD] = "TEXTURE_STORAGE_SPARSE_BIT_AMD"
	enumLookup[T2F_IUI_V3F_EXT] = "T2F_IUI_V3F_EXT"
	enumLookup[UNSIGNED_INT_24_8] = "UNSIGNED_INT_24_8"
	enumLookup[TRANSFORM_FEEDBACK_BUFFER_BINDING] = "TRANSFORM_FEEDBACK_BUFFER_BINDING"
	enumLookup[DOUBLE_MAT3_EXT] = "DOUBLE_MAT3_EXT"
	enumLookup[SMOOTH_QUADRATIC_CURVE_TO_NV] = "SMOOTH_QUADRATIC_CURVE_TO_NV"
	enumLookup[SYNC_CONDITION] = "SYNC_CONDITION"
	enumLookup[MAX_SAMPLES_IMG] = "MAX_SAMPLES_IMG"
	enumLookup[MAT_COLOR_INDEXES_BIT_PGI] = "MAT_COLOR_INDEXES_BIT_PGI"
	enumLookup[LINE_SMOOTH_HINT] = "LINE_SMOOTH_HINT"
	enumLookup[STENCIL_INDEX] = "STENCIL_INDEX"
	enumLookup[INDEX_MATERIAL_FACE_EXT] = "INDEX_MATERIAL_FACE_EXT"
	enumLookup[MAX_TESS_EVALUATION_UNIFORM_COMPONENTS] = "MAX_TESS_EVALUATION_UNIFORM_COMPONENTS"
	enumLookup[VIDEO_COLOR_CONVERSION_OFFSET_NV] = "VIDEO_COLOR_CONVERSION_OFFSET_NV"
	enumLookup[MIN_SPARSE_LEVEL_AMD] = "MIN_SPARSE_LEVEL_AMD"
	enumLookup[REFERENCED_BY_GEOMETRY_SHADER] = "REFERENCED_BY_GEOMETRY_SHADER"
	enumLookup[ANY_SAMPLES_PASSED_CONSERVATIVE_EXT] = "ANY_SAMPLES_PASSED_CONSERVATIVE_EXT"
	enumLookup[TRUE] = "TRUE"
	enumLookup[MAP1_TEXTURE_COORD_3] = "MAP1_TEXTURE_COORD_3"
	enumLookup[AND_INVERTED] = "AND_INVERTED"
	enumLookup[MIN_EXT] = "MIN_EXT"
	enumLookup[COMPRESSED_RG] = "COMPRESSED_RG"
	enumLookup[VERTEX_STREAM1_ATI] = "VERTEX_STREAM1_ATI"
	enumLookup[COLOR_ATTACHMENT8_NV] = "COLOR_ATTACHMENT8_NV"
	enumLookup[PRESENT_DURATION_NV] = "PRESENT_DURATION_NV"
	enumLookup[TEXTURE_SAMPLES] = "TEXTURE_SAMPLES"
	enumLookup[TEXTURE_4DSIZE_SGIS] = "TEXTURE_4DSIZE_SGIS"
	enumLookup[SECONDARY_COLOR_ARRAY_POINTER_EXT] = "SECONDARY_COLOR_ARRAY_POINTER_EXT"
	enumLookup[ATTRIB_ARRAY_STRIDE_NV] = "ATTRIB_ARRAY_STRIDE_NV"
	enumLookup[OP_EXP_BASE_2_EXT] = "OP_EXP_BASE_2_EXT"
	enumLookup[CON_11_ATI] = "CON_11_ATI"
	enumLookup[SKIP_COMPONENTS2_NV] = "SKIP_COMPONENTS2_NV"
	enumLookup[CIRCULAR_CCW_ARC_TO_NV] = "CIRCULAR_CCW_ARC_TO_NV"
	enumLookup[SYNC_FENCE_APPLE] = "SYNC_FENCE_APPLE"
	enumLookup[FOG_BIT] = "FOG_BIT"
	enumLookup[POLYGON_OFFSET_BIAS_EXT] = "POLYGON_OFFSET_BIAS_EXT"
	enumLookup[MATRIX3_NV] = "MATRIX3_NV"
	enumLookup[EVAL_VERTEX_ATTRIB14_NV] = "EVAL_VERTEX_ATTRIB14_NV"
	enumLookup[DEPTH_BOUNDS_EXT] = "DEPTH_BOUNDS_EXT"
	enumLookup[SHADER_OBJECT_ARB] = "SHADER_OBJECT_ARB"
	enumLookup[UNSIGNED_INT_SAMPLER_1D_ARRAY] = "UNSIGNED_INT_SAMPLER_1D_ARRAY"
	enumLookup[IMAGE_CUBE_MAP_ARRAY_EXT] = "IMAGE_CUBE_MAP_ARRAY_EXT"
	enumLookup[ACTIVE_ATOMIC_COUNTER_BUFFERS] = "ACTIVE_ATOMIC_COUNTER_BUFFERS"
	enumLookup[EDGE_FLAG] = "EDGE_FLAG"
	enumLookup[PIXEL_MAP_A_TO_A_SIZE] = "PIXEL_MAP_A_TO_A_SIZE"
	enumLookup[LINEAR_DETAIL_COLOR_SGIS] = "LINEAR_DETAIL_COLOR_SGIS"
	enumLookup[OUTPUT_TEXTURE_COORD29_EXT] = "OUTPUT_TEXTURE_COORD29_EXT"
	enumLookup[TEXTURE_FREE_MEMORY_ATI] = "TEXTURE_FREE_MEMORY_ATI"
	enumLookup[COLOR_ATTACHMENT9_NV] = "COLOR_ATTACHMENT9_NV"
	enumLookup[MAX_GEOMETRY_SHADER_INVOCATIONS] = "MAX_GEOMETRY_SHADER_INVOCATIONS"
	enumLookup[MULTISAMPLE_BIT_3DFX] = "MULTISAMPLE_BIT_3DFX"
	enumLookup[CONVOLUTION_FILTER_BIAS_EXT] = "CONVOLUTION_FILTER_BIAS_EXT"
	enumLookup[DUAL_LUMINANCE_ALPHA4_SGIS] = "DUAL_LUMINANCE_ALPHA4_SGIS"
	enumLookup[TEXTURE26] = "TEXTURE26"
	enumLookup[T2F_N3F_V3F] = "T2F_N3F_V3F"
	enumLookup[DOT_PRODUCT_PASS_THROUGH_NV] = "DOT_PRODUCT_PASS_THROUGH_NV"
	enumLookup[FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM] = "FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM"
	enumLookup[QUAD_INTENSITY4_SGIS] = "QUAD_INTENSITY4_SGIS"
	enumLookup[TEXTURE7] = "TEXTURE7"
	enumLookup[MATRIX3_ARB] = "MATRIX3_ARB"
	enumLookup[VERTEX_ID_NV] = "VERTEX_ID_NV"
	enumLookup[VIDEO_COLOR_CONVERSION_MAX_NV] = "VIDEO_COLOR_CONVERSION_MAX_NV"
	enumLookup[DEBUG_LOGGED_MESSAGES_AMD] = "DEBUG_LOGGED_MESSAGES_AMD"
	enumLookup[VERSION_1_3] = "VERSION_1_3"
	enumLookup[POLYGON_OFFSET_FACTOR] = "POLYGON_OFFSET_FACTOR"
	enumLookup[TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB] = "TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[UNIFORM_IS_ROW_MAJOR] = "UNIFORM_IS_ROW_MAJOR"
	enumLookup[FLOAT_32_UNSIGNED_INT_24_8_REV] = "FLOAT_32_UNSIGNED_INT_24_8_REV"
	enumLookup[TEXTURE_SWIZZLE_A_EXT] = "TEXTURE_SWIZZLE_A_EXT"
	enumLookup[TEXTURE_LIGHTING_MODE_HP] = "TEXTURE_LIGHTING_MODE_HP"
	enumLookup[MATRIX2_ARB] = "MATRIX2_ARB"
	enumLookup[RENDERBUFFER_INTERNAL_FORMAT_EXT] = "RENDERBUFFER_INTERNAL_FORMAT_EXT"
	enumLookup[TEXTURE25] = "TEXTURE25"
	enumLookup[SOURCE2_ALPHA] = "SOURCE2_ALPHA"
	enumLookup[DS_BIAS_NV] = "DS_BIAS_NV"
	enumLookup[MAGNITUDE_BIAS_NV] = "MAGNITUDE_BIAS_NV"
	enumLookup[BUFFER_SIZE_ARB] = "BUFFER_SIZE_ARB"
	enumLookup[OP_LOG_BASE_2_EXT] = "OP_LOG_BASE_2_EXT"
	enumLookup[MIN_PROGRAM_TEXEL_OFFSET] = "MIN_PROGRAM_TEXEL_OFFSET"
	enumLookup[NUM_FRAGMENT_CONSTANTS_ATI] = "NUM_FRAGMENT_CONSTANTS_ATI"
	enumLookup[FRAGMENT_SHADER] = "FRAGMENT_SHADER"
	enumLookup[SINGLE_COLOR] = "SINGLE_COLOR"
	enumLookup[SECONDARY_COLOR_ARRAY_STRIDE_EXT] = "SECONDARY_COLOR_ARRAY_STRIDE_EXT"
	enumLookup[TEXTURE_RECTANGLE] = "TEXTURE_RECTANGLE"
	enumLookup[EDGE_FLAG_ARRAY] = "EDGE_FLAG_ARRAY"
	enumLookup[FUNC_SUBTRACT] = "FUNC_SUBTRACT"
	enumLookup[COVERAGE_AUTOMATIC_NV] = "COVERAGE_AUTOMATIC_NV"
	enumLookup[UNSIGNED_INT_IMAGE_2D] = "UNSIGNED_INT_IMAGE_2D"
	enumLookup[ACTIVE_VARIABLES] = "ACTIVE_VARIABLES"
	enumLookup[BLEND_SRC] = "BLEND_SRC"
	enumLookup[MAX_EVAL_ORDER] = "MAX_EVAL_ORDER"
	enumLookup[INDEX_ARRAY_BUFFER_BINDING_ARB] = "INDEX_ARRAY_BUFFER_BINDING_ARB"
	enumLookup[REG_8_ATI] = "REG_8_ATI"
	enumLookup[FONT_DESCENDER_BIT_NV] = "FONT_DESCENDER_BIT_NV"
	enumLookup[POINT_SIZE_RANGE] = "POINT_SIZE_RANGE"
	enumLookup[DETAIL_TEXTURE_LEVEL_SGIS] = "DETAIL_TEXTURE_LEVEL_SGIS"
	enumLookup[MIRRORED_REPEAT_OES] = "MIRRORED_REPEAT_OES"
	enumLookup[SGIS_texture_border_clamp] = "SGIS_texture_border_clamp"
	enumLookup[FILL] = "FILL"
	enumLookup[UNSIGNED_INT_IMAGE_2D_ARRAY_EXT] = "UNSIGNED_INT_IMAGE_2D_ARRAY_EXT"
	enumLookup[BUFFER_UPDATE_BARRIER_BIT_EXT] = "BUFFER_UPDATE_BARRIER_BIT_EXT"
	enumLookup[LUMINANCE_FLOAT32_APPLE] = "LUMINANCE_FLOAT32_APPLE"
	enumLookup[FIXED] = "FIXED"
	enumLookup[POINT_SIZE_MIN_SGIS] = "POINT_SIZE_MIN_SGIS"
	enumLookup[EVAL_VERTEX_ATTRIB2_NV] = "EVAL_VERTEX_ATTRIB2_NV"
	enumLookup[EVAL_VERTEX_ATTRIB8_NV] = "EVAL_VERTEX_ATTRIB8_NV"
	enumLookup[Y_EXT] = "Y_EXT"
	enumLookup[CON_22_ATI] = "CON_22_ATI"
	enumLookup[COUNTER_TYPE_AMD] = "COUNTER_TYPE_AMD"
	enumLookup[PATH_COVER_DEPTH_FUNC_NV] = "PATH_COVER_DEPTH_FUNC_NV"
	enumLookup[MAP1_VERTEX_4] = "MAP1_VERTEX_4"
	enumLookup[RGBA_FLOAT32_APPLE] = "RGBA_FLOAT32_APPLE"
	enumLookup[SAMPLER_2D_RECT_SHADOW] = "SAMPLER_2D_RECT_SHADOW"
	enumLookup[TEXTURE_COORD_NV] = "TEXTURE_COORD_NV"
	enumLookup[RGB8I] = "RGB8I"
	enumLookup[DEBUG_CATEGORY_SHADER_COMPILER_AMD] = "DEBUG_CATEGORY_SHADER_COMPILER_AMD"
	enumLookup[VERTEX_ATTRIB_ARRAY12_NV] = "VERTEX_ATTRIB_ARRAY12_NV"
	enumLookup[NEGATIVE_X_EXT] = "NEGATIVE_X_EXT"
	enumLookup[MAX_PROGRAM_ATTRIBS_ARB] = "MAX_PROGRAM_ATTRIBS_ARB"
	enumLookup[TEXTURE_RENDERBUFFER_NV] = "TEXTURE_RENDERBUFFER_NV"
	enumLookup[COMBINER_MUX_SUM_NV] = "COMBINER_MUX_SUM_NV"
	enumLookup[TEXTURE_BUFFER_OFFSET] = "TEXTURE_BUFFER_OFFSET"
	enumLookup[CON_25_ATI] = "CON_25_ATI"
	enumLookup[PERFMON_RESULT_AVAILABLE_AMD] = "PERFMON_RESULT_AVAILABLE_AMD"
	enumLookup[UNIFORM_BLOCK] = "UNIFORM_BLOCK"
	enumLookup[UNPACK_IMAGE_DEPTH_SGIS] = "UNPACK_IMAGE_DEPTH_SGIS"
	enumLookup[SRGB] = "SRGB"
	enumLookup[PATH_MITER_LIMIT_NV] = "PATH_MITER_LIMIT_NV"
	enumLookup[VERSION] = "VERSION"
	enumLookup[ELEMENT_ARRAY_POINTER_APPLE] = "ELEMENT_ARRAY_POINTER_APPLE"
	enumLookup[TEXTURE_LUMINANCE_TYPE] = "TEXTURE_LUMINANCE_TYPE"
	enumLookup[MEDIUM_INT] = "MEDIUM_INT"
	enumLookup[RECT_NV] = "RECT_NV"
	enumLookup[SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE] = "SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE"
	enumLookup[TEXTURE30] = "TEXTURE30"
	enumLookup[COMBINER_SCALE_NV] = "COMBINER_SCALE_NV"
	enumLookup[MAX_PROGRAM_TEXEL_OFFSET_NV] = "MAX_PROGRAM_TEXEL_OFFSET_NV"
	enumLookup[VIDEO_COLOR_CONVERSION_MIN_NV] = "VIDEO_COLOR_CONVERSION_MIN_NV"
	enumLookup[COMPRESSED_RGBA_ASTC_10x8_KHR] = "COMPRESSED_RGBA_ASTC_10x8_KHR"
	enumLookup[AMBIENT_AND_DIFFUSE] = "AMBIENT_AND_DIFFUSE"
	enumLookup[READ_PIXELS] = "READ_PIXELS"
	enumLookup[CURRENT_FOG_COORDINATE_EXT] = "CURRENT_FOG_COORDINATE_EXT"
	enumLookup[EXPAND_NORMAL_NV] = "EXPAND_NORMAL_NV"
	enumLookup[VERTEX_ATTRIB_ARRAY5_NV] = "VERTEX_ATTRIB_ARRAY5_NV"
	enumLookup[TEXTURE_STENCIL_SIZE_EXT] = "TEXTURE_STENCIL_SIZE_EXT"
	enumLookup[REG_27_ATI] = "REG_27_ATI"
	enumLookup[LOWER_LEFT] = "LOWER_LEFT"
	enumLookup[DOUBLE_MAT3x2_EXT] = "DOUBLE_MAT3x2_EXT"
	enumLookup[PIXEL_MAP_I_TO_G_SIZE] = "PIXEL_MAP_I_TO_G_SIZE"
	enumLookup[GENERATE_MIPMAP_SGIS] = "GENERATE_MIPMAP_SGIS"
	enumLookup[OPERAND2_RGB_ARB] = "OPERAND2_RGB_ARB"
	enumLookup[OUTPUT_TEXTURE_COORD26_EXT] = "OUTPUT_TEXTURE_COORD26_EXT"
	enumLookup[TEXCOORD4_BIT_PGI] = "TEXCOORD4_BIT_PGI"
	enumLookup[CLIP_DISTANCE2] = "CLIP_DISTANCE2"
	enumLookup[SAMPLE_ALPHA_TO_MASK_EXT] = "SAMPLE_ALPHA_TO_MASK_EXT"
	enumLookup[COLOR_MATRIX_SGI] = "COLOR_MATRIX_SGI"
	enumLookup[DEBUG_CALLBACK_FUNCTION_ARB] = "DEBUG_CALLBACK_FUNCTION_ARB"
	enumLookup[MAX_LAYERS] = "MAX_LAYERS"
	enumLookup[LUMINANCE8] = "LUMINANCE8"
	enumLookup[MAX_PROGRAM_ALU_INSTRUCTIONS_ARB] = "MAX_PROGRAM_ALU_INSTRUCTIONS_ARB"
	enumLookup[QUERY_RESULT_AVAILABLE] = "QUERY_RESULT_AVAILABLE"
	enumLookup[POINT_SIZE_MIN_ARB] = "POINT_SIZE_MIN_ARB"
	enumLookup[TEXTURE_MIN_LOD_SGIS] = "TEXTURE_MIN_LOD_SGIS"
	enumLookup[VERTEX_SHADER_INSTRUCTIONS_EXT] = "VERTEX_SHADER_INSTRUCTIONS_EXT"
	enumLookup[DYNAMIC_DRAW_ARB] = "DYNAMIC_DRAW_ARB"
	enumLookup[SAMPLER_3D_OES] = "SAMPLER_3D_OES"
	enumLookup[DEBUG_TYPE_OTHER] = "DEBUG_TYPE_OTHER"
	enumLookup[SUBTRACT] = "SUBTRACT"
	enumLookup[VERTEX_ARRAY_RANGE_LENGTH_APPLE] = "VERTEX_ARRAY_RANGE_LENGTH_APPLE"
	enumLookup[NUM_GENERAL_COMBINERS_NV] = "NUM_GENERAL_COMBINERS_NV"
	enumLookup[PIXEL_UNPACK_BUFFER_EXT] = "PIXEL_UNPACK_BUFFER_EXT"
	enumLookup[RGBA16] = "RGBA16"
	enumLookup[GL_1PASS_SGIS] = "GL_1PASS_SGIS"
	enumLookup[MAX_DUAL_SOURCE_DRAW_BUFFERS] = "MAX_DUAL_SOURCE_DRAW_BUFFERS"
	enumLookup[MIN_SAMPLE_SHADING_VALUE_ARB] = "MIN_SAMPLE_SHADING_VALUE_ARB"
	enumLookup[RENDERBUFFER_BINDING_ANGLE] = "RENDERBUFFER_BINDING_ANGLE"
	enumLookup[DECAL] = "DECAL"
	enumLookup[CLAMP_TO_EDGE_SGIS] = "CLAMP_TO_EDGE_SGIS"
	enumLookup[MIRRORED_REPEAT] = "MIRRORED_REPEAT"
	enumLookup[REPLACEMENT_CODE_ARRAY_TYPE_SUN] = "REPLACEMENT_CODE_ARRAY_TYPE_SUN"
	enumLookup[VARIANT_EXT] = "VARIANT_EXT"
	enumLookup[FRAGMENT_SHADER_DERIVATIVE_HINT_ARB] = "FRAGMENT_SHADER_DERIVATIVE_HINT_ARB"
	enumLookup[TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV] = "TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV"
	enumLookup[ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER] = "ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER"
	enumLookup[UNIFORM] = "UNIFORM"
	enumLookup[SGI_color_table] = "SGI_color_table"
	enumLookup[GREEN_BITS] = "GREEN_BITS"
	enumLookup[BLUE] = "BLUE"
	enumLookup[R1UI_C3F_V3F_SUN] = "R1UI_C3F_V3F_SUN"
	enumLookup[DRAW_BUFFER6_ARB] = "DRAW_BUFFER6_ARB"
	enumLookup[PIXEL_MAP_I_TO_A] = "PIXEL_MAP_I_TO_A"
	enumLookup[UNPACK_CMYK_HINT_EXT] = "UNPACK_CMYK_HINT_EXT"
	enumLookup[DEPTH_COMPONENT16_SGIX] = "DEPTH_COMPONENT16_SGIX"
	enumLookup[X_EXT] = "X_EXT"
	enumLookup[COLOR_ATTACHMENT1_EXT] = "COLOR_ATTACHMENT1_EXT"
	enumLookup[SAMPLER_RENDERBUFFER_NV] = "SAMPLER_RENDERBUFFER_NV"
	enumLookup[Z6Y10Z6CB10Z6Y10Z6CR10_422_NV] = "Z6Y10Z6CB10Z6Y10Z6CR10_422_NV"
	enumLookup[SECONDARY_COLOR_ARRAY_LIST_IBM] = "SECONDARY_COLOR_ARRAY_LIST_IBM"
	enumLookup[MAX_DEBUG_GROUP_STACK_DEPTH] = "MAX_DEBUG_GROUP_STACK_DEPTH"
	enumLookup[INTERPOLATE] = "INTERPOLATE"
	enumLookup[OUTPUT_TEXTURE_COORD9_EXT] = "OUTPUT_TEXTURE_COORD9_EXT"
	enumLookup[TIME_ELAPSED] = "TIME_ELAPSED"
	enumLookup[TIMEOUT_EXPIRED] = "TIMEOUT_EXPIRED"
	enumLookup[ALL_BARRIER_BITS] = "ALL_BARRIER_BITS"
	enumLookup[DRAW_INDIRECT_BUFFER] = "DRAW_INDIRECT_BUFFER"
	enumLookup[INT_IMAGE_1D_ARRAY] = "INT_IMAGE_1D_ARRAY"
	enumLookup[STENCIL_BITS] = "STENCIL_BITS"
	enumLookup[CLIP_DISTANCE5] = "CLIP_DISTANCE5"
	enumLookup[DEBUG_TYPE_POP_GROUP] = "DEBUG_TYPE_POP_GROUP"
	enumLookup[VERTEX_WEIGHT_ARRAY_POINTER_EXT] = "VERTEX_WEIGHT_ARRAY_POINTER_EXT"
	enumLookup[MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB] = "MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB"
	enumLookup[FRAGMENT_SHADER_DERIVATIVE_HINT] = "FRAGMENT_SHADER_DERIVATIVE_HINT"
	enumLookup[CULL_VERTEX_IBM] = "CULL_VERTEX_IBM"
	enumLookup[POST_COLOR_MATRIX_COLOR_TABLE_SGI] = "POST_COLOR_MATRIX_COLOR_TABLE_SGI"
	enumLookup[FRAMEBUFFER_ATTACHMENT_BLUE_SIZE] = "FRAMEBUFFER_ATTACHMENT_BLUE_SIZE"
	enumLookup[PREVIOUS_ARB] = "PREVIOUS_ARB"
	enumLookup[DRAW_BUFFER7_ATI] = "DRAW_BUFFER7_ATI"
	enumLookup[INSTRUMENT_BUFFER_POINTER_SGIX] = "INSTRUMENT_BUFFER_POINTER_SGIX"
	enumLookup[SOURCE0_ALPHA] = "SOURCE0_ALPHA"
	enumLookup[VECTOR_EXT] = "VECTOR_EXT"
	enumLookup[MAX_COMBINED_UNIFORM_BLOCKS] = "MAX_COMBINED_UNIFORM_BLOCKS"
	enumLookup[TEXTURE_BINDING_BUFFER] = "TEXTURE_BINDING_BUFFER"
	enumLookup[DSDT8_MAG8_NV] = "DSDT8_MAG8_NV"
	enumLookup[RGB5] = "RGB5"
	enumLookup[UNPACK_COMPRESSED_SIZE_SGIX] = "UNPACK_COMPRESSED_SIZE_SGIX"
	enumLookup[RGBA16F_ARB] = "RGBA16F_ARB"
	enumLookup[COUNT_UP_NV] = "COUNT_UP_NV"
	enumLookup[COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR] = "COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR"
	enumLookup[DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB] = "DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB"
	enumLookup[TEXTURE30_ARB] = "TEXTURE30_ARB"
	enumLookup[ARRAY_BUFFER_ARB] = "ARRAY_BUFFER_ARB"
	enumLookup[GL_8X_BIT_ATI] = "GL_8X_BIT_ATI"
	enumLookup[TEXTURE_BUFFER_FORMAT_EXT] = "TEXTURE_BUFFER_FORMAT_EXT"
	enumLookup[COVERAGE_EDGE_FRAGMENTS_NV] = "COVERAGE_EDGE_FRAGMENTS_NV"
	enumLookup[PATH_TERMINAL_END_CAP_NV] = "PATH_TERMINAL_END_CAP_NV"
	enumLookup[RELATIVE_LARGE_CW_ARC_TO_NV] = "RELATIVE_LARGE_CW_ARC_TO_NV"
	enumLookup[SPOT_DIRECTION] = "SPOT_DIRECTION"
	enumLookup[MIN] = "MIN"
	enumLookup[IUI_N3F_V3F_EXT] = "IUI_N3F_V3F_EXT"
	enumLookup[VERTEX_STREAM6_ATI] = "VERTEX_STREAM6_ATI"
	enumLookup[FLOAT_R16_NV] = "FLOAT_R16_NV"
	enumLookup[TRANSFORM_FEEDBACK_NV] = "TRANSFORM_FEEDBACK_NV"
	enumLookup[MIN_PROGRAM_TEXTURE_GATHER_OFFSET] = "MIN_PROGRAM_TEXTURE_GATHER_OFFSET"
	enumLookup[PACK_SKIP_IMAGES] = "PACK_SKIP_IMAGES"
	enumLookup[INVALID_OPERATION] = "INVALID_OPERATION"
	enumLookup[MODELVIEW0_STACK_DEPTH_EXT] = "MODELVIEW0_STACK_DEPTH_EXT"
	enumLookup[EDGE_FLAG_ARRAY_COUNT_EXT] = "EDGE_FLAG_ARRAY_COUNT_EXT"
	enumLookup[OUTPUT_TEXTURE_COORD14_EXT] = "OUTPUT_TEXTURE_COORD14_EXT"
	enumLookup[TYPE] = "TYPE"
	enumLookup[RGBA_FLOAT16_ATI] = "RGBA_FLOAT16_ATI"
	enumLookup[LINK_STATUS] = "LINK_STATUS"
	enumLookup[SGIS_texture_edge_clamp] = "SGIS_texture_edge_clamp"
	enumLookup[MAX_VERTEX_ATTRIB_RELATIVE_OFFSET] = "MAX_VERTEX_ATTRIB_RELATIVE_OFFSET"
	enumLookup[SLICE_ACCUM_SUN] = "SLICE_ACCUM_SUN"
	enumLookup[DEBUG_OBJECT_MESA] = "DEBUG_OBJECT_MESA"
	enumLookup[PALETTE8_RGB8_OES] = "PALETTE8_RGB8_OES"
	enumLookup[MAX_GEOMETRY_VARYING_COMPONENTS_EXT] = "MAX_GEOMETRY_VARYING_COMPONENTS_EXT"
	enumLookup[DEBUG_CATEGORY_DEPRECATION_AMD] = "DEBUG_CATEGORY_DEPRECATION_AMD"
	enumLookup[TRIANGLES_ADJACENCY] = "TRIANGLES_ADJACENCY"
	enumLookup[SPRITE_EYE_ALIGNED_SGIX] = "SPRITE_EYE_ALIGNED_SGIX"
	enumLookup[FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX] = "FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX"
	enumLookup[MODELVIEW_PROJECTION_NV] = "MODELVIEW_PROJECTION_NV"
	enumLookup[MAP1_VERTEX_ATTRIB1_4_NV] = "MAP1_VERTEX_ATTRIB1_4_NV"
	enumLookup[MAP1_VERTEX_ATTRIB10_4_NV] = "MAP1_VERTEX_ATTRIB10_4_NV"
	enumLookup[PROGRAM_TEX_INSTRUCTIONS_ARB] = "PROGRAM_TEX_INSTRUCTIONS_ARB"
	enumLookup[ACTIVE_UNIFORMS] = "ACTIVE_UNIFORMS"
}
func (e Enum) String() string {
	name, ok := enumLookup[e]
	if ok {
		return name
	}
	return fmt.Sprintf("Enum(%d)", e)
}

type Context struct {
	access                    sync.Mutex
	context                   *C.gl12Context
	extensions                map[string]bool
	inBeginEnd                bool
	traceback                 []string
	Accum                     func(op Enum, value float32)
	AlphaFunc                 func(Func Enum, ref float32)
	Begin                     func(mode Enum)
	End                       func()
	Bitmap                    func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8)
	BlendFunc                 func(sfactor, dfactor Enum)
	CallList                  func(list uint32)
	CallLists                 func(n int32, Type Enum, lists unsafe.Pointer)
	Clear                     func(mask uint32)
	ClearAccum                func(red, green, blue, alpha float32)
	ClearColor                func(red, green, blue, alpha float32)
	ClearDepth                func(depth float64)
	ClearIndex                func(c float32)
	ClearStencil              func(s int32)
	ClipPlane                 func(plane Enum, equation *float64)
	Color3b                   func(red, green, blue int8)
	Color3d                   func(red, green, blue float64)
	Color3f                   func(red, green, blue float32)
	Color3i                   func(red, green, blue int32)
	Color3s                   func(red, green, blue int16)
	Color3ub                  func(red, green, blue uint8)
	Color3ui                  func(red, green, blue uint32)
	Color3us                  func(red, green, blue uint16)
	Color4b                   func(red, green, blue, alpha int8)
	Color4d                   func(red, green, blue, alpha float64)
	Color4f                   func(red, green, blue, alpha float32)
	Color4i                   func(red, green, blue, alpha int32)
	Color4s                   func(red, green, blue, alpha int16)
	Color4ub                  func(red, green, blue, alpha uint8)
	Color4ui                  func(red, green, blue, alpha uint32)
	Color4us                  func(red, green, blue, alpha uint16)
	Color3bv                  func(v *int8)
	Color3dv                  func(v *float64)
	Color3fv                  func(v *float32)
	Color3iv                  func(v *int32)
	Color3sv                  func(v *int16)
	Color3ubv                 func(v *uint8)
	Color3uiv                 func(v *uint32)
	Color3usv                 func(v *uint16)
	Color4bv                  func(v *int8)
	Color4dv                  func(v *float64)
	Color4fv                  func(v *float32)
	Color4iv                  func(v *int32)
	Color4sv                  func(v *int16)
	Color4ubv                 func(v *uint8)
	Color4uiv                 func(v *uint32)
	Color4usv                 func(v *uint16)
	ColorMask                 func(red, green, blue, alpha bool)
	ColorMaterial             func(face, mode Enum)
	CopyPixels                func(x, y int32, width, height int32, Type Enum)
	CullFace                  func(mode Enum)
	DeleteLists               func(list uint32, Range int32)
	DepthFunc                 func(Func Enum)
	DepthMask                 func(flag bool)
	DepthRange                func(zNear, zFar float64)
	Enable                    func(cap Enum)
	Disable                   func(cap Enum)
	DrawBuffer                func(mode Enum)
	DrawPixels                func(width, height int32, format, Type Enum, data unsafe.Pointer)
	EdgeFlag                  func(flag bool)
	EdgeFlagv                 func(flag *bool)
	EdgeFlagPointer           func(stride int32, pointer unsafe.Pointer)
	EvalCoord1d               func(u float64)
	EvalCoord1f               func(u float32)
	EvalCoord2d               func(u, v float64)
	EvalCoord2f               func(u, v float32)
	EvalCoord1dv              func(u *float64)
	EvalCoord1fv              func(u *float32)
	EvalCoord2dv              func(u *float64)
	EvalCoord2fv              func(u *float32)
	EvalMesh1                 func(mode Enum, i1, i2 int32)
	EvalMesh2                 func(mode Enum, i1, i2, j1, j2 int32)
	EvalPoint1                func(i int32)
	EvalPoint2                func(i, j int32)
	FeedbackBuffer            func(size int32, Type Enum, buffer *float32)
	Finish                    func()
	Flush                     func()
	Fogf                      func(pname Enum, param float32)
	Fogi                      func(pname Enum, param int32)
	Fogfv                     func(pname Enum, params *float32)
	Fogiv                     func(pname Enum, params *int32)
	FrontFace                 func(mode Enum)
	Frustum                   func(left, right, bottom, top, zNear, zFar float64)
	GenLists                  func(Range int32) uint32
	GetBooleanv               func(pname Enum, params *bool)
	GetDoublev                func(pname Enum, params *float64)
	GetFloatv                 func(pname Enum, params *float32)
	GetIntegerv               func(pname Enum, params *int32)
	GetClipPlane              func(plane Enum, equation *float64)
	GetError                  func() Enum
	GetLightfv                func(light, pname Enum, params *float32)
	GetLightiv                func(light, pname Enum, params *int32)
	GetMapdv                  func(target, query Enum, v *float64)
	GetMapfv                  func(target, query Enum, v *float32)
	GetMapiv                  func(target, query Enum, v *int32)
	GetMaterialfv             func(face, pname Enum, params *float32)
	GetMaterialiv             func(face, pname Enum, params *int32)
	GetPixelMapfv             func(Map Enum, values *float32)
	GetPixelMapuiv            func(Map Enum, values *uint32)
	GetPixelMapusv            func(Map Enum, values *uint16)
	GetPolygonStipple         func(pattern *uint8)
	GetString                 func(name Enum) string
	GetTexEnvfv               func(target, pname Enum, params *float32)
	GetTexEnviv               func(target, pname Enum, params *int32)
	GetTexGendv               func(coord, pname Enum, params *float64)
	GetTexGenfv               func(coord, pname Enum, params *float32)
	GetTexGeniv               func(coord, pname Enum, params *int32)
	GetTexImage               func(target Enum, level int32, format, Type Enum, pixels unsafe.Pointer)
	GetTexLevelParameterfv    func(target Enum, level int32, pname Enum, params *float32)
	GetTexLevelParameteriv    func(target Enum, level int32, pname Enum, params *int32)
	GetTexParameterfv         func(target, pname Enum, params *float32)
	GetTexParameteriv         func(target, pname Enum, params *int32)
	Hint                      func(target, mode Enum)
	Indexd                    func(c float64)
	Indexf                    func(c float32)
	Indexi                    func(c int32)
	Indexs                    func(c int16)
	Indexdv                   func(c *float64)
	Indexfv                   func(c *float32)
	Indexiv                   func(c *int32)
	Indexsv                   func(c *int16)
	IndexMask                 func(mask uint32)
	IndexPointer              func(Type Enum, stride int32, pointer unsafe.Pointer)
	InitNames                 func()
	IsEnabled                 func(cap Enum)
	IsList                    func(list uint32) bool
	Lightf                    func(light, pname Enum, param float32)
	Lighti                    func(light, pname Enum, param int32)
	Lightfv                   func(light, pname Enum, params *float32)
	Lightiv                   func(light, pname Enum, params *int32)
	LightModelf               func(pname Enum, param float32)
	LightModeli               func(pname Enum, param int32)
	LightModelfv              func(pname Enum, params *float32)
	LightModeliv              func(pname Enum, params *int32)
	LineStipple               func(factor int32, pattern uint16)
	LineWidth                 func(width float32)
	ListBase                  func(base uint32)
	LoadIdentity              func()
	LoadMatrixd               func(m *float64)
	LoadMatrixf               func(m *float32)
	LoadName                  func(name uint32)
	LogicOp                   func(opcode Enum)
	Map1d                     func(target Enum, u1, u2 float64, stride, order int32, points *float64)
	Map1f                     func(target Enum, u1, u2 float32, stride, order int32, points *float32)
	Map2d                     func(target Enum, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64)
	Map2f                     func(target Enum, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32)
	MapGrid1d                 func(un int32, u1, u2 float64)
	MapGrid1f                 func(un int32, u1, u2 float32)
	MapGrid2d                 func(un int32, u1, u2 float64, vn int32, v1, v2 float64)
	MapGrid2f                 func(un int32, u1, u2 float32, vn int32, v1, v2 float32)
	Materialf                 func(face, pname Enum, param float32)
	Materiali                 func(face, pname Enum, param int32)
	Materialfv                func(face, pname Enum, params *float32)
	Materialiv                func(face, pname Enum, params *int32)
	MatrixMode                func(mode Enum)
	MultMatrixd               func(m *float64)
	MultMatrixf               func(m *float32)
	NewList                   func(list uint32, mode Enum)
	EndList                   func()
	Normal3b                  func(nx, ny, nz int8)
	Normal3d                  func(nx, ny, nz float64)
	Normal3f                  func(nx, ny, nz float32)
	Normal3i                  func(nx, ny, nz int32)
	Normal3s                  func(nx, ny, nz int16)
	Normal3bv                 func(v *int8)
	Normal3dv                 func(v *float64)
	Normal3fv                 func(v *float32)
	Normal3iv                 func(v *int32)
	Normal3sv                 func(v *int16)
	Ortho                     func(left, right, bottom, top, zNear, zfar float64)
	PassThrough               func(token float32)
	PixelMapfv                func(Map Enum, mapsize int32, values *float32)
	PixelMapuiv               func(Map Enum, mapsize int32, values *uint32)
	PixelMapusv               func(Map Enum, mapsize int32, values *uint16)
	PixelStoref               func(pname Enum, param float32)
	PixelStorei               func(pname Enum, param int32)
	PixelTransferf            func(pname Enum, param float32)
	PixelTransferi            func(pname Enum, param int32)
	PixelZoom                 func(xfactor, yfactor float32)
	PointSize                 func(size float32)
	PolygonMode               func(face, mode Enum)
	PolygonStipple            func(mask *uint8)
	PushAttrib                func(mask uint32)
	PopAttrib                 func()
	PushMatrix                func()
	PopMatrix                 func()
	PushName                  func(name uint32)
	PopName                   func()
	RasterPos2d               func(x, y float64)
	RasterPos2f               func(x, y float32)
	RasterPos2i               func(x, y int32)
	RasterPos2s               func(x, y int16)
	RasterPos3d               func(x, y, z float64)
	RasterPos3f               func(x, y, z float32)
	RasterPos3i               func(x, y, z int32)
	RasterPos3s               func(x, y, z int16)
	RasterPos4d               func(x, y, z, w float64)
	RasterPos4f               func(x, y, z, w float32)
	RasterPos4i               func(x, y, z, w int32)
	RasterPos4s               func(x, y, z, w int16)
	RasterPos2dv              func(v *float64)
	RasterPos2fv              func(v *float32)
	RasterPos2iv              func(v *int32)
	RasterPos2sv              func(v *int16)
	RasterPos3dv              func(v *float64)
	RasterPos3fv              func(v *float32)
	RasterPos3iv              func(v *int32)
	RasterPos3sv              func(v *int16)
	RasterPos4dv              func(v *float64)
	RasterPos4fv              func(v *float32)
	RasterPos4iv              func(v *int32)
	RasterPos4sv              func(v *int16)
	ReadBuffer                func(mode Enum)
	ReadPixels                func(x, y int32, width, height int32, format, Type Enum, pixels unsafe.Pointer)
	Rectd                     func(x1, y1, x2, y2 float64)
	Rectf                     func(x1, y1, x2, y2 float32)
	Recti                     func(x1, y1, x2, y2 int32)
	Rects                     func(x1, y1, x2, y2 int16)
	Rectdv                    func(v1, v2 *float64)
	Rectfv                    func(v1, v2 *float32)
	Rectiv                    func(v1, v2 *int32)
	Rectsv                    func(v1, v2 *int16)
	RenderMode                func(mode Enum) int32
	Rotated                   func(angle, x, y, z float64)
	Rotatef                   func(angle, x, y, z float32)
	Scaled                    func(x, y, z float64)
	Scalef                    func(x, y, z float32)
	Scissor                   func(x, y int32, width, height int32)
	SelectBuffer              func(size int32, buffer *uint32)
	ShadeModel                func(mode Enum)
	StencilFunc               func(Func Enum, ref int32, mask uint32)
	StencilMask               func(mask uint32)
	StencilOp                 func(fail, zfail, zpass Enum)
	TexCoord1d                func(s float64)
	TexCoord1f                func(s float32)
	TexCoord1i                func(s int32)
	TexCoord1s                func(s int16)
	TexCoord2d                func(s, t float64)
	TexCoord2f                func(s, t float32)
	TexCoord2i                func(s, t int32)
	TexCoord2s                func(s, t int16)
	TexCoord3d                func(s, t, r float64)
	TexCoord3f                func(s, t, r float32)
	TexCoord3i                func(s, t, r int32)
	TexCoord3s                func(s, t, r int16)
	TexCoord4d                func(s, t, r, q float64)
	TexCoord4f                func(s, t, r, q float32)
	TexCoord4i                func(s, t, r, q int32)
	TexCoord4s                func(s, t, r, q int16)
	TexCoord1dv               func(v *float64)
	TexCoord1fv               func(v *float32)
	TexCoord1iv               func(v *int32)
	TexCoord1sv               func(v *int16)
	TexCoord2dv               func(v *float64)
	TexCoord2fv               func(v *float32)
	TexCoord2iv               func(v *int32)
	TexCoord2sv               func(v *int16)
	TexCoord3dv               func(v *float64)
	TexCoord3fv               func(v *float32)
	TexCoord3iv               func(v *int32)
	TexCoord3sv               func(v *int16)
	TexCoord4dv               func(v *float64)
	TexCoord4fv               func(v *float32)
	TexCoord4iv               func(v *int32)
	TexCoord4sv               func(v *int16)
	TexEnvf                   func(target, pname Enum, param float32)
	TexEnvi                   func(target, pname Enum, param int32)
	TexEnvfv                  func(target, pname Enum, params *float32)
	TexEnviv                  func(target, pname Enum, params *int32)
	TexGend                   func(coord, pname Enum, param float64)
	TexGenf                   func(coord, pname Enum, param float32)
	TexGeni                   func(coord, pname Enum, param int32)
	TexGendv                  func(coord, pname Enum, params *float64)
	TexGenfv                  func(coord, pname Enum, params *float32)
	TexGeniv                  func(coord, pname Enum, params *int32)
	TexImage1D                func(target Enum, level, internalformat int32, width int32, border int32, format, Type Enum, pixels unsafe.Pointer)
	TexImage2D                func(target Enum, level, internalformat int32, width, height int32, border int32, format, Type Enum, pixels unsafe.Pointer)
	TexParameterf             func(target, pname Enum, param float32)
	TexParameteri             func(target, pname Enum, param int32)
	TexParameterfv            func(target, pname Enum, params *float32)
	TexParameteriv            func(target, pname Enum, params *int32)
	Translated                func(x, y, z float64)
	Translatef                func(x, y, z float32)
	Vertex2s                  func(x, y int16)
	Vertex2i                  func(x, y int32)
	Vertex2f                  func(x, y float32)
	Vertex2d                  func(x, y float64)
	Vertex3s                  func(x, y, z int16)
	Vertex3i                  func(x, y, z int32)
	Vertex3f                  func(x, y, z float32)
	Vertex3d                  func(x, y, z float64)
	Vertex4s                  func(x, y, z, w int16)
	Vertex4i                  func(x, y, z, w int32)
	Vertex4f                  func(x, y, z, w float32)
	Vertex4d                  func(x, y, z, w float64)
	Viewport                  func(x, y int32, width, height int32)
	GetConvolutionParameterfv func(target, pname Enum, params *float32)
	GetConvolutionParameteriv func(target, pname Enum, params *int32)
	AreTexturesResident       func(textures []uint32) (status bool, residencies []bool)
	ArrayElement              func(i int32)
	DrawArrays                func(mode Enum, first int32, count int32)
	DrawElements              func(mode Enum, count int32, Type Enum, indices unsafe.Pointer)
	GetPointerv               func(pname Enum, params unsafe.Pointer)
	PolygonOffset             func(factor, units float32)
	CopyTexImage1D            func(target Enum, level int32, internalFormat Enum, x, y int32, width int32, border int32)
	CopyTexImage2D            func(target Enum, level int32, internalFormat Enum, x, y int32, width, height int32, border int32)
	CopyTexSubImage1D         func(target Enum, level, xoffset, x, y int32, width int32)
	CopyTexSubImage2D         func(target Enum, level, xoffset, yoffset, x, y int32, width, height int32)
	BindTexture               func(target Enum, texture uint32)
	DeleteTextures            func(n int32, textures *uint32)
	GenTextures               func(n int32, textures *uint32)
	IsTexture                 func(texture uint32) bool
	ColorPointer              func(size int32, Type Enum, stride int32, pointer unsafe.Pointer)
	EnableClientState         func(cap Enum)
	DisableClientState        func(cap Enum)
	Indexub                   func(c uint8)
	Indexubv                  func(c *uint8)
	InterleavedArrays         func(format Enum, stride int32, pointer unsafe.Pointer)
	NormalPointer             func(Type Enum, stride int32, pointer unsafe.Pointer)
	PushClientAttrib          func(mask uint32)
	PrioritizeTextures        func(n int32, textures *uint32, priorities *float32)
	PopClientAttrib           func()
	TexCoordPointer           func(size int32, Type Enum, stride int32, pointer unsafe.Pointer)
	TexSubImage1D             func(target Enum, level, xoffset int32, width int32, format, Type Enum, pixels unsafe.Pointer)
	TexSubImage2D             func(target Enum, level, xoffset, yoffset int32, width, height int32, format, Type Enum, pixels unsafe.Pointer)
	VertexPointer             func(size int32, Type Enum, stride int32, pointer unsafe.Pointer)
	ColorTable                func(target, internalformat Enum, width int32, format, Type Enum, data unsafe.Pointer)
	ColorTableParameterfv     func(target, pname Enum, params *float32)
	ColorTableParameteriv     func(target, pname Enum, params *int32)
	ColorSubTable             func(target Enum, start, count int32, format, Type Enum, data unsafe.Pointer)
	ConvolutionFilter1D       func(target, internalformat Enum, width int32, format, Type Enum, data unsafe.Pointer)
	ConvolutionFilter2D       func(target, internalformat Enum, width, height int32, format, Type Enum, data unsafe.Pointer)
	ConvolutionParameterf     func(target, pname Enum, params float32)
	ConvolutionParameteri     func(target, pname Enum, params int32)
	CopyColorTable            func(target, internalformat Enum, x, y int32, width int32)
	CopyColorSubTable         func(target Enum, start int32, x, y int32, width int32)
	CopyConvolutionFilter1D   func(target, internalformat Enum, x, y int32, width int32)
	CopyConvolutionFilter2D   func(target, internalformat Enum, x, y int32, width, height int32)
	GetColorTable             func(target, format, Type Enum, table unsafe.Pointer)
	GetColorTableParameterfv  func(target, pname Enum, params *float32)
	GetColorTableParameteriv  func(target, pname Enum, params *int32)
	GetConvolutionFilter      func(target, format, Type Enum, image unsafe.Pointer)
	GetHistogram              func(target Enum, reset bool, format, Type Enum, values unsafe.Pointer)
	GetHistogramParameterfv   func(target, pname Enum, params *float32)
	GetHistogramParameteriv   func(target, pname Enum, params *int32)
	GetSeparableFilter        func(target, format, Type Enum, row, column, span unsafe.Pointer)
	Histogram                 func(target Enum, width int32, internalformat Enum, sink bool)
	Minmax                    func(target, internalformat Enum, sink bool)
	MultiTexCoord1s           func(target Enum, s int16)
	MultiTexCoord1i           func(target Enum, s int32)
	MultiTexCoord1f           func(target Enum, s float32)
	MultiTexCoord1d           func(target Enum, s float64)
	MultiTexCoord2s           func(target Enum, s, t int16)
	MultiTexCoord2i           func(target Enum, s, t int32)
	MultiTexCoord2f           func(target Enum, s, t float32)
	MultiTexCoord2d           func(target Enum, s, t float64)
	MultiTexCoord3s           func(target Enum, s, t, r int16)
	MultiTexCoord3i           func(target Enum, s, t, r int32)
	MultiTexCoord3f           func(target Enum, s, t, r float32)
	MultiTexCoord3d           func(target Enum, s, t, r float64)
	MultiTexCoord4s           func(target Enum, s, t, r, q int16)
	MultiTexCoord4i           func(target Enum, s, t, r, q int32)
	MultiTexCoord4f           func(target Enum, s, t, r, q float32)
	MultiTexCoord4d           func(target Enum, s, t, r, q float64)
	MultiTexCoord1sv          func(target Enum, v *int16)
	MultiTexCoord1iv          func(target Enum, v *int32)
	MultiTexCoord1fv          func(target Enum, v *float32)
	MultiTexCoord1dv          func(target Enum, v *float64)
	MultiTexCoord2sv          func(target Enum, v *int16)
	MultiTexCoord2iv          func(target Enum, v *int32)
	MultiTexCoord2fv          func(target Enum, v *float32)
	MultiTexCoord2dv          func(target Enum, v *float64)
	MultiTexCoord3sv          func(target Enum, v *int16)
	MultiTexCoord3iv          func(target Enum, v *int32)
	MultiTexCoord3fv          func(target Enum, v *float32)
	MultiTexCoord3dv          func(target Enum, v *float64)
	MultiTexCoord4sv          func(target Enum, v *int16)
	MultiTexCoord4iv          func(target Enum, v *int32)
	MultiTexCoord4fv          func(target Enum, v *float32)
	MultiTexCoord4dv          func(target Enum, v *float64)
	ResetHistogram            func(target Enum)
	ResetMinmax               func(target Enum)
	SeparableFilter2D         func(target, internalformat Enum, width, height int32, format, Type Enum, row, column unsafe.Pointer)
	BlendColor                func(red, green, blue, alpha float32)
	BlendEquation             func(mode Enum)
	CopyTexSubImage3D         func(target Enum, level, xoffset, yoffset, zoffset, x, y int32, width, height int32)
	DrawRangeElements         func(mode Enum, start, end uint32, count int32, Type Enum, indices unsafe.Pointer)
	TexImage3D                func(target Enum, level, internalformat int32, width, height, depth int32, border int32, format, Type Enum, pixels unsafe.Pointer)
	TexSubImage3D             func(target Enum, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type Enum, pixels unsafe.Pointer)
	ActiveTexture             func(texture Enum)
	ClientActiveTexture       func(texture Enum)
	CompressedTexImage1D      func(target Enum, level int32, internalformat Enum, width int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage2D      func(target Enum, level int32, internalformat Enum, width, height int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage3D      func(target Enum, level int32, internalformat Enum, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage1D   func(target Enum, level, xoffset int32, width int32, format Enum, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage2D   func(target Enum, level, xoffset, yoffset int32, width, height int32, format Enum, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage3D   func(target Enum, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format Enum, imageSize int32, data unsafe.Pointer)
	GetCompressedTexImage     func(target Enum, lod int32, img unsafe.Pointer)
	LoadTransposeMatrixd      func(m *float64)
	LoadTransposeMatrixf      func(m *float64)
	MultTransposeMatrixd      func(m *float64)
	MultTransposeMatrixf      func(m *float32)
	SampleCoverage            func(value float32, invert bool)
	BlendFuncSeparate         func(srcRGB, dstRGB, srcAlpha, dstAlpha Enum)
	FogCoordPointer           func(Type Enum, stride int32, pointer unsafe.Pointer)
	FogCoordd                 func(coord float64)
	FogCoordf                 func(coord float32)
	FogCoorddv                func(coord *float64)
	FogCoordfv                func(coord *float32)
	MultiDrawArrays           func(mode Enum, first *int32, count *int32, primcount int32)
	MultiDrawElements         func(mode Enum, count *int32, Type Enum, indices unsafe.Pointer, primcount int32)
	PointParameterf           func(pname Enum, param float32)
	PointParameteri           func(pname Enum, param int32)
	SecondaryColor3b          func(red, green, blue int8)
	SecondaryColor3s          func(red, green, blue int16)
	SecondaryColor3i          func(red, green, blue int32)
	SecondaryColor3f          func(red, green, blue float32)
	SecondaryColor3d          func(red, green, blue float64)
	SecondaryColor3ub         func(red, green, blue uint8)
	SecondaryColor3us         func(red, green, blue uint16)
	SecondaryColor3ui         func(red, green, blue uint32)
	SecondaryColor3bv         func(v *int8)
	SecondaryColor3sv         func(v *int16)
	SecondaryColor3iv         func(v *int32)
	SecondaryColor3fv         func(v *float32)
	SecondaryColor3dv         func(v *float64)
	SecondaryColor3ubv        func(v *uint8)
	SecondaryColor3usv        func(v *uint16)
	SecondaryColor3uiv        func(v *uint32)
	SecondaryColorPointer     func(size int32, Type Enum, stride int32, pointer unsafe.Pointer)
	WindowPos2s               func(x, y int16)
	WindowPos2i               func(x, y int32)
	WindowPos2f               func(x, y float32)
	WindowPos2d               func(x, y float64)
	WindowPos3s               func(x, y, z int16)
	WindowPos3i               func(x, y, z int32)
	WindowPos3f               func(x, y, z float32)
	WindowPos3d               func(x, y, z float64)
	WindowPos2sv              func(v *int16)
	WindowPos2iv              func(v *int32)
	WindowPos2fv              func(v *float32)
	WindowPos2dv              func(v *float64)
	WindowPos3sv              func(v *int16)
	WindowPos3iv              func(v *int32)
	WindowPos3fv              func(v *float32)
	WindowPos3dv              func(v *float64)
	BeginQuery                func(target Enum, id uint32)
	BindBuffer                func(target Enum, buffer uint32)
	BufferData                func(target Enum, size int32, data unsafe.Pointer, usage Enum)
	BufferSubData             func(target, offset Enum, size int32, data unsafe.Pointer)
	DeleteBuffers             func(n int32, buffers *uint32)
	DeleteQueries             func(n int32, ids *uint32)
	GenBuffers                func(n int32, buffers *uint32)
	GenQueries                func(n int32, ids *uint32)
	GetBufferParameteriv      func(target, value Enum, data *int32)
	GetBufferPointerv         func(target, pname Enum, params unsafe.Pointer)
	GetBufferSubData          func(target Enum, offset int32, size int32, data unsafe.Pointer)
	GetQueryObjectiv          func(id uint32, pname Enum, params *int32)
	GetQueryObjectuiv         func(id uint32, pname Enum, params *uint32)
	GetQueryiv                func(target, pname Enum, params *int32)
	IsBuffer                  func(buffer uint32) bool
	IsQuery                   func(id uint32) bool
	MapBuffer                 func(target, access Enum) unsafe.Pointer
	UnmapBuffer               func(target Enum) bool
	AttachShader              func(program, shader uint32)
	BindAttribLocation        func(program, index uint32, name string)
	BlendEquationSeperate     func(modeRGB, modeAlpha Enum)
	CompileShader             func(shader uint32)
	CreateProgram             func() uint32
	CreateShader              func(shaderType Enum) uint32
	DeleteProgram             func(program uint32)
	DeleteShader              func(shader uint32)
	DetachShader              func(program, shader uint32)
	EnableVertexAttribArray   func(index uint32)
	DisableVertexAttribArray  func(index uint32)
	DrawBuffers               func(n int32, bufs *Enum)
	GetActiveAttrib           func(program, index uint32, bufSize int32) (length int32, size int32, Type Enum, name string)
	GetActiveUniform          func(program, index uint32, bufSize int32, length *int32, size *int32, Type *Enum, name *byte)
	GetAttachedShaders        func(program uint32, maxCount int32, count *int32, shaders *uint32)
	GetAttribLocation         func(program uint32, name *byte) int32
	GetProgramiv              func(program uint32, pname Enum, params *int32)
	GetProgramInfoLog         func(program uint32, maxLength int32, length *int32, infoLog *byte)
	GetShaderiv               func(program uint32, pname Enum, params *int32)
	GetShaderInfoLog          func(shader uint32, maxLength int32, length *int32, infoLog *byte)
	GetShaderSource           func(shader uint32, bufSize int32, length *int32, source *byte)
	GetUniformfv              func(program uint32, location int32, params *float32)
	GetUniformiv              func(program uint32, location int32, params *int32)
	GetUniformLocation        func(program uint32, name *byte) int32
	GetVertexAttribdv         func(index uint32, pname Enum, params *float64)
	GetVertexAttribfv         func(index uint32, pname Enum, params *float32)
	GetVertexAttribiv         func(index uint32, pname Enum, params *int32)
	GetVertexAttribPointerv   func(index uint32, pname Enum, pointer unsafe.Pointer)
	IsProgram                 func(program uint32) bool
	IsShader                  func(shader uint32) bool
	LinkProgram               func(program uint32)
	ShaderSource              func(shader uint32, count int32, string **byte, length *int32)
	StencilFuncSeparate       func(face, Func Enum, ref int32, mask uint32)
	StencilMaskSeparate       func(face Enum, mask uint32)
	StencilOpSeparate         func(face, sfail, dpfail, dppass Enum)
	Uniform1f                 func(location int32, v0 float32)
	Uniform2f                 func(location int32, v0, v1 float32)
	Uniform3f                 func(location int32, v0, v1, v2 float32)
	Uniform4f                 func(location int32, v0, v1, v2, v3 float32)
	Uniform1i                 func(location, v0 int32)
	Uniform2i                 func(location, v0, v1 int32)
	Uniform3i                 func(location, v0, v1, v2 int32)
	Uniform4i                 func(location, v0, v1, v2, v3 int32)
	Uniform1ui                func(location int32, v0 uint32)
	Uniform2ui                func(location int32, v0, v1 uint32)
	Uniform3ui                func(location int32, v0, v1, v2 uint32)
	Uniform4ui                func(location int32, v0, v1, v2, v3 uint32)
	Uniform1fv                func(location int32, count int32, value *float32)
	Uniform2fv                func(location int32, count int32, value *float32)
	Uniform3fv                func(location int32, count int32, value *float32)
	Uniform4fv                func(location int32, count int32, value *float32)
	Uniform1iv                func(location int32, count int32, value *int32)
	Uniform2iv                func(location int32, count int32, value *int32)
	Uniform3iv                func(location int32, count int32, value *int32)
	Uniform4iv                func(location int32, count int32, value *int32)
	Uniform1uiv               func(location int32, count int32, value *uint32)
	Uniform2uiv               func(location int32, count int32, value *uint32)
	Uniform3uiv               func(location int32, count int32, value *uint32)
	Uniform4uiv               func(location int32, count int32, value *uint32)
	UseProgram                func(program uint32)
	ValidateProgram           func(program uint32)
	VertexAttribPointer       func(index uint32, size int32, Type Enum, normalized bool, stride int32, pointer unsafe.Pointer)
	VertexAttrib1f            func(index uint32, v0 float32)
	VertexAttrib1s            func(index uint32, v0 int16)
	VertexAttrib1d            func(index uint32, v0 float64)
	VertexAttrib2f            func(index uint32, v0, v1 float32)
	VertexAttrib2s            func(index uint32, v0, v1 int16)
	VertexAttrib2d            func(index uint32, v0, v1 float64)
	VertexAttrib3f            func(index uint32, v0, v1, v2 float32)
	VertexAttrib3s            func(index uint32, v0, v1, v2 int16)
	VertexAttrib3d            func(index uint32, v0, v1, v2 float64)
	VertexAttrib4f            func(index uint32, v0, v1, v2, v3 float32)
	VertexAttrib4s            func(index uint32, v0, v1, v2, v3 int16)
	VertexAttrib4d            func(index uint32, v0, v1, v2, v3 float64)
	VertexAttrib4Nuv          func(index uint32, v0, v1, v2, v3 uint8)
	VertexAttrib1fv           func(index uint32, v *float32)
	VertexAttrib1sv           func(index uint32, v *int16)
	VertexAttrib1dv           func(index uint32, v *float64)
	VertexAttrib2fv           func(index uint32, v *float32)
	VertexAttrib2sv           func(index uint32, v *int16)
	VertexAttrib2dv           func(index uint32, v *float64)
	VertexAttrib3fv           func(index uint32, v *float32)
	VertexAttrib3sv           func(index uint32, v *int16)
	VertexAttrib3dv           func(index uint32, v *float64)
	VertexAttrib4fv           func(index uint32, v *float32)
	VertexAttrib4sv           func(index uint32, v *int16)
	VertexAttrib4dv           func(index uint32, v *float64)
	VertexAttrib4iv           func(index uint32, v *int32)
	VertexAttrib4bv           func(index uint32, v *int8)
	VertexAttrib4ubv          func(index uint32, v *uint8)
	VertexAttrib4usv          func(index uint32, v *uint16)
	VertexAttrib4uiv          func(index uint32, v *uint32)
	VertexAttrib4Nbv          func(index uint32, v *int8)
	VertexAttrib4Nsv          func(index uint32, v *int16)
	VertexAttrib4Niv          func(index uint32, v *int32)
	VertexAttrib4Nubv         func(index uint32, v *uint8)
	VertexAttrib4Nusv         func(index uint32, v *uint16)
	VertexAttrib4Nuiv         func(index uint32, v *uint32)
	UniformMatrix2fv          func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix3fv          func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix4fv          func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix2x3fv        func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix3x2fv        func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix2x4fv        func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix4x2fv        func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix3x4fv        func(location int32, count int32, transpose bool, value *float32)
	UniformMatrix4x3fv        func(location int32, count int32, transpose bool, value *float32)
}

func New() *Context {
	glc := new(Context)
	glc.context = C.gl12NewContext()

	glc.Accum = func(op Enum, value float32) {
		fmtCall := fmt.Sprintf("Accum(%v, %v)", op, value)
		defer glc.trace(fmtCall)
		C.gl12Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func Enum, ref float32) {
		fmtCall := fmt.Sprintf("AlphaFunc(%v, %v)", Func, ref)
		defer glc.trace(fmtCall)
		C.gl12AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode Enum) {
		fmtCall := fmt.Sprintf("Begin(%v)", mode)
		defer glc.trace(fmtCall)
		glc.inBeginEnd = true
		C.gl12Begin(glc.context, C.GLenum(mode))
		return
	}

	glc.End = func() {
		fmtCall := "End()"
		defer glc.trace(fmtCall)
		C.gl12End(glc.context)
		glc.inBeginEnd = false
		return
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8) {
		fmtCall := fmt.Sprintf("Bitmap(%v, %v, %v, %v, %v, %v, %v)", width, height, xorig, yorig, xmove, ymove, bitmap)
		defer glc.trace(fmtCall)
		C.gl12Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
	}

	glc.BlendFunc = func(sfactor, dfactor Enum) {
		fmtCall := fmt.Sprintf("BlendFunc(%v, %v)", sfactor, dfactor)
		defer glc.trace(fmtCall)
		C.gl12BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		fmtCall := fmt.Sprintf("CallList(%v)", list)
		defer glc.trace(fmtCall)
		C.gl12CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type Enum, lists unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CallLists(%v, %v, %v)", n, Type, lists)
		defer glc.trace(fmtCall)
		C.gl12CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		fmtCall := fmt.Sprintf("Clear(%v)", mask)
		defer glc.trace(fmtCall)
		C.gl12Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		fmtCall := fmt.Sprintf("ClearAccum(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		fmtCall := fmt.Sprintf("ClearColor(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		fmtCall := fmt.Sprintf("ClearDepth(%v)", depth)
		defer glc.trace(fmtCall)
		C.gl12ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		fmtCall := fmt.Sprintf("ClearIndex(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		fmtCall := fmt.Sprintf("ClearStencil(%v)", s)
		defer glc.trace(fmtCall)
		C.gl12ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane Enum, equation *float64) {
		fmtCall := fmt.Sprintf("ClipPlane(%v, %v)", plane, equation)
		defer glc.trace(fmtCall)
		C.gl12ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.Color3b = func(red, green, blue int8) {
		fmtCall := fmt.Sprintf("Color3b(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		fmtCall := fmt.Sprintf("Color3d(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		fmtCall := fmt.Sprintf("Color3f(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		fmtCall := fmt.Sprintf("Color3i(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		fmtCall := fmt.Sprintf("Color3s(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		fmtCall := fmt.Sprintf("Color3ub(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		fmtCall := fmt.Sprintf("Color3ui(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		fmtCall := fmt.Sprintf("Color3us(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		fmtCall := fmt.Sprintf("Color4b(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		fmtCall := fmt.Sprintf("Color4d(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		fmtCall := fmt.Sprintf("Color4f(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		fmtCall := fmt.Sprintf("Color4i(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		fmtCall := fmt.Sprintf("Color4s(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		fmtCall := fmt.Sprintf("Color4ub(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		fmtCall := fmt.Sprintf("Color4ui(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		fmtCall := fmt.Sprintf("Color4us(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v *int8) {
		fmtCall := fmt.Sprintf("Color3bv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color3dv = func(v *float64) {
		fmtCall := fmt.Sprintf("Color3dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color3fv = func(v *float32) {
		fmtCall := fmt.Sprintf("Color3fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color3iv = func(v *int32) {
		fmtCall := fmt.Sprintf("Color3iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color3sv = func(v *int16) {
		fmtCall := fmt.Sprintf("Color3sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color3ubv = func(v *uint8) {
		fmtCall := fmt.Sprintf("Color3ubv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color3uiv = func(v *uint32) {
		fmtCall := fmt.Sprintf("Color3uiv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color3usv = func(v *uint16) {
		fmtCall := fmt.Sprintf("Color3usv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.Color4bv = func(v *int8) {
		fmtCall := fmt.Sprintf("Color4bv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color4dv = func(v *float64) {
		fmtCall := fmt.Sprintf("Color4dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color4fv = func(v *float32) {
		fmtCall := fmt.Sprintf("Color4fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color4iv = func(v *int32) {
		fmtCall := fmt.Sprintf("Color4iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color4sv = func(v *int16) {
		fmtCall := fmt.Sprintf("Color4sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color4ubv = func(v *uint8) {
		fmtCall := fmt.Sprintf("Color4ubv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color4uiv = func(v *uint32) {
		fmtCall := fmt.Sprintf("Color4uiv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color4usv = func(v *uint16) {
		fmtCall := fmt.Sprintf("Color4usv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		fmtCall := fmt.Sprintf("ColorMask(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode Enum) {
		fmtCall := fmt.Sprintf("ColorMaterial(%v, %v)", face, mode)
		defer glc.trace(fmtCall)
		C.gl12ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type Enum) {
		fmtCall := fmt.Sprintf("CopyPixels(%v, %v, %v, %v, %v)", x, y, width, height, Type)
		defer glc.trace(fmtCall)
		C.gl12CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode Enum) {
		fmtCall := fmt.Sprintf("CullFace(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12CullFace(glc.context, C.GLenum(mode))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		fmtCall := fmt.Sprintf("DeleteLists(%v, %v)", list, Range)
		defer glc.trace(fmtCall)
		C.gl12DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func Enum) {
		fmtCall := fmt.Sprintf("DepthFunc(%v)", Func)
		defer glc.trace(fmtCall)
		C.gl12DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthMask = func(flag bool) {
		fmtCall := fmt.Sprintf("DepthMask(%v)", flag)
		defer glc.trace(fmtCall)
		C.gl12DepthMask(glc.context, boolToGL(flag))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		fmtCall := fmt.Sprintf("DepthRange(%v, %v)", zNear, zFar)
		defer glc.trace(fmtCall)
		C.gl12DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap Enum) {
		fmtCall := fmt.Sprintf("Enable(%v)", cap)
		defer glc.trace(fmtCall)
		C.gl12Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap Enum) {
		fmtCall := fmt.Sprintf("Disable(%v)", cap)
		defer glc.trace(fmtCall)
		C.gl12Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode Enum) {
		fmtCall := fmt.Sprintf("DrawBuffer(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type Enum, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("DrawPixels(%v, %v, %v, %v, %v)", width, height, format, Type, data)
		defer glc.trace(fmtCall)
		C.gl12DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.EdgeFlag = func(flag bool) {
		fmtCall := fmt.Sprintf("EdgeFlag(%v)", flag)
		defer glc.trace(fmtCall)
		C.gl12EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag *bool) {
		fmtCall := fmt.Sprintf("EdgeFlagv(%v)", flag)
		defer glc.trace(fmtCall)
		C.gl12EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(flag)))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("EdgeFlagPointer(%v, %v)", stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EvalCoord1d = func(u float64) {
		fmtCall := fmt.Sprintf("EvalCoord1d(%v)", u)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		fmtCall := fmt.Sprintf("EvalCoord1f(%v)", u)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		fmtCall := fmt.Sprintf("EvalCoord2d(%v, %v)", u, v)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		fmtCall := fmt.Sprintf("EvalCoord2f(%v, %v)", u, v)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u *float64) {
		fmtCall := fmt.Sprintf("EvalCoord1dv(%v)", u)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord1fv = func(u *float32) {
		fmtCall := fmt.Sprintf("EvalCoord1fv(%v)", u)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2dv = func(u *float64) {
		fmtCall := fmt.Sprintf("EvalCoord2dv(%v)", u)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2fv = func(u *float32) {
		fmtCall := fmt.Sprintf("EvalCoord2fv(%v)", u)
		defer glc.trace(fmtCall)
		C.gl12EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalMesh1 = func(mode Enum, i1, i2 int32) {
		fmtCall := fmt.Sprintf("EvalMesh1(%v, %v, %v)", mode, i1, i2)
		defer glc.trace(fmtCall)
		C.gl12EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode Enum, i1, i2, j1, j2 int32) {
		fmtCall := fmt.Sprintf("EvalMesh2(%v, %v, %v, %v, %v)", mode, i1, i2, j1, j2)
		defer glc.trace(fmtCall)
		C.gl12EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		fmtCall := fmt.Sprintf("EvalPoint1(%v)", i)
		defer glc.trace(fmtCall)
		C.gl12EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		fmtCall := fmt.Sprintf("EvalPoint2(%v, %v)", i, j)
		defer glc.trace(fmtCall)
		C.gl12EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type Enum, buffer *float32) {
		fmtCall := fmt.Sprintf("FeedbackBuffer(%v, %v, %v)", size, Type, buffer)
		defer glc.trace(fmtCall)
		C.gl12FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(buffer)))
	}

	glc.Finish = func() {
		fmtCall := "Finish()"
		defer glc.trace(fmtCall)
		C.gl12Finish(glc.context)
	}

	glc.Flush = func() {
		fmtCall := "Flush()"
		defer glc.trace(fmtCall)
		C.gl12Flush(glc.context)
	}

	glc.Fogf = func(pname Enum, param float32) {
		fmtCall := fmt.Sprintf("Fogf(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname Enum, param int32) {
		fmtCall := fmt.Sprintf("Fogi(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("Fogfv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Fogiv = func(pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("Fogiv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.FrontFace = func(mode Enum) {
		fmtCall := fmt.Sprintf("FrontFace(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		fmtCall := fmt.Sprintf("Frustum(%v, %v, %v, %v, %v, %v)", left, right, bottom, top, zNear, zFar)
		defer glc.trace(fmtCall)
		C.gl12Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		fmtCall := fmt.Sprintf("GenLists(%v)", Range)
		defer glc.trace(fmtCall)
		return uint32(C.gl12GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname Enum, params *bool) {
		fmtCall := fmt.Sprintf("GetBooleanv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(params)))
	}

	glc.GetDoublev = func(pname Enum, params *float64) {
		fmtCall := fmt.Sprintf("GetDoublev(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetFloatv = func(pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetFloatv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetIntegerv = func(pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetIntegerv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetClipPlane = func(plane Enum, equation *float64) {
		fmtCall := fmt.Sprintf("GetClipPlane(%v, %v)", plane, equation)
		defer glc.trace(fmtCall)
		C.gl12GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.GetError = func() Enum {
		return Enum(C.gl12GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetLightfv(%v, %v, %v)", light, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetLightiv = func(light, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetLightiv(%v, %v, %v)", light, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetMapdv = func(target, query Enum, v *float64) {
		fmtCall := fmt.Sprintf("GetMapdv(%v, %v, %v)", target, query, v)
		defer glc.trace(fmtCall)
		C.gl12GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.GetMapfv = func(target, query Enum, v *float32) {
		fmtCall := fmt.Sprintf("GetMapfv(%v, %v, %v)", target, query, v)
		defer glc.trace(fmtCall)
		C.gl12GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.GetMapiv = func(target, query Enum, v *int32) {
		fmtCall := fmt.Sprintf("GetMapiv(%v, %v, %v)", target, query, v)
		defer glc.trace(fmtCall)
		C.gl12GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.GetMaterialfv = func(face, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetMaterialfv(%v, %v, %v)", face, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetMaterialiv = func(face, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetMaterialiv(%v, %v, %v)", face, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetPixelMapfv = func(Map Enum, values *float32) {
		fmtCall := fmt.Sprintf("GetPixelMapfv(%v, %v)", Map, values)
		defer glc.trace(fmtCall)
		C.gl12GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapuiv = func(Map Enum, values *uint32) {
		fmtCall := fmt.Sprintf("GetPixelMapuiv(%v, %v)", Map, values)
		defer glc.trace(fmtCall)
		C.gl12GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapusv = func(Map Enum, values *uint16) {
		fmtCall := fmt.Sprintf("GetPixelMapusv(%v, %v)", Map, values)
		defer glc.trace(fmtCall)
		C.gl12GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.GetPolygonStipple = func(pattern *uint8) {
		fmtCall := fmt.Sprintf("GetPolygonStipple(%v)", pattern)
		defer glc.trace(fmtCall)
		C.gl12GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(pattern)))
	}

	glc.GetString = func(name Enum) string {
		fmtCall := fmt.Sprintf("GetString(%v)", name)
		defer glc.trace(fmtCall)
		cstr := C.gl12GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetTexEnvfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexEnviv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetTexEnviv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexGendv = func(coord, pname Enum, params *float64) {
		fmtCall := fmt.Sprintf("GetTexGendv(%v, %v, %v)", coord, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetTexGenfv = func(coord, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetTexGenfv(%v, %v, %v)", coord, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexGeniv = func(coord, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetTexGeniv(%v, %v, %v)", coord, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexImage = func(target Enum, level int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetTexImage(%v, %v, %v, %v, %v)", target, level, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target Enum, level int32, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetTexLevelParameterfv(%v, %v, %v, %v)", target, level, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexLevelParameteriv = func(target Enum, level int32, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetTexLevelParameteriv(%v, %v, %v, %v)", target, level, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexParameterfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetTexParameterfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexParameteriv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetTexParameteriv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Hint = func(target, mode Enum) {
		fmtCall := fmt.Sprintf("Hint(%v, %v)", target, mode)
		defer glc.trace(fmtCall)
		C.gl12Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		fmtCall := fmt.Sprintf("Indexd(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		fmtCall := fmt.Sprintf("Indexf(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		fmtCall := fmt.Sprintf("Indexi(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		fmtCall := fmt.Sprintf("Indexs(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexdv = func(c *float64) {
		fmtCall := fmt.Sprintf("Indexdv(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(c)))
	}

	glc.Indexfv = func(c *float32) {
		fmtCall := fmt.Sprintf("Indexfv(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(c)))
	}

	glc.Indexiv = func(c *int32) {
		fmtCall := fmt.Sprintf("Indexiv(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(c)))
	}

	glc.Indexsv = func(c *int16) {
		fmtCall := fmt.Sprintf("Indexsv(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(c)))
	}

	glc.IndexMask = func(mask uint32) {
		fmtCall := fmt.Sprintf("IndexMask(%v)", mask)
		defer glc.trace(fmtCall)
		C.gl12IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("IndexPointer(%v, %v, %v)", Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		fmtCall := "InitNames()"
		defer glc.trace(fmtCall)
		C.gl12InitNames(glc.context)
	}

	glc.IsEnabled = func(cap Enum) {
		fmtCall := fmt.Sprintf("IsEnabled(%v)", cap)
		defer glc.trace(fmtCall)
		C.gl12IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		fmtCall := fmt.Sprintf("IsList(%v)", list)
		defer glc.trace(fmtCall)
		return C.gl12IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname Enum, param float32) {
		fmtCall := fmt.Sprintf("Lightf(%v, %v, %v)", light, pname, param)
		defer glc.trace(fmtCall)
		C.gl12Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname Enum, param int32) {
		fmtCall := fmt.Sprintf("Lighti(%v, %v, %v)", light, pname, param)
		defer glc.trace(fmtCall)
		C.gl12Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("Lightfv(%v, %v, %v)", light, pname, params)
		defer glc.trace(fmtCall)
		C.gl12Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Lightiv = func(light, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("Lightiv(%v, %v, %v)", light, pname, params)
		defer glc.trace(fmtCall)
		C.gl12Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LightModelf = func(pname Enum, param float32) {
		fmtCall := fmt.Sprintf("LightModelf(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname Enum, param int32) {
		fmtCall := fmt.Sprintf("LightModeli(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("LightModelfv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.LightModeliv = func(pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("LightModeliv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		fmtCall := fmt.Sprintf("LineStipple(%v, %v)", factor, pattern)
		defer glc.trace(fmtCall)
		C.gl12LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		fmtCall := fmt.Sprintf("LineWidth(%v)", width)
		defer glc.trace(fmtCall)
		C.gl12LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		fmtCall := fmt.Sprintf("ListBase(%v)", base)
		defer glc.trace(fmtCall)
		C.gl12ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		fmtCall := "LoadIdentity()"
		defer glc.trace(fmtCall)
		C.gl12LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m *float64) {
		fmtCall := fmt.Sprintf("LoadMatrixd(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadMatrixf = func(m *float32) {
		fmtCall := fmt.Sprintf("LoadMatrixf(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.LoadName = func(name uint32) {
		fmtCall := fmt.Sprintf("LoadName(%v)", name)
		defer glc.trace(fmtCall)
		C.gl12LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode Enum) {
		fmtCall := fmt.Sprintf("LogicOp(%v)", opcode)
		defer glc.trace(fmtCall)
		C.gl12LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target Enum, u1, u2 float64, stride, order int32, points *float64) {
		fmtCall := fmt.Sprintf("Map1d(%v, %v, %v, %v, %v, %v)", target, u1, u2, stride, order, points)
		defer glc.trace(fmtCall)
		C.gl12Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map1f = func(target Enum, u1, u2 float32, stride, order int32, points *float32) {
		fmtCall := fmt.Sprintf("Map1f(%v, %v, %v, %v, %v, %v)", target, u1, u2, stride, order, points)
		defer glc.trace(fmtCall)
		C.gl12Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.Map2d = func(target Enum, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64) {
		fmtCall := fmt.Sprintf("Map2d(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points)
		defer glc.trace(fmtCall)
		C.gl12Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map2f = func(target Enum, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32) {
		fmtCall := fmt.Sprintf("Map2f(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points)
		defer glc.trace(fmtCall)
		C.gl12Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		fmtCall := fmt.Sprintf("MapGrid1d(%v, %v, %v)", un, u1, u2)
		defer glc.trace(fmtCall)
		C.gl12MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		fmtCall := fmt.Sprintf("MapGrid1f(%v, %v, %v)", un, u1, u2)
		defer glc.trace(fmtCall)
		C.gl12MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		fmtCall := fmt.Sprintf("MapGrid2d(%v, %v, %v, %v, %v, %v)", un, u1, u2, vn, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		fmtCall := fmt.Sprintf("MapGrid2f(%v, %v, %v, %v, %v, %v)", un, u1, u2, vn, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname Enum, param float32) {
		fmtCall := fmt.Sprintf("Materialf(%v, %v, %v)", face, pname, param)
		defer glc.trace(fmtCall)
		C.gl12Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname Enum, param int32) {
		fmtCall := fmt.Sprintf("Materiali(%v, %v, %v)", face, pname, param)
		defer glc.trace(fmtCall)
		C.gl12Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("Materialfv(%v, %v, %v)", face, pname, params)
		defer glc.trace(fmtCall)
		C.gl12Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Materialiv = func(face, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("Materialiv(%v, %v, %v)", face, pname, params)
		defer glc.trace(fmtCall)
		C.gl12Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.MatrixMode = func(mode Enum) {
		fmtCall := fmt.Sprintf("MatrixMode(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m *float64) {
		fmtCall := fmt.Sprintf("MultMatrixd(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultMatrixf = func(m *float32) {
		fmtCall := fmt.Sprintf("MultMatrixf(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.NewList = func(list uint32, mode Enum) {
		fmtCall := fmt.Sprintf("NewList(%v, %v)", list, mode)
		defer glc.trace(fmtCall)
		C.gl12NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		fmtCall := "EndList()"
		defer glc.trace(fmtCall)
		C.gl12EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		fmtCall := fmt.Sprintf("Normal3b(%v, %v, %v)", nx, ny, nz)
		defer glc.trace(fmtCall)
		C.gl12Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		fmtCall := fmt.Sprintf("Normal3d(%v, %v, %v)", nx, ny, nz)
		defer glc.trace(fmtCall)
		C.gl12Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		fmtCall := fmt.Sprintf("Normal3f(%v, %v, %v)", nx, ny, nz)
		defer glc.trace(fmtCall)
		C.gl12Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		fmtCall := fmt.Sprintf("Normal3i(%v, %v, %v)", nx, ny, nz)
		defer glc.trace(fmtCall)
		C.gl12Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		fmtCall := fmt.Sprintf("Normal3s(%v, %v, %v)", nx, ny, nz)
		defer glc.trace(fmtCall)
		C.gl12Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v *int8) {
		fmtCall := fmt.Sprintf("Normal3bv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Normal3dv = func(v *float64) {
		fmtCall := fmt.Sprintf("Normal3dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Normal3fv = func(v *float32) {
		fmtCall := fmt.Sprintf("Normal3fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Normal3iv = func(v *int32) {
		fmtCall := fmt.Sprintf("Normal3iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Normal3sv = func(v *int16) {
		fmtCall := fmt.Sprintf("Normal3sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		fmtCall := fmt.Sprintf("Ortho(%v, %v, %v, %v, %v, %v)", left, right, bottom, top, zNear, zfar)
		defer glc.trace(fmtCall)
		C.gl12Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		fmtCall := fmt.Sprintf("PassThrough(%v)", token)
		defer glc.trace(fmtCall)
		C.gl12PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map Enum, mapsize int32, values *float32) {
		fmtCall := fmt.Sprintf("PixelMapfv(%v, %v, %v)", Map, mapsize, values)
		defer glc.trace(fmtCall)
		C.gl12PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.PixelMapuiv = func(Map Enum, mapsize int32, values *uint32) {
		fmtCall := fmt.Sprintf("PixelMapuiv(%v, %v, %v)", Map, mapsize, values)
		defer glc.trace(fmtCall)
		C.gl12PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.PixelMapusv = func(Map Enum, mapsize int32, values *uint16) {
		fmtCall := fmt.Sprintf("PixelMapusv(%v, %v, %v)", Map, mapsize, values)
		defer glc.trace(fmtCall)
		C.gl12PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.PixelStoref = func(pname Enum, param float32) {
		fmtCall := fmt.Sprintf("PixelStoref(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname Enum, param int32) {
		fmtCall := fmt.Sprintf("PixelStorei(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname Enum, param float32) {
		fmtCall := fmt.Sprintf("PixelTransferf(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname Enum, param int32) {
		fmtCall := fmt.Sprintf("PixelTransferi(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		fmtCall := fmt.Sprintf("PixelZoom(%v, %v)", xfactor, yfactor)
		defer glc.trace(fmtCall)
		C.gl12PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		fmtCall := fmt.Sprintf("PointSize(%v)", size)
		defer glc.trace(fmtCall)
		C.gl12PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode Enum) {
		fmtCall := fmt.Sprintf("PolygonMode(%v, %v)", face, mode)
		defer glc.trace(fmtCall)
		C.gl12PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask *uint8) {
		fmtCall := fmt.Sprintf("PolygonStipple(%v)", mask)
		defer glc.trace(fmtCall)
		C.gl12PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(mask)))
	}

	glc.PushAttrib = func(mask uint32) {
		fmtCall := fmt.Sprintf("PushAttrib(%v)", mask)
		defer glc.trace(fmtCall)
		C.gl12PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		fmtCall := "PopAttrib()"
		defer glc.trace(fmtCall)
		C.gl12PopAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		fmtCall := "PushMatrix()"
		defer glc.trace(fmtCall)
		C.gl12PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		fmtCall := "PopMatrix()"
		defer glc.trace(fmtCall)
		C.gl12PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		fmtCall := fmt.Sprintf("PushName(%v)", name)
		defer glc.trace(fmtCall)
		C.gl12PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		fmtCall := "PopName()"
		defer glc.trace(fmtCall)
		C.gl12PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		fmtCall := fmt.Sprintf("RasterPos2d(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		fmtCall := fmt.Sprintf("RasterPos2f(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		fmtCall := fmt.Sprintf("RasterPos2i(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		fmtCall := fmt.Sprintf("RasterPos2s(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		fmtCall := fmt.Sprintf("RasterPos3d(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		fmtCall := fmt.Sprintf("RasterPos3f(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		fmtCall := fmt.Sprintf("RasterPos3i(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		fmtCall := fmt.Sprintf("RasterPos3s(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		fmtCall := fmt.Sprintf("RasterPos4d(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		fmtCall := fmt.Sprintf("RasterPos4f(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		fmtCall := fmt.Sprintf("RasterPos4i(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		fmtCall := fmt.Sprintf("RasterPos4s(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v *float64) {
		fmtCall := fmt.Sprintf("RasterPos2dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos2fv = func(v *float32) {
		fmtCall := fmt.Sprintf("RasterPos2fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos2iv = func(v *int32) {
		fmtCall := fmt.Sprintf("RasterPos2iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos2sv = func(v *int16) {
		fmtCall := fmt.Sprintf("RasterPos2sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos3dv = func(v *float64) {
		fmtCall := fmt.Sprintf("RasterPos3dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos3fv = func(v *float32) {
		fmtCall := fmt.Sprintf("RasterPos3fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos3iv = func(v *int32) {
		fmtCall := fmt.Sprintf("RasterPos3iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos3sv = func(v *int16) {
		fmtCall := fmt.Sprintf("RasterPos3sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos4dv = func(v *float64) {
		fmtCall := fmt.Sprintf("RasterPos4dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos4fv = func(v *float32) {
		fmtCall := fmt.Sprintf("RasterPos4fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos4iv = func(v *int32) {
		fmtCall := fmt.Sprintf("RasterPos4iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos4sv = func(v *int16) {
		fmtCall := fmt.Sprintf("RasterPos4sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.ReadBuffer = func(mode Enum) {
		fmtCall := fmt.Sprintf("ReadBuffer(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("ReadPixels(%v, %v, %v, %v, %v, %v, %v)", x, y, width, height, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		fmtCall := fmt.Sprintf("Rectd(%v, %v, %v, %v)", x1, y1, x2, y2)
		defer glc.trace(fmtCall)
		C.gl12Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		fmtCall := fmt.Sprintf("Rectf(%v, %v, %v, %v)", x1, y1, x2, y2)
		defer glc.trace(fmtCall)
		C.gl12Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		fmtCall := fmt.Sprintf("Recti(%v, %v, %v, %v)", x1, y1, x2, y2)
		defer glc.trace(fmtCall)
		C.gl12Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		fmtCall := fmt.Sprintf("Rects(%v, %v, %v, %v)", x1, y1, x2, y2)
		defer glc.trace(fmtCall)
		C.gl12Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 *float64) {
		fmtCall := fmt.Sprintf("Rectdv(%v, %v)", v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
	}

	glc.Rectfv = func(v1, v2 *float32) {
		fmtCall := fmt.Sprintf("Rectfv(%v, %v)", v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
	}

	glc.Rectiv = func(v1, v2 *int32) {
		fmtCall := fmt.Sprintf("Rectiv(%v, %v)", v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
	}

	glc.Rectsv = func(v1, v2 *int16) {
		fmtCall := fmt.Sprintf("Rectsv(%v, %v)", v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
	}

	glc.RenderMode = func(mode Enum) int32 {
		fmtCall := fmt.Sprintf("RenderMode(%v)", mode)
		defer glc.trace(fmtCall)
		return int32(C.gl12RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		fmtCall := fmt.Sprintf("Rotated(%v, %v, %v, %v)", angle, x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		fmtCall := fmt.Sprintf("Rotatef(%v, %v, %v, %v)", angle, x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		fmtCall := fmt.Sprintf("Scaled(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		fmtCall := fmt.Sprintf("Scalef(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		fmtCall := fmt.Sprintf("Scissor(%v, %v, %v, %v)", x, y, width, height)
		defer glc.trace(fmtCall)
		C.gl12Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer *uint32) {
		fmtCall := fmt.Sprintf("SelectBuffer(%v, %v)", size, buffer)
		defer glc.trace(fmtCall)
		C.gl12SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
	}

	glc.ShadeModel = func(mode Enum) {
		fmtCall := fmt.Sprintf("ShadeModel(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func Enum, ref int32, mask uint32) {
		fmtCall := fmt.Sprintf("StencilFunc(%v, %v, %v)", Func, ref, mask)
		defer glc.trace(fmtCall)
		C.gl12StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		fmtCall := fmt.Sprintf("StencilMask(%v)", mask)
		defer glc.trace(fmtCall)
		C.gl12StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass Enum) {
		fmtCall := fmt.Sprintf("StencilOp(%v, %v, %v)", fail, zfail, zpass)
		defer glc.trace(fmtCall)
		C.gl12StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		fmtCall := fmt.Sprintf("TexCoord1d(%v)", s)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		fmtCall := fmt.Sprintf("TexCoord1f(%v)", s)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		fmtCall := fmt.Sprintf("TexCoord1i(%v)", s)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		fmtCall := fmt.Sprintf("TexCoord1s(%v)", s)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		fmtCall := fmt.Sprintf("TexCoord2d(%v, %v)", s, t)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		fmtCall := fmt.Sprintf("TexCoord2f(%v, %v)", s, t)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		fmtCall := fmt.Sprintf("TexCoord2i(%v, %v)", s, t)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		fmtCall := fmt.Sprintf("TexCoord2s(%v, %v)", s, t)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		fmtCall := fmt.Sprintf("TexCoord3d(%v, %v, %v)", s, t, r)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		fmtCall := fmt.Sprintf("TexCoord3f(%v, %v, %v)", s, t, r)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		fmtCall := fmt.Sprintf("TexCoord3i(%v, %v, %v)", s, t, r)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		fmtCall := fmt.Sprintf("TexCoord3s(%v, %v, %v)", s, t, r)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		fmtCall := fmt.Sprintf("TexCoord4d(%v, %v, %v, %v)", s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		fmtCall := fmt.Sprintf("TexCoord4f(%v, %v, %v, %v)", s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		fmtCall := fmt.Sprintf("TexCoord4i(%v, %v, %v, %v)", s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		fmtCall := fmt.Sprintf("TexCoord4s(%v, %v, %v, %v)", s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v *float64) {
		fmtCall := fmt.Sprintf("TexCoord1dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord1fv = func(v *float32) {
		fmtCall := fmt.Sprintf("TexCoord1fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord1iv = func(v *int32) {
		fmtCall := fmt.Sprintf("TexCoord1iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord1sv = func(v *int16) {
		fmtCall := fmt.Sprintf("TexCoord1sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord2dv = func(v *float64) {
		fmtCall := fmt.Sprintf("TexCoord2dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord2fv = func(v *float32) {
		fmtCall := fmt.Sprintf("TexCoord2fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord2iv = func(v *int32) {
		fmtCall := fmt.Sprintf("TexCoord2iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord2sv = func(v *int16) {
		fmtCall := fmt.Sprintf("TexCoord2sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord3dv = func(v *float64) {
		fmtCall := fmt.Sprintf("TexCoord3dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord3fv = func(v *float32) {
		fmtCall := fmt.Sprintf("TexCoord3fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord3iv = func(v *int32) {
		fmtCall := fmt.Sprintf("TexCoord3iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord3sv = func(v *int16) {
		fmtCall := fmt.Sprintf("TexCoord3sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord4dv = func(v *float64) {
		fmtCall := fmt.Sprintf("TexCoord4dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord4fv = func(v *float32) {
		fmtCall := fmt.Sprintf("TexCoord4fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord4iv = func(v *int32) {
		fmtCall := fmt.Sprintf("TexCoord4iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord4sv = func(v *int16) {
		fmtCall := fmt.Sprintf("TexCoord4sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexEnvf = func(target, pname Enum, param float32) {
		fmtCall := fmt.Sprintf("TexEnvf(%v, %v, %v)", target, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname Enum, param int32) {
		fmtCall := fmt.Sprintf("TexEnvi(%v, %v, %v)", target, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("TexEnvfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexEnviv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("TexEnviv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexGend = func(coord, pname Enum, param float64) {
		fmtCall := fmt.Sprintf("TexGend(%v, %v, %v)", coord, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname Enum, param float32) {
		fmtCall := fmt.Sprintf("TexGenf(%v, %v, %v)", coord, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname Enum, param int32) {
		fmtCall := fmt.Sprintf("TexGeni(%v, %v, %v)", coord, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname Enum, params *float64) {
		fmtCall := fmt.Sprintf("TexGendv(%v, %v, %v)", coord, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.TexGenfv = func(coord, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("TexGenfv(%v, %v, %v)", coord, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexGeniv = func(coord, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("TexGeniv(%v, %v, %v)", coord, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexImage1D = func(target Enum, level, internalformat int32, width int32, border int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexImage1D(%v, %v, %v, %v, %v, %v, %v, %v)", target, level, internalformat, width, border, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target Enum, level, internalformat int32, width, height int32, border int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexImage2D(%v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, internalformat, width, height, border, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname Enum, param float32) {
		fmtCall := fmt.Sprintf("TexParameterf(%v, %v, %v)", target, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname Enum, param int32) {
		fmtCall := fmt.Sprintf("TexParameteri(%v, %v, %v)", target, pname, param)
		defer glc.trace(fmtCall)
		C.gl12TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("TexParameterfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexParameteriv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("TexParameteriv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Translated = func(x, y, z float64) {
		fmtCall := fmt.Sprintf("Translated(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		fmtCall := fmt.Sprintf("Translatef(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		fmtCall := fmt.Sprintf("Vertex2s(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		fmtCall := fmt.Sprintf("Vertex2i(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		fmtCall := fmt.Sprintf("Vertex2f(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		fmtCall := fmt.Sprintf("Vertex2d(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		fmtCall := fmt.Sprintf("Vertex3s(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		fmtCall := fmt.Sprintf("Vertex3i(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		fmtCall := fmt.Sprintf("Vertex3f(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		fmtCall := fmt.Sprintf("Vertex3d(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		fmtCall := fmt.Sprintf("Vertex4s(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		fmtCall := fmt.Sprintf("Vertex4i(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		fmtCall := fmt.Sprintf("Vertex4f(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		fmtCall := fmt.Sprintf("Vertex4d(%v, %v, %v, %v)", x, y, z, w)
		defer glc.trace(fmtCall)
		C.gl12Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		fmtCall := fmt.Sprintf("Viewport(%v, %v, %v, %v)", x, y, width, height)
		defer glc.trace(fmtCall)
		C.gl12Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetConvolutionParameterfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetConvolutionParameterfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetConvolutionParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionParameteriv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetConvolutionParameteriv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetConvolutionParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		fmtCall := fmt.Sprintf("AreTexturesResident(%v)", textures)
		defer glc.trace(fmtCall)
		var cRes *C.GLboolean
		status = C.gl12AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		fmtCall := fmt.Sprintf("ArrayElement(%v)", i)
		defer glc.trace(fmtCall)
		C.gl12ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode Enum, first int32, count int32) {
		fmtCall := fmt.Sprintf("DrawArrays(%v, %v, %v)", mode, first, count)
		defer glc.trace(fmtCall)
		C.gl12DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode Enum, count int32, Type Enum, indices unsafe.Pointer) {
		fmtCall := fmt.Sprintf("DrawElements(%v, %v, %v, %v)", mode, count, Type, indices)
		defer glc.trace(fmtCall)
		C.gl12DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname Enum, params unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetPointerv(%v, %v)", pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		fmtCall := fmt.Sprintf("PolygonOffset(%v, %v)", factor, units)
		defer glc.trace(fmtCall)
		C.gl12PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target Enum, level int32, internalFormat Enum, x, y int32, width int32, border int32) {
		fmtCall := fmt.Sprintf("CopyTexImage1D(%v, %v, %v, %v, %v, %v, %v)", target, level, internalFormat, x, y, width, border)
		defer glc.trace(fmtCall)
		C.gl12CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target Enum, level int32, internalFormat Enum, x, y int32, width, height int32, border int32) {
		fmtCall := fmt.Sprintf("CopyTexImage2D(%v, %v, %v, %v, %v, %v, %v, %v)", target, level, internalFormat, x, y, width, height, border)
		defer glc.trace(fmtCall)
		C.gl12CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target Enum, level, xoffset, x, y int32, width int32) {
		fmtCall := fmt.Sprintf("CopyTexSubImage1D(%v, %v, %v, %v, %v, %v)", target, level, xoffset, x, y, width)
		defer glc.trace(fmtCall)
		C.gl12CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target Enum, level, xoffset, yoffset, x, y int32, width, height int32) {
		fmtCall := fmt.Sprintf("CopyTexSubImage2D(%v, %v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, yoffset, x, y, width, height)
		defer glc.trace(fmtCall)
		C.gl12CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target Enum, texture uint32) {
		fmtCall := fmt.Sprintf("BindTexture(%v, %v)", target, texture)
		defer glc.trace(fmtCall)
		C.gl12BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures *uint32) {
		fmtCall := fmt.Sprintf("DeleteTextures(%v, %v)", n, textures)
		defer glc.trace(fmtCall)
		C.gl12DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.GenTextures = func(n int32, textures *uint32) {
		fmtCall := fmt.Sprintf("GenTextures(%v, %v)", n, textures)
		defer glc.trace(fmtCall)
		C.gl12GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.IsTexture = func(texture uint32) bool {
		fmtCall := fmt.Sprintf("IsTexture(%v)", texture)
		defer glc.trace(fmtCall)
		return C.gl12IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("ColorPointer(%v, %v, %v, %v)", size, Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap Enum) {
		fmtCall := fmt.Sprintf("EnableClientState(%v)", cap)
		defer glc.trace(fmtCall)
		C.gl12EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap Enum) {
		fmtCall := fmt.Sprintf("DisableClientState(%v)", cap)
		defer glc.trace(fmtCall)
		C.gl12DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.Indexub = func(c uint8) {
		fmtCall := fmt.Sprintf("Indexub(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexubv = func(c *uint8) {
		fmtCall := fmt.Sprintf("Indexubv(%v)", c)
		defer glc.trace(fmtCall)
		C.gl12Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(c)))
	}

	glc.InterleavedArrays = func(format Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("InterleavedArrays(%v, %v, %v)", format, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.NormalPointer = func(Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("NormalPointer(%v, %v, %v)", Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.PushClientAttrib = func(mask uint32) {
		fmtCall := fmt.Sprintf("PushClientAttrib(%v)", mask)
		defer glc.trace(fmtCall)
		C.gl12PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PrioritizeTextures = func(n int32, textures *uint32, priorities *float32) {
		fmtCall := fmt.Sprintf("PrioritizeTextures(%v, %v, %v)", n, textures, priorities)
		defer glc.trace(fmtCall)
		C.gl12PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLclampf)(unsafe.Pointer(priorities)))
	}

	glc.PopClientAttrib = func() {
		fmtCall := "PopClientAttrib()"
		defer glc.trace(fmtCall)
		C.gl12PopClientAttrib(glc.context)
	}

	glc.TexCoordPointer = func(size int32, Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexCoordPointer(%v, %v, %v, %v)", size, Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexSubImage1D = func(target Enum, level, xoffset int32, width int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexSubImage1D(%v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, width, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target Enum, level, xoffset, yoffset int32, width, height int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexSubImage2D(%v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, yoffset, width, height, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.VertexPointer = func(size int32, Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("VertexPointer(%v, %v, %v, %v)", size, Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.ColorTable = func(target, internalformat Enum, width int32, format, Type Enum, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("ColorTable(%v, %v, %v, %v, %v, %v)", target, internalformat, width, format, Type, data)
		defer glc.trace(fmtCall)
		C.gl12ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("ColorTableParameterfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.ColorTableParameteriv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("ColorTableParameteriv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorSubTable = func(target Enum, start, count int32, format, Type Enum, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("ColorSubTable(%v, %v, %v, %v, %v, %v)", target, start, count, format, Type, data)
		defer glc.trace(fmtCall)
		C.gl12ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter1D = func(target, internalformat Enum, width int32, format, Type Enum, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("ConvolutionFilter1D(%v, %v, %v, %v, %v, %v)", target, internalformat, width, format, Type, data)
		defer glc.trace(fmtCall)
		C.gl12ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat Enum, width, height int32, format, Type Enum, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("ConvolutionFilter2D(%v, %v, %v, %v, %v, %v, %v)", target, internalformat, width, height, format, Type, data)
		defer glc.trace(fmtCall)
		C.gl12ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname Enum, params float32) {
		fmtCall := fmt.Sprintf("ConvolutionParameterf(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname Enum, params int32) {
		fmtCall := fmt.Sprintf("ConvolutionParameteri(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat Enum, x, y int32, width int32) {
		fmtCall := fmt.Sprintf("CopyColorTable(%v, %v, %v, %v, %v)", target, internalformat, x, y, width)
		defer glc.trace(fmtCall)
		C.gl12CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target Enum, start int32, x, y int32, width int32) {
		fmtCall := fmt.Sprintf("CopyColorSubTable(%v, %v, %v, %v, %v)", target, start, x, y, width)
		defer glc.trace(fmtCall)
		C.gl12CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat Enum, x, y int32, width int32) {
		fmtCall := fmt.Sprintf("CopyConvolutionFilter1D(%v, %v, %v, %v, %v)", target, internalformat, x, y, width)
		defer glc.trace(fmtCall)
		C.gl12CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat Enum, x, y int32, width, height int32) {
		fmtCall := fmt.Sprintf("CopyConvolutionFilter2D(%v, %v, %v, %v, %v, %v)", target, internalformat, x, y, width, height)
		defer glc.trace(fmtCall)
		C.gl12CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetColorTable = func(target, format, Type Enum, table unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetColorTable(%v, %v, %v, %v)", target, format, Type, table)
		defer glc.trace(fmtCall)
		C.gl12GetColorTable(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), table)
	}

	glc.GetColorTableParameterfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetColorTableParameterfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetColorTableParameteriv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetColorTableParameteriv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionFilter = func(target, format, Type Enum, image unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetConvolutionFilter(%v, %v, %v, %v)", target, format, Type, image)
		defer glc.trace(fmtCall)
		C.gl12GetConvolutionFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), image)
	}

	glc.GetHistogram = func(target Enum, reset bool, format, Type Enum, values unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetHistogram(%v, %v, %v, %v, %v)", target, reset, format, Type, values)
		defer glc.trace(fmtCall)
		C.gl12GetHistogram(glc.context, C.GLenum(target), boolToGL(reset), C.GLenum(format), C.GLenum(Type), values)
	}

	glc.GetHistogramParameterfv = func(target, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetHistogramParameterfv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetHistogramParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetHistogramParameteriv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetHistogramParameteriv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetHistogramParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetSeparableFilter = func(target, format, Type Enum, row, column, span unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetSeparableFilter(%v, %v, %v, %v, %v, %v)", target, format, Type, row, column, span)
		defer glc.trace(fmtCall)
		C.gl12GetSeparableFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), row, column, span)
	}

	glc.Histogram = func(target Enum, width int32, internalformat Enum, sink bool) {
		fmtCall := fmt.Sprintf("Histogram(%v, %v, %v, %v)", target, width, internalformat, sink)
		defer glc.trace(fmtCall)
		C.gl12Histogram(glc.context, C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.Minmax = func(target, internalformat Enum, sink bool) {
		fmtCall := fmt.Sprintf("Minmax(%v, %v, %v)", target, internalformat, sink)
		defer glc.trace(fmtCall)
		C.gl12Minmax(glc.context, C.GLenum(target), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.MultiTexCoord1s = func(target Enum, s int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord1s(%v, %v)", target, s)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1s(glc.context, C.GLenum(target), C.GLshort(s))
	}

	glc.MultiTexCoord1i = func(target Enum, s int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord1i(%v, %v)", target, s)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1i(glc.context, C.GLenum(target), C.GLint(s))
	}

	glc.MultiTexCoord1f = func(target Enum, s float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord1f(%v, %v)", target, s)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1f(glc.context, C.GLenum(target), C.GLfloat(s))
	}

	glc.MultiTexCoord1d = func(target Enum, s float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord1d(%v, %v)", target, s)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1d(glc.context, C.GLenum(target), C.GLdouble(s))
	}

	glc.MultiTexCoord2s = func(target Enum, s, t int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord2s(%v, %v, %v)", target, s, t)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t))
	}

	glc.MultiTexCoord2i = func(target Enum, s, t int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord2i(%v, %v, %v)", target, s, t)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t))
	}

	glc.MultiTexCoord2f = func(target Enum, s, t float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord2f(%v, %v, %v)", target, s, t)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
	}

	glc.MultiTexCoord2d = func(target Enum, s, t float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord2d(%v, %v, %v)", target, s, t)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
	}

	glc.MultiTexCoord3s = func(target Enum, s, t, r int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord3s(%v, %v, %v, %v)", target, s, t, r)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.MultiTexCoord3i = func(target Enum, s, t, r int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord3i(%v, %v, %v, %v)", target, s, t, r)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.MultiTexCoord3f = func(target Enum, s, t, r float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord3f(%v, %v, %v, %v)", target, s, t, r)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.MultiTexCoord3d = func(target Enum, s, t, r float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord3d(%v, %v, %v, %v)", target, s, t, r)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.MultiTexCoord4s = func(target Enum, s, t, r, q int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord4s(%v, %v, %v, %v, %v)", target, s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.MultiTexCoord4i = func(target Enum, s, t, r, q int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord4i(%v, %v, %v, %v, %v)", target, s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.MultiTexCoord4f = func(target Enum, s, t, r, q float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord4f(%v, %v, %v, %v, %v)", target, s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.MultiTexCoord4d = func(target Enum, s, t, r, q float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord4d(%v, %v, %v, %v, %v)", target, s, t, r, q)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.MultiTexCoord1sv = func(target Enum, v *int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord1sv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1iv = func(target Enum, v *int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord1iv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1fv = func(target Enum, v *float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord1fv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1dv = func(target Enum, v *float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord1dv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord1dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2sv = func(target Enum, v *int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord2sv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2iv = func(target Enum, v *int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord2iv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2fv = func(target Enum, v *float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord2fv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2dv = func(target Enum, v *float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord2dv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord2dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3sv = func(target Enum, v *int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord3sv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3iv = func(target Enum, v *int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord3iv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3fv = func(target Enum, v *float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord3fv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3dv = func(target Enum, v *float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord3dv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord3dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4sv = func(target Enum, v *int16) {
		fmtCall := fmt.Sprintf("MultiTexCoord4sv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4iv = func(target Enum, v *int32) {
		fmtCall := fmt.Sprintf("MultiTexCoord4iv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4fv = func(target Enum, v *float32) {
		fmtCall := fmt.Sprintf("MultiTexCoord4fv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4dv = func(target Enum, v *float64) {
		fmtCall := fmt.Sprintf("MultiTexCoord4dv(%v, %v)", target, v)
		defer glc.trace(fmtCall)
		C.gl12MultiTexCoord4dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.ResetHistogram = func(target Enum) {
		fmtCall := fmt.Sprintf("ResetHistogram(%v)", target)
		defer glc.trace(fmtCall)
		C.gl12ResetHistogram(glc.context, C.GLenum(target))
	}

	glc.ResetMinmax = func(target Enum) {
		fmtCall := fmt.Sprintf("ResetMinmax(%v)", target)
		defer glc.trace(fmtCall)
		C.gl12ResetMinmax(glc.context, C.GLenum(target))
	}

	glc.SeparableFilter2D = func(target, internalformat Enum, width, height int32, format, Type Enum, row, column unsafe.Pointer) {
		fmtCall := fmt.Sprintf("SeparableFilter2D(%v, %v, %v, %v, %v, %v, %v, %v)", target, internalformat, width, height, format, Type, row, column)
		defer glc.trace(fmtCall)
		C.gl12SeparableFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), row, column)
	}

	glc.BlendColor = func(red, green, blue, alpha float32) {
		fmtCall := fmt.Sprintf("BlendColor(%v, %v, %v, %v)", red, green, blue, alpha)
		defer glc.trace(fmtCall)
		C.gl12BlendColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode Enum) {
		fmtCall := fmt.Sprintf("BlendEquation(%v)", mode)
		defer glc.trace(fmtCall)
		C.gl12BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target Enum, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		fmtCall := fmt.Sprintf("CopyTexSubImage3D(%v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, yoffset, zoffset, x, y, width, height)
		defer glc.trace(fmtCall)
		C.gl12CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DrawRangeElements = func(mode Enum, start, end uint32, count int32, Type Enum, indices unsafe.Pointer) {
		fmtCall := fmt.Sprintf("DrawRangeElements(%v, %v, %v, %v, %v, %v)", mode, start, end, count, Type, indices)
		defer glc.trace(fmtCall)
		C.gl12DrawRangeElements(glc.context, C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.TexImage3D = func(target Enum, level, internalformat int32, width, height, depth int32, border int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexImage3D(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, internalformat, width, height, depth, border, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12TexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3D = func(target Enum, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type Enum, pixels unsafe.Pointer) {
		fmtCall := fmt.Sprintf("TexSubImage3D(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, yoffset, zoffset, width, height, depth, format, Type, pixels)
		defer glc.trace(fmtCall)
		C.gl12TexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.ActiveTexture = func(texture Enum) {
		fmtCall := fmt.Sprintf("ActiveTexture(%v)", texture)
		defer glc.trace(fmtCall)
		C.gl12ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture Enum) {
		fmtCall := fmt.Sprintf("ClientActiveTexture(%v)", texture)
		defer glc.trace(fmtCall)
		C.gl12ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target Enum, level int32, internalformat Enum, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CompressedTexImage1D(%v, %v, %v, %v, %v, %v, %v)", target, level, internalformat, width, border, imageSize, data)
		defer glc.trace(fmtCall)
		C.gl12CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target Enum, level int32, internalformat Enum, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CompressedTexImage2D(%v, %v, %v, %v, %v, %v, %v, %v)", target, level, internalformat, width, height, border, imageSize, data)
		defer glc.trace(fmtCall)
		C.gl12CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target Enum, level int32, internalformat Enum, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CompressedTexImage3D(%v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, internalformat, width, height, depth, border, imageSize, data)
		defer glc.trace(fmtCall)
		C.gl12CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target Enum, level, xoffset int32, width int32, format Enum, imageSize int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CompressedTexSubImage1D(%v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, width, format, imageSize, data)
		defer glc.trace(fmtCall)
		C.gl12CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target Enum, level, xoffset, yoffset int32, width, height int32, format Enum, imageSize int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CompressedTexSubImage2D(%v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, yoffset, width, height, format, imageSize, data)
		defer glc.trace(fmtCall)
		C.gl12CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target Enum, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format Enum, imageSize int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("CompressedTexSubImage3D(%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v)", target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data)
		defer glc.trace(fmtCall)
		C.gl12CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.GetCompressedTexImage = func(target Enum, lod int32, img unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetCompressedTexImage(%v, %v, %v)", target, lod, img)
		defer glc.trace(fmtCall)
		C.gl12GetCompressedTexImage(glc.context, C.GLenum(target), C.GLint(lod), img)
	}

	glc.LoadTransposeMatrixd = func(m *float64) {
		fmtCall := fmt.Sprintf("LoadTransposeMatrixd(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12LoadTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadTransposeMatrixf = func(m *float64) {
		fmtCall := fmt.Sprintf("LoadTransposeMatrixf(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12LoadTransposeMatrixf(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixd = func(m *float64) {
		fmtCall := fmt.Sprintf("MultTransposeMatrixd(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12MultTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixf = func(m *float32) {
		fmtCall := fmt.Sprintf("MultTransposeMatrixf(%v)", m)
		defer glc.trace(fmtCall)
		C.gl12MultTransposeMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.SampleCoverage = func(value float32, invert bool) {
		fmtCall := fmt.Sprintf("SampleCoverage(%v, %v)", value, invert)
		defer glc.trace(fmtCall)
		C.gl12SampleCoverage(glc.context, C.GLclampf(value), boolToGL(invert))
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha Enum) {
		fmtCall := fmt.Sprintf("BlendFuncSeparate(%v, %v, %v, %v)", srcRGB, dstRGB, srcAlpha, dstAlpha)
		defer glc.trace(fmtCall)
		C.gl12BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.FogCoordPointer = func(Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("FogCoordPointer(%v, %v, %v)", Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12FogCoordPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.FogCoordd = func(coord float64) {
		fmtCall := fmt.Sprintf("FogCoordd(%v)", coord)
		defer glc.trace(fmtCall)
		C.gl12FogCoordd(glc.context, C.GLdouble(coord))
	}

	glc.FogCoordf = func(coord float32) {
		fmtCall := fmt.Sprintf("FogCoordf(%v)", coord)
		defer glc.trace(fmtCall)
		C.gl12FogCoordf(glc.context, C.GLfloat(coord))
	}

	glc.FogCoorddv = func(coord *float64) {
		fmtCall := fmt.Sprintf("FogCoorddv(%v)", coord)
		defer glc.trace(fmtCall)
		C.gl12FogCoorddv(glc.context, (*C.GLdouble)(unsafe.Pointer(coord)))
	}

	glc.FogCoordfv = func(coord *float32) {
		fmtCall := fmt.Sprintf("FogCoordfv(%v)", coord)
		defer glc.trace(fmtCall)
		C.gl12FogCoordfv(glc.context, (*C.GLfloat)(unsafe.Pointer(coord)))
	}

	glc.MultiDrawArrays = func(mode Enum, first *int32, count *int32, primcount int32) {
		fmtCall := fmt.Sprintf("MultiDrawArrays(%v, %v, %v, %v)", mode, first, count, primcount)
		defer glc.trace(fmtCall)
		C.gl12MultiDrawArrays(glc.context, C.GLenum(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), C.GLsizei(primcount))
	}

	glc.MultiDrawElements = func(mode Enum, count *int32, Type Enum, indices unsafe.Pointer, primcount int32) {
		fmtCall := fmt.Sprintf("MultiDrawElements(%v, %v, %v, %v, %v)", mode, count, Type, indices, primcount)
		defer glc.trace(fmtCall)
		C.gl12MultiDrawElements(glc.context, C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(count)), C.GLenum(Type), indices, C.GLsizei(primcount))
	}

	glc.PointParameterf = func(pname Enum, param float32) {
		fmtCall := fmt.Sprintf("PointParameterf(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12PointParameterf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PointParameteri = func(pname Enum, param int32) {
		fmtCall := fmt.Sprintf("PointParameteri(%v, %v)", pname, param)
		defer glc.trace(fmtCall)
		C.gl12PointParameteri(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.SecondaryColor3b = func(red, green, blue int8) {
		fmtCall := fmt.Sprintf("SecondaryColor3b(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.SecondaryColor3s = func(red, green, blue int16) {
		fmtCall := fmt.Sprintf("SecondaryColor3s(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.SecondaryColor3i = func(red, green, blue int32) {
		fmtCall := fmt.Sprintf("SecondaryColor3i(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.SecondaryColor3f = func(red, green, blue float32) {
		fmtCall := fmt.Sprintf("SecondaryColor3f(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.SecondaryColor3d = func(red, green, blue float64) {
		fmtCall := fmt.Sprintf("SecondaryColor3d(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.SecondaryColor3ub = func(red, green, blue uint8) {
		fmtCall := fmt.Sprintf("SecondaryColor3ub(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.SecondaryColor3us = func(red, green, blue uint16) {
		fmtCall := fmt.Sprintf("SecondaryColor3us(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.SecondaryColor3ui = func(red, green, blue uint32) {
		fmtCall := fmt.Sprintf("SecondaryColor3ui(%v, %v, %v)", red, green, blue)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.SecondaryColor3bv = func(v *int8) {
		fmtCall := fmt.Sprintf("SecondaryColor3bv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3sv = func(v *int16) {
		fmtCall := fmt.Sprintf("SecondaryColor3sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3iv = func(v *int32) {
		fmtCall := fmt.Sprintf("SecondaryColor3iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3fv = func(v *float32) {
		fmtCall := fmt.Sprintf("SecondaryColor3fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3dv = func(v *float64) {
		fmtCall := fmt.Sprintf("SecondaryColor3dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3ubv = func(v *uint8) {
		fmtCall := fmt.Sprintf("SecondaryColor3ubv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3usv = func(v *uint16) {
		fmtCall := fmt.Sprintf("SecondaryColor3usv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3uiv = func(v *uint32) {
		fmtCall := fmt.Sprintf("SecondaryColor3uiv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColor3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColorPointer = func(size int32, Type Enum, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("SecondaryColorPointer(%v, %v, %v, %v)", size, Type, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12SecondaryColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.WindowPos2s = func(x, y int16) {
		fmtCall := fmt.Sprintf("WindowPos2s(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.WindowPos2i = func(x, y int32) {
		fmtCall := fmt.Sprintf("WindowPos2i(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.WindowPos2f = func(x, y float32) {
		fmtCall := fmt.Sprintf("WindowPos2f(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.WindowPos2d = func(x, y float64) {
		fmtCall := fmt.Sprintf("WindowPos2d(%v, %v)", x, y)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.WindowPos3s = func(x, y, z int16) {
		fmtCall := fmt.Sprintf("WindowPos3s(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.WindowPos3i = func(x, y, z int32) {
		fmtCall := fmt.Sprintf("WindowPos3i(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.WindowPos3f = func(x, y, z float32) {
		fmtCall := fmt.Sprintf("WindowPos3f(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.WindowPos3d = func(x, y, z float64) {
		fmtCall := fmt.Sprintf("WindowPos3d(%v, %v, %v)", x, y, z)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.WindowPos2sv = func(v *int16) {
		fmtCall := fmt.Sprintf("WindowPos2sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos2iv = func(v *int32) {
		fmtCall := fmt.Sprintf("WindowPos2iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos2fv = func(v *float32) {
		fmtCall := fmt.Sprintf("WindowPos2fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos2dv = func(v *float64) {
		fmtCall := fmt.Sprintf("WindowPos2dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.WindowPos3sv = func(v *int16) {
		fmtCall := fmt.Sprintf("WindowPos3sv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos3iv = func(v *int32) {
		fmtCall := fmt.Sprintf("WindowPos3iv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos3fv = func(v *float32) {
		fmtCall := fmt.Sprintf("WindowPos3fv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos3dv = func(v *float64) {
		fmtCall := fmt.Sprintf("WindowPos3dv(%v)", v)
		defer glc.trace(fmtCall)
		C.gl12WindowPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.BeginQuery = func(target Enum, id uint32) {
		fmtCall := fmt.Sprintf("BeginQuery(%v, %v)", target, id)
		defer glc.trace(fmtCall)
		C.gl12BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target Enum, buffer uint32) {
		fmtCall := fmt.Sprintf("BindBuffer(%v, %v)", target, buffer)
		defer glc.trace(fmtCall)
		C.gl12BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target Enum, size int32, data unsafe.Pointer, usage Enum) {
		fmtCall := fmt.Sprintf("BufferData(%v, %v, %v, %v)", target, size, data, usage)
		defer glc.trace(fmtCall)
		C.gl12BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset Enum, size int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("BufferSubData(%v, %v, %v, %v)", target, offset, size, data)
		defer glc.trace(fmtCall)
		C.gl12BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.DeleteBuffers = func(n int32, buffers *uint32) {
		fmtCall := fmt.Sprintf("DeleteBuffers(%v, %v)", n, buffers)
		defer glc.trace(fmtCall)
		C.gl12DeleteBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.DeleteQueries = func(n int32, ids *uint32) {
		fmtCall := fmt.Sprintf("DeleteQueries(%v, %v)", n, ids)
		defer glc.trace(fmtCall)
		C.gl12DeleteQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GenBuffers = func(n int32, buffers *uint32) {
		fmtCall := fmt.Sprintf("GenBuffers(%v, %v)", n, buffers)
		defer glc.trace(fmtCall)
		C.gl12GenBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.GenQueries = func(n int32, ids *uint32) {
		fmtCall := fmt.Sprintf("GenQueries(%v, %v)", n, ids)
		defer glc.trace(fmtCall)
		C.gl12GenQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GetBufferParameteriv = func(target, value Enum, data *int32) {
		fmtCall := fmt.Sprintf("GetBufferParameteriv(%v, %v, %v)", target, value, data)
		defer glc.trace(fmtCall)
		C.gl12GetBufferParameteriv(glc.context, C.GLenum(target), C.GLenum(value), (*C.GLint)(unsafe.Pointer(data)))
	}

	glc.GetBufferPointerv = func(target, pname Enum, params unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetBufferPointerv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetBufferPointerv(glc.context, C.GLenum(target), C.GLenum(pname), params)
	}

	glc.GetBufferSubData = func(target Enum, offset int32, size int32, data unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetBufferSubData(%v, %v, %v, %v)", target, offset, size, data)
		defer glc.trace(fmtCall)
		C.gl12GetBufferSubData(glc.context, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}

	glc.GetQueryObjectiv = func(id uint32, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetQueryObjectiv(%v, %v, %v)", id, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetQueryObjectiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetQueryObjectuiv = func(id uint32, pname Enum, params *uint32) {
		fmtCall := fmt.Sprintf("GetQueryObjectuiv(%v, %v, %v)", id, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetQueryObjectuiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(params)))
	}

	glc.GetQueryiv = func(target, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetQueryiv(%v, %v, %v)", target, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetQueryiv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.IsBuffer = func(buffer uint32) bool {
		fmtCall := fmt.Sprintf("IsBuffer(%v)", buffer)
		defer glc.trace(fmtCall)
		return C.gl12IsBuffer(glc.context, C.GLuint(buffer)) != 0
	}

	glc.IsQuery = func(id uint32) bool {
		fmtCall := fmt.Sprintf("IsQuery(%v)", id)
		defer glc.trace(fmtCall)
		return C.gl12IsQuery(glc.context, C.GLuint(id)) != 0
	}

	glc.MapBuffer = func(target, access Enum) unsafe.Pointer {
		fmtCall := fmt.Sprintf("MapBuffer(%v, %v)", target, access)
		defer glc.trace(fmtCall)
		return unsafe.Pointer(C.gl12MapBuffer(glc.context, C.GLenum(target), C.GLenum(access)))
	}

	glc.UnmapBuffer = func(target Enum) bool {
		fmtCall := fmt.Sprintf("UnmapBuffer(%v)", target)
		defer glc.trace(fmtCall)
		return C.gl12UnmapBuffer(glc.context, C.GLenum(target)) != 0
	}

	glc.AttachShader = func(program, shader uint32) {
		fmtCall := fmt.Sprintf("AttachShader(%v, %v)", program, shader)
		defer glc.trace(fmtCall)
		C.gl12AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		fmtCall := fmt.Sprintf("BindAttribLocation(%v, %v, %v)", program, index, name)
		defer glc.trace(fmtCall)
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl12BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha Enum) {
		fmtCall := fmt.Sprintf("BlendEquationSeperate(%v, %v)", modeRGB, modeAlpha)
		defer glc.trace(fmtCall)
		C.gl12BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		fmtCall := fmt.Sprintf("CompileShader(%v)", shader)
		defer glc.trace(fmtCall)
		C.gl12CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		fmtCall := "CreateProgram()"
		defer glc.trace(fmtCall)
		return uint32(C.gl12CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType Enum) uint32 {
		fmtCall := fmt.Sprintf("CreateShader(%v)", shaderType)
		defer glc.trace(fmtCall)
		return uint32(C.gl12CreateShader(glc.context, C.GLenum(shaderType)))
	}

	glc.DeleteProgram = func(program uint32) {
		fmtCall := fmt.Sprintf("DeleteProgram(%v)", program)
		defer glc.trace(fmtCall)
		C.gl12DeleteProgram(glc.context, C.GLuint(program))
	}

	glc.DeleteShader = func(shader uint32) {
		fmtCall := fmt.Sprintf("DeleteShader(%v)", shader)
		defer glc.trace(fmtCall)
		C.gl12DeleteShader(glc.context, C.GLuint(shader))
	}

	glc.DetachShader = func(program, shader uint32) {
		fmtCall := fmt.Sprintf("DetachShader(%v, %v)", program, shader)
		defer glc.trace(fmtCall)
		C.gl12DetachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.EnableVertexAttribArray = func(index uint32) {
		fmtCall := fmt.Sprintf("EnableVertexAttribArray(%v)", index)
		defer glc.trace(fmtCall)
		C.gl12EnableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DisableVertexAttribArray = func(index uint32) {
		fmtCall := fmt.Sprintf("DisableVertexAttribArray(%v)", index)
		defer glc.trace(fmtCall)
		C.gl12DisableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DrawBuffers = func(n int32, bufs *Enum) {
		fmtCall := fmt.Sprintf("DrawBuffers(%v, %v)", n, bufs)
		defer glc.trace(fmtCall)
		C.gl12DrawBuffers(glc.context, C.GLsizei(n), (*C.GLenum)(unsafe.Pointer(bufs)))
	}

	glc.GetActiveAttrib = func(program, index uint32, bufSize int32) (length int32, size int32, Type Enum, name string) {
		fmtCall := fmt.Sprintf("GetActiveAttrib(%v, %v, %v)", program, index, bufSize)
		defer glc.trace(fmtCall)
		var (
			cname C.GLchar
		)
		C.gl12GetActiveAttrib(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
		name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
		return
	}

	glc.GetActiveUniform = func(program, index uint32, bufSize int32, length *int32, size *int32, Type *Enum, name *byte) {
		fmtCall := fmt.Sprintf("GetActiveUniform(%v, %v, %v, %v, %v, %v, %v)", program, index, bufSize, length, size, Type, name)
		defer glc.trace(fmtCall)
		C.gl12GetActiveUniform(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(Type)), (*C.GLchar)(unsafe.Pointer(name)))
	}

	glc.GetAttachedShaders = func(program uint32, maxCount int32, count *int32, shaders *uint32) {
		fmtCall := fmt.Sprintf("GetAttachedShaders(%v, %v, %v, %v)", program, maxCount, count, shaders)
		defer glc.trace(fmtCall)
		C.gl12GetAttachedShaders(glc.context, C.GLuint(program), C.GLsizei(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
	}

	glc.GetAttribLocation = func(program uint32, name *byte) int32 {
		fmtCall := fmt.Sprintf("GetAttribLocation(%v, %v)", program, name)
		defer glc.trace(fmtCall)
		return int32(C.gl12GetAttribLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetProgramiv = func(program uint32, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetProgramiv(%v, %v, %v)", program, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetProgramiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetProgramInfoLog = func(program uint32, maxLength int32, length *int32, infoLog *byte) {
		fmtCall := fmt.Sprintf("GetProgramInfoLog(%v, %v, %v, %v)", program, maxLength, length, infoLog)
		defer glc.trace(fmtCall)
		C.gl12GetProgramInfoLog(glc.context, C.GLuint(program), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderiv = func(program uint32, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetShaderiv(%v, %v, %v)", program, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetShaderiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetShaderInfoLog = func(shader uint32, maxLength int32, length *int32, infoLog *byte) {
		fmtCall := fmt.Sprintf("GetShaderInfoLog(%v, %v, %v, %v)", shader, maxLength, length, infoLog)
		defer glc.trace(fmtCall)
		C.gl12GetShaderInfoLog(glc.context, C.GLuint(shader), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderSource = func(shader uint32, bufSize int32, length *int32, source *byte) {
		fmtCall := fmt.Sprintf("GetShaderSource(%v, %v, %v, %v)", shader, bufSize, length, source)
		defer glc.trace(fmtCall)
		C.gl12GetShaderSource(glc.context, C.GLuint(shader), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
	}

	glc.GetUniformfv = func(program uint32, location int32, params *float32) {
		fmtCall := fmt.Sprintf("GetUniformfv(%v, %v, %v)", program, location, params)
		defer glc.trace(fmtCall)
		C.gl12GetUniformfv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetUniformiv = func(program uint32, location int32, params *int32) {
		fmtCall := fmt.Sprintf("GetUniformiv(%v, %v, %v)", program, location, params)
		defer glc.trace(fmtCall)
		C.gl12GetUniformiv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetUniformLocation = func(program uint32, name *byte) int32 {
		fmtCall := fmt.Sprintf("GetUniformLocation(%v, %v)", program, name)
		defer glc.trace(fmtCall)
		return int32(C.gl12GetUniformLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetVertexAttribdv = func(index uint32, pname Enum, params *float64) {
		fmtCall := fmt.Sprintf("GetVertexAttribdv(%v, %v, %v)", index, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetVertexAttribdv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribfv = func(index uint32, pname Enum, params *float32) {
		fmtCall := fmt.Sprintf("GetVertexAttribfv(%v, %v, %v)", index, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetVertexAttribfv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribiv = func(index uint32, pname Enum, params *int32) {
		fmtCall := fmt.Sprintf("GetVertexAttribiv(%v, %v, %v)", index, pname, params)
		defer glc.trace(fmtCall)
		C.gl12GetVertexAttribiv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribPointerv = func(index uint32, pname Enum, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("GetVertexAttribPointerv(%v, %v, %v)", index, pname, pointer)
		defer glc.trace(fmtCall)
		C.gl12GetVertexAttribPointerv(glc.context, C.GLuint(index), C.GLenum(pname), pointer)
	}

	glc.IsProgram = func(program uint32) bool {
		fmtCall := fmt.Sprintf("IsProgram(%v)", program)
		defer glc.trace(fmtCall)
		return C.gl12IsProgram(glc.context, C.GLuint(program)) != 0
	}

	glc.IsShader = func(shader uint32) bool {
		fmtCall := fmt.Sprintf("IsShader(%v)", shader)
		defer glc.trace(fmtCall)
		return C.gl12IsShader(glc.context, C.GLuint(shader)) != 0
	}

	glc.LinkProgram = func(program uint32) {
		fmtCall := fmt.Sprintf("LinkProgram(%v)", program)
		defer glc.trace(fmtCall)
		C.gl12LinkProgram(glc.context, C.GLuint(program))
	}

	glc.ShaderSource = func(shader uint32, count int32, string **byte, length *int32) {
		fmtCall := fmt.Sprintf("ShaderSource(%v, %v, %v, %v)", shader, count, string, length)
		defer glc.trace(fmtCall)
		C.gl12ShaderSource(glc.context, C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(string)), (*C.GLint)(unsafe.Pointer(length)))
	}

	glc.StencilFuncSeparate = func(face, Func Enum, ref int32, mask uint32) {
		fmtCall := fmt.Sprintf("StencilFuncSeparate(%v, %v, %v, %v)", face, Func, ref, mask)
		defer glc.trace(fmtCall)
		C.gl12StencilFuncSeparate(glc.context, C.GLenum(face), C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMaskSeparate = func(face Enum, mask uint32) {
		fmtCall := fmt.Sprintf("StencilMaskSeparate(%v, %v)", face, mask)
		defer glc.trace(fmtCall)
		C.gl12StencilMaskSeparate(glc.context, C.GLenum(face), C.GLuint(mask))
	}

	glc.StencilOpSeparate = func(face, sfail, dpfail, dppass Enum) {
		fmtCall := fmt.Sprintf("StencilOpSeparate(%v, %v, %v, %v)", face, sfail, dpfail, dppass)
		defer glc.trace(fmtCall)
		C.gl12StencilOpSeparate(glc.context, C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
	}

	glc.Uniform1f = func(location int32, v0 float32) {
		fmtCall := fmt.Sprintf("Uniform1f(%v, %v)", location, v0)
		defer glc.trace(fmtCall)
		C.gl12Uniform1f(glc.context, C.GLint(location), C.GLfloat(v0))
	}

	glc.Uniform2f = func(location int32, v0, v1 float32) {
		fmtCall := fmt.Sprintf("Uniform2f(%v, %v, %v)", location, v0, v1)
		defer glc.trace(fmtCall)
		C.gl12Uniform2f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.Uniform3f = func(location int32, v0, v1, v2 float32) {
		fmtCall := fmt.Sprintf("Uniform3f(%v, %v, %v, %v)", location, v0, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Uniform3f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Uniform4f = func(location int32, v0, v1, v2, v3 float32) {
		fmtCall := fmt.Sprintf("Uniform4f(%v, %v, %v, %v, %v)", location, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12Uniform4f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.Uniform1i = func(location, v0 int32) {
		fmtCall := fmt.Sprintf("Uniform1i(%v, %v)", location, v0)
		defer glc.trace(fmtCall)
		C.gl12Uniform1i(glc.context, C.GLint(location), C.GLint(v0))
	}

	glc.Uniform2i = func(location, v0, v1 int32) {
		fmtCall := fmt.Sprintf("Uniform2i(%v, %v, %v)", location, v0, v1)
		defer glc.trace(fmtCall)
		C.gl12Uniform2i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1))
	}

	glc.Uniform3i = func(location, v0, v1, v2 int32) {
		fmtCall := fmt.Sprintf("Uniform3i(%v, %v, %v, %v)", location, v0, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Uniform3i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
	}

	glc.Uniform4i = func(location, v0, v1, v2, v3 int32) {
		fmtCall := fmt.Sprintf("Uniform4i(%v, %v, %v, %v, %v)", location, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12Uniform4i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
	}

	glc.Uniform1ui = func(location int32, v0 uint32) {
		fmtCall := fmt.Sprintf("Uniform1ui(%v, %v)", location, v0)
		defer glc.trace(fmtCall)
		C.gl12Uniform1ui(glc.context, C.GLint(location), C.GLuint(v0))
	}

	glc.Uniform2ui = func(location int32, v0, v1 uint32) {
		fmtCall := fmt.Sprintf("Uniform2ui(%v, %v, %v)", location, v0, v1)
		defer glc.trace(fmtCall)
		C.gl12Uniform2ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1))
	}

	glc.Uniform3ui = func(location int32, v0, v1, v2 uint32) {
		fmtCall := fmt.Sprintf("Uniform3ui(%v, %v, %v, %v)", location, v0, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12Uniform3ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
	}

	glc.Uniform4ui = func(location int32, v0, v1, v2, v3 uint32) {
		fmtCall := fmt.Sprintf("Uniform4ui(%v, %v, %v, %v, %v)", location, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12Uniform4ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
	}

	glc.Uniform1fv = func(location int32, count int32, value *float32) {
		fmtCall := fmt.Sprintf("Uniform1fv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform1fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform2fv = func(location int32, count int32, value *float32) {
		fmtCall := fmt.Sprintf("Uniform2fv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform2fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform3fv = func(location int32, count int32, value *float32) {
		fmtCall := fmt.Sprintf("Uniform3fv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform3fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform4fv = func(location int32, count int32, value *float32) {
		fmtCall := fmt.Sprintf("Uniform4fv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform4fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform1iv = func(location int32, count int32, value *int32) {
		fmtCall := fmt.Sprintf("Uniform1iv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform1iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform2iv = func(location int32, count int32, value *int32) {
		fmtCall := fmt.Sprintf("Uniform2iv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform2iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform3iv = func(location int32, count int32, value *int32) {
		fmtCall := fmt.Sprintf("Uniform3iv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform3iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform4iv = func(location int32, count int32, value *int32) {
		fmtCall := fmt.Sprintf("Uniform4iv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform4iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform1uiv = func(location int32, count int32, value *uint32) {
		fmtCall := fmt.Sprintf("Uniform1uiv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform1uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform2uiv = func(location int32, count int32, value *uint32) {
		fmtCall := fmt.Sprintf("Uniform2uiv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform2uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform3uiv = func(location int32, count int32, value *uint32) {
		fmtCall := fmt.Sprintf("Uniform3uiv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform3uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform4uiv = func(location int32, count int32, value *uint32) {
		fmtCall := fmt.Sprintf("Uniform4uiv(%v, %v, %v)", location, count, value)
		defer glc.trace(fmtCall)
		C.gl12Uniform4uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.UseProgram = func(program uint32) {
		fmtCall := fmt.Sprintf("UseProgram(%v)", program)
		defer glc.trace(fmtCall)
		C.gl12UseProgram(glc.context, C.GLuint(program))
	}

	glc.ValidateProgram = func(program uint32) {
		fmtCall := fmt.Sprintf("ValidateProgram(%v)", program)
		defer glc.trace(fmtCall)
		C.gl12ValidateProgram(glc.context, C.GLuint(program))
	}

	glc.VertexAttribPointer = func(index uint32, size int32, Type Enum, normalized bool, stride int32, pointer unsafe.Pointer) {
		fmtCall := fmt.Sprintf("VertexAttribPointer(%v, %v, %v, %v, %v, %v)", index, size, Type, normalized, stride, pointer)
		defer glc.trace(fmtCall)
		C.gl12VertexAttribPointer(glc.context, C.GLuint(index), C.GLint(size), C.GLenum(Type), boolToGL(normalized), C.GLsizei(stride), pointer)
	}

	glc.VertexAttrib1f = func(index uint32, v0 float32) {
		fmtCall := fmt.Sprintf("VertexAttrib1f(%v, %v)", index, v0)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib1f(glc.context, C.GLuint(index), C.GLfloat(v0))
	}

	glc.VertexAttrib1s = func(index uint32, v0 int16) {
		fmtCall := fmt.Sprintf("VertexAttrib1s(%v, %v)", index, v0)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib1s(glc.context, C.GLuint(index), C.GLshort(v0))
	}

	glc.VertexAttrib1d = func(index uint32, v0 float64) {
		fmtCall := fmt.Sprintf("VertexAttrib1d(%v, %v)", index, v0)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib1d(glc.context, C.GLuint(index), C.GLdouble(v0))
	}

	glc.VertexAttrib2f = func(index uint32, v0, v1 float32) {
		fmtCall := fmt.Sprintf("VertexAttrib2f(%v, %v, %v)", index, v0, v1)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib2f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.VertexAttrib2s = func(index uint32, v0, v1 int16) {
		fmtCall := fmt.Sprintf("VertexAttrib2s(%v, %v, %v)", index, v0, v1)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib2s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
	}

	glc.VertexAttrib2d = func(index uint32, v0, v1 float64) {
		fmtCall := fmt.Sprintf("VertexAttrib2d(%v, %v, %v)", index, v0, v1)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib2d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
	}

	glc.VertexAttrib3f = func(index uint32, v0, v1, v2 float32) {
		fmtCall := fmt.Sprintf("VertexAttrib3f(%v, %v, %v, %v)", index, v0, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib3f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.VertexAttrib3s = func(index uint32, v0, v1, v2 int16) {
		fmtCall := fmt.Sprintf("VertexAttrib3s(%v, %v, %v, %v)", index, v0, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib3s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
	}

	glc.VertexAttrib3d = func(index uint32, v0, v1, v2 float64) {
		fmtCall := fmt.Sprintf("VertexAttrib3d(%v, %v, %v, %v)", index, v0, v1, v2)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib3d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.VertexAttrib4f = func(index uint32, v0, v1, v2, v3 float32) {
		fmtCall := fmt.Sprintf("VertexAttrib4f(%v, %v, %v, %v, %v)", index, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.VertexAttrib4s = func(index uint32, v0, v1, v2, v3 int16) {
		fmtCall := fmt.Sprintf("VertexAttrib4s(%v, %v, %v, %v, %v)", index, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
	}

	glc.VertexAttrib4d = func(index uint32, v0, v1, v2, v3 float64) {
		fmtCall := fmt.Sprintf("VertexAttrib4d(%v, %v, %v, %v, %v)", index, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
	}

	glc.VertexAttrib4Nuv = func(index uint32, v0, v1, v2, v3 uint8) {
		fmtCall := fmt.Sprintf("VertexAttrib4Nuv(%v, %v, %v, %v, %v)", index, v0, v1, v2, v3)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Nuv(glc.context, C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
	}

	glc.VertexAttrib1fv = func(index uint32, v *float32) {
		fmtCall := fmt.Sprintf("VertexAttrib1fv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib1fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1sv = func(index uint32, v *int16) {
		fmtCall := fmt.Sprintf("VertexAttrib1sv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib1sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1dv = func(index uint32, v *float64) {
		fmtCall := fmt.Sprintf("VertexAttrib1dv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib1dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2fv = func(index uint32, v *float32) {
		fmtCall := fmt.Sprintf("VertexAttrib2fv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib2fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2sv = func(index uint32, v *int16) {
		fmtCall := fmt.Sprintf("VertexAttrib2sv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib2sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2dv = func(index uint32, v *float64) {
		fmtCall := fmt.Sprintf("VertexAttrib2dv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib2dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3fv = func(index uint32, v *float32) {
		fmtCall := fmt.Sprintf("VertexAttrib3fv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib3fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3sv = func(index uint32, v *int16) {
		fmtCall := fmt.Sprintf("VertexAttrib3sv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib3sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3dv = func(index uint32, v *float64) {
		fmtCall := fmt.Sprintf("VertexAttrib3dv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib3dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4fv = func(index uint32, v *float32) {
		fmtCall := fmt.Sprintf("VertexAttrib4fv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4sv = func(index uint32, v *int16) {
		fmtCall := fmt.Sprintf("VertexAttrib4sv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4dv = func(index uint32, v *float64) {
		fmtCall := fmt.Sprintf("VertexAttrib4dv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4iv = func(index uint32, v *int32) {
		fmtCall := fmt.Sprintf("VertexAttrib4iv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4iv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4bv = func(index uint32, v *int8) {
		fmtCall := fmt.Sprintf("VertexAttrib4bv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4bv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4ubv = func(index uint32, v *uint8) {
		fmtCall := fmt.Sprintf("VertexAttrib4ubv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4ubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4usv = func(index uint32, v *uint16) {
		fmtCall := fmt.Sprintf("VertexAttrib4usv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4usv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4uiv = func(index uint32, v *uint32) {
		fmtCall := fmt.Sprintf("VertexAttrib4uiv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4uiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nbv = func(index uint32, v *int8) {
		fmtCall := fmt.Sprintf("VertexAttrib4Nbv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Nbv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nsv = func(index uint32, v *int16) {
		fmtCall := fmt.Sprintf("VertexAttrib4Nsv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Nsv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Niv = func(index uint32, v *int32) {
		fmtCall := fmt.Sprintf("VertexAttrib4Niv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Niv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nubv = func(index uint32, v *uint8) {
		fmtCall := fmt.Sprintf("VertexAttrib4Nubv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Nubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nusv = func(index uint32, v *uint16) {
		fmtCall := fmt.Sprintf("VertexAttrib4Nusv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Nusv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nuiv = func(index uint32, v *uint32) {
		fmtCall := fmt.Sprintf("VertexAttrib4Nuiv(%v, %v)", index, v)
		defer glc.trace(fmtCall)
		C.gl12VertexAttrib4Nuiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.UniformMatrix2fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix2fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix3fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix4fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x3fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix2x3fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix2x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x2fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix3x2fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix3x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x4fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix2x4fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix2x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x2fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix4x2fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix4x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x4fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix3x4fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix3x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x3fv = func(location int32, count int32, transpose bool, value *float32) {
		fmtCall := fmt.Sprintf("UniformMatrix4x3fv(%v, %v, %v, %v)", location, count, transpose, value)
		defer glc.trace(fmtCall)
		C.gl12UniformMatrix4x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	if !versionSupported(glc) {
		return nil
	}
	glc.queryExtensions()
	return glc
}
