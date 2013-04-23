#include <stdbool.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <windows.h>
#endif
#include "gl20.h"

#ifdef _WIN32
    HMODULE gl20OpenGL32;
    void* doGetProcAddress(char* name) {
        if(gl20OpenGL32 == NULL) {
            gl20OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
        }
        return GetProcAddress(gl20OpenGL32, TEXT(name));
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

void gl20CopyPixels(gl20Context* glc, GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
    return glc->fnCopyPixels(x, y, width, height, type);
}

void gl20CullFace(gl20Context* glc, GLenum mode) {
    return glc->fnCullFace(mode);
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

void gl20DeleteLists(gl20Context* glc, GLuint list, GLsizei range) {
    return glc->fnDeleteLists(list, range);
}

void gl20DepthFunc(gl20Context* glc, GLenum func) {
    return glc->fnDepthFunc(func);
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

void gl20DrawPixels(gl20Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnDrawPixels(width, height, format, type, pixels);
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

void gl20EnableClientState(gl20Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl20DisableClientState(gl20Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
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

void gl20GetPolygonStipple(gl20Context* glc, GLubyte* mask) {
    return glc->fnGetPolygonStipple(mask);
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

void gl20Indexub(gl20Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
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

void gl20Indexubv(gl20Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
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

void gl20InterleavedArrays(gl20Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
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

void gl20NormalPointer(gl20Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
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

void gl20PrioritizeTextures(gl20Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl20PushAttrib(gl20Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl20PopAttrib(gl20Context* glc) {
    return glc->fnPopAttrib();
}

void gl20PushClientAttrib(gl20Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl20PopClientAttrib(gl20Context* glc) {
    return glc->fnPopClientAttrib();
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

void gl20TexCoordPointer(gl20Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
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

void gl20TexImage3DEXT(gl20Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3DEXT(target, level, internalformat, width, height, depth, border, format, type, pixels);
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

void gl20TexSubImage1D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl20TexSubImage2D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl20TexSubImage3DEXT(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3DEXT(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
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

void gl20VertexPointer(gl20Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl20Viewport(gl20Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
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

void gl20BlendColorEXT(gl20Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColorEXT(red, green, blue, alpha);
}

void gl20BlendEquation(gl20Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl20CopyTexSubImage3D(gl20Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
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

void gl20BlendFuncSeparate(gl20Context* glc, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {
    return glc->fnBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha);
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

gl20Context* gl20NewContext() {
    gl20Context* glc = calloc(1, sizeof(gl20Context));

    // Preload all procedures
    glc->fnAccum = (gl20PAccum)doGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl20PAlphaFunc)doGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl20PBegin)doGetProcAddress("glBegin");
    glc->fnEnd = (gl20PEnd)doGetProcAddress("glEnd");
    glc->fnBitmap = (gl20PBitmap)doGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl20PBlendFunc)doGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl20PCallList)doGetProcAddress("glCallList");
    glc->fnCallLists = (gl20PCallLists)doGetProcAddress("glCallLists");
    glc->fnClear = (gl20PClear)doGetProcAddress("glClear");
    glc->fnClearAccum = (gl20PClearAccum)doGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl20PClearColor)doGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl20PClearDepth)doGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl20PClearIndex)doGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl20PClearStencil)doGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl20PClipPlane)doGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl20PColor3b)doGetProcAddress("glColor3b");
    glc->fnColor3d = (gl20PColor3d)doGetProcAddress("glColor3d");
    glc->fnColor3f = (gl20PColor3f)doGetProcAddress("glColor3f");
    glc->fnColor3i = (gl20PColor3i)doGetProcAddress("glColor3i");
    glc->fnColor3s = (gl20PColor3s)doGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl20PColor3ub)doGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl20PColor3ui)doGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl20PColor3us)doGetProcAddress("glColor3us");
    glc->fnColor4b = (gl20PColor4b)doGetProcAddress("glColor4b");
    glc->fnColor4d = (gl20PColor4d)doGetProcAddress("glColor4d");
    glc->fnColor4f = (gl20PColor4f)doGetProcAddress("glColor4f");
    glc->fnColor4i = (gl20PColor4i)doGetProcAddress("glColor4i");
    glc->fnColor4s = (gl20PColor4s)doGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl20PColor4ub)doGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl20PColor4ui)doGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl20PColor4us)doGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl20PColor3bv)doGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl20PColor3dv)doGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl20PColor3fv)doGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl20PColor3iv)doGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl20PColor3sv)doGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl20PColor3ubv)doGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl20PColor3uiv)doGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl20PColor3usv)doGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl20PColor4bv)doGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl20PColor4dv)doGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl20PColor4fv)doGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl20PColor4iv)doGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl20PColor4sv)doGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl20PColor4ubv)doGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl20PColor4uiv)doGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl20PColor4usv)doGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl20PColorMask)doGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl20PColorMaterial)doGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl20PColorTable)wglGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl20PColorTableParameterfv)wglGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl20PColorTableParameteriv)wglGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl20PColorSubTable)wglGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl20PCopyPixels)doGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl20PCullFace)doGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl20PConvolutionFilter1D)wglGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl20PConvolutionFilter2D)wglGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl20PConvolutionParameterf)wglGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl20PConvolutionParameteri)wglGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl20PCopyColorTable)wglGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl20PCopyColorSubTable)wglGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl20PCopyConvolutionFilter1D)wglGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl20PCopyConvolutionFilter2D)wglGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl20PDeleteLists)doGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl20PDepthFunc)doGetProcAddress("glDepthFunc");
    glc->fnDepthRange = (gl20PDepthRange)doGetProcAddress("glDepthRange");
    glc->fnEnable = (gl20PEnable)doGetProcAddress("glEnable");
    glc->fnDisable = (gl20PDisable)doGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl20PDrawBuffer)doGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl20PDrawPixels)doGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl20PEdgeFlag)doGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl20PEdgeFlagv)doGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl20PEdgeFlagPointer)doGetProcAddress("glEdgeFlagPointer");
    glc->fnEnableClientState = (gl20PEnableClientState)doGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl20PDisableClientState)doGetProcAddress("glDisableClientState");
    glc->fnEvalCoord1d = (gl20PEvalCoord1d)doGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl20PEvalCoord1f)doGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl20PEvalCoord2d)doGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl20PEvalCoord2f)doGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl20PEvalCoord1dv)doGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl20PEvalCoord1fv)doGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl20PEvalCoord2dv)doGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl20PEvalCoord2fv)doGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl20PEvalMesh1)doGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl20PEvalMesh2)doGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl20PEvalPoint1)doGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl20PEvalPoint2)doGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl20PFeedbackBuffer)doGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl20PFinish)doGetProcAddress("glFinish");
    glc->fnFlush = (gl20PFlush)doGetProcAddress("glFlush");
    glc->fnFogf = (gl20PFogf)doGetProcAddress("glFogf");
    glc->fnFogi = (gl20PFogi)doGetProcAddress("glFogi");
    glc->fnFogfv = (gl20PFogfv)doGetProcAddress("glFogfv");
    glc->fnFogiv = (gl20PFogiv)doGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl20PFrontFace)doGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl20PFrustum)doGetProcAddress("glFrustum");
    glc->fnGenLists = (gl20PGenLists)doGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl20PGetBooleanv)doGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl20PGetDoublev)doGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl20PGetFloatv)doGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl20PGetIntegerv)doGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl20PGetClipPlane)doGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl20PGetError)doGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl20PGetLightfv)doGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl20PGetLightiv)doGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl20PGetMapdv)doGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl20PGetMapfv)doGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl20PGetMapiv)doGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl20PGetMaterialfv)doGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl20PGetMaterialiv)doGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl20PGetPixelMapfv)doGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl20PGetPixelMapuiv)doGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl20PGetPixelMapusv)doGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl20PGetPolygonStipple)doGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl20PGetString)doGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl20PGetTexEnvfv)doGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl20PGetTexEnviv)doGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl20PGetTexGendv)doGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl20PGetTexGenfv)doGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl20PGetTexGeniv)doGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl20PGetTexImage)doGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl20PGetTexLevelParameterfv)doGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl20PGetTexLevelParameteriv)doGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl20PGetTexParameterfv)doGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl20PGetTexParameteriv)doGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl20PHint)doGetProcAddress("glHint");
    glc->fnIndexd = (gl20PIndexd)doGetProcAddress("glIndexd");
    glc->fnIndexf = (gl20PIndexf)doGetProcAddress("glIndexf");
    glc->fnIndexi = (gl20PIndexi)doGetProcAddress("glIndexi");
    glc->fnIndexs = (gl20PIndexs)doGetProcAddress("glIndexs");
    glc->fnIndexub = (gl20PIndexub)doGetProcAddress("glIndexub");
    glc->fnIndexdv = (gl20PIndexdv)doGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl20PIndexfv)doGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl20PIndexiv)doGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl20PIndexsv)doGetProcAddress("glIndexsv");
    glc->fnIndexubv = (gl20PIndexubv)doGetProcAddress("glIndexubv");
    glc->fnIndexMask = (gl20PIndexMask)doGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl20PIndexPointer)doGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl20PInitNames)doGetProcAddress("glInitNames");
    glc->fnInterleavedArrays = (gl20PInterleavedArrays)doGetProcAddress("glInterleavedArrays");
    glc->fnIsEnabled = (gl20PIsEnabled)doGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl20PIsList)doGetProcAddress("glIsList");
    glc->fnLightf = (gl20PLightf)doGetProcAddress("glLightf");
    glc->fnLighti = (gl20PLighti)doGetProcAddress("glLighti");
    glc->fnLightfv = (gl20PLightfv)doGetProcAddress("glLightfv");
    glc->fnLightiv = (gl20PLightiv)doGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl20PLightModelf)doGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl20PLightModeli)doGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl20PLightModelfv)doGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl20PLightModeliv)doGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl20PLineStipple)doGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl20PLineWidth)doGetProcAddress("glLineWidth");
    glc->fnListBase = (gl20PListBase)doGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl20PLoadIdentity)doGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl20PLoadMatrixd)doGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl20PLoadMatrixf)doGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl20PLoadName)doGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl20PLogicOp)doGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl20PMap1d)doGetProcAddress("glMap1d");
    glc->fnMap1f = (gl20PMap1f)doGetProcAddress("glMap1f");
    glc->fnMap2d = (gl20PMap2d)doGetProcAddress("glMap2d");
    glc->fnMap2f = (gl20PMap2f)doGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl20PMapGrid1d)doGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl20PMapGrid1f)doGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl20PMapGrid2d)doGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl20PMapGrid2f)doGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl20PMaterialf)doGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl20PMateriali)doGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl20PMaterialfv)doGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl20PMaterialiv)doGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl20PMatrixMode)doGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl20PMultMatrixd)doGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl20PMultMatrixf)doGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl20PNewList)doGetProcAddress("glNewList");
    glc->fnEndList = (gl20PEndList)doGetProcAddress("glEndList");
    glc->fnNormal3b = (gl20PNormal3b)doGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl20PNormal3d)doGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl20PNormal3f)doGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl20PNormal3i)doGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl20PNormal3s)doGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl20PNormal3bv)doGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl20PNormal3dv)doGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl20PNormal3fv)doGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl20PNormal3iv)doGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl20PNormal3sv)doGetProcAddress("glNormal3sv");
    glc->fnNormalPointer = (gl20PNormalPointer)doGetProcAddress("glNormalPointer");
    glc->fnOrtho = (gl20POrtho)doGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl20PPassThrough)doGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl20PPixelMapfv)doGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl20PPixelMapuiv)doGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl20PPixelMapusv)doGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl20PPixelStoref)doGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl20PPixelStorei)doGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl20PPixelTransferf)doGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl20PPixelTransferi)doGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl20PPixelZoom)doGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl20PPointSize)doGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl20PPolygonMode)doGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl20PPolygonStipple)doGetProcAddress("glPolygonStipple");
    glc->fnPrioritizeTextures = (gl20PPrioritizeTextures)doGetProcAddress("glPrioritizeTextures");
    glc->fnPushAttrib = (gl20PPushAttrib)doGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl20PPopAttrib)doGetProcAddress("glPopAttrib");
    glc->fnPushClientAttrib = (gl20PPushClientAttrib)doGetProcAddress("glPushClientAttrib");
    glc->fnPopClientAttrib = (gl20PPopClientAttrib)doGetProcAddress("glPopClientAttrib");
    glc->fnPushMatrix = (gl20PPushMatrix)doGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl20PPopMatrix)doGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl20PPushName)doGetProcAddress("glPushName");
    glc->fnPopName = (gl20PPopName)doGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl20PRasterPos2d)doGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl20PRasterPos2f)doGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl20PRasterPos2i)doGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl20PRasterPos2s)doGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl20PRasterPos3d)doGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl20PRasterPos3f)doGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl20PRasterPos3i)doGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl20PRasterPos3s)doGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl20PRasterPos4d)doGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl20PRasterPos4f)doGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl20PRasterPos4i)doGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl20PRasterPos4s)doGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl20PRasterPos2dv)doGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl20PRasterPos2fv)doGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl20PRasterPos2iv)doGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl20PRasterPos2sv)doGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl20PRasterPos3dv)doGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl20PRasterPos3fv)doGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl20PRasterPos3iv)doGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl20PRasterPos3sv)doGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl20PRasterPos4dv)doGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl20PRasterPos4fv)doGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl20PRasterPos4iv)doGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl20PRasterPos4sv)doGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl20PReadBuffer)doGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl20PReadPixels)doGetProcAddress("glReadPixels");
    glc->fnRectd = (gl20PRectd)doGetProcAddress("glRectd");
    glc->fnRectf = (gl20PRectf)doGetProcAddress("glRectf");
    glc->fnRecti = (gl20PRecti)doGetProcAddress("glRecti");
    glc->fnRects = (gl20PRects)doGetProcAddress("glRects");
    glc->fnRectdv = (gl20PRectdv)doGetProcAddress("glRectdv");
    glc->fnRectfv = (gl20PRectfv)doGetProcAddress("glRectfv");
    glc->fnRectiv = (gl20PRectiv)doGetProcAddress("glRectiv");
    glc->fnRectsv = (gl20PRectsv)doGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl20PRenderMode)doGetProcAddress("glRenderMode");
    glc->fnRotated = (gl20PRotated)doGetProcAddress("glRotated");
    glc->fnRotatef = (gl20PRotatef)doGetProcAddress("glRotatef");
    glc->fnScaled = (gl20PScaled)doGetProcAddress("glScaled");
    glc->fnScalef = (gl20PScalef)doGetProcAddress("glScalef");
    glc->fnScissor = (gl20PScissor)doGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl20PSelectBuffer)doGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl20PShadeModel)doGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl20PStencilFunc)doGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl20PStencilMask)doGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl20PStencilOp)doGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl20PTexCoord1d)doGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl20PTexCoord1f)doGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl20PTexCoord1i)doGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl20PTexCoord1s)doGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl20PTexCoord2d)doGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl20PTexCoord2f)doGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl20PTexCoord2i)doGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl20PTexCoord2s)doGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl20PTexCoord3d)doGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl20PTexCoord3f)doGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl20PTexCoord3i)doGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl20PTexCoord3s)doGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl20PTexCoord4d)doGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl20PTexCoord4f)doGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl20PTexCoord4i)doGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl20PTexCoord4s)doGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl20PTexCoord1dv)doGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl20PTexCoord1fv)doGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl20PTexCoord1iv)doGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl20PTexCoord1sv)doGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl20PTexCoord2dv)doGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl20PTexCoord2fv)doGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl20PTexCoord2iv)doGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl20PTexCoord2sv)doGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl20PTexCoord3dv)doGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl20PTexCoord3fv)doGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl20PTexCoord3iv)doGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl20PTexCoord3sv)doGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl20PTexCoord4dv)doGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl20PTexCoord4fv)doGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl20PTexCoord4iv)doGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl20PTexCoord4sv)doGetProcAddress("glTexCoord4sv");
    glc->fnTexCoordPointer = (gl20PTexCoordPointer)doGetProcAddress("glTexCoordPointer");
    glc->fnTexEnvf = (gl20PTexEnvf)doGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl20PTexEnvi)doGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl20PTexEnvfv)doGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl20PTexEnviv)doGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl20PTexGend)doGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl20PTexGenf)doGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl20PTexGeni)doGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl20PTexGendv)doGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl20PTexGenfv)doGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl20PTexGeniv)doGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl20PTexImage1D)doGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl20PTexImage2D)doGetProcAddress("glTexImage2D");
    glc->fnTexImage3DEXT = (gl20PTexImage3DEXT)wglGetProcAddress("glTexImage3DEXT");
    glc->fnTexParameterf = (gl20PTexParameterf)doGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl20PTexParameteri)doGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl20PTexParameterfv)doGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl20PTexParameteriv)doGetProcAddress("glTexParameteriv");
    glc->fnTexSubImage1D = (gl20PTexSubImage1D)doGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl20PTexSubImage2D)doGetProcAddress("glTexSubImage2D");
    glc->fnTexSubImage3DEXT = (gl20PTexSubImage3DEXT)wglGetProcAddress("glTexSubImage3DEXT");
    glc->fnTranslated = (gl20PTranslated)doGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl20PTranslatef)doGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl20PVertex2s)doGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl20PVertex2i)doGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl20PVertex2f)doGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl20PVertex2d)doGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl20PVertex3s)doGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl20PVertex3i)doGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl20PVertex3f)doGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl20PVertex3d)doGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl20PVertex4s)doGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl20PVertex4i)doGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl20PVertex4f)doGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl20PVertex4d)doGetProcAddress("glVertex4d");
    glc->fnVertexPointer = (gl20PVertexPointer)doGetProcAddress("glVertexPointer");
    glc->fnViewport = (gl20PViewport)doGetProcAddress("glViewport");
    glc->fnAreTexturesResident = (gl20PAreTexturesResident)doGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl20PArrayElement)doGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl20PDrawArrays)doGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl20PDrawElements)doGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl20PGetPointerv)doGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl20PPolygonOffset)doGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl20PCopyTexImage1D)doGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl20PCopyTexImage2D)doGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl20PCopyTexSubImage1D)doGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl20PCopyTexSubImage2D)doGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl20PBindTexture)doGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl20PDeleteTextures)doGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl20PGenTextures)doGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl20PIsTexture)doGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl20PColorPointer)doGetProcAddress("glColorPointer");
    glc->fnBlendColorEXT = (gl20PBlendColorEXT)wglGetProcAddress("glBlendColorEXT");
    glc->fnBlendEquation = (gl20PBlendEquation)wglGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl20PCopyTexSubImage3D)wglGetProcAddress("glCopyTexSubImage3D");
    glc->fnActiveTexture = (gl20PActiveTexture)wglGetProcAddress("glActiveTexture");
    glc->fnClientActiveTexture = (gl20PClientActiveTexture)wglGetProcAddress("glClientActiveTexture");
    glc->fnCompressedTexImage1D = (gl20PCompressedTexImage1D)wglGetProcAddress("glCompressedTexImage1D");
    glc->fnCompressedTexImage2D = (gl20PCompressedTexImage2D)wglGetProcAddress("glCompressedTexImage2D");
    glc->fnCompressedTexImage3D = (gl20PCompressedTexImage3D)wglGetProcAddress("glCompressedTexImage3D");
    glc->fnCompressedTexSubImage1D = (gl20PCompressedTexSubImage1D)wglGetProcAddress("glCompressedTexSubImage1D");
    glc->fnCompressedTexSubImage2D = (gl20PCompressedTexSubImage2D)wglGetProcAddress("glCompressedTexSubImage2D");
    glc->fnCompressedTexSubImage3D = (gl20PCompressedTexSubImage3D)wglGetProcAddress("glCompressedTexSubImage3D");
    glc->fnBlendFuncSeparate = (gl20PBlendFuncSeparate)wglGetProcAddress("glBlendFuncSeparate");
    glc->fnBeginQuery = (gl20PBeginQuery)wglGetProcAddress("glBeginQuery");
    glc->fnBindBuffer = (gl20PBindBuffer)wglGetProcAddress("glBindBuffer");
    glc->fnBufferData = (gl20PBufferData)wglGetProcAddress("glBufferData");
    glc->fnBufferSubData = (gl20PBufferSubData)wglGetProcAddress("glBufferSubData");
    glc->fnAttachShader = (gl20PAttachShader)doGetProcAddress("glAttachShader");
    glc->fnBindAttribLocation = (gl20PBindAttribLocation)doGetProcAddress("glBindAttribLocation");
    glc->fnBlendEquationSeperate = (gl20PBlendEquationSeperate)doGetProcAddress("glBlendEquationSeperate");
    glc->fnCompileShader = (gl20PCompileShader)doGetProcAddress("glCompileShader");
    glc->fnCreateProgram = (gl20PCreateProgram)doGetProcAddress("glCreateProgram");
    glc->fnCreateShader = (gl20PCreateShader)doGetProcAddress("glCreateShader");
    return glc;
}

