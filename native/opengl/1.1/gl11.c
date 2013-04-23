#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl11.h"

#ifdef _WIN32
    HMODULE gl11OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl11OpenGL32 == NULL) {
            gl11OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl11OpenGL32, TEXT(name));
    }
#endif

void gl11Accum(gl11Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl11AlphaFunc(gl11Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl11Begin(gl11Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl11End(gl11Context* glc) {
    return glc->fnEnd();
}

void gl11Bitmap(gl11Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl11BlendFunc(gl11Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl11CallList(gl11Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl11CallLists(gl11Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl11Clear(gl11Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl11ClearAccum(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl11ClearColor(gl11Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl11ClearDepth(gl11Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl11ClearIndex(gl11Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl11ClearStencil(gl11Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl11ClipPlane(gl11Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl11Color3b(gl11Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl11Color3d(gl11Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl11Color3f(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl11Color3i(gl11Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl11Color3s(gl11Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl11Color3ub(gl11Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl11Color3ui(gl11Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl11Color3us(gl11Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl11Color4b(gl11Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl11Color4d(gl11Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl11Color4f(gl11Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl11Color4i(gl11Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl11Color4s(gl11Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl11Color4ub(gl11Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl11Color4ui(gl11Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl11Color4us(gl11Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl11Color3bv(gl11Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl11Color3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl11Color3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl11Color3iv(gl11Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl11Color3sv(gl11Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl11Color3ubv(gl11Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl11Color3uiv(gl11Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl11Color3usv(gl11Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl11Color4bv(gl11Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl11Color4dv(gl11Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl11Color4fv(gl11Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl11Color4iv(gl11Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl11Color4sv(gl11Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl11Color4ubv(gl11Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl11Color4uiv(gl11Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl11Color4usv(gl11Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl11ColorMask(gl11Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl11ColorMaterial(gl11Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl11ColorTable(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl11ColorTableParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl11ColorTableParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl11ColorSubTable(gl11Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl11CopyPixels(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl11CullFace(gl11Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl11ConvolutionFilter1D(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl11ConvolutionFilter2D(gl11Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl11ConvolutionParameterf(gl11Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl11ConvolutionParameteri(gl11Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl11CopyColorTable(gl11Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl11CopyColorSubTable(gl11Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl11CopyConvolutionFilter1D(gl11Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl11CopyConvolutionFilter2D(gl11Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl11DeleteLists(gl11Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl11DepthFunc(gl11Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl11DepthRange(gl11Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl11Enable(gl11Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl11Disable(gl11Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl11DrawBuffer(gl11Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl11DrawPixels(gl11Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
}

void gl11EdgeFlag(gl11Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl11EdgeFlagv(gl11Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl11EdgeFlagPointer(gl11Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl11EnableClientState(gl11Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl11DisableClientState(gl11Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl11EvalCoord1d(gl11Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl11EvalCoord1f(gl11Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl11EvalCoord2d(gl11Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl11EvalCoord2f(gl11Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl11EvalCoord1dv(gl11Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl11EvalCoord1fv(gl11Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl11EvalCoord2dv(gl11Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl11EvalCoord2fv(gl11Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl11EvalMesh1(gl11Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl11EvalMesh2(gl11Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl11EvalPoint1(gl11Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl11EvalPoint2(gl11Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl11FeedbackBuffer(gl11Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl11Finish(gl11Context* glc) {
    return glc->fnFinish();
}

void gl11Flush(gl11Context* glc) {
    return glc->fnFlush();
}

void gl11Fogf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl11Fogi(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl11Fogfv(gl11Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl11Fogiv(gl11Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl11FrontFace(gl11Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl11Frustum(gl11Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl11GenLists(gl11Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl11GetBooleanv(gl11Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl11GetDoublev(gl11Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl11GetFloatv(gl11Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl11GetIntegerv(gl11Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl11GetClipPlane(gl11Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl11GetError(gl11Context* glc) {
    return glc->fnGetError();
}

void gl11GetLightfv(gl11Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl11GetLightiv(gl11Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl11GetMapdv(gl11Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl11GetMapfv(gl11Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl11GetMapiv(gl11Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl11GetMaterialfv(gl11Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl11GetMaterialiv(gl11Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl11GetPixelMapfv(gl11Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl11GetPixelMapuiv(gl11Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl11GetPixelMapusv(gl11Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl11GetPolygonStipple(gl11Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
}

GLubyte* gl11GetString(gl11Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl11GetTexEnvfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl11GetTexEnviv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl11GetTexGendv(gl11Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl11GetTexGenfv(gl11Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl11GetTexGeniv(gl11Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl11GetTexImage(gl11Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl11GetTexLevelParameterfv(gl11Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl11GetTexLevelParameteriv(gl11Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl11GetTexParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl11GetTexParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl11Hint(gl11Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl11Indexd(gl11Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl11Indexf(gl11Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl11Indexi(gl11Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl11Indexs(gl11Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl11Indexub(gl11Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl11Indexdv(gl11Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl11Indexfv(gl11Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl11Indexiv(gl11Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl11Indexsv(gl11Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl11Indexubv(gl11Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl11IndexMask(gl11Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl11IndexPointer(gl11Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl11InitNames(gl11Context* glc) {
    return glc->fnInitNames();
}

void gl11InterleavedArrays(gl11Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl11IsEnabled(gl11Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl11IsList(gl11Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl11Lightf(gl11Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl11Lighti(gl11Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl11Lightfv(gl11Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl11Lightiv(gl11Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl11LightModelf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl11LightModeli(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl11LightModelfv(gl11Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl11LightModeliv(gl11Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl11LineStipple(gl11Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl11LineWidth(gl11Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl11ListBase(gl11Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl11LoadIdentity(gl11Context* glc) {
    return glc->fnLoadIdentity();
}

void gl11LoadMatrixd(gl11Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl11LoadMatrixf(gl11Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl11LoadName(gl11Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl11LogicOp(gl11Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl11Map1d(gl11Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl11Map1f(gl11Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl11Map2d(gl11Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl11Map2f(gl11Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl11MapGrid1d(gl11Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl11MapGrid1f(gl11Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl11MapGrid2d(gl11Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl11MapGrid2f(gl11Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl11Materialf(gl11Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl11Materiali(gl11Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl11Materialfv(gl11Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl11Materialiv(gl11Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl11MatrixMode(gl11Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl11MultMatrixd(gl11Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl11MultMatrixf(gl11Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl11NewList(gl11Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl11EndList(gl11Context* glc) {
    return glc->fnEndList();
}

void gl11Normal3b(gl11Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl11Normal3d(gl11Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl11Normal3f(gl11Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl11Normal3i(gl11Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl11Normal3s(gl11Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl11Normal3bv(gl11Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl11Normal3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl11Normal3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl11Normal3iv(gl11Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl11Normal3sv(gl11Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl11NormalPointer(gl11Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl11Ortho(gl11Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl11PassThrough(gl11Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl11PixelMapfv(gl11Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl11PixelMapuiv(gl11Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl11PixelMapusv(gl11Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl11PixelStoref(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl11PixelStorei(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl11PixelTransferf(gl11Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl11PixelTransferi(gl11Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl11PixelZoom(gl11Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl11PointSize(gl11Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl11PolygonMode(gl11Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl11PolygonStipple(gl11Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl11PrioritizeTextures(gl11Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl11PushAttrib(gl11Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl11PopAttrib(gl11Context* glc) {
    return glc->fnPopAttrib();
}

void gl11PushClientAttrib(gl11Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl11PopClientAttrib(gl11Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl11PushMatrix(gl11Context* glc) {
    return glc->fnPushMatrix();
}

void gl11PopMatrix(gl11Context* glc) {
    return glc->fnPopMatrix();
}

void gl11PushName(gl11Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl11PopName(gl11Context* glc) {
    return glc->fnPopName();
}

void gl11RasterPos2d(gl11Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl11RasterPos2f(gl11Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl11RasterPos2i(gl11Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl11RasterPos2s(gl11Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl11RasterPos3d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl11RasterPos3f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl11RasterPos3i(gl11Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl11RasterPos3s(gl11Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl11RasterPos4d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl11RasterPos4f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl11RasterPos4i(gl11Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl11RasterPos4s(gl11Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl11RasterPos2dv(gl11Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl11RasterPos2fv(gl11Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl11RasterPos2iv(gl11Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl11RasterPos2sv(gl11Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl11RasterPos3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl11RasterPos3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl11RasterPos3iv(gl11Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl11RasterPos3sv(gl11Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl11RasterPos4dv(gl11Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl11RasterPos4fv(gl11Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl11RasterPos4iv(gl11Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl11RasterPos4sv(gl11Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl11ReadBuffer(gl11Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl11ReadPixels(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl11Rectd(gl11Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl11Rectf(gl11Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl11Recti(gl11Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl11Rects(gl11Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl11Rectdv(gl11Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl11Rectfv(gl11Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl11Rectiv(gl11Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl11Rectsv(gl11Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl11RenderMode(gl11Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl11Rotated(gl11Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl11Rotatef(gl11Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl11Scaled(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl11Scalef(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl11Scissor(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl11SelectBuffer(gl11Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl11ShadeModel(gl11Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl11StencilFunc(gl11Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl11StencilMask(gl11Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl11StencilOp(gl11Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl11TexCoord1d(gl11Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl11TexCoord1f(gl11Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl11TexCoord1i(gl11Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl11TexCoord1s(gl11Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl11TexCoord2d(gl11Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl11TexCoord2f(gl11Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl11TexCoord2i(gl11Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl11TexCoord2s(gl11Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl11TexCoord3d(gl11Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl11TexCoord3f(gl11Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl11TexCoord3i(gl11Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl11TexCoord3s(gl11Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl11TexCoord4d(gl11Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl11TexCoord4f(gl11Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl11TexCoord4i(gl11Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl11TexCoord4s(gl11Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl11TexCoord1dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl11TexCoord1fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl11TexCoord1iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl11TexCoord1sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl11TexCoord2dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl11TexCoord2fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl11TexCoord2iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl11TexCoord2sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl11TexCoord3dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl11TexCoord3fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl11TexCoord3iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl11TexCoord3sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl11TexCoord4dv(gl11Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl11TexCoord4fv(gl11Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl11TexCoord4iv(gl11Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl11TexCoord4sv(gl11Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl11TexCoordPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl11TexEnvf(gl11Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl11TexEnvi(gl11Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl11TexEnvfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl11TexEnviv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl11TexGend(gl11Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl11TexGenf(gl11Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl11TexGeni(gl11Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl11TexGendv(gl11Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl11TexGenfv(gl11Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl11TexGeniv(gl11Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl11TexImage1D(gl11Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl11TexImage2D(gl11Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl11TexImage3DEXT(gl11Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl11TexParameterf(gl11Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl11TexParameteri(gl11Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl11TexParameterfv(gl11Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl11TexParameteriv(gl11Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl11TexSubImage1D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl11TexSubImage2D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl11TexSubImage3DEXT(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl11Translated(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl11Translatef(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl11Vertex2s(gl11Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl11Vertex2i(gl11Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl11Vertex2f(gl11Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl11Vertex2d(gl11Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl11Vertex3s(gl11Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl11Vertex3i(gl11Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl11Vertex3f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl11Vertex3d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl11Vertex4s(gl11Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl11Vertex4i(gl11Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl11Vertex4f(gl11Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl11Vertex4d(gl11Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl11VertexPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl11Viewport(gl11Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

GLboolean gl11AreTexturesResident(gl11Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl11ArrayElement(gl11Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl11DrawArrays(gl11Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl11DrawElements(gl11Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl11GetPointerv(gl11Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl11PolygonOffset(gl11Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl11CopyTexImage1D(gl11Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl11CopyTexImage2D(gl11Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl11CopyTexSubImage1D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl11CopyTexSubImage2D(gl11Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl11BindTexture(gl11Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl11DeleteTextures(gl11Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl11GenTextures(gl11Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl11IsTexture(gl11Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl11ColorPointer(gl11Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

gl11Context* gl11NewContext() {
    gl11Context* glc = calloc(1, sizeof(gl11Context));

    // Preload all procedures
    glc->fnAccum = (gl11PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl11PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl11PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl11PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl11PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl11PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl11PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl11PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl11PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl11PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl11PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl11PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl11PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl11PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl11PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl11PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl11PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl11PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl11PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl11PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl11PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl11PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl11PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl11PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl11PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl11PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl11PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl11PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl11PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl11PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl11PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl11PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl11PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl11PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl11PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl11PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl11PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl11PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl11PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl11PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl11PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl11PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl11PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl11PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl11PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl11PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl11PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl11PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl11PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl11PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl11PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl11PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl11PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl11PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl11PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl11PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl11PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl11PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl11PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl11PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl11PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl11PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl11PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl11PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl11PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl11PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl11PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl11PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl11PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl11PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl11PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl11PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl11PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl11PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl11PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl11PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl11PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl11PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl11PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl11PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl11PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl11PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl11PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl11PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl11PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl11PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl11PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl11PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl11PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl11PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl11PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl11PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl11PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl11PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl11PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl11PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl11PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl11PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl11PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl11PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl11PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl11PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl11PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl11PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl11PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl11PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl11PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl11PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl11PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl11PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl11PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl11PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl11PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl11PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl11PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl11PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl11PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl11PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl11PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl11PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl11PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl11PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl11PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl11PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl11PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl11PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl11PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl11PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl11PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl11PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl11PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl11PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl11PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl11PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl11PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl11PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl11PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl11PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl11PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl11PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl11PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl11PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl11PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl11PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl11PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl11PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl11PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl11PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl11PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl11PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl11PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl11PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl11PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl11PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl11PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl11PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl11PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl11PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl11PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl11PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl11PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl11PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl11PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl11PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl11PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl11PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl11PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl11PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl11PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl11PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl11PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl11PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl11PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl11PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl11PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl11PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl11PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl11PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl11PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl11PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl11PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl11PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl11PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl11PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl11PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl11PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl11POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl11PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl11PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl11PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl11PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl11PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl11PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl11PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl11PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl11PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl11PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl11PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl11PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl11PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl11PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl11PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl11PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl11PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl11PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl11PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl11PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl11PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl11PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl11PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl11PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl11PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl11PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl11PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl11PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl11PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl11PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl11PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl11PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl11PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl11PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl11PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl11PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl11PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl11PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl11PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl11PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl11PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl11PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl11PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl11PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl11PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl11PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl11PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl11PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl11PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl11PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl11PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl11PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl11PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl11PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl11PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl11PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl11PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl11PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl11PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl11PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl11PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl11PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl11PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl11PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl11PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl11PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl11PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl11PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl11PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl11PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl11PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl11PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl11PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl11PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl11PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl11PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl11PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl11PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl11PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl11PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl11PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl11PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl11PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl11PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl11PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl11PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl11PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl11PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl11PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl11PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl11PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl11PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl11PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl11PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl11PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl11PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl11PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl11PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl11PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl11PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl11PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl11PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl11PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl11PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl11PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl11PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl11PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl11PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl11PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl11PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl11PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl11PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl11PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl11PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl11PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl11PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl11PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl11PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl11PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl11PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl11PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl11PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl11PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl11PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl11PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl11PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl11PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl11PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl11PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl11PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl11PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl11PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl11PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl11PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl11PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl11PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl11PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl11PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl11PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl11PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl11PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl11PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl11PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl11PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl11PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl11PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl11PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl11PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl11PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl11PColorPointer)doGetProcAddress("glColorPointer");
    return glc;
}

