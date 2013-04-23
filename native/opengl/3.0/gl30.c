#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl30.h"

#ifdef _WIN32
    HMODULE gl30OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl30OpenGL32 == NULL) {
            gl30OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl30OpenGL32, TEXT(name));
    }
#endif

void gl30Accum(gl30Context* glc, GLenum op, GLfloat value) {
    return glc->fnAccum(op, value);
}

void gl30AlphaFunc(gl30Context* glc, GLenum func, GLclampf ref) {
    return glc->fnAlphaFunc(func, ref);
}

void gl30Begin(gl30Context* glc, GLenum mode) {
    return glc->fnBegin(mode);
}

void gl30End(gl30Context* glc) {
    return glc->fnEnd();
}

void gl30Bitmap(gl30Context* glc, GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, GLubyte* bitmap) {
    return glc->fnBitmap(width, height, xorig, yorig, xmove, ymove, bitmap);
}

void gl30BlendFunc(gl30Context* glc, GLenum sfactor, GLenum dfactor) {
    return glc->fnBlendFunc(sfactor, dfactor);
}

void gl30CallList(gl30Context* glc, GLuint list) {
    return glc->fnCallList(list);
}

void gl30CallLists(gl30Context* glc, GLsizei n, GLenum type, GLvoid* lists) {
    return glc->fnCallLists(n, type, lists);
}

void gl30Clear(gl30Context* glc, GLbitfield mask) {
    return glc->fnClear(mask);
}

void gl30ClearAccum(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnClearAccum(red, green, blue, alpha);
}

void gl30ClearColor(gl30Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnClearColor(red, green, blue, alpha);
}

void gl30ClearDepth(gl30Context* glc, GLclampd depth) {
    return glc->fnClearDepth(depth);
}

void gl30ClearIndex(gl30Context* glc, GLfloat c) {
    return glc->fnClearIndex(c);
}

void gl30ClearStencil(gl30Context* glc, GLint s) {
    return glc->fnClearStencil(s);
}

void gl30ClipPlane(gl30Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnClipPlane(plane, equation);
}

void gl30Color3b(gl30Context* glc, GLbyte red, GLbyte green, GLbyte blue) {
    return glc->fnColor3b(red, green, blue);
}

void gl30Color3d(gl30Context* glc, GLdouble red, GLdouble green, GLdouble blue) {
    return glc->fnColor3d(red, green, blue);
}

void gl30Color3f(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue) {
    return glc->fnColor3f(red, green, blue);
}

void gl30Color3i(gl30Context* glc, GLint red, GLint green, GLint blue) {
    return glc->fnColor3i(red, green, blue);
}

void gl30Color3s(gl30Context* glc, GLshort red, GLshort green, GLshort blue) {
    return glc->fnColor3s(red, green, blue);
}

void gl30Color3ub(gl30Context* glc, GLubyte red, GLubyte green, GLubyte blue) {
    return glc->fnColor3ub(red, green, blue);
}

void gl30Color3ui(gl30Context* glc, GLuint red, GLuint green, GLuint blue) {
    return glc->fnColor3ui(red, green, blue);
}

void gl30Color3us(gl30Context* glc, GLushort red, GLushort green, GLushort blue) {
    return glc->fnColor3us(red, green, blue);
}

void gl30Color4b(gl30Context* glc, GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
    return glc->fnColor4b(red, green, blue, alpha);
}

void gl30Color4d(gl30Context* glc, GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
    return glc->fnColor4d(red, green, blue, alpha);
}

void gl30Color4f(gl30Context* glc, GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
    return glc->fnColor4f(red, green, blue, alpha);
}

void gl30Color4i(gl30Context* glc, GLint red, GLint green, GLint blue, GLint alpha) {
    return glc->fnColor4i(red, green, blue, alpha);
}

void gl30Color4s(gl30Context* glc, GLshort red, GLshort green, GLshort blue, GLshort alpha) {
    return glc->fnColor4s(red, green, blue, alpha);
}

void gl30Color4ub(gl30Context* glc, GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
    return glc->fnColor4ub(red, green, blue, alpha);
}

void gl30Color4ui(gl30Context* glc, GLuint red, GLuint green, GLuint blue, GLuint alpha) {
    return glc->fnColor4ui(red, green, blue, alpha);
}

void gl30Color4us(gl30Context* glc, GLushort red, GLushort green, GLushort blue, GLushort alpha) {
    return glc->fnColor4us(red, green, blue, alpha);
}

void gl30Color3bv(gl30Context* glc, GLbyte* v) {
    return glc->fnColor3bv(v);
}

void gl30Color3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnColor3dv(v);
}

void gl30Color3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnColor3fv(v);
}

void gl30Color3iv(gl30Context* glc, GLint* v) {
    return glc->fnColor3iv(v);
}

void gl30Color3sv(gl30Context* glc, GLshort* v) {
    return glc->fnColor3sv(v);
}

void gl30Color3ubv(gl30Context* glc, GLubyte* v) {
    return glc->fnColor3ubv(v);
}

void gl30Color3uiv(gl30Context* glc, GLuint* v) {
    return glc->fnColor3uiv(v);
}

void gl30Color3usv(gl30Context* glc, GLushort* v) {
    return glc->fnColor3usv(v);
}

void gl30Color4bv(gl30Context* glc, GLbyte* v) {
    return glc->fnColor4bv(v);
}

void gl30Color4dv(gl30Context* glc, GLdouble* v) {
    return glc->fnColor4dv(v);
}

