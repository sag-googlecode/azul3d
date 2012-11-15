// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// SUN_convolution_border_modes: http://www.opengl.org/registry/specs/SUN/convolution_border_modes.txt
// 
// SUN_global_alpha: http://www.opengl.org/registry/specs/SUN/global_alpha.txt
// 
// SUN_mesh_array: http://www.opengl.org/registry/specs/SUN/mesh_array.txt
// 
// SUN_slice_accum: http://www.opengl.org/registry/specs/SUN/slice_accum.txt
// 
// SUN_triangle_list: http://www.opengl.org/registry/specs/SUN/triangle_list.txt
// 
// SUN_vertex: http://www.opengl.org/registry/specs/SUN/vertex.txt
// 
package sun

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
// //  SUN_convolution_border_modes
// //  SUN_global_alpha
// void (APIENTRYP ptrglGlobalAlphaFactorbSUN)(GLbyte factor);
// void (APIENTRYP ptrglGlobalAlphaFactorsSUN)(GLshort factor);
// void (APIENTRYP ptrglGlobalAlphaFactoriSUN)(GLint factor);
// void (APIENTRYP ptrglGlobalAlphaFactorfSUN)(GLfloat factor);
// void (APIENTRYP ptrglGlobalAlphaFactordSUN)(GLdouble factor);
// void (APIENTRYP ptrglGlobalAlphaFactorubSUN)(GLubyte factor);
// void (APIENTRYP ptrglGlobalAlphaFactorusSUN)(GLushort factor);
// void (APIENTRYP ptrglGlobalAlphaFactoruiSUN)(GLuint factor);
// //  SUN_mesh_array
// void (APIENTRYP ptrglDrawMeshArraysSUN)(GLenum mode, GLint first, GLsizei count, GLsizei width);
// //  SUN_slice_accum
// //  SUN_triangle_list
// void (APIENTRYP ptrglReplacementCodeuiSUN)(GLuint code);
// void (APIENTRYP ptrglReplacementCodeusSUN)(GLushort code);
// void (APIENTRYP ptrglReplacementCodeubSUN)(GLubyte code);
// void (APIENTRYP ptrglReplacementCodeuivSUN)(GLuint* code);
// void (APIENTRYP ptrglReplacementCodeusvSUN)(GLushort* code);
// void (APIENTRYP ptrglReplacementCodeubvSUN)(GLubyte* code);
// void (APIENTRYP ptrglReplacementCodePointerSUN)(GLenum type, GLsizei stride, GLvoid** pointer);
// //  SUN_vertex
// void (APIENTRYP ptrglColor4ubVertex2fSUN)(GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y);
// void (APIENTRYP ptrglColor4ubVertex2fvSUN)(GLubyte* c, GLfloat* v);
// void (APIENTRYP ptrglColor4ubVertex3fSUN)(GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglColor4ubVertex3fvSUN)(GLubyte* c, GLfloat* v);
// void (APIENTRYP ptrglColor3fVertex3fSUN)(GLfloat r, GLfloat g, GLfloat b, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglColor3fVertex3fvSUN)(GLfloat* c, GLfloat* v);
// void (APIENTRYP ptrglNormal3fVertex3fSUN)(GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglNormal3fVertex3fvSUN)(GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglColor4fNormal3fVertex3fSUN)(GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglColor4fNormal3fVertex3fvSUN)(GLfloat* c, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglTexCoord2fVertex3fSUN)(GLfloat s, GLfloat t, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTexCoord2fVertex3fvSUN)(GLfloat* tc, GLfloat* v);
// void (APIENTRYP ptrglTexCoord4fVertex4fSUN)(GLfloat s, GLfloat t, GLfloat p, GLfloat q, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglTexCoord4fVertex4fvSUN)(GLfloat* tc, GLfloat* v);
// void (APIENTRYP ptrglTexCoord2fColor4ubVertex3fSUN)(GLfloat s, GLfloat t, GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTexCoord2fColor4ubVertex3fvSUN)(GLfloat* tc, GLubyte* c, GLfloat* v);
// void (APIENTRYP ptrglTexCoord2fColor3fVertex3fSUN)(GLfloat s, GLfloat t, GLfloat r, GLfloat g, GLfloat b, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTexCoord2fColor3fVertex3fvSUN)(GLfloat* tc, GLfloat* c, GLfloat* v);
// void (APIENTRYP ptrglTexCoord2fNormal3fVertex3fSUN)(GLfloat s, GLfloat t, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTexCoord2fNormal3fVertex3fvSUN)(GLfloat* tc, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglTexCoord2fColor4fNormal3fVertex3fSUN)(GLfloat s, GLfloat t, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTexCoord2fColor4fNormal3fVertex3fvSUN)(GLfloat* tc, GLfloat* c, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglTexCoord4fColor4fNormal3fVertex4fSUN)(GLfloat s, GLfloat t, GLfloat p, GLfloat q, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglTexCoord4fColor4fNormal3fVertex4fvSUN)(GLfloat* tc, GLfloat* c, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiVertex3fSUN)(GLuint rc, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiVertex3fvSUN)(GLuint* rc, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiColor4ubVertex3fSUN)(GLuint rc, GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiColor4ubVertex3fvSUN)(GLuint* rc, GLubyte* c, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiColor3fVertex3fSUN)(GLuint rc, GLfloat r, GLfloat g, GLfloat b, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiColor3fVertex3fvSUN)(GLuint* rc, GLfloat* c, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiNormal3fVertex3fSUN)(GLuint rc, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiNormal3fVertex3fvSUN)(GLuint* rc, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiColor4fNormal3fVertex3fSUN)(GLuint rc, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiColor4fNormal3fVertex3fvSUN)(GLuint* rc, GLfloat* c, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiTexCoord2fVertex3fSUN)(GLuint rc, GLfloat s, GLfloat t, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiTexCoord2fVertex3fvSUN)(GLuint* rc, GLfloat* tc, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fSUN)(GLuint rc, GLfloat s, GLfloat t, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN)(GLuint* rc, GLfloat* tc, GLfloat* n, GLfloat* v);
// void (APIENTRYP ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN)(GLuint rc, GLfloat s, GLfloat t, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN)(GLuint* rc, GLfloat* tc, GLfloat* c, GLfloat* n, GLfloat* v);
// 
// //  SUN_convolution_border_modes
// //  SUN_global_alpha
// void goglGlobalAlphaFactorbSUN(GLbyte factor) {
// 	(*ptrglGlobalAlphaFactorbSUN)(factor);
// }
// void goglGlobalAlphaFactorsSUN(GLshort factor) {
// 	(*ptrglGlobalAlphaFactorsSUN)(factor);
// }
// void goglGlobalAlphaFactoriSUN(GLint factor) {
// 	(*ptrglGlobalAlphaFactoriSUN)(factor);
// }
// void goglGlobalAlphaFactorfSUN(GLfloat factor) {
// 	(*ptrglGlobalAlphaFactorfSUN)(factor);
// }
// void goglGlobalAlphaFactordSUN(GLdouble factor) {
// 	(*ptrglGlobalAlphaFactordSUN)(factor);
// }
// void goglGlobalAlphaFactorubSUN(GLubyte factor) {
// 	(*ptrglGlobalAlphaFactorubSUN)(factor);
// }
// void goglGlobalAlphaFactorusSUN(GLushort factor) {
// 	(*ptrglGlobalAlphaFactorusSUN)(factor);
// }
// void goglGlobalAlphaFactoruiSUN(GLuint factor) {
// 	(*ptrglGlobalAlphaFactoruiSUN)(factor);
// }
// //  SUN_mesh_array
// void goglDrawMeshArraysSUN(GLenum mode, GLint first, GLsizei count, GLsizei width) {
// 	(*ptrglDrawMeshArraysSUN)(mode, first, count, width);
// }
// //  SUN_slice_accum
// //  SUN_triangle_list
// void goglReplacementCodeuiSUN(GLuint code) {
// 	(*ptrglReplacementCodeuiSUN)(code);
// }
// void goglReplacementCodeusSUN(GLushort code) {
// 	(*ptrglReplacementCodeusSUN)(code);
// }
// void goglReplacementCodeubSUN(GLubyte code) {
// 	(*ptrglReplacementCodeubSUN)(code);
// }
// void goglReplacementCodeuivSUN(GLuint* code) {
// 	(*ptrglReplacementCodeuivSUN)(code);
// }
// void goglReplacementCodeusvSUN(GLushort* code) {
// 	(*ptrglReplacementCodeusvSUN)(code);
// }
// void goglReplacementCodeubvSUN(GLubyte* code) {
// 	(*ptrglReplacementCodeubvSUN)(code);
// }
// void goglReplacementCodePointerSUN(GLenum type_, GLsizei stride, GLvoid** pointer) {
// 	(*ptrglReplacementCodePointerSUN)(type_, stride, pointer);
// }
// //  SUN_vertex
// void goglColor4ubVertex2fSUN(GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y) {
// 	(*ptrglColor4ubVertex2fSUN)(r, g, b, a, x, y);
// }
// void goglColor4ubVertex2fvSUN(GLubyte* c, GLfloat* v) {
// 	(*ptrglColor4ubVertex2fvSUN)(c, v);
// }
// void goglColor4ubVertex3fSUN(GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglColor4ubVertex3fSUN)(r, g, b, a, x, y, z);
// }
// void goglColor4ubVertex3fvSUN(GLubyte* c, GLfloat* v) {
// 	(*ptrglColor4ubVertex3fvSUN)(c, v);
// }
// void goglColor3fVertex3fSUN(GLfloat r, GLfloat g, GLfloat b, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglColor3fVertex3fSUN)(r, g, b, x, y, z);
// }
// void goglColor3fVertex3fvSUN(GLfloat* c, GLfloat* v) {
// 	(*ptrglColor3fVertex3fvSUN)(c, v);
// }
// void goglNormal3fVertex3fSUN(GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglNormal3fVertex3fSUN)(nx, ny, nz, x, y, z);
// }
// void goglNormal3fVertex3fvSUN(GLfloat* n, GLfloat* v) {
// 	(*ptrglNormal3fVertex3fvSUN)(n, v);
// }
// void goglColor4fNormal3fVertex3fSUN(GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglColor4fNormal3fVertex3fSUN)(r, g, b, a, nx, ny, nz, x, y, z);
// }
// void goglColor4fNormal3fVertex3fvSUN(GLfloat* c, GLfloat* n, GLfloat* v) {
// 	(*ptrglColor4fNormal3fVertex3fvSUN)(c, n, v);
// }
// void goglTexCoord2fVertex3fSUN(GLfloat s, GLfloat t, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTexCoord2fVertex3fSUN)(s, t, x, y, z);
// }
// void goglTexCoord2fVertex3fvSUN(GLfloat* tc, GLfloat* v) {
// 	(*ptrglTexCoord2fVertex3fvSUN)(tc, v);
// }
// void goglTexCoord4fVertex4fSUN(GLfloat s, GLfloat t, GLfloat p, GLfloat q, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglTexCoord4fVertex4fSUN)(s, t, p, q, x, y, z, w);
// }
// void goglTexCoord4fVertex4fvSUN(GLfloat* tc, GLfloat* v) {
// 	(*ptrglTexCoord4fVertex4fvSUN)(tc, v);
// }
// void goglTexCoord2fColor4ubVertex3fSUN(GLfloat s, GLfloat t, GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTexCoord2fColor4ubVertex3fSUN)(s, t, r, g, b, a, x, y, z);
// }
// void goglTexCoord2fColor4ubVertex3fvSUN(GLfloat* tc, GLubyte* c, GLfloat* v) {
// 	(*ptrglTexCoord2fColor4ubVertex3fvSUN)(tc, c, v);
// }
// void goglTexCoord2fColor3fVertex3fSUN(GLfloat s, GLfloat t, GLfloat r, GLfloat g, GLfloat b, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTexCoord2fColor3fVertex3fSUN)(s, t, r, g, b, x, y, z);
// }
// void goglTexCoord2fColor3fVertex3fvSUN(GLfloat* tc, GLfloat* c, GLfloat* v) {
// 	(*ptrglTexCoord2fColor3fVertex3fvSUN)(tc, c, v);
// }
// void goglTexCoord2fNormal3fVertex3fSUN(GLfloat s, GLfloat t, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTexCoord2fNormal3fVertex3fSUN)(s, t, nx, ny, nz, x, y, z);
// }
// void goglTexCoord2fNormal3fVertex3fvSUN(GLfloat* tc, GLfloat* n, GLfloat* v) {
// 	(*ptrglTexCoord2fNormal3fVertex3fvSUN)(tc, n, v);
// }
// void goglTexCoord2fColor4fNormal3fVertex3fSUN(GLfloat s, GLfloat t, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTexCoord2fColor4fNormal3fVertex3fSUN)(s, t, r, g, b, a, nx, ny, nz, x, y, z);
// }
// void goglTexCoord2fColor4fNormal3fVertex3fvSUN(GLfloat* tc, GLfloat* c, GLfloat* n, GLfloat* v) {
// 	(*ptrglTexCoord2fColor4fNormal3fVertex3fvSUN)(tc, c, n, v);
// }
// void goglTexCoord4fColor4fNormal3fVertex4fSUN(GLfloat s, GLfloat t, GLfloat p, GLfloat q, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglTexCoord4fColor4fNormal3fVertex4fSUN)(s, t, p, q, r, g, b, a, nx, ny, nz, x, y, z, w);
// }
// void goglTexCoord4fColor4fNormal3fVertex4fvSUN(GLfloat* tc, GLfloat* c, GLfloat* n, GLfloat* v) {
// 	(*ptrglTexCoord4fColor4fNormal3fVertex4fvSUN)(tc, c, n, v);
// }
// void goglReplacementCodeuiVertex3fSUN(GLuint rc, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiVertex3fSUN)(rc, x, y, z);
// }
// void goglReplacementCodeuiVertex3fvSUN(GLuint* rc, GLfloat* v) {
// 	(*ptrglReplacementCodeuiVertex3fvSUN)(rc, v);
// }
// void goglReplacementCodeuiColor4ubVertex3fSUN(GLuint rc, GLubyte r, GLubyte g, GLubyte b, GLubyte a, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiColor4ubVertex3fSUN)(rc, r, g, b, a, x, y, z);
// }
// void goglReplacementCodeuiColor4ubVertex3fvSUN(GLuint* rc, GLubyte* c, GLfloat* v) {
// 	(*ptrglReplacementCodeuiColor4ubVertex3fvSUN)(rc, c, v);
// }
// void goglReplacementCodeuiColor3fVertex3fSUN(GLuint rc, GLfloat r, GLfloat g, GLfloat b, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiColor3fVertex3fSUN)(rc, r, g, b, x, y, z);
// }
// void goglReplacementCodeuiColor3fVertex3fvSUN(GLuint* rc, GLfloat* c, GLfloat* v) {
// 	(*ptrglReplacementCodeuiColor3fVertex3fvSUN)(rc, c, v);
// }
// void goglReplacementCodeuiNormal3fVertex3fSUN(GLuint rc, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiNormal3fVertex3fSUN)(rc, nx, ny, nz, x, y, z);
// }
// void goglReplacementCodeuiNormal3fVertex3fvSUN(GLuint* rc, GLfloat* n, GLfloat* v) {
// 	(*ptrglReplacementCodeuiNormal3fVertex3fvSUN)(rc, n, v);
// }
// void goglReplacementCodeuiColor4fNormal3fVertex3fSUN(GLuint rc, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiColor4fNormal3fVertex3fSUN)(rc, r, g, b, a, nx, ny, nz, x, y, z);
// }
// void goglReplacementCodeuiColor4fNormal3fVertex3fvSUN(GLuint* rc, GLfloat* c, GLfloat* n, GLfloat* v) {
// 	(*ptrglReplacementCodeuiColor4fNormal3fVertex3fvSUN)(rc, c, n, v);
// }
// void goglReplacementCodeuiTexCoord2fVertex3fSUN(GLuint rc, GLfloat s, GLfloat t, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiTexCoord2fVertex3fSUN)(rc, s, t, x, y, z);
// }
// void goglReplacementCodeuiTexCoord2fVertex3fvSUN(GLuint* rc, GLfloat* tc, GLfloat* v) {
// 	(*ptrglReplacementCodeuiTexCoord2fVertex3fvSUN)(rc, tc, v);
// }
// void goglReplacementCodeuiTexCoord2fNormal3fVertex3fSUN(GLuint rc, GLfloat s, GLfloat t, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fSUN)(rc, s, t, nx, ny, nz, x, y, z);
// }
// void goglReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN(GLuint* rc, GLfloat* tc, GLfloat* n, GLfloat* v) {
// 	(*ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN)(rc, tc, n, v);
// }
// void goglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN(GLuint rc, GLfloat s, GLfloat t, GLfloat r, GLfloat g, GLfloat b, GLfloat a, GLfloat nx, GLfloat ny, GLfloat nz, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN)(rc, s, t, r, g, b, a, nx, ny, nz, x, y, z);
// }
// void goglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN(GLuint* rc, GLfloat* tc, GLfloat* c, GLfloat* n, GLfloat* v) {
// 	(*ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN)(rc, tc, c, n, v);
// }
// 
// int init_SUN_convolution_border_modes() {
// 	return 0;
// }
// int init_SUN_global_alpha() {
// 	ptrglGlobalAlphaFactorbSUN = goglGetProcAddress("glGlobalAlphaFactorbSUN");
// 	if(ptrglGlobalAlphaFactorbSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactorsSUN = goglGetProcAddress("glGlobalAlphaFactorsSUN");
// 	if(ptrglGlobalAlphaFactorsSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactoriSUN = goglGetProcAddress("glGlobalAlphaFactoriSUN");
// 	if(ptrglGlobalAlphaFactoriSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactorfSUN = goglGetProcAddress("glGlobalAlphaFactorfSUN");
// 	if(ptrglGlobalAlphaFactorfSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactordSUN = goglGetProcAddress("glGlobalAlphaFactordSUN");
// 	if(ptrglGlobalAlphaFactordSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactorubSUN = goglGetProcAddress("glGlobalAlphaFactorubSUN");
// 	if(ptrglGlobalAlphaFactorubSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactorusSUN = goglGetProcAddress("glGlobalAlphaFactorusSUN");
// 	if(ptrglGlobalAlphaFactorusSUN == NULL) return 1;
// 	ptrglGlobalAlphaFactoruiSUN = goglGetProcAddress("glGlobalAlphaFactoruiSUN");
// 	if(ptrglGlobalAlphaFactoruiSUN == NULL) return 1;
// 	return 0;
// }
// int init_SUN_mesh_array() {
// 	ptrglDrawMeshArraysSUN = goglGetProcAddress("glDrawMeshArraysSUN");
// 	if(ptrglDrawMeshArraysSUN == NULL) return 1;
// 	return 0;
// }
// int init_SUN_slice_accum() {
// 	return 0;
// }
// int init_SUN_triangle_list() {
// 	ptrglReplacementCodeuiSUN = goglGetProcAddress("glReplacementCodeuiSUN");
// 	if(ptrglReplacementCodeuiSUN == NULL) return 1;
// 	ptrglReplacementCodeusSUN = goglGetProcAddress("glReplacementCodeusSUN");
// 	if(ptrglReplacementCodeusSUN == NULL) return 1;
// 	ptrglReplacementCodeubSUN = goglGetProcAddress("glReplacementCodeubSUN");
// 	if(ptrglReplacementCodeubSUN == NULL) return 1;
// 	ptrglReplacementCodeuivSUN = goglGetProcAddress("glReplacementCodeuivSUN");
// 	if(ptrglReplacementCodeuivSUN == NULL) return 1;
// 	ptrglReplacementCodeusvSUN = goglGetProcAddress("glReplacementCodeusvSUN");
// 	if(ptrglReplacementCodeusvSUN == NULL) return 1;
// 	ptrglReplacementCodeubvSUN = goglGetProcAddress("glReplacementCodeubvSUN");
// 	if(ptrglReplacementCodeubvSUN == NULL) return 1;
// 	ptrglReplacementCodePointerSUN = goglGetProcAddress("glReplacementCodePointerSUN");
// 	if(ptrglReplacementCodePointerSUN == NULL) return 1;
// 	return 0;
// }
// int init_SUN_vertex() {
// 	ptrglColor4ubVertex2fSUN = goglGetProcAddress("glColor4ubVertex2fSUN");
// 	if(ptrglColor4ubVertex2fSUN == NULL) return 1;
// 	ptrglColor4ubVertex2fvSUN = goglGetProcAddress("glColor4ubVertex2fvSUN");
// 	if(ptrglColor4ubVertex2fvSUN == NULL) return 1;
// 	ptrglColor4ubVertex3fSUN = goglGetProcAddress("glColor4ubVertex3fSUN");
// 	if(ptrglColor4ubVertex3fSUN == NULL) return 1;
// 	ptrglColor4ubVertex3fvSUN = goglGetProcAddress("glColor4ubVertex3fvSUN");
// 	if(ptrglColor4ubVertex3fvSUN == NULL) return 1;
// 	ptrglColor3fVertex3fSUN = goglGetProcAddress("glColor3fVertex3fSUN");
// 	if(ptrglColor3fVertex3fSUN == NULL) return 1;
// 	ptrglColor3fVertex3fvSUN = goglGetProcAddress("glColor3fVertex3fvSUN");
// 	if(ptrglColor3fVertex3fvSUN == NULL) return 1;
// 	ptrglNormal3fVertex3fSUN = goglGetProcAddress("glNormal3fVertex3fSUN");
// 	if(ptrglNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglNormal3fVertex3fvSUN = goglGetProcAddress("glNormal3fVertex3fvSUN");
// 	if(ptrglNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglColor4fNormal3fVertex3fSUN = goglGetProcAddress("glColor4fNormal3fVertex3fSUN");
// 	if(ptrglColor4fNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglColor4fNormal3fVertex3fvSUN = goglGetProcAddress("glColor4fNormal3fVertex3fvSUN");
// 	if(ptrglColor4fNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglTexCoord2fVertex3fSUN = goglGetProcAddress("glTexCoord2fVertex3fSUN");
// 	if(ptrglTexCoord2fVertex3fSUN == NULL) return 1;
// 	ptrglTexCoord2fVertex3fvSUN = goglGetProcAddress("glTexCoord2fVertex3fvSUN");
// 	if(ptrglTexCoord2fVertex3fvSUN == NULL) return 1;
// 	ptrglTexCoord4fVertex4fSUN = goglGetProcAddress("glTexCoord4fVertex4fSUN");
// 	if(ptrglTexCoord4fVertex4fSUN == NULL) return 1;
// 	ptrglTexCoord4fVertex4fvSUN = goglGetProcAddress("glTexCoord4fVertex4fvSUN");
// 	if(ptrglTexCoord4fVertex4fvSUN == NULL) return 1;
// 	ptrglTexCoord2fColor4ubVertex3fSUN = goglGetProcAddress("glTexCoord2fColor4ubVertex3fSUN");
// 	if(ptrglTexCoord2fColor4ubVertex3fSUN == NULL) return 1;
// 	ptrglTexCoord2fColor4ubVertex3fvSUN = goglGetProcAddress("glTexCoord2fColor4ubVertex3fvSUN");
// 	if(ptrglTexCoord2fColor4ubVertex3fvSUN == NULL) return 1;
// 	ptrglTexCoord2fColor3fVertex3fSUN = goglGetProcAddress("glTexCoord2fColor3fVertex3fSUN");
// 	if(ptrglTexCoord2fColor3fVertex3fSUN == NULL) return 1;
// 	ptrglTexCoord2fColor3fVertex3fvSUN = goglGetProcAddress("glTexCoord2fColor3fVertex3fvSUN");
// 	if(ptrglTexCoord2fColor3fVertex3fvSUN == NULL) return 1;
// 	ptrglTexCoord2fNormal3fVertex3fSUN = goglGetProcAddress("glTexCoord2fNormal3fVertex3fSUN");
// 	if(ptrglTexCoord2fNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglTexCoord2fNormal3fVertex3fvSUN = goglGetProcAddress("glTexCoord2fNormal3fVertex3fvSUN");
// 	if(ptrglTexCoord2fNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglTexCoord2fColor4fNormal3fVertex3fSUN = goglGetProcAddress("glTexCoord2fColor4fNormal3fVertex3fSUN");
// 	if(ptrglTexCoord2fColor4fNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglTexCoord2fColor4fNormal3fVertex3fvSUN = goglGetProcAddress("glTexCoord2fColor4fNormal3fVertex3fvSUN");
// 	if(ptrglTexCoord2fColor4fNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglTexCoord4fColor4fNormal3fVertex4fSUN = goglGetProcAddress("glTexCoord4fColor4fNormal3fVertex4fSUN");
// 	if(ptrglTexCoord4fColor4fNormal3fVertex4fSUN == NULL) return 1;
// 	ptrglTexCoord4fColor4fNormal3fVertex4fvSUN = goglGetProcAddress("glTexCoord4fColor4fNormal3fVertex4fvSUN");
// 	if(ptrglTexCoord4fColor4fNormal3fVertex4fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiVertex3fSUN = goglGetProcAddress("glReplacementCodeuiVertex3fSUN");
// 	if(ptrglReplacementCodeuiVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiVertex3fvSUN");
// 	if(ptrglReplacementCodeuiVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiColor4ubVertex3fSUN = goglGetProcAddress("glReplacementCodeuiColor4ubVertex3fSUN");
// 	if(ptrglReplacementCodeuiColor4ubVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiColor4ubVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiColor4ubVertex3fvSUN");
// 	if(ptrglReplacementCodeuiColor4ubVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiColor3fVertex3fSUN = goglGetProcAddress("glReplacementCodeuiColor3fVertex3fSUN");
// 	if(ptrglReplacementCodeuiColor3fVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiColor3fVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiColor3fVertex3fvSUN");
// 	if(ptrglReplacementCodeuiColor3fVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiNormal3fVertex3fSUN = goglGetProcAddress("glReplacementCodeuiNormal3fVertex3fSUN");
// 	if(ptrglReplacementCodeuiNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiNormal3fVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiNormal3fVertex3fvSUN");
// 	if(ptrglReplacementCodeuiNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiColor4fNormal3fVertex3fSUN = goglGetProcAddress("glReplacementCodeuiColor4fNormal3fVertex3fSUN");
// 	if(ptrglReplacementCodeuiColor4fNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiColor4fNormal3fVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiColor4fNormal3fVertex3fvSUN");
// 	if(ptrglReplacementCodeuiColor4fNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiTexCoord2fVertex3fSUN = goglGetProcAddress("glReplacementCodeuiTexCoord2fVertex3fSUN");
// 	if(ptrglReplacementCodeuiTexCoord2fVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiTexCoord2fVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiTexCoord2fVertex3fvSUN");
// 	if(ptrglReplacementCodeuiTexCoord2fVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fSUN = goglGetProcAddress("glReplacementCodeuiTexCoord2fNormal3fVertex3fSUN");
// 	if(ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN");
// 	if(ptrglReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN == NULL) return 1;
// 	ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN = goglGetProcAddress("glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN");
// 	if(ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN == NULL) return 1;
// 	ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN = goglGetProcAddress("glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN");
// 	if(ptrglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN == NULL) return 1;
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

