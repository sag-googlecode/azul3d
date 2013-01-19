// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// IBM_cull_vertex: http://www.opengl.org/registry/specs/IBM/cull_vertex.txt
// 
// IBM_multimode_draw_arrays: http://www.opengl.org/registry/specs/IBM/multimode_draw_arrays.txt
// 
// IBM_rasterpos_clip: http://www.opengl.org/registry/specs/IBM/rasterpos_clip.txt
// 
// IBM_vertex_array_lists: http://www.opengl.org/registry/specs/IBM/vertex_array_lists.txt
// 
package ibm

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
// //  IBM_cull_vertex
// //  IBM_multimode_draw_arrays
// void (APIENTRYP ptrglMultiModeDrawArraysIBM)(GLenum* mode, GLint* first, GLsizei* count, GLsizei primcount, GLint modestride);
// void (APIENTRYP ptrglMultiModeDrawElementsIBM)(GLenum* mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei primcount, GLint modestride);
// //  IBM_rasterpos_clip
// //  IBM_vertex_array_lists
// void (APIENTRYP ptrglColorPointerListIBM)(GLint size, GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// void (APIENTRYP ptrglSecondaryColorPointerListIBM)(GLint size, GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// void (APIENTRYP ptrglEdgeFlagPointerListIBM)(GLint stride, GLboolean** pointer, GLint ptrstride);
// void (APIENTRYP ptrglFogCoordPointerListIBM)(GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// void (APIENTRYP ptrglIndexPointerListIBM)(GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// void (APIENTRYP ptrglNormalPointerListIBM)(GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// void (APIENTRYP ptrglTexCoordPointerListIBM)(GLint size, GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// void (APIENTRYP ptrglVertexPointerListIBM)(GLint size, GLenum type, GLint stride, GLvoid** pointer, GLint ptrstride);
// 
// //  IBM_cull_vertex
// //  IBM_multimode_draw_arrays
// void goglMultiModeDrawArraysIBM(GLenum* mode, GLint* first, GLsizei* count, GLsizei primcount, GLint modestride) {
// 	(*ptrglMultiModeDrawArraysIBM)(mode, first, count, primcount, modestride);
// }
// void goglMultiModeDrawElementsIBM(GLenum* mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei primcount, GLint modestride) {
// 	(*ptrglMultiModeDrawElementsIBM)(mode, count, type_, indices, primcount, modestride);
// }
// //  IBM_rasterpos_clip
// //  IBM_vertex_array_lists
// void goglColorPointerListIBM(GLint size, GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglColorPointerListIBM)(size, type_, stride, pointer, ptrstride);
// }
// void goglSecondaryColorPointerListIBM(GLint size, GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglSecondaryColorPointerListIBM)(size, type_, stride, pointer, ptrstride);
// }
// void goglEdgeFlagPointerListIBM(GLint stride, GLboolean** pointer, GLint ptrstride) {
// 	(*ptrglEdgeFlagPointerListIBM)(stride, pointer, ptrstride);
// }
// void goglFogCoordPointerListIBM(GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglFogCoordPointerListIBM)(type_, stride, pointer, ptrstride);
// }
// void goglIndexPointerListIBM(GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglIndexPointerListIBM)(type_, stride, pointer, ptrstride);
// }
// void goglNormalPointerListIBM(GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglNormalPointerListIBM)(type_, stride, pointer, ptrstride);
// }
// void goglTexCoordPointerListIBM(GLint size, GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglTexCoordPointerListIBM)(size, type_, stride, pointer, ptrstride);
// }
// void goglVertexPointerListIBM(GLint size, GLenum type_, GLint stride, GLvoid** pointer, GLint ptrstride) {
// 	(*ptrglVertexPointerListIBM)(size, type_, stride, pointer, ptrstride);
// }
// 
// int init_IBM_cull_vertex() {
// 	return 0;
// }
// int init_IBM_multimode_draw_arrays() {
// 	ptrglMultiModeDrawArraysIBM = goglGetProcAddress("glMultiModeDrawArraysIBM");
// 	if(ptrglMultiModeDrawArraysIBM == NULL) return 1;
// 	ptrglMultiModeDrawElementsIBM = goglGetProcAddress("glMultiModeDrawElementsIBM");
// 	if(ptrglMultiModeDrawElementsIBM == NULL) return 1;
// 	return 0;
// }
// int init_IBM_rasterpos_clip() {
// 	return 0;
// }
// int init_IBM_vertex_array_lists() {
// 	ptrglColorPointerListIBM = goglGetProcAddress("glColorPointerListIBM");
// 	if(ptrglColorPointerListIBM == NULL) return 1;
// 	ptrglSecondaryColorPointerListIBM = goglGetProcAddress("glSecondaryColorPointerListIBM");
// 	if(ptrglSecondaryColorPointerListIBM == NULL) return 1;
// 	ptrglEdgeFlagPointerListIBM = goglGetProcAddress("glEdgeFlagPointerListIBM");
// 	if(ptrglEdgeFlagPointerListIBM == NULL) return 1;
// 	ptrglFogCoordPointerListIBM = goglGetProcAddress("glFogCoordPointerListIBM");
// 	if(ptrglFogCoordPointerListIBM == NULL) return 1;
// 	ptrglIndexPointerListIBM = goglGetProcAddress("glIndexPointerListIBM");
// 	if(ptrglIndexPointerListIBM == NULL) return 1;
// 	ptrglNormalPointerListIBM = goglGetProcAddress("glNormalPointerListIBM");
// 	if(ptrglNormalPointerListIBM == NULL) return 1;
// 	ptrglTexCoordPointerListIBM = goglGetProcAddress("glTexCoordPointerListIBM");
// 	if(ptrglTexCoordPointerListIBM == NULL) return 1;
// 	ptrglVertexPointerListIBM = goglGetProcAddress("glVertexPointerListIBM");
// 	if(ptrglVertexPointerListIBM == NULL) return 1;
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