void gl30Color4fv(gl30Context* glc, GLfloat* v) {
    return glc->fnColor4fv(v);
}

void gl30Color4iv(gl30Context* glc, GLint* v) {
    return glc->fnColor4iv(v);
}

void gl30Color4sv(gl30Context* glc, GLshort* v) {
    return glc->fnColor4sv(v);
}

void gl30Color4ubv(gl30Context* glc, GLubyte* v) {
    return glc->fnColor4ubv(v);
}

void gl30Color4uiv(gl30Context* glc, GLuint* v) {
    return glc->fnColor4uiv(v);
}

void gl30Color4usv(gl30Context* glc, GLushort* v) {
    return glc->fnColor4usv(v);
}

void gl30ColorMask(gl30Context* glc, GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
    return glc->fnColorMask(red, green, blue, alpha);
}

void gl30ColorMaterial(gl30Context* glc, GLenum face, GLenum mode) {
    return glc->fnColorMaterial(face, mode);
}

void gl30ColorTable(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorTable(target, internalformat, width, format, type, data);
}

void gl30ColorTableParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnColorTableParameterfv(target, pname, params);
}

void gl30ColorTableParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnColorTableParameteriv(target, pname, params);
}

void gl30ColorSubTable(gl30Context* glc, GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnColorSubTable(target, start, count, format, type, data);
}

void gl30CopyPixels(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl30CullFace(gl30Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
}

void gl30ConvolutionFilter1D(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter1D(target, internalformat, width, format, type, data);
}

void gl30ConvolutionFilter2D(gl30Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnConvolutionFilter2D(target, internalformat, width, height, format, type, data);
}

void gl30ConvolutionParameterf(gl30Context* glc, GLenum target, GLenum pname, GLfloat params) {
    return glc->fnConvolutionParameterf(target, pname, params);
}

void gl30ConvolutionParameteri(gl30Context* glc, GLenum target, GLenum pname, GLint params) {
    return glc->fnConvolutionParameteri(target, pname, params);
}

void gl30CopyColorTable(gl30Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorTable(target, internalformat, x, y, width);
}

void gl30CopyColorSubTable(gl30Context* glc, GLenum target, GLsizei start, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyColorSubTable(target, start, x, y, width);
}

void gl30CopyConvolutionFilter1D(gl30Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyConvolutionFilter1D(target, internalformat, x, y, width);
}

void gl30CopyConvolutionFilter2D(gl30Context* glc, GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyConvolutionFilter2D(target, internalformat, x, y, width, height);
}

void gl30DeleteLists(gl30Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl30DepthFunc(gl30Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
}

void gl30DepthRange(gl30Context* glc, GLclampd zNear, GLclampd zFar) {
    return glc->fnDepthRange(zNear, zFar);
}

void gl30Enable(gl30Context* glc, GLenum cap) {
    return glc->fnEnable(cap);
}

void gl30Disable(gl30Context* glc, GLenum cap) {
    return glc->fnDisable(cap);
}

void gl30DrawBuffer(gl30Context* glc, GLenum mode) {
    return glc->fnDrawBuffer(mode);
}

void gl30DrawPixels(gl30Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
}

void gl30EdgeFlag(gl30Context* glc, GLboolean flag) {
    return glc->fnEdgeFlag(flag);
}

void gl30EdgeFlagv(gl30Context* glc, GLboolean* flag) {
    return glc->fnEdgeFlagv(flag);
}

void gl30EdgeFlagPointer(gl30Context* glc, GLsizei stride, GLvoid* pointer) {
    return glc->fnEdgeFlagPointer(stride, pointer);
}

void gl30EnableClientState(gl30Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl30DisableClientState(gl30Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl30EvalCoord1d(gl30Context* glc, GLdouble u) {
    return glc->fnEvalCoord1d(u);
}

void gl30EvalCoord1f(gl30Context* glc, GLfloat u) {
    return glc->fnEvalCoord1f(u);
}

void gl30EvalCoord2d(gl30Context* glc, GLdouble u, GLdouble v) {
    return glc->fnEvalCoord2d(u, v);
}

void gl30EvalCoord2f(gl30Context* glc, GLfloat u, GLfloat v) {
    return glc->fnEvalCoord2f(u, v);
}

void gl30EvalCoord1dv(gl30Context* glc, GLdouble* u) {
    return glc->fnEvalCoord1dv(u);
}

void gl30EvalCoord1fv(gl30Context* glc, GLfloat* u) {
    return glc->fnEvalCoord1fv(u);
}

void gl30EvalCoord2dv(gl30Context* glc, GLdouble* u) {
    return glc->fnEvalCoord2dv(u);
}

void gl30EvalCoord2fv(gl30Context* glc, GLfloat* u) {
    return glc->fnEvalCoord2fv(u);
}

void gl30EvalMesh1(gl30Context* glc, GLenum mode, GLint i1, GLint i2) {
    return glc->fnEvalMesh1(mode, i1, i2);
}

void gl30EvalMesh2(gl30Context* glc, GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
    return glc->fnEvalMesh2(mode, i1, i2, j1, j2);
}

void gl30EvalPoint1(gl30Context* glc, GLint i) {
    return glc->fnEvalPoint1(i);
}

void gl30EvalPoint2(gl30Context* glc, GLint i, GLint j) {
    return glc->fnEvalPoint2(i, j);
}

void gl30FeedbackBuffer(gl30Context* glc, GLsizei size, GLenum type, GLfloat* buffer) {
    return glc->fnFeedbackBuffer(size, type, buffer);
}

void gl30Finish(gl30Context* glc) {
    return glc->fnFinish();
}

void gl30Flush(gl30Context* glc) {
    return glc->fnFlush();
}

void gl30Fogf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnFogf(pname, param);
}

void gl30Fogi(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnFogi(pname, param);
}

void gl30Fogfv(gl30Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnFogfv(pname, params);
}

void gl30Fogiv(gl30Context* glc, GLenum pname, GLint* params) {
    return glc->fnFogiv(pname, params);
}

void gl30FrontFace(gl30Context* glc, GLenum mode) {
    return glc->fnFrontFace(mode);
}

void gl30Frustum(gl30Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
    return glc->fnFrustum(left, right, bottom, top, zNear, zFar);
}

GLuint gl30GenLists(gl30Context* glc, GLsizei range) {
    return glc->fnGenLists(range);
}

void gl30GetBooleanv(gl30Context* glc, GLenum pname, GLboolean* params) {
    return glc->fnGetBooleanv(pname, params);
}

void gl30GetDoublev(gl30Context* glc, GLenum pname, GLdouble* params) {
    return glc->fnGetDoublev(pname, params);
}

void gl30GetFloatv(gl30Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnGetFloatv(pname, params);
}

void gl30GetIntegerv(gl30Context* glc, GLenum pname, GLint* params) {
    return glc->fnGetIntegerv(pname, params);
}

void gl30GetClipPlane(gl30Context* glc, GLenum plane, GLdouble* equation) {
    return glc->fnGetClipPlane(plane, equation);
}

GLenum gl30GetError(gl30Context* glc) {
    return glc->fnGetError();
}

void gl30GetLightfv(gl30Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnGetLightfv(light, pname, params);
}

void gl30GetLightiv(gl30Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnGetLightiv(light, pname, params);
}

void gl30GetMapdv(gl30Context* glc, GLenum target, GLenum query, GLdouble* v) {
    return glc->fnGetMapdv(target, query, v);
}

void gl30GetMapfv(gl30Context* glc, GLenum target, GLenum query, GLfloat* v) {
    return glc->fnGetMapfv(target, query, v);
}

void gl30GetMapiv(gl30Context* glc, GLenum target, GLenum query, GLint* v) {
    return glc->fnGetMapiv(target, query, v);
}

void gl30GetMaterialfv(gl30Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnGetMaterialfv(face, pname, params);
}

void gl30GetMaterialiv(gl30Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnGetMaterialiv(face, pname, params);
}

void gl30GetPixelMapfv(gl30Context* glc, GLenum map, GLfloat* values) {
    return glc->fnGetPixelMapfv(map, values);
}

void gl30GetPixelMapuiv(gl30Context* glc, GLenum map, GLuint* values) {
    return glc->fnGetPixelMapuiv(map, values);
}

void gl30GetPixelMapusv(gl30Context* glc, GLenum map, GLushort* values) {
    return glc->fnGetPixelMapusv(map, values);
}

void gl30GetPolygonStipple(gl30Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
}

GLubyte* gl30GetString(gl30Context* glc, GLenum name) {
    return glc->fnGetString(name);
}

void gl30GetTexEnvfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexEnvfv(target, pname, params);
}

void gl30GetTexEnviv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexEnviv(target, pname, params);
}

void gl30GetTexGendv(gl30Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnGetTexGendv(coord, pname, params);
}

void gl30GetTexGenfv(gl30Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnGetTexGenfv(coord, pname, params);
}

void gl30GetTexGeniv(gl30Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnGetTexGeniv(coord, pname, params);
}

void gl30GetTexImage(gl30Context* glc, GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnGetTexImage(target, level, format, type, pixels);
}

void gl30GetTexLevelParameterfv(gl30Context* glc, GLenum target, GLint level, GLenum pname, GLfloat* params) {
    return glc->fnGetTexLevelParameterfv(target, level, pname, params);
}

void gl30GetTexLevelParameteriv(gl30Context* glc, GLenum target, GLint level, GLenum pname, GLint* params) {
    return glc->fnGetTexLevelParameteriv(target, level, pname, params);
}

void gl30GetTexParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetTexParameterfv(target, pname, params);
}

