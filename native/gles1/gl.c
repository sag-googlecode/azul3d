// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This source file was automatically generated using glwrap.

#include "gl.h"

typedef void (*__glwrap_func_ptr)(void);

#if defined(__WIN32) || defined(__WIN64)
	typedef void* HMODULE;
	typedef int (*FARPROC)(void);
	typedef int (*PROC)(void);

	//extern PROC wglGetProcAddress(const char* name);
	extern HMODULE LoadLibraryA(const char* name);
	extern FARPROC GetProcAddress(HMODULE, const char*);

	HMODULE glwrap_OpenGL32;


	typedef PROC (*__glwrap_PFNWGLGETPROCADDRESS)(const char*);
	__glwrap_PFNWGLGETPROCADDRESS __glwrap_wglGetProcAddressPtr;
	inline PROC __glwrap_wglGetProcAddress(const char* name) {
		if(__glwrap_wglGetProcAddressPtr == NULL) {
			if(glwrap_OpenGL32 == NULL) {
				glwrap_OpenGL32 = LoadLibraryA("opengl32.dll");
			}
			__glwrap_wglGetProcAddressPtr = (__glwrap_PFNWGLGETPROCADDRESS)GetProcAddress(glwrap_OpenGL32, "wglGetProcAddress");
		}
		return __glwrap_wglGetProcAddressPtr(name);
	}

#elif defined(__linux) || defined(__unix) || defined(__posix)
	// See http://dri.freedesktop.org/wiki/glXGetProcAddressNeverReturnsNULL
	//
	// glXGetProcAddressARB is *required* to be statically defined in libGL,
	// but glXGetProcAddress is not, and will fail to be found in nvidia's
	// libGL
	extern __glwrap_func_ptr glXGetProcAddressARB(const GLubyte * procName);

#elif defined(__APPLE__)
	// Todo: OS X support.
#endif

inline __glwrap_func_ptr gl_wrap_get_pointer(const char* name) {
	#if defined(__WIN32) || defined(__WIN64)
		void* ptr = __glwrap_wglGetProcAddress(name);
		intptr_t iptr = (intptr_t)ptr;

		if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
			// Could be a core function, then.

			// No need for this: because it's done in wglGetProcAddress()
			//if(glwrap_OpenGL32 == NULL) {
			//	glwrap_OpenGL32 = LoadLibraryA("opengl32.dll");
			//}
			return (__glwrap_func_ptr)GetProcAddress(glwrap_OpenGL32, name);
		}

		return ptr;

	#elif defined(__linux) || defined(__unix) || defined(__posix)
		return glXGetProcAddressARB(name);

	#elif defined(__APPLE__)
		// Todo: OS X support.
	#endif
}

