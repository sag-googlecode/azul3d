
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl15.h"

#ifdef _WIN32
HMODULE gl15OpenGL32;

void* gl15LibGetProcAddress(char* name) {
	if(gl15OpenGL32 == NULL) {
		gl15OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl15OpenGL32, TEXT(name));
}

void* gl15GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl15Accum(gl15Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl15AlphaFunc(gl15Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl15Begin(gl15Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl15End(gl15Context* glc) {
    return glc->fnEnd();
}

void gl15Bitmap(gl15Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl15BlendFunc(gl15Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl15CallList(gl15Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl15CallLists(gl15Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl15Clear(gl15Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl15ClearAccum(gl15Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl15ClearColor(gl15Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl15ClearDepth(gl15Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl15ClearDepthf(gl15Context* glc, GLclampf depth) {
    return glc->fnClearDepthf(depth);
}

void gl15ClearIndex(gl15Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl15ClearStencil(gl15Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl15ClipPlane(gl15Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl15Color3b(gl15Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl15Color3d(gl15Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl15Color3f(gl15Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl15Color3i(gl15Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl15Color3s(gl15Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl15Color3ub(gl15Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl15Color3ui(gl15Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl15Color3us(gl15Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl15Color4b(gl15Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl15Color4d(gl15Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl15Color4f(gl15Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl15Color4i(gl15Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl15Color4s(gl15Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl15Color4ub(gl15Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl15Color4ui(gl15Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl15Color4us(gl15Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl15Color3bv(gl15Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl15Color3dv(gl15Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl15Color3fv(gl15Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl15Color3iv(gl15Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl15Color3sv(gl15Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl15Color3ubv(gl15Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl15Color3uiv(gl15Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl15Color3usv(gl15Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl15Color4bv(gl15Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl15Color4dv(gl15Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl15Color4fv(gl15Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl15Color4iv(gl15Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl15Color4sv(gl15Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl15Color4ubv(gl15Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl15Color4uiv(gl15Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl15Color4usv(gl15Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl15ColorMask(gl15Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl15ColorMaterial(gl15Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl15CopyPixels(gl15Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl15CullFace(gl15Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl15DeleteLists(gl15Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl15DepthFunc(gl15Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl15DepthMask(gl15Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl15DepthRange(gl15Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl15Enable(gl15Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl15Disable(gl15Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl15DrawBuffer(gl15Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl15DrawPixels(gl15Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl15EdgeFlag(gl15Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl15EdgeFlagv(gl15Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl15EdgeFlagPointer(gl15Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl15EvalCoord1d(gl15Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl15EvalCoord1f(gl15Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl15EvalCoord2d(gl15Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl15EvalCoord2f(gl15Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl15EvalCoord1dv(gl15Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl15EvalCoord1fv(gl15Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl15EvalCoord2dv(gl15Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl15EvalCoord2fv(gl15Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl15EvalMesh1(gl15Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl15EvalMesh2(gl15Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl15EvalPoint1(gl15Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl15EvalPoint2(gl15Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl15FeedbackBuffer(gl15Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl15Finish(gl15Context* glc) {
    return glc->fnFinish();
}

void gl15Flush(gl15Context* glc) {
    return glc->fnFlush();
}

void gl15Fogf(gl15Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl15Fogi(gl15Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl15Fogfv(gl15Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl15Fogiv(gl15Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl15FrontFace(gl15Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl15Frustum(gl15Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl15GenLists(gl15Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl15GetBooleanv(gl15Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl15GetDoublev(gl15Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl15GetFloatv(gl15Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl15GetIntegerv(gl15Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl15GetClipPlane(gl15Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl15GetError(gl15Context* glc) {
    return glc->fnGetError();
}

void gl15GetLightfv(gl15Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl15GetLightiv(gl15Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl15GetMapdv(gl15Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl15GetMapfv(gl15Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl15GetMapiv(gl15Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl15GetMaterialfv(gl15Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl15GetMaterialiv(gl15Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl15GetPixelMapfv(gl15Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl15GetPixelMapuiv(gl15Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl15GetPixelMapusv(gl15Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl15GetPolygonStipple(gl15Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl15GetString(gl15Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl15GetTexEnvfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl15GetTexEnviv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl15GetTexGendv(gl15Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl15GetTexGenfv(gl15Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl15GetTexGeniv(gl15Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl15GetTexImage(gl15Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl15GetTexLevelParameterfv(gl15Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl15GetTexLevelParameteriv(gl15Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl15GetTexParameterfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl15GetTexParameteriv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl15Hint(gl15Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl15Indexd(gl15Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl15Indexf(gl15Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl15Indexi(gl15Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl15Indexs(gl15Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl15Indexdv(gl15Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl15Indexfv(gl15Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl15Indexiv(gl15Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl15Indexsv(gl15Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl15IndexMask(gl15Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl15IndexPointer(gl15Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl15InitNames(gl15Context* glc) {
    return glc->fnInitNames();
}

void gl15IsEnabled(gl15Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl15IsList(gl15Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl15Lightf(gl15Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl15Lighti(gl15Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl15Lightfv(gl15Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl15Lightiv(gl15Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl15LightModelf(gl15Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl15LightModeli(gl15Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl15LightModelfv(gl15Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl15LightModeliv(gl15Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl15LineStipple(gl15Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl15LineWidth(gl15Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl15ListBase(gl15Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl15LoadIdentity(gl15Context* glc) {
    return glc->fnLoadIdentity();
}

void gl15LoadMatrixd(gl15Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl15LoadMatrixf(gl15Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl15LoadName(gl15Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl15LogicOp(gl15Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl15Map1d(gl15Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl15Map1f(gl15Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl15Map2d(gl15Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl15Map2f(gl15Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl15MapGrid1d(gl15Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl15MapGrid1f(gl15Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl15MapGrid2d(gl15Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl15MapGrid2f(gl15Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl15Materialf(gl15Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl15Materiali(gl15Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl15Materialfv(gl15Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl15Materialiv(gl15Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl15MatrixMode(gl15Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl15MultMatrixd(gl15Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl15MultMatrixf(gl15Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl15NewList(gl15Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl15EndList(gl15Context* glc) {
    return glc->fnEndList();
}

void gl15Normal3b(gl15Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl15Normal3d(gl15Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl15Normal3f(gl15Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl15Normal3i(gl15Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl15Normal3s(gl15Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl15Normal3bv(gl15Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl15Normal3dv(gl15Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl15Normal3fv(gl15Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl15Normal3iv(gl15Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl15Normal3sv(gl15Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl15Ortho(gl15Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl15PassThrough(gl15Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl15PixelMapfv(gl15Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl15PixelMapuiv(gl15Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl15PixelMapusv(gl15Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl15PixelStoref(gl15Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl15PixelStorei(gl15Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl15PixelTransferf(gl15Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl15PixelTransferi(gl15Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl15PixelZoom(gl15Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl15PointSize(gl15Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl15PolygonMode(gl15Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl15PolygonStipple(gl15Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl15PushAttrib(gl15Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl15PopAttrib(gl15Context* glc) {
    return glc->fnPopAttrib();
}

void gl15PushMatrix(gl15Context* glc) {
    return glc->fnPushMatrix();
}

void gl15PopMatrix(gl15Context* glc) {
    return glc->fnPopMatrix();
}

void gl15PushName(gl15Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl15PopName(gl15Context* glc) {
    return glc->fnPopName();
}

void gl15RasterPos2d(gl15Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl15RasterPos2f(gl15Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl15RasterPos2i(gl15Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl15RasterPos2s(gl15Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl15RasterPos3d(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl15RasterPos3f(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl15RasterPos3i(gl15Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl15RasterPos3s(gl15Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl15RasterPos4d(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl15RasterPos4f(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl15RasterPos4i(gl15Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl15RasterPos4s(gl15Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl15RasterPos2dv(gl15Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl15RasterPos2fv(gl15Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl15RasterPos2iv(gl15Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl15RasterPos2sv(gl15Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl15RasterPos3dv(gl15Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl15RasterPos3fv(gl15Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl15RasterPos3iv(gl15Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl15RasterPos3sv(gl15Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl15RasterPos4dv(gl15Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl15RasterPos4fv(gl15Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl15RasterPos4iv(gl15Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl15RasterPos4sv(gl15Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl15ReadBuffer(gl15Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl15ReadPixels(gl15Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl15Rectd(gl15Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl15Rectf(gl15Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl15Recti(gl15Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl15Rects(gl15Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl15Rectdv(gl15Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl15Rectfv(gl15Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl15Rectiv(gl15Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl15Rectsv(gl15Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl15RenderMode(gl15Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl15Rotated(gl15Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl15Rotatef(gl15Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl15Scaled(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl15Scalef(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl15Scissor(gl15Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl15SelectBuffer(gl15Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl15ShadeModel(gl15Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl15StencilFunc(gl15Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl15StencilMask(gl15Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl15StencilOp(gl15Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl15TexCoord1d(gl15Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl15TexCoord1f(gl15Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl15TexCoord1i(gl15Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl15TexCoord1s(gl15Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl15TexCoord2d(gl15Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl15TexCoord2f(gl15Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl15TexCoord2i(gl15Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl15TexCoord2s(gl15Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl15TexCoord3d(gl15Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl15TexCoord3f(gl15Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl15TexCoord3i(gl15Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl15TexCoord3s(gl15Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl15TexCoord4d(gl15Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl15TexCoord4f(gl15Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl15TexCoord4i(gl15Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl15TexCoord4s(gl15Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl15TexCoord1dv(gl15Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl15TexCoord1fv(gl15Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl15TexCoord1iv(gl15Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl15TexCoord1sv(gl15Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl15TexCoord2dv(gl15Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl15TexCoord2fv(gl15Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl15TexCoord2iv(gl15Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl15TexCoord2sv(gl15Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl15TexCoord3dv(gl15Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl15TexCoord3fv(gl15Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl15TexCoord3iv(gl15Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl15TexCoord3sv(gl15Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl15TexCoord4dv(gl15Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl15TexCoord4fv(gl15Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl15TexCoord4iv(gl15Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl15TexCoord4sv(gl15Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl15TexEnvf(gl15Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl15TexEnvi(gl15Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl15TexEnvfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl15TexEnviv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl15TexGend(gl15Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl15TexGenf(gl15Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl15TexGeni(gl15Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl15TexGendv(gl15Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl15TexGenfv(gl15Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl15TexGeniv(gl15Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl15TexImage1D(gl15Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl15TexImage2D(gl15Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl15TexParameterf(gl15Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl15TexParameteri(gl15Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl15TexParameterfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl15TexParameteriv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl15Translated(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl15Translatef(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl15Vertex2s(gl15Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl15Vertex2i(gl15Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl15Vertex2f(gl15Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl15Vertex2d(gl15Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl15Vertex3s(gl15Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl15Vertex3i(gl15Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl15Vertex3f(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl15Vertex3d(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl15Vertex4s(gl15Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl15Vertex4i(gl15Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl15Vertex4f(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl15Vertex4d(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl15Viewport(gl15Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl15GetConvolutionParameterfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl15GetConvolutionParameteriv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl15AreTexturesResident(gl15Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl15ArrayElement(gl15Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl15DrawArrays(gl15Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl15DrawElements(gl15Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl15GetPointerv(gl15Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl15PolygonOffset(gl15Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl15CopyTexImage1D(gl15Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl15CopyTexImage2D(gl15Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl15CopyTexSubImage1D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl15CopyTexSubImage2D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl15BindTexture(gl15Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl15DeleteTextures(gl15Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl15GenTextures(gl15Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl15IsTexture(gl15Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl15ColorPointer(gl15Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl15EnableClientState(gl15Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl15DisableClientState(gl15Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl15Indexub(gl15Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl15Indexubv(gl15Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl15InterleavedArrays(gl15Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl15NormalPointer(gl15Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl15PushClientAttrib(gl15Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl15PrioritizeTextures(gl15Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl15PopClientAttrib(gl15Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl15TexCoordPointer(gl15Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl15TexSubImage1D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl15TexSubImage2D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl15VertexPointer(gl15Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl15ColorTable(gl15Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl15ColorTableParameterfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl15ColorTableParameteriv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl15ColorSubTable(gl15Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl15ConvolutionFilter1D(gl15Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl15ConvolutionFilter2D(gl15Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl15ConvolutionParameterf(gl15Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl15ConvolutionParameteri(gl15Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl15CopyColorTable(gl15Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl15CopyColorSubTable(gl15Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl15CopyConvolutionFilter1D(gl15Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl15CopyConvolutionFilter2D(gl15Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl15GetColorTable(gl15Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl15GetColorTableParameterfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl15GetColorTableParameteriv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl15GetConvolutionFilter(gl15Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl15GetHistogram(gl15Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl15GetHistogramParameterfv(gl15Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl15GetHistogramParameteriv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl15GetSeparableFilter(gl15Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl15Histogram(gl15Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl15Minmax(gl15Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl15MultiTexCoord1s(gl15Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl15MultiTexCoord1i(gl15Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl15MultiTexCoord1f(gl15Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl15MultiTexCoord1d(gl15Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl15MultiTexCoord2s(gl15Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl15MultiTexCoord2i(gl15Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl15MultiTexCoord2f(gl15Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl15MultiTexCoord2d(gl15Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl15MultiTexCoord3s(gl15Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl15MultiTexCoord3i(gl15Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl15MultiTexCoord3f(gl15Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl15MultiTexCoord3d(gl15Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl15MultiTexCoord4s(gl15Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl15MultiTexCoord4i(gl15Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl15MultiTexCoord4f(gl15Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl15MultiTexCoord4d(gl15Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl15MultiTexCoord1sv(gl15Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl15MultiTexCoord1iv(gl15Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl15MultiTexCoord1fv(gl15Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl15MultiTexCoord1dv(gl15Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl15MultiTexCoord2sv(gl15Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl15MultiTexCoord2iv(gl15Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl15MultiTexCoord2fv(gl15Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl15MultiTexCoord2dv(gl15Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl15MultiTexCoord3sv(gl15Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl15MultiTexCoord3iv(gl15Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl15MultiTexCoord3fv(gl15Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl15MultiTexCoord3dv(gl15Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl15MultiTexCoord4sv(gl15Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl15MultiTexCoord4iv(gl15Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl15MultiTexCoord4fv(gl15Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl15MultiTexCoord4dv(gl15Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl15ResetHistogram(gl15Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl15ResetMinmax(gl15Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl15SeparableFilter2D(gl15Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl15BlendColor(gl15Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl15BlendEquation(gl15Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl15CopyTexSubImage3D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl15DrawRangeElements(gl15Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl15TexImage3D(gl15Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl15TexSubImage3D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl15ActiveTexture(gl15Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl15ClientActiveTexture(gl15Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl15CompressedTexImage1D(gl15Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl15CompressedTexImage2D(gl15Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl15CompressedTexImage3D(gl15Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl15CompressedTexSubImage1D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl15CompressedTexSubImage2D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl15CompressedTexSubImage3D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl15GetCompressedTexImage(gl15Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl15LoadTransposeMatrixd(gl15Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl15LoadTransposeMatrixf(gl15Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl15MultTransposeMatrixd(gl15Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl15MultTransposeMatrixf(gl15Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl15SampleCoverage(gl15Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl15BlendFuncSeparate(gl15Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl15FogCoordPointer(gl15Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl15FogCoordd(gl15Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl15FogCoordf(gl15Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl15FogCoorddv(gl15Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl15FogCoordfv(gl15Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl15MultiDrawArrays(gl15Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl15MultiDrawElements(gl15Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl15PointParameterf(gl15Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl15PointParameteri(gl15Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl15SecondaryColor3b(gl15Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl15SecondaryColor3s(gl15Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl15SecondaryColor3i(gl15Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl15SecondaryColor3f(gl15Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl15SecondaryColor3d(gl15Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl15SecondaryColor3ub(gl15Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl15SecondaryColor3us(gl15Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl15SecondaryColor3ui(gl15Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl15SecondaryColor3bv(gl15Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl15SecondaryColor3sv(gl15Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl15SecondaryColor3iv(gl15Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl15SecondaryColor3fv(gl15Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl15SecondaryColor3dv(gl15Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl15SecondaryColor3ubv(gl15Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl15SecondaryColor3usv(gl15Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl15SecondaryColor3uiv(gl15Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl15SecondaryColorPointer(gl15Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl15WindowPos2s(gl15Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl15WindowPos2i(gl15Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl15WindowPos2f(gl15Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl15WindowPos2d(gl15Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl15WindowPos3s(gl15Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl15WindowPos3i(gl15Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl15WindowPos3f(gl15Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl15WindowPos3d(gl15Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl15WindowPos2sv(gl15Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl15WindowPos2iv(gl15Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl15WindowPos2fv(gl15Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl15WindowPos2dv(gl15Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl15WindowPos3sv(gl15Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl15WindowPos3iv(gl15Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl15WindowPos3fv(gl15Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl15WindowPos3dv(gl15Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl15BeginQuery(gl15Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl15BindBuffer(gl15Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl15BufferData(gl15Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl15BufferSubData(gl15Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl15DeleteBuffers(gl15Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl15DeleteQueries(gl15Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl15GenBuffers(gl15Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl15GenQueries(gl15Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl15GetBufferParameteriv(gl15Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl15GetBufferPointerv(gl15Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl15GetBufferSubData(gl15Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl15GetQueryObjectiv(gl15Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl15GetQueryObjectuiv(gl15Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl15GetQueryiv(gl15Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl15IsBuffer(gl15Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl15IsQuery(gl15Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl15MapBuffer(gl15Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl15UnmapBuffer(gl15Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl15AttachShader(gl15Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl15BindAttribLocation(gl15Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl15BlendEquationSeperate(gl15Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl15CompileShader(gl15Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl15CreateProgram(gl15Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl15CreateShader(gl15Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl15DeleteProgram(gl15Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl15DeleteShader(gl15Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl15DetachShader(gl15Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl15EnableVertexAttribArray(gl15Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl15DisableVertexAttribArray(gl15Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl15DrawBuffers(gl15Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl15GetActiveAttrib(gl15Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl15GetActiveUniform(gl15Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl15GetAttachedShaders(gl15Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl15GetAttribLocation(gl15Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl15GetProgramiv(gl15Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl15GetProgramInfoLog(gl15Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl15GetShaderiv(gl15Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl15GetShaderInfoLog(gl15Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl15GetShaderSource(gl15Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl15GetUniformfv(gl15Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl15GetUniformiv(gl15Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl15GetUniformLocation(gl15Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl15GetVertexAttribdv(gl15Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl15GetVertexAttribfv(gl15Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl15GetVertexAttribiv(gl15Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl15GetVertexAttribPointerv(gl15Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl15IsProgram(gl15Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl15IsShader(gl15Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl15LinkProgram(gl15Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl15ShaderSource(gl15Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl15StencilFuncSeparate(gl15Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl15StencilMaskSeparate(gl15Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl15StencilOpSeparate(gl15Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl15Uniform1f(gl15Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl15Uniform2f(gl15Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl15Uniform3f(gl15Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl15Uniform4f(gl15Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl15Uniform1i(gl15Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl15Uniform2i(gl15Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl15Uniform3i(gl15Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl15Uniform4i(gl15Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl15Uniform1fv(gl15Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl15Uniform2fv(gl15Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl15Uniform3fv(gl15Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl15Uniform4fv(gl15Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl15Uniform1iv(gl15Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl15Uniform2iv(gl15Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl15Uniform3iv(gl15Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl15Uniform4iv(gl15Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl15UseProgram(gl15Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl15ValidateProgram(gl15Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl15VertexAttribPointer(gl15Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl15VertexAttrib1f(gl15Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl15VertexAttrib1s(gl15Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl15VertexAttrib1d(gl15Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl15VertexAttrib2f(gl15Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl15VertexAttrib2s(gl15Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl15VertexAttrib2d(gl15Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl15VertexAttrib3f(gl15Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl15VertexAttrib3s(gl15Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl15VertexAttrib3d(gl15Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl15VertexAttrib4f(gl15Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl15VertexAttrib4s(gl15Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl15VertexAttrib4d(gl15Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl15VertexAttrib4Nuv(gl15Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl15VertexAttrib1fv(gl15Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl15VertexAttrib1sv(gl15Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl15VertexAttrib1dv(gl15Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl15VertexAttrib2fv(gl15Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl15VertexAttrib2sv(gl15Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl15VertexAttrib2dv(gl15Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl15VertexAttrib3fv(gl15Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl15VertexAttrib3sv(gl15Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl15VertexAttrib3dv(gl15Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl15VertexAttrib4fv(gl15Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl15VertexAttrib4sv(gl15Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl15VertexAttrib4dv(gl15Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl15VertexAttrib4iv(gl15Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl15VertexAttrib4bv(gl15Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl15VertexAttrib4ubv(gl15Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl15VertexAttrib4usv(gl15Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl15VertexAttrib4uiv(gl15Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl15VertexAttrib4Nbv(gl15Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl15VertexAttrib4Nsv(gl15Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl15VertexAttrib4Niv(gl15Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl15VertexAttrib4Nubv(gl15Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl15VertexAttrib4Nusv(gl15Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl15VertexAttrib4Nuiv(gl15Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl15UniformMatrix2fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl15UniformMatrix3fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl15UniformMatrix4fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl15UniformMatrix2x3fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl15UniformMatrix3x2fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl15UniformMatrix2x4fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl15UniformMatrix4x2fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl15UniformMatrix3x4fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl15UniformMatrix4x3fv(gl15Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl15Context* gl15NewContext() {
    gl15Context* glc = calloc(1, sizeof(gl15Context));

    // Preload all procedures
    glc->fnAccum = (gl15PAccum)gl15LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl15PAlphaFunc)gl15LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl15PBegin)gl15LibGetProcAddress("glBegin");
    glc->fnEnd = (gl15PEnd)gl15LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl15PBitmap)gl15LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl15PBlendFunc)gl15LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl15PCallList)gl15LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl15PCallLists)gl15LibGetProcAddress("glCallLists");
    glc->fnClear = (gl15PClear)gl15LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl15PClearAccum)gl15LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl15PClearColor)gl15LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl15PClearDepth)gl15LibGetProcAddress("glClearDepth");
    glc->fnClearDepthf = (gl15PClearDepthf)gl15LibGetProcAddress("glClearDepthf");
    glc->fnClearIndex = (gl15PClearIndex)gl15LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl15PClearStencil)gl15LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl15PClipPlane)gl15LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl15PColor3b)gl15LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl15PColor3d)gl15LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl15PColor3f)gl15LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl15PColor3i)gl15LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl15PColor3s)gl15LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl15PColor3ub)gl15LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl15PColor3ui)gl15LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl15PColor3us)gl15LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl15PColor4b)gl15LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl15PColor4d)gl15LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl15PColor4f)gl15LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl15PColor4i)gl15LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl15PColor4s)gl15LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl15PColor4ub)gl15LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl15PColor4ui)gl15LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl15PColor4us)gl15LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl15PColor3bv)gl15LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl15PColor3dv)gl15LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl15PColor3fv)gl15LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl15PColor3iv)gl15LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl15PColor3sv)gl15LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl15PColor3ubv)gl15LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl15PColor3uiv)gl15LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl15PColor3usv)gl15LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl15PColor4bv)gl15LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl15PColor4dv)gl15LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl15PColor4fv)gl15LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl15PColor4iv)gl15LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl15PColor4sv)gl15LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl15PColor4ubv)gl15LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl15PColor4uiv)gl15LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl15PColor4usv)gl15LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl15PColorMask)gl15LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl15PColorMaterial)gl15LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl15PCopyPixels)gl15LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl15PCullFace)gl15LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl15PDeleteLists)gl15LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl15PDepthFunc)gl15LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl15PDepthMask)gl15LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl15PDepthRange)gl15LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl15PEnable)gl15LibGetProcAddress("glEnable");
    glc->fnDisable = (gl15PDisable)gl15LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl15PDrawBuffer)gl15LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl15PDrawPixels)gl15LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl15PEdgeFlag)gl15LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl15PEdgeFlagv)gl15LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl15PEdgeFlagPointer)gl15LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl15PEvalCoord1d)gl15LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl15PEvalCoord1f)gl15LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl15PEvalCoord2d)gl15LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl15PEvalCoord2f)gl15LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl15PEvalCoord1dv)gl15LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl15PEvalCoord1fv)gl15LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl15PEvalCoord2dv)gl15LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl15PEvalCoord2fv)gl15LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl15PEvalMesh1)gl15LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl15PEvalMesh2)gl15LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl15PEvalPoint1)gl15LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl15PEvalPoint2)gl15LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl15PFeedbackBuffer)gl15LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl15PFinish)gl15LibGetProcAddress("glFinish");
    glc->fnFlush = (gl15PFlush)gl15LibGetProcAddress("glFlush");
    glc->fnFogf = (gl15PFogf)gl15LibGetProcAddress("glFogf");
    glc->fnFogi = (gl15PFogi)gl15LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl15PFogfv)gl15LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl15PFogiv)gl15LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl15PFrontFace)gl15LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl15PFrustum)gl15LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl15PGenLists)gl15LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl15PGetBooleanv)gl15LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl15PGetDoublev)gl15LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl15PGetFloatv)gl15LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl15PGetIntegerv)gl15LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl15PGetClipPlane)gl15LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl15PGetError)gl15LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl15PGetLightfv)gl15LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl15PGetLightiv)gl15LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl15PGetMapdv)gl15LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl15PGetMapfv)gl15LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl15PGetMapiv)gl15LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl15PGetMaterialfv)gl15LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl15PGetMaterialiv)gl15LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl15PGetPixelMapfv)gl15LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl15PGetPixelMapuiv)gl15LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl15PGetPixelMapusv)gl15LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl15PGetPolygonStipple)gl15LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl15PGetString)gl15LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl15PGetTexEnvfv)gl15LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl15PGetTexEnviv)gl15LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl15PGetTexGendv)gl15LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl15PGetTexGenfv)gl15LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl15PGetTexGeniv)gl15LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl15PGetTexImage)gl15LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl15PGetTexLevelParameterfv)gl15LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl15PGetTexLevelParameteriv)gl15LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl15PGetTexParameterfv)gl15LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl15PGetTexParameteriv)gl15LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl15PHint)gl15LibGetProcAddress("glHint");
    glc->fnIndexd = (gl15PIndexd)gl15LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl15PIndexf)gl15LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl15PIndexi)gl15LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl15PIndexs)gl15LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl15PIndexdv)gl15LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl15PIndexfv)gl15LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl15PIndexiv)gl15LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl15PIndexsv)gl15LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl15PIndexMask)gl15LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl15PIndexPointer)gl15LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl15PInitNames)gl15LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl15PIsEnabled)gl15LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl15PIsList)gl15LibGetProcAddress("glIsList");
    glc->fnLightf = (gl15PLightf)gl15LibGetProcAddress("glLightf");
    glc->fnLighti = (gl15PLighti)gl15LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl15PLightfv)gl15LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl15PLightiv)gl15LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl15PLightModelf)gl15LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl15PLightModeli)gl15LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl15PLightModelfv)gl15LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl15PLightModeliv)gl15LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl15PLineStipple)gl15LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl15PLineWidth)gl15LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl15PListBase)gl15LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl15PLoadIdentity)gl15LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl15PLoadMatrixd)gl15LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl15PLoadMatrixf)gl15LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl15PLoadName)gl15LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl15PLogicOp)gl15LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl15PMap1d)gl15LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl15PMap1f)gl15LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl15PMap2d)gl15LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl15PMap2f)gl15LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl15PMapGrid1d)gl15LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl15PMapGrid1f)gl15LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl15PMapGrid2d)gl15LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl15PMapGrid2f)gl15LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl15PMaterialf)gl15LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl15PMateriali)gl15LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl15PMaterialfv)gl15LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl15PMaterialiv)gl15LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl15PMatrixMode)gl15LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl15PMultMatrixd)gl15LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl15PMultMatrixf)gl15LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl15PNewList)gl15LibGetProcAddress("glNewList");
    glc->fnEndList = (gl15PEndList)gl15LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl15PNormal3b)gl15LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl15PNormal3d)gl15LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl15PNormal3f)gl15LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl15PNormal3i)gl15LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl15PNormal3s)gl15LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl15PNormal3bv)gl15LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl15PNormal3dv)gl15LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl15PNormal3fv)gl15LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl15PNormal3iv)gl15LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl15PNormal3sv)gl15LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl15POrtho)gl15LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl15PPassThrough)gl15LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl15PPixelMapfv)gl15LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl15PPixelMapuiv)gl15LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl15PPixelMapusv)gl15LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl15PPixelStoref)gl15LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl15PPixelStorei)gl15LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl15PPixelTransferf)gl15LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl15PPixelTransferi)gl15LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl15PPixelZoom)gl15LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl15PPointSize)gl15LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl15PPolygonMode)gl15LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl15PPolygonStipple)gl15LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl15PPushAttrib)gl15LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl15PPopAttrib)gl15LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl15PPushMatrix)gl15LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl15PPopMatrix)gl15LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl15PPushName)gl15LibGetProcAddress("glPushName");
    glc->fnPopName = (gl15PPopName)gl15LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl15PRasterPos2d)gl15LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl15PRasterPos2f)gl15LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl15PRasterPos2i)gl15LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl15PRasterPos2s)gl15LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl15PRasterPos3d)gl15LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl15PRasterPos3f)gl15LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl15PRasterPos3i)gl15LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl15PRasterPos3s)gl15LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl15PRasterPos4d)gl15LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl15PRasterPos4f)gl15LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl15PRasterPos4i)gl15LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl15PRasterPos4s)gl15LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl15PRasterPos2dv)gl15LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl15PRasterPos2fv)gl15LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl15PRasterPos2iv)gl15LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl15PRasterPos2sv)gl15LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl15PRasterPos3dv)gl15LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl15PRasterPos3fv)gl15LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl15PRasterPos3iv)gl15LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl15PRasterPos3sv)gl15LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl15PRasterPos4dv)gl15LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl15PRasterPos4fv)gl15LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl15PRasterPos4iv)gl15LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl15PRasterPos4sv)gl15LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl15PReadBuffer)gl15LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl15PReadPixels)gl15LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl15PRectd)gl15LibGetProcAddress("glRectd");
    glc->fnRectf = (gl15PRectf)gl15LibGetProcAddress("glRectf");
    glc->fnRecti = (gl15PRecti)gl15LibGetProcAddress("glRecti");
    glc->fnRects = (gl15PRects)gl15LibGetProcAddress("glRects");
    glc->fnRectdv = (gl15PRectdv)gl15LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl15PRectfv)gl15LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl15PRectiv)gl15LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl15PRectsv)gl15LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl15PRenderMode)gl15LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl15PRotated)gl15LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl15PRotatef)gl15LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl15PScaled)gl15LibGetProcAddress("glScaled");
    glc->fnScalef = (gl15PScalef)gl15LibGetProcAddress("glScalef");
    glc->fnScissor = (gl15PScissor)gl15LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl15PSelectBuffer)gl15LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl15PShadeModel)gl15LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl15PStencilFunc)gl15LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl15PStencilMask)gl15LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl15PStencilOp)gl15LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl15PTexCoord1d)gl15LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl15PTexCoord1f)gl15LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl15PTexCoord1i)gl15LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl15PTexCoord1s)gl15LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl15PTexCoord2d)gl15LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl15PTexCoord2f)gl15LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl15PTexCoord2i)gl15LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl15PTexCoord2s)gl15LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl15PTexCoord3d)gl15LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl15PTexCoord3f)gl15LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl15PTexCoord3i)gl15LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl15PTexCoord3s)gl15LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl15PTexCoord4d)gl15LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl15PTexCoord4f)gl15LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl15PTexCoord4i)gl15LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl15PTexCoord4s)gl15LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl15PTexCoord1dv)gl15LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl15PTexCoord1fv)gl15LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl15PTexCoord1iv)gl15LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl15PTexCoord1sv)gl15LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl15PTexCoord2dv)gl15LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl15PTexCoord2fv)gl15LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl15PTexCoord2iv)gl15LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl15PTexCoord2sv)gl15LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl15PTexCoord3dv)gl15LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl15PTexCoord3fv)gl15LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl15PTexCoord3iv)gl15LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl15PTexCoord3sv)gl15LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl15PTexCoord4dv)gl15LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl15PTexCoord4fv)gl15LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl15PTexCoord4iv)gl15LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl15PTexCoord4sv)gl15LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl15PTexEnvf)gl15LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl15PTexEnvi)gl15LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl15PTexEnvfv)gl15LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl15PTexEnviv)gl15LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl15PTexGend)gl15LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl15PTexGenf)gl15LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl15PTexGeni)gl15LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl15PTexGendv)gl15LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl15PTexGenfv)gl15LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl15PTexGeniv)gl15LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl15PTexImage1D)gl15LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl15PTexImage2D)gl15LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl15PTexParameterf)gl15LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl15PTexParameteri)gl15LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl15PTexParameterfv)gl15LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl15PTexParameteriv)gl15LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl15PTranslated)gl15LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl15PTranslatef)gl15LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl15PVertex2s)gl15LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl15PVertex2i)gl15LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl15PVertex2f)gl15LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl15PVertex2d)gl15LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl15PVertex3s)gl15LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl15PVertex3i)gl15LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl15PVertex3f)gl15LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl15PVertex3d)gl15LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl15PVertex4s)gl15LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl15PVertex4i)gl15LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl15PVertex4f)gl15LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl15PVertex4d)gl15LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl15PViewport)gl15LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl15PGetConvolutionParameterfv)gl15LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl15PGetConvolutionParameteriv)gl15LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl15PAreTexturesResident)gl15LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl15PArrayElement)gl15LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl15PDrawArrays)gl15LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl15PDrawElements)gl15LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl15PGetPointerv)gl15LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl15PPolygonOffset)gl15LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl15PCopyTexImage1D)gl15LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl15PCopyTexImage2D)gl15LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl15PCopyTexSubImage1D)gl15LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl15PCopyTexSubImage2D)gl15LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl15PBindTexture)gl15LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl15PDeleteTextures)gl15LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl15PGenTextures)gl15LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl15PIsTexture)gl15LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl15PColorPointer)gl15LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl15PEnableClientState)gl15LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl15PDisableClientState)gl15LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl15PIndexub)gl15LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl15PIndexubv)gl15LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl15PInterleavedArrays)gl15LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl15PNormalPointer)gl15LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl15PPushClientAttrib)gl15LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl15PPrioritizeTextures)gl15LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl15PPopClientAttrib)gl15LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl15PTexCoordPointer)gl15LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl15PTexSubImage1D)gl15LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl15PTexSubImage2D)gl15LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl15PVertexPointer)gl15LibGetProcAddress("glVertexPointer");
    glc->fnColorTable = (gl15PColorTable)gl15GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl15PColorTableParameterfv)gl15GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl15PColorTableParameteriv)gl15GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl15PColorSubTable)gl15GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl15PConvolutionFilter1D)gl15GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl15PConvolutionFilter2D)gl15GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl15PConvolutionParameterf)gl15GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl15PConvolutionParameteri)gl15GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl15PCopyColorTable)gl15GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl15PCopyColorSubTable)gl15GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl15PCopyConvolutionFilter1D)gl15GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl15PCopyConvolutionFilter2D)gl15GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl15PGetColorTable)gl15GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl15PGetColorTableParameterfv)gl15GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl15PGetColorTableParameteriv)gl15GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl15PGetConvolutionFilter)gl15GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl15PGetHistogram)gl15GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl15PGetHistogramParameterfv)gl15GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl15PGetHistogramParameteriv)gl15GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl15PGetSeparableFilter)gl15GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl15PHistogram)gl15GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl15PMinmax)gl15GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl15PMultiTexCoord1s)gl15GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl15PMultiTexCoord1i)gl15GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl15PMultiTexCoord1f)gl15GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl15PMultiTexCoord1d)gl15GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl15PMultiTexCoord2s)gl15GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl15PMultiTexCoord2i)gl15GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl15PMultiTexCoord2f)gl15GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl15PMultiTexCoord2d)gl15GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl15PMultiTexCoord3s)gl15GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl15PMultiTexCoord3i)gl15GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl15PMultiTexCoord3f)gl15GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl15PMultiTexCoord3d)gl15GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl15PMultiTexCoord4s)gl15GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl15PMultiTexCoord4i)gl15GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl15PMultiTexCoord4f)gl15GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl15PMultiTexCoord4d)gl15GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl15PMultiTexCoord1sv)gl15GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl15PMultiTexCoord1iv)gl15GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl15PMultiTexCoord1fv)gl15GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl15PMultiTexCoord1dv)gl15GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl15PMultiTexCoord2sv)gl15GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl15PMultiTexCoord2iv)gl15GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl15PMultiTexCoord2fv)gl15GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl15PMultiTexCoord2dv)gl15GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl15PMultiTexCoord3sv)gl15GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl15PMultiTexCoord3iv)gl15GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl15PMultiTexCoord3fv)gl15GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl15PMultiTexCoord3dv)gl15GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl15PMultiTexCoord4sv)gl15GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl15PMultiTexCoord4iv)gl15GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl15PMultiTexCoord4fv)gl15GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl15PMultiTexCoord4dv)gl15GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl15PResetHistogram)gl15GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl15PResetMinmax)gl15GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl15PSeparableFilter2D)gl15GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl15PBlendColor)gl15GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl15PBlendEquation)gl15GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl15PCopyTexSubImage3D)gl15GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl15PDrawRangeElements)gl15GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl15PTexImage3D)gl15GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl15PTexSubImage3D)gl15GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl15PActiveTexture)gl15GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl15PClientActiveTexture)gl15GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl15PCompressedTexImage1D)gl15GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl15PCompressedTexImage2D)gl15GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl15PCompressedTexImage3D)gl15GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl15PCompressedTexSubImage1D)gl15GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl15PCompressedTexSubImage2D)gl15GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl15PCompressedTexSubImage3D)gl15GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl15PGetCompressedTexImage)gl15GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl15PLoadTransposeMatrixd)gl15GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl15PLoadTransposeMatrixf)gl15GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl15PMultTransposeMatrixd)gl15GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl15PMultTransposeMatrixf)gl15GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl15PSampleCoverage)gl15GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl15PBlendFuncSeparate)gl15GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl15PFogCoordPointer)gl15GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl15PFogCoordd)gl15GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl15PFogCoordf)gl15GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl15PFogCoorddv)gl15GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl15PFogCoordfv)gl15GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl15PMultiDrawArrays)gl15GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl15PMultiDrawElements)gl15GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl15PPointParameterf)gl15GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl15PPointParameteri)gl15GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl15PSecondaryColor3b)gl15GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl15PSecondaryColor3s)gl15GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl15PSecondaryColor3i)gl15GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl15PSecondaryColor3f)gl15GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl15PSecondaryColor3d)gl15GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl15PSecondaryColor3ub)gl15GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl15PSecondaryColor3us)gl15GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl15PSecondaryColor3ui)gl15GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl15PSecondaryColor3bv)gl15GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl15PSecondaryColor3sv)gl15GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl15PSecondaryColor3iv)gl15GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl15PSecondaryColor3fv)gl15GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl15PSecondaryColor3dv)gl15GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl15PSecondaryColor3ubv)gl15GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl15PSecondaryColor3usv)gl15GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl15PSecondaryColor3uiv)gl15GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl15PSecondaryColorPointer)gl15GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl15PWindowPos2s)gl15GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl15PWindowPos2i)gl15GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl15PWindowPos2f)gl15GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl15PWindowPos2d)gl15GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl15PWindowPos3s)gl15GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl15PWindowPos3i)gl15GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl15PWindowPos3f)gl15GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl15PWindowPos3d)gl15GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl15PWindowPos2sv)gl15GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl15PWindowPos2iv)gl15GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl15PWindowPos2fv)gl15GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl15PWindowPos2dv)gl15GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl15PWindowPos3sv)gl15GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl15PWindowPos3iv)gl15GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl15PWindowPos3fv)gl15GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl15PWindowPos3dv)gl15GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl15PBeginQuery)gl15GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl15PBindBuffer)gl15GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl15PBufferData)gl15GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl15PBufferSubData)gl15GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl15PDeleteBuffers)gl15GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl15PDeleteQueries)gl15GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl15PGenBuffers)gl15GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl15PGenQueries)gl15GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl15PGetBufferParameteriv)gl15GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl15PGetBufferPointerv)gl15GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl15PGetBufferSubData)gl15GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl15PGetQueryObjectiv)gl15GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl15PGetQueryObjectuiv)gl15GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl15PGetQueryiv)gl15GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl15PIsBuffer)gl15GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl15PIsQuery)gl15GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl15PMapBuffer)gl15GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl15PUnmapBuffer)gl15GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl15PAttachShader)gl15GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl15PBindAttribLocation)gl15GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl15PBlendEquationSeperate)gl15GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl15PCompileShader)gl15GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl15PCreateProgram)gl15GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl15PCreateShader)gl15GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl15PDeleteProgram)gl15GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl15PDeleteShader)gl15GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl15PDetachShader)gl15GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl15PEnableVertexAttribArray)gl15GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl15PDisableVertexAttribArray)gl15GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl15PDrawBuffers)gl15LibGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl15PGetActiveAttrib)gl15LibGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl15PGetActiveUniform)gl15LibGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl15PGetAttachedShaders)gl15LibGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl15PGetAttribLocation)gl15LibGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl15PGetProgramiv)gl15LibGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl15PGetProgramInfoLog)gl15LibGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl15PGetShaderiv)gl15LibGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl15PGetShaderInfoLog)gl15LibGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl15PGetShaderSource)gl15LibGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl15PGetUniformfv)gl15LibGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl15PGetUniformiv)gl15LibGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl15PGetUniformLocation)gl15LibGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl15PGetVertexAttribdv)gl15LibGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl15PGetVertexAttribfv)gl15LibGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl15PGetVertexAttribiv)gl15LibGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl15PGetVertexAttribPointerv)gl15LibGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl15PIsProgram)gl15LibGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl15PIsShader)gl15LibGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl15PLinkProgram)gl15LibGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl15PShaderSource)gl15LibGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl15PStencilFuncSeparate)gl15LibGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl15PStencilMaskSeparate)gl15LibGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl15PStencilOpSeparate)gl15LibGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl15PUniform1f)gl15LibGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl15PUniform2f)gl15LibGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl15PUniform3f)gl15LibGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl15PUniform4f)gl15LibGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl15PUniform1i)gl15LibGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl15PUniform2i)gl15LibGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl15PUniform3i)gl15LibGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl15PUniform4i)gl15LibGetProcAddress("glUniform4i");
    glc->fnUniform1fv = (gl15PUniform1fv)gl15LibGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl15PUniform2fv)gl15LibGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl15PUniform3fv)gl15LibGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl15PUniform4fv)gl15LibGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl15PUniform1iv)gl15LibGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl15PUniform2iv)gl15LibGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl15PUniform3iv)gl15LibGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl15PUniform4iv)gl15LibGetProcAddress("glUniform4iv");
    glc->fnUseProgram = (gl15PUseProgram)gl15LibGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl15PValidateProgram)gl15LibGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl15PVertexAttribPointer)gl15LibGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl15PVertexAttrib1f)gl15LibGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl15PVertexAttrib1s)gl15LibGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl15PVertexAttrib1d)gl15LibGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl15PVertexAttrib2f)gl15LibGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl15PVertexAttrib2s)gl15LibGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl15PVertexAttrib2d)gl15LibGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl15PVertexAttrib3f)gl15LibGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl15PVertexAttrib3s)gl15LibGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl15PVertexAttrib3d)gl15LibGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl15PVertexAttrib4f)gl15LibGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl15PVertexAttrib4s)gl15LibGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl15PVertexAttrib4d)gl15LibGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl15PVertexAttrib4Nuv)gl15LibGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl15PVertexAttrib1fv)gl15LibGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl15PVertexAttrib1sv)gl15LibGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl15PVertexAttrib1dv)gl15LibGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl15PVertexAttrib2fv)gl15LibGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl15PVertexAttrib2sv)gl15LibGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl15PVertexAttrib2dv)gl15LibGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl15PVertexAttrib3fv)gl15LibGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl15PVertexAttrib3sv)gl15LibGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl15PVertexAttrib3dv)gl15LibGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl15PVertexAttrib4fv)gl15LibGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl15PVertexAttrib4sv)gl15LibGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl15PVertexAttrib4dv)gl15LibGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl15PVertexAttrib4iv)gl15LibGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl15PVertexAttrib4bv)gl15LibGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl15PVertexAttrib4ubv)gl15LibGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl15PVertexAttrib4usv)gl15LibGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl15PVertexAttrib4uiv)gl15LibGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl15PVertexAttrib4Nbv)gl15LibGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl15PVertexAttrib4Nsv)gl15LibGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl15PVertexAttrib4Niv)gl15LibGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl15PVertexAttrib4Nubv)gl15LibGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl15PVertexAttrib4Nusv)gl15LibGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl15PVertexAttrib4Nuiv)gl15LibGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl15PUniformMatrix2fv)gl15LibGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl15PUniformMatrix3fv)gl15LibGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl15PUniformMatrix4fv)gl15LibGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl15PUniformMatrix2x3fv)gl15LibGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl15PUniformMatrix3x2fv)gl15LibGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl15PUniformMatrix2x4fv)gl15LibGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl15PUniformMatrix4x2fv)gl15LibGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl15PUniformMatrix3x4fv)gl15LibGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl15PUniformMatrix4x3fv)gl15LibGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

