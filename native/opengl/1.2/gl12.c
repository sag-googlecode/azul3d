
#include <stdbool.h>
#include <stdlib.h>

#ifdef _WIN32
	#include <windows.h>
#endif

#include "gl12.h"

#ifdef _WIN32
	HMODULE gl12OpenGL32;

	void* gl12LibGetProcAddress(char* name) {
		if(gl12OpenGL32 == NULL) {
			gl12OpenGL32 = LoadLibrary(TEXT("opengl32.dll"));
		}
		return GetProcAddress(gl12OpenGL32, TEXT(name));
	}

	void* gl12GLGetProcAddress(char* name) {
		void* ptr = wglGetProcAddress(name);

		intptr_t iptr = (intptr_t)ptr;

		if(iptr == 0 || iptr == 1 || iptr == 2 || iptr == 3 || iptr == -1) {
			return NULL;
		}
		return ptr;
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

void gl12DepthMask(gl12Context* glc, GLboolean flag) {
    return glc->fnDepthMask(flag);
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

void gl12DrawPixels(gl12Context* glc, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* data) {
    return glc->fnDrawPixels(width, height, format, type, data);
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

void gl12GetPolygonStipple(gl12Context* glc, GLubyte* pattern) {
    return glc->fnGetPolygonStipple(pattern);
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

void gl12IndexMask(gl12Context* glc, GLuint mask) {
    return glc->fnIndexMask(mask);
}

void gl12IndexPointer(gl12Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnIndexPointer(type, stride, pointer);
}

void gl12InitNames(gl12Context* glc) {
    return glc->fnInitNames();
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

void gl12PushAttrib(gl12Context* glc, GLbitfield mask) {
    return glc->fnPushAttrib(mask);
}

void gl12PopAttrib(gl12Context* glc) {
    return glc->fnPopAttrib();
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

void gl12Viewport(gl12Context* glc, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnViewport(x, y, width, height);
}

void gl12GetColorTable(gl12Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* table) {
    return glc->fnGetColorTable(target, format, type, table);
}

void gl12GetColorTableParameterfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetColorTableParameterfv(target, pname, params);
}

void gl12GetColorTableParameteriv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetColorTableParameteriv(target, pname, params);
}

void gl12GetConvolutionFilter(gl12Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* image) {
    return glc->fnGetConvolutionFilter(target, format, type, image);
}

void gl12GetConvolutionParameterfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetConvolutionParameterfv(target, pname, params);
}

void gl12GetConvolutionParameteriv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetConvolutionParameteriv(target, pname, params);
}

void gl12GetHistogram(gl12Context* glc, GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid* values) {
    return glc->fnGetHistogram(target, reset, format, type, values);
}

void gl12GetHistogramParameterfv(gl12Context* glc, GLenum target, GLenum pname, GLfloat* params) {
    return glc->fnGetHistogramParameterfv(target, pname, params);
}

void gl12GetHistogramParameteriv(gl12Context* glc, GLenum target, GLenum pname, GLint* params) {
    return glc->fnGetHistogramParameteriv(target, pname, params);
}

void gl12GetSeparableFilter(gl12Context* glc, GLenum target, GLenum format, GLenum type, GLvoid* row, GLvoid* column, GLvoid* span) {
    return glc->fnGetSeparableFilter(target, format, type, row, column, span);
}

void gl12Histogram(gl12Context* glc, GLenum target, GLsizei width, GLenum internalformat, GLboolean sink) {
    return glc->fnHistogram(target, width, internalformat, sink);
}

void gl12Minmax(gl12Context* glc, GLenum target, GLenum internalformat, GLboolean sink) {
    return glc->fnMinmax(target, internalformat, sink);
}

void gl12MultiTexCoord1s(gl12Context* glc, GLenum target, GLshort s) {
    return glc->fnMultiTexCoord1s(target, s);
}

void gl12MultiTexCoord1i(gl12Context* glc, GLenum target, GLint s) {
    return glc->fnMultiTexCoord1i(target, s);
}

void gl12MultiTexCoord1f(gl12Context* glc, GLenum target, GLfloat s) {
    return glc->fnMultiTexCoord1f(target, s);
}

void gl12MultiTexCoord1d(gl12Context* glc, GLenum target, GLdouble s) {
    return glc->fnMultiTexCoord1d(target, s);
}

void gl12MultiTexCoord2s(gl12Context* glc, GLenum target, GLshort s, GLshort t) {
    return glc->fnMultiTexCoord2s(target, s, t);
}

void gl12MultiTexCoord2i(gl12Context* glc, GLenum target, GLint s, GLint t) {
    return glc->fnMultiTexCoord2i(target, s, t);
}

void gl12MultiTexCoord2f(gl12Context* glc, GLenum target, GLfloat s, GLfloat t) {
    return glc->fnMultiTexCoord2f(target, s, t);
}

void gl12MultiTexCoord2d(gl12Context* glc, GLenum target, GLdouble s, GLdouble t) {
    return glc->fnMultiTexCoord2d(target, s, t);
}

void gl12MultiTexCoord3s(gl12Context* glc, GLenum target, GLshort s, GLshort t, GLshort r) {
    return glc->fnMultiTexCoord3s(target, s, t, r);
}

void gl12MultiTexCoord3i(gl12Context* glc, GLenum target, GLint s, GLint t, GLint r) {
    return glc->fnMultiTexCoord3i(target, s, t, r);
}

void gl12MultiTexCoord3f(gl12Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r) {
    return glc->fnMultiTexCoord3f(target, s, t, r);
}

void gl12MultiTexCoord3d(gl12Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r) {
    return glc->fnMultiTexCoord3d(target, s, t, r);
}

void gl12MultiTexCoord4s(gl12Context* glc, GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
    return glc->fnMultiTexCoord4s(target, s, t, r, q);
}

void gl12MultiTexCoord4i(gl12Context* glc, GLenum target, GLint s, GLint t, GLint r, GLint q) {
    return glc->fnMultiTexCoord4i(target, s, t, r, q);
}

void gl12MultiTexCoord4f(gl12Context* glc, GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
    return glc->fnMultiTexCoord4f(target, s, t, r, q);
}

void gl12MultiTexCoord4d(gl12Context* glc, GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
    return glc->fnMultiTexCoord4d(target, s, t, r, q);
}

void gl12MultiTexCoord1sv(gl12Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord1sv(target, v);
}

void gl12MultiTexCoord1iv(gl12Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord1iv(target, v);
}

void gl12MultiTexCoord1fv(gl12Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord1fv(target, v);
}

void gl12MultiTexCoord1dv(gl12Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord1dv(target, v);
}

void gl12MultiTexCoord2sv(gl12Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord2sv(target, v);
}

void gl12MultiTexCoord2iv(gl12Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord2iv(target, v);
}

void gl12MultiTexCoord2fv(gl12Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord2fv(target, v);
}

void gl12MultiTexCoord2dv(gl12Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord2dv(target, v);
}

void gl12MultiTexCoord3sv(gl12Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord3sv(target, v);
}

void gl12MultiTexCoord3iv(gl12Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord3iv(target, v);
}

void gl12MultiTexCoord3fv(gl12Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord3fv(target, v);
}

void gl12MultiTexCoord3dv(gl12Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord3dv(target, v);
}

void gl12MultiTexCoord4sv(gl12Context* glc, GLenum target, GLshort* v) {
    return glc->fnMultiTexCoord4sv(target, v);
}

void gl12MultiTexCoord4iv(gl12Context* glc, GLenum target, GLint* v) {
    return glc->fnMultiTexCoord4iv(target, v);
}

void gl12MultiTexCoord4fv(gl12Context* glc, GLenum target, GLfloat* v) {
    return glc->fnMultiTexCoord4fv(target, v);
}

void gl12MultiTexCoord4dv(gl12Context* glc, GLenum target, GLdouble* v) {
    return glc->fnMultiTexCoord4dv(target, v);
}

void gl12ResetHistogram(gl12Context* glc, GLenum target) {
    return glc->fnResetHistogram(target);
}

void gl12ResetMinmax(gl12Context* glc, GLenum target) {
    return glc->fnResetMinmax(target);
}

void gl12SeparableFilter2D(gl12Context* glc, GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* row, GLvoid* column) {
    return glc->fnSeparableFilter2D(target, internalformat, width, height, format, type, row, column);
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

void gl12EnableClientState(gl12Context* glc, GLenum cap) {
    return glc->fnEnableClientState(cap);
}

void gl12DisableClientState(gl12Context* glc, GLenum cap) {
    return glc->fnDisableClientState(cap);
}

void gl12Indexub(gl12Context* glc, GLubyte c) {
    return glc->fnIndexub(c);
}

void gl12Indexubv(gl12Context* glc, GLubyte* c) {
    return glc->fnIndexubv(c);
}

void gl12InterleavedArrays(gl12Context* glc, GLenum format, GLsizei stride, GLvoid* pointer) {
    return glc->fnInterleavedArrays(format, stride, pointer);
}

void gl12NormalPointer(gl12Context* glc, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnNormalPointer(type, stride, pointer);
}

void gl12PushClientAttrib(gl12Context* glc, GLbitfield mask) {
    return glc->fnPushClientAttrib(mask);
}

void gl12PrioritizeTextures(gl12Context* glc, GLsizei n, GLuint* textures, GLclampf* priorities) {
    return glc->fnPrioritizeTextures(n, textures, priorities);
}

void gl12PopClientAttrib(gl12Context* glc) {
    return glc->fnPopClientAttrib();
}

void gl12TexCoordPointer(gl12Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnTexCoordPointer(size, type, stride, pointer);
}

void gl12TexSubImage1D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage1D(target, level, xoffset, width, format, type, pixels);
}

void gl12TexSubImage2D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage2D(target, level, xoffset, yoffset, width, height, format, type, pixels);
}

void gl12VertexPointer(gl12Context* glc, GLint size, GLenum type, GLsizei stride, GLvoid* pointer) {
    return glc->fnVertexPointer(size, type, stride, pointer);
}

void gl12BlendColor(gl12Context* glc, GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
    return glc->fnBlendColor(red, green, blue, alpha);
}

void gl12BlendEquation(gl12Context* glc, GLenum mode) {
    return glc->fnBlendEquation(mode);
}

void gl12CopyTexSubImage3D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
    return glc->fnCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}

