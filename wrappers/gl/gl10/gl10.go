// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// VERSION_1_0
// 
// VERSION_1_0_DEPRECATED
// 
// http://www.opengl.org/sdk/docs/man
// 
package gl10

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
// //  VERSION_1_0
// void (APIENTRYP ptrglCullFace)(GLenum mode);
// void (APIENTRYP ptrglFrontFace)(GLenum mode);
// void (APIENTRYP ptrglHint)(GLenum target, GLenum mode);
// void (APIENTRYP ptrglLineWidth)(GLfloat width);
// void (APIENTRYP ptrglPointSize)(GLfloat size);
// void (APIENTRYP ptrglPolygonMode)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglScissor)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglTexParameterf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglTexImage1D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglTexImage2D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglDrawBuffer)(GLenum mode);
// void (APIENTRYP ptrglClear)(GLbitfield mask);
// void (APIENTRYP ptrglClearColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglClearStencil)(GLint s);
// void (APIENTRYP ptrglClearDepth)(GLdouble depth);
// void (APIENTRYP ptrglStencilMask)(GLuint mask);
// void (APIENTRYP ptrglColorMask)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
// void (APIENTRYP ptrglDepthMask)(GLboolean flag);
// void (APIENTRYP ptrglDisable)(GLenum cap);
// void (APIENTRYP ptrglEnable)(GLenum cap);
// void (APIENTRYP ptrglFinish)();
// void (APIENTRYP ptrglFlush)();
// void (APIENTRYP ptrglBlendFunc)(GLenum sfactor, GLenum dfactor);
// void (APIENTRYP ptrglLogicOp)(GLenum opcode);
// void (APIENTRYP ptrglStencilFunc)(GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglStencilOp)(GLenum fail, GLenum zfail, GLenum zpass);
// void (APIENTRYP ptrglDepthFunc)(GLenum func);
// void (APIENTRYP ptrglPixelStoref)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPixelStorei)(GLenum pname, GLint param);
// void (APIENTRYP ptrglReadBuffer)(GLenum mode);
// void (APIENTRYP ptrglReadPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGetBooleanv)(GLenum pname, GLboolean* params);
// void (APIENTRYP ptrglGetDoublev)(GLenum pname, GLdouble* params);
// GLenum (APIENTRYP ptrglGetError)();
// void (APIENTRYP ptrglGetFloatv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetIntegerv)(GLenum pname, GLint* params);
// const GLubyte * (APIENTRYP ptrglGetString)(GLenum name);
// void (APIENTRYP ptrglGetTexImage)(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGetTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetTexLevelParameterfv)(GLenum target, GLint level, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexLevelParameteriv)(GLenum target, GLint level, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrglIsEnabled)(GLenum cap);
// void (APIENTRYP ptrglDepthRange)(GLdouble near, GLdouble far);
// void (APIENTRYP ptrglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
// //  VERSION_1_0_DEPRECATED
// void (APIENTRYP ptrglNewList)(GLuint list, GLenum mode);
// void (APIENTRYP ptrglEndList)();
// void (APIENTRYP ptrglCallList)(GLuint list);
// void (APIENTRYP ptrglCallLists)(GLsizei n, GLenum type, GLvoid* lists);
// void (APIENTRYP ptrglDeleteLists)(GLuint list, GLsizei range);
// GLuint (APIENTRYP ptrglGenLists)(GLsizei range);
// void (APIENTRYP ptrglListBase)(GLuint base);
// void (APIENTRYP ptrglBegin)(GLenum mode);
// void (APIENTRYP ptrglBitmap)(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap);
// void (APIENTRYP ptrglColor3b)(GLbyte red, GLbyte green, GLbyte blue);
// void (APIENTRYP ptrglColor3bv)(GLbyte* v);
// void (APIENTRYP ptrglColor3d)(GLdouble red, GLdouble green, GLdouble blue);
// void (APIENTRYP ptrglColor3dv)(GLdouble* v);
// void (APIENTRYP ptrglColor3f)(GLfloat red, GLfloat green, GLfloat blue);
// void (APIENTRYP ptrglColor3fv)(GLfloat* v);
// void (APIENTRYP ptrglColor3i)(GLint red, GLint green, GLint blue);
// void (APIENTRYP ptrglColor3iv)(GLint* v);
// void (APIENTRYP ptrglColor3s)(GLshort red, GLshort green, GLshort blue);
// void (APIENTRYP ptrglColor3sv)(GLshort* v);
// void (APIENTRYP ptrglColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
// void (APIENTRYP ptrglColor3ubv)(GLubyte* v);
// void (APIENTRYP ptrglColor3ui)(GLuint red, GLuint green, GLuint blue);
// void (APIENTRYP ptrglColor3uiv)(GLuint* v);
// void (APIENTRYP ptrglColor3us)(GLushort red, GLushort green, GLushort blue);
// void (APIENTRYP ptrglColor3usv)(GLushort* v);
// void (APIENTRYP ptrglColor4b)(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha);
// void (APIENTRYP ptrglColor4bv)(GLbyte* v);
// void (APIENTRYP ptrglColor4d)(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha);
// void (APIENTRYP ptrglColor4dv)(GLdouble* v);
// void (APIENTRYP ptrglColor4f)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglColor4fv)(GLfloat* v);
// void (APIENTRYP ptrglColor4i)(GLint red, GLint green, GLint blue, GLint alpha);
// void (APIENTRYP ptrglColor4iv)(GLint* v);
// void (APIENTRYP ptrglColor4s)(GLshort red, GLshort green, GLshort blue, GLshort alpha);
// void (APIENTRYP ptrglColor4sv)(GLshort* v);
// void (APIENTRYP ptrglColor4ub)(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha);
// void (APIENTRYP ptrglColor4ubv)(GLubyte* v);
// void (APIENTRYP ptrglColor4ui)(GLuint red, GLuint green, GLuint blue, GLuint alpha);
// void (APIENTRYP ptrglColor4uiv)(GLuint* v);
// void (APIENTRYP ptrglColor4us)(GLushort red, GLushort green, GLushort blue, GLushort alpha);
// void (APIENTRYP ptrglColor4usv)(GLushort* v);
// void (APIENTRYP ptrglEdgeFlag)(GLboolean flag);
// void (APIENTRYP ptrglEdgeFlagv)(GLboolean* flag);
// void (APIENTRYP ptrglEnd)();
// void (APIENTRYP ptrglIndexd)(GLdouble c);
// void (APIENTRYP ptrglIndexdv)(GLdouble* c);
// void (APIENTRYP ptrglIndexf)(GLfloat c);
// void (APIENTRYP ptrglIndexfv)(GLfloat* c);
// void (APIENTRYP ptrglIndexi)(GLint c);
// void (APIENTRYP ptrglIndexiv)(GLint* c);
// void (APIENTRYP ptrglIndexs)(GLshort c);
// void (APIENTRYP ptrglIndexsv)(GLshort* c);
// void (APIENTRYP ptrglNormal3b)(GLbyte nx, GLbyte ny, GLbyte nz);
// void (APIENTRYP ptrglNormal3bv)(GLbyte* v);
// void (APIENTRYP ptrglNormal3d)(GLdouble nx, GLdouble ny, GLdouble nz);
// void (APIENTRYP ptrglNormal3dv)(GLdouble* v);
// void (APIENTRYP ptrglNormal3f)(GLfloat nx, GLfloat ny, GLfloat nz);
// void (APIENTRYP ptrglNormal3fv)(GLfloat* v);
// void (APIENTRYP ptrglNormal3i)(GLint nx, GLint ny, GLint nz);
// void (APIENTRYP ptrglNormal3iv)(GLint* v);
// void (APIENTRYP ptrglNormal3s)(GLshort nx, GLshort ny, GLshort nz);
// void (APIENTRYP ptrglNormal3sv)(GLshort* v);
// void (APIENTRYP ptrglRasterPos2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrglRasterPos2dv)(GLdouble* v);
// void (APIENTRYP ptrglRasterPos2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrglRasterPos2fv)(GLfloat* v);
// void (APIENTRYP ptrglRasterPos2i)(GLint x, GLint y);
// void (APIENTRYP ptrglRasterPos2iv)(GLint* v);
// void (APIENTRYP ptrglRasterPos2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrglRasterPos2sv)(GLshort* v);
// void (APIENTRYP ptrglRasterPos3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglRasterPos3dv)(GLdouble* v);
// void (APIENTRYP ptrglRasterPos3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglRasterPos3fv)(GLfloat* v);
// void (APIENTRYP ptrglRasterPos3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglRasterPos3iv)(GLint* v);
// void (APIENTRYP ptrglRasterPos3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglRasterPos3sv)(GLshort* v);
// void (APIENTRYP ptrglRasterPos4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglRasterPos4dv)(GLdouble* v);
// void (APIENTRYP ptrglRasterPos4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglRasterPos4fv)(GLfloat* v);
// void (APIENTRYP ptrglRasterPos4i)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglRasterPos4iv)(GLint* v);
// void (APIENTRYP ptrglRasterPos4s)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglRasterPos4sv)(GLshort* v);
// void (APIENTRYP ptrglRectd)(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2);
// void (APIENTRYP ptrglRectdv)(GLdouble* v1, GLdouble* v2);
// void (APIENTRYP ptrglRectf)(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2);
// void (APIENTRYP ptrglRectfv)(GLfloat* v1, GLfloat* v2);
// void (APIENTRYP ptrglRecti)(GLint x1, GLint y1, GLint x2, GLint y2);
// void (APIENTRYP ptrglRectiv)(GLint* v1, GLint* v2);
// void (APIENTRYP ptrglRects)(GLshort x1, GLshort y1, GLshort x2, GLshort y2);
// void (APIENTRYP ptrglRectsv)(GLshort* v1, GLshort* v2);
// void (APIENTRYP ptrglTexCoord1d)(GLdouble s);
// void (APIENTRYP ptrglTexCoord1dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord1f)(GLfloat s);
// void (APIENTRYP ptrglTexCoord1fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord1i)(GLint s);
// void (APIENTRYP ptrglTexCoord1iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord1s)(GLshort s);
// void (APIENTRYP ptrglTexCoord1sv)(GLshort* v);
// void (APIENTRYP ptrglTexCoord2d)(GLdouble s, GLdouble t);
// void (APIENTRYP ptrglTexCoord2dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord2f)(GLfloat s, GLfloat t);
// void (APIENTRYP ptrglTexCoord2fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord2i)(GLint s, GLint t);
// void (APIENTRYP ptrglTexCoord2iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord2s)(GLshort s, GLshort t);
// void (APIENTRYP ptrglTexCoord2sv)(GLshort* v);
// void (APIENTRYP ptrglTexCoord3d)(GLdouble s, GLdouble t, GLdouble r);
// void (APIENTRYP ptrglTexCoord3dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord3f)(GLfloat s, GLfloat t, GLfloat r);
// void (APIENTRYP ptrglTexCoord3fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord3i)(GLint s, GLint t, GLint r);
// void (APIENTRYP ptrglTexCoord3iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord3s)(GLshort s, GLshort t, GLshort r);
// void (APIENTRYP ptrglTexCoord3sv)(GLshort* v);
// void (APIENTRYP ptrglTexCoord4d)(GLdouble s, GLdouble t, GLdouble r, GLdouble q);
// void (APIENTRYP ptrglTexCoord4dv)(GLdouble* v);
// void (APIENTRYP ptrglTexCoord4f)(GLfloat s, GLfloat t, GLfloat r, GLfloat q);
// void (APIENTRYP ptrglTexCoord4fv)(GLfloat* v);
// void (APIENTRYP ptrglTexCoord4i)(GLint s, GLint t, GLint r, GLint q);
// void (APIENTRYP ptrglTexCoord4iv)(GLint* v);
// void (APIENTRYP ptrglTexCoord4s)(GLshort s, GLshort t, GLshort r, GLshort q);
// void (APIENTRYP ptrglTexCoord4sv)(GLshort* v);
// void (APIENTRYP ptrglVertex2d)(GLdouble x, GLdouble y);
// void (APIENTRYP ptrglVertex2dv)(GLdouble* v);
// void (APIENTRYP ptrglVertex2f)(GLfloat x, GLfloat y);
// void (APIENTRYP ptrglVertex2fv)(GLfloat* v);
// void (APIENTRYP ptrglVertex2i)(GLint x, GLint y);
// void (APIENTRYP ptrglVertex2iv)(GLint* v);
// void (APIENTRYP ptrglVertex2s)(GLshort x, GLshort y);
// void (APIENTRYP ptrglVertex2sv)(GLshort* v);
// void (APIENTRYP ptrglVertex3d)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglVertex3dv)(GLdouble* v);
// void (APIENTRYP ptrglVertex3f)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglVertex3fv)(GLfloat* v);
// void (APIENTRYP ptrglVertex3i)(GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglVertex3iv)(GLint* v);
// void (APIENTRYP ptrglVertex3s)(GLshort x, GLshort y, GLshort z);
// void (APIENTRYP ptrglVertex3sv)(GLshort* v);
// void (APIENTRYP ptrglVertex4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertex4dv)(GLdouble* v);
// void (APIENTRYP ptrglVertex4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
// void (APIENTRYP ptrglVertex4fv)(GLfloat* v);
// void (APIENTRYP ptrglVertex4i)(GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglVertex4iv)(GLint* v);
// void (APIENTRYP ptrglVertex4s)(GLshort x, GLshort y, GLshort z, GLshort w);
// void (APIENTRYP ptrglVertex4sv)(GLshort* v);
// void (APIENTRYP ptrglClipPlane)(GLenum plane, GLdouble* equation);
// void (APIENTRYP ptrglColorMaterial)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglFogf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglFogfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglFogi)(GLenum pname, GLint param);
// void (APIENTRYP ptrglFogiv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglLightf)(GLenum light, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglLightfv)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglLighti)(GLenum light, GLenum pname, GLint param);
// void (APIENTRYP ptrglLightiv)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrglLightModelf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglLightModelfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglLightModeli)(GLenum pname, GLint param);
// void (APIENTRYP ptrglLightModeliv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglLineStipple)(GLint factor, GLushort pattern);
// void (APIENTRYP ptrglMaterialf)(GLenum face, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglMateriali)(GLenum face, GLenum pname, GLint param);
// void (APIENTRYP ptrglMaterialiv)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrglPolygonStipple)(GLubyte* mask);
// void (APIENTRYP ptrglShadeModel)(GLenum mode);
// void (APIENTRYP ptrglTexEnvf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexEnvi)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexEnviv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglTexGend)(GLenum coord, GLenum pname, GLdouble param);
// void (APIENTRYP ptrglTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglTexGenf)(GLenum coord, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexGeni)(GLenum coord, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexGeniv)(GLenum coord, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFeedbackBuffer)(GLsizei size, GLenum type, GLfloat* buffer);
// void (APIENTRYP ptrglSelectBuffer)(GLsizei size, GLuint* buffer);
// GLint (APIENTRYP ptrglRenderMode)(GLenum mode);
// void (APIENTRYP ptrglInitNames)();
// void (APIENTRYP ptrglLoadName)(GLuint name);
// void (APIENTRYP ptrglPassThrough)(GLfloat token);
// void (APIENTRYP ptrglPopName)();
// void (APIENTRYP ptrglPushName)(GLuint name);
// void (APIENTRYP ptrglClearAccum)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglClearIndex)(GLfloat c);
// void (APIENTRYP ptrglIndexMask)(GLuint mask);
// void (APIENTRYP ptrglAccum)(GLenum op, GLfloat value);
// void (APIENTRYP ptrglPopAttrib)();
// void (APIENTRYP ptrglPushAttrib)(GLbitfield mask);
// void (APIENTRYP ptrglMap1d)(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points);
// void (APIENTRYP ptrglMap1f)(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points);
// void (APIENTRYP ptrglMap2d)(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points);
// void (APIENTRYP ptrglMap2f)(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points);
// void (APIENTRYP ptrglMapGrid1d)(GLint un, GLdouble u1, GLdouble u2);
// void (APIENTRYP ptrglMapGrid1f)(GLint un, GLfloat u1, GLfloat u2);
// void (APIENTRYP ptrglMapGrid2d)(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2);
// void (APIENTRYP ptrglMapGrid2f)(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglEvalCoord1d)(GLdouble u);
// void (APIENTRYP ptrglEvalCoord1dv)(GLdouble* u);
// void (APIENTRYP ptrglEvalCoord1f)(GLfloat u);
// void (APIENTRYP ptrglEvalCoord1fv)(GLfloat* u);
// void (APIENTRYP ptrglEvalCoord2d)(GLdouble u, GLdouble v);
// void (APIENTRYP ptrglEvalCoord2dv)(GLdouble* u);
// void (APIENTRYP ptrglEvalCoord2f)(GLfloat u, GLfloat v);
// void (APIENTRYP ptrglEvalCoord2fv)(GLfloat* u);
// void (APIENTRYP ptrglEvalMesh1)(GLenum mode, GLint i1, GLint i2);
// void (APIENTRYP ptrglEvalPoint1)(GLint i);
// void (APIENTRYP ptrglEvalMesh2)(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2);
// void (APIENTRYP ptrglEvalPoint2)(GLint i, GLint j);
// void (APIENTRYP ptrglAlphaFunc)(GLenum func, GLfloat ref);
// void (APIENTRYP ptrglPixelZoom)(GLfloat xfactor, GLfloat yfactor);
// void (APIENTRYP ptrglPixelTransferf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglPixelTransferi)(GLenum pname, GLint param);
// void (APIENTRYP ptrglPixelMapfv)(GLenum map, GLsizei mapsize, GLfloat* values);
// void (APIENTRYP ptrglPixelMapuiv)(GLenum map, GLsizei mapsize, GLuint* values);
// void (APIENTRYP ptrglPixelMapusv)(GLenum map, GLsizei mapsize, GLushort* values);
// void (APIENTRYP ptrglCopyPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type);
// void (APIENTRYP ptrglDrawPixels)(GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGetClipPlane)(GLenum plane, GLdouble* equation);
// void (APIENTRYP ptrglGetLightfv)(GLenum light, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetLightiv)(GLenum light, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetMapdv)(GLenum target, GLenum query, GLdouble* v);
// void (APIENTRYP ptrglGetMapfv)(GLenum target, GLenum query, GLfloat* v);
// void (APIENTRYP ptrglGetMapiv)(GLenum target, GLenum query, GLint* v);
// void (APIENTRYP ptrglGetMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetMaterialiv)(GLenum face, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetPixelMapfv)(GLenum map, GLfloat* values);
// void (APIENTRYP ptrglGetPixelMapuiv)(GLenum map, GLuint* values);
// void (APIENTRYP ptrglGetPixelMapusv)(GLenum map, GLushort* values);
// void (APIENTRYP ptrglGetPolygonStipple)(GLubyte* mask);
// void (APIENTRYP ptrglGetTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexEnviv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglGetTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetTexGeniv)(GLenum coord, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrglIsList)(GLuint list);
// void (APIENTRYP ptrglFrustum)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrglLoadIdentity)();
// void (APIENTRYP ptrglLoadMatrixf)(GLfloat* m);
// void (APIENTRYP ptrglLoadMatrixd)(GLdouble* m);
// void (APIENTRYP ptrglMatrixMode)(GLenum mode);
// void (APIENTRYP ptrglMultMatrixf)(GLfloat* m);
// void (APIENTRYP ptrglMultMatrixd)(GLdouble* m);
// void (APIENTRYP ptrglOrtho)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
// void (APIENTRYP ptrglPopMatrix)();
// void (APIENTRYP ptrglPushMatrix)();
// void (APIENTRYP ptrglRotated)(GLdouble angle, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglRotatef)(GLfloat angle, GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglScaled)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglScalef)(GLfloat x, GLfloat y, GLfloat z);
// void (APIENTRYP ptrglTranslated)(GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglTranslatef)(GLfloat x, GLfloat y, GLfloat z);
// 
// //  VERSION_1_0
// void goglCullFace(GLenum mode) {
// 	(*ptrglCullFace)(mode);
// }
// void goglFrontFace(GLenum mode) {
// 	(*ptrglFrontFace)(mode);
// }
// void goglHint(GLenum target, GLenum mode) {
// 	(*ptrglHint)(target, mode);
// }
// void goglLineWidth(GLfloat width) {
// 	(*ptrglLineWidth)(width);
// }
// void goglPointSize(GLfloat size) {
// 	(*ptrglPointSize)(size);
// }
// void goglPolygonMode(GLenum face, GLenum mode) {
// 	(*ptrglPolygonMode)(face, mode);
// }
// void goglScissor(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglScissor)(x, y, width, height);
// }
// void goglTexParameterf(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrglTexParameterf)(target, pname, param);
// }
// void goglTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglTexParameterfv)(target, pname, params);
// }
// void goglTexParameteri(GLenum target, GLenum pname, GLint param) {
// 	(*ptrglTexParameteri)(target, pname, param);
// }
// void goglTexParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglTexParameteriv)(target, pname, params);
// }
// void goglTexImage1D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexImage1D)(target, level, internalformat, width, border, format, type_, pixels);
// }
// void goglTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglTexImage2D)(target, level, internalformat, width, height, border, format, type_, pixels);
// }
// void goglDrawBuffer(GLenum mode) {
// 	(*ptrglDrawBuffer)(mode);
// }
// void goglClear(GLbitfield mask) {
// 	(*ptrglClear)(mask);
// }
// void goglClearColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrglClearColor)(red, green, blue, alpha);
// }
// void goglClearStencil(GLint s) {
// 	(*ptrglClearStencil)(s);
// }
// void goglClearDepth(GLdouble depth) {
// 	(*ptrglClearDepth)(depth);
// }
// void goglStencilMask(GLuint mask) {
// 	(*ptrglStencilMask)(mask);
// }
// void goglColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
// 	(*ptrglColorMask)(red, green, blue, alpha);
// }
// void goglDepthMask(GLboolean flag) {
// 	(*ptrglDepthMask)(flag);
// }
// void goglDisable(GLenum cap) {
// 	(*ptrglDisable)(cap);
// }
// void goglEnable(GLenum cap) {
// 	(*ptrglEnable)(cap);
// }
// void goglFinish() {
// 	(*ptrglFinish)();
// }
// void goglFlush() {
// 	(*ptrglFlush)();
// }
// void goglBlendFunc(GLenum sfactor, GLenum dfactor) {
// 	(*ptrglBlendFunc)(sfactor, dfactor);
// }
// void goglLogicOp(GLenum opcode) {
// 	(*ptrglLogicOp)(opcode);
// }
// void goglStencilFunc(GLenum func_, GLint ref, GLuint mask) {
// 	(*ptrglStencilFunc)(func_, ref, mask);
// }
// void goglStencilOp(GLenum fail, GLenum zfail, GLenum zpass) {
// 	(*ptrglStencilOp)(fail, zfail, zpass);
// }
// void goglDepthFunc(GLenum func_) {
// 	(*ptrglDepthFunc)(func_);
// }
// void goglPixelStoref(GLenum pname, GLfloat param) {
// 	(*ptrglPixelStoref)(pname, param);
// }
// void goglPixelStorei(GLenum pname, GLint param) {
// 	(*ptrglPixelStorei)(pname, param);
// }
// void goglReadBuffer(GLenum mode) {
// 	(*ptrglReadBuffer)(mode);
// }
// void goglReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglReadPixels)(x, y, width, height, format, type_, pixels);
// }
// void goglGetBooleanv(GLenum pname, GLboolean* params) {
// 	(*ptrglGetBooleanv)(pname, params);
// }
// void goglGetDoublev(GLenum pname, GLdouble* params) {
// 	(*ptrglGetDoublev)(pname, params);
// }
// GLenum goglGetError() {
// 	return (*ptrglGetError)();
// }
// void goglGetFloatv(GLenum pname, GLfloat* params) {
// 	(*ptrglGetFloatv)(pname, params);
// }
// void goglGetIntegerv(GLenum pname, GLint* params) {
// 	(*ptrglGetIntegerv)(pname, params);
// }
// const GLubyte * goglGetString(GLenum name) {
// 	return (*ptrglGetString)(name);
// }
// void goglGetTexImage(GLenum target, GLint level, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglGetTexImage)(target, level, format, type_, pixels);
// }
// void goglGetTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexParameterfv)(target, pname, params);
// }
// void goglGetTexParameteriv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetTexParameteriv)(target, pname, params);
// }
// void goglGetTexLevelParameterfv(GLenum target, GLint level, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexLevelParameterfv)(target, level, pname, params);
// }
// void goglGetTexLevelParameteriv(GLenum target, GLint level, GLenum pname, GLint* params) {
// 	(*ptrglGetTexLevelParameteriv)(target, level, pname, params);
// }
// GLboolean goglIsEnabled(GLenum cap) {
// 	return (*ptrglIsEnabled)(cap);
// }
// void goglDepthRange(GLdouble near_, GLdouble far_) {
// 	(*ptrglDepthRange)(near_, far_);
// }
// void goglViewport(GLint x, GLint y, GLsizei width, GLsizei height) {
// 	(*ptrglViewport)(x, y, width, height);
// }
// //  VERSION_1_0_DEPRECATED
// void goglNewList(GLuint list, GLenum mode) {
// 	(*ptrglNewList)(list, mode);
// }
// void goglEndList() {
// 	(*ptrglEndList)();
// }
// void goglCallList(GLuint list) {
// 	(*ptrglCallList)(list);
// }
// void goglCallLists(GLsizei n, GLenum type_, GLvoid* lists) {
// 	(*ptrglCallLists)(n, type_, lists);
// }
// void goglDeleteLists(GLuint list, GLsizei range_) {
// 	(*ptrglDeleteLists)(list, range_);
// }
// GLuint goglGenLists(GLsizei range_) {
// 	return (*ptrglGenLists)(range_);
// }
// void goglListBase(GLuint base) {
// 	(*ptrglListBase)(base);
// }
// void goglBegin(GLenum mode) {
// 	(*ptrglBegin)(mode);
// }
// void goglBitmap(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
// 	(*ptrglBitmap)(width, height, xorig, yorig, xmove, ymove, bitmap);
// }
// void goglColor3b(GLbyte red, GLbyte green, GLbyte blue) {
// 	(*ptrglColor3b)(red, green, blue);
// }
// void goglColor3bv(GLbyte* v) {
// 	(*ptrglColor3bv)(v);
// }
// void goglColor3d(GLdouble red, GLdouble green, GLdouble blue) {
// 	(*ptrglColor3d)(red, green, blue);
// }
// void goglColor3dv(GLdouble* v) {
// 	(*ptrglColor3dv)(v);
// }
// void goglColor3f(GLfloat red, GLfloat green, GLfloat blue) {
// 	(*ptrglColor3f)(red, green, blue);
// }
// void goglColor3fv(GLfloat* v) {
// 	(*ptrglColor3fv)(v);
// }
// void goglColor3i(GLint red, GLint green, GLint blue) {
// 	(*ptrglColor3i)(red, green, blue);
// }
// void goglColor3iv(GLint* v) {
// 	(*ptrglColor3iv)(v);
// }
// void goglColor3s(GLshort red, GLshort green, GLshort blue) {
// 	(*ptrglColor3s)(red, green, blue);
// }
// void goglColor3sv(GLshort* v) {
// 	(*ptrglColor3sv)(v);
// }
// void goglColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
// 	(*ptrglColor3ub)(red, green, blue);
// }
// void goglColor3ubv(GLubyte* v) {
// 	(*ptrglColor3ubv)(v);
// }
// void goglColor3ui(GLuint red, GLuint green, GLuint blue) {
// 	(*ptrglColor3ui)(red, green, blue);
// }
// void goglColor3uiv(GLuint* v) {
// 	(*ptrglColor3uiv)(v);
// }
// void goglColor3us(GLushort red, GLushort green, GLushort blue) {
// 	(*ptrglColor3us)(red, green, blue);
// }
// void goglColor3usv(GLushort* v) {
// 	(*ptrglColor3usv)(v);
// }
// void goglColor4b(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
// 	(*ptrglColor4b)(red, green, blue, alpha);
// }
// void goglColor4bv(GLbyte* v) {
// 	(*ptrglColor4bv)(v);
// }
// void goglColor4d(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
// 	(*ptrglColor4d)(red, green, blue, alpha);
// }
// void goglColor4dv(GLdouble* v) {
// 	(*ptrglColor4dv)(v);
// }
// void goglColor4f(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrglColor4f)(red, green, blue, alpha);
// }
// void goglColor4fv(GLfloat* v) {
// 	(*ptrglColor4fv)(v);
// }
// void goglColor4i(GLint red, GLint green, GLint blue, GLint alpha) {
// 	(*ptrglColor4i)(red, green, blue, alpha);
// }
// void goglColor4iv(GLint* v) {
// 	(*ptrglColor4iv)(v);
// }
// void goglColor4s(GLshort red, GLshort green, GLshort blue, GLshort alpha) {
// 	(*ptrglColor4s)(red, green, blue, alpha);
// }
// void goglColor4sv(GLshort* v) {
// 	(*ptrglColor4sv)(v);
// }
// void goglColor4ub(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
// 	(*ptrglColor4ub)(red, green, blue, alpha);
// }
// void goglColor4ubv(GLubyte* v) {
// 	(*ptrglColor4ubv)(v);
// }
// void goglColor4ui(GLuint red, GLuint green, GLuint blue, GLuint alpha) {
// 	(*ptrglColor4ui)(red, green, blue, alpha);
// }
// void goglColor4uiv(GLuint* v) {
// 	(*ptrglColor4uiv)(v);
// }
// void goglColor4us(GLushort red, GLushort green, GLushort blue, GLushort alpha) {
// 	(*ptrglColor4us)(red, green, blue, alpha);
// }
// void goglColor4usv(GLushort* v) {
// 	(*ptrglColor4usv)(v);
// }
// void goglEdgeFlag(GLboolean flag) {
// 	(*ptrglEdgeFlag)(flag);
// }
// void goglEdgeFlagv(GLboolean* flag) {
// 	(*ptrglEdgeFlagv)(flag);
// }
// void goglEnd() {
// 	(*ptrglEnd)();
// }
// void goglIndexd(GLdouble c) {
// 	(*ptrglIndexd)(c);
// }
// void goglIndexdv(GLdouble* c) {
// 	(*ptrglIndexdv)(c);
// }
// void goglIndexf(GLfloat c) {
// 	(*ptrglIndexf)(c);
// }
// void goglIndexfv(GLfloat* c) {
// 	(*ptrglIndexfv)(c);
// }
// void goglIndexi(GLint c) {
// 	(*ptrglIndexi)(c);
// }
// void goglIndexiv(GLint* c) {
// 	(*ptrglIndexiv)(c);
// }
// void goglIndexs(GLshort c) {
// 	(*ptrglIndexs)(c);
// }
// void goglIndexsv(GLshort* c) {
// 	(*ptrglIndexsv)(c);
// }
// void goglNormal3b(GLbyte nx, GLbyte ny, GLbyte nz) {
// 	(*ptrglNormal3b)(nx, ny, nz);
// }
// void goglNormal3bv(GLbyte* v) {
// 	(*ptrglNormal3bv)(v);
// }
// void goglNormal3d(GLdouble nx, GLdouble ny, GLdouble nz) {
// 	(*ptrglNormal3d)(nx, ny, nz);
// }
// void goglNormal3dv(GLdouble* v) {
// 	(*ptrglNormal3dv)(v);
// }
// void goglNormal3f(GLfloat nx, GLfloat ny, GLfloat nz) {
// 	(*ptrglNormal3f)(nx, ny, nz);
// }
// void goglNormal3fv(GLfloat* v) {
// 	(*ptrglNormal3fv)(v);
// }
// void goglNormal3i(GLint nx, GLint ny, GLint nz) {
// 	(*ptrglNormal3i)(nx, ny, nz);
// }
// void goglNormal3iv(GLint* v) {
// 	(*ptrglNormal3iv)(v);
// }
// void goglNormal3s(GLshort nx, GLshort ny, GLshort nz) {
// 	(*ptrglNormal3s)(nx, ny, nz);
// }
// void goglNormal3sv(GLshort* v) {
// 	(*ptrglNormal3sv)(v);
// }
// void goglRasterPos2d(GLdouble x, GLdouble y) {
// 	(*ptrglRasterPos2d)(x, y);
// }
// void goglRasterPos2dv(GLdouble* v) {
// 	(*ptrglRasterPos2dv)(v);
// }
// void goglRasterPos2f(GLfloat x, GLfloat y) {
// 	(*ptrglRasterPos2f)(x, y);
// }
// void goglRasterPos2fv(GLfloat* v) {
// 	(*ptrglRasterPos2fv)(v);
// }
// void goglRasterPos2i(GLint x, GLint y) {
// 	(*ptrglRasterPos2i)(x, y);
// }
// void goglRasterPos2iv(GLint* v) {
// 	(*ptrglRasterPos2iv)(v);
// }
// void goglRasterPos2s(GLshort x, GLshort y) {
// 	(*ptrglRasterPos2s)(x, y);
// }
// void goglRasterPos2sv(GLshort* v) {
// 	(*ptrglRasterPos2sv)(v);
// }
// void goglRasterPos3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglRasterPos3d)(x, y, z);
// }
// void goglRasterPos3dv(GLdouble* v) {
// 	(*ptrglRasterPos3dv)(v);
// }
// void goglRasterPos3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglRasterPos3f)(x, y, z);
// }
// void goglRasterPos3fv(GLfloat* v) {
// 	(*ptrglRasterPos3fv)(v);
// }
// void goglRasterPos3i(GLint x, GLint y, GLint z) {
// 	(*ptrglRasterPos3i)(x, y, z);
// }
// void goglRasterPos3iv(GLint* v) {
// 	(*ptrglRasterPos3iv)(v);
// }
// void goglRasterPos3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrglRasterPos3s)(x, y, z);
// }
// void goglRasterPos3sv(GLshort* v) {
// 	(*ptrglRasterPos3sv)(v);
// }
// void goglRasterPos4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglRasterPos4d)(x, y, z, w);
// }
// void goglRasterPos4dv(GLdouble* v) {
// 	(*ptrglRasterPos4dv)(v);
// }
// void goglRasterPos4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglRasterPos4f)(x, y, z, w);
// }
// void goglRasterPos4fv(GLfloat* v) {
// 	(*ptrglRasterPos4fv)(v);
// }
// void goglRasterPos4i(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglRasterPos4i)(x, y, z, w);
// }
// void goglRasterPos4iv(GLint* v) {
// 	(*ptrglRasterPos4iv)(v);
// }
// void goglRasterPos4s(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglRasterPos4s)(x, y, z, w);
// }
// void goglRasterPos4sv(GLshort* v) {
// 	(*ptrglRasterPos4sv)(v);
// }
// void goglRectd(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
// 	(*ptrglRectd)(x1, y1, x2, y2);
// }
// void goglRectdv(GLdouble* v1, GLdouble* v2) {
// 	(*ptrglRectdv)(v1, v2);
// }
// void goglRectf(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
// 	(*ptrglRectf)(x1, y1, x2, y2);
// }
// void goglRectfv(GLfloat* v1, GLfloat* v2) {
// 	(*ptrglRectfv)(v1, v2);
// }
// void goglRecti(GLint x1, GLint y1, GLint x2, GLint y2) {
// 	(*ptrglRecti)(x1, y1, x2, y2);
// }
// void goglRectiv(GLint* v1, GLint* v2) {
// 	(*ptrglRectiv)(v1, v2);
// }
// void goglRects(GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
// 	(*ptrglRects)(x1, y1, x2, y2);
// }
// void goglRectsv(GLshort* v1, GLshort* v2) {
// 	(*ptrglRectsv)(v1, v2);
// }
// void goglTexCoord1d(GLdouble s) {
// 	(*ptrglTexCoord1d)(s);
// }
// void goglTexCoord1dv(GLdouble* v) {
// 	(*ptrglTexCoord1dv)(v);
// }
// void goglTexCoord1f(GLfloat s) {
// 	(*ptrglTexCoord1f)(s);
// }
// void goglTexCoord1fv(GLfloat* v) {
// 	(*ptrglTexCoord1fv)(v);
// }
// void goglTexCoord1i(GLint s) {
// 	(*ptrglTexCoord1i)(s);
// }
// void goglTexCoord1iv(GLint* v) {
// 	(*ptrglTexCoord1iv)(v);
// }
// void goglTexCoord1s(GLshort s) {
// 	(*ptrglTexCoord1s)(s);
// }
// void goglTexCoord1sv(GLshort* v) {
// 	(*ptrglTexCoord1sv)(v);
// }
// void goglTexCoord2d(GLdouble s, GLdouble t) {
// 	(*ptrglTexCoord2d)(s, t);
// }
// void goglTexCoord2dv(GLdouble* v) {
// 	(*ptrglTexCoord2dv)(v);
// }
// void goglTexCoord2f(GLfloat s, GLfloat t) {
// 	(*ptrglTexCoord2f)(s, t);
// }
// void goglTexCoord2fv(GLfloat* v) {
// 	(*ptrglTexCoord2fv)(v);
// }
// void goglTexCoord2i(GLint s, GLint t) {
// 	(*ptrglTexCoord2i)(s, t);
// }
// void goglTexCoord2iv(GLint* v) {
// 	(*ptrglTexCoord2iv)(v);
// }
// void goglTexCoord2s(GLshort s, GLshort t) {
// 	(*ptrglTexCoord2s)(s, t);
// }
// void goglTexCoord2sv(GLshort* v) {
// 	(*ptrglTexCoord2sv)(v);
// }
// void goglTexCoord3d(GLdouble s, GLdouble t, GLdouble r) {
// 	(*ptrglTexCoord3d)(s, t, r);
// }
// void goglTexCoord3dv(GLdouble* v) {
// 	(*ptrglTexCoord3dv)(v);
// }
// void goglTexCoord3f(GLfloat s, GLfloat t, GLfloat r) {
// 	(*ptrglTexCoord3f)(s, t, r);
// }
// void goglTexCoord3fv(GLfloat* v) {
// 	(*ptrglTexCoord3fv)(v);
// }
// void goglTexCoord3i(GLint s, GLint t, GLint r) {
// 	(*ptrglTexCoord3i)(s, t, r);
// }
// void goglTexCoord3iv(GLint* v) {
// 	(*ptrglTexCoord3iv)(v);
// }
// void goglTexCoord3s(GLshort s, GLshort t, GLshort r) {
// 	(*ptrglTexCoord3s)(s, t, r);
// }
// void goglTexCoord3sv(GLshort* v) {
// 	(*ptrglTexCoord3sv)(v);
// }
// void goglTexCoord4d(GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
// 	(*ptrglTexCoord4d)(s, t, r, q);
// }
// void goglTexCoord4dv(GLdouble* v) {
// 	(*ptrglTexCoord4dv)(v);
// }
// void goglTexCoord4f(GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
// 	(*ptrglTexCoord4f)(s, t, r, q);
// }
// void goglTexCoord4fv(GLfloat* v) {
// 	(*ptrglTexCoord4fv)(v);
// }
// void goglTexCoord4i(GLint s, GLint t, GLint r, GLint q) {
// 	(*ptrglTexCoord4i)(s, t, r, q);
// }
// void goglTexCoord4iv(GLint* v) {
// 	(*ptrglTexCoord4iv)(v);
// }
// void goglTexCoord4s(GLshort s, GLshort t, GLshort r, GLshort q) {
// 	(*ptrglTexCoord4s)(s, t, r, q);
// }
// void goglTexCoord4sv(GLshort* v) {
// 	(*ptrglTexCoord4sv)(v);
// }
// void goglVertex2d(GLdouble x, GLdouble y) {
// 	(*ptrglVertex2d)(x, y);
// }
// void goglVertex2dv(GLdouble* v) {
// 	(*ptrglVertex2dv)(v);
// }
// void goglVertex2f(GLfloat x, GLfloat y) {
// 	(*ptrglVertex2f)(x, y);
// }
// void goglVertex2fv(GLfloat* v) {
// 	(*ptrglVertex2fv)(v);
// }
// void goglVertex2i(GLint x, GLint y) {
// 	(*ptrglVertex2i)(x, y);
// }
// void goglVertex2iv(GLint* v) {
// 	(*ptrglVertex2iv)(v);
// }
// void goglVertex2s(GLshort x, GLshort y) {
// 	(*ptrglVertex2s)(x, y);
// }
// void goglVertex2sv(GLshort* v) {
// 	(*ptrglVertex2sv)(v);
// }
// void goglVertex3d(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglVertex3d)(x, y, z);
// }
// void goglVertex3dv(GLdouble* v) {
// 	(*ptrglVertex3dv)(v);
// }
// void goglVertex3f(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglVertex3f)(x, y, z);
// }
// void goglVertex3fv(GLfloat* v) {
// 	(*ptrglVertex3fv)(v);
// }
// void goglVertex3i(GLint x, GLint y, GLint z) {
// 	(*ptrglVertex3i)(x, y, z);
// }
// void goglVertex3iv(GLint* v) {
// 	(*ptrglVertex3iv)(v);
// }
// void goglVertex3s(GLshort x, GLshort y, GLshort z) {
// 	(*ptrglVertex3s)(x, y, z);
// }
// void goglVertex3sv(GLshort* v) {
// 	(*ptrglVertex3sv)(v);
// }
// void goglVertex4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
// 	(*ptrglVertex4d)(x, y, z, w);
// }
// void goglVertex4dv(GLdouble* v) {
// 	(*ptrglVertex4dv)(v);
// }
// void goglVertex4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	(*ptrglVertex4f)(x, y, z, w);
// }
// void goglVertex4fv(GLfloat* v) {
// 	(*ptrglVertex4fv)(v);
// }
// void goglVertex4i(GLint x, GLint y, GLint z, GLint w) {
// 	(*ptrglVertex4i)(x, y, z, w);
// }
// void goglVertex4iv(GLint* v) {
// 	(*ptrglVertex4iv)(v);
// }
// void goglVertex4s(GLshort x, GLshort y, GLshort z, GLshort w) {
// 	(*ptrglVertex4s)(x, y, z, w);
// }
// void goglVertex4sv(GLshort* v) {
// 	(*ptrglVertex4sv)(v);
// }
// void goglClipPlane(GLenum plane, GLdouble* equation) {
// 	(*ptrglClipPlane)(plane, equation);
// }
// void goglColorMaterial(GLenum face, GLenum mode) {
// 	(*ptrglColorMaterial)(face, mode);
// }
// void goglFogf(GLenum pname, GLfloat param) {
// 	(*ptrglFogf)(pname, param);
// }
// void goglFogfv(GLenum pname, GLfloat* params) {
// 	(*ptrglFogfv)(pname, params);
// }
// void goglFogi(GLenum pname, GLint param) {
// 	(*ptrglFogi)(pname, param);
// }
// void goglFogiv(GLenum pname, GLint* params) {
// 	(*ptrglFogiv)(pname, params);
// }
// void goglLightf(GLenum light, GLenum pname, GLfloat param) {
// 	(*ptrglLightf)(light, pname, param);
// }
// void goglLightfv(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrglLightfv)(light, pname, params);
// }
// void goglLighti(GLenum light, GLenum pname, GLint param) {
// 	(*ptrglLighti)(light, pname, param);
// }
// void goglLightiv(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrglLightiv)(light, pname, params);
// }
// void goglLightModelf(GLenum pname, GLfloat param) {
// 	(*ptrglLightModelf)(pname, param);
// }
// void goglLightModelfv(GLenum pname, GLfloat* params) {
// 	(*ptrglLightModelfv)(pname, params);
// }
// void goglLightModeli(GLenum pname, GLint param) {
// 	(*ptrglLightModeli)(pname, param);
// }
// void goglLightModeliv(GLenum pname, GLint* params) {
// 	(*ptrglLightModeliv)(pname, params);
// }
// void goglLineStipple(GLint factor, GLushort pattern) {
// 	(*ptrglLineStipple)(factor, pattern);
// }
// void goglMaterialf(GLenum face, GLenum pname, GLfloat param) {
// 	(*ptrglMaterialf)(face, pname, param);
// }
// void goglMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrglMaterialfv)(face, pname, params);
// }
// void goglMateriali(GLenum face, GLenum pname, GLint param) {
// 	(*ptrglMateriali)(face, pname, param);
// }
// void goglMaterialiv(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrglMaterialiv)(face, pname, params);
// }
// void goglPolygonStipple(GLubyte* mask) {
// 	(*ptrglPolygonStipple)(mask);
// }
// void goglShadeModel(GLenum mode) {
// 	(*ptrglShadeModel)(mode);
// }
// void goglTexEnvf(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrglTexEnvf)(target, pname, param);
// }
// void goglTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglTexEnvfv)(target, pname, params);
// }
// void goglTexEnvi(GLenum target, GLenum pname, GLint param) {
// 	(*ptrglTexEnvi)(target, pname, param);
// }
// void goglTexEnviv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglTexEnviv)(target, pname, params);
// }
// void goglTexGend(GLenum coord, GLenum pname, GLdouble param) {
// 	(*ptrglTexGend)(coord, pname, param);
// }
// void goglTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
// 	(*ptrglTexGendv)(coord, pname, params);
// }
// void goglTexGenf(GLenum coord, GLenum pname, GLfloat param) {
// 	(*ptrglTexGenf)(coord, pname, param);
// }
// void goglTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
// 	(*ptrglTexGenfv)(coord, pname, params);
// }
// void goglTexGeni(GLenum coord, GLenum pname, GLint param) {
// 	(*ptrglTexGeni)(coord, pname, param);
// }
// void goglTexGeniv(GLenum coord, GLenum pname, GLint* params) {
// 	(*ptrglTexGeniv)(coord, pname, params);
// }
// void goglFeedbackBuffer(GLsizei size, GLenum type_, GLfloat* buffer) {
// 	(*ptrglFeedbackBuffer)(size, type_, buffer);
// }
// void goglSelectBuffer(GLsizei size, GLuint* buffer) {
// 	(*ptrglSelectBuffer)(size, buffer);
// }
// GLint goglRenderMode(GLenum mode) {
// 	return (*ptrglRenderMode)(mode);
// }
// void goglInitNames() {
// 	(*ptrglInitNames)();
// }
// void goglLoadName(GLuint name) {
// 	(*ptrglLoadName)(name);
// }
// void goglPassThrough(GLfloat token) {
// 	(*ptrglPassThrough)(token);
// }
// void goglPopName() {
// 	(*ptrglPopName)();
// }
// void goglPushName(GLuint name) {
// 	(*ptrglPushName)(name);
// }
// void goglClearAccum(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
// 	(*ptrglClearAccum)(red, green, blue, alpha);
// }
// void goglClearIndex(GLfloat c) {
// 	(*ptrglClearIndex)(c);
// }
// void goglIndexMask(GLuint mask) {
// 	(*ptrglIndexMask)(mask);
// }
// void goglAccum(GLenum op, GLfloat value) {
// 	(*ptrglAccum)(op, value);
// }
// void goglPopAttrib() {
// 	(*ptrglPopAttrib)();
// }
// void goglPushAttrib(GLbitfield mask) {
// 	(*ptrglPushAttrib)(mask);
// }
// void goglMap1d(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
// 	(*ptrglMap1d)(target, u1, u2, stride, order, points);
// }
// void goglMap1f(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
// 	(*ptrglMap1f)(target, u1, u2, stride, order, points);
// }
// void goglMap2d(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
// 	(*ptrglMap2d)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMap2f(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
// 	(*ptrglMap2f)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// void goglMapGrid1d(GLint un, GLdouble u1, GLdouble u2) {
// 	(*ptrglMapGrid1d)(un, u1, u2);
// }
// void goglMapGrid1f(GLint un, GLfloat u1, GLfloat u2) {
// 	(*ptrglMapGrid1f)(un, u1, u2);
// }
// void goglMapGrid2d(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
// 	(*ptrglMapGrid2d)(un, u1, u2, vn, v1, v2);
// }
// void goglMapGrid2f(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
// 	(*ptrglMapGrid2f)(un, u1, u2, vn, v1, v2);
// }
// void goglEvalCoord1d(GLdouble u) {
// 	(*ptrglEvalCoord1d)(u);
// }
// void goglEvalCoord1dv(GLdouble* u) {
// 	(*ptrglEvalCoord1dv)(u);
// }
// void goglEvalCoord1f(GLfloat u) {
// 	(*ptrglEvalCoord1f)(u);
// }
// void goglEvalCoord1fv(GLfloat* u) {
// 	(*ptrglEvalCoord1fv)(u);
// }
// void goglEvalCoord2d(GLdouble u, GLdouble v) {
// 	(*ptrglEvalCoord2d)(u, v);
// }
// void goglEvalCoord2dv(GLdouble* u) {
// 	(*ptrglEvalCoord2dv)(u);
// }
// void goglEvalCoord2f(GLfloat u, GLfloat v) {
// 	(*ptrglEvalCoord2f)(u, v);
// }
// void goglEvalCoord2fv(GLfloat* u) {
// 	(*ptrglEvalCoord2fv)(u);
// }
// void goglEvalMesh1(GLenum mode, GLint i1, GLint i2) {
// 	(*ptrglEvalMesh1)(mode, i1, i2);
// }
// void goglEvalPoint1(GLint i) {
// 	(*ptrglEvalPoint1)(i);
// }
// void goglEvalMesh2(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
// 	(*ptrglEvalMesh2)(mode, i1, i2, j1, j2);
// }
// void goglEvalPoint2(GLint i, GLint j) {
// 	(*ptrglEvalPoint2)(i, j);
// }
// void goglAlphaFunc(GLenum func_, GLfloat ref) {
// 	(*ptrglAlphaFunc)(func_, ref);
// }
// void goglPixelZoom(GLfloat xfactor, GLfloat yfactor) {
// 	(*ptrglPixelZoom)(xfactor, yfactor);
// }
// void goglPixelTransferf(GLenum pname, GLfloat param) {
// 	(*ptrglPixelTransferf)(pname, param);
// }
// void goglPixelTransferi(GLenum pname, GLint param) {
// 	(*ptrglPixelTransferi)(pname, param);
// }
// void goglPixelMapfv(GLenum map_, GLsizei mapsize, GLfloat* values) {
// 	(*ptrglPixelMapfv)(map_, mapsize, values);
// }
// void goglPixelMapuiv(GLenum map_, GLsizei mapsize, GLuint* values) {
// 	(*ptrglPixelMapuiv)(map_, mapsize, values);
// }
// void goglPixelMapusv(GLenum map_, GLsizei mapsize, GLushort* values) {
// 	(*ptrglPixelMapusv)(map_, mapsize, values);
// }
// void goglCopyPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type_) {
// 	(*ptrglCopyPixels)(x, y, width, height, type_);
// }
// void goglDrawPixels(GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {
// 	(*ptrglDrawPixels)(width, height, format, type_, pixels);
// }
// void goglGetClipPlane(GLenum plane, GLdouble* equation) {
// 	(*ptrglGetClipPlane)(plane, equation);
// }
// void goglGetLightfv(GLenum light, GLenum pname, GLfloat* params) {
// 	(*ptrglGetLightfv)(light, pname, params);
// }
// void goglGetLightiv(GLenum light, GLenum pname, GLint* params) {
// 	(*ptrglGetLightiv)(light, pname, params);
// }
// void goglGetMapdv(GLenum target, GLenum query, GLdouble* v) {
// 	(*ptrglGetMapdv)(target, query, v);
// }
// void goglGetMapfv(GLenum target, GLenum query, GLfloat* v) {
// 	(*ptrglGetMapfv)(target, query, v);
// }
// void goglGetMapiv(GLenum target, GLenum query, GLint* v) {
// 	(*ptrglGetMapiv)(target, query, v);
// }
// void goglGetMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
// 	(*ptrglGetMaterialfv)(face, pname, params);
// }
// void goglGetMaterialiv(GLenum face, GLenum pname, GLint* params) {
// 	(*ptrglGetMaterialiv)(face, pname, params);
// }
// void goglGetPixelMapfv(GLenum map_, GLfloat* values) {
// 	(*ptrglGetPixelMapfv)(map_, values);
// }
// void goglGetPixelMapuiv(GLenum map_, GLuint* values) {
// 	(*ptrglGetPixelMapuiv)(map_, values);
// }
// void goglGetPixelMapusv(GLenum map_, GLushort* values) {
// 	(*ptrglGetPixelMapusv)(map_, values);
// }
// void goglGetPolygonStipple(GLubyte* mask) {
// 	(*ptrglGetPolygonStipple)(mask);
// }
// void goglGetTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexEnvfv)(target, pname, params);
// }
// void goglGetTexEnviv(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetTexEnviv)(target, pname, params);
// }
// void goglGetTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
// 	(*ptrglGetTexGendv)(coord, pname, params);
// }
// void goglGetTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
// 	(*ptrglGetTexGenfv)(coord, pname, params);
// }
// void goglGetTexGeniv(GLenum coord, GLenum pname, GLint* params) {
// 	(*ptrglGetTexGeniv)(coord, pname, params);
// }
// GLboolean goglIsList(GLuint list) {
// 	return (*ptrglIsList)(list);
// }
// void goglFrustum(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
// 	(*ptrglFrustum)(left, right, bottom, top, zNear, zFar);
// }
// void goglLoadIdentity() {
// 	(*ptrglLoadIdentity)();
// }
// void goglLoadMatrixf(GLfloat* m) {
// 	(*ptrglLoadMatrixf)(m);
// }
// void goglLoadMatrixd(GLdouble* m) {
// 	(*ptrglLoadMatrixd)(m);
// }
// void goglMatrixMode(GLenum mode) {
// 	(*ptrglMatrixMode)(mode);
// }
// void goglMultMatrixf(GLfloat* m) {
// 	(*ptrglMultMatrixf)(m);
// }
// void goglMultMatrixd(GLdouble* m) {
// 	(*ptrglMultMatrixd)(m);
// }
// void goglOrtho(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
// 	(*ptrglOrtho)(left, right, bottom, top, zNear, zFar);
// }
// void goglPopMatrix() {
// 	(*ptrglPopMatrix)();
// }
// void goglPushMatrix() {
// 	(*ptrglPushMatrix)();
// }
// void goglRotated(GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglRotated)(angle, x, y, z);
// }
// void goglRotatef(GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglRotatef)(angle, x, y, z);
// }
// void goglScaled(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglScaled)(x, y, z);
// }
// void goglScalef(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglScalef)(x, y, z);
// }
// void goglTranslated(GLdouble x, GLdouble y, GLdouble z) {
// 	(*ptrglTranslated)(x, y, z);
// }
// void goglTranslatef(GLfloat x, GLfloat y, GLfloat z) {
// 	(*ptrglTranslatef)(x, y, z);
// }
// 
// int init_VERSION_1_0() {
// 	ptrglCullFace = goglGetProcAddress("glCullFace");
// 	if(ptrglCullFace == NULL) return 1;
// 	ptrglFrontFace = goglGetProcAddress("glFrontFace");
// 	if(ptrglFrontFace == NULL) return 1;
// 	ptrglHint = goglGetProcAddress("glHint");
// 	if(ptrglHint == NULL) return 1;
// 	ptrglLineWidth = goglGetProcAddress("glLineWidth");
// 	if(ptrglLineWidth == NULL) return 1;
// 	ptrglPointSize = goglGetProcAddress("glPointSize");
// 	if(ptrglPointSize == NULL) return 1;
// 	ptrglPolygonMode = goglGetProcAddress("glPolygonMode");
// 	if(ptrglPolygonMode == NULL) return 1;
// 	ptrglScissor = goglGetProcAddress("glScissor");
// 	if(ptrglScissor == NULL) return 1;
// 	ptrglTexParameterf = goglGetProcAddress("glTexParameterf");
// 	if(ptrglTexParameterf == NULL) return 1;
// 	ptrglTexParameterfv = goglGetProcAddress("glTexParameterfv");
// 	if(ptrglTexParameterfv == NULL) return 1;
// 	ptrglTexParameteri = goglGetProcAddress("glTexParameteri");
// 	if(ptrglTexParameteri == NULL) return 1;
// 	ptrglTexParameteriv = goglGetProcAddress("glTexParameteriv");
// 	if(ptrglTexParameteriv == NULL) return 1;
// 	ptrglTexImage1D = goglGetProcAddress("glTexImage1D");
// 	if(ptrglTexImage1D == NULL) return 1;
// 	ptrglTexImage2D = goglGetProcAddress("glTexImage2D");
// 	if(ptrglTexImage2D == NULL) return 1;
// 	ptrglDrawBuffer = goglGetProcAddress("glDrawBuffer");
// 	if(ptrglDrawBuffer == NULL) return 1;
// 	ptrglClear = goglGetProcAddress("glClear");
// 	if(ptrglClear == NULL) return 1;
// 	ptrglClearColor = goglGetProcAddress("glClearColor");
// 	if(ptrglClearColor == NULL) return 1;
// 	ptrglClearStencil = goglGetProcAddress("glClearStencil");
// 	if(ptrglClearStencil == NULL) return 1;
// 	ptrglClearDepth = goglGetProcAddress("glClearDepth");
// 	if(ptrglClearDepth == NULL) return 1;
// 	ptrglStencilMask = goglGetProcAddress("glStencilMask");
// 	if(ptrglStencilMask == NULL) return 1;
// 	ptrglColorMask = goglGetProcAddress("glColorMask");
// 	if(ptrglColorMask == NULL) return 1;
// 	ptrglDepthMask = goglGetProcAddress("glDepthMask");
// 	if(ptrglDepthMask == NULL) return 1;
// 	ptrglDisable = goglGetProcAddress("glDisable");
// 	if(ptrglDisable == NULL) return 1;
// 	ptrglEnable = goglGetProcAddress("glEnable");
// 	if(ptrglEnable == NULL) return 1;
// 	ptrglFinish = goglGetProcAddress("glFinish");
// 	if(ptrglFinish == NULL) return 1;
// 	ptrglFlush = goglGetProcAddress("glFlush");
// 	if(ptrglFlush == NULL) return 1;
// 	ptrglBlendFunc = goglGetProcAddress("glBlendFunc");
// 	if(ptrglBlendFunc == NULL) return 1;
// 	ptrglLogicOp = goglGetProcAddress("glLogicOp");
// 	if(ptrglLogicOp == NULL) return 1;
// 	ptrglStencilFunc = goglGetProcAddress("glStencilFunc");
// 	if(ptrglStencilFunc == NULL) return 1;
// 	ptrglStencilOp = goglGetProcAddress("glStencilOp");
// 	if(ptrglStencilOp == NULL) return 1;
// 	ptrglDepthFunc = goglGetProcAddress("glDepthFunc");
// 	if(ptrglDepthFunc == NULL) return 1;
// 	ptrglPixelStoref = goglGetProcAddress("glPixelStoref");
// 	if(ptrglPixelStoref == NULL) return 1;
// 	ptrglPixelStorei = goglGetProcAddress("glPixelStorei");
// 	if(ptrglPixelStorei == NULL) return 1;
// 	ptrglReadBuffer = goglGetProcAddress("glReadBuffer");
// 	if(ptrglReadBuffer == NULL) return 1;
// 	ptrglReadPixels = goglGetProcAddress("glReadPixels");
// 	if(ptrglReadPixels == NULL) return 1;
// 	ptrglGetBooleanv = goglGetProcAddress("glGetBooleanv");
// 	if(ptrglGetBooleanv == NULL) return 1;
// 	ptrglGetDoublev = goglGetProcAddress("glGetDoublev");
// 	if(ptrglGetDoublev == NULL) return 1;
// 	ptrglGetError = goglGetProcAddress("glGetError");
// 	if(ptrglGetError == NULL) return 1;
// 	ptrglGetFloatv = goglGetProcAddress("glGetFloatv");
// 	if(ptrglGetFloatv == NULL) return 1;
// 	ptrglGetIntegerv = goglGetProcAddress("glGetIntegerv");
// 	if(ptrglGetIntegerv == NULL) return 1;
// 	ptrglGetString = goglGetProcAddress("glGetString");
// 	if(ptrglGetString == NULL) return 1;
// 	ptrglGetTexImage = goglGetProcAddress("glGetTexImage");
// 	if(ptrglGetTexImage == NULL) return 1;
// 	ptrglGetTexParameterfv = goglGetProcAddress("glGetTexParameterfv");
// 	if(ptrglGetTexParameterfv == NULL) return 1;
// 	ptrglGetTexParameteriv = goglGetProcAddress("glGetTexParameteriv");
// 	if(ptrglGetTexParameteriv == NULL) return 1;
// 	ptrglGetTexLevelParameterfv = goglGetProcAddress("glGetTexLevelParameterfv");
// 	if(ptrglGetTexLevelParameterfv == NULL) return 1;
// 	ptrglGetTexLevelParameteriv = goglGetProcAddress("glGetTexLevelParameteriv");
// 	if(ptrglGetTexLevelParameteriv == NULL) return 1;
// 	ptrglIsEnabled = goglGetProcAddress("glIsEnabled");
// 	if(ptrglIsEnabled == NULL) return 1;
// 	ptrglDepthRange = goglGetProcAddress("glDepthRange");
// 	if(ptrglDepthRange == NULL) return 1;
// 	ptrglViewport = goglGetProcAddress("glViewport");
// 	if(ptrglViewport == NULL) return 1;
// 	return 0;
// }
// int init_VERSION_1_0_DEPRECATED() {
// 	ptrglNewList = goglGetProcAddress("glNewList");
// 	if(ptrglNewList == NULL) return 1;
// 	ptrglEndList = goglGetProcAddress("glEndList");
// 	if(ptrglEndList == NULL) return 1;
// 	ptrglCallList = goglGetProcAddress("glCallList");
// 	if(ptrglCallList == NULL) return 1;
// 	ptrglCallLists = goglGetProcAddress("glCallLists");
// 	if(ptrglCallLists == NULL) return 1;
// 	ptrglDeleteLists = goglGetProcAddress("glDeleteLists");
// 	if(ptrglDeleteLists == NULL) return 1;
// 	ptrglGenLists = goglGetProcAddress("glGenLists");
// 	if(ptrglGenLists == NULL) return 1;
// 	ptrglListBase = goglGetProcAddress("glListBase");
// 	if(ptrglListBase == NULL) return 1;
// 	ptrglBegin = goglGetProcAddress("glBegin");
// 	if(ptrglBegin == NULL) return 1;
// 	ptrglBitmap = goglGetProcAddress("glBitmap");
// 	if(ptrglBitmap == NULL) return 1;
// 	ptrglColor3b = goglGetProcAddress("glColor3b");
// 	if(ptrglColor3b == NULL) return 1;
// 	ptrglColor3bv = goglGetProcAddress("glColor3bv");
// 	if(ptrglColor3bv == NULL) return 1;
// 	ptrglColor3d = goglGetProcAddress("glColor3d");
// 	if(ptrglColor3d == NULL) return 1;
// 	ptrglColor3dv = goglGetProcAddress("glColor3dv");
// 	if(ptrglColor3dv == NULL) return 1;
// 	ptrglColor3f = goglGetProcAddress("glColor3f");
// 	if(ptrglColor3f == NULL) return 1;
// 	ptrglColor3fv = goglGetProcAddress("glColor3fv");
// 	if(ptrglColor3fv == NULL) return 1;
// 	ptrglColor3i = goglGetProcAddress("glColor3i");
// 	if(ptrglColor3i == NULL) return 1;
// 	ptrglColor3iv = goglGetProcAddress("glColor3iv");
// 	if(ptrglColor3iv == NULL) return 1;
// 	ptrglColor3s = goglGetProcAddress("glColor3s");
// 	if(ptrglColor3s == NULL) return 1;
// 	ptrglColor3sv = goglGetProcAddress("glColor3sv");
// 	if(ptrglColor3sv == NULL) return 1;
// 	ptrglColor3ub = goglGetProcAddress("glColor3ub");
// 	if(ptrglColor3ub == NULL) return 1;
// 	ptrglColor3ubv = goglGetProcAddress("glColor3ubv");
// 	if(ptrglColor3ubv == NULL) return 1;
// 	ptrglColor3ui = goglGetProcAddress("glColor3ui");
// 	if(ptrglColor3ui == NULL) return 1;
// 	ptrglColor3uiv = goglGetProcAddress("glColor3uiv");
// 	if(ptrglColor3uiv == NULL) return 1;
// 	ptrglColor3us = goglGetProcAddress("glColor3us");
// 	if(ptrglColor3us == NULL) return 1;
// 	ptrglColor3usv = goglGetProcAddress("glColor3usv");
// 	if(ptrglColor3usv == NULL) return 1;
// 	ptrglColor4b = goglGetProcAddress("glColor4b");
// 	if(ptrglColor4b == NULL) return 1;
// 	ptrglColor4bv = goglGetProcAddress("glColor4bv");
// 	if(ptrglColor4bv == NULL) return 1;
// 	ptrglColor4d = goglGetProcAddress("glColor4d");
// 	if(ptrglColor4d == NULL) return 1;
// 	ptrglColor4dv = goglGetProcAddress("glColor4dv");
// 	if(ptrglColor4dv == NULL) return 1;
// 	ptrglColor4f = goglGetProcAddress("glColor4f");
// 	if(ptrglColor4f == NULL) return 1;
// 	ptrglColor4fv = goglGetProcAddress("glColor4fv");
// 	if(ptrglColor4fv == NULL) return 1;
// 	ptrglColor4i = goglGetProcAddress("glColor4i");
// 	if(ptrglColor4i == NULL) return 1;
// 	ptrglColor4iv = goglGetProcAddress("glColor4iv");
// 	if(ptrglColor4iv == NULL) return 1;
// 	ptrglColor4s = goglGetProcAddress("glColor4s");
// 	if(ptrglColor4s == NULL) return 1;
// 	ptrglColor4sv = goglGetProcAddress("glColor4sv");
// 	if(ptrglColor4sv == NULL) return 1;
// 	ptrglColor4ub = goglGetProcAddress("glColor4ub");
// 	if(ptrglColor4ub == NULL) return 1;
// 	ptrglColor4ubv = goglGetProcAddress("glColor4ubv");
// 	if(ptrglColor4ubv == NULL) return 1;
// 	ptrglColor4ui = goglGetProcAddress("glColor4ui");
// 	if(ptrglColor4ui == NULL) return 1;
// 	ptrglColor4uiv = goglGetProcAddress("glColor4uiv");
// 	if(ptrglColor4uiv == NULL) return 1;
// 	ptrglColor4us = goglGetProcAddress("glColor4us");
// 	if(ptrglColor4us == NULL) return 1;
// 	ptrglColor4usv = goglGetProcAddress("glColor4usv");
// 	if(ptrglColor4usv == NULL) return 1;
// 	ptrglEdgeFlag = goglGetProcAddress("glEdgeFlag");
// 	if(ptrglEdgeFlag == NULL) return 1;
// 	ptrglEdgeFlagv = goglGetProcAddress("glEdgeFlagv");
// 	if(ptrglEdgeFlagv == NULL) return 1;
// 	ptrglEnd = goglGetProcAddress("glEnd");
// 	if(ptrglEnd == NULL) return 1;
// 	ptrglIndexd = goglGetProcAddress("glIndexd");
// 	if(ptrglIndexd == NULL) return 1;
// 	ptrglIndexdv = goglGetProcAddress("glIndexdv");
// 	if(ptrglIndexdv == NULL) return 1;
// 	ptrglIndexf = goglGetProcAddress("glIndexf");
// 	if(ptrglIndexf == NULL) return 1;
// 	ptrglIndexfv = goglGetProcAddress("glIndexfv");
// 	if(ptrglIndexfv == NULL) return 1;
// 	ptrglIndexi = goglGetProcAddress("glIndexi");
// 	if(ptrglIndexi == NULL) return 1;
// 	ptrglIndexiv = goglGetProcAddress("glIndexiv");
// 	if(ptrglIndexiv == NULL) return 1;
// 	ptrglIndexs = goglGetProcAddress("glIndexs");
// 	if(ptrglIndexs == NULL) return 1;
// 	ptrglIndexsv = goglGetProcAddress("glIndexsv");
// 	if(ptrglIndexsv == NULL) return 1;
// 	ptrglNormal3b = goglGetProcAddress("glNormal3b");
// 	if(ptrglNormal3b == NULL) return 1;
// 	ptrglNormal3bv = goglGetProcAddress("glNormal3bv");
// 	if(ptrglNormal3bv == NULL) return 1;
// 	ptrglNormal3d = goglGetProcAddress("glNormal3d");
// 	if(ptrglNormal3d == NULL) return 1;
// 	ptrglNormal3dv = goglGetProcAddress("glNormal3dv");
// 	if(ptrglNormal3dv == NULL) return 1;
// 	ptrglNormal3f = goglGetProcAddress("glNormal3f");
// 	if(ptrglNormal3f == NULL) return 1;
// 	ptrglNormal3fv = goglGetProcAddress("glNormal3fv");
// 	if(ptrglNormal3fv == NULL) return 1;
// 	ptrglNormal3i = goglGetProcAddress("glNormal3i");
// 	if(ptrglNormal3i == NULL) return 1;
// 	ptrglNormal3iv = goglGetProcAddress("glNormal3iv");
// 	if(ptrglNormal3iv == NULL) return 1;
// 	ptrglNormal3s = goglGetProcAddress("glNormal3s");
// 	if(ptrglNormal3s == NULL) return 1;
// 	ptrglNormal3sv = goglGetProcAddress("glNormal3sv");
// 	if(ptrglNormal3sv == NULL) return 1;
// 	ptrglRasterPos2d = goglGetProcAddress("glRasterPos2d");
// 	if(ptrglRasterPos2d == NULL) return 1;
// 	ptrglRasterPos2dv = goglGetProcAddress("glRasterPos2dv");
// 	if(ptrglRasterPos2dv == NULL) return 1;
// 	ptrglRasterPos2f = goglGetProcAddress("glRasterPos2f");
// 	if(ptrglRasterPos2f == NULL) return 1;
// 	ptrglRasterPos2fv = goglGetProcAddress("glRasterPos2fv");
// 	if(ptrglRasterPos2fv == NULL) return 1;
// 	ptrglRasterPos2i = goglGetProcAddress("glRasterPos2i");
// 	if(ptrglRasterPos2i == NULL) return 1;
// 	ptrglRasterPos2iv = goglGetProcAddress("glRasterPos2iv");
// 	if(ptrglRasterPos2iv == NULL) return 1;
// 	ptrglRasterPos2s = goglGetProcAddress("glRasterPos2s");
// 	if(ptrglRasterPos2s == NULL) return 1;
// 	ptrglRasterPos2sv = goglGetProcAddress("glRasterPos2sv");
// 	if(ptrglRasterPos2sv == NULL) return 1;
// 	ptrglRasterPos3d = goglGetProcAddress("glRasterPos3d");
// 	if(ptrglRasterPos3d == NULL) return 1;
// 	ptrglRasterPos3dv = goglGetProcAddress("glRasterPos3dv");
// 	if(ptrglRasterPos3dv == NULL) return 1;
// 	ptrglRasterPos3f = goglGetProcAddress("glRasterPos3f");
// 	if(ptrglRasterPos3f == NULL) return 1;
// 	ptrglRasterPos3fv = goglGetProcAddress("glRasterPos3fv");
// 	if(ptrglRasterPos3fv == NULL) return 1;
// 	ptrglRasterPos3i = goglGetProcAddress("glRasterPos3i");
// 	if(ptrglRasterPos3i == NULL) return 1;
// 	ptrglRasterPos3iv = goglGetProcAddress("glRasterPos3iv");
// 	if(ptrglRasterPos3iv == NULL) return 1;
// 	ptrglRasterPos3s = goglGetProcAddress("glRasterPos3s");
// 	if(ptrglRasterPos3s == NULL) return 1;
// 	ptrglRasterPos3sv = goglGetProcAddress("glRasterPos3sv");
// 	if(ptrglRasterPos3sv == NULL) return 1;
// 	ptrglRasterPos4d = goglGetProcAddress("glRasterPos4d");
// 	if(ptrglRasterPos4d == NULL) return 1;
// 	ptrglRasterPos4dv = goglGetProcAddress("glRasterPos4dv");
// 	if(ptrglRasterPos4dv == NULL) return 1;
// 	ptrglRasterPos4f = goglGetProcAddress("glRasterPos4f");
// 	if(ptrglRasterPos4f == NULL) return 1;
// 	ptrglRasterPos4fv = goglGetProcAddress("glRasterPos4fv");
// 	if(ptrglRasterPos4fv == NULL) return 1;
// 	ptrglRasterPos4i = goglGetProcAddress("glRasterPos4i");
// 	if(ptrglRasterPos4i == NULL) return 1;
// 	ptrglRasterPos4iv = goglGetProcAddress("glRasterPos4iv");
// 	if(ptrglRasterPos4iv == NULL) return 1;
// 	ptrglRasterPos4s = goglGetProcAddress("glRasterPos4s");
// 	if(ptrglRasterPos4s == NULL) return 1;
// 	ptrglRasterPos4sv = goglGetProcAddress("glRasterPos4sv");
// 	if(ptrglRasterPos4sv == NULL) return 1;
// 	ptrglRectd = goglGetProcAddress("glRectd");
// 	if(ptrglRectd == NULL) return 1;
// 	ptrglRectdv = goglGetProcAddress("glRectdv");
// 	if(ptrglRectdv == NULL) return 1;
// 	ptrglRectf = goglGetProcAddress("glRectf");
// 	if(ptrglRectf == NULL) return 1;
// 	ptrglRectfv = goglGetProcAddress("glRectfv");
// 	if(ptrglRectfv == NULL) return 1;
// 	ptrglRecti = goglGetProcAddress("glRecti");
// 	if(ptrglRecti == NULL) return 1;
// 	ptrglRectiv = goglGetProcAddress("glRectiv");
// 	if(ptrglRectiv == NULL) return 1;
// 	ptrglRects = goglGetProcAddress("glRects");
// 	if(ptrglRects == NULL) return 1;
// 	ptrglRectsv = goglGetProcAddress("glRectsv");
// 	if(ptrglRectsv == NULL) return 1;
// 	ptrglTexCoord1d = goglGetProcAddress("glTexCoord1d");
// 	if(ptrglTexCoord1d == NULL) return 1;
// 	ptrglTexCoord1dv = goglGetProcAddress("glTexCoord1dv");
// 	if(ptrglTexCoord1dv == NULL) return 1;
// 	ptrglTexCoord1f = goglGetProcAddress("glTexCoord1f");
// 	if(ptrglTexCoord1f == NULL) return 1;
// 	ptrglTexCoord1fv = goglGetProcAddress("glTexCoord1fv");
// 	if(ptrglTexCoord1fv == NULL) return 1;
// 	ptrglTexCoord1i = goglGetProcAddress("glTexCoord1i");
// 	if(ptrglTexCoord1i == NULL) return 1;
// 	ptrglTexCoord1iv = goglGetProcAddress("glTexCoord1iv");
// 	if(ptrglTexCoord1iv == NULL) return 1;
// 	ptrglTexCoord1s = goglGetProcAddress("glTexCoord1s");
// 	if(ptrglTexCoord1s == NULL) return 1;
// 	ptrglTexCoord1sv = goglGetProcAddress("glTexCoord1sv");
// 	if(ptrglTexCoord1sv == NULL) return 1;
// 	ptrglTexCoord2d = goglGetProcAddress("glTexCoord2d");
// 	if(ptrglTexCoord2d == NULL) return 1;
// 	ptrglTexCoord2dv = goglGetProcAddress("glTexCoord2dv");
// 	if(ptrglTexCoord2dv == NULL) return 1;
// 	ptrglTexCoord2f = goglGetProcAddress("glTexCoord2f");
// 	if(ptrglTexCoord2f == NULL) return 1;
// 	ptrglTexCoord2fv = goglGetProcAddress("glTexCoord2fv");
// 	if(ptrglTexCoord2fv == NULL) return 1;
// 	ptrglTexCoord2i = goglGetProcAddress("glTexCoord2i");
// 	if(ptrglTexCoord2i == NULL) return 1;
// 	ptrglTexCoord2iv = goglGetProcAddress("glTexCoord2iv");
// 	if(ptrglTexCoord2iv == NULL) return 1;
// 	ptrglTexCoord2s = goglGetProcAddress("glTexCoord2s");
// 	if(ptrglTexCoord2s == NULL) return 1;
// 	ptrglTexCoord2sv = goglGetProcAddress("glTexCoord2sv");
// 	if(ptrglTexCoord2sv == NULL) return 1;
// 	ptrglTexCoord3d = goglGetProcAddress("glTexCoord3d");
// 	if(ptrglTexCoord3d == NULL) return 1;
// 	ptrglTexCoord3dv = goglGetProcAddress("glTexCoord3dv");
// 	if(ptrglTexCoord3dv == NULL) return 1;
// 	ptrglTexCoord3f = goglGetProcAddress("glTexCoord3f");
// 	if(ptrglTexCoord3f == NULL) return 1;
// 	ptrglTexCoord3fv = goglGetProcAddress("glTexCoord3fv");
// 	if(ptrglTexCoord3fv == NULL) return 1;
// 	ptrglTexCoord3i = goglGetProcAddress("glTexCoord3i");
// 	if(ptrglTexCoord3i == NULL) return 1;
// 	ptrglTexCoord3iv = goglGetProcAddress("glTexCoord3iv");
// 	if(ptrglTexCoord3iv == NULL) return 1;
// 	ptrglTexCoord3s = goglGetProcAddress("glTexCoord3s");
// 	if(ptrglTexCoord3s == NULL) return 1;
// 	ptrglTexCoord3sv = goglGetProcAddress("glTexCoord3sv");
// 	if(ptrglTexCoord3sv == NULL) return 1;
// 	ptrglTexCoord4d = goglGetProcAddress("glTexCoord4d");
// 	if(ptrglTexCoord4d == NULL) return 1;
// 	ptrglTexCoord4dv = goglGetProcAddress("glTexCoord4dv");
// 	if(ptrglTexCoord4dv == NULL) return 1;
// 	ptrglTexCoord4f = goglGetProcAddress("glTexCoord4f");
// 	if(ptrglTexCoord4f == NULL) return 1;
// 	ptrglTexCoord4fv = goglGetProcAddress("glTexCoord4fv");
// 	if(ptrglTexCoord4fv == NULL) return 1;
// 	ptrglTexCoord4i = goglGetProcAddress("glTexCoord4i");
// 	if(ptrglTexCoord4i == NULL) return 1;
// 	ptrglTexCoord4iv = goglGetProcAddress("glTexCoord4iv");
// 	if(ptrglTexCoord4iv == NULL) return 1;
// 	ptrglTexCoord4s = goglGetProcAddress("glTexCoord4s");
// 	if(ptrglTexCoord4s == NULL) return 1;
// 	ptrglTexCoord4sv = goglGetProcAddress("glTexCoord4sv");
// 	if(ptrglTexCoord4sv == NULL) return 1;
// 	ptrglVertex2d = goglGetProcAddress("glVertex2d");
// 	if(ptrglVertex2d == NULL) return 1;
// 	ptrglVertex2dv = goglGetProcAddress("glVertex2dv");
// 	if(ptrglVertex2dv == NULL) return 1;
// 	ptrglVertex2f = goglGetProcAddress("glVertex2f");
// 	if(ptrglVertex2f == NULL) return 1;
// 	ptrglVertex2fv = goglGetProcAddress("glVertex2fv");
// 	if(ptrglVertex2fv == NULL) return 1;
// 	ptrglVertex2i = goglGetProcAddress("glVertex2i");
// 	if(ptrglVertex2i == NULL) return 1;
// 	ptrglVertex2iv = goglGetProcAddress("glVertex2iv");
// 	if(ptrglVertex2iv == NULL) return 1;
// 	ptrglVertex2s = goglGetProcAddress("glVertex2s");
// 	if(ptrglVertex2s == NULL) return 1;
// 	ptrglVertex2sv = goglGetProcAddress("glVertex2sv");
// 	if(ptrglVertex2sv == NULL) return 1;
// 	ptrglVertex3d = goglGetProcAddress("glVertex3d");
// 	if(ptrglVertex3d == NULL) return 1;
// 	ptrglVertex3dv = goglGetProcAddress("glVertex3dv");
// 	if(ptrglVertex3dv == NULL) return 1;
// 	ptrglVertex3f = goglGetProcAddress("glVertex3f");
// 	if(ptrglVertex3f == NULL) return 1;
// 	ptrglVertex3fv = goglGetProcAddress("glVertex3fv");
// 	if(ptrglVertex3fv == NULL) return 1;
// 	ptrglVertex3i = goglGetProcAddress("glVertex3i");
// 	if(ptrglVertex3i == NULL) return 1;
// 	ptrglVertex3iv = goglGetProcAddress("glVertex3iv");
// 	if(ptrglVertex3iv == NULL) return 1;
// 	ptrglVertex3s = goglGetProcAddress("glVertex3s");
// 	if(ptrglVertex3s == NULL) return 1;
// 	ptrglVertex3sv = goglGetProcAddress("glVertex3sv");
// 	if(ptrglVertex3sv == NULL) return 1;
// 	ptrglVertex4d = goglGetProcAddress("glVertex4d");
// 	if(ptrglVertex4d == NULL) return 1;
// 	ptrglVertex4dv = goglGetProcAddress("glVertex4dv");
// 	if(ptrglVertex4dv == NULL) return 1;
// 	ptrglVertex4f = goglGetProcAddress("glVertex4f");
// 	if(ptrglVertex4f == NULL) return 1;
// 	ptrglVertex4fv = goglGetProcAddress("glVertex4fv");
// 	if(ptrglVertex4fv == NULL) return 1;
// 	ptrglVertex4i = goglGetProcAddress("glVertex4i");
// 	if(ptrglVertex4i == NULL) return 1;
// 	ptrglVertex4iv = goglGetProcAddress("glVertex4iv");
// 	if(ptrglVertex4iv == NULL) return 1;
// 	ptrglVertex4s = goglGetProcAddress("glVertex4s");
// 	if(ptrglVertex4s == NULL) return 1;
// 	ptrglVertex4sv = goglGetProcAddress("glVertex4sv");
// 	if(ptrglVertex4sv == NULL) return 1;
// 	ptrglClipPlane = goglGetProcAddress("glClipPlane");
// 	if(ptrglClipPlane == NULL) return 1;
// 	ptrglColorMaterial = goglGetProcAddress("glColorMaterial");
// 	if(ptrglColorMaterial == NULL) return 1;
// 	ptrglFogf = goglGetProcAddress("glFogf");
// 	if(ptrglFogf == NULL) return 1;
// 	ptrglFogfv = goglGetProcAddress("glFogfv");
// 	if(ptrglFogfv == NULL) return 1;
// 	ptrglFogi = goglGetProcAddress("glFogi");
// 	if(ptrglFogi == NULL) return 1;
// 	ptrglFogiv = goglGetProcAddress("glFogiv");
// 	if(ptrglFogiv == NULL) return 1;
// 	ptrglLightf = goglGetProcAddress("glLightf");
// 	if(ptrglLightf == NULL) return 1;
// 	ptrglLightfv = goglGetProcAddress("glLightfv");
// 	if(ptrglLightfv == NULL) return 1;
// 	ptrglLighti = goglGetProcAddress("glLighti");
// 	if(ptrglLighti == NULL) return 1;
// 	ptrglLightiv = goglGetProcAddress("glLightiv");
// 	if(ptrglLightiv == NULL) return 1;
// 	ptrglLightModelf = goglGetProcAddress("glLightModelf");
// 	if(ptrglLightModelf == NULL) return 1;
// 	ptrglLightModelfv = goglGetProcAddress("glLightModelfv");
// 	if(ptrglLightModelfv == NULL) return 1;
// 	ptrglLightModeli = goglGetProcAddress("glLightModeli");
// 	if(ptrglLightModeli == NULL) return 1;
// 	ptrglLightModeliv = goglGetProcAddress("glLightModeliv");
// 	if(ptrglLightModeliv == NULL) return 1;
// 	ptrglLineStipple = goglGetProcAddress("glLineStipple");
// 	if(ptrglLineStipple == NULL) return 1;
// 	ptrglMaterialf = goglGetProcAddress("glMaterialf");
// 	if(ptrglMaterialf == NULL) return 1;
// 	ptrglMaterialfv = goglGetProcAddress("glMaterialfv");
// 	if(ptrglMaterialfv == NULL) return 1;
// 	ptrglMateriali = goglGetProcAddress("glMateriali");
// 	if(ptrglMateriali == NULL) return 1;
// 	ptrglMaterialiv = goglGetProcAddress("glMaterialiv");
// 	if(ptrglMaterialiv == NULL) return 1;
// 	ptrglPolygonStipple = goglGetProcAddress("glPolygonStipple");
// 	if(ptrglPolygonStipple == NULL) return 1;
// 	ptrglShadeModel = goglGetProcAddress("glShadeModel");
// 	if(ptrglShadeModel == NULL) return 1;
// 	ptrglTexEnvf = goglGetProcAddress("glTexEnvf");
// 	if(ptrglTexEnvf == NULL) return 1;
// 	ptrglTexEnvfv = goglGetProcAddress("glTexEnvfv");
// 	if(ptrglTexEnvfv == NULL) return 1;
// 	ptrglTexEnvi = goglGetProcAddress("glTexEnvi");
// 	if(ptrglTexEnvi == NULL) return 1;
// 	ptrglTexEnviv = goglGetProcAddress("glTexEnviv");
// 	if(ptrglTexEnviv == NULL) return 1;
// 	ptrglTexGend = goglGetProcAddress("glTexGend");
// 	if(ptrglTexGend == NULL) return 1;
// 	ptrglTexGendv = goglGetProcAddress("glTexGendv");
// 	if(ptrglTexGendv == NULL) return 1;
// 	ptrglTexGenf = goglGetProcAddress("glTexGenf");
// 	if(ptrglTexGenf == NULL) return 1;
// 	ptrglTexGenfv = goglGetProcAddress("glTexGenfv");
// 	if(ptrglTexGenfv == NULL) return 1;
// 	ptrglTexGeni = goglGetProcAddress("glTexGeni");
// 	if(ptrglTexGeni == NULL) return 1;
// 	ptrglTexGeniv = goglGetProcAddress("glTexGeniv");
// 	if(ptrglTexGeniv == NULL) return 1;
// 	ptrglFeedbackBuffer = goglGetProcAddress("glFeedbackBuffer");
// 	if(ptrglFeedbackBuffer == NULL) return 1;
// 	ptrglSelectBuffer = goglGetProcAddress("glSelectBuffer");
// 	if(ptrglSelectBuffer == NULL) return 1;
// 	ptrglRenderMode = goglGetProcAddress("glRenderMode");
// 	if(ptrglRenderMode == NULL) return 1;
// 	ptrglInitNames = goglGetProcAddress("glInitNames");
// 	if(ptrglInitNames == NULL) return 1;
// 	ptrglLoadName = goglGetProcAddress("glLoadName");
// 	if(ptrglLoadName == NULL) return 1;
// 	ptrglPassThrough = goglGetProcAddress("glPassThrough");
// 	if(ptrglPassThrough == NULL) return 1;
// 	ptrglPopName = goglGetProcAddress("glPopName");
// 	if(ptrglPopName == NULL) return 1;
// 	ptrglPushName = goglGetProcAddress("glPushName");
// 	if(ptrglPushName == NULL) return 1;
// 	ptrglClearAccum = goglGetProcAddress("glClearAccum");
// 	if(ptrglClearAccum == NULL) return 1;
// 	ptrglClearIndex = goglGetProcAddress("glClearIndex");
// 	if(ptrglClearIndex == NULL) return 1;
// 	ptrglIndexMask = goglGetProcAddress("glIndexMask");
// 	if(ptrglIndexMask == NULL) return 1;
// 	ptrglAccum = goglGetProcAddress("glAccum");
// 	if(ptrglAccum == NULL) return 1;
// 	ptrglPopAttrib = goglGetProcAddress("glPopAttrib");
// 	if(ptrglPopAttrib == NULL) return 1;
// 	ptrglPushAttrib = goglGetProcAddress("glPushAttrib");
// 	if(ptrglPushAttrib == NULL) return 1;
// 	ptrglMap1d = goglGetProcAddress("glMap1d");
// 	if(ptrglMap1d == NULL) return 1;
// 	ptrglMap1f = goglGetProcAddress("glMap1f");
// 	if(ptrglMap1f == NULL) return 1;
// 	ptrglMap2d = goglGetProcAddress("glMap2d");
// 	if(ptrglMap2d == NULL) return 1;
// 	ptrglMap2f = goglGetProcAddress("glMap2f");
// 	if(ptrglMap2f == NULL) return 1;
// 	ptrglMapGrid1d = goglGetProcAddress("glMapGrid1d");
// 	if(ptrglMapGrid1d == NULL) return 1;
// 	ptrglMapGrid1f = goglGetProcAddress("glMapGrid1f");
// 	if(ptrglMapGrid1f == NULL) return 1;
// 	ptrglMapGrid2d = goglGetProcAddress("glMapGrid2d");
// 	if(ptrglMapGrid2d == NULL) return 1;
// 	ptrglMapGrid2f = goglGetProcAddress("glMapGrid2f");
// 	if(ptrglMapGrid2f == NULL) return 1;
// 	ptrglEvalCoord1d = goglGetProcAddress("glEvalCoord1d");
// 	if(ptrglEvalCoord1d == NULL) return 1;
// 	ptrglEvalCoord1dv = goglGetProcAddress("glEvalCoord1dv");
// 	if(ptrglEvalCoord1dv == NULL) return 1;
// 	ptrglEvalCoord1f = goglGetProcAddress("glEvalCoord1f");
// 	if(ptrglEvalCoord1f == NULL) return 1;
// 	ptrglEvalCoord1fv = goglGetProcAddress("glEvalCoord1fv");
// 	if(ptrglEvalCoord1fv == NULL) return 1;
// 	ptrglEvalCoord2d = goglGetProcAddress("glEvalCoord2d");
// 	if(ptrglEvalCoord2d == NULL) return 1;
// 	ptrglEvalCoord2dv = goglGetProcAddress("glEvalCoord2dv");
// 	if(ptrglEvalCoord2dv == NULL) return 1;
// 	ptrglEvalCoord2f = goglGetProcAddress("glEvalCoord2f");
// 	if(ptrglEvalCoord2f == NULL) return 1;
// 	ptrglEvalCoord2fv = goglGetProcAddress("glEvalCoord2fv");
// 	if(ptrglEvalCoord2fv == NULL) return 1;
// 	ptrglEvalMesh1 = goglGetProcAddress("glEvalMesh1");
// 	if(ptrglEvalMesh1 == NULL) return 1;
// 	ptrglEvalPoint1 = goglGetProcAddress("glEvalPoint1");
// 	if(ptrglEvalPoint1 == NULL) return 1;
// 	ptrglEvalMesh2 = goglGetProcAddress("glEvalMesh2");
// 	if(ptrglEvalMesh2 == NULL) return 1;
// 	ptrglEvalPoint2 = goglGetProcAddress("glEvalPoint2");
// 	if(ptrglEvalPoint2 == NULL) return 1;
// 	ptrglAlphaFunc = goglGetProcAddress("glAlphaFunc");
// 	if(ptrglAlphaFunc == NULL) return 1;
// 	ptrglPixelZoom = goglGetProcAddress("glPixelZoom");
// 	if(ptrglPixelZoom == NULL) return 1;
// 	ptrglPixelTransferf = goglGetProcAddress("glPixelTransferf");
// 	if(ptrglPixelTransferf == NULL) return 1;
// 	ptrglPixelTransferi = goglGetProcAddress("glPixelTransferi");
// 	if(ptrglPixelTransferi == NULL) return 1;
// 	ptrglPixelMapfv = goglGetProcAddress("glPixelMapfv");
// 	if(ptrglPixelMapfv == NULL) return 1;
// 	ptrglPixelMapuiv = goglGetProcAddress("glPixelMapuiv");
// 	if(ptrglPixelMapuiv == NULL) return 1;
// 	ptrglPixelMapusv = goglGetProcAddress("glPixelMapusv");
// 	if(ptrglPixelMapusv == NULL) return 1;
// 	ptrglCopyPixels = goglGetProcAddress("glCopyPixels");
// 	if(ptrglCopyPixels == NULL) return 1;
// 	ptrglDrawPixels = goglGetProcAddress("glDrawPixels");
// 	if(ptrglDrawPixels == NULL) return 1;
// 	ptrglGetClipPlane = goglGetProcAddress("glGetClipPlane");
// 	if(ptrglGetClipPlane == NULL) return 1;
// 	ptrglGetLightfv = goglGetProcAddress("glGetLightfv");
// 	if(ptrglGetLightfv == NULL) return 1;
// 	ptrglGetLightiv = goglGetProcAddress("glGetLightiv");
// 	if(ptrglGetLightiv == NULL) return 1;
// 	ptrglGetMapdv = goglGetProcAddress("glGetMapdv");
// 	if(ptrglGetMapdv == NULL) return 1;
// 	ptrglGetMapfv = goglGetProcAddress("glGetMapfv");
// 	if(ptrglGetMapfv == NULL) return 1;
// 	ptrglGetMapiv = goglGetProcAddress("glGetMapiv");
// 	if(ptrglGetMapiv == NULL) return 1;
// 	ptrglGetMaterialfv = goglGetProcAddress("glGetMaterialfv");
// 	if(ptrglGetMaterialfv == NULL) return 1;
// 	ptrglGetMaterialiv = goglGetProcAddress("glGetMaterialiv");
// 	if(ptrglGetMaterialiv == NULL) return 1;
// 	ptrglGetPixelMapfv = goglGetProcAddress("glGetPixelMapfv");
// 	if(ptrglGetPixelMapfv == NULL) return 1;
// 	ptrglGetPixelMapuiv = goglGetProcAddress("glGetPixelMapuiv");
// 	if(ptrglGetPixelMapuiv == NULL) return 1;
// 	ptrglGetPixelMapusv = goglGetProcAddress("glGetPixelMapusv");
// 	if(ptrglGetPixelMapusv == NULL) return 1;
// 	ptrglGetPolygonStipple = goglGetProcAddress("glGetPolygonStipple");
// 	if(ptrglGetPolygonStipple == NULL) return 1;
// 	ptrglGetTexEnvfv = goglGetProcAddress("glGetTexEnvfv");
// 	if(ptrglGetTexEnvfv == NULL) return 1;
// 	ptrglGetTexEnviv = goglGetProcAddress("glGetTexEnviv");
// 	if(ptrglGetTexEnviv == NULL) return 1;
// 	ptrglGetTexGendv = goglGetProcAddress("glGetTexGendv");
// 	if(ptrglGetTexGendv == NULL) return 1;
// 	ptrglGetTexGenfv = goglGetProcAddress("glGetTexGenfv");
// 	if(ptrglGetTexGenfv == NULL) return 1;
// 	ptrglGetTexGeniv = goglGetProcAddress("glGetTexGeniv");
// 	if(ptrglGetTexGeniv == NULL) return 1;
// 	ptrglIsList = goglGetProcAddress("glIsList");
// 	if(ptrglIsList == NULL) return 1;
// 	ptrglFrustum = goglGetProcAddress("glFrustum");
// 	if(ptrglFrustum == NULL) return 1;
// 	ptrglLoadIdentity = goglGetProcAddress("glLoadIdentity");
// 	if(ptrglLoadIdentity == NULL) return 1;
// 	ptrglLoadMatrixf = goglGetProcAddress("glLoadMatrixf");
// 	if(ptrglLoadMatrixf == NULL) return 1;
// 	ptrglLoadMatrixd = goglGetProcAddress("glLoadMatrixd");
// 	if(ptrglLoadMatrixd == NULL) return 1;
// 	ptrglMatrixMode = goglGetProcAddress("glMatrixMode");
// 	if(ptrglMatrixMode == NULL) return 1;
// 	ptrglMultMatrixf = goglGetProcAddress("glMultMatrixf");
// 	if(ptrglMultMatrixf == NULL) return 1;
// 	ptrglMultMatrixd = goglGetProcAddress("glMultMatrixd");
// 	if(ptrglMultMatrixd == NULL) return 1;
// 	ptrglOrtho = goglGetProcAddress("glOrtho");
// 	if(ptrglOrtho == NULL) return 1;
// 	ptrglPopMatrix = goglGetProcAddress("glPopMatrix");
// 	if(ptrglPopMatrix == NULL) return 1;
// 	ptrglPushMatrix = goglGetProcAddress("glPushMatrix");
// 	if(ptrglPushMatrix == NULL) return 1;
// 	ptrglRotated = goglGetProcAddress("glRotated");
// 	if(ptrglRotated == NULL) return 1;
// 	ptrglRotatef = goglGetProcAddress("glRotatef");
// 	if(ptrglRotatef == NULL) return 1;
// 	ptrglScaled = goglGetProcAddress("glScaled");
// 	if(ptrglScaled == NULL) return 1;
// 	ptrglScalef = goglGetProcAddress("glScalef");
// 	if(ptrglScalef == NULL) return 1;
// 	ptrglTranslated = goglGetProcAddress("glTranslated");
// 	if(ptrglTranslated == NULL) return 1;
// 	ptrglTranslatef = goglGetProcAddress("glTranslatef");
// 	if(ptrglTranslatef == NULL) return 1;
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

