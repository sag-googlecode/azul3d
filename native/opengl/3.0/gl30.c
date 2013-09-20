// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl30.h"

#ifdef _WIN32
HMODULE gl30OpenGL32;

void* gl30LibGetProcAddress(char* name) {
	if(gl30OpenGL32 == NULL) {
		gl30OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl30OpenGL32, TEXT(name));
}

void* gl30GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl30Accum(gl30Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl30AlphaFunc(gl30Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl30Begin(gl30Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl30End(gl30Context* glc) {
    return glc->fnEnd();
}

void gl30Bitmap(gl30Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl30BlendFunc(gl30Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl30CallList(gl30Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl30CallLists(gl30Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl30Clear(gl30Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl30ClearAccum(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl30ClearColor(gl30Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl30ClearDepth(gl30Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl30ClearIndex(gl30Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl30ClearStencil(gl30Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl30ClipPlane(gl30Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl30Color3b(gl30Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl30Color3d(gl30Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl30Color3f(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl30Color3i(gl30Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl30Color3s(gl30Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl30Color3ub(gl30Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl30Color3ui(gl30Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl30Color3us(gl30Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl30Color4b(gl30Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl30Color4d(gl30Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl30Color4f(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl30Color4i(gl30Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl30Color4s(gl30Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl30Color4ub(gl30Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl30Color4ui(gl30Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl30Color4us(gl30Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl30Color3bv(gl30Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl30Color3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl30Color3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl30Color3iv(gl30Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl30Color3sv(gl30Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl30Color3ubv(gl30Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl30Color3uiv(gl30Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl30Color3usv(gl30Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl30Color4bv(gl30Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl30Color4dv(gl30Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl30Color4fv(gl30Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl30Color4iv(gl30Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl30Color4sv(gl30Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl30Color4ubv(gl30Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl30Color4uiv(gl30Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl30Color4usv(gl30Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl30ColorMask(gl30Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl30ColorMaterial(gl30Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl30CopyPixels(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl30CullFace(gl30Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl30DeleteLists(gl30Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl30DepthFunc(gl30Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl30DepthMask(gl30Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl30DepthRange(gl30Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl30Enable(gl30Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl30Disable(gl30Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl30DrawBuffer(gl30Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl30DrawPixels(gl30Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl30EdgeFlag(gl30Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl30EdgeFlagv(gl30Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl30EdgeFlagPointer(gl30Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl30EvalCoord1d(gl30Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl30EvalCoord1f(gl30Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl30EvalCoord2d(gl30Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl30EvalCoord2f(gl30Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl30EvalCoord1dv(gl30Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl30EvalCoord1fv(gl30Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl30EvalCoord2dv(gl30Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl30EvalCoord2fv(gl30Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl30EvalMesh1(gl30Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl30EvalMesh2(gl30Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl30EvalPoint1(gl30Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl30EvalPoint2(gl30Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl30FeedbackBuffer(gl30Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl30Finish(gl30Context* glc) {
    return glc->fnFinish();
}

void gl30Flush(gl30Context* glc) {
    return glc->fnFlush();
}

void gl30Fogf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl30Fogi(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl30Fogfv(gl30Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl30Fogiv(gl30Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl30FrontFace(gl30Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl30Frustum(gl30Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl30GenLists(gl30Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl30GetBooleanv(gl30Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl30GetDoublev(gl30Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl30GetFloatv(gl30Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl30GetIntegerv(gl30Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl30GetClipPlane(gl30Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl30GetError(gl30Context* glc) {
    return glc->fnGetError();
}

void gl30GetLightfv(gl30Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl30GetLightiv(gl30Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl30GetMapdv(gl30Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl30GetMapfv(gl30Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl30GetMapiv(gl30Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl30GetMaterialfv(gl30Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl30GetMaterialiv(gl30Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl30GetPixelMapfv(gl30Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl30GetPixelMapuiv(gl30Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl30GetPixelMapusv(gl30Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl30GetPolygonStipple(gl30Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl30GetString(gl30Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl30GetTexEnvfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl30GetTexEnviv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl30GetTexGendv(gl30Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl30GetTexGenfv(gl30Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl30GetTexGeniv(gl30Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl30GetTexImage(gl30Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl30GetTexLevelParameterfv(gl30Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl30GetTexLevelParameteriv(gl30Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl30GetTexParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl30GetTexParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl30Hint(gl30Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl30Indexd(gl30Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl30Indexf(gl30Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl30Indexi(gl30Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl30Indexs(gl30Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl30Indexdv(gl30Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl30Indexfv(gl30Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl30Indexiv(gl30Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl30Indexsv(gl30Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl30IndexMask(gl30Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl30IndexPointer(gl30Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl30InitNames(gl30Context* glc) {
    return glc->fnInitNames();
}

void gl30IsEnabled(gl30Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl30IsList(gl30Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl30Lightf(gl30Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl30Lighti(gl30Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl30Lightfv(gl30Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl30Lightiv(gl30Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl30LightModelf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl30LightModeli(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl30LightModelfv(gl30Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl30LightModeliv(gl30Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl30LineStipple(gl30Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl30LineWidth(gl30Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl30ListBase(gl30Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl30LoadIdentity(gl30Context* glc) {
    return glc->fnLoadIdentity();
}

void gl30LoadMatrixd(gl30Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl30LoadMatrixf(gl30Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl30LoadName(gl30Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl30LogicOp(gl30Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl30Map1d(gl30Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl30Map1f(gl30Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl30Map2d(gl30Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl30Map2f(gl30Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl30MapGrid1d(gl30Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl30MapGrid1f(gl30Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl30MapGrid2d(gl30Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl30MapGrid2f(gl30Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl30Materialf(gl30Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl30Materiali(gl30Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl30Materialfv(gl30Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl30Materialiv(gl30Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl30MatrixMode(gl30Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl30MultMatrixd(gl30Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl30MultMatrixf(gl30Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl30NewList(gl30Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl30EndList(gl30Context* glc) {
    return glc->fnEndList();
}

void gl30Normal3b(gl30Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl30Normal3d(gl30Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl30Normal3f(gl30Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl30Normal3i(gl30Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl30Normal3s(gl30Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl30Normal3bv(gl30Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl30Normal3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl30Normal3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl30Normal3iv(gl30Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl30Normal3sv(gl30Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl30Ortho(gl30Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl30PassThrough(gl30Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl30PixelMapfv(gl30Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl30PixelMapuiv(gl30Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl30PixelMapusv(gl30Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl30PixelStoref(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl30PixelStorei(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl30PixelTransferf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl30PixelTransferi(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl30PixelZoom(gl30Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl30PointSize(gl30Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl30PolygonMode(gl30Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl30PolygonStipple(gl30Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl30PushAttrib(gl30Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl30PopAttrib(gl30Context* glc) {
    return glc->fnPopAttrib();
}

void gl30PushMatrix(gl30Context* glc) {
    return glc->fnPushMatrix();
}

void gl30PopMatrix(gl30Context* glc) {
    return glc->fnPopMatrix();
}

void gl30PushName(gl30Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl30PopName(gl30Context* glc) {
    return glc->fnPopName();
}

void gl30RasterPos2d(gl30Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl30RasterPos2f(gl30Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl30RasterPos2i(gl30Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl30RasterPos2s(gl30Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl30RasterPos3d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl30RasterPos3f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl30RasterPos3i(gl30Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl30RasterPos3s(gl30Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl30RasterPos4d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl30RasterPos4f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl30RasterPos4i(gl30Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl30RasterPos4s(gl30Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl30RasterPos2dv(gl30Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl30RasterPos2fv(gl30Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl30RasterPos2iv(gl30Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl30RasterPos2sv(gl30Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl30RasterPos3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl30RasterPos3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl30RasterPos3iv(gl30Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl30RasterPos3sv(gl30Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl30RasterPos4dv(gl30Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl30RasterPos4fv(gl30Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl30RasterPos4iv(gl30Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl30RasterPos4sv(gl30Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl30ReadBuffer(gl30Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl30ReadPixels(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl30Rectd(gl30Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl30Rectf(gl30Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl30Recti(gl30Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl30Rects(gl30Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl30Rectdv(gl30Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl30Rectfv(gl30Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl30Rectiv(gl30Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl30Rectsv(gl30Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl30RenderMode(gl30Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl30Rotated(gl30Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl30Rotatef(gl30Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl30Scaled(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl30Scalef(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl30Scissor(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl30SelectBuffer(gl30Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl30ShadeModel(gl30Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl30StencilFunc(gl30Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl30StencilMask(gl30Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl30StencilOp(gl30Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl30TexCoord1d(gl30Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl30TexCoord1f(gl30Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl30TexCoord1i(gl30Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl30TexCoord1s(gl30Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl30TexCoord2d(gl30Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl30TexCoord2f(gl30Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl30TexCoord2i(gl30Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl30TexCoord2s(gl30Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl30TexCoord3d(gl30Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl30TexCoord3f(gl30Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl30TexCoord3i(gl30Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl30TexCoord3s(gl30Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl30TexCoord4d(gl30Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl30TexCoord4f(gl30Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl30TexCoord4i(gl30Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl30TexCoord4s(gl30Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl30TexCoord1dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl30TexCoord1fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl30TexCoord1iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl30TexCoord1sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl30TexCoord2dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl30TexCoord2fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl30TexCoord2iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl30TexCoord2sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl30TexCoord3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl30TexCoord3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl30TexCoord3iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl30TexCoord3sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl30TexCoord4dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl30TexCoord4fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl30TexCoord4iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl30TexCoord4sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl30TexEnvf(gl30Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl30TexEnvi(gl30Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl30TexEnvfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl30TexEnviv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl30TexGend(gl30Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl30TexGenf(gl30Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl30TexGeni(gl30Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl30TexGendv(gl30Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl30TexGenfv(gl30Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl30TexGeniv(gl30Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl30TexImage1D(gl30Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl30TexImage2D(gl30Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl30TexParameterf(gl30Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl30TexParameteri(gl30Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl30TexParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl30TexParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl30Translated(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl30Translatef(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl30Vertex2s(gl30Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl30Vertex2i(gl30Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl30Vertex2f(gl30Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl30Vertex2d(gl30Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl30Vertex3s(gl30Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl30Vertex3i(gl30Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl30Vertex3f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl30Vertex3d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl30Vertex4s(gl30Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl30Vertex4i(gl30Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl30Vertex4f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl30Vertex4d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl30Viewport(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl30GetConvolutionParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl30GetConvolutionParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl30AreTexturesResident(gl30Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl30ArrayElement(gl30Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl30DrawArrays(gl30Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl30DrawElements(gl30Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl30GetPointerv(gl30Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl30PolygonOffset(gl30Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl30CopyTexImage1D(gl30Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl30CopyTexImage2D(gl30Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl30CopyTexSubImage1D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl30CopyTexSubImage2D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl30BindTexture(gl30Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl30DeleteTextures(gl30Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl30GenTextures(gl30Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl30IsTexture(gl30Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl30ColorPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl30EnableClientState(gl30Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl30DisableClientState(gl30Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl30Indexub(gl30Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl30Indexubv(gl30Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl30InterleavedArrays(gl30Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl30NormalPointer(gl30Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl30PushClientAttrib(gl30Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl30PrioritizeTextures(gl30Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl30PopClientAttrib(gl30Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl30TexCoordPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl30TexSubImage1D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl30TexSubImage2D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl30VertexPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl30ColorTable(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl30ColorTableParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl30ColorTableParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl30ColorSubTable(gl30Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl30ConvolutionFilter1D(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl30ConvolutionFilter2D(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl30ConvolutionParameterf(gl30Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl30ConvolutionParameteri(gl30Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl30CopyColorTable(gl30Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl30CopyColorSubTable(gl30Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl30CopyConvolutionFilter1D(gl30Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl30CopyConvolutionFilter2D(gl30Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl30GetColorTable(gl30Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl30GetColorTableParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl30GetColorTableParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl30GetConvolutionFilter(gl30Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl30GetHistogram(gl30Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl30GetHistogramParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl30GetHistogramParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl30GetSeparableFilter(gl30Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl30Histogram(gl30Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl30Minmax(gl30Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl30MultiTexCoord1s(gl30Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl30MultiTexCoord1i(gl30Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl30MultiTexCoord1f(gl30Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl30MultiTexCoord1d(gl30Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl30MultiTexCoord2s(gl30Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl30MultiTexCoord2i(gl30Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl30MultiTexCoord2f(gl30Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl30MultiTexCoord2d(gl30Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl30MultiTexCoord3s(gl30Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl30MultiTexCoord3i(gl30Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl30MultiTexCoord3f(gl30Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl30MultiTexCoord3d(gl30Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl30MultiTexCoord4s(gl30Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl30MultiTexCoord4i(gl30Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl30MultiTexCoord4f(gl30Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl30MultiTexCoord4d(gl30Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl30MultiTexCoord1sv(gl30Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl30MultiTexCoord1iv(gl30Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl30MultiTexCoord1fv(gl30Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl30MultiTexCoord1dv(gl30Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl30MultiTexCoord2sv(gl30Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl30MultiTexCoord2iv(gl30Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl30MultiTexCoord2fv(gl30Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl30MultiTexCoord2dv(gl30Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl30MultiTexCoord3sv(gl30Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl30MultiTexCoord3iv(gl30Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl30MultiTexCoord3fv(gl30Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl30MultiTexCoord3dv(gl30Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl30MultiTexCoord4sv(gl30Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl30MultiTexCoord4iv(gl30Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl30MultiTexCoord4fv(gl30Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl30MultiTexCoord4dv(gl30Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl30ResetHistogram(gl30Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl30ResetMinmax(gl30Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl30SeparableFilter2D(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl30BlendColor(gl30Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl30BlendEquation(gl30Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl30CopyTexSubImage3D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl30DrawRangeElements(gl30Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl30TexImage3D(gl30Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl30TexSubImage3D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl30ActiveTexture(gl30Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl30ClientActiveTexture(gl30Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl30CompressedTexImage1D(gl30Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl30CompressedTexImage2D(gl30Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl30CompressedTexImage3D(gl30Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl30CompressedTexSubImage1D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl30CompressedTexSubImage2D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl30CompressedTexSubImage3D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl30GetCompressedTexImage(gl30Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl30LoadTransposeMatrixd(gl30Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl30LoadTransposeMatrixf(gl30Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl30MultTransposeMatrixd(gl30Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl30MultTransposeMatrixf(gl30Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl30SampleCoverage(gl30Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl30BlendFuncSeparate(gl30Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl30FogCoordPointer(gl30Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl30FogCoordd(gl30Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl30FogCoordf(gl30Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl30FogCoorddv(gl30Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl30FogCoordfv(gl30Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl30MultiDrawArrays(gl30Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl30MultiDrawElements(gl30Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl30PointParameterf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl30PointParameteri(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl30SecondaryColor3b(gl30Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl30SecondaryColor3s(gl30Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl30SecondaryColor3i(gl30Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl30SecondaryColor3f(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl30SecondaryColor3d(gl30Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl30SecondaryColor3ub(gl30Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl30SecondaryColor3us(gl30Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl30SecondaryColor3ui(gl30Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl30SecondaryColor3bv(gl30Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl30SecondaryColor3sv(gl30Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl30SecondaryColor3iv(gl30Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl30SecondaryColor3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl30SecondaryColor3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl30SecondaryColor3ubv(gl30Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl30SecondaryColor3usv(gl30Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl30SecondaryColor3uiv(gl30Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl30SecondaryColorPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl30WindowPos2s(gl30Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl30WindowPos2i(gl30Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl30WindowPos2f(gl30Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl30WindowPos2d(gl30Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl30WindowPos3s(gl30Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl30WindowPos3i(gl30Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl30WindowPos3f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl30WindowPos3d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl30WindowPos2sv(gl30Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl30WindowPos2iv(gl30Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl30WindowPos2fv(gl30Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl30WindowPos2dv(gl30Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl30WindowPos3sv(gl30Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl30WindowPos3iv(gl30Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl30WindowPos3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl30WindowPos3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl30BeginQuery(gl30Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl30BindBuffer(gl30Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl30BufferData(gl30Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl30BufferSubData(gl30Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl30DeleteBuffers(gl30Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl30DeleteQueries(gl30Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl30GenBuffers(gl30Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl30GenQueries(gl30Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl30GetBufferParameteriv(gl30Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl30GetBufferPointerv(gl30Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl30GetBufferSubData(gl30Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl30GetQueryObjectiv(gl30Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl30GetQueryObjectuiv(gl30Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl30GetQueryiv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl30IsBuffer(gl30Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl30IsQuery(gl30Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl30MapBuffer(gl30Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl30UnmapBuffer(gl30Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl30AttachShader(gl30Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl30BindAttribLocation(gl30Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl30BlendEquationSeperate(gl30Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl30CompileShader(gl30Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl30CreateProgram(gl30Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl30CreateShader(gl30Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl30DeleteProgram(gl30Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl30DeleteShader(gl30Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl30DetachShader(gl30Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl30EnableVertexAttribArray(gl30Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl30DisableVertexAttribArray(gl30Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl30DrawBuffers(gl30Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl30GetActiveAttrib(gl30Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl30GetActiveUniform(gl30Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl30GetAttachedShaders(gl30Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl30GetAttribLocation(gl30Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl30GetProgramiv(gl30Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl30GetProgramInfoLog(gl30Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl30GetShaderiv(gl30Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl30GetShaderInfoLog(gl30Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl30GetShaderSource(gl30Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl30GetUniformfv(gl30Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl30GetUniformiv(gl30Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl30GetUniformLocation(gl30Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl30GetVertexAttribdv(gl30Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl30GetVertexAttribfv(gl30Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl30GetVertexAttribiv(gl30Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl30GetVertexAttribPointerv(gl30Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl30IsProgram(gl30Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl30IsShader(gl30Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl30LinkProgram(gl30Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl30ShaderSource(gl30Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl30StencilFuncSeparate(gl30Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl30StencilMaskSeparate(gl30Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl30StencilOpSeparate(gl30Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl30Uniform1f(gl30Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl30Uniform2f(gl30Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl30Uniform3f(gl30Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl30Uniform4f(gl30Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl30Uniform1i(gl30Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl30Uniform2i(gl30Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl30Uniform3i(gl30Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl30Uniform4i(gl30Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl30Uniform1ui(gl30Context* glc, GLint location, GLuint v0) {
    return glc->fnUniform1ui(location, v0);
}

void gl30Uniform2ui(gl30Context* glc, GLint location, GLuint v0, GLuint v1) {
    return glc->fnUniform2ui(location, v0, v1);
}

void gl30Uniform3ui(gl30Context* glc, GLint location, GLuint v0, GLuint v1, GLuint v2) {
    return glc->fnUniform3ui(location, v0, v1, v2);
}

void gl30Uniform4ui(gl30Context* glc, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {
    return glc->fnUniform4ui(location, v0, v1, v2, v3);
}

void gl30Uniform1fv(gl30Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl30Uniform2fv(gl30Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl30Uniform3fv(gl30Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl30Uniform4fv(gl30Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl30Uniform1iv(gl30Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl30Uniform2iv(gl30Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl30Uniform3iv(gl30Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl30Uniform4iv(gl30Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl30Uniform1uiv(gl30Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform1uiv(location, count, value);
}

void gl30Uniform2uiv(gl30Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform2uiv(location, count, value);
}

void gl30Uniform3uiv(gl30Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform3uiv(location, count, value);
}

void gl30Uniform4uiv(gl30Context* glc, GLint location, GLsizei count, GLuint* value) {
    return glc->fnUniform4uiv(location, count, value);
}

void gl30UseProgram(gl30Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl30ValidateProgram(gl30Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl30VertexAttribPointer(gl30Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl30VertexAttrib1f(gl30Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl30VertexAttrib1s(gl30Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl30VertexAttrib1d(gl30Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl30VertexAttrib2f(gl30Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl30VertexAttrib2s(gl30Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl30VertexAttrib2d(gl30Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl30VertexAttrib3f(gl30Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl30VertexAttrib3s(gl30Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl30VertexAttrib3d(gl30Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl30VertexAttrib4f(gl30Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl30VertexAttrib4s(gl30Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl30VertexAttrib4d(gl30Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl30VertexAttrib4Nuv(gl30Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl30VertexAttrib1fv(gl30Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl30VertexAttrib1sv(gl30Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl30VertexAttrib1dv(gl30Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl30VertexAttrib2fv(gl30Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl30VertexAttrib2sv(gl30Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl30VertexAttrib2dv(gl30Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl30VertexAttrib3fv(gl30Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl30VertexAttrib3sv(gl30Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl30VertexAttrib3dv(gl30Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl30VertexAttrib4fv(gl30Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl30VertexAttrib4sv(gl30Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl30VertexAttrib4dv(gl30Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl30VertexAttrib4iv(gl30Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl30VertexAttrib4bv(gl30Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl30VertexAttrib4ubv(gl30Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl30VertexAttrib4usv(gl30Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl30VertexAttrib4uiv(gl30Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl30VertexAttrib4Nbv(gl30Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl30VertexAttrib4Nsv(gl30Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl30VertexAttrib4Niv(gl30Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl30VertexAttrib4Nubv(gl30Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl30VertexAttrib4Nusv(gl30Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl30VertexAttrib4Nuiv(gl30Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl30UniformMatrix2fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl30UniformMatrix3fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl30UniformMatrix4fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl30UniformMatrix2x3fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl30UniformMatrix3x2fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl30UniformMatrix2x4fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl30UniformMatrix4x2fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl30UniformMatrix3x4fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl30UniformMatrix4x3fv(gl30Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl30Context* gl30NewContext() {
    gl30Context* glc = calloc(1, sizeof(gl30Context));

    // Preload all procedures
    glc->fnAccum = (gl30PAccum)gl30LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl30PAlphaFunc)gl30LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl30PBegin)gl30LibGetProcAddress("glBegin");
    glc->fnEnd = (gl30PEnd)gl30LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl30PBitmap)gl30LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl30PBlendFunc)gl30LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl30PCallList)gl30LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl30PCallLists)gl30LibGetProcAddress("glCallLists");
    glc->fnClear = (gl30PClear)gl30LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl30PClearAccum)gl30LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl30PClearColor)gl30LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl30PClearDepth)gl30LibGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl30PClearIndex)gl30LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl30PClearStencil)gl30LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl30PClipPlane)gl30LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl30PColor3b)gl30LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl30PColor3d)gl30LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl30PColor3f)gl30LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl30PColor3i)gl30LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl30PColor3s)gl30LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl30PColor3ub)gl30LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl30PColor3ui)gl30LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl30PColor3us)gl30LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl30PColor4b)gl30LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl30PColor4d)gl30LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl30PColor4f)gl30LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl30PColor4i)gl30LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl30PColor4s)gl30LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl30PColor4ub)gl30LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl30PColor4ui)gl30LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl30PColor4us)gl30LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl30PColor3bv)gl30LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl30PColor3dv)gl30LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl30PColor3fv)gl30LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl30PColor3iv)gl30LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl30PColor3sv)gl30LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl30PColor3ubv)gl30LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl30PColor3uiv)gl30LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl30PColor3usv)gl30LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl30PColor4bv)gl30LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl30PColor4dv)gl30LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl30PColor4fv)gl30LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl30PColor4iv)gl30LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl30PColor4sv)gl30LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl30PColor4ubv)gl30LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl30PColor4uiv)gl30LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl30PColor4usv)gl30LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl30PColorMask)gl30LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl30PColorMaterial)gl30LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl30PCopyPixels)gl30LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl30PCullFace)gl30LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl30PDeleteLists)gl30LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl30PDepthFunc)gl30LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl30PDepthMask)gl30LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl30PDepthRange)gl30LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl30PEnable)gl30LibGetProcAddress("glEnable");
    glc->fnDisable = (gl30PDisable)gl30LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl30PDrawBuffer)gl30LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl30PDrawPixels)gl30LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl30PEdgeFlag)gl30LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl30PEdgeFlagv)gl30LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl30PEdgeFlagPointer)gl30LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl30PEvalCoord1d)gl30LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl30PEvalCoord1f)gl30LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl30PEvalCoord2d)gl30LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl30PEvalCoord2f)gl30LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl30PEvalCoord1dv)gl30LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl30PEvalCoord1fv)gl30LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl30PEvalCoord2dv)gl30LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl30PEvalCoord2fv)gl30LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl30PEvalMesh1)gl30LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl30PEvalMesh2)gl30LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl30PEvalPoint1)gl30LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl30PEvalPoint2)gl30LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl30PFeedbackBuffer)gl30LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl30PFinish)gl30LibGetProcAddress("glFinish");
    glc->fnFlush = (gl30PFlush)gl30LibGetProcAddress("glFlush");
    glc->fnFogf = (gl30PFogf)gl30LibGetProcAddress("glFogf");
    glc->fnFogi = (gl30PFogi)gl30LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl30PFogfv)gl30LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl30PFogiv)gl30LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl30PFrontFace)gl30LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl30PFrustum)gl30LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl30PGenLists)gl30LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl30PGetBooleanv)gl30LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl30PGetDoublev)gl30LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl30PGetFloatv)gl30LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl30PGetIntegerv)gl30LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl30PGetClipPlane)gl30LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl30PGetError)gl30LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl30PGetLightfv)gl30LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl30PGetLightiv)gl30LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl30PGetMapdv)gl30LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl30PGetMapfv)gl30LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl30PGetMapiv)gl30LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl30PGetMaterialfv)gl30LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl30PGetMaterialiv)gl30LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl30PGetPixelMapfv)gl30LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl30PGetPixelMapuiv)gl30LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl30PGetPixelMapusv)gl30LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl30PGetPolygonStipple)gl30LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl30PGetString)gl30LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl30PGetTexEnvfv)gl30LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl30PGetTexEnviv)gl30LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl30PGetTexGendv)gl30LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl30PGetTexGenfv)gl30LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl30PGetTexGeniv)gl30LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl30PGetTexImage)gl30LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl30PGetTexLevelParameterfv)gl30LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl30PGetTexLevelParameteriv)gl30LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl30PGetTexParameterfv)gl30LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl30PGetTexParameteriv)gl30LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl30PHint)gl30LibGetProcAddress("glHint");
    glc->fnIndexd = (gl30PIndexd)gl30LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl30PIndexf)gl30LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl30PIndexi)gl30LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl30PIndexs)gl30LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl30PIndexdv)gl30LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl30PIndexfv)gl30LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl30PIndexiv)gl30LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl30PIndexsv)gl30LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl30PIndexMask)gl30LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl30PIndexPointer)gl30LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl30PInitNames)gl30LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl30PIsEnabled)gl30LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl30PIsList)gl30LibGetProcAddress("glIsList");
    glc->fnLightf = (gl30PLightf)gl30LibGetProcAddress("glLightf");
    glc->fnLighti = (gl30PLighti)gl30LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl30PLightfv)gl30LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl30PLightiv)gl30LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl30PLightModelf)gl30LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl30PLightModeli)gl30LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl30PLightModelfv)gl30LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl30PLightModeliv)gl30LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl30PLineStipple)gl30LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl30PLineWidth)gl30LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl30PListBase)gl30LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl30PLoadIdentity)gl30LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl30PLoadMatrixd)gl30LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl30PLoadMatrixf)gl30LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl30PLoadName)gl30LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl30PLogicOp)gl30LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl30PMap1d)gl30LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl30PMap1f)gl30LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl30PMap2d)gl30LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl30PMap2f)gl30LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl30PMapGrid1d)gl30LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl30PMapGrid1f)gl30LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl30PMapGrid2d)gl30LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl30PMapGrid2f)gl30LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl30PMaterialf)gl30LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl30PMateriali)gl30LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl30PMaterialfv)gl30LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl30PMaterialiv)gl30LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl30PMatrixMode)gl30LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl30PMultMatrixd)gl30LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl30PMultMatrixf)gl30LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl30PNewList)gl30LibGetProcAddress("glNewList");
    glc->fnEndList = (gl30PEndList)gl30LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl30PNormal3b)gl30LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl30PNormal3d)gl30LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl30PNormal3f)gl30LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl30PNormal3i)gl30LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl30PNormal3s)gl30LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl30PNormal3bv)gl30LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl30PNormal3dv)gl30LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl30PNormal3fv)gl30LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl30PNormal3iv)gl30LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl30PNormal3sv)gl30LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl30POrtho)gl30LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl30PPassThrough)gl30LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl30PPixelMapfv)gl30LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl30PPixelMapuiv)gl30LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl30PPixelMapusv)gl30LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl30PPixelStoref)gl30LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl30PPixelStorei)gl30LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl30PPixelTransferf)gl30LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl30PPixelTransferi)gl30LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl30PPixelZoom)gl30LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl30PPointSize)gl30LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl30PPolygonMode)gl30LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl30PPolygonStipple)gl30LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl30PPushAttrib)gl30LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl30PPopAttrib)gl30LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl30PPushMatrix)gl30LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl30PPopMatrix)gl30LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl30PPushName)gl30LibGetProcAddress("glPushName");
    glc->fnPopName = (gl30PPopName)gl30LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl30PRasterPos2d)gl30LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl30PRasterPos2f)gl30LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl30PRasterPos2i)gl30LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl30PRasterPos2s)gl30LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl30PRasterPos3d)gl30LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl30PRasterPos3f)gl30LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl30PRasterPos3i)gl30LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl30PRasterPos3s)gl30LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl30PRasterPos4d)gl30LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl30PRasterPos4f)gl30LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl30PRasterPos4i)gl30LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl30PRasterPos4s)gl30LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl30PRasterPos2dv)gl30LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl30PRasterPos2fv)gl30LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl30PRasterPos2iv)gl30LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl30PRasterPos2sv)gl30LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl30PRasterPos3dv)gl30LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl30PRasterPos3fv)gl30LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl30PRasterPos3iv)gl30LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl30PRasterPos3sv)gl30LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl30PRasterPos4dv)gl30LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl30PRasterPos4fv)gl30LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl30PRasterPos4iv)gl30LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl30PRasterPos4sv)gl30LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl30PReadBuffer)gl30LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl30PReadPixels)gl30LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl30PRectd)gl30LibGetProcAddress("glRectd");
    glc->fnRectf = (gl30PRectf)gl30LibGetProcAddress("glRectf");
    glc->fnRecti = (gl30PRecti)gl30LibGetProcAddress("glRecti");
    glc->fnRects = (gl30PRects)gl30LibGetProcAddress("glRects");
    glc->fnRectdv = (gl30PRectdv)gl30LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl30PRectfv)gl30LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl30PRectiv)gl30LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl30PRectsv)gl30LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl30PRenderMode)gl30LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl30PRotated)gl30LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl30PRotatef)gl30LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl30PScaled)gl30LibGetProcAddress("glScaled");
    glc->fnScalef = (gl30PScalef)gl30LibGetProcAddress("glScalef");
    glc->fnScissor = (gl30PScissor)gl30LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl30PSelectBuffer)gl30LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl30PShadeModel)gl30LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl30PStencilFunc)gl30LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl30PStencilMask)gl30LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl30PStencilOp)gl30LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl30PTexCoord1d)gl30LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl30PTexCoord1f)gl30LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl30PTexCoord1i)gl30LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl30PTexCoord1s)gl30LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl30PTexCoord2d)gl30LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl30PTexCoord2f)gl30LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl30PTexCoord2i)gl30LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl30PTexCoord2s)gl30LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl30PTexCoord3d)gl30LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl30PTexCoord3f)gl30LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl30PTexCoord3i)gl30LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl30PTexCoord3s)gl30LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl30PTexCoord4d)gl30LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl30PTexCoord4f)gl30LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl30PTexCoord4i)gl30LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl30PTexCoord4s)gl30LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl30PTexCoord1dv)gl30LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl30PTexCoord1fv)gl30LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl30PTexCoord1iv)gl30LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl30PTexCoord1sv)gl30LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl30PTexCoord2dv)gl30LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl30PTexCoord2fv)gl30LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl30PTexCoord2iv)gl30LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl30PTexCoord2sv)gl30LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl30PTexCoord3dv)gl30LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl30PTexCoord3fv)gl30LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl30PTexCoord3iv)gl30LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl30PTexCoord3sv)gl30LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl30PTexCoord4dv)gl30LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl30PTexCoord4fv)gl30LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl30PTexCoord4iv)gl30LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl30PTexCoord4sv)gl30LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl30PTexEnvf)gl30LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl30PTexEnvi)gl30LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl30PTexEnvfv)gl30LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl30PTexEnviv)gl30LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl30PTexGend)gl30LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl30PTexGenf)gl30LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl30PTexGeni)gl30LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl30PTexGendv)gl30LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl30PTexGenfv)gl30LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl30PTexGeniv)gl30LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl30PTexImage1D)gl30LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl30PTexImage2D)gl30LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl30PTexParameterf)gl30LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl30PTexParameteri)gl30LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl30PTexParameterfv)gl30LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl30PTexParameteriv)gl30LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl30PTranslated)gl30LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl30PTranslatef)gl30LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl30PVertex2s)gl30LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl30PVertex2i)gl30LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl30PVertex2f)gl30LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl30PVertex2d)gl30LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl30PVertex3s)gl30LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl30PVertex3i)gl30LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl30PVertex3f)gl30LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl30PVertex3d)gl30LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl30PVertex4s)gl30LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl30PVertex4i)gl30LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl30PVertex4f)gl30LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl30PVertex4d)gl30LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl30PViewport)gl30LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl30PGetConvolutionParameterfv)gl30LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl30PGetConvolutionParameteriv)gl30LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl30PAreTexturesResident)gl30LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl30PArrayElement)gl30LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl30PDrawArrays)gl30LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl30PDrawElements)gl30LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl30PGetPointerv)gl30LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl30PPolygonOffset)gl30LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl30PCopyTexImage1D)gl30LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl30PCopyTexImage2D)gl30LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl30PCopyTexSubImage1D)gl30LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl30PCopyTexSubImage2D)gl30LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl30PBindTexture)gl30LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl30PDeleteTextures)gl30LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl30PGenTextures)gl30LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl30PIsTexture)gl30LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl30PColorPointer)gl30LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl30PEnableClientState)gl30LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl30PDisableClientState)gl30LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl30PIndexub)gl30LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl30PIndexubv)gl30LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl30PInterleavedArrays)gl30LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl30PNormalPointer)gl30LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl30PPushClientAttrib)gl30LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl30PPrioritizeTextures)gl30LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl30PPopClientAttrib)gl30LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl30PTexCoordPointer)gl30LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl30PTexSubImage1D)gl30LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl30PTexSubImage2D)gl30LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl30PVertexPointer)gl30LibGetProcAddress("glVertexPointer");
    glc->fnColorTable = (gl30PColorTable)gl30GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl30PColorTableParameterfv)gl30GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl30PColorTableParameteriv)gl30GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl30PColorSubTable)gl30GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl30PConvolutionFilter1D)gl30GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl30PConvolutionFilter2D)gl30GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl30PConvolutionParameterf)gl30GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl30PConvolutionParameteri)gl30GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl30PCopyColorTable)gl30GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl30PCopyColorSubTable)gl30GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl30PCopyConvolutionFilter1D)gl30GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl30PCopyConvolutionFilter2D)gl30GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl30PGetColorTable)gl30GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl30PGetColorTableParameterfv)gl30GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl30PGetColorTableParameteriv)gl30GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl30PGetConvolutionFilter)gl30GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl30PGetHistogram)gl30GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl30PGetHistogramParameterfv)gl30GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl30PGetHistogramParameteriv)gl30GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl30PGetSeparableFilter)gl30GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl30PHistogram)gl30GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl30PMinmax)gl30GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl30PMultiTexCoord1s)gl30GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl30PMultiTexCoord1i)gl30GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl30PMultiTexCoord1f)gl30GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl30PMultiTexCoord1d)gl30GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl30PMultiTexCoord2s)gl30GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl30PMultiTexCoord2i)gl30GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl30PMultiTexCoord2f)gl30GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl30PMultiTexCoord2d)gl30GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl30PMultiTexCoord3s)gl30GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl30PMultiTexCoord3i)gl30GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl30PMultiTexCoord3f)gl30GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl30PMultiTexCoord3d)gl30GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl30PMultiTexCoord4s)gl30GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl30PMultiTexCoord4i)gl30GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl30PMultiTexCoord4f)gl30GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl30PMultiTexCoord4d)gl30GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl30PMultiTexCoord1sv)gl30GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl30PMultiTexCoord1iv)gl30GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl30PMultiTexCoord1fv)gl30GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl30PMultiTexCoord1dv)gl30GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl30PMultiTexCoord2sv)gl30GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl30PMultiTexCoord2iv)gl30GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl30PMultiTexCoord2fv)gl30GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl30PMultiTexCoord2dv)gl30GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl30PMultiTexCoord3sv)gl30GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl30PMultiTexCoord3iv)gl30GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl30PMultiTexCoord3fv)gl30GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl30PMultiTexCoord3dv)gl30GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl30PMultiTexCoord4sv)gl30GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl30PMultiTexCoord4iv)gl30GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl30PMultiTexCoord4fv)gl30GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl30PMultiTexCoord4dv)gl30GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl30PResetHistogram)gl30GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl30PResetMinmax)gl30GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl30PSeparableFilter2D)gl30GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl30PBlendColor)gl30GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl30PBlendEquation)gl30GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl30PCopyTexSubImage3D)gl30GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl30PDrawRangeElements)gl30GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl30PTexImage3D)gl30GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl30PTexSubImage3D)gl30GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl30PActiveTexture)gl30GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl30PClientActiveTexture)gl30GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl30PCompressedTexImage1D)gl30GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl30PCompressedTexImage2D)gl30GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl30PCompressedTexImage3D)gl30GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl30PCompressedTexSubImage1D)gl30GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl30PCompressedTexSubImage2D)gl30GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl30PCompressedTexSubImage3D)gl30GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl30PGetCompressedTexImage)gl30GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl30PLoadTransposeMatrixd)gl30GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl30PLoadTransposeMatrixf)gl30GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl30PMultTransposeMatrixd)gl30GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl30PMultTransposeMatrixf)gl30GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl30PSampleCoverage)gl30GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl30PBlendFuncSeparate)gl30GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl30PFogCoordPointer)gl30GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl30PFogCoordd)gl30GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl30PFogCoordf)gl30GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl30PFogCoorddv)gl30GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl30PFogCoordfv)gl30GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl30PMultiDrawArrays)gl30GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl30PMultiDrawElements)gl30GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl30PPointParameterf)gl30GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl30PPointParameteri)gl30GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl30PSecondaryColor3b)gl30GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl30PSecondaryColor3s)gl30GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl30PSecondaryColor3i)gl30GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl30PSecondaryColor3f)gl30GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl30PSecondaryColor3d)gl30GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl30PSecondaryColor3ub)gl30GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl30PSecondaryColor3us)gl30GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl30PSecondaryColor3ui)gl30GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl30PSecondaryColor3bv)gl30GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl30PSecondaryColor3sv)gl30GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl30PSecondaryColor3iv)gl30GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl30PSecondaryColor3fv)gl30GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl30PSecondaryColor3dv)gl30GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl30PSecondaryColor3ubv)gl30GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl30PSecondaryColor3usv)gl30GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl30PSecondaryColor3uiv)gl30GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl30PSecondaryColorPointer)gl30GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl30PWindowPos2s)gl30GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl30PWindowPos2i)gl30GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl30PWindowPos2f)gl30GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl30PWindowPos2d)gl30GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl30PWindowPos3s)gl30GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl30PWindowPos3i)gl30GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl30PWindowPos3f)gl30GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl30PWindowPos3d)gl30GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl30PWindowPos2sv)gl30GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl30PWindowPos2iv)gl30GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl30PWindowPos2fv)gl30GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl30PWindowPos2dv)gl30GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl30PWindowPos3sv)gl30GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl30PWindowPos3iv)gl30GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl30PWindowPos3fv)gl30GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl30PWindowPos3dv)gl30GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl30PBeginQuery)gl30GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl30PBindBuffer)gl30GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl30PBufferData)gl30GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl30PBufferSubData)gl30GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl30PDeleteBuffers)gl30GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl30PDeleteQueries)gl30GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl30PGenBuffers)gl30GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl30PGenQueries)gl30GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl30PGetBufferParameteriv)gl30GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl30PGetBufferPointerv)gl30GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl30PGetBufferSubData)gl30GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl30PGetQueryObjectiv)gl30GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl30PGetQueryObjectuiv)gl30GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl30PGetQueryiv)gl30GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl30PIsBuffer)gl30GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl30PIsQuery)gl30GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl30PMapBuffer)gl30GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl30PUnmapBuffer)gl30GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl30PAttachShader)gl30GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl30PBindAttribLocation)gl30GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl30PBlendEquationSeperate)gl30GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl30PCompileShader)gl30GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl30PCreateProgram)gl30GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl30PCreateShader)gl30GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl30PDeleteProgram)gl30GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl30PDeleteShader)gl30GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl30PDetachShader)gl30GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl30PEnableVertexAttribArray)gl30GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl30PDisableVertexAttribArray)gl30GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl30PDrawBuffers)gl30GLGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl30PGetActiveAttrib)gl30GLGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl30PGetActiveUniform)gl30GLGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl30PGetAttachedShaders)gl30GLGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl30PGetAttribLocation)gl30GLGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl30PGetProgramiv)gl30GLGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl30PGetProgramInfoLog)gl30GLGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl30PGetShaderiv)gl30GLGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl30PGetShaderInfoLog)gl30GLGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl30PGetShaderSource)gl30GLGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl30PGetUniformfv)gl30GLGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl30PGetUniformiv)gl30GLGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl30PGetUniformLocation)gl30GLGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl30PGetVertexAttribdv)gl30GLGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl30PGetVertexAttribfv)gl30GLGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl30PGetVertexAttribiv)gl30GLGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl30PGetVertexAttribPointerv)gl30GLGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl30PIsProgram)gl30GLGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl30PIsShader)gl30GLGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl30PLinkProgram)gl30GLGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl30PShaderSource)gl30GLGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl30PStencilFuncSeparate)gl30GLGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl30PStencilMaskSeparate)gl30GLGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl30PStencilOpSeparate)gl30GLGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl30PUniform1f)gl30GLGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl30PUniform2f)gl30GLGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl30PUniform3f)gl30GLGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl30PUniform4f)gl30GLGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl30PUniform1i)gl30GLGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl30PUniform2i)gl30GLGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl30PUniform3i)gl30GLGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl30PUniform4i)gl30GLGetProcAddress("glUniform4i");
    glc->fnUniform1ui = (gl30PUniform1ui)gl30GLGetProcAddress("glUniform1ui");
    glc->fnUniform2ui = (gl30PUniform2ui)gl30GLGetProcAddress("glUniform2ui");
    glc->fnUniform3ui = (gl30PUniform3ui)gl30GLGetProcAddress("glUniform3ui");
    glc->fnUniform4ui = (gl30PUniform4ui)gl30GLGetProcAddress("glUniform4ui");
    glc->fnUniform1fv = (gl30PUniform1fv)gl30GLGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl30PUniform2fv)gl30GLGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl30PUniform3fv)gl30GLGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl30PUniform4fv)gl30GLGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl30PUniform1iv)gl30GLGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl30PUniform2iv)gl30GLGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl30PUniform3iv)gl30GLGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl30PUniform4iv)gl30GLGetProcAddress("glUniform4iv");
    glc->fnUniform1uiv = (gl30PUniform1uiv)gl30GLGetProcAddress("glUniform1uiv");
    glc->fnUniform2uiv = (gl30PUniform2uiv)gl30GLGetProcAddress("glUniform2uiv");
    glc->fnUniform3uiv = (gl30PUniform3uiv)gl30GLGetProcAddress("glUniform3uiv");
    glc->fnUniform4uiv = (gl30PUniform4uiv)gl30GLGetProcAddress("glUniform4uiv");
    glc->fnUseProgram = (gl30PUseProgram)gl30GLGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl30PValidateProgram)gl30GLGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl30PVertexAttribPointer)gl30GLGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl30PVertexAttrib1f)gl30GLGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl30PVertexAttrib1s)gl30GLGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl30PVertexAttrib1d)gl30GLGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl30PVertexAttrib2f)gl30GLGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl30PVertexAttrib2s)gl30GLGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl30PVertexAttrib2d)gl30GLGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl30PVertexAttrib3f)gl30GLGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl30PVertexAttrib3s)gl30GLGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl30PVertexAttrib3d)gl30GLGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl30PVertexAttrib4f)gl30GLGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl30PVertexAttrib4s)gl30GLGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl30PVertexAttrib4d)gl30GLGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl30PVertexAttrib4Nuv)gl30GLGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl30PVertexAttrib1fv)gl30GLGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl30PVertexAttrib1sv)gl30GLGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl30PVertexAttrib1dv)gl30GLGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl30PVertexAttrib2fv)gl30GLGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl30PVertexAttrib2sv)gl30GLGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl30PVertexAttrib2dv)gl30GLGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl30PVertexAttrib3fv)gl30GLGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl30PVertexAttrib3sv)gl30GLGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl30PVertexAttrib3dv)gl30GLGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl30PVertexAttrib4fv)gl30GLGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl30PVertexAttrib4sv)gl30GLGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl30PVertexAttrib4dv)gl30GLGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl30PVertexAttrib4iv)gl30GLGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl30PVertexAttrib4bv)gl30GLGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl30PVertexAttrib4ubv)gl30GLGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl30PVertexAttrib4usv)gl30GLGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl30PVertexAttrib4uiv)gl30GLGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl30PVertexAttrib4Nbv)gl30GLGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl30PVertexAttrib4Nsv)gl30GLGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl30PVertexAttrib4Niv)gl30GLGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl30PVertexAttrib4Nubv)gl30GLGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl30PVertexAttrib4Nusv)gl30GLGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl30PVertexAttrib4Nuiv)gl30GLGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl30PUniformMatrix2fv)gl30GLGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl30PUniformMatrix3fv)gl30GLGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl30PUniformMatrix4fv)gl30GLGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl30PUniformMatrix2x3fv)gl30GLGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl30PUniformMatrix3x2fv)gl30GLGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl30PUniformMatrix2x4fv)gl30GLGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl30PUniformMatrix4x2fv)gl30GLGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl30PUniformMatrix3x4fv)gl30GLGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl30PUniformMatrix4x3fv)gl30GLGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