// SUN_convolution_border_modes
const (
	WRAP_BORDER_SUN = 0x81D4
)
// SUN_global_alpha
const (
	GLOBAL_ALPHA_FACTOR_SUN = 0x81DA
	GLOBAL_ALPHA_SUN = 0x81D9
)
// SUN_mesh_array
const (
	QUAD_MESH_SUN = 0x8614
	TRIANGLE_MESH_SUN = 0x8615
)
// SUN_slice_accum
const (
	SLICE_ACCUM_SUN = 0x85CC
)
// SUN_triangle_list
const (
	R1UI_C3F_V3F_SUN = 0x85C6
	R1UI_C4F_N3F_V3F_SUN = 0x85C8
	R1UI_C4UB_V3F_SUN = 0x85C5
	R1UI_N3F_V3F_SUN = 0x85C7
	R1UI_T2F_C4F_N3F_V3F_SUN = 0x85CB
	R1UI_T2F_N3F_V3F_SUN = 0x85CA
	R1UI_T2F_V3F_SUN = 0x85C9
	R1UI_V3F_SUN = 0x85C4
	REPLACEMENT_CODE_ARRAY_POINTER_SUN = 0x85C3
	REPLACEMENT_CODE_ARRAY_STRIDE_SUN = 0x85C2
	REPLACEMENT_CODE_ARRAY_SUN = 0x85C0
	REPLACEMENT_CODE_ARRAY_TYPE_SUN = 0x85C1
	REPLACEMENT_CODE_SUN = 0x81D8
	REPLACE_MIDDLE_SUN = 0x0002
	REPLACE_OLDEST_SUN = 0x0003
	RESTART_SUN = 0x0001
	TRIANGLE_LIST_SUN = 0x81D7
)
// SUN_vertex
const (
)
// SUN_convolution_border_modes

