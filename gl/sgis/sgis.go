// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// SGIS_detail_texture: http://www.opengl.org/registry/specs/SGIS/detail_texture.txt
// 
// SGIS_fog_function: http://www.opengl.org/registry/specs/SGIS/fog_function.txt
// 
// SGIS_generate_mipmap: http://www.opengl.org/registry/specs/SGIS/generate_mipmap.txt
// 
// SGIS_multisample: http://www.opengl.org/registry/specs/SGIS/multisample.txt
// 
// SGIS_pixel_texture: http://www.opengl.org/registry/specs/SGIS/pixel_texture.txt
// 
// SGIS_point_line_texgen: http://www.opengl.org/registry/specs/SGIS/point_line_texgen.txt
// 
// SGIS_point_parameters: http://www.opengl.org/registry/specs/SGIS/point_parameters.txt
// 
// SGIS_sharpen_texture: http://www.opengl.org/registry/specs/SGIS/sharpen_texture.txt
// 
// SGIS_texture4D: http://www.opengl.org/registry/specs/SGIS/texture4D.txt
// 
// SGIS_texture_border_clamp: http://www.opengl.org/registry/specs/SGIS/texture_border_clamp.txt
// 
// SGIS_texture_color_mask: http://www.opengl.org/registry/specs/SGIS/texture_color_mask.txt
// 
// SGIS_texture_edge_clamp: http://www.opengl.org/registry/specs/SGIS/texture_edge_clamp.txt
// 
// SGIS_texture_filter4: http://www.opengl.org/registry/specs/SGIS/texture_filter4.txt
// 
// SGIS_texture_lod: http://www.opengl.org/registry/specs/SGIS/texture_lod.txt
// 
package sgis

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// 
// #include <stdlib.h>
// #if defined(__APPLE__)
// #include <dlfcn.h>
// #elif defined(_WIN32)
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
// #else
// #include <X11/Xlib.h>
// #include <GL/glx.h>
// #endif
// 
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// 
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef int GLsizei;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef unsigned short GLhalf;
// typedef float GLfloat;
// typedef float GLclampf;
// typedef double GLdouble;
// typedef double GLclampd;
// typedef void GLvoid;
// 
// #include <stddef.h>
// #ifndef GL_VERSION_2_0
// /* GL type for program/shader text */
// typedef char GLchar;
// #endif
// 
// #ifndef GL_VERSION_1_5
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// #endif
// 
// #ifndef GL_ARB_vertex_buffer_object
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptrARB;
// typedef ptrdiff_t GLsizeiptrARB;
// #endif
// 
// #ifndef GL_ARB_shader_objects
// /* GL types for program/shader text and shader object handles */
// typedef char GLcharARB;
// typedef unsigned int GLhandleARB;
// #endif
// 
// /* GL type for "half" precision (s10e5) float data in host memory */
// #ifndef GL_ARB_half_float_pixel
// typedef unsigned short GLhalfARB;
// #endif
// 
// #ifndef GL_NV_half_float
// typedef unsigned short GLhalfNV;
// #endif
// 
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// 
// #ifndef GL_EXT_timer_query
// typedef int64_t GLint64EXT;
// typedef uint64_t GLuint64EXT;
// #endif
// 
// #ifndef GL_ARB_sync
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef struct __GLsync *GLsync;
// #endif
// 
// #ifndef GL_ARB_cl_event
// /* These incomplete types let us declare types compatible with OpenCL's cl_context and cl_event */
// struct _cl_context;
// struct _cl_event;
// #endif
// 
// #ifndef GL_ARB_debug_output
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// 
// #ifndef GL_AMD_debug_output
// typedef void (APIENTRY *GLDEBUGPROCAMD)(GLuint id,GLenum category,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// 
// #ifndef GL_NV_vdpau_interop
// typedef GLintptr GLvdpauSurfaceNV;
// #endif
// 
// #ifdef _WIN32
// static HMODULE opengl32 = NULL;
// #endif
// 
// static void* goglGetProcAddress(const char* name) { 
// #ifdef __APPLE__
// 	return dlsym(RTLD_DEFAULT, name);
// #elif _WIN32
// 	void* pf = wglGetProcAddress((LPCSTR)name);
// 	if(pf) {
// 		return pf;
// 	}
// 	if(opengl32 == NULL) {
// 		opengl32 = LoadLibraryA("opengl32.dll");
// 	}
// 	return GetProcAddress(opengl32, (LPCSTR)name);
// #else
// 	return glXGetProcAddress((const GLubyte*)name);
// #endif
// }
// 
// //  SGIS_detail_texture
// void (APIENTRYP ptrglDetailTexFuncSGIS)(GLenum target, GLsizei n, GLfloat* points);
// void (APIENTRYP ptrglGetDetailTexFuncSGIS)(GLenum target, GLfloat* points);
// //  SGIS_fog_function
// void (APIENTRYP ptrglFogFuncSGIS)(GLsizei n, GLfloat* points);
// void (APIENTRYP ptrglGetFogFuncSGIS)(GLfloat* points);
// //  SGIS_generate_mipmap
// //  SGIS_multisample
// void (APIENTRYP ptrglSampleMaskSGIS)(GLclampf value, GLboolean invert);
// void (APIENTRYP ptrglSamplePatternSGIS)(GLenum pattern);
// //  SGIS_pixel_texture
// void (APIENTRYP ptrglPixelTexGenParameteriSGIS)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPixelTexGenParameterivSGIS)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglPixelTexGenParameterfSGIS)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPixelTexGenParameterfvSGIS)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetPixelTexGenParameterivSGIS)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetPixelTexGenParameterfvSGIS)(GLenum pname, GLfloat* params);
// //  SGIS_point_line_texgen
// //  SGIS_point_parameters
// void (APIENTRYP ptrglPointParameterfSGIS)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPointParameterfvSGIS)(GLenum pname, GLfloat* params);
// //  SGIS_sharpen_texture
// void (APIENTRYP ptrglSharpenTexFuncSGIS)(GLenum target, GLsizei n, GLfloat* points);
// void (APIENTRYP ptrglGetSharpenTexFuncSGIS)(GLenum target, GLfloat* points);
// //  SGIS_texture4D
// void (APIENTRYP ptrglTexImage4DSGIS)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLsizei size4d, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglTexSubImage4DSGIS)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint woffset, GLsizei width, GLsizei height, GLsizei depth, GLsizei size4d, GLenum format, GLenum type, GLvoid* pixels);
// //  SGIS_texture_border_clamp
// //  SGIS_texture_color_mask
// void (APIENTRYP ptrglTextureColorMaskSGIS)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
// //  SGIS_texture_edge_clamp
// //  SGIS_texture_filter4
// void (APIENTRYP ptrglGetTexFilterFuncSGIS)(GLenum target, GLenum filter, GLfloat* weights);
// void (APIENTRYP ptrglTexFilterFuncSGIS)(GLenum target, GLenum filter, GLsizei n, GLfloat* weights);
// //  SGIS_texture_lod
// 
// //  SGIS_detail_texture
// void goglDetailTexFuncSGIS(GLenum target, GLsizei n, GLfloat* points) {
// 	(*ptrglDetailTexFuncSGIS)(target, n, points);
// }
// void goglGetDetailTexFuncSGIS(GLenum target, GLfloat* points) {
// 	(*ptrglGetDetailTexFuncSGIS)(target, points);
// }
// //  SGIS_fog_function
// void goglFogFuncSGIS(GLsizei n, GLfloat* points) {
// 	(*ptrglFogFuncSGIS)(n, points);
// }
// void goglGetFogFuncSGIS(GLfloat* points) {
// 	(*ptrglGetFogFuncSGIS)(points);
// }
// //  SGIS_generate_mipmap
// //  SGIS_multisample
// void goglSampleMaskSGIS(GLclampf value, GLboolean invert) {
// 	(*ptrglSampleMaskSGIS)(value, invert);
// }
// void goglSamplePatternSGIS(GLenum pattern) {
// 	(*ptrglSamplePatternSGIS)(pattern);
// }
// //  SGIS_pixel_texture
// void goglPixelTexGenParameteriSGIS(GLenum pname, GLint param) {
// 	(*ptrglPixelTexGenParameteriSGIS)(pname, param);
// }
// void goglPixelTexGenParameterivSGIS(GLenum pname, GLint* params) {
// 	(*ptrglPixelTexGenParameterivSGIS)(pname, params);
// }
// void goglPixelTexGenParameterfSGIS(GLenum pname, GLfloat param) {
// 	(*ptrglPixelTexGenParameterfSGIS)(pname, param);
// }
// void goglPixelTexGenParameterfvSGIS(GLenum pname, GLfloat* params) {
// 	(*ptrglPixelTexGenParameterfvSGIS)(pname, params);
// }
// void goglGetPixelTexGenParameterivSGIS(GLenum pname, GLint* params) {
// 	(*ptrglGetPixelTexGenParameterivSGIS)(pname, params);
// }
// void goglGetPixelTexGenParameterfvSGIS(GLenum pname, GLfloat* params) {
// 	(*ptrglGetPixelTexGenParameterfvSGIS)(pname, params);
// }
// //  SGIS_point_line_texgen
// //  SGIS_point_parameters
// void goglPointParameterfSGIS(GLenum pname, GLfloat param) {
// 	(*ptrglPointParameterfSGIS)(pname, param);
// }
// void goglPointParameterfvSGIS(GLenum pname, GLfloat* params) {
// 	(*ptrglPointParameterfvSGIS)(pname, params);
// }
// //  SGIS_sharpen_texture
// void goglSharpenTexFuncSGIS(GLenum target, GLsizei n, GLfloat* points) {
// 	(*ptrglSharpenTexFuncSGIS)(target, n, points);
// }
// void goglGetSharpenTexFuncSGIS(GLenum target, GLfloat* points) {
// 	(*ptrglGetSharpenTexFuncSGIS)(target, points);
// }
// //  SGIS_texture4D
// void goglTexImage4DSGIS(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLsizei size4d, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexImage4DSGIS)(target, level, internalformat, width, height, depth, size4d, border, format, type_, pixels);
// }
// void goglTexSubImage4DSGIS(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint woffset, GLsizei width, GLsizei height, GLsizei depth, GLsizei size4d, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexSubImage4DSGIS)(target, level, xoffset, yoffset, zoffset, woffset, width, height, depth, size4d, format, type_, pixels);
// }
// //  SGIS_texture_border_clamp
// //  SGIS_texture_color_mask
// void goglTextureColorMaskSGIS(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
// 	(*ptrglTextureColorMaskSGIS)(red, green, blue, alpha);
// }
// //  SGIS_texture_edge_clamp
// //  SGIS_texture_filter4
// void goglGetTexFilterFuncSGIS(GLenum target, GLenum filter, GLfloat* weights) {
// 	(*ptrglGetTexFilterFuncSGIS)(target, filter, weights);
// }
// void goglTexFilterFuncSGIS(GLenum target, GLenum filter, GLsizei n, GLfloat* weights) {
// 	(*ptrglTexFilterFuncSGIS)(target, filter, n, weights);
// }
// //  SGIS_texture_lod
// 
// int init_SGIS_detail_texture() {
// 	ptrglDetailTexFuncSGIS = goglGetProcAddress("glDetailTexFuncSGIS");
// 	if(ptrglDetailTexFuncSGIS == NULL) return 1;
// 	ptrglGetDetailTexFuncSGIS = goglGetProcAddress("glGetDetailTexFuncSGIS");
// 	if(ptrglGetDetailTexFuncSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_fog_function() {
// 	ptrglFogFuncSGIS = goglGetProcAddress("glFogFuncSGIS");
// 	if(ptrglFogFuncSGIS == NULL) return 1;
// 	ptrglGetFogFuncSGIS = goglGetProcAddress("glGetFogFuncSGIS");
// 	if(ptrglGetFogFuncSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_generate_mipmap() {
// 	return 0;
// }
// int init_SGIS_multisample() {
// 	ptrglSampleMaskSGIS = goglGetProcAddress("glSampleMaskSGIS");
// 	if(ptrglSampleMaskSGIS == NULL) return 1;
// 	ptrglSamplePatternSGIS = goglGetProcAddress("glSamplePatternSGIS");
// 	if(ptrglSamplePatternSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_pixel_texture() {
// 	ptrglPixelTexGenParameteriSGIS = goglGetProcAddress("glPixelTexGenParameteriSGIS");
// 	if(ptrglPixelTexGenParameteriSGIS == NULL) return 1;
// 	ptrglPixelTexGenParameterivSGIS = goglGetProcAddress("glPixelTexGenParameterivSGIS");
// 	if(ptrglPixelTexGenParameterivSGIS == NULL) return 1;
// 	ptrglPixelTexGenParameterfSGIS = goglGetProcAddress("glPixelTexGenParameterfSGIS");
// 	if(ptrglPixelTexGenParameterfSGIS == NULL) return 1;
// 	ptrglPixelTexGenParameterfvSGIS = goglGetProcAddress("glPixelTexGenParameterfvSGIS");
// 	if(ptrglPixelTexGenParameterfvSGIS == NULL) return 1;
// 	ptrglGetPixelTexGenParameterivSGIS = goglGetProcAddress("glGetPixelTexGenParameterivSGIS");
// 	if(ptrglGetPixelTexGenParameterivSGIS == NULL) return 1;
// 	ptrglGetPixelTexGenParameterfvSGIS = goglGetProcAddress("glGetPixelTexGenParameterfvSGIS");
// 	if(ptrglGetPixelTexGenParameterfvSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_point_line_texgen() {
// 	return 0;
// }
// int init_SGIS_point_parameters() {
// 	ptrglPointParameterfSGIS = goglGetProcAddress("glPointParameterfSGIS");
// 	if(ptrglPointParameterfSGIS == NULL) return 1;
// 	ptrglPointParameterfvSGIS = goglGetProcAddress("glPointParameterfvSGIS");
// 	if(ptrglPointParameterfvSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_sharpen_texture() {
// 	ptrglSharpenTexFuncSGIS = goglGetProcAddress("glSharpenTexFuncSGIS");
// 	if(ptrglSharpenTexFuncSGIS == NULL) return 1;
// 	ptrglGetSharpenTexFuncSGIS = goglGetProcAddress("glGetSharpenTexFuncSGIS");
// 	if(ptrglGetSharpenTexFuncSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_texture4D() {
// 	ptrglTexImage4DSGIS = goglGetProcAddress("glTexImage4DSGIS");
// 	if(ptrglTexImage4DSGIS == NULL) return 1;
// 	ptrglTexSubImage4DSGIS = goglGetProcAddress("glTexSubImage4DSGIS");
// 	if(ptrglTexSubImage4DSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_texture_border_clamp() {
// 	return 0;
// }
// int init_SGIS_texture_color_mask() {
// 	ptrglTextureColorMaskSGIS = goglGetProcAddress("glTextureColorMaskSGIS");
// 	if(ptrglTextureColorMaskSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_texture_edge_clamp() {
// 	return 0;
// }
// int init_SGIS_texture_filter4() {
// 	ptrglGetTexFilterFuncSGIS = goglGetProcAddress("glGetTexFilterFuncSGIS");
// 	if(ptrglGetTexFilterFuncSGIS == NULL) return 1;
// 	ptrglTexFilterFuncSGIS = goglGetProcAddress("glTexFilterFuncSGIS");
// 	if(ptrglTexFilterFuncSGIS == NULL) return 1;
// 	return 0;
// }
// int init_SGIS_texture_lod() {
// 	return 0;
// }
// 
import "C"
import "unsafe"
import "errors"