// Function definition for each appropriate OpenGL function.
//
// If the pointer in the context for the function is null; it is loaded
// immedietly (as such this is effectively lazy-loading).
void gl_wrap_context_glAlphaFunc(gl_wrap_context* ctx, GLenum func, GLfloat ref) {
	if(ctx->glAlphaFuncProc == NULL) {
		ctx->glAlphaFuncProc = (PFNGLALPHAFUNCPROC)gl_wrap_get_pointer("glAlphaFunc");
	}
	ctx->glAlphaFuncProc(func, ref);
};
void gl_wrap_context_glClearColor(gl_wrap_context* ctx, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
	if(ctx->glClearColorProc == NULL) {
		ctx->glClearColorProc = (PFNGLCLEARCOLORPROC)gl_wrap_get_pointer("glClearColor");
	}
	ctx->glClearColorProc(red, green, blue, alpha);
};
void gl_wrap_context_glClearDepthf(gl_wrap_context* ctx, GLfloat d) {
	if(ctx->glClearDepthfProc == NULL) {
		ctx->glClearDepthfProc = (PFNGLCLEARDEPTHFPROC)gl_wrap_get_pointer("glClearDepthf");
	}
	ctx->glClearDepthfProc(d);
};
void gl_wrap_context_glClipPlanef(gl_wrap_context* ctx, GLenum p, GLfloat* eqn) {
	if(ctx->glClipPlanefProc == NULL) {
		ctx->glClipPlanefProc = (PFNGLCLIPPLANEFPROC)gl_wrap_get_pointer("glClipPlanef");
	}
	ctx->glClipPlanefProc(p, eqn);
};
void gl_wrap_context_glColor4f(gl_wrap_context* ctx, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
	if(ctx->glColor4fProc == NULL) {
		ctx->glColor4fProc = (PFNGLCOLOR4FPROC)gl_wrap_get_pointer("glColor4f");
	}
	ctx->glColor4fProc(red, green, blue, alpha);
};
void gl_wrap_context_glDepthRangef(gl_wrap_context* ctx, GLfloat n, GLfloat f) {
	if(ctx->glDepthRangefProc == NULL) {
		ctx->glDepthRangefProc = (PFNGLDEPTHRANGEFPROC)gl_wrap_get_pointer("glDepthRangef");
	}
	ctx->glDepthRangefProc(n, f);
};
void gl_wrap_context_glFogf(gl_wrap_context* ctx, GLenum pname, GLfloat param) {
	if(ctx->glFogfProc == NULL) {
		ctx->glFogfProc = (PFNGLFOGFPROC)gl_wrap_get_pointer("glFogf");
	}
	ctx->glFogfProc(pname, param);
};
void gl_wrap_context_glFogfv(gl_wrap_context* ctx, GLenum pname, GLfloat* params) {
	if(ctx->glFogfvProc == NULL) {
		ctx->glFogfvProc = (PFNGLFOGFVPROC)gl_wrap_get_pointer("glFogfv");
	}
	ctx->glFogfvProc(pname, params);
};
void gl_wrap_context_glFrustumf(gl_wrap_context* ctx, GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f) {
	if(ctx->glFrustumfProc == NULL) {
		ctx->glFrustumfProc = (PFNGLFRUSTUMFPROC)gl_wrap_get_pointer("glFrustumf");
	}
	ctx->glFrustumfProc(l, r, b, t, n, f);
};
void gl_wrap_context_glGetClipPlanef(gl_wrap_context* ctx, GLenum plane, GLfloat* equation) {
	if(ctx->glGetClipPlanefProc == NULL) {
		ctx->glGetClipPlanefProc = (PFNGLGETCLIPPLANEFPROC)gl_wrap_get_pointer("glGetClipPlanef");
	}
	ctx->glGetClipPlanefProc(plane, equation);
};
void gl_wrap_context_glGetFloatv(gl_wrap_context* ctx, GLenum pname, GLfloat* data) {
	if(ctx->glGetFloatvProc == NULL) {
		ctx->glGetFloatvProc = (PFNGLGETFLOATVPROC)gl_wrap_get_pointer("glGetFloatv");
	}
	ctx->glGetFloatvProc(pname, data);
};
void gl_wrap_context_glGetLightfv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfloat* params) {
	if(ctx->glGetLightfvProc == NULL) {
		ctx->glGetLightfvProc = (PFNGLGETLIGHTFVPROC)gl_wrap_get_pointer("glGetLightfv");
	}
	ctx->glGetLightfvProc(light, pname, params);
};
void gl_wrap_context_glGetMaterialfv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfloat* params) {
	if(ctx->glGetMaterialfvProc == NULL) {
		ctx->glGetMaterialfvProc = (PFNGLGETMATERIALFVPROC)gl_wrap_get_pointer("glGetMaterialfv");
	}
	ctx->glGetMaterialfvProc(face, pname, params);
};
void gl_wrap_context_glGetTexEnvfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params) {
	if(ctx->glGetTexEnvfvProc == NULL) {
		ctx->glGetTexEnvfvProc = (PFNGLGETTEXENVFVPROC)gl_wrap_get_pointer("glGetTexEnvfv");
	}
	ctx->glGetTexEnvfvProc(target, pname, params);
};
void gl_wrap_context_glGetTexParameterfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params) {
	if(ctx->glGetTexParameterfvProc == NULL) {
		ctx->glGetTexParameterfvProc = (PFNGLGETTEXPARAMETERFVPROC)gl_wrap_get_pointer("glGetTexParameterfv");
	}
	ctx->glGetTexParameterfvProc(target, pname, params);
};
void gl_wrap_context_glLightModelf(gl_wrap_context* ctx, GLenum pname, GLfloat param) {
	if(ctx->glLightModelfProc == NULL) {
		ctx->glLightModelfProc = (PFNGLLIGHTMODELFPROC)gl_wrap_get_pointer("glLightModelf");
	}
	ctx->glLightModelfProc(pname, param);
};
void gl_wrap_context_glLightModelfv(gl_wrap_context* ctx, GLenum pname, GLfloat* params) {
	if(ctx->glLightModelfvProc == NULL) {
		ctx->glLightModelfvProc = (PFNGLLIGHTMODELFVPROC)gl_wrap_get_pointer("glLightModelfv");
	}
	ctx->glLightModelfvProc(pname, params);
};
void gl_wrap_context_glLightf(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfloat param) {
	if(ctx->glLightfProc == NULL) {
		ctx->glLightfProc = (PFNGLLIGHTFPROC)gl_wrap_get_pointer("glLightf");
	}
	ctx->glLightfProc(light, pname, param);
};
void gl_wrap_context_glLightfv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfloat* params) {
	if(ctx->glLightfvProc == NULL) {
		ctx->glLightfvProc = (PFNGLLIGHTFVPROC)gl_wrap_get_pointer("glLightfv");
	}
	ctx->glLightfvProc(light, pname, params);
};
void gl_wrap_context_glLineWidth(gl_wrap_context* ctx, GLfloat width) {
	if(ctx->glLineWidthProc == NULL) {
		ctx->glLineWidthProc = (PFNGLLINEWIDTHPROC)gl_wrap_get_pointer("glLineWidth");
	}
	ctx->glLineWidthProc(width);
};
void gl_wrap_context_glLoadMatrixf(gl_wrap_context* ctx, GLfloat* m) {
	if(ctx->glLoadMatrixfProc == NULL) {
		ctx->glLoadMatrixfProc = (PFNGLLOADMATRIXFPROC)gl_wrap_get_pointer("glLoadMatrixf");
	}
	ctx->glLoadMatrixfProc(m);
};
void gl_wrap_context_glMaterialf(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfloat param) {
	if(ctx->glMaterialfProc == NULL) {
		ctx->glMaterialfProc = (PFNGLMATERIALFPROC)gl_wrap_get_pointer("glMaterialf");
	}
	ctx->glMaterialfProc(face, pname, param);
};
void gl_wrap_context_glMaterialfv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfloat* params) {
	if(ctx->glMaterialfvProc == NULL) {
		ctx->glMaterialfvProc = (PFNGLMATERIALFVPROC)gl_wrap_get_pointer("glMaterialfv");
	}
	ctx->glMaterialfvProc(face, pname, params);
};
void gl_wrap_context_glMultMatrixf(gl_wrap_context* ctx, GLfloat* m) {
	if(ctx->glMultMatrixfProc == NULL) {
		ctx->glMultMatrixfProc = (PFNGLMULTMATRIXFPROC)gl_wrap_get_pointer("glMultMatrixf");
	}
	ctx->glMultMatrixfProc(m);
};
void gl_wrap_context_glMultiTexCoord4f(gl_wrap_context* ctx, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
	if(ctx->glMultiTexCoord4fProc == NULL) {
		ctx->glMultiTexCoord4fProc = (PFNGLMULTITEXCOORD4FPROC)gl_wrap_get_pointer("glMultiTexCoord4f");
	}
	ctx->glMultiTexCoord4fProc(target, s, t, r, q);
};
void gl_wrap_context_glNormal3f(gl_wrap_context* ctx, GLfloat nx, GLfloat ny, GLfloat nz) {
	if(ctx->glNormal3fProc == NULL) {
		ctx->glNormal3fProc = (PFNGLNORMAL3FPROC)gl_wrap_get_pointer("glNormal3f");
	}
	ctx->glNormal3fProc(nx, ny, nz);
};
void gl_wrap_context_glOrthof(gl_wrap_context* ctx, GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f) {
	if(ctx->glOrthofProc == NULL) {
		ctx->glOrthofProc = (PFNGLORTHOFPROC)gl_wrap_get_pointer("glOrthof");
	}
	ctx->glOrthofProc(l, r, b, t, n, f);
};
void gl_wrap_context_glPointParameterf(gl_wrap_context* ctx, GLenum pname, GLfloat param) {
	if(ctx->glPointParameterfProc == NULL) {
		ctx->glPointParameterfProc = (PFNGLPOINTPARAMETERFPROC)gl_wrap_get_pointer("glPointParameterf");
	}
	ctx->glPointParameterfProc(pname, param);
};
void gl_wrap_context_glPointParameterfv(gl_wrap_context* ctx, GLenum pname, GLfloat* params) {
	if(ctx->glPointParameterfvProc == NULL) {
		ctx->glPointParameterfvProc = (PFNGLPOINTPARAMETERFVPROC)gl_wrap_get_pointer("glPointParameterfv");
	}
	ctx->glPointParameterfvProc(pname, params);
};
void gl_wrap_context_glPointSize(gl_wrap_context* ctx, GLfloat size) {
	if(ctx->glPointSizeProc == NULL) {
		ctx->glPointSizeProc = (PFNGLPOINTSIZEPROC)gl_wrap_get_pointer("glPointSize");
	}
	ctx->glPointSizeProc(size);
};
void gl_wrap_context_glPolygonOffset(gl_wrap_context* ctx, GLfloat factor, GLfloat units) {
	if(ctx->glPolygonOffsetProc == NULL) {
		ctx->glPolygonOffsetProc = (PFNGLPOLYGONOFFSETPROC)gl_wrap_get_pointer("glPolygonOffset");
	}
	ctx->glPolygonOffsetProc(factor, units);
};
void gl_wrap_context_glRotatef(gl_wrap_context* ctx, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
	if(ctx->glRotatefProc == NULL) {
		ctx->glRotatefProc = (PFNGLROTATEFPROC)gl_wrap_get_pointer("glRotatef");
	}
	ctx->glRotatefProc(angle, x, y, z);
};
void gl_wrap_context_glScalef(gl_wrap_context* ctx, GLfloat x, GLfloat y, GLfloat z) {
	if(ctx->glScalefProc == NULL) {
		ctx->glScalefProc = (PFNGLSCALEFPROC)gl_wrap_get_pointer("glScalef");
	}
	ctx->glScalefProc(x, y, z);
};
void gl_wrap_context_glTexEnvf(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat param) {
	if(ctx->glTexEnvfProc == NULL) {
		ctx->glTexEnvfProc = (PFNGLTEXENVFPROC)gl_wrap_get_pointer("glTexEnvf");
	}
	ctx->glTexEnvfProc(target, pname, param);
};
void gl_wrap_context_glTexEnvfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params) {
	if(ctx->glTexEnvfvProc == NULL) {
		ctx->glTexEnvfvProc = (PFNGLTEXENVFVPROC)gl_wrap_get_pointer("glTexEnvfv");
	}
	ctx->glTexEnvfvProc(target, pname, params);
};
void gl_wrap_context_glTexParameterf(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat param) {
	if(ctx->glTexParameterfProc == NULL) {
		ctx->glTexParameterfProc = (PFNGLTEXPARAMETERFPROC)gl_wrap_get_pointer("glTexParameterf");
	}
	ctx->glTexParameterfProc(target, pname, param);
};
void gl_wrap_context_glTexParameterfv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfloat* params) {
	if(ctx->glTexParameterfvProc == NULL) {
		ctx->glTexParameterfvProc = (PFNGLTEXPARAMETERFVPROC)gl_wrap_get_pointer("glTexParameterfv");
	}
	ctx->glTexParameterfvProc(target, pname, params);
};
void gl_wrap_context_glTranslatef(gl_wrap_context* ctx, GLfloat x, GLfloat y, GLfloat z) {
	if(ctx->glTranslatefProc == NULL) {
		ctx->glTranslatefProc = (PFNGLTRANSLATEFPROC)gl_wrap_get_pointer("glTranslatef");
	}
	ctx->glTranslatefProc(x, y, z);
};
void gl_wrap_context_glActiveTexture(gl_wrap_context* ctx, GLenum texture) {
	if(ctx->glActiveTextureProc == NULL) {
		ctx->glActiveTextureProc = (PFNGLACTIVETEXTUREPROC)gl_wrap_get_pointer("glActiveTexture");
	}
	ctx->glActiveTextureProc(texture);
};
void gl_wrap_context_glAlphaFuncx(gl_wrap_context* ctx, GLenum func, GLfixed ref) {
	if(ctx->glAlphaFuncxProc == NULL) {
		ctx->glAlphaFuncxProc = (PFNGLALPHAFUNCXPROC)gl_wrap_get_pointer("glAlphaFuncx");
	}
	ctx->glAlphaFuncxProc(func, ref);
};
void gl_wrap_context_glBindBuffer(gl_wrap_context* ctx, GLenum target, GLuint buffer) {
	if(ctx->glBindBufferProc == NULL) {
		ctx->glBindBufferProc = (PFNGLBINDBUFFERPROC)gl_wrap_get_pointer("glBindBuffer");
	}
	ctx->glBindBufferProc(target, buffer);
};
void gl_wrap_context_glBindTexture(gl_wrap_context* ctx, GLenum target, GLuint texture) {
	if(ctx->glBindTextureProc == NULL) {
		ctx->glBindTextureProc = (PFNGLBINDTEXTUREPROC)gl_wrap_get_pointer("glBindTexture");
	}
	ctx->glBindTextureProc(target, texture);
};
void gl_wrap_context_glBlendFunc(gl_wrap_context* ctx, GLenum sfactor, GLenum dfactor) {
	if(ctx->glBlendFuncProc == NULL) {
		ctx->glBlendFuncProc = (PFNGLBLENDFUNCPROC)gl_wrap_get_pointer("glBlendFunc");
	}
	ctx->glBlendFuncProc(sfactor, dfactor);
};
void gl_wrap_context_glBufferData(gl_wrap_context* ctx, GLenum target, GLsizeiptr size, void* data, GLenum usage) {
	if(ctx->glBufferDataProc == NULL) {
		ctx->glBufferDataProc = (PFNGLBUFFERDATAPROC)gl_wrap_get_pointer("glBufferData");
	}
	ctx->glBufferDataProc(target, size, data, usage);
};
void gl_wrap_context_glBufferSubData(gl_wrap_context* ctx, GLenum target, GLintptr offset, GLsizeiptr size, void* data) {
	if(ctx->glBufferSubDataProc == NULL) {
		ctx->glBufferSubDataProc = (PFNGLBUFFERSUBDATAPROC)gl_wrap_get_pointer("glBufferSubData");
	}
	ctx->glBufferSubDataProc(target, offset, size, data);
};
void gl_wrap_context_glClear(gl_wrap_context* ctx, GLbitfield mask) {
	if(ctx->glClearProc == NULL) {
		ctx->glClearProc = (PFNGLCLEARPROC)gl_wrap_get_pointer("glClear");
	}
	ctx->glClearProc(mask);
};
void gl_wrap_context_glClearColorx(gl_wrap_context* ctx, GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha) {
	if(ctx->glClearColorxProc == NULL) {
		ctx->glClearColorxProc = (PFNGLCLEARCOLORXPROC)gl_wrap_get_pointer("glClearColorx");
	}
	ctx->glClearColorxProc(red, green, blue, alpha);
};
void gl_wrap_context_glClearDepthx(gl_wrap_context* ctx, GLfixed depth) {
	if(ctx->glClearDepthxProc == NULL) {
		ctx->glClearDepthxProc = (PFNGLCLEARDEPTHXPROC)gl_wrap_get_pointer("glClearDepthx");
	}
	ctx->glClearDepthxProc(depth);
};
void gl_wrap_context_glClearStencil(gl_wrap_context* ctx, GLint s) {
	if(ctx->glClearStencilProc == NULL) {
		ctx->glClearStencilProc = (PFNGLCLEARSTENCILPROC)gl_wrap_get_pointer("glClearStencil");
	}
	ctx->glClearStencilProc(s);
};
void gl_wrap_context_glClientActiveTexture(gl_wrap_context* ctx, GLenum texture) {
	if(ctx->glClientActiveTextureProc == NULL) {
		ctx->glClientActiveTextureProc = (PFNGLCLIENTACTIVETEXTUREPROC)gl_wrap_get_pointer("glClientActiveTexture");
	}
	ctx->glClientActiveTextureProc(texture);
};
void gl_wrap_context_glClipPlanex(gl_wrap_context* ctx, GLenum plane, GLfixed* equation) {
	if(ctx->glClipPlanexProc == NULL) {
		ctx->glClipPlanexProc = (PFNGLCLIPPLANEXPROC)gl_wrap_get_pointer("glClipPlanex");
	}
	ctx->glClipPlanexProc(plane, equation);
};
void gl_wrap_context_glColor4ub(gl_wrap_context* ctx, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
	if(ctx->glColor4ubProc == NULL) {
		ctx->glColor4ubProc = (PFNGLCOLOR4UBPROC)gl_wrap_get_pointer("glColor4ub");
	}
	ctx->glColor4ubProc(red, green, blue, alpha);
};
void gl_wrap_context_glColor4x(gl_wrap_context* ctx, GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha) {
	if(ctx->glColor4xProc == NULL) {
		ctx->glColor4xProc = (PFNGLCOLOR4XPROC)gl_wrap_get_pointer("glColor4x");
	}
	ctx->glColor4xProc(red, green, blue, alpha);
};
void gl_wrap_context_glColorMask(gl_wrap_context* ctx, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
	if(ctx->glColorMaskProc == NULL) {
		ctx->glColorMaskProc = (PFNGLCOLORMASKPROC)gl_wrap_get_pointer("glColorMask");
	}
	ctx->glColorMaskProc(red, green, blue, alpha);
};
void gl_wrap_context_glColorPointer(gl_wrap_context* ctx, GLint size, GLenum type, GLsizei stride, void* pointer) {
	if(ctx->glColorPointerProc == NULL) {
		ctx->glColorPointerProc = (PFNGLCOLORPOINTERPROC)gl_wrap_get_pointer("glColorPointer");
	}
	ctx->glColorPointerProc(size, type, stride, pointer);
};
void gl_wrap_context_glCompressedTexImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, void* data) {
	if(ctx->glCompressedTexImage2DProc == NULL) {
		ctx->glCompressedTexImage2DProc = (PFNGLCOMPRESSEDTEXIMAGE2DPROC)gl_wrap_get_pointer("glCompressedTexImage2D");
	}
	ctx->glCompressedTexImage2DProc(target, level, internalformat, width, height, border, imageSize, data);
};
void gl_wrap_context_glCompressedTexSubImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, void* data) {
	if(ctx->glCompressedTexSubImage2DProc == NULL) {
		ctx->glCompressedTexSubImage2DProc = (PFNGLCOMPRESSEDTEXSUBIMAGE2DPROC)gl_wrap_get_pointer("glCompressedTexSubImage2D");
	}
	ctx->glCompressedTexSubImage2DProc(target, level, xoffset, yoffset, width, height, format, imageSize, data);
};
void gl_wrap_context_glCopyTexImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
	if(ctx->glCopyTexImage2DProc == NULL) {
		ctx->glCopyTexImage2DProc = (PFNGLCOPYTEXIMAGE2DPROC)gl_wrap_get_pointer("glCopyTexImage2D");
	}
	ctx->glCopyTexImage2DProc(target, level, internalformat, x, y, width, height, border);
};
void gl_wrap_context_glCopyTexSubImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
	if(ctx->glCopyTexSubImage2DProc == NULL) {
		ctx->glCopyTexSubImage2DProc = (PFNGLCOPYTEXSUBIMAGE2DPROC)gl_wrap_get_pointer("glCopyTexSubImage2D");
	}
	ctx->glCopyTexSubImage2DProc(target, level, xoffset, yoffset, x, y, width, height);
};
void gl_wrap_context_glCullFace(gl_wrap_context* ctx, GLenum mode) {
	if(ctx->glCullFaceProc == NULL) {
		ctx->glCullFaceProc = (PFNGLCULLFACEPROC)gl_wrap_get_pointer("glCullFace");
	}
	ctx->glCullFaceProc(mode);
};
void gl_wrap_context_glDeleteBuffers(gl_wrap_context* ctx, GLsizei n, GLuint* buffers) {
	if(ctx->glDeleteBuffersProc == NULL) {
		ctx->glDeleteBuffersProc = (PFNGLDELETEBUFFERSPROC)gl_wrap_get_pointer("glDeleteBuffers");
	}
	ctx->glDeleteBuffersProc(n, buffers);
};
void gl_wrap_context_glDeleteTextures(gl_wrap_context* ctx, GLsizei n, GLuint* textures) {
	if(ctx->glDeleteTexturesProc == NULL) {
		ctx->glDeleteTexturesProc = (PFNGLDELETETEXTURESPROC)gl_wrap_get_pointer("glDeleteTextures");
	}
	ctx->glDeleteTexturesProc(n, textures);
};
void gl_wrap_context_glDepthFunc(gl_wrap_context* ctx, GLenum func) {
	if(ctx->glDepthFuncProc == NULL) {
		ctx->glDepthFuncProc = (PFNGLDEPTHFUNCPROC)gl_wrap_get_pointer("glDepthFunc");
	}
	ctx->glDepthFuncProc(func);
};
void gl_wrap_context_glDepthMask(gl_wrap_context* ctx, GLboolean flag) {
	if(ctx->glDepthMaskProc == NULL) {
		ctx->glDepthMaskProc = (PFNGLDEPTHMASKPROC)gl_wrap_get_pointer("glDepthMask");
	}
	ctx->glDepthMaskProc(flag);
};
void gl_wrap_context_glDepthRangex(gl_wrap_context* ctx, GLfixed n, GLfixed f) {
	if(ctx->glDepthRangexProc == NULL) {
		ctx->glDepthRangexProc = (PFNGLDEPTHRANGEXPROC)gl_wrap_get_pointer("glDepthRangex");
	}
	ctx->glDepthRangexProc(n, f);
};
void gl_wrap_context_glDisable(gl_wrap_context* ctx, GLenum cap) {
	if(ctx->glDisableProc == NULL) {
		ctx->glDisableProc = (PFNGLDISABLEPROC)gl_wrap_get_pointer("glDisable");
	}
	ctx->glDisableProc(cap);
};
void gl_wrap_context_glDisableClientState(gl_wrap_context* ctx, GLenum array) {
	if(ctx->glDisableClientStateProc == NULL) {
		ctx->glDisableClientStateProc = (PFNGLDISABLECLIENTSTATEPROC)gl_wrap_get_pointer("glDisableClientState");
	}
	ctx->glDisableClientStateProc(array);
};
void gl_wrap_context_glDrawArrays(gl_wrap_context* ctx, GLenum mode, GLint first, GLsizei count) {
	if(ctx->glDrawArraysProc == NULL) {
		ctx->glDrawArraysProc = (PFNGLDRAWARRAYSPROC)gl_wrap_get_pointer("glDrawArrays");
	}
	ctx->glDrawArraysProc(mode, first, count);
};
void gl_wrap_context_glDrawElements(gl_wrap_context* ctx, GLenum mode, GLsizei count, GLenum type, void* indices) {
	if(ctx->glDrawElementsProc == NULL) {
		ctx->glDrawElementsProc = (PFNGLDRAWELEMENTSPROC)gl_wrap_get_pointer("glDrawElements");
	}
	ctx->glDrawElementsProc(mode, count, type, indices);
};
void gl_wrap_context_glEnable(gl_wrap_context* ctx, GLenum cap) {
	if(ctx->glEnableProc == NULL) {
		ctx->glEnableProc = (PFNGLENABLEPROC)gl_wrap_get_pointer("glEnable");
	}
	ctx->glEnableProc(cap);
};
void gl_wrap_context_glEnableClientState(gl_wrap_context* ctx, GLenum array) {
	if(ctx->glEnableClientStateProc == NULL) {
		ctx->glEnableClientStateProc = (PFNGLENABLECLIENTSTATEPROC)gl_wrap_get_pointer("glEnableClientState");
	}
	ctx->glEnableClientStateProc(array);
};
void gl_wrap_context_glFinish(gl_wrap_context* ctx) {
	if(ctx->glFinishProc == NULL) {
		ctx->glFinishProc = (PFNGLFINISHPROC)gl_wrap_get_pointer("glFinish");
	}
	ctx->glFinishProc();
};
void gl_wrap_context_glFlush(gl_wrap_context* ctx) {
	if(ctx->glFlushProc == NULL) {
		ctx->glFlushProc = (PFNGLFLUSHPROC)gl_wrap_get_pointer("glFlush");
	}
	ctx->glFlushProc();
};
void gl_wrap_context_glFogx(gl_wrap_context* ctx, GLenum pname, GLfixed param) {
	if(ctx->glFogxProc == NULL) {
		ctx->glFogxProc = (PFNGLFOGXPROC)gl_wrap_get_pointer("glFogx");
	}
	ctx->glFogxProc(pname, param);
};
void gl_wrap_context_glFogxv(gl_wrap_context* ctx, GLenum pname, GLfixed* param) {
	if(ctx->glFogxvProc == NULL) {
		ctx->glFogxvProc = (PFNGLFOGXVPROC)gl_wrap_get_pointer("glFogxv");
	}
	ctx->glFogxvProc(pname, param);
};
void gl_wrap_context_glFrontFace(gl_wrap_context* ctx, GLenum mode) {
	if(ctx->glFrontFaceProc == NULL) {
		ctx->glFrontFaceProc = (PFNGLFRONTFACEPROC)gl_wrap_get_pointer("glFrontFace");
	}
	ctx->glFrontFaceProc(mode);
};
void gl_wrap_context_glFrustumx(gl_wrap_context* ctx, GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f) {
	if(ctx->glFrustumxProc == NULL) {
		ctx->glFrustumxProc = (PFNGLFRUSTUMXPROC)gl_wrap_get_pointer("glFrustumx");
	}
	ctx->glFrustumxProc(l, r, b, t, n, f);
};
void gl_wrap_context_glGetBooleanv(gl_wrap_context* ctx, GLenum pname, GLboolean* data) {
	if(ctx->glGetBooleanvProc == NULL) {
		ctx->glGetBooleanvProc = (PFNGLGETBOOLEANVPROC)gl_wrap_get_pointer("glGetBooleanv");
	}
	ctx->glGetBooleanvProc(pname, data);
};
void gl_wrap_context_glGetBufferParameteriv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params) {
	if(ctx->glGetBufferParameterivProc == NULL) {
		ctx->glGetBufferParameterivProc = (PFNGLGETBUFFERPARAMETERIVPROC)gl_wrap_get_pointer("glGetBufferParameteriv");
	}
	ctx->glGetBufferParameterivProc(target, pname, params);
};
void gl_wrap_context_glGetClipPlanex(gl_wrap_context* ctx, GLenum plane, GLfixed* equation) {
	if(ctx->glGetClipPlanexProc == NULL) {
		ctx->glGetClipPlanexProc = (PFNGLGETCLIPPLANEXPROC)gl_wrap_get_pointer("glGetClipPlanex");
	}
	ctx->glGetClipPlanexProc(plane, equation);
};
void gl_wrap_context_glGenBuffers(gl_wrap_context* ctx, GLsizei n, GLuint* buffers) {
	if(ctx->glGenBuffersProc == NULL) {
		ctx->glGenBuffersProc = (PFNGLGENBUFFERSPROC)gl_wrap_get_pointer("glGenBuffers");
	}
	ctx->glGenBuffersProc(n, buffers);
};
void gl_wrap_context_glGenTextures(gl_wrap_context* ctx, GLsizei n, GLuint* textures) {
	if(ctx->glGenTexturesProc == NULL) {
		ctx->glGenTexturesProc = (PFNGLGENTEXTURESPROC)gl_wrap_get_pointer("glGenTextures");
	}
	ctx->glGenTexturesProc(n, textures);
};
GLenum gl_wrap_context_glGetError(gl_wrap_context* ctx) {
	if(ctx->glGetErrorProc == NULL) {
		ctx->glGetErrorProc = (PFNGLGETERRORPROC)gl_wrap_get_pointer("glGetError");
	}
	return ctx->glGetErrorProc();
};
void gl_wrap_context_glGetFixedv(gl_wrap_context* ctx, GLenum pname, GLfixed* params) {
	if(ctx->glGetFixedvProc == NULL) {
		ctx->glGetFixedvProc = (PFNGLGETFIXEDVPROC)gl_wrap_get_pointer("glGetFixedv");
	}
	ctx->glGetFixedvProc(pname, params);
};
void gl_wrap_context_glGetIntegerv(gl_wrap_context* ctx, GLenum pname, GLint* data) {
	if(ctx->glGetIntegervProc == NULL) {
		ctx->glGetIntegervProc = (PFNGLGETINTEGERVPROC)gl_wrap_get_pointer("glGetIntegerv");
	}
	ctx->glGetIntegervProc(pname, data);
};
void gl_wrap_context_glGetLightxv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfixed* params) {
	if(ctx->glGetLightxvProc == NULL) {
		ctx->glGetLightxvProc = (PFNGLGETLIGHTXVPROC)gl_wrap_get_pointer("glGetLightxv");
	}
	ctx->glGetLightxvProc(light, pname, params);
};
void gl_wrap_context_glGetMaterialxv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfixed* params) {
	if(ctx->glGetMaterialxvProc == NULL) {
		ctx->glGetMaterialxvProc = (PFNGLGETMATERIALXVPROC)gl_wrap_get_pointer("glGetMaterialxv");
	}
	ctx->glGetMaterialxvProc(face, pname, params);
};
void gl_wrap_context_glGetPointerv(gl_wrap_context* ctx, GLenum pname, void** params) {
	if(ctx->glGetPointervProc == NULL) {
		ctx->glGetPointervProc = (PFNGLGETPOINTERVPROC)gl_wrap_get_pointer("glGetPointerv");
	}
	ctx->glGetPointervProc(pname, params);
};
GLubyte* gl_wrap_context_glGetString(gl_wrap_context* ctx, GLenum name) {
	if(ctx->glGetStringProc == NULL) {
		ctx->glGetStringProc = (PFNGLGETSTRINGPROC)gl_wrap_get_pointer("glGetString");
	}
	return ctx->glGetStringProc(name);
};
void gl_wrap_context_glGetTexEnviv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params) {
	if(ctx->glGetTexEnvivProc == NULL) {
		ctx->glGetTexEnvivProc = (PFNGLGETTEXENVIVPROC)gl_wrap_get_pointer("glGetTexEnviv");
	}
	ctx->glGetTexEnvivProc(target, pname, params);
};
void gl_wrap_context_glGetTexEnvxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params) {
	if(ctx->glGetTexEnvxvProc == NULL) {
		ctx->glGetTexEnvxvProc = (PFNGLGETTEXENVXVPROC)gl_wrap_get_pointer("glGetTexEnvxv");
	}
	ctx->glGetTexEnvxvProc(target, pname, params);
};
void gl_wrap_context_glGetTexParameteriv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params) {
	if(ctx->glGetTexParameterivProc == NULL) {
		ctx->glGetTexParameterivProc = (PFNGLGETTEXPARAMETERIVPROC)gl_wrap_get_pointer("glGetTexParameteriv");
	}
	ctx->glGetTexParameterivProc(target, pname, params);
};
void gl_wrap_context_glGetTexParameterxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params) {
	if(ctx->glGetTexParameterxvProc == NULL) {
		ctx->glGetTexParameterxvProc = (PFNGLGETTEXPARAMETERXVPROC)gl_wrap_get_pointer("glGetTexParameterxv");
	}
	ctx->glGetTexParameterxvProc(target, pname, params);
};
void gl_wrap_context_glHint(gl_wrap_context* ctx, GLenum target, GLenum mode) {
	if(ctx->glHintProc == NULL) {
		ctx->glHintProc = (PFNGLHINTPROC)gl_wrap_get_pointer("glHint");
	}
	ctx->glHintProc(target, mode);
};
GLboolean gl_wrap_context_glIsBuffer(gl_wrap_context* ctx, GLuint buffer) {
	if(ctx->glIsBufferProc == NULL) {
		ctx->glIsBufferProc = (PFNGLISBUFFERPROC)gl_wrap_get_pointer("glIsBuffer");
	}
	return ctx->glIsBufferProc(buffer);
};
GLboolean gl_wrap_context_glIsEnabled(gl_wrap_context* ctx, GLenum cap) {
	if(ctx->glIsEnabledProc == NULL) {
		ctx->glIsEnabledProc = (PFNGLISENABLEDPROC)gl_wrap_get_pointer("glIsEnabled");
	}
	return ctx->glIsEnabledProc(cap);
};
GLboolean gl_wrap_context_glIsTexture(gl_wrap_context* ctx, GLuint texture) {
	if(ctx->glIsTextureProc == NULL) {
		ctx->glIsTextureProc = (PFNGLISTEXTUREPROC)gl_wrap_get_pointer("glIsTexture");
	}
	return ctx->glIsTextureProc(texture);
};
void gl_wrap_context_glLightModelx(gl_wrap_context* ctx, GLenum pname, GLfixed param) {
	if(ctx->glLightModelxProc == NULL) {
		ctx->glLightModelxProc = (PFNGLLIGHTMODELXPROC)gl_wrap_get_pointer("glLightModelx");
	}
	ctx->glLightModelxProc(pname, param);
};
void gl_wrap_context_glLightModelxv(gl_wrap_context* ctx, GLenum pname, GLfixed* param) {
	if(ctx->glLightModelxvProc == NULL) {
		ctx->glLightModelxvProc = (PFNGLLIGHTMODELXVPROC)gl_wrap_get_pointer("glLightModelxv");
	}
	ctx->glLightModelxvProc(pname, param);
};
void gl_wrap_context_glLightx(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfixed param) {
	if(ctx->glLightxProc == NULL) {
		ctx->glLightxProc = (PFNGLLIGHTXPROC)gl_wrap_get_pointer("glLightx");
	}
	ctx->glLightxProc(light, pname, param);
};
void gl_wrap_context_glLightxv(gl_wrap_context* ctx, GLenum light, GLenum pname, GLfixed* params) {
	if(ctx->glLightxvProc == NULL) {
		ctx->glLightxvProc = (PFNGLLIGHTXVPROC)gl_wrap_get_pointer("glLightxv");
	}
	ctx->glLightxvProc(light, pname, params);
};
void gl_wrap_context_glLineWidthx(gl_wrap_context* ctx, GLfixed width) {
	if(ctx->glLineWidthxProc == NULL) {
		ctx->glLineWidthxProc = (PFNGLLINEWIDTHXPROC)gl_wrap_get_pointer("glLineWidthx");
	}
	ctx->glLineWidthxProc(width);
};
void gl_wrap_context_glLoadIdentity(gl_wrap_context* ctx) {
	if(ctx->glLoadIdentityProc == NULL) {
		ctx->glLoadIdentityProc = (PFNGLLOADIDENTITYPROC)gl_wrap_get_pointer("glLoadIdentity");
	}
	ctx->glLoadIdentityProc();
};
void gl_wrap_context_glLoadMatrixx(gl_wrap_context* ctx, GLfixed* m) {
	if(ctx->glLoadMatrixxProc == NULL) {
		ctx->glLoadMatrixxProc = (PFNGLLOADMATRIXXPROC)gl_wrap_get_pointer("glLoadMatrixx");
	}
	ctx->glLoadMatrixxProc(m);
};
void gl_wrap_context_glLogicOp(gl_wrap_context* ctx, GLenum opcode) {
	if(ctx->glLogicOpProc == NULL) {
		ctx->glLogicOpProc = (PFNGLLOGICOPPROC)gl_wrap_get_pointer("glLogicOp");
	}
	ctx->glLogicOpProc(opcode);
};
void gl_wrap_context_glMaterialx(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfixed param) {
	if(ctx->glMaterialxProc == NULL) {
		ctx->glMaterialxProc = (PFNGLMATERIALXPROC)gl_wrap_get_pointer("glMaterialx");
	}
	ctx->glMaterialxProc(face, pname, param);
};
void gl_wrap_context_glMaterialxv(gl_wrap_context* ctx, GLenum face, GLenum pname, GLfixed* param) {
	if(ctx->glMaterialxvProc == NULL) {
		ctx->glMaterialxvProc = (PFNGLMATERIALXVPROC)gl_wrap_get_pointer("glMaterialxv");
	}
	ctx->glMaterialxvProc(face, pname, param);
};
void gl_wrap_context_glMatrixMode(gl_wrap_context* ctx, GLenum mode) {
	if(ctx->glMatrixModeProc == NULL) {
		ctx->glMatrixModeProc = (PFNGLMATRIXMODEPROC)gl_wrap_get_pointer("glMatrixMode");
	}
	ctx->glMatrixModeProc(mode);
};
void gl_wrap_context_glMultMatrixx(gl_wrap_context* ctx, GLfixed* m) {
	if(ctx->glMultMatrixxProc == NULL) {
		ctx->glMultMatrixxProc = (PFNGLMULTMATRIXXPROC)gl_wrap_get_pointer("glMultMatrixx");
	}
	ctx->glMultMatrixxProc(m);
};
void gl_wrap_context_glMultiTexCoord4x(gl_wrap_context* ctx, GLenum texture, GLfixed s, GLfixed t, GLfixed r, GLfixed q) {
	if(ctx->glMultiTexCoord4xProc == NULL) {
		ctx->glMultiTexCoord4xProc = (PFNGLMULTITEXCOORD4XPROC)gl_wrap_get_pointer("glMultiTexCoord4x");
	}
	ctx->glMultiTexCoord4xProc(texture, s, t, r, q);
};
void gl_wrap_context_glNormal3x(gl_wrap_context* ctx, GLfixed nx, GLfixed ny, GLfixed nz) {
	if(ctx->glNormal3xProc == NULL) {
		ctx->glNormal3xProc = (PFNGLNORMAL3XPROC)gl_wrap_get_pointer("glNormal3x");
	}
	ctx->glNormal3xProc(nx, ny, nz);
};
void gl_wrap_context_glNormalPointer(gl_wrap_context* ctx, GLenum type, GLsizei stride, void* pointer) {
	if(ctx->glNormalPointerProc == NULL) {
		ctx->glNormalPointerProc = (PFNGLNORMALPOINTERPROC)gl_wrap_get_pointer("glNormalPointer");
	}
	ctx->glNormalPointerProc(type, stride, pointer);
};
void gl_wrap_context_glOrthox(gl_wrap_context* ctx, GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f) {
	if(ctx->glOrthoxProc == NULL) {
		ctx->glOrthoxProc = (PFNGLORTHOXPROC)gl_wrap_get_pointer("glOrthox");
	}
	ctx->glOrthoxProc(l, r, b, t, n, f);
};
void gl_wrap_context_glPixelStorei(gl_wrap_context* ctx, GLenum pname, GLint param) {
	if(ctx->glPixelStoreiProc == NULL) {
		ctx->glPixelStoreiProc = (PFNGLPIXELSTOREIPROC)gl_wrap_get_pointer("glPixelStorei");
	}
	ctx->glPixelStoreiProc(pname, param);
};
void gl_wrap_context_glPointParameterx(gl_wrap_context* ctx, GLenum pname, GLfixed param) {
	if(ctx->glPointParameterxProc == NULL) {
		ctx->glPointParameterxProc = (PFNGLPOINTPARAMETERXPROC)gl_wrap_get_pointer("glPointParameterx");
	}
	ctx->glPointParameterxProc(pname, param);
};
void gl_wrap_context_glPointParameterxv(gl_wrap_context* ctx, GLenum pname, GLfixed* params) {
	if(ctx->glPointParameterxvProc == NULL) {
		ctx->glPointParameterxvProc = (PFNGLPOINTPARAMETERXVPROC)gl_wrap_get_pointer("glPointParameterxv");
	}
	ctx->glPointParameterxvProc(pname, params);
};
void gl_wrap_context_glPointSizex(gl_wrap_context* ctx, GLfixed size) {
	if(ctx->glPointSizexProc == NULL) {
		ctx->glPointSizexProc = (PFNGLPOINTSIZEXPROC)gl_wrap_get_pointer("glPointSizex");
	}
	ctx->glPointSizexProc(size);
};
void gl_wrap_context_glPolygonOffsetx(gl_wrap_context* ctx, GLfixed factor, GLfixed units) {
	if(ctx->glPolygonOffsetxProc == NULL) {
		ctx->glPolygonOffsetxProc = (PFNGLPOLYGONOFFSETXPROC)gl_wrap_get_pointer("glPolygonOffsetx");
	}
	ctx->glPolygonOffsetxProc(factor, units);
};
void gl_wrap_context_glPopMatrix(gl_wrap_context* ctx) {
	if(ctx->glPopMatrixProc == NULL) {
		ctx->glPopMatrixProc = (PFNGLPOPMATRIXPROC)gl_wrap_get_pointer("glPopMatrix");
	}
	ctx->glPopMatrixProc();
};
void gl_wrap_context_glPushMatrix(gl_wrap_context* ctx) {
	if(ctx->glPushMatrixProc == NULL) {
		ctx->glPushMatrixProc = (PFNGLPUSHMATRIXPROC)gl_wrap_get_pointer("glPushMatrix");
	}
	ctx->glPushMatrixProc();
};
void gl_wrap_context_glReadPixels(gl_wrap_context* ctx, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, void* pixels) {
	if(ctx->glReadPixelsProc == NULL) {
		ctx->glReadPixelsProc = (PFNGLREADPIXELSPROC)gl_wrap_get_pointer("glReadPixels");
	}
	ctx->glReadPixelsProc(x, y, width, height, format, type, pixels);
};
void gl_wrap_context_glRotatex(gl_wrap_context* ctx, GLfixed angle, GLfixed x, GLfixed y, GLfixed z) {
	if(ctx->glRotatexProc == NULL) {
		ctx->glRotatexProc = (PFNGLROTATEXPROC)gl_wrap_get_pointer("glRotatex");
	}
	ctx->glRotatexProc(angle, x, y, z);
};
void gl_wrap_context_glSampleCoverage(gl_wrap_context* ctx, GLfloat value, GLboolean invert) {
	if(ctx->glSampleCoverageProc == NULL) {
		ctx->glSampleCoverageProc = (PFNGLSAMPLECOVERAGEPROC)gl_wrap_get_pointer("glSampleCoverage");
	}
	ctx->glSampleCoverageProc(value, invert);
};
void gl_wrap_context_glSampleCoveragex(gl_wrap_context* ctx, GLclampx value, GLboolean invert) {
	if(ctx->glSampleCoveragexProc == NULL) {
		ctx->glSampleCoveragexProc = (PFNGLSAMPLECOVERAGEXPROC)gl_wrap_get_pointer("glSampleCoveragex");
	}
	ctx->glSampleCoveragexProc(value, invert);
};
void gl_wrap_context_glScalex(gl_wrap_context* ctx, GLfixed x, GLfixed y, GLfixed z) {
	if(ctx->glScalexProc == NULL) {
		ctx->glScalexProc = (PFNGLSCALEXPROC)gl_wrap_get_pointer("glScalex");
	}
	ctx->glScalexProc(x, y, z);
};
void gl_wrap_context_glScissor(gl_wrap_context* ctx, GLint x, GLint y, GLsizei width, GLsizei height) {
	if(ctx->glScissorProc == NULL) {
		ctx->glScissorProc = (PFNGLSCISSORPROC)gl_wrap_get_pointer("glScissor");
	}
	ctx->glScissorProc(x, y, width, height);
};
void gl_wrap_context_glShadeModel(gl_wrap_context* ctx, GLenum mode) {
	if(ctx->glShadeModelProc == NULL) {
		ctx->glShadeModelProc = (PFNGLSHADEMODELPROC)gl_wrap_get_pointer("glShadeModel");
	}
	ctx->glShadeModelProc(mode);
};
void gl_wrap_context_glStencilFunc(gl_wrap_context* ctx, GLenum func, GLint ref, GLuint mask) {
	if(ctx->glStencilFuncProc == NULL) {
		ctx->glStencilFuncProc = (PFNGLSTENCILFUNCPROC)gl_wrap_get_pointer("glStencilFunc");
	}
	ctx->glStencilFuncProc(func, ref, mask);
};
void gl_wrap_context_glStencilMask(gl_wrap_context* ctx, GLuint mask) {
	if(ctx->glStencilMaskProc == NULL) {
		ctx->glStencilMaskProc = (PFNGLSTENCILMASKPROC)gl_wrap_get_pointer("glStencilMask");
	}
	ctx->glStencilMaskProc(mask);
};
void gl_wrap_context_glStencilOp(gl_wrap_context* ctx, GLenum fail, GLenum zfail, GLenum zpass) {
	if(ctx->glStencilOpProc == NULL) {
		ctx->glStencilOpProc = (PFNGLSTENCILOPPROC)gl_wrap_get_pointer("glStencilOp");
	}
	ctx->glStencilOpProc(fail, zfail, zpass);
};
void gl_wrap_context_glTexCoordPointer(gl_wrap_context* ctx, GLint size, GLenum type, GLsizei stride, void* pointer) {
	if(ctx->glTexCoordPointerProc == NULL) {
		ctx->glTexCoordPointerProc = (PFNGLTEXCOORDPOINTERPROC)gl_wrap_get_pointer("glTexCoordPointer");
	}
	ctx->glTexCoordPointerProc(size, type, stride, pointer);
};
void gl_wrap_context_glTexEnvi(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint param) {
	if(ctx->glTexEnviProc == NULL) {
		ctx->glTexEnviProc = (PFNGLTEXENVIPROC)gl_wrap_get_pointer("glTexEnvi");
	}
	ctx->glTexEnviProc(target, pname, param);
};
void gl_wrap_context_glTexEnvx(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed param) {
	if(ctx->glTexEnvxProc == NULL) {
		ctx->glTexEnvxProc = (PFNGLTEXENVXPROC)gl_wrap_get_pointer("glTexEnvx");
	}
	ctx->glTexEnvxProc(target, pname, param);
};
void gl_wrap_context_glTexEnviv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params) {
	if(ctx->glTexEnvivProc == NULL) {
		ctx->glTexEnvivProc = (PFNGLTEXENVIVPROC)gl_wrap_get_pointer("glTexEnviv");
	}
	ctx->glTexEnvivProc(target, pname, params);
};
void gl_wrap_context_glTexEnvxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params) {
	if(ctx->glTexEnvxvProc == NULL) {
		ctx->glTexEnvxvProc = (PFNGLTEXENVXVPROC)gl_wrap_get_pointer("glTexEnvxv");
	}
	ctx->glTexEnvxvProc(target, pname, params);
};
void gl_wrap_context_glTexImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, void* pixels) {
	if(ctx->glTexImage2DProc == NULL) {
		ctx->glTexImage2DProc = (PFNGLTEXIMAGE2DPROC)gl_wrap_get_pointer("glTexImage2D");
	}
	ctx->glTexImage2DProc(target, level, internalformat, width, height, border, format, type, pixels);
};
void gl_wrap_context_glTexParameteri(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint param) {
	if(ctx->glTexParameteriProc == NULL) {
		ctx->glTexParameteriProc = (PFNGLTEXPARAMETERIPROC)gl_wrap_get_pointer("glTexParameteri");
	}
	ctx->glTexParameteriProc(target, pname, param);
};
void gl_wrap_context_glTexParameterx(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed param) {
	if(ctx->glTexParameterxProc == NULL) {
		ctx->glTexParameterxProc = (PFNGLTEXPARAMETERXPROC)gl_wrap_get_pointer("glTexParameterx");
	}
	ctx->glTexParameterxProc(target, pname, param);
};
void gl_wrap_context_glTexParameteriv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLint* params) {
	if(ctx->glTexParameterivProc == NULL) {
		ctx->glTexParameterivProc = (PFNGLTEXPARAMETERIVPROC)gl_wrap_get_pointer("glTexParameteriv");
	}
	ctx->glTexParameterivProc(target, pname, params);
};
void gl_wrap_context_glTexParameterxv(gl_wrap_context* ctx, GLenum target, GLenum pname, GLfixed* params) {
	if(ctx->glTexParameterxvProc == NULL) {
		ctx->glTexParameterxvProc = (PFNGLTEXPARAMETERXVPROC)gl_wrap_get_pointer("glTexParameterxv");
	}
	ctx->glTexParameterxvProc(target, pname, params);
};
void gl_wrap_context_glTexSubImage2D(gl_wrap_context* ctx, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, void* pixels) {
	if(ctx->glTexSubImage2DProc == NULL) {
		ctx->glTexSubImage2DProc = (PFNGLTEXSUBIMAGE2DPROC)gl_wrap_get_pointer("glTexSubImage2D");
	}
	ctx->glTexSubImage2DProc(target, level, xoffset, yoffset, width, height, format, type, pixels);
};
void gl_wrap_context_glTranslatex(gl_wrap_context* ctx, GLfixed x, GLfixed y, GLfixed z) {
	if(ctx->glTranslatexProc == NULL) {
		ctx->glTranslatexProc = (PFNGLTRANSLATEXPROC)gl_wrap_get_pointer("glTranslatex");
	}
	ctx->glTranslatexProc(x, y, z);
};
void gl_wrap_context_glVertexPointer(gl_wrap_context* ctx, GLint size, GLenum type, GLsizei stride, void* pointer) {
	if(ctx->glVertexPointerProc == NULL) {
		ctx->glVertexPointerProc = (PFNGLVERTEXPOINTERPROC)gl_wrap_get_pointer("glVertexPointer");
	}
	ctx->glVertexPointerProc(size, type, stride, pointer);
};
void gl_wrap_context_glViewport(gl_wrap_context* ctx, GLint x, GLint y, GLsizei width, GLsizei height) {
	if(ctx->glViewportProc == NULL) {
		ctx->glViewportProc = (PFNGLVIEWPORTPROC)gl_wrap_get_pointer("glViewport");
	}
	ctx->glViewportProc(x, y, width, height);
};



