
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl14.h"

#ifdef _WIN32
HMODULE gl14OpenGL32;

void* gl14LibGetProcAddress(char* name) {
	if(gl14OpenGL32 == NULL) {
		gl14OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl14OpenGL32, TEXT(name));
}

void* gl14GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl14Accum(gl14Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl14AlphaFunc(gl14Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl14Begin(gl14Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl14End(gl14Context* glc) {
    return glc->fnEnd();
}

void gl14Bitmap(gl14Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl14BlendFunc(gl14Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl14CallList(gl14Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl14CallLists(gl14Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl14Clear(gl14Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl14ClearAccum(gl14Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl14ClearColor(gl14Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl14ClearDepth(gl14Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl14ClearDepthf(gl14Context* glc, GLclampf depth) {
    return glc->fnClearDepthf(depth);
}

void gl14ClearIndex(gl14Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl14ClearStencil(gl14Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl14ClipPlane(gl14Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl14Color3b(gl14Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl14Color3d(gl14Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl14Color3f(gl14Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl14Color3i(gl14Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl14Color3s(gl14Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl14Color3ub(gl14Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl14Color3ui(gl14Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl14Color3us(gl14Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl14Color4b(gl14Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl14Color4d(gl14Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl14Color4f(gl14Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl14Color4i(gl14Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl14Color4s(gl14Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl14Color4ub(gl14Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl14Color4ui(gl14Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl14Color4us(gl14Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl14Color3bv(gl14Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl14Color3dv(gl14Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl14Color3fv(gl14Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl14Color3iv(gl14Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl14Color3sv(gl14Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl14Color3ubv(gl14Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl14Color3uiv(gl14Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl14Color3usv(gl14Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl14Color4bv(gl14Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl14Color4dv(gl14Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl14Color4fv(gl14Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl14Color4iv(gl14Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl14Color4sv(gl14Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl14Color4ubv(gl14Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl14Color4uiv(gl14Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl14Color4usv(gl14Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl14ColorMask(gl14Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl14ColorMaterial(gl14Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl14CopyPixels(gl14Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl14CullFace(gl14Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl14DeleteLists(gl14Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl14DepthFunc(gl14Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl14DepthMask(gl14Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl14DepthRange(gl14Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl14Enable(gl14Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl14Disable(gl14Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl14DrawBuffer(gl14Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl14DrawPixels(gl14Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl14EdgeFlag(gl14Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl14EdgeFlagv(gl14Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl14EdgeFlagPointer(gl14Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl14EvalCoord1d(gl14Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl14EvalCoord1f(gl14Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl14EvalCoord2d(gl14Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl14EvalCoord2f(gl14Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl14EvalCoord1dv(gl14Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl14EvalCoord1fv(gl14Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl14EvalCoord2dv(gl14Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl14EvalCoord2fv(gl14Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl14EvalMesh1(gl14Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl14EvalMesh2(gl14Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl14EvalPoint1(gl14Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl14EvalPoint2(gl14Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl14FeedbackBuffer(gl14Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl14Finish(gl14Context* glc) {
    return glc->fnFinish();
}

void gl14Flush(gl14Context* glc) {
    return glc->fnFlush();
}

void gl14Fogf(gl14Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl14Fogi(gl14Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl14Fogfv(gl14Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl14Fogiv(gl14Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl14FrontFace(gl14Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl14Frustum(gl14Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl14GenLists(gl14Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl14GetBooleanv(gl14Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl14GetDoublev(gl14Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl14GetFloatv(gl14Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl14GetIntegerv(gl14Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl14GetClipPlane(gl14Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl14GetError(gl14Context* glc) {
    return glc->fnGetError();
}

void gl14GetLightfv(gl14Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl14GetLightiv(gl14Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl14GetMapdv(gl14Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl14GetMapfv(gl14Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl14GetMapiv(gl14Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl14GetMaterialfv(gl14Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl14GetMaterialiv(gl14Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl14GetPixelMapfv(gl14Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl14GetPixelMapuiv(gl14Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl14GetPixelMapusv(gl14Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl14GetPolygonStipple(gl14Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl14GetString(gl14Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl14GetTexEnvfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl14GetTexEnviv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl14GetTexGendv(gl14Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl14GetTexGenfv(gl14Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl14GetTexGeniv(gl14Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl14GetTexImage(gl14Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl14GetTexLevelParameterfv(gl14Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl14GetTexLevelParameteriv(gl14Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl14GetTexParameterfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl14GetTexParameteriv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl14Hint(gl14Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl14Indexd(gl14Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl14Indexf(gl14Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl14Indexi(gl14Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl14Indexs(gl14Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl14Indexdv(gl14Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl14Indexfv(gl14Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl14Indexiv(gl14Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl14Indexsv(gl14Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl14IndexMask(gl14Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl14IndexPointer(gl14Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl14InitNames(gl14Context* glc) {
    return glc->fnInitNames();
}

void gl14IsEnabled(gl14Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl14IsList(gl14Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl14Lightf(gl14Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl14Lighti(gl14Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl14Lightfv(gl14Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl14Lightiv(gl14Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl14LightModelf(gl14Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl14LightModeli(gl14Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl14LightModelfv(gl14Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl14LightModeliv(gl14Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl14LineStipple(gl14Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl14LineWidth(gl14Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl14ListBase(gl14Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl14LoadIdentity(gl14Context* glc) {
    return glc->fnLoadIdentity();
}

void gl14LoadMatrixd(gl14Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl14LoadMatrixf(gl14Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl14LoadName(gl14Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl14LogicOp(gl14Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl14Map1d(gl14Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl14Map1f(gl14Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl14Map2d(gl14Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl14Map2f(gl14Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl14MapGrid1d(gl14Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl14MapGrid1f(gl14Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl14MapGrid2d(gl14Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl14MapGrid2f(gl14Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl14Materialf(gl14Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl14Materiali(gl14Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl14Materialfv(gl14Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl14Materialiv(gl14Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl14MatrixMode(gl14Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl14MultMatrixd(gl14Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl14MultMatrixf(gl14Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl14NewList(gl14Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl14EndList(gl14Context* glc) {
    return glc->fnEndList();
}

void gl14Normal3b(gl14Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl14Normal3d(gl14Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl14Normal3f(gl14Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl14Normal3i(gl14Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl14Normal3s(gl14Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl14Normal3bv(gl14Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl14Normal3dv(gl14Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl14Normal3fv(gl14Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl14Normal3iv(gl14Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl14Normal3sv(gl14Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl14Ortho(gl14Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl14PassThrough(gl14Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl14PixelMapfv(gl14Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl14PixelMapuiv(gl14Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl14PixelMapusv(gl14Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl14PixelStoref(gl14Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl14PixelStorei(gl14Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl14PixelTransferf(gl14Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl14PixelTransferi(gl14Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl14PixelZoom(gl14Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl14PointSize(gl14Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl14PolygonMode(gl14Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl14PolygonStipple(gl14Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl14PushAttrib(gl14Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl14PopAttrib(gl14Context* glc) {
    return glc->fnPopAttrib();
}

void gl14PushMatrix(gl14Context* glc) {
    return glc->fnPushMatrix();
}

void gl14PopMatrix(gl14Context* glc) {
    return glc->fnPopMatrix();
}

void gl14PushName(gl14Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl14PopName(gl14Context* glc) {
    return glc->fnPopName();
}

void gl14RasterPos2d(gl14Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl14RasterPos2f(gl14Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl14RasterPos2i(gl14Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl14RasterPos2s(gl14Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl14RasterPos3d(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl14RasterPos3f(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl14RasterPos3i(gl14Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl14RasterPos3s(gl14Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl14RasterPos4d(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl14RasterPos4f(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl14RasterPos4i(gl14Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl14RasterPos4s(gl14Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl14RasterPos2dv(gl14Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl14RasterPos2fv(gl14Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl14RasterPos2iv(gl14Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl14RasterPos2sv(gl14Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl14RasterPos3dv(gl14Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl14RasterPos3fv(gl14Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl14RasterPos3iv(gl14Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl14RasterPos3sv(gl14Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl14RasterPos4dv(gl14Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl14RasterPos4fv(gl14Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl14RasterPos4iv(gl14Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl14RasterPos4sv(gl14Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl14ReadBuffer(gl14Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl14ReadPixels(gl14Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl14Rectd(gl14Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl14Rectf(gl14Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl14Recti(gl14Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl14Rects(gl14Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl14Rectdv(gl14Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl14Rectfv(gl14Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl14Rectiv(gl14Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl14Rectsv(gl14Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl14RenderMode(gl14Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl14Rotated(gl14Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl14Rotatef(gl14Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl14Scaled(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl14Scalef(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl14Scissor(gl14Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl14SelectBuffer(gl14Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl14ShadeModel(gl14Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl14StencilFunc(gl14Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl14StencilMask(gl14Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl14StencilOp(gl14Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl14TexCoord1d(gl14Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl14TexCoord1f(gl14Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl14TexCoord1i(gl14Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl14TexCoord1s(gl14Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl14TexCoord2d(gl14Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl14TexCoord2f(gl14Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl14TexCoord2i(gl14Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl14TexCoord2s(gl14Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl14TexCoord3d(gl14Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl14TexCoord3f(gl14Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl14TexCoord3i(gl14Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl14TexCoord3s(gl14Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl14TexCoord4d(gl14Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl14TexCoord4f(gl14Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl14TexCoord4i(gl14Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl14TexCoord4s(gl14Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl14TexCoord1dv(gl14Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl14TexCoord1fv(gl14Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl14TexCoord1iv(gl14Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl14TexCoord1sv(gl14Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl14TexCoord2dv(gl14Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl14TexCoord2fv(gl14Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl14TexCoord2iv(gl14Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl14TexCoord2sv(gl14Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl14TexCoord3dv(gl14Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl14TexCoord3fv(gl14Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl14TexCoord3iv(gl14Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl14TexCoord3sv(gl14Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl14TexCoord4dv(gl14Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl14TexCoord4fv(gl14Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl14TexCoord4iv(gl14Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl14TexCoord4sv(gl14Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl14TexEnvf(gl14Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl14TexEnvi(gl14Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl14TexEnvfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl14TexEnviv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl14TexGend(gl14Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl14TexGenf(gl14Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl14TexGeni(gl14Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl14TexGendv(gl14Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl14TexGenfv(gl14Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl14TexGeniv(gl14Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl14TexImage1D(gl14Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl14TexImage2D(gl14Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl14TexParameterf(gl14Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl14TexParameteri(gl14Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl14TexParameterfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl14TexParameteriv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl14Translated(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl14Translatef(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl14Vertex2s(gl14Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl14Vertex2i(gl14Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl14Vertex2f(gl14Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl14Vertex2d(gl14Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl14Vertex3s(gl14Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl14Vertex3i(gl14Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl14Vertex3f(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl14Vertex3d(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl14Vertex4s(gl14Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl14Vertex4i(gl14Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl14Vertex4f(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl14Vertex4d(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl14Viewport(gl14Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl14GetConvolutionParameterfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl14GetConvolutionParameteriv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl14AreTexturesResident(gl14Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl14ArrayElement(gl14Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl14DrawArrays(gl14Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl14DrawElements(gl14Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl14GetPointerv(gl14Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl14PolygonOffset(gl14Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl14CopyTexImage1D(gl14Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl14CopyTexImage2D(gl14Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl14CopyTexSubImage1D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl14CopyTexSubImage2D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl14BindTexture(gl14Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl14DeleteTextures(gl14Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl14GenTextures(gl14Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl14IsTexture(gl14Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl14ColorPointer(gl14Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl14EnableClientState(gl14Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl14DisableClientState(gl14Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl14Indexub(gl14Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl14Indexubv(gl14Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl14InterleavedArrays(gl14Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl14NormalPointer(gl14Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl14PushClientAttrib(gl14Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl14PrioritizeTextures(gl14Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl14PopClientAttrib(gl14Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl14TexCoordPointer(gl14Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl14TexSubImage1D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl14TexSubImage2D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl14VertexPointer(gl14Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl14ColorTable(gl14Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl14ColorTableParameterfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl14ColorTableParameteriv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl14ColorSubTable(gl14Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl14ConvolutionFilter1D(gl14Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl14ConvolutionFilter2D(gl14Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl14ConvolutionParameterf(gl14Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl14ConvolutionParameteri(gl14Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl14CopyColorTable(gl14Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl14CopyColorSubTable(gl14Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl14CopyConvolutionFilter1D(gl14Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl14CopyConvolutionFilter2D(gl14Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl14GetColorTable(gl14Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl14GetColorTableParameterfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl14GetColorTableParameteriv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl14GetConvolutionFilter(gl14Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl14GetHistogram(gl14Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl14GetHistogramParameterfv(gl14Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl14GetHistogramParameteriv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl14GetSeparableFilter(gl14Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl14Histogram(gl14Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl14Minmax(gl14Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl14MultiTexCoord1s(gl14Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl14MultiTexCoord1i(gl14Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl14MultiTexCoord1f(gl14Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl14MultiTexCoord1d(gl14Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl14MultiTexCoord2s(gl14Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl14MultiTexCoord2i(gl14Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl14MultiTexCoord2f(gl14Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl14MultiTexCoord2d(gl14Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl14MultiTexCoord3s(gl14Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl14MultiTexCoord3i(gl14Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl14MultiTexCoord3f(gl14Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl14MultiTexCoord3d(gl14Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl14MultiTexCoord4s(gl14Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl14MultiTexCoord4i(gl14Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl14MultiTexCoord4f(gl14Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl14MultiTexCoord4d(gl14Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl14MultiTexCoord1sv(gl14Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl14MultiTexCoord1iv(gl14Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl14MultiTexCoord1fv(gl14Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl14MultiTexCoord1dv(gl14Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl14MultiTexCoord2sv(gl14Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl14MultiTexCoord2iv(gl14Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl14MultiTexCoord2fv(gl14Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl14MultiTexCoord2dv(gl14Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl14MultiTexCoord3sv(gl14Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl14MultiTexCoord3iv(gl14Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl14MultiTexCoord3fv(gl14Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl14MultiTexCoord3dv(gl14Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl14MultiTexCoord4sv(gl14Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl14MultiTexCoord4iv(gl14Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl14MultiTexCoord4fv(gl14Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl14MultiTexCoord4dv(gl14Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl14ResetHistogram(gl14Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl14ResetMinmax(gl14Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl14SeparableFilter2D(gl14Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl14BlendColor(gl14Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl14BlendEquation(gl14Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl14CopyTexSubImage3D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl14DrawRangeElements(gl14Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl14TexImage3D(gl14Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl14TexSubImage3D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl14ActiveTexture(gl14Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl14ClientActiveTexture(gl14Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl14CompressedTexImage1D(gl14Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl14CompressedTexImage2D(gl14Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl14CompressedTexImage3D(gl14Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl14CompressedTexSubImage1D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl14CompressedTexSubImage2D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl14CompressedTexSubImage3D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl14GetCompressedTexImage(gl14Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl14LoadTransposeMatrixd(gl14Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl14LoadTransposeMatrixf(gl14Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl14MultTransposeMatrixd(gl14Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl14MultTransposeMatrixf(gl14Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl14SampleCoverage(gl14Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl14BlendFuncSeparate(gl14Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl14FogCoordPointer(gl14Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl14FogCoordd(gl14Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl14FogCoordf(gl14Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl14FogCoorddv(gl14Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl14FogCoordfv(gl14Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl14MultiDrawArrays(gl14Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl14MultiDrawElements(gl14Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl14PointParameterf(gl14Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl14PointParameteri(gl14Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl14SecondaryColor3b(gl14Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl14SecondaryColor3s(gl14Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl14SecondaryColor3i(gl14Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl14SecondaryColor3f(gl14Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl14SecondaryColor3d(gl14Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl14SecondaryColor3ub(gl14Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl14SecondaryColor3us(gl14Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl14SecondaryColor3ui(gl14Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl14SecondaryColor3bv(gl14Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl14SecondaryColor3sv(gl14Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl14SecondaryColor3iv(gl14Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl14SecondaryColor3fv(gl14Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl14SecondaryColor3dv(gl14Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl14SecondaryColor3ubv(gl14Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl14SecondaryColor3usv(gl14Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl14SecondaryColor3uiv(gl14Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl14SecondaryColorPointer(gl14Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl14WindowPos2s(gl14Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl14WindowPos2i(gl14Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl14WindowPos2f(gl14Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl14WindowPos2d(gl14Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl14WindowPos3s(gl14Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl14WindowPos3i(gl14Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl14WindowPos3f(gl14Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl14WindowPos3d(gl14Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl14WindowPos2sv(gl14Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl14WindowPos2iv(gl14Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl14WindowPos2fv(gl14Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl14WindowPos2dv(gl14Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl14WindowPos3sv(gl14Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl14WindowPos3iv(gl14Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl14WindowPos3fv(gl14Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl14WindowPos3dv(gl14Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl14BeginQuery(gl14Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl14BindBuffer(gl14Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl14BufferData(gl14Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl14BufferSubData(gl14Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl14DeleteBuffers(gl14Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl14DeleteQueries(gl14Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl14GenBuffers(gl14Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl14GenQueries(gl14Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl14GetBufferParameteriv(gl14Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl14GetBufferPointerv(gl14Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl14GetBufferSubData(gl14Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl14GetQueryObjectiv(gl14Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl14GetQueryObjectuiv(gl14Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl14GetQueryiv(gl14Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl14IsBuffer(gl14Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl14IsQuery(gl14Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl14MapBuffer(gl14Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl14UnmapBuffer(gl14Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl14AttachShader(gl14Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl14BindAttribLocation(gl14Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl14BlendEquationSeperate(gl14Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl14CompileShader(gl14Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl14CreateProgram(gl14Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl14CreateShader(gl14Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl14DeleteProgram(gl14Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl14DeleteShader(gl14Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl14DetachShader(gl14Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl14EnableVertexAttribArray(gl14Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl14DisableVertexAttribArray(gl14Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl14DrawBuffers(gl14Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl14GetActiveAttrib(gl14Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl14GetActiveUniform(gl14Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl14GetAttachedShaders(gl14Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl14GetAttribLocation(gl14Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl14GetProgramiv(gl14Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl14GetProgramInfoLog(gl14Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl14GetShaderiv(gl14Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl14GetShaderInfoLog(gl14Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl14GetShaderSource(gl14Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl14GetUniformfv(gl14Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl14GetUniformiv(gl14Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl14GetUniformLocation(gl14Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl14GetVertexAttribdv(gl14Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl14GetVertexAttribfv(gl14Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl14GetVertexAttribiv(gl14Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl14GetVertexAttribPointerv(gl14Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl14IsProgram(gl14Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl14IsShader(gl14Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl14LinkProgram(gl14Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl14ShaderSource(gl14Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl14StencilFuncSeparate(gl14Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl14StencilMaskSeparate(gl14Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl14StencilOpSeparate(gl14Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl14Uniform1f(gl14Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl14Uniform2f(gl14Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl14Uniform3f(gl14Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl14Uniform4f(gl14Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl14Uniform1i(gl14Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl14Uniform2i(gl14Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl14Uniform3i(gl14Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl14Uniform4i(gl14Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl14Uniform1fv(gl14Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl14Uniform2fv(gl14Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl14Uniform3fv(gl14Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl14Uniform4fv(gl14Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl14Uniform1iv(gl14Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl14Uniform2iv(gl14Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl14Uniform3iv(gl14Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl14Uniform4iv(gl14Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl14UseProgram(gl14Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl14ValidateProgram(gl14Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl14VertexAttribPointer(gl14Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl14VertexAttrib1f(gl14Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl14VertexAttrib1s(gl14Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl14VertexAttrib1d(gl14Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl14VertexAttrib2f(gl14Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl14VertexAttrib2s(gl14Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl14VertexAttrib2d(gl14Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl14VertexAttrib3f(gl14Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl14VertexAttrib3s(gl14Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl14VertexAttrib3d(gl14Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl14VertexAttrib4f(gl14Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl14VertexAttrib4s(gl14Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl14VertexAttrib4d(gl14Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl14VertexAttrib4Nuv(gl14Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl14VertexAttrib1fv(gl14Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl14VertexAttrib1sv(gl14Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl14VertexAttrib1dv(gl14Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl14VertexAttrib2fv(gl14Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl14VertexAttrib2sv(gl14Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl14VertexAttrib2dv(gl14Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl14VertexAttrib3fv(gl14Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl14VertexAttrib3sv(gl14Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl14VertexAttrib3dv(gl14Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl14VertexAttrib4fv(gl14Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl14VertexAttrib4sv(gl14Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl14VertexAttrib4dv(gl14Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl14VertexAttrib4iv(gl14Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl14VertexAttrib4bv(gl14Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl14VertexAttrib4ubv(gl14Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl14VertexAttrib4usv(gl14Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl14VertexAttrib4uiv(gl14Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl14VertexAttrib4Nbv(gl14Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl14VertexAttrib4Nsv(gl14Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl14VertexAttrib4Niv(gl14Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl14VertexAttrib4Nubv(gl14Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl14VertexAttrib4Nusv(gl14Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl14VertexAttrib4Nuiv(gl14Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl14UniformMatrix2fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl14UniformMatrix3fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl14UniformMatrix4fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl14UniformMatrix2x3fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl14UniformMatrix3x2fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl14UniformMatrix2x4fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl14UniformMatrix4x2fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl14UniformMatrix3x4fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl14UniformMatrix4x3fv(gl14Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl14Context* gl14NewContext() {
    gl14Context* glc = calloc(1, sizeof(gl14Context));

    // Preload all procedures
    glc->fnAccum = (gl14PAccum)gl14LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl14PAlphaFunc)gl14LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl14PBegin)gl14LibGetProcAddress("glBegin");
    glc->fnEnd = (gl14PEnd)gl14LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl14PBitmap)gl14LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl14PBlendFunc)gl14LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl14PCallList)gl14LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl14PCallLists)gl14LibGetProcAddress("glCallLists");
    glc->fnClear = (gl14PClear)gl14LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl14PClearAccum)gl14LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl14PClearColor)gl14LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl14PClearDepth)gl14LibGetProcAddress("glClearDepth");
    glc->fnClearDepthf = (gl14PClearDepthf)gl14LibGetProcAddress("glClearDepthf");
    glc->fnClearIndex = (gl14PClearIndex)gl14LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl14PClearStencil)gl14LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl14PClipPlane)gl14LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl14PColor3b)gl14LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl14PColor3d)gl14LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl14PColor3f)gl14LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl14PColor3i)gl14LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl14PColor3s)gl14LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl14PColor3ub)gl14LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl14PColor3ui)gl14LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl14PColor3us)gl14LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl14PColor4b)gl14LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl14PColor4d)gl14LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl14PColor4f)gl14LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl14PColor4i)gl14LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl14PColor4s)gl14LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl14PColor4ub)gl14LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl14PColor4ui)gl14LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl14PColor4us)gl14LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl14PColor3bv)gl14LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl14PColor3dv)gl14LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl14PColor3fv)gl14LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl14PColor3iv)gl14LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl14PColor3sv)gl14LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl14PColor3ubv)gl14LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl14PColor3uiv)gl14LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl14PColor3usv)gl14LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl14PColor4bv)gl14LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl14PColor4dv)gl14LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl14PColor4fv)gl14LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl14PColor4iv)gl14LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl14PColor4sv)gl14LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl14PColor4ubv)gl14LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl14PColor4uiv)gl14LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl14PColor4usv)gl14LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl14PColorMask)gl14LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl14PColorMaterial)gl14LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl14PCopyPixels)gl14LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl14PCullFace)gl14LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl14PDeleteLists)gl14LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl14PDepthFunc)gl14LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl14PDepthMask)gl14LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl14PDepthRange)gl14LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl14PEnable)gl14LibGetProcAddress("glEnable");
    glc->fnDisable = (gl14PDisable)gl14LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl14PDrawBuffer)gl14LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl14PDrawPixels)gl14LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl14PEdgeFlag)gl14LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl14PEdgeFlagv)gl14LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl14PEdgeFlagPointer)gl14LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl14PEvalCoord1d)gl14LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl14PEvalCoord1f)gl14LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl14PEvalCoord2d)gl14LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl14PEvalCoord2f)gl14LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl14PEvalCoord1dv)gl14LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl14PEvalCoord1fv)gl14LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl14PEvalCoord2dv)gl14LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl14PEvalCoord2fv)gl14LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl14PEvalMesh1)gl14LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl14PEvalMesh2)gl14LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl14PEvalPoint1)gl14LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl14PEvalPoint2)gl14LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl14PFeedbackBuffer)gl14LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl14PFinish)gl14LibGetProcAddress("glFinish");
    glc->fnFlush = (gl14PFlush)gl14LibGetProcAddress("glFlush");
    glc->fnFogf = (gl14PFogf)gl14LibGetProcAddress("glFogf");
    glc->fnFogi = (gl14PFogi)gl14LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl14PFogfv)gl14LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl14PFogiv)gl14LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl14PFrontFace)gl14LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl14PFrustum)gl14LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl14PGenLists)gl14LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl14PGetBooleanv)gl14LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl14PGetDoublev)gl14LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl14PGetFloatv)gl14LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl14PGetIntegerv)gl14LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl14PGetClipPlane)gl14LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl14PGetError)gl14LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl14PGetLightfv)gl14LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl14PGetLightiv)gl14LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl14PGetMapdv)gl14LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl14PGetMapfv)gl14LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl14PGetMapiv)gl14LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl14PGetMaterialfv)gl14LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl14PGetMaterialiv)gl14LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl14PGetPixelMapfv)gl14LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl14PGetPixelMapuiv)gl14LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl14PGetPixelMapusv)gl14LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl14PGetPolygonStipple)gl14LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl14PGetString)gl14LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl14PGetTexEnvfv)gl14LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl14PGetTexEnviv)gl14LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl14PGetTexGendv)gl14LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl14PGetTexGenfv)gl14LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl14PGetTexGeniv)gl14LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl14PGetTexImage)gl14LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl14PGetTexLevelParameterfv)gl14LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl14PGetTexLevelParameteriv)gl14LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl14PGetTexParameterfv)gl14LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl14PGetTexParameteriv)gl14LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl14PHint)gl14LibGetProcAddress("glHint");
    glc->fnIndexd = (gl14PIndexd)gl14LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl14PIndexf)gl14LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl14PIndexi)gl14LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl14PIndexs)gl14LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl14PIndexdv)gl14LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl14PIndexfv)gl14LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl14PIndexiv)gl14LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl14PIndexsv)gl14LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl14PIndexMask)gl14LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl14PIndexPointer)gl14LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl14PInitNames)gl14LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl14PIsEnabled)gl14LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl14PIsList)gl14LibGetProcAddress("glIsList");
    glc->fnLightf = (gl14PLightf)gl14LibGetProcAddress("glLightf");
    glc->fnLighti = (gl14PLighti)gl14LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl14PLightfv)gl14LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl14PLightiv)gl14LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl14PLightModelf)gl14LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl14PLightModeli)gl14LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl14PLightModelfv)gl14LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl14PLightModeliv)gl14LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl14PLineStipple)gl14LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl14PLineWidth)gl14LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl14PListBase)gl14LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl14PLoadIdentity)gl14LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl14PLoadMatrixd)gl14LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl14PLoadMatrixf)gl14LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl14PLoadName)gl14LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl14PLogicOp)gl14LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl14PMap1d)gl14LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl14PMap1f)gl14LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl14PMap2d)gl14LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl14PMap2f)gl14LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl14PMapGrid1d)gl14LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl14PMapGrid1f)gl14LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl14PMapGrid2d)gl14LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl14PMapGrid2f)gl14LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl14PMaterialf)gl14LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl14PMateriali)gl14LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl14PMaterialfv)gl14LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl14PMaterialiv)gl14LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl14PMatrixMode)gl14LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl14PMultMatrixd)gl14LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl14PMultMatrixf)gl14LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl14PNewList)gl14LibGetProcAddress("glNewList");
    glc->fnEndList = (gl14PEndList)gl14LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl14PNormal3b)gl14LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl14PNormal3d)gl14LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl14PNormal3f)gl14LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl14PNormal3i)gl14LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl14PNormal3s)gl14LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl14PNormal3bv)gl14LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl14PNormal3dv)gl14LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl14PNormal3fv)gl14LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl14PNormal3iv)gl14LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl14PNormal3sv)gl14LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl14POrtho)gl14LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl14PPassThrough)gl14LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl14PPixelMapfv)gl14LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl14PPixelMapuiv)gl14LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl14PPixelMapusv)gl14LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl14PPixelStoref)gl14LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl14PPixelStorei)gl14LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl14PPixelTransferf)gl14LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl14PPixelTransferi)gl14LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl14PPixelZoom)gl14LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl14PPointSize)gl14LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl14PPolygonMode)gl14LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl14PPolygonStipple)gl14LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl14PPushAttrib)gl14LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl14PPopAttrib)gl14LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl14PPushMatrix)gl14LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl14PPopMatrix)gl14LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl14PPushName)gl14LibGetProcAddress("glPushName");
    glc->fnPopName = (gl14PPopName)gl14LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl14PRasterPos2d)gl14LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl14PRasterPos2f)gl14LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl14PRasterPos2i)gl14LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl14PRasterPos2s)gl14LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl14PRasterPos3d)gl14LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl14PRasterPos3f)gl14LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl14PRasterPos3i)gl14LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl14PRasterPos3s)gl14LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl14PRasterPos4d)gl14LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl14PRasterPos4f)gl14LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl14PRasterPos4i)gl14LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl14PRasterPos4s)gl14LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl14PRasterPos2dv)gl14LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl14PRasterPos2fv)gl14LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl14PRasterPos2iv)gl14LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl14PRasterPos2sv)gl14LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl14PRasterPos3dv)gl14LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl14PRasterPos3fv)gl14LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl14PRasterPos3iv)gl14LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl14PRasterPos3sv)gl14LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl14PRasterPos4dv)gl14LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl14PRasterPos4fv)gl14LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl14PRasterPos4iv)gl14LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl14PRasterPos4sv)gl14LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl14PReadBuffer)gl14LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl14PReadPixels)gl14LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl14PRectd)gl14LibGetProcAddress("glRectd");
    glc->fnRectf = (gl14PRectf)gl14LibGetProcAddress("glRectf");
    glc->fnRecti = (gl14PRecti)gl14LibGetProcAddress("glRecti");
    glc->fnRects = (gl14PRects)gl14LibGetProcAddress("glRects");
    glc->fnRectdv = (gl14PRectdv)gl14LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl14PRectfv)gl14LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl14PRectiv)gl14LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl14PRectsv)gl14LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl14PRenderMode)gl14LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl14PRotated)gl14LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl14PRotatef)gl14LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl14PScaled)gl14LibGetProcAddress("glScaled");
    glc->fnScalef = (gl14PScalef)gl14LibGetProcAddress("glScalef");
    glc->fnScissor = (gl14PScissor)gl14LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl14PSelectBuffer)gl14LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl14PShadeModel)gl14LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl14PStencilFunc)gl14LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl14PStencilMask)gl14LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl14PStencilOp)gl14LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl14PTexCoord1d)gl14LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl14PTexCoord1f)gl14LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl14PTexCoord1i)gl14LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl14PTexCoord1s)gl14LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl14PTexCoord2d)gl14LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl14PTexCoord2f)gl14LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl14PTexCoord2i)gl14LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl14PTexCoord2s)gl14LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl14PTexCoord3d)gl14LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl14PTexCoord3f)gl14LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl14PTexCoord3i)gl14LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl14PTexCoord3s)gl14LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl14PTexCoord4d)gl14LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl14PTexCoord4f)gl14LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl14PTexCoord4i)gl14LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl14PTexCoord4s)gl14LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl14PTexCoord1dv)gl14LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl14PTexCoord1fv)gl14LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl14PTexCoord1iv)gl14LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl14PTexCoord1sv)gl14LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl14PTexCoord2dv)gl14LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl14PTexCoord2fv)gl14LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl14PTexCoord2iv)gl14LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl14PTexCoord2sv)gl14LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl14PTexCoord3dv)gl14LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl14PTexCoord3fv)gl14LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl14PTexCoord3iv)gl14LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl14PTexCoord3sv)gl14LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl14PTexCoord4dv)gl14LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl14PTexCoord4fv)gl14LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl14PTexCoord4iv)gl14LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl14PTexCoord4sv)gl14LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl14PTexEnvf)gl14LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl14PTexEnvi)gl14LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl14PTexEnvfv)gl14LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl14PTexEnviv)gl14LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl14PTexGend)gl14LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl14PTexGenf)gl14LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl14PTexGeni)gl14LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl14PTexGendv)gl14LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl14PTexGenfv)gl14LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl14PTexGeniv)gl14LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl14PTexImage1D)gl14LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl14PTexImage2D)gl14LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl14PTexParameterf)gl14LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl14PTexParameteri)gl14LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl14PTexParameterfv)gl14LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl14PTexParameteriv)gl14LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl14PTranslated)gl14LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl14PTranslatef)gl14LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl14PVertex2s)gl14LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl14PVertex2i)gl14LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl14PVertex2f)gl14LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl14PVertex2d)gl14LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl14PVertex3s)gl14LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl14PVertex3i)gl14LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl14PVertex3f)gl14LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl14PVertex3d)gl14LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl14PVertex4s)gl14LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl14PVertex4i)gl14LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl14PVertex4f)gl14LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl14PVertex4d)gl14LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl14PViewport)gl14LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl14PGetConvolutionParameterfv)gl14LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl14PGetConvolutionParameteriv)gl14LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl14PAreTexturesResident)gl14LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl14PArrayElement)gl14LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl14PDrawArrays)gl14LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl14PDrawElements)gl14LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl14PGetPointerv)gl14LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl14PPolygonOffset)gl14LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl14PCopyTexImage1D)gl14LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl14PCopyTexImage2D)gl14LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl14PCopyTexSubImage1D)gl14LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl14PCopyTexSubImage2D)gl14LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl14PBindTexture)gl14LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl14PDeleteTextures)gl14LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl14PGenTextures)gl14LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl14PIsTexture)gl14LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl14PColorPointer)gl14LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl14PEnableClientState)gl14LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl14PDisableClientState)gl14LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl14PIndexub)gl14LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl14PIndexubv)gl14LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl14PInterleavedArrays)gl14LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl14PNormalPointer)gl14LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl14PPushClientAttrib)gl14LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl14PPrioritizeTextures)gl14LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl14PPopClientAttrib)gl14LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl14PTexCoordPointer)gl14LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl14PTexSubImage1D)gl14LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl14PTexSubImage2D)gl14LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl14PVertexPointer)gl14LibGetProcAddress("glVertexPointer");
    glc->fnColorTable = (gl14PColorTable)gl14GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl14PColorTableParameterfv)gl14GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl14PColorTableParameteriv)gl14GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl14PColorSubTable)gl14GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl14PConvolutionFilter1D)gl14GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl14PConvolutionFilter2D)gl14GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl14PConvolutionParameterf)gl14GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl14PConvolutionParameteri)gl14GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl14PCopyColorTable)gl14GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl14PCopyColorSubTable)gl14GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl14PCopyConvolutionFilter1D)gl14GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl14PCopyConvolutionFilter2D)gl14GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl14PGetColorTable)gl14GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl14PGetColorTableParameterfv)gl14GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl14PGetColorTableParameteriv)gl14GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl14PGetConvolutionFilter)gl14GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl14PGetHistogram)gl14GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl14PGetHistogramParameterfv)gl14GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl14PGetHistogramParameteriv)gl14GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl14PGetSeparableFilter)gl14GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl14PHistogram)gl14GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl14PMinmax)gl14GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl14PMultiTexCoord1s)gl14GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl14PMultiTexCoord1i)gl14GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl14PMultiTexCoord1f)gl14GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl14PMultiTexCoord1d)gl14GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl14PMultiTexCoord2s)gl14GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl14PMultiTexCoord2i)gl14GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl14PMultiTexCoord2f)gl14GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl14PMultiTexCoord2d)gl14GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl14PMultiTexCoord3s)gl14GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl14PMultiTexCoord3i)gl14GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl14PMultiTexCoord3f)gl14GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl14PMultiTexCoord3d)gl14GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl14PMultiTexCoord4s)gl14GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl14PMultiTexCoord4i)gl14GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl14PMultiTexCoord4f)gl14GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl14PMultiTexCoord4d)gl14GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl14PMultiTexCoord1sv)gl14GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl14PMultiTexCoord1iv)gl14GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl14PMultiTexCoord1fv)gl14GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl14PMultiTexCoord1dv)gl14GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl14PMultiTexCoord2sv)gl14GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl14PMultiTexCoord2iv)gl14GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl14PMultiTexCoord2fv)gl14GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl14PMultiTexCoord2dv)gl14GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl14PMultiTexCoord3sv)gl14GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl14PMultiTexCoord3iv)gl14GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl14PMultiTexCoord3fv)gl14GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl14PMultiTexCoord3dv)gl14GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl14PMultiTexCoord4sv)gl14GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl14PMultiTexCoord4iv)gl14GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl14PMultiTexCoord4fv)gl14GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl14PMultiTexCoord4dv)gl14GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl14PResetHistogram)gl14GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl14PResetMinmax)gl14GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl14PSeparableFilter2D)gl14GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl14PBlendColor)gl14GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl14PBlendEquation)gl14GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl14PCopyTexSubImage3D)gl14GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl14PDrawRangeElements)gl14GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl14PTexImage3D)gl14GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl14PTexSubImage3D)gl14GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl14PActiveTexture)gl14GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl14PClientActiveTexture)gl14GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl14PCompressedTexImage1D)gl14GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl14PCompressedTexImage2D)gl14GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl14PCompressedTexImage3D)gl14GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl14PCompressedTexSubImage1D)gl14GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl14PCompressedTexSubImage2D)gl14GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl14PCompressedTexSubImage3D)gl14GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl14PGetCompressedTexImage)gl14GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl14PLoadTransposeMatrixd)gl14GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl14PLoadTransposeMatrixf)gl14GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl14PMultTransposeMatrixd)gl14GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl14PMultTransposeMatrixf)gl14GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl14PSampleCoverage)gl14GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl14PBlendFuncSeparate)gl14GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl14PFogCoordPointer)gl14GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl14PFogCoordd)gl14GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl14PFogCoordf)gl14GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl14PFogCoorddv)gl14GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl14PFogCoordfv)gl14GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl14PMultiDrawArrays)gl14GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl14PMultiDrawElements)gl14GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl14PPointParameterf)gl14GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl14PPointParameteri)gl14GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl14PSecondaryColor3b)gl14GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl14PSecondaryColor3s)gl14GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl14PSecondaryColor3i)gl14GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl14PSecondaryColor3f)gl14GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl14PSecondaryColor3d)gl14GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl14PSecondaryColor3ub)gl14GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl14PSecondaryColor3us)gl14GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl14PSecondaryColor3ui)gl14GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl14PSecondaryColor3bv)gl14GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl14PSecondaryColor3sv)gl14GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl14PSecondaryColor3iv)gl14GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl14PSecondaryColor3fv)gl14GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl14PSecondaryColor3dv)gl14GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl14PSecondaryColor3ubv)gl14GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl14PSecondaryColor3usv)gl14GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl14PSecondaryColor3uiv)gl14GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl14PSecondaryColorPointer)gl14GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl14PWindowPos2s)gl14GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl14PWindowPos2i)gl14GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl14PWindowPos2f)gl14GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl14PWindowPos2d)gl14GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl14PWindowPos3s)gl14GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl14PWindowPos3i)gl14GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl14PWindowPos3f)gl14GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl14PWindowPos3d)gl14GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl14PWindowPos2sv)gl14GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl14PWindowPos2iv)gl14GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl14PWindowPos2fv)gl14GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl14PWindowPos2dv)gl14GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl14PWindowPos3sv)gl14GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl14PWindowPos3iv)gl14GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl14PWindowPos3fv)gl14GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl14PWindowPos3dv)gl14GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl14PBeginQuery)gl14GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl14PBindBuffer)gl14GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl14PBufferData)gl14GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl14PBufferSubData)gl14GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl14PDeleteBuffers)gl14GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl14PDeleteQueries)gl14GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl14PGenBuffers)gl14GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl14PGenQueries)gl14GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl14PGetBufferParameteriv)gl14GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl14PGetBufferPointerv)gl14GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl14PGetBufferSubData)gl14GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl14PGetQueryObjectiv)gl14GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl14PGetQueryObjectuiv)gl14GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl14PGetQueryiv)gl14GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl14PIsBuffer)gl14GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl14PIsQuery)gl14GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl14PMapBuffer)gl14GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl14PUnmapBuffer)gl14GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl14PAttachShader)gl14GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl14PBindAttribLocation)gl14GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl14PBlendEquationSeperate)gl14GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl14PCompileShader)gl14GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl14PCreateProgram)gl14GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl14PCreateShader)gl14GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl14PDeleteProgram)gl14GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl14PDeleteShader)gl14GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl14PDetachShader)gl14GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl14PEnableVertexAttribArray)gl14GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl14PDisableVertexAttribArray)gl14GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl14PDrawBuffers)gl14LibGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl14PGetActiveAttrib)gl14LibGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl14PGetActiveUniform)gl14LibGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl14PGetAttachedShaders)gl14LibGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl14PGetAttribLocation)gl14LibGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl14PGetProgramiv)gl14LibGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl14PGetProgramInfoLog)gl14LibGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl14PGetShaderiv)gl14LibGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl14PGetShaderInfoLog)gl14LibGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl14PGetShaderSource)gl14LibGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl14PGetUniformfv)gl14LibGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl14PGetUniformiv)gl14LibGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl14PGetUniformLocation)gl14LibGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl14PGetVertexAttribdv)gl14LibGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl14PGetVertexAttribfv)gl14LibGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl14PGetVertexAttribiv)gl14LibGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl14PGetVertexAttribPointerv)gl14LibGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl14PIsProgram)gl14LibGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl14PIsShader)gl14LibGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl14PLinkProgram)gl14LibGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl14PShaderSource)gl14LibGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl14PStencilFuncSeparate)gl14LibGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl14PStencilMaskSeparate)gl14LibGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl14PStencilOpSeparate)gl14LibGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl14PUniform1f)gl14LibGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl14PUniform2f)gl14LibGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl14PUniform3f)gl14LibGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl14PUniform4f)gl14LibGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl14PUniform1i)gl14LibGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl14PUniform2i)gl14LibGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl14PUniform3i)gl14LibGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl14PUniform4i)gl14LibGetProcAddress("glUniform4i");
    glc->fnUniform1fv = (gl14PUniform1fv)gl14LibGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl14PUniform2fv)gl14LibGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl14PUniform3fv)gl14LibGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl14PUniform4fv)gl14LibGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl14PUniform1iv)gl14LibGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl14PUniform2iv)gl14LibGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl14PUniform3iv)gl14LibGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl14PUniform4iv)gl14LibGetProcAddress("glUniform4iv");
    glc->fnUseProgram = (gl14PUseProgram)gl14LibGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl14PValidateProgram)gl14LibGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl14PVertexAttribPointer)gl14LibGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl14PVertexAttrib1f)gl14LibGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl14PVertexAttrib1s)gl14LibGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl14PVertexAttrib1d)gl14LibGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl14PVertexAttrib2f)gl14LibGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl14PVertexAttrib2s)gl14LibGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl14PVertexAttrib2d)gl14LibGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl14PVertexAttrib3f)gl14LibGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl14PVertexAttrib3s)gl14LibGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl14PVertexAttrib3d)gl14LibGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl14PVertexAttrib4f)gl14LibGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl14PVertexAttrib4s)gl14LibGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl14PVertexAttrib4d)gl14LibGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl14PVertexAttrib4Nuv)gl14LibGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl14PVertexAttrib1fv)gl14LibGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl14PVertexAttrib1sv)gl14LibGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl14PVertexAttrib1dv)gl14LibGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl14PVertexAttrib2fv)gl14LibGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl14PVertexAttrib2sv)gl14LibGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl14PVertexAttrib2dv)gl14LibGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl14PVertexAttrib3fv)gl14LibGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl14PVertexAttrib3sv)gl14LibGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl14PVertexAttrib3dv)gl14LibGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl14PVertexAttrib4fv)gl14LibGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl14PVertexAttrib4sv)gl14LibGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl14PVertexAttrib4dv)gl14LibGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl14PVertexAttrib4iv)gl14LibGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl14PVertexAttrib4bv)gl14LibGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl14PVertexAttrib4ubv)gl14LibGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl14PVertexAttrib4usv)gl14LibGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl14PVertexAttrib4uiv)gl14LibGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl14PVertexAttrib4Nbv)gl14LibGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl14PVertexAttrib4Nsv)gl14LibGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl14PVertexAttrib4Niv)gl14LibGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl14PVertexAttrib4Nubv)gl14LibGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl14PVertexAttrib4Nusv)gl14LibGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl14PVertexAttrib4Nuiv)gl14LibGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl14PUniformMatrix2fv)gl14LibGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl14PUniformMatrix3fv)gl14LibGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl14PUniformMatrix4fv)gl14LibGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl14PUniformMatrix2x3fv)gl14LibGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl14PUniformMatrix3x2fv)gl14LibGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl14PUniformMatrix2x4fv)gl14LibGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl14PUniformMatrix4x2fv)gl14LibGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl14PUniformMatrix3x4fv)gl14LibGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl14PUniformMatrix4x3fv)gl14LibGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

