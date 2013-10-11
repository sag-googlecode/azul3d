// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// +build !opengl_debug

// Package 'opengl' implements OpenGL version 1.4
package opengl

// #cgo windows LDFLAGS: -lopengl32
// #cgo linux LDFLAGS: -lGL -ldl
// #include "gl14.h"
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
		n, wantedMajor, wantedMinor, wantedRev := parseVersions("1.4")
		if n < 2 {
			fmt.Printf("OpenGL: *** JSON version parsing failed for %q ***\n", "1.4")
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

const (
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	COMBINE_RGB_ARB                                            = 0x8571
	DRAW_BUFFER1                                               = 0x8826
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	INTENSITY16UI_EXT                                          = 0x8D79
	LOGIC_OP                                                   = 0x0BF1
	VERTEX_ARRAY_EXT                                           = 0x8074
	WRAP_BORDER_SUN                                            = 0x81D4
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	LINE_STRIP                                                 = 0x0003
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	SWIZZLE_STQ_ATI                                            = 0x8977
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	COMPRESSED_R11_EAC                                         = 0x9270
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	SRGB_EXT                                                   = 0x8C40
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	TEXTURE_3D_OES                                             = 0x806F
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	MAP2_BINORMAL_EXT                                          = 0x8447
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	CMYKA_EXT                                                  = 0x800D
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	FLOAT_RG_NV                                                = 0x8881
	REG_19_ATI                                                 = 0x8934
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	EXT_histogram                                              = 1
	HINT_BIT                                                   = 0x00008000
	OPERAND1_ALPHA_ARB                                         = 0x8599
	MATRIX0_NV                                                 = 0x8630
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	RG8_EXT                                                    = 0x822B
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	SOURCE2_RGB                                                = 0x8582
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	RGBA_SNORM                                                 = 0x8F93
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	INDEX_ARRAY_LIST_IBM                                       = 103073
	PROJECTION                                                 = 0x1701
	CONSTANT_ARB                                               = 0x8576
	RGB32F_ARB                                                 = 0x8815
	DOUBLE_MAT2_EXT                                            = 0x8F46
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	FOG_INDEX                                                  = 0x0B61
	RGBA8_EXT                                                  = 0x8058
	DSDT_MAG_VIB_NV                                            = 0x86F7
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	FLOAT_MAT4                                                 = 0x8B5C
	STENCIL_ATTACHMENT                                         = 0x8D20
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	LUMINANCE8_EXT                                             = 0x8040
	SRC1_ALPHA                                                 = 0x8589
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	QUARTER_BIT_ATI                                            = 0x00000010
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	TEXTURE8                                                   = 0x84C8
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	IMAGE_2D_EXT                                               = 0x904D
	QUERY_OBJECT_EXT                                           = 0x9153
	CON_25_ATI                                                 = 0x895A
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	QUADRATIC_ATTENUATION                                      = 0x1209
	T                                                          = 0x2001
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	NEGATIVE_W_EXT                                             = 0x87DC
	VARIANT_VALUE_EXT                                          = 0x87E4
	CON_6_ATI                                                  = 0x8947
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	DRAW_BUFFER12                                              = 0x8831
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	GREATER                                                    = 0x0204
	TEXTURE_COORD_NV                                           = 0x8C79
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	CONTEXT_PROFILE_MASK                                       = 0x9126
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	COLOR_ATTACHMENT10                                         = 0x8CEA
	STANDARD_FONT_NAME_NV                                      = 0x9072
	ZERO                                                       = 0
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	TRIANGLE_MESH_SUN                                          = 0x8615
	DRAW_BUFFER10_ATI                                          = 0x882F
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	REFLECTION_MAP_EXT                                         = 0x8512
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	STREAM_COPY_ARB                                            = 0x88E2
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	SIGNALED_APPLE                                             = 0x9119
	RGB5                                                       = 0x8050
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	SAMPLES_EXT                                                = 0x80A9
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	SAMPLER_3D_OES                                             = 0x8B5F
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	NICEST                                                     = 0x1102
	LIGHT1                                                     = 0x4001
	SIGNED_ALPHA8_NV                                           = 0x8706
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	INDEX_ARRAY_TYPE                                           = 0x8085
	FILL                                                       = 0x1B02
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	SAMPLE_MASK_EXT                                            = 0x80A0
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	DATA_BUFFER_AMD                                            = 0x9151
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	SAMPLES_PASSED_ARB                                         = 0x8914
	UNIFORM_BUFFER                                             = 0x8A11
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	EXT_point_parameters                                       = 1
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	PROXY_HISTOGRAM                                            = 0x8025
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	TEXTURE25                                                  = 0x84D9
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	VERSION_3_2                                                = 1
	COMBINER_SCALE_NV                                          = 0x8548
	LUMINANCE8_ALPHA8                                          = 0x8045
	DRAW_BUFFER2_NV                                            = 0x8827
	DOUBLE_MAT4                                                = 0x8F48
	MAP_WRITE_BIT_EXT                                          = 0x0002
	QUERY                                                      = 0x82E3
	SOURCE3_ALPHA_NV                                           = 0x858B
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	CONSTANT_BORDER                                            = 0x8151
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	STATE_RESTORE                                              = 0x8BDC
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	SGIX_tag_sample_buffer                                     = 1
	FRONT                                                      = 0x0404
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	PERCENTAGE_AMD                                             = 0x8BC3
	SGIX_impact_pixel_texture                                  = 1
	SCALAR_EXT                                                 = 0x87BE
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	SLUMINANCE8_NV                                             = 0x8C47
	COLOR_ATTACHMENT1                                          = 0x8CE1
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	INT_IMAGE_3D                                               = 0x9059
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	DOUBLE_MAT2                                                = 0x8F46
	POLYGON_STIPPLE                                            = 0x0B42
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	POLYGON_OFFSET_LINE                                        = 0x2A02
	TEXTURE_BLUE_SIZE                                          = 0x805E
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	INTENSITY8_EXT                                             = 0x804B
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	CURRENT_INDEX                                              = 0x0B01
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	ITALIC_BIT_NV                                              = 0x02
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	CON_2_ATI                                                  = 0x8943
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	RGBA8I_EXT                                                 = 0x8D8E
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	INCR_WRAP_OES                                              = 0x8507
	MODELVIEW13_ARB                                            = 0x872D
	INT8_VEC3_NV                                               = 0x8FE2
	INT_IMAGE_1D                                               = 0x9057
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	COLOR_ATTACHMENT3                                          = 0x8CE3
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	BOOL_VEC3                                                  = 0x8B58
	REPLICATE_BORDER                                           = 0x8153
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	FENCE_APPLE                                                = 0x8A0B
	FULL_SUPPORT                                               = 0x82B7
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	FRAGMENT_TEXTURE                                           = 0x829F
	COMBINE_RGB                                                = 0x8571
	LIGHTING_BIT                                               = 0x00000040
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	ONE_EXT                                                    = 0x87DE
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	BUFFER_MAP_OFFSET                                          = 0x9121
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	MAP2_TANGENT_EXT                                           = 0x8445
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	BIAS_BIT_ATI                                               = 0x00000008
	BGR_INTEGER_EXT                                            = 0x8D9A
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	MIRRORED_REPEAT_OES                                        = 0x8370
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	GL_4PASS_2_EXT                                             = 0x80A6
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	EXT_blend_logic_op                                         = 1
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	GL_1PASS_SGIS                                              = 0x80A1
	RGB_FLOAT16_ATI                                            = 0x881B
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	STENCIL_INDEX1_EXT                                         = 0x8D46
	LUMINANCE32I_EXT                                           = 0x8D86
	LIGHT2                                                     = 0x4002
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	R11F_G11F_B10F                                             = 0x8C3A
	INTENSITY16I_EXT                                           = 0x8D8B
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	LIGHT7                                                     = 0x4007
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	TEXTURE26                                                  = 0x84DA
	TEXTURE19_ARB                                              = 0x84D3
	GL_3D_COLOR                                                = 0x0602
	NAND                                                       = 0x150E
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	POINT_SIZE_MIN                                             = 0x8126
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	REG_7_ATI                                                  = 0x8928
	MAX_SAMPLES_EXT                                            = 0x8D57
	SAMPLE_BUFFERS                                             = 0x80A8
	SAMPLER_1D_SHADOW                                          = 0x8B61
	TIMESTAMP                                                  = 0x8E28
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	VERSION_1_4                                                = 1
	COLOR_INDEX16_EXT                                          = 0x80E7
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	COLOR_ARRAY                                                = 0x8076
	READ_BUFFER_EXT                                            = 0x0C02
	CONSTANT                                                   = 0x8576
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	TRIANGLE_FAN                                               = 0x0006
	STACK_OVERFLOW                                             = 0x0503
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	MAX_VARYING_VECTORS                                        = 0x8DFC
	LUMINANCE12_ALPHA4                                         = 0x8046
	CMYK_EXT                                                   = 0x800C
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	FOG_OFFSET_SGIX                                            = 0x8198
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	CON_11_ATI                                                 = 0x894C
	FLOAT_VEC2_ARB                                             = 0x8B50
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	PACK_SKIP_PIXELS                                           = 0x0D04
	SMOOTH                                                     = 0x1D01
	R                                                          = 0x2002
	FUNC_ADD_OES                                               = 0x8006
	RGB9_E5_EXT                                                = 0x8C3D
	QUERY_NO_WAIT                                              = 0x8E14
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	T2F_IUI_V3F_EXT                                            = 0x81B2
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	SGIX_shadow_ambient                                        = 1
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	INT16_NV                                                   = 0x8FE4
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	SGX_BINARY_IMG                                             = 0x8C0A
	MULTIVIEW_EXT                                              = 0x90F1
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	CURRENT_QUERY_EXT                                          = 0x8865
	BOOL                                                       = 0x8B56
	DEPTH_RANGE                                                = 0x0B70
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	INT_IMAGE_CUBE                                             = 0x905B
	AFFINE_2D_NV                                               = 0x9092
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	FIXED_ONLY                                                 = 0x891D
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	UNPACK_LSB_FIRST                                           = 0x0CF1
	MAP1_VERTEX_3                                              = 0x0D97
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	TEXTURE13_ARB                                              = 0x84CD
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	POLYGON_TOKEN                                              = 0x0703
	OPERAND0_RGB                                               = 0x8590
	RGB32I                                                     = 0x8D83
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	MAX_TEXTURE_UNITS                                          = 0x84E2
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	MAX_PATCH_VERTICES                                         = 0x8E7D
	EXT_blend_minmax                                           = 1
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	OPERAND0_ALPHA_ARB                                         = 0x8598
	MATRIX19_ARB                                               = 0x88D3
	PATH_FORMAT_PS_NV                                          = 0x9071
	UNPACK_ALIGNMENT                                           = 0x0CF5
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	FENCE_STATUS_NV                                            = 0x84F3
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	HIGH_FLOAT                                                 = 0x8DF2
	MULTISAMPLE_ARB                                            = 0x809D
	SRGB_ALPHA_EXT                                             = 0x8C42
	DEPTH_ATTACHMENT                                           = 0x8D00
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	BLUE_BITS                                                  = 0x0D54
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	NOTEQUAL                                                   = 0x0205
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	EYE_LINE_SGIS                                              = 0x81F6
	SYNC_OBJECT_APPLE                                          = 0x8A53
	MATRIX_STRIDE                                              = 0x92FF
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	TEXTURE23_ARB                                              = 0x84D7
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	TEXTURE_LOD_BIAS                                           = 0x8501
	TIME_ELAPSED_EXT                                           = 0x88BF
	LERP_ATI                                                   = 0x8969
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	VERTEX_STREAM6_ATI                                         = 0x8772
	DRAW_BUFFER0                                               = 0x8825
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	MAX_SAMPLES_IMG                                            = 0x9135
	FLOAT                                                      = 0x1406
	LUMINANCE6_ALPHA2                                          = 0x8044
	REPLACE_MIDDLE_SUN                                         = 0x0002
	ALPHA32I_EXT                                               = 0x8D84
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	SGIX_interlace                                             = 1
	SLIM12S_SGIX                                               = 0x831F
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	SHADER_OPERATION_NV                                        = 0x86DF
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	COMBINER3_NV                                               = 0x8553
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	INTENSITY_EXT                                              = 0x8049
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	COORD_REPLACE                                              = 0x8862
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	DOUBLE_VEC3                                                = 0x8FFD
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	RESCALE_NORMAL                                             = 0x803A
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	BGRA                                                       = 0x80E1
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	CON_7_ATI                                                  = 0x8948
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	COMPRESSED_SRGB                                            = 0x8C48
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	EYE_RADIAL_NV                                              = 0x855B
	MATRIX13_ARB                                               = 0x88CD
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	IMAGE_2D_RECT_EXT                                          = 0x904F
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	MODELVIEW6_ARB                                             = 0x8726
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	RGBA_MODE                                                  = 0x0C31
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	REPLACE_OLDEST_SUN                                         = 0x0003
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	VERSION_3_0                                                = 1
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	WRITE_ONLY_OES                                             = 0x88B9
	CLAMP_VERTEX_COLOR                                         = 0x891A
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	ARB_imaging                                                = 1
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	LAYOUT_DEFAULT_INTEL                                       = 0
	FOG_DENSITY                                                = 0x0B62
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	LINES_ADJACENCY                                            = 0x000A
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	MATRIX1_NV                                                 = 0x8631
	DRAW_BUFFER13_ATI                                          = 0x8832
	REG_15_ATI                                                 = 0x8930
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	C3F_V3F                                                    = 0x2A24
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	ATTENUATION_EXT                                            = 0x834D
	OP_SET_GE_EXT                                              = 0x878C
	OP_POWER_EXT                                               = 0x8793
	PROVOKING_VERTEX                                           = 0x8E4F
	COMPRESSED_INTENSITY                                       = 0x84EC
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	OCCLUSION_TEST_HP                                          = 0x8165
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	VARIABLE_D_NV                                              = 0x8526
	DRAW_BUFFER7_NV                                            = 0x882C
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	CON_9_ATI                                                  = 0x894A
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	BUFFER_MAP_POINTER                                         = 0x88BD
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	CONTINUOUS_AMD                                             = 0x9007
	TEXTURE_GREEN_SIZE                                         = 0x805D
	CLIP_DISTANCE6                                             = 0x3006
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	GREEN_INTEGER                                              = 0x8D95
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	GL_4_BYTES                                                 = 0x1409
	EYE_PLANE                                                  = 0x2502
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	CON_24_ATI                                                 = 0x8959
	SHADER_BINARY_VIV                                          = 0x8FC4
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	GL_422_AVERAGE_EXT                                         = 0x80CE
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	CLAMP_READ_COLOR                                           = 0x891C
	MAD_ATI                                                    = 0x8968
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	BLEND_DST_RGB_OES                                          = 0x80C8
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	TEXTURE3_ARB                                               = 0x84C3
	TEXTURE31_ARB                                              = 0x84DF
	COLOR_MATERIAL_FACE                                        = 0x0B55
	RESCALE_NORMAL_EXT                                         = 0x803A
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	LUMINANCE4_EXT                                             = 0x803F
	TEXTURE29                                                  = 0x84DD
	DRAW_BUFFER3                                               = 0x8828
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	MAX_SAMPLES                                                = 0x8D57
	PHONG_HINT_WIN                                             = 0x80EB
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	PRIMITIVE_RESTART_NV                                       = 0x8558
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	RGB16_SNORM                                                = 0x8F9A
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	SET                                                        = 0x150F
	FILTER                                                     = 0x829A
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	FLOAT16_NV                                                 = 0x8FF8
	UNSIGNALED_APPLE                                           = 0x9118
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	FRAMEBUFFER_OES                                            = 0x8D40
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	LINE_SMOOTH                                                = 0x0B20
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	DRAW_BUFFER8_ARB                                           = 0x882D
	MODELVIEW11_ARB                                            = 0x872B
	TRANSLATE_Y_NV                                             = 0x908F
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	W_EXT                                                      = 0x87D8
	ALPHA32F_ARB                                               = 0x8816
	RENDERBUFFER_BINDING                                       = 0x8CA7
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	CONSTANT_COLOR                                             = 0x8001
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	SOURCE2_RGB_EXT                                            = 0x8582
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	INT_VEC2_ARB                                               = 0x8B53
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	DUAL_ALPHA12_SGIS                                          = 0x8112
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	MAX_VERTEX_STREAMS                                         = 0x8E71
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	LINE_BIT                                                   = 0x00000004
	GL_3D                                                      = 0x0601
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	RGBA_S3TC                                                  = 0x83A2
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	PROGRAM_INPUT                                              = 0x92E3
	EXT_packed_pixels                                          = 1
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	ALPHA16_EXT                                                = 0x803E
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	DRAW_BUFFER0_ARB                                           = 0x8825
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	MODELVIEW27_ARB                                            = 0x873B
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	READ_BUFFER                                                = 0x0C02
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	REG_5_ATI                                                  = 0x8926
	PACK_SWAP_BYTES                                            = 0x0D00
	INT_IMAGE_2D                                               = 0x9058
	TEXTURE_BINDING_1D                                         = 0x8068
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	DRAW_BUFFER2_ATI                                           = 0x8827
	REG_14_ATI                                                 = 0x892F
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	REG_31_ATI                                                 = 0x8940
	BOOL_VEC2_ARB                                              = 0x8B57
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	IDENTITY_NV                                                = 0x862A
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	RGB12                                                      = 0x8053
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	CON_26_ATI                                                 = 0x895B
	SGIX_async                                                 = 1
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	COPY_INVERTED                                              = 0x150C
	REFLECTION_MAP_ARB                                         = 0x8512
	TRACK_MATRIX_NV                                            = 0x8648
	VERTEX_SOURCE_ATI                                          = 0x8774
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	CAVEAT_SUPPORT                                             = 0x82B8
	SLIM8U_SGIX                                                = 0x831D
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	REG_26_ATI                                                 = 0x893B
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	ALL_STATIC_DATA_IBM                                        = 103060
	EXT_convolution                                            = 1
	ALPHA4                                                     = 0x803B
	MAX_VIEWPORTS                                              = 0x825B
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	COLOR_MATRIX_SGI                                           = 0x80B1
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	ONE                                                        = 1
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	ALPHA8                                                     = 0x803C
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	RGBA_INTEGER_EXT                                           = 0x8D99
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	INT64_VEC3_NV                                              = 0x8FEA
	REFLECTION_MAP                                             = 0x8512
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	PRIMITIVES_GENERATED                                       = 0x8C87
	OFFSET                                                     = 0x92FC
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	DOUBLE_MAT2x4                                              = 0x8F4A
	EXT_shared_texture_palette                                 = 1
	INDEX_SHIFT                                                = 0x0D12
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	MODELVIEW14_ARB                                            = 0x872E
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	VARIABLE_A_NV                                              = 0x8523
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	COLOR_ATTACHMENT8                                          = 0x8CE8
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	FOG_COLOR                                                  = 0x0B66
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	SIGNED_NEGATE_NV                                           = 0x853D
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	INFO_LOG_LENGTH                                            = 0x8B84
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	SURFACE_MAPPED_NV                                          = 0x8700
	INT_VEC3                                                   = 0x8B54
	INTENSITY32UI_EXT                                          = 0x8D73
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	INT_IMAGE_2D_EXT                                           = 0x9058
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	OPERAND2_RGB_EXT                                           = 0x8592
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	FLOAT_RGBA_NV                                              = 0x8883
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	SYNC_CL_EVENT_ARB                                          = 0x8240
	MATRIX10_ARB                                               = 0x88CA
	MATRIX26_ARB                                               = 0x88DA
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	SYNC_CONDITION                                             = 0x9113
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	DEPTH_TEXTURE_MODE                                         = 0x884B
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	MIN_LOD_WARNING_AMD                                        = 0x919C
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	TEXTURE22                                                  = 0x84D6
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	RGBA32UI                                                   = 0x8D70
	RED_INTEGER                                                = 0x8D94
	ALPHA8_EXT                                                 = 0x803C
	POINT_SIZE_MAX_ARB                                         = 0x8127
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	RGBA16I                                                    = 0x8D88
	SGIX_pixel_tiles                                           = 1
	EXPAND_NEGATE_NV                                           = 0x8539
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	SLIM10U_SGIX                                               = 0x831E
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	CON_1_ATI                                                  = 0x8942
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	LINE_RESET_TOKEN                                           = 0x0707
	DRAW_BUFFER9_ATI                                           = 0x882E
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	PRESENT_TIME_NV                                            = 0x8E2A
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	MAX_CLIP_PLANES                                            = 0x0D32
	COMPRESSED_RGB                                             = 0x84ED
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	INDEX_OFFSET                                               = 0x0D13
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	SIGNED_IDENTITY_NV                                         = 0x853C
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	QUERY_WAIT                                                 = 0x8E13
	TRIANGLE_LIST_SUN                                          = 0x81D7
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	FRAMEBUFFER_BLEND                                          = 0x828B
	MATRIX2_NV                                                 = 0x8632
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	SOURCE0_ALPHA_ARB                                          = 0x8588
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	RIGHT                                                      = 0x0407
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	INDEX_WRITEMASK                                            = 0x0C21
	PACK_IMAGE_HEIGHT                                          = 0x806C
	VERTEX_BLEND_ARB                                           = 0x86A7
	GL_3_BYTES                                                 = 0x1408
	OPERAND0_RGB_ARB                                           = 0x8590
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	COLOR_ATTACHMENT0                                          = 0x8CE0
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	NOR                                                        = 0x1508
	TEXTURE18                                                  = 0x84D2
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	DRAW_BUFFER4_ARB                                           = 0x8829
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	TEXTURE2                                                   = 0x84C2
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	MAP2_NORMAL                                                = 0x0DB2
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	DEPTH_STENCIL_EXT                                          = 0x84F9
	SRC2_RGB                                                   = 0x8582
	PREVIOUS                                                   = 0x8578
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	R8_SNORM                                                   = 0x8F94
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	INT_IMAGE_3D_EXT                                           = 0x9059
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	GPU_ADDRESS_NV                                             = 0x8F34
	UNSIGNED_INT8_NV                                           = 0x8FEC
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	ZERO_EXT                                                   = 0x87DD
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	SGIX_depth_texture                                         = 1
	VERTEX_ARRAY_POINTER                                       = 0x808E
	SHADER_IMAGE_STORE                                         = 0x82A5
	YCRCB_SGIX                                                 = 0x8318
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	COMPUTE_SUBROUTINE                                         = 0x92ED
	STENCIL_BUFFER_BIT                                         = 0x00000400
	FUNC_SUBTRACT_EXT                                          = 0x800A
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	COLOR_INDEX8_EXT                                           = 0x80E5
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	RG_SNORM                                                   = 0x8F91
	VERSION_2_1                                                = 1
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	VERTEX_STREAM4_ATI                                         = 0x8770
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	STENCIL_INDEX                                              = 0x1901
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	BUFFER_SIZE_ARB                                            = 0x8764
	INT8_VEC2_NV                                               = 0x8FE1
	BLEND_COLOR                                                = 0x8005
	HILO16_NV                                                  = 0x86F8
	DRAW_BUFFER8_ATI                                           = 0x882D
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	FIXED_ONLY_ARB                                             = 0x891D
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	FAILURE_NV                                                 = 0x9030
	UNSIGNED_INVERT_NV                                         = 0x8537
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	RGB10                                                      = 0x8052
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	ALPHA_INTEGER                                              = 0x8D97
	COLOR_SAMPLES_NV                                           = 0x8E20
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	R32UI                                                      = 0x8236
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	VARIABLE_C_NV                                              = 0x8525
	HIGH_INT                                                   = 0x8DF5
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	DOUBLE_MAT4x3                                              = 0x8F4E
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	NEGATIVE_Z_EXT                                             = 0x87DB
	LINK_STATUS                                                = 0x8B82
	SGIX_instruments                                           = 1
	RGB16_EXT                                                  = 0x8054
	TRANSPOSE_NV                                               = 0x862C
	LINES                                                      = 0x0001
	R16F                                                       = 0x822D
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	BLUE_BIT_ATI                                               = 0x00000004
	EXT_blend_color                                            = 1
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	MATRIX5_NV                                                 = 0x8635
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	DRAW_BUFFER5_ARB                                           = 0x882A
	SAMPLE_POSITION_NV                                         = 0x8E50
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	RG32F                                                      = 0x8230
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	SYNC_STATUS_APPLE                                          = 0x9114
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	NORMALIZE                                                  = 0x0BA1
	CONVOLUTION_1D                                             = 0x8010
	CUBIC_HP                                                   = 0x815F
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	TRACE_MASK_MESA                                            = 0x8755
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	STEREO                                                     = 0x0C33
	BUFFER_MAPPED_OES                                          = 0x88BC
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	INT_2_10_10_10_REV                                         = 0x8D9F
	RGB32UI                                                    = 0x8D71
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	RGB16F                                                     = 0x881B
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	PACK_SKIP_IMAGES                                           = 0x806B
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	DEPTH_STENCIL_MESA                                         = 0x8750
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	TEXTURE10                                                  = 0x84CA
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	HI_BIAS_NV                                                 = 0x8714
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	MAX_WIDTH                                                  = 0x827E
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	ALPHA                                                      = 0x1906
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	UNSIGNED_INT64_NV                                          = 0x140F
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	R8UI                                                       = 0x8232
	READ_ONLY_ARB                                              = 0x88B8
	REG_18_ATI                                                 = 0x8933
	TEXTURE_ENV                                                = 0x2300
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	SYNC_FLAGS_APPLE                                           = 0x9115
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	SUBTRACT                                                   = 0x84E7
	VERTEX_STREAM5_ATI                                         = 0x8771
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	RGB32I_EXT                                                 = 0x8D83
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	PROGRAM_STRING_NV                                          = 0x8628
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	PRIMARY_COLOR                                              = 0x8577
	LUMINANCE8UI_EXT                                           = 0x8D80
	DEPTH_TEST                                                 = 0x0B71
	DEPTH_COMPONENTS                                           = 0x8284
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	HISTOGRAM_SINK                                             = 0x802D
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	STENCIL_INDEX16_EXT                                        = 0x8D49
	SYNC_FLAGS                                                 = 0x9115
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	R16_SNORM                                                  = 0x8F98
	TRANSFORM_BIT                                              = 0x00001000
	AND_INVERTED                                               = 0x1504
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	FUNC_ADD                                                   = 0x8006
	TESS_CONTROL_TEXTURE                                       = 0x829C
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	BUMP_TARGET_ATI                                            = 0x877C
	SAMPLER_3D_ARB                                             = 0x8B5F
	CLIP_DISTANCE_NV                                           = 0x8C7A
	DEPTH_COMPONENT32F                                         = 0x8CAC
	STENCIL_INDEX16                                            = 0x8D49
	MATRIX4_NV                                                 = 0x8634
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	DOT3_ATI                                                   = 0x8966
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	VERSION_1_2                                                = 1
	MAX_LIST_NESTING                                           = 0x0B31
	VERTEX_ARRAY_TYPE                                          = 0x807B
	INT64_NV                                                   = 0x140E
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	MINMAX                                                     = 0x802E
	MIRRORED_REPEAT                                            = 0x8370
	PASS_THROUGH_NV                                            = 0x86E6
	RGB5_A1_OES                                                = 0x8057
	CURRENT_TANGENT_EXT                                        = 0x843B
	NORMAL_MAP_EXT                                             = 0x8511
	VIBRANCE_BIAS_NV                                           = 0x8719
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	DISCARD_ATI                                                = 0x8763
	STENCIL_INDEX4_EXT                                         = 0x8D47
	EYE_LINEAR                                                 = 0x2400
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	SIGNED_LUMINANCE_NV                                        = 0x8701
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	COLOR_ATTACHMENT13                                         = 0x8CED
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	EXT_rescale_normal                                         = 1
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	STACK_UNDERFLOW                                            = 0x0504
	SRC_COLOR                                                  = 0x0300
	COMPRESSED_RG                                              = 0x8226
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	CON_15_ATI                                                 = 0x8950
	BOOL_VEC4                                                  = 0x8B59
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	BLEND_EQUATION_RGB                                         = 0x8009
	ALPHA4_EXT                                                 = 0x803B
	LUMINANCE16F_ARB                                           = 0x881E
	CURRENT_BIT                                                = 0x00000001
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	RGBA_DXT5_S3TC                                             = 0x83A4
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	ALPHA_TEST_REF                                             = 0x0BC2
	POINT_SIZE_MAX_EXT                                         = 0x8127
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	VIEW_CLASS_48_BITS                                         = 0x82C7
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	CLIP_DISTANCE4                                             = 0x3004
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	DRAW_BUFFER10                                              = 0x882F
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	COUNT_DOWN_NV                                              = 0x9089
	GL_4PASS_0_SGIS                                            = 0x80A4
	DEPTH_BOUNDS_EXT                                           = 0x8891
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	LAYER_NV                                                   = 0x8DAA
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	COMBINE_EXT                                                = 0x8570
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	ARRAY_SIZE                                                 = 0x92FB
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	SGIS_generate_mipmap                                       = 1
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	UNSIGNED_NORMALIZED                                        = 0x8C17
	BLEND_SRC_ALPHA                                            = 0x80CB
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	DRAW_BUFFER10_NV                                           = 0x882F
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	POINT_BIT                                                  = 0x00000002
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	FOG_HINT                                                   = 0x0C54
	SELECT                                                     = 0x1C02
	SEPARABLE_2D                                               = 0x8012
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	ALL_COMPLETED_NV                                           = 0x84F2
	RGB_422_APPLE                                              = 0x8A1F
	FLOAT_VEC2                                                 = 0x8B50
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	OP_MAX_EXT                                                 = 0x878A
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	ETC1_RGB8_OES                                              = 0x8D64
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	TEXTURE_SHADOW                                             = 0x82A1
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	RGBA4_S3TC                                                 = 0x83A3
	CLIP_PLANE4                                                = 0x3004
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	DOT4_ATI                                                   = 0x8967
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	MODELVIEW                                                  = 0x1700
	TEXTURE8_ARB                                               = 0x84C8
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	DISCRETE_AMD                                               = 0x9006
	EDGE_FLAG                                                  = 0x0B43
	BGR                                                        = 0x80E0
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	CURRENT_BINORMAL_EXT                                       = 0x843C
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	OP_RECIP_SQRT_EXT                                          = 0x8795
	RGB16F_ARB                                                 = 0x881B
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	DEPTH24_STENCIL8                                           = 0x88F0
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	TEXTURE_HEIGHT                                             = 0x1001
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	DRAW_BUFFER13                                              = 0x8832
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	COPY_WRITE_BUFFER                                          = 0x8F37
	LUMINANCE4                                                 = 0x803F
	TEXTURE6                                                   = 0x84C6
	RGB_SCALE_EXT                                              = 0x8573
	SAMPLES_3DFX                                               = 0x86B4
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	POLYGON_OFFSET_FILL                                        = 0x8037
	RGBA16_EXT                                                 = 0x805B
	SAMPLE_MASK_SGIS                                           = 0x80A0
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	MAGNITUDE_SCALE_NV                                         = 0x8712
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	BLUE_SCALE                                                 = 0x0D1A
	RGB4_EXT                                                   = 0x804F
	R16UI                                                      = 0x8234
	RG8I                                                       = 0x8237
	TEXTURE17_ARB                                              = 0x84D1
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	OP_FLOOR_EXT                                               = 0x878F
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	OUT_OF_MEMORY                                              = 0x0505
	RGBA_FLOAT16_ATI                                           = 0x881A
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	PRESENT_DURATION_NV                                        = 0x8E2B
	EXT_polygon_offset                                         = 1
	LINE_WIDTH                                                 = 0x0B21
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	COLOR_SUM                                                  = 0x8458
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	CONVEX_HULL_NV                                             = 0x908B
	TEXTURE_GEN_Q                                              = 0x0C63
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	TEXTURE19                                                  = 0x84D3
	INTERLACE_READ_OML                                         = 0x8981
	ALPHA16                                                    = 0x803E
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	RGBA_INTEGER                                               = 0x8D99
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	DEBUG_OBJECT_MESA                                          = 0x8759
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	OPERAND2_ALPHA_EXT                                         = 0x859A
	SET_AMD                                                    = 0x874A
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	DRAW_BUFFER15_NV                                           = 0x8834
	COLOR_INDEX4_EXT                                           = 0x80E4
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	SURFACE_REGISTERED_NV                                      = 0x86FD
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	DEPTH_COMPONENT32                                          = 0x81A7
	MAX_HEIGHT                                                 = 0x827F
	MODELVIEW7_ARB                                             = 0x8727
	DS_SCALE_NV                                                = 0x8710
	DRAW_BUFFER3_NV                                            = 0x8828
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	POINT_SMOOTH                                               = 0x0B10
	TEXTURE11                                                  = 0x84CB
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	OP_RECIP_EXT                                               = 0x8794
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	EXPAND_NORMAL_NV                                           = 0x8538
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	HALF_FLOAT                                                 = 0x140B
	TABLE_TOO_LARGE                                            = 0x8031
	CUBIC_EXT                                                  = 0x8334
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	TRIANGULAR_NV                                              = 0x90A5
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	MATRIX27_ARB                                               = 0x88DB
	LUMINANCE16_SNORM                                          = 0x9019
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	SPRITE_AXIAL_SGIX                                          = 0x814C
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	MATRIX18_ARB                                               = 0x88D2
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	IMAGE_2D                                                   = 0x904D
	DEBUG_SEVERITY_LOW                                         = 0x9148
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	SRGB_WRITE                                                 = 0x8298
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	INTENSITY8_SNORM                                           = 0x9017
	TEXTURE_BINDING_2D                                         = 0x8069
	CLIP_PLANE5                                                = 0x3005
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	MAX_LIGHTS                                                 = 0x0D31
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	REG_24_ATI                                                 = 0x8939
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	OUTPUT_FOG_EXT                                             = 0x87BD
	REG_12_ATI                                                 = 0x892D
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	RGBA32I_EXT                                                = 0x8D82
	TEXTURE_GEN_S                                              = 0x0C60
	SOURCE2_ALPHA_ARB                                          = 0x858A
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	STENCIL_INDEX8_OES                                         = 0x8D48
	NEXT_BUFFER_NV                                             = -2
	SKIP_COMPONENTS4_NV                                        = -3
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	IMAGE_2D_ARRAY                                             = 0x9053
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	BLEND_SRC_RGB                                              = 0x80C9
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	BYTE                                                       = 0x1400
	LIGHT0                                                     = 0x4000
	RGB8_EXT                                                   = 0x8051
	COLOR_ARRAY_EXT                                            = 0x8076
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	DRAW_BUFFER7                                               = 0x882C
	BEVEL_NV                                                   = 0x90A6
	GL_2D                                                      = 0x0600
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	LEQUAL                                                     = 0x0203
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	BOOL_VEC2                                                  = 0x8B57
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	LINE                                                       = 0x1B01
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	TESSELLATION_MODE_AMD                                      = 0x9004
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	COLOR_ATTACHMENT15                                         = 0x8CEF
	MATRIX_MODE                                                = 0x0BA0
	SOURCE0_RGB                                                = 0x8580
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	INTENSITY4_EXT                                             = 0x804A
	POINT_SIZE_MAX                                             = 0x8127
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	PURGEABLE_APPLE                                            = 0x8A1D
	SRGB8_NV                                                   = 0x8C41
	GL_2PASS_1_SGIS                                            = 0x80A3
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	DOT_PRODUCT_NV                                             = 0x86EC
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	CONSTANT_EXT                                               = 0x8576
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	CURRENT_QUERY_ARB                                          = 0x8865
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	PIXEL_COUNT_NV                                             = 0x8866
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	SAMPLE_COVERAGE                                            = 0x80A0
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	CON_28_ATI                                                 = 0x895D
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	SGIX_async_pixel                                           = 1
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	BOOL_ARB                                                   = 0x8B56
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	COMPRESSED_RG11_EAC                                        = 0x9272
	TEXTURE14                                                  = 0x84CE
	DRAW_BUFFER2_ARB                                           = 0x8827
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	PALETTE4_RGBA8_OES                                         = 0x8B91
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	RED_SNORM                                                  = 0x8F90
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	SIGNED_HILO8_NV                                            = 0x885F
	TEXTURE_BUFFER                                             = 0x8C2A
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	ACCUM_RED_BITS                                             = 0x0D58
	LUMINANCE16_ALPHA16                                        = 0x8048
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	SLUMINANCE_ALPHA                                           = 0x8C44
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	TEXTURE_3D                                                 = 0x806F
	TEXTURE25_ARB                                              = 0x84D9
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	DRAW_BUFFER14_NV                                           = 0x8833
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	TEXTURE_WIDTH                                              = 0x1000
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	QUERY_RESULT_EXT                                           = 0x8866
	CURRENT_TIME_NV                                            = 0x8E28
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	COMBINER4_NV                                               = 0x8554
	PROGRAM_PARAMETER_NV                                       = 0x8644
	UNPACK_RESAMPLE_OML                                        = 0x8985
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	LINEAR                                                     = 0x2601
	RGBA4_OES                                                  = 0x8056
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	SGIX_clipmap                                               = 1
	MAP_READ_BIT_EXT                                           = 0x0001
	POINT_SIZE_MIN_ARB                                         = 0x8126
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	MODELVIEW19_ARB                                            = 0x8733
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	FLOAT_MAT4_ARB                                             = 0x8B5C
	HISTOGRAM_SINK_EXT                                         = 0x802D
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	T2F_V3F                                                    = 0x2A27
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	DOT3_RGBA                                                  = 0x86AF
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	TEXTURE20_ARB                                              = 0x84D4
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	Q                                                          = 0x2003
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	AVERAGE_HP                                                 = 0x8160
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	LOWER_LEFT                                                 = 0x8CA1
	FOG_BIT                                                    = 0x00000080
	INDEX_ARRAY                                                = 0x8077
	TEXTURE_MIN_FILTER                                         = 0x2801
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	INTENSITY16_SNORM                                          = 0x901B
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	WAIT_FAILED                                                = 0x911D
	MAP_READ_BIT                                               = 0x0001
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	Z400_BINARY_AMD                                            = 0x8740
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	FRONT_FACE                                                 = 0x0B46
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	SPRITE_MODE_SGIX                                           = 0x8149
	DITHER                                                     = 0x0BD0
	SOURCE2_ALPHA_EXT                                          = 0x858A
	FOG                                                        = 0x0B60
	GENERATE_MIPMAP                                            = 0x8191
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	COLOR                                                      = 0x1800
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	ALPHA_BIAS                                                 = 0x0D1D
	RED_BITS                                                   = 0x0D52
	TEXTURE15                                                  = 0x84CF
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	TEXTURE_MAX_LOD                                            = 0x813B
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	COLOR_LOGIC_OP                                             = 0x0BF2
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	BLEND_EQUATION_ALPHA                                       = 0x883D
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	ACTIVE_VARIABLES                                           = 0x9305
	REDUCE_EXT                                                 = 0x8016
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	SLUMINANCE8                                                = 0x8C47
	TEXTURE_BORDER                                             = 0x1005
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	AMBIENT                                                    = 0x1200
	DOUBLE_MAT2x3                                              = 0x8F49
	VERTEX_SHADER_BIT                                          = 0x00000001
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	MODELVIEW5_ARB                                             = 0x8725
	MATRIX16_ARB                                               = 0x88D0
	TRUE                                                       = 1
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	PALETTE8_RGBA4_OES                                         = 0x8B98
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	RGBA                                                       = 0x1908
	OUTPUT_COLOR1_EXT                                          = 0x879C
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	OP_ROUND_EXT                                               = 0x8790
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	SGI_color_table                                            = 1
	SPARE1_NV                                                  = 0x852F
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	DEPTH_STENCIL_OES                                          = 0x84F9
	SOURCE0_ALPHA_EXT                                          = 0x8588
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	STENCIL_BACK_REF                                           = 0x8CA3
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	UNIFORM_SIZE                                               = 0x8A38
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	PATH_FORMAT_SVG_NV                                         = 0x9070
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	NEGATIVE_Y_EXT                                             = 0x87DA
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	LUMINANCE4_ALPHA4                                          = 0x8043
	DRAW_BUFFER14_ARB                                          = 0x8833
	REG_21_ATI                                                 = 0x8936
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	FIELDS_NV                                                  = 0x8E27
	MODELVIEW30_ARB                                            = 0x873E
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	REG_17_ATI                                                 = 0x8932
	CON_29_ATI                                                 = 0x895E
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	BLOCK_INDEX                                                = 0x92FD
	INTENSITY32I_EXT                                           = 0x8D85
	SGIX_shadow                                                = 1
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	MEDIUM_FLOAT                                               = 0x8DF1
	CURRENT_COLOR                                              = 0x0B00
	SAMPLES                                                    = 0x80A9
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	MATRIX3_ARB                                                = 0x88C3
	HALF_FLOAT_OES                                             = 0x8D61
	SKIP_COMPONENTS1_NV                                        = -6
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	TYPE                                                       = 0x92FA
	COLOR_ARRAY_TYPE                                           = 0x8082
	GENERATE_MIPMAP_HINT                                       = 0x8192
	ELEMENT_ARRAY_ATI                                          = 0x8768
	MATRIX_PALETTE_OES                                         = 0x8840
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	PRIMITIVE_RESTART                                          = 0x8F9D
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	TEXTURE_COMPARE_MODE                                       = 0x884C
	RENDERBUFFER                                               = 0x8D41
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	CURRENT_RASTER_INDEX                                       = 0x0B05
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	LUMINANCE16_EXT                                            = 0x8042
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	MIRROR_CLAMP_EXT                                           = 0x8742
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	RG16F                                                      = 0x822F
	COLOR_RENDERABLE                                           = 0x8286
	RGB9_E5                                                    = 0x8C3D
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	DUAL_ALPHA16_SGIS                                          = 0x8113
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	SGIX_ir_instrument1                                        = 1
	PROXY_TEXTURE_2D                                           = 0x8064
	SAMPLES_SGIS                                               = 0x80A9
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	COMBINE                                                    = 0x8570
	MODELVIEW25_ARB                                            = 0x8739
	VERTEX_STREAM0_ATI                                         = 0x876C
	CON_3_ATI                                                  = 0x8944
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	PACK_ALIGNMENT                                             = 0x0D05
	EMISSION                                                   = 0x1600
	OBJECT_POINT_SGIS                                          = 0x81F5
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	SIGNED_INTENSITY8_NV                                       = 0x8708
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	IMAGE_CUBE_EXT                                             = 0x9050
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	MATRIX30_ARB                                               = 0x88DE
	COLOR_ATTACHMENT6                                          = 0x8CE6
	IMAGE_3D_EXT                                               = 0x904E
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	SHADER                                                     = 0x82E1
	TEXTURE30                                                  = 0x84DE
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	TEXTURE5_ARB                                               = 0x84C5
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	AUX3                                                       = 0x040C
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	DOUBLEBUFFER                                               = 0x0C32
	CLIP_DISTANCE7                                             = 0x3007
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	MATRIX9_ARB                                                = 0x88C9
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	SHADER_COMPILER                                            = 0x8DFA
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	RG32UI                                                     = 0x823C
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	RGBA16F                                                    = 0x881A
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	RGBA8UI_EXT                                                = 0x8D7C
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	DOMAIN                                                     = 0x0A02
	SLUMINANCE8_EXT                                            = 0x8C47
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	RGBA8_OES                                                  = 0x8058
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	COLOR_CLEAR_VALUE                                          = 0x0C22
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	YCBCR_422_APPLE                                            = 0x85B9
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	READ_WRITE_ARB                                             = 0x88BA
	INVALID_INDEX                                              = 0xFFFFFFFF
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	BLEND_DST_ALPHA                                            = 0x80CA
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	COMBINE_ALPHA_ARB                                          = 0x8572
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	REG_8_ATI                                                  = 0x8929
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	MAP_STENCIL                                                = 0x0D11
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	STORAGE_SHARED_APPLE                                       = 0x85BF
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	STATIC_ATI                                                 = 0x8760
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	GL_422_REV_EXT                                             = 0x80CD
	RGBA32I                                                    = 0x8D82
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	NORMAL_ARRAY_TYPE                                          = 0x807E
	LINEAR_DETAIL_SGIS                                         = 0x8097
	COLOR_TABLE_SGI                                            = 0x80D0
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	DRAW_BUFFER9_NV                                            = 0x882E
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	MITER_REVERT_NV                                            = 0x90A7
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	COLOR_ATTACHMENT2                                          = 0x8CE2
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	INTENSITY                                                  = 0x8049
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	DOUBLE_MAT4_EXT                                            = 0x8F48
	ARC_TO_NV                                                  = 0xFE
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	DT_SCALE_NV                                                = 0x8711
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	FOG_END                                                    = 0x0B64
	TEXTURE_GEN_T                                              = 0x0C61
	GREEN_BIT_ATI                                              = 0x00000002
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	TEXTURE_CUBE_MAP                                           = 0x8513
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	LINES_ADJACENCY_ARB                                        = 0x000A
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	SKIP_DECODE_EXT                                            = 0x8A4A
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	RGB16UI                                                    = 0x8D77
	OBJECT_TYPE_APPLE                                          = 0x9112
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	DRAW_BUFFER7_ARB                                           = 0x882C
	READ_ONLY                                                  = 0x88B8
	RELEASED_APPLE                                             = 0x8A19
	VERTEX_ARRAY_SIZE                                          = 0x807A
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	COMBINE4_NV                                                = 0x8503
	OP_SUB_EXT                                                 = 0x8796
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	PERTURB_EXT                                                = 0x85AE
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	HISTOGRAM_WIDTH                                            = 0x8026
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	CON_8_ATI                                                  = 0x8949
	QUAD_STRIP                                                 = 0x0008
	R8I                                                        = 0x8231
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	SGIX_subsample                                             = 1
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	SGIX_framezoom                                             = 1
	POLYGON_MODE                                               = 0x0B40
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	FOG_START                                                  = 0x0B63
	R16                                                        = 0x822A
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	EXT_vertex_array                                           = 1
	RED_SCALE                                                  = 0x0D14
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	GL_4PASS_0_EXT                                             = 0x80A4
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	VARIABLE_B_NV                                              = 0x8524
	WEIGHT_ARRAY_OES                                           = 0x86AD
	LO_SCALE_NV                                                = 0x870F
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	QUERY_RESULT_ARB                                           = 0x8866
	MODELVIEW10_ARB                                            = 0x872A
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	MAP1_TANGENT_EXT                                           = 0x8444
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	UNSIGNED_BYTE                                              = 0x1401
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	VERSION_3_1                                                = 1
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	OUTPUT_COLOR0_EXT                                          = 0x879B
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	INVERTED_SCREEN_W_REND                                     = 0x8491
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	SIGNED_HILO_NV                                             = 0x86F9
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	ROUND_NV                                                   = 0x90A4
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	ADD                                                        = 0x0104
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	RGBA4_DXT5_S3TC                                            = 0x83A5
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	YCBCR_MESA                                                 = 0x8757
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	ARRAY_STRIDE                                               = 0x92FE
	INVALID_ENUM                                               = 0x0500
	AND                                                        = 0x1501
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	CON_4_ATI                                                  = 0x8945
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	WAIT_FAILED_APPLE                                          = 0x911D
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	PRIMARY_COLOR_ARB                                          = 0x8577
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	FLOAT_RGB_NV                                               = 0x8882
	MODELVIEW_MATRIX                                           = 0x0BA6
	OR                                                         = 0x1507
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	CONSTANT_BORDER_HP                                         = 0x8151
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	SIGNED_RGB_NV                                              = 0x86FE
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	AUX1                                                       = 0x040A
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	DRAW_BUFFER10_ARB                                          = 0x882F
	INTERLACE_OML                                              = 0x8980
	INTENSITY16                                                = 0x804D
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	PRIMARY_COLOR_NV                                           = 0x852C
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	IMAGE_BINDING_NAME                                         = 0x8F3A
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	SINGLE_COLOR_EXT                                           = 0x81F9
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	PATCH_VERTICES                                             = 0x8E72
	LOGIC_OP_MODE                                              = 0x0BF0
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	EMBOSS_LIGHT_NV                                            = 0x855D
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	DSDT8_NV                                                   = 0x8709
	DYNAMIC_READ                                               = 0x88E9
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	REG_27_ATI                                                 = 0x893C
	TESS_GEN_SPACING                                           = 0x8E77
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	LUMINANCE16                                                = 0x8042
	LUMINANCE12_EXT                                            = 0x8041
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	DRAW_BUFFER12_ATI                                          = 0x8831
	CON_17_ATI                                                 = 0x8952
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	SUCCESS_NV                                                 = 0x902F
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	QUERY_NO_WAIT_NV                                           = 0x8E14
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	TEXTURE_MAG_FILTER                                         = 0x2800
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	REG_11_ATI                                                 = 0x892C
	LINE_WIDTH_RANGE                                           = 0x0B22
	TEXTURE_SHADER_NV                                          = 0x86DE
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	QUADS                                                      = 0x0007
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	NEGATIVE_ONE_EXT                                           = 0x87DF
	YCBAYCR8A_4224_NV                                          = 0x9032
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	SOURCE1_RGB_ARB                                            = 0x8581
	MODELVIEW9_ARB                                             = 0x8729
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	RED_BIT_ATI                                                = 0x00000001
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	EVAL_BIT                                                   = 0x00010000
	COLOR_INDEX12_EXT                                          = 0x80E6
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	COMP_BIT_ATI                                               = 0x00000002
	INT_IMAGE_1D_EXT                                           = 0x9057
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	EMBOSS_CONSTANT_NV                                         = 0x855E
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	TEXTURE24_ARB                                              = 0x84D8
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	OP_LOG_BASE_2_EXT                                          = 0x8792
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	ACCUM_BUFFER_BIT                                           = 0x00000200
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	MODELVIEW21_ARB                                            = 0x8735
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	TESS_GEN_MODE                                              = 0x8E76
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	TEXTURE_BINDING_3D                                         = 0x806A
	VERSION                                                    = 0x1F02
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	FIRST_TO_REST_NV                                           = 0x90AF
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	BOOL_VEC3_ARB                                              = 0x8B58
	COLOR_ATTACHMENT14                                         = 0x8CEE
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	COMPUTE_SHADER                                             = 0x91B9
	DECAL                                                      = 0x2101
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	DSDT8_MAG8_NV                                              = 0x870A
	UNIFORM_OFFSET                                             = 0x8A3B
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	INT_IMAGE_BUFFER                                           = 0x905C
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	PATH_DASH_OFFSET_NV                                        = 0x907E
	PATH_FILL_MASK_NV                                          = 0x9081
	COLOR_ATTACHMENT9                                          = 0x8CE9
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	INT64_VEC4_NV                                              = 0x8FEB
	UTF16_NV                                                   = 0x909B
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	EIGHTH_BIT_ATI                                             = 0x00000020
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	SPECULAR                                                   = 0x1202
	OR_REVERSE                                                 = 0x150B
	DEBUG_SOURCE_OTHER                                         = 0x824B
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	POINT_SIZE                                                 = 0x0B11
	MULTISAMPLE_EXT                                            = 0x809D
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	DOUBLE_EXT                                                 = 0x140A
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	COLOR_ATTACHMENT4                                          = 0x8CE4
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	FUNC_SUBTRACT                                              = 0x800A
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	TEXTURE20                                                  = 0x84D4
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	LUMINANCE16I_EXT                                           = 0x8D8C
	SGIX_convolution_accuracy                                  = 1
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	COLOR_ARRAY_STRIDE                                         = 0x8083
	COMPRESSED_RED                                             = 0x8225
	RENDER                                                     = 0x1C00
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	DEPTH_CLAMP                                                = 0x864F
	INT_VEC3_ARB                                               = 0x8B54
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	MAP1_NORMAL                                                = 0x0D92
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	COLOR_TABLE_BIAS                                           = 0x80D7
	MAX_DEPTH                                                  = 0x8280
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	CONSTANT_ALPHA_EXT                                         = 0x8003
	REPLICATE_BORDER_HP                                        = 0x8153
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	FLOAT_R16_NV                                               = 0x8884
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	OPERAND2_RGB_ARB                                           = 0x8592
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	DISCARD_NV                                                 = 0x8530
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	TEXTURE_DEPTH                                              = 0x8071
	MATRIX11_ARB                                               = 0x88CB
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	SGIX_list_priority                                         = 1
	CURRENT_FOG_COORDINATE                                     = 0x8453
	TEXTURE_COMPRESSED                                         = 0x86A1
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	CLEAR                                                      = 0x1500
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	EXT_texture                                                = 1
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	V2F                                                        = 0x2A20
	SCALE_BY_FOUR_NV                                           = 0x853F
	MAX_TEXTURE_SIZE                                           = 0x0D33
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	INT8_VEC4_NV                                               = 0x8FE3
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	UNSIGNED_INT                                               = 0x1405
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	VERTEX4_BIT_PGI                                            = 0x00000008
	RGBA8                                                      = 0x8058
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	IMAGE_BINDING_FORMAT                                       = 0x906E
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	MVP_MATRIX_EXT                                             = 0x87E3
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	PROGRAM_LENGTH_ARB                                         = 0x8627
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	DRAW_BUFFER13_ARB                                          = 0x8832
	RGB565_OES                                                 = 0x8D62
	SGIS_point_line_texgen                                     = 1
	VERTEX_SHADER                                              = 0x8B31
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	DOUBLE_VEC2                                                = 0x8FFC
	PROGRAM_OUTPUT                                             = 0x92E4
	COMPUTE_TEXTURE                                            = 0x82A0
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	RGBA32UI_EXT                                               = 0x8D70
	VERTEX23_BIT_PGI                                           = 0x00000004
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	STENCIL                                                    = 0x1802
	BLEND_EQUATION_EXT                                         = 0x8009
	IMAGE_SCALE_Y_HP                                           = 0x8156
	TEXTURE10_ARB                                              = 0x84CA
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	PACK_SKIP_ROWS                                             = 0x0D03
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	SOURCE1_RGB_EXT                                            = 0x8581
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	MAX_SAMPLES_NV                                             = 0x8D57
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	DEPTH_BIAS                                                 = 0x0D1F
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	VERTEX_STREAM2_ATI                                         = 0x876E
	COMPILE_STATUS                                             = 0x8B81
	FILE_NAME_NV                                               = 0x9074
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	FIXED                                                      = 0x140C
	IGNORE_BORDER_HP                                           = 0x8150
	MAX_IMAGE_SAMPLES                                          = 0x906D
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	SHADER_OBJECT_ARB                                          = 0x8B48
	COMPILE_AND_EXECUTE                                        = 0x1301
	RGBA2                                                      = 0x8055
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	PASS_THROUGH_TOKEN                                         = 0x0700
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	REG_3_ATI                                                  = 0x8924
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	CLIP_PLANE0                                                = 0x3000
	REG_0_ATI                                                  = 0x8921
	CON_5_ATI                                                  = 0x8946
	BLUE_INTEGER                                               = 0x8D96
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	BUFFER_MAP_LENGTH                                          = 0x9120
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	MIPMAP                                                     = 0x8293
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	UPPER_LEFT                                                 = 0x8CA2
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	FLOAT_MAT2x3                                               = 0x8B65
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	FLOAT16_VEC3_NV                                            = 0x8FFA
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	COEFF                                                      = 0x0A00
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	TEXTURE_MAX_LEVEL                                          = 0x813D
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	SAMPLER_CUBE                                               = 0x8B60
	RGBA8I                                                     = 0x8D8E
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	V3F                                                        = 0x2A21
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	SRC1_COLOR                                                 = 0x88F9
	ALPHA_INTEGER_EXT                                          = 0x8D97
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	LIST_MODE                                                  = 0x0B30
	MINMAX_FORMAT_EXT                                          = 0x802F
	READ_PIXELS                                                = 0x828C
	TEXTURE7_ARB                                               = 0x84C7
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	Z_EXT                                                      = 0x87D7
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	MAX_NAME_LENGTH                                            = 0x92F6
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	COLOR_ARRAY_LIST_IBM                                       = 103072
	UNSIGNED_SHORT                                             = 0x1403
	RG16F_EXT                                                  = 0x822F
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	MODELVIEW8_ARB                                             = 0x8728
	NEGATIVE_X_EXT                                             = 0x87D9
	GREEN_SCALE                                                = 0x0D18
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	SGIX_fog_offset                                            = 1
	SGI_color_matrix                                           = 1
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	UNIFORM_BUFFER_START                                       = 0x8A29
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	POINT_SIZE_RANGE                                           = 0x0B12
	DEPTH_WRITEMASK                                            = 0x0B72
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	SLUMINANCE_NV                                              = 0x8C46
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	SRC_ALPHA_SATURATE                                         = 0x0308
	TEXTURE_DEPTH_EXT                                          = 0x8071
	DYNAMIC_DRAW                                               = 0x88E8
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	MODELVIEW17_ARB                                            = 0x8731
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	CULL_VERTEX_IBM                                            = 103050
	OPERAND1_RGB                                               = 0x8591
	VERTEX_PROGRAM_ARB                                         = 0x8620
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	NORMAL_ARRAY_POINTER                                       = 0x808F
	RED_EXT                                                    = 0x1903
	DEBUG_SOURCE_API                                           = 0x8246
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	FRACTIONAL_ODD                                             = 0x8E7B
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	SGIX_texture_coordinate_clamp                              = 1
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	RGB8_SNORM                                                 = 0x8F96
	SIGNALED                                                   = 0x9119
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	PACK_ROW_LENGTH                                            = 0x0D02
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	TEXTURE3                                                   = 0x84C3
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	C4UB_V2F                                                   = 0x2A22
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	ALPHA_BITS                                                 = 0x0D55
	BITMAP                                                     = 0x1A00
	BGRA_EXT                                                   = 0x80E1
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	SGIS_multisample                                           = 1
	SGIS_texture_select                                        = 1
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	GL_3DC_XY_AMD                                              = 0x87FA
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	MAX_SUBROUTINES                                            = 0x8DE7
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MULT                                                       = 0x0103
	RGBA32F                                                    = 0x8814
	NOOP                                                       = 0x1505
	BLEND_COLOR_EXT                                            = 0x8005
	TEXTURE_LIGHT_EXT                                          = 0x8350
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	COLOR_SUM_EXT                                              = 0x8458
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	PRIMITIVE_ID_NV                                            = 0x8C7C
	RGBA16I_EXT                                                = 0x8D88
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	VIEWPORT_BIT                                               = 0x00000800
	COLOR_TABLE                                                = 0x80D0
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	STATIC_READ_ARB                                            = 0x88E5
	REG_25_ATI                                                 = 0x893A
	RASTERIZER_DISCARD                                         = 0x8C89
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	LINE_TOKEN                                                 = 0x0702
	RGBA4                                                      = 0x8056
	TEXTURE0_ARB                                               = 0x84C0
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	MODELVIEW0_EXT                                             = 0x1700
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	SGIS_texture_edge_clamp                                    = 1
	IUI_N3F_V3F_EXT                                            = 0x81B0
	COMPRESSED_RGB_ARB                                         = 0x84ED
	OPERAND3_ALPHA_NV                                          = 0x859B
	CULL_FRAGMENT_NV                                           = 0x86E7
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	T4F_V4F                                                    = 0x2A28
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	REG_6_ATI                                                  = 0x8927
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	REDUCE                                                     = 0x8016
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	EYE_POINT_SGIS                                             = 0x81F4
	PN_TRIANGLES_ATI                                           = 0x87F0
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	DOT3_RGB                                                   = 0x86AE
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	REPLACE_EXT                                                = 0x8062
	GL_2PASS_1_EXT                                             = 0x80A3
	OBJECT_LINE_SGIS                                           = 0x81F7
	COLOR_TABLE_FORMAT                                         = 0x80D8
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	VIEW_CLASS_24_BITS                                         = 0x82C9
	INTENSITY8UI_EXT                                           = 0x8D7F
	LINES_ADJACENCY_EXT                                        = 0x000A
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	SHININESS                                                  = 0x1601
	MIN                                                        = 0x8007
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	COLOR_ATTACHMENT5                                          = 0x8CE5
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	DIFFUSE                                                    = 0x1201
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	POINT_SPRITE_ARB                                           = 0x8861
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	CULL_FACE                                                  = 0x0B44
	POINT_SIZE_MIN_EXT                                         = 0x8126
	COMBINER_BIAS_NV                                           = 0x8549
	POINT_SPRITE_NV                                            = 0x8861
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	POINT_SMOOTH_HINT                                          = 0x0C51
	VENDOR                                                     = 0x1F00
	QUAD_ALPHA8_SGIS                                           = 0x811F
	COLOR_ATTACHMENT12                                         = 0x8CEC
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	FLOAT_RGB16_NV                                             = 0x8888
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	BGRA_INTEGER_EXT                                           = 0x8D9B
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	POLYGON_SMOOTH                                             = 0x0B41
	OP_ADD_EXT                                                 = 0x8787
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	RGB10_A2UI                                                 = 0x906F
	EXP                                                        = 0x0800
	INCR                                                       = 0x1E02
	DRAW_BUFFER11_NV                                           = 0x8830
	STREAM_COPY                                                = 0x88E2
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	SGIX_reference_plane                                       = 1
	R16F_EXT                                                   = 0x822D
	PRESERVE_ATI                                               = 0x8762
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	CON_12_ATI                                                 = 0x894D
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	FUNC_ADD_EXT                                               = 0x8006
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	FLOAT16_VEC2_NV                                            = 0x8FF9
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	NEAREST                                                    = 0x2600
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	LIGHTING                                                   = 0x0B50
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	DEBUG_TYPE_MARKER                                          = 0x8268
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	SGIX_pixel_texture                                         = 1
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	TEXTURE21_ARB                                              = 0x84D5
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	DRAW_BUFFER12_ARB                                          = 0x8831
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	MATRIX_EXT                                                 = 0x87C0
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	EQUIV                                                      = 0x1509
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	TEXTURE15_ARB                                              = 0x84CF
	NORMAL_BIT_PGI                                             = 0x08000000
	PROJECTION_MATRIX                                          = 0x0BA7
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	FLOAT16_VEC4_NV                                            = 0x8FFB
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	CUBIC_CURVE_TO_NV                                          = 0x0C
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	DUAL_ALPHA4_SGIS                                           = 0x8110
	CLEAR_BUFFER                                               = 0x82B4
	DISPLAY_LIST                                               = 0x82E7
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	PACK_RESAMPLE_SGIX                                         = 0x842C
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	PATH_GEN_COEFF_NV                                          = 0x90B1
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	FALSE                                                      = 0
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	SRGB8_EXT                                                  = 0x8C41
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	SAMPLE_MASK_NV                                             = 0x8E51
	HISTOGRAM_RED_SIZE                                         = 0x8028
	TEXTURE13                                                  = 0x84CD
	FOG_MODE                                                   = 0x0B65
	DRAW_BUFFER_EXT                                            = 0x0C01
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	SGIS_point_parameters                                      = 1
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	COLOR_TABLE_WIDTH                                          = 0x80D9
	MIRRORED_REPEAT_IBM                                        = 0x8370
	MULTISAMPLE_3DFX                                           = 0x86B2
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	CON_18_ATI                                                 = 0x8953
	BUFFER_DATA_SIZE                                           = 0x9303
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	RGBA8UI                                                    = 0x8D7C
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	VERTEX_STREAM3_ATI                                         = 0x876F
	SEPARATE_ATTRIBS                                           = 0x8C8D
	INT16_VEC4_NV                                              = 0x8FE7
	SGIX_texture_scale_bias                                    = 1
	GREEN                                                      = 0x1904
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	PATH_JOIN_STYLE_NV                                         = 0x9079
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	SPOT_CUTOFF                                                = 0x1206
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	INT64_VEC2_NV                                              = 0x8FE9
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	DOUBLE_MAT3x2                                              = 0x8F4B
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	NAME_LENGTH                                                = 0x92F9
	SGIS_texture_filter4                                       = 1
	TEXTURE_BIT                                                = 0x00040000
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	SPOT_EXPONENT                                              = 0x1205
	CONVOLUTION_1D_EXT                                         = 0x8010
	STATIC_READ                                                = 0x88E5
	IMAGE_2D_RECT                                              = 0x904F
	CONSTANT_COLOR_EXT                                         = 0x8001
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	INT16_VEC3_NV                                              = 0x8FE6
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	REG_9_ATI                                                  = 0x892A
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	GL_4PASS_3_EXT                                             = 0x80A7
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	COMBINER_MUX_SUM_NV                                        = 0x8547
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	DRAW_BUFFER9                                               = 0x882E
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	FASTEST                                                    = 0x1101
	MINMAX_SINK_EXT                                            = 0x8030
	SOURCE0_RGB_ARB                                            = 0x8580
	DOT3_RGBA_EXT                                              = 0x8741
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	ACCUM_GREEN_BITS                                           = 0x0D59
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	VERTEX_TEXTURE                                             = 0x829B
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	SAMPLER_2D_ARB                                             = 0x8B5E
	PARTIAL_SUCCESS_NV                                         = 0x902E
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	AUX2                                                       = 0x040B
	INDEX_BITS                                                 = 0x0D51
	ALPHA12_EXT                                                = 0x803D
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	UNDEFINED_VERTEX                                           = 0x8260
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	TEXTURE4                                                   = 0x84C4
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	TEXTURE31                                                  = 0x84DF
	TEXTURE_COORD_ARRAY                                        = 0x8078
	CON_21_ATI                                                 = 0x8956
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	MODELVIEW20_ARB                                            = 0x8734
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	GL_2X_BIT_ATI                                              = 0x00000001
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	TRIANGLES_ADJACENCY                                        = 0x000C
	FRAGMENT_DEPTH                                             = 0x8452
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	DRAW_BUFFER13_NV                                           = 0x8832
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	TEXTURE_GEN_R                                              = 0x0C62
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	EXT_blend_subtract                                         = 1
	RGB_S3TC                                                   = 0x83A0
	PACK_RESAMPLE_OML                                          = 0x8984
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	LOCATION                                                   = 0x930E
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	T2F_N3F_V3F                                                = 0x2A2B
	CLIP_DISTANCE3                                             = 0x3003
	YCRCBA_SGIX                                                = 0x8319
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	TEXTURE22_ARB                                              = 0x84D6
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	FLOAT_MAT3x2                                               = 0x8B67
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	MULTISAMPLE_BIT                                            = 0x20000000
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	PACK_LSB_FIRST                                             = 0x0D01
	CONVOLUTION_WIDTH                                          = 0x8018
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	GL_4PASS_3_SGIS                                            = 0x80A7
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	FOG_COORDINATE_ARRAY                                       = 0x8457
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	BUFFER_VARIABLE                                            = 0x92E5
	PATCHES                                                    = 0x000E
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	MAX_DRAW_BUFFERS                                           = 0x8824
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	SYNC_FENCE_APPLE                                           = 0x9116
	PALETTE8_RGB8_OES                                          = 0x8B95
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	SGIX_sprite                                                = 1
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	DOT3_RGBA_ARB                                              = 0x86AF
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	PROGRAM_BINDING_ARB                                        = 0x8677
	SAMPLER_2D_SHADOW                                          = 0x8B62
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	STENCIL_RENDERABLE                                         = 0x8288
	OP_MIN_EXT                                                 = 0x878B
	CON_13_ATI                                                 = 0x894E
	SAMPLE_POSITION                                            = 0x8E50
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	GREEN_BITS                                                 = 0x0D53
	PACK_CMYK_HINT_EXT                                         = 0x800E
	MATRIX0_ARB                                                = 0x88C0
	MAX_IMAGE_UNITS                                            = 0x8F38
	C4UB_V3F                                                   = 0x2A23
	CONVOLUTION_HEIGHT                                         = 0x8019
	MIRROR_CLAMP_ATI                                           = 0x8742
	PALETTE8_RGBA8_OES                                         = 0x8B96
	LOW_INT                                                    = 0x8DF3
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	TEXTURE9_ARB                                               = 0x84C9
	MODELVIEW0_ARB                                             = 0x1700
	STREAM_DRAW_ARB                                            = 0x88E0
	SKIP_COMPONENTS2_NV                                        = -5
	TEXTURE_COMPONENTS                                         = 0x1003
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	RELATIVE_MOVE_TO_NV                                        = 0x03
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	GL_2PASS_0_EXT                                             = 0x80A2
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	SOURCE0_ALPHA                                              = 0x8588
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	RGB5_EXT                                                   = 0x8050
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	TEXTURE26_ARB                                              = 0x84DA
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	RGB8UI_EXT                                                 = 0x8D7D
	BOLD_BIT_NV                                                = 0x01
	SGIX_scalebias_hint                                        = 1
	RGB                                                        = 0x1907
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	TEXTURE12                                                  = 0x84CC
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	MATRIX28_ARB                                               = 0x88DC
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	QUERY_BUFFER_AMD                                           = 0x9192
	INTENSITY4                                                 = 0x804A
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	BLEND_EQUATION                                             = 0x8009
	TEXTURE17                                                  = 0x84D1
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	LIST_PRIORITY_SGIX                                         = 0x8182
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	BUFFER_MAPPED_ARB                                          = 0x88BC
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	DEPTH_RENDERABLE                                           = 0x8287
	HILO_NV                                                    = 0x86F4
	SGIX_texture_lod_bias                                      = 1
	SPHERE_MAP                                                 = 0x2402
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	MAP2_INDEX                                                 = 0x0DB1
	DEPTH_STENCIL_NV                                           = 0x84F9
	CONSTANT_COLOR0_NV                                         = 0x852A
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	BLEND                                                      = 0x0BE2
	LUMINANCE                                                  = 0x1909
	R32F                                                       = 0x822E
	MODULATE_ADD_ATI                                           = 0x8744
	NEGATE_BIT_ATI                                             = 0x00000004
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	SGI_texture_color_table                                    = 1
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	X_EXT                                                      = 0x87D5
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	SAMPLER_1D                                                 = 0x8B5D
	RGB16UI_EXT                                                = 0x8D77
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	SRC1_RGB                                                   = 0x8581
	OPERAND3_RGB_NV                                            = 0x8593
	VALIDATE_STATUS                                            = 0x8B83
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	SGIS_sharpen_texture                                       = 1
	DST_COLOR                                                  = 0x0306
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	ALPHA16F_ARB                                               = 0x881C
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	SAMPLE_MASK                                                = 0x8E51
	LUMINANCE_SNORM                                            = 0x9011
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	PREVIOUS_EXT                                               = 0x8578
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	SGIX_resample                                              = 1
	LINE_LOOP                                                  = 0x0002
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	SRC0_RGB                                                   = 0x8580
	PACK_INVERT_MESA                                           = 0x8758
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	NUM_SAMPLE_COUNTS                                          = 0x9380
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	REG_28_ATI                                                 = 0x893D
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	SGIS_texture4D                                             = 1
	BACK_RIGHT                                                 = 0x0403
	MATRIX20_ARB                                               = 0x88D4
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	MODELVIEW28_ARB                                            = 0x873C
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	TRANSLATE_X_NV                                             = 0x908E
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	PROGRAM_PIPELINE                                           = 0x82E4
	DRAW_BUFFER8_NV                                            = 0x882D
	ARRAY_BUFFER_BINDING                                       = 0x8894
	READ_WRITE                                                 = 0x88BA
	PERFMON_RESULT_AMD                                         = 0x8BC6
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	TRIANGLES                                                  = 0x0004
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	LUMINANCE12_ALPHA12                                        = 0x8047
	CND_ATI                                                    = 0x896A
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	FOG_COORD_SRC                                              = 0x8450
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	DRAW_BUFFER0_NV                                            = 0x8825
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	FLOAT_RG32_NV                                              = 0x8887
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	ONE_MINUS_DST_COLOR                                        = 0x0307
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	LUMINANCE8_SNORM                                           = 0x9015
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	POLYGON_OFFSET_EXT                                         = 0x8037
	GEOMETRY_SHADER                                            = 0x8DD9
	FRAME_NV                                                   = 0x8E26
	SMALL_CW_ARC_TO_NV                                         = 0x14
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	MODULATE                                                   = 0x2100
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	COLOR_COMPONENTS                                           = 0x8283
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	STREAM_READ                                                = 0x88E1
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	ACCUM                                                      = 0x0100
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	TEXTURE_VIEW                                               = 0x82B5
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	INVALID_VALUE                                              = 0x0501
	INT                                                        = 0x1404
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	RG32I                                                      = 0x823B
	OP_EXP_BASE_2_EXT                                          = 0x8791
	DEPTH_BITS                                                 = 0x0D56
	MAP2_VERTEX_3                                              = 0x0DB7
	TEXTURE_WRAP_R_OES                                         = 0x8072
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	VERTICAL_LINE_TO_NV                                        = 0x08
	LOAD                                                       = 0x0101
	FRONT_RIGHT                                                = 0x0401
	BLEND_SRC                                                  = 0x0BE1
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	CON_30_ATI                                                 = 0x895F
	LOW_FLOAT                                                  = 0x8DF0
	MITER_TRUNCATE_NV                                          = 0x90A8
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	WRITE_ONLY_ARB                                             = 0x88B9
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	PATH_MITER_LIMIT_NV                                        = 0x907A
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	SPARE0_NV                                                  = 0x852E
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	CONDITION_SATISFIED                                        = 0x911C
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	REG_16_ATI                                                 = 0x8931
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	COLOR_ATTACHMENT11                                         = 0x8CEB
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	STENCIL_TEST                                               = 0x0B90
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	SIGNED_INTENSITY_NV                                        = 0x8707
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	RELATIVE_ARC_TO_NV                                         = 0xFF
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	RGB8                                                       = 0x8051
	TEXTURE_MIN_LOD                                            = 0x813A
	FLOAT_MAT3x4                                               = 0x8B68
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	R1UI_V3F_SUN                                               = 0x85C4
	BUFFER_ACCESS                                              = 0x88BB
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	INDEX_CLEAR_VALUE                                          = 0x0C20
	CONSTANT_ATTENUATION                                       = 0x1207
	INTENSITY16_EXT                                            = 0x804D
	RGB_FLOAT32_APPLE                                          = 0x8815
	FLOAT_RGBA16_NV                                            = 0x888A
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	STATIC_DRAW_ARB                                            = 0x88E4
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	EQUAL                                                      = 0x0202
	LINE_SMOOTH_HINT                                           = 0x0C52
	FIXED_OES                                                  = 0x140C
	MIN_EXT                                                    = 0x8007
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	CULL_MODES_NV                                              = 0x86E0
	SUB_ATI                                                    = 0x8965
	COLOR_ATTACHMENT7                                          = 0x8CE7
	OP_CLAMP_EXT                                               = 0x878E
	STENCIL_INDEX4                                             = 0x8D47
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	RGB565                                                     = 0x8D62
	IMAGE_BUFFER_EXT                                           = 0x9051
	EMBOSS_MAP_NV                                              = 0x855F
	COPY_READ_BUFFER                                           = 0x8F36
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	STENCIL_BACK_FAIL                                          = 0x8801
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	SGIX_polynomial_ffd                                        = 1
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	COLOR4_BIT_PGI                                             = 0x00020000
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	ATTACHED_SHADERS                                           = 0x8B85
	READ_FRAMEBUFFER                                           = 0x8CA8
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	DS_BIAS_NV                                                 = 0x8716
	OP_NEGATE_EXT                                              = 0x8783
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	DOT3_RGB_EXT                                               = 0x8740
	FRAMEBUFFER_EXT                                            = 0x8D40
	EXT_copy_texture                                           = 1
	BGR_EXT                                                    = 0x80E0
	RGB8I_EXT                                                  = 0x8D8F
	LARGE_CW_ARC_TO_NV                                         = 0x18
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	SGIS_texture_lod                                           = 1
	INTERPOLATE_ARB                                            = 0x8575
	DRAW_BUFFER11_ATI                                          = 0x8830
	DYNAMIC_COPY_ARB                                           = 0x88EA
	FLOAT_MAT3                                                 = 0x8B5B
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	ISOLINES                                                   = 0x8E7A
	INDEX_LOGIC_OP                                             = 0x0BF1
	SAMPLER_CUBE_ARB                                           = 0x8B60
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	LIGHT3                                                     = 0x4003
	INTENSITY8I_EXT                                            = 0x8D91
	INT_SAMPLER_CUBE                                           = 0x8DCC
	TRANSFORM_FEEDBACK                                         = 0x8E22
	LOCATION_INDEX                                             = 0x930F
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	TEXTURE                                                    = 0x1702
	Y_EXT                                                      = 0x87D6
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	RGB4_S3TC                                                  = 0x83A1
	OP_FRAC_EXT                                                = 0x8789
	DYNAMIC_COPY                                               = 0x88EA
	TEXTURE24                                                  = 0x84D8
	DRAW_BUFFER0_ATI                                           = 0x8825
	CON_14_ATI                                                 = 0x894F
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	SCREEN_COORDINATES_REND                                    = 0x8490
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	REG_23_ATI                                                 = 0x8938
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	SGIX_ycrcb                                                 = 1
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	E_TIMES_F_NV                                               = 0x8531
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	ALPHA_SNORM                                                = 0x9010
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	SHADER_OBJECT_EXT                                          = 0x8B48
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	TEXTURE0                                                   = 0x84C0
	DOT2_ADD_ATI                                               = 0x896C
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	INT_SAMPLER_2D                                             = 0x8DCA
	NO_ERROR                                                   = 0
	MAP_COLOR                                                  = 0x0D10
	NORMAL_MAP_ARB                                             = 0x8511
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	REFERENCE_PLANE_SGIX                                       = 0x817D
	TEXTURE23                                                  = 0x84D7
	VERTEX_SHADER_EXT                                          = 0x8780
	SAMPLER_1D_ARB                                             = 0x8B5D
	SKIP_COMPONENTS3_NV                                        = -4
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	QUERY_COUNTER_BITS                                         = 0x8864
	ATC_RGB_AMD                                                = 0x8C92
	RED_INTEGER_EXT                                            = 0x8D94
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	RGB_SCALE_ARB                                              = 0x8573
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	TEXTURE_1D_ARRAY                                           = 0x8C18
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	PROXY_COLOR_TABLE                                          = 0x80D3
	OP_MUL_EXT                                                 = 0x8786
	TESS_CONTROL_SHADER                                        = 0x8E88
	BOUNDING_BOX_NV                                            = 0x908D
	SGIS_pixel_texture                                         = 1
	FUNC_SUBTRACT_OES                                          = 0x800A
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	SLUMINANCE_EXT                                             = 0x8C46
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	HISTOGRAM_FORMAT                                           = 0x8027
	MINMAX_SINK                                                = 0x8030
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	VARIABLE_F_NV                                              = 0x8528
	FULL_RANGE_EXT                                             = 0x87E1
	MATRIX1_ARB                                                = 0x88C1
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	TEXTURE_3D_EXT                                             = 0x806F
	OP_MOV_EXT                                                 = 0x8799
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	CURRENT_QUERY                                              = 0x8865
	REG_20_ATI                                                 = 0x8935
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	INT_VEC2                                                   = 0x8B53
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	MODELVIEW23_ARB                                            = 0x8737
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	MATRIX31_ARB                                               = 0x88DF
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	GREEN_INTEGER_EXT                                          = 0x8D95
	ALPHA_TEST_FUNC                                            = 0x0BC1
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	ALPHA8UI_EXT                                               = 0x8D7E
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	EXP2                                                       = 0x0801
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	MATRIX_PALETTE_ARB                                         = 0x8840
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	DRAW_BUFFER12_NV                                           = 0x8831
	CON_16_ATI                                                 = 0x8951
	STENCIL_INDEX1                                             = 0x8D46
	SGIX_blend_alpha_minmax                                    = 1
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	FOG_COORD                                                  = 0x8451
	TEXTURE6_ARB                                               = 0x84C6
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	SCALE_BY_TWO_NV                                            = 0x853E
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	CON_23_ATI                                                 = 0x8958
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	SRGB8_ALPHA8                                               = 0x8C43
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	RGBA16UI_EXT                                               = 0x8D76
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	BLEND_DST_RGB                                              = 0x80C8
	SECONDARY_COLOR_NV                                         = 0x852D
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	GL_4PASS_1_EXT                                             = 0x80A5
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	INT8_NV                                                    = 0x8FE0
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	RED_MIN_CLAMP_INGR                                         = 0x8560
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	UNSIGNALED                                                 = 0x9118
	READ_BUFFER_NV                                             = 0x0C02
	UNPACK_SKIP_IMAGES                                         = 0x806D
	R16I                                                       = 0x8233
	SRGB_DECODE_ARB                                            = 0x8299
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	VERTEX_STREAM7_ATI                                         = 0x8773
	RETURN                                                     = 0x0102
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	DOUBLE_VEC4                                                = 0x8FFE
	IMAGE_3D                                                   = 0x904E
	MAP2_COLOR_4                                               = 0x0DB0
	TEXTURE_ENV_MODE                                           = 0x2200
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	MATRIX25_ARB                                               = 0x88D9
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	UNSIGNED_INT16_NV                                          = 0x8FF0
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	SGIX_calligraphic_fragment                                 = 1
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	DECODE_EXT                                                 = 0x8A49
	SAMPLER_BUFFER                                             = 0x8DC2
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	TEXTURE29_ARB                                              = 0x84DD
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	VERSION_1_1                                                = 1
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	OPERAND0_ALPHA                                             = 0x8598
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	COMBINER_INPUT_NV                                          = 0x8542
	DRAW_BUFFER6                                               = 0x882B
	FLOAT_VEC3_ARB                                             = 0x8B51
	VERTEX_ID_NV                                               = 0x8C7B
	INTENSITY_SNORM                                            = 0x9013
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	RGB_SCALE                                                  = 0x8573
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	MODELVIEW16_ARB                                            = 0x8730
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	NEVER                                                      = 0x0200
	SAMPLES_ARB                                                = 0x80A9
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	ADD_SIGNED_EXT                                             = 0x8574
	STREAM_READ_ARB                                            = 0x88E1
	INT_VEC4                                                   = 0x8B55
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	MINMAX_FORMAT                                              = 0x802F
	PROGRAM_SEPARABLE                                          = 0x8258
	SOURCE3_RGB_NV                                             = 0x8583
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	REG_2_ATI                                                  = 0x8923
	SWIZZLE_STR_ATI                                            = 0x8976
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	AVERAGE_EXT                                                = 0x8335
	BUFFER_USAGE_ARB                                           = 0x8765
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	STENCIL_INDEX4_OES                                         = 0x8D47
	TANGENT_ARRAY_EXT                                          = 0x8439
	HI_SCALE_NV                                                = 0x870E
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	SGIX_async_histogram                                       = 1
	TEXTURE_ENV_COLOR                                          = 0x2201
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	FOG_COORDINATE_EXT                                         = 0x8451
	DRAW_BUFFER4_ATI                                           = 0x8829
	POINT                                                      = 0x1B00
	COLOR_TABLE_SCALE                                          = 0x80D6
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	DEPTH_SCALE                                                = 0x0D1E
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	AFFINE_3D_NV                                               = 0x9094
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	MAX                                                        = 0x8008
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	MAGNITUDE_BIAS_NV                                          = 0x8718
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	TEXTURE_SAMPLES                                            = 0x9106
	RGB10_EXT                                                  = 0x8052
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	PROGRAM                                                    = 0x82E2
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	INT_SAMPLER_1D                                             = 0x8DC9
	AUX0                                                       = 0x0409
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	VIEW_CLASS_8_BITS                                          = 0x82CB
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	MATRIX29_ARB                                               = 0x88DD
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	DRAW_PIXEL_TOKEN                                           = 0x0705
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	POINTS                                                     = 0x0000
	DECR                                                       = 0x1E03
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	TRIANGLE_STRIP                                             = 0x0005
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	QUERY_RESULT                                               = 0x8866
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	PATH_GEN_MODE_NV                                           = 0x90B0
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	INDEX_TEST_REF_EXT                                         = 0x81B7
	LO_BIAS_NV                                                 = 0x8715
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	COLOR_INDEXES                                              = 0x1603
	BLEND_EQUATION_OES                                         = 0x8009
	COMBINER2_NV                                               = 0x8552
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	TEXTURE2_ARB                                               = 0x84C2
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	CLIP_PLANE1                                                = 0x3001
	COMBINER0_NV                                               = 0x8550
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	DRAW_BUFFER11_ARB                                          = 0x8830
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	RENDERBUFFER_WIDTH                                         = 0x8D42
	MAP1_INDEX                                                 = 0x0D91
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	DEPTH_STENCIL                                              = 0x84F9
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	DRAW_BUFFER5_NV                                            = 0x882A
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	LUMINANCE12                                                = 0x8041
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	PROGRAM_POINT_SIZE                                         = 0x8642
	DRAW_BUFFER1_NV                                            = 0x8826
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	LIST_BASE                                                  = 0x0B32
	SPOT_DIRECTION                                             = 0x1204
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	COMBINE_RGB_EXT                                            = 0x8571
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	ACTIVE_TEXTURE                                             = 0x84E0
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	MAX_VARYING_FLOATS                                         = 0x8B4B
	RGB8I                                                      = 0x8D8F
	RG8_SNORM                                                  = 0x8F95
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	HISTOGRAM                                                  = 0x8024
	RG8                                                        = 0x822B
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	MEDIUM_INT                                                 = 0x8DF4
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	VERSION_2_0                                                = 1
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	MAX_LABEL_LENGTH                                           = 0x82E8
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	STENCIL_BACK_FUNC                                          = 0x8800
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	REG_13_ATI                                                 = 0x892E
	RGB32UI_EXT                                                = 0x8D71
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	MODELVIEW4_ARB                                             = 0x8724
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	MATRIX21_ARB                                               = 0x88D5
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	STENCIL_EXT                                                = 0x1802
	SOURCE2_RGB_ARB                                            = 0x8582
	SAMPLER_2D                                                 = 0x8B5E
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	COUNT_UP_NV                                                = 0x9088
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	TEXTURE4_ARB                                               = 0x84C4
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	INT16_VEC2_NV                                              = 0x8FE5
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	EXT_cmyka                                                  = 1
	HALF_APPLE                                                 = 0x140B
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	OPERAND1_ALPHA                                             = 0x8599
	DEPTH32F_STENCIL8                                          = 0x8CAD
	INTENSITY8                                                 = 0x804B
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	DUDV_ATI                                                   = 0x8779
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	SRGB                                                       = 0x8C40
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	SGIX_fragment_lighting                                     = 1
	LUMINANCE8                                                 = 0x8040
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	SIGNED_RGBA8_NV                                            = 0x86FC
	MATRIX15_ARB                                               = 0x88CF
	REPEAT                                                     = 0x2901
	PROXY_TEXTURE_3D                                           = 0x8070
	VIEW_CLASS_16_BITS                                         = 0x82CA
	COMBINE_ALPHA_EXT                                          = 0x8572
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	MODELVIEW12_ARB                                            = 0x872C
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	RGB_INTEGER                                                = 0x8D98
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	DEBUG_TYPE_ERROR                                           = 0x824C
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	AND_REVERSE                                                = 0x1502
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	CON_19_ATI                                                 = 0x8954
	SAMPLE_SHADING_ARB                                         = 0x8C36
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	TEXTURE12_ARB                                              = 0x84CC
	COMBINE_ARB                                                = 0x8570
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	CLAMP_TO_BORDER                                            = 0x812D
	CULL_VERTEX_EXT                                            = 0x81AA
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	COLOR_ARRAY_POINTER                                        = 0x8090
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	INCR_WRAP                                                  = 0x8507
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	MOVE_TO_RESETS_NV                                          = 0x90B5
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	R3_G3_B2                                                   = 0x2A10
	RG                                                         = 0x8227
	MAP1_BINORMAL_EXT                                          = 0x8446
	OPERAND2_ALPHA_ARB                                         = 0x859A
	RESTART_PATH_NV                                            = 0xF0
	RED_BIAS                                                   = 0x0D15
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	SRGB8                                                      = 0x8C41
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	ALPHA16_SNORM                                              = 0x9018
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	RGB16                                                      = 0x8054
	C4F_N3F_V3F                                                = 0x2A26
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	CON_22_ATI                                                 = 0x8957
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	ALPHA_TEST                                                 = 0x0BC0
	TEXTURE_RED_SIZE                                           = 0x805C
	RGBA12                                                     = 0x805A
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	ALPHA_FLOAT32_ATI                                          = 0x8816
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	INDEX_ARRAY_STRIDE                                         = 0x8086
	GL_422_EXT                                                 = 0x80CC
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	PREVIOUS_ARB                                               = 0x8578
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	DRAW_BUFFER5                                               = 0x882A
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	LINE_STRIP_ADJACENCY                                       = 0x000B
	NAME_STACK_DEPTH                                           = 0x0D70
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	DECR_WRAP                                                  = 0x8508
	SGIX_texture_multi_buffer                                  = 1
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	REPLACE_VALUE_AMD                                          = 0x874B
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	FENCE_CONDITION_NV                                         = 0x84F4
	SAMPLES_PASSED                                             = 0x8914
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	MODELVIEW22_ARB                                            = 0x8736
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	DRAW_BUFFER4_NV                                            = 0x8829
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	DST_ALPHA                                                  = 0x0304
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	MODELVIEW31_ARB                                            = 0x873F
	SIGNED_HILO16_NV                                           = 0x86FA
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	STENCIL_INDEX1_OES                                         = 0x8D46
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	RGB8UI                                                     = 0x8D7D
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	FILTER4_SGIS                                               = 0x8146
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	TEXTURE27                                                  = 0x84DB
	INCR_WRAP_EXT                                              = 0x8507
	TEXTURE_MATRIX                                             = 0x0BA8
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	VARIANT_EXT                                                = 0x87C1
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	ZOOM_Y                                                     = 0x0D17
	REG_30_ATI                                                 = 0x893F
	OPERAND2_ALPHA                                             = 0x859A
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	VERSION_1_3                                                = 1
	CURRENT_RASTER_POSITION                                    = 0x0B07
	INTERLACE_SGIX                                             = 0x8094
	MATRIX7_ARB                                                = 0x88C7
	DOUBLE_MAT4x2                                              = 0x8F4D
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	CLIP_DISTANCE1                                             = 0x3001
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	VARIABLE_G_NV                                              = 0x8529
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	STATIC_COPY_ARB                                            = 0x88E6
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	VIDEO_BUFFER_NV                                            = 0x9020
	SUBPIXEL_BITS                                              = 0x0D50
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	SWIZZLE_STRQ_ATI                                           = 0x897A
	IMAGE_1D_EXT                                               = 0x904C
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	DEBUG_TYPE_OTHER                                           = 0x8251
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	CULL_FACE_MODE                                             = 0x0B45
	ALPHA_TEST_QCOM                                            = 0x0BC0
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	DRAW_BUFFER3_ARB                                           = 0x8828
	ADD_ATI                                                    = 0x8963
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	RGBA12_EXT                                                 = 0x805A
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	FRACTIONAL_EVEN                                            = 0x8E7C
	VERSION_1_5                                                = 1
	CCW                                                        = 0x0901
	FEEDBACK                                                   = 0x1C01
	OPERAND0_RGB_EXT                                           = 0x8590
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	IMAGE_1D_ARRAY                                             = 0x9052
	COLOR_WRITEMASK                                            = 0x0C23
	TEXTURE_1D                                                 = 0x0DE0
	MODELVIEW26_ARB                                            = 0x873A
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	CLOSE_PATH_NV                                              = 0x00
	MULTISAMPLE                                                = 0x809D
	IUI_V2F_EXT                                                = 0x81AD
	VARIANT_ARRAY_EXT                                          = 0x87E8
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	RGB_SNORM                                                  = 0x8F92
	UNIFORM_BLOCK                                              = 0x92E2
	EXT_abgr                                                   = 1
	INTERPOLATE_EXT                                            = 0x8575
	FACTOR_MIN_AMD                                             = 0x901C
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	INTENSITY12                                                = 0x804C
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	REGISTER_COMBINERS_NV                                      = 0x8522
	INVARIANT_VALUE_EXT                                        = 0x87EA
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	GL_2PASS_0_SGIS                                            = 0x80A2
	R8_EXT                                                     = 0x8229
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	MAX_SHININESS_NV                                           = 0x8504
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	STENCIL_VALUE_MASK                                         = 0x0B93
	ACCUM_BLUE_BITS                                            = 0x0D5A
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	LUMINANCE16UI_EXT                                          = 0x8D7A
	MINOR_VERSION                                              = 0x821C
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	DOT3_RGB_ARB                                               = 0x86AE
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	BLUE_BIAS                                                  = 0x0D1B
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	RGB16F_EXT                                                 = 0x881B
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	FRONT_AND_BACK                                             = 0x0408
	MAP1_COLOR_4                                               = 0x0D90
	GL_1PASS_EXT                                               = 0x80A1
	TEXTURE_BASE_LEVEL                                         = 0x813C
	TEXTURE_RECTANGLE                                          = 0x84F5
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	LUMINANCE32F_ARB                                           = 0x8818
	FLOAT_R32_NV                                               = 0x8885
	FLOAT_RGBA32_NV                                            = 0x888B
	FLOAT_MAT2x4                                               = 0x8B66
	CURRENT_PROGRAM                                            = 0x8B8D
	EXT_subtexture                                             = 1
	LAYOUT_LINEAR_INTEL                                        = 1
	INDEX_MATERIAL_EXT                                         = 0x81B8
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	CLAMP_TO_EDGE                                              = 0x812F
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	STENCIL_FUNC                                               = 0x0B92
	COLOR_INDEX                                                = 0x1900
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	MAP_TESSELLATION_NV                                        = 0x86C2
	LOCAL_EXT                                                  = 0x87C4
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	RGB_INTEGER_EXT                                            = 0x8D98
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	FOG_COORDINATE_SOURCE                                      = 0x8450
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	MATRIX14_ARB                                               = 0x88CE
	VIEW_CLASS_128_BITS                                        = 0x82C4
	TEXTURE_WRAP_S                                             = 0x2802
	ALPHA_MAX_SGIX                                             = 0x8321
	INVARIANT_EXT                                              = 0x87C2
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	AUX_BUFFERS                                                = 0x0C00
	CLIP_PLANE2                                                = 0x3002
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	MODELVIEW24_ARB                                            = 0x8738
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	TEXTURE_BORDER_COLOR                                       = 0x1004
	MIRRORED_REPEAT_ARB                                        = 0x8370
	PRIMARY_COLOR_EXT                                          = 0x8577
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	BITMAP_TOKEN                                               = 0x0704
	COMPILE                                                    = 0x1300
	SRC2_ALPHA                                                 = 0x858A
	PROGRAM_FORMAT_ARB                                         = 0x8876
	TEXTURE_2D                                                 = 0x0DE1
	T2F_C4UB_V3F                                               = 0x2A29
	T2F_C3F_V3F                                                = 0x2A2A
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	NORMAL_MAP_OES                                             = 0x8511
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	DRAW_BUFFER15                                              = 0x8834
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	INDEX_ARRAY_POINTER                                        = 0x8091
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	OBJECT_TYPE                                                = 0x9112
	CW                                                         = 0x0900
	ALPHA_SCALE                                                = 0x0D1C
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	DUAL_ALPHA8_SGIS                                           = 0x8111
	ALPHA_MIN_SGIX                                             = 0x8320
	DOT3_RGBA_IMG                                              = 0x86AF
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	COMPRESSED_RGBA                                            = 0x84EE
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	MOVE_TO_NV                                                 = 0x02
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	REG_4_ATI                                                  = 0x8925
	REG_22_ATI                                                 = 0x8937
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	DRAW_BUFFER                                                = 0x0C01
	CONVOLUTION_FORMAT                                         = 0x8017
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	TRACE_NAME_MESA                                            = 0x8756
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	ACTIVE_UNIFORMS                                            = 0x8B86
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	PATH_FILL_MODE_NV                                          = 0x9080
	PATH_STROKE_MASK_NV                                        = 0x9084
	RENDERER                                                   = 0x1F01
	NUM_EXTENSIONS                                             = 0x821D
	DEBUG_PRINT_MESA                                           = 0x875A
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	ADD_BLEND_IMG                                              = 0x8C09
	LUMINANCE8I_EXT                                            = 0x8D92
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	POLYGON                                                    = 0x0009
	INTENSITY12_EXT                                            = 0x804C
	FLOAT_MAT4x3                                               = 0x8B6A
	TEXTURE_RED_TYPE                                           = 0x8C10
	COLOR_BUFFER_BIT                                           = 0x00004000
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	BUFFER                                                     = 0x82E0
	MODELVIEW1_EXT                                             = 0x850A
	COMBINER1_NV                                               = 0x8551
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	R32I                                                       = 0x8235
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	RGBA_FLOAT32_APPLE                                         = 0x8814
	UNIFORM                                                    = 0x92E1
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	MATRIX24_ARB                                               = 0x88D8
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	COMPRESSED_LUMINANCE                                       = 0x84EA
	DRAW_BUFFER14_ATI                                          = 0x8833
	SCISSOR_TEST                                               = 0x0C11
	N3F_V3F                                                    = 0x2A25
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	BACK_LEFT                                                  = 0x0402
	DEPTH                                                      = 0x1801
	RGBA16                                                     = 0x805B
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	VOLATILE_APPLE                                             = 0x8A1A
	RGBA8_SNORM                                                = 0x8F97
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	SUBTRACT_ARB                                               = 0x84E7
	SLICE_ACCUM_SUN                                            = 0x85CC
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	TESS_GEN_POINT_MODE                                        = 0x8E79
	IMAGE_1D                                                   = 0x904C
	SGIX_icc_texture                                           = 1
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	IMAGE_CUBE                                                 = 0x9050
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	TEXTURE_4D_SGIS                                            = 0x8134
	RG8UI                                                      = 0x8238
	MATRIX2_ARB                                                = 0x88C2
	STENCIL_INDEX8                                             = 0x8D48
	ADJACENT_PAIRS_NV                                          = 0x90AE
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	SAMPLER                                                    = 0x82E6
	GL_3DC_X_AMD                                               = 0x87F9
	ALPHA8I_EXT                                                = 0x8D90
	FLOAT_RG16_NV                                              = 0x8886
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	CURRENT_RASTER_COLOR                                       = 0x0B04
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	FRAGMENT_SHADER                                            = 0x8B30
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	RELATIVE_LINE_TO_NV                                        = 0x05
	EXT_texture_object                                         = 1
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	GL_4PASS_1_SGIS                                            = 0x80A5
	SOURCE1_RGB                                                = 0x8581
	BUFFER_MAPPED                                              = 0x88BC
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	EDGE_FLAG_ARRAY                                            = 0x8079
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	DEPTH_COMPONENT16                                          = 0x81A5
	RG_INTEGER                                                 = 0x8228
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	SHADER_BINARY_DMP                                          = 0x9250
	SRC_ALPHA                                                  = 0x0302
	STENCIL_BITS                                               = 0x0D57
	RGB10_A2_EXT                                               = 0x8059
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	BUFFER_ACCESS_ARB                                          = 0x88BB
	FLOAT_MAT3_ARB                                             = 0x8B5B
	FIELD_UPPER_NV                                             = 0x9022
	MAP1_VERTEX_4                                              = 0x0D98
	MAX_EXT                                                    = 0x8008
	MAX_LAYERS                                                 = 0x8281
	SHADER_IMAGE_LOAD                                          = 0x82A4
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	MATRIX23_ARB                                               = 0x88D7
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	MAP_WRITE_BIT                                              = 0x0002
	STENCIL_FAIL                                               = 0x0B94
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	SPRITE_SGIX                                                = 0x8148
	TEXTURE18_ARB                                              = 0x84D2
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	DT_BIAS_NV                                                 = 0x8717
	DRAW_BUFFER8                                               = 0x882D
	STATIC_DRAW                                                = 0x88E4
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	OR_INVERTED                                                = 0x150D
	ALPHA12                                                    = 0x803D
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	SRC0_ALPHA                                                 = 0x8588
	DRAW_BUFFER9_ARB                                           = 0x882E
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	SGIS_texture_border_clamp                                  = 1
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	MATRIX8_ARB                                                = 0x88C8
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	DOUBLE_MAT3x4                                              = 0x8F4C
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	XOR                                                        = 0x1506
	LIGHT6                                                     = 0x4006
	VIEW_CLASS_32_BITS                                         = 0x82C8
	COMBINER6_NV                                               = 0x8556
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	MATRIX12_ARB                                               = 0x88CC
	GL_4X_BIT_ATI                                              = 0x00000002
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	CONST_EYE_NV                                               = 0x86E5
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	MULTISAMPLE_SGIS                                           = 0x809D
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	DRAW_BUFFER1_ATI                                           = 0x8826
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	RENDERBUFFER_EXT                                           = 0x8D41
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	MATRIX22_ARB                                               = 0x88D6
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	DRAW_BUFFER6_NV                                            = 0x882B
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	TEXTURE16_ARB                                              = 0x84D0
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	RGB5_A1                                                    = 0x8057
	DSDT_NV                                                    = 0x86F5
	ETC1_SRGB8_NV                                              = 0x88EE
	SAMPLER_3D                                                 = 0x8B5F
	IUI_V3F_EXT                                                = 0x81AE
	WRITE_DISCARD_NV                                           = 0x88BE
	RGB5_A1_EXT                                                = 0x8057
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	RGBA16F_ARB                                                = 0x881A
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	CND0_ATI                                                   = 0x896B
	CONVOLUTION_2D                                             = 0x8011
	REFLECTION_MAP_OES                                         = 0x8512
	PROGRAM_TARGET_NV                                          = 0x8646
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	LUMINANCE32UI_EXT                                          = 0x8D74
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	TEXTURE16                                                  = 0x84D0
	INTERPOLATE                                                = 0x8575
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	VERTEX_PROGRAM_NV                                          = 0x8620
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	SEPARABLE_2D_EXT                                           = 0x8012
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	INDEX_BIT_PGI                                              = 0x00080000
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	COUNTER_RANGE_AMD                                          = 0x8BC1
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	DEPTH_EXT                                                  = 0x1801
	MODELVIEW3_ARB                                             = 0x8723
	MAX_INTEGER_SAMPLES                                        = 0x9110
	STENCIL_WRITEMASK                                          = 0x0B98
	PROXY_TEXTURE_1D                                           = 0x8063
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	INVERT                                                     = 0x150A
	TEXTURE11_ARB                                              = 0x84CB
	GL_8X_BIT_ATI                                              = 0x00000004
	SAMPLE_MASK_VALUE                                          = 0x8E52
	TRANSLATE_2D_NV                                            = 0x9090
	RED                                                        = 0x1903
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	MATRIX3_NV                                                 = 0x8633
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	STORAGE_CACHED_APPLE                                       = 0x85BE
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	BLEND_DST                                                  = 0x0BE0
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	SAMPLER_BUFFER_AMD                                         = 0x9001
	PATH_COORD_COUNT_NV                                        = 0x909E
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	IUI_N3F_V2F_EXT                                            = 0x81AF
	OPERAND1_RGB_EXT                                           = 0x8591
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	QUERY_WAIT_NV                                              = 0x8E13
	FOG_COORDINATE                                             = 0x8451
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	BLEND_DST_RGB_EXT                                          = 0x80C8
	BGRA_INTEGER                                               = 0x8D9B
	MAX_EVAL_ORDER                                             = 0x0D30
	MODELVIEW2_ARB                                             = 0x8722
	PALETTE4_RGB8_OES                                          = 0x8B90
	RENDERBUFFER_OES                                           = 0x8D41
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	TEXTURE_RESIDENT                                           = 0x8067
	RGBA2_EXT                                                  = 0x8055
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	COMBINER_MAPPING_NV                                        = 0x8543
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	OP_DOT3_EXT                                                = 0x8784
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	READ_PIXELS_FORMAT                                         = 0x828D
	TEXTURE_GATHER                                             = 0x82A2
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	ACTIVE_RESOURCES                                           = 0x92F5
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	TEXTURE1_ARB                                               = 0x84C1
	DRAW_BUFFER4                                               = 0x8829
	MATRIX17_ARB                                               = 0x88D1
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	RGBA16UI                                                   = 0x8D76
	COMPUTE_SHADER_BIT                                         = 0x00000020
	SCISSOR_BOX                                                = 0x0C10
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	POLYGON_BIT                                                = 0x00000008
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	MINMAX_EXT                                                 = 0x802E
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	ALWAYS                                                     = 0x0207
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	DRAW_BUFFER6_ATI                                           = 0x882B
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	GREEN_BIAS                                                 = 0x0D19
	LIGHT4                                                     = 0x4004
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	NORMAL_MAP_NV                                              = 0x8511
	FLOAT_VEC4_ARB                                             = 0x8B52
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	UTF8_NV                                                    = 0x909A
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	CONTEXT_FLAGS                                              = 0x821E
	PROGRAM_RESIDENT_NV                                        = 0x8647
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	CURRENT_NORMAL                                             = 0x0B02
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	OPERAND1_RGB_ARB                                           = 0x8591
	MODELVIEW15_ARB                                            = 0x872F
	RGBA_FLOAT32_ATI                                           = 0x8814
	DRAW_BUFFER15_ATI                                          = 0x8834
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	INDEX                                                      = 0x8222
	RG16UI                                                     = 0x823A
	DRAW_BUFFER14                                              = 0x8833
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	FLOAT_MAT2_ARB                                             = 0x8B5A
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	RGB_FLOAT16_APPLE                                          = 0x881B
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	OUTPUT_VERTEX_EXT                                          = 0x879A
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	PIXEL_MODE_BIT                                             = 0x00000020
	TEXTURE5                                                   = 0x84C5
	CONSTANT_COLOR1_NV                                         = 0x852B
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	POSITION                                                   = 0x1203
	CLIP_PLANE3                                                = 0x3003
	TIME_ELAPSED                                               = 0x88BF
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	TEXTURE7                                                   = 0x84C7
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	MOV_ATI                                                    = 0x8961
	VERTEX_SHADER_ARB                                          = 0x8B31
	RGB16I                                                     = 0x8D89
	SGIS_fog_function                                          = 1
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	TEXTURE1                                                   = 0x84C1
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	DOUBLE_MAT3                                                = 0x8F47
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	TEXTURE_PRIORITY                                           = 0x8066
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	IS_ROW_MAJOR                                               = 0x9300
	VERTEX_ARRAY                                               = 0x8074
	TEXTURE_WRAP_R                                             = 0x8072
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	SAMPLER_2D_RECT                                            = 0x8B63
	OBJECT_TYPE_ARB                                            = 0x8B4E
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	SIGNED_ALPHA_NV                                            = 0x8705
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	BUFFER_OBJECT_EXT                                          = 0x9151
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	READ_PIXELS_TYPE                                           = 0x828E
	SIGNED_RGB8_NV                                             = 0x86FF
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	FLOAT_VEC4                                                 = 0x8B52
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	ZOOM_X                                                     = 0x0D16
	ASYNC_MARKER_SGIX                                          = 0x8329
	SOURCE1_ALPHA                                              = 0x8589
	PROGRAM_STRING_ARB                                         = 0x8628
	FLOAT_RGB32_NV                                             = 0x8889
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	PROGRAM_LENGTH_NV                                          = 0x8627
	MODELVIEW1_ARB                                             = 0x850A
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	POINT_TOKEN                                                = 0x0701
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	COMBINE_ALPHA                                              = 0x8572
	RGBA32F_ARB                                                = 0x8814
	MUL_ATI                                                    = 0x8964
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	COPY_PIXEL_TOKEN                                           = 0x0706
	VIEWPORT                                                   = 0x0BA2
	CURRENT_MATRIX_NV                                          = 0x8641
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	NORMAL_ARRAY_EXT                                           = 0x8075
	CON_31_ATI                                                 = 0x8960
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	COLOR_SUM_ARB                                              = 0x8458
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	INT_IMAGE_2D_RECT                                          = 0x905A
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	YCRCB_444_SGIX                                             = 0x81BC
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	DRAW_BUFFER1_ARB                                           = 0x8826
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	DU8DV8_ATI                                                 = 0x877A
	STATIC_COPY                                                = 0x88E6
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	EXTENSIONS                                                 = 0x1F03
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	BLUE                                                       = 0x1905
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	DEPTH_CLAMP_NV                                             = 0x864F
	VIBRANCE_SCALE_NV                                          = 0x8713
	FLOAT_MAT4x2                                               = 0x8B69
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	RESTART_SUN                                                = 0x0001
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	TEXTURE28_ARB                                              = 0x84DC
	QUAD_MESH_SUN                                              = 0x8614
	OP_MADD_EXT                                                = 0x8788
	REG_10_ATI                                                 = 0x892B
	CON_0_ATI                                                  = 0x8941
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	DRAW_BUFFER5_ATI                                           = 0x882A
	RG16_SNORM                                                 = 0x8F99
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	CLAMP                                                      = 0x2900
	ACTIVE_PROGRAM                                             = 0x8259
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	COMPRESSED_ALPHA                                           = 0x84E9
	REFLECTION_MAP_NV                                          = 0x8512
	DYNAMIC_ATI                                                = 0x8761
	DRAW_BUFFER15_ARB                                          = 0x8834
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	ADD_SIGNED                                                 = 0x8574
	FRAGMENT_SHADER_ATI                                        = 0x8920
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	ALPHA16UI_EXT                                              = 0x8D78
	HALF_FLOAT_ARB                                             = 0x140B
	SPRITE_AXIS_SGIX                                           = 0x814A
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	SRGB_ALPHA                                                 = 0x8C42
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	DYNAMIC_READ_ARB                                           = 0x88E9
	CON_10_ATI                                                 = 0x894B
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	SGIX_flush_raster                                          = 1
	DEPTH_FUNC                                                 = 0x0B74
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	GL_4PASS_2_SGIS                                            = 0x80A6
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	CURRENT_FOG_COORD                                          = 0x8453
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	RGB_FLOAT32_ATI                                            = 0x8815
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	BUFFER_BINDING                                             = 0x9302
	INVALID_OPERATION                                          = 0x0502
	AUTO_NORMAL                                                = 0x0D80
	GL_2_BYTES                                                 = 0x1407
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	MAJOR_VERSION                                              = 0x821B
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	INTENSITY16F_ARB                                           = 0x881D
	HILO8_NV                                                   = 0x885E
	WRITE_ONLY                                                 = 0x88B9
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	DELETE_STATUS                                              = 0x8B80
	FACTOR_MAX_AMD                                             = 0x901D
	RECT_NV                                                    = 0xF6
	CONSTANT_ALPHA                                             = 0x8003
	RGB12_EXT                                                  = 0x8053
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	INTERLACE_READ_INGR                                        = 0x8568
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	DRAW_BUFFER3_ATI                                           = 0x8828
	ARRAY_BUFFER                                               = 0x8892
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	OBJECT_PLANE                                               = 0x2501
	SOURCE0_RGB_EXT                                            = 0x8580
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	PIXEL_PACK_BUFFER                                          = 0x88EB
	COMMAND_BARRIER_BIT                                        = 0x00000040
	NORMAL_ARRAY                                               = 0x8075
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	CLIP_DISTANCE5                                             = 0x3005
	ADD_SIGNED_ARB                                             = 0x8574
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	IS_PER_PATCH                                               = 0x92E7
	NONE                                                       = 0
	LIGHT5                                                     = 0x4005
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	MATRIX7_NV                                                 = 0x8637
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	COLOR_ARRAY_SIZE                                           = 0x8081
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	TEXTURE9                                                   = 0x84C9
	TEXTURE30_ARB                                              = 0x84DE
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	ORDER                                                      = 0x0A01
	DEPTH_COMPONENT                                            = 0x1902
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	OPERAND2_RGB                                               = 0x8592
	CURRENT_MATRIX_ARB                                         = 0x8641
	OP_DOT4_EXT                                                = 0x8785
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	TEXTURE27_ARB                                              = 0x84DB
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	DEBUG_OUTPUT                                               = 0x92E0
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	COLOR_ENCODING                                             = 0x8296
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	BUMP_ENVMAP_ATI                                            = 0x877B
	DRAW_BUFFER2                                               = 0x8827
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	PATH_END_CAPS_NV                                           = 0x9076
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	PHONG_WIN                                                  = 0x80EA
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	NUM_PASSES_ATI                                             = 0x8970
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	SGIX_texture_add_env                                       = 1
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	KEEP                                                       = 0x1E00
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	INVERSE_NV                                                 = 0x862B
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	MATRIX4_ARB                                                = 0x88C4
	ALPHA16I_EXT                                               = 0x8D8A
	SHADE_MODEL                                                = 0x0B54
	INDEX_MODE                                                 = 0x0C30
	MAP2_VERTEX_4                                              = 0x0DB8
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	ARRAY_BUFFER_ARB                                           = 0x8892
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	FOG_FUNC_SGIS                                              = 0x812A
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	STENCIL_REF                                                = 0x0B97
	DEPTH_COMPONENT24                                          = 0x81A6
	FLOAT_MAT2                                                 = 0x8B5A
	RETAINED_APPLE                                             = 0x8A1B
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	HISTOGRAM_EXT                                              = 0x8024
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	LIST_BIT                                                   = 0x00020000
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	COLOR_INDEX1_EXT                                           = 0x80E2
	STENCIL_COMPONENTS                                         = 0x8285
	GEOMETRY_TEXTURE                                           = 0x829E
	DRAW_BUFFER7_ATI                                           = 0x882C
	BLUE_INTEGER_EXT                                           = 0x8D96
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	RENDER_MODE                                                = 0x0C40
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	QUERY_OBJECT_AMD                                           = 0x9153
	DEPTH_BUFFER_BIT                                           = 0x00000100
	FRONT_LEFT                                                 = 0x0400
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	CURRENT_ATTRIB_NV                                          = 0x8626
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	LINEAR_ATTENUATION                                         = 0x1208
	CLAMP_TO_BORDER_NV                                         = 0x812D
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	MAX_CLIP_DISTANCES                                         = 0x0D32
	IMAGE_SCALE_X_HP                                           = 0x8155
	TEXTURE14_ARB                                              = 0x84CE
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	STENCIL_INDEX8_EXT                                         = 0x8D48
	TEXTURE28                                                  = 0x84DC
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	VECTOR_EXT                                                 = 0x87BF
	MATRIX5_ARB                                                = 0x88C5
	SHADER_TYPE                                                = 0x8B4F
	LEFT                                                       = 0x0406
	SHORT                                                      = 0x1402
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	COLOR_INDEX2_EXT                                           = 0x80E3
	OP_SET_LT_EXT                                              = 0x878D
	RGB32F                                                     = 0x8815
	ALPHA_FLOAT16_ATI                                          = 0x881C
	DRAW_BUFFER6_ARB                                           = 0x882B
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	EXT_texture3D                                              = 1
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	RED_MAX_CLAMP_INGR                                         = 0x8564
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	INTENSITY32F_ARB                                           = 0x8817
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	DRAW_BUFFER11                                              = 0x8830
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	DOUBLE_MAT3_EXT                                            = 0x8F47
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	DSDT_MAG_NV                                                = 0x86F6
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	DOUBLE                                                     = 0x140A
	DECR_WRAP_OES                                              = 0x8508
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	REG_1_ATI                                                  = 0x8922
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	SGIS_detail_texture                                        = 1
	LINE_STIPPLE                                               = 0x0B24
	LIST_INDEX                                                 = 0x0B33
	R8                                                         = 0x8229
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	SGIX_vertex_preclip                                        = 1
	COLOR_EXT                                                  = 0x1800
	TEXTURE_WRAP_T                                             = 0x2803
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	BUFFER_USAGE                                               = 0x8765
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	LESS                                                       = 0x0201
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	BLEND_SRC_RGB_OES                                          = 0x80C9
	FOG_COORD_ARRAY                                            = 0x8457
	COMBINER5_NV                                               = 0x8555
	SURFACE_STATE_NV                                           = 0x86EB
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	CURRENT_VERTEX_EXT                                         = 0x87E2
	MATRIX6_NV                                                 = 0x8636
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	RGBA4_EXT                                                  = 0x8056
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	COORD_REPLACE_NV                                           = 0x8862
	MODULATE_COLOR_IMG                                         = 0x8C04
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	SLUMINANCE                                                 = 0x8C46
	PATH_STENCIL_REF_NV                                        = 0x90B8
	SAMPLER_OBJECT_AMD                                         = 0x9155
	COLOR3_BIT_PGI                                             = 0x00010000
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	SINGLE_COLOR                                               = 0x81F9
	SOURCE1_ALPHA_ARB                                          = 0x8589
	FLOAT_R_NV                                                 = 0x8880
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	COLOR_MATRIX                                               = 0x80B1
	VIEW_CLASS_64_BITS                                         = 0x82C6
	UNDEFINED_APPLE                                            = 0x8A1C
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	COPY                                                       = 0x1503
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	ENABLE_BIT                                                 = 0x00002000
	FLAT                                                       = 0x1D00
	CLIP_DISTANCE2                                             = 0x3002
	ABGR_EXT                                                   = 0x8000
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	OPERAND0_ALPHA_EXT                                         = 0x8598
	OPERAND1_ALPHA_EXT                                         = 0x8599
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	MODELVIEW29_ARB                                            = 0x873D
	SAMPLER_BINDING                                            = 0x8919
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	DONT_CARE                                                  = 0x1100
	FRAMEZOOM_SGIX                                             = 0x818B
	TEXTURE21                                                  = 0x84D5
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	SOURCE2_ALPHA                                              = 0x858A
	BGR_INTEGER                                                = 0x8D9A
	PATH_DASH_CAPS_NV                                          = 0x907B
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	BOOL_VEC4_ARB                                              = 0x8B59
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	RGB10_A2                                                   = 0x8059
	RG_EXT                                                     = 0x8227
	INT_10_10_10_2_OES                                         = 0x8DF7
	CONVOLUTION_2D_EXT                                         = 0x8011
	SRGB_READ                                                  = 0x8297
	OP_INDEX_EXT                                               = 0x8782
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	VERTEX_SUBROUTINE                                          = 0x92E8
	REPLACE                                                    = 0x1E01
	SOURCE1_ALPHA_EXT                                          = 0x8589
	SATURATE_BIT_ATI                                           = 0x00000040
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	PALETTE4_RGBA4_OES                                         = 0x8B93
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	COLOR_MATERIAL                                             = 0x0B57
	RG16I                                                      = 0x8239
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	RGBA16F_EXT                                                = 0x881A
	REG_29_ATI                                                 = 0x893E
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	YCRCB_422_SGIX                                             = 0x81BB
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	GEQUAL                                                     = 0x0206
	BACK                                                       = 0x0405
	CON_27_ATI                                                 = 0x895C
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	RGB4                                                       = 0x804F
	INDEX_TEST_EXT                                             = 0x81B5
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	S                                                          = 0x2000
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	FRAGMENT_COLOR_EXT                                         = 0x834C
	BINORMAL_ARRAY_EXT                                         = 0x843A
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	TRANSLATE_3D_NV                                            = 0x9091
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	SQUARE_NV                                                  = 0x90A3
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	POLYGON_OFFSET_POINT                                       = 0x2A01
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	SYNC_STATUS                                                = 0x9114
	SYNC_CONDITION_APPLE                                       = 0x9113
	INDEX_ARRAY_EXT                                            = 0x8077
	VIEW_CLASS_96_BITS                                         = 0x82C5
	COMBINER7_NV                                               = 0x8557
	SHADER_CONSISTENT_NV                                       = 0x86DD
	SIGNED_RGBA_NV                                             = 0x86FB
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	RGB16I_EXT                                                 = 0x8D89
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	SCISSOR_BIT                                                = 0x00080000
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	YCBYCR8_422_NV                                             = 0x9031
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	LINE_TO_NV                                                 = 0x04
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	T2F_IUI_V2F_EXT                                            = 0x81B1
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	RGBA16_SNORM                                               = 0x8F9B
	RG16                                                       = 0x822C
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	VARIABLE_E_NV                                              = 0x8527
	EVAL_2D_NV                                                 = 0x86C0
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	BUFFER_ACCESS_OES                                          = 0x88BB
	CON_20_ATI                                                 = 0x8955
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	FLOAT_VEC3                                                 = 0x8B51
	SIGNED_NORMALIZED                                          = 0x8F9C
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	OBJECT_LINEAR                                              = 0x2401
	QUAD_ALPHA4_SGIS                                           = 0x811E
	BUFFER_SIZE                                                = 0x8764
	MAX_TEXTURE_COORDS                                         = 0x8871
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	LUMINANCE_ALPHA                                            = 0x190A
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	TIMEOUT_EXPIRED                                            = 0x911B
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	HALF_FLOAT_NV                                              = 0x140B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	ALPHA32UI_EXT                                              = 0x8D72
	ALREADY_SIGNALED                                           = 0x911A
	TEXTURE_GEN_MODE                                           = 0x2500
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	INT_SAMPLER_3D                                             = 0x8DCB
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	MODELVIEW18_ARB                                            = 0x8732
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	SYNC_FENCE                                                 = 0x9116
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	RGBA_FLOAT16_APPLE                                         = 0x881A
	STREAM_DRAW                                                = 0x88E0
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	DECR_WRAP_EXT                                              = 0x8508
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	INT_VEC4_ARB                                               = 0x8B55
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	TESS_EVALUATION_SHADER                                     = 0x8E87
	CLIP_DISTANCE0                                             = 0x3000
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	UNIFORM_TYPE                                               = 0x8A37
	IMAGE_BUFFER                                               = 0x9051
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	RGB2_EXT                                                   = 0x804E
	UNSIGNED_INT_24_8                                          = 0x84FA
	POINT_SPRITE                                               = 0x8861
	COORD_REPLACE_ARB                                          = 0x8862
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	ALPHA8_SNORM                                               = 0x9014
	DEBUG_ASSERT_MESA                                          = 0x875B
	MATRIX6_ARB                                                = 0x88C6
	COUNTER_TYPE_AMD                                           = 0x8BC0
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	FIELD_LOWER_NV                                             = 0x9023
	NORMAL_MAP                                                 = 0x8511
	HALF_BIT_ATI                                               = 0x00000008
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	VERTEX_STREAM1_ATI                                         = 0x876D
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	FRAMEBUFFER                                                = 0x8D40
	TEXTURE_SAMPLES_IMG                                        = 0x9136
)

type Context struct {
	access                    sync.Mutex
	context                   *C.gl14Context
	extensions                map[string]bool
	inBeginEnd                bool
	traceback                 []string
	Accum                     func(op uint32, value float32)
	AlphaFunc                 func(Func uint32, ref float32)
	Begin                     func(mode uint32)
	End                       func()
	Bitmap                    func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8)
	BlendFunc                 func(sfactor, dfactor uint32)
	CallList                  func(list uint32)
	CallLists                 func(n int32, Type uint32, lists unsafe.Pointer)
	Clear                     func(mask uint32)
	ClearAccum                func(red, green, blue, alpha float32)
	ClearColor                func(red, green, blue, alpha float32)
	ClearDepth                func(depth float64)
	ClearIndex                func(c float32)
	ClearStencil              func(s int32)
	ClipPlane                 func(plane uint32, equation *float64)
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
	ColorMaterial             func(face, mode uint32)
	CopyPixels                func(x, y int32, width, height int32, Type uint32)
	CullFace                  func(mode uint32)
	DeleteLists               func(list uint32, Range int32)
	DepthFunc                 func(Func uint32)
	DepthMask                 func(flag bool)
	DepthRange                func(zNear, zFar float64)
	Enable                    func(cap uint32)
	Disable                   func(cap uint32)
	DrawBuffer                func(mode uint32)
	DrawPixels                func(width, height int32, format, Type uint32, data unsafe.Pointer)
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
	EvalMesh1                 func(mode uint32, i1, i2 int32)
	EvalMesh2                 func(mode uint32, i1, i2, j1, j2 int32)
	EvalPoint1                func(i int32)
	EvalPoint2                func(i, j int32)
	FeedbackBuffer            func(size int32, Type uint32, buffer *float32)
	Finish                    func()
	Flush                     func()
	Fogf                      func(pname uint32, param float32)
	Fogi                      func(pname uint32, param int32)
	Fogfv                     func(pname uint32, params *float32)
	Fogiv                     func(pname uint32, params *int32)
	FrontFace                 func(mode uint32)
	Frustum                   func(left, right, bottom, top, zNear, zFar float64)
	GenLists                  func(Range int32) uint32
	GetBooleanv               func(pname uint32, params *bool)
	GetDoublev                func(pname uint32, params *float64)
	GetFloatv                 func(pname uint32, params *float32)
	GetIntegerv               func(pname uint32, params *int32)
	GetClipPlane              func(plane uint32, equation *float64)
	GetError                  func() uint32
	GetLightfv                func(light, pname uint32, params *float32)
	GetLightiv                func(light, pname uint32, params *int32)
	GetMapdv                  func(target, query uint32, v *float64)
	GetMapfv                  func(target, query uint32, v *float32)
	GetMapiv                  func(target, query uint32, v *int32)
	GetMaterialfv             func(face, pname uint32, params *float32)
	GetMaterialiv             func(face, pname uint32, params *int32)
	GetPixelMapfv             func(Map uint32, values *float32)
	GetPixelMapuiv            func(Map uint32, values *uint32)
	GetPixelMapusv            func(Map uint32, values *uint16)
	GetPolygonStipple         func(pattern *uint8)
	GetString                 func(name uint32) string
	GetTexEnvfv               func(target, pname uint32, params *float32)
	GetTexEnviv               func(target, pname uint32, params *int32)
	GetTexGendv               func(coord, pname uint32, params *float64)
	GetTexGenfv               func(coord, pname uint32, params *float32)
	GetTexGeniv               func(coord, pname uint32, params *int32)
	GetTexImage               func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer)
	GetTexLevelParameterfv    func(target uint32, level int32, pname uint32, params *float32)
	GetTexLevelParameteriv    func(target uint32, level int32, pname uint32, params *int32)
	GetTexParameterfv         func(target, pname uint32, params *float32)
	GetTexParameteriv         func(target, pname uint32, params *int32)
	Hint                      func(target, mode uint32)
	Indexd                    func(c float64)
	Indexf                    func(c float32)
	Indexi                    func(c int32)
	Indexs                    func(c int16)
	Indexdv                   func(c *float64)
	Indexfv                   func(c *float32)
	Indexiv                   func(c *int32)
	Indexsv                   func(c *int16)
	IndexMask                 func(mask uint32)
	IndexPointer              func(Type uint32, stride int32, pointer unsafe.Pointer)
	InitNames                 func()
	IsEnabled                 func(cap uint32)
	IsList                    func(list uint32) bool
	Lightf                    func(light, pname uint32, param float32)
	Lighti                    func(light, pname uint32, param int32)
	Lightfv                   func(light, pname uint32, params *float32)
	Lightiv                   func(light, pname uint32, params *int32)
	LightModelf               func(pname uint32, param float32)
	LightModeli               func(pname uint32, param int32)
	LightModelfv              func(pname uint32, params *float32)
	LightModeliv              func(pname uint32, params *int32)
	LineStipple               func(factor int32, pattern uint16)
	LineWidth                 func(width float32)
	ListBase                  func(base uint32)
	LoadIdentity              func()
	LoadMatrixd               func(m *float64)
	LoadMatrixf               func(m *float32)
	LoadName                  func(name uint32)
	LogicOp                   func(opcode uint32)
	Map1d                     func(target uint32, u1, u2 float64, stride, order int32, points *float64)
	Map1f                     func(target uint32, u1, u2 float32, stride, order int32, points *float32)
	Map2d                     func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64)
	Map2f                     func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32)
	MapGrid1d                 func(un int32, u1, u2 float64)
	MapGrid1f                 func(un int32, u1, u2 float32)
	MapGrid2d                 func(un int32, u1, u2 float64, vn int32, v1, v2 float64)
	MapGrid2f                 func(un int32, u1, u2 float32, vn int32, v1, v2 float32)
	Materialf                 func(face, pname uint32, param float32)
	Materiali                 func(face, pname uint32, param int32)
	Materialfv                func(face, pname uint32, params *float32)
	Materialiv                func(face, pname uint32, params *int32)
	MatrixMode                func(mode uint32)
	MultMatrixd               func(m *float64)
	MultMatrixf               func(m *float32)
	NewList                   func(list uint32, mode uint32)
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
	PixelMapfv                func(Map uint32, mapsize int32, values *float32)
	PixelMapuiv               func(Map uint32, mapsize int32, values *uint32)
	PixelMapusv               func(Map uint32, mapsize int32, values *uint16)
	PixelStoref               func(pname uint32, param float32)
	PixelStorei               func(pname uint32, param int32)
	PixelTransferf            func(pname uint32, param float32)
	PixelTransferi            func(pname uint32, param int32)
	PixelZoom                 func(xfactor, yfactor float32)
	PointSize                 func(size float32)
	PolygonMode               func(face, mode uint32)
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
	ReadBuffer                func(mode uint32)
	ReadPixels                func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer)
	Rectd                     func(x1, y1, x2, y2 float64)
	Rectf                     func(x1, y1, x2, y2 float32)
	Recti                     func(x1, y1, x2, y2 int32)
	Rects                     func(x1, y1, x2, y2 int16)
	Rectdv                    func(v1, v2 *float64)
	Rectfv                    func(v1, v2 *float32)
	Rectiv                    func(v1, v2 *int32)
	Rectsv                    func(v1, v2 *int16)
	RenderMode                func(mode uint32) int32
	Rotated                   func(angle, x, y, z float64)
	Rotatef                   func(angle, x, y, z float32)
	Scaled                    func(x, y, z float64)
	Scalef                    func(x, y, z float32)
	Scissor                   func(x, y int32, width, height int32)
	SelectBuffer              func(size int32, buffer *uint32)
	ShadeModel                func(mode uint32)
	StencilFunc               func(Func uint32, ref int32, mask uint32)
	StencilMask               func(mask uint32)
	StencilOp                 func(fail, zfail, zpass uint32)
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
	TexEnvf                   func(target, pname uint32, param float32)
	TexEnvi                   func(target, pname uint32, param int32)
	TexEnvfv                  func(target, pname uint32, params *float32)
	TexEnviv                  func(target, pname uint32, params *int32)
	TexGend                   func(coord, pname uint32, param float64)
	TexGenf                   func(coord, pname uint32, param float32)
	TexGeni                   func(coord, pname uint32, param int32)
	TexGendv                  func(coord, pname uint32, params *float64)
	TexGenfv                  func(coord, pname uint32, params *float32)
	TexGeniv                  func(coord, pname uint32, params *int32)
	TexImage1D                func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer)
	TexImage2D                func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer)
	TexParameterf             func(target, pname uint32, param float32)
	TexParameteri             func(target, pname uint32, param int32)
	TexParameterfv            func(target, pname uint32, params *float32)
	TexParameteriv            func(target, pname uint32, params *int32)
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
	GetConvolutionParameterfv func(target, pname uint32, params *float32)
	GetConvolutionParameteriv func(target, pname uint32, params *int32)
	AreTexturesResident       func(textures []uint32) (status bool, residencies []bool)
	ArrayElement              func(i int32)
	DrawArrays                func(mode uint32, first int32, count int32)
	DrawElements              func(mode uint32, count int32, Type uint32, indices unsafe.Pointer)
	GetPointerv               func(pname uint32, params unsafe.Pointer)
	PolygonOffset             func(factor, units float32)
	CopyTexImage1D            func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32)
	CopyTexImage2D            func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32)
	CopyTexSubImage1D         func(target uint32, level, xoffset, x, y int32, width int32)
	CopyTexSubImage2D         func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32)
	BindTexture               func(target uint32, texture uint32)
	DeleteTextures            func(n int32, textures *uint32)
	GenTextures               func(n int32, textures *uint32)
	IsTexture                 func(texture uint32) bool
	ColorPointer              func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
	EnableClientState         func(cap uint32)
	DisableClientState        func(cap uint32)
	Indexub                   func(c uint8)
	Indexubv                  func(c *uint8)
	InterleavedArrays         func(format uint32, stride int32, pointer unsafe.Pointer)
	NormalPointer             func(Type uint32, stride int32, pointer unsafe.Pointer)
	PushClientAttrib          func(mask uint32)
	PrioritizeTextures        func(n int32, textures *uint32, priorities *float32)
	PopClientAttrib           func()
	TexCoordPointer           func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
	TexSubImage1D             func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer)
	TexSubImage2D             func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer)
	VertexPointer             func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
	ColorTable                func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer)
	ColorTableParameterfv     func(target, pname uint32, params *float32)
	ColorTableParameteriv     func(target, pname uint32, params *int32)
	ColorSubTable             func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionFilter1D       func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionFilter2D       func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionParameterf     func(target, pname uint32, params float32)
	ConvolutionParameteri     func(target, pname uint32, params int32)
	CopyColorTable            func(target, internalformat uint32, x, y int32, width int32)
	CopyColorSubTable         func(target uint32, start int32, x, y int32, width int32)
	CopyConvolutionFilter1D   func(target, internalformat uint32, x, y int32, width int32)
	CopyConvolutionFilter2D   func(target, internalformat uint32, x, y int32, width, height int32)
	GetColorTable             func(target, format, Type uint32, table unsafe.Pointer)
	GetColorTableParameterfv  func(target, pname uint32, params *float32)
	GetColorTableParameteriv  func(target, pname uint32, params *int32)
	GetConvolutionFilter      func(target, format, Type uint32, image unsafe.Pointer)
	GetHistogram              func(target uint32, reset bool, format, Type uint32, values unsafe.Pointer)
	GetHistogramParameterfv   func(target, pname uint32, params *float32)
	GetHistogramParameteriv   func(target, pname uint32, params *int32)
	GetSeparableFilter        func(target, format, Type uint32, row, column, span unsafe.Pointer)
	Histogram                 func(target uint32, width int32, internalformat uint32, sink bool)
	Minmax                    func(target, internalformat uint32, sink bool)
	MultiTexCoord1s           func(target uint32, s int16)
	MultiTexCoord1i           func(target uint32, s int32)
	MultiTexCoord1f           func(target uint32, s float32)
	MultiTexCoord1d           func(target uint32, s float64)
	MultiTexCoord2s           func(target uint32, s, t int16)
	MultiTexCoord2i           func(target uint32, s, t int32)
	MultiTexCoord2f           func(target uint32, s, t float32)
	MultiTexCoord2d           func(target uint32, s, t float64)
	MultiTexCoord3s           func(target uint32, s, t, r int16)
	MultiTexCoord3i           func(target uint32, s, t, r int32)
	MultiTexCoord3f           func(target uint32, s, t, r float32)
	MultiTexCoord3d           func(target uint32, s, t, r float64)
	MultiTexCoord4s           func(target uint32, s, t, r, q int16)
	MultiTexCoord4i           func(target uint32, s, t, r, q int32)
	MultiTexCoord4f           func(target uint32, s, t, r, q float32)
	MultiTexCoord4d           func(target uint32, s, t, r, q float64)
	MultiTexCoord1sv          func(target uint32, v *int16)
	MultiTexCoord1iv          func(target uint32, v *int32)
	MultiTexCoord1fv          func(target uint32, v *float32)
	MultiTexCoord1dv          func(target uint32, v *float64)
	MultiTexCoord2sv          func(target uint32, v *int16)
	MultiTexCoord2iv          func(target uint32, v *int32)
	MultiTexCoord2fv          func(target uint32, v *float32)
	MultiTexCoord2dv          func(target uint32, v *float64)
	MultiTexCoord3sv          func(target uint32, v *int16)
	MultiTexCoord3iv          func(target uint32, v *int32)
	MultiTexCoord3fv          func(target uint32, v *float32)
	MultiTexCoord3dv          func(target uint32, v *float64)
	MultiTexCoord4sv          func(target uint32, v *int16)
	MultiTexCoord4iv          func(target uint32, v *int32)
	MultiTexCoord4fv          func(target uint32, v *float32)
	MultiTexCoord4dv          func(target uint32, v *float64)
	ResetHistogram            func(target uint32)
	ResetMinmax               func(target uint32)
	SeparableFilter2D         func(target, internalformat uint32, width, height int32, format, Type uint32, row, column unsafe.Pointer)
	BlendColor                func(red, green, blue, alpha float32)
	BlendEquation             func(mode uint32)
	CopyTexSubImage3D         func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32)
	DrawRangeElements         func(mode uint32, start, end uint32, count int32, Type uint32, indices unsafe.Pointer)
	TexImage3D                func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer)
	TexSubImage3D             func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer)
	ActiveTexture             func(texture uint32)
	ClientActiveTexture       func(texture uint32)
	CompressedTexImage1D      func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage2D      func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage3D      func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage1D   func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage2D   func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage3D   func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer)
	GetCompressedTexImage     func(target uint32, lod int32, img unsafe.Pointer)
	LoadTransposeMatrixd      func(m *float64)
	LoadTransposeMatrixf      func(m *float64)
	MultTransposeMatrixd      func(m *float64)
	MultTransposeMatrixf      func(m *float32)
	SampleCoverage            func(value float32, invert bool)
	BlendFuncSeparate         func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32)
	FogCoordPointer           func(Type uint32, stride int32, pointer unsafe.Pointer)
	FogCoordd                 func(coord float64)
	FogCoordf                 func(coord float32)
	FogCoorddv                func(coord *float64)
	FogCoordfv                func(coord *float32)
	MultiDrawArrays           func(mode uint32, first *int32, count *int32, primcount int32)
	MultiDrawElements         func(mode uint32, count *int32, Type uint32, indices unsafe.Pointer, primcount int32)
	PointParameterf           func(pname uint32, param float32)
	PointParameteri           func(pname uint32, param int32)
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
	SecondaryColorPointer     func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
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
	BeginQuery                func(target uint32, id uint32)
	BindBuffer                func(target uint32, buffer uint32)
	BufferData                func(target uint32, size int32, data unsafe.Pointer, usage uint32)
	BufferSubData             func(target, offset uint32, size int32, data unsafe.Pointer)
	DeleteBuffers             func(n int32, buffers *uint32)
	DeleteQueries             func(n int32, ids *uint32)
	GenBuffers                func(n int32, buffers *uint32)
	GenQueries                func(n int32, ids *uint32)
	GetBufferParameteriv      func(target, value uint32, data *int32)
	GetBufferPointerv         func(target, pname uint32, params unsafe.Pointer)
	GetBufferSubData          func(target uint32, offset int32, size int32, data unsafe.Pointer)
	GetQueryObjectiv          func(id uint32, pname uint32, params *int32)
	GetQueryObjectuiv         func(id uint32, pname uint32, params *uint32)
	GetQueryiv                func(target, pname uint32, params *int32)
	IsBuffer                  func(buffer uint32) bool
	IsQuery                   func(id uint32) bool
	MapBuffer                 func(target, access uint32) unsafe.Pointer
	UnmapBuffer               func(target uint32) bool
	AttachShader              func(program, shader uint32)
	BindAttribLocation        func(program, index uint32, name string)
	BlendEquationSeperate     func(modeRGB, modeAlpha uint32)
	CompileShader             func(shader uint32)
	CreateProgram             func() uint32
	CreateShader              func(shaderType uint32) uint32
	DeleteProgram             func(program uint32)
	DeleteShader              func(shader uint32)
	DetachShader              func(program, shader uint32)
	EnableVertexAttribArray   func(index uint32)
	DisableVertexAttribArray  func(index uint32)
	DrawBuffers               func(n int32, bufs *uint32)
	GetActiveAttrib           func(program, index uint32, bufSize int32) (length int32, size int32, Type uint32, name string)
	GetActiveUniform          func(program, index uint32, bufSize int32, length *int32, size *int32, Type *uint32, name *byte)
	GetAttachedShaders        func(program uint32, maxCount int32, count *int32, shaders *uint32)
	GetAttribLocation         func(program uint32, name *byte) int32
	GetProgramiv              func(program uint32, pname uint32, params *int32)
	GetProgramInfoLog         func(program uint32, maxLength int32, length *int32, infoLog *byte)
	GetShaderiv               func(program uint32, pname uint32, params *int32)
	GetShaderInfoLog          func(shader uint32, maxLength int32, length *int32, infoLog *byte)
	GetShaderSource           func(shader uint32, bufSize int32, length *int32, source *byte)
	GetUniformfv              func(program uint32, location int32, params *float32)
	GetUniformiv              func(program uint32, location int32, params *int32)
	GetUniformLocation        func(program uint32, name *byte) int32
	GetVertexAttribdv         func(index uint32, pname uint32, params *float64)
	GetVertexAttribfv         func(index uint32, pname uint32, params *float32)
	GetVertexAttribiv         func(index uint32, pname uint32, params *int32)
	GetVertexAttribPointerv   func(index uint32, pname uint32, pointer unsafe.Pointer)
	IsProgram                 func(program uint32) bool
	IsShader                  func(shader uint32) bool
	LinkProgram               func(program uint32)
	ShaderSource              func(shader uint32, count int32, string **byte, length *int32)
	StencilFuncSeparate       func(face, Func uint32, ref int32, mask uint32)
	StencilMaskSeparate       func(face uint32, mask uint32)
	StencilOpSeparate         func(face, sfail, dpfail, dppass uint32)
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
	VertexAttribPointer       func(index uint32, size int32, Type uint32, normalized bool, stride int32, pointer unsafe.Pointer)
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
	glc.context = C.gl14NewContext()

	glc.Accum = func(op uint32, value float32) {
		C.gl14Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		C.gl14AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		glc.inBeginEnd = true
		C.gl14Begin(glc.context, C.GLenum(mode))
		return
	}

	glc.End = func() {
		C.gl14End(glc.context)
		glc.inBeginEnd = false
		return
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8) {
		C.gl14Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		C.gl14BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		C.gl14CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		C.gl14CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		C.gl14Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		C.gl14ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		C.gl14ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		C.gl14ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		C.gl14ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		C.gl14ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation *float64) {
		C.gl14ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.Color3b = func(red, green, blue int8) {
		C.gl14Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		C.gl14Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		C.gl14Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		C.gl14Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		C.gl14Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		C.gl14Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		C.gl14Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		C.gl14Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		C.gl14Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		C.gl14Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		C.gl14Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		C.gl14Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		C.gl14Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		C.gl14Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		C.gl14Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		C.gl14Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v *int8) {
		C.gl14Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color3dv = func(v *float64) {
		C.gl14Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color3fv = func(v *float32) {
		C.gl14Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color3iv = func(v *int32) {
		C.gl14Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color3sv = func(v *int16) {
		C.gl14Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color3ubv = func(v *uint8) {
		C.gl14Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color3uiv = func(v *uint32) {
		C.gl14Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color3usv = func(v *uint16) {
		C.gl14Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.Color4bv = func(v *int8) {
		C.gl14Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color4dv = func(v *float64) {
		C.gl14Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color4fv = func(v *float32) {
		C.gl14Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color4iv = func(v *int32) {
		C.gl14Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color4sv = func(v *int16) {
		C.gl14Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color4ubv = func(v *uint8) {
		C.gl14Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color4uiv = func(v *uint32) {
		C.gl14Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color4usv = func(v *uint16) {
		C.gl14Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		C.gl14ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		C.gl14ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		C.gl14CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		C.gl14CullFace(glc.context, C.GLenum(mode))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		C.gl14DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		C.gl14DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthMask = func(flag bool) {
		C.gl14DepthMask(glc.context, boolToGL(flag))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		C.gl14DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		C.gl14Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		C.gl14Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		C.gl14DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl14DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.EdgeFlag = func(flag bool) {
		C.gl14EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag *bool) {
		C.gl14EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(flag)))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		C.gl14EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EvalCoord1d = func(u float64) {
		C.gl14EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		C.gl14EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		C.gl14EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		C.gl14EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u *float64) {
		C.gl14EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord1fv = func(u *float32) {
		C.gl14EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2dv = func(u *float64) {
		C.gl14EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2fv = func(u *float32) {
		C.gl14EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		C.gl14EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		C.gl14EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		C.gl14EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		C.gl14EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer *float32) {
		C.gl14FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(buffer)))
	}

	glc.Finish = func() {
		C.gl14Finish(glc.context)
	}

	glc.Flush = func() {
		C.gl14Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		C.gl14Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		C.gl14Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params *float32) {
		C.gl14Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Fogiv = func(pname uint32, params *int32) {
		C.gl14Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.FrontFace = func(mode uint32) {
		C.gl14FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		C.gl14Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		return uint32(C.gl14GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params *bool) {
		C.gl14GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(params)))
	}

	glc.GetDoublev = func(pname uint32, params *float64) {
		C.gl14GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetFloatv = func(pname uint32, params *float32) {
		C.gl14GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetIntegerv = func(pname uint32, params *int32) {
		C.gl14GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetClipPlane = func(plane uint32, equation *float64) {
		C.gl14GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl14GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params *float32) {
		C.gl14GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetLightiv = func(light, pname uint32, params *int32) {
		C.gl14GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetMapdv = func(target, query uint32, v *float64) {
		C.gl14GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.GetMapfv = func(target, query uint32, v *float32) {
		C.gl14GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.GetMapiv = func(target, query uint32, v *int32) {
		C.gl14GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.GetMaterialfv = func(face, pname uint32, params *float32) {
		C.gl14GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetMaterialiv = func(face, pname uint32, params *int32) {
		C.gl14GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetPixelMapfv = func(Map uint32, values *float32) {
		C.gl14GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapuiv = func(Map uint32, values *uint32) {
		C.gl14GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapusv = func(Map uint32, values *uint16) {
		C.gl14GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.GetPolygonStipple = func(pattern *uint8) {
		C.gl14GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(pattern)))
	}

	glc.GetString = func(name uint32) string {
		cstr := C.gl14GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params *float32) {
		C.gl14GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexEnviv = func(target, pname uint32, params *int32) {
		C.gl14GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexGendv = func(coord, pname uint32, params *float64) {
		C.gl14GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params *float32) {
		C.gl14GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params *int32) {
		C.gl14GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params *float32) {
		C.gl14GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params *int32) {
		C.gl14GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params *float32) {
		C.gl14GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params *int32) {
		C.gl14GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Hint = func(target, mode uint32) {
		C.gl14Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		C.gl14Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		C.gl14Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		C.gl14Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		C.gl14Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexdv = func(c *float64) {
		C.gl14Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(c)))
	}

	glc.Indexfv = func(c *float32) {
		C.gl14Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(c)))
	}

	glc.Indexiv = func(c *int32) {
		C.gl14Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(c)))
	}

	glc.Indexsv = func(c *int16) {
		C.gl14Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(c)))
	}

	glc.IndexMask = func(mask uint32) {
		C.gl14IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		C.gl14InitNames(glc.context)
	}

	glc.IsEnabled = func(cap uint32) {
		C.gl14IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		return C.gl14IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		C.gl14Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		C.gl14Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params *float32) {
		C.gl14Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Lightiv = func(light, pname uint32, params *int32) {
		C.gl14Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		C.gl14LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		C.gl14LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params *float32) {
		C.gl14LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.LightModeliv = func(pname uint32, params *int32) {
		C.gl14LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		C.gl14LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		C.gl14LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		C.gl14ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		C.gl14LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m *float64) {
		C.gl14LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadMatrixf = func(m *float32) {
		C.gl14LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.LoadName = func(name uint32) {
		C.gl14LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		C.gl14LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points *float64) {
		C.gl14Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points *float32) {
		C.gl14Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64) {
		C.gl14Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32) {
		C.gl14Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		C.gl14MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		C.gl14MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		C.gl14MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		C.gl14MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		C.gl14Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		C.gl14Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params *float32) {
		C.gl14Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Materialiv = func(face, pname uint32, params *int32) {
		C.gl14Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.MatrixMode = func(mode uint32) {
		C.gl14MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m *float64) {
		C.gl14MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultMatrixf = func(m *float32) {
		C.gl14MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.NewList = func(list uint32, mode uint32) {
		C.gl14NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		C.gl14EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		C.gl14Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		C.gl14Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		C.gl14Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		C.gl14Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		C.gl14Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v *int8) {
		C.gl14Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Normal3dv = func(v *float64) {
		C.gl14Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Normal3fv = func(v *float32) {
		C.gl14Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Normal3iv = func(v *int32) {
		C.gl14Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Normal3sv = func(v *int16) {
		C.gl14Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		C.gl14Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		C.gl14PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values *float32) {
		C.gl14PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values *uint32) {
		C.gl14PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values *uint16) {
		C.gl14PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		C.gl14PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		C.gl14PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		C.gl14PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		C.gl14PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		C.gl14PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		C.gl14PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		C.gl14PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask *uint8) {
		C.gl14PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(mask)))
	}

	glc.PushAttrib = func(mask uint32) {
		C.gl14PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		C.gl14PopAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		C.gl14PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		C.gl14PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		C.gl14PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		C.gl14PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		C.gl14RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		C.gl14RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		C.gl14RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		C.gl14RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		C.gl14RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		C.gl14RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		C.gl14RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		C.gl14RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		C.gl14RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		C.gl14RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		C.gl14RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		C.gl14RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v *float64) {
		C.gl14RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos2fv = func(v *float32) {
		C.gl14RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos2iv = func(v *int32) {
		C.gl14RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos2sv = func(v *int16) {
		C.gl14RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos3dv = func(v *float64) {
		C.gl14RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos3fv = func(v *float32) {
		C.gl14RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos3iv = func(v *int32) {
		C.gl14RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos3sv = func(v *int16) {
		C.gl14RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos4dv = func(v *float64) {
		C.gl14RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos4fv = func(v *float32) {
		C.gl14RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos4iv = func(v *int32) {
		C.gl14RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos4sv = func(v *int16) {
		C.gl14RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.ReadBuffer = func(mode uint32) {
		C.gl14ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		C.gl14Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		C.gl14Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		C.gl14Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		C.gl14Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 *float64) {
		C.gl14Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
	}

	glc.Rectfv = func(v1, v2 *float32) {
		C.gl14Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
	}

	glc.Rectiv = func(v1, v2 *int32) {
		C.gl14Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
	}

	glc.Rectsv = func(v1, v2 *int16) {
		C.gl14Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
	}

	glc.RenderMode = func(mode uint32) int32 {
		return int32(C.gl14RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		C.gl14Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		C.gl14Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		C.gl14Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		C.gl14Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		C.gl14Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer *uint32) {
		C.gl14SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
	}

	glc.ShadeModel = func(mode uint32) {
		C.gl14ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		C.gl14StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		C.gl14StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		C.gl14StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		C.gl14TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		C.gl14TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		C.gl14TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		C.gl14TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		C.gl14TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		C.gl14TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		C.gl14TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		C.gl14TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		C.gl14TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		C.gl14TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		C.gl14TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		C.gl14TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		C.gl14TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		C.gl14TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		C.gl14TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		C.gl14TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v *float64) {
		C.gl14TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord1fv = func(v *float32) {
		C.gl14TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord1iv = func(v *int32) {
		C.gl14TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord1sv = func(v *int16) {
		C.gl14TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord2dv = func(v *float64) {
		C.gl14TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord2fv = func(v *float32) {
		C.gl14TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord2iv = func(v *int32) {
		C.gl14TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord2sv = func(v *int16) {
		C.gl14TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord3dv = func(v *float64) {
		C.gl14TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord3fv = func(v *float32) {
		C.gl14TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord3iv = func(v *int32) {
		C.gl14TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord3sv = func(v *int16) {
		C.gl14TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord4dv = func(v *float64) {
		C.gl14TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord4fv = func(v *float32) {
		C.gl14TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord4iv = func(v *int32) {
		C.gl14TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord4sv = func(v *int16) {
		C.gl14TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		C.gl14TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		C.gl14TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params *float32) {
		C.gl14TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexEnviv = func(target, pname uint32, params *int32) {
		C.gl14TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		C.gl14TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		C.gl14TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		C.gl14TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params *float64) {
		C.gl14TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.TexGenfv = func(coord, pname uint32, params *float32) {
		C.gl14TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexGeniv = func(coord, pname uint32, params *int32) {
		C.gl14TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		C.gl14TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		C.gl14TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params *float32) {
		C.gl14TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexParameteriv = func(target, pname uint32, params *int32) {
		C.gl14TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Translated = func(x, y, z float64) {
		C.gl14Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		C.gl14Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		C.gl14Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		C.gl14Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		C.gl14Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		C.gl14Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		C.gl14Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		C.gl14Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		C.gl14Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		C.gl14Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		C.gl14Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		C.gl14Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		C.gl14Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		C.gl14Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		C.gl14Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetConvolutionParameterfv = func(target, pname uint32, params *float32) {
		C.gl14GetConvolutionParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionParameteriv = func(target, pname uint32, params *int32) {
		C.gl14GetConvolutionParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		var cRes *C.GLboolean
		status = C.gl14AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		C.gl14ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		C.gl14DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl14DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		C.gl14GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		C.gl14PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		C.gl14CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		C.gl14CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		C.gl14CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		C.gl14CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		C.gl14BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures *uint32) {
		C.gl14DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.GenTextures = func(n int32, textures *uint32) {
		C.gl14GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.IsTexture = func(texture uint32) bool {
		return C.gl14IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		C.gl14EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		C.gl14DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.Indexub = func(c uint8) {
		C.gl14Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexubv = func(c *uint8) {
		C.gl14Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(c)))
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.PushClientAttrib = func(mask uint32) {
		C.gl14PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PrioritizeTextures = func(n int32, textures *uint32, priorities *float32) {
		C.gl14PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLclampf)(unsafe.Pointer(priorities)))
	}

	glc.PopClientAttrib = func() {
		C.gl14PopClientAttrib(glc.context)
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl14ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params *float32) {
		C.gl14ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params *int32) {
		C.gl14ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		C.gl14ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl14ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl14ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		C.gl14ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		C.gl14ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl14CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		C.gl14CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl14CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		C.gl14CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetColorTable = func(target, format, Type uint32, table unsafe.Pointer) {
		C.gl14GetColorTable(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), table)
	}

	glc.GetColorTableParameterfv = func(target, pname uint32, params *float32) {
		C.gl14GetColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetColorTableParameteriv = func(target, pname uint32, params *int32) {
		C.gl14GetColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionFilter = func(target, format, Type uint32, image unsafe.Pointer) {
		C.gl14GetConvolutionFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), image)
	}

	glc.GetHistogram = func(target uint32, reset bool, format, Type uint32, values unsafe.Pointer) {
		C.gl14GetHistogram(glc.context, C.GLenum(target), boolToGL(reset), C.GLenum(format), C.GLenum(Type), values)
	}

	glc.GetHistogramParameterfv = func(target, pname uint32, params *float32) {
		C.gl14GetHistogramParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetHistogramParameteriv = func(target, pname uint32, params *int32) {
		C.gl14GetHistogramParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetSeparableFilter = func(target, format, Type uint32, row, column, span unsafe.Pointer) {
		C.gl14GetSeparableFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), row, column, span)
	}

	glc.Histogram = func(target uint32, width int32, internalformat uint32, sink bool) {
		C.gl14Histogram(glc.context, C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.Minmax = func(target, internalformat uint32, sink bool) {
		C.gl14Minmax(glc.context, C.GLenum(target), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.MultiTexCoord1s = func(target uint32, s int16) {
		C.gl14MultiTexCoord1s(glc.context, C.GLenum(target), C.GLshort(s))
	}

	glc.MultiTexCoord1i = func(target uint32, s int32) {
		C.gl14MultiTexCoord1i(glc.context, C.GLenum(target), C.GLint(s))
	}

	glc.MultiTexCoord1f = func(target uint32, s float32) {
		C.gl14MultiTexCoord1f(glc.context, C.GLenum(target), C.GLfloat(s))
	}

	glc.MultiTexCoord1d = func(target uint32, s float64) {
		C.gl14MultiTexCoord1d(glc.context, C.GLenum(target), C.GLdouble(s))
	}

	glc.MultiTexCoord2s = func(target uint32, s, t int16) {
		C.gl14MultiTexCoord2s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t))
	}

	glc.MultiTexCoord2i = func(target uint32, s, t int32) {
		C.gl14MultiTexCoord2i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t))
	}

	glc.MultiTexCoord2f = func(target uint32, s, t float32) {
		C.gl14MultiTexCoord2f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
	}

	glc.MultiTexCoord2d = func(target uint32, s, t float64) {
		C.gl14MultiTexCoord2d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
	}

	glc.MultiTexCoord3s = func(target uint32, s, t, r int16) {
		C.gl14MultiTexCoord3s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.MultiTexCoord3i = func(target uint32, s, t, r int32) {
		C.gl14MultiTexCoord3i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.MultiTexCoord3f = func(target uint32, s, t, r float32) {
		C.gl14MultiTexCoord3f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.MultiTexCoord3d = func(target uint32, s, t, r float64) {
		C.gl14MultiTexCoord3d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.MultiTexCoord4s = func(target uint32, s, t, r, q int16) {
		C.gl14MultiTexCoord4s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.MultiTexCoord4i = func(target uint32, s, t, r, q int32) {
		C.gl14MultiTexCoord4i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.MultiTexCoord4f = func(target uint32, s, t, r, q float32) {
		C.gl14MultiTexCoord4f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.MultiTexCoord4d = func(target uint32, s, t, r, q float64) {
		C.gl14MultiTexCoord4d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.MultiTexCoord1sv = func(target uint32, v *int16) {
		C.gl14MultiTexCoord1sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1iv = func(target uint32, v *int32) {
		C.gl14MultiTexCoord1iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1fv = func(target uint32, v *float32) {
		C.gl14MultiTexCoord1fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1dv = func(target uint32, v *float64) {
		C.gl14MultiTexCoord1dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2sv = func(target uint32, v *int16) {
		C.gl14MultiTexCoord2sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2iv = func(target uint32, v *int32) {
		C.gl14MultiTexCoord2iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2fv = func(target uint32, v *float32) {
		C.gl14MultiTexCoord2fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2dv = func(target uint32, v *float64) {
		C.gl14MultiTexCoord2dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3sv = func(target uint32, v *int16) {
		C.gl14MultiTexCoord3sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3iv = func(target uint32, v *int32) {
		C.gl14MultiTexCoord3iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3fv = func(target uint32, v *float32) {
		C.gl14MultiTexCoord3fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3dv = func(target uint32, v *float64) {
		C.gl14MultiTexCoord3dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4sv = func(target uint32, v *int16) {
		C.gl14MultiTexCoord4sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4iv = func(target uint32, v *int32) {
		C.gl14MultiTexCoord4iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4fv = func(target uint32, v *float32) {
		C.gl14MultiTexCoord4fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4dv = func(target uint32, v *float64) {
		C.gl14MultiTexCoord4dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.ResetHistogram = func(target uint32) {
		C.gl14ResetHistogram(glc.context, C.GLenum(target))
	}

	glc.ResetMinmax = func(target uint32) {
		C.gl14ResetMinmax(glc.context, C.GLenum(target))
	}

	glc.SeparableFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, row, column unsafe.Pointer) {
		C.gl14SeparableFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), row, column)
	}

	glc.BlendColor = func(red, green, blue, alpha float32) {
		C.gl14BlendColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		C.gl14BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		C.gl14CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DrawRangeElements = func(mode uint32, start, end uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl14DrawRangeElements(glc.context, C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.TexImage3D = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14TexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl14TexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.ActiveTexture = func(texture uint32) {
		C.gl14ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		C.gl14ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl14CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl14CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl14CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl14CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl14CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl14CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.GetCompressedTexImage = func(target uint32, lod int32, img unsafe.Pointer) {
		C.gl14GetCompressedTexImage(glc.context, C.GLenum(target), C.GLint(lod), img)
	}

	glc.LoadTransposeMatrixd = func(m *float64) {
		C.gl14LoadTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadTransposeMatrixf = func(m *float64) {
		C.gl14LoadTransposeMatrixf(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixd = func(m *float64) {
		C.gl14MultTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixf = func(m *float32) {
		C.gl14MultTransposeMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.SampleCoverage = func(value float32, invert bool) {
		C.gl14SampleCoverage(glc.context, C.GLclampf(value), boolToGL(invert))
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		C.gl14BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.FogCoordPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14FogCoordPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.FogCoordd = func(coord float64) {
		C.gl14FogCoordd(glc.context, C.GLdouble(coord))
	}

	glc.FogCoordf = func(coord float32) {
		C.gl14FogCoordf(glc.context, C.GLfloat(coord))
	}

	glc.FogCoorddv = func(coord *float64) {
		C.gl14FogCoorddv(glc.context, (*C.GLdouble)(unsafe.Pointer(coord)))
	}

	glc.FogCoordfv = func(coord *float32) {
		C.gl14FogCoordfv(glc.context, (*C.GLfloat)(unsafe.Pointer(coord)))
	}

	glc.MultiDrawArrays = func(mode uint32, first *int32, count *int32, primcount int32) {
		C.gl14MultiDrawArrays(glc.context, C.GLenum(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), C.GLsizei(primcount))
	}

	glc.MultiDrawElements = func(mode uint32, count *int32, Type uint32, indices unsafe.Pointer, primcount int32) {
		C.gl14MultiDrawElements(glc.context, C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(count)), C.GLenum(Type), indices, C.GLsizei(primcount))
	}

	glc.PointParameterf = func(pname uint32, param float32) {
		C.gl14PointParameterf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PointParameteri = func(pname uint32, param int32) {
		C.gl14PointParameteri(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.SecondaryColor3b = func(red, green, blue int8) {
		C.gl14SecondaryColor3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.SecondaryColor3s = func(red, green, blue int16) {
		C.gl14SecondaryColor3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.SecondaryColor3i = func(red, green, blue int32) {
		C.gl14SecondaryColor3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.SecondaryColor3f = func(red, green, blue float32) {
		C.gl14SecondaryColor3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.SecondaryColor3d = func(red, green, blue float64) {
		C.gl14SecondaryColor3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.SecondaryColor3ub = func(red, green, blue uint8) {
		C.gl14SecondaryColor3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.SecondaryColor3us = func(red, green, blue uint16) {
		C.gl14SecondaryColor3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.SecondaryColor3ui = func(red, green, blue uint32) {
		C.gl14SecondaryColor3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.SecondaryColor3bv = func(v *int8) {
		C.gl14SecondaryColor3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3sv = func(v *int16) {
		C.gl14SecondaryColor3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3iv = func(v *int32) {
		C.gl14SecondaryColor3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3fv = func(v *float32) {
		C.gl14SecondaryColor3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3dv = func(v *float64) {
		C.gl14SecondaryColor3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3ubv = func(v *uint8) {
		C.gl14SecondaryColor3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3usv = func(v *uint16) {
		C.gl14SecondaryColor3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3uiv = func(v *uint32) {
		C.gl14SecondaryColor3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl14SecondaryColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.WindowPos2s = func(x, y int16) {
		C.gl14WindowPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.WindowPos2i = func(x, y int32) {
		C.gl14WindowPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.WindowPos2f = func(x, y float32) {
		C.gl14WindowPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.WindowPos2d = func(x, y float64) {
		C.gl14WindowPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.WindowPos3s = func(x, y, z int16) {
		C.gl14WindowPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.WindowPos3i = func(x, y, z int32) {
		C.gl14WindowPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.WindowPos3f = func(x, y, z float32) {
		C.gl14WindowPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.WindowPos3d = func(x, y, z float64) {
		C.gl14WindowPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.WindowPos2sv = func(v *int16) {
		C.gl14WindowPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos2iv = func(v *int32) {
		C.gl14WindowPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos2fv = func(v *float32) {
		C.gl14WindowPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos2dv = func(v *float64) {
		C.gl14WindowPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.WindowPos3sv = func(v *int16) {
		C.gl14WindowPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos3iv = func(v *int32) {
		C.gl14WindowPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos3fv = func(v *float32) {
		C.gl14WindowPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos3dv = func(v *float64) {
		C.gl14WindowPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		C.gl14BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		C.gl14BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		C.gl14BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		C.gl14BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.DeleteBuffers = func(n int32, buffers *uint32) {
		C.gl14DeleteBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.DeleteQueries = func(n int32, ids *uint32) {
		C.gl14DeleteQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GenBuffers = func(n int32, buffers *uint32) {
		C.gl14GenBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.GenQueries = func(n int32, ids *uint32) {
		C.gl14GenQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GetBufferParameteriv = func(target, value uint32, data *int32) {
		C.gl14GetBufferParameteriv(glc.context, C.GLenum(target), C.GLenum(value), (*C.GLint)(unsafe.Pointer(data)))
	}

	glc.GetBufferPointerv = func(target, pname uint32, params unsafe.Pointer) {
		C.gl14GetBufferPointerv(glc.context, C.GLenum(target), C.GLenum(pname), params)
	}

	glc.GetBufferSubData = func(target uint32, offset int32, size int32, data unsafe.Pointer) {
		C.gl14GetBufferSubData(glc.context, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}

	glc.GetQueryObjectiv = func(id uint32, pname uint32, params *int32) {
		C.gl14GetQueryObjectiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetQueryObjectuiv = func(id uint32, pname uint32, params *uint32) {
		C.gl14GetQueryObjectuiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(params)))
	}

	glc.GetQueryiv = func(target, pname uint32, params *int32) {
		C.gl14GetQueryiv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.IsBuffer = func(buffer uint32) bool {
		return C.gl14IsBuffer(glc.context, C.GLuint(buffer)) != 0
	}

	glc.IsQuery = func(id uint32) bool {
		return C.gl14IsQuery(glc.context, C.GLuint(id)) != 0
	}

	glc.MapBuffer = func(target, access uint32) unsafe.Pointer {
		return unsafe.Pointer(C.gl14MapBuffer(glc.context, C.GLenum(target), C.GLenum(access)))
	}

	glc.UnmapBuffer = func(target uint32) bool {
		return C.gl14UnmapBuffer(glc.context, C.GLenum(target)) != 0
	}

	glc.AttachShader = func(program, shader uint32) {
		C.gl14AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl14BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha uint32) {
		C.gl14BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		C.gl14CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		return uint32(C.gl14CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType uint32) uint32 {
		return uint32(C.gl14CreateShader(glc.context, C.GLenum(shaderType)))
	}

	glc.DeleteProgram = func(program uint32) {
		C.gl14DeleteProgram(glc.context, C.GLuint(program))
	}

	glc.DeleteShader = func(shader uint32) {
		C.gl14DeleteShader(glc.context, C.GLuint(shader))
	}

	glc.DetachShader = func(program, shader uint32) {
		C.gl14DetachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.EnableVertexAttribArray = func(index uint32) {
		C.gl14EnableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DisableVertexAttribArray = func(index uint32) {
		C.gl14DisableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DrawBuffers = func(n int32, bufs *uint32) {
		C.gl14DrawBuffers(glc.context, C.GLsizei(n), (*C.GLenum)(unsafe.Pointer(bufs)))
	}

	glc.GetActiveAttrib = func(program, index uint32, bufSize int32) (length int32, size int32, Type uint32, name string) {
		var (
			cname C.GLchar
		)
		C.gl14GetActiveAttrib(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
		name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
		return
	}

	glc.GetActiveUniform = func(program, index uint32, bufSize int32, length *int32, size *int32, Type *uint32, name *byte) {
		C.gl14GetActiveUniform(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(Type)), (*C.GLchar)(unsafe.Pointer(name)))
	}

	glc.GetAttachedShaders = func(program uint32, maxCount int32, count *int32, shaders *uint32) {
		C.gl14GetAttachedShaders(glc.context, C.GLuint(program), C.GLsizei(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
	}

	glc.GetAttribLocation = func(program uint32, name *byte) int32 {
		return int32(C.gl14GetAttribLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetProgramiv = func(program uint32, pname uint32, params *int32) {
		C.gl14GetProgramiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetProgramInfoLog = func(program uint32, maxLength int32, length *int32, infoLog *byte) {
		C.gl14GetProgramInfoLog(glc.context, C.GLuint(program), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderiv = func(program uint32, pname uint32, params *int32) {
		C.gl14GetShaderiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetShaderInfoLog = func(shader uint32, maxLength int32, length *int32, infoLog *byte) {
		C.gl14GetShaderInfoLog(glc.context, C.GLuint(shader), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderSource = func(shader uint32, bufSize int32, length *int32, source *byte) {
		C.gl14GetShaderSource(glc.context, C.GLuint(shader), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
	}

	glc.GetUniformfv = func(program uint32, location int32, params *float32) {
		C.gl14GetUniformfv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetUniformiv = func(program uint32, location int32, params *int32) {
		C.gl14GetUniformiv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetUniformLocation = func(program uint32, name *byte) int32 {
		return int32(C.gl14GetUniformLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetVertexAttribdv = func(index uint32, pname uint32, params *float64) {
		C.gl14GetVertexAttribdv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribfv = func(index uint32, pname uint32, params *float32) {
		C.gl14GetVertexAttribfv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribiv = func(index uint32, pname uint32, params *int32) {
		C.gl14GetVertexAttribiv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribPointerv = func(index uint32, pname uint32, pointer unsafe.Pointer) {
		C.gl14GetVertexAttribPointerv(glc.context, C.GLuint(index), C.GLenum(pname), pointer)
	}

	glc.IsProgram = func(program uint32) bool {
		return C.gl14IsProgram(glc.context, C.GLuint(program)) != 0
	}

	glc.IsShader = func(shader uint32) bool {
		return C.gl14IsShader(glc.context, C.GLuint(shader)) != 0
	}

	glc.LinkProgram = func(program uint32) {
		C.gl14LinkProgram(glc.context, C.GLuint(program))
	}

	glc.ShaderSource = func(shader uint32, count int32, string **byte, length *int32) {
		C.gl14ShaderSource(glc.context, C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(string)), (*C.GLint)(unsafe.Pointer(length)))
	}

	glc.StencilFuncSeparate = func(face, Func uint32, ref int32, mask uint32) {
		C.gl14StencilFuncSeparate(glc.context, C.GLenum(face), C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMaskSeparate = func(face uint32, mask uint32) {
		C.gl14StencilMaskSeparate(glc.context, C.GLenum(face), C.GLuint(mask))
	}

	glc.StencilOpSeparate = func(face, sfail, dpfail, dppass uint32) {
		C.gl14StencilOpSeparate(glc.context, C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
	}

	glc.Uniform1f = func(location int32, v0 float32) {
		C.gl14Uniform1f(glc.context, C.GLint(location), C.GLfloat(v0))
	}

	glc.Uniform2f = func(location int32, v0, v1 float32) {
		C.gl14Uniform2f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.Uniform3f = func(location int32, v0, v1, v2 float32) {
		C.gl14Uniform3f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Uniform4f = func(location int32, v0, v1, v2, v3 float32) {
		C.gl14Uniform4f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.Uniform1i = func(location, v0 int32) {
		C.gl14Uniform1i(glc.context, C.GLint(location), C.GLint(v0))
	}

	glc.Uniform2i = func(location, v0, v1 int32) {
		C.gl14Uniform2i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1))
	}

	glc.Uniform3i = func(location, v0, v1, v2 int32) {
		C.gl14Uniform3i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
	}

	glc.Uniform4i = func(location, v0, v1, v2, v3 int32) {
		C.gl14Uniform4i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
	}

	glc.Uniform1ui = func(location int32, v0 uint32) {
		C.gl14Uniform1ui(glc.context, C.GLint(location), C.GLuint(v0))
	}

	glc.Uniform2ui = func(location int32, v0, v1 uint32) {
		C.gl14Uniform2ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1))
	}

	glc.Uniform3ui = func(location int32, v0, v1, v2 uint32) {
		C.gl14Uniform3ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
	}

	glc.Uniform4ui = func(location int32, v0, v1, v2, v3 uint32) {
		C.gl14Uniform4ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
	}

	glc.Uniform1fv = func(location int32, count int32, value *float32) {
		C.gl14Uniform1fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform2fv = func(location int32, count int32, value *float32) {
		C.gl14Uniform2fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform3fv = func(location int32, count int32, value *float32) {
		C.gl14Uniform3fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform4fv = func(location int32, count int32, value *float32) {
		C.gl14Uniform4fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform1iv = func(location int32, count int32, value *int32) {
		C.gl14Uniform1iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform2iv = func(location int32, count int32, value *int32) {
		C.gl14Uniform2iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform3iv = func(location int32, count int32, value *int32) {
		C.gl14Uniform3iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform4iv = func(location int32, count int32, value *int32) {
		C.gl14Uniform4iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform1uiv = func(location int32, count int32, value *uint32) {
		C.gl14Uniform1uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform2uiv = func(location int32, count int32, value *uint32) {
		C.gl14Uniform2uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform3uiv = func(location int32, count int32, value *uint32) {
		C.gl14Uniform3uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform4uiv = func(location int32, count int32, value *uint32) {
		C.gl14Uniform4uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.UseProgram = func(program uint32) {
		C.gl14UseProgram(glc.context, C.GLuint(program))
	}

	glc.ValidateProgram = func(program uint32) {
		C.gl14ValidateProgram(glc.context, C.GLuint(program))
	}

	glc.VertexAttribPointer = func(index uint32, size int32, Type uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
		C.gl14VertexAttribPointer(glc.context, C.GLuint(index), C.GLint(size), C.GLenum(Type), boolToGL(normalized), C.GLsizei(stride), pointer)
	}

	glc.VertexAttrib1f = func(index uint32, v0 float32) {
		C.gl14VertexAttrib1f(glc.context, C.GLuint(index), C.GLfloat(v0))
	}

	glc.VertexAttrib1s = func(index uint32, v0 int16) {
		C.gl14VertexAttrib1s(glc.context, C.GLuint(index), C.GLshort(v0))
	}

	glc.VertexAttrib1d = func(index uint32, v0 float64) {
		C.gl14VertexAttrib1d(glc.context, C.GLuint(index), C.GLdouble(v0))
	}

	glc.VertexAttrib2f = func(index uint32, v0, v1 float32) {
		C.gl14VertexAttrib2f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.VertexAttrib2s = func(index uint32, v0, v1 int16) {
		C.gl14VertexAttrib2s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
	}

	glc.VertexAttrib2d = func(index uint32, v0, v1 float64) {
		C.gl14VertexAttrib2d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
	}

	glc.VertexAttrib3f = func(index uint32, v0, v1, v2 float32) {
		C.gl14VertexAttrib3f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.VertexAttrib3s = func(index uint32, v0, v1, v2 int16) {
		C.gl14VertexAttrib3s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
	}

	glc.VertexAttrib3d = func(index uint32, v0, v1, v2 float64) {
		C.gl14VertexAttrib3d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.VertexAttrib4f = func(index uint32, v0, v1, v2, v3 float32) {
		C.gl14VertexAttrib4f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.VertexAttrib4s = func(index uint32, v0, v1, v2, v3 int16) {
		C.gl14VertexAttrib4s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
	}

	glc.VertexAttrib4d = func(index uint32, v0, v1, v2, v3 float64) {
		C.gl14VertexAttrib4d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
	}

	glc.VertexAttrib4Nuv = func(index uint32, v0, v1, v2, v3 uint8) {
		C.gl14VertexAttrib4Nuv(glc.context, C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
	}

	glc.VertexAttrib1fv = func(index uint32, v *float32) {
		C.gl14VertexAttrib1fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1sv = func(index uint32, v *int16) {
		C.gl14VertexAttrib1sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1dv = func(index uint32, v *float64) {
		C.gl14VertexAttrib1dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2fv = func(index uint32, v *float32) {
		C.gl14VertexAttrib2fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2sv = func(index uint32, v *int16) {
		C.gl14VertexAttrib2sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2dv = func(index uint32, v *float64) {
		C.gl14VertexAttrib2dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3fv = func(index uint32, v *float32) {
		C.gl14VertexAttrib3fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3sv = func(index uint32, v *int16) {
		C.gl14VertexAttrib3sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3dv = func(index uint32, v *float64) {
		C.gl14VertexAttrib3dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4fv = func(index uint32, v *float32) {
		C.gl14VertexAttrib4fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4sv = func(index uint32, v *int16) {
		C.gl14VertexAttrib4sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4dv = func(index uint32, v *float64) {
		C.gl14VertexAttrib4dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4iv = func(index uint32, v *int32) {
		C.gl14VertexAttrib4iv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4bv = func(index uint32, v *int8) {
		C.gl14VertexAttrib4bv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4ubv = func(index uint32, v *uint8) {
		C.gl14VertexAttrib4ubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4usv = func(index uint32, v *uint16) {
		C.gl14VertexAttrib4usv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4uiv = func(index uint32, v *uint32) {
		C.gl14VertexAttrib4uiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nbv = func(index uint32, v *int8) {
		C.gl14VertexAttrib4Nbv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nsv = func(index uint32, v *int16) {
		C.gl14VertexAttrib4Nsv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Niv = func(index uint32, v *int32) {
		C.gl14VertexAttrib4Niv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nubv = func(index uint32, v *uint8) {
		C.gl14VertexAttrib4Nubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nusv = func(index uint32, v *uint16) {
		C.gl14VertexAttrib4Nusv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nuiv = func(index uint32, v *uint32) {
		C.gl14VertexAttrib4Nuiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.UniformMatrix2fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x3fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix2x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x2fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix3x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x4fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix2x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x2fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix4x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x4fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix3x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x3fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl14UniformMatrix4x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	if !versionSupported(glc) {
		return nil
	}
	glc.queryExtensions()
	return glc
}
