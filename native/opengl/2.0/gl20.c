
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#endif

#include "gl20.h"

#ifdef _WIN32
HMODULE gl20OpenGL32;

void* gl20LibGetProcAddress(char* name) {
	if(gl20OpenGL32 == NULL) {
		gl20OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
	}
	return GetProcAddress(gl20OpenGL32, TEXT(name));
}

void* gl20GLGetProcAddress(char* name) {
	void* ptr = wglGetProcAddress(name);

	intptr_t iptr = (intptr_t)ptr;

	if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
		return NULL;
	}
	return ptr;
}
#endif


void gl20Accum(gl20Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl20AlphaFunc(gl20Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl20Begin(gl20Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl20End(gl20Context* glc) {
    return glc->fnEnd();
}

void gl20Bitmap(gl20Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl20BlendFunc(gl20Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl20CallList(gl20Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl20CallLists(gl20Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl20Clear(gl20Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl20ClearAccum(gl20Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl20ClearColor(gl20Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl20ClearDepth(gl20Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl20ClearDepthf(gl20Context* glc, GLclampf depth) {
    return glc->fnClearDepthf(depth);
}

void gl20ClearIndex(gl20Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl20ClearStencil(gl20Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl20ClipPlane(gl20Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl20Color3b(gl20Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl20Color3d(gl20Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl20Color3f(gl20Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl20Color3i(gl20Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl20Color3s(gl20Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl20Color3ub(gl20Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl20Color3ui(gl20Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl20Color3us(gl20Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl20Color4b(gl20Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl20Color4d(gl20Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl20Color4f(gl20Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl20Color4i(gl20Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl20Color4s(gl20Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl20Color4ub(gl20Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl20Color4ui(gl20Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl20Color4us(gl20Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl20Color3bv(gl20Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl20Color3dv(gl20Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl20Color3fv(gl20Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl20Color3iv(gl20Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl20Color3sv(gl20Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl20Color3ubv(gl20Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl20Color3uiv(gl20Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl20Color3usv(gl20Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl20Color4bv(gl20Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl20Color4dv(gl20Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl20Color4fv(gl20Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl20Color4iv(gl20Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl20Color4sv(gl20Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl20Color4ubv(gl20Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl20Color4uiv(gl20Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl20Color4usv(gl20Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl20ColorMask(gl20Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl20ColorMaterial(gl20Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl20CopyPixels(gl20Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl20CullFace(gl20Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl20DeleteLists(gl20Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl20DepthFunc(gl20Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl20DepthMask(gl20Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
}

void gl20DepthRange(gl20Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl20Enable(gl20Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl20Disable(gl20Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl20DrawBuffer(gl20Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl20DrawPixels(gl20Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
}

void gl20EdgeFlag(gl20Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl20EdgeFlagv(gl20Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl20EdgeFlagPointer(gl20Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl20EvalCoord1d(gl20Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl20EvalCoord1f(gl20Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl20EvalCoord2d(gl20Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl20EvalCoord2f(gl20Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl20EvalCoord1dv(gl20Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl20EvalCoord1fv(gl20Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl20EvalCoord2dv(gl20Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl20EvalCoord2fv(gl20Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl20EvalMesh1(gl20Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl20EvalMesh2(gl20Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl20EvalPoint1(gl20Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl20EvalPoint2(gl20Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl20FeedbackBuffer(gl20Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl20Finish(gl20Context* glc) {
    return glc->fnFinish();
}

void gl20Flush(gl20Context* glc) {
    return glc->fnFlush();
}

void gl20Fogf(gl20Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl20Fogi(gl20Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl20Fogfv(gl20Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl20Fogiv(gl20Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl20FrontFace(gl20Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl20Frustum(gl20Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl20GenLists(gl20Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl20GetBooleanv(gl20Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl20GetDoublev(gl20Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl20GetFloatv(gl20Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl20GetIntegerv(gl20Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl20GetClipPlane(gl20Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl20GetError(gl20Context* glc) {
    return glc->fnGetError();
}

void gl20GetLightfv(gl20Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl20GetLightiv(gl20Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl20GetMapdv(gl20Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl20GetMapfv(gl20Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl20GetMapiv(gl20Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl20GetMaterialfv(gl20Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl20GetMaterialiv(gl20Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl20GetPixelMapfv(gl20Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl20GetPixelMapuiv(gl20Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl20GetPixelMapusv(gl20Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl20GetPolygonStipple(gl20Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
}

GLubyte* gl20GetString(gl20Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl20GetTexEnvfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl20GetTexEnviv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl20GetTexGendv(gl20Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl20GetTexGenfv(gl20Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl20GetTexGeniv(gl20Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl20GetTexImage(gl20Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl20GetTexLevelParameterfv(gl20Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl20GetTexLevelParameteriv(gl20Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl20GetTexParameterfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl20GetTexParameteriv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl20Hint(gl20Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl20Indexd(gl20Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl20Indexf(gl20Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl20Indexi(gl20Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl20Indexs(gl20Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl20Indexdv(gl20Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl20Indexfv(gl20Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl20Indexiv(gl20Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl20Indexsv(gl20Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl20IndexMask(gl20Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl20IndexPointer(gl20Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl20InitNames(gl20Context* glc) {
    return glc->fnInitNames();
}

void gl20IsEnabled(gl20Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl20IsList(gl20Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl20Lightf(gl20Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl20Lighti(gl20Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl20Lightfv(gl20Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl20Lightiv(gl20Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl20LightModelf(gl20Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl20LightModeli(gl20Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl20LightModelfv(gl20Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl20LightModeliv(gl20Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl20LineStipple(gl20Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl20LineWidth(gl20Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl20ListBase(gl20Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl20LoadIdentity(gl20Context* glc) {
    return glc->fnLoadIdentity();
}

void gl20LoadMatrixd(gl20Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl20LoadMatrixf(gl20Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl20LoadName(gl20Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl20LogicOp(gl20Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl20Map1d(gl20Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl20Map1f(gl20Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl20Map2d(gl20Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl20Map2f(gl20Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl20MapGrid1d(gl20Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl20MapGrid1f(gl20Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl20MapGrid2d(gl20Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl20MapGrid2f(gl20Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl20Materialf(gl20Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl20Materiali(gl20Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl20Materialfv(gl20Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl20Materialiv(gl20Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl20MatrixMode(gl20Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl20MultMatrixd(gl20Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl20MultMatrixf(gl20Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl20NewList(gl20Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl20EndList(gl20Context* glc) {
    return glc->fnEndList();
}

void gl20Normal3b(gl20Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl20Normal3d(gl20Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl20Normal3f(gl20Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl20Normal3i(gl20Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl20Normal3s(gl20Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl20Normal3bv(gl20Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl20Normal3dv(gl20Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl20Normal3fv(gl20Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl20Normal3iv(gl20Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl20Normal3sv(gl20Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl20Ortho(gl20Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl20PassThrough(gl20Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl20PixelMapfv(gl20Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl20PixelMapuiv(gl20Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl20PixelMapusv(gl20Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl20PixelStoref(gl20Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl20PixelStorei(gl20Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl20PixelTransferf(gl20Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl20PixelTransferi(gl20Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl20PixelZoom(gl20Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl20PointSize(gl20Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl20PolygonMode(gl20Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl20PolygonStipple(gl20Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl20PushAttrib(gl20Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl20PopAttrib(gl20Context* glc) {
    return glc->fnPopAttrib();
}

void gl20PushMatrix(gl20Context* glc) {
    return glc->fnPushMatrix();
}

void gl20PopMatrix(gl20Context* glc) {
    return glc->fnPopMatrix();
}

void gl20PushName(gl20Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl20PopName(gl20Context* glc) {
    return glc->fnPopName();
}

void gl20RasterPos2d(gl20Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl20RasterPos2f(gl20Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl20RasterPos2i(gl20Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl20RasterPos2s(gl20Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl20RasterPos3d(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl20RasterPos3f(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl20RasterPos3i(gl20Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl20RasterPos3s(gl20Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl20RasterPos4d(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl20RasterPos4f(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl20RasterPos4i(gl20Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl20RasterPos4s(gl20Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl20RasterPos2dv(gl20Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl20RasterPos2fv(gl20Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl20RasterPos2iv(gl20Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl20RasterPos2sv(gl20Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl20RasterPos3dv(gl20Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl20RasterPos3fv(gl20Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl20RasterPos3iv(gl20Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl20RasterPos3sv(gl20Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl20RasterPos4dv(gl20Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl20RasterPos4fv(gl20Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl20RasterPos4iv(gl20Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl20RasterPos4sv(gl20Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl20ReadBuffer(gl20Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl20ReadPixels(gl20Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl20Rectd(gl20Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl20Rectf(gl20Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl20Recti(gl20Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl20Rects(gl20Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl20Rectdv(gl20Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl20Rectfv(gl20Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl20Rectiv(gl20Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl20Rectsv(gl20Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl20RenderMode(gl20Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl20Rotated(gl20Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl20Rotatef(gl20Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl20Scaled(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl20Scalef(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl20Scissor(gl20Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl20SelectBuffer(gl20Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl20ShadeModel(gl20Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl20StencilFunc(gl20Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl20StencilMask(gl20Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl20StencilOp(gl20Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl20TexCoord1d(gl20Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl20TexCoord1f(gl20Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl20TexCoord1i(gl20Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl20TexCoord1s(gl20Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl20TexCoord2d(gl20Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl20TexCoord2f(gl20Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl20TexCoord2i(gl20Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl20TexCoord2s(gl20Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl20TexCoord3d(gl20Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl20TexCoord3f(gl20Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl20TexCoord3i(gl20Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl20TexCoord3s(gl20Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl20TexCoord4d(gl20Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl20TexCoord4f(gl20Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl20TexCoord4i(gl20Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl20TexCoord4s(gl20Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl20TexCoord1dv(gl20Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl20TexCoord1fv(gl20Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl20TexCoord1iv(gl20Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl20TexCoord1sv(gl20Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl20TexCoord2dv(gl20Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl20TexCoord2fv(gl20Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl20TexCoord2iv(gl20Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl20TexCoord2sv(gl20Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl20TexCoord3dv(gl20Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl20TexCoord3fv(gl20Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl20TexCoord3iv(gl20Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl20TexCoord3sv(gl20Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl20TexCoord4dv(gl20Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl20TexCoord4fv(gl20Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl20TexCoord4iv(gl20Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl20TexCoord4sv(gl20Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl20TexEnvf(gl20Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl20TexEnvi(gl20Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl20TexEnvfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl20TexEnviv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl20TexGend(gl20Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl20TexGenf(gl20Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl20TexGeni(gl20Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl20TexGendv(gl20Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl20TexGenfv(gl20Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl20TexGeniv(gl20Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl20TexImage1D(gl20Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl20TexImage2D(gl20Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl20TexParameterf(gl20Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl20TexParameteri(gl20Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl20TexParameterfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl20TexParameteriv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl20Translated(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl20Translatef(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl20Vertex2s(gl20Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl20Vertex2i(gl20Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl20Vertex2f(gl20Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl20Vertex2d(gl20Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl20Vertex3s(gl20Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl20Vertex3i(gl20Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl20Vertex3f(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl20Vertex3d(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl20Vertex4s(gl20Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl20Vertex4i(gl20Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl20Vertex4f(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl20Vertex4d(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl20Viewport(gl20Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl20GetConvolutionParameterfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl20GetConvolutionParameteriv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

GLboolean gl20AreTexturesResident(gl20Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl20ArrayElement(gl20Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl20DrawArrays(gl20Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl20DrawElements(gl20Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl20GetPointerv(gl20Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl20PolygonOffset(gl20Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl20CopyTexImage1D(gl20Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl20CopyTexImage2D(gl20Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl20CopyTexSubImage1D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl20CopyTexSubImage2D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl20BindTexture(gl20Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl20DeleteTextures(gl20Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl20GenTextures(gl20Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl20IsTexture(gl20Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl20ColorPointer(gl20Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl20EnableClientState(gl20Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl20DisableClientState(gl20Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl20Indexub(gl20Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl20Indexubv(gl20Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl20InterleavedArrays(gl20Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl20NormalPointer(gl20Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl20PushClientAttrib(gl20Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl20PrioritizeTextures(gl20Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl20PopClientAttrib(gl20Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl20TexCoordPointer(gl20Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl20TexSubImage1D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl20TexSubImage2D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl20VertexPointer(gl20Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl20DrawBuffers(gl20Context* glc, GLsizei n, GLenum* bufs) {
    return glc->fnDrawBuffers(n, bufs);
}

void gl20GetActiveAttrib(gl20Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveAttrib(program, index, bufSize, length, size, type, name);
}

void gl20GetActiveUniform(gl20Context* glc, GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
    return glc->fnGetActiveUniform(program, index, bufSize, length, size, type, name);
}

void gl20GetAttachedShaders(gl20Context* glc, GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
    return glc->fnGetAttachedShaders(program, maxCount, count, shaders);
}

GLint gl20GetAttribLocation(gl20Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetAttribLocation(program, name);
}

void gl20GetProgramiv(gl20Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetProgramiv(program, pname, params);
}

void gl20GetProgramInfoLog(gl20Context* glc, GLuint program, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetProgramInfoLog(program, maxLength, length, infoLog);
}

void gl20GetShaderiv(gl20Context* glc, GLuint program, GLenum pname, GLint* params) {
    return glc->fnGetShaderiv(program, pname, params);
}

void gl20GetShaderInfoLog(gl20Context* glc, GLuint shader, GLsizei maxLength, GLsizei* length, GLchar* infoLog) {
    return glc->fnGetShaderInfoLog(shader, maxLength, length, infoLog);
}

void gl20GetShaderSource(gl20Context* glc, GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
    return glc->fnGetShaderSource(shader, bufSize, length, source);
}

void gl20GetUniformfv(gl20Context* glc, GLuint program, GLint location, GLfloat* params) {
    return glc->fnGetUniformfv(program, location, params);
}

void gl20GetUniformiv(gl20Context* glc, GLuint program, GLint location, GLint* params) {
    return glc->fnGetUniformiv(program, location, params);
}

GLint gl20GetUniformLocation(gl20Context* glc, GLuint program, GLchar* name) {
    return glc->fnGetUniformLocation(program, name);
}

void gl20GetVertexAttribdv(gl20Context* glc, GLuint index, GLenum pname, GLdouble* params) {
    return glc->fnGetVertexAttribdv(index, pname, params);
}

void gl20GetVertexAttribfv(gl20Context* glc, GLuint index, GLenum pname, GLfloat* params) {
    return glc->fnGetVertexAttribfv(index, pname, params);
}

void gl20GetVertexAttribiv(gl20Context* glc, GLuint index, GLenum pname, GLint* params) {
    return glc->fnGetVertexAttribiv(index, pname, params);
}

void gl20GetVertexAttribPointerv(gl20Context* glc, GLuint index, GLenum pname, GLvoid* pointer) {
    return glc->fnGetVertexAttribPointerv(index, pname, pointer);
}

GLboolean gl20IsProgram(gl20Context* glc, GLuint program) {
    return glc->fnIsProgram(program);
}

GLboolean gl20IsShader(gl20Context* glc, GLuint shader) {
    return glc->fnIsShader(shader);
}

void gl20LinkProgram(gl20Context* glc, GLuint program) {
    return glc->fnLinkProgram(program);
}

void gl20ShaderSource(gl20Context* glc, GLuint shader, GLsizei count, GLchar** string, GLint* length) {
    return glc->fnShaderSource(shader, count, string, length);
}

void gl20StencilFuncSeparate(gl20Context* glc, GLenum face, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFuncSeparate(face, func, ref, mask);
}

void gl20StencilMaskSeparate(gl20Context* glc, GLenum face, GLuint mask) {
    return glc->fnStencilMaskSeparate(face, mask);
}

void gl20StencilOpSeparate(gl20Context* glc, GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
    return glc->fnStencilOpSeparate(face, sfail, dpfail, dppass);
}

void gl20Uniform1f(gl20Context* glc, GLint location, GLfloat v0) {
    return glc->fnUniform1f(location, v0);
}

void gl20Uniform2f(gl20Context* glc, GLint location, GLfloat v0, GLfloat v1) {
    return glc->fnUniform2f(location, v0, v1);
}

void gl20Uniform3f(gl20Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnUniform3f(location, v0, v1, v2);
}

void gl20Uniform4f(gl20Context* glc, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnUniform4f(location, v0, v1, v2, v3);
}

void gl20Uniform1i(gl20Context* glc, GLint location, GLint v0) {
    return glc->fnUniform1i(location, v0);
}

void gl20Uniform2i(gl20Context* glc, GLint location, GLint v0, GLint v1) {
    return glc->fnUniform2i(location, v0, v1);
}

void gl20Uniform3i(gl20Context* glc, GLint location, GLint v0, GLint v1, GLint v2) {
    return glc->fnUniform3i(location, v0, v1, v2);
}

void gl20Uniform4i(gl20Context* glc, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
    return glc->fnUniform4i(location, v0, v1, v2, v3);
}

void gl20Uniform1fv(gl20Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform1fv(location, count, value);
}

void gl20Uniform2fv(gl20Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform2fv(location, count, value);
}

void gl20Uniform3fv(gl20Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform3fv(location, count, value);
}

void gl20Uniform4fv(gl20Context* glc, GLint location, GLsizei count, GLfloat* value) {
    return glc->fnUniform4fv(location, count, value);
}

void gl20Uniform1iv(gl20Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform1iv(location, count, value);
}

void gl20Uniform2iv(gl20Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform2iv(location, count, value);
}

void gl20Uniform3iv(gl20Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform3iv(location, count, value);
}

void gl20Uniform4iv(gl20Context* glc, GLint location, GLsizei count, GLint* value) {
    return glc->fnUniform4iv(location, count, value);
}

void gl20UseProgram(gl20Context* glc, GLuint program) {
    return glc->fnUseProgram(program);
}

void gl20ValidateProgram(gl20Context* glc, GLuint program) {
    return glc->fnValidateProgram(program);
}

void gl20VertexAttribPointer(gl20Context* glc, GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexAttribPointer(index, size, type, normalized, stride, pointer);
}

void gl20VertexAttrib1f(gl20Context* glc, GLuint index, GLfloat v0) {
    return glc->fnVertexAttrib1f(index, v0);
}

void gl20VertexAttrib1s(gl20Context* glc, GLuint index, GLshort v0) {
    return glc->fnVertexAttrib1s(index, v0);
}

void gl20VertexAttrib1d(gl20Context* glc, GLuint index, GLdouble v0) {
    return glc->fnVertexAttrib1d(index, v0);
}

void gl20VertexAttrib2f(gl20Context* glc, GLuint index, GLfloat v0, GLfloat v1) {
    return glc->fnVertexAttrib2f(index, v0, v1);
}

void gl20VertexAttrib2s(gl20Context* glc, GLuint index, GLshort v0, GLshort v1) {
    return glc->fnVertexAttrib2s(index, v0, v1);
}

void gl20VertexAttrib2d(gl20Context* glc, GLuint index, GLdouble v0, GLdouble v1) {
    return glc->fnVertexAttrib2d(index, v0, v1);
}

void gl20VertexAttrib3f(gl20Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2) {
    return glc->fnVertexAttrib3f(index, v0, v1, v2);
}

void gl20VertexAttrib3s(gl20Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2) {
    return glc->fnVertexAttrib3s(index, v0, v1, v2);
}

void gl20VertexAttrib3d(gl20Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2) {
    return glc->fnVertexAttrib3d(index, v0, v1, v2);
}

void gl20VertexAttrib4f(gl20Context* glc, GLuint index, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
    return glc->fnVertexAttrib4f(index, v0, v1, v2, v3);
}

void gl20VertexAttrib4s(gl20Context* glc, GLuint index, GLshort v0, GLshort v1, GLshort v2, GLshort v3) {
    return glc->fnVertexAttrib4s(index, v0, v1, v2, v3);
}

void gl20VertexAttrib4d(gl20Context* glc, GLuint index, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {
    return glc->fnVertexAttrib4d(index, v0, v1, v2, v3);
}

void gl20VertexAttrib4Nuv(gl20Context* glc, GLuint index, GLubyte v0, GLubyte v1, GLubyte v2, GLubyte v3) {
    return glc->fnVertexAttrib4Nuv(index, v0, v1, v2, v3);
}

void gl20VertexAttrib1fv(gl20Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib1fv(index, v);
}

void gl20VertexAttrib1sv(gl20Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib1sv(index, v);
}

void gl20VertexAttrib1dv(gl20Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib1dv(index, v);
}

void gl20VertexAttrib2fv(gl20Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib2fv(index, v);
}

void gl20VertexAttrib2sv(gl20Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib2sv(index, v);
}

void gl20VertexAttrib2dv(gl20Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib2dv(index, v);
}

void gl20VertexAttrib3fv(gl20Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib3fv(index, v);
}

void gl20VertexAttrib3sv(gl20Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib3sv(index, v);
}

void gl20VertexAttrib3dv(gl20Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib3dv(index, v);
}

void gl20VertexAttrib4fv(gl20Context* glc, GLuint index, GLfloat* v) {
    return glc->fnVertexAttrib4fv(index, v);
}

void gl20VertexAttrib4sv(gl20Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4sv(index, v);
}

void gl20VertexAttrib4dv(gl20Context* glc, GLuint index, GLdouble* v) {
    return glc->fnVertexAttrib4dv(index, v);
}

void gl20VertexAttrib4iv(gl20Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4iv(index, v);
}

void gl20VertexAttrib4bv(gl20Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4bv(index, v);
}

void gl20VertexAttrib4ubv(gl20Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4ubv(index, v);
}

void gl20VertexAttrib4usv(gl20Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4usv(index, v);
}

void gl20VertexAttrib4uiv(gl20Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4uiv(index, v);
}

void gl20VertexAttrib4Nbv(gl20Context* glc, GLuint index, GLbyte* v) {
    return glc->fnVertexAttrib4Nbv(index, v);
}

void gl20VertexAttrib4Nsv(gl20Context* glc, GLuint index, GLshort* v) {
    return glc->fnVertexAttrib4Nsv(index, v);
}

void gl20VertexAttrib4Niv(gl20Context* glc, GLuint index, GLint* v) {
    return glc->fnVertexAttrib4Niv(index, v);
}

void gl20VertexAttrib4Nubv(gl20Context* glc, GLuint index, GLubyte* v) {
    return glc->fnVertexAttrib4Nubv(index, v);
}

void gl20VertexAttrib4Nusv(gl20Context* glc, GLuint index, GLushort* v) {
    return glc->fnVertexAttrib4Nusv(index, v);
}

void gl20VertexAttrib4Nuiv(gl20Context* glc, GLuint index, GLuint* v) {
    return glc->fnVertexAttrib4Nuiv(index, v);
}

void gl20ColorTable(gl20Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl20ColorTableParameterfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl20ColorTableParameteriv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl20ColorSubTable(gl20Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl20ConvolutionFilter1D(gl20Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl20ConvolutionFilter2D(gl20Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl20ConvolutionParameterf(gl20Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl20ConvolutionParameteri(gl20Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl20CopyColorTable(gl20Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl20CopyColorSubTable(gl20Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl20CopyConvolutionFilter1D(gl20Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl20CopyConvolutionFilter2D(gl20Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl20GetColorTable(gl20Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl20GetColorTableParameterfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl20GetColorTableParameteriv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl20GetConvolutionFilter(gl20Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl20GetHistogram(gl20Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl20GetHistogramParameterfv(gl20Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl20GetHistogramParameteriv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl20GetSeparableFilter(gl20Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl20Histogram(gl20Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl20Minmax(gl20Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl20MultiTexCoord1s(gl20Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl20MultiTexCoord1i(gl20Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl20MultiTexCoord1f(gl20Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl20MultiTexCoord1d(gl20Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl20MultiTexCoord2s(gl20Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl20MultiTexCoord2i(gl20Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl20MultiTexCoord2f(gl20Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl20MultiTexCoord2d(gl20Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl20MultiTexCoord3s(gl20Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl20MultiTexCoord3i(gl20Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl20MultiTexCoord3f(gl20Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl20MultiTexCoord3d(gl20Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl20MultiTexCoord4s(gl20Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl20MultiTexCoord4i(gl20Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl20MultiTexCoord4f(gl20Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl20MultiTexCoord4d(gl20Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl20MultiTexCoord1sv(gl20Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl20MultiTexCoord1iv(gl20Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl20MultiTexCoord1fv(gl20Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl20MultiTexCoord1dv(gl20Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl20MultiTexCoord2sv(gl20Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl20MultiTexCoord2iv(gl20Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl20MultiTexCoord2fv(gl20Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl20MultiTexCoord2dv(gl20Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl20MultiTexCoord3sv(gl20Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl20MultiTexCoord3iv(gl20Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl20MultiTexCoord3fv(gl20Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl20MultiTexCoord3dv(gl20Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl20MultiTexCoord4sv(gl20Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl20MultiTexCoord4iv(gl20Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl20MultiTexCoord4fv(gl20Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl20MultiTexCoord4dv(gl20Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl20ResetHistogram(gl20Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl20ResetMinmax(gl20Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl20SeparableFilter2D(gl20Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
}

void gl20BlendColor(gl20Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl20BlendEquation(gl20Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl20CopyTexSubImage3D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl20DrawRangeElements(gl20Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl20TexImage3D(gl20Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl20TexSubImage3D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl20ActiveTexture(gl20Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl20ClientActiveTexture(gl20Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl20CompressedTexImage1D(gl20Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl20CompressedTexImage2D(gl20Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl20CompressedTexImage3D(gl20Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl20CompressedTexSubImage1D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl20CompressedTexSubImage2D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl20CompressedTexSubImage3D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl20GetCompressedTexImage(gl20Context* glc, GLenum target, GLint lod, GLvoid* img) {
    return glc->fnGetCompressedTexImage(target, lod, img);
}

void gl20LoadTransposeMatrixd(gl20Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixd(m);
}

void gl20LoadTransposeMatrixf(gl20Context* glc, GLdouble* m) {
    return glc->fnLoadTransposeMatrixf(m);
}

void gl20MultTransposeMatrixd(gl20Context* glc, GLdouble* m) {
    return glc->fnMultTransposeMatrixd(m);
}

void gl20MultTransposeMatrixf(gl20Context* glc, GLfloat* m) {
    return glc->fnMultTransposeMatrixf(m);
}

void gl20SampleCoverage(gl20Context* glc, GLclampf value, GLboolean invert) {
    return glc->fnSampleCoverage(value, invert);
}

void gl20BlendFuncSeparate(gl20Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

void gl20FogCoordPointer(gl20Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnFogCoordPointer(type, stride, pointer);
}

void gl20FogCoordd(gl20Context* glc, GLdouble coord) {
    return glc->fnFogCoordd(coord);
}

void gl20FogCoordf(gl20Context* glc, GLfloat coord) {
    return glc->fnFogCoordf(coord);
}

void gl20FogCoorddv(gl20Context* glc, GLdouble* coord) {
    return glc->fnFogCoorddv(coord);
}

void gl20FogCoordfv(gl20Context* glc, GLfloat* coord) {
    return glc->fnFogCoordfv(coord);
}

void gl20MultiDrawArrays(gl20Context* glc, GLenum mode, GLint* first, GLsizei* count, GLsizei primcount) {
    return glc->fnMultiDrawArrays(mode, first, count, primcount);
}

void gl20MultiDrawElements(gl20Context* glc, GLenum mode, GLsizei* count, GLenum type, GLvoid* indices, GLsizei primcount) {
    return glc->fnMultiDrawElements(mode, count, type, indices, primcount);
}

void gl20PointParameterf(gl20Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPointParameterf(pname, param);
}

void gl20PointParameteri(gl20Context* glc, GLenum pname, GLint param) {
    return glc->fnPointParameteri(pname, param);
}

void gl20SecondaryColor3b(gl20Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnSecondaryColor3b(red, green, blue);
}

void gl20SecondaryColor3s(gl20Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnSecondaryColor3s(red, green, blue);
}

void gl20SecondaryColor3i(gl20Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnSecondaryColor3i(red, green, blue);
}

void gl20SecondaryColor3f(gl20Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnSecondaryColor3f(red, green, blue);
}

void gl20SecondaryColor3d(gl20Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnSecondaryColor3d(red, green, blue);
}

void gl20SecondaryColor3ub(gl20Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnSecondaryColor3ub(red, green, blue);
}

void gl20SecondaryColor3us(gl20Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnSecondaryColor3us(red, green, blue);
}

void gl20SecondaryColor3ui(gl20Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnSecondaryColor3ui(red, green, blue);
}

void gl20SecondaryColor3bv(gl20Context* glc, GLbyte* v) {
    return glc->fnSecondaryColor3bv(v);
}

void gl20SecondaryColor3sv(gl20Context* glc, GLshort* v) {
    return glc->fnSecondaryColor3sv(v);
}

void gl20SecondaryColor3iv(gl20Context* glc, GLint* v) {
    return glc->fnSecondaryColor3iv(v);
}

void gl20SecondaryColor3fv(gl20Context* glc, GLfloat* v) {
    return glc->fnSecondaryColor3fv(v);
}

void gl20SecondaryColor3dv(gl20Context* glc, GLdouble* v) {
    return glc->fnSecondaryColor3dv(v);
}

void gl20SecondaryColor3ubv(gl20Context* glc, GLubyte* v) {
    return glc->fnSecondaryColor3ubv(v);
}

void gl20SecondaryColor3usv(gl20Context* glc, GLushort* v) {
    return glc->fnSecondaryColor3usv(v);
}

void gl20SecondaryColor3uiv(gl20Context* glc, GLuint* v) {
    return glc->fnSecondaryColor3uiv(v);
}

void gl20SecondaryColorPointer(gl20Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnSecondaryColorPointer(size, type, stride, pointer);
}

void gl20WindowPos2s(gl20Context* glc, GLshort x, GLshort y) {
    return glc->fnWindowPos2s(x, y);
}

void gl20WindowPos2i(gl20Context* glc, GLint x, GLint y) {
    return glc->fnWindowPos2i(x, y);
}

void gl20WindowPos2f(gl20Context* glc, GLfloat x, GLfloat y) {
    return glc->fnWindowPos2f(x, y);
}

void gl20WindowPos2d(gl20Context* glc, GLdouble x, GLdouble y) {
    return glc->fnWindowPos2d(x, y);
}

void gl20WindowPos3s(gl20Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnWindowPos3s(x, y, z);
}

void gl20WindowPos3i(gl20Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnWindowPos3i(x, y, z);
}

void gl20WindowPos3f(gl20Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnWindowPos3f(x, y, z);
}

void gl20WindowPos3d(gl20Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnWindowPos3d(x, y, z);
}

void gl20WindowPos2sv(gl20Context* glc, GLshort* v) {
    return glc->fnWindowPos2sv(v);
}

void gl20WindowPos2iv(gl20Context* glc, GLint* v) {
    return glc->fnWindowPos2iv(v);
}

void gl20WindowPos2fv(gl20Context* glc, GLfloat* v) {
    return glc->fnWindowPos2fv(v);
}

void gl20WindowPos2dv(gl20Context* glc, GLdouble* v) {
    return glc->fnWindowPos2dv(v);
}

void gl20WindowPos3sv(gl20Context* glc, GLshort* v) {
    return glc->fnWindowPos3sv(v);
}

void gl20WindowPos3iv(gl20Context* glc, GLint* v) {
    return glc->fnWindowPos3iv(v);
}

void gl20WindowPos3fv(gl20Context* glc, GLfloat* v) {
    return glc->fnWindowPos3fv(v);
}

void gl20WindowPos3dv(gl20Context* glc, GLdouble* v) {
    return glc->fnWindowPos3dv(v);
}

void gl20BeginQuery(gl20Context* glc, GLenum target, GLuint id) {
    return glc->fnBeginQuery(target, id);
}

void gl20BindBuffer(gl20Context* glc, GLenum target, GLuint buffer) {
    return glc->fnBindBuffer(target, buffer);
}

void gl20BufferData(gl20Context* glc, GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {
    return glc->fnBufferData(target, size, data, usage);
}

void gl20BufferSubData(gl20Context* glc, GLenum target, GLenum offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnBufferSubData(target, offset, size, data);
}

void gl20DeleteBuffers(gl20Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnDeleteBuffers(n, buffers);
}

void gl20DeleteQueries(gl20Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnDeleteQueries(n, ids);
}

void gl20GenBuffers(gl20Context* glc, GLsizei n, GLuint* buffers) {
    return glc->fnGenBuffers(n, buffers);
}

void gl20GenQueries(gl20Context* glc, GLsizei n, GLuint* ids) {
    return glc->fnGenQueries(n, ids);
}

void gl20GetBufferParameteriv(gl20Context* glc, GLenum target, GLenum value, GLint* data) {
    return glc->fnGetBufferParameteriv(target, value, data);
}

void gl20GetBufferPointerv(gl20Context* glc, GLenum target, GLenum pname, GLvoid* params) {
    return glc->fnGetBufferPointerv(target, pname, params);
}

void gl20GetBufferSubData(gl20Context* glc, GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
    return glc->fnGetBufferSubData(target, offset, size, data);
}

void gl20GetQueryObjectiv(gl20Context* glc, GLuint id, GLenum pname, GLint* params) {
    return glc->fnGetQueryObjectiv(id, pname, params);
}

void gl20GetQueryObjectuiv(gl20Context* glc, GLuint id, GLenum pname, GLuint* params) {
    return glc->fnGetQueryObjectuiv(id, pname, params);
}

void gl20GetQueryiv(gl20Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetQueryiv(target, pname, params);
}

GLboolean gl20IsBuffer(gl20Context* glc, GLuint buffer) {
    return glc->fnIsBuffer(buffer);
}

GLboolean gl20IsQuery(gl20Context* glc, GLuint id) {
    return glc->fnIsQuery(id);
}

GLvoid* gl20MapBuffer(gl20Context* glc, GLenum target, GLenum access) {
    return glc->fnMapBuffer(target, access);
}

GLboolean gl20UnmapBuffer(gl20Context* glc, GLenum target) {
    return glc->fnUnmapBuffer(target);
}

void gl20AttachShader(gl20Context* glc, GLuint program, GLuint shader) {
    return glc->fnAttachShader(program, shader);
}

void gl20BindAttribLocation(gl20Context* glc, GLuint program, GLuint index, GLchar* name) {
    return glc->fnBindAttribLocation(program, index, name);
}

void gl20BlendEquationSeperate(gl20Context* glc, GLenum modeRGB, GLenum modeAlpha) {
    return glc->fnBlendEquationSeperate(modeRGB, modeAlpha);
}

void gl20CompileShader(gl20Context* glc, GLuint shader) {
    return glc->fnCompileShader(shader);
}

GLuint gl20CreateProgram(gl20Context* glc) {
    return glc->fnCreateProgram();
}

GLuint gl20CreateShader(gl20Context* glc, GLenum shaderType) {
    return glc->fnCreateShader(shaderType);
}

void gl20DeleteProgram(gl20Context* glc, GLuint program) {
    return glc->fnDeleteProgram(program);
}

void gl20DeleteShader(gl20Context* glc, GLuint shader) {
    return glc->fnDeleteShader(shader);
}

void gl20DetachShader(gl20Context* glc, GLuint program, GLuint shader) {
    return glc->fnDetachShader(program, shader);
}

void gl20EnableVertexAttribArray(gl20Context* glc, GLuint index) {
    return glc->fnEnableVertexAttribArray(index);
}

void gl20DisableVertexAttribArray(gl20Context* glc, GLuint index) {
    return glc->fnDisableVertexAttribArray(index);
}

void gl20UniformMatrix2fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2fv(location, count, transpose, value);
}

void gl20UniformMatrix3fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3fv(location, count, transpose, value);
}

void gl20UniformMatrix4fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4fv(location, count, transpose, value);
}

void gl20UniformMatrix2x3fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x3fv(location, count, transpose, value);
}

void gl20UniformMatrix3x2fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x2fv(location, count, transpose, value);
}

void gl20UniformMatrix2x4fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix2x4fv(location, count, transpose, value);
}

void gl20UniformMatrix4x2fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x2fv(location, count, transpose, value);
}

void gl20UniformMatrix3x4fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix3x4fv(location, count, transpose, value);
}

void gl20UniformMatrix4x3fv(gl20Context* glc, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {
    return glc->fnUniformMatrix4x3fv(location, count, transpose, value);
}

gl20Context* gl20NewContext() {
    gl20Context* glc = calloc(1, sizeof(gl20Context));

    // Preload all procedures
    glc->fnAccum = (gl20PAccum)gl20LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl20PAlphaFunc)gl20LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl20PBegin)gl20LibGetProcAddress("glBegin");
    glc->fnEnd = (gl20PEnd)gl20LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl20PBitmap)gl20LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl20PBlendFunc)gl20LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl20PCallList)gl20LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl20PCallLists)gl20LibGetProcAddress("glCallLists");
    glc->fnClear = (gl20PClear)gl20LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl20PClearAccum)gl20LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl20PClearColor)gl20LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl20PClearDepth)gl20LibGetProcAddress("glClearDepth");
    glc->fnClearDepthf = (gl20PClearDepthf)gl20LibGetProcAddress("glClearDepthf");
    glc->fnClearIndex = (gl20PClearIndex)gl20LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl20PClearStencil)gl20LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl20PClipPlane)gl20LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl20PColor3b)gl20LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl20PColor3d)gl20LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl20PColor3f)gl20LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl20PColor3i)gl20LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl20PColor3s)gl20LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl20PColor3ub)gl20LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl20PColor3ui)gl20LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl20PColor3us)gl20LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl20PColor4b)gl20LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl20PColor4d)gl20LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl20PColor4f)gl20LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl20PColor4i)gl20LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl20PColor4s)gl20LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl20PColor4ub)gl20LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl20PColor4ui)gl20LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl20PColor4us)gl20LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl20PColor3bv)gl20LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl20PColor3dv)gl20LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl20PColor3fv)gl20LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl20PColor3iv)gl20LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl20PColor3sv)gl20LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl20PColor3ubv)gl20LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl20PColor3uiv)gl20LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl20PColor3usv)gl20LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl20PColor4bv)gl20LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl20PColor4dv)gl20LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl20PColor4fv)gl20LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl20PColor4iv)gl20LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl20PColor4sv)gl20LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl20PColor4ubv)gl20LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl20PColor4uiv)gl20LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl20PColor4usv)gl20LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl20PColorMask)gl20LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl20PColorMaterial)gl20LibGetProcAddress("glColorMaterial");
    glc->fnCopyPixels = (gl20PCopyPixels)gl20LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl20PCullFace)gl20LibGetProcAddress("glCullFace");
    glc->fnDeleteLists = (gl20PDeleteLists)gl20LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl20PDepthFunc)gl20LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl20PDepthMask)gl20LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl20PDepthRange)gl20LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl20PEnable)gl20LibGetProcAddress("glEnable");
    glc->fnDisable = (gl20PDisable)gl20LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl20PDrawBuffer)gl20LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl20PDrawPixels)gl20LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl20PEdgeFlag)gl20LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl20PEdgeFlagv)gl20LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl20PEdgeFlagPointer)gl20LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl20PEvalCoord1d)gl20LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl20PEvalCoord1f)gl20LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl20PEvalCoord2d)gl20LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl20PEvalCoord2f)gl20LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl20PEvalCoord1dv)gl20LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl20PEvalCoord1fv)gl20LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl20PEvalCoord2dv)gl20LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl20PEvalCoord2fv)gl20LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl20PEvalMesh1)gl20LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl20PEvalMesh2)gl20LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl20PEvalPoint1)gl20LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl20PEvalPoint2)gl20LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl20PFeedbackBuffer)gl20LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl20PFinish)gl20LibGetProcAddress("glFinish");
    glc->fnFlush = (gl20PFlush)gl20LibGetProcAddress("glFlush");
    glc->fnFogf = (gl20PFogf)gl20LibGetProcAddress("glFogf");
    glc->fnFogi = (gl20PFogi)gl20LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl20PFogfv)gl20LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl20PFogiv)gl20LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl20PFrontFace)gl20LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl20PFrustum)gl20LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl20PGenLists)gl20LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl20PGetBooleanv)gl20LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl20PGetDoublev)gl20LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl20PGetFloatv)gl20LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl20PGetIntegerv)gl20LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl20PGetClipPlane)gl20LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl20PGetError)gl20LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl20PGetLightfv)gl20LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl20PGetLightiv)gl20LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl20PGetMapdv)gl20LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl20PGetMapfv)gl20LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl20PGetMapiv)gl20LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl20PGetMaterialfv)gl20LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl20PGetMaterialiv)gl20LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl20PGetPixelMapfv)gl20LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl20PGetPixelMapuiv)gl20LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl20PGetPixelMapusv)gl20LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl20PGetPolygonStipple)gl20LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl20PGetString)gl20LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl20PGetTexEnvfv)gl20LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl20PGetTexEnviv)gl20LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl20PGetTexGendv)gl20LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl20PGetTexGenfv)gl20LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl20PGetTexGeniv)gl20LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl20PGetTexImage)gl20LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl20PGetTexLevelParameterfv)gl20LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl20PGetTexLevelParameteriv)gl20LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl20PGetTexParameterfv)gl20LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl20PGetTexParameteriv)gl20LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl20PHint)gl20LibGetProcAddress("glHint");
    glc->fnIndexd = (gl20PIndexd)gl20LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl20PIndexf)gl20LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl20PIndexi)gl20LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl20PIndexs)gl20LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl20PIndexdv)gl20LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl20PIndexfv)gl20LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl20PIndexiv)gl20LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl20PIndexsv)gl20LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl20PIndexMask)gl20LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl20PIndexPointer)gl20LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl20PInitNames)gl20LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl20PIsEnabled)gl20LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl20PIsList)gl20LibGetProcAddress("glIsList");
    glc->fnLightf = (gl20PLightf)gl20LibGetProcAddress("glLightf");
    glc->fnLighti = (gl20PLighti)gl20LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl20PLightfv)gl20LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl20PLightiv)gl20LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl20PLightModelf)gl20LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl20PLightModeli)gl20LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl20PLightModelfv)gl20LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl20PLightModeliv)gl20LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl20PLineStipple)gl20LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl20PLineWidth)gl20LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl20PListBase)gl20LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl20PLoadIdentity)gl20LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl20PLoadMatrixd)gl20LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl20PLoadMatrixf)gl20LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl20PLoadName)gl20LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl20PLogicOp)gl20LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl20PMap1d)gl20LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl20PMap1f)gl20LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl20PMap2d)gl20LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl20PMap2f)gl20LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl20PMapGrid1d)gl20LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl20PMapGrid1f)gl20LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl20PMapGrid2d)gl20LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl20PMapGrid2f)gl20LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl20PMaterialf)gl20LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl20PMateriali)gl20LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl20PMaterialfv)gl20LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl20PMaterialiv)gl20LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl20PMatrixMode)gl20LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl20PMultMatrixd)gl20LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl20PMultMatrixf)gl20LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl20PNewList)gl20LibGetProcAddress("glNewList");
    glc->fnEndList = (gl20PEndList)gl20LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl20PNormal3b)gl20LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl20PNormal3d)gl20LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl20PNormal3f)gl20LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl20PNormal3i)gl20LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl20PNormal3s)gl20LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl20PNormal3bv)gl20LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl20PNormal3dv)gl20LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl20PNormal3fv)gl20LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl20PNormal3iv)gl20LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl20PNormal3sv)gl20LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl20POrtho)gl20LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl20PPassThrough)gl20LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl20PPixelMapfv)gl20LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl20PPixelMapuiv)gl20LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl20PPixelMapusv)gl20LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl20PPixelStoref)gl20LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl20PPixelStorei)gl20LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl20PPixelTransferf)gl20LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl20PPixelTransferi)gl20LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl20PPixelZoom)gl20LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl20PPointSize)gl20LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl20PPolygonMode)gl20LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl20PPolygonStipple)gl20LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl20PPushAttrib)gl20LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl20PPopAttrib)gl20LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl20PPushMatrix)gl20LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl20PPopMatrix)gl20LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl20PPushName)gl20LibGetProcAddress("glPushName");
    glc->fnPopName = (gl20PPopName)gl20LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl20PRasterPos2d)gl20LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl20PRasterPos2f)gl20LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl20PRasterPos2i)gl20LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl20PRasterPos2s)gl20LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl20PRasterPos3d)gl20LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl20PRasterPos3f)gl20LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl20PRasterPos3i)gl20LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl20PRasterPos3s)gl20LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl20PRasterPos4d)gl20LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl20PRasterPos4f)gl20LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl20PRasterPos4i)gl20LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl20PRasterPos4s)gl20LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl20PRasterPos2dv)gl20LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl20PRasterPos2fv)gl20LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl20PRasterPos2iv)gl20LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl20PRasterPos2sv)gl20LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl20PRasterPos3dv)gl20LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl20PRasterPos3fv)gl20LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl20PRasterPos3iv)gl20LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl20PRasterPos3sv)gl20LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl20PRasterPos4dv)gl20LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl20PRasterPos4fv)gl20LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl20PRasterPos4iv)gl20LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl20PRasterPos4sv)gl20LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl20PReadBuffer)gl20LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl20PReadPixels)gl20LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl20PRectd)gl20LibGetProcAddress("glRectd");
    glc->fnRectf = (gl20PRectf)gl20LibGetProcAddress("glRectf");
    glc->fnRecti = (gl20PRecti)gl20LibGetProcAddress("glRecti");
    glc->fnRects = (gl20PRects)gl20LibGetProcAddress("glRects");
    glc->fnRectdv = (gl20PRectdv)gl20LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl20PRectfv)gl20LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl20PRectiv)gl20LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl20PRectsv)gl20LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl20PRenderMode)gl20LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl20PRotated)gl20LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl20PRotatef)gl20LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl20PScaled)gl20LibGetProcAddress("glScaled");
    glc->fnScalef = (gl20PScalef)gl20LibGetProcAddress("glScalef");
    glc->fnScissor = (gl20PScissor)gl20LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl20PSelectBuffer)gl20LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl20PShadeModel)gl20LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl20PStencilFunc)gl20LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl20PStencilMask)gl20LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl20PStencilOp)gl20LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl20PTexCoord1d)gl20LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl20PTexCoord1f)gl20LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl20PTexCoord1i)gl20LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl20PTexCoord1s)gl20LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl20PTexCoord2d)gl20LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl20PTexCoord2f)gl20LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl20PTexCoord2i)gl20LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl20PTexCoord2s)gl20LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl20PTexCoord3d)gl20LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl20PTexCoord3f)gl20LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl20PTexCoord3i)gl20LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl20PTexCoord3s)gl20LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl20PTexCoord4d)gl20LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl20PTexCoord4f)gl20LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl20PTexCoord4i)gl20LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl20PTexCoord4s)gl20LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl20PTexCoord1dv)gl20LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl20PTexCoord1fv)gl20LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl20PTexCoord1iv)gl20LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl20PTexCoord1sv)gl20LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl20PTexCoord2dv)gl20LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl20PTexCoord2fv)gl20LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl20PTexCoord2iv)gl20LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl20PTexCoord2sv)gl20LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl20PTexCoord3dv)gl20LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl20PTexCoord3fv)gl20LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl20PTexCoord3iv)gl20LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl20PTexCoord3sv)gl20LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl20PTexCoord4dv)gl20LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl20PTexCoord4fv)gl20LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl20PTexCoord4iv)gl20LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl20PTexCoord4sv)gl20LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl20PTexEnvf)gl20LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl20PTexEnvi)gl20LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl20PTexEnvfv)gl20LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl20PTexEnviv)gl20LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl20PTexGend)gl20LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl20PTexGenf)gl20LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl20PTexGeni)gl20LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl20PTexGendv)gl20LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl20PTexGenfv)gl20LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl20PTexGeniv)gl20LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl20PTexImage1D)gl20LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl20PTexImage2D)gl20LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl20PTexParameterf)gl20LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl20PTexParameteri)gl20LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl20PTexParameterfv)gl20LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl20PTexParameteriv)gl20LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl20PTranslated)gl20LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl20PTranslatef)gl20LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl20PVertex2s)gl20LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl20PVertex2i)gl20LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl20PVertex2f)gl20LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl20PVertex2d)gl20LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl20PVertex3s)gl20LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl20PVertex3i)gl20LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl20PVertex3f)gl20LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl20PVertex3d)gl20LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl20PVertex4s)gl20LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl20PVertex4i)gl20LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl20PVertex4f)gl20LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl20PVertex4d)gl20LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl20PViewport)gl20LibGetProcAddress("glViewport");
    glc->fnGetConvolutionParameterfv = (gl20PGetConvolutionParameterfv)gl20LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl20PGetConvolutionParameteriv)gl20LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnAreTexturesResident = (gl20PAreTexturesResident)gl20LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl20PArrayElement)gl20LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl20PDrawArrays)gl20LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl20PDrawElements)gl20LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl20PGetPointerv)gl20LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl20PPolygonOffset)gl20LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl20PCopyTexImage1D)gl20LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl20PCopyTexImage2D)gl20LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl20PCopyTexSubImage1D)gl20LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl20PCopyTexSubImage2D)gl20LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl20PBindTexture)gl20LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl20PDeleteTextures)gl20LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl20PGenTextures)gl20LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl20PIsTexture)gl20LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl20PColorPointer)gl20LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl20PEnableClientState)gl20LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl20PDisableClientState)gl20LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl20PIndexub)gl20LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl20PIndexubv)gl20LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl20PInterleavedArrays)gl20LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl20PNormalPointer)gl20LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl20PPushClientAttrib)gl20LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl20PPrioritizeTextures)gl20LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl20PPopClientAttrib)gl20LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl20PTexCoordPointer)gl20LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl20PTexSubImage1D)gl20LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl20PTexSubImage2D)gl20LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl20PVertexPointer)gl20LibGetProcAddress("glVertexPointer");
    glc->fnDrawBuffers = (gl20PDrawBuffers)gl20LibGetProcAddress("glDrawBuffers");
    glc->fnGetActiveAttrib = (gl20PGetActiveAttrib)gl20LibGetProcAddress("glGetActiveAttrib");
    glc->fnGetActiveUniform = (gl20PGetActiveUniform)gl20LibGetProcAddress("glGetActiveUniform");
    glc->fnGetAttachedShaders = (gl20PGetAttachedShaders)gl20LibGetProcAddress("glGetAttachedShaders");
    glc->fnGetAttribLocation = (gl20PGetAttribLocation)gl20LibGetProcAddress("glGetAttribLocation");
    glc->fnGetProgramiv = (gl20PGetProgramiv)gl20LibGetProcAddress("glGetProgramiv");
    glc->fnGetProgramInfoLog = (gl20PGetProgramInfoLog)gl20LibGetProcAddress("glGetProgramInfoLog");
    glc->fnGetShaderiv = (gl20PGetShaderiv)gl20LibGetProcAddress("glGetShaderiv");
    glc->fnGetShaderInfoLog = (gl20PGetShaderInfoLog)gl20LibGetProcAddress("glGetShaderInfoLog");
    glc->fnGetShaderSource = (gl20PGetShaderSource)gl20LibGetProcAddress("glGetShaderSource");
    glc->fnGetUniformfv = (gl20PGetUniformfv)gl20LibGetProcAddress("glGetUniformfv");
    glc->fnGetUniformiv = (gl20PGetUniformiv)gl20LibGetProcAddress("glGetUniformiv");
    glc->fnGetUniformLocation = (gl20PGetUniformLocation)gl20LibGetProcAddress("glGetUniformLocation");
    glc->fnGetVertexAttribdv = (gl20PGetVertexAttribdv)gl20LibGetProcAddress("glGetVertexAttribdv");
    glc->fnGetVertexAttribfv = (gl20PGetVertexAttribfv)gl20LibGetProcAddress("glGetVertexAttribfv");
    glc->fnGetVertexAttribiv = (gl20PGetVertexAttribiv)gl20LibGetProcAddress("glGetVertexAttribiv");
    glc->fnGetVertexAttribPointerv = (gl20PGetVertexAttribPointerv)gl20LibGetProcAddress("glGetVertexAttribPointerv");
    glc->fnIsProgram = (gl20PIsProgram)gl20LibGetProcAddress("glIsProgram");
    glc->fnIsShader = (gl20PIsShader)gl20LibGetProcAddress("glIsShader");
    glc->fnLinkProgram = (gl20PLinkProgram)gl20LibGetProcAddress("glLinkProgram");
    glc->fnShaderSource = (gl20PShaderSource)gl20LibGetProcAddress("glShaderSource");
    glc->fnStencilFuncSeparate = (gl20PStencilFuncSeparate)gl20LibGetProcAddress("glStencilFuncSeparate");
    glc->fnStencilMaskSeparate = (gl20PStencilMaskSeparate)gl20LibGetProcAddress("glStencilMaskSeparate");
    glc->fnStencilOpSeparate = (gl20PStencilOpSeparate)gl20LibGetProcAddress("glStencilOpSeparate");
    glc->fnUniform1f = (gl20PUniform1f)gl20LibGetProcAddress("glUniform1f");
    glc->fnUniform2f = (gl20PUniform2f)gl20LibGetProcAddress("glUniform2f");
    glc->fnUniform3f = (gl20PUniform3f)gl20LibGetProcAddress("glUniform3f");
    glc->fnUniform4f = (gl20PUniform4f)gl20LibGetProcAddress("glUniform4f");
    glc->fnUniform1i = (gl20PUniform1i)gl20LibGetProcAddress("glUniform1i");
    glc->fnUniform2i = (gl20PUniform2i)gl20LibGetProcAddress("glUniform2i");
    glc->fnUniform3i = (gl20PUniform3i)gl20LibGetProcAddress("glUniform3i");
    glc->fnUniform4i = (gl20PUniform4i)gl20LibGetProcAddress("glUniform4i");
    glc->fnUniform1fv = (gl20PUniform1fv)gl20LibGetProcAddress("glUniform1fv");
    glc->fnUniform2fv = (gl20PUniform2fv)gl20LibGetProcAddress("glUniform2fv");
    glc->fnUniform3fv = (gl20PUniform3fv)gl20LibGetProcAddress("glUniform3fv");
    glc->fnUniform4fv = (gl20PUniform4fv)gl20LibGetProcAddress("glUniform4fv");
    glc->fnUniform1iv = (gl20PUniform1iv)gl20LibGetProcAddress("glUniform1iv");
    glc->fnUniform2iv = (gl20PUniform2iv)gl20LibGetProcAddress("glUniform2iv");
    glc->fnUniform3iv = (gl20PUniform3iv)gl20LibGetProcAddress("glUniform3iv");
    glc->fnUniform4iv = (gl20PUniform4iv)gl20LibGetProcAddress("glUniform4iv");
    glc->fnUseProgram = (gl20PUseProgram)gl20LibGetProcAddress("glUseProgram");
    glc->fnValidateProgram = (gl20PValidateProgram)gl20LibGetProcAddress("glValidateProgram");
    glc->fnVertexAttribPointer = (gl20PVertexAttribPointer)gl20LibGetProcAddress("glVertexAttribPointer");
    glc->fnVertexAttrib1f = (gl20PVertexAttrib1f)gl20LibGetProcAddress("glVertexAttrib1f");
    glc->fnVertexAttrib1s = (gl20PVertexAttrib1s)gl20LibGetProcAddress("glVertexAttrib1s");
    glc->fnVertexAttrib1d = (gl20PVertexAttrib1d)gl20LibGetProcAddress("glVertexAttrib1d");
    glc->fnVertexAttrib2f = (gl20PVertexAttrib2f)gl20LibGetProcAddress("glVertexAttrib2f");
    glc->fnVertexAttrib2s = (gl20PVertexAttrib2s)gl20LibGetProcAddress("glVertexAttrib2s");
    glc->fnVertexAttrib2d = (gl20PVertexAttrib2d)gl20LibGetProcAddress("glVertexAttrib2d");
    glc->fnVertexAttrib3f = (gl20PVertexAttrib3f)gl20LibGetProcAddress("glVertexAttrib3f");
    glc->fnVertexAttrib3s = (gl20PVertexAttrib3s)gl20LibGetProcAddress("glVertexAttrib3s");
    glc->fnVertexAttrib3d = (gl20PVertexAttrib3d)gl20LibGetProcAddress("glVertexAttrib3d");
    glc->fnVertexAttrib4f = (gl20PVertexAttrib4f)gl20LibGetProcAddress("glVertexAttrib4f");
    glc->fnVertexAttrib4s = (gl20PVertexAttrib4s)gl20LibGetProcAddress("glVertexAttrib4s");
    glc->fnVertexAttrib4d = (gl20PVertexAttrib4d)gl20LibGetProcAddress("glVertexAttrib4d");
    glc->fnVertexAttrib4Nuv = (gl20PVertexAttrib4Nuv)gl20LibGetProcAddress("glVertexAttrib4Nuv");
    glc->fnVertexAttrib1fv = (gl20PVertexAttrib1fv)gl20LibGetProcAddress("glVertexAttrib1fv");
    glc->fnVertexAttrib1sv = (gl20PVertexAttrib1sv)gl20LibGetProcAddress("glVertexAttrib1sv");
    glc->fnVertexAttrib1dv = (gl20PVertexAttrib1dv)gl20LibGetProcAddress("glVertexAttrib1dv");
    glc->fnVertexAttrib2fv = (gl20PVertexAttrib2fv)gl20LibGetProcAddress("glVertexAttrib2fv");
    glc->fnVertexAttrib2sv = (gl20PVertexAttrib2sv)gl20LibGetProcAddress("glVertexAttrib2sv");
    glc->fnVertexAttrib2dv = (gl20PVertexAttrib2dv)gl20LibGetProcAddress("glVertexAttrib2dv");
    glc->fnVertexAttrib3fv = (gl20PVertexAttrib3fv)gl20LibGetProcAddress("glVertexAttrib3fv");
    glc->fnVertexAttrib3sv = (gl20PVertexAttrib3sv)gl20LibGetProcAddress("glVertexAttrib3sv");
    glc->fnVertexAttrib3dv = (gl20PVertexAttrib3dv)gl20LibGetProcAddress("glVertexAttrib3dv");
    glc->fnVertexAttrib4fv = (gl20PVertexAttrib4fv)gl20LibGetProcAddress("glVertexAttrib4fv");
    glc->fnVertexAttrib4sv = (gl20PVertexAttrib4sv)gl20LibGetProcAddress("glVertexAttrib4sv");
    glc->fnVertexAttrib4dv = (gl20PVertexAttrib4dv)gl20LibGetProcAddress("glVertexAttrib4dv");
    glc->fnVertexAttrib4iv = (gl20PVertexAttrib4iv)gl20LibGetProcAddress("glVertexAttrib4iv");
    glc->fnVertexAttrib4bv = (gl20PVertexAttrib4bv)gl20LibGetProcAddress("glVertexAttrib4bv");
    glc->fnVertexAttrib4ubv = (gl20PVertexAttrib4ubv)gl20LibGetProcAddress("glVertexAttrib4ubv");
    glc->fnVertexAttrib4usv = (gl20PVertexAttrib4usv)gl20LibGetProcAddress("glVertexAttrib4usv");
    glc->fnVertexAttrib4uiv = (gl20PVertexAttrib4uiv)gl20LibGetProcAddress("glVertexAttrib4uiv");
    glc->fnVertexAttrib4Nbv = (gl20PVertexAttrib4Nbv)gl20LibGetProcAddress("glVertexAttrib4Nbv");
    glc->fnVertexAttrib4Nsv = (gl20PVertexAttrib4Nsv)gl20LibGetProcAddress("glVertexAttrib4Nsv");
    glc->fnVertexAttrib4Niv = (gl20PVertexAttrib4Niv)gl20LibGetProcAddress("glVertexAttrib4Niv");
    glc->fnVertexAttrib4Nubv = (gl20PVertexAttrib4Nubv)gl20LibGetProcAddress("glVertexAttrib4Nubv");
    glc->fnVertexAttrib4Nusv = (gl20PVertexAttrib4Nusv)gl20LibGetProcAddress("glVertexAttrib4Nusv");
    glc->fnVertexAttrib4Nuiv = (gl20PVertexAttrib4Nuiv)gl20LibGetProcAddress("glVertexAttrib4Nuiv");
    glc->fnColorTable = (gl20PColorTable)gl20GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl20PColorTableParameterfv)gl20GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl20PColorTableParameteriv)gl20GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl20PColorSubTable)gl20GLGetProcAddress("glColorSubTable");
    glc->fnConvolutionFilter1D = (gl20PConvolutionFilter1D)gl20GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl20PConvolutionFilter2D)gl20GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl20PConvolutionParameterf)gl20GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl20PConvolutionParameteri)gl20GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl20PCopyColorTable)gl20GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl20PCopyColorSubTable)gl20GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl20PCopyConvolutionFilter1D)gl20GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl20PCopyConvolutionFilter2D)gl20GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnGetColorTable = (gl20PGetColorTable)gl20GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl20PGetColorTableParameterfv)gl20GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl20PGetColorTableParameteriv)gl20GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl20PGetConvolutionFilter)gl20GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetHistogram = (gl20PGetHistogram)gl20GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl20PGetHistogramParameterfv)gl20GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl20PGetHistogramParameteriv)gl20GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl20PGetSeparableFilter)gl20GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl20PHistogram)gl20GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl20PMinmax)gl20GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl20PMultiTexCoord1s)gl20GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl20PMultiTexCoord1i)gl20GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl20PMultiTexCoord1f)gl20GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl20PMultiTexCoord1d)gl20GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl20PMultiTexCoord2s)gl20GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl20PMultiTexCoord2i)gl20GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl20PMultiTexCoord2f)gl20GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl20PMultiTexCoord2d)gl20GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl20PMultiTexCoord3s)gl20GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl20PMultiTexCoord3i)gl20GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl20PMultiTexCoord3f)gl20GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl20PMultiTexCoord3d)gl20GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl20PMultiTexCoord4s)gl20GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl20PMultiTexCoord4i)gl20GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl20PMultiTexCoord4f)gl20GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl20PMultiTexCoord4d)gl20GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl20PMultiTexCoord1sv)gl20GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl20PMultiTexCoord1iv)gl20GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl20PMultiTexCoord1fv)gl20GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl20PMultiTexCoord1dv)gl20GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl20PMultiTexCoord2sv)gl20GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl20PMultiTexCoord2iv)gl20GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl20PMultiTexCoord2fv)gl20GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl20PMultiTexCoord2dv)gl20GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl20PMultiTexCoord3sv)gl20GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl20PMultiTexCoord3iv)gl20GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl20PMultiTexCoord3fv)gl20GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl20PMultiTexCoord3dv)gl20GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl20PMultiTexCoord4sv)gl20GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl20PMultiTexCoord4iv)gl20GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl20PMultiTexCoord4fv)gl20GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl20PMultiTexCoord4dv)gl20GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl20PResetHistogram)gl20GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl20PResetMinmax)gl20GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl20PSeparableFilter2D)gl20GLGetProcAddress("glSeparableFilter2D");
    glc->fnBlendColor = (gl20PBlendColor)gl20GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl20PBlendEquation)gl20GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl20PCopyTexSubImage3D)gl20GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl20PDrawRangeElements)gl20GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl20PTexImage3D)gl20GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl20PTexSubImage3D)gl20GLGetProcAddress("glTexSubImage3D");
    glc->fnActiveTexture = (gl20PActiveTexture)gl20GLGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl20PClientActiveTexture)gl20GLGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl20PCompressedTexImage1D)gl20GLGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl20PCompressedTexImage2D)gl20GLGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl20PCompressedTexImage3D)gl20GLGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl20PCompressedTexSubImage1D)gl20GLGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl20PCompressedTexSubImage2D)gl20GLGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl20PCompressedTexSubImage3D)gl20GLGetProcAddress("glCompressedTexSubImage3D");
    glc->fnGetCompressedTexImage = (gl20PGetCompressedTexImage)gl20GLGetProcAddress("glGetCompressedTexImage");
    glc->fnLoadTransposeMatrixd = (gl20PLoadTransposeMatrixd)gl20GLGetProcAddress("glLoadTransposeMatrixd");
    glc->fnLoadTransposeMatrixf = (gl20PLoadTransposeMatrixf)gl20GLGetProcAddress("glLoadTransposeMatrixf");
    glc->fnMultTransposeMatrixd = (gl20PMultTransposeMatrixd)gl20GLGetProcAddress("glMultTransposeMatrixd");
    glc->fnMultTransposeMatrixf = (gl20PMultTransposeMatrixf)gl20GLGetProcAddress("glMultTransposeMatrixf");
    glc->fnSampleCoverage = (gl20PSampleCoverage)gl20GLGetProcAddress("glSampleCoverage");
    glc->fnBlendFuncSeparate = (gl20PBlendFuncSeparate)gl20GLGetProcAddress("glBlendFuncSeparate");
    glc->fnFogCoordPointer = (gl20PFogCoordPointer)gl20GLGetProcAddress("glFogCoordPointer");
    glc->fnFogCoordd = (gl20PFogCoordd)gl20GLGetProcAddress("glFogCoordd");
    glc->fnFogCoordf = (gl20PFogCoordf)gl20GLGetProcAddress("glFogCoordf");
    glc->fnFogCoorddv = (gl20PFogCoorddv)gl20GLGetProcAddress("glFogCoorddv");
    glc->fnFogCoordfv = (gl20PFogCoordfv)gl20GLGetProcAddress("glFogCoordfv");
    glc->fnMultiDrawArrays = (gl20PMultiDrawArrays)gl20GLGetProcAddress("glMultiDrawArrays");
    glc->fnMultiDrawElements = (gl20PMultiDrawElements)gl20GLGetProcAddress("glMultiDrawElements");
    glc->fnPointParameterf = (gl20PPointParameterf)gl20GLGetProcAddress("glPointParameterf");
    glc->fnPointParameteri = (gl20PPointParameteri)gl20GLGetProcAddress("glPointParameteri");
    glc->fnSecondaryColor3b = (gl20PSecondaryColor3b)gl20GLGetProcAddress("glSecondaryColor3b");
    glc->fnSecondaryColor3s = (gl20PSecondaryColor3s)gl20GLGetProcAddress("glSecondaryColor3s");
    glc->fnSecondaryColor3i = (gl20PSecondaryColor3i)gl20GLGetProcAddress("glSecondaryColor3i");
    glc->fnSecondaryColor3f = (gl20PSecondaryColor3f)gl20GLGetProcAddress("glSecondaryColor3f");
    glc->fnSecondaryColor3d = (gl20PSecondaryColor3d)gl20GLGetProcAddress("glSecondaryColor3d");
    glc->fnSecondaryColor3ub = (gl20PSecondaryColor3ub)gl20GLGetProcAddress("glSecondaryColor3ub");
    glc->fnSecondaryColor3us = (gl20PSecondaryColor3us)gl20GLGetProcAddress("glSecondaryColor3us");
    glc->fnSecondaryColor3ui = (gl20PSecondaryColor3ui)gl20GLGetProcAddress("glSecondaryColor3ui");
    glc->fnSecondaryColor3bv = (gl20PSecondaryColor3bv)gl20GLGetProcAddress("glSecondaryColor3bv");
    glc->fnSecondaryColor3sv = (gl20PSecondaryColor3sv)gl20GLGetProcAddress("glSecondaryColor3sv");
    glc->fnSecondaryColor3iv = (gl20PSecondaryColor3iv)gl20GLGetProcAddress("glSecondaryColor3iv");
    glc->fnSecondaryColor3fv = (gl20PSecondaryColor3fv)gl20GLGetProcAddress("glSecondaryColor3fv");
    glc->fnSecondaryColor3dv = (gl20PSecondaryColor3dv)gl20GLGetProcAddress("glSecondaryColor3dv");
    glc->fnSecondaryColor3ubv = (gl20PSecondaryColor3ubv)gl20GLGetProcAddress("glSecondaryColor3ubv");
    glc->fnSecondaryColor3usv = (gl20PSecondaryColor3usv)gl20GLGetProcAddress("glSecondaryColor3usv");
    glc->fnSecondaryColor3uiv = (gl20PSecondaryColor3uiv)gl20GLGetProcAddress("glSecondaryColor3uiv");
    glc->fnSecondaryColorPointer = (gl20PSecondaryColorPointer)gl20GLGetProcAddress("glSecondaryColorPointer");
    glc->fnWindowPos2s = (gl20PWindowPos2s)gl20GLGetProcAddress("glWindowPos2s");
    glc->fnWindowPos2i = (gl20PWindowPos2i)gl20GLGetProcAddress("glWindowPos2i");
    glc->fnWindowPos2f = (gl20PWindowPos2f)gl20GLGetProcAddress("glWindowPos2f");
    glc->fnWindowPos2d = (gl20PWindowPos2d)gl20GLGetProcAddress("glWindowPos2d");
    glc->fnWindowPos3s = (gl20PWindowPos3s)gl20GLGetProcAddress("glWindowPos3s");
    glc->fnWindowPos3i = (gl20PWindowPos3i)gl20GLGetProcAddress("glWindowPos3i");
    glc->fnWindowPos3f = (gl20PWindowPos3f)gl20GLGetProcAddress("glWindowPos3f");
    glc->fnWindowPos3d = (gl20PWindowPos3d)gl20GLGetProcAddress("glWindowPos3d");
    glc->fnWindowPos2sv = (gl20PWindowPos2sv)gl20GLGetProcAddress("glWindowPos2sv");
    glc->fnWindowPos2iv = (gl20PWindowPos2iv)gl20GLGetProcAddress("glWindowPos2iv");
    glc->fnWindowPos2fv = (gl20PWindowPos2fv)gl20GLGetProcAddress("glWindowPos2fv");
    glc->fnWindowPos2dv = (gl20PWindowPos2dv)gl20GLGetProcAddress("glWindowPos2dv");
    glc->fnWindowPos3sv = (gl20PWindowPos3sv)gl20GLGetProcAddress("glWindowPos3sv");
    glc->fnWindowPos3iv = (gl20PWindowPos3iv)gl20GLGetProcAddress("glWindowPos3iv");
    glc->fnWindowPos3fv = (gl20PWindowPos3fv)gl20GLGetProcAddress("glWindowPos3fv");
    glc->fnWindowPos3dv = (gl20PWindowPos3dv)gl20GLGetProcAddress("glWindowPos3dv");
    glc->fnBeginQuery = (gl20PBeginQuery)gl20GLGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl20PBindBuffer)gl20GLGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl20PBufferData)gl20GLGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl20PBufferSubData)gl20GLGetProcAddress("glBufferSubData");
    glc->fnDeleteBuffers = (gl20PDeleteBuffers)gl20GLGetProcAddress("glDeleteBuffers");
    glc->fnDeleteQueries = (gl20PDeleteQueries)gl20GLGetProcAddress("glDeleteQueries");
    glc->fnGenBuffers = (gl20PGenBuffers)gl20GLGetProcAddress("glGenBuffers");
    glc->fnGenQueries = (gl20PGenQueries)gl20GLGetProcAddress("glGenQueries");
    glc->fnGetBufferParameteriv = (gl20PGetBufferParameteriv)gl20GLGetProcAddress("glGetBufferParameteriv");
    glc->fnGetBufferPointerv = (gl20PGetBufferPointerv)gl20GLGetProcAddress("glGetBufferPointerv");
    glc->fnGetBufferSubData = (gl20PGetBufferSubData)gl20GLGetProcAddress("glGetBufferSubData");
    glc->fnGetQueryObjectiv = (gl20PGetQueryObjectiv)gl20GLGetProcAddress("glGetQueryObjectiv");
    glc->fnGetQueryObjectuiv = (gl20PGetQueryObjectuiv)gl20GLGetProcAddress("glGetQueryObjectuiv");
    glc->fnGetQueryiv = (gl20PGetQueryiv)gl20GLGetProcAddress("glGetQueryiv");
    glc->fnIsBuffer = (gl20PIsBuffer)gl20GLGetProcAddress("glIsBuffer");
    glc->fnIsQuery = (gl20PIsQuery)gl20GLGetProcAddress("glIsQuery");
    glc->fnMapBuffer = (gl20PMapBuffer)gl20GLGetProcAddress("glMapBuffer");
    glc->fnUnmapBuffer = (gl20PUnmapBuffer)gl20GLGetProcAddress("glUnmapBuffer");
    glc->fnAttachShader = (gl20PAttachShader)gl20GLGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl20PBindAttribLocation)gl20GLGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl20PBlendEquationSeperate)gl20GLGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl20PCompileShader)gl20GLGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl20PCreateProgram)gl20GLGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl20PCreateShader)gl20GLGetProcAddress("glCreateShader");
    glc->fnDeleteProgram = (gl20PDeleteProgram)gl20GLGetProcAddress("glDeleteProgram");
    glc->fnDeleteShader = (gl20PDeleteShader)gl20GLGetProcAddress("glDeleteShader");
    glc->fnDetachShader = (gl20PDetachShader)gl20GLGetProcAddress("glDetachShader");
    glc->fnEnableVertexAttribArray = (gl20PEnableVertexAttribArray)gl20GLGetProcAddress("glEnableVertexAttribArray");
    glc->fnDisableVertexAttribArray = (gl20PDisableVertexAttribArray)gl20GLGetProcAddress("glDisableVertexAttribArray");
    glc->fnUniformMatrix2fv = (gl20PUniformMatrix2fv)gl20LibGetProcAddress("glUniformMatrix2fv");
    glc->fnUniformMatrix3fv = (gl20PUniformMatrix3fv)gl20LibGetProcAddress("glUniformMatrix3fv");
    glc->fnUniformMatrix4fv = (gl20PUniformMatrix4fv)gl20LibGetProcAddress("glUniformMatrix4fv");
    glc->fnUniformMatrix2x3fv = (gl20PUniformMatrix2x3fv)gl20LibGetProcAddress("glUniformMatrix2x3fv");
    glc->fnUniformMatrix3x2fv = (gl20PUniformMatrix3x2fv)gl20LibGetProcAddress("glUniformMatrix3x2fv");
    glc->fnUniformMatrix2x4fv = (gl20PUniformMatrix2x4fv)gl20LibGetProcAddress("glUniformMatrix2x4fv");
    glc->fnUniformMatrix4x2fv = (gl20PUniformMatrix4x2fv)gl20LibGetProcAddress("glUniformMatrix4x2fv");
    glc->fnUniformMatrix3x4fv = (gl20PUniformMatrix3x4fv)gl20LibGetProcAddress("glUniformMatrix3x4fv");
    glc->fnUniformMatrix4x3fv = (gl20PUniformMatrix4x3fv)gl20LibGetProcAddress("glUniformMatrix4x3fv");
    return glc;
}

