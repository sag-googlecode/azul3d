// Automatically generated OpenGL binding.
// 
// Categories in this package: 
// 
// HP_convolution_border_modes: http://www.opengl.org/registry/specs/HP/convolution_border_modes.txt
// 
// HP_image_transform: http://www.opengl.org/registry/specs/HP/image_transform.txt
// 
// HP_occlusion_test: http://www.opengl.org/registry/specs/HP/occlusion_test.txt
// 
// HP_texture_lighting: http://www.opengl.org/registry/specs/HP/texture_lighting.txt
// 
package hp

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// 
// #include <stdlib.h>
// #if defined(__APPLE__)
// #include <dlfcn.h>
// #elif defined(_WIN32)
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
// #else
// #include <X11/Xlib.h>
// #include <GL/glx.h>
// #endif
// 
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// 
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef int GLsizei;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef unsigned short GLhalf;
// typedef float GLfloat;
// typedef float GLclampf;
// typedef double GLdouble;
// typedef double GLclampd;
// typedef void GLvoid;
// 
// #include <stddef.h>
// #ifndef GL_VERSION_2_0
// /* GL type for program/shader text */
// typedef char GLchar;
// #endif
// 
// #ifndef GL_VERSION_1_5
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// #endif
// 
// #ifndef GL_ARB_vertex_buffer_object
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptrARB;
// typedef ptrdiff_t GLsizeiptrARB;
// #endif
// 
// #ifndef GL_ARB_shader_objects
// /* GL types for program/shader text and shader object handles */
// typedef char GLcharARB;
// typedef unsigned int GLhandleARB;
// #endif
// 
// /* GL type for "half" precision (s10e5) float data in host memory */
// #ifndef GL_ARB_half_float_pixel
// typedef unsigned short GLhalfARB;
// #endif
// 
// #ifndef GL_NV_half_float
// typedef unsigned short GLhalfNV;
// #endif
// 
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// 
// #ifndef GL_EXT_timer_query
// typedef int64_t GLint64EXT;
// typedef uint64_t GLuint64EXT;
// #endif
// 
// #ifndef GL_ARB_sync
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef struct __GLsync *GLsync;
// #endif
// 
// #ifndef GL_ARB_cl_event
// /* These incomplete types let us declare types compatible with OpenCL's cl_context and cl_event */
// struct _cl_context;
// struct _cl_event;
// #endif
// 
// #ifndef GL_ARB_debug_output
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// 
// #ifndef GL_AMD_debug_output
// typedef void (APIENTRY *GLDEBUGPROCAMD)(GLuint id,GLenum category,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// 
// #ifndef GL_NV_vdpau_interop
// typedef GLintptr GLvdpauSurfaceNV;
// #endif
// 
// #ifdef _WIN32
// static HMODULE opengl32 = NULL;
// #endif
// 
// static void* goglGetProcAddress(const char* name) { 
// #ifdef __APPLE__
// 	return dlsym(RTLD_DEFAULT, name);
// #elif _WIN32
// 	void* pf = wglGetProcAddress((LPCSTR)name);
// 	if(pf) {
// 		return pf;
// 	}
// 	if(opengl32 == NULL) {
// 		opengl32 = LoadLibraryA("opengl32.dll");
// 	}
// 	return GetProcAddress(opengl32, (LPCSTR)name);
// #else
// 	return glXGetProcAddress((const GLubyte*)name);
// #endif
// }
// 
// //  HP_convolution_border_modes
// //  HP_image_transform
// void (APIENTRYP ptrglImageTransformParameteriHP)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglImageTransformParameterfHP)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglImageTransformParameterivHP)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglImageTransformParameterfvHP)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetImageTransformParameterivHP)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetImageTransformParameterfvHP)(GLenum target, GLenum pname, GLfloat* params);
// //  HP_occlusion_test
// //  HP_texture_lighting
// 
// //  HP_convolution_border_modes
// //  HP_image_transform
// void goglImageTransformParameteriHP(GLenum target, GLenum pname, GLint param) {
// 	(*ptrglImageTransformParameteriHP)(target, pname, param);
// }
// void goglImageTransformParameterfHP(GLenum target, GLenum pname, GLfloat param) {
// 	(*ptrglImageTransformParameterfHP)(target, pname, param);
// }
// void goglImageTransformParameterivHP(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglImageTransformParameterivHP)(target, pname, params);
// }
// void goglImageTransformParameterfvHP(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglImageTransformParameterfvHP)(target, pname, params);
// }
// void goglGetImageTransformParameterivHP(GLenum target, GLenum pname, GLint* params) {
// 	(*ptrglGetImageTransformParameterivHP)(target, pname, params);
// }
// void goglGetImageTransformParameterfvHP(GLenum target, GLenum pname, GLfloat* params) {
// 	(*ptrglGetImageTransformParameterfvHP)(target, pname, params);
// }
// //  HP_occlusion_test
// //  HP_texture_lighting
// 
// int init_HP_convolution_border_modes() {
// 	return 0;
// }
// int init_HP_image_transform() {
// 	ptrglImageTransformParameteriHP = goglGetProcAddress("glImageTransformParameteriHP");
// 	if(ptrglImageTransformParameteriHP == NULL) return 1;
// 	ptrglImageTransformParameterfHP = goglGetProcAddress("glImageTransformParameterfHP");
// 	if(ptrglImageTransformParameterfHP == NULL) return 1;
// 	ptrglImageTransformParameterivHP = goglGetProcAddress("glImageTransformParameterivHP");
// 	if(ptrglImageTransformParameterivHP == NULL) return 1;
// 	ptrglImageTransformParameterfvHP = goglGetProcAddress("glImageTransformParameterfvHP");
// 	if(ptrglImageTransformParameterfvHP == NULL) return 1;
// 	ptrglGetImageTransformParameterivHP = goglGetProcAddress("glGetImageTransformParameterivHP");
// 	if(ptrglGetImageTransformParameterivHP == NULL) return 1;
// 	ptrglGetImageTransformParameterfvHP = goglGetProcAddress("glGetImageTransformParameterfvHP");
// 	if(ptrglGetImageTransformParameterfvHP == NULL) return 1;
// 	return 0;
// }
// int init_HP_occlusion_test() {
// 	return 0;
// }
// int init_HP_texture_lighting() {
// 	return 0;
// }
// 
import "C"
import "unsafe"
import "errors"