type (
	Enum     C.GLenum
	Boolean  C.GLboolean
	Bitfield C.GLbitfield
	Byte     C.GLbyte
	Short    C.GLshort
	Int      C.GLint
	Sizei    C.GLsizei
	Ubyte    C.GLubyte
	Ushort   C.GLushort
	Uint     C.GLuint
	Half     C.GLhalf
	Float    C.GLfloat
	Clampf   C.GLclampf
	Double   C.GLdouble
	Clampd   C.GLclampd
	Char     C.GLchar
	Pointer  unsafe.Pointer
	Sync     C.GLsync
	Int64    C.GLint64
	Uint64   C.GLuint64
	Intptr   C.GLintptr
	Sizeiptr C.GLsizeiptr
)

// SGIS_detail_texture
const (
	DETAIL_TEXTURE_2D_BINDING_SGIS = 0x8096
	DETAIL_TEXTURE_2D_SGIS = 0x8095
	DETAIL_TEXTURE_FUNC_POINTS_SGIS = 0x809C
	DETAIL_TEXTURE_LEVEL_SGIS = 0x809A
	DETAIL_TEXTURE_MODE_SGIS = 0x809B
	LINEAR_DETAIL_ALPHA_SGIS = 0x8098
	LINEAR_DETAIL_COLOR_SGIS = 0x8099
	LINEAR_DETAIL_SGIS = 0x8097
)
// SGIS_fog_function
const (
	FOG_FUNC_POINTS_SGIS = 0x812B
	FOG_FUNC_SGIS = 0x812A
	MAX_FOG_FUNC_POINTS_SGIS = 0x812C
)
// SGIS_generate_mipmap
const (
	GENERATE_MIPMAP_HINT_SGIS = 0x8192
	GENERATE_MIPMAP_SGIS = 0x8191
)
// SGIS_multisample
const (
	X1PASS_SGIS = 0x80A1
	X2PASS_0_SGIS = 0x80A2
	X2PASS_1_SGIS = 0x80A3
	X4PASS_0_SGIS = 0x80A4
	X4PASS_1_SGIS = 0x80A5
	X4PASS_2_SGIS = 0x80A6
	X4PASS_3_SGIS = 0x80A7
	MULTISAMPLE_SGIS = 0x809D
	SAMPLES_SGIS = 0x80A9
	SAMPLE_ALPHA_TO_MASK_SGIS = 0x809E
	SAMPLE_ALPHA_TO_ONE_SGIS = 0x809F
	SAMPLE_BUFFERS_SGIS = 0x80A8
	SAMPLE_MASK_INVERT_SGIS = 0x80AB
	SAMPLE_MASK_SGIS = 0x80A0
	SAMPLE_MASK_VALUE_SGIS = 0x80AA
	SAMPLE_PATTERN_SGIS = 0x80AC
)
// SGIS_pixel_texture
const (
	PIXEL_FRAGMENT_ALPHA_SOURCE_SGIS = 0x8355
	PIXEL_FRAGMENT_RGB_SOURCE_SGIS = 0x8354
	PIXEL_GROUP_COLOR_SGIS = 0x8356
	PIXEL_TEXTURE_SGIS = 0x8353
)
// SGIS_point_line_texgen
const (
	EYE_DISTANCE_TO_LINE_SGIS = 0x81F2
	EYE_DISTANCE_TO_POINT_SGIS = 0x81F0
	EYE_LINE_SGIS = 0x81F6
	EYE_POINT_SGIS = 0x81F4
	OBJECT_DISTANCE_TO_LINE_SGIS = 0x81F3
	OBJECT_DISTANCE_TO_POINT_SGIS = 0x81F1
	OBJECT_LINE_SGIS = 0x81F7
	OBJECT_POINT_SGIS = 0x81F5
)
// SGIS_point_parameters
const (
	DISTANCE_ATTENUATION_SGIS = 0x8129
	POINT_FADE_THRESHOLD_SIZE_SGIS = 0x8128
	POINT_SIZE_MAX_SGIS = 0x8127
	POINT_SIZE_MIN_SGIS = 0x8126
)
// SGIS_sharpen_texture
const (
	LINEAR_SHARPEN_ALPHA_SGIS = 0x80AE
	LINEAR_SHARPEN_COLOR_SGIS = 0x80AF
	LINEAR_SHARPEN_SGIS = 0x80AD
	SHARPEN_TEXTURE_FUNC_POINTS_SGIS = 0x80B0
)
// SGIS_texture4D
const (
	MAX_4D_TEXTURE_SIZE_SGIS = 0x8138
	PACK_IMAGE_DEPTH_SGIS = 0x8131
	PACK_SKIP_VOLUMES_SGIS = 0x8130
	PROXY_TEXTURE_4D_SGIS = 0x8135
	TEXTURE_4DSIZE_SGIS = 0x8136
	TEXTURE_4D_BINDING_SGIS = 0x814F
	TEXTURE_4D_SGIS = 0x8134
	TEXTURE_WRAP_Q_SGIS = 0x8137
	UNPACK_IMAGE_DEPTH_SGIS = 0x8133
	UNPACK_SKIP_VOLUMES_SGIS = 0x8132
)
// SGIS_texture_border_clamp
const (
	CLAMP_TO_BORDER_SGIS = 0x812D
)
// SGIS_texture_color_mask
const (
	TEXTURE_COLOR_WRITEMASK_SGIS = 0x81EF
)
// SGIS_texture_edge_clamp
const (
	CLAMP_TO_EDGE_SGIS = 0x812F
)
// SGIS_texture_filter4
const (
	FILTER4_SGIS = 0x8146
	TEXTURE_FILTER4_SIZE_SGIS = 0x8147
)
// SGIS_texture_lod
const (
	TEXTURE_BASE_LEVEL_SGIS = 0x813C
	TEXTURE_MAX_LEVEL_SGIS = 0x813D
	TEXTURE_MAX_LOD_SGIS = 0x813B
	TEXTURE_MIN_LOD_SGIS = 0x813A
)
// SGIS_texture_select
const (
	DUAL_ALPHA12_SGIS = 0x8112
	DUAL_ALPHA16_SGIS = 0x8113
	DUAL_ALPHA4_SGIS = 0x8110
	DUAL_ALPHA8_SGIS = 0x8111
	DUAL_INTENSITY12_SGIS = 0x811A
	DUAL_INTENSITY16_SGIS = 0x811B
	DUAL_INTENSITY4_SGIS = 0x8118
	DUAL_INTENSITY8_SGIS = 0x8119
	DUAL_LUMINANCE12_SGIS = 0x8116
	DUAL_LUMINANCE16_SGIS = 0x8117
	DUAL_LUMINANCE4_SGIS = 0x8114
	DUAL_LUMINANCE8_SGIS = 0x8115
	DUAL_LUMINANCE_ALPHA4_SGIS = 0x811C
	DUAL_LUMINANCE_ALPHA8_SGIS = 0x811D
	DUAL_TEXTURE_SELECT_SGIS = 0x8124
	QUAD_ALPHA4_SGIS = 0x811E
	QUAD_ALPHA8_SGIS = 0x811F
	QUAD_INTENSITY4_SGIS = 0x8122
	QUAD_INTENSITY8_SGIS = 0x8123
	QUAD_LUMINANCE4_SGIS = 0x8120
	QUAD_LUMINANCE8_SGIS = 0x8121
	QUAD_TEXTURE_SELECT_SGIS = 0x8125
)
// SGIS_detail_texture