void gl30GetTexParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetTexParameteriv(target, pname, params);
}

void gl30Hint(gl30Context* glc, GLenum target, GLenum mode) {
    return glc->fnHint(target, mode);
}

void gl30Indexd(gl30Context* glc, GLdouble c) {
    return glc->fnIndexd(c);
}

void gl30Indexf(gl30Context* glc, GLfloat c) {
    return glc->fnIndexf(c);
}

void gl30Indexi(gl30Context* glc, GLint c) {
    return glc->fnIndexi(c);
}

void gl30Indexs(gl30Context* glc, GLshort c) {
    return glc->fnIndexs(c);
}

void gl30Indexub(gl30Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl30Indexdv(gl30Context* glc, GLdouble* c) {
    return glc->fnIndexdv(c);
}

void gl30Indexfv(gl30Context* glc, GLfloat* c) {
    return glc->fnIndexfv(c);
}

void gl30Indexiv(gl30Context* glc, GLint* c) {
    return glc->fnIndexiv(c);
}

void gl30Indexsv(gl30Context* glc, GLshort* c) {
    return glc->fnIndexsv(c);
}

void gl30Indexubv(gl30Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl30IndexMask(gl30Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl30IndexPointer(gl30Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl30InitNames(gl30Context* glc) {
    return glc->fnInitNames();
}

void gl30InterleavedArrays(gl30Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl30IsEnabled(gl30Context* glc, GLenum cap) {
    return glc->fnIsEnabled(cap);
}

GLboolean gl30IsList(gl30Context* glc, GLuint list) {
    return glc->fnIsList(list);
}

void gl30Lightf(gl30Context* glc, GLenum light, GLenum pname, GLfloat param) {
    return glc->fnLightf(light, pname, param);
}

void gl30Lighti(gl30Context* glc, GLenum light, GLenum pname, GLint param) {
    return glc->fnLighti(light, pname, param);
}

void gl30Lightfv(gl30Context* glc, GLenum light, GLenum pname, GLfloat* params) {
    return glc->fnLightfv(light, pname, params);
}

void gl30Lightiv(gl30Context* glc, GLenum light, GLenum pname, GLint* params) {
    return glc->fnLightiv(light, pname, params);
}

void gl30LightModelf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnLightModelf(pname, param);
}

void gl30LightModeli(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnLightModeli(pname, param);
}

void gl30LightModelfv(gl30Context* glc, GLenum pname, GLfloat* params) {
    return glc->fnLightModelfv(pname, params);
}

void gl30LightModeliv(gl30Context* glc, GLenum pname, GLint* params) {
    return glc->fnLightModeliv(pname, params);
}

void gl30LineStipple(gl30Context* glc, GLint factor, GLushort pattern) {
    return glc->fnLineStipple(factor, pattern);
}

void gl30LineWidth(gl30Context* glc, GLfloat width) {
    return glc->fnLineWidth(width);
}

void gl30ListBase(gl30Context* glc, GLuint base) {
    return glc->fnListBase(base);
}

void gl30LoadIdentity(gl30Context* glc) {
    return glc->fnLoadIdentity();
}

void gl30LoadMatrixd(gl30Context* glc, GLdouble* m) {
    return glc->fnLoadMatrixd(m);
}

void gl30LoadMatrixf(gl30Context* glc, GLfloat* m) {
    return glc->fnLoadMatrixf(m);
}

void gl30LoadName(gl30Context* glc, GLuint name) {
    return glc->fnLoadName(name);
}

void gl30LogicOp(gl30Context* glc, GLenum opcode) {
    return glc->fnLogicOp(opcode);
}

void gl30Map1d(gl30Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, GLdouble* points) {
    return glc->fnMap1d(target, u1, u2, stride, order, points);
}

void gl30Map1f(gl30Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, GLfloat* points) {
    return glc->fnMap1f(target, u1, u2, stride, order, points);
}

void gl30Map2d(gl30Context* glc, GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, GLdouble* points) {
    return glc->fnMap2d(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl30Map2f(gl30Context* glc, GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, GLfloat* points) {
    return glc->fnMap2f(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}

void gl30MapGrid1d(gl30Context* glc, GLint un, GLdouble u1, GLdouble u2) {
    return glc->fnMapGrid1d(un, u1, u2);
}

void gl30MapGrid1f(gl30Context* glc, GLint un, GLfloat u1, GLfloat u2) {
    return glc->fnMapGrid1f(un, u1, u2);
}

void gl30MapGrid2d(gl30Context* glc, GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
    return glc->fnMapGrid2d(un, u1, u2, vn, v1, v2);
}

void gl30MapGrid2f(gl30Context* glc, GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
    return glc->fnMapGrid2f(un, u1, u2, vn, v1, v2);
}

void gl30Materialf(gl30Context* glc, GLenum face, GLenum pname, GLfloat param) {
    return glc->fnMaterialf(face, pname, param);
}

void gl30Materiali(gl30Context* glc, GLenum face, GLenum pname, GLint param) {
    return glc->fnMateriali(face, pname, param);
}

void gl30Materialfv(gl30Context* glc, GLenum face, GLenum pname, GLfloat* params) {
    return glc->fnMaterialfv(face, pname, params);
}

void gl30Materialiv(gl30Context* glc, GLenum face, GLenum pname, GLint* params) {
    return glc->fnMaterialiv(face, pname, params);
}

void gl30MatrixMode(gl30Context* glc, GLenum mode) {
    return glc->fnMatrixMode(mode);
}

void gl30MultMatrixd(gl30Context* glc, GLdouble* m) {
    return glc->fnMultMatrixd(m);
}

void gl30MultMatrixf(gl30Context* glc, GLfloat* m) {
    return glc->fnMultMatrixf(m);
}

void gl30NewList(gl30Context* glc, GLuint list, GLenum mode) {
    return glc->fnNewList(list, mode);
}

void gl30EndList(gl30Context* glc) {
    return glc->fnEndList();
}

void gl30Normal3b(gl30Context* glc, GLbyte nx, GLbyte ny, GLbyte nz) {
    return glc->fnNormal3b(nx, ny, nz);
}

void gl30Normal3d(gl30Context* glc, GLdouble nx, GLdouble ny, GLdouble nz) {
    return glc->fnNormal3d(nx, ny, nz);
}

void gl30Normal3f(gl30Context* glc, GLfloat nx, GLfloat ny, GLfloat nz) {
    return glc->fnNormal3f(nx, ny, nz);
}

void gl30Normal3i(gl30Context* glc, GLint nx, GLint ny, GLint nz) {
    return glc->fnNormal3i(nx, ny, nz);
}

void gl30Normal3s(gl30Context* glc, GLshort nx, GLshort ny, GLshort nz) {
    return glc->fnNormal3s(nx, ny, nz);
}

void gl30Normal3bv(gl30Context* glc, GLbyte* v) {
    return glc->fnNormal3bv(v);
}

void gl30Normal3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnNormal3dv(v);
}

void gl30Normal3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnNormal3fv(v);
}

void gl30Normal3iv(gl30Context* glc, GLint* v) {
    return glc->fnNormal3iv(v);
}

void gl30Normal3sv(gl30Context* glc, GLshort* v) {
    return glc->fnNormal3sv(v);
}

void gl30NormalPointer(gl30Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl30Ortho(gl30Context* glc, GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zfar) {
    return glc->fnOrtho(left, right, bottom, top, zNear, zfar);
}

void gl30PassThrough(gl30Context* glc, GLfloat token) {
    return glc->fnPassThrough(token);
}

void gl30PixelMapfv(gl30Context* glc, GLenum map, GLsizei mapsize, GLfloat* values) {
    return glc->fnPixelMapfv(map, mapsize, values);
}

void gl30PixelMapuiv(gl30Context* glc, GLenum map, GLsizei mapsize, GLuint* values) {
    return glc->fnPixelMapuiv(map, mapsize, values);
}

void gl30PixelMapusv(gl30Context* glc, GLenum map, GLsizei mapsize, GLushort* values) {
    return glc->fnPixelMapusv(map, mapsize, values);
}

void gl30PixelStoref(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelStoref(pname, param);
}

void gl30PixelStorei(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelStorei(pname, param);
}

void gl30PixelTransferf(gl30Context* glc, GLenum pname, GLfloat param) {
    return glc->fnPixelTransferf(pname, param);
}

void gl30PixelTransferi(gl30Context* glc, GLenum pname, GLint param) {
    return glc->fnPixelTransferi(pname, param);
}

void gl30PixelZoom(gl30Context* glc, GLfloat xfactor, GLfloat yfactor) {
    return glc->fnPixelZoom(xfactor, yfactor);
}

void gl30PointSize(gl30Context* glc, GLfloat size) {
    return glc->fnPointSize(size);
}

void gl30PolygonMode(gl30Context* glc, GLenum face, GLenum mode) {
    return glc->fnPolygonMode(face, mode);
}

void gl30PolygonStipple(gl30Context* glc, GLubyte* mask) {
    return glc->fnPolygonStipple(mask);
}

void gl30PrioritizeTextures(gl30Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl30PushAttrib(gl30Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl30PopAttrib(gl30Context* glc) {
    return glc->fnPopAttrib();
}

void gl30PushClientAttrib(gl30Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl30PopClientAttrib(gl30Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl30PushMatrix(gl30Context* glc) {
    return glc->fnPushMatrix();
}

void gl30PopMatrix(gl30Context* glc) {
    return glc->fnPopMatrix();
}

void gl30PushName(gl30Context* glc, GLuint name) {
    return glc->fnPushName(name);
}

void gl30PopName(gl30Context* glc) {
    return glc->fnPopName();
}

void gl30RasterPos2d(gl30Context* glc, GLdouble x, GLdouble y) {
    return glc->fnRasterPos2d(x, y);
}

void gl30RasterPos2f(gl30Context* glc, GLfloat x, GLfloat y) {
    return glc->fnRasterPos2f(x, y);
}

void gl30RasterPos2i(gl30Context* glc, GLint x, GLint y) {
    return glc->fnRasterPos2i(x, y);
}

void gl30RasterPos2s(gl30Context* glc, GLshort x, GLshort y) {
    return glc->fnRasterPos2s(x, y);
}

void gl30RasterPos3d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRasterPos3d(x, y, z);
}

void gl30RasterPos3f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRasterPos3f(x, y, z);
}

void gl30RasterPos3i(gl30Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnRasterPos3i(x, y, z);
}

void gl30RasterPos3s(gl30Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnRasterPos3s(x, y, z);
}

void gl30RasterPos4d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnRasterPos4d(x, y, z, w);
}

void gl30RasterPos4f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnRasterPos4f(x, y, z, w);
}

void gl30RasterPos4i(gl30Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnRasterPos4i(x, y, z, w);
}

void gl30RasterPos4s(gl30Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnRasterPos4s(x, y, z, w);
}

void gl30RasterPos2dv(gl30Context* glc, GLdouble* v) {
    return glc->fnRasterPos2dv(v);
}

void gl30RasterPos2fv(gl30Context* glc, GLfloat* v) {
    return glc->fnRasterPos2fv(v);
}

void gl30RasterPos2iv(gl30Context* glc, GLint* v) {
    return glc->fnRasterPos2iv(v);
}

void gl30RasterPos2sv(gl30Context* glc, GLshort* v) {
    return glc->fnRasterPos2sv(v);
}

void gl30RasterPos3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnRasterPos3dv(v);
}

void gl30RasterPos3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnRasterPos3fv(v);
}

void gl30RasterPos3iv(gl30Context* glc, GLint* v) {
    return glc->fnRasterPos3iv(v);
}

void gl30RasterPos3sv(gl30Context* glc, GLshort* v) {
    return glc->fnRasterPos3sv(v);
}

void gl30RasterPos4dv(gl30Context* glc, GLdouble* v) {
    return glc->fnRasterPos4dv(v);
}

void gl30RasterPos4fv(gl30Context* glc, GLfloat* v) {
    return glc->fnRasterPos4fv(v);
}

void gl30RasterPos4iv(gl30Context* glc, GLint* v) {
    return glc->fnRasterPos4iv(v);
}

void gl30RasterPos4sv(gl30Context* glc, GLshort* v) {
    return glc->fnRasterPos4sv(v);
}

void gl30ReadBuffer(gl30Context* glc, GLenum mode) {
    return glc->fnReadBuffer(mode);
}

void gl30ReadPixels(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnReadPixels(x, y, width, height, format, type, pixels);
}

void gl30Rectd(gl30Context* glc, GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
    return glc->fnRectd(x1, y1, x2, y2);
}

void gl30Rectf(gl30Context* glc, GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
    return glc->fnRectf(x1, y1, x2, y2);
}

void gl30Recti(gl30Context* glc, GLint x1, GLint y1, GLint x2, GLint y2) {
    return glc->fnRecti(x1, y1, x2, y2);
}

void gl30Rects(gl30Context* glc, GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
    return glc->fnRects(x1, y1, x2, y2);
}

void gl30Rectdv(gl30Context* glc, GLdouble* v1, GLdouble* v2) {
    return glc->fnRectdv(v1, v2);
}

void gl30Rectfv(gl30Context* glc, GLfloat* v1, GLfloat* v2) {
    return glc->fnRectfv(v1, v2);
}

void gl30Rectiv(gl30Context* glc, GLint* v1, GLint* v2) {
    return glc->fnRectiv(v1, v2);
}

void gl30Rectsv(gl30Context* glc, GLshort* v1, GLshort* v2) {
    return glc->fnRectsv(v1, v2);
}

GLint gl30RenderMode(gl30Context* glc, GLenum mode) {
    return glc->fnRenderMode(mode);
}

void gl30Rotated(gl30Context* glc, GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnRotated(angle, x, y, z);
}

void gl30Rotatef(gl30Context* glc, GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnRotatef(angle, x, y, z);
}

void gl30Scaled(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnScaled(x, y, z);
}

void gl30Scalef(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnScalef(x, y, z);
}

void gl30Scissor(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnScissor(x, y, width, height);
}

void gl30SelectBuffer(gl30Context* glc, GLsizei size, GLuint* buffer) {
    return glc->fnSelectBuffer(size, buffer);
}

void gl30ShadeModel(gl30Context* glc, GLenum mode) {
    return glc->fnShadeModel(mode);
}

void gl30StencilFunc(gl30Context* glc, GLenum func, GLint ref, GLuint mask) {
    return glc->fnStencilFunc(func, ref, mask);
}

void gl30StencilMask(gl30Context* glc, GLuint mask) {
    return glc->fnStencilMask(mask);
}

void gl30StencilOp(gl30Context* glc, GLenum fail, GLenum zfail, GLenum zpass) {
    return glc->fnStencilOp(fail, zfail, zpass);
}

void gl30TexCoord1d(gl30Context* glc, GLdouble s) {
    return glc->fnTexCoord1d(s);
}

void gl30TexCoord1f(gl30Context* glc, GLfloat s) {
    return glc->fnTexCoord1f(s);
}

void gl30TexCoord1i(gl30Context* glc, GLint s) {
    return glc->fnTexCoord1i(s);
}

void gl30TexCoord1s(gl30Context* glc, GLshort s) {
    return glc->fnTexCoord1s(s);
}

void gl30TexCoord2d(gl30Context* glc, GLdouble s, GLdouble t) {
    return glc->fnTexCoord2d(s, t);
}

void gl30TexCoord2f(gl30Context* glc, GLfloat s, GLfloat t) {
    return glc->fnTexCoord2f(s, t);
}

void gl30TexCoord2i(gl30Context* glc, GLint s, GLint t) {
    return glc->fnTexCoord2i(s, t);
}

void gl30TexCoord2s(gl30Context* glc, GLshort s, GLshort t) {
    return glc->fnTexCoord2s(s, t);
}

void gl30TexCoord3d(gl30Context* glc, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnTexCoord3d(s, t, r);
}

void gl30TexCoord3f(gl30Context* glc, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnTexCoord3f(s, t, r);
}

void gl30TexCoord3i(gl30Context* glc, GLint s, GLint t, GLint r) {
    return glc->fnTexCoord3i(s, t, r);
}

void gl30TexCoord3s(gl30Context* glc, GLshort s, GLshort t, GLshort r) {
    return glc->fnTexCoord3s(s, t, r);
}

void gl30TexCoord4d(gl30Context* glc, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnTexCoord4d(s, t, r, q);
}

void gl30TexCoord4f(gl30Context* glc, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnTexCoord4f(s, t, r, q);
}

void gl30TexCoord4i(gl30Context* glc, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnTexCoord4i(s, t, r, q);
}

void gl30TexCoord4s(gl30Context* glc, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnTexCoord4s(s, t, r, q);
}

void gl30TexCoord1dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord1dv(v);
}

void gl30TexCoord1fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord1fv(v);
}

void gl30TexCoord1iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord1iv(v);
}

void gl30TexCoord1sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord1sv(v);
}

