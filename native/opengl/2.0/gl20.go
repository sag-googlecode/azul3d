// Package 'opengl' implements OpenGL version 2.0
package opengl

// #cgo LDFLAGS: -lopengl32
// #include "gl20.h"
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
		n, wantedMajor, wantedMinor, wantedRev := parseVersions("2.0")
		if n < 2 {
			fmt.Printf("OpenGL: *** JSON version parsing failed for %q ***\n", "2.0")
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
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	ALPHA16_EXT                                                = 0x803E
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	TEXTURE29                                                  = 0x84DD
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	SLUMINANCE_ALPHA                                           = 0x8C44
	COMBINE_ALPHA_ARB                                          = 0x8572
	DRAW_BUFFER5_NV                                            = 0x882A
	UNIFORM_SIZE                                               = 0x8A38
	LUMINANCE16_SNORM                                          = 0x9019
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	REG_6_ATI                                                  = 0x8927
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	SGIX_fog_offset                                            = 1
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	VERTEX_STREAM0_ATI                                         = 0x876C
	DRAW_BUFFER15                                              = 0x8834
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	PACK_ROW_LENGTH                                            = 0x0D02
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	MATRIX19_ARB                                               = 0x88D3
	MAX_SUBROUTINES                                            = 0x8DE7
	TESS_GEN_POINT_MODE                                        = 0x8E79
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	HISTOGRAM_WIDTH                                            = 0x8026
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	INDEX_SHIFT                                                = 0x0D12
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	TEXTURE10_ARB                                              = 0x84CA
	DEPTH_STENCIL_EXT                                          = 0x84F9
	SIGNED_ALPHA_NV                                            = 0x8705
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	MAP_READ_BIT                                               = 0x0001
	TEXTURE_BLUE_SIZE                                          = 0x805E
	CLIP_PLANE5                                                = 0x3005
	COLOR_TABLE_SGI                                            = 0x80D0
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	VARIABLE_B_NV                                              = 0x8524
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	CUBIC_CURVE_TO_NV                                          = 0x0C
	DRAW_BUFFER4                                               = 0x8829
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	COLOR_INDEX2_EXT                                           = 0x80E3
	DRAW_BUFFER3_ATI                                           = 0x8828
	DRAW_BUFFER13_ATI                                          = 0x8832
	BOOL_VEC3                                                  = 0x8B58
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	COLOR_ATTACHMENT8                                          = 0x8CE8
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	V2F                                                        = 0x2A20
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	INTENSITY16F_ARB                                           = 0x881D
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	WRITE_ONLY_ARB                                             = 0x88B9
	DATA_BUFFER_AMD                                            = 0x9151
	EMISSION                                                   = 0x1600
	SIGNED_HILO16_NV                                           = 0x86FA
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	ALPHA4_EXT                                                 = 0x803B
	GL_422_EXT                                                 = 0x80CC
	TEXTURE11_ARB                                              = 0x84CB
	SIGNED_INTENSITY8_NV                                       = 0x8708
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	COLOR_TABLE_BIAS                                           = 0x80D7
	MATRIX26_ARB                                               = 0x88DA
	SAMPLE_BUFFERS                                             = 0x80A8
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	STENCIL_INDEX8                                             = 0x8D48
	RGBA16UI                                                   = 0x8D76
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	RGBA8                                                      = 0x8058
	RESCALE_NORMAL_EXT                                         = 0x803A
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	OCCLUSION_TEST_HP                                          = 0x8165
	INVERSE_NV                                                 = 0x862B
	PROXY_TEXTURE_1D                                           = 0x8063
	LUMINANCE12_ALPHA4                                         = 0x8046
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	OP_SUB_EXT                                                 = 0x8796
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	TESS_GEN_SPACING                                           = 0x8E77
	DRAW_PIXEL_TOKEN                                           = 0x0705
	UNSIGNED_BYTE                                              = 0x1401
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	MIRROR_CLAMP_EXT                                           = 0x8742
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	FALSE                                                      = 0
	IGNORE_BORDER_HP                                           = 0x8150
	SKIP_COMPONENTS2_NV                                        = -5
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	BLEND_DST                                                  = 0x0BE0
	RG32UI                                                     = 0x823C
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	RGB_SNORM                                                  = 0x8F92
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	READ_ONLY                                                  = 0x88B8
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	MATRIX12_ARB                                               = 0x88CC
	RELATIVE_ARC_TO_NV                                         = 0xFF
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	DRAW_BUFFER5_ATI                                           = 0x882A
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	EIGHTH_BIT_ATI                                             = 0x00000020
	EXT_abgr                                                   = 1
	DECAL                                                      = 0x2101
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	FLOAT_R32_NV                                               = 0x8885
	ROUND_NV                                                   = 0x90A4
	BYTE                                                       = 0x1400
	TEXTURE_MAX_LEVEL                                          = 0x813D
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	INTENSITY                                                  = 0x8049
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	SGIX_list_priority                                         = 1
	AUX1                                                       = 0x040A
	TEXTURE_WRAP_S                                             = 0x2802
	TEXTURE6_ARB                                               = 0x84C6
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	POLYGON_OFFSET_POINT                                       = 0x2A01
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	VERTEX_PROGRAM_ARB                                         = 0x8620
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	COMBINE                                                    = 0x8570
	REG_17_ATI                                                 = 0x8932
	TEXTURE_COORD_NV                                           = 0x8C79
	SGIX_vertex_preclip                                        = 1
	SRGB_READ                                                  = 0x8297
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	LO_BIAS_NV                                                 = 0x8715
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	TIMESTAMP                                                  = 0x8E28
	DOUBLE_MAT4x3                                              = 0x8F4E
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	VERSION_3_1                                                = 1
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	FLOAT_MAT3                                                 = 0x8B5B
	RGB32UI_EXT                                                = 0x8D71
	RGB8I                                                      = 0x8D8F
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	SRC_ALPHA                                                  = 0x0302
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	RGBA_FLOAT16_ATI                                           = 0x881A
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	SAMPLER                                                    = 0x82E6
	SLUMINANCE_NV                                              = 0x8C46
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	TEXTURE16                                                  = 0x84D0
	MATRIX_PALETTE_ARB                                         = 0x8840
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	PACK_RESAMPLE_SGIX                                         = 0x842C
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	MAP1_COLOR_4                                               = 0x0D90
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	REFERENCE_PLANE_SGIX                                       = 0x817D
	TEXTURE12_ARB                                              = 0x84CC
	COMPRESSED_LUMINANCE                                       = 0x84EA
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	SYNC_STATUS                                                = 0x9114
	GL_4_BYTES                                                 = 0x1409
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	ETC1_RGB8_OES                                              = 0x8D64
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	PALETTE4_RGBA8_OES                                         = 0x8B91
	CLIP_DISTANCE_NV                                           = 0x8C7A
	LUMINANCE8I_EXT                                            = 0x8D92
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	VARIANT_EXT                                                = 0x87C1
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	TRANSLATE_X_NV                                             = 0x908E
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	MAX_DEPTH                                                  = 0x8280
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	MAX                                                        = 0x8008
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	DRAW_BUFFER6_ARB                                           = 0x882B
	DRAW_BUFFER4_NV                                            = 0x8829
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	ACTIVE_RESOURCES                                           = 0x92F5
	MAX_SHININESS_NV                                           = 0x8504
	DRAW_BUFFER1_NV                                            = 0x8826
	BOOL_VEC2_ARB                                              = 0x8B57
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	MAX_TEXTURE_SIZE                                           = 0x0D33
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	MAP2_NORMAL                                                = 0x0DB2
	DYNAMIC_READ_ARB                                           = 0x88E9
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	CLIP_PLANE4                                                = 0x3004
	R16                                                        = 0x822A
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	PROGRAM_FORMAT_ARB                                         = 0x8876
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	INT_IMAGE_CUBE                                             = 0x905B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	TRIANGLES_ADJACENCY                                        = 0x000C
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	NUM_SAMPLE_COUNTS                                          = 0x9380
	LIGHTING_BIT                                               = 0x00000040
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	SRGB8_EXT                                                  = 0x8C41
	RED_SCALE                                                  = 0x0D14
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	FLOAT16_VEC3_NV                                            = 0x8FFA
	BLUE_BIAS                                                  = 0x0D1B
	FRAGMENT_SHADER_ATI                                        = 0x8920
	READ_FRAMEBUFFER                                           = 0x8CA8
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	ALPHA_SCALE                                                = 0x0D1C
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	WRITE_ONLY_OES                                             = 0x88B9
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	STENCIL_EXT                                                = 0x1802
	GL_4PASS_2_SGIS                                            = 0x80A6
	INTERPOLATE_EXT                                            = 0x8575
	SOURCE1_ALPHA_EXT                                          = 0x8589
	REG_21_ATI                                                 = 0x8936
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	EXT_vertex_array                                           = 1
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	PREVIOUS                                                   = 0x8578
	HILO_NV                                                    = 0x86F4
	DRAW_BUFFER9_NV                                            = 0x882E
	ARRAY_BUFFER_BINDING                                       = 0x8894
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	INDEX_MATERIAL_EXT                                         = 0x81B8
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	SAMPLER_CUBE_ARB                                           = 0x8B60
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	HIGH_INT                                                   = 0x8DF5
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	UNSIGNED_INT64_NV                                          = 0x140F
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	TEXTURE19                                                  = 0x84D3
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	SGIX_texture_lod_bias                                      = 1
	MIN                                                        = 0x8007
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	TEXTURE4_ARB                                               = 0x84C4
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	HI_SCALE_NV                                                = 0x870E
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	LUMINANCE_SNORM                                            = 0x9011
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	SGIX_ycrcb                                                 = 1
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	LIGHT5                                                     = 0x4005
	DEPTH24_STENCIL8                                           = 0x88F0
	REG_5_ATI                                                  = 0x8926
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	RGBA16_SNORM                                               = 0x8F9B
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	DRAW_BUFFER7_ARB                                           = 0x882C
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	STENCIL_WRITEMASK                                          = 0x0B98
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	POSITION                                                   = 0x1203
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	DEPTH_CLAMP_NV                                             = 0x864F
	SUB_ATI                                                    = 0x8965
	OFFSET                                                     = 0x92FC
	ALPHA_BIAS                                                 = 0x0D1D
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	NEGATIVE_X_EXT                                             = 0x87D9
	CURRENT_QUERY_EXT                                          = 0x8865
	MAP_COLOR                                                  = 0x0D10
	FRAMEBUFFER_BLEND                                          = 0x828B
	UNSIGNED_SHORT                                             = 0x1403
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	BGRA                                                       = 0x80E1
	UNPACK_RESAMPLE_OML                                        = 0x8985
	PURGEABLE_APPLE                                            = 0x8A1D
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	EQUAL                                                      = 0x0202
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	MATRIX7_NV                                                 = 0x8637
	CON_9_ATI                                                  = 0x894A
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	INT_SAMPLER_1D                                             = 0x8DC9
	IMAGE_2D                                                   = 0x904D
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	LUMINANCE8_SNORM                                           = 0x9015
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	TRACE_NAME_MESA                                            = 0x8756
	LUMINANCE32F_ARB                                           = 0x8818
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	COMBINER_INPUT_NV                                          = 0x8542
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	LEQUAL                                                     = 0x0203
	CURRENT_COLOR                                              = 0x0B00
	BLUE_SCALE                                                 = 0x0D1A
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	REFLECTION_MAP_NV                                          = 0x8512
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	LIGHT0                                                     = 0x4000
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	NEGATIVE_Y_EXT                                             = 0x87DA
	FLOAT_R16_NV                                               = 0x8884
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	CON_25_ATI                                                 = 0x895A
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	HINT_BIT                                                   = 0x00008000
	CONSTANT_ALPHA                                             = 0x8003
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	ADD_SIGNED_ARB                                             = 0x8574
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	SHADER_BINARY_DMP                                          = 0x9250
	STENCIL_BUFFER_BIT                                         = 0x00000400
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	TEXTURE31                                                  = 0x84DF
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	MAX_IMAGE_SAMPLES                                          = 0x906D
	RELATIVE_MOVE_TO_NV                                        = 0x03
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	FUNC_SUBTRACT                                              = 0x800A
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	DRAW_BUFFER5_ARB                                           = 0x882A
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	INT8_VEC3_NV                                               = 0x8FE2
	INTENSITY12                                                = 0x804C
	RG8_EXT                                                    = 0x822B
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	LESS                                                       = 0x0201
	REG_13_ATI                                                 = 0x892E
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	AUX3                                                       = 0x040C
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	RGB16I_EXT                                                 = 0x8D89
	BUFFER_DATA_SIZE                                           = 0x9303
	CCW                                                        = 0x0901
	BLEND_EQUATION_EXT                                         = 0x8009
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	DRAW_BUFFER11_ARB                                          = 0x8830
	INT_VEC4_ARB                                               = 0x8B55
	TEXTURE29_ARB                                              = 0x84DD
	OPERAND2_ALPHA                                             = 0x859A
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	CON_29_ATI                                                 = 0x895E
	ITALIC_BIT_NV                                              = 0x02
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	R3_G3_B2                                                   = 0x2A10
	RG16F_EXT                                                  = 0x822F
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	CONSTANT_BORDER_HP                                         = 0x8151
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	C4F_N3F_V3F                                                = 0x2A26
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	FLOAT_MAT2x4                                               = 0x8B66
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	IMAGE_1D                                                   = 0x904C
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	MULTISAMPLE                                                = 0x809D
	DRAW_BUFFER6_ATI                                           = 0x882B
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	IUI_V3F_EXT                                                = 0x81AE
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	IS_PER_PATCH                                               = 0x92E7
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	MATRIX6_NV                                                 = 0x8636
	SURFACE_REGISTERED_NV                                      = 0x86FD
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	FLOAT_RGBA32_NV                                            = 0x888B
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	COLOR_ATTACHMENT11                                         = 0x8CEB
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	VARIABLE_G_NV                                              = 0x8529
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	SPRITE_MODE_SGIX                                           = 0x8149
	STENCIL_BACK_FUNC                                          = 0x8800
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	SMOOTH                                                     = 0x1D01
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	QUERY_WAIT                                                 = 0x8E13
	IMAGE_BUFFER_EXT                                           = 0x9051
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	DEPTH_TEST                                                 = 0x0B71
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	WEIGHT_ARRAY_OES                                           = 0x86AD
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	PRIMARY_COLOR_EXT                                          = 0x8577
	ALPHA32F_ARB                                               = 0x8816
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	DOUBLE                                                     = 0x140A
	BLEND_DST_RGB_OES                                          = 0x80C8
	OPERAND0_ALPHA_EXT                                         = 0x8598
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	NAME_LENGTH                                                = 0x92F9
	LIGHT4                                                     = 0x4004
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	MATRIX_STRIDE                                              = 0x92FF
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	DOUBLE_MAT3                                                = 0x8F47
	EYE_LINE_SGIS                                              = 0x81F6
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	E_TIMES_F_NV                                               = 0x8531
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	VOLATILE_APPLE                                             = 0x8A1A
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	VERSION_3_2                                                = 1
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	CONSTANT_COLOR0_NV                                         = 0x852A
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	MATRIX2_NV                                                 = 0x8632
	VECTOR_EXT                                                 = 0x87BF
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	COLOR_INDEX4_EXT                                           = 0x80E4
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	RGBA_FLOAT32_APPLE                                         = 0x8814
	UNIFORM_TYPE                                               = 0x8A37
	QUERY_WAIT_NV                                              = 0x8E13
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	BLEND_SRC                                                  = 0x0BE1
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	SOURCE1_RGB_EXT                                            = 0x8581
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	SGIX_texture_coordinate_clamp                              = 1
	MAX_LIST_NESTING                                           = 0x0B31
	VENDOR                                                     = 0x1F00
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	OUTPUT_VERTEX_EXT                                          = 0x879A
	INT64_VEC4_NV                                              = 0x8FEB
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	CW                                                         = 0x0900
	RGB_FLOAT16_APPLE                                          = 0x881B
	RGB16UI_EXT                                                = 0x8D77
	RGBA32I_EXT                                                = 0x8D82
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	RIGHT                                                      = 0x0407
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	ALREADY_SIGNALED                                           = 0x911A
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	RG                                                         = 0x8227
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	RED_MAX_CLAMP_INGR                                         = 0x8564
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	READ_BUFFER                                                = 0x0C02
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	DYNAMIC_DRAW                                               = 0x88E8
	OBJECT_TYPE_APPLE                                          = 0x9112
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	LINES_ADJACENCY_EXT                                        = 0x000A
	MATRIX15_ARB                                               = 0x88CF
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	RGB9_E5                                                    = 0x8C3D
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	STENCIL_INDEX4_EXT                                         = 0x8D47
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	REPLACE_EXT                                                = 0x8062
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	MODELVIEW1_ARB                                             = 0x850A
	BUFFER_ACCESS                                              = 0x88BB
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	INT_10_10_10_2_OES                                         = 0x8DF7
	MODELVIEW7_ARB                                             = 0x8727
	DRAW_BUFFER10                                              = 0x882F
	R16_SNORM                                                  = 0x8F98
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	PROGRAM_POINT_SIZE                                         = 0x8642
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	OP_DOT3_EXT                                                = 0x8784
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	FACTOR_MIN_AMD                                             = 0x901C
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	BGR_EXT                                                    = 0x80E0
	SHADER_OPERATION_NV                                        = 0x86DF
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	CND_ATI                                                    = 0x896A
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	UNPACK_SKIP_IMAGES                                         = 0x806D
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	DEBUG_ASSERT_MESA                                          = 0x875B
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	LINEAR_DETAIL_SGIS                                         = 0x8097
	RGB_SCALE_EXT                                              = 0x8573
	DEPTH_CLAMP                                                = 0x864F
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	MAX_VARYING_VECTORS                                        = 0x8DFC
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	PERTURB_EXT                                                = 0x85AE
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	DRAW_BUFFER0_ATI                                           = 0x8825
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	TEXTURE12                                                  = 0x84CC
	DOT4_ATI                                                   = 0x8967
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	SOURCE2_RGB                                                = 0x8582
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	WRITE_ONLY                                                 = 0x88B9
	MATRIX2_ARB                                                = 0x88C2
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	RED_SNORM                                                  = 0x8F90
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	DRAW_BUFFER2_ARB                                           = 0x8827
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	UNIFORM_OFFSET                                             = 0x8A3B
	SHADER_OBJECT_ARB                                          = 0x8B48
	RENDERBUFFER_WIDTH                                         = 0x8D42
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	INTENSITY8_SNORM                                           = 0x9017
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	SCALE_BY_TWO_NV                                            = 0x853E
	VERTEX_PROGRAM_NV                                          = 0x8620
	MODELVIEW6_ARB                                             = 0x8726
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	CON_22_ATI                                                 = 0x8957
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	ALPHA16                                                    = 0x803E
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	EXT_blend_logic_op                                         = 1
	VERTEX_ARRAY_SIZE                                          = 0x807A
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	COMBINER2_NV                                               = 0x8552
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	TEXTURE10                                                  = 0x84CA
	MATRIX7_ARB                                                = 0x88C7
	POINT_TOKEN                                                = 0x0701
	ALPHA8                                                     = 0x803C
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	DYNAMIC_ATI                                                = 0x8761
	EYE_RADIAL_NV                                              = 0x855B
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	NOTEQUAL                                                   = 0x0205
	LUMINANCE8                                                 = 0x8040
	OP_MADD_EXT                                                = 0x8788
	DRAW_BUFFER2_ATI                                           = 0x8827
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	FILTER                                                     = 0x829A
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	VERTEX_SHADER_ARB                                          = 0x8B31
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	RGBA_SNORM                                                 = 0x8F93
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	INDEX_OFFSET                                               = 0x0D13
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	SAMPLER_BINDING                                            = 0x8919
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	INTERLACE_READ_OML                                         = 0x8981
	MAX_SAMPLES_NV                                             = 0x8D57
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	DYNAMIC_COPY                                               = 0x88EA
	GREEN_INTEGER                                              = 0x8D95
	ACTIVE_VARIABLES                                           = 0x9305
	LINE_STRIP_ADJACENCY                                       = 0x000B
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	ETC1_SRGB8_NV                                              = 0x88EE
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	COPY_PIXEL_TOKEN                                           = 0x0706
	LINE_WIDTH                                                 = 0x0B21
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	SAMPLES_PASSED                                             = 0x8914
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	MAX_CLIP_PLANES                                            = 0x0D32
	CONST_EYE_NV                                               = 0x86E5
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	FLOAT_R_NV                                                 = 0x8880
	SGIX_scalebias_hint                                        = 1
	SIGNED_LUMINANCE_NV                                        = 0x8701
	FLOAT_VEC4                                                 = 0x8B52
	INCR                                                       = 0x1E02
	CONVOLUTION_FORMAT                                         = 0x8017
	TEXTURE23_ARB                                              = 0x84D7
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	STATIC_DRAW_ARB                                            = 0x88E4
	BIAS_BIT_ATI                                               = 0x00000008
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	CON_12_ATI                                                 = 0x894D
	SKIP_COMPONENTS3_NV                                        = -4
	POLYGON_OFFSET_LINE                                        = 0x2A02
	RGB10_A2                                                   = 0x8059
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	INDEX_CLEAR_VALUE                                          = 0x0C20
	COMBINER0_NV                                               = 0x8550
	TEXTURE_MIN_LOD                                            = 0x813A
	MODELVIEW8_ARB                                             = 0x8728
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	CON_13_ATI                                                 = 0x894E
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	COLOR3_BIT_PGI                                             = 0x00010000
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	INTENSITY16UI_EXT                                          = 0x8D79
	CLAMP_TO_BORDER_NV                                         = 0x812D
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	RG16F                                                      = 0x822F
	FOG_COORD                                                  = 0x8451
	TEXTURE30_ARB                                              = 0x84DE
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	SGI_texture_color_table                                    = 1
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	ALPHA_INTEGER_EXT                                          = 0x8D97
	EXT_packed_pixels                                          = 1
	STENCIL_FAIL                                               = 0x0B94
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	LINE_BIT                                                   = 0x00000004
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	TEXTURE22_ARB                                              = 0x84D6
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	DEPTH_RENDERABLE                                           = 0x8287
	PROGRAM_PARAMETER_NV                                       = 0x8644
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	TEXTURE15_ARB                                              = 0x84CF
	CONSTANT_EXT                                               = 0x8576
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	TEXTURE_3D                                                 = 0x806F
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	OPERAND1_RGB_EXT                                           = 0x8591
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	COMBINER7_NV                                               = 0x8557
	CULL_FRAGMENT_NV                                           = 0x86E7
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	SRGB                                                       = 0x8C40
	TRANSFORM_BIT                                              = 0x00001000
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	QUERY                                                      = 0x82E3
	MATRIX4_NV                                                 = 0x8634
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	MAP1_VERTEX_3                                              = 0x0D97
	SIGNED_NEGATE_NV                                           = 0x853D
	SYNC_OBJECT_APPLE                                          = 0x8A53
	TEXTURE3_ARB                                               = 0x84C3
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	RENDER_MODE                                                = 0x0C40
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	FRONT_RIGHT                                                = 0x0401
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	SOURCE0_RGB                                                = 0x8580
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	YCBAYCR8A_4224_NV                                          = 0x9032
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	TRIANGLES                                                  = 0x0004
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	BACK_LEFT                                                  = 0x0402
	RGB2_EXT                                                   = 0x804E
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	IMAGE_BINDING_FORMAT                                       = 0x906E
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	FLOAT_RG16_NV                                              = 0x8886
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	INDEX_ARRAY                                                = 0x8077
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	COLOR_INDEX1_EXT                                           = 0x80E2
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	SLIM12S_SGIX                                               = 0x831F
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	COMPILE_AND_EXECUTE                                        = 0x1301
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	DS_SCALE_NV                                                = 0x8710
	BUFFER_MAPPED                                              = 0x88BC
	SWIZZLE_STR_ATI                                            = 0x8976
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	REG_27_ATI                                                 = 0x893C
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	BLUE                                                       = 0x1905
	POINT_SIZE_MAX                                             = 0x8127
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	INT_VEC2_ARB                                               = 0x8B53
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	STENCIL_BITS                                               = 0x0D57
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	R16UI                                                      = 0x8234
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	SUBTRACT_ARB                                               = 0x84E7
	C4UB_V3F                                                   = 0x2A23
	FUNC_SUBTRACT_OES                                          = 0x800A
	COLOR_TABLE_WIDTH                                          = 0x80D9
	POINT_SIZE_MIN_EXT                                         = 0x8126
	STENCIL_INDEX4_OES                                         = 0x8D47
	ALPHA8I_EXT                                                = 0x8D90
	MINMAX_EXT                                                 = 0x802E
	INDEX_TEST_EXT                                             = 0x81B5
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	PALETTE8_RGB8_OES                                          = 0x8B95
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	SPARE1_NV                                                  = 0x852F
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	GREEN_INTEGER_EXT                                          = 0x8D95
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	OPERAND1_RGB_ARB                                           = 0x8591
	DSDT8_NV                                                   = 0x8709
	BUFFER_USAGE                                               = 0x8765
	RGBA_FLOAT16_APPLE                                         = 0x881A
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	TEXTURE13_ARB                                              = 0x84CD
	SOURCE2_RGB_ARB                                            = 0x8582
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	LINK_STATUS                                                = 0x8B82
	PALETTE8_RGBA8_OES                                         = 0x8B96
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	IMAGE_BINDING_NAME                                         = 0x8F3A
	OUT_OF_MEMORY                                              = 0x0505
	DOUBLEBUFFER                                               = 0x0C32
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	RENDER                                                     = 0x1C00
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	MATRIX27_ARB                                               = 0x88DB
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	SAMPLE_MASK_NV                                             = 0x8E51
	TEXTURE19_ARB                                              = 0x84D3
	MVP_MATRIX_EXT                                             = 0x87E3
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	RGB8                                                       = 0x8051
	SAMPLE_MASK_SGIS                                           = 0x80A0
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	FRAGMENT_DEPTH                                             = 0x8452
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	VERTEX_STREAM3_ATI                                         = 0x876F
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	READ_PIXELS                                                = 0x828C
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	INTENSITY4                                                 = 0x804A
	TEXTURE0                                                   = 0x84C0
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	CON_28_ATI                                                 = 0x895D
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	IMAGE_BUFFER                                               = 0x9051
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	POINT_SMOOTH                                               = 0x0B10
	STENCIL_TEST                                               = 0x0B90
	NAME_STACK_DEPTH                                           = 0x0D70
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	RGB_FLOAT16_ATI                                            = 0x881B
	MATRIX30_ARB                                               = 0x88DE
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	INDEX_ARRAY_STRIDE                                         = 0x8086
	FIXED_OES                                                  = 0x140C
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	COORD_REPLACE_ARB                                          = 0x8862
	RGB_422_APPLE                                              = 0x8A1F
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	PATH_JOIN_STYLE_NV                                         = 0x9079
	COMBINE_ALPHA_EXT                                          = 0x8572
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	SAMPLER_2D_ARB                                             = 0x8B5E
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	FOG_MODE                                                   = 0x0B65
	REPLACE                                                    = 0x1E01
	FLOAT_RG_NV                                                = 0x8881
	BOOL_VEC3_ARB                                              = 0x8B58
	INT64_VEC2_NV                                              = 0x8FE9
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	DEBUG_OBJECT_MESA                                          = 0x8759
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	IMAGE_CUBE                                                 = 0x9050
	VARIABLE_F_NV                                              = 0x8528
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	PATH_FILL_MASK_NV                                          = 0x9081
	DECR_WRAP_OES                                              = 0x8508
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	GEOMETRY_SHADER                                            = 0x8DD9
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	RGBA4_EXT                                                  = 0x8056
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	MODELVIEW                                                  = 0x1700
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	STATIC_COPY_ARB                                            = 0x88E6
	LUMINANCE32UI_EXT                                          = 0x8D74
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	SPRITE_AXIAL_SGIX                                          = 0x814C
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	STENCIL_INDEX8_EXT                                         = 0x8D48
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	X_EXT                                                      = 0x87D5
	BLUE_INTEGER_EXT                                           = 0x8D96
	FIRST_TO_REST_NV                                           = 0x90AF
	MIPMAP                                                     = 0x8293
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	GREEN_BIAS                                                 = 0x0D19
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	MODELVIEW20_ARB                                            = 0x8734
	SET_AMD                                                    = 0x874A
	INT8_VEC2_NV                                               = 0x8FE1
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	GL_3D                                                      = 0x0601
	TEXTURE25                                                  = 0x84D9
	SLUMINANCE8_NV                                             = 0x8C47
	IMAGE_1D_ARRAY                                             = 0x9052
	TEXTURE_SAMPLES                                            = 0x9106
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	CON_16_ATI                                                 = 0x8951
	SLUMINANCE                                                 = 0x8C46
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	SGIX_polynomial_ffd                                        = 1
	LINES_ADJACENCY_ARB                                        = 0x000A
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	PRIMITIVE_RESTART_NV                                       = 0x8558
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	DRAW_BUFFER12                                              = 0x8831
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	VERTEX_SUBROUTINE                                          = 0x92E8
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	FILTER4_SGIS                                               = 0x8146
	T2F_IUI_V2F_EXT                                            = 0x81B1
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	MATRIX25_ARB                                               = 0x88D9
	COMPRESSED_SRGB                                            = 0x8C48
	MATRIX_MODE                                                = 0x0BA0
	MAP1_VERTEX_4                                              = 0x0D98
	AMBIENT                                                    = 0x1200
	COMBINE4_NV                                                = 0x8503
	REFLECTION_MAP_ARB                                         = 0x8512
	SRC1_RGB                                                   = 0x8581
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	SGIX_subsample                                             = 1
	DEBUG_PRINT_MESA                                           = 0x875A
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	BUFFER_ACCESS_OES                                          = 0x88BB
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	BUFFER_MAP_OFFSET                                          = 0x9121
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	SGIS_pixel_texture                                         = 1
	SGIS_point_parameters                                      = 1
	FOG_COORD_ARRAY                                            = 0x8457
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	COMBINE_ALPHA                                              = 0x8572
	VERTEX_SOURCE_ATI                                          = 0x8774
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	CLIP_DISTANCE6                                             = 0x3006
	MODELVIEW19_ARB                                            = 0x8733
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	PERCENTAGE_AMD                                             = 0x8BC3
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	POINTS                                                     = 0x0000
	DRAW_BUFFER14_ATI                                          = 0x8833
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	CONTINUOUS_AMD                                             = 0x9007
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	INT_IMAGE_2D_RECT                                          = 0x905A
	SYNC_FLAGS_APPLE                                           = 0x9115
	COLOR_ARRAY_EXT                                            = 0x8076
	REPLACE_MIDDLE_SUN                                         = 0x0002
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	STATIC_ATI                                                 = 0x8760
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	LIST_BIT                                                   = 0x00020000
	NEVER                                                      = 0x0200
	FOG_HINT                                                   = 0x0C54
	PACK_LSB_FIRST                                             = 0x0D01
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	SUBTRACT                                                   = 0x84E7
	COMBINER_BIAS_NV                                           = 0x8549
	SGIS_generate_mipmap                                       = 1
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	LUMINANCE32I_EXT                                           = 0x8D86
	RGBA_INTEGER_EXT                                           = 0x8D99
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	DOUBLE_MAT4_EXT                                            = 0x8F48
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	RGB8UI                                                     = 0x8D7D
	CONTEXT_PROFILE_MASK                                       = 0x9126
	UNIFORM                                                    = 0x92E1
	VIEWPORT                                                   = 0x0BA2
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	MODELVIEW15_ARB                                            = 0x872F
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	SGIX_flush_raster                                          = 1
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	FIXED                                                      = 0x140C
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	FOG_COORD_SRC                                              = 0x8450
	SOURCE1_RGB                                                = 0x8581
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	UNIFORM_BLOCK                                              = 0x92E2
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	UTF16_NV                                                   = 0x909B
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	TESS_GEN_MODE                                              = 0x8E76
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	SOURCE2_ALPHA_EXT                                          = 0x858A
	CON_5_ATI                                                  = 0x8946
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	TESS_EVALUATION_SHADER                                     = 0x8E87
	LARGE_CW_ARC_TO_NV                                         = 0x18
	RGBA                                                       = 0x1908
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	SLIM10U_SGIX                                               = 0x831E
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	GPU_ADDRESS_NV                                             = 0x8F34
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	EXP2                                                       = 0x0801
	AUX_BUFFERS                                                = 0x0C00
	STENCIL                                                    = 0x1802
	CON_18_ATI                                                 = 0x8953
	LOWER_LEFT                                                 = 0x8CA1
	COMBINE_RGB_EXT                                            = 0x8571
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	REG_24_ATI                                                 = 0x8939
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	FUNC_SUBTRACT_EXT                                          = 0x800A
	INTENSITY_EXT                                              = 0x8049
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	PARTIAL_SUCCESS_NV                                         = 0x902E
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	VERTEX_BLEND_ARB                                           = 0x86A7
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	DRAW_BUFFER13_ARB                                          = 0x8832
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	FLOAT_VEC2                                                 = 0x8B50
	RED_INTEGER                                                = 0x8D94
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	COLOR_ATTACHMENT2                                          = 0x8CE2
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	RG_EXT                                                     = 0x8227
	HALF_BIT_ATI                                               = 0x00000008
	BGR_INTEGER                                                = 0x8D9A
	CURRENT_TIME_NV                                            = 0x8E28
	SGIX_blend_alpha_minmax                                    = 1
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	QUERY_RESULT_EXT                                           = 0x8866
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	COLOR_ATTACHMENT3                                          = 0x8CE3
	BUFFER_BINDING                                             = 0x9302
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	VARIANT_ARRAY_EXT                                          = 0x87E8
	DOT3_ATI                                                   = 0x8966
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	LINE_STIPPLE                                               = 0x0B24
	INTENSITY12_EXT                                            = 0x804C
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	INTENSITY8                                                 = 0x804B
	CONSTANT_ALPHA_EXT                                         = 0x8003
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	DECR_WRAP_EXT                                              = 0x8508
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	INDEX_ARRAY_LIST_IBM                                       = 103073
	FOG_COLOR                                                  = 0x0B66
	OPERAND0_RGB_ARB                                           = 0x8590
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	BUFFER_OBJECT_EXT                                          = 0x9151
	DOT3_RGB                                                   = 0x86AE
	DRAW_BUFFER6                                               = 0x882B
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	SIGNED_NORMALIZED                                          = 0x8F9C
	VIEW_CLASS_48_BITS                                         = 0x82C7
	NORMAL_MAP_EXT                                             = 0x8511
	TRACE_MASK_MESA                                            = 0x8755
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	SGIS_texture_lod                                           = 1
	DEPTH_EXT                                                  = 0x1801
	SPHERE_MAP                                                 = 0x2402
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	DRAW_BUFFER7                                               = 0x882C
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	FEEDBACK                                                   = 0x1C01
	BUFFER                                                     = 0x82E0
	MUL_ATI                                                    = 0x8964
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	RETURN                                                     = 0x0102
	GL_422_AVERAGE_EXT                                         = 0x80CE
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	SGIX_async_pixel                                           = 1
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	TEXTURE24                                                  = 0x84D8
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	ARB_imaging                                                = 1
	COLOR_INDEX8_EXT                                           = 0x80E5
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	REG_29_ATI                                                 = 0x893E
	PATH_GEN_MODE_NV                                           = 0x90B0
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	PREVIOUS_ARB                                               = 0x8578
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	INVALID_OPERATION                                          = 0x0502
	NORMALIZE                                                  = 0x0BA1
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	DEPTH_STENCIL_MESA                                         = 0x8750
	COLOR_ARRAY_LIST_IBM                                       = 103072
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	ALPHA_INTEGER                                              = 0x8D97
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	RG8I                                                       = 0x8237
	COMBINE_RGB_ARB                                            = 0x8571
	MODELVIEW27_ARB                                            = 0x873B
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	R32UI                                                      = 0x8236
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	TEXTURE17_ARB                                              = 0x84D1
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	STREAM_DRAW                                                = 0x88E0
	GREEN_BIT_ATI                                              = 0x00000002
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	REPLACE_VALUE_AMD                                          = 0x874B
	MATRIX4_ARB                                                = 0x88C4
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	VARIABLE_C_NV                                              = 0x8525
	RGB16F_ARB                                                 = 0x881B
	LINE_SMOOTH_HINT                                           = 0x0C52
	MAP1_TANGENT_EXT                                           = 0x8444
	REG_18_ATI                                                 = 0x8933
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	DT_SCALE_NV                                                = 0x8711
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	REG_4_ATI                                                  = 0x8925
	DOT2_ADD_ATI                                               = 0x896C
	ACCUM_BLUE_BITS                                            = 0x0D5A
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	LINE_STRIP                                                 = 0x0003
	RG8UI                                                      = 0x8238
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	OPERAND1_ALPHA_EXT                                         = 0x8599
	HILO16_NV                                                  = 0x86F8
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	STENCIL_FUNC                                               = 0x0B92
	COLOR_ARRAY_STRIDE                                         = 0x8083
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	COMPUTE_SUBROUTINE                                         = 0x92ED
	MAX_EVAL_ORDER                                             = 0x0D30
	QUAD_ALPHA8_SGIS                                           = 0x811F
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	PACK_SWAP_BYTES                                            = 0x0D00
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	CONSTANT_COLOR_EXT                                         = 0x8001
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	CLIP_DISTANCE4                                             = 0x3004
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	TRACK_MATRIX_NV                                            = 0x8648
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	SWIZZLE_STRQ_ATI                                           = 0x897A
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	TEXTURE_BORDER                                             = 0x1005
	COLOR_INDEXES                                              = 0x1603
	FOG_FUNC_SGIS                                              = 0x812A
	OBJECT_LINE_SGIS                                           = 0x81F7
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	DISCARD_ATI                                                = 0x8763
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	FRAME_NV                                                   = 0x8E26
	RGBA16                                                     = 0x805B
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	SRGB8_NV                                                   = 0x8C41
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	CLIP_PLANE1                                                = 0x3001
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	SAMPLES_3DFX                                               = 0x86B4
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	DOUBLE_EXT                                                 = 0x140A
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	MAP1_NORMAL                                                = 0x0D92
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	AFFINE_3D_NV                                               = 0x9094
	FOG_COORDINATE_SOURCE                                      = 0x8450
	CONDITION_SATISFIED                                        = 0x911C
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	BOOL_ARB                                                   = 0x8B56
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	YCBCR_MESA                                                 = 0x8757
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	PATH_COORD_COUNT_NV                                        = 0x909E
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	LINE                                                       = 0x1B01
	LUMINANCE16                                                = 0x8042
	CMYK_EXT                                                   = 0x800C
	FOG_OFFSET_SGIX                                            = 0x8198
	LUMINANCE16F_ARB                                           = 0x881E
	FLOAT_MAT4                                                 = 0x8B5C
	LEFT                                                       = 0x0406
	RGBA_S3TC                                                  = 0x83A2
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	COMBINER3_NV                                               = 0x8553
	STREAM_DRAW_ARB                                            = 0x88E0
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	IMAGE_2D_ARRAY                                             = 0x9053
	ACCUM_BUFFER_BIT                                           = 0x00000200
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	SAMPLES_SGIS                                               = 0x80A9
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	LO_SCALE_NV                                                = 0x870F
	CON_7_ATI                                                  = 0x8948
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	R16F_EXT                                                   = 0x822D
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	STREAM_READ                                                = 0x88E1
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	FLOAT_MAT2x3                                               = 0x8B65
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	UNPACK_LSB_FIRST                                           = 0x0CF1
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	RGB10                                                      = 0x8052
	MAP2_BINORMAL_EXT                                          = 0x8447
	FIXED_ONLY_ARB                                             = 0x891D
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	TEXTURE_RED_SIZE                                           = 0x805C
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	ZERO_EXT                                                   = 0x87DD
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	INT16_VEC3_NV                                              = 0x8FE6
	COLOR_BUFFER_BIT                                           = 0x00004000
	RGBA_MODE                                                  = 0x0C31
	GL_422_REV_EXT                                             = 0x80CD
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	VIBRANCE_BIAS_NV                                           = 0x8719
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	SGIS_point_line_texgen                                     = 1
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	DEPTH_COMPONENT24                                          = 0x81A6
	RG8                                                        = 0x822B
	MATRIX3_NV                                                 = 0x8633
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	INTENSITY32F_ARB                                           = 0x8817
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	CON_31_ATI                                                 = 0x8960
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	SYNC_STATUS_APPLE                                          = 0x9114
	SGIX_resample                                              = 1
	SAMPLER_1D_SHADOW                                          = 0x8B61
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	MAX_NAME_LENGTH                                            = 0x92F6
	CONSTANT_ARB                                               = 0x8576
	DRAW_BUFFER8_ATI                                           = 0x882D
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	MAX_INTEGER_SAMPLES                                        = 0x9110
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	COLOR_TABLE                                                = 0x80D0
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	GL_4PASS_1_EXT                                             = 0x80A5
	COUNT_DOWN_NV                                              = 0x9089
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	RGB_SCALE_ARB                                              = 0x8573
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	FUNC_ADD_EXT                                               = 0x8006
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	DOT_PRODUCT_NV                                             = 0x86EC
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	SGIS_texture_filter4                                       = 1
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	PACK_RESAMPLE_OML                                          = 0x8984
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	HISTOGRAM                                                  = 0x8024
	VERTEX_STREAM4_ATI                                         = 0x8770
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	CLAMP_READ_COLOR                                           = 0x891C
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	QUERY_OBJECT_EXT                                           = 0x9153
	R8UI                                                       = 0x8232
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	FENCE_STATUS_NV                                            = 0x84F3
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	FLOAT_VEC3                                                 = 0x8B51
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	VERSION_1_2                                                = 1
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	REG_1_ATI                                                  = 0x8922
	REG_15_ATI                                                 = 0x8930
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	SGIX_tag_sample_buffer                                     = 1
	EVAL_BIT                                                   = 0x00010000
	TEXTURE_MAX_LOD                                            = 0x813B
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	SLIM8U_SGIX                                                = 0x831D
	MATRIX0_ARB                                                = 0x88C0
	UNDEFINED_APPLE                                            = 0x8A1C
	RGBA32I                                                    = 0x8D82
	EXT_cmyka                                                  = 1
	FOG_BIT                                                    = 0x00000080
	NAND                                                       = 0x150E
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	TEXTURE_RECTANGLE                                          = 0x84F5
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	COLOR_ENCODING                                             = 0x8296
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	REG_23_ATI                                                 = 0x8938
	SAMPLER_BUFFER_AMD                                         = 0x9001
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	SET                                                        = 0x150F
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	COLOR_ATTACHMENT9                                          = 0x8CE9
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	CLIP_DISTANCE3                                             = 0x3003
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	EXT_texture_object                                         = 1
	COPY                                                       = 0x1503
	CURRENT_TANGENT_EXT                                        = 0x843B
	R1UI_V3F_SUN                                               = 0x85C4
	MEDIUM_FLOAT                                               = 0x8DF1
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	MODELVIEW31_ARB                                            = 0x873F
	RGB32F_ARB                                                 = 0x8815
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	FENCE_CONDITION_NV                                         = 0x84F4
	EXPAND_NEGATE_NV                                           = 0x8539
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	FRAGMENT_SHADER                                            = 0x8B30
	DEPTH_WRITEMASK                                            = 0x0B72
	R8                                                         = 0x8229
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	PRESENT_DURATION_NV                                        = 0x8E2B
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	MATRIX0_NV                                                 = 0x8630
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	VERTEX23_BIT_PGI                                           = 0x00000004
	DOT3_RGBA                                                  = 0x86AF
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	CON_30_ATI                                                 = 0x895F
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	MAX_IMAGE_UNITS                                            = 0x8F38
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	OPERAND2_ALPHA_EXT                                         = 0x859A
	SAMPLER_2D                                                 = 0x8B5E
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	OPERAND1_ALPHA_ARB                                         = 0x8599
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	RGBA_FLOAT32_ATI                                           = 0x8814
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	EXP                                                        = 0x0800
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	COLOR_ATTACHMENT14                                         = 0x8CEE
	EXT_polygon_offset                                         = 1
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	COLOR_INDEX16_EXT                                          = 0x80E7
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	ALPHA8_EXT                                                 = 0x803C
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	TESSELLATION_MODE_AMD                                      = 0x9004
	RENDERER                                                   = 0x1F01
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	PATH_FORMAT_PS_NV                                          = 0x9071
	AFFINE_2D_NV                                               = 0x9092
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	N3F_V3F                                                    = 0x2A25
	OPERAND1_RGB                                               = 0x8591
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	INT_SAMPLER_3D                                             = 0x8DCB
	TEXTURE30                                                  = 0x84DE
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	ZOOM_Y                                                     = 0x0D17
	RGBA2                                                      = 0x8055
	SHADER                                                     = 0x82E1
	ALPHA_MAX_SGIX                                             = 0x8321
	CON_23_ATI                                                 = 0x8958
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	SAMPLER_2D_SHADOW                                          = 0x8B62
	FOG_INDEX                                                  = 0x0B61
	AND_INVERTED                                               = 0x1504
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	SPRITE_AXIS_SGIX                                           = 0x814A
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	SAMPLES_ARB                                                = 0x80A9
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	TEXTURE18                                                  = 0x84D2
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	BLEND_SRC_RGB_OES                                          = 0x80C9
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	FLOAT_VEC3_ARB                                             = 0x8B51
	BUMP_ENVMAP_ATI                                            = 0x877B
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	ADJACENT_PAIRS_NV                                          = 0x90AE
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	TRIANGLE_LIST_SUN                                          = 0x81D7
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	SRC2_RGB                                                   = 0x8582
	INVARIANT_EXT                                              = 0x87C2
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	INT64_NV                                                   = 0x140E
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	COLOR_ATTACHMENT5                                          = 0x8CE5
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	PATH_MITER_LIMIT_NV                                        = 0x907A
	SAMPLE_COVERAGE                                            = 0x80A0
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	DECR                                                       = 0x1E03
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	BOOL_VEC4_ARB                                              = 0x8B59
	TEXTURE_1D_ARRAY                                           = 0x8C18
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	GL_1PASS_SGIS                                              = 0x80A1
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	LINEAR_ATTENUATION                                         = 0x1208
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	DRAW_BUFFER3                                               = 0x8828
	DRAW_BUFFER11_ATI                                          = 0x8830
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	RGB5_A1_EXT                                                = 0x8057
	IUI_N3F_V3F_EXT                                            = 0x81B0
	R8I                                                        = 0x8231
	SGIX_impact_pixel_texture                                  = 1
	QUADRATIC_ATTENUATION                                      = 0x1209
	SWIZZLE_STQ_ATI                                            = 0x8977
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	ALPHA_MIN_SGIX                                             = 0x8320
	GL_2D                                                      = 0x0600
	FOG_END                                                    = 0x0B64
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	IUI_N3F_V2F_EXT                                            = 0x81AF
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	UNPACK_ALIGNMENT                                           = 0x0CF5
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	VIEW_CLASS_96_BITS                                         = 0x82C5
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	LINE_TOKEN                                                 = 0x0702
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	SAMPLER_2D_RECT                                            = 0x8B63
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	COPY_WRITE_BUFFER                                          = 0x8F37
	AND_REVERSE                                                = 0x1502
	CONVOLUTION_WIDTH                                          = 0x8018
	CONTEXT_FLAGS                                              = 0x821E
	STENCIL_INDEX1_EXT                                         = 0x8D46
	LAYER_NV                                                   = 0x8DAA
	MODELVIEW12_ARB                                            = 0x872C
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	LINE_WIDTH_RANGE                                           = 0x0B22
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	MAP_WRITE_BIT                                              = 0x0002
	SUBPIXEL_BITS                                              = 0x0D50
	REDUCE                                                     = 0x8016
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	CMYKA_EXT                                                  = 0x800D
	GL_4PASS_0_SGIS                                            = 0x80A4
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	SGIX_sprite                                                = 1
	TEXTURE_GEN_R                                              = 0x0C62
	INTENSITY8_EXT                                             = 0x804B
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	SIGNED_RGBA8_NV                                            = 0x86FC
	DRAW_BUFFER14                                              = 0x8833
	READ_ONLY_ARB                                              = 0x88B8
	RGB_S3TC                                                   = 0x83A0
	COMBINER1_NV                                               = 0x8551
	MATRIX13_ARB                                               = 0x88CD
	SGIX_convolution_accuracy                                  = 1
	LINE_LOOP                                                  = 0x0002
	VERTEX_ARRAY_TYPE                                          = 0x807B
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	MODELVIEW22_ARB                                            = 0x8736
	MATRIX21_ARB                                               = 0x88D5
	SRGB_EXT                                                   = 0x8C40
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	MAX_LAYERS                                                 = 0x8281
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	INDEX_ARRAY_EXT                                            = 0x8077
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	PROVOKING_VERTEX                                           = 0x8E4F
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	YCRCBA_SGIX                                                = 0x8319
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	STENCIL_INDEX                                              = 0x1901
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	EXT_point_parameters                                       = 1
	STEREO                                                     = 0x0C33
	OBJECT_POINT_SGIS                                          = 0x81F5
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	FLOAT_VEC2_ARB                                             = 0x8B50
	INT16_VEC4_NV                                              = 0x8FE7
	TEXTURE_COMPRESSED                                         = 0x86A1
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	SEPARABLE_2D                                               = 0x8012
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	SLUMINANCE_EXT                                             = 0x8C46
	TEXTURE0_ARB                                               = 0x84C0
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	INDEX_ARRAY_POINTER                                        = 0x8091
	CONVOLUTION_2D_EXT                                         = 0x8011
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	MATRIX16_ARB                                               = 0x88D0
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	DOUBLE_MAT3x2                                              = 0x8F4B
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	RGB8I_EXT                                                  = 0x8D8F
	GREEN_SCALE                                                = 0x0D18
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	RG32F                                                      = 0x8230
	TESS_CONTROL_TEXTURE                                       = 0x829C
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	MINOR_VERSION                                              = 0x821C
	DSDT_MAG_NV                                                = 0x86F6
	SIGNED_HILO_NV                                             = 0x86F9
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	DEPTH_BUFFER_BIT                                           = 0x00000100
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	SPARE0_NV                                                  = 0x852E
	SRC2_ALPHA                                                 = 0x858A
	MODELVIEW11_ARB                                            = 0x872B
	MODELVIEW24_ARB                                            = 0x8738
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	COLOR_ARRAY                                                = 0x8076
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	COLOR_EXT                                                  = 0x1800
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	POINT_SPRITE                                               = 0x8861
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	MATRIX31_ARB                                               = 0x88DF
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	T2F_C3F_V3F                                                = 0x2A2A
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	RG16                                                       = 0x822C
	VIEW_CLASS_24_BITS                                         = 0x82C9
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	PROGRAM_LENGTH_NV                                          = 0x8627
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	RGBA16I_EXT                                                = 0x8D88
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	TIMEOUT_EXPIRED                                            = 0x911B
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	RGBA32UI_EXT                                               = 0x8D70
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	CURRENT_PROGRAM                                            = 0x8B8D
	PRIMITIVE_ID_NV                                            = 0x8C7C
	DEPTH_BITS                                                 = 0x0D56
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	STORAGE_CACHED_APPLE                                       = 0x85BE
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	ADD_ATI                                                    = 0x8963
	UNIFORM_BUFFER_START                                       = 0x8A29
	IMAGE_CUBE_EXT                                             = 0x9050
	TIME_ELAPSED                                               = 0x88BF
	MATRIX1_ARB                                                = 0x88C1
	SAMPLE_SHADING_ARB                                         = 0x8C36
	FIXED_ONLY                                                 = 0x891D
	RESTART_PATH_NV                                            = 0xF0
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	RGB5                                                       = 0x8050
	ASYNC_MARKER_SGIX                                          = 0x8329
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	ATC_RGB_AMD                                                = 0x8C92
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	PATCHES                                                    = 0x000E
	GL_2PASS_1_SGIS                                            = 0x80A3
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	PRIMARY_COLOR_ARB                                          = 0x8577
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	SAMPLE_POSITION                                            = 0x8E50
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	ACCUM_GREEN_BITS                                           = 0x0D59
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	RGB32UI                                                    = 0x8D71
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	SCALE_BY_FOUR_NV                                           = 0x853F
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	PACK_INVERT_MESA                                           = 0x8758
	TIME_ELAPSED_EXT                                           = 0x88BF
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	CON_0_ATI                                                  = 0x8941
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	PACK_ALIGNMENT                                             = 0x0D05
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	OPERAND1_ALPHA                                             = 0x8599
	COLOR_ATTACHMENT7                                          = 0x8CE7
	MAX_SAMPLES_EXT                                            = 0x8D57
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	RGB16_EXT                                                  = 0x8054
	TEXTURE_WRAP_R_OES                                         = 0x8072
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	DEPTH_ATTACHMENT                                           = 0x8D00
	LOAD                                                       = 0x0101
	QUAD_MESH_SUN                                              = 0x8614
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	FRAMEBUFFER_EXT                                            = 0x8D40
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	SPOT_EXPONENT                                              = 0x1205
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	FLOAT_MAT4x2                                               = 0x8B69
	SRGB8                                                      = 0x8C41
	STENCIL_INDEX1                                             = 0x8D46
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	MAP2_VERTEX_4                                              = 0x0DB8
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	DOUBLE_MAT2x3                                              = 0x8F49
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	LIGHTING                                                   = 0x0B50
	LUMINANCE4_ALPHA4                                          = 0x8043
	TEXTURE8_ARB                                               = 0x84C8
	IMAGE_2D_RECT                                              = 0x904F
	OBJECT_TYPE                                                = 0x9112
	CLAMP                                                      = 0x2900
	MODELVIEW23_ARB                                            = 0x8737
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	SIGNED_IDENTITY_NV                                         = 0x853C
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	BINORMAL_ARRAY_EXT                                         = 0x843A
	INTERPOLATE_ARB                                            = 0x8575
	PROGRAM_RESIDENT_NV                                        = 0x8647
	SRGB_ALPHA                                                 = 0x8C42
	FACTOR_MAX_AMD                                             = 0x901D
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	SAMPLE_MASK_EXT                                            = 0x80A0
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	ADD_BLEND_IMG                                              = 0x8C09
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	RECT_NV                                                    = 0xF6
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	TEXTURE9                                                   = 0x84C9
	TEXTURE2_ARB                                               = 0x84C2
	PALETTE8_RGBA4_OES                                         = 0x8B98
	Z400_BINARY_AMD                                            = 0x8740
	SELECT                                                     = 0x1C02
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	OPERAND2_RGB_EXT                                           = 0x8592
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	RED_BITS                                                   = 0x0D52
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	LINES_ADJACENCY                                            = 0x000A
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	VERTEX_ARRAY                                               = 0x8074
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	OP_ADD_EXT                                                 = 0x8787
	TRANSFORM_FEEDBACK                                         = 0x8E22
	DOUBLE_VEC2                                                = 0x8FFC
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	INT_VEC3                                                   = 0x8B54
	NEXT_BUFFER_NV                                             = -2
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	EXT_texture3D                                              = 1
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	TEXTURE_ENV                                                = 0x2300
	MITER_REVERT_NV                                            = 0x90A7
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	ALPHA_TEST                                                 = 0x0BC0
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	SRGB_DECODE_ARB                                            = 0x8299
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	EXTENSIONS                                                 = 0x1F03
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	REG_28_ATI                                                 = 0x893D
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	RELATIVE_LINE_TO_NV                                        = 0x05
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	RGB16                                                      = 0x8054
	LIGHT6                                                     = 0x4006
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	DUAL_ALPHA4_SGIS                                           = 0x8110
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	REG_0_ATI                                                  = 0x8921
	PATH_STENCIL_REF_NV                                        = 0x90B8
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	VARIANT_VALUE_EXT                                          = 0x87E4
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	VERTICAL_LINE_TO_NV                                        = 0x08
	TEXTURE_BORDER_COLOR                                       = 0x1004
	TEXTURE_GREEN_SIZE                                         = 0x805D
	UNSIGNED_INT                                               = 0x1405
	PROXY_TEXTURE_3D                                           = 0x8070
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	STATIC_DRAW                                                = 0x88E4
	CON_27_ATI                                                 = 0x895C
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	MIN_LOD_WARNING_AMD                                        = 0x919C
	COEFF                                                      = 0x0A00
	MAX_EXT                                                    = 0x8008
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	RGB16F_EXT                                                 = 0x881B
	DRAW_BUFFER9_ATI                                           = 0x882E
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	FLOAT_RG32_NV                                              = 0x8887
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	LUMINANCE16UI_EXT                                          = 0x8D7A
	KEEP                                                       = 0x1E00
	COLOR_SUM_ARB                                              = 0x8458
	TEXTURE11                                                  = 0x84CB
	COMBINER_MAPPING_NV                                        = 0x8543
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	FRACTIONAL_ODD                                             = 0x8E7B
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	TEXTURE26                                                  = 0x84DA
	INFO_LOG_LENGTH                                            = 0x8B84
	IMAGE_3D_EXT                                               = 0x904E
	RGBA4_OES                                                  = 0x8056
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	READ_BUFFER_EXT                                            = 0x0C02
	TEXTURE_COMPONENTS                                         = 0x1003
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	COLOR_MATRIX                                               = 0x80B1
	INTENSITY8UI_EXT                                           = 0x8D7F
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	DOUBLE_MAT2                                                = 0x8F46
	RGBA8_SNORM                                                = 0x8F97
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	LUMINANCE                                                  = 0x1909
	DT_BIAS_NV                                                 = 0x8717
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	POINT_SIZE_MIN_ARB                                         = 0x8126
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	MAP1_BINORMAL_EXT                                          = 0x8446
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	DST_COLOR                                                  = 0x0306
	TEXTURE_BASE_LEVEL                                         = 0x813C
	MAX_LABEL_LENGTH                                           = 0x82E8
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	MATRIX10_ARB                                               = 0x88CA
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	OP_RECIP_EXT                                               = 0x8794
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	INCR_WRAP_OES                                              = 0x8507
	DISCARD_NV                                                 = 0x8530
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	TEXTURE_ENV_COLOR                                          = 0x2201
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	STENCIL_INDEX16                                            = 0x8D49
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	CON_26_ATI                                                 = 0x895B
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	VERTEX_ID_NV                                               = 0x8C7B
	INVERT                                                     = 0x150A
	DRAW_BUFFER11_NV                                           = 0x8830
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	TEXTURE20_ARB                                              = 0x84D4
	CON_11_ATI                                                 = 0x894C
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	AUX0                                                       = 0x0409
	VERTEX_ARRAY_POINTER                                       = 0x808E
	FRAGMENT_TEXTURE                                           = 0x829F
	DRAW_BUFFER1_ATI                                           = 0x8826
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	RGB8UI_EXT                                                 = 0x8D7D
	OR_REVERSE                                                 = 0x150B
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	AND                                                        = 0x1501
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	FULL_SUPPORT                                               = 0x82B7
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	SHADER_COMPILER                                            = 0x8DFA
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	VERTEX_SHADER_BIT                                          = 0x00000001
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	OP_CLAMP_EXT                                               = 0x878E
	SGIS_sharpen_texture                                       = 1
	BLEND_EQUATION                                             = 0x8009
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	EXT_copy_texture                                           = 1
	VIEWPORT_BIT                                               = 0x00000800
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	RGBA8I                                                     = 0x8D8E
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	DEPTH_BIAS                                                 = 0x0D1F
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	GL_4PASS_3_SGIS                                            = 0x80A7
	VERTEX_STREAM6_ATI                                         = 0x8772
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	ALPHA_SNORM                                                = 0x9010
	CURRENT_RASTER_INDEX                                       = 0x0B05
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	ACCUM                                                      = 0x0100
	TEXTURE2                                                   = 0x84C2
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	LUMINANCE12_ALPHA12                                        = 0x8047
	VERTEX_ARRAY_EXT                                           = 0x8074
	MATRIX_EXT                                                 = 0x87C0
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	FLOAT16_NV                                                 = 0x8FF8
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	ACTIVE_UNIFORMS                                            = 0x8B86
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	MODELVIEW4_ARB                                             = 0x8724
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	VARIABLE_E_NV                                              = 0x8527
	FLOAT_RGBA16_NV                                            = 0x888A
	INT_VEC4                                                   = 0x8B55
	DEBUG_TYPE_ERROR                                           = 0x824C
	SKIP_COMPONENTS1_NV                                        = -6
	DEBUG_SEVERITY_LOW                                         = 0x9148
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	SGIX_async_histogram                                       = 1
	RGBA4                                                      = 0x8056
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	PRIMITIVES_GENERATED                                       = 0x8C87
	DEPTH_COMPONENT                                            = 0x1902
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	RGBA4_DXT5_S3TC                                            = 0x83A5
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	NEGATE_BIT_ATI                                             = 0x00000004
	STENCIL_REF                                                = 0x0B97
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	COLOR_SAMPLES_NV                                           = 0x8E20
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	POLYGON_OFFSET_FILL                                        = 0x8037
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	CONSTANT_COLOR1_NV                                         = 0x852B
	INT64_VEC3_NV                                              = 0x8FEA
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	OBJECT_PLANE                                               = 0x2501
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	SHADER_IMAGE_STORE                                         = 0x82A5
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	MODELVIEW26_ARB                                            = 0x873A
	CON_3_ATI                                                  = 0x8944
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	MIRROR_CLAMP_ATI                                           = 0x8742
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	DRAW_BUFFER15_ATI                                          = 0x8834
	LUMINANCE8UI_EXT                                           = 0x8D80
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	MITER_TRUNCATE_NV                                          = 0x90A8
	SCISSOR_BIT                                                = 0x00080000
	CLIP_PLANE0                                                = 0x3000
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	INTENSITY16_SNORM                                          = 0x901B
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	COMBINER_SCALE_NV                                          = 0x8548
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	DOT3_RGBA_EXT                                              = 0x8741
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	INT_SAMPLER_CUBE                                           = 0x8DCC
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	DRAW_BUFFER3_ARB                                           = 0x8828
	DRAW_BUFFER7_ATI                                           = 0x882C
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	OPERAND0_ALPHA                                             = 0x8598
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	TEXTURE15                                                  = 0x84CF
	DRAW_BUFFER12_ATI                                          = 0x8831
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	SRGB_ALPHA_EXT                                             = 0x8C42
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	PACK_IMAGE_HEIGHT                                          = 0x806C
	DEPTH_COMPONENT16                                          = 0x81A5
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	DRAW_BUFFER0_ARB                                           = 0x8825
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	INT_IMAGE_2D                                               = 0x9058
	POINT_SMOOTH_HINT                                          = 0x0C51
	NORMAL_ARRAY_EXT                                           = 0x8075
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	INTENSITY32I_EXT                                           = 0x8D85
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	CLAMP_TO_BORDER                                            = 0x812D
	QUERY_RESULT                                               = 0x8866
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	SRC0_RGB                                                   = 0x8580
	MODELVIEW13_ARB                                            = 0x872D
	TEXTURE_COMPARE_MODE                                       = 0x884C
	BLUE_BIT_ATI                                               = 0x00000004
	INT_IMAGE_BUFFER                                           = 0x905C
	DEPTH_COMPONENT32                                          = 0x81A7
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	VERTEX_TEXTURE                                             = 0x829B
	MODELVIEW14_ARB                                            = 0x872E
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	VIEW_CLASS_32_BITS                                         = 0x82C8
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	SECONDARY_COLOR_NV                                         = 0x852D
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	GL_2PASS_0_SGIS                                            = 0x80A2
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	INDEX_TEST_REF_EXT                                         = 0x81B7
	WRAP_BORDER_SUN                                            = 0x81D4
	SCALAR_EXT                                                 = 0x87BE
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	PN_TRIANGLES_ATI                                           = 0x87F0
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	SCISSOR_TEST                                               = 0x0C11
	OP_NEGATE_EXT                                              = 0x8783
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	DYNAMIC_COPY_ARB                                           = 0x88EA
	TRANSLATE_3D_NV                                            = 0x9091
	PACK_SKIP_PIXELS                                           = 0x0D04
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	MAP_READ_BIT_EXT                                           = 0x0001
	RGB10_EXT                                                  = 0x8052
	STENCIL_INDEX16_EXT                                        = 0x8D49
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	LOGIC_OP                                                   = 0x0BF1
	Q                                                          = 0x2003
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	BLEND                                                      = 0x0BE2
	BLEND_DST_RGB                                              = 0x80C8
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	PROGRAM_BINDING_ARB                                        = 0x8677
	NEGATIVE_Z_EXT                                             = 0x87DB
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	FIELD_LOWER_NV                                             = 0x9023
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	COLOR_TABLE_FORMAT                                         = 0x80D8
	MAX_VIEWPORTS                                              = 0x825B
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	READ_WRITE_ARB                                             = 0x88BA
	STREAM_READ_ARB                                            = 0x88E1
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	PROXY_HISTOGRAM                                            = 0x8025
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	SGIX_ir_instrument1                                        = 1
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	CON_19_ATI                                                 = 0x8954
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	BACK                                                       = 0x0405
	SAMPLES_EXT                                                = 0x80A9
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	CON_21_ATI                                                 = 0x8956
	MODELVIEW17_ARB                                            = 0x8731
	LERP_ATI                                                   = 0x8969
	FRONT                                                      = 0x0404
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	PROXY_TEXTURE_2D                                           = 0x8064
	SINGLE_COLOR_EXT                                           = 0x81F9
	COLOR_ATTACHMENT1                                          = 0x8CE1
	ALPHA_TEST_REF                                             = 0x0BC2
	TEXTURE_GEN_Q                                              = 0x0C63
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	IS_ROW_MAJOR                                               = 0x9300
	RGB                                                        = 0x1907
	RGB12_EXT                                                  = 0x8053
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	STENCIL_INDEX4                                             = 0x8D47
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	CLOSE_PATH_NV                                              = 0x00
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	BLEND_COLOR_EXT                                            = 0x8005
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	SGIX_instruments                                           = 1
	LIGHT3                                                     = 0x4003
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	FRAMEZOOM_SGIX                                             = 0x818B
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	NORMAL_MAP_NV                                              = 0x8511
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	TESS_CONTROL_SHADER                                        = 0x8E88
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	TEXTURE_GEN_T                                              = 0x0C61
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	YCRCB_SGIX                                                 = 0x8318
	SOURCE3_RGB_NV                                             = 0x8583
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	MATRIX17_ARB                                               = 0x88D1
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	UNSIGNALED                                                 = 0x9118
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	BLEND_EQUATION_RGB                                         = 0x8009
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	VERSION_1_5                                                = 1
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	POINT_SIZE                                                 = 0x0B11
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	CULL_MODES_NV                                              = 0x86E0
	GL_2X_BIT_ATI                                              = 0x00000001
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	RG_INTEGER                                                 = 0x8228
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	LIST_MODE                                                  = 0x0B30
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	SIGNED_RGB8_NV                                             = 0x86FF
	MAX_TEXTURE_COORDS                                         = 0x8871
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	ALPHA_TEST_QCOM                                            = 0x0BC0
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	DOT3_RGBA_IMG                                              = 0x86AF
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	SGIX_shadow                                                = 1
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	RETAINED_APPLE                                             = 0x8A1B
	FRAMEBUFFER                                                = 0x8D40
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	REG_8_ATI                                                  = 0x8929
	RGB9_E5_EXT                                                = 0x8C3D
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	REPLACE_OLDEST_SUN                                         = 0x0003
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	QUADS                                                      = 0x0007
	DIFFUSE                                                    = 0x1201
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	OP_MIN_EXT                                                 = 0x878B
	CON_17_ATI                                                 = 0x8952
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	CLIP_DISTANCE1                                             = 0x3001
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	UNSIGNED_INVERT_NV                                         = 0x8537
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	BUFFER_SIZE_ARB                                            = 0x8764
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	SPECULAR                                                   = 0x1202
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	COMPRESSED_R11_EAC                                         = 0x9270
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	BOLD_BIT_NV                                                = 0x01
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	FLOAT_MAT3_ARB                                             = 0x8B5B
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	TRIANGLE_STRIP                                             = 0x0005
	SHININESS                                                  = 0x1601
	LIGHT2                                                     = 0x4002
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	LUMINANCE_ALPHA                                            = 0x190A
	GEOMETRY_TEXTURE                                           = 0x829E
	DSDT8_MAG8_NV                                              = 0x870A
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	COMBINE_ARB                                                = 0x8570
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	EXT_blend_color                                            = 1
	ZERO                                                       = 0
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	NORMAL_MAP                                                 = 0x8511
	POINT_SPRITE_NV                                            = 0x8861
	SLUMINANCE8                                                = 0x8C47
	GL_4PASS_0_EXT                                             = 0x80A4
	BLEND_DST_RGB_EXT                                          = 0x80C8
	TEXTURE28_ARB                                              = 0x84DC
	SRC1_ALPHA                                                 = 0x8589
	SAMPLER_1D_ARB                                             = 0x8B5D
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	CURRENT_FOG_COORDINATE                                     = 0x8453
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	TEXTURE7_ARB                                               = 0x84C7
	INCR_WRAP                                                  = 0x8507
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	ARRAY_BUFFER                                               = 0x8892
	CON_1_ATI                                                  = 0x8942
	MIN_EXT                                                    = 0x8007
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	AUX2                                                       = 0x040B
	FOG_START                                                  = 0x0B63
	GL_2PASS_1_EXT                                             = 0x80A3
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	SOURCE0_ALPHA                                              = 0x8588
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	REG_20_ATI                                                 = 0x8935
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	DOUBLE_MAT2x4                                              = 0x8F4A
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	MATRIX29_ARB                                               = 0x88DD
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	OUTPUT_FOG_EXT                                             = 0x87BD
	STATIC_READ_ARB                                            = 0x88E5
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	LUMINANCE8_EXT                                             = 0x8040
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	TEXTURE_HEIGHT                                             = 0x1001
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	DELETE_STATUS                                              = 0x8B80
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	MULTISAMPLE_SGIS                                           = 0x809D
	FOG_COORDINATE_EXT                                         = 0x8451
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	SLUMINANCE8_EXT                                            = 0x8C47
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	BEVEL_NV                                                   = 0x90A6
	VERTEX4_BIT_PGI                                            = 0x00000008
	NICEST                                                     = 0x1102
	VERTEX_SHADER                                              = 0x8B31
	TEXTURE_ENV_MODE                                           = 0x2200
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	DONT_CARE                                                  = 0x1100
	TEXTURE_CUBE_MAP                                           = 0x8513
	COMBINER4_NV                                               = 0x8554
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	CON_10_ATI                                                 = 0x894B
	LUMINANCE4_EXT                                             = 0x803F
	VIEW_CLASS_8_BITS                                          = 0x82CB
	FLOAT_RGB32_NV                                             = 0x8889
	SRC1_COLOR                                                 = 0x88F9
	TEXTURE_MIN_FILTER                                         = 0x2801
	GENERATE_MIPMAP                                            = 0x8191
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	DEPTH_STENCIL_OES                                          = 0x84F9
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	SGIX_depth_texture                                         = 1
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	GL_3_BYTES                                                 = 0x1408
	TEXTURE28                                                  = 0x84DC
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	CONVOLUTION_HEIGHT                                         = 0x8019
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	COMBINE_EXT                                                = 0x8570
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	HISTOGRAM_EXT                                              = 0x8024
	TEXTURE22                                                  = 0x84D6
	SOURCE1_ALPHA_ARB                                          = 0x8589
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	OP_MUL_EXT                                                 = 0x8786
	NUM_PASSES_ATI                                             = 0x8970
	PALETTE4_RGBA4_OES                                         = 0x8B93
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	LAYOUT_DEFAULT_INTEL                                       = 0
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	ALL_STATIC_DATA_IBM                                        = 103060
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	REFLECTION_MAP_EXT                                         = 0x8512
	TRANSPOSE_NV                                               = 0x862C
	RENDERBUFFER_EXT                                           = 0x8D41
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	UNSIGNED_INT_24_8                                          = 0x84FA
	OP_POWER_EXT                                               = 0x8793
	RGB8_SNORM                                                 = 0x8F96
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	DOT3_RGB_EXT                                               = 0x8740
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	MATRIX3_ARB                                                = 0x88C3
	EXT_rescale_normal                                         = 1
	RGBA16_EXT                                                 = 0x805B
	RED_MIN_CLAMP_INGR                                         = 0x8560
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	ADD_SIGNED_EXT                                             = 0x8574
	RENDERBUFFER_BINDING                                       = 0x8CA7
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	MULTIVIEW_EXT                                              = 0x90F1
	DEPTH_FUNC                                                 = 0x0B74
	COLOR_WRITEMASK                                            = 0x0C23
	MODELVIEW1_EXT                                             = 0x850A
	RGBA16UI_EXT                                               = 0x8D76
	TEXTURE_BINDING_2D                                         = 0x8069
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	DRAW_BUFFER5                                               = 0x882A
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	NORMAL_ARRAY_POINTER                                       = 0x808F
	NEAREST                                                    = 0x2600
	MINMAX                                                     = 0x802E
	SHADER_BINARY_VIV                                          = 0x8FC4
	LIGHT1                                                     = 0x4001
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	R8_EXT                                                     = 0x8229
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	DOT3_RGB_ARB                                               = 0x86AE
	UNSIGNED_NORMALIZED                                        = 0x8C17
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	CURRENT_QUERY                                              = 0x8865
	SAMPLER_1D                                                 = 0x8B5D
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	SUCCESS_NV                                                 = 0x902F
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	INDEX_BITS                                                 = 0x0D51
	TEXTURE_MAG_FILTER                                         = 0x2800
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	COMPRESSED_INTENSITY                                       = 0x84EC
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	FUNC_ADD_OES                                               = 0x8006
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	SAMPLER_CUBE                                               = 0x8B60
	VERSION_2_1                                                = 1
	DEPTH_SCALE                                                = 0x0D1E
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	SRC0_ALPHA                                                 = 0x8588
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	BLUE_INTEGER                                               = 0x8D96
	INT8_NV                                                    = 0x8FE0
	RGB12                                                      = 0x8053
	LUMINANCE12_EXT                                            = 0x8041
	ACTIVE_PROGRAM                                             = 0x8259
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	INDEX_MODE                                                 = 0x0C30
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	SOURCE2_ALPHA_ARB                                          = 0x858A
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	OBJECT_TYPE_ARB                                            = 0x8B4E
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	ONE_MINUS_DST_COLOR                                        = 0x0307
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	MULTISAMPLE_ARB                                            = 0x809D
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	RGBA8I_EXT                                                 = 0x8D8E
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	OP_EXP_BASE_2_EXT                                          = 0x8791
	IMAGE_2D_EXT                                               = 0x904D
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	CURRENT_RASTER_POSITION                                    = 0x0B07
	COPY_INVERTED                                              = 0x150C
	LUMINANCE4                                                 = 0x803F
	TEXTURE18_ARB                                              = 0x84D2
	RGBA32F_ARB                                                = 0x8814
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	YCBYCR8_422_NV                                             = 0x9031
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	MAX_DRAW_BUFFERS                                           = 0x8824
	PRESENT_TIME_NV                                            = 0x8E2A
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	SATURATE_BIT_ATI                                           = 0x00000040
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	DRAW_BUFFER8_ARB                                           = 0x882D
	REG_3_ATI                                                  = 0x8924
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	BOOL_VEC4                                                  = 0x8B59
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	FLOAT16_VEC2_NV                                            = 0x8FF9
	TYPE                                                       = 0x92FA
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	PROGRAM_STRING_ARB                                         = 0x8628
	SURFACE_STATE_NV                                           = 0x86EB
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	ALPHA_FLOAT32_ATI                                          = 0x8816
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	DSDT_MAG_VIB_NV                                            = 0x86F7
	HISTOGRAM_SINK_EXT                                         = 0x802D
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	TEXTURE5                                                   = 0x84C5
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	DISCRETE_AMD                                               = 0x9006
	EXT_texture                                                = 1
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	VERSION_1_4                                                = 1
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	SOURCE0_RGB_ARB                                            = 0x8580
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	DRAW_BUFFER11                                              = 0x8830
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	RASTERIZER_DISCARD                                         = 0x8C89
	STACK_UNDERFLOW                                            = 0x0504
	CLIP_DISTANCE7                                             = 0x3007
	RGBA_DXT5_S3TC                                             = 0x83A4
	DRAW_BUFFER9_ARB                                           = 0x882E
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	PIXEL_PACK_BUFFER                                          = 0x88EB
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	RGB_INTEGER_EXT                                            = 0x8D98
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	OP_SET_GE_EXT                                              = 0x878C
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	HISTOGRAM_FORMAT                                           = 0x8027
	DEPTH_COMPONENTS                                           = 0x8284
	MODELVIEW16_ARB                                            = 0x8730
	MATRIX6_ARB                                                = 0x88C6
	ALPHA16UI_EXT                                              = 0x8D78
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	TEXTURE_SHADOW                                             = 0x82A1
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	COLOR_CLEAR_VALUE                                          = 0x0C22
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	PERFMON_RESULT_AMD                                         = 0x8BC6
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	PACK_SKIP_IMAGES                                           = 0x806B
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	CULL_FACE                                                  = 0x0B44
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	STACK_OVERFLOW                                             = 0x0503
	NORMAL_MAP_OES                                             = 0x8511
	EMBOSS_MAP_NV                                              = 0x855F
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	REG_10_ATI                                                 = 0x892B
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	LUMINANCE16_EXT                                            = 0x8042
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	EXT_convolution                                            = 1
	NOOP                                                       = 0x1505
	OBJECT_LINEAR                                              = 0x2401
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	SHADER_OBJECT_EXT                                          = 0x8B48
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	PALETTE4_RGB8_OES                                          = 0x8B90
	RG_SNORM                                                   = 0x8F91
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	PROGRAM_PIPELINE                                           = 0x82E4
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	RGBA16I                                                    = 0x8D88
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	TEXTURE27                                                  = 0x84DB
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	DEPTH_STENCIL_NV                                           = 0x84F9
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	MATRIX28_ARB                                               = 0x88DC
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	CURRENT_VERTEX_EXT                                         = 0x87E2
	DOUBLE_MAT4x2                                              = 0x8F4D
	TEXTURE_BINDING_3D                                         = 0x806A
	READ_PIXELS_FORMAT                                         = 0x828D
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	ACTIVE_TEXTURE                                             = 0x84E0
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	ADD_SIGNED                                                 = 0x8574
	REG_16_ATI                                                 = 0x8931
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	SAMPLER_OBJECT_AMD                                         = 0x9155
	ORDER                                                      = 0x0A01
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	FLOAT_MAT3x4                                               = 0x8B68
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	SAMPLE_MASK                                                = 0x8E51
	SPOT_DIRECTION                                             = 0x1204
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	IMAGE_SCALE_Y_HP                                           = 0x8156
	COMPRESSED_RED                                             = 0x8225
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	ADD                                                        = 0x0104
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	ALPHA_TEST_FUNC                                            = 0x0BC1
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	ARRAY_BUFFER_ARB                                           = 0x8892
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	LINEAR                                                     = 0x2601
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	QUARTER_BIT_ATI                                            = 0x00000010
	RGB16I                                                     = 0x8D89
	VERSION_2_0                                                = 1
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	BACK_RIGHT                                                 = 0x0403
	SEPARABLE_2D_EXT                                           = 0x8012
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	IMAGE_SCALE_X_HP                                           = 0x8155
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	INT_VEC2                                                   = 0x8B53
	FLOAT_MAT3x2                                               = 0x8B67
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	RGBA12_EXT                                                 = 0x805A
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	TEXTURE6                                                   = 0x84C6
	CON_8_ATI                                                  = 0x8949
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	CURRENT_MATRIX_ARB                                         = 0x8641
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	UNSIGNED_INT16_NV                                          = 0x8FF0
	REPEAT                                                     = 0x2901
	RGB8_EXT                                                   = 0x8051
	BGR                                                        = 0x80E0
	REGISTER_COMBINERS_NV                                      = 0x8522
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	DRAW_BUFFER14_ARB                                          = 0x8833
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	TRUE                                                       = 1
	GL_3D_COLOR                                                = 0x0602
	PROGRAM_TARGET_NV                                          = 0x8646
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	RG32I                                                      = 0x823B
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	UPPER_LEFT                                                 = 0x8CA2
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	INDEX_BIT_PGI                                              = 0x00080000
	TEXTURE_GATHER                                             = 0x82A2
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	NORMAL_ARRAY_TYPE                                          = 0x807E
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	TEXTURE21                                                  = 0x84D5
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	INTERLACE_OML                                              = 0x8980
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	DEBUG_OUTPUT                                               = 0x92E0
	SGIX_shadow_ambient                                        = 1
	COLOR_ARRAY_POINTER                                        = 0x8090
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	GL_4PASS_2_EXT                                             = 0x80A6
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	SURFACE_MAPPED_NV                                          = 0x8700
	OP_MOV_EXT                                                 = 0x8799
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	CONVOLUTION_1D                                             = 0x8010
	CURRENT_FOG_COORD                                          = 0x8453
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	TEXTURE_BUFFER                                             = 0x8C2A
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	TEXTURE8                                                   = 0x84C8
	TEXTURE_RED_TYPE                                           = 0x8C10
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	WAIT_FAILED_APPLE                                          = 0x911D
	MODULATE                                                   = 0x2100
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	DOT3_RGBA_ARB                                              = 0x86AF
	MAP_TESSELLATION_NV                                        = 0x86C2
	MAP2_COLOR_4                                               = 0x0DB0
	GREEN                                                      = 0x1904
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	FLOAT16_VEC4_NV                                            = 0x8FFB
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	MODELVIEW_MATRIX                                           = 0x0BA6
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	IDENTITY_NV                                                = 0x862A
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	VERTEX_STREAM5_ATI                                         = 0x8771
	DRAW_BUFFER                                                = 0x0C01
	CURRENT_BINORMAL_EXT                                       = 0x843C
	STATE_RESTORE                                              = 0x8BDC
	CUBIC_EXT                                                  = 0x8334
	MIRRORED_REPEAT_OES                                        = 0x8370
	TEXTURE24_ARB                                              = 0x84D8
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	ZOOM_X                                                     = 0x0D16
	TEXTURE_RESIDENT                                           = 0x8067
	DUAL_ALPHA12_SGIS                                          = 0x8112
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	MODELVIEW10_ARB                                            = 0x872A
	CON_14_ATI                                                 = 0x894F
	SIGNALED_APPLE                                             = 0x9119
	AUTO_NORMAL                                                = 0x0D80
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	RGBA4_S3TC                                                 = 0x83A3
	FOG_COORDINATE_ARRAY                                       = 0x8457
	MODELVIEW21_ARB                                            = 0x8735
	OP_ROUND_EXT                                               = 0x8790
	TRANSLATE_2D_NV                                            = 0x9090
	LINE_RESET_TOKEN                                           = 0x0707
	DRAW_BUFFER1                                               = 0x8826
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	SGIX_pixel_tiles                                           = 1
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	LIST_INDEX                                                 = 0x0B33
	TEXTURE27_ARB                                              = 0x84DB
	FLOAT_RGB_NV                                               = 0x8882
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	VERTEX_SHADER_EXT                                          = 0x8780
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	REG_9_ATI                                                  = 0x892A
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	SKIP_DECODE_EXT                                            = 0x8A4A
	SHADER_TYPE                                                = 0x8B4F
	QUAD_STRIP                                                 = 0x0008
	CONVOLUTION_1D_EXT                                         = 0x8010
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	DRAW_BUFFER15_NV                                           = 0x8834
	ALPHA32UI_EXT                                              = 0x8D72
	LINE_TO_NV                                                 = 0x04
	COMBINER6_NV                                               = 0x8556
	VIBRANCE_SCALE_NV                                          = 0x8713
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	FLOAT_MAT2                                                 = 0x8B5A
	BGRA_INTEGER_EXT                                           = 0x8D9B
	EXT_histogram                                              = 1
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	CON_6_ATI                                                  = 0x8947
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	MODELVIEW25_ARB                                            = 0x8739
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	STATIC_READ                                                = 0x88E5
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	COLOR_MATERIAL                                             = 0x0B57
	READ_BUFFER_NV                                             = 0x0C02
	COMBINER_MUX_SUM_NV                                        = 0x8547
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	RESTART_SUN                                                = 0x0001
	MODELVIEW2_ARB                                             = 0x8722
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	SGIS_detail_texture                                        = 1
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	CURRENT_ATTRIB_NV                                          = 0x8626
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	R8_SNORM                                                   = 0x8F94
	TEXTURE_BINDING_1D                                         = 0x8068
	INT_IMAGE_1D                                               = 0x9057
	SYNC_FLAGS                                                 = 0x9115
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	REPLICATE_BORDER_HP                                        = 0x8153
	TEXTURE_LOD_BIAS                                           = 0x8501
	QUERY_RESULT_ARB                                           = 0x8866
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	OUTPUT_COLOR0_EXT                                          = 0x879B
	RGB10_A2UI                                                 = 0x906F
	NO_ERROR                                                   = 0
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	TEXTURE_SHADER_NV                                          = 0x86DE
	COLOR_ATTACHMENT12                                         = 0x8CEC
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	DOUBLE_MAT2_EXT                                            = 0x8F46
	DOUBLE_VEC4                                                = 0x8FFE
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	MAP1_INDEX                                                 = 0x0D91
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	ALPHA4                                                     = 0x803B
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	RGB32I                                                     = 0x8D83
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	FILL                                                       = 0x1B02
	DRAW_BUFFER0_NV                                            = 0x8825
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	REG_22_ATI                                                 = 0x8937
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	SOURCE1_ALPHA                                              = 0x8589
	DUDV_ATI                                                   = 0x8779
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	SGIS_texture4D                                             = 1
	OPERAND2_RGB_ARB                                           = 0x8592
	BUMP_TARGET_ATI                                            = 0x877C
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	RGBA32UI                                                   = 0x8D70
	LOW_INT                                                    = 0x8DF3
	LIST_PRIORITY_SGIX                                         = 0x8182
	NORMAL_MAP_ARB                                             = 0x8511
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	INVALID_VALUE                                              = 0x0501
	HALF_FLOAT                                                 = 0x140B
	Y_EXT                                                      = 0x87D6
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	MIRRORED_REPEAT_IBM                                        = 0x8370
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	RENDERBUFFER_OES                                           = 0x8D41
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	MODELVIEW5_ARB                                             = 0x8725
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	RGB10_A2_EXT                                               = 0x8059
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	QUERY_NO_WAIT                                              = 0x8E14
	REFLECTION_MAP                                             = 0x8512
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	COORD_REPLACE_NV                                           = 0x8862
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	SOURCE0_ALPHA_ARB                                          = 0x8588
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	BUFFER_MAPPED_OES                                          = 0x88BC
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	TEXTURE_WRAP_T                                             = 0x2803
	NUM_EXTENSIONS                                             = 0x821D
	MIRRORED_REPEAT                                            = 0x8370
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	COLOR_ATTACHMENT0                                          = 0x8CE0
	SKIP_COMPONENTS4_NV                                        = -3
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	MOVE_TO_NV                                                 = 0x02
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	AVERAGE_EXT                                                = 0x8335
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	SGIX_texture_multi_buffer                                  = 1
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	SHADER_CONSISTENT_NV                                       = 0x86DD
	RGB16F                                                     = 0x881B
	REG_12_ATI                                                 = 0x892D
	PATH_GEN_COEFF_NV                                          = 0x90B1
	POLYGON_MODE                                               = 0x0B40
	MAX_CLIP_DISTANCES                                         = 0x0D32
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	DRAW_BUFFER0                                               = 0x8825
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	COMPUTE_SHADER                                             = 0x91B9
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	OPERAND0_RGB_EXT                                           = 0x8590
	SLICE_ACCUM_SUN                                            = 0x85CC
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	VERSION_1_1                                                = 1
	SGIX_clipmap                                               = 1
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	PHONG_HINT_WIN                                             = 0x80EB
	SIGNED_ALPHA8_NV                                           = 0x8706
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	SYNC_FENCE_APPLE                                           = 0x9116
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	EDGE_FLAG_ARRAY                                            = 0x8079
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	SGIX_reference_plane                                       = 1
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	COLOR_RENDERABLE                                           = 0x8286
	INVARIANT_VALUE_EXT                                        = 0x87EA
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	LUMINANCE16I_EXT                                           = 0x8D8C
	POINT                                                      = 0x1B00
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	WRITE_DISCARD_NV                                           = 0x88BE
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	V3F                                                        = 0x2A21
	VARIABLE_D_NV                                              = 0x8526
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	POINT_SIZE_MIN                                             = 0x8126
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	MATRIX8_ARB                                                = 0x88C8
	EXT_subtexture                                             = 1
	FRONT_FACE                                                 = 0x0B46
	LIGHT7                                                     = 0x4007
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	CULL_VERTEX_EXT                                            = 0x81AA
	OP_FLOOR_EXT                                               = 0x878F
	DRAW_BUFFER10_NV                                           = 0x882F
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	SAMPLER_3D_ARB                                             = 0x8B5F
	CONSTANT_BORDER                                            = 0x8151
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	CON_15_ATI                                                 = 0x8950
	SAMPLER_BUFFER                                             = 0x8DC2
	SGIX_interlace                                             = 1
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	DRAW_BUFFER_EXT                                            = 0x0C01
	DRAW_BUFFER7_NV                                            = 0x882C
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	READ_PIXELS_TYPE                                           = 0x828E
	TEXTURE31_ARB                                              = 0x84DF
	EXPAND_NORMAL_NV                                           = 0x8538
	DRAW_BUFFER10_ARB                                          = 0x882F
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	MATRIX1_NV                                                 = 0x8631
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	MINMAX_SINK_EXT                                            = 0x8030
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	CON_2_ATI                                                  = 0x8943
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	RGB_INTEGER                                                = 0x8D98
	CLIP_DISTANCE2                                             = 0x3002
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	IUI_V2F_EXT                                                = 0x81AD
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	INTERPOLATE                                                = 0x8575
	Z_EXT                                                      = 0x87D7
	READ_WRITE                                                 = 0x88BA
	SRGB8_ALPHA8                                               = 0x8C43
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	INTENSITY16_EXT                                            = 0x804D
	YCBCR_422_APPLE                                            = 0x85B9
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	COLOR_INDEX12_EXT                                          = 0x80E6
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	TANGENT_ARRAY_EXT                                          = 0x8439
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	DOUBLE_MAT3x4                                              = 0x8F4C
	LOCATION_INDEX                                             = 0x930F
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	RGBA2_EXT                                                  = 0x8055
	SCREEN_COORDINATES_REND                                    = 0x8490
	MODELVIEW0_ARB                                             = 0x1700
	ELEMENT_ARRAY_ATI                                          = 0x8768
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	HALF_FLOAT_OES                                             = 0x8D61
	SMALL_CW_ARC_TO_NV                                         = 0x14
	ARRAY_SIZE                                                 = 0x92FB
	ALL_COMPLETED_NV                                           = 0x84F2
	COMBINE_RGB                                                = 0x8571
	CONSTANT                                                   = 0x8576
	TRIANGLE_MESH_SUN                                          = 0x8615
	NEGATIVE_W_EXT                                             = 0x87DC
	DRAW_BUFFER15_ARB                                          = 0x8834
	DRAW_BUFFER4_ATI                                           = 0x8829
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	FAILURE_NV                                                 = 0x9030
	SHADE_MODEL                                                = 0x0B54
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	SGIS_fog_function                                          = 1
	TABLE_TOO_LARGE                                            = 0x8031
	MAX_WIDTH                                                  = 0x827E
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	LUMINANCE16_ALPHA16                                        = 0x8048
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	LOGIC_OP_MODE                                              = 0x0BF0
	BGRA_EXT                                                   = 0x80E1
	INT8_VEC4_NV                                               = 0x8FE3
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	ALPHA16F_ARB                                               = 0x881C
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	RENDERBUFFER                                               = 0x8D41
	COMPUTE_SHADER_BIT                                         = 0x00000020
	FLOAT                                                      = 0x1406
	GL_2_BYTES                                                 = 0x1407
	BLEND_DST_ALPHA                                            = 0x80CA
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	MATRIX5_NV                                                 = 0x8635
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	COPY_READ_BUFFER                                           = 0x8F36
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	DRAW_BUFFER12_ARB                                          = 0x8831
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	RGB32I_EXT                                                 = 0x8D83
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	RGBA8_OES                                                  = 0x8058
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	MINMAX_FORMAT                                              = 0x802F
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	PROXY_COLOR_TABLE                                          = 0x80D3
	INTERLACE_READ_INGR                                        = 0x8568
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	CURRENT_INDEX                                              = 0x0B01
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	POINT_SIZE_RANGE                                           = 0x0B12
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	INTENSITY32UI_EXT                                          = 0x8D73
	BLUE_BITS                                                  = 0x0D54
	RGB_SCALE                                                  = 0x8573
	SOURCE3_ALPHA_NV                                           = 0x858B
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	POLYGON                                                    = 0x0009
	C4UB_V2F                                                   = 0x2A22
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	COLOR_ATTACHMENT13                                         = 0x8CED
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	MODULATE_COLOR_IMG                                         = 0x8C04
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	RELEASED_APPLE                                             = 0x8A19
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	VERSION_3_0                                                = 1
	STENCIL_VALUE_MASK                                         = 0x0B93
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	DRAW_BUFFER8                                               = 0x882D
	MAD_ATI                                                    = 0x8968
	INT_2_10_10_10_REV                                         = 0x8D9F
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	OR                                                         = 0x1507
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	INTENSITY8I_EXT                                            = 0x8D91
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	INT_IMAGE_2D_EXT                                           = 0x9058
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	PROJECTION_MATRIX                                          = 0x0BA7
	ALPHA12_EXT                                                = 0x803D
	TEXTURE26_ARB                                              = 0x84DA
	CON_20_ATI                                                 = 0x8955
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	INDEX_LOGIC_OP                                             = 0x0BF1
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	LUMINANCE12                                                = 0x8041
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	COMPRESSED_RGB                                             = 0x84ED
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	COLOR_TABLE_SCALE                                          = 0x80D6
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	ARC_TO_NV                                                  = 0xFE
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	INTERLACE_SGIX                                             = 0x8094
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	COMMAND_BARRIER_BIT                                        = 0x00000040
	BITMAP_TOKEN                                               = 0x0704
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	GL_3DC_X_AMD                                               = 0x87F9
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	COLOR_ATTACHMENT4                                          = 0x8CE4
	FRACTIONAL_EVEN                                            = 0x8E7C
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	RGB32F                                                     = 0x8815
	ENABLE_BIT                                                 = 0x00002000
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	PROGRAM_OUTPUT                                             = 0x92E4
	EXT_blend_subtract                                         = 1
	DUAL_ALPHA16_SGIS                                          = 0x8113
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	FRONT_AND_BACK                                             = 0x0408
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	RED_INTEGER_EXT                                            = 0x8D94
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	LOCAL_EXT                                                  = 0x87C4
	ONE_EXT                                                    = 0x87DE
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	BUFFER_VARIABLE                                            = 0x92E5
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	EYE_LINEAR                                                 = 0x2400
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	CLIP_PLANE3                                                = 0x3003
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	RGB565_OES                                                 = 0x8D62
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	SGI_color_matrix                                           = 1
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	CONVOLUTION_2D                                             = 0x8011
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	POINT_SPRITE_ARB                                           = 0x8861
	CON_24_ATI                                                 = 0x8959
	PATH_END_CAPS_NV                                           = 0x9076
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	SPOT_CUTOFF                                                = 0x1206
	INTENSITY16                                                = 0x804D
	RGBA16F                                                    = 0x881A
	DRAW_BUFFER2                                               = 0x8827
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	RGB5_EXT                                                   = 0x8050
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	RGB16UI                                                    = 0x8D77
	BLEND_EQUATION_OES                                         = 0x8009
	REDUCE_EXT                                                 = 0x8016
	RGB5_A1_OES                                                = 0x8057
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	REG_31_ATI                                                 = 0x8940
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	MAP2_INDEX                                                 = 0x0DB1
	OR_INVERTED                                                = 0x150D
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	TEXTURE14                                                  = 0x84CE
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	QUERY_COUNTER_BITS                                         = 0x8864
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	MAX_LIGHTS                                                 = 0x0D31
	C3F_V3F                                                    = 0x2A24
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	DRAW_BUFFER3_NV                                            = 0x8828
	BLOCK_INDEX                                                = 0x92FD
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	UNDEFINED_VERTEX                                           = 0x8260
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	BITMAP                                                     = 0x1A00
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	PROGRAM_SEPARABLE                                          = 0x8258
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	DEBUG_SOURCE_OTHER                                         = 0x824B
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	MIRRORED_REPEAT_ARB                                        = 0x8370
	TEXTURE16_ARB                                              = 0x84D0
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	SEPARATE_ATTRIBS                                           = 0x8C8D
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	DRAW_BUFFER6_NV                                            = 0x882B
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	COMPRESSED_RGBA                                            = 0x84EE
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	REG_2_ATI                                                  = 0x8923
	SAMPLER_3D                                                 = 0x8B5F
	COMPILE_STATUS                                             = 0x8B81
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	PATH_STROKE_MASK_NV                                        = 0x9084
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	POLYGON_STIPPLE                                            = 0x0B42
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	MATRIX20_ARB                                               = 0x88D4
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	INT_IMAGE_3D_EXT                                           = 0x9059
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	MAP_STENCIL                                                = 0x0D11
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	WAIT_FAILED                                                = 0x911D
	QUERY_OBJECT_AMD                                           = 0x9153
	COMPRESSED_RG11_EAC                                        = 0x9272
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	CURRENT_QUERY_ARB                                          = 0x8865
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	COLOR_SUM_EXT                                              = 0x8458
	SGI_color_table                                            = 1
	EQUIV                                                      = 0x1509
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	ALPHA32I_EXT                                               = 0x8D84
	SYNC_CONDITION_APPLE                                       = 0x9113
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	INDEX_WRITEMASK                                            = 0x0C21
	XOR                                                        = 0x1506
	T4F_V4F                                                    = 0x2A28
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	T2F_IUI_V3F_EXT                                            = 0x81B2
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	GREATER                                                    = 0x0204
	INVALID_ENUM                                               = 0x0500
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	BUFFER_MAP_POINTER                                         = 0x88BD
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	DST_ALPHA                                                  = 0x0304
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	REG_11_ATI                                                 = 0x892C
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	DRAW_BUFFER14_NV                                           = 0x8833
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	SGIS_multisample                                           = 1
	SHADER_IMAGE_LOAD                                          = 0x82A4
	PATCH_VERTICES                                             = 0x8E72
	MODELVIEW3_ARB                                             = 0x8723
	ONE                                                        = 1
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	VERSION                                                    = 0x1F02
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	REG_14_ATI                                                 = 0x892F
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	HIGH_FLOAT                                                 = 0x8DF2
	RED_EXT                                                    = 0x1903
	REFLECTION_MAP_OES                                         = 0x8512
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	LUMINANCE6_ALPHA2                                          = 0x8044
	T2F_V3F                                                    = 0x2A27
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	CONVEX_HULL_NV                                             = 0x908B
	TEXTURE_WIDTH                                              = 0x1000
	MODELVIEW0_EXT                                             = 0x1700
	MINMAX_FORMAT_EXT                                          = 0x802F
	CUBIC_HP                                                   = 0x815F
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	COLOR_MATERIAL_FACE                                        = 0x0B55
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	MATRIX18_ARB                                               = 0x88D2
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	CURRENT_NORMAL                                             = 0x0B02
	POLYGON_OFFSET_EXT                                         = 0x8037
	TEXTURE4                                                   = 0x84C4
	PROGRAM_STRING_NV                                          = 0x8628
	MODELVIEW9_ARB                                             = 0x8729
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	IMAGE_2D_RECT_EXT                                          = 0x904F
	RGB5_A1                                                    = 0x8057
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	OP_SET_LT_EXT                                              = 0x878D
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	FLOAT_MAT4_ARB                                             = 0x8B5C
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	OP_INDEX_EXT                                               = 0x8782
	GL_3DC_XY_AMD                                              = 0x87FA
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	FILE_NAME_NV                                               = 0x9074
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	COUNT_UP_NV                                                = 0x9088
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	FRAMEBUFFER_OES                                            = 0x8D40
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	INT_IMAGE_3D                                               = 0x9059
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	CND0_ATI                                                   = 0x896B
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	CLEAR_BUFFER                                               = 0x82B4
	STORAGE_SHARED_APPLE                                       = 0x85BF
	CONSTANT_COLOR                                             = 0x8001
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	INT_SAMPLER_2D                                             = 0x8DCA
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	ARRAY_STRIDE                                               = 0x92FE
	CLIP_DISTANCE0                                             = 0x3000
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	RG8_SNORM                                                  = 0x8F95
	SRC_ALPHA_SATURATE                                         = 0x0308
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	VARIABLE_A_NV                                              = 0x8523
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	HISTOGRAM_SINK                                             = 0x802D
	VIEW_CLASS_128_BITS                                        = 0x82C4
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	TRIANGLE_FAN                                               = 0x0006
	RGBA8_EXT                                                  = 0x8058
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	FLOAT_RGBA_NV                                              = 0x8883
	VALIDATE_STATUS                                            = 0x8B83
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	DRAW_BUFFER2_NV                                            = 0x8827
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	DITHER                                                     = 0x0BD0
	TEXTURE17                                                  = 0x84D1
	DEPTH                                                      = 0x1801
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	SAMPLES_PASSED_ARB                                         = 0x8914
	COMP_BIT_ATI                                               = 0x00000002
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	RGB4                                                       = 0x804F
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	EMBOSS_LIGHT_NV                                            = 0x855D
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	LINE_SMOOTH                                                = 0x0B20
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	PACK_CMYK_HINT_EXT                                         = 0x800E
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	CONSTANT_ATTENUATION                                       = 0x1207
	BLEND_COLOR                                                = 0x8005
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	TEXTURE20                                                  = 0x84D4
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	CURRENT_MATRIX_NV                                          = 0x8641
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	MODELVIEW30_ARB                                            = 0x873E
	MATRIX14_ARB                                               = 0x88CE
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	INT16_NV                                                   = 0x8FE4
	TEXTURE_MATRIX                                             = 0x0BA8
	TEXTURE_WRAP_R                                             = 0x8072
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	SGIS_texture_select                                        = 1
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	SOURCE2_ALPHA                                              = 0x858A
	OP_FRAC_EXT                                                = 0x8789
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	CULL_FACE_MODE                                             = 0x0B45
	RED_BIAS                                                   = 0x0D15
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	COMPRESSED_ALPHA                                           = 0x84E9
	VERTEX_STREAM7_ATI                                         = 0x8773
	INT_VEC3_ARB                                               = 0x8B54
	TEXTURE_1D                                                 = 0x0DE0
	R16F                                                       = 0x822D
	SOURCE0_ALPHA_EXT                                          = 0x8588
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	SGIX_async                                                 = 1
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	SCISSOR_BOX                                                = 0x0C10
	RGBA16F_ARB                                                = 0x881A
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	MAX_VERTEX_STREAMS                                         = 0x8E71
	LIST_BASE                                                  = 0x0B32
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	SAMPLE_POSITION_NV                                         = 0x8E50
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	VERTEX_STREAM2_ATI                                         = 0x876E
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	MULTISAMPLE_BIT                                            = 0x20000000
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	SIGNED_RGBA_NV                                             = 0x86FB
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	T2F_N3F_V3F                                                = 0x2A2B
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	PREVIOUS_EXT                                               = 0x8578
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	DOUBLE_MAT3_EXT                                            = 0x8F47
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	EXT_shared_texture_palette                                 = 1
	TEXTURE                                                    = 0x1702
	R                                                          = 0x2002
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	MAX_TEXTURE_UNITS                                          = 0x84E2
	DS_BIAS_NV                                                 = 0x8716
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	TEXTURE_GEN_MODE                                           = 0x2500
	GL_2PASS_0_EXT                                             = 0x80A2
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	DRAW_BUFFER13                                              = 0x8832
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	COLOR_LOGIC_OP                                             = 0x0BF2
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	MATRIX9_ARB                                                = 0x88C9
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	BLEND_SRC_ALPHA                                            = 0x80CB
	TEXTURE_4D_SGIS                                            = 0x8134
	TEXTURE_VIEW                                               = 0x82B5
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	LOW_FLOAT                                                  = 0x8DF0
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	TEXTURE_DEPTH                                              = 0x8071
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	STREAM_COPY                                                = 0x88E2
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	SGIX_texture_scale_bias                                    = 1
	PRIMARY_COLOR                                              = 0x8577
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	DRAW_BUFFER1_ARB                                           = 0x8826
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	DYNAMIC_READ                                               = 0x88E9
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	STENCIL_ATTACHMENT                                         = 0x8D20
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	SGIX_pixel_texture                                         = 1
	MATRIX_PALETTE_OES                                         = 0x8840
	BOUNDING_BOX_NV                                            = 0x908D
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	LAYOUT_LINEAR_INTEL                                        = 1
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	OPERAND2_ALPHA_ARB                                         = 0x859A
	ALPHA16I_EXT                                               = 0x8D8A
	ALPHA8_SNORM                                               = 0x9014
	TRIANGULAR_NV                                              = 0x90A5
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	GL_4PASS_1_SGIS                                            = 0x80A5
	VERTEX_STREAM1_ATI                                         = 0x876D
	OP_RECIP_SQRT_EXT                                          = 0x8795
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	VERSION_1_3                                                = 1
	S                                                          = 0x2000
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	REG_30_ATI                                                 = 0x893F
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	MULT                                                       = 0x0103
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	RGB16_SNORM                                                = 0x8F9A
	SAMPLES                                                    = 0x80A9
	REG_7_ATI                                                  = 0x8928
	FIELDS_NV                                                  = 0x8E27
	HALF_FLOAT_NV                                              = 0x140B
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	BLEND_SRC_RGB                                              = 0x80C9
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	MULTISAMPLE_3DFX                                           = 0x86B2
	BLEND_EQUATION_ALPHA                                       = 0x883D
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	FLOAT_VEC4_ARB                                             = 0x8B52
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	HI_BIAS_NV                                                 = 0x8714
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	FASTEST                                                    = 0x1101
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	SGIX_fragment_lighting                                     = 1
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	MODULATE_ADD_ATI                                           = 0x8744
	DRAW_BUFFER13_NV                                           = 0x8832
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	OPERAND0_RGB                                               = 0x8590
	OPERAND3_RGB_NV                                            = 0x8593
	ALPHA16_SNORM                                              = 0x9018
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	INTENSITY_SNORM                                            = 0x9013
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	MAJOR_VERSION                                              = 0x821B
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	FOG                                                        = 0x0B60
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	MAX_VARYING_FLOATS                                         = 0x8B4B
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	TEXTURE_DEPTH_EXT                                          = 0x8071
	GL_4PASS_3_EXT                                             = 0x80A7
	INCR_WRAP_EXT                                              = 0x8507
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	DUAL_ALPHA8_SGIS                                           = 0x8111
	FIELD_UPPER_NV                                             = 0x9022
	SGIX_texture_add_env                                       = 1
	TEXTURE23                                                  = 0x84D7
	COMPRESSED_RGB_ARB                                         = 0x84ED
	BGRA_INTEGER                                               = 0x8D9B
	LUMINANCE8_ALPHA8                                          = 0x8045
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	DRAW_BUFFER9                                               = 0x882E
	STREAM_COPY_ARB                                            = 0x88E2
	PRIMITIVE_RESTART                                          = 0x8F9D
	PASS_THROUGH_TOKEN                                         = 0x0700
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	RGB_FLOAT32_APPLE                                          = 0x8815
	RGBA16F_EXT                                                = 0x881A
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	TEXTURE_GEN_S                                              = 0x0C60
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	SOURCE0_RGB_EXT                                            = 0x8580
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	FLOAT_RGB16_NV                                             = 0x8888
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	FLAT                                                       = 0x1D00
	BUFFER_MAPPED_ARB                                          = 0x88BC
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	SGIX_icc_texture                                           = 1
	MAP_WRITE_BIT_EXT                                          = 0x0002
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	PROGRAM_LENGTH_ARB                                         = 0x8627
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	COUNTER_RANGE_AMD                                          = 0x8BC1
	VIEW_CLASS_64_BITS                                         = 0x82C6
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	MAX_SAMPLES                                                = 0x8D57
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	BUFFER_SIZE                                                = 0x8764
	ALPHA_FLOAT16_ATI                                          = 0x881C
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	DEBUG_TYPE_OTHER                                           = 0x8251
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	BUFFER_USAGE_ARB                                           = 0x8765
	UTF8_NV                                                    = 0x909A
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	REG_19_ATI                                                 = 0x8934
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	BUFFER_MAP_LENGTH                                          = 0x9120
	COLOR_ARRAY_SIZE                                           = 0x8081
	INT                                                        = 0x1404
	BUFFER_ACCESS_ARB                                          = 0x88BB
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	ALPHA_BITS                                                 = 0x0D55
	RGBA12                                                     = 0x805A
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	SIGNED_INTENSITY_NV                                        = 0x8707
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	TEXTURE_BIT                                                = 0x00040000
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	CURRENT_RASTER_COLOR                                       = 0x0B04
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	OPERAND3_ALPHA_NV                                          = 0x859B
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	SGIX_calligraphic_fragment                                 = 1
	DOMAIN                                                     = 0x0A02
	EDGE_FLAG                                                  = 0x0B43
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	SIGNED_RGB_NV                                              = 0x86FE
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	QUAD_ALPHA4_SGIS                                           = 0x811E
	COLOR_ATTACHMENT10                                         = 0x8CEA
	STENCIL_INDEX8_OES                                         = 0x8D48
	R32I                                                       = 0x8235
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	INVERTED_SCREEN_W_REND                                     = 0x8491
	TEXTURE1                                                   = 0x84C1
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	TEXTURE9_ARB                                               = 0x84C9
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	COLOR                                                      = 0x1800
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	TEXTURE1_ARB                                               = 0x84C1
	REG_25_ATI                                                 = 0x893A
	MOV_ATI                                                    = 0x8961
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	POLYGON_SMOOTH                                             = 0x0B41
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	MATRIX24_ARB                                               = 0x88D8
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	COLOR_SUM                                                  = 0x8458
	SYNC_CONDITION                                             = 0x9113
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	PROGRAM                                                    = 0x82E2
	DU8DV8_ATI                                                 = 0x877A
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	SYNC_FENCE                                                 = 0x9116
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	SPRITE_SGIX                                                = 0x8148
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	RGBA8UI_EXT                                                = 0x8D7C
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	HILO8_NV                                                   = 0x885E
	DEPTH_BOUNDS_EXT                                           = 0x8891
	CON_4_ATI                                                  = 0x8945
	COLOR_ATTACHMENT6                                          = 0x8CE6
	CLEAR                                                      = 0x1500
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	CLAMP_TO_EDGE                                              = 0x812F
	BOOL                                                       = 0x8B56
	COLOR_ATTACHMENT15                                         = 0x8CEF
	RGBA_INTEGER                                               = 0x8D99
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	INTENSITY4_EXT                                             = 0x804A
	GL_1PASS_EXT                                               = 0x80A1
	PATH_DASH_OFFSET_NV                                        = 0x907E
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	ACCUM_RED_BITS                                             = 0x0D58
	MAP2_VERTEX_3                                              = 0x0DB7
	SYNC_CL_EVENT_ARB                                          = 0x8240
	DEBUG_SOURCE_API                                           = 0x8246
	ATTACHED_SHADERS                                           = 0x8B85
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	STENCIL_RENDERABLE                                         = 0x8288
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	PHONG_WIN                                                  = 0x80EA
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	SQUARE_NV                                                  = 0x90A3
	ALPHA                                                      = 0x1906
	MATRIX11_ARB                                               = 0x88CB
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	MAP2_TANGENT_EXT                                           = 0x8445
	OP_LOG_BASE_2_EXT                                          = 0x8792
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	IMAGE_1D_EXT                                               = 0x904C
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	POINT_SIZE_MAX_ARB                                         = 0x8127
	OP_DOT4_EXT                                                = 0x8785
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	TEXTURE_3D_EXT                                             = 0x806F
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	RGB565                                                     = 0x8D62
	RESCALE_NORMAL                                             = 0x803A
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	MEDIUM_INT                                                 = 0x8DF4
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	MOVE_TO_RESETS_NV                                          = 0x90B5
	FOG_DENSITY                                                = 0x0B62
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	ATTENUATION_EXT                                            = 0x834D
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	PATH_FORMAT_SVG_NV                                         = 0x9070
	NORMAL_ARRAY                                               = 0x8075
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	COMPUTE_TEXTURE                                            = 0x82A0
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	DRAW_BUFFER8_NV                                            = 0x882D
	POLYGON_TOKEN                                              = 0x0703
	ABGR_EXT                                                   = 0x8000
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	VIEW_CLASS_16_BITS                                         = 0x82CA
	COMBINER5_NV                                               = 0x8555
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	SAMPLE_MASK_VALUE                                          = 0x8E52
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	MAGNITUDE_SCALE_NV                                         = 0x8712
	FULL_RANGE_EXT                                             = 0x87E1
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	TEXTURE5_ARB                                               = 0x84C5
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	SAMPLER_3D_OES                                             = 0x8B5F
	T2F_C4UB_V3F                                               = 0x2A29
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	UNSIGNALED_APPLE                                           = 0x9118
	SGIX_framezoom                                             = 1
	ALWAYS                                                     = 0x0207
	EYE_PLANE                                                  = 0x2502
	AVERAGE_HP                                                 = 0x8160
	RG16UI                                                     = 0x823A
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	SOURCE2_RGB_EXT                                            = 0x8582
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	SIGNED_HILO8_NV                                            = 0x885F
	POLYGON_BIT                                                = 0x00000008
	T                                                          = 0x2001
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	HISTOGRAM_RED_SIZE                                         = 0x8028
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	SGX_BINARY_IMG                                             = 0x8C0A
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	DOUBLE_MAT4                                                = 0x8F48
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	COORD_REPLACE                                              = 0x8862
	IMAGE_3D                                                   = 0x904E
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	NONE                                                       = 0
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	RGBA32F                                                    = 0x8814
	POINT_BIT                                                  = 0x00000002
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	MAGNITUDE_BIAS_NV                                          = 0x8718
	OP_MAX_EXT                                                 = 0x878A
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	DECR_WRAP                                                  = 0x8508
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	DEPTH32F_STENCIL8                                          = 0x8CAD
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	COMPILE                                                    = 0x1300
	DISPLAY_LIST                                               = 0x82E7
	DSDT_NV                                                    = 0x86F5
	OUTPUT_COLOR1_EXT                                          = 0x879C
	RGB_FLOAT32_ATI                                            = 0x8815
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	MAX_SAMPLES_IMG                                            = 0x9135
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	ALPHA8UI_EXT                                               = 0x8D7E
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	CAVEAT_SUPPORT                                             = 0x82B8
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	CLIP_PLANE2                                                = 0x3002
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	TEXTURE_COORD_ARRAY                                        = 0x8078
	MINMAX_SINK                                                = 0x8030
	TEXTURE21_ARB                                              = 0x84D5
	RED_BIT_ATI                                                = 0x00000001
	LOCATION                                                   = 0x930E
	OPERAND0_ALPHA_ARB                                         = 0x8598
	TRANSLATE_Y_NV                                             = 0x908F
	QUERY_BUFFER_AMD                                           = 0x9192
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	NOR                                                        = 0x1508
	MAX_HEIGHT                                                 = 0x827F
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	LINES                                                      = 0x0001
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	STATIC_COPY                                                = 0x88E6
	REG_26_ATI                                                 = 0x893B
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	CURRENT_BIT                                                = 0x00000001
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	DEPTH_COMPONENT32F                                         = 0x8CAC
	DOUBLE_VEC3                                                = 0x8FFD
	COLOR4_BIT_PGI                                             = 0x00020000
	FRONT_LEFT                                                 = 0x0400
	GL_8X_BIT_ATI                                              = 0x00000004
	INVALID_INDEX                                              = 0xFFFFFFFF
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	PACK_SKIP_ROWS                                             = 0x0D03
	FUNC_ADD                                                   = 0x8006
	RGB4_EXT                                                   = 0x804F
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	DEBUG_TYPE_MARKER                                          = 0x8268
	DEPTH_STENCIL                                              = 0x84F9
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	DRAW_BUFFER12_NV                                           = 0x8831
	INT16_VEC2_NV                                              = 0x8FE5
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	TEXTURE_PRIORITY                                           = 0x8066
	YCRCB_444_SGIX                                             = 0x81BC
	OPERAND2_RGB                                               = 0x8592
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	TEXTURE_3D_OES                                             = 0x806F
	COMPRESSED_RG                                              = 0x8226
	INDEX_ARRAY_TYPE                                           = 0x8085
	FRAGMENT_COLOR_EXT                                         = 0x834C
	TEXTURE14_ARB                                              = 0x84CE
	MODELVIEW28_ARB                                            = 0x873C
	STENCIL_BACK_REF                                           = 0x8CA3
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	R16I                                                       = 0x8233
	RG16I                                                      = 0x8239
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	DRAW_BUFFER10_ATI                                          = 0x882F
	MATRIX23_ARB                                               = 0x88D7
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	EYE_POINT_SGIS                                             = 0x81F4
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	HALF_APPLE                                                 = 0x140B
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	NEGATIVE_ONE_EXT                                           = 0x87DF
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	MULTISAMPLE_EXT                                            = 0x809D
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	R32F                                                       = 0x822E
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	MODELVIEW18_ARB                                            = 0x8732
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	COLOR_INDEX                                                = 0x1900
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	MAX_PATCH_VERTICES                                         = 0x8E7D
	PATH_DASH_CAPS_NV                                          = 0x907B
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	SRGB_WRITE                                                 = 0x8298
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	TEXTURE25_ARB                                              = 0x84D9
	SOURCE1_RGB_ARB                                            = 0x8581
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	CULL_VERTEX_IBM                                            = 103050
	SGIS_texture_border_clamp                                  = 1
	PIXEL_MODE_BIT                                             = 0x00000020
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	INDEX                                                      = 0x8222
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	VIDEO_BUFFER_NV                                            = 0x9020
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	ALPHA12                                                    = 0x803D
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	ISOLINES                                                   = 0x8E7A
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	SGIS_texture_edge_clamp                                    = 1
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	TEXTURE_2D                                                 = 0x0DE1
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	REPLICATE_BORDER                                           = 0x8153
	MATRIX22_ARB                                               = 0x88D6
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	GENERATE_MIPMAP_HINT                                       = 0x8192
	SINGLE_COLOR                                               = 0x81F9
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	GL_4X_BIT_ATI                                              = 0x00000002
	BOOL_VEC2                                                  = 0x8B57
	PROGRAM_INPUT                                              = 0x92E3
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	MATRIX5_ARB                                                = 0x88C5
	CLAMP_VERTEX_COLOR                                         = 0x891A
	FLOAT_MAT4x3                                               = 0x8B6A
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	STANDARD_FONT_NAME_NV                                      = 0x9072
	PASS_THROUGH_NV                                            = 0x86E6
	R11F_G11F_B10F                                             = 0x8C3A
	STENCIL_INDEX1_OES                                         = 0x8D46
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	PROJECTION                                                 = 0x1701
	POINT_SIZE_MAX_EXT                                         = 0x8127
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	EVAL_2D_NV                                                 = 0x86C0
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	YCRCB_422_SGIX                                             = 0x81BB
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	BGR_INTEGER_EXT                                            = 0x8D9A
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	SIGNALED                                                   = 0x9119
	EXT_blend_minmax                                           = 1
	EMBOSS_CONSTANT_NV                                         = 0x855E
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	INT_IMAGE_1D_EXT                                           = 0x9057
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	DECODE_EXT                                                 = 0x8A49
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	DEPTH_RANGE                                                = 0x0B70
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	TEXTURE7                                                   = 0x84C7
	TEXTURE13                                                  = 0x84CD
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	RGBA8UI                                                    = 0x8D7C
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	FOG_COORDINATE                                             = 0x8451
	COLOR_MATRIX_SGI                                           = 0x80B1
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	TEXTURE_LIGHT_EXT                                          = 0x8350
	PRESERVE_ATI                                               = 0x8762
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	RG16_SNORM                                                 = 0x8F99
	GEQUAL                                                     = 0x0206
	NORMAL_BIT_PGI                                             = 0x08000000
	COLOR_ARRAY_TYPE                                           = 0x8082
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	HALF_FLOAT_ARB                                             = 0x140B
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	COLOR_COMPONENTS                                           = 0x8283
	W_EXT                                                      = 0x87D8
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	FENCE_APPLE                                                = 0x8A0B
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	TEXTURE3                                                   = 0x84C3
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	PIXEL_COUNT_NV                                             = 0x8866
	UNSIGNED_INT8_NV                                           = 0x8FEC
	SHORT                                                      = 0x1402
	CLIP_DISTANCE5                                             = 0x3005
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	PRIMARY_COLOR_NV                                           = 0x852C
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	DRAW_BUFFER4_ARB                                           = 0x8829
	DEPTH_TEXTURE_MODE                                         = 0x884B
	UNIFORM_BUFFER                                             = 0x8A11
	FLOAT_MAT2_ARB                                             = 0x8B5A
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	STENCIL_COMPONENTS                                         = 0x8285
	MODELVIEW29_ARB                                            = 0x873D
	COUNTER_TYPE_AMD                                           = 0x8BC0
	INTENSITY16I_EXT                                           = 0x8D8B
	RGB4_S3TC                                                  = 0x83A1
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	STENCIL_BACK_FAIL                                          = 0x8801
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	PATH_FILL_MODE_NV                                          = 0x9080
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	QUERY_NO_WAIT_NV                                           = 0x8E14
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	SRC_COLOR                                                  = 0x0300
	GREEN_BITS                                                 = 0x0D53
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	RED                                                        = 0x1903
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
)