// Handler functions are defined for each OpenGL call; each handler function
// takes the OpenGL context struct and a pointer to the same OpenGL function's
// arguments stored in a struct.
//
// Each handler function is responsible for invoking the OpenGL function with
// the proper parameters.
//
// All function handlers are placed with respect to order in the defined jump
// table (see below), which allows batched OpenGL calls to be made without
// using a large (and costly) switch statement.
//
// Handler functions are not defined for OpenGL functions which return any
// value, as these function calls cannot be batched (see the Go documentation
// for this package, which explains this in more detail).
inline void gl_wrap_handler_glAlphaFunc(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glAlphaFunc_args args = *(gl_wrap_handler_glAlphaFunc_args*)argsPtr;
	gl_wrap_context_glAlphaFunc(ctx, args.func, args.ref);
}

inline void gl_wrap_handler_glClearColor(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClearColor_args args = *(gl_wrap_handler_glClearColor_args*)argsPtr;
	gl_wrap_context_glClearColor(ctx, args.red, args.green, args.blue, args.alpha);
}

inline void gl_wrap_handler_glClearDepthf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClearDepthf_args args = *(gl_wrap_handler_glClearDepthf_args*)argsPtr;
	gl_wrap_context_glClearDepthf(ctx, args.d);
}

inline void gl_wrap_handler_glClipPlanef(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClipPlanef_args args = *(gl_wrap_handler_glClipPlanef_args*)argsPtr;
	gl_wrap_context_glClipPlanef(ctx, args.p, args.eqn);
}

