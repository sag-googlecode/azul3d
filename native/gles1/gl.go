// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// This source file was automatically generated using glwrap.
// +build arm android
// +build !opengl_debug

// Package opengl implements Go bindings to OpenGL ES
//
// Debugging OpenGL applications is made easy by using the build tag for this
// package "opengl_debug", that is:
//
// go install -tags "opengl_debug" my/package/path
//
// Debugging using this build tag has both advantages and disadvantages:
//
// Advantage: glGetError() is ran directly after each OpenGL function call for
// you, and in the event of an error a panic will occur in addition to a full
// OpenGL stack trace being dumped to stdout before the panic occurs.
//
// Advantage: Batching cannot be turned on, this can help identify batching
// related bugs (due to incorrect memory management).
//
// Disadvantage: Applications compiled with the opengl_debug build tag are slow
// and will not run at full speed due to both disabled batching and the
// additional glGetError call for each OpenGL call.
//
package opengl

/*
#cgo linux LDFLAGS: -lGL
#include "gl.h"
*/
import "C"

import(
	"unsafe"
	"strconv"
	"strings"
)

// Context represents a single OpenGL context's API access.
//
// Since CGO (Go to C) calls are expensive in large quantities as they require
// a context switch due to Go's and C's incompatibile ABI, and since even
// modern OpenGL applications can make several hundred thousands of calls per
// frame, it can become a large bottleneck.
//
// To help remedy this specific bottleneck, batching can be used. When OpenGL
// calls are made and batching is turned on, instead of calls being executed
// immedietly they are stored in a queue and executed later upon a single call
// to the Execute() method.
//
// When batching is turned off and OpenGL calls are made, they are executed
// immedietly, the Execute() method is no-op, and large amounts of OpenGL calls
// will suffer from the performance degration described above.
//
// Additionally, batching may be turned on or off at runtime (even in-between
// OpenGL calls) using the SetBatching() method at runtime.
//
// Note that since batched OpenGL calls are only truly made upon invocation of
// the Execute() method, the memory for pointer arguments to OpenGL function
// calls must remain valid at least untill Execute() is called or else memory
// corruption is possible.
//
// Also note that OpenGL functions which return anything, such as the GetString
// function, cannot be queued for batching because they require that the value
// be immedietly returned. Also, if the function in question relies on previous
// OpenGL commands being executed first, you are responsible for calling the
// Execute() method.
//
// Since batching requires some extra effort on the programmer's side, and
// since debugging applications with batching turned on is more difficult, it
// is by default turned off.
type Context struct {
	c *C.gl_wrap_context
	batch []C.gl_wrap_batch_func
	batching bool
	loadedShaderVersion, loadedVersion bool
	major, minor, release int
	shaderMajor, shaderMinor, shaderRelease int
	vendorVersion, vendorShaderVersion string
	extensions map[string]bool
}

// New returns a new initialized Context with batching turned on.
func New() *Context {
	c := new(Context)
	c.c = new(C.gl_wrap_context)
	c.batching = false
	return c
}

// SetBatching turns on or off batching of this Context.
//
// If turning off batching using SetBatching(false), then an implicit call to
// Execute() which executes all the currently pending OpenGL function calls on
// this Context occurs.
//
// Default value is off (false).
func (c *Context) SetBatching(batching bool) {
	if !batching && len(c.batch) > 0 {
		c.Execute()
	}
	c.batching = batching
}

// Batching tells weather or not batching is on or off for this Context.
func (c *Context) Batching() bool {
	return c.batching
}

func (c *Context) push(f C.gl_wrap_batch_func) {
	c.batch = append(c.batch, f)
}

// Execute executes all the currently pending OpenGL function calls for this
// Context.
//
// This function executes pending calls regardless of weather or not batching
// is turned on.
func (c *Context) Execute() {
	if len(c.batch) > 0 {
		C.gl_wrap_batch_exec(
			c.c,
			(*C.gl_wrap_batch_func)(unsafe.Pointer(&c.batch[0])),
			C.int(len(c.batch)),
		)

		// Re-slice the batch
		c.batch = c.batch[:0]
	}
}

