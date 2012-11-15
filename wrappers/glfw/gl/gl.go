package gl

/*
#cgo LDFLAGS: -lGL
#include <GL/gl.h>
*/
import "C"

// TYPE MAPPING
// GLenum     -> uint
// GLboolean  -> byte
// GLbitfield -> uint
// GLvoid     -> void
// GLbyte     -> int8
// GLshort    -> int16
// GLint      -> int
// GLubyte    -> byte
// GLushort   -> uint16
// GLuint     -> uint
// GLsizei    -> int
// GLfloat    -> float32
// GLclampf   -> float32
// GLdouble   -> float32
// GLclampd   -> float32




/*
//
 * Datatypes
 //
typedef unsigned int	GLenum;
typedef unsigned char	GLboolean;
typedef unsigned int	GLbitfield;
typedef void		GLvoid;
typedef signed char	GLbyte;		// 1-byte signed //
typedef short		GLshort;	// 2-byte signed //
typedef int		GLint;		// 4-byte signed //
typedef unsigned char	GLubyte;	// 1-byte unsigned //
typedef unsigned short	GLushort;	// 2-byte unsigned //
typedef unsigned int	GLuint;		// 4-byte unsigned //
typedef int		GLsizei;	// 4-byte signed //
typedef float		GLfloat;	// single precision float //
typedef float		GLclampf;	// single precision float in [0,1] //
typedef double		GLdouble;	// double precision float //
typedef double		GLclampd;	// double precision float in [0,1] //



//
 * Constants
 //

// Boolean values //
#define GL_FALSE				0x0
#define GL_TRUE					0x1

// Data types //
#define GL_BYTE					0x1400
#define GL_UNSIGNED_BYTE			0x1401
#define GL_SHORT				0x1402
#define GL_UNSIGNED_SHORT			0x1403
#define GL_INT					0x1404
#define GL_UNSIGNED_INT				0x1405
#define GL_FLOAT				0x1406
#define GL_2_BYTES				0x1407
#define GL_3_BYTES				0x1408
#define GL_4_BYTES				0x1409
#define GL_DOUBLE				0x140A
*/

// ------------------------------------------------------------------------- //
// Miscellaneous
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glClearIndex( GLfloat c );
func ClearIndex(c float32) {
    C.glClearIndex(C.GLfloat(c))
}

