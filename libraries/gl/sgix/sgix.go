// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// SGIX_async: http://www.opengl.org/registry/specs/SGIX/async.txt
// 
// SGIX_async_histogram: http://www.opengl.org/registry/specs/SGIX/async_histogram.txt
// 
// SGIX_async_pixel: http://www.opengl.org/registry/specs/SGIX/async_pixel.txt
// 
// SGIX_blend_alpha_minmax: http://www.opengl.org/registry/specs/SGIX/blend_alpha_minmax.txt
// 
// SGIX_calligraphic_fragment: http://www.opengl.org/registry/specs/SGIX/calligraphic_fragment.txt
// 
// SGIX_clipmap: http://www.opengl.org/registry/specs/SGIX/clipmap.txt
// 
// SGIX_convolution_accuracy: http://www.opengl.org/registry/specs/SGIX/convolution_accuracy.txt
// 
// SGIX_depth_pass_instrument: http://www.opengl.org/registry/specs/SGIX/depth_pass_instrument.txt
// 
// SGIX_depth_texture: http://www.opengl.org/registry/specs/SGIX/depth_texture.txt
// 
// SGIX_flush_raster: http://www.opengl.org/registry/specs/SGIX/flush_raster.txt
// 
// SGIX_fog_offset: http://www.opengl.org/registry/specs/SGIX/fog_offset.txt
// 
// SGIX_fog_scale: http://www.opengl.org/registry/specs/SGIX/fog_scale.txt
// 
// SGIX_fragment_lighting: http://www.opengl.org/registry/specs/SGIX/fragment_lighting.txt
// 
// SGIX_framezoom: http://www.opengl.org/registry/specs/SGIX/framezoom.txt
// 
// SGIX_igloo_interface: http://www.opengl.org/registry/specs/SGIX/igloo_interface.txt
// 
// SGIX_instruments: http://www.opengl.org/registry/specs/SGIX/instruments.txt
// 
// SGIX_interlace: http://www.opengl.org/registry/specs/SGIX/interlace.txt
// 
// SGIX_ir_instrument1: http://www.opengl.org/registry/specs/SGIX/ir_instrument1.txt
// 
// SGIX_list_priority: http://www.opengl.org/registry/specs/SGIX/list_priority.txt
// 
// SGIX_pixel_texture: http://www.opengl.org/registry/specs/SGIX/pixel_texture.txt
// 
// SGIX_pixel_tiles: http://www.opengl.org/registry/specs/SGIX/pixel_tiles.txt
// 
// SGIX_polynomial_ffd: http://www.opengl.org/registry/specs/SGIX/polynomial_ffd.txt
// 
// SGIX_reference_plane: http://www.opengl.org/registry/specs/SGIX/reference_plane.txt
// 
// SGIX_resample: http://www.opengl.org/registry/specs/SGIX/resample.txt
// 
// SGIX_scalebias_hint: http://www.opengl.org/registry/specs/SGIX/scalebias_hint.txt
// 
// SGIX_shadow: http://www.opengl.org/registry/specs/SGIX/shadow.txt
// 
// SGIX_shadow_ambient: http://www.opengl.org/registry/specs/SGIX/shadow_ambient.txt
// 
// SGIX_sprite: http://www.opengl.org/registry/specs/SGIX/sprite.txt
// 
// SGIX_subsample: http://www.opengl.org/registry/specs/SGIX/subsample.txt
// 
// SGIX_tag_sample_buffer: http://www.opengl.org/registry/specs/SGIX/tag_sample_buffer.txt
// 
// SGIX_texture_add_env: http://www.opengl.org/registry/specs/SGIX/texture_add_env.txt
// 
// SGIX_texture_coordinate_clamp: http://www.opengl.org/registry/specs/SGIX/texture_coordinate_clamp.txt
// 
// SGIX_texture_lod_bias: http://www.opengl.org/registry/specs/SGIX/texture_lod_bias.txt
// 
// SGIX_texture_multi_buffer: http://www.opengl.org/registry/specs/SGIX/texture_multi_buffer.txt
// 
// SGIX_texture_scale_bias: http://www.opengl.org/registry/specs/SGIX/texture_scale_bias.txt
// 
// SGIX_texture_select: http://www.opengl.org/registry/specs/SGIX/texture_select.txt
// 
// SGIX_vertex_preclip: http://www.opengl.org/registry/specs/SGIX/vertex_preclip.txt
// 
// SGIX_ycrcb: http://www.opengl.org/registry/specs/SGIX/ycrcb.txt
// 
// SGIX_ycrcb_subsample: http://www.opengl.org/registry/specs/SGIX/ycrcb_subsample.txt
// 
// SGIX_ycrcba: http://www.opengl.org/registry/specs/SGIX/ycrcba.txt
// 
package sgix

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
// //  SGIX_async
// void (APIENTRYP ptrglAsyncMarkerSGIX)(GLuint marker);
// GLint (APIENTRYP ptrglFinishAsyncSGIX)(GLuint* markerp);
// GLint (APIENTRYP ptrglPollAsyncSGIX)(GLuint* markerp);
// GLuint (APIENTRYP ptrglGenAsyncMarkersSGIX)(GLsizei range);
// void (APIENTRYP ptrglDeleteAsyncMarkersSGIX)(GLuint marker, GLsizei range);
// GLboolean (APIENTRYP ptrglIsAsyncMarkerSGIX)(GLuint marker);
// //  SGIX_async_histogram
// //  SGIX_async_pixel
// //  SGIX_blend_alpha_minmax
// //  SGIX_calligraphic_fragment
// //  SGIX_clipmap
// //  SGIX_convolution_accuracy
// //  SGIX_depth_pass_instrument
// //  SGIX_depth_texture
// //  SGIX_flush_raster
// void (APIENTRYP ptrglFlushRasterSGIX)();
// //  SGIX_fog_offset
// //  SGIX_fog_scale
// //  SGIX_fragment_lighting
// void (APIENTRYP ptrglFragmentColorMaterialSGIX)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglFragmentLightfSGIX)(GLenum light, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglFragmentLightfvSGIX)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglFragmentLightiSGIX)(GLenum light, GLenum pname, GLint param);
// void (APIENTRYP ptrglFragmentLightivSGIX)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFragmentLightModelfSGIX)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglFragmentLightModelfvSGIX)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglFragmentLightModeliSGIX)(GLenum pname, GLint param);
// void (APIENTRYP ptrglFragmentLightModelivSGIX)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglFragmentMaterialfSGIX)(GLenum face, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglFragmentMaterialfvSGIX)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglFragmentMaterialiSGIX)(GLenum face, GLenum pname, GLint param);
// void (APIENTRYP ptrglFragmentMaterialivSGIX)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetFragmentLightfvSGIX)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetFragmentLightivSGIX)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetFragmentMaterialfvSGIX)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetFragmentMaterialivSGIX)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrglLightEnviSGIX)(GLenum pname, GLint param);
// //  SGIX_framezoom
// void (APIENTRYP ptrglFrameZoomSGIX)(GLint factor);
// //  SGIX_igloo_interface
// void (APIENTRYP ptrglIglooInterfaceSGIX)(GLenum pname, GLvoid* params);
// //  SGIX_instruments
// GLint (APIENTRYP ptrglGetInstrumentsSGIX)();
// void (APIENTRYP ptrglInstrumentsBufferSGIX)(GLsizei size, GLint* buffer);
// GLint (APIENTRYP ptrglPollInstrumentsSGIX)(GLint* marker_p);
// void (APIENTRYP ptrglReadInstrumentsSGIX)(GLint marker);
// void (APIENTRYP ptrglStartInstrumentsSGIX)();
// void (APIENTRYP ptrglStopInstrumentsSGIX)(GLint marker);
// //  SGIX_interlace
// //  SGIX_ir_instrument1
// //  SGIX_list_priority
// void (APIENTRYP ptrglGetListParameterfvSGIX)(GLuint list, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetListParameterivSGIX)(GLuint list, GLenum pname, GLint* params);
// void (APIENTRYP ptrglListParameterfSGIX)(GLuint list, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglListParameterfvSGIX)(GLuint list, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglListParameteriSGIX)(GLuint list, GLenum pname, GLint param);
// void (APIENTRYP ptrglListParameterivSGIX)(GLuint list, GLenum pname, GLint* params);
// //  SGIX_pixel_texture
// void (APIENTRYP ptrglPixelTexGenSGIX)(GLenum mode);
// //  SGIX_pixel_tiles
// //  SGIX_polynomial_ffd
// void (APIENTRYP ptrglDeformationMap3dSGIX)(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble w1, GLdouble w2, GLint wstride, GLint worder, GLdouble* points);
// void (APIENTRYP ptrglDeformationMap3fSGIX)(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat w1, GLfloat w2, GLint wstride, GLint worder, GLfloat* points);
// void (APIENTRYP ptrglDeformSGIX)(GLbitfield mask);
// void (APIENTRYP ptrglLoadIdentityDeformationMapSGIX)(GLbitfield mask);
// //  SGIX_reference_plane
// void (APIENTRYP ptrglReferencePlaneSGIX)(GLdouble* equation);
// //  SGIX_resample
// //  SGIX_scalebias_hint
// //  SGIX_shadow
// //  SGIX_shadow_ambient
// //  SGIX_sprite
// void (APIENTRYP ptrglSpriteParameterfSGIX)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglSpriteParameterfvSGIX)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglSpriteParameteriSGIX)(GLenum pname, GLint param);
// void (APIENTRYP ptrglSpriteParameterivSGIX)(GLenum pname, GLint* params);
// //  SGIX_subsample
// //  SGIX_tag_sample_buffer
// void (APIENTRYP ptrglTagSampleBufferSGIX)();
// //  SGIX_texture_add_env
// //  SGIX_texture_coordinate_clamp
// //  SGIX_texture_lod_bias
// //  SGIX_texture_multi_buffer
// //  SGIX_texture_scale_bias
// //  SGIX_texture_select
// //  SGIX_vertex_preclip
// //  SGIX_ycrcb
// //  SGIX_ycrcb_subsample
// //  SGIX_ycrcba
// 
// //  SGIX_async
// void goglAsyncMarkerSGIX(GLuint marker) {
// 	(*ptrglAsyncMarkerSGIX)(marker);
// }
// GLint goglFinishAsyncSGIX(GLuint* markerp) {
// 	return (*ptrglFinishAsyncSGIX)(markerp);
// }
// GLint goglPollAsyncSGIX(GLuint* markerp) {
// 	return (*ptrglPollAsyncSGIX)(markerp);
// }
// GLuint goglGenAsyncMarkersSGIX(GLsizei range_) {
// 	return (*ptrglGenAsyncMarkersSGIX)(range_);
// }
// void goglDeleteAsyncMarkersSGIX(GLuint marker, GLsizei range_) {
// 	(*ptrglDeleteAsyncMarkersSGIX)(marker, range_);
// }
// GLboolean goglIsAsyncMarkerSGIX(GLuint marker) {
// 	return (*ptrglIsAsyncMarkerSGIX)(marker);
// }
// //  SGIX_async_histogram
// //  SGIX_async_pixel
// //  SGIX_blend_alpha_minmax
// //  SGIX_calligraphic_fragment
// //  SGIX_clipmap
// //  SGIX_convolution_accuracy
// //  SGIX_depth_pass_instrument
// //  SGIX_depth_texture
// //  SGIX_flush_raster
// void goglFlushRasterSGIX() {
// 	(*ptrglFlushRasterSGIX)();
// }
// //  SGIX_fog_offset
// //  SGIX_fog_scale
// //  SGIX_fragment_lighting
// void goglFragmentColorMaterialSGIX(GLenum face, GLenum mode) {
// 	(*ptrglFragmentColorMaterialSGIX)(face, mode);
// }
// void goglFragmentLightfSGIX(GLenum light, GLenum pname, GLfloat param) {
// 	(*ptrglFragmentLightfSGIX)(light, pname, param);
// }
// void goglFragmentLightfvSGIX(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrglFragmentLightfvSGIX)(light, pname, params);
// }
// void goglFragmentLightiSGIX(GLenum light, GLenum pname, GLint param) {
// 	(*ptrglFragmentLightiSGIX)(light, pname, param);
// }
// void goglFragmentLightivSGIX(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrglFragmentLightivSGIX)(light, pname, params);
// }
// void goglFragmentLightModelfSGIX(GLenum pname, GLfloat param) {
// 	(*ptrglFragmentLightModelfSGIX)(pname, param);
// }
// void goglFragmentLightModelfvSGIX(GLenum pname, GLfloat* params) {
// 	(*ptrglFragmentLightModelfvSGIX)(pname, params);
// }
// void goglFragmentLightModeliSGIX(GLenum pname, GLint param) {
// 	(*ptrglFragmentLightModeliSGIX)(pname, param);
// }
// void goglFragmentLightModelivSGIX(GLenum pname, GLint* params) {
// 	(*ptrglFragmentLightModelivSGIX)(pname, params);
// }
// void goglFragmentMaterialfSGIX(GLenum face, GLenum pname, GLfloat param) {
// 	(*ptrglFragmentMaterialfSGIX)(face, pname, param);
// }
// void goglFragmentMaterialfvSGIX(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrglFragmentMaterialfvSGIX)(face, pname, params);
// }
// void goglFragmentMaterialiSGIX(GLenum face, GLenum pname, GLint param) {
// 	(*ptrglFragmentMaterialiSGIX)(face, pname, param);
// }
// void goglFragmentMaterialivSGIX(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrglFragmentMaterialivSGIX)(face, pname, params);
// }
// void goglGetFragmentLightfvSGIX(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrglGetFragmentLightfvSGIX)(light, pname, params);
// }
// void goglGetFragmentLightivSGIX(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrglGetFragmentLightivSGIX)(light, pname, params);
// }
// void goglGetFragmentMaterialfvSGIX(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrglGetFragmentMaterialfvSGIX)(face, pname, params);
// }
// void goglGetFragmentMaterialivSGIX(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrglGetFragmentMaterialivSGIX)(face, pname, params);
// }
// void goglLightEnviSGIX(GLenum pname, GLint param) {
// 	(*ptrglLightEnviSGIX)(pname, param);
// }
// //  SGIX_framezoom
// void goglFrameZoomSGIX(GLint factor) {
// 	(*ptrglFrameZoomSGIX)(factor);
// }
// //  SGIX_igloo_interface
// void goglIglooInterfaceSGIX(GLenum pname, GLvoid* params) {
// 	(*ptrglIglooInterfaceSGIX)(pname, params);
// }
// //  SGIX_instruments
// GLint goglGetInstrumentsSGIX() {
// 	return (*ptrglGetInstrumentsSGIX)();
// }
// void goglInstrumentsBufferSGIX(GLsizei size, GLint* buffer) {
// 	(*ptrglInstrumentsBufferSGIX)(size, buffer);
// }
// GLint goglPollInstrumentsSGIX(GLint* marker_p) {
// 	return (*ptrglPollInstrumentsSGIX)(marker_p);
// }
// void goglReadInstrumentsSGIX(GLint marker) {
// 	(*ptrglReadInstrumentsSGIX)(marker);
// }
// void goglStartInstrumentsSGIX() {
// 	(*ptrglStartInstrumentsSGIX)();
// }
// void goglStopInstrumentsSGIX(GLint marker) {
// 	(*ptrglStopInstrumentsSGIX)(marker);
// }
// //  SGIX_interlace
// //  SGIX_ir_instrument1
// //  SGIX_list_priority
// void goglGetListParameterfvSGIX(GLuint list, GLenum pname, GLfloat* params) {
// 	(*ptrglGetListParameterfvSGIX)(list, pname, params);
// }
// void goglGetListParameterivSGIX(GLuint list, GLenum pname, GLint* params) {
// 	(*ptrglGetListParameterivSGIX)(list, pname, params);
// }
// void goglListParameterfSGIX(GLuint list, GLenum pname, GLfloat param) {
// 	(*ptrglListParameterfSGIX)(list, pname, param);
// }
// void goglListParameterfvSGIX(GLuint list, GLenum pname, GLfloat* params) {
// 	(*ptrglListParameterfvSGIX)(list, pname, params);
// }
// void goglListParameteriSGIX(GLuint list, GLenum pname, GLint param) {
// 	(*ptrglListParameteriSGIX)(list, pname, param);
// }
// void goglListParameterivSGIX(GLuint list, GLenum pname, GLint* params) {
// 	(*ptrglListParameterivSGIX)(list, pname, params);
// }
// //  SGIX_pixel_texture
// void goglPixelTexGenSGIX(GLenum mode) {
// 	(*ptrglPixelTexGenSGIX)(mode);
// }
// //  SGIX_pixel_tiles
// //  SGIX_polynomial_ffd
// void goglDeformationMap3dSGIX(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble w1, GLdouble w2, GLint wstride, GLint worder, GLdouble* points) {
// 	(*ptrglDeformationMap3dSGIX)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, w1, w2, wstride, worder, points);
// }
// void goglDeformationMap3fSGIX(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat w1, GLfloat w2, GLint wstride, GLint worder, GLfloat* points) {
// 	(*ptrglDeformationMap3fSGIX)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, w1, w2, wstride, worder, points);
// }
// void goglDeformSGIX(GLbitfield mask) {
// 	(*ptrglDeformSGIX)(mask);
// }
// void goglLoadIdentityDeformationMapSGIX(GLbitfield mask) {
// 	(*ptrglLoadIdentityDeformationMapSGIX)(mask);
// }
// //  SGIX_reference_plane
// void goglReferencePlaneSGIX(GLdouble* equation) {
// 	(*ptrglReferencePlaneSGIX)(equation);
// }
// //  SGIX_resample
// //  SGIX_scalebias_hint
// //  SGIX_shadow
// //  SGIX_shadow_ambient
// //  SGIX_sprite
// void goglSpriteParameterfSGIX(GLenum pname, GLfloat param) {
// 	(*ptrglSpriteParameterfSGIX)(pname, param);
// }
// void goglSpriteParameterfvSGIX(GLenum pname, GLfloat* params) {
// 	(*ptrglSpriteParameterfvSGIX)(pname, params);
// }
// void goglSpriteParameteriSGIX(GLenum pname, GLint param) {
// 	(*ptrglSpriteParameteriSGIX)(pname, param);
// }
// void goglSpriteParameterivSGIX(GLenum pname, GLint* params) {
// 	(*ptrglSpriteParameterivSGIX)(pname, params);
// }
// //  SGIX_subsample
// //  SGIX_tag_sample_buffer
// void goglTagSampleBufferSGIX() {
// 	(*ptrglTagSampleBufferSGIX)();
// }
// //  SGIX_texture_add_env
// //  SGIX_texture_coordinate_clamp
// //  SGIX_texture_lod_bias
// //  SGIX_texture_multi_buffer
// //  SGIX_texture_scale_bias
// //  SGIX_texture_select
// //  SGIX_vertex_preclip
// //  SGIX_ycrcb
// //  SGIX_ycrcb_subsample
// //  SGIX_ycrcba
// 
// int init_SGIX_async() {
// 	ptrglAsyncMarkerSGIX = goglGetProcAddress("glAsyncMarkerSGIX");
// 	if(ptrglAsyncMarkerSGIX == NULL) return 1;
// 	ptrglFinishAsyncSGIX = goglGetProcAddress("glFinishAsyncSGIX");
// 	if(ptrglFinishAsyncSGIX == NULL) return 1;
// 	ptrglPollAsyncSGIX = goglGetProcAddress("glPollAsyncSGIX");
// 	if(ptrglPollAsyncSGIX == NULL) return 1;
// 	ptrglGenAsyncMarkersSGIX = goglGetProcAddress("glGenAsyncMarkersSGIX");
// 	if(ptrglGenAsyncMarkersSGIX == NULL) return 1;
// 	ptrglDeleteAsyncMarkersSGIX = goglGetProcAddress("glDeleteAsyncMarkersSGIX");
// 	if(ptrglDeleteAsyncMarkersSGIX == NULL) return 1;
// 	ptrglIsAsyncMarkerSGIX = goglGetProcAddress("glIsAsyncMarkerSGIX");
// 	if(ptrglIsAsyncMarkerSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_async_histogram() {
// 	return 0;
// }
// int init_SGIX_async_pixel() {
// 	return 0;
// }
// int init_SGIX_blend_alpha_minmax() {
// 	return 0;
// }
// int init_SGIX_calligraphic_fragment() {
// 	return 0;
// }
// int init_SGIX_clipmap() {
// 	return 0;
// }
// int init_SGIX_convolution_accuracy() {
// 	return 0;
// }
// int init_SGIX_depth_pass_instrument() {
// 	return 0;
// }
// int init_SGIX_depth_texture() {
// 	return 0;
// }
// int init_SGIX_flush_raster() {
// 	ptrglFlushRasterSGIX = goglGetProcAddress("glFlushRasterSGIX");
// 	if(ptrglFlushRasterSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_fog_offset() {
// 	return 0;
// }
// int init_SGIX_fog_scale() {
// 	return 0;
// }
// int init_SGIX_fragment_lighting() {
// 	ptrglFragmentColorMaterialSGIX = goglGetProcAddress("glFragmentColorMaterialSGIX");
// 	if(ptrglFragmentColorMaterialSGIX == NULL) return 1;
// 	ptrglFragmentLightfSGIX = goglGetProcAddress("glFragmentLightfSGIX");
// 	if(ptrglFragmentLightfSGIX == NULL) return 1;
// 	ptrglFragmentLightfvSGIX = goglGetProcAddress("glFragmentLightfvSGIX");
// 	if(ptrglFragmentLightfvSGIX == NULL) return 1;
// 	ptrglFragmentLightiSGIX = goglGetProcAddress("glFragmentLightiSGIX");
// 	if(ptrglFragmentLightiSGIX == NULL) return 1;
// 	ptrglFragmentLightivSGIX = goglGetProcAddress("glFragmentLightivSGIX");
// 	if(ptrglFragmentLightivSGIX == NULL) return 1;
// 	ptrglFragmentLightModelfSGIX = goglGetProcAddress("glFragmentLightModelfSGIX");
// 	if(ptrglFragmentLightModelfSGIX == NULL) return 1;
// 	ptrglFragmentLightModelfvSGIX = goglGetProcAddress("glFragmentLightModelfvSGIX");
// 	if(ptrglFragmentLightModelfvSGIX == NULL) return 1;
// 	ptrglFragmentLightModeliSGIX = goglGetProcAddress("glFragmentLightModeliSGIX");
// 	if(ptrglFragmentLightModeliSGIX == NULL) return 1;
// 	ptrglFragmentLightModelivSGIX = goglGetProcAddress("glFragmentLightModelivSGIX");
// 	if(ptrglFragmentLightModelivSGIX == NULL) return 1;
// 	ptrglFragmentMaterialfSGIX = goglGetProcAddress("glFragmentMaterialfSGIX");
// 	if(ptrglFragmentMaterialfSGIX == NULL) return 1;
// 	ptrglFragmentMaterialfvSGIX = goglGetProcAddress("glFragmentMaterialfvSGIX");
// 	if(ptrglFragmentMaterialfvSGIX == NULL) return 1;
// 	ptrglFragmentMaterialiSGIX = goglGetProcAddress("glFragmentMaterialiSGIX");
// 	if(ptrglFragmentMaterialiSGIX == NULL) return 1;
// 	ptrglFragmentMaterialivSGIX = goglGetProcAddress("glFragmentMaterialivSGIX");
// 	if(ptrglFragmentMaterialivSGIX == NULL) return 1;
// 	ptrglGetFragmentLightfvSGIX = goglGetProcAddress("glGetFragmentLightfvSGIX");
// 	if(ptrglGetFragmentLightfvSGIX == NULL) return 1;
// 	ptrglGetFragmentLightivSGIX = goglGetProcAddress("glGetFragmentLightivSGIX");
// 	if(ptrglGetFragmentLightivSGIX == NULL) return 1;
// 	ptrglGetFragmentMaterialfvSGIX = goglGetProcAddress("glGetFragmentMaterialfvSGIX");
// 	if(ptrglGetFragmentMaterialfvSGIX == NULL) return 1;
// 	ptrglGetFragmentMaterialivSGIX = goglGetProcAddress("glGetFragmentMaterialivSGIX");
// 	if(ptrglGetFragmentMaterialivSGIX == NULL) return 1;
// 	ptrglLightEnviSGIX = goglGetProcAddress("glLightEnviSGIX");
// 	if(ptrglLightEnviSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_framezoom() {
// 	ptrglFrameZoomSGIX = goglGetProcAddress("glFrameZoomSGIX");
// 	if(ptrglFrameZoomSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_igloo_interface() {
// 	ptrglIglooInterfaceSGIX = goglGetProcAddress("glIglooInterfaceSGIX");
// 	if(ptrglIglooInterfaceSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_instruments() {
// 	ptrglGetInstrumentsSGIX = goglGetProcAddress("glGetInstrumentsSGIX");
// 	if(ptrglGetInstrumentsSGIX == NULL) return 1;
// 	ptrglInstrumentsBufferSGIX = goglGetProcAddress("glInstrumentsBufferSGIX");
// 	if(ptrglInstrumentsBufferSGIX == NULL) return 1;
// 	ptrglPollInstrumentsSGIX = goglGetProcAddress("glPollInstrumentsSGIX");
// 	if(ptrglPollInstrumentsSGIX == NULL) return 1;
// 	ptrglReadInstrumentsSGIX = goglGetProcAddress("glReadInstrumentsSGIX");
// 	if(ptrglReadInstrumentsSGIX == NULL) return 1;
// 	ptrglStartInstrumentsSGIX = goglGetProcAddress("glStartInstrumentsSGIX");
// 	if(ptrglStartInstrumentsSGIX == NULL) return 1;
// 	ptrglStopInstrumentsSGIX = goglGetProcAddress("glStopInstrumentsSGIX");
// 	if(ptrglStopInstrumentsSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_interlace() {
// 	return 0;
// }
// int init_SGIX_ir_instrument1() {
// 	return 0;
// }
// int init_SGIX_list_priority() {
// 	ptrglGetListParameterfvSGIX = goglGetProcAddress("glGetListParameterfvSGIX");
// 	if(ptrglGetListParameterfvSGIX == NULL) return 1;
// 	ptrglGetListParameterivSGIX = goglGetProcAddress("glGetListParameterivSGIX");
// 	if(ptrglGetListParameterivSGIX == NULL) return 1;
// 	ptrglListParameterfSGIX = goglGetProcAddress("glListParameterfSGIX");
// 	if(ptrglListParameterfSGIX == NULL) return 1;
// 	ptrglListParameterfvSGIX = goglGetProcAddress("glListParameterfvSGIX");
// 	if(ptrglListParameterfvSGIX == NULL) return 1;
// 	ptrglListParameteriSGIX = goglGetProcAddress("glListParameteriSGIX");
// 	if(ptrglListParameteriSGIX == NULL) return 1;
// 	ptrglListParameterivSGIX = goglGetProcAddress("glListParameterivSGIX");
// 	if(ptrglListParameterivSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_pixel_texture() {
// 	ptrglPixelTexGenSGIX = goglGetProcAddress("glPixelTexGenSGIX");
// 	if(ptrglPixelTexGenSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_pixel_tiles() {
// 	return 0;
// }
// int init_SGIX_polynomial_ffd() {
// 	ptrglDeformationMap3dSGIX = goglGetProcAddress("glDeformationMap3dSGIX");
// 	if(ptrglDeformationMap3dSGIX == NULL) return 1;
// 	ptrglDeformationMap3fSGIX = goglGetProcAddress("glDeformationMap3fSGIX");
// 	if(ptrglDeformationMap3fSGIX == NULL) return 1;
// 	ptrglDeformSGIX = goglGetProcAddress("glDeformSGIX");
// 	if(ptrglDeformSGIX == NULL) return 1;
// 	ptrglLoadIdentityDeformationMapSGIX = goglGetProcAddress("glLoadIdentityDeformationMapSGIX");
// 	if(ptrglLoadIdentityDeformationMapSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_reference_plane() {
// 	ptrglReferencePlaneSGIX = goglGetProcAddress("glReferencePlaneSGIX");
// 	if(ptrglReferencePlaneSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_resample() {
// 	return 0;
// }
// int init_SGIX_scalebias_hint() {
// 	return 0;
// }
// int init_SGIX_shadow() {
// 	return 0;
// }
// int init_SGIX_shadow_ambient() {
// 	return 0;
// }
// int init_SGIX_sprite() {
// 	ptrglSpriteParameterfSGIX = goglGetProcAddress("glSpriteParameterfSGIX");
// 	if(ptrglSpriteParameterfSGIX == NULL) return 1;
// 	ptrglSpriteParameterfvSGIX = goglGetProcAddress("glSpriteParameterfvSGIX");
// 	if(ptrglSpriteParameterfvSGIX == NULL) return 1;
// 	ptrglSpriteParameteriSGIX = goglGetProcAddress("glSpriteParameteriSGIX");
// 	if(ptrglSpriteParameteriSGIX == NULL) return 1;
// 	ptrglSpriteParameterivSGIX = goglGetProcAddress("glSpriteParameterivSGIX");
// 	if(ptrglSpriteParameterivSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_subsample() {
// 	return 0;
// }
// int init_SGIX_tag_sample_buffer() {
// 	ptrglTagSampleBufferSGIX = goglGetProcAddress("glTagSampleBufferSGIX");
// 	if(ptrglTagSampleBufferSGIX == NULL) return 1;
// 	return 0;
// }
// int init_SGIX_texture_add_env() {
// 	return 0;
// }
// int init_SGIX_texture_coordinate_clamp() {
// 	return 0;
// }
// int init_SGIX_texture_lod_bias() {
// 	return 0;
// }
// int init_SGIX_texture_multi_buffer() {
// 	return 0;
// }
// int init_SGIX_texture_scale_bias() {
// 	return 0;
// }
// int init_SGIX_texture_select() {
// 	return 0;
// }
// int init_SGIX_vertex_preclip() {
// 	return 0;
// }
// int init_SGIX_ycrcb() {
// 	return 0;
// }
// int init_SGIX_ycrcb_subsample() {
// 	return 0;
// }
// int init_SGIX_ycrcba() {
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