void gl30TexCoord2dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord2dv(v);
}

void gl30TexCoord2fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord2fv(v);
}

void gl30TexCoord2iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord2iv(v);
}

void gl30TexCoord2sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord2sv(v);
}

void gl30TexCoord3dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord3dv(v);
}

void gl30TexCoord3fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord3fv(v);
}

void gl30TexCoord3iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord3iv(v);
}

void gl30TexCoord3sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord3sv(v);
}

void gl30TexCoord4dv(gl30Context* glc, GLdouble* v) {
    return glc->fnTexCoord4dv(v);
}

void gl30TexCoord4fv(gl30Context* glc, GLfloat* v) {
    return glc->fnTexCoord4fv(v);
}

void gl30TexCoord4iv(gl30Context* glc, GLint* v) {
    return glc->fnTexCoord4iv(v);
}

void gl30TexCoord4sv(gl30Context* glc, GLshort* v) {
    return glc->fnTexCoord4sv(v);
}

void gl30TexCoordPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl30TexEnvf(gl30Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexEnvf(target, pname, param);
}

void gl30TexEnvi(gl30Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexEnvi(target, pname, param);
}

void gl30TexEnvfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexEnvfv(target, pname, params);
}

void gl30TexEnviv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexEnviv(target, pname, params);
}

