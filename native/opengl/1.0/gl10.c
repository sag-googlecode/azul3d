#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl10.h"

#ifdef _WIN32
    HMODULE gl10OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl10OpenGL32 == NULL) {
            gl10OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl10OpenGL32, TEXT(name));
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

void gl10CopyPixels(gl10Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl10CullFace(gl10Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl10DeleteLists(gl10Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl10DepthFunc(gl10Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl10DrawPixels(gl10Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl10EnableClientState(gl10Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl10DisableClientState(gl10Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl10GetPolygonStipple(gl10Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl10Indexub(gl10Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl10Indexubv(gl10Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl10InterleavedArrays(gl10Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl10NormalPointer(gl10Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl10PrioritizeTextures(gl10Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl10PushAttrib(gl10Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl10PopAttrib(gl10Context* glc) {
    return glc->fnPopAttrib();
}

void gl10PushClientAttrib(gl10Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl10PopClientAttrib(gl10Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl10TexCoordPointer(gl10Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl10TexImage3DEXT(gl10Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl10TexSubImage1D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl10TexSubImage2D(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl10TexSubImage3DEXT(gl10Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl10VertexPointer(gl10Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl10Viewport(gl10Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

gl10Context* gl10NewContext() {
    gl10Context* glc = calloc(1, sizeof(gl10Context));

    // Preload all procedures
    glc->fnAccum = (gl10PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl10PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl10PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl10PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl10PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl10PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl10PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl10PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl10PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl10PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl10PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl10PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl10PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl10PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl10PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl10PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl10PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl10PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl10PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl10PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl10PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl10PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl10PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl10PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl10PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl10PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl10PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl10PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl10PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl10PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl10PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl10PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl10PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl10PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl10PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl10PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl10PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl10PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl10PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl10PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl10PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl10PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl10PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl10PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl10PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl10PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl10PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl10PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl10PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl10PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl10PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl10PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl10PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl10PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl10PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl10PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl10PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl10PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl10PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl10PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl10PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl10PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl10PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl10PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl10PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl10PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl10PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl10PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl10PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl10PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl10PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl10PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl10PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl10PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl10PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl10PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl10PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl10PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl10PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl10PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl10PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl10PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl10PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl10PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl10PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl10PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl10PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl10PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl10PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl10PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl10PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl10PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl10PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl10PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl10PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl10PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl10PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl10PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl10PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl10PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl10PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl10PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl10PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl10PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl10PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl10PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl10PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl10PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl10PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl10PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl10PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl10PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl10PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl10PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl10PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl10PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl10PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl10PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl10PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl10PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl10PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl10PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl10PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl10PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl10PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl10PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl10PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl10PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl10PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl10PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl10PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl10PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl10PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl10PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl10PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl10PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl10PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl10PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl10PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl10PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl10PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl10PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl10PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl10PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl10PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl10PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl10PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl10PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl10PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl10PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl10PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl10PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl10PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl10PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl10PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl10PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl10PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl10PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl10PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl10PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl10PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl10PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl10PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl10PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl10PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl10PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl10PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl10PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl10PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl10PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl10PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl10PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl10PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl10PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl10PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl10PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl10PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl10PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl10PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl10PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl10PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl10PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl10PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl10PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl10PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl10PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl10POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl10PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl10PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl10PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl10PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl10PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl10PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl10PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl10PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl10PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl10PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl10PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl10PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl10PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl10PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl10PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl10PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl10PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl10PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl10PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl10PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl10PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl10PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl10PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl10PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl10PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl10PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl10PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl10PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl10PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl10PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl10PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl10PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl10PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl10PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl10PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl10PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl10PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl10PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl10PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl10PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl10PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl10PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl10PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl10PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl10PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl10PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl10PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl10PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl10PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl10PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl10PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl10PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl10PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl10PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl10PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl10PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl10PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl10PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl10PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl10PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl10PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl10PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl10PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl10PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl10PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl10PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl10PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl10PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl10PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl10PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl10PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl10PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl10PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl10PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl10PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl10PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl10PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl10PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl10PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl10PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl10PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl10PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl10PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl10PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl10PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl10PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl10PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl10PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl10PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl10PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl10PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl10PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl10PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl10PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl10PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl10PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl10PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl10PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl10PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl10PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl10PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl10PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl10PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl10PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl10PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl10PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl10PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl10PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl10PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl10PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl10PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl10PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl10PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl10PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl10PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl10PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl10PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl10PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl10PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl10PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl10PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl10PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl10PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl10PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl10PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl10PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl10PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl10PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl10PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl10PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl10PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl10PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl10PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl10PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl10PViewport)doGetProcAddress("glViewport");
    return glc;
}

