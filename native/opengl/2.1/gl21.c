#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl21.h"

#ifdef _WIN32
    HMODULE gl21OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl21OpenGL32 == NULL) {
            gl21OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl21OpenGL32, TEXT(name));
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

void gl21CopyPixels(gl21Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl21CullFace(gl21Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl21DeleteLists(gl21Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl21DepthFunc(gl21Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl21DrawPixels(gl21Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl21EnableClientState(gl21Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl21DisableClientState(gl21Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl21GetPolygonStipple(gl21Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl21Indexub(gl21Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl21Indexubv(gl21Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl21InterleavedArrays(gl21Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl21NormalPointer(gl21Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl21PrioritizeTextures(gl21Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl21PushAttrib(gl21Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl21PopAttrib(gl21Context* glc) {
    return glc->fnPopAttrib();
}

void gl21PushClientAttrib(gl21Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl21PopClientAttrib(gl21Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl21TexCoordPointer(gl21Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl21TexImage3DEXT(gl21Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl21TexSubImage1D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl21TexSubImage2D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl21TexSubImage3DEXT(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl21VertexPointer(gl21Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl21Viewport(gl21Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
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

void gl21BlendColorEXT(gl21Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl21BlendEquation(gl21Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl21CopyTexSubImage3D(gl21Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
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

void gl21BlendFuncSeparate(gl21Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
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

gl21Context* gl21NewContext() {
    gl21Context* glc = calloc(1, sizeof(gl21Context));

    // Preload all procedures
    glc->fnAccum = (gl21PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl21PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl21PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl21PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl21PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl21PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl21PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl21PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl21PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl21PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl21PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl21PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl21PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl21PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl21PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl21PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl21PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl21PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl21PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl21PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl21PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl21PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl21PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl21PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl21PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl21PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl21PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl21PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl21PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl21PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl21PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl21PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl21PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl21PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl21PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl21PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl21PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl21PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl21PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl21PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl21PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl21PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl21PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl21PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl21PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl21PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl21PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl21PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl21PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl21PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl21PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl21PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl21PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl21PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl21PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl21PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl21PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl21PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl21PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl21PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl21PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl21PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl21PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl21PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl21PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl21PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl21PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl21PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl21PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl21PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl21PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl21PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl21PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl21PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl21PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl21PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl21PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl21PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl21PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl21PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl21PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl21PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl21PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl21PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl21PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl21PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl21PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl21PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl21PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl21PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl21PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl21PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl21PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl21PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl21PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl21PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl21PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl21PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl21PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl21PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl21PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl21PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl21PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl21PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl21PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl21PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl21PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl21PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl21PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl21PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl21PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl21PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl21PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl21PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl21PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl21PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl21PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl21PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl21PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl21PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl21PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl21PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl21PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl21PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl21PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl21PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl21PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl21PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl21PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl21PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl21PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl21PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl21PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl21PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl21PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl21PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl21PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl21PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl21PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl21PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl21PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl21PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl21PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl21PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl21PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl21PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl21PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl21PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl21PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl21PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl21PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl21PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl21PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl21PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl21PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl21PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl21PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl21PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl21PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl21PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl21PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl21PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl21PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl21PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl21PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl21PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl21PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl21PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl21PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl21PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl21PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl21PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl21PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl21PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl21PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl21PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl21PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl21PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl21PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl21PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl21PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl21PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl21PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl21PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl21PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl21PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl21POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl21PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl21PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl21PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl21PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl21PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl21PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl21PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl21PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl21PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl21PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl21PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl21PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl21PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl21PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl21PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl21PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl21PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl21PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl21PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl21PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl21PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl21PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl21PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl21PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl21PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl21PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl21PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl21PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl21PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl21PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl21PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl21PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl21PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl21PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl21PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl21PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl21PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl21PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl21PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl21PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl21PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl21PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl21PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl21PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl21PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl21PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl21PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl21PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl21PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl21PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl21PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl21PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl21PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl21PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl21PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl21PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl21PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl21PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl21PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl21PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl21PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl21PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl21PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl21PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl21PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl21PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl21PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl21PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl21PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl21PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl21PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl21PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl21PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl21PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl21PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl21PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl21PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl21PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl21PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl21PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl21PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl21PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl21PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl21PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl21PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl21PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl21PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl21PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl21PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl21PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl21PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl21PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl21PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl21PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl21PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl21PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl21PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl21PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl21PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl21PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl21PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl21PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl21PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl21PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl21PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl21PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl21PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl21PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl21PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl21PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl21PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl21PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl21PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl21PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl21PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl21PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl21PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl21PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl21PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl21PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl21PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl21PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl21PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl21PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl21PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl21PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl21PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl21PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl21PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl21PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl21PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl21PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl21PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl21PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl21PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl21PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl21PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl21PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl21PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl21PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl21PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl21PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl21PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl21PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl21PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl21PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl21PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl21PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl21PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl21PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl21PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl21PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl21PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    glc->fnActiveTexture = (gl21PActiveTexture)wglGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl21PClientActiveTexture)wglGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl21PCompressedTexImage1D)wglGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl21PCompressedTexImage2D)wglGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl21PCompressedTexImage3D)wglGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl21PCompressedTexSubImage1D)wglGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl21PCompressedTexSubImage2D)wglGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl21PCompressedTexSubImage3D)wglGetProcAddress("glCompressedTexSubImage3D");
    glc->fnBlendFuncSeparate = (gl21PBlendFuncSeparate)wglGetProcAddress("glBlendFuncSeparate");
    glc->fnBeginQuery = (gl21PBeginQuery)wglGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl21PBindBuffer)wglGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl21PBufferData)wglGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl21PBufferSubData)wglGetProcAddress("glBufferSubData");
    glc->fnAttachShader = (gl21PAttachShader)doGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl21PBindAttribLocation)doGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl21PBlendEquationSeperate)doGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl21PCompileShader)doGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl21PCreateProgram)doGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl21PCreateShader)doGetProcAddress("glCreateShader");
    return glc;
}