inline void gl_wrap_handler_glColor4f(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glColor4f_args args = *(gl_wrap_handler_glColor4f_args*)argsPtr;
	gl_wrap_context_glColor4f(ctx, args.red, args.green, args.blue, args.alpha);
}

inline void gl_wrap_handler_glDepthRangef(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDepthRangef_args args = *(gl_wrap_handler_glDepthRangef_args*)argsPtr;
	gl_wrap_context_glDepthRangef(ctx, args.n, args.f);
}

inline void gl_wrap_handler_glFogf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFogf_args args = *(gl_wrap_handler_glFogf_args*)argsPtr;
	gl_wrap_context_glFogf(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glFogfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFogfv_args args = *(gl_wrap_handler_glFogfv_args*)argsPtr;
	gl_wrap_context_glFogfv(ctx, args.pname, args.params);
}

inline void gl_wrap_handler_glFrustumf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFrustumf_args args = *(gl_wrap_handler_glFrustumf_args*)argsPtr;
	gl_wrap_context_glFrustumf(ctx, args.l, args.r, args.b, args.t, args.n, args.f);
}

inline void gl_wrap_handler_glGetClipPlanef(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetClipPlanef_args args = *(gl_wrap_handler_glGetClipPlanef_args*)argsPtr;
	gl_wrap_context_glGetClipPlanef(ctx, args.plane, args.equation);
}