void gl30TexGend(gl30Context* glc, GLenum coord, GLenum pname, GLdouble param) {
    return glc->fnTexGend(coord, pname, param);
}

void gl30TexGenf(gl30Context* glc, GLenum coord, GLenum pname, GLfloat param) {
    return glc->fnTexGenf(coord, pname, param);
}

void gl30TexGeni(gl30Context* glc, GLenum coord, GLenum pname, GLint param) {
    return glc->fnTexGeni(coord, pname, param);
}

void gl30TexGendv(gl30Context* glc, GLenum coord, GLenum pname, GLdouble* params) {
    return glc->fnTexGendv(coord, pname, params);
}

void gl30TexGenfv(gl30Context* glc, GLenum coord, GLenum pname, GLfloat* params) {
    return glc->fnTexGenfv(coord, pname, params);
}

void gl30TexGeniv(gl30Context* glc, GLenum coord, GLenum pname, GLint* params) {
    return glc->fnTexGeniv(coord, pname, params);
}

void gl30TexImage1D(gl30Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage1D(target, level, internalformat, width, border, format, type, pixels);
}

void gl30TexImage2D(gl30Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage2D(target, level, internalformat, width, height, border, format, type, pixels);
}

void gl30TexImage3DEXT(gl30Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl30TexParameterf(gl30Context* glc, GLenum target, GLenum pname, GLfloat param) {
    return glc->fnTexParameterf(target, pname, param);
}

void gl30TexParameteri(gl30Context* glc, GLenum target, GLenum pname, GLint param) {
    return glc->fnTexParameteri(target, pname, param);
}

void gl30TexParameterfv(gl30Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnTexParameterfv(target, pname, params);
}

void gl30TexParameteriv(gl30Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnTexParameteriv(target, pname, params);
}

void gl30TexSubImage1D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl30TexSubImage2D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl30TexSubImage3DEXT(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

void gl30Translated(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnTranslated(x, y, z);
}

void gl30Translatef(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnTranslatef(x, y, z);
}

void gl30Vertex2s(gl30Context* glc, GLshort x, GLshort y) {
    return glc->fnVertex2s(x, y);
}

void gl30Vertex2i(gl30Context* glc, GLint x, GLint y) {
    return glc->fnVertex2i(x, y);
}

void gl30Vertex2f(gl30Context* glc, GLfloat x, GLfloat y) {
    return glc->fnVertex2f(x, y);
}

void gl30Vertex2d(gl30Context* glc, GLdouble x, GLdouble y) {
    return glc->fnVertex2d(x, y);
}

void gl30Vertex3s(gl30Context* glc, GLshort x, GLshort y, GLshort z) {
    return glc->fnVertex3s(x, y, z);
}

void gl30Vertex3i(gl30Context* glc, GLint x, GLint y, GLint z) {
    return glc->fnVertex3i(x, y, z);
}

void gl30Vertex3f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z) {
    return glc->fnVertex3f(x, y, z);
}

void gl30Vertex3d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z) {
    return glc->fnVertex3d(x, y, z);
}

