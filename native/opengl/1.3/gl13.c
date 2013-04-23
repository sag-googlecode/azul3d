#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl13.h"

#ifdef _WIN32
    HMODULE gl13OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl13OpenGL32 == NULL) {
            gl13OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl13OpenGL32, TEXT(name));
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

void gl13CopyPixels(gl13Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl13CullFace(gl13Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl13DeleteLists(gl13Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl13DepthFunc(gl13Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl13DrawPixels(gl13Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl13EnableClientState(gl13Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl13DisableClientState(gl13Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl13GetPolygonStipple(gl13Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl13Indexub(gl13Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl13Indexubv(gl13Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl13InterleavedArrays(gl13Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl13NormalPointer(gl13Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl13PrioritizeTextures(gl13Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl13PushAttrib(gl13Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl13PopAttrib(gl13Context* glc) {
    return glc->fnPopAttrib();
}

void gl13PushClientAttrib(gl13Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl13PopClientAttrib(gl13Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl13TexCoordPointer(gl13Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl13TexImage3DEXT(gl13Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl13TexSubImage1D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl13TexSubImage2D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl13TexSubImage3DEXT(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl13VertexPointer(gl13Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl13Viewport(gl13Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
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

void gl13BlendColorEXT(gl13Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl13BlendEquation(gl13Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl13CopyTexSubImage3D(gl13Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
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

gl13Context* gl13NewContext() {
    gl13Context* glc = calloc(1, sizeof(gl13Context));

    // Preload all procedures
    glc->fnAccum = (gl13PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl13PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl13PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl13PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl13PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl13PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl13PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl13PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl13PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl13PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl13PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl13PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl13PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl13PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl13PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl13PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl13PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl13PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl13PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl13PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl13PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl13PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl13PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl13PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl13PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl13PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl13PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl13PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl13PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl13PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl13PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl13PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl13PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl13PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl13PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl13PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl13PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl13PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl13PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl13PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl13PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl13PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl13PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl13PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl13PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl13PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl13PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl13PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl13PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl13PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl13PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl13PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl13PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl13PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl13PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl13PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl13PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl13PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl13PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl13PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl13PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl13PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl13PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl13PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl13PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl13PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl13PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl13PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl13PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl13PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl13PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl13PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl13PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl13PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl13PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl13PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl13PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl13PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl13PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl13PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl13PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl13PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl13PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl13PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl13PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl13PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl13PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl13PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl13PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl13PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl13PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl13PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl13PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl13PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl13PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl13PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl13PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl13PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl13PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl13PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl13PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl13PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl13PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl13PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl13PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl13PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl13PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl13PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl13PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl13PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl13PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl13PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl13PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl13PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl13PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl13PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl13PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl13PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl13PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl13PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl13PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl13PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl13PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl13PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl13PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl13PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl13PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl13PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl13PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl13PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl13PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl13PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl13PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl13PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl13PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl13PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl13PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl13PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl13PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl13PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl13PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl13PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl13PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl13PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl13PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl13PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl13PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl13PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl13PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl13PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl13PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl13PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl13PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl13PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl13PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl13PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl13PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl13PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl13PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl13PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl13PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl13PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl13PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl13PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl13PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl13PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl13PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl13PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl13PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl13PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl13PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl13PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl13PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl13PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl13PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl13PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl13PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl13PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl13PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl13PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl13PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl13PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl13PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl13PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl13PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl13PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl13POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl13PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl13PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl13PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl13PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl13PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl13PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl13PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl13PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl13PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl13PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl13PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl13PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl13PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl13PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl13PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl13PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl13PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl13PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl13PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl13PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl13PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl13PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl13PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl13PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl13PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl13PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl13PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl13PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl13PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl13PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl13PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl13PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl13PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl13PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl13PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl13PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl13PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl13PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl13PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl13PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl13PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl13PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl13PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl13PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl13PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl13PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl13PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl13PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl13PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl13PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl13PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl13PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl13PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl13PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl13PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl13PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl13PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl13PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl13PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl13PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl13PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl13PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl13PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl13PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl13PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl13PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl13PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl13PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl13PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl13PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl13PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl13PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl13PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl13PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl13PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl13PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl13PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl13PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl13PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl13PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl13PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl13PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl13PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl13PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl13PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl13PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl13PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl13PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl13PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl13PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl13PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl13PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl13PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl13PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl13PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl13PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl13PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl13PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl13PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl13PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl13PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl13PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl13PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl13PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl13PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl13PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl13PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl13PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl13PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl13PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl13PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl13PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl13PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl13PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl13PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl13PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl13PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl13PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl13PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl13PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl13PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl13PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl13PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl13PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl13PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl13PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl13PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl13PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl13PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl13PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl13PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl13PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl13PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl13PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl13PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl13PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl13PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl13PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl13PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl13PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl13PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl13PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl13PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl13PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl13PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl13PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl13PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl13PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl13PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl13PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl13PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl13PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl13PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    glc->fnActiveTexture = (gl13PActiveTexture)wglGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl13PClientActiveTexture)wglGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl13PCompressedTexImage1D)wglGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl13PCompressedTexImage2D)wglGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl13PCompressedTexImage3D)wglGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl13PCompressedTexSubImage1D)wglGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl13PCompressedTexSubImage2D)wglGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl13PCompressedTexSubImage3D)wglGetProcAddress("glCompressedTexSubImage3D");
    return glc;
}

