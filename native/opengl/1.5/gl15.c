#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl15.h"

#ifdef _WIN32
    HMODULE gl15OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl15OpenGL32 == NULL) {
            gl15OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl15OpenGL32, TEXT(name));
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

void gl15CopyPixels(gl15Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl15CullFace(gl15Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl15DeleteLists(gl15Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl15DepthFunc(gl15Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl15DrawPixels(gl15Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl15EnableClientState(gl15Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl15DisableClientState(gl15Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl15GetPolygonStipple(gl15Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl15Indexub(gl15Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl15Indexubv(gl15Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl15InterleavedArrays(gl15Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl15NormalPointer(gl15Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl15PrioritizeTextures(gl15Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl15PushAttrib(gl15Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl15PopAttrib(gl15Context* glc) {
    return glc->fnPopAttrib();
}

void gl15PushClientAttrib(gl15Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl15PopClientAttrib(gl15Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl15TexCoordPointer(gl15Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl15TexImage3DEXT(gl15Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl15TexSubImage1D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl15TexSubImage2D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl15TexSubImage3DEXT(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl15VertexPointer(gl15Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl15Viewport(gl15Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
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

void gl15BlendColorEXT(gl15Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl15BlendEquation(gl15Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl15CopyTexSubImage3D(gl15Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
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

void gl15BlendFuncSeparate(gl15Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
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

gl15Context* gl15NewContext() {
    gl15Context* glc = calloc(1, sizeof(gl15Context));

    // Preload all procedures
    glc->fnAccum = (gl15PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl15PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl15PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl15PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl15PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl15PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl15PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl15PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl15PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl15PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl15PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl15PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl15PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl15PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl15PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl15PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl15PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl15PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl15PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl15PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl15PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl15PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl15PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl15PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl15PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl15PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl15PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl15PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl15PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl15PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl15PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl15PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl15PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl15PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl15PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl15PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl15PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl15PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl15PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl15PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl15PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl15PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl15PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl15PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl15PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl15PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl15PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl15PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl15PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl15PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl15PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl15PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl15PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl15PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl15PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl15PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl15PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl15PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl15PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl15PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl15PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl15PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl15PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl15PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl15PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl15PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl15PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl15PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl15PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl15PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl15PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl15PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl15PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl15PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl15PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl15PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl15PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl15PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl15PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl15PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl15PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl15PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl15PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl15PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl15PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl15PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl15PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl15PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl15PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl15PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl15PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl15PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl15PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl15PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl15PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl15PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl15PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl15PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl15PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl15PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl15PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl15PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl15PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl15PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl15PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl15PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl15PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl15PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl15PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl15PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl15PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl15PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl15PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl15PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl15PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl15PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl15PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl15PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl15PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl15PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl15PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl15PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl15PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl15PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl15PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl15PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl15PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl15PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl15PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl15PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl15PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl15PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl15PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl15PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl15PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl15PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl15PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl15PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl15PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl15PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl15PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl15PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl15PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl15PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl15PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl15PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl15PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl15PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl15PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl15PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl15PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl15PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl15PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl15PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl15PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl15PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl15PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl15PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl15PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl15PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl15PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl15PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl15PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl15PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl15PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl15PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl15PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl15PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl15PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl15PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl15PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl15PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl15PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl15PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl15PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl15PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl15PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl15PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl15PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl15PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl15PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl15PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl15PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl15PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl15PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl15PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl15POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl15PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl15PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl15PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl15PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl15PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl15PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl15PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl15PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl15PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl15PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl15PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl15PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl15PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl15PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl15PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl15PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl15PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl15PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl15PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl15PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl15PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl15PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl15PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl15PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl15PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl15PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl15PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl15PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl15PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl15PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl15PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl15PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl15PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl15PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl15PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl15PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl15PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl15PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl15PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl15PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl15PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl15PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl15PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl15PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl15PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl15PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl15PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl15PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl15PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl15PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl15PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl15PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl15PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl15PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl15PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl15PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl15PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl15PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl15PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl15PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl15PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl15PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl15PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl15PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl15PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl15PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl15PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl15PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl15PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl15PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl15PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl15PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl15PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl15PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl15PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl15PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl15PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl15PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl15PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl15PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl15PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl15PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl15PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl15PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl15PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl15PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl15PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl15PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl15PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl15PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl15PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl15PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl15PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl15PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl15PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl15PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl15PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl15PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl15PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl15PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl15PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl15PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl15PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl15PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl15PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl15PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl15PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl15PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl15PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl15PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl15PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl15PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl15PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl15PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl15PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl15PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl15PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl15PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl15PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl15PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl15PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl15PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl15PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl15PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl15PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl15PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl15PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl15PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl15PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl15PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl15PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl15PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl15PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl15PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl15PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl15PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl15PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl15PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl15PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl15PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl15PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl15PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl15PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl15PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl15PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl15PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl15PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl15PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl15PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl15PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl15PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl15PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl15PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    glc->fnActiveTexture = (gl15PActiveTexture)wglGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl15PClientActiveTexture)wglGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl15PCompressedTexImage1D)wglGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl15PCompressedTexImage2D)wglGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl15PCompressedTexImage3D)wglGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl15PCompressedTexSubImage1D)wglGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl15PCompressedTexSubImage2D)wglGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl15PCompressedTexSubImage3D)wglGetProcAddress("glCompressedTexSubImage3D");
    glc->fnBlendFuncSeparate = (gl15PBlendFuncSeparate)wglGetProcAddress("glBlendFuncSeparate");
    glc->fnBeginQuery = (gl15PBeginQuery)wglGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl15PBindBuffer)wglGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl15PBufferData)wglGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl15PBufferSubData)wglGetProcAddress("glBufferSubData");
    return glc;
}

