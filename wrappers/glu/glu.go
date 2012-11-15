package glu

/*
#cgo LDFLAGS: -lGLU
#include <GL/glu.h>
*/
import "C"

import "fmt"

func init() {
    fmt.Println("WARNING: glu is deprectated and should only be used in test")
    fmt.Println("WARNING: programs. In addition to that, this wrapper only")
    fmt.Println("WARNING: contains a limited amount of bindings.")
}

/*
typedef void (GLAPIENTRYP _GLUfuncptr)(void);

GLAPI void GLAPIENTRY gluBeginCurve (GLUnurbs* nurb);
GLAPI void GLAPIENTRY gluBeginPolygon (GLUtesselator* tess);
GLAPI void GLAPIENTRY gluBeginSurface (GLUnurbs* nurb);
GLAPI void GLAPIENTRY gluBeginTrim (GLUnurbs* nurb);
GLAPI GLint GLAPIENTRY gluBuild1DMipmapLevels (GLenum target, GLint internalFormat, GLsizei width, GLenum format, GLenum type, GLint level, GLint base, GLint max, const void *data);
GLAPI GLint GLAPIENTRY gluBuild1DMipmaps (GLenum target, GLint internalFormat, GLsizei width, GLenum format, GLenum type, const void *data);
GLAPI GLint GLAPIENTRY gluBuild2DMipmapLevels (GLenum target, GLint internalFormat, GLsizei width, GLsizei height, GLenum format, GLenum type, GLint level, GLint base, GLint max, const void *data);
GLAPI GLint GLAPIENTRY gluBuild2DMipmaps (GLenum target, GLint internalFormat, GLsizei width, GLsizei height, GLenum format, GLenum type, const void *data);
GLAPI GLint GLAPIENTRY gluBuild3DMipmapLevels (GLenum target, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLint level, GLint base, GLint max, const void *data);
GLAPI GLint GLAPIENTRY gluBuild3DMipmaps (GLenum target, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const void *data);
GLAPI GLboolean GLAPIENTRY gluCheckExtension (const GLubyte *extName, const GLubyte *extString);
GLAPI void GLAPIENTRY gluCylinder (GLUquadric* quad, GLdouble base, GLdouble top, GLdouble height, GLint slices, GLint stacks);
GLAPI void GLAPIENTRY gluDeleteNurbsRenderer (GLUnurbs* nurb);
GLAPI void GLAPIENTRY gluDeleteQuadric (GLUquadric* quad);
GLAPI void GLAPIENTRY gluDeleteTess (GLUtesselator* tess);
GLAPI void GLAPIENTRY gluDisk (GLUquadric* quad, GLdouble inner, GLdouble outer, GLint slices, GLint loops);
GLAPI void GLAPIENTRY gluEndCurve (GLUnurbs* nurb);
GLAPI void GLAPIENTRY gluEndPolygon (GLUtesselator* tess);
GLAPI void GLAPIENTRY gluEndSurface (GLUnurbs* nurb);
GLAPI void GLAPIENTRY gluEndTrim (GLUnurbs* nurb);
GLAPI const GLubyte * GLAPIENTRY gluErrorString (GLenum error);
GLAPI void GLAPIENTRY gluGetNurbsProperty (GLUnurbs* nurb, GLenum property, GLfloat* data);
GLAPI const GLubyte * GLAPIENTRY gluGetString (GLenum name);
GLAPI void GLAPIENTRY gluGetTessProperty (GLUtesselator* tess, GLenum which, GLdouble* data);
GLAPI void GLAPIENTRY gluLoadSamplingMatrices (GLUnurbs* nurb, const GLfloat *model, const GLfloat *perspective, const GLint *view);
*/

//GLAPI void GLAPIENTRY gluLookAt (GLdouble eyeX, GLdouble eyeY, GLdouble eyeZ, GLdouble centerX, GLdouble centerY, GLdouble centerZ, GLdouble upX, GLdouble upY, GLdouble upZ);
func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) {
    C.gluLookAt(C.GLdouble(eyeX), C.GLdouble(eyeY), C.GLdouble(eyeZ), C.GLdouble(centerX), C.GLdouble(centerY), C.GLdouble(centerZ), C.GLdouble(upX), C.GLdouble(upY), C.GLdouble(upZ))
}

