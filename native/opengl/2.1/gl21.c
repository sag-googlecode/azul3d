
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl21.h"

#ifdef _WIN32
HMODULE gl21OpenGL32;

void* gl21LibGetProcAddress(char* name) {
	if(gl21OpenGL32 == NULL) {
		gl21OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl21OpenGL32, TEXT(name));
}

void* gl21GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl21Accum(gl21Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl21AlphaFunc(gl21Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl21Begin(gl21Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl21End(gl21Context* glc) {
    return glc->fnEnd();
}

void gl21Bitmap(gl21Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl21BlendFunc(gl21Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl21CallList(gl21Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl21CallLists(gl21Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl21Clear(gl21Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl21ClearAccum(gl21Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl21ClearColor(gl21Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl21ClearDepth(gl21Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl21ClearDepthf(gl21Context* glc, GLclampf depth) {
    return glc->fnClearDepthf(depth);
}

void gl21ClearIndex(gl21Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl21ClearStencil(gl21Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl21ClipPlane(gl21Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl21Color3b(gl21Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl21Color3d(gl21Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl21Color3f(gl21Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl21Color3i(gl21Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl21Color3s(gl21Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl21Color3ub(gl21Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl21Color3ui(gl21Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl21Color3us(gl21Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl21Color4b(gl21Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl21Color4d(gl21Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl21Color4f(gl21Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl21Color4i(gl21Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl21Color4s(gl21Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl21Color4ub(gl21Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl21Color4ui(gl21Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl21Color4us(gl21Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl21Color3bv(gl21Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl21Color3dv(gl21Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl21Color3fv(gl21Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl21Color3iv(gl21Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl21Color3sv(gl21Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl21Color3ubv(gl21Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl21Color3uiv(gl21Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl21Color3usv(gl21Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl21Color4bv(gl21Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl21Color4dv(gl21Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl21Color4fv(gl21Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl21Color4iv(gl21Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl21Color4sv(gl21Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl21Color4ubv(gl21Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl21Color4uiv(gl21Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl21Color4usv(gl21Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl21ColorMask(gl21Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl21ColorMaterial(gl21Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl21CopyPixels(gl21Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl21CullFace(gl21Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl21DeleteLists(gl21Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl21DepthFunc(gl21Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl21DepthMask(gl21Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl21DepthRange(gl21Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl21Enable(gl21Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl21Disable(gl21Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl21DrawBuffer(gl21Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl21DrawPixels(gl21Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl21EdgeFlag(gl21Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl21EdgeFlagv(gl21Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl21EdgeFlagPointer(gl21Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl21EvalCoord1d(gl21Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl21EvalCoord1f(gl21Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl21EvalCoord2d(gl21Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl21EvalCoord2f(gl21Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl21EvalCoord1dv(gl21Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl21EvalCoord1fv(gl21Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl21EvalCoord2dv(gl21Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl21EvalCoord2fv(gl21Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl21EvalMesh1(gl21Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl21EvalMesh2(gl21Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl21EvalPoint1(gl21Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl21EvalPoint2(gl21Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl21FeedbackBuffer(gl21Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl21Finish(gl21Context* glc) {
    return glc->fnFinish();
}

void gl21Flush(gl21Context* glc) {
    return glc->fnFlush();
}

void gl21Fogf(gl21Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl21Fogi(gl21Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl21Fogfv(gl21Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl21Fogiv(gl21Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl21FrontFace(gl21Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl21Frustum(gl21Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl21GenLists(gl21Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl21GetBooleanv(gl21Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl21GetDoublev(gl21Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl21GetFloatv(gl21Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl21GetIntegerv(gl21Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl21GetClipPlane(gl21Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl21GetError(gl21Context* glc) {
    return glc->fnGetError();
}

void gl21GetLightfv(gl21Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl21GetLightiv(gl21Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl21GetMapdv(gl21Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl21GetMapfv(gl21Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl21GetMapiv(gl21Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl21GetMaterialfv(gl21Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl21GetMaterialiv(gl21Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl21GetPixelMapfv(gl21Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl21GetPixelMapuiv(gl21Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl21GetPixelMapusv(gl21Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl21GetPolygonStipple(gl21Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl21GetString(gl21Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl21GetTexEnvfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl21GetTexEnviv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl21GetTexGendv(gl21Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl21GetTexGenfv(gl21Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl21GetTexGeniv(gl21Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl21GetTexImage(gl21Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl21GetTexLevelParameterfv(gl21Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl21GetTexLevelParameteriv(gl21Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl21GetTexParameterfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl21GetTexParameteriv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl21Hint(gl21Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl21Indexd(gl21Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl21Indexf(gl21Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl21Indexi(gl21Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl21Indexs(gl21Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl21Indexdv(gl21Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl21Indexfv(gl21Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl21Indexiv(gl21Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl21Indexsv(gl21Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl21IndexMask(gl21Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl21IndexPointer(gl21Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl21InitNames(gl21Context* glc) {
    return glc->fnInitNames();
}

void gl21IsEnabled(gl21Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl21IsList(gl21Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl21Lightf(gl21Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl21Lighti(gl21Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl21Lightfv(gl21Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl21Lightiv(gl21Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl21LightModelf(gl21Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl21LightModeli(gl21Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl21LightModelfv(gl21Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl21LightModeliv(gl21Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl21LineStipple(gl21Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl21LineWidth(gl21Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl21ListBase(gl21Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl21LoadIdentity(gl21Context* glc) {
    return glc->fnLoadIdentity();
}

void gl21LoadMatrixd(gl21Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl21LoadMatrixf(gl21Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl21LoadName(gl21Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl21LogicOp(gl21Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl21Map1d(gl21Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl21Map1f(gl21Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl21Map2d(gl21Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl21Map2f(gl21Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl21MapGrid1d(gl21Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl21MapGrid1f(gl21Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl21MapGrid2d(gl21Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl21MapGrid2f(gl21Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl21Materialf(gl21Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl21Materiali(gl21Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl21Materialfv(gl21Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl21Materialiv(gl21Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl21MatrixMode(gl21Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl21MultMatrixd(gl21Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl21MultMatrixf(gl21Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl21NewList(gl21Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl21EndList(gl21Context* glc) {
    return glc->fnEndList();
}

void gl21Normal3b(gl21Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl21Normal3d(gl21Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl21Normal3f(gl21Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl21Normal3i(gl21Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl21Normal3s(gl21Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl21Normal3bv(gl21Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl21Normal3dv(gl21Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl21Normal3fv(gl21Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl21Normal3iv(gl21Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl21Normal3sv(gl21Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl21Ortho(gl21Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl21PassThrough(gl21Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl21PixelMapfv(gl21Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl21PixelMapuiv(gl21Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl21PixelMapusv(gl21Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl21PixelStoref(gl21Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl21PixelStorei(gl21Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl21PixelTransferf(gl21Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl21PixelTransferi(gl21Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl21PixelZoom(gl21Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl21PointSize(gl21Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl21PolygonMode(gl21Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl21PolygonStipple(gl21Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl21PushAttrib(gl21Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl21PopAttrib(gl21Context* glc) {
    return glc->fnPopAttrib();
}

void gl21PushMatrix(gl21Context* glc) {
    return glc->fnPushMatrix();
}

void gl21PopMatrix(gl21Context* glc) {
    return glc->fnPopMatrix();
}

void gl21PushName(gl21Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl21PopName(gl21Context* glc) {
    return glc->fnPopName();
}

void gl21RasterPos2d(gl21Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl21RasterPos2f(gl21Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl21RasterPos2i(gl21Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl21RasterPos2s(gl21Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl21RasterPos3d(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl21RasterPos3f(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl21RasterPos3i(gl21Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl21RasterPos3s(gl21Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl21RasterPos4d(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl21RasterPos4f(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl21RasterPos4i(gl21Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl21RasterPos4s(gl21Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl21RasterPos2dv(gl21Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl21RasterPos2fv(gl21Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl21RasterPos2iv(gl21Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl21RasterPos2sv(gl21Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl21RasterPos3dv(gl21Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl21RasterPos3fv(gl21Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl21RasterPos3iv(gl21Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl21RasterPos3sv(gl21Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl21RasterPos4dv(gl21Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl21RasterPos4fv(gl21Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl21RasterPos4iv(gl21Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl21RasterPos4sv(gl21Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl21ReadBuffer(gl21Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl21ReadPixels(gl21Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl21Rectd(gl21Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl21Rectf(gl21Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl21Recti(gl21Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl21Rects(gl21Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl21Rectdv(gl21Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl21Rectfv(gl21Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl21Rectiv(gl21Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl21Rectsv(gl21Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl21RenderMode(gl21Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl21Rotated(gl21Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl21Rotatef(gl21Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl21Scaled(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl21Scalef(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl21Scissor(gl21Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl21SelectBuffer(gl21Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl21ShadeModel(gl21Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl21StencilFunc(gl21Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl21StencilMask(gl21Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl21StencilOp(gl21Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl21TexCoord1d(gl21Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl21TexCoord1f(gl21Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl21TexCoord1i(gl21Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl21TexCoord1s(gl21Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl21TexCoord2d(gl21Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl21TexCoord2f(gl21Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl21TexCoord2i(gl21Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl21TexCoord2s(gl21Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl21TexCoord3d(gl21Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl21TexCoord3f(gl21Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl21TexCoord3i(gl21Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl21TexCoord3s(gl21Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl21TexCoord4d(gl21Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl21TexCoord4f(gl21Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl21TexCoord4i(gl21Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl21TexCoord4s(gl21Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl21TexCoord1dv(gl21Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl21TexCoord1fv(gl21Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl21TexCoord1iv(gl21Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl21TexCoord1sv(gl21Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl21TexCoord2dv(gl21Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl21TexCoord2fv(gl21Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl21TexCoord2iv(gl21Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl21TexCoord2sv(gl21Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl21TexCoord3dv(gl21Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl21TexCoord3fv(gl21Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl21TexCoord3iv(gl21Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl21TexCoord3sv(gl21Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl21TexCoord4dv(gl21Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl21TexCoord4fv(gl21Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl21TexCoord4iv(gl21Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl21TexCoord4sv(gl21Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl21TexEnvf(gl21Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl21TexEnvi(gl21Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl21TexEnvfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl21TexEnviv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl21TexGend(gl21Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl21TexGenf(gl21Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl21TexGeni(gl21Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl21TexGendv(gl21Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl21TexGenfv(gl21Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl21TexGeniv(gl21Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl21TexImage1D(gl21Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl21TexImage2D(gl21Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl21TexParameterf(gl21Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl21TexParameteri(gl21Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl21TexParameterfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl21TexParameteriv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl21Translated(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl21Translatef(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl21Vertex2s(gl21Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl21Vertex2i(gl21Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl21Vertex2f(gl21Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl21Vertex2d(gl21Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl21Vertex3s(gl21Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl21Vertex3i(gl21Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl21Vertex3f(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl21Vertex3d(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl21Vertex4s(gl21Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl21Vertex4i(gl21Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl21Vertex4f(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl21Vertex4d(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl21Viewport(gl21Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl21GetConvolutionParameterfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl21GetConvolutionParameteriv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl21AreTexturesResident(gl21Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl21ArrayElement(gl21Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl21DrawArrays(gl21Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl21DrawElements(gl21Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl21GetPointerv(gl21Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl21PolygonOffset(gl21Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl21CopyTexImage1D(gl21Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl21CopyTexImage2D(gl21Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl21CopyTexSubImage1D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl21CopyTexSubImage2D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl21BindTexture(gl21Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl21DeleteTextures(gl21Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl21GenTextures(gl21Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl21IsTexture(gl21Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl21ColorPointer(gl21Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl21EnableClientState(gl21Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl21DisableClientState(gl21Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl21Indexub(gl21Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl21Indexubv(gl21Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl21InterleavedArrays(gl21Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl21NormalPointer(gl21Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl21PushClientAttrib(gl21Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl21PrioritizeTextures(gl21Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl21PopClientAttrib(gl21Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl21TexCoordPointer(gl21Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl21TexSubImage1D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl21TexSubImage2D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl21VertexPointer(gl21Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl21DrawBuffers(gl21Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl21GetActiveAttrib(gl21Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl21GetActiveUniform(gl21Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl21GetAttachedShaders(gl21Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl21GetAttribLocation(gl21Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl21GetProgramiv(gl21Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl21GetProgramInfoLog(gl21Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl21GetShaderiv(gl21Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl21GetShaderInfoLog(gl21Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl21GetShaderSource(gl21Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl21GetUniformfv(gl21Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl21GetUniformiv(gl21Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl21GetUniformLocation(gl21Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl21GetVertexAttribdv(gl21Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl21GetVertexAttribfv(gl21Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl21GetVertexAttribiv(gl21Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl21GetVertexAttribPointerv(gl21Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl21IsProgram(gl21Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl21IsShader(gl21Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl21LinkProgram(gl21Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl21ShaderSource(gl21Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl21StencilFuncSeparate(gl21Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl21StencilMaskSeparate(gl21Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl21StencilOpSeparate(gl21Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl21Uniform1f(gl21Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl21Uniform2f(gl21Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl21Uniform3f(gl21Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl21Uniform4f(gl21Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl21Uniform1i(gl21Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl21Uniform2i(gl21Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl21Uniform3i(gl21Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl21Uniform4i(gl21Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl21Uniform1fv(gl21Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl21Uniform2fv(gl21Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl21Uniform3fv(gl21Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl21Uniform4fv(gl21Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl21Uniform1iv(gl21Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl21Uniform2iv(gl21Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl21Uniform3iv(gl21Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl21Uniform4iv(gl21Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl21UseProgram(gl21Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl21ValidateProgram(gl21Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl21VertexAttribPointer(gl21Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl21VertexAttrib1f(gl21Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl21VertexAttrib1s(gl21Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl21VertexAttrib1d(gl21Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl21VertexAttrib2f(gl21Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl21VertexAttrib2s(gl21Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl21VertexAttrib2d(gl21Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl21VertexAttrib3f(gl21Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl21VertexAttrib3s(gl21Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl21VertexAttrib3d(gl21Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl21VertexAttrib4f(gl21Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl21VertexAttrib4s(gl21Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl21VertexAttrib4d(gl21Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl21VertexAttrib4Nuv(gl21Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl21VertexAttrib1fv(gl21Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl21VertexAttrib1sv(gl21Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl21VertexAttrib1dv(gl21Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl21VertexAttrib2fv(gl21Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl21VertexAttrib2sv(gl21Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl21VertexAttrib2dv(gl21Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl21VertexAttrib3fv(gl21Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl21VertexAttrib3sv(gl21Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl21VertexAttrib3dv(gl21Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl21VertexAttrib4fv(gl21Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl21VertexAttrib4sv(gl21Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl21VertexAttrib4dv(gl21Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl21VertexAttrib4iv(gl21Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl21VertexAttrib4bv(gl21Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl21VertexAttrib4ubv(gl21Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl21VertexAttrib4usv(gl21Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl21VertexAttrib4uiv(gl21Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl21VertexAttrib4Nbv(gl21Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl21VertexAttrib4Nsv(gl21Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl21VertexAttrib4Niv(gl21Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl21VertexAttrib4Nubv(gl21Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl21VertexAttrib4Nusv(gl21Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl21VertexAttrib4Nuiv(gl21Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl21UniformMatrix2fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl21UniformMatrix3fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl21UniformMatrix4fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl21UniformMatrix2x3fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl21UniformMatrix3x2fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl21UniformMatrix2x4fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl21UniformMatrix4x2fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl21UniformMatrix3x4fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl21UniformMatrix4x3fv(gl21Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

void gl21ColorTable(gl21Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl21ColorTableParameterfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl21ColorTableParameteriv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl21ColorSubTable(gl21Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl21ConvolutionFilter1D(gl21Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl21ConvolutionFilter2D(gl21Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl21ConvolutionParameterf(gl21Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl21ConvolutionParameteri(gl21Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl21CopyColorTable(gl21Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl21CopyColorSubTable(gl21Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl21CopyConvolutionFilter1D(gl21Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl21CopyConvolutionFilter2D(gl21Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl21GetColorTable(gl21Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl21GetColorTableParameterfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl21GetColorTableParameteriv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl21GetConvolutionFilter(gl21Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl21GetHistogram(gl21Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl21GetHistogramParameterfv(gl21Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl21GetHistogramParameteriv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl21GetSeparableFilter(gl21Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl21Histogram(gl21Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl21Minmax(gl21Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl21MultiTexCoord1s(gl21Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl21MultiTexCoord1i(gl21Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl21MultiTexCoord1f(gl21Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl21MultiTexCoord1d(gl21Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl21MultiTexCoord2s(gl21Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl21MultiTexCoord2i(gl21Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl21MultiTexCoord2f(gl21Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl21MultiTexCoord2d(gl21Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl21MultiTexCoord3s(gl21Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl21MultiTexCoord3i(gl21Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl21MultiTexCoord3f(gl21Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl21MultiTexCoord3d(gl21Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl21MultiTexCoord4s(gl21Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl21MultiTexCoord4i(gl21Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl21MultiTexCoord4f(gl21Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl21MultiTexCoord4d(gl21Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl21MultiTexCoord1sv(gl21Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl21MultiTexCoord1iv(gl21Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl21MultiTexCoord1fv(gl21Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl21MultiTexCoord1dv(gl21Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl21MultiTexCoord2sv(gl21Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl21MultiTexCoord2iv(gl21Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl21MultiTexCoord2fv(gl21Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl21MultiTexCoord2dv(gl21Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl21MultiTexCoord3sv(gl21Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl21MultiTexCoord3iv(gl21Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl21MultiTexCoord3fv(gl21Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl21MultiTexCoord3dv(gl21Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl21MultiTexCoord4sv(gl21Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl21MultiTexCoord4iv(gl21Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl21MultiTexCoord4fv(gl21Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl21MultiTexCoord4dv(gl21Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl21ResetHistogram(gl21Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl21ResetMinmax(gl21Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl21SeparableFilter2D(gl21Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl21BlendColor(gl21Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl21BlendEquation(gl21Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl21CopyTexSubImage3D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl21DrawRangeElements(gl21Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl21TexImage3D(gl21Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl21TexSubImage3D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl21ActiveTexture(gl21Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl21ClientActiveTexture(gl21Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl21CompressedTexImage1D(gl21Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl21CompressedTexImage2D(gl21Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl21CompressedTexImage3D(gl21Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl21CompressedTexSubImage1D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl21CompressedTexSubImage2D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl21CompressedTexSubImage3D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl21GetCompressedTexImage(gl21Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl21LoadTransposeMatrixd(gl21Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl21LoadTransposeMatrixf(gl21Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl21MultTransposeMatrixd(gl21Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl21MultTransposeMatrixf(gl21Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl21SampleCoverage(gl21Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl21BlendFuncSeparate(gl21Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl21FogCoordPointer(gl21Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl21FogCoordd(gl21Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl21FogCoordf(gl21Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl21FogCoorddv(gl21Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl21FogCoordfv(gl21Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl21MultiDrawArrays(gl21Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl21MultiDrawElements(gl21Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl21PointParameterf(gl21Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl21PointParameteri(gl21Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl21SecondaryColor3b(gl21Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl21SecondaryColor3s(gl21Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl21SecondaryColor3i(gl21Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl21SecondaryColor3f(gl21Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl21SecondaryColor3d(gl21Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl21SecondaryColor3ub(gl21Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl21SecondaryColor3us(gl21Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl21SecondaryColor3ui(gl21Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl21SecondaryColor3bv(gl21Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl21SecondaryColor3sv(gl21Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl21SecondaryColor3iv(gl21Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl21SecondaryColor3fv(gl21Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl21SecondaryColor3dv(gl21Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl21SecondaryColor3ubv(gl21Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl21SecondaryColor3usv(gl21Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl21SecondaryColor3uiv(gl21Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl21SecondaryColorPointer(gl21Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl21WindowPos2s(gl21Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl21WindowPos2i(gl21Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl21WindowPos2f(gl21Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl21WindowPos2d(gl21Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl21WindowPos3s(gl21Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl21WindowPos3i(gl21Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl21WindowPos3f(gl21Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl21WindowPos3d(gl21Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl21WindowPos2sv(gl21Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl21WindowPos2iv(gl21Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl21WindowPos2fv(gl21Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl21WindowPos2dv(gl21Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl21WindowPos3sv(gl21Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl21WindowPos3iv(gl21Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl21WindowPos3fv(gl21Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl21WindowPos3dv(gl21Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl21BeginQuery(gl21Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl21BindBuffer(gl21Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl21BufferData(gl21Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl21BufferSubData(gl21Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl21DeleteBuffers(gl21Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl21DeleteQueries(gl21Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl21GenBuffers(gl21Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl21GenQueries(gl21Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl21GetBufferParameteriv(gl21Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl21GetBufferPointerv(gl21Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl21GetBufferSubData(gl21Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl21GetQueryObjectiv(gl21Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl21GetQueryObjectuiv(gl21Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl21GetQueryiv(gl21Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl21IsBuffer(gl21Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl21IsQuery(gl21Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl21MapBuffer(gl21Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl21UnmapBuffer(gl21Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl21AttachShader(gl21Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl21BindAttribLocation(gl21Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl21BlendEquationSeperate(gl21Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl21CompileShader(gl21Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl21CreateProgram(gl21Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl21CreateShader(gl21Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl21DeleteProgram(gl21Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl21DeleteShader(gl21Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl21DetachShader(gl21Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl21EnableVertexAttribArray(gl21Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl21DisableVertexAttribArray(gl21Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

gl21Context* gl21NewContext() {
    gl21Context* glc = calloc(1, sizeof(gl21Context));

    // Preload all procedures
    glc->fnAccum = (gl21PAccum)gl21LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl21PAlphaFunc)gl21LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl21PBegin)gl21LibGetProcAddress("glBegin");
    glc->fnEnd = (gl21PEnd)gl21LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl21PBitmap)gl21LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl21PBlendFunc)gl21LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl21PCallList)gl21LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl21PCallLists)gl21LibGetProcAddress("glCallLists");
    glc->fnClear = (gl21PClear)gl21LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl21PClearAccum)gl21LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl21PClearColor)gl21LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl21PClearDepth)gl21LibGetProcAddress("glClearDepth");
    glc->fnClearDepthf = (gl21PClearDepthf)gl21LibGetProcAddress("glClearDepthf");
    glc->fnClearIndex = (gl21PClearIndex)gl21LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl21PClearStencil)gl21LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl21PClipPlane)gl21LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl21PColor3b)gl21LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl21PColor3d)gl21LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl21PColor3f)gl21LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl21PColor3i)gl21LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl21PColor3s)gl21LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl21PColor3ub)gl21LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl21PColor3ui)gl21LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl21PColor3us)gl21LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl21PColor4b)gl21LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl21PColor4d)gl21LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl21PColor4f)gl21LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl21PColor4i)gl21LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl21PColor4s)gl21LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl21PColor4ub)gl21LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl21PColor4ui)gl21LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl21PColor4us)gl21LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl21PColor3bv)gl21LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl21PColor3dv)gl21LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl21PColor3fv)gl21LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl21PColor3iv)gl21LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl21PColor3sv)gl21LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl21PColor3ubv)gl21LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl21PColor3uiv)gl21LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl21PColor3usv)gl21LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl21PColor4bv)gl21LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl21PColor4dv)gl21LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl21PColor4fv)gl21LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl21PColor4iv)gl21LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl21PColor4sv)gl21LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl21PColor4ubv)gl21LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl21PColor4uiv)gl21LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl21PColor4usv)gl21LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl21PColorMask)gl21LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl21PColorMaterial)gl21LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl21PCopyPixels)gl21LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl21PCullFace)gl21LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl21PDeleteLists)gl21LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl21PDepthFunc)gl21LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl21PDepthMask)gl21LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl21PDepthRange)gl21LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl21PEnable)gl21LibGetProcAddress("glEnable");
    glc->fnDisable = (gl21PDisable)gl21LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl21PDrawBuffer)gl21LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl21PDrawPixels)gl21LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl21PEdgeFlag)gl21LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl21PEdgeFlagv)gl21LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl21PEdgeFlagPointer)gl21LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl21PEvalCoord1d)gl21LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl21PEvalCoord1f)gl21LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl21PEvalCoord2d)gl21LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl21PEvalCoord2f)gl21LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl21PEvalCoord1dv)gl21LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl21PEvalCoord1fv)gl21LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl21PEvalCoord2dv)gl21LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl21PEvalCoord2fv)gl21LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl21PEvalMesh1)gl21LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl21PEvalMesh2)gl21LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl21PEvalPoint1)gl21LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl21PEvalPoint2)gl21LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl21PFeedbackBuffer)gl21LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl21PFinish)gl21LibGetProcAddress("glFinish");
    glc->fnFlush = (gl21PFlush)gl21LibGetProcAddress("glFlush");
    glc->fnFogf = (gl21PFogf)gl21LibGetProcAddress("glFogf");
    glc->fnFogi = (gl21PFogi)gl21LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl21PFogfv)gl21LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl21PFogiv)gl21LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl21PFrontFace)gl21LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl21PFrustum)gl21LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl21PGenLists)gl21LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl21PGetBooleanv)gl21LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl21PGetDoublev)gl21LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl21PGetFloatv)gl21LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl21PGetIntegerv)gl21LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl21PGetClipPlane)gl21LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl21PGetError)gl21LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl21PGetLightfv)gl21LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl21PGetLightiv)gl21LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl21PGetMapdv)gl21LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl21PGetMapfv)gl21LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl21PGetMapiv)gl21LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl21PGetMaterialfv)gl21LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl21PGetMaterialiv)gl21LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl21PGetPixelMapfv)gl21LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl21PGetPixelMapuiv)gl21LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl21PGetPixelMapusv)gl21LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl21PGetPolygonStipple)gl21LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl21PGetString)gl21LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl21PGetTexEnvfv)gl21LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl21PGetTexEnviv)gl21LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl21PGetTexGendv)gl21LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl21PGetTexGenfv)gl21LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl21PGetTexGeniv)gl21LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl21PGetTexImage)gl21LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl21PGetTexLevelParameterfv)gl21LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl21PGetTexLevelParameteriv)gl21LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl21PGetTexParameterfv)gl21LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl21PGetTexParameteriv)gl21LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl21PHint)gl21LibGetProcAddress("glHint");
    glc->fnIndexd = (gl21PIndexd)gl21LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl21PIndexf)gl21LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl21PIndexi)gl21LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl21PIndexs)gl21LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl21PIndexdv)gl21LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl21PIndexfv)gl21LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl21PIndexiv)gl21LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl21PIndexsv)gl21LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl21PIndexMask)gl21LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl21PIndexPointer)gl21LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl21PInitNames)gl21LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl21PIsEnabled)gl21LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl21PIsList)gl21LibGetProcAddress("glIsList");
    glc->fnLightf = (gl21PLightf)gl21LibGetProcAddress("glLightf");
    glc->fnLighti = (gl21PLighti)gl21LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl21PLightfv)gl21LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl21PLightiv)gl21LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl21PLightModelf)gl21LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl21PLightModeli)gl21LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl21PLightModelfv)gl21LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl21PLightModeliv)gl21LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl21PLineStipple)gl21LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl21PLineWidth)gl21LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl21PListBase)gl21LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl21PLoadIdentity)gl21LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl21PLoadMatrixd)gl21LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl21PLoadMatrixf)gl21LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl21PLoadName)gl21LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl21PLogicOp)gl21LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl21PMap1d)gl21LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl21PMap1f)gl21LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl21PMap2d)gl21LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl21PMap2f)gl21LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl21PMapGrid1d)gl21LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl21PMapGrid1f)gl21LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl21PMapGrid2d)gl21LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl21PMapGrid2f)gl21LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl21PMaterialf)gl21LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl21PMateriali)gl21LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl21PMaterialfv)gl21LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl21PMaterialiv)gl21LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl21PMatrixMode)gl21LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl21PMultMatrixd)gl21LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl21PMultMatrixf)gl21LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl21PNewList)gl21LibGetProcAddress("glNewList");
    glc->fnEndList = (gl21PEndList)gl21LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl21PNormal3b)gl21LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl21PNormal3d)gl21LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl21PNormal3f)gl21LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl21PNormal3i)gl21LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl21PNormal3s)gl21LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl21PNormal3bv)gl21LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl21PNormal3dv)gl21LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl21PNormal3fv)gl21LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl21PNormal3iv)gl21LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl21PNormal3sv)gl21LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl21POrtho)gl21LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl21PPassThrough)gl21LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl21PPixelMapfv)gl21LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl21PPixelMapuiv)gl21LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl21PPixelMapusv)gl21LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl21PPixelStoref)gl21LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl21PPixelStorei)gl21LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl21PPixelTransferf)gl21LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl21PPixelTransferi)gl21LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl21PPixelZoom)gl21LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl21PPointSize)gl21LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl21PPolygonMode)gl21LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl21PPolygonStipple)gl21LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl21PPushAttrib)gl21LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl21PPopAttrib)gl21LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl21PPushMatrix)gl21LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl21PPopMatrix)gl21LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl21PPushName)gl21LibGetProcAddress("glPushName");
    glc->fnPopName = (gl21PPopName)gl21LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl21PRasterPos2d)gl21LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl21PRasterPos2f)gl21LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl21PRasterPos2i)gl21LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl21PRasterPos2s)gl21LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl21PRasterPos3d)gl21LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl21PRasterPos3f)gl21LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl21PRasterPos3i)gl21LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl21PRasterPos3s)gl21LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl21PRasterPos4d)gl21LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl21PRasterPos4f)gl21LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl21PRasterPos4i)gl21LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl21PRasterPos4s)gl21LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl21PRasterPos2dv)gl21LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl21PRasterPos2fv)gl21LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl21PRasterPos2iv)gl21LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl21PRasterPos2sv)gl21LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl21PRasterPos3dv)gl21LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl21PRasterPos3fv)gl21LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl21PRasterPos3iv)gl21LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl21PRasterPos3sv)gl21LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl21PRasterPos4dv)gl21LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl21PRasterPos4fv)gl21LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl21PRasterPos4iv)gl21LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl21PRasterPos4sv)gl21LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl21PReadBuffer)gl21LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl21PReadPixels)gl21LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl21PRectd)gl21LibGetProcAddress("glRectd");
    glc->fnRectf = (gl21PRectf)gl21LibGetProcAddress("glRectf");
    glc->fnRecti = (gl21PRecti)gl21LibGetProcAddress("glRecti");
    glc->fnRects = (gl21PRects)gl21LibGetProcAddress("glRects");
    glc->fnRectdv = (gl21PRectdv)gl21LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl21PRectfv)gl21LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl21PRectiv)gl21LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl21PRectsv)gl21LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl21PRenderMode)gl21LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl21PRotated)gl21LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl21PRotatef)gl21LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl21PScaled)gl21LibGetProcAddress("glScaled");
    glc->fnScalef = (gl21PScalef)gl21LibGetProcAddress("glScalef");
    glc->fnScissor = (gl21PScissor)gl21LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl21PSelectBuffer)gl21LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl21PShadeModel)gl21LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl21PStencilFunc)gl21LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl21PStencilMask)gl21LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl21PStencilOp)gl21LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl21PTexCoord1d)gl21LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl21PTexCoord1f)gl21LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl21PTexCoord1i)gl21LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl21PTexCoord1s)gl21LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl21PTexCoord2d)gl21LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl21PTexCoord2f)gl21LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl21PTexCoord2i)gl21LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl21PTexCoord2s)gl21LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl21PTexCoord3d)gl21LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl21PTexCoord3f)gl21LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl21PTexCoord3i)gl21LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl21PTexCoord3s)gl21LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl21PTexCoord4d)gl21LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl21PTexCoord4f)gl21LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl21PTexCoord4i)gl21LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl21PTexCoord4s)gl21LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl21PTexCoord1dv)gl21LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl21PTexCoord1fv)gl21LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl21PTexCoord1iv)gl21LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl21PTexCoord1sv)gl21LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl21PTexCoord2dv)gl21LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl21PTexCoord2fv)gl21LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl21PTexCoord2iv)gl21LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl21PTexCoord2sv)gl21LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl21PTexCoord3dv)gl21LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl21PTexCoord3fv)gl21LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl21PTexCoord3iv)gl21LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl21PTexCoord3sv)gl21LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl21PTexCoord4dv)gl21LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl21PTexCoord4fv)gl21LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl21PTexCoord4iv)gl21LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl21PTexCoord4sv)gl21LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl21PTexEnvf)gl21LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl21PTexEnvi)gl21LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl21PTexEnvfv)gl21LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl21PTexEnviv)gl21LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl21PTexGend)gl21LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl21PTexGenf)gl21LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl21PTexGeni)gl21LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl21PTexGendv)gl21LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl21PTexGenfv)gl21LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl21PTexGeniv)gl21LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl21PTexImage1D)gl21LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl21PTexImage2D)gl21LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl21PTexParameterf)gl21LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl21PTexParameteri)gl21LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl21PTexParameterfv)gl21LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl21PTexParameteriv)gl21LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl21PTranslated)gl21LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl21PTranslatef)gl21LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl21PVertex2s)gl21LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl21PVertex2i)gl21LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl21PVertex2f)gl21LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl21PVertex2d)gl21LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl21PVertex3s)gl21LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl21PVertex3i)gl21LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl21PVertex3f)gl21LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl21PVertex3d)gl21LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl21PVertex4s)gl21LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl21PVertex4i)gl21LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl21PVertex4f)gl21LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl21PVertex4d)gl21LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl21PViewport)gl21LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl21PGetConvolutionParameterfv)gl21LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl21PGetConvolutionParameteriv)gl21LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl21PAreTexturesResident)gl21LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl21PArrayElement)gl21LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl21PDrawArrays)gl21LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl21PDrawElements)gl21LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl21PGetPointerv)gl21LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl21PPolygonOffset)gl21LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl21PCopyTexImage1D)gl21LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl21PCopyTexImage2D)gl21LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl21PCopyTexSubImage1D)gl21LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl21PCopyTexSubImage2D)gl21LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl21PBindTexture)gl21LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl21PDeleteTextures)gl21LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl21PGenTextures)gl21LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl21PIsTexture)gl21LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl21PColorPointer)gl21LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl21PEnableClientState)gl21LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl21PDisableClientState)gl21LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl21PIndexub)gl21LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl21PIndexubv)gl21LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl21PInterleavedArrays)gl21LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl21PNormalPointer)gl21LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl21PPushClientAttrib)gl21LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl21PPrioritizeTextures)gl21LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl21PPopClientAttrib)gl21LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl21PTexCoordPointer)gl21LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl21PTexSubImage1D)gl21LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl21PTexSubImage2D)gl21LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl21PVertexPointer)gl21LibGetProcAddress("glVertexPointer");
    glc->fnDrawBuffers = (gl21PDrawBuffers)gl21LibGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl21PGetActiveAttrib)gl21LibGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl21PGetActiveUniform)gl21LibGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl21PGetAttachedShaders)gl21LibGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl21PGetAttribLocation)gl21LibGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl21PGetProgramiv)gl21LibGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl21PGetProgramInfoLog)gl21LibGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl21PGetShaderiv)gl21LibGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl21PGetShaderInfoLog)gl21LibGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl21PGetShaderSource)gl21LibGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl21PGetUniformfv)gl21LibGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl21PGetUniformiv)gl21LibGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl21PGetUniformLocation)gl21LibGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl21PGetVertexAttribdv)gl21LibGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl21PGetVertexAttribfv)gl21LibGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl21PGetVertexAttribiv)gl21LibGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl21PGetVertexAttribPointerv)gl21LibGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl21PIsProgram)gl21LibGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl21PIsShader)gl21LibGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl21PLinkProgram)gl21LibGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl21PShaderSource)gl21LibGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl21PStencilFuncSeparate)gl21LibGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl21PStencilMaskSeparate)gl21LibGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl21PStencilOpSeparate)gl21LibGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl21PUniform1f)gl21LibGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl21PUniform2f)gl21LibGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl21PUniform3f)gl21LibGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl21PUniform4f)gl21LibGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl21PUniform1i)gl21LibGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl21PUniform2i)gl21LibGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl21PUniform3i)gl21LibGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl21PUniform4i)gl21LibGetProcAddress("glUniform4i");
    glc->fnUniform1fv = (gl21PUniform1fv)gl21LibGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl21PUniform2fv)gl21LibGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl21PUniform3fv)gl21LibGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl21PUniform4fv)gl21LibGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl21PUniform1iv)gl21LibGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl21PUniform2iv)gl21LibGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl21PUniform3iv)gl21LibGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl21PUniform4iv)gl21LibGetProcAddress("glUniform4iv");
    glc->fnUseProgram = (gl21PUseProgram)gl21LibGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl21PValidateProgram)gl21LibGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl21PVertexAttribPointer)gl21LibGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl21PVertexAttrib1f)gl21LibGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl21PVertexAttrib1s)gl21LibGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl21PVertexAttrib1d)gl21LibGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl21PVertexAttrib2f)gl21LibGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl21PVertexAttrib2s)gl21LibGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl21PVertexAttrib2d)gl21LibGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl21PVertexAttrib3f)gl21LibGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl21PVertexAttrib3s)gl21LibGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl21PVertexAttrib3d)gl21LibGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl21PVertexAttrib4f)gl21LibGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl21PVertexAttrib4s)gl21LibGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl21PVertexAttrib4d)gl21LibGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl21PVertexAttrib4Nuv)gl21LibGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl21PVertexAttrib1fv)gl21LibGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl21PVertexAttrib1sv)gl21LibGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl21PVertexAttrib1dv)gl21LibGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl21PVertexAttrib2fv)gl21LibGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl21PVertexAttrib2sv)gl21LibGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl21PVertexAttrib2dv)gl21LibGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl21PVertexAttrib3fv)gl21LibGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl21PVertexAttrib3sv)gl21LibGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl21PVertexAttrib3dv)gl21LibGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl21PVertexAttrib4fv)gl21LibGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl21PVertexAttrib4sv)gl21LibGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl21PVertexAttrib4dv)gl21LibGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl21PVertexAttrib4iv)gl21LibGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl21PVertexAttrib4bv)gl21LibGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl21PVertexAttrib4ubv)gl21LibGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl21PVertexAttrib4usv)gl21LibGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl21PVertexAttrib4uiv)gl21LibGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl21PVertexAttrib4Nbv)gl21LibGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl21PVertexAttrib4Nsv)gl21LibGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl21PVertexAttrib4Niv)gl21LibGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl21PVertexAttrib4Nubv)gl21LibGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl21PVertexAttrib4Nusv)gl21LibGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl21PVertexAttrib4Nuiv)gl21LibGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnUniformMatrix2fv = (gl21PUniformMatrix2fv)gl21LibGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl21PUniformMatrix3fv)gl21LibGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl21PUniformMatrix4fv)gl21LibGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl21PUniformMatrix2x3fv)gl21LibGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl21PUniformMatrix3x2fv)gl21LibGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl21PUniformMatrix2x4fv)gl21LibGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl21PUniformMatrix4x2fv)gl21LibGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl21PUniformMatrix3x4fv)gl21LibGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl21PUniformMatrix4x3fv)gl21LibGetProcAddress("glUniformMatrix4x3fv");
    glc->fnColorTable = (gl21PColorTable)gl21GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl21PColorTableParameterfv)gl21GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl21PColorTableParameteriv)gl21GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl21PColorSubTable)gl21GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl21PConvolutionFilter1D)gl21GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl21PConvolutionFilter2D)gl21GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl21PConvolutionParameterf)gl21GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl21PConvolutionParameteri)gl21GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl21PCopyColorTable)gl21GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl21PCopyColorSubTable)gl21GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl21PCopyConvolutionFilter1D)gl21GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl21PCopyConvolutionFilter2D)gl21GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl21PGetColorTable)gl21GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl21PGetColorTableParameterfv)gl21GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl21PGetColorTableParameteriv)gl21GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl21PGetConvolutionFilter)gl21GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl21PGetHistogram)gl21GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl21PGetHistogramParameterfv)gl21GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl21PGetHistogramParameteriv)gl21GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl21PGetSeparableFilter)gl21GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl21PHistogram)gl21GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl21PMinmax)gl21GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl21PMultiTexCoord1s)gl21GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl21PMultiTexCoord1i)gl21GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl21PMultiTexCoord1f)gl21GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl21PMultiTexCoord1d)gl21GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl21PMultiTexCoord2s)gl21GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl21PMultiTexCoord2i)gl21GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl21PMultiTexCoord2f)gl21GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl21PMultiTexCoord2d)gl21GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl21PMultiTexCoord3s)gl21GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl21PMultiTexCoord3i)gl21GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl21PMultiTexCoord3f)gl21GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl21PMultiTexCoord3d)gl21GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl21PMultiTexCoord4s)gl21GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl21PMultiTexCoord4i)gl21GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl21PMultiTexCoord4f)gl21GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl21PMultiTexCoord4d)gl21GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl21PMultiTexCoord1sv)gl21GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl21PMultiTexCoord1iv)gl21GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl21PMultiTexCoord1fv)gl21GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl21PMultiTexCoord1dv)gl21GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl21PMultiTexCoord2sv)gl21GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl21PMultiTexCoord2iv)gl21GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl21PMultiTexCoord2fv)gl21GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl21PMultiTexCoord2dv)gl21GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl21PMultiTexCoord3sv)gl21GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl21PMultiTexCoord3iv)gl21GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl21PMultiTexCoord3fv)gl21GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl21PMultiTexCoord3dv)gl21GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl21PMultiTexCoord4sv)gl21GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl21PMultiTexCoord4iv)gl21GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl21PMultiTexCoord4fv)gl21GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl21PMultiTexCoord4dv)gl21GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl21PResetHistogram)gl21GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl21PResetMinmax)gl21GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl21PSeparableFilter2D)gl21GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl21PBlendColor)gl21GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl21PBlendEquation)gl21GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl21PCopyTexSubImage3D)gl21GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl21PDrawRangeElements)gl21GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl21PTexImage3D)gl21GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl21PTexSubImage3D)gl21GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl21PActiveTexture)gl21GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl21PClientActiveTexture)gl21GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl21PCompressedTexImage1D)gl21GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl21PCompressedTexImage2D)gl21GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl21PCompressedTexImage3D)gl21GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl21PCompressedTexSubImage1D)gl21GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl21PCompressedTexSubImage2D)gl21GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl21PCompressedTexSubImage3D)gl21GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl21PGetCompressedTexImage)gl21GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl21PLoadTransposeMatrixd)gl21GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl21PLoadTransposeMatrixf)gl21GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl21PMultTransposeMatrixd)gl21GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl21PMultTransposeMatrixf)gl21GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl21PSampleCoverage)gl21GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl21PBlendFuncSeparate)gl21GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl21PFogCoordPointer)gl21GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl21PFogCoordd)gl21GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl21PFogCoordf)gl21GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl21PFogCoorddv)gl21GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl21PFogCoordfv)gl21GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl21PMultiDrawArrays)gl21GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl21PMultiDrawElements)gl21GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl21PPointParameterf)gl21GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl21PPointParameteri)gl21GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl21PSecondaryColor3b)gl21GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl21PSecondaryColor3s)gl21GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl21PSecondaryColor3i)gl21GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl21PSecondaryColor3f)gl21GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl21PSecondaryColor3d)gl21GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl21PSecondaryColor3ub)gl21GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl21PSecondaryColor3us)gl21GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl21PSecondaryColor3ui)gl21GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl21PSecondaryColor3bv)gl21GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl21PSecondaryColor3sv)gl21GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl21PSecondaryColor3iv)gl21GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl21PSecondaryColor3fv)gl21GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl21PSecondaryColor3dv)gl21GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl21PSecondaryColor3ubv)gl21GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl21PSecondaryColor3usv)gl21GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl21PSecondaryColor3uiv)gl21GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl21PSecondaryColorPointer)gl21GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl21PWindowPos2s)gl21GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl21PWindowPos2i)gl21GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl21PWindowPos2f)gl21GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl21PWindowPos2d)gl21GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl21PWindowPos3s)gl21GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl21PWindowPos3i)gl21GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl21PWindowPos3f)gl21GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl21PWindowPos3d)gl21GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl21PWindowPos2sv)gl21GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl21PWindowPos2iv)gl21GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl21PWindowPos2fv)gl21GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl21PWindowPos2dv)gl21GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl21PWindowPos3sv)gl21GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl21PWindowPos3iv)gl21GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl21PWindowPos3fv)gl21GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl21PWindowPos3dv)gl21GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl21PBeginQuery)gl21GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl21PBindBuffer)gl21GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl21PBufferData)gl21GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl21PBufferSubData)gl21GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl21PDeleteBuffers)gl21GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl21PDeleteQueries)gl21GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl21PGenBuffers)gl21GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl21PGenQueries)gl21GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl21PGetBufferParameteriv)gl21GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl21PGetBufferPointerv)gl21GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl21PGetBufferSubData)gl21GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl21PGetQueryObjectiv)gl21GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl21PGetQueryObjectuiv)gl21GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl21PGetQueryiv)gl21GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl21PIsBuffer)gl21GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl21PIsQuery)gl21GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl21PMapBuffer)gl21GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl21PUnmapBuffer)gl21GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl21PAttachShader)gl21GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl21PBindAttribLocation)gl21GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl21PBlendEquationSeperate)gl21GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl21PCompileShader)gl21GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl21PCreateProgram)gl21GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl21PCreateShader)gl21GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl21PDeleteProgram)gl21GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl21PDeleteShader)gl21GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl21PDetachShader)gl21GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl21PEnableVertexAttribArray)gl21GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl21PDisableVertexAttribArray)gl21GLGetProcAddress("glDisableVertexAttribArray");
    return glc;
}