//GLAPI void GLAPIENTRY glClearColor( GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha );
func ClearColor(red, green, blue, alpha float32) {
    C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

//GLAPI void GLAPIENTRY glClear( GLbitfield mask );
func Clear(mask uint) {
    C.glClear(C.GLbitfield(mask))
}

//GLAPI void GLAPIENTRY glIndexMask( GLuint mask );
func IndexMask(mask uint) {
    C.glIndexMask(C.GLuint(mask))
}

//GLAPI void GLAPIENTRY glColorMask( GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha );
func ColorMask(red, green, blue, alpha uint8) {
    C.glColorMask(C.GLboolean(red), C.GLboolean(green), C.GLboolean(blue), C.GLboolean(alpha))
}

//GLAPI void GLAPIENTRY glAlphaFunc( GLenum func, GLclampf ref );
func AlphaFunc(fun uint, ref float32) {
    C.glAlphaFunc(C.GLenum(fun), C.GLclampf(ref))
}

//GLAPI void GLAPIENTRY glBlendFunc( GLenum sfactor, GLenum dfactor );
func BlendFunc(sfactor uint, dfactor uint) {
    C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

//GLAPI void GLAPIENTRY glLogicOp( GLenum opcode );
func LogicOp(opcode uint) {
    C.glLogicOp(C.GLenum(opcode))
}

//GLAPI void GLAPIENTRY glCullFace( GLenum mode );
func CullFace(mode uint) {
    C.glCullFace(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glFrontFace( GLenum mode );
func FrontFace(mode uint) {
    C.glFrontFace(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glPointSize( GLfloat size );
func PointSize(size float32) {
    C.glPointSize(C.GLfloat(size))
}

//GLAPI void GLAPIENTRY glLineWidth( GLfloat width );
func LineWidth(width float32) {
    C.glLineWidth(C.GLfloat(width))
}

//GLAPI void GLAPIENTRY glLineStipple( GLint factor, GLushort pattern );
func LineStipple(factor int, pattern uint16) {
    C.glLineStipple(C.GLint(factor), C.GLushort(pattern))
}

// FIXME:
//GLAPI void GLAPIENTRY glPolygonMode( GLenum face, GLenum mode );
func PolygonMode(face, mode uint) {
    C.glPolygonMode(C.GLenum(face), C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glPolygonOffset( GLfloat factor, GLfloat units );
// the above two arguments are actually Glenum... whatever
func PolygonOffset(factor, units uint) {
    C.glPolygonMode(C.GLenum(factor), C.GLenum(units))
}

// FIXME
//GLAPI void GLAPIENTRY glPolygonStipple( const GLubyte *mask );
//GLAPI void GLAPIENTRY glGetPolygonStipple( GLubyte *mask );

//GLAPI void GLAPIENTRY glEdgeFlag( GLboolean flag );
func EdgeFlag(flag uint8) {
    C.glEdgeFlag(C.GLboolean(flag))
}

//GLAPI void GLAPIENTRY glEdgeFlagv( const GLboolean *flag );
func EdgeFlagv() (uint8) {
    var flag C.GLboolean
    C.glEdgeFlagv(&flag)
    return uint8(flag)
}

//GLAPI void GLAPIENTRY glScissor( GLint x, GLint y, GLsizei width, GLsizei height);
func Scissor(x, y, width, height int) {
    C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// FIXME
//GLAPI void GLAPIENTRY glClipPlane( GLenum plane, const GLdouble *equation );
//GLAPI void GLAPIENTRY glGetClipPlane( GLenum plane, GLdouble *equation );

//GLAPI void GLAPIENTRY glDrawBuffer( GLenum mode );
func DrawBuffer(mode uint) {
    C.glDrawBuffer(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glReadBuffer( GLenum mode );
func ReadBuffer(mode uint) {
    C.glReadBuffer(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glEnable( GLenum cap );
func Enable(cap uint) {
    C.glEnable(C.GLenum(cap))
}

//GLAPI void GLAPIENTRY glDisable( GLenum cap );
func Disable(cap uint) {
    C.glDisable(C.GLenum(cap))
}

//GLAPI GLboolean GLAPIENTRY glIsEnabled( GLenum cap );
func IsEnabled(cap uint) (bool) {
    return C.glIsEnabled(C.GLenum(cap)) != 0
}

//GLAPI void GLAPIENTRY glEnableClientState( GLenum cap );  // 1.1 //
func EnableClientState(cap uint) {
    C.glEnableClientState(C.GLenum(cap))
}

//GLAPI void GLAPIENTRY glDisableClientState( GLenum cap );  // 1.1 //
func DisableClientState(cap uint) {
    C.glDisableClientState(C.GLenum(cap))
}

// FIXME
//GLAPI void GLAPIENTRY glGetBooleanv( GLenum pname, GLboolean *params );
//GLAPI void GLAPIENTRY glGetDoublev( GLenum pname, GLdouble *params );
//GLAPI void GLAPIENTRY glGetFloatv( GLenum pname, GLfloat *params );
//GLAPI void GLAPIENTRY glGetIntegerv( GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glPushAttrib( GLbitfield mask );
func PushAttrib(mask uint) {
    C.glPushAttrib(C.GLbitfield(mask))
}

//GLAPI void GLAPIENTRY glPopAttrib( void );
func PopAttrib() {
    C.glPopAttrib()
}

//GLAPI void GLAPIENTRY glPushClientAttrib( GLbitfield mask );  // 1.1 //
func PushClientAttrib(mask uint) {
    C.glPushClientAttrib(C.GLbitfield(mask))
}

//GLAPI void GLAPIENTRY glPopClientAttrib( void );  // 1.1 //
func PopClientAttrib() {
    C.glPopClientAttrib()
}

//GLAPI GLint GLAPIENTRY glRenderMode( GLenum mode );
func RenderMode(mode uint) {
    C.glRenderMode(C.GLenum(mode))
}

//GLAPI GLenum GLAPIENTRY glGetError( void );
func GetError() (uint) {
    return uint(C.glGetError())
}

//GLAPI const GLubyte * GLAPIENTRY glGetString( GLenum name );

//GLAPI void GLAPIENTRY glFinish( void );
func Finish() {
    C.glFinish()
}

//GLAP void GLAPIENTRY glFlush( void );
func Flush() {
    C.glFlush()
}

//GLAP void GLAPIENTRY glHint( GLenum target, GLenum mode );
func Hint(target, mode uint) {
    C.glHint(C.GLenum(target), C.GLenum(mode))
}

// ------------------------------------------------------------------------- //
// Depth Buffer
// ------------------------------------------------------------------------- //
//GLAP void GLAPIENTRY glClearDepth( GLclampd depth );
func ClearDepth(depth float32) {
    C.glClearDepth(C.GLclampd(depth))
}

//GLAP void GLAPIENTRY glDepthFunc( GLenum func );
func DepthFunc(fun uint) {
    C.glDepthFunc(C.GLenum(fun))
}

//GLAP void GLAPIENTRY glDepthMask( GLboolean flag );
func DepthMask(flag byte) {
    C.glDepthMask(C.GLboolean(flag))
}

//GLAP void GLAPIENTRY glDepthRange( GLclampd near_val, GLclampd far_val );
func DepthRange(near_val, far_val float32) {
    C.glDepthRange(C.GLclampd(near_val), C.GLclampd(far_val))
}

// ------------------------------------------------------------------------- //
// Accumulation Buffer
// ------------------------------------------------------------------------- //
//GLAP void GLAPIENTRY glClearAccum( GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha );
func ClearAccum(red, green, blue, alpha float32) {
    C.glClearAccum(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

//GLAP void GLAPIENTRY glAccum( GLenum op, GLfloat value );
func Accum(op uint, value float32) {
    C.glAccum(C.GLenum(op), C.GLfloat(value))
}

// ------------------------------------------------------------------------- //
// Transformation
// ------------------------------------------------------------------------- //
//GLAP void GLAPIENTRY glMatrixMode( GLenum mode );
func MatrixMode(mode uint) {
    C.glMatrixMode(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glOrtho( GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble near_val, GLdouble far_val );
func Ortho(left, right, bottom, top, near_val, far_val float32) {
    C.glOrtho(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(near_val), C.GLdouble(far_val))
}

//GLAPI void GLAPIENTRY glFrustum( GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble near_val, GLdouble far_val );


//GLAPI void GLAPIENTRY glViewport( GLint x, GLint y, GLsizei width, GLsizei height );
func Viewport(x, y, width, height int) {
    C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

//GLAPI void GLAPIENTRY glPushMatrix( void );
func PushMatrix() {
    C.glPushMatrix()
}

//GLAPI void GLAPIENTRY glPopMatrix( void );
func PopMatrix() {
    C.glPopMatrix()
}

//GLAPI void GLAPIENTRY glLoadIdentity( void );
func LoadIdentity() {
    C.glLoadIdentity()
}

// FIXME
//GLAPI void GLAPIENTRY glLoadMatrixd( const GLdouble *m );
//GLAPI void GLAPIENTRY glLoadMatrixf( const GLfloat *m );
//GLAPI void GLAPIENTRY glMultMatrixd( const GLdouble *m );
//GLAPI void GLAPIENTRY glMultMatrixf( const GLfloat *m );

//GLAPI void GLAPIENTRY glRotated( GLdouble angle, GLdouble x, GLdouble y, GLdouble z );
func Rotated(angle, x, y, z float32) {
    C.glRotated(C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//GLAPI void GLAPIENTRY glRotatef( GLfloat angle, GLfloat x, GLfloat y, GLfloat z );
func Rotatef(angle, x, y, z float32) {
    C.glRotatef(C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//GLAPI void GLAPIENTRY glScaled( GLdouble x, GLdouble y, GLdouble z );
func Scaled(x, y, z float32) {
    C.glScaled(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//GLAPI void GLAPIENTRY glScalef( GLfloat x, GLfloat y, GLfloat z );
func Scalef(x, y, z float32) {
    C.glScaled(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//GLAPI void GLAPIENTRY glTranslated( GLdouble x, GLdouble y, GLdouble z );
func Translated(x, y, z float32) {
    C.glTranslated(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//GLAPI void GLAPIENTRY glTranslatef( GLfloat x, GLfloat y, GLfloat z );
func Translatef(x, y, z float32) {
    C.glTranslatef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}


// ------------------------------------------------------------------------- //
// Display Lists
// ------------------------------------------------------------------------- //
//GLAPI GLboolean GLAPIENTRY glIsList( GLuint list );
func IsList(list uint) (bool) {
    return C.glIsList(C.GLuint(list)) != 0
}

//GLAPI void GLAPIENTRY glDeleteLists( GLuint list, GLsizei range );
func DeleteLists(list uint, r int) {
    C.glDeleteLists(C.GLuint(list), C.GLsizei(r))
}

//GLAPI GLuint GLAPIENTRY glGenLists( GLsizei range );
func GenLists(r int) {
    C.glGenLists(C.GLsizei(r))
}

//GLAPI void GLAPIENTRY glNewList( GLuint list, GLenum mode );
func NewList(list, mode uint) {
    C.glNewList(C.GLuint(list), C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glEndList( void );
func EndList() {
    C.glEndList()
}

//GLAPI void GLAPIENTRY glCallList( GLuint list );
func CallList(list uint) {
    C.glCallList(C.GLuint(list))
}

// FIXME
//GLAPI void GLAPIENTRY glCallLists( GLsizei n, GLenum type, const GLvoid *lists );

//GLAPI void GLAPIENTRY glListBase( GLuint base );
func ListBase(base uint) {
    C.glListBase(C.GLuint(base))
}

// ------------------------------------------------------------------------- //
// Drawing Functions
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glBegin( GLenum mode );
func Begin(mode uint) {
    C.glBegin(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glEnd( void );
func End() {
    C.glEnd()
}

//GLAPI void GLAPIENTRY glVertex2d( GLdouble x, GLdouble y );
func Vertex2d(x, y float32) {
    C.glVertex2d(C.GLdouble(x), C.GLdouble(y))
}

//GLAPI void GLAPIENTRY glVertex2f( GLfloat x, GLfloat y );
func Vertex2f(x, y float32) {
    C.glVertex2f(C.GLfloat(x), C.GLfloat(y))
}

//GLAPI void GLAPIENTRY glVertex2i( GLint x, GLint y );
func Vertex2i(x, y int) {
    C.glVertex2i(C.GLint(x), C.GLint(y))
}

//GLAPI void GLAPIENTRY glVertex2s( GLshort x, GLshort y );
func Vertex2s(x, y int16) {
    C.glVertex2s(C.GLshort(x), C.GLshort(y))
}

//GLAPI void GLAPIENTRY glVertex3d( GLdouble x, GLdouble y, GLdouble z );
func Vertex3d(x, y, z float32) {
    C.glVertex3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//GLAPI void GLAPIENTRY glVertex3f( GLfloat x, GLfloat y, GLfloat z );
func Vertex3f(x, y, z float32) {
    C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//GLAPI void GLAPIENTRY glVertex3i( GLint x, GLint y, GLint z );
func Vertex3i(x, y, z int) {
    C.glVertex3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

//GLAPI void GLAPIENTRY glVertex3s( GLshort x, GLshort y, GLshort z );
func Vertex3s(x, y, z int16) {
    C.glVertex3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

//GLAPI void GLAPIENTRY glVertex4d( GLdouble x, GLdouble y, GLdouble z, GLdouble w );
func Vertex4d(x, y, z, w float32) {
    C.glVertex4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

//GLAPI void GLAPIENTRY glVertex4f( GLfloat x, GLfloat y, GLfloat z, GLfloat w );
func Vertex4f(x, y, z, w float32) {
    C.glVertex4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

//GLAPI void GLAPIENTRY glVertex4i( GLint x, GLint y, GLint z, GLint w );
func Vertex4i(x, y, z, w int) {
    C.glVertex4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

//GLAPI void GLAPIENTRY glVertex4s( GLshort x, GLshort y, GLshort z, GLshort w );
func Vertex4s(x, y, z, w int16) {
    C.glVertex4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

// FIXME
//GLAPI void GLAPIENTRY glVertex2dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glVertex2fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glVertex2iv( const GLint *v );
//GLAPI void GLAPIENTRY glVertex2sv( const GLshort *v );

//GLAPI void GLAPIENTRY glVertex3dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glVertex3fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glVertex3iv( const GLint *v );
//GLAPI void GLAPIENTRY glVertex3sv( const GLshort *v );

//GLAPI void GLAPIENTRY glVertex4dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glVertex4fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glVertex4iv( const GLint *v );
//GLAPI void GLAPIENTRY glVertex4sv( const GLshort *v );


//GLAPI void GLAPIENTRY glNormal3b( GLbyte nx, GLbyte ny, GLbyte nz );
func Normal3b(nx, ny, nz int8) {
    C.glNormal3b(C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
}

//GLAPI void GLAPIENTRY glNormal3d( GLdouble nx, GLdouble ny, GLdouble nz );
func Normal3d(nx, ny, nz float32) {
    C.glNormal3d(C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
}

//GLAPI void GLAPIENTRY glNormal3f( GLfloat nx, GLfloat ny, GLfloat nz );
func Normal3f(nx, ny, nz float32) {
    C.glNormal3f(C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
}

//GLAPI void GLAPIENTRY glNormal3i( GLint nx, GLint ny, GLint nz );
func Normal3i(nx, ny, nz int) {
    C.glNormal3i(C.GLint(nx), C.GLint(ny), C.GLint(nz))
}

//GLAPI void GLAPIENTRY glNormal3s( GLshort nx, GLshort ny, GLshort nz );
func Normal3s(nx, ny, nz int16) {
    C.glNormal3s(C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
}

// FIXME
//GLAPI void GLAPIENTRY glNormal3bv( const GLbyte *v );
//GLAPI void GLAPIENTRY glNormal3dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glNormal3fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glNormal3iv( const GLint *v );
//GLAPI void GLAPIENTRY glNormal3sv( const GLshort *v );


//GLAPI void GLAPIENTRY glIndexd( GLdouble c );
func Indexd(c float32) {
    C.glIndexd(C.GLdouble(c))
}

//GLAPI void GLAPIENTRY glIndexf( GLfloat c );
func Indexf(c float32) {
    C.glIndexf(C.GLfloat(c))
}

//GLAPI void GLAPIENTRY glIndexi( GLint c );
func Indexi(c int) {
    C.glIndexi(C.GLint(c))
}

//GLAPI void GLAPIENTRY glIndexs( GLshort c );
func Indexs(c int16) {
    C.glIndexs(C.GLshort(c))
}

//GLAPI void GLAPIENTRY glIndexub( GLubyte c );  // 1.1 //
func Indexub(c uint8) {
    C.glIndexub(C.GLubyte(c))
}

// FIXME
//GLAPI void GLAPIENTRY glIndexdv( const GLdouble *c );
//GLAPI void GLAPIENTRY glIndexfv( const GLfloat *c );
//GLAPI void GLAPIENTRY glIndexiv( const GLint *c );
//GLAPI void GLAPIENTRY glIndexsv( const GLshort *c );
//GLAPI void GLAPIENTRY glIndexubv( const GLubyte *c );  // 1.1 //

//GLAPI void GLAPIENTRY glColor3b( GLbyte red, GLbyte green, GLbyte blue );
func Color3b(red, green, blue int8) {
    C.glColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
}

//GLAPI void GLAPIENTRY glColor3d( GLdouble red, GLdouble green, GLdouble blue );
func Color3d(red, green, blue float32) {
    C.glColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
}

//GLAPI void GLAPIENTRY glColor3f( GLfloat red, GLfloat green, GLfloat blue );
func Color3f(red, green, blue float32) {
    C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

//GLAPI void GLAPIENTRY glColor3i( GLint red, GLint green, GLint blue );
func Color3i(red, green, blue int) {
    C.glColor3i(C.GLint(red), C.GLint(green), C.GLint(blue))
}

//GLAPI void GLAPIENTRY glColor3s( GLshort red, GLshort green, GLshort blue );
func Color3s(red, green, blue int16) {
    C.glColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue))
}

//GLAPI void GLAPIENTRY glColor3ub( GLubyte red, GLubyte green, GLubyte blue );
func Color3ub(red, green, blue uint8) {
    C.glColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
}

//GLAPI void GLAPIENTRY glColor3ui( GLuint red, GLuint green, GLuint blue );
func Color3ui(red, green, blue uint) {
    C.glColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue))
}

//GLAPI void GLAPIENTRY glColor3us( GLushort red, GLushort green, GLushort blue );
func Color3us(red, green, blue uint16) {
    C.glColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue))
}


//GLAPI void GLAPIENTRY glColor4b( GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha );
func Color4b(red, green, blue, alpha int8) {
    C.glColor4b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
}

//GLAPI void GLAPIENTRY glColor4d( GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha );
func Color4d(red, green, blue, alpha float32) {
    C.glColor4d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
}

//GLAPI void GLAPIENTRY glColor4f( GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha );
func Color4f(red, green, blue, alpha float32) {
    C.glColor4f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

//GLAPI void GLAPIENTRY glColor4i( GLint red, GLint green, GLint blue, GLint alpha );
func Color4i(red, green, blue, alpha int) {
    C.glColor4i(C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
}

//GLAPI void GLAPIENTRY glColor4s( GLshort red, GLshort green, GLshort blue, GLshort alpha );
func Color4s(red, green, blue, alpha int16) {
    C.glColor4s(C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
}

//GLAPI void GLAPIENTRY glColor4ub( GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha );
func Color4ub(red, green, blue, alpha uint8) {
    C.glColor4ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
}

//GLAPI void GLAPIENTRY glColor4ui( GLuint red, GLuint green, GLuint blue, GLuint alpha );
func Color4ui(red, green, blue, alpha uint) {
    C.glColor4ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
}

//GLAPI void GLAPIENTRY glColor4us( GLushort red, GLushort green, GLushort blue, GLushort alpha );
func Color4us(red, green, blue, alpha uint16) {
    C.glColor4us(C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
}

// FIXME
//GLAPI void GLAPIENTRY glColor3bv( const GLbyte *v );
//GLAPI void GLAPIENTRY glColor3dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glColor3fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glColor3iv( const GLint *v );
//GLAPI void GLAPIENTRY glColor3sv( const GLshort *v );
//GLAPI void GLAPIENTRY glColor3ubv( const GLubyte *v );
//GLAPI void GLAPIENTRY glColor3uiv( const GLuint *v );
//GLAPI void GLAPIENTRY glColor3usv( const GLushort *v );

//GLAPI void GLAPIENTRY glColor4bv( const GLbyte *v );
//GLAPI void GLAPIENTRY glColor4dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glColor4fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glColor4iv( const GLint *v );
//GLAPI void GLAPIENTRY glColor4sv( const GLshort *v );
//GLAPI void GLAPIENTRY glColor4ubv( const GLubyte *v );
//GLAPI void GLAPIENTRY glColor4uiv( const GLuint *v );
//GLAPI void GLAPIENTRY glColor4usv( const GLushort *v );


//GLAPI void GLAPIENTRY glTexCoord1d( GLdouble s );
func TexCoord1d(s float32) {
    C.glTexCoord1d(C.GLdouble(s))
}

//GLAPI void GLAPIENTRY glTexCoord1f( GLfloat s );
func TexCoord1f(s float32) {
    C.glTexCoord1f(C.GLfloat(s))
}

//GLAPI void GLAPIENTRY glTexCoord1i( GLint s );
func TexCoord1i(s int) {
    C.glTexCoord1i(C.GLint(s))
}

//GLAPI void GLAPIENTRY glTexCoord1s( GLshort s );
func TexCoord1s(s int16) {
    C.glTexCoord1s(C.GLshort(s))
}

//GLAPI void GLAPIENTRY glTexCoord2d( GLdouble s, GLdouble t );
func TexCoord2d(s, t float32) {
    C.glTexCoord2d(C.GLdouble(s), C.GLdouble(t))
}

//GLAPI void GLAPIENTRY glTexCoord2f( GLfloat s, GLfloat t );
func TexCoord2f(s, t float32) {
    C.glTexCoord2f(C.GLfloat(s), C.GLfloat(t))
}

//GLAPI void GLAPIENTRY glTexCoord2i( GLint s, GLint t );
func TexCoord2i(s, t int) {
    C.glTexCoord2i(C.GLint(s), C.GLint(t))
}

//GLAPI void GLAPIENTRY glTexCoord2s( GLshort s, GLshort t );
func TexCoord2s(s, t int16) {
    C.glTexCoord2s(C.GLshort(s), C.GLshort(t))
}

//GLAPI void GLAPIENTRY glTexCoord3d( GLdouble s, GLdouble t, GLdouble r );
func TexCoord3d(s, t, r float32) {
    C.glTexCoord3d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
}

//GLAPI void GLAPIENTRY glTexCoord3f( GLfloat s, GLfloat t, GLfloat r );
func TexCoord3f(s, t, r float32) {
    C.glTexCoord3f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
}

//GLAPI void GLAPIENTRY glTexCoord3i( GLint s, GLint t, GLint r );
func TexCoord3i(s, t, r int) {
    C.glTexCoord3i(C.GLint(s), C.GLint(t), C.GLint(r))
}

//GLAPI void GLAPIENTRY glTexCoord3s( GLshort s, GLshort t, GLshort r );
func TexCoord3s(s, t, r int16) {
    C.glTexCoord3s(C.GLshort(s), C.GLshort(t), C.GLshort(r))
}

//GLAPI void GLAPIENTRY glTexCoord4d( GLdouble s, GLdouble t, GLdouble r, GLdouble q );
func TexCoord4d(s, t, r, q float32) {
    C.glTexCoord4d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
}

//GLAPI void GLAPIENTRY glTexCoord4f( GLfloat s, GLfloat t, GLfloat r, GLfloat q );
func TexCoord4f(s, t, r, q float32) {
    C.glTexCoord4f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
}

//GLAPI void GLAPIENTRY glTexCoord4i( GLint s, GLint t, GLint r, GLint q );
func TexCoord4i(s, t, r, q int) {
    C.glTexCoord4i(C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
}

//GLAPI void GLAPIENTRY glTexCoord4s( GLshort s, GLshort t, GLshort r, GLshort q );
func TexCoord4s(s, t, r, q int) {
    C.glTexCoord4s(C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
}

// FIXME
//GLAPI void GLAPIENTRY glTexCoord1dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glTexCoord1fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glTexCoord1iv( const GLint *v );
//GLAPI void GLAPIENTRY glTexCoord1sv( const GLshort *v );

//GLAPI void GLAPIENTRY glTexCoord2dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glTexCoord2fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glTexCoord2iv( const GLint *v );
//GLAPI void GLAPIENTRY glTexCoord2sv( const GLshort *v );

//GLAPI void GLAPIENTRY glTexCoord3dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glTexCoord3fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glTexCoord3iv( const GLint *v );
//GLAPI void GLAPIENTRY glTexCoord3sv( const GLshort *v );

//GLAPI void GLAPIENTRY glTexCoord4dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glTexCoord4fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glTexCoord4iv( const GLint *v );
//GLAPI void GLAPIENTRY glTexCoord4sv( const GLshort *v );


//GLAPI void GLAPIENTRY glRasterPos2d( GLdouble x, GLdouble y );
func RasterPos2d(x, y float32) {
    C.glRasterPos2d(C.GLdouble(x), C.GLdouble(y))
}

//GLAPI void GLAPIENTRY glRasterPos2f( GLfloat x, GLfloat y );
func RasterPos2f(x, y float32) {
    C.glRasterPos2f(C.GLfloat(x), C.GLfloat(y))
}

//GLAPI void GLAPIENTRY glRasterPos2i( GLint x, GLint y );
func RasterPos2i(x, y int) {
    C.glRasterPos2i(C.GLint(x), C.GLint(y))
}

//GLAPI void GLAPIENTRY glRasterPos2s( GLshort x, GLshort y );
func RasterPos2s(x, y int16) {
    C.glRasterPos2s(C.GLshort(x), C.GLshort(y))
}

//GLAPI void GLAPIENTRY glRasterPos3d( GLdouble x, GLdouble y, GLdouble z );
func RasterPos3d(x, y, z float32) {
    C.glRasterPos3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//GLAPI void GLAPIENTRY glRasterPos3f( GLfloat x, GLfloat y, GLfloat z );
func RasterPos3f(x, y, z float32) {
    C.glRasterPos3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//GLAPI void GLAPIENTRY glRasterPos3i( GLint x, GLint y, GLint z );
func RasterPos3i(x, y, z int) {
    C.glRasterPos3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

//GLAPI void GLAPIENTRY glRasterPos3s( GLshort x, GLshort y, GLshort z );
func RasterPos3s(x, y, z int16) {
    C.glRasterPos3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

//GLAPI void GLAPIENTRY glRasterPos4d( GLdouble x, GLdouble y, GLdouble z, GLdouble w );
func RasterPos4d(x, y, z, w float32) {
    C.glRasterPos4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

//GLAPI void GLAPIENTRY glRasterPos4f( GLfloat x, GLfloat y, GLfloat z, GLfloat w );
func RasterPos4f(x, y, z, w float32) {
    C.glRasterPos4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

//GLAPI void GLAPIENTRY glRasterPos4i( GLint x, GLint y, GLint z, GLint w );
func RasterPos4i(x, y, z, w int) {
    C.glRasterPos4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

//GLAPI void GLAPIENTRY glRasterPos4s( GLshort x, GLshort y, GLshort z, GLshort w );
func RasterPos4s(x, y, z, w int16) {
    C.glRasterPos4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

// FIXME
//GLAPI void GLAPIENTRY glRasterPos2dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glRasterPos2fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glRasterPos2iv( const GLint *v );
//GLAPI void GLAPIENTRY glRasterPos2sv( const GLshort *v );

//GLAPI void GLAPIENTRY glRasterPos3dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glRasterPos3fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glRasterPos3iv( const GLint *v );
//GLAPI void GLAPIENTRY glRasterPos3sv( const GLshort *v );

//GLAPI void GLAPIENTRY glRasterPos4dv( const GLdouble *v );
//GLAPI void GLAPIENTRY glRasterPos4fv( const GLfloat *v );
//GLAPI void GLAPIENTRY glRasterPos4iv( const GLint *v );
//GLAPI void GLAPIENTRY glRasterPos4sv( const GLshort *v );


//GLAPI void GLAPIENTRY glRectd( GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2 );
func Rectd(x1, y1, x2, y2 float32) {
    C.glRectd(C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
}

//GLAPI void GLAPIENTRY glRectf( GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2 );
func Rectf(x1, y1, x2, y2 float32) {
    C.glRectf(C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
}

//GLAPI void GLAPIENTRY glRecti( GLint x1, GLint y1, GLint x2, GLint y2 );
func Recti(x1, y1, x2, y2 int) {
    C.glRecti(C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
}

//GLAPI void GLAPIENTRY glRects( GLshort x1, GLshort y1, GLshort x2, GLshort y2 );
func Rects(x1, y1, x2, y2 int16) {
    C.glRects(C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
}

// FIXME
//GLAPI void GLAPIENTRY glRectdv( const GLdouble *v1, const GLdouble *v2 );
//GLAPI void GLAPIENTRY glRectfv( const GLfloat *v1, const GLfloat *v2 );
//GLAPI void GLAPIENTRY glRectiv( const GLint *v1, const GLint *v2 );
//GLAPI void GLAPIENTRY glRectsv( const GLshort *v1, const GLshort *v2 );


// ------------------------------------------------------------------------- //
// Vertex Arrays  (1.1)
// ------------------------------------------------------------------------- //

// FIXME
//GLAPI void GLAPIENTRY glVertexPointer( GLint size, GLenum type, GLsizei stride, const GLvoid *ptr );

//GLAPI void GLAPIENTRY glNormalPointer( GLenum type, GLsizei stride, const GLvoid *ptr );

//GLAPI void GLAPIENTRY glColorPointer( GLint size, GLenum type, GLsizei stride, const GLvoid *ptr );

//GLAPI void GLAPIENTRY glIndexPointer( GLenum type, GLsizei stride, const GLvoid *ptr );

//GLAPI void GLAPIENTRY glTexCoordPointer( GLint size, GLenum type, GLsizei stride, const GLvoid *ptr );

//GLAPI void GLAPIENTRY glEdgeFlagPointer( GLsizei stride, const GLvoid *ptr );

//GLAPI void GLAPIENTRY glGetPointerv( GLenum pname, GLvoid **params );

//GLAPI void GLAPIENTRY glArrayElement( GLint i );
func ArrayElement(i int) {
    C.glArrayElement(C.GLint(i))
}

//GLAPI void GLAPIENTRY glDrawArrays( GLenum mode, GLint first, GLsizei count );
func DrawArrays(mode uint, first, count int) {
    C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

// FIXME
//GLAPI void GLAPIENTRY glDrawElements( GLenum mode, GLsizei count, GLenum type, const GLvoid *indices );

//GLAPI void GLAPIENTRY glInterleavedArrays( GLenum format, GLsizei stride, const GLvoid *pointer );

// ------------------------------------------------------------------------- //
// Lighting
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glShadeModel( GLenum mode );
func ShadeModel(mode uint) {
    C.glShadeModel(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glLightf( GLenum light, GLenum pname, GLfloat param );
func Lightf(light, pname uint, param float32) {
    C.glLightf(C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glLighti( GLenum light, GLenum pname, GLint param );
func Lighti(light, pname uint, param int) {
    C.glLighti(C.GLenum(light), C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glLightfv( GLenum light, GLenum pname, const GLfloat *params );
//GLAPI void GLAPIENTRY glLightiv( GLenum light, GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glGetLightfv( GLenum light, GLenum pname, GLfloat *params );
//GLAPI void GLAPIENTRY glGetLightiv( GLenum light, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glLightModelf( GLenum pname, GLfloat param );
func LightModelf(pname uint, param float32) {
    C.glLightModelf(C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glLightModeli( GLenum pname, GLint param );
func LightModeli(pname uint, param int) {
    C.glLightModeli(C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glLightModelfv( GLenum pname, const GLfloat *params );
//GLAPI void GLAPIENTRY glLightModeliv( GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glMaterialf( GLenum face, GLenum pname, GLfloat param );
func glMaterialf(face, pname uint, param float32) {
    C.glMaterialf(C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glMateriali( GLenum face, GLenum pname, GLint param );
func glMateriali(face, pname uint, param int) {
    C.glMateriali(C.GLenum(face), C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glMaterialfv( GLenum face, GLenum pname, const GLfloat *params );
//GLAPI void GLAPIENTRY glMaterialiv( GLenum face, GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glGetMaterialfv( GLenum face, GLenum pname, GLfloat *params );
//GLAPI void GLAPIENTRY glGetMaterialiv( GLenum face, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glColorMaterial( GLenum face, GLenum mode );
func ColorMaterial(face, mode uint) {
    C.glColorMaterial(C.GLenum(face), C.GLenum(mode))
}

// ------------------------------------------------------------------------- //
// Raster functions
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glPixelZoom( GLfloat xfactor, GLfloat yfactor );
func PixelZoom(xfactor, yfactor float32) {
    C.glPixelZoom(C.GLfloat(xfactor), C.GLfloat(yfactor))
}

//GLAPI void GLAPIENTRY glPixelStoref( GLenum pname, GLfloat param );
func PixelStoref(pname uint, param float32) {
    C.glPixelStoref(C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glPixelStorei( GLenum pname, GLint param );
func PixelStorei(pname uint, param int) {
    C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

//GLAPI void GLAPIENTRY glPixelTransferf( GLenum pname, GLfloat param );
func PixelTransferf(pname uint, param float32) {
    C.glPixelTransferf(C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glPixelTransferi( GLenum pname, GLint param );
func PixelTransferi(pname uint, param int) {
    C.glPixelTransferi(C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glPixelMapfv( GLenum map, GLsizei mapsize, const GLfloat *values );
//GLAPI void GLAPIENTRY glPixelMapuiv( GLenum map, GLsizei mapsize, const GLuint *values );
//GLAPI void GLAPIENTRY glPixelMapusv( GLenum map, GLsizei mapsize, const GLushort *values );

//GLAPI void GLAPIENTRY glGetPixelMapfv( GLenum map, GLfloat *values );
//GLAPI void GLAPIENTRY glGetPixelMapuiv( GLenum map, GLuint *values );
//GLAPI void GLAPIENTRY glGetPixelMapusv( GLenum map, GLushort *values );

//GLAPI void GLAPIENTRY glBitmap( GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, const GLubyte *bitmap );

//GLAPI void GLAPIENTRY glReadPixels( GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid *pixels );

//GLAPI void GLAPIENTRY glDrawPixels( GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid *pixels );

//GLAPI void GLAPIENTRY glCopyPixels( GLint x, GLint y, GLsizei width, GLsizei height, GLenum type );
func CopyPixels(x, y, width, height int, t uint) {
    C.glCopyPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(t))
}

// ------------------------------------------------------------------------- //
// Stenciling mapping
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glStencilFunc( GLenum func, GLint ref, GLuint mask );
func StencilFunc(f uint, ref int, mask uint) {
    C.glStencilFunc(C.GLenum(f), C.GLint(ref), C.GLuint(mask))
}

//GLAPI void GLAPIENTRY glStencilMask( GLuint mask );
func StencilMask(mask uint) {
    C.glStencilMask(C.GLuint(mask))
}

//GLAPI void GLAPIENTRY glStencilOp( GLenum fail, GLenum zfail, GLenum zpass );
func StencilOp(fail, zfail, zpass uint) {
    C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

//GLAPI void GLAPIENTRY glClearStencil( GLint s );
func ClearStencil(s int) {
    C.glClearStencil(C.GLint(s))
}

// ------------------------------------------------------------------------- //
// Texture mapping
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glTexGend( GLenum coord, GLenum pname, GLdouble param );
func TexGend(coord, pname uint, param float32) {
    C.glTexGend(C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
}

//GLAPI void GLAPIENTRY glTexGenf( GLenum coord, GLenum pname, GLfloat param );
func TexGenf(coord, pname uint, param float32) {
    C.glTexGenf(C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glTexGeni( GLenum coord, GLenum pname, GLint param );
func TexGeni(coord, pname uint, param int) {
    C.glTexGeni(C.GLenum(coord), C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glTexGendv( GLenum coord, GLenum pname, const GLdouble *params );
//GLAPI void GLAPIENTRY glTexGenfv( GLenum coord, GLenum pname, const GLfloat *params );
//GLAPI void GLAPIENTRY glTexGeniv( GLenum coord, GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glGetTexGendv( GLenum coord, GLenum pname, GLdouble *params );
//GLAPI void GLAPIENTRY glGetTexGenfv( GLenum coord, GLenum pname, GLfloat *params );
//GLAPI void GLAPIENTRY glGetTexGeniv( GLenum coord, GLenum pname, GLint *params );


//GLAPI void GLAPIENTRY glTexEnvf( GLenum target, GLenum pname, GLfloat param );
func TexEnvf(target, pname uint, param float32) {
    C.glTexEnvf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glTexEnvi( GLenum target, GLenum pname, GLint param );
func TexEnvi(target, pname uint, param int) {
    C.glTexEnvi(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glTexEnvfv( GLenum target, GLenum pname, const GLfloat *params );
//GLAPI void GLAPIENTRY glTexEnviv( GLenum target, GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glGetTexEnvfv( GLenum target, GLenum pname, GLfloat *params );
//GLAPI void GLAPIENTRY glGetTexEnviv( GLenum target, GLenum pname, GLint *params );


//GLAPI void GLAPIENTRY glTexParameterf( GLenum target, GLenum pname, GLfloat param );
func TexParameterf(target, pname uint, param float32) {
    C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glTexParameteri( GLenum target, GLenum pname, GLint param );
func TexParameteri(target, pname uint, param int) {
    C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glTexParameterfv( GLenum target, GLenum pname, const GLfloat *params );
//GLAPI void GLAPIENTRY glTexParameteriv( GLenum target, GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glGetTexParameterfv( GLenum target, GLenum pname, GLfloat *params);
//GLAPI void GLAPIENTRY glGetTexParameteriv( GLenum target, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glGetTexLevelParameterfv( GLenum target, GLint level, GLenum pname, GLfloat *params );
//GLAPI void GLAPIENTRY glGetTexLevelParameteriv( GLenum target, GLint level, GLenum pname, GLint *params );


//GLAPI void GLAPIENTRY glTexImage1D( GLenum target, GLint level, GLint internalFormat, GLsizei width, GLint border, GLenum format, GLenum type, const GLvoid *pixels );

//GLAPI void GLAPIENTRY glTexImage2D( GLenum target, GLint level, GLint internalFormat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid *pixels );

//GLAPI void GLAPIENTRY glGetTexImage( GLenum target, GLint level, GLenum format, GLenum type, GLvoid *pixels );


// ------------------------------------------------------------------------- //
// 1.1 functions
// ------------------------------------------------------------------------- //

// FIXME
//GLAPI void GLAPIENTRY glGenTextures( GLsizei n, GLuint *textures );

//GLAPI void GLAPIENTRY glDeleteTextures( GLsizei n, const GLuint *textures);

//GLAPI void GLAPIENTRY glBindTexture( GLenum target, GLuint texture );
func BindTexture(target, texture uint) {
    C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

// FIXME
//GLAPI void GLAPIENTRY glPrioritizeTextures( GLsizei n, const GLuint *textures, const GLclampf *priorities );

//GLAPI GLboolean GLAPIENTRY glAreTexturesResident( GLsizei n, const GLuint *textures, GLboolean *residences );

//GLAPI GLboolean GLAPIENTRY glIsTexture( GLuint texture );
func IsTexture(texture uint) (bool) {
    return C.glIsTexture(C.GLuint(texture)) != 0
}

// FIXME
//GLAPI void GLAPIENTRY glTexSubImage1D( GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, const GLvoid *pixels );


//GLAPI void GLAPIENTRY glTexSubImage2D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid *pixels );


//GLAPI void GLAPIENTRY glCopyTexImage1D( GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border );
func CopyTexImage1D(target uint, level int, internalformat uint, x, y, width, border int) {
    C.glCopyTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

//GLAPI void GLAPIENTRY glCopyTexImage2D( GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border );
func CopyTexImage2D(target uint, level int, internalformat uint, x, y, width, height, border int) {
    C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}


//GLAPI void GLAPIENTRY glCopyTexSubImage1D( GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width );
func CopyTexSubImage1D(target uint, level, xoffset, x, y, width int) {
    C.glCopyTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

//GLAPI void GLAPIENTRY glCopyTexSubImage2D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
func CopyTexSubImage2D(target uint, level, xoffset, yoffset, x, y, width, height int) {
    C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}


// ------------------------------------------------------------------------- //
// Evaluators
// ------------------------------------------------------------------------- //

// FIXME
//GLAPI void GLAPIENTRY glMap1d( GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, const GLdouble *points );
//GLAPI void GLAPIENTRY glMap1f( GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, const GLfloat *points );

//GLAPI void GLAPIENTRY glMap2d( GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, const GLdouble *points );
//GLAPI void GLAPIENTRY glMap2f( GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, const GLfloat *points );

//GLAPI void GLAPIENTRY glGetMapdv( GLenum target, GLenum query, GLdouble *v );
//GLAPI void GLAPIENTRY glGetMapfv( GLenum target, GLenum query, GLfloat *v );
//GLAPI void GLAPIENTRY glGetMapiv( GLenum target, GLenum query, GLint *v );

//GLAPI void GLAPIENTRY glEvalCoord1d( GLdouble u );
func EvalCoord1d(u float32) {
    C.glEvalCoord1d(C.GLdouble(u))
}

//GLAPI void GLAPIENTRY glEvalCoord1f( GLfloat u );
func EvalCoord1f(u float32) {
    C.glEvalCoord1f(C.GLfloat(u))
}

// FIXME
//GLAPI void GLAPIENTRY glEvalCoord1dv( const GLdouble *u );
//GLAPI void GLAPIENTRY glEvalCoord1fv( const GLfloat *u );

//GLAPI void GLAPIENTRY glEvalCoord2d( GLdouble u, GLdouble v );
func EvalCoord2d(u, v float32) {
    C.glEvalCoord2d(C.GLdouble(u), C.GLdouble(v))
}

//GLAPI void GLAPIENTRY glEvalCoord2f( GLfloat u, GLfloat v );
func EvalCoord2f(u, v float32) {
    C.glEvalCoord2f(C.GLfloat(u), C.GLfloat(v))
}

// FIXME
//GLAPI void GLAPIENTRY glEvalCoord2dv( const GLdouble *u );
//GLAPI void GLAPIENTRY glEvalCoord2fv( const GLfloat *u );

//GLAPI void GLAPIENTRY glMapGrid1d( GLint un, GLdouble u1, GLdouble u2 );
func MapGrid1d(un int, u1, u2 float32) {
    C.glMapGrid1d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
}

//GLAPI void GLAPIENTRY glMapGrid1f( GLint un, GLfloat u1, GLfloat u2 );
func MapGrid1f(un int, u1, u2 float32) {
    C.glMapGrid1f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
}

//GLAPI void GLAPIENTRY glMapGrid2d( GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2 );
func MapGrid2d(un int, u1, u2 float32, vn int, v1, v2 float32) {
    C.glMapGrid2d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
}

//GLAPI void GLAPIENTRY glMapGrid2f( GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2 );
func MapGrid2f(un int, u1, u2 float32, vn int, v1, v2 float32) {
    C.glMapGrid2f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
}

//GLAPI void GLAPIENTRY glEvalPoint1( GLint i );
func EvalPoint1(i int) {
    C.glEvalPoint1(C.GLint(i))
}

//GLAPI void GLAPIENTRY glEvalPoint2( GLint i, GLint j );
func EvalPoint2(i, j int) {
    C.glEvalPoint2(C.GLint(i), C.GLint(j))
}

//GLAPI void GLAPIENTRY glEvalMesh1( GLenum mode, GLint i1, GLint i2 );
func EvalMesh1(mode uint, i1, i2 int) {
    C.glEvalMesh1(C.GLenum(mode), C.GLint(i1), C.GLint(i2))
}

//GLAPI void GLAPIENTRY glEvalMesh2( GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2 );
func EvalMesh2(mode uint, i1, i2, j1, j2 int) {
    C.glEvalMesh2(C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
}

// ------------------------------------------------------------------------- //
// Fog
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glFogf( GLenum pname, GLfloat param );
func Fogf(pname uint, param float32) {
    C.glFogf(C.GLenum(pname), C.GLfloat(param))
}

//GLAPI void GLAPIENTRY glFogi( GLenum pname, GLint param );
func Fogi(pname uint, param int) {
    C.glFogi(C.GLenum(pname), C.GLint(param))
}

// FIXME
//GLAPI void GLAPIENTRY glFogfv( GLenum pname, const GLfloat *params );

//GLAPI void GLAPIENTRY glFogiv( GLenum pname, const GLint *params );


// ------------------------------------------------------------------------- //
// Selection and Feedback
// ------------------------------------------------------------------------- //

// FIXME
//GLAPI void GLAPIENTRY glFeedbackBuffer( GLsizei size, GLenum type, GLfloat *buffer );

//GLAPI void GLAPIENTRY glPassThrough( GLfloat token );
func PassThrough(token float32) {
    C.glPassThrough(C.GLfloat(token))
}

// FIXME
//GLAPI void GLAPIENTRY glSelectBuffer( GLsizei size, GLuint *buffer );

//GLAPI void GLAPIENTRY glInitNames( void );
func InitNames() {
    C.glInitNames()
}

//GLAPI void GLAPIENTRY glLoadName( GLuint name );
func LoadName(name uint) {
    C.glLoadName(C.GLuint(name))
}

//GLAPI void GLAPIENTRY glPushName( GLuint name );
func PushName(name uint) {
    C.glPushName(C.GLuint(name))
}

//GLAPI void GLAPIENTRY glPopName( void );
func PopName() {
    C.glPopName()
}

// ------------------------------------------------------------------------- //
// OpenGL 1.2
// ------------------------------------------------------------------------- //

// FIXME
//GLAPI void GLAPIENTRY glDrawRangeElements( GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, const GLvoid *indices );

//GLAPI void GLAPIENTRY glTexImage3D( GLenum target, GLint level, GLint internalFormat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, const GLvoid *pixels );

//GLAPI void GLAPIENTRY glTexSubImage3D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const GLvoid *pixels);

//GLAPI void GLAPIENTRY glCopyTexSubImage3D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height );
func CopyTexSubImage3D(target uint, level, xoffset, yoffset, zoffset, x, y, width, height int) {
    C.glCopyTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

/*
typedef void (APIENTRYP PFNGLDRAWRANGEELEMENTSPROC) (GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, const GLvoid *indices);
typedef void (APIENTRYP PFNGLTEXIMAGE3DPROC) (GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, const GLvoid *pixels);
typedef void (APIENTRYP PFNGLTEXSUBIMAGE3DPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const GLvoid *pixels);
typedef void (APIENTRYP PFNGLCOPYTEXSUBIMAGE3DPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
*/

// ------------------------------------------------------------------------- //
// GL_ARB_imaging
// ------------------------------------------------------------------------- //

// FIXME
//GLAPI void GLAPIENTRY glColorTable( GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, const GLvoid *table );

//GLAPI void GLAPIENTRY glColorSubTable( GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, const GLvoid *data );

//GLAPI void GLAPIENTRY glColorTableParameteriv(GLenum target, GLenum pname, const GLint *params);

//GLAPI void GLAPIENTRY glColorTableParameterfv(GLenum target, GLenum pname, const GLfloat *params);

//GLAPI void GLAPIENTRY glCopyColorSubTable( GLenum target, GLsizei start, GLint x, GLint y, GLsizei width );
func CopyColorSubTable(target uint, start, x, y, width int) {
    C.glCopyColorSubTable(C.GLenum(target), C.GLsizei(start), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

//GLAPI void GLAPIENTRY glCopyColorTable( GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width );
func CopyColorTable(target, internalformat uint, x, y, width int) {
    C.glCopyColorTable(C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

// FIXME
//GLAPI void GLAPIENTRY glGetColorTable( GLenum target, GLenum format, GLenum type, GLvoid *table );

//GLAPI void GLAPIENTRY glGetColorTableParameterfv( GLenum target, GLenum pname, GLfloat *params );

//GLAPI void GLAPIENTRY glGetColorTableParameteriv( GLenum target, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glBlendEquation( GLenum mode );
func BlendEquation(mode uint) {
    C.glBlendEquation(C.GLenum(mode))
}

//GLAPI void GLAPIENTRY glBlendColor( GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha );
func BlendColor(red, green, blue, alpha float32) {
    C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

//GLAPI void GLAPIENTRY glHistogram( GLenum target, GLsizei width, GLenum internalformat, GLboolean sink );
func Histogram(target uint, width int, internalformat uint, sink int8) {
    C.glHistogram(C.GLenum(target), C.GLsizei(width), C.GLenum(internalformat), C.GLboolean(sink))
}

//GLAPI void GLAPIENTRY glResetHistogram( GLenum target );
func ResetHistogram(target uint) {
    C.glResetHistogram(C.GLenum(target))
}

// FIXME
//GLAPI void GLAPIENTRY glGetHistogram( GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid *values );

//GLAPI void GLAPIENTRY glGetHistogramParameterfv( GLenum target, GLenum pname, GLfloat *params );

//GLAPI void GLAPIENTRY glGetHistogramParameteriv( GLenum target, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glMinmax( GLenum target, GLenum internalformat, GLboolean sink );
func Minmax(target, internalformat uint, sink int8) {
    C.glMinmax(C.GLenum(target), C.GLenum(internalformat), C.GLboolean(sink))
}

//GLAPI void GLAPIENTRY glResetMinmax( GLenum target );
func ResetMinmax(target uint) {
    C.glResetMinmax(C.GLenum(target))
}

// FIXME
//GLAPI void GLAPIENTRY glGetMinmax( GLenum target, GLboolean reset, GLenum format, GLenum types, GLvoid *values );

//GLAPI void GLAPIENTRY glGetMinmaxParameterfv( GLenum target, GLenum pname, GLfloat *params );

//GLAPI void GLAPIENTRY glGetMinmaxParameteriv( GLenum target, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glConvolutionFilter1D( GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, const GLvoid *image );

//GLAPI void GLAPIENTRY glConvolutionFilter2D( GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid *image );

//GLAPI void GLAPIENTRY glConvolutionParameterf( GLenum target, GLenum pname, GLfloat params );
func ConvolutionParameterf(target, pname uint, params float32) {
    C.glConvolutionParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(params))
}

// FIXME
//GLAPI void GLAPIENTRY glConvolutionParameterfv( GLenum target, GLenum pname, const GLfloat *params );

//GLAPI void GLAPIENTRY glConvolutionParameteri( GLenum target, GLenum pname, GLint params );
func ConvolutionParameteri(target, pname uint, params int) {
    C.glConvolutionParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(params))
}

// FIXME
//GLAPI void GLAPIENTRY glConvolutionParameteriv( GLenum target, GLenum pname, const GLint *params );

//GLAPI void GLAPIENTRY glCopyConvolutionFilter1D( GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width );
func CopyConvolutionFilter1D(target, internalformat uint, x, y, width int) {
    C.glCopyConvolutionFilter1D(C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

//GLAPI void GLAPIENTRY glCopyConvolutionFilter2D( GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height);
func CopyConvolutionFilter2D(target, internalformat uint, x, y, width, height int) {
    C.glCopyConvolutionFilter2D(C.GLenum(target), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// FIXME
//GLAPI void GLAPIENTRY glGetConvolutionFilter( GLenum target, GLenum format, GLenum type, GLvoid *image );

//GLAPI void GLAPIENTRY glGetConvolutionParameterfv( GLenum target, GLenum pname, GLfloat *params );

//GLAPI void GLAPIENTRY glGetConvolutionParameteriv( GLenum target, GLenum pname, GLint *params );

//GLAPI void GLAPIENTRY glSeparableFilter2D( GLenum target, GLenum internalformat, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid *row, const GLvoid *column );

//GLAPI void GLAPIENTRY glGetSeparableFilter( GLenum target, GLenum format, GLenum type, GLvoid *row, GLvoid *column, GLvoid *span );

/*
typedef void (APIENTRYP PFNGLBLENDCOLORPROC) (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
typedef void (APIENTRYP PFNGLBLENDEQUATIONPROC) (GLenum mode);
*/


// ------------------------------------------------------------------------- //
// OpenGL 1.3
// ------------------------------------------------------------------------- //

//GLAPI void GLAPIENTRY glActiveTexture( GLenum texture );
func ActiveTexture(texture uint) {
    C.glActiveTexture(C.GLenum(texture))
}

//GLAPI void GLAPIENTRY glClientActiveTexture( GLenum texture );
func ClientActiveTexture(texture uint) {
    C.glClientActiveTexture(C.GLenum(texture))
}

// FIXME
//GLAPI void GLAPIENTRY glCompressedTexImage1D( GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, const GLvoid *data );

//GLAPI void GLAPIENTRY glCompressedTexImage2D( GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const GLvoid *data );

//GLAPI void GLAPIENTRY glCompressedTexImage3D( GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid *data );

//GLAPI void GLAPIENTRY glCompressedTexSubImage1D( GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, const GLvoid *data );

//GLAPI void GLAPIENTRY glCompressedTexSubImage2D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const GLvoid *data );

//GLAPI void GLAPIENTRY glCompressedTexSubImage3D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid *data );

//GLAPI void GLAPIENTRY glGetCompressedTexImage( GLenum target, GLint lod, GLvoid *img );

//GLAPI void GLAPIENTRY glMultiTexCoord1d( GLenum target, GLdouble s );
func MultiTexCoord1d(target uint, s float32) {
    C.glMultiTexCoord1d(C.GLenum(target), C.GLdouble(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1dv( GLenum target, const GLdouble *v );

//GLAPI void GLAPIENTRY glMultiTexCoord1f( GLenum target, GLfloat s );
func MultiTexCoord1f(target uint, s float32) {
    C.glMultiTexCoord1f(C.GLenum(target), C.GLfloat(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1fv( GLenum target, const GLfloat *v );

//GLAPI void GLAPIENTRY glMultiTexCoord1i( GLenum target, GLint s );
func MultiTexCoord1i(target uint, s int) {
    C.glMultiTexCoord1i(C.GLenum(target), C.GLint(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1iv( GLenum target, const GLint *v );

//GLAPI void GLAPIENTRY glMultiTexCoord1s( GLenum target, GLshort s );
func MultiTexCoord1s(target uint, s int16) {
    C.glMultiTexCoord1s(C.GLenum(target), C.GLshort(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1sv( GLenum target, const GLshort *v );

//GLAPI void GLAPIENTRY glMultiTexCoord2d( GLenum target, GLdouble s, GLdouble t );
func MultiTexCoord2d(target uint, s, t float32) {
    C.glMultiTexCoord2d(C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2dv( GLenum target, const GLdouble *v );

//GLAPI void GLAPIENTRY glMultiTexCoord2f( GLenum target, GLfloat s, GLfloat t );
func MultiTexCoord2f(target uint, s, t float32) {
    C.glMultiTexCoord2f(C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2fv( GLenum target, const GLfloat *v );

//GLAPI void GLAPIENTRY glMultiTexCoord2i( GLenum target, GLint s, GLint t );
func MultiTexCoord2i(target uint, s, t int) {
    C.glMultiTexCoord2i(C.GLenum(target), C.GLint(s), C.GLint(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2iv( GLenum target, const GLint *v );

//GLAPI void GLAPIENTRY glMultiTexCoord2s( GLenum target, GLshort s, GLshort t );
func MultiTexCoord2s(target uint, s, t int16) {
    C.glMultiTexCoord2s(C.GLenum(target), C.GLshort(s), C.GLshort(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2sv( GLenum target, const GLshort *v );

//GLAPI void GLAPIENTRY glMultiTexCoord3d( GLenum target, GLdouble s, GLdouble t, GLdouble r );
func MultiTexCoord3d(target uint, s, t, r float32) {
    C.glMultiTexCoord3d(C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3dv( GLenum target, const GLdouble *v );

//GLAPI void GLAPIENTRY glMultiTexCoord3f( GLenum target, GLfloat s, GLfloat t, GLfloat r );
func MultiTexCoord3f(target uint, s, t, r float32) {
    C.glMultiTexCoord3f(C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3fv( GLenum target, const GLfloat *v );

//GLAPI void GLAPIENTRY glMultiTexCoord3i( GLenum target, GLint s, GLint t, GLint r );
func MultiTexCoord3i(target uint, s, t, r int) {
    C.glMultiTexCoord3i(C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3iv( GLenum target, const GLint *v );

//GLAPI void GLAPIENTRY glMultiTexCoord3s( GLenum target, GLshort s, GLshort t, GLshort r );
func MultiTexCoord3s(target uint, s, t, r int16) {
    C.glMultiTexCoord3s(C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3sv( GLenum target, const GLshort *v );

//GLAPI void GLAPIENTRY glMultiTexCoord4d( GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q );
func MultiTexCoord4d(target uint, s, t, r, q float32) {
    C.glMultiTexCoord4d(C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4dv( GLenum target, const GLdouble *v );

//GLAPI void GLAPIENTRY glMultiTexCoord4f( GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q );
func MultiTexCoord4f(target uint, s, t, r, q float32) {
    C.glMultiTexCoord4f(C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4fv( GLenum target, const GLfloat *v );

//GLAPI void GLAPIENTRY glMultiTexCoord4i( GLenum target, GLint s, GLint t, GLint r, GLint q );
func MultiTexCoord4i(target uint, s, t, r, q int) {
    C.glMultiTexCoord4i(C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4iv( GLenum target, const GLint *v );

//GLAPI void GLAPIENTRY glMultiTexCoord4s( GLenum target, GLshort s, GLshort t, GLshort r, GLshort q );
func MultiTexCoord4s(target uint, s, t, r, q int16) {
    C.glMultiTexCoord4s(C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4sv( GLenum target, const GLshort *v );

// FIXME
//GLAPI void GLAPIENTRY glLoadTransposeMatrixd( const GLdouble m[16] );

//GLAPI void GLAPIENTRY glLoadTransposeMatrixf( const GLfloat m[16] );

//GLAPI void GLAPIENTRY glMultTransposeMatrixd( const GLdouble m[16] );

//GLAPI void GLAPIENTRY glMultTransposeMatrixf( const GLfloat m[16] );

//GLAPI void GLAPIENTRY glSampleCoverage( GLclampf value, GLboolean invert );
func SampleCoverage(value float32, invert int8) {
    C.glSampleCoverage(C.GLclampf(value), C.GLboolean(invert))
}

/*
typedef void (APIENTRYP PFNGLACTIVETEXTUREPROC) (GLenum texture);
typedef void (APIENTRYP PFNGLSAMPLECOVERAGEPROC) (GLclampf value, GLboolean invert);
typedef void (APIENTRYP PFNGLCOMPRESSEDTEXIMAGE3DPROC) (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid *data);
typedef void (APIENTRYP PFNGLCOMPRESSEDTEXIMAGE2DPROC) (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const GLvoid *data);
typedef void (APIENTRYP PFNGLCOMPRESSEDTEXIMAGE1DPROC) (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, const GLvoid *data);
typedef void (APIENTRYP PFNGLCOMPRESSEDTEXSUBIMAGE3DPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid *data);
typedef void (APIENTRYP PFNGLCOMPRESSEDTEXSUBIMAGE2DPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const GLvoid *data);
typedef void (APIENTRYP PFNGLCOMPRESSEDTEXSUBIMAGE1DPROC) (GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, const GLvoid *data);
typedef void (APIENTRYP PFNGLGETCOMPRESSEDTEXIMAGEPROC) (GLenum target, GLint level, GLvoid *img);
*/

//GLAPI void GLAPIENTRY glActiveTextureARB(GLenum texture);
func ActiveTextureARB(texture uint) {
    C.glActiveTextureARB(C.GLenum(texture))
}

//GLAPI void GLAPIENTRY glClientActiveTextureARB(GLenum texture);
func ClientActiveTextureARB(texture uint) {
    C.glClientActiveTextureARB(C.GLenum(texture))
}

//GLAPI void GLAPIENTRY glMultiTexCoord1dARB(GLenum target, GLdouble s);
func MultiTexCoord1dARB(target uint, s float32) {
    C.glMultiTexCoord1dARB(C.GLenum(target), C.GLdouble(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1dvARB(GLenum target, const GLdouble *v);
//GLAPI void GLAPIENTRY glMultiTexCoord1fARB(GLenum target, GLfloat s);
func MultiTexCoord1fARB(target uint, s float32) {
    C.glMultiTexCoord1fARB(C.GLenum(target), C.GLfloat(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1fvARB(GLenum target, const GLfloat *v);
//GLAPI void GLAPIENTRY glMultiTexCoord1iARB(GLenum target, GLint s);
func MultiTexCoord1iARB(target uint, s int) {
    C.glMultiTexCoord1iARB(C.GLenum(target), C.GLint(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1ivARB(GLenum target, const GLint *v);
//GLAPI void GLAPIENTRY glMultiTexCoord1sARB(GLenum target, GLshort s);
func MultiTexCoord1sARB(target uint, s int16) {
    C.glMultiTexCoord1sARB(C.GLenum(target), C.GLshort(s))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord1svARB(GLenum target, const GLshort *v);
//GLAPI void GLAPIENTRY glMultiTexCoord2dARB(GLenum target, GLdouble s, GLdouble t);
func MultiTexCoord2dARB(target uint, s, t float32) {
    C.glMultiTexCoord2dARB(C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2dvARB(GLenum target, const GLdouble *v);
//GLAPI void GLAPIENTRY glMultiTexCoord2fARB(GLenum target, GLfloat s, GLfloat t);
func MultiTexCoord2fARB(target uint, s, t float32) {
    C.glMultiTexCoord2fARB(C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2fvARB(GLenum target, const GLfloat *v);
//GLAPI void GLAPIENTRY glMultiTexCoord2iARB(GLenum target, GLint s, GLint t);
func MultiTexCoord2iARB(target uint, s, t float32) {
    C.glMultiTexCoord2iARB(C.GLenum(target), C.GLint(s), C.GLint(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2ivARB(GLenum target, const GLint *v);
//GLAPI void GLAPIENTRY glMultiTexCoord2sARB(GLenum target, GLshort s, GLshort t);
func MultiTexCoord2sARB(target uint, s, t float32) {
    C.glMultiTexCoord2sARB(C.GLenum(target), C.GLshort(s), C.GLshort(t))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord2svARB(GLenum target, const GLshort *v);
//GLAPI void GLAPIENTRY glMultiTexCoord3dARB(GLenum target, GLdouble s, GLdouble t, GLdouble r);
func MultiTexCoord3dARB(target uint, s, t, r float32) {
    C.glMultiTexCoord3dARB(C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3dvARB(GLenum target, const GLdouble *v);
//GLAPI void GLAPIENTRY glMultiTexCoord3fARB(GLenum target, GLfloat s, GLfloat t, GLfloat r);
func MultiTexCoord3fARB(target uint, s, t, r float32) {
    C.glMultiTexCoord3fARB(C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3fvARB(GLenum target, const GLfloat *v);
//GLAPI void GLAPIENTRY glMultiTexCoord3iARB(GLenum target, GLint s, GLint t, GLint r);
func MultiTexCoord3iARB(target uint, s, t, r int) {
    C.glMultiTexCoord3iARB(C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3ivARB(GLenum target, const GLint *v);
//GLAPI void GLAPIENTRY glMultiTexCoord3sARB(GLenum target, GLshort s, GLshort t, GLshort r);
func MultiTexCoord3sARB(target uint, s, t, r int16) {
    C.glMultiTexCoord3sARB(C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord3svARB(GLenum target, const GLshort *v);
//GLAPI void GLAPIENTRY glMultiTexCoord4dARB(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q);
func MultiTexCoord4dARB(target uint, s, t, r, q float32) {
    C.glMultiTexCoord4dARB(C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4dvARB(GLenum target, const GLdouble *v);
//GLAPI void GLAPIENTRY glMultiTexCoord4fARB(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q);
func MultiTexCoord4fARB(target uint, s, t, r, q float32) {
    C.glMultiTexCoord4fARB(C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4fvARB(GLenum target, const GLfloat *v);
//GLAPI void GLAPIENTRY glMultiTexCoord4iARB(GLenum target, GLint s, GLint t, GLint r, GLint q);
func MultiTexCoord4iARB(target uint, s, t, r, q int) {
    C.glMultiTexCoord4iARB(C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4ivARB(GLenum target, const GLint *v);
//GLAPI void GLAPIENTRY glMultiTexCoord4sARB(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q);
func MultiTexCoord4sARB(target uint, s, t, r, q int16) {
    C.glMultiTexCoord4sARB(C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
}

// FIXME
//GLAPI void GLAPIENTRY glMultiTexCoord4svARB(GLenum target, const GLshort *v);









































// Dunno if anything below is even needed, or supported on other GL platforms.

/*
typedef void (APIENTRYP PFNGLACTIVETEXTUREARBPROC) (GLenum texture);
typedef void (APIENTRYP PFNGLCLIENTACTIVETEXTUREARBPROC) (GLenum texture);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1DARBPROC) (GLenum target, GLdouble s);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1DVARBPROC) (GLenum target, const GLdouble *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1FARBPROC) (GLenum target, GLfloat s);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1FVARBPROC) (GLenum target, const GLfloat *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1IARBPROC) (GLenum target, GLint s);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1IVARBPROC) (GLenum target, const GLint *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1SARBPROC) (GLenum target, GLshort s);
typedef void (APIENTRYP PFNGLMULTITEXCOORD1SVARBPROC) (GLenum target, const GLshort *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2DARBPROC) (GLenum target, GLdouble s, GLdouble t);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2DVARBPROC) (GLenum target, const GLdouble *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2FARBPROC) (GLenum target, GLfloat s, GLfloat t);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2FVARBPROC) (GLenum target, const GLfloat *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2IARBPROC) (GLenum target, GLint s, GLint t);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2IVARBPROC) (GLenum target, const GLint *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2SARBPROC) (GLenum target, GLshort s, GLshort t);
typedef void (APIENTRYP PFNGLMULTITEXCOORD2SVARBPROC) (GLenum target, const GLshort *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3DARBPROC) (GLenum target, GLdouble s, GLdouble t, GLdouble r);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3DVARBPROC) (GLenum target, const GLdouble *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3FARBPROC) (GLenum target, GLfloat s, GLfloat t, GLfloat r);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3FVARBPROC) (GLenum target, const GLfloat *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3IARBPROC) (GLenum target, GLint s, GLint t, GLint r);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3IVARBPROC) (GLenum target, const GLint *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3SARBPROC) (GLenum target, GLshort s, GLshort t, GLshort r);
typedef void (APIENTRYP PFNGLMULTITEXCOORD3SVARBPROC) (GLenum target, const GLshort *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4DARBPROC) (GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4DVARBPROC) (GLenum target, const GLdouble *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4FARBPROC) (GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4FVARBPROC) (GLenum target, const GLfloat *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4IARBPROC) (GLenum target, GLint s, GLint t, GLint r, GLint q);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4IVARBPROC) (GLenum target, const GLint *v);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4SARBPROC) (GLenum target, GLshort s, GLshort t, GLshort r, GLshort q);
typedef void (APIENTRYP PFNGLMULTITEXCOORD4SVARBPROC) (GLenum target, const GLshort *v);

#endif // GL_ARB_multitexture //



//
 * Define this token if you want "old-style" header file behaviour (extensions
 * defined in gl.h).  Otherwise, extensions will be included from glext.h.
 //
#if defined(GL_GLEXT_LEGACY)

// All extensions that used to be here are now found in glext.h //

#else  // GL_GLEXT_LEGACY //

#include <GL/glext.h>

#endif  // GL_GLEXT_LEGACY //



#if GL_ARB_shader_objects

#ifndef GL_MESA_shader_debug
#define GL_MESA_shader_debug 1

#define GL_DEBUG_OBJECT_MESA              0x8759
#define GL_DEBUG_PRINT_MESA               0x875A
#define GL_DEBUG_ASSERT_MESA              0x875B

//GLAPI GLhandleARB GLAPIENTRY glCreateDebugObjectMESA (void);
//GLAPI void GLAPIENTRY glClearDebugLogMESA (GLhandleARB obj, GLenum logType, GLenum shaderType);
//GLAPI void GLAPIENTRY glGetDebugLogMESA (GLhandleARB obj, GLenum logType, GLenum shaderType, GLsizei maxLength,
                                         GLsizei *length, GLcharARB *debugLog);
//GLAPI GLsizei GLAPIENTRY glGetDebugLogLengthMESA (GLhandleARB obj, GLenum logType, GLenum shaderType);

#endif // GL_MESA_shader_debug //

#endif // GL_ARB_shader_objects //


//
 * ???. GL_MESA_packed_depth_stencil
 * XXX obsolete
 //
#ifndef GL_MESA_packed_depth_stencil
#define GL_MESA_packed_depth_stencil 1

#define GL_DEPTH_STENCIL_MESA			0x8750
#define GL_UNSIGNED_INT_24_8_MESA		0x8751
#define GL_UNSIGNED_INT_8_24_REV_MESA		0x8752
#define GL_UNSIGNED_SHORT_15_1_MESA		0x8753
#define GL_UNSIGNED_SHORT_1_15_REV_MESA		0x8754

#endif // GL_MESA_packed_depth_stencil //


#ifndef GL_MESA_program_debug
#define GL_MESA_program_debug 1

#define GL_FRAGMENT_PROGRAM_POSITION_MESA       0x8bb0
#define GL_FRAGMENT_PROGRAM_CALLBACK_MESA       0x8bb1
#define GL_FRAGMENT_PROGRAM_CALLBACK_FUNC_MESA  0x8bb2
#define GL_FRAGMENT_PROGRAM_CALLBACK_DATA_MESA  0x8bb3
#define GL_VERTEX_PROGRAM_POSITION_MESA         0x8bb4
#define GL_VERTEX_PROGRAM_CALLBACK_MESA         0x8bb5
#define GL_VERTEX_PROGRAM_CALLBACK_FUNC_MESA    0x8bb6
#define GL_VERTEX_PROGRAM_CALLBACK_DATA_MESA    0x8bb7

typedef void (*GLprogramcallbackMESA)(GLenum target, GLvoid *data);

//GLAPI void GLAPIENTRY glProgramCallbackMESA(GLenum target, GLprogramcallbackMESA callback, GLvoid *data);

//GLAPI void GLAPIENTRY glGetProgramRegisterfvMESA(GLenum target, GLsizei len, const GLubyte *name, GLfloat *v);

#endif // GL_MESA_program_debug //


#ifndef GL_MESA_texture_array
#define GL_MESA_texture_array 1

// GL_MESA_texture_array uses the same enum values as GL_EXT_texture_array.
 //
#ifndef GL_EXT_texture_array

#ifdef GL_GLEXT_PROTOTYPES
//GLAPI void APIENTRY glFramebufferTextureLayerEXT(GLenum target,
    GLenum attachment, GLuint texture, GLint level, GLint layer);
#endif // GL_GLEXT_PROTOTYPES //

#if 0
// (temporarily) disabled because of collision with typedef in glext.h
 * that happens if apps include both gl.h and glext.h
 //
typedef void (APIENTRYP PFNGLFRAMEBUFFERTEXTURELAYEREXTPROC) (GLenum target,
    GLenum attachment, GLuint texture, GLint level, GLint layer);
#endif

#define GL_TEXTURE_1D_ARRAY_EXT         0x8C18
#define GL_PROXY_TEXTURE_1D_ARRAY_EXT   0x8C19
#define GL_TEXTURE_2D_ARRAY_EXT         0x8C1A
#define GL_PROXY_TEXTURE_2D_ARRAY_EXT   0x8C1B
#define GL_TEXTURE_BINDING_1D_ARRAY_EXT 0x8C1C
#define GL_TEXTURE_BINDING_2D_ARRAY_EXT 0x8C1D
#define GL_MAX_ARRAY_TEXTURE_LAYERS_EXT 0x88FF
#define GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER_EXT 0x8CD4
#endif

#endif


#ifndef GL_ATI_blend_equation_separate
#define GL_ATI_blend_equation_separate 1

#define GL_ALPHA_BLEND_EQUATION_ATI	        0x883D

//GLAPI void GLAPIENTRY glBlendEquationSeparateATI( GLenum modeRGB, GLenum modeA );
typedef void (APIENTRYP PFNGLBLENDEQUATIONSEPARATEATIPROC) (GLenum modeRGB, GLenum modeA);

#endif // GL_ATI_blend_equation_separate //


// GL_OES_EGL_image //
#ifndef GL_OES_EGL_image
typedef void* GLeglImageOES;
#endif

#ifndef GL_OES_EGL_image
#define GL_OES_EGL_image 1
#ifdef GL_GLEXT_PROTOTYPES
//GLAPI void APIENTRY glEGLImageTargetTexture2DOES (GLenum target, GLeglImageOES image);
//GLAPI void APIENTRY glEGLImageTargetRenderbufferStorageOES (GLenum target, GLeglImageOES image);
#endif
typedef void (APIENTRYP PFNGLEGLIMAGETARGETTEXTURE2DOESPROC) (GLenum target, GLeglImageOES image);
typedef void (APIENTRYP PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOESPROC) (GLenum target, GLeglImageOES image);
#endif
*/
