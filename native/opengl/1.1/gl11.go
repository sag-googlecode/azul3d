// Package 'opengl' implements OpenGL version 1.1
package opengl

// #cgo LDFLAGS: -lopengl32
// #include "gl11.h"
import "C"

import (
	"fmt"
	"strconv"
	"strings"
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
		n, wantedMajor, wantedMinor, wantedRev := parseVersions("1.1")
		if n < 2 {
			fmt.Printf("OpenGL: *** JSON version parsing failed for %q ***\n", "1.1")
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
			if major == wantedMajor && minor >= wantedMinor {
				return true
			}
		} else {
			if major == wantedMajor && minor == wantedMinor && rev >= wantedRev {
				return true
			}
		}
	}
	return false
}

const (
	MIRRORED_REPEAT_ARB                                        = 0x8370
	EMBOSS_LIGHT_NV                                            = 0x855D
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	DRAW_BUFFER3_NV                                            = 0x8828
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	LUMINANCE12                                                = 0x8041
	COLOR_INDEX4_EXT                                           = 0x80E4
	RGB_SCALE_EXT                                              = 0x8573
	MAP_TESSELLATION_NV                                        = 0x86C2
	PIXEL_PACK_BUFFER                                          = 0x88EB
	REG_4_ATI                                                  = 0x8925
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	RGB4                                                       = 0x804F
	MAX                                                        = 0x8008
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	COMPILE                                                    = 0x1300
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	GL_2PASS_0_EXT                                             = 0x80A2
	COLOR_INDEX16_EXT                                          = 0x80E7
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	DECR_WRAP_EXT                                              = 0x8508
	SAMPLER_1D_SHADOW                                          = 0x8B61
	EXT_histogram                                              = 1
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	T2F_IUI_V3F_EXT                                            = 0x81B2
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	MIPMAP                                                     = 0x8293
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	FRONT_FACE                                                 = 0x0B46
	TEXTURE_ENV_COLOR                                          = 0x2201
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	SRGB_EXT                                                   = 0x8C40
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	ADD_SIGNED_ARB                                             = 0x8574
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	RGB16F_ARB                                                 = 0x881B
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	BLEND                                                      = 0x0BE2
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	MAX_VARYING_VECTORS                                        = 0x8DFC
	SGIX_clipmap                                               = 1
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	BGR                                                        = 0x80E0
	RGBA_FLOAT16_APPLE                                         = 0x881A
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	STENCIL_ATTACHMENT                                         = 0x8D20
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	REPLACE                                                    = 0x1E01
	LUMINANCE4_EXT                                             = 0x803F
	COLOR_TABLE_FORMAT                                         = 0x80D8
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	MULT                                                       = 0x0103
	TEXTURE_ENV                                                = 0x2300
	DRAW_BUFFER4_ATI                                           = 0x8829
	CON_24_ATI                                                 = 0x8959
	PROGRAM_INPUT                                              = 0x92E3
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	ADD_SIGNED                                                 = 0x8574
	OUTPUT_COLOR0_EXT                                          = 0x879B
	BIAS_BIT_ATI                                               = 0x00000008
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	YCBYCR8_422_NV                                             = 0x9031
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	BUFFER_ACCESS                                              = 0x88BB
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	FACTOR_MIN_AMD                                             = 0x901C
	SGIX_texture_lod_bias                                      = 1
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	NORMAL_ARRAY                                               = 0x8075
	ALPHA16_EXT                                                = 0x803E
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	RGB16F                                                     = 0x881B
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	EXT_abgr                                                   = 1
	SAMPLES_SGIS                                               = 0x80A9
	DUAL_ALPHA16_SGIS                                          = 0x8113
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	VARIABLE_A_NV                                              = 0x8523
	DOUBLE_MAT3                                                = 0x8F47
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	SLIM12S_SGIX                                               = 0x831F
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	AFFINE_3D_NV                                               = 0x9094
	STORAGE_CACHED_APPLE                                       = 0x85BE
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	READ_FRAMEBUFFER                                           = 0x8CA8
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	FACTOR_MAX_AMD                                             = 0x901D
	COPY_PIXEL_TOKEN                                           = 0x0706
	LIGHT1                                                     = 0x4001
	TEXTURE16                                                  = 0x84D0
	HALF_FLOAT_NV                                              = 0x140B
	ALPHA12_EXT                                                = 0x803D
	READ_PIXELS_TYPE                                           = 0x828E
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	COMPRESSED_RG11_EAC                                        = 0x9272
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	RG16F                                                      = 0x822F
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	MATRIX_EXT                                                 = 0x87C0
	SGIS_texture4D                                             = 1
	INVALID_ENUM                                               = 0x0500
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	UNSIGNED_NORMALIZED                                        = 0x8C17
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	LINE_TO_NV                                                 = 0x04
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	FOG_INDEX                                                  = 0x0B61
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	UNIFORM_OFFSET                                             = 0x8A3B
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	FUNC_SUBTRACT_EXT                                          = 0x800A
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	PROGRAM_SEPARABLE                                          = 0x8258
	EXT_convolution                                            = 1
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	ACCUM                                                      = 0x0100
	PROGRAM_RESIDENT_NV                                        = 0x8647
	DU8DV8_ATI                                                 = 0x877A
	RED_BIT_ATI                                                = 0x00000001
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	OPERAND0_ALPHA_ARB                                         = 0x8598
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	SRC0_RGB                                                   = 0x8580
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	FENCE_APPLE                                                = 0x8A0B
	FLOAT_VEC3                                                 = 0x8B51
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	MINMAX_SINK_EXT                                            = 0x8030
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	RG                                                         = 0x8227
	MAX_VARYING_FLOATS                                         = 0x8B4B
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	V2F                                                        = 0x2A20
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	COORD_REPLACE_ARB                                          = 0x8862
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	TEXTURE_RED_TYPE                                           = 0x8C10
	PATH_DASH_OFFSET_NV                                        = 0x907E
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	SHADER_IMAGE_LOAD                                          = 0x82A4
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	ALPHA8_SNORM                                               = 0x9014
	GL_2D                                                      = 0x0600
	CONSTANT_BORDER                                            = 0x8151
	CAVEAT_SUPPORT                                             = 0x82B8
	SOURCE0_RGB_ARB                                            = 0x8580
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	LAYOUT_LINEAR_INTEL                                        = 1
	CONVOLUTION_2D                                             = 0x8011
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	GREEN_INTEGER                                              = 0x8D95
	T                                                          = 0x2001
	ALPHA_INTEGER                                              = 0x8D97
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	SGIX_ycrcb                                                 = 1
	OR                                                         = 0x1507
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	VARIABLE_B_NV                                              = 0x8524
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	OP_POWER_EXT                                               = 0x8793
	CLAMP_READ_COLOR                                           = 0x891C
	RGB2_EXT                                                   = 0x804E
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	INTENSITY32F_ARB                                           = 0x8817
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	COMPRESSED_R11_EAC                                         = 0x9270
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	OP_MOV_EXT                                                 = 0x8799
	SLUMINANCE8_EXT                                            = 0x8C47
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	KEEP                                                       = 0x1E00
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	R32F                                                       = 0x822E
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	RGBA8UI_EXT                                                = 0x8D7C
	S                                                          = 0x2000
	MODELVIEW28_ARB                                            = 0x873C
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	TEXTURE31                                                  = 0x84DF
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	CURRENT_TANGENT_EXT                                        = 0x843B
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	INCR_WRAP                                                  = 0x8507
	COMBINER0_NV                                               = 0x8550
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	PERFMON_RESULT_AMD                                         = 0x8BC6
	FUNC_SUBTRACT                                              = 0x800A
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	CONSTANT_BORDER_HP                                         = 0x8151
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	QUAD_STRIP                                                 = 0x0008
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	SEPARABLE_2D_EXT                                           = 0x8012
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	FRAGMENT_SHADER_ATI                                        = 0x8920
	INT_VEC2_ARB                                               = 0x8B53
	UNSIGNALED                                                 = 0x9118
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	PROGRAM_STRING_NV                                          = 0x8628
	PROGRAM_PARAMETER_NV                                       = 0x8644
	DYNAMIC_ATI                                                = 0x8761
	STATIC_DRAW                                                = 0x88E4
	PALETTE4_RGBA4_OES                                         = 0x8B93
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	LERP_ATI                                                   = 0x8969
	REPLACE_OLDEST_SUN                                         = 0x0003
	REG_18_ATI                                                 = 0x8933
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	COLOR_ARRAY_LIST_IBM                                       = 103072
	STENCIL_EXT                                                = 0x1802
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	RGBA4_OES                                                  = 0x8056
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	TEXTURE_COORD_NV                                           = 0x8C79
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	DEPTH_COMPONENT                                            = 0x1902
	INDEX_ARRAY_EXT                                            = 0x8077
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	TEXTURE21                                                  = 0x84D5
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	STENCIL_INDEX16_EXT                                        = 0x8D49
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	ALPHA_SCALE                                                = 0x0D1C
	SET                                                        = 0x150F
	PROJECTION                                                 = 0x1701
	SPRITE_AXIS_SGIX                                           = 0x814A
	SLIM10U_SGIX                                               = 0x831E
	PROGRAM_STRING_ARB                                         = 0x8628
	HILO16_NV                                                  = 0x86F8
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	GL_3DC_X_AMD                                               = 0x87F9
	DOUBLE_MAT3x4                                              = 0x8F4C
	PATH_FILL_MASK_NV                                          = 0x9081
	REDUCE                                                     = 0x8016
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	SOURCE1_ALPHA_EXT                                          = 0x8589
	COUNTER_TYPE_AMD                                           = 0x8BC0
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	SRGB_ALPHA_EXT                                             = 0x8C42
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	RGB10_EXT                                                  = 0x8052
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	COLOR_ATTACHMENT10                                         = 0x8CEA
	INT_10_10_10_2_OES                                         = 0x8DF7
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	FOG                                                        = 0x0B60
	BLEND_COLOR                                                = 0x8005
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	TEXTURE_MAX_LEVEL                                          = 0x813D
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	DONT_CARE                                                  = 0x1100
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	COORD_REPLACE_NV                                           = 0x8862
	TEXTURE14                                                  = 0x84CE
	OPERAND1_RGB                                               = 0x8591
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	STENCIL_INDEX1_EXT                                         = 0x8D46
	COLOR3_BIT_PGI                                             = 0x00010000
	DRAW_BUFFER15                                              = 0x8834
	MATRIX21_ARB                                               = 0x88D5
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	INT_IMAGE_2D_RECT                                          = 0x905A
	X_EXT                                                      = 0x87D5
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	LUMINANCE8_ALPHA8                                          = 0x8045
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	VERSION_1_3                                                = 1
	TEXTURE_HEIGHT                                             = 0x1001
	GL_4PASS_1_EXT                                             = 0x80A5
	MATRIX6_NV                                                 = 0x8636
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	CW                                                         = 0x0900
	TEXTURE_COMPONENTS                                         = 0x1003
	COLOR_INDEX2_EXT                                           = 0x80E3
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	INTERLACE_OML                                              = 0x8980
	UNIFORM_TYPE                                               = 0x8A37
	SATURATE_BIT_ATI                                           = 0x00000040
	OBJECT_TYPE_ARB                                            = 0x8B4E
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	EXT_texture3D                                              = 1
	STENCIL_BUFFER_BIT                                         = 0x00000400
	T2F_V3F                                                    = 0x2A27
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	SRGB_DECODE_ARB                                            = 0x8299
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	ALPHA32F_ARB                                               = 0x8816
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	AND_INVERTED                                               = 0x1504
	DEBUG_TYPE_OTHER                                           = 0x8251
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	RGBA16UI                                                   = 0x8D76
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	BLEND_DST_RGB_EXT                                          = 0x80C8
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	SINGLE_COLOR                                               = 0x81F9
	BUFFER                                                     = 0x82E0
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	SAMPLES_PASSED                                             = 0x8914
	CURRENT_PROGRAM                                            = 0x8B8D
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	VERTEX_SUBROUTINE                                          = 0x92E8
	EXT_subtexture                                             = 1
	SIGNED_INTENSITY8_NV                                       = 0x8708
	COLOR_ATTACHMENT6                                          = 0x8CE6
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	WAIT_FAILED_APPLE                                          = 0x911D
	ONE                                                        = 1
	COLOR_ARRAY_POINTER                                        = 0x8090
	RGB8UI_EXT                                                 = 0x8D7D
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	PROXY_TEXTURE_1D                                           = 0x8063
	EYE_POINT_SGIS                                             = 0x81F4
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	MIN                                                        = 0x8007
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	TEXTURE17_ARB                                              = 0x84D1
	HI_BIAS_NV                                                 = 0x8714
	RENDERBUFFER_EXT                                           = 0x8D41
	T4F_V4F                                                    = 0x2A28
	QUAD_ALPHA4_SGIS                                           = 0x811E
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	CON_10_ATI                                                 = 0x894B
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	INT_SAMPLER_3D                                             = 0x8DCB
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	REG_21_ATI                                                 = 0x8936
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	FRONT_AND_BACK                                             = 0x0408
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	ALPHA_MAX_SGIX                                             = 0x8321
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	DEPTH_STENCIL_MESA                                         = 0x8750
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	STENCIL_INDEX8                                             = 0x8D48
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	CONSTANT_COLOR_EXT                                         = 0x8001
	TEXTURE_GATHER                                             = 0x82A2
	HI_SCALE_NV                                                = 0x870E
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	DYNAMIC_READ                                               = 0x88E9
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	RGB16F_EXT                                                 = 0x881B
	SRGB_ALPHA                                                 = 0x8C42
	SAMPLER_2D_RECT                                            = 0x8B63
	SHADER_COMPILER                                            = 0x8DFA
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	SCALAR_EXT                                                 = 0x87BE
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	COLOR_ARRAY_TYPE                                           = 0x8082
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	CLIP_PLANE3                                                = 0x3003
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	RGB4_EXT                                                   = 0x804F
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	GREATER                                                    = 0x0204
	FOG_START                                                  = 0x0B63
	HALF_APPLE                                                 = 0x140B
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	TEXTURE_1D_ARRAY                                           = 0x8C18
	SLIM8U_SGIX                                                = 0x831D
	INTERPOLATE_EXT                                            = 0x8575
	BUFFER_MAPPED                                              = 0x88BC
	REG_16_ATI                                                 = 0x8931
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	SHADER_BINARY_DMP                                          = 0x9250
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	RGB32F                                                     = 0x8815
	NORMAL_ARRAY_POINTER                                       = 0x808F
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	CONVOLUTION_FORMAT                                         = 0x8017
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	RGBA32I                                                    = 0x8D82
	ROUND_NV                                                   = 0x90A4
	LINE_SMOOTH                                                = 0x0B20
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	CLIP_PLANE2                                                = 0x3002
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	SOURCE0_RGB_EXT                                            = 0x8580
	REG_29_ATI                                                 = 0x893E
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	CON_4_ATI                                                  = 0x8945
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	TESS_GEN_SPACING                                           = 0x8E77
	SPECULAR                                                   = 0x1202
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	OP_ADD_EXT                                                 = 0x8787
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	MODELVIEW18_ARB                                            = 0x8732
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	TANGENT_ARRAY_EXT                                          = 0x8439
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	FIELDS_NV                                                  = 0x8E27
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	TEXTURE24                                                  = 0x84D8
	SIGNED_RGB_NV                                              = 0x86FE
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	DOUBLE_MAT2x3                                              = 0x8F49
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	BLEND_EQUATION_EXT                                         = 0x8009
	COMBINER_BIAS_NV                                           = 0x8549
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	BLEND_COLOR_EXT                                            = 0x8005
	FUNC_ADD_OES                                               = 0x8006
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	MAP_READ_BIT_EXT                                           = 0x0001
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	TEXTURE12_ARB                                              = 0x84CC
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	RED_EXT                                                    = 0x1903
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	IUI_V3F_EXT                                                = 0x81AE
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	REG_27_ATI                                                 = 0x893C
	RG8_EXT                                                    = 0x822B
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	POINT_SPRITE                                               = 0x8861
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	MAP_WRITE_BIT_EXT                                          = 0x0002
	ONE_MINUS_DST_COLOR                                        = 0x0307
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	TEXTURE23                                                  = 0x84D7
	CLIP_PLANE5                                                = 0x3005
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	DOT3_RGB_EXT                                               = 0x8740
	PURGEABLE_APPLE                                            = 0x8A1D
	FRACTIONAL_EVEN                                            = 0x8E7C
	VERTEX_STREAM0_ATI                                         = 0x876C
	PIXEL_COUNT_NV                                             = 0x8866
	SAMPLE_MASK_NV                                             = 0x8E51
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	FLOAT_RGBA_NV                                              = 0x8883
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	BUFFER_DATA_SIZE                                           = 0x9303
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	FIXED_ONLY                                                 = 0x891D
	DOUBLE_MAT4_EXT                                            = 0x8F48
	BUFFER_BINDING                                             = 0x9302
	COMPILE_AND_EXECUTE                                        = 0x1301
	CONVOLUTION_1D                                             = 0x8010
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	FRAGMENT_SHADER                                            = 0x8B30
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	RGB_INTEGER                                                = 0x8D98
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	OP_RECIP_SQRT_EXT                                          = 0x8795
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	PRESENT_TIME_NV                                            = 0x8E2A
	TESS_EVALUATION_SHADER                                     = 0x8E87
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	SAMPLER                                                    = 0x82E6
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	MEDIUM_FLOAT                                               = 0x8DF1
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	OBJECT_LINEAR                                              = 0x2401
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	FLOAT16_VEC4_NV                                            = 0x8FFB
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	INTENSITY8                                                 = 0x804B
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	MATRIX16_ARB                                               = 0x88D0
	LUMINANCE16I_EXT                                           = 0x8D8C
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	CURRENT_COLOR                                              = 0x0B00
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	MITER_TRUNCATE_NV                                          = 0x90A8
	READ_PIXELS                                                = 0x828C
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	COMBINE                                                    = 0x8570
	INTERPOLATE_ARB                                            = 0x8575
	TRACE_NAME_MESA                                            = 0x8756
	SAMPLER_1D_ARB                                             = 0x8B5D
	UNSIGNED_INT8_NV                                           = 0x8FEC
	POINT_SIZE                                                 = 0x0B11
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	T2F_C3F_V3F                                                = 0x2A2A
	MODELVIEW8_ARB                                             = 0x8728
	MAX_TEXTURE_SIZE                                           = 0x0D33
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	PROGRAM_OUTPUT                                             = 0x92E4
	SGIX_texture_scale_bias                                    = 1
	STENCIL_FUNC                                               = 0x0B92
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	MAP1_COLOR_4                                               = 0x0D90
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	RGB16UI                                                    = 0x8D77
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	VERTEX23_BIT_PGI                                           = 0x00000004
	TEXTURE_WRAP_T                                             = 0x2803
	VARIABLE_E_NV                                              = 0x8527
	OPERAND0_ALPHA                                             = 0x8598
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	COPY                                                       = 0x1503
	SOURCE1_ALPHA_ARB                                          = 0x8589
	OP_LOG_BASE_2_EXT                                          = 0x8792
	MATRIX22_ARB                                               = 0x88D6
	FLOAT_MAT2_ARB                                             = 0x8B5A
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	IUI_N3F_V3F_EXT                                            = 0x81B0
	TEXTURE_LOD_BIAS                                           = 0x8501
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	LIGHT0                                                     = 0x4000
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	BLEND_EQUATION                                             = 0x8009
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	SOURCE2_ALPHA_EXT                                          = 0x858A
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	SGIS_texture_select                                        = 1
	PRESERVE_ATI                                               = 0x8762
	COORD_REPLACE                                              = 0x8862
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	IMAGE_2D_ARRAY                                             = 0x9053
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	RG_EXT                                                     = 0x8227
	TEXTURE0                                                   = 0x84C0
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	DSDT_NV                                                    = 0x86F5
	NO_ERROR                                                   = 0
	GL_4PASS_3_SGIS                                            = 0x80A7
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	ALPHA4                                                     = 0x803B
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	VERTEX_STREAM1_ATI                                         = 0x876D
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	CONSTANT_COLOR1_NV                                         = 0x852B
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	TESSELLATION_MODE_AMD                                      = 0x9004
	COLOR_EXT                                                  = 0x1800
	T2F_C4UB_V3F                                               = 0x2A29
	NORMAL_MAP_ARB                                             = 0x8511
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	CON_28_ATI                                                 = 0x895D
	REDUCE_EXT                                                 = 0x8016
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	RGBA8UI                                                    = 0x8D7C
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	EMBOSS_CONSTANT_NV                                         = 0x855E
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	CON_22_ATI                                                 = 0x8957
	INT_SAMPLER_CUBE                                           = 0x8DCC
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	SYNC_CONDITION                                             = 0x9113
	TEXTURE20_ARB                                              = 0x84D4
	TEXTURE21_ARB                                              = 0x84D5
	SRC2_RGB                                                   = 0x8582
	SOURCE2_ALPHA                                              = 0x858A
	CON_2_ATI                                                  = 0x8943
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	TEXTURE_GEN_S                                              = 0x0C60
	VERSION                                                    = 0x1F02
	LUMINANCE16_ALPHA16                                        = 0x8048
	POLYGON_OFFSET_EXT                                         = 0x8037
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	REG_7_ATI                                                  = 0x8928
	RGB8I_EXT                                                  = 0x8D8F
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	MAX_LIST_NESTING                                           = 0x0B31
	TEXTURE7_ARB                                               = 0x84C7
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	READ_WRITE                                                 = 0x88BA
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	MODELVIEW27_ARB                                            = 0x873B
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	DRAW_BUFFER5_ARB                                           = 0x882A
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	RGB12_EXT                                                  = 0x8053
	MODELVIEW16_ARB                                            = 0x8730
	DRAW_BUFFER9_ATI                                           = 0x882E
	STATIC_COPY_ARB                                            = 0x88E6
	SLUMINANCE_EXT                                             = 0x8C46
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	SPOT_CUTOFF                                                = 0x1206
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	TEXTURE6                                                   = 0x84C6
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	OP_DOT3_EXT                                                = 0x8784
	CCW                                                        = 0x0901
	EYE_LINEAR                                                 = 0x2400
	MATRIX4_ARB                                                = 0x88C4
	FRAME_NV                                                   = 0x8E26
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	SAMPLER_CUBE                                               = 0x8B60
	FLOAT_MAT3x2                                               = 0x8B67
	EXT_packed_pixels                                          = 1
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	TEXTURE_BLUE_SIZE                                          = 0x805E
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	COLOR_MATERIAL                                             = 0x0B57
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	BUFFER_SIZE                                                = 0x8764
	OP_FLOOR_EXT                                               = 0x878F
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	QUERY_RESULT_ARB                                           = 0x8866
	LUMINANCE_SNORM                                            = 0x9011
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	MAP2_BINORMAL_EXT                                          = 0x8447
	OPERAND2_RGB                                               = 0x8592
	CURRENT_QUERY_EXT                                          = 0x8865
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	VERTEX_SHADER                                              = 0x8B31
	BOUNDING_BOX_NV                                            = 0x908D
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	MOVE_TO_NV                                                 = 0x02
	RESTART_PATH_NV                                            = 0xF0
	VERTEX4_BIT_PGI                                            = 0x00000008
	UNSIGNED_INT_24_8                                          = 0x84FA
	CON_15_ATI                                                 = 0x8950
	GEOMETRY_SHADER                                            = 0x8DD9
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	DSDT_MAG_NV                                                = 0x86F6
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	FLOAT_VEC2                                                 = 0x8B50
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	OR_INVERTED                                                = 0x150D
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	LESS                                                       = 0x0201
	TEXTURE_BINDING_3D                                         = 0x806A
	SRGB_READ                                                  = 0x8297
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	RGB8                                                       = 0x8051
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	PASS_THROUGH_NV                                            = 0x86E6
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	CURRENT_VERTEX_EXT                                         = 0x87E2
	FILE_NAME_NV                                               = 0x9074
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	CONVOLUTION_HEIGHT                                         = 0x8019
	ALPHA12                                                    = 0x803D
	CUBIC_EXT                                                  = 0x8334
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	RGB_FLOAT32_ATI                                            = 0x8815
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	INT_IMAGE_2D_EXT                                           = 0x9058
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	TRIANGLE_STRIP                                             = 0x0005
	INDEX_CLEAR_VALUE                                          = 0x0C20
	VERTEX_ARRAY_SIZE                                          = 0x807A
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	OP_MAX_EXT                                                 = 0x878A
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	MATRIX20_ARB                                               = 0x88D4
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	VERSION_2_1                                                = 1
	FALSE                                                      = 0
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	SYNC_FLAGS_APPLE                                           = 0x9115
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	RG8I                                                       = 0x8237
	NORMAL_MAP_NV                                              = 0x8511
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	MATRIX3_ARB                                                = 0x88C3
	SLUMINANCE8                                                = 0x8C47
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	OP_NEGATE_EXT                                              = 0x8783
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	MATRIX19_ARB                                               = 0x88D3
	REG_30_ATI                                                 = 0x893F
	VERTEX_ARRAY_TYPE                                          = 0x807B
	GL_1PASS_EXT                                               = 0x80A1
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	COMBINE4_NV                                                = 0x8503
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	BOOL_VEC3_ARB                                              = 0x8B58
	CURRENT_NORMAL                                             = 0x0B02
	SAMPLES_EXT                                                = 0x80A9
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	TEXTURE27_ARB                                              = 0x84DB
	BUFFER_MAP_POINTER                                         = 0x88BD
	MATRIX15_ARB                                               = 0x88CF
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	HIGH_INT                                                   = 0x8DF5
	FLOAT_MAT4                                                 = 0x8B5C
	FLOAT_MAT4x3                                               = 0x8B6A
	LIGHT3                                                     = 0x4003
	VIEW_CLASS_64_BITS                                         = 0x82C6
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	DRAW_BUFFER3_ATI                                           = 0x8828
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	IMAGE_BUFFER                                               = 0x9051
	SGIX_list_priority                                         = 1
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	DRAW_BUFFER2_ATI                                           = 0x8827
	MATRIX_PALETTE_ARB                                         = 0x8840
	CON_25_ATI                                                 = 0x895A
	CON_31_ATI                                                 = 0x8960
	LINEAR                                                     = 0x2601
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	COMBINER7_NV                                               = 0x8557
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	MIRRORED_REPEAT                                            = 0x8370
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	RED_INTEGER                                                = 0x8D94
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	SGIS_sharpen_texture                                       = 1
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	MIRROR_CLAMP_EXT                                           = 0x8742
	RGB565                                                     = 0x8D62
	RGBA16_SNORM                                               = 0x8F9B
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	COMBINER_SCALE_NV                                          = 0x8548
	EYE_RADIAL_NV                                              = 0x855B
	VERTEX_STREAM7_ATI                                         = 0x8773
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	REG_11_ATI                                                 = 0x892C
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	WEIGHT_ARRAY_OES                                           = 0x86AD
	SGIS_generate_mipmap                                       = 1
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	REFERENCE_PLANE_SGIX                                       = 0x817D
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	IMAGE_CUBE_EXT                                             = 0x9050
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	ARRAY_SIZE                                                 = 0x92FB
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	OPERAND1_RGB_ARB                                           = 0x8591
	RGB_FLOAT32_APPLE                                          = 0x8815
	COLOR_SAMPLES_NV                                           = 0x8E20
	TEXTURE_1D                                                 = 0x0DE0
	CUBIC_HP                                                   = 0x815F
	DOUBLE                                                     = 0x140A
	NOR                                                        = 0x1508
	DEBUG_TYPE_ERROR                                           = 0x824C
	DRAW_BUFFER0_ATI                                           = 0x8825
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	SWIZZLE_STRQ_ATI                                           = 0x897A
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	EXT_blend_subtract                                         = 1
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	TRANSLATE_2D_NV                                            = 0x9090
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	RENDERER                                                   = 0x1F01
	LIGHT2                                                     = 0x4002
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	EXPAND_NORMAL_NV                                           = 0x8538
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	LARGE_CW_ARC_TO_NV                                         = 0x18
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	RENDER                                                     = 0x1C00
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	PATH_FORMAT_PS_NV                                          = 0x9071
	SCISSOR_BOX                                                = 0x0C10
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	OP_MADD_EXT                                                = 0x8788
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	TYPE                                                       = 0x92FA
	INDEX_SHIFT                                                = 0x0D12
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	PROXY_COLOR_TABLE                                          = 0x80D3
	R32UI                                                      = 0x8236
	DEPTH_STENCIL_OES                                          = 0x84F9
	PERTURB_EXT                                                = 0x85AE
	REG_5_ATI                                                  = 0x8926
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	DISCRETE_AMD                                               = 0x9006
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	COMBINER2_NV                                               = 0x8552
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	AUX_BUFFERS                                                = 0x0C00
	RGBA4_DXT5_S3TC                                            = 0x83A5
	SKIP_COMPONENTS3_NV                                        = -4
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	STENCIL_COMPONENTS                                         = 0x8285
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	DEPTH_WRITEMASK                                            = 0x0B72
	DEPTH_COMPONENTS                                           = 0x8284
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	SGIX_pixel_tiles                                           = 1
	DEPTH_STENCIL                                              = 0x84F9
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	COMPUTE_SHADER_BIT                                         = 0x00000020
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	COLOR_SUM                                                  = 0x8458
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	CLEAR                                                      = 0x1500
	CMYKA_EXT                                                  = 0x800D
	RGB_SCALE_ARB                                              = 0x8573
	TESS_CONTROL_TEXTURE                                       = 0x829C
	PROGRAM                                                    = 0x82E2
	MIRRORED_REPEAT_OES                                        = 0x8370
	MODELVIEW6_ARB                                             = 0x8726
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	EXT_texture                                                = 1
	LIGHTING                                                   = 0x0B50
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	MODELVIEW22_ARB                                            = 0x8736
	TEXTURE5_ARB                                               = 0x84C5
	INT_VEC3_ARB                                               = 0x8B54
	LINE_STIPPLE                                               = 0x0B24
	OR_REVERSE                                                 = 0x150B
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	OCCLUSION_TEST_HP                                          = 0x8165
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	BOOL_VEC3                                                  = 0x8B58
	LUMINANCE8                                                 = 0x8040
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	SGX_BINARY_IMG                                             = 0x8C0A
	QUERY_NO_WAIT_NV                                           = 0x8E14
	TEXTURE30                                                  = 0x84DE
	REG_15_ATI                                                 = 0x8930
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	COMPRESSED_LUMINANCE                                       = 0x84EA
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	SGIX_texture_multi_buffer                                  = 1
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	DUAL_ALPHA4_SGIS                                           = 0x8110
	DYNAMIC_COPY                                               = 0x88EA
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	COLOR_ATTACHMENT1                                          = 0x8CE1
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	FUNC_SUBTRACT_OES                                          = 0x800A
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	SIGNED_ALPHA8_NV                                           = 0x8706
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	MATRIX12_ARB                                               = 0x88CC
	DEBUG_SEVERITY_LOW                                         = 0x9148
	COMPRESSED_ALPHA                                           = 0x84E9
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	DUAL_ALPHA12_SGIS                                          = 0x8112
	CONSTANT_COLOR0_NV                                         = 0x852A
	VIBRANCE_BIAS_NV                                           = 0x8719
	RGB_FLOAT16_ATI                                            = 0x881B
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	CON_26_ATI                                                 = 0x895B
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	REPEAT                                                     = 0x2901
	FOG_COORDINATE_ARRAY                                       = 0x8457
	TEXTURE3_ARB                                               = 0x84C3
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	RED_MIN_CLAMP_INGR                                         = 0x8560
	SOURCE3_RGB_NV                                             = 0x8583
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	RGBA2_EXT                                                  = 0x8055
	FLOAT_MAT3x4                                               = 0x8B68
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	CURRENT_TIME_NV                                            = 0x8E28
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	COLOR_TABLE_BIAS                                           = 0x80D7
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	DEPTH_ATTACHMENT                                           = 0x8D00
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	ZOOM_X                                                     = 0x0D16
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	NEGATIVE_X_EXT                                             = 0x87D9
	CON_19_ATI                                                 = 0x8954
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	R16                                                        = 0x822A
	TEXTURE4                                                   = 0x84C4
	COMBINE_ARB                                                = 0x8570
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	LOCATION_INDEX                                             = 0x930F
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	LIST_INDEX                                                 = 0x0B33
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	NEGATIVE_Y_EXT                                             = 0x87DA
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	TEXTURE_BORDER                                             = 0x1005
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	DOUBLE_MAT2                                                = 0x8F46
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	VALIDATE_STATUS                                            = 0x8B83
	TESS_GEN_POINT_MODE                                        = 0x8E79
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	SAMPLE_BUFFERS                                             = 0x80A8
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	RED_INTEGER_EXT                                            = 0x8D94
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	RELATIVE_MOVE_TO_NV                                        = 0x03
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	IGNORE_BORDER_HP                                           = 0x8150
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	VERTEX_PROGRAM_ARB                                         = 0x8620
	OP_ROUND_EXT                                               = 0x8790
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	SGIX_depth_texture                                         = 1
	COLOR_ARRAY_SIZE                                           = 0x8081
	SIGNED_IDENTITY_NV                                         = 0x853C
	MODELVIEW23_ARB                                            = 0x8737
	UNSIGNALED_APPLE                                           = 0x9118
	IUI_V2F_EXT                                                = 0x81AD
	COMBINE_RGB                                                = 0x8571
	VERTEX_SOURCE_ATI                                          = 0x8774
	CON_21_ATI                                                 = 0x8956
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	RGB32UI                                                    = 0x8D71
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	TEXTURE24_ARB                                              = 0x84D8
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	MULTISAMPLE_BIT                                            = 0x20000000
	RESCALE_NORMAL                                             = 0x803A
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	SOURCE1_RGB_EXT                                            = 0x8581
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	REG_1_ATI                                                  = 0x8922
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	MAX_CLIP_DISTANCES                                         = 0x0D32
	EDGE_FLAG_ARRAY                                            = 0x8079
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	DISCARD_ATI                                                = 0x8763
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	PATCHES                                                    = 0x000E
	QUAD_ALPHA8_SGIS                                           = 0x811F
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	INT64_VEC4_NV                                              = 0x8FEB
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	SEPARABLE_2D                                               = 0x8012
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	HIGH_FLOAT                                                 = 0x8DF2
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	COMPUTE_SUBROUTINE                                         = 0x92ED
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	PROGRAM_PIPELINE                                           = 0x82E4
	Z_EXT                                                      = 0x87D7
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	SAMPLER_2D_SHADOW                                          = 0x8B62
	COLOR_ATTACHMENT7                                          = 0x8CE7
	RGB16_SNORM                                                = 0x8F9A
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	INT64_VEC2_NV                                              = 0x8FE9
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	BITMAP                                                     = 0x1A00
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	MAX_WIDTH                                                  = 0x827E
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	MATRIX6_ARB                                                = 0x88C6
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	INTENSITY16UI_EXT                                          = 0x8D79
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	ACCUM_BUFFER_BIT                                           = 0x00000200
	ENABLE_BIT                                                 = 0x00002000
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	COLOR_ATTACHMENT5                                          = 0x8CE5
	LINES_ADJACENCY_EXT                                        = 0x000A
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	TEXTURE1                                                   = 0x84C1
	DRAW_BUFFER15_ATI                                          = 0x8834
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	TEXTURE29_ARB                                              = 0x84DD
	DATA_BUFFER_AMD                                            = 0x9151
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	SGIX_shadow_ambient                                        = 1
	DEPTH_EXT                                                  = 0x1801
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	MAX_SAMPLES_NV                                             = 0x8D57
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	ALL_STATIC_DATA_IBM                                        = 103060
	R                                                          = 0x2002
	C4F_N3F_V3F                                                = 0x2A26
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	VERTEX_PROGRAM_NV                                          = 0x8620
	DT_SCALE_NV                                                = 0x8711
	FLOAT_RGB_NV                                               = 0x8882
	GREEN_BITS                                                 = 0x0D53
	TEXTURE4_ARB                                               = 0x84C4
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	RGBA4_EXT                                                  = 0x8056
	COLOR_RENDERABLE                                           = 0x8286
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	MATRIX0_NV                                                 = 0x8630
	MODELVIEW15_ARB                                            = 0x872F
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	EXP2                                                       = 0x0801
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	IMAGE_2D_EXT                                               = 0x904D
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	BUFFER_MAPPED_ARB                                          = 0x88BC
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	SAMPLER_2D                                                 = 0x8B5E
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	RGB32I_EXT                                                 = 0x8D83
	INT8_VEC3_NV                                               = 0x8FE2
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	COMBINE_RGB_ARB                                            = 0x8571
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	PREVIOUS_EXT                                               = 0x8578
	MATRIX8_ARB                                                = 0x88C8
	BOOL_VEC4_ARB                                              = 0x8B59
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	ALPHA_FLOAT16_ATI                                          = 0x881C
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	INCR_WRAP_EXT                                              = 0x8507
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	LAYER_NV                                                   = 0x8DAA
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	TRANSFORM_BIT                                              = 0x00001000
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	MODELVIEW_MATRIX                                           = 0x0BA6
	SAMPLE_COVERAGE                                            = 0x80A0
	WRAP_BORDER_SUN                                            = 0x81D4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	DEBUG_OBJECT_MESA                                          = 0x8759
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	VERTEX_STREAM2_ATI                                         = 0x876E
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	UTF8_NV                                                    = 0x909A
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	MATRIX_MODE                                                = 0x0BA0
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	MATRIX4_NV                                                 = 0x8634
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	TEXTURE25_ARB                                              = 0x84D9
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	FLOAT_RG32_NV                                              = 0x8887
	SGIS_texture_edge_clamp                                    = 1
	SCISSOR_TEST                                               = 0x0C11
	TEXTURE_BORDER_COLOR                                       = 0x1004
	COLOR_INDEXES                                              = 0x1603
	CLIP_PLANE1                                                = 0x3001
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	GL_1PASS_SGIS                                              = 0x80A1
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	INT_SAMPLER_1D                                             = 0x8DC9
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	UNIFORM_SIZE                                               = 0x8A38
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	MATRIX10_ARB                                               = 0x88CA
	MODULATE_COLOR_IMG                                         = 0x8C04
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	COLOR_TABLE_SGI                                            = 0x80D0
	COLOR_INDEX12_EXT                                          = 0x80E6
	DUAL_ALPHA8_SGIS                                           = 0x8111
	RG8UI                                                      = 0x8238
	VERSION_1_1                                                = 1
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	COMPRESSED_RGBA                                            = 0x84EE
	IDENTITY_NV                                                = 0x862A
	SAMPLER_CUBE_ARB                                           = 0x8B60
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	R8_EXT                                                     = 0x8229
	CURRENT_FOG_COORDINATE                                     = 0x8453
	DRAW_BUFFER15_ARB                                          = 0x8834
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	MAX_NAME_LENGTH                                            = 0x92F6
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	FIXED_ONLY_ARB                                             = 0x891D
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	IMAGE_2D_RECT_EXT                                          = 0x904F
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	SRC_ALPHA_SATURATE                                         = 0x0308
	BLEND_SRC                                                  = 0x0BE1
	GL_2PASS_0_SGIS                                            = 0x80A2
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	INT_IMAGE_1D_EXT                                           = 0x9057
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	SMALL_CW_ARC_TO_NV                                         = 0x14
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	FRONT                                                      = 0x0404
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	LUMINANCE16_SNORM                                          = 0x9019
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	SGIX_scalebias_hint                                        = 1
	PACK_ROW_LENGTH                                            = 0x0D02
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	BLUE_BIAS                                                  = 0x0D1B
	SPOT_DIRECTION                                             = 0x1204
	RGBA12                                                     = 0x805A
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	COLOR_ENCODING                                             = 0x8296
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	DS_BIAS_NV                                                 = 0x8716
	MOV_ATI                                                    = 0x8961
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	SGIX_texture_coordinate_clamp                              = 1
	LINE_BIT                                                   = 0x00000004
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	RGBA16UI_EXT                                               = 0x8D76
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	TRUE                                                       = 1
	UNSIGNED_INVERT_NV                                         = 0x8537
	STORAGE_SHARED_APPLE                                       = 0x85BF
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	R16UI                                                      = 0x8234
	REPLACE_VALUE_AMD                                          = 0x874B
	GL_3DC_XY_AMD                                              = 0x87FA
	CND_ATI                                                    = 0x896A
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	FLOAT_R32_NV                                               = 0x8885
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	RGBA_S3TC                                                  = 0x83A2
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	SAMPLER_3D_ARB                                             = 0x8B5F
	COLOR_ATTACHMENT3                                          = 0x8CE3
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	ONE_EXT                                                    = 0x87DE
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	REG_28_ATI                                                 = 0x893D
	LUMINANCE8UI_EXT                                           = 0x8D80
	TEXTURE_BINDING_1D                                         = 0x8068
	DECR_WRAP_OES                                              = 0x8508
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	INDEX_ARRAY_POINTER                                        = 0x8091
	POLYGON_MODE                                               = 0x0B40
	DRAW_BUFFER_EXT                                            = 0x0C01
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	MODELVIEW1_EXT                                             = 0x850A
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	RGB_FLOAT16_APPLE                                          = 0x881B
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	IMAGE_CUBE                                                 = 0x9050
	RGBA2                                                      = 0x8055
	VIEW_CLASS_32_BITS                                         = 0x82C8
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	C4UB_V3F                                                   = 0x2A23
	LIGHT5                                                     = 0x4005
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	INT_IMAGE_CUBE                                             = 0x905B
	SIGNALED_APPLE                                             = 0x9119
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	COMPRESSED_RGB                                             = 0x84ED
	COPY_READ_BUFFER                                           = 0x8F36
	VERSION_3_1                                                = 1
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	HISTOGRAM_RED_SIZE                                         = 0x8028
	RGBA_FLOAT16_ATI                                           = 0x881A
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	COLOR_TABLE_SCALE                                          = 0x80D6
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	CON_0_ATI                                                  = 0x8941
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	VERSION_3_0                                                = 1
	MAP_READ_BIT                                               = 0x0001
	DITHER                                                     = 0x0BD0
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	CURRENT_MATRIX_NV                                          = 0x8641
	DOT3_RGB                                                   = 0x86AE
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	NEXT_BUFFER_NV                                             = -2
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	DRAW_BUFFER14                                              = 0x8833
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	SOURCE2_RGB                                                = 0x8582
	DRAW_BUFFER13_NV                                           = 0x8832
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	INDEX_BIT_PGI                                              = 0x00080000
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	ALPHA16I_EXT                                               = 0x8D8A
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	BGR_EXT                                                    = 0x80E0
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	EMISSION                                                   = 0x1600
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	MAX_VERTEX_STREAMS                                         = 0x8E71
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	ALPHA16_SNORM                                              = 0x9018
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	NUM_EXTENSIONS                                             = 0x821D
	TEXTURE15                                                  = 0x84CF
	PACK_ALIGNMENT                                             = 0x0D05
	TEXTURE_3D_OES                                             = 0x806F
	MODELVIEW13_ARB                                            = 0x872D
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	FRAMEBUFFER_OES                                            = 0x8D40
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	INVALID_OPERATION                                          = 0x0502
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	FOG_COORD                                                  = 0x8451
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	POLYGON_STIPPLE                                            = 0x0B42
	TEXTURE_COORD_ARRAY                                        = 0x8078
	POINT                                                      = 0x1B00
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	OUTPUT_FOG_EXT                                             = 0x87BD
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	RGB32UI_EXT                                                = 0x8D71
	SAMPLER_3D_OES                                             = 0x8B5F
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	TEXTURE_GREEN_SIZE                                         = 0x805D
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	LUMINANCE32I_EXT                                           = 0x8D86
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	SRC_COLOR                                                  = 0x0300
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	DEPTH_CLAMP                                                = 0x864F
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	TRIANGLE_LIST_SUN                                          = 0x81D7
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MAP1_VERTEX_4                                              = 0x0D98
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	PROGRAM_BINDING_ARB                                        = 0x8677
	FLOAT_MAT3_ARB                                             = 0x8B5B
	INFO_LOG_LENGTH                                            = 0x8B84
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	CURRENT_RASTER_INDEX                                       = 0x0B05
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	PATH_JOIN_STYLE_NV                                         = 0x9079
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	TEXTURE_GEN_R                                              = 0x0C62
	SOURCE0_ALPHA_EXT                                          = 0x8588
	QUARTER_BIT_ATI                                            = 0x00000010
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	MAP1_TANGENT_EXT                                           = 0x8444
	TEXTURE11                                                  = 0x84CB
	OPERAND0_RGB_EXT                                           = 0x8590
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	MODELVIEW20_ARB                                            = 0x8734
	E_TIMES_F_NV                                               = 0x8531
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	EYE_PLANE                                                  = 0x2502
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	CON_30_ATI                                                 = 0x895F
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	LUMINANCE16F_ARB                                           = 0x881E
	POINT_SPRITE_NV                                            = 0x8861
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	TEXTURE_DEPTH                                              = 0x8071
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	ACTIVE_TEXTURE                                             = 0x84E0
	SET_AMD                                                    = 0x874A
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	REG_26_ATI                                                 = 0x893B
	ALPHA8UI_EXT                                               = 0x8D7E
	PATH_STROKE_MASK_NV                                        = 0x9084
	PHONG_HINT_WIN                                             = 0x80EB
	TEXTURE_MAX_LOD                                            = 0x813B
	MATRIX5_NV                                                 = 0x8635
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	READ_ONLY_ARB                                              = 0x88B8
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	R16F                                                       = 0x822D
	SCREEN_COORDINATES_REND                                    = 0x8490
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	SIGNED_INTENSITY_NV                                        = 0x8707
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	READ_BUFFER                                                = 0x0C02
	RED_SCALE                                                  = 0x0D14
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	DEPTH24_STENCIL8                                           = 0x88F0
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	MAX_SAMPLES                                                = 0x8D57
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	SGIS_multisample                                           = 1
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	PROGRAM_TARGET_NV                                          = 0x8646
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	PATCH_VERTICES                                             = 0x8E72
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	POLYGON_SMOOTH                                             = 0x0B41
	MAP1_BINORMAL_EXT                                          = 0x8446
	MATRIX9_ARB                                                = 0x88C9
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	ALPHA32UI_EXT                                              = 0x8D72
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	CURRENT_BINORMAL_EXT                                       = 0x843C
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	SKIP_DECODE_EXT                                            = 0x8A4A
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	QUERY_COUNTER_BITS                                         = 0x8864
	R3_G3_B2                                                   = 0x2A10
	HISTOGRAM_SINK_EXT                                         = 0x802D
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	YCBCR_422_APPLE                                            = 0x85B9
	MATRIX29_ARB                                               = 0x88DD
	DOUBLE_MAT4x3                                              = 0x8F4E
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	POINT_SMOOTH                                               = 0x0B10
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	DRAW_BUFFER12                                              = 0x8831
	COMPILE_STATUS                                             = 0x8B81
	TESS_GEN_MODE                                              = 0x8E76
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	MODELVIEW1_ARB                                             = 0x850A
	HILO8_NV                                                   = 0x885E
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	SHADER_IMAGE_STORE                                         = 0x82A5
	MODELVIEW3_ARB                                             = 0x8723
	VERTEX_STREAM4_ATI                                         = 0x8770
	LAYOUT_DEFAULT_INTEL                                       = 0
	LINES                                                      = 0x0001
	CURRENT_RASTER_COLOR                                       = 0x0B04
	TEXTURE_MIN_LOD                                            = 0x813A
	TEXTURE26_ARB                                              = 0x84DA
	MATRIX13_ARB                                               = 0x88CD
	REG_23_ATI                                                 = 0x8938
	QUERY_OBJECT_EXT                                           = 0x9153
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	LIST_BASE                                                  = 0x0B32
	TABLE_TOO_LARGE                                            = 0x8031
	TRANSPOSE_NV                                               = 0x862C
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	BYTE                                                       = 0x1400
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	RED_BIAS                                                   = 0x0D15
	SIGNED_LUMINANCE_NV                                        = 0x8701
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	READ_PIXELS_FORMAT                                         = 0x828D
	TEXTURE_SHADOW                                             = 0x82A1
	RGBA4_S3TC                                                 = 0x83A3
	Y_EXT                                                      = 0x87D6
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	FRAMEZOOM_SGIX                                             = 0x818B
	MIRROR_CLAMP_ATI                                           = 0x8742
	LUMINANCE16_EXT                                            = 0x8042
	COLOR_ATTACHMENT9                                          = 0x8CE9
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	R16F_EXT                                                   = 0x822D
	RELATIVE_LINE_TO_NV                                        = 0x05
	EXT_point_parameters                                       = 1
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	BLEND_DST_RGB_OES                                          = 0x80C8
	REPLICATE_BORDER                                           = 0x8153
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	ABGR_EXT                                                   = 0x8000
	GENERATE_MIPMAP                                            = 0x8191
	RG16                                                       = 0x822C
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	STENCIL_INDEX8_OES                                         = 0x8D48
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	UNIFORM_BUFFER_START                                       = 0x8A29
	BOOL_VEC2_ARB                                              = 0x8B57
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	BITMAP_TOKEN                                               = 0x0704
	TEXTURE_COMPRESSED                                         = 0x86A1
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	QUERY_OBJECT_AMD                                           = 0x9153
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	INT64_NV                                                   = 0x140E
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	REPLICATE_BORDER_HP                                        = 0x8153
	IMAGE_1D_EXT                                               = 0x904C
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	UNIFORM_BUFFER                                             = 0x8A11
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	EXT_texture_object                                         = 1
	COLOR_INDEX                                                = 0x1900
	SGIX_fragment_lighting                                     = 1
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	FILTER                                                     = 0x829A
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	COLOR_TABLE                                                = 0x80D0
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	COLOR4_BIT_PGI                                             = 0x00020000
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	INCR                                                       = 0x1E02
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	MAX_SHININESS_NV                                           = 0x8504
	NEGATIVE_Z_EXT                                             = 0x87DB
	DRAW_BUFFER2                                               = 0x8827
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	TIME_ELAPSED                                               = 0x88BF
	INDEX_ARRAY_LIST_IBM                                       = 103073
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	TEXTURE27                                                  = 0x84DB
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	WAIT_FAILED                                                = 0x911D
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	SIGNED_RGBA_NV                                             = 0x86FB
	STREAM_COPY                                                = 0x88E2
	RETAINED_APPLE                                             = 0x8A1B
	PATH_MITER_LIMIT_NV                                        = 0x907A
	EXT_blend_minmax                                           = 1
	TEXTURE19                                                  = 0x84D3
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	GL_422_AVERAGE_EXT                                         = 0x80CE
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	CONTEXT_FLAGS                                              = 0x821E
	MATRIX5_ARB                                                = 0x88C5
	REG_13_ATI                                                 = 0x892E
	FIRST_TO_REST_NV                                           = 0x90AF
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	LIST_BIT                                                   = 0x00020000
	CLIP_DISTANCE0                                             = 0x3000
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	SOURCE0_ALPHA                                              = 0x8588
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	STATIC_COPY                                                = 0x88E6
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	LINE_SMOOTH_HINT                                           = 0x0C52
	ACCUM_RED_BITS                                             = 0x0D58
	TEXTURE_PRIORITY                                           = 0x8066
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	COUNT_UP_NV                                                = 0x9088
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	NONE                                                       = 0
	XOR                                                        = 0x1506
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	RGBA32F                                                    = 0x8814
	SAMPLER_1D                                                 = 0x8B5D
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	CURRENT_MATRIX_ARB                                         = 0x8641
	CLAMP_VERTEX_COLOR                                         = 0x891A
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	INT_IMAGE_3D_EXT                                           = 0x9059
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	ADD_SIGNED_EXT                                             = 0x8574
	MITER_REVERT_NV                                            = 0x90A7
	ALPHA8                                                     = 0x803C
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	NOTEQUAL                                                   = 0x0205
	INVALID_VALUE                                              = 0x0501
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	PRIMARY_COLOR_EXT                                          = 0x8577
	DEPTH_COMPONENT32F                                         = 0x8CAC
	DEPTH_RENDERABLE                                           = 0x8287
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	CON_16_ATI                                                 = 0x8951
	LIGHT6                                                     = 0x4006
	VERTEX_TEXTURE                                             = 0x829B
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	VIDEO_BUFFER_NV                                            = 0x9020
	GL_422_EXT                                                 = 0x80CC
	SCALE_BY_TWO_NV                                            = 0x853E
	CON_29_ATI                                                 = 0x895E
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	RESCALE_NORMAL_EXT                                         = 0x803A
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	CONDITION_SATISFIED                                        = 0x911C
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	QUADS                                                      = 0x0007
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	OFFSET                                                     = 0x92FC
	RGBA8                                                      = 0x8058
	INTENSITY_EXT                                              = 0x8049
	BLEND_SRC_RGB_OES                                          = 0x80C9
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	STENCIL_BACK_FUNC                                          = 0x8800
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	COLOR_WRITEMASK                                            = 0x0C23
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	SIGNED_HILO16_NV                                           = 0x86FA
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	FLOAT_RGB32_NV                                             = 0x8889
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	MODELVIEW31_ARB                                            = 0x873F
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	LINE_STRIP                                                 = 0x0003
	SPRITE_MODE_SGIX                                           = 0x8149
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	MAX_DEPTH                                                  = 0x8280
	TEXTURE_WRAP_R                                             = 0x8072
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	DEBUG_TYPE_MARKER                                          = 0x8268
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	INDEX_ARRAY_TYPE                                           = 0x8085
	RESTART_SUN                                                = 0x0001
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	READ_ONLY                                                  = 0x88B8
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	MATRIX7_NV                                                 = 0x8637
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	FOG_COLOR                                                  = 0x0B66
	SYNC_CL_EVENT_ARB                                          = 0x8240
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	COMPRESSED_SRGB                                            = 0x8C48
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	AUX3                                                       = 0x040C
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	GL_3D                                                      = 0x0601
	LUMINANCE12_ALPHA12                                        = 0x8047
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	IMAGE_3D_EXT                                               = 0x904E
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	TEXTURE_BINDING_2D                                         = 0x8069
	MODELVIEW11_ARB                                            = 0x872B
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	DECR_WRAP                                                  = 0x8508
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	VIEW_CLASS_24_BITS                                         = 0x82C9
	YCRCBA_SGIX                                                = 0x8319
	REG_20_ATI                                                 = 0x8935
	RG_SNORM                                                   = 0x8F91
	TEXTURE_MAG_FILTER                                         = 0x2800
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	VERTEX_SHADER_EXT                                          = 0x8780
	REG_8_ATI                                                  = 0x8929
	POLYGON_OFFSET_FILL                                        = 0x8037
	RED                                                        = 0x1903
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	SRGB                                                       = 0x8C40
	MAX_SAMPLES_EXT                                            = 0x8D57
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	RGB16UI_EXT                                                = 0x8D77
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	MATRIX2_ARB                                                = 0x88C2
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	DELETE_STATUS                                              = 0x8B80
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	CONSTANT_ATTENUATION                                       = 0x1207
	OPERAND3_RGB_NV                                            = 0x8593
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	RGB10_A2                                                   = 0x8059
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	VARIABLE_D_NV                                              = 0x8526
	COMBINE_ALPHA_EXT                                          = 0x8572
	MATRIX24_ARB                                               = 0x88D8
	STEREO                                                     = 0x0C33
	RGB8I                                                      = 0x8D8F
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	SHADER_OBJECT_EXT                                          = 0x8B48
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	RGB4_S3TC                                                  = 0x83A1
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	NEGATE_BIT_ATI                                             = 0x00000004
	TRANSFORM_FEEDBACK                                         = 0x8E22
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	DOUBLE_MAT4x2                                              = 0x8F4D
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	RGBA16I                                                    = 0x8D88
	LOW_FLOAT                                                  = 0x8DF0
	MEDIUM_INT                                                 = 0x8DF4
	INT16_NV                                                   = 0x8FE4
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	DRAW_BUFFER11                                              = 0x8830
	SHADER_TYPE                                                = 0x8B4F
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	MVP_MATRIX_EXT                                             = 0x87E3
	RGBA16F_EXT                                                = 0x881A
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	DEPTH_BITS                                                 = 0x0D56
	CLIP_DISTANCE4                                             = 0x3004
	SPRITE_SGIX                                                = 0x8148
	OPERAND2_ALPHA                                             = 0x859A
	DOT3_RGB_ARB                                               = 0x86AE
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	ADJACENT_PAIRS_NV                                          = 0x90AE
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	INDEX_WRITEMASK                                            = 0x0C21
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	DRAW_BUFFER10_ATI                                          = 0x882F
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	PREVIOUS_ARB                                               = 0x8578
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	COLOR_ATTACHMENT0                                          = 0x8CE0
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	POLYGON_OFFSET_LINE                                        = 0x2A02
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	INTENSITY8_SNORM                                           = 0x9017
	PACK_INVERT_MESA                                           = 0x8758
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	VARIABLE_F_NV                                              = 0x8528
	TRACE_MASK_MESA                                            = 0x8755
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	DOUBLE_EXT                                                 = 0x140A
	TEXTURE19_ARB                                              = 0x84D3
	RED_MAX_CLAMP_INGR                                         = 0x8564
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	GPU_ADDRESS_NV                                             = 0x8F34
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	MINMAX_FORMAT_EXT                                          = 0x802F
	COMPUTE_TEXTURE                                            = 0x82A0
	MODELVIEW9_ARB                                             = 0x8729
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	T2F_N3F_V3F                                                = 0x2A2B
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	RGB8_EXT                                                   = 0x8051
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	VECTOR_EXT                                                 = 0x87BF
	LOCAL_EXT                                                  = 0x87C4
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	COLOR_ARRAY                                                = 0x8076
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	SIGNED_RGB8_NV                                             = 0x86FF
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	QUAD_MESH_SUN                                              = 0x8614
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	SLUMINANCE8_NV                                             = 0x8C47
	INTENSITY8I_EXT                                            = 0x8D91
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	COLOR_INDEX1_EXT                                           = 0x80E2
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	VERTEX_SHADER_ARB                                          = 0x8B31
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	OBJECT_POINT_SGIS                                          = 0x81F5
	MAX_VIEWPORTS                                              = 0x825B
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	VARIABLE_G_NV                                              = 0x8529
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	STENCIL_FAIL                                               = 0x0B94
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	DOUBLE_MAT2x4                                              = 0x8F4A
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	MINOR_VERSION                                              = 0x821C
	FOG_COORDINATE                                             = 0x8451
	DRAW_BUFFER7_ARB                                           = 0x882C
	PATH_GEN_MODE_NV                                           = 0x90B0
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	MATRIX7_ARB                                                = 0x88C7
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	POINT_SIZE_MAX_ARB                                         = 0x8127
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	COMP_BIT_ATI                                               = 0x00000002
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	ARRAY_BUFFER_ARB                                           = 0x8892
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	HILO_NV                                                    = 0x86F4
	SWIZZLE_STQ_ATI                                            = 0x8977
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	ADD_ATI                                                    = 0x8963
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	CLIP_DISTANCE_NV                                           = 0x8C7A
	RGBA_INTEGER                                               = 0x8D99
	ALPHA_TEST                                                 = 0x0BC0
	PACK_IMAGE_HEIGHT                                          = 0x806C
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	COMBINER1_NV                                               = 0x8551
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	PROGRAM_POINT_SIZE                                         = 0x8642
	CND0_ATI                                                   = 0x896B
	INT_IMAGE_3D                                               = 0x9059
	RGBA4                                                      = 0x8056
	RGB10_A2_EXT                                               = 0x8059
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	R1UI_V3F_SUN                                               = 0x85C4
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	INT_VEC3                                                   = 0x8B54
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	STENCIL_INDEX16                                            = 0x8D49
	SGIX_blend_alpha_minmax                                    = 1
	TEXTURE8                                                   = 0x84C8
	INVERSE_NV                                                 = 0x862B
	VERTEX_STREAM3_ATI                                         = 0x876F
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	TEXTURE_GEN_T                                              = 0x0C61
	CURRENT_ATTRIB_NV                                          = 0x8626
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	STENCIL_BACK_FAIL                                          = 0x8801
	QUERY_WAIT_NV                                              = 0x8E13
	UNSIGNED_INT16_NV                                          = 0x8FF0
	SGIX_impact_pixel_texture                                  = 1
	FOG_END                                                    = 0x0B64
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	MULTISAMPLE_SGIS                                           = 0x809D
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	OP_RECIP_EXT                                               = 0x8794
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	LIST_PRIORITY_SGIX                                         = 0x8182
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	ALPHA8I_EXT                                                = 0x8D90
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	COMBINER6_NV                                               = 0x8556
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	EXTENSIONS                                                 = 0x1F03
	DT_BIAS_NV                                                 = 0x8717
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	PROXY_HISTOGRAM                                            = 0x8025
	INTENSITY12_EXT                                            = 0x804C
	NEGATIVE_ONE_EXT                                           = 0x87DF
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	UNPACK_SKIP_IMAGES                                         = 0x806D
	DOT3_RGBA_ARB                                              = 0x86AF
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	OUTPUT_VERTEX_EXT                                          = 0x879A
	DRAW_BUFFER5                                               = 0x882A
	POINT_SPRITE_ARB                                           = 0x8861
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	T2F_IUI_V2F_EXT                                            = 0x81B1
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	OP_SET_GE_EXT                                              = 0x878C
	DRAW_BUFFER10_NV                                           = 0x882F
	RGBA32UI_EXT                                               = 0x8D70
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	DOUBLEBUFFER                                               = 0x0C32
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	TEXTURE31_ARB                                              = 0x84DF
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	REG_19_ATI                                                 = 0x8934
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	PRIMITIVE_RESTART                                          = 0x8F9D
	SGIS_detail_texture                                        = 1
	IMAGE_SCALE_X_HP                                           = 0x8155
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	QUERY_RESULT_EXT                                           = 0x8866
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	MODELVIEW4_ARB                                             = 0x8724
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	SAMPLER_BUFFER                                             = 0x8DC2
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	PATH_COORD_COUNT_NV                                        = 0x909E
	MAP2_NORMAL                                                = 0x0DB2
	TEXTURE9_ARB                                               = 0x84C9
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	MAX_IMAGE_UNITS                                            = 0x8F38
	SGIX_interlace                                             = 1
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	DSDT8_NV                                                   = 0x8709
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	TEXTURE10_ARB                                              = 0x84CA
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	LO_SCALE_NV                                                = 0x870F
	MODELVIEW0_ARB                                             = 0x1700
	SAMPLES_3DFX                                               = 0x86B4
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	EXT_copy_texture                                           = 1
	RGBA12_EXT                                                 = 0x805A
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	VARIANT_EXT                                                = 0x87C1
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	DRAW_BUFFER6                                               = 0x882B
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	DOUBLE_VEC3                                                = 0x8FFD
	STENCIL_BITS                                               = 0x0D57
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	SIGNED_HILO_NV                                             = 0x86F9
	BUFFER_ACCESS_OES                                          = 0x88BB
	POINT_BIT                                                  = 0x00000002
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	GL_3_BYTES                                                 = 0x1408
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	SHADER_BINARY_VIV                                          = 0x8FC4
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	ADD                                                        = 0x0104
	MAP_COLOR                                                  = 0x0D10
	BLUE_SCALE                                                 = 0x0D1A
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	TEXTURE_VIEW                                               = 0x82B5
	TEXTURE0_ARB                                               = 0x84C0
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	TEXTURE28_ARB                                              = 0x84DC
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	RGBA_SNORM                                                 = 0x8F93
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	INTENSITY32UI_EXT                                          = 0x8D73
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	QUERY_BUFFER_AMD                                           = 0x9192
	TEXTURE_GEN_MODE                                           = 0x2500
	DEBUG_SOURCE_API                                           = 0x8246
	OP_SUB_EXT                                                 = 0x8796
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	LOAD                                                       = 0x0101
	MATRIX0_ARB                                                = 0x88C0
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	COPY_WRITE_BUFFER                                          = 0x8F37
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	ATTENUATION_EXT                                            = 0x834D
	MATRIX1_ARB                                                = 0x88C1
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	CURRENT_BIT                                                = 0x00000001
	LINE_TOKEN                                                 = 0x0702
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	TRANSLATE_X_NV                                             = 0x908E
	CONSTANT_ALPHA                                             = 0x8003
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	FRAMEBUFFER_BLEND                                          = 0x828B
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	INTENSITY16I_EXT                                           = 0x8D8B
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	LUMINANCE16                                                = 0x8042
	RGB16_EXT                                                  = 0x8054
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	PROVOKING_VERTEX                                           = 0x8E4F
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	MATRIX3_NV                                                 = 0x8633
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	AND_REVERSE                                                = 0x1502
	CMYK_EXT                                                   = 0x800C
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	MODELVIEW26_ARB                                            = 0x873A
	DYNAMIC_READ_ARB                                           = 0x88E9
	FLOAT_VEC3_ARB                                             = 0x8B51
	SAMPLES_ARB                                                = 0x80A9
	YCRCB_422_SGIX                                             = 0x81BB
	RASTERIZER_DISCARD                                         = 0x8C89
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	CLIP_DISTANCE5                                             = 0x3005
	LUMINANCE8_EXT                                             = 0x8040
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	INVARIANT_VALUE_EXT                                        = 0x87EA
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	PACK_SWAP_BYTES                                            = 0x0D00
	LINEAR_DETAIL_SGIS                                         = 0x8097
	MAX_TEXTURE_UNITS                                          = 0x84E2
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	CLIP_DISTANCE3                                             = 0x3003
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	TIME_ELAPSED_EXT                                           = 0x88BF
	COLOR_CLEAR_VALUE                                          = 0x0C22
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	SAMPLES_PASSED_ARB                                         = 0x8914
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	TEXTURE11_ARB                                              = 0x84CB
	INCR_WRAP_OES                                              = 0x8507
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	CON_20_ATI                                                 = 0x8955
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	MAX_SAMPLES_IMG                                            = 0x9135
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	MATRIX18_ARB                                               = 0x88D2
	SAMPLER_3D                                                 = 0x8B5F
	BGRA_INTEGER                                               = 0x8D9B
	CLOSE_PATH_NV                                              = 0x00
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	CON_1_ATI                                                  = 0x8942
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	FAILURE_NV                                                 = 0x9030
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	CULL_VERTEX_IBM                                            = 103050
	DEPTH                                                      = 0x1801
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	SGIX_fog_offset                                            = 1
	COLOR_BUFFER_BIT                                           = 0x00004000
	SUBPIXEL_BITS                                              = 0x0D50
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	Z400_BINARY_AMD                                            = 0x8740
	DRAW_BUFFER14_ARB                                          = 0x8833
	BLUE_BIT_ATI                                               = 0x00000004
	FRAMEBUFFER_EXT                                            = 0x8D40
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	MAX_EXT                                                    = 0x8008
	TEXTURE17                                                  = 0x84D1
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	COLOR_TABLE_WIDTH                                          = 0x80D9
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	TRIANGLE_MESH_SUN                                          = 0x8615
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	STACK_OVERFLOW                                             = 0x0503
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	DEPTH_CLAMP_NV                                             = 0x864F
	CON_5_ATI                                                  = 0x8946
	CON_17_ATI                                                 = 0x8952
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	ZERO_EXT                                                   = 0x87DD
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	EQUAL                                                      = 0x0202
	PACK_SKIP_PIXELS                                           = 0x0D04
	NORMAL_MAP_OES                                             = 0x8511
	ACCUM_GREEN_BITS                                           = 0x0D59
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	VERSION_1_2                                                = 1
	SOURCE2_ALPHA_ARB                                          = 0x858A
	SIGNED_RGBA8_NV                                            = 0x86FC
	AUX2                                                       = 0x040B
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	WRITE_ONLY_ARB                                             = 0x88B9
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	NORMAL_ARRAY_TYPE                                          = 0x807E
	REG_2_ATI                                                  = 0x8923
	PATH_GEN_COEFF_NV                                          = 0x90B1
	COMPUTE_SHADER                                             = 0x91B9
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	FULL_SUPPORT                                               = 0x82B7
	DRAW_BUFFER12_NV                                           = 0x8831
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	FUNC_ADD                                                   = 0x8006
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	DRAW_BUFFER11_NV                                           = 0x8830
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	LUMINANCE4                                                 = 0x803F
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	MODELVIEW                                                  = 0x1700
	DRAW_BUFFER12_ATI                                          = 0x8831
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	SGIX_pixel_texture                                         = 1
	COLOR_MATERIAL_FACE                                        = 0x0B55
	DECODE_EXT                                                 = 0x8A49
	FLOAT_MAT4_ARB                                             = 0x8B5C
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	POLYGON_OFFSET_POINT                                       = 0x2A01
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	HINT_BIT                                                   = 0x00008000
	REFLECTION_MAP_EXT                                         = 0x8512
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	SGIX_vertex_preclip                                        = 1
	CONSTANT_COLOR                                             = 0x8001
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	EXP                                                        = 0x0800
	RGBA16_EXT                                                 = 0x805B
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	SCALE_BY_FOUR_NV                                           = 0x853F
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	RGBA_FLOAT32_ATI                                           = 0x8814
	DRAW_BUFFER8                                               = 0x882D
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	VERTEX_ID_NV                                               = 0x8C7B
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	CON_13_ATI                                                 = 0x894E
	RELEASED_APPLE                                             = 0x8A19
	PATH_STENCIL_REF_NV                                        = 0x90B8
	FRAGMENT_TEXTURE                                           = 0x829F
	MODELVIEW12_ARB                                            = 0x872C
	MAGNITUDE_SCALE_NV                                         = 0x8712
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	AMBIENT                                                    = 0x1200
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	QUERY_RESULT                                               = 0x8866
	RELATIVE_ARC_TO_NV                                         = 0xFF
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	CON_18_ATI                                                 = 0x8953
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	INVARIANT_EXT                                              = 0x87C2
	SKIP_COMPONENTS1_NV                                        = -6
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	NORMAL_MAP_EXT                                             = 0x8511
	COLOR_ATTACHMENT15                                         = 0x8CEF
	EXT_shared_texture_palette                                 = 1
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	ALPHA4_EXT                                                 = 0x803B
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	SUB_ATI                                                    = 0x8965
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	INTENSITY12                                                = 0x804C
	RGB_SCALE                                                  = 0x8573
	STREAM_READ                                                = 0x88E1
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	FRAMEBUFFER                                                = 0x8D40
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	RGB12                                                      = 0x8053
	OP_FRAC_EXT                                                = 0x8789
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	STENCIL_INDEX1                                             = 0x8D46
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	TEXTURE1_ARB                                               = 0x84C1
	VERSION_2_0                                                = 1
	SGIS_point_parameters                                      = 1
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	PROGRAM_LENGTH_ARB                                         = 0x8627
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	DSDT_MAG_VIB_NV                                            = 0x86F7
	FRAGMENT_DEPTH                                             = 0x8452
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	UNDEFINED_APPLE                                            = 0x8A1C
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	NEAREST                                                    = 0x2600
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	DRAW_BUFFER11_ATI                                          = 0x8830
	SAMPLE_MASK_VALUE                                          = 0x8E52
	SGIS_pixel_texture                                         = 1
	UNSIGNED_INT64_NV                                          = 0x140F
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	UNIFORM_BLOCK                                              = 0x92E2
	DS_SCALE_NV                                                = 0x8710
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	FOG_COORDINATE_EXT                                         = 0x8451
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	SGIX_async_histogram                                       = 1
	EIGHTH_BIT_ATI                                             = 0x00000020
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	STACK_UNDERFLOW                                            = 0x0504
	INDEX                                                      = 0x8222
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	ORDER                                                      = 0x0A01
	RGB10                                                      = 0x8052
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	GL_4X_BIT_ATI                                              = 0x00000002
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	MAP_WRITE_BIT                                              = 0x0002
	COLOR_ARRAY_STRIDE                                         = 0x8083
	NAND                                                       = 0x150E
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	STREAM_DRAW_ARB                                            = 0x88E0
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	FOG_OFFSET_SGIX                                            = 0x8198
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	RGBA_MODE                                                  = 0x0C31
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	TEXTURE_3D_EXT                                             = 0x806F
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	OPERAND1_ALPHA_EXT                                         = 0x8599
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	DEPTH32F_STENCIL8                                          = 0x8CAD
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	SOURCE0_RGB                                                = 0x8580
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	MUL_ATI                                                    = 0x8964
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	COLOR_MATRIX                                               = 0x80B1
	RG32UI                                                     = 0x823C
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	NUM_SAMPLE_COUNTS                                          = 0x9380
	AUX1                                                       = 0x040A
	GL_3D_COLOR                                                = 0x0602
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	SGIS_texture_border_clamp                                  = 1
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	DRAW_BUFFER7_NV                                            = 0x882C
	REG_31_ATI                                                 = 0x8940
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	VENDOR                                                     = 0x1F00
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	EXT_blend_logic_op                                         = 1
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	LUMINANCE32F_ARB                                           = 0x8818
	FLOAT_MAT2x3                                               = 0x8B65
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	LINE_STRIP_ADJACENCY                                       = 0x000B
	DEPTH_BIAS                                                 = 0x0D1F
	BLEND_SRC_RGB                                              = 0x80C9
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	MAX_IMAGE_SAMPLES                                          = 0x906D
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	TRIANGLE_FAN                                               = 0x0006
	STENCIL                                                    = 0x1802
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	VIEW_CLASS_96_BITS                                         = 0x82C5
	TEXTURE14_ARB                                              = 0x84CE
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	COUNT_DOWN_NV                                              = 0x9089
	TEXTURE5                                                   = 0x84C5
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	COLOR_INDEX8_EXT                                           = 0x80E5
	OPERAND1_RGB_EXT                                           = 0x8591
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	SRGB8_ALPHA8                                               = 0x8C43
	CUBIC_CURVE_TO_NV                                          = 0x0C
	YCRCB_SGIX                                                 = 0x8318
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	COUNTER_RANGE_AMD                                          = 0x8BC1
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	STENCIL_REF                                                = 0x0B97
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	RGBA16F                                                    = 0x881A
	FOG_HINT                                                   = 0x0C54
	CLAMP_TO_BORDER_NV                                         = 0x812D
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	ARRAY_BUFFER_BINDING                                       = 0x8894
	IMAGE_SCALE_Y_HP                                           = 0x8156
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	POINT_SIZE_MIN_ARB                                         = 0x8126
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	OP_CLAMP_EXT                                               = 0x878E
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	FENCE_CONDITION_NV                                         = 0x84F4
	TEXTURE_RECTANGLE                                          = 0x84F5
	RGB32F_ARB                                                 = 0x8815
	POINT_SIZE_RANGE                                           = 0x0B12
	SOURCE3_ALPHA_NV                                           = 0x858B
	CULL_MODES_NV                                              = 0x86E0
	COLOR_ATTACHMENT11                                         = 0x8CEB
	EXT_rescale_normal                                         = 1
	ACCUM_BLUE_BITS                                            = 0x0D5A
	GL_4PASS_1_SGIS                                            = 0x80A5
	DOUBLE_VEC2                                                = 0x8FFC
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	TRANSLATE_Y_NV                                             = 0x908F
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	RGB                                                        = 0x1907
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	DEBUG_OUTPUT                                               = 0x92E0
	TEXTURE_ENV_MODE                                           = 0x2200
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	COMBINER_MAPPING_NV                                        = 0x8543
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	RGBA8I                                                     = 0x8D8E
	CULL_FACE                                                  = 0x0B44
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	PALETTE8_RGBA4_OES                                         = 0x8B98
	RGBA8I_EXT                                                 = 0x8D8E
	LIGHTING_BIT                                               = 0x00000040
	STATIC_DRAW_ARB                                            = 0x88E4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	LIGHT7                                                     = 0x4007
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	SPARE1_NV                                                  = 0x852F
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	RGBA32F_ARB                                                = 0x8814
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	SURFACE_REGISTERED_NV                                      = 0x86FD
	PALETTE4_RGB8_OES                                          = 0x8B90
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	SPARE0_NV                                                  = 0x852E
	WRITE_ONLY_OES                                             = 0x88B9
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	MAD_ATI                                                    = 0x8968
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	MAX_LIGHTS                                                 = 0x0D31
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	DRAW_BUFFER3_ARB                                           = 0x8828
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	UNPACK_ALIGNMENT                                           = 0x0CF5
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	CONVOLUTION_WIDTH                                          = 0x8018
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	NUM_PASSES_ATI                                             = 0x8970
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	POLYGON_TOKEN                                              = 0x0703
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	RG16I                                                      = 0x8239
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	COMBINER5_NV                                               = 0x8555
	SLICE_ACCUM_SUN                                            = 0x85CC
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	CON_8_ATI                                                  = 0x8949
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	QUERY                                                      = 0x82E3
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	SPOT_EXPONENT                                              = 0x1205
	RGB565_OES                                                 = 0x8D62
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	LUMINANCE32UI_EXT                                          = 0x8D74
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	DRAW_BUFFER4_ARB                                           = 0x8829
	DOT2_ADD_ATI                                               = 0x896C
	GL_2PASS_1_SGIS                                            = 0x80A3
	SECONDARY_COLOR_NV                                         = 0x852D
	RGBA_FLOAT32_APPLE                                         = 0x8814
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	MAX_PATCH_VERTICES                                         = 0x8E7D
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	SGIX_reference_plane                                       = 1
	INTENSITY8_EXT                                             = 0x804B
	RGB5_EXT                                                   = 0x8050
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	BUFFER_VARIABLE                                            = 0x92E5
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	ALL_COMPLETED_NV                                           = 0x84F2
	OPERAND0_RGB                                               = 0x8590
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	CLIP_PLANE4                                                = 0x3004
	FLOAT_RGBA16_NV                                            = 0x888A
	DOT4_ATI                                                   = 0x8967
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	EDGE_FLAG                                                  = 0x0B43
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	RENDERBUFFER_WIDTH                                         = 0x8D42
	LINE_LOOP                                                  = 0x0002
	POSITION                                                   = 0x1203
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	SURFACE_MAPPED_NV                                          = 0x8700
	MATRIX26_ARB                                               = 0x88DA
	IMAGE_1D                                                   = 0x904C
	SGIX_icc_texture                                           = 1
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	VERTEX_ARRAY                                               = 0x8074
	BUFFER_USAGE                                               = 0x8765
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	REG_25_ATI                                                 = 0x893A
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	INT_IMAGE_1D                                               = 0x9057
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	SQUARE_NV                                                  = 0x90A3
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	COMPRESSED_INTENSITY                                       = 0x84EC
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	R8I                                                        = 0x8231
	RG32I                                                      = 0x823B
	REFLECTION_MAP_NV                                          = 0x8512
	BLEND_DST_ALPHA                                            = 0x80CA
	REFLECTION_MAP_OES                                         = 0x8512
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	MAP2_TANGENT_EXT                                           = 0x8445
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	SUBTRACT                                                   = 0x84E7
	DOT_PRODUCT_NV                                             = 0x86EC
	DRAW_BUFFER1_NV                                            = 0x8826
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	ASYNC_MARKER_SGIX                                          = 0x8329
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	YCBAYCR8A_4224_NV                                          = 0x9032
	COMMAND_BARRIER_BIT                                        = 0x00000040
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	DRAW_BUFFER3                                               = 0x8828
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	RENDER_MODE                                                = 0x0C40
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	MULTISAMPLE_3DFX                                           = 0x86B2
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	SPHERE_MAP                                                 = 0x2402
	REGISTER_COMBINERS_NV                                      = 0x8522
	VARIABLE_C_NV                                              = 0x8525
	MATRIX11_ARB                                               = 0x88CB
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	RGB16                                                      = 0x8054
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	DRAW_BUFFER13                                              = 0x8832
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	PRIMARY_COLOR_NV                                           = 0x852C
	COLOR_ATTACHMENT12                                         = 0x8CEC
	LINES_ADJACENCY                                            = 0x000A
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	RGB8_SNORM                                                 = 0x8F96
	PIXEL_MODE_BIT                                             = 0x00000020
	BLEND_DST                                                  = 0x0BE0
	PRIMARY_COLOR_ARB                                          = 0x8577
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	NOOP                                                       = 0x1505
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	GREEN                                                      = 0x1904
	INTENSITY4                                                 = 0x804A
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	CON_14_ATI                                                 = 0x894F
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	SLUMINANCE_NV                                              = 0x8C46
	INTENSITY32I_EXT                                           = 0x8D85
	SUCCESS_NV                                                 = 0x902F
	VERTEX_SHADER_BIT                                          = 0x00000001
	RG_INTEGER                                                 = 0x8228
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	RIGHT                                                      = 0x0407
	HISTOGRAM                                                  = 0x8024
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	SRC1_RGB                                                   = 0x8581
	SYNC_FENCE                                                 = 0x9116
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	CONSTANT_ARB                                               = 0x8576
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	ALPHA_INTEGER_EXT                                          = 0x8D97
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	DRAW_BUFFER6_NV                                            = 0x882B
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	MAP2_VERTEX_4                                              = 0x0DB8
	ALPHA                                                      = 0x1906
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	SHADE_MODEL                                                = 0x0B54
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	BOOL_ARB                                                   = 0x8B56
	IMAGE_1D_ARRAY                                             = 0x9052
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	OP_EXP_BASE_2_EXT                                          = 0x8791
	STREAM_READ_ARB                                            = 0x88E1
	STENCIL_INDEX4                                             = 0x8D47
	MATRIX_STRIDE                                              = 0x92FF
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	STENCIL_BACK_REF                                           = 0x8CA3
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	RGB5_A1_EXT                                                = 0x8057
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	YCBCR_MESA                                                 = 0x8757
	ALPHA16UI_EXT                                              = 0x8D78
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	CLAMP_TO_EDGE                                              = 0x812F
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	SOURCE1_ALPHA                                              = 0x8589
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	RG16F_EXT                                                  = 0x822F
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	SHADER_OPERATION_NV                                        = 0x86DF
	MAX_TEXTURE_COORDS                                         = 0x8871
	TEXTURE18                                                  = 0x84D2
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	PACK_CMYK_HINT_EXT                                         = 0x800E
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	V3F                                                        = 0x2A21
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	GL_8X_BIT_ATI                                              = 0x00000004
	LOGIC_OP                                                   = 0x0BF1
	TEXTURE_4D_SGIS                                            = 0x8134
	COLOR_SUM_ARB                                              = 0x8458
	OP_MIN_EXT                                                 = 0x878B
	DEPTH_TEXTURE_MODE                                         = 0x884B
	DST_ALPHA                                                  = 0x0304
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	CULL_VERTEX_EXT                                            = 0x81AA
	SOURCE1_RGB_ARB                                            = 0x8581
	REG_12_ATI                                                 = 0x892D
	EVAL_BIT                                                   = 0x00010000
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	READ_WRITE_ARB                                             = 0x88BA
	INT_VEC2                                                   = 0x8B53
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	ACTIVE_RESOURCES                                           = 0x92F5
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	LUMINANCE8_SNORM                                           = 0x9015
	BLUE_BITS                                                  = 0x0D54
	UNSIGNED_BYTE                                              = 0x1401
	LINE                                                       = 0x1B01
	R16I                                                       = 0x8233
	SIGNED_ALPHA_NV                                            = 0x8705
	PACK_RESAMPLE_OML                                          = 0x8984
	ALPHA32I_EXT                                               = 0x8D84
	LOGIC_OP_MODE                                              = 0x0BF0
	UPPER_LEFT                                                 = 0x8CA2
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	ITALIC_BIT_NV                                              = 0x02
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	RGB10_A2UI                                                 = 0x906F
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	VIEW_CLASS_16_BITS                                         = 0x82CA
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	DRAW_BUFFER                                                = 0x0C01
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	UNPACK_LSB_FIRST                                           = 0x0CF1
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	LUMINANCE16UI_EXT                                          = 0x8D7A
	TRIANGULAR_NV                                              = 0x90A5
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	MAP2_VERTEX_3                                              = 0x0DB7
	TEXTURE_RESIDENT                                           = 0x8067
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	REG_9_ATI                                                  = 0x892A
	ALPHA_TEST_REF                                             = 0x0BC2
	C3F_V3F                                                    = 0x2A24
	DRAW_BUFFER15_NV                                           = 0x8834
	ARB_imaging                                                = 1
	REG_22_ATI                                                 = 0x8937
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	LIGHT4                                                     = 0x4004
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	DOT3_RGBA                                                  = 0x86AF
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	DRAW_PIXEL_TOKEN                                           = 0x0705
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	TEXTURE7                                                   = 0x84C7
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	FLOAT16_VEC2_NV                                            = 0x8FF9
	PACK_LSB_FIRST                                             = 0x0D01
	OPERAND3_ALPHA_NV                                          = 0x859B
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	POINT_SMOOTH_HINT                                          = 0x0C51
	LUMINANCE_ALPHA                                            = 0x190A
	COLOR_COMPONENTS                                           = 0x8283
	COLOR_ATTACHMENT2                                          = 0x8CE2
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	FOG_COORD_ARRAY                                            = 0x8457
	TEXTURE30_ARB                                              = 0x84DE
	ETC1_SRGB8_NV                                              = 0x88EE
	GREEN_BIT_ATI                                              = 0x00000002
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	REG_0_ATI                                                  = 0x8921
	CON_9_ATI                                                  = 0x894A
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	R8UI                                                       = 0x8232
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	MODELVIEW5_ARB                                             = 0x8725
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	DYNAMIC_DRAW                                               = 0x88E8
	HALF_FLOAT_OES                                             = 0x8D61
	RG8_SNORM                                                  = 0x8F95
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	FRONT_LEFT                                                 = 0x0400
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	CLIP_PLANE0                                                = 0x3000
	CLIP_DISTANCE1                                             = 0x3001
	TEXTURE_3D                                                 = 0x806F
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	DRAW_BUFFER8_ATI                                           = 0x882D
	STREAM_COPY_ARB                                            = 0x88E2
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	DEPTH_COMPONENT16                                          = 0x81A5
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	FLOAT_VEC4_ARB                                             = 0x8B52
	STENCIL_INDEX1_OES                                         = 0x8D46
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	NORMAL_MAP                                                 = 0x8511
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	DRAW_BUFFER10                                              = 0x882F
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	MULTIVIEW_EXT                                              = 0x90F1
	BUFFER_MAP_LENGTH                                          = 0x9120
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	DRAW_BUFFER5_NV                                            = 0x882A
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	INT                                                        = 0x1404
	MULTISAMPLE                                                = 0x809D
	TEXTURE22_ARB                                              = 0x84D6
	DRAW_BUFFER9_NV                                            = 0x882E
	PROGRAM_FORMAT_ARB                                         = 0x8876
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	ISOLINES                                                   = 0x8E7A
	INT64_VEC3_NV                                              = 0x8FEA
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	MODELVIEW17_ARB                                            = 0x8731
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	BOOL_VEC4                                                  = 0x8B59
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	BGR_INTEGER_EXT                                            = 0x8D9A
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	RGB5                                                       = 0x8050
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	CLEAR_BUFFER                                               = 0x82B4
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	MATRIX17_ARB                                               = 0x88D1
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	PATH_END_CAPS_NV                                           = 0x9076
	SGIX_subsample                                             = 1
	PROGRAM_LENGTH_NV                                          = 0x8627
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	COMBINE_RGB_EXT                                            = 0x8571
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	SRGB8_NV                                                   = 0x8C41
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	SYNC_FENCE_APPLE                                           = 0x9116
	ALPHA16                                                    = 0x803E
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	MAJOR_VERSION                                              = 0x821B
	UNDEFINED_VERTEX                                           = 0x8260
	TEXTURE26                                                  = 0x84DA
	INTERPOLATE                                                = 0x8575
	MODELVIEW14_ARB                                            = 0x872E
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	PASS_THROUGH_TOKEN                                         = 0x0700
	INDEX_LOGIC_OP                                             = 0x0BF1
	INT8_VEC2_NV                                               = 0x8FE1
	ALPHA_SNORM                                                = 0x9010
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	MODELVIEW30_ARB                                            = 0x873E
	DRAW_BUFFER9_ARB                                           = 0x882E
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	TEXTURE2_ARB                                               = 0x84C2
	RGB16I_EXT                                                 = 0x8D89
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	SGIX_polynomial_ffd                                        = 1
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	FULL_RANGE_EXT                                             = 0x87E1
	RECT_NV                                                    = 0xF6
	ALPHA_TEST_QCOM                                            = 0x0BC0
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	ALWAYS                                                     = 0x0207
	MAP2_COLOR_4                                               = 0x0DB0
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	TEXTURE10                                                  = 0x84CA
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	MAX_HEIGHT                                                 = 0x827F
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	PARTIAL_SUCCESS_NV                                         = 0x902E
	BLOCK_INDEX                                                = 0x92FD
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	MODELVIEW25_ARB                                            = 0x8739
	BUMP_ENVMAP_ATI                                            = 0x877B
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	FIXED                                                      = 0x140C
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	INT_IMAGE_BUFFER                                           = 0x905C
	AND                                                        = 0x1501
	MODULATE                                                   = 0x2100
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	INT_VEC4                                                   = 0x8B55
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	CON_3_ATI                                                  = 0x8944
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	POINT_SIZE_MIN_EXT                                         = 0x8126
	EYE_LINE_SGIS                                              = 0x81F6
	VARIANT_ARRAY_EXT                                          = 0x87E8
	SRGB8                                                      = 0x8C41
	SGI_color_table                                            = 1
	SPRITE_AXIAL_SGIX                                          = 0x814C
	DEBUG_SOURCE_OTHER                                         = 0x824B
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	ALPHA_FLOAT32_ATI                                          = 0x8816
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	CULL_FRAGMENT_NV                                           = 0x86E7
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	SGIX_texture_add_env                                       = 1
	BLEND_EQUATION_OES                                         = 0x8009
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	PREVIOUS                                                   = 0x8578
	BLEND_EQUATION_ALPHA                                       = 0x883D
	TIMESTAMP                                                  = 0x8E28
	VERSION_1_5                                                = 1
	C4UB_V2F                                                   = 0x2A22
	FILTER4_SGIS                                               = 0x8146
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	PACK_SKIP_ROWS                                             = 0x0D03
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	HISTOGRAM_EXT                                              = 0x8024
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	INDEX_ARRAY_STRIDE                                         = 0x8086
	TEXTURE_WRAP_S                                             = 0x2802
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	TEXTURE18_ARB                                              = 0x84D2
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	RGBA32UI                                                   = 0x8D70
	FOG_MODE                                                   = 0x0B65
	VIEWPORT                                                   = 0x0BA2
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	BUMP_TARGET_ATI                                            = 0x877C
	BLUE                                                       = 0x1905
	LUMINANCE4_ALPHA4                                          = 0x8043
	RGB_S3TC                                                   = 0x83A0
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	BGRA_EXT                                                   = 0x80E1
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	COMBINER4_NV                                               = 0x8554
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	STENCIL_WRITEMASK                                          = 0x0B98
	MULTISAMPLE_ARB                                            = 0x809D
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	OP_DOT4_EXT                                                = 0x8785
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	GL_4PASS_0_SGIS                                            = 0x80A4
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	INT8_NV                                                    = 0x8FE0
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	LINK_STATUS                                                = 0x8B82
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	TEXTURE_SAMPLES                                            = 0x9106
	STENCIL_VALUE_MASK                                         = 0x0B93
	STENCIL_INDEX                                              = 0x1901
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	RETURN                                                     = 0x0102
	REG_6_ATI                                                  = 0x8927
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	BACK                                                       = 0x0405
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	PATH_FORMAT_SVG_NV                                         = 0x9070
	ALPHA_BITS                                                 = 0x0D55
	COLOR_ARRAY_EXT                                            = 0x8076
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	REFLECTION_MAP_ARB                                         = 0x8512
	DRAW_BUFFER4                                               = 0x8829
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	BUFFER_SIZE_ARB                                            = 0x8764
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	QUERY_WAIT                                                 = 0x8E13
	LUMINANCE12_EXT                                            = 0x8041
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	FOG_FUNC_SGIS                                              = 0x812A
	MODELVIEW10_ARB                                            = 0x872A
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	CULL_FACE_MODE                                             = 0x0B45
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	BGR_INTEGER                                                = 0x8D9A
	PRESENT_DURATION_NV                                        = 0x8E2B
	DOMAIN                                                     = 0x0A02
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	REG_3_ATI                                                  = 0x8924
	BLUE_INTEGER                                               = 0x8D96
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	TEXTURE_LIGHT_EXT                                          = 0x8350
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	TEXTURE3                                                   = 0x84C3
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	FIXED_OES                                                  = 0x140C
	HISTOGRAM_WIDTH                                            = 0x8026
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	INVALID_INDEX                                              = 0xFFFFFFFF
	ARC_TO_NV                                                  = 0xFE
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	PN_TRIANGLES_ATI                                           = 0x87F0
	REG_10_ATI                                                 = 0x892B
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	FRACTIONAL_ODD                                             = 0x8E7B
	TRIANGLES                                                  = 0x0004
	LUMINANCE                                                  = 0x1909
	POINT_SIZE_MAX                                             = 0x8127
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	TEXTURE22                                                  = 0x84D6
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	SGIX_convolution_accuracy                                  = 1
	INDEX_MODE                                                 = 0x0C30
	RGB9_E5_EXT                                                = 0x8C3D
	TRANSLATE_3D_NV                                            = 0x9091
	RED_BITS                                                   = 0x0D52
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	UNSIGNED_INT                                               = 0x1405
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	SHADER                                                     = 0x82E1
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	CONVEX_HULL_NV                                             = 0x908B
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	MIN_EXT                                                    = 0x8007
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	SAMPLE_MASK_EXT                                            = 0x80A0
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	AVERAGE_EXT                                                = 0x8335
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	IMAGE_BINDING_NAME                                         = 0x8F3A
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	R16_SNORM                                                  = 0x8F98
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	MAP2_INDEX                                                 = 0x0DB1
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	OUTPUT_COLOR1_EXT                                          = 0x879C
	STATIC_READ                                                = 0x88E5
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	SEPARATE_ATTRIBS                                           = 0x8C8D
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	TEXTURE_SHADER_NV                                          = 0x86DE
	COLOR_ATTACHMENT14                                         = 0x8CEE
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	MODELVIEW7_ARB                                             = 0x8727
	SYNC_OBJECT_APPLE                                          = 0x8A53
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	SYNC_FLAGS                                                 = 0x9115
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	SGIX_flush_raster                                          = 1
	PROJECTION_MATRIX                                          = 0x0BA7
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	CURRENT_QUERY                                              = 0x8865
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	UNIFORM                                                    = 0x92E1
	SGIX_calligraphic_fragment                                 = 1
	READ_BUFFER_EXT                                            = 0x0C02
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	GENERATE_MIPMAP_HINT                                       = 0x8192
	TEXTURE25                                                  = 0x84D9
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	TEXTURE_BUFFER                                             = 0x8C2A
	SLUMINANCE                                                 = 0x8C46
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	SKIP_COMPONENTS2_NV                                        = -5
	MAP_STENCIL                                                = 0x0D11
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	IMAGE_2D                                                   = 0x904D
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	COMPRESSED_RG                                              = 0x8226
	PRIMITIVES_GENERATED                                       = 0x8C87
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	PATH_DASH_CAPS_NV                                          = 0x907B
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	RGBA_DXT5_S3TC                                             = 0x83A4
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	ALREADY_SIGNALED                                           = 0x911A
	FUNC_ADD_EXT                                               = 0x8006
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	ACTIVE_PROGRAM                                             = 0x8259
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	INTENSITY16F_ARB                                           = 0x881D
	MATRIX23_ARB                                               = 0x88D7
	BOLD_BIT_NV                                                = 0x01
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	COMBINER_MUX_SUM_NV                                        = 0x8547
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	CON_11_ATI                                                 = 0x894C
	AUX0                                                       = 0x0409
	WRITE_DISCARD_NV                                           = 0x88BE
	DOT3_ATI                                                   = 0x8966
	COPY_INVERTED                                              = 0x150C
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	GREEN_INTEGER_EXT                                          = 0x8D95
	TRACK_MATRIX_NV                                            = 0x8648
	DEPTH_BOUNDS_EXT                                           = 0x8891
	R8_SNORM                                                   = 0x8F94
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	CONVOLUTION_1D_EXT                                         = 0x8010
	LO_BIAS_NV                                                 = 0x8715
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	RG16UI                                                     = 0x823A
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	INTENSITY_SNORM                                            = 0x9013
	EMBOSS_MAP_NV                                              = 0x855F
	OP_INDEX_EXT                                               = 0x8782
	COLOR_ATTACHMENT13                                         = 0x8CED
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	POLYGON_BIT                                                = 0x00000008
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	STATIC_READ_ARB                                            = 0x88E5
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	SGIX_tag_sample_buffer                                     = 1
	MINMAX_EXT                                                 = 0x802E
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	EVAL_2D_NV                                                 = 0x86C0
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	CON_12_ATI                                                 = 0x894D
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	TEXTURE8_ARB                                               = 0x84C8
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	VARIANT_VALUE_EXT                                          = 0x87E4
	CON_23_ATI                                                 = 0x8958
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	SAMPLER_OBJECT_AMD                                         = 0x9155
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	RGBA8_EXT                                                  = 0x8058
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	FLOAT_MAT4x2                                               = 0x8B69
	RGBA32I_EXT                                                = 0x8D82
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	DYNAMIC_COPY_ARB                                           = 0x88EA
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	MULTISAMPLE_EXT                                            = 0x809D
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	REG_24_ATI                                                 = 0x8939
	UNPACK_RESAMPLE_OML                                        = 0x8985
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	RGBA8_SNORM                                                = 0x8F97
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	MAX_LABEL_LENGTH                                           = 0x82E8
	SUBTRACT_ARB                                               = 0x84E7
	INT8_VEC4_NV                                               = 0x8FE3
	MAX_INTEGER_SAMPLES                                        = 0x9110
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	READ_BUFFER_NV                                             = 0x0C02
	MAX_LAYERS                                                 = 0x8281
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	DRAW_BUFFER11_ARB                                          = 0x8830
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	DOUBLE_MAT2_EXT                                            = 0x8F46
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	PACK_SKIP_IMAGES                                           = 0x806B
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	VERTEX_ARRAY_EXT                                           = 0x8074
	DRAW_BUFFER12_ARB                                          = 0x8831
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	FIELD_UPPER_NV                                             = 0x9022
	COMPRESSED_RED                                             = 0x8225
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	VIBRANCE_SCALE_NV                                          = 0x8713
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	SGIX_sprite                                                = 1
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	ZOOM_Y                                                     = 0x0D17
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	FLOAT_RGBA32_NV                                            = 0x888B
	INT16_VEC3_NV                                              = 0x8FE6
	AVERAGE_HP                                                 = 0x8160
	CURRENT_QUERY_ARB                                          = 0x8865
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	FOG_COORDINATE_SOURCE                                      = 0x8450
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	MATRIX27_ARB                                               = 0x88DB
	STREAM_DRAW                                                = 0x88E0
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	RG32F                                                      = 0x8230
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	FLAT                                                       = 0x1D00
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	SRGB_WRITE                                                 = 0x8298
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	SGIS_texture_lod                                           = 1
	SELECT                                                     = 0x1C02
	CLIP_DISTANCE2                                             = 0x3002
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	DRAW_BUFFER8_NV                                            = 0x882D
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	SKIP_COMPONENTS4_NV                                        = -3
	OBJECT_PLANE                                               = 0x2501
	TEXTURE20                                                  = 0x84D4
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	DEPTH_RANGE                                                = 0x0B70
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	DEPTH_STENCIL_EXT                                          = 0x84F9
	PALETTE8_RGBA8_OES                                         = 0x8B96
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	FLOAT                                                      = 0x1406
	W_EXT                                                      = 0x87D8
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	VIEWPORT_BIT                                               = 0x00000800
	TEXTURE_GEN_Q                                              = 0x0C63
	FRAGMENT_COLOR_EXT                                         = 0x834C
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	MIRRORED_REPEAT_IBM                                        = 0x8370
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	ARRAY_BUFFER                                               = 0x8892
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	CON_6_ATI                                                  = 0x8947
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	FLOAT16_NV                                                 = 0x8FF8
	FOG_BIT                                                    = 0x00000080
	EQUIV                                                      = 0x1509
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	GL_2PASS_1_EXT                                             = 0x80A3
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	REG_14_ATI                                                 = 0x892F
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	INT_IMAGE_2D                                               = 0x9058
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	SYNC_STATUS_APPLE                                          = 0x9114
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	ZERO                                                       = 0
	REPLACE_MIDDLE_SUN                                         = 0x0002
	FLOAT_MAT2x4                                               = 0x8B66
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	SGI_texture_color_table                                    = 1
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	FLOAT_RG_NV                                                = 0x8881
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	DEBUG_PRINT_MESA                                           = 0x875A
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	SYNC_CONDITION_APPLE                                       = 0x9113
	COMBINE_ALPHA_ARB                                          = 0x8572
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	FEEDBACK                                                   = 0x1C01
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	BOOL                                                       = 0x8B56
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	LOCATION                                                   = 0x930E
	BLEND_EQUATION_RGB                                         = 0x8009
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	YCRCB_444_SGIX                                             = 0x81BC
	STENCIL_RENDERABLE                                         = 0x8288
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	POINT_SIZE_MIN                                             = 0x8126
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	FLOAT_VEC4                                                 = 0x8B52
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	FLOAT16_VEC3_NV                                            = 0x8FFA
	OPERAND2_RGB_EXT                                           = 0x8592
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	BLEND_SRC_ALPHA                                            = 0x80CB
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	VERSION_3_2                                                = 1
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	DISPLAY_LIST                                               = 0x82E7
	FLOAT_RG16_NV                                              = 0x8886
	QUADRATIC_ATTENUATION                                      = 0x1209
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	Q                                                          = 0x2003
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	GL_4PASS_2_EXT                                             = 0x80A6
	DRAW_BUFFER1                                               = 0x8826
	BUFFER_ACCESS_ARB                                          = 0x88BB
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	SRC_ALPHA                                                  = 0x0302
	SAMPLES                                                    = 0x80A9
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	INTENSITY                                                  = 0x8049
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	OBJECT_LINE_SGIS                                           = 0x81F7
	ARRAY_STRIDE                                               = 0x92FE
	LINEAR_ATTENUATION                                         = 0x1208
	GL_4_BYTES                                                 = 0x1409
	LUMINANCE6_ALPHA2                                          = 0x8044
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	OPERAND0_RGB_ARB                                           = 0x8590
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	LUMINANCE12_ALPHA4                                         = 0x8046
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	SRC0_ALPHA                                                 = 0x8588
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	CONSTANT                                                   = 0x8576
	SOURCE1_RGB                                                = 0x8581
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	MOVE_TO_RESETS_NV                                          = 0x90B5
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	MAP1_VERTEX_3                                              = 0x0D97
	COLOR_SUM_EXT                                              = 0x8458
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	MODELVIEW21_ARB                                            = 0x8735
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	EXPAND_NEGATE_NV                                           = 0x8539
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	RENDERBUFFER_OES                                           = 0x8D41
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	EXT_polygon_offset                                         = 1
	R32I                                                       = 0x8235
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	VERSION_1_4                                                = 1
	INDEX_BITS                                                 = 0x0D51
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	DEPTH_COMPONENT24                                          = 0x81A6
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	TEXTURE2                                                   = 0x84C2
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	BGRA_INTEGER_EXT                                           = 0x8D9B
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	SIGNED_NORMALIZED                                          = 0x8F9C
	COEFF                                                      = 0x0A00
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	SAMPLER_2D_ARB                                             = 0x8B5E
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	INTENSITY16_SNORM                                          = 0x901B
	EXT_cmyka                                                  = 1
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	LINE_WIDTH                                                 = 0x0B21
	TEXTURE9                                                   = 0x84C9
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	GL_2_BYTES                                                 = 0x1407
	HALF_FLOAT                                                 = 0x140B
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	SGIX_shadow                                                = 1
	POINTS                                                     = 0x0000
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	COMBINE_EXT                                                = 0x8570
	OPERAND0_ALPHA_EXT                                         = 0x8598
	NEGATIVE_W_EXT                                             = 0x87DC
	PRIMITIVE_ID_NV                                            = 0x8C7C
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	DEPTH_FUNC                                                 = 0x0B74
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	TEXTURE15_ARB                                              = 0x84CF
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	ALPHA_BIAS                                                 = 0x0D1D
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	TEXTURE12                                                  = 0x84CC
	INTENSITY4_EXT                                             = 0x804A
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	RENDERBUFFER                                               = 0x8D41
	IMAGE_BUFFER_EXT                                           = 0x9051
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	FLOAT_R_NV                                                 = 0x8880
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	R11F_G11F_B10F                                             = 0x8C3A
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	MATRIX2_NV                                                 = 0x8632
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	OP_MUL_EXT                                                 = 0x8786
	DRAW_BUFFER4_NV                                            = 0x8829
	COLOR_ATTACHMENT4                                          = 0x8CE4
	INTENSITY8UI_EXT                                           = 0x8D7F
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	MODELVIEW19_ARB                                            = 0x8733
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	LINE_RESET_TOKEN                                           = 0x0707
	TEXTURE28                                                  = 0x84DC
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	QUERY_NO_WAIT                                              = 0x8E14
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	SOURCE2_RGB_ARB                                            = 0x8582
	SGIX_instruments                                           = 1
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	LUMINANCE8I_EXT                                            = 0x8D92
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	PROXY_TEXTURE_2D                                           = 0x8064
	PROXY_TEXTURE_3D                                           = 0x8070
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	DRAW_BUFFER9                                               = 0x882E
	MATRIX28_ARB                                               = 0x88DC
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	STENCIL_INDEX8_EXT                                         = 0x8D48
	LEFT                                                       = 0x0406
	MAP1_NORMAL                                                = 0x0D92
	REFLECTION_MAP                                             = 0x8512
	SGIS_texture_filter4                                       = 1
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	MODULATE_ADD_ATI                                           = 0x8744
	SRC1_COLOR                                                 = 0x88F9
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	SIGNALED                                                   = 0x9119
	ALPHA8_EXT                                                 = 0x803C
	ATC_RGB_AMD                                                = 0x8C92
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	PALETTE4_RGBA8_OES                                         = 0x8B91
	RENDERBUFFER_BINDING                                       = 0x8CA7
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	INDEX_OFFSET                                               = 0x0D13
	TEXTURE13                                                  = 0x84CD
	FLOAT_MAT3                                                 = 0x8B5B
	ADD_BLEND_IMG                                              = 0x8C09
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	OPERAND2_RGB_ARB                                           = 0x8592
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	ATTACHED_SHADERS                                           = 0x8B85
	STENCIL_INDEX4_EXT                                         = 0x8D47
	LINE_WIDTH_RANGE                                           = 0x0B22
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	RG16_SNORM                                                 = 0x8F99
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	ETC1_RGB8_OES                                              = 0x8D64
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	LEQUAL                                                     = 0x0203
	TEXTURE                                                    = 0x1702
	HISTOGRAM_FORMAT                                           = 0x8027
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	DRAW_BUFFER13_ATI                                          = 0x8832
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	INT16_VEC2_NV                                              = 0x8FE5
	EXT_blend_color                                            = 1
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	PRIMITIVE_RESTART_NV                                       = 0x8558
	DOT3_RGBA_IMG                                              = 0x86AF
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	DRAW_BUFFER0_NV                                            = 0x8825
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	IMAGE_BINDING_FORMAT                                       = 0x906E
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	GL_4PASS_3_EXT                                             = 0x80A7
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	OP_SET_LT_EXT                                              = 0x878D
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	RGB_422_APPLE                                              = 0x8A1F
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	AFFINE_2D_NV                                               = 0x9092
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	R8                                                         = 0x8229
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	MAX_SUBROUTINES                                            = 0x8DE7
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	VOLATILE_APPLE                                             = 0x8A1A
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	DEPTH_BUFFER_BIT                                           = 0x00000100
	LINES_ADJACENCY_ARB                                        = 0x000A
	TEXTURE23_ARB                                              = 0x84D7
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	GEQUAL                                                     = 0x0206
	TEXTURE_RED_SIZE                                           = 0x805C
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	COMBINER_INPUT_NV                                          = 0x8542
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	ALPHA_MIN_SGIX                                             = 0x8320
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	DRAW_BUFFER6_ARB                                           = 0x882B
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	DOUBLE_MAT3x2                                              = 0x8F4B
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	GL_4PASS_2_SGIS                                            = 0x80A6
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	DRAW_BUFFER6_ATI                                           = 0x882B
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	SIGNED_NEGATE_NV                                           = 0x853D
	DRAW_BUFFER0                                               = 0x8825
	DRAW_BUFFER8_ARB                                           = 0x882D
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	TIMEOUT_EXPIRED                                            = 0x911B
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	INTENSITY16_EXT                                            = 0x804D
	INTERLACE_SGIX                                             = 0x8094
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	FOG_DENSITY                                                = 0x0B62
	SRC1_ALPHA                                                 = 0x8589
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	BUFFER_MAPPED_OES                                          = 0x88BC
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	CONTINUOUS_AMD                                             = 0x9007
	SMOOTH                                                     = 0x1D01
	COMBINE_ALPHA                                              = 0x8572
	SURFACE_STATE_NV                                           = 0x86EB
	BLUE_INTEGER_EXT                                           = 0x8D96
	INT_SAMPLER_2D                                             = 0x8DCA
	CONTEXT_PROFILE_MASK                                       = 0x9126
	SGIX_framezoom                                             = 1
	CLIP_DISTANCE6                                             = 0x3006
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	TEXTURE16_ARB                                              = 0x84D0
	DRAW_BUFFER14_ATI                                          = 0x8833
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	FLOAT_RGB16_NV                                             = 0x8888
	COLOR                                                      = 0x1800
	GEOMETRY_TEXTURE                                           = 0x829E
	DUDV_ATI                                                   = 0x8779
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	ELEMENT_ARRAY_ATI                                          = 0x8768
	MATRIX30_ARB                                               = 0x88DE
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	NAME_STACK_DEPTH                                           = 0x0D70
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	INDEX_MATERIAL_EXT                                         = 0x81B8
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	STANDARD_FONT_NAME_NV                                      = 0x9072
	SCISSOR_BIT                                                = 0x00080000
	FRONT_RIGHT                                                = 0x0401
	DEPTH_SCALE                                                = 0x0D1E
	CLAMP                                                      = 0x2900
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	DRAW_BUFFER7                                               = 0x882C
	CONVOLUTION_2D_EXT                                         = 0x8011
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	NORMAL_BIT_PGI                                             = 0x08000000
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	MINMAX                                                     = 0x802E
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	DRAW_BUFFER2_NV                                            = 0x8827
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	RGB_SNORM                                                  = 0x8F92
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	SINGLE_COLOR_EXT                                           = 0x81F9
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	BEVEL_NV                                                   = 0x90A6
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	OPERAND2_ALPHA_EXT                                         = 0x859A
	RG8                                                        = 0x822B
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	DRAW_BUFFER1_ATI                                           = 0x8826
	SGIX_async_pixel                                           = 1
	POLYGON                                                    = 0x0009
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	DEPTH_COMPONENT32                                          = 0x81A7
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	DECR                                                       = 0x1E03
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	OPERAND1_ALPHA                                             = 0x8599
	SOURCE0_ALPHA_ARB                                          = 0x8588
	ALPHA_TEST_FUNC                                            = 0x0BC1
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	CON_7_ATI                                                  = 0x8948
	SAMPLER_BUFFER_AMD                                         = 0x9001
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	INDEX_TEST_REF_EXT                                         = 0x81B7
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	TEXTURE13_ARB                                              = 0x84CD
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	AUTO_NORMAL                                                = 0x0D80
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	TEXTURE_COMPARE_MODE                                       = 0x884C
	SIGNED_HILO8_NV                                            = 0x885F
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	BOOL_VEC2                                                  = 0x8B57
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	MODELVIEW2_ARB                                             = 0x8722
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	NEVER                                                      = 0x0200
	TEXTURE_MIN_FILTER                                         = 0x2801
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	NORMALIZE                                                  = 0x0BA1
	MATRIX1_NV                                                 = 0x8631
	ALPHA16F_ARB                                               = 0x881C
	DRAW_BUFFER1_ARB                                           = 0x8826
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	DSDT8_MAG8_NV                                              = 0x870A
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	PERCENTAGE_AMD                                             = 0x8BC3
	BUFFER_MAP_OFFSET                                          = 0x9121
	COLOR_LOGIC_OP                                             = 0x0BF2
	UNSIGNED_SHORT                                             = 0x1403
	BGRA                                                       = 0x80E1
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	TEXTURE_BIT                                                = 0x00040000
	TEXTURE_DEPTH_EXT                                          = 0x8071
	STATIC_ATI                                                 = 0x8760
	DRAW_BUFFER2_ARB                                           = 0x8827
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	INT_2_10_10_10_REV                                         = 0x8D9F
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	LOWER_LEFT                                                 = 0x8CA1
	DOUBLE_MAT3_EXT                                            = 0x8F47
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	N3F_V3F                                                    = 0x2A25
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	MAX_EVAL_ORDER                                             = 0x0D30
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	OBJECT_TYPE                                                = 0x9112
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	INDEX_ARRAY                                                = 0x8077
	COLOR_MATRIX_SGI                                           = 0x80B1
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	POINT_TOKEN                                                = 0x0701
	IUI_N3F_V2F_EXT                                            = 0x81AF
	CONSTANT_EXT                                               = 0x8576
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	FASTEST                                                    = 0x1101
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	DRAW_BUFFER5_ATI                                           = 0x882A
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	CLIP_DISTANCE7                                             = 0x3007
	INDEX_TEST_EXT                                             = 0x81B5
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	ACTIVE_VARIABLES                                           = 0x9305
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	RGB5_A1_OES                                                = 0x8057
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	COMBINER3_NV                                               = 0x8553
	PALETTE8_RGB8_OES                                          = 0x8B95
	STATE_RESTORE                                              = 0x8BDC
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	MINMAX_SINK                                                = 0x8030
	DEPTH_STENCIL_NV                                           = 0x84F9
	DRAW_BUFFER0_ARB                                           = 0x8825
	HALF_BIT_ATI                                               = 0x00000008
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	RED_SNORM                                                  = 0x8F90
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	EXT_vertex_array                                           = 1
	DEPTH_TEST                                                 = 0x0B71
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	FIELD_LOWER_NV                                             = 0x9023
	POINT_SIZE_MAX_EXT                                         = 0x8127
	SOURCE2_RGB_EXT                                            = 0x8582
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	VERTEX_STREAM5_ATI                                         = 0x8771
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	SHORT                                                      = 0x1402
	MINMAX_FORMAT                                              = 0x802F
	CON_27_ATI                                                 = 0x895C
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	RGB9_E5                                                    = 0x8C3D
	STENCIL_TEST                                               = 0x0B90
	VIEW_CLASS_128_BITS                                        = 0x82C4
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	BACK_RIGHT                                                 = 0x0403
	TEXTURE_MATRIX                                             = 0x0BA8
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	DRAW_BUFFER13_ARB                                          = 0x8832
	FLOAT_VEC2_ARB                                             = 0x8B50
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	SRGB8_EXT                                                  = 0x8C41
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	LOW_INT                                                    = 0x8DF3
	OBJECT_TYPE_APPLE                                          = 0x9112
	RGBA16F_ARB                                                = 0x881A
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	SWIZZLE_STR_ATI                                            = 0x8976
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	MODELVIEW0_EXT                                             = 0x1700
	TEXTURE29                                                  = 0x84DD
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	SGIX_async                                                 = 1
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	MAX_DRAW_BUFFERS                                           = 0x8824
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	DOT3_RGBA_EXT                                              = 0x8741
	TRIANGLES_ADJACENCY                                        = 0x000C
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	COLOR_ATTACHMENT8                                          = 0x8CE8
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	GREEN_BIAS                                                 = 0x0D19
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	PATH_FILL_MODE_NV                                          = 0x9080
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	DOUBLE_VEC4                                                = 0x8FFE
	BUFFER_OBJECT_EXT                                          = 0x9151
	TEXTURE_2D                                                 = 0x0DE1
	OPERAND2_ALPHA_ARB                                         = 0x859A
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	BLEND_DST_RGB                                              = 0x80C8
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	VIEW_CLASS_8_BITS                                          = 0x82CB
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	TEXTURE_BASE_LEVEL                                         = 0x813C
	TEXTURE_CUBE_MAP                                           = 0x8513
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	DRAW_BUFFER10_ARB                                          = 0x882F
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	SAMPLE_POSITION_NV                                         = 0x8E50
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	INTENSITY16                                                = 0x804D
	CONSTANT_ALPHA_EXT                                         = 0x8003
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	MATRIX31_ARB                                               = 0x88DF
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	DECAL                                                      = 0x2101
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	RGB8UI                                                     = 0x8D7D
	RGB16I                                                     = 0x8D89
	HISTOGRAM_SINK                                             = 0x802D
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	INT16_VEC4_NV                                              = 0x8FE7
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	NICEST                                                     = 0x1102
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	FLOAT_MAT2                                                 = 0x8B5A
	SGIS_point_line_texgen                                     = 1
	GL_4PASS_0_EXT                                             = 0x80A4
	CLAMP_TO_BORDER                                            = 0x812D
	MAGNITUDE_BIAS_NV                                          = 0x8718
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	RGBA_INTEGER_EXT                                           = 0x8D99
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	TEXTURE_WRAP_R_OES                                         = 0x8072
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	SAMPLE_SHADING_ARB                                         = 0x8C36
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	SAMPLE_MASK_SGIS                                           = 0x80A0
	COMPRESSED_RGB_ARB                                         = 0x84ED
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	REPLACE_EXT                                                = 0x8062
	RGBA8_OES                                                  = 0x8058
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	PACK_RESAMPLE_SGIX                                         = 0x842C
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	CONST_EYE_NV                                               = 0x86E5
	SGIX_ir_instrument1                                        = 1
	DST_COLOR                                                  = 0x0306
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	CURRENT_FOG_COORD                                          = 0x8453
	CURRENT_INDEX                                              = 0x0B01
	MAX_CLIP_PLANES                                            = 0x0D32
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	IMAGE_2D_RECT                                              = 0x904F
	MAP1_INDEX                                                 = 0x0D91
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	VERTEX_BLEND_ARB                                           = 0x86A7
	SHADER_OBJECT_ARB                                          = 0x8B48
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	RGB_INTEGER_EXT                                            = 0x8D98
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	FILL                                                       = 0x1B02
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	GL_422_REV_EXT                                             = 0x80CD
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	GL_2X_BIT_ATI                                              = 0x00000001
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	SGIS_fog_function                                          = 1
	BACK_LEFT                                                  = 0x0402
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	MATRIX14_ARB                                               = 0x88CE
	RGB32I                                                     = 0x8D83
	SAMPLE_MASK                                                = 0x8E51
	VERTICAL_LINE_TO_NV                                        = 0x08
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	SGIX_resample                                              = 1
	DISCARD_NV                                                 = 0x8530
	BUFFER_USAGE_ARB                                           = 0x8765
	MATRIX_PALETTE_OES                                         = 0x8840
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	DOUBLE_MAT4                                                = 0x8F48
	TEXTURE_WIDTH                                              = 0x1000
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	SHADER_CONSISTENT_NV                                       = 0x86DD
	ACTIVE_UNIFORMS                                            = 0x8B86
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	STENCIL_INDEX4_OES                                         = 0x8D47
	IS_PER_PATCH                                               = 0x92E7
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	DEBUG_ASSERT_MESA                                          = 0x875B
	FLOAT_R16_NV                                               = 0x8884
	INVERT                                                     = 0x150A
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	SRC2_ALPHA                                                 = 0x858A
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	INT_VEC4_ARB                                               = 0x8B55
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	CURRENT_RASTER_POSITION                                    = 0x0B07
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	HALF_FLOAT_ARB                                             = 0x140B
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	LIST_MODE                                                  = 0x0B30
	SHININESS                                                  = 0x1601
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	NORMAL_ARRAY_EXT                                           = 0x8075
	OPERAND1_ALPHA_ARB                                         = 0x8599
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	IMAGE_3D                                                   = 0x904E
	RGBA                                                       = 0x1908
	VERTEX_STREAM6_ATI                                         = 0x8772
	DRAW_BUFFER14_NV                                           = 0x8833
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	PHONG_WIN                                                  = 0x80EA
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	RGB5_A1                                                    = 0x8057
	SAMPLE_POSITION                                            = 0x8E50
	SYNC_STATUS                                                = 0x9114
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	RGBA16                                                     = 0x805B
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	TEXTURE6_ARB                                               = 0x84C6
	PRIMARY_COLOR                                              = 0x8577
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	MATRIX25_ARB                                               = 0x88D9
	SAMPLER_BINDING                                            = 0x8919
	INTERLACE_READ_OML                                         = 0x8981
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	SGI_color_matrix                                           = 1
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	DRAW_BUFFER7_ATI                                           = 0x882C
	TESS_CONTROL_SHADER                                        = 0x8E88
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	FOG_COORD_SRC                                              = 0x8450
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	OUT_OF_MEMORY                                              = 0x0505
	INVERTED_SCREEN_W_REND                                     = 0x8491
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	MODELVIEW24_ARB                                            = 0x8738
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	IS_ROW_MAJOR                                               = 0x9300
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	MIN_LOD_WARNING_AMD                                        = 0x919C
	GREEN_SCALE                                                = 0x0D18
	DIFFUSE                                                    = 0x1201
	MODELVIEW29_ARB                                            = 0x873D
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	SLUMINANCE_ALPHA                                           = 0x8C44
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	RGBA16I_EXT                                                = 0x8D88
	NAME_LENGTH                                                = 0x92F9
	VIEW_CLASS_48_BITS                                         = 0x82C7
	BINORMAL_ARRAY_EXT                                         = 0x843A
	WRITE_ONLY                                                 = 0x88B9
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	VERTEX_ARRAY_POINTER                                       = 0x808E
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	FENCE_STATUS_NV                                            = 0x84F3
	INTERLACE_READ_INGR                                        = 0x8568
	REG_17_ATI                                                 = 0x8932
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	UTF16_NV                                                   = 0x909B
)