// IBM_cull_vertex
const (
	CULL_VERTEX_IBM = 103050
)
// IBM_multimode_draw_arrays
const (
)
// IBM_rasterpos_clip
const (
	RASTER_POSITION_UNCLIPPED_IBM = 0x19262
)
// IBM_texture_mirrored_repeat
const (
	MIRRORED_REPEAT_IBM = 0x8370
)
// IBM_vertex_array_lists
const (
	COLOR_ARRAY_LIST_IBM = 103072
	COLOR_ARRAY_LIST_STRIDE_IBM = 103082
	EDGE_FLAG_ARRAY_LIST_IBM = 103075
	EDGE_FLAG_ARRAY_LIST_STRIDE_IBM = 103085
	FOG_COORDINATE_ARRAY_LIST_IBM = 103076
	FOG_COORDINATE_ARRAY_LIST_STRIDE_IBM = 103086
	INDEX_ARRAY_LIST_IBM = 103073
	INDEX_ARRAY_LIST_STRIDE_IBM = 103083
	NORMAL_ARRAY_LIST_IBM = 103071
	NORMAL_ARRAY_LIST_STRIDE_IBM = 103081
	SECONDARY_COLOR_ARRAY_LIST_IBM = 103077
	SECONDARY_COLOR_ARRAY_LIST_STRIDE_IBM = 103087
	TEXTURE_COORD_ARRAY_LIST_IBM = 103074
	TEXTURE_COORD_ARRAY_LIST_STRIDE_IBM = 103084
	VERTEX_ARRAY_LIST_IBM = 103070
	VERTEX_ARRAY_LIST_STRIDE_IBM = 103080
)
// IBM_cull_vertex

// IBM_multimode_draw_arrays

func MultiModeDrawArraysIBM(mode *Enum, first *Int, count *Sizei, primcount Sizei, modestride Int)  {
	C.goglMultiModeDrawArraysIBM((*C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount), (C.GLint)(modestride))
}
func MultiModeDrawElementsIBM(mode *Enum, count *Sizei, type_ Enum, indices *Pointer, primcount Sizei, modestride Int)  {
	C.goglMultiModeDrawElementsIBM((*C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(primcount), (C.GLint)(modestride))
}
// IBM_rasterpos_clip

// IBM_vertex_array_lists

func ColorPointerListIBM(size Int, type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglColorPointerListIBM((C.GLint)(size), (C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func SecondaryColorPointerListIBM(size Int, type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglSecondaryColorPointerListIBM((C.GLint)(size), (C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func EdgeFlagPointerListIBM(stride Int, pointer **Boolean, ptrstride Int)  {
	C.goglEdgeFlagPointerListIBM((C.GLint)(stride), (**C.GLboolean)(unsafe.Pointer(pointer)), (C.GLint)(ptrstride))
}
func FogCoordPointerListIBM(type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglFogCoordPointerListIBM((C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func IndexPointerListIBM(type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglIndexPointerListIBM((C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func NormalPointerListIBM(type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglNormalPointerListIBM((C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func TexCoordPointerListIBM(size Int, type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglTexCoordPointerListIBM((C.GLint)(size), (C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func VertexPointerListIBM(size Int, type_ Enum, stride Int, pointer *Pointer, ptrstride Int)  {
	C.goglVertexPointerListIBM((C.GLint)(size), (C.GLenum)(type_), (C.GLint)(stride), (*unsafe.Pointer)(pointer), (C.GLint)(ptrstride))
}
func InitIbmCullVertex() error {
	var ret C.int
	if ret = C.init_IBM_cull_vertex(); ret != 0 {
		return errors.New("unable to initialize IBM_cull_vertex")
	}
	return nil
}
func InitIbmMultimodeDrawArrays() error {
	var ret C.int
	if ret = C.init_IBM_multimode_draw_arrays(); ret != 0 {
		return errors.New("unable to initialize IBM_multimode_draw_arrays")
	}
	return nil
}
func InitIbmRasterposClip() error {
	var ret C.int
	if ret = C.init_IBM_rasterpos_clip(); ret != 0 {
		return errors.New("unable to initialize IBM_rasterpos_clip")
	}
	return nil
}
func InitIbmVertexArrayLists() error {
	var ret C.int
	if ret = C.init_IBM_vertex_array_lists(); ret != 0 {
		return errors.New("unable to initialize IBM_vertex_array_lists")
	}
	return nil
}
// EOF