// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// +build opengl_debug

// Package 'opengl' implements OpenGL version 1.0
package opengl

// #cgo windows LDFLAGS: -lopengl32
// #cgo linux LDFLAGS: -lGL -ldl
// #include "gl10.h"
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
		n, wantedMajor, wantedMinor, wantedRev := parseVersions("1.0")
		if n < 2 {
			fmt.Printf("OpenGL: *** JSON version parsing failed for %q ***\n", "1.0")
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
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	TEXTURE23                                                  = 0x84D7
	IMAGE_3D_EXT                                               = 0x904E
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	VERTEX_ARRAY_EXT                                           = 0x8074
	MODELVIEW18_ARB                                            = 0x8732
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	AVERAGE_EXT                                                = 0x8335
	ALPHA_FLOAT32_ATI                                          = 0x8816
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	LINE_TOKEN                                                 = 0x0702
	MINMAX                                                     = 0x802E
	BLEND_DST_RGB                                              = 0x80C8
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	DSDT8_NV                                                   = 0x8709
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	VOLATILE_APPLE                                             = 0x8A1A
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	RGB_422_APPLE                                              = 0x8A1F
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	TEXTURE_BORDER                                             = 0x1005
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	EVAL_2D_NV                                                 = 0x86C0
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	HI_BIAS_NV                                                 = 0x8714
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	SAMPLE_MASK_EXT                                            = 0x80A0
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	DRAW_BUFFER8_ARB                                           = 0x882D
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	UNSIGNALED_APPLE                                           = 0x9118
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	DSDT_MAG_NV                                                = 0x86F6
	REG_23_ATI                                                 = 0x8938
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	TEXTURE_BINDING_2D                                         = 0x8069
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	UPPER_LEFT                                                 = 0x8CA2
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	FRAGMENT_SHADER                                            = 0x8B30
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	LAYOUT_DEFAULT_INTEL                                       = 0
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	FRONT_LEFT                                                 = 0x0400
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	AVERAGE_HP                                                 = 0x8160
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	SURFACE_MAPPED_NV                                          = 0x8700
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	SGIX_texture_add_env                                       = 1
	MAX_LIGHTS                                                 = 0x0D31
	T4F_V4F                                                    = 0x2A28
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	NONE                                                       = 0
	LUMINANCE16_ALPHA16                                        = 0x8048
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	RGB12_EXT                                                  = 0x8053
	SINGLE_COLOR                                               = 0x81F9
	CURRENT_BINORMAL_EXT                                       = 0x843C
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	CURRENT_INDEX                                              = 0x0B01
	LUMINANCE16                                                = 0x8042
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	SIGNED_LUMINANCE_NV                                        = 0x8701
	DRAW_BUFFER6                                               = 0x882B
	QUARTER_BIT_ATI                                            = 0x00000010
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	QUERY                                                      = 0x82E3
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	DRAW_BUFFER9_ARB                                           = 0x882E
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	STENCIL_REF                                                = 0x0B97
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	VERTEX_STREAM3_ATI                                         = 0x876F
	VERTEX_STREAM4_ATI                                         = 0x8770
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	ENABLE_BIT                                                 = 0x00002000
	AUX3                                                       = 0x040C
	RGBA8                                                      = 0x8058
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	R8                                                         = 0x8229
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	FLOAT_RG32_NV                                              = 0x8887
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	LIGHTING                                                   = 0x0B50
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	INT_SAMPLER_2D                                             = 0x8DCA
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	SHADER_CONSISTENT_NV                                       = 0x86DD
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	UNIFORM_OFFSET                                             = 0x8A3B
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	HIGH_INT                                                   = 0x8DF5
	COMPRESSED_RG                                              = 0x8226
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	MATRIX4_NV                                                 = 0x8634
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	YCBAYCR8A_4224_NV                                          = 0x9032
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	FILTER4_SGIS                                               = 0x8146
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	ORDER                                                      = 0x0A01
	RGBA16_EXT                                                 = 0x805B
	MODELVIEW15_ARB                                            = 0x872F
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	SMALL_CW_ARC_TO_NV                                         = 0x14
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	TRIANGLE_STRIP                                             = 0x0005
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	TEXTURE_MIN_LOD                                            = 0x813A
	DRAW_BUFFER10_ATI                                          = 0x882F
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	CURRENT_TANGENT_EXT                                        = 0x843B
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	VERSION_1_1                                                = 1
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	DATA_BUFFER_AMD                                            = 0x9151
	N3F_V3F                                                    = 0x2A25
	ALPHA4_EXT                                                 = 0x803B
	STENCIL_RENDERABLE                                         = 0x8288
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	INDEX_BITS                                                 = 0x0D51
	LIGHT7                                                     = 0x4007
	ALPHA16F_ARB                                               = 0x881C
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	MODELVIEW19_ARB                                            = 0x8733
	COMPRESSED_RG11_EAC                                        = 0x9272
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	SHADER                                                     = 0x82E1
	FLOAT_R32_NV                                               = 0x8885
	FLOAT16_NV                                                 = 0x8FF8
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	VERSION_3_1                                                = 1
	ALPHA4                                                     = 0x803B
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	COLOR_ATTACHMENT13                                         = 0x8CED
	BOUNDING_BOX_NV                                            = 0x908D
	MAP2_INDEX                                                 = 0x0DB1
	COLOR_TABLE_BIAS                                           = 0x80D7
	CULL_VERTEX_EXT                                            = 0x81AA
	RG                                                         = 0x8227
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	DYNAMIC_READ_ARB                                           = 0x88E9
	MAX_INTEGER_SAMPLES                                        = 0x9110
	INTENSITY4_EXT                                             = 0x804A
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	DOT3_RGBA_EXT                                              = 0x8741
	DRAW_BUFFER1_NV                                            = 0x8826
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	PROXY_TEXTURE_3D                                           = 0x8070
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	RESTART_SUN                                                = 0x0001
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	RELATIVE_LINE_TO_NV                                        = 0x05
	RED_EXT                                                    = 0x1903
	IUI_V2F_EXT                                                = 0x81AD
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	CONTEXT_FLAGS                                              = 0x821E
	DOT3_ATI                                                   = 0x8966
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	TEXTURE8                                                   = 0x84C8
	OP_MAX_EXT                                                 = 0x878A
	INTENSITY32F_ARB                                           = 0x8817
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	RGBA4                                                      = 0x8056
	CLIP_PLANE1                                                = 0x3001
	TEXTURE15_ARB                                              = 0x84CF
	PREVIOUS_ARB                                               = 0x8578
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	RGBA16F_EXT                                                = 0x881A
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	SRGB_ALPHA                                                 = 0x8C42
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	ATC_RGB_AMD                                                = 0x8C92
	EXT_texture_object                                         = 1
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	INCR_WRAP_OES                                              = 0x8507
	INTERPOLATE                                                = 0x8575
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	YCRCB_SGIX                                                 = 0x8318
	DS_BIAS_NV                                                 = 0x8716
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	TEXTURE_COMPONENTS                                         = 0x1003
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	MATRIX_MODE                                                = 0x0BA0
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	PATH_FORMAT_SVG_NV                                         = 0x9070
	ACTIVE_VARIABLES                                           = 0x9305
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	FULL_RANGE_EXT                                             = 0x87E1
	BLUE_BIT_ATI                                               = 0x00000004
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	SGIX_impact_pixel_texture                                  = 1
	STENCIL_INDEX                                              = 0x1901
	SAMPLER_BINDING                                            = 0x8919
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	GL_4PASS_1_SGIS                                            = 0x80A5
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	VENDOR                                                     = 0x1F00
	POINT_SIZE_MAX_EXT                                         = 0x8127
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	FRAMEBUFFER_OES                                            = 0x8D40
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	TEXTURE_3D_OES                                             = 0x806F
	COMBINE_RGB                                                = 0x8571
	MODELVIEW10_ARB                                            = 0x872A
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	ARRAY_SIZE                                                 = 0x92FB
	CLEAR                                                      = 0x1500
	SAMPLES_ARB                                                = 0x80A9
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	VIEW_CLASS_48_BITS                                         = 0x82C7
	SRGB8_ALPHA8                                               = 0x8C43
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	RGBA_INTEGER                                               = 0x8D99
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	SYNC_CONDITION_APPLE                                       = 0x9113
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	BUMP_TARGET_ATI                                            = 0x877C
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	SATURATE_BIT_ATI                                           = 0x00000040
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	DRAW_BUFFER3                                               = 0x8828
	MAX_IMAGE_UNITS                                            = 0x8F38
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	TEXTURE_4D_SGIS                                            = 0x8134
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	DRAW_BUFFER3_ATI                                           = 0x8828
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	ACTIVE_RESOURCES                                           = 0x92F5
	DRAW_PIXEL_TOKEN                                           = 0x0705
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	STATIC_COPY_ARB                                            = 0x88E6
	ACCUM                                                      = 0x0100
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	STENCIL_VALUE_MASK                                         = 0x0B93
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	SGIX_convolution_accuracy                                  = 1
	INDEX_LOGIC_OP                                             = 0x0BF1
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	COMBINE_EXT                                                = 0x8570
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	ZOOM_Y                                                     = 0x0D17
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	TEXTURE_BASE_LEVEL                                         = 0x813C
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	TEXTURE_WRAP_R                                             = 0x8072
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	SIGNED_HILO8_NV                                            = 0x885F
	REG_31_ATI                                                 = 0x8940
	INT_VEC4                                                   = 0x8B55
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	INVALID_OPERATION                                          = 0x0502
	DEPTH_WRITEMASK                                            = 0x0B72
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	EXT_histogram                                              = 1
	UNSIGNED_SHORT                                             = 0x1403
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	OPERAND0_ALPHA                                             = 0x8598
	LO_SCALE_NV                                                = 0x870F
	OP_MADD_EXT                                                = 0x8788
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	ALPHA8_EXT                                                 = 0x803C
	MATRIX3_ARB                                                = 0x88C3
	LUMINANCE8I_EXT                                            = 0x8D92
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	TRANSLATE_3D_NV                                            = 0x9091
	TEXTURE_WRAP_T                                             = 0x2803
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	CCW                                                        = 0x0901
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	UNIFORM_TYPE                                               = 0x8A37
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	COMPUTE_SHADER                                             = 0x91B9
	RGBA4_EXT                                                  = 0x8056
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	DECR_WRAP_OES                                              = 0x8508
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	SWIZZLE_STRQ_ATI                                           = 0x897A
	VERTEX_SHADER                                              = 0x8B31
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	COPY_READ_BUFFER                                           = 0x8F36
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	INT64_VEC4_NV                                              = 0x8FEB
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	MAX_VIEWPORTS                                              = 0x825B
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	REG_16_ATI                                                 = 0x8931
	FLOAT_MAT3x2                                               = 0x8B67
	LUMINANCE4                                                 = 0x803F
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	RGBA4_S3TC                                                 = 0x83A3
	TEXTURE26                                                  = 0x84DA
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	FOG_COORD_SRC                                              = 0x8450
	VARIABLE_D_NV                                              = 0x8526
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	COLOR_ATTACHMENT15                                         = 0x8CEF
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	DRAW_BUFFER11_ARB                                          = 0x8830
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	OBJECT_POINT_SGIS                                          = 0x81F5
	MAX_LAYERS                                                 = 0x8281
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	READ_WRITE_ARB                                             = 0x88BA
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	DRAW_BUFFER_EXT                                            = 0x0C01
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	RG8_SNORM                                                  = 0x8F95
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	CONVOLUTION_WIDTH                                          = 0x8018
	YCRCB_444_SGIX                                             = 0x81BC
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	DRAW_BUFFER2_NV                                            = 0x8827
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	OP_MOV_EXT                                                 = 0x8799
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	PERCENTAGE_AMD                                             = 0x8BC3
	VERTEX_SHADER_BIT                                          = 0x00000001
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	RG16UI                                                     = 0x823A
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	TEXTURE_BUFFER                                             = 0x8C2A
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	BLEND_DST_ALPHA                                            = 0x80CA
	SHADER_OBJECT_ARB                                          = 0x8B48
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	LINE_STRIP                                                 = 0x0003
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	LUMINANCE_SNORM                                            = 0x9011
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	LINE_SMOOTH                                                = 0x0B20
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	COLOR_INDEX12_EXT                                          = 0x80E6
	REG_5_ATI                                                  = 0x8926
	SLUMINANCE8_EXT                                            = 0x8C47
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	PATCH_VERTICES                                             = 0x8E72
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	NO_ERROR                                                   = 0
	COLOR_CLEAR_VALUE                                          = 0x0C22
	NICEST                                                     = 0x1102
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	BLEND_EQUATION_OES                                         = 0x8009
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	TEXTURE_RECTANGLE                                          = 0x84F5
	COMBINE                                                    = 0x8570
	FLOAT_VEC2_ARB                                             = 0x8B50
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	AFFINE_3D_NV                                               = 0x9094
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	CW                                                         = 0x0900
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	COLOR_ATTACHMENT3                                          = 0x8CE3
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	BLEND_SRC_ALPHA                                            = 0x80CB
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	R                                                          = 0x2002
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	T2F_C3F_V3F                                                = 0x2A2A
	HISTOGRAM_WIDTH                                            = 0x8026
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	PATH_STENCIL_REF_NV                                        = 0x90B8
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	BUFFER_USAGE_ARB                                           = 0x8765
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	BUFFER_VARIABLE                                            = 0x92E5
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	SAMPLER_BUFFER                                             = 0x8DC2
	SAMPLE_MASK_VALUE                                          = 0x8E52
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	EXP                                                        = 0x0800
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	ACTIVE_TEXTURE                                             = 0x84E0
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	RGBA32F                                                    = 0x8814
	TEXTURE_COMPARE_MODE                                       = 0x884C
	QUERY_NO_WAIT_NV                                           = 0x8E14
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	TANGENT_ARRAY_EXT                                          = 0x8439
	R11F_G11F_B10F                                             = 0x8C3A
	LINE_WIDTH                                                 = 0x0B21
	MAP_COLOR                                                  = 0x0D10
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	MAP1_BINORMAL_EXT                                          = 0x8446
	DSDT8_MAG8_NV                                              = 0x870A
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	CON_15_ATI                                                 = 0x8950
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	OP_INDEX_EXT                                               = 0x8782
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	VIEWPORT_BIT                                               = 0x00000800
	Q                                                          = 0x2003
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	COLOR_TABLE                                                = 0x80D0
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	OPERAND1_RGB                                               = 0x8591
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	LOCATION                                                   = 0x930E
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	RGB10_A2UI                                                 = 0x906F
	MITER_REVERT_NV                                            = 0x90A7
	FRONT_FACE                                                 = 0x0B46
	VARIABLE_B_NV                                              = 0x8524
	INT64_VEC3_NV                                              = 0x8FEA
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	DEPTH_STENCIL_NV                                           = 0x84F9
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	BUFFER_ACCESS                                              = 0x88BB
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	EYE_PLANE                                                  = 0x2502
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	MIN_LOD_WARNING_AMD                                        = 0x919C
	POLYGON_STIPPLE                                            = 0x0B42
	INDEX_ARRAY_TYPE                                           = 0x8085
	COLOR_TABLE_WIDTH                                          = 0x80D9
	REPLACE_MIDDLE_SUN                                         = 0x0002
	DEPTH_STENCIL                                              = 0x84F9
	COMBINE4_NV                                                = 0x8503
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	STREAM_READ                                                = 0x88E1
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	SLIM10U_SGIX                                               = 0x831E
	PROGRAM_LENGTH_NV                                          = 0x8627
	DU8DV8_ATI                                                 = 0x877A
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	LIGHTING_BIT                                               = 0x00000040
	INDEX_MATERIAL_EXT                                         = 0x81B8
	FILTER                                                     = 0x829A
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	OPERAND1_ALPHA_EXT                                         = 0x8599
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	MODELVIEW_MATRIX                                           = 0x0BA6
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	TEXTURE8_ARB                                               = 0x84C8
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	MATRIX15_ARB                                               = 0x88CF
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	EYE_POINT_SGIS                                             = 0x81F4
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	BGRA_INTEGER                                               = 0x8D9B
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	INCR_WRAP_EXT                                              = 0x8507
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	MODELVIEW2_ARB                                             = 0x8722
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	PATH_STROKE_MASK_NV                                        = 0x9084
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	MODULATE_ADD_ATI                                           = 0x8744
	TRACE_NAME_MESA                                            = 0x8756
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	REG_3_ATI                                                  = 0x8924
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	MAX_SHININESS_NV                                           = 0x8504
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	DOUBLE_MAT3x2                                              = 0x8F4B
	SGIS_sharpen_texture                                       = 1
	SHADER_OBJECT_EXT                                          = 0x8B48
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	R8I                                                        = 0x8231
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	CURRENT_ATTRIB_NV                                          = 0x8626
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	GL_4PASS_0_EXT                                             = 0x80A4
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	DRAW_BUFFER14                                              = 0x8833
	COMP_BIT_ATI                                               = 0x00000002
	ALPHA16UI_EXT                                              = 0x8D78
	RGB16_SNORM                                                = 0x8F9A
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	RELATIVE_MOVE_TO_NV                                        = 0x03
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	ALPHA_TEST                                                 = 0x0BC0
	VARIANT_VALUE_EXT                                          = 0x87E4
	DOUBLE_MAT4x3                                              = 0x8F4E
	SGIS_texture_filter4                                       = 1
	TEXTURE21                                                  = 0x84D5
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	EXT_polygon_offset                                         = 1
	PROGRAM_BINDING_ARB                                        = 0x8677
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	STENCIL_BACK_FUNC                                          = 0x8800
	CLAMP_VERTEX_COLOR                                         = 0x891A
	SRGB_ALPHA_EXT                                             = 0x8C42
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	ALPHA_TEST_FUNC                                            = 0x0BC1
	CND0_ATI                                                   = 0x896B
	BOOL_VEC4_ARB                                              = 0x8B59
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	SGIX_pixel_tiles                                           = 1
	DITHER                                                     = 0x0BD0
	NEGATIVE_ONE_EXT                                           = 0x87DF
	CON_18_ATI                                                 = 0x8953
	CLIP_DISTANCE2                                             = 0x3002
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	SAMPLE_BUFFERS                                             = 0x80A8
	MODELVIEW31_ARB                                            = 0x873F
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	MITER_TRUNCATE_NV                                          = 0x90A8
	IS_PER_PATCH                                               = 0x92E7
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	PIXEL_PACK_BUFFER                                          = 0x88EB
	SAMPLES_PASSED_ARB                                         = 0x8914
	PATH_DASH_OFFSET_NV                                        = 0x907E
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	VERTEX_STREAM6_ATI                                         = 0x8772
	STENCIL_BACK_FAIL                                          = 0x8801
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	DOUBLE_MAT2x4                                              = 0x8F4A
	YCBYCR8_422_NV                                             = 0x9031
	NAME_STACK_DEPTH                                           = 0x0D70
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	PROGRAM_SEPARABLE                                          = 0x8258
	LAYER_NV                                                   = 0x8DAA
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	RGBA4_DXT5_S3TC                                            = 0x83A5
	EMBOSS_CONSTANT_NV                                         = 0x855E
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	DEPTH_RANGE                                                = 0x0B70
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	SRC2_RGB                                                   = 0x8582
	RGB8UI_EXT                                                 = 0x8D7D
	RGB8I                                                      = 0x8D8F
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	R8_SNORM                                                   = 0x8F94
	TEXTURE_GATHER                                             = 0x82A2
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	MAX_IMAGE_SAMPLES                                          = 0x906D
	FOG_HINT                                                   = 0x0C54
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	CURRENT_RASTER_COLOR                                       = 0x0B04
	REG_29_ATI                                                 = 0x893E
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	ALPHA32I_EXT                                               = 0x8D84
	NOR                                                        = 0x1508
	INTENSITY16                                                = 0x804D
	QUAD_MESH_SUN                                              = 0x8614
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	CON_23_ATI                                                 = 0x8958
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	DRAW_BUFFER4_NV                                            = 0x8829
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	PATH_MITER_LIMIT_NV                                        = 0x907A
	FRAMEZOOM_SGIX                                             = 0x818B
	R32UI                                                      = 0x8236
	TEXTURE30_ARB                                              = 0x84DE
	MATRIX31_ARB                                               = 0x88DF
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	SHADER_BINARY_DMP                                          = 0x9250
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	INDEX_TEST_REF_EXT                                         = 0x81B7
	SIGNED_RGB8_NV                                             = 0x86FF
	REPEAT                                                     = 0x2901
	BGRA_EXT                                                   = 0x80E1
	FRAGMENT_DEPTH                                             = 0x8452
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	LUMINANCE16UI_EXT                                          = 0x8D7A
	TESS_GEN_MODE                                              = 0x8E76
	PACK_SKIP_ROWS                                             = 0x0D03
	INTERPOLATE_ARB                                            = 0x8575
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	BEVEL_NV                                                   = 0x90A6
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	SHADE_MODEL                                                = 0x0B54
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	INTENSITY32I_EXT                                           = 0x8D85
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	SMOOTH                                                     = 0x1D01
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	MODELVIEW23_ARB                                            = 0x8737
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	REG_6_ATI                                                  = 0x8927
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	IMAGE_1D_ARRAY                                             = 0x9052
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	ADD_SIGNED_EXT                                             = 0x8574
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	ABGR_EXT                                                   = 0x8000
	MAX_DEPTH                                                  = 0x8280
	FLOAT_RGBA16_NV                                            = 0x888A
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	REG_30_ATI                                                 = 0x893F
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	SGIX_scalebias_hint                                        = 1
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	DRAW_BUFFER10_ARB                                          = 0x882F
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	SWIZZLE_STR_ATI                                            = 0x8976
	SKIP_COMPONENTS4_NV                                        = -3
	POINT_SIZE                                                 = 0x0B11
	FOG_COLOR                                                  = 0x0B66
	BLEND_SRC_RGB_OES                                          = 0x80C9
	RG8_EXT                                                    = 0x822B
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	MODELVIEW22_ARB                                            = 0x8736
	HILO_NV                                                    = 0x86F4
	TIMEOUT_EXPIRED                                            = 0x911B
	TEXTURE_MAG_FILTER                                         = 0x2800
	RGB16                                                      = 0x8054
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	AND_REVERSE                                                = 0x1502
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	MAP1_TANGENT_EXT                                           = 0x8444
	SLUMINANCE_EXT                                             = 0x8C46
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	DOUBLE_MAT2x3                                              = 0x8F49
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	RESCALE_NORMAL_EXT                                         = 0x803A
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	CULL_MODES_NV                                              = 0x86E0
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	RENDERBUFFER_WIDTH                                         = 0x8D42
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	LIST_MODE                                                  = 0x0B30
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	Z_EXT                                                      = 0x87D7
	FLOAT_VEC4                                                 = 0x8B52
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	FOG_END                                                    = 0x0B64
	OPERAND3_RGB_NV                                            = 0x8593
	RGB9_E5                                                    = 0x8C3D
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	PACK_SWAP_BYTES                                            = 0x0D00
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	DOT_PRODUCT_NV                                             = 0x86EC
	SAMPLER_2D_ARB                                             = 0x8B5E
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	SIGNED_NEGATE_NV                                           = 0x853D
	OUTPUT_VERTEX_EXT                                          = 0x879A
	BUFFER_ACCESS_ARB                                          = 0x88BB
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	GENERATE_MIPMAP                                            = 0x8191
	SAMPLER                                                    = 0x82E6
	RGB_FLOAT32_APPLE                                          = 0x8815
	MATRIX24_ARB                                               = 0x88D8
	CONDITION_SATISFIED                                        = 0x911C
	RGB4_S3TC                                                  = 0x83A1
	MODELVIEW13_ARB                                            = 0x872D
	FRONT_AND_BACK                                             = 0x0408
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	TEXTURE_LIGHT_EXT                                          = 0x8350
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	DEBUG_OBJECT_MESA                                          = 0x8759
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	SGIX_clipmap                                               = 1
	MAX_HEIGHT                                                 = 0x827F
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	MATRIX6_NV                                                 = 0x8636
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	OP_ROUND_EXT                                               = 0x8790
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	MAP2_VERTEX_3                                              = 0x0DB7
	GREEN                                                      = 0x1904
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	ZOOM_X                                                     = 0x0D16
	C4UB_V2F                                                   = 0x2A22
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	ROUND_NV                                                   = 0x90A4
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	VERTEX_BLEND_ARB                                           = 0x86A7
	DRAW_BUFFER15_NV                                           = 0x8834
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	SGIS_texture_lod                                           = 1
	PATCHES                                                    = 0x000E
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	LUMINANCE16_SNORM                                          = 0x9019
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	INT_IMAGE_1D_EXT                                           = 0x9057
	SGIX_flush_raster                                          = 1
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	OPERAND2_ALPHA                                             = 0x859A
	FLOAT_RGBA32_NV                                            = 0x888B
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	CON_8_ATI                                                  = 0x8949
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	ZERO                                                       = 0
	DOT3_RGB_EXT                                               = 0x8740
	FRAGMENT_SHADER_ATI                                        = 0x8920
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	FRAMEBUFFER_BLEND                                          = 0x828B
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	BLEND_DST_RGB_EXT                                          = 0x80C8
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	INT_IMAGE_BUFFER                                           = 0x905C
	GL_422_AVERAGE_EXT                                         = 0x80CE
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	OBJECT_LINE_SGIS                                           = 0x81F7
	STENCIL_COMPONENTS                                         = 0x8285
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	BLEND_EQUATION_EXT                                         = 0x8009
	VIBRANCE_SCALE_NV                                          = 0x8713
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	LOWER_LEFT                                                 = 0x8CA1
	COEFF                                                      = 0x0A00
	GL_3DC_X_AMD                                               = 0x87F9
	SAMPLES_PASSED                                             = 0x8914
	DOT3_RGB                                                   = 0x86AE
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	COLOR_LOGIC_OP                                             = 0x0BF2
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	BOLD_BIT_NV                                                = 0x01
	LINES_ADJACENCY                                            = 0x000A
	EDGE_FLAG_ARRAY                                            = 0x8079
	NEAREST                                                    = 0x2600
	DISPLAY_LIST                                               = 0x82E7
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	CON_22_ATI                                                 = 0x8957
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	ADD_SIGNED                                                 = 0x8574
	GL_2X_BIT_ATI                                              = 0x00000001
	PATH_FILL_MASK_NV                                          = 0x9081
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	SGIX_calligraphic_fragment                                 = 1
	RGB16F_EXT                                                 = 0x881B
	RGBA8I                                                     = 0x8D8E
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	TEXTURE0                                                   = 0x84C0
	COMBINER2_NV                                               = 0x8552
	PRIMARY_COLOR_EXT                                          = 0x8577
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	DEBUG_PRINT_MESA                                           = 0x875A
	PALETTE8_RGBA4_OES                                         = 0x8B98
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	SPRITE_AXIAL_SGIX                                          = 0x814C
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	UNPACK_RESAMPLE_OML                                        = 0x8985
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	EXT_packed_pixels                                          = 1
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	ZERO_EXT                                                   = 0x87DD
	LUMINANCE8_ALPHA8                                          = 0x8045
	DT_BIAS_NV                                                 = 0x8717
	RENDERBUFFER_EXT                                           = 0x8D41
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	TEXTURE1_ARB                                               = 0x84C1
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	WRITE_ONLY_ARB                                             = 0x88B9
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	INT                                                        = 0x1404
	PROJECTION                                                 = 0x1701
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	STENCIL_BACK_REF                                           = 0x8CA3
	STENCIL_INDEX4_OES                                         = 0x8D47
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	MINMAX_FORMAT                                              = 0x802F
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	TEXTURE9_ARB                                               = 0x84C9
	DOT3_RGBA_ARB                                              = 0x86AF
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	TEXTURE9                                                   = 0x84C9
	PREVIOUS_EXT                                               = 0x8578
	RGB_FLOAT16_ATI                                            = 0x881B
	COLOR_ATTACHMENT4                                          = 0x8CE4
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	INT_IMAGE_2D_RECT                                          = 0x905A
	PROGRAM_OUTPUT                                             = 0x92E4
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	SRGB_EXT                                                   = 0x8C40
	OP_RECIP_SQRT_EXT                                          = 0x8795
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	VERSION_3_0                                                = 1
	LIST_BIT                                                   = 0x00020000
	DEBUG_TYPE_ERROR                                           = 0x824C
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	SAMPLES                                                    = 0x80A9
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	RG16I                                                      = 0x8239
	SOURCE1_RGB                                                = 0x8581
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	LOW_FLOAT                                                  = 0x8DF0
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	PRESERVE_ATI                                               = 0x8762
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	ITALIC_BIT_NV                                              = 0x02
	GL_2D                                                      = 0x0600
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	SIGNALED_APPLE                                             = 0x9119
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	SRC_ALPHA                                                  = 0x0302
	FLOAT                                                      = 0x1406
	BLEND_EQUATION                                             = 0x8009
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	INFO_LOG_LENGTH                                            = 0x8B84
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	INTERLACE_SGIX                                             = 0x8094
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	SPARE1_NV                                                  = 0x852F
	SIGNED_IDENTITY_NV                                         = 0x853C
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	REG_22_ATI                                                 = 0x8937
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	SAMPLER_3D_ARB                                             = 0x8B5F
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	MAP1_INDEX                                                 = 0x0D91
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	TEXTURE14                                                  = 0x84CE
	MATRIX9_ARB                                                = 0x88C9
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	RGB565                                                     = 0x8D62
	RGB_INTEGER_EXT                                            = 0x8D98
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	TEXTURE28                                                  = 0x84DC
	FLOAT_MAT3x4                                               = 0x8B68
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	HINT_BIT                                                   = 0x00008000
	REPLICATE_BORDER_HP                                        = 0x8153
	REG_11_ATI                                                 = 0x892C
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	READ_BUFFER_EXT                                            = 0x0C02
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	DRAW_BUFFER14_NV                                           = 0x8833
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	FRACTIONAL_ODD                                             = 0x8E7B
	LINE_LOOP                                                  = 0x0002
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	CLIP_DISTANCE6                                             = 0x3006
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	MATRIX12_ARB                                               = 0x88CC
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	SIGNED_HILO_NV                                             = 0x86F9
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	SGIX_texture_lod_bias                                      = 1
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	DEPTH_COMPONENT32F                                         = 0x8CAC
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	ALPHA16I_EXT                                               = 0x8D8A
	CURRENT_RASTER_POSITION                                    = 0x0B07
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	SAMPLER_CUBE                                               = 0x8B60
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	VERTEX4_BIT_PGI                                            = 0x00000008
	INTENSITY_EXT                                              = 0x8049
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	TEXTURE17_ARB                                              = 0x84D1
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	RGB8_EXT                                                   = 0x8051
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	OCCLUSION_TEST_HP                                          = 0x8165
	SLUMINANCE8_NV                                             = 0x8C47
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	TEXTURE_GEN_S                                              = 0x0C60
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	GL_1PASS_SGIS                                              = 0x80A1
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	BLEND_DST_RGB_OES                                          = 0x80C8
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	LEFT                                                       = 0x0406
	DIFFUSE                                                    = 0x1201
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	DEPTH_FUNC                                                 = 0x0B74
	C4F_N3F_V3F                                                = 0x2A26
	R16                                                        = 0x822A
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	RGBA_FLOAT32_ATI                                           = 0x8814
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	REG_28_ATI                                                 = 0x893D
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	DUAL_ALPHA12_SGIS                                          = 0x8112
	FENCE_STATUS_NV                                            = 0x84F3
	COLOR_ATTACHMENT9                                          = 0x8CE9
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	PACK_ALIGNMENT                                             = 0x0D05
	TEXTURE_GREEN_SIZE                                         = 0x805D
	PROXY_TEXTURE_2D                                           = 0x8064
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	FOG_INDEX                                                  = 0x0B61
	STENCIL_EXT                                                = 0x1802
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	DRAW_BUFFER9                                               = 0x882E
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	MINMAX_SINK_EXT                                            = 0x8030
	LUMINANCE8_EXT                                             = 0x8040
	COMPUTE_TEXTURE                                            = 0x82A0
	TEXTURE12                                                  = 0x84CC
	DRAW_BUFFER14_ARB                                          = 0x8833
	COLOR_ATTACHMENT11                                         = 0x8CEB
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	ALPHA_SCALE                                                = 0x0D1C
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	MATRIX11_ARB                                               = 0x88CB
	MATRIX14_ARB                                               = 0x88CE
	DOUBLE_MAT3                                                = 0x8F47
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	IGNORE_BORDER_HP                                           = 0x8150
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	RGBA8I_EXT                                                 = 0x8D8E
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	CLIP_DISTANCE4                                             = 0x3004
	PHONG_HINT_WIN                                             = 0x80EB
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	RGBA16UI                                                   = 0x8D76
	COLOR_MATERIAL                                             = 0x0B57
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	AFFINE_2D_NV                                               = 0x9092
	R3_G3_B2                                                   = 0x2A10
	REDUCE_EXT                                                 = 0x8016
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	TRIANGLES                                                  = 0x0004
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	FLOAT_VEC3                                                 = 0x8B51
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	MAP2_VERTEX_4                                              = 0x0DB8
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	TEXTURE10                                                  = 0x84CA
	OPERAND2_ALPHA_EXT                                         = 0x859A
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	MAX_PATCH_VERTICES                                         = 0x8E7D
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	LIGHT0                                                     = 0x4000
	PACK_CMYK_HINT_EXT                                         = 0x800E
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	INT16_VEC4_NV                                              = 0x8FE7
	IMAGE_1D                                                   = 0x904C
	INTENSITY16F_ARB                                           = 0x881D
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	HALF_FLOAT_NV                                              = 0x140B
	COPY_INVERTED                                              = 0x150C
	SHADER_IMAGE_LOAD                                          = 0x82A4
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	STENCIL_FUNC                                               = 0x0B92
	W_EXT                                                      = 0x87D8
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	MODELVIEW25_ARB                                            = 0x8739
	DRAW_BUFFER8_ATI                                           = 0x882D
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	COLOR_ARRAY_EXT                                            = 0x8076
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	EXT_subtexture                                             = 1
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	VERTEX_ARRAY_SIZE                                          = 0x807A
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	DRAW_BUFFER11_ATI                                          = 0x8830
	ARRAY_BUFFER_ARB                                           = 0x8892
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	BLEND_SRC                                                  = 0x0BE1
	OP_RECIP_EXT                                               = 0x8794
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	REG_26_ATI                                                 = 0x893B
	DEPTH_ATTACHMENT                                           = 0x8D00
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	DOUBLE_MAT3_EXT                                            = 0x8F47
	GL_4PASS_3_EXT                                             = 0x80A7
	TEXTURE22_ARB                                              = 0x84D6
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	SRC0_RGB                                                   = 0x8580
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	CON_5_ATI                                                  = 0x8946
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	COLOR_ARRAY_TYPE                                           = 0x8082
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	DRAW_BUFFER15_ATI                                          = 0x8834
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	COMMAND_BARRIER_BIT                                        = 0x00000040
	FULL_SUPPORT                                               = 0x82B7
	FLOAT_R_NV                                                 = 0x8880
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	POINT_SPRITE_NV                                            = 0x8861
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	MAX_TEXTURE_COORDS                                         = 0x8871
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	TEXTURE7                                                   = 0x84C7
	DEBUG_ASSERT_MESA                                          = 0x875B
	RGB_FLOAT32_ATI                                            = 0x8815
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	INT_10_10_10_2_OES                                         = 0x8DF7
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	SAMPLER_1D                                                 = 0x8B5D
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	SRC1_COLOR                                                 = 0x88F9
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	BLUE_INTEGER                                               = 0x8D96
	FACTOR_MIN_AMD                                             = 0x901C
	BUFFER_OBJECT_EXT                                          = 0x9151
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	INT16_NV                                                   = 0x8FE4
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	SGIX_reference_plane                                       = 1
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	OUTPUT_COLOR1_EXT                                          = 0x879C
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	COLOR_MATERIAL_FACE                                        = 0x0B55
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	CON_20_ATI                                                 = 0x8955
	INT8_VEC4_NV                                               = 0x8FE3
	LUMINANCE4_ALPHA4                                          = 0x8043
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	SUBTRACT_ARB                                               = 0x84E7
	FLOAT_RGB16_NV                                             = 0x8888
	POINT_BIT                                                  = 0x00000002
	COLOR_INDEX                                                = 0x1900
	LUMINANCE16_EXT                                            = 0x8042
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	READ_PIXELS_TYPE                                           = 0x828E
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	DRAW_BUFFER14_ATI                                          = 0x8833
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	INVALID_ENUM                                               = 0x0500
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	INT_IMAGE_2D_EXT                                           = 0x9058
	TEXTURE_ENV                                                = 0x2300
	MATRIX7_NV                                                 = 0x8637
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	DEPTH                                                      = 0x1801
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	LUMINANCE8UI_EXT                                           = 0x8D80
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	VARIANT_ARRAY_EXT                                          = 0x87E8
	CLOSE_PATH_NV                                              = 0x00
	MULTISAMPLE                                                = 0x809D
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	ETC1_SRGB8_NV                                              = 0x88EE
	MAP_WRITE_BIT_EXT                                          = 0x0002
	SPOT_DIRECTION                                             = 0x1204
	MATRIX30_ARB                                               = 0x88DE
	FLOAT_VEC2                                                 = 0x8B50
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	LUMINANCE12_EXT                                            = 0x8041
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	TEXTURE6_ARB                                               = 0x84C6
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	LINE_WIDTH_RANGE                                           = 0x0B22
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	EVAL_BIT                                                   = 0x00010000
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	CLIP_DISTANCE5                                             = 0x3005
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	VIEW_CLASS_32_BITS                                         = 0x82C8
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	VERTEX_ARRAY_TYPE                                          = 0x807B
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	RGBA32UI_EXT                                               = 0x8D70
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	CLEAR_BUFFER                                               = 0x82B4
	VIEW_CLASS_96_BITS                                         = 0x82C5
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	BUFFER_MAPPED                                              = 0x88BC
	REG_17_ATI                                                 = 0x8932
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	VALIDATE_STATUS                                            = 0x8B83
	ONE                                                        = 1
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	SIGNED_RGB_NV                                              = 0x86FE
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	SPRITE_AXIS_SGIX                                           = 0x814A
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	RGB16UI_EXT                                                = 0x8D77
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	DUAL_ALPHA4_SGIS                                           = 0x8110
	COMBINE_ALPHA_ARB                                          = 0x8572
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	COORD_REPLACE_NV                                           = 0x8862
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	DOUBLE_MAT4                                                = 0x8F48
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	OP_MUL_EXT                                                 = 0x8786
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	DST_ALPHA                                                  = 0x0304
	AUX1                                                       = 0x040A
	COMPRESSED_RGB_ARB                                         = 0x84ED
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	REG_24_ATI                                                 = 0x8939
	FLOAT_MAT2_ARB                                             = 0x8B5A
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	STENCIL_BUFFER_BIT                                         = 0x00000400
	NOTEQUAL                                                   = 0x0205
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	SGIS_texture4D                                             = 1
	CONSTANT_EXT                                               = 0x8576
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	BITMAP_TOKEN                                               = 0x0704
	MIPMAP                                                     = 0x8293
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	PRIMITIVES_GENERATED                                       = 0x8C87
	ALPHA8_SNORM                                               = 0x9014
	BLEND                                                      = 0x0BE2
	DEPTH_SCALE                                                = 0x0D1E
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	MATRIX21_ARB                                               = 0x88D5
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	SOURCE0_ALPHA_ARB                                          = 0x8588
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	ISOLINES                                                   = 0x8E7A
	RGB5_A1_OES                                                = 0x8057
	CLAMP_TO_BORDER_NV                                         = 0x812D
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	FLOAT_MAT4_ARB                                             = 0x8B5C
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	POLYGON_OFFSET_LINE                                        = 0x2A02
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	MAX_SUBROUTINES                                            = 0x8DE7
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	NORMAL_MAP                                                 = 0x8511
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	FIXED_ONLY_ARB                                             = 0x891D
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	SGIX_list_priority                                         = 1
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	CONVOLUTION_FORMAT                                         = 0x8017
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	DEPTH_COMPONENT16                                          = 0x81A5
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	FLOAT_VEC3_ARB                                             = 0x8B51
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	SKIP_COMPONENTS1_NV                                        = -6
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	LINEAR                                                     = 0x2601
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	RGBA16I_EXT                                                = 0x8D88
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	STENCIL_FAIL                                               = 0x0B94
	COMBINER_MUX_SUM_NV                                        = 0x8547
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	VERSION_3_2                                                = 1
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	R1UI_V3F_SUN                                               = 0x85C4
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	POLYGON_TOKEN                                              = 0x0703
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	SPARE0_NV                                                  = 0x852E
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	MODELVIEW27_ARB                                            = 0x873B
	OP_CLAMP_EXT                                               = 0x878E
	REG_2_ATI                                                  = 0x8923
	SPOT_CUTOFF                                                = 0x1206
	LOCAL_EXT                                                  = 0x87C4
	MATRIX4_ARB                                                = 0x88C4
	DOT4_ATI                                                   = 0x8967
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	SOURCE1_RGB_ARB                                            = 0x8581
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	MULTIVIEW_EXT                                              = 0x90F1
	CLIP_PLANE2                                                = 0x3002
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	BOOL_VEC3                                                  = 0x8B58
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	SGIS_texture_border_clamp                                  = 1
	CUBIC_EXT                                                  = 0x8334
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	FRONT                                                      = 0x0404
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	CULL_VERTEX_IBM                                            = 103050
	GL_3_BYTES                                                 = 0x1408
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	RGBA8UI                                                    = 0x8D7C
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	RGBA16                                                     = 0x805B
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	POINTS                                                     = 0x0000
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	RGB8UI                                                     = 0x8D7D
	RGBA12_EXT                                                 = 0x805A
	WEIGHT_ARRAY_OES                                           = 0x86AD
	REG_9_ATI                                                  = 0x892A
	CONTEXT_PROFILE_MASK                                       = 0x9126
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	MAX_VARYING_VECTORS                                        = 0x8DFC
	COLOR_COMPONENTS                                           = 0x8283
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	BIAS_BIT_ATI                                               = 0x00000008
	SHADER_TYPE                                                = 0x8B4F
	SGIX_ycrcb                                                 = 1
	UNDEFINED_VERTEX                                           = 0x8260
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	SRC0_ALPHA                                                 = 0x8588
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	PATH_JOIN_STYLE_NV                                         = 0x9079
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	STENCIL_BITS                                               = 0x0D57
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	TEXTURE17                                                  = 0x84D1
	Z400_BINARY_AMD                                            = 0x8740
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	FIXED_OES                                                  = 0x140C
	VIEW_CLASS_64_BITS                                         = 0x82C6
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	MAX_SAMPLES                                                = 0x8D57
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	RGBA12                                                     = 0x805A
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	GREEN_BIT_ATI                                              = 0x00000002
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	SAMPLER_1D_SHADOW                                          = 0x8B61
	RGBA_SNORM                                                 = 0x8F93
	LEQUAL                                                     = 0x0203
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	VARIANT_EXT                                                = 0x87C1
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	T2F_V3F                                                    = 0x2A27
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	SAMPLE_MASK                                                = 0x8E51
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	TEXTURE_BINDING_1D                                         = 0x8068
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	DEPTH_STENCIL_EXT                                          = 0x84F9
	PERTURB_EXT                                                = 0x85AE
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	FOG_FUNC_SGIS                                              = 0x812A
	TEXTURE15                                                  = 0x84CF
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	UNIFORM_SIZE                                               = 0x8A38
	MODULATE_COLOR_IMG                                         = 0x8C04
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	HALF_FLOAT_OES                                             = 0x8D61
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	QUERY_OBJECT_AMD                                           = 0x9153
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	DRAW_BUFFER7_ATI                                           = 0x882C
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	SAMPLE_POSITION                                            = 0x8E50
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	DRAW_BUFFER8_NV                                            = 0x882D
	QUERY_COUNTER_BITS                                         = 0x8864
	INT_VEC4_ARB                                               = 0x8B55
	RG_SNORM                                                   = 0x8F91
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	SELECT                                                     = 0x1C02
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	DRAW_BUFFER                                                = 0x0C01
	EYE_LINEAR                                                 = 0x2400
	LIST_PRIORITY_SGIX                                         = 0x8182
	OP_SUB_EXT                                                 = 0x8796
	ALPHA16_EXT                                                = 0x803E
	FRAGMENT_COLOR_EXT                                         = 0x834C
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	INCR_WRAP                                                  = 0x8507
	COMBINER1_NV                                               = 0x8551
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	RGBA32UI                                                   = 0x8D70
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	COMPILE                                                    = 0x1300
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	MULT                                                       = 0x0103
	PN_TRIANGLES_ATI                                           = 0x87F0
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	LINK_STATUS                                                = 0x8B82
	ARC_TO_NV                                                  = 0xFE
	FALSE                                                      = 0
	RG16                                                       = 0x822C
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	RG32I                                                      = 0x823B
	RG_EXT                                                     = 0x8227
	SRGB_READ                                                  = 0x8297
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	CONSTANT_ALPHA_EXT                                         = 0x8003
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	RENDERBUFFER_BINDING                                       = 0x8CA7
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	FIELDS_NV                                                  = 0x8E27
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	IMAGE_2D_EXT                                               = 0x904D
	PATH_GEN_MODE_NV                                           = 0x90B0
	TEXTURE_LOD_BIAS                                           = 0x8501
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	FLOAT_VEC4_ARB                                             = 0x8B52
	MAX_CLIP_PLANES                                            = 0x0D32
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	ALPHA_INTEGER                                              = 0x8D97
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	OR                                                         = 0x1507
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	HIGH_FLOAT                                                 = 0x8DF2
	BLUE_BIAS                                                  = 0x0D1B
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	CND_ATI                                                    = 0x896A
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	DRAW_BUFFER1_ATI                                           = 0x8826
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	DRAW_BUFFER13_ATI                                          = 0x8832
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	BUFFER_MAP_POINTER                                         = 0x88BD
	REG_0_ATI                                                  = 0x8921
	SGIX_shadow_ambient                                        = 1
	ALPHA_TEST_REF                                             = 0x0BC2
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	COMPILE_STATUS                                             = 0x8B81
	INTENSITY16UI_EXT                                          = 0x8D79
	INDEX_CLEAR_VALUE                                          = 0x0C20
	CLAMP_TO_EDGE                                              = 0x812F
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	BLEND_EQUATION_ALPHA                                       = 0x883D
	POINT_SPRITE                                               = 0x8861
	GL_4X_BIT_ATI                                              = 0x00000002
	POLYGON                                                    = 0x0009
	MATRIX0_NV                                                 = 0x8630
	LUMINANCE16F_ARB                                           = 0x881E
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	EXT_texture3D                                              = 1
	MAX_CLIP_DISTANCES                                         = 0x0D32
	DECR                                                       = 0x1E03
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	MODULATE                                                   = 0x2100
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	TESS_GEN_POINT_MODE                                        = 0x8E79
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	DUAL_ALPHA16_SGIS                                          = 0x8113
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	INVALID_VALUE                                              = 0x0501
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	ALPHA32UI_EXT                                              = 0x8D72
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	LUMINANCE_ALPHA                                            = 0x190A
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	DRAW_BUFFER12_NV                                           = 0x8831
	FENCE_APPLE                                                = 0x8A0B
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	POSITION                                                   = 0x1203
	RGB10_EXT                                                  = 0x8052
	HALF_BIT_ATI                                               = 0x00000008
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	R16UI                                                      = 0x8234
	PASS_THROUGH_NV                                            = 0x86E6
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	RGB16I                                                     = 0x8D89
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	FUNC_ADD_EXT                                               = 0x8006
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	ONE_EXT                                                    = 0x87DE
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	MATRIX5_NV                                                 = 0x8635
	DRAW_BUFFER3_ARB                                           = 0x8828
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	FRAMEBUFFER                                                = 0x8D40
	SGI_texture_color_table                                    = 1
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	MAX_SAMPLES_IMG                                            = 0x9135
	SGIX_fragment_lighting                                     = 1
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	SGIS_detail_texture                                        = 1
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	COMPRESSED_RGB                                             = 0x84ED
	DYNAMIC_DRAW                                               = 0x88E8
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	STREAM_COPY                                                = 0x88E2
	STENCIL_TEST                                               = 0x0B90
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	SAMPLER_3D                                                 = 0x8B5F
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	HALF_FLOAT_ARB                                             = 0x140B
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	FRAME_NV                                                   = 0x8E26
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	DRAW_BUFFER11                                              = 0x8830
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	STENCIL_INDEX8                                             = 0x8D48
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	DYNAMIC_READ                                               = 0x88E9
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	CONVOLUTION_1D                                             = 0x8010
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	ALREADY_SIGNALED                                           = 0x911A
	SGIS_fog_function                                          = 1
	DOUBLEBUFFER                                               = 0x0C32
	NORMAL_ARRAY                                               = 0x8075
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	CON_2_ATI                                                  = 0x8943
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	REFLECTION_MAP                                             = 0x8512
	RGB565_OES                                                 = 0x8D62
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	MAX_EVAL_ORDER                                             = 0x0D30
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	DEPTH_BOUNDS_EXT                                           = 0x8891
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	TEXTURE_COMPRESSED                                         = 0x86A1
	EXPAND_NEGATE_NV                                           = 0x8539
	EMBOSS_MAP_NV                                              = 0x855F
	X_EXT                                                      = 0x87D5
	DRAW_BUFFER6_ARB                                           = 0x882B
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	INT64_VEC2_NV                                              = 0x8FE9
	LUMINANCE12_ALPHA12                                        = 0x8047
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	CURRENT_FOG_COORDINATE                                     = 0x8453
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	TEXTURE_COORD_NV                                           = 0x8C79
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	SHORT                                                      = 0x1402
	RGB5_EXT                                                   = 0x8050
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	DEPTH_COMPONENTS                                           = 0x8284
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	RENDER_MODE                                                = 0x0C40
	RGB10_A2_EXT                                               = 0x8059
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	ATTENUATION_EXT                                            = 0x834D
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	SGIX_async                                                 = 1
	LINE_RESET_TOKEN                                           = 0x0707
	QUAD_ALPHA4_SGIS                                           = 0x811E
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	SIGNED_ALPHA8_NV                                           = 0x8706
	OP_DOT3_EXT                                                = 0x8784
	DRAW_BUFFER5_ATI                                           = 0x882A
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	FUNC_SUBTRACT                                              = 0x800A
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	STATIC_DRAW_ARB                                            = 0x88E4
	RG16_SNORM                                                 = 0x8F99
	MATRIX2_ARB                                                = 0x88C2
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	OPERAND0_RGB_ARB                                           = 0x8590
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	COLOR_BUFFER_BIT                                           = 0x00004000
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	RGB16F_ARB                                                 = 0x881B
	MAX_VARYING_FLOATS                                         = 0x8B4B
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	INDEX_ARRAY_EXT                                            = 0x8077
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	POINT_SPRITE_ARB                                           = 0x8861
	CURRENT_TIME_NV                                            = 0x8E28
	BACK                                                       = 0x0405
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	COMBINE_ALPHA_EXT                                          = 0x8572
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	TEXTURE13_ARB                                              = 0x84CD
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	STEREO                                                     = 0x0C33
	COUNTER_RANGE_AMD                                          = 0x8BC1
	OPERAND1_RGB_ARB                                           = 0x8591
	RASTERIZER_DISCARD                                         = 0x8C89
	WRAP_BORDER_SUN                                            = 0x81D4
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	HILO8_NV                                                   = 0x885E
	MATRIX1_ARB                                                = 0x88C1
	REG_7_ATI                                                  = 0x8928
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	CURRENT_RASTER_INDEX                                       = 0x0B05
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	RGB4                                                       = 0x804F
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	PRESENT_DURATION_NV                                        = 0x8E2B
	IMAGE_2D_RECT                                              = 0x904F
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	OP_SET_GE_EXT                                              = 0x878C
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	UNPACK_LSB_FIRST                                           = 0x0CF1
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	SOURCE2_RGB_ARB                                            = 0x8582
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	GL_4PASS_2_EXT                                             = 0x80A6
	READ_PIXELS                                                = 0x828C
	DOT3_RGBA_IMG                                              = 0x86AF
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	RGBA16UI_EXT                                               = 0x8D76
	FIELD_UPPER_NV                                             = 0x9022
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	UNSIGNED_INT_24_8                                          = 0x84FA
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	REDUCE                                                     = 0x8016
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	OBJECT_LINEAR                                              = 0x2401
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	MODELVIEW0_ARB                                             = 0x1700
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	BLUE_BITS                                                  = 0x0D54
	TEXTURE_RESIDENT                                           = 0x8067
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	SIGNED_INTENSITY8_NV                                       = 0x8708
	LUMINANCE32F_ARB                                           = 0x8818
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	TEXTURE_HEIGHT                                             = 0x1001
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	FUNC_ADD_OES                                               = 0x8006
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	VERTEX_STREAM2_ATI                                         = 0x876E
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	READ_ONLY_ARB                                              = 0x88B8
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	IMAGE_BUFFER                                               = 0x9051
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	COLOR_INDEX4_EXT                                           = 0x80E4
	RG32UI                                                     = 0x823C
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	SPRITE_MODE_SGIX                                           = 0x8149
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	PACK_RESAMPLE_SGIX                                         = 0x842C
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	VERSION_1_2                                                = 1
	MATRIX26_ARB                                               = 0x88DA
	CON_10_ATI                                                 = 0x894B
	LOW_INT                                                    = 0x8DF3
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	PRIMITIVE_RESTART_NV                                       = 0x8558
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	INVARIANT_VALUE_EXT                                        = 0x87EA
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	COLOR_MATRIX                                               = 0x80B1
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	INTERPOLATE_EXT                                            = 0x8575
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	FLOAT_RGB_NV                                               = 0x8882
	SAMPLER_3D_OES                                             = 0x8B5F
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	FLOAT16_VEC4_NV                                            = 0x8FFB
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	EXT_vertex_array                                           = 1
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	DSDT_MAG_VIB_NV                                            = 0x86F7
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	DRAW_BUFFER15_ARB                                          = 0x8834
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	NEGATE_BIT_ATI                                             = 0x00000004
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	NAME_LENGTH                                                = 0x92F9
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	SKIP_COMPONENTS3_NV                                        = -4
	FRACTIONAL_EVEN                                            = 0x8E7C
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	ARRAY_BUFFER_BINDING                                       = 0x8894
	VERSION_1_4                                                = 1
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	GEQUAL                                                     = 0x0206
	MAP1_COLOR_4                                               = 0x0D90
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	RGBA32F_ARB                                                = 0x8814
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	DOMAIN                                                     = 0x0A02
	LINE_STIPPLE                                               = 0x0B24
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	MODELVIEW4_ARB                                             = 0x8724
	GENERATE_MIPMAP_HINT                                       = 0x8192
	TEXTURE30                                                  = 0x84DE
	TEXTURE5_ARB                                               = 0x84C5
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	READ_PIXELS_FORMAT                                         = 0x828D
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	RELEASED_APPLE                                             = 0x8A19
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	PROGRAM_POINT_SIZE                                         = 0x8642
	ADD_BLEND_IMG                                              = 0x8C09
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	PREVIOUS                                                   = 0x8578
	ADD                                                        = 0x0104
	GREATER                                                    = 0x0204
	ONE_MINUS_DST_COLOR                                        = 0x0307
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	MOVE_TO_NV                                                 = 0x02
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	VERSION                                                    = 0x1F02
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	CON_19_ATI                                                 = 0x8954
	BOOL_VEC2_ARB                                              = 0x8B57
	STENCIL_ATTACHMENT                                         = 0x8D20
	SYNC_STATUS                                                = 0x9114
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	MIRRORED_REPEAT                                            = 0x8370
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	GL_3D                                                      = 0x0601
	CONVOLUTION_HEIGHT                                         = 0x8019
	REGISTER_COMBINERS_NV                                      = 0x8522
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	MODELVIEW8_ARB                                             = 0x8728
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	SRC2_ALPHA                                                 = 0x858A
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	EQUAL                                                      = 0x0202
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	TEXTURE16_ARB                                              = 0x84D0
	DRAW_BUFFER13_ARB                                          = 0x8832
	RGB32UI                                                    = 0x8D71
	AUX0                                                       = 0x0409
	SOURCE2_ALPHA_ARB                                          = 0x858A
	ALPHA_FLOAT16_ATI                                          = 0x881C
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	TRACE_MASK_MESA                                            = 0x8755
	DRAW_BUFFER9_NV                                            = 0x882E
	MATRIX_PALETTE_OES                                         = 0x8840
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	SAMPLES_SGIS                                               = 0x80A9
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	CONSTANT                                                   = 0x8576
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	GL_4PASS_3_SGIS                                            = 0x80A7
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	CONSTANT_COLOR1_NV                                         = 0x852B
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	DRAW_BUFFER9_ATI                                           = 0x882E
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	ACCUM_BUFFER_BIT                                           = 0x00000200
	ALPHA_BITS                                                 = 0x0D55
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	RGB_SCALE_ARB                                              = 0x8573
	MODELVIEW16_ARB                                            = 0x8730
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	COLOR_INDEX16_EXT                                          = 0x80E7
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	ARB_imaging                                                = 1
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	RED_MIN_CLAMP_INGR                                         = 0x8560
	VERTEX_PROGRAM_NV                                          = 0x8620
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	SGIX_ir_instrument1                                        = 1
	MAP_WRITE_BIT                                              = 0x0002
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	RENDER                                                     = 0x1C00
	TEXTURE_SHADOW                                             = 0x82A1
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	TEXTURE_1D                                                 = 0x0DE0
	ALPHA16                                                    = 0x803E
	COLOR_ENCODING                                             = 0x8296
	COMBINER4_NV                                               = 0x8554
	COMBINE_ARB                                                = 0x8570
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	AUTO_NORMAL                                                = 0x0D80
	LUMINANCE12_ALPHA4                                         = 0x8046
	C3F_V3F                                                    = 0x2A24
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	INDEX_ARRAY_LIST_IBM                                       = 103073
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	PROGRAM_LENGTH_ARB                                         = 0x8627
	INT_VEC2                                                   = 0x8B53
	STATIC_READ_ARB                                            = 0x88E5
	SRGB8                                                      = 0x8C41
	BLUE_INTEGER_EXT                                           = 0x8D96
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	FOG_BIT                                                    = 0x00000080
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	FOG_COORDINATE                                             = 0x8451
	SOURCE2_RGB                                                = 0x8582
	MATRIX7_ARB                                                = 0x88C7
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	STENCIL_INDEX1_EXT                                         = 0x8D46
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	LINES_ADJACENCY_EXT                                        = 0x000A
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	BUFFER_ACCESS_OES                                          = 0x88BB
	DEPTH_BIAS                                                 = 0x0D1F
	ALPHA_TEST_QCOM                                            = 0x0BC0
	SEPARABLE_2D                                               = 0x8012
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	FLOAT16_VEC2_NV                                            = 0x8FF9
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	COORD_REPLACE                                              = 0x8862
	TEXTURE_1D_ARRAY                                           = 0x8C18
	AMBIENT                                                    = 0x1200
	LUMINANCE6_ALPHA2                                          = 0x8044
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	ETC1_RGB8_OES                                              = 0x8D64
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	SGIX_instruments                                           = 1
	COLOR_WRITEMASK                                            = 0x0C23
	SOURCE1_RGB_EXT                                            = 0x8581
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	BGRA_INTEGER_EXT                                           = 0x8D9B
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	SRGB_DECODE_ARB                                            = 0x8299
	EIGHTH_BIT_ATI                                             = 0x00000020
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	SAMPLE_POSITION_NV                                         = 0x8E50
	IMAGE_3D                                                   = 0x904E
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	RGBA4_OES                                                  = 0x8056
	INT16_VEC3_NV                                              = 0x8FE6
	OBJECT_TYPE_APPLE                                          = 0x9112
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	GL_8X_BIT_ATI                                              = 0x00000004
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	SGIS_texture_edge_clamp                                    = 1
	HISTOGRAM_FORMAT                                           = 0x8027
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	COLOR_SAMPLES_NV                                           = 0x8E20
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	TEXTURE_DEPTH                                              = 0x8071
	UNSIGNED_INVERT_NV                                         = 0x8537
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	UNDEFINED_APPLE                                            = 0x8A1C
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	IMAGE_BUFFER_EXT                                           = 0x9051
	POINT_SMOOTH                                               = 0x0B10
	MAX_TEXTURE_SIZE                                           = 0x0D33
	ALPHA12_EXT                                                = 0x803D
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	DEBUG_TYPE_OTHER                                           = 0x8251
	SRC1_RGB                                                   = 0x8581
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	SGX_BINARY_IMG                                             = 0x8C0A
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	NUM_EXTENSIONS                                             = 0x821D
	TEXTURE25                                                  = 0x84D9
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	OUTPUT_COLOR0_EXT                                          = 0x879B
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	CON_26_ATI                                                 = 0x895B
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	TEXTURE_3D                                                 = 0x806F
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	REPLACE_VALUE_AMD                                          = 0x874B
	FLOAT_MAT2                                                 = 0x8B5A
	INTENSITY8UI_EXT                                           = 0x8D7F
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	PATH_DASH_CAPS_NV                                          = 0x907B
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	RGB10_A2                                                   = 0x8059
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	POINT_SIZE_MIN_EXT                                         = 0x8126
	TEXTURE27_ARB                                              = 0x84DB
	PALETTE4_RGBA8_OES                                         = 0x8B91
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	ALPHA12                                                    = 0x803D
	CLIP_PLANE0                                                = 0x3000
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	PATH_GEN_COEFF_NV                                          = 0x90B1
	TEXTURE_BIT                                                = 0x00040000
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	INVARIANT_EXT                                              = 0x87C2
	READ_WRITE                                                 = 0x88BA
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	NORMAL_MAP_EXT                                             = 0x8511
	MAGNITUDE_SCALE_NV                                         = 0x8712
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	XOR                                                        = 0x1506
	MIN                                                        = 0x8007
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	TEXTURE22                                                  = 0x84D6
	LOCATION_INDEX                                             = 0x930F
	TEXTURE_GEN_T                                              = 0x0C61
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	COMPRESSED_RED                                             = 0x8225
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	CONVEX_HULL_NV                                             = 0x908B
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	CON_6_ATI                                                  = 0x8947
	INTENSITY8_SNORM                                           = 0x9017
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	DST_COLOR                                                  = 0x0306
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	MODELVIEW28_ARB                                            = 0x873C
	CURRENT_QUERY                                              = 0x8865
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	COLOR_ATTACHMENT12                                         = 0x8CEC
	EXT_cmyka                                                  = 1
	INTENSITY8                                                 = 0x804B
	PROGRAM_RESIDENT_NV                                        = 0x8647
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	ALL_STATIC_DATA_IBM                                        = 103060
	ALPHA8                                                     = 0x803C
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	IMAGE_SCALE_Y_HP                                           = 0x8156
	RGB16F                                                     = 0x881B
	DRAW_BUFFER4_ARB                                           = 0x8829
	REG_21_ATI                                                 = 0x8936
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	DRAW_BUFFER5_ARB                                           = 0x882A
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	SGIX_vertex_preclip                                        = 1
	COLOR_MATRIX_SGI                                           = 0x80B1
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	TEXTURE0_ARB                                               = 0x84C0
	FLOAT_MAT2x4                                               = 0x8B66
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	STREAM_DRAW_ARB                                            = 0x88E0
	DOUBLE_MAT4_EXT                                            = 0x8F48
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	SIGNALED                                                   = 0x9119
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	DRAW_BUFFER1_ARB                                           = 0x8826
	LUMINANCE12                                                = 0x8041
	CONSTANT_COLOR_EXT                                         = 0x8001
	TEXTURE31                                                  = 0x84DF
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	V3F                                                        = 0x2A21
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	DEBUG_SOURCE_OTHER                                         = 0x824B
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	SIGNED_RGBA8_NV                                            = 0x86FC
	REG_19_ATI                                                 = 0x8934
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	DEBUG_SEVERITY_LOW                                         = 0x9148
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	DEPTH_COMPONENT32                                          = 0x81A7
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	REG_13_ATI                                                 = 0x892E
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	CONVOLUTION_1D_EXT                                         = 0x8010
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	LINES_ADJACENCY_ARB                                        = 0x000A
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	MATRIX29_ARB                                               = 0x88DD
	REG_1_ATI                                                  = 0x8922
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	FLOAT_MAT4                                                 = 0x8B5C
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	DOUBLE_MAT2_EXT                                            = 0x8F46
	IS_ROW_MAJOR                                               = 0x9300
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	UNSIGNED_INT64_NV                                          = 0x140F
	COMPRESSED_LUMINANCE                                       = 0x84EA
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	SGIX_icc_texture                                           = 1
	CONSTANT_ATTENUATION                                       = 0x1207
	POINT_SIZE_MIN                                             = 0x8126
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	IUI_N3F_V2F_EXT                                            = 0x81AF
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	DEPTH32F_STENCIL8                                          = 0x8CAD
	LINE_TO_NV                                                 = 0x04
	TEXTURE_BLUE_SIZE                                          = 0x805E
	MAP2_BINORMAL_EXT                                          = 0x8447
	PALETTE8_RGBA8_OES                                         = 0x8B96
	R16_SNORM                                                  = 0x8F98
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	RGB_SCALE                                                  = 0x8573
	YCBCR_MESA                                                 = 0x8757
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	SGIX_fog_offset                                            = 1
	SUB_ATI                                                    = 0x8965
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	TEXTURE_WIDTH                                              = 0x1000
	IDENTITY_NV                                                = 0x862A
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	COPY                                                       = 0x1503
	RGBA_S3TC                                                  = 0x83A2
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	MATRIX5_ARB                                                = 0x88C5
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	TEXTURE_MAX_LEVEL                                          = 0x813D
	MATRIX3_NV                                                 = 0x8633
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	CONTINUOUS_AMD                                             = 0x9007
	UNIFORM                                                    = 0x92E1
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	ACCUM_GREEN_BITS                                           = 0x0D59
	RGB                                                        = 0x1907
	DISCARD_ATI                                                = 0x8763
	RGB_FLOAT16_APPLE                                          = 0x881B
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	VECTOR_EXT                                                 = 0x87BF
	CON_16_ATI                                                 = 0x8951
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	DOUBLE_MAT3x4                                              = 0x8F4C
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	SOURCE2_ALPHA                                              = 0x858A
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	MATRIX25_ARB                                               = 0x88D9
	RGB16I_EXT                                                 = 0x8D89
	TEXTURE14_ARB                                              = 0x84CE
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	CURRENT_FOG_COORD                                          = 0x8453
	DEPTH_CLAMP_NV                                             = 0x864F
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	SOURCE0_ALPHA                                              = 0x8588
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	CLIP_DISTANCE0                                             = 0x3000
	TEXTURE_CUBE_MAP                                           = 0x8513
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	PATH_FORMAT_PS_NV                                          = 0x9071
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	COMBINER6_NV                                               = 0x8556
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	NEGATIVE_Z_EXT                                             = 0x87DB
	RGB_INTEGER                                                = 0x8D98
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	RENDERER                                                   = 0x1F01
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	MATRIX17_ARB                                               = 0x88D1
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	LINEAR_DETAIL_SGIS                                         = 0x8097
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	GL_3DC_XY_AMD                                              = 0x87FA
	BUFFER_MAPPED_OES                                          = 0x88BC
	DECODE_EXT                                                 = 0x8A49
	DOUBLE_VEC3                                                = 0x8FFD
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	TIME_ELAPSED                                               = 0x88BF
	RGBA_INTEGER_EXT                                           = 0x8D99
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	STREAM_READ_ARB                                            = 0x88E1
	FOG                                                        = 0x0B60
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	MODELVIEW7_ARB                                             = 0x8727
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	PRIMITIVE_ID_NV                                            = 0x8C7C
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	RGB12                                                      = 0x8053
	QUAD_ALPHA8_SGIS                                           = 0x811F
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	TEXTURE_SHADER_NV                                          = 0x86DE
	DRAW_BUFFER4_ATI                                           = 0x8829
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	DEPTH_COMPONENT                                            = 0x1902
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	FOG_OFFSET_SGIX                                            = 0x8198
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	RGB8_SNORM                                                 = 0x8F96
	CURRENT_MATRIX_NV                                          = 0x8641
	EXT_blend_logic_op                                         = 1
	POINT_SIZE_RANGE                                           = 0x0B12
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	TEXTURE19_ARB                                              = 0x84D3
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	UNSIGNALED                                                 = 0x9118
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	EXT_blend_subtract                                         = 1
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	SPRITE_SGIX                                                = 0x8148
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	DSDT_NV                                                    = 0x86F5
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	COMPUTE_SUBROUTINE                                         = 0x92ED
	BACK_RIGHT                                                 = 0x0403
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	CURRENT_VERTEX_EXT                                         = 0x87E2
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	V2F                                                        = 0x2A20
	TEXTURE23_ARB                                              = 0x84D7
	MODELVIEW17_ARB                                            = 0x8731
	DRAW_BUFFER12_ATI                                          = 0x8831
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	SOURCE1_ALPHA_EXT                                          = 0x8589
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	DRAW_BUFFER0_ARB                                           = 0x8825
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	CON_12_ATI                                                 = 0x894D
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	GREEN_INTEGER_EXT                                          = 0x8D95
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	T2F_IUI_V2F_EXT                                            = 0x81B1
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	BOOL_ARB                                                   = 0x8B56
	SYNC_FENCE                                                 = 0x9116
	KEEP                                                       = 0x1E00
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	OPERAND2_RGB                                               = 0x8592
	RGB32UI_EXT                                                = 0x8D71
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	INTERLACE_OML                                              = 0x8980
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	RED_INTEGER_EXT                                            = 0x8D94
	MIRRORED_REPEAT_ARB                                        = 0x8370
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	OP_LOG_BASE_2_EXT                                          = 0x8792
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	BUFFER_DATA_SIZE                                           = 0x9303
	FUNC_SUBTRACT_EXT                                          = 0x800A
	MAX_SAMPLES_NV                                             = 0x8D57
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	INT8_VEC3_NV                                               = 0x8FE2
	TRANSLATE_X_NV                                             = 0x908E
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	CONSTANT_BORDER                                            = 0x8151
	BINORMAL_ARRAY_EXT                                         = 0x843A
	PRIMARY_COLOR_NV                                           = 0x852C
	DYNAMIC_COPY                                               = 0x88EA
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	COLOR_ATTACHMENT10                                         = 0x8CEA
	SGIX_depth_texture                                         = 1
	COLOR_RENDERABLE                                           = 0x8286
	OPERAND0_RGB_EXT                                           = 0x8590
	Y_EXT                                                      = 0x87D6
	REG_25_ATI                                                 = 0x893A
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	INDEX_BIT_PGI                                              = 0x00080000
	ALWAYS                                                     = 0x0207
	TEXTURE_BINDING_3D                                         = 0x806A
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	VERTEX_ID_NV                                               = 0x8C7B
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	TEXTURE1                                                   = 0x84C1
	VARIABLE_G_NV                                              = 0x8529
	QUADS                                                      = 0x0007
	CLIP_DISTANCE7                                             = 0x3007
	REG_15_ATI                                                 = 0x8930
	COMPRESSED_SRGB                                            = 0x8C48
	HISTOGRAM_SINK                                             = 0x802D
	POINT_SIZE_MAX                                             = 0x8127
	INVERTED_SCREEN_W_REND                                     = 0x8491
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	STREAM_COPY_ARB                                            = 0x88E2
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	SPHERE_MAP                                                 = 0x2402
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	INDEX                                                      = 0x8222
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	DECAL                                                      = 0x2101
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	RED_SCALE                                                  = 0x0D14
	MAP2_NORMAL                                                = 0x0DB2
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	LUMINANCE32I_EXT                                           = 0x8D86
	RED_SNORM                                                  = 0x8F90
	VERTEX_ARRAY                                               = 0x8074
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	COLOR_INDEX2_EXT                                           = 0x80E3
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	WRITE_ONLY_OES                                             = 0x88B9
	RGBA                                                       = 0x1908
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	SRC1_ALPHA                                                 = 0x8589
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	RGB32I                                                     = 0x8D83
	IMAGE_CUBE_EXT                                             = 0x9050
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	POLYGON_OFFSET_FILL                                        = 0x8037
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	DYNAMIC_ATI                                                = 0x8761
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	SLUMINANCE8                                                = 0x8C47
	MEDIUM_FLOAT                                               = 0x8DF1
	SCISSOR_TEST                                               = 0x0C11
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	FEEDBACK                                                   = 0x1C01
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	INDEX_MODE                                                 = 0x0C30
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	SYNC_CONDITION                                             = 0x9113
	EXT_blend_minmax                                           = 1
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	MAD_ATI                                                    = 0x8968
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	DT_SCALE_NV                                                = 0x8711
	PROGRAM_FORMAT_ARB                                         = 0x8876
	SAMPLER_1D_ARB                                             = 0x8B5D
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	CLIP_PLANE3                                                = 0x3003
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	FASTEST                                                    = 0x1101
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	SAMPLES_EXT                                                = 0x80A9
	SOURCE2_ALPHA_EXT                                          = 0x858A
	VERTEX_SOURCE_ATI                                          = 0x8774
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	COLOR_TABLE_SGI                                            = 0x80D0
	STACK_UNDERFLOW                                            = 0x0504
	SRGB8_EXT                                                  = 0x8C41
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	GREEN_BITS                                                 = 0x0D53
	IMAGE_SCALE_X_HP                                           = 0x8155
	SIGNED_HILO16_NV                                           = 0x86FA
	OP_FLOOR_EXT                                               = 0x878F
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	SRC_ALPHA_SATURATE                                         = 0x0308
	REFLECTION_MAP_NV                                          = 0x8512
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	VERSION_2_0                                                = 1
	COLOR_TABLE_FORMAT                                         = 0x80D8
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	MODELVIEW21_ARB                                            = 0x8735
	FLOAT_MAT4x2                                               = 0x8B69
	RGBA16_SNORM                                               = 0x8F9B
	VIDEO_BUFFER_NV                                            = 0x9020
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	RGBA8_EXT                                                  = 0x8058
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	NORMAL_MAP_OES                                             = 0x8511
	DEPTH_CLAMP                                                = 0x864F
	INT_VEC3_ARB                                               = 0x8B54
	ACTIVE_UNIFORMS                                            = 0x8B86
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	INTENSITY8I_EXT                                            = 0x8D91
	READ_BUFFER                                                = 0x0C02
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	TEXTURE                                                    = 0x1702
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	TEXTURE19                                                  = 0x84D3
	RED_BIT_ATI                                                = 0x00000001
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	RGBA_FLOAT16_ATI                                           = 0x881A
	CLAMP_READ_COLOR                                           = 0x891C
	SAMPLER_BUFFER_AMD                                         = 0x9001
	SCALE_BY_TWO_NV                                            = 0x853E
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	BGRA                                                       = 0x80E1
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	VARIABLE_E_NV                                              = 0x8527
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	WRITE_DISCARD_NV                                           = 0x88BE
	RGB5_A1_EXT                                                = 0x8057
	T2F_IUI_V3F_EXT                                            = 0x81B2
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	DRAW_BUFFER2_ATI                                           = 0x8827
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	SHADER_BINARY_VIV                                          = 0x8FC4
	FIXED                                                      = 0x140C
	AND                                                        = 0x1501
	S                                                          = 0x2000
	PROGRAM_STRING_ARB                                         = 0x8628
	OP_POWER_EXT                                               = 0x8793
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	UNPACK_ALIGNMENT                                           = 0x0CF5
	RG16F_EXT                                                  = 0x822F
	TRIANGLE_MESH_SUN                                          = 0x8615
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	LINES                                                      = 0x0001
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	MINMAX_EXT                                                 = 0x802E
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	COLOR_INDEXES                                              = 0x1603
	COORD_REPLACE_ARB                                          = 0x8862
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	VERTEX_PROGRAM_ARB                                         = 0x8620
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	BOOL_VEC4                                                  = 0x8B59
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	ALPHA_SNORM                                                = 0x9010
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	PHONG_WIN                                                  = 0x80EA
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	CON_24_ATI                                                 = 0x8959
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	QUERY_OBJECT_EXT                                           = 0x9153
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	INTENSITY4                                                 = 0x804A
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	RENDERBUFFER_OES                                           = 0x8D41
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	VIEW_CLASS_16_BITS                                         = 0x82CA
	INTERLACE_READ_INGR                                        = 0x8568
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	NORMAL_MAP_ARB                                             = 0x8511
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	LUMINANCE4_EXT                                             = 0x803F
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	TRIANGLE_LIST_SUN                                          = 0x81D7
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	C4UB_V3F                                                   = 0x2A23
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	RGBA16F                                                    = 0x881A
	STATIC_COPY                                                = 0x88E6
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	RGBA_DXT5_S3TC                                             = 0x83A4
	SLUMINANCE_NV                                              = 0x8C46
	RED_MAX_CLAMP_INGR                                         = 0x8564
	TEXTURE21_ARB                                              = 0x84D5
	SOURCE3_RGB_NV                                             = 0x8583
	RELATIVE_ARC_TO_NV                                         = 0xFF
	MAX_NAME_LENGTH                                            = 0x92F6
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	COLOR_SUM_EXT                                              = 0x8458
	MATRIX_EXT                                                 = 0x87C0
	CON_0_ATI                                                  = 0x8941
	MUL_ATI                                                    = 0x8964
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	INDEX_OFFSET                                               = 0x0D13
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	BUMP_ENVMAP_ATI                                            = 0x877B
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	SWIZZLE_STQ_ATI                                            = 0x8977
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	LESS                                                       = 0x0201
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	VERTEX_STREAM5_ATI                                         = 0x8771
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	FRAMEBUFFER_EXT                                            = 0x8D40
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	NORMAL_ARRAY_EXT                                           = 0x8075
	PROGRAM                                                    = 0x82E2
	VERTEX_STREAM0_ATI                                         = 0x876C
	CON_4_ATI                                                  = 0x8945
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	DONT_CARE                                                  = 0x1100
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	MODELVIEW6_ARB                                             = 0x8726
	MATRIX6_ARB                                                = 0x88C6
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	VIEW_CLASS_8_BITS                                          = 0x82CB
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	STANDARD_FONT_NAME_NV                                      = 0x9072
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	CURRENT_QUERY_ARB                                          = 0x8865
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	INT8_NV                                                    = 0x8FE0
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	CONSTANT_COLOR                                             = 0x8001
	ASYNC_MARKER_SGIX                                          = 0x8329
	PIXEL_COUNT_NV                                             = 0x8866
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	NORMAL_BIT_PGI                                             = 0x08000000
	RGB10                                                      = 0x8052
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	INT16_VEC2_NV                                              = 0x8FE5
	POLYGON_SMOOTH                                             = 0x0B41
	MODELVIEW3_ARB                                             = 0x8723
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	DRAW_BUFFER13                                              = 0x8832
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	MAX_LABEL_LENGTH                                           = 0x82E8
	FOG_COORD                                                  = 0x8451
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	INTENSITY16_SNORM                                          = 0x901B
	HALF_FLOAT                                                 = 0x140B
	T2F_C4UB_V3F                                               = 0x2A29
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	MAX_VERTEX_STREAMS                                         = 0x8E71
	GL_4_BYTES                                                 = 0x1409
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	NUM_PASSES_ATI                                             = 0x8970
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	CON_13_ATI                                                 = 0x894E
	COMBINER_BIAS_NV                                           = 0x8549
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	LERP_ATI                                                   = 0x8969
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	PASS_THROUGH_TOKEN                                         = 0x0700
	RESCALE_NORMAL                                             = 0x803A
	TEXTURE25_ARB                                              = 0x84D9
	BUFFER_BINDING                                             = 0x9302
	SGIX_tag_sample_buffer                                     = 1
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	DRAW_BUFFER7_ARB                                           = 0x882C
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	FOG_COORDINATE_EXT                                         = 0x8451
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	SAMPLER_2D_SHADOW                                          = 0x8B62
	INT_SAMPLER_1D                                             = 0x8DC9
	CLIP_DISTANCE3                                             = 0x3003
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	CUBIC_CURVE_TO_NV                                          = 0x0C
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	POINT_SMOOTH_HINT                                          = 0x0C51
	ACCUM_BLUE_BITS                                            = 0x0D5A
	VARIABLE_A_NV                                              = 0x8523
	EMBOSS_LIGHT_NV                                            = 0x855D
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	CON_25_ATI                                                 = 0x895A
	GL_2_BYTES                                                 = 0x1407
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	COMBINER_MAPPING_NV                                        = 0x8543
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	SAMPLER_2D                                                 = 0x8B5E
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	INT8_VEC2_NV                                               = 0x8FE1
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	SGIS_texture_select                                        = 1
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	BITMAP                                                     = 0x1A00
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	COMBINER5_NV                                               = 0x8555
	CURRENT_MATRIX_ARB                                         = 0x8641
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	RETAINED_APPLE                                             = 0x8A1B
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	IMAGE_CUBE                                                 = 0x9050
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	TEXTURE27                                                  = 0x84DB
	TEXTURE31_ARB                                              = 0x84DF
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	COLOR_ATTACHMENT7                                          = 0x8CE7
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	INCR                                                       = 0x1E02
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	INVERSE_NV                                                 = 0x862B
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	CMYKA_EXT                                                  = 0x800D
	HISTOGRAM_RED_SIZE                                         = 0x8028
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	VERTEX_STREAM1_ATI                                         = 0x876D
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	MODELVIEW                                                  = 0x1700
	COLOR                                                      = 0x1800
	MATRIX2_NV                                                 = 0x8632
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	MATRIX_STRIDE                                              = 0x92FF
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	OP_ADD_EXT                                                 = 0x8787
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	TEXTURE18                                                  = 0x84D2
	COMBINER3_NV                                               = 0x8553
	SLUMINANCE_ALPHA                                           = 0x8C44
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	CON_7_ATI                                                  = 0x8948
	LARGE_CW_ARC_TO_NV                                         = 0x18
	ARRAY_STRIDE                                               = 0x92FE
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	ALPHA16_SNORM                                              = 0x9018
	SAMPLER_OBJECT_AMD                                         = 0x9155
	SGIX_async_pixel                                           = 1
	SAMPLE_COVERAGE                                            = 0x80A0
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	SGIS_point_parameters                                      = 1
	LINE_BIT                                                   = 0x00000004
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	NORMALIZE                                                  = 0x0BA1
	TEXTURE_GEN_R                                              = 0x0C62
	DOT3_RGBA                                                  = 0x86AF
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	FRONT_RIGHT                                                = 0x0401
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	VERTEX_SUBROUTINE                                          = 0x92E8
	MIRRORED_REPEAT_OES                                        = 0x8370
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	LIST_BASE                                                  = 0x0B32
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	TEXTURE28_ARB                                              = 0x84DC
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	MULTISAMPLE_BIT                                            = 0x20000000
	MAP1_VERTEX_4                                              = 0x0D98
	RED                                                        = 0x1903
	RG8                                                        = 0x822B
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	SCREEN_COORDINATES_REND                                    = 0x8490
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	FOG_COORDINATE_ARRAY                                       = 0x8457
	MOVE_TO_RESETS_NV                                          = 0x90B5
	DEPTH_EXT                                                  = 0x1801
	LUMINANCE8                                                 = 0x8040
	COLOR_TABLE_SCALE                                          = 0x80D6
	SCALAR_EXT                                                 = 0x87BE
	RGB9_E5_EXT                                                = 0x8C3D
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	MODELVIEW0_EXT                                             = 0x1700
	RGB4_EXT                                                   = 0x804F
	COLOR_ATTACHMENT14                                         = 0x8CEE
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	TEXTURE_BORDER_COLOR                                       = 0x1004
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	SAMPLE_MASK_NV                                             = 0x8E51
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	BGR_INTEGER                                                = 0x8D9A
	VERSION_1_3                                                = 1
	COMBINER7_NV                                               = 0x8557
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	SGIS_point_line_texgen                                     = 1
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	SHADER_COMPILER                                            = 0x8DFA
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	COLOR_EXT                                                  = 0x1800
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	MATRIX13_ARB                                               = 0x88CD
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	COUNT_UP_NV                                                = 0x9088
	MAP1_VERTEX_3                                              = 0x0D97
	CONVOLUTION_2D                                             = 0x8011
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	CONSTANT_ARB                                               = 0x8576
	MAGNITUDE_BIAS_NV                                          = 0x8718
	FLOAT_RG_NV                                                = 0x8881
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	VERTICAL_LINE_TO_NV                                        = 0x08
	PIXEL_MODE_BIT                                             = 0x00000020
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	STENCIL_INDEX8_EXT                                         = 0x8D48
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	INDEX_ARRAY_POINTER                                        = 0x8091
	DEBUG_TYPE_MARKER                                          = 0x8268
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	R32I                                                       = 0x8235
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	TEXTURE16                                                  = 0x84D0
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	INDEX_ARRAY                                                = 0x8077
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	LOGIC_OP_MODE                                              = 0x0BF0
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	COMPRESSED_R11_EAC                                         = 0x9270
	RIGHT                                                      = 0x0407
	CONSTANT_ALPHA                                             = 0x8003
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	BGR_EXT                                                    = 0x80E0
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	TEXTURE11                                                  = 0x84CB
	FLOAT_R16_NV                                               = 0x8884
	BUFFER_MAPPED_ARB                                          = 0x88BC
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	CLIP_PLANE4                                                = 0x3004
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	IUI_V3F_EXT                                                = 0x81AE
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	RG8I                                                       = 0x8237
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	RGBA8UI_EXT                                                = 0x8D7C
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	SQUARE_NV                                                  = 0x90A3
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	CURRENT_QUERY_EXT                                          = 0x8865
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	VARIABLE_C_NV                                              = 0x8525
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	INT_IMAGE_CUBE                                             = 0x905B
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	REPLACE                                                    = 0x1E01
	SAMPLES_3DFX                                               = 0x86B4
	DRAW_BUFFER13_NV                                           = 0x8832
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	CLIP_DISTANCE_NV                                           = 0x8C7A
	IMAGE_2D_ARRAY                                             = 0x9053
	STORAGE_CACHED_APPLE                                       = 0x85BE
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	DRAW_BUFFER10_NV                                           = 0x882F
	MATRIX_PALETTE_ARB                                         = 0x8840
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	WRITE_ONLY                                                 = 0x88B9
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	LOGIC_OP                                                   = 0x0BF1
	OR_INVERTED                                                = 0x150D
	HI_SCALE_NV                                                = 0x870E
	REG_18_ATI                                                 = 0x8933
	COLOR_ATTACHMENT6                                          = 0x8CE6
	DOUBLE                                                     = 0x140A
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	MAX_EXT                                                    = 0x8008
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	DRAW_BUFFER6_ATI                                           = 0x882B
	OP_EXP_BASE_2_EXT                                          = 0x8791
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	SURFACE_REGISTERED_NV                                      = 0x86FD
	RGB32F                                                     = 0x8815
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	DRAW_BUFFER2_ARB                                           = 0x8827
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	TESS_CONTROL_SHADER                                        = 0x8E88
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	QUADRATIC_ATTENUATION                                      = 0x1209
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	GEOMETRY_TEXTURE                                           = 0x829E
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	COLOR4_BIT_PGI                                             = 0x00020000
	CON_9_ATI                                                  = 0x894A
	UNSIGNED_INT8_NV                                           = 0x8FEC
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	OBJECT_TYPE_ARB                                            = 0x8B4E
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	BUFFER_MAP_LENGTH                                          = 0x9120
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	DEPTH_TEXTURE_MODE                                         = 0x884B
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	ALPHA8UI_EXT                                               = 0x8D7E
	SUCCESS_NV                                                 = 0x902F
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	FLOAT_RGBA_NV                                              = 0x8883
	MATRIX27_ARB                                               = 0x88DB
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	STENCIL_INDEX1_OES                                         = 0x8D46
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	SOURCE0_ALPHA_EXT                                          = 0x8588
	MAX_DRAW_BUFFERS                                           = 0x8824
	REG_4_ATI                                                  = 0x8925
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	UTF16_NV                                                   = 0x909B
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	OPERAND0_ALPHA_ARB                                         = 0x8598
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	SYNC_CL_EVENT_ARB                                          = 0x8240
	DEPTH_STENCIL_MESA                                         = 0x8750
	FLOAT_RG16_NV                                              = 0x8886
	RGB16UI                                                    = 0x8D77
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	EXT_point_parameters                                       = 1
	CURRENT_NORMAL                                             = 0x0B02
	INTENSITY12                                                = 0x804C
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	VERTEX_SHADER_EXT                                          = 0x8780
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	LUMINANCE                                                  = 0x1909
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	SLIM8U_SGIX                                                = 0x831D
	TEXTURE11_ARB                                              = 0x84CB
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	MATRIX23_ARB                                               = 0x88D7
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	FLAT                                                       = 0x1D00
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	PRIMARY_COLOR_ARB                                          = 0x8577
	PROGRAM_TARGET_NV                                          = 0x8646
	MIRROR_CLAMP_ATI                                           = 0x8742
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	INTENSITY16I_EXT                                           = 0x8D8B
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	CON_3_ATI                                                  = 0x8944
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	EXT_rescale_normal                                         = 1
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	COLOR_ATTACHMENT1                                          = 0x8CE1
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	LAYOUT_LINEAR_INTEL                                        = 1
	LIST_INDEX                                                 = 0x0B33
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	GL_422_REV_EXT                                             = 0x80CD
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	TEXTURE29_ARB                                              = 0x84DD
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	DRAW_BUFFER12                                              = 0x8831
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	INVALID_INDEX                                              = 0xFFFFFFFF
	QUERY_WAIT                                                 = 0x8E13
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	INT64_NV                                                   = 0x140E
	PACK_ROW_LENGTH                                            = 0x0D02
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	FLOAT_MAT2x3                                               = 0x8B65
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	BLEND_COLOR                                                = 0x8005
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	SGIX_texture_multi_buffer                                  = 1
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	NEVER                                                      = 0x0200
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	TEXTURE3_ARB                                               = 0x84C3
	CON_31_ATI                                                 = 0x8960
	MULTISAMPLE_ARB                                            = 0x809D
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	SYNC_FENCE_APPLE                                           = 0x9116
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	CON_11_ATI                                                 = 0x894C
	PURGEABLE_APPLE                                            = 0x8A1D
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	INT_IMAGE_2D                                               = 0x9058
	UTF8_NV                                                    = 0x909A
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	TEXTURE12_ARB                                              = 0x84CC
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	CON_17_ATI                                                 = 0x8952
	CURRENT_PROGRAM                                            = 0x8B8D
	TEXTURE_GEN_Q                                              = 0x0C63
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	TEXTURE10_ARB                                              = 0x84CA
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	COLOR_INDEX8_EXT                                           = 0x80E5
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	PERFMON_RESULT_AMD                                         = 0x8BC6
	STACK_OVERFLOW                                             = 0x0503
	NORMAL_MAP_NV                                              = 0x8511
	MODELVIEW9_ARB                                             = 0x8729
	MULTISAMPLE_3DFX                                           = 0x86B2
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	GEOMETRY_SHADER                                            = 0x8DD9
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	REPLACE_EXT                                                = 0x8062
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	NEGATIVE_W_EXT                                             = 0x87DC
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	VERTEX23_BIT_PGI                                           = 0x00000004
	TEXTURE_MIN_FILTER                                         = 0x2801
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	TIMESTAMP                                                  = 0x8E28
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	POINT_SIZE_MAX_ARB                                         = 0x8127
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MODELVIEW24_ARB                                            = 0x8738
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	DYNAMIC_COPY_ARB                                           = 0x88EA
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	BLEND_EQUATION_RGB                                         = 0x8009
	MODELVIEW20_ARB                                            = 0x8734
	LIGHT2                                                     = 0x4002
	VARIABLE_F_NV                                              = 0x8528
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	FLOAT_RGB32_NV                                             = 0x8889
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	R16I                                                       = 0x8233
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	RGBA32I                                                    = 0x8D82
	NEXT_BUFFER_NV                                             = -2
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	FOG_START                                                  = 0x0B63
	RED_BITS                                                   = 0x0D52
	COLOR_ARRAY                                                = 0x8076
	EYE_RADIAL_NV                                              = 0x855B
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	ATTACHED_SHADERS                                           = 0x8B85
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	PATH_FILL_MODE_NV                                          = 0x9080
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	ACCUM_RED_BITS                                             = 0x0D58
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	TRUE                                                       = 1
	COLOR_ARRAY_STRIDE                                         = 0x8083
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	MINMAX_SINK                                                = 0x8030
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	SGIS_multisample                                           = 1
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	FOG_DENSITY                                                = 0x0B62
	TEXTURE4                                                   = 0x84C4
	DRAW_BUFFER3_NV                                            = 0x8828
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	READ_FRAMEBUFFER                                           = 0x8CA8
	PRIMITIVE_RESTART                                          = 0x8F9D
	INT_IMAGE_1D                                               = 0x9057
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	SIGNED_ALPHA_NV                                            = 0x8705
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	DEPTH_BUFFER_BIT                                           = 0x00000100
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	TEXTURE13                                                  = 0x84CD
	RGBA_FLOAT32_APPLE                                         = 0x8814
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	CULL_FRAGMENT_NV                                           = 0x86E7
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	BLUE_SCALE                                                 = 0x0D1A
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	CUBIC_HP                                                   = 0x815F
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	CAVEAT_SUPPORT                                             = 0x82B8
	REFLECTION_MAP_OES                                         = 0x8512
	OPERAND0_RGB                                               = 0x8590
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	PROXY_HISTOGRAM                                            = 0x8025
	INTENSITY16_EXT                                            = 0x804D
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	SRC_COLOR                                                  = 0x0300
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	LIGHT1                                                     = 0x4001
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	MATRIX28_ARB                                               = 0x88DC
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	DEBUG_SOURCE_API                                           = 0x8246
	SOURCE0_RGB                                                = 0x8580
	MODELVIEW11_ARB                                            = 0x872B
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	PATH_END_CAPS_NV                                           = 0x9076
	SYNC_FLAGS_APPLE                                           = 0x9115
	MIN_EXT                                                    = 0x8007
	DECR_WRAP_EXT                                              = 0x8508
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	TEXTURE_RED_TYPE                                           = 0x8C10
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	POLYGON_OFFSET_EXT                                         = 0x8037
	EXT_copy_texture                                           = 1
	PROJECTION_MATRIX                                          = 0x0BA7
	MIRRORED_REPEAT_IBM                                        = 0x8370
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	EXT_blend_color                                            = 1
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	SCISSOR_BIT                                                = 0x00080000
	TEXTURE_ENV_COLOR                                          = 0x2201
	IMAGE_BINDING_NAME                                         = 0x8F3A
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	TEXTURE20                                                  = 0x84D4
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	STENCIL_INDEX4_EXT                                         = 0x8D47
	STENCIL_INDEX16_EXT                                        = 0x8D49
	T                                                          = 0x2001
	ALPHA_MIN_SGIX                                             = 0x8320
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	GPU_ADDRESS_NV                                             = 0x8F34
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	PROXY_COLOR_TABLE                                          = 0x80D3
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	DEPTH_STENCIL_OES                                          = 0x84F9
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	STENCIL                                                    = 0x1802
	MODELVIEW1_EXT                                             = 0x850A
	OPERAND2_RGB_EXT                                           = 0x8592
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	QUERY_RESULT                                               = 0x8866
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	NAND                                                       = 0x150E
	OBJECT_PLANE                                               = 0x2501
	BLEND_COLOR_EXT                                            = 0x8005
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	TEXTURE6                                                   = 0x84C6
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	RGB_SNORM                                                  = 0x8F92
	TEXTURE_MAX_LOD                                            = 0x813B
	CONSTANT_COLOR0_NV                                         = 0x852A
	REG_14_ATI                                                 = 0x892F
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	TEXTURE2                                                   = 0x84C2
	OP_FRAC_EXT                                                = 0x8789
	DRAW_BUFFER4                                               = 0x8829
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	INT_SAMPLER_CUBE                                           = 0x8DCC
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	HILO16_NV                                                  = 0x86F8
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	UNIFORM_BUFFER                                             = 0x8A11
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	TRANSLATE_Y_NV                                             = 0x908F
	SYNC_STATUS_APPLE                                          = 0x9114
	MAP1_NORMAL                                                = 0x0D92
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	DEPTH_RENDERABLE                                           = 0x8287
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	DRAW_BUFFER0                                               = 0x8825
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	EXPAND_NORMAL_NV                                           = 0x8538
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	TESS_GEN_SPACING                                           = 0x8E77
	RG_INTEGER                                                 = 0x8228
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	RG8UI                                                      = 0x8238
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	BLUE                                                       = 0x1905
	ALPHA_MAX_SGIX                                             = 0x8321
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	CURRENT_COLOR                                              = 0x0B00
	LIGHT6                                                     = 0x4006
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	MATRIX20_ARB                                               = 0x88D4
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	GL_1PASS_EXT                                               = 0x80A1
	TEXTURE24                                                  = 0x84D8
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	GL_422_EXT                                                 = 0x80CC
	DRAW_BUFFER5                                               = 0x882A
	FIXED_ONLY                                                 = 0x891D
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	INTENSITY32UI_EXT                                          = 0x8D73
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	RECT_NV                                                    = 0xF6
	DRAW_BUFFER2                                               = 0x8827
	DRAW_BUFFER12_ARB                                          = 0x8831
	PACK_IMAGE_HEIGHT                                          = 0x806C
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	SYNC_FLAGS                                                 = 0x9115
	HISTOGRAM                                                  = 0x8024
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	OPERAND0_ALPHA_EXT                                         = 0x8598
	OP_SET_LT_EXT                                              = 0x878D
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	UNSIGNED_INT16_NV                                          = 0x8FF0
	WAIT_FAILED                                                = 0x911D
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	COLOR_ATTACHMENT2                                          = 0x8CE2
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	BGR_INTEGER_EXT                                            = 0x8D9A
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	SGIX_shadow                                                = 1
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	DISCRETE_AMD                                               = 0x9006
	SGIX_blend_alpha_minmax                                    = 1
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	FILE_NAME_NV                                               = 0x9074
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	EXP2                                                       = 0x0801
	FOG_MODE                                                   = 0x0B65
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	MAX                                                        = 0x8008
	R32F                                                       = 0x822E
	SIGNED_RGBA_NV                                             = 0x86FB
	MATRIX10_ARB                                               = 0x88CA
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	PROGRAM_INPUT                                              = 0x92E3
	SAMPLE_MASK_SGIS                                           = 0x80A0
	TEXTURE5                                                   = 0x84C5
	MODELVIEW29_ARB                                            = 0x873D
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	CURRENT_BIT                                                = 0x00000001
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	FUNC_ADD                                                   = 0x8006
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	REG_20_ATI                                                 = 0x8935
	RENDERBUFFER                                               = 0x8D41
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	GL_4PASS_2_SGIS                                            = 0x80A6
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	RGBA8_OES                                                  = 0x8058
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	MVP_MATRIX_EXT                                             = 0x87E3
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	MATRIX22_ARB                                               = 0x88D6
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	ALPHA32F_ARB                                               = 0x8816
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	OP_MIN_EXT                                                 = 0x878B
	FLOAT_MAT4x3                                               = 0x8B6A
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	TEXTURE_WRAP_R_OES                                         = 0x8072
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	FRAGMENT_TEXTURE                                           = 0x829F
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	SGIX_texture_coordinate_clamp                              = 1
	RGBA2_EXT                                                  = 0x8055
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	MAP_TESSELLATION_NV                                        = 0x86C2
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	MATRIX0_ARB                                                = 0x88C0
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	MATRIX16_ARB                                               = 0x88D0
	IMAGE_2D                                                   = 0x904D
	FIRST_TO_REST_NV                                           = 0x90AF
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	SGIX_polynomial_ffd                                        = 1
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	SUBTRACT                                                   = 0x84E7
	YCBCR_422_APPLE                                            = 0x85B9
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	TEXTURE_3D_EXT                                             = 0x806F
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	DEPTH_COMPONENT24                                          = 0x81A6
	ADD_SIGNED_ARB                                             = 0x8574
	BUFFER_SIZE                                                = 0x8764
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	SYNC_OBJECT_APPLE                                          = 0x8A53
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	RED_BIAS                                                   = 0x0D15
	HISTOGRAM_SINK_EXT                                         = 0x802D
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	COLOR_SUM_ARB                                              = 0x8458
	OPERAND3_ALPHA_NV                                          = 0x859B
	CON_28_ATI                                                 = 0x895D
	BLOCK_INDEX                                                = 0x92FD
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	EXT_texture                                                = 1
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	STENCIL_INDEX16                                            = 0x8D49
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	RGB32F_ARB                                                 = 0x8815
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	POINT                                                      = 0x1B00
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	COMBINE_RGB_EXT                                            = 0x8571
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	OFFSET                                                     = 0x92FC
	TEXTURE_WRAP_S                                             = 0x2802
	PROXY_TEXTURE_1D                                           = 0x8063
	FENCE_CONDITION_NV                                         = 0x84F4
	COMBINE_ALPHA                                              = 0x8572
	OPERAND1_ALPHA                                             = 0x8599
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	FILL                                                       = 0x1B02
	GL_2PASS_0_EXT                                             = 0x80A2
	MINOR_VERSION                                              = 0x821C
	REFLECTION_MAP_EXT                                         = 0x8512
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	SLUMINANCE                                                 = 0x8C46
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	BUFFER_SIZE_ARB                                            = 0x8764
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	DRAW_BUFFER1                                               = 0x8826
	TIME_ELAPSED_EXT                                           = 0x88BF
	PACK_RESAMPLE_OML                                          = 0x8984
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	EQUIV                                                      = 0x1509
	DRAW_BUFFER10                                              = 0x882F
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	READ_ONLY                                                  = 0x88B8
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	STENCIL_WRITEMASK                                          = 0x0B98
	TABLE_TOO_LARGE                                            = 0x8031
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	LUMINANCE32UI_EXT                                          = 0x8D74
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	SOURCE1_ALPHA                                              = 0x8589
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	POLYGON_MODE                                               = 0x0B40
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	COMPRESSED_RGBA                                            = 0x84EE
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	TRANSFORM_FEEDBACK                                         = 0x8E22
	POLYGON_BIT                                                = 0x00000008
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	MODELVIEW5_ARB                                             = 0x8725
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	QUERY_RESULT_ARB                                           = 0x8866
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	TEXTURE_2D                                                 = 0x0DE1
	SINGLE_COLOR_EXT                                           = 0x81F9
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	COMPUTE_SHADER_BIT                                         = 0x00000020
	SIGNED_INTENSITY_NV                                        = 0x8707
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	RGBA16I                                                    = 0x8D88
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	LINE_STRIP_ADJACENCY                                       = 0x000B
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COMBINE_RGB_ARB                                            = 0x8571
	COLOR_ATTACHMENT8                                          = 0x8CE8
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	IUI_N3F_V3F_EXT                                            = 0x81B0
	TEXTURE7_ARB                                               = 0x84C7
	TEXTURE20_ARB                                              = 0x84D4
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	PRESENT_TIME_NV                                            = 0x8E2A
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	DOT2_ADD_ATI                                               = 0x896C
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	SET_AMD                                                    = 0x874A
	MATRIX8_ARB                                                = 0x88C8
	PALETTE8_RGB8_OES                                          = 0x8B95
	RESTART_PATH_NV                                            = 0xF0
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	VERTEX_TEXTURE                                             = 0x829B
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	DEPTH24_STENCIL8                                           = 0x88F0
	PALETTE4_RGBA4_OES                                         = 0x8B93
	SGIX_sprite                                                = 1
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	MAP2_COLOR_4                                               = 0x0DB0
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	OUT_OF_MEMORY                                              = 0x0505
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	INDEX_SHIFT                                                = 0x0D12
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	DUDV_ATI                                                   = 0x8779
	DRAW_BUFFER0_NV                                            = 0x8825
	RGB5_A1                                                    = 0x8057
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	TRANSPOSE_NV                                               = 0x862C
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	WAIT_FAILED_APPLE                                          = 0x911D
	YCRCB_422_SGIX                                             = 0x81BB
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	UNSIGNED_INT                                               = 0x1405
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	SLIM12S_SGIX                                               = 0x831F
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	ADJACENT_PAIRS_NV                                          = 0x90AE
	MAP_READ_BIT_EXT                                           = 0x0001
	CLAMP                                                      = 0x2900
	FOG_COORD_ARRAY                                            = 0x8457
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	RETURN                                                     = 0x0102
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	TEXTURE_COORD_ARRAY                                        = 0x8078
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	STATE_RESTORE                                              = 0x8BDC
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	SPOT_EXPONENT                                              = 0x1205
	TEXTURE3                                                   = 0x84C3
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	COMBINER_INPUT_NV                                          = 0x8542
	OP_NEGATE_EXT                                              = 0x8783
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	LOAD                                                       = 0x0101
	COLOR_ARRAY_LIST_IBM                                       = 103072
	SET                                                        = 0x150F
	MODELVIEW14_ARB                                            = 0x872E
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	CON_14_ATI                                                 = 0x894F
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	INTENSITY                                                  = 0x8049
	RGB32I_EXT                                                 = 0x8D83
	INT_2_10_10_10_REV                                         = 0x8D9F
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	REFERENCE_PLANE_SGIX                                       = 0x817D
	OPERAND2_RGB_ARB                                           = 0x8592
	DOT3_RGB_ARB                                               = 0x86AE
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	PROVOKING_VERTEX                                           = 0x8E4F
	TEXTURE_GEN_MODE                                           = 0x2500
	RGB16_EXT                                                  = 0x8054
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	TEXTURE_VIEW                                               = 0x82B5
	REFLECTION_MAP_ARB                                         = 0x8512
	MEDIUM_INT                                                 = 0x8DF4
	RGBA8_SNORM                                                = 0x8F97
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	COLOR_ATTACHMENT5                                          = 0x8CE5
	SGIX_framezoom                                             = 1
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	LIGHT3                                                     = 0x4003
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	NORMAL_ARRAY_TYPE                                          = 0x807E
	MODELVIEW30_ARB                                            = 0x873E
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	OPERAND2_ALPHA_ARB                                         = 0x859A
	UNIFORM_BUFFER_START                                       = 0x8A29
	RG16F                                                      = 0x822F
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	EXT_convolution                                            = 1
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	REG_8_ATI                                                  = 0x8929
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	DOUBLE_MAT4x2                                              = 0x8F4D
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	LO_BIAS_NV                                                 = 0x8715
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	TRIANGULAR_NV                                              = 0x90A5
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	RGBA16F_ARB                                                = 0x881A
	BOOL                                                       = 0x8B56
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	COMPILE_AND_EXECUTE                                        = 0x1301
	EXTENSIONS                                                 = 0x1F03
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	TEXTURE26_ARB                                              = 0x84DA
	ALL_COMPLETED_NV                                           = 0x84F2
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	RGBA_MODE                                                  = 0x0C31
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	RGBA32I_EXT                                                = 0x8D82
	MULTISAMPLE_EXT                                            = 0x809D
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	SGIS_generate_mipmap                                       = 1
	COMPRESSED_ALPHA                                           = 0x84E9
	SOURCE2_RGB_EXT                                            = 0x8582
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	STATIC_ATI                                                 = 0x8760
	ELEMENT_ARRAY_ATI                                          = 0x8768
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	REG_12_ATI                                                 = 0x892D
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	RED_INTEGER                                                = 0x8D94
	TEXTURE_SAMPLES                                            = 0x9106
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	SGI_color_table                                            = 1
	PROGRAM_PIPELINE                                           = 0x82E4
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	SRGB8_NV                                                   = 0x8C41
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	INDEX_TEST_EXT                                             = 0x81B5
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	OR_REVERSE                                                 = 0x150B
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	SOURCE3_ALPHA_NV                                           = 0x858B
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	GL_3D_COLOR                                                = 0x0602
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	NORMAL_ARRAY_POINTER                                       = 0x808F
	BUFFER                                                     = 0x82E0
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	FLOAT_MAT3                                                 = 0x8B5B
	DELETE_STATUS                                              = 0x8B80
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	SHADER_IMAGE_STORE                                         = 0x82A5
	PRIMARY_COLOR                                              = 0x8577
	VIBRANCE_BIAS_NV                                           = 0x8719
	SIGNED_NORMALIZED                                          = 0x8F9C
	FAILURE_NV                                                 = 0x9030
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	QUERY_BUFFER_AMD                                           = 0x9192
	UNSIGNED_BYTE                                              = 0x1401
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	POLYGON_OFFSET_POINT                                       = 0x2A01
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	SOURCE0_RGB_EXT                                            = 0x8580
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	PACK_LSB_FIRST                                             = 0x0D01
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	RGB_S3TC                                                   = 0x83A0
	UNSIGNED_NORMALIZED                                        = 0x8C17
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	POINT_TOKEN                                                = 0x0701
	SAMPLER_2D_RECT                                            = 0x8B63
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	REPLACE_OLDEST_SUN                                         = 0x0003
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	SAMPLE_SHADING_ARB                                         = 0x8C36
	BLEND_DST                                                  = 0x0BE0
	TEXTURE29                                                  = 0x84DD
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	MAX_LIST_NESTING                                           = 0x0B31
	HALF_APPLE                                                 = 0x140B
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	DRAW_BUFFER6_NV                                            = 0x882B
	STENCIL_INDEX1                                             = 0x8D46
	LINE_SMOOTH_HINT                                           = 0x0C52
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	COUNTER_TYPE_AMD                                           = 0x8BC0
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	CULL_FACE                                                  = 0x0B44
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	QUERY_NO_WAIT                                              = 0x8E14
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	MULTISAMPLE_SGIS                                           = 0x809D
	COLOR_INDEX1_EXT                                           = 0x80E2
	R8UI                                                       = 0x8232
	SHADER_OPERATION_NV                                        = 0x86DF
	BUFFER_USAGE                                               = 0x8765
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	EXT_shared_texture_palette                                 = 1
	BLEND_SRC_RGB                                              = 0x80C9
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	POINT_SIZE_MIN_ARB                                         = 0x8126
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	DEBUG_OUTPUT                                               = 0x92E0
	COLOR3_BIT_PGI                                             = 0x00010000
	TEXTURE_PRIORITY                                           = 0x8066
	SLICE_ACCUM_SUN                                            = 0x85CC
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	EXT_abgr                                                   = 1
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	QUAD_STRIP                                                 = 0x0008
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	COLOR_ATTACHMENT0                                          = 0x8CE0
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	CONSTANT_BORDER_HP                                         = 0x8151
	RGBA_FLOAT16_APPLE                                         = 0x881A
	REG_27_ATI                                                 = 0x893C
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	SGIX_pixel_texture                                         = 1
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	MAP_STENCIL                                                = 0x0D11
	SEPARABLE_2D_EXT                                           = 0x8012
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	CON_29_ATI                                                 = 0x895E
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	BYTE                                                       = 0x1400
	SECONDARY_COLOR_NV                                         = 0x852D
	TEXTURE_RED_SIZE                                           = 0x805C
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	VERSION_2_1                                                = 1
	DOUBLE_EXT                                                 = 0x140A
	MAP2_TANGENT_EXT                                           = 0x8445
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	IMAGE_2D_RECT_EXT                                          = 0x904F
	SGIX_async_histogram                                       = 1
	DEPTH_BITS                                                 = 0x0D56
	UNPACK_SKIP_IMAGES                                         = 0x806D
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	BOOL_VEC2                                                  = 0x8B57
	SPECULAR                                                   = 0x1202
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	DECR_WRAP                                                  = 0x8508
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	CON_21_ATI                                                 = 0x8956
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	OUTPUT_FOG_EXT                                             = 0x87BD
	REG_10_ATI                                                 = 0x892B
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	DISCARD_NV                                                 = 0x8530
	LUMINANCE16I_EXT                                           = 0x8D8C
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	PACK_SKIP_IMAGES                                           = 0x806B
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	R16F                                                       = 0x822D
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	NEGATIVE_Y_EXT                                             = 0x87DA
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	OP_DOT4_EXT                                                = 0x8785
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	MAX_SAMPLES_EXT                                            = 0x8D57
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	SGIX_texture_scale_bias                                    = 1
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	GL_2PASS_1_SGIS                                            = 0x80A3
	ALPHA8I_EXT                                                = 0x8D90
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	AUX_BUFFERS                                                = 0x0C00
	GL_2PASS_0_SGIS                                            = 0x80A2
	VIEW_CLASS_24_BITS                                         = 0x82C9
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	FOG_COORDINATE_SOURCE                                      = 0x8450
	MODELVIEW1_ARB                                             = 0x850A
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	PALETTE4_RGB8_OES                                          = 0x8B90
	COPY_WRITE_BUFFER                                          = 0x8F37
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	HISTOGRAM_EXT                                              = 0x8024
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	OBJECT_TYPE                                                = 0x9112
	TRANSFORM_BIT                                              = 0x00001000
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	GL_4PASS_1_EXT                                             = 0x80A5
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	TESSELLATION_MODE_AMD                                      = 0x9004
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	SURFACE_STATE_NV                                           = 0x86EB
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	TYPE                                                       = 0x92FA
	MAP_READ_BIT                                               = 0x0001
	GL_4PASS_0_SGIS                                            = 0x80A4
	STORAGE_SHARED_APPLE                                       = 0x85BF
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	DOUBLE_VEC2                                                = 0x8FFC
	UNIFORM_BLOCK                                              = 0x92E2
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	T2F_N3F_V3F                                                = 0x2A2B
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	CONST_EYE_NV                                               = 0x86E5
	NEGATIVE_X_EXT                                             = 0x87D9
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	TEXTURE_ENV_MODE                                           = 0x2200
	MAX_WIDTH                                                  = 0x827E
	RGB_SCALE_EXT                                              = 0x8573
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	COMPRESSED_INTENSITY                                       = 0x84EC
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	SCISSOR_BOX                                                = 0x0C10
	PACK_SKIP_PIXELS                                           = 0x0D04
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	TEXTURE4_ARB                                               = 0x84C4
	TEXTURE24_ARB                                              = 0x84D8
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	CULL_FACE_MODE                                             = 0x0B45
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	INT_IMAGE_3D_EXT                                           = 0x9059
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	MATRIX19_ARB                                               = 0x88D3
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	VERSION_1_5                                                = 1
	DUAL_ALPHA8_SGIS                                           = 0x8111
	DRAW_BUFFER11_NV                                           = 0x8830
	RGB8I_EXT                                                  = 0x8D8F
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	BUFFER_MAP_OFFSET                                          = 0x9121
	INDEX_WRITEMASK                                            = 0x0C21
	DRAW_BUFFER7_NV                                            = 0x882C
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	FIELD_LOWER_NV                                             = 0x9023
	DEPTH_TEST                                                 = 0x0B71
	CLIP_PLANE5                                                = 0x3005
	INTENSITY12_EXT                                            = 0x804C
	SKIP_COMPONENTS2_NV                                        = -5
	PATH_COORD_COUNT_NV                                        = 0x909E
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	COLOR_ARRAY_POINTER                                        = 0x8090
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	REPLICATE_BORDER                                           = 0x8153
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	COMBINER_SCALE_NV                                          = 0x8548
	MOV_ATI                                                    = 0x8961
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	QUERY_WAIT_NV                                              = 0x8E13
	GREEN_BIAS                                                 = 0x0D19
	MODELVIEW12_ARB                                            = 0x872C
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	LUMINANCE8_SNORM                                           = 0x9015
	INT_IMAGE_3D                                               = 0x9059
	TRANSLATE_2D_NV                                            = 0x9090
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	COUNT_DOWN_NV                                              = 0x9089
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	ALPHA_BIAS                                                 = 0x0D1D
	TESS_CONTROL_TEXTURE                                       = 0x829C
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	CON_30_ATI                                                 = 0x895F
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	DRAW_BUFFER15                                              = 0x8834
	CON_1_ATI                                                  = 0x8942
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	EDGE_FLAG                                                  = 0x0B43
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	RGBA2                                                      = 0x8055
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	AUX2                                                       = 0x040B
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	COPY_PIXEL_TOKEN                                           = 0x0706
	STENCIL_INDEX8_OES                                         = 0x8D48
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	RGB5                                                       = 0x8050
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	YCRCBA_SGIX                                                = 0x8319
	E_TIMES_F_NV                                               = 0x8531
	SAMPLER_CUBE_ARB                                           = 0x8B60
	NUM_SAMPLE_COUNTS                                          = 0x9380
	READ_BUFFER_NV                                             = 0x0C02
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	TEXTURE_MATRIX                                             = 0x0BA8
	RGB8                                                       = 0x8051
	FUNC_SUBTRACT_OES                                          = 0x800A
	STREAM_DRAW                                                = 0x88E0
	CON_27_ATI                                                 = 0x895C
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	IMAGE_1D_EXT                                               = 0x904C
	IMAGE_BINDING_FORMAT                                       = 0x906E
	COLOR_ARRAY_SIZE                                           = 0x8081
	NOOP                                                       = 0x1505
	RGB2_EXT                                                   = 0x804E
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	VERTEX_SHADER_ARB                                          = 0x8B31
	INT_VEC2_ARB                                               = 0x8B53
	FLOAT_MAT3_ARB                                             = 0x8B5B
	SRGB                                                       = 0x8C40
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	BACK_LEFT                                                  = 0x0402
	CLIP_DISTANCE1                                             = 0x3001
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	STENCIL_INDEX4                                             = 0x8D47
	PARTIAL_SUCCESS_NV                                         = 0x902E
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	SGIX_resample                                              = 1
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	GL_2PASS_1_EXT                                             = 0x80A3
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	TESS_EVALUATION_SHADER                                     = 0x8E87
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	R16F_EXT                                                   = 0x822D
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	SOURCE1_ALPHA_ARB                                          = 0x8589
	INTERLACE_READ_OML                                         = 0x8981
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	DOUBLE_VEC4                                                = 0x8FFE
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	INVERT                                                     = 0x150A
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	STATIC_READ                                                = 0x88E5
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	CLAMP_TO_BORDER                                            = 0x812D
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	QUERY_RESULT_EXT                                           = 0x8866
	MATRIX18_ARB                                               = 0x88D2
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	VERTEX_ARRAY_POINTER                                       = 0x808E
	INDEX_ARRAY_STRIDE                                         = 0x8086
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	EYE_LINE_SGIS                                              = 0x81F6
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	FLOAT16_VEC3_NV                                            = 0x8FFA
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	ACTIVE_PROGRAM                                             = 0x8259
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	FACTOR_MAX_AMD                                             = 0x901D
	DRAW_BUFFER8                                               = 0x882D
	ARRAY_BUFFER                                               = 0x8892
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	INT_SAMPLER_3D                                             = 0x8DCB
	SGIX_subsample                                             = 1
	ALPHA                                                      = 0x1906
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	MAX_TEXTURE_UNITS                                          = 0x84E2
	OPERAND1_RGB_EXT                                           = 0x8591
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	SKIP_DECODE_EXT                                            = 0x8A4A
	SUBPIXEL_BITS                                              = 0x0D50
	INTENSITY8_EXT                                             = 0x804B
	RG32F                                                      = 0x8230
	SOURCE0_RGB_ARB                                            = 0x8580
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	SGI_color_matrix                                           = 1
	AND_INVERTED                                               = 0x1504
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	TEXTURE18_ARB                                              = 0x84D2
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	DRAW_BUFFER5_NV                                            = 0x882A
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	INT_VEC3                                                   = 0x8B54
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	BOOL_VEC3_ARB                                              = 0x8B58
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	DOUBLE_MAT2                                                = 0x8F46
	SGIS_pixel_texture                                         = 1
	VIEWPORT                                                   = 0x0BA2
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	MATRIX1_NV                                                 = 0x8631
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	LINEAR_ATTENUATION                                         = 0x1208
	CMYK_EXT                                                   = 0x800C
	SRGB_WRITE                                                 = 0x8298
	TRACK_MATRIX_NV                                            = 0x8648
	VERTEX_STREAM7_ATI                                         = 0x8773
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	GREEN_SCALE                                                = 0x0D18
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	SGIX_interlace                                             = 1
	PROGRAM_STRING_NV                                          = 0x8628
	DRAW_BUFFER7                                               = 0x882C
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	MINMAX_FORMAT_EXT                                          = 0x802F
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	ALPHA_INTEGER_EXT                                          = 0x8D97
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	LIGHT5                                                     = 0x4005
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	DRAW_BUFFER0_ATI                                           = 0x8825
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	SEPARATE_ATTRIBS                                           = 0x8C8D
	GREEN_INTEGER                                              = 0x8D95
	TRIANGLE_FAN                                               = 0x0006
	EMISSION                                                   = 0x1600
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	ADD_ATI                                                    = 0x8963
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	SHININESS                                                  = 0x1601
	CONVOLUTION_2D_EXT                                         = 0x8011
	TEXTURE_DEPTH_EXT                                          = 0x8071
	VIEW_CLASS_128_BITS                                        = 0x82C4
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	COMBINER0_NV                                               = 0x8550
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	TRIANGLES_ADJACENCY                                        = 0x000C
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	DS_SCALE_NV                                                = 0x8710
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	STATIC_DRAW                                                = 0x88E4
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	SCALE_BY_FOUR_NV                                           = 0x853F
	MODELVIEW26_ARB                                            = 0x873A
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	LIGHT4                                                     = 0x4004
	BGR                                                        = 0x80E0
	R8_EXT                                                     = 0x8229
	OPERAND1_ALPHA_ARB                                         = 0x8599
	MIRROR_CLAMP_EXT                                           = 0x8742
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	MAJOR_VERSION                                              = 0x821B
	PROGRAM_PARAMETER_NV                                       = 0x8644
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	INTENSITY_SNORM                                            = 0x9013
	LINE                                                       = 0x1B01
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	COLOR_SUM                                                  = 0x8458
	TEXTURE2_ARB                                               = 0x84C2
	PACK_INVERT_MESA                                           = 0x8758
	TEXTURE_SWIZZLE_R                                          = 0x8E42
)