void gl30Vertex4s(gl30Context* glc, GLshort x, GLshort y, GLshort z, GLshort w) {
    return glc->fnVertex4s(x, y, z, w);
}

void gl30Vertex4i(gl30Context* glc, GLint x, GLint y, GLint z, GLint w) {
    return glc->fnVertex4i(x, y, z, w);
}

void gl30Vertex4f(gl30Context* glc, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
    return glc->fnVertex4f(x, y, z, w);
}

void gl30Vertex4d(gl30Context* glc, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
    return glc->fnVertex4d(x, y, z, w);
}

void gl30VertexPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl30Viewport(gl30Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

GLboolean gl30AreTexturesResident(gl30Context* glc, GLsizei n, GLuint* textures, GLboolean* residences) {
    return glc->fnAreTexturesResident(n, textures, residences);
}

void gl30ArrayElement(gl30Context* glc, GLint i) {
    return glc->fnArrayElement(i);
}

void gl30DrawArrays(gl30Context* glc, GLenum mode, GLint first, GLsizei count) {
    return glc->fnDrawArrays(mode, first, count);
}

void gl30DrawElements(gl30Context* glc, GLenum mode, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawElements(mode, count, type, indices);
}

void gl30GetPointerv(gl30Context* glc, GLenum pname, GLvoid* params) {
    return glc->fnGetPointerv(pname, params);
}

void gl30PolygonOffset(gl30Context* glc, GLfloat factor, GLfloat units) {
    return glc->fnPolygonOffset(factor, units);
}

void gl30CopyTexImage1D(gl30Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border) {
    return glc->fnCopyTexImage1D(target, level, internalFormat, x, y, width, border);
}

void gl30CopyTexImage2D(gl30Context* glc, GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
    return glc->fnCopyTexImage2D(target, level, internalFormat, x, y, width, height, border);
}

void gl30CopyTexSubImage1D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
    return glc->fnCopyTexSubImage1D(target, level, xoffset, x, y, width);
}

