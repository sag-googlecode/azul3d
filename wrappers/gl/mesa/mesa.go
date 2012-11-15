// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// MESA_pack_invert: http://www.opengl.org/registry/specs/MESA/pack_invert.txt
// 
// MESA_resize_buffers: http://www.opengl.org/registry/specs/MESA/resize_buffers.txt
// 
// MESA_window_pos: http://www.opengl.org/registry/specs/MESA/window_pos.txt
// 
// MESA_ycbcr_texture: http://www.opengl.org/registry/specs/MESA/ycbcr_texture.txt
// 
package mesa

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
// //  MESA_pack_invert
// //  MESA_resize_buffers
// void (APIENTRYP ptrglResizeBuffersMESA)();
// //  MESA_window_pos
// void (APIENTRYP ptrglWindowPos2dMESA)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrglWindowPos2dvMESA)(GLdouble* v);
// void (APIENTRYP ptrglWindowPos2fMESA)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrglWindowPos2fvMESA)(GLfloat* v);
// void (APIENTRYP ptrglWindowPos2iMESA)(GLint x, GLint y);
// void (APIENTRYP ptrglWindowPos2ivMESA)(GLint* v);
// void (APIENTRYP ptrglWindowPos2sMESA)(GLshort x, GLshort y);
// void (APIENTRYP ptrglWindowPos2svMESA)(GLshort* v);
// void (APIENTRYP ptrglWindowPos3dMESA)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglWindowPos3dvMESA)(GLdouble* v);
// void (APIENTRYP ptrglWindowPos3fMESA)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglWindowPos3fvMESA)(GLfloat* v);
// void (APIENTRYP ptrglWindowPos3iMESA)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglWindowPos3ivMESA)(GLint* v);
// void (APIENTRYP ptrglWindowPos3sMESA)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglWindowPos3svMESA)(GLshort* v);
// void (APIENTRYP ptrglWindowPos4dMESA)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglWindowPos4dvMESA)(GLdouble* v);
// void (APIENTRYP ptrglWindowPos4fMESA)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglWindowPos4fvMESA)(GLfloat* v);
// void (APIENTRYP ptrglWindowPos4iMESA)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglWindowPos4ivMESA)(GLint* v);
// void (APIENTRYP ptrglWindowPos4sMESA)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglWindowPos4svMESA)(GLshort* v);
// //  MESA_ycbcr_texture
// 
// //  MESA_pack_invert
// //  MESA_resize_buffers
// void goglResizeBuffersMESA() {
// 	(*ptrglResizeBuffersMESA)();
// }
// //  MESA_window_pos
// void goglWindowPos2dMESA(GLdouble x, GLdouble y) {
// 	(*ptrglWindowPos2dMESA)(x, y);
// }
// void goglWindowPos2dvMESA(GLdouble* v) {
// 	(*ptrglWindowPos2dvMESA)(v);
// }
// void goglWindowPos2fMESA(GLfloat x, GLfloat y) {
// 	(*ptrglWindowPos2fMESA)(x, y);
// }
// void goglWindowPos2fvMESA(GLfloat* v) {
// 	(*ptrglWindowPos2fvMESA)(v);
// }
// void goglWindowPos2iMESA(GLint x, GLint y) {
// 	(*ptrglWindowPos2iMESA)(x, y);
// }
// void goglWindowPos2ivMESA(GLint* v) {
// 	(*ptrglWindowPos2ivMESA)(v);
// }
// void goglWindowPos2sMESA(GLshort x, GLshort y) {
// 	(*ptrglWindowPos2sMESA)(x, y);
// }
// void goglWindowPos2svMESA(GLshort* v) {
// 	(*ptrglWindowPos2svMESA)(v);
// }
// void goglWindowPos3dMESA(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglWindowPos3dMESA)(x, y, z);
// }
// void goglWindowPos3dvMESA(GLdouble* v) {
// 	(*ptrglWindowPos3dvMESA)(v);
// }
// void goglWindowPos3fMESA(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglWindowPos3fMESA)(x, y, z);
// }
// void goglWindowPos3fvMESA(GLfloat* v) {
// 	(*ptrglWindowPos3fvMESA)(v);
// }
// void goglWindowPos3iMESA(GLint x, GLint y, GLint z) {
// 	(*ptrglWindowPos3iMESA)(x, y, z);
// }
// void goglWindowPos3ivMESA(GLint* v) {
// 	(*ptrglWindowPos3ivMESA)(v);
// }
// void goglWindowPos3sMESA(GLshort x, GLshort y, GLshort z) {
// 	(*ptrglWindowPos3sMESA)(x, y, z);
// }
// void goglWindowPos3svMESA(GLshort* v) {
// 	(*ptrglWindowPos3svMESA)(v);
// }
// void goglWindowPos4dMESA(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglWindowPos4dMESA)(x, y, z, w);
// }
// void goglWindowPos4dvMESA(GLdouble* v) {
// 	(*ptrglWindowPos4dvMESA)(v);
// }
// void goglWindowPos4fMESA(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglWindowPos4fMESA)(x, y, z, w);
// }
// void goglWindowPos4fvMESA(GLfloat* v) {
// 	(*ptrglWindowPos4fvMESA)(v);
// }
// void goglWindowPos4iMESA(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglWindowPos4iMESA)(x, y, z, w);
// }
// void goglWindowPos4ivMESA(GLint* v) {
// 	(*ptrglWindowPos4ivMESA)(v);
// }
// void goglWindowPos4sMESA(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglWindowPos4sMESA)(x, y, z, w);
// }
// void goglWindowPos4svMESA(GLshort* v) {
// 	(*ptrglWindowPos4svMESA)(v);
// }
// //  MESA_ycbcr_texture
// 
// int init_MESA_pack_invert() {
// 	return 0;
// }
// int init_MESA_resize_buffers() {
// 	ptrglResizeBuffersMESA = goglGetProcAddress("glResizeBuffersMESA");
// 	if(ptrglResizeBuffersMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_window_pos() {
// 	ptrglWindowPos2dMESA = goglGetProcAddress("glWindowPos2dMESA");
// 	if(ptrglWindowPos2dMESA == NULL) return 1;
// 	ptrglWindowPos2dvMESA = goglGetProcAddress("glWindowPos2dvMESA");
// 	if(ptrglWindowPos2dvMESA == NULL) return 1;
// 	ptrglWindowPos2fMESA = goglGetProcAddress("glWindowPos2fMESA");
// 	if(ptrglWindowPos2fMESA == NULL) return 1;
// 	ptrglWindowPos2fvMESA = goglGetProcAddress("glWindowPos2fvMESA");
// 	if(ptrglWindowPos2fvMESA == NULL) return 1;
// 	ptrglWindowPos2iMESA = goglGetProcAddress("glWindowPos2iMESA");
// 	if(ptrglWindowPos2iMESA == NULL) return 1;
// 	ptrglWindowPos2ivMESA = goglGetProcAddress("glWindowPos2ivMESA");
// 	if(ptrglWindowPos2ivMESA == NULL) return 1;
// 	ptrglWindowPos2sMESA = goglGetProcAddress("glWindowPos2sMESA");
// 	if(ptrglWindowPos2sMESA == NULL) return 1;
// 	ptrglWindowPos2svMESA = goglGetProcAddress("glWindowPos2svMESA");
// 	if(ptrglWindowPos2svMESA == NULL) return 1;
// 	ptrglWindowPos3dMESA = goglGetProcAddress("glWindowPos3dMESA");
// 	if(ptrglWindowPos3dMESA == NULL) return 1;
// 	ptrglWindowPos3dvMESA = goglGetProcAddress("glWindowPos3dvMESA");
// 	if(ptrglWindowPos3dvMESA == NULL) return 1;
// 	ptrglWindowPos3fMESA = goglGetProcAddress("glWindowPos3fMESA");
// 	if(ptrglWindowPos3fMESA == NULL) return 1;
// 	ptrglWindowPos3fvMESA = goglGetProcAddress("glWindowPos3fvMESA");
// 	if(ptrglWindowPos3fvMESA == NULL) return 1;
// 	ptrglWindowPos3iMESA = goglGetProcAddress("glWindowPos3iMESA");
// 	if(ptrglWindowPos3iMESA == NULL) return 1;
// 	ptrglWindowPos3ivMESA = goglGetProcAddress("glWindowPos3ivMESA");
// 	if(ptrglWindowPos3ivMESA == NULL) return 1;
// 	ptrglWindowPos3sMESA = goglGetProcAddress("glWindowPos3sMESA");
// 	if(ptrglWindowPos3sMESA == NULL) return 1;
// 	ptrglWindowPos3svMESA = goglGetProcAddress("glWindowPos3svMESA");
// 	if(ptrglWindowPos3svMESA == NULL) return 1;
// 	ptrglWindowPos4dMESA = goglGetProcAddress("glWindowPos4dMESA");
// 	if(ptrglWindowPos4dMESA == NULL) return 1;
// 	ptrglWindowPos4dvMESA = goglGetProcAddress("glWindowPos4dvMESA");
// 	if(ptrglWindowPos4dvMESA == NULL) return 1;
// 	ptrglWindowPos4fMESA = goglGetProcAddress("glWindowPos4fMESA");
// 	if(ptrglWindowPos4fMESA == NULL) return 1;
// 	ptrglWindowPos4fvMESA = goglGetProcAddress("glWindowPos4fvMESA");
// 	if(ptrglWindowPos4fvMESA == NULL) return 1;
// 	ptrglWindowPos4iMESA = goglGetProcAddress("glWindowPos4iMESA");
// 	if(ptrglWindowPos4iMESA == NULL) return 1;
// 	ptrglWindowPos4ivMESA = goglGetProcAddress("glWindowPos4ivMESA");
// 	if(ptrglWindowPos4ivMESA == NULL) return 1;
// 	ptrglWindowPos4sMESA = goglGetProcAddress("glWindowPos4sMESA");
// 	if(ptrglWindowPos4sMESA == NULL) return 1;
// 	ptrglWindowPos4svMESA = goglGetProcAddress("glWindowPos4svMESA");
// 	if(ptrglWindowPos4svMESA == NULL) return 1;
// 	return 0;
// }
// int init_MESA_ycbcr_texture() {
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

