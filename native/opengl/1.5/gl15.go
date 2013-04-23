package opengl

// #cgo LDFLAGS: -lopengl32
// #include "gl15.h"
import "C"

import "unsafe"

func boolToGL(b bool) C.GLboolean {
	if b {
		return C.GLboolean(1)
	}
	return C.GLboolean(0)
}

const (
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	DEBUG_OBJECT_MESA                                          = 0x8759
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	VERSION_2_1                                                = 1
	MAP_READ_BIT_EXT                                           = 0x0001
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	STENCIL_INDEX1                                             = 0x8D46
	DECR                                                       = 0x1E03
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	OPERAND1_RGB_EXT                                           = 0x8591
	SKIP_COMPONENTS2_NV                                        = -5
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	MODELVIEW29_ARB                                            = 0x873D
	FLOAT_MAT4_ARB                                             = 0x8B5C
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	DOUBLE_MAT4                                                = 0x8F48
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	DRAW_BUFFER9_NV                                            = 0x882E
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	CLIP_PLANE1                                                = 0x3001
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	COLOR4_BIT_PGI                                             = 0x00020000
	SGIS_sharpen_texture                                       = 1
	LIST_PRIORITY_SGIX                                         = 0x8182
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	REG_15_ATI                                                 = 0x8930
	MITER_REVERT_NV                                            = 0x90A7
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	FUNC_SUBTRACT_OES                                          = 0x800A
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	SGIX_tag_sample_buffer                                     = 1
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	SGIX_texture_lod_bias                                      = 1
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	SEPARABLE_2D_EXT                                           = 0x8012
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	MODULATE_ADD_ATI                                           = 0x8744
	DRAW_BUFFER15_ARB                                          = 0x8834
	DEPTH32F_STENCIL8                                          = 0x8CAD
	LUMINANCE4                                                 = 0x803F
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	MODELVIEW21_ARB                                            = 0x8735
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	CON_23_ATI                                                 = 0x8958
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	DISCRETE_AMD                                               = 0x9006
	CURRENT_RASTER_POSITION                                    = 0x0B07
	S                                                          = 0x2000
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	RGB10_A2                                                   = 0x8059
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	DRAW_BUFFER8                                               = 0x882D
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	DEPTH_EXT                                                  = 0x1801
	INT_VEC4                                                   = 0x8B55
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	FIRST_TO_REST_NV                                           = 0x90AF
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	DOUBLE_MAT3                                                = 0x8F47
	SGIS_texture4D                                             = 1
	CURRENT_BIT                                                = 0x00000001
	OR                                                         = 0x1507
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	TEXTURE11_ARB                                              = 0x84CB
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	MODELVIEW9_ARB                                             = 0x8729
	BUFFER_USAGE                                               = 0x8765
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	OBJECT_POINT_SGIS                                          = 0x81F5
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	AUX0                                                       = 0x0409
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	MAP2_NORMAL                                                = 0x0DB2
	RGB5_A1                                                    = 0x8057
	BLEND_DST_ALPHA                                            = 0x80CA
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	LOW_INT                                                    = 0x8DF3
	POINT_SIZE_RANGE                                           = 0x0B12
	MAP1_VERTEX_3                                              = 0x0D97
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	SGIX_fragment_lighting                                     = 1
	COLOR_EXT                                                  = 0x1800
	PROXY_COLOR_TABLE                                          = 0x80D3
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	ALPHA                                                      = 0x1906
	MAX_WIDTH                                                  = 0x827E
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	DRAW_BUFFER9_ATI                                           = 0x882E
	RENDERBUFFER                                               = 0x8D41
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	MAX_NAME_LENGTH                                            = 0x92F6
	LUMINANCE                                                  = 0x1909
	SCALE_BY_TWO_NV                                            = 0x853E
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	TEXTURE_BINDING_2D                                         = 0x8069
	RGB4                                                       = 0x804F
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	TEXTURE2                                                   = 0x84C2
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	TEXTURE_ENV_MODE                                           = 0x2200
	LIGHT3                                                     = 0x4003
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	SIGNED_HILO16_NV                                           = 0x86FA
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	FRAMEBUFFER_OES                                            = 0x8D40
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	OPERAND2_RGB_ARB                                           = 0x8592
	WEIGHT_ARRAY_OES                                           = 0x86AD
	VERTEX_STREAM5_ATI                                         = 0x8771
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	BUFFER_DATA_SIZE                                           = 0x9303
	DEPTH_WRITEMASK                                            = 0x0B72
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	TESS_CONTROL_SHADER                                        = 0x8E88
	DATA_BUFFER_AMD                                            = 0x9151
	CLEAR                                                      = 0x1500
	DEPTH_COMPONENT16                                          = 0x81A5
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	YCBAYCR8A_4224_NV                                          = 0x9032
	PATH_JOIN_STYLE_NV                                         = 0x9079
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	ARB_imaging                                                = 1
	TEXTURE_BLUE_SIZE                                          = 0x805E
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	MATRIX1_ARB                                                = 0x88C1
	TEXTURE_BINDING_3D                                         = 0x806A
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	FRAMEBUFFER                                                = 0x8D40
	DEBUG_OUTPUT                                               = 0x92E0
	TEXTURE28_ARB                                              = 0x84DC
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	BOOL_ARB                                                   = 0x8B56
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	RECT_NV                                                    = 0xF6
	LOAD                                                       = 0x0101
	MAX                                                        = 0x8008
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	MAP1_BINORMAL_EXT                                          = 0x8446
	STENCIL_INDEX8                                             = 0x8D48
	RGBA32F                                                    = 0x8814
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	SGIX_interlace                                             = 1
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	OP_SET_GE_EXT                                              = 0x878C
	READ_ONLY                                                  = 0x88B8
	T2F_IUI_V2F_EXT                                            = 0x81B1
	ADD_SIGNED                                                 = 0x8574
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	DRAW_BUFFER0_ATI                                           = 0x8825
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	IMAGE_BINDING_NAME                                         = 0x8F3A
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	R                                                          = 0x2002
	TEXTURE18                                                  = 0x84D2
	BLEND_SRC_RGB_OES                                          = 0x80C9
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	ORDER                                                      = 0x0A01
	COLOR_LOGIC_OP                                             = 0x0BF2
	MIRRORED_REPEAT_OES                                        = 0x8370
	PN_TRIANGLES_ATI                                           = 0x87F0
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	INVARIANT_VALUE_EXT                                        = 0x87EA
	CURRENT_QUERY                                              = 0x8865
	BOOL_VEC3                                                  = 0x8B58
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	SGIS_texture_lod                                           = 1
	SGIX_calligraphic_fragment                                 = 1
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	RGB8UI                                                     = 0x8D7D
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	LINE_TO_NV                                                 = 0x04
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	ALPHA_SCALE                                                = 0x0D1C
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	ARRAY_BUFFER                                               = 0x8892
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	GL_3DC_XY_AMD                                              = 0x87FA
	DRAW_BUFFER15_NV                                           = 0x8834
	TIME_ELAPSED                                               = 0x88BF
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	INT                                                        = 0x1404
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	UNSIGNED_NORMALIZED                                        = 0x8C17
	MAX_SAMPLES_EXT                                            = 0x8D57
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	FENCE_CONDITION_NV                                         = 0x84F4
	MODELVIEW17_ARB                                            = 0x8731
	OBJECT_TYPE_ARB                                            = 0x8B4E
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	ALPHA_SNORM                                                = 0x9010
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	SECONDARY_COLOR_NV                                         = 0x852D
	COLOR_ATTACHMENT4                                          = 0x8CE4
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	GREATER                                                    = 0x0204
	SAMPLER                                                    = 0x82E6
	E_TIMES_F_NV                                               = 0x8531
	SOURCE0_ALPHA_EXT                                          = 0x8588
	QUERY_RESULT                                               = 0x8866
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	RENDER                                                     = 0x1C00
	DRAW_BUFFER13                                              = 0x8832
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	BUFFER_VARIABLE                                            = 0x92E5
	QUERY_WAIT                                                 = 0x8E13
	NORMAL_ARRAY_TYPE                                          = 0x807E
	RGB_FLOAT16_APPLE                                          = 0x881B
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	VERTEX_STREAM3_ATI                                         = 0x876F
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	CON_12_ATI                                                 = 0x894D
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	SAMPLES_SGIS                                               = 0x80A9
	SRGB_WRITE                                                 = 0x8298
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	DRAW_BUFFER12_ARB                                          = 0x8831
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	OBJECT_TYPE                                                = 0x9112
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	SGIS_texture_select                                        = 1
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	REPLACE_VALUE_AMD                                          = 0x874B
	SLUMINANCE                                                 = 0x8C46
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	REFLECTION_MAP_OES                                         = 0x8512
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	FLOAT_RGBA_NV                                              = 0x8883
	SEPARATE_ATTRIBS                                           = 0x8C8D
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	DYNAMIC_READ                                               = 0x88E9
	RENDERBUFFER_EXT                                           = 0x8D41
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	LIST_BASE                                                  = 0x0B32
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	INT8_VEC3_NV                                               = 0x8FE2
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	MAX_LABEL_LENGTH                                           = 0x82E8
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	DRAW_BUFFER13_NV                                           = 0x8832
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	BOUNDING_BOX_NV                                            = 0x908D
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	STREAM_READ_ARB                                            = 0x88E1
	INT16_VEC4_NV                                              = 0x8FE7
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	FOG_COORDINATE                                             = 0x8451
	SAMPLER_2D_SHADOW                                          = 0x8B62
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	DS_SCALE_NV                                                = 0x8710
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	MODELVIEW0_ARB                                             = 0x1700
	POINT_SPRITE_NV                                            = 0x8861
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	CONSTANT_ALPHA                                             = 0x8003
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	GL_4PASS_2_EXT                                             = 0x80A6
	LINE_BIT                                                   = 0x00000004
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	CONST_EYE_NV                                               = 0x86E5
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	V3F                                                        = 0x2A21
	COLOR_TABLE                                                = 0x80D0
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	BACK                                                       = 0x0405
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	C4UB_V3F                                                   = 0x2A23
	INTENSITY16_EXT                                            = 0x804D
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	SRC0_ALPHA                                                 = 0x8588
	SHADER_CONSISTENT_NV                                       = 0x86DD
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	COLOR_SUM                                                  = 0x8458
	TEXTURE13                                                  = 0x84CD
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	INTERPOLATE_ARB                                            = 0x8575
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	STENCIL_BACK_FAIL                                          = 0x8801
	COLOR_ATTACHMENT14                                         = 0x8CEE
	LUMINANCE32I_EXT                                           = 0x8D86
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	TEXTURE5_ARB                                               = 0x84C5
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	RGBA32I                                                    = 0x8D82
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	FLAT                                                       = 0x1D00
	MATRIX3_NV                                                 = 0x8633
	ONE_EXT                                                    = 0x87DE
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	MODELVIEW_MATRIX                                           = 0x0BA6
	SCISSOR_TEST                                               = 0x0C11
	POINT_SIZE_MIN                                             = 0x8126
	RGB32F_ARB                                                 = 0x8815
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	INTENSITY8UI_EXT                                           = 0x8D7F
	MINMAX_SINK_EXT                                            = 0x8030
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	IMAGE_2D_RECT                                              = 0x904F
	ENABLE_BIT                                                 = 0x00002000
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	TYPE                                                       = 0x92FA
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	SRC_ALPHA_SATURATE                                         = 0x0308
	RGBA8_EXT                                                  = 0x8058
	MODELVIEW30_ARB                                            = 0x873E
	FRONT_LEFT                                                 = 0x0400
	FOG_COORDINATE_ARRAY                                       = 0x8457
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	CURRENT_PROGRAM                                            = 0x8B8D
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	LUMINANCE12_ALPHA4                                         = 0x8046
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	SWIZZLE_STR_ATI                                            = 0x8976
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	QUERY_OBJECT_AMD                                           = 0x9153
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	SRGB_READ                                                  = 0x8297
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	COLOR_INDEXES                                              = 0x1603
	CONSTANT_COLOR0_NV                                         = 0x852A
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	SGIX_vertex_preclip                                        = 1
	AMBIENT                                                    = 0x1200
	BLEND_EQUATION_EXT                                         = 0x8009
	COMBINE_EXT                                                = 0x8570
	DS_BIAS_NV                                                 = 0x8716
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	SUBTRACT                                                   = 0x84E7
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	DRAW_BUFFER3                                               = 0x8828
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	INVALID_OPERATION                                          = 0x0502
	TEXTURE_DEPTH_EXT                                          = 0x8071
	DUAL_ALPHA16_SGIS                                          = 0x8113
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	DT_SCALE_NV                                                = 0x8711
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	VALIDATE_STATUS                                            = 0x8B83
	SGIX_impact_pixel_texture                                  = 1
	STENCIL_FAIL                                               = 0x0B94
	PACK_SKIP_PIXELS                                           = 0x0D04
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	RGB10_A2_EXT                                               = 0x8059
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	DOUBLE_VEC3                                                = 0x8FFD
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	SKIP_COMPONENTS4_NV                                        = -3
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	ZERO                                                       = 0
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	RGB12_EXT                                                  = 0x8053
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	NEGATIVE_Y_EXT                                             = 0x87DA
	SGIX_clipmap                                               = 1
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	UNSIGNED_INT16_NV                                          = 0x8FF0
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	GEOMETRY_TEXTURE                                           = 0x829E
	WAIT_FAILED_APPLE                                          = 0x911D
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	OUTPUT_VERTEX_EXT                                          = 0x879A
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	MIN_LOD_WARNING_AMD                                        = 0x919C
	SAMPLE_MASK_EXT                                            = 0x80A0
	CULL_VERTEX_EXT                                            = 0x81AA
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	PRESERVE_ATI                                               = 0x8762
	QUERY_COUNTER_BITS                                         = 0x8864
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	MAX_INTEGER_SAMPLES                                        = 0x9110
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	VERTEX_PROGRAM_ARB                                         = 0x8620
	NEGATIVE_W_EXT                                             = 0x87DC
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	SIGNED_RGBA_NV                                             = 0x86FB
	OP_MOV_EXT                                                 = 0x8799
	SHADER_OBJECT_ARB                                          = 0x8B48
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	C3F_V3F                                                    = 0x2A24
	INTENSITY4_EXT                                             = 0x804A
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	GL_4PASS_0_EXT                                             = 0x80A4
	INCR_WRAP_EXT                                              = 0x8507
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	BOOL                                                       = 0x8B56
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	POLYGON_OFFSET_FILL                                        = 0x8037
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	RGB8_EXT                                                   = 0x8051
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	QUADRATIC_ATTENUATION                                      = 0x1209
	RED_EXT                                                    = 0x1903
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	R16_SNORM                                                  = 0x8F98
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	DT_BIAS_NV                                                 = 0x8717
	MATRIX15_ARB                                               = 0x88CF
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	STENCIL_EXT                                                = 0x1802
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	STENCIL_BACK_FUNC                                          = 0x8800
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	ALPHA12                                                    = 0x803D
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	ITALIC_BIT_NV                                              = 0x02
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	DRAW_BUFFER7_ARB                                           = 0x882C
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	DEPTH_ATTACHMENT                                           = 0x8D00
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	LINEAR_DETAIL_SGIS                                         = 0x8097
	TEXTURE29_ARB                                              = 0x84DD
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	ALPHA32I_EXT                                               = 0x8D84
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	MATRIX24_ARB                                               = 0x88D8
	REG_10_ATI                                                 = 0x892B
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	FRONT_FACE                                                 = 0x0B46
	NORMAL_ARRAY_EXT                                           = 0x8075
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	TEXTURE22                                                  = 0x84D6
	MAP_WRITE_BIT_EXT                                          = 0x0002
	RED_BIAS                                                   = 0x0D15
	MVP_MATRIX_EXT                                             = 0x87E3
	BLUE_BIT_ATI                                               = 0x00000004
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	INDEX_ARRAY_POINTER                                        = 0x8091
	CLAMP_TO_BORDER                                            = 0x812D
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	POINT_SIZE_MAX_ARB                                         = 0x8127
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	SOURCE0_ALPHA_ARB                                          = 0x8588
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	RGB16UI_EXT                                                = 0x8D77
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	GL_2PASS_1_EXT                                             = 0x80A3
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	IUI_V3F_EXT                                                = 0x81AE
	FLOAT_R_NV                                                 = 0x8880
	COLOR_INDEX16_EXT                                          = 0x80E7
	DOT3_ATI                                                   = 0x8966
	SWIZZLE_STRQ_ATI                                           = 0x897A
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	DRAW_BUFFER_EXT                                            = 0x0C01
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	DRAW_BUFFER2                                               = 0x8827
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	SAMPLE_MASK_VALUE                                          = 0x8E52
	R8_EXT                                                     = 0x8229
	ADD_SIGNED_EXT                                             = 0x8574
	MODELVIEW8_ARB                                             = 0x8728
	RELEASED_APPLE                                             = 0x8A19
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	NEVER                                                      = 0x0200
	LEQUAL                                                     = 0x0203
	TESS_GEN_MODE                                              = 0x8E76
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	INDEX_WRITEMASK                                            = 0x0C21
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	MATRIX_EXT                                                 = 0x87C0
	ALPHA16_EXT                                                = 0x803E
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	MATRIX23_ARB                                               = 0x88D7
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	LUMINANCE4_ALPHA4                                          = 0x8043
	RGB9_E5                                                    = 0x8C3D
	BGR_INTEGER                                                = 0x8D9A
	RG16_SNORM                                                 = 0x8F99
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	STENCIL_BUFFER_BIT                                         = 0x00000400
	SPECULAR                                                   = 0x1202
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	RGBA_MODE                                                  = 0x0C31
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	PATH_DASH_OFFSET_NV                                        = 0x907E
	ALPHA_TEST_REF                                             = 0x0BC2
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	COMPRESSED_RG11_EAC                                        = 0x9272
	LIGHT4                                                     = 0x4004
	TEXTURE28                                                  = 0x84DC
	SOURCE3_RGB_NV                                             = 0x8583
	DRAW_BUFFER3_ATI                                           = 0x8828
	MATRIX6_ARB                                                = 0x88C6
	PATH_END_CAPS_NV                                           = 0x9076
	COMPUTE_SUBROUTINE                                         = 0x92ED
	PATCHES                                                    = 0x000E
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	EXPAND_NORMAL_NV                                           = 0x8538
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	MATRIX31_ARB                                               = 0x88DF
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	PACK_ALIGNMENT                                             = 0x0D05
	STENCIL                                                    = 0x1802
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	DEPTH_TEXTURE_MODE                                         = 0x884B
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	DECAL                                                      = 0x2101
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	SYNC_FENCE_APPLE                                           = 0x9116
	PRIMITIVE_RESTART_NV                                       = 0x8558
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	R16F                                                       = 0x822D
	REG_22_ATI                                                 = 0x8937
	AND_REVERSE                                                = 0x1502
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	CULL_FACE                                                  = 0x0B44
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	MAX_TEXTURE_COORDS                                         = 0x8871
	MATRIX25_ARB                                               = 0x88D9
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	SGIX_pixel_tiles                                           = 1
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	CON_7_ATI                                                  = 0x8948
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	ALPHA_TEST_FUNC                                            = 0x0BC1
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	PRESENT_TIME_NV                                            = 0x8E2A
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	RESTART_SUN                                                = 0x0001
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	FLOAT_MAT2_ARB                                             = 0x8B5A
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	MAP1_NORMAL                                                = 0x0D92
	GL_3_BYTES                                                 = 0x1408
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	TEXTURE4_ARB                                               = 0x84C4
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	COMPRESSED_SRGB                                            = 0x8C48
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	EXT_convolution                                            = 1
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	COMBINER3_NV                                               = 0x8553
	NEGATE_BIT_ATI                                             = 0x00000004
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	R11F_G11F_B10F                                             = 0x8C3A
	CW                                                         = 0x0900
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	FIELDS_NV                                                  = 0x8E27
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	ALPHA_TEST_QCOM                                            = 0x0BC0
	INTENSITY4                                                 = 0x804A
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	READ_PIXELS_TYPE                                           = 0x828E
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	MEDIUM_INT                                                 = 0x8DF4
	COMPUTE_SHADER_BIT                                         = 0x00000020
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	MATRIX6_NV                                                 = 0x8636
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	SRGB8_EXT                                                  = 0x8C41
	LUMINANCE32UI_EXT                                          = 0x8D74
	BLEND                                                      = 0x0BE2
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	NUM_PASSES_ATI                                             = 0x8970
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	POINTS                                                     = 0x0000
	TEXTURE21_ARB                                              = 0x84D5
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	SIGNED_ALPHA_NV                                            = 0x8705
	RGBA_FLOAT32_ATI                                           = 0x8814
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	IMAGE_2D                                                   = 0x904D
	PATH_STROKE_MASK_NV                                        = 0x9084
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	ALREADY_SIGNALED                                           = 0x911A
	VIEWPORT                                                   = 0x0BA2
	TEXTURE_BORDER_COLOR                                       = 0x1004
	TEXTURE14_ARB                                              = 0x84CE
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	SINGLE_COLOR                                               = 0x81F9
	SPARE1_NV                                                  = 0x852F
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	INT_IMAGE_1D                                               = 0x9057
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	MATRIX8_ARB                                                = 0x88C8
	FOG_COLOR                                                  = 0x0B66
	DOUBLEBUFFER                                               = 0x0C32
	TEXTURE_ENV                                                = 0x2300
	GL_4PASS_1_SGIS                                            = 0x80A5
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	RED_INTEGER                                                = 0x8D94
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	RGB32I_EXT                                                 = 0x8D83
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	TEXTURE                                                    = 0x1702
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	VERTEX_BLEND_ARB                                           = 0x86A7
	TRANSLATE_Y_NV                                             = 0x908F
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	VERSION_2_0                                                = 1
	SGIX_depth_texture                                         = 1
	TEXTURE1_ARB                                               = 0x84C1
	VERTEX_STREAM1_ATI                                         = 0x876D
	CON_3_ATI                                                  = 0x8944
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	VERTEX_ARRAY_POINTER                                       = 0x808E
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	SAMPLE_COVERAGE                                            = 0x80A0
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	MODELVIEW1_EXT                                             = 0x850A
	RGBA_FLOAT16_APPLE                                         = 0x881A
	REG_12_ATI                                                 = 0x892D
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	VARIABLE_C_NV                                              = 0x8525
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	SPARE0_NV                                                  = 0x852E
	CON_15_ATI                                                 = 0x8950
	CON_24_ATI                                                 = 0x8959
	FLOAT_VEC2_ARB                                             = 0x8B50
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	COLOR_ATTACHMENT1                                          = 0x8CE1
	ADJACENT_PAIRS_NV                                          = 0x90AE
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	MAP_WRITE_BIT                                              = 0x0002
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	COLOR_ATTACHMENT15                                         = 0x8CEF
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	INT_IMAGE_2D_RECT                                          = 0x905A
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	LOCATION                                                   = 0x930E
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	INDEX                                                      = 0x8222
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	OP_LOG_BASE_2_EXT                                          = 0x8792
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	DRAW_BUFFER8_ARB                                           = 0x882D
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	STENCIL_VALUE_MASK                                         = 0x0B93
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	OPERAND1_RGB_ARB                                           = 0x8591
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	FOG_END                                                    = 0x0B64
	CON_4_ATI                                                  = 0x8945
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	PERFMON_RESULT_AMD                                         = 0x8BC6
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	QUARTER_BIT_ATI                                            = 0x00000010
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	COLOR_ARRAY_POINTER                                        = 0x8090
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	TEXTURE16_ARB                                              = 0x84D0
	COMPRESSED_LUMINANCE                                       = 0x84EA
	OP_RECIP_SQRT_EXT                                          = 0x8795
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	PRIMITIVE_ID_NV                                            = 0x8C7C
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	EXT_cmyka                                                  = 1
	INVERTED_SCREEN_W_REND                                     = 0x8491
	OPERAND2_RGB_EXT                                           = 0x8592
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	ALPHA16_SNORM                                              = 0x9018
	LINEAR                                                     = 0x2601
	EYE_POINT_SGIS                                             = 0x81F4
	TEXTURE5                                                   = 0x84C5
	MODELVIEW1_ARB                                             = 0x850A
	UNPACK_RESAMPLE_OML                                        = 0x8985
	GREEN_SCALE                                                = 0x0D18
	R32F                                                       = 0x822E
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	STATIC_COPY                                                = 0x88E6
	HALF_BIT_ATI                                               = 0x00000008
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	SRC_COLOR                                                  = 0x0300
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	BYTE                                                       = 0x1400
	CLIP_DISTANCE4                                             = 0x3004
	CONVOLUTION_FORMAT                                         = 0x8017
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	EVAL_2D_NV                                                 = 0x86C0
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	DRAW_BUFFER2_NV                                            = 0x8827
	NOOP                                                       = 0x1505
	EMBOSS_CONSTANT_NV                                         = 0x855E
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	MODELVIEW15_ARB                                            = 0x872F
	NEGATIVE_X_EXT                                             = 0x87D9
	MATRIX_PALETTE_ARB                                         = 0x8840
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	BGRA_INTEGER_EXT                                           = 0x8D9B
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	DRAW_BUFFER5_NV                                            = 0x882A
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	PROGRAM_STRING_NV                                          = 0x8628
	REG_21_ATI                                                 = 0x8936
	SAMPLER_CUBE_ARB                                           = 0x8B60
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	STENCIL_INDEX16                                            = 0x8D49
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	SHADER                                                     = 0x82E1
	MODELVIEW27_ARB                                            = 0x873B
	RENDERBUFFER_OES                                           = 0x8D41
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	AUX2                                                       = 0x040B
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	IS_PER_PATCH                                               = 0x92E7
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	SCALAR_EXT                                                 = 0x87BE
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	VERSION_3_0                                                = 1
	NOR                                                        = 0x1508
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	EYE_LINE_SGIS                                              = 0x81F6
	MODELVIEW14_ARB                                            = 0x872E
	WRITE_DISCARD_NV                                           = 0x88BE
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	HISTOGRAM_SINK                                             = 0x802D
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	REG_18_ATI                                                 = 0x8933
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	SAMPLE_POSITION                                            = 0x8E50
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	POINT_SIZE                                                 = 0x0B11
	OPERAND0_RGB                                               = 0x8590
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	RGB_FLOAT32_APPLE                                          = 0x8815
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	RGBA8UI_EXT                                                = 0x8D7C
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	EXT_subtexture                                             = 1
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	TEXTURE_3D_OES                                             = 0x806F
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	CONVOLUTION_WIDTH                                          = 0x8018
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	OP_NEGATE_EXT                                              = 0x8783
	TEXTURE_COORD_NV                                           = 0x8C79
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	PATH_DASH_CAPS_NV                                          = 0x907B
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	MAX_SHININESS_NV                                           = 0x8504
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	COMBINER0_NV                                               = 0x8550
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	VERSION_1_3                                                = 1
	GL_3D_COLOR                                                = 0x0602
	EXP2                                                       = 0x0801
	INDEX_ARRAY                                                = 0x8077
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	DEPTH_BOUNDS_EXT                                           = 0x8891
	ISOLINES                                                   = 0x8E7A
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	SAMPLER_OBJECT_AMD                                         = 0x9155
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	TEXTURE24                                                  = 0x84D8
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	DRAW_BUFFER8_NV                                            = 0x882D
	CCW                                                        = 0x0901
	COLOR_INDEX12_EXT                                          = 0x80E6
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	MATRIX22_ARB                                               = 0x88D6
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	PATH_FORMAT_SVG_NV                                         = 0x9070
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	BLEND_COLOR_EXT                                            = 0x8005
	REPLACE_OLDEST_SUN                                         = 0x0003
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	STENCIL_BITS                                               = 0x0D57
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	DSDT_MAG_NV                                                = 0x86F6
	DUDV_ATI                                                   = 0x8779
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	SAMPLE_MASK_NV                                             = 0x8E51
	DOUBLE_MAT2                                                = 0x8F46
	R16                                                        = 0x822A
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	OBJECT_PLANE                                               = 0x2501
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	EMBOSS_MAP_NV                                              = 0x855F
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	R16I                                                       = 0x8233
	TEXTURE13_ARB                                              = 0x84CD
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	MODELVIEW5_ARB                                             = 0x8725
	CND0_ATI                                                   = 0x896B
	RGB32I                                                     = 0x8D83
	SPHERE_MAP                                                 = 0x2402
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	TIMEOUT_EXPIRED                                            = 0x911B
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	NEAREST                                                    = 0x2600
	TEXTURE22_ARB                                              = 0x84D6
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	LUMINANCE8_ALPHA8                                          = 0x8045
	COMBINER_BIAS_NV                                           = 0x8549
	SOURCE0_RGB_EXT                                            = 0x8580
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	LUMINANCE16F_ARB                                           = 0x881E
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	MIRRORED_REPEAT_ARB                                        = 0x8370
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	QUERY_RESULT_ARB                                           = 0x8866
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	VERTEX_PROGRAM_NV                                          = 0x8620
	ZERO_EXT                                                   = 0x87DD
	LUMINANCE16UI_EXT                                          = 0x8D7A
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	HALF_FLOAT                                                 = 0x140B
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	VENDOR                                                     = 0x1F00
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	INDEX_TEST_EXT                                             = 0x81B5
	EYE_RADIAL_NV                                              = 0x855B
	RGB16UI                                                    = 0x8D77
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	TEXTURE4                                                   = 0x84C4
	FLOAT_RG16_NV                                              = 0x8886
	READ_ONLY_ARB                                              = 0x88B8
	RGBA_INTEGER_EXT                                           = 0x8D99
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	SURFACE_REGISTERED_NV                                      = 0x86FD
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	ALPHA16F_ARB                                               = 0x881C
	TEXTURE_WRAP_S                                             = 0x2802
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	TEXTURE12                                                  = 0x84CC
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	FLOAT_MAT2x3                                               = 0x8B65
	HALF_FLOAT_ARB                                             = 0x140B
	TEXTURE_SHADOW                                             = 0x82A1
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	DRAW_BUFFER1                                               = 0x8826
	DOUBLE_MAT2x3                                              = 0x8F49
	NUM_SAMPLE_COUNTS                                          = 0x9380
	POLYGON_MODE                                               = 0x0B40
	INDEX_BITS                                                 = 0x0D51
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	SCREEN_COORDINATES_REND                                    = 0x8490
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	LINE_STRIP                                                 = 0x0003
	INDEX_ARRAY_TYPE                                           = 0x8085
	TEXTURE_WRAP_R                                             = 0x8072
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	VERTEX_STREAM6_ATI                                         = 0x8772
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	NORMAL_MAP                                                 = 0x8511
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	LIGHTING                                                   = 0x0B50
	RGB8                                                       = 0x8051
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	ARRAY_SIZE                                                 = 0x92FB
	PRIMARY_COLOR_NV                                           = 0x852C
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	GL_4PASS_1_EXT                                             = 0x80A5
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	REG_4_ATI                                                  = 0x8925
	CON_9_ATI                                                  = 0x894A
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	SMALL_CW_ARC_TO_NV                                         = 0x14
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	MAP_READ_BIT                                               = 0x0001
	FUNC_SUBTRACT_EXT                                          = 0x800A
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	DOT3_RGB                                                   = 0x86AE
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	LUMINANCE8I_EXT                                            = 0x8D92
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	T2F_V3F                                                    = 0x2A27
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	MATRIX7_ARB                                                = 0x88C7
	SRC1_COLOR                                                 = 0x88F9
	FENCE_APPLE                                                = 0x8A0B
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	LUMINANCE8_EXT                                             = 0x8040
	MODELVIEW18_ARB                                            = 0x8732
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	ARC_TO_NV                                                  = 0xFE
	LIST_BIT                                                   = 0x00020000
	EMISSION                                                   = 0x1600
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	MODELVIEW19_ARB                                            = 0x8733
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	VERSION_3_2                                                = 1
	INDEX_MATERIAL_EXT                                         = 0x81B8
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	COLOR_WRITEMASK                                            = 0x0C23
	SHADER_OPERATION_NV                                        = 0x86DF
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	RG16F                                                      = 0x822F
	COMBINE4_NV                                                = 0x8503
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	AFFINE_2D_NV                                               = 0x9092
	SIGNALED                                                   = 0x9119
	RGBA4_OES                                                  = 0x8056
	MULTISAMPLE                                                = 0x809D
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	MATRIX11_ARB                                               = 0x88CB
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	MAX_SAMPLES_IMG                                            = 0x9135
	TEXTURE_GEN_MODE                                           = 0x2500
	DISCARD_NV                                                 = 0x8530
	MAX_VARYING_VECTORS                                        = 0x8DFC
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	DRAW_BUFFER12_ATI                                          = 0x8831
	FLOAT_R16_NV                                               = 0x8884
	WRITE_ONLY_ARB                                             = 0x88B9
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	EXT_point_parameters                                       = 1
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	MATRIX21_ARB                                               = 0x88D5
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	TEXTURE_RESIDENT                                           = 0x8067
	LINEAR_ATTENUATION                                         = 0x1208
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	LUMINANCE32F_ARB                                           = 0x8818
	DRAW_BUFFER4_ATI                                           = 0x8829
	BOOL_VEC4_ARB                                              = 0x8B59
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	QUERY_WAIT_NV                                              = 0x8E13
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	TESSELLATION_MODE_AMD                                      = 0x9004
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	RGB5_A1_OES                                                = 0x8057
	COLOR_COMPONENTS                                           = 0x8283
	BUFFER_MAP_LENGTH                                          = 0x9120
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	VERTEX_ARRAY                                               = 0x8074
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	TEXTURE_MAX_LOD                                            = 0x813B
	ALPHA_MIN_SGIX                                             = 0x8320
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	SGIS_texture_border_clamp                                  = 1
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	STATE_RESTORE                                              = 0x8BDC
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	PARTIAL_SUCCESS_NV                                         = 0x902E
	DIFFUSE                                                    = 0x1201
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	VERTEX_STREAM4_ATI                                         = 0x8770
	T2F_C3F_V3F                                                = 0x2A2A
	TEXTURE_3D_EXT                                             = 0x806F
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	FOG                                                        = 0x0B60
	TEXTURE_GEN_R                                              = 0x0C62
	DRAW_BUFFER6_ATI                                           = 0x882B
	MATRIX4_ARB                                                = 0x88C4
	SWIZZLE_STQ_ATI                                            = 0x8977
	LUMINANCE16I_EXT                                           = 0x8D8C
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	TRIANGLES                                                  = 0x0004
	POLYGON_SMOOTH                                             = 0x0B41
	RGBA_S3TC                                                  = 0x83A2
	DEPTH24_STENCIL8                                           = 0x88F0
	ADD_BLEND_IMG                                              = 0x8C09
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	ONE                                                        = 1
	GREEN_BITS                                                 = 0x0D53
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	PREVIOUS                                                   = 0x8578
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	ALPHA_INTEGER_EXT                                          = 0x8D97
	BLEND_EQUATION                                             = 0x8009
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	PROGRAM_LENGTH_NV                                          = 0x8627
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	CLIP_DISTANCE0                                             = 0x3000
	REPLACE_MIDDLE_SUN                                         = 0x0002
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	SGIX_shadow_ambient                                        = 1
	COMPRESSED_RGBA                                            = 0x84EE
	OPERAND1_ALPHA_EXT                                         = 0x8599
	DRAW_BUFFER3_ARB                                           = 0x8828
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	TRIANGLE_FAN                                               = 0x0006
	COMPUTE_TEXTURE                                            = 0x82A0
	PRIMARY_COLOR                                              = 0x8577
	OPERAND0_ALPHA_EXT                                         = 0x8598
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	MAX_SUBROUTINES                                            = 0x8DE7
	QUAD_ALPHA4_SGIS                                           = 0x811E
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	SKIP_COMPONENTS3_NV                                        = -4
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	MATRIX2_ARB                                                = 0x88C2
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	COLOR_ATTACHMENT3                                          = 0x8CE3
	DST_ALPHA                                                  = 0x0304
	COMPRESSED_RED                                             = 0x8225
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	CURRENT_TANGENT_EXT                                        = 0x843B
	OUTPUT_COLOR0_EXT                                          = 0x879B
	PALETTE8_RGBA4_OES                                         = 0x8B98
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	DRAW_BUFFER                                                = 0x0C01
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	BUFFER_MAP_POINTER                                         = 0x88BD
	WRITE_ONLY_OES                                             = 0x88B9
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	LARGE_CW_ARC_TO_NV                                         = 0x18
	INTENSITY16I_EXT                                           = 0x8D8B
	INT_10_10_10_2_OES                                         = 0x8DF7
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	ACTIVE_PROGRAM                                             = 0x8259
	BITMAP_TOKEN                                               = 0x0704
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	REFERENCE_PLANE_SGIX                                       = 0x817D
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	OP_MIN_EXT                                                 = 0x878B
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	TEXTURE_MAG_FILTER                                         = 0x2800
	SAMPLE_MASK_SGIS                                           = 0x80A0
	MAX_HEIGHT                                                 = 0x827F
	SUBTRACT_ARB                                               = 0x84E7
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	SAMPLER_1D_SHADOW                                          = 0x8B61
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	LOGIC_OP                                                   = 0x0BF1
	TEXTURE_VIEW                                               = 0x82B5
	COMBINER6_NV                                               = 0x8556
	CONSTANT_ARB                                               = 0x8576
	DEPTH_COMPONENT32F                                         = 0x8CAC
	COLOR_ATTACHMENT8                                          = 0x8CE8
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	STENCIL_WRITEMASK                                          = 0x0B98
	RGB5_EXT                                                   = 0x8050
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	MAP_COLOR                                                  = 0x0D10
	COMBINE_ALPHA_ARB                                          = 0x8572
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	COLOR_ATTACHMENT7                                          = 0x8CE7
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	SAMPLE_POSITION_NV                                         = 0x8E50
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	CON_18_ATI                                                 = 0x8953
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	MODELVIEW28_ARB                                            = 0x873C
	COORD_REPLACE_NV                                           = 0x8862
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	CONTEXT_PROFILE_MASK                                       = 0x9126
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	BLEND_DST_RGB_OES                                          = 0x80C8
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	CON_29_ATI                                                 = 0x895E
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	QUERY_BUFFER_AMD                                           = 0x9192
	VIEWPORT_BIT                                               = 0x00000800
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	ALWAYS                                                     = 0x0207
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	FIXED_ONLY_ARB                                             = 0x891D
	SRGB_ALPHA                                                 = 0x8C42
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	EXT_texture_object                                         = 1
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	UNPACK_LSB_FIRST                                           = 0x0CF1
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	SYNC_STATUS_APPLE                                          = 0x9114
	PHONG_WIN                                                  = 0x80EA
	EIGHTH_BIT_ATI                                             = 0x00000020
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	LINES_ADJACENCY_ARB                                        = 0x000A
	TEXTURE_BORDER                                             = 0x1005
	TESS_CONTROL_TEXTURE                                       = 0x829C
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	FOG_COORD_ARRAY                                            = 0x8457
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	SAMPLES                                                    = 0x80A9
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	FRAGMENT_DEPTH                                             = 0x8452
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	MINMAX_FORMAT                                              = 0x802F
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	TRACE_MASK_MESA                                            = 0x8755
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	COLOR_ARRAY_LIST_IBM                                       = 103072
	DISPLAY_LIST                                               = 0x82E7
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	STANDARD_FONT_NAME_NV                                      = 0x9072
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	BOOL_VEC2_ARB                                              = 0x8B57
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	BEVEL_NV                                                   = 0x90A6
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	SGIX_instruments                                           = 1
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	AND_INVERTED                                               = 0x1504
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	READ_WRITE_ARB                                             = 0x88BA
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	MIPMAP                                                     = 0x8293
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	MAP2_TANGENT_EXT                                           = 0x8445
	TEXTURE8_ARB                                               = 0x84C8
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	DEPTH_STENCIL_EXT                                          = 0x84F9
	MODELVIEW10_ARB                                            = 0x872A
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	LINES                                                      = 0x0001
	RGBA16F_EXT                                                = 0x881A
	MATRIX_STRIDE                                              = 0x92FF
	RGBA12_EXT                                                 = 0x805A
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	MODELVIEW16_ARB                                            = 0x8730
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	HINT_BIT                                                   = 0x00008000
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	TEXTURE21                                                  = 0x84D5
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	DU8DV8_ATI                                                 = 0x877A
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	STACK_OVERFLOW                                             = 0x0503
	DRAW_BUFFER11_ATI                                          = 0x8830
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	BUFFER_ACCESS                                              = 0x88BB
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	POINT_BIT                                                  = 0x00000002
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	MAX_VIEWPORTS                                              = 0x825B
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	CUBIC_HP                                                   = 0x815F
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	NORMAL_ARRAY                                               = 0x8075
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	BLEND_COLOR                                                = 0x8005
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	BUFFER_SIZE                                                = 0x8764
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	MAGNITUDE_BIAS_NV                                          = 0x8718
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	MULTISAMPLE_BIT                                            = 0x20000000
	SPOT_EXPONENT                                              = 0x1205
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	TEXTURE_COMPRESSED                                         = 0x86A1
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	SGIS_fog_function                                          = 1
	PROGRAM_PARAMETER_NV                                       = 0x8644
	ABGR_EXT                                                   = 0x8000
	MODELVIEW4_ARB                                             = 0x8724
	MODELVIEW6_ARB                                             = 0x8726
	INT8_NV                                                    = 0x8FE0
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	TEXTURE11                                                  = 0x84CB
	BUFFER_ACCESS_OES                                          = 0x88BB
	STATIC_DRAW                                                = 0x88E4
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	FLOAT_VEC4                                                 = 0x8B52
	TRIANGULAR_NV                                              = 0x90A5
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	SGIX_resample                                              = 1
	EDGE_FLAG                                                  = 0x0B43
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	PROGRAM                                                    = 0x82E2
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	SYNC_FLAGS                                                 = 0x9115
	CLAMP_TO_BORDER_NV                                         = 0x812D
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	CON_22_ATI                                                 = 0x8957
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	RGBA32UI_EXT                                               = 0x8D70
	TEXTURE_BASE_LEVEL                                         = 0x813C
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	LINE_LOOP                                                  = 0x0002
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	CONDITION_SATISFIED                                        = 0x911C
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	FOG_OFFSET_SGIX                                            = 0x8198
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	SGIS_pixel_texture                                         = 1
	CONSTANT_COLOR_EXT                                         = 0x8001
	SRGB_DECODE_ARB                                            = 0x8299
	REG_0_ATI                                                  = 0x8921
	LOWER_LEFT                                                 = 0x8CA1
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	DRAW_BUFFER15_ATI                                          = 0x8834
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	RGB8I                                                      = 0x8D8F
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	FLOAT_VEC3_ARB                                             = 0x8B51
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	VIEW_CLASS_96_BITS                                         = 0x82C5
	FULL_RANGE_EXT                                             = 0x87E1
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	COMPILE                                                    = 0x1300
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	VARIABLE_F_NV                                              = 0x8528
	MAX_IMAGE_UNITS                                            = 0x8F38
	DEPTH_FUNC                                                 = 0x0B74
	RGBA16                                                     = 0x805B
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	REG_7_ATI                                                  = 0x8928
	LAYER_NV                                                   = 0x8DAA
	MINMAX_SINK                                                = 0x8030
	SPRITE_AXIS_SGIX                                           = 0x814A
	NORMAL_MAP_OES                                             = 0x8511
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	ACTIVE_RESOURCES                                           = 0x92F5
	RGBA2                                                      = 0x8055
	VERTEX_SOURCE_ATI                                          = 0x8774
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	SHADE_MODEL                                                = 0x0B54
	COLOR_ENCODING                                             = 0x8296
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	DRAW_BUFFER14_ATI                                          = 0x8833
	TIMESTAMP                                                  = 0x8E28
	QUADS                                                      = 0x0007
	TEXTURE_GEN_Q                                              = 0x0C63
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	MODELVIEW12_ARB                                            = 0x872C
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	RESCALE_NORMAL                                             = 0x803A
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	DOT3_RGB_EXT                                               = 0x8740
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	INT64_VEC4_NV                                              = 0x8FEB
	OPERAND2_ALPHA_ARB                                         = 0x859A
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	DSDT8_MAG8_NV                                              = 0x870A
	RGBA_FLOAT16_ATI                                           = 0x881A
	REG_11_ATI                                                 = 0x892C
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	EXT_polygon_offset                                         = 1
	POLYGON_TOKEN                                              = 0x0703
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	TEXTURE_GEN_T                                              = 0x0C61
	RGB_SCALE_ARB                                              = 0x8573
	INVERSE_NV                                                 = 0x862B
	FLOAT_MAT3_ARB                                             = 0x8B5B
	VERTEX_TEXTURE                                             = 0x829B
	PALETTE4_RGBA8_OES                                         = 0x8B91
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	BGRA_INTEGER                                               = 0x8D9B
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	PASS_THROUGH_TOKEN                                         = 0x0700
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	SOURCE0_ALPHA                                              = 0x8588
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	STENCIL_INDEX1_OES                                         = 0x8D46
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	MAX_LIGHTS                                                 = 0x0D31
	INT64_NV                                                   = 0x140E
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	PROGRAM_FORMAT_ARB                                         = 0x8876
	MAX_SAMPLES                                                = 0x8D57
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	BUMP_TARGET_ATI                                            = 0x877C
	MODULATE_COLOR_IMG                                         = 0x8C04
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	SGIS_point_line_texgen                                     = 1
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	IMAGE_BINDING_FORMAT                                       = 0x906E
	UNSIGNALED                                                 = 0x9118
	C4UB_V2F                                                   = 0x2A22
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	COMBINER4_NV                                               = 0x8554
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	MAX_CLIP_DISTANCES                                         = 0x0D32
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	YCRCB_444_SGIX                                             = 0x81BC
	FLOAT_R32_NV                                               = 0x8885
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	SCISSOR_BOX                                                = 0x0C10
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	RGBA32F_ARB                                                = 0x8814
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	COMBINE                                                    = 0x8570
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	SPOT_CUTOFF                                                = 0x1206
	DOUBLE_EXT                                                 = 0x140A
	LIGHT6                                                     = 0x4006
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	TEXTURE27                                                  = 0x84DB
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	GL_3D                                                      = 0x0601
	STENCIL_TEST                                               = 0x0B90
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	COPY                                                       = 0x1503
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	BLEND_SRC_RGB                                              = 0x80C9
	BLEND_SRC_ALPHA                                            = 0x80CB
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	SCALE_BY_FOUR_NV                                           = 0x853F
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	STENCIL_BACK_REF                                           = 0x8CA3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	HISTOGRAM_SINK_EXT                                         = 0x802D
	TEXTURE0_ARB                                               = 0x84C0
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	OPERAND0_ALPHA_ARB                                         = 0x8598
	MAX_VARYING_FLOATS                                         = 0x8B4B
	PATCH_VERTICES                                             = 0x8E72
	TEXTURE_LIGHT_EXT                                          = 0x8350
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	SKIP_COMPONENTS1_NV                                        = -6
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	OBJECT_LINE_SGIS                                           = 0x81F7
	YCBCR_422_APPLE                                            = 0x85B9
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	OPERAND0_RGB_ARB                                           = 0x8590
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	REG_5_ATI                                                  = 0x8926
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	ASYNC_MARKER_SGIX                                          = 0x8329
	NORMAL_MAP_NV                                              = 0x8511
	OPERAND1_ALPHA_ARB                                         = 0x8599
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	MAX_TEXTURE_SIZE                                           = 0x0D33
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	CLIP_DISTANCE5                                             = 0x3005
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	SOURCE2_RGB_ARB                                            = 0x8582
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	PATH_GEN_COEFF_NV                                          = 0x90B1
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	MAD_ATI                                                    = 0x8968
	INT_VEC4_ARB                                               = 0x8B55
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	CURRENT_NORMAL                                             = 0x0B02
	ADD_SIGNED_ARB                                             = 0x8574
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	MATRIX_MODE                                                = 0x0BA0
	GL_4PASS_3_SGIS                                            = 0x80A7
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	IGNORE_BORDER_HP                                           = 0x8150
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	AUX3                                                       = 0x040C
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	INDEX_ARRAY_STRIDE                                         = 0x8086
	MINMAX                                                     = 0x802E
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	PROGRAM_TARGET_NV                                          = 0x8646
	OP_MUL_EXT                                                 = 0x8786
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	DOUBLE_MAT3_EXT                                            = 0x8F47
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	CONSTANT_COLOR                                             = 0x8001
	NUM_EXTENSIONS                                             = 0x821D
	RG_EXT                                                     = 0x8227
	BLUE_INTEGER_EXT                                           = 0x8D96
	LINES_ADJACENCY_EXT                                        = 0x000A
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	BOOL_VEC3_ARB                                              = 0x8B58
	EXT_histogram                                              = 1
	STENCIL_INDEX                                              = 0x1901
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	COMPRESSED_ALPHA                                           = 0x84E9
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	PACK_SKIP_IMAGES                                           = 0x806B
	GENERATE_MIPMAP                                            = 0x8191
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	DOT3_RGBA_ARB                                              = 0x86AF
	COMPRESSED_RGB_ARB                                         = 0x84ED
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	RGBA16_EXT                                                 = 0x805B
	GL_4PASS_0_SGIS                                            = 0x80A4
	CURRENT_FOG_COORDINATE                                     = 0x8453
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	RGBA32UI                                                   = 0x8D70
	BOLD_BIT_NV                                                = 0x01
	INTENSITY12_EXT                                            = 0x804C
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	SLIM10U_SGIX                                               = 0x831E
	RGB_S3TC                                                   = 0x83A0
	SIGNED_NEGATE_NV                                           = 0x853D
	PROGRAM_OUTPUT                                             = 0x92E4
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	SOURCE0_RGB_ARB                                            = 0x8580
	DEBUG_PRINT_MESA                                           = 0x875A
	FRACTIONAL_ODD                                             = 0x8E7B
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	POLYGON_OFFSET_POINT                                       = 0x2A01
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	OP_EXP_BASE_2_EXT                                          = 0x8791
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	UNSIGNED_INT64_NV                                          = 0x140F
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	QUAD_MESH_SUN                                              = 0x8614
	BGR_INTEGER_EXT                                            = 0x8D9A
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	REG_6_ATI                                                  = 0x8927
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	REFLECTION_MAP_ARB                                         = 0x8512
	SET_AMD                                                    = 0x874A
	GREEN                                                      = 0x1904
	TEXTURE25                                                  = 0x84D9
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	FACTOR_MIN_AMD                                             = 0x901C
	IMAGE_3D_EXT                                               = 0x904E
	POINT                                                      = 0x1B00
	POINT_SIZE_MIN_EXT                                         = 0x8126
	SRC1_RGB                                                   = 0x8581
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	SRGB_ALPHA_EXT                                             = 0x8C42
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	DECR_WRAP_EXT                                              = 0x8508
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	STENCIL_INDEX4_OES                                         = 0x8D47
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	DRAW_BUFFER7_ATI                                           = 0x882C
	MATRIX16_ARB                                               = 0x88D0
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	CLIP_PLANE5                                                = 0x3005
	TEXTURE_WRAP_R_OES                                         = 0x8072
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	RG8I                                                       = 0x8237
	FOG_COORD                                                  = 0x8451
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	CON_28_ATI                                                 = 0x895D
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	DECR_WRAP                                                  = 0x8508
	SRGB8_NV                                                   = 0x8C41
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	LUMINANCE6_ALPHA2                                          = 0x8044
	CONVOLUTION_1D                                             = 0x8010
	TEXTURE_CUBE_MAP                                           = 0x8513
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	GL_8X_BIT_ATI                                              = 0x00000004
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	STREAM_DRAW_ARB                                            = 0x88E0
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	TEXTURE_GEN_S                                              = 0x0C60
	CLIP_DISTANCE2                                             = 0x3002
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	SIGNED_RGB_NV                                              = 0x86FE
	NEGATIVE_Z_EXT                                             = 0x87DB
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	VERTEX_SUBROUTINE                                          = 0x92E8
	COLOR_BUFFER_BIT                                           = 0x00004000
	RENDERER                                                   = 0x1F01
	CONSTANT_ALPHA_EXT                                         = 0x8003
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	REG_3_ATI                                                  = 0x8924
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	SUCCESS_NV                                                 = 0x902F
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	ALPHA_BITS                                                 = 0x0D55
	MAP2_VERTEX_3                                              = 0x0DB7
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	EMBOSS_LIGHT_NV                                            = 0x855D
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	MAX_LIST_NESTING                                           = 0x0B31
	ZOOM_Y                                                     = 0x0D17
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	TANGENT_ARRAY_EXT                                          = 0x8439
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	OUTPUT_COLOR1_EXT                                          = 0x879C
	VECTOR_EXT                                                 = 0x87BF
	INTERLACE_OML                                              = 0x8980
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	ALPHA16UI_EXT                                              = 0x8D78
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	SOURCE3_ALPHA_NV                                           = 0x858B
	DOT3_RGBA_IMG                                              = 0x86AF
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	ALPHA12_EXT                                                = 0x803D
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	WRAP_BORDER_SUN                                            = 0x81D4
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	MAX_CLIP_PLANES                                            = 0x0D32
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	SGIX_async                                                 = 1
	TEXTURE12_ARB                                              = 0x84CC
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	ALPHA_INTEGER                                              = 0x8D97
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	MAP1_INDEX                                                 = 0x0D91
	OR_REVERSE                                                 = 0x150B
	LINE                                                       = 0x1B01
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	AUX1                                                       = 0x040A
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	STENCIL_INDEX8_EXT                                         = 0x8D48
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	FIXED_ONLY                                                 = 0x891D
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	COUNTER_RANGE_AMD                                          = 0x8BC1
	RG16                                                       = 0x822C
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	INT_VEC2_ARB                                               = 0x8B53
	UPPER_LEFT                                                 = 0x8CA2
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	SGIX_fog_offset                                            = 1
	DRAW_BUFFER1_ARB                                           = 0x8826
	REG_27_ATI                                                 = 0x893C
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	VERTEX_STREAM7_ATI                                         = 0x8773
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	IUI_V2F_EXT                                                = 0x81AD
	PALETTE4_RGBA4_OES                                         = 0x8B93
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	COMPUTE_SHADER                                             = 0x91B9
	ZOOM_X                                                     = 0x0D16
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	MODULATE                                                   = 0x2100
	FUNC_ADD                                                   = 0x8006
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	INT_IMAGE_2D                                               = 0x9058
	RETURN                                                     = 0x0102
	VIEW_CLASS_8_BITS                                          = 0x82CB
	OP_FRAC_EXT                                                = 0x8789
	DRAW_BUFFER1_NV                                            = 0x8826
	MODELVIEW23_ARB                                            = 0x8737
	RG_SNORM                                                   = 0x8F91
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	COLOR_CLEAR_VALUE                                          = 0x0C22
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	COORD_REPLACE                                              = 0x8862
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	MIN_EXT                                                    = 0x8007
	IMAGE_SCALE_Y_HP                                           = 0x8156
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	COUNTER_TYPE_AMD                                           = 0x8BC0
	LINE_STIPPLE                                               = 0x0B24
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	PIXEL_COUNT_NV                                             = 0x8866
	STREAM_READ                                                = 0x88E1
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	DEBUG_SOURCE_API                                           = 0x8246
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	UNSIGNED_INVERT_NV                                         = 0x8537
	BUFFER_SIZE_ARB                                            = 0x8764
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	LUMINANCE16_SNORM                                          = 0x9019
	AUTO_NORMAL                                                = 0x0D80
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	SLUMINANCE_EXT                                             = 0x8C46
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	RGBA8I_EXT                                                 = 0x8D8E
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	DOUBLE_VEC2                                                = 0x8FFC
	RGB4_EXT                                                   = 0x804F
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	VERTEX_ARRAY_EXT                                           = 0x8074
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	TRANSLATE_3D_NV                                            = 0x9091
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	FLOAT_MAT3x2                                               = 0x8B67
	LINE_SMOOTH                                                = 0x0B20
	SRC1_ALPHA                                                 = 0x8589
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	MATRIX20_ARB                                               = 0x88D4
	EXT_texture3D                                              = 1
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	INT_VEC3_ARB                                               = 0x8B54
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	VARIABLE_A_NV                                              = 0x8523
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	REG_17_ATI                                                 = 0x8932
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	ALPHA_TEST                                                 = 0x0BC0
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	Y_EXT                                                      = 0x87D6
	TESS_GEN_POINT_MODE                                        = 0x8E79
	IMAGE_CUBE_EXT                                             = 0x9050
	MULTIVIEW_EXT                                              = 0x90F1
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	SIGNED_NORMALIZED                                          = 0x8F9C
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	COMBINE_ALPHA                                              = 0x8572
	REG_30_ATI                                                 = 0x893F
	INTERLACE_READ_OML                                         = 0x8981
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	SAMPLER_BUFFER                                             = 0x8DC2
	RG32UI                                                     = 0x823C
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	MODELVIEW3_ARB                                             = 0x8723
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	NAND                                                       = 0x150E
	NORMAL_MAP_EXT                                             = 0x8511
	INT_SAMPLER_3D                                             = 0x8DCB
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	LIST_MODE                                                  = 0x0B30
	INVERT                                                     = 0x150A
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DEBUG_TYPE_MARKER                                          = 0x8268
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	COMBINER2_NV                                               = 0x8552
	STORAGE_CACHED_APPLE                                       = 0x85BE
	OP_SUB_EXT                                                 = 0x8796
	FLOAT_MAT4x2                                               = 0x8B69
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	RED_MAX_CLAMP_INGR                                         = 0x8564
	DOT3_RGBA_EXT                                              = 0x8741
	FLOAT_MAT4x3                                               = 0x8B6A
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	DEPTH_SCALE                                                = 0x0D1E
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	OPERAND2_ALPHA_EXT                                         = 0x859A
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	FOG_HINT                                                   = 0x0C54
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	SPOT_DIRECTION                                             = 0x1204
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	UNIFORM_BLOCK                                              = 0x92E2
	ALPHA_FLOAT32_ATI                                          = 0x8816
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	SIGNED_IDENTITY_NV                                         = 0x853C
	STENCIL_INDEX4                                             = 0x8D47
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	FRONT_RIGHT                                                = 0x0401
	EYE_LINEAR                                                 = 0x2400
	CMYKA_EXT                                                  = 0x800D
	CLAMP_TO_EDGE                                              = 0x812F
	CONSTANT_BORDER                                            = 0x8151
	RGBA_DXT5_S3TC                                             = 0x83A4
	COMPRESSED_INTENSITY                                       = 0x84EC
	DSDT_MAG_VIB_NV                                            = 0x86F7
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	SAMPLER_3D                                                 = 0x8B5F
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	TEXTURE_SHADER_NV                                          = 0x86DE
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	FLOAT16_VEC4_NV                                            = 0x8FFB
	COUNT_UP_NV                                                = 0x9088
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	CURRENT_BINORMAL_EXT                                       = 0x843C
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	COLOR_ATTACHMENT12                                         = 0x8CEC
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	INT16_NV                                                   = 0x8FE4
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	LIGHT7                                                     = 0x4007
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	CAVEAT_SUPPORT                                             = 0x82B8
	ATTENUATION_EXT                                            = 0x834D
	TEXTURE27_ARB                                              = 0x84DB
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	SGI_color_table                                            = 1
	INCR                                                       = 0x1E02
	MODELVIEW2_ARB                                             = 0x8722
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	BLUE_SCALE                                                 = 0x0D1A
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	BIAS_BIT_ATI                                               = 0x00000008
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	INTENSITY_SNORM                                            = 0x9013
	SGIS_generate_mipmap                                       = 1
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	IMAGE_2D_EXT                                               = 0x904D
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	PRIMARY_COLOR_EXT                                          = 0x8577
	DEPTH_STENCIL_MESA                                         = 0x8750
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	DEPTH_BUFFER_BIT                                           = 0x00000100
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	COLOR_ARRAY_TYPE                                           = 0x8082
	PACK_IMAGE_HEIGHT                                          = 0x806C
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	TEXTURE_GATHER                                             = 0x82A2
	DRAW_BUFFER14_ARB                                          = 0x8833
	OP_INDEX_EXT                                               = 0x8782
	MATRIX13_ARB                                               = 0x88CD
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	SLUMINANCE8                                                = 0x8C47
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	DRAW_BUFFER0_NV                                            = 0x8825
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	STENCIL_INDEX1_EXT                                         = 0x8D46
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	GL_422_AVERAGE_EXT                                         = 0x80CE
	QUERY_RESULT_EXT                                           = 0x8866
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	GL_422_EXT                                                 = 0x80CC
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	CURRENT_MATRIX_NV                                          = 0x8641
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	RGBA16F                                                    = 0x881A
	NORMALIZE                                                  = 0x0BA1
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	LUMINANCE8UI_EXT                                           = 0x8D80
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	MODELVIEW26_ARB                                            = 0x873A
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	PACK_ROW_LENGTH                                            = 0x0D02
	RGB12                                                      = 0x8053
	ALPHA4_EXT                                                 = 0x803B
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MIRROR_CLAMP_ATI                                           = 0x8742
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	SGI_color_matrix                                           = 1
	UNSIGNED_BYTE                                              = 0x1401
	FUNC_ADD_OES                                               = 0x8006
	INTENSITY_EXT                                              = 0x8049
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	DRAW_BUFFER12                                              = 0x8831
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	RGBA4                                                      = 0x8056
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	ALPHA_MAX_SGIX                                             = 0x8321
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	DOUBLE_MAT3x4                                              = 0x8F4C
	INDEX_MODE                                                 = 0x0C30
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	TEXTURE0                                                   = 0x84C0
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	FILE_NAME_NV                                               = 0x9074
	BLOCK_INDEX                                                = 0x92FD
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	OP_RECIP_EXT                                               = 0x8794
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	TEXTURE3                                                   = 0x84C3
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	MULTISAMPLE_SGIS                                           = 0x809D
	TEXTURE17                                                  = 0x84D1
	SIGNED_RGBA8_NV                                            = 0x86FC
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	SOURCE0_RGB                                                = 0x8580
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	OP_FLOOR_EXT                                               = 0x878F
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	MATRIX5_ARB                                                = 0x88C5
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	OPERAND1_ALPHA                                             = 0x8599
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	COLOR_ATTACHMENT5                                          = 0x8CE5
	STENCIL_INDEX4_EXT                                         = 0x8D47
	R16F_EXT                                                   = 0x822D
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	MATRIX4_NV                                                 = 0x8634
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	RGB565_OES                                                 = 0x8D62
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	DRAW_BUFFER8_ATI                                           = 0x882D
	REG_24_ATI                                                 = 0x8939
	EXT_blend_subtract                                         = 1
	TEXTURE_3D                                                 = 0x806F
	RGB4_S3TC                                                  = 0x83A1
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	LEFT                                                       = 0x0406
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	MAX_TEXTURE_UNITS                                          = 0x84E2
	TEXTURE19_ARB                                              = 0x84D3
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	SHADER_BINARY_VIV                                          = 0x8FC4
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	CON_2_ATI                                                  = 0x8943
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	STREAM_COPY                                                = 0x88E2
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	RENDERBUFFER_BINDING                                       = 0x8CA7
	PRESENT_DURATION_NV                                        = 0x8E2B
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	LINE_STRIP_ADJACENCY                                       = 0x000B
	INDEX_LOGIC_OP                                             = 0x0BF1
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	BUFFER_MAPPED                                              = 0x88BC
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	MAP2_VERTEX_4                                              = 0x0DB8
	HI_BIAS_NV                                                 = 0x8714
	CURRENT_QUERY_ARB                                          = 0x8865
	SATURATE_BIT_ATI                                           = 0x00000040
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	TEXTURE26_ARB                                              = 0x84DA
	REG_1_ATI                                                  = 0x8922
	SRGB                                                       = 0x8C40
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	PACK_INVERT_MESA                                           = 0x8758
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	MITER_TRUNCATE_NV                                          = 0x90A8
	MAX_LAYERS                                                 = 0x8281
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	VARIANT_VALUE_EXT                                          = 0x87E4
	VERTEX_SHADER_ARB                                          = 0x8B31
	STENCIL_INDEX8_OES                                         = 0x8D48
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	R3_G3_B2                                                   = 0x2A10
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	BGR_EXT                                                    = 0x80E0
	POINT_SIZE_MAX                                             = 0x8127
	DRAW_BUFFER10_NV                                           = 0x882F
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	CULL_FACE_MODE                                             = 0x0B45
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	FLOAT_RG_NV                                                = 0x8881
	BOOL_VEC4                                                  = 0x8B59
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	ALPHA32UI_EXT                                              = 0x8D72
	INTENSITY8                                                 = 0x804B
	TEXTURE_MIN_LOD                                            = 0x813A
	STENCIL_COMPONENTS                                         = 0x8285
	AVERAGE_HP                                                 = 0x8160
	SOURCE1_ALPHA                                              = 0x8589
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	SELECT                                                     = 0x1C02
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	HISTOGRAM_WIDTH                                            = 0x8026
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	AVERAGE_EXT                                                = 0x8335
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	DRAW_BUFFER12_NV                                           = 0x8831
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	VERSION_1_1                                                = 1
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	SIGNED_INTENSITY_NV                                        = 0x8707
	STATIC_ATI                                                 = 0x8760
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	DRAW_BUFFER2_ATI                                           = 0x8827
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	PALETTE8_RGBA8_OES                                         = 0x8B96
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	POLYGON_OFFSET_LINE                                        = 0x2A02
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	TEXTURE_RECTANGLE                                          = 0x84F5
	INT_IMAGE_3D                                               = 0x9059
	OBJECT_TYPE_APPLE                                          = 0x9112
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	GL_3DC_X_AMD                                               = 0x87F9
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	ARRAY_BUFFER_ARB                                           = 0x8892
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	DEPTH_COMPONENT24                                          = 0x81A6
	POLYGON_STIPPLE                                            = 0x0B42
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	YCRCBA_SGIX                                                = 0x8319
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	SURFACE_STATE_NV                                           = 0x86EB
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	NEXT_BUFFER_NV                                             = -2
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	RENDERBUFFER_WIDTH                                         = 0x8D42
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	VERTEX_ID_NV                                               = 0x8C7B
	T                                                          = 0x2001
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	FRACTIONAL_EVEN                                            = 0x8E7C
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	SOURCE1_RGB_EXT                                            = 0x8581
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	AUX_BUFFERS                                                = 0x0C00
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	OPERAND0_ALPHA                                             = 0x8598
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	RG_INTEGER                                                 = 0x8228
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	DYNAMIC_COPY_ARB                                           = 0x88EA
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	RG8UI                                                      = 0x8238
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	CON_6_ATI                                                  = 0x8947
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	SAMPLER_3D_ARB                                             = 0x8B5F
	SAMPLER_3D_OES                                             = 0x8B5F
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	CURRENT_TIME_NV                                            = 0x8E28
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	PALETTE4_RGB8_OES                                          = 0x8B90
	DOUBLE_MAT2x4                                              = 0x8F4A
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	PACK_SWAP_BYTES                                            = 0x0D00
	TEXTURE3_ARB                                               = 0x84C3
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	YCBCR_MESA                                                 = 0x8757
	TEXTURE_COMPARE_MODE                                       = 0x884C
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	IMAGE_1D                                                   = 0x904C
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	ADD                                                        = 0x0104
	COMBINE_ARB                                                = 0x8570
	MATRIX2_NV                                                 = 0x8632
	FEEDBACK                                                   = 0x1C01
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	TEXTURE10                                                  = 0x84CA
	DRAW_BUFFER10_ARB                                          = 0x882F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	VIDEO_BUFFER_NV                                            = 0x9020
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	BLUE_BITS                                                  = 0x0D54
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	PACK_SKIP_ROWS                                             = 0x0D03
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	RG8                                                        = 0x822B
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	COLOR_MATERIAL                                             = 0x0B57
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	SAMPLES_ARB                                                = 0x80A9
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	DECODE_EXT                                                 = 0x8A49
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	LUMINANCE16                                                = 0x8042
	T2F_IUI_V3F_EXT                                            = 0x81B2
	MATRIX28_ARB                                               = 0x88DC
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	MEDIUM_FLOAT                                               = 0x8DF1
	DOUBLE_MAT4x2                                              = 0x8F4D
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	TEXTURE16                                                  = 0x84D0
	FLOAT_VEC2                                                 = 0x8B50
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	INTENSITY32I_EXT                                           = 0x8D85
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	PROXY_TEXTURE_1D                                           = 0x8063
	LUMINANCE12                                                = 0x8041
	SURFACE_MAPPED_NV                                          = 0x8700
	CON_10_ATI                                                 = 0x894B
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	STACK_UNDERFLOW                                            = 0x0504
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	R32UI                                                      = 0x8236
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	OPERAND3_RGB_NV                                            = 0x8593
	CON_1_ATI                                                  = 0x8942
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	DYNAMIC_READ_ARB                                           = 0x88E9
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	PROGRAM_SEPARABLE                                          = 0x8258
	SOURCE2_ALPHA                                              = 0x858A
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	LINES_ADJACENCY                                            = 0x000A
	ALPHA8_EXT                                                 = 0x803C
	PROVOKING_VERTEX                                           = 0x8E4F
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	TEXTURE6                                                   = 0x84C6
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	IMAGE_1D_EXT                                               = 0x904C
	UTF8_NV                                                    = 0x909A
	INVALID_ENUM                                               = 0x0500
	VIEW_CLASS_48_BITS                                         = 0x82C7
	TEXTURE19                                                  = 0x84D3
	MATRIX5_NV                                                 = 0x8635
	SLUMINANCE8_EXT                                            = 0x8C47
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	COPY_WRITE_BUFFER                                          = 0x8F37
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	CONSTANT                                                   = 0x8576
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	BACK_LEFT                                                  = 0x0402
	CURRENT_FOG_COORD                                          = 0x8453
	LUMINANCE8                                                 = 0x8040
	SAMPLES_EXT                                                = 0x80A9
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	DRAW_BUFFER11                                              = 0x8830
	BLEND_EQUATION_ALPHA                                       = 0x883D
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	MATRIX10_ARB                                               = 0x88CA
	RELATIVE_ARC_TO_NV                                         = 0xFF
	PIXEL_MODE_BIT                                             = 0x00000020
	N3F_V3F                                                    = 0x2A25
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	VERTEX_ARRAY_TYPE                                          = 0x807B
	TEXTURE9                                                   = 0x84C9
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	READ_BUFFER_EXT                                            = 0x0C02
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	PROXY_TEXTURE_2D                                           = 0x8064
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	FLOAT_RG32_NV                                              = 0x8887
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	PATH_FILL_MASK_NV                                          = 0x9081
	EXT_vertex_array                                           = 1
	RGBA12                                                     = 0x805A
	MULTISAMPLE_EXT                                            = 0x809D
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	HALF_FLOAT_OES                                             = 0x8D61
	INT_IMAGE_2D_EXT                                           = 0x9058
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	RED_BITS                                                   = 0x0D52
	SMOOTH                                                     = 0x1D01
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	VARIABLE_G_NV                                              = 0x8529
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	CONSTANT_BORDER_HP                                         = 0x8151
	R16UI                                                      = 0x8234
	DOT_PRODUCT_NV                                             = 0x86EC
	RGB8_SNORM                                                 = 0x8F96
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	RGBA_SNORM                                                 = 0x8F93
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	SRGB8_ALPHA8                                               = 0x8C43
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	DEPTH_RANGE                                                = 0x0B70
	CLIP_PLANE4                                                = 0x3004
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	PROGRAM_LENGTH_ARB                                         = 0x8627
	INTENSITY32UI_EXT                                          = 0x8D73
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	CURRENT_INDEX                                              = 0x0B01
	YCRCB_422_SGIX                                             = 0x81BB
	MODELVIEW22_ARB                                            = 0x8736
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	CONTINUOUS_AMD                                             = 0x9007
	DEPTH_TEST                                                 = 0x0B71
	DOUBLE                                                     = 0x140A
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	PACK_CMYK_HINT_EXT                                         = 0x800E
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	CON_16_ATI                                                 = 0x8951
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	EXT_texture                                                = 1
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	SEPARABLE_2D                                               = 0x8012
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	CONVOLUTION_1D_EXT                                         = 0x8010
	PREVIOUS_EXT                                               = 0x8578
	MATRIX27_ARB                                               = 0x88DB
	UNDEFINED_APPLE                                            = 0x8A1C
	TEXTURE_RED_SIZE                                           = 0x805C
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	TEXTURE25_ARB                                              = 0x84D9
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	INT8_VEC2_NV                                               = 0x8FE1
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	RGB2_EXT                                                   = 0x804E
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	CON_25_ATI                                                 = 0x895A
	RGB_422_APPLE                                              = 0x8A1F
	ACTIVE_VARIABLES                                           = 0x9305
	CURRENT_MATRIX_ARB                                         = 0x8641
	DRAW_BUFFER14                                              = 0x8833
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	SAMPLES_PASSED                                             = 0x8914
	PALETTE8_RGB8_OES                                          = 0x8B95
	SRGB8                                                      = 0x8C41
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	TRACK_MATRIX_NV                                            = 0x8648
	DRAW_BUFFER1_ATI                                           = 0x8826
	CLIP_DISTANCE_NV                                           = 0x8C7A
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	RGB32UI                                                    = 0x8D71
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	DRAW_BUFFER4                                               = 0x8829
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	LINE_RESET_TOKEN                                           = 0x0707
	XOR                                                        = 0x1506
	LUMINANCE16_EXT                                            = 0x8042
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	VARIANT_ARRAY_EXT                                          = 0x87E8
	SAMPLER_BINDING                                            = 0x8919
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	COLOR_SAMPLES_NV                                           = 0x8E20
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	INVARIANT_EXT                                              = 0x87C2
	COLOR_ATTACHMENT6                                          = 0x8CE6
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	DRAW_BUFFER4_NV                                            = 0x8829
	INT_IMAGE_1D_EXT                                           = 0x9057
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	TRUE                                                       = 1
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	OR_INVERTED                                                = 0x150D
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	MATRIX3_ARB                                                = 0x88C3
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	INTENSITY8_EXT                                             = 0x804B
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	IS_ROW_MAJOR                                               = 0x9300
	DEPTH_COMPONENT                                            = 0x1902
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	SIGNED_HILO8_NV                                            = 0x885F
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	DOUBLE_MAT2_EXT                                            = 0x8F46
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	OP_MAX_EXT                                                 = 0x878A
	DRAW_BUFFER4_ARB                                           = 0x8829
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	TRANSLATE_2D_NV                                            = 0x9090
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	CLAMP_VERTEX_COLOR                                         = 0x891A
	CONVEX_HULL_NV                                             = 0x908B
	FLOAT_MAT4                                                 = 0x8B5C
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	COLOR_ARRAY_EXT                                            = 0x8076
	FRAGMENT_TEXTURE                                           = 0x829F
	MAP2_BINORMAL_EXT                                          = 0x8447
	COMPRESSED_RGB                                             = 0x84ED
	DOUBLE_MAT4_EXT                                            = 0x8F48
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	SLIM8U_SGIX                                                = 0x831D
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	BLEND_DST_RGB                                              = 0x80C8
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	FLOAT_RGBA16_NV                                            = 0x888A
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	RGB_FLOAT32_ATI                                            = 0x8815
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	COMPILE_STATUS                                             = 0x8B81
	SLUMINANCE8_NV                                             = 0x8C47
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	TEXTURE_MIN_FILTER                                         = 0x2801
	HISTOGRAM_FORMAT                                           = 0x8027
	SHININESS                                                  = 0x1601
	EYE_PLANE                                                  = 0x2502
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	DRAW_PIXEL_TOKEN                                           = 0x0705
	INDEX_CLEAR_VALUE                                          = 0x0C20
	POSITION                                                   = 0x1203
	RGB10                                                      = 0x8052
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	SGIX_sprite                                                = 1
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	CONSTANT_COLOR1_NV                                         = 0x852B
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	INT64_VEC3_NV                                              = 0x8FEA
	V2F                                                        = 0x2A20
	HISTOGRAM                                                  = 0x8024
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	REG_29_ATI                                                 = 0x893E
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	LIGHT5                                                     = 0x4005
	UNIFORM_SIZE                                               = 0x8A38
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	TEXTURE_HEIGHT                                             = 0x1001
	MIN                                                        = 0x8007
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	POLYGON                                                    = 0x0009
	REPLACE                                                    = 0x1E01
	CLIP_DISTANCE7                                             = 0x3007
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	R8UI                                                       = 0x8232
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	IMAGE_2D_ARRAY                                             = 0x9053
	COLOR_TABLE_FORMAT                                         = 0x80D8
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	BGR                                                        = 0x80E0
	PACK_RESAMPLE_OML                                          = 0x8984
	TESS_GEN_SPACING                                           = 0x8E77
	INT8_VEC4_NV                                               = 0x8FE3
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	HISTOGRAM_RED_SIZE                                         = 0x8028
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	FLOAT16_VEC2_NV                                            = 0x8FF9
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	SGIS_texture_filter4                                       = 1
	R1UI_V3F_SUN                                               = 0x85C4
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	OP_POWER_EXT                                               = 0x8793
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	MAP_TESSELLATION_NV                                        = 0x86C2
	CURRENT_VERTEX_EXT                                         = 0x87E2
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	DEBUG_SOURCE_OTHER                                         = 0x824B
	OPERAND1_RGB                                               = 0x8591
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	DEPTH_BITS                                                 = 0x0D56
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	MATRIX26_ARB                                               = 0x88DA
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	SLIM12S_SGIX                                               = 0x831F
	MODELVIEW11_ARB                                            = 0x872B
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	MULT                                                       = 0x0103
	SPRITE_MODE_SGIX                                           = 0x8149
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	INT16_VEC3_NV                                              = 0x8FE6
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	SGIX_texture_scale_bias                                    = 1
	PROJECTION                                                 = 0x1701
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	PROXY_HISTOGRAM                                            = 0x8025
	RG16F_EXT                                                  = 0x822F
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	UTF16_NV                                                   = 0x909B
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	SCISSOR_BIT                                                = 0x00080000
	TEXTURE_PRIORITY                                           = 0x8066
	FIXED                                                      = 0x140C
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	SIGNED_RGB8_NV                                             = 0x86FF
	INVALID_INDEX                                              = 0xFFFFFFFF
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	RGBA8                                                      = 0x8058
	VERTEX_SHADER_EXT                                          = 0x8780
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	FIXED_OES                                                  = 0x140C
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	TEXTURE26                                                  = 0x84DA
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	RGBA16F_ARB                                                = 0x881A
	WRITE_ONLY                                                 = 0x88B9
	BUFFER_ACCESS_ARB                                          = 0x88BB
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	STORAGE_SHARED_APPLE                                       = 0x85BF
	SIGNED_LUMINANCE_NV                                        = 0x8701
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	REPEAT                                                     = 0x2901
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	RGB565                                                     = 0x8D62
	TEXTURE_COMPONENTS                                         = 0x1003
	PROGRAM_BINDING_ARB                                        = 0x8677
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	LINE_TOKEN                                                 = 0x0702
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	CURRENT_ATTRIB_NV                                          = 0x8626
	REG_19_ATI                                                 = 0x8934
	ALPHA4                                                     = 0x803B
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	TEXTURE_BUFFER                                             = 0x8C2A
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	SHADER_COMPILER                                            = 0x8DFA
	INTENSITY8_SNORM                                           = 0x9017
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	SGIS_detail_texture                                        = 1
	INTERLACE_SGIX                                             = 0x8094
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	ATC_RGB_AMD                                                = 0x8C92
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	FIELD_UPPER_NV                                             = 0x9022
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	SGIX_texture_coordinate_clamp                              = 1
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	BUFFER_MAPPED_ARB                                          = 0x88BC
	MATRIX17_ARB                                               = 0x88D1
	UNIFORM_BUFFER                                             = 0x8A11
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	SRC2_ALPHA                                                 = 0x858A
	SIGNED_HILO_NV                                             = 0x86F9
	DYNAMIC_DRAW                                               = 0x88E8
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	SRC2_RGB                                                   = 0x8582
	PATH_STENCIL_REF_NV                                        = 0x90B8
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	UNSIGNED_SHORT                                             = 0x1403
	EQUIV                                                      = 0x1509
	DUAL_ALPHA12_SGIS                                          = 0x8112
	TEXTURE29                                                  = 0x84DD
	DRAW_BUFFER10                                              = 0x882F
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	AFFINE_3D_NV                                               = 0x9094
	VARIABLE_B_NV                                              = 0x8524
	FRAGMENT_SHADER                                            = 0x8B30
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	RGB9_E5_EXT                                                = 0x8C3D
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	CUBIC_CURVE_TO_NV                                          = 0x0C
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	REPLICATE_BORDER_HP                                        = 0x8153
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	FOG_START                                                  = 0x0B63
	MATRIX19_ARB                                               = 0x88D3
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	VERTEX_SHADER_BIT                                          = 0x00000001
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	TEXTURE2_ARB                                               = 0x84C2
	TEXTURE9_ARB                                               = 0x84C9
	HALF_APPLE                                                 = 0x140B
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	DRAW_BUFFER5_ATI                                           = 0x882A
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	LUMINANCE8_SNORM                                           = 0x9015
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	SYNC_CONDITION_APPLE                                       = 0x9113
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	SOURCE1_RGB                                                = 0x8581
	MODELVIEW24_ARB                                            = 0x8738
	MAX_IMAGE_SAMPLES                                          = 0x906D
	UNSIGNED_INT                                               = 0x1405
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	RGB                                                        = 0x1907
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	GREEN_INTEGER_EXT                                          = 0x8D95
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	GL_2PASS_0_SGIS                                            = 0x80A2
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	INTERPOLATE_EXT                                            = 0x8575
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	BGRA                                                       = 0x80E1
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	ALPHA32F_ARB                                               = 0x8816
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	UNIFORM_OFFSET                                             = 0x8A3B
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	SGIX_convolution_accuracy                                  = 1
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	HISTOGRAM_EXT                                              = 0x8024
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	VARIABLE_D_NV                                              = 0x8526
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	COLOR_INDEX1_EXT                                           = 0x80E2
	FOG_COORD_SRC                                              = 0x8450
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	COLOR_MATERIAL_FACE                                        = 0x0B55
	FOG_DENSITY                                                = 0x0B62
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	SAMPLES_PASSED_ARB                                         = 0x8914
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	CON_11_ATI                                                 = 0x894C
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	VERTEX_ARRAY_SIZE                                          = 0x807A
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	RGBA16_SNORM                                               = 0x8F9B
	MAX_DRAW_BUFFERS                                           = 0x8824
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	EXT_copy_texture                                           = 1
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	MINOR_VERSION                                              = 0x821C
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	CON_26_ATI                                                 = 0x895B
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	ACCUM_BUFFER_BIT                                           = 0x00000200
	COPY_PIXEL_TOKEN                                           = 0x0706
	MAJOR_VERSION                                              = 0x821B
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	DRAW_BUFFER7_NV                                            = 0x882C
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	CLIP_PLANE2                                                = 0x3002
	FUNC_ADD_EXT                                               = 0x8006
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	DYNAMIC_ATI                                                = 0x8761
	CON_8_ATI                                                  = 0x8949
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	COLOR_ATTACHMENT13                                         = 0x8CED
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	SGIS_texture_edge_clamp                                    = 1
	IUI_N3F_V3F_EXT                                            = 0x81B0
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	PROGRAM_RESIDENT_NV                                        = 0x8647
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	FLOAT_RGBA32_NV                                            = 0x888B
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	VERSION_1_2                                                = 1
	IMAGE_SCALE_X_HP                                           = 0x8155
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	MINMAX_EXT                                                 = 0x802E
	R8I                                                        = 0x8231
	TEXTURE10_ARB                                              = 0x84CA
	TEXTURE18_ARB                                              = 0x84D2
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	RG8_SNORM                                                  = 0x8F95
	VIEW_CLASS_16_BITS                                         = 0x82CA
	Z400_BINARY_AMD                                            = 0x8740
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	RGBA4_EXT                                                  = 0x8056
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	CURRENT_COLOR                                              = 0x0B00
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	RG16UI                                                     = 0x823A
	SOURCE2_ALPHA_ARB                                          = 0x858A
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	LOCATION_INDEX                                             = 0x930F
	INDEX_ARRAY_LIST_IBM                                       = 103073
	NICEST                                                     = 0x1102
	CONVOLUTION_2D                                             = 0x8011
	MAP1_TANGENT_EXT                                           = 0x8444
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	DISCARD_ATI                                                = 0x8763
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	SGIX_scalebias_hint                                        = 1
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	LUMINANCE12_EXT                                            = 0x8041
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	INDEX_ARRAY_EXT                                            = 0x8077
	MULTISAMPLE_ARB                                            = 0x809D
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	TEXTURE_SAMPLES                                            = 0x9106
	FASTEST                                                    = 0x1101
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	TESS_EVALUATION_SHADER                                     = 0x8E87
	RG16I                                                      = 0x8239
	DEPTH_CLAMP_NV                                             = 0x864F
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	FLOAT_VEC3                                                 = 0x8B51
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	LO_SCALE_NV                                                = 0x870F
	FLOAT_MAT2x4                                               = 0x8B66
	FLOAT_MAT3x4                                               = 0x8B68
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	RGB16I                                                     = 0x8D89
	MAP2_INDEX                                                 = 0x0DB1
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	DEPTH_STENCIL_OES                                          = 0x84F9
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	ELEMENT_ARRAY_ATI                                          = 0x8768
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	COLOR_ARRAY_STRIDE                                         = 0x8083
	LUMINANCE16_ALPHA16                                        = 0x8048
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	SHADER_OBJECT_EXT                                          = 0x8B48
	CULL_VERTEX_IBM                                            = 103050
	LIST_INDEX                                                 = 0x0B33
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	MODELVIEW31_ARB                                            = 0x873F
	OP_ADD_EXT                                                 = 0x8787
	DRAW_BUFFER2_ARB                                           = 0x8827
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	SYNC_OBJECT_APPLE                                          = 0x8A53
	ACCUM_BLUE_BITS                                            = 0x0D5A
	DEBUG_TYPE_OTHER                                           = 0x8251
	OP_SET_LT_EXT                                              = 0x878D
	SGIX_shadow                                                = 1
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	CLIP_PLANE3                                                = 0x3003
	MIRRORED_REPEAT_IBM                                        = 0x8370
	FOG_COORDINATE_EXT                                         = 0x8451
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	SGIX_pixel_texture                                         = 1
	TEXTURE30_ARB                                              = 0x84DE
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	R8_SNORM                                                   = 0x8F94
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	SPRITE_SGIX                                                = 0x8148
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	EXT_packed_pixels                                          = 1
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	DRAW_BUFFER5_ARB                                           = 0x882A
	CURRENT_QUERY_EXT                                          = 0x8865
	MOV_ATI                                                    = 0x8961
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	BUFFER                                                     = 0x82E0
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	INCR_WRAP                                                  = 0x8507
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	STENCIL_INDEX16_EXT                                        = 0x8D49
	VERTEX23_BIT_PGI                                           = 0x00000004
	BLEND_SRC                                                  = 0x0BE1
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	VOLATILE_APPLE                                             = 0x8A1A
	DELETE_STATUS                                              = 0x8B80
	LINK_STATUS                                                = 0x8B82
	NO_ERROR                                                   = 0
	PURGEABLE_APPLE                                            = 0x8A1D
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	RED                                                        = 0x1903
	RESCALE_NORMAL_EXT                                         = 0x803A
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	SOURCE2_RGB                                                = 0x8582
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	RG8_EXT                                                    = 0x822B
	STATIC_COPY_ARB                                            = 0x88E6
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	ONE_MINUS_DST_COLOR                                        = 0x0307
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	VIEW_CLASS_64_BITS                                         = 0x82C6
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	SAMPLER_2D                                                 = 0x8B5E
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	FALSE                                                      = 0
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	MODELVIEW20_ARB                                            = 0x8734
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	HILO8_NV                                                   = 0x885E
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	COLOR_ATTACHMENT10                                         = 0x8CEA
	QUERY_NO_WAIT_NV                                           = 0x8E14
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	DUAL_ALPHA4_SGIS                                           = 0x8110
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	DRAW_BUFFER10_ATI                                          = 0x882F
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	PATH_FORMAT_PS_NV                                          = 0x9071
	READ_BUFFER                                                = 0x0C02
	CLAMP                                                      = 0x2900
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	FRAME_NV                                                   = 0x8E26
	PATH_MITER_LIMIT_NV                                        = 0x907A
	TRANSLATE_X_NV                                             = 0x908E
	EXT_blend_color                                            = 1
	SGIX_icc_texture                                           = 1
	TRANSFORM_BIT                                              = 0x00001000
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	OCCLUSION_TEST_HP                                          = 0x8165
	TEXTURE20                                                  = 0x84D4
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	TRACE_NAME_MESA                                            = 0x8756
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	VERTEX_STREAM0_ATI                                         = 0x876C
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	RELATIVE_MOVE_TO_NV                                        = 0x03
	LIGHT1                                                     = 0x4001
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	DEBUG_SEVERITY_LOW                                         = 0x9148
	MAP1_VERTEX_4                                              = 0x0D98
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	QUERY_OBJECT_EXT                                           = 0x9153
	UNDEFINED_VERTEX                                           = 0x8260
	PRIMARY_COLOR_ARB                                          = 0x8577
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	ARRAY_STRIDE                                               = 0x92FE
	TEXTURE_MAX_LEVEL                                          = 0x813D
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	COMBINER5_NV                                               = 0x8555
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	FAILURE_NV                                                 = 0x9030
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	CON_27_ATI                                                 = 0x895C
	ALPHA8UI_EXT                                               = 0x8D7E
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	DRAW_BUFFER6                                               = 0x882B
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	REGISTER_COMBINERS_NV                                      = 0x8522
	SAMPLER_CUBE                                               = 0x8B60
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	LAYOUT_LINEAR_INTEL                                        = 1
	DOMAIN                                                     = 0x0A02
	MAP_STENCIL                                                = 0x0D11
	DEPTH_COMPONENT32                                          = 0x81A7
	TRIANGLE_LIST_SUN                                          = 0x81D7
	TEXTURE23                                                  = 0x84D7
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	W_EXT                                                      = 0x87D8
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	ETC1_SRGB8_NV                                              = 0x88EE
	TEXTURE_1D_ARRAY                                           = 0x8C18
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	FIELD_LOWER_NV                                             = 0x9023
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	INT64_VEC2_NV                                              = 0x8FE9
	UNIFORM_TYPE                                               = 0x8A37
	VERTEX_SHADER                                              = 0x8B31
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	TEXTURE8                                                   = 0x84C8
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	SAMPLE_BUFFERS                                             = 0x80A8
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	ACCUM_RED_BITS                                             = 0x0D58
	PROGRAM_PIPELINE                                           = 0x82E4
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	SHADER_BINARY_DMP                                          = 0x9250
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	OBJECT_LINEAR                                              = 0x2401
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	BINORMAL_ARRAY_EXT                                         = 0x843A
	COPY_INVERTED                                              = 0x150C
	CON_5_ATI                                                  = 0x8946
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	SGIX_reference_plane                                       = 1
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	Q                                                          = 0x2003
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	QUERY_NO_WAIT                                              = 0x8E14
	BUFFER_OBJECT_EXT                                          = 0x9151
	SGIX_texture_add_env                                       = 1
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	POINT_SMOOTH                                               = 0x0B10
	TEXTURE_WRAP_T                                             = 0x2803
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	MATRIX14_ARB                                               = 0x88CE
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	SOURCE2_RGB_EXT                                            = 0x8582
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	PROXY_TEXTURE_3D                                           = 0x8070
	COMBINER_MUX_SUM_NV                                        = 0x8547
	COLOR_ATTACHMENT0                                          = 0x8CE0
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	NONE                                                       = 0
	FOG_INDEX                                                  = 0x0B61
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	LIGHT0                                                     = 0x4000
	RED_MIN_CLAMP_INGR                                         = 0x8560
	ALPHA_FLOAT16_ATI                                          = 0x881C
	IMAGE_3D                                                   = 0x904E
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	READ_WRITE                                                 = 0x88BA
	STATIC_DRAW_ARB                                            = 0x88E4
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	R32I                                                       = 0x8235
	RGB_SCALE_EXT                                              = 0x8573
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	ALPHA8_SNORM                                               = 0x9014
	T2F_N3F_V3F                                                = 0x2A2B
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	NAME_LENGTH                                                = 0x92F9
	EDGE_FLAG_ARRAY                                            = 0x8079
	GL_2_BYTES                                                 = 0x1407
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	DRAW_BUFFER13_ARB                                          = 0x8832
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	COLOR_RENDERABLE                                           = 0x8286
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	GENERATE_MIPMAP_HINT                                       = 0x8192
	VIEW_CLASS_24_BITS                                         = 0x82C9
	OPERAND2_RGB                                               = 0x8592
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	INT_VEC3                                                   = 0x8B54
	RASTERIZER_DISCARD                                         = 0x8C89
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	EVAL_BIT                                                   = 0x00010000
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	LOCAL_EXT                                                  = 0x87C4
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	DRAW_BUFFER0                                               = 0x8825
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	DONT_CARE                                                  = 0x1100
	RGB_SCALE                                                  = 0x8573
	INTERPOLATE                                                = 0x8575
	DOT4_ATI                                                   = 0x8967
	FRAMEBUFFER_EXT                                            = 0x8D40
	INTENSITY16UI_EXT                                          = 0x8D79
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	TEXTURE_BIT                                                = 0x00040000
	FILTER                                                     = 0x829A
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	RGBA_FLOAT32_APPLE                                         = 0x8814
	UNSIGNALED_APPLE                                           = 0x9118
	TRIANGLE_STRIP                                             = 0x0005
	C4F_N3F_V3F                                                = 0x2A26
	DRAW_BUFFER3_NV                                            = 0x8828
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	READ_BUFFER_NV                                             = 0x0C02
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	GL_4_BYTES                                                 = 0x1409
	VARIABLE_E_NV                                              = 0x8527
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	CON_17_ATI                                                 = 0x8952
	CND_ATI                                                    = 0x896A
	SAMPLER_1D_ARB                                             = 0x8B5D
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	RED_SCALE                                                  = 0x0D14
	INTENSITY12                                                = 0x804C
	PHONG_HINT_WIN                                             = 0x80EB
	RGBA4_S3TC                                                 = 0x83A3
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	RGBA8_OES                                                  = 0x8058
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	PACK_RESAMPLE_SGIX                                         = 0x842C
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	REG_13_ATI                                                 = 0x892E
	RGBA16I_EXT                                                = 0x8D88
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	INTENSITY32F_ARB                                           = 0x8817
	DRAW_BUFFER9                                               = 0x882E
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	RGBA8_SNORM                                                = 0x8F97
	INT_IMAGE_BUFFER                                           = 0x905C
	INDEX_SHIFT                                                = 0x0D12
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	INT16_VEC2_NV                                              = 0x8FE5
	SLICE_ACCUM_SUN                                            = 0x85CC
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	LERP_ATI                                                   = 0x8969
	TEXTURE_ENV_COLOR                                          = 0x2201
	DEBUG_TYPE_ERROR                                           = 0x824C
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	TEXTURE_LOD_BIAS                                           = 0x8501
	DECR_WRAP_OES                                              = 0x8508
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	MODELVIEW7_ARB                                             = 0x8727
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	REG_16_ATI                                                 = 0x8931
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	RG                                                         = 0x8227
	SOURCE2_ALPHA_EXT                                          = 0x858A
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	SAMPLER_2D_ARB                                             = 0x8B5E
	LUMINANCE_SNORM                                            = 0x9011
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	COLOR_SUM_EXT                                              = 0x8458
	FLOAT_VEC4_ARB                                             = 0x8B52
	GPU_ADDRESS_NV                                             = 0x8F34
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	COMPRESSED_R11_EAC                                         = 0x9270
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	TEXTURE23_ARB                                              = 0x84D7
	REG_25_ATI                                                 = 0x893A
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	MAX_EXT                                                    = 0x8008
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	POINT_SIZE_MIN_ARB                                         = 0x8126
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	SET                                                        = 0x150F
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	INT_SAMPLER_2D                                             = 0x8DCA
	SQUARE_NV                                                  = 0x90A3
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	CULL_FRAGMENT_NV                                           = 0x86E7
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	CLAMP_READ_COLOR                                           = 0x891C
	CONSTANT_ATTENUATION                                       = 0x1207
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	RGBA8UI                                                    = 0x8D7C
	INT_SAMPLER_1D                                             = 0x8DC9
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	SGIX_texture_multi_buffer                                  = 1
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	INTENSITY16                                                = 0x804D
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	TEXTURE17_ARB                                              = 0x84D1
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	HI_SCALE_NV                                                = 0x870E
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	OFFSET                                                     = 0x92FC
	INDEX_BIT_PGI                                              = 0x00080000
	LINE_WIDTH_RANGE                                           = 0x0B22
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	TEXTURE6_ARB                                               = 0x84C6
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	COUNT_DOWN_NV                                              = 0x9089
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	PROGRAM_POINT_SIZE                                         = 0x8642
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	RGB32F                                                     = 0x8815
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	VERSION_1_4                                                = 1
	EXTENSIONS                                                 = 0x1F03
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	OPERAND3_ALPHA_NV                                          = 0x859B
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	SHADER_IMAGE_LOAD                                          = 0x82A4
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	RGB16F_EXT                                                 = 0x881B
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	PRIMITIVE_RESTART                                          = 0x8F9D
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	COMBINER7_NV                                               = 0x8557
	TRANSPOSE_NV                                               = 0x862C
	SAMPLER_2D_RECT                                            = 0x8B63
	RGB_INTEGER                                                = 0x8D98
	ACCUM                                                      = 0x0100
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	GL_1PASS_SGIS                                              = 0x80A1
	FENCE_STATUS_NV                                            = 0x84F3
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	RELATIVE_LINE_TO_NV                                        = 0x05
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	VERTEX4_BIT_PGI                                            = 0x00000008
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	LESS                                                       = 0x0201
	COLOR_ARRAY                                                = 0x8076
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	MAX_EVAL_ORDER                                             = 0x0D30
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	CONTEXT_FLAGS                                              = 0x821E
	TEXTURE15                                                  = 0x84CF
	DOT3_RGBA                                                  = 0x86AF
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	BLEND_DST                                                  = 0x0BE0
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	RGBA4_DXT5_S3TC                                            = 0x83A5
	MATRIX12_ARB                                               = 0x88CC
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	PROJECTION_MATRIX                                          = 0x0BA7
	COLOR                                                      = 0x1800
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	LIGHTING_BIT                                               = 0x00000040
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	COEFF                                                      = 0x0A00
	BLUE_BIAS                                                  = 0x0D1B
	TEXTURE30                                                  = 0x84DE
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	TIME_ELAPSED_EXT                                           = 0x88BF
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	CURRENT_RASTER_INDEX                                       = 0x0B05
	MODELVIEW                                                  = 0x1700
	GL_2X_BIT_ATI                                              = 0x00000001
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	FACTOR_MAX_AMD                                             = 0x901D
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	X_EXT                                                      = 0x87D5
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	CON_20_ATI                                                 = 0x8955
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	TEXTURE_MATRIX                                             = 0x0BA8
	VIEW_CLASS_128_BITS                                        = 0x82C4
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	IMAGE_BUFFER                                               = 0x9051
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	COLOR_ATTACHMENT9                                          = 0x8CE9
	MOVE_TO_RESETS_NV                                          = 0x90B5
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	EXT_shared_texture_palette                                 = 1
	DST_COLOR                                                  = 0x0306
	DEPTH_BIAS                                                 = 0x0D1F
	READ_PIXELS_FORMAT                                         = 0x828D
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	VARIANT_EXT                                                = 0x87C1
	REG_26_ATI                                                 = 0x893B
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	STENCIL_FUNC                                               = 0x0B92
	ATTACHED_SHADERS                                           = 0x8B85
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	INDEX_OFFSET                                               = 0x0D13
	DRAW_BUFFER5                                               = 0x882A
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	SGIX_async_pixel                                           = 1
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	RGBA16UI                                                   = 0x8D76
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	MATRIX_PALETTE_OES                                         = 0x8840
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	FRONT                                                      = 0x0404
	INT_VEC2                                                   = 0x8B53
	FLOAT16_NV                                                 = 0x8FF8
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	REG_2_ATI                                                  = 0x8923
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	SHORT                                                      = 0x1402
	OP_MADD_EXT                                                = 0x8788
	SAMPLE_SHADING_ARB                                         = 0x8C36
	READ_FRAMEBUFFER                                           = 0x8CA8
	RGBA32I_EXT                                                = 0x8D82
	MAX_VERTEX_STREAMS                                         = 0x8E71
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	INT_IMAGE_3D_EXT                                           = 0x9059
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	LIGHT2                                                     = 0x4002
	GL_4PASS_2_SGIS                                            = 0x80A6
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	SGIX_list_priority                                         = 1
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	TRIANGLE_MESH_SUN                                          = 0x8615
	GL_4X_BIT_ATI                                              = 0x00000002
	FRONT_AND_BACK                                             = 0x0408
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	DEPTH_RENDERABLE                                           = 0x8287
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	TEXTURE31_ARB                                              = 0x84DF
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	TEXTURE_DEPTH                                              = 0x8071
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	EXT_blend_logic_op                                         = 1
	ACCUM_GREEN_BITS                                           = 0x0D59
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	CLEAR_BUFFER                                               = 0x82B4
	STATIC_READ_ARB                                            = 0x88E5
	CLOSE_PATH_NV                                              = 0x00
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	SIGNED_INTENSITY8_NV                                       = 0x8708
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	BITMAP                                                     = 0x1A00
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	MIRRORED_REPEAT                                            = 0x8370
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	SGIS_point_parameters                                      = 1
	COLOR_ATTACHMENT11                                         = 0x8CEB
	HIGH_FLOAT                                                 = 0x8DF2
	SAMPLER_BUFFER_AMD                                         = 0x9001
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	MATRIX18_ARB                                               = 0x88D2
	REG_28_ATI                                                 = 0x893D
	CON_21_ATI                                                 = 0x8956
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	BUFFER_MAP_OFFSET                                          = 0x9121
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	FLOAT_MAT3                                                 = 0x8B5B
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	SYNC_CONDITION                                             = 0x9113
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	REDUCE_EXT                                                 = 0x8016
	POINT_SIZE_MAX_EXT                                         = 0x8127
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	PIXEL_PACK_BUFFER                                          = 0x88EB
	ALPHA8I_EXT                                                = 0x8D90
	ROUND_NV                                                   = 0x90A4
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	COMBINER_MAPPING_NV                                        = 0x8543
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	COMPILE_AND_EXECUTE                                        = 0x1301
	RGB16F                                                     = 0x881B
	ALPHA_BIAS                                                 = 0x0D1D
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	TEXTURE31                                                  = 0x84DF
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	FLOAT_RGB32_NV                                             = 0x8889
	CON_14_ATI                                                 = 0x894F
	MUL_ATI                                                    = 0x8964
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	UNPACK_ALIGNMENT                                           = 0x0CF5
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	RGB16_EXT                                                  = 0x8054
	COMPRESSED_RG                                              = 0x8226
	NORMAL_MAP_ARB                                             = 0x8511
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	REDUCE                                                     = 0x8016
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	CLIP_DISTANCE3                                             = 0x3003
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	COMMAND_BARRIER_BIT                                        = 0x00000040
	TEXTURE_2D                                                 = 0x0DE1
	CON_31_ATI                                                 = 0x8960
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	MINMAX_FORMAT_EXT                                          = 0x802F
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	INTENSITY16_SNORM                                          = 0x901B
	LAYOUT_DEFAULT_INTEL                                       = 0
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	RGBA16UI_EXT                                               = 0x8D76
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	SGI_texture_color_table                                    = 1
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	MAP2_COLOR_4                                               = 0x0DB0
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	COPY_READ_BUFFER                                           = 0x8F36
	GL_422_REV_EXT                                             = 0x80CD
	DEPTH_CLAMP                                                = 0x864F
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	RGB16F_ARB                                                 = 0x881B
	CON_0_ATI                                                  = 0x8941
	UNIFORM_BUFFER_START                                       = 0x8A29
	SLUMINANCE_ALPHA                                           = 0x8C44
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	COLOR_SUM_ARB                                              = 0x8458
	REFLECTION_MAP                                             = 0x8512
	EXPAND_NEGATE_NV                                           = 0x8539
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	CLIP_DISTANCE1                                             = 0x3001
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	SKIP_DECODE_EXT                                            = 0x8A4A
	STENCIL_ATTACHMENT                                         = 0x8D20
	CUBIC_EXT                                                  = 0x8334
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	OP_DOT4_EXT                                                = 0x8785
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	RENDER_MODE                                                = 0x0C40
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	INVALID_VALUE                                              = 0x0501
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	DRAW_BUFFER6_ARB                                           = 0x882B
	RGBA                                                       = 0x1908
	RGB5_A1_EXT                                                = 0x8057
	GL_2PASS_0_EXT                                             = 0x80A2
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	BUFFER_USAGE_ARB                                           = 0x8765
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	LOW_FLOAT                                                  = 0x8DF0
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	COMBINER_SCALE_NV                                          = 0x8548
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	POINT_SMOOTH_HINT                                          = 0x0C51
	HALF_FLOAT_NV                                              = 0x140B
	VERSION                                                    = 0x1F02
	RGB10_EXT                                                  = 0x8052
	FLOAT_MAT2                                                 = 0x8B5A
	BACK_RIGHT                                                 = 0x0403
	COLOR_TABLE_SCALE                                          = 0x80D6
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	DEPTH_STENCIL                                              = 0x84F9
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	TRANSFORM_FEEDBACK                                         = 0x8E22
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	COMBINE_RGB_ARB                                            = 0x8571
	SOURCE1_ALPHA_ARB                                          = 0x8589
	POINT_SPRITE_ARB                                           = 0x8861
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	FOG_FUNC_SGIS                                              = 0x812A
	CONSTANT_EXT                                               = 0x8576
	EXT_rescale_normal                                         = 1
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	RG32I                                                      = 0x823B
	REFLECTION_MAP_NV                                          = 0x8512
	REG_14_ATI                                                 = 0x892F
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	MAX_SAMPLES_NV                                             = 0x8D57
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	PROGRAM_INPUT                                              = 0x92E3
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	CLIP_DISTANCE6                                             = 0x3006
	GL_2PASS_1_SGIS                                            = 0x80A3
	DOUBLE_VEC4                                                = 0x8FFE
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	INCR_WRAP_OES                                              = 0x8507
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	REG_8_ATI                                                  = 0x8929
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	SYNC_STATUS                                                = 0x9114
	SGIX_subsample                                             = 1
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	RGBA_INTEGER                                               = 0x8D99
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	RGB10_A2UI                                                 = 0x906F
	DOT3_RGB_ARB                                               = 0x86AE
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	DEBUG_ASSERT_MESA                                          = 0x875B
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	DOUBLE_MAT3x2                                              = 0x8F4B
	SIGNALED_APPLE                                             = 0x9119
	ALL_STATIC_DATA_IBM                                        = 103060
	VIBRANCE_BIAS_NV                                           = 0x8719
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	READ_PIXELS                                                = 0x828C
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	DRAW_BUFFER11_NV                                           = 0x8830
	STATIC_READ                                                = 0x88E5
	HIGH_INT                                                   = 0x8DF5
	LUMINANCE12_ALPHA12                                        = 0x8047
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	OPERAND2_ALPHA                                             = 0x859A
	REG_23_ATI                                                 = 0x8938
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	SGIX_framezoom                                             = 1
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	ALL_COMPLETED_NV                                           = 0x84F2
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	CON_19_ATI                                                 = 0x8954
	SRGB_EXT                                                   = 0x8C40
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	GEQUAL                                                     = 0x0206
	CURRENT_RASTER_COLOR                                       = 0x0B04
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	INTENSITY8I_EXT                                            = 0x8D91
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	COLOR_INDEX8_EXT                                           = 0x80E5
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	BLUE_INTEGER                                               = 0x8D96
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	DITHER                                                     = 0x0BD0
	GL_4PASS_3_EXT                                             = 0x80A7
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	SOURCE1_ALPHA_EXT                                          = 0x8589
	MATRIX0_NV                                                 = 0x8630
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	GL_2D                                                      = 0x0600
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	FRAGMENT_COLOR_EXT                                         = 0x834C
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	OUT_OF_MEMORY                                              = 0x0505
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	T4F_V4F                                                    = 0x2A28
	HILO16_NV                                                  = 0x86F8
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	DEPTH_STENCIL_NV                                           = 0x84F9
	TEXTURE_WIDTH                                              = 0x1000
	UNPACK_SKIP_IMAGES                                         = 0x806D
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	TEXTURE15_ARB                                              = 0x84CF
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	SUBPIXEL_BITS                                              = 0x0D50
	MATRIX30_ARB                                               = 0x88DE
	RED_INTEGER_EXT                                            = 0x8D94
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	COMBINER1_NV                                               = 0x8551
	MAX_DEPTH                                                  = 0x8280
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	GL_1PASS_EXT                                               = 0x80A1
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	BOOL_VEC2                                                  = 0x8B57
	PATH_COORD_COUNT_NV                                        = 0x909E
	DEPTH                                                      = 0x1801
	REPLACE_EXT                                                = 0x8062
	YCRCB_SGIX                                                 = 0x8318
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	UNSIGNED_INT_24_8                                          = 0x84FA
	LO_BIAS_NV                                                 = 0x8715
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	BLEND_EQUATION_RGB                                         = 0x8009
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	CULL_MODES_NV                                              = 0x86E0
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	RGB5                                                       = 0x8050
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	POLYGON_OFFSET_EXT                                         = 0x8037
	COLOR3_BIT_PGI                                             = 0x00010000
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	SOURCE1_RGB_ARB                                            = 0x8581
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	PRIMITIVES_GENERATED                                       = 0x8C87
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	PASS_THROUGH_NV                                            = 0x86E6
	GREEN_BIT_ATI                                              = 0x00000002
	IMAGE_BUFFER_EXT                                           = 0x9051
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	HILO_NV                                                    = 0x86F4
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	SGX_BINARY_IMG                                             = 0x8C0A
	GEOMETRY_SHADER                                            = 0x8DD9
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	EQUAL                                                      = 0x0202
	OP_CLAMP_EXT                                               = 0x878E
	COMP_BIT_ATI                                               = 0x00000002
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	IMAGE_1D_ARRAY                                             = 0x9052
	EXT_abgr                                                   = 1
	SGIX_ycrcb                                                 = 1
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	CONVOLUTION_2D_EXT                                         = 0x8011
	PREVIOUS_ARB                                               = 0x8578
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	STREAM_DRAW                                                = 0x88E0
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	SGIS_multisample                                           = 1
	COLOR_ARRAY_SIZE                                           = 0x8081
	CMYK_EXT                                                   = 0x800C
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	ACTIVE_UNIFORMS                                            = 0x8B86
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	INTENSITY                                                  = 0x8049
	QUAD_ALPHA8_SGIS                                           = 0x811F
	TEXTURE_4D_SGIS                                            = 0x8134
	TEXTURE1                                                   = 0x84C1
	RETAINED_APPLE                                             = 0x8A1B
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	SGIX_flush_raster                                          = 1
	RGB16                                                      = 0x8054
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	OP_DOT3_EXT                                                = 0x8784
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	SGIX_ir_instrument1                                        = 1
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	SAMPLES_3DFX                                               = 0x86B4
	COORD_REPLACE_ARB                                          = 0x8862
	YCBYCR8_422_NV                                             = 0x9031
	FLOAT                                                      = 0x1406
	CONVOLUTION_HEIGHT                                         = 0x8019
	R8                                                         = 0x8229
	TEXTURE7                                                   = 0x84C7
	POINT_TOKEN                                                = 0x0701
	FOG_MODE                                                   = 0x0B65
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	RGB_FLOAT16_ATI                                            = 0x881B
	RGB_INTEGER_EXT                                            = 0x8D98
	TRIANGLES_ADJACENCY                                        = 0x000C
	PERTURB_EXT                                                = 0x85AE
	DSDT_NV                                                    = 0x86F5
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	FLOAT16_VEC3_NV                                            = 0x8FFA
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	BUFFER_BINDING                                             = 0x9302
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	ALPHA8                                                     = 0x803C
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	DOUBLE_MAT4x3                                              = 0x8F4E
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	TEXTURE_COORD_ARRAY                                        = 0x8078
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	BUMP_ENVMAP_ATI                                            = 0x877B
	REG_31_ATI                                                 = 0x8940
	ETC1_RGB8_OES                                              = 0x8D64
	RGBA8I                                                     = 0x8D8E
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	REFLECTION_MAP_EXT                                         = 0x8512
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	RED_SNORM                                                  = 0x8F90
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	POINT_SPRITE                                               = 0x8861
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	UNIFORM                                                    = 0x92E1
	FOG_BIT                                                    = 0x00000080
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	DOT2_ADD_ATI                                               = 0x896C
	SYNC_FENCE                                                 = 0x9116
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	MATRIX1_NV                                                 = 0x8631
	FILL                                                       = 0x1B02
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	NORMAL_BIT_PGI                                             = 0x08000000
	LINE_SMOOTH_HINT                                           = 0x0C52
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	RGB8I_EXT                                                  = 0x8D8F
	QUERY                                                      = 0x82E3
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	DRAW_BUFFER9_ARB                                           = 0x882E
	SLUMINANCE_NV                                              = 0x8C46
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	SYNC_CL_EVENT_ARB                                          = 0x8240
	SPRITE_AXIAL_SGIX                                          = 0x814C
	FOG_COORDINATE_SOURCE                                      = 0x8450
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	COLOR_ATTACHMENT2                                          = 0x8CE2
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	IUI_N3F_V2F_EXT                                            = 0x81AF
	VIEW_CLASS_32_BITS                                         = 0x82C8
	SRC0_RGB                                                   = 0x8580
	CON_13_ATI                                                 = 0x894E
	RGBA16I                                                    = 0x8D88
	SGIX_polynomial_ffd                                        = 1
	STREAM_COPY_ARB                                            = 0x88E2
	NORMAL_ARRAY_POINTER                                       = 0x808F
	NAME_STACK_DEPTH                                           = 0x0D70
	FILTER4_SGIS                                               = 0x8146
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	LINE_WIDTH                                                 = 0x0B21
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	BLEND_EQUATION_OES                                         = 0x8009
	COMBINER_INPUT_NV                                          = 0x8542
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	KEEP                                                       = 0x1E00
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	MODELVIEW25_ARB                                            = 0x8739
	DRAW_BUFFER14_NV                                           = 0x8833
	BUFFER_MAPPED_OES                                          = 0x88BC
	RGB8UI_EXT                                                 = 0x8D7D
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	LUMINANCE_ALPHA                                            = 0x190A
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	ADD_ATI                                                    = 0x8963
	RESTART_PATH_NV                                            = 0xF0
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	VERSION_1_5                                                = 1
	POLYGON_BIT                                                = 0x00000008
	COLOR_TABLE_BIAS                                           = 0x80D7
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	INT_SAMPLER_CUBE                                           = 0x8DCC
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	SINGLE_COLOR_EXT                                           = 0x81F9
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	SGIX_async_histogram                                       = 1
	NEGATIVE_ONE_EXT                                           = 0x87DF
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	PERCENTAGE_AMD                                             = 0x8BC3
	IMAGE_2D_RECT_EXT                                          = 0x904F
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	INT_IMAGE_CUBE                                             = 0x905B
	COMBINE_RGB_EXT                                            = 0x8571
	MATRIX29_ARB                                               = 0x88DD
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	ALPHA16                                                    = 0x803E
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	TEXTURE7_ARB                                               = 0x84C7
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	TEXTURE14                                                  = 0x84CE
	MODELVIEW0_EXT                                             = 0x1700
	FRAGMENT_SHADER_ATI                                        = 0x8920
	TEXTURE_GREEN_SIZE                                         = 0x805D
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	DRAW_BUFFER11_ARB                                          = 0x8830
	ARRAY_BUFFER_BINDING                                       = 0x8894
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	PACK_LSB_FIRST                                             = 0x0D01
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	DUAL_ALPHA8_SGIS                                           = 0x8111
	TEXTURE20_ARB                                              = 0x84D4
	MODELVIEW13_ARB                                            = 0x872D
	VIBRANCE_SCALE_NV                                          = 0x8713
	IMAGE_CUBE                                                 = 0x9050
	EXP                                                        = 0x0800
	AND                                                        = 0x1501
	STENCIL_REF                                                = 0x0B97
	COLOR_MATRIX_SGI                                           = 0x80B1
	MATRIX0_ARB                                                = 0x88C0
	REG_20_ATI                                                 = 0x8935
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	RGB16I_EXT                                                 = 0x8D89
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	REPLICATE_BORDER                                           = 0x8153
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	SYNC_FLAGS_APPLE                                           = 0x9115
	SRC_ALPHA                                                  = 0x0302
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	OUTPUT_FOG_EXT                                             = 0x87BD
	CON_30_ATI                                                 = 0x895F
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	COLOR_MATRIX                                               = 0x80B1
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	PATH_GEN_MODE_NV                                           = 0x90B0
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	INDEX_TEST_REF_EXT                                         = 0x81B7
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	SIGNED_ALPHA8_NV                                           = 0x8706
	MAGNITUDE_SCALE_NV                                         = 0x8712
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	RGB16_SNORM                                                = 0x8F9A
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	EXT_blend_minmax                                           = 1
	TEXTURE_1D                                                 = 0x0DE0
	MULTISAMPLE_3DFX                                           = 0x86B2
	UNSIGNED_INT8_NV                                           = 0x8FEC
	QUAD_STRIP                                                 = 0x0008
	IDENTITY_NV                                                = 0x862A
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	MIRROR_CLAMP_EXT                                           = 0x8742
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	REG_9_ATI                                                  = 0x892A
	GREEN_BIAS                                                 = 0x0D19
	ACTIVE_TEXTURE                                             = 0x84E0
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	DRAW_BUFFER6_NV                                            = 0x882B
	FLOAT_RGB_NV                                               = 0x8882
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	COLOR_INDEX4_EXT                                           = 0x80E4
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	INTERLACE_READ_INGR                                        = 0x8568
	COMBINE_RGB                                                = 0x8571
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	BLUE                                                       = 0x1905
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	Z_EXT                                                      = 0x87D7
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	COLOR_INDEX2_EXT                                           = 0x80E3
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	RGB32UI_EXT                                                = 0x8D71
	SHADER_IMAGE_STORE                                         = 0x82A5
	INFO_LOG_LENGTH                                            = 0x8B84
	ALPHA16I_EXT                                               = 0x8D8A
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	LUMINANCE4_EXT                                             = 0x803F
	COLOR_TABLE_WIDTH                                          = 0x80D9
	TEXTURE_RED_TYPE                                           = 0x8C10
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	MAP1_COLOR_4                                               = 0x0D90
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	DRAW_BUFFER13_ATI                                          = 0x8832
	RGB_SNORM                                                  = 0x8F92
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	BLEND_DST_RGB_EXT                                          = 0x80C8
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	SAMPLER_1D                                                 = 0x8B5D
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	GREEN_INTEGER                                              = 0x8D95
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	WAIT_FAILED                                                = 0x911D
	OPERAND0_RGB_EXT                                           = 0x8590
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	INTENSITY16F_ARB                                           = 0x881D
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	CLIP_PLANE0                                                = 0x3000
	FUNC_SUBTRACT                                              = 0x800A
	FLOAT_RGB16_NV                                             = 0x8888
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	COLOR_TABLE_SGI                                            = 0x80D0
	BGRA_EXT                                                   = 0x80E1
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	RG32F                                                      = 0x8230
	FRAMEBUFFER_BLEND                                          = 0x828B
	VERSION_3_1                                                = 1
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	DRAW_BUFFER7                                               = 0x882C
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	T2F_C4UB_V3F                                               = 0x2A29
	STENCIL_RENDERABLE                                         = 0x8288
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	NOTEQUAL                                                   = 0x0205
	LOGIC_OP_MODE                                              = 0x0BF0
	DYNAMIC_COPY                                               = 0x88EA
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	PROGRAM_STRING_ARB                                         = 0x8628
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	INT_2_10_10_10_REV                                         = 0x8D9F
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	OP_ROUND_EXT                                               = 0x8790
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	TEXTURE24_ARB                                              = 0x84D8
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	SGIX_blend_alpha_minmax                                    = 1
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	DEPTH_COMPONENTS                                           = 0x8284
	SHADER_TYPE                                                = 0x8B4F
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	MOVE_TO_NV                                                 = 0x02
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	RIGHT                                                      = 0x0407
	FULL_SUPPORT                                               = 0x82B7
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	SUB_ATI                                                    = 0x8965
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	FRAMEZOOM_SGIX                                             = 0x818B
	VERTEX_STREAM2_ATI                                         = 0x876E
	DRAW_BUFFER0_ARB                                           = 0x8825
	MATRIX9_ARB                                                = 0x88C9
	TABLE_TOO_LARGE                                            = 0x8031
	COMBINE_ALPHA_EXT                                          = 0x8572
	MATRIX7_NV                                                 = 0x8637
	RED_BIT_ATI                                                = 0x00000001
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	TEXTURE_BINDING_1D                                         = 0x8068
	COLOR_INDEX                                                = 0x1900
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	DSDT8_NV                                                   = 0x8709
	DRAW_BUFFER15                                              = 0x8834
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	SAMPLE_MASK                                                = 0x8E51
	VERTICAL_LINE_TO_NV                                        = 0x08
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	STEREO                                                     = 0x0C33
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	RGBA2_EXT                                                  = 0x8055
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	PATH_FILL_MODE_NV                                          = 0x9080
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	MAX_PATCH_VERTICES                                         = 0x8E7D
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
)