inline void gl_wrap_handler_glGetFloatv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetFloatv_args args = *(gl_wrap_handler_glGetFloatv_args*)argsPtr;
	gl_wrap_context_glGetFloatv(ctx, args.pname, args.data);
}

inline void gl_wrap_handler_glGetLightfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetLightfv_args args = *(gl_wrap_handler_glGetLightfv_args*)argsPtr;
	gl_wrap_context_glGetLightfv(ctx, args.light, args.pname, args.params);
}

inline void gl_wrap_handler_glGetMaterialfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetMaterialfv_args args = *(gl_wrap_handler_glGetMaterialfv_args*)argsPtr;
	gl_wrap_context_glGetMaterialfv(ctx, args.face, args.pname, args.params);
}

inline void gl_wrap_handler_glGetTexEnvfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetTexEnvfv_args args = *(gl_wrap_handler_glGetTexEnvfv_args*)argsPtr;
	gl_wrap_context_glGetTexEnvfv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glGetTexParameterfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetTexParameterfv_args args = *(gl_wrap_handler_glGetTexParameterfv_args*)argsPtr;
	gl_wrap_context_glGetTexParameterfv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glLightModelf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightModelf_args args = *(gl_wrap_handler_glLightModelf_args*)argsPtr;
	gl_wrap_context_glLightModelf(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glLightModelfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightModelfv_args args = *(gl_wrap_handler_glLightModelfv_args*)argsPtr;
	gl_wrap_context_glLightModelfv(ctx, args.pname, args.params);
}

inline void gl_wrap_handler_glLightf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightf_args args = *(gl_wrap_handler_glLightf_args*)argsPtr;
	gl_wrap_context_glLightf(ctx, args.light, args.pname, args.param);
}

inline void gl_wrap_handler_glLightfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightfv_args args = *(gl_wrap_handler_glLightfv_args*)argsPtr;
	gl_wrap_context_glLightfv(ctx, args.light, args.pname, args.params);
}