// SGIX_async
const (
	ASYNC_MARKER_SGIX = 0x8329
)
// SGIX_async_histogram
const (
	ASYNC_HISTOGRAM_SGIX = 0x832C
	MAX_ASYNC_HISTOGRAM_SGIX = 0x832D
)
// SGIX_async_pixel
const (
	ASYNC_DRAW_PIXELS_SGIX = 0x835D
	ASYNC_READ_PIXELS_SGIX = 0x835E
	ASYNC_TEX_IMAGE_SGIX = 0x835C
	MAX_ASYNC_DRAW_PIXELS_SGIX = 0x8360
	MAX_ASYNC_READ_PIXELS_SGIX = 0x8361
	MAX_ASYNC_TEX_IMAGE_SGIX = 0x835F
)
// SGIX_blend_alpha_minmax
const (
	ALPHA_MAX_SGIX = 0x8321
	ALPHA_MIN_SGIX = 0x8320
)
// SGIX_calligraphic_fragment
const (
	CALLIGRAPHIC_FRAGMENT_SGIX = 0x8183
)
// SGIX_clipmap
const (
	LINEAR_CLIPMAP_LINEAR_SGIX = 0x8170
	LINEAR_CLIPMAP_NEAREST_SGIX = 0x844F
	MAX_CLIPMAP_DEPTH_SGIX = 0x8177
	MAX_CLIPMAP_VIRTUAL_DEPTH_SGIX = 0x8178
	NEAREST_CLIPMAP_LINEAR_SGIX = 0x844E
	NEAREST_CLIPMAP_NEAREST_SGIX = 0x844D
	TEXTURE_CLIPMAP_CENTER_SGIX = 0x8171
	TEXTURE_CLIPMAP_DEPTH_SGIX = 0x8176
	TEXTURE_CLIPMAP_FRAME_SGIX = 0x8172
	TEXTURE_CLIPMAP_LOD_OFFSET_SGIX = 0x8175
	TEXTURE_CLIPMAP_OFFSET_SGIX = 0x8173
	TEXTURE_CLIPMAP_VIRTUAL_DEPTH_SGIX = 0x8174
)
// SGIX_convolution_accuracy
const (
	CONVOLUTION_HINT_SGIX = 0x8316
)
// SGIX_depth_texture
const (
	DEPTH_COMPONENT16_SGIX = 0x81A5
	DEPTH_COMPONENT24_SGIX = 0x81A6
	DEPTH_COMPONENT32_SGIX = 0x81A7
)
// SGIX_flush_raster
const (
)
// SGIX_fog_offset
const (
	FOG_OFFSET_SGIX = 0x8198
	FOG_OFFSET_VALUE_SGIX = 0x8199
)
// SGIX_fog_scale
const (
	FOG_SCALE_SGIX = 0x81FC
	FOG_SCALE_VALUE_SGIX = 0x81FD
)
// SGIX_fragment_lighting
const (
	CURRENT_RASTER_NORMAL_SGIX = 0x8406
	FRAGMENT_COLOR_MATERIAL_FACE_SGIX = 0x8402
	FRAGMENT_COLOR_MATERIAL_PARAMETER_SGIX = 0x8403
	FRAGMENT_COLOR_MATERIAL_SGIX = 0x8401
	FRAGMENT_LIGHT0_SGIX = 0x840C
	FRAGMENT_LIGHT1_SGIX = 0x840D
	FRAGMENT_LIGHT2_SGIX = 0x840E
	FRAGMENT_LIGHT3_SGIX = 0x840F
	FRAGMENT_LIGHT4_SGIX = 0x8410
	FRAGMENT_LIGHT5_SGIX = 0x8411
	FRAGMENT_LIGHT6_SGIX = 0x8412
	FRAGMENT_LIGHT7_SGIX = 0x8413
	FRAGMENT_LIGHTING_SGIX = 0x8400
	FRAGMENT_LIGHT_MODEL_AMBIENT_SGIX = 0x840A
	FRAGMENT_LIGHT_MODEL_LOCAL_VIEWER_SGIX = 0x8408
	FRAGMENT_LIGHT_MODEL_NORMAL_INTERPOLATION_SGIX = 0x840B
	FRAGMENT_LIGHT_MODEL_TWO_SIDE_SGIX = 0x8409
	LIGHT_ENV_MODE_SGIX = 0x8407
	MAX_ACTIVE_LIGHTS_SGIX = 0x8405
	MAX_FRAGMENT_LIGHTS_SGIX = 0x8404
)
// SGIX_framezoom
const (
	FRAMEZOOM_FACTOR_SGIX = 0x818C
	FRAMEZOOM_SGIX = 0x818B
	MAX_FRAMEZOOM_FACTOR_SGIX = 0x818D
)
// SGIX_impact_pixel_texture
const (
	PIXEL_TEX_GEN_ALPHA_LS_SGIX = 0x8189
	PIXEL_TEX_GEN_ALPHA_MS_SGIX = 0x818A
	PIXEL_TEX_GEN_ALPHA_NO_REPLACE_SGIX = 0x8188
	PIXEL_TEX_GEN_ALPHA_REPLACE_SGIX = 0x8187
	PIXEL_TEX_GEN_Q_CEILING_SGIX = 0x8184
	PIXEL_TEX_GEN_Q_FLOOR_SGIX = 0x8186
	PIXEL_TEX_GEN_Q_ROUND_SGIX = 0x8185
)
// SGIX_instruments
const (
	INSTRUMENT_BUFFER_POINTER_SGIX = 0x8180
	INSTRUMENT_MEASUREMENTS_SGIX = 0x8181
)
// SGIX_interlace
const (
	INTERLACE_SGIX = 0x8094
)
// SGIX_ir_instrument1
const (
	IR_INSTRUMENT1_SGIX = 0x817F
)
// SGIX_list_priority
const (
	LIST_PRIORITY_SGIX = 0x8182
)
// SGIX_pixel_texture
const (
	PIXEL_TEX_GEN_MODE_SGIX = 0x832B
	PIXEL_TEX_GEN_SGIX = 0x8139
)
// SGIX_pixel_tiles
const (
	PIXEL_TILE_BEST_ALIGNMENT_SGIX = 0x813E
	PIXEL_TILE_CACHE_INCREMENT_SGIX = 0x813F
	PIXEL_TILE_CACHE_SIZE_SGIX = 0x8145
	PIXEL_TILE_GRID_DEPTH_SGIX = 0x8144
	PIXEL_TILE_GRID_HEIGHT_SGIX = 0x8143
	PIXEL_TILE_GRID_WIDTH_SGIX = 0x8142
	PIXEL_TILE_HEIGHT_SGIX = 0x8141
	PIXEL_TILE_WIDTH_SGIX = 0x8140
)
// SGIX_polynomial_ffd
const (
	DEFORMATIONS_MASK_SGIX = 0x8196
	GEOMETRY_DEFORMATION_SGIX = 0x8194
	MAX_DEFORMATION_ORDER_SGIX = 0x8197
	TEXTURE_DEFORMATION_SGIX = 0x8195
)
// SGIX_reference_plane
const (
	REFERENCE_PLANE_EQUATION_SGIX = 0x817E
	REFERENCE_PLANE_SGIX = 0x817D
)
// SGIX_resample
const (
	PACK_RESAMPLE_SGIX = 0x842C
	RESAMPLE_DECIMATE_SGIX = 0x8430
	RESAMPLE_REPLICATE_SGIX = 0x842E
	RESAMPLE_ZERO_FILL_SGIX = 0x842F
	UNPACK_RESAMPLE_SGIX = 0x842D
)
// SGIX_scalebias_hint
const (
	SCALEBIAS_HINT_SGIX = 0x8322
)
// SGIX_shadow
const (
	TEXTURE_COMPARE_OPERATOR_SGIX = 0x819B
	TEXTURE_COMPARE_SGIX = 0x819A
	TEXTURE_GEQUAL_R_SGIX = 0x819D
	TEXTURE_LEQUAL_R_SGIX = 0x819C
)
// SGIX_shadow_ambient
const (
	SHADOW_AMBIENT_SGIX = 0x80BF
)
// SGIX_sprite
const (
	SPRITE_AXIAL_SGIX = 0x814C
	SPRITE_AXIS_SGIX = 0x814A
	SPRITE_EYE_ALIGNED_SGIX = 0x814E
	SPRITE_MODE_SGIX = 0x8149
	SPRITE_OBJECT_ALIGNED_SGIX = 0x814D
	SPRITE_SGIX = 0x8148
	SPRITE_TRANSLATION_SGIX = 0x814B
)
// SGIX_subsample
const (
)
// SGIX_tag_sample_buffer
const (
)
// SGIX_texture_add_env
const (
	TEXTURE_ENV_BIAS_SGIX = 0x80BE
)
// SGIX_texture_coordinate_clamp
const (
	TEXTURE_MAX_CLAMP_R_SGIX = 0x836B
	TEXTURE_MAX_CLAMP_S_SGIX = 0x8369
	TEXTURE_MAX_CLAMP_T_SGIX = 0x836A
)
// SGIX_texture_lod_bias
const (
	TEXTURE_LOD_BIAS_R_SGIX = 0x8190
	TEXTURE_LOD_BIAS_S_SGIX = 0x818E
	TEXTURE_LOD_BIAS_T_SGIX = 0x818F
)
// SGIX_texture_multi_buffer
const (
	TEXTURE_MULTI_BUFFER_HINT_SGIX = 0x812E
)
// SGIX_texture_scale_bias
const (
	POST_TEXTURE_FILTER_BIAS_RANGE_SGIX = 0x817B
	POST_TEXTURE_FILTER_BIAS_SGIX = 0x8179
	POST_TEXTURE_FILTER_SCALE_RANGE_SGIX = 0x817C
	POST_TEXTURE_FILTER_SCALE_SGIX = 0x817A
)
// SGIX_vertex_preclip
const (
	VERTEX_PRECLIP_HINT_SGIX = 0x83EF
	VERTEX_PRECLIP_SGIX = 0x83EE
)
// SGIX_ycrcb
const (
	YCRCB_422_SGIX = 0x81BB
	YCRCB_444_SGIX = 0x81BC
)
// SGIX_ycrcb_subsample
const (
)
// SGIX_ycrcba
const (
	YCRCBA_SGIX = 0x8319
	YCRCB_SGIX = 0x8318
)
// SGIX_async