type Context struct {
	context                 *C.gl15Context
	Accum                   func(op uint32, value float32)
	AlphaFunc               func(Func uint32, ref float32)
	Begin                   func(mode uint32)
	End                     func()
	Bitmap                  func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap []uint8)
	BlendFunc               func(sfactor, dfactor uint32)
	CallList                func(list uint32)
	CallLists               func(n int32, Type uint32, lists unsafe.Pointer)
	Clear                   func(mask uint32)
	ClearAccum              func(red, green, blue, alpha float32)
	ClearColor              func(red, green, blue, alpha float32)
	ClearDepth              func(depth float64)
	ClearIndex              func(c float32)
	ClearStencil            func(s int32)
	ClipPlane               func(plane uint32, equation []float64)
	Color3b                 func(red, green, blue int8)
	Color3d                 func(red, green, blue float64)
	Color3f                 func(red, green, blue float32)
	Color3i                 func(red, green, blue int32)
	Color3s                 func(red, green, blue int16)
	Color3ub                func(red, green, blue uint8)
	Color3ui                func(red, green, blue uint32)
	Color3us                func(red, green, blue uint16)
	Color4b                 func(red, green, blue, alpha int8)
	Color4d                 func(red, green, blue, alpha float64)
	Color4f                 func(red, green, blue, alpha float32)
	Color4i                 func(red, green, blue, alpha int32)
	Color4s                 func(red, green, blue, alpha int16)
	Color4ub                func(red, green, blue, alpha uint8)
	Color4ui                func(red, green, blue, alpha uint32)
	Color4us                func(red, green, blue, alpha uint16)
	Color3bv                func(v []int8)
	Color3dv                func(v []float64)
	Color3fv                func(v []float32)
	Color3iv                func(v []int32)
	Color3sv                func(v []int16)
	Color3ubv               func(v []uint8)
	Color3uiv               func(v []uint32)
	Color3usv               func(v []uint16)
	Color4bv                func(v []int8)
	Color4dv                func(v []float64)
	Color4fv                func(v []float32)
	Color4iv                func(v []int32)
	Color4sv                func(v []int16)
	Color4ubv               func(v []uint8)
	Color4uiv               func(v []uint32)
	Color4usv               func(v []uint16)
	ColorMask               func(red, green, blue, alpha bool)
	ColorMaterial           func(face, mode uint32)
	ColorTable              func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer)
	ColorTableParameterfv   func(target, pname uint32, params []float32)
	ColorTableParameteriv   func(target, pname uint32, params []int32)
	ColorSubTable           func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer)
	CopyPixels              func(x, y int32, width, height int32, Type uint32)
	CullFace                func(mode uint32)
	ConvolutionFilter1D     func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionFilter2D     func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer)
	ConvolutionParameterf   func(target, pname uint32, params float32)
	ConvolutionParameteri   func(target, pname uint32, params int32)
	CopyColorTable          func(target, internalformat uint32, x, y int32, width int32)
	CopyColorSubTable       func(target uint32, start int32, x, y int32, width int32)
	CopyConvolutionFilter1D func(target, internalformat uint32, x, y int32, width int32)
	CopyConvolutionFilter2D func(target, internalformat uint32, x, y int32, width, height int32)
	DeleteLists             func(list uint32, Range int32)
	DepthFunc               func(Func uint32)
	DepthRange              func(zNear, zFar float64)
	Enable                  func(cap uint32)
	Disable                 func(cap uint32)
	DrawBuffer              func(mode uint32)
	DrawPixels              func(width, height int32, format, Type uint32, pixels unsafe.Pointer)
	EdgeFlag                func(flag bool)
	EdgeFlagv               func(flag []bool)
	EdgeFlagPointer         func(stride int32, pointer unsafe.Pointer)
	EnableClientState       func(cap uint32)
	DisableClientState      func(cap uint32)
	EvalCoord1d             func(u float64)
	EvalCoord1f             func(u float32)
	EvalCoord2d             func(u, v float64)
	EvalCoord2f             func(u, v float32)
	EvalCoord1dv            func(u []float64)
	EvalCoord1fv            func(u []float32)
	EvalCoord2dv            func(u []float64)
	EvalCoord2fv            func(u []float32)
	EvalMesh1               func(mode uint32, i1, i2 int32)
	EvalMesh2               func(mode uint32, i1, i2, j1, j2 int32)
	EvalPoint1              func(i int32)
	EvalPoint2              func(i, j int32)
	FeedbackBuffer          func(size int32, Type uint32, buffer []float32)
	Finish                  func()
	Flush                   func()
	Fogf                    func(pname uint32, param float32)
	Fogi                    func(pname uint32, param int32)
	Fogfv                   func(pname uint32, params []float32)
	Fogiv                   func(pname uint32, params []int32)
	FrontFace               func(mode uint32)
	Frustum                 func(left, right, bottom, top, zNear, zFar float64)
	GenLists                func(Range int32) uint32
	GetBooleanv             func(pname uint32, params []bool)
	GetDoublev              func(pname uint32, params []float64)
	GetFloatv               func(pname uint32, params []float32)
	GetIntegerv             func(pname uint32, params []int32)
	GetClipPlane            func(plane uint32, equation []float64)
	GetError                func() uint32
	GetLightfv              func(light, pname uint32, params []float32)
	GetLightiv              func(light, pname uint32, params []int32)
	GetMapdv                func(target, query uint32, v []float64)
	GetMapfv                func(target, query uint32, v []float32)
	GetMapiv                func(target, query uint32, v []int32)
	GetMaterialfv           func(face, pname uint32, params []float32)
	GetMaterialiv           func(face, pname uint32, params []int32)
	GetPixelMapfv           func(Map uint32, values []float32)
	GetPixelMapuiv          func(Map uint32, values []uint32)
	GetPixelMapusv          func(Map uint32, values []uint16)
	GetPolygonStipple       func(mask []uint8)
	GetString               func(name uint32) string
	GetTexEnvfv             func(target, pname uint32, params []float32)
	GetTexEnviv             func(target, pname uint32, params []int32)
	GetTexGendv             func(coord, pname uint32, params []float64)
	GetTexGenfv             func(coord, pname uint32, params []float32)
	GetTexGeniv             func(coord, pname uint32, params []int32)
	GetTexImage             func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer)
	GetTexLevelParameterfv  func(target uint32, level int32, pname uint32, params []float32)
	GetTexLevelParameteriv  func(target uint32, level int32, pname uint32, params []int32)
	GetTexParameterfv       func(target, pname uint32, params []float32)
	GetTexParameteriv       func(target, pname uint32, params []int32)
	Hint                    func(target, mode uint32)
	Indexd                  func(c float64)
	Indexf                  func(c float32)
	Indexi                  func(c int32)
	Indexs                  func(c int16)
	Indexub                 func(c uint8)
	Indexdv                 func(c []float64)
	Indexfv                 func(c []float32)
	Indexiv                 func(c []int32)
	Indexsv                 func(c []int16)
	Indexubv                func(c []uint8)
	IndexMask               func(mask uint32)
	IndexPointer            func(Type uint32, stride int32, pointer unsafe.Pointer)
	InitNames               func()
	InterleavedArrays       func(format uint32, stride int32, pointer unsafe.Pointer)
	IsEnabled               func(cap uint32)
	IsList                  func(list uint32) bool
	Lightf                  func(light, pname uint32, param float32)
	Lighti                  func(light, pname uint32, param int32)
	Lightfv                 func(light, pname uint32, params []float32)
	Lightiv                 func(light, pname uint32, params []int32)
	LightModelf             func(pname uint32, param float32)
	LightModeli             func(pname uint32, param int32)
	LightModelfv            func(pname uint32, params []float32)
	LightModeliv            func(pname uint32, params []int32)
	LineStipple             func(factor int32, pattern uint16)
	LineWidth               func(width float32)
	ListBase                func(base uint32)
	LoadIdentity            func()
	LoadMatrixd             func(m []float64)
	LoadMatrixf             func(m []float32)
	LoadName                func(name uint32)
	LogicOp                 func(opcode uint32)
	Map1d                   func(target uint32, u1, u2 float64, stride, order int32, points []float64)
	Map1f                   func(target uint32, u1, u2 float32, stride, order int32, points []float32)
	Map2d                   func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points []float64)
	Map2f                   func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points []float32)
	MapGrid1d               func(un int32, u1, u2 float64)
	MapGrid1f               func(un int32, u1, u2 float32)
	MapGrid2d               func(un int32, u1, u2 float64, vn int32, v1, v2 float64)
	MapGrid2f               func(un int32, u1, u2 float32, vn int32, v1, v2 float32)
	Materialf               func(face, pname uint32, param float32)
	Materiali               func(face, pname uint32, param int32)
	Materialfv              func(face, pname uint32, params []float32)
	Materialiv              func(face, pname uint32, params []int32)
	MatrixMode              func(mode uint32)
	MultMatrixd             func(m []float64)
	MultMatrixf             func(m []float32)
	NewList                 func(list uint32, mode uint32)
	EndList                 func()
	Normal3b                func(nx, ny, nz int8)
	Normal3d                func(nx, ny, nz float64)
	Normal3f                func(nx, ny, nz float32)
	Normal3i                func(nx, ny, nz int32)
	Normal3s                func(nx, ny, nz int16)
	Normal3bv               func(v []int8)
	Normal3dv               func(v []float64)
	Normal3fv               func(v []float32)
	Normal3iv               func(v []int32)
	Normal3sv               func(v []int16)
	NormalPointer           func(Type uint32, stride int32, pointer unsafe.Pointer)
	Ortho                   func(left, right, bottom, top, zNear, zfar float64)
	PassThrough             func(token float32)
	PixelMapfv              func(Map uint32, mapsize int32, values []float32)
	PixelMapuiv             func(Map uint32, mapsize int32, values []uint32)
	PixelMapusv             func(Map uint32, mapsize int32, values []uint16)
	PixelStoref             func(pname uint32, param float32)
	PixelStorei             func(pname uint32, param int32)
	PixelTransferf          func(pname uint32, param float32)
	PixelTransferi          func(pname uint32, param int32)
	PixelZoom               func(xfactor, yfactor float32)
	PointSize               func(size float32)
	PolygonMode             func(face, mode uint32)
	PolygonStipple          func(mask []uint8)
	PrioritizeTextures      func(n int32, textures []uint32, priorities []float32)
	PushAttrib              func(mask uint32)
	PopAttrib               func()
	PushClientAttrib        func(mask uint32)
	PopClientAttrib         func()
	PushMatrix              func()
	PopMatrix               func()
	PushName                func(name uint32)
	PopName                 func()
	RasterPos2d             func(x, y float64)
	RasterPos2f             func(x, y float32)
	RasterPos2i             func(x, y int32)
	RasterPos2s             func(x, y int16)
	RasterPos3d             func(x, y, z float64)
	RasterPos3f             func(x, y, z float32)
	RasterPos3i             func(x, y, z int32)
	RasterPos3s             func(x, y, z int16)
	RasterPos4d             func(x, y, z, w float64)
	RasterPos4f             func(x, y, z, w float32)
	RasterPos4i             func(x, y, z, w int32)
	RasterPos4s             func(x, y, z, w int16)
	RasterPos2dv            func(v []float64)
	RasterPos2fv            func(v []float32)
	RasterPos2iv            func(v []int32)
	RasterPos2sv            func(v []int16)
	RasterPos3dv            func(v []float64)
	RasterPos3fv            func(v []float32)
	RasterPos3iv            func(v []int32)
	RasterPos3sv            func(v []int16)
	RasterPos4dv            func(v []float64)
	RasterPos4fv            func(v []float32)
	RasterPos4iv            func(v []int32)
	RasterPos4sv            func(v []int16)
	ReadBuffer              func(mode uint32)
	ReadPixels              func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer)
	Rectd                   func(x1, y1, x2, y2 float64)
	Rectf                   func(x1, y1, x2, y2 float32)
	Recti                   func(x1, y1, x2, y2 int32)
	Rects                   func(x1, y1, x2, y2 int16)
	Rectdv                  func(v1, v2 []float64)
	Rectfv                  func(v1, v2 []float32)
	Rectiv                  func(v1, v2 []int32)
	Rectsv                  func(v1, v2 []int16)
	RenderMode              func(mode uint32) int32
	Rotated                 func(angle, x, y, z float64)
	Rotatef                 func(angle, x, y, z float32)
	Scaled                  func(x, y, z float64)
	Scalef                  func(x, y, z float32)
	Scissor                 func(x, y int32, width, height int32)
	SelectBuffer            func(size int32, buffer []uint32)
	ShadeModel              func(mode uint32)
	StencilFunc             func(Func uint32, ref int32, mask uint32)
	StencilMask             func(mask uint32)
	StencilOp               func(fail, zfail, zpass uint32)
	TexCoord1d              func(s float64)
	TexCoord1f              func(s float32)
	TexCoord1i              func(s int32)
	TexCoord1s              func(s int16)
	TexCoord2d              func(s, t float64)
	TexCoord2f              func(s, t float32)
	TexCoord2i              func(s, t int32)
	TexCoord2s              func(s, t int16)
	TexCoord3d              func(s, t, r float64)
	TexCoord3f              func(s, t, r float32)
	TexCoord3i              func(s, t, r int32)
	TexCoord3s              func(s, t, r int16)
	TexCoord4d              func(s, t, r, q float64)
	TexCoord4f              func(s, t, r, q float32)
	TexCoord4i              func(s, t, r, q int32)
	TexCoord4s              func(s, t, r, q int16)
	TexCoord1dv             func(v []float64)
	TexCoord1fv             func(v []float32)
	TexCoord1iv             func(v []int32)
	TexCoord1sv             func(v []int16)
	TexCoord2dv             func(v []float64)
	TexCoord2fv             func(v []float32)
	TexCoord2iv             func(v []int32)
	TexCoord2sv             func(v []int16)
	TexCoord3dv             func(v []float64)
	TexCoord3fv             func(v []float32)
	TexCoord3iv             func(v []int32)
	TexCoord3sv             func(v []int16)
	TexCoord4dv             func(v []float64)
	TexCoord4fv             func(v []float32)
	TexCoord4iv             func(v []int32)
	TexCoord4sv             func(v []int16)
	TexCoordPointer         func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
	TexEnvf                 func(target, pname uint32, param float32)
	TexEnvi                 func(target, pname uint32, param int32)
	TexEnvfv                func(target, pname uint32, params []float32)
	TexEnviv                func(target, pname uint32, params []int32)
	TexGend                 func(coord, pname uint32, param float64)
	TexGenf                 func(coord, pname uint32, param float32)
	TexGeni                 func(coord, pname uint32, param int32)
	TexGendv                func(coord, pname uint32, params []float64)
	TexGenfv                func(coord, pname uint32, params []float32)
	TexGeniv                func(coord, pname uint32, params []int32)
	TexImage1D              func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer)
	TexImage2D              func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer)
	TexImage3DEXT           func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer)
	TexParameterf           func(target, pname uint32, param float32)
	TexParameteri           func(target, pname uint32, param int32)
	TexParameterfv          func(target, pname uint32, params []float32)
	TexParameteriv          func(target, pname uint32, params []int32)
	TexSubImage1D           func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer)
	TexSubImage2D           func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer)
	TexSubImage3DEXT        func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer)
	Translated              func(x, y, z float64)
	Translatef              func(x, y, z float32)
	Vertex2s                func(x, y int16)
	Vertex2i                func(x, y int32)
	Vertex2f                func(x, y float32)
	Vertex2d                func(x, y float64)
	Vertex3s                func(x, y, z int16)
	Vertex3i                func(x, y, z int32)
	Vertex3f                func(x, y, z float32)
	Vertex3d                func(x, y, z float64)
	Vertex4s                func(x, y, z, w int16)
	Vertex4i                func(x, y, z, w int32)
	Vertex4f                func(x, y, z, w float32)
	Vertex4d                func(x, y, z, w float64)
	VertexPointer           func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
	Viewport                func(x, y int32, width, height int32)
	AreTexturesResident     func(textures []uint32) (status bool, residencies []bool)
	ArrayElement            func(i int32)
	DrawArrays              func(mode uint32, first int32, count int32)
	DrawElements            func(mode uint32, count int32, Type uint32, indices unsafe.Pointer)
	GetPointerv             func(pname uint32, params unsafe.Pointer)
	PolygonOffset           func(factor, units float32)
	CopyTexImage1D          func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32)
	CopyTexImage2D          func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32)
	CopyTexSubImage1D       func(target uint32, level, xoffset, x, y int32, width int32)
	CopyTexSubImage2D       func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32)
	BindTexture             func(target uint32, texture uint32)
	DeleteTextures          func(n int32, textures []uint32)
	GenTextures             func(n int32, textures []uint32)
	IsTexture               func(texture uint32) bool
	ColorPointer            func(size int32, Type uint32, stride int32, pointer unsafe.Pointer)
	BlendColorEXT           func(red, green, blue, alpha float32)
	BlendEquation           func(mode uint32)
	CopyTexSubImage3D       func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32)
	ActiveTexture           func(texture uint32)
	ClientActiveTexture     func(texture uint32)
	CompressedTexImage1D    func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage2D    func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexImage3D    func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage1D func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage2D func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer)
	CompressedTexSubImage3D func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer)
	BlendFuncSeparate       func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32)
	BeginQuery              func(target uint32, id uint32)
	BindBuffer              func(target uint32, buffer uint32)
	BufferData              func(target uint32, size int32, data unsafe.Pointer, usage uint32)
	BufferSubData           func(target, offset uint32, size int32, data unsafe.Pointer)
}