// VERSION_1_0

// http://www.opengl.org/sdk/docs/man/xhtml/glCullFace.xml
func CullFace(mode Enum)  {
	C.goglCullFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFrontFace.xml
func FrontFace(mode Enum)  {
	C.goglFrontFace((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glHint.xml
func Hint(target Enum, mode Enum)  {
	C.goglHint((C.GLenum)(target), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLineWidth.xml
func LineWidth(width Float)  {
	C.goglLineWidth((C.GLfloat)(width))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPointSize.xml
func PointSize(size Float)  {
	C.goglPointSize((C.GLfloat)(size))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum)  {
	C.goglPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei)  {
	C.goglScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float)  {
	C.goglTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int)  {
	C.goglTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum)  {
	C.goglDrawBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClear.xml
func Clear(mask Bitfield)  {
	C.goglClear((C.GLbitfield)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearColor.xml
func ClearColor(red Float, green Float, blue Float, alpha Float)  {
	C.goglClearColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearStencil.xml
func ClearStencil(s Int)  {
	C.goglClearStencil((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearDepth.xml
func ClearDepth(depth Double)  {
	C.goglClearDepth((C.GLdouble)(depth))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilMask.xml
func StencilMask(mask Uint)  {
	C.goglStencilMask((C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)  {
	C.goglColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDepthMask.xml
func DepthMask(flag Boolean)  {
	C.goglDepthMask((C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDisable.xml
func Disable(cap Enum)  {
	C.goglDisable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnable.xml
func Enable(cap Enum)  {
	C.goglEnable((C.GLenum)(cap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFinish.xml
func Finish()  {
	C.goglFinish()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFlush.xml
func Flush()  {
	C.goglFlush()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum)  {
	C.goglBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLogicOp.xml
func LogicOp(opcode Enum)  {
	C.goglLogicOp((C.GLenum)(opcode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint)  {
	C.goglStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum)  {
	C.goglStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum)  {
	C.goglDepthFunc((C.GLenum)(func_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float)  {
	C.goglPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int)  {
	C.goglPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum)  {
	C.goglReadBuffer((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean)  {
	C.goglGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double)  {
	C.goglGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.goglGetError())
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float)  {
	C.goglGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int)  {
	C.goglGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.goglGetString((C.GLenum)(name)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Pointer)  {
	C.goglGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)  {
	C.goglGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)  {
	C.goglGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.goglIsEnabled((C.GLenum)(cap)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDepthRange.xml
func DepthRange(near_ Double, far_ Double)  {
	C.goglDepthRange((C.GLdouble)(near_), (C.GLdouble)(far_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei)  {
	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
// VERSION_1_0_DEPRECATED

// http://www.opengl.org/sdk/docs/man/xhtml/glNewList.xml
func NewList(list Uint, mode Enum)  {
	C.goglNewList((C.GLuint)(list), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEndList.xml
func EndList()  {
	C.goglEndList()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCallList.xml
func CallList(list Uint)  {
	C.goglCallList((C.GLuint)(list))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCallLists.xml
func CallLists(n Sizei, type_ Enum, lists Pointer)  {
	C.goglCallLists((C.GLsizei)(n), (C.GLenum)(type_), (unsafe.Pointer)(lists))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDeleteLists.xml
func DeleteLists(list Uint, range_ Sizei)  {
	C.goglDeleteLists((C.GLuint)(list), (C.GLsizei)(range_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGenLists.xml
func GenLists(range_ Sizei) Uint {
	return (Uint)(C.goglGenLists((C.GLsizei)(range_)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glListBase.xml
func ListBase(base Uint)  {
	C.goglListBase((C.GLuint)(base))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBegin.xml
func Begin(mode Enum)  {
	C.goglBegin((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glBitmap.xml
func Bitmap(width Sizei, height Sizei, xorig Float, yorig Float, xmove Float, ymove Float, bitmap *Ubyte)  {
	C.goglBitmap((C.GLsizei)(width), (C.GLsizei)(height), (C.GLfloat)(xorig), (C.GLfloat)(yorig), (C.GLfloat)(xmove), (C.GLfloat)(ymove), (*C.GLubyte)(bitmap))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3b.xml
func Color3b(red Byte, green Byte, blue Byte)  {
	C.goglColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3bv.xml
func Color3bv(v *Byte)  {
	C.goglColor3bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3d.xml
func Color3d(red Double, green Double, blue Double)  {
	C.goglColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3dv.xml
func Color3dv(v *Double)  {
	C.goglColor3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3f.xml
func Color3f(red Float, green Float, blue Float)  {
	C.goglColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3fv.xml
func Color3fv(v *Float)  {
	C.goglColor3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3i.xml
func Color3i(red Int, green Int, blue Int)  {
	C.goglColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3iv.xml
func Color3iv(v *Int)  {
	C.goglColor3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3s.xml
func Color3s(red Short, green Short, blue Short)  {
	C.goglColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3sv.xml
func Color3sv(v *Short)  {
	C.goglColor3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3ub.xml
func Color3ub(red Ubyte, green Ubyte, blue Ubyte)  {
	C.goglColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3ubv.xml
func Color3ubv(v *Ubyte)  {
	C.goglColor3ubv((*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3ui.xml
func Color3ui(red Uint, green Uint, blue Uint)  {
	C.goglColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3uiv.xml
func Color3uiv(v *Uint)  {
	C.goglColor3uiv((*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3us.xml
func Color3us(red Ushort, green Ushort, blue Ushort)  {
	C.goglColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor3usv.xml
func Color3usv(v *Ushort)  {
	C.goglColor3usv((*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4b.xml
func Color4b(red Byte, green Byte, blue Byte, alpha Byte)  {
	C.goglColor4b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue), (C.GLbyte)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4bv.xml
func Color4bv(v *Byte)  {
	C.goglColor4bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4d.xml
func Color4d(red Double, green Double, blue Double, alpha Double)  {
	C.goglColor4d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue), (C.GLdouble)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4dv.xml
func Color4dv(v *Double)  {
	C.goglColor4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4f.xml
func Color4f(red Float, green Float, blue Float, alpha Float)  {
	C.goglColor4f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4fv.xml
func Color4fv(v *Float)  {
	C.goglColor4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4i.xml
func Color4i(red Int, green Int, blue Int, alpha Int)  {
	C.goglColor4i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue), (C.GLint)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4iv.xml
func Color4iv(v *Int)  {
	C.goglColor4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4s.xml
func Color4s(red Short, green Short, blue Short, alpha Short)  {
	C.goglColor4s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue), (C.GLshort)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4sv.xml
func Color4sv(v *Short)  {
	C.goglColor4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4ub.xml
func Color4ub(red Ubyte, green Ubyte, blue Ubyte, alpha Ubyte)  {
	C.goglColor4ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue), (C.GLubyte)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4ubv.xml
func Color4ubv(v *Ubyte)  {
	C.goglColor4ubv((*C.GLubyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4ui.xml
func Color4ui(red Uint, green Uint, blue Uint, alpha Uint)  {
	C.goglColor4ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue), (C.GLuint)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4uiv.xml
func Color4uiv(v *Uint)  {
	C.goglColor4uiv((*C.GLuint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4us.xml
func Color4us(red Ushort, green Ushort, blue Ushort, alpha Ushort)  {
	C.goglColor4us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue), (C.GLushort)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColor4usv.xml
func Color4usv(v *Ushort)  {
	C.goglColor4usv((*C.GLushort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEdgeFlag.xml
func EdgeFlag(flag Boolean)  {
	C.goglEdgeFlag((C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEdgeFlagv.xml
func EdgeFlagv(flag *Boolean)  {
	C.goglEdgeFlagv((*C.GLboolean)(flag))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEnd.xml
func End()  {
	C.goglEnd()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexd.xml
func Indexd(c Double)  {
	C.goglIndexd((C.GLdouble)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexdv.xml
func Indexdv(c *Double)  {
	C.goglIndexdv((*C.GLdouble)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexf.xml
func Indexf(c Float)  {
	C.goglIndexf((C.GLfloat)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexfv.xml
func Indexfv(c *Float)  {
	C.goglIndexfv((*C.GLfloat)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexi.xml
func Indexi(c Int)  {
	C.goglIndexi((C.GLint)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexiv.xml
func Indexiv(c *Int)  {
	C.goglIndexiv((*C.GLint)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexs.xml
func Indexs(c Short)  {
	C.goglIndexs((C.GLshort)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexsv.xml
func Indexsv(c *Short)  {
	C.goglIndexsv((*C.GLshort)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3b.xml
func Normal3b(nx Byte, ny Byte, nz Byte)  {
	C.goglNormal3b((C.GLbyte)(nx), (C.GLbyte)(ny), (C.GLbyte)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3bv.xml
func Normal3bv(v *Byte)  {
	C.goglNormal3bv((*C.GLbyte)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3d.xml
func Normal3d(nx Double, ny Double, nz Double)  {
	C.goglNormal3d((C.GLdouble)(nx), (C.GLdouble)(ny), (C.GLdouble)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3dv.xml
func Normal3dv(v *Double)  {
	C.goglNormal3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3f.xml
func Normal3f(nx Float, ny Float, nz Float)  {
	C.goglNormal3f((C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3fv.xml
func Normal3fv(v *Float)  {
	C.goglNormal3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3i.xml
func Normal3i(nx Int, ny Int, nz Int)  {
	C.goglNormal3i((C.GLint)(nx), (C.GLint)(ny), (C.GLint)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3iv.xml
func Normal3iv(v *Int)  {
	C.goglNormal3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3s.xml
func Normal3s(nx Short, ny Short, nz Short)  {
	C.goglNormal3s((C.GLshort)(nx), (C.GLshort)(ny), (C.GLshort)(nz))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glNormal3sv.xml
func Normal3sv(v *Short)  {
	C.goglNormal3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2d.xml
func RasterPos2d(x Double, y Double)  {
	C.goglRasterPos2d((C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2dv.xml
func RasterPos2dv(v *Double)  {
	C.goglRasterPos2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2f.xml
func RasterPos2f(x Float, y Float)  {
	C.goglRasterPos2f((C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2fv.xml
func RasterPos2fv(v *Float)  {
	C.goglRasterPos2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2i.xml
func RasterPos2i(x Int, y Int)  {
	C.goglRasterPos2i((C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2iv.xml
func RasterPos2iv(v *Int)  {
	C.goglRasterPos2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2s.xml
func RasterPos2s(x Short, y Short)  {
	C.goglRasterPos2s((C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos2sv.xml
func RasterPos2sv(v *Short)  {
	C.goglRasterPos2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3d.xml
func RasterPos3d(x Double, y Double, z Double)  {
	C.goglRasterPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3dv.xml
func RasterPos3dv(v *Double)  {
	C.goglRasterPos3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3f.xml
func RasterPos3f(x Float, y Float, z Float)  {
	C.goglRasterPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3fv.xml
func RasterPos3fv(v *Float)  {
	C.goglRasterPos3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3i.xml
func RasterPos3i(x Int, y Int, z Int)  {
	C.goglRasterPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3iv.xml
func RasterPos3iv(v *Int)  {
	C.goglRasterPos3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3s.xml
func RasterPos3s(x Short, y Short, z Short)  {
	C.goglRasterPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos3sv.xml
func RasterPos3sv(v *Short)  {
	C.goglRasterPos3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4d.xml
func RasterPos4d(x Double, y Double, z Double, w Double)  {
	C.goglRasterPos4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4dv.xml
func RasterPos4dv(v *Double)  {
	C.goglRasterPos4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4f.xml
func RasterPos4f(x Float, y Float, z Float, w Float)  {
	C.goglRasterPos4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4fv.xml
func RasterPos4fv(v *Float)  {
	C.goglRasterPos4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4i.xml
func RasterPos4i(x Int, y Int, z Int, w Int)  {
	C.goglRasterPos4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4iv.xml
func RasterPos4iv(v *Int)  {
	C.goglRasterPos4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4s.xml
func RasterPos4s(x Short, y Short, z Short, w Short)  {
	C.goglRasterPos4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRasterPos4sv.xml
func RasterPos4sv(v *Short)  {
	C.goglRasterPos4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectd.xml
func Rectd(x1 Double, y1 Double, x2 Double, y2 Double)  {
	C.goglRectd((C.GLdouble)(x1), (C.GLdouble)(y1), (C.GLdouble)(x2), (C.GLdouble)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectdv.xml
func Rectdv(v1 *Double, v2 *Double)  {
	C.goglRectdv((*C.GLdouble)(v1), (*C.GLdouble)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectf.xml
func Rectf(x1 Float, y1 Float, x2 Float, y2 Float)  {
	C.goglRectf((C.GLfloat)(x1), (C.GLfloat)(y1), (C.GLfloat)(x2), (C.GLfloat)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectfv.xml
func Rectfv(v1 *Float, v2 *Float)  {
	C.goglRectfv((*C.GLfloat)(v1), (*C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRecti.xml
func Recti(x1 Int, y1 Int, x2 Int, y2 Int)  {
	C.goglRecti((C.GLint)(x1), (C.GLint)(y1), (C.GLint)(x2), (C.GLint)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectiv.xml
func Rectiv(v1 *Int, v2 *Int)  {
	C.goglRectiv((*C.GLint)(v1), (*C.GLint)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRects.xml
func Rects(x1 Short, y1 Short, x2 Short, y2 Short)  {
	C.goglRects((C.GLshort)(x1), (C.GLshort)(y1), (C.GLshort)(x2), (C.GLshort)(y2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRectsv.xml
func Rectsv(v1 *Short, v2 *Short)  {
	C.goglRectsv((*C.GLshort)(v1), (*C.GLshort)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1d.xml
func TexCoord1d(s Double)  {
	C.goglTexCoord1d((C.GLdouble)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1dv.xml
func TexCoord1dv(v *Double)  {
	C.goglTexCoord1dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1f.xml
func TexCoord1f(s Float)  {
	C.goglTexCoord1f((C.GLfloat)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1fv.xml
func TexCoord1fv(v *Float)  {
	C.goglTexCoord1fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1i.xml
func TexCoord1i(s Int)  {
	C.goglTexCoord1i((C.GLint)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1iv.xml
func TexCoord1iv(v *Int)  {
	C.goglTexCoord1iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1s.xml
func TexCoord1s(s Short)  {
	C.goglTexCoord1s((C.GLshort)(s))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord1sv.xml
func TexCoord1sv(v *Short)  {
	C.goglTexCoord1sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2d.xml
func TexCoord2d(s Double, t Double)  {
	C.goglTexCoord2d((C.GLdouble)(s), (C.GLdouble)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2dv.xml
func TexCoord2dv(v *Double)  {
	C.goglTexCoord2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2f.xml
func TexCoord2f(s Float, t Float)  {
	C.goglTexCoord2f((C.GLfloat)(s), (C.GLfloat)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2fv.xml
func TexCoord2fv(v *Float)  {
	C.goglTexCoord2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2i.xml
func TexCoord2i(s Int, t Int)  {
	C.goglTexCoord2i((C.GLint)(s), (C.GLint)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2iv.xml
func TexCoord2iv(v *Int)  {
	C.goglTexCoord2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2s.xml
func TexCoord2s(s Short, t Short)  {
	C.goglTexCoord2s((C.GLshort)(s), (C.GLshort)(t))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord2sv.xml
func TexCoord2sv(v *Short)  {
	C.goglTexCoord2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3d.xml
func TexCoord3d(s Double, t Double, r Double)  {
	C.goglTexCoord3d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3dv.xml
func TexCoord3dv(v *Double)  {
	C.goglTexCoord3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3f.xml
func TexCoord3f(s Float, t Float, r Float)  {
	C.goglTexCoord3f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3fv.xml
func TexCoord3fv(v *Float)  {
	C.goglTexCoord3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3i.xml
func TexCoord3i(s Int, t Int, r Int)  {
	C.goglTexCoord3i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3iv.xml
func TexCoord3iv(v *Int)  {
	C.goglTexCoord3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3s.xml
func TexCoord3s(s Short, t Short, r Short)  {
	C.goglTexCoord3s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord3sv.xml
func TexCoord3sv(v *Short)  {
	C.goglTexCoord3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4d.xml
func TexCoord4d(s Double, t Double, r Double, q Double)  {
	C.goglTexCoord4d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4dv.xml
func TexCoord4dv(v *Double)  {
	C.goglTexCoord4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4f.xml
func TexCoord4f(s Float, t Float, r Float, q Float)  {
	C.goglTexCoord4f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4fv.xml
func TexCoord4fv(v *Float)  {
	C.goglTexCoord4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4i.xml
func TexCoord4i(s Int, t Int, r Int, q Int)  {
	C.goglTexCoord4i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4iv.xml
func TexCoord4iv(v *Int)  {
	C.goglTexCoord4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4s.xml
func TexCoord4s(s Short, t Short, r Short, q Short)  {
	C.goglTexCoord4s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexCoord4sv.xml
func TexCoord4sv(v *Short)  {
	C.goglTexCoord4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2d.xml
func Vertex2d(x Double, y Double)  {
	C.goglVertex2d((C.GLdouble)(x), (C.GLdouble)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2dv.xml
func Vertex2dv(v *Double)  {
	C.goglVertex2dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2f.xml
func Vertex2f(x Float, y Float)  {
	C.goglVertex2f((C.GLfloat)(x), (C.GLfloat)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2fv.xml
func Vertex2fv(v *Float)  {
	C.goglVertex2fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2i.xml
func Vertex2i(x Int, y Int)  {
	C.goglVertex2i((C.GLint)(x), (C.GLint)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2iv.xml
func Vertex2iv(v *Int)  {
	C.goglVertex2iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2s.xml
func Vertex2s(x Short, y Short)  {
	C.goglVertex2s((C.GLshort)(x), (C.GLshort)(y))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex2sv.xml
func Vertex2sv(v *Short)  {
	C.goglVertex2sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3d.xml
func Vertex3d(x Double, y Double, z Double)  {
	C.goglVertex3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3dv.xml
func Vertex3dv(v *Double)  {
	C.goglVertex3dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3f.xml
func Vertex3f(x Float, y Float, z Float)  {
	C.goglVertex3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3fv.xml
func Vertex3fv(v *Float)  {
	C.goglVertex3fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3i.xml
func Vertex3i(x Int, y Int, z Int)  {
	C.goglVertex3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3iv.xml
func Vertex3iv(v *Int)  {
	C.goglVertex3iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3s.xml
func Vertex3s(x Short, y Short, z Short)  {
	C.goglVertex3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex3sv.xml
func Vertex3sv(v *Short)  {
	C.goglVertex3sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4d.xml
func Vertex4d(x Double, y Double, z Double, w Double)  {
	C.goglVertex4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4dv.xml
func Vertex4dv(v *Double)  {
	C.goglVertex4dv((*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4f.xml
func Vertex4f(x Float, y Float, z Float, w Float)  {
	C.goglVertex4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4fv.xml
func Vertex4fv(v *Float)  {
	C.goglVertex4fv((*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4i.xml
func Vertex4i(x Int, y Int, z Int, w Int)  {
	C.goglVertex4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4iv.xml
func Vertex4iv(v *Int)  {
	C.goglVertex4iv((*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4s.xml
func Vertex4s(x Short, y Short, z Short, w Short)  {
	C.goglVertex4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glVertex4sv.xml
func Vertex4sv(v *Short)  {
	C.goglVertex4sv((*C.GLshort)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClipPlane.xml
func ClipPlane(plane Enum, equation *Double)  {
	C.goglClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glColorMaterial.xml
func ColorMaterial(face Enum, mode Enum)  {
	C.goglColorMaterial((C.GLenum)(face), (C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogf.xml
func Fogf(pname Enum, param Float)  {
	C.goglFogf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogfv.xml
func Fogfv(pname Enum, params *Float)  {
	C.goglFogfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogi.xml
func Fogi(pname Enum, param Int)  {
	C.goglFogi((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFogiv.xml
func Fogiv(pname Enum, params *Int)  {
	C.goglFogiv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightf.xml
func Lightf(light Enum, pname Enum, param Float)  {
	C.goglLightf((C.GLenum)(light), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightfv.xml
func Lightfv(light Enum, pname Enum, params *Float)  {
	C.goglLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLighti.xml
func Lighti(light Enum, pname Enum, param Int)  {
	C.goglLighti((C.GLenum)(light), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightiv.xml
func Lightiv(light Enum, pname Enum, params *Int)  {
	C.goglLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModelf.xml
func LightModelf(pname Enum, param Float)  {
	C.goglLightModelf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModelfv.xml
func LightModelfv(pname Enum, params *Float)  {
	C.goglLightModelfv((C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModeli.xml
func LightModeli(pname Enum, param Int)  {
	C.goglLightModeli((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLightModeliv.xml
func LightModeliv(pname Enum, params *Int)  {
	C.goglLightModeliv((C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLineStipple.xml
func LineStipple(factor Int, pattern Ushort)  {
	C.goglLineStipple((C.GLint)(factor), (C.GLushort)(pattern))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMaterialf.xml
func Materialf(face Enum, pname Enum, param Float)  {
	C.goglMaterialf((C.GLenum)(face), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMaterialfv.xml
func Materialfv(face Enum, pname Enum, params *Float)  {
	C.goglMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMateriali.xml
func Materiali(face Enum, pname Enum, param Int)  {
	C.goglMateriali((C.GLenum)(face), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMaterialiv.xml
func Materialiv(face Enum, pname Enum, params *Int)  {
	C.goglMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPolygonStipple.xml
func PolygonStipple(mask *Ubyte)  {
	C.goglPolygonStipple((*C.GLubyte)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glShadeModel.xml
func ShadeModel(mode Enum)  {
	C.goglShadeModel((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnvf.xml
func TexEnvf(target Enum, pname Enum, param Float)  {
	C.goglTexEnvf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnvfv.xml
func TexEnvfv(target Enum, pname Enum, params *Float)  {
	C.goglTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnvi.xml
func TexEnvi(target Enum, pname Enum, param Int)  {
	C.goglTexEnvi((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexEnviv.xml
func TexEnviv(target Enum, pname Enum, params *Int)  {
	C.goglTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGend.xml
func TexGend(coord Enum, pname Enum, param Double)  {
	C.goglTexGend((C.GLenum)(coord), (C.GLenum)(pname), (C.GLdouble)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGendv.xml
func TexGendv(coord Enum, pname Enum, params *Double)  {
	C.goglTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGenf.xml
func TexGenf(coord Enum, pname Enum, param Float)  {
	C.goglTexGenf((C.GLenum)(coord), (C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGenfv.xml
func TexGenfv(coord Enum, pname Enum, params *Float)  {
	C.goglTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGeni.xml
func TexGeni(coord Enum, pname Enum, param Int)  {
	C.goglTexGeni((C.GLenum)(coord), (C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTexGeniv.xml
func TexGeniv(coord Enum, pname Enum, params *Int)  {
	C.goglTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFeedbackBuffer.xml
func FeedbackBuffer(size Sizei, type_ Enum, buffer *Float)  {
	C.goglFeedbackBuffer((C.GLsizei)(size), (C.GLenum)(type_), (*C.GLfloat)(buffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glSelectBuffer.xml
func SelectBuffer(size Sizei, buffer *Uint)  {
	C.goglSelectBuffer((C.GLsizei)(size), (*C.GLuint)(buffer))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRenderMode.xml
func RenderMode(mode Enum) Int {
	return (Int)(C.goglRenderMode((C.GLenum)(mode)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glInitNames.xml
func InitNames()  {
	C.goglInitNames()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadName.xml
func LoadName(name Uint)  {
	C.goglLoadName((C.GLuint)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPassThrough.xml
func PassThrough(token Float)  {
	C.goglPassThrough((C.GLfloat)(token))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopName.xml
func PopName()  {
	C.goglPopName()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushName.xml
func PushName(name Uint)  {
	C.goglPushName((C.GLuint)(name))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearAccum.xml
func ClearAccum(red Float, green Float, blue Float, alpha Float)  {
	C.goglClearAccum((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glClearIndex.xml
func ClearIndex(c Float)  {
	C.goglClearIndex((C.GLfloat)(c))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIndexMask.xml
func IndexMask(mask Uint)  {
	C.goglIndexMask((C.GLuint)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glAccum.xml
func Accum(op Enum, value Float)  {
	C.goglAccum((C.GLenum)(op), (C.GLfloat)(value))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopAttrib.xml
func PopAttrib()  {
	C.goglPopAttrib()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushAttrib.xml
func PushAttrib(mask Bitfield)  {
	C.goglPushAttrib((C.GLbitfield)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap1d.xml
func Map1d(target Enum, u1 Double, u2 Double, stride Int, order Int, points *Double)  {
	C.goglMap1d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLdouble)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap1f.xml
func Map1f(target Enum, u1 Float, u2 Float, stride Int, order Int, points *Float)  {
	C.goglMap1f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLfloat)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap2d.xml
func Map2d(target Enum, u1 Double, u2 Double, ustride Int, uorder Int, v1 Double, v2 Double, vstride Int, vorder Int, points *Double)  {
	C.goglMap2d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLdouble)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMap2f.xml
func Map2f(target Enum, u1 Float, u2 Float, ustride Int, uorder Int, v1 Float, v2 Float, vstride Int, vorder Int, points *Float)  {
	C.goglMap2f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLfloat)(points))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid1d.xml
func MapGrid1d(un Int, u1 Double, u2 Double)  {
	C.goglMapGrid1d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid1f.xml
func MapGrid1f(un Int, u1 Float, u2 Float)  {
	C.goglMapGrid1f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid2d.xml
func MapGrid2d(un Int, u1 Double, u2 Double, vn Int, v1 Double, v2 Double)  {
	C.goglMapGrid2d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(vn), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMapGrid2f.xml
func MapGrid2f(un Int, u1 Float, u2 Float, vn Int, v1 Float, v2 Float)  {
	C.goglMapGrid2f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(vn), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1d.xml
func EvalCoord1d(u Double)  {
	C.goglEvalCoord1d((C.GLdouble)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1dv.xml
func EvalCoord1dv(u *Double)  {
	C.goglEvalCoord1dv((*C.GLdouble)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1f.xml
func EvalCoord1f(u Float)  {
	C.goglEvalCoord1f((C.GLfloat)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord1fv.xml
func EvalCoord1fv(u *Float)  {
	C.goglEvalCoord1fv((*C.GLfloat)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2d.xml
func EvalCoord2d(u Double, v Double)  {
	C.goglEvalCoord2d((C.GLdouble)(u), (C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2dv.xml
func EvalCoord2dv(u *Double)  {
	C.goglEvalCoord2dv((*C.GLdouble)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2f.xml
func EvalCoord2f(u Float, v Float)  {
	C.goglEvalCoord2f((C.GLfloat)(u), (C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalCoord2fv.xml
func EvalCoord2fv(u *Float)  {
	C.goglEvalCoord2fv((*C.GLfloat)(u))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalMesh1.xml
func EvalMesh1(mode Enum, i1 Int, i2 Int)  {
	C.goglEvalMesh1((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalPoint1.xml
func EvalPoint1(i Int)  {
	C.goglEvalPoint1((C.GLint)(i))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalMesh2.xml
func EvalMesh2(mode Enum, i1 Int, i2 Int, j1 Int, j2 Int)  {
	C.goglEvalMesh2((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2), (C.GLint)(j1), (C.GLint)(j2))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glEvalPoint2.xml
func EvalPoint2(i Int, j Int)  {
	C.goglEvalPoint2((C.GLint)(i), (C.GLint)(j))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glAlphaFunc.xml
func AlphaFunc(func_ Enum, ref Float)  {
	C.goglAlphaFunc((C.GLenum)(func_), (C.GLfloat)(ref))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelZoom.xml
func PixelZoom(xfactor Float, yfactor Float)  {
	C.goglPixelZoom((C.GLfloat)(xfactor), (C.GLfloat)(yfactor))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelTransferf.xml
func PixelTransferf(pname Enum, param Float)  {
	C.goglPixelTransferf((C.GLenum)(pname), (C.GLfloat)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelTransferi.xml
func PixelTransferi(pname Enum, param Int)  {
	C.goglPixelTransferi((C.GLenum)(pname), (C.GLint)(param))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelMapfv.xml
func PixelMapfv(map_ Enum, mapsize Sizei, values *Float)  {
	C.goglPixelMapfv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLfloat)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelMapuiv.xml
func PixelMapuiv(map_ Enum, mapsize Sizei, values *Uint)  {
	C.goglPixelMapuiv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLuint)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPixelMapusv.xml
func PixelMapusv(map_ Enum, mapsize Sizei, values *Ushort)  {
	C.goglPixelMapusv((C.GLenum)(map_), (C.GLsizei)(mapsize), (*C.GLushort)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glCopyPixels.xml
func CopyPixels(x Int, y Int, width Sizei, height Sizei, type_ Enum)  {
	C.goglCopyPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(type_))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glDrawPixels.xml
func DrawPixels(width Sizei, height Sizei, format Enum, type_ Enum, pixels Pointer)  {
	C.goglDrawPixels((C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetClipPlane.xml
func GetClipPlane(plane Enum, equation *Double)  {
	C.goglGetClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetLightfv.xml
func GetLightfv(light Enum, pname Enum, params *Float)  {
	C.goglGetLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetLightiv.xml
func GetLightiv(light Enum, pname Enum, params *Int)  {
	C.goglGetLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMapdv.xml
func GetMapdv(target Enum, query Enum, v *Double)  {
	C.goglGetMapdv((C.GLenum)(target), (C.GLenum)(query), (*C.GLdouble)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMapfv.xml
func GetMapfv(target Enum, query Enum, v *Float)  {
	C.goglGetMapfv((C.GLenum)(target), (C.GLenum)(query), (*C.GLfloat)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMapiv.xml
func GetMapiv(target Enum, query Enum, v *Int)  {
	C.goglGetMapiv((C.GLenum)(target), (C.GLenum)(query), (*C.GLint)(v))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMaterialfv.xml
func GetMaterialfv(face Enum, pname Enum, params *Float)  {
	C.goglGetMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetMaterialiv.xml
func GetMaterialiv(face Enum, pname Enum, params *Int)  {
	C.goglGetMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPixelMapfv.xml
func GetPixelMapfv(map_ Enum, values *Float)  {
	C.goglGetPixelMapfv((C.GLenum)(map_), (*C.GLfloat)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPixelMapuiv.xml
func GetPixelMapuiv(map_ Enum, values *Uint)  {
	C.goglGetPixelMapuiv((C.GLenum)(map_), (*C.GLuint)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPixelMapusv.xml
func GetPixelMapusv(map_ Enum, values *Ushort)  {
	C.goglGetPixelMapusv((C.GLenum)(map_), (*C.GLushort)(values))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetPolygonStipple.xml
func GetPolygonStipple(mask *Ubyte)  {
	C.goglGetPolygonStipple((*C.GLubyte)(mask))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexEnvfv.xml
func GetTexEnvfv(target Enum, pname Enum, params *Float)  {
	C.goglGetTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexEnviv.xml
func GetTexEnviv(target Enum, pname Enum, params *Int)  {
	C.goglGetTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexGendv.xml
func GetTexGendv(coord Enum, pname Enum, params *Double)  {
	C.goglGetTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexGenfv.xml
func GetTexGenfv(coord Enum, pname Enum, params *Float)  {
	C.goglGetTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glGetTexGeniv.xml
func GetTexGeniv(coord Enum, pname Enum, params *Int)  {
	C.goglGetTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glIsList.xml
func IsList(list Uint) Boolean {
	return (Boolean)(C.goglIsList((C.GLuint)(list)))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glFrustum.xml
func Frustum(left Double, right Double, bottom Double, top Double, zNear Double, zFar Double)  {
	C.goglFrustum((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadIdentity.xml
func LoadIdentity()  {
	C.goglLoadIdentity()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadMatrixf.xml
func LoadMatrixf(m *Float)  {
	C.goglLoadMatrixf((*C.GLfloat)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glLoadMatrixd.xml
func LoadMatrixd(m *Double)  {
	C.goglLoadMatrixd((*C.GLdouble)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMatrixMode.xml
func MatrixMode(mode Enum)  {
	C.goglMatrixMode((C.GLenum)(mode))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultMatrixf.xml
func MultMatrixf(m *Float)  {
	C.goglMultMatrixf((*C.GLfloat)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glMultMatrixd.xml
func MultMatrixd(m *Double)  {
	C.goglMultMatrixd((*C.GLdouble)(m))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glOrtho.xml
func Ortho(left Double, right Double, bottom Double, top Double, zNear Double, zFar Double)  {
	C.goglOrtho((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPopMatrix.xml
func PopMatrix()  {
	C.goglPopMatrix()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glPushMatrix.xml
func PushMatrix()  {
	C.goglPushMatrix()
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRotated.xml
func Rotated(angle Double, x Double, y Double, z Double)  {
	C.goglRotated((C.GLdouble)(angle), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glRotatef.xml
func Rotatef(angle Float, x Float, y Float, z Float)  {
	C.goglRotatef((C.GLfloat)(angle), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glScaled.xml
func Scaled(x Double, y Double, z Double)  {
	C.goglScaled((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glScalef.xml
func Scalef(x Float, y Float, z Float)  {
	C.goglScalef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTranslated.xml
func Translated(x Double, y Double, z Double)  {
	C.goglTranslated((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
// http://www.opengl.org/sdk/docs/man/xhtml/glTranslatef.xml
func Translatef(x Float, y Float, z Float)  {
	C.goglTranslatef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func InitVersion10() error {
	var ret C.int
	if ret = C.init_VERSION_1_0(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0")
	}
	return nil
}
func InitVersion10Deprecated() error {
	var ret C.int
	if ret = C.init_VERSION_1_0_DEPRECATED(); ret != 0 {
		return errors.New("unable to initialize VERSION_1_0_DEPRECATED")
	}
	return nil
}
func Init() error {
	var err error
	if err = InitVersion10(); err != nil {
		return err
	}
	if err = InitVersion10Deprecated(); err != nil {
		return err
	}
	return nil
}
//Go bool to GL boolean.
func GLBool(b bool) Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

// GL boolean to Go bool.
func GoBool(b Boolean) bool {
	return b == TRUE
}

// Go string to GL string.
func GLString(str string) *Char {
	return (*Char)(C.CString(str))
}

// Allocates a GL string.
func GLStringAlloc(length Sizei) *Char {
	return (*Char)(C.malloc(C.size_t(length)))
}

// Frees GL string.
func GLStringFree(str *Char) {
	C.free(unsafe.Pointer(str))
}

// GL string (GLchar*) to Go string.
func GoString(str *Char) string {
	return C.GoString((*C.char)(str))
}

// GL string (GLubyte*) to Go string.
func GoStringUb(str *Ubyte) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

// GL string (GLchar*) with length to Go string.
func GoStringN(str *Char, length Sizei) string {
	return C.GoStringN((*C.char)(str), C.int(length))
}

// Converts a list of Go strings to a slice of GL strings.
// Usefull for ShaderSource().
func GLStringArray(strs ...string) []*Char {
	strSlice := make([]*Char, len(strs))
	for i, s := range strs {
		strSlice[i] = (*Char)(C.CString(s))
	}
	return strSlice
}

// Free GL string slice allocated by GLStringArray().
func GLStringArrayFree(strs []*Char) {
	for _, s := range strs {
		C.free(unsafe.Pointer(s))
	}
}

// Add offset to a pointer. Usefull for VertexAttribPointer, TexCoordPointer, NormalPointer, ... 
func Offset(p Pointer, o uintptr) Pointer {
	return Pointer(uintptr(p) + o)
}

// EOF