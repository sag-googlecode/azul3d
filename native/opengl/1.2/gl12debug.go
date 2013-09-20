// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// +build opengl_debug

// Package 'opengl' implements OpenGL version 1.2
package opengl

// #cgo LDFLAGS: -lopengl32
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

func (glc *Context) trace(name string) {
	glc.access.Lock()
	defer glc.access.Unlock()

	glc.traceback = append(glc.traceback, name)
	l := len(glc.traceback)
	if l > 100 {
		glc.traceback = glc.traceback[l-100 : l]
	}

	if glc.inBeginEnd {
		return
	}
	err := glc.GetError()
	if err != NO_ERROR {
		fmt.Println("OpenGL call stack (last 100 - most recent first).")

		// Print stack now
		count := 0
		for i := len(glc.traceback); i > 0; i-- {
			count++
			fmt.Printf("%3.d. %s\n", count, glc.traceback[i-1])
		}

		switch err {
		case INVALID_ENUM:
			panic("GL_INVALID_ENUM: An unacceptable value was specified for an enumerated argument.")
		case INVALID_VALUE:
			panic("GL_INVALID_VALUE: A numeric argument is out of range.")
		case INVALID_OPERATION:
			panic("GL_INVALID_OPERATION: The specified operation is not allowed in the current state.")
		case INVALID_FRAMEBUFFER_OPERATION:
			panic("GL_INVALID_FRAMEBUFFER_OPERATION: The framebuffer object is not complete.")
		case OUT_OF_MEMORY:
			panic("GL_OUT_OF_MEMORY: There is not enough memory left to execute the command.")
		case STACK_UNDERFLOW:
			panic("GL_STACK_UNDERFLOW: An attempt has been made to perform an operation that would cause an internal stack to underflow.")
		case STACK_OVERFLOW:
			panic("GL_STACK_OVERFLOW: An attempt has been made to perform an operation that would cause an internal stack to overflow.")
		}
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
	MODELVIEW27_ARB                                            = 0x873B
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	INT_IMAGE_1D_EXT                                           = 0x9057
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	VERTEX_PROGRAM_NV                                          = 0x8620
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	CULL_VERTEX_EXT                                            = 0x81AA
	OPERAND3_RGB_NV                                            = 0x8593
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	PARTIAL_SUCCESS_NV                                         = 0x902E
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	TEXTURE_GEN_S                                              = 0x0C60
	LINEAR_ATTENUATION                                         = 0x1208
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	CONSTANT_ARB                                               = 0x8576
	SHADER_COMPILER                                            = 0x8DFA
	CLOSE_PATH_NV                                              = 0x00
	RGBA16                                                     = 0x805B
	RGB_SCALE                                                  = 0x8573
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	SGIX_calligraphic_fragment                                 = 1
	GL_3DC_XY_AMD                                              = 0x87FA
	FRACTIONAL_ODD                                             = 0x8E7B
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	DEBUG_OUTPUT                                               = 0x92E0
	ARRAY_STRIDE                                               = 0x92FE
	HISTOGRAM                                                  = 0x8024
	TEXTURE21                                                  = 0x84D5
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	SGIX_subsample                                             = 1
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	TEXTURE9                                                   = 0x84C9
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	SHADER_OBJECT_EXT                                          = 0x8B48
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	LIST_BASE                                                  = 0x0B32
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	TRIANGLE_FAN                                               = 0x0006
	RG8I                                                       = 0x8237
	VERTEX_PROGRAM_ARB                                         = 0x8620
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	PROGRAM_POINT_SIZE                                         = 0x8642
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	INT64_NV                                                   = 0x140E
	TEXTURE28_ARB                                              = 0x84DC
	COMBINER6_NV                                               = 0x8556
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	DYNAMIC_READ                                               = 0x88E9
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	RGBA16I                                                    = 0x8D88
	TESS_GEN_SPACING                                           = 0x8E77
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	FIELDS_NV                                                  = 0x8E27
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	LUMINANCE12_EXT                                            = 0x8041
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	DRAW_BUFFER10_ARB                                          = 0x882F
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	DRAW_BUFFER8_ARB                                           = 0x882D
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	COLOR_INDEX2_EXT                                           = 0x80E3
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	COLOR_ENCODING                                             = 0x8296
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	WAIT_FAILED_APPLE                                          = 0x911D
	SGI_texture_color_table                                    = 1
	NICEST                                                     = 0x1102
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	COLOR_LOGIC_OP                                             = 0x0BF2
	R32I                                                       = 0x8235
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	RENDER_MODE                                                = 0x0C40
	COMPRESSED_RED                                             = 0x8225
	DRAW_BUFFER8_NV                                            = 0x882D
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	VERSION_1_4                                                = 1
	SGIX_vertex_preclip                                        = 1
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	STENCIL_INDEX1_OES                                         = 0x8D46
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	DRAW_BUFFER0_ARB                                           = 0x8825
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	C3F_V3F                                                    = 0x2A24
	FENCE_CONDITION_NV                                         = 0x84F4
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	SIGNED_ALPHA8_NV                                           = 0x8706
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	DISCARD_NV                                                 = 0x8530
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	VARIABLE_A_NV                                              = 0x8523
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	DUDV_ATI                                                   = 0x8779
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	REG_4_ATI                                                  = 0x8925
	COLOR_ATTACHMENT2                                          = 0x8CE2
	CURRENT_RASTER_INDEX                                       = 0x0B05
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	RGBA_FLOAT32_APPLE                                         = 0x8814
	FALSE                                                      = 0
	ZOOM_X                                                     = 0x0D16
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	MINOR_VERSION                                              = 0x821C
	CULL_FRAGMENT_NV                                           = 0x86E7
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	CURRENT_BIT                                                = 0x00000001
	LUMINANCE                                                  = 0x1909
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	MODELVIEW26_ARB                                            = 0x873A
	MOV_ATI                                                    = 0x8961
	RGBA8UI_EXT                                                = 0x8D7C
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	FACTOR_MAX_AMD                                             = 0x901D
	GL_3D                                                      = 0x0601
	R3_G3_B2                                                   = 0x2A10
	INTENSITY16_EXT                                            = 0x804D
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	SGIS_detail_texture                                        = 1
	AUX0                                                       = 0x0409
	TEXTURE8_ARB                                               = 0x84C8
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	COLOR_TABLE_FORMAT                                         = 0x80D8
	PHONG_HINT_WIN                                             = 0x80EB
	MAX_VIEWPORTS                                              = 0x825B
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	RGBA8I                                                     = 0x8D8E
	UNPACK_SKIP_IMAGES                                         = 0x806D
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	LOCATION_INDEX                                             = 0x930F
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	LIGHT0                                                     = 0x4000
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	MODELVIEW14_ARB                                            = 0x872E
	MODELVIEW25_ARB                                            = 0x8739
	DOT2_ADD_ATI                                               = 0x896C
	STENCIL_WRITEMASK                                          = 0x0B98
	SAMPLE_COVERAGE                                            = 0x80A0
	COLOR_BUFFER_BIT                                           = 0x00004000
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	TEXTURE_VIEW                                               = 0x82B5
	DOT3_RGBA                                                  = 0x86AF
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	DOUBLE_MAT4                                                = 0x8F48
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	TEXTURE11_ARB                                              = 0x84CB
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	OPERAND2_ALPHA_ARB                                         = 0x859A
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	DRAW_BUFFER11_ATI                                          = 0x8830
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	HISTOGRAM_EXT                                              = 0x8024
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	MATRIX12_ARB                                               = 0x88CC
	OBJECT_TYPE_ARB                                            = 0x8B4E
	BLUE_INTEGER                                               = 0x8D96
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	WEIGHT_ARRAY_OES                                           = 0x86AD
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	STATIC_DRAW                                                = 0x88E4
	FLOAT_VEC4                                                 = 0x8B52
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	MUL_ATI                                                    = 0x8964
	FUNC_SUBTRACT_EXT                                          = 0x800A
	DRAW_BUFFER2                                               = 0x8827
	PIXEL_COUNT_NV                                             = 0x8866
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	RGB8I_EXT                                                  = 0x8D8F
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	PACK_SKIP_PIXELS                                           = 0x0D04
	MAP1_NORMAL                                                = 0x0D92
	MIRRORED_REPEAT_IBM                                        = 0x8370
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	EXT_blend_subtract                                         = 1
	NAND                                                       = 0x150E
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	TEXTURE20                                                  = 0x84D4
	SAMPLES_3DFX                                               = 0x86B4
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	ACTIVE_RESOURCES                                           = 0x92F5
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	COMBINER_BIAS_NV                                           = 0x8549
	OP_MADD_EXT                                                = 0x8788
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	CLIP_PLANE5                                                = 0x3005
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	VIEWPORT_BIT                                               = 0x00000800
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	FLOAT_RG_NV                                                = 0x8881
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	FLOAT16_VEC4_NV                                            = 0x8FFB
	BLUE_SCALE                                                 = 0x0D1A
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	BUFFER_DATA_SIZE                                           = 0x9303
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	LUMINANCE8_EXT                                             = 0x8040
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	VIEWPORT                                                   = 0x0BA2
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	BGR_INTEGER_EXT                                            = 0x8D9A
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	GL_2D                                                      = 0x0600
	TEXTURE_BINDING_2D                                         = 0x8069
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	ASYNC_MARKER_SGIX                                          = 0x8329
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	SGIS_texture_select                                        = 1
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	RG16UI                                                     = 0x823A
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	CONTEXT_PROFILE_MASK                                       = 0x9126
	COLOR_MATERIAL_FACE                                        = 0x0B55
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	BUFFER_SIZE                                                = 0x8764
	REG_16_ATI                                                 = 0x8931
	R11F_G11F_B10F                                             = 0x8C3A
	ALPHA8I_EXT                                                = 0x8D90
	PATH_FORMAT_PS_NV                                          = 0x9071
	SHORT                                                      = 0x1402
	TEXTURE6                                                   = 0x84C6
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	MATRIX_EXT                                                 = 0x87C0
	BGR_INTEGER                                                = 0x8D9A
	TRIANGLE_STRIP                                             = 0x0005
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	OPERAND1_ALPHA                                             = 0x8599
	ADD_SIGNED_ARB                                             = 0x8574
	UTF8_NV                                                    = 0x909A
	DEBUG_SEVERITY_LOW                                         = 0x9148
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	STENCIL_BITS                                               = 0x0D57
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	SIGNED_IDENTITY_NV                                         = 0x853C
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	MAP2_COLOR_4                                               = 0x0DB0
	AND_INVERTED                                               = 0x1504
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	CAVEAT_SUPPORT                                             = 0x82B8
	DECR_WRAP                                                  = 0x8508
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	MATRIX7_NV                                                 = 0x8637
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	SGIX_list_priority                                         = 1
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	RESCALE_NORMAL_EXT                                         = 0x803A
	COMBINER_INPUT_NV                                          = 0x8542
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	HISTOGRAM_FORMAT                                           = 0x8027
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	OP_SUB_EXT                                                 = 0x8796
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	FRAMEBUFFER_OES                                            = 0x8D40
	VERTEX_SHADER_BIT                                          = 0x00000001
	TEXTURE_MATRIX                                             = 0x0BA8
	CONVOLUTION_1D                                             = 0x8010
	REPLACE_EXT                                                = 0x8062
	DRAW_BUFFER4_ATI                                           = 0x8829
	EXT_copy_texture                                           = 1
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	DRAW_BUFFER1_ATI                                           = 0x8826
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	SGIX_async_histogram                                       = 1
	ETC1_SRGB8_NV                                              = 0x88EE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	T2F_C4UB_V3F                                               = 0x2A29
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	DUAL_ALPHA8_SGIS                                           = 0x8111
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	DRAW_BUFFER1_NV                                            = 0x8826
	INTENSITY8I_EXT                                            = 0x8D91
	RED_SNORM                                                  = 0x8F90
	QUERY_OBJECT_AMD                                           = 0x9153
	LIGHT2                                                     = 0x4002
	ZOOM_Y                                                     = 0x0D17
	TEXTURE_BINDING_3D                                         = 0x806A
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	PRIMITIVE_ID_NV                                            = 0x8C7C
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	RGBA_MODE                                                  = 0x0C31
	NEGATIVE_ONE_EXT                                           = 0x87DF
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	BGR_EXT                                                    = 0x80E0
	OBJECT_POINT_SGIS                                          = 0x81F5
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	REG_25_ATI                                                 = 0x893A
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	EVAL_2D_NV                                                 = 0x86C0
	ALPHA32F_ARB                                               = 0x8816
	BOOL_VEC2                                                  = 0x8B57
	SLUMINANCE8_EXT                                            = 0x8C47
	RASTERIZER_DISCARD                                         = 0x8C89
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	GL_4X_BIT_ATI                                              = 0x00000002
	UNDEFINED_APPLE                                            = 0x8A1C
	FLOAT_MAT3x2                                               = 0x8B67
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	HISTOGRAM_SINK                                             = 0x802D
	MATRIX3_NV                                                 = 0x8633
	INTERPOLATE                                                = 0x8575
	OPERAND1_ALPHA_EXT                                         = 0x8599
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	DRAW_BUFFER6                                               = 0x882B
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	SLUMINANCE_NV                                              = 0x8C46
	FRAMEBUFFER_EXT                                            = 0x8D40
	AFFINE_2D_NV                                               = 0x9092
	BUFFER_MAP_OFFSET                                          = 0x9121
	STEREO                                                     = 0x0C33
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	R16F                                                       = 0x822D
	DRAW_BUFFER13_ARB                                          = 0x8832
	DEPTH_RANGE                                                = 0x0B70
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	LIST_MODE                                                  = 0x0B30
	MIRRORED_REPEAT_ARB                                        = 0x8370
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	GREEN_BIT_ATI                                              = 0x00000002
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	RG_SNORM                                                   = 0x8F91
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	FOG_FUNC_SGIS                                              = 0x812A
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	OPERAND0_RGB                                               = 0x8590
	RGB10_A2UI                                                 = 0x906F
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	COMPRESSED_R11_EAC                                         = 0x9270
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	TEXTURE0_ARB                                               = 0x84C0
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	FUNC_SUBTRACT                                              = 0x800A
	RGBA4_S3TC                                                 = 0x83A3
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	COLOR_ATTACHMENT7                                          = 0x8CE7
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	READ_BUFFER_EXT                                            = 0x0C02
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	REG_28_ATI                                                 = 0x893D
	FOG_COORDINATE_EXT                                         = 0x8451
	MATRIX19_ARB                                               = 0x88D3
	RGB_422_APPLE                                              = 0x8A1F
	SAMPLER_3D                                                 = 0x8B5F
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	SOURCE2_ALPHA_EXT                                          = 0x858A
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	READ_FRAMEBUFFER                                           = 0x8CA8
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	LUMINANCE8_SNORM                                           = 0x9015
	EDGE_FLAG                                                  = 0x0B43
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	MATRIX21_ARB                                               = 0x88D5
	HALF_FLOAT_NV                                              = 0x140B
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	SIGNED_RGB_NV                                              = 0x86FE
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	MATRIX26_ARB                                               = 0x88DA
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	MATRIX7_ARB                                                = 0x88C7
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	RGBA16F_EXT                                                = 0x881A
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	RGBA_INTEGER                                               = 0x8D99
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	COMBINE_ARB                                                = 0x8570
	MODELVIEW28_ARB                                            = 0x873C
	DRAW_BUFFER2_NV                                            = 0x8827
	PALETTE4_RGB8_OES                                          = 0x8B90
	RGB_INTEGER                                                = 0x8D98
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	COMBINER3_NV                                               = 0x8553
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	DOT_PRODUCT_NV                                             = 0x86EC
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	DU8DV8_ATI                                                 = 0x877A
	DRAW_BUFFER9_NV                                            = 0x882E
	CON_5_ATI                                                  = 0x8946
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	LINE_LOOP                                                  = 0x0002
	CLAMP                                                      = 0x2900
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	NORMAL_MAP_OES                                             = 0x8511
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	EXPAND_NEGATE_NV                                           = 0x8539
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	FILE_NAME_NV                                               = 0x9074
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	ADD                                                        = 0x0104
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	SRGB8                                                      = 0x8C41
	AUX2                                                       = 0x040B
	YCRCB_422_SGIX                                             = 0x81BB
	COLOR_SUM                                                  = 0x8458
	DELETE_STATUS                                              = 0x8B80
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	STENCIL_ATTACHMENT                                         = 0x8D20
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	LOAD                                                       = 0x0101
	EXTENSIONS                                                 = 0x1F03
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	INTENSITY16UI_EXT                                          = 0x8D79
	TRANSLATE_Y_NV                                             = 0x908F
	QUERY_BUFFER_AMD                                           = 0x9192
	QUADS                                                      = 0x0007
	DEPTH_FUNC                                                 = 0x0B74
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	RG16F                                                      = 0x822F
	SRGB_WRITE                                                 = 0x8298
	LUMINANCE8                                                 = 0x8040
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	RGB_S3TC                                                   = 0x83A0
	CON_22_ATI                                                 = 0x8957
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	SPOT_EXPONENT                                              = 0x1205
	CON_15_ATI                                                 = 0x8950
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	TEXTURE0                                                   = 0x84C0
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	EMBOSS_CONSTANT_NV                                         = 0x855E
	WRITE_ONLY_OES                                             = 0x88B9
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	RGB4_S3TC                                                  = 0x83A1
	DEPTH_CLAMP                                                = 0x864F
	INVALID_OPERATION                                          = 0x0502
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	DUAL_ALPHA16_SGIS                                          = 0x8113
	YCRCB_444_SGIX                                             = 0x81BC
	DEPTH_STENCIL                                              = 0x84F9
	E_TIMES_F_NV                                               = 0x8531
	N3F_V3F                                                    = 0x2A25
	FUNC_SUBTRACT_OES                                          = 0x800A
	SAMPLES_PASSED_ARB                                         = 0x8914
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	PATH_DASH_CAPS_NV                                          = 0x907B
	EQUIV                                                      = 0x1509
	REG_6_ATI                                                  = 0x8927
	DOT3_ATI                                                   = 0x8966
	FRONT_AND_BACK                                             = 0x0408
	ALPHA_SCALE                                                = 0x0D1C
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	TEXTURE30                                                  = 0x84DE
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	STREAM_DRAW_ARB                                            = 0x88E0
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	RGB5_A1_EXT                                                = 0x8057
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	YCBCR_422_APPLE                                            = 0x85B9
	COMPUTE_SHADER                                             = 0x91B9
	LINES_ADJACENCY_ARB                                        = 0x000A
	OR_REVERSE                                                 = 0x150B
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	TRACK_MATRIX_NV                                            = 0x8648
	PROGRAM_STRING_ARB                                         = 0x8628
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	RG16F_EXT                                                  = 0x822F
	SHADER_IMAGE_LOAD                                          = 0x82A4
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	OPERAND2_RGB                                               = 0x8592
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	STATIC_DRAW_ARB                                            = 0x88E4
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	NORMAL_ARRAY                                               = 0x8075
	AND_REVERSE                                                = 0x1502
	PROXY_TEXTURE_1D                                           = 0x8063
	PROGRAM_PIPELINE                                           = 0x82E4
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	BLEND_SRC_RGB                                              = 0x80C9
	FILTER4_SGIS                                               = 0x8146
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	MOVE_TO_RESETS_NV                                          = 0x90B5
	ARC_TO_NV                                                  = 0xFE
	EXT_abgr                                                   = 1
	LINE_STIPPLE                                               = 0x0B24
	PROGRAM_LENGTH_NV                                          = 0x8627
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	INTENSITY8_EXT                                             = 0x804B
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	MIRROR_CLAMP_EXT                                           = 0x8742
	RELATIVE_ARC_TO_NV                                         = 0xFF
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	MULT                                                       = 0x0103
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	SLIM10U_SGIX                                               = 0x831E
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	STENCIL_INDEX4_OES                                         = 0x8D47
	POLYGON_OFFSET_POINT                                       = 0x2A01
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	TEXTURE17_ARB                                              = 0x84D1
	SOURCE2_RGB_EXT                                            = 0x8582
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	CND_ATI                                                    = 0x896A
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	SIGNALED                                                   = 0x9119
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	POINT_SIZE_MAX                                             = 0x8127
	DS_BIAS_NV                                                 = 0x8716
	LUMINANCE32F_ARB                                           = 0x8818
	STENCIL_INDEX8                                             = 0x8D48
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	NAME_LENGTH                                                = 0x92F9
	RGB8                                                       = 0x8051
	CONVOLUTION_1D_EXT                                         = 0x8010
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	INTENSITY                                                  = 0x8049
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	TESS_CONTROL_TEXTURE                                       = 0x829C
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	SOURCE1_ALPHA                                              = 0x8589
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	MATRIX9_ARB                                                = 0x88C9
	CON_20_ATI                                                 = 0x8955
	BOOL_VEC4                                                  = 0x8B59
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	RGBA16UI_EXT                                               = 0x8D76
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	YCBYCR8_422_NV                                             = 0x9031
	SYNC_STATUS_APPLE                                          = 0x9114
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	SLUMINANCE8                                                = 0x8C47
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	MODELVIEW                                                  = 0x1700
	INDEX_TEST_REF_EXT                                         = 0x81B7
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	OPERAND0_ALPHA_ARB                                         = 0x8598
	CURRENT_ATTRIB_NV                                          = 0x8626
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	REFLECTION_MAP                                             = 0x8512
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	DIFFUSE                                                    = 0x1201
	PACK_CMYK_HINT_EXT                                         = 0x800E
	STREAM_COPY                                                = 0x88E2
	TEXTURE_RED_TYPE                                           = 0x8C10
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	MAX_TEXTURE_UNITS                                          = 0x84E2
	OP_DOT4_EXT                                                = 0x8785
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	EXP2                                                       = 0x0801
	SOURCE0_RGB_EXT                                            = 0x8580
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	ZERO                                                       = 0
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	RED                                                        = 0x1903
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	ALPHA16F_ARB                                               = 0x881C
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	ACTIVE_PROGRAM                                             = 0x8259
	TEXTURE12                                                  = 0x84CC
	TESS_GEN_POINT_MODE                                        = 0x8E79
	SLICE_ACCUM_SUN                                            = 0x85CC
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	ALPHA_SNORM                                                = 0x9010
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	SPOT_CUTOFF                                                = 0x1206
	IUI_N3F_V3F_EXT                                            = 0x81B0
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	NEXT_BUFFER_NV                                             = -2
	UNSIGNALED                                                 = 0x9118
	CURRENT_RASTER_POSITION                                    = 0x0B07
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	COMPRESSED_LUMINANCE                                       = 0x84EA
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	TEXTURE15_ARB                                              = 0x84CF
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	FOG_DENSITY                                                = 0x0B62
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	SAMPLE_MASK_NV                                             = 0x8E51
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	DRAW_BUFFER0                                               = 0x8825
	DEPTH_TEXTURE_MODE                                         = 0x884B
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	TEXTURE26                                                  = 0x84DA
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	EXT_texture                                                = 1
	LUMINANCE12_ALPHA4                                         = 0x8046
	LUMINANCE16_ALPHA16                                        = 0x8048
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	SCALE_BY_TWO_NV                                            = 0x853E
	OPERAND1_RGB_ARB                                           = 0x8591
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	TEXTURE_MIN_LOD                                            = 0x813A
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	SHADER_BINARY_VIV                                          = 0x8FC4
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	RGBA_FLOAT32_ATI                                           = 0x8814
	PALETTE8_RGBA4_OES                                         = 0x8B98
	BGRA                                                       = 0x80E1
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	CONSTANT_BORDER                                            = 0x8151
	INCR_WRAP_EXT                                              = 0x8507
	INVERSE_NV                                                 = 0x862B
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	FRONT                                                      = 0x0404
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	BLEND_SRC_ALPHA                                            = 0x80CB
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	COMBINER_MUX_SUM_NV                                        = 0x8547
	TRACE_NAME_MESA                                            = 0x8756
	EYE_LINEAR                                                 = 0x2400
	OP_MUL_EXT                                                 = 0x8786
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	TRIANGLES_ADJACENCY                                        = 0x000C
	LINEAR_DETAIL_SGIS                                         = 0x8097
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	ACTIVE_UNIFORMS                                            = 0x8B86
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	DRAW_BUFFER4                                               = 0x8829
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	LEFT                                                       = 0x0406
	PROGRAM_STRING_NV                                          = 0x8628
	DOT3_RGB_EXT                                               = 0x8740
	CON_18_ATI                                                 = 0x8953
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	INT64_VEC4_NV                                              = 0x8FEB
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	CURRENT_QUERY                                              = 0x8865
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	RGBA2                                                      = 0x8055
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	TEXTURE_LOD_BIAS                                           = 0x8501
	SOURCE0_ALPHA_ARB                                          = 0x8588
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	WAIT_FAILED                                                = 0x911D
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	RGB_FLOAT16_ATI                                            = 0x881B
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	DEBUG_TYPE_OTHER                                           = 0x8251
	NORMAL_MAP_NV                                              = 0x8511
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	MAX_IMAGE_SAMPLES                                          = 0x906D
	MAP1_COLOR_4                                               = 0x0D90
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	ALREADY_SIGNALED                                           = 0x911A
	QUAD_STRIP                                                 = 0x0008
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	STENCIL_INDEX                                              = 0x1901
	SRC1_COLOR                                                 = 0x88F9
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	FRAGMENT_DEPTH                                             = 0x8452
	SHADER_OPERATION_NV                                        = 0x86DF
	VIBRANCE_SCALE_NV                                          = 0x8713
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	FLAT                                                       = 0x1D00
	GL_4PASS_1_EXT                                             = 0x80A5
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	LIST_PRIORITY_SGIX                                         = 0x8182
	GL_2_BYTES                                                 = 0x1407
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	DRAW_BUFFER7_NV                                            = 0x882C
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	RGBA32I_EXT                                                = 0x8D82
	POLYGON_TOKEN                                              = 0x0703
	RETAINED_APPLE                                             = 0x8A1B
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	QUAD_ALPHA8_SGIS                                           = 0x811F
	SINGLE_COLOR_EXT                                           = 0x81F9
	R8                                                         = 0x8229
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	FOG_COORDINATE                                             = 0x8451
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	INDEX_ARRAY_POINTER                                        = 0x8091
	RED_BITS                                                   = 0x0D52
	MATRIX3_ARB                                                = 0x88C3
	SAMPLE_POSITION_NV                                         = 0x8E50
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	COLOR_ARRAY                                                = 0x8076
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	CLIP_DISTANCE0                                             = 0x3000
	RG8                                                        = 0x822B
	OPERAND1_ALPHA_ARB                                         = 0x8599
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	SOURCE0_ALPHA                                              = 0x8588
	MODELVIEW21_ARB                                            = 0x8735
	DYNAMIC_COPY                                               = 0x88EA
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	DEPTH_BUFFER_BIT                                           = 0x00000100
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	COMPRESSED_RG11_EAC                                        = 0x9272
	LUMINANCE4                                                 = 0x803F
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	TEXTURE30_ARB                                              = 0x84DE
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	AND                                                        = 0x1501
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	INTENSITY12                                                = 0x804C
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	DRAW_BUFFER15_ATI                                          = 0x8834
	QUERY_COUNTER_BITS                                         = 0x8864
	SAMPLER_1D_ARB                                             = 0x8B5D
	HALF_FLOAT_ARB                                             = 0x140B
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	FLOAT                                                      = 0x1406
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	PRIMITIVE_RESTART_NV                                       = 0x8558
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	COLOR_ATTACHMENT10                                         = 0x8CEA
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	TEXTURE_WIDTH                                              = 0x1000
	INTENSITY16F_ARB                                           = 0x881D
	QUERY_RESULT_ARB                                           = 0x8866
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	DRAW_BUFFER10_NV                                           = 0x882F
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	RGB_SCALE_ARB                                              = 0x8573
	ATC_RGB_AMD                                                = 0x8C92
	TYPE                                                       = 0x92FA
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	MODELVIEW20_ARB                                            = 0x8734
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	MAX_CLIP_DISTANCES                                         = 0x0D32
	SPRITE_MODE_SGIX                                           = 0x8149
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	RGBA8I_EXT                                                 = 0x8D8E
	RG8_SNORM                                                  = 0x8F95
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	READ_WRITE                                                 = 0x88BA
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	MAX_IMAGE_UNITS                                            = 0x8F38
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	ISOLINES                                                   = 0x8E7A
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	BLEND_DST_RGB_EXT                                          = 0x80C8
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	NOTEQUAL                                                   = 0x0205
	RGB5                                                       = 0x8050
	CONVOLUTION_FORMAT                                         = 0x8017
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	MODELVIEW6_ARB                                             = 0x8726
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	CONSTANT_COLOR                                             = 0x8001
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	TEXTURE_ENV_COLOR                                          = 0x2201
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	SOURCE3_RGB_NV                                             = 0x8583
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	CON_31_ATI                                                 = 0x8960
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	IMAGE_1D_EXT                                               = 0x904C
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	SAMPLER_2D_RECT                                            = 0x8B63
	DOUBLE_MAT3                                                = 0x8F47
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	SKIP_COMPONENTS2_NV                                        = -5
	LIST_INDEX                                                 = 0x0B33
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	STENCIL_INDEX1                                             = 0x8D46
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	LUMINANCE8_ALPHA8                                          = 0x8045
	SRGB_DECODE_ARB                                            = 0x8299
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	INTENSITY32UI_EXT                                          = 0x8D73
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	MULTISAMPLE_EXT                                            = 0x809D
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	TEXTURE_BUFFER                                             = 0x8C2A
	OP_MAX_EXT                                                 = 0x878A
	DRAW_BUFFER13_NV                                           = 0x8832
	SYNC_CONDITION_APPLE                                       = 0x9113
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	VARIANT_EXT                                                = 0x87C1
	GL_3DC_X_AMD                                               = 0x87F9
	MATRIX2_ARB                                                = 0x88C2
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	POINT_SPRITE_NV                                            = 0x8861
	PACK_RESAMPLE_OML                                          = 0x8984
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	REG_30_ATI                                                 = 0x893F
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	AMBIENT                                                    = 0x1200
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	GL_2PASS_0_SGIS                                            = 0x80A2
	SAMPLE_MASK_EXT                                            = 0x80A0
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	CON_11_ATI                                                 = 0x894C
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	SWIZZLE_STQ_ATI                                            = 0x8977
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	PRESENT_TIME_NV                                            = 0x8E2A
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	PROGRAM_OUTPUT                                             = 0x92E4
	MODELVIEW18_ARB                                            = 0x8732
	EXT_cmyka                                                  = 1
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	TRUE                                                       = 1
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	DRAW_BUFFER2_ARB                                           = 0x8827
	REG_14_ATI                                                 = 0x892F
	CON_2_ATI                                                  = 0x8943
	CND0_ATI                                                   = 0x896B
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	TEXTURE7                                                   = 0x84C7
	TEXTURE2_ARB                                               = 0x84C2
	TEXTURE21_ARB                                              = 0x84D5
	TEXTURE_MAX_LOD                                            = 0x813B
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	S                                                          = 0x2000
	ATTENUATION_EXT                                            = 0x834D
	DOUBLE_MAT3x4                                              = 0x8F4C
	HISTOGRAM_WIDTH                                            = 0x8026
	POINT_SIZE_MIN                                             = 0x8126
	UNSIGNED_INT_24_8                                          = 0x84FA
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	SRC_ALPHA                                                  = 0x0302
	LUMINANCE_ALPHA                                            = 0x190A
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	VARIABLE_E_NV                                              = 0x8527
	OPERAND2_ALPHA_EXT                                         = 0x859A
	ACCUM                                                      = 0x0100
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	LUMINANCE8UI_EXT                                           = 0x8D80
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	LIGHTING                                                   = 0x0B50
	POINT_SIZE_MAX_EXT                                         = 0x8127
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	OBJECT_LINE_SGIS                                           = 0x81F7
	DEBUG_SOURCE_API                                           = 0x8246
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	RGB16I_EXT                                                 = 0x8D89
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	LIGHTING_BIT                                               = 0x00000040
	READ_BUFFER_NV                                             = 0x0C02
	RGB8I                                                      = 0x8D8F
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	SGIX_interlace                                             = 1
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	COPY_READ_BUFFER                                           = 0x8F36
	INT_IMAGE_1D                                               = 0x9057
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	GL_3D_COLOR                                                = 0x0602
	DRAW_BUFFER10_ATI                                          = 0x882F
	INDEX_BITS                                                 = 0x0D51
	CONSTANT_ALPHA_EXT                                         = 0x8003
	REG_7_ATI                                                  = 0x8928
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	PACK_LSB_FIRST                                             = 0x0D01
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	MULTISAMPLE_3DFX                                           = 0x86B2
	VERSION_1_3                                                = 1
	SGIX_blend_alpha_minmax                                    = 1
	BLUE_BITS                                                  = 0x0D54
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	FLOAT_VEC4_ARB                                             = 0x8B52
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	INTERLACE_SGIX                                             = 0x8094
	R16UI                                                      = 0x8234
	FILL                                                       = 0x1B02
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	COLOR                                                      = 0x1800
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	INTENSITY8_SNORM                                           = 0x9017
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	R8I                                                        = 0x8231
	REG_11_ATI                                                 = 0x892C
	FLOAT_MAT3x4                                               = 0x8B68
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	RGB2_EXT                                                   = 0x804E
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	MAX_SHININESS_NV                                           = 0x8504
	TRIANGLE_MESH_SUN                                          = 0x8615
	FLOAT_MAT4_ARB                                             = 0x8B5C
	IMAGE_2D                                                   = 0x904D
	MAX_NAME_LENGTH                                            = 0x92F6
	COMBINER0_NV                                               = 0x8550
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	TRANSLATE_X_NV                                             = 0x908E
	MAX_EVAL_ORDER                                             = 0x0D30
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	SYNC_STATUS                                                = 0x9114
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	RED_BIAS                                                   = 0x0D15
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	RGB16F                                                     = 0x881B
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	DOUBLE_VEC3                                                = 0x8FFD
	TANGENT_ARRAY_EXT                                          = 0x8439
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	DT_SCALE_NV                                                = 0x8711
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	CLIP_DISTANCE2                                             = 0x3002
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	INT64_VEC3_NV                                              = 0x8FEA
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	RGB9_E5_EXT                                                = 0x8C3D
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	RGB32UI                                                    = 0x8D71
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	GREEN_BIAS                                                 = 0x0D19
	GL_4_BYTES                                                 = 0x1409
	PROXY_HISTOGRAM                                            = 0x8025
	POINT_SIZE_MIN_EXT                                         = 0x8126
	FOG_OFFSET_SGIX                                            = 0x8198
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	MODELVIEW7_ARB                                             = 0x8727
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	VERTEX_ARRAY_TYPE                                          = 0x807B
	TEXTURE_BORDER                                             = 0x1005
	CLIP_DISTANCE4                                             = 0x3004
	RGB4                                                       = 0x804F
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	MATRIX29_ARB                                               = 0x88DD
	TABLE_TOO_LARGE                                            = 0x8031
	SAMPLES_ARB                                                = 0x80A9
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	STREAM_READ                                                = 0x88E1
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	COLOR_TABLE_SCALE                                          = 0x80D6
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	OPERAND2_RGB_EXT                                           = 0x8592
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	SGIS_texture4D                                             = 1
	BLEND_DST                                                  = 0x0BE0
	INCR                                                       = 0x1E02
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	DRAW_BUFFER3                                               = 0x8828
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	SRC0_RGB                                                   = 0x8580
	PACK_SWAP_BYTES                                            = 0x0D00
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	MINMAX_SINK                                                = 0x8030
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	SAMPLES_PASSED                                             = 0x8914
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	REPEAT                                                     = 0x2901
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	TIME_ELAPSED                                               = 0x88BF
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	OP_SET_GE_EXT                                              = 0x878C
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	EXT_texture3D                                              = 1
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	OPERAND1_RGB_EXT                                           = 0x8591
	VERSION_3_2                                                = 1
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	REG_1_ATI                                                  = 0x8922
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	SRC1_ALPHA                                                 = 0x8589
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	FLOAT_VEC2_ARB                                             = 0x8B50
	LINK_STATUS                                                = 0x8B82
	NORMAL_ARRAY_POINTER                                       = 0x808F
	FEEDBACK                                                   = 0x1C01
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	FLOAT_VEC3                                                 = 0x8B51
	RGB12                                                      = 0x8053
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	INTENSITY32F_ARB                                           = 0x8817
	ADD_ATI                                                    = 0x8963
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	DRAW_BUFFER11                                              = 0x8830
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	MODELVIEW2_ARB                                             = 0x8722
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	T2F_N3F_V3F                                                = 0x2A2B
	CLIP_DISTANCE6                                             = 0x3006
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	SPARE1_NV                                                  = 0x852F
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	LO_SCALE_NV                                                = 0x870F
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	RGBA8_SNORM                                                = 0x8F97
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	C4UB_V3F                                                   = 0x2A23
	BGR                                                        = 0x80E0
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	ADD_SIGNED_EXT                                             = 0x8574
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	COLOR_ATTACHMENT13                                         = 0x8CED
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	LINE_TOKEN                                                 = 0x0702
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	ALPHA                                                      = 0x1906
	RGB10_EXT                                                  = 0x8052
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	INT_IMAGE_2D                                               = 0x9058
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	MODELVIEW30_ARB                                            = 0x873E
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	OP_RECIP_SQRT_EXT                                          = 0x8795
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	TEXTURE13                                                  = 0x84CD
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	SEPARATE_ATTRIBS                                           = 0x8C8D
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	QUERY                                                      = 0x82E3
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	GREEN_SCALE                                                = 0x0D18
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	RGB9_E5                                                    = 0x8C3D
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	TEXTURE22_ARB                                              = 0x84D6
	SOURCE1_RGB                                                = 0x8581
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	NEGATIVE_Y_EXT                                             = 0x87DA
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	BOUNDING_BOX_NV                                            = 0x908D
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	RG                                                         = 0x8227
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	MATRIX0_NV                                                 = 0x8630
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	HI_BIAS_NV                                                 = 0x8714
	DRAW_BUFFER9_ARB                                           = 0x882E
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	AUX_BUFFERS                                                = 0x0C00
	LUMINANCE12_ALPHA12                                        = 0x8047
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	CURRENT_NORMAL                                             = 0x0B02
	VERTEX_STREAM7_ATI                                         = 0x8773
	MATRIX24_ARB                                               = 0x88D8
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	EXT_blend_logic_op                                         = 1
	T                                                          = 0x2001
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	RGBA8                                                      = 0x8058
	LIGHT7                                                     = 0x4007
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	MAX_SAMPLES                                                = 0x8D57
	ALPHA32UI_EXT                                              = 0x8D72
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	INTENSITY8                                                 = 0x804B
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	SGIX_scalebias_hint                                        = 1
	MODELVIEW4_ARB                                             = 0x8724
	NEGATIVE_X_EXT                                             = 0x87D9
	CLAMP_VERTEX_COLOR                                         = 0x891A
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SGIS_texture_border_clamp                                  = 1
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	V3F                                                        = 0x2A21
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	ALPHA8UI_EXT                                               = 0x8D7E
	MODELVIEW1_ARB                                             = 0x850A
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	AUX3                                                       = 0x040C
	MATRIX_PALETTE_OES                                         = 0x8840
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	MATRIX28_ARB                                               = 0x88DC
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	MAP_WRITE_BIT                                              = 0x0002
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	INTENSITY12_EXT                                            = 0x804C
	COLOR_MATRIX                                               = 0x80B1
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	SYNC_CONDITION                                             = 0x9113
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	REG_8_ATI                                                  = 0x8929
	VERTEX_SHADER                                              = 0x8B31
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	DST_ALPHA                                                  = 0x0304
	DRAW_BUFFER7                                               = 0x882C
	ARRAY_BUFFER                                               = 0x8892
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	UNIFORM_BUFFER                                             = 0x8A11
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	FENCE_STATUS_NV                                            = 0x84F3
	MODELVIEW24_ARB                                            = 0x8738
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	GL_1PASS_EXT                                               = 0x80A1
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	IGNORE_BORDER_HP                                           = 0x8150
	YCRCBA_SGIX                                                = 0x8319
	IMAGE_BUFFER                                               = 0x9051
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	COLOR_SUM_EXT                                              = 0x8458
	DRAW_BUFFER14                                              = 0x8833
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	WRITE_DISCARD_NV                                           = 0x88BE
	STATIC_READ_ARB                                            = 0x88E5
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	VERSION_3_0                                                = 1
	POINT                                                      = 0x1B00
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	PREVIOUS_ARB                                               = 0x8578
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	DEPTH                                                      = 0x1801
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	TEXTURE_3D_EXT                                             = 0x806F
	SAMPLES_SGIS                                               = 0x80A9
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	VERTEX_ARRAY_EXT                                           = 0x8074
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	SUBTRACT_ARB                                               = 0x84E7
	DRAW_BUFFER8_ATI                                           = 0x882D
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	RGB_SCALE_EXT                                              = 0x8573
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	CW                                                         = 0x0900
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	SAMPLER                                                    = 0x82E6
	TEXTURE26_ARB                                              = 0x84DA
	DEPTH_STENCIL_NV                                           = 0x84F9
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	REG_22_ATI                                                 = 0x8937
	TIMESTAMP                                                  = 0x8E28
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	CON_25_ATI                                                 = 0x895A
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	LUMINANCE4_ALPHA4                                          = 0x8043
	BLEND_EQUATION_EXT                                         = 0x8009
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	RGB16UI                                                    = 0x8D77
	GREEN_INTEGER_EXT                                          = 0x8D95
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	MIPMAP                                                     = 0x8293
	RIGHT                                                      = 0x0407
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	VERTEX_STREAM4_ATI                                         = 0x8770
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	GPU_ADDRESS_NV                                             = 0x8F34
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	REFLECTION_MAP_EXT                                         = 0x8512
	ONE_EXT                                                    = 0x87DE
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	SKIP_COMPONENTS4_NV                                        = -3
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	DEPTH_COMPONENTS                                           = 0x8284
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	TEXTURE29                                                  = 0x84DD
	LUMINANCE32I_EXT                                           = 0x8D86
	IMAGE_1D_ARRAY                                             = 0x9052
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	SHADER_IMAGE_STORE                                         = 0x82A5
	VIEW_CLASS_24_BITS                                         = 0x82C9
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	TEXTURE_COMPARE_MODE                                       = 0x884C
	INT_2_10_10_10_REV                                         = 0x8D9F
	TEXTURE_GEN_R                                              = 0x0C62
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	DISCRETE_AMD                                               = 0x9006
	INT_IMAGE_BUFFER                                           = 0x905C
	UNPACK_ALIGNMENT                                           = 0x0CF5
	RGBA32F_ARB                                                = 0x8814
	DRAW_BUFFER5_ARB                                           = 0x882A
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	DYNAMIC_DRAW                                               = 0x88E8
	REG_17_ATI                                                 = 0x8932
	STENCIL_INDEX4_EXT                                         = 0x8D47
	PRESENT_DURATION_NV                                        = 0x8E2B
	RELATIVE_MOVE_TO_NV                                        = 0x03
	FIXED_OES                                                  = 0x140C
	ALPHA4                                                     = 0x803B
	GL_4PASS_0_SGIS                                            = 0x80A4
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	TEXTURE22                                                  = 0x84D6
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	PACK_INVERT_MESA                                           = 0x8758
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	DEPTH_COMPONENT                                            = 0x1902
	COLOR_INDEX1_EXT                                           = 0x80E2
	R16F_EXT                                                   = 0x822D
	DEPTH_ATTACHMENT                                           = 0x8D00
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	SGIX_sprite                                                = 1
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	BITMAP_TOKEN                                               = 0x0704
	RGB5_A1_OES                                                = 0x8057
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	DEPTH32F_STENCIL8                                          = 0x8CAD
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	SYNC_FLAGS                                                 = 0x9115
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	PASS_THROUGH_NV                                            = 0x86E6
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	TEXTURE8                                                   = 0x84C8
	CON_30_ATI                                                 = 0x895F
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	DOUBLE_MAT3_EXT                                            = 0x8F47
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	BINORMAL_ARRAY_EXT                                         = 0x843A
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	CUBIC_HP                                                   = 0x815F
	VERTEX_BLEND_ARB                                           = 0x86A7
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	VERSION_3_1                                                = 1
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	TEXTURE5_ARB                                               = 0x84C5
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	STENCIL_INDEX4                                             = 0x8D47
	SAMPLER_BUFFER_AMD                                         = 0x9001
	INT_IMAGE_2D_EXT                                           = 0x9058
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	GL_2PASS_1_SGIS                                            = 0x80A3
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	FLOAT_MAT2                                                 = 0x8B5A
	CLAMP_TO_EDGE                                              = 0x812F
	OP_NEGATE_EXT                                              = 0x8783
	COORD_REPLACE_NV                                           = 0x8862
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	TEXTURE25_ARB                                              = 0x84D9
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	VERTEX_SHADER_EXT                                          = 0x8780
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	DSDT8_MAG8_NV                                              = 0x870A
	ZERO_EXT                                                   = 0x87DD
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	LINE_WIDTH_RANGE                                           = 0x0B22
	RED_MAX_CLAMP_INGR                                         = 0x8564
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	COMBINE_ALPHA                                              = 0x8572
	STENCIL_BACK_FAIL                                          = 0x8801
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	COLOR_ATTACHMENT1                                          = 0x8CE1
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	ARB_imaging                                                = 1
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	OPERAND3_ALPHA_NV                                          = 0x859B
	MATRIX6_NV                                                 = 0x8636
	MODELVIEW16_ARB                                            = 0x8730
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	IMAGE_2D_RECT                                              = 0x904F
	PATH_JOIN_STYLE_NV                                         = 0x9079
	SGIX_clipmap                                               = 1
	CONVOLUTION_2D                                             = 0x8011
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	TEXTURE12_ARB                                              = 0x84CC
	DRAW_BUFFER7_ATI                                           = 0x882C
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	TEXTURE4                                                   = 0x84C4
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	POINT_SPRITE_ARB                                           = 0x8861
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	TEXTURE_PRIORITY                                           = 0x8066
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	FAILURE_NV                                                 = 0x9030
	PATH_GEN_MODE_NV                                           = 0x90B0
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	SRGB8_ALPHA8                                               = 0x8C43
	LINE_WIDTH                                                 = 0x0B21
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	LINE_SMOOTH                                                = 0x0B20
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	SRC_COLOR                                                  = 0x0300
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	BLEND_SRC                                                  = 0x0BE1
	VERTEX_ARRAY                                               = 0x8074
	DECR                                                       = 0x1E03
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	MAX_VARYING_VECTORS                                        = 0x8DFC
	SAMPLER_CUBE_ARB                                           = 0x8B60
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	SPRITE_AXIS_SGIX                                           = 0x814A
	PROGRAM                                                    = 0x82E2
	OPERAND0_RGB_EXT                                           = 0x8590
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	PERFMON_RESULT_AMD                                         = 0x8BC6
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	SGIX_shadow                                                = 1
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	OP_MOV_EXT                                                 = 0x8799
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	DOUBLE_MAT3x2                                              = 0x8F4B
	EXT_packed_pixels                                          = 1
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	TEXTURE_DEPTH_EXT                                          = 0x8071
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	FIXED_ONLY_ARB                                             = 0x891D
	SAMPLER_2D_SHADOW                                          = 0x8B62
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	BACK_LEFT                                                  = 0x0402
	RGBA_DXT5_S3TC                                             = 0x83A4
	COMPRESSED_RGB_ARB                                         = 0x84ED
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	DECODE_EXT                                                 = 0x8A49
	ATTACHED_SHADERS                                           = 0x8B85
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	INVARIANT_EXT                                              = 0x87C2
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	FOG_INDEX                                                  = 0x0B61
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	ARRAY_BUFFER_BINDING                                       = 0x8894
	REG_10_ATI                                                 = 0x892B
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	RENDERBUFFER                                               = 0x8D41
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	RGB_FLOAT16_APPLE                                          = 0x881B
	MATRIX18_ARB                                               = 0x88D2
	VERTEX_ID_NV                                               = 0x8C7B
	MATRIX4_ARB                                                = 0x88C4
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	LAYOUT_LINEAR_INTEL                                        = 1
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	POLYGON_OFFSET_EXT                                         = 0x8037
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	COLOR_SAMPLES_NV                                           = 0x8E20
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	OPERAND1_RGB                                               = 0x8591
	FLOAT_RGBA32_NV                                            = 0x888B
	POLYGON_MODE                                               = 0x0B40
	MAX_CLIP_PLANES                                            = 0x0D32
	IMAGE_SCALE_X_HP                                           = 0x8155
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	UNIFORM_BUFFER_START                                       = 0x8A29
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	RGBA16UI                                                   = 0x8D76
	RGBA32I                                                    = 0x8D82
	RENDERER                                                   = 0x1F01
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	SGIX_framezoom                                             = 1
	LINE_BIT                                                   = 0x00000004
	EDGE_FLAG_ARRAY                                            = 0x8079
	COLOR_INDEXES                                              = 0x1603
	MINMAX_FORMAT                                              = 0x802F
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	BUFFER_ACCESS_ARB                                          = 0x88BB
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	INT_IMAGE_3D_EXT                                           = 0x9059
	R8UI                                                       = 0x8232
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	IMAGE_CUBE_EXT                                             = 0x9050
	CLIP_PLANE2                                                = 0x3002
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	SAMPLER_3D_ARB                                             = 0x8B5F
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	UNSIGNED_INT64_NV                                          = 0x140F
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	OUTPUT_VERTEX_EXT                                          = 0x879A
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	VALIDATE_STATUS                                            = 0x8B83
	VERTEX_TEXTURE                                             = 0x829B
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	FRAGMENT_COLOR_EXT                                         = 0x834C
	MODELVIEW15_ARB                                            = 0x872F
	LUMINANCE16_SNORM                                          = 0x9019
	BACK_RIGHT                                                 = 0x0403
	NOR                                                        = 0x1508
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	INDEX_ARRAY_STRIDE                                         = 0x8086
	MIN_EXT                                                    = 0x8007
	TEXTURE_GATHER                                             = 0x82A2
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	RED_EXT                                                    = 0x1903
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	REPLICATE_BORDER_HP                                        = 0x8153
	SECONDARY_COLOR_NV                                         = 0x852D
	PRIMARY_COLOR_EXT                                          = 0x8577
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	RGB565_OES                                                 = 0x8D62
	TEXTURE18_ARB                                              = 0x84D2
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	INVARIANT_VALUE_EXT                                        = 0x87EA
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	CON_21_ATI                                                 = 0x8956
	FRONT_RIGHT                                                = 0x0401
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	REG_31_ATI                                                 = 0x8940
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	COLOR_ATTACHMENT8                                          = 0x8CE8
	DOUBLE_MAT2x3                                              = 0x8F49
	READ_BUFFER                                                = 0x0C02
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	SYNC_OBJECT_APPLE                                          = 0x8A53
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	PATH_MITER_LIMIT_NV                                        = 0x907A
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	MAJOR_VERSION                                              = 0x821B
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	BUFFER_ACCESS                                              = 0x88BB
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	COUNT_DOWN_NV                                              = 0x9089
	TRANSLATE_2D_NV                                            = 0x9090
	REG_26_ATI                                                 = 0x893B
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	DATA_BUFFER_AMD                                            = 0x9151
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	COLOR3_BIT_PGI                                             = 0x00010000
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	RGBA_FLOAT16_APPLE                                         = 0x881A
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	SATURATE_BIT_ATI                                           = 0x00000040
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	NONE                                                       = 0
	VARIABLE_G_NV                                              = 0x8529
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	CON_0_ATI                                                  = 0x8941
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	INT8_NV                                                    = 0x8FE0
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	DSDT_NV                                                    = 0x86F5
	EQUAL                                                      = 0x0202
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	PRIMARY_COLOR                                              = 0x8577
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	QUERY_NO_WAIT_NV                                           = 0x8E14
	POLYGON_BIT                                                = 0x00000008
	CULL_FACE                                                  = 0x0B44
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	LUMINANCE16UI_EXT                                          = 0x8D7A
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	DRAW_PIXEL_TOKEN                                           = 0x0705
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	RGBA4_EXT                                                  = 0x8056
	INTERLACE_READ_OML                                         = 0x8981
	DEPTH_TEST                                                 = 0x0B71
	MATRIX_MODE                                                = 0x0BA0
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	DRAW_BUFFER1                                               = 0x8826
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	INVERTED_SCREEN_W_REND                                     = 0x8491
	SIGNED_NORMALIZED                                          = 0x8F9C
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	RGBA16_EXT                                                 = 0x805B
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	COLOR_SUM_ARB                                              = 0x8458
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	FIXED_ONLY                                                 = 0x891D
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	SOURCE1_RGB_EXT                                            = 0x8581
	SGIX_resample                                              = 1
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	SGIS_pixel_texture                                         = 1
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	ROUND_NV                                                   = 0x90A4
	NOOP                                                       = 0x1505
	SPHERE_MAP                                                 = 0x2402
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	SAMPLER_2D_ARB                                             = 0x8B5E
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	SURFACE_REGISTERED_NV                                      = 0x86FD
	W_EXT                                                      = 0x87D8
	NEGATIVE_Z_EXT                                             = 0x87DB
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	DEPTH_RENDERABLE                                           = 0x8287
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	FRAGMENT_SHADER                                            = 0x8B30
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	EXT_convolution                                            = 1
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	TEXTURE10                                                  = 0x84CA
	MIRROR_CLAMP_ATI                                           = 0x8742
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	TEXTURE_GEN_Q                                              = 0x0C63
	OBJECT_LINEAR                                              = 0x2401
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	SOURCE1_ALPHA_EXT                                          = 0x8589
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	INCR_WRAP_OES                                              = 0x8507
	READ_ONLY                                                  = 0x88B8
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	SHADER_TYPE                                                = 0x8B4F
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	CURRENT_BINORMAL_EXT                                       = 0x843C
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	PATCH_VERTICES                                             = 0x8E72
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	EXT_polygon_offset                                         = 1
	SGIX_tag_sample_buffer                                     = 1
	TRANSFORM_BIT                                              = 0x00001000
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	SPECULAR                                                   = 0x1202
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	NORMAL_ARRAY_EXT                                           = 0x8075
	TEXTURE_CUBE_MAP                                           = 0x8513
	INT_VEC3_ARB                                               = 0x8B54
	FACTOR_MIN_AMD                                             = 0x901C
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	T2F_C3F_V3F                                                = 0x2A2A
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	DONT_CARE                                                  = 0x1100
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	UNDEFINED_VERTEX                                           = 0x8260
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	STENCIL_BACK_REF                                           = 0x8CA3
	DOUBLE_MAT4_EXT                                            = 0x8F48
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	ALPHA_BIAS                                                 = 0x0D1D
	GL_4PASS_3_SGIS                                            = 0x80A7
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	RGBA                                                       = 0x1908
	INTENSITY_EXT                                              = 0x8049
	ALPHA_MAX_SGIX                                             = 0x8321
	IDENTITY_NV                                                = 0x862A
	DEPTH_COMPONENT32F                                         = 0x8CAC
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	SGIS_texture_lod                                           = 1
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	SOURCE0_ALPHA_EXT                                          = 0x8588
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	TEXTURE_WRAP_R_OES                                         = 0x8072
	SGIX_ir_instrument1                                        = 1
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	REG_21_ATI                                                 = 0x8936
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	POSITION                                                   = 0x1203
	GL_4PASS_1_SGIS                                            = 0x80A5
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	BOOL_VEC4_ARB                                              = 0x8B59
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	INTENSITY16_SNORM                                          = 0x901B
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	INTENSITY16                                                = 0x804D
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	RGBA4                                                      = 0x8056
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	RGB32F                                                     = 0x8815
	TEXTURE_RESIDENT                                           = 0x8067
	VIEW_CLASS_8_BITS                                          = 0x82CB
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	PROXY_TEXTURE_2D                                           = 0x8064
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	IMAGE_SCALE_Y_HP                                           = 0x8156
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	AVERAGE_EXT                                                = 0x8335
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	COMBINE_RGB_ARB                                            = 0x8571
	DRAW_BUFFER2_ATI                                           = 0x8827
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	FIELD_UPPER_NV                                             = 0x9022
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	UTF16_NV                                                   = 0x909B
	MAX                                                        = 0x8008
	R32F                                                       = 0x822E
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	STENCIL_VALUE_MASK                                         = 0x0B93
	COLOR_INDEX16_EXT                                          = 0x80E7
	DEBUG_ASSERT_MESA                                          = 0x875B
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	COPY_WRITE_BUFFER                                          = 0x8F37
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	UNIFORM                                                    = 0x92E1
	COMBINER5_NV                                               = 0x8555
	DS_SCALE_NV                                                = 0x8710
	SET_AMD                                                    = 0x874A
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	SIGNED_INTENSITY_NV                                        = 0x8707
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	SGX_BINARY_IMG                                             = 0x8C0A
	R8_SNORM                                                   = 0x8F94
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	BGRA_EXT                                                   = 0x80E1
	DRAW_BUFFER11_NV                                           = 0x8830
	LERP_ATI                                                   = 0x8969
	VERTEX_SHADER_ARB                                          = 0x8B31
	DOUBLE_MAT4x2                                              = 0x8F4D
	QUARTER_BIT_ATI                                            = 0x00000010
	LIST_BIT                                                   = 0x00020000
	STENCIL_BACK_FUNC                                          = 0x8800
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	IUI_V3F_EXT                                                = 0x81AE
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	TEXTURE23                                                  = 0x84D7
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	ACCUM_RED_BITS                                             = 0x0D58
	POINT_SIZE_MIN_ARB                                         = 0x8126
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	FLOAT_RGB_NV                                               = 0x8882
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	MAX_PATCH_VERTICES                                         = 0x8E7D
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	SLIM12S_SGIX                                               = 0x831F
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	EIGHTH_BIT_ATI                                             = 0x00000020
	MAX_SUBROUTINES                                            = 0x8DE7
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	SPRITE_AXIAL_SGIX                                          = 0x814C
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	PROGRAM_LENGTH_ARB                                         = 0x8627
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	INT_IMAGE_CUBE                                             = 0x905B
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	SPOT_DIRECTION                                             = 0x1204
	TEXTURE_ENV                                                = 0x2300
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	PALETTE4_RGBA4_OES                                         = 0x8B93
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	GEOMETRY_SHADER                                            = 0x8DD9
	MAX_INTEGER_SAMPLES                                        = 0x9110
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	FLOAT_MAT2x4                                               = 0x8B66
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	SGIX_depth_texture                                         = 1
	ELEMENT_ARRAY_ATI                                          = 0x8768
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	LINE_RESET_TOKEN                                           = 0x0707
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	GL_4PASS_2_SGIS                                            = 0x80A6
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	MODELVIEW31_ARB                                            = 0x873F
	STATIC_COPY_ARB                                            = 0x88E6
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	PROXY_COLOR_TABLE                                          = 0x80D3
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	GEQUAL                                                     = 0x0206
	HISTOGRAM_SINK_EXT                                         = 0x802D
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	ACCUM_BLUE_BITS                                            = 0x0D5A
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	CLIP_DISTANCE5                                             = 0x3005
	DSDT_MAG_VIB_NV                                            = 0x86F7
	FRACTIONAL_EVEN                                            = 0x8E7C
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	PROGRAM_RESIDENT_NV                                        = 0x8647
	DRAW_BUFFER5                                               = 0x882A
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	MODELVIEW22_ARB                                            = 0x8736
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	CMYKA_EXT                                                  = 0x800D
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	INVALID_ENUM                                               = 0x0500
	STENCIL_FAIL                                               = 0x0B94
	FUNC_ADD_OES                                               = 0x8006
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	FLOAT_MAT3                                                 = 0x8B5B
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	CON_9_ATI                                                  = 0x894A
	BLUE_BIAS                                                  = 0x0D1B
	T4F_V4F                                                    = 0x2A28
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	DUAL_ALPHA12_SGIS                                          = 0x8112
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	BLEND_DST_ALPHA                                            = 0x80CA
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	NORMAL_MAP                                                 = 0x8511
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	SAMPLE_POSITION                                            = 0x8E50
	RETURN                                                     = 0x0102
	POINT_SMOOTH_HINT                                          = 0x0C51
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	MATRIX20_ARB                                               = 0x88D4
	VERTEX_SUBROUTINE                                          = 0x92E8
	FOG_START                                                  = 0x0B63
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	STREAM_COPY_ARB                                            = 0x88E2
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	DISCARD_ATI                                                = 0x8763
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	TEXTURE_DEPTH                                              = 0x8071
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	SAMPLE_MASK                                                = 0x8E51
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	ORDER                                                      = 0x0A01
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	BYTE                                                       = 0x1400
	FOG_COORD_ARRAY                                            = 0x8457
	STATIC_ATI                                                 = 0x8760
	GL_8X_BIT_ATI                                              = 0x00000004
	COLOR_ATTACHMENT14                                         = 0x8CEE
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	RESTART_PATH_NV                                            = 0xF0
	MULTIVIEW_EXT                                              = 0x90F1
	MATRIX_STRIDE                                              = 0x92FF
	EXP                                                        = 0x0800
	MAP2_VERTEX_3                                              = 0x0DB7
	PACK_RESAMPLE_SGIX                                         = 0x842C
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	DRAW_BUFFER10                                              = 0x882F
	MATRIX31_ARB                                               = 0x88DF
	TRANSFORM_FEEDBACK                                         = 0x8E22
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	POINT_TOKEN                                                = 0x0701
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	MITER_REVERT_NV                                            = 0x90A7
	FOG_HINT                                                   = 0x0C54
	FOG_COORD                                                  = 0x8451
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	FLOAT16_VEC2_NV                                            = 0x8FF9
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	FASTEST                                                    = 0x1101
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	SGI_color_table                                            = 1
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	MULTISAMPLE_SGIS                                           = 0x809D
	DRAW_BUFFER9                                               = 0x882E
	DRAW_BUFFER12_NV                                           = 0x8831
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	GL_422_EXT                                                 = 0x80CC
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	SHADER_BINARY_DMP                                          = 0x9250
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	DOT3_RGBA_ARB                                              = 0x86AF
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	TRACE_MASK_MESA                                            = 0x8755
	PALETTE4_RGBA8_OES                                         = 0x8B91
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	COMPILE                                                    = 0x1300
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	TEXTURE_RECTANGLE                                          = 0x84F5
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	COPY_INVERTED                                              = 0x150C
	LUMINANCE6_ALPHA2                                          = 0x8044
	OCCLUSION_TEST_HP                                          = 0x8165
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	SWIZZLE_STR_ATI                                            = 0x8976
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	POLYGON_OFFSET_LINE                                        = 0x2A02
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	INT_VEC2                                                   = 0x8B53
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	RESCALE_NORMAL                                             = 0x803A
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	TEXTURE_WRAP_R                                             = 0x8072
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	TIME_ELAPSED_EXT                                           = 0x88BF
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	DST_COLOR                                                  = 0x0306
	BITMAP                                                     = 0x1A00
	GL_4PASS_2_EXT                                             = 0x80A6
	NUM_EXTENSIONS                                             = 0x821D
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	MAX_SAMPLES_IMG                                            = 0x9135
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	SRC0_ALPHA                                                 = 0x8588
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	UPPER_LEFT                                                 = 0x8CA2
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	VERSION_1_1                                                = 1
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	INDEX_MATERIAL_EXT                                         = 0x81B8
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	SIGNED_HILO16_NV                                           = 0x86FA
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	FOG_MODE                                                   = 0x0B65
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	SOURCE0_RGB                                                = 0x8580
	INTERPOLATE_ARB                                            = 0x8575
	TEXTURE_1D_ARRAY                                           = 0x8C18
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	UNSIGNED_INT8_NV                                           = 0x8FEC
	SGIS_point_line_texgen                                     = 1
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	INTENSITY16I_EXT                                           = 0x8D8B
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	EXT_shared_texture_palette                                 = 1
	INVALID_VALUE                                              = 0x0501
	DITHER                                                     = 0x0BD0
	TEXTURE1_ARB                                               = 0x84C1
	DRAW_BUFFER4_NV                                            = 0x8829
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	SHADER                                                     = 0x82E1
	OPERAND2_RGB_ARB                                           = 0x8592
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	INFO_LOG_LENGTH                                            = 0x8B84
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	RECT_NV                                                    = 0xF6
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	LINE_STRIP                                                 = 0x0003
	T2F_V3F                                                    = 0x2A27
	RGB12_EXT                                                  = 0x8053
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	COMBINE_RGB_EXT                                            = 0x8571
	VERTEX_STREAM3_ATI                                         = 0x876F
	SGIX_convolution_accuracy                                  = 1
	SGIX_fog_offset                                            = 1
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	MAP_TESSELLATION_NV                                        = 0x86C2
	CON_12_ATI                                                 = 0x894D
	INT_VEC4_ARB                                               = 0x8B55
	ALPHA_INTEGER                                              = 0x8D97
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	MATRIX22_ARB                                               = 0x88D6
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	REG_27_ATI                                                 = 0x893C
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	RGB565                                                     = 0x8D62
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	BUFFER_MAP_LENGTH                                          = 0x9120
	ALL_STATIC_DATA_IBM                                        = 103060
	UNPACK_LSB_FIRST                                           = 0x0CF1
	BLEND_EQUATION_OES                                         = 0x8009
	ACTIVE_TEXTURE                                             = 0x84E0
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	ALPHA4_EXT                                                 = 0x803B
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	VERSION_2_0                                                = 1
	RG_INTEGER                                                 = 0x8228
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	STACK_OVERFLOW                                             = 0x0503
	PACK_ALIGNMENT                                             = 0x0D05
	READ_ONLY_ARB                                              = 0x88B8
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	ALPHA_TEST_QCOM                                            = 0x0BC0
	IMAGE_2D_EXT                                               = 0x904D
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	CONSTANT_COLOR0_NV                                         = 0x852A
	INDEX_ARRAY_TYPE                                           = 0x8085
	RESTART_SUN                                                = 0x0001
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	LINES_ADJACENCY                                            = 0x000A
	ALPHA_TEST_FUNC                                            = 0x0BC1
	IUI_N3F_V2F_EXT                                            = 0x81AF
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	R16I                                                       = 0x8233
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	FLOAT_MAT4                                                 = 0x8B5C
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	TEXTURE27                                                  = 0x84DB
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	REG_18_ATI                                                 = 0x8933
	STENCIL_INDEX1_EXT                                         = 0x8D46
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	COMPRESSED_ALPHA                                           = 0x84E9
	MAGNITUDE_SCALE_NV                                         = 0x8712
	SGIX_texture_scale_bias                                    = 1
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	COMBINER_MAPPING_NV                                        = 0x8543
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	STENCIL_RENDERABLE                                         = 0x8288
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	COMPRESSED_RGB                                             = 0x84ED
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	CONSTANT_EXT                                               = 0x8576
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	ALPHA_MIN_SGIX                                             = 0x8320
	TEXTURE2                                                   = 0x84C2
	RGBA16F                                                    = 0x881A
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	OP_FRAC_EXT                                                = 0x8789
	CON_3_ATI                                                  = 0x8944
	LOCATION                                                   = 0x930E
	MAP2_VERTEX_4                                              = 0x0DB8
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	MODELVIEW11_ARB                                            = 0x872B
	DRAW_BUFFER14_ATI                                          = 0x8833
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	MODELVIEW3_ARB                                             = 0x8723
	SIGNED_HILO8_NV                                            = 0x885F
	PATH_STROKE_MASK_NV                                        = 0x9084
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	OR_INVERTED                                                = 0x150D
	RGBA_FLOAT16_ATI                                           = 0x881A
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	TEXTURE17                                                  = 0x84D1
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	STENCIL_BUFFER_BIT                                         = 0x00000400
	NORMAL_MAP_ARB                                             = 0x8511
	BOOL_VEC3                                                  = 0x8B58
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	STANDARD_FONT_NAME_NV                                      = 0x9072
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	DECR_WRAP_OES                                              = 0x8508
	LOCAL_EXT                                                  = 0x87C4
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	COUNTER_RANGE_AMD                                          = 0x8BC1
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	CURRENT_RASTER_COLOR                                       = 0x0B04
	PACK_ROW_LENGTH                                            = 0x0D02
	TEXTURE_WRAP_S                                             = 0x2802
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	SLUMINANCE8_NV                                             = 0x8C47
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	COMBINE_RGB                                                = 0x8571
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	REG_20_ATI                                                 = 0x8935
	LARGE_CW_ARC_TO_NV                                         = 0x18
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	BUFFER_SIZE_ARB                                            = 0x8764
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	EMBOSS_LIGHT_NV                                            = 0x855D
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	PATH_FORMAT_SVG_NV                                         = 0x9070
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	LUMINANCE8I_EXT                                            = 0x8D92
	NEVER                                                      = 0x0200
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	INT8_VEC3_NV                                               = 0x8FE2
	SHADE_MODEL                                                = 0x0B54
	LUMINANCE16_EXT                                            = 0x8042
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	FLOAT_MAT4x3                                               = 0x8B6A
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	DEPTH_STENCIL_EXT                                          = 0x84F9
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	MINMAX_FORMAT_EXT                                          = 0x802F
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	CURRENT_FOG_COORDINATE                                     = 0x8453
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	POINT_SPRITE                                               = 0x8861
	RGB16UI_EXT                                                = 0x8D77
	SAMPLER_BUFFER                                             = 0x8DC2
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	OBJECT_TYPE                                                = 0x9112
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	C4F_N3F_V3F                                                = 0x2A26
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	MIRRORED_REPEAT_OES                                        = 0x8370
	MODULATE_COLOR_IMG                                         = 0x8C04
	HIGH_INT                                                   = 0x8DF5
	TEXTURE_WRAP_T                                             = 0x2803
	OPERAND0_ALPHA                                             = 0x8598
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	MODELVIEW23_ARB                                            = 0x8737
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	RGB_FLOAT32_APPLE                                          = 0x8815
	PROGRAM_FORMAT_ARB                                         = 0x8876
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	DEPTH_BOUNDS_EXT                                           = 0x8891
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	MINMAX_EXT                                                 = 0x802E
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	RG32F                                                      = 0x8230
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	DEPTH_WRITEMASK                                            = 0x0B72
	GL_3_BYTES                                                 = 0x1408
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	DRAW_BUFFER6_ATI                                           = 0x882B
	MATRIX30_ARB                                               = 0x88DE
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	ADJACENT_PAIRS_NV                                          = 0x90AE
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	CURRENT_COLOR                                              = 0x0B00
	TEXTURE_4D_SGIS                                            = 0x8134
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	OUTPUT_COLOR1_EXT                                          = 0x879C
	SAMPLER_BINDING                                            = 0x8919
	INT_VEC3                                                   = 0x8B54
	COUNTER_TYPE_AMD                                           = 0x8BC0
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	ADD_BLEND_IMG                                              = 0x8C09
	MEDIUM_INT                                                 = 0x8DF4
	PATCHES                                                    = 0x000E
	HALF_FLOAT                                                 = 0x140B
	HALF_APPLE                                                 = 0x140B
	SHININESS                                                  = 0x1601
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	ALPHA_FLOAT32_ATI                                          = 0x8816
	CURRENT_INDEX                                              = 0x0B01
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	PATH_FILL_MASK_NV                                          = 0x9081
	REPLACE                                                    = 0x1E01
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	FOG_COORDINATE_ARRAY                                       = 0x8457
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	TEXTURE3                                                   = 0x84C3
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	MODELVIEW0_ARB                                             = 0x1700
	VERTEX_SOURCE_ATI                                          = 0x8774
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	CON_26_ATI                                                 = 0x895B
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	TEXTURE_COORD_ARRAY                                        = 0x8078
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	COMBINE_ALPHA_EXT                                          = 0x8572
	SAMPLER_2D                                                 = 0x8B5E
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	RGB16_EXT                                                  = 0x8054
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	READ_PIXELS_TYPE                                           = 0x828E
	BUFFER_USAGE_ARB                                           = 0x8765
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	GREEN                                                      = 0x1904
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	DOT3_RGBA_IMG                                              = 0x86AF
	DRAW_BUFFER9_ATI                                           = 0x882E
	MATRIX1_ARB                                                = 0x88C1
	COMPILE_STATUS                                             = 0x8B81
	SYNC_FENCE_APPLE                                           = 0x9116
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	MAX_DRAW_BUFFERS                                           = 0x8824
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	MATRIX1_NV                                                 = 0x8631
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	BUFFER_MAP_POINTER                                         = 0x88BD
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	INT64_VEC2_NV                                              = 0x8FE9
	VENDOR                                                     = 0x1F00
	VIEW_CLASS_64_BITS                                         = 0x82C6
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	TEXTURE5                                                   = 0x84C5
	COMBINE4_NV                                                = 0x8503
	FUNC_ADD_EXT                                               = 0x8006
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	STENCIL_INDEX8_EXT                                         = 0x8D48
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	MODELVIEW_MATRIX                                           = 0x0BA6
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	COMBINER_SCALE_NV                                          = 0x8548
	POLYGON                                                    = 0x0009
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	STENCIL                                                    = 0x1802
	GL_422_REV_EXT                                             = 0x80CD
	MAP2_TANGENT_EXT                                           = 0x8445
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	DRAW_BUFFER12_ATI                                          = 0x8831
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	INT16_NV                                                   = 0x8FE4
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	ALPHA_TEST_REF                                             = 0x0BC2
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	READ_PIXELS_FORMAT                                         = 0x828D
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	MATRIX13_ARB                                               = 0x88CD
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	TEXTURE_BIT                                                = 0x00040000
	AUTO_NORMAL                                                = 0x0D80
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	COMP_BIT_ATI                                               = 0x00000002
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	PACK_SKIP_ROWS                                             = 0x0D03
	MULTISAMPLE                                                = 0x809D
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	LEQUAL                                                     = 0x0203
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	ACCUM_GREEN_BITS                                           = 0x0D59
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	FLOAT_RG16_NV                                              = 0x8886
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	REPLICATE_BORDER                                           = 0x8153
	TEXTURE15                                                  = 0x84CF
	TEXTURE_COMPRESSED                                         = 0x86A1
	VARIABLE_B_NV                                              = 0x8524
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	IMAGE_1D                                                   = 0x904C
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	SGIX_async_pixel                                           = 1
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	RGBA12                                                     = 0x805A
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	CLAMP_TO_BORDER                                            = 0x812D
	DT_BIAS_NV                                                 = 0x8717
	STATIC_READ                                                = 0x88E5
	PALETTE8_RGB8_OES                                          = 0x8B95
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	SIGNED_NEGATE_NV                                           = 0x853D
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	PATH_STENCIL_REF_NV                                        = 0x90B8
	MATRIX27_ARB                                               = 0x88DB
	BOOL_VEC2_ARB                                              = 0x8B57
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	REFLECTION_MAP_NV                                          = 0x8512
	MAP1_VERTEX_4                                              = 0x0D98
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	CURRENT_MATRIX_NV                                          = 0x8641
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	CON_29_ATI                                                 = 0x895E
	RGB32I_EXT                                                 = 0x8D83
	BGRA_INTEGER                                               = 0x8D9B
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	SKIP_DECODE_EXT                                            = 0x8A4A
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	DEPTH_BIAS                                                 = 0x0D1F
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	TEXTURE31                                                  = 0x84DF
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	MOVE_TO_NV                                                 = 0x02
	GL_4PASS_0_EXT                                             = 0x80A4
	PN_TRIANGLES_ATI                                           = 0x87F0
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	COLOR_WRITEMASK                                            = 0x0C23
	DOUBLE_EXT                                                 = 0x140A
	XOR                                                        = 0x1506
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	ALPHA12_EXT                                                = 0x803D
	MODELVIEW10_ARB                                            = 0x872A
	DRAW_BUFFER3_ARB                                           = 0x8828
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	YCBAYCR8A_4224_NV                                          = 0x9032
	SGIX_impact_pixel_texture                                  = 1
	COLOR_ARRAY_SIZE                                           = 0x8081
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	MAX_TEXTURE_COORDS                                         = 0x8871
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	DEBUG_OBJECT_MESA                                          = 0x8759
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MIN_LOD_WARNING_AMD                                        = 0x919C
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	DRAW_BUFFER1_ARB                                           = 0x8826
	FLOAT_RGBA_NV                                              = 0x8883
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	SAMPLE_SHADING_ARB                                         = 0x8C36
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	POLYGON_OFFSET_FILL                                        = 0x8037
	INTENSITY4_EXT                                             = 0x804A
	TEXTURE24_ARB                                              = 0x84D8
	R1UI_V3F_SUN                                               = 0x85C4
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	COLOR_ATTACHMENT11                                         = 0x8CEB
	ONE                                                        = 1
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	CURRENT_QUERY_EXT                                          = 0x8865
	IMAGE_3D                                                   = 0x904E
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	VERSION                                                    = 0x1F02
	SCALAR_EXT                                                 = 0x87BE
	RGB32F_ARB                                                 = 0x8815
	COORD_REPLACE_ARB                                          = 0x8862
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	SGIX_pixel_texture                                         = 1
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	CURRENT_PROGRAM                                            = 0x8B8D
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	STENCIL_FUNC                                               = 0x0B92
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	TEXTURE_SAMPLES                                            = 0x9106
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	SRGB_ALPHA                                                 = 0x8C42
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	BLEND_EQUATION_ALPHA                                       = 0x883D
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	HI_SCALE_NV                                                = 0x870E
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	INDEX_OFFSET                                               = 0x0D13
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	ALPHA32I_EXT                                               = 0x8D84
	MODELVIEW8_ARB                                             = 0x8728
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	SKIP_COMPONENTS1_NV                                        = -6
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	RED_BIT_ATI                                                = 0x00000001
	TEXTURE_3D                                                 = 0x806F
	T2F_IUI_V2F_EXT                                            = 0x81B1
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	CURRENT_FOG_COORD                                          = 0x8453
	VERTEX_STREAM2_ATI                                         = 0x876E
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	MODELVIEW1_EXT                                             = 0x850A
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	CONTINUOUS_AMD                                             = 0x9007
	RG32I                                                      = 0x823B
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	CON_7_ATI                                                  = 0x8948
	SUCCESS_NV                                                 = 0x902F
	LO_BIAS_NV                                                 = 0x8715
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	STENCIL_INDEX16                                            = 0x8D49
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	STATIC_COPY                                                = 0x88E6
	POINT_BIT                                                  = 0x00000002
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	COLOR_MATERIAL                                             = 0x0B57
	PREVIOUS_EXT                                               = 0x8578
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	PATH_COORD_COUNT_NV                                        = 0x909E
	SEPARABLE_2D_EXT                                           = 0x8012
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	MATRIX17_ARB                                               = 0x88D1
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	SHADER_CONSISTENT_NV                                       = 0x86DD
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	RGB32I                                                     = 0x8D83
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	MAX_WIDTH                                                  = 0x827E
	TEXTURE1                                                   = 0x84C1
	TRANSPOSE_NV                                               = 0x862C
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	IS_ROW_MAJOR                                               = 0x9300
	FRAMEZOOM_SGIX                                             = 0x818B
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	COLOR_ATTACHMENT6                                          = 0x8CE6
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	COMBINER4_NV                                               = 0x8554
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	SIGNALED_APPLE                                             = 0x9119
	IUI_V2F_EXT                                                = 0x81AD
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	OBJECT_TYPE_APPLE                                          = 0x9112
	BUFFER_VARIABLE                                            = 0x92E5
	IS_PER_PATCH                                               = 0x92E7
	TEXTURE_ENV_MODE                                           = 0x2200
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	OUTPUT_FOG_EXT                                             = 0x87BD
	RELATIVE_LINE_TO_NV                                        = 0x05
	GL_1PASS_SGIS                                              = 0x80A1
	COLOR_TABLE_WIDTH                                          = 0x80D9
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	SRGB                                                       = 0x8C40
	RENDERBUFFER_EXT                                           = 0x8D41
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	CON_16_ATI                                                 = 0x8951
	COPY                                                       = 0x1503
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	LIGHT6                                                     = 0x4006
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	REG_24_ATI                                                 = 0x8939
	MAP_STENCIL                                                = 0x0D11
	OBJECT_PLANE                                               = 0x2501
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	UNPACK_RESAMPLE_OML                                        = 0x8985
	OR                                                         = 0x1507
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	LOGIC_OP                                                   = 0x0BF1
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	OP_SET_LT_EXT                                              = 0x878D
	CON_23_ATI                                                 = 0x8958
	FUNC_ADD                                                   = 0x8006
	MATRIX15_ARB                                               = 0x88CF
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	EXT_subtexture                                             = 1
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	STENCIL_COMPONENTS                                         = 0x8285
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	INT_VEC4                                                   = 0x8B55
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	LINES                                                      = 0x0001
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	POINTS                                                     = 0x0000
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	MAX_TEXTURE_SIZE                                           = 0x0D33
	MIN                                                        = 0x8007
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	TEXTURE29_ARB                                              = 0x84DD
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	TRIANGULAR_NV                                              = 0x90A5
	LAYOUT_DEFAULT_INTEL                                       = 0
	COMPRESSED_RGBA                                            = 0x84EE
	OP_POWER_EXT                                               = 0x8793
	HALF_BIT_ATI                                               = 0x00000008
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	CON_17_ATI                                                 = 0x8952
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	BOOL_ARB                                                   = 0x8B56
	NORMALIZE                                                  = 0x0BA1
	CLIP_PLANE4                                                = 0x3004
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	GL_4PASS_3_EXT                                             = 0x80A7
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	TEXTURE13_ARB                                              = 0x84CD
	HILO16_NV                                                  = 0x86F8
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	SGIX_texture_coordinate_clamp                              = 1
	CULL_FACE_MODE                                             = 0x0B45
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	ACCUM_BUFFER_BIT                                           = 0x00000200
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	COMBINER1_NV                                               = 0x8551
	ALPHA16_EXT                                                = 0x803E
	TEXTURE_SHADOW                                             = 0x82A1
	MATRIX2_NV                                                 = 0x8632
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	ABGR_EXT                                                   = 0x8000
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	SGIS_texture_filter4                                       = 1
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	PREVIOUS                                                   = 0x8578
	RGBA8UI                                                    = 0x8D7C
	RGB16_SNORM                                                = 0x8F9A
	BEVEL_NV                                                   = 0x90A6
	SYNC_FLAGS_APPLE                                           = 0x9115
	COLOR4_BIT_PGI                                             = 0x00020000
	SAMPLE_MASK_SGIS                                           = 0x80A0
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	VIBRANCE_BIAS_NV                                           = 0x8719
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	FLOAT_MAT2x3                                               = 0x8B65
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	PIXEL_PACK_BUFFER                                          = 0x88EB
	FLOAT16_VEC3_NV                                            = 0x8FFA
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	COLOR_ATTACHMENT9                                          = 0x8CE9
	INTENSITY4                                                 = 0x804A
	SEPARABLE_2D                                               = 0x8012
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	SYNC_FENCE                                                 = 0x9116
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	DRAW_BUFFER7_ARB                                           = 0x882C
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	MAX_VARYING_FLOATS                                         = 0x8B4B
	FLOAT_MAT4x2                                               = 0x8B69
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	ETC1_RGB8_OES                                              = 0x8D64
	DOUBLE_VEC4                                                = 0x8FFE
	INDEX_BIT_PGI                                              = 0x00080000
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	DRAW_BUFFER13_ATI                                          = 0x8832
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	SAMPLER_1D_SHADOW                                          = 0x8B61
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	CONSTANT_COLOR_EXT                                         = 0x8001
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	QUERY_OBJECT_EXT                                           = 0x9153
	HINT_BIT                                                   = 0x00008000
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	EXT_histogram                                              = 1
	READ_WRITE_ARB                                             = 0x88BA
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	SURFACE_MAPPED_NV                                          = 0x8700
	FLOAT_RGB16_NV                                             = 0x8888
	PROVOKING_VERTEX                                           = 0x8E4F
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	SAMPLE_MASK_VALUE                                          = 0x8E52
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	MODELVIEW9_ARB                                             = 0x8729
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	BOOL                                                       = 0x8B56
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	FLOAT_R16_NV                                               = 0x8884
	BOOL_VEC3_ARB                                              = 0x8B58
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	COLOR_ATTACHMENT5                                          = 0x8CE5
	GREEN_INTEGER                                              = 0x8D95
	BGRA_INTEGER_EXT                                           = 0x8D9B
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	DEBUG_TYPE_ERROR                                           = 0x824C
	CON_1_ATI                                                  = 0x8942
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	INCR_WRAP                                                  = 0x8507
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	NORMAL_BIT_PGI                                             = 0x08000000
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	RGB8_SNORM                                                 = 0x8F96
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	SGIX_texture_lod_bias                                      = 1
	DEPTH_BITS                                                 = 0x0D56
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	VIEW_CLASS_32_BITS                                         = 0x82C8
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	RG16                                                       = 0x822C
	INTENSITY32I_EXT                                           = 0x8D85
	VIDEO_BUFFER_NV                                            = 0x9020
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	TEXTURE_2D                                                 = 0x0DE1
	EMISSION                                                   = 0x1600
	LINE                                                       = 0x1B01
	DRAW_BUFFER6_ARB                                           = 0x882B
	REG_9_ATI                                                  = 0x892A
	V2F                                                        = 0x2A20
	COMPUTE_TEXTURE                                            = 0x82A0
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	TEXTURE_RED_SIZE                                           = 0x805C
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	UNIFORM_SIZE                                               = 0x8A38
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	BUFFER_OBJECT_EXT                                          = 0x9151
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	GENERATE_MIPMAP                                            = 0x8191
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	INT_VEC2_ARB                                               = 0x8B53
	COLOR_ATTACHMENT12                                         = 0x8CEC
	PATH_GEN_COEFF_NV                                          = 0x90B1
	MODULATE                                                   = 0x2100
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	SMALL_CW_ARC_TO_NV                                         = 0x14
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	DRAW_BUFFER14_NV                                           = 0x8833
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	DOUBLE_VEC2                                                = 0x8FFC
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	POINT_SIZE                                                 = 0x0B11
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	DOT3_RGB                                                   = 0x86AE
	Y_EXT                                                      = 0x87D6
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	CON_27_ATI                                                 = 0x895C
	SRGB_ALPHA_EXT                                             = 0x8C42
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	REG_29_ATI                                                 = 0x893E
	PATH_FILL_MODE_NV                                          = 0x9080
	COEFF                                                      = 0x0A00
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	MAX_HEIGHT                                                 = 0x827F
	MAX_DEPTH                                                  = 0x8280
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	RGB16F_EXT                                                 = 0x881B
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	CONSTANT_COLOR1_NV                                         = 0x852B
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	REFLECTION_MAP_ARB                                         = 0x8512
	STATE_RESTORE                                              = 0x8BDC
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	OP_DOT3_EXT                                                = 0x8784
	RGB16F_ARB                                                 = 0x881B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	CLIP_DISTANCE1                                             = 0x3001
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	GL_2PASS_1_EXT                                             = 0x80A3
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	INDEX_ARRAY_LIST_IBM                                       = 103073
	LINES_ADJACENCY_EXT                                        = 0x000A
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	REPLACE_OLDEST_SUN                                         = 0x0003
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	PACK_IMAGE_HEIGHT                                          = 0x806C
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	WRAP_BORDER_SUN                                            = 0x81D4
	R8_EXT                                                     = 0x8229
	COORD_REPLACE                                              = 0x8862
	UNSIGNED_INT                                               = 0x1405
	OPERAND2_ALPHA                                             = 0x859A
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	OP_RECIP_EXT                                               = 0x8794
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	LIGHT1                                                     = 0x4001
	SUBTRACT                                                   = 0x84E7
	VARIABLE_C_NV                                              = 0x8525
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	VARIANT_VALUE_EXT                                          = 0x87E4
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	ALPHA16_SNORM                                              = 0x9018
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	VERTEX4_BIT_PGI                                            = 0x00000008
	ONE_MINUS_DST_COLOR                                        = 0x0307
	CLIP_PLANE3                                                = 0x3003
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	DEPTH_CLAMP_NV                                             = 0x864F
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	ALPHA_INTEGER_EXT                                          = 0x8D97
	R16_SNORM                                                  = 0x8F98
	EXT_rescale_normal                                         = 1
	OP_MIN_EXT                                                 = 0x878B
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	INDEX_TEST_EXT                                             = 0x81B5
	MAX_LAYERS                                                 = 0x8281
	DSDT_MAG_NV                                                = 0x86F6
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	NEGATIVE_W_EXT                                             = 0x87DC
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	PURGEABLE_APPLE                                            = 0x8A1D
	TEXTURE_COORD_NV                                           = 0x8C79
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	SGIS_sharpen_texture                                       = 1
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	CONTEXT_FLAGS                                              = 0x821E
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	SGIX_reference_plane                                       = 1
	RGB10_A2_EXT                                               = 0x8059
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	INTERPOLATE_EXT                                            = 0x8575
	VERTEX_ARRAY_SIZE                                          = 0x807A
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	PROXY_TEXTURE_3D                                           = 0x8070
	SRC2_RGB                                                   = 0x8582
	MATRIX16_ARB                                               = 0x88D0
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	SRGB8_NV                                                   = 0x8C41
	INTENSITY_SNORM                                            = 0x9013
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	LOWER_LEFT                                                 = 0x8CA1
	TESS_EVALUATION_SHADER                                     = 0x8E87
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	MODELVIEW17_ARB                                            = 0x8731
	MATRIX6_ARB                                                = 0x88C6
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	INDEX_ARRAY                                                = 0x8077
	PROGRAM_PARAMETER_NV                                       = 0x8644
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	SGIX_shadow_ambient                                        = 1
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	COMPUTE_SUBROUTINE                                         = 0x92ED
	MAP_COLOR                                                  = 0x0D10
	VIEW_CLASS_48_BITS                                         = 0x82C7
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	RG16_SNORM                                                 = 0x8F99
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	LESS                                                       = 0x0201
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	COMPRESSED_INTENSITY                                       = 0x84EC
	PRIMARY_COLOR_NV                                           = 0x852C
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	IMAGE_BINDING_FORMAT                                       = 0x906E
	INT                                                        = 0x1404
	CONSTANT_BORDER_HP                                         = 0x8151
	BUFFER                                                     = 0x82E0
	CON_8_ATI                                                  = 0x8949
	UNIFORM_OFFSET                                             = 0x8A3B
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	SCREEN_COORDINATES_REND                                    = 0x8490
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	DRAW_BUFFER8                                               = 0x882D
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	MULTISAMPLE_ARB                                            = 0x809D
	STREAM_DRAW                                                = 0x88E0
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	TEXTURE_3D_OES                                             = 0x806F
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	RED_MIN_CLAMP_INGR                                         = 0x8560
	INTERLACE_OML                                              = 0x8980
	DOUBLE_MAT2                                                = 0x8F46
	INT16_VEC4_NV                                              = 0x8FE7
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	BLEND_EQUATION                                             = 0x8009
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	ENABLE_BIT                                                 = 0x00002000
	EYE_PLANE                                                  = 0x2502
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	PRIMITIVES_GENERATED                                       = 0x8C87
	STENCIL_TEST                                               = 0x0B90
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	VIEW_CLASS_16_BITS                                         = 0x82CA
	TEXTURE19_ARB                                              = 0x84D3
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	MODULATE_ADD_ATI                                           = 0x8744
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	LINE_TO_NV                                                 = 0x04
	SGIX_flush_raster                                          = 1
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	SLUMINANCE                                                 = 0x8C46
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	PROJECTION_MATRIX                                          = 0x0BA7
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	GL_2X_BIT_ATI                                              = 0x00000001
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	UNSIGNED_BYTE                                              = 0x1401
	Q                                                          = 0x2003
	DEBUG_SOURCE_OTHER                                         = 0x824B
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	FOG_COORDINATE_SOURCE                                      = 0x8450
	FLOAT_R_NV                                                 = 0x8880
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	QUAD_ALPHA4_SGIS                                           = 0x811E
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	MIRRORED_REPEAT                                            = 0x8370
	TEXTURE28                                                  = 0x84DC
	CONSTANT                                                   = 0x8576
	SIGNED_INTENSITY8_NV                                       = 0x8708
	BLUE_BIT_ATI                                               = 0x00000004
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	COMPRESSED_RG                                              = 0x8226
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	OUTPUT_COLOR0_EXT                                          = 0x879B
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	CLIP_DISTANCE3                                             = 0x3003
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	TEXTURE_GEN_MODE                                           = 0x2500
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	TEXTURE20_ARB                                              = 0x84D4
	TEXTURE_SHADER_NV                                          = 0x86DE
	BUFFER_MAPPED_OES                                          = 0x88BC
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	DECAL                                                      = 0x2101
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	SOURCE2_RGB                                                = 0x8582
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	TEXTURE_1D                                                 = 0x0DE0
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	FLOAT16_NV                                                 = 0x8FF8
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	MATRIX5_ARB                                                = 0x88C5
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	FOG                                                        = 0x0B60
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	ALL_COMPLETED_NV                                           = 0x84F2
	COMBINE_EXT                                                = 0x8570
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	HALF_FLOAT_OES                                             = 0x8D61
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	BACK                                                       = 0x0405
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	IMAGE_BINDING_NAME                                         = 0x8F3A
	DRAW_BUFFER12_ARB                                          = 0x8831
	WRITE_ONLY                                                 = 0x88B9
	STENCIL_INDEX16_EXT                                        = 0x8D49
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	T2F_IUI_V3F_EXT                                            = 0x81B2
	RGBA32F                                                    = 0x8814
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	IMAGE_2D_ARRAY                                             = 0x9053
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	RED_SCALE                                                  = 0x0D14
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	REG_19_ATI                                                 = 0x8934
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	CUBIC_EXT                                                  = 0x8334
	RGBA32UI                                                   = 0x8D70
	ALPHA_TEST                                                 = 0x0BC0
	CLEAR                                                      = 0x1500
	CLIP_PLANE0                                                = 0x3000
	COLOR_INDEX12_EXT                                          = 0x80E6
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	TEXTURE27_ARB                                              = 0x84DB
	DECR_WRAP_EXT                                              = 0x8508
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	CON_28_ATI                                                 = 0x895D
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	FOG_END                                                    = 0x0B64
	LINE_SMOOTH_HINT                                           = 0x0C52
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	CON_10_ATI                                                 = 0x894B
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	PIXEL_MODE_BIT                                             = 0x00000020
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	COLOR_INDEX                                                = 0x1900
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	SGIS_point_parameters                                      = 1
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	COMPILE_AND_EXECUTE                                        = 0x1301
	LINEAR                                                     = 0x2601
	TRIANGLE_LIST_SUN                                          = 0x81D7
	COLOR_ARRAY_POINTER                                        = 0x8090
	LOGIC_OP_MODE                                              = 0x0BF0
	BLEND_COLOR                                                = 0x8005
	CONVOLUTION_WIDTH                                          = 0x8018
	COLOR_TABLE_SGI                                            = 0x80D0
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	MAX_LIST_NESTING                                           = 0x0B31
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	IMAGE_2D_RECT_EXT                                          = 0x904F
	MAP1_INDEX                                                 = 0x0D91
	PROJECTION                                                 = 0x1701
	POINT_SIZE_MAX_ARB                                         = 0x8127
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	MODELVIEW0_EXT                                             = 0x1700
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	PHONG_WIN                                                  = 0x80EA
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	CLIP_PLANE1                                                = 0x3001
	SUB_ATI                                                    = 0x8965
	FLOAT_MAT3_ARB                                             = 0x8B5B
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	MAX_EXT                                                    = 0x8008
	PROGRAM_TARGET_NV                                          = 0x8646
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	CLAMP_READ_COLOR                                           = 0x891C
	LUMINANCE16I_EXT                                           = 0x8D8C
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	RGB10_A2                                                   = 0x8059
	SAMPLER_1D                                                 = 0x8B5D
	TESSELLATION_MODE_AMD                                      = 0x9004
	DEBUG_TYPE_MARKER                                          = 0x8268
	STORAGE_CACHED_APPLE                                       = 0x85BE
	SQUARE_NV                                                  = 0x90A3
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	REFERENCE_PLANE_SGIX                                       = 0x817D
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	ADD_SIGNED                                                 = 0x8574
	CLIP_DISTANCE_NV                                           = 0x8C7A
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	SGIX_async                                                 = 1
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	TESS_GEN_MODE                                              = 0x8E76
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	INDEX_CLEAR_VALUE                                          = 0x0C20
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	CURRENT_TANGENT_EXT                                        = 0x843B
	NUM_PASSES_ATI                                             = 0x8970
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	SGIX_icc_texture                                           = 1
	GREEN_BITS                                                 = 0x0D53
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	COLOR_ATTACHMENT4                                          = 0x8CE4
	BLUE_INTEGER_EXT                                           = 0x8D96
	IMAGE_BUFFER_EXT                                           = 0x9051
	CONVEX_HULL_NV                                             = 0x908B
	MAP2_NORMAL                                                = 0x0DB2
	REPLACE_MIDDLE_SUN                                         = 0x0002
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	MATRIX14_ARB                                               = 0x88CE
	CON_14_ATI                                                 = 0x894F
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	LUMINANCE_SNORM                                            = 0x9011
	BOLD_BIT_NV                                                = 0x01
	ACTIVE_VARIABLES                                           = 0x9305
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	OP_ROUND_EXT                                               = 0x8790
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	COLOR_ATTACHMENT3                                          = 0x8CE3
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	MAP2_BINORMAL_EXT                                          = 0x8447
	MODELVIEW12_ARB                                            = 0x872C
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	RENDERBUFFER_WIDTH                                         = 0x8D42
	TEXTURE_HEIGHT                                             = 0x1001
	MATRIX4_NV                                                 = 0x8634
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	LOW_INT                                                    = 0x8DF3
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	STACK_UNDERFLOW                                            = 0x0504
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	COMBINE                                                    = 0x8570
	SOURCE2_RGB_ARB                                            = 0x8582
	OP_ADD_EXT                                                 = 0x8787
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	BLUE                                                       = 0x1905
	MODELVIEW19_ARB                                            = 0x8733
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	MAX_LIGHTS                                                 = 0x0D31
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	COLOR_RENDERABLE                                           = 0x8286
	SIGNED_RGBA_NV                                             = 0x86FB
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	FOG_COORD_SRC                                              = 0x8450
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	DRAW_BUFFER0_ATI                                           = 0x8825
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	DEPTH_EXT                                                  = 0x1801
	PROGRAM_BINDING_ARB                                        = 0x8677
	PRESERVE_ATI                                               = 0x8762
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	PROGRAM_SEPARABLE                                          = 0x8258
	OPERAND0_ALPHA_EXT                                         = 0x8598
	IMAGE_3D_EXT                                               = 0x904E
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	RG32UI                                                     = 0x823C
	DYNAMIC_COPY_ARB                                           = 0x88EA
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	SET                                                        = 0x150F
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	INVALID_INDEX                                              = 0xFFFFFFFF
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	MAP1_TANGENT_EXT                                           = 0x8444
	VECTOR_EXT                                                 = 0x87BF
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	UNIFORM_TYPE                                               = 0x8A37
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	VERTEX23_BIT_PGI                                           = 0x00000004
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	BLOCK_INDEX                                                = 0x92FD
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	DRAW_BUFFER15_ARB                                          = 0x8834
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	MEDIUM_FLOAT                                               = 0x8DF1
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	SAMPLES                                                    = 0x80A9
	EYE_POINT_SGIS                                             = 0x81F4
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	READ_PIXELS                                                = 0x828C
	SOURCE2_ALPHA_ARB                                          = 0x858A
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	RGBA8_OES                                                  = 0x8058
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	DEPTH_STENCIL_MESA                                         = 0x8750
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	COLOR_ARRAY_TYPE                                           = 0x8082
	TEXTURE_MAX_LEVEL                                          = 0x813D
	BLEND_COLOR_EXT                                            = 0x8005
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	RG8_EXT                                                    = 0x822B
	SIGNED_HILO_NV                                             = 0x86F9
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	MITER_TRUNCATE_NV                                          = 0x90A8
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	EXT_blend_minmax                                           = 1
	CLIP_DISTANCE7                                             = 0x3007
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	RGB_SNORM                                                  = 0x8F92
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	COLOR_INDEX4_EXT                                           = 0x80E4
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	TEXTURE14_ARB                                              = 0x84CE
	RGBA16F_ARB                                                = 0x881A
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	SGIX_fragment_lighting                                     = 1
	COLOR_ARRAY_STRIDE                                         = 0x8083
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	DYNAMIC_ATI                                                = 0x8761
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	SAMPLE_BUFFERS                                             = 0x80A8
	RENDERBUFFER_OES                                           = 0x8D41
	LOW_FLOAT                                                  = 0x8DF0
	ALPHA8_SNORM                                               = 0x9014
	MULTISAMPLE_BIT                                            = 0x20000000
	GREATER                                                    = 0x0204
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	INT_SAMPLER_2D                                             = 0x8DCA
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	DRAW_BUFFER_EXT                                            = 0x0C01
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	PRIMITIVE_RESTART                                          = 0x8F9D
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	FULL_RANGE_EXT                                             = 0x87E1
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	STENCIL_INDEX8_OES                                         = 0x8D48
	SAMPLER_OBJECT_AMD                                         = 0x9155
	ALWAYS                                                     = 0x0207
	Z_EXT                                                      = 0x87D7
	FLOAT_MAT2_ARB                                             = 0x8B5A
	INT16_VEC2_NV                                              = 0x8FE5
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	RG8UI                                                      = 0x8238
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	FRONT_FACE                                                 = 0x0B46
	CONVOLUTION_2D_EXT                                         = 0x8011
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	CON_4_ATI                                                  = 0x8945
	C4UB_V2F                                                   = 0x2A22
	RGB4_EXT                                                   = 0x804F
	RGB_INTEGER_EXT                                            = 0x8D98
	RGBA8_EXT                                                  = 0x8058
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	TEXTURE31_ARB                                              = 0x84DF
	SMOOTH                                                     = 0x1D01
	LIGHT3                                                     = 0x4003
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	DOUBLE_MAT4x3                                              = 0x8F4E
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	FENCE_APPLE                                                = 0x8A0B
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	UNSIGNED_SHORT                                             = 0x1403
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	DRAW_BUFFER6_NV                                            = 0x882B
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	DOUBLEBUFFER                                               = 0x0C32
	SURFACE_STATE_NV                                           = 0x86EB
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	YCBCR_MESA                                                 = 0x8757
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	BLEND_SRC_RGB_OES                                          = 0x80C9
	RG16I                                                      = 0x8239
	FRAMEBUFFER_BLEND                                          = 0x828B
	DRAW_BUFFER11_ARB                                          = 0x8830
	FLOAT_RGBA16_NV                                            = 0x888A
	REDUCE                                                     = 0x8016
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	SGIX_ycrcb                                                 = 1
	RGB10                                                      = 0x8052
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	QUERY_RESULT_EXT                                           = 0x8866
	REG_5_ATI                                                  = 0x8926
	SLUMINANCE_EXT                                             = 0x8C46
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	HISTOGRAM_RED_SIZE                                         = 0x8028
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	GENERATE_MIPMAP_HINT                                       = 0x8192
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	INVERT                                                     = 0x150A
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	FIELD_LOWER_NV                                             = 0x9023
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	COMPUTE_SHADER_BIT                                         = 0x00000020
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	RED_INTEGER_EXT                                            = 0x8D94
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	TEXTURE6_ARB                                               = 0x84C6
	DRAW_BUFFER3_NV                                            = 0x8828
	FRONT_LEFT                                                 = 0x0400
	POLYGON_SMOOTH                                             = 0x0B41
	COLOR_MATRIX_SGI                                           = 0x80B1
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	PRIMARY_COLOR_ARB                                          = 0x8577
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	DRAW_BUFFER4_ARB                                           = 0x8829
	DRAW_BUFFER0_NV                                            = 0x8825
	MAX_SAMPLES_NV                                             = 0x8D57
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	COLOR_TABLE_BIAS                                           = 0x80D7
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	SGI_color_matrix                                           = 1
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	DISPLAY_LIST                                               = 0x82E7
	EMBOSS_MAP_NV                                              = 0x855F
	SOURCE1_ALPHA_ARB                                          = 0x8589
	VOLATILE_APPLE                                             = 0x8A1A
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	TEXTURE_GEN_T                                              = 0x0C61
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	COLOR_CLEAR_VALUE                                          = 0x0C22
	RGB                                                        = 0x1907
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	CURRENT_VERTEX_EXT                                         = 0x87E2
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	SOURCE0_RGB_ARB                                            = 0x8580
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	VERTEX_ARRAY_POINTER                                       = 0x808E
	RGB5_EXT                                                   = 0x8050
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	CULL_VERTEX_IBM                                            = 103050
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	INDEX_SHIFT                                                = 0x0D12
	SYNC_CL_EVENT_ARB                                          = 0x8240
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	LUMINANCE12                                                = 0x8041
	MODELVIEW5_ARB                                             = 0x8725
	ARRAY_BUFFER_ARB                                           = 0x8892
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	REG_23_ATI                                                 = 0x8938
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	LUMINANCE32UI_EXT                                          = 0x8D74
	CLEAR_BUFFER                                               = 0x82B4
	COLOR_ATTACHMENT15                                         = 0x8CEF
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	REDUCE_EXT                                                 = 0x8016
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	RED_INTEGER                                                = 0x8D94
	TIMEOUT_EXPIRED                                            = 0x911B
	TEXTURE_MIN_FILTER                                         = 0x2801
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	SGIX_polynomial_ffd                                        = 1
	RENDER                                                     = 0x1C00
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	TEXTURE23_ARB                                              = 0x84D7
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	MAX_VERTEX_STREAMS                                         = 0x8E71
	SGIX_texture_multi_buffer                                  = 1
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	VERSION_1_5                                                = 1
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	SIGNED_ALPHA_NV                                            = 0x8705
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	MATRIX25_ARB                                               = 0x88D9
	COMPRESSED_SRGB                                            = 0x8C48
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	LAYER_NV                                                   = 0x8DAA
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	INT_IMAGE_3D                                               = 0x9059
	FOG_BIT                                                    = 0x00000080
	OUT_OF_MEMORY                                              = 0x0505
	FILTER                                                     = 0x829A
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	EYE_LINE_SGIS                                              = 0x81F6
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	UNSIGNED_NORMALIZED                                        = 0x8C17
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	COMMAND_BARRIER_BIT                                        = 0x00000040
	FIXED                                                      = 0x140C
	SPRITE_SGIX                                                = 0x8148
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	OP_CLAMP_EXT                                               = 0x878E
	UNSIGNALED_APPLE                                           = 0x9118
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	GL_2PASS_0_EXT                                             = 0x80A2
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	SLIM8U_SGIX                                                = 0x831D
	DEPTH_STENCIL_OES                                          = 0x84F9
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	COLOR_COMPONENTS                                           = 0x8283
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	AVERAGE_HP                                                 = 0x8160
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	BUFFER_ACCESS_OES                                          = 0x88BB
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	SCISSOR_TEST                                               = 0x0C11
	TEXTURE_BLUE_SIZE                                          = 0x805E
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	SOURCE3_ALPHA_NV                                           = 0x858B
	INT16_VEC3_NV                                              = 0x8FE6
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	SAMPLES_EXT                                                = 0x80A9
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	DRAW_BUFFER15_NV                                           = 0x8834
	MATRIX23_ARB                                               = 0x88D7
	SAMPLER_3D_OES                                             = 0x8B5F
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	ALPHA16I_EXT                                               = 0x8D8A
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	TEXTURE7_ARB                                               = 0x84C7
	VARIABLE_F_NV                                              = 0x8528
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	POINT_SMOOTH                                               = 0x0B10
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	OPERAND0_RGB_ARB                                           = 0x8590
	VERTEX_STREAM1_ATI                                         = 0x876D
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	GL_422_AVERAGE_EXT                                         = 0x80CE
	QUERY_WAIT_NV                                              = 0x8E13
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	SCISSOR_BOX                                                = 0x0C10
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	FRAGMENT_TEXTURE                                           = 0x829F
	TEXTURE16_ARB                                              = 0x84D0
	DRAW_BUFFER15                                              = 0x8834
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	BUFFER_USAGE                                               = 0x8765
	BUFFER_MAPPED                                              = 0x88BC
	MATRIX0_ARB                                                = 0x88C0
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	POINT_SIZE_RANGE                                           = 0x0B12
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	SAMPLER_CUBE                                               = 0x8B60
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	QUADRATIC_ATTENUATION                                      = 0x1209
	RGB5_A1                                                    = 0x8057
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	COMBINE_ALPHA_ARB                                          = 0x8572
	CURRENT_MATRIX_ARB                                         = 0x8641
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	REG_15_ATI                                                 = 0x8930
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	MATRIX8_ARB                                                = 0x88C8
	SRGB_EXT                                                   = 0x8C40
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	RGBA4_OES                                                  = 0x8056
	COLOR_ARRAY_EXT                                            = 0x8076
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	STREAM_READ_ARB                                            = 0x88E1
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	REFLECTION_MAP_OES                                         = 0x8512
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	PERTURB_EXT                                                = 0x85AE
	BUMP_ENVMAP_ATI                                            = 0x877B
	ALPHA8                                                     = 0x803C
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	SPARE0_NV                                                  = 0x852E
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	REGISTER_COMBINERS_NV                                      = 0x8522
	NEGATE_BIT_ATI                                             = 0x00000004
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	UNIFORM_BLOCK                                              = 0x92E2
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	IMAGE_CUBE                                                 = 0x9050
	LUMINANCE16                                                = 0x8042
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	VIEW_CLASS_96_BITS                                         = 0x82C5
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	CON_19_ATI                                                 = 0x8954
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	TEXTURE10_ARB                                              = 0x84CA
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	TEXTURE9_ARB                                               = 0x84C9
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	OP_EXP_BASE_2_EXT                                          = 0x8791
	REG_12_ATI                                                 = 0x892D
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	INT_IMAGE_2D_RECT                                          = 0x905A
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	TEXTURE18                                                  = 0x84D2
	DOT3_RGBA_EXT                                              = 0x8741
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	SKIP_COMPONENTS3_NV                                        = -4
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	SRC1_RGB                                                   = 0x8581
	QUAD_MESH_SUN                                              = 0x8614
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	SIGNED_RGBA8_NV                                            = 0x86FC
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	RENDERBUFFER_BINDING                                       = 0x8CA7
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	RGB8UI                                                     = 0x8D7D
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	REG_13_ATI                                                 = 0x892E
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	TEXTURE_BINDING_1D                                         = 0x8068
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	ALPHA_FLOAT16_ATI                                          = 0x881C
	DRAW_BUFFER3_ATI                                           = 0x8828
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	VARIABLE_D_NV                                              = 0x8526
	BUFFER_MAPPED_ARB                                          = 0x88BC
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	TESS_CONTROL_SHADER                                        = 0x8E88
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	SOURCE2_ALPHA                                              = 0x858A
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	RELEASED_APPLE                                             = 0x8A19
	SGIS_texture_edge_clamp                                    = 1
	CONVOLUTION_HEIGHT                                         = 0x8019
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	R16                                                        = 0x822A
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	EXPAND_NORMAL_NV                                           = 0x8538
	PALETTE8_RGBA8_OES                                         = 0x8B96
	COLOR_ATTACHMENT0                                          = 0x8CE0
	PROGRAM_INPUT                                              = 0x92E3
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	ALPHA_BITS                                                 = 0x0D55
	DEPTH_COMPONENT32                                          = 0x81A7
	FRAMEBUFFER                                                = 0x8D40
	INT_SAMPLER_3D                                             = 0x8DCB
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	TRIANGLES                                                  = 0x0004
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	ALPHA12                                                    = 0x803D
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	VERTEX_STREAM6_ATI                                         = 0x8772
	DRAW_BUFFER13                                              = 0x8832
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	POLYGON_STIPPLE                                            = 0x0B42
	COLOR_EXT                                                  = 0x1800
	MINMAX                                                     = 0x802E
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	SIGNED_RGB8_NV                                             = 0x86FF
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	CON_6_ATI                                                  = 0x8947
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	TEXTURE_GREEN_SIZE                                         = 0x805D
	DEPTH_COMPONENT24                                          = 0x81A6
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	FLOAT_RGB32_NV                                             = 0x8889
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	INDEX_WRITEMASK                                            = 0x0C21
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	RGBA4_DXT5_S3TC                                            = 0x83A5
	COMBINER2_NV                                               = 0x8552
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	CUBIC_CURVE_TO_NV                                          = 0x0C
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	EXT_point_parameters                                       = 1
	MAP_READ_BIT                                               = 0x0001
	DSDT8_NV                                                   = 0x8709
	MAD_ATI                                                    = 0x8968
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	PATH_END_CAPS_NV                                           = 0x9076
	OP_FLOOR_EXT                                               = 0x878F
	MAX_SAMPLES_EXT                                            = 0x8D57
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	TEXTURE_MAG_FILTER                                         = 0x2800
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	ITALIC_BIT_NV                                              = 0x02
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	GEOMETRY_TEXTURE                                           = 0x829E
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	FLOAT_VEC2                                                 = 0x8B50
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	NORMAL_ARRAY_TYPE                                          = 0x807E
	LIGHT4                                                     = 0x4004
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	HIGH_FLOAT                                                 = 0x8DF2
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	STENCIL_REF                                                = 0x0B97
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	INT8_VEC2_NV                                               = 0x8FE1
	BLEND_DST_RGB                                              = 0x80C8
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	HILO8_NV                                                   = 0x885E
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	INT_SAMPLER_1D                                             = 0x8DC9
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	DOT3_RGB_ARB                                               = 0x86AE
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	TEXTURE14                                                  = 0x84CE
	Z400_BINARY_AMD                                            = 0x8740
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	MAP2_INDEX                                                 = 0x0DB1
	PACK_SKIP_IMAGES                                           = 0x806B
	FRAGMENT_SHADER_ATI                                        = 0x8920
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	QUERY_NO_WAIT                                              = 0x8E14
	SRC_ALPHA_SATURATE                                         = 0x0308
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	RG_EXT                                                     = 0x8227
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	NO_ERROR                                                   = 0
	SHADER_OBJECT_ARB                                          = 0x8B48
	PERCENTAGE_AMD                                             = 0x8BC3
	REG_0_ATI                                                  = 0x8921
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	DEBUG_PRINT_MESA                                           = 0x875A
	RGBA_INTEGER_EXT                                           = 0x8D99
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	ARRAY_SIZE                                                 = 0x92FB
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	TEXTURE25                                                  = 0x84D9
	X_EXT                                                      = 0x87D5
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	INDEX_MODE                                                 = 0x0C30
	DEPTH_SCALE                                                = 0x0D1E
	RGB8_EXT                                                   = 0x8051
	SCALE_BY_FOUR_NV                                           = 0x853F
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	TEXTURE                                                    = 0x1702
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	MODELVIEW29_ARB                                            = 0x873D
	BUMP_TARGET_ATI                                            = 0x877C
	QUERY_RESULT                                               = 0x8866
	RGBA16_SNORM                                               = 0x8F9B
	INT8_VEC4_NV                                               = 0x8FE3
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	DRAW_BUFFER12                                              = 0x8831
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	PATH_DASH_OFFSET_NV                                        = 0x907E
	SGIS_generate_mipmap                                       = 1
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	INTENSITY8UI_EXT                                           = 0x8D7F
	DOMAIN                                                     = 0x0A02
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	SINGLE_COLOR                                               = 0x81F9
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	NORMAL_MAP_EXT                                             = 0x8511
	UNSIGNED_INT16_NV                                          = 0x8FF0
	COUNT_UP_NV                                                = 0x9088
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	YCRCB_SGIX                                                 = 0x8318
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	DYNAMIC_READ_ARB                                           = 0x88E9
	CON_13_ATI                                                 = 0x894E
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	BLEND_EQUATION_RGB                                         = 0x8009
	DUAL_ALPHA4_SGIS                                           = 0x8110
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	WRITE_ONLY_ARB                                             = 0x88B9
	AFFINE_3D_NV                                               = 0x9094
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	FLOAT_VEC3_ARB                                             = 0x8B51
	RGBA16I_EXT                                                = 0x8D88
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	SGIX_pixel_tiles                                           = 1
	KEEP                                                       = 0x1E00
	RGB_FLOAT32_ATI                                            = 0x8815
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	MVP_MATRIX_EXT                                             = 0x87E3
	SRGB8_EXT                                                  = 0x8C41
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	SGIX_instruments                                           = 1
	CLAMP_TO_BORDER_NV                                         = 0x812D
	LINE_STRIP_ADJACENCY                                       = 0x000B
	RGBA_S3TC                                                  = 0x83A2
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	MAP1_BINORMAL_EXT                                          = 0x8446
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	AUX1                                                       = 0x040A
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	DEPTH24_STENCIL8                                           = 0x88F0
	RGBA32UI_EXT                                               = 0x8D70
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	TEXTURE_COMPONENTS                                         = 0x1003
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	TEXTURE_BASE_LEVEL                                         = 0x813C
	VARIANT_ARRAY_EXT                                          = 0x87E8
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	RGB16                                                      = 0x8054
	FULL_SUPPORT                                               = 0x82B7
	VERTEX_STREAM5_ATI                                         = 0x8771
	OP_LOG_BASE_2_EXT                                          = 0x8792
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	TEXTURE16                                                  = 0x84D0
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	BUFFER_BINDING                                             = 0x9302
	CMYK_EXT                                                   = 0x800C
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	EYE_RADIAL_NV                                              = 0x855B
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	MAP_READ_BIT_EXT                                           = 0x0001
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	MAP_WRITE_BIT_EXT                                          = 0x0002
	R                                                          = 0x2002
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	VERTICAL_LINE_TO_NV                                        = 0x08
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	SELECT                                                     = 0x1C02
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	SRGB_READ                                                  = 0x8297
	LUMINANCE16F_ARB                                           = 0x881E
	REG_2_ATI                                                  = 0x8923
	INT_10_10_10_2_OES                                         = 0x8DF7
	OFFSET                                                     = 0x92FC
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	STENCIL_EXT                                                = 0x1802
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	CON_24_ATI                                                 = 0x8959
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	RGB32UI_EXT                                                = 0x8D71
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	CURRENT_QUERY_ARB                                          = 0x8865
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	ALPHA16                                                    = 0x803E
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	COLOR_INDEX8_EXT                                           = 0x80E5
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	DOUBLE_MAT2_EXT                                            = 0x8F46
	RGBA_SNORM                                                 = 0x8F93
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	DRAW_BUFFER                                                = 0x0C01
	SIGNED_LUMINANCE_NV                                        = 0x8701
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	QUERY_WAIT                                                 = 0x8E13
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	EXT_blend_color                                            = 1
	MAGNITUDE_BIAS_NV                                          = 0x8718
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	TRANSLATE_3D_NV                                            = 0x9091
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	NAME_STACK_DEPTH                                           = 0x0D70
	MATRIX11_ARB                                               = 0x88CB
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	EVAL_BIT                                                   = 0x00010000
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	SUBPIXEL_BITS                                              = 0x0D50
	BLEND_DST_RGB_OES                                          = 0x80C8
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	REG_3_ATI                                                  = 0x8924
	VERSION_2_1                                                = 1
	COPY_PIXEL_TOKEN                                           = 0x0706
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	CONST_EYE_NV                                               = 0x86E5
	FLOAT_R32_NV                                               = 0x8885
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	NEAREST                                                    = 0x2600
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	SWIZZLE_STRQ_ATI                                           = 0x897A
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	CONDITION_SATISFIED                                        = 0x911C
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	COLOR_TABLE                                                = 0x80D0
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	DOUBLE_MAT2x4                                              = 0x8F4A
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	TEXTURE4_ARB                                               = 0x84C4
	MATRIX_PALETTE_ARB                                         = 0x8840
	SGIX_texture_add_env                                       = 1
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	EXT_texture_object                                         = 1
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	VERSION_1_2                                                = 1
	INDEX_ARRAY_EXT                                            = 0x8077
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	BIAS_BIT_ATI                                               = 0x00000008
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	MINMAX_SINK_EXT                                            = 0x8030
	RGBA2_EXT                                                  = 0x8055
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	INTERLACE_READ_INGR                                        = 0x8568
	SRC2_ALPHA                                                 = 0x858A
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	DOT4_ATI                                                   = 0x8967
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	RGBA12_EXT                                                 = 0x805A
	INDEX                                                      = 0x8222
	MAX_LABEL_LENGTH                                           = 0x82E8
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	SGIS_multisample                                           = 1
	MODELVIEW13_ARB                                            = 0x872D
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	SCISSOR_BIT                                                = 0x00080000
	BLEND                                                      = 0x0BE2
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	DOUBLE                                                     = 0x140A
	R32UI                                                      = 0x8236
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	UNSIGNED_INVERT_NV                                         = 0x8537
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	COMBINER7_NV                                               = 0x8557
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	CULL_MODES_NV                                              = 0x86E0
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	DEPTH_COMPONENT16                                          = 0x81A5
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	TEXTURE3_ARB                                               = 0x84C3
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	FRAME_NV                                                   = 0x8E26
	LUMINANCE4_EXT                                             = 0x803F
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	OP_INDEX_EXT                                               = 0x8782
	RGB8UI_EXT                                                 = 0x8D7D
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	DRAW_BUFFER5_ATI                                           = 0x882A
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	DRAW_BUFFER14_ARB                                          = 0x8833
	SLUMINANCE_ALPHA                                           = 0x8C44
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	ALPHA8_EXT                                                 = 0x803C
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	CURRENT_TIME_NV                                            = 0x8E28
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	TEXTURE_LIGHT_EXT                                          = 0x8350
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	MATRIX5_NV                                                 = 0x8635
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	STORAGE_SHARED_APPLE                                       = 0x85BF
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	RGB16I                                                     = 0x8D89
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	FIRST_TO_REST_NV                                           = 0x90AF
	SGIS_fog_function                                          = 1
	VIEW_CLASS_128_BITS                                        = 0x82C4
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	INDEX_LOGIC_OP                                             = 0x0BF1
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	TEXTURE19                                                  = 0x84D3
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	TEXTURE24                                                  = 0x84D8
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	CONSTANT_ALPHA                                             = 0x8003
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	REPLACE_VALUE_AMD                                          = 0x874B
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	LIGHT5                                                     = 0x4005
	VERTEX_STREAM0_ATI                                         = 0x876C
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	CCW                                                        = 0x0901
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	SOURCE1_RGB_ARB                                            = 0x8581
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	FLOAT_RG32_NV                                              = 0x8887
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	NUM_SAMPLE_COUNTS                                          = 0x9380
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	TEXTURE_BORDER_COLOR                                       = 0x1004
	TEXTURE11                                                  = 0x84CB
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	EXT_vertex_array                                           = 1
	PASS_THROUGH_TOKEN                                         = 0x0700
	FOG_COLOR                                                  = 0x0B66
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	ALPHA16UI_EXT                                              = 0x8D78
	INT_SAMPLER_CUBE                                           = 0x8DCC
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	MAP1_VERTEX_3                                              = 0x0D97
	DRAW_BUFFER5_NV                                            = 0x882A
	MATRIX10_ARB                                               = 0x88CA
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	CONSTANT_ATTENUATION                                       = 0x1207
	HILO_NV                                                    = 0x86F4
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	COLOR_ARRAY_LIST_IBM                                       = 103072
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
)

type Context struct {
	access                    sync.Mutex
	context                   *C.gl12Context
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
	glc.context = C.gl12NewContext()

	glc.Accum = func(op uint32, value float32) {
		defer glc.trace("Accum")
		C.gl12Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		defer glc.trace("AlphaFunc")
		C.gl12AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		defer glc.trace("Begin")
		glc.inBeginEnd = true
		C.gl12Begin(glc.context, C.GLenum(mode))
		return
	}

	glc.End = func() {
		defer glc.trace("End")
		C.gl12End(glc.context)
		glc.inBeginEnd = false
		return
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8) {
		defer glc.trace("Bitmap")
		C.gl12Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		defer glc.trace("BlendFunc")
		C.gl12BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		defer glc.trace("CallList")
		C.gl12CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		defer glc.trace("CallLists")
		C.gl12CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		defer glc.trace("Clear")
		C.gl12Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		defer glc.trace("ClearAccum")
		C.gl12ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		defer glc.trace("ClearColor")
		C.gl12ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		defer glc.trace("ClearDepth")
		C.gl12ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		defer glc.trace("ClearIndex")
		C.gl12ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		defer glc.trace("ClearStencil")
		C.gl12ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation *float64) {
		defer glc.trace("ClipPlane")
		C.gl12ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.Color3b = func(red, green, blue int8) {
		defer glc.trace("Color3b")
		C.gl12Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		defer glc.trace("Color3d")
		C.gl12Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		defer glc.trace("Color3f")
		C.gl12Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		defer glc.trace("Color3i")
		C.gl12Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		defer glc.trace("Color3s")
		C.gl12Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		defer glc.trace("Color3ub")
		C.gl12Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		defer glc.trace("Color3ui")
		C.gl12Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		defer glc.trace("Color3us")
		C.gl12Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		defer glc.trace("Color4b")
		C.gl12Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		defer glc.trace("Color4d")
		C.gl12Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		defer glc.trace("Color4f")
		C.gl12Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		defer glc.trace("Color4i")
		C.gl12Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		defer glc.trace("Color4s")
		C.gl12Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		defer glc.trace("Color4ub")
		C.gl12Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		defer glc.trace("Color4ui")
		C.gl12Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		defer glc.trace("Color4us")
		C.gl12Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v *int8) {
		defer glc.trace("Color3bv")
		C.gl12Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color3dv = func(v *float64) {
		defer glc.trace("Color3dv")
		C.gl12Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color3fv = func(v *float32) {
		defer glc.trace("Color3fv")
		C.gl12Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color3iv = func(v *int32) {
		defer glc.trace("Color3iv")
		C.gl12Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color3sv = func(v *int16) {
		defer glc.trace("Color3sv")
		C.gl12Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color3ubv = func(v *uint8) {
		defer glc.trace("Color3ubv")
		C.gl12Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color3uiv = func(v *uint32) {
		defer glc.trace("Color3uiv")
		C.gl12Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color3usv = func(v *uint16) {
		defer glc.trace("Color3usv")
		C.gl12Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.Color4bv = func(v *int8) {
		defer glc.trace("Color4bv")
		C.gl12Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color4dv = func(v *float64) {
		defer glc.trace("Color4dv")
		C.gl12Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color4fv = func(v *float32) {
		defer glc.trace("Color4fv")
		C.gl12Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color4iv = func(v *int32) {
		defer glc.trace("Color4iv")
		C.gl12Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color4sv = func(v *int16) {
		defer glc.trace("Color4sv")
		C.gl12Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color4ubv = func(v *uint8) {
		defer glc.trace("Color4ubv")
		C.gl12Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color4uiv = func(v *uint32) {
		defer glc.trace("Color4uiv")
		C.gl12Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color4usv = func(v *uint16) {
		defer glc.trace("Color4usv")
		C.gl12Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		defer glc.trace("ColorMask")
		C.gl12ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		defer glc.trace("ColorMaterial")
		C.gl12ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		defer glc.trace("CopyPixels")
		C.gl12CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		defer glc.trace("CullFace")
		C.gl12CullFace(glc.context, C.GLenum(mode))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		defer glc.trace("DeleteLists")
		C.gl12DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		defer glc.trace("DepthFunc")
		C.gl12DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthMask = func(flag bool) {
		defer glc.trace("DepthMask")
		C.gl12DepthMask(glc.context, boolToGL(flag))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		defer glc.trace("DepthRange")
		C.gl12DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		defer glc.trace("Enable")
		C.gl12Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		defer glc.trace("Disable")
		C.gl12Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		defer glc.trace("DrawBuffer")
		C.gl12DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("DrawPixels")
		C.gl12DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.EdgeFlag = func(flag bool) {
		defer glc.trace("EdgeFlag")
		C.gl12EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag *bool) {
		defer glc.trace("EdgeFlagv")
		C.gl12EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(flag)))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		defer glc.trace("EdgeFlagPointer")
		C.gl12EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EvalCoord1d = func(u float64) {
		defer glc.trace("EvalCoord1d")
		C.gl12EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		defer glc.trace("EvalCoord1f")
		C.gl12EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		defer glc.trace("EvalCoord2d")
		C.gl12EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		defer glc.trace("EvalCoord2f")
		C.gl12EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u *float64) {
		defer glc.trace("EvalCoord1dv")
		C.gl12EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord1fv = func(u *float32) {
		defer glc.trace("EvalCoord1fv")
		C.gl12EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2dv = func(u *float64) {
		defer glc.trace("EvalCoord2dv")
		C.gl12EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2fv = func(u *float32) {
		defer glc.trace("EvalCoord2fv")
		C.gl12EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		defer glc.trace("EvalMesh1")
		C.gl12EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		defer glc.trace("EvalMesh2")
		C.gl12EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		defer glc.trace("EvalPoint1")
		C.gl12EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		defer glc.trace("EvalPoint2")
		C.gl12EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer *float32) {
		defer glc.trace("FeedbackBuffer")
		C.gl12FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(buffer)))
	}

	glc.Finish = func() {
		defer glc.trace("Finish")
		C.gl12Finish(glc.context)
	}

	glc.Flush = func() {
		defer glc.trace("Flush")
		C.gl12Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		defer glc.trace("Fogf")
		C.gl12Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		defer glc.trace("Fogi")
		C.gl12Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params *float32) {
		defer glc.trace("Fogfv")
		C.gl12Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Fogiv = func(pname uint32, params *int32) {
		defer glc.trace("Fogiv")
		C.gl12Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.FrontFace = func(mode uint32) {
		defer glc.trace("FrontFace")
		C.gl12FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		defer glc.trace("Frustum")
		C.gl12Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		defer glc.trace("GenLists")
		return uint32(C.gl12GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params *bool) {
		defer glc.trace("GetBooleanv")
		C.gl12GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(params)))
	}

	glc.GetDoublev = func(pname uint32, params *float64) {
		defer glc.trace("GetDoublev")
		C.gl12GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetFloatv = func(pname uint32, params *float32) {
		defer glc.trace("GetFloatv")
		C.gl12GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetIntegerv = func(pname uint32, params *int32) {
		defer glc.trace("GetIntegerv")
		C.gl12GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetClipPlane = func(plane uint32, equation *float64) {
		defer glc.trace("GetClipPlane")
		C.gl12GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl12GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params *float32) {
		defer glc.trace("GetLightfv")
		C.gl12GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetLightiv = func(light, pname uint32, params *int32) {
		defer glc.trace("GetLightiv")
		C.gl12GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetMapdv = func(target, query uint32, v *float64) {
		defer glc.trace("GetMapdv")
		C.gl12GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.GetMapfv = func(target, query uint32, v *float32) {
		defer glc.trace("GetMapfv")
		C.gl12GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.GetMapiv = func(target, query uint32, v *int32) {
		defer glc.trace("GetMapiv")
		C.gl12GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.GetMaterialfv = func(face, pname uint32, params *float32) {
		defer glc.trace("GetMaterialfv")
		C.gl12GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetMaterialiv = func(face, pname uint32, params *int32) {
		defer glc.trace("GetMaterialiv")
		C.gl12GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetPixelMapfv = func(Map uint32, values *float32) {
		defer glc.trace("GetPixelMapfv")
		C.gl12GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapuiv = func(Map uint32, values *uint32) {
		defer glc.trace("GetPixelMapuiv")
		C.gl12GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapusv = func(Map uint32, values *uint16) {
		defer glc.trace("GetPixelMapusv")
		C.gl12GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.GetPolygonStipple = func(pattern *uint8) {
		defer glc.trace("GetPolygonStipple")
		C.gl12GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(pattern)))
	}

	glc.GetString = func(name uint32) string {
		defer glc.trace("GetString")
		cstr := C.gl12GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetTexEnvfv")
		C.gl12GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexEnviv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetTexEnviv")
		C.gl12GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexGendv = func(coord, pname uint32, params *float64) {
		defer glc.trace("GetTexGendv")
		C.gl12GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params *float32) {
		defer glc.trace("GetTexGenfv")
		C.gl12GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params *int32) {
		defer glc.trace("GetTexGeniv")
		C.gl12GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("GetTexImage")
		C.gl12GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params *float32) {
		defer glc.trace("GetTexLevelParameterfv")
		C.gl12GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params *int32) {
		defer glc.trace("GetTexLevelParameteriv")
		C.gl12GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetTexParameterfv")
		C.gl12GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetTexParameteriv")
		C.gl12GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Hint = func(target, mode uint32) {
		defer glc.trace("Hint")
		C.gl12Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		defer glc.trace("Indexd")
		C.gl12Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		defer glc.trace("Indexf")
		C.gl12Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		defer glc.trace("Indexi")
		C.gl12Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		defer glc.trace("Indexs")
		C.gl12Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexdv = func(c *float64) {
		defer glc.trace("Indexdv")
		C.gl12Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(c)))
	}

	glc.Indexfv = func(c *float32) {
		defer glc.trace("Indexfv")
		C.gl12Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(c)))
	}

	glc.Indexiv = func(c *int32) {
		defer glc.trace("Indexiv")
		C.gl12Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(c)))
	}

	glc.Indexsv = func(c *int16) {
		defer glc.trace("Indexsv")
		C.gl12Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(c)))
	}

	glc.IndexMask = func(mask uint32) {
		defer glc.trace("IndexMask")
		C.gl12IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("IndexPointer")
		C.gl12IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		defer glc.trace("InitNames")
		C.gl12InitNames(glc.context)
	}

	glc.IsEnabled = func(cap uint32) {
		defer glc.trace("IsEnabled")
		C.gl12IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		defer glc.trace("IsList")
		return C.gl12IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		defer glc.trace("Lightf")
		C.gl12Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		defer glc.trace("Lighti")
		C.gl12Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params *float32) {
		defer glc.trace("Lightfv")
		C.gl12Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Lightiv = func(light, pname uint32, params *int32) {
		defer glc.trace("Lightiv")
		C.gl12Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		defer glc.trace("LightModelf")
		C.gl12LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		defer glc.trace("LightModeli")
		C.gl12LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params *float32) {
		defer glc.trace("LightModelfv")
		C.gl12LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.LightModeliv = func(pname uint32, params *int32) {
		defer glc.trace("LightModeliv")
		C.gl12LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		defer glc.trace("LineStipple")
		C.gl12LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		defer glc.trace("LineWidth")
		C.gl12LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		defer glc.trace("ListBase")
		C.gl12ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		defer glc.trace("LoadIdentity")
		C.gl12LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m *float64) {
		defer glc.trace("LoadMatrixd")
		C.gl12LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadMatrixf = func(m *float32) {
		defer glc.trace("LoadMatrixf")
		C.gl12LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.LoadName = func(name uint32) {
		defer glc.trace("LoadName")
		C.gl12LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		defer glc.trace("LogicOp")
		C.gl12LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points *float64) {
		defer glc.trace("Map1d")
		C.gl12Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points *float32) {
		defer glc.trace("Map1f")
		C.gl12Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64) {
		defer glc.trace("Map2d")
		C.gl12Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32) {
		defer glc.trace("Map2f")
		C.gl12Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		defer glc.trace("MapGrid1d")
		C.gl12MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		defer glc.trace("MapGrid1f")
		C.gl12MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		defer glc.trace("MapGrid2d")
		C.gl12MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		defer glc.trace("MapGrid2f")
		C.gl12MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		defer glc.trace("Materialf")
		C.gl12Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		defer glc.trace("Materiali")
		C.gl12Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params *float32) {
		defer glc.trace("Materialfv")
		C.gl12Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Materialiv = func(face, pname uint32, params *int32) {
		defer glc.trace("Materialiv")
		C.gl12Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.MatrixMode = func(mode uint32) {
		defer glc.trace("MatrixMode")
		C.gl12MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m *float64) {
		defer glc.trace("MultMatrixd")
		C.gl12MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultMatrixf = func(m *float32) {
		defer glc.trace("MultMatrixf")
		C.gl12MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.NewList = func(list uint32, mode uint32) {
		defer glc.trace("NewList")
		C.gl12NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		defer glc.trace("EndList")
		C.gl12EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		defer glc.trace("Normal3b")
		C.gl12Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		defer glc.trace("Normal3d")
		C.gl12Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		defer glc.trace("Normal3f")
		C.gl12Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		defer glc.trace("Normal3i")
		C.gl12Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		defer glc.trace("Normal3s")
		C.gl12Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v *int8) {
		defer glc.trace("Normal3bv")
		C.gl12Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Normal3dv = func(v *float64) {
		defer glc.trace("Normal3dv")
		C.gl12Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Normal3fv = func(v *float32) {
		defer glc.trace("Normal3fv")
		C.gl12Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Normal3iv = func(v *int32) {
		defer glc.trace("Normal3iv")
		C.gl12Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Normal3sv = func(v *int16) {
		defer glc.trace("Normal3sv")
		C.gl12Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		defer glc.trace("Ortho")
		C.gl12Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		defer glc.trace("PassThrough")
		C.gl12PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values *float32) {
		defer glc.trace("PixelMapfv")
		C.gl12PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values *uint32) {
		defer glc.trace("PixelMapuiv")
		C.gl12PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values *uint16) {
		defer glc.trace("PixelMapusv")
		C.gl12PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		defer glc.trace("PixelStoref")
		C.gl12PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		defer glc.trace("PixelStorei")
		C.gl12PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		defer glc.trace("PixelTransferf")
		C.gl12PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		defer glc.trace("PixelTransferi")
		C.gl12PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		defer glc.trace("PixelZoom")
		C.gl12PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		defer glc.trace("PointSize")
		C.gl12PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		defer glc.trace("PolygonMode")
		C.gl12PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask *uint8) {
		defer glc.trace("PolygonStipple")
		C.gl12PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(mask)))
	}

	glc.PushAttrib = func(mask uint32) {
		defer glc.trace("PushAttrib")
		C.gl12PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		defer glc.trace("PopAttrib")
		C.gl12PopAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		defer glc.trace("PushMatrix")
		C.gl12PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		defer glc.trace("PopMatrix")
		C.gl12PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		defer glc.trace("PushName")
		C.gl12PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		defer glc.trace("PopName")
		C.gl12PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		defer glc.trace("RasterPos2d")
		C.gl12RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		defer glc.trace("RasterPos2f")
		C.gl12RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		defer glc.trace("RasterPos2i")
		C.gl12RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		defer glc.trace("RasterPos2s")
		C.gl12RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		defer glc.trace("RasterPos3d")
		C.gl12RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		defer glc.trace("RasterPos3f")
		C.gl12RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		defer glc.trace("RasterPos3i")
		C.gl12RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		defer glc.trace("RasterPos3s")
		C.gl12RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		defer glc.trace("RasterPos4d")
		C.gl12RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		defer glc.trace("RasterPos4f")
		C.gl12RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		defer glc.trace("RasterPos4i")
		C.gl12RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		defer glc.trace("RasterPos4s")
		C.gl12RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v *float64) {
		defer glc.trace("RasterPos2dv")
		C.gl12RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos2fv = func(v *float32) {
		defer glc.trace("RasterPos2fv")
		C.gl12RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos2iv = func(v *int32) {
		defer glc.trace("RasterPos2iv")
		C.gl12RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos2sv = func(v *int16) {
		defer glc.trace("RasterPos2sv")
		C.gl12RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos3dv = func(v *float64) {
		defer glc.trace("RasterPos3dv")
		C.gl12RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos3fv = func(v *float32) {
		defer glc.trace("RasterPos3fv")
		C.gl12RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos3iv = func(v *int32) {
		defer glc.trace("RasterPos3iv")
		C.gl12RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos3sv = func(v *int16) {
		defer glc.trace("RasterPos3sv")
		C.gl12RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos4dv = func(v *float64) {
		defer glc.trace("RasterPos4dv")
		C.gl12RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos4fv = func(v *float32) {
		defer glc.trace("RasterPos4fv")
		C.gl12RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos4iv = func(v *int32) {
		defer glc.trace("RasterPos4iv")
		C.gl12RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos4sv = func(v *int16) {
		defer glc.trace("RasterPos4sv")
		C.gl12RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.ReadBuffer = func(mode uint32) {
		defer glc.trace("ReadBuffer")
		C.gl12ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("ReadPixels")
		C.gl12ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		defer glc.trace("Rectd")
		C.gl12Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		defer glc.trace("Rectf")
		C.gl12Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		defer glc.trace("Recti")
		C.gl12Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		defer glc.trace("Rects")
		C.gl12Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 *float64) {
		defer glc.trace("Rectdv")
		C.gl12Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
	}

	glc.Rectfv = func(v1, v2 *float32) {
		defer glc.trace("Rectfv")
		C.gl12Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
	}

	glc.Rectiv = func(v1, v2 *int32) {
		defer glc.trace("Rectiv")
		C.gl12Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
	}

	glc.Rectsv = func(v1, v2 *int16) {
		defer glc.trace("Rectsv")
		C.gl12Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
	}

	glc.RenderMode = func(mode uint32) int32 {
		defer glc.trace("RenderMode")
		return int32(C.gl12RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		defer glc.trace("Rotated")
		C.gl12Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		defer glc.trace("Rotatef")
		C.gl12Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		defer glc.trace("Scaled")
		C.gl12Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		defer glc.trace("Scalef")
		C.gl12Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		defer glc.trace("Scissor")
		C.gl12Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer *uint32) {
		defer glc.trace("SelectBuffer")
		C.gl12SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
	}

	glc.ShadeModel = func(mode uint32) {
		defer glc.trace("ShadeModel")
		C.gl12ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		defer glc.trace("StencilFunc")
		C.gl12StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		defer glc.trace("StencilMask")
		C.gl12StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		defer glc.trace("StencilOp")
		C.gl12StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		defer glc.trace("TexCoord1d")
		C.gl12TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		defer glc.trace("TexCoord1f")
		C.gl12TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		defer glc.trace("TexCoord1i")
		C.gl12TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		defer glc.trace("TexCoord1s")
		C.gl12TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		defer glc.trace("TexCoord2d")
		C.gl12TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		defer glc.trace("TexCoord2f")
		C.gl12TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		defer glc.trace("TexCoord2i")
		C.gl12TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		defer glc.trace("TexCoord2s")
		C.gl12TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		defer glc.trace("TexCoord3d")
		C.gl12TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		defer glc.trace("TexCoord3f")
		C.gl12TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		defer glc.trace("TexCoord3i")
		C.gl12TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		defer glc.trace("TexCoord3s")
		C.gl12TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		defer glc.trace("TexCoord4d")
		C.gl12TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		defer glc.trace("TexCoord4f")
		C.gl12TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		defer glc.trace("TexCoord4i")
		C.gl12TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		defer glc.trace("TexCoord4s")
		C.gl12TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v *float64) {
		defer glc.trace("TexCoord1dv")
		C.gl12TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord1fv = func(v *float32) {
		defer glc.trace("TexCoord1fv")
		C.gl12TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord1iv = func(v *int32) {
		defer glc.trace("TexCoord1iv")
		C.gl12TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord1sv = func(v *int16) {
		defer glc.trace("TexCoord1sv")
		C.gl12TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord2dv = func(v *float64) {
		defer glc.trace("TexCoord2dv")
		C.gl12TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord2fv = func(v *float32) {
		defer glc.trace("TexCoord2fv")
		C.gl12TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord2iv = func(v *int32) {
		defer glc.trace("TexCoord2iv")
		C.gl12TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord2sv = func(v *int16) {
		defer glc.trace("TexCoord2sv")
		C.gl12TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord3dv = func(v *float64) {
		defer glc.trace("TexCoord3dv")
		C.gl12TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord3fv = func(v *float32) {
		defer glc.trace("TexCoord3fv")
		C.gl12TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord3iv = func(v *int32) {
		defer glc.trace("TexCoord3iv")
		C.gl12TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord3sv = func(v *int16) {
		defer glc.trace("TexCoord3sv")
		C.gl12TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord4dv = func(v *float64) {
		defer glc.trace("TexCoord4dv")
		C.gl12TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord4fv = func(v *float32) {
		defer glc.trace("TexCoord4fv")
		C.gl12TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord4iv = func(v *int32) {
		defer glc.trace("TexCoord4iv")
		C.gl12TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord4sv = func(v *int16) {
		defer glc.trace("TexCoord4sv")
		C.gl12TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		defer glc.trace("TexEnvf")
		C.gl12TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		defer glc.trace("TexEnvi")
		C.gl12TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params *float32) {
		defer glc.trace("TexEnvfv")
		C.gl12TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexEnviv = func(target, pname uint32, params *int32) {
		defer glc.trace("TexEnviv")
		C.gl12TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		defer glc.trace("TexGend")
		C.gl12TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		defer glc.trace("TexGenf")
		C.gl12TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		defer glc.trace("TexGeni")
		C.gl12TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params *float64) {
		defer glc.trace("TexGendv")
		C.gl12TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.TexGenfv = func(coord, pname uint32, params *float32) {
		defer glc.trace("TexGenfv")
		C.gl12TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexGeniv = func(coord, pname uint32, params *int32) {
		defer glc.trace("TexGeniv")
		C.gl12TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexImage1D")
		C.gl12TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexImage2D")
		C.gl12TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		defer glc.trace("TexParameterf")
		C.gl12TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		defer glc.trace("TexParameteri")
		C.gl12TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("TexParameterfv")
		C.gl12TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("TexParameteriv")
		C.gl12TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Translated = func(x, y, z float64) {
		defer glc.trace("Translated")
		C.gl12Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		defer glc.trace("Translatef")
		C.gl12Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		defer glc.trace("Vertex2s")
		C.gl12Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		defer glc.trace("Vertex2i")
		C.gl12Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		defer glc.trace("Vertex2f")
		C.gl12Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		defer glc.trace("Vertex2d")
		C.gl12Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		defer glc.trace("Vertex3s")
		C.gl12Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		defer glc.trace("Vertex3i")
		C.gl12Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		defer glc.trace("Vertex3f")
		C.gl12Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		defer glc.trace("Vertex3d")
		C.gl12Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		defer glc.trace("Vertex4s")
		C.gl12Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		defer glc.trace("Vertex4i")
		C.gl12Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		defer glc.trace("Vertex4f")
		C.gl12Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		defer glc.trace("Vertex4d")
		C.gl12Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		defer glc.trace("Viewport")
		C.gl12Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetConvolutionParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetConvolutionParameterfv")
		C.gl12GetConvolutionParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetConvolutionParameteriv")
		C.gl12GetConvolutionParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		defer glc.trace("AreTexturesResident")
		var cRes *C.GLboolean
		status = C.gl12AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		defer glc.trace("ArrayElement")
		C.gl12ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		defer glc.trace("DrawArrays")
		C.gl12DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		defer glc.trace("DrawElements")
		C.gl12DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		defer glc.trace("GetPointerv")
		C.gl12GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		defer glc.trace("PolygonOffset")
		C.gl12PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		defer glc.trace("CopyTexImage1D")
		C.gl12CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		defer glc.trace("CopyTexImage2D")
		C.gl12CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		defer glc.trace("CopyTexSubImage1D")
		C.gl12CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		defer glc.trace("CopyTexSubImage2D")
		C.gl12CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		defer glc.trace("BindTexture")
		C.gl12BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures *uint32) {
		defer glc.trace("DeleteTextures")
		C.gl12DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.GenTextures = func(n int32, textures *uint32) {
		defer glc.trace("GenTextures")
		C.gl12GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.IsTexture = func(texture uint32) bool {
		defer glc.trace("IsTexture")
		return C.gl12IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("ColorPointer")
		C.gl12ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		defer glc.trace("EnableClientState")
		C.gl12EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		defer glc.trace("DisableClientState")
		C.gl12DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.Indexub = func(c uint8) {
		defer glc.trace("Indexub")
		C.gl12Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexubv = func(c *uint8) {
		defer glc.trace("Indexubv")
		C.gl12Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(c)))
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("InterleavedArrays")
		C.gl12InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("NormalPointer")
		C.gl12NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.PushClientAttrib = func(mask uint32) {
		defer glc.trace("PushClientAttrib")
		C.gl12PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PrioritizeTextures = func(n int32, textures *uint32, priorities *float32) {
		defer glc.trace("PrioritizeTextures")
		C.gl12PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLclampf)(unsafe.Pointer(priorities)))
	}

	glc.PopClientAttrib = func() {
		defer glc.trace("PopClientAttrib")
		C.gl12PopClientAttrib(glc.context)
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("TexCoordPointer")
		C.gl12TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexSubImage1D")
		C.gl12TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexSubImage2D")
		C.gl12TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("VertexPointer")
		C.gl12VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ColorTable")
		C.gl12ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("ColorTableParameterfv")
		C.gl12ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("ColorTableParameteriv")
		C.gl12ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ColorSubTable")
		C.gl12ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ConvolutionFilter1D")
		C.gl12ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ConvolutionFilter2D")
		C.gl12ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		defer glc.trace("ConvolutionParameterf")
		C.gl12ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		defer glc.trace("ConvolutionParameteri")
		C.gl12ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		defer glc.trace("CopyColorTable")
		C.gl12CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		defer glc.trace("CopyColorSubTable")
		C.gl12CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		defer glc.trace("CopyConvolutionFilter1D")
		C.gl12CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		defer glc.trace("CopyConvolutionFilter2D")
		C.gl12CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetColorTable = func(target, format, Type uint32, table unsafe.Pointer) {
		defer glc.trace("GetColorTable")
		C.gl12GetColorTable(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), table)
	}

	glc.GetColorTableParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetColorTableParameterfv")
		C.gl12GetColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetColorTableParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetColorTableParameteriv")
		C.gl12GetColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionFilter = func(target, format, Type uint32, image unsafe.Pointer) {
		defer glc.trace("GetConvolutionFilter")
		C.gl12GetConvolutionFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), image)
	}

	glc.GetHistogram = func(target uint32, reset bool, format, Type uint32, values unsafe.Pointer) {
		defer glc.trace("GetHistogram")
		C.gl12GetHistogram(glc.context, C.GLenum(target), boolToGL(reset), C.GLenum(format), C.GLenum(Type), values)
	}

	glc.GetHistogramParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetHistogramParameterfv")
		C.gl12GetHistogramParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetHistogramParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetHistogramParameteriv")
		C.gl12GetHistogramParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetSeparableFilter = func(target, format, Type uint32, row, column, span unsafe.Pointer) {
		defer glc.trace("GetSeparableFilter")
		C.gl12GetSeparableFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), row, column, span)
	}

	glc.Histogram = func(target uint32, width int32, internalformat uint32, sink bool) {
		defer glc.trace("Histogram")
		C.gl12Histogram(glc.context, C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.Minmax = func(target, internalformat uint32, sink bool) {
		defer glc.trace("Minmax")
		C.gl12Minmax(glc.context, C.GLenum(target), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.MultiTexCoord1s = func(target uint32, s int16) {
		defer glc.trace("MultiTexCoord1s")
		C.gl12MultiTexCoord1s(glc.context, C.GLenum(target), C.GLshort(s))
	}

	glc.MultiTexCoord1i = func(target uint32, s int32) {
		defer glc.trace("MultiTexCoord1i")
		C.gl12MultiTexCoord1i(glc.context, C.GLenum(target), C.GLint(s))
	}

	glc.MultiTexCoord1f = func(target uint32, s float32) {
		defer glc.trace("MultiTexCoord1f")
		C.gl12MultiTexCoord1f(glc.context, C.GLenum(target), C.GLfloat(s))
	}

	glc.MultiTexCoord1d = func(target uint32, s float64) {
		defer glc.trace("MultiTexCoord1d")
		C.gl12MultiTexCoord1d(glc.context, C.GLenum(target), C.GLdouble(s))
	}

	glc.MultiTexCoord2s = func(target uint32, s, t int16) {
		defer glc.trace("MultiTexCoord2s")
		C.gl12MultiTexCoord2s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t))
	}

	glc.MultiTexCoord2i = func(target uint32, s, t int32) {
		defer glc.trace("MultiTexCoord2i")
		C.gl12MultiTexCoord2i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t))
	}

	glc.MultiTexCoord2f = func(target uint32, s, t float32) {
		defer glc.trace("MultiTexCoord2f")
		C.gl12MultiTexCoord2f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
	}

	glc.MultiTexCoord2d = func(target uint32, s, t float64) {
		defer glc.trace("MultiTexCoord2d")
		C.gl12MultiTexCoord2d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
	}

	glc.MultiTexCoord3s = func(target uint32, s, t, r int16) {
		defer glc.trace("MultiTexCoord3s")
		C.gl12MultiTexCoord3s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.MultiTexCoord3i = func(target uint32, s, t, r int32) {
		defer glc.trace("MultiTexCoord3i")
		C.gl12MultiTexCoord3i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.MultiTexCoord3f = func(target uint32, s, t, r float32) {
		defer glc.trace("MultiTexCoord3f")
		C.gl12MultiTexCoord3f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.MultiTexCoord3d = func(target uint32, s, t, r float64) {
		defer glc.trace("MultiTexCoord3d")
		C.gl12MultiTexCoord3d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.MultiTexCoord4s = func(target uint32, s, t, r, q int16) {
		defer glc.trace("MultiTexCoord4s")
		C.gl12MultiTexCoord4s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.MultiTexCoord4i = func(target uint32, s, t, r, q int32) {
		defer glc.trace("MultiTexCoord4i")
		C.gl12MultiTexCoord4i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.MultiTexCoord4f = func(target uint32, s, t, r, q float32) {
		defer glc.trace("MultiTexCoord4f")
		C.gl12MultiTexCoord4f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.MultiTexCoord4d = func(target uint32, s, t, r, q float64) {
		defer glc.trace("MultiTexCoord4d")
		C.gl12MultiTexCoord4d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.MultiTexCoord1sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord1sv")
		C.gl12MultiTexCoord1sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord1iv")
		C.gl12MultiTexCoord1iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord1fv")
		C.gl12MultiTexCoord1fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord1dv")
		C.gl12MultiTexCoord1dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord2sv")
		C.gl12MultiTexCoord2sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord2iv")
		C.gl12MultiTexCoord2iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord2fv")
		C.gl12MultiTexCoord2fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord2dv")
		C.gl12MultiTexCoord2dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord3sv")
		C.gl12MultiTexCoord3sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord3iv")
		C.gl12MultiTexCoord3iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord3fv")
		C.gl12MultiTexCoord3fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord3dv")
		C.gl12MultiTexCoord3dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord4sv")
		C.gl12MultiTexCoord4sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord4iv")
		C.gl12MultiTexCoord4iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord4fv")
		C.gl12MultiTexCoord4fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord4dv")
		C.gl12MultiTexCoord4dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.ResetHistogram = func(target uint32) {
		defer glc.trace("ResetHistogram")
		C.gl12ResetHistogram(glc.context, C.GLenum(target))
	}

	glc.ResetMinmax = func(target uint32) {
		defer glc.trace("ResetMinmax")
		C.gl12ResetMinmax(glc.context, C.GLenum(target))
	}

	glc.SeparableFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, row, column unsafe.Pointer) {
		defer glc.trace("SeparableFilter2D")
		C.gl12SeparableFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), row, column)
	}

	glc.BlendColor = func(red, green, blue, alpha float32) {
		defer glc.trace("BlendColor")
		C.gl12BlendColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		defer glc.trace("BlendEquation")
		C.gl12BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		defer glc.trace("CopyTexSubImage3D")
		C.gl12CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DrawRangeElements = func(mode uint32, start, end uint32, count int32, Type uint32, indices unsafe.Pointer) {
		defer glc.trace("DrawRangeElements")
		C.gl12DrawRangeElements(glc.context, C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.TexImage3D = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexImage3D")
		C.gl12TexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexSubImage3D")
		C.gl12TexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.ActiveTexture = func(texture uint32) {
		defer glc.trace("ActiveTexture")
		C.gl12ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		defer glc.trace("ClientActiveTexture")
		C.gl12ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexImage1D")
		C.gl12CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexImage2D")
		C.gl12CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexImage3D")
		C.gl12CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexSubImage1D")
		C.gl12CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexSubImage2D")
		C.gl12CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexSubImage3D")
		C.gl12CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.GetCompressedTexImage = func(target uint32, lod int32, img unsafe.Pointer) {
		defer glc.trace("GetCompressedTexImage")
		C.gl12GetCompressedTexImage(glc.context, C.GLenum(target), C.GLint(lod), img)
	}

	glc.LoadTransposeMatrixd = func(m *float64) {
		defer glc.trace("LoadTransposeMatrixd")
		C.gl12LoadTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadTransposeMatrixf = func(m *float64) {
		defer glc.trace("LoadTransposeMatrixf")
		C.gl12LoadTransposeMatrixf(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixd = func(m *float64) {
		defer glc.trace("MultTransposeMatrixd")
		C.gl12MultTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixf = func(m *float32) {
		defer glc.trace("MultTransposeMatrixf")
		C.gl12MultTransposeMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.SampleCoverage = func(value float32, invert bool) {
		defer glc.trace("SampleCoverage")
		C.gl12SampleCoverage(glc.context, C.GLclampf(value), boolToGL(invert))
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		defer glc.trace("BlendFuncSeparate")
		C.gl12BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.FogCoordPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("FogCoordPointer")
		C.gl12FogCoordPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.FogCoordd = func(coord float64) {
		defer glc.trace("FogCoordd")
		C.gl12FogCoordd(glc.context, C.GLdouble(coord))
	}

	glc.FogCoordf = func(coord float32) {
		defer glc.trace("FogCoordf")
		C.gl12FogCoordf(glc.context, C.GLfloat(coord))
	}

	glc.FogCoorddv = func(coord *float64) {
		defer glc.trace("FogCoorddv")
		C.gl12FogCoorddv(glc.context, (*C.GLdouble)(unsafe.Pointer(coord)))
	}

	glc.FogCoordfv = func(coord *float32) {
		defer glc.trace("FogCoordfv")
		C.gl12FogCoordfv(glc.context, (*C.GLfloat)(unsafe.Pointer(coord)))
	}

	glc.MultiDrawArrays = func(mode uint32, first *int32, count *int32, primcount int32) {
		defer glc.trace("MultiDrawArrays")
		C.gl12MultiDrawArrays(glc.context, C.GLenum(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), C.GLsizei(primcount))
	}

	glc.MultiDrawElements = func(mode uint32, count *int32, Type uint32, indices unsafe.Pointer, primcount int32) {
		defer glc.trace("MultiDrawElements")
		C.gl12MultiDrawElements(glc.context, C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(count)), C.GLenum(Type), indices, C.GLsizei(primcount))
	}

	glc.PointParameterf = func(pname uint32, param float32) {
		defer glc.trace("PointParameterf")
		C.gl12PointParameterf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PointParameteri = func(pname uint32, param int32) {
		defer glc.trace("PointParameteri")
		C.gl12PointParameteri(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.SecondaryColor3b = func(red, green, blue int8) {
		defer glc.trace("SecondaryColor3b")
		C.gl12SecondaryColor3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.SecondaryColor3s = func(red, green, blue int16) {
		defer glc.trace("SecondaryColor3s")
		C.gl12SecondaryColor3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.SecondaryColor3i = func(red, green, blue int32) {
		defer glc.trace("SecondaryColor3i")
		C.gl12SecondaryColor3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.SecondaryColor3f = func(red, green, blue float32) {
		defer glc.trace("SecondaryColor3f")
		C.gl12SecondaryColor3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.SecondaryColor3d = func(red, green, blue float64) {
		defer glc.trace("SecondaryColor3d")
		C.gl12SecondaryColor3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.SecondaryColor3ub = func(red, green, blue uint8) {
		defer glc.trace("SecondaryColor3ub")
		C.gl12SecondaryColor3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.SecondaryColor3us = func(red, green, blue uint16) {
		defer glc.trace("SecondaryColor3us")
		C.gl12SecondaryColor3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.SecondaryColor3ui = func(red, green, blue uint32) {
		defer glc.trace("SecondaryColor3ui")
		C.gl12SecondaryColor3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.SecondaryColor3bv = func(v *int8) {
		defer glc.trace("SecondaryColor3bv")
		C.gl12SecondaryColor3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3sv = func(v *int16) {
		defer glc.trace("SecondaryColor3sv")
		C.gl12SecondaryColor3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3iv = func(v *int32) {
		defer glc.trace("SecondaryColor3iv")
		C.gl12SecondaryColor3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3fv = func(v *float32) {
		defer glc.trace("SecondaryColor3fv")
		C.gl12SecondaryColor3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3dv = func(v *float64) {
		defer glc.trace("SecondaryColor3dv")
		C.gl12SecondaryColor3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3ubv = func(v *uint8) {
		defer glc.trace("SecondaryColor3ubv")
		C.gl12SecondaryColor3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3usv = func(v *uint16) {
		defer glc.trace("SecondaryColor3usv")
		C.gl12SecondaryColor3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3uiv = func(v *uint32) {
		defer glc.trace("SecondaryColor3uiv")
		C.gl12SecondaryColor3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("SecondaryColorPointer")
		C.gl12SecondaryColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.WindowPos2s = func(x, y int16) {
		defer glc.trace("WindowPos2s")
		C.gl12WindowPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.WindowPos2i = func(x, y int32) {
		defer glc.trace("WindowPos2i")
		C.gl12WindowPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.WindowPos2f = func(x, y float32) {
		defer glc.trace("WindowPos2f")
		C.gl12WindowPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.WindowPos2d = func(x, y float64) {
		defer glc.trace("WindowPos2d")
		C.gl12WindowPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.WindowPos3s = func(x, y, z int16) {
		defer glc.trace("WindowPos3s")
		C.gl12WindowPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.WindowPos3i = func(x, y, z int32) {
		defer glc.trace("WindowPos3i")
		C.gl12WindowPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.WindowPos3f = func(x, y, z float32) {
		defer glc.trace("WindowPos3f")
		C.gl12WindowPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.WindowPos3d = func(x, y, z float64) {
		defer glc.trace("WindowPos3d")
		C.gl12WindowPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.WindowPos2sv = func(v *int16) {
		defer glc.trace("WindowPos2sv")
		C.gl12WindowPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos2iv = func(v *int32) {
		defer glc.trace("WindowPos2iv")
		C.gl12WindowPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos2fv = func(v *float32) {
		defer glc.trace("WindowPos2fv")
		C.gl12WindowPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos2dv = func(v *float64) {
		defer glc.trace("WindowPos2dv")
		C.gl12WindowPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.WindowPos3sv = func(v *int16) {
		defer glc.trace("WindowPos3sv")
		C.gl12WindowPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos3iv = func(v *int32) {
		defer glc.trace("WindowPos3iv")
		C.gl12WindowPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos3fv = func(v *float32) {
		defer glc.trace("WindowPos3fv")
		C.gl12WindowPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos3dv = func(v *float64) {
		defer glc.trace("WindowPos3dv")
		C.gl12WindowPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		defer glc.trace("BeginQuery")
		C.gl12BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		defer glc.trace("BindBuffer")
		C.gl12BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		defer glc.trace("BufferData")
		C.gl12BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		defer glc.trace("BufferSubData")
		C.gl12BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.DeleteBuffers = func(n int32, buffers *uint32) {
		defer glc.trace("DeleteBuffers")
		C.gl12DeleteBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.DeleteQueries = func(n int32, ids *uint32) {
		defer glc.trace("DeleteQueries")
		C.gl12DeleteQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GenBuffers = func(n int32, buffers *uint32) {
		defer glc.trace("GenBuffers")
		C.gl12GenBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.GenQueries = func(n int32, ids *uint32) {
		defer glc.trace("GenQueries")
		C.gl12GenQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GetBufferParameteriv = func(target, value uint32, data *int32) {
		defer glc.trace("GetBufferParameteriv")
		C.gl12GetBufferParameteriv(glc.context, C.GLenum(target), C.GLenum(value), (*C.GLint)(unsafe.Pointer(data)))
	}

	glc.GetBufferPointerv = func(target, pname uint32, params unsafe.Pointer) {
		defer glc.trace("GetBufferPointerv")
		C.gl12GetBufferPointerv(glc.context, C.GLenum(target), C.GLenum(pname), params)
	}

	glc.GetBufferSubData = func(target uint32, offset int32, size int32, data unsafe.Pointer) {
		defer glc.trace("GetBufferSubData")
		C.gl12GetBufferSubData(glc.context, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}

	glc.GetQueryObjectiv = func(id uint32, pname uint32, params *int32) {
		defer glc.trace("GetQueryObjectiv")
		C.gl12GetQueryObjectiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetQueryObjectuiv = func(id uint32, pname uint32, params *uint32) {
		defer glc.trace("GetQueryObjectuiv")
		C.gl12GetQueryObjectuiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(params)))
	}

	glc.GetQueryiv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetQueryiv")
		C.gl12GetQueryiv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.IsBuffer = func(buffer uint32) bool {
		defer glc.trace("IsBuffer")
		return C.gl12IsBuffer(glc.context, C.GLuint(buffer)) != 0
	}

	glc.IsQuery = func(id uint32) bool {
		defer glc.trace("IsQuery")
		return C.gl12IsQuery(glc.context, C.GLuint(id)) != 0
	}

	glc.MapBuffer = func(target, access uint32) unsafe.Pointer {
		defer glc.trace("MapBuffer")
		return unsafe.Pointer(C.gl12MapBuffer(glc.context, C.GLenum(target), C.GLenum(access)))
	}

	glc.UnmapBuffer = func(target uint32) bool {
		defer glc.trace("UnmapBuffer")
		return C.gl12UnmapBuffer(glc.context, C.GLenum(target)) != 0
	}

	glc.AttachShader = func(program, shader uint32) {
		defer glc.trace("AttachShader")
		C.gl12AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		defer glc.trace("BindAttribLocation")
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl12BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha uint32) {
		defer glc.trace("BlendEquationSeperate")
		C.gl12BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		defer glc.trace("CompileShader")
		C.gl12CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		defer glc.trace("CreateProgram")
		return uint32(C.gl12CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType uint32) uint32 {
		defer glc.trace("CreateShader")
		return uint32(C.gl12CreateShader(glc.context, C.GLenum(shaderType)))
	}

	glc.DeleteProgram = func(program uint32) {
		defer glc.trace("DeleteProgram")
		C.gl12DeleteProgram(glc.context, C.GLuint(program))
	}

	glc.DeleteShader = func(shader uint32) {
		defer glc.trace("DeleteShader")
		C.gl12DeleteShader(glc.context, C.GLuint(shader))
	}

	glc.DetachShader = func(program, shader uint32) {
		defer glc.trace("DetachShader")
		C.gl12DetachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.EnableVertexAttribArray = func(index uint32) {
		defer glc.trace("EnableVertexAttribArray")
		C.gl12EnableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DisableVertexAttribArray = func(index uint32) {
		defer glc.trace("DisableVertexAttribArray")
		C.gl12DisableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DrawBuffers = func(n int32, bufs *uint32) {
		defer glc.trace("DrawBuffers")
		C.gl12DrawBuffers(glc.context, C.GLsizei(n), (*C.GLenum)(unsafe.Pointer(bufs)))
	}

	glc.GetActiveAttrib = func(program, index uint32, bufSize int32) (length int32, size int32, Type uint32, name string) {
		defer glc.trace("GetActiveAttrib")
		var (
			cname C.GLchar
		)
		C.gl12GetActiveAttrib(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
		name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
		return
	}

	glc.GetActiveUniform = func(program, index uint32, bufSize int32, length *int32, size *int32, Type *uint32, name *byte) {
		defer glc.trace("GetActiveUniform")
		C.gl12GetActiveUniform(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(Type)), (*C.GLchar)(unsafe.Pointer(name)))
	}

	glc.GetAttachedShaders = func(program uint32, maxCount int32, count *int32, shaders *uint32) {
		defer glc.trace("GetAttachedShaders")
		C.gl12GetAttachedShaders(glc.context, C.GLuint(program), C.GLsizei(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
	}

	glc.GetAttribLocation = func(program uint32, name *byte) int32 {
		defer glc.trace("GetAttribLocation")
		return int32(C.gl12GetAttribLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetProgramiv = func(program uint32, pname uint32, params *int32) {
		defer glc.trace("GetProgramiv")
		C.gl12GetProgramiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetProgramInfoLog = func(program uint32, maxLength int32, length *int32, infoLog *byte) {
		defer glc.trace("GetProgramInfoLog")
		C.gl12GetProgramInfoLog(glc.context, C.GLuint(program), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderiv = func(program uint32, pname uint32, params *int32) {
		defer glc.trace("GetShaderiv")
		C.gl12GetShaderiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetShaderInfoLog = func(shader uint32, maxLength int32, length *int32, infoLog *byte) {
		defer glc.trace("GetShaderInfoLog")
		C.gl12GetShaderInfoLog(glc.context, C.GLuint(shader), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderSource = func(shader uint32, bufSize int32, length *int32, source *byte) {
		defer glc.trace("GetShaderSource")
		C.gl12GetShaderSource(glc.context, C.GLuint(shader), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
	}

	glc.GetUniformfv = func(program uint32, location int32, params *float32) {
		defer glc.trace("GetUniformfv")
		C.gl12GetUniformfv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetUniformiv = func(program uint32, location int32, params *int32) {
		defer glc.trace("GetUniformiv")
		C.gl12GetUniformiv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetUniformLocation = func(program uint32, name *byte) int32 {
		defer glc.trace("GetUniformLocation")
		return int32(C.gl12GetUniformLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetVertexAttribdv = func(index uint32, pname uint32, params *float64) {
		defer glc.trace("GetVertexAttribdv")
		C.gl12GetVertexAttribdv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribfv = func(index uint32, pname uint32, params *float32) {
		defer glc.trace("GetVertexAttribfv")
		C.gl12GetVertexAttribfv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribiv = func(index uint32, pname uint32, params *int32) {
		defer glc.trace("GetVertexAttribiv")
		C.gl12GetVertexAttribiv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribPointerv = func(index uint32, pname uint32, pointer unsafe.Pointer) {
		defer glc.trace("GetVertexAttribPointerv")
		C.gl12GetVertexAttribPointerv(glc.context, C.GLuint(index), C.GLenum(pname), pointer)
	}

	glc.IsProgram = func(program uint32) bool {
		defer glc.trace("IsProgram")
		return C.gl12IsProgram(glc.context, C.GLuint(program)) != 0
	}

	glc.IsShader = func(shader uint32) bool {
		defer glc.trace("IsShader")
		return C.gl12IsShader(glc.context, C.GLuint(shader)) != 0
	}

	glc.LinkProgram = func(program uint32) {
		defer glc.trace("LinkProgram")
		C.gl12LinkProgram(glc.context, C.GLuint(program))
	}

	glc.ShaderSource = func(shader uint32, count int32, string **byte, length *int32) {
		defer glc.trace("ShaderSource")
		C.gl12ShaderSource(glc.context, C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(string)), (*C.GLint)(unsafe.Pointer(length)))
	}

	glc.StencilFuncSeparate = func(face, Func uint32, ref int32, mask uint32) {
		defer glc.trace("StencilFuncSeparate")
		C.gl12StencilFuncSeparate(glc.context, C.GLenum(face), C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMaskSeparate = func(face uint32, mask uint32) {
		defer glc.trace("StencilMaskSeparate")
		C.gl12StencilMaskSeparate(glc.context, C.GLenum(face), C.GLuint(mask))
	}

	glc.StencilOpSeparate = func(face, sfail, dpfail, dppass uint32) {
		defer glc.trace("StencilOpSeparate")
		C.gl12StencilOpSeparate(glc.context, C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
	}

	glc.Uniform1f = func(location int32, v0 float32) {
		defer glc.trace("Uniform1f")
		C.gl12Uniform1f(glc.context, C.GLint(location), C.GLfloat(v0))
	}

	glc.Uniform2f = func(location int32, v0, v1 float32) {
		defer glc.trace("Uniform2f")
		C.gl12Uniform2f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.Uniform3f = func(location int32, v0, v1, v2 float32) {
		defer glc.trace("Uniform3f")
		C.gl12Uniform3f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Uniform4f = func(location int32, v0, v1, v2, v3 float32) {
		defer glc.trace("Uniform4f")
		C.gl12Uniform4f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.Uniform1i = func(location, v0 int32) {
		defer glc.trace("Uniform1i")
		C.gl12Uniform1i(glc.context, C.GLint(location), C.GLint(v0))
	}

	glc.Uniform2i = func(location, v0, v1 int32) {
		defer glc.trace("Uniform2i")
		C.gl12Uniform2i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1))
	}

	glc.Uniform3i = func(location, v0, v1, v2 int32) {
		defer glc.trace("Uniform3i")
		C.gl12Uniform3i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
	}

	glc.Uniform4i = func(location, v0, v1, v2, v3 int32) {
		defer glc.trace("Uniform4i")
		C.gl12Uniform4i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
	}

	glc.Uniform1ui = func(location int32, v0 uint32) {
		defer glc.trace("Uniform1ui")
		C.gl12Uniform1ui(glc.context, C.GLint(location), C.GLuint(v0))
	}

	glc.Uniform2ui = func(location int32, v0, v1 uint32) {
		defer glc.trace("Uniform2ui")
		C.gl12Uniform2ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1))
	}

	glc.Uniform3ui = func(location int32, v0, v1, v2 uint32) {
		defer glc.trace("Uniform3ui")
		C.gl12Uniform3ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
	}

	glc.Uniform4ui = func(location int32, v0, v1, v2, v3 uint32) {
		defer glc.trace("Uniform4ui")
		C.gl12Uniform4ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
	}

	glc.Uniform1fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform1fv")
		C.gl12Uniform1fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform2fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform2fv")
		C.gl12Uniform2fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform3fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform3fv")
		C.gl12Uniform3fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform4fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform4fv")
		C.gl12Uniform4fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform1iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform1iv")
		C.gl12Uniform1iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform2iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform2iv")
		C.gl12Uniform2iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform3iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform3iv")
		C.gl12Uniform3iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform4iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform4iv")
		C.gl12Uniform4iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform1uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform1uiv")
		C.gl12Uniform1uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform2uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform2uiv")
		C.gl12Uniform2uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform3uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform3uiv")
		C.gl12Uniform3uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform4uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform4uiv")
		C.gl12Uniform4uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.UseProgram = func(program uint32) {
		defer glc.trace("UseProgram")
		C.gl12UseProgram(glc.context, C.GLuint(program))
	}

	glc.ValidateProgram = func(program uint32) {
		defer glc.trace("ValidateProgram")
		C.gl12ValidateProgram(glc.context, C.GLuint(program))
	}

	glc.VertexAttribPointer = func(index uint32, size int32, Type uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("VertexAttribPointer")
		C.gl12VertexAttribPointer(glc.context, C.GLuint(index), C.GLint(size), C.GLenum(Type), boolToGL(normalized), C.GLsizei(stride), pointer)
	}

	glc.VertexAttrib1f = func(index uint32, v0 float32) {
		defer glc.trace("VertexAttrib1f")
		C.gl12VertexAttrib1f(glc.context, C.GLuint(index), C.GLfloat(v0))
	}

	glc.VertexAttrib1s = func(index uint32, v0 int16) {
		defer glc.trace("VertexAttrib1s")
		C.gl12VertexAttrib1s(glc.context, C.GLuint(index), C.GLshort(v0))
	}

	glc.VertexAttrib1d = func(index uint32, v0 float64) {
		defer glc.trace("VertexAttrib1d")
		C.gl12VertexAttrib1d(glc.context, C.GLuint(index), C.GLdouble(v0))
	}

	glc.VertexAttrib2f = func(index uint32, v0, v1 float32) {
		defer glc.trace("VertexAttrib2f")
		C.gl12VertexAttrib2f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.VertexAttrib2s = func(index uint32, v0, v1 int16) {
		defer glc.trace("VertexAttrib2s")
		C.gl12VertexAttrib2s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
	}

	glc.VertexAttrib2d = func(index uint32, v0, v1 float64) {
		defer glc.trace("VertexAttrib2d")
		C.gl12VertexAttrib2d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
	}

	glc.VertexAttrib3f = func(index uint32, v0, v1, v2 float32) {
		defer glc.trace("VertexAttrib3f")
		C.gl12VertexAttrib3f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.VertexAttrib3s = func(index uint32, v0, v1, v2 int16) {
		defer glc.trace("VertexAttrib3s")
		C.gl12VertexAttrib3s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
	}

	glc.VertexAttrib3d = func(index uint32, v0, v1, v2 float64) {
		defer glc.trace("VertexAttrib3d")
		C.gl12VertexAttrib3d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.VertexAttrib4f = func(index uint32, v0, v1, v2, v3 float32) {
		defer glc.trace("VertexAttrib4f")
		C.gl12VertexAttrib4f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.VertexAttrib4s = func(index uint32, v0, v1, v2, v3 int16) {
		defer glc.trace("VertexAttrib4s")
		C.gl12VertexAttrib4s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
	}

	glc.VertexAttrib4d = func(index uint32, v0, v1, v2, v3 float64) {
		defer glc.trace("VertexAttrib4d")
		C.gl12VertexAttrib4d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
	}

	glc.VertexAttrib4Nuv = func(index uint32, v0, v1, v2, v3 uint8) {
		defer glc.trace("VertexAttrib4Nuv")
		C.gl12VertexAttrib4Nuv(glc.context, C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
	}

	glc.VertexAttrib1fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib1fv")
		C.gl12VertexAttrib1fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib1sv")
		C.gl12VertexAttrib1sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib1dv")
		C.gl12VertexAttrib1dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib2fv")
		C.gl12VertexAttrib2fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib2sv")
		C.gl12VertexAttrib2sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib2dv")
		C.gl12VertexAttrib2dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib3fv")
		C.gl12VertexAttrib3fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib3sv")
		C.gl12VertexAttrib3sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib3dv")
		C.gl12VertexAttrib3dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib4fv")
		C.gl12VertexAttrib4fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib4sv")
		C.gl12VertexAttrib4sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib4dv")
		C.gl12VertexAttrib4dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4iv = func(index uint32, v *int32) {
		defer glc.trace("VertexAttrib4iv")
		C.gl12VertexAttrib4iv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4bv = func(index uint32, v *int8) {
		defer glc.trace("VertexAttrib4bv")
		C.gl12VertexAttrib4bv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4ubv = func(index uint32, v *uint8) {
		defer glc.trace("VertexAttrib4ubv")
		C.gl12VertexAttrib4ubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4usv = func(index uint32, v *uint16) {
		defer glc.trace("VertexAttrib4usv")
		C.gl12VertexAttrib4usv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4uiv = func(index uint32, v *uint32) {
		defer glc.trace("VertexAttrib4uiv")
		C.gl12VertexAttrib4uiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nbv = func(index uint32, v *int8) {
		defer glc.trace("VertexAttrib4Nbv")
		C.gl12VertexAttrib4Nbv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nsv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib4Nsv")
		C.gl12VertexAttrib4Nsv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Niv = func(index uint32, v *int32) {
		defer glc.trace("VertexAttrib4Niv")
		C.gl12VertexAttrib4Niv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nubv = func(index uint32, v *uint8) {
		defer glc.trace("VertexAttrib4Nubv")
		C.gl12VertexAttrib4Nubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nusv = func(index uint32, v *uint16) {
		defer glc.trace("VertexAttrib4Nusv")
		C.gl12VertexAttrib4Nusv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nuiv = func(index uint32, v *uint32) {
		defer glc.trace("VertexAttrib4Nuiv")
		C.gl12VertexAttrib4Nuiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.UniformMatrix2fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix2fv")
		C.gl12UniformMatrix2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix3fv")
		C.gl12UniformMatrix3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix4fv")
		C.gl12UniformMatrix4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x3fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix2x3fv")
		C.gl12UniformMatrix2x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x2fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix3x2fv")
		C.gl12UniformMatrix3x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x4fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix2x4fv")
		C.gl12UniformMatrix2x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x2fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix4x2fv")
		C.gl12UniformMatrix4x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x4fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix3x4fv")
		C.gl12UniformMatrix3x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x3fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix4x3fv")
		C.gl12UniformMatrix4x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	if !versionSupported(glc) {
		return nil
	}
	glc.queryExtensions()
	return glc
}