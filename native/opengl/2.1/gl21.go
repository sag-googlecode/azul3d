// Package 'opengl' implements OpenGL version 2.1
package opengl

// #cgo LDFLAGS: -lopengl32
// #include "gl21.h"
import "C"

import "unsafe"

func boolToGL(b bool) C.GLboolean {
	if b {
		return C.GLboolean(1)
	}
	return C.GLboolean(0)
}

const (
	T                                                          = 0x2001
	FOG_COORD_SRC                                              = 0x8450
	FOG_COORD                                                  = 0x8451
	MATRIX7_ARB                                                = 0x88C7
	OBJECT_SUBTYPE_ARB                                         = 0x8B4F
	SAMPLER_CUBE_SHADOW_EXT                                    = 0x8DC5
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	STENCIL                                                    = 0x1802
	POINT_SIZE_MIN_ARB                                         = 0x8126
	RESTART_SUN                                                = 0x0001
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	NUM_COMPRESSED_TEXTURE_FORMATS_ARB                         = 0x86A2
	FLOAT_VEC3                                                 = 0x8B51
	COLOR_ATTACHMENT6_NV                                       = 0x8CE6
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	VERTEX_SHADER_BIT                                          = 0x00000001
	SPRITE_TRANSLATION_SGIX                                    = 0x814B
	COLOR_COMPONENTS                                           = 0x8283
	RGB_FLOAT16_APPLE                                          = 0x881B
	REG_29_ATI                                                 = 0x893E
	PIXEL_TILE_GRID_HEIGHT_SGIX                                = 0x8143
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	IMAGE_BINDING_NAME                                         = 0x8F3A
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	LINE_SMOOTH                                                = 0x0B20
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	MAX_FRAGMENT_LIGHTS_SGIX                                   = 0x8404
	TEXTURE_RANGE_POINTER_APPLE                                = 0x85B8
	DOT_PRODUCT_DEPTH_REPLACE_NV                               = 0x86ED
	DRAW_BUFFER7_ATI                                           = 0x882C
	DRAW_BUFFER10_ATI                                          = 0x882F
	MAX_CONVOLUTION_HEIGHT_EXT                                 = 0x801B
	TEXTURE_LUMINANCE_TYPE_ARB                                 = 0x8C14
	FRACTIONAL_ODD                                             = 0x8E7B
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	MIPMAP                                                     = 0x8293
	TEXTURE1_ARB                                               = 0x84C1
	TEXTURE_INTENSITY_TYPE_ARB                                 = 0x8C15
	UNSIGNED_INT_SAMPLER_3D_EXT                                = 0x8DD3
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	DYNAMIC_ATI                                                = 0x8761
	MAX_PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                   = 0x88B3
	FRAMEBUFFER_BINDING_ANGLE                                  = 0x8CA6
	RGBA32UI                                                   = 0x8D70
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	ARRAY_ELEMENT_LOCK_COUNT_EXT                               = 0x81A9
	MATRIX7_NV                                                 = 0x8637
	PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                        = 0x8808
	DEPTH24_STENCIL8                                           = 0x88F0
	RENDERBUFFER_WIDTH_EXT                                     = 0x8D42
	POINTS                                                     = 0x0000
	POLYGON_OFFSET_LINE                                        = 0x2A02
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	NORMAL_ARRAY_BUFFER_BINDING_ARB                            = 0x8897
	UNIFORM_TYPE                                               = 0x8A37
	COLOR_ATTACHMENT8                                          = 0x8CE8
	COMPRESSED_SIGNED_RED_RGTC1_EXT                            = 0x8DBC
	RGBA_SNORM                                                 = 0x8F93
	EXT_vertex_array                                           = 1
	BUFFER_UPDATE_BARRIER_BIT_EXT                              = 0x00000200
	STENCIL_FAIL                                               = 0x0B94
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	VIEW_CLASS_48_BITS                                         = 0x82C7
	COLOR_ARRAY_PARALLEL_POINTERS_INTEL                        = 0x83F7
	VERTEX_ARRAY_RANGE_LENGTH_NV                               = 0x851E
	OFFSET_PROJECTIVE_TEXTURE_2D_SCALE_NV                      = 0x8851
	POINT_SIZE_ARRAY_POINTER_OES                               = 0x898C
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	ALLOW_DRAW_FRG_HINT_PGI                                    = 0x1A210
	MATRIX_MODE                                                = 0x0BA0
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	MAX_TEXTURE_IMAGE_UNITS_ARB                                = 0x8872
	PERFMON_RESULT_AMD                                         = 0x8BC6
	READ_FRAMEBUFFER_BINDING_EXT                               = 0x8CAA
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	SGIX_polynomial_ffd                                        = 1
	SCISSOR_TEST                                               = 0x0C11
	ACCUM_BLUE_BITS                                            = 0x0D5A
	INVERTED_SCREEN_W_REND                                     = 0x8491
	OBJECT_ACTIVE_UNIFORMS_ARB                                 = 0x8B86
	SHADE_MODEL                                                = 0x0B54
	CONSTANT_COLOR                                             = 0x8001
	COLOR_TABLE_RED_SIZE_SGI                                   = 0x80DA
	TEXTURE21_ARB                                              = 0x84D5
	ATTRIB_ARRAY_SIZE_NV                                       = 0x8623
	SECONDARY_INTERPOLATOR_ATI                                 = 0x896D
	COLOR_ATTACHMENT11_NV                                      = 0x8CEB
	DOUBLE_MAT4x2_EXT                                          = 0x8F4D
	INDEX_ARRAY_POINTER                                        = 0x8091
	BLEND_SRC_ALPHA                                            = 0x80CB
	VERTEX_WEIGHT_ARRAY_EXT                                    = 0x850C
	QUAD_MESH_SUN                                              = 0x8614
	LINE_STIPPLE                                               = 0x0B24
	HALF_BIT_ATI                                               = 0x00000008
	DRAW_FRAMEBUFFER_BINDING_EXT                               = 0x8CA6
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	OPERAND0_RGB_ARB                                           = 0x8590
	VERTEX_PROGRAM_BINDING_NV                                  = 0x864A
	OUTPUT_TEXTURE_COORD26_EXT                                 = 0x87B7
	VIRTUAL_PAGE_SIZE_Z_AMD                                    = 0x9197
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	VARIABLE_A_NV                                              = 0x8523
	MALI_SHADER_BINARY_ARM                                     = 0x8F60
	BITMAP_TOKEN                                               = 0x0704
	BOOL_VEC3                                                  = 0x8B58
	MALI_PROGRAM_BINARY_ARM                                    = 0x8F61
	MATRIX14_ARB                                               = 0x88CE
	BUFFER_OBJECT_EXT                                          = 0x9151
	UNSIGNED_SHORT_4_4_4_4_REV_EXT                             = 0x8365
	TEXTURE_WRAP_R_EXT                                         = 0x8072
	VERTEX_STREAM4_ATI                                         = 0x8770
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV                         = 0x8C4E
	ALPHA_MIN_CLAMP_INGR                                       = 0x8563
	SAMPLER_CUBE                                               = 0x8B60
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	EXT_texture3D                                              = 1
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	OPERAND2_ALPHA                                             = 0x859A
	WEIGHT_SUM_UNITY_ARB                                       = 0x86A6
	OP_SET_LT_EXT                                              = 0x878D
	MAX_PROGRAM_OUTPUT_VERTICES_NV                             = 0x8C27
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	FULL_STIPPLE_HINT_PGI                                      = 0x1A219
	TRACE_PRIMITIVES_BIT_MESA                                  = 0x0002
	REG_19_ATI                                                 = 0x8934
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	IMAGE_1D_EXT                                               = 0x904C
	SGIS_point_line_texgen                                     = 1
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	MODELVIEW23_ARB                                            = 0x8737
	MATRIX23_ARB                                               = 0x88D7
	OBJECT_LINK_STATUS_ARB                                     = 0x8B82
	SPOT_CUTOFF                                                = 0x1206
	SELECT                                                     = 0x1C02
	PN_TRIANGLES_POINT_MODE_LINEAR_ATI                         = 0x87F5
	INTENSITY32UI_EXT                                          = 0x8D73
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	PATH_JOIN_STYLE_NV                                         = 0x9079
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	COLOR_TABLE_LUMINANCE_SIZE                                 = 0x80DE
	COMPRESSED_RG                                              = 0x8226
	MATRIX_INDEX_ARRAY_TYPE_OES                                = 0x8847
	ALPHA8UI_EXT                                               = 0x8D7E
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	SMOOTH                                                     = 0x1D01
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	DEBUG_CATEGORY_OTHER_AMD                                   = 0x9150
	MAX_OPTIMIZED_VERTEX_SHADER_VARIANTS_EXT                   = 0x87CB
	MATRIX10_ARB                                               = 0x88CA
	FLOAT16_VEC2_NV                                            = 0x8FF9
	NUM_EXTENSIONS                                             = 0x821D
	SRC2_ALPHA                                                 = 0x858A
	ALPHA16F_ARB                                               = 0x881C
	OBJECT_COMPILE_STATUS_ARB                                  = 0x8B81
	SLUMINANCE                                                 = 0x8C46
	MAX_PROGRAM_GENERIC_ATTRIBS_NV                             = 0x8DA5
	SYNC_FENCE_APPLE                                           = 0x9116
	FUNC_REVERSE_SUBTRACT_EXT                                  = 0x800B
	MAX_TEXTURE_COORDS_ARB                                     = 0x8871
	ORDER                                                      = 0x0A01
	COLOR_ARRAY_SIZE                                           = 0x8081
	DEPTH_COMPONENT32_OES                                      = 0x81A7
	RG16                                                       = 0x822C
	STATIC_ATI                                                 = 0x8760
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	ADJACENT_PAIRS_NV                                          = 0x90AE
	BLEND_DST_RGB_EXT                                          = 0x80C8
	RG8UI                                                      = 0x8238
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	FRAGMENT_LIGHT3_SGIX                                       = 0x840F
	TEXTURE25_ARB                                              = 0x84D9
	REG_9_ATI                                                  = 0x892A
	GREEN                                                      = 0x1904
	DEPTH_STENCIL                                              = 0x84F9
	COMPRESSED_LUMINANCE_ALPHA_LATC2_EXT                       = 0x8C72
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_EXT                  = 0x8C88
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	NAME_LENGTH                                                = 0x92F9
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	ACCUM_GREEN_BITS                                           = 0x0D59
	OBJECT_POINT_SGIS                                          = 0x81F5
	PRIMARY_COLOR_NV                                           = 0x852C
	INTENSITY16F_ARB                                           = 0x881D
	DRAW_BUFFER10_ARB                                          = 0x882F
	CONDITION_SATISFIED_APPLE                                  = 0x911C
	COLOR_ARRAY_LIST_STRIDE_IBM                                = 103082
	TEXTURE_CUBE_MAP_NEGATIVE_Y_OES                            = 0x8518
	COMBINER_INPUT_NV                                          = 0x8542
	VERTEX_ATTRIB_ARRAY_STRIDE_ARB                             = 0x8624
	PALETTE8_RGB8_OES                                          = 0x8B95
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	DEPTH_STENCIL_MESA                                         = 0x8750
	IMAGE_SCALE_X_HP                                           = 0x8155
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	VERTEX_STREAM6_ATI                                         = 0x8772
	LUMINANCE16_ALPHA16_SNORM                                  = 0x901A
	RED_BITS                                                   = 0x0D52
	VIEWPORT                                                   = 0x0BA2
	COLOR_CLEAR_VALUE                                          = 0x0C22
	SHORT                                                      = 0x1402
	INCR_WRAP_OES                                              = 0x8507
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	RENDERBUFFER_INTERNAL_FORMAT_EXT                           = 0x8D44
	FIELDS_NV                                                  = 0x8E27
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	TEXTURE_BINDING_RECTANGLE_ARB                              = 0x84F6
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	SECONDARY_COLOR_ARRAY_POINTER_EXT                          = 0x845D
	ALPHA12_EXT                                                = 0x803D
	MATRIX_INDEX_ARRAY_OES                                     = 0x8844
	CON_11_ATI                                                 = 0x894C
	FRAGMENT_PROGRAM_CALLBACK_MESA                             = 0x8BB1
	OUT_OF_MEMORY                                              = 0x0505
	RG32F                                                      = 0x8230
	FRAMEBUFFER_BINDING_OES                                    = 0x8CA6
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	SAMPLE_ALPHA_TO_MASK_SGIS                                  = 0x809E
	POST_COLOR_MATRIX_BLUE_SCALE                               = 0x80B6
	VERTEX_ATTRIB_MAP2_ORDER_APPLE                             = 0x8A08
	COLOR_ATTACHMENT1                                          = 0x8CE1
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	DOUBLE_VEC4_EXT                                            = 0x8FFE
	SMALL_CW_ARC_TO_NV                                         = 0x14
	INT_IMAGE_3D_EXT                                           = 0x9059
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	SAMPLES_EXT                                                = 0x80A9
	COMBINER4_NV                                               = 0x8554
	RENDERBUFFER_EXT                                           = 0x8D41
	SIGNED_NORMALIZED                                          = 0x8F9C
	SGIS_multisample                                           = 1
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	MAX_VIEWPORTS                                              = 0x825B
	COMPILE_STATUS                                             = 0x8B81
	MATRIX_INDEX_ARRAY_BUFFER_BINDING_OES                      = 0x8B9E
	VIDEO_CAPTURE_SURFACE_ORIGIN_NV                            = 0x903C
	LUMINANCE4_ALPHA4_EXT                                      = 0x8043
	SOURCE0_ALPHA_EXT                                          = 0x8588
	INVERSE_TRANSPOSE_NV                                       = 0x862D
	CULL_MODES_NV                                              = 0x86E0
	BUFFER_MAP_POINTER                                         = 0x88BD
	RG_SNORM                                                   = 0x8F91
	DEBUG_SEVERITY_MEDIUM_AMD                                  = 0x9147
	TEXTURE_WRAP_Q_SGIS                                        = 0x8137
	PN_TRIANGLES_ATI                                           = 0x87F0
	DOUBLE_MAT4x2                                              = 0x8F4D
	UNSIGNED_BYTE_3_3_2_EXT                                    = 0x8032
	R8UI                                                       = 0x8232
	OFFSET_TEXTURE_2D_SCALE_NV                                 = 0x86E2
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	SGIX_interlace                                             = 1
	MIRRORED_REPEAT_IBM                                        = 0x8370
	FRAMEBUFFER_EXT                                            = 0x8D40
	UNSIGNED_INT_SAMPLER_2D_RECT_EXT                           = 0x8DD5
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	FRAGMENT_LIGHTING_SGIX                                     = 0x8400
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	MAX_TEXTURE_BUFFER_SIZE_EXT                                = 0x8C2B
	T2F_IUI_N3F_V3F_EXT                                        = 0x81B4
	R32F                                                       = 0x822E
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	MODELVIEW1_ARB                                             = 0x850A
	TIMEOUT_EXPIRED_APPLE                                      = 0x911B
	FENCE_APPLE                                                = 0x8A0B
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                         = 0x9061
	UNSIGNALED                                                 = 0x9118
	INTENSITY12                                                = 0x804C
	CLIP_DISTANCE7                                             = 0x3007
	VERTEX_TEXTURE                                             = 0x829B
	ATTACHED_SHADERS                                           = 0x8B85
	ACTIVE_VARYING_MAX_LENGTH_NV                               = 0x8C82
	STENCIL_INDEX8_EXT                                         = 0x8D48
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	TEXTURE_MAX_LOD                                            = 0x813B
	CON_26_ATI                                                 = 0x895B
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	VERTEX_ATTRIB_ARRAY8_NV                                    = 0x8658
	MODELVIEW31_ARB                                            = 0x873F
	INTERLACE_OML                                              = 0x8980
	UNSIGNED_NORMALIZED_EXT                                    = 0x8C17
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	INVERT                                                     = 0x150A
	LUMINANCE8_ALPHA8_EXT                                      = 0x8045
	MAX_OPTIMIZED_VERTEX_SHADER_LOCALS_EXT                     = 0x87CE
	CON_9_ATI                                                  = 0x894A
	INDEX                                                      = 0x8222
	MAP1_TANGENT_EXT                                           = 0x8444
	COMBINER_COMPONENT_USAGE_NV                                = 0x8544
	VIDEO_BUFFER_BINDING_NV                                    = 0x9021
	VERSION_3_0                                                = 1
	TEXTURE_3D_OES                                             = 0x806F
	UNSIGNED_NORMALIZED_ARB                                    = 0x8C17
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	GL_422_REV_AVERAGE_EXT                                     = 0x80CF
	DISTANCE_ATTENUATION_SGIS                                  = 0x8129
	TEXTURE_LOD_BIAS_S_SGIX                                    = 0x818E
	TRANSPOSE_CURRENT_MATRIX_ARB                               = 0x88B7
	TRANSFORM_FEEDBACK_BUFFER_MODE_NV                          = 0x8C7F
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5E
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	PROGRAM_STRING_ARB                                         = 0x8628
	MATRIX29_ARB                                               = 0x88DD
	COLOR_ATTACHMENT2_EXT                                      = 0x8CE2
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	ONE_MINUS_CONSTANT_ALPHA_EXT                               = 0x8004
	DOUBLE_VEC3                                                = 0x8FFD
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	PROXY_TEXTURE_3D_EXT                                       = 0x8070
	ACTIVE_STENCIL_FACE_EXT                                    = 0x8911
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	VIDEO_CAPTURE_TO_422_SUPPORTED_NV                          = 0x9026
	COMPRESSED_R11_EAC                                         = 0x9270
	MAP2_VERTEX_3                                              = 0x0DB7
	STENCIL_BACK_FAIL                                          = 0x8801
	COLOR_ATTACHMENT4_NV                                       = 0x8CE4
	SGIX_async                                                 = 1
	CON_8_ATI                                                  = 0x8949
	QUERY_NO_WAIT                                              = 0x8E14
	TEXTURE_CUBE_MAP_EXT                                       = 0x8513
	SLIM12S_SGIX                                               = 0x831F
	TEXTURE_CUBE_MAP                                           = 0x8513
	TEXTURE_DT_SIZE_NV                                         = 0x871E
	REG_23_ATI                                                 = 0x8938
	SLUMINANCE8_EXT                                            = 0x8C47
	SKIP_COMPONENTS2_NV                                        = -5
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	COMPRESSED_RGBA_S3TC_DXT3_ANGLE                            = 0x83F2
	R1UI_T2F_V3F_SUN                                           = 0x85C9
	MAP1_VERTEX_ATTRIB2_4_NV                                   = 0x8662
	MAX_VARYING_COMPONENTS_EXT                                 = 0x8B4B
	CLIP_DISTANCE_NV                                           = 0x8C7A
	MAX_PROGRAM_PARAMETER_BUFFER_SIZE_NV                       = 0x8DA1
	VERSION_2_1                                                = 1
	REDUCE_EXT                                                 = 0x8016
	TRACK_MATRIX_TRANSFORM_NV                                  = 0x8649
	DRAW_BUFFER10_NV                                           = 0x882F
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	SGIX_calligraphic_fragment                                 = 1
	TEXTURE8_ARB                                               = 0x84C8
	UNIFORM_BUFFER                                             = 0x8A11
	MAX_NAME_LENGTH                                            = 0x92F6
	HISTOGRAM_FORMAT_EXT                                       = 0x8027
	TEXTURE_COORD_ARRAY_STRIDE_EXT                             = 0x808A
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	COLOR_ATTACHMENT7_NV                                       = 0x8CE7
	MAX_IMAGE_UNITS_EXT                                        = 0x8F38
	SLUMINANCE8_NV                                             = 0x8C47
	RENDER_MODE                                                = 0x0C40
	RGB16                                                      = 0x8054
	COLOR_ARRAY_STRIDE_EXT                                     = 0x8083
	R16F_EXT                                                   = 0x822D
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	VIDEO_CAPTURE_FIELD_LOWER_HEIGHT_NV                        = 0x903B
	SLUMINANCE_NV                                              = 0x8C46
	INDEX_BIT_PGI                                              = 0x00080000
	DOT3_ATI                                                   = 0x8966
	TEXTURE_COVERAGE_SAMPLES_NV                                = 0x9045
	PACK_RESAMPLE_OML                                          = 0x8984
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	READ_BUFFER                                                = 0x0C02
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	FRAGMENT_MATERIAL_EXT                                      = 0x8349
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	TEXTURE_WIDTH_QCOM                                         = 0x8BD2
	STENCIL_INDEX4                                             = 0x8D47
	PATH_FILL_MODE_NV                                          = 0x9080
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI                            = 0x00200000
	GREATER                                                    = 0x0204
	POINT_SPRITE                                               = 0x8861
	RG8I                                                       = 0x8237
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB                            = 0x8517
	VECTOR_EXT                                                 = 0x87BF
	MAX_FRAGMENT_BINDABLE_UNIFORMS_EXT                         = 0x8DE3
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	C3F_V3F                                                    = 0x2A24
	COLOR_ATTACHMENT7                                          = 0x8CE7
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	CW                                                         = 0x0900
	STENCIL_BACK_FUNC_ATI                                      = 0x8800
	MATRIX11_ARB                                               = 0x88CB
	BLUE                                                       = 0x1905
	REFLECTION_MAP_OES                                         = 0x8512
	SCALE_BY_ONE_HALF_NV                                       = 0x8540
	MODELVIEW9_ARB                                             = 0x8729
	LUMINANCE_FLOAT16_APPLE                                    = 0x881E
	SUBSAMPLE_DISTANCE_AMD                                     = 0x883F
	INT_IMAGE_2D_RECT                                          = 0x905A
	MINMAX                                                     = 0x802E
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	PIXEL_PACK_BUFFER                                          = 0x88EB
	UNSIGNED_INT_SAMPLER_BUFFER_EXT                            = 0x8DD8
	CURRENT_BINORMAL_EXT                                       = 0x843C
	PIXEL_TILE_WIDTH_SGIX                                      = 0x8140
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	NORMAL_MAP_EXT                                             = 0x8511
	ARRAY_STRIDE                                               = 0x92FE
	SPRITE_AXIS_SGIX                                           = 0x814A
	SAMPLER                                                    = 0x82E6
	TEXTURE17                                                  = 0x84D1
	FENCE_CONDITION_NV                                         = 0x84F4
	SOURCE1_RGB_EXT                                            = 0x8581
	PROGRAM_POINT_SIZE_EXT                                     = 0x8642
	SAMPLER_2D_RECT                                            = 0x8B63
	POLYGON_OFFSET_POINT                                       = 0x2A01
	CONVOLUTION_2D_EXT                                         = 0x8011
	PIXEL_GROUP_COLOR_SGIS                                     = 0x8356
	TEXTURE14_ARB                                              = 0x84CE
	IDENTITY_NV                                                = 0x862A
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	MODELVIEW0_ARB                                             = 0x1700
	MODELVIEW21_ARB                                            = 0x8735
	RENDERBUFFER_WIDTH                                         = 0x8D42
	PATCHES                                                    = 0x000E
	POST_COLOR_MATRIX_BLUE_BIAS                                = 0x80BA
	FRAMEBUFFER_UNDEFINED_OES                                  = 0x8219
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	TEXTURE_COMPRESSED                                         = 0x86A1
	SOURCE1_ALPHA                                              = 0x8589
	MAP2_VERTEX_ATTRIB0_4_NV                                   = 0x8670
	READ_ONLY                                                  = 0x88B8
	REG_26_ATI                                                 = 0x893B
	FRAGMENT_SHADER                                            = 0x8B30
	UNSIGNED_SHORT_5_6_5_EXT                                   = 0x8363
	PROXY_COLOR_TABLE_SGI                                      = 0x80D3
	INVARIANT_DATATYPE_EXT                                     = 0x87EB
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	MAX_GEOMETRY_BINDABLE_UNIFORMS_EXT                         = 0x8DE4
	INT_IMAGE_2D_RECT_EXT                                      = 0x905A
	PROGRAM_OUTPUT                                             = 0x92E4
	ALWAYS_FAST_HINT_PGI                                       = 0x1A20C
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	CLAMP_TO_BORDER                                            = 0x812D
	TEXCOORD4_BIT_PGI                                          = 0x80000000
	TEXTURE_RED_SIZE                                           = 0x805C
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	MAX_PALETTE_MATRICES_OES                                   = 0x8842
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	COLOR_SAMPLES_NV                                           = 0x8E20
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	ONE                                                        = 1
	LINEAR_SHARPEN_SGIS                                        = 0x80AD
	UNPACK_SKIP_VOLUMES_SGIS                                   = 0x8132
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	SOURCE0_RGB_ARB                                            = 0x8580
	DRAW_BUFFER11_ARB                                          = 0x8830
	POST_CONVOLUTION_BLUE_SCALE                                = 0x801E
	VARIABLE_B_NV                                              = 0x8524
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	SGI_color_table                                            = 1
	MAX_CLIP_PLANES                                            = 0x0D32
	DRAW_BUFFER10                                              = 0x882F
	FRAMEBUFFER_INCOMPLETE_FORMATS_OES                         = 0x8CDA
	DOUBLE_MAT4x3_EXT                                          = 0x8F4E
	SYNC_FENCE                                                 = 0x9116
	MAX_SPARSE_3D_TEXTURE_SIZE_AMD                             = 0x9199
	EXT_shared_texture_palette                                 = 1
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV                      = 0x8D56
	INT_SAMPLER_2D                                             = 0x8DCA
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	TRIANGLES                                                  = 0x0004
	MATRIX1_ARB                                                = 0x88C1
	COEFF                                                      = 0x0A00
	BYTE                                                       = 0x1400
	TEXTURE_BLUE_SIZE_EXT                                      = 0x805E
	PROXY_TEXTURE_1D_EXT                                       = 0x8063
	COLOR_ARRAY_TYPE_EXT                                       = 0x8082
	COLOR_INDEX2_EXT                                           = 0x80E3
	TEXTURE_MIN_LOD_SGIS                                       = 0x813A
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	UNSIGNED_INT_S8_S8_8_8_NV                                  = 0x86DA
	BUFFER_MAPPED                                              = 0x88BC
	MULTISAMPLE_BIT_EXT                                        = 0x20000000
	DRAW_BUFFER                                                = 0x0C01
	MAP_STENCIL                                                = 0x0D11
	MINMAX_FORMAT_EXT                                          = 0x802F
	PIXEL_TEX_GEN_Q_CEILING_SGIX                               = 0x8184
	FOG_COORDINATE_ARRAY_POINTER_EXT                           = 0x8456
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	RGB5_A1_OES                                                = 0x8057
	TEXTURE_CUBE_MAP_POSITIVE_X_EXT                            = 0x8515
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	ALPHA_BITS                                                 = 0x0D55
	GL_4PASS_0_SGIS                                            = 0x80A4
	POINT_FADE_THRESHOLD_SIZE_ARB                              = 0x8128
	RG16F_EXT                                                  = 0x822F
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	TEXTURE_LIGHT_EXT                                          = 0x8350
	MODULATE_ADD_ATI                                           = 0x8744
	OUTPUT_TEXTURE_COORD27_EXT                                 = 0x87B8
	DRAW_BUFFER11_ATI                                          = 0x8830
	PATH_FILL_MASK_NV                                          = 0x9081
	AUX1                                                       = 0x040A
	POST_COLOR_MATRIX_GREEN_SCALE_SGI                          = 0x80B5
	TEXTURE26                                                  = 0x84DA
	PROGRAM_ERROR_POSITION_NV                                  = 0x864B
	VERTEX_ATTRIB_ARRAY7_NV                                    = 0x8657
	SAMPLE_POSITION_NV                                         = 0x8E50
	SGIS_detail_texture                                        = 1
	CUBIC_HP                                                   = 0x815F
	MAX_PROGRAM_PARAMETERS_ARB                                 = 0x88A9
	TEXTURE_BINDING_CUBE_MAP_EXT                               = 0x8514
	HISTOGRAM_RED_SIZE                                         = 0x8028
	POINT_SIZE_MIN_EXT                                         = 0x8126
	MAX_DEPTH                                                  = 0x8280
	RENDERBUFFER_BINDING_OES                                   = 0x8CA7
	VIDEO_COLOR_CONVERSION_MAX_NV                              = 0x902A
	VERTEX_DATA_HINT_PGI                                       = 0x1A22A
	TEXTURE_WIDTH                                              = 0x1000
	DEPTH_PASS_INSTRUMENT_MAX_SGIX                             = 0x8312
	MAP_ATTRIB_V_ORDER_NV                                      = 0x86C4
	PROGRAM_PARAMETERS_ARB                                     = 0x88A8
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	RGB10                                                      = 0x8052
	TEXTURE_MAX_LOD_SGIS                                       = 0x813B
	TEXTURE24_ARB                                              = 0x84D8
	VERTEX_SHADER                                              = 0x8B31
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	STENCIL_ATTACHMENT_OES                                     = 0x8D20
	COVERAGE_BUFFER_BIT_NV                                     = 0x00008000
	POST_COLOR_MATRIX_BLUE_BIAS_SGI                            = 0x80BA
	FILTER4_SGIS                                               = 0x8146
	FOG_OFFSET_VALUE_SGIX                                      = 0x8199
	VERTEX_STREAM3_ATI                                         = 0x876F
	DRAW_BUFFER4_NV                                            = 0x8829
	FRACTIONAL_EVEN                                            = 0x8E7C
	QUERY                                                      = 0x82E3
	SIGNED_RGB_NV                                              = 0x86FE
	CON_7_ATI                                                  = 0x8948
	TEXTURE_SWIZZLE_RGBA_EXT                                   = 0x8E46
	INT_IMAGE_3D                                               = 0x9059
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	UNSIGNED_INVERT_NV                                         = 0x8537
	TEXTURE_COMPARE_MODE_ARB                                   = 0x884C
	READ_FRAMEBUFFER_BINDING_NV                                = 0x8CAA
	FOG_INDEX                                                  = 0x0B61
	VERTEX_ATTRIB_ARRAY15_NV                                   = 0x865F
	OP_MADD_EXT                                                = 0x8788
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_EXT                     = 0x8CD0
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	STENCIL_BUFFER_BIT                                         = 0x00000400
	POST_CONVOLUTION_BLUE_BIAS_EXT                             = 0x8022
	EDGE_FLAG_ARRAY_STRIDE_EXT                                 = 0x808C
	OP_CROSS_PRODUCT_EXT                                       = 0x8797
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	CON_19_ATI                                                 = 0x8954
	COUNTER_RANGE_AMD                                          = 0x8BC1
	TIMESTAMP                                                  = 0x8E28
	IMAGE_TRANSFORM_2D_HP                                      = 0x8161
	RGBA32F                                                    = 0x8814
	UTF16_NV                                                   = 0x909B
	RGB32F                                                     = 0x8815
	FLOAT_RG32_NV                                              = 0x8887
	PERFMON_GLOBAL_MODE_QCOM                                   = 0x8FA0
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI                     = 0x80D4
	DEBUG_TYPE_ERROR                                           = 0x824C
	TEXTURE_RANGE_LENGTH_APPLE                                 = 0x85B7
	TRACE_ERRORS_BIT_MESA                                      = 0x0020
	PRESENT_DURATION_NV                                        = 0x8E2B
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	SGIX_async_pixel                                           = 1
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	SAMPLE_BUFFERS_SGIS                                        = 0x80A8
	PIXEL_TEX_GEN_SGIX                                         = 0x8139
	TEXTURE4_ARB                                               = 0x84C4
	PURGEABLE_APPLE                                            = 0x8A1D
	SAMPLER_2D_SHADOW                                          = 0x8B62
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	DEPTH_BIAS                                                 = 0x0D1F
	T2F_C3F_V3F                                                = 0x2A2A
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	TRANSFORM_FEEDBACK_BUFFER_BINDING_EXT                      = 0x8C8F
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	GL_422_REV_EXT                                             = 0x80CD
	ASYNC_DRAW_PIXELS_SGIX                                     = 0x835D
	ALPHA_FLOAT32_APPLE                                        = 0x8816
	GREEN_BITS                                                 = 0x0D53
	VERTEX_ATTRIB_ARRAY_TYPE_ARB                               = 0x8625
	TEXTURE_HI_SIZE_NV                                         = 0x871B
	LUMINANCE_FLOAT16_ATI                                      = 0x881E
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	SGIS_texture4D                                             = 1
	TEXTURE_BORDER                                             = 0x1005
	LIGHT7                                                     = 0x4007
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	INT64_VEC3_NV                                              = 0x8FEA
	COLOR_TABLE_FORMAT_SGI                                     = 0x80D8
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	SAMPLE_POSITION                                            = 0x8E50
	TEXTURE_RENDERBUFFER_NV                                    = 0x8E55
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	DRAW_BUFFER13_ATI                                          = 0x8832
	DITHER                                                     = 0x0BD0
	MODELVIEW5_ARB                                             = 0x8725
	EVAL_VERTEX_ATTRIB12_NV                                    = 0x86D2
	TRANSFORM_FEEDBACK_ATTRIBS_NV                              = 0x8C7E
	RED_SNORM                                                  = 0x8F90
	IMAGE_BUFFER                                               = 0x9051
	SCALED_RESOLVE_NICEST_EXT                                  = 0x90BB
	AND_REVERSE                                                = 0x1502
	Q                                                          = 0x2003
	CONVOLUTION_BORDER_MODE_EXT                                = 0x8013
	VERTEX_ATTRIB_ARRAY_INTEGER_NV                             = 0x88FD
	UNSIGNED_INT_SAMPLER_CUBE_EXT                              = 0x8DD4
	SYNC_X11_FENCE_EXT                                         = 0x90E1
	HISTOGRAM_SINK_EXT                                         = 0x802D
	SAMPLE_BUFFERS                                             = 0x80A8
	IMAGE_ROTATE_ORIGIN_Y_HP                                   = 0x815B
	TEXTURE_MEMORY_LAYOUT_INTEL                                = 0x83FF
	MAX_TEXTURE_UNITS                                          = 0x84E2
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE                            = 0x8A09
	ACTIVE_UNIFORMS                                            = 0x8B86
	HISTOGRAM_RED_SIZE_EXT                                     = 0x8028
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	MAP2_VERTEX_ATTRIB2_4_NV                                   = 0x8672
	MODELVIEW2_ARB                                             = 0x8722
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	RED                                                        = 0x1903
	VARIANT_ARRAY_EXT                                          = 0x87E8
	RECIP_ADD_SIGNED_ALPHA_IMG                                 = 0x8C05
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	NOTEQUAL                                                   = 0x0205
	TEXTURE_DEFORMATION_SGIX                                   = 0x8195
	REFLECTION_MAP_NV                                          = 0x8512
	DSDT_MAG_NV                                                = 0x86F6
	OUTPUT_TEXTURE_COORD9_EXT                                  = 0x87A6
	SAMPLER_3D_OES                                             = 0x8B5F
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	TEXTURE_COMPARE_MODE                                       = 0x884C
	DYNAMIC_DRAW                                               = 0x88E8
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_EXT                     = 0x8DA9
	SYNC_FLUSH_COMMANDS_BIT_APPLE                              = 0x00000001
	SCALE_BY_TWO_NV                                            = 0x853E
	OUTPUT_TEXTURE_COORD6_EXT                                  = 0x87A3
	SGIX_shadow_ambient                                        = 1
	VERTEX_ARRAY_BINDING_APPLE                                 = 0x85B5
	MAP2_VERTEX_ATTRIB7_4_NV                                   = 0x8677
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	MAP2_VERTEX_ATTRIB11_4_NV                                  = 0x867B
	COMPUTE_SHADER_BIT                                         = 0x00000020
	TRANSPOSE_COLOR_MATRIX_ARB                                 = 0x84E6
	MAX_PROGRAM_LOCAL_PARAMETERS_ARB                           = 0x88B4
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	ALREADY_SIGNALED                                           = 0x911A
	OUTPUT_COLOR1_EXT                                          = 0x879C
	RESAMPLE_AVERAGE_OML                                       = 0x8988
	COMPRESSED_RED_GREEN_RGTC2_EXT                             = 0x8DBD
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	VERSION_1_4                                                = 1
	INDEX_SHIFT                                                = 0x0D12
	CLIP_PLANE3                                                = 0x3003
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	TEXTURE_MIN_FILTER                                         = 0x2801
	PRIMITIVE_RESTART_NV                                       = 0x8558
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	IS_PER_PATCH                                               = 0x92E7
	TEXTURE6_ARB                                               = 0x84C6
	ALPHA16I_EXT                                               = 0x8D8A
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT_EXT                        = 0x00000001
	PROXY_HISTOGRAM_EXT                                        = 0x8025
	QUAD_ALPHA8_SGIS                                           = 0x811F
	SPRITE_SGIX                                                = 0x8148
	LIST_PRIORITY_SGIX                                         = 0x8182
	NUM_GENERAL_COMBINERS_NV                                   = 0x854E
	ATOMIC_COUNTER_BARRIER_BIT_EXT                             = 0x00001000
	HALF_APPLE                                                 = 0x140B
	MODELVIEW10_ARB                                            = 0x872A
	MIRROR_CLAMP_TO_EDGE_ATI                                   = 0x8743
	COMPRESSED_SIGNED_LUMINANCE_LATC1_EXT                      = 0x8C71
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	SGIX_pixel_tiles                                           = 1
	POST_COLOR_MATRIX_COLOR_TABLE_SGI                          = 0x80D2
	VERTEX_PRECLIP_SGIX                                        = 0x83EE
	SECONDARY_COLOR_NV                                         = 0x852D
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	DEPTH_CLAMP_NEAR_AMD                                       = 0x901E
	RGB4                                                       = 0x804F
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	ALPHA16_EXT                                                = 0x803E
	RESAMPLE_ZERO_FILL_SGIX                                    = 0x842F
	INCR_WRAP                                                  = 0x8507
	VARIANT_DATATYPE_EXT                                       = 0x87E5
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	FRAMEBUFFER_SRGB_CAPABLE_EXT                               = 0x8DBA
	FUNC_ADD                                                   = 0x8006
	VERTEX_SOURCE_ATI                                          = 0x8774
	MAX_DRAW_BUFFERS                                           = 0x8824
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_OES           = 0x8CD3
	TEXTURE_COORD_ARRAY_SIZE_EXT                               = 0x8088
	FLOAT_RGBA_NV                                              = 0x8883
	RENDERBUFFER_OES                                           = 0x8D41
	INT_SAMPLER_2D_EXT                                         = 0x8DCA
	BLEND_EQUATION                                             = 0x8009
	COLOR_TABLE_INTENSITY_SIZE_SGI                             = 0x80DF
	OUTPUT_VERTEX_EXT                                          = 0x879A
	RGB                                                        = 0x1907
	COLOR_INDEX12_EXT                                          = 0x80E6
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX                           = 0x8187
	DECR_WRAP_EXT                                              = 0x8508
	RGBA_UNSIGNED_DOT_PRODUCT_MAPPING_NV                       = 0x86D9
	PASS_THROUGH_NV                                            = 0x86E6
	CON_4_ATI                                                  = 0x8945
	FRAGMENT_SHADER_ARB                                        = 0x8B30
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_EXT              = 0x8CD7
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	GL_4PASS_0_EXT                                             = 0x80A4
	MIRRORED_REPEAT_ARB                                        = 0x8370
	EVAL_VERTEX_ATTRIB4_NV                                     = 0x86CA
	X_EXT                                                      = 0x87D5
	WRITEONLY_RENDERING_QCOM                                   = 0x8823
	MAX_PROGRAM_SUBROUTINE_NUM_NV                              = 0x8F45
	TEXTURE_BASE_LEVEL                                         = 0x813C
	CLIENT_ACTIVE_TEXTURE_ARB                                  = 0x84E1
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	TEXTURE_COORD_ARRAY_LIST_IBM                               = 103074
	MAP2_VERTEX_ATTRIB4_4_NV                                   = 0x8674
	INT_VEC4                                                   = 0x8B55
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	BLUE_INTEGER_EXT                                           = 0x8D96
	MAX_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5C
	MAX_DEBUG_MESSAGE_LENGTH_AMD                               = 0x9143
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	DST_ALPHA                                                  = 0x0304
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	CLIP_DISTANCE3                                             = 0x3003
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	SIGNED_RGBA_NV                                             = 0x86FB
	INT_SAMPLER_3D                                             = 0x8DCB
	TEXTURE_3D_BINDING_OES                                     = 0x806A
	TEXTURE9                                                   = 0x84C9
	OUTPUT_TEXTURE_COORD29_EXT                                 = 0x87BA
	TRANSFORM_BIT                                              = 0x00001000
	COLOR_TABLE_GREEN_SIZE                                     = 0x80DB
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX                        = 0x8188
	TEXTURE15_ARB                                              = 0x84CF
	MAX_RENDERBUFFER_SIZE_OES                                  = 0x84E8
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	DEPTH_BOUNDS_EXT                                           = 0x8891
	TESS_EVALUATION_SHADER                                     = 0x8E87
	IMAGE_BINDING_ACCESS_EXT                                   = 0x8F3E
	FOG_COORDINATE_ARRAY_EXT                                   = 0x8457
	EDGE_FLAG_ARRAY_EXT                                        = 0x8079
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	TEXTURE18_ARB                                              = 0x84D2
	OFFSET_PROJECTIVE_TEXTURE_2D_NV                            = 0x8850
	CON_2_ATI                                                  = 0x8943
	ENABLE_BIT                                                 = 0x00002000
	GL_4PASS_2_SGIS                                            = 0x80A6
	LO_BIAS_NV                                                 = 0x8715
	QUERY_RESULT                                               = 0x8866
	PRESENT_TIME_NV                                            = 0x8E2A
	FOG_DENSITY                                                = 0x0B62
	COLOR_MATRIX_STACK_DEPTH_SGI                               = 0x80B2
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	SGIX_clipmap                                               = 1
	S                                                          = 0x2000
	NORMAL_ARRAY_EXT                                           = 0x8075
	CLIP_VOLUME_CLIPPING_HINT_EXT                              = 0x80F0
	TEXTURE18                                                  = 0x84D2
	MATRIX5_NV                                                 = 0x8635
	VARIANT_ARRAY_TYPE_EXT                                     = 0x87E7
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	PROGRAM_ALU_INSTRUCTIONS_ARB                               = 0x8805
	PIXEL_COUNT_AVAILABLE_NV                                   = 0x8867
	CON_23_ATI                                                 = 0x8958
	NUM_FRAGMENT_REGISTERS_ATI                                 = 0x896E
	WEIGHT_ARRAY_ARB                                           = 0x86AD
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	MAT_AMBIENT_BIT_PGI                                        = 0x00100000
	INTENSITY8_EXT                                             = 0x804B
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	LUMINANCE_ALPHA_FLOAT32_APPLE                              = 0x8819
	INTENSITY_FLOAT16_APPLE                                    = 0x881D
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	STENCIL_INDEX4_OES                                         = 0x8D47
	UTF8_NV                                                    = 0x909A
	SUBTRACT_ARB                                               = 0x84E7
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB                            = 0x8516
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	BLEND_SRC_ALPHA_EXT                                        = 0x80CB
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	TEXTURE_SAMPLES                                            = 0x9106
	STENCIL_TEST                                               = 0x0B90
	UNSIGNED_SHORT_1_15_REV_MESA                               = 0x8754
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	MAP1_COLOR_4                                               = 0x0D90
	MODELVIEW0_EXT                                             = 0x1700
	VERTEX_ARRAY_COUNT_EXT                                     = 0x807D
	SRC2_RGB                                                   = 0x8582
	REG_24_ATI                                                 = 0x8939
	COVERAGE_ALL_FRAGMENTS_NV                                  = 0x8ED5
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	SHININESS                                                  = 0x1601
	COMPRESSED_ALPHA                                           = 0x84E9
	CONSTANT                                                   = 0x8576
	PROXY_TEXTURE_1D_STACK_MESAX                               = 0x875B
	DRAW_BUFFER2_ARB                                           = 0x8827
	MATRIX0_ARB                                                = 0x88C0
	SAMPLER_2D_SHADOW_ARB                                      = 0x8B62
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS                           = 0x80B0
	COLOR_TABLE_ALPHA_SIZE                                     = 0x80DD
	IMAGE_TRANSLATE_X_HP                                       = 0x8157
	FRAMEBUFFER_UNSUPPORTED_OES                                = 0x8CDD
	PATH_END_CAPS_NV                                           = 0x9076
	COLOR_ARRAY_BUFFER_BINDING_ARB                             = 0x8898
	CONVOLUTION_1D_EXT                                         = 0x8010
	TEXTURE_TOO_LARGE_EXT                                      = 0x8065
	POINT_SIZE_MAX_EXT                                         = 0x8127
	MAX_ASYNC_HISTOGRAM_SGIX                                   = 0x832D
	OPERAND0_RGB                                               = 0x8590
	READ_WRITE                                                 = 0x88BA
	CON_27_ATI                                                 = 0x895C
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SOURCE1_ALPHA_ARB                                          = 0x8589
	OUTPUT_TEXTURE_COORD20_EXT                                 = 0x87B1
	BUFFER_MAP_POINTER_ARB                                     = 0x88BD
	COLOR_ARRAY_POINTER                                        = 0x8090
	COLOR_TABLE_BLUE_SIZE                                      = 0x80DC
	COORD_REPLACE_ARB                                          = 0x8862
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING_ARB                   = 0x889C
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	LUMINANCE12_ALPHA12                                        = 0x8047
	NORMAL_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F6
	PROGRAM_BINARY_LENGTH_OES                                  = 0x8741
	PN_TRIANGLES_NORMAL_MODE_QUADRATIC_ATI                     = 0x87F8
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	LINES_ADJACENCY_EXT                                        = 0x000A
	AVERAGE_HP                                                 = 0x8160
	VERTEX_WEIGHT_ARRAY_SIZE_EXT                               = 0x850D
	OBJECT_BUFFER_USAGE_ATI                                    = 0x8765
	MVP_MATRIX_EXT                                             = 0x87E3
	UNSIGNED_NORMALIZED                                        = 0x8C17
	SLUMINANCE_ALPHA_NV                                        = 0x8C44
	COLOR_ATTACHMENT11                                         = 0x8CEB
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	POLYGON_MODE                                               = 0x0B40
	INDEX_ARRAY_TYPE_EXT                                       = 0x8085
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	OPERAND1_ALPHA                                             = 0x8599
	MODELVIEW16_ARB                                            = 0x8730
	NUM_INSTRUCTIONS_PER_PASS_ATI                              = 0x8971
	MAX_VERTEX_TEXTURE_IMAGE_UNITS_ARB                         = 0x8B4C
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	OBJECT_PLANE                                               = 0x2501
	DOUBLE_VEC4                                                = 0x8FFE
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	AUTO_NORMAL                                                = 0x0D80
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	PROGRAM_INSTRUCTIONS_ARB                                   = 0x88A0
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	LINE_STRIP                                                 = 0x0003
	LOAD                                                       = 0x0101
	FLAT                                                       = 0x1D00
	ALL_COMPLETED_NV                                           = 0x84F2
	INCR_WRAP_EXT                                              = 0x8507
	OUTPUT_TEXTURE_COORD3_EXT                                  = 0x87A0
	COLOR_ATTACHMENT13                                         = 0x8CED
	POST_CONVOLUTION_ALPHA_SCALE_EXT                           = 0x801F
	VERTEX_ARRAY_STRIDE_EXT                                    = 0x807C
	FRAGMENT_LIGHT5_SGIX                                       = 0x8411
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	DYNAMIC_READ                                               = 0x88E9
	DRAW_FRAMEBUFFER_ANGLE                                     = 0x8CA9
	FLOAT                                                      = 0x1406
	XOR                                                        = 0x1506
	TEXTURE_CUBE_MAP_POSITIVE_Z_EXT                            = 0x8519
	WRITE_ONLY_ARB                                             = 0x88B9
	RASTERIZER_DISCARD_EXT                                     = 0x8C89
	LUMINANCE32I_EXT                                           = 0x8D86
	ACTIVE_RESOURCES                                           = 0x92F5
	SGIX_async_histogram                                       = 1
	CLAMP_VERTEX_COLOR_ARB                                     = 0x891A
	LINK_STATUS                                                = 0x8B82
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_NV              = 0x8C80
	QUAD_STRIP                                                 = 0x0008
	KEEP                                                       = 0x1E00
	PROGRAM_UNDER_NATIVE_LIMITS_ARB                            = 0x88B6
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	EXTERNAL_VIRTUAL_MEMORY_BUFFER_AMD                         = 0x9160
	CLIP_PLANE2                                                = 0x3002
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	WEIGHT_ARRAY_SIZE_ARB                                      = 0x86AB
	W_EXT                                                      = 0x87D8
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	PROXY_TEXTURE_1D_ARRAY_EXT                                 = 0x8C19
	EMISSION                                                   = 0x1600
	VERTEX_ARRAY_EXT                                           = 0x8074
	SAMPLE_MASK_SGIS                                           = 0x80A0
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX                            = 0x8175
	MAP1_VERTEX_ATTRIB10_4_NV                                  = 0x866A
	EVAL_TRIANGULAR_2D_NV                                      = 0x86C1
	DOT_PRODUCT_CONST_EYE_REFLECT_CUBE_MAP_NV                  = 0x86F3
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	COMPRESSED_RGBA_PVRTC_4BPPV2_IMG                           = 0x9138
	LOCATION_INDEX                                             = 0x930F
	R1UI_N3F_V3F_SUN                                           = 0x85C7
	FRAGMENT_SHADER_ATI                                        = 0x8920
	REG_1_ATI                                                  = 0x8922
	TEXTURE_1D_ARRAY_EXT                                       = 0x8C18
	RGBA16I_EXT                                                = 0x8D88
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	COMPRESSED_RGB_FXT1_3DFX                                   = 0x86B0
	AND_INVERTED                                               = 0x1504
	SYNC_CL_EVENT_ARB                                          = 0x8240
	MODELVIEW27_ARB                                            = 0x873B
	DOT_PRODUCT_TEXTURE_CUBE_MAP_NV                            = 0x86F0
	PROGRAM_FORMAT_ASCII_ARB                                   = 0x8875
	MAX_TEXTURE_SIZE                                           = 0x0D33
	TEXTURE_CUBE_MAP_POSITIVE_Z_OES                            = 0x8519
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	SLIM8U_SGIX                                                = 0x831D
	TEXTURE_RED_TYPE_ARB                                       = 0x8C10
	COLOR_ATTACHMENT12_EXT                                     = 0x8CEC
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	RG_EXT                                                     = 0x8227
	MODELVIEW4_ARB                                             = 0x8724
	FULL_RANGE_EXT                                             = 0x87E1
	RENDERBUFFER_INTERNAL_FORMAT_OES                           = 0x8D44
	SAMPLE_MASK_VALUE                                          = 0x8E52
	TRIANGLES_ADJACENCY                                        = 0x000C
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	EDGEFLAG_BIT_PGI                                           = 0x00040000
	UNSIGNED_INT64_NV                                          = 0x140F
	FRAGMENT_LIGHT0_SGIX                                       = 0x840C
	DRAW_BUFFER5_NV                                            = 0x882A
	IMAGE_BINDING_LAYER_EXT                                    = 0x8F3D
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	PROXY_COLOR_TABLE                                          = 0x80D3
	BIAS_BY_NEGATIVE_ONE_HALF_NV                               = 0x8541
	MODELVIEW3_ARB                                             = 0x8723
	SGX_BINARY_IMG                                             = 0x8C0A
	RENDER                                                     = 0x1C00
	T4F_V4F                                                    = 0x2A28
	ALPHA_MIN_SGIX                                             = 0x8320
	TEXTURE_BINDING_CUBE_MAP_ARB                               = 0x8514
	SOURCE0_RGB_EXT                                            = 0x8580
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	STENCIL_TEST_TWO_SIDE_EXT                                  = 0x8910
	BUFFER_MAP_LENGTH                                          = 0x9120
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	TEXTURE_CUBE_MAP_POSITIVE_Y_OES                            = 0x8517
	VARIANT_EXT                                                = 0x87C1
	TEXTURE_ENV_COLOR                                          = 0x2201
	MAX_COMBINED_TEXTURE_IMAGE_UNITS_ARB                       = 0x8B4D
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	AFFINE_2D_NV                                               = 0x9092
	VERSION_3_1                                                = 1
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	RGB12                                                      = 0x8053
	CONVOLUTION_1D                                             = 0x8010
	UNSIGNED_SHORT_5_6_5_REV_EXT                               = 0x8364
	SPARE1_NV                                                  = 0x852F
	OP_NEGATE_EXT                                              = 0x8783
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_EXT                   = 0x8DE1
	RGB8                                                       = 0x8051
	DEBUG_SOURCE_OTHER                                         = 0x824B
	OFFSET_TEXTURE_BIAS_NV                                     = 0x86E3
	RGBA_FLOAT_MODE_ARB                                        = 0x8820
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	STENCIL_INDEX1_EXT                                         = 0x8D46
	UNSIGNED_INT_VEC3_EXT                                      = 0x8DC7
	SAMPLE_MASK_INVERT_SGIS                                    = 0x80AB
	TRIANGLE_LIST_SUN                                          = 0x81D7
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	OUTPUT_TEXTURE_COORD23_EXT                                 = 0x87B4
	DOUBLE_VEC3_EXT                                            = 0x8FFD
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	VERTEX_ARRAY_RANGE_WITHOUT_FLUSH_NV                        = 0x8533
	ADD_SIGNED                                                 = 0x8574
	UNSIGNED_INT_8_24_REV_MESA                                 = 0x8752
	STRICT_DEPTHFUNC_HINT_PGI                                  = 0x1A216
	FRONT_AND_BACK                                             = 0x0408
	ELEMENT_ARRAY_ATI                                          = 0x8768
	MAX_DRAW_BUFFERS_ARB                                       = 0x8824
	GEOMETRY_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA3
	INDEX_TEST_FUNC_EXT                                        = 0x81B6
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	LO_SCALE_NV                                                = 0x870F
	IMAGE_CUBE_EXT                                             = 0x9050
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	DUAL_TEXTURE_SELECT_SGIS                                   = 0x8124
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB                            = 0x851A
	OP_EXP_BASE_2_EXT                                          = 0x8791
	UNSIGNED_INT_10F_11F_11F_REV_EXT                           = 0x8C3B
	TEXTURE_USAGE_ANGLE                                        = 0x93A2
	EDGE_FLAG_ARRAY                                            = 0x8079
	SAMPLES                                                    = 0x80A9
	QUARTER_BIT_ATI                                            = 0x00000010
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	FLOAT_RGBA_MODE_NV                                         = 0x888E
	TEXTURE_NUM_LEVELS_QCOM                                    = 0x8BD9
	LINE_WIDTH                                                 = 0x0B21
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	LUMINANCE12_EXT                                            = 0x8041
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	FIXED_ONLY_ARB                                             = 0x891D
	REG_0_ATI                                                  = 0x8921
	MEDIUM_INT                                                 = 0x8DF4
	WEIGHT_ARRAY_TYPE_ARB                                      = 0x86A9
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	QUERY_RESULT_AVAILABLE_ARB                                 = 0x8867
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	RG                                                         = 0x8227
	RGBA4_DXT5_S3TC                                            = 0x83A5
	PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                        = 0x8809
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	COLOR_ARRAY_LIST_IBM                                       = 103072
	NEVER                                                      = 0x0200
	RGBA_S3TC                                                  = 0x83A2
	INTERLACE_READ_INGR                                        = 0x8568
	MAP1_VERTEX_ATTRIB14_4_NV                                  = 0x866E
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX                       = 0x817C
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	COLOR_ATTACHMENT1_EXT                                      = 0x8CE1
	LINEAR                                                     = 0x2601
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	LUMINANCE_ALPHA_INTEGER_EXT                                = 0x8D9D
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	RGB16F                                                     = 0x881B
	REG_25_ATI                                                 = 0x893A
	CON_17_ATI                                                 = 0x8952
	COLOR_ATTACHMENT0_EXT                                      = 0x8CE0
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	LESS                                                       = 0x0201
	TEXTURE_COLOR_TABLE_SGI                                    = 0x80BC
	FRAMEZOOM_FACTOR_SGIX                                      = 0x818C
	TEXTURE4                                                   = 0x84C4
	OPERAND3_ALPHA_NV                                          = 0x859B
	MATRIX20_ARB                                               = 0x88D4
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	DOT3_RGBA_EXT                                              = 0x8741
	BUFFER_MAPPED_ARB                                          = 0x88BC
	UNIFORM_BLOCK                                              = 0x92E2
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	MAX_PROGRAM_TEXEL_OFFSET_NV                                = 0x8905
	DEPTH_SCALE                                                = 0x0D1E
	TEXTURE_COMPARE_OPERATOR_SGIX                              = 0x819B
	TEXTURE30_ARB                                              = 0x84DE
	CONSTANT_EXT                                               = 0x8576
	INT_VEC3_ARB                                               = 0x8B54
	IMAGE_BINDING_FORMAT_EXT                                   = 0x906E
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	MAX_CUBE_MAP_TEXTURE_SIZE_OES                              = 0x851C
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	TRUE                                                       = 1
	REPLICATE_BORDER_HP                                        = 0x8153
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	VERTEX23_BIT_PGI                                           = 0x00000004
	CONSTANT_ATTENUATION                                       = 0x1207
	TEXTURE_MAX_LEVEL_SGIS                                     = 0x813D
	EMBOSS_CONSTANT_NV                                         = 0x855E
	MAX_TRACK_MATRICES_NV                                      = 0x862F
	RGB_FLOAT16_ATI                                            = 0x881B
	TEXTURE_BUFFER_DATA_STORE_BINDING_EXT                      = 0x8C2D
	DOUBLE_MAT2x4_EXT                                          = 0x8F4A
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	COMBINER_AB_OUTPUT_NV                                      = 0x854A
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	TEXTURE_CUBE_MAP_NEGATIVE_Z_EXT                            = 0x851A
	INTERPOLATE_ARB                                            = 0x8575
	VERTEX_ATTRIB_ARRAY4_NV                                    = 0x8654
	TRACE_ALL_BITS_MESA                                        = 0xFFFF
	PIXEL_COUNTER_BITS_NV                                      = 0x8864
	SGIX_blend_alpha_minmax                                    = 1
	R16                                                        = 0x822A
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	DRAW_BUFFER15_ATI                                          = 0x8834
	N3F_V3F                                                    = 0x2A25
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	VERTEX_PROGRAM_ARB                                         = 0x8620
	INTENSITY_FLOAT16_ATI                                      = 0x881D
	UNSIGNALED_APPLE                                           = 0x9118
	POST_COLOR_MATRIX_ALPHA_BIAS                               = 0x80BB
	TEXTURE_PRE_SPECULAR_HP                                    = 0x8169
	SAMPLER_2D_ARRAY_EXT                                       = 0x8DC1
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	BACK_LEFT                                                  = 0x0402
	OFFSET_HILO_PROJECTIVE_TEXTURE_2D_NV                       = 0x8856
	UNSIGNED_INT_VEC4_EXT                                      = 0x8DC8
	MAX_GEOMETRY_VARYING_COMPONENTS_EXT                        = 0x8DDD
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	MULTISAMPLE_BIT                                            = 0x20000000
	TEXTURE_WRAP_R                                             = 0x8072
	SPRITE_AXIAL_SGIX                                          = 0x814C
	COMBINE_ALPHA                                              = 0x8572
	RGB16F_EXT                                                 = 0x881B
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	QUAD_LUMINANCE8_SGIS                                       = 0x8121
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	RGBA8I_EXT                                                 = 0x8D8E
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION_EXT               = 0x8E4C
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	TEXTURE_SHADER_NV                                          = 0x86DE
	CURRENT_MATRIX_INDEX_ARB                                   = 0x8845
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	TEXTURE_RED_TYPE                                           = 0x8C10
	TEXTURE3_ARB                                               = 0x84C3
	DSDT8_MAG8_NV                                              = 0x870A
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	RESAMPLE_DECIMATE_OML                                      = 0x8989
	MIN_FRAGMENT_INTERPOLATION_OFFSET_NV                       = 0x8E5B
	NAND                                                       = 0x150E
	LUMINANCE12                                                = 0x8041
	LUMINANCE16_EXT                                            = 0x8042
	MAX_ELEMENTS_VERTICES_EXT                                  = 0x80E8
	POINT_SPRITE_ARB                                           = 0x8861
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	INDEX_ARRAY_POINTER_EXT                                    = 0x8091
	RGB16F_ARB                                                 = 0x881B
	MAX_PROGRAM_ENV_PARAMETERS_ARB                             = 0x88B5
	RGB8I_EXT                                                  = 0x8D8F
	QUERY_BUFFER_AMD                                           = 0x9192
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	INDEX_ARRAY_STRIDE                                         = 0x8086
	UNPACK_IMAGE_HEIGHT_EXT                                    = 0x806E
	POST_COLOR_MATRIX_RED_BIAS_SGI                             = 0x80B8
	BLEND_SRC_RGB                                              = 0x80C9
	STREAM_COPY                                                = 0x88E2
	EXT_copy_texture                                           = 1
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	RGB16I_EXT                                                 = 0x8D89
	FRAMEBUFFER_ATTACHMENT_LAYERED_EXT                         = 0x8DA7
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	FRONT                                                      = 0x0404
	BLEND_EQUATION_RGB                                         = 0x8009
	DYNAMIC_COPY_ARB                                           = 0x88EA
	SYNC_CONDITION                                             = 0x9113
	INDEX_ARRAY                                                = 0x8077
	SHADER_IMAGE_LOAD                                          = 0x82A4
	SAMPLER_3D_ARB                                             = 0x8B5F
	RELATIVE_ARC_TO_NV                                         = 0xFF
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	COLOR_LOGIC_OP                                             = 0x0BF2
	TEXTURE_GREEN_SIZE_EXT                                     = 0x805D
	MAX_VERTEX_ARRAY_RANGE_ELEMENT_NV                          = 0x8520
	TEXTURE_2D_STACK_MESAX                                     = 0x875A
	VERTEX_PROGRAM_PARAMETER_BUFFER_NV                         = 0x8DA2
	UNSIGNED_INT_SAMPLER_1D_EXT                                = 0x8DD1
	NONE                                                       = 0
	GL_3D                                                      = 0x0601
	REPEAT                                                     = 0x2901
	FRAMEBUFFER_OES                                            = 0x8D40
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	TEXTURE_STORAGE_SPARSE_BIT_AMD                             = 0x00000001
	IMAGE_SCALE_Y_HP                                           = 0x8156
	UNDEFINED_VERTEX                                           = 0x8260
	FOG_COORDINATE                                             = 0x8451
	REG_16_ATI                                                 = 0x8931
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	EQUAL                                                      = 0x0202
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	PIXEL_TEX_GEN_Q_ROUND_SGIX                                 = 0x8185
	SOURCE1_ALPHA_EXT                                          = 0x8589
	VERTEX_ATTRIB_ARRAY12_NV                                   = 0x865C
	MAP2_VERTEX_ATTRIB9_4_NV                                   = 0x8679
	TEXTURE_MAG_SIZE_NV                                        = 0x871F
	OUTPUT_TEXTURE_COORD30_EXT                                 = 0x87BB
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	EXT_texture                                                = 1
	COMMAND_BARRIER_BIT                                        = 0x00000040
	RGB16_EXT                                                  = 0x8054
	COMPARE_REF_DEPTH_TO_TEXTURE_EXT                           = 0x884E
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	MAT_EMISSION_BIT_PGI                                       = 0x00800000
	CONVOLUTION_FILTER_SCALE                                   = 0x8014
	INTENSITY4_EXT                                             = 0x804A
	CLAMP_TO_EDGE_SGIS                                         = 0x812F
	OPERAND1_RGB                                               = 0x8591
	RGBA_FLOAT16_ATI                                           = 0x881A
	RGB8I                                                      = 0x8D8F
	RGBA32I_EXT                                                = 0x8D82
	UNSIGNED_INT_VEC2_EXT                                      = 0x8DC6
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	FOG_COORDINATE_ARRAY                                       = 0x8457
	TEXTURE25                                                  = 0x84D9
	MODELVIEW0_STACK_DEPTH_EXT                                 = 0x0BA3
	TEXTURE_NORMAL_EXT                                         = 0x85AF
	VERTEX_PROGRAM_TWO_SIDE_ARB                                = 0x8643
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	TEXTURE2                                                   = 0x84C2
	MIRROR_CLAMP_TO_EDGE_EXT                                   = 0x8743
	PIXEL_PACK_BUFFER_EXT                                      = 0x88EB
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	FULL_SUPPORT                                               = 0x82B7
	DRAW_BUFFER3                                               = 0x8828
	COMPRESSED_RG11_EAC                                        = 0x9272
	BLEND_DST_ALPHA_EXT                                        = 0x80CA
	STATE_RESTORE                                              = 0x8BDC
	COMPUTE_PROGRAM_PARAMETER_BUFFER_NV                        = 0x90FC
	TEXTURE_POST_SPECULAR_HP                                   = 0x8168
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	INDEX_ARRAY_TYPE                                           = 0x8085
	TIME_ELAPSED                                               = 0x88BF
	LAYOUT_LINEAR_INTEL                                        = 1
	BUFFER_USAGE_ARB                                           = 0x8765
	DRAW_BUFFER7                                               = 0x882C
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	TEXTURE_BINDING_RENDERBUFFER_NV                            = 0x8E53
	RELATIVE_MOVE_TO_NV                                        = 0x03
	VERTEX_ARRAY_STORAGE_HINT_APPLE                            = 0x851F
	DRAW_BUFFER5_ATI                                           = 0x882A
	LUMINANCE_ALPHA8I_EXT                                      = 0x8D93
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	FEEDBACK                                                   = 0x1C01
	UNSIGNED_SHORT_15_1_MESA                                   = 0x8753
	RENDER_DIRECT_TO_FRAMEBUFFER_QCOM                          = 0x8FB3
	EXTENSIONS                                                 = 0x1F03
	SURFACE_REGISTERED_NV                                      = 0x86FD
	RGBA_FLOAT32_APPLE                                         = 0x8814
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_EXT                   = 0x8CD2
	LUMINANCE_ALPHA_SNORM                                      = 0x9012
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	BLEND_DST_ALPHA                                            = 0x80CA
	ONE_EXT                                                    = 0x87DE
	DRAW_BUFFER3_ATI                                           = 0x8828
	Z4Y12Z4CB12Z4A12Z4Y12Z4CR12Z4A12_4224_NV                   = 0x9036
	ALPHA_TEST_QCOM                                            = 0x0BC0
	FRAGMENTS_INSTRUMENT_MAX_SGIX                              = 0x8315
	GREEN_MAX_CLAMP_INGR                                       = 0x8565
	OP_ROUND_EXT                                               = 0x8790
	DRAW_BUFFER3_NV                                            = 0x8828
	FLOAT_MAT2x4                                               = 0x8B66
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	POLYGON_TOKEN                                              = 0x0703
	DECR                                                       = 0x1E03
	REDUCE                                                     = 0x8016
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	MAX_ASYNC_READ_PIXELS_SGIX                                 = 0x8361
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE                            = 0x8A05
	QUERY_WAIT_NV                                              = 0x8E13
	UNSIGNED_INT16_NV                                          = 0x8FF0
	SGIX_scalebias_hint                                        = 1
	MAX_SPOT_EXPONENT_NV                                       = 0x8505
	TEXTURE_PRIORITY                                           = 0x8066
	UNSIGNED_SHORT                                             = 0x1403
	NUM_PASSES_ATI                                             = 0x8970
	DOUBLE_MAT4x3                                              = 0x8F4E
	LIST_BIT                                                   = 0x00020000
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	HISTOGRAM_GREEN_SIZE                                       = 0x8029
	MAP1_VERTEX_ATTRIB7_4_NV                                   = 0x8667
	DRAW_BUFFER9_ATI                                           = 0x882E
	TRANSFORM_FEEDBACK_VARYINGS_NV                             = 0x8C83
	VERTEX_STREAM7_ATI                                         = 0x8773
	READ_BUFFER_EXT                                            = 0x0C02
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	ACTIVE_PROGRAM                                             = 0x8259
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	TEXTURE_GATHER                                             = 0x82A2
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	SOURCE0_ALPHA                                              = 0x8588
	VERTEX_ATTRIB_ARRAY13_NV                                   = 0x865D
	SHADER_OBJECT_ARB                                          = 0x8B48
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	AUX_DEPTH_STENCIL_APPLE                                    = 0x8A14
	SHADING_LANGUAGE_VERSION_ARB                               = 0x8B8C
	GL_3_BYTES                                                 = 0x1408
	TEXTURE_MULTI_BUFFER_HINT_SGIX                             = 0x812E
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                        = 0x8162
	ACTIVE_VARIABLES                                           = 0x9305
	LUMINANCE6_ALPHA2_EXT                                      = 0x8044
	MAX_DEFORMATION_ORDER_SGIX                                 = 0x8197
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB                            = 0x8515
	GEOMETRY_PROGRAM_NV                                        = 0x8C26
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	LOGIC_OP                                                   = 0x0BF1
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	VERTEX_ARRAY_RANGE_POINTER_NV                              = 0x8521
	PERTURB_EXT                                                = 0x85AE
	SAMPLER_2D_ARRAY_SHADOW_NV                                 = 0x8DC4
	DETAIL_TEXTURE_2D_BINDING_SGIS                             = 0x8096
	PACK_SKIP_VOLUMES_SGIS                                     = 0x8130
	LINEAR_CLIPMAP_LINEAR_SGIX                                 = 0x8170
	DEPENDENT_RGB_TEXTURE_CUBE_MAP_NV                          = 0x885A
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MAX_VERTEX_HINT_PGI                                        = 0x1A22D
	TEXTURE_RED_SIZE_EXT                                       = 0x805C
	VERTEX_ARRAY_POINTER_EXT                                   = 0x808E
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT                  = 0x8211
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	STATIC_COPY                                                = 0x88E6
	RESAMPLE_REPLICATE_OML                                     = 0x8986
	BUFFER_FLUSHING_UNMAP_APPLE                                = 0x8A13
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	VERSION_1_5                                                = 1
	VERTEX_PROGRAM_NV                                          = 0x8620
	HILO_NV                                                    = 0x86F4
	RGBA_SIGNED_COMPONENTS_EXT                                 = 0x8C3C
	RENDERBUFFER_SAMPLES_ANGLE                                 = 0x8CAB
	MAX_PROGRAM_GENERIC_RESULTS_NV                             = 0x8DA6
	DOT_PRODUCT_TEXTURE_2D_NV                                  = 0x86EE
	GEOMETRY_SHADER_EXT                                        = 0x8DD9
	GCCSO_SHADER_BINARY_FJ                                     = 0x9260
	MATRIX24_ARB                                               = 0x88D8
	RENDERBUFFER_STENCIL_SIZE_OES                              = 0x8D55
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	LUMINANCE12_ALPHA12_EXT                                    = 0x8047
	OPERAND1_RGB_ARB                                           = 0x8591
	DEPTH_BOUNDS_TEST_EXT                                      = 0x8890
	NUM_INPUT_INTERPOLATOR_COMPONENTS_ATI                      = 0x8973
	TESS_GEN_POINT_MODE                                        = 0x8E79
	TEXTURE_4DSIZE_SGIS                                        = 0x8136
	LAST_VERTEX_CONVENTION_EXT                                 = 0x8E4E
	MITER_REVERT_NV                                            = 0x90A7
	TEXTURE_MATRIX                                             = 0x0BA8
	HISTOGRAM_WIDTH                                            = 0x8026
	RGB10_EXT                                                  = 0x8052
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	DOT3_RGBA_ARB                                              = 0x86AF
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	DOUBLE_MAT2x3                                              = 0x8F49
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	MAX_DEBUG_LOGGED_MESSAGES_AMD                              = 0x9144
	POST_COLOR_MATRIX_RED_BIAS                                 = 0x80B8
	TANGENT_ARRAY_EXT                                          = 0x8439
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	CONSTANT_COLOR0_NV                                         = 0x852A
	MAX_VERTEX_SHADER_INSTRUCTIONS_EXT                         = 0x87C5
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	DRAW_BUFFER8_NV                                            = 0x882D
	COLOR_ATTACHMENT5                                          = 0x8CE5
	ARC_TO_NV                                                  = 0xFE
	SGIX_fog_offset                                            = 1
	ASYNC_READ_PIXELS_SGIX                                     = 0x835E
	Z400_BINARY_AMD                                            = 0x8740
	PROGRAM_ADDRESS_REGISTERS_ARB                              = 0x88B0
	YCBYCR8_422_NV                                             = 0x9031
	SYNC_STATUS                                                = 0x9114
	LIST_INDEX                                                 = 0x0B33
	MINMAX_SINK_EXT                                            = 0x8030
	SPARE0_NV                                                  = 0x852E
	SRGB                                                       = 0x8C40
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	CONSTANT_ALPHA_EXT                                         = 0x8003
	BLEND_EQUATION_RGB_EXT                                     = 0x8009
	PIXEL_TEX_GEN_ALPHA_MS_SGIX                                = 0x818A
	DEPTH_COMPONENT24_ARB                                      = 0x81A6
	LINE_QUALITY_HINT_SGIX                                     = 0x835B
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	DEPTH_ATTACHMENT_EXT                                       = 0x8D00
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	POST_COLOR_MATRIX_GREEN_BIAS_SGI                           = 0x80B9
	OBJECT_LINE_SGIS                                           = 0x81F7
	PREVIOUS_TEXTURE_INPUT_NV                                  = 0x86E4
	MAX_PROGRAM_TEMPORARIES_ARB                                = 0x88A5
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	BUFFER_VARIABLE                                            = 0x92E5
	EXT_blend_color                                            = 1
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	PIXEL_SUBSAMPLE_2424_SGIX                                  = 0x85A3
	MAX_VERTEX_UNITS_OES                                       = 0x86A4
	EVAL_VERTEX_ATTRIB6_NV                                     = 0x86CC
	SIGNED_RGBA8_NV                                            = 0x86FC
	NEGATIVE_Z_EXT                                             = 0x87DB
	MATRIX_INDEX_ARRAY_SIZE_ARB                                = 0x8846
	MATRIX2_ARB                                                = 0x88C2
	RGB16UI                                                    = 0x8D77
	IMAGE_2D_MULTISAMPLE_EXT                                   = 0x9055
	CLIP_PLANE0                                                = 0x3000
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	READ_PIXELS                                                = 0x828C
	MAX_PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                     = 0x8337
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	PROJECTION_MATRIX                                          = 0x0BA7
	RGBA16F_ARB                                                = 0x881A
	NORMAL_ARRAY_COUNT_EXT                                     = 0x8080
	TEXTURE23_ARB                                              = 0x84D7
	RGBA16F_EXT                                                = 0x881A
	INT_VEC3                                                   = 0x8B54
	RG16F                                                      = 0x822F
	INTENSITY8UI_EXT                                           = 0x8D7F
	POINT_SMOOTH_HINT                                          = 0x0C51
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	ALPHA8                                                     = 0x803C
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	CURRENT_QUERY                                              = 0x8865
	MAX_SUBROUTINES                                            = 0x8DE7
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	UNSIGNED_INT_IMAGE_BUFFER_EXT                              = 0x9067
	LINE_STRIP_ADJACENCY_EXT                                   = 0x000B
	TABLE_TOO_LARGE                                            = 0x8031
	GL_1PASS_EXT                                               = 0x80A1
	DEPTH_COMPONENT16                                          = 0x81A5
	VERTEX_STATE_PROGRAM_NV                                    = 0x8621
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	SRC0_RGB                                                   = 0x8580
	COLOR_ATTACHMENT0                                          = 0x8CE0
	RGB_INTEGER_EXT                                            = 0x8D98
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	DISCARD_NV                                                 = 0x8530
	EVAL_VERTEX_ATTRIB14_NV                                    = 0x86D4
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	OFFSET                                                     = 0x92FC
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	OP_RECIP_EXT                                               = 0x8794
	FIXED_ONLY                                                 = 0x891D
	ALPHA16_SNORM                                              = 0x9018
	NEXT_VIDEO_CAPTURE_BUFFER_STATUS_NV                        = 0x9025
	SECONDARY_COLOR_ARRAY_LIST_IBM                             = 103077
	FRAGMENT_PROGRAM_BINDING_NV                                = 0x8873
	DRAW_BUFFER15_NV                                           = 0x8834
	UNIFORM_BUFFER_START                                       = 0x8A29
	VIEWPORT_BIT                                               = 0x00000800
	PACK_MAX_COMPRESSED_SIZE_SGIX                              = 0x831B
	VARIABLE_F_NV                                              = 0x8528
	CAVEAT_SUPPORT                                             = 0x82B8
	COMBINE_RGB_EXT                                            = 0x8571
	UNDEFINED_APPLE                                            = 0x8A1C
	SRGB_EXT                                                   = 0x8C40
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	DUAL_INTENSITY4_SGIS                                       = 0x8118
	CONSTANT_BORDER_HP                                         = 0x8151
	MAP2_TANGENT_EXT                                           = 0x8445
	CLAMP_READ_COLOR_ARB                                       = 0x891C
	MAX_FRAGMENT_UNIFORM_COMPONENTS_ARB                        = 0x8B49
	SRGB8                                                      = 0x8C41
	TRANSFORM_FEEDBACK_BUFFER_SIZE_EXT                         = 0x8C85
	LUMINANCE_ALPHA8UI_EXT                                     = 0x8D81
	INT_SAMPLER_BUFFER_EXT                                     = 0x8DD0
	NORMAL_ARRAY_POINTER                                       = 0x808F
	CLEAR                                                      = 0x1500
	T2F_V3F                                                    = 0x2A27
	ZERO_EXT                                                   = 0x87DD
	FLOAT_R_NV                                                 = 0x8880
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	SAMPLE_MASK_NV                                             = 0x8E51
	TEXTURE_UPDATE_BARRIER_BIT_EXT                             = 0x00000100
	TEXTURE_BINDING_3D                                         = 0x806A
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	PRIMITIVES_GENERATED_NV                                    = 0x8C87
	FLOAT16_NV                                                 = 0x8FF8
	SGIX_texture_coordinate_clamp                              = 1
	LIST_MODE                                                  = 0x0B30
	TEXTURE_BUFFER_FORMAT_EXT                                  = 0x8C2E
	RGB_SNORM                                                  = 0x8F92
	PACK_IMAGE_DEPTH_SGIS                                      = 0x8131
	BUFFER                                                     = 0x82E0
	CONSTANT_COLOR1_NV                                         = 0x852B
	NEGATIVE_X_EXT                                             = 0x87D9
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	EYE_LINEAR                                                 = 0x2400
	COMBINER_CD_OUTPUT_NV                                      = 0x854B
	PROGRAM_NATIVE_TEMPORARIES_ARB                             = 0x88A6
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	MAT_SPECULAR_BIT_PGI                                       = 0x04000000
	QUADRATIC_ATTENUATION                                      = 0x1209
	REPLACE_OLDEST_SUN                                         = 0x0003
	R8_EXT                                                     = 0x8229
	EXT_convolution                                            = 1
	LUMINANCE16_ALPHA16                                        = 0x8048
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	COLOR_TABLE_BLUE_SIZE_SGI                                  = 0x80DC
	TEXTURE_TARGET_QCOM                                        = 0x8BDA
	FRAMEBUFFER_COMPLETE_EXT                                   = 0x8CD5
	SKIP_COMPONENTS1_NV                                        = -6
	INT_SAMPLER_1D                                             = 0x8DC9
	SGIX_instruments                                           = 1
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	FRAGMENT_PROGRAM_INTERPOLATION_OFFSET_BITS_NV              = 0x8E5D
	FIELD_UPPER_NV                                             = 0x9022
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	VIEW_CLASS_8_BITS                                          = 0x82CB
	SOURCE2_ALPHA                                              = 0x858A
	DEPTH_CLAMP_NV                                             = 0x864F
	RGBA32F_ARB                                                = 0x8814
	MAX_MATRIX_PALETTE_STACK_DEPTH_ARB                         = 0x8841
	FLOAT_VEC2_ARB                                             = 0x8B50
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	COLOR_ATTACHMENT9_EXT                                      = 0x8CE9
	UNSIGNED_INT_SAMPLER_2D_ARRAY_EXT                          = 0x8DD7
	IMAGE_BUFFER_EXT                                           = 0x9051
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	NUM_SAMPLE_COUNTS                                          = 0x9380
	TEXTURE_GREEN_SIZE                                         = 0x805D
	NORMAL_MAP_NV                                              = 0x8511
	MAX_PROGRAM_LOOP_DEPTH_NV                                  = 0x88F7
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE_EXT           = 0x8CD3
	COVERAGE_EDGE_FRAGMENTS_NV                                 = 0x8ED6
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	MAP_READ_BIT_EXT                                           = 0x0001
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	RGB5_A1_EXT                                                = 0x8057
	MAJOR_VERSION                                              = 0x821B
	SOURCE2_RGB_ARB                                            = 0x8582
	VERTEX_PROGRAM_CALLBACK_FUNC_MESA                          = 0x8BB6
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	TEXTURE_GREEN_TYPE_ARB                                     = 0x8C11
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	PIXEL_TILE_CACHE_SIZE_SGIX                                 = 0x8145
	YCBCR_MESA                                                 = 0x8757
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	VERTEX_SUBROUTINE                                          = 0x92E8
	ALPHA_TEST_FUNC                                            = 0x0BC1
	LINE_RESET_TOKEN                                           = 0x0707
	GENERATE_MIPMAP_HINT_SGIS                                  = 0x8192
	VERTEX_ARRAY_PARALLEL_POINTERS_INTEL                       = 0x83F5
	COMBINER2_NV                                               = 0x8552
	TEXTURE_DEPTH_SIZE_ARB                                     = 0x884A
	MAX_GEOMETRY_PROGRAM_INVOCATIONS_NV                        = 0x8E5A
	UNSIGNED_INT_SAMPLER_BUFFER_AMD                            = 0x9003
	LINE_TOKEN                                                 = 0x0702
	TEXTURE_RECTANGLE_NV                                       = 0x84F5
	DOT3_RGB_EXT                                               = 0x8740
	WEIGHT_ARRAY_BUFFER_BINDING_OES                            = 0x889E
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	REQUIRED_TEXTURE_IMAGE_UNITS_OES                           = 0x8D68
	HIGH_INT                                                   = 0x8DF5
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	EXPAND_NORMAL_NV                                           = 0x8538
	INT_SAMPLER_CUBE_EXT                                       = 0x8DCC
	BOLD_BIT_NV                                                = 0x01
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	BGR_EXT                                                    = 0x80E0
	MAP1_VERTEX_ATTRIB11_4_NV                                  = 0x866B
	MAP_ATTRIB_U_ORDER_NV                                      = 0x86C3
	OUTPUT_TEXTURE_COORD22_EXT                                 = 0x87B3
	DRAW_BUFFER12_NV                                           = 0x8831
	RENDERBUFFER_HEIGHT_OES                                    = 0x8D43
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	UNSIGNED_INT_IMAGE_CUBE_EXT                                = 0x9066
	TIMEOUT_EXPIRED                                            = 0x911B
	MODELVIEW7_ARB                                             = 0x8727
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	INDEX_CLEAR_VALUE                                          = 0x0C20
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	STENCIL_INDEX8                                             = 0x8D48
	MAX_PROGRAM_SUBROUTINE_PARAMETERS_NV                       = 0x8F44
	INT16_NV                                                   = 0x8FE4
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	LINE                                                       = 0x1B01
	LUMINANCE16                                                = 0x8042
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	ALPHA_MAX_SGIX                                             = 0x8321
	LUMINANCE_ALPHA_FLOAT32_ATI                                = 0x8819
	RGBA_MODE                                                  = 0x0C31
	RESCALE_NORMAL_EXT                                         = 0x803A
	DETAIL_TEXTURE_FUNC_POINTS_SGIS                            = 0x809C
	OUTPUT_FOG_EXT                                             = 0x87BD
	RENDERBUFFER_RED_SIZE_OES                                  = 0x8D50
	UNSIGNED_INT_IMAGE_2D_ARRAY_EXT                            = 0x9069
	MAP_UNSYNCHRONIZED_BIT_EXT                                 = 0x0020
	FRAGMENT_LIGHT2_SGIX                                       = 0x840E
	OP_FRAC_EXT                                                = 0x8789
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	SHADER_BINARY_DMP                                          = 0x9250
	INTENSITY16_EXT                                            = 0x804D
	SAMPLE_COVERAGE                                            = 0x80A0
	COLOR_INDEX1_EXT                                           = 0x80E2
	MAX_PROGRAM_EXEC_INSTRUCTIONS_NV                           = 0x88F4
	TEXTURE_2D_ARRAY_EXT                                       = 0x8C1A
	LINE_STRIP_ADJACENCY                                       = 0x000B
	MAP1_INDEX                                                 = 0x0D91
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	ALPHA_MAX_CLAMP_INGR                                       = 0x8567
	TEXTURE_SWIZZLE_R_EXT                                      = 0x8E42
	DETAIL_TEXTURE_2D_SGIS                                     = 0x8095
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	GL_3D_COLOR                                                = 0x0602
	EYE_DISTANCE_TO_POINT_SGIS                                 = 0x81F0
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	TEXTURE19                                                  = 0x84D3
	SATURATE_BIT_ATI                                           = 0x00000040
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_EXT                      = 0x8CD6
	PACK_ALIGNMENT                                             = 0x0D05
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	MATRIX1_NV                                                 = 0x8631
	MATRIX_PALETTE_OES                                         = 0x8840
	CON_22_ATI                                                 = 0x8957
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_NV                       = 0x8E5F
	MAX_INTEGER_SAMPLES                                        = 0x9110
	PIXEL_MODE_BIT                                             = 0x00000020
	RGBA8_EXT                                                  = 0x8058
	R1UI_C3F_V3F_SUN                                           = 0x85C6
	OUTPUT_TEXTURE_COORD24_EXT                                 = 0x87B5
	PROGRAM_ATTRIB_COMPONENTS_NV                               = 0x8906
	AUX3                                                       = 0x040C
	UNSIGNED_INT_8_8_S8_S8_REV_NV                              = 0x86DB
	EXT_subtexture                                             = 1
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	POINT_SMOOTH                                               = 0x0B10
	FILL                                                       = 0x1B02
	RGB12_EXT                                                  = 0x8053
	RG8                                                        = 0x822B
	MAP1_VERTEX_ATTRIB15_4_NV                                  = 0x866F
	OFFSET_HILO_TEXTURE_2D_NV                                  = 0x8854
	DYNAMIC_READ_ARB                                           = 0x88E9
	TEXTURE_BIT                                                = 0x00040000
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	UNSIGNED_BYTE_2_3_3_REV_EXT                                = 0x8362
	MAX_TEXTURE_UNITS_ARB                                      = 0x84E2
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT                     = 0x8D56
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	IMAGE_1D_ARRAY                                             = 0x9052
	INVALID_ENUM                                               = 0x0500
	READ_BUFFER_NV                                             = 0x0C02
	FILTER                                                     = 0x829A
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	PROGRAM_INPUT                                              = 0x92E3
	OR                                                         = 0x1507
	DEPTH_COMPONENT                                            = 0x1902
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	QUERY_RESULT_AVAILABLE_EXT                                 = 0x8867
	MAX_SAMPLE_MASK_WORDS_NV                                   = 0x8E59
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	RESAMPLE_REPLICATE_SGIX                                    = 0x842E
	DEPTH24_STENCIL8_EXT                                       = 0x88F0
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	PATH_DASH_CAPS_NV                                          = 0x907B
	QUERY_OBJECT_AMD                                           = 0x9153
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	TEXTURE_1D                                                 = 0x0DE0
	COLOR_ARRAY_POINTER_EXT                                    = 0x8090
	DEBUG_PRINT_MESA                                           = 0x875A
	GREEN_BIT_ATI                                              = 0x00000002
	CONVOLUTION_FILTER_SCALE_EXT                               = 0x8014
	SRGB_WRITE                                                 = 0x8298
	LINE_WIDTH_RANGE                                           = 0x0B22
	POST_CONVOLUTION_ALPHA_BIAS                                = 0x8023
	SINGLE_COLOR                                               = 0x81F9
	DSDT8_MAG8_INTENSITY8_NV                                   = 0x870B
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	RGB4_S3TC                                                  = 0x83A1
	CURRENT_VERTEX_EXT                                         = 0x87E2
	DELETE_STATUS                                              = 0x8B80
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	PACK_IMAGE_HEIGHT                                          = 0x806C
	ACTIVE_TEXTURE                                             = 0x84E0
	COMBINE_ARB                                                = 0x8570
	GL_3DC_XY_AMD                                              = 0x87FA
	SRC1_COLOR                                                 = 0x88F9
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	GL_2_BYTES                                                 = 0x1407
	ETC1_SRGB8_NV                                              = 0x88EE
	SGIX_shadow                                                = 1
	FUNC_SUBTRACT_OES                                          = 0x800A
	MAX_WIDTH                                                  = 0x827E
	VERTEX_ARRAY_RANGE_LENGTH_APPLE                            = 0x851E
	MATRIX12_ARB                                               = 0x88CC
	TEXTURE_BINDING_EXTERNAL_OES                               = 0x8D67
	BLUE_INTEGER                                               = 0x8D96
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM                        = 103084
	SGIS_fog_function                                          = 1
	ALPHA12                                                    = 0x803D
	MAP2_VERTEX_ATTRIB10_4_NV                                  = 0x867A
	STREAM_READ_ARB                                            = 0x88E1
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	ELEMENT_ARRAY_BARRIER_BIT_EXT                              = 0x00000002
	FOG_OFFSET_SGIX                                            = 0x8198
	TEXTURE_UNSIGNED_REMAP_MODE_NV                             = 0x888F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV                         = 0x8C4F
	PROJECTION                                                 = 0x1701
	MODELVIEW20_ARB                                            = 0x8734
	EVAL_VERTEX_ATTRIB9_NV                                     = 0x86CF
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	TEXTURE_2D_BINDING_EXT                                     = 0x8069
	COLOR_CLEAR_UNCLAMPED_VALUE_ATI                            = 0x8835
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	RED_SCALE                                                  = 0x0D14
	BUMP_ROT_MATRIX_ATI                                        = 0x8775
	READ_WRITE_ARB                                             = 0x88BA
	ACTIVE_VERTEX_UNITS_ARB                                    = 0x86A5
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	ALWAYS                                                     = 0x0207
	DUAL_ALPHA8_SGIS                                           = 0x8111
	MATRIX6_NV                                                 = 0x8636
	BLEND_EQUATION_ALPHA_OES                                   = 0x883D
	TEXTURE_COMPARE_MODE_EXT                                   = 0x884C
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING_ARB                     = 0x889F
	CND0_ATI                                                   = 0x896B
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	VERTEX_ATTRIB_ARRAY_NORMALIZED_ARB                         = 0x886A
	BIAS_BIT_ATI                                               = 0x00000008
	COMPRESSED_RGBA_PVRTC_2BPPV1_IMG                           = 0x8C03
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH_EXT                  = 0x8C76
	DRAW_FRAMEBUFFER_NV                                        = 0x8CA9
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	RGBA8_OES                                                  = 0x8058
	QUERY_COUNTER_BITS_ARB                                     = 0x8864
	DOT4_ATI                                                   = 0x8967
	GL_2X_BIT_ATI                                              = 0x00000001
	INTENSITY16UI_EXT                                          = 0x8D79
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	MAX_OPTIMIZED_VERTEX_SHADER_LOCAL_CONSTANTS_EXT            = 0x87CC
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	BOOL_VEC2                                                  = 0x8B57
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG                     = 0x9134
	VERTEX_ARRAY_RANGE_VALID_NV                                = 0x851F
	CON_0_ATI                                                  = 0x8941
	SRGB8_ALPHA8                                               = 0x8C43
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	TEXTURE_VIEW                                               = 0x82B5
	UNPACK_COMPRESSED_SIZE_SGIX                                = 0x831A
	TEXTURE_COMPRESSED_ARB                                     = 0x86A1
	SIGNED_ALPHA_NV                                            = 0x8705
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	NORMAL_BIT_PGI                                             = 0x08000000
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	SECONDARY_COLOR_ARRAY_SIZE_EXT                             = 0x845A
	OFFSET_TEXTURE_MATRIX_NV                                   = 0x86E1
	TEXTURE_COORD_ARRAY_BUFFER_BINDING_ARB                     = 0x889A
	STREAM_COPY_ARB                                            = 0x88E2
	MATRIX_STRIDE                                              = 0x92FF
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	STENCIL_BITS                                               = 0x0D57
	NORMAL_ARRAY                                               = 0x8075
	LUMINANCE12_ALPHA4                                         = 0x8046
	DUAL_INTENSITY12_SGIS                                      = 0x811A
	TEXTURE16_ARB                                              = 0x84D0
	Z_EXT                                                      = 0x87D7
	LUMINANCE8_SNORM                                           = 0x9015
	UNIFORM_BARRIER_BIT_EXT                                    = 0x00000004
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	PROGRAM_RESIDENT_NV                                        = 0x8647
	QUAD_INTENSITY4_SGIS                                       = 0x8122
	TEXTURE27                                                  = 0x84DB
	MAP2_VERTEX_4                                              = 0x0DB8
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	ALPHA8I_EXT                                                = 0x8D90
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	TEXTURE5                                                   = 0x84C5
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	PACK_SKIP_PIXELS                                           = 0x0D04
	HALF_FLOAT                                                 = 0x140B
	SOURCE3_ALPHA_NV                                           = 0x858B
	LUMINANCE_ALPHA32F_ARB                                     = 0x8819
	POINT_SPRITE_NV                                            = 0x8861
	SAMPLE_SHADING_ARB                                         = 0x8C36
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	WRITE_PIXEL_DATA_RANGE_LENGTH_NV                           = 0x887A
	UNSIGNED_INT_5_9_9_9_REV_EXT                               = 0x8C3E
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	OPERAND1_ALPHA_EXT                                         = 0x8599
	MATRIX3_ARB                                                = 0x88C3
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_EXT              = 0x8CD4
	OPERAND1_RGB_EXT                                           = 0x8591
	SIGNED_INTENSITY_NV                                        = 0x8707
	SGX_PROGRAM_BINARY_IMG                                     = 0x9130
	BLOCK_INDEX                                                = 0x92FD
	GREEN_BIAS                                                 = 0x0D19
	ONE_MINUS_CONSTANT_COLOR_EXT                               = 0x8002
	CONVOLUTION_FILTER_BIAS_EXT                                = 0x8015
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB                            = 0x8519
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	BUFFER_DATA_SIZE                                           = 0x9303
	RGBA4                                                      = 0x8056
	DUAL_LUMINANCE16_SGIS                                      = 0x8117
	DUAL_INTENSITY16_SGIS                                      = 0x811B
	MODELVIEW12_ARB                                            = 0x872C
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	VIDEO_COLOR_CONVERSION_MIN_NV                              = 0x902B
	NOOP                                                       = 0x1505
	VERTEX_PRECLIP_HINT_SGIX                                   = 0x83EF
	BINORMAL_ARRAY_POINTER_EXT                                 = 0x8443
	ALPHA4                                                     = 0x803B
	CLIP_DISTANCE1                                             = 0x3001
	DOT_PRODUCT_REFLECT_CUBE_MAP_NV                            = 0x86F2
	DRAW_BUFFER13_NV                                           = 0x8832
	TEXTURE9_ARB                                               = 0x84C9
	COMPRESSED_INTENSITY                                       = 0x84EC
	FRAMEBUFFER                                                = 0x8D40
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	HI_SCALE_NV                                                = 0x870E
	ITALIC_BIT_NV                                              = 0x02
	NORMAL_ARRAY_TYPE                                          = 0x807E
	WEIGHT_ARRAY_SIZE_OES                                      = 0x86AB
	DRAW_BUFFER6_ARB                                           = 0x882B
	SAMPLER_1D                                                 = 0x8B5D
	PATH_GEN_COLOR_FORMAT_NV                                   = 0x90B2
	POSITION                                                   = 0x1203
	UNPACK_CMYK_HINT_EXT                                       = 0x800F
	EDGE_FLAG_ARRAY_COUNT_EXT                                  = 0x808D
	DRAW_BUFFER13_ARB                                          = 0x8832
	MOV_ATI                                                    = 0x8961
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	TRIANGLE_FAN                                               = 0x0006
	TEXTURE_GEN_R                                              = 0x0C62
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	PROGRAM_TARGET_NV                                          = 0x8646
	PROGRAM_NATIVE_ATTRIBS_ARB                                 = 0x88AE
	FLOAT_MAT4_ARB                                             = 0x8B5C
	LOW_FLOAT                                                  = 0x8DF0
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MAX_CONVOLUTION_WIDTH_EXT                                  = 0x801A
	INT_SAMPLER_RENDERBUFFER_NV                                = 0x8E57
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	DISTANCE_ATTENUATION_EXT                                   = 0x8129
	CLAMP_TO_BORDER_NV                                         = 0x812D
	INTENSITY_SNORM                                            = 0x9013
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	ARRAY_OBJECT_OFFSET_ATI                                    = 0x8767
	FLOAT_32_UNSIGNED_INT_24_8_REV_NV                          = 0x8DAD
	MULTISAMPLE_ARB                                            = 0x809D
	COLOR_TABLE_LUMINANCE_SIZE_SGI                             = 0x80DE
	READ_PIXELS_TYPE                                           = 0x828E
	PROGRAM_POINT_SIZE                                         = 0x8642
	OP_MUL_EXT                                                 = 0x8786
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	POST_CONVOLUTION_ALPHA_SCALE                               = 0x801F
	SAMPLE_MASK_VALUE_SGIS                                     = 0x80AA
	TEXTURE_LOD_BIAS                                           = 0x8501
	GENERIC_ATTRIB_NV                                          = 0x8C7D
	FRAMEBUFFER_SRGB_EXT                                       = 0x8DB9
	UNIFORM_BUFFER_EXT                                         = 0x8DEE
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	SGIS_texture_filter4                                       = 1
	SCISSOR_BIT                                                = 0x00080000
	DUAL_LUMINANCE12_SGIS                                      = 0x8116
	TEXTURE_COMPRESSED_IMAGE_SIZE_ARB                          = 0x86A0
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	INTERLACE_READ_OML                                         = 0x8981
	TEXTURE_GEN_S                                              = 0x0C60
	NEAREST_CLIPMAP_LINEAR_SGIX                                = 0x844E
	OPERAND0_ALPHA_ARB                                         = 0x8598
	COLOR_ATTACHMENT13_EXT                                     = 0x8CED
	MAX_PROGRAM_PARAMETER_BUFFER_BINDINGS_NV                   = 0x8DA0
	ONE_MINUS_DST_COLOR                                        = 0x0307
	OP_RECIP_SQRT_EXT                                          = 0x8795
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN_NV                   = 0x8C88
	RGB32UI                                                    = 0x8D71
	MAX_GEOMETRY_UNIFORM_COMPONENTS_EXT                        = 0x8DDF
	EMBOSS_MAP_NV                                              = 0x855F
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	CONVOLUTION_2D                                             = 0x8011
	POST_CONVOLUTION_COLOR_TABLE_SGI                           = 0x80D1
	COLOR_TABLE_BIAS_SGI                                       = 0x80D7
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	COLOR_SUM_CLAMP_NV                                         = 0x854F
	SIGNED_INTENSITY8_NV                                       = 0x8708
	GEOMETRY_SHADER                                            = 0x8DD9
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	COLOR_TABLE_SCALE                                          = 0x80D6
	COLOR_RENDERABLE                                           = 0x8286
	LIGHT_ENV_MODE_SGIX                                        = 0x8407
	DRAW_BUFFER8_ATI                                           = 0x882D
	POINT_SIZE_ARRAY_BUFFER_BINDING_OES                        = 0x8B9F
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	COMPRESSED_RGBA_PVRTC_2BPPV2_IMG                           = 0x9137
	NAME_STACK_DEPTH                                           = 0x0D70
	UNSIGNED_INT_8_8_8_8_EXT                                   = 0x8035
	UNPACK_CONSTANT_DATA_SUNX                                  = 0x81D5
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	LUMINANCE4_EXT                                             = 0x803F
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	VOLATILE_APPLE                                             = 0x8A1A
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	LUMINANCE_ALPHA16I_EXT                                     = 0x8D8D
	BUFFER_MAP_OFFSET                                          = 0x9121
	GEQUAL                                                     = 0x0206
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	GREEN_INTEGER_EXT                                          = 0x8D95
	FRAMEBUFFER_COMPLETE_OES                                   = 0x8CD5
	LUMINANCE16I_EXT                                           = 0x8D8C
	REPLACE                                                    = 0x1E01
	REPLICATE_BORDER                                           = 0x8153
	PIXEL_TEX_GEN_Q_FLOOR_SGIX                                 = 0x8186
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	OUTPUT_TEXTURE_COORD11_EXT                                 = 0x87A8
	BACK_SECONDARY_COLOR_NV                                    = 0x8C78
	STENCIL_INDEX8_OES                                         = 0x8D48
	TEXTURE_ENV                                                = 0x2300
	PROXY_TEXTURE_RECTANGLE_NV                                 = 0x84F7
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	FRAGMENT_PROGRAM_POSITION_MESA                             = 0x8BB0
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	GENERATE_MIPMAP_HINT                                       = 0x8192
	MAX_TEXTURE_IMAGE_UNITS_NV                                 = 0x8872
	POINT_SIZE_ARRAY_TYPE_OES                                  = 0x898A
	LUMINANCE4                                                 = 0x803F
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	MATRIX16_ARB                                               = 0x88D0
	STENCIL_EXT                                                = 0x1802
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	FLOAT_MAT2                                                 = 0x8B5A
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	MULTISAMPLE_BIT_3DFX                                       = 0x20000000
	MIN                                                        = 0x8007
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE                   = 0x8D56
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	INT16_VEC4_NV                                              = 0x8FE7
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	ALL_STATIC_DATA_IBM                                        = 103060
	UNSIGNED_INT_2_10_10_10_REV_EXT                            = 0x8368
	BGRA                                                       = 0x80E1
	ACCUM                                                      = 0x0100
	OFFSET_TEXTURE_2D_MATRIX_NV                                = 0x86E1
	OBJECT_ACTIVE_ATTRIBUTES_ARB                               = 0x8B89
	INT_IMAGE_CUBE                                             = 0x905B
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX                     = 0x8403
	BINORMAL_ARRAY_STRIDE_EXT                                  = 0x8441
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	FUNC_ADD_OES                                               = 0x8006
	TEXTURE_MAX_CLAMP_T_SGIX                                   = 0x836A
	FLOAT_VEC2                                                 = 0x8B50
	ETC1_RGB8_OES                                              = 0x8D64
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY_EXT                      = 0x906A
	COPY_INVERTED                                              = 0x150C
	MAX_OPTIMIZED_VERTEX_SHADER_INVARIANTS_EXT                 = 0x87CD
	STENCIL_BACK_FUNC                                          = 0x8800
	CON_25_ATI                                                 = 0x895A
	SPOT_EXPONENT                                              = 0x1205
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	R16UI                                                      = 0x8234
	SURFACE_MAPPED_NV                                          = 0x8700
	OUTPUT_TEXTURE_COORD17_EXT                                 = 0x87AE
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	YCRCB_444_SGIX                                             = 0x81BC
	DEPTH_CLAMP                                                = 0x864F
	MAX_PROGRAM_ATTRIB_COMPONENTS_NV                           = 0x8908
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	QUERY_OBJECT_EXT                                           = 0x9153
	HISTOGRAM_WIDTH_EXT                                        = 0x8026
	DEFORMATIONS_MASK_SGIX                                     = 0x8196
	RG16UI                                                     = 0x823A
	MAP1_VERTEX_ATTRIB4_4_NV                                   = 0x8664
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	HI_BIAS_NV                                                 = 0x8714
	DUDV_ATI                                                   = 0x8779
	CON_30_ATI                                                 = 0x895F
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	MAX_VERTEX_STREAMS                                         = 0x8E71
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	TEXTURE_MIN_LOD                                            = 0x813A
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	PROJECTION_MATRIX_FLOAT_AS_INT_BITS_OES                    = 0x898E
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	IS_ROW_MAJOR                                               = 0x9300
	GENERATE_MIPMAP                                            = 0x8191
	GLOBAL_ALPHA_FACTOR_SUN                                    = 0x81DA
	TEXTURE_STORAGE_HINT_APPLE                                 = 0x85BC
	DRAW_BUFFER4_ATI                                           = 0x8829
	DRAW_BUFFER2_NV                                            = 0x8827
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	PROGRAM_NATIVE_INSTRUCTIONS_ARB                            = 0x88A2
	MATRIX9_ARB                                                = 0x88C9
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	RGB4_EXT                                                   = 0x804F
	MAP2_VERTEX_ATTRIB1_4_NV                                   = 0x8671
	BUFFER_MAPPED_OES                                          = 0x88BC
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	TIMEOUT_IGNORED_APPLE                                      = 0xFFFFFFFF
	LIST_BASE                                                  = 0x0B32
	TEXTURE_MAX_CLAMP_R_SGIX                                   = 0x836B
	FRAGMENT_DEPTH_EXT                                         = 0x8452
	SCISSOR_BOX                                                = 0x0C10
	LUMINANCE6_ALPHA2                                          = 0x8044
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	MODELVIEW1_STACK_DEPTH_EXT                                 = 0x8502
	RGB_FLOAT32_ATI                                            = 0x8815
	MAX_SAMPLES_ANGLE                                          = 0x8D57
	SAMPLER_BUFFER_AMD                                         = 0x9001
	PATH_FORMAT_PS_NV                                          = 0x9071
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	MAX_3D_TEXTURE_SIZE_EXT                                    = 0x8073
	CALLIGRAPHIC_FRAGMENT_SGIX                                 = 0x8183
	OPERAND2_RGB_EXT                                           = 0x8592
	TEXTURE_HEIGHT_QCOM                                        = 0x8BD3
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_OES                     = 0x8CDC
	ISOLINES                                                   = 0x8E7A
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	CURRENT_VERTEX_ATTRIB_ARB                                  = 0x8626
	DOT3_RGBA                                                  = 0x86AF
	TEXTURE_ALPHA_MODULATE_IMG                                 = 0x8C06
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	RG8_SNORM                                                  = 0x8F95
	CURRENT_BIT                                                = 0x00000001
	POST_CONVOLUTION_GREEN_BIAS                                = 0x8021
	BACK                                                       = 0x0405
	IGNORE_BORDER_HP                                           = 0x8150
	COLOR_SUM_ARB                                              = 0x8458
	DRAW_BUFFER8_ARB                                           = 0x882D
	FOG_COORDINATE_ARRAY_LIST_IBM                              = 103076
	CLIP_PLANE4                                                = 0x3004
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX                        = 0x8311
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	TRANSFORM_FEEDBACK_BARRIER_BIT_EXT                         = 0x00000800
	SAMPLES_SGIS                                               = 0x80A9
	PIXEL_TILE_GRID_WIDTH_SGIX                                 = 0x8142
	FRAGMENT_DEPTH                                             = 0x8452
	TEXTURE_BINDING_2D_ARRAY_EXT                               = 0x8C1D
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	RENDERBUFFER_BINDING_EXT                                   = 0x8CA7
	VERTICAL_LINE_TO_NV                                        = 0x08
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	SGIS_texture_edge_clamp                                    = 1
	CONVOLUTION_HEIGHT_EXT                                     = 0x8019
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	OFFSET_TEXTURE_RECTANGLE_SCALE_NV                          = 0x864D
	DRAW_BUFFER13                                              = 0x8832
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	VENDOR                                                     = 0x1F00
	FUNC_ADD_EXT                                               = 0x8006
	TRANSFORM_HINT_APPLE                                       = 0x85B1
	LINEAR_DETAIL_COLOR_SGIS                                   = 0x8099
	CONVOLUTION_BORDER_COLOR_HP                                = 0x8154
	IUI_V2F_EXT                                                = 0x81AD
	TEXTURE17_ARB                                              = 0x84D1
	NORMAL_MAP_ARB                                             = 0x8511
	SOURCE2_ALPHA_EXT                                          = 0x858A
	DONT_CARE                                                  = 0x1100
	COMBINE                                                    = 0x8570
	MAP1_VERTEX_ATTRIB1_4_NV                                   = 0x8661
	DEBUG_CATEGORY_PERFORMANCE_AMD                             = 0x914D
	MAP_COLOR                                                  = 0x0D10
	SPHERE_MAP                                                 = 0x2402
	ALPHA8_EXT                                                 = 0x803C
	FLOAT_MAT3x4                                               = 0x8B68
	OBJECT_DELETE_STATUS_ARB                                   = 0x8B80
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	POINT                                                      = 0x1B00
	CURRENT_MATRIX_STACK_DEPTH_ARB                             = 0x8640
	TEXTURE_LO_SIZE_NV                                         = 0x871C
	MAX_PROGRAM_CALL_DEPTH_NV                                  = 0x88F5
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_NV                 = 0x8C8B
	VIDEO_BUFFER_INTERNAL_FORMAT_NV                            = 0x902D
	TEXCOORD3_BIT_PGI                                          = 0x40000000
	CONVOLUTION_WIDTH_EXT                                      = 0x8018
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	BLUE_BITS                                                  = 0x0D54
	Z6Y10Z6CB10Z6A10Z6Y10Z6CR10Z6A10_4224_NV                   = 0x9034
	HISTOGRAM_ALPHA_SIZE                                       = 0x802B
	DEBUG_TYPE_MARKER                                          = 0x8268
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN                          = 0x85C2
	MODELVIEW14_ARB                                            = 0x872E
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	INT_IMAGE_2D_ARRAY_EXT                                     = 0x905E
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	EYE_PLANE_ABSOLUTE_NV                                      = 0x855C
	MODELVIEW19_ARB                                            = 0x8733
	CONST_EYE_NV                                               = 0x86E5
	MIRROR_CLAMP_ATI                                           = 0x8742
	STATIC_DRAW_ARB                                            = 0x88E4
	RED_INTEGER_EXT                                            = 0x8D94
	ALPHA_SNORM                                                = 0x9010
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	ALPHA16UI_EXT                                              = 0x8D78
	OP_MAX_EXT                                                 = 0x878A
	Y_EXT                                                      = 0x87D6
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	SRGB8_NV                                                   = 0x8C41
	BGRA_INTEGER_EXT                                           = 0x8D9B
	DUAL_LUMINANCE8_SGIS                                       = 0x8115
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	SKIP_DECODE_EXT                                            = 0x8A4A
	TESS_CONTROL_SHADER                                        = 0x8E88
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	SHADOW_AMBIENT_SGIX                                        = 0x80BF
	SRC1_RGB                                                   = 0x8581
	BLEND_EQUATION_ALPHA                                       = 0x883D
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	SGIX_tag_sample_buffer                                     = 1
	TEXTURE_HEIGHT                                             = 0x1001
	POST_CONVOLUTION_RED_BIAS_EXT                              = 0x8020
	SAMPLE_COVERAGE_ARB                                        = 0x80A0
	PN_TRIANGLES_NORMAL_MODE_LINEAR_ATI                        = 0x87F7
	DRAW_BUFFER11_NV                                           = 0x8830
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	IMAGE_2D_MULTISAMPLE_ARRAY_EXT                             = 0x9056
	STRICT_SCISSOR_HINT_PGI                                    = 0x1A218
	NEAREST                                                    = 0x2600
	TEXTURE_4D_BINDING_SGIS                                    = 0x814F
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	MAX_PROGRAM_ADDRESS_REGISTERS_ARB                          = 0x88B1
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	TEXTURE_INTENSITY_TYPE                                     = 0x8C15
	DRAW_FRAMEBUFFER_BINDING_NV                                = 0x8CA6
	GEOMETRY_VERTICES_OUT_EXT                                  = 0x8DDA
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	MAX_DEEP_3D_TEXTURE_WIDTH_HEIGHT_NV                        = 0x90D0
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	ACCUM_RED_BITS                                             = 0x0D58
	POLYGON_OFFSET_EXT                                         = 0x8037
	UNPACK_RESAMPLE_SGIX                                       = 0x842D
	ARRAY_BUFFER_BINDING                                       = 0x8894
	OBJECT_TYPE_ARB                                            = 0x8B4E
	HISTOGRAM_FORMAT                                           = 0x8027
	BUMP_NUM_TEX_UNITS_ATI                                     = 0x8777
	DRAW_BUFFER8                                               = 0x882D
	VERTEX_ATTRIB_MAP2_APPLE                                   = 0x8A01
	RGBA16UI                                                   = 0x8D76
	BLEND_DST                                                  = 0x0BE0
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	CURRENT_TANGENT_EXT                                        = 0x843B
	MODULATE_SIGNED_ADD_ATI                                    = 0x8745
	PROGRAM_TEMPORARIES_ARB                                    = 0x88A4
	STREAM_DRAW                                                = 0x88E0
	VIRTUAL_PAGE_SIZE_Y_AMD                                    = 0x9196
	SGIX_texture_scale_bias                                    = 1
	MAP1_NORMAL                                                = 0x0D92
	VIEW_CLASS_64_BITS                                         = 0x82C6
	REG_10_ATI                                                 = 0x892B
	LINEAR_SHARPEN_COLOR_SGIS                                  = 0x80AF
	TEXTURE7                                                   = 0x84C7
	MAX_RENDERBUFFER_SIZE_EXT                                  = 0x84E8
	WEIGHT_ARRAY_STRIDE_ARB                                    = 0x86AA
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	UNSIGNED_INT_IMAGE_1D_EXT                                  = 0x9062
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	TEXTURE_RECTANGLE_ARB                                      = 0x84F5
	BUFFER_ACCESS_OES                                          = 0x88BB
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	DUAL_ALPHA12_SGIS                                          = 0x8112
	MAP2_VERTEX_ATTRIB6_4_NV                                   = 0x8676
	LUMINANCE_FLOAT32_ATI                                      = 0x8818
	ANY_SAMPLES_PASSED_EXT                                     = 0x8C2F
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	COLOR_INDEX16_EXT                                          = 0x80E7
	SWIZZLE_STRQ_ATI                                           = 0x897A
	SLUMINANCE_ALPHA                                           = 0x8C44
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	R16F                                                       = 0x822D
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	LUMINANCE8I_EXT                                            = 0x8D92
	PATH_STROKE_MASK_NV                                        = 0x9084
	RGB_SCALE_EXT                                              = 0x8573
	MODELVIEW24_ARB                                            = 0x8738
	MATRIX_INDEX_ARRAY_STRIDE_ARB                              = 0x8848
	MAX_ARRAY_TEXTURE_LAYERS_EXT                               = 0x88FF
	NEGATE_BIT_ATI                                             = 0x00000004
	IMAGE_2D_EXT                                               = 0x904D
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	ALPHA_TEST_REF                                             = 0x0BC2
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	COLOR_TABLE_ALPHA_SIZE_SGI                                 = 0x80DD
	DRAW_BUFFER1_NV                                            = 0x8826
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	TEXTURE12                                                  = 0x84CC
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	DOUBLE_MAT2x4                                              = 0x8F4A
	CLIP_FAR_HINT_PGI                                          = 0x1A221
	DEPTH_RANGE                                                = 0x0B70
	OUTPUT_TEXTURE_COORD2_EXT                                  = 0x879F
	GL_8X_BIT_ATI                                              = 0x00000004
	DRAW_FRAMEBUFFER_EXT                                       = 0x8CA9
	DEPTH_ATTACHMENT                                           = 0x8D00
	RGB10_A2UI                                                 = 0x906F
	MINOR_VERSION                                              = 0x821C
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	MODELVIEW30_ARB                                            = 0x873E
	VERTEX_ARRAY_BUFFER_BINDING_ARB                            = 0x8896
	ACTIVE_VARYINGS_NV                                         = 0x8C81
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_NV           = 0x8C8A
	UNSIGNED_INT                                               = 0x1405
	TEXTURE_CUBE_MAP_OES                                       = 0x8513
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_OES                     = 0x8CDB
	MAX_SERVER_WAIT_TIMEOUT_APPLE                              = 0x9111
	LUMINANCE                                                  = 0x1909
	TRANSFORM_FEEDBACK_BUFFER_PAUSED_NV                        = 0x8E23
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	DEPTH_EXT                                                  = 0x1801
	POST_CONVOLUTION_RED_BIAS                                  = 0x8020
	MAX_ASYNC_TEX_IMAGE_SGIX                                   = 0x835F
	TEXTURE1                                                   = 0x84C1
	SLICE_ACCUM_SUN                                            = 0x85CC
	SURFACE_STATE_NV                                           = 0x86EB
	SCALAR_EXT                                                 = 0x87BE
	INT_VEC4_ARB                                               = 0x8B55
	POST_COLOR_MATRIX_GREEN_SCALE                              = 0x80B5
	DOT_PRODUCT_TEXTURE_RECTANGLE_NV                           = 0x864E
	PIXEL_COUNT_NV                                             = 0x8866
	SWIZZLE_STR_DR_ATI                                         = 0x8978
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	UNSIGNED_INT_SAMPLER_2D_EXT                                = 0x8DD2
	VERSION_3_2                                                = 1
	RGBA8                                                      = 0x8058
	MODELVIEW29_ARB                                            = 0x873D
	TEXTURE_BORDER_VALUES_NV                                   = 0x871A
	RGB_FLOAT32_APPLE                                          = 0x8815
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	FAILURE_NV                                                 = 0x9030
	TEXTURE_ENV_MODE                                           = 0x2200
	PROVOKING_VERTEX_EXT                                       = 0x8E4F
	RGBA8_SNORM                                                = 0x8F97
	POST_CONVOLUTION_GREEN_SCALE_EXT                           = 0x801D
	STENCIL_BACK_FAIL_ATI                                      = 0x8801
	SAMPLER_2D_SHADOW_EXT                                      = 0x8B62
	LIGHT2                                                     = 0x4002
	PHONG_WIN                                                  = 0x80EA
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	TEXTURE_BINDING_1D_ARRAY_EXT                               = 0x8C1C
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	POLYGON                                                    = 0x0009
	TEXTURE_COMPONENTS                                         = 0x1003
	UNSIGNED_INT_24_8_NV                                       = 0x84FA
	TEXTURE_CONSTANT_DATA_SUNX                                 = 0x81D6
	MODELVIEW15_ARB                                            = 0x872F
	EVAL_VERTEX_ATTRIB10_NV                                    = 0x86D0
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	TRIANGULAR_NV                                              = 0x90A5
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	FOG_COORDINATE_SOURCE_EXT                                  = 0x8450
	REG_13_ATI                                                 = 0x892E
	CON_13_ATI                                                 = 0x894E
	ELEMENT_ARRAY_TYPE_APPLE                                   = 0x8A0D
	CMYKA_EXT                                                  = 0x800D
	QUAD_INTENSITY8_SGIS                                       = 0x8123
	PIXEL_TEX_GEN_ALPHA_LS_SGIX                                = 0x8189
	MODULATE_SUBTRACT_ATI                                      = 0x8746
	REG_15_ATI                                                 = 0x8930
	MAX_VERTEX_UNIFORM_COMPONENTS_ARB                          = 0x8B4A
	TRANSFORM_FEEDBACK_BUFFER_NV                               = 0x8C8E
	POST_CONVOLUTION_GREEN_BIAS_EXT                            = 0x8021
	TEXTURE7_ARB                                               = 0x84C7
	MAP2_VERTEX_ATTRIB14_4_NV                                  = 0x867E
	FLOAT_RGB32_NV                                             = 0x8889
	MATRIX4_ARB                                                = 0x88C4
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	CLIP_DISTANCE0                                             = 0x3000
	OUTPUT_COLOR0_EXT                                          = 0x879B
	POINT_SIZE_ARRAY_OES                                       = 0x8B9C
	TEXTURE_GEN_STR_OES                                        = 0x8D60
	SRC_ALPHA_SATURATE                                         = 0x0308
	MIRROR_CLAMP_TO_BORDER_EXT                                 = 0x8912
	REG_14_ATI                                                 = 0x892F
	FRAGMENT_PROGRAM_CALLBACK_DATA_MESA                        = 0x8BB3
	IMAGE_2D_RECT_EXT                                          = 0x904F
	FRAMEBUFFER_BARRIER_BIT_EXT                                = 0x00000400
	RGB5_EXT                                                   = 0x8050
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	VERTEX_ARRAY                                               = 0x8074
	TEXTURE0_ARB                                               = 0x84C0
	BLUE_MAX_CLAMP_INGR                                        = 0x8566
	INDEX_ARRAY_BUFFER_BINDING_ARB                             = 0x8899
	CLAMP_FRAGMENT_COLOR_ARB                                   = 0x891B
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	ALPHA16                                                    = 0x803E
	SIGNED_LUMINANCE8_ALPHA8_NV                                = 0x8704
	POINT_SPRITE_R_MODE_NV                                     = 0x8863
	MATRIX25_ARB                                               = 0x88D9
	STREAM_DRAW_ARB                                            = 0x88E0
	MAP_INVALIDATE_RANGE_BIT_EXT                               = 0x0004
	OPERAND0_ALPHA                                             = 0x8598
	AUX0                                                       = 0x0409
	DEPTH                                                      = 0x1801
	TEXTURE27_ARB                                              = 0x84DB
	INDEX_ARRAY_LIST_STRIDE_IBM                                = 103083
	PACK_SKIP_IMAGES                                           = 0x806B
	PROGRAM_PIPELINE                                           = 0x82E4
	MAX_VERTEX_STREAMS_ATI                                     = 0x876B
	ALPHA32I_EXT                                               = 0x8D84
	QUERY_BUFFER_BINDING_AMD                                   = 0x9193
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_EXT                   = 0x8DA8
	LIGHTING                                                   = 0x0B50
	AUX_BUFFERS                                                = 0x0C00
	LINEAR_DETAIL_ALPHA_SGIS                                   = 0x8098
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	UNSIGNED_INT_IMAGE_2D_EXT                                  = 0x9063
	DEPTH_COMPONENT16_OES                                      = 0x81A5
	SAMPLE_BUFFERS_3DFX                                        = 0x86B3
	EVAL_VERTEX_ATTRIB0_NV                                     = 0x86C6
	COMPRESSED_SRGB_S3TC_DXT1_NV                               = 0x8C4C
	UNSIGNED_INT_SAMPLER_1D_ARRAY_EXT                          = 0x8DD6
	INT8_VEC4_NV                                               = 0x8FE3
	MAP_INVALIDATE_BUFFER_BIT_EXT                              = 0x0008
	EYE_LINE_SGIS                                              = 0x81F6
	PALETTE4_RGBA8_OES                                         = 0x8B91
	COLOR_ATTACHMENT_EXT                                       = 0x90F0
	DEBUG_SEVERITY_LOW_AMD                                     = 0x9148
	COLOR_TABLE_WIDTH                                          = 0x80D9
	DS_SCALE_NV                                                = 0x8710
	SAMPLES_PASSED_ARB                                         = 0x8914
	TRANSFORM_FEEDBACK_NV                                      = 0x8E22
	NATIVE_GRAPHICS_END_HINT_PGI                               = 0x1A204
	TEXTURE_COMPARE_FAIL_VALUE_ARB                             = 0x80BF
	GEOMETRY_TEXTURE                                           = 0x829E
	MIRRORED_REPEAT_OES                                        = 0x8370
	VERTEX_ATTRIB_ARRAY0_NV                                    = 0x8650
	BOOL_ARB                                                   = 0x8B56
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	RGB8UI_EXT                                                 = 0x8D7D
	GREEN_SCALE                                                = 0x0D18
	COLOR_TABLE_FORMAT                                         = 0x80D8
	IMAGE_CUBIC_WEIGHT_HP                                      = 0x815E
	VERTEX_SHADER_VARIANTS_EXT                                 = 0x87D0
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	VERTEX_ID_NV                                               = 0x8C7B
	UPPER_LEFT                                                 = 0x8CA2
	BLEND_SRC_ALPHA_OES                                        = 0x80CB
	MATRIX_INDEX_ARRAY_SIZE_OES                                = 0x8846
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	DOUBLE_VEC2                                                = 0x8FFC
	CONVOLUTION_HEIGHT                                         = 0x8019
	REG_18_ATI                                                 = 0x8933
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	DEBUG_OUTPUT                                               = 0x92E0
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	LUMINANCE_ALPHA                                            = 0x190A
	BLEND_DST_RGB_OES                                          = 0x80C8
	CURRENT_MATRIX_NV                                          = 0x8641
	DEBUG_OBJECT_MESA                                          = 0x8759
	REG_12_ATI                                                 = 0x892D
	TEXTURE_INTERNAL_FORMAT_QCOM                               = 0x8BD5
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	MAX_IMAGE_SAMPLES                                          = 0x906D
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	INDEX_WRITEMASK                                            = 0x0C21
	PATH_MITER_LIMIT_NV                                        = 0x907A
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	UNSIGNED_SHORT_4_4_4_4_EXT                                 = 0x8033
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	COMPRESSED_RED_RGTC1_EXT                                   = 0x8DBB
	COMPRESSED_LUMINANCE                                       = 0x84EA
	COORD_REPLACE                                              = 0x8862
	QUERY_RESULT_ARB                                           = 0x8866
	BUFFER_ACCESS                                              = 0x88BB
	RGB16_SNORM                                                = 0x8F9A
	SGIS_texture_lod                                           = 1
	COMBINE_ALPHA_ARB                                          = 0x8572
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	LINES_ADJACENCY                                            = 0x000A
	STORAGE_CLIENT_APPLE                                       = 0x85B4
	PRIMITIVE_ID_NV                                            = 0x8C7C
	COLOR_ATTACHMENT14                                         = 0x8CEE
	INTENSITY16I_EXT                                           = 0x8D8B
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	BITMAP                                                     = 0x1A00
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	CURRENT_RASTER_NORMAL_SGIX                                 = 0x8406
	MODELVIEW1_EXT                                             = 0x850A
	DISCARD_ATI                                                = 0x8763
	REG_31_ATI                                                 = 0x8940
	NEAREST_CLIPMAP_NEAREST_SGIX                               = 0x844D
	TRACE_NAME_MESA                                            = 0x8756
	VERTEX_PROGRAM_CALLBACK_MESA                               = 0x8BB4
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_EXT                      = 0x8CD9
	VIDEO_BUFFER_PITCH_NV                                      = 0x9028
	SGIS_pixel_texture                                         = 1
	RGB10_A2_EXT                                               = 0x8059
	TEXTURE_COORD_ARRAY_TYPE_EXT                               = 0x8089
	DUAL_LUMINANCE_ALPHA4_SGIS                                 = 0x811C
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	TEXTURE_MATRIX_FLOAT_AS_INT_BITS_OES                       = 0x898F
	VERTEX_PROGRAM_CALLBACK_DATA_MESA                          = 0x8BB7
	COMBINER_BIAS_NV                                           = 0x8549
	OUTPUT_TEXTURE_COORD5_EXT                                  = 0x87A2
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	EXT_packed_pixels                                          = 1
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	FASTEST                                                    = 0x1101
	MAX_FRAMEZOOM_FACTOR_SGIX                                  = 0x818D
	SLIM10U_SGIX                                               = 0x831E
	ASYNC_MARKER_SGIX                                          = 0x8329
	MAX_PROGRAM_RESULT_COMPONENTS_NV                           = 0x8909
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	CULL_FACE_MODE                                             = 0x0B45
	TEXTURE_PRIORITY_EXT                                       = 0x8066
	CLEAR_BUFFER                                               = 0x82B4
	SAMPLES_3DFX                                               = 0x86B4
	TESS_EVALUATION_PROGRAM_NV                                 = 0x891F
	SWIZZLE_STQ_ATI                                            = 0x8977
	SEPARATE_ATTRIBS_NV                                        = 0x8C8D
	INTENSITY8I_EXT                                            = 0x8D91
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	PACK_SKIP_IMAGES_EXT                                       = 0x806B
	SAMPLE_ALPHA_TO_ONE_SGIS                                   = 0x809F
	MAX_PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                    = 0x8810
	FOG_SPECULAR_TEXTURE_WIN                                   = 0x80EC
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	TEXTURE20_ARB                                              = 0x84D4
	INDEX_ARRAY_STRIDE_EXT                                     = 0x8086
	GL_422_EXT                                                 = 0x80CC
	EYE_RADIAL_NV                                              = 0x855B
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	COMPRESSED_RGB_PVRTC_2BPPV1_IMG                            = 0x8C01
	BGR_INTEGER                                                = 0x8D9A
	TEXCOORD1_BIT_PGI                                          = 0x10000000
	TEXTURE_3D                                                 = 0x806F
	T2F_IUI_N3F_V2F_EXT                                        = 0x81B3
	MULTISAMPLE_3DFX                                           = 0x86B2
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_EXT                       = 0x8C29
	SPRITE_MODE_SGIX                                           = 0x8149
	MODELVIEW8_ARB                                             = 0x8728
	RENDERBUFFER_GREEN_SIZE_EXT                                = 0x8D51
	RESAMPLE_ZERO_FILL_OML                                     = 0x8987
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	SECONDARY_COLOR_ARRAY_TYPE_EXT                             = 0x845B
	DRAW_BUFFER6                                               = 0x882B
	SCREEN_COORDINATES_REND                                    = 0x8490
	VIEW_CLASS_16_BITS                                         = 0x82CA
	DECR_WRAP                                                  = 0x8508
	SRC1_ALPHA                                                 = 0x8589
	WRITE_DISCARD_NV                                           = 0x88BE
	LOGIC_OP_MODE                                              = 0x0BF0
	CMYK_EXT                                                   = 0x800C
	TEXTURE_FILTER4_SIZE_SGIS                                  = 0x8147
	PER_STAGE_CONSTANTS_NV                                     = 0x8535
	OPERAND2_ALPHA_EXT                                         = 0x859A
	RG16I                                                      = 0x8239
	TEXTURE_BINDING_RECTANGLE_NV                               = 0x84F6
	SIGNED_RGB8_UNSIGNED_ALPHA8_NV                             = 0x870D
	DEPTH_COMPONENT32F_NV                                      = 0x8DAB
	VERTEX_ATTRIB_ARRAY6_NV                                    = 0x8656
	DU8DV8_ATI                                                 = 0x877A
	REG_5_ATI                                                  = 0x8926
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	SGIX_vertex_preclip                                        = 1
	ELEMENT_ARRAY_POINTER_ATI                                  = 0x876A
	DRAW_BUFFER14_NV                                           = 0x8833
	READ_PIXEL_DATA_RANGE_NV                                   = 0x8879
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	VARIABLE_E_NV                                              = 0x8527
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	UNSIGNED_INT_IMAGE_1D_ARRAY_EXT                            = 0x9068
	SGI_color_matrix                                           = 1
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	GL_2PASS_1_SGIS                                            = 0x80A3
	RGB9_E5_EXT                                                = 0x8C3D
	BLEND_SRC_RGB_OES                                          = 0x80C9
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	TEXTURE10_ARB                                              = 0x84CA
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	TEXTURE_MAG_FILTER                                         = 0x2800
	R3_G3_B2                                                   = 0x2A10
	DOT3_RGB                                                   = 0x86AE
	PRESERVE_ATI                                               = 0x8762
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	INTENSITY16                                                = 0x804D
	RGB5                                                       = 0x8050
	VIEW_CLASS_128_BITS                                        = 0x82C4
	OUTPUT_TEXTURE_COORD14_EXT                                 = 0x87AB
	INT_IMAGE_BUFFER_EXT                                       = 0x905C
	CULL_VERTEX_EXT                                            = 0x81AA
	SIGNALED_APPLE                                             = 0x9119
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	HALF_FLOAT_NV                                              = 0x140B
	VERTEX_ARRAY_SIZE_EXT                                      = 0x807A
	MAX_HEIGHT                                                 = 0x827F
	EVAL_VERTEX_ATTRIB5_NV                                     = 0x86CB
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	RIGHT                                                      = 0x0407
	EXP                                                        = 0x0800
	REPLACE_MIDDLE_SUN                                         = 0x0002
	VIEW_CLASS_96_BITS                                         = 0x82C5
	CONVOLUTION_HINT_SGIX                                      = 0x8316
	OUTPUT_TEXTURE_COORD25_EXT                                 = 0x87B6
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	TRANSPOSE_TEXTURE_MATRIX_ARB                               = 0x84E5
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS_EXT             = 0x8C80
	SAMPLER_EXTERNAL_OES                                       = 0x8D66
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	TEXTURE6                                                   = 0x84C6
	SIGNED_HILO8_NV                                            = 0x885F
	UNIFORM_OFFSET                                             = 0x8A3B
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	POLYGON_OFFSET_FACTOR_EXT                                  = 0x8038
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	COMBINER_MAPPING_NV                                        = 0x8543
	COLOR_ALPHA_PAIRING_ATI                                    = 0x8975
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	FOG_HINT                                                   = 0x0C54
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX                        = 0x817B
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	QUERY_WAIT                                                 = 0x8E13
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	FIXED_OES                                                  = 0x140C
	TEXTURE_BASE_LEVEL_SGIS                                    = 0x813C
	PROGRAM_LENGTH_NV                                          = 0x8627
	MIRROR_CLAMP_EXT                                           = 0x8742
	TRACE_ARRAYS_BIT_MESA                                      = 0x0004
	PASS_THROUGH_TOKEN                                         = 0x0700
	CLAMP_TO_BORDER_SGIS                                       = 0x812D
	DEPTH_PASS_INSTRUMENT_SGIX                                 = 0x8310
	TEXTURE21                                                  = 0x84D5
	VERTEX_PROGRAM_POINT_SIZE_NV                               = 0x8642
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	COLOR_ATTACHMENT8_EXT                                      = 0x8CE8
	DEPTH_CLAMP_FAR_AMD                                        = 0x901F
	IMAGE_3D                                                   = 0x904E
	ATTRIB_ARRAY_POINTER_NV                                    = 0x8645
	COLOR_ATTACHMENT11_EXT                                     = 0x8CEB
	LUMINANCE32UI_EXT                                          = 0x8D74
	VIDEO_CAPTURE_FRAME_WIDTH_NV                               = 0x9038
	INT_IMAGE_2D                                               = 0x9058
	COLOR_EXT                                                  = 0x1800
	DRAW_BUFFER0_ARB                                           = 0x8825
	OBJECT_SHADER_SOURCE_LENGTH_ARB                            = 0x8B88
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER_EXT                     = 0x8CDB
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	OP_MULTIPLY_MATRIX_EXT                                     = 0x8798
	PROGRAM_BINARY_FORMATS_OES                                 = 0x87FF
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	INTENSITY                                                  = 0x8049
	GL_422_AVERAGE_EXT                                         = 0x80CE
	COLOR_TABLE                                                = 0x80D0
	COMPRESSED_RED                                             = 0x8225
	LAYER_NV                                                   = 0x8DAA
	VIDEO_CAPTURE_FRAME_HEIGHT_NV                              = 0x9039
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	TEXTURE_BORDER_COLOR                                       = 0x1004
	T2F_C4UB_V3F                                               = 0x2A29
	SHADER_OPERATION_NV                                        = 0x86DF
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	CONTINUOUS_AMD                                             = 0x9007
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	C4UB_V2F                                                   = 0x2A22
	RGB8_EXT                                                   = 0x8051
	UNPACK_SUBSAMPLE_RATE_SGIX                                 = 0x85A1
	TRANSFORM_FEEDBACK_BUFFER_START_EXT                        = 0x8C84
	STENCIL_ATTACHMENT_EXT                                     = 0x8D20
	ALLOW_DRAW_MEM_HINT_PGI                                    = 0x1A211
	SET                                                        = 0x150F
	TEXTURE_LOD_BIAS_T_SGIX                                    = 0x818F
	DRAW_BUFFER9                                               = 0x882E
	MATRIX15_ARB                                               = 0x88CF
	FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT                 = 0x8D6C
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	STATIC_VERTEX_ARRAY_IBM                                    = 103061
	MAX_PROGRAM_PATCH_ATTRIBS_NV                               = 0x86D8
	OBJECT_TYPE_APPLE                                          = 0x9112
	CULL_VERTEX_IBM                                            = 103050
	TEXTURE_1D_BINDING_EXT                                     = 0x8068
	BUMP_TEX_UNITS_ATI                                         = 0x8778
	SYNC_GPU_COMMANDS_COMPLETE_APPLE                           = 0x9117
	BLEND_COLOR                                                = 0x8005
	POST_CONVOLUTION_RED_SCALE_EXT                             = 0x801C
	MAP2_VERTEX_ATTRIB3_4_NV                                   = 0x8673
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	VERTEX_SHADER_OPTIMIZED_EXT                                = 0x87D4
	DRAW_BUFFER0_NV                                            = 0x8825
	COLOR_ATTACHMENT2                                          = 0x8CE2
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS                            = 0x919A
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	BUFFER_BINDING                                             = 0x9302
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	DRAW_BUFFER12_ARB                                          = 0x8831
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	CONTEXT_PROFILE_MASK                                       = 0x9126
	POINT_SIZE_RANGE                                           = 0x0B12
	STENCIL_REF                                                = 0x0B97
	BGR                                                        = 0x80E0
	OP_FLOOR_EXT                                               = 0x878F
	DEPTH_COMPONENT24_SGIX                                     = 0x81A6
	COLOR_ATTACHMENT5_EXT                                      = 0x8CE5
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	NICEST                                                     = 0x1102
	MODULATE                                                   = 0x2100
	CLIP_PLANE1                                                = 0x3001
	RGB_S3TC                                                   = 0x83A0
	OPERAND2_ALPHA_ARB                                         = 0x859A
	REG_27_ATI                                                 = 0x893C
	FLOAT_MAT2x3                                               = 0x8B65
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	TEXTURE_CLIPMAP_DEPTH_SGIX                                 = 0x8176
	STENCIL_COMPONENTS                                         = 0x8285
	SECONDARY_COLOR_ARRAY_EXT                                  = 0x845E
	TEXTURE13_ARB                                              = 0x84CD
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	PROXY_TEXTURE_2D_EXT                                       = 0x8064
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	ASYNC_TEX_IMAGE_SGIX                                       = 0x835C
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	ARRAY_SIZE                                                 = 0x92FB
	STATIC_COPY_ARB                                            = 0x88E6
	TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE                      = 0x93A0
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	COMBINER7_NV                                               = 0x8557
	LOCAL_CONSTANT_VALUE_EXT                                   = 0x87EC
	COVERAGE_BUFFERS_NV                                        = 0x8ED3
	POINT_TOKEN                                                = 0x0701
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	DEPTH_STENCIL_OES                                          = 0x84F9
	FRAGMENT_PROGRAM_NV                                        = 0x8870
	BOOL_VEC3_ARB                                              = 0x8B58
	MAX_LABEL_LENGTH                                           = 0x82E8
	COMBINE4_NV                                                = 0x8503
	MAX_PROGRAM_TEX_INSTRUCTIONS_ARB                           = 0x880C
	MATRIX18_ARB                                               = 0x88D2
	RGBA8UI                                                    = 0x8D7C
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	SIGNED_RGB_UNSIGNED_ALPHA_NV                               = 0x870C
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	MAX_SPARSE_TEXTURE_SIZE_AMD                                = 0x9198
	COLOR_BUFFER_BIT                                           = 0x00004000
	POLYGON_SMOOTH                                             = 0x0B41
	LINEAR_ATTENUATION                                         = 0x1208
	RED_EXT                                                    = 0x1903
	INTENSITY8                                                 = 0x804B
	COLOR_TABLE_RED_SIZE                                       = 0x80DA
	DRAW_BUFFER0                                               = 0x8825
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	FIRST_TO_REST_NV                                           = 0x90AF
	STACK_OVERFLOW                                             = 0x0503
	MODELVIEW                                                  = 0x1700
	ADD_SIGNED_EXT                                             = 0x8574
	NEGATIVE_W_EXT                                             = 0x87DC
	CON_24_ATI                                                 = 0x8959
	CLIP_NEAR_HINT_PGI                                         = 0x1A220
	TEXTURE_LEQUAL_R_SGIX                                      = 0x819C
	SGIX_icc_texture                                           = 1
	GL_4PASS_1_SGIS                                            = 0x80A5
	OUTPUT_TEXTURE_COORD19_EXT                                 = 0x87B0
	ELEMENT_ARRAY_APPLE                                        = 0x8A0C
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	DOUBLE_MAT2x3_EXT                                          = 0x8F49
	SAMPLE_BUFFERS_ARB                                         = 0x80A8
	POINT_FADE_THRESHOLD_SIZE_EXT                              = 0x8128
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	TEXTURE_BINDING_CUBE_MAP_OES                               = 0x8514
	DRAW_BUFFER1_ATI                                           = 0x8826
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	VERTEX_ATTRIB_ARRAY14_NV                                   = 0x865E
	INVALID_INDEX                                              = 0xFFFFFFFF
	MIN_LOD_WARNING_AMD                                        = 0x919C
	TRANSPOSE_NV                                               = 0x862C
	TEXTURE_FREE_MEMORY_ATI                                    = 0x87FC
	MAX_PROGRAM_TEX_INDIRECTIONS_ARB                           = 0x880D
	STATIC_READ_ARB                                            = 0x88E5
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	POST_CONVOLUTION_GREEN_SCALE                               = 0x801D
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	COMPRESSED_RGBA_ARB                                        = 0x84EE
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	IMAGE_BINDING_LEVEL_EXT                                    = 0x8F3B
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	SHADER_IMAGE_STORE                                         = 0x82A5
	OUTPUT_TEXTURE_COORD15_EXT                                 = 0x87AC
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	RGB_INTEGER                                                = 0x8D98
	LUMINANCE_SNORM                                            = 0x9011
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	PROGRAM_TEX_INDIRECTIONS_ARB                               = 0x8807
	FLOAT_RG16_NV                                              = 0x8886
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	HIGH_FLOAT                                                 = 0x8DF2
	PN_TRIANGLES_POINT_MODE_CUBIC_ATI                          = 0x87F6
	REG_4_ATI                                                  = 0x8925
	SQUARE_NV                                                  = 0x90A3
	EXT_cmyka                                                  = 1
	RETURN                                                     = 0x0102
	TEXTURE14                                                  = 0x84CE
	INVARIANT_VALUE_EXT                                        = 0x87EA
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	TEXTURE_FLOAT_COMPONENTS_NV                                = 0x888C
	COMPRESSED_RGB_PVRTC_4BPPV1_IMG                            = 0x8C00
	PROXY_TEXTURE_COLOR_TABLE_SGI                              = 0x80BD
	VERTEX_ATTRIB_MAP2_SIZE_APPLE                              = 0x8A06
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV                         = 0x8C4D
	SYNC_CONDITION_APPLE                                       = 0x9113
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB                            = 0x8518
	SHADER_CONSISTENT_NV                                       = 0x86DD
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_NV                     = 0x8852
	OFFSET_PROJECTIVE_TEXTURE_RECTANGLE_SCALE_NV               = 0x8853
	IMAGE_1D_ARRAY_EXT                                         = 0x9052
	DETAIL_TEXTURE_LEVEL_SGIS                                  = 0x809A
	VERTEX_WEIGHT_ARRAY_TYPE_EXT                               = 0x850E
	RENDERBUFFER_RED_SIZE_EXT                                  = 0x8D50
	FRONT_FACE                                                 = 0x0B46
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	SAMPLER_2D_RECT_SHADOW_ARB                                 = 0x8B64
	SHADER_IMAGE_ACCESS_BARRIER_BIT_EXT                        = 0x00000020
	CURRENT_SECONDARY_COLOR_EXT                                = 0x8459
	COMBINER_CD_DOT_PRODUCT_NV                                 = 0x8546
	COMBINE_EXT                                                = 0x8570
	MATRIX26_ARB                                               = 0x88DA
	REG_20_ATI                                                 = 0x8935
	SAMPLER_CUBE_SHADOW_NV                                     = 0x8DC5
	IMAGE_3D_EXT                                               = 0x904E
	PACK_COMPRESSED_SIZE_SGIX                                  = 0x831C
	TEXTURE26_ARB                                              = 0x84DA
	VERTEX_ATTRIB_ARRAY11_NV                                   = 0x865B
	OP_MIN_EXT                                                 = 0x878B
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	LINE_SMOOTH_HINT                                           = 0x0C52
	LIGHT3                                                     = 0x4003
	FOG_FUNC_POINTS_SGIS                                       = 0x812B
	PIXEL_TRANSFORM_2D_STACK_DEPTH_EXT                         = 0x8336
	STORAGE_CACHED_APPLE                                       = 0x85BE
	DRAW_BUFFER1_ARB                                           = 0x8826
	UNPACK_RESAMPLE_OML                                        = 0x8985
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	TRIANGLE_STRIP_ADJACENCY_EXT                               = 0x000D
	LIGHT1                                                     = 0x4001
	TEXTURE_ALPHA_SIZE_EXT                                     = 0x805F
	GL_2PASS_0_SGIS                                            = 0x80A2
	COMBINER_SCALE_NV                                          = 0x8548
	REG_6_ATI                                                  = 0x8927
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	UNIFORM                                                    = 0x92E1
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	UNSIGNED_INT_24_8_EXT                                      = 0x84FA
	DEPTH32F_STENCIL8_NV                                       = 0x8DAC
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	INTENSITY_EXT                                              = 0x8049
	GL_4PASS_3_EXT                                             = 0x80A7
	TEXTURE_1D_STACK_MESAX                                     = 0x8759
	DEPTH_COMPONENT32_ARB                                      = 0x81A7
	NUM_PROGRAM_BINARY_FORMATS_OES                             = 0x87FE
	COMPUTE_SHADER                                             = 0x91B9
	EXT_polygon_offset                                         = 1
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	TEXTURE10                                                  = 0x84CA
	SOURCE2_RGB_EXT                                            = 0x8582
	OUTPUT_TEXTURE_COORD13_EXT                                 = 0x87AA
	MAX_PALETTE_MATRICES_ARB                                   = 0x8842
	MATRIX_INDEX_ARRAY_POINTER_ARB                             = 0x8849
	FRAMEBUFFER_UNSUPPORTED_EXT                                = 0x8CDD
	COMBINER_AB_DOT_PRODUCT_NV                                 = 0x8545
	PRIMARY_COLOR_EXT                                          = 0x8577
	DRAW_BUFFER12                                              = 0x8831
	SAMPLER_1D_ARRAY_EXT                                       = 0x8DC0
	DISCRETE_AMD                                               = 0x9006
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	SUCCESS_NV                                                 = 0x902F
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	OUTPUT_TEXTURE_COORD31_EXT                                 = 0x87BC
	CON_12_ATI                                                 = 0x894D
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	MAX_MAP_TESSELLATION_NV                                    = 0x86D6
	VERTEX_SHADER_LOCALS_EXT                                   = 0x87D3
	WRITE_ONLY                                                 = 0x88B9
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	CURRENT_RASTER_POSITION                                    = 0x0B07
	RESAMPLE_DECIMATE_SGIX                                     = 0x8430
	EVAL_VERTEX_ATTRIB11_NV                                    = 0x86D1
	MATRIX_INDEX_ARRAY_STRIDE_OES                              = 0x8848
	MAX_MULTIVIEW_BUFFERS_EXT                                  = 0x90F2
	SRGB_DECODE_ARB                                            = 0x8299
	LEFT                                                       = 0x0406
	TEXTURE_DEFORMATION_BIT_SGIX                               = 0x00000001
	PIXEL_TEX_GEN_MODE_SGIX                                    = 0x832B
	WRAP_BORDER_SUN                                            = 0x81D4
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	WRITE_ONLY_OES                                             = 0x88B9
	EXT_texture_object                                         = 1
	INDEX_TEST_REF_EXT                                         = 0x81B7
	MATRIX0_NV                                                 = 0x8630
	SIGNED_HILO16_NV                                           = 0x86FA
	PN_TRIANGLES_NORMAL_MODE_ATI                               = 0x87F3
	DOT_PRODUCT_PASS_THROUGH_NV                                = 0x885B
	CLAMP_VERTEX_COLOR                                         = 0x891A
	TRANSFORM_FEEDBACK_BUFFER_BINDING_NV                       = 0x8C8F
	CONVEX_HULL_NV                                             = 0x908B
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	TEXTURE_1D_ARRAY                                           = 0x8C18
	BACK_PRIMARY_COLOR_NV                                      = 0x8C77
	COLOR_ATTACHMENT3_EXT                                      = 0x8CE3
	COLOR_ATTACHMENT2_NV                                       = 0x8CE2
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	BLEND_SRC_RGB_EXT                                          = 0x80C9
	PIXEL_TRANSFORM_2D_EXT                                     = 0x8330
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	INVALID_OPERATION                                          = 0x0502
	DOUBLE_EXT                                                 = 0x140A
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP                  = 0x8163
	REFLECTION_MAP                                             = 0x8512
	STENCIL_BACK_OP_VALUE_AMD                                  = 0x874D
	LUMINANCE16F_ARB                                           = 0x881E
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	COLOR                                                      = 0x1800
	DUAL_INTENSITY8_SGIS                                       = 0x8119
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS                           = 0x8355
	OFFSET_HILO_TEXTURE_RECTANGLE_NV                           = 0x8855
	FLOAT_RGBA16_NV                                            = 0x888A
	ELEMENT_ARRAY_BUFFER_BINDING_ARB                           = 0x8895
	SRGB_ALPHA_EXT                                             = 0x8C42
	MOVE_TO_NV                                                 = 0x02
	TEXTURE_RESIDENT_EXT                                       = 0x8067
	TEXTURE_COMPARE_SGIX                                       = 0x819A
	PALETTE4_RGBA4_OES                                         = 0x8B93
	VIDEO_CAPTURE_FIELD_UPPER_HEIGHT_NV                        = 0x903A
	PATH_GEN_COEFF_NV                                          = 0x90B1
	IUI_V3F_EXT                                                = 0x81AE
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	FRAGMENT_PROGRAM_ARB                                       = 0x8804
	SAMPLES_PASSED                                             = 0x8914
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM                      = 103087
	SLUMINANCE8                                                = 0x8C47
	ALPHA_INTEGER                                              = 0x8D97
	INTENSITY16_SNORM                                          = 0x901B
	FACTOR_MAX_AMD                                             = 0x901D
	TYPE                                                       = 0x92FA
	DRAW_BUFFER15_ARB                                          = 0x8834
	PIXEL_UNPACK_BUFFER_BINDING_EXT                            = 0x88EF
	FRAGMENT_SHADER_DERIVATIVE_HINT_ARB                        = 0x8B8B
	COMPRESSED_LUMINANCE_LATC1_EXT                             = 0x8C70
	IMAGE_2D_RECT                                              = 0x904F
	DEPTH_WRITEMASK                                            = 0x0B72
	MAP1_VERTEX_4                                              = 0x0D98
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX                         = 0x8174
	ATTENUATION_EXT                                            = 0x834D
	DRAW_BUFFER6_ATI                                           = 0x882B
	R8_SNORM                                                   = 0x8F94
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	MAT_SHININESS_BIT_PGI                                      = 0x02000000
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	TEXTURE_ALPHA_TYPE_ARB                                     = 0x8C13
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	PROGRAM_BINARY_ANGLE                                       = 0x93A6
	EXT_rescale_normal                                         = 1
	SGIS_texture_border_clamp                                  = 1
	LIGHT0                                                     = 0x4000
	PIXEL_TILE_BEST_ALIGNMENT_SGIX                             = 0x813E
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	COMPRESSED_ALPHA_ARB                                       = 0x84E9
	DRAW_BUFFER15                                              = 0x8834
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	INDEX_BITS                                                 = 0x0D51
	POST_COLOR_MATRIX_ALPHA_SCALE                              = 0x80B7
	COLOR_TABLE_SCALE_SGI                                      = 0x80D6
	MAX_ELEMENTS_INDICES_EXT                                   = 0x80E9
	EYE_DISTANCE_TO_LINE_SGIS                                  = 0x81F2
	OUTPUT_TEXTURE_COORD8_EXT                                  = 0x87A5
	MEDIUM_FLOAT                                               = 0x8DF1
	COVERAGE_AUTOMATIC_NV                                      = 0x8ED7
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	INTERLACE_SGIX                                             = 0x8094
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	CURRENT_WEIGHT_ARB                                         = 0x86A8
	MODULATE_COLOR_IMG                                         = 0x8C04
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	VERTEX4_BIT_PGI                                            = 0x00000008
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	SCALE_BY_FOUR_NV                                           = 0x853F
	DEPENDENT_RGB_TEXTURE_3D_NV                                = 0x8859
	SYNC_OBJECT_APPLE                                          = 0x8A53
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	LIGHTING_BIT                                               = 0x00000040
	CURRENT_RASTER_INDEX                                       = 0x0B05
	CURRENT_FOG_COORDINATE_EXT                                 = 0x8453
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	MATRIX28_ARB                                               = 0x88DC
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	COLOR_ATTACHMENT9                                          = 0x8CE9
	INVALID_VALUE                                              = 0x0501
	IMAGE_ROTATE_ORIGIN_X_HP                                   = 0x815A
	VERTEX_ARRAY_RANGE_APPLE                                   = 0x851D
	DSDT_MAG_VIB_NV                                            = 0x86F7
	ARRAY_OBJECT_BUFFER_ATI                                    = 0x8766
	RGBA_FLOAT_MODE_ATI                                        = 0x8820
	BACK_RIGHT                                                 = 0x0403
	PACK_CMYK_HINT_EXT                                         = 0x800E
	PN_TRIANGLES_POINT_MODE_ATI                                = 0x87F2
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	SGIX_ir_instrument1                                        = 1
	EYE_PLANE                                                  = 0x2502
	RGBA16F                                                    = 0x881A
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	POINT_SIZE_MAX_SGIS                                        = 0x8127
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	COMPRESSED_RGB                                             = 0x84ED
	LOWER_LEFT                                                 = 0x8CA1
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	MAP2_VERTEX_ATTRIB15_4_NV                                  = 0x867F
	FLOAT_RG_NV                                                = 0x8881
	VERTEX_ATTRIB_MAP1_SIZE_APPLE                              = 0x8A02
	SLUMINANCE_EXT                                             = 0x8C46
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	DEBUG_CATEGORY_API_ERROR_AMD                               = 0x9149
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	MODELVIEW_MATRIX                                           = 0x0BA6
	MAX_PROGRAM_NATIVE_ALU_INSTRUCTIONS_ARB                    = 0x880E
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	RGB32I                                                     = 0x8D83
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	FILE_NAME_NV                                               = 0x9074
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	VERSION_1_1                                                = 1
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	FLOAT_RGB_NV                                               = 0x8882
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	DOUBLE_MAT4                                                = 0x8F48
	VERSION_1_3                                                = 1
	FLOAT_R16_NV                                               = 0x8884
	CONDITION_SATISFIED                                        = 0x911C
	SGIX_flush_raster                                          = 1
	GEOMETRY_DEFORMATION_BIT_SGIX                              = 0x00000002
	FIXED                                                      = 0x140C
	NOR                                                        = 0x1508
	TEXTURE_WRAP_S                                             = 0x2802
	IMAGE_ROTATE_ANGLE_HP                                      = 0x8159
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	SIGNED_LUMINANCE_ALPHA_NV                                  = 0x8703
	STENCIL_ATTACHMENT                                         = 0x8D20
	TEXTURE_2D                                                 = 0x0DE1
	FRAGMENT_LIGHT4_SGIX                                       = 0x8410
	NORMAL_MAP                                                 = 0x8511
	ARRAY_BUFFER                                               = 0x8892
	LAST_VIDEO_CAPTURE_STATUS_NV                               = 0x9027
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	DEPTH_COMPONENTS                                           = 0x8284
	FENCE_STATUS_NV                                            = 0x84F3
	SHADER_BINARY_VIV                                          = 0x8FC4
	FUNC_SUBTRACT_EXT                                          = 0x800A
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	INT_IMAGE_2D_EXT                                           = 0x9058
	IMAGE_2D                                                   = 0x904D
	FRAGMENT_NORMAL_EXT                                        = 0x834A
	PIXEL_SUBSAMPLE_4444_SGIX                                  = 0x85A2
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	TRIANGLES_ADJACENCY_EXT                                    = 0x000C
	RGBA                                                       = 0x1908
	LUMINANCE_ALPHA_FLOAT16_APPLE                              = 0x881F
	CON_5_ATI                                                  = 0x8946
	COMP_BIT_ATI                                               = 0x00000002
	TEXTURE_SWIZZLE_A_EXT                                      = 0x8E45
	INT64_VEC4_NV                                              = 0x8FEB
	SGIX_subsample                                             = 1
	LUMINANCE8                                                 = 0x8040
	UNPACK_IMAGE_DEPTH_SGIS                                    = 0x8133
	INDEX_MATERIAL_FACE_EXT                                    = 0x81BA
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	FOG_COORDINATE_ARRAY_TYPE_EXT                              = 0x8454
	TRACE_TEXTURES_BIT_MESA                                    = 0x0008
	FLOAT_MAT2_ARB                                             = 0x8B5A
	TRIANGLE_STRIP                                             = 0x0005
	NORMAL_ARRAY_TYPE_EXT                                      = 0x807E
	YCRCB_422_SGIX                                             = 0x81BB
	PALETTE8_RGBA4_OES                                         = 0x8B98
	TEXTURE_FORMAT_QCOM                                        = 0x8BD6
	COLOR_ATTACHMENT3                                          = 0x8CE3
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	OBJECT_LINEAR                                              = 0x2401
	DRAW_BUFFER14_ARB                                          = 0x8833
	CONSERVE_MEMORY_HINT_PGI                                   = 0x1A1FD
	DOMAIN                                                     = 0x0A02
	RELEASED_APPLE                                             = 0x8A19
	TEXTURE_CROP_RECT_OES                                      = 0x8B9D
	INT_SAMPLER_2D_RECT_EXT                                    = 0x8DCD
	PATH_FOG_GEN_MODE_NV                                       = 0x90AC
	ARB_imaging                                                = 1
	FOG_END                                                    = 0x0B64
	DISPLAY_LIST                                               = 0x82E7
	COMPRESSED_RGBA                                            = 0x84EE
	NEGATIVE_ONE_EXT                                           = 0x87DF
	RGBA32UI_EXT                                               = 0x8D70
	MAX_IMAGE_SAMPLES_EXT                                      = 0x906D
	PROXY_TEXTURE_2D                                           = 0x8064
	COLOR_ARRAY_EXT                                            = 0x8076
	OBJECT_BUFFER_SIZE_ATI                                     = 0x8764
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	POST_COLOR_MATRIX_RED_SCALE_SGI                            = 0x80B4
	DEPTH_COMPONENT16_ARB                                      = 0x81A5
	REG_11_ATI                                                 = 0x892C
	FLOAT_MAT3                                                 = 0x8B5B
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	MAX_SAMPLES                                                = 0x8D57
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_OES                     = 0x8CD1
	OBJECT_INFO_LOG_LENGTH_ARB                                 = 0x8B84
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	MAX_VARYING_VECTORS                                        = 0x8DFC
	AVERAGE_EXT                                                = 0x8335
	ALPHA32F_ARB                                               = 0x8816
	SEPARATE_ATTRIBS                                           = 0x8C8D
	PACK_REVERSE_ROW_ORDER_ANGLE                               = 0x93A4
	EDGE_FLAG                                                  = 0x0B43
	RENDERER                                                   = 0x1F01
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	OP_DOT4_EXT                                                = 0x8785
	OBJECT_VALIDATE_STATUS_ARB                                 = 0x8B83
	RGBA16UI_EXT                                               = 0x8D76
	TRANSLATE_3D_NV                                            = 0x9091
	COMPUTE_PROGRAM_NV                                         = 0x90FB
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	INTERPOLATE                                                = 0x8575
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	ARRAY_BUFFER_ARB                                           = 0x8892
	BUFFER_MAP_POINTER_OES                                     = 0x88BD
	COLOR_ATTACHMENT14_EXT                                     = 0x8CEE
	SGIX_list_priority                                         = 1
	NORMALIZE                                                  = 0x0BA1
	PACK_LSB_FIRST                                             = 0x0D01
	INDEX_ARRAY_EXT                                            = 0x8077
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	LINEAR_CLIPMAP_NEAREST_SGIX                                = 0x844F
	OFFSET_TEXTURE_2D_NV                                       = 0x86E8
	INT_SAMPLER_CUBE                                           = 0x8DCC
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS_EXT          = 0x8F39
	DEBUG_LOGGED_MESSAGES_AMD                                  = 0x9145
	INDEX_MODE                                                 = 0x0C30
	MAX_EXT                                                    = 0x8008
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	SAMPLE_MASK_VALUE_EXT                                      = 0x80AA
	INDEX_MATERIAL_PARAMETER_EXT                               = 0x81B9
	OPERAND3_RGB_NV                                            = 0x8593
	DOT_PRODUCT_TEXTURE_1D_NV                                  = 0x885C
	MAX_PROGRAM_IF_DEPTH_NV                                    = 0x88F6
	TANGENT_ARRAY_TYPE_EXT                                     = 0x843E
	RGB565                                                     = 0x8D62
	COUNT_UP_NV                                                = 0x9088
	TEXTURE_COMPRESSION_HINT_ARB                               = 0x84EF
	COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT                      = 0x8DBE
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	AND                                                        = 0x1501
	SEPARABLE_2D_EXT                                           = 0x8012
	PACK_IMAGE_HEIGHT_EXT                                      = 0x806C
	MAP1_VERTEX_ATTRIB6_4_NV                                   = 0x8666
	INTENSITY32F_ARB                                           = 0x8817
	ADD_ATI                                                    = 0x8963
	OUTPUT_TEXTURE_COORD16_EXT                                 = 0x87AD
	CONSTANT_ALPHA                                             = 0x8003
	REG_28_ATI                                                 = 0x893D
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	VIDEO_COLOR_CONVERSION_OFFSET_NV                           = 0x902C
	MAT_COLOR_INDEXES_BIT_PGI                                  = 0x01000000
	FOG_DISTANCE_MODE_NV                                       = 0x855A
	MODELVIEW18_ARB                                            = 0x8732
	INT16_VEC3_NV                                              = 0x8FE6
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	MAX_SAMPLES_IMG                                            = 0x9135
	GL_4PASS_3_SGIS                                            = 0x80A7
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX                         = 0x8409
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	COLOR_ATTACHMENT0_NV                                       = 0x8CE0
	BGR_INTEGER_EXT                                            = 0x8D9A
	OPERAND1_ALPHA_ARB                                         = 0x8599
	MATRIX4_NV                                                 = 0x8634
	DRAW_BUFFER0_ATI                                           = 0x8825
	UNSIGNED_INT8_NV                                           = 0x8FEC
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	RENDERBUFFER                                               = 0x8D41
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT                   = 0x8CD4
	HISTOGRAM                                                  = 0x8024
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	INT_VEC2_ARB                                               = 0x8B53
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	TEXTURE_COORD_ARRAY_POINTER_EXT                            = 0x8092
	IMAGE_TRANSLATE_Y_HP                                       = 0x8158
	OPERAND0_RGB_EXT                                           = 0x8590
	MATRIX_PALETTE_ARB                                         = 0x8840
	PERCENTAGE_AMD                                             = 0x8BC3
	RENDERBUFFER_SAMPLES_IMG                                   = 0x9133
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	TEXTURE8                                                   = 0x84C8
	COMPRESSED_RGB_ARB                                         = 0x84ED
	TEXTURE_DEPTH_TYPE_ARB                                     = 0x8C16
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	ALL_BARRIER_BITS_EXT                                       = 0xFFFFFFFF
	VERTEX_WEIGHT_ARRAY_STRIDE_EXT                             = 0x850F
	STORAGE_PRIVATE_APPLE                                      = 0x85BD
	DEPTH_TEXTURE_MODE                                         = 0x884B
	MATRIX31_ARB                                               = 0x88DF
	SUB_ATI                                                    = 0x8965
	FORMAT_SUBSAMPLE_244_244_OML                               = 0x8983
	TEXTURE_COORD_NV                                           = 0x8C79
	COLOR_ATTACHMENT3_NV                                       = 0x8CE3
	COMMAND_BARRIER_BIT_EXT                                    = 0x00000040
	TEXTURE_WRAP_R_OES                                         = 0x8072
	COLOR_TABLE_GREEN_SIZE_SGI                                 = 0x80DB
	LOCAL_CONSTANT_DATATYPE_EXT                                = 0x87ED
	CON_1_ATI                                                  = 0x8942
	PATH_DASH_OFFSET_NV                                        = 0x907E
	FOG_START                                                  = 0x0B63
	TEXTURE_4D_SGIS                                            = 0x8134
	INDEX_MATERIAL_EXT                                         = 0x81B8
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	TEXTURE0                                                   = 0x84C0
	LINE_TO_NV                                                 = 0x04
	ALPHA_SCALE                                                = 0x0D1C
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	VARIABLE_G_NV                                              = 0x8529
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	OPERAND2_RGB                                               = 0x8592
	ATTRIB_ARRAY_STRIDE_NV                                     = 0x8624
	RGB8_SNORM                                                 = 0x8F96
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	WRITE_PIXEL_DATA_RANGE_NV                                  = 0x8878
	OBJECT_ACTIVE_ATTRIBUTE_MAX_LENGTH_ARB                     = 0x8B8A
	MAX_PROGRAM_TOTAL_OUTPUT_COMPONENTS_NV                     = 0x8C28
	Z6Y10Z6CB10Z6Y10Z6CR10_422_NV                              = 0x9033
	COMPILE                                                    = 0x1300
	MULTISAMPLE                                                = 0x809D
	SAMPLE_COVERAGE_VALUE_ARB                                  = 0x80AA
	FOG_COORDINATE_EXT                                         = 0x8451
	MITER_TRUNCATE_NV                                          = 0x90A8
	DEPTH_COMPONENT24_OES                                      = 0x81A6
	RGB_SCALE                                                  = 0x8573
	OUTPUT_TEXTURE_COORD0_EXT                                  = 0x879D
	INDEX_ARRAY_LIST_IBM                                       = 103073
	COLOR4_BIT_PGI                                             = 0x00020000
	LOW_INT                                                    = 0x8DF3
	NATIVE_GRAPHICS_HANDLE_PGI                                 = 0x1A202
	MULT                                                       = 0x0103
	POST_CONVOLUTION_COLOR_TABLE                               = 0x80D1
	VERTEX_ATTRIB_ARRAY5_NV                                    = 0x8655
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	RGBA32I                                                    = 0x8D82
	ALPHA_INTEGER_EXT                                          = 0x8D97
	TEXTURE_COLOR_SAMPLES_NV                                   = 0x9046
	TRANSLATE_X_NV                                             = 0x908E
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	TEXTURE_FETCH_BARRIER_BIT_EXT                              = 0x00000008
	GL_1PASS_SGIS                                              = 0x80A1
	T2F_IUI_V3F_EXT                                            = 0x81B2
	CUBIC_EXT                                                  = 0x8334
	INVARIANT_EXT                                              = 0x87C2
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	TEXTURE_3D_EXT                                             = 0x806F
	TEXTURE29                                                  = 0x84DD
	UNSIGNED_INT_24_8                                          = 0x84FA
	UNSIGNED_INT_24_8_OES                                      = 0x84FA
	TRACE_OPERATIONS_BIT_MESA                                  = 0x0001
	CND_ATI                                                    = 0x896A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS_EXT                = 0x8C8B
	COLOR_ATTACHMENT1_NV                                       = 0x8CE1
	MAX_COLOR_ATTACHMENTS_NV                                   = 0x8CDF
	COLOR_WRITEMASK                                            = 0x0C23
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	VERTEX_ARRAY_LIST_STRIDE_IBM                               = 103080
	RGB2_EXT                                                   = 0x804E
	VERTEX_STREAM0_ATI                                         = 0x876C
	INTENSITY32I_EXT                                           = 0x8D85
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	SGIX_texture_lod_bias                                      = 1
	VERTEX_WEIGHTING_EXT                                       = 0x8509
	INTERLEAVED_ATTRIBS_NV                                     = 0x8C8C
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	POINT_SIZE                                                 = 0x0B11
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	LUMINANCE16_ALPHA16_EXT                                    = 0x8048
	TEXTURE31                                                  = 0x84DF
	MODELVIEW11_ARB                                            = 0x872B
	COLOR_ATTACHMENT10_NV                                      = 0x8CEA
	LUMINANCE8UI_EXT                                           = 0x8D80
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	EQUIV                                                      = 0x1509
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	REG_8_ATI                                                  = 0x8929
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	MODELVIEW25_ARB                                            = 0x8739
	MODELVIEW26_ARB                                            = 0x873A
	RGB_422_APPLE                                              = 0x8A1F
	RGBA_INTEGER_EXT                                           = 0x8D99
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	PACK_ROW_LENGTH                                            = 0x0D02
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	BLUE_MIN_CLAMP_INGR                                        = 0x8562
	OFFSET_TEXTURE_RECTANGLE_NV                                = 0x864C
	PROGRAM_TEX_INSTRUCTIONS_ARB                               = 0x8806
	RENDERBUFFER_DEPTH_SIZE_EXT                                = 0x8D54
	MAX_IMAGE_UNITS                                            = 0x8F38
	STENCIL_VALUE_MASK                                         = 0x0B93
	BUMP_ENVMAP_ATI                                            = 0x877B
	DEPTH24_STENCIL8_OES                                       = 0x88F0
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	CONSTANT_ARB                                               = 0x8576
	PROGRAM_ERROR_STRING_NV                                    = 0x8874
	COLOR_ATTACHMENT7_EXT                                      = 0x8CE7
	COMPRESSED_LUMINANCE_ARB                                   = 0x84EA
	RED_MAX_CLAMP_INGR                                         = 0x8564
	MAX_PROGRAM_NATIVE_INSTRUCTIONS_ARB                        = 0x88A3
	TEXTURE_COORD_ARRAY                                        = 0x8078
	VARIANT_ARRAY_STRIDE_EXT                                   = 0x87E6
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	UNSIGNED_INT_IMAGE_3D_EXT                                  = 0x9064
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	MAP_WRITE_BIT_EXT                                          = 0x0002
	TEXTURE23                                                  = 0x84D7
	CON_15_ATI                                                 = 0x8950
	MAX_DEEP_3D_TEXTURE_DEPTH_NV                               = 0x90D1
	SGIX_reference_plane                                       = 1
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	OP_MOV_EXT                                                 = 0x8799
	MAX_PROGRAM_NATIVE_PARAMETERS_ARB                          = 0x88AB
	CON_16_ATI                                                 = 0x8951
	INT_2_10_10_10_REV                                         = 0x8D9F
	IMAGE_BINDING_NAME_EXT                                     = 0x8F3A
	NORMAL_ARRAY_POINTER_EXT                                   = 0x808F
	ARRAY_ELEMENT_LOCK_FIRST_EXT                               = 0x81A8
	REG_3_ATI                                                  = 0x8924
	R32I                                                       = 0x8235
	MAP2_BINORMAL_EXT                                          = 0x8447
	VERTEX_ATTRIB_ARRAY_POINTER_ARB                            = 0x8645
	COMPARE_R_TO_TEXTURE_ARB                                   = 0x884E
	TEXTURE_COMPARE_FUNC_EXT                                   = 0x884D
	INT8_VEC3_NV                                               = 0x8FE2
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY_EXT                = 0x906C
	TEXTURE_CUBE_MAP_NEGATIVE_Y_EXT                            = 0x8518
	REFLECTION_MAP_ARB                                         = 0x8512
	PREVIOUS_EXT                                               = 0x8578
	MATRIX30_ARB                                               = 0x88DE
	PROGRAM_RESULT_COMPONENTS_NV                               = 0x8907
	PRIMITIVES_GENERATED_EXT                                   = 0x8C87
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	WAIT_FAILED_APPLE                                          = 0x911D
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	LUMINANCE_ALPHA32UI_EXT                                    = 0x8D75
	WIDE_LINE_HINT_PGI                                         = 0x1A222
	PROXY_TEXTURE_4D_SGIS                                      = 0x8135
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	EVAL_FRACTIONAL_TESSELLATION_NV                            = 0x86C5
	STENCIL_TAG_BITS_EXT                                       = 0x88F2
	SAMPLER_2D_RECT_ARB                                        = 0x8B63
	SRGB_ALPHA                                                 = 0x8C42
	DOUBLE_MAT4_EXT                                            = 0x8F48
	CUBIC_CURVE_TO_NV                                          = 0x0C
	BLEND_EQUATION_ALPHA_EXT                                   = 0x883D
	HALF_FLOAT_ARB                                             = 0x140B
	CLIP_DISTANCE4                                             = 0x3004
	COLOR_INDEX4_EXT                                           = 0x80E4
	QUAD_ALPHA4_SGIS                                           = 0x811E
	PREVIOUS_ARB                                               = 0x8578
	VERTEX_ATTRIB_ARRAY_ENABLED_ARB                            = 0x8622
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	VERTEX_STREAM2_ATI                                         = 0x876E
	OFFSET_HILO_PROJECTIVE_TEXTURE_RECTANGLE_NV                = 0x8857
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	HISTOGRAM_EXT                                              = 0x8024
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	DSDT_NV                                                    = 0x86F5
	ELEMENT_ARRAY_BUFFER_ARB                                   = 0x8893
	RENDERBUFFER_DEPTH_SIZE_OES                                = 0x8D54
	TEXT_FRAGMENT_SHADER_ATI                                   = 0x8200
	TEXTURE_CUBE_MAP_NEGATIVE_X_EXT                            = 0x8516
	OP_SUB_EXT                                                 = 0x8796
	SAMPLER_CUBE_ARB                                           = 0x8B60
	TRANSFORM_FEEDBACK_BUFFER_EXT                              = 0x8C8E
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	MAP2_INDEX                                                 = 0x0DB1
	DEPTH_STENCIL_NV                                           = 0x84F9
	MAX_PROGRAM_ALU_INSTRUCTIONS_ARB                           = 0x880B
	COLOR_ATTACHMENT9_NV                                       = 0x8CE9
	MAP_FLUSH_EXPLICIT_BIT_EXT                                 = 0x0010
	SAMPLE_COVERAGE_INVERT_ARB                                 = 0x80AB
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	PROGRAM_NATIVE_ADDRESS_REGISTERS_ARB                       = 0x88B2
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	FLOAT_MAT4x3                                               = 0x8B6A
	INT64_VEC2_NV                                              = 0x8FE9
	NORMAL_ARRAY_LIST_STRIDE_IBM                               = 103081
	MAX_LIST_NESTING                                           = 0x0B31
	HISTOGRAM_BLUE_SIZE_EXT                                    = 0x802A
	SAMPLE_MASK_EXT                                            = 0x80A0
	OP_CLAMP_EXT                                               = 0x878E
	TEXTURE_SWIZZLE_B_EXT                                      = 0x8E44
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	LUMINANCE32F_ARB                                           = 0x8818
	VERTEX_ATTRIB_ARRAY_DIVISOR_NV                             = 0x88FE
	NUM_FRAGMENT_CONSTANTS_ATI                                 = 0x896F
	ELEMENT_ARRAY_POINTER_APPLE                                = 0x8A0E
	TEXTURE22_ARB                                              = 0x84D6
	SIGNED_NEGATE_NV                                           = 0x853D
	PRIMARY_COLOR_ARB                                          = 0x8577
	CULL_VERTEX_OBJECT_POSITION_EXT                            = 0x81AC
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	PROXY_TEXTURE_2D_ARRAY_EXT                                 = 0x8C1B
	R11F_G11F_B10F                                             = 0x8C3A
	MULTISAMPLE_BIT_ARB                                        = 0x20000000
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	COLOR_MATRIX                                               = 0x80B1
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	PROGRAM_PARAMETER_NV                                       = 0x8644
	VIBRANCE_BIAS_NV                                           = 0x8719
	ANY_SAMPLES_PASSED_CONSERVATIVE_EXT                        = 0x8D6A
	DOUBLE_MAT3x2                                              = 0x8F4B
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	DRAW_PIXEL_TOKEN                                           = 0x0705
	RGBA16_EXT                                                 = 0x805B
	LERP_ATI                                                   = 0x8969
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	DECR_WRAP_OES                                              = 0x8508
	REG_21_ATI                                                 = 0x8936
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM                       = 103086
	INT64_NV                                                   = 0x140E
	SEPARABLE_2D                                               = 0x8012
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	PROXY_TEXTURE_CUBE_MAP_EXT                                 = 0x851B
	REPLACEMENT_CODE_ARRAY_SUN                                 = 0x85C0
	MODELVIEW22_ARB                                            = 0x8736
	DEPTH_COMPONENT32F                                         = 0x8CAC
	FIRST_VERTEX_CONVENTION_EXT                                = 0x8E4D
	EXT_blend_logic_op                                         = 1
	STEREO                                                     = 0x0C33
	MINMAX_FORMAT                                              = 0x802F
	PROGRAM                                                    = 0x82E2
	TEXTURE_CUBE_MAP_NEGATIVE_Z_OES                            = 0x851A
	SWIZZLE_STRQ_DQ_ATI                                        = 0x897B
	BOOL_VEC2_ARB                                              = 0x8B57
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	COPY_PIXEL_TOKEN                                           = 0x0706
	COMBINER0_NV                                               = 0x8550
	INT_IMAGE_1D                                               = 0x9057
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	ACTIVE_TEXTURE_ARB                                         = 0x84E0
	FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES              = 0x8CD4
	STENCIL_FUNC                                               = 0x0B92
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	RGBA_DXT5_S3TC                                             = 0x83A4
	DRAW_BUFFER3_ARB                                           = 0x8828
	TRANSLATE_Y_NV                                             = 0x908F
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	RGBA4_EXT                                                  = 0x8056
	OBJECT_DISTANCE_TO_POINT_SGIS                              = 0x81F1
	UNSIGNED_INT_24_8_MESA                                     = 0x8751
	STENCIL_INDEX1_OES                                         = 0x8D46
	CURRENT_TIME_NV                                            = 0x8E28
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_EXT                      = 0x906B
	RGBA2                                                      = 0x8055
	RG_INTEGER                                                 = 0x8228
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	PIXEL_TRANSFORM_2D_MATRIX_EXT                              = 0x8338
	MATRIX19_ARB                                               = 0x88D3
	R11F_G11F_B10F_EXT                                         = 0x8C3A
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	VERTEX_ARRAY_OBJECT_AMD                                    = 0x9154
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	RG8_EXT                                                    = 0x822B
	COLOR_ATTACHMENT15_EXT                                     = 0x8CEF
	RGBA_INTEGER_MODE_EXT                                      = 0x8D9E
	LINE_LOOP                                                  = 0x0002
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	SAMPLE_PATTERN_SGIS                                        = 0x80AC
	SOURCE2_ALPHA_ARB                                          = 0x858A
	VERTEX_ATTRIB_ARRAY10_NV                                   = 0x865A
	OP_DOT3_EXT                                                = 0x8784
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	EDGE_FLAG_ARRAY_BUFFER_BINDING_ARB                         = 0x889B
	READ_ONLY_ARB                                              = 0x88B8
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	LUMINANCE12_ALPHA4_EXT                                     = 0x8046
	POINT_SIZE_MAX                                             = 0x8127
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	QUERY_RESULT_EXT                                           = 0x8866
	COLOR_ATTACHMENT6                                          = 0x8CE6
	PROVOKING_VERTEX                                           = 0x8E4F
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	SPRITE_OBJECT_ALIGNED_SGIX                                 = 0x814D
	FRAME_NV                                                   = 0x8E26
	TESS_GEN_SPACING                                           = 0x8E77
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	LINES                                                      = 0x0001
	DST_COLOR                                                  = 0x0306
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	LIGHT6                                                     = 0x4006
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI                           = 0x80B3
	TRANSPOSE_PROJECTION_MATRIX_ARB                            = 0x84E4
	TESSELLATION_FACTOR_AMD                                    = 0x9005
	FOG_COORD_ARRAY                                            = 0x8457
	VERTEX_ARRAY_RANGE_NV                                      = 0x851D
	REG_30_ATI                                                 = 0x893F
	ACCUM_BUFFER_BIT                                           = 0x00000200
	TEXTURE_LIGHTING_MODE_HP                                   = 0x8167
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	DRAW_BUFFER_EXT                                            = 0x0C01
	TEXTURE                                                    = 0x1702
	BUFFER_ACCESS_ARB                                          = 0x88BB
	LEQUAL                                                     = 0x0203
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	CLAMP                                                      = 0x2900
	EVAL_VERTEX_ATTRIB15_NV                                    = 0x86D5
	REG_2_ATI                                                  = 0x8923
	IUI_N3F_V2F_EXT                                            = 0x81AF
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	FLOAT_VEC3_ARB                                             = 0x8B51
	INT_SAMPLER_1D_EXT                                         = 0x8DC9
	YCBAYCR8A_4224_NV                                          = 0x9032
	DOT_PRODUCT_DIFFUSE_CUBE_MAP_NV                            = 0x86F1
	STANDARD_FONT_NAME_NV                                      = 0x9072
	EVAL_BIT                                                   = 0x00010000
	TEXTURE29_ARB                                              = 0x84DD
	COMPARE_REF_TO_TEXTURE_EXT                                 = 0x884E
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	TRACE_PIXELS_BIT_MESA                                      = 0x0010
	FOG_COORDINATE_ARRAY_BUFFER_BINDING_ARB                    = 0x889D
	MAX_COLOR_ATTACHMENTS_EXT                                  = 0x8CDF
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM                            = 103085
	UNPACK_ALIGNMENT                                           = 0x0CF5
	TEXTURE_RESIDENT                                           = 0x8067
	SAMPLE_BUFFERS_EXT                                         = 0x80A8
	MODELVIEW6_ARB                                             = 0x8726
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	SHADOW_ATTENUATION_EXT                                     = 0x834E
	TEXTURE_BUFFER_EXT                                         = 0x8C2A
	CLIP_PLANE5                                                = 0x3005
	SAMPLE_ALPHA_TO_COVERAGE_ARB                               = 0x809E
	DEBUG_SOURCE_API                                           = 0x8246
	RECT_NV                                                    = 0xF6
	VERSION                                                    = 0x1F02
	DEPTH_COMPONENT24                                          = 0x81A6
	SECONDARY_COLOR_ARRAY_STRIDE_EXT                           = 0x845C
	MIN_SPARSE_LEVEL_AMD                                       = 0x919B
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	INVALID_FRAMEBUFFER_OPERATION_EXT                          = 0x0506
	SOURCE1_RGB                                                = 0x8581
	COPY_WRITE_BUFFER                                          = 0x8F37
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	GREEN_INTEGER                                              = 0x8D95
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	VERTEX_BLEND_ARB                                           = 0x86A7
	WEIGHT_ARRAY_POINTER_OES                                   = 0x86AC
	BUFFER_SIZE                                                = 0x8764
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	PROGRAM_ERROR_STRING_ARB                                   = 0x8874
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	HALF_BIAS_NEGATE_NV                                        = 0x853B
	MAX_VARYING_FLOATS                                         = 0x8B4B
	NORMAL_ARRAY_LIST_IBM                                      = 103071
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	PATH_FORMAT_SVG_NV                                         = 0x9070
	LAYOUT_LINEAR_CPU_CACHED_INTEL                             = 2
	TEXTURE_BORDER_COLOR_NV                                    = 0x1004
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	NORMALIZED_RANGE_EXT                                       = 0x87E0
	PERFORMANCE_MONITOR_AMD                                    = 0x9152
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI                    = 0x80D5
	MODELVIEW17_ARB                                            = 0x8731
	TEXTURE_STENCIL_SIZE_EXT                                   = 0x88F1
	NUM_INSTRUCTIONS_TOTAL_ATI                                 = 0x8972
	FOG_COORDINATE_SOURCE                                      = 0x8450
	EVAL_VERTEX_ATTRIB1_NV                                     = 0x86C7
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	FRONT_RIGHT                                                = 0x0401
	TEXTURE_COORD_ARRAY_COUNT_EXT                              = 0x808B
	DUAL_ALPHA16_SGIS                                          = 0x8113
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	ASYNC_HISTOGRAM_SGIX                                       = 0x832C
	READ_PIXEL_DATA_RANGE_LENGTH_NV                            = 0x887B
	QUERY_RESULT_NO_WAIT_AMD                                   = 0x9194
	STACK_UNDERFLOW                                            = 0x0504
	INTENSITY4                                                 = 0x804A
	EMBOSS_LIGHT_NV                                            = 0x855D
	MODELVIEW_MATRIX_FLOAT_AS_INT_BITS_OES                     = 0x898D
	QUERY_NO_WAIT_NV                                           = 0x8E14
	SAMPLE_MASK_INVERT_EXT                                     = 0x80AB
	VARIABLE_D_NV                                              = 0x8526
	MATRIX22_ARB                                               = 0x88D6
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	CONSTANT_BORDER                                            = 0x8151
	SRGB8_EXT                                                  = 0x8C41
	HALF_FLOAT_OES                                             = 0x8D61
	COVERAGE_SAMPLES_NV                                        = 0x8ED4
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	STRICT_LIGHTING_HINT_PGI                                   = 0x1A217
	R                                                          = 0x2002
	CLIP_DISTANCE5                                             = 0x3005
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	MAX_PROGRAM_NATIVE_TEMPORARIES_ARB                         = 0x88A7
	MAX_VERTEX_SHADER_LOCAL_CONSTANTS_EXT                      = 0x87C8
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX                     = 0x8408
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	DEPTH_TEXTURE_MODE_ARB                                     = 0x884B
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	READ_FRAMEBUFFER_NV                                        = 0x8CA8
	HINT_BIT                                                   = 0x00008000
	CCW                                                        = 0x0901
	TEXTURE_BLUE_SIZE                                          = 0x805E
	PIXEL_TILE_HEIGHT_SGIX                                     = 0x8141
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	DYNAMIC_COPY                                               = 0x88EA
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	RENDERBUFFER_BLUE_SIZE_EXT                                 = 0x8D52
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	POINT_SIZE_MAX_ARB                                         = 0x8127
	TEXTURE_CLIPMAP_OFFSET_SGIX                                = 0x8173
	FRAGMENTS_INSTRUMENT_COUNTERS_SGIX                         = 0x8314
	MAX_SHININESS_NV                                           = 0x8504
	MIN_PROGRAM_TEXEL_OFFSET_NV                                = 0x8904
	DOUBLE_MAT3x2_EXT                                          = 0x8F4B
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	VERTEX_ARRAY_SIZE                                          = 0x807A
	REPLACE_EXT                                                = 0x8062
	INDEX_ARRAY_COUNT_EXT                                      = 0x8087
	DEPTH_BUFFER_FLOAT_MODE_NV                                 = 0x8DAF
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	FOG_BIT                                                    = 0x00000080
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	INCR                                                       = 0x1E02
	HISTOGRAM_LUMINANCE_SIZE                                   = 0x802C
	FLOAT_MAT3_ARB                                             = 0x8B5B
	STENCIL_INDEX4_EXT                                         = 0x8D47
	TESS_GEN_MODE                                              = 0x8E76
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS                             = 0x8354
	MATRIX3_NV                                                 = 0x8633
	BOOL                                                       = 0x8B56
	CURRENT_NORMAL                                             = 0x0B02
	DOT_PRODUCT_NV                                             = 0x86EC
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	ALLOW_DRAW_OBJ_HINT_PGI                                    = 0x1A20E
	POST_CONVOLUTION_ALPHA_BIAS_EXT                            = 0x8023
	STENCIL_OP_VALUE_AMD                                       = 0x874C
	UNSIGNED_INT_IMAGE_2D_RECT_EXT                             = 0x9065
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	PROXY_TEXTURE_3D                                           = 0x8070
	RGBA_FLOAT16_APPLE                                         = 0x881A
	MATRIX_INDEX_ARRAY_POINTER_OES                             = 0x8849
	DECODE_EXT                                                 = 0x8A49
	STORAGE_SHARED_APPLE                                       = 0x85BF
	OUTPUT_TEXTURE_COORD1_EXT                                  = 0x879E
	OUTPUT_TEXTURE_COORD18_EXT                                 = 0x87AF
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT_OES                      = 0x8CD6
	TEXCOORD2_BIT_PGI                                          = 0x20000000
	DEPTH_TEST                                                 = 0x0B71
	ZOOM_X                                                     = 0x0D16
	T2F_N3F_V3F                                                = 0x2A2B
	LUMINANCE8_EXT                                             = 0x8040
	TEXTURE_MAX_LEVEL                                          = 0x813D
	CULL_VERTEX_EYE_POSITION_EXT                               = 0x81AB
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	COMPRESSED_RGBA_S3TC_DXT5_ANGLE                            = 0x83F3
	PRIMARY_COLOR                                              = 0x8577
	MAP1_VERTEX_ATTRIB5_4_NV                                   = 0x8665
	BUMP_ROT_MATRIX_SIZE_ATI                                   = 0x8776
	DOUBLE_VEC2_EXT                                            = 0x8FFC
	LINES_ADJACENCY_ARB                                        = 0x000A
	PIXEL_MIN_FILTER_EXT                                       = 0x8332
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	VERTEX_STREAM5_ATI                                         = 0x8771
	SHADER_OBJECT_EXT                                          = 0x8B48
	TRANSFORM_FEEDBACK_VARYINGS_EXT                            = 0x8C83
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	FUNC_SUBTRACT                                              = 0x800A
	COLOR_INDEX8_EXT                                           = 0x80E5
	BUFFER_SIZE_ARB                                            = 0x8764
	MAX_VERTEX_BINDABLE_UNIFORMS_EXT                           = 0x8DE2
	SIGNALED                                                   = 0x9119
	SGIX_ycrcb                                                 = 1
	VERTEX_ARRAY_TYPE                                          = 0x807B
	INDEX_TEST_EXT                                             = 0x81B5
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	CONVOLUTION_WIDTH                                          = 0x8018
	MAX_CONVOLUTION_HEIGHT                                     = 0x801B
	SIGNED_HILO_NV                                             = 0x86F9
	MAGNITUDE_BIAS_NV                                          = 0x8718
	ALPHA_FLOAT16_APPLE                                        = 0x881C
	QUERY_COUNTER_BITS                                         = 0x8864
	CON_18_ATI                                                 = 0x8953
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	ALWAYS_SOFT_HINT_PGI                                       = 0x1A20D
	LIGHT5                                                     = 0x4005
	PROXY_HISTOGRAM                                            = 0x8025
	DETAIL_TEXTURE_MODE_SGIS                                   = 0x809B
	VERTEX_SHADER_INVARIANTS_EXT                               = 0x87D1
	CURRENT_PALETTE_MATRIX_OES                                 = 0x8843
	SAMPLER_1D_ARB                                             = 0x8B5D
	ATC_RGB_AMD                                                = 0x8C92
	RENDERBUFFER_ALPHA_SIZE_OES                                = 0x8D53
	RGB_SCALE_ARB                                              = 0x8573
	SOURCE0_ALPHA_ARB                                          = 0x8588
	BOOL_VEC4_ARB                                              = 0x8B59
	GEOMETRY_OUTPUT_TYPE_EXT                                   = 0x8DDC
	MAX_PROGRAM_MATRIX_STACK_DEPTH_ARB                         = 0x862E
	MAX_VERTEX_SHADER_INVARIANTS_EXT                           = 0x87C7
	FIELD_LOWER_NV                                             = 0x9023
	MAX_VERTEX_UNITS_ARB                                       = 0x86A4
	COLOR_ATTACHMENT15                                         = 0x8CEF
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	STENCIL_BACK_REF                                           = 0x8CA3
	DEBUG_ASSERT_MESA                                          = 0x875B
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT_OES              = 0x8CD7
	EXT_blend_minmax                                           = 1
	VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE                          = 0x88FE
	TEXTURE_EXTERNAL_OES                                       = 0x8D65
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	RESCALE_NORMAL                                             = 0x803A
	TEXTURE_ENV_BIAS_SGIX                                      = 0x80BE
	COLOR_TABLE_BIAS                                           = 0x80D7
	TEXTURE_COORD_ARRAY_PARALLEL_POINTERS_INTEL                = 0x83F8
	TRANSFORM_FEEDBACK_BUFFER_MODE_EXT                         = 0x8C7F
	INTERLEAVED_ATTRIBS_EXT                                    = 0x8C8C
	COMBINER6_NV                                               = 0x8556
	REPLACEMENT_CODE_ARRAY_POINTER_SUN                         = 0x85C3
	MAX_TEXTURE_COORDS_NV                                      = 0x8871
	SGIX_pixel_texture                                         = 1
	MAP2_NORMAL                                                = 0x0DB2
	FUNC_REVERSE_SUBTRACT_OES                                  = 0x800B
	COMBINER5_NV                                               = 0x8555
	PIXEL_SUBSAMPLE_4242_SGIX                                  = 0x85A4
	SWIZZLE_STR_ATI                                            = 0x8976
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	OFFSET_TEXTURE_SCALE_NV                                    = 0x86E2
	SET_AMD                                                    = 0x874A
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	COVERAGE_COMPONENT_NV                                      = 0x8ED0
	SRC_COLOR                                                  = 0x0300
	V3F                                                        = 0x2A21
	OCCLUSION_TEST_HP                                          = 0x8165
	MODELVIEW1_MATRIX_EXT                                      = 0x8506
	MAX_CUBE_MAP_TEXTURE_SIZE_ARB                              = 0x851C
	INT_VEC2                                                   = 0x8B53
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	OR_REVERSE                                                 = 0x150B
	R8I                                                        = 0x8231
	LUMINANCE_ALPHA_FLOAT16_ATI                                = 0x881F
	MATRIX8_ARB                                                = 0x88C8
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	SGIX_convolution_accuracy                                  = 1
	UNSIGNED_BYTE                                              = 0x1401
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	DEBUG_CATEGORY_DEPRECATION_AMD                             = 0x914B
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	BLUE_BIAS                                                  = 0x0D1B
	BGRA_EXT                                                   = 0x80E1
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX                             = 0x8178
	PIXEL_MAG_FILTER_EXT                                       = 0x8331
	TEXTURE_SAMPLES_IMG                                        = 0x9136
	DATA_BUFFER_AMD                                            = 0x9151
	GL_2PASS_0_EXT                                             = 0x80A2
	GLOBAL_ALPHA_SUN                                           = 0x81D9
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	SAMPLER_1D_SHADOW_ARB                                      = 0x8B61
	RENDERBUFFER_BINDING_ANGLE                                 = 0x8CA7
	LUMINANCE_INTEGER_EXT                                      = 0x8D9C
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	SIGNED_IDENTITY_NV                                         = 0x853C
	COMBINER_SUM_OUTPUT_NV                                     = 0x854C
	BUFFER_USAGE                                               = 0x8765
	MAX_VERTEX_ATTRIBS_ARB                                     = 0x8869
	TEXTURE_DEPTH_QCOM                                         = 0x8BD4
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	PRIMITIVES_GENERATED                                       = 0x8C87
	PRIMITIVE_RESTART_INDEX_NV                                 = 0x8559
	CON_3_ATI                                                  = 0x8944
	ALREADY_SIGNALED_APPLE                                     = 0x911A
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	BLEND                                                      = 0x0BE2
	CLIP_DISTANCE6                                             = 0x3006
	OP_INDEX_EXT                                               = 0x8782
	FLOAT_MAT4x2                                               = 0x8B69
	TEXTURE_COMPARE_FUNC_ARB                                   = 0x884D
	PIXEL_UNPACK_BUFFER_EXT                                    = 0x88EC
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	EXT_blend_subtract                                         = 1
	PACK_SWAP_BYTES                                            = 0x0D00
	TEXTURE_CUBE_MAP_POSITIVE_X_OES                            = 0x8515
	WEIGHT_ARRAY_TYPE_OES                                      = 0x86A9
	INT_SAMPLER_2D_ARRAY_EXT                                   = 0x8DCF
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	DOUBLE_MAT3x4_EXT                                          = 0x8F4C
	COMBINE_RGB                                                = 0x8571
	MAP2_VERTEX_ATTRIB12_4_NV                                  = 0x867C
	VERTEX_SHADER_INSTRUCTIONS_EXT                             = 0x87CF
	STENCIL_INDEX16                                            = 0x8D49
	DEPTH_ATTACHMENT_OES                                       = 0x8D00
	RG16_SNORM                                                 = 0x8F99
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	REFLECTION_MAP_EXT                                         = 0x8512
	ADD_SIGNED_ARB                                             = 0x8574
	MAX_RATIONAL_EVAL_ORDER_NV                                 = 0x86D7
	LUMINANCE16UI_EXT                                          = 0x8D7A
	SKIP_COMPONENTS3_NV                                        = -4
	NUM_FILL_STREAMS_NV                                        = 0x8E29
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	SRC_ALPHA                                                  = 0x0302
	LIGHT4                                                     = 0x4004
	OCCLUSION_TEST_RESULT_HP                                   = 0x8166
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX                          = 0x840A
	WEIGHT_ARRAY_POINTER_ARB                                   = 0x86AC
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	NO_ERROR                                                   = 0
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	YCBCR_422_APPLE                                            = 0x85B9
	UNSIGNED_SHORT_8_8_MESA                                    = 0x85BA
	PROGRAM_LENGTH_ARB                                         = 0x8627
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	INT_IMAGE_2D_MULTISAMPLE_EXT                               = 0x9060
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	LOCATION                                                   = 0x930E
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	POST_COLOR_MATRIX_BLUE_SCALE_SGI                           = 0x80B6
	PROGRAM_ATTRIBS_ARB                                        = 0x88AC
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	IMAGE_2D_ARRAY_EXT                                         = 0x9053
	UNPACK_LSB_FIRST                                           = 0x0CF1
	OUTPUT_TEXTURE_COORD12_EXT                                 = 0x87A9
	RASTERIZER_DISCARD                                         = 0x8C89
	TEXTURE_SWIZZLE_G_EXT                                      = 0x8E43
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	DIFFUSE                                                    = 0x1201
	CURRENT_FOG_COORDINATE                                     = 0x8453
	WEIGHT_ARRAY_STRIDE_OES                                    = 0x86AA
	MAX_OPTIMIZED_VERTEX_SHADER_INSTRUCTIONS_EXT               = 0x87CA
	STENCIL_BACK_PASS_DEPTH_PASS_ATI                           = 0x8803
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	TEXTURE_MATERIAL_PARAMETER_EXT                             = 0x8352
	CURRENT_PALETTE_MATRIX_ARB                                 = 0x8843
	READ_PIXEL_DATA_RANGE_POINTER_NV                           = 0x887D
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	READ_FRAMEBUFFER_EXT                                       = 0x8CA8
	MAX_BINDABLE_UNIFORM_SIZE_EXT                              = 0x8DED
	UNIFORM_BUFFER_BINDING_EXT                                 = 0x8DEF
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	STENCIL_RENDERABLE                                         = 0x8288
	COLOR_ATTACHMENT12_NV                                      = 0x8CEC
	POINT_DISTANCE_ATTENUATION_ARB                             = 0x8129
	TEXTURE_CLIPMAP_FRAME_SGIX                                 = 0x8172
	OPERAND0_ALPHA_EXT                                         = 0x8598
	SLUMINANCE8_ALPHA8_NV                                      = 0x8C45
	ALPHA                                                      = 0x1906
	C4F_N3F_V3F                                                = 0x2A26
	COLOR_ENCODING                                             = 0x8296
	R16I                                                       = 0x8233
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	PROXY_TEXTURE_CUBE_MAP_ARB                                 = 0x851B
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	DECAL                                                      = 0x2101
	PREVIOUS                                                   = 0x8578
	VERTEX_ATTRIB_ARRAY2_NV                                    = 0x8652
	R1UI_T2F_C4F_N3F_V3F_SUN                                   = 0x85CB
	VARIANT_ARRAY_POINTER_EXT                                  = 0x87E9
	GEOMETRY_INPUT_TYPE_EXT                                    = 0x8DDB
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	COMPUTE_SUBROUTINE                                         = 0x92ED
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	PROGRAM_STRING_NV                                          = 0x8628
	EVAL_VERTEX_ATTRIB3_NV                                     = 0x86C9
	VBO_FREE_MEMORY_ATI                                        = 0x87FB
	FLOAT_RGBA32_NV                                            = 0x888B
	READ_FRAMEBUFFER                                           = 0x8CA8
	TEXTURE_INDEX_SIZE_EXT                                     = 0x80ED
	SAMPLER_BUFFER_EXT                                         = 0x8DC2
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	SGIS_point_parameters                                      = 1
	MIN_EXT                                                    = 0x8007
	LOCAL_EXT                                                  = 0x87C4
	DRAW_BUFFER5_ARB                                           = 0x882A
	SPOT_DIRECTION                                             = 0x1204
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	NORMAL_MAP_OES                                             = 0x8511
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	COMBINE_RGB_ARB                                            = 0x8571
	SOURCE3_RGB_NV                                             = 0x8583
	DRAW_BUFFER2                                               = 0x8827
	RGB16UI_EXT                                                = 0x8D77
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	RASTER_POSITION_UNCLIPPED_IBM                              = 0x19262
	LINE_BIT                                                   = 0x00000004
	EYE_POINT_SGIS                                             = 0x81F4
	DOT2_ADD_ATI                                               = 0x896C
	RGBA16I                                                    = 0x8D88
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	CURRENT_MATRIX_STACK_DEPTH_NV                              = 0x8640
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	SIGNED_RGB8_NV                                             = 0x86FF
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	SGIX_resample                                              = 1
	DEPTH_FUNC                                                 = 0x0B74
	RGBA16                                                     = 0x805B
	TEXTURE_SHADOW                                             = 0x82A1
	BUMP_TARGET_ATI                                            = 0x877C
	CURRENT_QUERY_EXT                                          = 0x8865
	FRAGMENT_COLOR_EXT                                         = 0x834C
	RGB32F_ARB                                                 = 0x8815
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	DUAL_LUMINANCE4_SGIS                                       = 0x8114
	VERTEX_ATTRIB_ARRAY3_NV                                    = 0x8653
	SYNC_FLAGS                                                 = 0x9115
	DEBUG_SEVERITY_LOW                                         = 0x9148
	VIRTUAL_PAGE_SIZE_X_AMD                                    = 0x9195
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	INT16_VEC2_NV                                              = 0x8FE5
	POLYGON_OFFSET_BIAS_EXT                                    = 0x8039
	POST_COLOR_MATRIX_COLOR_TABLE                              = 0x80D2
	POST_TEXTURE_FILTER_BIAS_SGIX                              = 0x8179
	READ_PIXELS_FORMAT                                         = 0x828D
	CON_20_ATI                                                 = 0x8955
	PIXEL_TILE_CACHE_INCREMENT_SGIX                            = 0x813F
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	COMPRESSED_TEXTURE_FORMATS_ARB                             = 0x86A3
	TRIANGLE_MESH_SUN                                          = 0x8615
	MAX_DRAW_BUFFERS_NV                                        = 0x8824
	DOT_PRODUCT_AFFINE_DEPTH_REPLACE_NV                        = 0x885D
	IMAGE_BINDING_FORMAT                                       = 0x906E
	RECLAIM_MEMORY_HINT_PGI                                    = 0x1A1FE
	TEXTURE_BINDING_1D                                         = 0x8068
	RGBA12                                                     = 0x805A
	TEXTURE_DEPTH                                              = 0x8071
	REFERENCE_PLANE_EQUATION_SGIX                              = 0x817E
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	LIGHT_MODEL_COLOR_CONTROL_EXT                              = 0x81F8
	RGB32UI_EXT                                                = 0x8D71
	IR_INSTRUMENT1_SGIX                                        = 0x817F
	CONTEXT_FLAGS                                              = 0x821E
	MAX_CONVOLUTION_WIDTH                                      = 0x801A
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA                        = 0x8BB2
	ATC_RGBA_EXPLICIT_ALPHA_AMD                                = 0x8C93
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	VIEW_CLASS_32_BITS                                         = 0x82C8
	MODELVIEW_PROJECTION_NV                                    = 0x8629
	SIGNED_LUMINANCE_NV                                        = 0x8701
	DEPTH_STENCIL_TO_RGBA_NV                                   = 0x886E
	EXT_histogram                                              = 1
	UNSIGNED_SHORT_1_5_5_5_REV_EXT                             = 0x8366
	FOG                                                        = 0x0B60
	ALPHA_TEST                                                 = 0x0BC0
	FRAGMENT_LIGHT7_SGIX                                       = 0x8413
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	TRACE_MASK_MESA                                            = 0x8755
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	EVAL_VERTEX_ATTRIB8_NV                                     = 0x86CE
	COLOR_ATTACHMENT15_NV                                      = 0x8CEF
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	COPY                                                       = 0x1503
	DOT_PRODUCT_TEXTURE_3D_NV                                  = 0x86EF
	FLOAT_MAT4                                                 = 0x8B5C
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	MINMAX_EXT                                                 = 0x802E
	EDGE_FLAG_ARRAY_POINTER_EXT                                = 0x8093
	COMPRESSED_LUMINANCE_ALPHA_ARB                             = 0x84EB
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE                          = 0x85B0
	TEXTURE_DS_SIZE_NV                                         = 0x871D
	VERTEX_ATTRIB_MAP1_APPLE                                   = 0x8A00
	COLOR_ATTACHMENT6_EXT                                      = 0x8CE6
	SKIP_COMPONENTS4_NV                                        = -3
	UNSIGNED_INT_SAMPLER_RENDERBUFFER_NV                       = 0x8E58
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	DRAW_BUFFER5                                               = 0x882A
	REG_22_ATI                                                 = 0x8937
	CON_28_ATI                                                 = 0x895D
	INT_10_10_10_2_OES                                         = 0x8DF7
	DEBUG_CATEGORY_UNDEFINED_BEHAVIOR_AMD                      = 0x914C
	ZERO                                                       = 0
	SPECULAR                                                   = 0x1202
	ATC_RGBA_INTERPOLATED_ALPHA_AMD                            = 0x87EE
	R16_SNORM                                                  = 0x8F98
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	CONVOLUTION_FILTER_BIAS                                    = 0x8015
	PROGRAM_NATIVE_TEX_INDIRECTIONS_ARB                        = 0x880A
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	RENDERBUFFER_SAMPLES_EXT                                   = 0x8CAB
	COVERAGE_ATTACHMENT_NV                                     = 0x8ED2
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	POINT_FADE_THRESHOLD_SIZE_SGIS                             = 0x8128
	TRACK_MATRIX_NV                                            = 0x8648
	OBJECT_DISTANCE_TO_LINE_SGIS                               = 0x81F3
	TEXTURE_CUBE_MAP_POSITIVE_Y_EXT                            = 0x8517
	COLOR_ATTACHMENT4_EXT                                      = 0x8CE4
	RGBA16_SNORM                                               = 0x8F9B
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	SGI_texture_color_table                                    = 1
	MAP_WRITE_BIT                                              = 0x0002
	R8                                                         = 0x8229
	TEXTURE_CUBE_MAP_NEGATIVE_X_OES                            = 0x8516
	SOURCE0_RGB                                                = 0x8580
	TIME_ELAPSED_EXT                                           = 0x88BF
	RENDERBUFFER_BLUE_SIZE_OES                                 = 0x8D52
	RGBA_INTEGER                                               = 0x8D99
	BLUE_SCALE                                                 = 0x0D1A
	POINT_SIZE_MIN                                             = 0x8126
	TEXTURE16                                                  = 0x84D0
	WAIT_FAILED                                                = 0x911D
	VERTEX_CONSISTENT_HINT_PGI                                 = 0x1A22B
	INVALID_FRAMEBUFFER_OPERATION_OES                          = 0x0506
	COLOR_MATRIX_SGI                                           = 0x80B1
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	VIBRANCE_SCALE_NV                                          = 0x8713
	NEGATIVE_Y_EXT                                             = 0x87DA
	PN_TRIANGLES_TESSELATION_LEVEL_ATI                         = 0x87F4
	ALPHA_FLOAT16_ATI                                          = 0x881C
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	LUMINANCE16_SNORM                                          = 0x9019
	COLOR3_BIT_PGI                                             = 0x00010000
	ALPHA_TEST_FUNC_QCOM                                       = 0x0BC1
	PHONG_HINT_WIN                                             = 0x80EB
	ELEMENT_ARRAY_TYPE_ATI                                     = 0x8769
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	ROUND_NV                                                   = 0x90A4
	RELATIVE_LINE_TO_NV                                        = 0x05
	POST_CONVOLUTION_RED_SCALE                                 = 0x801C
	MAX_ASYNC_DRAW_PIXELS_SGIX                                 = 0x8360
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	VIDEO_COLOR_CONVERSION_MATRIX_NV                           = 0x9029
	TEXTURE_GEN_T                                              = 0x0C61
	BLEND_EQUATION_EXT                                         = 0x8009
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX                          = 0x8402
	COMBINER3_NV                                               = 0x8553
	RGB9_E5                                                    = 0x8C3D
	EXT_point_parameters                                       = 1
	COORD_REPLACE_NV                                           = 0x8862
	MATRIX13_ARB                                               = 0x88CD
	POLYGON_BIT                                                = 0x00000008
	PROGRAM_BINDING_ARB                                        = 0x8677
	VERTEX_SHADER_EXT                                          = 0x8780
	FACTOR_MIN_AMD                                             = 0x901C
	TEXTURE12_ARB                                              = 0x84CC
	MAX_PROGRAM_ATTRIBS_ARB                                    = 0x88AD
	CON_31_ATI                                                 = 0x8960
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	TEXTURE_GEN_Q                                              = 0x0C63
	PROGRAM_OBJECT_ARB                                         = 0x8B40
	SAMPLER_1D_SHADOW                                          = 0x8B61
	SYNC_STATUS_APPLE                                          = 0x9114
	CULL_FACE                                                  = 0x0B44
	COLOR_MATERIAL_FACE                                        = 0x0B55
	RGB10_A2                                                   = 0x8059
	MULTISAMPLE_SGIS                                           = 0x809D
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	TEXTURE_OBJECT_VALID_QCOM                                  = 0x8BDB
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	TIMEOUT_IGNORED                                            = 0xFFFFFFFF
	EDGE_FLAG_ARRAY_LIST_IBM                                   = 103075
	EXT_abgr                                                   = 1
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	FOG_COLOR                                                  = 0x0B66
	CONVOLUTION_FORMAT                                         = 0x8017
	POST_CONVOLUTION_BLUE_BIAS                                 = 0x8022
	LINEAR_SHARPEN_ALPHA_SGIS                                  = 0x80AE
	INT_SAMPLER_1D_ARRAY_EXT                                   = 0x8DCE
	SPARE0_PLUS_SECONDARY_COLOR_NV                             = 0x8532
	IMAGE_2D_ARRAY                                             = 0x9053
	POINT_SIZE_MIN_SGIS                                        = 0x8126
	MAP2_VERTEX_ATTRIB5_4_NV                                   = 0x8675
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI                           = 0x80BB
	TRANSLATE_2D_NV                                            = 0x9090
	VERTEX_ARRAY_POINTER                                       = 0x808E
	COLOR_TABLE_INTENSITY_SIZE                                 = 0x80DF
	DEBUG_TYPE_OTHER                                           = 0x8251
	UNSIGNED_SHORT_8_8_REV_MESA                                = 0x85BB
	DSDT8_NV                                                   = 0x8709
	OBJECT_ACTIVE_UNIFORM_MAX_LENGTH_ARB                       = 0x8B87
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	V2F                                                        = 0x2A20
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	SUBTRACT                                                   = 0x84E7
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	MATRIX21_ARB                                               = 0x88D5
	MUL_ATI                                                    = 0x8964
	COUNTER_TYPE_AMD                                           = 0x8BC0
	FRAMEBUFFER_BLEND                                          = 0x828B
	RENDERBUFFER_WIDTH_OES                                     = 0x8D42
	INDEX_LOGIC_OP                                             = 0x0BF1
	SINGLE_COLOR_EXT                                           = 0x81F9
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	FRAGMENT_COLOR_MATERIAL_SGIX                               = 0x8401
	TEXTURE5_ARB                                               = 0x84C5
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	POLYGON_STIPPLE                                            = 0x0B42
	FRAGMENTS_INSTRUMENT_SGIX                                  = 0x8313
	EVAL_2D_NV                                                 = 0x86C0
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	VERTEX_PROGRAM_POINT_SIZE_ARB                              = 0x8642
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	VIEW_CLASS_24_BITS                                         = 0x82C9
	EXP2                                                       = 0x0801
	UNSIGNED_INT_10_10_10_2_EXT                                = 0x8036
	MULTISAMPLE_EXT                                            = 0x809D
	TEXTURE11_ARB                                              = 0x84CB
	OUTPUT_TEXTURE_COORD4_EXT                                  = 0x87A1
	DRAW_PIXELS_APPLE                                          = 0x8A0A
	OPERAND2_RGB_ARB                                           = 0x8592
	LOCAL_CONSTANT_EXT                                         = 0x87C3
	READ_FRAMEBUFFER_ANGLE                                     = 0x8CA8
	COLOR_ATTACHMENT12                                         = 0x8CEC
	INTENSITY8_SNORM                                           = 0x9017
	AUX2                                                       = 0x040B
	COLOR_ARRAY_COUNT_EXT                                      = 0x8084
	CURRENT_FOG_COORD                                          = 0x8453
	TEXTURE2_ARB                                               = 0x84C2
	ARRAY_BUFFER_BINDING_ARB                                   = 0x8894
	TRANSFORM_FEEDBACK_BUFFER_START_NV                         = 0x8C84
	DEBUG_SEVERITY_HIGH_AMD                                    = 0x9146
	INDEX_OFFSET                                               = 0x0D13
	LUMINANCE8_ALPHA8                                          = 0x8045
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	REPLACE_VALUE_AMD                                          = 0x874B
	TESS_CONTROL_PROGRAM_NV                                    = 0x891E
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	DRAW_BUFFER4                                               = 0x8829
	DEPTH32F_STENCIL8                                          = 0x8CAD
	FLOAT16_VEC3_NV                                            = 0x8FFA
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	CLAMP_TO_EDGE                                              = 0x812F
	BUFFER_SERIALIZED_MODIFY_APPLE                             = 0x8A12
	SAMPLER_BUFFER                                             = 0x8DC2
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	INFO_LOG_LENGTH                                            = 0x8B84
	PALETTE8_RGBA8_OES                                         = 0x8B96
	ALLOW_DRAW_WIN_HINT_PGI                                    = 0x1A20F
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	SHARED_TEXTURE_PALETTE_EXT                                 = 0x81FB
	VERTEX_ATTRIB_ARRAY1_NV                                    = 0x8651
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS_EXT          = 0x8C8A
	COPY_READ_BUFFER                                           = 0x8F36
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	FOG_COORDINATE_ARRAY_STRIDE_EXT                            = 0x8455
	OUTPUT_TEXTURE_COORD7_EXT                                  = 0x87A4
	SHADER_TYPE                                                = 0x8B4F
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	COLOR_ATTACHMENT10_EXT                                     = 0x8CEA
	RENDERBUFFER_STENCIL_SIZE_EXT                              = 0x8D55
	SAMPLE_MASK                                                = 0x8E51
	STATIC_READ                                                = 0x88E5
	REG_17_ATI                                                 = 0x8932
	IMAGE_CUBE                                                 = 0x9050
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	RGB5_A1                                                    = 0x8057
	UNPACK_ROW_BYTES_APPLE                                     = 0x8A16
	FACTOR_ALPHA_MODULATE_IMG                                  = 0x8C07
	VERSION_2_0                                                = 1
	SAMPLE_ALPHA_TO_ONE_ARB                                    = 0x809F
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	WRITE_PIXEL_DATA_RANGE_POINTER_NV                          = 0x887C
	IMAGE_1D                                                   = 0x904C
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	SYNC_FLAGS_APPLE                                           = 0x9115
	TEXTURE_GEQUAL_R_SGIX                                      = 0x819D
	PIXEL_TEXTURE_SGIS                                         = 0x8353
	TEXTURE31_ARB                                              = 0x84DF
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	TANGENT_ARRAY_STRIDE_EXT                                   = 0x843F
	SOURCE2_RGB                                                = 0x8582
	PROGRAM_ERROR_POSITION_ARB                                 = 0x864B
	DRAW_BUFFER7_NV                                            = 0x882C
	DYNAMIC_DRAW_ARB                                           = 0x88E8
	VERTEX_ATTRIB_MAP1_COEFF_APPLE                             = 0x8A03
	RGBA8I                                                     = 0x8D8E
	Z4Y12Z4CB12Z4Y12Z4CR12_422_NV                              = 0x9035
	PIXEL_TILE_GRID_DEPTH_SGIX                                 = 0x8144
	LUMINANCE_ALPHA16F_ARB                                     = 0x881F
	RGBA8UI_EXT                                                = 0x8D7C
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	GL_4PASS_1_EXT                                             = 0x80A5
	TEXTURE20                                                  = 0x84D4
	COLOR_INDEXES                                              = 0x1603
	PACK_RESAMPLE_SGIX                                         = 0x842C
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	ATTRIB_ARRAY_TYPE_NV                                       = 0x8625
	TEXTURE_SHARED_SIZE_EXT                                    = 0x8C3F
	FRAGMENT_PROGRAM_PARAMETER_BUFFER_NV                       = 0x8DA4
	GPU_ADDRESS_NV                                             = 0x8F34
	SCALED_RESOLVE_FASTEST_EXT                                 = 0x90BA
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	SIGNED_LUMINANCE8_NV                                       = 0x8702
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	INT8_VEC2_NV                                               = 0x8FE1
	BLEND_SRC                                                  = 0x0BE1
	TEXTURE_COLOR_WRITEMASK_SGIS                               = 0x81EF
	PROGRAM_SEPARABLE                                          = 0x8258
	TEXTURE3                                                   = 0x84C3
	MATRIX27_ARB                                               = 0x88DB
	MAD_ATI                                                    = 0x8968
	RETAINED_APPLE                                             = 0x8A1B
	UNSIGNED_INT_8_8_8_8_REV_EXT                               = 0x8367
	WEIGHT_ARRAY_OES                                           = 0x86AD
	VERTEX_SHADER_BINDING_EXT                                  = 0x8781
	COLOR_FLOAT_APPLE                                          = 0x8A0F
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	DOUBLEBUFFER                                               = 0x0C32
	DEPTH_COMPONENT16_SGIX                                     = 0x81A5
	MAX_VERTEX_VARYING_COMPONENTS_EXT                          = 0x8DDE
	VIDEO_BUFFER_NV                                            = 0x9020
	MULTIVIEW_EXT                                              = 0x90F1
	DEPENDENT_HILO_TEXTURE_2D_NV                               = 0x8858
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	SGIS_generate_mipmap                                       = 1
	MIRRORED_REPEAT                                            = 0x8370
	FRAGMENT_LIGHT6_SGIX                                       = 0x8412
	MAX_CUBE_MAP_TEXTURE_SIZE_EXT                              = 0x851C
	OP_SET_GE_EXT                                              = 0x878C
	BGRA_INTEGER                                               = 0x8D9B
	INT_SAMPLER_3D_EXT                                         = 0x8DCB
	IMAGE_CUBE_MAP_ARRAY_EXT                                   = 0x9054
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	PIXEL_CUBIC_WEIGHT_EXT                                     = 0x8333
	BINORMAL_ARRAY_EXT                                         = 0x843A
	SOURCE1_RGB_ARB                                            = 0x8581
	COMBINE_ALPHA_EXT                                          = 0x8572
	FLOAT_CLEAR_COLOR_VALUE_NV                                 = 0x888D
	POINT_BIT                                                  = 0x00000002
	COLOR_ARRAY                                                = 0x8076
	RGBA4_OES                                                  = 0x8056
	COLOR_ARRAY_SIZE_EXT                                       = 0x8081
	LUMINANCE8_ALPHA8_SNORM                                    = 0x9016
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	BLEND_COLOR_EXT                                            = 0x8005
	RGBA2_EXT                                                  = 0x8055
	TEXTURE_INTENSITY_SIZE_EXT                                 = 0x8061
	COMBINER_MUX_SUM_NV                                        = 0x8547
	FRAGMENT_SHADER_DERIVATIVE_HINT_OES                        = 0x8B8B
	SEPARATE_ATTRIBS_EXT                                       = 0x8C8D
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	DEBUG_CATEGORY_APPLICATION_AMD                             = 0x914F
	EXPAND_NEGATE_NV                                           = 0x8539
	MAX_PROGRAM_MATRICES_ARB                                   = 0x862F
	VERTEX_ATTRIB_MAP1_ORDER_APPLE                             = 0x8A04
	PREFER_DOUBLEBUFFER_HINT_PGI                               = 0x1A1F8
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	GL_2PASS_1_EXT                                             = 0x80A3
	PACK_ROW_BYTES_APPLE                                       = 0x8A15
	ADD_BLEND_IMG                                              = 0x8C09
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	DOUBLE_MAT2                                                = 0x8F46
	MAX_EVAL_ORDER                                             = 0x0D30
	DEPENDENT_AR_TEXTURE_2D_NV                                 = 0x86E9
	TEXTURE_WRAP_T                                             = 0x2803
	RGBA12_EXT                                                 = 0x805A
	VERTEX_ARRAY_RANGE_POINTER_APPLE                           = 0x8521
	CURRENT_MATRIX_ARB                                         = 0x8641
	DRAW_BUFFER9_ARB                                           = 0x882E
	DT_BIAS_NV                                                 = 0x8717
	OP_ADD_EXT                                                 = 0x8787
	RED_BIT_ATI                                                = 0x00000001
	RENDERBUFFER_HEIGHT_EXT                                    = 0x8D43
	DEPTH_BITS                                                 = 0x0D56
	MAX_PROGRAM_NATIVE_TEX_INSTRUCTIONS_ARB                    = 0x880F
	COLOR_TABLE_WIDTH_SGI                                      = 0x80D9
	RGBA4_S3TC                                                 = 0x83A3
	PACK_SUBSAMPLE_RATE_SGIX                                   = 0x85A0
	EVAL_VERTEX_ATTRIB7_NV                                     = 0x86CD
	MAX_PN_TRIANGLES_TESSELATION_LEVEL_ATI                     = 0x87F1
	FLOAT_VEC4                                                 = 0x8B52
	PIXEL_BUFFER_BARRIER_BIT_EXT                               = 0x00000080
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	UNSIGNED_IDENTITY_NV                                       = 0x8536
	TRANSFORM_FEEDBACK_BUFFER_SIZE_NV                          = 0x8C85
	FRAMEBUFFER_ATTACHMENT_ANGLE                               = 0x93A3
	DOUBLE                                                     = 0x140A
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	MAX_RECTANGLE_TEXTURE_SIZE_NV                              = 0x84F8
	COMBINER1_NV                                               = 0x8551
	OUTPUT_TEXTURE_COORD10_EXT                                 = 0x87A7
	RENDERBUFFER_BINDING                                       = 0x8CA7
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	MAP1_VERTEX_ATTRIB13_4_NV                                  = 0x866D
	HILO16_NV                                                  = 0x86F8
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	COVERAGE_COMPONENT4_NV                                     = 0x8ED1
	ALPHA_BIAS                                                 = 0x0D1D
	POST_CONVOLUTION_BLUE_SCALE_EXT                            = 0x801E
	UNPACK_SKIP_IMAGES_EXT                                     = 0x806D
	MATRIX2_NV                                                 = 0x8632
	DSDT_MAG_INTENSITY_NV                                      = 0x86DC
	COLOR_ATTACHMENT4                                          = 0x8CE4
	MAT_DIFFUSE_BIT_PGI                                        = 0x00400000
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	LUMINANCE4_ALPHA4                                          = 0x8043
	NORMAL_ARRAY_STRIDE_EXT                                    = 0x807F
	GENERATE_MIPMAP_SGIS                                       = 0x8191
	DT_SCALE_NV                                                = 0x8711
	MAX_PROGRAM_INSTRUCTIONS_ARB                               = 0x88A1
	FORMAT_SUBSAMPLE_24_24_OML                                 = 0x8982
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	TESSELLATION_MODE_AMD                                      = 0x9004
	RGB565_OES                                                 = 0x8D62
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	MAX_CLIP_DISTANCES                                         = 0x0D32
	GL_4PASS_2_EXT                                             = 0x80A6
	SIGNED_ALPHA8_NV                                           = 0x8706
	DRAW_BUFFER2_ATI                                           = 0x8827
	CLAMP_READ_COLOR                                           = 0x891C
	RENDERBUFFER_GREEN_SIZE_OES                                = 0x8D51
	TEXTURE_RECTANGLE                                          = 0x84F5
	OUTPUT_TEXTURE_COORD28_EXT                                 = 0x87B9
	RGB16I                                                     = 0x8D89
	BACK_NORMALS_HINT_PGI                                      = 0x1A223
	BLEND_DST_ALPHA_OES                                        = 0x80CA
	TEXTURE_MAX_CLAMP_S_SGIX                                   = 0x8369
	PROXY_TEXTURE_RECTANGLE_ARB                                = 0x84F7
	TEXTURE_2D_STACK_BINDING_MESAX                             = 0x875E
	SGIX_fragment_lighting                                     = 1
	SAMPLE_PATTERN_EXT                                         = 0x80AC
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	MAGNITUDE_SCALE_NV                                         = 0x8712
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	INT_SAMPLER_BUFFER_AMD                                     = 0x9002
	OR_INVERTED                                                = 0x150D
	OFFSET_TEXTURE_2D_BIAS_NV                                  = 0x86E3
	COLOR_ATTACHMENT5_NV                                       = 0x8CE5
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	PATH_COORD_COUNT_NV                                        = 0x909E
	DEBUG_CATEGORY_SHADER_COMPILER_AMD                         = 0x914E
	AMBIENT                                                    = 0x1200
	UNPACK_SKIP_IMAGES                                         = 0x806D
	COMPUTE_TEXTURE                                            = 0x82A0
	TRANSFORM_FEEDBACK_BINDING_NV                              = 0x8E25
	MOVE_TO_RESETS_NV                                          = 0x90B5
	HISTOGRAM_SINK                                             = 0x802D
	MAX_RECTANGLE_TEXTURE_SIZE_ARB                             = 0x84F8
	VERTEX_WEIGHT_ARRAY_POINTER_EXT                            = 0x8510
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	MAX_GENERAL_COMBINERS_NV                                   = 0x854D
	DOT3_RGBA_IMG                                              = 0x86AF
	DOUBLE_MAT3_EXT                                            = 0x8F47
	CONVOLUTION_BORDER_MODE                                    = 0x8013
	MODELVIEW28_ARB                                            = 0x873C
	OP_LOG_BASE_2_EXT                                          = 0x8792
	WEIGHT_ARRAY_BUFFER_BINDING_ARB                            = 0x889E
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	MODELVIEW0_MATRIX_EXT                                      = 0x0BA6
	R32UI                                                      = 0x8236
	TEXTURE_CUBE_MAP_ARB                                       = 0x8513
	SRC0_ALPHA                                                 = 0x8588
	CON_29_ATI                                                 = 0x895E
	COLOR_ATTACHMENT8_NV                                       = 0x8CE8
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	OBJECT_TYPE                                                = 0x9112
	ABGR_EXT                                                   = 0x8000
	MAP_TESSELLATION_NV                                        = 0x86C2
	REG_7_ATI                                                  = 0x8928
	UNIFORM_SIZE                                               = 0x8A38
	FRAMEBUFFER_BINDING_EXT                                    = 0x8CA6
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	TEXTURE_APPLICATION_MODE_EXT                               = 0x834F
	ALPHA_FLOAT32_ATI                                          = 0x8816
	DRAW_BUFFER7_ARB                                           = 0x882C
	VERTEX_ARRAY_LIST_IBM                                      = 103070
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	ALPHA_TEST_REF_QCOM                                        = 0x0BC2
	CURRENT_QUERY_ARB                                          = 0x8865
	VERTEX_PROGRAM_POSITION_MESA                               = 0x8BB4
	PATCH_VERTICES                                             = 0x8E72
	DOUBLE_MAT2_EXT                                            = 0x8F46
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	ZOOM_Y                                                     = 0x0D17
	TEXTURE_LUMINANCE_SIZE_EXT                                 = 0x8060
	MAX_VERTEX_SHADER_VARIANTS_EXT                             = 0x87C6
	DRAW_BUFFER1                                               = 0x8826
	FORCE_BLUE_TO_ONE_NV                                       = 0x8860
	ALPHA8_SNORM                                               = 0x9014
	LARGE_CW_ARC_TO_NV                                         = 0x18
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	RG32I                                                      = 0x823B
	TRANSPOSE_MODELVIEW_MATRIX_ARB                             = 0x84E3
	MAP2_VERTEX_ATTRIB8_4_NV                                   = 0x8678
	MATRIX6_ARB                                                = 0x88C6
	MAX_LIGHTS                                                 = 0x0D31
	TEXTURE_COORD_ARRAY_EXT                                    = 0x8078
	COLOR_MATRIX_STACK_DEPTH                                   = 0x80B2
	MAX_PROGRAM_LOOP_COUNT_NV                                  = 0x88F8
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	MINMAX_SINK                                                = 0x8030
	HALF_BIAS_NORMAL_NV                                        = 0x853A
	CON_14_ATI                                                 = 0x894F
	RGB32I_EXT                                                 = 0x8D83
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	MAP1_VERTEX_ATTRIB9_4_NV                                   = 0x8669
	RED_INTEGER                                                = 0x8D94
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	COLOR_INDEX                                                = 0x1900
	TEXTURE_LOD_BIAS_R_SGIX                                    = 0x8190
	GPU_OPTIMIZED_QCOM                                         = 0x8FB2
	MATRIX_EXT                                                 = 0x87C0
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER_EXT                     = 0x8CDC
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE_NV                        = 0x8E24
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	CURRENT_COLOR                                              = 0x0B00
	INTENSITY12_EXT                                            = 0x804C
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	INT8_NV                                                    = 0x8FE0
	PARTIAL_SUCCESS_NV                                         = 0x902E
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	SUBPIXEL_BITS                                              = 0x0D50
	OUTPUT_TEXTURE_COORD21_EXT                                 = 0x87B2
	GL_3DC_X_AMD                                               = 0x87F9
	PROGRAM_FORMAT_ARB                                         = 0x8876
	TEXTURE_BLUE_TYPE_ARB                                      = 0x8C12
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	PATH_STENCIL_REF_NV                                        = 0x90B8
	CLIP_DISTANCE2                                             = 0x3002
	POST_COLOR_MATRIX_RED_SCALE                                = 0x80B4
	COLOR_SUM_EXT                                              = 0x8458
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	TRANSFORM_FEEDBACK_RECORD_NV                               = 0x8C86
	FRAMEZOOM_SGIX                                             = 0x818B
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	MAP_READ_BIT                                               = 0x0001
	PACK_SKIP_ROWS                                             = 0x0D03
	COMPILE_AND_EXECUTE                                        = 0x1301
	SAMPLE_ALPHA_TO_ONE_EXT                                    = 0x809F
	OP_POWER_EXT                                               = 0x8793
	CON_6_ATI                                                  = 0x8947
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	RED_BIAS                                                   = 0x0D15
	TEXTURE_BINDING_2D                                         = 0x8069
	CONSTANT_COLOR_EXT                                         = 0x8001
	MATRIX5_ARB                                                = 0x88C5
	TEXTURE_BUFFER                                             = 0x8C2A
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	COLOR_ATTACHMENT10                                         = 0x8CEA
	HISTOGRAM_BLUE_SIZE                                        = 0x802A
	POST_COLOR_MATRIX_GREEN_BIAS                               = 0x80B9
	YCRCBA_SGIX                                                = 0x8319
	MAP1_VERTEX_ATTRIB3_4_NV                                   = 0x8663
	DRAW_BUFFER14                                              = 0x8833
	COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_EXT                = 0x8C73
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	DEPTH_STENCIL_EXT                                          = 0x84F9
	RED_MIN_CLAMP_INGR                                         = 0x8560
	R1UI_V3F_SUN                                               = 0x85C4
	MAX_FRAGMENT_PROGRAM_LOCAL_PARAMETERS_NV                   = 0x8868
	RASTERIZER_DISCARD_NV                                      = 0x8C89
	RENDERBUFFER_SAMPLES_NV                                    = 0x8CAB
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	ADD                                                        = 0x0104
	C4UB_V3F                                                   = 0x2A23
	MAX                                                        = 0x8008
	BUFFER_OBJECT_APPLE                                        = 0x85B3
	CON_10_ATI                                                 = 0x894B
	FRAGMENT_ALPHA_MODULATE_IMG                                = 0x8C08
	LUMINANCE_ALPHA16UI_EXT                                    = 0x8D7B
	FRONT_LEFT                                                 = 0x0400
	NEXT_BUFFER_NV                                             = -2
	SAMPLE_MASK_VALUE_NV                                       = 0x8E52
	BEVEL_NV                                                   = 0x90A6
	VARIANT_VALUE_EXT                                          = 0x87E4
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	MATERIAL_SIDE_HINT_PGI                                     = 0x1A22C
	COMPRESSED_INTENSITY_ARB                                   = 0x84EC
	CULL_FRAGMENT_NV                                           = 0x86E7
	EIGHTH_BIT_ATI                                             = 0x00000020
	POINT_SIZE_ARRAY_STRIDE_OES                                = 0x898B
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	STENCIL_WRITEMASK                                          = 0x0B98
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	VERTEX_ATTRIB_MAP2_COEFF_APPLE                             = 0x8A07
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	SGIS_sharpen_texture                                       = 1
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	STATIC_DRAW                                                = 0x88E4
	FRAMEBUFFER_INCOMPLETE_FORMATS_EXT                         = 0x8CDA
	STENCIL_INDEX1                                             = 0x8D46
	LUMINANCE_ALPHA32I_EXT                                     = 0x8D87
	SAMPLER_2D_ARRAY_SHADOW_EXT                                = 0x8DC4
	INT_IMAGE_CUBE_MAP_ARRAY_EXT                               = 0x905F
	PATH_GEN_MODE_NV                                           = 0x90B0
	SAMPLER_OBJECT_AMD                                         = 0x9155
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	BINORMAL_ARRAY_TYPE_EXT                                    = 0x8440
	DRAW_BUFFER11                                              = 0x8830
	FLOAT_VEC4_ARB                                             = 0x8B52
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	Z4Y12Z4CB12Z4CR12_444_NV                                   = 0x9037
	TEXTURE_GEN_MODE                                           = 0x2500
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	BOOL_VEC4                                                  = 0x8B59
	COMPRESSED_SRGB                                            = 0x8C48
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	COLOR_TABLE_SGI                                            = 0x80D0
	TEXTURE28                                                  = 0x84DC
	E_TIMES_F_NV                                               = 0x8531
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	COLOR_ARRAY_STRIDE                                         = 0x8083
	FOG_FUNC_SGIS                                              = 0x812A
	FRAGMENT_TEXTURE                                           = 0x829F
	INTENSITY_FLOAT32_APPLE                                    = 0x8817
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	RENDERBUFFER_ALPHA_SIZE_EXT                                = 0x8D53
	SAMPLER_1D_ARRAY_SHADOW_EXT                                = 0x8DC3
	FLOAT16_VEC4_NV                                            = 0x8FFB
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	MAP1_VERTEX_ATTRIB12_4_NV                                  = 0x866C
	FLOAT_R32_NV                                               = 0x8885
	ALPHA32UI_EXT                                              = 0x8D72
	INT_IMAGE_1D_EXT                                           = 0x9057
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	MAP2_VERTEX_ATTRIB13_4_NV                                  = 0x867D
	INSTRUMENT_BUFFER_POINTER_SGIX                             = 0x8180
	IUI_N3F_V3F_EXT                                            = 0x81B0
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	TEXTURE_RENDERBUFFER_DATA_STORE_BINDING_NV                 = 0x8E54
	TEXTURE_CLIPMAP_CENTER_SGIX                                = 0x8171
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	BOUNDING_BOX_NV                                            = 0x908D
	PROXY_TEXTURE_1D                                           = 0x8063
	DEPENDENT_GB_TEXTURE_2D_NV                                 = 0x86EA
	SGIX_texture_add_env                                       = 1
	GL_2D                                                      = 0x0600
	DEPTH_RENDERABLE                                           = 0x8287
	SAMPLER_3D                                                 = 0x8B5F
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	YCRCB_SGIX                                                 = 0x8318
	DS_BIAS_NV                                                 = 0x8716
	HISTOGRAM_ALPHA_SIZE_EXT                                   = 0x802B
	T2F_IUI_V2F_EXT                                            = 0x81B1
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	R1UI_C4UB_V3F_SUN                                          = 0x85C5
	SAMPLER_RENDERBUFFER_NV                                    = 0x8E56
	IMAGE_BINDING_LAYERED_EXT                                  = 0x8F3C
	COLOR_MATERIAL                                             = 0x0B57
	FRAGMENT_LIGHT1_SGIX                                       = 0x840D
	GREEN_MIN_CLAMP_INGR                                       = 0x8561
	INTERPOLATE_EXT                                            = 0x8575
	REPLACEMENT_CODE_ARRAY_TYPE_SUN                            = 0x85C1
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	MATRIX_INDEX_ARRAY_ARB                                     = 0x8844
	FLOAT_RGB16_NV                                             = 0x8888
	FLOAT_MAT3x2                                               = 0x8B67
	COMPRESSED_RGBA_PVRTC_4BPPV1_IMG                           = 0x8C02
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	INT_IMAGE_BUFFER                                           = 0x905C
	MAX_LAYERS                                                 = 0x8281
	INTENSITY_FLOAT32_ATI                                      = 0x8817
	TEXTURE_LUMINANCE_TYPE                                     = 0x8C14
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	STREAM_READ                                                = 0x88E1
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE_OES                     = 0x8CD0
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE                        = 0x80D5
	IMAGE_MIN_FILTER_HP                                        = 0x815D
	MULTISAMPLE_FILTER_HINT_NV                                 = 0x8534
	R1UI_C4F_N3F_V3F_SUN                                       = 0x85C8
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	CONVOLUTION_FORMAT_EXT                                     = 0x8017
	TABLE_TOO_LARGE_EXT                                        = 0x8031
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	TEXTURE13                                                  = 0x84CD
	PRIMITIVE_RESTART                                          = 0x8F9D
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	BLEND_DST_RGB                                              = 0x80C8
	VERTEX_STREAM1_ATI                                         = 0x876D
	MAX_PATCH_VERTICES                                         = 0x8E7D
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	VERTEX_SHADER_ARB                                          = 0x8B31
	SHADER_COMPILER                                            = 0x8DFA
	TRANSFORM_FEEDBACK                                         = 0x8E22
	INT_IMAGE_CUBE_EXT                                         = 0x905B
	QUADS                                                      = 0x0007
	SAMPLES_ARB                                                = 0x80A9
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	CON_21_ATI                                                 = 0x8956
	SGIX_texture_multi_buffer                                  = 1
	TEXTURE_3D_BINDING_EXT                                     = 0x806A
	VERTEX_ARRAY_TYPE_EXT                                      = 0x807B
	GEOMETRY_DEFORMATION_SGIX                                  = 0x8194
	MAP1_BINORMAL_EXT                                          = 0x8446
	VERTEX_ATTRIB_ARRAY9_NV                                    = 0x8659
	MAX_VERTEX_SHADER_LOCALS_EXT                               = 0x87C9
	DRAW_BUFFER14_ATI                                          = 0x8833
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS_OES                      = 0x8CD9
	MAX_SAMPLES_NV                                             = 0x8D57
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	CPU_OPTIMIZED_QCOM                                         = 0x8FB1
	POLYGON_OFFSET_FILL                                        = 0x8037
	REPLACEMENT_CODE_SUN                                       = 0x81D8
	VERSION_1_2                                                = 1
	REFERENCE_PLANE_SGIX                                       = 0x817D
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	UNPACK_CLIENT_STORAGE_APPLE                                = 0x85B2
	INVERSE_NV                                                 = 0x862B
	DRAW_BUFFER6_NV                                            = 0x882B
	VALIDATE_STATUS                                            = 0x8B83
	INT_IMAGE_1D_ARRAY_EXT                                     = 0x905D
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	LAYOUT_DEFAULT_INTEL                                       = 0
	BLEND_EQUATION_OES                                         = 0x8009
	CONVOLUTION_BORDER_COLOR                                   = 0x8154
	TEXTURE22                                                  = 0x84D6
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	HILO8_NV                                                   = 0x885E
	CURRENT_OCCLUSION_QUERY_ID_NV                              = 0x8865
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	MAP2_COLOR_4                                               = 0x0DB0
	IMAGE_MAG_FILTER_HP                                        = 0x815C
	RENDERBUFFER_FREE_MEMORY_ATI                               = 0x87FD
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	MAX_GEOMETRY_OUTPUT_VERTICES_EXT                           = 0x8DE0
	CURRENT_RASTER_COLOR                                       = 0x0B04
	COLOR_ARRAY_TYPE                                           = 0x8082
	GL_4_BYTES                                                 = 0x1409
	PARALLEL_ARRAYS_INTEL                                      = 0x83F4
	RGBA_FLOAT32_ATI                                           = 0x8814
	DOUBLE_MAT3                                                = 0x8F47
	ALPHA4_EXT                                                 = 0x803B
	CURRENT_PROGRAM                                            = 0x8B8D
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	MAX_COLOR_MATRIX_STACK_DEPTH                               = 0x80B3
	DEPTH_STENCIL_TO_BGRA_NV                                   = 0x886F
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	DEPTH_BUFFER_BIT                                           = 0x00000100
	FOG_MODE                                                   = 0x0B65
	HISTOGRAM_GREEN_SIZE_EXT                                   = 0x8029
	TEXTURE_DEPTH_EXT                                          = 0x8071
	TESS_CONTROL_TEXTURE                                       = 0x829C
	SCALEBIAS_HINT_SGIX                                        = 0x8322
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX             = 0x840B
	NUM_LOOPBACK_COMPONENTS_ATI                                = 0x8974
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	DEPTH_COMPONENT32                                          = 0x81A7
	DOT3_RGB_ARB                                               = 0x86AE
	VERTEX_SHADER_LOCAL_CONSTANTS_EXT                          = 0x87D2
	TESS_EVALUATION_PROGRAM_PARAMETER_BUFFER_NV                = 0x8C75
	COLOR_SUM                                                  = 0x8458
	TEXTURE19_ARB                                              = 0x84D3
	MAX_SAMPLES_EXT                                            = 0x8D57
	STENCIL_INDEX                                              = 0x1901
	TEXTURE_MATERIAL_FACE_EXT                                  = 0x8351
	TANGENT_ARRAY_POINTER_EXT                                  = 0x8442
	VARIABLE_C_NV                                              = 0x8525
	SAMPLER_2D                                                 = 0x8B5E
	UNSIGNED_SHORT_5_5_5_1_EXT                                 = 0x8034
	MAX_4D_TEXTURE_SIZE_SGIS                                   = 0x8138
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME_EXT                     = 0x8CD1
	MAX_ACTIVE_LIGHTS_SGIX                                     = 0x8405
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	SGIS_texture_select                                        = 1
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	HISTOGRAM_LUMINANCE_SIZE_EXT                               = 0x802C
	MODELVIEW13_ARB                                            = 0x872D
	TEXTURE_1D_STACK_BINDING_MESAX                             = 0x875D
	SAMPLER_BINDING                                            = 0x8919
	TEXTURE_BINDING_BUFFER_EXT                                 = 0x8C2C
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	INSTRUMENT_MEASUREMENTS_SGIX                               = 0x8181
	SAMPLER_2D_ARB                                             = 0x8B5E
	MAX_CLIPMAP_DEPTH_SGIX                                     = 0x8177
	REGISTER_COMBINERS_NV                                      = 0x8522
	EVAL_VERTEX_ATTRIB2_NV                                     = 0x86C8
	MAX_TEXTURE_COORDS                                         = 0x8871
	MAX_PROGRAM_NATIVE_ATTRIBS_ARB                             = 0x88AF
	DOUBLE_MAT3x4                                              = 0x8F4C
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	DRAW_BUFFER9_NV                                            = 0x882E
	TEXTURE_TYPE_QCOM                                          = 0x8BD7
	PROXY_POST_CONVOLUTION_COLOR_TABLE                         = 0x80D4
	MAX_DRAW_BUFFERS_ATI                                       = 0x8824
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	SGIX_sprite                                                = 1
	SAMPLE_ALPHA_TO_MASK_EXT                                   = 0x809E
	DUAL_ALPHA4_SGIS                                           = 0x8110
	STENCIL_INDEX16_EXT                                        = 0x8D49
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	SHADER                                                     = 0x82E1
	MATRIX17_ARB                                               = 0x88D1
	PALETTE4_RGB8_OES                                          = 0x8B90
	RESTART_PATH_NV                                            = 0xF0
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	SGIX_framezoom                                             = 1
	BLEND_EQUATION_RGB_OES                                     = 0x8009
	PROGRAM_NATIVE_PARAMETERS_ARB                              = 0x88AA
	TEXTURE_IMAGE_VALID_QCOM                                   = 0x8BD8
	UNSIGNED_INT_10_10_10_2_OES                                = 0x8DF6
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	COMPRESSED_RGBA_FXT1_3DFX                                  = 0x86B1
	OBJECT_ATTACHED_OBJECTS_ARB                                = 0x8B85
	CURRENT_INDEX                                              = 0x0B01
	CURRENT_ATTRIB_NV                                          = 0x8626
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL_OES                   = 0x8CD2
	TEXTURE28_ARB                                              = 0x84DC
	CURRENT_VERTEX_WEIGHT_EXT                                  = 0x850B
	R1UI_T2F_N3F_V3F_SUN                                       = 0x85CA
	SWIZZLE_STQ_DQ_ATI                                         = 0x8979
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	DEPTH_COMPONENT16_NONLINEAR_NV                             = 0x8E2C
	BINNING_CONTROL_HINT_QCOM                                  = 0x8FB0
	CLOSE_PATH_NV                                              = 0x00
	MAX_3D_TEXTURE_SIZE_OES                                    = 0x8073
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI                          = 0x80B7
	DUAL_LUMINANCE_ALPHA8_SGIS                                 = 0x811D
	TEXTURE15                                                  = 0x84CF
	LUMINANCE_FLOAT32_APPLE                                    = 0x8818
	INT                                                        = 0x1404
	SPRITE_EYE_ALIGNED_SGIX                                    = 0x814E
	EVAL_VERTEX_ATTRIB13_NV                                    = 0x86D3
	MATRIX_INDEX_ARRAY_TYPE_ARB                                = 0x8847
	STENCIL_CLEAR_TAG_VALUE_EXT                                = 0x88F3
	BLUE_BIT_ATI                                               = 0x00000004
	TESS_CONTROL_PROGRAM_PARAMETER_BUFFER_NV                   = 0x8C74
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	TEXTURE30                                                  = 0x84DE
	PACK_INVERT_MESA                                           = 0x8758
	COLOR_ATTACHMENT0_OES                                      = 0x8CE0
	RGB8UI                                                     = 0x8D7D
	FALSE                                                      = 0
	LINEAR_DETAIL_SGIS                                         = 0x8097
	DRAW_BUFFER4_ARB                                           = 0x8829
	GL_4X_BIT_ATI                                              = 0x00000002
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	SGIX_impact_pixel_texture                                  = 1
	POST_TEXTURE_FILTER_SCALE_SGIX                             = 0x817A
	MAP1_VERTEX_ATTRIB8_4_NV                                   = 0x8668
	STENCIL_BACK_PASS_DEPTH_FAIL_ATI                           = 0x8802
	COLOR_ATTACHMENT13_NV                                      = 0x8CED
	SGIX_depth_texture                                         = 1
	DEPTH_COMPONENT32_SGIX                                     = 0x81A7
	SEPARATE_SPECULAR_COLOR_EXT                                = 0x81FA
	TEXTURE11                                                  = 0x84CB
	MAX_VARYING_FLOATS_ARB                                     = 0x8B4B
	QUAD_LUMINANCE4_SGIS                                       = 0x8120
	RG32UI                                                     = 0x823C
	MAX_TRACK_MATRIX_STACK_DEPTH_NV                            = 0x862E
	VERTEX_ATTRIB_ARRAY_SIZE_ARB                               = 0x8623
	COLOR_ATTACHMENT14_NV                                      = 0x8CEE
	NATIVE_GRAPHICS_BEGIN_HINT_PGI                             = 0x1A203
	MAP1_VERTEX_3                                              = 0x0D97
	SRGB_READ                                                  = 0x8297
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	QUAD_TEXTURE_SELECT_SGIS                                   = 0x8125
	MAX_FOG_FUNC_POINTS_SGIS                                   = 0x812C
	PIXEL_PACK_BUFFER_BINDING_EXT                              = 0x88ED
	AFFINE_3D_NV                                               = 0x9094
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	TEXTURE24                                                  = 0x84D8
	VERTEX_PROGRAM_TWO_SIDE_NV                                 = 0x8643
	MAP1_VERTEX_ATTRIB0_4_NV                                   = 0x8660
	PROXY_TEXTURE_2D_STACK_MESAX                               = 0x875C
	DRAW_BUFFER12_ATI                                          = 0x8831
	NUM_VIDEO_CAPTURE_STREAMS_NV                               = 0x9024
	DEBUG_CATEGORY_WINDOW_SYSTEM_AMD                           = 0x914A
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	COUNT_DOWN_NV                                              = 0x9089
)

