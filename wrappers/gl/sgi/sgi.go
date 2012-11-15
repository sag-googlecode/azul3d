// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// SGI_color_matrix: http://www.opengl.org/registry/specs/SGI/color_matrix.txt
// 
// SGI_color_table: http://www.opengl.org/registry/specs/SGI/color_table.txt
// 
// SGI_texture_color_table: http://www.opengl.org/registry/specs/SGI/texture_color_table.txt
// 
package sgi

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
// //  SGI_color_matrix
// //  SGI_color_table
// void (APIENTRYP ptrglColorTableSGI)(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* table);
// void (APIENTRYP ptrglColorTableParameterfvSGI)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglColorTableParameterivSGI)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglCopyColorTableSGI)(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglGetColorTableSGI)(GLenum target, GLenum format, GLenum type, GLvoid* table);
// void (APIENTRYP ptrglGetColorTableParameterfvSGI)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetColorTableParameterivSGI)(GLenum target, GLenum pname, GLint* params);
// //  SGI_texture_color_table
// 
// //  SGI_color_matrix
// //  SGI_color_table
// void goglColorTableSGI(GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type_, GLvoid* table) {
// 	(*ptrglColorTableSGI)(target, internalformat, width, format, type_, table);
// }
// void goglColorTableParameterfvSGI(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglColorTableParameterfvSGI)(target, pname, params);
// }
// void goglColorTableParameterivSGI(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglColorTableParameterivSGI)(target, pname, params);
// }
// void goglCopyColorTableSGI(GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
// 	(*ptrglCopyColorTableSGI)(target, internalformat, x, y, width);
// }
// void goglGetColorTableSGI(GLenum target, GLenum format, GLenum type_, GLvoid* table) {
// 	(*ptrglGetColorTableSGI)(target, format, type_, table);
// }
// void goglGetColorTableParameterfvSGI(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetColorTableParameterfvSGI)(target, pname, params);
// }
// void goglGetColorTableParameterivSGI(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetColorTableParameterivSGI)(target, pname, params);
// }
// //  SGI_texture_color_table
// 
// int init_SGI_color_matrix() {
// 	return 0;
// }
// int init_SGI_color_table() {
// 	ptrglColorTableSGI = goglGetProcAddress("glColorTableSGI");
// 	if(ptrglColorTableSGI == NULL) return 1;
// 	ptrglColorTableParameterfvSGI = goglGetProcAddress("glColorTableParameterfvSGI");
// 	if(ptrglColorTableParameterfvSGI == NULL) return 1;
// 	ptrglColorTableParameterivSGI = goglGetProcAddress("glColorTableParameterivSGI");
// 	if(ptrglColorTableParameterivSGI == NULL) return 1;
// 	ptrglCopyColorTableSGI = goglGetProcAddress("glCopyColorTableSGI");
// 	if(ptrglCopyColorTableSGI == NULL) return 1;
// 	ptrglGetColorTableSGI = goglGetProcAddress("glGetColorTableSGI");
// 	if(ptrglGetColorTableSGI == NULL) return 1;
// 	ptrglGetColorTableParameterfvSGI = goglGetProcAddress("glGetColorTableParameterfvSGI");
// 	if(ptrglGetColorTableParameterfvSGI == NULL) return 1;
// 	ptrglGetColorTableParameterivSGI = goglGetProcAddress("glGetColorTableParameterivSGI");
// 	if(ptrglGetColorTableParameterivSGI == NULL) return 1;
// 	return 0;
// }
// int init_SGI_texture_color_table() {
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

