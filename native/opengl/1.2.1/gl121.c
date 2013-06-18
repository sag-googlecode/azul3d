
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl121.h"

#ifdef _WIN32
HMODULE gl121OpenGL32;

void* gl121LibGetProcAddress(char* name) {
	if(gl121OpenGL32 == NULL) {
		gl121OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl121OpenGL32, TEXT(name));
}

void* gl121GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl121Accum(gl121Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl121AlphaFunc(gl121Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl121Begin(gl121Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl121End(gl121Context* glc) {
    return glc->fnEnd();
}

void gl121Bitmap(gl121Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl121BlendFunc(gl121Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl121CallList(gl121Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl121CallLists(gl121Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl121Clear(gl121Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl121ClearAccum(gl121Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl121ClearColor(gl121Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl121ClearDepth(gl121Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl121ClearIndex(gl121Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl121ClearStencil(gl121Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl121ClipPlane(gl121Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl121Color3b(gl121Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl121Color3d(gl121Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl121Color3f(gl121Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl121Color3i(gl121Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl121Color3s(gl121Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl121Color3ub(gl121Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl121Color3ui(gl121Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl121Color3us(gl121Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl121Color4b(gl121Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl121Color4d(gl121Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl121Color4f(gl121Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl121Color4i(gl121Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl121Color4s(gl121Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl121Color4ub(gl121Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl121Color4ui(gl121Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl121Color4us(gl121Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl121Color3bv(gl121Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl121Color3dv(gl121Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl121Color3fv(gl121Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl121Color3iv(gl121Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl121Color3sv(gl121Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl121Color3ubv(gl121Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl121Color3uiv(gl121Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl121Color3usv(gl121Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl121Color4bv(gl121Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl121Color4dv(gl121Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl121Color4fv(gl121Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl121Color4iv(gl121Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl121Color4sv(gl121Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl121Color4ubv(gl121Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl121Color4uiv(gl121Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl121Color4usv(gl121Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl121ColorMask(gl121Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl121ColorMaterial(gl121Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl121CopyPixels(gl121Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl121CullFace(gl121Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl121DeleteLists(gl121Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl121DepthFunc(gl121Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl121DepthMask(gl121Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl121DepthRange(gl121Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl121Enable(gl121Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl121Disable(gl121Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl121DrawBuffer(gl121Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl121DrawPixels(gl121Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl121EdgeFlag(gl121Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl121EdgeFlagv(gl121Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl121EdgeFlagPointer(gl121Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl121EvalCoord1d(gl121Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl121EvalCoord1f(gl121Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl121EvalCoord2d(gl121Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl121EvalCoord2f(gl121Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl121EvalCoord1dv(gl121Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl121EvalCoord1fv(gl121Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl121EvalCoord2dv(gl121Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl121EvalCoord2fv(gl121Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl121EvalMesh1(gl121Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl121EvalMesh2(gl121Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl121EvalPoint1(gl121Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl121EvalPoint2(gl121Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl121FeedbackBuffer(gl121Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl121Finish(gl121Context* glc) {
    return glc->fnFinish();
}

void gl121Flush(gl121Context* glc) {
    return glc->fnFlush();
}

void gl121Fogf(gl121Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl121Fogi(gl121Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl121Fogfv(gl121Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl121Fogiv(gl121Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl121FrontFace(gl121Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl121Frustum(gl121Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl121GenLists(gl121Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl121GetBooleanv(gl121Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl121GetDoublev(gl121Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl121GetFloatv(gl121Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl121GetIntegerv(gl121Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl121GetClipPlane(gl121Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl121GetError(gl121Context* glc) {
    return glc->fnGetError();
}

void gl121GetLightfv(gl121Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl121GetLightiv(gl121Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl121GetMapdv(gl121Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl121GetMapfv(gl121Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl121GetMapiv(gl121Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl121GetMaterialfv(gl121Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl121GetMaterialiv(gl121Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl121GetPixelMapfv(gl121Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl121GetPixelMapuiv(gl121Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl121GetPixelMapusv(gl121Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl121GetPolygonStipple(gl121Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl121GetString(gl121Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl121GetTexEnvfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl121GetTexEnviv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl121GetTexGendv(gl121Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl121GetTexGenfv(gl121Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl121GetTexGeniv(gl121Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl121GetTexImage(gl121Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl121GetTexLevelParameterfv(gl121Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl121GetTexLevelParameteriv(gl121Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl121GetTexParameterfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl121GetTexParameteriv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl121Hint(gl121Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl121Indexd(gl121Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl121Indexf(gl121Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl121Indexi(gl121Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl121Indexs(gl121Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl121Indexdv(gl121Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl121Indexfv(gl121Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl121Indexiv(gl121Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl121Indexsv(gl121Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl121IndexMask(gl121Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl121IndexPointer(gl121Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl121InitNames(gl121Context* glc) {
    return glc->fnInitNames();
}

void gl121IsEnabled(gl121Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl121IsList(gl121Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl121Lightf(gl121Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl121Lighti(gl121Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl121Lightfv(gl121Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl121Lightiv(gl121Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl121LightModelf(gl121Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl121LightModeli(gl121Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl121LightModelfv(gl121Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl121LightModeliv(gl121Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl121LineStipple(gl121Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl121LineWidth(gl121Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl121ListBase(gl121Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl121LoadIdentity(gl121Context* glc) {
    return glc->fnLoadIdentity();
}

void gl121LoadMatrixd(gl121Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl121LoadMatrixf(gl121Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl121LoadName(gl121Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl121LogicOp(gl121Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl121Map1d(gl121Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl121Map1f(gl121Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl121Map2d(gl121Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl121Map2f(gl121Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl121MapGrid1d(gl121Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl121MapGrid1f(gl121Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl121MapGrid2d(gl121Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl121MapGrid2f(gl121Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl121Materialf(gl121Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl121Materiali(gl121Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl121Materialfv(gl121Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl121Materialiv(gl121Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl121MatrixMode(gl121Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl121MultMatrixd(gl121Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl121MultMatrixf(gl121Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl121NewList(gl121Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl121EndList(gl121Context* glc) {
    return glc->fnEndList();
}

void gl121Normal3b(gl121Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl121Normal3d(gl121Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl121Normal3f(gl121Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl121Normal3i(gl121Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl121Normal3s(gl121Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl121Normal3bv(gl121Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl121Normal3dv(gl121Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl121Normal3fv(gl121Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl121Normal3iv(gl121Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl121Normal3sv(gl121Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl121Ortho(gl121Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl121PassThrough(gl121Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl121PixelMapfv(gl121Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl121PixelMapuiv(gl121Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl121PixelMapusv(gl121Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl121PixelStoref(gl121Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl121PixelStorei(gl121Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl121PixelTransferf(gl121Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl121PixelTransferi(gl121Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl121PixelZoom(gl121Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl121PointSize(gl121Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl121PolygonMode(gl121Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl121PolygonStipple(gl121Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl121PushAttrib(gl121Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl121PopAttrib(gl121Context* glc) {
    return glc->fnPopAttrib();
}

void gl121PushMatrix(gl121Context* glc) {
    return glc->fnPushMatrix();
}

void gl121PopMatrix(gl121Context* glc) {
    return glc->fnPopMatrix();
}

void gl121PushName(gl121Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl121PopName(gl121Context* glc) {
    return glc->fnPopName();
}

void gl121RasterPos2d(gl121Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl121RasterPos2f(gl121Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl121RasterPos2i(gl121Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl121RasterPos2s(gl121Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl121RasterPos3d(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl121RasterPos3f(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl121RasterPos3i(gl121Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl121RasterPos3s(gl121Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl121RasterPos4d(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl121RasterPos4f(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl121RasterPos4i(gl121Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl121RasterPos4s(gl121Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl121RasterPos2dv(gl121Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl121RasterPos2fv(gl121Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl121RasterPos2iv(gl121Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl121RasterPos2sv(gl121Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl121RasterPos3dv(gl121Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl121RasterPos3fv(gl121Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl121RasterPos3iv(gl121Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl121RasterPos3sv(gl121Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl121RasterPos4dv(gl121Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl121RasterPos4fv(gl121Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl121RasterPos4iv(gl121Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl121RasterPos4sv(gl121Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl121ReadBuffer(gl121Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl121ReadPixels(gl121Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl121Rectd(gl121Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl121Rectf(gl121Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl121Recti(gl121Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl121Rects(gl121Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl121Rectdv(gl121Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl121Rectfv(gl121Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl121Rectiv(gl121Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl121Rectsv(gl121Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl121RenderMode(gl121Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl121Rotated(gl121Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl121Rotatef(gl121Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl121Scaled(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl121Scalef(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl121Scissor(gl121Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl121SelectBuffer(gl121Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl121ShadeModel(gl121Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl121StencilFunc(gl121Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl121StencilMask(gl121Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl121StencilOp(gl121Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl121TexCoord1d(gl121Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl121TexCoord1f(gl121Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl121TexCoord1i(gl121Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl121TexCoord1s(gl121Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl121TexCoord2d(gl121Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl121TexCoord2f(gl121Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl121TexCoord2i(gl121Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl121TexCoord2s(gl121Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl121TexCoord3d(gl121Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl121TexCoord3f(gl121Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl121TexCoord3i(gl121Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl121TexCoord3s(gl121Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl121TexCoord4d(gl121Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl121TexCoord4f(gl121Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl121TexCoord4i(gl121Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl121TexCoord4s(gl121Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl121TexCoord1dv(gl121Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl121TexCoord1fv(gl121Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl121TexCoord1iv(gl121Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl121TexCoord1sv(gl121Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl121TexCoord2dv(gl121Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl121TexCoord2fv(gl121Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl121TexCoord2iv(gl121Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl121TexCoord2sv(gl121Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl121TexCoord3dv(gl121Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl121TexCoord3fv(gl121Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl121TexCoord3iv(gl121Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl121TexCoord3sv(gl121Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl121TexCoord4dv(gl121Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl121TexCoord4fv(gl121Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl121TexCoord4iv(gl121Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl121TexCoord4sv(gl121Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl121TexEnvf(gl121Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl121TexEnvi(gl121Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl121TexEnvfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl121TexEnviv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl121TexGend(gl121Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl121TexGenf(gl121Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl121TexGeni(gl121Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl121TexGendv(gl121Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl121TexGenfv(gl121Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl121TexGeniv(gl121Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl121TexImage1D(gl121Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl121TexImage2D(gl121Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl121TexParameterf(gl121Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl121TexParameteri(gl121Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl121TexParameterfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl121TexParameteriv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl121Translated(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl121Translatef(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl121Vertex2s(gl121Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl121Vertex2i(gl121Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl121Vertex2f(gl121Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl121Vertex2d(gl121Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl121Vertex3s(gl121Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl121Vertex3i(gl121Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl121Vertex3f(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl121Vertex3d(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl121Vertex4s(gl121Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl121Vertex4i(gl121Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl121Vertex4f(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl121Vertex4d(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl121Viewport(gl121Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl121GetConvolutionParameterfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl121GetConvolutionParameteriv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl121AreTexturesResident(gl121Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl121ArrayElement(gl121Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl121DrawArrays(gl121Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl121DrawElements(gl121Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl121GetPointerv(gl121Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl121PolygonOffset(gl121Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl121CopyTexImage1D(gl121Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl121CopyTexImage2D(gl121Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl121CopyTexSubImage1D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl121CopyTexSubImage2D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl121BindTexture(gl121Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl121DeleteTextures(gl121Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl121GenTextures(gl121Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl121IsTexture(gl121Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl121ColorPointer(gl121Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl121EnableClientState(gl121Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl121DisableClientState(gl121Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl121Indexub(gl121Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl121Indexubv(gl121Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl121InterleavedArrays(gl121Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl121NormalPointer(gl121Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl121PushClientAttrib(gl121Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl121PrioritizeTextures(gl121Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl121PopClientAttrib(gl121Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl121TexCoordPointer(gl121Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl121TexSubImage1D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl121TexSubImage2D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl121VertexPointer(gl121Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl121ColorTable(gl121Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl121ColorTableParameterfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl121ColorTableParameteriv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl121ColorSubTable(gl121Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl121ConvolutionFilter1D(gl121Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl121ConvolutionFilter2D(gl121Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl121ConvolutionParameterf(gl121Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl121ConvolutionParameteri(gl121Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl121CopyColorTable(gl121Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl121CopyColorSubTable(gl121Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl121CopyConvolutionFilter1D(gl121Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl121CopyConvolutionFilter2D(gl121Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl121GetColorTable(gl121Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl121GetColorTableParameterfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl121GetColorTableParameteriv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl121GetConvolutionFilter(gl121Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl121GetHistogram(gl121Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl121GetHistogramParameterfv(gl121Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl121GetHistogramParameteriv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl121GetSeparableFilter(gl121Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl121Histogram(gl121Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl121Minmax(gl121Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl121MultiTexCoord1s(gl121Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl121MultiTexCoord1i(gl121Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl121MultiTexCoord1f(gl121Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl121MultiTexCoord1d(gl121Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl121MultiTexCoord2s(gl121Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl121MultiTexCoord2i(gl121Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl121MultiTexCoord2f(gl121Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl121MultiTexCoord2d(gl121Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl121MultiTexCoord3s(gl121Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl121MultiTexCoord3i(gl121Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl121MultiTexCoord3f(gl121Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl121MultiTexCoord3d(gl121Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl121MultiTexCoord4s(gl121Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl121MultiTexCoord4i(gl121Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl121MultiTexCoord4f(gl121Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl121MultiTexCoord4d(gl121Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl121MultiTexCoord1sv(gl121Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl121MultiTexCoord1iv(gl121Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl121MultiTexCoord1fv(gl121Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl121MultiTexCoord1dv(gl121Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl121MultiTexCoord2sv(gl121Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl121MultiTexCoord2iv(gl121Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl121MultiTexCoord2fv(gl121Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl121MultiTexCoord2dv(gl121Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl121MultiTexCoord3sv(gl121Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl121MultiTexCoord3iv(gl121Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl121MultiTexCoord3fv(gl121Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl121MultiTexCoord3dv(gl121Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl121MultiTexCoord4sv(gl121Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl121MultiTexCoord4iv(gl121Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl121MultiTexCoord4fv(gl121Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl121MultiTexCoord4dv(gl121Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl121ResetHistogram(gl121Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl121ResetMinmax(gl121Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl121SeparableFilter2D(gl121Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl121BlendColor(gl121Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl121BlendEquation(gl121Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl121CopyTexSubImage3D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl121DrawRangeElements(gl121Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl121TexImage3D(gl121Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl121TexSubImage3D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl121ActiveTexture(gl121Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl121ClientActiveTexture(gl121Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl121CompressedTexImage1D(gl121Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl121CompressedTexImage2D(gl121Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl121CompressedTexImage3D(gl121Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl121CompressedTexSubImage1D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl121CompressedTexSubImage2D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl121CompressedTexSubImage3D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl121GetCompressedTexImage(gl121Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl121LoadTransposeMatrixd(gl121Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl121LoadTransposeMatrixf(gl121Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl121MultTransposeMatrixd(gl121Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl121MultTransposeMatrixf(gl121Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl121SampleCoverage(gl121Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl121BlendFuncSeparate(gl121Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl121FogCoordPointer(gl121Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl121FogCoordd(gl121Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl121FogCoordf(gl121Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl121FogCoorddv(gl121Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl121FogCoordfv(gl121Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl121MultiDrawArrays(gl121Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl121MultiDrawElements(gl121Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl121PointParameterf(gl121Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl121PointParameteri(gl121Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl121SecondaryColor3b(gl121Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl121SecondaryColor3s(gl121Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl121SecondaryColor3i(gl121Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl121SecondaryColor3f(gl121Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl121SecondaryColor3d(gl121Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl121SecondaryColor3ub(gl121Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl121SecondaryColor3us(gl121Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl121SecondaryColor3ui(gl121Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl121SecondaryColor3bv(gl121Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl121SecondaryColor3sv(gl121Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl121SecondaryColor3iv(gl121Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl121SecondaryColor3fv(gl121Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl121SecondaryColor3dv(gl121Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl121SecondaryColor3ubv(gl121Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl121SecondaryColor3usv(gl121Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl121SecondaryColor3uiv(gl121Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl121SecondaryColorPointer(gl121Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl121WindowPos2s(gl121Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl121WindowPos2i(gl121Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl121WindowPos2f(gl121Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl121WindowPos2d(gl121Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl121WindowPos3s(gl121Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl121WindowPos3i(gl121Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl121WindowPos3f(gl121Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl121WindowPos3d(gl121Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl121WindowPos2sv(gl121Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl121WindowPos2iv(gl121Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl121WindowPos2fv(gl121Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl121WindowPos2dv(gl121Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl121WindowPos3sv(gl121Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl121WindowPos3iv(gl121Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl121WindowPos3fv(gl121Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl121WindowPos3dv(gl121Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl121BeginQuery(gl121Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl121BindBuffer(gl121Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl121BufferData(gl121Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl121BufferSubData(gl121Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl121DeleteBuffers(gl121Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl121DeleteQueries(gl121Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl121GenBuffers(gl121Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl121GenQueries(gl121Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl121GetBufferParameteriv(gl121Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl121GetBufferPointerv(gl121Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl121GetBufferSubData(gl121Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl121GetQueryObjectiv(gl121Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl121GetQueryObjectuiv(gl121Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl121GetQueryiv(gl121Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl121IsBuffer(gl121Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl121IsQuery(gl121Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl121MapBuffer(gl121Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl121UnmapBuffer(gl121Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl121AttachShader(gl121Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl121BindAttribLocation(gl121Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl121BlendEquationSeperate(gl121Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl121CompileShader(gl121Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl121CreateProgram(gl121Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl121CreateShader(gl121Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl121DeleteProgram(gl121Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl121DeleteShader(gl121Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl121DetachShader(gl121Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl121EnableVertexAttribArray(gl121Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl121DisableVertexAttribArray(gl121Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl121DrawBuffers(gl121Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl121GetActiveAttrib(gl121Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl121GetActiveUniform(gl121Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl121GetAttachedShaders(gl121Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl121GetAttribLocation(gl121Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl121GetProgramiv(gl121Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl121GetProgramInfoLog(gl121Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl121GetShaderiv(gl121Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl121GetShaderInfoLog(gl121Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl121GetShaderSource(gl121Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl121GetUniformfv(gl121Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl121GetUniformiv(gl121Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl121GetUniformLocation(gl121Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl121GetVertexAttribdv(gl121Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl121GetVertexAttribfv(gl121Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl121GetVertexAttribiv(gl121Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl121GetVertexAttribPointerv(gl121Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl121IsProgram(gl121Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl121IsShader(gl121Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl121LinkProgram(gl121Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl121ShaderSource(gl121Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl121StencilFuncSeparate(gl121Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl121StencilMaskSeparate(gl121Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl121StencilOpSeparate(gl121Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl121Uniform1f(gl121Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl121Uniform2f(gl121Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl121Uniform3f(gl121Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl121Uniform4f(gl121Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl121Uniform1i(gl121Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl121Uniform2i(gl121Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl121Uniform3i(gl121Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl121Uniform4i(gl121Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl121Uniform1fv(gl121Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl121Uniform2fv(gl121Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl121Uniform3fv(gl121Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl121Uniform4fv(gl121Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl121Uniform1iv(gl121Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl121Uniform2iv(gl121Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl121Uniform3iv(gl121Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl121Uniform4iv(gl121Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl121UseProgram(gl121Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl121ValidateProgram(gl121Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl121VertexAttribPointer(gl121Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl121VertexAttrib1f(gl121Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl121VertexAttrib1s(gl121Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl121VertexAttrib1d(gl121Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl121VertexAttrib2f(gl121Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl121VertexAttrib2s(gl121Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl121VertexAttrib2d(gl121Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl121VertexAttrib3f(gl121Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl121VertexAttrib3s(gl121Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl121VertexAttrib3d(gl121Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl121VertexAttrib4f(gl121Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl121VertexAttrib4s(gl121Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl121VertexAttrib4d(gl121Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl121VertexAttrib4Nuv(gl121Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl121VertexAttrib1fv(gl121Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl121VertexAttrib1sv(gl121Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl121VertexAttrib1dv(gl121Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl121VertexAttrib2fv(gl121Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl121VertexAttrib2sv(gl121Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl121VertexAttrib2dv(gl121Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl121VertexAttrib3fv(gl121Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl121VertexAttrib3sv(gl121Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl121VertexAttrib3dv(gl121Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl121VertexAttrib4fv(gl121Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl121VertexAttrib4sv(gl121Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl121VertexAttrib4dv(gl121Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl121VertexAttrib4iv(gl121Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl121VertexAttrib4bv(gl121Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl121VertexAttrib4ubv(gl121Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl121VertexAttrib4usv(gl121Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl121VertexAttrib4uiv(gl121Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl121VertexAttrib4Nbv(gl121Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl121VertexAttrib4Nsv(gl121Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl121VertexAttrib4Niv(gl121Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl121VertexAttrib4Nubv(gl121Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl121VertexAttrib4Nusv(gl121Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl121VertexAttrib4Nuiv(gl121Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl121UniformMatrix2fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl121UniformMatrix3fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl121UniformMatrix4fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl121UniformMatrix2x3fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl121UniformMatrix3x2fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl121UniformMatrix2x4fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl121UniformMatrix4x2fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl121UniformMatrix3x4fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl121UniformMatrix4x3fv(gl121Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl121Context* gl121NewContext() {
    gl121Context* glc = calloc(1, sizeof(gl121Context));

    // Preload all procedures
    glc->fnAccum = (gl121PAccum)gl121LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl121PAlphaFunc)gl121LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl121PBegin)gl121LibGetProcAddress("glBegin");
    glc->fnEnd = (gl121PEnd)gl121LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl121PBitmap)gl121LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl121PBlendFunc)gl121LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl121PCallList)gl121LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl121PCallLists)gl121LibGetProcAddress("glCallLists");
    glc->fnClear = (gl121PClear)gl121LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl121PClearAccum)gl121LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl121PClearColor)gl121LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl121PClearDepth)gl121LibGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl121PClearIndex)gl121LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl121PClearStencil)gl121LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl121PClipPlane)gl121LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl121PColor3b)gl121LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl121PColor3d)gl121LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl121PColor3f)gl121LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl121PColor3i)gl121LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl121PColor3s)gl121LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl121PColor3ub)gl121LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl121PColor3ui)gl121LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl121PColor3us)gl121LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl121PColor4b)gl121LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl121PColor4d)gl121LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl121PColor4f)gl121LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl121PColor4i)gl121LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl121PColor4s)gl121LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl121PColor4ub)gl121LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl121PColor4ui)gl121LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl121PColor4us)gl121LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl121PColor3bv)gl121LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl121PColor3dv)gl121LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl121PColor3fv)gl121LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl121PColor3iv)gl121LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl121PColor3sv)gl121LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl121PColor3ubv)gl121LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl121PColor3uiv)gl121LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl121PColor3usv)gl121LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl121PColor4bv)gl121LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl121PColor4dv)gl121LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl121PColor4fv)gl121LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl121PColor4iv)gl121LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl121PColor4sv)gl121LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl121PColor4ubv)gl121LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl121PColor4uiv)gl121LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl121PColor4usv)gl121LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl121PColorMask)gl121LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl121PColorMaterial)gl121LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl121PCopyPixels)gl121LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl121PCullFace)gl121LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl121PDeleteLists)gl121LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl121PDepthFunc)gl121LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl121PDepthMask)gl121LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl121PDepthRange)gl121LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl121PEnable)gl121LibGetProcAddress("glEnable");
    glc->fnDisable = (gl121PDisable)gl121LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl121PDrawBuffer)gl121LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl121PDrawPixels)gl121LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl121PEdgeFlag)gl121LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl121PEdgeFlagv)gl121LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl121PEdgeFlagPointer)gl121LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl121PEvalCoord1d)gl121LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl121PEvalCoord1f)gl121LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl121PEvalCoord2d)gl121LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl121PEvalCoord2f)gl121LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl121PEvalCoord1dv)gl121LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl121PEvalCoord1fv)gl121LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl121PEvalCoord2dv)gl121LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl121PEvalCoord2fv)gl121LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl121PEvalMesh1)gl121LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl121PEvalMesh2)gl121LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl121PEvalPoint1)gl121LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl121PEvalPoint2)gl121LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl121PFeedbackBuffer)gl121LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl121PFinish)gl121LibGetProcAddress("glFinish");
    glc->fnFlush = (gl121PFlush)gl121LibGetProcAddress("glFlush");
    glc->fnFogf = (gl121PFogf)gl121LibGetProcAddress("glFogf");
    glc->fnFogi = (gl121PFogi)gl121LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl121PFogfv)gl121LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl121PFogiv)gl121LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl121PFrontFace)gl121LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl121PFrustum)gl121LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl121PGenLists)gl121LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl121PGetBooleanv)gl121LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl121PGetDoublev)gl121LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl121PGetFloatv)gl121LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl121PGetIntegerv)gl121LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl121PGetClipPlane)gl121LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl121PGetError)gl121LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl121PGetLightfv)gl121LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl121PGetLightiv)gl121LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl121PGetMapdv)gl121LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl121PGetMapfv)gl121LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl121PGetMapiv)gl121LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl121PGetMaterialfv)gl121LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl121PGetMaterialiv)gl121LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl121PGetPixelMapfv)gl121LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl121PGetPixelMapuiv)gl121LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl121PGetPixelMapusv)gl121LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl121PGetPolygonStipple)gl121LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl121PGetString)gl121LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl121PGetTexEnvfv)gl121LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl121PGetTexEnviv)gl121LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl121PGetTexGendv)gl121LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl121PGetTexGenfv)gl121LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl121PGetTexGeniv)gl121LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl121PGetTexImage)gl121LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl121PGetTexLevelParameterfv)gl121LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl121PGetTexLevelParameteriv)gl121LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl121PGetTexParameterfv)gl121LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl121PGetTexParameteriv)gl121LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl121PHint)gl121LibGetProcAddress("glHint");
    glc->fnIndexd = (gl121PIndexd)gl121LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl121PIndexf)gl121LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl121PIndexi)gl121LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl121PIndexs)gl121LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl121PIndexdv)gl121LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl121PIndexfv)gl121LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl121PIndexiv)gl121LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl121PIndexsv)gl121LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl121PIndexMask)gl121LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl121PIndexPointer)gl121LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl121PInitNames)gl121LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl121PIsEnabled)gl121LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl121PIsList)gl121LibGetProcAddress("glIsList");
    glc->fnLightf = (gl121PLightf)gl121LibGetProcAddress("glLightf");
    glc->fnLighti = (gl121PLighti)gl121LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl121PLightfv)gl121LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl121PLightiv)gl121LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl121PLightModelf)gl121LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl121PLightModeli)gl121LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl121PLightModelfv)gl121LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl121PLightModeliv)gl121LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl121PLineStipple)gl121LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl121PLineWidth)gl121LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl121PListBase)gl121LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl121PLoadIdentity)gl121LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl121PLoadMatrixd)gl121LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl121PLoadMatrixf)gl121LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl121PLoadName)gl121LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl121PLogicOp)gl121LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl121PMap1d)gl121LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl121PMap1f)gl121LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl121PMap2d)gl121LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl121PMap2f)gl121LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl121PMapGrid1d)gl121LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl121PMapGrid1f)gl121LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl121PMapGrid2d)gl121LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl121PMapGrid2f)gl121LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl121PMaterialf)gl121LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl121PMateriali)gl121LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl121PMaterialfv)gl121LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl121PMaterialiv)gl121LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl121PMatrixMode)gl121LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl121PMultMatrixd)gl121LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl121PMultMatrixf)gl121LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl121PNewList)gl121LibGetProcAddress("glNewList");
    glc->fnEndList = (gl121PEndList)gl121LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl121PNormal3b)gl121LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl121PNormal3d)gl121LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl121PNormal3f)gl121LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl121PNormal3i)gl121LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl121PNormal3s)gl121LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl121PNormal3bv)gl121LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl121PNormal3dv)gl121LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl121PNormal3fv)gl121LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl121PNormal3iv)gl121LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl121PNormal3sv)gl121LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl121POrtho)gl121LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl121PPassThrough)gl121LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl121PPixelMapfv)gl121LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl121PPixelMapuiv)gl121LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl121PPixelMapusv)gl121LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl121PPixelStoref)gl121LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl121PPixelStorei)gl121LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl121PPixelTransferf)gl121LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl121PPixelTransferi)gl121LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl121PPixelZoom)gl121LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl121PPointSize)gl121LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl121PPolygonMode)gl121LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl121PPolygonStipple)gl121LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl121PPushAttrib)gl121LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl121PPopAttrib)gl121LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl121PPushMatrix)gl121LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl121PPopMatrix)gl121LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl121PPushName)gl121LibGetProcAddress("glPushName");
    glc->fnPopName = (gl121PPopName)gl121LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl121PRasterPos2d)gl121LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl121PRasterPos2f)gl121LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl121PRasterPos2i)gl121LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl121PRasterPos2s)gl121LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl121PRasterPos3d)gl121LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl121PRasterPos3f)gl121LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl121PRasterPos3i)gl121LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl121PRasterPos3s)gl121LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl121PRasterPos4d)gl121LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl121PRasterPos4f)gl121LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl121PRasterPos4i)gl121LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl121PRasterPos4s)gl121LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl121PRasterPos2dv)gl121LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl121PRasterPos2fv)gl121LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl121PRasterPos2iv)gl121LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl121PRasterPos2sv)gl121LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl121PRasterPos3dv)gl121LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl121PRasterPos3fv)gl121LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl121PRasterPos3iv)gl121LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl121PRasterPos3sv)gl121LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl121PRasterPos4dv)gl121LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl121PRasterPos4fv)gl121LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl121PRasterPos4iv)gl121LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl121PRasterPos4sv)gl121LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl121PReadBuffer)gl121LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl121PReadPixels)gl121LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl121PRectd)gl121LibGetProcAddress("glRectd");
    glc->fnRectf = (gl121PRectf)gl121LibGetProcAddress("glRectf");
    glc->fnRecti = (gl121PRecti)gl121LibGetProcAddress("glRecti");
    glc->fnRects = (gl121PRects)gl121LibGetProcAddress("glRects");
    glc->fnRectdv = (gl121PRectdv)gl121LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl121PRectfv)gl121LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl121PRectiv)gl121LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl121PRectsv)gl121LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl121PRenderMode)gl121LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl121PRotated)gl121LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl121PRotatef)gl121LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl121PScaled)gl121LibGetProcAddress("glScaled");
    glc->fnScalef = (gl121PScalef)gl121LibGetProcAddress("glScalef");
    glc->fnScissor = (gl121PScissor)gl121LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl121PSelectBuffer)gl121LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl121PShadeModel)gl121LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl121PStencilFunc)gl121LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl121PStencilMask)gl121LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl121PStencilOp)gl121LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl121PTexCoord1d)gl121LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl121PTexCoord1f)gl121LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl121PTexCoord1i)gl121LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl121PTexCoord1s)gl121LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl121PTexCoord2d)gl121LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl121PTexCoord2f)gl121LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl121PTexCoord2i)gl121LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl121PTexCoord2s)gl121LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl121PTexCoord3d)gl121LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl121PTexCoord3f)gl121LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl121PTexCoord3i)gl121LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl121PTexCoord3s)gl121LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl121PTexCoord4d)gl121LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl121PTexCoord4f)gl121LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl121PTexCoord4i)gl121LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl121PTexCoord4s)gl121LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl121PTexCoord1dv)gl121LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl121PTexCoord1fv)gl121LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl121PTexCoord1iv)gl121LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl121PTexCoord1sv)gl121LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl121PTexCoord2dv)gl121LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl121PTexCoord2fv)gl121LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl121PTexCoord2iv)gl121LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl121PTexCoord2sv)gl121LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl121PTexCoord3dv)gl121LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl121PTexCoord3fv)gl121LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl121PTexCoord3iv)gl121LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl121PTexCoord3sv)gl121LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl121PTexCoord4dv)gl121LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl121PTexCoord4fv)gl121LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl121PTexCoord4iv)gl121LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl121PTexCoord4sv)gl121LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl121PTexEnvf)gl121LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl121PTexEnvi)gl121LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl121PTexEnvfv)gl121LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl121PTexEnviv)gl121LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl121PTexGend)gl121LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl121PTexGenf)gl121LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl121PTexGeni)gl121LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl121PTexGendv)gl121LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl121PTexGenfv)gl121LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl121PTexGeniv)gl121LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl121PTexImage1D)gl121LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl121PTexImage2D)gl121LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl121PTexParameterf)gl121LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl121PTexParameteri)gl121LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl121PTexParameterfv)gl121LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl121PTexParameteriv)gl121LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl121PTranslated)gl121LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl121PTranslatef)gl121LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl121PVertex2s)gl121LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl121PVertex2i)gl121LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl121PVertex2f)gl121LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl121PVertex2d)gl121LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl121PVertex3s)gl121LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl121PVertex3i)gl121LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl121PVertex3f)gl121LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl121PVertex3d)gl121LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl121PVertex4s)gl121LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl121PVertex4i)gl121LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl121PVertex4f)gl121LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl121PVertex4d)gl121LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl121PViewport)gl121LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl121PGetConvolutionParameterfv)gl121LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl121PGetConvolutionParameteriv)gl121LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl121PAreTexturesResident)gl121LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl121PArrayElement)gl121LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl121PDrawArrays)gl121LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl121PDrawElements)gl121LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl121PGetPointerv)gl121LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl121PPolygonOffset)gl121LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl121PCopyTexImage1D)gl121LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl121PCopyTexImage2D)gl121LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl121PCopyTexSubImage1D)gl121LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl121PCopyTexSubImage2D)gl121LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl121PBindTexture)gl121LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl121PDeleteTextures)gl121LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl121PGenTextures)gl121LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl121PIsTexture)gl121LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl121PColorPointer)gl121LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl121PEnableClientState)gl121LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl121PDisableClientState)gl121LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl121PIndexub)gl121LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl121PIndexubv)gl121LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl121PInterleavedArrays)gl121LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl121PNormalPointer)gl121LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl121PPushClientAttrib)gl121LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl121PPrioritizeTextures)gl121LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl121PPopClientAttrib)gl121LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl121PTexCoordPointer)gl121LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl121PTexSubImage1D)gl121LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl121PTexSubImage2D)gl121LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl121PVertexPointer)gl121LibGetProcAddress("glVertexPointer");
    glc->fnColorTable = (gl121PColorTable)gl121GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl121PColorTableParameterfv)gl121GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl121PColorTableParameteriv)gl121GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl121PColorSubTable)gl121GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl121PConvolutionFilter1D)gl121GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl121PConvolutionFilter2D)gl121GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl121PConvolutionParameterf)gl121GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl121PConvolutionParameteri)gl121GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl121PCopyColorTable)gl121GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl121PCopyColorSubTable)gl121GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl121PCopyConvolutionFilter1D)gl121GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl121PCopyConvolutionFilter2D)gl121GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl121PGetColorTable)gl121GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl121PGetColorTableParameterfv)gl121GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl121PGetColorTableParameteriv)gl121GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl121PGetConvolutionFilter)gl121GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl121PGetHistogram)gl121GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl121PGetHistogramParameterfv)gl121GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl121PGetHistogramParameteriv)gl121GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl121PGetSeparableFilter)gl121GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl121PHistogram)gl121GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl121PMinmax)gl121GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl121PMultiTexCoord1s)gl121GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl121PMultiTexCoord1i)gl121GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl121PMultiTexCoord1f)gl121GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl121PMultiTexCoord1d)gl121GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl121PMultiTexCoord2s)gl121GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl121PMultiTexCoord2i)gl121GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl121PMultiTexCoord2f)gl121GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl121PMultiTexCoord2d)gl121GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl121PMultiTexCoord3s)gl121GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl121PMultiTexCoord3i)gl121GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl121PMultiTexCoord3f)gl121GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl121PMultiTexCoord3d)gl121GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl121PMultiTexCoord4s)gl121GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl121PMultiTexCoord4i)gl121GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl121PMultiTexCoord4f)gl121GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl121PMultiTexCoord4d)gl121GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl121PMultiTexCoord1sv)gl121GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl121PMultiTexCoord1iv)gl121GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl121PMultiTexCoord1fv)gl121GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl121PMultiTexCoord1dv)gl121GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl121PMultiTexCoord2sv)gl121GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl121PMultiTexCoord2iv)gl121GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl121PMultiTexCoord2fv)gl121GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl121PMultiTexCoord2dv)gl121GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl121PMultiTexCoord3sv)gl121GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl121PMultiTexCoord3iv)gl121GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl121PMultiTexCoord3fv)gl121GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl121PMultiTexCoord3dv)gl121GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl121PMultiTexCoord4sv)gl121GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl121PMultiTexCoord4iv)gl121GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl121PMultiTexCoord4fv)gl121GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl121PMultiTexCoord4dv)gl121GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl121PResetHistogram)gl121GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl121PResetMinmax)gl121GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl121PSeparableFilter2D)gl121GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl121PBlendColor)gl121GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl121PBlendEquation)gl121GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl121PCopyTexSubImage3D)gl121GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl121PDrawRangeElements)gl121GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl121PTexImage3D)gl121GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl121PTexSubImage3D)gl121GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl121PActiveTexture)gl121GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl121PClientActiveTexture)gl121GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl121PCompressedTexImage1D)gl121GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl121PCompressedTexImage2D)gl121GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl121PCompressedTexImage3D)gl121GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl121PCompressedTexSubImage1D)gl121GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl121PCompressedTexSubImage2D)gl121GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl121PCompressedTexSubImage3D)gl121GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl121PGetCompressedTexImage)gl121GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl121PLoadTransposeMatrixd)gl121GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl121PLoadTransposeMatrixf)gl121GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl121PMultTransposeMatrixd)gl121GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl121PMultTransposeMatrixf)gl121GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl121PSampleCoverage)gl121GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl121PBlendFuncSeparate)gl121GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl121PFogCoordPointer)gl121GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl121PFogCoordd)gl121GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl121PFogCoordf)gl121GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl121PFogCoorddv)gl121GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl121PFogCoordfv)gl121GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl121PMultiDrawArrays)gl121GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl121PMultiDrawElements)gl121GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl121PPointParameterf)gl121GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl121PPointParameteri)gl121GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl121PSecondaryColor3b)gl121GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl121PSecondaryColor3s)gl121GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl121PSecondaryColor3i)gl121GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl121PSecondaryColor3f)gl121GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl121PSecondaryColor3d)gl121GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl121PSecondaryColor3ub)gl121GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl121PSecondaryColor3us)gl121GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl121PSecondaryColor3ui)gl121GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl121PSecondaryColor3bv)gl121GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl121PSecondaryColor3sv)gl121GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl121PSecondaryColor3iv)gl121GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl121PSecondaryColor3fv)gl121GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl121PSecondaryColor3dv)gl121GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl121PSecondaryColor3ubv)gl121GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl121PSecondaryColor3usv)gl121GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl121PSecondaryColor3uiv)gl121GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl121PSecondaryColorPointer)gl121GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl121PWindowPos2s)gl121GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl121PWindowPos2i)gl121GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl121PWindowPos2f)gl121GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl121PWindowPos2d)gl121GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl121PWindowPos3s)gl121GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl121PWindowPos3i)gl121GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl121PWindowPos3f)gl121GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl121PWindowPos3d)gl121GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl121PWindowPos2sv)gl121GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl121PWindowPos2iv)gl121GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl121PWindowPos2fv)gl121GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl121PWindowPos2dv)gl121GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl121PWindowPos3sv)gl121GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl121PWindowPos3iv)gl121GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl121PWindowPos3fv)gl121GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl121PWindowPos3dv)gl121GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl121PBeginQuery)gl121GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl121PBindBuffer)gl121GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl121PBufferData)gl121GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl121PBufferSubData)gl121GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl121PDeleteBuffers)gl121GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl121PDeleteQueries)gl121GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl121PGenBuffers)gl121GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl121PGenQueries)gl121GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl121PGetBufferParameteriv)gl121GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl121PGetBufferPointerv)gl121GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl121PGetBufferSubData)gl121GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl121PGetQueryObjectiv)gl121GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl121PGetQueryObjectuiv)gl121GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl121PGetQueryiv)gl121GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl121PIsBuffer)gl121GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl121PIsQuery)gl121GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl121PMapBuffer)gl121GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl121PUnmapBuffer)gl121GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl121PAttachShader)gl121GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl121PBindAttribLocation)gl121GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl121PBlendEquationSeperate)gl121GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl121PCompileShader)gl121GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl121PCreateProgram)gl121GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl121PCreateShader)gl121GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl121PDeleteProgram)gl121GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl121PDeleteShader)gl121GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl121PDetachShader)gl121GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl121PEnableVertexAttribArray)gl121GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl121PDisableVertexAttribArray)gl121GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnDrawBuffers = (gl121PDrawBuffers)gl121LibGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl121PGetActiveAttrib)gl121LibGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl121PGetActiveUniform)gl121LibGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl121PGetAttachedShaders)gl121LibGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl121PGetAttribLocation)gl121LibGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl121PGetProgramiv)gl121LibGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl121PGetProgramInfoLog)gl121LibGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl121PGetShaderiv)gl121LibGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl121PGetShaderInfoLog)gl121LibGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl121PGetShaderSource)gl121LibGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl121PGetUniformfv)gl121LibGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl121PGetUniformiv)gl121LibGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl121PGetUniformLocation)gl121LibGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl121PGetVertexAttribdv)gl121LibGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl121PGetVertexAttribfv)gl121LibGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl121PGetVertexAttribiv)gl121LibGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl121PGetVertexAttribPointerv)gl121LibGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl121PIsProgram)gl121LibGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl121PIsShader)gl121LibGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl121PLinkProgram)gl121LibGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl121PShaderSource)gl121LibGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl121PStencilFuncSeparate)gl121LibGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl121PStencilMaskSeparate)gl121LibGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl121PStencilOpSeparate)gl121LibGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl121PUniform1f)gl121LibGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl121PUniform2f)gl121LibGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl121PUniform3f)gl121LibGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl121PUniform4f)gl121LibGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl121PUniform1i)gl121LibGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl121PUniform2i)gl121LibGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl121PUniform3i)gl121LibGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl121PUniform4i)gl121LibGetProcAddress("glUniform4i");
    glc->fnUniform1fv = (gl121PUniform1fv)gl121LibGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl121PUniform2fv)gl121LibGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl121PUniform3fv)gl121LibGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl121PUniform4fv)gl121LibGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl121PUniform1iv)gl121LibGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl121PUniform2iv)gl121LibGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl121PUniform3iv)gl121LibGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl121PUniform4iv)gl121LibGetProcAddress("glUniform4iv");
    glc->fnUseProgram = (gl121PUseProgram)gl121LibGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl121PValidateProgram)gl121LibGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl121PVertexAttribPointer)gl121LibGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl121PVertexAttrib1f)gl121LibGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl121PVertexAttrib1s)gl121LibGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl121PVertexAttrib1d)gl121LibGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl121PVertexAttrib2f)gl121LibGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl121PVertexAttrib2s)gl121LibGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl121PVertexAttrib2d)gl121LibGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl121PVertexAttrib3f)gl121LibGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl121PVertexAttrib3s)gl121LibGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl121PVertexAttrib3d)gl121LibGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl121PVertexAttrib4f)gl121LibGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl121PVertexAttrib4s)gl121LibGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl121PVertexAttrib4d)gl121LibGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl121PVertexAttrib4Nuv)gl121LibGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl121PVertexAttrib1fv)gl121LibGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl121PVertexAttrib1sv)gl121LibGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl121PVertexAttrib1dv)gl121LibGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl121PVertexAttrib2fv)gl121LibGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl121PVertexAttrib2sv)gl121LibGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl121PVertexAttrib2dv)gl121LibGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl121PVertexAttrib3fv)gl121LibGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl121PVertexAttrib3sv)gl121LibGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl121PVertexAttrib3dv)gl121LibGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl121PVertexAttrib4fv)gl121LibGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl121PVertexAttrib4sv)gl121LibGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl121PVertexAttrib4dv)gl121LibGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl121PVertexAttrib4iv)gl121LibGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl121PVertexAttrib4bv)gl121LibGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl121PVertexAttrib4ubv)gl121LibGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl121PVertexAttrib4usv)gl121LibGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl121PVertexAttrib4uiv)gl121LibGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl121PVertexAttrib4Nbv)gl121LibGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl121PVertexAttrib4Nsv)gl121LibGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl121PVertexAttrib4Niv)gl121LibGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl121PVertexAttrib4Nubv)gl121LibGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl121PVertexAttrib4Nusv)gl121LibGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl121PVertexAttrib4Nuiv)gl121LibGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl121PUniformMatrix2fv)gl121LibGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl121PUniformMatrix3fv)gl121LibGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl121PUniformMatrix4fv)gl121LibGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl121PUniformMatrix2x3fv)gl121LibGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl121PUniformMatrix3x2fv)gl121LibGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl121PUniformMatrix2x4fv)gl121LibGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl121PUniformMatrix4x2fv)gl121LibGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl121PUniformMatrix3x4fv)gl121LibGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl121PUniformMatrix4x3fv)gl121LibGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

