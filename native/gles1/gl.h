// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// This source file was automatically generated using glwrap.
// +build<no value>

#ifndef H_GLWRAP_OPENGL_ES_1
#define H_GLWRAP_OPENGL_ES_1


// OpenGL type definitions
#include <stddef.h>
#include <KHR/khrplatform.h>
#ifndef GLEXT_64_TYPES_DEFINED
/* This code block is duplicated in glxext.h, so must be protected */
#define GLEXT_64_TYPES_DEFINED
/* Define int32_t, int64_t, and uint64_t types for UST/MSC */
/* (as used in the GL_EXT_timer_query extension). */
#if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
#include <inttypes.h>
#elif defined(__sun__) || defined(__digital__)
#include <inttypes.h>
#if defined(__STDC__)
#if defined(__arch64__) || defined(_LP64)
typedef long int int64_t;
typedef unsigned long int uint64_t;
#else
typedef long long int int64_t;
typedef unsigned long long int uint64_t;
#endif /* __arch64__ */
#endif /* __STDC__ */
#elif defined( __VMS ) || defined(__sgi)
#include <inttypes.h>
#elif defined(__SCO__) || defined(__USLC__)
#include <stdint.h>
#elif defined(__UNIXOS2__) || defined(__SOL64__)
typedef long int int32_t;
typedef long long int int64_t;
typedef unsigned long long int uint64_t;
#elif defined(_WIN32) && defined(__GNUC__)
#include <stdint.h>
#elif defined(_WIN32)
typedef __int32 int32_t;
typedef __int64 int64_t;
typedef unsigned __int64 uint64_t;
#else
/* Fallback if nothing above works */
#include <inttypes.h>
#endif
#endif
typedef unsigned int GLenum;
typedef unsigned char GLboolean;
typedef unsigned int GLbitfield;
typedef void GLvoid;
typedef signed char GLbyte;
typedef short GLshort;
typedef int GLint;
typedef int GLclampx;
typedef unsigned char GLubyte;
typedef unsigned short GLushort;
typedef unsigned int GLuint;
typedef int GLsizei;
typedef float GLfloat;
typedef float GLclampf;
typedef double GLdouble;
typedef double GLclampd;
typedef void *GLeglImageOES;
typedef char GLchar;
typedef char GLcharARB;
#ifdef __APPLE__
typedef void *GLhandleARB;
#else
typedef unsigned int GLhandleARB;
#endif
typedef unsigned short GLhalfARB;
typedef unsigned short GLhalf;
typedef GLint GLfixed;
typedef ptrdiff_t GLintptr;
typedef ptrdiff_t GLsizeiptr;
typedef int64_t GLint64;
typedef uint64_t GLuint64;
typedef ptrdiff_t GLintptrARB;
typedef ptrdiff_t GLsizeiptrARB;
typedef int64_t GLint64EXT;
typedef uint64_t GLuint64EXT;
typedef struct __GLsync *GLsync;
struct _cl_context;
struct _cl_event;
typedef void ( *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
typedef void ( *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
typedef void ( *GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
typedef khronos_int32_t GLclampx;
typedef khronos_int8_t GLbyte;
typedef khronos_uint8_t GLubyte;
typedef khronos_float_t GLfloat;
typedef khronos_float_t GLclampf;
typedef khronos_int32_t GLfixed;
typedef khronos_int64_t GLint64;
typedef khronos_uint64_t GLuint64;
typedef khronos_intptr_t GLintptr;
typedef khronos_ssize_t GLsizeiptr;
typedef void ( *GLDEBUGPROCAMD)(GLuint id,GLenum category,GLenum severity,GLsizei length,const GLchar *message,void *userParam);
typedef unsigned short GLhalfNV;
typedef GLintptr GLvdpauSurfaceNV;



// Define GLAPIENTRY for platforms that need it.
#ifndef GLAPIENTRY
#define GLAPIENTRY
#endif


// Typedefs for each OpenGL pointer function; GLAPIENTRY is required on Windows
// OS (but for other platforms is simply nothing).
typedef void (GLAPIENTRY* PFNGLALPHAFUNCPROC) (GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLCLEARCOLORPROC) (GLfloat, GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLCLEARDEPTHFPROC) (GLfloat);
typedef void (GLAPIENTRY* PFNGLCLIPPLANEFPROC) (GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLCOLOR4FPROC) (GLfloat, GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLDEPTHRANGEFPROC) (GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLFOGFPROC) (GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLFOGFVPROC) (GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLFRUSTUMFPROC) (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLGETCLIPPLANEFPROC) (GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLGETFLOATVPROC) (GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLGETLIGHTFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLGETMATERIALFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLGETTEXENVFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLGETTEXPARAMETERFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLLIGHTMODELFPROC) (GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLLIGHTMODELFVPROC) (GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLLIGHTFPROC) (GLenum, GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLLIGHTFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLLINEWIDTHPROC) (GLfloat);
typedef void (GLAPIENTRY* PFNGLLOADMATRIXFPROC) (GLfloat*);
typedef void (GLAPIENTRY* PFNGLMATERIALFPROC) (GLenum, GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLMATERIALFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLMULTMATRIXFPROC) (GLfloat*);
typedef void (GLAPIENTRY* PFNGLMULTITEXCOORD4FPROC) (GLenum, GLfloat, GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLNORMAL3FPROC) (GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLORTHOFPROC) (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLPOINTPARAMETERFPROC) (GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLPOINTPARAMETERFVPROC) (GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLPOINTSIZEPROC) (GLfloat);
typedef void (GLAPIENTRY* PFNGLPOLYGONOFFSETPROC) (GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLROTATEFPROC) (GLfloat, GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLSCALEFPROC) (GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLTEXENVFPROC) (GLenum, GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLTEXENVFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLTEXPARAMETERFPROC) (GLenum, GLenum, GLfloat);
typedef void (GLAPIENTRY* PFNGLTEXPARAMETERFVPROC) (GLenum, GLenum, GLfloat*);
typedef void (GLAPIENTRY* PFNGLTRANSLATEFPROC) (GLfloat, GLfloat, GLfloat);
typedef void (GLAPIENTRY* PFNGLACTIVETEXTUREPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLALPHAFUNCXPROC) (GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLBINDBUFFERPROC) (GLenum, GLuint);
typedef void (GLAPIENTRY* PFNGLBINDTEXTUREPROC) (GLenum, GLuint);
typedef void (GLAPIENTRY* PFNGLBLENDFUNCPROC) (GLenum, GLenum);
typedef void (GLAPIENTRY* PFNGLBUFFERDATAPROC) (GLenum, GLsizeiptr, void*, GLenum);
typedef void (GLAPIENTRY* PFNGLBUFFERSUBDATAPROC) (GLenum, GLintptr, GLsizeiptr, void*);
typedef void (GLAPIENTRY* PFNGLCLEARPROC) (GLbitfield);
typedef void (GLAPIENTRY* PFNGLCLEARCOLORXPROC) (GLfixed, GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLCLEARDEPTHXPROC) (GLfixed);
typedef void (GLAPIENTRY* PFNGLCLEARSTENCILPROC) (GLint);
typedef void (GLAPIENTRY* PFNGLCLIENTACTIVETEXTUREPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLCLIPPLANEXPROC) (GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLCOLOR4UBPROC) (GLubyte, GLubyte, GLubyte, GLubyte);
typedef void (GLAPIENTRY* PFNGLCOLOR4XPROC) (GLfixed, GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLCOLORMASKPROC) (GLboolean, GLboolean, GLboolean, GLboolean);
typedef void (GLAPIENTRY* PFNGLCOLORPOINTERPROC) (GLint, GLenum, GLsizei, void*);
typedef void (GLAPIENTRY* PFNGLCOMPRESSEDTEXIMAGE2DPROC) (GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLsizei, void*);
typedef void (GLAPIENTRY* PFNGLCOMPRESSEDTEXSUBIMAGE2DPROC) (GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLsizei, void*);
typedef void (GLAPIENTRY* PFNGLCOPYTEXIMAGE2DPROC) (GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLsizei, GLint);
typedef void (GLAPIENTRY* PFNGLCOPYTEXSUBIMAGE2DPROC) (GLenum, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei);
typedef void (GLAPIENTRY* PFNGLCULLFACEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLDELETEBUFFERSPROC) (GLsizei, GLuint*);
typedef void (GLAPIENTRY* PFNGLDELETETEXTURESPROC) (GLsizei, GLuint*);
typedef void (GLAPIENTRY* PFNGLDEPTHFUNCPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLDEPTHMASKPROC) (GLboolean);
typedef void (GLAPIENTRY* PFNGLDEPTHRANGEXPROC) (GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLDISABLEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLDISABLECLIENTSTATEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLDRAWARRAYSPROC) (GLenum, GLint, GLsizei);
typedef void (GLAPIENTRY* PFNGLDRAWELEMENTSPROC) (GLenum, GLsizei, GLenum, void*);
typedef void (GLAPIENTRY* PFNGLENABLEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLENABLECLIENTSTATEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLFINISHPROC) (void);
typedef void (GLAPIENTRY* PFNGLFLUSHPROC) (void);
typedef void (GLAPIENTRY* PFNGLFOGXPROC) (GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLFOGXVPROC) (GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLFRONTFACEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLFRUSTUMXPROC) (GLfixed, GLfixed, GLfixed, GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLGETBOOLEANVPROC) (GLenum, GLboolean*);
typedef void (GLAPIENTRY* PFNGLGETBUFFERPARAMETERIVPROC) (GLenum, GLenum, GLint*);
typedef void (GLAPIENTRY* PFNGLGETCLIPPLANEXPROC) (GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLGENBUFFERSPROC) (GLsizei, GLuint*);
typedef void (GLAPIENTRY* PFNGLGENTEXTURESPROC) (GLsizei, GLuint*);
typedef GLenum (GLAPIENTRY* PFNGLGETERRORPROC) (void);
typedef void (GLAPIENTRY* PFNGLGETFIXEDVPROC) (GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLGETINTEGERVPROC) (GLenum, GLint*);
typedef void (GLAPIENTRY* PFNGLGETLIGHTXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLGETMATERIALXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLGETPOINTERVPROC) (GLenum, void**);
typedef GLubyte* (GLAPIENTRY* PFNGLGETSTRINGPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLGETTEXENVIVPROC) (GLenum, GLenum, GLint*);
typedef void (GLAPIENTRY* PFNGLGETTEXENVXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLGETTEXPARAMETERIVPROC) (GLenum, GLenum, GLint*);
typedef void (GLAPIENTRY* PFNGLGETTEXPARAMETERXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLHINTPROC) (GLenum, GLenum);
typedef GLboolean (GLAPIENTRY* PFNGLISBUFFERPROC) (GLuint);
typedef GLboolean (GLAPIENTRY* PFNGLISENABLEDPROC) (GLenum);
typedef GLboolean (GLAPIENTRY* PFNGLISTEXTUREPROC) (GLuint);
typedef void (GLAPIENTRY* PFNGLLIGHTMODELXPROC) (GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLLIGHTMODELXVPROC) (GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLLIGHTXPROC) (GLenum, GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLLIGHTXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLLINEWIDTHXPROC) (GLfixed);
typedef void (GLAPIENTRY* PFNGLLOADIDENTITYPROC) (void);
typedef void (GLAPIENTRY* PFNGLLOADMATRIXXPROC) (GLfixed*);
typedef void (GLAPIENTRY* PFNGLLOGICOPPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLMATERIALXPROC) (GLenum, GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLMATERIALXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLMATRIXMODEPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLMULTMATRIXXPROC) (GLfixed*);
typedef void (GLAPIENTRY* PFNGLMULTITEXCOORD4XPROC) (GLenum, GLfixed, GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLNORMAL3XPROC) (GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLNORMALPOINTERPROC) (GLenum, GLsizei, void*);
typedef void (GLAPIENTRY* PFNGLORTHOXPROC) (GLfixed, GLfixed, GLfixed, GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLPIXELSTOREIPROC) (GLenum, GLint);
typedef void (GLAPIENTRY* PFNGLPOINTPARAMETERXPROC) (GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLPOINTPARAMETERXVPROC) (GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLPOINTSIZEXPROC) (GLfixed);
typedef void (GLAPIENTRY* PFNGLPOLYGONOFFSETXPROC) (GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLPOPMATRIXPROC) (void);
typedef void (GLAPIENTRY* PFNGLPUSHMATRIXPROC) (void);
typedef void (GLAPIENTRY* PFNGLREADPIXELSPROC) (GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, void*);
typedef void (GLAPIENTRY* PFNGLROTATEXPROC) (GLfixed, GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLSAMPLECOVERAGEPROC) (GLfloat, GLboolean);
typedef void (GLAPIENTRY* PFNGLSAMPLECOVERAGEXPROC) (GLclampx, GLboolean);
typedef void (GLAPIENTRY* PFNGLSCALEXPROC) (GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLSCISSORPROC) (GLint, GLint, GLsizei, GLsizei);
typedef void (GLAPIENTRY* PFNGLSHADEMODELPROC) (GLenum);
typedef void (GLAPIENTRY* PFNGLSTENCILFUNCPROC) (GLenum, GLint, GLuint);
typedef void (GLAPIENTRY* PFNGLSTENCILMASKPROC) (GLuint);
typedef void (GLAPIENTRY* PFNGLSTENCILOPPROC) (GLenum, GLenum, GLenum);
typedef void (GLAPIENTRY* PFNGLTEXCOORDPOINTERPROC) (GLint, GLenum, GLsizei, void*);
typedef void (GLAPIENTRY* PFNGLTEXENVIPROC) (GLenum, GLenum, GLint);
typedef void (GLAPIENTRY* PFNGLTEXENVXPROC) (GLenum, GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLTEXENVIVPROC) (GLenum, GLenum, GLint*);
typedef void (GLAPIENTRY* PFNGLTEXENVXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLTEXIMAGE2DPROC) (GLenum, GLint, GLint, GLsizei, GLsizei, GLint, GLenum, GLenum, void*);
typedef void (GLAPIENTRY* PFNGLTEXPARAMETERIPROC) (GLenum, GLenum, GLint);
typedef void (GLAPIENTRY* PFNGLTEXPARAMETERXPROC) (GLenum, GLenum, GLfixed);
typedef void (GLAPIENTRY* PFNGLTEXPARAMETERIVPROC) (GLenum, GLenum, GLint*);
typedef void (GLAPIENTRY* PFNGLTEXPARAMETERXVPROC) (GLenum, GLenum, GLfixed*);
typedef void (GLAPIENTRY* PFNGLTEXSUBIMAGE2DPROC) (GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, void*);
typedef void (GLAPIENTRY* PFNGLTRANSLATEXPROC) (GLfixed, GLfixed, GLfixed);
typedef void (GLAPIENTRY* PFNGLVERTEXPOINTERPROC) (GLint, GLenum, GLsizei, void*);
typedef void (GLAPIENTRY* PFNGLVIEWPORTPROC) (GLint, GLint, GLsizei, GLsizei);


// gl_wrap_context holds the loaded function pointers for a given OpenGL
// context. Pointers are lazily-loaded upon first invocation of each function.
// through their appropriate gl_wrap_context_glFunction() counterparts.
typedef struct {
	PFNGLALPHAFUNCPROC glAlphaFuncProc;
	PFNGLCLEARCOLORPROC glClearColorProc;
	PFNGLCLEARDEPTHFPROC glClearDepthfProc;
	PFNGLCLIPPLANEFPROC glClipPlanefProc;
	PFNGLCOLOR4FPROC glColor4fProc;
	PFNGLDEPTHRANGEFPROC glDepthRangefProc;
	PFNGLFOGFPROC glFogfProc;
	PFNGLFOGFVPROC glFogfvProc;
	PFNGLFRUSTUMFPROC glFrustumfProc;
	PFNGLGETCLIPPLANEFPROC glGetClipPlanefProc;
	PFNGLGETFLOATVPROC glGetFloatvProc;
	PFNGLGETLIGHTFVPROC glGetLightfvProc;
	PFNGLGETMATERIALFVPROC glGetMaterialfvProc;
	PFNGLGETTEXENVFVPROC glGetTexEnvfvProc;
	PFNGLGETTEXPARAMETERFVPROC glGetTexParameterfvProc;
	PFNGLLIGHTMODELFPROC glLightModelfProc;
	PFNGLLIGHTMODELFVPROC glLightModelfvProc;
	PFNGLLIGHTFPROC glLightfProc;
	PFNGLLIGHTFVPROC glLightfvProc;
	PFNGLLINEWIDTHPROC glLineWidthProc;
	PFNGLLOADMATRIXFPROC glLoadMatrixfProc;
	PFNGLMATERIALFPROC glMaterialfProc;
	PFNGLMATERIALFVPROC glMaterialfvProc;
	PFNGLMULTMATRIXFPROC glMultMatrixfProc;
	PFNGLMULTITEXCOORD4FPROC glMultiTexCoord4fProc;
	PFNGLNORMAL3FPROC glNormal3fProc;
	PFNGLORTHOFPROC glOrthofProc;
	PFNGLPOINTPARAMETERFPROC glPointParameterfProc;
	PFNGLPOINTPARAMETERFVPROC glPointParameterfvProc;
	PFNGLPOINTSIZEPROC glPointSizeProc;
	PFNGLPOLYGONOFFSETPROC glPolygonOffsetProc;
	PFNGLROTATEFPROC glRotatefProc;
	PFNGLSCALEFPROC glScalefProc;
	PFNGLTEXENVFPROC glTexEnvfProc;
	PFNGLTEXENVFVPROC glTexEnvfvProc;
	PFNGLTEXPARAMETERFPROC glTexParameterfProc;
	PFNGLTEXPARAMETERFVPROC glTexParameterfvProc;
	PFNGLTRANSLATEFPROC glTranslatefProc;
	PFNGLACTIVETEXTUREPROC glActiveTextureProc;
	PFNGLALPHAFUNCXPROC glAlphaFuncxProc;
	PFNGLBINDBUFFERPROC glBindBufferProc;
	PFNGLBINDTEXTUREPROC glBindTextureProc;
	PFNGLBLENDFUNCPROC glBlendFuncProc;
	PFNGLBUFFERDATAPROC glBufferDataProc;
	PFNGLBUFFERSUBDATAPROC glBufferSubDataProc;
	PFNGLCLEARPROC glClearProc;
	PFNGLCLEARCOLORXPROC glClearColorxProc;
	PFNGLCLEARDEPTHXPROC glClearDepthxProc;
	PFNGLCLEARSTENCILPROC glClearStencilProc;
	PFNGLCLIENTACTIVETEXTUREPROC glClientActiveTextureProc;
	PFNGLCLIPPLANEXPROC glClipPlanexProc;
	PFNGLCOLOR4UBPROC glColor4ubProc;
	PFNGLCOLOR4XPROC glColor4xProc;
	PFNGLCOLORMASKPROC glColorMaskProc;
	PFNGLCOLORPOINTERPROC glColorPointerProc;
	PFNGLCOMPRESSEDTEXIMAGE2DPROC glCompressedTexImage2DProc;
	PFNGLCOMPRESSEDTEXSUBIMAGE2DPROC glCompressedTexSubImage2DProc;
	PFNGLCOPYTEXIMAGE2DPROC glCopyTexImage2DProc;
	PFNGLCOPYTEXSUBIMAGE2DPROC glCopyTexSubImage2DProc;
	PFNGLCULLFACEPROC glCullFaceProc;
	PFNGLDELETEBUFFERSPROC glDeleteBuffersProc;
	PFNGLDELETETEXTURESPROC glDeleteTexturesProc;
	PFNGLDEPTHFUNCPROC glDepthFuncProc;
	PFNGLDEPTHMASKPROC glDepthMaskProc;
	PFNGLDEPTHRANGEXPROC glDepthRangexProc;
	PFNGLDISABLEPROC glDisableProc;
	PFNGLDISABLECLIENTSTATEPROC glDisableClientStateProc;
	PFNGLDRAWARRAYSPROC glDrawArraysProc;
	PFNGLDRAWELEMENTSPROC glDrawElementsProc;
	PFNGLENABLEPROC glEnableProc;
	PFNGLENABLECLIENTSTATEPROC glEnableClientStateProc;
	PFNGLFINISHPROC glFinishProc;
	PFNGLFLUSHPROC glFlushProc;
	PFNGLFOGXPROC glFogxProc;
	PFNGLFOGXVPROC glFogxvProc;
	PFNGLFRONTFACEPROC glFrontFaceProc;
	PFNGLFRUSTUMXPROC glFrustumxProc;
	PFNGLGETBOOLEANVPROC glGetBooleanvProc;
	PFNGLGETBUFFERPARAMETERIVPROC glGetBufferParameterivProc;
	PFNGLGETCLIPPLANEXPROC glGetClipPlanexProc;
	PFNGLGENBUFFERSPROC glGenBuffersProc;
	PFNGLGENTEXTURESPROC glGenTexturesProc;
	PFNGLGETERRORPROC glGetErrorProc;
	PFNGLGETFIXEDVPROC glGetFixedvProc;
	PFNGLGETINTEGERVPROC glGetIntegervProc;
	PFNGLGETLIGHTXVPROC glGetLightxvProc;
	PFNGLGETMATERIALXVPROC glGetMaterialxvProc;
	PFNGLGETPOINTERVPROC glGetPointervProc;
	PFNGLGETSTRINGPROC glGetStringProc;
	PFNGLGETTEXENVIVPROC glGetTexEnvivProc;
	PFNGLGETTEXENVXVPROC glGetTexEnvxvProc;
	PFNGLGETTEXPARAMETERIVPROC glGetTexParameterivProc;
	PFNGLGETTEXPARAMETERXVPROC glGetTexParameterxvProc;
	PFNGLHINTPROC glHintProc;
	PFNGLISBUFFERPROC glIsBufferProc;
	PFNGLISENABLEDPROC glIsEnabledProc;
	PFNGLISTEXTUREPROC glIsTextureProc;
	PFNGLLIGHTMODELXPROC glLightModelxProc;
	PFNGLLIGHTMODELXVPROC glLightModelxvProc;
	PFNGLLIGHTXPROC glLightxProc;
	PFNGLLIGHTXVPROC glLightxvProc;
	PFNGLLINEWIDTHXPROC glLineWidthxProc;
	PFNGLLOADIDENTITYPROC glLoadIdentityProc;
	PFNGLLOADMATRIXXPROC glLoadMatrixxProc;
	PFNGLLOGICOPPROC glLogicOpProc;
	PFNGLMATERIALXPROC glMaterialxProc;
	PFNGLMATERIALXVPROC glMaterialxvProc;
	PFNGLMATRIXMODEPROC glMatrixModeProc;
	PFNGLMULTMATRIXXPROC glMultMatrixxProc;
	PFNGLMULTITEXCOORD4XPROC glMultiTexCoord4xProc;
	PFNGLNORMAL3XPROC glNormal3xProc;
	PFNGLNORMALPOINTERPROC glNormalPointerProc;
	PFNGLORTHOXPROC glOrthoxProc;
	PFNGLPIXELSTOREIPROC glPixelStoreiProc;
	PFNGLPOINTPARAMETERXPROC glPointParameterxProc;
	PFNGLPOINTPARAMETERXVPROC glPointParameterxvProc;
	PFNGLPOINTSIZEXPROC glPointSizexProc;
	PFNGLPOLYGONOFFSETXPROC glPolygonOffsetxProc;
	PFNGLPOPMATRIXPROC glPopMatrixProc;
	PFNGLPUSHMATRIXPROC glPushMatrixProc;
	PFNGLREADPIXELSPROC glReadPixelsProc;
	PFNGLROTATEXPROC glRotatexProc;
	PFNGLSAMPLECOVERAGEPROC glSampleCoverageProc;
	PFNGLSAMPLECOVERAGEXPROC glSampleCoveragexProc;
	PFNGLSCALEXPROC glScalexProc;
	PFNGLSCISSORPROC glScissorProc;
	PFNGLSHADEMODELPROC glShadeModelProc;
	PFNGLSTENCILFUNCPROC glStencilFuncProc;
	PFNGLSTENCILMASKPROC glStencilMaskProc;
	PFNGLSTENCILOPPROC glStencilOpProc;
	PFNGLTEXCOORDPOINTERPROC glTexCoordPointerProc;
	PFNGLTEXENVIPROC glTexEnviProc;
	PFNGLTEXENVXPROC glTexEnvxProc;
	PFNGLTEXENVIVPROC glTexEnvivProc;
	PFNGLTEXENVXVPROC glTexEnvxvProc;
	PFNGLTEXIMAGE2DPROC glTexImage2DProc;
	PFNGLTEXPARAMETERIPROC glTexParameteriProc;
	PFNGLTEXPARAMETERXPROC glTexParameterxProc;
	PFNGLTEXPARAMETERIVPROC glTexParameterivProc;
	PFNGLTEXPARAMETERXVPROC glTexParameterxvProc;
	PFNGLTEXSUBIMAGE2DPROC glTexSubImage2DProc;
	PFNGLTRANSLATEXPROC glTranslatexProc;
	PFNGLVERTEXPOINTERPROC glVertexPointerProc;
	PFNGLVIEWPORTPROC glViewportProc;

} gl_wrap_context;

// Function prototype definitions for each appropriate OpenGL function.
inline void gl_wrap_context_glAlphaFunc(gl_wrap_context* ctx, GLenum func, GLfloat ref);
inline void gl_wrap_context_glClearColor(gl_wrap_context* ctx, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
inline void gl_wrap_context_glClearDepthf(gl_wrap_context* ctx, GLfloat d);
inline void gl_wrap_context_glClipPlanef(gl_wrap_context* ctx, GLenum p, GLfloat* eqn);
inline void gl_wrap_context_glColor4f(gl_wrap_context* ctx, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
inline void gl_wrap_context_glDepthRangef(gl_wrap_context* ctx, GLfloat n, GLfloat f);
inline void gl_wrap_context_glFogf(gl_wrap_context* ctx, GLenum pname, GLfloat param);
inline void gl_wrap_context_glFogfv(gl_wrap_context* ctx, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glFrustumf(gl_wrap_context* ctx, GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f);
inline void gl_wrap_context_glGetClipPlanef(gl_wrap_context* ctx, GLenum plane, GLfloat* equation);
inline void gl_wrap_context_glGetFloatv(gl_wrap_context* ctx, GLenum pname, GLfloat* data);
inline void gl_wrap_context_glGetLightfv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glGetMaterialfv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glGetTexEnvfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glGetTexParameterfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glLightModelf(gl_wrap_context* ctx, GLenum pname, GLfloat param);
inline void gl_wrap_context_glLightModelfv(gl_wrap_context* ctx, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glLightf(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfloat param);
inline void gl_wrap_context_glLightfv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glLineWidth(gl_wrap_context* ctx, GLfloat width);
inline void gl_wrap_context_glLoadMatrixf(gl_wrap_context* ctx, GLfloat* m);
inline void gl_wrap_context_glMaterialf(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfloat param);
inline void gl_wrap_context_glMaterialfv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glMultMatrixf(gl_wrap_context* ctx, GLfloat* m);
inline void gl_wrap_context_glMultiTexCoord4f(gl_wrap_context* ctx, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q);
inline void gl_wrap_context_glNormal3f(gl_wrap_context* ctx, GLfloat nx, GLfloat ny, GLfloat nz);
inline void gl_wrap_context_glOrthof(gl_wrap_context* ctx, GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f);
inline void gl_wrap_context_glPointParameterf(gl_wrap_context* ctx, GLenum pname, GLfloat param);
inline void gl_wrap_context_glPointParameterfv(gl_wrap_context* ctx, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glPointSize(gl_wrap_context* ctx, GLfloat size);
inline void gl_wrap_context_glPolygonOffset(gl_wrap_context* ctx, GLfloat factor, GLfloat units);
inline void gl_wrap_context_glRotatef(gl_wrap_context* ctx, GLfloat angle, GLfloat x, GLfloat y, GLfloat z);
inline void gl_wrap_context_glScalef(gl_wrap_context* ctx, GLfloat x, GLfloat y, GLfloat z);
inline void gl_wrap_context_glTexEnvf(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat param);
inline void gl_wrap_context_glTexEnvfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glTexParameterf(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat param);
inline void gl_wrap_context_glTexParameterfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params);
inline void gl_wrap_context_glTranslatef(gl_wrap_context* ctx, GLfloat x, GLfloat y, GLfloat z);
inline void gl_wrap_context_glActiveTexture(gl_wrap_context* ctx, GLenum texture);
inline void gl_wrap_context_glAlphaFuncx(gl_wrap_context* ctx, GLenum func, GLfixed ref);
inline void gl_wrap_context_glBindBuffer(gl_wrap_context* ctx, GLenum target, GLuint buffer);
inline void gl_wrap_context_glBindTexture(gl_wrap_context* ctx, GLenum target, GLuint texture);
inline void gl_wrap_context_glBlendFunc(gl_wrap_context* ctx, GLenum sfactor, GLenum dfactor);
inline void gl_wrap_context_glBufferData(gl_wrap_context* ctx, GLenum target, GLsizeiptr size, void* data, GLenum usage);
inline void gl_wrap_context_glBufferSubData(gl_wrap_context* ctx, GLenum target, GLintptr offset, GLsizeiptr size, void* data);
inline void gl_wrap_context_glClear(gl_wrap_context* ctx, GLbitfield mask);
inline void gl_wrap_context_glClearColorx(gl_wrap_context* ctx, GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha);
inline void gl_wrap_context_glClearDepthx(gl_wrap_context* ctx, GLfixed depth);
inline void gl_wrap_context_glClearStencil(gl_wrap_context* ctx, GLint s);
inline void gl_wrap_context_glClientActiveTexture(gl_wrap_context* ctx, GLenum texture);
inline void gl_wrap_context_glClipPlanex(gl_wrap_context* ctx, GLenum plane, GLfixed* equation);
inline void gl_wrap_context_glColor4ub(gl_wrap_context* ctx, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha);
inline void gl_wrap_context_glColor4x(gl_wrap_context* ctx, GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha);
inline void gl_wrap_context_glColorMask(gl_wrap_context* ctx, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
inline void gl_wrap_context_glColorPointer(gl_wrap_context* ctx, GLint size, GLenum type, GLsizei stride, void* pointer);
inline void gl_wrap_context_glCompressedTexImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, void* data);
inline void gl_wrap_context_glCompressedTexSubImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, void* data);
inline void gl_wrap_context_glCopyTexImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
inline void gl_wrap_context_glCopyTexSubImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
inline void gl_wrap_context_glCullFace(gl_wrap_context* ctx, GLenum mode);
inline void gl_wrap_context_glDeleteBuffers(gl_wrap_context* ctx, GLsizei n, GLuint* buffers);
inline void gl_wrap_context_glDeleteTextures(gl_wrap_context* ctx, GLsizei n, GLuint* textures);
inline void gl_wrap_context_glDepthFunc(gl_wrap_context* ctx, GLenum func);
inline void gl_wrap_context_glDepthMask(gl_wrap_context* ctx, GLboolean flag);
inline void gl_wrap_context_glDepthRangex(gl_wrap_context* ctx, GLfixed n, GLfixed f);
inline void gl_wrap_context_glDisable(gl_wrap_context* ctx, GLenum cap);
inline void gl_wrap_context_glDisableClientState(gl_wrap_context* ctx, GLenum array);
inline void gl_wrap_context_glDrawArrays(gl_wrap_context* ctx, GLenum mode, GLint first, GLsizei count);
inline void gl_wrap_context_glDrawElements(gl_wrap_context* ctx, GLenum mode, GLsizei count, GLenum type, void* indices);
inline void gl_wrap_context_glEnable(gl_wrap_context* ctx, GLenum cap);
inline void gl_wrap_context_glEnableClientState(gl_wrap_context* ctx, GLenum array);
inline void gl_wrap_context_glFinish(gl_wrap_context* ctxvoid);
inline void gl_wrap_context_glFlush(gl_wrap_context* ctxvoid);
inline void gl_wrap_context_glFogx(gl_wrap_context* ctx, GLenum pname, GLfixed param);
inline void gl_wrap_context_glFogxv(gl_wrap_context* ctx, GLenum pname, GLfixed* param);
inline void gl_wrap_context_glFrontFace(gl_wrap_context* ctx, GLenum mode);
inline void gl_wrap_context_glFrustumx(gl_wrap_context* ctx, GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f);
inline void gl_wrap_context_glGetBooleanv(gl_wrap_context* ctx, GLenum pname, GLboolean* data);
inline void gl_wrap_context_glGetBufferParameteriv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params);
inline void gl_wrap_context_glGetClipPlanex(gl_wrap_context* ctx, GLenum plane, GLfixed* equation);
inline void gl_wrap_context_glGenBuffers(gl_wrap_context* ctx, GLsizei n, GLuint* buffers);
inline void gl_wrap_context_glGenTextures(gl_wrap_context* ctx, GLsizei n, GLuint* textures);
inline GLenum gl_wrap_context_glGetError(gl_wrap_context* ctxvoid);
inline void gl_wrap_context_glGetFixedv(gl_wrap_context* ctx, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glGetIntegerv(gl_wrap_context* ctx, GLenum pname, GLint* data);
inline void gl_wrap_context_glGetLightxv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glGetMaterialxv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glGetPointerv(gl_wrap_context* ctx, GLenum pname, void** params);
inline GLubyte* gl_wrap_context_glGetString(gl_wrap_context* ctx, GLenum name);
inline void gl_wrap_context_glGetTexEnviv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params);
inline void gl_wrap_context_glGetTexEnvxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glGetTexParameteriv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params);
inline void gl_wrap_context_glGetTexParameterxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glHint(gl_wrap_context* ctx, GLenum target, GLenum mode);
inline GLboolean gl_wrap_context_glIsBuffer(gl_wrap_context* ctx, GLuint buffer);
inline GLboolean gl_wrap_context_glIsEnabled(gl_wrap_context* ctx, GLenum cap);
inline GLboolean gl_wrap_context_glIsTexture(gl_wrap_context* ctx, GLuint texture);
inline void gl_wrap_context_glLightModelx(gl_wrap_context* ctx, GLenum pname, GLfixed param);
inline void gl_wrap_context_glLightModelxv(gl_wrap_context* ctx, GLenum pname, GLfixed* param);
inline void gl_wrap_context_glLightx(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfixed param);
inline void gl_wrap_context_glLightxv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glLineWidthx(gl_wrap_context* ctx, GLfixed width);
inline void gl_wrap_context_glLoadIdentity(gl_wrap_context* ctxvoid);
inline void gl_wrap_context_glLoadMatrixx(gl_wrap_context* ctx, GLfixed* m);
inline void gl_wrap_context_glLogicOp(gl_wrap_context* ctx, GLenum opcode);
inline void gl_wrap_context_glMaterialx(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfixed param);
inline void gl_wrap_context_glMaterialxv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfixed* param);
inline void gl_wrap_context_glMatrixMode(gl_wrap_context* ctx, GLenum mode);
inline void gl_wrap_context_glMultMatrixx(gl_wrap_context* ctx, GLfixed* m);
inline void gl_wrap_context_glMultiTexCoord4x(gl_wrap_context* ctx, GLenum texture, GLfixed s, GLfixed t, GLfixed r, GLfixed q);
inline void gl_wrap_context_glNormal3x(gl_wrap_context* ctx, GLfixed nx, GLfixed ny, GLfixed nz);
inline void gl_wrap_context_glNormalPointer(gl_wrap_context* ctx, GLenum type, GLsizei stride, void* pointer);
inline void gl_wrap_context_glOrthox(gl_wrap_context* ctx, GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f);
inline void gl_wrap_context_glPixelStorei(gl_wrap_context* ctx, GLenum pname, GLint param);
inline void gl_wrap_context_glPointParameterx(gl_wrap_context* ctx, GLenum pname, GLfixed param);
inline void gl_wrap_context_glPointParameterxv(gl_wrap_context* ctx, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glPointSizex(gl_wrap_context* ctx, GLfixed size);
inline void gl_wrap_context_glPolygonOffsetx(gl_wrap_context* ctx, GLfixed factor, GLfixed units);
inline void gl_wrap_context_glPopMatrix(gl_wrap_context* ctxvoid);
inline void gl_wrap_context_glPushMatrix(gl_wrap_context* ctxvoid);
inline void gl_wrap_context_glReadPixels(gl_wrap_context* ctx, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, void* pixels);
inline void gl_wrap_context_glRotatex(gl_wrap_context* ctx, GLfixed angle, GLfixed x, GLfixed y, GLfixed z);
inline void gl_wrap_context_glSampleCoverage(gl_wrap_context* ctx, GLfloat value, GLboolean invert);
inline void gl_wrap_context_glSampleCoveragex(gl_wrap_context* ctx, GLclampx value, GLboolean invert);
inline void gl_wrap_context_glScalex(gl_wrap_context* ctx, GLfixed x, GLfixed y, GLfixed z);
inline void gl_wrap_context_glScissor(gl_wrap_context* ctx, GLint x, GLint y, GLsizei width, GLsizei height);
inline void gl_wrap_context_glShadeModel(gl_wrap_context* ctx, GLenum mode);
inline void gl_wrap_context_glStencilFunc(gl_wrap_context* ctx, GLenum func, GLint ref, GLuint mask);
inline void gl_wrap_context_glStencilMask(gl_wrap_context* ctx, GLuint mask);
inline void gl_wrap_context_glStencilOp(gl_wrap_context* ctx, GLenum fail, GLenum zfail, GLenum zpass);
inline void gl_wrap_context_glTexCoordPointer(gl_wrap_context* ctx, GLint size, GLenum type, GLsizei stride, void* pointer);
inline void gl_wrap_context_glTexEnvi(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint param);
inline void gl_wrap_context_glTexEnvx(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed param);
inline void gl_wrap_context_glTexEnviv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params);
inline void gl_wrap_context_glTexEnvxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glTexImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, void* pixels);
inline void gl_wrap_context_glTexParameteri(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint param);
inline void gl_wrap_context_glTexParameterx(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed param);
inline void gl_wrap_context_glTexParameteriv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params);
inline void gl_wrap_context_glTexParameterxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params);
inline void gl_wrap_context_glTexSubImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, void* pixels);
inline void gl_wrap_context_glTranslatex(gl_wrap_context* ctx, GLfixed x, GLfixed y, GLfixed z);
inline void gl_wrap_context_glVertexPointer(gl_wrap_context* ctx, GLint size, GLenum type, GLsizei stride, void* pointer);
inline void gl_wrap_context_glViewport(gl_wrap_context* ctx, GLint x, GLint y, GLsizei width, GLsizei height);