type (
	Enum     C.GLenum
	Boolean  C.GLboolean
	Bitfield C.GLbitfield
	Byte     C.GLbyte
	Short    C.GLshort
	Int      C.GLint
	Sizei    C.GLsizei
	Ubyte    C.GLubyte
	Ushort   C.GLushort
	Uint     C.GLuint
	Half     C.GLhalf
	Float    C.GLfloat
	Clampf   C.GLclampf
	Double   C.GLdouble
	Clampd   C.GLclampd
	Char     C.GLchar
	Pointer  unsafe.Pointer
	Sync     C.GLsync
	Int64    C.GLint64
	Uint64   C.GLuint64
	Intptr   C.GLintptr
	Sizeiptr C.GLsizeiptr
)

// HP_convolution_border_modes
const (
	CONSTANT_BORDER_HP = 0x8151
	CONVOLUTION_BORDER_COLOR_HP = 0x8154
	IGNORE_BORDER_HP = 0x8150
	REPLICATE_BORDER_HP = 0x8153
)
// HP_image_transform
const (
	AVERAGE_HP = 0x8160
	CUBIC_HP = 0x815F
	IMAGE_CUBIC_WEIGHT_HP = 0x815E
	IMAGE_MAG_FILTER_HP = 0x815C
	IMAGE_MIN_FILTER_HP = 0x815D
	IMAGE_ROTATE_ANGLE_HP = 0x8159
	IMAGE_ROTATE_ORIGIN_X_HP = 0x815A
	IMAGE_ROTATE_ORIGIN_Y_HP = 0x815B
	IMAGE_SCALE_X_HP = 0x8155
	IMAGE_SCALE_Y_HP = 0x8156
	IMAGE_TRANSFORM_2D_HP = 0x8161
	IMAGE_TRANSLATE_X_HP = 0x8157
	IMAGE_TRANSLATE_Y_HP = 0x8158
	POST_IMAGE_TRANSFORM_COLOR_TABLE_HP = 0x8162
	PROXY_POST_IMAGE_TRANSFORM_COLOR_TABLE_HP = 0x8163
)
// HP_occlusion_test
const (
	OCCLUSION_TEST_HP = 0x8165
	OCCLUSION_TEST_RESULT_HP = 0x8166
)
// HP_texture_lighting
const (
	TEXTURE_LIGHTING_MODE_HP = 0x8167
	TEXTURE_POST_SPECULAR_HP = 0x8168
	TEXTURE_PRE_SPECULAR_HP = 0x8169
)
// HP_convolution_border_modes

// HP_image_transform

func ImageTransformParameteriHP(target Enum, pname Enum, param Int)  {
	C.goglImageTransformParameteriHP((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func ImageTransformParameterfHP(target Enum, pname Enum, param Float)  {
	C.goglImageTransformParameterfHP((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func ImageTransformParameterivHP(target Enum, pname Enum, params *Int)  {
	C.goglImageTransformParameterivHP((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func ImageTransformParameterfvHP(target Enum, pname Enum, params *Float)  {
	C.goglImageTransformParameterfvHP((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
func GetImageTransformParameterivHP(target Enum, pname Enum, params *Int)  {
	C.goglGetImageTransformParameterivHP((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}
func GetImageTransformParameterfvHP(target Enum, pname Enum, params *Float)  {
	C.goglGetImageTransformParameterfvHP((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}
// HP_occlusion_test

// HP_texture_lighting

func InitHpConvolutionBorderModes() error {
	var ret C.int
	if ret = C.init_HP_convolution_border_modes(); ret != 0 {
		return errors.New("unable to initialize HP_convolution_border_modes")
	}
	return nil
}
func InitHpImageTransform() error {
	var ret C.int
	if ret = C.init_HP_image_transform(); ret != 0 {
		return errors.New("unable to initialize HP_image_transform")
	}
	return nil
}
func InitHpOcclusionTest() error {
	var ret C.int
	if ret = C.init_HP_occlusion_test(); ret != 0 {
		return errors.New("unable to initialize HP_occlusion_test")
	}
	return nil
}
func InitHpTextureLighting() error {
	var ret C.int
	if ret = C.init_HP_texture_lighting(); ret != 0 {
		return errors.New("unable to initialize HP_texture_lighting")
	}
	return nil
}
// EOF