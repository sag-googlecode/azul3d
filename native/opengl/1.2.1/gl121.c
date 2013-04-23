#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl121.h"

#ifdef _WIN32
    HMODULE gl121OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl121OpenGL32 == NULL) {
            gl121OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl121OpenGL32, TEXT(name));
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

void gl121CopyPixels(gl121Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl121CullFace(gl121Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl121DeleteLists(gl121Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl121DepthFunc(gl121Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl121DrawPixels(gl121Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl121EnableClientState(gl121Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl121DisableClientState(gl121Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl121GetPolygonStipple(gl121Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl121Indexub(gl121Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl121Indexubv(gl121Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl121InterleavedArrays(gl121Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl121NormalPointer(gl121Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl121PrioritizeTextures(gl121Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl121PushAttrib(gl121Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl121PopAttrib(gl121Context* glc) {
    return glc->fnPopAttrib();
}

void gl121PushClientAttrib(gl121Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl121PopClientAttrib(gl121Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl121TexCoordPointer(gl121Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl121TexImage3DEXT(gl121Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl121TexSubImage1D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl121TexSubImage2D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl121TexSubImage3DEXT(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl121VertexPointer(gl121Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl121Viewport(gl121Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
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

void gl121BlendColorEXT(gl121Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl121BlendEquation(gl121Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl121CopyTexSubImage3D(gl121Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

gl121Context* gl121NewContext() {
    gl121Context* glc = calloc(1, sizeof(gl121Context));

    // Preload all procedures
    glc->fnAccum = (gl121PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl121PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl121PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl121PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl121PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl121PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl121PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl121PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl121PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl121PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl121PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl121PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl121PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl121PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl121PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl121PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl121PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl121PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl121PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl121PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl121PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl121PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl121PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl121PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl121PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl121PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl121PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl121PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl121PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl121PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl121PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl121PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl121PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl121PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl121PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl121PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl121PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl121PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl121PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl121PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl121PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl121PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl121PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl121PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl121PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl121PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl121PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl121PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl121PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl121PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl121PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl121PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl121PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl121PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl121PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl121PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl121PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl121PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl121PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl121PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl121PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl121PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl121PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl121PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl121PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl121PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl121PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl121PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl121PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl121PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl121PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl121PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl121PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl121PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl121PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl121PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl121PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl121PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl121PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl121PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl121PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl121PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl121PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl121PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl121PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl121PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl121PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl121PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl121PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl121PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl121PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl121PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl121PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl121PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl121PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl121PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl121PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl121PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl121PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl121PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl121PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl121PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl121PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl121PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl121PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl121PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl121PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl121PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl121PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl121PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl121PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl121PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl121PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl121PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl121PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl121PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl121PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl121PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl121PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl121PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl121PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl121PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl121PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl121PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl121PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl121PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl121PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl121PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl121PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl121PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl121PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl121PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl121PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl121PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl121PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl121PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl121PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl121PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl121PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl121PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl121PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl121PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl121PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl121PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl121PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl121PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl121PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl121PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl121PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl121PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl121PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl121PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl121PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl121PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl121PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl121PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl121PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl121PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl121PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl121PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl121PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl121PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl121PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl121PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl121PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl121PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl121PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl121PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl121PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl121PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl121PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl121PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl121PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl121PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl121PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl121PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl121PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl121PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl121PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl121PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl121PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl121PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl121PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl121PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl121PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl121PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl121POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl121PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl121PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl121PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl121PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl121PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl121PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl121PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl121PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl121PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl121PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl121PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl121PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl121PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl121PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl121PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl121PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl121PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl121PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl121PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl121PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl121PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl121PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl121PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl121PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl121PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl121PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl121PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl121PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl121PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl121PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl121PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl121PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl121PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl121PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl121PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl121PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl121PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl121PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl121PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl121PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl121PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl121PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl121PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl121PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl121PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl121PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl121PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl121PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl121PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl121PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl121PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl121PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl121PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl121PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl121PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl121PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl121PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl121PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl121PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl121PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl121PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl121PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl121PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl121PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl121PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl121PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl121PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl121PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl121PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl121PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl121PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl121PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl121PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl121PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl121PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl121PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl121PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl121PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl121PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl121PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl121PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl121PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl121PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl121PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl121PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl121PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl121PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl121PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl121PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl121PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl121PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl121PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl121PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl121PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl121PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl121PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl121PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl121PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl121PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl121PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl121PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl121PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl121PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl121PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl121PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl121PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl121PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl121PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl121PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl121PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl121PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl121PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl121PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl121PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl121PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl121PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl121PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl121PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl121PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl121PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl121PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl121PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl121PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl121PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl121PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl121PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl121PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl121PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl121PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl121PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl121PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl121PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl121PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl121PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl121PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl121PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl121PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl121PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl121PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl121PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl121PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl121PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl121PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl121PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl121PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl121PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl121PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl121PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl121PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl121PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl121PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl121PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl121PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    return glc;
}

