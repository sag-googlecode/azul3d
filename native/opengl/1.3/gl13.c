// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl13.h"

#ifdef _WIN32
HMODULE gl13OpenGL32;

void* gl13LibGetProcAddress(char* name) {
	if(gl13OpenGL32 == NULL) {
		gl13OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl13OpenGL32, TEXT(name));
}

void* gl13GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl13Accum(gl13Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl13AlphaFunc(gl13Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl13Begin(gl13Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl13End(gl13Context* glc) {
    return glc->fnEnd();
}

void gl13Bitmap(gl13Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl13BlendFunc(gl13Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl13CallList(gl13Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl13CallLists(gl13Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl13Clear(gl13Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl13ClearAccum(gl13Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl13ClearColor(gl13Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl13ClearDepth(gl13Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl13ClearIndex(gl13Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl13ClearStencil(gl13Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl13ClipPlane(gl13Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl13Color3b(gl13Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl13Color3d(gl13Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl13Color3f(gl13Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl13Color3i(gl13Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl13Color3s(gl13Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl13Color3ub(gl13Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl13Color3ui(gl13Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl13Color3us(gl13Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl13Color4b(gl13Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl13Color4d(gl13Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl13Color4f(gl13Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl13Color4i(gl13Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl13Color4s(gl13Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl13Color4ub(gl13Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl13Color4ui(gl13Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl13Color4us(gl13Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl13Color3bv(gl13Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl13Color3dv(gl13Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl13Color3fv(gl13Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl13Color3iv(gl13Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl13Color3sv(gl13Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl13Color3ubv(gl13Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl13Color3uiv(gl13Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl13Color3usv(gl13Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl13Color4bv(gl13Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl13Color4dv(gl13Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl13Color4fv(gl13Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl13Color4iv(gl13Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl13Color4sv(gl13Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl13Color4ubv(gl13Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl13Color4uiv(gl13Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl13Color4usv(gl13Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl13ColorMask(gl13Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl13ColorMaterial(gl13Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl13CopyPixels(gl13Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl13CullFace(gl13Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl13DeleteLists(gl13Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl13DepthFunc(gl13Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl13DepthMask(gl13Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl13DepthRange(gl13Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl13Enable(gl13Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl13Disable(gl13Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl13DrawBuffer(gl13Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl13DrawPixels(gl13Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl13EdgeFlag(gl13Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl13EdgeFlagv(gl13Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl13EdgeFlagPointer(gl13Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl13EvalCoord1d(gl13Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl13EvalCoord1f(gl13Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl13EvalCoord2d(gl13Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl13EvalCoord2f(gl13Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl13EvalCoord1dv(gl13Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl13EvalCoord1fv(gl13Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl13EvalCoord2dv(gl13Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl13EvalCoord2fv(gl13Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl13EvalMesh1(gl13Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl13EvalMesh2(gl13Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl13EvalPoint1(gl13Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl13EvalPoint2(gl13Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl13FeedbackBuffer(gl13Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl13Finish(gl13Context* glc) {
    return glc->fnFinish();
}

void gl13Flush(gl13Context* glc) {
    return glc->fnFlush();
}

void gl13Fogf(gl13Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl13Fogi(gl13Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl13Fogfv(gl13Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl13Fogiv(gl13Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl13FrontFace(gl13Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl13Frustum(gl13Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl13GenLists(gl13Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl13GetBooleanv(gl13Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl13GetDoublev(gl13Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl13GetFloatv(gl13Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl13GetIntegerv(gl13Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl13GetClipPlane(gl13Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl13GetError(gl13Context* glc) {
    return glc->fnGetError();
}

void gl13GetLightfv(gl13Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl13GetLightiv(gl13Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl13GetMapdv(gl13Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl13GetMapfv(gl13Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl13GetMapiv(gl13Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl13GetMaterialfv(gl13Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl13GetMaterialiv(gl13Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl13GetPixelMapfv(gl13Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl13GetPixelMapuiv(gl13Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl13GetPixelMapusv(gl13Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl13GetPolygonStipple(gl13Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl13GetString(gl13Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl13GetTexEnvfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl13GetTexEnviv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl13GetTexGendv(gl13Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl13GetTexGenfv(gl13Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl13GetTexGeniv(gl13Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl13GetTexImage(gl13Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl13GetTexLevelParameterfv(gl13Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl13GetTexLevelParameteriv(gl13Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl13GetTexParameterfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl13GetTexParameteriv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl13Hint(gl13Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl13Indexd(gl13Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl13Indexf(gl13Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl13Indexi(gl13Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl13Indexs(gl13Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl13Indexdv(gl13Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl13Indexfv(gl13Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl13Indexiv(gl13Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl13Indexsv(gl13Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl13IndexMask(gl13Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl13IndexPointer(gl13Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl13InitNames(gl13Context* glc) {
    return glc->fnInitNames();
}

void gl13IsEnabled(gl13Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl13IsList(gl13Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl13Lightf(gl13Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl13Lighti(gl13Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl13Lightfv(gl13Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl13Lightiv(gl13Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl13LightModelf(gl13Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl13LightModeli(gl13Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl13LightModelfv(gl13Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl13LightModeliv(gl13Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl13LineStipple(gl13Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl13LineWidth(gl13Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl13ListBase(gl13Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl13LoadIdentity(gl13Context* glc) {
    return glc->fnLoadIdentity();
}

void gl13LoadMatrixd(gl13Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl13LoadMatrixf(gl13Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl13LoadName(gl13Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl13LogicOp(gl13Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl13Map1d(gl13Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl13Map1f(gl13Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl13Map2d(gl13Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl13Map2f(gl13Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl13MapGrid1d(gl13Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl13MapGrid1f(gl13Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl13MapGrid2d(gl13Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl13MapGrid2f(gl13Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl13Materialf(gl13Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl13Materiali(gl13Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl13Materialfv(gl13Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl13Materialiv(gl13Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl13MatrixMode(gl13Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl13MultMatrixd(gl13Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl13MultMatrixf(gl13Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl13NewList(gl13Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl13EndList(gl13Context* glc) {
    return glc->fnEndList();
}

void gl13Normal3b(gl13Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl13Normal3d(gl13Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl13Normal3f(gl13Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl13Normal3i(gl13Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl13Normal3s(gl13Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl13Normal3bv(gl13Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl13Normal3dv(gl13Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl13Normal3fv(gl13Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl13Normal3iv(gl13Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl13Normal3sv(gl13Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl13Ortho(gl13Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl13PassThrough(gl13Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl13PixelMapfv(gl13Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl13PixelMapuiv(gl13Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl13PixelMapusv(gl13Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl13PixelStoref(gl13Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl13PixelStorei(gl13Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl13PixelTransferf(gl13Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl13PixelTransferi(gl13Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl13PixelZoom(gl13Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl13PointSize(gl13Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl13PolygonMode(gl13Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl13PolygonStipple(gl13Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl13PushAttrib(gl13Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl13PopAttrib(gl13Context* glc) {
    return glc->fnPopAttrib();
}

void gl13PushMatrix(gl13Context* glc) {
    return glc->fnPushMatrix();
}

void gl13PopMatrix(gl13Context* glc) {
    return glc->fnPopMatrix();
}

void gl13PushName(gl13Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl13PopName(gl13Context* glc) {
    return glc->fnPopName();
}

void gl13RasterPos2d(gl13Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl13RasterPos2f(gl13Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl13RasterPos2i(gl13Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl13RasterPos2s(gl13Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl13RasterPos3d(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl13RasterPos3f(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl13RasterPos3i(gl13Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl13RasterPos3s(gl13Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl13RasterPos4d(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl13RasterPos4f(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl13RasterPos4i(gl13Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl13RasterPos4s(gl13Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl13RasterPos2dv(gl13Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl13RasterPos2fv(gl13Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl13RasterPos2iv(gl13Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl13RasterPos2sv(gl13Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl13RasterPos3dv(gl13Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl13RasterPos3fv(gl13Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl13RasterPos3iv(gl13Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl13RasterPos3sv(gl13Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl13RasterPos4dv(gl13Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl13RasterPos4fv(gl13Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl13RasterPos4iv(gl13Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl13RasterPos4sv(gl13Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl13ReadBuffer(gl13Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl13ReadPixels(gl13Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl13Rectd(gl13Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl13Rectf(gl13Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl13Recti(gl13Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl13Rects(gl13Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl13Rectdv(gl13Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl13Rectfv(gl13Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl13Rectiv(gl13Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl13Rectsv(gl13Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl13RenderMode(gl13Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl13Rotated(gl13Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl13Rotatef(gl13Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl13Scaled(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl13Scalef(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl13Scissor(gl13Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl13SelectBuffer(gl13Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl13ShadeModel(gl13Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl13StencilFunc(gl13Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl13StencilMask(gl13Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl13StencilOp(gl13Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl13TexCoord1d(gl13Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl13TexCoord1f(gl13Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl13TexCoord1i(gl13Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl13TexCoord1s(gl13Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl13TexCoord2d(gl13Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl13TexCoord2f(gl13Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl13TexCoord2i(gl13Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl13TexCoord2s(gl13Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl13TexCoord3d(gl13Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl13TexCoord3f(gl13Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl13TexCoord3i(gl13Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl13TexCoord3s(gl13Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl13TexCoord4d(gl13Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl13TexCoord4f(gl13Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl13TexCoord4i(gl13Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl13TexCoord4s(gl13Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl13TexCoord1dv(gl13Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl13TexCoord1fv(gl13Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl13TexCoord1iv(gl13Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl13TexCoord1sv(gl13Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl13TexCoord2dv(gl13Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl13TexCoord2fv(gl13Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl13TexCoord2iv(gl13Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl13TexCoord2sv(gl13Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl13TexCoord3dv(gl13Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl13TexCoord3fv(gl13Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl13TexCoord3iv(gl13Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl13TexCoord3sv(gl13Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl13TexCoord4dv(gl13Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl13TexCoord4fv(gl13Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl13TexCoord4iv(gl13Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl13TexCoord4sv(gl13Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl13TexEnvf(gl13Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl13TexEnvi(gl13Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl13TexEnvfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl13TexEnviv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl13TexGend(gl13Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl13TexGenf(gl13Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl13TexGeni(gl13Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl13TexGendv(gl13Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl13TexGenfv(gl13Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl13TexGeniv(gl13Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl13TexImage1D(gl13Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl13TexImage2D(gl13Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl13TexParameterf(gl13Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl13TexParameteri(gl13Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl13TexParameterfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl13TexParameteriv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl13Translated(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl13Translatef(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl13Vertex2s(gl13Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl13Vertex2i(gl13Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl13Vertex2f(gl13Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl13Vertex2d(gl13Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl13Vertex3s(gl13Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl13Vertex3i(gl13Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl13Vertex3f(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl13Vertex3d(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl13Vertex4s(gl13Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl13Vertex4i(gl13Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl13Vertex4f(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl13Vertex4d(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl13Viewport(gl13Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl13GetConvolutionParameterfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl13GetConvolutionParameteriv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl13AreTexturesResident(gl13Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl13ArrayElement(gl13Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl13DrawArrays(gl13Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl13DrawElements(gl13Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl13GetPointerv(gl13Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl13PolygonOffset(gl13Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl13CopyTexImage1D(gl13Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl13CopyTexImage2D(gl13Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl13CopyTexSubImage1D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl13CopyTexSubImage2D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl13BindTexture(gl13Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl13DeleteTextures(gl13Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl13GenTextures(gl13Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl13IsTexture(gl13Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl13ColorPointer(gl13Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl13EnableClientState(gl13Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl13DisableClientState(gl13Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl13Indexub(gl13Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl13Indexubv(gl13Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl13InterleavedArrays(gl13Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl13NormalPointer(gl13Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl13PushClientAttrib(gl13Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl13PrioritizeTextures(gl13Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl13PopClientAttrib(gl13Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl13TexCoordPointer(gl13Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl13TexSubImage1D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl13TexSubImage2D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl13VertexPointer(gl13Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl13ColorTable(gl13Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl13ColorTableParameterfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl13ColorTableParameteriv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl13ColorSubTable(gl13Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl13ConvolutionFilter1D(gl13Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl13ConvolutionFilter2D(gl13Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl13ConvolutionParameterf(gl13Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl13ConvolutionParameteri(gl13Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl13CopyColorTable(gl13Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl13CopyColorSubTable(gl13Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl13CopyConvolutionFilter1D(gl13Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl13CopyConvolutionFilter2D(gl13Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl13GetColorTable(gl13Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl13GetColorTableParameterfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl13GetColorTableParameteriv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl13GetConvolutionFilter(gl13Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl13GetHistogram(gl13Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl13GetHistogramParameterfv(gl13Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl13GetHistogramParameteriv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl13GetSeparableFilter(gl13Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl13Histogram(gl13Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl13Minmax(gl13Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl13MultiTexCoord1s(gl13Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl13MultiTexCoord1i(gl13Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl13MultiTexCoord1f(gl13Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl13MultiTexCoord1d(gl13Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl13MultiTexCoord2s(gl13Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl13MultiTexCoord2i(gl13Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl13MultiTexCoord2f(gl13Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl13MultiTexCoord2d(gl13Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl13MultiTexCoord3s(gl13Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl13MultiTexCoord3i(gl13Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl13MultiTexCoord3f(gl13Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl13MultiTexCoord3d(gl13Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl13MultiTexCoord4s(gl13Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl13MultiTexCoord4i(gl13Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl13MultiTexCoord4f(gl13Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl13MultiTexCoord4d(gl13Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl13MultiTexCoord1sv(gl13Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl13MultiTexCoord1iv(gl13Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl13MultiTexCoord1fv(gl13Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl13MultiTexCoord1dv(gl13Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl13MultiTexCoord2sv(gl13Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl13MultiTexCoord2iv(gl13Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl13MultiTexCoord2fv(gl13Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl13MultiTexCoord2dv(gl13Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl13MultiTexCoord3sv(gl13Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl13MultiTexCoord3iv(gl13Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl13MultiTexCoord3fv(gl13Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl13MultiTexCoord3dv(gl13Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl13MultiTexCoord4sv(gl13Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl13MultiTexCoord4iv(gl13Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl13MultiTexCoord4fv(gl13Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl13MultiTexCoord4dv(gl13Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl13ResetHistogram(gl13Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl13ResetMinmax(gl13Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl13SeparableFilter2D(gl13Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl13BlendColor(gl13Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl13BlendEquation(gl13Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl13CopyTexSubImage3D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl13DrawRangeElements(gl13Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl13TexImage3D(gl13Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl13TexSubImage3D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl13ActiveTexture(gl13Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl13ClientActiveTexture(gl13Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl13CompressedTexImage1D(gl13Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl13CompressedTexImage2D(gl13Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl13CompressedTexImage3D(gl13Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl13CompressedTexSubImage1D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl13CompressedTexSubImage2D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl13CompressedTexSubImage3D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl13GetCompressedTexImage(gl13Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl13LoadTransposeMatrixd(gl13Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl13LoadTransposeMatrixf(gl13Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl13MultTransposeMatrixd(gl13Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl13MultTransposeMatrixf(gl13Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl13SampleCoverage(gl13Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl13BlendFuncSeparate(gl13Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl13FogCoordPointer(gl13Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl13FogCoordd(gl13Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl13FogCoordf(gl13Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl13FogCoorddv(gl13Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl13FogCoordfv(gl13Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl13MultiDrawArrays(gl13Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl13MultiDrawElements(gl13Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl13PointParameterf(gl13Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl13PointParameteri(gl13Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl13SecondaryColor3b(gl13Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl13SecondaryColor3s(gl13Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl13SecondaryColor3i(gl13Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl13SecondaryColor3f(gl13Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl13SecondaryColor3d(gl13Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl13SecondaryColor3ub(gl13Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl13SecondaryColor3us(gl13Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl13SecondaryColor3ui(gl13Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl13SecondaryColor3bv(gl13Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl13SecondaryColor3sv(gl13Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl13SecondaryColor3iv(gl13Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl13SecondaryColor3fv(gl13Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl13SecondaryColor3dv(gl13Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl13SecondaryColor3ubv(gl13Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl13SecondaryColor3usv(gl13Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl13SecondaryColor3uiv(gl13Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl13SecondaryColorPointer(gl13Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl13WindowPos2s(gl13Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl13WindowPos2i(gl13Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl13WindowPos2f(gl13Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl13WindowPos2d(gl13Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl13WindowPos3s(gl13Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl13WindowPos3i(gl13Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl13WindowPos3f(gl13Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl13WindowPos3d(gl13Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl13WindowPos2sv(gl13Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl13WindowPos2iv(gl13Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl13WindowPos2fv(gl13Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl13WindowPos2dv(gl13Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl13WindowPos3sv(gl13Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl13WindowPos3iv(gl13Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl13WindowPos3fv(gl13Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl13WindowPos3dv(gl13Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl13BeginQuery(gl13Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl13BindBuffer(gl13Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl13BufferData(gl13Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl13BufferSubData(gl13Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl13DeleteBuffers(gl13Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl13DeleteQueries(gl13Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl13GenBuffers(gl13Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl13GenQueries(gl13Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl13GetBufferParameteriv(gl13Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl13GetBufferPointerv(gl13Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl13GetBufferSubData(gl13Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl13GetQueryObjectiv(gl13Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl13GetQueryObjectuiv(gl13Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl13GetQueryiv(gl13Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl13IsBuffer(gl13Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl13IsQuery(gl13Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl13MapBuffer(gl13Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl13UnmapBuffer(gl13Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl13AttachShader(gl13Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl13BindAttribLocation(gl13Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl13BlendEquationSeperate(gl13Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl13CompileShader(gl13Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl13CreateProgram(gl13Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl13CreateShader(gl13Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl13DeleteProgram(gl13Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl13DeleteShader(gl13Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl13DetachShader(gl13Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl13EnableVertexAttribArray(gl13Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl13DisableVertexAttribArray(gl13Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl13DrawBuffers(gl13Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl13GetActiveAttrib(gl13Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl13GetActiveUniform(gl13Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl13GetAttachedShaders(gl13Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl13GetAttribLocation(gl13Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl13GetProgramiv(gl13Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl13GetProgramInfoLog(gl13Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl13GetShaderiv(gl13Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl13GetShaderInfoLog(gl13Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl13GetShaderSource(gl13Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl13GetUniformfv(gl13Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl13GetUniformiv(gl13Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl13GetUniformLocation(gl13Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl13GetVertexAttribdv(gl13Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl13GetVertexAttribfv(gl13Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl13GetVertexAttribiv(gl13Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl13GetVertexAttribPointerv(gl13Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl13IsProgram(gl13Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl13IsShader(gl13Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl13LinkProgram(gl13Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl13ShaderSource(gl13Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl13StencilFuncSeparate(gl13Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl13StencilMaskSeparate(gl13Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl13StencilOpSeparate(gl13Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl13Uniform1f(gl13Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl13Uniform2f(gl13Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl13Uniform3f(gl13Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl13Uniform4f(gl13Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl13Uniform1i(gl13Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl13Uniform2i(gl13Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl13Uniform3i(gl13Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl13Uniform4i(gl13Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl13Uniform1ui(gl13Context* glc, GLint location, GLuint v0) {
    return glc->fnUniform1ui(location, v0);
}

void gl13Uniform2ui(gl13Context* glc, GLint location, GLuint v0, GLuint v1) {
    return glc->fnUniform2ui(location, v0, v1);
}

void gl13Uniform3ui(gl13Context* glc, GLint location, GLuint v0, GLuint v1, GLuint v2) {
    return glc->fnUniform3ui(location, v0, v1, v2);
}

void gl13Uniform4ui(gl13Context* glc, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {
    return glc->fnUniform4ui(location, v0, v1, v2, v3);
}

void gl13Uniform1fv(gl13Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl13Uniform2fv(gl13Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl13Uniform3fv(gl13Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl13Uniform4fv(gl13Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl13Uniform1iv(gl13Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl13Uniform2iv(gl13Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl13Uniform3iv(gl13Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl13Uniform4iv(gl13Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl13Uniform1uiv(gl13Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform1uiv(location, count, value);
}

void gl13Uniform2uiv(gl13Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform2uiv(location, count, value);
}

void gl13Uniform3uiv(gl13Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform3uiv(location, count, value);
}

void gl13Uniform4uiv(gl13Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform4uiv(location, count, value);
}

void gl13UseProgram(gl13Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl13ValidateProgram(gl13Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl13VertexAttribPointer(gl13Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl13VertexAttrib1f(gl13Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl13VertexAttrib1s(gl13Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl13VertexAttrib1d(gl13Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl13VertexAttrib2f(gl13Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl13VertexAttrib2s(gl13Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl13VertexAttrib2d(gl13Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl13VertexAttrib3f(gl13Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl13VertexAttrib3s(gl13Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl13VertexAttrib3d(gl13Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl13VertexAttrib4f(gl13Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl13VertexAttrib4s(gl13Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl13VertexAttrib4d(gl13Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl13VertexAttrib4Nuv(gl13Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl13VertexAttrib1fv(gl13Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl13VertexAttrib1sv(gl13Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl13VertexAttrib1dv(gl13Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl13VertexAttrib2fv(gl13Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl13VertexAttrib2sv(gl13Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl13VertexAttrib2dv(gl13Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl13VertexAttrib3fv(gl13Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl13VertexAttrib3sv(gl13Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl13VertexAttrib3dv(gl13Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl13VertexAttrib4fv(gl13Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl13VertexAttrib4sv(gl13Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl13VertexAttrib4dv(gl13Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl13VertexAttrib4iv(gl13Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl13VertexAttrib4bv(gl13Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl13VertexAttrib4ubv(gl13Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl13VertexAttrib4usv(gl13Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl13VertexAttrib4uiv(gl13Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl13VertexAttrib4Nbv(gl13Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl13VertexAttrib4Nsv(gl13Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl13VertexAttrib4Niv(gl13Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl13VertexAttrib4Nubv(gl13Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl13VertexAttrib4Nusv(gl13Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl13VertexAttrib4Nuiv(gl13Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl13UniformMatrix2fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl13UniformMatrix3fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl13UniformMatrix4fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl13UniformMatrix2x3fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl13UniformMatrix3x2fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl13UniformMatrix2x4fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl13UniformMatrix4x2fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl13UniformMatrix3x4fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl13UniformMatrix4x3fv(gl13Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl13Context* gl13NewContext() {
    gl13Context* glc = calloc(1, sizeof(gl13Context));

    // Preload all procedures
    glc->fnAccum = (gl13PAccum)gl13LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl13PAlphaFunc)gl13LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl13PBegin)gl13LibGetProcAddress("glBegin");
    glc->fnEnd = (gl13PEnd)gl13LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl13PBitmap)gl13LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl13PBlendFunc)gl13LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl13PCallList)gl13LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl13PCallLists)gl13LibGetProcAddress("glCallLists");
    glc->fnClear = (gl13PClear)gl13LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl13PClearAccum)gl13LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl13PClearColor)gl13LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl13PClearDepth)gl13LibGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl13PClearIndex)gl13LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl13PClearStencil)gl13LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl13PClipPlane)gl13LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl13PColor3b)gl13LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl13PColor3d)gl13LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl13PColor3f)gl13LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl13PColor3i)gl13LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl13PColor3s)gl13LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl13PColor3ub)gl13LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl13PColor3ui)gl13LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl13PColor3us)gl13LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl13PColor4b)gl13LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl13PColor4d)gl13LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl13PColor4f)gl13LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl13PColor4i)gl13LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl13PColor4s)gl13LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl13PColor4ub)gl13LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl13PColor4ui)gl13LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl13PColor4us)gl13LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl13PColor3bv)gl13LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl13PColor3dv)gl13LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl13PColor3fv)gl13LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl13PColor3iv)gl13LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl13PColor3sv)gl13LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl13PColor3ubv)gl13LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl13PColor3uiv)gl13LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl13PColor3usv)gl13LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl13PColor4bv)gl13LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl13PColor4dv)gl13LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl13PColor4fv)gl13LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl13PColor4iv)gl13LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl13PColor4sv)gl13LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl13PColor4ubv)gl13LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl13PColor4uiv)gl13LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl13PColor4usv)gl13LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl13PColorMask)gl13LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl13PColorMaterial)gl13LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl13PCopyPixels)gl13LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl13PCullFace)gl13LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl13PDeleteLists)gl13LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl13PDepthFunc)gl13LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl13PDepthMask)gl13LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl13PDepthRange)gl13LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl13PEnable)gl13LibGetProcAddress("glEnable");
    glc->fnDisable = (gl13PDisable)gl13LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl13PDrawBuffer)gl13LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl13PDrawPixels)gl13LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl13PEdgeFlag)gl13LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl13PEdgeFlagv)gl13LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl13PEdgeFlagPointer)gl13LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl13PEvalCoord1d)gl13LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl13PEvalCoord1f)gl13LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl13PEvalCoord2d)gl13LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl13PEvalCoord2f)gl13LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl13PEvalCoord1dv)gl13LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl13PEvalCoord1fv)gl13LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl13PEvalCoord2dv)gl13LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl13PEvalCoord2fv)gl13LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl13PEvalMesh1)gl13LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl13PEvalMesh2)gl13LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl13PEvalPoint1)gl13LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl13PEvalPoint2)gl13LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl13PFeedbackBuffer)gl13LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl13PFinish)gl13LibGetProcAddress("glFinish");
    glc->fnFlush = (gl13PFlush)gl13LibGetProcAddress("glFlush");
    glc->fnFogf = (gl13PFogf)gl13LibGetProcAddress("glFogf");
    glc->fnFogi = (gl13PFogi)gl13LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl13PFogfv)gl13LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl13PFogiv)gl13LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl13PFrontFace)gl13LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl13PFrustum)gl13LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl13PGenLists)gl13LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl13PGetBooleanv)gl13LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl13PGetDoublev)gl13LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl13PGetFloatv)gl13LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl13PGetIntegerv)gl13LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl13PGetClipPlane)gl13LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl13PGetError)gl13LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl13PGetLightfv)gl13LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl13PGetLightiv)gl13LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl13PGetMapdv)gl13LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl13PGetMapfv)gl13LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl13PGetMapiv)gl13LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl13PGetMaterialfv)gl13LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl13PGetMaterialiv)gl13LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl13PGetPixelMapfv)gl13LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl13PGetPixelMapuiv)gl13LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl13PGetPixelMapusv)gl13LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl13PGetPolygonStipple)gl13LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl13PGetString)gl13LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl13PGetTexEnvfv)gl13LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl13PGetTexEnviv)gl13LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl13PGetTexGendv)gl13LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl13PGetTexGenfv)gl13LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl13PGetTexGeniv)gl13LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl13PGetTexImage)gl13LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl13PGetTexLevelParameterfv)gl13LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl13PGetTexLevelParameteriv)gl13LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl13PGetTexParameterfv)gl13LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl13PGetTexParameteriv)gl13LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl13PHint)gl13LibGetProcAddress("glHint");
    glc->fnIndexd = (gl13PIndexd)gl13LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl13PIndexf)gl13LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl13PIndexi)gl13LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl13PIndexs)gl13LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl13PIndexdv)gl13LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl13PIndexfv)gl13LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl13PIndexiv)gl13LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl13PIndexsv)gl13LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl13PIndexMask)gl13LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl13PIndexPointer)gl13LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl13PInitNames)gl13LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl13PIsEnabled)gl13LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl13PIsList)gl13LibGetProcAddress("glIsList");
    glc->fnLightf = (gl13PLightf)gl13LibGetProcAddress("glLightf");
    glc->fnLighti = (gl13PLighti)gl13LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl13PLightfv)gl13LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl13PLightiv)gl13LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl13PLightModelf)gl13LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl13PLightModeli)gl13LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl13PLightModelfv)gl13LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl13PLightModeliv)gl13LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl13PLineStipple)gl13LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl13PLineWidth)gl13LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl13PListBase)gl13LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl13PLoadIdentity)gl13LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl13PLoadMatrixd)gl13LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl13PLoadMatrixf)gl13LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl13PLoadName)gl13LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl13PLogicOp)gl13LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl13PMap1d)gl13LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl13PMap1f)gl13LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl13PMap2d)gl13LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl13PMap2f)gl13LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl13PMapGrid1d)gl13LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl13PMapGrid1f)gl13LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl13PMapGrid2d)gl13LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl13PMapGrid2f)gl13LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl13PMaterialf)gl13LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl13PMateriali)gl13LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl13PMaterialfv)gl13LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl13PMaterialiv)gl13LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl13PMatrixMode)gl13LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl13PMultMatrixd)gl13LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl13PMultMatrixf)gl13LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl13PNewList)gl13LibGetProcAddress("glNewList");
    glc->fnEndList = (gl13PEndList)gl13LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl13PNormal3b)gl13LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl13PNormal3d)gl13LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl13PNormal3f)gl13LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl13PNormal3i)gl13LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl13PNormal3s)gl13LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl13PNormal3bv)gl13LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl13PNormal3dv)gl13LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl13PNormal3fv)gl13LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl13PNormal3iv)gl13LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl13PNormal3sv)gl13LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl13POrtho)gl13LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl13PPassThrough)gl13LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl13PPixelMapfv)gl13LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl13PPixelMapuiv)gl13LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl13PPixelMapusv)gl13LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl13PPixelStoref)gl13LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl13PPixelStorei)gl13LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl13PPixelTransferf)gl13LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl13PPixelTransferi)gl13LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl13PPixelZoom)gl13LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl13PPointSize)gl13LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl13PPolygonMode)gl13LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl13PPolygonStipple)gl13LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl13PPushAttrib)gl13LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl13PPopAttrib)gl13LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl13PPushMatrix)gl13LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl13PPopMatrix)gl13LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl13PPushName)gl13LibGetProcAddress("glPushName");
    glc->fnPopName = (gl13PPopName)gl13LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl13PRasterPos2d)gl13LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl13PRasterPos2f)gl13LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl13PRasterPos2i)gl13LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl13PRasterPos2s)gl13LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl13PRasterPos3d)gl13LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl13PRasterPos3f)gl13LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl13PRasterPos3i)gl13LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl13PRasterPos3s)gl13LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl13PRasterPos4d)gl13LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl13PRasterPos4f)gl13LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl13PRasterPos4i)gl13LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl13PRasterPos4s)gl13LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl13PRasterPos2dv)gl13LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl13PRasterPos2fv)gl13LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl13PRasterPos2iv)gl13LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl13PRasterPos2sv)gl13LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl13PRasterPos3dv)gl13LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl13PRasterPos3fv)gl13LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl13PRasterPos3iv)gl13LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl13PRasterPos3sv)gl13LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl13PRasterPos4dv)gl13LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl13PRasterPos4fv)gl13LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl13PRasterPos4iv)gl13LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl13PRasterPos4sv)gl13LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl13PReadBuffer)gl13LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl13PReadPixels)gl13LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl13PRectd)gl13LibGetProcAddress("glRectd");
    glc->fnRectf = (gl13PRectf)gl13LibGetProcAddress("glRectf");
    glc->fnRecti = (gl13PRecti)gl13LibGetProcAddress("glRecti");
    glc->fnRects = (gl13PRects)gl13LibGetProcAddress("glRects");
    glc->fnRectdv = (gl13PRectdv)gl13LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl13PRectfv)gl13LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl13PRectiv)gl13LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl13PRectsv)gl13LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl13PRenderMode)gl13LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl13PRotated)gl13LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl13PRotatef)gl13LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl13PScaled)gl13LibGetProcAddress("glScaled");
    glc->fnScalef = (gl13PScalef)gl13LibGetProcAddress("glScalef");
    glc->fnScissor = (gl13PScissor)gl13LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl13PSelectBuffer)gl13LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl13PShadeModel)gl13LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl13PStencilFunc)gl13LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl13PStencilMask)gl13LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl13PStencilOp)gl13LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl13PTexCoord1d)gl13LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl13PTexCoord1f)gl13LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl13PTexCoord1i)gl13LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl13PTexCoord1s)gl13LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl13PTexCoord2d)gl13LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl13PTexCoord2f)gl13LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl13PTexCoord2i)gl13LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl13PTexCoord2s)gl13LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl13PTexCoord3d)gl13LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl13PTexCoord3f)gl13LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl13PTexCoord3i)gl13LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl13PTexCoord3s)gl13LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl13PTexCoord4d)gl13LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl13PTexCoord4f)gl13LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl13PTexCoord4i)gl13LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl13PTexCoord4s)gl13LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl13PTexCoord1dv)gl13LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl13PTexCoord1fv)gl13LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl13PTexCoord1iv)gl13LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl13PTexCoord1sv)gl13LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl13PTexCoord2dv)gl13LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl13PTexCoord2fv)gl13LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl13PTexCoord2iv)gl13LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl13PTexCoord2sv)gl13LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl13PTexCoord3dv)gl13LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl13PTexCoord3fv)gl13LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl13PTexCoord3iv)gl13LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl13PTexCoord3sv)gl13LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl13PTexCoord4dv)gl13LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl13PTexCoord4fv)gl13LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl13PTexCoord4iv)gl13LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl13PTexCoord4sv)gl13LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl13PTexEnvf)gl13LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl13PTexEnvi)gl13LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl13PTexEnvfv)gl13LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl13PTexEnviv)gl13LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl13PTexGend)gl13LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl13PTexGenf)gl13LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl13PTexGeni)gl13LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl13PTexGendv)gl13LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl13PTexGenfv)gl13LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl13PTexGeniv)gl13LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl13PTexImage1D)gl13LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl13PTexImage2D)gl13LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl13PTexParameterf)gl13LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl13PTexParameteri)gl13LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl13PTexParameterfv)gl13LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl13PTexParameteriv)gl13LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl13PTranslated)gl13LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl13PTranslatef)gl13LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl13PVertex2s)gl13LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl13PVertex2i)gl13LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl13PVertex2f)gl13LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl13PVertex2d)gl13LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl13PVertex3s)gl13LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl13PVertex3i)gl13LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl13PVertex3f)gl13LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl13PVertex3d)gl13LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl13PVertex4s)gl13LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl13PVertex4i)gl13LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl13PVertex4f)gl13LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl13PVertex4d)gl13LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl13PViewport)gl13LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl13PGetConvolutionParameterfv)gl13LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl13PGetConvolutionParameteriv)gl13LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl13PAreTexturesResident)gl13LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl13PArrayElement)gl13LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl13PDrawArrays)gl13LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl13PDrawElements)gl13LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl13PGetPointerv)gl13LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl13PPolygonOffset)gl13LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl13PCopyTexImage1D)gl13LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl13PCopyTexImage2D)gl13LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl13PCopyTexSubImage1D)gl13LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl13PCopyTexSubImage2D)gl13LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl13PBindTexture)gl13LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl13PDeleteTextures)gl13LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl13PGenTextures)gl13LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl13PIsTexture)gl13LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl13PColorPointer)gl13LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl13PEnableClientState)gl13LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl13PDisableClientState)gl13LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl13PIndexub)gl13LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl13PIndexubv)gl13LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl13PInterleavedArrays)gl13LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl13PNormalPointer)gl13LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl13PPushClientAttrib)gl13LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl13PPrioritizeTextures)gl13LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl13PPopClientAttrib)gl13LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl13PTexCoordPointer)gl13LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl13PTexSubImage1D)gl13LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl13PTexSubImage2D)gl13LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl13PVertexPointer)gl13LibGetProcAddress("glVertexPointer");
    glc->fnColorTable = (gl13PColorTable)gl13GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl13PColorTableParameterfv)gl13GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl13PColorTableParameteriv)gl13GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl13PColorSubTable)gl13GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl13PConvolutionFilter1D)gl13GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl13PConvolutionFilter2D)gl13GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl13PConvolutionParameterf)gl13GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl13PConvolutionParameteri)gl13GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl13PCopyColorTable)gl13GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl13PCopyColorSubTable)gl13GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl13PCopyConvolutionFilter1D)gl13GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl13PCopyConvolutionFilter2D)gl13GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl13PGetColorTable)gl13GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl13PGetColorTableParameterfv)gl13GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl13PGetColorTableParameteriv)gl13GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl13PGetConvolutionFilter)gl13GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl13PGetHistogram)gl13GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl13PGetHistogramParameterfv)gl13GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl13PGetHistogramParameteriv)gl13GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl13PGetSeparableFilter)gl13GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl13PHistogram)gl13GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl13PMinmax)gl13GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl13PMultiTexCoord1s)gl13GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl13PMultiTexCoord1i)gl13GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl13PMultiTexCoord1f)gl13GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl13PMultiTexCoord1d)gl13GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl13PMultiTexCoord2s)gl13GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl13PMultiTexCoord2i)gl13GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl13PMultiTexCoord2f)gl13GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl13PMultiTexCoord2d)gl13GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl13PMultiTexCoord3s)gl13GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl13PMultiTexCoord3i)gl13GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl13PMultiTexCoord3f)gl13GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl13PMultiTexCoord3d)gl13GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl13PMultiTexCoord4s)gl13GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl13PMultiTexCoord4i)gl13GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl13PMultiTexCoord4f)gl13GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl13PMultiTexCoord4d)gl13GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl13PMultiTexCoord1sv)gl13GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl13PMultiTexCoord1iv)gl13GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl13PMultiTexCoord1fv)gl13GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl13PMultiTexCoord1dv)gl13GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl13PMultiTexCoord2sv)gl13GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl13PMultiTexCoord2iv)gl13GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl13PMultiTexCoord2fv)gl13GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl13PMultiTexCoord2dv)gl13GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl13PMultiTexCoord3sv)gl13GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl13PMultiTexCoord3iv)gl13GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl13PMultiTexCoord3fv)gl13GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl13PMultiTexCoord3dv)gl13GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl13PMultiTexCoord4sv)gl13GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl13PMultiTexCoord4iv)gl13GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl13PMultiTexCoord4fv)gl13GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl13PMultiTexCoord4dv)gl13GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl13PResetHistogram)gl13GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl13PResetMinmax)gl13GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl13PSeparableFilter2D)gl13GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl13PBlendColor)gl13GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl13PBlendEquation)gl13GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl13PCopyTexSubImage3D)gl13GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl13PDrawRangeElements)gl13GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl13PTexImage3D)gl13GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl13PTexSubImage3D)gl13GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl13PActiveTexture)gl13GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl13PClientActiveTexture)gl13GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl13PCompressedTexImage1D)gl13GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl13PCompressedTexImage2D)gl13GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl13PCompressedTexImage3D)gl13GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl13PCompressedTexSubImage1D)gl13GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl13PCompressedTexSubImage2D)gl13GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl13PCompressedTexSubImage3D)gl13GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl13PGetCompressedTexImage)gl13GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl13PLoadTransposeMatrixd)gl13GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl13PLoadTransposeMatrixf)gl13GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl13PMultTransposeMatrixd)gl13GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl13PMultTransposeMatrixf)gl13GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl13PSampleCoverage)gl13GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl13PBlendFuncSeparate)gl13GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl13PFogCoordPointer)gl13GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl13PFogCoordd)gl13GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl13PFogCoordf)gl13GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl13PFogCoorddv)gl13GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl13PFogCoordfv)gl13GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl13PMultiDrawArrays)gl13GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl13PMultiDrawElements)gl13GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl13PPointParameterf)gl13GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl13PPointParameteri)gl13GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl13PSecondaryColor3b)gl13GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl13PSecondaryColor3s)gl13GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl13PSecondaryColor3i)gl13GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl13PSecondaryColor3f)gl13GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl13PSecondaryColor3d)gl13GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl13PSecondaryColor3ub)gl13GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl13PSecondaryColor3us)gl13GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl13PSecondaryColor3ui)gl13GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl13PSecondaryColor3bv)gl13GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl13PSecondaryColor3sv)gl13GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl13PSecondaryColor3iv)gl13GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl13PSecondaryColor3fv)gl13GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl13PSecondaryColor3dv)gl13GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl13PSecondaryColor3ubv)gl13GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl13PSecondaryColor3usv)gl13GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl13PSecondaryColor3uiv)gl13GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl13PSecondaryColorPointer)gl13GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl13PWindowPos2s)gl13GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl13PWindowPos2i)gl13GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl13PWindowPos2f)gl13GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl13PWindowPos2d)gl13GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl13PWindowPos3s)gl13GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl13PWindowPos3i)gl13GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl13PWindowPos3f)gl13GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl13PWindowPos3d)gl13GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl13PWindowPos2sv)gl13GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl13PWindowPos2iv)gl13GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl13PWindowPos2fv)gl13GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl13PWindowPos2dv)gl13GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl13PWindowPos3sv)gl13GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl13PWindowPos3iv)gl13GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl13PWindowPos3fv)gl13GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl13PWindowPos3dv)gl13GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl13PBeginQuery)gl13GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl13PBindBuffer)gl13GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl13PBufferData)gl13GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl13PBufferSubData)gl13GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl13PDeleteBuffers)gl13GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl13PDeleteQueries)gl13GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl13PGenBuffers)gl13GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl13PGenQueries)gl13GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl13PGetBufferParameteriv)gl13GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl13PGetBufferPointerv)gl13GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl13PGetBufferSubData)gl13GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl13PGetQueryObjectiv)gl13GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl13PGetQueryObjectuiv)gl13GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl13PGetQueryiv)gl13GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl13PIsBuffer)gl13GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl13PIsQuery)gl13GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl13PMapBuffer)gl13GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl13PUnmapBuffer)gl13GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl13PAttachShader)gl13GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl13PBindAttribLocation)gl13GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl13PBlendEquationSeperate)gl13GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl13PCompileShader)gl13GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl13PCreateProgram)gl13GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl13PCreateShader)gl13GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl13PDeleteProgram)gl13GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl13PDeleteShader)gl13GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl13PDetachShader)gl13GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl13PEnableVertexAttribArray)gl13GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl13PDisableVertexAttribArray)gl13GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl13PDrawBuffers)gl13GLGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl13PGetActiveAttrib)gl13GLGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl13PGetActiveUniform)gl13GLGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl13PGetAttachedShaders)gl13GLGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl13PGetAttribLocation)gl13GLGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl13PGetProgramiv)gl13GLGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl13PGetProgramInfoLog)gl13GLGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl13PGetShaderiv)gl13GLGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl13PGetShaderInfoLog)gl13GLGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl13PGetShaderSource)gl13GLGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl13PGetUniformfv)gl13GLGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl13PGetUniformiv)gl13GLGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl13PGetUniformLocation)gl13GLGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl13PGetVertexAttribdv)gl13GLGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl13PGetVertexAttribfv)gl13GLGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl13PGetVertexAttribiv)gl13GLGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl13PGetVertexAttribPointerv)gl13GLGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl13PIsProgram)gl13GLGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl13PIsShader)gl13GLGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl13PLinkProgram)gl13GLGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl13PShaderSource)gl13GLGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl13PStencilFuncSeparate)gl13GLGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl13PStencilMaskSeparate)gl13GLGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl13PStencilOpSeparate)gl13GLGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl13PUniform1f)gl13GLGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl13PUniform2f)gl13GLGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl13PUniform3f)gl13GLGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl13PUniform4f)gl13GLGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl13PUniform1i)gl13GLGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl13PUniform2i)gl13GLGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl13PUniform3i)gl13GLGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl13PUniform4i)gl13GLGetProcAddress("glUniform4i");
    glc->fnUniform1ui = (gl13PUniform1ui)gl13GLGetProcAddress("glUniform1ui");
    glc->fnUniform2ui = (gl13PUniform2ui)gl13GLGetProcAddress("glUniform2ui");
    glc->fnUniform3ui = (gl13PUniform3ui)gl13GLGetProcAddress("glUniform3ui");
    glc->fnUniform4ui = (gl13PUniform4ui)gl13GLGetProcAddress("glUniform4ui");
    glc->fnUniform1fv = (gl13PUniform1fv)gl13GLGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl13PUniform2fv)gl13GLGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl13PUniform3fv)gl13GLGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl13PUniform4fv)gl13GLGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl13PUniform1iv)gl13GLGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl13PUniform2iv)gl13GLGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl13PUniform3iv)gl13GLGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl13PUniform4iv)gl13GLGetProcAddress("glUniform4iv");
    glc->fnUniform1uiv = (gl13PUniform1uiv)gl13GLGetProcAddress("glUniform1uiv");
    glc->fnUniform2uiv = (gl13PUniform2uiv)gl13GLGetProcAddress("glUniform2uiv");
    glc->fnUniform3uiv = (gl13PUniform3uiv)gl13GLGetProcAddress("glUniform3uiv");
    glc->fnUniform4uiv = (gl13PUniform4uiv)gl13GLGetProcAddress("glUniform4uiv");
    glc->fnUseProgram = (gl13PUseProgram)gl13GLGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl13PValidateProgram)gl13GLGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl13PVertexAttribPointer)gl13GLGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl13PVertexAttrib1f)gl13GLGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl13PVertexAttrib1s)gl13GLGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl13PVertexAttrib1d)gl13GLGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl13PVertexAttrib2f)gl13GLGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl13PVertexAttrib2s)gl13GLGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl13PVertexAttrib2d)gl13GLGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl13PVertexAttrib3f)gl13GLGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl13PVertexAttrib3s)gl13GLGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl13PVertexAttrib3d)gl13GLGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl13PVertexAttrib4f)gl13GLGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl13PVertexAttrib4s)gl13GLGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl13PVertexAttrib4d)gl13GLGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl13PVertexAttrib4Nuv)gl13GLGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl13PVertexAttrib1fv)gl13GLGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl13PVertexAttrib1sv)gl13GLGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl13PVertexAttrib1dv)gl13GLGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl13PVertexAttrib2fv)gl13GLGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl13PVertexAttrib2sv)gl13GLGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl13PVertexAttrib2dv)gl13GLGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl13PVertexAttrib3fv)gl13GLGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl13PVertexAttrib3sv)gl13GLGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl13PVertexAttrib3dv)gl13GLGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl13PVertexAttrib4fv)gl13GLGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl13PVertexAttrib4sv)gl13GLGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl13PVertexAttrib4dv)gl13GLGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl13PVertexAttrib4iv)gl13GLGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl13PVertexAttrib4bv)gl13GLGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl13PVertexAttrib4ubv)gl13GLGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl13PVertexAttrib4usv)gl13GLGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl13PVertexAttrib4uiv)gl13GLGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl13PVertexAttrib4Nbv)gl13GLGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl13PVertexAttrib4Nsv)gl13GLGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl13PVertexAttrib4Niv)gl13GLGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl13PVertexAttrib4Nubv)gl13GLGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl13PVertexAttrib4Nusv)gl13GLGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl13PVertexAttrib4Nuiv)gl13GLGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl13PUniformMatrix2fv)gl13GLGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl13PUniformMatrix3fv)gl13GLGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl13PUniformMatrix4fv)gl13GLGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl13PUniformMatrix2x3fv)gl13GLGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl13PUniformMatrix3x2fv)gl13GLGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl13PUniformMatrix2x4fv)gl13GLGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl13PUniformMatrix4x2fv)gl13GLGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl13PUniformMatrix3x4fv)gl13GLGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl13PUniformMatrix4x3fv)gl13GLGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