// Argument struct definitions for batch function handlers.
typedef struct {
	GLenum func;
	GLfloat ref;
} gl_wrap_handler_glAlphaFunc_args;

typedef struct {
	GLfloat red;
	GLfloat green;
	GLfloat blue;
	GLfloat alpha;
} gl_wrap_handler_glClearColor_args;

typedef struct {
	GLfloat d;
} gl_wrap_handler_glClearDepthf_args;

typedef struct {
	GLenum p;
	GLfloat* eqn;
} gl_wrap_handler_glClipPlanef_args;

typedef struct {
	GLfloat red;
	GLfloat green;
	GLfloat blue;
	GLfloat alpha;
} gl_wrap_handler_glColor4f_args;

typedef struct {
	GLfloat n;
	GLfloat f;
} gl_wrap_handler_glDepthRangef_args;

typedef struct {
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glFogf_args;

typedef struct {
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glFogfv_args;

typedef struct {
	GLfloat l;
	GLfloat r;
	GLfloat b;
	GLfloat t;
	GLfloat n;
	GLfloat f;
} gl_wrap_handler_glFrustumf_args;

typedef struct {
	GLenum plane;
	GLfloat* equation;
} gl_wrap_handler_glGetClipPlanef_args;

typedef struct {
	GLenum pname;
	GLfloat* data;
} gl_wrap_handler_glGetFloatv_args;

typedef struct {
	GLenum light;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glGetLightfv_args;

typedef struct {
	GLenum face;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glGetMaterialfv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glGetTexEnvfv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glGetTexParameterfv_args;

typedef struct {
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glLightModelf_args;

typedef struct {
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glLightModelfv_args;

typedef struct {
	GLenum light;
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glLightf_args;

typedef struct {
	GLenum light;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glLightfv_args;

typedef struct {
	GLfloat width;
} gl_wrap_handler_glLineWidth_args;

typedef struct {
	GLfloat* m;
} gl_wrap_handler_glLoadMatrixf_args;

typedef struct {
	GLenum face;
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glMaterialf_args;

typedef struct {
	GLenum face;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glMaterialfv_args;

typedef struct {
	GLfloat* m;
} gl_wrap_handler_glMultMatrixf_args;

typedef struct {
	GLenum target;
	GLfloat s;
	GLfloat t;
	GLfloat r;
	GLfloat q;
} gl_wrap_handler_glMultiTexCoord4f_args;

typedef struct {
	GLfloat nx;
	GLfloat ny;
	GLfloat nz;
} gl_wrap_handler_glNormal3f_args;

typedef struct {
	GLfloat l;
	GLfloat r;
	GLfloat b;
	GLfloat t;
	GLfloat n;
	GLfloat f;
} gl_wrap_handler_glOrthof_args;

typedef struct {
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glPointParameterf_args;

typedef struct {
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glPointParameterfv_args;

typedef struct {
	GLfloat size;
} gl_wrap_handler_glPointSize_args;

typedef struct {
	GLfloat factor;
	GLfloat units;
} gl_wrap_handler_glPolygonOffset_args;

typedef struct {
	GLfloat angle;
	GLfloat x;
	GLfloat y;
	GLfloat z;
} gl_wrap_handler_glRotatef_args;

typedef struct {
	GLfloat x;
	GLfloat y;
	GLfloat z;
} gl_wrap_handler_glScalef_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glTexEnvf_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glTexEnvfv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfloat param;
} gl_wrap_handler_glTexParameterf_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfloat* params;
} gl_wrap_handler_glTexParameterfv_args;

typedef struct {
	GLfloat x;
	GLfloat y;
	GLfloat z;
} gl_wrap_handler_glTranslatef_args;

typedef struct {
	GLenum texture;
} gl_wrap_handler_glActiveTexture_args;

typedef struct {
	GLenum func;
	GLfixed ref;
} gl_wrap_handler_glAlphaFuncx_args;

typedef struct {
	GLenum target;
	GLuint buffer;
} gl_wrap_handler_glBindBuffer_args;

typedef struct {
	GLenum target;
	GLuint texture;
} gl_wrap_handler_glBindTexture_args;

typedef struct {
	GLenum sfactor;
	GLenum dfactor;
} gl_wrap_handler_glBlendFunc_args;

typedef struct {
	GLenum target;
	GLsizeiptr size;
	void* data;
	GLenum usage;
} gl_wrap_handler_glBufferData_args;

typedef struct {
	GLenum target;
	GLintptr offset;
	GLsizeiptr size;
	void* data;
} gl_wrap_handler_glBufferSubData_args;

typedef struct {
	GLbitfield mask;
} gl_wrap_handler_glClear_args;

typedef struct {
	GLfixed red;
	GLfixed green;
	GLfixed blue;
	GLfixed alpha;
} gl_wrap_handler_glClearColorx_args;

typedef struct {
	GLfixed depth;
} gl_wrap_handler_glClearDepthx_args;

typedef struct {
	GLint s;
} gl_wrap_handler_glClearStencil_args;

typedef struct {
	GLenum texture;
} gl_wrap_handler_glClientActiveTexture_args;

typedef struct {
	GLenum plane;
	GLfixed* equation;
} gl_wrap_handler_glClipPlanex_args;

typedef struct {
	GLubyte red;
	GLubyte green;
	GLubyte blue;
	GLubyte alpha;
} gl_wrap_handler_glColor4ub_args;

typedef struct {
	GLfixed red;
	GLfixed green;
	GLfixed blue;
	GLfixed alpha;
} gl_wrap_handler_glColor4x_args;

typedef struct {
	GLboolean red;
	GLboolean green;
	GLboolean blue;
	GLboolean alpha;
} gl_wrap_handler_glColorMask_args;

typedef struct {
	GLint size;
	GLenum type;
	GLsizei stride;
	void* pointer;
} gl_wrap_handler_glColorPointer_args;

typedef struct {
	GLenum target;
	GLint level;
	GLenum internalformat;
	GLsizei width;
	GLsizei height;
	GLint border;
	GLsizei imageSize;
	void* data;
} gl_wrap_handler_glCompressedTexImage2D_args;

typedef struct {
	GLenum target;
	GLint level;
	GLint xoffset;
	GLint yoffset;
	GLsizei width;
	GLsizei height;
	GLenum format;
	GLsizei imageSize;
	void* data;
} gl_wrap_handler_glCompressedTexSubImage2D_args;

typedef struct {
	GLenum target;
	GLint level;
	GLenum internalformat;
	GLint x;
	GLint y;
	GLsizei width;
	GLsizei height;
	GLint border;
} gl_wrap_handler_glCopyTexImage2D_args;

typedef struct {
	GLenum target;
	GLint level;
	GLint xoffset;
	GLint yoffset;
	GLint x;
	GLint y;
	GLsizei width;
	GLsizei height;
} gl_wrap_handler_glCopyTexSubImage2D_args;

typedef struct {
	GLenum mode;
} gl_wrap_handler_glCullFace_args;

typedef struct {
	GLsizei n;
	GLuint* buffers;
} gl_wrap_handler_glDeleteBuffers_args;

typedef struct {
	GLsizei n;
	GLuint* textures;
} gl_wrap_handler_glDeleteTextures_args;

typedef struct {
	GLenum func;
} gl_wrap_handler_glDepthFunc_args;

typedef struct {
	GLboolean flag;
} gl_wrap_handler_glDepthMask_args;

typedef struct {
	GLfixed n;
	GLfixed f;
} gl_wrap_handler_glDepthRangex_args;

typedef struct {
	GLenum cap;
} gl_wrap_handler_glDisable_args;

typedef struct {
	GLenum array;
} gl_wrap_handler_glDisableClientState_args;

typedef struct {
	GLenum mode;
	GLint first;
	GLsizei count;
} gl_wrap_handler_glDrawArrays_args;

typedef struct {
	GLenum mode;
	GLsizei count;
	GLenum type;
	void* indices;
} gl_wrap_handler_glDrawElements_args;

typedef struct {
	GLenum cap;
} gl_wrap_handler_glEnable_args;

typedef struct {
	GLenum array;
} gl_wrap_handler_glEnableClientState_args;

typedef struct {
} gl_wrap_handler_glFinish_args;

typedef struct {
} gl_wrap_handler_glFlush_args;

typedef struct {
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glFogx_args;

typedef struct {
	GLenum pname;
	GLfixed* param;
} gl_wrap_handler_glFogxv_args;

typedef struct {
	GLenum mode;
} gl_wrap_handler_glFrontFace_args;

typedef struct {
	GLfixed l;
	GLfixed r;
	GLfixed b;
	GLfixed t;
	GLfixed n;
	GLfixed f;
} gl_wrap_handler_glFrustumx_args;

typedef struct {
	GLenum pname;
	GLboolean* data;
} gl_wrap_handler_glGetBooleanv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint* params;
} gl_wrap_handler_glGetBufferParameteriv_args;

typedef struct {
	GLenum plane;
	GLfixed* equation;
} gl_wrap_handler_glGetClipPlanex_args;

typedef struct {
	GLsizei n;
	GLuint* buffers;
} gl_wrap_handler_glGenBuffers_args;

typedef struct {
	GLsizei n;
	GLuint* textures;
} gl_wrap_handler_glGenTextures_args;

typedef struct {
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glGetFixedv_args;

typedef struct {
	GLenum pname;
	GLint* data;
} gl_wrap_handler_glGetIntegerv_args;

typedef struct {
	GLenum light;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glGetLightxv_args;

typedef struct {
	GLenum face;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glGetMaterialxv_args;

typedef struct {
	GLenum pname;
	void** params;
} gl_wrap_handler_glGetPointerv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint* params;
} gl_wrap_handler_glGetTexEnviv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glGetTexEnvxv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint* params;
} gl_wrap_handler_glGetTexParameteriv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glGetTexParameterxv_args;

typedef struct {
	GLenum target;
	GLenum mode;
} gl_wrap_handler_glHint_args;

typedef struct {
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glLightModelx_args;

typedef struct {
	GLenum pname;
	GLfixed* param;
} gl_wrap_handler_glLightModelxv_args;

typedef struct {
	GLenum light;
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glLightx_args;

typedef struct {
	GLenum light;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glLightxv_args;

typedef struct {
	GLfixed width;
} gl_wrap_handler_glLineWidthx_args;

typedef struct {
} gl_wrap_handler_glLoadIdentity_args;

typedef struct {
	GLfixed* m;
} gl_wrap_handler_glLoadMatrixx_args;

typedef struct {
	GLenum opcode;
} gl_wrap_handler_glLogicOp_args;

typedef struct {
	GLenum face;
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glMaterialx_args;

typedef struct {
	GLenum face;
	GLenum pname;
	GLfixed* param;
} gl_wrap_handler_glMaterialxv_args;

typedef struct {
	GLenum mode;
} gl_wrap_handler_glMatrixMode_args;

typedef struct {
	GLfixed* m;
} gl_wrap_handler_glMultMatrixx_args;

typedef struct {
	GLenum texture;
	GLfixed s;
	GLfixed t;
	GLfixed r;
	GLfixed q;
} gl_wrap_handler_glMultiTexCoord4x_args;

typedef struct {
	GLfixed nx;
	GLfixed ny;
	GLfixed nz;
} gl_wrap_handler_glNormal3x_args;

typedef struct {
	GLenum type;
	GLsizei stride;
	void* pointer;
} gl_wrap_handler_glNormalPointer_args;

typedef struct {
	GLfixed l;
	GLfixed r;
	GLfixed b;
	GLfixed t;
	GLfixed n;
	GLfixed f;
} gl_wrap_handler_glOrthox_args;

typedef struct {
	GLenum pname;
	GLint param;
} gl_wrap_handler_glPixelStorei_args;

typedef struct {
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glPointParameterx_args;

typedef struct {
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glPointParameterxv_args;

typedef struct {
	GLfixed size;
} gl_wrap_handler_glPointSizex_args;

typedef struct {
	GLfixed factor;
	GLfixed units;
} gl_wrap_handler_glPolygonOffsetx_args;

typedef struct {
} gl_wrap_handler_glPopMatrix_args;

typedef struct {
} gl_wrap_handler_glPushMatrix_args;

typedef struct {
	GLint x;
	GLint y;
	GLsizei width;
	GLsizei height;
	GLenum format;
	GLenum type;
	void* pixels;
} gl_wrap_handler_glReadPixels_args;

typedef struct {
	GLfixed angle;
	GLfixed x;
	GLfixed y;
	GLfixed z;
} gl_wrap_handler_glRotatex_args;

typedef struct {
	GLfloat value;
	GLboolean invert;
} gl_wrap_handler_glSampleCoverage_args;

typedef struct {
	GLclampx value;
	GLboolean invert;
} gl_wrap_handler_glSampleCoveragex_args;

typedef struct {
	GLfixed x;
	GLfixed y;
	GLfixed z;
} gl_wrap_handler_glScalex_args;

typedef struct {
	GLint x;
	GLint y;
	GLsizei width;
	GLsizei height;
} gl_wrap_handler_glScissor_args;

typedef struct {
	GLenum mode;
} gl_wrap_handler_glShadeModel_args;

typedef struct {
	GLenum func;
	GLint ref;
	GLuint mask;
} gl_wrap_handler_glStencilFunc_args;

typedef struct {
	GLuint mask;
} gl_wrap_handler_glStencilMask_args;

typedef struct {
	GLenum fail;
	GLenum zfail;
	GLenum zpass;
} gl_wrap_handler_glStencilOp_args;

typedef struct {
	GLint size;
	GLenum type;
	GLsizei stride;
	void* pointer;
} gl_wrap_handler_glTexCoordPointer_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint param;
} gl_wrap_handler_glTexEnvi_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glTexEnvx_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint* params;
} gl_wrap_handler_glTexEnviv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glTexEnvxv_args;

typedef struct {
	GLenum target;
	GLint level;
	GLint internalformat;
	GLsizei width;
	GLsizei height;
	GLint border;
	GLenum format;
	GLenum type;
	void* pixels;
} gl_wrap_handler_glTexImage2D_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint param;
} gl_wrap_handler_glTexParameteri_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfixed param;
} gl_wrap_handler_glTexParameterx_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLint* params;
} gl_wrap_handler_glTexParameteriv_args;

typedef struct {
	GLenum target;
	GLenum pname;
	GLfixed* params;
} gl_wrap_handler_glTexParameterxv_args;

typedef struct {
	GLenum target;
	GLint level;
	GLint xoffset;
	GLint yoffset;
	GLsizei width;
	GLsizei height;
	GLenum format;
	GLenum type;
	void* pixels;
} gl_wrap_handler_glTexSubImage2D_args;

typedef struct {
	GLfixed x;
	GLfixed y;
	GLfixed z;
} gl_wrap_handler_glTranslatex_args;

typedef struct {
	GLint size;
	GLenum type;
	GLsizei stride;
	void* pointer;
} gl_wrap_handler_glVertexPointer_args;

typedef struct {
	GLint x;
	GLint y;
	GLsizei width;
	GLsizei height;
} gl_wrap_handler_glViewport_args;



// Other custom definitions
typedef void (*gl_wrap_jump_handler)(gl_wrap_context* ctx, void* args);

typedef struct {
	int jump_index;
	void* args;
} gl_wrap_batch_func;

void gl_wrap_batch_exec(gl_wrap_context* ctx, gl_wrap_batch_func* funcs, int numFuncs);

#endif