void gl30CopyTexSubImage2D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height);
}

void gl30BindTexture(gl30Context* glc, GLenum target, GLuint texture) {
    return glc->fnBindTexture(target, texture);
}

void gl30DeleteTextures(gl30Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnDeleteTextures(n, textures);
}

void gl30GenTextures(gl30Context* glc, GLsizei n, GLuint* textures) {
    return glc->fnGenTextures(n, textures);
}

GLboolean gl30IsTexture(gl30Context* glc, GLuint texture) {
    return glc->fnIsTexture(texture);
}

void gl30ColorPointer(gl30Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnColorPointer(size, type, stride, pointer);
}

void gl30BlendColorEXT(gl30Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl30BlendEquation(gl30Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl30CopyTexSubImage3D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl30ActiveTexture(gl30Context* glc, GLenum texture) {
    return glc->fnActiveTexture(texture);
}

void gl30ClientActiveTexture(gl30Context* glc, GLenum texture) {
    return glc->fnClientActiveTexture(texture);
}

void gl30CompressedTexImage1D(gl30Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage1D(target, level, internalformat, width, border, imageSize, data);
}

void gl30CompressedTexImage2D(gl30Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data);
}

void gl30CompressedTexImage3D(gl30Context* glc, GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data);
}

void gl30CompressedTexSubImage1D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data);
}