// SUN_global_alpha

func GlobalAlphaFactorbSUN(factor Byte)  {
	C.goglGlobalAlphaFactorbSUN((C.GLbyte)(factor))
}
func GlobalAlphaFactorsSUN(factor Short)  {
	C.goglGlobalAlphaFactorsSUN((C.GLshort)(factor))
}
func GlobalAlphaFactoriSUN(factor Int)  {
	C.goglGlobalAlphaFactoriSUN((C.GLint)(factor))
}
func GlobalAlphaFactorfSUN(factor Float)  {
	C.goglGlobalAlphaFactorfSUN((C.GLfloat)(factor))
}
func GlobalAlphaFactordSUN(factor Double)  {
	C.goglGlobalAlphaFactordSUN((C.GLdouble)(factor))
}
func GlobalAlphaFactorubSUN(factor Ubyte)  {
	C.goglGlobalAlphaFactorubSUN((C.GLubyte)(factor))
}
func GlobalAlphaFactorusSUN(factor Ushort)  {
	C.goglGlobalAlphaFactorusSUN((C.GLushort)(factor))
}
func GlobalAlphaFactoruiSUN(factor Uint)  {
	C.goglGlobalAlphaFactoruiSUN((C.GLuint)(factor))
}
// SUN_mesh_array