void gl12DrawRangeElements(gl12Context* glc, GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices) {
    return glc->fnDrawRangeElements(mode, start, end, count, type, indices);
}

void gl12TexImage3D(gl12Context* glc, GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexImage3D(target, level, internalformat, width, height, depth, border, format, type, pixels);
}

void gl12TexSubImage3D(gl12Context* glc, GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels) {
    return glc->fnTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

gl12Context* gl12NewContext() {
    gl12Context* glc = calloc(1, sizeof(gl12Context));

    // Preload all procedures
    glc->fnAccum = (gl12PAccum)gl12LibGetProcAddress("glAccum");
    glc->fnAlphaFunc = (gl12PAlphaFunc)gl12LibGetProcAddress("glAlphaFunc");
    glc->fnBegin = (gl12PBegin)gl12LibGetProcAddress("glBegin");
    glc->fnEnd = (gl12PEnd)gl12LibGetProcAddress("glEnd");
    glc->fnBitmap = (gl12PBitmap)gl12LibGetProcAddress("glBitmap");
    glc->fnBlendFunc = (gl12PBlendFunc)gl12LibGetProcAddress("glBlendFunc");
    glc->fnCallList = (gl12PCallList)gl12LibGetProcAddress("glCallList");
    glc->fnCallLists = (gl12PCallLists)gl12LibGetProcAddress("glCallLists");
    glc->fnClear = (gl12PClear)gl12LibGetProcAddress("glClear");
    glc->fnClearAccum = (gl12PClearAccum)gl12LibGetProcAddress("glClearAccum");
    glc->fnClearColor = (gl12PClearColor)gl12LibGetProcAddress("glClearColor");
    glc->fnClearDepth = (gl12PClearDepth)gl12LibGetProcAddress("glClearDepth");
    glc->fnClearIndex = (gl12PClearIndex)gl12LibGetProcAddress("glClearIndex");
    glc->fnClearStencil = (gl12PClearStencil)gl12LibGetProcAddress("glClearStencil");
    glc->fnClipPlane = (gl12PClipPlane)gl12LibGetProcAddress("glClipPlane");
    glc->fnColor3b = (gl12PColor3b)gl12LibGetProcAddress("glColor3b");
    glc->fnColor3d = (gl12PColor3d)gl12LibGetProcAddress("glColor3d");
    glc->fnColor3f = (gl12PColor3f)gl12LibGetProcAddress("glColor3f");
    glc->fnColor3i = (gl12PColor3i)gl12LibGetProcAddress("glColor3i");
    glc->fnColor3s = (gl12PColor3s)gl12LibGetProcAddress("glColor3s");
    glc->fnColor3ub = (gl12PColor3ub)gl12LibGetProcAddress("glColor3ub");
    glc->fnColor3ui = (gl12PColor3ui)gl12LibGetProcAddress("glColor3ui");
    glc->fnColor3us = (gl12PColor3us)gl12LibGetProcAddress("glColor3us");
    glc->fnColor4b = (gl12PColor4b)gl12LibGetProcAddress("glColor4b");
    glc->fnColor4d = (gl12PColor4d)gl12LibGetProcAddress("glColor4d");
    glc->fnColor4f = (gl12PColor4f)gl12LibGetProcAddress("glColor4f");
    glc->fnColor4i = (gl12PColor4i)gl12LibGetProcAddress("glColor4i");
    glc->fnColor4s = (gl12PColor4s)gl12LibGetProcAddress("glColor4s");
    glc->fnColor4ub = (gl12PColor4ub)gl12LibGetProcAddress("glColor4ub");
    glc->fnColor4ui = (gl12PColor4ui)gl12LibGetProcAddress("glColor4ui");
    glc->fnColor4us = (gl12PColor4us)gl12LibGetProcAddress("glColor4us");
    glc->fnColor3bv = (gl12PColor3bv)gl12LibGetProcAddress("glColor3bv");
    glc->fnColor3dv = (gl12PColor3dv)gl12LibGetProcAddress("glColor3dv");
    glc->fnColor3fv = (gl12PColor3fv)gl12LibGetProcAddress("glColor3fv");
    glc->fnColor3iv = (gl12PColor3iv)gl12LibGetProcAddress("glColor3iv");
    glc->fnColor3sv = (gl12PColor3sv)gl12LibGetProcAddress("glColor3sv");
    glc->fnColor3ubv = (gl12PColor3ubv)gl12LibGetProcAddress("glColor3ubv");
    glc->fnColor3uiv = (gl12PColor3uiv)gl12LibGetProcAddress("glColor3uiv");
    glc->fnColor3usv = (gl12PColor3usv)gl12LibGetProcAddress("glColor3usv");
    glc->fnColor4bv = (gl12PColor4bv)gl12LibGetProcAddress("glColor4bv");
    glc->fnColor4dv = (gl12PColor4dv)gl12LibGetProcAddress("glColor4dv");
    glc->fnColor4fv = (gl12PColor4fv)gl12LibGetProcAddress("glColor4fv");
    glc->fnColor4iv = (gl12PColor4iv)gl12LibGetProcAddress("glColor4iv");
    glc->fnColor4sv = (gl12PColor4sv)gl12LibGetProcAddress("glColor4sv");
    glc->fnColor4ubv = (gl12PColor4ubv)gl12LibGetProcAddress("glColor4ubv");
    glc->fnColor4uiv = (gl12PColor4uiv)gl12LibGetProcAddress("glColor4uiv");
    glc->fnColor4usv = (gl12PColor4usv)gl12LibGetProcAddress("glColor4usv");
    glc->fnColorMask = (gl12PColorMask)gl12LibGetProcAddress("glColorMask");
    glc->fnColorMaterial = (gl12PColorMaterial)gl12LibGetProcAddress("glColorMaterial");
    glc->fnColorTable = (gl12PColorTable)gl12GLGetProcAddress("glColorTable");
    glc->fnColorTableParameterfv = (gl12PColorTableParameterfv)gl12GLGetProcAddress("glColorTableParameterfv");
    glc->fnColorTableParameteriv = (gl12PColorTableParameteriv)gl12GLGetProcAddress("glColorTableParameteriv");
    glc->fnColorSubTable = (gl12PColorSubTable)gl12GLGetProcAddress("glColorSubTable");
    glc->fnCopyPixels = (gl12PCopyPixels)gl12LibGetProcAddress("glCopyPixels");
    glc->fnCullFace = (gl12PCullFace)gl12LibGetProcAddress("glCullFace");
    glc->fnConvolutionFilter1D = (gl12PConvolutionFilter1D)gl12GLGetProcAddress("glConvolutionFilter1D");
    glc->fnConvolutionFilter2D = (gl12PConvolutionFilter2D)gl12GLGetProcAddress("glConvolutionFilter2D");
    glc->fnConvolutionParameterf = (gl12PConvolutionParameterf)gl12GLGetProcAddress("glConvolutionParameterf");
    glc->fnConvolutionParameteri = (gl12PConvolutionParameteri)gl12GLGetProcAddress("glConvolutionParameteri");
    glc->fnCopyColorTable = (gl12PCopyColorTable)gl12GLGetProcAddress("glCopyColorTable");
    glc->fnCopyColorSubTable = (gl12PCopyColorSubTable)gl12GLGetProcAddress("glCopyColorSubTable");
    glc->fnCopyConvolutionFilter1D = (gl12PCopyConvolutionFilter1D)gl12GLGetProcAddress("glCopyConvolutionFilter1D");
    glc->fnCopyConvolutionFilter2D = (gl12PCopyConvolutionFilter2D)gl12GLGetProcAddress("glCopyConvolutionFilter2D");
    glc->fnDeleteLists = (gl12PDeleteLists)gl12LibGetProcAddress("glDeleteLists");
    glc->fnDepthFunc = (gl12PDepthFunc)gl12LibGetProcAddress("glDepthFunc");
    glc->fnDepthMask = (gl12PDepthMask)gl12LibGetProcAddress("glDepthMask");
    glc->fnDepthRange = (gl12PDepthRange)gl12LibGetProcAddress("glDepthRange");
    glc->fnEnable = (gl12PEnable)gl12LibGetProcAddress("glEnable");
    glc->fnDisable = (gl12PDisable)gl12LibGetProcAddress("glDisable");
    glc->fnDrawBuffer = (gl12PDrawBuffer)gl12LibGetProcAddress("glDrawBuffer");
    glc->fnDrawPixels = (gl12PDrawPixels)gl12LibGetProcAddress("glDrawPixels");
    glc->fnEdgeFlag = (gl12PEdgeFlag)gl12LibGetProcAddress("glEdgeFlag");
    glc->fnEdgeFlagv = (gl12PEdgeFlagv)gl12LibGetProcAddress("glEdgeFlagv");
    glc->fnEdgeFlagPointer = (gl12PEdgeFlagPointer)gl12LibGetProcAddress("glEdgeFlagPointer");
    glc->fnEvalCoord1d = (gl12PEvalCoord1d)gl12LibGetProcAddress("glEvalCoord1d");
    glc->fnEvalCoord1f = (gl12PEvalCoord1f)gl12LibGetProcAddress("glEvalCoord1f");
    glc->fnEvalCoord2d = (gl12PEvalCoord2d)gl12LibGetProcAddress("glEvalCoord2d");
    glc->fnEvalCoord2f = (gl12PEvalCoord2f)gl12LibGetProcAddress("glEvalCoord2f");
    glc->fnEvalCoord1dv = (gl12PEvalCoord1dv)gl12LibGetProcAddress("glEvalCoord1dv");
    glc->fnEvalCoord1fv = (gl12PEvalCoord1fv)gl12LibGetProcAddress("glEvalCoord1fv");
    glc->fnEvalCoord2dv = (gl12PEvalCoord2dv)gl12LibGetProcAddress("glEvalCoord2dv");
    glc->fnEvalCoord2fv = (gl12PEvalCoord2fv)gl12LibGetProcAddress("glEvalCoord2fv");
    glc->fnEvalMesh1 = (gl12PEvalMesh1)gl12LibGetProcAddress("glEvalMesh1");
    glc->fnEvalMesh2 = (gl12PEvalMesh2)gl12LibGetProcAddress("glEvalMesh2");
    glc->fnEvalPoint1 = (gl12PEvalPoint1)gl12LibGetProcAddress("glEvalPoint1");
    glc->fnEvalPoint2 = (gl12PEvalPoint2)gl12LibGetProcAddress("glEvalPoint2");
    glc->fnFeedbackBuffer = (gl12PFeedbackBuffer)gl12LibGetProcAddress("glFeedbackBuffer");
    glc->fnFinish = (gl12PFinish)gl12LibGetProcAddress("glFinish");
    glc->fnFlush = (gl12PFlush)gl12LibGetProcAddress("glFlush");
    glc->fnFogf = (gl12PFogf)gl12LibGetProcAddress("glFogf");
    glc->fnFogi = (gl12PFogi)gl12LibGetProcAddress("glFogi");
    glc->fnFogfv = (gl12PFogfv)gl12LibGetProcAddress("glFogfv");
    glc->fnFogiv = (gl12PFogiv)gl12LibGetProcAddress("glFogiv");
    glc->fnFrontFace = (gl12PFrontFace)gl12LibGetProcAddress("glFrontFace");
    glc->fnFrustum = (gl12PFrustum)gl12LibGetProcAddress("glFrustum");
    glc->fnGenLists = (gl12PGenLists)gl12LibGetProcAddress("glGenLists");
    glc->fnGetBooleanv = (gl12PGetBooleanv)gl12LibGetProcAddress("glGetBooleanv");
    glc->fnGetDoublev = (gl12PGetDoublev)gl12LibGetProcAddress("glGetDoublev");
    glc->fnGetFloatv = (gl12PGetFloatv)gl12LibGetProcAddress("glGetFloatv");
    glc->fnGetIntegerv = (gl12PGetIntegerv)gl12LibGetProcAddress("glGetIntegerv");
    glc->fnGetClipPlane = (gl12PGetClipPlane)gl12LibGetProcAddress("glGetClipPlane");
    glc->fnGetError = (gl12PGetError)gl12LibGetProcAddress("glGetError");
    glc->fnGetLightfv = (gl12PGetLightfv)gl12LibGetProcAddress("glGetLightfv");
    glc->fnGetLightiv = (gl12PGetLightiv)gl12LibGetProcAddress("glGetLightiv");
    glc->fnGetMapdv = (gl12PGetMapdv)gl12LibGetProcAddress("glGetMapdv");
    glc->fnGetMapfv = (gl12PGetMapfv)gl12LibGetProcAddress("glGetMapfv");
    glc->fnGetMapiv = (gl12PGetMapiv)gl12LibGetProcAddress("glGetMapiv");
    glc->fnGetMaterialfv = (gl12PGetMaterialfv)gl12LibGetProcAddress("glGetMaterialfv");
    glc->fnGetMaterialiv = (gl12PGetMaterialiv)gl12LibGetProcAddress("glGetMaterialiv");
    glc->fnGetPixelMapfv = (gl12PGetPixelMapfv)gl12LibGetProcAddress("glGetPixelMapfv");
    glc->fnGetPixelMapuiv = (gl12PGetPixelMapuiv)gl12LibGetProcAddress("glGetPixelMapuiv");
    glc->fnGetPixelMapusv = (gl12PGetPixelMapusv)gl12LibGetProcAddress("glGetPixelMapusv");
    glc->fnGetPolygonStipple = (gl12PGetPolygonStipple)gl12LibGetProcAddress("glGetPolygonStipple");
    glc->fnGetString = (gl12PGetString)gl12LibGetProcAddress("glGetString");
    glc->fnGetTexEnvfv = (gl12PGetTexEnvfv)gl12LibGetProcAddress("glGetTexEnvfv");
    glc->fnGetTexEnviv = (gl12PGetTexEnviv)gl12LibGetProcAddress("glGetTexEnviv");
    glc->fnGetTexGendv = (gl12PGetTexGendv)gl12LibGetProcAddress("glGetTexGendv");
    glc->fnGetTexGenfv = (gl12PGetTexGenfv)gl12LibGetProcAddress("glGetTexGenfv");
    glc->fnGetTexGeniv = (gl12PGetTexGeniv)gl12LibGetProcAddress("glGetTexGeniv");
    glc->fnGetTexImage = (gl12PGetTexImage)gl12LibGetProcAddress("glGetTexImage");
    glc->fnGetTexLevelParameterfv = (gl12PGetTexLevelParameterfv)gl12LibGetProcAddress("glGetTexLevelParameterfv");
    glc->fnGetTexLevelParameteriv = (gl12PGetTexLevelParameteriv)gl12LibGetProcAddress("glGetTexLevelParameteriv");
    glc->fnGetTexParameterfv = (gl12PGetTexParameterfv)gl12LibGetProcAddress("glGetTexParameterfv");
    glc->fnGetTexParameteriv = (gl12PGetTexParameteriv)gl12LibGetProcAddress("glGetTexParameteriv");
    glc->fnHint = (gl12PHint)gl12LibGetProcAddress("glHint");
    glc->fnIndexd = (gl12PIndexd)gl12LibGetProcAddress("glIndexd");
    glc->fnIndexf = (gl12PIndexf)gl12LibGetProcAddress("glIndexf");
    glc->fnIndexi = (gl12PIndexi)gl12LibGetProcAddress("glIndexi");
    glc->fnIndexs = (gl12PIndexs)gl12LibGetProcAddress("glIndexs");
    glc->fnIndexdv = (gl12PIndexdv)gl12LibGetProcAddress("glIndexdv");
    glc->fnIndexfv = (gl12PIndexfv)gl12LibGetProcAddress("glIndexfv");
    glc->fnIndexiv = (gl12PIndexiv)gl12LibGetProcAddress("glIndexiv");
    glc->fnIndexsv = (gl12PIndexsv)gl12LibGetProcAddress("glIndexsv");
    glc->fnIndexMask = (gl12PIndexMask)gl12LibGetProcAddress("glIndexMask");
    glc->fnIndexPointer = (gl12PIndexPointer)gl12LibGetProcAddress("glIndexPointer");
    glc->fnInitNames = (gl12PInitNames)gl12LibGetProcAddress("glInitNames");
    glc->fnIsEnabled = (gl12PIsEnabled)gl12LibGetProcAddress("glIsEnabled");
    glc->fnIsList = (gl12PIsList)gl12LibGetProcAddress("glIsList");
    glc->fnLightf = (gl12PLightf)gl12LibGetProcAddress("glLightf");
    glc->fnLighti = (gl12PLighti)gl12LibGetProcAddress("glLighti");
    glc->fnLightfv = (gl12PLightfv)gl12LibGetProcAddress("glLightfv");
    glc->fnLightiv = (gl12PLightiv)gl12LibGetProcAddress("glLightiv");
    glc->fnLightModelf = (gl12PLightModelf)gl12LibGetProcAddress("glLightModelf");
    glc->fnLightModeli = (gl12PLightModeli)gl12LibGetProcAddress("glLightModeli");
    glc->fnLightModelfv = (gl12PLightModelfv)gl12LibGetProcAddress("glLightModelfv");
    glc->fnLightModeliv = (gl12PLightModeliv)gl12LibGetProcAddress("glLightModeliv");
    glc->fnLineStipple = (gl12PLineStipple)gl12LibGetProcAddress("glLineStipple");
    glc->fnLineWidth = (gl12PLineWidth)gl12LibGetProcAddress("glLineWidth");
    glc->fnListBase = (gl12PListBase)gl12LibGetProcAddress("glListBase");
    glc->fnLoadIdentity = (gl12PLoadIdentity)gl12LibGetProcAddress("glLoadIdentity");
    glc->fnLoadMatrixd = (gl12PLoadMatrixd)gl12LibGetProcAddress("glLoadMatrixd");
    glc->fnLoadMatrixf = (gl12PLoadMatrixf)gl12LibGetProcAddress("glLoadMatrixf");
    glc->fnLoadName = (gl12PLoadName)gl12LibGetProcAddress("glLoadName");
    glc->fnLogicOp = (gl12PLogicOp)gl12LibGetProcAddress("glLogicOp");
    glc->fnMap1d = (gl12PMap1d)gl12LibGetProcAddress("glMap1d");
    glc->fnMap1f = (gl12PMap1f)gl12LibGetProcAddress("glMap1f");
    glc->fnMap2d = (gl12PMap2d)gl12LibGetProcAddress("glMap2d");
    glc->fnMap2f = (gl12PMap2f)gl12LibGetProcAddress("glMap2f");
    glc->fnMapGrid1d = (gl12PMapGrid1d)gl12LibGetProcAddress("glMapGrid1d");
    glc->fnMapGrid1f = (gl12PMapGrid1f)gl12LibGetProcAddress("glMapGrid1f");
    glc->fnMapGrid2d = (gl12PMapGrid2d)gl12LibGetProcAddress("glMapGrid2d");
    glc->fnMapGrid2f = (gl12PMapGrid2f)gl12LibGetProcAddress("glMapGrid2f");
    glc->fnMaterialf = (gl12PMaterialf)gl12LibGetProcAddress("glMaterialf");
    glc->fnMateriali = (gl12PMateriali)gl12LibGetProcAddress("glMateriali");
    glc->fnMaterialfv = (gl12PMaterialfv)gl12LibGetProcAddress("glMaterialfv");
    glc->fnMaterialiv = (gl12PMaterialiv)gl12LibGetProcAddress("glMaterialiv");
    glc->fnMatrixMode = (gl12PMatrixMode)gl12LibGetProcAddress("glMatrixMode");
    glc->fnMultMatrixd = (gl12PMultMatrixd)gl12LibGetProcAddress("glMultMatrixd");
    glc->fnMultMatrixf = (gl12PMultMatrixf)gl12LibGetProcAddress("glMultMatrixf");
    glc->fnNewList = (gl12PNewList)gl12LibGetProcAddress("glNewList");
    glc->fnEndList = (gl12PEndList)gl12LibGetProcAddress("glEndList");
    glc->fnNormal3b = (gl12PNormal3b)gl12LibGetProcAddress("glNormal3b");
    glc->fnNormal3d = (gl12PNormal3d)gl12LibGetProcAddress("glNormal3d");
    glc->fnNormal3f = (gl12PNormal3f)gl12LibGetProcAddress("glNormal3f");
    glc->fnNormal3i = (gl12PNormal3i)gl12LibGetProcAddress("glNormal3i");
    glc->fnNormal3s = (gl12PNormal3s)gl12LibGetProcAddress("glNormal3s");
    glc->fnNormal3bv = (gl12PNormal3bv)gl12LibGetProcAddress("glNormal3bv");
    glc->fnNormal3dv = (gl12PNormal3dv)gl12LibGetProcAddress("glNormal3dv");
    glc->fnNormal3fv = (gl12PNormal3fv)gl12LibGetProcAddress("glNormal3fv");
    glc->fnNormal3iv = (gl12PNormal3iv)gl12LibGetProcAddress("glNormal3iv");
    glc->fnNormal3sv = (gl12PNormal3sv)gl12LibGetProcAddress("glNormal3sv");
    glc->fnOrtho = (gl12POrtho)gl12LibGetProcAddress("glOrtho");
    glc->fnPassThrough = (gl12PPassThrough)gl12LibGetProcAddress("glPassThrough");
    glc->fnPixelMapfv = (gl12PPixelMapfv)gl12LibGetProcAddress("glPixelMapfv");
    glc->fnPixelMapuiv = (gl12PPixelMapuiv)gl12LibGetProcAddress("glPixelMapuiv");
    glc->fnPixelMapusv = (gl12PPixelMapusv)gl12LibGetProcAddress("glPixelMapusv");
    glc->fnPixelStoref = (gl12PPixelStoref)gl12LibGetProcAddress("glPixelStoref");
    glc->fnPixelStorei = (gl12PPixelStorei)gl12LibGetProcAddress("glPixelStorei");
    glc->fnPixelTransferf = (gl12PPixelTransferf)gl12LibGetProcAddress("glPixelTransferf");
    glc->fnPixelTransferi = (gl12PPixelTransferi)gl12LibGetProcAddress("glPixelTransferi");
    glc->fnPixelZoom = (gl12PPixelZoom)gl12LibGetProcAddress("glPixelZoom");
    glc->fnPointSize = (gl12PPointSize)gl12LibGetProcAddress("glPointSize");
    glc->fnPolygonMode = (gl12PPolygonMode)gl12LibGetProcAddress("glPolygonMode");
    glc->fnPolygonStipple = (gl12PPolygonStipple)gl12LibGetProcAddress("glPolygonStipple");
    glc->fnPushAttrib = (gl12PPushAttrib)gl12LibGetProcAddress("glPushAttrib");
    glc->fnPopAttrib = (gl12PPopAttrib)gl12LibGetProcAddress("glPopAttrib");
    glc->fnPushMatrix = (gl12PPushMatrix)gl12LibGetProcAddress("glPushMatrix");
    glc->fnPopMatrix = (gl12PPopMatrix)gl12LibGetProcAddress("glPopMatrix");
    glc->fnPushName = (gl12PPushName)gl12LibGetProcAddress("glPushName");
    glc->fnPopName = (gl12PPopName)gl12LibGetProcAddress("glPopName");
    glc->fnRasterPos2d = (gl12PRasterPos2d)gl12LibGetProcAddress("glRasterPos2d");
    glc->fnRasterPos2f = (gl12PRasterPos2f)gl12LibGetProcAddress("glRasterPos2f");
    glc->fnRasterPos2i = (gl12PRasterPos2i)gl12LibGetProcAddress("glRasterPos2i");
    glc->fnRasterPos2s = (gl12PRasterPos2s)gl12LibGetProcAddress("glRasterPos2s");
    glc->fnRasterPos3d = (gl12PRasterPos3d)gl12LibGetProcAddress("glRasterPos3d");
    glc->fnRasterPos3f = (gl12PRasterPos3f)gl12LibGetProcAddress("glRasterPos3f");
    glc->fnRasterPos3i = (gl12PRasterPos3i)gl12LibGetProcAddress("glRasterPos3i");
    glc->fnRasterPos3s = (gl12PRasterPos3s)gl12LibGetProcAddress("glRasterPos3s");
    glc->fnRasterPos4d = (gl12PRasterPos4d)gl12LibGetProcAddress("glRasterPos4d");
    glc->fnRasterPos4f = (gl12PRasterPos4f)gl12LibGetProcAddress("glRasterPos4f");
    glc->fnRasterPos4i = (gl12PRasterPos4i)gl12LibGetProcAddress("glRasterPos4i");
    glc->fnRasterPos4s = (gl12PRasterPos4s)gl12LibGetProcAddress("glRasterPos4s");
    glc->fnRasterPos2dv = (gl12PRasterPos2dv)gl12LibGetProcAddress("glRasterPos2dv");
    glc->fnRasterPos2fv = (gl12PRasterPos2fv)gl12LibGetProcAddress("glRasterPos2fv");
    glc->fnRasterPos2iv = (gl12PRasterPos2iv)gl12LibGetProcAddress("glRasterPos2iv");
    glc->fnRasterPos2sv = (gl12PRasterPos2sv)gl12LibGetProcAddress("glRasterPos2sv");
    glc->fnRasterPos3dv = (gl12PRasterPos3dv)gl12LibGetProcAddress("glRasterPos3dv");
    glc->fnRasterPos3fv = (gl12PRasterPos3fv)gl12LibGetProcAddress("glRasterPos3fv");
    glc->fnRasterPos3iv = (gl12PRasterPos3iv)gl12LibGetProcAddress("glRasterPos3iv");
    glc->fnRasterPos3sv = (gl12PRasterPos3sv)gl12LibGetProcAddress("glRasterPos3sv");
    glc->fnRasterPos4dv = (gl12PRasterPos4dv)gl12LibGetProcAddress("glRasterPos4dv");
    glc->fnRasterPos4fv = (gl12PRasterPos4fv)gl12LibGetProcAddress("glRasterPos4fv");
    glc->fnRasterPos4iv = (gl12PRasterPos4iv)gl12LibGetProcAddress("glRasterPos4iv");
    glc->fnRasterPos4sv = (gl12PRasterPos4sv)gl12LibGetProcAddress("glRasterPos4sv");
    glc->fnReadBuffer = (gl12PReadBuffer)gl12LibGetProcAddress("glReadBuffer");
    glc->fnReadPixels = (gl12PReadPixels)gl12LibGetProcAddress("glReadPixels");
    glc->fnRectd = (gl12PRectd)gl12LibGetProcAddress("glRectd");
    glc->fnRectf = (gl12PRectf)gl12LibGetProcAddress("glRectf");
    glc->fnRecti = (gl12PRecti)gl12LibGetProcAddress("glRecti");
    glc->fnRects = (gl12PRects)gl12LibGetProcAddress("glRects");
    glc->fnRectdv = (gl12PRectdv)gl12LibGetProcAddress("glRectdv");
    glc->fnRectfv = (gl12PRectfv)gl12LibGetProcAddress("glRectfv");
    glc->fnRectiv = (gl12PRectiv)gl12LibGetProcAddress("glRectiv");
    glc->fnRectsv = (gl12PRectsv)gl12LibGetProcAddress("glRectsv");
    glc->fnRenderMode = (gl12PRenderMode)gl12LibGetProcAddress("glRenderMode");
    glc->fnRotated = (gl12PRotated)gl12LibGetProcAddress("glRotated");
    glc->fnRotatef = (gl12PRotatef)gl12LibGetProcAddress("glRotatef");
    glc->fnScaled = (gl12PScaled)gl12LibGetProcAddress("glScaled");
    glc->fnScalef = (gl12PScalef)gl12LibGetProcAddress("glScalef");
    glc->fnScissor = (gl12PScissor)gl12LibGetProcAddress("glScissor");
    glc->fnSelectBuffer = (gl12PSelectBuffer)gl12LibGetProcAddress("glSelectBuffer");
    glc->fnShadeModel = (gl12PShadeModel)gl12LibGetProcAddress("glShadeModel");
    glc->fnStencilFunc = (gl12PStencilFunc)gl12LibGetProcAddress("glStencilFunc");
    glc->fnStencilMask = (gl12PStencilMask)gl12LibGetProcAddress("glStencilMask");
    glc->fnStencilOp = (gl12PStencilOp)gl12LibGetProcAddress("glStencilOp");
    glc->fnTexCoord1d = (gl12PTexCoord1d)gl12LibGetProcAddress("glTexCoord1d");
    glc->fnTexCoord1f = (gl12PTexCoord1f)gl12LibGetProcAddress("glTexCoord1f");
    glc->fnTexCoord1i = (gl12PTexCoord1i)gl12LibGetProcAddress("glTexCoord1i");
    glc->fnTexCoord1s = (gl12PTexCoord1s)gl12LibGetProcAddress("glTexCoord1s");
    glc->fnTexCoord2d = (gl12PTexCoord2d)gl12LibGetProcAddress("glTexCoord2d");
    glc->fnTexCoord2f = (gl12PTexCoord2f)gl12LibGetProcAddress("glTexCoord2f");
    glc->fnTexCoord2i = (gl12PTexCoord2i)gl12LibGetProcAddress("glTexCoord2i");
    glc->fnTexCoord2s = (gl12PTexCoord2s)gl12LibGetProcAddress("glTexCoord2s");
    glc->fnTexCoord3d = (gl12PTexCoord3d)gl12LibGetProcAddress("glTexCoord3d");
    glc->fnTexCoord3f = (gl12PTexCoord3f)gl12LibGetProcAddress("glTexCoord3f");
    glc->fnTexCoord3i = (gl12PTexCoord3i)gl12LibGetProcAddress("glTexCoord3i");
    glc->fnTexCoord3s = (gl12PTexCoord3s)gl12LibGetProcAddress("glTexCoord3s");
    glc->fnTexCoord4d = (gl12PTexCoord4d)gl12LibGetProcAddress("glTexCoord4d");
    glc->fnTexCoord4f = (gl12PTexCoord4f)gl12LibGetProcAddress("glTexCoord4f");
    glc->fnTexCoord4i = (gl12PTexCoord4i)gl12LibGetProcAddress("glTexCoord4i");
    glc->fnTexCoord4s = (gl12PTexCoord4s)gl12LibGetProcAddress("glTexCoord4s");
    glc->fnTexCoord1dv = (gl12PTexCoord1dv)gl12LibGetProcAddress("glTexCoord1dv");
    glc->fnTexCoord1fv = (gl12PTexCoord1fv)gl12LibGetProcAddress("glTexCoord1fv");
    glc->fnTexCoord1iv = (gl12PTexCoord1iv)gl12LibGetProcAddress("glTexCoord1iv");
    glc->fnTexCoord1sv = (gl12PTexCoord1sv)gl12LibGetProcAddress("glTexCoord1sv");
    glc->fnTexCoord2dv = (gl12PTexCoord2dv)gl12LibGetProcAddress("glTexCoord2dv");
    glc->fnTexCoord2fv = (gl12PTexCoord2fv)gl12LibGetProcAddress("glTexCoord2fv");
    glc->fnTexCoord2iv = (gl12PTexCoord2iv)gl12LibGetProcAddress("glTexCoord2iv");
    glc->fnTexCoord2sv = (gl12PTexCoord2sv)gl12LibGetProcAddress("glTexCoord2sv");
    glc->fnTexCoord3dv = (gl12PTexCoord3dv)gl12LibGetProcAddress("glTexCoord3dv");
    glc->fnTexCoord3fv = (gl12PTexCoord3fv)gl12LibGetProcAddress("glTexCoord3fv");
    glc->fnTexCoord3iv = (gl12PTexCoord3iv)gl12LibGetProcAddress("glTexCoord3iv");
    glc->fnTexCoord3sv = (gl12PTexCoord3sv)gl12LibGetProcAddress("glTexCoord3sv");
    glc->fnTexCoord4dv = (gl12PTexCoord4dv)gl12LibGetProcAddress("glTexCoord4dv");
    glc->fnTexCoord4fv = (gl12PTexCoord4fv)gl12LibGetProcAddress("glTexCoord4fv");
    glc->fnTexCoord4iv = (gl12PTexCoord4iv)gl12LibGetProcAddress("glTexCoord4iv");
    glc->fnTexCoord4sv = (gl12PTexCoord4sv)gl12LibGetProcAddress("glTexCoord4sv");
    glc->fnTexEnvf = (gl12PTexEnvf)gl12LibGetProcAddress("glTexEnvf");
    glc->fnTexEnvi = (gl12PTexEnvi)gl12LibGetProcAddress("glTexEnvi");
    glc->fnTexEnvfv = (gl12PTexEnvfv)gl12LibGetProcAddress("glTexEnvfv");
    glc->fnTexEnviv = (gl12PTexEnviv)gl12LibGetProcAddress("glTexEnviv");
    glc->fnTexGend = (gl12PTexGend)gl12LibGetProcAddress("glTexGend");
    glc->fnTexGenf = (gl12PTexGenf)gl12LibGetProcAddress("glTexGenf");
    glc->fnTexGeni = (gl12PTexGeni)gl12LibGetProcAddress("glTexGeni");
    glc->fnTexGendv = (gl12PTexGendv)gl12LibGetProcAddress("glTexGendv");
    glc->fnTexGenfv = (gl12PTexGenfv)gl12LibGetProcAddress("glTexGenfv");
    glc->fnTexGeniv = (gl12PTexGeniv)gl12LibGetProcAddress("glTexGeniv");
    glc->fnTexImage1D = (gl12PTexImage1D)gl12LibGetProcAddress("glTexImage1D");
    glc->fnTexImage2D = (gl12PTexImage2D)gl12LibGetProcAddress("glTexImage2D");
    glc->fnTexParameterf = (gl12PTexParameterf)gl12LibGetProcAddress("glTexParameterf");
    glc->fnTexParameteri = (gl12PTexParameteri)gl12LibGetProcAddress("glTexParameteri");
    glc->fnTexParameterfv = (gl12PTexParameterfv)gl12LibGetProcAddress("glTexParameterfv");
    glc->fnTexParameteriv = (gl12PTexParameteriv)gl12LibGetProcAddress("glTexParameteriv");
    glc->fnTranslated = (gl12PTranslated)gl12LibGetProcAddress("glTranslated");
    glc->fnTranslatef = (gl12PTranslatef)gl12LibGetProcAddress("glTranslatef");
    glc->fnVertex2s = (gl12PVertex2s)gl12LibGetProcAddress("glVertex2s");
    glc->fnVertex2i = (gl12PVertex2i)gl12LibGetProcAddress("glVertex2i");
    glc->fnVertex2f = (gl12PVertex2f)gl12LibGetProcAddress("glVertex2f");
    glc->fnVertex2d = (gl12PVertex2d)gl12LibGetProcAddress("glVertex2d");
    glc->fnVertex3s = (gl12PVertex3s)gl12LibGetProcAddress("glVertex3s");
    glc->fnVertex3i = (gl12PVertex3i)gl12LibGetProcAddress("glVertex3i");
    glc->fnVertex3f = (gl12PVertex3f)gl12LibGetProcAddress("glVertex3f");
    glc->fnVertex3d = (gl12PVertex3d)gl12LibGetProcAddress("glVertex3d");
    glc->fnVertex4s = (gl12PVertex4s)gl12LibGetProcAddress("glVertex4s");
    glc->fnVertex4i = (gl12PVertex4i)gl12LibGetProcAddress("glVertex4i");
    glc->fnVertex4f = (gl12PVertex4f)gl12LibGetProcAddress("glVertex4f");
    glc->fnVertex4d = (gl12PVertex4d)gl12LibGetProcAddress("glVertex4d");
    glc->fnViewport = (gl12PViewport)gl12LibGetProcAddress("glViewport");
    glc->fnGetColorTable = (gl12PGetColorTable)gl12GLGetProcAddress("glGetColorTable");
    glc->fnGetColorTableParameterfv = (gl12PGetColorTableParameterfv)gl12GLGetProcAddress("glGetColorTableParameterfv");
    glc->fnGetColorTableParameteriv = (gl12PGetColorTableParameteriv)gl12GLGetProcAddress("glGetColorTableParameteriv");
    glc->fnGetConvolutionFilter = (gl12PGetConvolutionFilter)gl12GLGetProcAddress("glGetConvolutionFilter");
    glc->fnGetConvolutionParameterfv = (gl12PGetConvolutionParameterfv)gl12LibGetProcAddress("glGetConvolutionParameterfv");
    glc->fnGetConvolutionParameteriv = (gl12PGetConvolutionParameteriv)gl12LibGetProcAddress("glGetConvolutionParameteriv");
    glc->fnGetHistogram = (gl12PGetHistogram)gl12GLGetProcAddress("glGetHistogram");
    glc->fnGetHistogramParameterfv = (gl12PGetHistogramParameterfv)gl12GLGetProcAddress("glGetHistogramParameterfv");
    glc->fnGetHistogramParameteriv = (gl12PGetHistogramParameteriv)gl12GLGetProcAddress("glGetHistogramParameteriv");
    glc->fnGetSeparableFilter = (gl12PGetSeparableFilter)gl12GLGetProcAddress("glGetSeparableFilter");
    glc->fnHistogram = (gl12PHistogram)gl12GLGetProcAddress("glHistogram");
    glc->fnMinmax = (gl12PMinmax)gl12GLGetProcAddress("glMinmax");
    glc->fnMultiTexCoord1s = (gl12PMultiTexCoord1s)gl12GLGetProcAddress("glMultiTexCoord1s");
    glc->fnMultiTexCoord1i = (gl12PMultiTexCoord1i)gl12GLGetProcAddress("glMultiTexCoord1i");
    glc->fnMultiTexCoord1f = (gl12PMultiTexCoord1f)gl12GLGetProcAddress("glMultiTexCoord1f");
    glc->fnMultiTexCoord1d = (gl12PMultiTexCoord1d)gl12GLGetProcAddress("glMultiTexCoord1d");
    glc->fnMultiTexCoord2s = (gl12PMultiTexCoord2s)gl12GLGetProcAddress("glMultiTexCoord2s");
    glc->fnMultiTexCoord2i = (gl12PMultiTexCoord2i)gl12GLGetProcAddress("glMultiTexCoord2i");
    glc->fnMultiTexCoord2f = (gl12PMultiTexCoord2f)gl12GLGetProcAddress("glMultiTexCoord2f");
    glc->fnMultiTexCoord2d = (gl12PMultiTexCoord2d)gl12GLGetProcAddress("glMultiTexCoord2d");
    glc->fnMultiTexCoord3s = (gl12PMultiTexCoord3s)gl12GLGetProcAddress("glMultiTexCoord3s");
    glc->fnMultiTexCoord3i = (gl12PMultiTexCoord3i)gl12GLGetProcAddress("glMultiTexCoord3i");
    glc->fnMultiTexCoord3f = (gl12PMultiTexCoord3f)gl12GLGetProcAddress("glMultiTexCoord3f");
    glc->fnMultiTexCoord3d = (gl12PMultiTexCoord3d)gl12GLGetProcAddress("glMultiTexCoord3d");
    glc->fnMultiTexCoord4s = (gl12PMultiTexCoord4s)gl12GLGetProcAddress("glMultiTexCoord4s");
    glc->fnMultiTexCoord4i = (gl12PMultiTexCoord4i)gl12GLGetProcAddress("glMultiTexCoord4i");
    glc->fnMultiTexCoord4f = (gl12PMultiTexCoord4f)gl12GLGetProcAddress("glMultiTexCoord4f");
    glc->fnMultiTexCoord4d = (gl12PMultiTexCoord4d)gl12GLGetProcAddress("glMultiTexCoord4d");
    glc->fnMultiTexCoord1sv = (gl12PMultiTexCoord1sv)gl12GLGetProcAddress("glMultiTexCoord1sv");
    glc->fnMultiTexCoord1iv = (gl12PMultiTexCoord1iv)gl12GLGetProcAddress("glMultiTexCoord1iv");
    glc->fnMultiTexCoord1fv = (gl12PMultiTexCoord1fv)gl12GLGetProcAddress("glMultiTexCoord1fv");
    glc->fnMultiTexCoord1dv = (gl12PMultiTexCoord1dv)gl12GLGetProcAddress("glMultiTexCoord1dv");
    glc->fnMultiTexCoord2sv = (gl12PMultiTexCoord2sv)gl12GLGetProcAddress("glMultiTexCoord2sv");
    glc->fnMultiTexCoord2iv = (gl12PMultiTexCoord2iv)gl12GLGetProcAddress("glMultiTexCoord2iv");
    glc->fnMultiTexCoord2fv = (gl12PMultiTexCoord2fv)gl12GLGetProcAddress("glMultiTexCoord2fv");
    glc->fnMultiTexCoord2dv = (gl12PMultiTexCoord2dv)gl12GLGetProcAddress("glMultiTexCoord2dv");
    glc->fnMultiTexCoord3sv = (gl12PMultiTexCoord3sv)gl12GLGetProcAddress("glMultiTexCoord3sv");
    glc->fnMultiTexCoord3iv = (gl12PMultiTexCoord3iv)gl12GLGetProcAddress("glMultiTexCoord3iv");
    glc->fnMultiTexCoord3fv = (gl12PMultiTexCoord3fv)gl12GLGetProcAddress("glMultiTexCoord3fv");
    glc->fnMultiTexCoord3dv = (gl12PMultiTexCoord3dv)gl12GLGetProcAddress("glMultiTexCoord3dv");
    glc->fnMultiTexCoord4sv = (gl12PMultiTexCoord4sv)gl12GLGetProcAddress("glMultiTexCoord4sv");
    glc->fnMultiTexCoord4iv = (gl12PMultiTexCoord4iv)gl12GLGetProcAddress("glMultiTexCoord4iv");
    glc->fnMultiTexCoord4fv = (gl12PMultiTexCoord4fv)gl12GLGetProcAddress("glMultiTexCoord4fv");
    glc->fnMultiTexCoord4dv = (gl12PMultiTexCoord4dv)gl12GLGetProcAddress("glMultiTexCoord4dv");
    glc->fnResetHistogram = (gl12PResetHistogram)gl12GLGetProcAddress("glResetHistogram");
    glc->fnResetMinmax = (gl12PResetMinmax)gl12GLGetProcAddress("glResetMinmax");
    glc->fnSeparableFilter2D = (gl12PSeparableFilter2D)gl12GLGetProcAddress("glSeparableFilter2D");
    glc->fnAreTexturesResident = (gl12PAreTexturesResident)gl12LibGetProcAddress("glAreTexturesResident");
    glc->fnArrayElement = (gl12PArrayElement)gl12LibGetProcAddress("glArrayElement");
    glc->fnDrawArrays = (gl12PDrawArrays)gl12LibGetProcAddress("glDrawArrays");
    glc->fnDrawElements = (gl12PDrawElements)gl12LibGetProcAddress("glDrawElements");
    glc->fnGetPointerv = (gl12PGetPointerv)gl12LibGetProcAddress("glGetPointerv");
    glc->fnPolygonOffset = (gl12PPolygonOffset)gl12LibGetProcAddress("glPolygonOffset");
    glc->fnCopyTexImage1D = (gl12PCopyTexImage1D)gl12LibGetProcAddress("glCopyTexImage1D");
    glc->fnCopyTexImage2D = (gl12PCopyTexImage2D)gl12LibGetProcAddress("glCopyTexImage2D");
    glc->fnCopyTexSubImage1D = (gl12PCopyTexSubImage1D)gl12LibGetProcAddress("glCopyTexSubImage1D");
    glc->fnCopyTexSubImage2D = (gl12PCopyTexSubImage2D)gl12LibGetProcAddress("glCopyTexSubImage2D");
    glc->fnBindTexture = (gl12PBindTexture)gl12LibGetProcAddress("glBindTexture");
    glc->fnDeleteTextures = (gl12PDeleteTextures)gl12LibGetProcAddress("glDeleteTextures");
    glc->fnGenTextures = (gl12PGenTextures)gl12LibGetProcAddress("glGenTextures");
    glc->fnIsTexture = (gl12PIsTexture)gl12LibGetProcAddress("glIsTexture");
    glc->fnColorPointer = (gl12PColorPointer)gl12LibGetProcAddress("glColorPointer");
    glc->fnEnableClientState = (gl12PEnableClientState)gl12LibGetProcAddress("glEnableClientState");
    glc->fnDisableClientState = (gl12PDisableClientState)gl12LibGetProcAddress("glDisableClientState");
    glc->fnIndexub = (gl12PIndexub)gl12LibGetProcAddress("glIndexub");
    glc->fnIndexubv = (gl12PIndexubv)gl12LibGetProcAddress("glIndexubv");
    glc->fnInterleavedArrays = (gl12PInterleavedArrays)gl12LibGetProcAddress("glInterleavedArrays");
    glc->fnNormalPointer = (gl12PNormalPointer)gl12LibGetProcAddress("glNormalPointer");
    glc->fnPushClientAttrib = (gl12PPushClientAttrib)gl12LibGetProcAddress("glPushClientAttrib");
    glc->fnPrioritizeTextures = (gl12PPrioritizeTextures)gl12LibGetProcAddress("glPrioritizeTextures");
    glc->fnPopClientAttrib = (gl12PPopClientAttrib)gl12LibGetProcAddress("glPopClientAttrib");
    glc->fnTexCoordPointer = (gl12PTexCoordPointer)gl12LibGetProcAddress("glTexCoordPointer");
    glc->fnTexSubImage1D = (gl12PTexSubImage1D)gl12LibGetProcAddress("glTexSubImage1D");
    glc->fnTexSubImage2D = (gl12PTexSubImage2D)gl12LibGetProcAddress("glTexSubImage2D");
    glc->fnVertexPointer = (gl12PVertexPointer)gl12LibGetProcAddress("glVertexPointer");
    glc->fnBlendColor = (gl12PBlendColor)gl12GLGetProcAddress("glBlendColor");
    glc->fnBlendEquation = (gl12PBlendEquation)gl12GLGetProcAddress("glBlendEquation");
    glc->fnCopyTexSubImage3D = (gl12PCopyTexSubImage3D)gl12GLGetProcAddress("glCopyTexSubImage3D");
    glc->fnDrawRangeElements = (gl12PDrawRangeElements)gl12GLGetProcAddress("glDrawRangeElements");
    glc->fnTexImage3D = (gl12PTexImage3D)gl12GLGetProcAddress("glTexImage3D");
    glc->fnTexSubImage3D = (gl12PTexSubImage3D)gl12GLGetProcAddress("glTexSubImage3D");
    return glc;
}