type Context struct {
	access                    sync.Mutex
	context                   *C.gl10Context
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
	glc.context = C.gl10NewContext()

	glc.Accum = func(op uint32, value float32) {
		defer glc.trace("Accum")
		C.gl10Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		defer glc.trace("AlphaFunc")
		C.gl10AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		defer glc.trace("Begin")
		glc.inBeginEnd = true
		C.gl10Begin(glc.context, C.GLenum(mode))
		return
	}

	glc.End = func() {
		defer glc.trace("End")
		C.gl10End(glc.context)
		glc.inBeginEnd = false
		return
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap *uint8) {
		defer glc.trace("Bitmap")
		C.gl10Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		defer glc.trace("BlendFunc")
		C.gl10BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		defer glc.trace("CallList")
		C.gl10CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		defer glc.trace("CallLists")
		C.gl10CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		defer glc.trace("Clear")
		C.gl10Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		defer glc.trace("ClearAccum")
		C.gl10ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		defer glc.trace("ClearColor")
		C.gl10ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		defer glc.trace("ClearDepth")
		C.gl10ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		defer glc.trace("ClearIndex")
		C.gl10ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		defer glc.trace("ClearStencil")
		C.gl10ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation *float64) {
		defer glc.trace("ClipPlane")
		C.gl10ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.Color3b = func(red, green, blue int8) {
		defer glc.trace("Color3b")
		C.gl10Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		defer glc.trace("Color3d")
		C.gl10Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		defer glc.trace("Color3f")
		C.gl10Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		defer glc.trace("Color3i")
		C.gl10Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		defer glc.trace("Color3s")
		C.gl10Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		defer glc.trace("Color3ub")
		C.gl10Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		defer glc.trace("Color3ui")
		C.gl10Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		defer glc.trace("Color3us")
		C.gl10Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		defer glc.trace("Color4b")
		C.gl10Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		defer glc.trace("Color4d")
		C.gl10Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		defer glc.trace("Color4f")
		C.gl10Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		defer glc.trace("Color4i")
		C.gl10Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		defer glc.trace("Color4s")
		C.gl10Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		defer glc.trace("Color4ub")
		C.gl10Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		defer glc.trace("Color4ui")
		C.gl10Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		defer glc.trace("Color4us")
		C.gl10Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v *int8) {
		defer glc.trace("Color3bv")
		C.gl10Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color3dv = func(v *float64) {
		defer glc.trace("Color3dv")
		C.gl10Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color3fv = func(v *float32) {
		defer glc.trace("Color3fv")
		C.gl10Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color3iv = func(v *int32) {
		defer glc.trace("Color3iv")
		C.gl10Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color3sv = func(v *int16) {
		defer glc.trace("Color3sv")
		C.gl10Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color3ubv = func(v *uint8) {
		defer glc.trace("Color3ubv")
		C.gl10Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color3uiv = func(v *uint32) {
		defer glc.trace("Color3uiv")
		C.gl10Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color3usv = func(v *uint16) {
		defer glc.trace("Color3usv")
		C.gl10Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.Color4bv = func(v *int8) {
		defer glc.trace("Color4bv")
		C.gl10Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Color4dv = func(v *float64) {
		defer glc.trace("Color4dv")
		C.gl10Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Color4fv = func(v *float32) {
		defer glc.trace("Color4fv")
		C.gl10Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Color4iv = func(v *int32) {
		defer glc.trace("Color4iv")
		C.gl10Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Color4sv = func(v *int16) {
		defer glc.trace("Color4sv")
		C.gl10Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Color4ubv = func(v *uint8) {
		defer glc.trace("Color4ubv")
		C.gl10Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.Color4uiv = func(v *uint32) {
		defer glc.trace("Color4uiv")
		C.gl10Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.Color4usv = func(v *uint16) {
		defer glc.trace("Color4usv")
		C.gl10Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		defer glc.trace("ColorMask")
		C.gl10ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		defer glc.trace("ColorMaterial")
		C.gl10ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		defer glc.trace("CopyPixels")
		C.gl10CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		defer glc.trace("CullFace")
		C.gl10CullFace(glc.context, C.GLenum(mode))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		defer glc.trace("DeleteLists")
		C.gl10DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		defer glc.trace("DepthFunc")
		C.gl10DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthMask = func(flag bool) {
		defer glc.trace("DepthMask")
		C.gl10DepthMask(glc.context, boolToGL(flag))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		defer glc.trace("DepthRange")
		C.gl10DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		defer glc.trace("Enable")
		C.gl10Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		defer glc.trace("Disable")
		C.gl10Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		defer glc.trace("DrawBuffer")
		C.gl10DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("DrawPixels")
		C.gl10DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.EdgeFlag = func(flag bool) {
		defer glc.trace("EdgeFlag")
		C.gl10EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag *bool) {
		defer glc.trace("EdgeFlagv")
		C.gl10EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(flag)))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		defer glc.trace("EdgeFlagPointer")
		C.gl10EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EvalCoord1d = func(u float64) {
		defer glc.trace("EvalCoord1d")
		C.gl10EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		defer glc.trace("EvalCoord1f")
		C.gl10EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		defer glc.trace("EvalCoord2d")
		C.gl10EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		defer glc.trace("EvalCoord2f")
		C.gl10EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u *float64) {
		defer glc.trace("EvalCoord1dv")
		C.gl10EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord1fv = func(u *float32) {
		defer glc.trace("EvalCoord1fv")
		C.gl10EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2dv = func(u *float64) {
		defer glc.trace("EvalCoord2dv")
		C.gl10EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(u)))
	}

	glc.EvalCoord2fv = func(u *float32) {
		defer glc.trace("EvalCoord2fv")
		C.gl10EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(u)))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		defer glc.trace("EvalMesh1")
		C.gl10EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		defer glc.trace("EvalMesh2")
		C.gl10EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		defer glc.trace("EvalPoint1")
		C.gl10EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		defer glc.trace("EvalPoint2")
		C.gl10EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer *float32) {
		defer glc.trace("FeedbackBuffer")
		C.gl10FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(buffer)))
	}

	glc.Finish = func() {
		defer glc.trace("Finish")
		C.gl10Finish(glc.context)
	}

	glc.Flush = func() {
		defer glc.trace("Flush")
		C.gl10Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		defer glc.trace("Fogf")
		C.gl10Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		defer glc.trace("Fogi")
		C.gl10Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params *float32) {
		defer glc.trace("Fogfv")
		C.gl10Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Fogiv = func(pname uint32, params *int32) {
		defer glc.trace("Fogiv")
		C.gl10Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.FrontFace = func(mode uint32) {
		defer glc.trace("FrontFace")
		C.gl10FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		defer glc.trace("Frustum")
		C.gl10Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		defer glc.trace("GenLists")
		return uint32(C.gl10GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params *bool) {
		defer glc.trace("GetBooleanv")
		C.gl10GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(params)))
	}

	glc.GetDoublev = func(pname uint32, params *float64) {
		defer glc.trace("GetDoublev")
		C.gl10GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetFloatv = func(pname uint32, params *float32) {
		defer glc.trace("GetFloatv")
		C.gl10GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetIntegerv = func(pname uint32, params *int32) {
		defer glc.trace("GetIntegerv")
		C.gl10GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetClipPlane = func(plane uint32, equation *float64) {
		defer glc.trace("GetClipPlane")
		C.gl10GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(equation)))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl10GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params *float32) {
		defer glc.trace("GetLightfv")
		C.gl10GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetLightiv = func(light, pname uint32, params *int32) {
		defer glc.trace("GetLightiv")
		C.gl10GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetMapdv = func(target, query uint32, v *float64) {
		defer glc.trace("GetMapdv")
		C.gl10GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.GetMapfv = func(target, query uint32, v *float32) {
		defer glc.trace("GetMapfv")
		C.gl10GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.GetMapiv = func(target, query uint32, v *int32) {
		defer glc.trace("GetMapiv")
		C.gl10GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.GetMaterialfv = func(face, pname uint32, params *float32) {
		defer glc.trace("GetMaterialfv")
		C.gl10GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetMaterialiv = func(face, pname uint32, params *int32) {
		defer glc.trace("GetMaterialiv")
		C.gl10GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetPixelMapfv = func(Map uint32, values *float32) {
		defer glc.trace("GetPixelMapfv")
		C.gl10GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapuiv = func(Map uint32, values *uint32) {
		defer glc.trace("GetPixelMapuiv")
		C.gl10GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.GetPixelMapusv = func(Map uint32, values *uint16) {
		defer glc.trace("GetPixelMapusv")
		C.gl10GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.GetPolygonStipple = func(pattern *uint8) {
		defer glc.trace("GetPolygonStipple")
		C.gl10GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(pattern)))
	}

	glc.GetString = func(name uint32) string {
		defer glc.trace("GetString")
		cstr := C.gl10GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetTexEnvfv")
		C.gl10GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexEnviv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetTexEnviv")
		C.gl10GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexGendv = func(coord, pname uint32, params *float64) {
		defer glc.trace("GetTexGendv")
		C.gl10GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params *float32) {
		defer glc.trace("GetTexGenfv")
		C.gl10GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params *int32) {
		defer glc.trace("GetTexGeniv")
		C.gl10GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("GetTexImage")
		C.gl10GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params *float32) {
		defer glc.trace("GetTexLevelParameterfv")
		C.gl10GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params *int32) {
		defer glc.trace("GetTexLevelParameteriv")
		C.gl10GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetTexParameterfv")
		C.gl10GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetTexParameteriv")
		C.gl10GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Hint = func(target, mode uint32) {
		defer glc.trace("Hint")
		C.gl10Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		defer glc.trace("Indexd")
		C.gl10Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		defer glc.trace("Indexf")
		C.gl10Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		defer glc.trace("Indexi")
		C.gl10Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		defer glc.trace("Indexs")
		C.gl10Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexdv = func(c *float64) {
		defer glc.trace("Indexdv")
		C.gl10Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(c)))
	}

	glc.Indexfv = func(c *float32) {
		defer glc.trace("Indexfv")
		C.gl10Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(c)))
	}

	glc.Indexiv = func(c *int32) {
		defer glc.trace("Indexiv")
		C.gl10Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(c)))
	}

	glc.Indexsv = func(c *int16) {
		defer glc.trace("Indexsv")
		C.gl10Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(c)))
	}

	glc.IndexMask = func(mask uint32) {
		defer glc.trace("IndexMask")
		C.gl10IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("IndexPointer")
		C.gl10IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		defer glc.trace("InitNames")
		C.gl10InitNames(glc.context)
	}

	glc.IsEnabled = func(cap uint32) {
		defer glc.trace("IsEnabled")
		C.gl10IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		defer glc.trace("IsList")
		return C.gl10IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		defer glc.trace("Lightf")
		C.gl10Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		defer glc.trace("Lighti")
		C.gl10Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params *float32) {
		defer glc.trace("Lightfv")
		C.gl10Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Lightiv = func(light, pname uint32, params *int32) {
		defer glc.trace("Lightiv")
		C.gl10Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		defer glc.trace("LightModelf")
		C.gl10LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		defer glc.trace("LightModeli")
		C.gl10LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params *float32) {
		defer glc.trace("LightModelfv")
		C.gl10LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.LightModeliv = func(pname uint32, params *int32) {
		defer glc.trace("LightModeliv")
		C.gl10LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		defer glc.trace("LineStipple")
		C.gl10LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		defer glc.trace("LineWidth")
		C.gl10LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		defer glc.trace("ListBase")
		C.gl10ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		defer glc.trace("LoadIdentity")
		C.gl10LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m *float64) {
		defer glc.trace("LoadMatrixd")
		C.gl10LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadMatrixf = func(m *float32) {
		defer glc.trace("LoadMatrixf")
		C.gl10LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.LoadName = func(name uint32) {
		defer glc.trace("LoadName")
		C.gl10LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		defer glc.trace("LogicOp")
		C.gl10LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points *float64) {
		defer glc.trace("Map1d")
		C.gl10Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points *float32) {
		defer glc.trace("Map1f")
		C.gl10Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points *float64) {
		defer glc.trace("Map2d")
		C.gl10Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(points)))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points *float32) {
		defer glc.trace("Map2f")
		C.gl10Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(points)))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		defer glc.trace("MapGrid1d")
		C.gl10MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		defer glc.trace("MapGrid1f")
		C.gl10MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		defer glc.trace("MapGrid2d")
		C.gl10MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		defer glc.trace("MapGrid2f")
		C.gl10MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		defer glc.trace("Materialf")
		C.gl10Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		defer glc.trace("Materiali")
		C.gl10Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params *float32) {
		defer glc.trace("Materialfv")
		C.gl10Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.Materialiv = func(face, pname uint32, params *int32) {
		defer glc.trace("Materialiv")
		C.gl10Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.MatrixMode = func(mode uint32) {
		defer glc.trace("MatrixMode")
		C.gl10MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m *float64) {
		defer glc.trace("MultMatrixd")
		C.gl10MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultMatrixf = func(m *float32) {
		defer glc.trace("MultMatrixf")
		C.gl10MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.NewList = func(list uint32, mode uint32) {
		defer glc.trace("NewList")
		C.gl10NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		defer glc.trace("EndList")
		C.gl10EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		defer glc.trace("Normal3b")
		C.gl10Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		defer glc.trace("Normal3d")
		C.gl10Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		defer glc.trace("Normal3f")
		C.gl10Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		defer glc.trace("Normal3i")
		C.gl10Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		defer glc.trace("Normal3s")
		C.gl10Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v *int8) {
		defer glc.trace("Normal3bv")
		C.gl10Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.Normal3dv = func(v *float64) {
		defer glc.trace("Normal3dv")
		C.gl10Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.Normal3fv = func(v *float32) {
		defer glc.trace("Normal3fv")
		C.gl10Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.Normal3iv = func(v *int32) {
		defer glc.trace("Normal3iv")
		C.gl10Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.Normal3sv = func(v *int16) {
		defer glc.trace("Normal3sv")
		C.gl10Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		defer glc.trace("Ortho")
		C.gl10Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		defer glc.trace("PassThrough")
		C.gl10PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values *float32) {
		defer glc.trace("PixelMapfv")
		C.gl10PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(values)))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values *uint32) {
		defer glc.trace("PixelMapuiv")
		C.gl10PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(values)))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values *uint16) {
		defer glc.trace("PixelMapusv")
		C.gl10PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(values)))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		defer glc.trace("PixelStoref")
		C.gl10PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		defer glc.trace("PixelStorei")
		C.gl10PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		defer glc.trace("PixelTransferf")
		C.gl10PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		defer glc.trace("PixelTransferi")
		C.gl10PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		defer glc.trace("PixelZoom")
		C.gl10PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		defer glc.trace("PointSize")
		C.gl10PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		defer glc.trace("PolygonMode")
		C.gl10PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask *uint8) {
		defer glc.trace("PolygonStipple")
		C.gl10PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(mask)))
	}

	glc.PushAttrib = func(mask uint32) {
		defer glc.trace("PushAttrib")
		C.gl10PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		defer glc.trace("PopAttrib")
		C.gl10PopAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		defer glc.trace("PushMatrix")
		C.gl10PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		defer glc.trace("PopMatrix")
		C.gl10PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		defer glc.trace("PushName")
		C.gl10PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		defer glc.trace("PopName")
		C.gl10PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		defer glc.trace("RasterPos2d")
		C.gl10RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		defer glc.trace("RasterPos2f")
		C.gl10RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		defer glc.trace("RasterPos2i")
		C.gl10RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		defer glc.trace("RasterPos2s")
		C.gl10RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		defer glc.trace("RasterPos3d")
		C.gl10RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		defer glc.trace("RasterPos3f")
		C.gl10RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		defer glc.trace("RasterPos3i")
		C.gl10RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		defer glc.trace("RasterPos3s")
		C.gl10RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		defer glc.trace("RasterPos4d")
		C.gl10RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		defer glc.trace("RasterPos4f")
		C.gl10RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		defer glc.trace("RasterPos4i")
		C.gl10RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		defer glc.trace("RasterPos4s")
		C.gl10RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v *float64) {
		defer glc.trace("RasterPos2dv")
		C.gl10RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos2fv = func(v *float32) {
		defer glc.trace("RasterPos2fv")
		C.gl10RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos2iv = func(v *int32) {
		defer glc.trace("RasterPos2iv")
		C.gl10RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos2sv = func(v *int16) {
		defer glc.trace("RasterPos2sv")
		C.gl10RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos3dv = func(v *float64) {
		defer glc.trace("RasterPos3dv")
		C.gl10RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos3fv = func(v *float32) {
		defer glc.trace("RasterPos3fv")
		C.gl10RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos3iv = func(v *int32) {
		defer glc.trace("RasterPos3iv")
		C.gl10RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos3sv = func(v *int16) {
		defer glc.trace("RasterPos3sv")
		C.gl10RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.RasterPos4dv = func(v *float64) {
		defer glc.trace("RasterPos4dv")
		C.gl10RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.RasterPos4fv = func(v *float32) {
		defer glc.trace("RasterPos4fv")
		C.gl10RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.RasterPos4iv = func(v *int32) {
		defer glc.trace("RasterPos4iv")
		C.gl10RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.RasterPos4sv = func(v *int16) {
		defer glc.trace("RasterPos4sv")
		C.gl10RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.ReadBuffer = func(mode uint32) {
		defer glc.trace("ReadBuffer")
		C.gl10ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("ReadPixels")
		C.gl10ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		defer glc.trace("Rectd")
		C.gl10Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		defer glc.trace("Rectf")
		C.gl10Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		defer glc.trace("Recti")
		C.gl10Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		defer glc.trace("Rects")
		C.gl10Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 *float64) {
		defer glc.trace("Rectdv")
		C.gl10Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(v1)), (*C.GLdouble)(unsafe.Pointer(v2)))
	}

	glc.Rectfv = func(v1, v2 *float32) {
		defer glc.trace("Rectfv")
		C.gl10Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(v1)), (*C.GLfloat)(unsafe.Pointer(v2)))
	}

	glc.Rectiv = func(v1, v2 *int32) {
		defer glc.trace("Rectiv")
		C.gl10Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(v1)), (*C.GLint)(unsafe.Pointer(v2)))
	}

	glc.Rectsv = func(v1, v2 *int16) {
		defer glc.trace("Rectsv")
		C.gl10Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(v1)), (*C.GLshort)(unsafe.Pointer(v2)))
	}

	glc.RenderMode = func(mode uint32) int32 {
		defer glc.trace("RenderMode")
		return int32(C.gl10RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		defer glc.trace("Rotated")
		C.gl10Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		defer glc.trace("Rotatef")
		C.gl10Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		defer glc.trace("Scaled")
		C.gl10Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		defer glc.trace("Scalef")
		C.gl10Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		defer glc.trace("Scissor")
		C.gl10Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer *uint32) {
		defer glc.trace("SelectBuffer")
		C.gl10SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(buffer)))
	}

	glc.ShadeModel = func(mode uint32) {
		defer glc.trace("ShadeModel")
		C.gl10ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		defer glc.trace("StencilFunc")
		C.gl10StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		defer glc.trace("StencilMask")
		C.gl10StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		defer glc.trace("StencilOp")
		C.gl10StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		defer glc.trace("TexCoord1d")
		C.gl10TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		defer glc.trace("TexCoord1f")
		C.gl10TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		defer glc.trace("TexCoord1i")
		C.gl10TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		defer glc.trace("TexCoord1s")
		C.gl10TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		defer glc.trace("TexCoord2d")
		C.gl10TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		defer glc.trace("TexCoord2f")
		C.gl10TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		defer glc.trace("TexCoord2i")
		C.gl10TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		defer glc.trace("TexCoord2s")
		C.gl10TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		defer glc.trace("TexCoord3d")
		C.gl10TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		defer glc.trace("TexCoord3f")
		C.gl10TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		defer glc.trace("TexCoord3i")
		C.gl10TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		defer glc.trace("TexCoord3s")
		C.gl10TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		defer glc.trace("TexCoord4d")
		C.gl10TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		defer glc.trace("TexCoord4f")
		C.gl10TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		defer glc.trace("TexCoord4i")
		C.gl10TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		defer glc.trace("TexCoord4s")
		C.gl10TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v *float64) {
		defer glc.trace("TexCoord1dv")
		C.gl10TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord1fv = func(v *float32) {
		defer glc.trace("TexCoord1fv")
		C.gl10TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord1iv = func(v *int32) {
		defer glc.trace("TexCoord1iv")
		C.gl10TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord1sv = func(v *int16) {
		defer glc.trace("TexCoord1sv")
		C.gl10TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord2dv = func(v *float64) {
		defer glc.trace("TexCoord2dv")
		C.gl10TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord2fv = func(v *float32) {
		defer glc.trace("TexCoord2fv")
		C.gl10TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord2iv = func(v *int32) {
		defer glc.trace("TexCoord2iv")
		C.gl10TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord2sv = func(v *int16) {
		defer glc.trace("TexCoord2sv")
		C.gl10TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord3dv = func(v *float64) {
		defer glc.trace("TexCoord3dv")
		C.gl10TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord3fv = func(v *float32) {
		defer glc.trace("TexCoord3fv")
		C.gl10TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord3iv = func(v *int32) {
		defer glc.trace("TexCoord3iv")
		C.gl10TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord3sv = func(v *int16) {
		defer glc.trace("TexCoord3sv")
		C.gl10TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexCoord4dv = func(v *float64) {
		defer glc.trace("TexCoord4dv")
		C.gl10TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.TexCoord4fv = func(v *float32) {
		defer glc.trace("TexCoord4fv")
		C.gl10TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.TexCoord4iv = func(v *int32) {
		defer glc.trace("TexCoord4iv")
		C.gl10TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.TexCoord4sv = func(v *int16) {
		defer glc.trace("TexCoord4sv")
		C.gl10TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		defer glc.trace("TexEnvf")
		C.gl10TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		defer glc.trace("TexEnvi")
		C.gl10TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params *float32) {
		defer glc.trace("TexEnvfv")
		C.gl10TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexEnviv = func(target, pname uint32, params *int32) {
		defer glc.trace("TexEnviv")
		C.gl10TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		defer glc.trace("TexGend")
		C.gl10TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		defer glc.trace("TexGenf")
		C.gl10TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		defer glc.trace("TexGeni")
		C.gl10TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params *float64) {
		defer glc.trace("TexGendv")
		C.gl10TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.TexGenfv = func(coord, pname uint32, params *float32) {
		defer glc.trace("TexGenfv")
		C.gl10TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexGeniv = func(coord, pname uint32, params *int32) {
		defer glc.trace("TexGeniv")
		C.gl10TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexImage1D")
		C.gl10TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexImage2D")
		C.gl10TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		defer glc.trace("TexParameterf")
		C.gl10TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		defer glc.trace("TexParameteri")
		C.gl10TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("TexParameterfv")
		C.gl10TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.TexParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("TexParameteriv")
		C.gl10TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.Translated = func(x, y, z float64) {
		defer glc.trace("Translated")
		C.gl10Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		defer glc.trace("Translatef")
		C.gl10Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		defer glc.trace("Vertex2s")
		C.gl10Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		defer glc.trace("Vertex2i")
		C.gl10Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		defer glc.trace("Vertex2f")
		C.gl10Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		defer glc.trace("Vertex2d")
		C.gl10Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		defer glc.trace("Vertex3s")
		C.gl10Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		defer glc.trace("Vertex3i")
		C.gl10Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		defer glc.trace("Vertex3f")
		C.gl10Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		defer glc.trace("Vertex3d")
		C.gl10Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		defer glc.trace("Vertex4s")
		C.gl10Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		defer glc.trace("Vertex4i")
		C.gl10Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		defer glc.trace("Vertex4f")
		C.gl10Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		defer glc.trace("Vertex4d")
		C.gl10Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		defer glc.trace("Viewport")
		C.gl10Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetConvolutionParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetConvolutionParameterfv")
		C.gl10GetConvolutionParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetConvolutionParameteriv")
		C.gl10GetConvolutionParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ColorTable")
		C.gl10ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("ColorTableParameterfv")
		C.gl10ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("ColorTableParameteriv")
		C.gl10ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ColorSubTable")
		C.gl10ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ConvolutionFilter1D")
		C.gl10ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		defer glc.trace("ConvolutionFilter2D")
		C.gl10ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		defer glc.trace("ConvolutionParameterf")
		C.gl10ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		defer glc.trace("ConvolutionParameteri")
		C.gl10ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		defer glc.trace("CopyColorTable")
		C.gl10CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		defer glc.trace("CopyColorSubTable")
		C.gl10CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		defer glc.trace("CopyConvolutionFilter1D")
		C.gl10CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		defer glc.trace("CopyConvolutionFilter2D")
		C.gl10CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.GetColorTable = func(target, format, Type uint32, table unsafe.Pointer) {
		defer glc.trace("GetColorTable")
		C.gl10GetColorTable(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), table)
	}

	glc.GetColorTableParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetColorTableParameterfv")
		C.gl10GetColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetColorTableParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetColorTableParameteriv")
		C.gl10GetColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetConvolutionFilter = func(target, format, Type uint32, image unsafe.Pointer) {
		defer glc.trace("GetConvolutionFilter")
		C.gl10GetConvolutionFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), image)
	}

	glc.GetHistogram = func(target uint32, reset bool, format, Type uint32, values unsafe.Pointer) {
		defer glc.trace("GetHistogram")
		C.gl10GetHistogram(glc.context, C.GLenum(target), boolToGL(reset), C.GLenum(format), C.GLenum(Type), values)
	}

	glc.GetHistogramParameterfv = func(target, pname uint32, params *float32) {
		defer glc.trace("GetHistogramParameterfv")
		C.gl10GetHistogramParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetHistogramParameteriv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetHistogramParameteriv")
		C.gl10GetHistogramParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetSeparableFilter = func(target, format, Type uint32, row, column, span unsafe.Pointer) {
		defer glc.trace("GetSeparableFilter")
		C.gl10GetSeparableFilter(glc.context, C.GLenum(target), C.GLenum(format), C.GLenum(Type), row, column, span)
	}

	glc.Histogram = func(target uint32, width int32, internalformat uint32, sink bool) {
		defer glc.trace("Histogram")
		C.gl10Histogram(glc.context, C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.Minmax = func(target, internalformat uint32, sink bool) {
		defer glc.trace("Minmax")
		C.gl10Minmax(glc.context, C.GLenum(target), C.GLenum(internalformat), boolToGL(sink))
	}

	glc.MultiTexCoord1s = func(target uint32, s int16) {
		defer glc.trace("MultiTexCoord1s")
		C.gl10MultiTexCoord1s(glc.context, C.GLenum(target), C.GLshort(s))
	}

	glc.MultiTexCoord1i = func(target uint32, s int32) {
		defer glc.trace("MultiTexCoord1i")
		C.gl10MultiTexCoord1i(glc.context, C.GLenum(target), C.GLint(s))
	}

	glc.MultiTexCoord1f = func(target uint32, s float32) {
		defer glc.trace("MultiTexCoord1f")
		C.gl10MultiTexCoord1f(glc.context, C.GLenum(target), C.GLfloat(s))
	}

	glc.MultiTexCoord1d = func(target uint32, s float64) {
		defer glc.trace("MultiTexCoord1d")
		C.gl10MultiTexCoord1d(glc.context, C.GLenum(target), C.GLdouble(s))
	}

	glc.MultiTexCoord2s = func(target uint32, s, t int16) {
		defer glc.trace("MultiTexCoord2s")
		C.gl10MultiTexCoord2s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t))
	}

	glc.MultiTexCoord2i = func(target uint32, s, t int32) {
		defer glc.trace("MultiTexCoord2i")
		C.gl10MultiTexCoord2i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t))
	}

	glc.MultiTexCoord2f = func(target uint32, s, t float32) {
		defer glc.trace("MultiTexCoord2f")
		C.gl10MultiTexCoord2f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
	}

	glc.MultiTexCoord2d = func(target uint32, s, t float64) {
		defer glc.trace("MultiTexCoord2d")
		C.gl10MultiTexCoord2d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
	}

	glc.MultiTexCoord3s = func(target uint32, s, t, r int16) {
		defer glc.trace("MultiTexCoord3s")
		C.gl10MultiTexCoord3s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.MultiTexCoord3i = func(target uint32, s, t, r int32) {
		defer glc.trace("MultiTexCoord3i")
		C.gl10MultiTexCoord3i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.MultiTexCoord3f = func(target uint32, s, t, r float32) {
		defer glc.trace("MultiTexCoord3f")
		C.gl10MultiTexCoord3f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.MultiTexCoord3d = func(target uint32, s, t, r float64) {
		defer glc.trace("MultiTexCoord3d")
		C.gl10MultiTexCoord3d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.MultiTexCoord4s = func(target uint32, s, t, r, q int16) {
		defer glc.trace("MultiTexCoord4s")
		C.gl10MultiTexCoord4s(glc.context, C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.MultiTexCoord4i = func(target uint32, s, t, r, q int32) {
		defer glc.trace("MultiTexCoord4i")
		C.gl10MultiTexCoord4i(glc.context, C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.MultiTexCoord4f = func(target uint32, s, t, r, q float32) {
		defer glc.trace("MultiTexCoord4f")
		C.gl10MultiTexCoord4f(glc.context, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.MultiTexCoord4d = func(target uint32, s, t, r, q float64) {
		defer glc.trace("MultiTexCoord4d")
		C.gl10MultiTexCoord4d(glc.context, C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.MultiTexCoord1sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord1sv")
		C.gl10MultiTexCoord1sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord1iv")
		C.gl10MultiTexCoord1iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord1fv")
		C.gl10MultiTexCoord1fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord1dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord1dv")
		C.gl10MultiTexCoord1dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord2sv")
		C.gl10MultiTexCoord2sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord2iv")
		C.gl10MultiTexCoord2iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord2fv")
		C.gl10MultiTexCoord2fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord2dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord2dv")
		C.gl10MultiTexCoord2dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord3sv")
		C.gl10MultiTexCoord3sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord3iv")
		C.gl10MultiTexCoord3iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord3fv")
		C.gl10MultiTexCoord3fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord3dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord3dv")
		C.gl10MultiTexCoord3dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4sv = func(target uint32, v *int16) {
		defer glc.trace("MultiTexCoord4sv")
		C.gl10MultiTexCoord4sv(glc.context, C.GLenum(target), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4iv = func(target uint32, v *int32) {
		defer glc.trace("MultiTexCoord4iv")
		C.gl10MultiTexCoord4iv(glc.context, C.GLenum(target), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4fv = func(target uint32, v *float32) {
		defer glc.trace("MultiTexCoord4fv")
		C.gl10MultiTexCoord4fv(glc.context, C.GLenum(target), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.MultiTexCoord4dv = func(target uint32, v *float64) {
		defer glc.trace("MultiTexCoord4dv")
		C.gl10MultiTexCoord4dv(glc.context, C.GLenum(target), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.ResetHistogram = func(target uint32) {
		defer glc.trace("ResetHistogram")
		C.gl10ResetHistogram(glc.context, C.GLenum(target))
	}

	glc.ResetMinmax = func(target uint32) {
		defer glc.trace("ResetMinmax")
		C.gl10ResetMinmax(glc.context, C.GLenum(target))
	}

	glc.SeparableFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, row, column unsafe.Pointer) {
		defer glc.trace("SeparableFilter2D")
		C.gl10SeparableFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), row, column)
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		defer glc.trace("AreTexturesResident")
		var cRes *C.GLboolean
		status = C.gl10AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		defer glc.trace("ArrayElement")
		C.gl10ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		defer glc.trace("DrawArrays")
		C.gl10DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		defer glc.trace("DrawElements")
		C.gl10DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		defer glc.trace("GetPointerv")
		C.gl10GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		defer glc.trace("PolygonOffset")
		C.gl10PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		defer glc.trace("CopyTexImage1D")
		C.gl10CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		defer glc.trace("CopyTexImage2D")
		C.gl10CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		defer glc.trace("CopyTexSubImage1D")
		C.gl10CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		defer glc.trace("CopyTexSubImage2D")
		C.gl10CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		defer glc.trace("BindTexture")
		C.gl10BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures *uint32) {
		defer glc.trace("DeleteTextures")
		C.gl10DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.GenTextures = func(n int32, textures *uint32) {
		defer glc.trace("GenTextures")
		C.gl10GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}

	glc.IsTexture = func(texture uint32) bool {
		defer glc.trace("IsTexture")
		return C.gl10IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("ColorPointer")
		C.gl10ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		defer glc.trace("EnableClientState")
		C.gl10EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		defer glc.trace("DisableClientState")
		C.gl10DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.Indexub = func(c uint8) {
		defer glc.trace("Indexub")
		C.gl10Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexubv = func(c *uint8) {
		defer glc.trace("Indexubv")
		C.gl10Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(c)))
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("InterleavedArrays")
		C.gl10InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("NormalPointer")
		C.gl10NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.PushClientAttrib = func(mask uint32) {
		defer glc.trace("PushClientAttrib")
		C.gl10PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PrioritizeTextures = func(n int32, textures *uint32, priorities *float32) {
		defer glc.trace("PrioritizeTextures")
		C.gl10PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLclampf)(unsafe.Pointer(priorities)))
	}

	glc.PopClientAttrib = func() {
		defer glc.trace("PopClientAttrib")
		C.gl10PopClientAttrib(glc.context)
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("TexCoordPointer")
		C.gl10TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexSubImage1D")
		C.gl10TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexSubImage2D")
		C.gl10TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("VertexPointer")
		C.gl10VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.BlendColor = func(red, green, blue, alpha float32) {
		defer glc.trace("BlendColor")
		C.gl10BlendColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		defer glc.trace("BlendEquation")
		C.gl10BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		defer glc.trace("CopyTexSubImage3D")
		C.gl10CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DrawRangeElements = func(mode uint32, start, end uint32, count int32, Type uint32, indices unsafe.Pointer) {
		defer glc.trace("DrawRangeElements")
		C.gl10DrawRangeElements(glc.context, C.GLenum(mode), C.GLuint(start), C.GLuint(end), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.TexImage3D = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexImage3D")
		C.gl10TexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		defer glc.trace("TexSubImage3D")
		C.gl10TexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.ActiveTexture = func(texture uint32) {
		defer glc.trace("ActiveTexture")
		C.gl10ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		defer glc.trace("ClientActiveTexture")
		C.gl10ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexImage1D")
		C.gl10CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexImage2D")
		C.gl10CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexImage3D")
		C.gl10CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexSubImage1D")
		C.gl10CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexSubImage2D")
		C.gl10CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		defer glc.trace("CompressedTexSubImage3D")
		C.gl10CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.GetCompressedTexImage = func(target uint32, lod int32, img unsafe.Pointer) {
		defer glc.trace("GetCompressedTexImage")
		C.gl10GetCompressedTexImage(glc.context, C.GLenum(target), C.GLint(lod), img)
	}

	glc.LoadTransposeMatrixd = func(m *float64) {
		defer glc.trace("LoadTransposeMatrixd")
		C.gl10LoadTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.LoadTransposeMatrixf = func(m *float64) {
		defer glc.trace("LoadTransposeMatrixf")
		C.gl10LoadTransposeMatrixf(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixd = func(m *float64) {
		defer glc.trace("MultTransposeMatrixd")
		C.gl10MultTransposeMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(m)))
	}

	glc.MultTransposeMatrixf = func(m *float32) {
		defer glc.trace("MultTransposeMatrixf")
		C.gl10MultTransposeMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(m)))
	}

	glc.SampleCoverage = func(value float32, invert bool) {
		defer glc.trace("SampleCoverage")
		C.gl10SampleCoverage(glc.context, C.GLclampf(value), boolToGL(invert))
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		defer glc.trace("BlendFuncSeparate")
		C.gl10BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.FogCoordPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("FogCoordPointer")
		C.gl10FogCoordPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.FogCoordd = func(coord float64) {
		defer glc.trace("FogCoordd")
		C.gl10FogCoordd(glc.context, C.GLdouble(coord))
	}

	glc.FogCoordf = func(coord float32) {
		defer glc.trace("FogCoordf")
		C.gl10FogCoordf(glc.context, C.GLfloat(coord))
	}

	glc.FogCoorddv = func(coord *float64) {
		defer glc.trace("FogCoorddv")
		C.gl10FogCoorddv(glc.context, (*C.GLdouble)(unsafe.Pointer(coord)))
	}

	glc.FogCoordfv = func(coord *float32) {
		defer glc.trace("FogCoordfv")
		C.gl10FogCoordfv(glc.context, (*C.GLfloat)(unsafe.Pointer(coord)))
	}

	glc.MultiDrawArrays = func(mode uint32, first *int32, count *int32, primcount int32) {
		defer glc.trace("MultiDrawArrays")
		C.gl10MultiDrawArrays(glc.context, C.GLenum(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), C.GLsizei(primcount))
	}

	glc.MultiDrawElements = func(mode uint32, count *int32, Type uint32, indices unsafe.Pointer, primcount int32) {
		defer glc.trace("MultiDrawElements")
		C.gl10MultiDrawElements(glc.context, C.GLenum(mode), (*C.GLsizei)(unsafe.Pointer(count)), C.GLenum(Type), indices, C.GLsizei(primcount))
	}

	glc.PointParameterf = func(pname uint32, param float32) {
		defer glc.trace("PointParameterf")
		C.gl10PointParameterf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PointParameteri = func(pname uint32, param int32) {
		defer glc.trace("PointParameteri")
		C.gl10PointParameteri(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.SecondaryColor3b = func(red, green, blue int8) {
		defer glc.trace("SecondaryColor3b")
		C.gl10SecondaryColor3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.SecondaryColor3s = func(red, green, blue int16) {
		defer glc.trace("SecondaryColor3s")
		C.gl10SecondaryColor3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.SecondaryColor3i = func(red, green, blue int32) {
		defer glc.trace("SecondaryColor3i")
		C.gl10SecondaryColor3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.SecondaryColor3f = func(red, green, blue float32) {
		defer glc.trace("SecondaryColor3f")
		C.gl10SecondaryColor3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.SecondaryColor3d = func(red, green, blue float64) {
		defer glc.trace("SecondaryColor3d")
		C.gl10SecondaryColor3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.SecondaryColor3ub = func(red, green, blue uint8) {
		defer glc.trace("SecondaryColor3ub")
		C.gl10SecondaryColor3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.SecondaryColor3us = func(red, green, blue uint16) {
		defer glc.trace("SecondaryColor3us")
		C.gl10SecondaryColor3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.SecondaryColor3ui = func(red, green, blue uint32) {
		defer glc.trace("SecondaryColor3ui")
		C.gl10SecondaryColor3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.SecondaryColor3bv = func(v *int8) {
		defer glc.trace("SecondaryColor3bv")
		C.gl10SecondaryColor3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3sv = func(v *int16) {
		defer glc.trace("SecondaryColor3sv")
		C.gl10SecondaryColor3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3iv = func(v *int32) {
		defer glc.trace("SecondaryColor3iv")
		C.gl10SecondaryColor3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3fv = func(v *float32) {
		defer glc.trace("SecondaryColor3fv")
		C.gl10SecondaryColor3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3dv = func(v *float64) {
		defer glc.trace("SecondaryColor3dv")
		C.gl10SecondaryColor3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3ubv = func(v *uint8) {
		defer glc.trace("SecondaryColor3ubv")
		C.gl10SecondaryColor3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3usv = func(v *uint16) {
		defer glc.trace("SecondaryColor3usv")
		C.gl10SecondaryColor3usv(glc.context, (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.SecondaryColor3uiv = func(v *uint32) {
		defer glc.trace("SecondaryColor3uiv")
		C.gl10SecondaryColor3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.SecondaryColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("SecondaryColorPointer")
		C.gl10SecondaryColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.WindowPos2s = func(x, y int16) {
		defer glc.trace("WindowPos2s")
		C.gl10WindowPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.WindowPos2i = func(x, y int32) {
		defer glc.trace("WindowPos2i")
		C.gl10WindowPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.WindowPos2f = func(x, y float32) {
		defer glc.trace("WindowPos2f")
		C.gl10WindowPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.WindowPos2d = func(x, y float64) {
		defer glc.trace("WindowPos2d")
		C.gl10WindowPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.WindowPos3s = func(x, y, z int16) {
		defer glc.trace("WindowPos3s")
		C.gl10WindowPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.WindowPos3i = func(x, y, z int32) {
		defer glc.trace("WindowPos3i")
		C.gl10WindowPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.WindowPos3f = func(x, y, z float32) {
		defer glc.trace("WindowPos3f")
		C.gl10WindowPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.WindowPos3d = func(x, y, z float64) {
		defer glc.trace("WindowPos3d")
		C.gl10WindowPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.WindowPos2sv = func(v *int16) {
		defer glc.trace("WindowPos2sv")
		C.gl10WindowPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos2iv = func(v *int32) {
		defer glc.trace("WindowPos2iv")
		C.gl10WindowPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos2fv = func(v *float32) {
		defer glc.trace("WindowPos2fv")
		C.gl10WindowPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos2dv = func(v *float64) {
		defer glc.trace("WindowPos2dv")
		C.gl10WindowPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.WindowPos3sv = func(v *int16) {
		defer glc.trace("WindowPos3sv")
		C.gl10WindowPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.WindowPos3iv = func(v *int32) {
		defer glc.trace("WindowPos3iv")
		C.gl10WindowPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.WindowPos3fv = func(v *float32) {
		defer glc.trace("WindowPos3fv")
		C.gl10WindowPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.WindowPos3dv = func(v *float64) {
		defer glc.trace("WindowPos3dv")
		C.gl10WindowPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		defer glc.trace("BeginQuery")
		C.gl10BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		defer glc.trace("BindBuffer")
		C.gl10BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		defer glc.trace("BufferData")
		C.gl10BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		defer glc.trace("BufferSubData")
		C.gl10BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.DeleteBuffers = func(n int32, buffers *uint32) {
		defer glc.trace("DeleteBuffers")
		C.gl10DeleteBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.DeleteQueries = func(n int32, ids *uint32) {
		defer glc.trace("DeleteQueries")
		C.gl10DeleteQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GenBuffers = func(n int32, buffers *uint32) {
		defer glc.trace("GenBuffers")
		C.gl10GenBuffers(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}

	glc.GenQueries = func(n int32, ids *uint32) {
		defer glc.trace("GenQueries")
		C.gl10GenQueries(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(ids)))
	}

	glc.GetBufferParameteriv = func(target, value uint32, data *int32) {
		defer glc.trace("GetBufferParameteriv")
		C.gl10GetBufferParameteriv(glc.context, C.GLenum(target), C.GLenum(value), (*C.GLint)(unsafe.Pointer(data)))
	}

	glc.GetBufferPointerv = func(target, pname uint32, params unsafe.Pointer) {
		defer glc.trace("GetBufferPointerv")
		C.gl10GetBufferPointerv(glc.context, C.GLenum(target), C.GLenum(pname), params)
	}

	glc.GetBufferSubData = func(target uint32, offset int32, size int32, data unsafe.Pointer) {
		defer glc.trace("GetBufferSubData")
		C.gl10GetBufferSubData(glc.context, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}

	glc.GetQueryObjectiv = func(id uint32, pname uint32, params *int32) {
		defer glc.trace("GetQueryObjectiv")
		C.gl10GetQueryObjectiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetQueryObjectuiv = func(id uint32, pname uint32, params *uint32) {
		defer glc.trace("GetQueryObjectuiv")
		C.gl10GetQueryObjectuiv(glc.context, C.GLuint(id), C.GLenum(pname), (*C.GLuint)(unsafe.Pointer(params)))
	}

	glc.GetQueryiv = func(target, pname uint32, params *int32) {
		defer glc.trace("GetQueryiv")
		C.gl10GetQueryiv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.IsBuffer = func(buffer uint32) bool {
		defer glc.trace("IsBuffer")
		return C.gl10IsBuffer(glc.context, C.GLuint(buffer)) != 0
	}

	glc.IsQuery = func(id uint32) bool {
		defer glc.trace("IsQuery")
		return C.gl10IsQuery(glc.context, C.GLuint(id)) != 0
	}

	glc.MapBuffer = func(target, access uint32) unsafe.Pointer {
		defer glc.trace("MapBuffer")
		return unsafe.Pointer(C.gl10MapBuffer(glc.context, C.GLenum(target), C.GLenum(access)))
	}

	glc.UnmapBuffer = func(target uint32) bool {
		defer glc.trace("UnmapBuffer")
		return C.gl10UnmapBuffer(glc.context, C.GLenum(target)) != 0
	}

	glc.AttachShader = func(program, shader uint32) {
		defer glc.trace("AttachShader")
		C.gl10AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		defer glc.trace("BindAttribLocation")
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl10BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha uint32) {
		defer glc.trace("BlendEquationSeperate")
		C.gl10BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		defer glc.trace("CompileShader")
		C.gl10CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		defer glc.trace("CreateProgram")
		return uint32(C.gl10CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType uint32) uint32 {
		defer glc.trace("CreateShader")
		return uint32(C.gl10CreateShader(glc.context, C.GLenum(shaderType)))
	}

	glc.DeleteProgram = func(program uint32) {
		defer glc.trace("DeleteProgram")
		C.gl10DeleteProgram(glc.context, C.GLuint(program))
	}

	glc.DeleteShader = func(shader uint32) {
		defer glc.trace("DeleteShader")
		C.gl10DeleteShader(glc.context, C.GLuint(shader))
	}

	glc.DetachShader = func(program, shader uint32) {
		defer glc.trace("DetachShader")
		C.gl10DetachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.EnableVertexAttribArray = func(index uint32) {
		defer glc.trace("EnableVertexAttribArray")
		C.gl10EnableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DisableVertexAttribArray = func(index uint32) {
		defer glc.trace("DisableVertexAttribArray")
		C.gl10DisableVertexAttribArray(glc.context, C.GLuint(index))
	}

	glc.DrawBuffers = func(n int32, bufs *uint32) {
		defer glc.trace("DrawBuffers")
		C.gl10DrawBuffers(glc.context, C.GLsizei(n), (*C.GLenum)(unsafe.Pointer(bufs)))
	}

	glc.GetActiveAttrib = func(program, index uint32, bufSize int32) (length int32, size int32, Type uint32, name string) {
		defer glc.trace("GetActiveAttrib")
		var (
			cname C.GLchar
		)
		C.gl10GetActiveAttrib(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(&length)), (*C.GLint)(unsafe.Pointer(&size)), (*C.GLenum)(unsafe.Pointer(&Type)), &cname)
		name = C.GoString((*C.char)(unsafe.Pointer(&cname)))
		return
	}

	glc.GetActiveUniform = func(program, index uint32, bufSize int32, length *int32, size *int32, Type *uint32, name *byte) {
		defer glc.trace("GetActiveUniform")
		C.gl10GetActiveUniform(glc.context, C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(Type)), (*C.GLchar)(unsafe.Pointer(name)))
	}

	glc.GetAttachedShaders = func(program uint32, maxCount int32, count *int32, shaders *uint32) {
		defer glc.trace("GetAttachedShaders")
		C.gl10GetAttachedShaders(glc.context, C.GLuint(program), C.GLsizei(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
	}

	glc.GetAttribLocation = func(program uint32, name *byte) int32 {
		defer glc.trace("GetAttribLocation")
		return int32(C.gl10GetAttribLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetProgramiv = func(program uint32, pname uint32, params *int32) {
		defer glc.trace("GetProgramiv")
		C.gl10GetProgramiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetProgramInfoLog = func(program uint32, maxLength int32, length *int32, infoLog *byte) {
		defer glc.trace("GetProgramInfoLog")
		C.gl10GetProgramInfoLog(glc.context, C.GLuint(program), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderiv = func(program uint32, pname uint32, params *int32) {
		defer glc.trace("GetShaderiv")
		C.gl10GetShaderiv(glc.context, C.GLuint(program), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetShaderInfoLog = func(shader uint32, maxLength int32, length *int32, infoLog *byte) {
		defer glc.trace("GetShaderInfoLog")
		C.gl10GetShaderInfoLog(glc.context, C.GLuint(shader), C.GLsizei(maxLength), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
	}

	glc.GetShaderSource = func(shader uint32, bufSize int32, length *int32, source *byte) {
		defer glc.trace("GetShaderSource")
		C.gl10GetShaderSource(glc.context, C.GLuint(shader), C.GLsizei(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
	}

	glc.GetUniformfv = func(program uint32, location int32, params *float32) {
		defer glc.trace("GetUniformfv")
		C.gl10GetUniformfv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetUniformiv = func(program uint32, location int32, params *int32) {
		defer glc.trace("GetUniformiv")
		C.gl10GetUniformiv(glc.context, C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetUniformLocation = func(program uint32, name *byte) int32 {
		defer glc.trace("GetUniformLocation")
		return int32(C.gl10GetUniformLocation(glc.context, C.GLuint(program), (*C.GLchar)(unsafe.Pointer(name))))
	}

	glc.GetVertexAttribdv = func(index uint32, pname uint32, params *float64) {
		defer glc.trace("GetVertexAttribdv")
		C.gl10GetVertexAttribdv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribfv = func(index uint32, pname uint32, params *float32) {
		defer glc.trace("GetVertexAttribfv")
		C.gl10GetVertexAttribfv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribiv = func(index uint32, pname uint32, params *int32) {
		defer glc.trace("GetVertexAttribiv")
		C.gl10GetVertexAttribiv(glc.context, C.GLuint(index), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}

	glc.GetVertexAttribPointerv = func(index uint32, pname uint32, pointer unsafe.Pointer) {
		defer glc.trace("GetVertexAttribPointerv")
		C.gl10GetVertexAttribPointerv(glc.context, C.GLuint(index), C.GLenum(pname), pointer)
	}

	glc.IsProgram = func(program uint32) bool {
		defer glc.trace("IsProgram")
		return C.gl10IsProgram(glc.context, C.GLuint(program)) != 0
	}

	glc.IsShader = func(shader uint32) bool {
		defer glc.trace("IsShader")
		return C.gl10IsShader(glc.context, C.GLuint(shader)) != 0
	}

	glc.LinkProgram = func(program uint32) {
		defer glc.trace("LinkProgram")
		C.gl10LinkProgram(glc.context, C.GLuint(program))
	}

	glc.ShaderSource = func(shader uint32, count int32, string **byte, length *int32) {
		defer glc.trace("ShaderSource")
		C.gl10ShaderSource(glc.context, C.GLuint(shader), C.GLsizei(count), (**C.GLchar)(unsafe.Pointer(string)), (*C.GLint)(unsafe.Pointer(length)))
	}

	glc.StencilFuncSeparate = func(face, Func uint32, ref int32, mask uint32) {
		defer glc.trace("StencilFuncSeparate")
		C.gl10StencilFuncSeparate(glc.context, C.GLenum(face), C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMaskSeparate = func(face uint32, mask uint32) {
		defer glc.trace("StencilMaskSeparate")
		C.gl10StencilMaskSeparate(glc.context, C.GLenum(face), C.GLuint(mask))
	}

	glc.StencilOpSeparate = func(face, sfail, dpfail, dppass uint32) {
		defer glc.trace("StencilOpSeparate")
		C.gl10StencilOpSeparate(glc.context, C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
	}

	glc.Uniform1f = func(location int32, v0 float32) {
		defer glc.trace("Uniform1f")
		C.gl10Uniform1f(glc.context, C.GLint(location), C.GLfloat(v0))
	}

	glc.Uniform2f = func(location int32, v0, v1 float32) {
		defer glc.trace("Uniform2f")
		C.gl10Uniform2f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.Uniform3f = func(location int32, v0, v1, v2 float32) {
		defer glc.trace("Uniform3f")
		C.gl10Uniform3f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Uniform4f = func(location int32, v0, v1, v2, v3 float32) {
		defer glc.trace("Uniform4f")
		C.gl10Uniform4f(glc.context, C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.Uniform1i = func(location, v0 int32) {
		defer glc.trace("Uniform1i")
		C.gl10Uniform1i(glc.context, C.GLint(location), C.GLint(v0))
	}

	glc.Uniform2i = func(location, v0, v1 int32) {
		defer glc.trace("Uniform2i")
		C.gl10Uniform2i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1))
	}

	glc.Uniform3i = func(location, v0, v1, v2 int32) {
		defer glc.trace("Uniform3i")
		C.gl10Uniform3i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
	}

	glc.Uniform4i = func(location, v0, v1, v2, v3 int32) {
		defer glc.trace("Uniform4i")
		C.gl10Uniform4i(glc.context, C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
	}

	glc.Uniform1ui = func(location int32, v0 uint32) {
		defer glc.trace("Uniform1ui")
		C.gl10Uniform1ui(glc.context, C.GLint(location), C.GLuint(v0))
	}

	glc.Uniform2ui = func(location int32, v0, v1 uint32) {
		defer glc.trace("Uniform2ui")
		C.gl10Uniform2ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1))
	}

	glc.Uniform3ui = func(location int32, v0, v1, v2 uint32) {
		defer glc.trace("Uniform3ui")
		C.gl10Uniform3ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
	}

	glc.Uniform4ui = func(location int32, v0, v1, v2, v3 uint32) {
		defer glc.trace("Uniform4ui")
		C.gl10Uniform4ui(glc.context, C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
	}

	glc.Uniform1fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform1fv")
		C.gl10Uniform1fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform2fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform2fv")
		C.gl10Uniform2fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform3fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform3fv")
		C.gl10Uniform3fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform4fv = func(location int32, count int32, value *float32) {
		defer glc.trace("Uniform4fv")
		C.gl10Uniform4fv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.Uniform1iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform1iv")
		C.gl10Uniform1iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform2iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform2iv")
		C.gl10Uniform2iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform3iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform3iv")
		C.gl10Uniform3iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform4iv = func(location int32, count int32, value *int32) {
		defer glc.trace("Uniform4iv")
		C.gl10Uniform4iv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLint)(unsafe.Pointer(value)))
	}

	glc.Uniform1uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform1uiv")
		C.gl10Uniform1uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform2uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform2uiv")
		C.gl10Uniform2uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform3uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform3uiv")
		C.gl10Uniform3uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.Uniform4uiv = func(location int32, count int32, value *uint32) {
		defer glc.trace("Uniform4uiv")
		C.gl10Uniform4uiv(glc.context, C.GLint(location), C.GLsizei(count), (*C.GLuint)(unsafe.Pointer(value)))
	}

	glc.UseProgram = func(program uint32) {
		defer glc.trace("UseProgram")
		C.gl10UseProgram(glc.context, C.GLuint(program))
	}

	glc.ValidateProgram = func(program uint32) {
		defer glc.trace("ValidateProgram")
		C.gl10ValidateProgram(glc.context, C.GLuint(program))
	}

	glc.VertexAttribPointer = func(index uint32, size int32, Type uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
		defer glc.trace("VertexAttribPointer")
		C.gl10VertexAttribPointer(glc.context, C.GLuint(index), C.GLint(size), C.GLenum(Type), boolToGL(normalized), C.GLsizei(stride), pointer)
	}

	glc.VertexAttrib1f = func(index uint32, v0 float32) {
		defer glc.trace("VertexAttrib1f")
		C.gl10VertexAttrib1f(glc.context, C.GLuint(index), C.GLfloat(v0))
	}

	glc.VertexAttrib1s = func(index uint32, v0 int16) {
		defer glc.trace("VertexAttrib1s")
		C.gl10VertexAttrib1s(glc.context, C.GLuint(index), C.GLshort(v0))
	}

	glc.VertexAttrib1d = func(index uint32, v0 float64) {
		defer glc.trace("VertexAttrib1d")
		C.gl10VertexAttrib1d(glc.context, C.GLuint(index), C.GLdouble(v0))
	}

	glc.VertexAttrib2f = func(index uint32, v0, v1 float32) {
		defer glc.trace("VertexAttrib2f")
		C.gl10VertexAttrib2f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1))
	}

	glc.VertexAttrib2s = func(index uint32, v0, v1 int16) {
		defer glc.trace("VertexAttrib2s")
		C.gl10VertexAttrib2s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1))
	}

	glc.VertexAttrib2d = func(index uint32, v0, v1 float64) {
		defer glc.trace("VertexAttrib2d")
		C.gl10VertexAttrib2d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1))
	}

	glc.VertexAttrib3f = func(index uint32, v0, v1, v2 float32) {
		defer glc.trace("VertexAttrib3f")
		C.gl10VertexAttrib3f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.VertexAttrib3s = func(index uint32, v0, v1, v2 int16) {
		defer glc.trace("VertexAttrib3s")
		C.gl10VertexAttrib3s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2))
	}

	glc.VertexAttrib3d = func(index uint32, v0, v1, v2 float64) {
		defer glc.trace("VertexAttrib3d")
		C.gl10VertexAttrib3d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.VertexAttrib4f = func(index uint32, v0, v1, v2, v3 float32) {
		defer glc.trace("VertexAttrib4f")
		C.gl10VertexAttrib4f(glc.context, C.GLuint(index), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
	}

	glc.VertexAttrib4s = func(index uint32, v0, v1, v2, v3 int16) {
		defer glc.trace("VertexAttrib4s")
		C.gl10VertexAttrib4s(glc.context, C.GLuint(index), C.GLshort(v0), C.GLshort(v1), C.GLshort(v2), C.GLshort(v3))
	}

	glc.VertexAttrib4d = func(index uint32, v0, v1, v2, v3 float64) {
		defer glc.trace("VertexAttrib4d")
		C.gl10VertexAttrib4d(glc.context, C.GLuint(index), C.GLdouble(v0), C.GLdouble(v1), C.GLdouble(v2), C.GLdouble(v3))
	}

	glc.VertexAttrib4Nuv = func(index uint32, v0, v1, v2, v3 uint8) {
		defer glc.trace("VertexAttrib4Nuv")
		C.gl10VertexAttrib4Nuv(glc.context, C.GLuint(index), C.GLubyte(v0), C.GLubyte(v1), C.GLubyte(v2), C.GLubyte(v3))
	}

	glc.VertexAttrib1fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib1fv")
		C.gl10VertexAttrib1fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib1sv")
		C.gl10VertexAttrib1sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib1dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib1dv")
		C.gl10VertexAttrib1dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib2fv")
		C.gl10VertexAttrib2fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib2sv")
		C.gl10VertexAttrib2sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib2dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib2dv")
		C.gl10VertexAttrib2dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib3fv")
		C.gl10VertexAttrib3fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib3sv")
		C.gl10VertexAttrib3sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib3dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib3dv")
		C.gl10VertexAttrib3dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4fv = func(index uint32, v *float32) {
		defer glc.trace("VertexAttrib4fv")
		C.gl10VertexAttrib4fv(glc.context, C.GLuint(index), (*C.GLfloat)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4sv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib4sv")
		C.gl10VertexAttrib4sv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4dv = func(index uint32, v *float64) {
		defer glc.trace("VertexAttrib4dv")
		C.gl10VertexAttrib4dv(glc.context, C.GLuint(index), (*C.GLdouble)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4iv = func(index uint32, v *int32) {
		defer glc.trace("VertexAttrib4iv")
		C.gl10VertexAttrib4iv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4bv = func(index uint32, v *int8) {
		defer glc.trace("VertexAttrib4bv")
		C.gl10VertexAttrib4bv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4ubv = func(index uint32, v *uint8) {
		defer glc.trace("VertexAttrib4ubv")
		C.gl10VertexAttrib4ubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4usv = func(index uint32, v *uint16) {
		defer glc.trace("VertexAttrib4usv")
		C.gl10VertexAttrib4usv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4uiv = func(index uint32, v *uint32) {
		defer glc.trace("VertexAttrib4uiv")
		C.gl10VertexAttrib4uiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nbv = func(index uint32, v *int8) {
		defer glc.trace("VertexAttrib4Nbv")
		C.gl10VertexAttrib4Nbv(glc.context, C.GLuint(index), (*C.GLbyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nsv = func(index uint32, v *int16) {
		defer glc.trace("VertexAttrib4Nsv")
		C.gl10VertexAttrib4Nsv(glc.context, C.GLuint(index), (*C.GLshort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Niv = func(index uint32, v *int32) {
		defer glc.trace("VertexAttrib4Niv")
		C.gl10VertexAttrib4Niv(glc.context, C.GLuint(index), (*C.GLint)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nubv = func(index uint32, v *uint8) {
		defer glc.trace("VertexAttrib4Nubv")
		C.gl10VertexAttrib4Nubv(glc.context, C.GLuint(index), (*C.GLubyte)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nusv = func(index uint32, v *uint16) {
		defer glc.trace("VertexAttrib4Nusv")
		C.gl10VertexAttrib4Nusv(glc.context, C.GLuint(index), (*C.GLushort)(unsafe.Pointer(v)))
	}

	glc.VertexAttrib4Nuiv = func(index uint32, v *uint32) {
		defer glc.trace("VertexAttrib4Nuiv")
		C.gl10VertexAttrib4Nuiv(glc.context, C.GLuint(index), (*C.GLuint)(unsafe.Pointer(v)))
	}

	glc.UniformMatrix2fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix2fv")
		C.gl10UniformMatrix2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix3fv")
		C.gl10UniformMatrix3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix4fv")
		C.gl10UniformMatrix4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x3fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix2x3fv")
		C.gl10UniformMatrix2x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x2fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix3x2fv")
		C.gl10UniformMatrix3x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix2x4fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix2x4fv")
		C.gl10UniformMatrix2x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x2fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix4x2fv")
		C.gl10UniformMatrix4x2fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix3x4fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix3x4fv")
		C.gl10UniformMatrix3x4fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	glc.UniformMatrix4x3fv = func(location int32, count int32, transpose bool, value *float32) {
		defer glc.trace("UniformMatrix4x3fv")
		C.gl10UniformMatrix4x3fv(glc.context, C.GLint(location), C.GLsizei(count), boolToGL(transpose), (*C.GLfloat)(unsafe.Pointer(value)))
	}

	if !versionSupported(glc) {
		return nil
	}
	glc.queryExtensions()
	return glc
}