inline void gl_wrap_handler_glLineWidth(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLineWidth_args args = *(gl_wrap_handler_glLineWidth_args*)argsPtr;
	gl_wrap_context_glLineWidth(ctx, args.width);
}

inline void gl_wrap_handler_glLoadMatrixf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLoadMatrixf_args args = *(gl_wrap_handler_glLoadMatrixf_args*)argsPtr;
	gl_wrap_context_glLoadMatrixf(ctx, args.m);
}

inline void gl_wrap_handler_glMaterialf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMaterialf_args args = *(gl_wrap_handler_glMaterialf_args*)argsPtr;
	gl_wrap_context_glMaterialf(ctx, args.face, args.pname, args.param);
}

inline void gl_wrap_handler_glMaterialfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMaterialfv_args args = *(gl_wrap_handler_glMaterialfv_args*)argsPtr;
	gl_wrap_context_glMaterialfv(ctx, args.face, args.pname, args.params);
}

inline void gl_wrap_handler_glMultMatrixf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMultMatrixf_args args = *(gl_wrap_handler_glMultMatrixf_args*)argsPtr;
	gl_wrap_context_glMultMatrixf(ctx, args.m);
}

inline void gl_wrap_handler_glMultiTexCoord4f(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMultiTexCoord4f_args args = *(gl_wrap_handler_glMultiTexCoord4f_args*)argsPtr;
	gl_wrap_context_glMultiTexCoord4f(ctx, args.target, args.s, args.t, args.r, args.q);
}

inline void gl_wrap_handler_glNormal3f(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glNormal3f_args args = *(gl_wrap_handler_glNormal3f_args*)argsPtr;
	gl_wrap_context_glNormal3f(ctx, args.nx, args.ny, args.nz);
}

inline void gl_wrap_handler_glOrthof(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glOrthof_args args = *(gl_wrap_handler_glOrthof_args*)argsPtr;
	gl_wrap_context_glOrthof(ctx, args.l, args.r, args.b, args.t, args.n, args.f);
}

inline void gl_wrap_handler_glPointParameterf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPointParameterf_args args = *(gl_wrap_handler_glPointParameterf_args*)argsPtr;
	gl_wrap_context_glPointParameterf(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glPointParameterfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPointParameterfv_args args = *(gl_wrap_handler_glPointParameterfv_args*)argsPtr;
	gl_wrap_context_glPointParameterfv(ctx, args.pname, args.params);
}

inline void gl_wrap_handler_glPointSize(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPointSize_args args = *(gl_wrap_handler_glPointSize_args*)argsPtr;
	gl_wrap_context_glPointSize(ctx, args.size);
}

inline void gl_wrap_handler_glPolygonOffset(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPolygonOffset_args args = *(gl_wrap_handler_glPolygonOffset_args*)argsPtr;
	gl_wrap_context_glPolygonOffset(ctx, args.factor, args.units);
}

inline void gl_wrap_handler_glRotatef(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glRotatef_args args = *(gl_wrap_handler_glRotatef_args*)argsPtr;
	gl_wrap_context_glRotatef(ctx, args.angle, args.x, args.y, args.z);
}

inline void gl_wrap_handler_glScalef(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glScalef_args args = *(gl_wrap_handler_glScalef_args*)argsPtr;
	gl_wrap_context_glScalef(ctx, args.x, args.y, args.z);
}

inline void gl_wrap_handler_glTexEnvf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexEnvf_args args = *(gl_wrap_handler_glTexEnvf_args*)argsPtr;
	gl_wrap_context_glTexEnvf(ctx, args.target, args.pname, args.param);
}

inline void gl_wrap_handler_glTexEnvfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexEnvfv_args args = *(gl_wrap_handler_glTexEnvfv_args*)argsPtr;
	gl_wrap_context_glTexEnvfv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glTexParameterf(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexParameterf_args args = *(gl_wrap_handler_glTexParameterf_args*)argsPtr;
	gl_wrap_context_glTexParameterf(ctx, args.target, args.pname, args.param);
}

inline void gl_wrap_handler_glTexParameterfv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexParameterfv_args args = *(gl_wrap_handler_glTexParameterfv_args*)argsPtr;
	gl_wrap_context_glTexParameterfv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glTranslatef(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTranslatef_args args = *(gl_wrap_handler_glTranslatef_args*)argsPtr;
	gl_wrap_context_glTranslatef(ctx, args.x, args.y, args.z);
}

inline void gl_wrap_handler_glActiveTexture(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glActiveTexture_args args = *(gl_wrap_handler_glActiveTexture_args*)argsPtr;
	gl_wrap_context_glActiveTexture(ctx, args.texture);
}

inline void gl_wrap_handler_glAlphaFuncx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glAlphaFuncx_args args = *(gl_wrap_handler_glAlphaFuncx_args*)argsPtr;
	gl_wrap_context_glAlphaFuncx(ctx, args.func, args.ref);
}

inline void gl_wrap_handler_glBindBuffer(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glBindBuffer_args args = *(gl_wrap_handler_glBindBuffer_args*)argsPtr;
	gl_wrap_context_glBindBuffer(ctx, args.target, args.buffer);
}

inline void gl_wrap_handler_glBindTexture(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glBindTexture_args args = *(gl_wrap_handler_glBindTexture_args*)argsPtr;
	gl_wrap_context_glBindTexture(ctx, args.target, args.texture);
}

inline void gl_wrap_handler_glBlendFunc(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glBlendFunc_args args = *(gl_wrap_handler_glBlendFunc_args*)argsPtr;
	gl_wrap_context_glBlendFunc(ctx, args.sfactor, args.dfactor);
}

inline void gl_wrap_handler_glBufferData(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glBufferData_args args = *(gl_wrap_handler_glBufferData_args*)argsPtr;
	gl_wrap_context_glBufferData(ctx, args.target, args.size, args.data, args.usage);
}

inline void gl_wrap_handler_glBufferSubData(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glBufferSubData_args args = *(gl_wrap_handler_glBufferSubData_args*)argsPtr;
	gl_wrap_context_glBufferSubData(ctx, args.target, args.offset, args.size, args.data);
}

inline void gl_wrap_handler_glClear(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClear_args args = *(gl_wrap_handler_glClear_args*)argsPtr;
	gl_wrap_context_glClear(ctx, args.mask);
}

inline void gl_wrap_handler_glClearColorx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClearColorx_args args = *(gl_wrap_handler_glClearColorx_args*)argsPtr;
	gl_wrap_context_glClearColorx(ctx, args.red, args.green, args.blue, args.alpha);
}

inline void gl_wrap_handler_glClearDepthx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClearDepthx_args args = *(gl_wrap_handler_glClearDepthx_args*)argsPtr;
	gl_wrap_context_glClearDepthx(ctx, args.depth);
}

inline void gl_wrap_handler_glClearStencil(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClearStencil_args args = *(gl_wrap_handler_glClearStencil_args*)argsPtr;
	gl_wrap_context_glClearStencil(ctx, args.s);
}

inline void gl_wrap_handler_glClientActiveTexture(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClientActiveTexture_args args = *(gl_wrap_handler_glClientActiveTexture_args*)argsPtr;
	gl_wrap_context_glClientActiveTexture(ctx, args.texture);
}

inline void gl_wrap_handler_glClipPlanex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glClipPlanex_args args = *(gl_wrap_handler_glClipPlanex_args*)argsPtr;
	gl_wrap_context_glClipPlanex(ctx, args.plane, args.equation);
}

inline void gl_wrap_handler_glColor4ub(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glColor4ub_args args = *(gl_wrap_handler_glColor4ub_args*)argsPtr;
	gl_wrap_context_glColor4ub(ctx, args.red, args.green, args.blue, args.alpha);
}

inline void gl_wrap_handler_glColor4x(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glColor4x_args args = *(gl_wrap_handler_glColor4x_args*)argsPtr;
	gl_wrap_context_glColor4x(ctx, args.red, args.green, args.blue, args.alpha);
}

inline void gl_wrap_handler_glColorMask(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glColorMask_args args = *(gl_wrap_handler_glColorMask_args*)argsPtr;
	gl_wrap_context_glColorMask(ctx, args.red, args.green, args.blue, args.alpha);
}

inline void gl_wrap_handler_glColorPointer(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glColorPointer_args args = *(gl_wrap_handler_glColorPointer_args*)argsPtr;
	gl_wrap_context_glColorPointer(ctx, args.size, args.type, args.stride, args.pointer);
}

inline void gl_wrap_handler_glCompressedTexImage2D(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glCompressedTexImage2D_args args = *(gl_wrap_handler_glCompressedTexImage2D_args*)argsPtr;
	gl_wrap_context_glCompressedTexImage2D(ctx, args.target, args.level, args.internalformat, args.width, args.height, args.border, args.imageSize, args.data);
}

inline void gl_wrap_handler_glCompressedTexSubImage2D(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glCompressedTexSubImage2D_args args = *(gl_wrap_handler_glCompressedTexSubImage2D_args*)argsPtr;
	gl_wrap_context_glCompressedTexSubImage2D(ctx, args.target, args.level, args.xoffset, args.yoffset, args.width, args.height, args.format, args.imageSize, args.data);
}

inline void gl_wrap_handler_glCopyTexImage2D(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glCopyTexImage2D_args args = *(gl_wrap_handler_glCopyTexImage2D_args*)argsPtr;
	gl_wrap_context_glCopyTexImage2D(ctx, args.target, args.level, args.internalformat, args.x, args.y, args.width, args.height, args.border);
}

inline void gl_wrap_handler_glCopyTexSubImage2D(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glCopyTexSubImage2D_args args = *(gl_wrap_handler_glCopyTexSubImage2D_args*)argsPtr;
	gl_wrap_context_glCopyTexSubImage2D(ctx, args.target, args.level, args.xoffset, args.yoffset, args.x, args.y, args.width, args.height);
}

inline void gl_wrap_handler_glCullFace(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glCullFace_args args = *(gl_wrap_handler_glCullFace_args*)argsPtr;
	gl_wrap_context_glCullFace(ctx, args.mode);
}

inline void gl_wrap_handler_glDeleteBuffers(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDeleteBuffers_args args = *(gl_wrap_handler_glDeleteBuffers_args*)argsPtr;
	gl_wrap_context_glDeleteBuffers(ctx, args.n, args.buffers);
}

inline void gl_wrap_handler_glDeleteTextures(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDeleteTextures_args args = *(gl_wrap_handler_glDeleteTextures_args*)argsPtr;
	gl_wrap_context_glDeleteTextures(ctx, args.n, args.textures);
}

inline void gl_wrap_handler_glDepthFunc(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDepthFunc_args args = *(gl_wrap_handler_glDepthFunc_args*)argsPtr;
	gl_wrap_context_glDepthFunc(ctx, args.func);
}

inline void gl_wrap_handler_glDepthMask(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDepthMask_args args = *(gl_wrap_handler_glDepthMask_args*)argsPtr;
	gl_wrap_context_glDepthMask(ctx, args.flag);
}

inline void gl_wrap_handler_glDepthRangex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDepthRangex_args args = *(gl_wrap_handler_glDepthRangex_args*)argsPtr;
	gl_wrap_context_glDepthRangex(ctx, args.n, args.f);
}

inline void gl_wrap_handler_glDisable(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDisable_args args = *(gl_wrap_handler_glDisable_args*)argsPtr;
	gl_wrap_context_glDisable(ctx, args.cap);
}

inline void gl_wrap_handler_glDisableClientState(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDisableClientState_args args = *(gl_wrap_handler_glDisableClientState_args*)argsPtr;
	gl_wrap_context_glDisableClientState(ctx, args.array);
}

inline void gl_wrap_handler_glDrawArrays(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDrawArrays_args args = *(gl_wrap_handler_glDrawArrays_args*)argsPtr;
	gl_wrap_context_glDrawArrays(ctx, args.mode, args.first, args.count);
}