type Context struct {
	context                   *C.gl11Context
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
	ClearDepthf               func(depth float32)
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
	Uniform1fv                func(location int32, count int32, value *float32)
	Uniform2fv                func(location int32, count int32, value *float32)
	Uniform3fv                func(location int32, count int32, value *float32)
	Uniform4fv                func(location int32, count int32, value *float32)
	Uniform1iv                func(location int32, count int32, value *int32)
	Uniform2iv                func(location int32, count int32, value *int32)
	Uniform3iv                func(location int32, count int32, value *int32)
	Uniform4iv                func(location int32, count int32, value *int32)
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
	glc.context = C.gl11NewContext()

	glc.Accum = func(op uint32, value float32) {
		C.gl11Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		C.gl11AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		C.gl11Begin(glc.context, C.GLenum(mode))
	}

	glc.End = func() {
		C.gl11End(glc.context)
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8) {
		C.gl11Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		C.gl11BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		C.gl11CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		C.gl11CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		C.gl11Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		C.gl11ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		C.gl11ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		C.gl11ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearDepthf = func(depth float32) {
		C.gl11ClearDepthf(glc.context, C.GLclampf(depth))
	}

	glc.ClearIndex = func(c float32) {
		C.gl11ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		C.gl11ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation *float64) {
		C.gl11ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.Color3b = func(red, green, blue int8) {
		C.gl11Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		C.gl11Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		C.gl11Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		C.gl11Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		C.gl11Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		C.gl11Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		C.gl11Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		C.gl11Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		C.gl11Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		C.gl11Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		C.gl11Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		C.gl11Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		C.gl11Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		C.gl11Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		C.gl11Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		C.gl11Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v *int8) {
		C.gl11Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color3dv = func(v *float64) {
		C.gl11Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color3fv = func(v *float32) {
		C.gl11Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color3iv = func(v *int32) {
		C.gl11Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color3sv = func(v *int16) {
		C.gl11Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color3ubv = func(v *uint8) {
		C.gl11Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color3uiv = func(v *uint32) {
		C.gl11Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color3usv = func(v *uint16) {
		C.gl11Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.Color4bv = func(v *int8) {
		C.gl11Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color4dv = func(v *float64) {
		C.gl11Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color4fv = func(v *float32) {
		C.gl11Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color4iv = func(v *int32) {
		C.gl11Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color4sv = func(v *int16) {
		C.gl11Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color4ubv = func(v *uint8) {
		C.gl11Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color4uiv = func(v *uint32) {
		C.gl11Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color4usv = func(v *uint16) {
		C.gl11Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		C.gl11ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		C.gl11ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		C.gl11CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		C.gl11CullFace(glc.context, C.GLenum(mode))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		C.gl11DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		C.gl11DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthMask = func(flag bool) {
		C.gl11DepthMask(glc.context, boolToGL(flag))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		C.gl11DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		C.gl11Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		C.gl11Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		C.gl11DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl11DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.EdgeFlag = func(flag bool) {
		C.gl11EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag *bool) {
		C.gl11EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(flag)))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		C.gl11EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EvalCoord1d = func(u float64) {
		C.gl11EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		C.gl11EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		C.gl11EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		C.gl11EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u *float64) {
		C.gl11EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord1fv = func(u *float32) {
		C.gl11EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2dv = func(u *float64) {
		C.gl11EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2fv = func(u *float32) {
		C.gl11EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		C.gl11EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		C.gl11EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		C.gl11EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		C.gl11EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer *float32) {
		C.gl11FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(buffer)))
	}

	glc.Finish = func() {
		C.gl11Finish(glc.context)
	}

	glc.Flush = func() {
		C.gl11Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		C.gl11Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		C.gl11Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params *float32) {
		C.gl11Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Fogiv = func(pname uint32, params *int32) {
		C.gl11Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.FrontFace = func(mode uint32) {
		C.gl11FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		C.gl11Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		return uint32(C.gl11GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params *bool) {
		C.gl11GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(params)))
	}

	glc.GetDoublev = func(pname uint32, params *float64) {
		C.gl11GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetFloatv = func(pname uint32, params *float32) {
		C.gl11GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetIntegerv = func(pname uint32, params *int32) {
		C.gl11GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetClipPlane = func(plane uint32, equation *float64) {
		C.gl11GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl11GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params *float32) {
		C.gl11GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetLightiv = func(light, pname uint32, params *int32) {
		C.gl11GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetMapdv = func(target, query uint32, v *float64) {
		C.gl11GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.GetMapfv = func(target, query uint32, v *float32) {
		C.gl11GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.GetMapiv = func(target, query uint32, v *int32) {
		C.gl11GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.GetMaterialfv = func(face, pname uint32, params *float32) {
		C.gl11GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetMaterialiv = func(face, pname uint32, params *int32) {
		C.gl11GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetPixelMapfv = func(Map uint32, values *float32) {
		C.gl11GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapuiv = func(Map uint32, values *uint32) {
		C.gl11GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapusv = func(Map uint32, values *uint16) {
		C.gl11GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.GetPolygonStipple = func(pattern *uint8) {
		C.gl11GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(pattern)))
	}

	glc.GetString = func(name uint32) string {
		cstr := C.gl11GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params *float32) {
		C.gl11GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexEnviv = func(target, pname uint32, params *int32) {
		C.gl11GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexGendv = func(coord, pname uint32, params *float64) {
		C.gl11GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params *float32) {
		C.gl11GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params *int32) {
		C.gl11GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params *float32) {
		C.gl11GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params *int32) {
		C.gl11GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params *float32) {
		C.gl11GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params *int32) {
		C.gl11GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Hint = func(target, mode uint32) {
		C.gl11Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		C.gl11Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		C.gl11Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		C.gl11Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		C.gl11Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexdv = func(c *float64) {
		C.gl11Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(c)))
	}

	glc.Indexfv = func(c *float32) {
		C.gl11Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(c)))
	}

	glc.Indexiv = func(c *int32) {
		C.gl11Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(c)))
	}

	glc.Indexsv = func(c *int16) {
		C.gl11Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(c)))
	}

	glc.IndexMask = func(mask uint32) {
		C.gl11IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		C.gl11InitNames(glc.context)
	}

	glc.IsEnabled = func(cap uint32) {
		C.gl11IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		return C.gl11IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		C.gl11Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		C.gl11Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params *float32) {
		C.gl11Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Lightiv = func(light, pname uint32, params *int32) {
		C.gl11Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		C.gl11LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		C.gl11LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params *float32) {
		C.gl11LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.LightModeliv = func(pname uint32, params *int32) {
		C.gl11LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		C.gl11LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		C.gl11LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		C.gl11ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		C.gl11LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m *float64) {
		C.gl11LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadMatrixf = func(m *float32) {
		C.gl11LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.LoadName = func(name uint32) {
		C.gl11LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		C.gl11LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points *float64) {
		C.gl11Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points *float32) {
		C.gl11Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64) {
		C.gl11Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32) {
		C.gl11Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		C.gl11MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		C.gl11MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		C.gl11MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		C.gl11MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		C.gl11Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		C.gl11Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params *float32) {
		C.gl11Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Materialiv = func(face, pname uint32, params *int32) {
		C.gl11Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.MatrixMode = func(mode uint32) {
		C.gl11MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m *float64) {
		C.gl11MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultMatrixf = func(m *float32) {
		C.gl11MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.NewList = func(list uint32, mode uint32) {
		C.gl11NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		C.gl11EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		C.gl11Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		C.gl11Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		C.gl11Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		C.gl11Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		C.gl11Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v *int8) {
		C.gl11Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Normal3dv = func(v *float64) {
		C.gl11Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Normal3fv = func(v *float32) {
		C.gl11Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Normal3iv = func(v *int32) {
		C.gl11Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Normal3sv = func(v *int16) {
		C.gl11Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		C.gl11Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		C.gl11PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values *float32) {
		C.gl11PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values *uint32) {
		C.gl11PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values *uint16) {
		C.gl11PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		C.gl11PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		C.gl11PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		C.gl11PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		C.gl11PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		C.gl11PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		C.gl11PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		C.gl11PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask *uint8) {
		C.gl11PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(mask)))
	}

	glc.PushAttrib = func(mask uint32) {
		C.gl11PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		C.gl11PopAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		C.gl11PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		C.gl11PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		C.gl11PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		C.gl11PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		C.gl11RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		C.gl11RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		C.gl11RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		C.gl11RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		C.gl11RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		C.gl11RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		C.gl11RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		C.gl11RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		C.gl11RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		C.gl11RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		C.gl11RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		C.gl11RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v *float64) {
		C.gl11RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos2fv = func(v *float32) {
		C.gl11RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos2iv = func(v *int32) {
		C.gl11RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos2sv = func(v *int16) {
		C.gl11RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos3dv = func(v *float64) {
		C.gl11RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos3fv = func(v *float32) {
		C.gl11RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos3iv = func(v *int32) {
		C.gl11RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos3sv = func(v *int16) {
		C.gl11RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos4dv = func(v *float64) {
		C.gl11RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos4fv = func(v *float32) {
		C.gl11RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos4iv = func(v *int32) {
		C.gl11RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos4sv = func(v *int16) {
		C.gl11RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.ReadBuffer = func(mode uint32) {
		C.gl11ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		C.gl11Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		C.gl11Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		C.gl11Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		C.gl11Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 *float64) {
		C.gl11Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
	}

	glc.Rectfv = func(v1, v2 *float32) {
		C.gl11Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
	}

	glc.Rectiv = func(v1, v2 *int32) {
		C.gl11Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
	}

	glc.Rectsv = func(v1, v2 *int16) {
		C.gl11Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
	}

	glc.RenderMode = func(mode uint32) int32 {
		return int32(C.gl11RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		C.gl11Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		C.gl11Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		C.gl11Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		C.gl11Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		C.gl11Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer *uint32) {
		C.gl11SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
	}

	glc.ShadeModel = func(mode uint32) {
		C.gl11ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		C.gl11StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		C.gl11StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		C.gl11StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		C.gl11TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		C.gl11TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		C.gl11TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		C.gl11TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		C.gl11TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		C.gl11TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		C.gl11TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		C.gl11TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		C.gl11TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		C.gl11TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		C.gl11TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		C.gl11TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		C.gl11TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		C.gl11TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		C.gl11TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		C.gl11TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v *float64) {
		C.gl11TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord1fv = func(v *float32) {
		C.gl11TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord1iv = func(v *int32) {
		C.gl11TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord1sv = func(v *int16) {
		C.gl11TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord2dv = func(v *float64) {
		C.gl11TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord2fv = func(v *float32) {
		C.gl11TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord2iv = func(v *int32) {
		C.gl11TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord2sv = func(v *int16) {
		C.gl11TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord3dv = func(v *float64) {
		C.gl11TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord3fv = func(v *float32) {
		C.gl11TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord3iv = func(v *int32) {
		C.gl11TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord3sv = func(v *int16) {
		C.gl11TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord4dv = func(v *float64) {
		C.gl11TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord4fv = func(v *float32) {
		C.gl11TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord4iv = func(v *int32) {
		C.gl11TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord4sv = func(v *int16) {
		C.gl11TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		C.gl11TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		C.gl11TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params *float32) {
		C.gl11TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexEnviv = func(target, pname uint32, params *int32) {
		C.gl11TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		C.gl11TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		C.gl11TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		C.gl11TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params *float64) {
		C.gl11TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.TexGenfv = func(coord, pname uint32, params *float32) {
		C.gl11TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexGeniv = func(coord, pname uint32, params *int32) {
		C.gl11TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		C.gl11TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		C.gl11TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params *float32) {
		C.gl11TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexParameteriv = func(target, pname uint32, params *int32) {
		C.gl11TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Translated = func(x, y, z float64) {
		C.gl11Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		C.gl11Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		C.gl11Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		C.gl11Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		C.gl11Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		C.gl11Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		C.gl11Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		C.gl11Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		C.gl11Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		C.gl11Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		C.gl11Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		C.gl11Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		C.gl11Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		C.gl11Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		C.gl11Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetConvolutionParameterfv = func(target, pname uint32, params *float32) {
		C.gl11GetConvolutionParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionParameteriv = func(target, pname uint32, params *int32) {
		C.gl11GetConvolutionParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		var cRes *C.GLboolean
		status = C.gl11AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		C.gl11ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		C.gl11DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl11DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		C.gl11GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		C.gl11PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		C.gl11CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		C.gl11CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		C.gl11CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		C.gl11CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		C.gl11BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures *uint32) {
		C.gl11DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.GenTextures = func(n int32, textures *uint32) {
		C.gl11GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.IsTexture = func(texture uint32) bool {
		return C.gl11IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		C.gl11EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		C.gl11DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.Indexub = func(c uint8) {
		C.gl11Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexubv = func(c *uint8) {
		C.gl11Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(c)))
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.PushClientAttrib = func(mask uint32) {
		C.gl11PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PrioritizeTextures = func(n int32, textures *uint32, priorities *float32) {
		C.gl11PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLclampf)(unsafe.Pointer(priorities)))
	}

	glc.PopClientAttrib = func() {
		C.gl11PopClientAttrib(glc.context)
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl11ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params *float32) {
		C.gl11ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params *int32) {
		C.gl11ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		C.gl11ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl11ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl11ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		C.gl11ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		C.gl11ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl11CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		C.gl11CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl11CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		C.gl11CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetColorTable = func(target, format, Type uint32, table unsafe.Pointer) {
		C.gl11GetColorTable(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), table)
	}

	glc.GetColorTableParameterfv = func(target, pname uint32, params *float32) {
		C.gl11GetColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetColorTableParameteriv = func(target, pname uint32, params *int32) {
		C.gl11GetColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionFilter = func(target, format, Type uint32, image unsafe.Pointer) {
		C.gl11GetConvolutionFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), image)
	}

	glc.GetHistogram = func(target uint32, reset bool, format, Type uint32, values unsafe.Pointer) {
		C.gl11GetHistogram(glc.context, C.GLenum(target), boolToGL(reset), C.GLenum(format), C.GLenum(Type), values)
	}

	glc.GetHistogramParameterfv = func(target, pname uint32, params *float32) {
		C.gl11GetHistogramParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetHistogramParameteriv = func(target, pname uint32, params *int32) {
		C.gl11GetHistogramParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetSeparableFilter = func(target, format, Type uint32, row, column, span unsafe.Pointer) {
		C.gl11GetSeparableFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), row, column, span)
	}

	glc.Histogram = func(target uint32, width int32, internalformat uint32, sink bool) {
		C.gl11Histogram(glc.context, C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.Minmax = func(target, internalformat uint32, sink bool) {
		C.gl11Minmax(glc.context, C.GLenum(target), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.MultiTexCoord1s = func(target uint32, s int16) {
		C.gl11MultiTexCoord1s(glc.context, C.GLenum(target), C.GLshort(s))
	}

	glc.MultiTexCoord1i = func(target uint32, s int32) {
		C.gl11MultiTexCoord1i(glc.context, C.GLenum(target), C.GLint(s))
	}

	glc.MultiTexCoord1f = func(target uint32, s float32) {
		C.gl11MultiTexCoord1f(glc.context, C.GLenum(target), C.GLfloat(s))
	}

	glc.MultiTexCoord1d = func(target uint32, s float64) {
		C.gl11MultiTexCoord1d(glc.context, C.GLenum(target), C.GLdouble(s))
	}

	glc.MultiTexCoord2s = func(target uint32, s, t int16) {
		C.gl11MultiTexCoord2s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t))
	}

	glc.MultiTexCoord2i = func(target uint32, s, t int32) {
		C.gl11MultiTexCoord2i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t))
	}

	glc.MultiTexCoord2f = func(target uint32, s, t float32) {
		C.gl11MultiTexCoord2f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
	}

	glc.MultiTexCoord2d = func(target uint32, s, t float64) {
		C.gl11MultiTexCoord2d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
	}

	glc.MultiTexCoord3s = func(target uint32, s, t, r int16) {
		C.gl11MultiTexCoord3s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.MultiTexCoord3i = func(target uint32, s, t, r int32) {
		C.gl11MultiTexCoord3i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.MultiTexCoord3f = func(target uint32, s, t, r float32) {
		C.gl11MultiTexCoord3f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.MultiTexCoord3d = func(target uint32, s, t, r float64) {
		C.gl11MultiTexCoord3d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.MultiTexCoord4s = func(target uint32, s, t, r, q int16) {
		C.gl11MultiTexCoord4s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.MultiTexCoord4i = func(target uint32, s, t, r, q int32) {
		C.gl11MultiTexCoord4i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.MultiTexCoord4f = func(target uint32, s, t, r, q float32) {
		C.gl11MultiTexCoord4f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.MultiTexCoord4d = func(target uint32, s, t, r, q float64) {
		C.gl11MultiTexCoord4d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.MultiTexCoord1sv = func(target uint32, v *int16) {
		C.gl11MultiTexCoord1sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1iv = func(target uint32, v *int32) {
		C.gl11MultiTexCoord1iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1fv = func(target uint32, v *float32) {
		C.gl11MultiTexCoord1fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1dv = func(target uint32, v *float64) {
		C.gl11MultiTexCoord1dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2sv = func(target uint32, v *int16) {
		C.gl11MultiTexCoord2sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2iv = func(target uint32, v *int32) {
		C.gl11MultiTexCoord2iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2fv = func(target uint32, v *float32) {
		C.gl11MultiTexCoord2fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2dv = func(target uint32, v *float64) {
		C.gl11MultiTexCoord2dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3sv = func(target uint32, v *int16) {
		C.gl11MultiTexCoord3sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3iv = func(target uint32, v *int32) {
		C.gl11MultiTexCoord3iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3fv = func(target uint32, v *float32) {
		C.gl11MultiTexCoord3fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3dv = func(target uint32, v *float64) {
		C.gl11MultiTexCoord3dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4sv = func(target uint32, v *int16) {
		C.gl11MultiTexCoord4sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4iv = func(target uint32, v *int32) {
		C.gl11MultiTexCoord4iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4fv = func(target uint32, v *float32) {
		C.gl11MultiTexCoord4fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4dv = func(target uint32, v *float64) {
		C.gl11MultiTexCoord4dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.ResetHistogram = func(target uint32) {
		C.gl11ResetHistogram(glc.context, C.GLenum(target))
	}

	glc.ResetMinmax = func(target uint32) {
		C.gl11ResetMinmax(glc.context, C.GLenum(target))
	}

	glc.SeparableFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, row, column unsafe.Pointer) {
		C.gl11SeparableFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), row, column)
	}

	glc.BlendColor = func(red, green, blue, alpha float32) {
		C.gl11BlendColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		C.gl11BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		C.gl11CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DrawRangeElements = func(mode uint32, start, end uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl11DrawRangeElements(glc.context, C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.TexImage3D = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11TexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl11TexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.ActiveTexture = func(texture uint32) {
		C.gl11ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		C.gl11ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl11CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl11CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl11CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl11CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl11CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl11CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.GetCompressedTexImage = func(target uint32, lod int32, img unsafe.Pointer) {
		C.gl11GetCompressedTexImage(glc.context, C.GLenum(target), C.GLint(lod), img)
	}

	glc.LoadTransposeMatrixd = func(m *float64) {
		C.gl11LoadTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadTransposeMatrixf = func(m *float64) {
		C.gl11LoadTransposeMatrixf(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixd = func(m *float64) {
		C.gl11MultTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixf = func(m *float32) {
		C.gl11MultTransposeMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.SampleCoverage = func(value float32, invert bool) {
		C.gl11SampleCoverage(glc.context, C.GLclampf(value), boolToGL(invert))
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		C.gl11BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.FogCoordPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11FogCoordPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.FogCoordd = func(coord float64) {
		C.gl11FogCoordd(glc.context, C.GLdouble(coord))
	}

	glc.FogCoordf = func(coord float32) {
		C.gl11FogCoordf(glc.context, C.GLfloat(coord))
	}

	glc.FogCoorddv = func(coord *float64) {
		C.gl11FogCoorddv(glc.context, (*C.GLdouble)(unsafe.Pointer(coord)))
	}

	glc.FogCoordfv = func(coord *float32) {
		C.gl11FogCoordfv(glc.context, (*C.GLfloat)(unsafe.Pointer(coord)))
	}

	glc.MultiDrawArrays = func(mode uint32, first *int32, count *int32, primcount int32) {
		C.gl11MultiDrawArrays(glc.context, C.GLenum(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), C.GLsizei(primcount))
	}

	glc.MultiDrawElements = func(mode uint32, count *int32, Type uint32, indices unsafe.Pointer, primcount int32) {
		C.gl11MultiDrawElements(glc.context, C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(count)), C.GLenum(Type), indices, C.GLsizei(primcount))
	}

	glc.PointParameterf = func(pname uint32, param float32) {
		C.gl11PointParameterf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PointParameteri = func(pname uint32, param int32) {
		C.gl11PointParameteri(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.SecondaryColor3b = func(red, green, blue int8) {
		C.gl11SecondaryColor3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.SecondaryColor3s = func(red, green, blue int16) {
		C.gl11SecondaryColor3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.SecondaryColor3i = func(red, green, blue int32) {
		C.gl11SecondaryColor3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.SecondaryColor3f = func(red, green, blue float32) {
		C.gl11SecondaryColor3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.SecondaryColor3d = func(red, green, blue float64) {
		C.gl11SecondaryColor3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.SecondaryColor3ub = func(red, green, blue uint8) {
		C.gl11SecondaryColor3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.SecondaryColor3us = func(red, green, blue uint16) {
		C.gl11SecondaryColor3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.SecondaryColor3ui = func(red, green, blue uint32) {
		C.gl11SecondaryColor3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.SecondaryColor3bv = func(v *int8) {
		C.gl11SecondaryColor3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3sv = func(v *int16) {
		C.gl11SecondaryColor3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3iv = func(v *int32) {
		C.gl11SecondaryColor3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3fv = func(v *float32) {
		C.gl11SecondaryColor3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3dv = func(v *float64) {
		C.gl11SecondaryColor3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3ubv = func(v *uint8) {
		C.gl11SecondaryColor3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3usv = func(v *uint16) {
		C.gl11SecondaryColor3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3uiv = func(v *uint32) {
		C.gl11SecondaryColor3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl11SecondaryColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.WindowPos2s = func(x, y int16) {
		C.gl11WindowPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.WindowPos2i = func(x, y int32) {
		C.gl11WindowPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.WindowPos2f = func(x, y float32) {
		C.gl11WindowPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.WindowPos2d = func(x, y float64) {
		C.gl11WindowPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.WindowPos3s = func(x, y, z int16) {
		C.gl11WindowPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.WindowPos3i = func(x, y, z int32) {
		C.gl11WindowPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.WindowPos3f = func(x, y, z float32) {
		C.gl11WindowPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.WindowPos3d = func(x, y, z float64) {
		C.gl11WindowPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.WindowPos2sv = func(v *int16) {
		C.gl11WindowPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos2iv = func(v *int32) {
		C.gl11WindowPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos2fv = func(v *float32) {
		C.gl11WindowPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos2dv = func(v *float64) {
		C.gl11WindowPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.WindowPos3sv = func(v *int16) {
		C.gl11WindowPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos3iv = func(v *int32) {
		C.gl11WindowPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos3fv = func(v *float32) {
		C.gl11WindowPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos3dv = func(v *float64) {
		C.gl11WindowPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		C.gl11BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		C.gl11BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		C.gl11BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		C.gl11BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.DeleteBuffers = func(n int32, buffers *uint32) {
		C.gl11DeleteBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.DeleteQueries = func(n int32, ids *uint32) {
		C.gl11DeleteQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GenBuffers = func(n int32, buffers *uint32) {
		C.gl11GenBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.GenQueries = func(n int32, ids *uint32) {
		C.gl11GenQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GetBufferParameteriv = func(target, value uint32, data *int32) {
		C.gl11GetBufferParameteriv(glc.context, C.GLenum(target), C.GLenum(value), (*C.GLint)(unsafe.Pointer(data)))
	}

	glc.GetBufferPointerv = func(target, pname uint32, params unsafe.Pointer) {
		C.gl11GetBufferPointerv(glc.context, C.GLenum(target), C.GLenum(pname), params)
	}

	glc.GetBufferSubData = func(target uint32, offset int32, size int32, data unsafe.Pointer) {
		C.gl11GetBufferSubData(glc.context, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}

	glc.GetQueryObjectiv = func(id uint32, pname uint32, params *int32) {
		C.gl11GetQueryObjectiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetQueryObjectuiv = func(id uint32, pname uint32, params *uint32) {
		C.gl11GetQueryObjectuiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(params)))
	}

	glc.GetQueryiv = func(target, pname uint32, params *int32) {
		C.gl11GetQueryiv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.IsBuffer = func(buffer uint32) bool {
		return C.gl11IsBuffer(glc.context, C.GLuint(buffer)) != 0
	}

	glc.IsQuery = func(id uint32) bool {
		return C.gl11IsQuery(glc.context, C.GLuint(id)) != 0
	}

	glc.MapBuffer = func(target, access uint32) unsafe.Pointer {
		return unsafe.Pointer(C.gl11MapBuffer(glc.context, C.GLenum(target), C.GLenum(access)))
	}

	glc.UnmapBuffer = func(target uint32) bool {
		return C.gl11UnmapBuffer(glc.context, C.GLenum(target)) != 0
	}

	glc.AttachShader = func(program, shader uint32) {
		C.gl11AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl11BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha uint32) {
		C.gl11BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		C.gl11CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		return uint32(C.gl11CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType uint32) uint32 {
		return uint32(C.gl11CreateShader(glc.context, C.GLenum(shaderType)))
	}

	glc.DeleteProgram = func(program uint32) {
		C.gl11DeleteProgram(glc.context, C.GLuint(program))
	}

	glc.DeleteShader = func(shader uint32) {
		C.gl11DeleteShader(glc.context, C.GLuint(shader))
	}

	glc.DetachShader = func(program, shader uint32) {
		C.gl11DetachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.EnableVertexAttribArray = func(index uint32) {
		C.gl11EnableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DisableVertexAttribArray = func(index uint32) {
		C.gl11DisableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DrawBuffers = func(n int32, bufs *uint32) {
		C.gl11DrawBuffers(glc.context, C.GLsizei(n), (*C.GLenum)(unsafe.Pointer(bufs)))
	}

	glc.GetActiveAttrib = func(program, index uint32, bufSize int32) (length int32, size int32, Type uint32, name string) {
		var (
			cname C.GLchar
		)
		C.gl11GetActiveAttrib(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
		name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
		return
	}

	glc.GetActiveUniform = func(program, index uint32, bufSize int32, length *int32, size *int32, Type *uint32, name *byte) {
		C.gl11GetActiveUniform(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(Type)), (*C.GLchar)(unsafe.Pointer(name)))
	}

	glc.GetAttachedShaders = func(program uint32, maxCount int32, count *int32, shaders *uint32) {
		C.gl11GetAttachedShaders(glc.context, C.GLuint(program), C.GLsizei(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
	}

	glc.GetAttribLocation = func(program uint32, name *byte) int32 {
		return int32(C.gl11GetAttribLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetProgramiv = func(program uint32, pname uint32, params *int32) {
		C.gl11GetProgramiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetProgramInfoLog = func(program uint32, maxLength int32, length *int32, infoLog *byte) {
		C.gl11GetProgramInfoLog(glc.context, C.GLuint(program), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderiv = func(program uint32, pname uint32, params *int32) {
		C.gl11GetShaderiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetShaderInfoLog = func(shader uint32, maxLength int32, length *int32, infoLog *byte) {
		C.gl11GetShaderInfoLog(glc.context, C.GLuint(shader), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderSource = func(shader uint32, bufSize int32, length *int32, source *byte) {
		C.gl11GetShaderSource(glc.context, C.GLuint(shader), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
	}

	glc.GetUniformfv = func(program uint32, location int32, params *float32) {
		C.gl11GetUniformfv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetUniformiv = func(program uint32, location int32, params *int32) {
		C.gl11GetUniformiv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetUniformLocation = func(program uint32, name *byte) int32 {
		return int32(C.gl11GetUniformLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetVertexAttribdv = func(index uint32, pname uint32, params *float64) {
		C.gl11GetVertexAttribdv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribfv = func(index uint32, pname uint32, params *float32) {
		C.gl11GetVertexAttribfv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribiv = func(index uint32, pname uint32, params *int32) {
		C.gl11GetVertexAttribiv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribPointerv = func(index uint32, pname uint32, pointer unsafe.Pointer) {
		C.gl11GetVertexAttribPointerv(glc.context, C.GLuint(index), C.GLenum(pname), pointer)
	}

	glc.IsProgram = func(program uint32) bool {
		return C.gl11IsProgram(glc.context, C.GLuint(program)) != 0
	}

	glc.IsShader = func(shader uint32) bool {
		return C.gl11IsShader(glc.context, C.GLuint(shader)) != 0
	}

	glc.LinkProgram = func(program uint32) {
		C.gl11LinkProgram(glc.context, C.GLuint(program))
	}

	glc.ShaderSource = func(shader uint32, count int32, string **byte, length *int32) {
		C.gl11ShaderSource(glc.context, C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(string)), (*C.GLint)(unsafe.Pointer(length)))
	}

	glc.StencilFuncSeparate = func(face, Func uint32, ref int32, mask uint32) {
		C.gl11StencilFuncSeparate(glc.context, C.GLenum(face), C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMaskSeparate = func(face uint32, mask uint32) {
		C.gl11StencilMaskSeparate(glc.context, C.GLenum(face), C.GLuint(mask))
	}

	glc.StencilOpSeparate = func(face, sfail, dpfail, dppass uint32) {
		C.gl11StencilOpSeparate(glc.context, C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
	}

	glc.Uniform1f = func(location int32, v0 float32) {
		C.gl11Uniform1f(glc.context, C.GLint(location), C.GLfloat(v0))
	}

	glc.Uniform2f = func(location int32, v0, v1 float32) {
		C.gl11Uniform2f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.Uniform3f = func(location int32, v0, v1, v2 float32) {
		C.gl11Uniform3f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Uniform4f = func(location int32, v0, v1, v2, v3 float32) {
		C.gl11Uniform4f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.Uniform1i = func(location, v0 int32) {
		C.gl11Uniform1i(glc.context, C.GLint(location), C.GLint(v0))
	}

	glc.Uniform2i = func(location, v0, v1 int32) {
		C.gl11Uniform2i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1))
	}

	glc.Uniform3i = func(location, v0, v1, v2 int32) {
		C.gl11Uniform3i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
	}

	glc.Uniform4i = func(location, v0, v1, v2, v3 int32) {
		C.gl11Uniform4i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
	}

	glc.Uniform1fv = func(location int32, count int32, value *float32) {
		C.gl11Uniform1fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform2fv = func(location int32, count int32, value *float32) {
		C.gl11Uniform2fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform3fv = func(location int32, count int32, value *float32) {
		C.gl11Uniform3fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform4fv = func(location int32, count int32, value *float32) {
		C.gl11Uniform4fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform1iv = func(location int32, count int32, value *int32) {
		C.gl11Uniform1iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform2iv = func(location int32, count int32, value *int32) {
		C.gl11Uniform2iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform3iv = func(location int32, count int32, value *int32) {
		C.gl11Uniform3iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform4iv = func(location int32, count int32, value *int32) {
		C.gl11Uniform4iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.UseProgram = func(program uint32) {
		C.gl11UseProgram(glc.context, C.GLuint(program))
	}

	glc.ValidateProgram = func(program uint32) {
		C.gl11ValidateProgram(glc.context, C.GLuint(program))
	}

	glc.VertexAttribPointer = func(index uint32, size int32, Type uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
		C.gl11VertexAttribPointer(glc.context, C.GLuint(index), C.GLint(size), C.GLenum(Type), boolToGL(normalized), C.GLsizei(stride), pointer)
	}

	glc.VertexAttrib1f = func(index uint32, v0 float32) {
		C.gl11VertexAttrib1f(glc.context, C.GLuint(index), C.GLfloat(v0))
	}

	glc.VertexAttrib1s = func(index uint32, v0 int16) {
		C.gl11VertexAttrib1s(glc.context, C.GLuint(index), C.GLshort(v0))
	}

	glc.VertexAttrib1d = func(index uint32, v0 float64) {
		C.gl11VertexAttrib1d(glc.context, C.GLuint(index), C.GLdouble(v0))
	}

	glc.VertexAttrib2f = func(index uint32, v0, v1 float32) {
		C.gl11VertexAttrib2f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.VertexAttrib2s = func(index uint32, v0, v1 int16) {
		C.gl11VertexAttrib2s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
	}

	glc.VertexAttrib2d = func(index uint32, v0, v1 float64) {
		C.gl11VertexAttrib2d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
	}

	glc.VertexAttrib3f = func(index uint32, v0, v1, v2 float32) {
		C.gl11VertexAttrib3f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.VertexAttrib3s = func(index uint32, v0, v1, v2 int16) {
		C.gl11VertexAttrib3s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
	}

	glc.VertexAttrib3d = func(index uint32, v0, v1, v2 float64) {
		C.gl11VertexAttrib3d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.VertexAttrib4f = func(index uint32, v0, v1, v2, v3 float32) {
		C.gl11VertexAttrib4f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.VertexAttrib4s = func(index uint32, v0, v1, v2, v3 int16) {
		C.gl11VertexAttrib4s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
	}

	glc.VertexAttrib4d = func(index uint32, v0, v1, v2, v3 float64) {
		C.gl11VertexAttrib4d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
	}

	glc.VertexAttrib4Nuv = func(index uint32, v0, v1, v2, v3 uint8) {
		C.gl11VertexAttrib4Nuv(glc.context, C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
	}

	glc.VertexAttrib1fv = func(index uint32, v *float32) {
		C.gl11VertexAttrib1fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1sv = func(index uint32, v *int16) {
		C.gl11VertexAttrib1sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1dv = func(index uint32, v *float64) {
		C.gl11VertexAttrib1dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2fv = func(index uint32, v *float32) {
		C.gl11VertexAttrib2fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2sv = func(index uint32, v *int16) {
		C.gl11VertexAttrib2sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2dv = func(index uint32, v *float64) {
		C.gl11VertexAttrib2dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3fv = func(index uint32, v *float32) {
		C.gl11VertexAttrib3fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3sv = func(index uint32, v *int16) {
		C.gl11VertexAttrib3sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3dv = func(index uint32, v *float64) {
		C.gl11VertexAttrib3dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4fv = func(index uint32, v *float32) {
		C.gl11VertexAttrib4fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4sv = func(index uint32, v *int16) {
		C.gl11VertexAttrib4sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4dv = func(index uint32, v *float64) {
		C.gl11VertexAttrib4dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4iv = func(index uint32, v *int32) {
		C.gl11VertexAttrib4iv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4bv = func(index uint32, v *int8) {
		C.gl11VertexAttrib4bv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4ubv = func(index uint32, v *uint8) {
		C.gl11VertexAttrib4ubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4usv = func(index uint32, v *uint16) {
		C.gl11VertexAttrib4usv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4uiv = func(index uint32, v *uint32) {
		C.gl11VertexAttrib4uiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nbv = func(index uint32, v *int8) {
		C.gl11VertexAttrib4Nbv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nsv = func(index uint32, v *int16) {
		C.gl11VertexAttrib4Nsv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Niv = func(index uint32, v *int32) {
		C.gl11VertexAttrib4Niv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nubv = func(index uint32, v *uint8) {
		C.gl11VertexAttrib4Nubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nusv = func(index uint32, v *uint16) {
		C.gl11VertexAttrib4Nusv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nuiv = func(index uint32, v *uint32) {
		C.gl11VertexAttrib4Nuiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.UniformMatrix2fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x3fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix2x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x2fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix3x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x4fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix2x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x2fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix4x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x4fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix3x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x3fv = func(location int32, count int32, transpose bool, value *float32) {
		C.gl11UniformMatrix4x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	if !versionSupported(glc) {
		return nil
	}
	return glc
}