type Context struct {
	context                 *C.gl21Context
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
	AttachShader            func(program, shader uint32)
	BindAttribLocation      func(program, index uint32, name string)
	BlendEquationSeperate   func(modeRGB, modeAlpha uint32)
	CompileShader           func(shader uint32)
	CreateProgram           func() uint32
	CreateShader            func(shaderType uint32) uint32
}

func New() *Context {
	glc := new(Context)
	glc.context = C.gl21NewContext()

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
	if glc.context.fnAttachShader == nil {
		return nil
	}
	if glc.context.fnBindAttribLocation == nil {
		return nil
	}
	if glc.context.fnBlendEquationSeperate == nil {
		return nil
	}
	if glc.context.fnCompileShader == nil {
		return nil
	}
	if glc.context.fnCreateProgram == nil {
		return nil
	}
	if glc.context.fnCreateShader == nil {
		return nil
	}

	glc.Accum = func(op uint32, value float32) {
		C.gl21Accum(glc.context, C.GLenum(op), C.GLfloat(value))
	}

	glc.AlphaFunc = func(Func uint32, ref float32) {
		C.gl21AlphaFunc(glc.context, C.GLenum(Func), C.GLclampf(ref))
	}

	glc.Begin = func(mode uint32) {
		C.gl21Begin(glc.context, C.GLenum(mode))
	}

	glc.End = func() {
		C.gl21End(glc.context)
	}

	glc.Bitmap = func(width, height int32, xorig, yorig, xmove, ymove float32, bitmap []uint8) {
		C.gl21Bitmap(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(&bitmap[0])))
	}

	glc.BlendFunc = func(sfactor, dfactor uint32) {
		C.gl21BlendFunc(glc.context, C.GLenum(sfactor), C.GLenum(dfactor))
	}

	glc.CallList = func(list uint32) {
		C.gl21CallList(glc.context, C.GLuint(list))
	}

	glc.CallLists = func(n int32, Type uint32, lists unsafe.Pointer) {
		C.gl21CallLists(glc.context, C.GLsizei(n), C.GLenum(Type), lists)
	}

	glc.Clear = func(mask uint32) {
		C.gl21Clear(glc.context, C.GLbitfield(mask))
	}

	glc.ClearAccum = func(red, green, blue, alpha float32) {
		C.gl21ClearAccum(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.ClearColor = func(red, green, blue, alpha float32) {
		C.gl21ClearColor(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.ClearDepth = func(depth float64) {
		C.gl21ClearDepth(glc.context, C.GLclampd(depth))
	}

	glc.ClearIndex = func(c float32) {
		C.gl21ClearIndex(glc.context, C.GLfloat(c))
	}

	glc.ClearStencil = func(s int32) {
		C.gl21ClearStencil(glc.context, C.GLint(s))
	}

	glc.ClipPlane = func(plane uint32, equation []float64) {
		C.gl21ClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(&equation[0])))
	}

	glc.Color3b = func(red, green, blue int8) {
		C.gl21Color3b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
	}

	glc.Color3d = func(red, green, blue float64) {
		C.gl21Color3d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
	}

	glc.Color3f = func(red, green, blue float32) {
		C.gl21Color3f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
	}

	glc.Color3i = func(red, green, blue int32) {
		C.gl21Color3i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue))
	}

	glc.Color3s = func(red, green, blue int16) {
		C.gl21Color3s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue))
	}

	glc.Color3ub = func(red, green, blue uint8) {
		C.gl21Color3ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
	}

	glc.Color3ui = func(red, green, blue uint32) {
		C.gl21Color3ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue))
	}

	glc.Color3us = func(red, green, blue uint16) {
		C.gl21Color3us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue))
	}

	glc.Color4b = func(red, green, blue, alpha int8) {
		C.gl21Color4b(glc.context, C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
	}

	glc.Color4d = func(red, green, blue, alpha float64) {
		C.gl21Color4d(glc.context, C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
	}

	glc.Color4f = func(red, green, blue, alpha float32) {
		C.gl21Color4f(glc.context, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}

	glc.Color4i = func(red, green, blue, alpha int32) {
		C.gl21Color4i(glc.context, C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
	}

	glc.Color4s = func(red, green, blue, alpha int16) {
		C.gl21Color4s(glc.context, C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
	}

	glc.Color4ub = func(red, green, blue, alpha uint8) {
		C.gl21Color4ub(glc.context, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}

	glc.Color4ui = func(red, green, blue, alpha uint32) {
		C.gl21Color4ui(glc.context, C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
	}

	glc.Color4us = func(red, green, blue, alpha uint16) {
		C.gl21Color4us(glc.context, C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
	}

	glc.Color3bv = func(v []int8) {
		C.gl21Color3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color3dv = func(v []float64) {
		C.gl21Color3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.Color3fv = func(v []float32) {
		C.gl21Color3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.Color3iv = func(v []int32) {
		C.gl21Color3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.Color3sv = func(v []int16) {
		C.gl21Color3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.Color3ubv = func(v []uint8) {
		C.gl21Color3ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color3uiv = func(v []uint32) {
		C.gl21Color3uiv(glc.context, (*C.GLuint)(unsafe.Pointer(&v[0])))
	}

	glc.Color3usv = func(v []uint16) {
		C.gl21Color3usv(glc.context, (*C.GLushort)(unsafe.Pointer(&v[0])))
	}

	glc.Color4bv = func(v []int8) {
		C.gl21Color4bv(glc.context, (*C.GLbyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color4dv = func(v []float64) {
		C.gl21Color4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.Color4fv = func(v []float32) {
		C.gl21Color4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.Color4iv = func(v []int32) {
		C.gl21Color4iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.Color4sv = func(v []int16) {
		C.gl21Color4sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.Color4ubv = func(v []uint8) {
		C.gl21Color4ubv(glc.context, (*C.GLubyte)(unsafe.Pointer(&v[0])))
	}

	glc.Color4uiv = func(v []uint32) {
		C.gl21Color4uiv(glc.context, (*C.GLuint)(unsafe.Pointer(&v[0])))
	}

	glc.Color4usv = func(v []uint16) {
		C.gl21Color4usv(glc.context, (*C.GLushort)(unsafe.Pointer(&v[0])))
	}

	glc.ColorMask = func(red, green, blue, alpha bool) {
		C.gl21ColorMask(glc.context, boolToGL(red), boolToGL(green), boolToGL(blue), boolToGL(alpha))
	}

	glc.ColorMaterial = func(face, mode uint32) {
		C.gl21ColorMaterial(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.ColorTable = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl21ColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ColorTableParameterfv = func(target, pname uint32, params []float32) {
		C.gl21ColorTableParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.ColorTableParameteriv = func(target, pname uint32, params []int32) {
		C.gl21ColorTableParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.ColorSubTable = func(target uint32, start, count int32, format, Type uint32, data unsafe.Pointer) {
		C.gl21ColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLsizei(count), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.CopyPixels = func(x, y int32, width, height int32, Type uint32) {
		C.gl21CopyPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(Type))
	}

	glc.CullFace = func(mode uint32) {
		C.gl21CullFace(glc.context, C.GLenum(mode))
	}

	glc.ConvolutionFilter1D = func(target, internalformat uint32, width int32, format, Type uint32, data unsafe.Pointer) {
		C.gl21ConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionFilter2D = func(target, internalformat uint32, width, height int32, format, Type uint32, data unsafe.Pointer) {
		C.gl21ConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), data)
	}

	glc.ConvolutionParameterf = func(target, pname uint32, params float32) {
		C.gl21ConvolutionParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
	}

	glc.ConvolutionParameteri = func(target, pname uint32, params int32) {
		C.gl21ConvolutionParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(params))
	}

	glc.CopyColorTable = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl21CopyColorTable(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyColorSubTable = func(target uint32, start int32, x, y int32, width int32) {
		C.gl21CopyColorSubTable(glc.context, C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter1D = func(target, internalformat uint32, x, y int32, width int32) {
		C.gl21CopyConvolutionFilter1D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyConvolutionFilter2D = func(target, internalformat uint32, x, y int32, width, height int32) {
		C.gl21CopyConvolutionFilter2D(glc.context, C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.DeleteLists = func(list uint32, Range int32) {
		C.gl21DeleteLists(glc.context, C.GLuint(list), C.GLsizei(Range))
	}

	glc.DepthFunc = func(Func uint32) {
		C.gl21DepthFunc(glc.context, C.GLenum(Func))
	}

	glc.DepthRange = func(zNear, zFar float64) {
		C.gl21DepthRange(glc.context, C.GLclampd(zNear), C.GLclampd(zFar))
	}

	glc.Enable = func(cap uint32) {
		C.gl21Enable(glc.context, C.GLenum(cap))
	}

	glc.Disable = func(cap uint32) {
		C.gl21Disable(glc.context, C.GLenum(cap))
	}

	glc.DrawBuffer = func(mode uint32) {
		C.gl21DrawBuffer(glc.context, C.GLenum(mode))
	}

	glc.DrawPixels = func(width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21DrawPixels(glc.context, C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.EdgeFlag = func(flag bool) {
		C.gl21EdgeFlag(glc.context, boolToGL(flag))
	}

	glc.EdgeFlagv = func(flag []bool) {
		C.gl21EdgeFlagv(glc.context, (*C.GLboolean)(unsafe.Pointer(&flag[0])))
	}

	glc.EdgeFlagPointer = func(stride int32, pointer unsafe.Pointer) {
		C.gl21EdgeFlagPointer(glc.context, C.GLsizei(stride), pointer)
	}

	glc.EnableClientState = func(cap uint32) {
		C.gl21EnableClientState(glc.context, C.GLenum(cap))
	}

	glc.DisableClientState = func(cap uint32) {
		C.gl21DisableClientState(glc.context, C.GLenum(cap))
	}

	glc.EvalCoord1d = func(u float64) {
		C.gl21EvalCoord1d(glc.context, C.GLdouble(u))
	}

	glc.EvalCoord1f = func(u float32) {
		C.gl21EvalCoord1f(glc.context, C.GLfloat(u))
	}

	glc.EvalCoord2d = func(u, v float64) {
		C.gl21EvalCoord2d(glc.context, C.GLdouble(u), C.GLdouble(v))
	}

	glc.EvalCoord2f = func(u, v float32) {
		C.gl21EvalCoord2f(glc.context, C.GLfloat(u), C.GLfloat(v))
	}

	glc.EvalCoord1dv = func(u []float64) {
		C.gl21EvalCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&u[0])))
	}

	glc.EvalCoord1fv = func(u []float32) {
		C.gl21EvalCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&u[0])))
	}

	glc.EvalCoord2dv = func(u []float64) {
		C.gl21EvalCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&u[0])))
	}

	glc.EvalCoord2fv = func(u []float32) {
		C.gl21EvalCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&u[0])))
	}

	glc.EvalMesh1 = func(mode uint32, i1, i2 int32) {
		C.gl21EvalMesh1(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2))
	}

	glc.EvalMesh2 = func(mode uint32, i1, i2, j1, j2 int32) {
		C.gl21EvalMesh2(glc.context, C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
	}

	glc.EvalPoint1 = func(i int32) {
		C.gl21EvalPoint1(glc.context, C.GLint(i))
	}

	glc.EvalPoint2 = func(i, j int32) {
		C.gl21EvalPoint2(glc.context, C.GLint(i), C.GLint(j))
	}

	glc.FeedbackBuffer = func(size int32, Type uint32, buffer []float32) {
		C.gl21FeedbackBuffer(glc.context, C.GLsizei(size), C.GLenum(Type), (*C.GLfloat)(unsafe.Pointer(&buffer[0])))
	}

	glc.Finish = func() {
		C.gl21Finish(glc.context)
	}

	glc.Flush = func() {
		C.gl21Flush(glc.context)
	}

	glc.Fogf = func(pname uint32, param float32) {
		C.gl21Fogf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.Fogi = func(pname uint32, param int32) {
		C.gl21Fogi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.Fogfv = func(pname uint32, params []float32) {
		C.gl21Fogfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.Fogiv = func(pname uint32, params []int32) {
		C.gl21Fogiv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.FrontFace = func(mode uint32) {
		C.gl21FrontFace(glc.context, C.GLenum(mode))
	}

	glc.Frustum = func(left, right, bottom, top, zNear, zFar float64) {
		C.gl21Frustum(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
	}

	glc.GenLists = func(Range int32) uint32 {
		return uint32(C.gl21GenLists(glc.context, C.GLsizei(Range)))
	}

	glc.GetBooleanv = func(pname uint32, params []bool) {
		C.gl21GetBooleanv(glc.context, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(&params[0])))
	}

	glc.GetDoublev = func(pname uint32, params []float64) {
		C.gl21GetDoublev(glc.context, C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))
	}

	glc.GetFloatv = func(pname uint32, params []float32) {
		C.gl21GetFloatv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetIntegerv = func(pname uint32, params []int32) {
		C.gl21GetIntegerv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetClipPlane = func(plane uint32, equation []float64) {
		C.gl21GetClipPlane(glc.context, C.GLenum(plane), (*C.GLdouble)(unsafe.Pointer(&equation[0])))
	}

	glc.GetError = func() uint32 {
		return uint32(C.gl21GetError(glc.context))
	}

	glc.GetLightfv = func(light, pname uint32, params []float32) {
		C.gl21GetLightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetLightiv = func(light, pname uint32, params []int32) {
		C.gl21GetLightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetMapdv = func(target, query uint32, v []float64) {
		C.gl21GetMapdv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.GetMapfv = func(target, query uint32, v []float32) {
		C.gl21GetMapfv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.GetMapiv = func(target, query uint32, v []int32) {
		C.gl21GetMapiv(glc.context, C.GLenum(target), C.GLenum(query), (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.GetMaterialfv = func(face, pname uint32, params []float32) {
		C.gl21GetMaterialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetMaterialiv = func(face, pname uint32, params []int32) {
		C.gl21GetMaterialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetPixelMapfv = func(Map uint32, values []float32) {
		C.gl21GetPixelMapfv(glc.context, C.GLenum(Map), (*C.GLfloat)(unsafe.Pointer(&values[0])))
	}

	glc.GetPixelMapuiv = func(Map uint32, values []uint32) {
		C.gl21GetPixelMapuiv(glc.context, C.GLenum(Map), (*C.GLuint)(unsafe.Pointer(&values[0])))
	}

	glc.GetPixelMapusv = func(Map uint32, values []uint16) {
		C.gl21GetPixelMapusv(glc.context, C.GLenum(Map), (*C.GLushort)(unsafe.Pointer(&values[0])))
	}

	glc.GetPolygonStipple = func(mask []uint8) {
		C.gl21GetPolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(&mask[0])))
	}

	glc.GetString = func(name uint32) string {
		cstr := C.gl21GetString(glc.context, C.GLenum(name))
		return C.GoString((*C.char)(unsafe.Pointer(cstr)))
	}

	glc.GetTexEnvfv = func(target, pname uint32, params []float32) {
		C.gl21GetTexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexEnviv = func(target, pname uint32, params []int32) {
		C.gl21GetTexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexGendv = func(coord, pname uint32, params []float64) {
		C.gl21GetTexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexGenfv = func(coord, pname uint32, params []float32) {
		C.gl21GetTexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexGeniv = func(coord, pname uint32, params []int32) {
		C.gl21GetTexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexImage = func(target uint32, level int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21GetTexImage(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.GetTexLevelParameterfv = func(target uint32, level int32, pname uint32, params []float32) {
		C.gl21GetTexLevelParameterfv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexLevelParameteriv = func(target uint32, level int32, pname uint32, params []int32) {
		C.gl21GetTexLevelParameteriv(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexParameterfv = func(target, pname uint32, params []float32) {
		C.gl21GetTexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.GetTexParameteriv = func(target, pname uint32, params []int32) {
		C.gl21GetTexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.Hint = func(target, mode uint32) {
		C.gl21Hint(glc.context, C.GLenum(target), C.GLenum(mode))
	}

	glc.Indexd = func(c float64) {
		C.gl21Indexd(glc.context, C.GLdouble(c))
	}

	glc.Indexf = func(c float32) {
		C.gl21Indexf(glc.context, C.GLfloat(c))
	}

	glc.Indexi = func(c int32) {
		C.gl21Indexi(glc.context, C.GLint(c))
	}

	glc.Indexs = func(c int16) {
		C.gl21Indexs(glc.context, C.GLshort(c))
	}

	glc.Indexub = func(c uint8) {
		C.gl21Indexub(glc.context, C.GLubyte(c))
	}

	glc.Indexdv = func(c []float64) {
		C.gl21Indexdv(glc.context, (*C.GLdouble)(unsafe.Pointer(&c[0])))
	}

	glc.Indexfv = func(c []float32) {
		C.gl21Indexfv(glc.context, (*C.GLfloat)(unsafe.Pointer(&c[0])))
	}

	glc.Indexiv = func(c []int32) {
		C.gl21Indexiv(glc.context, (*C.GLint)(unsafe.Pointer(&c[0])))
	}

	glc.Indexsv = func(c []int16) {
		C.gl21Indexsv(glc.context, (*C.GLshort)(unsafe.Pointer(&c[0])))
	}

	glc.Indexubv = func(c []uint8) {
		C.gl21Indexubv(glc.context, (*C.GLubyte)(unsafe.Pointer(&c[0])))
	}

	glc.IndexMask = func(mask uint32) {
		C.gl21IndexMask(glc.context, C.GLuint(mask))
	}

	glc.IndexPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl21IndexPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.InitNames = func() {
		C.gl21InitNames(glc.context)
	}

	glc.InterleavedArrays = func(format uint32, stride int32, pointer unsafe.Pointer) {
		C.gl21InterleavedArrays(glc.context, C.GLenum(format), C.GLsizei(stride), pointer)
	}

	glc.IsEnabled = func(cap uint32) {
		C.gl21IsEnabled(glc.context, C.GLenum(cap))
	}

	glc.IsList = func(list uint32) bool {
		return C.gl21IsList(glc.context, C.GLuint(list)) != 0
	}

	glc.Lightf = func(light, pname uint32, param float32) {
		C.gl21Lightf(glc.context, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Lighti = func(light, pname uint32, param int32) {
		C.gl21Lighti(glc.context, C.GLenum(light), C.GLenum(pname), C.GLint(param))
	}

	glc.Lightfv = func(light, pname uint32, params []float32) {
		C.gl21Lightfv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.Lightiv = func(light, pname uint32, params []int32) {
		C.gl21Lightiv(glc.context, C.GLenum(light), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.LightModelf = func(pname uint32, param float32) {
		C.gl21LightModelf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.LightModeli = func(pname uint32, param int32) {
		C.gl21LightModeli(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.LightModelfv = func(pname uint32, params []float32) {
		C.gl21LightModelfv(glc.context, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.LightModeliv = func(pname uint32, params []int32) {
		C.gl21LightModeliv(glc.context, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.LineStipple = func(factor int32, pattern uint16) {
		C.gl21LineStipple(glc.context, C.GLint(factor), C.GLushort(pattern))
	}

	glc.LineWidth = func(width float32) {
		C.gl21LineWidth(glc.context, C.GLfloat(width))
	}

	glc.ListBase = func(base uint32) {
		C.gl21ListBase(glc.context, C.GLuint(base))
	}

	glc.LoadIdentity = func() {
		C.gl21LoadIdentity(glc.context)
	}

	glc.LoadMatrixd = func(m []float64) {
		C.gl21LoadMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(&m[0])))
	}

	glc.LoadMatrixf = func(m []float32) {
		C.gl21LoadMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(&m[0])))
	}

	glc.LoadName = func(name uint32) {
		C.gl21LoadName(glc.context, C.GLuint(name))
	}

	glc.LogicOp = func(opcode uint32) {
		C.gl21LogicOp(glc.context, C.GLenum(opcode))
	}

	glc.Map1d = func(target uint32, u1, u2 float64, stride, order int32, points []float64) {
		C.gl21Map1d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(unsafe.Pointer(&points[0])))
	}

	glc.Map1f = func(target uint32, u1, u2 float32, stride, order int32, points []float32) {
		C.gl21Map1f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(unsafe.Pointer(&points[0])))
	}

	glc.Map2d = func(target uint32, u1, u2 float64, ustride, uorder int32, v1, v2 float64, vstride, vorder int32, points []float64) {
		C.gl21Map2d(glc.context, C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(unsafe.Pointer(&points[0])))
	}

	glc.Map2f = func(target uint32, u1, u2 float32, ustride, uorder int32, v1, v2 float32, vstride, vorder int32, points []float32) {
		C.gl21Map2f(glc.context, C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(unsafe.Pointer(&points[0])))
	}

	glc.MapGrid1d = func(un int32, u1, u2 float64) {
		C.gl21MapGrid1d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
	}

	glc.MapGrid1f = func(un int32, u1, u2 float32) {
		C.gl21MapGrid1f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
	}

	glc.MapGrid2d = func(un int32, u1, u2 float64, vn int32, v1, v2 float64) {
		C.gl21MapGrid2d(glc.context, C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
	}

	glc.MapGrid2f = func(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
		C.gl21MapGrid2f(glc.context, C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
	}

	glc.Materialf = func(face, pname uint32, param float32) {
		C.gl21Materialf(glc.context, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}

	glc.Materiali = func(face, pname uint32, param int32) {
		C.gl21Materiali(glc.context, C.GLenum(face), C.GLenum(pname), C.GLint(param))
	}

	glc.Materialfv = func(face, pname uint32, params []float32) {
		C.gl21Materialfv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.Materialiv = func(face, pname uint32, params []int32) {
		C.gl21Materialiv(glc.context, C.GLenum(face), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.MatrixMode = func(mode uint32) {
		C.gl21MatrixMode(glc.context, C.GLenum(mode))
	}

	glc.MultMatrixd = func(m []float64) {
		C.gl21MultMatrixd(glc.context, (*C.GLdouble)(unsafe.Pointer(&m[0])))
	}

	glc.MultMatrixf = func(m []float32) {
		C.gl21MultMatrixf(glc.context, (*C.GLfloat)(unsafe.Pointer(&m[0])))
	}

	glc.NewList = func(list uint32, mode uint32) {
		C.gl21NewList(glc.context, C.GLuint(list), C.GLenum(mode))
	}

	glc.EndList = func() {
		C.gl21EndList(glc.context)
	}

	glc.Normal3b = func(nx, ny, nz int8) {
		C.gl21Normal3b(glc.context, C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
	}

	glc.Normal3d = func(nx, ny, nz float64) {
		C.gl21Normal3d(glc.context, C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
	}

	glc.Normal3f = func(nx, ny, nz float32) {
		C.gl21Normal3f(glc.context, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}

	glc.Normal3i = func(nx, ny, nz int32) {
		C.gl21Normal3i(glc.context, C.GLint(nx), C.GLint(ny), C.GLint(nz))
	}

	glc.Normal3s = func(nx, ny, nz int16) {
		C.gl21Normal3s(glc.context, C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
	}

	glc.Normal3bv = func(v []int8) {
		C.gl21Normal3bv(glc.context, (*C.GLbyte)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3dv = func(v []float64) {
		C.gl21Normal3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3fv = func(v []float32) {
		C.gl21Normal3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3iv = func(v []int32) {
		C.gl21Normal3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.Normal3sv = func(v []int16) {
		C.gl21Normal3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.NormalPointer = func(Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl21NormalPointer(glc.context, C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.Ortho = func(left, right, bottom, top, zNear, zfar float64) {
		C.gl21Ortho(glc.context, C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zfar))
	}

	glc.PassThrough = func(token float32) {
		C.gl21PassThrough(glc.context, C.GLfloat(token))
	}

	glc.PixelMapfv = func(Map uint32, mapsize int32, values []float32) {
		C.gl21PixelMapfv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLfloat)(unsafe.Pointer(&values[0])))
	}

	glc.PixelMapuiv = func(Map uint32, mapsize int32, values []uint32) {
		C.gl21PixelMapuiv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLuint)(unsafe.Pointer(&values[0])))
	}

	glc.PixelMapusv = func(Map uint32, mapsize int32, values []uint16) {
		C.gl21PixelMapusv(glc.context, C.GLenum(Map), C.GLsizei(mapsize), (*C.GLushort)(unsafe.Pointer(&values[0])))
	}

	glc.PixelStoref = func(pname uint32, param float32) {
		C.gl21PixelStoref(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelStorei = func(pname uint32, param int32) {
		C.gl21PixelStorei(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelTransferf = func(pname uint32, param float32) {
		C.gl21PixelTransferf(glc.context, C.GLenum(pname), C.GLfloat(param))
	}

	glc.PixelTransferi = func(pname uint32, param int32) {
		C.gl21PixelTransferi(glc.context, C.GLenum(pname), C.GLint(param))
	}

	glc.PixelZoom = func(xfactor, yfactor float32) {
		C.gl21PixelZoom(glc.context, C.GLfloat(xfactor), C.GLfloat(yfactor))
	}

	glc.PointSize = func(size float32) {
		C.gl21PointSize(glc.context, C.GLfloat(size))
	}

	glc.PolygonMode = func(face, mode uint32) {
		C.gl21PolygonMode(glc.context, C.GLenum(face), C.GLenum(mode))
	}

	glc.PolygonStipple = func(mask []uint8) {
		C.gl21PolygonStipple(glc.context, (*C.GLubyte)(unsafe.Pointer(&mask[0])))
	}

	glc.PrioritizeTextures = func(n int32, textures []uint32, priorities []float32) {
		C.gl21PrioritizeTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(&textures[0])), (*C.GLclampf)(unsafe.Pointer(&priorities[0])))
	}

	glc.PushAttrib = func(mask uint32) {
		C.gl21PushAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopAttrib = func() {
		C.gl21PopAttrib(glc.context)
	}

	glc.PushClientAttrib = func(mask uint32) {
		C.gl21PushClientAttrib(glc.context, C.GLbitfield(mask))
	}

	glc.PopClientAttrib = func() {
		C.gl21PopClientAttrib(glc.context)
	}

	glc.PushMatrix = func() {
		C.gl21PushMatrix(glc.context)
	}

	glc.PopMatrix = func() {
		C.gl21PopMatrix(glc.context)
	}

	glc.PushName = func(name uint32) {
		C.gl21PushName(glc.context, C.GLuint(name))
	}

	glc.PopName = func() {
		C.gl21PopName(glc.context)
	}

	glc.RasterPos2d = func(x, y float64) {
		C.gl21RasterPos2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.RasterPos2f = func(x, y float32) {
		C.gl21RasterPos2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.RasterPos2i = func(x, y int32) {
		C.gl21RasterPos2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.RasterPos2s = func(x, y int16) {
		C.gl21RasterPos2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.RasterPos3d = func(x, y, z float64) {
		C.gl21RasterPos3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.RasterPos3f = func(x, y, z float32) {
		C.gl21RasterPos3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.RasterPos3i = func(x, y, z int32) {
		C.gl21RasterPos3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.RasterPos3s = func(x, y, z int16) {
		C.gl21RasterPos3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.RasterPos4d = func(x, y, z, w float64) {
		C.gl21RasterPos4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.RasterPos4f = func(x, y, z, w float32) {
		C.gl21RasterPos4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.RasterPos4i = func(x, y, z, w int32) {
		C.gl21RasterPos4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.RasterPos4s = func(x, y, z, w int16) {
		C.gl21RasterPos4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.RasterPos2dv = func(v []float64) {
		C.gl21RasterPos2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos2fv = func(v []float32) {
		C.gl21RasterPos2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos2iv = func(v []int32) {
		C.gl21RasterPos2iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos2sv = func(v []int16) {
		C.gl21RasterPos2sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3dv = func(v []float64) {
		C.gl21RasterPos3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3fv = func(v []float32) {
		C.gl21RasterPos3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3iv = func(v []int32) {
		C.gl21RasterPos3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos3sv = func(v []int16) {
		C.gl21RasterPos3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4dv = func(v []float64) {
		C.gl21RasterPos4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4fv = func(v []float32) {
		C.gl21RasterPos4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4iv = func(v []int32) {
		C.gl21RasterPos4iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.RasterPos4sv = func(v []int16) {
		C.gl21RasterPos4sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.ReadBuffer = func(mode uint32) {
		C.gl21ReadBuffer(glc.context, C.GLenum(mode))
	}

	glc.ReadPixels = func(x, y int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21ReadPixels(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Rectd = func(x1, y1, x2, y2 float64) {
		C.gl21Rectd(glc.context, C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
	}

	glc.Rectf = func(x1, y1, x2, y2 float32) {
		C.gl21Rectf(glc.context, C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
	}

	glc.Recti = func(x1, y1, x2, y2 int32) {
		C.gl21Recti(glc.context, C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
	}

	glc.Rects = func(x1, y1, x2, y2 int16) {
		C.gl21Rects(glc.context, C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
	}

	glc.Rectdv = func(v1, v2 []float64) {
		C.gl21Rectdv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v1[0])), (*C.GLdouble)(unsafe.Pointer(&v2[0])))
	}

	glc.Rectfv = func(v1, v2 []float32) {
		C.gl21Rectfv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v1[0])), (*C.GLfloat)(unsafe.Pointer(&v2[0])))
	}

	glc.Rectiv = func(v1, v2 []int32) {
		C.gl21Rectiv(glc.context, (*C.GLint)(unsafe.Pointer(&v1[0])), (*C.GLint)(unsafe.Pointer(&v2[0])))
	}

	glc.Rectsv = func(v1, v2 []int16) {
		C.gl21Rectsv(glc.context, (*C.GLshort)(unsafe.Pointer(&v1[0])), (*C.GLshort)(unsafe.Pointer(&v2[0])))
	}

	glc.RenderMode = func(mode uint32) int32 {
		return int32(C.gl21RenderMode(glc.context, C.GLenum(mode)))
	}

	glc.Rotated = func(angle, x, y, z float64) {
		C.gl21Rotated(glc.context, C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Rotatef = func(angle, x, y, z float32) {
		C.gl21Rotatef(glc.context, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scaled = func(x, y, z float64) {
		C.gl21Scaled(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Scalef = func(x, y, z float32) {
		C.gl21Scalef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Scissor = func(x, y int32, width, height int32) {
		C.gl21Scissor(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.SelectBuffer = func(size int32, buffer []uint32) {
		C.gl21SelectBuffer(glc.context, C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(&buffer[0])))
	}

	glc.ShadeModel = func(mode uint32) {
		C.gl21ShadeModel(glc.context, C.GLenum(mode))
	}

	glc.StencilFunc = func(Func uint32, ref int32, mask uint32) {
		C.gl21StencilFunc(glc.context, C.GLenum(Func), C.GLint(ref), C.GLuint(mask))
	}

	glc.StencilMask = func(mask uint32) {
		C.gl21StencilMask(glc.context, C.GLuint(mask))
	}

	glc.StencilOp = func(fail, zfail, zpass uint32) {
		C.gl21StencilOp(glc.context, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}

	glc.TexCoord1d = func(s float64) {
		C.gl21TexCoord1d(glc.context, C.GLdouble(s))
	}

	glc.TexCoord1f = func(s float32) {
		C.gl21TexCoord1f(glc.context, C.GLfloat(s))
	}

	glc.TexCoord1i = func(s int32) {
		C.gl21TexCoord1i(glc.context, C.GLint(s))
	}

	glc.TexCoord1s = func(s int16) {
		C.gl21TexCoord1s(glc.context, C.GLshort(s))
	}

	glc.TexCoord2d = func(s, t float64) {
		C.gl21TexCoord2d(glc.context, C.GLdouble(s), C.GLdouble(t))
	}

	glc.TexCoord2f = func(s, t float32) {
		C.gl21TexCoord2f(glc.context, C.GLfloat(s), C.GLfloat(t))
	}

	glc.TexCoord2i = func(s, t int32) {
		C.gl21TexCoord2i(glc.context, C.GLint(s), C.GLint(t))
	}

	glc.TexCoord2s = func(s, t int16) {
		C.gl21TexCoord2s(glc.context, C.GLshort(s), C.GLshort(t))
	}

	glc.TexCoord3d = func(s, t, r float64) {
		C.gl21TexCoord3d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
	}

	glc.TexCoord3f = func(s, t, r float32) {
		C.gl21TexCoord3f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
	}

	glc.TexCoord3i = func(s, t, r int32) {
		C.gl21TexCoord3i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r))
	}

	glc.TexCoord3s = func(s, t, r int16) {
		C.gl21TexCoord3s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r))
	}

	glc.TexCoord4d = func(s, t, r, q float64) {
		C.gl21TexCoord4d(glc.context, C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
	}

	glc.TexCoord4f = func(s, t, r, q float32) {
		C.gl21TexCoord4f(glc.context, C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}

	glc.TexCoord4i = func(s, t, r, q int32) {
		C.gl21TexCoord4i(glc.context, C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
	}

	glc.TexCoord4s = func(s, t, r, q int16) {
		C.gl21TexCoord4s(glc.context, C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
	}

	glc.TexCoord1dv = func(v []float64) {
		C.gl21TexCoord1dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord1fv = func(v []float32) {
		C.gl21TexCoord1fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord1iv = func(v []int32) {
		C.gl21TexCoord1iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord1sv = func(v []int16) {
		C.gl21TexCoord1sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2dv = func(v []float64) {
		C.gl21TexCoord2dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2fv = func(v []float32) {
		C.gl21TexCoord2fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2iv = func(v []int32) {
		C.gl21TexCoord2iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord2sv = func(v []int16) {
		C.gl21TexCoord2sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3dv = func(v []float64) {
		C.gl21TexCoord3dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3fv = func(v []float32) {
		C.gl21TexCoord3fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3iv = func(v []int32) {
		C.gl21TexCoord3iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord3sv = func(v []int16) {
		C.gl21TexCoord3sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4dv = func(v []float64) {
		C.gl21TexCoord4dv(glc.context, (*C.GLdouble)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4fv = func(v []float32) {
		C.gl21TexCoord4fv(glc.context, (*C.GLfloat)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4iv = func(v []int32) {
		C.gl21TexCoord4iv(glc.context, (*C.GLint)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoord4sv = func(v []int16) {
		C.gl21TexCoord4sv(glc.context, (*C.GLshort)(unsafe.Pointer(&v[0])))
	}

	glc.TexCoordPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl21TexCoordPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.TexEnvf = func(target, pname uint32, param float32) {
		C.gl21TexEnvf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexEnvi = func(target, pname uint32, param int32) {
		C.gl21TexEnvi(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexEnvfv = func(target, pname uint32, params []float32) {
		C.gl21TexEnvfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.TexEnviv = func(target, pname uint32, params []int32) {
		C.gl21TexEnviv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.TexGend = func(coord, pname uint32, param float64) {
		C.gl21TexGend(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
	}

	glc.TexGenf = func(coord, pname uint32, param float32) {
		C.gl21TexGenf(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexGeni = func(coord, pname uint32, param int32) {
		C.gl21TexGeni(glc.context, C.GLenum(coord), C.GLenum(pname), C.GLint(param))
	}

	glc.TexGendv = func(coord, pname uint32, params []float64) {
		C.gl21TexGendv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(unsafe.Pointer(&params[0])))
	}

	glc.TexGenfv = func(coord, pname uint32, params []float32) {
		C.gl21TexGenfv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.TexGeniv = func(coord, pname uint32, params []int32) {
		C.gl21TexGeniv(glc.context, C.GLenum(coord), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.TexImage1D = func(target uint32, level, internalformat int32, width int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21TexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage2D = func(target uint32, level, internalformat int32, width, height int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21TexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexImage3DEXT = func(target uint32, level, internalformat int32, width, height, depth int32, border int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21TexImage3DEXT(glc.context, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexParameterf = func(target, pname uint32, param float32) {
		C.gl21TexParameterf(glc.context, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}

	glc.TexParameteri = func(target, pname uint32, param int32) {
		C.gl21TexParameteri(glc.context, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}

	glc.TexParameterfv = func(target, pname uint32, params []float32) {
		C.gl21TexParameterfv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(&params[0])))
	}

	glc.TexParameteriv = func(target, pname uint32, params []int32) {
		C.gl21TexParameteriv(glc.context, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(&params[0])))
	}

	glc.TexSubImage1D = func(target uint32, level, xoffset int32, width int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21TexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21TexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.TexSubImage3DEXT = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format, Type uint32, pixels unsafe.Pointer) {
		C.gl21TexSubImage3DEXT(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(Type), pixels)
	}

	glc.Translated = func(x, y, z float64) {
		C.gl21Translated(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Translatef = func(x, y, z float32) {
		C.gl21Translatef(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex2s = func(x, y int16) {
		C.gl21Vertex2s(glc.context, C.GLshort(x), C.GLshort(y))
	}

	glc.Vertex2i = func(x, y int32) {
		C.gl21Vertex2i(glc.context, C.GLint(x), C.GLint(y))
	}

	glc.Vertex2f = func(x, y float32) {
		C.gl21Vertex2f(glc.context, C.GLfloat(x), C.GLfloat(y))
	}

	glc.Vertex2d = func(x, y float64) {
		C.gl21Vertex2d(glc.context, C.GLdouble(x), C.GLdouble(y))
	}

	glc.Vertex3s = func(x, y, z int16) {
		C.gl21Vertex3s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z))
	}

	glc.Vertex3i = func(x, y, z int32) {
		C.gl21Vertex3i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z))
	}

	glc.Vertex3f = func(x, y, z float32) {
		C.gl21Vertex3f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}

	glc.Vertex3d = func(x, y, z float64) {
		C.gl21Vertex3d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
	}

	glc.Vertex4s = func(x, y, z, w int16) {
		C.gl21Vertex4s(glc.context, C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
	}

	glc.Vertex4i = func(x, y, z, w int32) {
		C.gl21Vertex4i(glc.context, C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
	}

	glc.Vertex4f = func(x, y, z, w float32) {
		C.gl21Vertex4f(glc.context, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
	}

	glc.Vertex4d = func(x, y, z, w float64) {
		C.gl21Vertex4d(glc.context, C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
	}

	glc.VertexPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl21VertexPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.Viewport = func(x, y int32, width, height int32) {
		C.gl21Viewport(glc.context, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.AreTexturesResident = func(textures []uint32) (status bool, residencies []bool) {
		var cRes *C.GLboolean
		status = C.gl21AreTexturesResident(glc.context, C.GLsizei(len(textures)), (*C.GLuint)(unsafe.Pointer(&textures[0])), cRes) != 0
		residencies = make([]bool, len(textures))
		for i := 0; i < len(textures); i++ {
			residencies[i] = (*(*C.GLboolean)(unsafe.Pointer(uintptr(unsafe.Pointer(cRes)) + uintptr(i)))) != 0
		}
		return
	}

	glc.ArrayElement = func(i int32) {
		C.gl21ArrayElement(glc.context, C.GLint(i))
	}

	glc.DrawArrays = func(mode uint32, first int32, count int32) {
		C.gl21DrawArrays(glc.context, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}

	glc.DrawElements = func(mode uint32, count int32, Type uint32, indices unsafe.Pointer) {
		C.gl21DrawElements(glc.context, C.GLenum(mode), C.GLsizei(count), C.GLenum(Type), indices)
	}

	glc.GetPointerv = func(pname uint32, params unsafe.Pointer) {
		C.gl21GetPointerv(glc.context, C.GLenum(pname), params)
	}

	glc.PolygonOffset = func(factor, units float32) {
		C.gl21PolygonOffset(glc.context, C.GLfloat(factor), C.GLfloat(units))
	}

	glc.CopyTexImage1D = func(target uint32, level int32, internalFormat uint32, x, y int32, width int32, border int32) {
		C.gl21CopyTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
	}

	glc.CopyTexImage2D = func(target uint32, level int32, internalFormat uint32, x, y int32, width, height int32, border int32) {
		C.gl21CopyTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}

	glc.CopyTexSubImage1D = func(target uint32, level, xoffset, x, y int32, width int32) {
		C.gl21CopyTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
	}

	glc.CopyTexSubImage2D = func(target uint32, level, xoffset, yoffset, x, y int32, width, height int32) {
		C.gl21CopyTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.BindTexture = func(target uint32, texture uint32) {
		C.gl21BindTexture(glc.context, C.GLenum(target), C.GLuint(texture))
	}

	glc.DeleteTextures = func(n int32, textures []uint32) {
		C.gl21DeleteTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(&textures[0])))
	}

	glc.GenTextures = func(n int32, textures []uint32) {
		C.gl21GenTextures(glc.context, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(&textures[0])))
	}

	glc.IsTexture = func(texture uint32) bool {
		return C.gl21IsTexture(glc.context, C.GLuint(texture)) != 0
	}

	glc.ColorPointer = func(size int32, Type uint32, stride int32, pointer unsafe.Pointer) {
		C.gl21ColorPointer(glc.context, C.GLint(size), C.GLenum(Type), C.GLsizei(stride), pointer)
	}

	glc.BlendColorEXT = func(red, green, blue, alpha float32) {
		C.gl21BlendColorEXT(glc.context, C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	}

	glc.BlendEquation = func(mode uint32) {
		C.gl21BlendEquation(glc.context, C.GLenum(mode))
	}

	glc.CopyTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset, x, y int32, width, height int32) {
		C.gl21CopyTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}

	glc.ActiveTexture = func(texture uint32) {
		C.gl21ActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.ClientActiveTexture = func(texture uint32) {
		C.gl21ClientActiveTexture(glc.context, C.GLenum(texture))
	}

	glc.CompressedTexImage1D = func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl21CompressedTexImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage2D = func(target uint32, level int32, internalformat uint32, width, height int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl21CompressedTexImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexImage3D = func(target uint32, level int32, internalformat uint32, width, height, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
		C.gl21CompressedTexImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage1D = func(target uint32, level, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl21CompressedTexSubImage1D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage2D = func(target uint32, level, xoffset, yoffset int32, width, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl21CompressedTexSubImage2D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.CompressedTexSubImage3D = func(target uint32, level, xoffset, yoffset, zoffset int32, width, height, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
		C.gl21CompressedTexSubImage3D(glc.context, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), data)
	}

	glc.BlendFuncSeparate = func(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
		C.gl21BlendFuncSeparate(glc.context, C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
	}

	glc.BeginQuery = func(target uint32, id uint32) {
		C.gl21BeginQuery(glc.context, C.GLenum(target), C.GLuint(id))
	}

	glc.BindBuffer = func(target uint32, buffer uint32) {
		C.gl21BindBuffer(glc.context, C.GLenum(target), C.GLuint(buffer))
	}

	glc.BufferData = func(target uint32, size int32, data unsafe.Pointer, usage uint32) {
		C.gl21BufferData(glc.context, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}

	glc.BufferSubData = func(target, offset uint32, size int32, data unsafe.Pointer) {
		C.gl21BufferSubData(glc.context, C.GLenum(target), C.GLenum(offset), C.GLsizeiptr(size), data)
	}

	glc.AttachShader = func(program, shader uint32) {
		C.gl21AttachShader(glc.context, C.GLuint(program), C.GLuint(shader))
	}

	glc.BindAttribLocation = func(program, index uint32, name string) {
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(&cstr))
		C.gl21BindAttribLocation(glc.context, C.GLuint(program), C.GLuint(index), (*C.GLchar)(unsafe.Pointer(cstr)))
		return
	}

	glc.BlendEquationSeperate = func(modeRGB, modeAlpha uint32) {
		C.gl21BlendEquationSeperate(glc.context, C.GLenum(modeRGB), C.GLenum(modeAlpha))
	}

	glc.CompileShader = func(shader uint32) {
		C.gl21CompileShader(glc.context, C.GLuint(shader))
	}

	glc.CreateProgram = func() uint32 {
		return uint32(C.gl21CreateProgram(glc.context))
	}

	glc.CreateShader = func(shaderType uint32) uint32 {
		return uint32(C.gl21CreateShader(glc.context, C.GLenum(shaderType)))
	}

	return glc
}
