#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl14.h"

#ifdef _WIN32
    HMODULE gl14OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl14OpenGL32 == NULL) {
            gl14OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl14OpenGL32, TEXT(name));
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

void gl14CopyPixels(gl14Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl14CullFace(gl14Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl14DeleteLists(gl14Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl14DepthFunc(gl14Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl14DrawPixels(gl14Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl14EnableClientState(gl14Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl14DisableClientState(gl14Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl14GetPolygonStipple(gl14Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl14Indexub(gl14Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl14Indexubv(gl14Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl14InterleavedArrays(gl14Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl14NormalPointer(gl14Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl14PrioritizeTextures(gl14Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl14PushAttrib(gl14Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl14PopAttrib(gl14Context* glc) {
    return glc->fnPopAttrib();
}

void gl14PushClientAttrib(gl14Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl14PopClientAttrib(gl14Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl14TexCoordPointer(gl14Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl14TexImage3DEXT(gl14Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl14TexSubImage1D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl14TexSubImage2D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl14TexSubImage3DEXT(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl14VertexPointer(gl14Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl14Viewport(gl14Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
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

void gl14BlendColorEXT(gl14Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl14BlendEquation(gl14Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl14CopyTexSubImage3D(gl14Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
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

void gl14BlendFuncSeparate(gl14Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

gl14Context* gl14NewContext() {
    gl14Context* glc = calloc(1, sizeof(gl14Context));

    // Preload all procedures
    glc->fnAccum = (gl14PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl14PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl14PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl14PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl14PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl14PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl14PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl14PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl14PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl14PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl14PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl14PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl14PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl14PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl14PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl14PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl14PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl14PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl14PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl14PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl14PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl14PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl14PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl14PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl14PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl14PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl14PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl14PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl14PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl14PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl14PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl14PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl14PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl14PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl14PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl14PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl14PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl14PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl14PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl14PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl14PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl14PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl14PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl14PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl14PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl14PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl14PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl14PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl14PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl14PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl14PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl14PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl14PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl14PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl14PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl14PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl14PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl14PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl14PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl14PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl14PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl14PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl14PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl14PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl14PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl14PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl14PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl14PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl14PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl14PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl14PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl14PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl14PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl14PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl14PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl14PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl14PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl14PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl14PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl14PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl14PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl14PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl14PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl14PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl14PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl14PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl14PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl14PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl14PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl14PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl14PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl14PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl14PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl14PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl14PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl14PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl14PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl14PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl14PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl14PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl14PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl14PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl14PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl14PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl14PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl14PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl14PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl14PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl14PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl14PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl14PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl14PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl14PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl14PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl14PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl14PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl14PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl14PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl14PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl14PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl14PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl14PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl14PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl14PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl14PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl14PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl14PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl14PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl14PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl14PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl14PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl14PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl14PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl14PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl14PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl14PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl14PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl14PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl14PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl14PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl14PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl14PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl14PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl14PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl14PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl14PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl14PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl14PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl14PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl14PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl14PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl14PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl14PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl14PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl14PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl14PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl14PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl14PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl14PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl14PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl14PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl14PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl14PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl14PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl14PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl14PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl14PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl14PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl14PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl14PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl14PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl14PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl14PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl14PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl14PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl14PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl14PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl14PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl14PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl14PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl14PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl14PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl14PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl14PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl14PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl14PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl14POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl14PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl14PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl14PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl14PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl14PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl14PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl14PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl14PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl14PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl14PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl14PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl14PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl14PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl14PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl14PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl14PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl14PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl14PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl14PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl14PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl14PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl14PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl14PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl14PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl14PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl14PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl14PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl14PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl14PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl14PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl14PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl14PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl14PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl14PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl14PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl14PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl14PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl14PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl14PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl14PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl14PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl14PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl14PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl14PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl14PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl14PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl14PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl14PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl14PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl14PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl14PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl14PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl14PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl14PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl14PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl14PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl14PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl14PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl14PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl14PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl14PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl14PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl14PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl14PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl14PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl14PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl14PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl14PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl14PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl14PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl14PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl14PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl14PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl14PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl14PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl14PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl14PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl14PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl14PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl14PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl14PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl14PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl14PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl14PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl14PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl14PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl14PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl14PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl14PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl14PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl14PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl14PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl14PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl14PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl14PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl14PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl14PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl14PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl14PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl14PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl14PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl14PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl14PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl14PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl14PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl14PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl14PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl14PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl14PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl14PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl14PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl14PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl14PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl14PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl14PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl14PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl14PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl14PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl14PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl14PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl14PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl14PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl14PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl14PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl14PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl14PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl14PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl14PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl14PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl14PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl14PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl14PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl14PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl14PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl14PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl14PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl14PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl14PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl14PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl14PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl14PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl14PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl14PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl14PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl14PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl14PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl14PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl14PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl14PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl14PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl14PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl14PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl14PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    glc->fnActiveTexture = (gl14PActiveTexture)wglGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl14PClientActiveTexture)wglGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl14PCompressedTexImage1D)wglGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl14PCompressedTexImage2D)wglGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl14PCompressedTexImage3D)wglGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl14PCompressedTexSubImage1D)wglGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl14PCompressedTexSubImage2D)wglGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl14PCompressedTexSubImage3D)wglGetProcAddress("glCompressedTexSubImage3D");
    glc->fnBlendFuncSeparate = (gl14PBlendFuncSeparate)wglGetProcAddress("glBlendFuncSeparate");
    return glc;
}

