#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl12.h"

#ifdef _WIN32
    HMODULE gl12OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl12OpenGL32 == NULL) {
            gl12OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl12OpenGL32, TEXT(name));
    }
#endif

void gl12Accum(gl12Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl12AlphaFunc(gl12Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl12Begin(gl12Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl12End(gl12Context* glc) {
    return glc->fnEnd();
}

void gl12Bitmap(gl12Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl12BlendFunc(gl12Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl12CallList(gl12Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl12CallLists(gl12Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl12Clear(gl12Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl12ClearAccum(gl12Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl12ClearColor(gl12Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl12ClearDepth(gl12Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl12ClearIndex(gl12Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl12ClearStencil(gl12Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl12ClipPlane(gl12Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl12Color3b(gl12Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl12Color3d(gl12Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl12Color3f(gl12Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl12Color3i(gl12Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl12Color3s(gl12Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl12Color3ub(gl12Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl12Color3ui(gl12Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl12Color3us(gl12Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl12Color4b(gl12Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl12Color4d(gl12Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl12Color4f(gl12Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl12Color4i(gl12Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl12Color4s(gl12Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl12Color4ub(gl12Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl12Color4ui(gl12Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl12Color4us(gl12Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl12Color3bv(gl12Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl12Color3dv(gl12Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl12Color3fv(gl12Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl12Color3iv(gl12Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl12Color3sv(gl12Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl12Color3ubv(gl12Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl12Color3uiv(gl12Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl12Color3usv(gl12Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl12Color4bv(gl12Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl12Color4dv(gl12Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl12Color4fv(gl12Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl12Color4iv(gl12Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl12Color4sv(gl12Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl12Color4ubv(gl12Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl12Color4uiv(gl12Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl12Color4usv(gl12Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl12ColorMask(gl12Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl12ColorMaterial(gl12Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl12ColorTable(gl12Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl12ColorTableParameterfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl12ColorTableParameteriv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl12ColorSubTable(gl12Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl12CopyPixels(gl12Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl12CullFace(gl12Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl12ConvolutionFilter1D(gl12Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl12ConvolutionFilter2D(gl12Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl12ConvolutionParameterf(gl12Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl12ConvolutionParameteri(gl12Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl12CopyColorTable(gl12Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl12CopyColorSubTable(gl12Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl12CopyConvolutionFilter1D(gl12Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl12CopyConvolutionFilter2D(gl12Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl12DeleteLists(gl12Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl12DepthFunc(gl12Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl12DepthRange(gl12Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl12Enable(gl12Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl12Disable(gl12Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl12DrawBuffer(gl12Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl12DrawPixels(gl12Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
}

void gl12EdgeFlag(gl12Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl12EdgeFlagv(gl12Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl12EdgeFlagPointer(gl12Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl12EnableClientState(gl12Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl12DisableClientState(gl12Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl12EvalCoord1d(gl12Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl12EvalCoord1f(gl12Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl12EvalCoord2d(gl12Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl12EvalCoord2f(gl12Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl12EvalCoord1dv(gl12Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl12EvalCoord1fv(gl12Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl12EvalCoord2dv(gl12Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl12EvalCoord2fv(gl12Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl12EvalMesh1(gl12Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl12EvalMesh2(gl12Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl12EvalPoint1(gl12Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl12EvalPoint2(gl12Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl12FeedbackBuffer(gl12Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl12Finish(gl12Context* glc) {
    return glc->fnFinish();
}

void gl12Flush(gl12Context* glc) {
    return glc->fnFlush();
}

void gl12Fogf(gl12Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl12Fogi(gl12Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl12Fogfv(gl12Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl12Fogiv(gl12Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl12FrontFace(gl12Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl12Frustum(gl12Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl12GenLists(gl12Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl12GetBooleanv(gl12Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl12GetDoublev(gl12Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl12GetFloatv(gl12Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl12GetIntegerv(gl12Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl12GetClipPlane(gl12Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl12GetError(gl12Context* glc) {
    return glc->fnGetError();
}

void gl12GetLightfv(gl12Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl12GetLightiv(gl12Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl12GetMapdv(gl12Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl12GetMapfv(gl12Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl12GetMapiv(gl12Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl12GetMaterialfv(gl12Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl12GetMaterialiv(gl12Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl12GetPixelMapfv(gl12Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl12GetPixelMapuiv(gl12Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl12GetPixelMapusv(gl12Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl12GetPolygonStipple(gl12Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
}

GLubyte* gl12GetString(gl12Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl12GetTexEnvfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl12GetTexEnviv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl12GetTexGendv(gl12Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl12GetTexGenfv(gl12Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl12GetTexGeniv(gl12Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl12GetTexImage(gl12Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl12GetTexLevelParameterfv(gl12Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl12GetTexLevelParameteriv(gl12Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl12GetTexParameterfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl12GetTexParameteriv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl12Hint(gl12Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl12Indexd(gl12Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl12Indexf(gl12Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl12Indexi(gl12Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl12Indexs(gl12Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl12Indexub(gl12Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl12Indexdv(gl12Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl12Indexfv(gl12Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl12Indexiv(gl12Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl12Indexsv(gl12Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl12Indexubv(gl12Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl12IndexMask(gl12Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl12IndexPointer(gl12Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl12InitNames(gl12Context* glc) {
    return glc->fnInitNames();
}

void gl12InterleavedArrays(gl12Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl12IsEnabled(gl12Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl12IsList(gl12Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl12Lightf(gl12Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl12Lighti(gl12Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl12Lightfv(gl12Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl12Lightiv(gl12Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl12LightModelf(gl12Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl12LightModeli(gl12Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl12LightModelfv(gl12Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl12LightModeliv(gl12Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl12LineStipple(gl12Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl12LineWidth(gl12Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl12ListBase(gl12Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl12LoadIdentity(gl12Context* glc) {
    return glc->fnLoadIdentity();
}

void gl12LoadMatrixd(gl12Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl12LoadMatrixf(gl12Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl12LoadName(gl12Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl12LogicOp(gl12Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl12Map1d(gl12Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl12Map1f(gl12Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl12Map2d(gl12Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl12Map2f(gl12Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl12MapGrid1d(gl12Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl12MapGrid1f(gl12Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl12MapGrid2d(gl12Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl12MapGrid2f(gl12Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl12Materialf(gl12Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl12Materiali(gl12Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl12Materialfv(gl12Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl12Materialiv(gl12Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl12MatrixMode(gl12Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl12MultMatrixd(gl12Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl12MultMatrixf(gl12Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl12NewList(gl12Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl12EndList(gl12Context* glc) {
    return glc->fnEndList();
}

void gl12Normal3b(gl12Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl12Normal3d(gl12Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl12Normal3f(gl12Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl12Normal3i(gl12Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl12Normal3s(gl12Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl12Normal3bv(gl12Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl12Normal3dv(gl12Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl12Normal3fv(gl12Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl12Normal3iv(gl12Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl12Normal3sv(gl12Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl12NormalPointer(gl12Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl12Ortho(gl12Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl12PassThrough(gl12Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl12PixelMapfv(gl12Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl12PixelMapuiv(gl12Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl12PixelMapusv(gl12Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl12PixelStoref(gl12Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl12PixelStorei(gl12Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl12PixelTransferf(gl12Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl12PixelTransferi(gl12Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl12PixelZoom(gl12Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl12PointSize(gl12Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl12PolygonMode(gl12Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl12PolygonStipple(gl12Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl12PrioritizeTextures(gl12Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl12PushAttrib(gl12Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl12PopAttrib(gl12Context* glc) {
    return glc->fnPopAttrib();
}

void gl12PushClientAttrib(gl12Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl12PopClientAttrib(gl12Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl12PushMatrix(gl12Context* glc) {
    return glc->fnPushMatrix();
}

void gl12PopMatrix(gl12Context* glc) {
    return glc->fnPopMatrix();
}

void gl12PushName(gl12Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl12PopName(gl12Context* glc) {
    return glc->fnPopName();
}

void gl12RasterPos2d(gl12Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl12RasterPos2f(gl12Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl12RasterPos2i(gl12Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl12RasterPos2s(gl12Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl12RasterPos3d(gl12Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl12RasterPos3f(gl12Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl12RasterPos3i(gl12Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl12RasterPos3s(gl12Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl12RasterPos4d(gl12Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl12RasterPos4f(gl12Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl12RasterPos4i(gl12Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl12RasterPos4s(gl12Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl12RasterPos2dv(gl12Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl12RasterPos2fv(gl12Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl12RasterPos2iv(gl12Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl12RasterPos2sv(gl12Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl12RasterPos3dv(gl12Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl12RasterPos3fv(gl12Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl12RasterPos3iv(gl12Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl12RasterPos3sv(gl12Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl12RasterPos4dv(gl12Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl12RasterPos4fv(gl12Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl12RasterPos4iv(gl12Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl12RasterPos4sv(gl12Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl12ReadBuffer(gl12Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl12ReadPixels(gl12Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl12Rectd(gl12Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl12Rectf(gl12Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl12Recti(gl12Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl12Rects(gl12Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl12Rectdv(gl12Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl12Rectfv(gl12Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl12Rectiv(gl12Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl12Rectsv(gl12Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl12RenderMode(gl12Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl12Rotated(gl12Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl12Rotatef(gl12Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl12Scaled(gl12Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl12Scalef(gl12Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl12Scissor(gl12Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl12SelectBuffer(gl12Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl12ShadeModel(gl12Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl12StencilFunc(gl12Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl12StencilMask(gl12Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl12StencilOp(gl12Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl12TexCoord1d(gl12Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl12TexCoord1f(gl12Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl12TexCoord1i(gl12Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl12TexCoord1s(gl12Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl12TexCoord2d(gl12Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl12TexCoord2f(gl12Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl12TexCoord2i(gl12Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl12TexCoord2s(gl12Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl12TexCoord3d(gl12Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl12TexCoord3f(gl12Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl12TexCoord3i(gl12Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl12TexCoord3s(gl12Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl12TexCoord4d(gl12Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl12TexCoord4f(gl12Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl12TexCoord4i(gl12Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl12TexCoord4s(gl12Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl12TexCoord1dv(gl12Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl12TexCoord1fv(gl12Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl12TexCoord1iv(gl12Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl12TexCoord1sv(gl12Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl12TexCoord2dv(gl12Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl12TexCoord2fv(gl12Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl12TexCoord2iv(gl12Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl12TexCoord2sv(gl12Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl12TexCoord3dv(gl12Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl12TexCoord3fv(gl12Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl12TexCoord3iv(gl12Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl12TexCoord3sv(gl12Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl12TexCoord4dv(gl12Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl12TexCoord4fv(gl12Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl12TexCoord4iv(gl12Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl12TexCoord4sv(gl12Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl12TexCoordPointer(gl12Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl12TexEnvf(gl12Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl12TexEnvi(gl12Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl12TexEnvfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl12TexEnviv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl12TexGend(gl12Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl12TexGenf(gl12Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl12TexGeni(gl12Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl12TexGendv(gl12Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl12TexGenfv(gl12Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl12TexGeniv(gl12Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl12TexImage1D(gl12Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl12TexImage2D(gl12Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl12TexImage3DEXT(gl12Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl12TexParameterf(gl12Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl12TexParameteri(gl12Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl12TexParameterfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl12TexParameteriv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl12TexSubImage1D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl12TexSubImage2D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl12TexSubImage3DEXT(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl12Translated(gl12Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl12Translatef(gl12Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl12Vertex2s(gl12Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl12Vertex2i(gl12Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl12Vertex2f(gl12Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl12Vertex2d(gl12Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl12Vertex3s(gl12Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl12Vertex3i(gl12Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl12Vertex3f(gl12Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl12Vertex3d(gl12Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl12Vertex4s(gl12Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl12Vertex4i(gl12Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl12Vertex4f(gl12Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl12Vertex4d(gl12Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl12VertexPointer(gl12Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl12Viewport(gl12Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

GLboolean gl12AreTexturesResident(gl12Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl12ArrayElement(gl12Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl12DrawArrays(gl12Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl12DrawElements(gl12Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl12GetPointerv(gl12Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl12PolygonOffset(gl12Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl12CopyTexImage1D(gl12Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl12CopyTexImage2D(gl12Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl12CopyTexSubImage1D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl12CopyTexSubImage2D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl12BindTexture(gl12Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl12DeleteTextures(gl12Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl12GenTextures(gl12Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl12IsTexture(gl12Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl12ColorPointer(gl12Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl12BlendColorEXT(gl12Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl12BlendEquation(gl12Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl12CopyTexSubImage3D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

gl12Context* gl12NewContext() {
    gl12Context* glc = calloc(1, sizeof(gl12Context));

    // Preload all procedures
    glc->fnAccum = (gl12PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl12PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl12PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl12PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl12PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl12PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl12PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl12PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl12PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl12PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl12PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl12PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl12PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl12PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl12PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl12PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl12PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl12PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl12PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl12PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl12PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl12PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl12PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl12PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl12PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl12PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl12PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl12PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl12PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl12PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl12PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl12PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl12PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl12PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl12PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl12PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl12PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl12PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl12PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl12PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl12PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl12PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl12PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl12PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl12PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl12PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl12PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl12PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl12PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl12PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl12PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl12PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl12PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl12PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl12PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl12PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl12PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl12PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl12PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl12PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl12PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl12PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl12PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl12PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl12PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl12PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl12PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl12PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl12PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl12PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl12PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl12PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl12PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl12PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl12PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl12PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl12PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl12PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl12PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl12PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl12PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl12PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl12PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl12PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl12PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl12PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl12PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl12PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl12PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl12PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl12PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl12PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl12PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl12PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl12PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl12PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl12PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl12PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl12PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl12PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl12PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl12PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl12PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl12PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl12PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl12PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl12PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl12PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl12PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl12PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl12PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl12PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl12PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl12PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl12PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl12PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl12PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl12PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl12PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl12PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl12PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl12PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl12PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl12PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl12PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl12PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl12PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl12PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl12PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl12PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl12PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl12PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl12PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl12PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl12PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl12PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl12PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl12PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl12PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl12PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl12PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl12PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl12PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl12PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl12PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl12PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl12PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl12PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl12PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl12PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl12PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl12PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl12PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl12PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl12PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl12PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl12PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl12PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl12PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl12PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl12PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl12PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl12PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl12PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl12PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl12PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl12PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl12PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl12PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl12PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl12PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl12PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl12PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl12PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl12PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl12PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl12PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl12PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl12PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl12PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl12PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl12PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl12PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl12PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl12PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl12PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl12POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl12PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl12PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl12PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl12PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl12PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl12PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl12PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl12PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl12PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl12PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl12PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl12PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl12PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl12PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl12PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl12PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl12PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl12PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl12PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl12PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl12PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl12PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl12PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl12PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl12PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl12PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl12PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl12PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl12PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl12PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl12PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl12PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl12PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl12PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl12PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl12PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl12PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl12PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl12PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl12PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl12PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl12PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl12PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl12PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl12PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl12PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl12PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl12PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl12PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl12PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl12PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl12PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl12PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl12PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl12PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl12PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl12PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl12PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl12PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl12PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl12PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl12PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl12PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl12PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl12PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl12PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl12PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl12PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl12PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl12PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl12PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl12PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl12PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl12PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl12PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl12PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl12PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl12PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl12PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl12PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl12PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl12PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl12PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl12PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl12PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl12PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl12PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl12PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl12PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl12PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl12PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl12PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl12PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl12PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl12PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl12PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl12PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl12PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl12PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl12PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl12PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl12PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl12PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl12PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl12PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl12PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl12PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl12PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl12PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl12PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl12PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl12PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl12PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl12PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl12PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl12PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl12PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl12PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl12PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl12PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl12PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl12PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl12PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl12PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl12PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl12PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl12PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl12PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl12PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl12PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl12PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl12PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl12PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl12PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl12PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl12PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl12PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl12PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl12PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl12PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl12PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl12PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl12PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl12PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl12PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl12PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl12PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl12PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl12PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl12PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl12PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl12PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl12PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    return glc;
}

