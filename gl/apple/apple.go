// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// APPLE_aux_depth_stencil: http://www.opengl.org/registry/specs/APPLE/aux_depth_stencil.txt
// 
// APPLE_client_storage: http://www.opengl.org/registry/specs/APPLE/client_storage.txt
// 
// APPLE_element_array: http://www.opengl.org/registry/specs/APPLE/element_array.txt
// 
// APPLE_fence: http://www.opengl.org/registry/specs/APPLE/fence.txt
// 
// APPLE_float_pixels: http://www.opengl.org/registry/specs/APPLE/float_pixels.txt
// 
// APPLE_flush_buffer_range: http://www.opengl.org/registry/specs/APPLE/flush_buffer_range.txt
// 
// APPLE_object_purgeable: http://www.opengl.org/registry/specs/APPLE/object_purgeable.txt
// 
// APPLE_rgb_422: http://www.opengl.org/registry/specs/APPLE/rgb_422.txt
// 
// APPLE_row_bytes: http://www.opengl.org/registry/specs/APPLE/row_bytes.txt
// 
// APPLE_specular_vector: http://www.opengl.org/registry/specs/APPLE/specular_vector.txt
// 
// APPLE_texture_range: http://www.opengl.org/registry/specs/APPLE/texture_range.txt
// 
// APPLE_transform_hint: http://www.opengl.org/registry/specs/APPLE/transform_hint.txt
// 
// APPLE_vertex_array_object: http://www.opengl.org/registry/specs/APPLE/vertex_array_object.txt
// 
// APPLE_vertex_array_range: http://www.opengl.org/registry/specs/APPLE/vertex_array_range.txt
// 
// APPLE_vertex_program_evaluators: http://www.opengl.org/registry/specs/APPLE/vertex_program_evaluators.txt
// 
// APPLE_ycbcr_422: http://www.opengl.org/registry/specs/APPLE/ycbcr_422.txt
// 
package apple

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
// //  APPLE_aux_depth_stencil
// //  APPLE_client_storage
// //  APPLE_element_array
// void (APIENTRYP ptrglElementPointerAPPLE)(GLenum type, GLvoid* pointer);
// void (APIENTRYP ptrglDrawElementArrayAPPLE)(GLenum mode, GLint first, GLsizei count);
// void (APIENTRYP ptrglDrawRangeElementArrayAPPLE)(GLenum mode, GLuint start, GLuint end, GLint first, GLsizei count);
// void (APIENTRYP ptrglMultiDrawElementArrayAPPLE)(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount);
// void (APIENTRYP ptrglMultiDrawRangeElementArrayAPPLE)(GLenum mode, GLuint start, GLuint end, GLint* first, GLsizei* count, GLsizei primcount);
// //  APPLE_fence
// void (APIENTRYP ptrglGenFencesAPPLE)(GLsizei n, GLuint* fences);
// void (APIENTRYP ptrglDeleteFencesAPPLE)(GLsizei n, GLuint* fences);
// void (APIENTRYP ptrglSetFenceAPPLE)(GLuint fence);
// GLboolean (APIENTRYP ptrglIsFenceAPPLE)(GLuint fence);
// GLboolean (APIENTRYP ptrglTestFenceAPPLE)(GLuint fence);
// void (APIENTRYP ptrglFinishFenceAPPLE)(GLuint fence);
// GLboolean (APIENTRYP ptrglTestObjectAPPLE)(GLenum object, GLuint name);
// void (APIENTRYP ptrglFinishObjectAPPLE)(GLenum object, GLint name);
// //  APPLE_float_pixels
// //  APPLE_flush_buffer_range
// void (APIENTRYP ptrglBufferParameteriAPPLE)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglFlushMappedBufferRangeAPPLE)(GLenum target, GLintptr offset, GLsizeiptr size);
// //  APPLE_object_purgeable
// GLenum (APIENTRYP ptrglObjectPurgeableAPPLE)(GLenum objectType, GLuint name, GLenum option);
// GLenum (APIENTRYP ptrglObjectUnpurgeableAPPLE)(GLenum objectType, GLuint name, GLenum option);
// void (APIENTRYP ptrglGetObjectParameterivAPPLE)(GLenum objectType, GLuint name, GLenum pname, GLint* params);
// //  APPLE_rgb_422
// //  APPLE_row_bytes
// //  APPLE_specular_vector
// //  APPLE_texture_range
// void (APIENTRYP ptrglTextureRangeAPPLE)(GLenum target, GLsizei length, GLvoid* pointer);
// void (APIENTRYP ptrglGetTexParameterPointervAPPLE)(GLenum target, GLenum pname, GLvoid** params);
// //  APPLE_transform_hint
// //  APPLE_vertex_array_object
// void (APIENTRYP ptrglBindVertexArrayAPPLE)(GLuint array);
// void (APIENTRYP ptrglDeleteVertexArraysAPPLE)(GLsizei n, GLuint* arrays);
// void (APIENTRYP ptrglGenVertexArraysAPPLE)(GLsizei n, GLuint* arrays);
// GLboolean (APIENTRYP ptrglIsVertexArrayAPPLE)(GLuint array);
// //  APPLE_vertex_array_range
// void (APIENTRYP ptrglVertexArrayRangeAPPLE)(GLsizei length, GLvoid* pointer);
// void (APIENTRYP ptrglFlushVertexArrayRangeAPPLE)(GLsizei length, GLvoid* pointer);
// void (APIENTRYP ptrglVertexArrayParameteriAPPLE)(GLenum pname, GLint param);
// //  APPLE_vertex_program_evaluators
// void (APIENTRYP ptrglEnableVertexAttribAPPLE)(GLuint index, GLenum pname);
// void (APIENTRYP ptrglDisableVertexAttribAPPLE)(GLuint index, GLenum pname);
// GLboolean (APIENTRYP ptrglIsVertexAttribEnabledAPPLE)(GLuint index, GLenum pname);
// void (APIENTRYP ptrglMapVertexAttrib1dAPPLE)(GLuint index, GLuint size, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points);
// void (APIENTRYP ptrglMapVertexAttrib1fAPPLE)(GLuint index, GLuint size, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points);
// void (APIENTRYP ptrglMapVertexAttrib2dAPPLE)(GLuint index, GLuint size, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points);
// void (APIENTRYP ptrglMapVertexAttrib2fAPPLE)(GLuint index, GLuint size, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points);
// //  APPLE_ycbcr_422
// 
// //  APPLE_aux_depth_stencil
// //  APPLE_client_storage
// //  APPLE_element_array
// void goglElementPointerAPPLE(GLenum type_, GLvoid* pointer) {
// 	(*ptrglElementPointerAPPLE)(type_, pointer);
// }
// void goglDrawElementArrayAPPLE(GLenum mode, GLint first, GLsizei count) {
// 	(*ptrglDrawElementArrayAPPLE)(mode, first, count);
// }
// void goglDrawRangeElementArrayAPPLE(GLenum mode, GLuint start, GLuint end, GLint first, GLsizei count) {
// 	(*ptrglDrawRangeElementArrayAPPLE)(mode, start, end, first, count);
// }
// void goglMultiDrawElementArrayAPPLE(GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
// 	(*ptrglMultiDrawElementArrayAPPLE)(mode, first, count, primcount);
// }
// void goglMultiDrawRangeElementArrayAPPLE(GLenum mode, GLuint start, GLuint end, GLint* first, GLsizei* count, GLsizei primcount) {
// 	(*ptrglMultiDrawRangeElementArrayAPPLE)(mode, start, end, first, count, primcount);
// }
// //  APPLE_fence
// void goglGenFencesAPPLE(GLsizei n, GLuint* fences) {
// 	(*ptrglGenFencesAPPLE)(n, fences);
// }
// void goglDeleteFencesAPPLE(GLsizei n, GLuint* fences) {
// 	(*ptrglDeleteFencesAPPLE)(n, fences);
// }
// void goglSetFenceAPPLE(GLuint fence) {
// 	(*ptrglSetFenceAPPLE)(fence);
// }
// GLboolean goglIsFenceAPPLE(GLuint fence) {
// 	return (*ptrglIsFenceAPPLE)(fence);
// }
// GLboolean goglTestFenceAPPLE(GLuint fence) {
// 	return (*ptrglTestFenceAPPLE)(fence);
// }
// void goglFinishFenceAPPLE(GLuint fence) {
// 	(*ptrglFinishFenceAPPLE)(fence);
// }
// GLboolean goglTestObjectAPPLE(GLenum object, GLuint name) {
// 	return (*ptrglTestObjectAPPLE)(object, name);
// }
// void goglFinishObjectAPPLE(GLenum object, GLint name) {
// 	(*ptrglFinishObjectAPPLE)(object, name);
// }
// //  APPLE_float_pixels
// //  APPLE_flush_buffer_range
// void goglBufferParameteriAPPLE(GLenum target, GLenum pname, GLint param) {
// 	(*ptrglBufferParameteriAPPLE)(target, pname, param);
// }
// void goglFlushMappedBufferRangeAPPLE(GLenum target, GLintptr offset, GLsizeiptr size) {
// 	(*ptrglFlushMappedBufferRangeAPPLE)(target, offset, size);
// }
// //  APPLE_object_purgeable
// GLenum goglObjectPurgeableAPPLE(GLenum objectType, GLuint name, GLenum option) {
// 	return (*ptrglObjectPurgeableAPPLE)(objectType, name, option);
// }
// GLenum goglObjectUnpurgeableAPPLE(GLenum objectType, GLuint name, GLenum option) {
// 	return (*ptrglObjectUnpurgeableAPPLE)(objectType, name, option);
// }
// void goglGetObjectParameterivAPPLE(GLenum objectType, GLuint name, GLenum pname, GLint* params) {
// 	(*ptrglGetObjectParameterivAPPLE)(objectType, name, pname, params);
// }
// //  APPLE_rgb_422
// //  APPLE_row_bytes
// //  APPLE_specular_vector
// //  APPLE_texture_range
// void goglTextureRangeAPPLE(GLenum target, GLsizei length, GLvoid* pointer) {
// 	(*ptrglTextureRangeAPPLE)(target, length, pointer);
// }
// void goglGetTexParameterPointervAPPLE(GLenum target, GLenum pname, GLvoid** params) {
// 	(*ptrglGetTexParameterPointervAPPLE)(target, pname, params);
// }
// //  APPLE_transform_hint
// //  APPLE_vertex_array_object
// void goglBindVertexArrayAPPLE(GLuint array) {
// 	(*ptrglBindVertexArrayAPPLE)(array);
// }
// void goglDeleteVertexArraysAPPLE(GLsizei n, GLuint* arrays) {
// 	(*ptrglDeleteVertexArraysAPPLE)(n, arrays);
// }
// void goglGenVertexArraysAPPLE(GLsizei n, GLuint* arrays) {
// 	(*ptrglGenVertexArraysAPPLE)(n, arrays);
// }
// GLboolean goglIsVertexArrayAPPLE(GLuint array) {
// 	return (*ptrglIsVertexArrayAPPLE)(array);
// }
// //  APPLE_vertex_array_range
// void goglVertexArrayRangeAPPLE(GLsizei length, GLvoid* pointer) {
// 	(*ptrglVertexArrayRangeAPPLE)(length, pointer);
// }
// void goglFlushVertexArrayRangeAPPLE(GLsizei length, GLvoid* pointer) {
// 	(*ptrglFlushVertexArrayRangeAPPLE)(length, pointer);
// }
// void goglVertexArrayParameteriAPPLE(GLenum pname, GLint param) {
// 	(*ptrglVertexArrayParameteriAPPLE)(pname, param);
// }
// //  APPLE_vertex_program_evaluators
// void goglEnableVertexAttribAPPLE(GLuint index, GLenum pname) {
// 	(*ptrglEnableVertexAttribAPPLE)(index, pname);
// }
// void goglDisableVertexAttribAPPLE(GLuint index, GLenum pname) {
// 	(*ptrglDisableVertexAttribAPPLE)(index, pname);
// }
// GLboolean goglIsVertexAttribEnabledAPPLE(GLuint index, GLenum pname) {
// 	return (*ptrglIsVertexAttribEnabledAPPLE)(index, pname);
// }
// void goglMapVertexAttrib1dAPPLE(GLuint index, GLuint size, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
// 	(*ptrglMapVertexAttrib1dAPPLE)(index, size, u1, u2, stride, order, points);
// }
// void goglMapVertexAttrib1fAPPLE(GLuint index, GLuint size, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
// 	(*ptrglMapVertexAttrib1fAPPLE)(index, size, u1, u2, stride, order, points);
// }
// void goglMapVertexAttrib2dAPPLE(GLuint index, GLuint size, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
// 	(*ptrglMapVertexAttrib2dAPPLE)(index, size, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMapVertexAttrib2fAPPLE(GLuint index, GLuint size, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
// 	(*ptrglMapVertexAttrib2fAPPLE)(index, size, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// //  APPLE_ycbcr_422
// 
// int init_APPLE_aux_depth_stencil() {
// 	return 0;
// }
// int init_APPLE_client_storage() {
// 	return 0;
// }
// int init_APPLE_element_array() {
// 	ptrglElementPointerAPPLE = goglGetProcAddress("glElementPointerAPPLE");
// 	if(ptrglElementPointerAPPLE == NULL) return 1;
// 	ptrglDrawElementArrayAPPLE = goglGetProcAddress("glDrawElementArrayAPPLE");
// 	if(ptrglDrawElementArrayAPPLE == NULL) return 1;
// 	ptrglDrawRangeElementArrayAPPLE = goglGetProcAddress("glDrawRangeElementArrayAPPLE");
// 	if(ptrglDrawRangeElementArrayAPPLE == NULL) return 1;
// 	ptrglMultiDrawElementArrayAPPLE = goglGetProcAddress("glMultiDrawElementArrayAPPLE");
// 	if(ptrglMultiDrawElementArrayAPPLE == NULL) return 1;
// 	ptrglMultiDrawRangeElementArrayAPPLE = goglGetProcAddress("glMultiDrawRangeElementArrayAPPLE");
// 	if(ptrglMultiDrawRangeElementArrayAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_fence() {
// 	ptrglGenFencesAPPLE = goglGetProcAddress("glGenFencesAPPLE");
// 	if(ptrglGenFencesAPPLE == NULL) return 1;
// 	ptrglDeleteFencesAPPLE = goglGetProcAddress("glDeleteFencesAPPLE");
// 	if(ptrglDeleteFencesAPPLE == NULL) return 1;
// 	ptrglSetFenceAPPLE = goglGetProcAddress("glSetFenceAPPLE");
// 	if(ptrglSetFenceAPPLE == NULL) return 1;
// 	ptrglIsFenceAPPLE = goglGetProcAddress("glIsFenceAPPLE");
// 	if(ptrglIsFenceAPPLE == NULL) return 1;
// 	ptrglTestFenceAPPLE = goglGetProcAddress("glTestFenceAPPLE");
// 	if(ptrglTestFenceAPPLE == NULL) return 1;
// 	ptrglFinishFenceAPPLE = goglGetProcAddress("glFinishFenceAPPLE");
// 	if(ptrglFinishFenceAPPLE == NULL) return 1;
// 	ptrglTestObjectAPPLE = goglGetProcAddress("glTestObjectAPPLE");
// 	if(ptrglTestObjectAPPLE == NULL) return 1;
// 	ptrglFinishObjectAPPLE = goglGetProcAddress("glFinishObjectAPPLE");
// 	if(ptrglFinishObjectAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_float_pixels() {
// 	return 0;
// }
// int init_APPLE_flush_buffer_range() {
// 	ptrglBufferParameteriAPPLE = goglGetProcAddress("glBufferParameteriAPPLE");
// 	if(ptrglBufferParameteriAPPLE == NULL) return 1;
// 	ptrglFlushMappedBufferRangeAPPLE = goglGetProcAddress("glFlushMappedBufferRangeAPPLE");
// 	if(ptrglFlushMappedBufferRangeAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_object_purgeable() {
// 	ptrglObjectPurgeableAPPLE = goglGetProcAddress("glObjectPurgeableAPPLE");
// 	if(ptrglObjectPurgeableAPPLE == NULL) return 1;
// 	ptrglObjectUnpurgeableAPPLE = goglGetProcAddress("glObjectUnpurgeableAPPLE");
// 	if(ptrglObjectUnpurgeableAPPLE == NULL) return 1;
// 	ptrglGetObjectParameterivAPPLE = goglGetProcAddress("glGetObjectParameterivAPPLE");
// 	if(ptrglGetObjectParameterivAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_rgb_422() {
// 	return 0;
// }
// int init_APPLE_row_bytes() {
// 	return 0;
// }
// int init_APPLE_specular_vector() {
// 	return 0;
// }
// int init_APPLE_texture_range() {
// 	ptrglTextureRangeAPPLE = goglGetProcAddress("glTextureRangeAPPLE");
// 	if(ptrglTextureRangeAPPLE == NULL) return 1;
// 	ptrglGetTexParameterPointervAPPLE = goglGetProcAddress("glGetTexParameterPointervAPPLE");
// 	if(ptrglGetTexParameterPointervAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_transform_hint() {
// 	return 0;
// }
// int init_APPLE_vertex_array_object() {
// 	ptrglBindVertexArrayAPPLE = goglGetProcAddress("glBindVertexArrayAPPLE");
// 	if(ptrglBindVertexArrayAPPLE == NULL) return 1;
// 	ptrglDeleteVertexArraysAPPLE = goglGetProcAddress("glDeleteVertexArraysAPPLE");
// 	if(ptrglDeleteVertexArraysAPPLE == NULL) return 1;
// 	ptrglGenVertexArraysAPPLE = goglGetProcAddress("glGenVertexArraysAPPLE");
// 	if(ptrglGenVertexArraysAPPLE == NULL) return 1;
// 	ptrglIsVertexArrayAPPLE = goglGetProcAddress("glIsVertexArrayAPPLE");
// 	if(ptrglIsVertexArrayAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_vertex_array_range() {
// 	ptrglVertexArrayRangeAPPLE = goglGetProcAddress("glVertexArrayRangeAPPLE");
// 	if(ptrglVertexArrayRangeAPPLE == NULL) return 1;
// 	ptrglFlushVertexArrayRangeAPPLE = goglGetProcAddress("glFlushVertexArrayRangeAPPLE");
// 	if(ptrglFlushVertexArrayRangeAPPLE == NULL) return 1;
// 	ptrglVertexArrayParameteriAPPLE = goglGetProcAddress("glVertexArrayParameteriAPPLE");
// 	if(ptrglVertexArrayParameteriAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_vertex_program_evaluators() {
// 	ptrglEnableVertexAttribAPPLE = goglGetProcAddress("glEnableVertexAttribAPPLE");
// 	if(ptrglEnableVertexAttribAPPLE == NULL) return 1;
// 	ptrglDisableVertexAttribAPPLE = goglGetProcAddress("glDisableVertexAttribAPPLE");
// 	if(ptrglDisableVertexAttribAPPLE == NULL) return 1;
// 	ptrglIsVertexAttribEnabledAPPLE = goglGetProcAddress("glIsVertexAttribEnabledAPPLE");
// 	if(ptrglIsVertexAttribEnabledAPPLE == NULL) return 1;
// 	ptrglMapVertexAttrib1dAPPLE = goglGetProcAddress("glMapVertexAttrib1dAPPLE");
// 	if(ptrglMapVertexAttrib1dAPPLE == NULL) return 1;
// 	ptrglMapVertexAttrib1fAPPLE = goglGetProcAddress("glMapVertexAttrib1fAPPLE");
// 	if(ptrglMapVertexAttrib1fAPPLE == NULL) return 1;
// 	ptrglMapVertexAttrib2dAPPLE = goglGetProcAddress("glMapVertexAttrib2dAPPLE");
// 	if(ptrglMapVertexAttrib2dAPPLE == NULL) return 1;
// 	ptrglMapVertexAttrib2fAPPLE = goglGetProcAddress("glMapVertexAttrib2fAPPLE");
// 	if(ptrglMapVertexAttrib2fAPPLE == NULL) return 1;
// 	return 0;
// }
// int init_APPLE_ycbcr_422() {
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