type Context struct {
	context                   *C.gl20Context
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
	ColorTable                func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer)
	ColorTableParameterfv     func(target, pname uint32, params *float32)
	ColorTableParameteriv     func(target, pname uint32, params *int32)
	ColorSubTable             func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer)
	CopyPixels                func(x, y int32, width, height int32, Type uint32)
	CullFace                  func(mode uint32)
	ConvolutionFilter1D       func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionFilter2D       func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionParameterf     func(target, pname uint32, params float32)
	ConvolutionParameteri     func(target, pname uint32, params int32)
	CopyColorTable            func(target, internalformat uint32, x, y int32, width int32)
	CopyColorSubTable         func(target uint32, start int32, x, y int32, width int32)
	CopyConvolutionFilter1D   func(target, internalformat uint32, x, y int32, width int32)
	CopyConvolutionFilter2D   func(target, internalformat uint32, x, y int32, width, height int32)
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
	GetColorTable             func(target, format, Type uint32, table unsafe.Pointer)
	GetColorTableParameterfv  func(target, pname uint32, params *float32)
	GetColorTableParameteriv  func(target, pname uint32, params *int32)
	GetConvolutionFilter      func(target, format, Type uint32, image unsafe.Pointer)
	GetConvolutionParameterfv func(target, pname uint32, params *float32)
	GetConvolutionParameteriv func(target, pname uint32, params *int32)
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
}