// SGI_color_matrix
const (
	COLOR_MATRIX_SGI = 0x80B1
	COLOR_MATRIX_STACK_DEPTH_SGI = 0x80B2
	MAX_COLOR_MATRIX_STACK_DEPTH_SGI = 0x80B3
	POST_COLOR_MATRIX_ALPHA_BIAS_SGI = 0x80BB
	POST_COLOR_MATRIX_ALPHA_SCALE_SGI = 0x80B7
	POST_COLOR_MATRIX_BLUE_BIAS_SGI = 0x80BA
	POST_COLOR_MATRIX_BLUE_SCALE_SGI = 0x80B6
	POST_COLOR_MATRIX_GREEN_BIAS_SGI = 0x80B9
	POST_COLOR_MATRIX_GREEN_SCALE_SGI = 0x80B5
	POST_COLOR_MATRIX_RED_BIAS_SGI = 0x80B8
	POST_COLOR_MATRIX_RED_SCALE_SGI = 0x80B4
)
// SGI_color_table
const (
	COLOR_TABLE_ALPHA_SIZE_SGI = 0x80DD
	COLOR_TABLE_BIAS_SGI = 0x80D7
	COLOR_TABLE_BLUE_SIZE_SGI = 0x80DC
	COLOR_TABLE_FORMAT_SGI = 0x80D8
	COLOR_TABLE_GREEN_SIZE_SGI = 0x80DB
	COLOR_TABLE_INTENSITY_SIZE_SGI = 0x80DF
	COLOR_TABLE_LUMINANCE_SIZE_SGI = 0x80DE
	COLOR_TABLE_RED_SIZE_SGI = 0x80DA
	COLOR_TABLE_SCALE_SGI = 0x80D6
	COLOR_TABLE_SGI = 0x80D0
	COLOR_TABLE_WIDTH_SGI = 0x80D9
	POST_COLOR_MATRIX_COLOR_TABLE_SGI = 0x80D2
	POST_CONVOLUTION_COLOR_TABLE_SGI = 0x80D1
	PROXY_COLOR_TABLE_SGI = 0x80D3
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE_SGI = 0x80D5
	PROXY_POST_CONVOLUTION_COLOR_TABLE_SGI = 0x80D4
)
// SGI_depth_pass_instrument
const (
	DEPTH_PASS_INSTRUMENT_COUNTERS_SGIX = 0x8311
	DEPTH_PASS_INSTRUMENT_MAX_SGIX = 0x8312
	DEPTH_PASS_INSTRUMENT_SGIX = 0x8310
)
// SGI_texture_color_table
const (
	PROXY_TEXTURE_COLOR_TABLE_SGI = 0x80BD
	TEXTURE_COLOR_TABLE_SGI = 0x80BC
)
// SGI_color_matrix

// SGI_color_table

func ColorTableSGI(target Enum, internalformat Enum, width Sizei, format Enum, type_ Enum, table Pointer)  {
	C.goglColorTableSGI((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(table))
}
func ColorTableParameterfvSGI(target Enum, pname Enum, params *Float)  {
	C.goglColorTableParameterfvSGI((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func ColorTableParameterivSGI(target Enum, pname Enum, params *Int)  {
	C.goglColorTableParameterivSGI((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func CopyColorTableSGI(target Enum, internalformat Enum, x Int, y Int, width Sizei)  {
	C.goglCopyColorTableSGI((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func GetColorTableSGI(target Enum, format Enum, type_ Enum, table Pointer)  {
	C.goglGetColorTableSGI((C.GLenum)(target), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(table))
}
func GetColorTableParameterfvSGI(target Enum, pname Enum, params *Float)  {
	C.goglGetColorTableParameterfvSGI((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetColorTableParameterivSGI(target Enum, pname Enum, params *Int)  {
	C.goglGetColorTableParameterivSGI((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// SGI_texture_color_table

func InitSgiColorMatrix() error {
	var ret C.int
	if ret = C.init_SGI_color_matrix(); ret != 0 {
		return errors.New("unable to initialize SGI_color_matrix")
	}
	return nil
}
func InitSgiColorTable() error {
	var ret C.int
	if ret = C.init_SGI_color_table(); ret != 0 {
		return errors.New("unable to initialize SGI_color_table")
	}
	return nil
}
func InitSgiTextureColorTable() error {
	var ret C.int
	if ret = C.init_SGI_texture_color_table(); ret != 0 {
		return errors.New("unable to initialize SGI_texture_color_table")
	}
	return nil
}
// EOF