func AsyncMarkerSGIX(marker Uint)  {
	C.goglAsyncMarkerSGIX((C.GLuint)(marker))
}
func FinishAsyncSGIX(markerp *Uint) Int {
	return (Int)(C.goglFinishAsyncSGIX((*C.GLuint)(markerp)))
}
func PollAsyncSGIX(markerp *Uint) Int {
	return (Int)(C.goglPollAsyncSGIX((*C.GLuint)(markerp)))
}
func GenAsyncMarkersSGIX(range_ Sizei) Uint {
	return (Uint)(C.goglGenAsyncMarkersSGIX((C.GLsizei)(range_)))
}
func DeleteAsyncMarkersSGIX(marker Uint, range_ Sizei)  {
	C.goglDeleteAsyncMarkersSGIX((C.GLuint)(marker), (C.GLsizei)(range_))
}
func IsAsyncMarkerSGIX(marker Uint) Boolean {
	return (Boolean)(C.goglIsAsyncMarkerSGIX((C.GLuint)(marker)))
}
// SGIX_async_histogram

// SGIX_async_pixel

// SGIX_blend_alpha_minmax

// SGIX_calligraphic_fragment

// SGIX_clipmap

// SGIX_convolution_accuracy

// SGIX_depth_pass_instrument

// SGIX_depth_texture