// MESA_pack_invert
const (
	PACK_INVERT_MESA = 0x8758
)
// MESA_resize_buffers
const (
)
// MESA_window_pos
const (
)
// MESA_ycbcr_texture
const (
	UNSIGNED_SHORT_8_8_MESA = 0x85BA
	UNSIGNED_SHORT_8_8_REV_MESA = 0x85BB
	YCBCR_MESA = 0x8757
)
// MESA_pack_invert

// MESA_resize_buffers

func ResizeBuffersMESA()  {
	C.goglResizeBuffersMESA()
}
// MESA_window_pos

func WindowPos2dMESA(x Double, y Double)  {
	C.goglWindowPos2dMESA((C.GLdouble)(x), (C.GLdouble)(y))
}
func WindowPos2dvMESA(v *Double)  {
	C.goglWindowPos2dvMESA((*C.GLdouble)(v))
}
func WindowPos2fMESA(x Float, y Float)  {
	C.goglWindowPos2fMESA((C.GLfloat)(x), (C.GLfloat)(y))
}
func WindowPos2fvMESA(v *Float)  {
	C.goglWindowPos2fvMESA((*C.GLfloat)(v))
}
func WindowPos2iMESA(x Int, y Int)  {
	C.goglWindowPos2iMESA((C.GLint)(x), (C.GLint)(y))
}
func WindowPos2ivMESA(v *Int)  {
	C.goglWindowPos2ivMESA((*C.GLint)(v))
}
func WindowPos2sMESA(x Short, y Short)  {
	C.goglWindowPos2sMESA((C.GLshort)(x), (C.GLshort)(y))
}
func WindowPos2svMESA(v *Short)  {
	C.goglWindowPos2svMESA((*C.GLshort)(v))
}
func WindowPos3dMESA(x Double, y Double, z Double)  {
	C.goglWindowPos3dMESA((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func WindowPos3dvMESA(v *Double)  {
	C.goglWindowPos3dvMESA((*C.GLdouble)(v))
}
func WindowPos3fMESA(x Float, y Float, z Float)  {
	C.goglWindowPos3fMESA((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func WindowPos3fvMESA(v *Float)  {
	C.goglWindowPos3fvMESA((*C.GLfloat)(v))
}
func WindowPos3iMESA(x Int, y Int, z Int)  {
	C.goglWindowPos3iMESA((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func WindowPos3ivMESA(v *Int)  {
	C.goglWindowPos3ivMESA((*C.GLint)(v))
}
func WindowPos3sMESA(x Short, y Short, z Short)  {
	C.goglWindowPos3sMESA((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func WindowPos3svMESA(v *Short)  {
	C.goglWindowPos3svMESA((*C.GLshort)(v))
}
func WindowPos4dMESA(x Double, y Double, z Double, w Double)  {
	C.goglWindowPos4dMESA((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func WindowPos4dvMESA(v *Double)  {
	C.goglWindowPos4dvMESA((*C.GLdouble)(v))
}
func WindowPos4fMESA(x Float, y Float, z Float, w Float)  {
	C.goglWindowPos4fMESA((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func WindowPos4fvMESA(v *Float)  {
	C.goglWindowPos4fvMESA((*C.GLfloat)(v))
}
func WindowPos4iMESA(x Int, y Int, z Int, w Int)  {
	C.goglWindowPos4iMESA((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func WindowPos4ivMESA(v *Int)  {
	C.goglWindowPos4ivMESA((*C.GLint)(v))
}
func WindowPos4sMESA(x Short, y Short, z Short, w Short)  {
	C.goglWindowPos4sMESA((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func WindowPos4svMESA(v *Short)  {
	C.goglWindowPos4svMESA((*C.GLshort)(v))
}
// MESA_ycbcr_texture

func InitMesaPackInvert() error {
	var ret C.int
	if ret = C.init_MESA_pack_invert(); ret != 0 {
		return errors.New("unable to initialize MESA_pack_invert")
	}
	return nil
}
func InitMesaResizeBuffers() error {
	var ret C.int
	if ret = C.init_MESA_resize_buffers(); ret != 0 {
		return errors.New("unable to initialize MESA_resize_buffers")
	}
	return nil
}
func InitMesaWindowPos() error {
	var ret C.int
	if ret = C.init_MESA_window_pos(); ret != 0 {
		return errors.New("unable to initialize MESA_window_pos")
	}
	return nil
}
func InitMesaYcbcrTexture() error {
	var ret C.int
	if ret = C.init_MESA_ycbcr_texture(); ret != 0 {
		return errors.New("unable to initialize MESA_ycbcr_texture")
	}
	return nil
}
// EOF