inline void gl_wrap_handler_glDrawElements(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glDrawElements_args args = *(gl_wrap_handler_glDrawElements_args*)argsPtr;
	gl_wrap_context_glDrawElements(ctx, args.mode, args.count, args.type, args.indices);
}

inline void gl_wrap_handler_glEnable(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glEnable_args args = *(gl_wrap_handler_glEnable_args*)argsPtr;
	gl_wrap_context_glEnable(ctx, args.cap);
}

inline void gl_wrap_handler_glEnableClientState(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glEnableClientState_args args = *(gl_wrap_handler_glEnableClientState_args*)argsPtr;
	gl_wrap_context_glEnableClientState(ctx, args.array);
}

inline void gl_wrap_handler_glFinish(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFinish_args args = *(gl_wrap_handler_glFinish_args*)argsPtr;
	gl_wrap_context_glFinish(ctx);
}

inline void gl_wrap_handler_glFlush(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFlush_args args = *(gl_wrap_handler_glFlush_args*)argsPtr;
	gl_wrap_context_glFlush(ctx);
}

inline void gl_wrap_handler_glFogx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFogx_args args = *(gl_wrap_handler_glFogx_args*)argsPtr;
	gl_wrap_context_glFogx(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glFogxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFogxv_args args = *(gl_wrap_handler_glFogxv_args*)argsPtr;
	gl_wrap_context_glFogxv(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glFrontFace(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFrontFace_args args = *(gl_wrap_handler_glFrontFace_args*)argsPtr;
	gl_wrap_context_glFrontFace(ctx, args.mode);
}

inline void gl_wrap_handler_glFrustumx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glFrustumx_args args = *(gl_wrap_handler_glFrustumx_args*)argsPtr;
	gl_wrap_context_glFrustumx(ctx, args.l, args.r, args.b, args.t, args.n, args.f);
}

inline void gl_wrap_handler_glGetBooleanv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetBooleanv_args args = *(gl_wrap_handler_glGetBooleanv_args*)argsPtr;
	gl_wrap_context_glGetBooleanv(ctx, args.pname, args.data);
}

inline void gl_wrap_handler_glGetBufferParameteriv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetBufferParameteriv_args args = *(gl_wrap_handler_glGetBufferParameteriv_args*)argsPtr;
	gl_wrap_context_glGetBufferParameteriv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glGetClipPlanex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetClipPlanex_args args = *(gl_wrap_handler_glGetClipPlanex_args*)argsPtr;
	gl_wrap_context_glGetClipPlanex(ctx, args.plane, args.equation);
}

inline void gl_wrap_handler_glGenBuffers(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGenBuffers_args args = *(gl_wrap_handler_glGenBuffers_args*)argsPtr;
	gl_wrap_context_glGenBuffers(ctx, args.n, args.buffers);
}

inline void gl_wrap_handler_glGenTextures(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGenTextures_args args = *(gl_wrap_handler_glGenTextures_args*)argsPtr;
	gl_wrap_context_glGenTextures(ctx, args.n, args.textures);
}

inline void gl_wrap_handler_glGetFixedv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetFixedv_args args = *(gl_wrap_handler_glGetFixedv_args*)argsPtr;
	gl_wrap_context_glGetFixedv(ctx, args.pname, args.params);
}

inline void gl_wrap_handler_glGetIntegerv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetIntegerv_args args = *(gl_wrap_handler_glGetIntegerv_args*)argsPtr;
	gl_wrap_context_glGetIntegerv(ctx, args.pname, args.data);
}

inline void gl_wrap_handler_glGetLightxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetLightxv_args args = *(gl_wrap_handler_glGetLightxv_args*)argsPtr;
	gl_wrap_context_glGetLightxv(ctx, args.light, args.pname, args.params);
}

inline void gl_wrap_handler_glGetMaterialxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetMaterialxv_args args = *(gl_wrap_handler_glGetMaterialxv_args*)argsPtr;
	gl_wrap_context_glGetMaterialxv(ctx, args.face, args.pname, args.params);
}

inline void gl_wrap_handler_glGetPointerv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetPointerv_args args = *(gl_wrap_handler_glGetPointerv_args*)argsPtr;
	gl_wrap_context_glGetPointerv(ctx, args.pname, args.params);
}

inline void gl_wrap_handler_glGetTexEnviv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetTexEnviv_args args = *(gl_wrap_handler_glGetTexEnviv_args*)argsPtr;
	gl_wrap_context_glGetTexEnviv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glGetTexEnvxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetTexEnvxv_args args = *(gl_wrap_handler_glGetTexEnvxv_args*)argsPtr;
	gl_wrap_context_glGetTexEnvxv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glGetTexParameteriv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetTexParameteriv_args args = *(gl_wrap_handler_glGetTexParameteriv_args*)argsPtr;
	gl_wrap_context_glGetTexParameteriv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glGetTexParameterxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glGetTexParameterxv_args args = *(gl_wrap_handler_glGetTexParameterxv_args*)argsPtr;
	gl_wrap_context_glGetTexParameterxv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glHint(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glHint_args args = *(gl_wrap_handler_glHint_args*)argsPtr;
	gl_wrap_context_glHint(ctx, args.target, args.mode);
}

inline void gl_wrap_handler_glLightModelx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightModelx_args args = *(gl_wrap_handler_glLightModelx_args*)argsPtr;
	gl_wrap_context_glLightModelx(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glLightModelxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightModelxv_args args = *(gl_wrap_handler_glLightModelxv_args*)argsPtr;
	gl_wrap_context_glLightModelxv(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glLightx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightx_args args = *(gl_wrap_handler_glLightx_args*)argsPtr;
	gl_wrap_context_glLightx(ctx, args.light, args.pname, args.param);
}

inline void gl_wrap_handler_glLightxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLightxv_args args = *(gl_wrap_handler_glLightxv_args*)argsPtr;
	gl_wrap_context_glLightxv(ctx, args.light, args.pname, args.params);
}

inline void gl_wrap_handler_glLineWidthx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLineWidthx_args args = *(gl_wrap_handler_glLineWidthx_args*)argsPtr;
	gl_wrap_context_glLineWidthx(ctx, args.width);
}

inline void gl_wrap_handler_glLoadIdentity(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLoadIdentity_args args = *(gl_wrap_handler_glLoadIdentity_args*)argsPtr;
	gl_wrap_context_glLoadIdentity(ctx);
}

inline void gl_wrap_handler_glLoadMatrixx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLoadMatrixx_args args = *(gl_wrap_handler_glLoadMatrixx_args*)argsPtr;
	gl_wrap_context_glLoadMatrixx(ctx, args.m);
}

inline void gl_wrap_handler_glLogicOp(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glLogicOp_args args = *(gl_wrap_handler_glLogicOp_args*)argsPtr;
	gl_wrap_context_glLogicOp(ctx, args.opcode);
}

inline void gl_wrap_handler_glMaterialx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMaterialx_args args = *(gl_wrap_handler_glMaterialx_args*)argsPtr;
	gl_wrap_context_glMaterialx(ctx, args.face, args.pname, args.param);
}

inline void gl_wrap_handler_glMaterialxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMaterialxv_args args = *(gl_wrap_handler_glMaterialxv_args*)argsPtr;
	gl_wrap_context_glMaterialxv(ctx, args.face, args.pname, args.param);
}

inline void gl_wrap_handler_glMatrixMode(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMatrixMode_args args = *(gl_wrap_handler_glMatrixMode_args*)argsPtr;
	gl_wrap_context_glMatrixMode(ctx, args.mode);
}

inline void gl_wrap_handler_glMultMatrixx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMultMatrixx_args args = *(gl_wrap_handler_glMultMatrixx_args*)argsPtr;
	gl_wrap_context_glMultMatrixx(ctx, args.m);
}

inline void gl_wrap_handler_glMultiTexCoord4x(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glMultiTexCoord4x_args args = *(gl_wrap_handler_glMultiTexCoord4x_args*)argsPtr;
	gl_wrap_context_glMultiTexCoord4x(ctx, args.texture, args.s, args.t, args.r, args.q);
}

inline void gl_wrap_handler_glNormal3x(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glNormal3x_args args = *(gl_wrap_handler_glNormal3x_args*)argsPtr;
	gl_wrap_context_glNormal3x(ctx, args.nx, args.ny, args.nz);
}

inline void gl_wrap_handler_glNormalPointer(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glNormalPointer_args args = *(gl_wrap_handler_glNormalPointer_args*)argsPtr;
	gl_wrap_context_glNormalPointer(ctx, args.type, args.stride, args.pointer);
}

inline void gl_wrap_handler_glOrthox(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glOrthox_args args = *(gl_wrap_handler_glOrthox_args*)argsPtr;
	gl_wrap_context_glOrthox(ctx, args.l, args.r, args.b, args.t, args.n, args.f);
}

inline void gl_wrap_handler_glPixelStorei(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPixelStorei_args args = *(gl_wrap_handler_glPixelStorei_args*)argsPtr;
	gl_wrap_context_glPixelStorei(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glPointParameterx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPointParameterx_args args = *(gl_wrap_handler_glPointParameterx_args*)argsPtr;
	gl_wrap_context_glPointParameterx(ctx, args.pname, args.param);
}

inline void gl_wrap_handler_glPointParameterxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPointParameterxv_args args = *(gl_wrap_handler_glPointParameterxv_args*)argsPtr;
	gl_wrap_context_glPointParameterxv(ctx, args.pname, args.params);
}

inline void gl_wrap_handler_glPointSizex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPointSizex_args args = *(gl_wrap_handler_glPointSizex_args*)argsPtr;
	gl_wrap_context_glPointSizex(ctx, args.size);
}

inline void gl_wrap_handler_glPolygonOffsetx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPolygonOffsetx_args args = *(gl_wrap_handler_glPolygonOffsetx_args*)argsPtr;
	gl_wrap_context_glPolygonOffsetx(ctx, args.factor, args.units);
}

inline void gl_wrap_handler_glPopMatrix(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPopMatrix_args args = *(gl_wrap_handler_glPopMatrix_args*)argsPtr;
	gl_wrap_context_glPopMatrix(ctx);
}

inline void gl_wrap_handler_glPushMatrix(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glPushMatrix_args args = *(gl_wrap_handler_glPushMatrix_args*)argsPtr;
	gl_wrap_context_glPushMatrix(ctx);
}

inline void gl_wrap_handler_glReadPixels(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glReadPixels_args args = *(gl_wrap_handler_glReadPixels_args*)argsPtr;
	gl_wrap_context_glReadPixels(ctx, args.x, args.y, args.width, args.height, args.format, args.type, args.pixels);
}

inline void gl_wrap_handler_glRotatex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glRotatex_args args = *(gl_wrap_handler_glRotatex_args*)argsPtr;
	gl_wrap_context_glRotatex(ctx, args.angle, args.x, args.y, args.z);
}

inline void gl_wrap_handler_glSampleCoverage(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glSampleCoverage_args args = *(gl_wrap_handler_glSampleCoverage_args*)argsPtr;
	gl_wrap_context_glSampleCoverage(ctx, args.value, args.invert);
}

inline void gl_wrap_handler_glSampleCoveragex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glSampleCoveragex_args args = *(gl_wrap_handler_glSampleCoveragex_args*)argsPtr;
	gl_wrap_context_glSampleCoveragex(ctx, args.value, args.invert);
}

inline void gl_wrap_handler_glScalex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glScalex_args args = *(gl_wrap_handler_glScalex_args*)argsPtr;
	gl_wrap_context_glScalex(ctx, args.x, args.y, args.z);
}

inline void gl_wrap_handler_glScissor(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glScissor_args args = *(gl_wrap_handler_glScissor_args*)argsPtr;
	gl_wrap_context_glScissor(ctx, args.x, args.y, args.width, args.height);
}