func (c *Context) parseVersionString(ver string) (major, minor, release int, vendor string) {
	if len(ver) == 0 {
		// Version string must not be empty
		return
	}

	// According to http://www.opengl.org/sdk/docs/man/xhtml/glGetString.xml
	//
	// the string returned may be 'major.minor' or 'major.minor.release'
	// and may be following by a space and any vendor specific information.

	// First locate a proper version string without vendor specific
	// information.
	var(
		versionString string
		err error
	)
	if strings.Contains(ver, " ") {
		// It must have vendor information
		split := strings.Split(ver, " ")
		if len(split) > 0 || len(split[0]) > 0 {
			// Everything looks good.
			versionString = split[0]
		} else {
			// Something must be wrong with their vendor string.
			return
		}

		// Store the vendor version information.
		vendor = ver[len(versionString):]
	} else {
		// No vendor information.
		versionString = ver
	}

	// We have a proper version string now without vendor information.
	dots := strings.Count(versionString, ".")
	if dots == 1 {
		// It's a 'major.minor' style string
		versions := strings.Split(versionString, ".")
		if len(versions) == 2 {
			major, err = strconv.Atoi(versions[0])
			if err != nil {
				return
			}

			minor, err = strconv.Atoi(versions[1])
			if err != nil {
				return
			}

		} else {
			return
		}

	} else if dots == 2 {
		// It's a 'major.minor.release' style string
		versions := strings.Split(versionString, ".")
		if len(versions) == 3 {
			major, err = strconv.Atoi(versions[0])
			if err != nil {
				return
			}

			minor, err = strconv.Atoi(versions[1])
			if err != nil {
				return
			}

			release, err = strconv.Atoi(versions[2])
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	return
}

func (c *Context) initVersion() {
	c.loadedVersion = true
	versionString := String(c.GetString(VERSION))
	c.major, c.minor, c.release, c.vendorVersion = c.parseVersionString(versionString)
}

func (c *Context) initShaderVersion() {
	c.loadedShaderVersion = true
	versionString := String(c.GetString(SHADING_LANGUAGE_VERSION))
	c.shaderMajor, c.shaderMinor, c.shaderRelease, c.vendorShaderVersion = c.parseVersionString(versionString)
}

// Version returns the major and minor OpenGL version supported by the
// implementation.
//
// Additionally provided is the vendor-specific release version (E.g. a driver
// version number).
func (c *Context) Version() (major, minor, release int) {
	if !c.loadedVersion {
		c.initVersion()
	}
	return c.major, c.minor, c.release
}

// AtLeastVersion tells if the major and minor OpenGL version reported by
// c.Version() is at least minMajor, minMinor.
//
// It is implemented as:
//
//  if major > minMajor {
//      return true
//  } else if major == minMajor && minor >= minMinor {
//      return true
//  }
//  return false
//
func (c *Context) AtLeastVersion(minMajor, minMinor int) bool {
	if !c.loadedVersion {
		c.initVersion()
	}
	major, minor := c.major, c.minor
	if major > minMajor {
		return true
	} else if major == minMajor && minor >= minMinor {
		return true
	}
	return false
}

// AtLeastRelease tells if the release OpenGL version reported by c.Version()
// is at least minRelease.
//
// It is implemented as:
//
//  if release == 0 {
//      return true
//  } else if release >= minRelease {
//      return true
//  }
//  return false
//
func (c *Context) AtLeastRelease(minRelease int) bool {
	if !c.loadedVersion {
		c.initVersion()
	}
	release := c.release
	if release == 0 {
		return true
	} else if release >= minRelease {
		return true
	}
	return false
}

// VendorVersion returns the vendor-specific OpenGL version information.
//
// This string depends on the vendor of the graphics card/driver.
func (c *Context) VendorVersion() string {
	if !c.loadedVersion {
		c.initVersion()
	}
	return c.vendorVersion
}

// ShaderVersion returns the major and minor version of the supported GLSL
// shading language per the OpenGL implementation.
//
// Additionally provided is the vendor-specific release version (E.g. a driver
// version number).
func (c *Context) ShaderVersion() (major, minor, release int) {
	if !c.loadedShaderVersion {
		c.initShaderVersion()
	}
	return c.shaderMajor, c.shaderMinor, c.shaderRelease
}

// AtLeastShaderVersion tells if the major and minor GLSL version reported by
// c.ShaderVersion() is at least minMajor, minMinor.
//
// It is implemented as:
//
//  if major > minMajor {
//      return true
//  } else if major == minMajor && minor >= minMinor {
//      return true
//  }
//  return false
//
func (c *Context) AtLeastShaderVersion(minMajor, minMinor int) bool {
	if !c.loadedShaderVersion {
		c.initShaderVersion()
	}
	major, minor := c.shaderMajor, c.shaderMinor
	if major > minMajor {
		return true
	} else if major == minMajor && minor >= minMinor {
		return true
	}
	return false
}

// AtLeastShaderRelease tells if the release GLSL version reported by
// c.ShaderVersion() is at least minRelease.
//
// It is implemented as:
//
//  if release == 0 {
//      return true
//  } else if release >= minRelease {
//      return true
//  }
//  return false
//
func (c *Context) AtLeastShaderRelease(minRelease int) bool {
	if !c.loadedShaderVersion {
		c.initShaderVersion()
	}
	release := c.shaderRelease
	if release == 0 {
		return true
	} else if release >= minRelease {
		return true
	}
	return false
}

// VendorShaderVersion returns the vendor-specific shading language version
// information.
//
// This string depends on the vendor of the graphics card/driver.
func (c *Context) VendorShaderVersion() string {
	if !c.loadedShaderVersion {
		c.initShaderVersion()
	}
	return c.vendorShaderVersion
}

func (c *Context) initExtensions() {
	// Initialize extensions map
	extString := String(c.GetString(EXTENSIONS))
	if len(extString) > 0 {
		split := strings.Split(extString, " ")
		c.extensions = make(map[string]bool, len(split))
		for _, ext := range split {
			if len(ext) > 0 {
				c.extensions[ext] = true
			}
		}
	} else {
		c.extensions = make(map[string]bool)
	}
}

// Extensions returns a slice of strings which represents all the extensions
// supported by the OpenGL implementation.
//
// Internally a map is used to store them, so a copy in the form of a slice is
// returned.
func (c *Context) Extensions() []string {
	if c.extensions == nil {
		c.initExtensions()
	}
	cpy := make([]string, len(c.extensions))
	i := 0
	for ext, _ := range c.extensions {
		cpy[i] = ext
		i++
	}
	return cpy
}

// Extension tells if the specified extension is supported by the OpenGL
// implementation.
func (c *Context) Extension(name string) (supported bool) {
	if c.extensions == nil {
		c.initExtensions()
	}
	_, supported = c.extensions[name]
	return
}

// Declare non-batchable functions

func (c *Context) GetError() int32 {
	return int32(C.gl_wrap_context_glGetError(c.c))
}

func (c *Context) GetString(name int32) *uint8 {
	return (*uint8)(unsafe.Pointer(C.gl_wrap_context_glGetString(c.c, C.GLenum(name))))
}

func (c *Context) IsBuffer(buffer uint32) uint8 {
	return uint8(C.gl_wrap_context_glIsBuffer(c.c, C.GLuint(buffer)))
}

func (c *Context) IsEnabled(cap int32) uint8 {
	return uint8(C.gl_wrap_context_glIsEnabled(c.c, C.GLenum(cap)))
}

func (c *Context) IsTexture(texture uint32) uint8 {
	return uint8(C.gl_wrap_context_glIsTexture(c.c, C.GLuint(texture)))
}


// Declare batchable functions

func (c *Context) AlphaFunc(pFunc int32, ref float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glAlphaFunc_args
		glWrapHandlerArgs._func = C.GLenum(pFunc);
		glWrapHandlerArgs.ref = C.GLfloat(ref);

		c.push(C.gl_wrap_batch_func{
			jump_index: 0,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glAlphaFunc(c.c, C.GLenum(pFunc), C.GLfloat(ref))
	}
}

func (c *Context) ClearColor(red float32, green float32, blue float32, alpha float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClearColor_args
		glWrapHandlerArgs.red = C.GLfloat(red);
		glWrapHandlerArgs.green = C.GLfloat(green);
		glWrapHandlerArgs.blue = C.GLfloat(blue);
		glWrapHandlerArgs.alpha = C.GLfloat(alpha);

		c.push(C.gl_wrap_batch_func{
			jump_index: 1,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClearColor(c.c, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}
}

func (c *Context) ClearDepthf(d float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClearDepthf_args
		glWrapHandlerArgs.d = C.GLfloat(d);

		c.push(C.gl_wrap_batch_func{
			jump_index: 2,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClearDepthf(c.c, C.GLfloat(d))
	}
}

func (c *Context) ClipPlanef(p int32, eqn *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClipPlanef_args
		glWrapHandlerArgs.p = C.GLenum(p);
		glWrapHandlerArgs.eqn = (*C.GLfloat)(unsafe.Pointer(eqn));

		c.push(C.gl_wrap_batch_func{
			jump_index: 3,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClipPlanef(c.c, C.GLenum(p), (*C.GLfloat)(unsafe.Pointer(eqn)))
	}
}

func (c *Context) Color4f(red float32, green float32, blue float32, alpha float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glColor4f_args
		glWrapHandlerArgs.red = C.GLfloat(red);
		glWrapHandlerArgs.green = C.GLfloat(green);
		glWrapHandlerArgs.blue = C.GLfloat(blue);
		glWrapHandlerArgs.alpha = C.GLfloat(alpha);

		c.push(C.gl_wrap_batch_func{
			jump_index: 4,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glColor4f(c.c, C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
	}
}

func (c *Context) DepthRangef(n float32, f float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDepthRangef_args
		glWrapHandlerArgs.n = C.GLfloat(n);
		glWrapHandlerArgs.f = C.GLfloat(f);

		c.push(C.gl_wrap_batch_func{
			jump_index: 5,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDepthRangef(c.c, C.GLfloat(n), C.GLfloat(f))
	}
}

func (c *Context) Fogf(pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFogf_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 6,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFogf(c.c, C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) Fogfv(pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFogfv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 7,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFogfv(c.c, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) Frustumf(l float32, r float32, b float32, t float32, n float32, f float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFrustumf_args
		glWrapHandlerArgs.l = C.GLfloat(l);
		glWrapHandlerArgs.r = C.GLfloat(r);
		glWrapHandlerArgs.b = C.GLfloat(b);
		glWrapHandlerArgs.t = C.GLfloat(t);
		glWrapHandlerArgs.n = C.GLfloat(n);
		glWrapHandlerArgs.f = C.GLfloat(f);

		c.push(C.gl_wrap_batch_func{
			jump_index: 8,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFrustumf(c.c, C.GLfloat(l), C.GLfloat(r), C.GLfloat(b), C.GLfloat(t), C.GLfloat(n), C.GLfloat(f))
	}
}

func (c *Context) GetClipPlanef(plane int32, equation *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetClipPlanef_args
		glWrapHandlerArgs.plane = C.GLenum(plane);
		glWrapHandlerArgs.equation = (*C.GLfloat)(unsafe.Pointer(equation));

		c.push(C.gl_wrap_batch_func{
			jump_index: 9,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetClipPlanef(c.c, C.GLenum(plane), (*C.GLfloat)(unsafe.Pointer(equation)))
	}
}

func (c *Context) GetFloatv(pname int32, data *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetFloatv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.data = (*C.GLfloat)(unsafe.Pointer(data));

		c.push(C.gl_wrap_batch_func{
			jump_index: 10,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetFloatv(c.c, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(data)))
	}
}

func (c *Context) GetLightfv(light int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetLightfv_args
		glWrapHandlerArgs.light = C.GLenum(light);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 11,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetLightfv(c.c, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetMaterialfv(face int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetMaterialfv_args
		glWrapHandlerArgs.face = C.GLenum(face);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 12,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetMaterialfv(c.c, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetTexEnvfv(target int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetTexEnvfv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 13,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetTexEnvfv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetTexParameterfv(target int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetTexParameterfv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 14,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetTexParameterfv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) LightModelf(pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightModelf_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 15,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightModelf(c.c, C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) LightModelfv(pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightModelfv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 16,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightModelfv(c.c, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) Lightf(light int32, pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightf_args
		glWrapHandlerArgs.light = C.GLenum(light);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 17,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightf(c.c, C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) Lightfv(light int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightfv_args
		glWrapHandlerArgs.light = C.GLenum(light);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 18,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightfv(c.c, C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) LineWidth(width float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLineWidth_args
		glWrapHandlerArgs.width = C.GLfloat(width);

		c.push(C.gl_wrap_batch_func{
			jump_index: 19,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLineWidth(c.c, C.GLfloat(width))
	}
}

func (c *Context) LoadMatrixf(m *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLoadMatrixf_args
		glWrapHandlerArgs.m = (*C.GLfloat)(unsafe.Pointer(m));

		c.push(C.gl_wrap_batch_func{
			jump_index: 20,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLoadMatrixf(c.c, (*C.GLfloat)(unsafe.Pointer(m)))
	}
}

func (c *Context) Materialf(face int32, pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMaterialf_args
		glWrapHandlerArgs.face = C.GLenum(face);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 21,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMaterialf(c.c, C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) Materialfv(face int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMaterialfv_args
		glWrapHandlerArgs.face = C.GLenum(face);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 22,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMaterialfv(c.c, C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) MultMatrixf(m *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMultMatrixf_args
		glWrapHandlerArgs.m = (*C.GLfloat)(unsafe.Pointer(m));

		c.push(C.gl_wrap_batch_func{
			jump_index: 23,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMultMatrixf(c.c, (*C.GLfloat)(unsafe.Pointer(m)))
	}
}

func (c *Context) MultiTexCoord4f(target int32, s float32, t float32, r float32, q float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMultiTexCoord4f_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.s = C.GLfloat(s);
		glWrapHandlerArgs.t = C.GLfloat(t);
		glWrapHandlerArgs.r = C.GLfloat(r);
		glWrapHandlerArgs.q = C.GLfloat(q);

		c.push(C.gl_wrap_batch_func{
			jump_index: 24,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMultiTexCoord4f(c.c, C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
	}
}

func (c *Context) Normal3f(nx float32, ny float32, nz float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glNormal3f_args
		glWrapHandlerArgs.nx = C.GLfloat(nx);
		glWrapHandlerArgs.ny = C.GLfloat(ny);
		glWrapHandlerArgs.nz = C.GLfloat(nz);

		c.push(C.gl_wrap_batch_func{
			jump_index: 25,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glNormal3f(c.c, C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
	}
}

func (c *Context) Orthof(l float32, r float32, b float32, t float32, n float32, f float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glOrthof_args
		glWrapHandlerArgs.l = C.GLfloat(l);
		glWrapHandlerArgs.r = C.GLfloat(r);
		glWrapHandlerArgs.b = C.GLfloat(b);
		glWrapHandlerArgs.t = C.GLfloat(t);
		glWrapHandlerArgs.n = C.GLfloat(n);
		glWrapHandlerArgs.f = C.GLfloat(f);

		c.push(C.gl_wrap_batch_func{
			jump_index: 26,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glOrthof(c.c, C.GLfloat(l), C.GLfloat(r), C.GLfloat(b), C.GLfloat(t), C.GLfloat(n), C.GLfloat(f))
	}
}

func (c *Context) PointParameterf(pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPointParameterf_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 27,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPointParameterf(c.c, C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) PointParameterfv(pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPointParameterfv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 28,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPointParameterfv(c.c, C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) PointSize(size float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPointSize_args
		glWrapHandlerArgs.size = C.GLfloat(size);

		c.push(C.gl_wrap_batch_func{
			jump_index: 29,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPointSize(c.c, C.GLfloat(size))
	}
}

func (c *Context) PolygonOffset(factor float32, units float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPolygonOffset_args
		glWrapHandlerArgs.factor = C.GLfloat(factor);
		glWrapHandlerArgs.units = C.GLfloat(units);

		c.push(C.gl_wrap_batch_func{
			jump_index: 30,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPolygonOffset(c.c, C.GLfloat(factor), C.GLfloat(units))
	}
}

func (c *Context) Rotatef(angle float32, x float32, y float32, z float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glRotatef_args
		glWrapHandlerArgs.angle = C.GLfloat(angle);
		glWrapHandlerArgs.x = C.GLfloat(x);
		glWrapHandlerArgs.y = C.GLfloat(y);
		glWrapHandlerArgs.z = C.GLfloat(z);

		c.push(C.gl_wrap_batch_func{
			jump_index: 31,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glRotatef(c.c, C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}
}

func (c *Context) Scalef(x float32, y float32, z float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glScalef_args
		glWrapHandlerArgs.x = C.GLfloat(x);
		glWrapHandlerArgs.y = C.GLfloat(y);
		glWrapHandlerArgs.z = C.GLfloat(z);

		c.push(C.gl_wrap_batch_func{
			jump_index: 32,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glScalef(c.c, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}
}

func (c *Context) TexEnvf(target int32, pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexEnvf_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 33,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexEnvf(c.c, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) TexEnvfv(target int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexEnvfv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 34,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexEnvfv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) TexParameterf(target int32, pname int32, param float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexParameterf_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfloat(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 35,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexParameterf(c.c, C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
	}
}

func (c *Context) TexParameterfv(target int32, pname int32, params *float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexParameterfv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfloat)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 36,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexParameterfv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(unsafe.Pointer(params)))
	}
}

func (c *Context) Translatef(x float32, y float32, z float32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTranslatef_args
		glWrapHandlerArgs.x = C.GLfloat(x);
		glWrapHandlerArgs.y = C.GLfloat(y);
		glWrapHandlerArgs.z = C.GLfloat(z);

		c.push(C.gl_wrap_batch_func{
			jump_index: 37,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTranslatef(c.c, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
	}
}

func (c *Context) ActiveTexture(texture int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glActiveTexture_args
		glWrapHandlerArgs.texture = C.GLenum(texture);

		c.push(C.gl_wrap_batch_func{
			jump_index: 38,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glActiveTexture(c.c, C.GLenum(texture))
	}
}

func (c *Context) AlphaFuncx(pFunc int32, ref int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glAlphaFuncx_args
		glWrapHandlerArgs._func = C.GLenum(pFunc);
		glWrapHandlerArgs.ref = C.GLfixed(ref);

		c.push(C.gl_wrap_batch_func{
			jump_index: 39,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glAlphaFuncx(c.c, C.GLenum(pFunc), C.GLfixed(ref))
	}
}

func (c *Context) BindBuffer(target int32, buffer uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glBindBuffer_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.buffer = C.GLuint(buffer);

		c.push(C.gl_wrap_batch_func{
			jump_index: 40,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glBindBuffer(c.c, C.GLenum(target), C.GLuint(buffer))
	}
}

func (c *Context) BindTexture(target int32, texture uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glBindTexture_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.texture = C.GLuint(texture);

		c.push(C.gl_wrap_batch_func{
			jump_index: 41,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glBindTexture(c.c, C.GLenum(target), C.GLuint(texture))
	}
}

func (c *Context) BlendFunc(sfactor int32, dfactor int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glBlendFunc_args
		glWrapHandlerArgs.sfactor = C.GLenum(sfactor);
		glWrapHandlerArgs.dfactor = C.GLenum(dfactor);

		c.push(C.gl_wrap_batch_func{
			jump_index: 42,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glBlendFunc(c.c, C.GLenum(sfactor), C.GLenum(dfactor))
	}
}

func (c *Context) BufferData(target int32, size uintptr, data unsafe.Pointer, usage int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glBufferData_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.size = C.GLsizeiptr(size);
		glWrapHandlerArgs.data = data;
		glWrapHandlerArgs.usage = C.GLenum(usage);

		c.push(C.gl_wrap_batch_func{
			jump_index: 43,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glBufferData(c.c, C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
	}
}

func (c *Context) BufferSubData(target int32, offset uintptr, size uintptr, data unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glBufferSubData_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.offset = C.GLintptr(offset);
		glWrapHandlerArgs.size = C.GLsizeiptr(size);
		glWrapHandlerArgs.data = data;

		c.push(C.gl_wrap_batch_func{
			jump_index: 44,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glBufferSubData(c.c, C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data)
	}
}

func (c *Context) Clear(mask uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClear_args
		glWrapHandlerArgs.mask = C.GLbitfield(mask);

		c.push(C.gl_wrap_batch_func{
			jump_index: 45,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClear(c.c, C.GLbitfield(mask))
	}
}

func (c *Context) ClearColorx(red int32, green int32, blue int32, alpha int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClearColorx_args
		glWrapHandlerArgs.red = C.GLfixed(red);
		glWrapHandlerArgs.green = C.GLfixed(green);
		glWrapHandlerArgs.blue = C.GLfixed(blue);
		glWrapHandlerArgs.alpha = C.GLfixed(alpha);

		c.push(C.gl_wrap_batch_func{
			jump_index: 46,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClearColorx(c.c, C.GLfixed(red), C.GLfixed(green), C.GLfixed(blue), C.GLfixed(alpha))
	}
}

func (c *Context) ClearDepthx(depth int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClearDepthx_args
		glWrapHandlerArgs.depth = C.GLfixed(depth);

		c.push(C.gl_wrap_batch_func{
			jump_index: 47,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClearDepthx(c.c, C.GLfixed(depth))
	}
}

func (c *Context) ClearStencil(s int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClearStencil_args
		glWrapHandlerArgs.s = C.GLint(s);

		c.push(C.gl_wrap_batch_func{
			jump_index: 48,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClearStencil(c.c, C.GLint(s))
	}
}

func (c *Context) ClientActiveTexture(texture int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClientActiveTexture_args
		glWrapHandlerArgs.texture = C.GLenum(texture);

		c.push(C.gl_wrap_batch_func{
			jump_index: 49,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClientActiveTexture(c.c, C.GLenum(texture))
	}
}

func (c *Context) ClipPlanex(plane int32, equation *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glClipPlanex_args
		glWrapHandlerArgs.plane = C.GLenum(plane);
		glWrapHandlerArgs.equation = (*C.GLfixed)(unsafe.Pointer(equation));

		c.push(C.gl_wrap_batch_func{
			jump_index: 50,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glClipPlanex(c.c, C.GLenum(plane), (*C.GLfixed)(unsafe.Pointer(equation)))
	}
}

func (c *Context) Color4ub(red uint8, green uint8, blue uint8, alpha uint8) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glColor4ub_args
		glWrapHandlerArgs.red = C.GLubyte(red);
		glWrapHandlerArgs.green = C.GLubyte(green);
		glWrapHandlerArgs.blue = C.GLubyte(blue);
		glWrapHandlerArgs.alpha = C.GLubyte(alpha);

		c.push(C.gl_wrap_batch_func{
			jump_index: 51,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glColor4ub(c.c, C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
	}
}

func (c *Context) Color4x(red int32, green int32, blue int32, alpha int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glColor4x_args
		glWrapHandlerArgs.red = C.GLfixed(red);
		glWrapHandlerArgs.green = C.GLfixed(green);
		glWrapHandlerArgs.blue = C.GLfixed(blue);
		glWrapHandlerArgs.alpha = C.GLfixed(alpha);

		c.push(C.gl_wrap_batch_func{
			jump_index: 52,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glColor4x(c.c, C.GLfixed(red), C.GLfixed(green), C.GLfixed(blue), C.GLfixed(alpha))
	}
}

func (c *Context) ColorMask(red uint8, green uint8, blue uint8, alpha uint8) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glColorMask_args
		glWrapHandlerArgs.red = C.GLboolean(red);
		glWrapHandlerArgs.green = C.GLboolean(green);
		glWrapHandlerArgs.blue = C.GLboolean(blue);
		glWrapHandlerArgs.alpha = C.GLboolean(alpha);

		c.push(C.gl_wrap_batch_func{
			jump_index: 53,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glColorMask(c.c, C.GLboolean(red), C.GLboolean(green), C.GLboolean(blue), C.GLboolean(alpha))
	}
}

func (c *Context) ColorPointer(size int32, pType int32, stride uint32, pointer unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glColorPointer_args
		glWrapHandlerArgs.size = C.GLint(size);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.stride = C.GLsizei(stride);
		glWrapHandlerArgs.pointer = pointer;

		c.push(C.gl_wrap_batch_func{
			jump_index: 54,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glColorPointer(c.c, C.GLint(size), C.GLenum(pType), C.GLsizei(stride), pointer)
	}
}

func (c *Context) CompressedTexImage2D(target int32, level int32, internalformat int32, width uint32, height uint32, border int32, imageSize uint32, data unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glCompressedTexImage2D_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.level = C.GLint(level);
		glWrapHandlerArgs.internalformat = C.GLenum(internalformat);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);
		glWrapHandlerArgs.border = C.GLint(border);
		glWrapHandlerArgs.imageSize = C.GLsizei(imageSize);
		glWrapHandlerArgs.data = data;

		c.push(C.gl_wrap_batch_func{
			jump_index: 55,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glCompressedTexImage2D(c.c, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), data)
	}
}

func (c *Context) CompressedTexSubImage2D(target int32, level int32, xoffset int32, yoffset int32, width uint32, height uint32, format int32, imageSize uint32, data unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glCompressedTexSubImage2D_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.level = C.GLint(level);
		glWrapHandlerArgs.xoffset = C.GLint(xoffset);
		glWrapHandlerArgs.yoffset = C.GLint(yoffset);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);
		glWrapHandlerArgs.format = C.GLenum(format);
		glWrapHandlerArgs.imageSize = C.GLsizei(imageSize);
		glWrapHandlerArgs.data = data;

		c.push(C.gl_wrap_batch_func{
			jump_index: 56,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glCompressedTexSubImage2D(c.c, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), data)
	}
}

func (c *Context) CopyTexImage2D(target int32, level int32, internalformat int32, x int32, y int32, width uint32, height uint32, border int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glCopyTexImage2D_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.level = C.GLint(level);
		glWrapHandlerArgs.internalformat = C.GLenum(internalformat);
		glWrapHandlerArgs.x = C.GLint(x);
		glWrapHandlerArgs.y = C.GLint(y);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);
		glWrapHandlerArgs.border = C.GLint(border);

		c.push(C.gl_wrap_batch_func{
			jump_index: 57,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glCopyTexImage2D(c.c, C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
	}
}

func (c *Context) CopyTexSubImage2D(target int32, level int32, xoffset int32, yoffset int32, x int32, y int32, width uint32, height uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glCopyTexSubImage2D_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.level = C.GLint(level);
		glWrapHandlerArgs.xoffset = C.GLint(xoffset);
		glWrapHandlerArgs.yoffset = C.GLint(yoffset);
		glWrapHandlerArgs.x = C.GLint(x);
		glWrapHandlerArgs.y = C.GLint(y);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);

		c.push(C.gl_wrap_batch_func{
			jump_index: 58,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glCopyTexSubImage2D(c.c, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}
}

func (c *Context) CullFace(mode int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glCullFace_args
		glWrapHandlerArgs.mode = C.GLenum(mode);

		c.push(C.gl_wrap_batch_func{
			jump_index: 59,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glCullFace(c.c, C.GLenum(mode))
	}
}

func (c *Context) DeleteBuffers(n uint32, buffers *uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDeleteBuffers_args
		glWrapHandlerArgs.n = C.GLsizei(n);
		glWrapHandlerArgs.buffers = (*C.GLuint)(unsafe.Pointer(buffers));

		c.push(C.gl_wrap_batch_func{
			jump_index: 60,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDeleteBuffers(c.c, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}
}

func (c *Context) DeleteTextures(n uint32, textures *uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDeleteTextures_args
		glWrapHandlerArgs.n = C.GLsizei(n);
		glWrapHandlerArgs.textures = (*C.GLuint)(unsafe.Pointer(textures));

		c.push(C.gl_wrap_batch_func{
			jump_index: 61,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDeleteTextures(c.c, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}
}

func (c *Context) DepthFunc(pFunc int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDepthFunc_args
		glWrapHandlerArgs._func = C.GLenum(pFunc);

		c.push(C.gl_wrap_batch_func{
			jump_index: 62,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDepthFunc(c.c, C.GLenum(pFunc))
	}
}

func (c *Context) DepthMask(flag uint8) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDepthMask_args
		glWrapHandlerArgs.flag = C.GLboolean(flag);

		c.push(C.gl_wrap_batch_func{
			jump_index: 63,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDepthMask(c.c, C.GLboolean(flag))
	}
}

func (c *Context) DepthRangex(n int32, f int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDepthRangex_args
		glWrapHandlerArgs.n = C.GLfixed(n);
		glWrapHandlerArgs.f = C.GLfixed(f);

		c.push(C.gl_wrap_batch_func{
			jump_index: 64,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDepthRangex(c.c, C.GLfixed(n), C.GLfixed(f))
	}
}

func (c *Context) Disable(cap int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDisable_args
		glWrapHandlerArgs.cap = C.GLenum(cap);

		c.push(C.gl_wrap_batch_func{
			jump_index: 65,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDisable(c.c, C.GLenum(cap))
	}
}

func (c *Context) DisableClientState(array int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDisableClientState_args
		glWrapHandlerArgs.array = C.GLenum(array);

		c.push(C.gl_wrap_batch_func{
			jump_index: 66,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDisableClientState(c.c, C.GLenum(array))
	}
}

func (c *Context) DrawArrays(mode int32, first int32, count uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDrawArrays_args
		glWrapHandlerArgs.mode = C.GLenum(mode);
		glWrapHandlerArgs.first = C.GLint(first);
		glWrapHandlerArgs.count = C.GLsizei(count);

		c.push(C.gl_wrap_batch_func{
			jump_index: 67,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDrawArrays(c.c, C.GLenum(mode), C.GLint(first), C.GLsizei(count))
	}
}

func (c *Context) DrawElements(mode int32, count uint32, pType int32, indices unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glDrawElements_args
		glWrapHandlerArgs.mode = C.GLenum(mode);
		glWrapHandlerArgs.count = C.GLsizei(count);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.indices = indices;

		c.push(C.gl_wrap_batch_func{
			jump_index: 68,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glDrawElements(c.c, C.GLenum(mode), C.GLsizei(count), C.GLenum(pType), indices)
	}
}

func (c *Context) Enable(cap int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glEnable_args
		glWrapHandlerArgs.cap = C.GLenum(cap);

		c.push(C.gl_wrap_batch_func{
			jump_index: 69,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glEnable(c.c, C.GLenum(cap))
	}
}

func (c *Context) EnableClientState(array int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glEnableClientState_args
		glWrapHandlerArgs.array = C.GLenum(array);

		c.push(C.gl_wrap_batch_func{
			jump_index: 70,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glEnableClientState(c.c, C.GLenum(array))
	}
}

func (c *Context) Finish() {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFinish_args

		c.push(C.gl_wrap_batch_func{
			jump_index: 71,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFinish(c.c)
	}
}

func (c *Context) Flush() {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFlush_args

		c.push(C.gl_wrap_batch_func{
			jump_index: 72,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFlush(c.c)
	}
}

func (c *Context) Fogx(pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFogx_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 73,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFogx(c.c, C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) Fogxv(pname int32, param *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFogxv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = (*C.GLfixed)(unsafe.Pointer(param));

		c.push(C.gl_wrap_batch_func{
			jump_index: 74,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFogxv(c.c, C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(param)))
	}
}

func (c *Context) FrontFace(mode int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFrontFace_args
		glWrapHandlerArgs.mode = C.GLenum(mode);

		c.push(C.gl_wrap_batch_func{
			jump_index: 75,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFrontFace(c.c, C.GLenum(mode))
	}
}

func (c *Context) Frustumx(l int32, r int32, b int32, t int32, n int32, f int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glFrustumx_args
		glWrapHandlerArgs.l = C.GLfixed(l);
		glWrapHandlerArgs.r = C.GLfixed(r);
		glWrapHandlerArgs.b = C.GLfixed(b);
		glWrapHandlerArgs.t = C.GLfixed(t);
		glWrapHandlerArgs.n = C.GLfixed(n);
		glWrapHandlerArgs.f = C.GLfixed(f);

		c.push(C.gl_wrap_batch_func{
			jump_index: 76,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glFrustumx(c.c, C.GLfixed(l), C.GLfixed(r), C.GLfixed(b), C.GLfixed(t), C.GLfixed(n), C.GLfixed(f))
	}
}

func (c *Context) GetBooleanv(pname int32, data *uint8) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetBooleanv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.data = (*C.GLboolean)(unsafe.Pointer(data));

		c.push(C.gl_wrap_batch_func{
			jump_index: 77,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetBooleanv(c.c, C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(data)))
	}
}

func (c *Context) GetBufferParameteriv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetBufferParameteriv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLint)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 78,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetBufferParameteriv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetClipPlanex(plane int32, equation *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetClipPlanex_args
		glWrapHandlerArgs.plane = C.GLenum(plane);
		glWrapHandlerArgs.equation = (*C.GLfixed)(unsafe.Pointer(equation));

		c.push(C.gl_wrap_batch_func{
			jump_index: 79,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetClipPlanex(c.c, C.GLenum(plane), (*C.GLfixed)(unsafe.Pointer(equation)))
	}
}

func (c *Context) GenBuffers(n uint32, buffers *uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGenBuffers_args
		glWrapHandlerArgs.n = C.GLsizei(n);
		glWrapHandlerArgs.buffers = (*C.GLuint)(unsafe.Pointer(buffers));

		c.push(C.gl_wrap_batch_func{
			jump_index: 80,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGenBuffers(c.c, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
	}
}

func (c *Context) GenTextures(n uint32, textures *uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGenTextures_args
		glWrapHandlerArgs.n = C.GLsizei(n);
		glWrapHandlerArgs.textures = (*C.GLuint)(unsafe.Pointer(textures));

		c.push(C.gl_wrap_batch_func{
			jump_index: 81,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGenTextures(c.c, C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
	}
}

func (c *Context) GetFixedv(pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetFixedv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 82,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetFixedv(c.c, C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetIntegerv(pname int32, data *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetIntegerv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.data = (*C.GLint)(unsafe.Pointer(data));

		c.push(C.gl_wrap_batch_func{
			jump_index: 83,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetIntegerv(c.c, C.GLenum(pname), (*C.GLint)(unsafe.Pointer(data)))
	}
}

func (c *Context) GetLightxv(light int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetLightxv_args
		glWrapHandlerArgs.light = C.GLenum(light);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 84,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetLightxv(c.c, C.GLenum(light), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetMaterialxv(face int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetMaterialxv_args
		glWrapHandlerArgs.face = C.GLenum(face);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 85,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetMaterialxv(c.c, C.GLenum(face), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetPointerv(pname int32, params *unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetPointerv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = params;

		c.push(C.gl_wrap_batch_func{
			jump_index: 86,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetPointerv(c.c, C.GLenum(pname), params)
	}
}

func (c *Context) GetTexEnviv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetTexEnviv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLint)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 87,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetTexEnviv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetTexEnvxv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetTexEnvxv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 88,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetTexEnvxv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetTexParameteriv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetTexParameteriv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLint)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 89,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetTexParameteriv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}
}

func (c *Context) GetTexParameterxv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glGetTexParameterxv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 90,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glGetTexParameterxv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) Hint(target int32, mode int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glHint_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.mode = C.GLenum(mode);

		c.push(C.gl_wrap_batch_func{
			jump_index: 91,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glHint(c.c, C.GLenum(target), C.GLenum(mode))
	}
}

func (c *Context) LightModelx(pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightModelx_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 92,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightModelx(c.c, C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) LightModelxv(pname int32, param *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightModelxv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = (*C.GLfixed)(unsafe.Pointer(param));

		c.push(C.gl_wrap_batch_func{
			jump_index: 93,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightModelxv(c.c, C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(param)))
	}
}

func (c *Context) Lightx(light int32, pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightx_args
		glWrapHandlerArgs.light = C.GLenum(light);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 94,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightx(c.c, C.GLenum(light), C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) Lightxv(light int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLightxv_args
		glWrapHandlerArgs.light = C.GLenum(light);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 95,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLightxv(c.c, C.GLenum(light), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) LineWidthx(width int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLineWidthx_args
		glWrapHandlerArgs.width = C.GLfixed(width);

		c.push(C.gl_wrap_batch_func{
			jump_index: 96,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLineWidthx(c.c, C.GLfixed(width))
	}
}

func (c *Context) LoadIdentity() {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLoadIdentity_args

		c.push(C.gl_wrap_batch_func{
			jump_index: 97,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLoadIdentity(c.c)
	}
}

func (c *Context) LoadMatrixx(m *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLoadMatrixx_args
		glWrapHandlerArgs.m = (*C.GLfixed)(unsafe.Pointer(m));

		c.push(C.gl_wrap_batch_func{
			jump_index: 98,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLoadMatrixx(c.c, (*C.GLfixed)(unsafe.Pointer(m)))
	}
}

func (c *Context) LogicOp(opcode int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glLogicOp_args
		glWrapHandlerArgs.opcode = C.GLenum(opcode);

		c.push(C.gl_wrap_batch_func{
			jump_index: 99,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glLogicOp(c.c, C.GLenum(opcode))
	}
}

func (c *Context) Materialx(face int32, pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMaterialx_args
		glWrapHandlerArgs.face = C.GLenum(face);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 100,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMaterialx(c.c, C.GLenum(face), C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) Materialxv(face int32, pname int32, param *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMaterialxv_args
		glWrapHandlerArgs.face = C.GLenum(face);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = (*C.GLfixed)(unsafe.Pointer(param));

		c.push(C.gl_wrap_batch_func{
			jump_index: 101,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMaterialxv(c.c, C.GLenum(face), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(param)))
	}
}

func (c *Context) MatrixMode(mode int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMatrixMode_args
		glWrapHandlerArgs.mode = C.GLenum(mode);

		c.push(C.gl_wrap_batch_func{
			jump_index: 102,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMatrixMode(c.c, C.GLenum(mode))
	}
}

func (c *Context) MultMatrixx(m *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMultMatrixx_args
		glWrapHandlerArgs.m = (*C.GLfixed)(unsafe.Pointer(m));

		c.push(C.gl_wrap_batch_func{
			jump_index: 103,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMultMatrixx(c.c, (*C.GLfixed)(unsafe.Pointer(m)))
	}
}

func (c *Context) MultiTexCoord4x(texture int32, s int32, t int32, r int32, q int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glMultiTexCoord4x_args
		glWrapHandlerArgs.texture = C.GLenum(texture);
		glWrapHandlerArgs.s = C.GLfixed(s);
		glWrapHandlerArgs.t = C.GLfixed(t);
		glWrapHandlerArgs.r = C.GLfixed(r);
		glWrapHandlerArgs.q = C.GLfixed(q);

		c.push(C.gl_wrap_batch_func{
			jump_index: 104,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glMultiTexCoord4x(c.c, C.GLenum(texture), C.GLfixed(s), C.GLfixed(t), C.GLfixed(r), C.GLfixed(q))
	}
}

func (c *Context) Normal3x(nx int32, ny int32, nz int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glNormal3x_args
		glWrapHandlerArgs.nx = C.GLfixed(nx);
		glWrapHandlerArgs.ny = C.GLfixed(ny);
		glWrapHandlerArgs.nz = C.GLfixed(nz);

		c.push(C.gl_wrap_batch_func{
			jump_index: 105,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glNormal3x(c.c, C.GLfixed(nx), C.GLfixed(ny), C.GLfixed(nz))
	}
}

func (c *Context) NormalPointer(pType int32, stride uint32, pointer unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glNormalPointer_args
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.stride = C.GLsizei(stride);
		glWrapHandlerArgs.pointer = pointer;

		c.push(C.gl_wrap_batch_func{
			jump_index: 106,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glNormalPointer(c.c, C.GLenum(pType), C.GLsizei(stride), pointer)
	}
}

func (c *Context) Orthox(l int32, r int32, b int32, t int32, n int32, f int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glOrthox_args
		glWrapHandlerArgs.l = C.GLfixed(l);
		glWrapHandlerArgs.r = C.GLfixed(r);
		glWrapHandlerArgs.b = C.GLfixed(b);
		glWrapHandlerArgs.t = C.GLfixed(t);
		glWrapHandlerArgs.n = C.GLfixed(n);
		glWrapHandlerArgs.f = C.GLfixed(f);

		c.push(C.gl_wrap_batch_func{
			jump_index: 107,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glOrthox(c.c, C.GLfixed(l), C.GLfixed(r), C.GLfixed(b), C.GLfixed(t), C.GLfixed(n), C.GLfixed(f))
	}
}

func (c *Context) PixelStorei(pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPixelStorei_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLint(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 108,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPixelStorei(c.c, C.GLenum(pname), C.GLint(param))
	}
}

func (c *Context) PointParameterx(pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPointParameterx_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 109,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPointParameterx(c.c, C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) PointParameterxv(pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPointParameterxv_args
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 110,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPointParameterxv(c.c, C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) PointSizex(size int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPointSizex_args
		glWrapHandlerArgs.size = C.GLfixed(size);

		c.push(C.gl_wrap_batch_func{
			jump_index: 111,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPointSizex(c.c, C.GLfixed(size))
	}
}

func (c *Context) PolygonOffsetx(factor int32, units int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPolygonOffsetx_args
		glWrapHandlerArgs.factor = C.GLfixed(factor);
		glWrapHandlerArgs.units = C.GLfixed(units);

		c.push(C.gl_wrap_batch_func{
			jump_index: 112,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPolygonOffsetx(c.c, C.GLfixed(factor), C.GLfixed(units))
	}
}

func (c *Context) PopMatrix() {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPopMatrix_args

		c.push(C.gl_wrap_batch_func{
			jump_index: 113,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPopMatrix(c.c)
	}
}

func (c *Context) PushMatrix() {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glPushMatrix_args

		c.push(C.gl_wrap_batch_func{
			jump_index: 114,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glPushMatrix(c.c)
	}
}

func (c *Context) ReadPixels(x int32, y int32, width uint32, height uint32, format int32, pType int32, pixels unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glReadPixels_args
		glWrapHandlerArgs.x = C.GLint(x);
		glWrapHandlerArgs.y = C.GLint(y);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);
		glWrapHandlerArgs.format = C.GLenum(format);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.pixels = pixels;

		c.push(C.gl_wrap_batch_func{
			jump_index: 115,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glReadPixels(c.c, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(pType), pixels)
	}
}

func (c *Context) Rotatex(angle int32, x int32, y int32, z int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glRotatex_args
		glWrapHandlerArgs.angle = C.GLfixed(angle);
		glWrapHandlerArgs.x = C.GLfixed(x);
		glWrapHandlerArgs.y = C.GLfixed(y);
		glWrapHandlerArgs.z = C.GLfixed(z);

		c.push(C.gl_wrap_batch_func{
			jump_index: 116,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glRotatex(c.c, C.GLfixed(angle), C.GLfixed(x), C.GLfixed(y), C.GLfixed(z))
	}
}

func (c *Context) SampleCoverage(value float32, invert uint8) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glSampleCoverage_args
		glWrapHandlerArgs.value = C.GLfloat(value);
		glWrapHandlerArgs.invert = C.GLboolean(invert);

		c.push(C.gl_wrap_batch_func{
			jump_index: 117,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glSampleCoverage(c.c, C.GLfloat(value), C.GLboolean(invert))
	}
}

func (c *Context) SampleCoveragex(value int32, invert uint8) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glSampleCoveragex_args
		glWrapHandlerArgs.value = C.GLclampx(value);
		glWrapHandlerArgs.invert = C.GLboolean(invert);

		c.push(C.gl_wrap_batch_func{
			jump_index: 118,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glSampleCoveragex(c.c, C.GLclampx(value), C.GLboolean(invert))
	}
}

func (c *Context) Scalex(x int32, y int32, z int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glScalex_args
		glWrapHandlerArgs.x = C.GLfixed(x);
		glWrapHandlerArgs.y = C.GLfixed(y);
		glWrapHandlerArgs.z = C.GLfixed(z);

		c.push(C.gl_wrap_batch_func{
			jump_index: 119,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glScalex(c.c, C.GLfixed(x), C.GLfixed(y), C.GLfixed(z))
	}
}

func (c *Context) Scissor(x int32, y int32, width uint32, height uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glScissor_args
		glWrapHandlerArgs.x = C.GLint(x);
		glWrapHandlerArgs.y = C.GLint(y);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);

		c.push(C.gl_wrap_batch_func{
			jump_index: 120,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glScissor(c.c, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}
}

func (c *Context) ShadeModel(mode int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glShadeModel_args
		glWrapHandlerArgs.mode = C.GLenum(mode);

		c.push(C.gl_wrap_batch_func{
			jump_index: 121,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glShadeModel(c.c, C.GLenum(mode))
	}
}

func (c *Context) StencilFunc(pFunc int32, ref int32, mask uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glStencilFunc_args
		glWrapHandlerArgs._func = C.GLenum(pFunc);
		glWrapHandlerArgs.ref = C.GLint(ref);
		glWrapHandlerArgs.mask = C.GLuint(mask);

		c.push(C.gl_wrap_batch_func{
			jump_index: 122,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glStencilFunc(c.c, C.GLenum(pFunc), C.GLint(ref), C.GLuint(mask))
	}
}

func (c *Context) StencilMask(mask uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glStencilMask_args
		glWrapHandlerArgs.mask = C.GLuint(mask);

		c.push(C.gl_wrap_batch_func{
			jump_index: 123,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glStencilMask(c.c, C.GLuint(mask))
	}
}

func (c *Context) StencilOp(fail int32, zfail int32, zpass int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glStencilOp_args
		glWrapHandlerArgs.fail = C.GLenum(fail);
		glWrapHandlerArgs.zfail = C.GLenum(zfail);
		glWrapHandlerArgs.zpass = C.GLenum(zpass);

		c.push(C.gl_wrap_batch_func{
			jump_index: 124,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glStencilOp(c.c, C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
	}
}

func (c *Context) TexCoordPointer(size int32, pType int32, stride uint32, pointer unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexCoordPointer_args
		glWrapHandlerArgs.size = C.GLint(size);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.stride = C.GLsizei(stride);
		glWrapHandlerArgs.pointer = pointer;

		c.push(C.gl_wrap_batch_func{
			jump_index: 125,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexCoordPointer(c.c, C.GLint(size), C.GLenum(pType), C.GLsizei(stride), pointer)
	}
}

func (c *Context) TexEnvi(target int32, pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexEnvi_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLint(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 126,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexEnvi(c.c, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}
}

func (c *Context) TexEnvx(target int32, pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexEnvx_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 127,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexEnvx(c.c, C.GLenum(target), C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) TexEnviv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexEnviv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLint)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 128,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexEnviv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}
}

func (c *Context) TexEnvxv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexEnvxv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 129,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexEnvxv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) TexImage2D(target int32, level int32, internalformat int32, width uint32, height uint32, border int32, format int32, pType int32, pixels unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexImage2D_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.level = C.GLint(level);
		glWrapHandlerArgs.internalformat = C.GLint(internalformat);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);
		glWrapHandlerArgs.border = C.GLint(border);
		glWrapHandlerArgs.format = C.GLenum(format);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.pixels = pixels;

		c.push(C.gl_wrap_batch_func{
			jump_index: 130,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexImage2D(c.c, C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(pType), pixels)
	}
}

func (c *Context) TexParameteri(target int32, pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexParameteri_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLint(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 131,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexParameteri(c.c, C.GLenum(target), C.GLenum(pname), C.GLint(param))
	}
}

func (c *Context) TexParameterx(target int32, pname int32, param int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexParameterx_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.param = C.GLfixed(param);

		c.push(C.gl_wrap_batch_func{
			jump_index: 132,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexParameterx(c.c, C.GLenum(target), C.GLenum(pname), C.GLfixed(param))
	}
}

func (c *Context) TexParameteriv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexParameteriv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLint)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 133,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexParameteriv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLint)(unsafe.Pointer(params)))
	}
}

func (c *Context) TexParameterxv(target int32, pname int32, params *int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexParameterxv_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.pname = C.GLenum(pname);
		glWrapHandlerArgs.params = (*C.GLfixed)(unsafe.Pointer(params));

		c.push(C.gl_wrap_batch_func{
			jump_index: 134,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexParameterxv(c.c, C.GLenum(target), C.GLenum(pname), (*C.GLfixed)(unsafe.Pointer(params)))
	}
}

func (c *Context) TexSubImage2D(target int32, level int32, xoffset int32, yoffset int32, width uint32, height uint32, format int32, pType int32, pixels unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTexSubImage2D_args
		glWrapHandlerArgs.target = C.GLenum(target);
		glWrapHandlerArgs.level = C.GLint(level);
		glWrapHandlerArgs.xoffset = C.GLint(xoffset);
		glWrapHandlerArgs.yoffset = C.GLint(yoffset);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);
		glWrapHandlerArgs.format = C.GLenum(format);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.pixels = pixels;

		c.push(C.gl_wrap_batch_func{
			jump_index: 135,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTexSubImage2D(c.c, C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(pType), pixels)
	}
}

func (c *Context) Translatex(x int32, y int32, z int32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glTranslatex_args
		glWrapHandlerArgs.x = C.GLfixed(x);
		glWrapHandlerArgs.y = C.GLfixed(y);
		glWrapHandlerArgs.z = C.GLfixed(z);

		c.push(C.gl_wrap_batch_func{
			jump_index: 136,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glTranslatex(c.c, C.GLfixed(x), C.GLfixed(y), C.GLfixed(z))
	}
}

func (c *Context) VertexPointer(size int32, pType int32, stride uint32, pointer unsafe.Pointer) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glVertexPointer_args
		glWrapHandlerArgs.size = C.GLint(size);
		glWrapHandlerArgs._type = C.GLenum(pType);
		glWrapHandlerArgs.stride = C.GLsizei(stride);
		glWrapHandlerArgs.pointer = pointer;

		c.push(C.gl_wrap_batch_func{
			jump_index: 137,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glVertexPointer(c.c, C.GLint(size), C.GLenum(pType), C.GLsizei(stride), pointer)
	}
}

func (c *Context) Viewport(x int32, y int32, width uint32, height uint32) {
	if c.batching {
		var glWrapHandlerArgs C.gl_wrap_handler_glViewport_args
		glWrapHandlerArgs.x = C.GLint(x);
		glWrapHandlerArgs.y = C.GLint(y);
		glWrapHandlerArgs.width = C.GLsizei(width);
		glWrapHandlerArgs.height = C.GLsizei(height);

		c.push(C.gl_wrap_batch_func{
			jump_index: 138,
			args: unsafe.Pointer(&glWrapHandlerArgs),
		})
	} else {
		C.gl_wrap_context_glViewport(c.c, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
	}
}


