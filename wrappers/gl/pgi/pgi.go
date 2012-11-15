// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// PGI_misc_hints: http://www.opengl.org/registry/specs/PGI/misc_hints.txt
// 
// PGI_vertex_hints: http://www.opengl.org/registry/specs/PGI/vertex_hints.txt
// 
package pgi

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
// //  PGI_misc_hints
// void (APIENTRYP ptrglHintPGI)(GLenum target, GLint mode);
// //  PGI_vertex_hints
// 
// //  PGI_misc_hints
// void goglHintPGI(GLenum target, GLint mode) {
// 	(*ptrglHintPGI)(target, mode);
// }
// //  PGI_vertex_hints
// 
// int init_PGI_misc_hints() {
// 	ptrglHintPGI = goglGetProcAddress("glHintPGI");
// 	if(ptrglHintPGI == NULL) return 1;
// 	return 0;
// }
// int init_PGI_vertex_hints() {
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

// PGI_misc_hints
const (
	ALLOW_DRAW_FRG_HINT_PGI = 0x1A210
	ALLOW_DRAW_MEM_HINT_PGI = 0x1A211
	ALLOW_DRAW_OBJ_HINT_PGI = 0x1A20E
	ALLOW_DRAW_WIN_HINT_PGI = 0x1A20F
	ALWAYS_FAST_HINT_PGI = 0x1A20C
	ALWAYS_SOFT_HINT_PGI = 0x1A20D
	BACK_NORMALS_HINT_PGI = 0x1A223
	CLIP_FAR_HINT_PGI = 0x1A221
	CLIP_NEAR_HINT_PGI = 0x1A220
	CONSERVE_MEMORY_HINT_PGI = 0x1A1FD
	FULL_STIPPLE_HINT_PGI = 0x1A219
	NATIVE_GRAPHICS_BEGIN_HINT_PGI = 0x1A203
	NATIVE_GRAPHICS_END_HINT_PGI = 0x1A204
	NATIVE_GRAPHICS_HANDLE_PGI = 0x1A202
	PREFER_DOUBLEBUFFER_HINT_PGI = 0x1A1F8
	RECLAIM_MEMORY_HINT_PGI = 0x1A1FE
	STRICT_DEPTHFUNC_HINT_PGI = 0x1A216
	STRICT_LIGHTING_HINT_PGI = 0x1A217
	STRICT_SCISSOR_HINT_PGI = 0x1A218
	WIDE_LINE_HINT_PGI = 0x1A222
)
// PGI_vertex_hints
const (
	COLOR3_BIT_PGI = 0x00010000
	COLOR4_BIT_PGI = 0x00020000
	EDGEFLAG_BIT_PGI = 0x00040000
	INDEX_BIT_PGI = 0x00080000
	MATERIAL_SIDE_HINT_PGI = 0x1A22C
	MAT_AMBIENT_AND_DIFFUSE_BIT_PGI = 0x00200000
	MAT_AMBIENT_BIT_PGI = 0x00100000
	MAT_COLOR_INDEXES_BIT_PGI = 0x01000000
	MAT_DIFFUSE_BIT_PGI = 0x00400000
	MAT_EMISSION_BIT_PGI = 0x00800000
	MAT_SHININESS_BIT_PGI = 0x02000000
	MAT_SPECULAR_BIT_PGI = 0x04000000
	MAX_VERTEX_HINT_PGI = 0x1A22D
	NORMAL_BIT_PGI = 0x08000000
	TEXCOORD1_BIT_PGI = 0x10000000
	TEXCOORD2_BIT_PGI = 0x20000000
	TEXCOORD3_BIT_PGI = 0x40000000
	TEXCOORD4_BIT_PGI = 0x80000000
	VERTEX23_BIT_PGI = 0x00000004
	VERTEX4_BIT_PGI = 0x00000008
	VERTEX_CONSISTENT_HINT_PGI = 0x1A22B
	VERTEX_DATA_HINT_PGI = 0x1A22A
)
// PGI_misc_hints

func HintPGI(target Enum, mode Int)  {
	C.goglHintPGI((C.GLenum)(target), (C.GLint)(mode))
}
// PGI_vertex_hints

func InitPgiMiscHints() error {
	var ret C.int
	if ret = C.init_PGI_misc_hints(); ret != 0 {
		return errors.New("unable to initialize PGI_misc_hints")
	}
	return nil
}
func InitPgiVertexHints() error {
	var ret C.int
	if ret = C.init_PGI_vertex_hints(); ret != 0 {
		return errors.New("unable to initialize PGI_vertex_hints")
	}
	return nil
}
// EOF