inline void gl_wrap_handler_glShadeModel(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glShadeModel_args args = *(gl_wrap_handler_glShadeModel_args*)argsPtr;
	gl_wrap_context_glShadeModel(ctx, args.mode);
}

inline void gl_wrap_handler_glStencilFunc(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glStencilFunc_args args = *(gl_wrap_handler_glStencilFunc_args*)argsPtr;
	gl_wrap_context_glStencilFunc(ctx, args.func, args.ref, args.mask);
}

inline void gl_wrap_handler_glStencilMask(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glStencilMask_args args = *(gl_wrap_handler_glStencilMask_args*)argsPtr;
	gl_wrap_context_glStencilMask(ctx, args.mask);
}

inline void gl_wrap_handler_glStencilOp(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glStencilOp_args args = *(gl_wrap_handler_glStencilOp_args*)argsPtr;
	gl_wrap_context_glStencilOp(ctx, args.fail, args.zfail, args.zpass);
}

inline void gl_wrap_handler_glTexCoordPointer(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexCoordPointer_args args = *(gl_wrap_handler_glTexCoordPointer_args*)argsPtr;
	gl_wrap_context_glTexCoordPointer(ctx, args.size, args.type, args.stride, args.pointer);
}

inline void gl_wrap_handler_glTexEnvi(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexEnvi_args args = *(gl_wrap_handler_glTexEnvi_args*)argsPtr;
	gl_wrap_context_glTexEnvi(ctx, args.target, args.pname, args.param);
}

inline void gl_wrap_handler_glTexEnvx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexEnvx_args args = *(gl_wrap_handler_glTexEnvx_args*)argsPtr;
	gl_wrap_context_glTexEnvx(ctx, args.target, args.pname, args.param);
}

inline void gl_wrap_handler_glTexEnviv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexEnviv_args args = *(gl_wrap_handler_glTexEnviv_args*)argsPtr;
	gl_wrap_context_glTexEnviv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glTexEnvxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexEnvxv_args args = *(gl_wrap_handler_glTexEnvxv_args*)argsPtr;
	gl_wrap_context_glTexEnvxv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glTexImage2D(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexImage2D_args args = *(gl_wrap_handler_glTexImage2D_args*)argsPtr;
	gl_wrap_context_glTexImage2D(ctx, args.target, args.level, args.internalformat, args.width, args.height, args.border, args.format, args.type, args.pixels);
}

inline void gl_wrap_handler_glTexParameteri(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexParameteri_args args = *(gl_wrap_handler_glTexParameteri_args*)argsPtr;
	gl_wrap_context_glTexParameteri(ctx, args.target, args.pname, args.param);
}

inline void gl_wrap_handler_glTexParameterx(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexParameterx_args args = *(gl_wrap_handler_glTexParameterx_args*)argsPtr;
	gl_wrap_context_glTexParameterx(ctx, args.target, args.pname, args.param);
}

inline void gl_wrap_handler_glTexParameteriv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexParameteriv_args args = *(gl_wrap_handler_glTexParameteriv_args*)argsPtr;
	gl_wrap_context_glTexParameteriv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glTexParameterxv(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexParameterxv_args args = *(gl_wrap_handler_glTexParameterxv_args*)argsPtr;
	gl_wrap_context_glTexParameterxv(ctx, args.target, args.pname, args.params);
}

inline void gl_wrap_handler_glTexSubImage2D(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTexSubImage2D_args args = *(gl_wrap_handler_glTexSubImage2D_args*)argsPtr;
	gl_wrap_context_glTexSubImage2D(ctx, args.target, args.level, args.xoffset, args.yoffset, args.width, args.height, args.format, args.type, args.pixels);
}

inline void gl_wrap_handler_glTranslatex(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glTranslatex_args args = *(gl_wrap_handler_glTranslatex_args*)argsPtr;
	gl_wrap_context_glTranslatex(ctx, args.x, args.y, args.z);
}

inline void gl_wrap_handler_glVertexPointer(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glVertexPointer_args args = *(gl_wrap_handler_glVertexPointer_args*)argsPtr;
	gl_wrap_context_glVertexPointer(ctx, args.size, args.type, args.stride, args.pointer);
}

inline void gl_wrap_handler_glViewport(gl_wrap_context* ctx, void* argsPtr) {
	gl_wrap_handler_glViewport_args args = *(gl_wrap_handler_glViewport_args*)argsPtr;
	gl_wrap_context_glViewport(ctx, args.x, args.y, args.width, args.height);
}

// This is the jump table used for executing each batched OpenGL function
// without doing a large (approx. 1k cases) and costly switch statement.
gl_wrap_jump_handler gl_wrap_jump_table[] = {
	gl_wrap_handler_glAlphaFunc,
	gl_wrap_handler_glClearColor,
	gl_wrap_handler_glClearDepthf,
	gl_wrap_handler_glClipPlanef,
	gl_wrap_handler_glColor4f,
	gl_wrap_handler_glDepthRangef,
	gl_wrap_handler_glFogf,
	gl_wrap_handler_glFogfv,
	gl_wrap_handler_glFrustumf,
	gl_wrap_handler_glGetClipPlanef,
	gl_wrap_handler_glGetFloatv,
	gl_wrap_handler_glGetLightfv,
	gl_wrap_handler_glGetMaterialfv,
	gl_wrap_handler_glGetTexEnvfv,
	gl_wrap_handler_glGetTexParameterfv,
	gl_wrap_handler_glLightModelf,
	gl_wrap_handler_glLightModelfv,
	gl_wrap_handler_glLightf,
	gl_wrap_handler_glLightfv,
	gl_wrap_handler_glLineWidth,
	gl_wrap_handler_glLoadMatrixf,
	gl_wrap_handler_glMaterialf,
	gl_wrap_handler_glMaterialfv,
	gl_wrap_handler_glMultMatrixf,
	gl_wrap_handler_glMultiTexCoord4f,
	gl_wrap_handler_glNormal3f,
	gl_wrap_handler_glOrthof,
	gl_wrap_handler_glPointParameterf,
	gl_wrap_handler_glPointParameterfv,
	gl_wrap_handler_glPointSize,
	gl_wrap_handler_glPolygonOffset,
	gl_wrap_handler_glRotatef,
	gl_wrap_handler_glScalef,
	gl_wrap_handler_glTexEnvf,
	gl_wrap_handler_glTexEnvfv,
	gl_wrap_handler_glTexParameterf,
	gl_wrap_handler_glTexParameterfv,
	gl_wrap_handler_glTranslatef,
	gl_wrap_handler_glActiveTexture,
	gl_wrap_handler_glAlphaFuncx,
	gl_wrap_handler_glBindBuffer,
	gl_wrap_handler_glBindTexture,
	gl_wrap_handler_glBlendFunc,
	gl_wrap_handler_glBufferData,
	gl_wrap_handler_glBufferSubData,
	gl_wrap_handler_glClear,
	gl_wrap_handler_glClearColorx,
	gl_wrap_handler_glClearDepthx,
	gl_wrap_handler_glClearStencil,
	gl_wrap_handler_glClientActiveTexture,
	gl_wrap_handler_glClipPlanex,
	gl_wrap_handler_glColor4ub,
	gl_wrap_handler_glColor4x,
	gl_wrap_handler_glColorMask,
	gl_wrap_handler_glColorPointer,
	gl_wrap_handler_glCompressedTexImage2D,
	gl_wrap_handler_glCompressedTexSubImage2D,
	gl_wrap_handler_glCopyTexImage2D,
	gl_wrap_handler_glCopyTexSubImage2D,
	gl_wrap_handler_glCullFace,
	gl_wrap_handler_glDeleteBuffers,
	gl_wrap_handler_glDeleteTextures,
	gl_wrap_handler_glDepthFunc,
	gl_wrap_handler_glDepthMask,
	gl_wrap_handler_glDepthRangex,
	gl_wrap_handler_glDisable,
	gl_wrap_handler_glDisableClientState,
	gl_wrap_handler_glDrawArrays,
	gl_wrap_handler_glDrawElements,
	gl_wrap_handler_glEnable,
	gl_wrap_handler_glEnableClientState,
	gl_wrap_handler_glFinish,
	gl_wrap_handler_glFlush,
	gl_wrap_handler_glFogx,
	gl_wrap_handler_glFogxv,
	gl_wrap_handler_glFrontFace,
	gl_wrap_handler_glFrustumx,
	gl_wrap_handler_glGetBooleanv,
	gl_wrap_handler_glGetBufferParameteriv,
	gl_wrap_handler_glGetClipPlanex,
	gl_wrap_handler_glGenBuffers,
	gl_wrap_handler_glGenTextures,
	gl_wrap_handler_glGetFixedv,
	gl_wrap_handler_glGetIntegerv,
	gl_wrap_handler_glGetLightxv,
	gl_wrap_handler_glGetMaterialxv,
	gl_wrap_handler_glGetPointerv,
	gl_wrap_handler_glGetTexEnviv,
	gl_wrap_handler_glGetTexEnvxv,
	gl_wrap_handler_glGetTexParameteriv,
	gl_wrap_handler_glGetTexParameterxv,
	gl_wrap_handler_glHint,
	gl_wrap_handler_glLightModelx,
	gl_wrap_handler_glLightModelxv,
	gl_wrap_handler_glLightx,
	gl_wrap_handler_glLightxv,
	gl_wrap_handler_glLineWidthx,
	gl_wrap_handler_glLoadIdentity,
	gl_wrap_handler_glLoadMatrixx,
	gl_wrap_handler_glLogicOp,
	gl_wrap_handler_glMaterialx,
	gl_wrap_handler_glMaterialxv,
	gl_wrap_handler_glMatrixMode,
	gl_wrap_handler_glMultMatrixx,
	gl_wrap_handler_glMultiTexCoord4x,
	gl_wrap_handler_glNormal3x,
	gl_wrap_handler_glNormalPointer,
	gl_wrap_handler_glOrthox,
	gl_wrap_handler_glPixelStorei,
	gl_wrap_handler_glPointParameterx,
	gl_wrap_handler_glPointParameterxv,
	gl_wrap_handler_glPointSizex,
	gl_wrap_handler_glPolygonOffsetx,
	gl_wrap_handler_glPopMatrix,
	gl_wrap_handler_glPushMatrix,
	gl_wrap_handler_glReadPixels,
	gl_wrap_handler_glRotatex,
	gl_wrap_handler_glSampleCoverage,
	gl_wrap_handler_glSampleCoveragex,
	gl_wrap_handler_glScalex,
	gl_wrap_handler_glScissor,
	gl_wrap_handler_glShadeModel,
	gl_wrap_handler_glStencilFunc,
	gl_wrap_handler_glStencilMask,
	gl_wrap_handler_glStencilOp,
	gl_wrap_handler_glTexCoordPointer,
	gl_wrap_handler_glTexEnvi,
	gl_wrap_handler_glTexEnvx,
	gl_wrap_handler_glTexEnviv,
	gl_wrap_handler_glTexEnvxv,
	gl_wrap_handler_glTexImage2D,
	gl_wrap_handler_glTexParameteri,
	gl_wrap_handler_glTexParameterx,
	gl_wrap_handler_glTexParameteriv,
	gl_wrap_handler_glTexParameterxv,
	gl_wrap_handler_glTexSubImage2D,
	gl_wrap_handler_glTranslatex,
	gl_wrap_handler_glVertexPointer,
	gl_wrap_handler_glViewport,
};

// Executes the functions in a batch for a given context.
void gl_wrap_batch_exec(gl_wrap_context* ctx, gl_wrap_batch_func* funcs, int numFuncs) {
	int i;
	for(i = 0; i < numFuncs; i++) {
		// Grab the function from the array
		gl_wrap_batch_func func = funcs[i];

		// Locate the handler function in the jump table at jump_index, execute
		// it using the context and function arguments.
		gl_wrap_jump_table[func.jump_index](ctx, func.args);
	}
}