func New() *Context {
	glc := new(Context)
	glc.context = C.gl20NewContext()

	glc.Accum = func(op uint32, value float32) {
		C.gl20Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		C.gl20AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		C.gl20Begin(glc.context, C.GLenum(mode))
	}

	glc.End = func() {
		C.gl20End(glc.context)
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8) {
		C.gl20Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		C.gl20BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		C.gl20CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		C.gl20CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		C.gl20Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		C.gl20ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		C.gl20ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		C.gl20ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		C.gl20ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		C.gl20ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation *float64) {
		C.gl20ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.Color3b = func(red, green, blue int8) {
		C.gl20Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		C.gl20Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		C.gl20Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		C.gl20Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		C.gl20Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		C.gl20Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		C.gl20Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		C.gl20Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		C.gl20Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		C.gl20Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		C.gl20Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		C.gl20Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		C.gl20Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		C.gl20Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		C.gl20Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		C.gl20Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v *int8) {
		C.gl20Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color3dv = func(v *float64) {
		C.gl20Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color3fv = func(v *float32) {
		C.gl20Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color3iv = func(v *int32) {
		C.gl20Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color3sv = func(v *int16) {
		C.gl20Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color3ubv = func(v *uint8) {
		C.gl20Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color3uiv = func(v *uint32) {
		C.gl20Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color3usv = func(v *uint16) {
		C.gl20Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.Color4bv = func(v *int8) {
		C.gl20Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color4dv = func(v *float64) {
		C.gl20Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color4fv = func(v *float32) {
		C.gl20Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color4iv = func(v *int32) {
		C.gl20Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color4sv = func(v *int16) {
		C.gl20Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color4ubv = func(v *uint8) {
		C.gl20Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color4uiv = func(v *uint32) {
		C.gl20Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color4usv = func(v *uint16) {
		C.gl20Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		C.gl20ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		C.gl20ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl20ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params *float32) {
		C.gl20ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params *int32) {
		C.gl20ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		C.gl20ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		C.gl20CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		C.gl20CullFace(glc.context, C.GLenum(mode))
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl20ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl20ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		C.gl20ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		C.gl20ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl20CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		C.gl20CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl20CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		C.gl20CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		C.gl20DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		C.gl20DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthMask = func(flag bool) {
		C.gl20DepthMask(glc.context, boolToGL(flag))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		C.gl20DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		C.gl20Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		C.gl20Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		C.gl20DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl20DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.EdgeFlag = func(flag bool) {
		C.gl20EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag *bool) {
		C.gl20EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(flag)))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		C.gl20EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EvalCoord1d = func(u float64) {
		C.gl20EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		C.gl20EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		C.gl20EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		C.gl20EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u *float64) {
		C.gl20EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord1fv = func(u *float32) {
		C.gl20EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2dv = func(u *float64) {
		C.gl20EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2fv = func(u *float32) {
		C.gl20EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		C.gl20EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		C.gl20EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		C.gl20EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		C.gl20EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer *float32) {
		C.gl20FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(buffer)))
	}

	glc.Finish = func() {
		C.gl20Finish(glc.context)
	}

	glc.Flush = func() {
		C.gl20Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		C.gl20Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		C.gl20Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params *float32) {
		C.gl20Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Fogiv = func(pname uint32, params *int32) {
		C.gl20Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.FrontFace = func(mode uint32) {
		C.gl20FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		C.gl20Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		return uint32(C.gl20GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params *bool) {
		C.gl20GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(params)))
	}

	glc.GetDoublev = func(pname uint32, params *float64) {
		C.gl20GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetFloatv = func(pname uint32, params *float32) {
		C.gl20GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetIntegerv = func(pname uint32, params *int32) {
		C.gl20GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetClipPlane = func(plane uint32, equation *float64) {
		C.gl20GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl20GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params *float32) {
		C.gl20GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetLightiv = func(light, pname uint32, params *int32) {
		C.gl20GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetMapdv = func(target, query uint32, v *float64) {
		C.gl20GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.GetMapfv = func(target, query uint32, v *float32) {
		C.gl20GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.GetMapiv = func(target, query uint32, v *int32) {
		C.gl20GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.GetMaterialfv = func(face, pname uint32, params *float32) {
		C.gl20GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetMaterialiv = func(face, pname uint32, params *int32) {
		C.gl20GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetPixelMapfv = func(Map uint32, values *float32) {
		C.gl20GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapuiv = func(Map uint32, values *uint32) {
		C.gl20GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapusv = func(Map uint32, values *uint16) {
		C.gl20GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.GetPolygonStipple = func(pattern *uint8) {
		C.gl20GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(pattern)))
	}

	glc.GetString = func(name uint32) string {
		cstr := C.gl20GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params *float32) {
		C.gl20GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexEnviv = func(target, pname uint32, params *int32) {
		C.gl20GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexGendv = func(coord, pname uint32, params *float64) {
		C.gl20GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params *float32) {
		C.gl20GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params *int32) {
		C.gl20GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params *float32) {
		C.gl20GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params *int32) {
		C.gl20GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params *float32) {
		C.gl20GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params *int32) {
		C.gl20GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Hint = func(target, mode uint32) {
		C.gl20Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		C.gl20Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		C.gl20Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		C.gl20Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		C.gl20Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexdv = func(c *float64) {
		C.gl20Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(c)))
	}

	glc.Indexfv = func(c *float32) {
		C.gl20Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(c)))
	}

	glc.Indexiv = func(c *int32) {
		C.gl20Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(c)))
	}

	glc.Indexsv = func(c *int16) {
		C.gl20Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(c)))
	}

	glc.IndexMask = func(mask uint32) {
		C.gl20IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		C.gl20InitNames(glc.context)
	}

	glc.IsEnabled = func(cap uint32) {
		C.gl20IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		return C.gl20IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		C.gl20Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		C.gl20Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params *float32) {
		C.gl20Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Lightiv = func(light, pname uint32, params *int32) {
		C.gl20Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		C.gl20LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		C.gl20LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params *float32) {
		C.gl20LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.LightModeliv = func(pname uint32, params *int32) {
		C.gl20LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		C.gl20LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		C.gl20LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		C.gl20ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		C.gl20LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m *float64) {
		C.gl20LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadMatrixf = func(m *float32) {
		C.gl20LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.LoadName = func(name uint32) {
		C.gl20LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		C.gl20LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points *float64) {
		C.gl20Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points *float32) {
		C.gl20Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64) {
		C.gl20Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32) {
		C.gl20Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		C.gl20MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		C.gl20MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		C.gl20MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		C.gl20MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		C.gl20Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		C.gl20Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params *float32) {
		C.gl20Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Materialiv = func(face, pname uint32, params *int32) {
		C.gl20Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.MatrixMode = func(mode uint32) {
		C.gl20MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m *float64) {
		C.gl20MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultMatrixf = func(m *float32) {
		C.gl20MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.NewList = func(list uint32, mode uint32) {
		C.gl20NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		C.gl20EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		C.gl20Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		C.gl20Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		C.gl20Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		C.gl20Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		C.gl20Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v *int8) {
		C.gl20Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Normal3dv = func(v *float64) {
		C.gl20Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Normal3fv = func(v *float32) {
		C.gl20Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Normal3iv = func(v *int32) {
		C.gl20Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Normal3sv = func(v *int16) {
		C.gl20Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		C.gl20Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		C.gl20PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values *float32) {
		C.gl20PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values *uint32) {
		C.gl20PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values *uint16) {
		C.gl20PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		C.gl20PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		C.gl20PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		C.gl20PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		C.gl20PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		C.gl20PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		C.gl20PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		C.gl20PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask *uint8) {
		C.gl20PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(mask)))
	}

	glc.PushAttrib = func(mask uint32) {
		C.gl20PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		C.gl20PopAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		C.gl20PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		C.gl20PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		C.gl20PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		C.gl20PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		C.gl20RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		C.gl20RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		C.gl20RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		C.gl20RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		C.gl20RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		C.gl20RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		C.gl20RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		C.gl20RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		C.gl20RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		C.gl20RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		C.gl20RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		C.gl20RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v *float64) {
		C.gl20RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos2fv = func(v *float32) {
		C.gl20RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos2iv = func(v *int32) {
		C.gl20RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos2sv = func(v *int16) {
		C.gl20RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos3dv = func(v *float64) {
		C.gl20RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos3fv = func(v *float32) {
		C.gl20RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos3iv = func(v *int32) {
		C.gl20RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos3sv = func(v *int16) {
		C.gl20RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos4dv = func(v *float64) {
		C.gl20RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos4fv = func(v *float32) {
		C.gl20RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos4iv = func(v *int32) {
		C.gl20RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos4sv = func(v *int16) {
		C.gl20RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.ReadBuffer = func(mode uint32) {
		C.gl20ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		C.gl20Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		C.gl20Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		C.gl20Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		C.gl20Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 *float64) {
		C.gl20Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
	}

	glc.Rectfv = func(v1, v2 *float32) {
		C.gl20Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
	}

	glc.Rectiv = func(v1, v2 *int32) {
		C.gl20Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
	}

	glc.Rectsv = func(v1, v2 *int16) {
		C.gl20Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
	}

	glc.RenderMode = func(mode uint32) int32 {
		return int32(C.gl20RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		C.gl20Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		C.gl20Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		C.gl20Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		C.gl20Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		C.gl20Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer *uint32) {
		C.gl20SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
	}

	glc.ShadeModel = func(mode uint32) {
		C.gl20ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		C.gl20StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		C.gl20StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		C.gl20StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		C.gl20TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		C.gl20TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		C.gl20TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		C.gl20TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		C.gl20TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		C.gl20TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		C.gl20TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		C.gl20TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		C.gl20TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		C.gl20TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		C.gl20TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		C.gl20TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		C.gl20TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		C.gl20TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		C.gl20TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		C.gl20TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v *float64) {
		C.gl20TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord1fv = func(v *float32) {
		C.gl20TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord1iv = func(v *int32) {
		C.gl20TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord1sv = func(v *int16) {
		C.gl20TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord2dv = func(v *float64) {
		C.gl20TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord2fv = func(v *float32) {
		C.gl20TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord2iv = func(v *int32) {
		C.gl20TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord2sv = func(v *int16) {
		C.gl20TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord3dv = func(v *float64) {
		C.gl20TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord3fv = func(v *float32) {
		C.gl20TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord3iv = func(v *int32) {
		C.gl20TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord3sv = func(v *int16) {
		C.gl20TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord4dv = func(v *float64) {
		C.gl20TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord4fv = func(v *float32) {
		C.gl20TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord4iv = func(v *int32) {
		C.gl20TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord4sv = func(v *int16) {
		C.gl20TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		C.gl20TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		C.gl20TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params *float32) {
		C.gl20TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexEnviv = func(target, pname uint32, params *int32) {
		C.gl20TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		C.gl20TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		C.gl20TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		C.gl20TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params *float64) {
		C.gl20TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.TexGenfv = func(coord, pname uint32, params *float32) {
		C.gl20TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexGeniv = func(coord, pname uint32, params *int32) {
		C.gl20TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		C.gl20TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		C.gl20TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params *float32) {
		C.gl20TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexParameteriv = func(target, pname uint32, params *int32) {
		C.gl20TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Translated = func(x, y, z float64) {
		C.gl20Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		C.gl20Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		C.gl20Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		C.gl20Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		C.gl20Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		C.gl20Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		C.gl20Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		C.gl20Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		C.gl20Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		C.gl20Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		C.gl20Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		C.gl20Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		C.gl20Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		C.gl20Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		C.gl20Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetColorTable = func(target, format, Type uint32, table unsafe.Pointer) {
		C.gl20GetColorTable(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), table)
	}

	glc.GetColorTableParameterfv = func(target, pname uint32, params *float32) {
		C.gl20GetColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetColorTableParameteriv = func(target, pname uint32, params *int32) {
		C.gl20GetColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionFilter = func(target, format, Type uint32, image unsafe.Pointer) {
		C.gl20GetConvolutionFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), image)
	}

	glc.GetConvolutionParameterfv = func(target, pname uint32, params *float32) {
		C.gl20GetConvolutionParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionParameteriv = func(target, pname uint32, params *int32) {
		C.gl20GetConvolutionParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetHistogram = func(target uint32, reset bool, format, Type uint32, values unsafe.Pointer) {
		C.gl20GetHistogram(glc.context, C.GLenum(target), boolToGL(reset), C.GLenum(format), C.GLenum(Type), values)
	}

	glc.GetHistogramParameterfv = func(target, pname uint32, params *float32) {
		C.gl20GetHistogramParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetHistogramParameteriv = func(target, pname uint32, params *int32) {
		C.gl20GetHistogramParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetSeparableFilter = func(target, format, Type uint32, row, column, span unsafe.Pointer) {
		C.gl20GetSeparableFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), row, column, span)
	}

	glc.Histogram = func(target uint32, width int32, internalformat uint32, sink bool) {
		C.gl20Histogram(glc.context, C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.Minmax = func(target, internalformat uint32, sink bool) {
		C.gl20Minmax(glc.context, C.GLenum(target), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.MultiTexCoord1s = func(target uint32, s int16) {
		C.gl20MultiTexCoord1s(glc.context, C.GLenum(target), C.GLshort(s))
	}

	glc.MultiTexCoord1i = func(target uint32, s int32) {
		C.gl20MultiTexCoord1i(glc.context, C.GLenum(target), C.GLint(s))
	}

	glc.MultiTexCoord1f = func(target uint32, s float32) {
		C.gl20MultiTexCoord1f(glc.context, C.GLenum(target), C.GLfloat(s))
	}

	glc.MultiTexCoord1d = func(target uint32, s float64) {
		C.gl20MultiTexCoord1d(glc.context, C.GLenum(target), C.GLdouble(s))
	}

	glc.MultiTexCoord2s = func(target uint32, s, t int16) {
		C.gl20MultiTexCoord2s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t))
	}

	glc.MultiTexCoord2i = func(target uint32, s, t int32) {
		C.gl20MultiTexCoord2i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t))
	}

	glc.MultiTexCoord2f = func(target uint32, s, t float32) {
		C.gl20MultiTexCoord2f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
	}

	glc.MultiTexCoord2d = func(target uint32, s, t float64) {
		C.gl20MultiTexCoord2d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
	}

	glc.MultiTexCoord3s = func(target uint32, s, t, r int16) {
		C.gl20MultiTexCoord3s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.MultiTexCoord3i = func(target uint32, s, t, r int32) {
		C.gl20MultiTexCoord3i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.MultiTexCoord3f = func(target uint32, s, t, r float32) {
		C.gl20MultiTexCoord3f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.MultiTexCoord3d = func(target uint32, s, t, r float64) {
		C.gl20MultiTexCoord3d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.MultiTexCoord4s = func(target uint32, s, t, r, q int16) {
		C.gl20MultiTexCoord4s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.MultiTexCoord4i = func(target uint32, s, t, r, q int32) {
		C.gl20MultiTexCoord4i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.MultiTexCoord4f = func(target uint32, s, t, r, q float32) {
		C.gl20MultiTexCoord4f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.MultiTexCoord4d = func(target uint32, s, t, r, q float64) {
		C.gl20MultiTexCoord4d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.MultiTexCoord1sv = func(target uint32, v *int16) {
		C.gl20MultiTexCoord1sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1iv = func(target uint32, v *int32) {
		C.gl20MultiTexCoord1iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1fv = func(target uint32, v *float32) {
		C.gl20MultiTexCoord1fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1dv = func(target uint32, v *float64) {
		C.gl20MultiTexCoord1dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2sv = func(target uint32, v *int16) {
		C.gl20MultiTexCoord2sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2iv = func(target uint32, v *int32) {
		C.gl20MultiTexCoord2iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2fv = func(target uint32, v *float32) {
		C.gl20MultiTexCoord2fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2dv = func(target uint32, v *float64) {
		C.gl20MultiTexCoord2dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3sv = func(target uint32, v *int16) {
		C.gl20MultiTexCoord3sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3iv = func(target uint32, v *int32) {
		C.gl20MultiTexCoord3iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3fv = func(target uint32, v *float32) {
		C.gl20MultiTexCoord3fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3dv = func(target uint32, v *float64) {
		C.gl20MultiTexCoord3dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4sv = func(target uint32, v *int16) {
		C.gl20MultiTexCoord4sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4iv = func(target uint32, v *int32) {
		C.gl20MultiTexCoord4iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4fv = func(target uint32, v *float32) {
		C.gl20MultiTexCoord4fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4dv = func(target uint32, v *float64) {
		C.gl20MultiTexCoord4dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.ResetHistogram = func(target uint32) {
		C.gl20ResetHistogram(glc.context, C.GLenum(target))
	}

	glc.ResetMinmax = func(target uint32) {
		C.gl20ResetMinmax(glc.context, C.GLenum(target))
	}

	glc.SeparableFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, row, column unsafe.Pointer) {
		C.gl20SeparableFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), row, column)
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		var cRes *C.GLboolean
		status = C.gl20AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		C.gl20ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		C.gl20DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl20DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		C.gl20GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		C.gl20PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		C.gl20CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		C.gl20CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		C.gl20CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		C.gl20CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		C.gl20BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures *uint32) {
		C.gl20DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.GenTextures = func(n int32, textures *uint32) {
		C.gl20GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.IsTexture = func(texture uint32) bool {
		return C.gl20IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		C.gl20EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		C.gl20DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.Indexub = func(c uint8) {
		C.gl20Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexubv = func(c *uint8) {
		C.gl20Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(c)))
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.PushClientAttrib = func(mask uint32) {
		C.gl20PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PrioritizeTextures = func(n int32, textures *uint32, priorities *float32) {
		C.gl20PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLclampf)(unsafe.Pointer(priorities)))
	}

	glc.PopClientAttrib = func() {
		C.gl20PopClientAttrib(glc.context)
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.BlendColor = func(red, green, blue, alpha float32) {
		C.gl20BlendColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		C.gl20BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		C.gl20CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DrawRangeElements = func(mode uint32, start, end uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl20DrawRangeElements(glc.context, C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.TexImage3D = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20TexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl20TexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.ActiveTexture = func(texture uint32) {
		C.gl20ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		C.gl20ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl20CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl20CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl20CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl20CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl20CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl20CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.GetCompressedTexImage = func(target uint32, lod int32, img unsafe.Pointer) {
		C.gl20GetCompressedTexImage(glc.context, C.GLenum(target), C.GLint(lod), img)
	}

	glc.LoadTransposeMatrixd = func(m *float64) {
		C.gl20LoadTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadTransposeMatrixf = func(m *float64) {
		C.gl20LoadTransposeMatrixf(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixd = func(m *float64) {
		C.gl20MultTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixf = func(m *float32) {
		C.gl20MultTransposeMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.SampleCoverage = func(value float32, invert bool) {
		C.gl20SampleCoverage(glc.context, C.GLclampf(value), boolToGL(invert))
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		C.gl20BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.FogCoordPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20FogCoordPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.FogCoordd = func(coord float64) {
		C.gl20FogCoordd(glc.context, C.GLdouble(coord))
	}

	glc.FogCoordf = func(coord float32) {
		C.gl20FogCoordf(glc.context, C.GLfloat(coord))
	}

	glc.FogCoorddv = func(coord *float64) {
		C.gl20FogCoorddv(glc.context, (*C.GLdouble)(unsafe.Pointer(coord)))
	}

	glc.FogCoordfv = func(coord *float32) {
		C.gl20FogCoordfv(glc.context, (*C.GLfloat)(unsafe.Pointer(coord)))
	}

	glc.MultiDrawArrays = func(mode uint32, first *int32, count *int32, primcount int32) {
		C.gl20MultiDrawArrays(glc.context, C.GLenum(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), C.GLsizei(primcount))
	}

	glc.MultiDrawElements = func(mode uint32, count *int32, Type uint32, indices unsafe.Pointer, primcount int32) {
		C.gl20MultiDrawElements(glc.context, C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(count)), C.GLenum(Type), indices, C.GLsizei(primcount))
	}

	glc.PointParameterf = func(pname uint32, param float32) {
		C.gl20PointParameterf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PointParameteri = func(pname uint32, param int32) {
		C.gl20PointParameteri(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.SecondaryColor3b = func(red, green, blue int8) {
		C.gl20SecondaryColor3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.SecondaryColor3s = func(red, green, blue int16) {
		C.gl20SecondaryColor3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.SecondaryColor3i = func(red, green, blue int32) {
		C.gl20SecondaryColor3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.SecondaryColor3f = func(red, green, blue float32) {
		C.gl20SecondaryColor3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.SecondaryColor3d = func(red, green, blue float64) {
		C.gl20SecondaryColor3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.SecondaryColor3ub = func(red, green, blue uint8) {
		C.gl20SecondaryColor3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.SecondaryColor3us = func(red, green, blue uint16) {
		C.gl20SecondaryColor3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.SecondaryColor3ui = func(red, green, blue uint32) {
		C.gl20SecondaryColor3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.SecondaryColor3bv = func(v *int8) {
		C.gl20SecondaryColor3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3sv = func(v *int16) {
		C.gl20SecondaryColor3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3iv = func(v *int32) {
		C.gl20SecondaryColor3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3fv = func(v *float32) {
		C.gl20SecondaryColor3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3dv = func(v *float64) {
		C.gl20SecondaryColor3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3ubv = func(v *uint8) {
		C.gl20SecondaryColor3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3usv = func(v *uint16) {
		C.gl20SecondaryColor3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3uiv = func(v *uint32) {
		C.gl20SecondaryColor3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl20SecondaryColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.WindowPos2s = func(x, y int16) {
		C.gl20WindowPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.WindowPos2i = func(x, y int32) {
		C.gl20WindowPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.WindowPos2f = func(x, y float32) {
		C.gl20WindowPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.WindowPos2d = func(x, y float64) {
		C.gl20WindowPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.WindowPos3s = func(x, y, z int16) {
		C.gl20WindowPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.WindowPos3i = func(x, y, z int32) {
		C.gl20WindowPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.WindowPos3f = func(x, y, z float32) {
		C.gl20WindowPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.WindowPos3d = func(x, y, z float64) {
		C.gl20WindowPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.WindowPos2sv = func(v *int16) {
		C.gl20WindowPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos2iv = func(v *int32) {
		C.gl20WindowPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos2fv = func(v *float32) {
		C.gl20WindowPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos2dv = func(v *float64) {
		C.gl20WindowPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.WindowPos3sv = func(v *int16) {
		C.gl20WindowPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos3iv = func(v *int32) {
		C.gl20WindowPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos3fv = func(v *float32) {
		C.gl20WindowPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos3dv = func(v *float64) {
		C.gl20WindowPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		C.gl20BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		C.gl20BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		C.gl20BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		C.gl20BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.DeleteBuffers = func(n int32, buffers *uint32) {
		C.gl20DeleteBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.DeleteQueries = func(n int32, ids *uint32) {
		C.gl20DeleteQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GenBuffers = func(n int32, buffers *uint32) {
		C.gl20GenBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.GenQueries = func(n int32, ids *uint32) {
		C.gl20GenQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GetBufferParameteriv = func(target, value uint32, data *int32) {
		C.gl20GetBufferParameteriv(glc.context, C.GLenum(target), C.GLenum(value), (*C.GLint)(unsafe.Pointer(data)))
	}

	glc.GetBufferPointerv = func(target, pname uint32, params unsafe.Pointer) {
		C.gl20GetBufferPointerv(glc.context, C.GLenum(target), C.GLenum(pname), params)
	}

	glc.GetBufferSubData = func(target uint32, offset int32, size int32, data unsafe.Pointer) {
		C.gl20GetBufferSubData(glc.context, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}

	glc.GetQueryObjectiv = func(id uint32, pname uint32, params *int32) {
		C.gl20GetQueryObjectiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetQueryObjectuiv = func(id uint32, pname uint32, params *uint32) {
		C.gl20GetQueryObjectuiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(params)))
	}

	glc.GetQueryiv = func(target, pname uint32, params *int32) {
		C.gl20GetQueryiv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.IsBuffer = func(buffer uint32) bool {
		return C.gl20IsBuffer(glc.context, C.GLuint(buffer)) != 0
	}

	glc.IsQuery = func(id uint32) bool {
		return C.gl20IsQuery(glc.context, C.GLuint(id)) != 0
	}

	glc.MapBuffer = func(target, access uint32) unsafe.Pointer {
		return unsafe.Pointer(C.gl20MapBuffer(glc.context, C.GLenum(target), C.GLenum(access)))
	}

	glc.UnmapBuffer = func(target uint32) bool {
		return C.gl20UnmapBuffer(glc.context, C.GLenum(target)) != 0
	}

	glc.AttachShader = func(program, shader uint32) {
		C.gl20AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl20BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha uint32) {
		C.gl20BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		C.gl20CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		return uint32(C.gl20CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType uint32) uint32 {
		return uint32(C.gl20CreateShader(glc.context, C.GLenum(shaderType)))
	}

	glc.DeleteProgram = func(program uint32) {
		C.gl20DeleteProgram(glc.context, C.GLuint(program))
	}

	glc.DeleteShader = func(shader uint32) {
		C.gl20DeleteShader(glc.context, C.GLuint(shader))
	}

	glc.DetachShader = func(program, shader uint32) {
		C.gl20DetachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.EnableVertexAttribArray = func(index uint32) {
		C.gl20EnableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DisableVertexAttribArray = func(index uint32) {
		C.gl20DisableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DrawBuffers = func(n int32, bufs *uint32) {
		C.gl20DrawBuffers(glc.context, C.GLsizei(n), (*C.GLenum)(unsafe.Pointer(bufs)))
	}

	glc.GetActiveAttrib = func(program, index uint32, bufSize int32) (length int32, size int32, Type uint32, name string) {
		var (
			cname C.GLchar
		)
		C.gl20GetActiveAttrib(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
		name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
		return
	}

	glc.GetActiveUniform = func(program, index uint32, bufSize int32, length *int32, size *int32, Type *uint32, name *byte) {
		C.gl20GetActiveUniform(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(Type)), (*C.GLchar)(unsafe.Pointer(name)))
	}

	glc.GetAttachedShaders = func(program uint32, maxCount int32, count *int32, shaders *uint32) {
		C.gl20GetAttachedShaders(glc.context, C.GLuint(program), C.GLsizei(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
	}

	glc.GetAttribLocation = func(program uint32, name *byte) int32 {
		return int32(C.gl20GetAttribLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetProgramiv = func(program uint32, pname uint32, params *int32) {
		C.gl20GetProgramiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetProgramInfoLog = func(program uint32, maxLength int32, length *int32, infoLog *byte) {
		C.gl20GetProgramInfoLog(glc.context, C.GLuint(program), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderiv = func(program uint32, pname uint32, params *int32) {
		C.gl20GetShaderiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetShaderInfoLog = func(shader uint32, maxLength int32, length *int32, infoLog *byte) {
		C.gl20GetShaderInfoLog(glc.context, C.GLuint(shader), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderSource = func(shader uint32, bufSize int32, length *int32, source *byte) {
		C.gl20GetShaderSource(glc.context, C.GLuint(shader), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
	}

	glc.GetUniformfv = func(program uint32, location int32, params *float32) {
		C.gl20GetUniformfv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetUniformiv = func(program uint32, location int32, params *int32) {
		C.gl20GetUniformiv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetUniformLocation = func(program uint32, name *byte) int32 {
		return int32(C.gl20GetUniformLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetVertexAttribdv = func(index uint32, pname uint32, params *float64) {
		C.gl20GetVertexAttribdv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribfv = func(index uint32, pname uint32, params *float32) {
		C.gl20GetVertexAttribfv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribiv = func(index uint32, pname uint32, params *int32) {
		C.gl20GetVertexAttribiv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribPointerv = func(index uint32, pname uint32, pointer unsafe.Pointer) {
		C.gl20GetVertexAttribPointerv(glc.context, C.GLuint(index), C.GLenum(pname), pointer)
	}

	glc.IsProgram = func(program uint32) bool {
		return C.gl20IsProgram(glc.context, C.GLuint(program)) != 0
	}

	glc.IsShader = func(shader uint32) bool {
		return C.gl20IsShader(glc.context, C.GLuint(shader)) != 0
	}

	glc.LinkProgram = func(program uint32) {
		C.gl20LinkProgram(glc.context, C.GLuint(program))
	}

	glc.ShaderSource = func(shader uint32, count int32, string **byte, length *int32) {
		C.gl20ShaderSource(glc.context, C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(string)), (*C.GLint)(unsafe.Pointer(length)))
	}

	glc.StencilFuncSeparate = func(face, Func uint32, ref int32, mask uint32) {
		C.gl20StencilFuncSeparate(glc.context, C.GLenum(face), C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMaskSeparate = func(face uint32, mask uint32) {
		C.gl20StencilMaskSeparate(glc.context, C.GLenum(face), C.GLuint(mask))
	}

	glc.StencilOpSeparate = func(face, sfail, dpfail, dppass uint32) {
		C.gl20StencilOpSeparate(glc.context, C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
	}

	glc.Uniform1f = func(location int32, v0 float32) {
		C.gl20Uniform1f(glc.context, C.GLint(location), C.GLfloat(v0))
	}

	glc.Uniform2f = func(location int32, v0, v1 float32) {
		C.gl20Uniform2f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.Uniform3f = func(location int32, v0, v1, v2 float32) {
		C.gl20Uniform3f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Uniform4f = func(location int32, v0, v1, v2, v3 float32) {
		C.gl20Uniform4f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.Uniform1i = func(location, v0 int32) {
		C.gl20Uniform1i(glc.context, C.GLint(location), C.GLint(v0))
	}

	glc.Uniform2i = func(location, v0, v1 int32) {
		C.gl20Uniform2i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1))
	}

	glc.Uniform3i = func(location, v0, v1, v2 int32) {
		C.gl20Uniform3i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
	}

	glc.Uniform4i = func(location, v0, v1, v2, v3 int32) {
		C.gl20Uniform4i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
	}

	glc.Uniform1fv = func(location int32, count int32, value *float32) {
		C.gl20Uniform1fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform2fv = func(location int32, count int32, value *float32) {
		C.gl20Uniform2fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform3fv = func(location int32, count int32, value *float32) {
		C.gl20Uniform3fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform4fv = func(location int32, count int32, value *float32) {
		C.gl20Uniform4fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform1iv = func(location int32, count int32, value *int32) {
		C.gl20Uniform1iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform2iv = func(location int32, count int32, value *int32) {
		C.gl20Uniform2iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform3iv = func(location int32, count int32, value *int32) {
		C.gl20Uniform3iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform4iv = func(location int32, count int32, value *int32) {
		C.gl20Uniform4iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.UseProgram = func(program uint32) {
		C.gl20UseProgram(glc.context, C.GLuint(program))
	}

	glc.ValidateProgram = func(program uint32) {
		C.gl20ValidateProgram(glc.context, C.GLuint(program))
	}

	glc.VertexAttribPointer = func(index uint32, size int32, Type uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
		C.gl20VertexAttribPointer(glc.context, C.GLuint(index), C.GLint(size), C.GLenum(Type), boolToGL(normalized), C.GLsizei(stride), pointer)
	}

	glc.VertexAttrib1f = func(index uint32, v0 float32) {
		C.gl20VertexAttrib1f(glc.context, C.GLuint(index), C.GLfloat(v0))
	}

	glc.VertexAttrib1s = func(index uint32, v0 int16) {
		C.gl20VertexAttrib1s(glc.context, C.GLuint(index), C.GLshort(v0))
	}

	glc.VertexAttrib1d = func(index uint32, v0 float64) {
		C.gl20VertexAttrib1d(glc.context, C.GLuint(index), C.GLdouble(v0))
	}

	glc.VertexAttrib2f = func(index uint32, v0, v1 float32) {
		C.gl20VertexAttrib2f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.VertexAttrib2s = func(index uint32, v0, v1 int16) {
		C.gl20VertexAttrib2s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
	}

	glc.VertexAttrib2d = func(index uint32, v0, v1 float64) {
		C.gl20VertexAttrib2d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
	}

	glc.VertexAttrib3f = func(index uint32, v0, v1, v2 float32) {
		C.gl20VertexAttrib3f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.VertexAttrib3s = func(index uint32, v0, v1, v2 int16) {
		C.gl20VertexAttrib3s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
	}

	glc.VertexAttrib3d = func(index uint32, v0, v1, v2 float64) {
		C.gl20VertexAttrib3d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.VertexAttrib4f = func(index uint32, v0, v1, v2, v3 float32) {
		C.gl20VertexAttrib4f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.VertexAttrib4s = func(index uint32, v0, v1, v2, v3 int16) {
		C.gl20VertexAttrib4s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
	}

	glc.VertexAttrib4d = func(index uint32, v0, v1, v2, v3 float64) {
		C.gl20VertexAttrib4d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
	}

	glc.VertexAttrib4Nuv = func(index uint32, v0, v1, v2, v3 uint8) {
		C.gl20VertexAttrib4Nuv(glc.context, C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
	}

	glc.VertexAttrib1fv = func(index uint32, v *float32) {
		C.gl20VertexAttrib1fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1sv = func(index uint32, v *int16) {
		C.gl20VertexAttrib1sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1dv = func(index uint32, v *float64) {
		C.gl20VertexAttrib1dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2fv = func(index uint32, v *float32) {
		C.gl20VertexAttrib2fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2sv = func(index uint32, v *int16) {
		C.gl20VertexAttrib2sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2dv = func(index uint32, v *float64) {
		C.gl20VertexAttrib2dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3fv = func(index uint32, v *float32) {
		C.gl20VertexAttrib3fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3sv = func(index uint32, v *int16) {
		C.gl20VertexAttrib3sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3dv = func(index uint32, v *float64) {
		C.gl20VertexAttrib3dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4fv = func(index uint32, v *float32) {
		C.gl20VertexAttrib4fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4sv = func(index uint32, v *int16) {
		C.gl20VertexAttrib4sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4dv = func(index uint32, v *float64) {
		C.gl20VertexAttrib4dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4iv = func(index uint32, v *int32) {
		C.gl20VertexAttrib4iv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4bv = func(index uint32, v *int8) {
		C.gl20VertexAttrib4bv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4ubv = func(index uint32, v *uint8) {
		C.gl20VertexAttrib4ubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4usv = func(index uint32, v *uint16) {
		C.gl20VertexAttrib4usv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4uiv = func(index uint32, v *uint32) {
		C.gl20VertexAttrib4uiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nbv = func(index uint32, v *int8) {
		C.gl20VertexAttrib4Nbv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nsv = func(index uint32, v *int16) {
		C.gl20VertexAttrib4Nsv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Niv = func(index uint32, v *int32) {
		C.gl20VertexAttrib4Niv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nubv = func(index uint32, v *uint8) {
		C.gl20VertexAttrib4Nubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nusv = func(index uint32, v *uint16) {
		C.gl20VertexAttrib4Nusv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nuiv = func(index uint32, v *uint32) {
		C.gl20VertexAttrib4Nuiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	if !versionSupported(glc) {
		return nil
	}
	return glc
}
