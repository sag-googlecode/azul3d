// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl11.h"

#ifdef _WIN32
HMODULE gl11OpenGL32;

void* gl11LibGetProcAddress(char* name) {
	if(gl11OpenGL32 == NULL) {
		gl11OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl11OpenGL32, TEXT(name));
}

void* gl11GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl11Accum(gl11Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl11AlphaFunc(gl11Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl11Begin(gl11Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl11End(gl11Context* glc) {
    return glc->fnEnd();
}

void gl11Bitmap(gl11Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl11BlendFunc(gl11Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl11CallList(gl11Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl11CallLists(gl11Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl11Clear(gl11Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl11ClearAccum(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl11ClearColor(gl11Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl11ClearDepth(gl11Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl11ClearIndex(gl11Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl11ClearStencil(gl11Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl11ClipPlane(gl11Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl11Color3b(gl11Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl11Color3d(gl11Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl11Color3f(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl11Color3i(gl11Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl11Color3s(gl11Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl11Color3ub(gl11Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl11Color3ui(gl11Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl11Color3us(gl11Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl11Color4b(gl11Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl11Color4d(gl11Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl11Color4f(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl11Color4i(gl11Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl11Color4s(gl11Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl11Color4ub(gl11Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl11Color4ui(gl11Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl11Color4us(gl11Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl11Color3bv(gl11Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl11Color3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl11Color3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl11Color3iv(gl11Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl11Color3sv(gl11Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl11Color3ubv(gl11Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl11Color3uiv(gl11Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl11Color3usv(gl11Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl11Color4bv(gl11Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl11Color4dv(gl11Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl11Color4fv(gl11Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl11Color4iv(gl11Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl11Color4sv(gl11Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl11Color4ubv(gl11Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl11Color4uiv(gl11Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl11Color4usv(gl11Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl11ColorMask(gl11Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl11ColorMaterial(gl11Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl11CopyPixels(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl11CullFace(gl11Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl11DeleteLists(gl11Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl11DepthFunc(gl11Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl11DepthMask(gl11Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl11DepthRange(gl11Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl11Enable(gl11Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl11Disable(gl11Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl11DrawBuffer(gl11Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl11DrawPixels(gl11Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl11EdgeFlag(gl11Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl11EdgeFlagv(gl11Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl11EdgeFlagPointer(gl11Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl11EvalCoord1d(gl11Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl11EvalCoord1f(gl11Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl11EvalCoord2d(gl11Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl11EvalCoord2f(gl11Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl11EvalCoord1dv(gl11Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl11EvalCoord1fv(gl11Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl11EvalCoord2dv(gl11Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl11EvalCoord2fv(gl11Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl11EvalMesh1(gl11Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl11EvalMesh2(gl11Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl11EvalPoint1(gl11Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl11EvalPoint2(gl11Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl11FeedbackBuffer(gl11Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl11Finish(gl11Context* glc) {
    return glc->fnFinish();
}

void gl11Flush(gl11Context* glc) {
    return glc->fnFlush();
}

void gl11Fogf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl11Fogi(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl11Fogfv(gl11Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl11Fogiv(gl11Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl11FrontFace(gl11Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl11Frustum(gl11Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl11GenLists(gl11Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl11GetBooleanv(gl11Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl11GetDoublev(gl11Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl11GetFloatv(gl11Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl11GetIntegerv(gl11Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl11GetClipPlane(gl11Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl11GetError(gl11Context* glc) {
    return glc->fnGetError();
}

void gl11GetLightfv(gl11Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl11GetLightiv(gl11Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl11GetMapdv(gl11Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl11GetMapfv(gl11Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl11GetMapiv(gl11Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl11GetMaterialfv(gl11Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl11GetMaterialiv(gl11Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl11GetPixelMapfv(gl11Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl11GetPixelMapuiv(gl11Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl11GetPixelMapusv(gl11Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl11GetPolygonStipple(gl11Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl11GetString(gl11Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl11GetTexEnvfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl11GetTexEnviv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl11GetTexGendv(gl11Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl11GetTexGenfv(gl11Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl11GetTexGeniv(gl11Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl11GetTexImage(gl11Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl11GetTexLevelParameterfv(gl11Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl11GetTexLevelParameteriv(gl11Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl11GetTexParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl11GetTexParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl11Hint(gl11Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl11Indexd(gl11Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl11Indexf(gl11Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl11Indexi(gl11Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl11Indexs(gl11Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl11Indexdv(gl11Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl11Indexfv(gl11Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl11Indexiv(gl11Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl11Indexsv(gl11Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl11IndexMask(gl11Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl11IndexPointer(gl11Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl11InitNames(gl11Context* glc) {
    return glc->fnInitNames();
}

void gl11IsEnabled(gl11Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl11IsList(gl11Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl11Lightf(gl11Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl11Lighti(gl11Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl11Lightfv(gl11Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl11Lightiv(gl11Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl11LightModelf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl11LightModeli(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl11LightModelfv(gl11Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl11LightModeliv(gl11Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl11LineStipple(gl11Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl11LineWidth(gl11Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl11ListBase(gl11Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl11LoadIdentity(gl11Context* glc) {
    return glc->fnLoadIdentity();
}

void gl11LoadMatrixd(gl11Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl11LoadMatrixf(gl11Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl11LoadName(gl11Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl11LogicOp(gl11Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl11Map1d(gl11Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl11Map1f(gl11Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl11Map2d(gl11Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl11Map2f(gl11Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl11MapGrid1d(gl11Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl11MapGrid1f(gl11Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl11MapGrid2d(gl11Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl11MapGrid2f(gl11Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl11Materialf(gl11Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl11Materiali(gl11Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl11Materialfv(gl11Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl11Materialiv(gl11Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl11MatrixMode(gl11Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl11MultMatrixd(gl11Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl11MultMatrixf(gl11Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl11NewList(gl11Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl11EndList(gl11Context* glc) {
    return glc->fnEndList();
}

void gl11Normal3b(gl11Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl11Normal3d(gl11Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl11Normal3f(gl11Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl11Normal3i(gl11Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl11Normal3s(gl11Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl11Normal3bv(gl11Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl11Normal3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl11Normal3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl11Normal3iv(gl11Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl11Normal3sv(gl11Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl11Ortho(gl11Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl11PassThrough(gl11Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl11PixelMapfv(gl11Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl11PixelMapuiv(gl11Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl11PixelMapusv(gl11Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl11PixelStoref(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl11PixelStorei(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl11PixelTransferf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl11PixelTransferi(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl11PixelZoom(gl11Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl11PointSize(gl11Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl11PolygonMode(gl11Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl11PolygonStipple(gl11Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl11PushAttrib(gl11Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl11PopAttrib(gl11Context* glc) {
    return glc->fnPopAttrib();
}

void gl11PushMatrix(gl11Context* glc) {
    return glc->fnPushMatrix();
}

void gl11PopMatrix(gl11Context* glc) {
    return glc->fnPopMatrix();
}

void gl11PushName(gl11Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl11PopName(gl11Context* glc) {
    return glc->fnPopName();
}

void gl11RasterPos2d(gl11Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl11RasterPos2f(gl11Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl11RasterPos2i(gl11Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl11RasterPos2s(gl11Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl11RasterPos3d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl11RasterPos3f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl11RasterPos3i(gl11Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl11RasterPos3s(gl11Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl11RasterPos4d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl11RasterPos4f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl11RasterPos4i(gl11Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl11RasterPos4s(gl11Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl11RasterPos2dv(gl11Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl11RasterPos2fv(gl11Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl11RasterPos2iv(gl11Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl11RasterPos2sv(gl11Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl11RasterPos3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl11RasterPos3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl11RasterPos3iv(gl11Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl11RasterPos3sv(gl11Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl11RasterPos4dv(gl11Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl11RasterPos4fv(gl11Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl11RasterPos4iv(gl11Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl11RasterPos4sv(gl11Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl11ReadBuffer(gl11Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl11ReadPixels(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl11Rectd(gl11Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl11Rectf(gl11Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl11Recti(gl11Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl11Rects(gl11Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl11Rectdv(gl11Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl11Rectfv(gl11Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl11Rectiv(gl11Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl11Rectsv(gl11Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl11RenderMode(gl11Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl11Rotated(gl11Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl11Rotatef(gl11Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl11Scaled(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl11Scalef(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl11Scissor(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl11SelectBuffer(gl11Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl11ShadeModel(gl11Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl11StencilFunc(gl11Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl11StencilMask(gl11Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl11StencilOp(gl11Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl11TexCoord1d(gl11Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl11TexCoord1f(gl11Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl11TexCoord1i(gl11Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl11TexCoord1s(gl11Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl11TexCoord2d(gl11Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl11TexCoord2f(gl11Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl11TexCoord2i(gl11Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl11TexCoord2s(gl11Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl11TexCoord3d(gl11Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl11TexCoord3f(gl11Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl11TexCoord3i(gl11Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl11TexCoord3s(gl11Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl11TexCoord4d(gl11Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl11TexCoord4f(gl11Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl11TexCoord4i(gl11Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl11TexCoord4s(gl11Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl11TexCoord1dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl11TexCoord1fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl11TexCoord1iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl11TexCoord1sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl11TexCoord2dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl11TexCoord2fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl11TexCoord2iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl11TexCoord2sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl11TexCoord3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl11TexCoord3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl11TexCoord3iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl11TexCoord3sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl11TexCoord4dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl11TexCoord4fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl11TexCoord4iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl11TexCoord4sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl11TexEnvf(gl11Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl11TexEnvi(gl11Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl11TexEnvfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl11TexEnviv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl11TexGend(gl11Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl11TexGenf(gl11Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl11TexGeni(gl11Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl11TexGendv(gl11Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl11TexGenfv(gl11Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl11TexGeniv(gl11Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl11TexImage1D(gl11Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl11TexImage2D(gl11Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl11TexParameterf(gl11Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl11TexParameteri(gl11Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl11TexParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl11TexParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl11Translated(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl11Translatef(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl11Vertex2s(gl11Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl11Vertex2i(gl11Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl11Vertex2f(gl11Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl11Vertex2d(gl11Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl11Vertex3s(gl11Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl11Vertex3i(gl11Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl11Vertex3f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl11Vertex3d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl11Vertex4s(gl11Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl11Vertex4i(gl11Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl11Vertex4f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl11Vertex4d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl11Viewport(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl11GetConvolutionParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl11GetConvolutionParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl11AreTexturesResident(gl11Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl11ArrayElement(gl11Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl11DrawArrays(gl11Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl11DrawElements(gl11Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl11GetPointerv(gl11Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl11PolygonOffset(gl11Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl11CopyTexImage1D(gl11Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl11CopyTexImage2D(gl11Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl11CopyTexSubImage1D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl11CopyTexSubImage2D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl11BindTexture(gl11Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl11DeleteTextures(gl11Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl11GenTextures(gl11Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl11IsTexture(gl11Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl11ColorPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl11EnableClientState(gl11Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl11DisableClientState(gl11Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl11Indexub(gl11Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl11Indexubv(gl11Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl11InterleavedArrays(gl11Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl11NormalPointer(gl11Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl11PushClientAttrib(gl11Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl11PrioritizeTextures(gl11Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl11PopClientAttrib(gl11Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl11TexCoordPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl11TexSubImage1D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl11TexSubImage2D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl11VertexPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl11ColorTable(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl11ColorTableParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl11ColorTableParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl11ColorSubTable(gl11Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl11ConvolutionFilter1D(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl11ConvolutionFilter2D(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl11ConvolutionParameterf(gl11Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl11ConvolutionParameteri(gl11Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl11CopyColorTable(gl11Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl11CopyColorSubTable(gl11Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl11CopyConvolutionFilter1D(gl11Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl11CopyConvolutionFilter2D(gl11Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl11GetColorTable(gl11Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl11GetColorTableParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl11GetColorTableParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl11GetConvolutionFilter(gl11Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl11GetHistogram(gl11Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl11GetHistogramParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl11GetHistogramParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl11GetSeparableFilter(gl11Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl11Histogram(gl11Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl11Minmax(gl11Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl11MultiTexCoord1s(gl11Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl11MultiTexCoord1i(gl11Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl11MultiTexCoord1f(gl11Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl11MultiTexCoord1d(gl11Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl11MultiTexCoord2s(gl11Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl11MultiTexCoord2i(gl11Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl11MultiTexCoord2f(gl11Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl11MultiTexCoord2d(gl11Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl11MultiTexCoord3s(gl11Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl11MultiTexCoord3i(gl11Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl11MultiTexCoord3f(gl11Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl11MultiTexCoord3d(gl11Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl11MultiTexCoord4s(gl11Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl11MultiTexCoord4i(gl11Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl11MultiTexCoord4f(gl11Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl11MultiTexCoord4d(gl11Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl11MultiTexCoord1sv(gl11Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl11MultiTexCoord1iv(gl11Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl11MultiTexCoord1fv(gl11Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl11MultiTexCoord1dv(gl11Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl11MultiTexCoord2sv(gl11Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl11MultiTexCoord2iv(gl11Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl11MultiTexCoord2fv(gl11Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl11MultiTexCoord2dv(gl11Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl11MultiTexCoord3sv(gl11Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl11MultiTexCoord3iv(gl11Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl11MultiTexCoord3fv(gl11Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl11MultiTexCoord3dv(gl11Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl11MultiTexCoord4sv(gl11Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl11MultiTexCoord4iv(gl11Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl11MultiTexCoord4fv(gl11Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl11MultiTexCoord4dv(gl11Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl11ResetHistogram(gl11Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl11ResetMinmax(gl11Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl11SeparableFilter2D(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl11BlendColor(gl11Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl11BlendEquation(gl11Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl11CopyTexSubImage3D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl11DrawRangeElements(gl11Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl11TexImage3D(gl11Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl11TexSubImage3D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl11ActiveTexture(gl11Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl11ClientActiveTexture(gl11Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl11CompressedTexImage1D(gl11Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl11CompressedTexImage2D(gl11Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl11CompressedTexImage3D(gl11Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl11CompressedTexSubImage1D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl11CompressedTexSubImage2D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl11CompressedTexSubImage3D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl11GetCompressedTexImage(gl11Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl11LoadTransposeMatrixd(gl11Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl11LoadTransposeMatrixf(gl11Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl11MultTransposeMatrixd(gl11Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl11MultTransposeMatrixf(gl11Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl11SampleCoverage(gl11Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl11BlendFuncSeparate(gl11Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl11FogCoordPointer(gl11Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl11FogCoordd(gl11Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl11FogCoordf(gl11Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl11FogCoorddv(gl11Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl11FogCoordfv(gl11Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl11MultiDrawArrays(gl11Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl11MultiDrawElements(gl11Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl11PointParameterf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl11PointParameteri(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl11SecondaryColor3b(gl11Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl11SecondaryColor3s(gl11Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl11SecondaryColor3i(gl11Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl11SecondaryColor3f(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl11SecondaryColor3d(gl11Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl11SecondaryColor3ub(gl11Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl11SecondaryColor3us(gl11Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl11SecondaryColor3ui(gl11Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl11SecondaryColor3bv(gl11Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl11SecondaryColor3sv(gl11Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl11SecondaryColor3iv(gl11Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl11SecondaryColor3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl11SecondaryColor3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl11SecondaryColor3ubv(gl11Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl11SecondaryColor3usv(gl11Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl11SecondaryColor3uiv(gl11Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl11SecondaryColorPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl11WindowPos2s(gl11Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl11WindowPos2i(gl11Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl11WindowPos2f(gl11Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl11WindowPos2d(gl11Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl11WindowPos3s(gl11Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl11WindowPos3i(gl11Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl11WindowPos3f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl11WindowPos3d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl11WindowPos2sv(gl11Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl11WindowPos2iv(gl11Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl11WindowPos2fv(gl11Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl11WindowPos2dv(gl11Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl11WindowPos3sv(gl11Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl11WindowPos3iv(gl11Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl11WindowPos3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl11WindowPos3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl11BeginQuery(gl11Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl11BindBuffer(gl11Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl11BufferData(gl11Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl11BufferSubData(gl11Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl11DeleteBuffers(gl11Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl11DeleteQueries(gl11Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl11GenBuffers(gl11Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl11GenQueries(gl11Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl11GetBufferParameteriv(gl11Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl11GetBufferPointerv(gl11Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl11GetBufferSubData(gl11Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl11GetQueryObjectiv(gl11Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl11GetQueryObjectuiv(gl11Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl11GetQueryiv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl11IsBuffer(gl11Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl11IsQuery(gl11Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl11MapBuffer(gl11Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl11UnmapBuffer(gl11Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl11AttachShader(gl11Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl11BindAttribLocation(gl11Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl11BlendEquationSeperate(gl11Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl11CompileShader(gl11Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl11CreateProgram(gl11Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl11CreateShader(gl11Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl11DeleteProgram(gl11Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl11DeleteShader(gl11Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl11DetachShader(gl11Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl11EnableVertexAttribArray(gl11Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl11DisableVertexAttribArray(gl11Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl11DrawBuffers(gl11Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl11GetActiveAttrib(gl11Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl11GetActiveUniform(gl11Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl11GetAttachedShaders(gl11Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl11GetAttribLocation(gl11Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl11GetProgramiv(gl11Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl11GetProgramInfoLog(gl11Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl11GetShaderiv(gl11Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl11GetShaderInfoLog(gl11Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl11GetShaderSource(gl11Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl11GetUniformfv(gl11Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl11GetUniformiv(gl11Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl11GetUniformLocation(gl11Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl11GetVertexAttribdv(gl11Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl11GetVertexAttribfv(gl11Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl11GetVertexAttribiv(gl11Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl11GetVertexAttribPointerv(gl11Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl11IsProgram(gl11Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl11IsShader(gl11Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl11LinkProgram(gl11Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl11ShaderSource(gl11Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl11StencilFuncSeparate(gl11Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl11StencilMaskSeparate(gl11Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl11StencilOpSeparate(gl11Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl11Uniform1f(gl11Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl11Uniform2f(gl11Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl11Uniform3f(gl11Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl11Uniform4f(gl11Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl11Uniform1i(gl11Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl11Uniform2i(gl11Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl11Uniform3i(gl11Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl11Uniform4i(gl11Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl11Uniform1ui(gl11Context* glc, GLint location, GLuint v0) {
    return glc->fnUniform1ui(location, v0);
}

void gl11Uniform2ui(gl11Context* glc, GLint location, GLuint v0, GLuint v1) {
    return glc->fnUniform2ui(location, v0, v1);
}

void gl11Uniform3ui(gl11Context* glc, GLint location, GLuint v0, GLuint v1, GLuint v2) {
    return glc->fnUniform3ui(location, v0, v1, v2);
}

void gl11Uniform4ui(gl11Context* glc, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {
    return glc->fnUniform4ui(location, v0, v1, v2, v3);
}

void gl11Uniform1fv(gl11Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl11Uniform2fv(gl11Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl11Uniform3fv(gl11Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl11Uniform4fv(gl11Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl11Uniform1iv(gl11Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl11Uniform2iv(gl11Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl11Uniform3iv(gl11Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl11Uniform4iv(gl11Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl11Uniform1uiv(gl11Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform1uiv(location, count, value);
}

void gl11Uniform2uiv(gl11Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform2uiv(location, count, value);
}

void gl11Uniform3uiv(gl11Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform3uiv(location, count, value);
}

void gl11Uniform4uiv(gl11Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform4uiv(location, count, value);
}

void gl11UseProgram(gl11Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl11ValidateProgram(gl11Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl11VertexAttribPointer(gl11Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl11VertexAttrib1f(gl11Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl11VertexAttrib1s(gl11Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl11VertexAttrib1d(gl11Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl11VertexAttrib2f(gl11Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl11VertexAttrib2s(gl11Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl11VertexAttrib2d(gl11Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl11VertexAttrib3f(gl11Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl11VertexAttrib3s(gl11Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl11VertexAttrib3d(gl11Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl11VertexAttrib4f(gl11Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl11VertexAttrib4s(gl11Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl11VertexAttrib4d(gl11Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl11VertexAttrib4Nuv(gl11Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl11VertexAttrib1fv(gl11Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl11VertexAttrib1sv(gl11Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl11VertexAttrib1dv(gl11Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl11VertexAttrib2fv(gl11Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl11VertexAttrib2sv(gl11Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl11VertexAttrib2dv(gl11Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl11VertexAttrib3fv(gl11Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl11VertexAttrib3sv(gl11Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl11VertexAttrib3dv(gl11Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl11VertexAttrib4fv(gl11Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl11VertexAttrib4sv(gl11Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl11VertexAttrib4dv(gl11Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl11VertexAttrib4iv(gl11Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl11VertexAttrib4bv(gl11Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl11VertexAttrib4ubv(gl11Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl11VertexAttrib4usv(gl11Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl11VertexAttrib4uiv(gl11Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl11VertexAttrib4Nbv(gl11Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl11VertexAttrib4Nsv(gl11Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl11VertexAttrib4Niv(gl11Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl11VertexAttrib4Nubv(gl11Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl11VertexAttrib4Nusv(gl11Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl11VertexAttrib4Nuiv(gl11Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl11UniformMatrix2fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl11UniformMatrix3fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl11UniformMatrix4fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl11UniformMatrix2x3fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl11UniformMatrix3x2fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl11UniformMatrix2x4fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl11UniformMatrix4x2fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl11UniformMatrix3x4fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl11UniformMatrix4x3fv(gl11Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl11Context* gl11NewContext() {
    gl11Context* glc = calloc(1, sizeof(gl11Context));

    // Preload all procedures
    glc->fnAccum = (gl11PAccum)gl11LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl11PAlphaFunc)gl11LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl11PBegin)gl11LibGetProcAddress("glBegin");
    glc->fnEnd = (gl11PEnd)gl11LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl11PBitmap)gl11LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl11PBlendFunc)gl11LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl11PCallList)gl11LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl11PCallLists)gl11LibGetProcAddress("glCallLists");
    glc->fnClear = (gl11PClear)gl11LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl11PClearAccum)gl11LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl11PClearColor)gl11LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl11PClearDepth)gl11LibGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl11PClearIndex)gl11LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl11PClearStencil)gl11LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl11PClipPlane)gl11LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl11PColor3b)gl11LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl11PColor3d)gl11LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl11PColor3f)gl11LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl11PColor3i)gl11LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl11PColor3s)gl11LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl11PColor3ub)gl11LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl11PColor3ui)gl11LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl11PColor3us)gl11LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl11PColor4b)gl11LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl11PColor4d)gl11LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl11PColor4f)gl11LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl11PColor4i)gl11LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl11PColor4s)gl11LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl11PColor4ub)gl11LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl11PColor4ui)gl11LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl11PColor4us)gl11LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl11PColor3bv)gl11LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl11PColor3dv)gl11LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl11PColor3fv)gl11LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl11PColor3iv)gl11LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl11PColor3sv)gl11LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl11PColor3ubv)gl11LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl11PColor3uiv)gl11LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl11PColor3usv)gl11LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl11PColor4bv)gl11LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl11PColor4dv)gl11LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl11PColor4fv)gl11LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl11PColor4iv)gl11LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl11PColor4sv)gl11LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl11PColor4ubv)gl11LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl11PColor4uiv)gl11LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl11PColor4usv)gl11LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl11PColorMask)gl11LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl11PColorMaterial)gl11LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl11PCopyPixels)gl11LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl11PCullFace)gl11LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl11PDeleteLists)gl11LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl11PDepthFunc)gl11LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl11PDepthMask)gl11LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl11PDepthRange)gl11LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl11PEnable)gl11LibGetProcAddress("glEnable");
    glc->fnDisable = (gl11PDisable)gl11LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl11PDrawBuffer)gl11LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl11PDrawPixels)gl11LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl11PEdgeFlag)gl11LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl11PEdgeFlagv)gl11LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl11PEdgeFlagPointer)gl11LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl11PEvalCoord1d)gl11LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl11PEvalCoord1f)gl11LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl11PEvalCoord2d)gl11LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl11PEvalCoord2f)gl11LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl11PEvalCoord1dv)gl11LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl11PEvalCoord1fv)gl11LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl11PEvalCoord2dv)gl11LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl11PEvalCoord2fv)gl11LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl11PEvalMesh1)gl11LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl11PEvalMesh2)gl11LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl11PEvalPoint1)gl11LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl11PEvalPoint2)gl11LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl11PFeedbackBuffer)gl11LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl11PFinish)gl11LibGetProcAddress("glFinish");
    glc->fnFlush = (gl11PFlush)gl11LibGetProcAddress("glFlush");
    glc->fnFogf = (gl11PFogf)gl11LibGetProcAddress("glFogf");
    glc->fnFogi = (gl11PFogi)gl11LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl11PFogfv)gl11LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl11PFogiv)gl11LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl11PFrontFace)gl11LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl11PFrustum)gl11LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl11PGenLists)gl11LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl11PGetBooleanv)gl11LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl11PGetDoublev)gl11LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl11PGetFloatv)gl11LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl11PGetIntegerv)gl11LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl11PGetClipPlane)gl11LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl11PGetError)gl11LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl11PGetLightfv)gl11LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl11PGetLightiv)gl11LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl11PGetMapdv)gl11LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl11PGetMapfv)gl11LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl11PGetMapiv)gl11LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl11PGetMaterialfv)gl11LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl11PGetMaterialiv)gl11LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl11PGetPixelMapfv)gl11LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl11PGetPixelMapuiv)gl11LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl11PGetPixelMapusv)gl11LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl11PGetPolygonStipple)gl11LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl11PGetString)gl11LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl11PGetTexEnvfv)gl11LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl11PGetTexEnviv)gl11LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl11PGetTexGendv)gl11LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl11PGetTexGenfv)gl11LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl11PGetTexGeniv)gl11LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl11PGetTexImage)gl11LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl11PGetTexLevelParameterfv)gl11LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl11PGetTexLevelParameteriv)gl11LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl11PGetTexParameterfv)gl11LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl11PGetTexParameteriv)gl11LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl11PHint)gl11LibGetProcAddress("glHint");
    glc->fnIndexd = (gl11PIndexd)gl11LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl11PIndexf)gl11LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl11PIndexi)gl11LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl11PIndexs)gl11LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl11PIndexdv)gl11LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl11PIndexfv)gl11LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl11PIndexiv)gl11LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl11PIndexsv)gl11LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl11PIndexMask)gl11LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl11PIndexPointer)gl11LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl11PInitNames)gl11LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl11PIsEnabled)gl11LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl11PIsList)gl11LibGetProcAddress("glIsList");
    glc->fnLightf = (gl11PLightf)gl11LibGetProcAddress("glLightf");
    glc->fnLighti = (gl11PLighti)gl11LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl11PLightfv)gl11LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl11PLightiv)gl11LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl11PLightModelf)gl11LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl11PLightModeli)gl11LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl11PLightModelfv)gl11LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl11PLightModeliv)gl11LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl11PLineStipple)gl11LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl11PLineWidth)gl11LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl11PListBase)gl11LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl11PLoadIdentity)gl11LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl11PLoadMatrixd)gl11LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl11PLoadMatrixf)gl11LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl11PLoadName)gl11LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl11PLogicOp)gl11LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl11PMap1d)gl11LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl11PMap1f)gl11LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl11PMap2d)gl11LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl11PMap2f)gl11LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl11PMapGrid1d)gl11LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl11PMapGrid1f)gl11LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl11PMapGrid2d)gl11LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl11PMapGrid2f)gl11LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl11PMaterialf)gl11LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl11PMateriali)gl11LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl11PMaterialfv)gl11LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl11PMaterialiv)gl11LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl11PMatrixMode)gl11LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl11PMultMatrixd)gl11LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl11PMultMatrixf)gl11LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl11PNewList)gl11LibGetProcAddress("glNewList");
    glc->fnEndList = (gl11PEndList)gl11LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl11PNormal3b)gl11LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl11PNormal3d)gl11LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl11PNormal3f)gl11LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl11PNormal3i)gl11LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl11PNormal3s)gl11LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl11PNormal3bv)gl11LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl11PNormal3dv)gl11LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl11PNormal3fv)gl11LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl11PNormal3iv)gl11LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl11PNormal3sv)gl11LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl11POrtho)gl11LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl11PPassThrough)gl11LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl11PPixelMapfv)gl11LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl11PPixelMapuiv)gl11LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl11PPixelMapusv)gl11LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl11PPixelStoref)gl11LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl11PPixelStorei)gl11LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl11PPixelTransferf)gl11LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl11PPixelTransferi)gl11LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl11PPixelZoom)gl11LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl11PPointSize)gl11LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl11PPolygonMode)gl11LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl11PPolygonStipple)gl11LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl11PPushAttrib)gl11LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl11PPopAttrib)gl11LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl11PPushMatrix)gl11LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl11PPopMatrix)gl11LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl11PPushName)gl11LibGetProcAddress("glPushName");
    glc->fnPopName = (gl11PPopName)gl11LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl11PRasterPos2d)gl11LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl11PRasterPos2f)gl11LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl11PRasterPos2i)gl11LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl11PRasterPos2s)gl11LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl11PRasterPos3d)gl11LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl11PRasterPos3f)gl11LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl11PRasterPos3i)gl11LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl11PRasterPos3s)gl11LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl11PRasterPos4d)gl11LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl11PRasterPos4f)gl11LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl11PRasterPos4i)gl11LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl11PRasterPos4s)gl11LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl11PRasterPos2dv)gl11LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl11PRasterPos2fv)gl11LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl11PRasterPos2iv)gl11LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl11PRasterPos2sv)gl11LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl11PRasterPos3dv)gl11LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl11PRasterPos3fv)gl11LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl11PRasterPos3iv)gl11LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl11PRasterPos3sv)gl11LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl11PRasterPos4dv)gl11LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl11PRasterPos4fv)gl11LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl11PRasterPos4iv)gl11LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl11PRasterPos4sv)gl11LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl11PReadBuffer)gl11LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl11PReadPixels)gl11LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl11PRectd)gl11LibGetProcAddress("glRectd");
    glc->fnRectf = (gl11PRectf)gl11LibGetProcAddress("glRectf");
    glc->fnRecti = (gl11PRecti)gl11LibGetProcAddress("glRecti");
    glc->fnRects = (gl11PRects)gl11LibGetProcAddress("glRects");
    glc->fnRectdv = (gl11PRectdv)gl11LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl11PRectfv)gl11LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl11PRectiv)gl11LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl11PRectsv)gl11LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl11PRenderMode)gl11LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl11PRotated)gl11LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl11PRotatef)gl11LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl11PScaled)gl11LibGetProcAddress("glScaled");
    glc->fnScalef = (gl11PScalef)gl11LibGetProcAddress("glScalef");
    glc->fnScissor = (gl11PScissor)gl11LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl11PSelectBuffer)gl11LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl11PShadeModel)gl11LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl11PStencilFunc)gl11LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl11PStencilMask)gl11LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl11PStencilOp)gl11LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl11PTexCoord1d)gl11LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl11PTexCoord1f)gl11LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl11PTexCoord1i)gl11LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl11PTexCoord1s)gl11LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl11PTexCoord2d)gl11LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl11PTexCoord2f)gl11LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl11PTexCoord2i)gl11LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl11PTexCoord2s)gl11LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl11PTexCoord3d)gl11LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl11PTexCoord3f)gl11LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl11PTexCoord3i)gl11LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl11PTexCoord3s)gl11LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl11PTexCoord4d)gl11LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl11PTexCoord4f)gl11LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl11PTexCoord4i)gl11LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl11PTexCoord4s)gl11LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl11PTexCoord1dv)gl11LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl11PTexCoord1fv)gl11LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl11PTexCoord1iv)gl11LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl11PTexCoord1sv)gl11LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl11PTexCoord2dv)gl11LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl11PTexCoord2fv)gl11LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl11PTexCoord2iv)gl11LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl11PTexCoord2sv)gl11LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl11PTexCoord3dv)gl11LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl11PTexCoord3fv)gl11LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl11PTexCoord3iv)gl11LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl11PTexCoord3sv)gl11LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl11PTexCoord4dv)gl11LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl11PTexCoord4fv)gl11LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl11PTexCoord4iv)gl11LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl11PTexCoord4sv)gl11LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl11PTexEnvf)gl11LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl11PTexEnvi)gl11LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl11PTexEnvfv)gl11LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl11PTexEnviv)gl11LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl11PTexGend)gl11LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl11PTexGenf)gl11LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl11PTexGeni)gl11LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl11PTexGendv)gl11LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl11PTexGenfv)gl11LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl11PTexGeniv)gl11LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl11PTexImage1D)gl11LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl11PTexImage2D)gl11LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl11PTexParameterf)gl11LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl11PTexParameteri)gl11LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl11PTexParameterfv)gl11LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl11PTexParameteriv)gl11LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl11PTranslated)gl11LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl11PTranslatef)gl11LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl11PVertex2s)gl11LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl11PVertex2i)gl11LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl11PVertex2f)gl11LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl11PVertex2d)gl11LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl11PVertex3s)gl11LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl11PVertex3i)gl11LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl11PVertex3f)gl11LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl11PVertex3d)gl11LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl11PVertex4s)gl11LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl11PVertex4i)gl11LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl11PVertex4f)gl11LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl11PVertex4d)gl11LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl11PViewport)gl11LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl11PGetConvolutionParameterfv)gl11LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl11PGetConvolutionParameteriv)gl11LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl11PAreTexturesResident)gl11LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl11PArrayElement)gl11LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl11PDrawArrays)gl11LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl11PDrawElements)gl11LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl11PGetPointerv)gl11LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl11PPolygonOffset)gl11LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl11PCopyTexImage1D)gl11LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl11PCopyTexImage2D)gl11LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl11PCopyTexSubImage1D)gl11LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl11PCopyTexSubImage2D)gl11LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl11PBindTexture)gl11LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl11PDeleteTextures)gl11LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl11PGenTextures)gl11LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl11PIsTexture)gl11LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl11PColorPointer)gl11LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl11PEnableClientState)gl11LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl11PDisableClientState)gl11LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl11PIndexub)gl11LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl11PIndexubv)gl11LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl11PInterleavedArrays)gl11LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl11PNormalPointer)gl11LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl11PPushClientAttrib)gl11LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl11PPrioritizeTextures)gl11LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl11PPopClientAttrib)gl11LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl11PTexCoordPointer)gl11LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl11PTexSubImage1D)gl11LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl11PTexSubImage2D)gl11LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl11PVertexPointer)gl11LibGetProcAddress("glVertexPointer");
    glc->fnColorTable = (gl11PColorTable)gl11GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl11PColorTableParameterfv)gl11GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl11PColorTableParameteriv)gl11GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl11PColorSubTable)gl11GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl11PConvolutionFilter1D)gl11GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl11PConvolutionFilter2D)gl11GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl11PConvolutionParameterf)gl11GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl11PConvolutionParameteri)gl11GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl11PCopyColorTable)gl11GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl11PCopyColorSubTable)gl11GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl11PCopyConvolutionFilter1D)gl11GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl11PCopyConvolutionFilter2D)gl11GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl11PGetColorTable)gl11GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl11PGetColorTableParameterfv)gl11GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl11PGetColorTableParameteriv)gl11GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl11PGetConvolutionFilter)gl11GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl11PGetHistogram)gl11GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl11PGetHistogramParameterfv)gl11GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl11PGetHistogramParameteriv)gl11GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl11PGetSeparableFilter)gl11GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl11PHistogram)gl11GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl11PMinmax)gl11GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl11PMultiTexCoord1s)gl11GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl11PMultiTexCoord1i)gl11GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl11PMultiTexCoord1f)gl11GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl11PMultiTexCoord1d)gl11GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl11PMultiTexCoord2s)gl11GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl11PMultiTexCoord2i)gl11GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl11PMultiTexCoord2f)gl11GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl11PMultiTexCoord2d)gl11GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl11PMultiTexCoord3s)gl11GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl11PMultiTexCoord3i)gl11GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl11PMultiTexCoord3f)gl11GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl11PMultiTexCoord3d)gl11GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl11PMultiTexCoord4s)gl11GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl11PMultiTexCoord4i)gl11GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl11PMultiTexCoord4f)gl11GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl11PMultiTexCoord4d)gl11GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl11PMultiTexCoord1sv)gl11GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl11PMultiTexCoord1iv)gl11GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl11PMultiTexCoord1fv)gl11GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl11PMultiTexCoord1dv)gl11GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl11PMultiTexCoord2sv)gl11GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl11PMultiTexCoord2iv)gl11GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl11PMultiTexCoord2fv)gl11GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl11PMultiTexCoord2dv)gl11GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl11PMultiTexCoord3sv)gl11GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl11PMultiTexCoord3iv)gl11GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl11PMultiTexCoord3fv)gl11GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl11PMultiTexCoord3dv)gl11GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl11PMultiTexCoord4sv)gl11GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl11PMultiTexCoord4iv)gl11GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl11PMultiTexCoord4fv)gl11GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl11PMultiTexCoord4dv)gl11GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl11PResetHistogram)gl11GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl11PResetMinmax)gl11GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl11PSeparableFilter2D)gl11GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl11PBlendColor)gl11GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl11PBlendEquation)gl11GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl11PCopyTexSubImage3D)gl11GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl11PDrawRangeElements)gl11GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl11PTexImage3D)gl11GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl11PTexSubImage3D)gl11GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl11PActiveTexture)gl11GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl11PClientActiveTexture)gl11GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl11PCompressedTexImage1D)gl11GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl11PCompressedTexImage2D)gl11GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl11PCompressedTexImage3D)gl11GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl11PCompressedTexSubImage1D)gl11GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl11PCompressedTexSubImage2D)gl11GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl11PCompressedTexSubImage3D)gl11GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl11PGetCompressedTexImage)gl11GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl11PLoadTransposeMatrixd)gl11GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl11PLoadTransposeMatrixf)gl11GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl11PMultTransposeMatrixd)gl11GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl11PMultTransposeMatrixf)gl11GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl11PSampleCoverage)gl11GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl11PBlendFuncSeparate)gl11GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl11PFogCoordPointer)gl11GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl11PFogCoordd)gl11GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl11PFogCoordf)gl11GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl11PFogCoorddv)gl11GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl11PFogCoordfv)gl11GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl11PMultiDrawArrays)gl11GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl11PMultiDrawElements)gl11GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl11PPointParameterf)gl11GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl11PPointParameteri)gl11GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl11PSecondaryColor3b)gl11GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl11PSecondaryColor3s)gl11GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl11PSecondaryColor3i)gl11GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl11PSecondaryColor3f)gl11GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl11PSecondaryColor3d)gl11GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl11PSecondaryColor3ub)gl11GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl11PSecondaryColor3us)gl11GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl11PSecondaryColor3ui)gl11GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl11PSecondaryColor3bv)gl11GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl11PSecondaryColor3sv)gl11GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl11PSecondaryColor3iv)gl11GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl11PSecondaryColor3fv)gl11GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl11PSecondaryColor3dv)gl11GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl11PSecondaryColor3ubv)gl11GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl11PSecondaryColor3usv)gl11GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl11PSecondaryColor3uiv)gl11GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl11PSecondaryColorPointer)gl11GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl11PWindowPos2s)gl11GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl11PWindowPos2i)gl11GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl11PWindowPos2f)gl11GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl11PWindowPos2d)gl11GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl11PWindowPos3s)gl11GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl11PWindowPos3i)gl11GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl11PWindowPos3f)gl11GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl11PWindowPos3d)gl11GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl11PWindowPos2sv)gl11GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl11PWindowPos2iv)gl11GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl11PWindowPos2fv)gl11GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl11PWindowPos2dv)gl11GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl11PWindowPos3sv)gl11GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl11PWindowPos3iv)gl11GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl11PWindowPos3fv)gl11GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl11PWindowPos3dv)gl11GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl11PBeginQuery)gl11GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl11PBindBuffer)gl11GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl11PBufferData)gl11GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl11PBufferSubData)gl11GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl11PDeleteBuffers)gl11GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl11PDeleteQueries)gl11GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl11PGenBuffers)gl11GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl11PGenQueries)gl11GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl11PGetBufferParameteriv)gl11GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl11PGetBufferPointerv)gl11GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl11PGetBufferSubData)gl11GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl11PGetQueryObjectiv)gl11GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl11PGetQueryObjectuiv)gl11GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl11PGetQueryiv)gl11GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl11PIsBuffer)gl11GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl11PIsQuery)gl11GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl11PMapBuffer)gl11GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl11PUnmapBuffer)gl11GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl11PAttachShader)gl11GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl11PBindAttribLocation)gl11GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl11PBlendEquationSeperate)gl11GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl11PCompileShader)gl11GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl11PCreateProgram)gl11GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl11PCreateShader)gl11GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl11PDeleteProgram)gl11GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl11PDeleteShader)gl11GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl11PDetachShader)gl11GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl11PEnableVertexAttribArray)gl11GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl11PDisableVertexAttribArray)gl11GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl11PDrawBuffers)gl11GLGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl11PGetActiveAttrib)gl11GLGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl11PGetActiveUniform)gl11GLGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl11PGetAttachedShaders)gl11GLGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl11PGetAttribLocation)gl11GLGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl11PGetProgramiv)gl11GLGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl11PGetProgramInfoLog)gl11GLGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl11PGetShaderiv)gl11GLGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl11PGetShaderInfoLog)gl11GLGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl11PGetShaderSource)gl11GLGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl11PGetUniformfv)gl11GLGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl11PGetUniformiv)gl11GLGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl11PGetUniformLocation)gl11GLGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl11PGetVertexAttribdv)gl11GLGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl11PGetVertexAttribfv)gl11GLGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl11PGetVertexAttribiv)gl11GLGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl11PGetVertexAttribPointerv)gl11GLGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl11PIsProgram)gl11GLGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl11PIsShader)gl11GLGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl11PLinkProgram)gl11GLGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl11PShaderSource)gl11GLGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl11PStencilFuncSeparate)gl11GLGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl11PStencilMaskSeparate)gl11GLGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl11PStencilOpSeparate)gl11GLGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl11PUniform1f)gl11GLGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl11PUniform2f)gl11GLGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl11PUniform3f)gl11GLGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl11PUniform4f)gl11GLGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl11PUniform1i)gl11GLGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl11PUniform2i)gl11GLGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl11PUniform3i)gl11GLGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl11PUniform4i)gl11GLGetProcAddress("glUniform4i");
    glc->fnUniform1ui = (gl11PUniform1ui)gl11GLGetProcAddress("glUniform1ui");
    glc->fnUniform2ui = (gl11PUniform2ui)gl11GLGetProcAddress("glUniform2ui");
    glc->fnUniform3ui = (gl11PUniform3ui)gl11GLGetProcAddress("glUniform3ui");
    glc->fnUniform4ui = (gl11PUniform4ui)gl11GLGetProcAddress("glUniform4ui");
    glc->fnUniform1fv = (gl11PUniform1fv)gl11GLGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl11PUniform2fv)gl11GLGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl11PUniform3fv)gl11GLGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl11PUniform4fv)gl11GLGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl11PUniform1iv)gl11GLGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl11PUniform2iv)gl11GLGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl11PUniform3iv)gl11GLGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl11PUniform4iv)gl11GLGetProcAddress("glUniform4iv");
    glc->fnUniform1uiv = (gl11PUniform1uiv)gl11GLGetProcAddress("glUniform1uiv");
    glc->fnUniform2uiv = (gl11PUniform2uiv)gl11GLGetProcAddress("glUniform2uiv");
    glc->fnUniform3uiv = (gl11PUniform3uiv)gl11GLGetProcAddress("glUniform3uiv");
    glc->fnUniform4uiv = (gl11PUniform4uiv)gl11GLGetProcAddress("glUniform4uiv");
    glc->fnUseProgram = (gl11PUseProgram)gl11GLGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl11PValidateProgram)gl11GLGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl11PVertexAttribPointer)gl11GLGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl11PVertexAttrib1f)gl11GLGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl11PVertexAttrib1s)gl11GLGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl11PVertexAttrib1d)gl11GLGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl11PVertexAttrib2f)gl11GLGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl11PVertexAttrib2s)gl11GLGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl11PVertexAttrib2d)gl11GLGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl11PVertexAttrib3f)gl11GLGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl11PVertexAttrib3s)gl11GLGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl11PVertexAttrib3d)gl11GLGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl11PVertexAttrib4f)gl11GLGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl11PVertexAttrib4s)gl11GLGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl11PVertexAttrib4d)gl11GLGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl11PVertexAttrib4Nuv)gl11GLGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl11PVertexAttrib1fv)gl11GLGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl11PVertexAttrib1sv)gl11GLGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl11PVertexAttrib1dv)gl11GLGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl11PVertexAttrib2fv)gl11GLGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl11PVertexAttrib2sv)gl11GLGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl11PVertexAttrib2dv)gl11GLGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl11PVertexAttrib3fv)gl11GLGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl11PVertexAttrib3sv)gl11GLGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl11PVertexAttrib3dv)gl11GLGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl11PVertexAttrib4fv)gl11GLGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl11PVertexAttrib4sv)gl11GLGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl11PVertexAttrib4dv)gl11GLGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl11PVertexAttrib4iv)gl11GLGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl11PVertexAttrib4bv)gl11GLGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl11PVertexAttrib4ubv)gl11GLGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl11PVertexAttrib4usv)gl11GLGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl11PVertexAttrib4uiv)gl11GLGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl11PVertexAttrib4Nbv)gl11GLGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl11PVertexAttrib4Nsv)gl11GLGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl11PVertexAttrib4Niv)gl11GLGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl11PVertexAttrib4Nubv)gl11GLGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl11PVertexAttrib4Nusv)gl11GLGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl11PVertexAttrib4Nuiv)gl11GLGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl11PUniformMatrix2fv)gl11GLGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl11PUniformMatrix3fv)gl11GLGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl11PUniformMatrix4fv)gl11GLGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl11PUniformMatrix2x3fv)gl11GLGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl11PUniformMatrix3x2fv)gl11GLGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl11PUniformMatrix2x4fv)gl11GLGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl11PUniformMatrix4x2fv)gl11GLGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl11PUniformMatrix3x4fv)gl11GLGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl11PUniformMatrix4x3fv)gl11GLGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