func New() *Context {
	glc := new(Context)
	glc.context = C.gl15NewContext()

	if glc.context.fnAccum == nil {
		return nil
	}
	if glc.context.fnAlphaFunc == nil {
		return nil
	}
	if glc.context.fnBegin == nil {
		return nil
	}
	if glc.context.fnEnd == nil {
		return nil
	}
	if glc.context.fnBitmap == nil {
		return nil
	}
	if glc.context.fnBlendFunc == nil {
		return nil
	}
	if glc.context.fnCallList == nil {
		return nil
	}
	if glc.context.fnCallLists == nil {
		return nil
	}
	if glc.context.fnClear == nil {
		return nil
	}
	if glc.context.fnClearAccum == nil {
		return nil
	}
	if glc.context.fnClearColor == nil {
		return nil
	}
	if glc.context.fnClearDepth == nil {
		return nil
	}
	if glc.context.fnClearIndex == nil {
		return nil
	}
	if glc.context.fnClearStencil == nil {
		return nil
	}
	if glc.context.fnClipPlane == nil {
		return nil
	}
	if glc.context.fnColor3b == nil {
		return nil
	}
	if glc.context.fnColor3d == nil {
		return nil
	}
	if glc.context.fnColor3f == nil {
		return nil
	}
	if glc.context.fnColor3i == nil {
		return nil
	}
	if glc.context.fnColor3s == nil {
		return nil
	}
	if glc.context.fnColor3ub == nil {
		return nil
	}
	if glc.context.fnColor3ui == nil {
		return nil
	}
	if glc.context.fnColor3us == nil {
		return nil
	}
	if glc.context.fnColor4b == nil {
		return nil
	}
	if glc.context.fnColor4d == nil {
		return nil
	}
	if glc.context.fnColor4f == nil {
		return nil
	}
	if glc.context.fnColor4i == nil {
		return nil
	}
	if glc.context.fnColor4s == nil {
		return nil
	}
	if glc.context.fnColor4ub == nil {
		return nil
	}
	if glc.context.fnColor4ui == nil {
		return nil
	}
	if glc.context.fnColor4us == nil {
		return nil
	}
	if glc.context.fnColor3bv == nil {
		return nil
	}
	if glc.context.fnColor3dv == nil {
		return nil
	}
	if glc.context.fnColor3fv == nil {
		return nil
	}
	if glc.context.fnColor3iv == nil {
		return nil
	}
	if glc.context.fnColor3sv == nil {
		return nil
	}
	if glc.context.fnColor3ubv == nil {
		return nil
	}
	if glc.context.fnColor3uiv == nil {
		return nil
	}
	if glc.context.fnColor3usv == nil {
		return nil
	}
	if glc.context.fnColor4bv == nil {
		return nil
	}
	if glc.context.fnColor4dv == nil {
		return nil
	}
	if glc.context.fnColor4fv == nil {
		return nil
	}
	if glc.context.fnColor4iv == nil {
		return nil
	}
	if glc.context.fnColor4sv == nil {
		return nil
	}
	if glc.context.fnColor4ubv == nil {
		return nil
	}
	if glc.context.fnColor4uiv == nil {
		return nil
	}
	if glc.context.fnColor4usv == nil {
		return nil
	}
	if glc.context.fnColorMask == nil {
		return nil
	}
	if glc.context.fnColorMaterial == nil {
		return nil
	}
	if glc.context.fnCopyPixels == nil {
		return nil
	}
	if glc.context.fnCullFace == nil {
		return nil
	}
	if glc.context.fnDeleteLists == nil {
		return nil
	}
	if glc.context.fnDepthFunc == nil {
		return nil
	}
	if glc.context.fnDepthRange == nil {
		return nil
	}
	if glc.context.fnEnable == nil {
		return nil
	}
	if glc.context.fnDisable == nil {
		return nil
	}
	if glc.context.fnDrawBuffer == nil {
		return nil
	}
	if glc.context.fnDrawPixels == nil {
		return nil
	}
	if glc.context.fnEdgeFlag == nil {
		return nil
	}
	if glc.context.fnEdgeFlagv == nil {
		return nil
	}
	if glc.context.fnEdgeFlagPointer == nil {
		return nil
	}
	if glc.context.fnEnableClientState == nil {
		return nil
	}
	if glc.context.fnDisableClientState == nil {
		return nil
	}
	if glc.context.fnEvalCoord1d == nil {
		return nil
	}
	if glc.context.fnEvalCoord1f == nil {
		return nil
	}
	if glc.context.fnEvalCoord2d == nil {
		return nil
	}
	if glc.context.fnEvalCoord2f == nil {
		return nil
	}
	if glc.context.fnEvalCoord1dv == nil {
		return nil
	}
	if glc.context.fnEvalCoord1fv == nil {
		return nil
	}
	if glc.context.fnEvalCoord2dv == nil {
		return nil
	}
	if glc.context.fnEvalCoord2fv == nil {
		return nil
	}
	if glc.context.fnEvalMesh1 == nil {
		return nil
	}
	if glc.context.fnEvalMesh2 == nil {
		return nil
	}
	if glc.context.fnEvalPoint1 == nil {
		return nil
	}
	if glc.context.fnEvalPoint2 == nil {
		return nil
	}
	if glc.context.fnFeedbackBuffer == nil {
		return nil
	}
	if glc.context.fnFinish == nil {
		return nil
	}
	if glc.context.fnFlush == nil {
		return nil
	}
	if glc.context.fnFogf == nil {
		return nil
	}
	if glc.context.fnFogi == nil {
		return nil
	}
	if glc.context.fnFogfv == nil {
		return nil
	}
	if glc.context.fnFogiv == nil {
		return nil
	}
	if glc.context.fnFrontFace == nil {
		return nil
	}
	if glc.context.fnFrustum == nil {
		return nil
	}
	if glc.context.fnGenLists == nil {
		return nil
	}
	if glc.context.fnGetBooleanv == nil {
		return nil
	}
	if glc.context.fnGetDoublev == nil {
		return nil
	}
	if glc.context.fnGetFloatv == nil {
		return nil
	}
	if glc.context.fnGetIntegerv == nil {
		return nil
	}
	if glc.context.fnGetClipPlane == nil {
		return nil
	}
	if glc.context.fnGetError == nil {
		return nil
	}
	if glc.context.fnGetLightfv == nil {
		return nil
	}
	if glc.context.fnGetLightiv == nil {
		return nil
	}
	if glc.context.fnGetMapdv == nil {
		return nil
	}
	if glc.context.fnGetMapfv == nil {
		return nil
	}
	if glc.context.fnGetMapiv == nil {
		return nil
	}
	if glc.context.fnGetMaterialfv == nil {
		return nil
	}
	if glc.context.fnGetMaterialiv == nil {
		return nil
	}
	if glc.context.fnGetPixelMapfv == nil {
		return nil
	}
	if glc.context.fnGetPixelMapuiv == nil {
		return nil
	}
	if glc.context.fnGetPixelMapusv == nil {
		return nil
	}
	if glc.context.fnGetPolygonStipple == nil {
		return nil
	}
	if glc.context.fnGetString == nil {
		return nil
	}
	if glc.context.fnGetTexEnvfv == nil {
		return nil
	}
	if glc.context.fnGetTexEnviv == nil {
		return nil
	}
	if glc.context.fnGetTexGendv == nil {
		return nil
	}
	if glc.context.fnGetTexGenfv == nil {
		return nil
	}
	if glc.context.fnGetTexGeniv == nil {
		return nil
	}
	if glc.context.fnGetTexImage == nil {
		return nil
	}
	if glc.context.fnGetTexLevelParameterfv == nil {
		return nil
	}
	if glc.context.fnGetTexLevelParameteriv == nil {
		return nil
	}
	if glc.context.fnGetTexParameterfv == nil {
		return nil
	}
	if glc.context.fnGetTexParameteriv == nil {
		return nil
	}
	if glc.context.fnHint == nil {
		return nil
	}
	if glc.context.fnIndexd == nil {
		return nil
	}
	if glc.context.fnIndexf == nil {
		return nil
	}
	if glc.context.fnIndexi == nil {
		return nil
	}
	if glc.context.fnIndexs == nil {
		return nil
	}
	if glc.context.fnIndexub == nil {
		return nil
	}
	if glc.context.fnIndexdv == nil {
		return nil
	}
	if glc.context.fnIndexfv == nil {
		return nil
	}
	if glc.context.fnIndexiv == nil {
		return nil
	}
	if glc.context.fnIndexsv == nil {
		return nil
	}
	if glc.context.fnIndexubv == nil {
		return nil
	}
	if glc.context.fnIndexMask == nil {
		return nil
	}
	if glc.context.fnIndexPointer == nil {
		return nil
	}
	if glc.context.fnInitNames == nil {
		return nil
	}
	if glc.context.fnInterleavedArrays == nil {
		return nil
	}
	if glc.context.fnIsEnabled == nil {
		return nil
	}
	if glc.context.fnIsList == nil {
		return nil
	}
	if glc.context.fnLightf == nil {
		return nil
	}
	if glc.context.fnLighti == nil {
		return nil
	}
	if glc.context.fnLightfv == nil {
		return nil
	}
	if glc.context.fnLightiv == nil {
		return nil
	}
	if glc.context.fnLightModelf == nil {
		return nil
	}
	if glc.context.fnLightModeli == nil {
		return nil
	}
	if glc.context.fnLightModelfv == nil {
		return nil
	}
	if glc.context.fnLightModeliv == nil {
		return nil
	}
	if glc.context.fnLineStipple == nil {
		return nil
	}
	if glc.context.fnLineWidth == nil {
		return nil
	}
	if glc.context.fnListBase == nil {
		return nil
	}
	if glc.context.fnLoadIdentity == nil {
		return nil
	}
	if glc.context.fnLoadMatrixd == nil {
		return nil
	}
	if glc.context.fnLoadMatrixf == nil {
		return nil
	}
	if glc.context.fnLoadName == nil {
		return nil
	}
	if glc.context.fnLogicOp == nil {
		return nil
	}
	if glc.context.fnMap1d == nil {
		return nil
	}
	if glc.context.fnMap1f == nil {
		return nil
	}
	if glc.context.fnMap2d == nil {
		return nil
	}
	if glc.context.fnMap2f == nil {
		return nil
	}
	if glc.context.fnMapGrid1d == nil {
		return nil
	}
	if glc.context.fnMapGrid1f == nil {
		return nil
	}
	if glc.context.fnMapGrid2d == nil {
		return nil
	}
	if glc.context.fnMapGrid2f == nil {
		return nil
	}
	if glc.context.fnMaterialf == nil {
		return nil
	}
	if glc.context.fnMateriali == nil {
		return nil
	}
	if glc.context.fnMaterialfv == nil {
		return nil
	}
	if glc.context.fnMaterialiv == nil {
		return nil
	}
	if glc.context.fnMatrixMode == nil {
		return nil
	}
	if glc.context.fnMultMatrixd == nil {
		return nil
	}
	if glc.context.fnMultMatrixf == nil {
		return nil
	}
	if glc.context.fnNewList == nil {
		return nil
	}
	if glc.context.fnEndList == nil {
		return nil
	}
	if glc.context.fnNormal3b == nil {
		return nil
	}
	if glc.context.fnNormal3d == nil {
		return nil
	}
	if glc.context.fnNormal3f == nil {
		return nil
	}
	if glc.context.fnNormal3i == nil {
		return nil
	}
	if glc.context.fnNormal3s == nil {
		return nil
	}
	if glc.context.fnNormal3bv == nil {
		return nil
	}
	if glc.context.fnNormal3dv == nil {
		return nil
	}
	if glc.context.fnNormal3fv == nil {
		return nil
	}
	if glc.context.fnNormal3iv == nil {
		return nil
	}
	if glc.context.fnNormal3sv == nil {
		return nil
	}
	if glc.context.fnNormalPointer == nil {
		return nil
	}
	if glc.context.fnOrtho == nil {
		return nil
	}
	if glc.context.fnPassThrough == nil {
		return nil
	}
	if glc.context.fnPixelMapfv == nil {
		return nil
	}
	if glc.context.fnPixelMapuiv == nil {
		return nil
	}
	if glc.context.fnPixelMapusv == nil {
		return nil
	}
	if glc.context.fnPixelStoref == nil {
		return nil
	}
	if glc.context.fnPixelStorei == nil {
		return nil
	}
	if glc.context.fnPixelTransferf == nil {
		return nil
	}
	if glc.context.fnPixelTransferi == nil {
		return nil
	}
	if glc.context.fnPixelZoom == nil {
		return nil
	}
	if glc.context.fnPointSize == nil {
		return nil
	}
	if glc.context.fnPolygonMode == nil {
		return nil
	}
	if glc.context.fnPolygonStipple == nil {
		return nil
	}
	if glc.context.fnPrioritizeTextures == nil {
		return nil
	}
	if glc.context.fnPushAttrib == nil {
		return nil
	}
	if glc.context.fnPopAttrib == nil {
		return nil
	}
	if glc.context.fnPushClientAttrib == nil {
		return nil
	}
	if glc.context.fnPopClientAttrib == nil {
		return nil
	}
	if glc.context.fnPushMatrix == nil {
		return nil
	}
	if glc.context.fnPopMatrix == nil {
		return nil
	}
	if glc.context.fnPushName == nil {
		return nil
	}
	if glc.context.fnPopName == nil {
		return nil
	}
	if glc.context.fnRasterPos2d == nil {
		return nil
	}
	if glc.context.fnRasterPos2f == nil {
		return nil
	}
	if glc.context.fnRasterPos2i == nil {
		return nil
	}
	if glc.context.fnRasterPos2s == nil {
		return nil
	}
	if glc.context.fnRasterPos3d == nil {
		return nil
	}
	if glc.context.fnRasterPos3f == nil {
		return nil
	}
	if glc.context.fnRasterPos3i == nil {
		return nil
	}
	if glc.context.fnRasterPos3s == nil {
		return nil
	}
	if glc.context.fnRasterPos4d == nil {
		return nil
	}
	if glc.context.fnRasterPos4f == nil {
		return nil
	}
	if glc.context.fnRasterPos4i == nil {
		return nil
	}
	if glc.context.fnRasterPos4s == nil {
		return nil
	}
	if glc.context.fnRasterPos2dv == nil {
		return nil
	}
	if glc.context.fnRasterPos2fv == nil {
		return nil
	}
	if glc.context.fnRasterPos2iv == nil {
		return nil
	}
	if glc.context.fnRasterPos2sv == nil {
		return nil
	}
	if glc.context.fnRasterPos3dv == nil {
		return nil
	}
	if glc.context.fnRasterPos3fv == nil {
		return nil
	}
	if glc.context.fnRasterPos3iv == nil {
		return nil
	}
	if glc.context.fnRasterPos3sv == nil {
		return nil
	}
	if glc.context.fnRasterPos4dv == nil {
		return nil
	}
	if glc.context.fnRasterPos4fv == nil {
		return nil
	}
	if glc.context.fnRasterPos4iv == nil {
		return nil
	}
	if glc.context.fnRasterPos4sv == nil {
		return nil
	}
	if glc.context.fnReadBuffer == nil {
		return nil
	}
	if glc.context.fnReadPixels == nil {
		return nil
	}
	if glc.context.fnRectd == nil {
		return nil
	}
	if glc.context.fnRectf == nil {
		return nil
	}
	if glc.context.fnRecti == nil {
		return nil
	}
	if glc.context.fnRects == nil {
		return nil
	}
	if glc.context.fnRectdv == nil {
		return nil
	}
	if glc.context.fnRectfv == nil {
		return nil
	}
	if glc.context.fnRectiv == nil {
		return nil
	}
	if glc.context.fnRectsv == nil {
		return nil
	}
	if glc.context.fnRenderMode == nil {
		return nil
	}
	if glc.context.fnRotated == nil {
		return nil
	}
	if glc.context.fnRotatef == nil {
		return nil
	}
	if glc.context.fnScaled == nil {
		return nil
	}
	if glc.context.fnScalef == nil {
		return nil
	}
	if glc.context.fnScissor == nil {
		return nil
	}
	if glc.context.fnSelectBuffer == nil {
		return nil
	}
	if glc.context.fnShadeModel == nil {
		return nil
	}
	if glc.context.fnStencilFunc == nil {
		return nil
	}
	if glc.context.fnStencilMask == nil {
		return nil
	}
	if glc.context.fnStencilOp == nil {
		return nil
	}
	if glc.context.fnTexCoord1d == nil {
		return nil
	}
	if glc.context.fnTexCoord1f == nil {
		return nil
	}
	if glc.context.fnTexCoord1i == nil {
		return nil
	}
	if glc.context.fnTexCoord1s == nil {
		return nil
	}
	if glc.context.fnTexCoord2d == nil {
		return nil
	}
	if glc.context.fnTexCoord2f == nil {
		return nil
	}
	if glc.context.fnTexCoord2i == nil {
		return nil
	}
	if glc.context.fnTexCoord2s == nil {
		return nil
	}
	if glc.context.fnTexCoord3d == nil {
		return nil
	}
	if glc.context.fnTexCoord3f == nil {
		return nil
	}
	if glc.context.fnTexCoord3i == nil {
		return nil
	}
	if glc.context.fnTexCoord3s == nil {
		return nil
	}
	if glc.context.fnTexCoord4d == nil {
		return nil
	}
	if glc.context.fnTexCoord4f == nil {
		return nil
	}
	if glc.context.fnTexCoord4i == nil {
		return nil
	}
	if glc.context.fnTexCoord4s == nil {
		return nil
	}
	if glc.context.fnTexCoord1dv == nil {
		return nil
	}
	if glc.context.fnTexCoord1fv == nil {
		return nil
	}
	if glc.context.fnTexCoord1iv == nil {
		return nil
	}
	if glc.context.fnTexCoord1sv == nil {
		return nil
	}
	if glc.context.fnTexCoord2dv == nil {
		return nil
	}
	if glc.context.fnTexCoord2fv == nil {
		return nil
	}
	if glc.context.fnTexCoord2iv == nil {
		return nil
	}
	if glc.context.fnTexCoord2sv == nil {
		return nil
	}
	if glc.context.fnTexCoord3dv == nil {
		return nil
	}
	if glc.context.fnTexCoord3fv == nil {
		return nil
	}
	if glc.context.fnTexCoord3iv == nil {
		return nil
	}
	if glc.context.fnTexCoord3sv == nil {
		return nil
	}
	if glc.context.fnTexCoord4dv == nil {
		return nil
	}
	if glc.context.fnTexCoord4fv == nil {
		return nil
	}
	if glc.context.fnTexCoord4iv == nil {
		return nil
	}
	if glc.context.fnTexCoord4sv == nil {
		return nil
	}
	if glc.context.fnTexCoordPointer == nil {
		return nil
	}
	if glc.context.fnTexEnvf == nil {
		return nil
	}
	if glc.context.fnTexEnvi == nil {
		return nil
	}
	if glc.context.fnTexEnvfv == nil {
		return nil
	}
	if glc.context.fnTexEnviv == nil {
		return nil
	}
	if glc.context.fnTexGend == nil {
		return nil
	}
	if glc.context.fnTexGenf == nil {
		return nil
	}
	if glc.context.fnTexGeni == nil {
		return nil
	}
	if glc.context.fnTexGendv == nil {
		return nil
	}
	if glc.context.fnTexGenfv == nil {
		return nil
	}
	if glc.context.fnTexGeniv == nil {
		return nil
	}
	if glc.context.fnTexImage1D == nil {
		return nil
	}
	if glc.context.fnTexImage2D == nil {
		return nil
	}
	if glc.context.fnTexParameterf == nil {
		return nil
	}
	if glc.context.fnTexParameteri == nil {
		return nil
	}
	if glc.context.fnTexParameterfv == nil {
		return nil
	}
	if glc.context.fnTexParameteriv == nil {
		return nil
	}
	if glc.context.fnTexSubImage1D == nil {
		return nil
	}
	if glc.context.fnTexSubImage2D == nil {
		return nil
	}
	if glc.context.fnTranslated == nil {
		return nil
	}
	if glc.context.fnTranslatef == nil {
		return nil
	}
	if glc.context.fnVertex2s == nil {
		return nil
	}
	if glc.context.fnVertex2i == nil {
		return nil
	}
	if glc.context.fnVertex2f == nil {
		return nil
	}
	if glc.context.fnVertex2d == nil {
		return nil
	}
	if glc.context.fnVertex3s == nil {
		return nil
	}
	if glc.context.fnVertex3i == nil {
		return nil
	}
	if glc.context.fnVertex3f == nil {
		return nil
	}
	if glc.context.fnVertex3d == nil {
		return nil
	}
	if glc.context.fnVertex4s == nil {
		return nil
	}
	if glc.context.fnVertex4i == nil {
		return nil
	}
	if glc.context.fnVertex4f == nil {
		return nil
	}
	if glc.context.fnVertex4d == nil {
		return nil
	}
	if glc.context.fnVertexPointer == nil {
		return nil
	}
	if glc.context.fnViewport == nil {
		return nil
	}
	if glc.context.fnAreTexturesResident == nil {
		return nil
	}
	if glc.context.fnArrayElement == nil {
		return nil
	}
	if glc.context.fnDrawArrays == nil {
		return nil
	}
	if glc.context.fnDrawElements == nil {
		return nil
	}
	if glc.context.fnGetPointerv == nil {
		return nil
	}
	if glc.context.fnPolygonOffset == nil {
		return nil
	}
	if glc.context.fnCopyTexImage1D == nil {
		return nil
	}
	if glc.context.fnCopyTexImage2D == nil {
		return nil
	}
	if glc.context.fnCopyTexSubImage1D == nil {
		return nil
	}
	if glc.context.fnCopyTexSubImage2D == nil {
		return nil
	}
	if glc.context.fnBindTexture == nil {
		return nil
	}
	if glc.context.fnDeleteTextures == nil {
		return nil
	}
	if glc.context.fnGenTextures == nil {
		return nil
	}
	if glc.context.fnIsTexture == nil {
		return nil
	}
	if glc.context.fnColorPointer == nil {
		return nil
	}

	glc.Accum = func(op uint32, value float32) {
		C.gl15Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		C.gl15AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		C.gl15Begin(glc.context, C.GLenum(mode))
	}

	glc.End = func() {
		C.gl15End(glc.context)
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap []uint8) {
		C.gl15Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(&bitmap[0])))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		C.gl15BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		C.gl15CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		C.gl15CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		C.gl15Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		C.gl15ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		C.gl15ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		C.gl15ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		C.gl15ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		C.gl15ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation []float64) {
		C.gl15ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(&equation[0])))
	}

	glc.Color3b = func(red, green, blue int8) {
		C.gl15Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		C.gl15Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		C.gl15Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		C.gl15Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		C.gl15Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		C.gl15Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		C.gl15Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		C.gl15Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		C.gl15Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		C.gl15Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		C.gl15Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		C.gl15Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		C.gl15Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		C.gl15Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		C.gl15Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		C.gl15Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v []int8) {
		C.gl15Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color3dv = func(v []float64) {
		C.gl15Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.Color3fv = func(v []float32) {
		C.gl15Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.Color3iv = func(v []int32) {
		C.gl15Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.Color3sv = func(v []int16) {
		C.gl15Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.Color3ubv = func(v []uint8) {
		C.gl15Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color3uiv = func(v []uint32) {
		C.gl15Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(&v[0])))
	}

	glc.Color3usv = func(v []uint16) {
		C.gl15Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(&v[0])))
	}

	glc.Color4bv = func(v []int8) {
		C.gl15Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color4dv = func(v []float64) {
		C.gl15Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.Color4fv = func(v []float32) {
		C.gl15Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.Color4iv = func(v []int32) {
		C.gl15Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.Color4sv = func(v []int16) {
		C.gl15Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.Color4ubv = func(v []uint8) {
		C.gl15Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color4uiv = func(v []uint32) {
		C.gl15Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(&v[0])))
	}

	glc.Color4usv = func(v []uint16) {
		C.gl15Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(&v[0])))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		C.gl15ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		C.gl15ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl15ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params []float32) {
		C.gl15ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params []int32) {
		C.gl15ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		C.gl15ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		C.gl15CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		C.gl15CullFace(glc.context, C.GLenum(mode))
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl15ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl15ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		C.gl15ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		C.gl15ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl15CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		C.gl15CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl15CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		C.gl15CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		C.gl15DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		C.gl15DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		C.gl15DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		C.gl15Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		C.gl15Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		C.gl15DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.EdgeFlag = func(flag bool) {
		C.gl15EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag []bool) {
		C.gl15EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(&flag[0])))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		C.gl15EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		C.gl15EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		C.gl15DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.EvalCoord1d = func(u float64) {
		C.gl15EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		C.gl15EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		C.gl15EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		C.gl15EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u []float64) {
		C.gl15EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&u[0])))
	}

	glc.EvalCoord1fv = func(u []float32) {
		C.gl15EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&u[0])))
	}

	glc.EvalCoord2dv = func(u []float64) {
		C.gl15EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&u[0])))
	}

	glc.EvalCoord2fv = func(u []float32) {
		C.gl15EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&u[0])))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		C.gl15EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		C.gl15EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		C.gl15EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		C.gl15EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer []float32) {
		C.gl15FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(&buffer[0])))
	}

	glc.Finish = func() {
		C.gl15Finish(glc.context)
	}

	glc.Flush = func() {
		C.gl15Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		C.gl15Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		C.gl15Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params []float32) {
		C.gl15Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.Fogiv = func(pname uint32, params []int32) {
		C.gl15Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.FrontFace = func(mode uint32) {
		C.gl15FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		C.gl15Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		return uint32(C.gl15GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params []bool) {
		C.gl15GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(&params[0])))
	}

	glc.GetDoublev = func(pname uint32, params []float64) {
		C.gl15GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))
	}

	glc.GetFloatv = func(pname uint32, params []float32) {
		C.gl15GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetIntegerv = func(pname uint32, params []int32) {
		C.gl15GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetClipPlane = func(plane uint32, equation []float64) {
		C.gl15GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(&equation[0])))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl15GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params []float32) {
		C.gl15GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetLightiv = func(light, pname uint32, params []int32) {
		C.gl15GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetMapdv = func(target, query uint32, v []float64) {
		C.gl15GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.GetMapfv = func(target, query uint32, v []float32) {
		C.gl15GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.GetMapiv = func(target, query uint32, v []int32) {
		C.gl15GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.GetMaterialfv = func(face, pname uint32, params []float32) {
		C.gl15GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetMaterialiv = func(face, pname uint32, params []int32) {
		C.gl15GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetPixelMapfv = func(Map uint32, values []float32) {
		C.gl15GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(&values[0])))
	}

	glc.GetPixelMapuiv = func(Map uint32, values []uint32) {
		C.gl15GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(&values[0])))
	}

	glc.GetPixelMapusv = func(Map uint32, values []uint16) {
		C.gl15GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(&values[0])))
	}

	glc.GetPolygonStipple = func(mask []uint8) {
		C.gl15GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(&mask[0])))
	}

	glc.GetString = func(name uint32) string {
		cstr := C.gl15GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params []float32) {
		C.gl15GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexEnviv = func(target, pname uint32, params []int32) {
		C.gl15GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexGendv = func(coord, pname uint32, params []float64) {
		C.gl15GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params []float32) {
		C.gl15GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params []int32) {
		C.gl15GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params []float32) {
		C.gl15GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params []int32) {
		C.gl15GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params []float32) {
		C.gl15GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params []int32) {
		C.gl15GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.Hint = func(target, mode uint32) {
		C.gl15Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		C.gl15Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		C.gl15Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		C.gl15Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		C.gl15Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexub = func(c uint8) {
		C.gl15Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexdv = func(c []float64) {
		C.gl15Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(&c[0])))
	}

	glc.Indexfv = func(c []float32) {
		C.gl15Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(&c[0])))
	}

	glc.Indexiv = func(c []int32) {
		C.gl15Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(&c[0])))
	}

	glc.Indexsv = func(c []int16) {
		C.gl15Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(&c[0])))
	}

	glc.Indexubv = func(c []uint8) {
		C.gl15Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(&c[0])))
	}

	glc.IndexMask = func(mask uint32) {
		C.gl15IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl15IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		C.gl15InitNames(glc.context)
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		C.gl15InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.IsEnabled = func(cap uint32) {
		C.gl15IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		return C.gl15IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		C.gl15Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		C.gl15Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params []float32) {
		C.gl15Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.Lightiv = func(light, pname uint32, params []int32) {
		C.gl15Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		C.gl15LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		C.gl15LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params []float32) {
		C.gl15LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.LightModeliv = func(pname uint32, params []int32) {
		C.gl15LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		C.gl15LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		C.gl15LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		C.gl15ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		C.gl15LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m []float64) {
		C.gl15LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(&m[0])))
	}

	glc.LoadMatrixf = func(m []float32) {
		C.gl15LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(&m[0])))
	}

	glc.LoadName = func(name uint32) {
		C.gl15LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		C.gl15LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points []float64) {
		C.gl15Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(&points[0])))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points []float32) {
		C.gl15Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(&points[0])))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points []float64) {
		C.gl15Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(&points[0])))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points []float32) {
		C.gl15Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(&points[0])))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		C.gl15MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		C.gl15MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		C.gl15MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		C.gl15MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		C.gl15Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		C.gl15Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params []float32) {
		C.gl15Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.Materialiv = func(face, pname uint32, params []int32) {
		C.gl15Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.MatrixMode = func(mode uint32) {
		C.gl15MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m []float64) {
		C.gl15MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(&m[0])))
	}

	glc.MultMatrixf = func(m []float32) {
		C.gl15MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(&m[0])))
	}

	glc.NewList = func(list uint32, mode uint32) {
		C.gl15NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		C.gl15EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		C.gl15Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		C.gl15Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		C.gl15Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		C.gl15Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		C.gl15Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v []int8) {
		C.gl15Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3dv = func(v []float64) {
		C.gl15Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3fv = func(v []float32) {
		C.gl15Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3iv = func(v []int32) {
		C.gl15Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3sv = func(v []int16) {
		C.gl15Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl15NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		C.gl15Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		C.gl15PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values []float32) {
		C.gl15PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(&values[0])))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values []uint32) {
		C.gl15PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(&values[0])))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values []uint16) {
		C.gl15PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(&values[0])))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		C.gl15PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		C.gl15PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		C.gl15PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		C.gl15PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		C.gl15PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		C.gl15PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		C.gl15PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask []uint8) {
		C.gl15PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(&mask[0])))
	}

	glc.PrioritizeTextures = func(n int32, textures []uint32, priorities []float32) {
		C.gl15PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(&textures[0])), (*C.GLclampf)(unsafe.Pointer(&priorities[0])))
	}

	glc.PushAttrib = func(mask uint32) {
		C.gl15PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		C.gl15PopAttrib(glc.context)
	}

	glc.PushClientAttrib = func(mask uint32) {
		C.gl15PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopClientAttrib = func() {
		C.gl15PopClientAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		C.gl15PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		C.gl15PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		C.gl15PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		C.gl15PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		C.gl15RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		C.gl15RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		C.gl15RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		C.gl15RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		C.gl15RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		C.gl15RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		C.gl15RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		C.gl15RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		C.gl15RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		C.gl15RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		C.gl15RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		C.gl15RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v []float64) {
		C.gl15RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos2fv = func(v []float32) {
		C.gl15RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos2iv = func(v []int32) {
		C.gl15RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos2sv = func(v []int16) {
		C.gl15RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3dv = func(v []float64) {
		C.gl15RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3fv = func(v []float32) {
		C.gl15RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3iv = func(v []int32) {
		C.gl15RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3sv = func(v []int16) {
		C.gl15RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4dv = func(v []float64) {
		C.gl15RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4fv = func(v []float32) {
		C.gl15RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4iv = func(v []int32) {
		C.gl15RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4sv = func(v []int16) {
		C.gl15RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.ReadBuffer = func(mode uint32) {
		C.gl15ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		C.gl15Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		C.gl15Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		C.gl15Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		C.gl15Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 []float64) {
		C.gl15Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v1[0])), (*C.GLdouble)(unsafe.Pointer(&v2[0])))
	}

	glc.Rectfv = func(v1, v2 []float32) {
		C.gl15Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v1[0])), (*C.GLfloat)(unsafe.Pointer(&v2[0])))
	}

	glc.Rectiv = func(v1, v2 []int32) {
		C.gl15Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(&v1[0])), (*C.GLint)(unsafe.Pointer(&v2[0])))
	}

	glc.Rectsv = func(v1, v2 []int16) {
		C.gl15Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(&v1[0])), (*C.GLshort)(unsafe.Pointer(&v2[0])))
	}

	glc.RenderMode = func(mode uint32) int32 {
		return int32(C.gl15RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		C.gl15Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		C.gl15Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		C.gl15Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		C.gl15Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		C.gl15Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer []uint32) {
		C.gl15SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(&buffer[0])))
	}

	glc.ShadeModel = func(mode uint32) {
		C.gl15ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		C.gl15StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		C.gl15StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		C.gl15StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		C.gl15TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		C.gl15TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		C.gl15TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		C.gl15TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		C.gl15TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		C.gl15TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		C.gl15TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		C.gl15TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		C.gl15TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		C.gl15TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		C.gl15TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		C.gl15TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		C.gl15TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		C.gl15TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		C.gl15TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		C.gl15TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v []float64) {
		C.gl15TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord1fv = func(v []float32) {
		C.gl15TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord1iv = func(v []int32) {
		C.gl15TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord1sv = func(v []int16) {
		C.gl15TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2dv = func(v []float64) {
		C.gl15TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2fv = func(v []float32) {
		C.gl15TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2iv = func(v []int32) {
		C.gl15TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2sv = func(v []int16) {
		C.gl15TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3dv = func(v []float64) {
		C.gl15TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3fv = func(v []float32) {
		C.gl15TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3iv = func(v []int32) {
		C.gl15TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3sv = func(v []int16) {
		C.gl15TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4dv = func(v []float64) {
		C.gl15TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4fv = func(v []float32) {
		C.gl15TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4iv = func(v []int32) {
		C.gl15TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4sv = func(v []int16) {
		C.gl15TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl15TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		C.gl15TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		C.gl15TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params []float32) {
		C.gl15TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.TexEnviv = func(target, pname uint32, params []int32) {
		C.gl15TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		C.gl15TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		C.gl15TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		C.gl15TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params []float64) {
		C.gl15TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))
	}

	glc.TexGenfv = func(coord, pname uint32, params []float32) {
		C.gl15TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.TexGeniv = func(coord, pname uint32, params []int32) {
		C.gl15TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage3DEXT = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15TexImage3DEXT(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		C.gl15TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		C.gl15TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params []float32) {
		C.gl15TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.TexParameteriv = func(target, pname uint32, params []int32) {
		C.gl15TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3DEXT = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl15TexSubImage3DEXT(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Translated = func(x, y, z float64) {
		C.gl15Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		C.gl15Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		C.gl15Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		C.gl15Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		C.gl15Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		C.gl15Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		C.gl15Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		C.gl15Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		C.gl15Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		C.gl15Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		C.gl15Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		C.gl15Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		C.gl15Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		C.gl15Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl15VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		C.gl15Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		var cRes *C.GLboolean
		status = C.gl15AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		C.gl15ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		C.gl15DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl15DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		C.gl15GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		C.gl15PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		C.gl15CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		C.gl15CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		C.gl15CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		C.gl15CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		C.gl15BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures []uint32) {
		C.gl15DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(&textures[0])))
	}

	glc.GenTextures = func(n int32, textures []uint32) {
		C.gl15GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(&textures[0])))
	}

	glc.IsTexture = func(texture uint32) bool {
		return C.gl15IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl15ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.BlendColorEXT = func(red, green, blue, alpha float32) {
		C.gl15BlendColorEXT(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		C.gl15BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		C.gl15CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.ActiveTexture = func(texture uint32) {
		C.gl15ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		C.gl15ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl15CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl15CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl15CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl15CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl15CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl15CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		C.gl15BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		C.gl15BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		C.gl15BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		C.gl15BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		C.gl15BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	return glc
}