// APPLE_aux_depth_stencil
const (
	AUX_DEPTH_STENCIL_APPLE = 0x8A14
)
// APPLE_client_storage
const (
	UNPACK_CLIENT_STORAGE_APPLE = 0x85B2
)
// APPLE_element_array
const (
	ELEMENT_ARRAY_APPLE = 0x8A0C
	ELEMENT_ARRAY_POINTER_APPLE = 0x8A0E
	ELEMENT_ARRAY_TYPE_APPLE = 0x8A0D
)
// APPLE_fence
const (
	DRAW_PIXELS_APPLE = 0x8A0A
	FENCE_APPLE = 0x8A0B
)
// APPLE_float_pixels
const (
	ALPHA_FLOAT16_APPLE = 0x881C
	ALPHA_FLOAT32_APPLE = 0x8816
	COLOR_FLOAT_APPLE = 0x8A0F
	HALF_APPLE = 0x140B
	INTENSITY_FLOAT16_APPLE = 0x881D
	INTENSITY_FLOAT32_APPLE = 0x8817
	LUMINANCE_ALPHA_FLOAT16_APPLE = 0x881F
	LUMINANCE_ALPHA_FLOAT32_APPLE = 0x8819
	LUMINANCE_FLOAT16_APPLE = 0x881E
	LUMINANCE_FLOAT32_APPLE = 0x8818
	RGBA_FLOAT16_APPLE = 0x881A
	RGBA_FLOAT32_APPLE = 0x8814
	RGB_FLOAT16_APPLE = 0x881B
	RGB_FLOAT32_APPLE = 0x8815
)
// APPLE_flush_buffer_range
const (
	BUFFER_FLUSHING_UNMAP_APPLE = 0x8A13
	BUFFER_SERIALIZED_MODIFY_APPLE = 0x8A12
)
// APPLE_object_purgeable
const (
	BUFFER_OBJECT_APPLE = 0x85B3
	PURGEABLE_APPLE = 0x8A1D
	RELEASED_APPLE = 0x8A19
	RETAINED_APPLE = 0x8A1B
	UNDEFINED_APPLE = 0x8A1C
	VOLATILE_APPLE = 0x8A1A
)
// APPLE_rgb_422
const (
	RGB_422_APPLE = 0x8A1F
)
// APPLE_row_bytes
const (
	PACK_ROW_BYTES_APPLE = 0x8A15
	UNPACK_ROW_BYTES_APPLE = 0x8A16
)
// APPLE_specular_vector
const (
	LIGHT_MODEL_SPECULAR_VECTOR_APPLE = 0x85B0
)
// APPLE_texture_range
const (
	STORAGE_PRIVATE_APPLE = 0x85BD
	TEXTURE_RANGE_LENGTH_APPLE = 0x85B7
	TEXTURE_RANGE_POINTER_APPLE = 0x85B8
	TEXTURE_STORAGE_HINT_APPLE = 0x85BC
)
// APPLE_transform_hint
const (
	TRANSFORM_HINT_APPLE = 0x85B1
)
// APPLE_vertex_array_object
const (
	VERTEX_ARRAY_BINDING_APPLE = 0x85B5
)
// APPLE_vertex_array_range
const (
	STORAGE_CLIENT_APPLE = 0x85B4
	VERTEX_ARRAY_RANGE_APPLE = 0x851D
	VERTEX_ARRAY_RANGE_LENGTH_APPLE = 0x851E
	VERTEX_ARRAY_RANGE_POINTER_APPLE = 0x8521
	VERTEX_ARRAY_STORAGE_HINT_APPLE = 0x851F
)
// APPLE_vertex_program_evaluators
const (
	VERTEX_ATTRIB_MAP1_APPLE = 0x8A00
	VERTEX_ATTRIB_MAP1_COEFF_APPLE = 0x8A03
	VERTEX_ATTRIB_MAP1_DOMAIN_APPLE = 0x8A05
	VERTEX_ATTRIB_MAP1_ORDER_APPLE = 0x8A04
	VERTEX_ATTRIB_MAP1_SIZE_APPLE = 0x8A02
	VERTEX_ATTRIB_MAP2_APPLE = 0x8A01
	VERTEX_ATTRIB_MAP2_COEFF_APPLE = 0x8A07
	VERTEX_ATTRIB_MAP2_DOMAIN_APPLE = 0x8A09
	VERTEX_ATTRIB_MAP2_ORDER_APPLE = 0x8A08
	VERTEX_ATTRIB_MAP2_SIZE_APPLE = 0x8A06
)
// APPLE_ycbcr_422
const (
	YCBCR_422_APPLE = 0x85B9
)
// APPLE_aux_depth_stencil

