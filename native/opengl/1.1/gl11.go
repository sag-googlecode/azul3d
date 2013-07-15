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
	VIEW_CLASS_32_BITS                                         = 0x82C8
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	R16UI                                                      = 0x8234
	MATRIX12_ARB                                               = 0x88CC
	REG_23_ATI                                                 = 0x8938
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	PN_TRIANGLES_ATI                                           = 0x87F0
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	GL_2_BYTES                                                 = 0x1407
	INVERT                                                     = 0x150A
	N3F_V3F                                                    = 0x2A25
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	HILO8_NV                                                   = 0x885E
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	POLYGON_OFFSET_LINE                                        = 0x2A02
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	RGB8_EXT                                                   = 0x8051
	STENCIL_BUFFER_BIT                                         = 0x00000400
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	COPY                                                       = 0x1503
	REFLECTION_MAP_OES                                         = 0x8512
	ONE_EXT                                                    = 0x87DE
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	FLOAT16_VEC4_NV                                            = 0x8FFB
	INVERTED_SCREEN_W_REND                                     = 0x8491
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	BLEND_SRC                                                  = 0x0BE1
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	DRAW_BUFFER15_ARB                                          = 0x8834
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	TEXTURE21_ARB                                              = 0x84D5
	PERCENTAGE_AMD                                             = 0x8BC3
	SEPARABLE_2D                                               = 0x8012
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	UNIFORM_BUFFER                                             = 0x8A11
	COLOR_ATTACHMENT5                                          = 0x8CE5
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	COLOR4_BIT_PGI                                             = 0x00020000
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	CONSTANT_ALPHA_EXT                                         = 0x8003
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	FILL                                                       = 0x1B02
	RGBA_DXT5_S3TC                                             = 0x83A4
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	VARIABLE_C_NV                                              = 0x8525
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	CLAMP_READ_COLOR                                           = 0x891C
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	TEXTURE6                                                   = 0x84C6
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	HISTOGRAM_SINK_EXT                                         = 0x802D
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	DEPTH_COMPONENT32F                                         = 0x8CAC
	MAX_CLIP_DISTANCES                                         = 0x0D32
	CONVOLUTION_HEIGHT                                         = 0x8019
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	RG32F                                                      = 0x8230
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	VARIANT_ARRAY_EXT                                          = 0x87E8
	MAX_IMAGE_UNITS                                            = 0x8F38
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	SAMPLER_BUFFER                                             = 0x8DC2
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	OPERAND0_ALPHA_EXT                                         = 0x8598
	OBJECT_TYPE                                                = 0x9112
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	DRAW_BUFFER6_NV                                            = 0x882B
	REG_5_ATI                                                  = 0x8926
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	R                                                          = 0x2002
	CONVOLUTION_1D                                             = 0x8010
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	X_EXT                                                      = 0x87D5
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	GL_3D                                                      = 0x0601
	DSDT_NV                                                    = 0x86F5
	READ_ONLY_ARB                                              = 0x88B8
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	UNDEFINED_APPLE                                            = 0x8A1C
	INTENSITY                                                  = 0x8049
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	TEXTURE15                                                  = 0x84CF
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	TIMEOUT_EXPIRED                                            = 0x911B
	LINE_RESET_TOKEN                                           = 0x0707
	COLOR_ARRAY_SIZE                                           = 0x8081
	SUB_ATI                                                    = 0x8965
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	COMBINE_RGB_ARB                                            = 0x8571
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	SAMPLER_1D_ARB                                             = 0x8B5D
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	RGBA8UI_EXT                                                = 0x8D7C
	TRIANGLE_FAN                                               = 0x0006
	RENDER_MODE                                                = 0x0C40
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	CLAMP_TO_BORDER_NV                                         = 0x812D
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	QUERY_RESULT                                               = 0x8866
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	SRC2_ALPHA                                                 = 0x858A
	RGBA16F_ARB                                                = 0x881A
	VOLATILE_APPLE                                             = 0x8A1A
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	LINEAR_DETAIL_SGIS                                         = 0x8097
	MULTISAMPLE_SGIS                                           = 0x809D
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	MAX_SHININESS_NV                                           = 0x8504
	R1UI_V3F_SUN                                               = 0x85C4
	DSDT8_NV                                                   = 0x8709
	VERTEX_STREAM5_ATI                                         = 0x8771
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	BYTE                                                       = 0x1400
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	UNIFORM_SIZE                                               = 0x8A38
	VERSION_1_4                                                = 1
	DRAW_BUFFER2_NV                                            = 0x8827
	REG_8_ATI                                                  = 0x8929
	VENDOR                                                     = 0x1F00
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	EXPAND_NEGATE_NV                                           = 0x8539
	OPERAND0_RGB_ARB                                           = 0x8590
	FLOAT_VEC4_ARB                                             = 0x8B52
	DONT_CARE                                                  = 0x1100
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	QUERY_OBJECT_EXT                                           = 0x9153
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	DEPTH_COMPONENT                                            = 0x1902
	RESCALE_NORMAL                                             = 0x803A
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	SPECULAR                                                   = 0x1202
	LIGHT0                                                     = 0x4000
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	TEXTURE30                                                  = 0x84DE
	VARIABLE_G_NV                                              = 0x8529
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	QUERY_WAIT                                                 = 0x8E13
	SUCCESS_NV                                                 = 0x902F
	NAME_STACK_DEPTH                                           = 0x0D70
	Y_EXT                                                      = 0x87D6
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	ORDER                                                      = 0x0A01
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	WRITE_ONLY                                                 = 0x88B9
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	SOURCE0_RGB_EXT                                            = 0x8580
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	DRAW_BUFFER14_ARB                                          = 0x8833
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	ALPHA_INTEGER                                              = 0x8D97
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	SET_AMD                                                    = 0x874A
	MVP_MATRIX_EXT                                             = 0x87E3
	RGBA32F                                                    = 0x8814
	RGB_FLOAT32_ATI                                            = 0x8815
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	BLEND_EQUATION_ALPHA                                       = 0x883D
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	RGBA4_EXT                                                  = 0x8056
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	STENCIL_BACK_FUNC                                          = 0x8800
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	RGBA16F_EXT                                                = 0x881A
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	LUMINANCE16I_EXT                                           = 0x8D8C
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	PATH_FORMAT_PS_NV                                          = 0x9071
	R16F_EXT                                                   = 0x822D
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	BUFFER_MAPPED_ARB                                          = 0x88BC
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	VERSION_2_0                                                = 1
	RED_BIAS                                                   = 0x0D15
	RGB5_EXT                                                   = 0x8050
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	UNSIGNED_INT_24_8                                          = 0x84FA
	COMBINE4_NV                                                = 0x8503
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	INT16_NV                                                   = 0x8FE4
	BLEND_COLOR                                                = 0x8005
	COLOR_TABLE_BIAS                                           = 0x80D7
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	DRAW_BUFFER10_NV                                           = 0x882F
	MATRIX20_ARB                                               = 0x88D4
	QUARTER_BIT_ATI                                            = 0x00000010
	FLOAT_VEC4                                                 = 0x8B52
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	FOG_COORDINATE_ARRAY                                       = 0x8457
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	FRONT_RIGHT                                                = 0x0401
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	FLOAT16_NV                                                 = 0x8FF8
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	OPERAND1_RGB                                               = 0x8591
	DEPTH_CLAMP                                                = 0x864F
	IMAGE_CUBE_EXT                                             = 0x9050
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	DRAW_BUFFER0_ATI                                           = 0x8825
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	ALPHA32I_EXT                                               = 0x8D84
	SGIS_texture4D                                             = 1
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	READ_WRITE                                                 = 0x88BA
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	FULL_SUPPORT                                               = 0x82B7
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	CLOSE_PATH_NV                                              = 0x00
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	CLAMP                                                      = 0x2900
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	TEXTURE8                                                   = 0x84C8
	MODELVIEW18_ARB                                            = 0x8732
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	RG_SNORM                                                   = 0x8F91
	IMAGE_BINDING_FORMAT                                       = 0x906E
	OFFSET                                                     = 0x92FC
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	FRONT                                                      = 0x0404
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	SHADER_BINARY_VIV                                          = 0x8FC4
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	UNIFORM                                                    = 0x92E1
	QUADRATIC_ATTENUATION                                      = 0x1209
	QUERY_WAIT_NV                                              = 0x8E13
	COLOR_COMPONENTS                                           = 0x8283
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	AUTO_NORMAL                                                = 0x0D80
	RGB12                                                      = 0x8053
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	DECR_WRAP_EXT                                              = 0x8508
	REPLACE_VALUE_AMD                                          = 0x874B
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	FIXED_ONLY                                                 = 0x891D
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	LOCATION_INDEX                                             = 0x930F
	COORD_REPLACE_NV                                           = 0x8862
	FLOAT_MAT2x3                                               = 0x8B65
	SYNC_FENCE                                                 = 0x9116
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	READ_BUFFER                                                = 0x0C02
	TEXTURE19_ARB                                              = 0x84D3
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	POINT_SPRITE_NV                                            = 0x8861
	NEVER                                                      = 0x0200
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	OP_LOG_BASE_2_EXT                                          = 0x8792
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	MODELVIEW27_ARB                                            = 0x873B
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	COLOR_ATTACHMENT4                                          = 0x8CE4
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	BLEND_DST_RGB_OES                                          = 0x80C8
	PATH_FILL_MASK_NV                                          = 0x9081
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	SGIX_pixel_tiles                                           = 1
	INTENSITY8                                                 = 0x804B
	IUI_N3F_V3F_EXT                                            = 0x81B0
	SRC2_RGB                                                   = 0x8582
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	NEXT_BUFFER_NV                                             = -2
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	SGI_texture_color_table                                    = 1
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	OPERAND0_RGB_EXT                                           = 0x8590
	BUFFER_ACCESS_OES                                          = 0x88BB
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	SIGNALED_APPLE                                             = 0x9119
	C4UB_V2F                                                   = 0x2A22
	PACK_RESAMPLE_SGIX                                         = 0x842C
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	VARIABLE_E_NV                                              = 0x8527
	CURRENT_VERTEX_EXT                                         = 0x87E2
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	MINMAX_FORMAT_EXT                                          = 0x802F
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	MULTISAMPLE_EXT                                            = 0x809D
	SPARE1_NV                                                  = 0x852F
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	SWIZZLE_STQ_ATI                                            = 0x8977
	CONTEXT_PROFILE_MASK                                       = 0x9126
	INCR                                                       = 0x1E02
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	CURRENT_RASTER_INDEX                                       = 0x0B05
	RGBA4_OES                                                  = 0x8056
	AVERAGE_EXT                                                = 0x8335
	PREVIOUS                                                   = 0x8578
	DRAW_BUFFER7_NV                                            = 0x882C
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	YCRCB_422_SGIX                                             = 0x81BB
	CONSTANT_COLOR0_NV                                         = 0x852A
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	MODELVIEW2_ARB                                             = 0x8722
	MODELVIEW16_ARB                                            = 0x8730
	SGIX_convolution_accuracy                                  = 1
	UNSIGNED_BYTE                                              = 0x1401
	SOURCE2_ALPHA_ARB                                          = 0x858A
	REG_17_ATI                                                 = 0x8932
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	XOR                                                        = 0x1506
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	DEPTH_BUFFER_BIT                                           = 0x00000100
	CON_23_ATI                                                 = 0x8958
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	INT_IMAGE_1D                                               = 0x9057
	SRGB_READ                                                  = 0x8297
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	PROXY_HISTOGRAM                                            = 0x8025
	TEXTURE_GATHER                                             = 0x82A2
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	RGB4_S3TC                                                  = 0x83A1
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	DRAW_BUFFER8_ARB                                           = 0x882D
	MAD_ATI                                                    = 0x8968
	PROVOKING_VERTEX                                           = 0x8E4F
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	ZOOM_X                                                     = 0x0D16
	LIST_PRIORITY_SGIX                                         = 0x8182
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	LUMINANCE_SNORM                                            = 0x9011
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	STREAM_COPY_ARB                                            = 0x88E2
	SKIP_DECODE_EXT                                            = 0x8A4A
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	HISTOGRAM_EXT                                              = 0x8024
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	CURRENT_MATRIX_NV                                          = 0x8641
	QUERY_COUNTER_BITS                                         = 0x8864
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	TEXTURE_RECTANGLE                                          = 0x84F5
	SIGNED_NEGATE_NV                                           = 0x853D
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	SPOT_EXPONENT                                              = 0x1205
	HISTOGRAM_FORMAT                                           = 0x8027
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	COUNT_DOWN_NV                                              = 0x9089
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	WRAP_BORDER_SUN                                            = 0x81D4
	SAMPLES_PASSED                                             = 0x8914
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	POLYGON_OFFSET_POINT                                       = 0x2A01
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	DRAW_BUFFER9_NV                                            = 0x882E
	PROGRAM_FORMAT_ARB                                         = 0x8876
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	LIST_MODE                                                  = 0x0B30
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	WRITE_DISCARD_NV                                           = 0x88BE
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	SGIX_sprite                                                = 1
	COLOR_ARRAY_POINTER                                        = 0x8090
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	INDEX_CLEAR_VALUE                                          = 0x0C20
	TEXTURE_MIN_FILTER                                         = 0x2801
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	INT64_NV                                                   = 0x140E
	INTENSITY16UI_EXT                                          = 0x8D79
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	PASS_THROUGH_TOKEN                                         = 0x0700
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	REDUCE_EXT                                                 = 0x8016
	UNSIGNED_NORMALIZED                                        = 0x8C17
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	T2F_C3F_V3F                                                = 0x2A2A
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	SPRITE_AXIS_SGIX                                           = 0x814A
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	MODELVIEW30_ARB                                            = 0x873E
	RELATIVE_MOVE_TO_NV                                        = 0x03
	SAMPLE_MASK_SGIS                                           = 0x80A0
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	IMAGE_2D                                                   = 0x904D
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	TEXTURE17_ARB                                              = 0x84D1
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	TEXTURE20                                                  = 0x84D4
	COORD_REPLACE_ARB                                          = 0x8862
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	MAX_SAMPLES                                                = 0x8D57
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	COMBINER0_NV                                               = 0x8550
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	EXT_texture_object                                         = 1
	COLOR_INDEX8_EXT                                           = 0x80E5
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	COLOR_TABLE_WIDTH                                          = 0x80D9
	MODELVIEW3_ARB                                             = 0x8723
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	TEXTURE_GEN_T                                              = 0x0C61
	MODELVIEW29_ARB                                            = 0x873D
	COLOR_ATTACHMENT1                                          = 0x8CE1
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	LUMINANCE16_ALPHA16                                        = 0x8048
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	DRAW_BUFFER10                                              = 0x882F
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	HIGH_INT                                                   = 0x8DF5
	ARC_TO_NV                                                  = 0xFE
	SGIX_reference_plane                                       = 1
	ALPHA16_EXT                                                = 0x803E
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	TEXTURE5_ARB                                               = 0x84C5
	SAMPLER_3D_ARB                                             = 0x8B5F
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	CLIP_PLANE5                                                = 0x3005
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	YCBCR_MESA                                                 = 0x8757
	AFFINE_2D_NV                                               = 0x9092
	LINE_STRIP_ADJACENCY                                       = 0x000B
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	LUMINANCE8_ALPHA8                                          = 0x8045
	REPLACE_OLDEST_SUN                                         = 0x0003
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	INT_VEC4_ARB                                               = 0x8B55
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	CURRENT_RASTER_COLOR                                       = 0x0B04
	SPARE0_NV                                                  = 0x852E
	MODELVIEW21_ARB                                            = 0x8735
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	REG_20_ATI                                                 = 0x8935
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	DEPTH_STENCIL                                              = 0x84F9
	STENCIL_INDEX1_OES                                         = 0x8D46
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	MAP2_COLOR_4                                               = 0x0DB0
	ACCUM_BUFFER_BIT                                           = 0x00000200
	OBJECT_LINEAR                                              = 0x2401
	FLOAT_MAT2x4                                               = 0x8B66
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	DEBUG_SEVERITY_LOW                                         = 0x9148
	TEXTURE29                                                  = 0x84DD
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	INCR_WRAP_EXT                                              = 0x8507
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	STENCIL_WRITEMASK                                          = 0x0B98
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	CURRENT_MATRIX_ARB                                         = 0x8641
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	BLUE_BIAS                                                  = 0x0D1B
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	TRIANGLE_STRIP                                             = 0x0005
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	TEXTURE20_ARB                                              = 0x84D4
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	SGIX_list_priority                                         = 1
	SCISSOR_TEST                                               = 0x0C11
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	OPERAND2_RGB                                               = 0x8592
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	FOG_COORDINATE                                             = 0x8451
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	GREATER                                                    = 0x0204
	CON_1_ATI                                                  = 0x8942
	SAMPLE_SHADING_ARB                                         = 0x8C36
	INT64_VEC2_NV                                              = 0x8FE9
	VERSION_1_1                                                = 1
	EYE_LINEAR                                                 = 0x2400
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	TRIANGLE_MESH_SUN                                          = 0x8615
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	DEPTH_RANGE                                                = 0x0B70
	GL_1PASS_EXT                                               = 0x80A1
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	MODELVIEW11_ARB                                            = 0x872B
	UNIFORM_OFFSET                                             = 0x8A3B
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	FRONT_AND_BACK                                             = 0x0408
	AND_REVERSE                                                = 0x1502
	CLIP_PLANE2                                                = 0x3002
	FLOAT_VEC2                                                 = 0x8B50
	SHADER_COMPILER                                            = 0x8DFA
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	GL_3D_COLOR                                                = 0x0602
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	VIEW_CLASS_96_BITS                                         = 0x82C5
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	SIGNED_HILO16_NV                                           = 0x86FA
	DRAW_BUFFER12_NV                                           = 0x8831
	MATRIX_PALETTE_ARB                                         = 0x8840
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	TEXTURE_BINDING_1D                                         = 0x8068
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	SRGB_WRITE                                                 = 0x8298
	VERTEX_SHADER_ARB                                          = 0x8B31
	TRANSLATE_X_NV                                             = 0x908E
	TEXTURE_ENV_MODE                                           = 0x2200
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	MAX_WIDTH                                                  = 0x827E
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	RGBA32I                                                    = 0x8D82
	BEVEL_NV                                                   = 0x90A6
	ALPHA32F_ARB                                               = 0x8816
	STREAM_READ_ARB                                            = 0x88E1
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	GL_4PASS_1_EXT                                             = 0x80A5
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	BUFFER_SIZE                                                = 0x8764
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	IS_PER_PATCH                                               = 0x92E7
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	INCR_WRAP_OES                                              = 0x8507
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	SHADER_TYPE                                                = 0x8B4F
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	REG_19_ATI                                                 = 0x8934
	R16_SNORM                                                  = 0x8F98
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	REPLICATE_BORDER                                           = 0x8153
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	INTERPOLATE_ARB                                            = 0x8575
	DRAW_BUFFER8_NV                                            = 0x882D
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	BUFFER_MAP_LENGTH                                          = 0x9120
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	DEPTH_STENCIL_OES                                          = 0x84F9
	Z400_BINARY_AMD                                            = 0x8740
	DRAW_BUFFER14_NV                                           = 0x8833
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	MITER_TRUNCATE_NV                                          = 0x90A8
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	TEXTURE1                                                   = 0x84C1
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	RGB8                                                       = 0x8051
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	ACTIVE_RESOURCES                                           = 0x92F5
	TEXTURE25                                                  = 0x84D9
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	DSDT8_MAG8_NV                                              = 0x870A
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	RGB10_EXT                                                  = 0x8052
	ACTIVE_VARIABLES                                           = 0x9305
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	PRIMARY_COLOR_NV                                           = 0x852C
	INT_VEC2_ARB                                               = 0x8B53
	RGB565                                                     = 0x8D62
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	TRANSPOSE_NV                                               = 0x862C
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	EXTENSIONS                                                 = 0x1F03
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	MAX_EVAL_ORDER                                             = 0x0D30
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	MODELVIEW1_ARB                                             = 0x850A
	DRAW_BUFFER5_ATI                                           = 0x882A
	REG_10_ATI                                                 = 0x892B
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	RENDERER                                                   = 0x1F01
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	COLOR_ATTACHMENT0                                          = 0x8CE0
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	ROUND_NV                                                   = 0x90A4
	UNSIGNED_INT64_NV                                          = 0x140F
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	CLIP_DISTANCE_NV                                           = 0x8C7A
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	RGBA8UI                                                    = 0x8D7C
	DOUBLE_MAT4                                                = 0x8F48
	RGBA_SNORM                                                 = 0x8F93
	IMAGE_CUBE                                                 = 0x9050
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	PREVIOUS_EXT                                               = 0x8578
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	ACTIVE_TEXTURE                                             = 0x84E0
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	INT_SAMPLER_2D                                             = 0x8DCA
	DEPTH_TEST                                                 = 0x0B71
	CONSTANT_ARB                                               = 0x8576
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	DECODE_EXT                                                 = 0x8A49
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	CAVEAT_SUPPORT                                             = 0x82B8
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	FRAMEBUFFER_EXT                                            = 0x8D40
	SGIS_detail_texture                                        = 1
	MAP2_VERTEX_4                                              = 0x0DB8
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	DRAW_BUFFER1_NV                                            = 0x8826
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	DISCRETE_AMD                                               = 0x9006
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	CON_6_ATI                                                  = 0x8947
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	MODELVIEW19_ARB                                            = 0x8733
	DRAW_BUFFER10_ATI                                          = 0x882F
	POLYGON_TOKEN                                              = 0x0703
	LIST_BASE                                                  = 0x0B32
	COLOR_ARRAY_STRIDE                                         = 0x8083
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	CON_24_ATI                                                 = 0x8959
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	TESS_GEN_POINT_MODE                                        = 0x8E79
	IMAGE_2D_EXT                                               = 0x904D
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	PACK_SKIP_PIXELS                                           = 0x0D04
	QUAD_ALPHA4_SGIS                                           = 0x811E
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	TEXTURE_4D_SGIS                                            = 0x8134
	R8I                                                        = 0x8231
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	SOURCE2_ALPHA                                              = 0x858A
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	NO_ERROR                                                   = 0
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	REG_16_ATI                                                 = 0x8931
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	ALPHA32UI_EXT                                              = 0x8D72
	MEDIUM_INT                                                 = 0x8DF4
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	EXT_histogram                                              = 1
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	GL_422_EXT                                                 = 0x80CC
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	BGRA_INTEGER                                               = 0x8D9B
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	TEXTURE_VIEW                                               = 0x82B5
	TEXTURE7_ARB                                               = 0x84C7
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	OR_REVERSE                                                 = 0x150B
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	RGB8UI_EXT                                                 = 0x8D7D
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	MODELVIEW5_ARB                                             = 0x8725
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	BGRA_EXT                                                   = 0x80E1
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	EXT_point_parameters                                       = 1
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	MIPMAP                                                     = 0x8293
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	CLIP_DISTANCE0                                             = 0x3000
	PROGRAM_SEPARABLE                                          = 0x8258
	E_TIMES_F_NV                                               = 0x8531
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	RGBA4_DXT5_S3TC                                            = 0x83A5
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	PROXY_TEXTURE_1D                                           = 0x8063
	MODELVIEW0_ARB                                             = 0x1700
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	OPERAND1_ALPHA_ARB                                         = 0x8599
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	DEPTH_ATTACHMENT                                           = 0x8D00
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	CONST_EYE_NV                                               = 0x86E5
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	BUFFER_ACCESS                                              = 0x88BB
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	DOT3_RGB_EXT                                               = 0x8740
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	RGB_INTEGER_EXT                                            = 0x8D98
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	RGB10_A2                                                   = 0x8059
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	DEPTH_COMPONENT32                                          = 0x81A7
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	BUMP_TARGET_ATI                                            = 0x877C
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	INTENSITY4_EXT                                             = 0x804A
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	MIRROR_CLAMP_ATI                                           = 0x8742
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	MAX_INTEGER_SAMPLES                                        = 0x9110
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	SGIS_pixel_texture                                         = 1
	VERTEX_STREAM3_ATI                                         = 0x876F
	CND0_ATI                                                   = 0x896B
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	CON_21_ATI                                                 = 0x8956
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	MATRIX0_ARB                                                = 0x88C0
	TESS_GEN_SPACING                                           = 0x8E77
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	MULTISAMPLE                                                = 0x809D
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	DEPTH_BIAS                                                 = 0x0D1F
	TEXTURE28_ARB                                              = 0x84DC
	EYE_RADIAL_NV                                              = 0x855B
	DSDT_MAG_NV                                                = 0x86F6
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	RGB_422_APPLE                                              = 0x8A1F
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	FRAMEBUFFER_OES                                            = 0x8D40
	RGBA8_SNORM                                                = 0x8F97
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	UNIFORM_BLOCK                                              = 0x92E2
	DRAW_BUFFER                                                = 0x0C01
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	SLUMINANCE8_NV                                             = 0x8C47
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	CLIP_PLANE0                                                = 0x3000
	COLOR_INDEX1_EXT                                           = 0x80E2
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	INTENSITY8_SNORM                                           = 0x9017
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	TEXTURE_RED_SIZE                                           = 0x805C
	DOUBLE                                                     = 0x140A
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	CURRENT_FOG_COORD                                          = 0x8453
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	BUFFER_VARIABLE                                            = 0x92E5
	MAX_LIST_NESTING                                           = 0x0B31
	RGBA12_EXT                                                 = 0x805A
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	CON_27_ATI                                                 = 0x895C
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	MAP1_VERTEX_4                                              = 0x0D98
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	BGRA                                                       = 0x80E1
	MODELVIEW4_ARB                                             = 0x8724
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	NEGATIVE_Z_EXT                                             = 0x87DB
	BGR_INTEGER                                                = 0x8D9A
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	PATH_FORMAT_SVG_NV                                         = 0x9070
	EXT_cmyka                                                  = 1
	SGIS_texture_select                                        = 1
	TEXTURE_GEN_R                                              = 0x0C62
	COLOR_TABLE_FORMAT                                         = 0x80D8
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	MODELVIEW14_ARB                                            = 0x872E
	DT_BIAS_NV                                                 = 0x8717
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	BOOL_VEC3                                                  = 0x8B58
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	RED_EXT                                                    = 0x1903
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	STENCIL_ATTACHMENT                                         = 0x8D20
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	MIN_EXT                                                    = 0x8007
	MINMAX_SINK                                                = 0x8030
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	SCALE_BY_TWO_NV                                            = 0x853E
	DRAW_BUFFER15                                              = 0x8834
	GREEN_BIT_ATI                                              = 0x00000002
	BOOL_VEC4                                                  = 0x8B59
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	FRAGMENT_COLOR_EXT                                         = 0x834C
	SOURCE1_ALPHA                                              = 0x8589
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	SMOOTH                                                     = 0x1D01
	LUMINANCE6_ALPHA2                                          = 0x8044
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	VECTOR_EXT                                                 = 0x87BF
	MATRIX17_ARB                                               = 0x88D1
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	SLICE_ACCUM_SUN                                            = 0x85CC
	INTENSITY32UI_EXT                                          = 0x8D73
	INT16_VEC4_NV                                              = 0x8FE7
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	TEXTURE27                                                  = 0x84DB
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	ALPHA12                                                    = 0x803D
	TEXTURE18                                                  = 0x84D2
	MATRIX4_ARB                                                = 0x88C4
	CON_0_ATI                                                  = 0x8941
	MAP_READ_BIT                                               = 0x0001
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	MATRIX2_ARB                                                = 0x88C2
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	TIMESTAMP                                                  = 0x8E28
	IMAGE_1D_EXT                                               = 0x904C
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	DRAW_BUFFER11_ARB                                          = 0x8830
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	INFO_LOG_LENGTH                                            = 0x8B84
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	EXT_rescale_normal                                         = 1
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	MODELVIEW17_ARB                                            = 0x8731
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	CON_15_ATI                                                 = 0x8950
	SAMPLER_2D_RECT                                            = 0x8B63
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	ABGR_EXT                                                   = 0x8000
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	VERTEX_STREAM7_ATI                                         = 0x8773
	VERTEX_SHADER_EXT                                          = 0x8780
	DOUBLE_VEC2                                                = 0x8FFC
	EDGE_FLAG                                                  = 0x0B43
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	TRANSLATE_Y_NV                                             = 0x908F
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	FLOAT_MAT3x4                                               = 0x8B68
	ALREADY_SIGNALED                                           = 0x911A
	COPY_PIXEL_TOKEN                                           = 0x0706
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	VERTEX_SOURCE_ATI                                          = 0x8774
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	PIXEL_PACK_BUFFER                                          = 0x88EB
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	SPRITE_MODE_SGIX                                           = 0x8149
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	TEXTURE_BINDING_3D                                         = 0x806A
	VERTEX_PROGRAM_NV                                          = 0x8620
	CURRENT_QUERY_EXT                                          = 0x8865
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	NOTEQUAL                                                   = 0x0205
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	NOR                                                        = 0x1508
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	DRAW_BUFFER1_ATI                                           = 0x8826
	MATRIX11_ARB                                               = 0x88CB
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	ALPHA_BIAS                                                 = 0x0D1D
	MINMAX_SINK_EXT                                            = 0x8030
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	MAX_TEXTURE_COORDS                                         = 0x8871
	RELATIVE_LINE_TO_NV                                        = 0x05
	KEEP                                                       = 0x1E00
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	EIGHTH_BIT_ATI                                             = 0x00000020
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	MAX_CLIP_PLANES                                            = 0x0D32
	COLOR_INDEX                                                = 0x1900
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	SIGNED_INTENSITY8_NV                                       = 0x8708
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	HINT_BIT                                                   = 0x00008000
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	INTENSITY16_EXT                                            = 0x804D
	CUBIC_EXT                                                  = 0x8334
	MAP1_BINORMAL_EXT                                          = 0x8446
	OPERAND0_RGB                                               = 0x8590
	OP_SET_GE_EXT                                              = 0x878C
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	SIGNED_HILO_NV                                             = 0x86F9
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	MODELVIEW1_EXT                                             = 0x850A
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	RGBA                                                       = 0x1908
	DEPTH_STENCIL_MESA                                         = 0x8750
	COLOR_ATTACHMENT9                                          = 0x8CE9
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	BUFFER_OBJECT_EXT                                          = 0x9151
	VERTEX_TEXTURE                                             = 0x829B
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	IMAGE_2D_RECT                                              = 0x904F
	SYNC_STATUS                                                = 0x9114
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	TEXTURE_LIGHT_EXT                                          = 0x8350
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	BACK_RIGHT                                                 = 0x0403
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	IGNORE_BORDER_HP                                           = 0x8150
	COLOR_ATTACHMENT3                                          = 0x8CE3
	RGB32I_EXT                                                 = 0x8D83
	FRACTIONAL_EVEN                                            = 0x8E7C
	ALL_STATIC_DATA_IBM                                        = 103060
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	R3_G3_B2                                                   = 0x2A10
	SOURCE1_ALPHA_ARB                                          = 0x8589
	FRAGMENT_SHADER_ATI                                        = 0x8920
	REG_12_ATI                                                 = 0x892D
	ALPHA8UI_EXT                                               = 0x8D7E
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	STENCIL_TEST                                               = 0x0B90
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	VERSION_3_0                                                = 1
	SGIX_subsample                                             = 1
	LINE_LOOP                                                  = 0x0002
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	WRITE_ONLY_OES                                             = 0x88B9
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	CON_25_ATI                                                 = 0x895A
	COLOR_ATTACHMENT13                                         = 0x8CED
	TEXTURE_ENV_COLOR                                          = 0x2201
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	SAMPLER_1D_SHADOW                                          = 0x8B61
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	CONSTANT_BORDER_HP                                         = 0x8151
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	NEGATIVE_ONE_EXT                                           = 0x87DF
	SAMPLES_PASSED_ARB                                         = 0x8914
	REG_3_ATI                                                  = 0x8924
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	RGB32UI                                                    = 0x8D71
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	CW                                                         = 0x0900
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	BLOCK_INDEX                                                = 0x92FD
	FLOAT_R32_NV                                               = 0x8885
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	BOUNDING_BOX_NV                                            = 0x908D
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	PACK_INVERT_MESA                                           = 0x8758
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	NORMAL_ARRAY                                               = 0x8075
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	PRIMARY_COLOR_ARB                                          = 0x8577
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	SRGB_ALPHA_EXT                                             = 0x8C42
	ENABLE_BIT                                                 = 0x00002000
	RGB_SNORM                                                  = 0x8F92
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	CURRENT_COLOR                                              = 0x0B00
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	FRAMEBUFFER                                                = 0x8D40
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	SIGNED_IDENTITY_NV                                         = 0x853C
	YCBCR_422_APPLE                                            = 0x85B9
	DEPTH24_STENCIL8                                           = 0x88F0
	RGB_INTEGER                                                = 0x8D98
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	CONSTANT_COLOR                                             = 0x8001
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	DRAW_BUFFER4_ATI                                           = 0x8829
	TESS_EVALUATION_SHADER                                     = 0x8E87
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	TEXTURE26_ARB                                              = 0x84DA
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	TRANSFORM_FEEDBACK                                         = 0x8E22
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	SIGNED_ALPHA8_NV                                           = 0x8706
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	STENCIL_INDEX16_EXT                                        = 0x8D49
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	LINE                                                       = 0x1B01
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	LOGIC_OP_MODE                                              = 0x0BF0
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	FRAGMENT_DEPTH                                             = 0x8452
	TEXTURE23                                                  = 0x84D7
	SOURCE0_ALPHA_ARB                                          = 0x8588
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	VERTEX_ARRAY_EXT                                           = 0x8074
	RGB16F                                                     = 0x881B
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	MAP1_TANGENT_EXT                                           = 0x8444
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	RG_EXT                                                     = 0x8227
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	DSDT_MAG_VIB_NV                                            = 0x86F7
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	COMPILE                                                    = 0x1300
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	FLOAT_RGBA16_NV                                            = 0x888A
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	MUL_ATI                                                    = 0x8964
	GL_8X_BIT_ATI                                              = 0x00000004
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	RGB                                                        = 0x1907
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	LUMINANCE8                                                 = 0x8040
	REFERENCE_PLANE_SGIX                                       = 0x817D
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	MODELVIEW28_ARB                                            = 0x873C
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	COMPRESSED_SRGB                                            = 0x8C48
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	SGIX_flush_raster                                          = 1
	ADD                                                        = 0x0104
	COLOR_TABLE                                                = 0x80D0
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	CONSTANT_COLOR1_NV                                         = 0x852B
	FLOAT_MAT3x2                                               = 0x8B67
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	NORMAL_MAP_EXT                                             = 0x8511
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	BUFFER_DATA_SIZE                                           = 0x9303
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	ACCUM_BLUE_BITS                                            = 0x0D5A
	INT                                                        = 0x1404
	SAMPLES_ARB                                                = 0x80A9
	FOG_FUNC_SGIS                                              = 0x812A
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	STREAM_READ                                                = 0x88E1
	SAMPLE_BUFFERS                                             = 0x80A8
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	MODELVIEW9_ARB                                             = 0x8729
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	MATRIX1_NV                                                 = 0x8631
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	MODELVIEW7_ARB                                             = 0x8727
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	CULL_FACE                                                  = 0x0B44
	COMPRESSED_RGB                                             = 0x84ED
	SOURCE0_RGB_ARB                                            = 0x8580
	MIRROR_CLAMP_EXT                                           = 0x8742
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	RGBA16UI_EXT                                               = 0x8D76
	TEXTURE_SAMPLES                                            = 0x9106
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	IS_ROW_MAJOR                                               = 0x9300
	EXT_blend_subtract                                         = 1
	SGIX_async                                                 = 1
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	NEGATIVE_Y_EXT                                             = 0x87DA
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	REG_11_ATI                                                 = 0x892C
	RASTERIZER_DISCARD                                         = 0x8C89
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	ALPHA_SNORM                                                = 0x9010
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	SAMPLER                                                    = 0x82E6
	CURRENT_FOG_COORDINATE                                     = 0x8453
	OPERAND3_RGB_NV                                            = 0x8593
	DYNAMIC_ATI                                                = 0x8761
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	LAYER_NV                                                   = 0x8DAA
	IUI_N3F_V2F_EXT                                            = 0x81AF
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	LEFT                                                       = 0x0406
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	DECR                                                       = 0x1E03
	FUNC_SUBTRACT_EXT                                          = 0x800A
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	POINT_SIZE_MAX                                             = 0x8127
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	CMYK_EXT                                                   = 0x800C
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	LUMINANCE12                                                = 0x8041
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	DRAW_BUFFER13                                              = 0x8832
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	SAMPLER_2D_SHADOW                                          = 0x8B62
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	RG8UI                                                      = 0x8238
	MIRRORED_REPEAT_OES                                        = 0x8370
	OPERAND2_RGB_ARB                                           = 0x8592
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	DRAW_BUFFER13_NV                                           = 0x8832
	TEXTURE5                                                   = 0x84C5
	OP_MOV_EXT                                                 = 0x8799
	REG_21_ATI                                                 = 0x8936
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	LO_SCALE_NV                                                = 0x870F
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	INTENSITY32I_EXT                                           = 0x8D85
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	MULT                                                       = 0x0103
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	DOUBLE_MAT4_EXT                                            = 0x8F48
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	MAP2_TANGENT_EXT                                           = 0x8445
	DRAW_BUFFER1_ARB                                           = 0x8826
	FLOAT_VEC3                                                 = 0x8B51
	INDEX_BITS                                                 = 0x0D51
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	TEXTURE29_ARB                                              = 0x84DD
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	IUI_V3F_EXT                                                = 0x81AE
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	RGB9_E5_EXT                                                = 0x8C3D
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	SAMPLES                                                    = 0x80A9
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	CURRENT_BIT                                                = 0x00000001
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	VERSION_1_3                                                = 1
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	BOOL_VEC2                                                  = 0x8B57
	FLOAT_MAT2_ARB                                             = 0x8B5A
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	EXT_texture                                                = 1
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	DRAW_BUFFER12_ATI                                          = 0x8831
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	MIRRORED_REPEAT_ARB                                        = 0x8370
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	RENDERBUFFER                                               = 0x8D41
	STEREO                                                     = 0x0C33
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	SCALAR_EXT                                                 = 0x87BE
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	NORMAL_MAP_OES                                             = 0x8511
	FIXED_ONLY_ARB                                             = 0x891D
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	VERTEX_SUBROUTINE                                          = 0x92E8
	TRIANGLES_ADJACENCY                                        = 0x000C
	RGB_SCALE                                                  = 0x8573
	SRC_COLOR                                                  = 0x0300
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	SYNC_FENCE_APPLE                                           = 0x9116
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	EXT_abgr                                                   = 1
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	R16F                                                       = 0x822D
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	SIGNED_RGBA8_NV                                            = 0x86FC
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	PROGRAM_POINT_SIZE                                         = 0x8642
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	INTENSITY16_SNORM                                          = 0x901B
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	COMBINER5_NV                                               = 0x8555
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	MATRIX24_ARB                                               = 0x88D8
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	NORMAL_MAP_ARB                                             = 0x8511
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	RGB565_OES                                                 = 0x8D62
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	DOT_PRODUCT_NV                                             = 0x86EC
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	MAX_SAMPLES_NV                                             = 0x8D57
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	VIEWPORT                                                   = 0x0BA2
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	UNSIGNED_INT8_NV                                           = 0x8FEC
	IMAGE_3D_EXT                                               = 0x904E
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	MATRIX7_NV                                                 = 0x8637
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	FASTEST                                                    = 0x1101
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	MAP2_BINORMAL_EXT                                          = 0x8447
	STREAM_COPY                                                = 0x88E2
	YCBAYCR8A_4224_NV                                          = 0x9032
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	DYNAMIC_COPY                                               = 0x88EA
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	RGBA_FLOAT16_ATI                                           = 0x881A
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	SOURCE0_RGB                                                = 0x8580
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	POLYGON                                                    = 0x0009
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	MATRIX19_ARB                                               = 0x88D3
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	AMBIENT                                                    = 0x1200
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	VERTEX_ARRAY_SIZE                                          = 0x807A
	BGR                                                        = 0x80E0
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	DOT3_ATI                                                   = 0x8966
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	SGIX_impact_pixel_texture                                  = 1
	MAX                                                        = 0x8008
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	CONVOLUTION_FORMAT                                         = 0x8017
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	READ_ONLY                                                  = 0x88B8
	PALETTE8_RGBA8_OES                                         = 0x8B96
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	TEXTURE_MAX_LEVEL                                          = 0x813D
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	TEXTURE_BIT                                                = 0x00040000
	MAX_VIEWPORTS                                              = 0x825B
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	FLOAT_RGB16_NV                                             = 0x8888
	ARRAY_BUFFER                                               = 0x8892
	LINE_TOKEN                                                 = 0x0702
	SPRITE_AXIAL_SGIX                                          = 0x814C
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	BLEND_EQUATION                                             = 0x8009
	TEXTURE28                                                  = 0x84DC
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	TRACE_NAME_MESA                                            = 0x8756
	DEBUG_ASSERT_MESA                                          = 0x875B
	SGX_BINARY_IMG                                             = 0x8C0A
	DOUBLE_MAT2x3                                              = 0x8F49
	HALF_FLOAT_NV                                              = 0x140B
	DRAW_BUFFER1                                               = 0x8826
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	INDEX_BIT_PGI                                              = 0x00080000
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	COMBINER7_NV                                               = 0x8557
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	COLOR_ARRAY                                                = 0x8076
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	TEXTURE15_ARB                                              = 0x84CF
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	DRAW_BUFFER2_ARB                                           = 0x8827
	BLUE                                                       = 0x1905
	RG8I                                                       = 0x8237
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	REG_6_ATI                                                  = 0x8927
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	COMPUTE_SUBROUTINE                                         = 0x92ED
	SGIX_shadow_ambient                                        = 1
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	UNPACK_ALIGNMENT                                           = 0x0CF5
	ALPHA                                                      = 0x1906
	PACK_IMAGE_HEIGHT                                          = 0x806C
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	TEXTURE_HEIGHT                                             = 0x1001
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	MATRIX30_ARB                                               = 0x88DE
	SKIP_COMPONENTS2_NV                                        = -5
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	LUMINANCE12_ALPHA4                                         = 0x8046
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	TESS_CONTROL_TEXTURE                                       = 0x829C
	VIEW_CLASS_8_BITS                                          = 0x82CB
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	SGIX_ycrcb                                                 = 1
	POINT_SIZE_MAX_EXT                                         = 0x8127
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	BINORMAL_ARRAY_EXT                                         = 0x843A
	MAX_VERTEX_STREAMS                                         = 0x8E71
	COPY_READ_BUFFER                                           = 0x8F36
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	DEPTH_COMPONENT16                                          = 0x81A5
	MAX_HEIGHT                                                 = 0x827F
	TEXTURE14_ARB                                              = 0x84CE
	TEXTURE22_ARB                                              = 0x84D6
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	SGIX_vertex_preclip                                        = 1
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	GL_2PASS_0_SGIS                                            = 0x80A2
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	MAX_PATCH_VERTICES                                         = 0x8E7D
	YCRCB_444_SGIX                                             = 0x81BC
	SAMPLES_EXT                                                = 0x80A9
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	AUX_BUFFERS                                                = 0x0C00
	TABLE_TOO_LARGE                                            = 0x8031
	TEXTURE24_ARB                                              = 0x84D8
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	PREVIOUS_ARB                                               = 0x8578
	INVARIANT_EXT                                              = 0x87C2
	CURRENT_PROGRAM                                            = 0x8B8D
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	SGIX_depth_texture                                         = 1
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	COLOR_MATRIX_SGI                                           = 0x80B1
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	MAP_READ_BIT_EXT                                           = 0x0001
	LIGHT1                                                     = 0x4001
	POINT_SIZE_MIN                                             = 0x8126
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	RG8_SNORM                                                  = 0x8F95
	DRAW_BUFFER_EXT                                            = 0x0C01
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	COLOR_ATTACHMENT15                                         = 0x8CEF
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	TEXTURE13_ARB                                              = 0x84CD
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	TEXTURE_WRAP_S                                             = 0x2802
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	MATRIX15_ARB                                               = 0x88CF
	EXT_texture3D                                              = 1
	NICEST                                                     = 0x1102
	OP_POWER_EXT                                               = 0x8793
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	TEXTURE_GREEN_SIZE                                         = 0x805D
	VARIABLE_A_NV                                              = 0x8523
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	INT_10_10_10_2_OES                                         = 0x8DF7
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	CON_20_ATI                                                 = 0x8955
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	SHADER_CONSISTENT_NV                                       = 0x86DD
	DRAW_BUFFER0                                               = 0x8825
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	FLOAT_MAT4x3                                               = 0x8B6A
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	MATRIX_EXT                                                 = 0x87C0
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	STENCIL_EXT                                                = 0x1802
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	SIGNED_RGBA_NV                                             = 0x86FB
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	POINT_BIT                                                  = 0x00000002
	DRAW_BUFFER0_NV                                            = 0x8825
	CON_26_ATI                                                 = 0x895B
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	CULL_FRAGMENT_NV                                           = 0x86E7
	FULL_RANGE_EXT                                             = 0x87E1
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	PROGRAM                                                    = 0x82E2
	BUFFER_ACCESS_ARB                                          = 0x88BB
	BLUE_INTEGER                                               = 0x8D96
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	PROGRAM_BINDING_ARB                                        = 0x8677
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	AUX0                                                       = 0x0409
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	INTENSITY32F_ARB                                           = 0x8817
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	PROGRAM_OUTPUT                                             = 0x92E4
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	LOWER_LEFT                                                 = 0x8CA1
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	LOGIC_OP                                                   = 0x0BF1
	MAP1_INDEX                                                 = 0x0D91
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	HALF_BIT_ATI                                               = 0x00000008
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	INT16_VEC2_NV                                              = 0x8FE5
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	DOT3_RGB                                                   = 0x86AE
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	LIGHT7                                                     = 0x4007
	RG8                                                        = 0x822B
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	INT_SAMPLER_1D                                             = 0x8DC9
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	SGIS_texture_edge_clamp                                    = 1
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	DRAW_BUFFER2                                               = 0x8827
	ACTIVE_UNIFORMS                                            = 0x8B86
	CON_17_ATI                                                 = 0x8952
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	MATRIX9_ARB                                                = 0x88C9
	INT_VEC2                                                   = 0x8B53
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	DECAL                                                      = 0x2101
	TEXTURE7                                                   = 0x84C7
	VERTEX_STREAM4_ATI                                         = 0x8770
	CON_7_ATI                                                  = 0x8948
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	INCR_WRAP                                                  = 0x8507
	GL_3DC_XY_AMD                                              = 0x87FA
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	LINE_BIT                                                   = 0x00000004
	RED                                                        = 0x1903
	CUBIC_HP                                                   = 0x815F
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	GREEN_INTEGER                                              = 0x8D95
	ONE_MINUS_DST_COLOR                                        = 0x0307
	PACK_ROW_LENGTH                                            = 0x0D02
	C4F_N3F_V3F                                                = 0x2A26
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	TEXTURE_MIN_LOD                                            = 0x813A
	TRACK_MATRIX_NV                                            = 0x8648
	FLOAT_RGB32_NV                                             = 0x8889
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	OP_SET_LT_EXT                                              = 0x878D
	MATRIX28_ARB                                               = 0x88DC
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	PRIMITIVES_GENERATED                                       = 0x8C87
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	RGBA16_SNORM                                               = 0x8F9B
	MINOR_VERSION                                              = 0x821C
	SIGNED_RGB_NV                                              = 0x86FE
	OP_CLAMP_EXT                                               = 0x878E
	STENCIL_INDEX8_EXT                                         = 0x8D48
	INTENSITY16I_EXT                                           = 0x8D8B
	INT8_NV                                                    = 0x8FE0
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	ZERO                                                       = 0
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	BLEND_DST                                                  = 0x0BE0
	GL_4PASS_0_SGIS                                            = 0x80A4
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	DOT3_RGBA_IMG                                              = 0x86AF
	DRAW_BUFFER11                                              = 0x8830
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	REPEAT                                                     = 0x2901
	RGBA_INTEGER_EXT                                           = 0x8D99
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	FRONT_FACE                                                 = 0x0B46
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	TEXTURE_COMPARE_MODE                                       = 0x884C
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	BUMP_ENVMAP_ATI                                            = 0x877B
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	VARIABLE_F_NV                                              = 0x8528
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	LUMINANCE16_EXT                                            = 0x8042
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	GL_4PASS_3_EXT                                             = 0x80A7
	MAX_LABEL_LENGTH                                           = 0x82E8
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	OP_RECIP_EXT                                               = 0x8794
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	QUERY_RESULT_ARB                                           = 0x8866
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	FRAME_NV                                                   = 0x8E26
	BOLD_BIT_NV                                                = 0x01
	DEBUG_TYPE_OTHER                                           = 0x8251
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	UNPACK_LSB_FIRST                                           = 0x0CF1
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	RGBA32I_EXT                                                = 0x8D82
	GREEN                                                      = 0x1904
	CLAMP_TO_BORDER                                            = 0x812D
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	COMBINER_SCALE_NV                                          = 0x8548
	QUERY_RESULT_EXT                                           = 0x8866
	RED_BIT_ATI                                                = 0x00000001
	CONTINUOUS_AMD                                             = 0x9007
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	CCW                                                        = 0x0901
	COMPILE_AND_EXECUTE                                        = 0x1301
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	DEBUG_TYPE_MARKER                                          = 0x8268
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	SAMPLER_3D                                                 = 0x8B5F
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	BLUE_SCALE                                                 = 0x0D1A
	COMPRESSED_LUMINANCE                                       = 0x84EA
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	MODELVIEW13_ARB                                            = 0x872D
	STATIC_READ                                                = 0x88E5
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	FOG_BIT                                                    = 0x00000080
	MAP1_COLOR_4                                               = 0x0D90
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	CLAMP_TO_EDGE                                              = 0x812F
	SGIS_point_line_texgen                                     = 1
	NAND                                                       = 0x150E
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	DEBUG_SOURCE_OTHER                                         = 0x824B
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	DRAW_BUFFER9_ARB                                           = 0x882E
	EXT_shared_texture_palette                                 = 1
	ALPHA4_EXT                                                 = 0x803B
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	GL_2D                                                      = 0x0600
	LINEAR_ATTENUATION                                         = 0x1208
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	PRIMITIVE_RESTART                                          = 0x8F9D
	INDEX_MODE                                                 = 0x0C30
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	STENCIL_INDEX4                                             = 0x8D47
	INT_IMAGE_BUFFER                                           = 0x905C
	STENCIL_VALUE_MASK                                         = 0x0B93
	REDUCE                                                     = 0x8016
	DOUBLE_EXT                                                 = 0x140A
	HALF_APPLE                                                 = 0x140B
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	MODELVIEW12_ARB                                            = 0x872C
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	DRAW_BUFFER15_NV                                           = 0x8834
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	IMAGE_BINDING_NAME                                         = 0x8F3A
	BACK                                                       = 0x0405
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	RGBA_FLOAT32_APPLE                                         = 0x8814
	CON_11_ATI                                                 = 0x894C
	FLOAT_MAT3_ARB                                             = 0x8B5B
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	MAX_VARYING_VECTORS                                        = 0x8DFC
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	SAMPLER_BUFFER_AMD                                         = 0x9001
	LUMINANCE8_EXT                                             = 0x8040
	ELEMENT_ARRAY_ATI                                          = 0x8768
	INT_VEC3_ARB                                               = 0x8B54
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	SMALL_CW_ARC_TO_NV                                         = 0x14
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	POINT_SIZE                                                 = 0x0B11
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	FOG_COORD                                                  = 0x8451
	COMBINE_ALPHA_EXT                                          = 0x8572
	PRESERVE_ATI                                               = 0x8762
	DRAW_BUFFER3_NV                                            = 0x8828
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	READ_FRAMEBUFFER                                           = 0x8CA8
	LUMINANCE16F_ARB                                           = 0x881E
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	OP_MADD_EXT                                                = 0x8788
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	EXT_blend_minmax                                           = 1
	INDEX_WRITEMASK                                            = 0x0C21
	NORMAL_MAP                                                 = 0x8511
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	MODELVIEW23_ARB                                            = 0x8737
	CON_29_ATI                                                 = 0x895E
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	RGBA_INTEGER                                               = 0x8D99
	READ_PIXELS                                                = 0x828C
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	RED_SNORM                                                  = 0x8F90
	HISTOGRAM_RED_SIZE                                         = 0x8028
	MINMAX_FORMAT                                              = 0x802F
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	COLOR_INDEXES                                              = 0x1603
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	FRAMEZOOM_SGIX                                             = 0x818B
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	YCBYCR8_422_NV                                             = 0x9031
	TEXTURE_BINDING_2D                                         = 0x8069
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	RGBA2_EXT                                                  = 0x8055
	COMBINER_MUX_SUM_NV                                        = 0x8547
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	EYE_PLANE                                                  = 0x2502
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	SGIX_texture_scale_bias                                    = 1
	R32F                                                       = 0x822E
	ADD_SIGNED_EXT                                             = 0x8574
	CON_10_ATI                                                 = 0x894B
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	SRC0_RGB                                                   = 0x8580
	TIME_ELAPSED                                               = 0x88BF
	BOOL_VEC4_ARB                                              = 0x8B59
	MAX_SAMPLES_EXT                                            = 0x8D57
	MIRRORED_REPEAT                                            = 0x8370
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	TESSELLATION_MODE_AMD                                      = 0x9004
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	ATC_RGB_AMD                                                = 0x8C92
	COMBINE_ALPHA                                              = 0x8572
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	DRAW_BUFFER13_ATI                                          = 0x8832
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	FOG_COORD_ARRAY                                            = 0x8457
	WAIT_FAILED                                                = 0x911D
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	ALPHA4                                                     = 0x803B
	OBJECT_POINT_SGIS                                          = 0x81F5
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	DECR_WRAP                                                  = 0x8508
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	TEXTURE_PRIORITY                                           = 0x8066
	GL_4PASS_0_EXT                                             = 0x80A4
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	COUNTER_RANGE_AMD                                          = 0x8BC1
	MODELVIEW10_ARB                                            = 0x872A
	OP_NEGATE_EXT                                              = 0x8783
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	LUMINANCE16UI_EXT                                          = 0x8D7A
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	LESS                                                       = 0x0201
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	DRAW_BUFFER7_ARB                                           = 0x882C
	LUMINANCE8I_EXT                                            = 0x8D92
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	QUERY_NO_WAIT                                              = 0x8E14
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	UNSIGNED_INT16_NV                                          = 0x8FF0
	FAILURE_NV                                                 = 0x9030
	GL_4PASS_1_SGIS                                            = 0x80A5
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	INDEX                                                      = 0x8222
	COMBINER3_NV                                               = 0x8553
	REG_13_ATI                                                 = 0x892E
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	SOURCE2_RGB_EXT                                            = 0x8582
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	DRAW_BUFFER9                                               = 0x882E
	RGB9_E5                                                    = 0x8C3D
	SGIS_texture_filter4                                       = 1
	NORMAL_ARRAY_POINTER                                       = 0x808F
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	PROXY_COLOR_TABLE                                          = 0x80D3
	COLOR_SUM_EXT                                              = 0x8458
	STATIC_DRAW_ARB                                            = 0x88E4
	TEXTURE_BUFFER                                             = 0x8C2A
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	MATRIX_STRIDE                                              = 0x92FF
	STENCIL_RENDERABLE                                         = 0x8288
	STATIC_DRAW                                                = 0x88E4
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	PERTURB_EXT                                                = 0x85AE
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	PALETTE8_RGB8_OES                                          = 0x8B95
	CURRENT_TIME_NV                                            = 0x8E28
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	PACK_RESAMPLE_OML                                          = 0x8984
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	STENCIL_INDEX8_OES                                         = 0x8D48
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	STENCIL_INDEX                                              = 0x1901
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	COLOR_ATTACHMENT12                                         = 0x8CEC
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	ALPHA_SCALE                                                = 0x0D1C
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	RGBA8                                                      = 0x8058
	CULL_VERTEX_EXT                                            = 0x81AA
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	SOURCE1_RGB_EXT                                            = 0x8581
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	DOT3_RGBA_ARB                                              = 0x86AF
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	TEXTURE_WIDTH                                              = 0x1000
	DT_SCALE_NV                                                = 0x8711
	MAX_DRAW_BUFFERS                                           = 0x8824
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	SURFACE_REGISTERED_NV                                      = 0x86FD
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	SGIX_scalebias_hint                                        = 1
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	SAMPLE_COVERAGE                                            = 0x80A0
	TEXTURE6_ARB                                               = 0x84C6
	CON_18_ATI                                                 = 0x8953
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	CLIP_PLANE3                                                = 0x3003
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	VERTEX_SHADER_BIT                                          = 0x00000001
	CLIP_DISTANCE1                                             = 0x3001
	BLEND_COLOR_EXT                                            = 0x8005
	RGB16_EXT                                                  = 0x8054
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	GL_2X_BIT_ATI                                              = 0x00000001
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	COLOR                                                      = 0x1800
	TEXTURE_GEN_MODE                                           = 0x2500
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	DISCARD_ATI                                                = 0x8763
	ETC1_RGB8_OES                                              = 0x8D64
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	ALPHA16_SNORM                                              = 0x9018
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	PRIMITIVE_ID_NV                                            = 0x8C7C
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	INT_IMAGE_2D_RECT                                          = 0x905A
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	REPLACE_EXT                                                = 0x8062
	IMAGE_SCALE_Y_HP                                           = 0x8156
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	FLOAT_MAT3                                                 = 0x8B5B
	QUERY_BUFFER_AMD                                           = 0x9192
	POLYGON_SMOOTH                                             = 0x0B41
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	COMPILE_STATUS                                             = 0x8B81
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	COLOR_TABLE_SCALE                                          = 0x80D6
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	GREEN_BIAS                                                 = 0x0D19
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	RG16UI                                                     = 0x823A
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	LARGE_CW_ARC_TO_NV                                         = 0x18
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	DEPTH_COMPONENTS                                           = 0x8284
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	HI_SCALE_NV                                                = 0x870E
	SRGB                                                       = 0x8C40
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	DEPTH_WRITEMASK                                            = 0x0B72
	RGBA_FLOAT32_ATI                                           = 0x8814
	CON_12_ATI                                                 = 0x894D
	ALPHA_TEST                                                 = 0x0BC0
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	RGB16F_EXT                                                 = 0x881B
	SAMPLE_POSITION                                            = 0x8E50
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	ALWAYS                                                     = 0x0207
	POLYGON_STIPPLE                                            = 0x0B42
	NORMAL_ARRAY_TYPE                                          = 0x807E
	FUNC_SUBTRACT_OES                                          = 0x800A
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	PURGEABLE_APPLE                                            = 0x8A1D
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	AND_INVERTED                                               = 0x1504
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	TEXTURE22                                                  = 0x84D6
	TEXTURE16_ARB                                              = 0x84D0
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	MAX_SUBROUTINES                                            = 0x8DE7
	SGIX_interlace                                             = 1
	RGB2_EXT                                                   = 0x804E
	FLOAT_MAT4x2                                               = 0x8B69
	COLOR_ATTACHMENT2                                          = 0x8CE2
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	SPRITE_SGIX                                                = 0x8148
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	DOUBLE_VEC4                                                = 0x8FFE
	PROJECTION                                                 = 0x1701
	QUAD_ALPHA8_SGIS                                           = 0x811F
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	INDEX_ARRAY_STRIDE                                         = 0x8086
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	SAMPLER_2D_ARB                                             = 0x8B5E
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	GEOMETRY_TEXTURE                                           = 0x829E
	COMBINER6_NV                                               = 0x8556
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	DEPTH_BOUNDS_EXT                                           = 0x8891
	TEXTURE_RESIDENT                                           = 0x8067
	ALPHA16F_ARB                                               = 0x881C
	CLAMP_VERTEX_COLOR                                         = 0x891A
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	BLEND_DST_ALPHA                                            = 0x80CA
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	COLOR_INDEX4_EXT                                           = 0x80E4
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	R8_EXT                                                     = 0x8229
	ALPHA_FLOAT16_ATI                                          = 0x881C
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	SRC1_COLOR                                                 = 0x88F9
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	ACCUM                                                      = 0x0100
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	SGIX_ir_instrument1                                        = 1
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	OP_FLOOR_EXT                                               = 0x878F
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	LINE_SMOOTH                                                = 0x0B20
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	COMPRESSED_RGB_ARB                                         = 0x84ED
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	BLEND_SRC_RGB                                              = 0x80C9
	SAMPLE_MASK_NV                                             = 0x8E51
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	TEXTURE27_ARB                                              = 0x84DB
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	INDEX_ARRAY_POINTER                                        = 0x8091
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	DRAW_BUFFER0_ARB                                           = 0x8825
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	CON_2_ATI                                                  = 0x8943
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	PATCH_VERTICES                                             = 0x8E72
	OR                                                         = 0x1507
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	EXT_polygon_offset                                         = 1
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	OPERAND2_ALPHA                                             = 0x859A
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	WEIGHT_ARRAY_OES                                           = 0x86AD
	LUMINANCE32F_ARB                                           = 0x8818
	REG_7_ATI                                                  = 0x8928
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	VERTEX_ARRAY_POINTER                                       = 0x808E
	DEPTH_BITS                                                 = 0x0D56
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	FOG_COORDINATE_EXT                                         = 0x8451
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	CULL_VERTEX_IBM                                            = 103050
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	RGB16F_ARB                                                 = 0x881B
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	SYNC_CONDITION                                             = 0x9113
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	VERTEX_STREAM2_ATI                                         = 0x876E
	IMAGE_3D                                                   = 0x904E
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	ZOOM_Y                                                     = 0x0D17
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	TEXTURE_3D_OES                                             = 0x806F
	BLEND_DST_RGB                                              = 0x80C8
	R32UI                                                      = 0x8236
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	ALPHA8                                                     = 0x803C
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	TRACE_MASK_MESA                                            = 0x8755
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	DISPLAY_LIST                                               = 0x82E7
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	GL_4PASS_3_SGIS                                            = 0x80A7
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	ADD_SIGNED                                                 = 0x8574
	VIBRANCE_SCALE_NV                                          = 0x8713
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	RGB16UI_EXT                                                = 0x8D77
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	MAP1_NORMAL                                                = 0x0D92
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	SHADER_IMAGE_STORE                                         = 0x82A5
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	RENDERBUFFER_BINDING                                       = 0x8CA7
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	SYNC_CONDITION_APPLE                                       = 0x9113
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	SRGB8_EXT                                                  = 0x8C41
	LOW_INT                                                    = 0x8DF3
	TRIANGLES                                                  = 0x0004
	FIXED_OES                                                  = 0x140C
	POINT_SIZE_MAX_ARB                                         = 0x8127
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	OP_ADD_EXT                                                 = 0x8787
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	MAP2_VERTEX_3                                              = 0x0DB7
	REFLECTION_MAP_NV                                          = 0x8512
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	TEXTURE18_ARB                                              = 0x84D2
	DRAW_BUFFER3_ATI                                           = 0x8828
	DRAW_BUFFER7_ATI                                           = 0x882C
	MAX_TEXTURE_SIZE                                           = 0x0D33
	TEXTURE_3D_EXT                                             = 0x806F
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	PALETTE8_RGBA4_OES                                         = 0x8B98
	DIFFUSE                                                    = 0x1201
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	INVERSE_NV                                                 = 0x862B
	WRITE_ONLY_ARB                                             = 0x88B9
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	LIGHTING_BIT                                               = 0x00000040
	FLOAT_RG32_NV                                              = 0x8887
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	STACK_OVERFLOW                                             = 0x0503
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	GENERATE_MIPMAP_HINT                                       = 0x8192
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	SIGNED_RGB8_NV                                             = 0x86FF
	DRAW_BUFFER5_ARB                                           = 0x882A
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	POINT_SMOOTH                                               = 0x0B10
	DOUBLE_MAT3_EXT                                            = 0x8F47
	IMAGE_1D                                                   = 0x904C
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	PACK_LSB_FIRST                                             = 0x0D01
	VIEW_CLASS_24_BITS                                         = 0x82C9
	REG_28_ATI                                                 = 0x893D
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	POINT_SMOOTH_HINT                                          = 0x0C51
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	RED_BITS                                                   = 0x0D52
	CONSTANT_COLOR_EXT                                         = 0x8001
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	DOT3_RGBA                                                  = 0x86AF
	REG_0_ATI                                                  = 0x8921
	SKIP_COMPONENTS1_NV                                        = -6
	INT64_VEC4_NV                                              = 0x8FEB
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	ATTENUATION_EXT                                            = 0x834D
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	SHADE_MODEL                                                = 0x0B54
	TRIANGLE_LIST_SUN                                          = 0x81D7
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	UNSIGNALED_APPLE                                           = 0x9118
	BLEND                                                      = 0x0BE2
	RGBA4_S3TC                                                 = 0x83A3
	DRAW_BUFFER14                                              = 0x8833
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	HALF_FLOAT                                                 = 0x140B
	ADD_SIGNED_ARB                                             = 0x8574
	RGB10_A2_EXT                                               = 0x8059
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	CLEAR_BUFFER                                               = 0x82B4
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	LUMINANCE4                                                 = 0x803F
	DELETE_STATUS                                              = 0x8B80
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	INTENSITY12_EXT                                            = 0x804C
	SHADER_IMAGE_LOAD                                          = 0x82A4
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	INT8_VEC4_NV                                               = 0x8FE3
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	SIGNED_HILO8_NV                                            = 0x885F
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	INT_SAMPLER_3D                                             = 0x8DCB
	UNPACK_SKIP_IMAGES                                         = 0x806D
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	CURRENT_ATTRIB_NV                                          = 0x8626
	ADD_BLEND_IMG                                              = 0x8C09
	FUNC_ADD                                                   = 0x8006
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	TEXTURE2                                                   = 0x84C2
	TEXTURE24                                                  = 0x84D8
	HIGH_FLOAT                                                 = 0x8DF2
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	DOUBLE_MAT3x2                                              = 0x8F4B
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	MAX_LAYERS                                                 = 0x8281
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	SWIZZLE_STRQ_ATI                                           = 0x897A
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	COLOR_CLEAR_VALUE                                          = 0x0C22
	SET                                                        = 0x150F
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	DEPTH_STENCIL_NV                                           = 0x84F9
	VARIABLE_D_NV                                              = 0x8526
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	RGBA8I                                                     = 0x8D8E
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	INT_IMAGE_3D_EXT                                           = 0x9059
	SGIX_icc_texture                                           = 1
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	QUAD_MESH_SUN                                              = 0x8614
	OP_INDEX_EXT                                               = 0x8782
	MATRIX10_ARB                                               = 0x88CA
	SAMPLER_CUBE_ARB                                           = 0x8B60
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	ARB_imaging                                                = 1
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	SUBTRACT_ARB                                               = 0x84E7
	STORAGE_CACHED_APPLE                                       = 0x85BE
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	MAX_LIGHTS                                                 = 0x0D31
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	DOT2_ADD_ATI                                               = 0x896C
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	COUNT_UP_NV                                                = 0x9088
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	DST_COLOR                                                  = 0x0306
	FILTER                                                     = 0x829A
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	T4F_V4F                                                    = 0x2A28
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	PATH_DASH_OFFSET_NV                                        = 0x907E
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	MODELVIEW20_ARB                                            = 0x8734
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	DRAW_BUFFER10_ARB                                          = 0x882F
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	DOUBLE_MAT2_EXT                                            = 0x8F46
	SIGNED_NORMALIZED                                          = 0x8F9C
	HILO_NV                                                    = 0x86F4
	SCREEN_COORDINATES_REND                                    = 0x8490
	SAMPLE_MASK_VALUE                                          = 0x8E52
	EXT_subtexture                                             = 1
	COLOR_SUM_ARB                                              = 0x8458
	PROGRAM_RESIDENT_NV                                        = 0x8647
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	DYNAMIC_COPY_ARB                                           = 0x88EA
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	MATRIX6_NV                                                 = 0x8636
	BGR_INTEGER_EXT                                            = 0x8D9A
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	READ_BUFFER_EXT                                            = 0x0C02
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	TEXTURE_MAX_LOD                                            = 0x813B
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	MATRIX16_ARB                                               = 0x88D0
	NUM_PASSES_ATI                                             = 0x8970
	ALPHA_INTEGER_EXT                                          = 0x8D97
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	COLOR_ARRAY_LIST_IBM                                       = 103072
	SGIX_texture_add_env                                       = 1
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	MAP_STENCIL                                                = 0x0D11
	MODELVIEW26_ARB                                            = 0x873A
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	FUNC_ADD_EXT                                               = 0x8006
	VIEW_CLASS_16_BITS                                         = 0x82CA
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	CUBIC_CURVE_TO_NV                                          = 0x0C
	VERTEX_ARRAY_TYPE                                          = 0x807B
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	COMBINE_EXT                                                = 0x8570
	R11F_G11F_B10F                                             = 0x8C3A
	PROGRAM_INPUT                                              = 0x92E3
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	COMBINE_RGB                                                = 0x8571
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	NORMAL_MAP_NV                                              = 0x8511
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	ARRAY_SIZE                                                 = 0x92FB
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	UNIFORM_TYPE                                               = 0x8A37
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	EXT_convolution                                            = 1
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	INT_VEC4                                                   = 0x8B55
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	ALPHA_TEST_FUNC                                            = 0x0BC1
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	STENCIL_INDEX16                                            = 0x8D49
	LINES_ADJACENCY_EXT                                        = 0x000A
	MAP_TESSELLATION_NV                                        = 0x86C2
	SLIM10U_SGIX                                               = 0x831E
	DS_BIAS_NV                                                 = 0x8716
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	FENCE_APPLE                                                = 0x8A0B
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	DYNAMIC_READ                                               = 0x88E9
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	TEXTURE_2D                                                 = 0x0DE1
	TEXTURE0_ARB                                               = 0x84C0
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	COLOR_MATERIAL_FACE                                        = 0x0B55
	BLUE_BITS                                                  = 0x0D54
	CONVOLUTION_2D_EXT                                         = 0x8011
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	CLIP_DISTANCE7                                             = 0x3007
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	RED_MAX_CLAMP_INGR                                         = 0x8564
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	MATRIX29_ARB                                               = 0x88DD
	CLIP_PLANE4                                                = 0x3004
	SRGB8_NV                                                   = 0x8C41
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	LUMINANCE_ALPHA                                            = 0x190A
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	SHADER_OBJECT_ARB                                          = 0x8B48
	MIN_LOD_WARNING_AMD                                        = 0x919C
	LINE_WIDTH_RANGE                                           = 0x0B22
	SIGNED_ALPHA_NV                                            = 0x8705
	SGIS_generate_mipmap                                       = 1
	EXP2                                                       = 0x0801
	Q                                                          = 0x2003
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	VERSION_3_1                                                = 1
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	ALPHA_FLOAT32_ATI                                          = 0x8816
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	COEFF                                                      = 0x0A00
	SAMPLE_MASK_EXT                                            = 0x80A0
	W_EXT                                                      = 0x87D8
	DRAW_BUFFER8                                               = 0x882D
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	DEBUG_TYPE_ERROR                                           = 0x824C
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	MATRIX31_ARB                                               = 0x88DF
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	MAX_IMAGE_SAMPLES                                          = 0x906D
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	BUFFER_USAGE_ARB                                           = 0x8765
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	RGB10                                                      = 0x8052
	GL_2PASS_1_SGIS                                            = 0x80A3
	VARIANT_EXT                                                = 0x87C1
	BUFFER_MAP_OFFSET                                          = 0x9121
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	CONSTANT                                                   = 0x8576
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	INT_VEC3                                                   = 0x8B54
	PACK_CMYK_HINT_EXT                                         = 0x800E
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	SAMPLES_SGIS                                               = 0x80A9
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	QUERY                                                      = 0x82E3
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	SUBTRACT                                                   = 0x84E7
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	TEXTURE_RED_TYPE                                           = 0x8C10
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	OPERAND0_ALPHA_ARB                                         = 0x8598
	OUTPUT_VERTEX_EXT                                          = 0x879A
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	GL_4PASS_2_EXT                                             = 0x80A6
	SGIX_texture_multi_buffer                                  = 1
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	FLOAT16_VEC2_NV                                            = 0x8FF9
	MULTIVIEW_EXT                                              = 0x90F1
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	PROGRAM_PARAMETER_NV                                       = 0x8644
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	TEXTURE_1D                                                 = 0x0DE0
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	OP_MAX_EXT                                                 = 0x878A
	DOUBLEBUFFER                                               = 0x0C32
	PATH_DASH_CAPS_NV                                          = 0x907B
	BUFFER_BINDING                                             = 0x9302
	STENCIL_BITS                                               = 0x0D57
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	OUTPUT_COLOR1_EXT                                          = 0x879C
	MODELVIEW0_EXT                                             = 0x1700
	MATRIX21_ARB                                               = 0x88D5
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	SLUMINANCE_EXT                                             = 0x8C46
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	COLOR_ARRAY_EXT                                            = 0x8076
	UNIFORM_BUFFER_START                                       = 0x8A29
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	COMPRESSED_RG                                              = 0x8226
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	RGB_SCALE_ARB                                              = 0x8573
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	RGBA32UI_EXT                                               = 0x8D70
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	INTERLACE_OML                                              = 0x8980
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	NORMALIZE                                                  = 0x0BA1
	ALPHA_TEST_REF                                             = 0x0BC2
	MODELVIEW15_ARB                                            = 0x872F
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	COLOR_TABLE_SGI                                            = 0x80D0
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	DRAW_BUFFER9_ATI                                           = 0x882E
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	SYNC_FLAGS_APPLE                                           = 0x9115
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	TEXTURE_MATRIX                                             = 0x0BA8
	TANGENT_ARRAY_EXT                                          = 0x8439
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	MODELVIEW8_ARB                                             = 0x8728
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	RGB16                                                      = 0x8054
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	BLEND_EQUATION_RGB                                         = 0x8009
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	REG_1_ATI                                                  = 0x8922
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	T2F_IUI_V3F_EXT                                            = 0x81B2
	OPERAND2_RGB_EXT                                           = 0x8592
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	DOUBLE_VEC3                                                = 0x8FFD
	PACK_ALIGNMENT                                             = 0x0D05
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	VERTICAL_LINE_TO_NV                                        = 0x08
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	OCCLUSION_TEST_HP                                          = 0x8165
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	REG_14_ATI                                                 = 0x892F
	RGB8I_EXT                                                  = 0x8D8F
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	T2F_C4UB_V3F                                               = 0x2A29
	IDENTITY_NV                                                = 0x862A
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	PIXEL_COUNT_NV                                             = 0x8866
	FLOAT_R16_NV                                               = 0x8884
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	LUMINANCE16                                                = 0x8042
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	FACTOR_MIN_AMD                                             = 0x901C
	PROXY_TEXTURE_3D                                           = 0x8070
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	SRGB_ALPHA                                                 = 0x8C42
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	TEXTURE                                                    = 0x1702
	RESCALE_NORMAL_EXT                                         = 0x803A
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	FALSE                                                      = 0
	AUX2                                                       = 0x040B
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	VIEW_CLASS_64_BITS                                         = 0x82C6
	MAGNITUDE_SCALE_NV                                         = 0x8712
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	INDEX_TEST_EXT                                             = 0x81B5
	RG16                                                       = 0x822C
	STENCIL_INDEX8                                             = 0x8D48
	FIELD_LOWER_NV                                             = 0x9023
	DEBUG_OUTPUT                                               = 0x92E0
	CLIP_DISTANCE6                                             = 0x3006
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	REFLECTION_MAP                                             = 0x8512
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	TEXTURE17                                                  = 0x84D1
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	EVAL_BIT                                                   = 0x00010000
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	RGB12_EXT                                                  = 0x8053
	HILO16_NV                                                  = 0x86F8
	CON_8_ATI                                                  = 0x8949
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	RG16F_EXT                                                  = 0x822F
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	CON_30_ATI                                                 = 0x895F
	RETAINED_APPLE                                             = 0x8A1B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	TEXTURE_WRAP_T                                             = 0x2803
	INDEX_ARRAY_EXT                                            = 0x8077
	TEXTURE11                                                  = 0x84CB
	CON_22_ATI                                                 = 0x8957
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	DEPTH_COMPONENT24                                          = 0x81A6
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	TEXTURE_1D_ARRAY                                           = 0x8C18
	TEXTURE21                                                  = 0x84D5
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	DYNAMIC_READ_ARB                                           = 0x88E9
	VERTEX_ID_NV                                               = 0x8C7B
	BACK_LEFT                                                  = 0x0402
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	OBJECT_PLANE                                               = 0x2501
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	MATRIX13_ARB                                               = 0x88CD
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	MATRIX_MODE                                                = 0x0BA0
	OBJECT_LINE_SGIS                                           = 0x81F7
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	CLEAR                                                      = 0x1500
	TEXTURE_3D                                                 = 0x806F
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	REG_31_ATI                                                 = 0x8940
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	SINGLE_COLOR_EXT                                           = 0x81F9
	VERTEX_PROGRAM_ARB                                         = 0x8620
	OP_MUL_EXT                                                 = 0x8786
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	ALPHA_BITS                                                 = 0x0D55
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	BUFFER                                                     = 0x82E0
	PROGRAM_LENGTH_NV                                          = 0x8627
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	CURRENT_QUERY_ARB                                          = 0x8865
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	COMPRESSED_INTENSITY                                       = 0x84EC
	INT8_VEC2_NV                                               = 0x8FE1
	IMAGE_2D_RECT_EXT                                          = 0x904F
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	FOG_END                                                    = 0x0B64
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	INTENSITY_SNORM                                            = 0x9013
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	STENCIL_INDEX4_EXT                                         = 0x8D47
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	AUX1                                                       = 0x040A
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	TRUE                                                       = 1
	SPOT_DIRECTION                                             = 0x1204
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	FLOAT_MAT4_ARB                                             = 0x8B5C
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	R8                                                         = 0x8229
	VIEW_CLASS_48_BITS                                         = 0x82C7
	SKIP_COMPONENTS3_NV                                        = -4
	PATH_FILL_MODE_NV                                          = 0x9080
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	SPHERE_MAP                                                 = 0x2402
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	TEXTURE11_ARB                                              = 0x84CB
	STREAM_DRAW_ARB                                            = 0x88E0
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	ALPHA_TEST_QCOM                                            = 0x0BC0
	EQUIV                                                      = 0x1509
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	VARIANT_VALUE_EXT                                          = 0x87E4
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	GL_2PASS_0_EXT                                             = 0x80A2
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	SOURCE1_RGB_ARB                                            = 0x8581
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	RG8_EXT                                                    = 0x822B
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	REG_22_ATI                                                 = 0x8937
	BITMAP_TOKEN                                               = 0x0704
	VIEW_CLASS_128_BITS                                        = 0x82C4
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	NOOP                                                       = 0x1505
	SLUMINANCE8_EXT                                            = 0x8C47
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	POLYGON_BIT                                                = 0x00000008
	EQUAL                                                      = 0x0202
	INTENSITY16F_ARB                                           = 0x881D
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	RGBA16I_EXT                                                = 0x8D88
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	LINES_ADJACENCY_ARB                                        = 0x000A
	YCRCBA_SGIX                                                = 0x8319
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	LAYOUT_DEFAULT_INTEL                                       = 0
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	RGB4                                                       = 0x804F
	T2F_IUI_V2F_EXT                                            = 0x81B1
	CONTEXT_FLAGS                                              = 0x821E
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	MODULATE_COLOR_IMG                                         = 0x8C04
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	RGB16_SNORM                                                = 0x8F9A
	VERSION_1_2                                                = 1
	GL_3_BYTES                                                 = 0x1408
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	SEPARABLE_2D_EXT                                           = 0x8012
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	EVAL_2D_NV                                                 = 0x86C0
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	SLIM12S_SGIX                                               = 0x831F
	COMBINE                                                    = 0x8570
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	RGB_FLOAT16_APPLE                                          = 0x881B
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	TEXTURE1_ARB                                               = 0x84C1
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	EXT_blend_color                                            = 1
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	VALIDATE_STATUS                                            = 0x8B83
	T2F_V3F                                                    = 0x2A27
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	DRAW_BUFFER5                                               = 0x882A
	DRAW_BUFFER7                                               = 0x882C
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	FLOAT_R_NV                                                 = 0x8880
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	VARIABLE_B_NV                                              = 0x8524
	FIRST_TO_REST_NV                                           = 0x90AF
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	IUI_V2F_EXT                                                = 0x81AD
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	COMBINER_INPUT_NV                                          = 0x8542
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	RGB8UI                                                     = 0x8D7D
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	SGIS_texture_border_clamp                                  = 1
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	AFFINE_3D_NV                                               = 0x9094
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	TEXTURE25_ARB                                              = 0x84D9
	STREAM_DRAW                                                = 0x88E0
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	TESS_GEN_MODE                                              = 0x8E76
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	INVALID_OPERATION                                          = 0x0502
	POLYGON_OFFSET_FILL                                        = 0x8037
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	SATURATE_BIT_ATI                                           = 0x00000040
	RED_INTEGER                                                = 0x8D94
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	TEXTURE8_ARB                                               = 0x84C8
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	UNSIGNED_SHORT                                             = 0x1403
	GL_422_AVERAGE_EXT                                         = 0x80CE
	RG16F                                                      = 0x822F
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DUDV_ATI                                                   = 0x8779
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	SGIS_point_parameters                                      = 1
	NONE                                                       = 0
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	LERP_ATI                                                   = 0x8969
	COPY_INVERTED                                              = 0x150C
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	RGB16I                                                     = 0x8D89
	RGBA8I_EXT                                                 = 0x8D8E
	PIXEL_MODE_BIT                                             = 0x00000020
	LIST_BIT                                                   = 0x00020000
	RETURN                                                     = 0x0102
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	UNDEFINED_VERTEX                                           = 0x8260
	INTERLACE_READ_INGR                                        = 0x8568
	MODELVIEW6_ARB                                             = 0x8726
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	DUAL_ALPHA16_SGIS                                          = 0x8113
	UNPACK_RESAMPLE_OML                                        = 0x8985
	COMPUTE_SHADER                                             = 0x91B9
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	RGBA8_EXT                                                  = 0x8058
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	TEXTURE12_ARB                                              = 0x84CC
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	SGIX_pixel_texture                                         = 1
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	GL_3DC_X_AMD                                               = 0x87F9
	ALPHA8_SNORM                                               = 0x9014
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	UPPER_LEFT                                                 = 0x8CA2
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	INVALID_ENUM                                               = 0x0500
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	RGBA32UI                                                   = 0x8D70
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	INT_SAMPLER_CUBE                                           = 0x8DCC
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	COMBINE_ARB                                                = 0x8570
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	READ_PIXELS_TYPE                                           = 0x828E
	TEXTURE_SHADOW                                             = 0x82A1
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	VERTEX_STREAM0_ATI                                         = 0x876C
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	POINT_SIZE_MIN_EXT                                         = 0x8126
	DRAW_BUFFER6_ARB                                           = 0x882B
	REG_18_ATI                                                 = 0x8933
	SYNC_OBJECT_APPLE                                          = 0x8A53
	GL_1PASS_SGIS                                              = 0x80A1
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	DRAW_BUFFER3                                               = 0x8828
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	COMP_BIT_ATI                                               = 0x00000002
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	TEXTURE_DEPTH_EXT                                          = 0x8071
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	REG_26_ATI                                                 = 0x893B
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	INDEX_MATERIAL_EXT                                         = 0x81B8
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	CON_3_ATI                                                  = 0x8944
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	FLOAT_VEC3_ARB                                             = 0x8B51
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	COLOR_BUFFER_BIT                                           = 0x00004000
	CMYKA_EXT                                                  = 0x800D
	REPLICATE_BORDER_HP                                        = 0x8153
	REFLECTION_MAP_ARB                                         = 0x8512
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	LUMINANCE4_ALPHA4                                          = 0x8043
	COLOR_ENCODING                                             = 0x8296
	DU8DV8_ATI                                                 = 0x877A
	FLOAT_RGBA_NV                                              = 0x8883
	ARRAY_BUFFER_ARB                                           = 0x8892
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	FLOAT_VEC2_ARB                                             = 0x8B50
	RGB32I                                                     = 0x8D83
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	MAP2_NORMAL                                                = 0x0DB2
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	PROGRAM_LENGTH_ARB                                         = 0x8627
	FOG_START                                                  = 0x0B63
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	MATRIX1_ARB                                                = 0x88C1
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	MINMAX                                                     = 0x802E
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	CON_14_ATI                                                 = 0x894F
	DOUBLE_MAT4x3                                              = 0x8F4E
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	INVARIANT_VALUE_EXT                                        = 0x87EA
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	RGB5_A1_OES                                                = 0x8057
	SAMPLE_MASK                                                = 0x8E51
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	IMAGE_1D_ARRAY                                             = 0x9052
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	INDEX_TEST_REF_EXT                                         = 0x81B7
	SRGB_DECODE_ARB                                            = 0x8299
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	TEXTURE_CUBE_MAP                                           = 0x8513
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	INT8_VEC3_NV                                               = 0x8FE2
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	EXPAND_NORMAL_NV                                           = 0x8538
	LO_BIAS_NV                                                 = 0x8715
	SRC_ALPHA_SATURATE                                         = 0x0308
	FOG_DENSITY                                                = 0x0B62
	FOG_MODE                                                   = 0x0B65
	LINE_SMOOTH_HINT                                           = 0x0C52
	RGBA12                                                     = 0x805A
	SOURCE1_RGB                                                = 0x8581
	RENDERBUFFER_OES                                           = 0x8D41
	RGB8_SNORM                                                 = 0x8F96
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	VIBRANCE_BIAS_NV                                           = 0x8719
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	BUFFER_MAPPED_OES                                          = 0x88BC
	INDEX_SHIFT                                                = 0x0D12
	TEXTURE4                                                   = 0x84C4
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	TESS_CONTROL_SHADER                                        = 0x8E88
	LOAD                                                       = 0x0101
	DOMAIN                                                     = 0x0A02
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	OPERAND1_ALPHA                                             = 0x8599
	CON_28_ATI                                                 = 0x895D
	DST_ALPHA                                                  = 0x0304
	VERSION_2_1                                                = 1
	DRAW_BUFFER12_ARB                                          = 0x8831
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	RGB32UI_EXT                                                = 0x8D71
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	FOG_INDEX                                                  = 0x0B61
	COMPUTE_TEXTURE                                            = 0x82A0
	ATTACHED_SHADERS                                           = 0x8B85
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	POINT_SIZE_MIN_ARB                                         = 0x8126
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	MATRIX4_NV                                                 = 0x8634
	Z_EXT                                                      = 0x87D7
	FIXED                                                      = 0x140C
	INTENSITY12                                                = 0x804C
	STENCIL_BACK_FAIL                                          = 0x8801
	CND_ATI                                                    = 0x896A
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	MAP_WRITE_BIT_EXT                                          = 0x0002
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	CON_5_ATI                                                  = 0x8946
	COLOR_ATTACHMENT7                                          = 0x8CE7
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	MATRIX_PALETTE_OES                                         = 0x8840
	VIDEO_BUFFER_NV                                            = 0x9020
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	TIME_ELAPSED_EXT                                           = 0x88BF
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	EXT_packed_pixels                                          = 1
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	HISTOGRAM_SINK                                             = 0x802D
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	LIGHT6                                                     = 0x4006
	STENCIL_INDEX1                                             = 0x8D46
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	PRIMARY_COLOR_EXT                                          = 0x8577
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	MODULATE                                                   = 0x2100
	COMPRESSED_RGBA                                            = 0x84EE
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	RGB_SCALE_EXT                                              = 0x8573
	MATRIX3_NV                                                 = 0x8633
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	DITHER                                                     = 0x0BD0
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	DRAW_BUFFER2_ATI                                           = 0x8827
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	R8UI                                                       = 0x8232
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	ETC1_SRGB8_NV                                              = 0x88EE
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	OBJECT_TYPE_APPLE                                          = 0x9112
	LAYOUT_LINEAR_INTEL                                        = 1
	SRC_ALPHA                                                  = 0x0302
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	COMPRESSED_R11_EAC                                         = 0x9270
	COLOR3_BIT_PGI                                             = 0x00010000
	LUMINANCE4_EXT                                             = 0x803F
	PATH_GEN_MODE_NV                                           = 0x90B0
	SHADER_BINARY_DMP                                          = 0x9250
	BGR_EXT                                                    = 0x80E0
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	RED_INTEGER_EXT                                            = 0x8D94
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	TEXTURE_BORDER                                             = 0x1005
	LIGHT2                                                     = 0x4002
	DECR_WRAP_OES                                              = 0x8508
	SIGNED_INTENSITY_NV                                        = 0x8707
	GL_4X_BIT_ATI                                              = 0x00000002
	CONSTANT_ATTENUATION                                       = 0x1207
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	LINES                                                      = 0x0001
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	SOURCE2_RGB                                                = 0x8582
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	SLUMINANCE_ALPHA                                           = 0x8C44
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	CLIP_DISTANCE5                                             = 0x3005
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	EXP                                                        = 0x0800
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	SGIX_texture_lod_bias                                      = 1
	NORMAL_ARRAY_EXT                                           = 0x8075
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	SLIM8U_SGIX                                                = 0x831D
	PERFMON_RESULT_AMD                                         = 0x8BC6
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	MATRIX5_NV                                                 = 0x8635
	MODELVIEW25_ARB                                            = 0x8739
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	PATH_STROKE_MASK_NV                                        = 0x9084
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	CURRENT_INDEX                                              = 0x0B01
	STENCIL_REF                                                = 0x0B97
	REPLACE_MIDDLE_SUN                                         = 0x0002
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	NEGATE_BIT_ATI                                             = 0x00000004
	SRGB8                                                      = 0x8C41
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	MATRIX25_ARB                                               = 0x88D9
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	TEXTURE_BORDER_COLOR                                       = 0x1004
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	SAMPLES_3DFX                                               = 0x86B4
	NEGATIVE_W_EXT                                             = 0x87DC
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	SGIX_resample                                              = 1
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	TEXTURE_BASE_LEVEL                                         = 0x813C
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	FILE_NAME_NV                                               = 0x9074
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	COMBINER_MAPPING_NV                                        = 0x8543
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	SOURCE0_ALPHA                                              = 0x8588
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	OUT_OF_MEMORY                                              = 0x0505
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	SGIX_polynomial_ffd                                        = 1
	INDEX_ARRAY                                                = 0x8077
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	TEXTURE_COMPRESSED                                         = 0x86A1
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	TEXTURE14                                                  = 0x84CE
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	MEDIUM_FLOAT                                               = 0x8DF1
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	POINT_SIZE_RANGE                                           = 0x0B12
	PROXY_TEXTURE_2D                                           = 0x8064
	CLIP_DISTANCE4                                             = 0x3004
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	DRAW_BUFFER13_ARB                                          = 0x8832
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	SPOT_CUTOFF                                                = 0x1206
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	FIELDS_NV                                                  = 0x8E27
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	VERTEX_STREAM1_ATI                                         = 0x876D
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	LUMINANCE12_EXT                                            = 0x8041
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	INTERLACE_READ_OML                                         = 0x8981
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	LOCATION                                                   = 0x930E
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	STENCIL_FAIL                                               = 0x0B94
	SAMPLER_2D                                                 = 0x8B5E
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	FENCE_STATUS_NV                                            = 0x84F3
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	IMAGE_BUFFER_EXT                                           = 0x9051
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	EMBOSS_MAP_NV                                              = 0x855F
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	RGB16I_EXT                                                 = 0x8D89
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	PACK_SKIP_ROWS                                             = 0x0D03
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	NEGATIVE_X_EXT                                             = 0x87D9
	DRAW_BUFFER6                                               = 0x882B
	SAMPLER_3D_OES                                             = 0x8B5F
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	HALF_FLOAT_OES                                             = 0x8D61
	OP_DOT4_EXT                                                = 0x8785
	LOCAL_EXT                                                  = 0x87C4
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	V2F                                                        = 0x2A20
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	TEXTURE19                                                  = 0x84D3
	MOV_ATI                                                    = 0x8961
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	CLIP_PLANE1                                                = 0x3001
	CONVOLUTION_WIDTH                                          = 0x8018
	REGISTER_COMBINERS_NV                                      = 0x8522
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	S                                                          = 0x2000
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	OP_FRAC_EXT                                                = 0x8789
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	INDEX_LOGIC_OP                                             = 0x0BF1
	HALF_FLOAT_ARB                                             = 0x140B
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	PRESENT_DURATION_NV                                        = 0x8E2B
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	COMPUTE_SHADER_BIT                                         = 0x00000020
	MIN                                                        = 0x8007
	RESTART_SUN                                                = 0x0001
	CURRENT_TANGENT_EXT                                        = 0x843B
	TEXTURE2_ARB                                               = 0x84C2
	BOOL_VEC3_ARB                                              = 0x8B58
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	DOUBLE_MAT2x4                                              = 0x8F4A
	INT64_VEC3_NV                                              = 0x8FEA
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	MULTISAMPLE_BIT                                            = 0x20000000
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	RIGHT                                                      = 0x0407
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	RGB4_EXT                                                   = 0x804F
	FENCE_CONDITION_NV                                         = 0x84F4
	OP_ROUND_EXT                                               = 0x8790
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	TEXTURE9_ARB                                               = 0x84C9
	REG_25_ATI                                                 = 0x893A
	SQUARE_NV                                                  = 0x90A3
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MINMAX_EXT                                                 = 0x802E
	RG                                                         = 0x8227
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	SRC1_ALPHA                                                 = 0x8589
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	MAX_TEXTURE_UNITS                                          = 0x84E2
	CONSTANT_EXT                                               = 0x8576
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	ONE                                                        = 1
	NUM_EXTENSIONS                                             = 0x821D
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	MAGNITUDE_BIAS_NV                                          = 0x8718
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	STATIC_COPY                                                = 0x88E6
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	SECONDARY_COLOR_NV                                         = 0x852D
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	ARRAY_STRIDE                                               = 0x92FE
	VERSION_1_5                                                = 1
	COLOR_LOGIC_OP                                             = 0x0BF2
	ALPHA8_EXT                                                 = 0x803C
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	FIELD_UPPER_NV                                             = 0x9022
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	EXT_copy_texture                                           = 1
	MULTISAMPLE_ARB                                            = 0x809D
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	TEXTURE10_ARB                                              = 0x84CA
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	INT_IMAGE_3D                                               = 0x9059
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	DRAW_BUFFER4_NV                                            = 0x8829
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	INDEX_ARRAY_LIST_IBM                                       = 103073
	FLOAT                                                      = 0x1406
	RGB5                                                       = 0x8050
	C4UB_V3F                                                   = 0x2A23
	SOURCE3_RGB_NV                                             = 0x8583
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	SEPARATE_ATTRIBS                                           = 0x8C8D
	IMAGE_2D_ARRAY                                             = 0x9053
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	RG32UI                                                     = 0x823C
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	MAP2_INDEX                                                 = 0x0DB1
	INDEX_ARRAY_TYPE                                           = 0x8085
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	COMBINE_RGB_EXT                                            = 0x8571
	SIGNED_LUMINANCE_NV                                        = 0x8701
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	EMISSION                                                   = 0x1600
	T                                                          = 0x2001
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	COMBINE_ALPHA_ARB                                          = 0x8572
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	SHORT                                                      = 0x1402
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	DRAW_BUFFER5_NV                                            = 0x882A
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	SGIX_calligraphic_fragment                                 = 1
	FRONT_LEFT                                                 = 0x0400
	TEXTURE_GEN_Q                                              = 0x0C63
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	DEPTH_SCALE                                                = 0x0D1E
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	DUAL_ALPHA8_SGIS                                           = 0x8111
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	OR_INVERTED                                                = 0x150D
	V3F                                                        = 0x2A21
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	POLYGON_MODE                                               = 0x0B40
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	UTF8_NV                                                    = 0x909A
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	AND                                                        = 0x1501
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	VERTEX_SHADER                                              = 0x8B31
	COLOR_SAMPLES_NV                                           = 0x8E20
	SHININESS                                                  = 0x1601
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	INT_2_10_10_10_REV                                         = 0x8D9F
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	COLOR_ARRAY_TYPE                                           = 0x8082
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	CURRENT_QUERY                                              = 0x8865
	REG_4_ATI                                                  = 0x8925
	ALPHA8I_EXT                                                = 0x8D90
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	GREEN_BITS                                                 = 0x0D53
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	OPERAND0_ALPHA                                             = 0x8598
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	FUNC_SUBTRACT                                              = 0x800A
	DRAW_BUFFER14_ATI                                          = 0x8833
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	STENCIL                                                    = 0x1802
	RGB5_A1                                                    = 0x8057
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	TEXTURE12                                                  = 0x84CC
	TEXTURE30_ARB                                              = 0x84DE
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	PALETTE4_RGBA8_OES                                         = 0x8B91
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	SGIX_framezoom                                             = 1
	LINE_WIDTH                                                 = 0x0B21
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	R16                                                        = 0x822A
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	DS_SCALE_NV                                                = 0x8710
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	STACK_UNDERFLOW                                            = 0x0504
	DEPTH_FUNC                                                 = 0x0B74
	PROJECTION_MATRIX                                          = 0x0BA7
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	COMBINER2_NV                                               = 0x8552
	MODELVIEW22_ARB                                            = 0x8736
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	GL_4_BYTES                                                 = 0x1409
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	IMAGE_SCALE_X_HP                                           = 0x8155
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	TEXTURE4_ARB                                               = 0x84C4
	RGB32F_ARB                                                 = 0x8815
	REG_15_ATI                                                 = 0x8930
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	DUAL_ALPHA4_SGIS                                           = 0x8110
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	INTENSITY8I_EXT                                            = 0x8D91
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	INTENSITY4                                                 = 0x804A
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	RGBA_FLOAT16_APPLE                                         = 0x881A
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	ADJACENT_PAIRS_NV                                          = 0x90AE
	NEAREST                                                    = 0x2600
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	MAX_NAME_LENGTH                                            = 0x92F6
	SGIX_fog_offset                                            = 1
	GL_2PASS_1_EXT                                             = 0x80A3
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	DEBUG_OBJECT_MESA                                          = 0x8759
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	INDEX_OFFSET                                               = 0x0D13
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	VERSION                                                    = 0x1F02
	EMBOSS_LIGHT_NV                                            = 0x855D
	STATIC_ATI                                                 = 0x8760
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	DOUBLE_MAT2                                                = 0x8F46
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	INTENSITY8UI_EXT                                           = 0x8D7F
	WAIT_FAILED_APPLE                                          = 0x911D
	INTERLACE_SGIX                                             = 0x8094
	COLOR_MATRIX                                               = 0x80B1
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	STENCIL_INDEX4_OES                                         = 0x8D47
	MAP1_VERTEX_3                                              = 0x0D97
	DEPTH_CLAMP_NV                                             = 0x864F
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	MATRIX2_NV                                                 = 0x8632
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	SIGNALED                                                   = 0x9119
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	ALPHA_MAX_SGIX                                             = 0x8321
	BLUE_BIT_ATI                                               = 0x00000004
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	VERTEX_ARRAY                                               = 0x8074
	INVALID_INDEX                                              = 0xFFFFFFFF
	NORMAL_BIT_PGI                                             = 0x08000000
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	MATRIX18_ARB                                               = 0x88D2
	BOOL_VEC2_ARB                                              = 0x8B57
	CONVEX_HULL_NV                                             = 0x908B
	MATRIX7_ARB                                                = 0x88C7
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	TEXTURE_COMPONENTS                                         = 0x1003
	COLOR_INDEX16_EXT                                          = 0x80E7
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	QUERY_NO_WAIT_NV                                           = 0x8E14
	ARRAY_BUFFER_BINDING                                       = 0x8894
	CON_16_ATI                                                 = 0x8951
	SYNC_STATUS_APPLE                                          = 0x9114
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	ALPHA16UI_EXT                                              = 0x8D78
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	RGBA_MODE                                                  = 0x0C31
	ALPHA_MIN_SGIX                                             = 0x8320
	TEXTURE0                                                   = 0x84C0
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	GREEN_SCALE                                                = 0x0D18
	OPERAND1_ALPHA_EXT                                         = 0x8599
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	LIGHTING                                                   = 0x0B50
	MAX_DEPTH                                                  = 0x8280
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	FACTOR_MAX_AMD                                             = 0x901D
	SUBPIXEL_BITS                                              = 0x0D50
	LINEAR                                                     = 0x2601
	LUMINANCE12_ALPHA12                                        = 0x8047
	DEPTH_RENDERABLE                                           = 0x8287
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	OUTPUT_FOG_EXT                                             = 0x87BD
	SGIX_blend_alpha_minmax                                    = 1
	ASYNC_MARKER_SGIX                                          = 0x8329
	LINE_STIPPLE                                               = 0x0B24
	OPERAND3_ALPHA_NV                                          = 0x859B
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	COMMAND_BARRIER_BIT                                        = 0x00000040
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	PRIMITIVE_RESTART_NV                                       = 0x8558
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	R8_SNORM                                                   = 0x8F94
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	PACK_SKIP_IMAGES                                           = 0x806B
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	MAX_VARYING_FLOATS                                         = 0x8B4B
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	UNSIGNALED                                                 = 0x9118
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	RED_MIN_CLAMP_INGR                                         = 0x8560
	MOVE_TO_NV                                                 = 0x02
	LINE_STRIP                                                 = 0x0003
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	FLOAT_MAT2                                                 = 0x8B5A
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	RG16_SNORM                                                 = 0x8F99
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	PALETTE4_RGB8_OES                                          = 0x8B90
	INVALID_VALUE                                              = 0x0501
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	PATCHES                                                    = 0x000E
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	SOURCE0_ALPHA_EXT                                          = 0x8588
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	MATRIX14_ARB                                               = 0x88CE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	PACK_SWAP_BYTES                                            = 0x0D00
	TEXTURE31_ARB                                              = 0x84DF
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	UTF16_NV                                                   = 0x909B
	FLAT                                                       = 0x1D00
	OP_RECIP_SQRT_EXT                                          = 0x8795
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	NAME_LENGTH                                                = 0x92F9
	VERSION_3_2                                                = 1
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	DRAW_BUFFER11_NV                                           = 0x8830
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	MATRIX27_ARB                                               = 0x88DB
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	DRAW_BUFFER11_ATI                                          = 0x8830
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	SLUMINANCE                                                 = 0x8C46
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	SKIP_COMPONENTS4_NV                                        = -3
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	QUAD_STRIP                                                 = 0x0008
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	FILTER4_SGIS                                               = 0x8146
	TEXTURE23_ARB                                              = 0x84D7
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	RENDER                                                     = 0x1C00
	INTENSITY16                                                = 0x804D
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	MATRIX22_ARB                                               = 0x88D6
	REG_29_ATI                                                 = 0x893E
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	PATH_MITER_LIMIT_NV                                        = 0x907A
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	POINT                                                      = 0x1B00
	PHONG_WIN                                                  = 0x80EA
	SINGLE_COLOR                                               = 0x81F9
	REFLECTION_MAP_EXT                                         = 0x8512
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	ZERO_EXT                                                   = 0x87DD
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	SGIX_instruments                                           = 1
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	STATE_RESTORE                                              = 0x8BDC
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	SGIS_fog_function                                          = 1
	AUX3                                                       = 0x040C
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	BUFFER_USAGE                                               = 0x8765
	REG_2_ATI                                                  = 0x8923
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	TEXTURE13                                                  = 0x84CD
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	TEXTURE_GEN_S                                              = 0x0C60
	COMPRESSED_RED                                             = 0x8225
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	UNSIGNED_INVERT_NV                                         = 0x8537
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	BGRA_INTEGER_EXT                                           = 0x8D9B
	COLOR_ATTACHMENT11                                         = 0x8CEB
	EXT_blend_logic_op                                         = 1
	COLOR_WRITEMASK                                            = 0x0C23
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	MATRIX5_ARB                                                = 0x88C5
	MITER_REVERT_NV                                            = 0x90A7
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	SGIS_texture_lod                                           = 1
	MAP_COLOR                                                  = 0x0D10
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	RGB32F                                                     = 0x8815
	DOUBLE_MAT4x2                                              = 0x8F4D
	CULL_FACE_MODE                                             = 0x0B45
	COLOR_MATERIAL                                             = 0x0B57
	TEXTURE_MAG_FILTER                                         = 0x2800
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	SGIX_async_histogram                                       = 1
	RGBA2                                                      = 0x8055
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	PROGRAM_PIPELINE                                           = 0x82E4
	DEPTH_STENCIL_EXT                                          = 0x84F9
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	CONSTANT_ALPHA                                             = 0x8003
	EMBOSS_CONSTANT_NV                                         = 0x855E
	OPERAND1_RGB_EXT                                           = 0x8591
	CONDITION_SATISFIED                                        = 0x911C
	COLOR_INDEX12_EXT                                          = 0x80E6
	RG_INTEGER                                                 = 0x8228
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	SOURCE2_ALPHA_EXT                                          = 0x858A
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	INTERPOLATE_EXT                                            = 0x8575
	FLOAT_MAT4                                                 = 0x8B5C
	TEXTURE_BLUE_SIZE                                          = 0x805E
	FEEDBACK                                                   = 0x1C01
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	PRIMARY_COLOR                                              = 0x8577
	MATRIX3_ARB                                                = 0x88C3
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	SCISSOR_BOX                                                = 0x0C10
	MIRRORED_REPEAT_IBM                                        = 0x8370
	SAMPLER_1D                                                 = 0x8B5D
	ITALIC_BIT_NV                                              = 0x02
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	FOG_COORDINATE_SOURCE                                      = 0x8450
	BIAS_BIT_ATI                                               = 0x00000008
	VERTEX23_BIT_PGI                                           = 0x00000004
	CLIP_DISTANCE2                                             = 0x3002
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	DRAW_BUFFER12                                              = 0x8831
	STANDARD_FONT_NAME_NV                                      = 0x9072
	DRAW_PIXEL_TOKEN                                           = 0x0705
	CONVOLUTION_1D_EXT                                         = 0x8010
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	PASS_THROUGH_NV                                            = 0x86E6
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	BLEND_EQUATION_EXT                                         = 0x8009
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	MULTISAMPLE_3DFX                                           = 0x86B2
	RGBA32F_ARB                                                = 0x8814
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	TEXTURE_WRAP_R_OES                                         = 0x8072
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	POINT_SPRITE_ARB                                           = 0x8861
	COPY_WRITE_BUFFER                                          = 0x8F37
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	RECT_NV                                                    = 0xF6
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	OP_EXP_BASE_2_EXT                                          = 0x8791
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	DOUBLE_MAT3x4                                              = 0x8F4C
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	LIGHT4                                                     = 0x4004
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	COLOR_INDEX2_EXT                                           = 0x80E3
	YCRCB_SGIX                                                 = 0x8318
	SRC0_ALPHA                                                 = 0x8588
	SURFACE_STATE_NV                                           = 0x86EB
	FRAGMENT_SHADER                                            = 0x8B30
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	RGBA16_EXT                                                 = 0x805B
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	SRGB_EXT                                                   = 0x8C40
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	RGB_S3TC                                                   = 0x83A0
	INTERPOLATE                                                = 0x8575
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	T2F_N3F_V3F                                                = 0x2A2B
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	FRACTIONAL_ODD                                             = 0x8E7B
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	CON_19_ATI                                                 = 0x8954
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	RGB8I                                                      = 0x8D8F
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	POINTS                                                     = 0x0000
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	CURRENT_BINORMAL_EXT                                       = 0x843C
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	LUMINANCE8_SNORM                                           = 0x9015
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	GL_422_REV_EXT                                             = 0x80CD
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	CONSTANT_BORDER                                            = 0x8151
	AVERAGE_HP                                                 = 0x8160
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	RGB_FLOAT16_ATI                                            = 0x881B
	C3F_V3F                                                    = 0x2A24
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	BOOL                                                       = 0x8B56
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	GEOMETRY_SHADER                                            = 0x8DD9
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	SAMPLER_OBJECT_AMD                                         = 0x9155
	CONVOLUTION_2D                                             = 0x8011
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	LOW_FLOAT                                                  = 0x8DF0
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	SAMPLER_CUBE                                               = 0x8B60
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	RGB5_A1_EXT                                                = 0x8057
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	CON_9_ATI                                                  = 0x894A
	TRANSLATE_2D_NV                                            = 0x9090
	CON_31_ATI                                                 = 0x8960
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	SGIX_clipmap                                               = 1
	ACCUM_GREEN_BITS                                           = 0x0D59
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	OPERAND2_ALPHA_EXT                                         = 0x859A
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	COLOR_RENDERABLE                                           = 0x8286
	DRAW_BUFFER4_ARB                                           = 0x8829
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	COMBINER4_NV                                               = 0x8554
	BOOL_ARB                                                   = 0x8B56
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	SAMPLE_POSITION_NV                                         = 0x8E50
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	UNSIGNED_INT                                               = 0x1405
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	COLOR_ATTACHMENT6                                          = 0x8CE6
	COLOR_ATTACHMENT14                                         = 0x8CEE
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	FOG                                                        = 0x0B60
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	RG16I                                                      = 0x8239
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	MODELVIEW                                                  = 0x1700
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	MATRIX8_ARB                                                = 0x88C8
	DEPTH32F_STENCIL8                                          = 0x8CAD
	PRESENT_TIME_NV                                            = 0x8E2A
	RGB10_A2UI                                                 = 0x906F
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	FOG_COLOR                                                  = 0x0B66
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	SHADER                                                     = 0x82E1
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	RGB_FLOAT32_APPLE                                          = 0x8815
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	LUMINANCE8UI_EXT                                           = 0x8D80
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	LIGHT3                                                     = 0x4003
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	INT_IMAGE_1D_EXT                                           = 0x9057
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	COMBINER_BIAS_NV                                           = 0x8549
	SOURCE2_RGB_ARB                                            = 0x8582
	DOT3_RGBA_EXT                                              = 0x8741
	OP_DOT3_EXT                                                = 0x8784
	RELEASED_APPLE                                             = 0x8A19
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	SCISSOR_BIT                                                = 0x00080000
	LIGHT5                                                     = 0x4005
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	VERTEX_BLEND_ARB                                           = 0x86A7
	REG_30_ATI                                                 = 0x893F
	OBJECT_TYPE_ARB                                            = 0x8B4E
	TEXTURE_COORD_NV                                           = 0x8C79
	INT16_VEC3_NV                                              = 0x8FE6
	BITMAP                                                     = 0x1A00
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	MATRIX23_ARB                                               = 0x88D7
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	LINK_STATUS                                                = 0x8B82
	PATH_JOIN_STYLE_NV                                         = 0x9079
	LUMINANCE                                                  = 0x1909
	OP_MIN_EXT                                                 = 0x878B
	CON_13_ATI                                                 = 0x894E
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	COLOR_ATTACHMENT10                                         = 0x8CEA
	LINE_TO_NV                                                 = 0x04
	CULL_MODES_NV                                              = 0x86E0
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	FLOAT_RG16_NV                                              = 0x8886
	MOVE_TO_RESETS_NV                                          = 0x90B5
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	LUMINANCE32UI_EXT                                          = 0x8D74
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	R32I                                                       = 0x8235
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	INT_IMAGE_CUBE                                             = 0x905B
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	REG_27_ATI                                                 = 0x893C
	ADD_ATI                                                    = 0x8963
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	BLEND_DST_RGB_EXT                                          = 0x80C8
	STENCIL_COMPONENTS                                         = 0x8285
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	PALETTE4_RGBA4_OES                                         = 0x8B93
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	TEXTURE_WRAP_R                                             = 0x8072
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	TEXTURE9                                                   = 0x84C9
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	RGB16UI                                                    = 0x8D77
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	FLOAT16_VEC3_NV                                            = 0x8FFA
	CURRENT_NORMAL                                             = 0x0B02
	ALPHA12_EXT                                                = 0x803D
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	SYNC_FLAGS                                                 = 0x9115
	FRAMEBUFFER_BLEND                                          = 0x828B
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	RED_SCALE                                                  = 0x0D14
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	TEXTURE26                                                  = 0x84DA
	MODULATE_ADD_ATI                                           = 0x8744
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	TEXTURE16                                                  = 0x84D0
	TEXTURE31                                                  = 0x84DF
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	RESTART_PATH_NV                                            = 0xF0
	MAX_SAMPLES_IMG                                            = 0x9135
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	BUFFER_MAPPED                                              = 0x88BC
	READ_WRITE_ARB                                             = 0x88BA
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	HISTOGRAM_WIDTH                                            = 0x8026
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	SRC1_RGB                                                   = 0x8581
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	SHADER_OBJECT_EXT                                          = 0x8B48
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	TRANSFORM_BIT                                              = 0x00001000
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	DEPTH_TEXTURE_MODE                                         = 0x884B
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	RGBA16F                                                    = 0x881A
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	RGBA4                                                      = 0x8056
	BLEND_EQUATION_OES                                         = 0x8009
	PHONG_HINT_WIN                                             = 0x80EB
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	COLOR_SUM                                                  = 0x8458
	DISCARD_NV                                                 = 0x8530
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	DYNAMIC_DRAW                                               = 0x88E8
	NUM_SAMPLE_COUNTS                                          = 0x9380
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	SGIS_sharpen_texture                                       = 1
	TEXTURE_ENV                                                = 0x2300
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	SGI_color_table                                            = 1
	LINES_ADJACENCY                                            = 0x000A
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	POINT_SPRITE                                               = 0x8861
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	STENCIL_INDEX1_EXT                                         = 0x8D46
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	REPLACE                                                    = 0x1E01
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	BLUE_INTEGER_EXT                                           = 0x8D96
	PATH_GEN_COEFF_NV                                          = 0x90B1
	VIEWPORT_BIT                                               = 0x00000800
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	TEXTURE_DEPTH                                              = 0x8071
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	R16I                                                       = 0x8233
	FRAGMENT_TEXTURE                                           = 0x829F
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	FOG_COORD_SRC                                              = 0x8450
	STORAGE_SHARED_APPLE                                       = 0x85BF
	PROGRAM_TARGET_NV                                          = 0x8646
	MODELVIEW24_ARB                                            = 0x8738
	DRAW_BUFFER4                                               = 0x8829
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	POLYGON_OFFSET_EXT                                         = 0x8037
	SGIX_async_pixel                                           = 1
	SGIX_shadow                                                = 1
	EYE_POINT_SGIS                                             = 0x81F4
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	BUFFER_MAP_POINTER                                         = 0x88BD
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MODELVIEW_MATRIX                                           = 0x0BA6
	READ_BUFFER_NV                                             = 0x0C02
	SCALE_BY_FOUR_NV                                           = 0x853F
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	DEPTH_EXT                                                  = 0x1801
	GL_4PASS_2_SGIS                                            = 0x80A6
	INT_IMAGE_2D_EXT                                           = 0x9058
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	RG32I                                                      = 0x823B
	SOURCE1_ALPHA_EXT                                          = 0x8589
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	STATIC_READ_ARB                                            = 0x88E5
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	DEPTH                                                      = 0x1801
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	MAJOR_VERSION                                              = 0x821B
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	LUMINANCE32I_EXT                                           = 0x8D86
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	TRANSLATE_3D_NV                                            = 0x9091
	SGIX_tag_sample_buffer                                     = 1
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	PATH_STENCIL_REF_NV                                        = 0x90B8
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	FOG_OFFSET_SGIX                                            = 0x8198
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	PARTIAL_SUCCESS_NV                                         = 0x902E
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	SGIX_texture_coordinate_clamp                              = 1
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	TEXTURE_COORD_ARRAY                                        = 0x8078
	HISTOGRAM                                                  = 0x8024
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	GEQUAL                                                     = 0x0206
	SOURCE3_ALPHA_NV                                           = 0x858B
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	RGBA16UI                                                   = 0x8D76
	EXT_vertex_array                                           = 1
	SGI_color_matrix                                           = 1
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	DRAW_BUFFER3_ARB                                           = 0x8828
	MATRIX6_ARB                                                = 0x88C6
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	PATH_COORD_COUNT_NV                                        = 0x909E
	INTENSITY8_EXT                                             = 0x804B
	ACTIVE_PROGRAM                                             = 0x8259
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	RGBA_S3TC                                                  = 0x83A2
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	SURFACE_MAPPED_NV                                          = 0x8700
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	CON_4_ATI                                                  = 0x8945
	DOUBLE_MAT3                                                = 0x8F47
	MAX_EXT                                                    = 0x8008
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	BUFFER_SIZE_ARB                                            = 0x8764
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	TEXTURE_LOD_BIAS                                           = 0x8501
	STATIC_COPY_ARB                                            = 0x88E6
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	LIST_INDEX                                                 = 0x0B33
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	TEXTURE3                                                   = 0x84C3
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	ISOLINES                                                   = 0x8E7A
	RELATIVE_ARC_TO_NV                                         = 0xFF
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	TEXTURE10                                                  = 0x84CA
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	INT_IMAGE_2D                                               = 0x9058
	BLEND_SRC_ALPHA                                            = 0x80CB
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	SWIZZLE_STR_ATI                                            = 0x8976
	ACCUM_RED_BITS                                             = 0x0D58
	READ_PIXELS_FORMAT                                         = 0x828D
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	SGIS_multisample                                           = 1
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	DUAL_ALPHA12_SGIS                                          = 0x8112
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	MODELVIEW31_ARB                                            = 0x873F
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	LUMINANCE16_SNORM                                          = 0x9019
	DATA_BUFFER_AMD                                            = 0x9151
	BLEND_SRC_RGB_OES                                          = 0x80C9
	COMBINER1_NV                                               = 0x8551
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	FLOAT_RGB_NV                                               = 0x8882
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	MAP_WRITE_BIT                                              = 0x0002
	LEQUAL                                                     = 0x0203
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	STENCIL_BACK_REF                                           = 0x8CA3
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	TRIANGULAR_NV                                              = 0x90A5
	GENERATE_MIPMAP                                            = 0x8191
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	OUTPUT_COLOR0_EXT                                          = 0x879B
	PATH_END_CAPS_NV                                           = 0x9076
	OPERAND1_RGB_ARB                                           = 0x8591
	FLOAT_RG_NV                                                = 0x8881
	RGBA16I                                                    = 0x8D88
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	ALPHA16                                                    = 0x803E
	RGBA8_OES                                                  = 0x8058
	MATRIX26_ARB                                               = 0x88DA
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	DRAW_BUFFER8_ATI                                           = 0x882D
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	MATRIX0_NV                                                 = 0x8630
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	RGBA16                                                     = 0x805B
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	COMPRESSED_ALPHA                                           = 0x84E9
	DOT3_RGB_ARB                                               = 0x86AE
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	SLUMINANCE8                                                = 0x8C47
	ALPHA16I_EXT                                               = 0x8D8A
	FUNC_ADD_OES                                               = 0x8006
	TEXTURE_SHADER_NV                                          = 0x86DE
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	DRAW_BUFFER15_ATI                                          = 0x8834
	RENDERBUFFER_WIDTH                                         = 0x8D42
	CLIP_DISTANCE3                                             = 0x3003
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	REG_24_ATI                                                 = 0x8939
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	POINT_TOKEN                                                = 0x0701
	POSITION                                                   = 0x1203
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	OP_SUB_EXT                                                 = 0x8796
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	HI_BIAS_NV                                                 = 0x8714
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	IMAGE_BUFFER                                               = 0x9051
	STENCIL_FUNC                                               = 0x0B92
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	ALL_COMPLETED_NV                                           = 0x84F2
	OPERAND2_ALPHA_ARB                                         = 0x859A
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	VERTEX4_BIT_PGI                                            = 0x00000008
	COLOR_EXT                                                  = 0x1800
	DEBUG_PRINT_MESA                                           = 0x875A
	SAMPLER_BINDING                                            = 0x8919
	SLUMINANCE_NV                                              = 0x8C46
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	SELECT                                                     = 0x1C02
	FLOAT_RGBA32_NV                                            = 0x888B
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	SGIX_fragment_lighting                                     = 1
	TEXTURE3_ARB                                               = 0x84C3
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	GREEN_INTEGER_EXT                                          = 0x8D95
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	SHADER_OPERATION_NV                                        = 0x86DF
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	RENDERBUFFER_EXT                                           = 0x8D41
	QUADS                                                      = 0x0007
	COORD_REPLACE                                              = 0x8862
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	QUERY_OBJECT_AMD                                           = 0x9153
	SYNC_CL_EVENT_ARB                                          = 0x8240
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	PROGRAM_STRING_NV                                          = 0x8628
	DRAW_BUFFER6_ATI                                           = 0x882B
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	COLOR_ATTACHMENT8                                          = 0x8CE8
	EDGE_FLAG_ARRAY                                            = 0x8079
	REG_9_ATI                                                  = 0x892A
	DOT4_ATI                                                   = 0x8967
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	COMPRESSED_RG11_EAC                                        = 0x9272
	DEBUG_SOURCE_API                                           = 0x8246
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	FOG_HINT                                                   = 0x0C54
	EYE_LINE_SGIS                                              = 0x81F6
	TYPE                                                       = 0x92FA
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	CURRENT_RASTER_POSITION                                    = 0x0B07
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	PROGRAM_STRING_ARB                                         = 0x8628
	SRGB8_ALPHA8                                               = 0x8C43
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	GPU_ADDRESS_NV                                             = 0x8F34
	VERTEX_STREAM6_ATI                                         = 0x8772
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	COUNTER_TYPE_AMD                                           = 0x8BC0
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	INTENSITY_EXT                                              = 0x8049
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
)

type Context struct {
	context                   *C.gl11Context
	extensions                map[string]bool
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
	glc.queryExtensions()
	return glc
}
