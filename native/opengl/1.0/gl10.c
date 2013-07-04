
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl10.h"

#ifdef _WIN32
HMODULE gl10OpenGL32;

void* gl10LibGetProcAddress(char* name) {
	if(gl10OpenGL32 == NULL) {
		gl10OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl10OpenGL32, TEXT(name));
}

void* gl10GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl10Accum(gl10Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl10AlphaFunc(gl10Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl10Begin(gl10Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl10End(gl10Context* glc) {
    return glc->fnEnd();
}

void gl10Bitmap(gl10Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl10BlendFunc(gl10Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl10CallList(gl10Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl10CallLists(gl10Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl10Clear(gl10Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl10ClearAccum(gl10Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl10ClearColor(gl10Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl10ClearDepth(gl10Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl10ClearDepthf(gl10Context* glc, GLclampf depth) {
    return glc->fnClearDepthf(depth);
}

void gl10ClearIndex(gl10Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl10ClearStencil(gl10Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl10ClipPlane(gl10Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl10Color3b(gl10Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl10Color3d(gl10Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl10Color3f(gl10Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl10Color3i(gl10Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl10Color3s(gl10Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl10Color3ub(gl10Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl10Color3ui(gl10Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl10Color3us(gl10Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl10Color4b(gl10Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl10Color4d(gl10Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl10Color4f(gl10Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl10Color4i(gl10Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl10Color4s(gl10Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl10Color4ub(gl10Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl10Color4ui(gl10Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl10Color4us(gl10Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl10Color3bv(gl10Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl10Color3dv(gl10Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl10Color3fv(gl10Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl10Color3iv(gl10Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl10Color3sv(gl10Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl10Color3ubv(gl10Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl10Color3uiv(gl10Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl10Color3usv(gl10Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl10Color4bv(gl10Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl10Color4dv(gl10Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl10Color4fv(gl10Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl10Color4iv(gl10Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl10Color4sv(gl10Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl10Color4ubv(gl10Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl10Color4uiv(gl10Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl10Color4usv(gl10Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl10ColorMask(gl10Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl10ColorMaterial(gl10Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl10CopyPixels(gl10Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl10CullFace(gl10Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl10DeleteLists(gl10Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl10DepthFunc(gl10Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl10DepthMask(gl10Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl10DepthRange(gl10Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl10Enable(gl10Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl10Disable(gl10Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl10DrawBuffer(gl10Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl10DrawPixels(gl10Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl10EdgeFlag(gl10Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl10EdgeFlagv(gl10Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl10EdgeFlagPointer(gl10Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl10EvalCoord1d(gl10Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl10EvalCoord1f(gl10Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl10EvalCoord2d(gl10Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl10EvalCoord2f(gl10Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl10EvalCoord1dv(gl10Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl10EvalCoord1fv(gl10Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl10EvalCoord2dv(gl10Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl10EvalCoord2fv(gl10Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl10EvalMesh1(gl10Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl10EvalMesh2(gl10Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl10EvalPoint1(gl10Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl10EvalPoint2(gl10Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl10FeedbackBuffer(gl10Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl10Finish(gl10Context* glc) {
    return glc->fnFinish();
}

void gl10Flush(gl10Context* glc) {
    return glc->fnFlush();
}

void gl10Fogf(gl10Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl10Fogi(gl10Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl10Fogfv(gl10Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl10Fogiv(gl10Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl10FrontFace(gl10Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl10Frustum(gl10Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl10GenLists(gl10Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl10GetBooleanv(gl10Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl10GetDoublev(gl10Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl10GetFloatv(gl10Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl10GetIntegerv(gl10Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl10GetClipPlane(gl10Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl10GetError(gl10Context* glc) {
    return glc->fnGetError();
}

void gl10GetLightfv(gl10Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl10GetLightiv(gl10Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl10GetMapdv(gl10Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl10GetMapfv(gl10Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl10GetMapiv(gl10Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl10GetMaterialfv(gl10Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl10GetMaterialiv(gl10Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl10GetPixelMapfv(gl10Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl10GetPixelMapuiv(gl10Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl10GetPixelMapusv(gl10Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl10GetPolygonStipple(gl10Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl10GetString(gl10Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl10GetTexEnvfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl10GetTexEnviv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl10GetTexGendv(gl10Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl10GetTexGenfv(gl10Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl10GetTexGeniv(gl10Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl10GetTexImage(gl10Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl10GetTexLevelParameterfv(gl10Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl10GetTexLevelParameteriv(gl10Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl10GetTexParameterfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl10GetTexParameteriv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl10Hint(gl10Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl10Indexd(gl10Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl10Indexf(gl10Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl10Indexi(gl10Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl10Indexs(gl10Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl10Indexdv(gl10Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl10Indexfv(gl10Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl10Indexiv(gl10Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl10Indexsv(gl10Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl10IndexMask(gl10Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl10IndexPointer(gl10Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl10InitNames(gl10Context* glc) {
    return glc->fnInitNames();
}

void gl10IsEnabled(gl10Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl10IsList(gl10Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl10Lightf(gl10Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl10Lighti(gl10Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl10Lightfv(gl10Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl10Lightiv(gl10Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl10LightModelf(gl10Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl10LightModeli(gl10Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl10LightModelfv(gl10Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl10LightModeliv(gl10Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl10LineStipple(gl10Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl10LineWidth(gl10Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl10ListBase(gl10Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl10LoadIdentity(gl10Context* glc) {
    return glc->fnLoadIdentity();
}

void gl10LoadMatrixd(gl10Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl10LoadMatrixf(gl10Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl10LoadName(gl10Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl10LogicOp(gl10Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl10Map1d(gl10Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl10Map1f(gl10Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl10Map2d(gl10Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl10Map2f(gl10Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl10MapGrid1d(gl10Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl10MapGrid1f(gl10Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl10MapGrid2d(gl10Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl10MapGrid2f(gl10Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl10Materialf(gl10Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl10Materiali(gl10Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl10Materialfv(gl10Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl10Materialiv(gl10Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl10MatrixMode(gl10Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl10MultMatrixd(gl10Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl10MultMatrixf(gl10Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl10NewList(gl10Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl10EndList(gl10Context* glc) {
    return glc->fnEndList();
}

void gl10Normal3b(gl10Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl10Normal3d(gl10Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl10Normal3f(gl10Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl10Normal3i(gl10Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl10Normal3s(gl10Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl10Normal3bv(gl10Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl10Normal3dv(gl10Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl10Normal3fv(gl10Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl10Normal3iv(gl10Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl10Normal3sv(gl10Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl10Ortho(gl10Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl10PassThrough(gl10Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl10PixelMapfv(gl10Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl10PixelMapuiv(gl10Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl10PixelMapusv(gl10Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl10PixelStoref(gl10Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl10PixelStorei(gl10Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl10PixelTransferf(gl10Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl10PixelTransferi(gl10Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl10PixelZoom(gl10Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl10PointSize(gl10Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl10PolygonMode(gl10Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl10PolygonStipple(gl10Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl10PushAttrib(gl10Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl10PopAttrib(gl10Context* glc) {
    return glc->fnPopAttrib();
}

void gl10PushMatrix(gl10Context* glc) {
    return glc->fnPushMatrix();
}

void gl10PopMatrix(gl10Context* glc) {
    return glc->fnPopMatrix();
}

void gl10PushName(gl10Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl10PopName(gl10Context* glc) {
    return glc->fnPopName();
}

void gl10RasterPos2d(gl10Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl10RasterPos2f(gl10Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl10RasterPos2i(gl10Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl10RasterPos2s(gl10Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl10RasterPos3d(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl10RasterPos3f(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl10RasterPos3i(gl10Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl10RasterPos3s(gl10Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl10RasterPos4d(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl10RasterPos4f(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl10RasterPos4i(gl10Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl10RasterPos4s(gl10Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl10RasterPos2dv(gl10Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl10RasterPos2fv(gl10Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl10RasterPos2iv(gl10Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl10RasterPos2sv(gl10Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl10RasterPos3dv(gl10Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl10RasterPos3fv(gl10Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl10RasterPos3iv(gl10Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl10RasterPos3sv(gl10Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl10RasterPos4dv(gl10Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl10RasterPos4fv(gl10Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl10RasterPos4iv(gl10Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl10RasterPos4sv(gl10Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl10ReadBuffer(gl10Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl10ReadPixels(gl10Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl10Rectd(gl10Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl10Rectf(gl10Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl10Recti(gl10Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl10Rects(gl10Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl10Rectdv(gl10Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl10Rectfv(gl10Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl10Rectiv(gl10Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl10Rectsv(gl10Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl10RenderMode(gl10Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl10Rotated(gl10Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl10Rotatef(gl10Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl10Scaled(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl10Scalef(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl10Scissor(gl10Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl10SelectBuffer(gl10Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl10ShadeModel(gl10Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl10StencilFunc(gl10Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl10StencilMask(gl10Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl10StencilOp(gl10Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl10TexCoord1d(gl10Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl10TexCoord1f(gl10Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl10TexCoord1i(gl10Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl10TexCoord1s(gl10Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl10TexCoord2d(gl10Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl10TexCoord2f(gl10Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl10TexCoord2i(gl10Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl10TexCoord2s(gl10Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl10TexCoord3d(gl10Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl10TexCoord3f(gl10Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl10TexCoord3i(gl10Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl10TexCoord3s(gl10Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl10TexCoord4d(gl10Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl10TexCoord4f(gl10Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl10TexCoord4i(gl10Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl10TexCoord4s(gl10Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl10TexCoord1dv(gl10Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl10TexCoord1fv(gl10Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl10TexCoord1iv(gl10Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl10TexCoord1sv(gl10Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl10TexCoord2dv(gl10Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl10TexCoord2fv(gl10Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl10TexCoord2iv(gl10Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl10TexCoord2sv(gl10Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl10TexCoord3dv(gl10Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl10TexCoord3fv(gl10Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl10TexCoord3iv(gl10Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl10TexCoord3sv(gl10Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl10TexCoord4dv(gl10Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl10TexCoord4fv(gl10Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl10TexCoord4iv(gl10Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl10TexCoord4sv(gl10Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl10TexEnvf(gl10Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl10TexEnvi(gl10Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl10TexEnvfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl10TexEnviv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl10TexGend(gl10Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl10TexGenf(gl10Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl10TexGeni(gl10Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl10TexGendv(gl10Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl10TexGenfv(gl10Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl10TexGeniv(gl10Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl10TexImage1D(gl10Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl10TexImage2D(gl10Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl10TexParameterf(gl10Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl10TexParameteri(gl10Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl10TexParameterfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl10TexParameteriv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl10Translated(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl10Translatef(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl10Vertex2s(gl10Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl10Vertex2i(gl10Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl10Vertex2f(gl10Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl10Vertex2d(gl10Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl10Vertex3s(gl10Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl10Vertex3i(gl10Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl10Vertex3f(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl10Vertex3d(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl10Vertex4s(gl10Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl10Vertex4i(gl10Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl10Vertex4f(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl10Vertex4d(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl10Viewport(gl10Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl10GetConvolutionParameterfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl10GetConvolutionParameteriv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

void gl10ColorTable(gl10Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl10ColorTableParameterfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl10ColorTableParameteriv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl10ColorSubTable(gl10Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl10ConvolutionFilter1D(gl10Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl10ConvolutionFilter2D(gl10Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl10ConvolutionParameterf(gl10Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl10ConvolutionParameteri(gl10Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl10CopyColorTable(gl10Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl10CopyColorSubTable(gl10Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl10CopyConvolutionFilter1D(gl10Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl10CopyConvolutionFilter2D(gl10Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl10GetColorTable(gl10Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl10GetColorTableParameterfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl10GetColorTableParameteriv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl10GetConvolutionFilter(gl10Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl10GetHistogram(gl10Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl10GetHistogramParameterfv(gl10Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl10GetHistogramParameteriv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl10GetSeparableFilter(gl10Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl10Histogram(gl10Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl10Minmax(gl10Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl10MultiTexCoord1s(gl10Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl10MultiTexCoord1i(gl10Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl10MultiTexCoord1f(gl10Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl10MultiTexCoord1d(gl10Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl10MultiTexCoord2s(gl10Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl10MultiTexCoord2i(gl10Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl10MultiTexCoord2f(gl10Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl10MultiTexCoord2d(gl10Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl10MultiTexCoord3s(gl10Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl10MultiTexCoord3i(gl10Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl10MultiTexCoord3f(gl10Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl10MultiTexCoord3d(gl10Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl10MultiTexCoord4s(gl10Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl10MultiTexCoord4i(gl10Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl10MultiTexCoord4f(gl10Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl10MultiTexCoord4d(gl10Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl10MultiTexCoord1sv(gl10Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl10MultiTexCoord1iv(gl10Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl10MultiTexCoord1fv(gl10Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl10MultiTexCoord1dv(gl10Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl10MultiTexCoord2sv(gl10Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl10MultiTexCoord2iv(gl10Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl10MultiTexCoord2fv(gl10Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl10MultiTexCoord2dv(gl10Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl10MultiTexCoord3sv(gl10Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl10MultiTexCoord3iv(gl10Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl10MultiTexCoord3fv(gl10Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl10MultiTexCoord3dv(gl10Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl10MultiTexCoord4sv(gl10Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl10MultiTexCoord4iv(gl10Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl10MultiTexCoord4fv(gl10Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl10MultiTexCoord4dv(gl10Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl10ResetHistogram(gl10Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl10ResetMinmax(gl10Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl10SeparableFilter2D(gl10Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

GLboolean gl10AreTexturesResident(gl10Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl10ArrayElement(gl10Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl10DrawArrays(gl10Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl10DrawElements(gl10Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl10GetPointerv(gl10Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl10PolygonOffset(gl10Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl10CopyTexImage1D(gl10Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl10CopyTexImage2D(gl10Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl10CopyTexSubImage1D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl10CopyTexSubImage2D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl10BindTexture(gl10Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl10DeleteTextures(gl10Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl10GenTextures(gl10Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl10IsTexture(gl10Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl10ColorPointer(gl10Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl10EnableClientState(gl10Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl10DisableClientState(gl10Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl10Indexub(gl10Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl10Indexubv(gl10Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl10InterleavedArrays(gl10Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl10NormalPointer(gl10Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl10PushClientAttrib(gl10Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl10PrioritizeTextures(gl10Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl10PopClientAttrib(gl10Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl10TexCoordPointer(gl10Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl10TexSubImage1D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl10TexSubImage2D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl10VertexPointer(gl10Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl10BlendColor(gl10Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl10BlendEquation(gl10Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl10CopyTexSubImage3D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl10DrawRangeElements(gl10Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl10TexImage3D(gl10Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl10TexSubImage3D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl10ActiveTexture(gl10Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl10ClientActiveTexture(gl10Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl10CompressedTexImage1D(gl10Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl10CompressedTexImage2D(gl10Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl10CompressedTexImage3D(gl10Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl10CompressedTexSubImage1D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl10CompressedTexSubImage2D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl10CompressedTexSubImage3D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl10GetCompressedTexImage(gl10Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl10LoadTransposeMatrixd(gl10Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl10LoadTransposeMatrixf(gl10Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl10MultTransposeMatrixd(gl10Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl10MultTransposeMatrixf(gl10Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl10SampleCoverage(gl10Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl10BlendFuncSeparate(gl10Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl10FogCoordPointer(gl10Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl10FogCoordd(gl10Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl10FogCoordf(gl10Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl10FogCoorddv(gl10Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl10FogCoordfv(gl10Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl10MultiDrawArrays(gl10Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl10MultiDrawElements(gl10Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl10PointParameterf(gl10Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl10PointParameteri(gl10Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl10SecondaryColor3b(gl10Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl10SecondaryColor3s(gl10Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl10SecondaryColor3i(gl10Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl10SecondaryColor3f(gl10Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl10SecondaryColor3d(gl10Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl10SecondaryColor3ub(gl10Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl10SecondaryColor3us(gl10Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl10SecondaryColor3ui(gl10Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl10SecondaryColor3bv(gl10Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl10SecondaryColor3sv(gl10Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl10SecondaryColor3iv(gl10Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl10SecondaryColor3fv(gl10Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl10SecondaryColor3dv(gl10Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl10SecondaryColor3ubv(gl10Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl10SecondaryColor3usv(gl10Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl10SecondaryColor3uiv(gl10Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl10SecondaryColorPointer(gl10Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl10WindowPos2s(gl10Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl10WindowPos2i(gl10Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl10WindowPos2f(gl10Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl10WindowPos2d(gl10Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl10WindowPos3s(gl10Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl10WindowPos3i(gl10Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl10WindowPos3f(gl10Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl10WindowPos3d(gl10Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl10WindowPos2sv(gl10Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl10WindowPos2iv(gl10Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl10WindowPos2fv(gl10Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl10WindowPos2dv(gl10Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl10WindowPos3sv(gl10Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl10WindowPos3iv(gl10Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl10WindowPos3fv(gl10Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl10WindowPos3dv(gl10Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl10BeginQuery(gl10Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl10BindBuffer(gl10Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl10BufferData(gl10Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl10BufferSubData(gl10Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl10DeleteBuffers(gl10Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl10DeleteQueries(gl10Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl10GenBuffers(gl10Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl10GenQueries(gl10Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl10GetBufferParameteriv(gl10Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl10GetBufferPointerv(gl10Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl10GetBufferSubData(gl10Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl10GetQueryObjectiv(gl10Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl10GetQueryObjectuiv(gl10Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl10GetQueryiv(gl10Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl10IsBuffer(gl10Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl10IsQuery(gl10Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl10MapBuffer(gl10Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl10UnmapBuffer(gl10Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl10AttachShader(gl10Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl10BindAttribLocation(gl10Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl10BlendEquationSeperate(gl10Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl10CompileShader(gl10Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl10CreateProgram(gl10Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl10CreateShader(gl10Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl10DeleteProgram(gl10Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl10DeleteShader(gl10Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl10DetachShader(gl10Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl10EnableVertexAttribArray(gl10Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl10DisableVertexAttribArray(gl10Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl10DrawBuffers(gl10Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl10GetActiveAttrib(gl10Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl10GetActiveUniform(gl10Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl10GetAttachedShaders(gl10Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl10GetAttribLocation(gl10Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl10GetProgramiv(gl10Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl10GetProgramInfoLog(gl10Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl10GetShaderiv(gl10Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl10GetShaderInfoLog(gl10Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl10GetShaderSource(gl10Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl10GetUniformfv(gl10Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl10GetUniformiv(gl10Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl10GetUniformLocation(gl10Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl10GetVertexAttribdv(gl10Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl10GetVertexAttribfv(gl10Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl10GetVertexAttribiv(gl10Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl10GetVertexAttribPointerv(gl10Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl10IsProgram(gl10Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl10IsShader(gl10Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl10LinkProgram(gl10Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl10ShaderSource(gl10Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl10StencilFuncSeparate(gl10Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl10StencilMaskSeparate(gl10Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl10StencilOpSeparate(gl10Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl10Uniform1f(gl10Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl10Uniform2f(gl10Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl10Uniform3f(gl10Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl10Uniform4f(gl10Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl10Uniform1i(gl10Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl10Uniform2i(gl10Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl10Uniform3i(gl10Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl10Uniform4i(gl10Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl10Uniform1fv(gl10Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl10Uniform2fv(gl10Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl10Uniform3fv(gl10Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl10Uniform4fv(gl10Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl10Uniform1iv(gl10Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl10Uniform2iv(gl10Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl10Uniform3iv(gl10Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl10Uniform4iv(gl10Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl10UseProgram(gl10Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl10ValidateProgram(gl10Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl10VertexAttribPointer(gl10Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl10VertexAttrib1f(gl10Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl10VertexAttrib1s(gl10Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl10VertexAttrib1d(gl10Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl10VertexAttrib2f(gl10Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl10VertexAttrib2s(gl10Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl10VertexAttrib2d(gl10Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl10VertexAttrib3f(gl10Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl10VertexAttrib3s(gl10Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl10VertexAttrib3d(gl10Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl10VertexAttrib4f(gl10Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl10VertexAttrib4s(gl10Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl10VertexAttrib4d(gl10Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl10VertexAttrib4Nuv(gl10Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl10VertexAttrib1fv(gl10Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl10VertexAttrib1sv(gl10Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl10VertexAttrib1dv(gl10Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl10VertexAttrib2fv(gl10Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl10VertexAttrib2sv(gl10Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl10VertexAttrib2dv(gl10Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl10VertexAttrib3fv(gl10Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl10VertexAttrib3sv(gl10Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl10VertexAttrib3dv(gl10Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl10VertexAttrib4fv(gl10Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl10VertexAttrib4sv(gl10Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl10VertexAttrib4dv(gl10Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl10VertexAttrib4iv(gl10Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl10VertexAttrib4bv(gl10Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl10VertexAttrib4ubv(gl10Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl10VertexAttrib4usv(gl10Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl10VertexAttrib4uiv(gl10Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl10VertexAttrib4Nbv(gl10Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl10VertexAttrib4Nsv(gl10Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl10VertexAttrib4Niv(gl10Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl10VertexAttrib4Nubv(gl10Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl10VertexAttrib4Nusv(gl10Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl10VertexAttrib4Nuiv(gl10Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl10UniformMatrix2fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl10UniformMatrix3fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl10UniformMatrix4fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl10UniformMatrix2x3fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl10UniformMatrix3x2fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl10UniformMatrix2x4fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl10UniformMatrix4x2fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl10UniformMatrix3x4fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl10UniformMatrix4x3fv(gl10Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl10Context* gl10NewContext() {
    gl10Context* glc = calloc(1, sizeof(gl10Context));

    // Preload all procedures
    glc->fnAccum = (gl10PAccum)gl10LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl10PAlphaFunc)gl10LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl10PBegin)gl10LibGetProcAddress("glBegin");
    glc->fnEnd = (gl10PEnd)gl10LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl10PBitmap)gl10LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl10PBlendFunc)gl10LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl10PCallList)gl10LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl10PCallLists)gl10LibGetProcAddress("glCallLists");
    glc->fnClear = (gl10PClear)gl10LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl10PClearAccum)gl10LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl10PClearColor)gl10LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl10PClearDepth)gl10LibGetProcAddress("glClearDepth");
    glc->fnClearDepthf = (gl10PClearDepthf)gl10LibGetProcAddress("glClearDepthf");
    glc->fnClearIndex = (gl10PClearIndex)gl10LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl10PClearStencil)gl10LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl10PClipPlane)gl10LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl10PColor3b)gl10LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl10PColor3d)gl10LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl10PColor3f)gl10LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl10PColor3i)gl10LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl10PColor3s)gl10LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl10PColor3ub)gl10LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl10PColor3ui)gl10LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl10PColor3us)gl10LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl10PColor4b)gl10LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl10PColor4d)gl10LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl10PColor4f)gl10LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl10PColor4i)gl10LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl10PColor4s)gl10LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl10PColor4ub)gl10LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl10PColor4ui)gl10LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl10PColor4us)gl10LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl10PColor3bv)gl10LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl10PColor3dv)gl10LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl10PColor3fv)gl10LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl10PColor3iv)gl10LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl10PColor3sv)gl10LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl10PColor3ubv)gl10LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl10PColor3uiv)gl10LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl10PColor3usv)gl10LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl10PColor4bv)gl10LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl10PColor4dv)gl10LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl10PColor4fv)gl10LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl10PColor4iv)gl10LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl10PColor4sv)gl10LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl10PColor4ubv)gl10LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl10PColor4uiv)gl10LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl10PColor4usv)gl10LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl10PColorMask)gl10LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl10PColorMaterial)gl10LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl10PCopyPixels)gl10LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl10PCullFace)gl10LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl10PDeleteLists)gl10LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl10PDepthFunc)gl10LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl10PDepthMask)gl10LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl10PDepthRange)gl10LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl10PEnable)gl10LibGetProcAddress("glEnable");
    glc->fnDisable = (gl10PDisable)gl10LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl10PDrawBuffer)gl10LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl10PDrawPixels)gl10LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl10PEdgeFlag)gl10LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl10PEdgeFlagv)gl10LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl10PEdgeFlagPointer)gl10LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl10PEvalCoord1d)gl10LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl10PEvalCoord1f)gl10LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl10PEvalCoord2d)gl10LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl10PEvalCoord2f)gl10LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl10PEvalCoord1dv)gl10LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl10PEvalCoord1fv)gl10LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl10PEvalCoord2dv)gl10LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl10PEvalCoord2fv)gl10LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl10PEvalMesh1)gl10LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl10PEvalMesh2)gl10LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl10PEvalPoint1)gl10LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl10PEvalPoint2)gl10LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl10PFeedbackBuffer)gl10LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl10PFinish)gl10LibGetProcAddress("glFinish");
    glc->fnFlush = (gl10PFlush)gl10LibGetProcAddress("glFlush");
    glc->fnFogf = (gl10PFogf)gl10LibGetProcAddress("glFogf");
    glc->fnFogi = (gl10PFogi)gl10LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl10PFogfv)gl10LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl10PFogiv)gl10LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl10PFrontFace)gl10LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl10PFrustum)gl10LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl10PGenLists)gl10LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl10PGetBooleanv)gl10LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl10PGetDoublev)gl10LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl10PGetFloatv)gl10LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl10PGetIntegerv)gl10LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl10PGetClipPlane)gl10LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl10PGetError)gl10LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl10PGetLightfv)gl10LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl10PGetLightiv)gl10LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl10PGetMapdv)gl10LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl10PGetMapfv)gl10LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl10PGetMapiv)gl10LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl10PGetMaterialfv)gl10LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl10PGetMaterialiv)gl10LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl10PGetPixelMapfv)gl10LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl10PGetPixelMapuiv)gl10LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl10PGetPixelMapusv)gl10LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl10PGetPolygonStipple)gl10LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl10PGetString)gl10LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl10PGetTexEnvfv)gl10LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl10PGetTexEnviv)gl10LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl10PGetTexGendv)gl10LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl10PGetTexGenfv)gl10LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl10PGetTexGeniv)gl10LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl10PGetTexImage)gl10LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl10PGetTexLevelParameterfv)gl10LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl10PGetTexLevelParameteriv)gl10LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl10PGetTexParameterfv)gl10LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl10PGetTexParameteriv)gl10LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl10PHint)gl10LibGetProcAddress("glHint");
    glc->fnIndexd = (gl10PIndexd)gl10LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl10PIndexf)gl10LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl10PIndexi)gl10LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl10PIndexs)gl10LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl10PIndexdv)gl10LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl10PIndexfv)gl10LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl10PIndexiv)gl10LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl10PIndexsv)gl10LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl10PIndexMask)gl10LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl10PIndexPointer)gl10LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl10PInitNames)gl10LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl10PIsEnabled)gl10LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl10PIsList)gl10LibGetProcAddress("glIsList");
    glc->fnLightf = (gl10PLightf)gl10LibGetProcAddress("glLightf");
    glc->fnLighti = (gl10PLighti)gl10LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl10PLightfv)gl10LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl10PLightiv)gl10LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl10PLightModelf)gl10LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl10PLightModeli)gl10LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl10PLightModelfv)gl10LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl10PLightModeliv)gl10LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl10PLineStipple)gl10LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl10PLineWidth)gl10LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl10PListBase)gl10LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl10PLoadIdentity)gl10LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl10PLoadMatrixd)gl10LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl10PLoadMatrixf)gl10LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl10PLoadName)gl10LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl10PLogicOp)gl10LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl10PMap1d)gl10LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl10PMap1f)gl10LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl10PMap2d)gl10LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl10PMap2f)gl10LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl10PMapGrid1d)gl10LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl10PMapGrid1f)gl10LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl10PMapGrid2d)gl10LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl10PMapGrid2f)gl10LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl10PMaterialf)gl10LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl10PMateriali)gl10LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl10PMaterialfv)gl10LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl10PMaterialiv)gl10LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl10PMatrixMode)gl10LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl10PMultMatrixd)gl10LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl10PMultMatrixf)gl10LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl10PNewList)gl10LibGetProcAddress("glNewList");
    glc->fnEndList = (gl10PEndList)gl10LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl10PNormal3b)gl10LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl10PNormal3d)gl10LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl10PNormal3f)gl10LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl10PNormal3i)gl10LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl10PNormal3s)gl10LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl10PNormal3bv)gl10LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl10PNormal3dv)gl10LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl10PNormal3fv)gl10LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl10PNormal3iv)gl10LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl10PNormal3sv)gl10LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl10POrtho)gl10LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl10PPassThrough)gl10LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl10PPixelMapfv)gl10LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl10PPixelMapuiv)gl10LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl10PPixelMapusv)gl10LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl10PPixelStoref)gl10LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl10PPixelStorei)gl10LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl10PPixelTransferf)gl10LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl10PPixelTransferi)gl10LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl10PPixelZoom)gl10LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl10PPointSize)gl10LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl10PPolygonMode)gl10LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl10PPolygonStipple)gl10LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl10PPushAttrib)gl10LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl10PPopAttrib)gl10LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl10PPushMatrix)gl10LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl10PPopMatrix)gl10LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl10PPushName)gl10LibGetProcAddress("glPushName");
    glc->fnPopName = (gl10PPopName)gl10LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl10PRasterPos2d)gl10LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl10PRasterPos2f)gl10LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl10PRasterPos2i)gl10LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl10PRasterPos2s)gl10LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl10PRasterPos3d)gl10LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl10PRasterPos3f)gl10LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl10PRasterPos3i)gl10LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl10PRasterPos3s)gl10LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl10PRasterPos4d)gl10LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl10PRasterPos4f)gl10LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl10PRasterPos4i)gl10LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl10PRasterPos4s)gl10LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl10PRasterPos2dv)gl10LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl10PRasterPos2fv)gl10LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl10PRasterPos2iv)gl10LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl10PRasterPos2sv)gl10LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl10PRasterPos3dv)gl10LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl10PRasterPos3fv)gl10LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl10PRasterPos3iv)gl10LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl10PRasterPos3sv)gl10LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl10PRasterPos4dv)gl10LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl10PRasterPos4fv)gl10LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl10PRasterPos4iv)gl10LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl10PRasterPos4sv)gl10LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl10PReadBuffer)gl10LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl10PReadPixels)gl10LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl10PRectd)gl10LibGetProcAddress("glRectd");
    glc->fnRectf = (gl10PRectf)gl10LibGetProcAddress("glRectf");
    glc->fnRecti = (gl10PRecti)gl10LibGetProcAddress("glRecti");
    glc->fnRects = (gl10PRects)gl10LibGetProcAddress("glRects");
    glc->fnRectdv = (gl10PRectdv)gl10LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl10PRectfv)gl10LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl10PRectiv)gl10LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl10PRectsv)gl10LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl10PRenderMode)gl10LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl10PRotated)gl10LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl10PRotatef)gl10LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl10PScaled)gl10LibGetProcAddress("glScaled");
    glc->fnScalef = (gl10PScalef)gl10LibGetProcAddress("glScalef");
    glc->fnScissor = (gl10PScissor)gl10LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl10PSelectBuffer)gl10LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl10PShadeModel)gl10LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl10PStencilFunc)gl10LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl10PStencilMask)gl10LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl10PStencilOp)gl10LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl10PTexCoord1d)gl10LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl10PTexCoord1f)gl10LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl10PTexCoord1i)gl10LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl10PTexCoord1s)gl10LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl10PTexCoord2d)gl10LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl10PTexCoord2f)gl10LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl10PTexCoord2i)gl10LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl10PTexCoord2s)gl10LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl10PTexCoord3d)gl10LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl10PTexCoord3f)gl10LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl10PTexCoord3i)gl10LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl10PTexCoord3s)gl10LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl10PTexCoord4d)gl10LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl10PTexCoord4f)gl10LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl10PTexCoord4i)gl10LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl10PTexCoord4s)gl10LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl10PTexCoord1dv)gl10LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl10PTexCoord1fv)gl10LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl10PTexCoord1iv)gl10LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl10PTexCoord1sv)gl10LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl10PTexCoord2dv)gl10LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl10PTexCoord2fv)gl10LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl10PTexCoord2iv)gl10LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl10PTexCoord2sv)gl10LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl10PTexCoord3dv)gl10LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl10PTexCoord3fv)gl10LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl10PTexCoord3iv)gl10LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl10PTexCoord3sv)gl10LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl10PTexCoord4dv)gl10LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl10PTexCoord4fv)gl10LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl10PTexCoord4iv)gl10LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl10PTexCoord4sv)gl10LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl10PTexEnvf)gl10LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl10PTexEnvi)gl10LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl10PTexEnvfv)gl10LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl10PTexEnviv)gl10LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl10PTexGend)gl10LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl10PTexGenf)gl10LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl10PTexGeni)gl10LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl10PTexGendv)gl10LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl10PTexGenfv)gl10LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl10PTexGeniv)gl10LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl10PTexImage1D)gl10LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl10PTexImage2D)gl10LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl10PTexParameterf)gl10LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl10PTexParameteri)gl10LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl10PTexParameterfv)gl10LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl10PTexParameteriv)gl10LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl10PTranslated)gl10LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl10PTranslatef)gl10LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl10PVertex2s)gl10LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl10PVertex2i)gl10LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl10PVertex2f)gl10LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl10PVertex2d)gl10LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl10PVertex3s)gl10LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl10PVertex3i)gl10LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl10PVertex3f)gl10LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl10PVertex3d)gl10LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl10PVertex4s)gl10LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl10PVertex4i)gl10LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl10PVertex4f)gl10LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl10PVertex4d)gl10LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl10PViewport)gl10LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl10PGetConvolutionParameterfv)gl10LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl10PGetConvolutionParameteriv)gl10LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnColorTable = (gl10PColorTable)gl10GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl10PColorTableParameterfv)gl10GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl10PColorTableParameteriv)gl10GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl10PColorSubTable)gl10GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl10PConvolutionFilter1D)gl10GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl10PConvolutionFilter2D)gl10GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl10PConvolutionParameterf)gl10GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl10PConvolutionParameteri)gl10GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl10PCopyColorTable)gl10GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl10PCopyColorSubTable)gl10GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl10PCopyConvolutionFilter1D)gl10GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl10PCopyConvolutionFilter2D)gl10GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl10PGetColorTable)gl10GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl10PGetColorTableParameterfv)gl10GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl10PGetColorTableParameteriv)gl10GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl10PGetConvolutionFilter)gl10GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl10PGetHistogram)gl10GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl10PGetHistogramParameterfv)gl10GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl10PGetHistogramParameteriv)gl10GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl10PGetSeparableFilter)gl10GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl10PHistogram)gl10GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl10PMinmax)gl10GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl10PMultiTexCoord1s)gl10GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl10PMultiTexCoord1i)gl10GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl10PMultiTexCoord1f)gl10GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl10PMultiTexCoord1d)gl10GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl10PMultiTexCoord2s)gl10GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl10PMultiTexCoord2i)gl10GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl10PMultiTexCoord2f)gl10GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl10PMultiTexCoord2d)gl10GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl10PMultiTexCoord3s)gl10GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl10PMultiTexCoord3i)gl10GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl10PMultiTexCoord3f)gl10GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl10PMultiTexCoord3d)gl10GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl10PMultiTexCoord4s)gl10GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl10PMultiTexCoord4i)gl10GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl10PMultiTexCoord4f)gl10GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl10PMultiTexCoord4d)gl10GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl10PMultiTexCoord1sv)gl10GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl10PMultiTexCoord1iv)gl10GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl10PMultiTexCoord1fv)gl10GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl10PMultiTexCoord1dv)gl10GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl10PMultiTexCoord2sv)gl10GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl10PMultiTexCoord2iv)gl10GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl10PMultiTexCoord2fv)gl10GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl10PMultiTexCoord2dv)gl10GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl10PMultiTexCoord3sv)gl10GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl10PMultiTexCoord3iv)gl10GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl10PMultiTexCoord3fv)gl10GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl10PMultiTexCoord3dv)gl10GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl10PMultiTexCoord4sv)gl10GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl10PMultiTexCoord4iv)gl10GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl10PMultiTexCoord4fv)gl10GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl10PMultiTexCoord4dv)gl10GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl10PResetHistogram)gl10GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl10PResetMinmax)gl10GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl10PSeparableFilter2D)gl10GLGetProcAddress("glSeparableFilter2D");
    glc->fnAreTexturesResident = (gl10PAreTexturesResident)gl10LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl10PArrayElement)gl10LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl10PDrawArrays)gl10LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl10PDrawElements)gl10LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl10PGetPointerv)gl10LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl10PPolygonOffset)gl10LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl10PCopyTexImage1D)gl10LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl10PCopyTexImage2D)gl10LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl10PCopyTexSubImage1D)gl10LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl10PCopyTexSubImage2D)gl10LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl10PBindTexture)gl10LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl10PDeleteTextures)gl10LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl10PGenTextures)gl10LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl10PIsTexture)gl10LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl10PColorPointer)gl10LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl10PEnableClientState)gl10LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl10PDisableClientState)gl10LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl10PIndexub)gl10LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl10PIndexubv)gl10LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl10PInterleavedArrays)gl10LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl10PNormalPointer)gl10LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl10PPushClientAttrib)gl10LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl10PPrioritizeTextures)gl10LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl10PPopClientAttrib)gl10LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl10PTexCoordPointer)gl10LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl10PTexSubImage1D)gl10LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl10PTexSubImage2D)gl10LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl10PVertexPointer)gl10LibGetProcAddress("glVertexPointer");
    glc->fnBlendColor = (gl10PBlendColor)gl10GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl10PBlendEquation)gl10GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl10PCopyTexSubImage3D)gl10GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl10PDrawRangeElements)gl10GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl10PTexImage3D)gl10GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl10PTexSubImage3D)gl10GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl10PActiveTexture)gl10GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl10PClientActiveTexture)gl10GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl10PCompressedTexImage1D)gl10GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl10PCompressedTexImage2D)gl10GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl10PCompressedTexImage3D)gl10GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl10PCompressedTexSubImage1D)gl10GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl10PCompressedTexSubImage2D)gl10GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl10PCompressedTexSubImage3D)gl10GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl10PGetCompressedTexImage)gl10GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl10PLoadTransposeMatrixd)gl10GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl10PLoadTransposeMatrixf)gl10GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl10PMultTransposeMatrixd)gl10GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl10PMultTransposeMatrixf)gl10GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl10PSampleCoverage)gl10GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl10PBlendFuncSeparate)gl10GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl10PFogCoordPointer)gl10GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl10PFogCoordd)gl10GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl10PFogCoordf)gl10GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl10PFogCoorddv)gl10GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl10PFogCoordfv)gl10GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl10PMultiDrawArrays)gl10GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl10PMultiDrawElements)gl10GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl10PPointParameterf)gl10GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl10PPointParameteri)gl10GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl10PSecondaryColor3b)gl10GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl10PSecondaryColor3s)gl10GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl10PSecondaryColor3i)gl10GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl10PSecondaryColor3f)gl10GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl10PSecondaryColor3d)gl10GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl10PSecondaryColor3ub)gl10GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl10PSecondaryColor3us)gl10GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl10PSecondaryColor3ui)gl10GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl10PSecondaryColor3bv)gl10GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl10PSecondaryColor3sv)gl10GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl10PSecondaryColor3iv)gl10GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl10PSecondaryColor3fv)gl10GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl10PSecondaryColor3dv)gl10GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl10PSecondaryColor3ubv)gl10GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl10PSecondaryColor3usv)gl10GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl10PSecondaryColor3uiv)gl10GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl10PSecondaryColorPointer)gl10GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl10PWindowPos2s)gl10GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl10PWindowPos2i)gl10GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl10PWindowPos2f)gl10GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl10PWindowPos2d)gl10GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl10PWindowPos3s)gl10GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl10PWindowPos3i)gl10GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl10PWindowPos3f)gl10GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl10PWindowPos3d)gl10GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl10PWindowPos2sv)gl10GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl10PWindowPos2iv)gl10GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl10PWindowPos2fv)gl10GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl10PWindowPos2dv)gl10GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl10PWindowPos3sv)gl10GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl10PWindowPos3iv)gl10GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl10PWindowPos3fv)gl10GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl10PWindowPos3dv)gl10GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl10PBeginQuery)gl10GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl10PBindBuffer)gl10GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl10PBufferData)gl10GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl10PBufferSubData)gl10GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl10PDeleteBuffers)gl10GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl10PDeleteQueries)gl10GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl10PGenBuffers)gl10GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl10PGenQueries)gl10GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl10PGetBufferParameteriv)gl10GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl10PGetBufferPointerv)gl10GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl10PGetBufferSubData)gl10GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl10PGetQueryObjectiv)gl10GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl10PGetQueryObjectuiv)gl10GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl10PGetQueryiv)gl10GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl10PIsBuffer)gl10GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl10PIsQuery)gl10GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl10PMapBuffer)gl10GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl10PUnmapBuffer)gl10GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl10PAttachShader)gl10GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl10PBindAttribLocation)gl10GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl10PBlendEquationSeperate)gl10GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl10PCompileShader)gl10GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl10PCreateProgram)gl10GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl10PCreateShader)gl10GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl10PDeleteProgram)gl10GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl10PDeleteShader)gl10GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl10PDetachShader)gl10GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl10PEnableVertexAttribArray)gl10GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl10PDisableVertexAttribArray)gl10GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl10PDrawBuffers)gl10LibGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl10PGetActiveAttrib)gl10LibGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl10PGetActiveUniform)gl10LibGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl10PGetAttachedShaders)gl10LibGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl10PGetAttribLocation)gl10LibGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl10PGetProgramiv)gl10LibGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl10PGetProgramInfoLog)gl10LibGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl10PGetShaderiv)gl10LibGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl10PGetShaderInfoLog)gl10LibGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl10PGetShaderSource)gl10LibGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl10PGetUniformfv)gl10LibGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl10PGetUniformiv)gl10LibGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl10PGetUniformLocation)gl10LibGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl10PGetVertexAttribdv)gl10LibGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl10PGetVertexAttribfv)gl10LibGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl10PGetVertexAttribiv)gl10LibGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl10PGetVertexAttribPointerv)gl10LibGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl10PIsProgram)gl10LibGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl10PIsShader)gl10LibGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl10PLinkProgram)gl10LibGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl10PShaderSource)gl10LibGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl10PStencilFuncSeparate)gl10LibGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl10PStencilMaskSeparate)gl10LibGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl10PStencilOpSeparate)gl10LibGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl10PUniform1f)gl10LibGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl10PUniform2f)gl10LibGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl10PUniform3f)gl10LibGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl10PUniform4f)gl10LibGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl10PUniform1i)gl10LibGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl10PUniform2i)gl10LibGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl10PUniform3i)gl10LibGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl10PUniform4i)gl10LibGetProcAddress("glUniform4i");
    glc->fnUniform1fv = (gl10PUniform1fv)gl10LibGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl10PUniform2fv)gl10LibGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl10PUniform3fv)gl10LibGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl10PUniform4fv)gl10LibGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl10PUniform1iv)gl10LibGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl10PUniform2iv)gl10LibGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl10PUniform3iv)gl10LibGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl10PUniform4iv)gl10LibGetProcAddress("glUniform4iv");
    glc->fnUseProgram = (gl10PUseProgram)gl10LibGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl10PValidateProgram)gl10LibGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl10PVertexAttribPointer)gl10LibGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl10PVertexAttrib1f)gl10LibGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl10PVertexAttrib1s)gl10LibGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl10PVertexAttrib1d)gl10LibGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl10PVertexAttrib2f)gl10LibGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl10PVertexAttrib2s)gl10LibGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl10PVertexAttrib2d)gl10LibGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl10PVertexAttrib3f)gl10LibGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl10PVertexAttrib3s)gl10LibGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl10PVertexAttrib3d)gl10LibGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl10PVertexAttrib4f)gl10LibGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl10PVertexAttrib4s)gl10LibGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl10PVertexAttrib4d)gl10LibGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl10PVertexAttrib4Nuv)gl10LibGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl10PVertexAttrib1fv)gl10LibGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl10PVertexAttrib1sv)gl10LibGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl10PVertexAttrib1dv)gl10LibGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl10PVertexAttrib2fv)gl10LibGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl10PVertexAttrib2sv)gl10LibGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl10PVertexAttrib2dv)gl10LibGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl10PVertexAttrib3fv)gl10LibGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl10PVertexAttrib3sv)gl10LibGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl10PVertexAttrib3dv)gl10LibGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl10PVertexAttrib4fv)gl10LibGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl10PVertexAttrib4sv)gl10LibGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl10PVertexAttrib4dv)gl10LibGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl10PVertexAttrib4iv)gl10LibGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl10PVertexAttrib4bv)gl10LibGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl10PVertexAttrib4ubv)gl10LibGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl10PVertexAttrib4usv)gl10LibGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl10PVertexAttrib4uiv)gl10LibGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl10PVertexAttrib4Nbv)gl10LibGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl10PVertexAttrib4Nsv)gl10LibGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl10PVertexAttrib4Niv)gl10LibGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl10PVertexAttrib4Nubv)gl10LibGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl10PVertexAttrib4Nusv)gl10LibGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl10PVertexAttrib4Nuiv)gl10LibGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl10PUniformMatrix2fv)gl10LibGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl10PUniformMatrix3fv)gl10LibGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl10PUniformMatrix4fv)gl10LibGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl10PUniformMatrix2x3fv)gl10LibGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl10PUniformMatrix3x2fv)gl10LibGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl10PUniformMatrix2x4fv)gl10LibGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl10PUniformMatrix4x2fv)gl10LibGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl10PUniformMatrix3x4fv)gl10LibGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl10PUniformMatrix4x3fv)gl10LibGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