func DetailTexFuncSGIS(target Enum, n Sizei, points *Float)  {
	C.goglDetailTexFuncSGIS((C.GLenum)(target), (C.GLsizei)(n), (*C.GLfloat)(points))
}
func GetDetailTexFuncSGIS(target Enum, points *Float)  {
	C.goglGetDetailTexFuncSGIS((C.GLenum)(target), (*C.GLfloat)(points))
}
// SGIS_fog_function

func FogFuncSGIS(n Sizei, points *Float)  {
	C.goglFogFuncSGIS((C.GLsizei)(n), (*C.GLfloat)(points))
}
func GetFogFuncSGIS(points *Float)  {
	C.goglGetFogFuncSGIS((*C.GLfloat)(points))
}
// SGIS_generate_mipmap

// SGIS_multisample

func SampleMaskSGIS(value Clampf, invert Boolean)  {
	C.goglSampleMaskSGIS((C.GLclampf)(value), (C.GLboolean)(invert))
}
func SamplePatternSGIS(pattern Enum)  {
	C.goglSamplePatternSGIS((C.GLenum)(pattern))
}
// SGIS_pixel_texture

func PixelTexGenParameteriSGIS(pname Enum, param Int)  {
	C.goglPixelTexGenParameteriSGIS((C.GLenum)(pname), (C.GLint)(param))
}
func PixelTexGenParameterivSGIS(pname Enum, params *Int)  {
	C.goglPixelTexGenParameterivSGIS((C.GLenum)(pname), (*C.GLint)(params))
}
func PixelTexGenParameterfSGIS(pname Enum, param Float)  {
	C.goglPixelTexGenParameterfSGIS((C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelTexGenParameterfvSGIS(pname Enum, params *Float)  {
	C.goglPixelTexGenParameterfvSGIS((C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetPixelTexGenParameterivSGIS(pname Enum, params *Int)  {
	C.goglGetPixelTexGenParameterivSGIS((C.GLenum)(pname), (*C.GLint)(params))
}
func GetPixelTexGenParameterfvSGIS(pname Enum, params *Float)  {
	C.goglGetPixelTexGenParameterfvSGIS((C.GLenum)(pname), (*C.GLfloat)(params))
}
// SGIS_point_line_texgen

// SGIS_point_parameters

func PointParameterfSGIS(pname Enum, param Float)  {
	C.goglPointParameterfSGIS((C.GLenum)(pname), (C.GLfloat)(param))
}
func PointParameterfvSGIS(pname Enum, params *Float)  {
	C.goglPointParameterfvSGIS((C.GLenum)(pname), (*C.GLfloat)(params))
}
// SGIS_sharpen_texture

func SharpenTexFuncSGIS(target Enum, n Sizei, points *Float)  {
	C.goglSharpenTexFuncSGIS((C.GLenum)(target), (C.GLsizei)(n), (*C.GLfloat)(points))
}
func GetSharpenTexFuncSGIS(target Enum, points *Float)  {
	C.goglGetSharpenTexFuncSGIS((C.GLenum)(target), (*C.GLfloat)(points))
}
// SGIS_texture4D

func TexImage4DSGIS(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, size4d Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage4DSGIS((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLsizei)(size4d), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
func TexSubImage4DSGIS(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, woffset Int, width Sizei, height Sizei, depth Sizei, size4d Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexSubImage4DSGIS((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(woffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLsizei)(size4d), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// SGIS_texture_border_clamp

// SGIS_texture_color_mask

func TextureColorMaskSGIS(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.goglTextureColorMaskSGIS((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
// SGIS_texture_edge_clamp

// SGIS_texture_filter4

func GetTexFilterFuncSGIS(target Enum, filter Enum, weights *Float)  {
	C.goglGetTexFilterFuncSGIS((C.GLenum)(target), (C.GLenum)(filter), (*C.GLfloat)(weights))
}
func TexFilterFuncSGIS(target Enum, filter Enum, n Sizei, weights *Float)  {
	C.goglTexFilterFuncSGIS((C.GLenum)(target), (C.GLenum)(filter), (C.GLsizei)(n), (*C.GLfloat)(weights))
}
// SGIS_texture_lod

func InitSgisDetailTexture() error {
	var ret C.int
	if ret = C.init_SGIS_detail_texture(); ret != 0 {
		return errors.New("unable to initialize SGIS_detail_texture")
	}
	return nil
}
func InitSgisFogFunction() error {
	var ret C.int
	if ret = C.init_SGIS_fog_function(); ret != 0 {
		return errors.New("unable to initialize SGIS_fog_function")
	}
	return nil
}
func InitSgisGenerateMipmap() error {
	var ret C.int
	if ret = C.init_SGIS_generate_mipmap(); ret != 0 {
		return errors.New("unable to initialize SGIS_generate_mipmap")
	}
	return nil
}
func InitSgisMultisample() error {
	var ret C.int
	if ret = C.init_SGIS_multisample(); ret != 0 {
		return errors.New("unable to initialize SGIS_multisample")
	}
	return nil
}
func InitSgisPixelTexture() error {
	var ret C.int
	if ret = C.init_SGIS_pixel_texture(); ret != 0 {
		return errors.New("unable to initialize SGIS_pixel_texture")
	}
	return nil
}
func InitSgisPointLineTexgen() error {
	var ret C.int
	if ret = C.init_SGIS_point_line_texgen(); ret != 0 {
		return errors.New("unable to initialize SGIS_point_line_texgen")
	}
	return nil
}
func InitSgisPointParameters() error {
	var ret C.int
	if ret = C.init_SGIS_point_parameters(); ret != 0 {
		return errors.New("unable to initialize SGIS_point_parameters")
	}
	return nil
}
func InitSgisSharpenTexture() error {
	var ret C.int
	if ret = C.init_SGIS_sharpen_texture(); ret != 0 {
		return errors.New("unable to initialize SGIS_sharpen_texture")
	}
	return nil
}
func InitSgisTexture4d() error {
	var ret C.int
	if ret = C.init_SGIS_texture4D(); ret != 0 {
		return errors.New("unable to initialize SGIS_texture4D")
	}
	return nil
}
func InitSgisTextureBorderClamp() error {
	var ret C.int
	if ret = C.init_SGIS_texture_border_clamp(); ret != 0 {
		return errors.New("unable to initialize SGIS_texture_border_clamp")
	}
	return nil
}
func InitSgisTextureColorMask() error {
	var ret C.int
	if ret = C.init_SGIS_texture_color_mask(); ret != 0 {
		return errors.New("unable to initialize SGIS_texture_color_mask")
	}
	return nil
}
func InitSgisTextureEdgeClamp() error {
	var ret C.int
	if ret = C.init_SGIS_texture_edge_clamp(); ret != 0 {
		return errors.New("unable to initialize SGIS_texture_edge_clamp")
	}
	return nil
}
func InitSgisTextureFilter4() error {
	var ret C.int
	if ret = C.init_SGIS_texture_filter4(); ret != 0 {
		return errors.New("unable to initialize SGIS_texture_filter4")
	}
	return nil
}
func InitSgisTextureLod() error {
	var ret C.int
	if ret = C.init_SGIS_texture_lod(); ret != 0 {
		return errors.New("unable to initialize SGIS_texture_lod")
	}
	return nil
}
// EOF