// SGIX_flush_raster

func FlushRasterSGIX()  {
	C.goglFlushRasterSGIX()
}
// SGIX_fog_offset

// SGIX_fog_scale

// SGIX_fragment_lighting

func FragmentColorMaterialSGIX(face Enum, mode Enum)  {
	C.goglFragmentColorMaterialSGIX((C.GLenum)(face), (C.GLenum)(mode))
}
func FragmentLightfSGIX(light Enum, pname Enum, param Float)  {
	C.goglFragmentLightfSGIX((C.GLenum)(light), (C.GLenum)(pname), (C.GLfloat)(param))
}
func FragmentLightfvSGIX(light Enum, pname Enum, params *Float)  {
	C.goglFragmentLightfvSGIX((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func FragmentLightiSGIX(light Enum, pname Enum, param Int)  {
	C.goglFragmentLightiSGIX((C.GLenum)(light), (C.GLenum)(pname), (C.GLint)(param))
}
func FragmentLightivSGIX(light Enum, pname Enum, params *Int)  {
	C.goglFragmentLightivSGIX((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
func FragmentLightModelfSGIX(pname Enum, param Float)  {
	C.goglFragmentLightModelfSGIX((C.GLenum)(pname), (C.GLfloat)(param))
}
func FragmentLightModelfvSGIX(pname Enum, params *Float)  {
	C.goglFragmentLightModelfvSGIX((C.GLenum)(pname), (*C.GLfloat)(params))
}
func FragmentLightModeliSGIX(pname Enum, param Int)  {
	C.goglFragmentLightModeliSGIX((C.GLenum)(pname), (C.GLint)(param))
}
func FragmentLightModelivSGIX(pname Enum, params *Int)  {
	C.goglFragmentLightModelivSGIX((C.GLenum)(pname), (*C.GLint)(params))
}
func FragmentMaterialfSGIX(face Enum, pname Enum, param Float)  {
	C.goglFragmentMaterialfSGIX((C.GLenum)(face), (C.GLenum)(pname), (C.GLfloat)(param))
}
func FragmentMaterialfvSGIX(face Enum, pname Enum, params *Float)  {
	C.goglFragmentMaterialfvSGIX((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func FragmentMaterialiSGIX(face Enum, pname Enum, param Int)  {
	C.goglFragmentMaterialiSGIX((C.GLenum)(face), (C.GLenum)(pname), (C.GLint)(param))
}
func FragmentMaterialivSGIX(face Enum, pname Enum, params *Int)  {
	C.goglFragmentMaterialivSGIX((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetFragmentLightfvSGIX(light Enum, pname Enum, params *Float)  {
	C.goglGetFragmentLightfvSGIX((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetFragmentLightivSGIX(light Enum, pname Enum, params *Int)  {
	C.goglGetFragmentLightivSGIX((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetFragmentMaterialfvSGIX(face Enum, pname Enum, params *Float)  {
	C.goglGetFragmentMaterialfvSGIX((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetFragmentMaterialivSGIX(face Enum, pname Enum, params *Int)  {
	C.goglGetFragmentMaterialivSGIX((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
func LightEnviSGIX(pname Enum, param Int)  {
	C.goglLightEnviSGIX((C.GLenum)(pname), (C.GLint)(param))
}
// SGIX_framezoom

func FrameZoomSGIX(factor Int)  {
	C.goglFrameZoomSGIX((C.GLint)(factor))
}
// SGIX_igloo_interface

func IglooInterfaceSGIX(pname Enum, params Pointer)  {
	C.goglIglooInterfaceSGIX((C.GLenum)(pname), (unsafe.Pointer)(params))
}
// SGIX_instruments

func GetInstrumentsSGIX() Int {
	return (Int)(C.goglGetInstrumentsSGIX())
}
func InstrumentsBufferSGIX(size Sizei, buffer *Int)  {
	C.goglInstrumentsBufferSGIX((C.GLsizei)(size), (*C.GLint)(buffer))
}
func PollInstrumentsSGIX(marker_p *Int) Int {
	return (Int)(C.goglPollInstrumentsSGIX((*C.GLint)(marker_p)))
}
func ReadInstrumentsSGIX(marker Int)  {
	C.goglReadInstrumentsSGIX((C.GLint)(marker))
}
func StartInstrumentsSGIX()  {
	C.goglStartInstrumentsSGIX()
}
func StopInstrumentsSGIX(marker Int)  {
	C.goglStopInstrumentsSGIX((C.GLint)(marker))
}
// SGIX_interlace

// SGIX_ir_instrument1

// SGIX_list_priority

func GetListParameterfvSGIX(list Uint, pname Enum, params *Float)  {
	C.goglGetListParameterfvSGIX((C.GLuint)(list), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetListParameterivSGIX(list Uint, pname Enum, params *Int)  {
	C.goglGetListParameterivSGIX((C.GLuint)(list), (C.GLenum)(pname), (*C.GLint)(params))
}
func ListParameterfSGIX(list Uint, pname Enum, param Float)  {
	C.goglListParameterfSGIX((C.GLuint)(list), (C.GLenum)(pname), (C.GLfloat)(param))
}
func ListParameterfvSGIX(list Uint, pname Enum, params *Float)  {
	C.goglListParameterfvSGIX((C.GLuint)(list), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func ListParameteriSGIX(list Uint, pname Enum, param Int)  {
	C.goglListParameteriSGIX((C.GLuint)(list), (C.GLenum)(pname), (C.GLint)(param))
}
func ListParameterivSGIX(list Uint, pname Enum, params *Int)  {
	C.goglListParameterivSGIX((C.GLuint)(list), (C.GLenum)(pname), (*C.GLint)(params))
}
// SGIX_pixel_texture

func PixelTexGenSGIX(mode Enum)  {
	C.goglPixelTexGenSGIX((C.GLenum)(mode))
}
// SGIX_pixel_tiles

// SGIX_polynomial_ffd

func DeformationMap3dSGIX(target Enum, u1 Double, u2 Double, ustride Int, uorder Int, v1 Double, v2 Double, vstride Int, vorder Int, w1 Double, w2 Double, wstride Int, worder Int, points *Double)  {
	C.goglDeformationMap3dSGIX((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (C.GLdouble)(w1), (C.GLdouble)(w2), (C.GLint)(wstride), (C.GLint)(worder), (*C.GLdouble)(points))
}
func DeformationMap3fSGIX(target Enum, u1 Float, u2 Float, ustride Int, uorder Int, v1 Float, v2 Float, vstride Int, vorder Int, w1 Float, w2 Float, wstride Int, worder Int, points *Float)  {
	C.goglDeformationMap3fSGIX((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (C.GLfloat)(w1), (C.GLfloat)(w2), (C.GLint)(wstride), (C.GLint)(worder), (*C.GLfloat)(points))
}
func DeformSGIX(mask Bitfield)  {
	C.goglDeformSGIX((C.GLbitfield)(mask))
}
func LoadIdentityDeformationMapSGIX(mask Bitfield)  {
	C.goglLoadIdentityDeformationMapSGIX((C.GLbitfield)(mask))
}
// SGIX_reference_plane

func ReferencePlaneSGIX(equation *Double)  {
	C.goglReferencePlaneSGIX((*C.GLdouble)(equation))
}
// SGIX_resample

// SGIX_scalebias_hint

// SGIX_shadow

// SGIX_shadow_ambient

// SGIX_sprite

func SpriteParameterfSGIX(pname Enum, param Float)  {
	C.goglSpriteParameterfSGIX((C.GLenum)(pname), (C.GLfloat)(param))
}
func SpriteParameterfvSGIX(pname Enum, params *Float)  {
	C.goglSpriteParameterfvSGIX((C.GLenum)(pname), (*C.GLfloat)(params))
}
func SpriteParameteriSGIX(pname Enum, param Int)  {
	C.goglSpriteParameteriSGIX((C.GLenum)(pname), (C.GLint)(param))
}
func SpriteParameterivSGIX(pname Enum, params *Int)  {
	C.goglSpriteParameterivSGIX((C.GLenum)(pname), (*C.GLint)(params))
}
// SGIX_subsample

// SGIX_tag_sample_buffer

func TagSampleBufferSGIX()  {
	C.goglTagSampleBufferSGIX()
}
// SGIX_texture_add_env

// SGIX_texture_coordinate_clamp

// SGIX_texture_lod_bias

// SGIX_texture_multi_buffer

// SGIX_texture_scale_bias

// SGIX_texture_select

// SGIX_vertex_preclip

// SGIX_ycrcb

// SGIX_ycrcb_subsample

// SGIX_ycrcba

func InitSgixAsync() error {
	var ret C.int
	if ret = C.init_SGIX_async(); ret != 0 {
		return errors.New("unable to initialize SGIX_async")
	}
	return nil
}
func InitSgixAsyncHistogram() error {
	var ret C.int
	if ret = C.init_SGIX_async_histogram(); ret != 0 {
		return errors.New("unable to initialize SGIX_async_histogram")
	}
	return nil
}
func InitSgixAsyncPixel() error {
	var ret C.int
	if ret = C.init_SGIX_async_pixel(); ret != 0 {
		return errors.New("unable to initialize SGIX_async_pixel")
	}
	return nil
}
func InitSgixBlendAlphaMinmax() error {
	var ret C.int
	if ret = C.init_SGIX_blend_alpha_minmax(); ret != 0 {
		return errors.New("unable to initialize SGIX_blend_alpha_minmax")
	}
	return nil
}
func InitSgixCalligraphicFragment() error {
	var ret C.int
	if ret = C.init_SGIX_calligraphic_fragment(); ret != 0 {
		return errors.New("unable to initialize SGIX_calligraphic_fragment")
	}
	return nil
}
func InitSgixClipmap() error {
	var ret C.int
	if ret = C.init_SGIX_clipmap(); ret != 0 {
		return errors.New("unable to initialize SGIX_clipmap")
	}
	return nil
}
func InitSgixConvolutionAccuracy() error {
	var ret C.int
	if ret = C.init_SGIX_convolution_accuracy(); ret != 0 {
		return errors.New("unable to initialize SGIX_convolution_accuracy")
	}
	return nil
}
func InitSgixDepthPassInstrument() error {
	var ret C.int
	if ret = C.init_SGIX_depth_pass_instrument(); ret != 0 {
		return errors.New("unable to initialize SGIX_depth_pass_instrument")
	}
	return nil
}
func InitSgixDepthTexture() error {
	var ret C.int
	if ret = C.init_SGIX_depth_texture(); ret != 0 {
		return errors.New("unable to initialize SGIX_depth_texture")
	}
	return nil
}
func InitSgixFlushRaster() error {
	var ret C.int
	if ret = C.init_SGIX_flush_raster(); ret != 0 {
		return errors.New("unable to initialize SGIX_flush_raster")
	}
	return nil
}
func InitSgixFogOffset() error {
	var ret C.int
	if ret = C.init_SGIX_fog_offset(); ret != 0 {
		return errors.New("unable to initialize SGIX_fog_offset")
	}
	return nil
}
func InitSgixFogScale() error {
	var ret C.int
	if ret = C.init_SGIX_fog_scale(); ret != 0 {
		return errors.New("unable to initialize SGIX_fog_scale")
	}
	return nil
}
func InitSgixFragmentLighting() error {
	var ret C.int
	if ret = C.init_SGIX_fragment_lighting(); ret != 0 {
		return errors.New("unable to initialize SGIX_fragment_lighting")
	}
	return nil
}
func InitSgixFramezoom() error {
	var ret C.int
	if ret = C.init_SGIX_framezoom(); ret != 0 {
		return errors.New("unable to initialize SGIX_framezoom")
	}
	return nil
}
func InitSgixIglooInterface() error {
	var ret C.int
	if ret = C.init_SGIX_igloo_interface(); ret != 0 {
		return errors.New("unable to initialize SGIX_igloo_interface")
	}
	return nil
}
func InitSgixInstruments() error {
	var ret C.int
	if ret = C.init_SGIX_instruments(); ret != 0 {
		return errors.New("unable to initialize SGIX_instruments")
	}
	return nil
}
func InitSgixInterlace() error {
	var ret C.int
	if ret = C.init_SGIX_interlace(); ret != 0 {
		return errors.New("unable to initialize SGIX_interlace")
	}
	return nil
}
func InitSgixIrInstrument1() error {
	var ret C.int
	if ret = C.init_SGIX_ir_instrument1(); ret != 0 {
		return errors.New("unable to initialize SGIX_ir_instrument1")
	}
	return nil
}
func InitSgixListPriority() error {
	var ret C.int
	if ret = C.init_SGIX_list_priority(); ret != 0 {
		return errors.New("unable to initialize SGIX_list_priority")
	}
	return nil
}
func InitSgixPixelTexture() error {
	var ret C.int
	if ret = C.init_SGIX_pixel_texture(); ret != 0 {
		return errors.New("unable to initialize SGIX_pixel_texture")
	}
	return nil
}
func InitSgixPixelTiles() error {
	var ret C.int
	if ret = C.init_SGIX_pixel_tiles(); ret != 0 {
		return errors.New("unable to initialize SGIX_pixel_tiles")
	}
	return nil
}
func InitSgixPolynomialFfd() error {
	var ret C.int
	if ret = C.init_SGIX_polynomial_ffd(); ret != 0 {
		return errors.New("unable to initialize SGIX_polynomial_ffd")
	}
	return nil
}
func InitSgixReferencePlane() error {
	var ret C.int
	if ret = C.init_SGIX_reference_plane(); ret != 0 {
		return errors.New("unable to initialize SGIX_reference_plane")
	}
	return nil
}
func InitSgixResample() error {
	var ret C.int
	if ret = C.init_SGIX_resample(); ret != 0 {
		return errors.New("unable to initialize SGIX_resample")
	}
	return nil
}
func InitSgixScalebiasHint() error {
	var ret C.int
	if ret = C.init_SGIX_scalebias_hint(); ret != 0 {
		return errors.New("unable to initialize SGIX_scalebias_hint")
	}
	return nil
}
func InitSgixShadow() error {
	var ret C.int
	if ret = C.init_SGIX_shadow(); ret != 0 {
		return errors.New("unable to initialize SGIX_shadow")
	}
	return nil
}
func InitSgixShadowAmbient() error {
	var ret C.int
	if ret = C.init_SGIX_shadow_ambient(); ret != 0 {
		return errors.New("unable to initialize SGIX_shadow_ambient")
	}
	return nil
}
func InitSgixSprite() error {
	var ret C.int
	if ret = C.init_SGIX_sprite(); ret != 0 {
		return errors.New("unable to initialize SGIX_sprite")
	}
	return nil
}
func InitSgixSubsample() error {
	var ret C.int
	if ret = C.init_SGIX_subsample(); ret != 0 {
		return errors.New("unable to initialize SGIX_subsample")
	}
	return nil
}
func InitSgixTagSampleBuffer() error {
	var ret C.int
	if ret = C.init_SGIX_tag_sample_buffer(); ret != 0 {
		return errors.New("unable to initialize SGIX_tag_sample_buffer")
	}
	return nil
}
func InitSgixTextureAddEnv() error {
	var ret C.int
	if ret = C.init_SGIX_texture_add_env(); ret != 0 {
		return errors.New("unable to initialize SGIX_texture_add_env")
	}
	return nil
}
func InitSgixTextureCoordinateClamp() error {
	var ret C.int
	if ret = C.init_SGIX_texture_coordinate_clamp(); ret != 0 {
		return errors.New("unable to initialize SGIX_texture_coordinate_clamp")
	}
	return nil
}
func InitSgixTextureLodBias() error {
	var ret C.int
	if ret = C.init_SGIX_texture_lod_bias(); ret != 0 {
		return errors.New("unable to initialize SGIX_texture_lod_bias")
	}
	return nil
}
func InitSgixTextureMultiBuffer() error {
	var ret C.int
	if ret = C.init_SGIX_texture_multi_buffer(); ret != 0 {
		return errors.New("unable to initialize SGIX_texture_multi_buffer")
	}
	return nil
}
func InitSgixTextureScaleBias() error {
	var ret C.int
	if ret = C.init_SGIX_texture_scale_bias(); ret != 0 {
		return errors.New("unable to initialize SGIX_texture_scale_bias")
	}
	return nil
}
func InitSgixTextureSelect() error {
	var ret C.int
	if ret = C.init_SGIX_texture_select(); ret != 0 {
		return errors.New("unable to initialize SGIX_texture_select")
	}
	return nil
}
func InitSgixVertexPreclip() error {
	var ret C.int
	if ret = C.init_SGIX_vertex_preclip(); ret != 0 {
		return errors.New("unable to initialize SGIX_vertex_preclip")
	}
	return nil
}
func InitSgixYcrcb() error {
	var ret C.int
	if ret = C.init_SGIX_ycrcb(); ret != 0 {
		return errors.New("unable to initialize SGIX_ycrcb")
	}
	return nil
}
func InitSgixYcrcbSubsample() error {
	var ret C.int
	if ret = C.init_SGIX_ycrcb_subsample(); ret != 0 {
		return errors.New("unable to initialize SGIX_ycrcb_subsample")
	}
	return nil
}
func InitSgixYcrcba() error {
	var ret C.int
	if ret = C.init_SGIX_ycrcba(); ret != 0 {
		return errors.New("unable to initialize SGIX_ycrcba")
	}
	return nil
}
// EOF