/*
GLAPI GLUnurbs* GLAPIENTRY gluNewNurbsRenderer (void);
GLAPI GLUquadric* GLAPIENTRY gluNewQuadric (void);
GLAPI GLUtesselator* GLAPIENTRY gluNewTess (void);
GLAPI void GLAPIENTRY gluNextContour (GLUtesselator* tess, GLenum type);
GLAPI void GLAPIENTRY gluNurbsCallback (GLUnurbs* nurb, GLenum which, _GLUfuncptr CallBackFunc);
GLAPI void GLAPIENTRY gluNurbsCallbackData (GLUnurbs* nurb, GLvoid* userData);
GLAPI void GLAPIENTRY gluNurbsCallbackDataEXT (GLUnurbs* nurb, GLvoid* userData);
GLAPI void GLAPIENTRY gluNurbsCurve (GLUnurbs* nurb, GLint knotCount, GLfloat *knots, GLint stride, GLfloat *control, GLint order, GLenum type);
GLAPI void GLAPIENTRY gluNurbsProperty (GLUnurbs* nurb, GLenum property, GLfloat value);
GLAPI void GLAPIENTRY gluNurbsSurface (GLUnurbs* nurb, GLint sKnotCount, GLfloat* sKnots, GLint tKnotCount, GLfloat* tKnots, GLint sStride, GLint tStride, GLfloat* control, GLint sOrder, GLint tOrder, GLenum type);
*/

//GLAPI void GLAPIENTRY gluOrtho2D (GLdouble left, GLdouble right, GLdouble bottom, GLdouble top);
func Ortho2D(left, right, bottom, top float32) {
    C.gluOrtho2D(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top))
}

//GLAPI void GLAPIENTRY gluPartialDisk (GLUquadric* quad, GLdouble inner, GLdouble outer, GLint slices, GLint loops, GLdouble start, GLdouble sweep);


//GLAPI void GLAPIENTRY gluPerspective (GLdouble fovy, GLdouble aspect, GLdouble zNear, GLdouble zFar);
func Perspective(fovy, aspect, zNear, zFar float32) {
    C.gluPerspective(C.GLdouble(fovy), C.GLdouble(aspect), C.GLdouble(zNear), C.GLdouble(zFar))
}

/*
GLAPI void GLAPIENTRY gluPickMatrix (GLdouble x, GLdouble y, GLdouble delX, GLdouble delY, GLint *viewport);
GLAPI GLint GLAPIENTRY gluProject (GLdouble objX, GLdouble objY, GLdouble objZ, const GLdouble *model, const GLdouble *proj, const GLint *view, GLdouble* winX, GLdouble* winY, GLdouble* winZ);
GLAPI void GLAPIENTRY gluPwlCurve (GLUnurbs* nurb, GLint count, GLfloat* data, GLint stride, GLenum type);
GLAPI void GLAPIENTRY gluQuadricCallback (GLUquadric* quad, GLenum which, _GLUfuncptr CallBackFunc);
GLAPI void GLAPIENTRY gluQuadricDrawStyle (GLUquadric* quad, GLenum draw);
GLAPI void GLAPIENTRY gluQuadricNormals (GLUquadric* quad, GLenum normal);
GLAPI void GLAPIENTRY gluQuadricOrientation (GLUquadric* quad, GLenum orientation);
GLAPI void GLAPIENTRY gluQuadricTexture (GLUquadric* quad, GLboolean texture);
GLAPI GLint GLAPIENTRY gluScaleImage (GLenum format, GLsizei wIn, GLsizei hIn, GLenum typeIn, const void *dataIn, GLsizei wOut, GLsizei hOut, GLenum typeOut, GLvoid* dataOut);
GLAPI void GLAPIENTRY gluSphere (GLUquadric* quad, GLdouble radius, GLint slices, GLint stacks);
GLAPI void GLAPIENTRY gluTessBeginContour (GLUtesselator* tess);
GLAPI void GLAPIENTRY gluTessBeginPolygon (GLUtesselator* tess, GLvoid* data);
GLAPI void GLAPIENTRY gluTessCallback (GLUtesselator* tess, GLenum which, _GLUfuncptr CallBackFunc);
GLAPI void GLAPIENTRY gluTessEndContour (GLUtesselator* tess);
GLAPI void GLAPIENTRY gluTessEndPolygon (GLUtesselator* tess);
GLAPI void GLAPIENTRY gluTessNormal (GLUtesselator* tess, GLdouble valueX, GLdouble valueY, GLdouble valueZ);
GLAPI void GLAPIENTRY gluTessProperty (GLUtesselator* tess, GLenum which, GLdouble data);
GLAPI void GLAPIENTRY gluTessVertex (GLUtesselator* tess, GLdouble *location, GLvoid* data);
GLAPI GLint GLAPIENTRY gluUnProject (GLdouble winX, GLdouble winY, GLdouble winZ, const GLdouble *model, const GLdouble *proj, const GLint *view, GLdouble* objX, GLdouble* objY, GLdouble* objZ);
GLAPI GLint GLAPIENTRY gluUnProject4 (GLdouble winX, GLdouble winY, GLdouble winZ, GLdouble clipW, const GLdouble *model, const GLdouble *proj, const GLint *view, GLdouble nearVal, GLdouble farVal, GLdouble* objX, GLdouble* objY, GLdouble* objZ, GLdouble* objW);
*/