func DrawMeshArraysSUN(mode Enum, first Int, count Sizei, width Sizei)  {
	C.goglDrawMeshArraysSUN((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(width))
}
// SUN_slice_accum

// SUN_triangle_list

func ReplacementCodeuiSUN(code Uint)  {
	C.goglReplacementCodeuiSUN((C.GLuint)(code))
}
func ReplacementCodeusSUN(code Ushort)  {
	C.goglReplacementCodeusSUN((C.GLushort)(code))
}
func ReplacementCodeubSUN(code Ubyte)  {
	C.goglReplacementCodeubSUN((C.GLubyte)(code))
}
func ReplacementCodeuivSUN(code *Uint)  {
	C.goglReplacementCodeuivSUN((*C.GLuint)(code))
}
func ReplacementCodeusvSUN(code *Ushort)  {
	C.goglReplacementCodeusvSUN((*C.GLushort)(code))
}
func ReplacementCodeubvSUN(code *Ubyte)  {
	C.goglReplacementCodeubvSUN((*C.GLubyte)(code))
}
func ReplacementCodePointerSUN(type_ Enum, stride Sizei, pointer *Pointer)  {
	C.goglReplacementCodePointerSUN((C.GLenum)(type_), (C.GLsizei)(stride), (*unsafe.Pointer)(pointer))
}
// SUN_vertex

func Color4ubVertex2fSUN(r Ubyte, g Ubyte, b Ubyte, a Ubyte, x Float, y Float)  {
	C.goglColor4ubVertex2fSUN((C.GLubyte)(r), (C.GLubyte)(g), (C.GLubyte)(b), (C.GLubyte)(a), (C.GLfloat)(x), (C.GLfloat)(y))
}
func Color4ubVertex2fvSUN(c *Ubyte, v *Float)  {
	C.goglColor4ubVertex2fvSUN((*C.GLubyte)(c), (*C.GLfloat)(v))
}
func Color4ubVertex3fSUN(r Ubyte, g Ubyte, b Ubyte, a Ubyte, x Float, y Float, z Float)  {
	C.goglColor4ubVertex3fSUN((C.GLubyte)(r), (C.GLubyte)(g), (C.GLubyte)(b), (C.GLubyte)(a), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Color4ubVertex3fvSUN(c *Ubyte, v *Float)  {
	C.goglColor4ubVertex3fvSUN((*C.GLubyte)(c), (*C.GLfloat)(v))
}
func Color3fVertex3fSUN(r Float, g Float, b Float, x Float, y Float, z Float)  {
	C.goglColor3fVertex3fSUN((C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Color3fVertex3fvSUN(c *Float, v *Float)  {
	C.goglColor3fVertex3fvSUN((*C.GLfloat)(c), (*C.GLfloat)(v))
}
func Normal3fVertex3fSUN(nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglNormal3fVertex3fSUN((C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Normal3fVertex3fvSUN(n *Float, v *Float)  {
	C.goglNormal3fVertex3fvSUN((*C.GLfloat)(n), (*C.GLfloat)(v))
}
func Color4fNormal3fVertex3fSUN(r Float, g Float, b Float, a Float, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglColor4fNormal3fVertex3fSUN((C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(a), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func Color4fNormal3fVertex3fvSUN(c *Float, n *Float, v *Float)  {
	C.goglColor4fNormal3fVertex3fvSUN((*C.GLfloat)(c), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func TexCoord2fVertex3fSUN(s Float, t Float, x Float, y Float, z Float)  {
	C.goglTexCoord2fVertex3fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func TexCoord2fVertex3fvSUN(tc *Float, v *Float)  {
	C.goglTexCoord2fVertex3fvSUN((*C.GLfloat)(tc), (*C.GLfloat)(v))
}
func TexCoord4fVertex4fSUN(s Float, t Float, p Float, q Float, x Float, y Float, z Float, w Float)  {
	C.goglTexCoord4fVertex4fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(p), (C.GLfloat)(q), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func TexCoord4fVertex4fvSUN(tc *Float, v *Float)  {
	C.goglTexCoord4fVertex4fvSUN((*C.GLfloat)(tc), (*C.GLfloat)(v))
}
func TexCoord2fColor4ubVertex3fSUN(s Float, t Float, r Ubyte, g Ubyte, b Ubyte, a Ubyte, x Float, y Float, z Float)  {
	C.goglTexCoord2fColor4ubVertex3fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLubyte)(r), (C.GLubyte)(g), (C.GLubyte)(b), (C.GLubyte)(a), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func TexCoord2fColor4ubVertex3fvSUN(tc *Float, c *Ubyte, v *Float)  {
	C.goglTexCoord2fColor4ubVertex3fvSUN((*C.GLfloat)(tc), (*C.GLubyte)(c), (*C.GLfloat)(v))
}
func TexCoord2fColor3fVertex3fSUN(s Float, t Float, r Float, g Float, b Float, x Float, y Float, z Float)  {
	C.goglTexCoord2fColor3fVertex3fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func TexCoord2fColor3fVertex3fvSUN(tc *Float, c *Float, v *Float)  {
	C.goglTexCoord2fColor3fVertex3fvSUN((*C.GLfloat)(tc), (*C.GLfloat)(c), (*C.GLfloat)(v))
}
func TexCoord2fNormal3fVertex3fSUN(s Float, t Float, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglTexCoord2fNormal3fVertex3fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func TexCoord2fNormal3fVertex3fvSUN(tc *Float, n *Float, v *Float)  {
	C.goglTexCoord2fNormal3fVertex3fvSUN((*C.GLfloat)(tc), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func TexCoord2fColor4fNormal3fVertex3fSUN(s Float, t Float, r Float, g Float, b Float, a Float, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglTexCoord2fColor4fNormal3fVertex3fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(a), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func TexCoord2fColor4fNormal3fVertex3fvSUN(tc *Float, c *Float, n *Float, v *Float)  {
	C.goglTexCoord2fColor4fNormal3fVertex3fvSUN((*C.GLfloat)(tc), (*C.GLfloat)(c), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func TexCoord4fColor4fNormal3fVertex4fSUN(s Float, t Float, p Float, q Float, r Float, g Float, b Float, a Float, nx Float, ny Float, nz Float, x Float, y Float, z Float, w Float)  {
	C.goglTexCoord4fColor4fNormal3fVertex4fSUN((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(p), (C.GLfloat)(q), (C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(a), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func TexCoord4fColor4fNormal3fVertex4fvSUN(tc *Float, c *Float, n *Float, v *Float)  {
	C.goglTexCoord4fColor4fNormal3fVertex4fvSUN((*C.GLfloat)(tc), (*C.GLfloat)(c), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func ReplacementCodeuiVertex3fSUN(rc Uint, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiVertex3fvSUN(rc *Uint, v *Float)  {
	C.goglReplacementCodeuiVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(v))
}
func ReplacementCodeuiColor4ubVertex3fSUN(rc Uint, r Ubyte, g Ubyte, b Ubyte, a Ubyte, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiColor4ubVertex3fSUN((C.GLuint)(rc), (C.GLubyte)(r), (C.GLubyte)(g), (C.GLubyte)(b), (C.GLubyte)(a), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiColor4ubVertex3fvSUN(rc *Uint, c *Ubyte, v *Float)  {
	C.goglReplacementCodeuiColor4ubVertex3fvSUN((*C.GLuint)(rc), (*C.GLubyte)(c), (*C.GLfloat)(v))
}
func ReplacementCodeuiColor3fVertex3fSUN(rc Uint, r Float, g Float, b Float, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiColor3fVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiColor3fVertex3fvSUN(rc *Uint, c *Float, v *Float)  {
	C.goglReplacementCodeuiColor3fVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(c), (*C.GLfloat)(v))
}
func ReplacementCodeuiNormal3fVertex3fSUN(rc Uint, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiNormal3fVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiNormal3fVertex3fvSUN(rc *Uint, n *Float, v *Float)  {
	C.goglReplacementCodeuiNormal3fVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func ReplacementCodeuiColor4fNormal3fVertex3fSUN(rc Uint, r Float, g Float, b Float, a Float, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiColor4fNormal3fVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(a), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiColor4fNormal3fVertex3fvSUN(rc *Uint, c *Float, n *Float, v *Float)  {
	C.goglReplacementCodeuiColor4fNormal3fVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(c), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func ReplacementCodeuiTexCoord2fVertex3fSUN(rc Uint, s Float, t Float, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiTexCoord2fVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiTexCoord2fVertex3fvSUN(rc *Uint, tc *Float, v *Float)  {
	C.goglReplacementCodeuiTexCoord2fVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(tc), (*C.GLfloat)(v))
}
func ReplacementCodeuiTexCoord2fNormal3fVertex3fSUN(rc Uint, s Float, t Float, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiTexCoord2fNormal3fVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN(rc *Uint, tc *Float, n *Float, v *Float)  {
	C.goglReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(tc), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN(rc Uint, s Float, t Float, r Float, g Float, b Float, a Float, nx Float, ny Float, nz Float, x Float, y Float, z Float)  {
	C.goglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN((C.GLuint)(rc), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(g), (C.GLfloat)(b), (C.GLfloat)(a), (C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN(rc *Uint, tc *Float, c *Float, n *Float, v *Float)  {
	C.goglReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN((*C.GLuint)(rc), (*C.GLfloat)(tc), (*C.GLfloat)(c), (*C.GLfloat)(n), (*C.GLfloat)(v))
}
func InitSunConvolutionBorderModes() error {
	var ret C.int
	if ret = C.init_SUN_convolution_border_modes(); ret != 0 {
		return errors.New("unable to initialize SUN_convolution_border_modes")
	}
	return nil
}
func InitSunGlobalAlpha() error {
	var ret C.int
	if ret = C.init_SUN_global_alpha(); ret != 0 {
		return errors.New("unable to initialize SUN_global_alpha")
	}
	return nil
}
func InitSunMeshArray() error {
	var ret C.int
	if ret = C.init_SUN_mesh_array(); ret != 0 {
		return errors.New("unable to initialize SUN_mesh_array")
	}
	return nil
}
func InitSunSliceAccum() error {
	var ret C.int
	if ret = C.init_SUN_slice_accum(); ret != 0 {
		return errors.New("unable to initialize SUN_slice_accum")
	}
	return nil
}
func InitSunTriangleList() error {
	var ret C.int
	if ret = C.init_SUN_triangle_list(); ret != 0 {
		return errors.New("unable to initialize SUN_triangle_list")
	}
	return nil
}
func InitSunVertex() error {
	var ret C.int
	if ret = C.init_SUN_vertex(); ret != 0 {
		return errors.New("unable to initialize SUN_vertex")
	}
	return nil
}
// EOF