// APPLE_client_storage

// APPLE_element_array

func ElementPointerAPPLE(type_ Enum, pointer Pointer)  {
	C.goglElementPointerAPPLE((C.GLenum)(type_), (unsafe.Pointer)(pointer))
}
func DrawElementArrayAPPLE(mode Enum, first Int, count Sizei)  {
	C.goglDrawElementArrayAPPLE((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
func DrawRangeElementArrayAPPLE(mode Enum, start Uint, end Uint, first Int, count Sizei)  {
	C.goglDrawRangeElementArrayAPPLE((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLint)(first), (C.GLsizei)(count))
}
func MultiDrawElementArrayAPPLE(mode Enum, first *Int, count *Sizei, primcount Sizei)  {
	C.goglMultiDrawElementArrayAPPLE((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount))
}
func MultiDrawRangeElementArrayAPPLE(mode Enum, start Uint, end Uint, first *Int, count *Sizei, primcount Sizei)  {
	C.goglMultiDrawRangeElementArrayAPPLE((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(primcount))
}
// APPLE_fence

func GenFencesAPPLE(n Sizei, fences *Uint)  {
	C.goglGenFencesAPPLE((C.GLsizei)(n), (*C.GLuint)(fences))
}
func DeleteFencesAPPLE(n Sizei, fences *Uint)  {
	C.goglDeleteFencesAPPLE((C.GLsizei)(n), (*C.GLuint)(fences))
}
func SetFenceAPPLE(fence Uint)  {
	C.goglSetFenceAPPLE((C.GLuint)(fence))
}
func IsFenceAPPLE(fence Uint) Boolean {
	return (Boolean)(C.goglIsFenceAPPLE((C.GLuint)(fence)))
}
func TestFenceAPPLE(fence Uint) Boolean {
	return (Boolean)(C.goglTestFenceAPPLE((C.GLuint)(fence)))
}
func FinishFenceAPPLE(fence Uint)  {
	C.goglFinishFenceAPPLE((C.GLuint)(fence))
}
func TestObjectAPPLE(object Enum, name Uint) Boolean {
	return (Boolean)(C.goglTestObjectAPPLE((C.GLenum)(object), (C.GLuint)(name)))
}
func FinishObjectAPPLE(object Enum, name Int)  {
	C.goglFinishObjectAPPLE((C.GLenum)(object), (C.GLint)(name))
}
// APPLE_float_pixels

// APPLE_flush_buffer_range

func BufferParameteriAPPLE(target Enum, pname Enum, param Int)  {
	C.goglBufferParameteriAPPLE((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func FlushMappedBufferRangeAPPLE(target Enum, offset Intptr, size Sizeiptr)  {
	C.goglFlushMappedBufferRangeAPPLE((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
// APPLE_object_purgeable

func ObjectPurgeableAPPLE(objectType Enum, name Uint, option Enum) Enum {
	return (Enum)(C.goglObjectPurgeableAPPLE((C.GLenum)(objectType), (C.GLuint)(name), (C.GLenum)(option)))
}
func ObjectUnpurgeableAPPLE(objectType Enum, name Uint, option Enum) Enum {
	return (Enum)(C.goglObjectUnpurgeableAPPLE((C.GLenum)(objectType), (C.GLuint)(name), (C.GLenum)(option)))
}
func GetObjectParameterivAPPLE(objectType Enum, name Uint, pname Enum, params *Int)  {
	C.goglGetObjectParameterivAPPLE((C.GLenum)(objectType), (C.GLuint)(name), (C.GLenum)(pname), (*C.GLint)(params))
}
// APPLE_rgb_422

// APPLE_row_bytes

// APPLE_specular_vector

// APPLE_texture_range

func TextureRangeAPPLE(target Enum, length Sizei, pointer Pointer)  {
	C.goglTextureRangeAPPLE((C.GLenum)(target), (C.GLsizei)(length), (unsafe.Pointer)(pointer))
}
func GetTexParameterPointervAPPLE(target Enum, pname Enum, params *Pointer)  {
	C.goglGetTexParameterPointervAPPLE((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}
// APPLE_transform_hint

// APPLE_vertex_array_object

func BindVertexArrayAPPLE(array Uint)  {
	C.goglBindVertexArrayAPPLE((C.GLuint)(array))
}
func DeleteVertexArraysAPPLE(n Sizei, arrays *Uint)  {
	C.goglDeleteVertexArraysAPPLE((C.GLsizei)(n), (*C.GLuint)(arrays))
}
func GenVertexArraysAPPLE(n Sizei, arrays *Uint)  {
	C.goglGenVertexArraysAPPLE((C.GLsizei)(n), (*C.GLuint)(arrays))
}
func IsVertexArrayAPPLE(array Uint) Boolean {
	return (Boolean)(C.goglIsVertexArrayAPPLE((C.GLuint)(array)))
}
// APPLE_vertex_array_range

func VertexArrayRangeAPPLE(length Sizei, pointer Pointer)  {
	C.goglVertexArrayRangeAPPLE((C.GLsizei)(length), (unsafe.Pointer)(pointer))
}
func FlushVertexArrayRangeAPPLE(length Sizei, pointer Pointer)  {
	C.goglFlushVertexArrayRangeAPPLE((C.GLsizei)(length), (unsafe.Pointer)(pointer))
}
func VertexArrayParameteriAPPLE(pname Enum, param Int)  {
	C.goglVertexArrayParameteriAPPLE((C.GLenum)(pname), (C.GLint)(param))
}
// APPLE_vertex_program_evaluators

func EnableVertexAttribAPPLE(index Uint, pname Enum)  {
	C.goglEnableVertexAttribAPPLE((C.GLuint)(index), (C.GLenum)(pname))
}
func DisableVertexAttribAPPLE(index Uint, pname Enum)  {
	C.goglDisableVertexAttribAPPLE((C.GLuint)(index), (C.GLenum)(pname))
}
func IsVertexAttribEnabledAPPLE(index Uint, pname Enum) Boolean {
	return (Boolean)(C.goglIsVertexAttribEnabledAPPLE((C.GLuint)(index), (C.GLenum)(pname)))
}
func MapVertexAttrib1dAPPLE(index Uint, size Uint, u1 Double, u2 Double, stride Int, order Int, points *Double)  {
	C.goglMapVertexAttrib1dAPPLE((C.GLuint)(index), (C.GLuint)(size), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLdouble)(points))
}
func MapVertexAttrib1fAPPLE(index Uint, size Uint, u1 Float, u2 Float, stride Int, order Int, points *Float)  {
	C.goglMapVertexAttrib1fAPPLE((C.GLuint)(index), (C.GLuint)(size), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLfloat)(points))
}
func MapVertexAttrib2dAPPLE(index Uint, size Uint, u1 Double, u2 Double, ustride Int, uorder Int, v1 Double, v2 Double, vstride Int, vorder Int, points *Double)  {
	C.goglMapVertexAttrib2dAPPLE((C.GLuint)(index), (C.GLuint)(size), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLdouble)(points))
}
func MapVertexAttrib2fAPPLE(index Uint, size Uint, u1 Float, u2 Float, ustride Int, uorder Int, v1 Float, v2 Float, vstride Int, vorder Int, points *Float)  {
	C.goglMapVertexAttrib2fAPPLE((C.GLuint)(index), (C.GLuint)(size), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLfloat)(points))
}
// APPLE_ycbcr_422

func InitAppleAuxDepthStencil() error {
	var ret C.int
	if ret = C.init_APPLE_aux_depth_stencil(); ret != 0 {
		return errors.New("unable to initialize APPLE_aux_depth_stencil")
	}
	return nil
}
func InitAppleClientStorage() error {
	var ret C.int
	if ret = C.init_APPLE_client_storage(); ret != 0 {
		return errors.New("unable to initialize APPLE_client_storage")
	}
	return nil
}
func InitAppleElementArray() error {
	var ret C.int
	if ret = C.init_APPLE_element_array(); ret != 0 {
		return errors.New("unable to initialize APPLE_element_array")
	}
	return nil
}
func InitAppleFence() error {
	var ret C.int
	if ret = C.init_APPLE_fence(); ret != 0 {
		return errors.New("unable to initialize APPLE_fence")
	}
	return nil
}
func InitAppleFloatPixels() error {
	var ret C.int
	if ret = C.init_APPLE_float_pixels(); ret != 0 {
		return errors.New("unable to initialize APPLE_float_pixels")
	}
	return nil
}
func InitAppleFlushBufferRange() error {
	var ret C.int
	if ret = C.init_APPLE_flush_buffer_range(); ret != 0 {
		return errors.New("unable to initialize APPLE_flush_buffer_range")
	}
	return nil
}
func InitAppleObjectPurgeable() error {
	var ret C.int
	if ret = C.init_APPLE_object_purgeable(); ret != 0 {
		return errors.New("unable to initialize APPLE_object_purgeable")
	}
	return nil
}
func InitAppleRgb422() error {
	var ret C.int
	if ret = C.init_APPLE_rgb_422(); ret != 0 {
		return errors.New("unable to initialize APPLE_rgb_422")
	}
	return nil
}
func InitAppleRowBytes() error {
	var ret C.int
	if ret = C.init_APPLE_row_bytes(); ret != 0 {
		return errors.New("unable to initialize APPLE_row_bytes")
	}
	return nil
}
func InitAppleSpecularVector() error {
	var ret C.int
	if ret = C.init_APPLE_specular_vector(); ret != 0 {
		return errors.New("unable to initialize APPLE_specular_vector")
	}
	return nil
}
func InitAppleTextureRange() error {
	var ret C.int
	if ret = C.init_APPLE_texture_range(); ret != 0 {
		return errors.New("unable to initialize APPLE_texture_range")
	}
	return nil
}
func InitAppleTransformHint() error {
	var ret C.int
	if ret = C.init_APPLE_transform_hint(); ret != 0 {
		return errors.New("unable to initialize APPLE_transform_hint")
	}
	return nil
}
func InitAppleVertexArrayObject() error {
	var ret C.int
	if ret = C.init_APPLE_vertex_array_object(); ret != 0 {
		return errors.New("unable to initialize APPLE_vertex_array_object")
	}
	return nil
}
func InitAppleVertexArrayRange() error {
	var ret C.int
	if ret = C.init_APPLE_vertex_array_range(); ret != 0 {
		return errors.New("unable to initialize APPLE_vertex_array_range")
	}
	return nil
}
func InitAppleVertexProgramEvaluators() error {
	var ret C.int
	if ret = C.init_APPLE_vertex_program_evaluators(); ret != 0 {
		return errors.New("unable to initialize APPLE_vertex_program_evaluators")
	}
	return nil
}
func InitAppleYcbcr422() error {
	var ret C.int
	if ret = C.init_APPLE_ycbcr_422(); ret != 0 {
		return errors.New("unable to initialize APPLE_ycbcr_422")
	}
	return nil
}
// EOF