void gl30CompressedTexSubImage2D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}

void gl30CompressedTexSubImage3D(gl30Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {
    return glc->fnCompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}

void gl30BlendFuncSeparate(gl30Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
}

gl30Context* gl30NewContext() {
    gl30Context* glc = calloc(1, sizeof(gl30Context));

    // Preload all procedures
    glc->fnAccum = (gl30PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl30PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl30PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl30PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl30PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl30PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl30PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl30PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl30PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl30PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl30PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl30PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl30PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl30PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl30PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl30PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl30PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl30PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl30PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl30PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl30PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl30PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl30PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl30PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl30PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl30PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl30PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl30PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl30PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl30PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl30PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl30PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl30PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl30PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl30PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl30PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl30PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl30PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl30PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl30PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl30PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl30PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl30PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl30PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl30PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl30PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl30PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl30PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl30PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl30PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl30PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl30PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl30PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl30PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl30PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl30PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl30PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl30PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl30PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl30PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl30PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl30PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl30PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl30PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl30PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl30PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl30PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl30PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl30PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl30PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl30PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl30PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl30PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl30PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl30PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl30PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl30PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl30PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl30PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl30PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl30PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl30PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl30PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl30PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl30PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl30PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl30PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl30PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl30PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl30PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl30PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl30PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl30PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl30PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl30PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl30PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl30PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl30PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl30PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl30PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl30PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl30PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl30PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl30PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl30PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl30PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl30PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl30PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl30PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl30PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl30PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl30PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl30PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl30PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl30PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl30PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl30PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl30PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl30PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl30PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl30PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl30PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl30PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl30PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl30PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl30PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl30PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl30PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl30PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl30PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl30PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl30PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl30PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl30PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl30PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl30PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl30PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl30PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl30PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl30PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl30PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl30PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl30PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl30PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl30PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl30PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl30PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl30PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl30PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl30PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl30PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl30PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl30PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl30PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl30PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl30PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl30PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl30PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl30PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl30PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl30PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl30PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl30PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl30PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl30PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl30PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl30PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl30PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl30PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl30PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl30PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl30PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl30PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl30PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl30PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl30PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl30PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl30PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl30PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl30PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl30PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl30PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl30PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl30PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl30PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl30PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl30POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl30PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl30PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl30PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl30PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl30PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl30PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl30PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl30PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl30PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl30PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl30PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl30PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl30PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl30PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl30PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl30PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl30PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl30PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl30PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl30PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl30PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl30PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl30PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl30PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl30PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl30PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl30PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl30PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl30PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl30PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl30PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl30PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl30PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl30PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl30PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl30PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl30PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl30PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl30PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl30PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl30PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl30PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl30PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl30PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl30PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl30PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl30PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl30PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl30PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl30PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl30PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl30PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl30PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl30PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl30PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl30PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl30PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl30PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl30PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl30PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl30PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl30PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl30PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl30PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl30PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl30PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl30PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl30PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl30PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl30PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl30PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl30PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl30PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl30PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl30PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl30PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl30PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl30PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl30PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl30PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl30PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl30PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl30PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl30PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl30PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl30PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl30PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl30PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl30PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl30PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl30PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl30PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl30PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl30PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl30PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl30PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl30PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl30PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl30PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl30PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl30PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl30PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl30PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl30PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl30PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl30PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl30PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl30PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl30PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl30PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl30PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl30PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl30PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl30PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl30PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl30PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl30PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl30PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl30PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl30PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl30PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl30PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl30PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl30PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl30PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl30PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl30PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl30PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl30PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl30PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl30PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl30PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl30PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl30PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl30PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl30PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl30PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl30PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl30PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl30PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl30PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl30PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl30PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl30PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl30PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl30PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl30PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl30PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl30PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl30PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl30PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl30PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl30PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    glc->fnActiveTexture = (gl30PActiveTexture)wglGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl30PClientActiveTexture)wglGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl30PCompressedTexImage1D)wglGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl30PCompressedTexImage2D)wglGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl30PCompressedTexImage3D)wglGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl30PCompressedTexSubImage1D)wglGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl30PCompressedTexSubImage2D)wglGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl30PCompressedTexSubImage3D)wglGetProcAddress("glCompressedTexSubImage3D");
    glc->fnBlendFuncSeparate = (gl30PBlendFuncSeparate)wglGetProcAddress("glBlendFuncSeparate");
    return glc;
}

