// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build linux

// Do not use - Wrappers around very specific x11 APIs.
package x11

/*
#include <stdlib.h>
#include <GL/glx.h>
#include <X11/Xlib-xcb.h>

#cgo LDFLAGS: -lX11 -lGL

int chippy_handle_glx_secret_event(Display* d, xcb_generic_event_t* ev);

GLXContext chippy_glXCreateContextAttribsARB(void* p, Display* dpy, GLXFBConfig config, GLXContext share, Bool direct, const int* attribs);
void chippy_glXSwapIntervalEXT(void* p, Display* dpy, GLXDrawable d, int interval);
int chippy_glXSwapIntervalMESA(void* p, int interval);
int chippy_glXSwapIntervalSGI(void* p, int interval);
*/
import "C"

import (
	"reflect"
	"unsafe"
)

const (
	GLX_DOUBLEBUFFER     = C.GLX_DOUBLEBUFFER
	GLX_STEREO           = C.GLX_STEREO
	GLX_AUX_BUFFERS      = C.GLX_AUX_BUFFERS
	GLX_RED_SIZE         = C.GLX_RED_SIZE
	GLX_GREEN_SIZE       = C.GLX_GREEN_SIZE
	GLX_BLUE_SIZE        = C.GLX_BLUE_SIZE
	GLX_ALPHA_SIZE       = C.GLX_ALPHA_SIZE
	GLX_DEPTH_SIZE       = C.GLX_DEPTH_SIZE
	GLX_STENCIL_SIZE     = C.GLX_STENCIL_SIZE
	GLX_ACCUM_RED_SIZE   = C.GLX_ACCUM_RED_SIZE
	GLX_ACCUM_GREEN_SIZE = C.GLX_ACCUM_GREEN_SIZE
	GLX_ACCUM_BLUE_SIZE  = C.GLX_ACCUM_BLUE_SIZE
	GLX_ACCUM_ALPHA_SIZE = C.GLX_ACCUM_ALPHA_SIZE

	GLX_SAMPLE_BUFFERS = C.GLX_SAMPLE_BUFFERS
	GLX_SAMPLES        = C.GLX_SAMPLES

	GLX_TRANSPARENT_TYPE  = C.GLX_TRANSPARENT_TYPE
	GLX_NONE              = C.GLX_NONE
	GLX_TRANSPARENT_RGB   = C.GLX_TRANSPARENT_RGB
	GLX_TRANSPARENT_INDEX = C.GLX_TRANSPARENT_INDEX

	GLX_TRANSPARENT_INDEX_VALUE = C.GLX_TRANSPARENT_INDEX_VALUE
	GLX_TRANSPARENT_RED_VALUE   = C.GLX_TRANSPARENT_RED_VALUE
	GLX_TRANSPARENT_GREEN_VALUE = C.GLX_TRANSPARENT_GREEN_VALUE
	GLX_TRANSPARENT_BLUE_VALUE  = C.GLX_TRANSPARENT_BLUE_VALUE
	GLX_TRANSPARENT_ALPHA_VALUE = C.GLX_TRANSPARENT_ALPHA_VALUE

	GLX_RENDER_TYPE      = C.GLX_RENDER_TYPE
	GLX_RGBA_TYPE        = C.GLX_RGBA_TYPE
	GLX_COLOR_INDEX_TYPE = C.GLX_COLOR_INDEX_TYPE

	GLX_X_VISUAL_TYPE = C.GLX_X_VISUAL_TYPE
	GLX_TRUE_COLOR    = C.GLX_TRUE_COLOR

	GLX_CONFIG_CAVEAT         = C.GLX_CONFIG_CAVEAT
	GLX_SLOW_CONFIG           = C.GLX_SLOW_CONFIG
	GLX_NON_CONFORMANT_CONFIG = C.GLX_NON_CONFORMANT_CONFIG

	GLX_CONTEXT_MAJOR_VERSION_ARB             = C.GLX_CONTEXT_MAJOR_VERSION_ARB
	GLX_CONTEXT_MINOR_VERSION_ARB             = C.GLX_CONTEXT_MINOR_VERSION_ARB
	GLX_CONTEXT_FORWARD_COMPATIBLE_BIT_ARB    = C.GLX_CONTEXT_FORWARD_COMPATIBLE_BIT_ARB
	GLX_CONTEXT_DEBUG_BIT_ARB                 = C.GLX_CONTEXT_DEBUG_BIT_ARB
	GLX_CONTEXT_FLAGS_ARB                     = C.GLX_CONTEXT_FLAGS_ARB
	GLX_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB = C.GLX_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB
	GLX_CONTEXT_CORE_PROFILE_BIT_ARB          = C.GLX_CONTEXT_CORE_PROFILE_BIT_ARB
	GLX_CONTEXT_PROFILE_MASK_ARB              = C.GLX_CONTEXT_PROFILE_MASK_ARB
)

type (
	GLXContext  C.GLXContext
	GLXDrawable C.GLXDrawable
	GLXFBConfig C.GLXFBConfig
	GLXWindow   C.GLXWindow
)

func (d *Display) HandleGLXSecretEvent(ev *GenericEvent) bool {
	return C.chippy_handle_glx_secret_event(
		(*[0]byte)(unsafe.Pointer(d)),
		(*C.xcb_generic_event_t)(unsafe.Pointer(ev.EGenericEvent)),
	) != 0
}

func (d *Display) GLXCreateNewContext(config GLXFBConfig, renderType int, shareList GLXContext, direct bool) GLXContext {
	cDirect := C.Bool(0)
	if direct {
		cDirect = 1
	}
	return GLXContext(C.glXCreateNewContext(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXFBConfig(config),
		C.int(renderType),
		C.GLXContext(shareList),
		cDirect,
	))
}

func (d *Display) GLXDestroyContext(ctx GLXContext) {
	C.glXDestroyContext(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXContext(ctx),
	)
}

func (d *Display) GLXMakeContextCurrent(draw, read GLXDrawable, ctx GLXContext) int {
	return int(C.glXMakeContextCurrent(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXDrawable(draw),
		C.GLXDrawable(read),
		C.GLXContext(ctx),
	))
}

func (d *Display) GLXSwapBuffers(drawable GLXDrawable) {
	C.glXSwapBuffers(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXDrawable(drawable),
	)
}

type Int C.int

func (d *Display) GLXQueryVersion(maj, min *Int) bool {
	return C.glXQueryVersion(
		(*[0]byte)(unsafe.Pointer(d)),
		(*C.int)(unsafe.Pointer(maj)),
		(*C.int)(unsafe.Pointer(min)),
	) != 0
}

func (d *Display) GLXQueryExtensionsString(screen int) string {
	data := C.glXQueryExtensionsString(
		(*[0]byte)(unsafe.Pointer(d)),
		C.int(screen),
	)
	return C.GoString(data)
}

func (d *Display) GLXGetFBConfigAttrib(config GLXFBConfig, attrib int) (value Int, ret int) {
	ret = int(C.glXGetFBConfigAttrib(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXFBConfig(config),
		C.int(attrib),
		(*C.int)(unsafe.Pointer(&value)),
	))
	return
}

func (d *Display) GLXGetFBConfigs(screen int) (configs []GLXFBConfig) {
	var nConfigs C.int
	cConfigs := C.glXGetFBConfigs(
		(*[0]byte)(unsafe.Pointer(d)),
		C.int(screen),
		&nConfigs,
	)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&configs))
	sliceHeader.Len = int(nConfigs)
	sliceHeader.Cap = int(nConfigs)
	sliceHeader.Data = uintptr(unsafe.Pointer(cConfigs))
	return
}

func (d *Display) GLXCreateWindow(config GLXFBConfig, win Window) GLXWindow {
	return GLXWindow(C.glXCreateWindow(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXFBConfig(config),
		C.Window(win),
		nil,
	))
}

func (d *Display) GLXDestroyWindow(win GLXWindow) {
	C.glXDestroyWindow(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXWindow(win),
	)
}

func (d *Display) GLXGetVisualFromFBConfig(config GLXFBConfig) *XVisualInfo {
	return (*XVisualInfo)(unsafe.Pointer(C.glXGetVisualFromFBConfig(
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXFBConfig(config),
	)))
}

func GLXGetCurrentContext() GLXContext {
	return GLXContext(C.glXGetCurrentContext())
}

func (d *Display) GLXGetProcAddressARB(p string) unsafe.Pointer {
	cstr := C.CString(p)
	defer C.free(unsafe.Pointer(cstr))
	return unsafe.Pointer(C.glXGetProcAddressARB(
		(*C.GLubyte)(unsafe.Pointer(cstr)),
	))
}

var glXCreateContextAttribsARBPtr unsafe.Pointer

func (d *Display) GLXCreateContextAttribsARB(config GLXFBConfig, share GLXContext, direct bool, attribs *Int) GLXContext {
	if glXCreateContextAttribsARBPtr == nil {
		glXCreateContextAttribsARBPtr = d.GLXGetProcAddressARB("glXCreateContextAttribsARB")
	}

	cDirect := C.Bool(0)
	if direct {
		cDirect = 1
	}
	return GLXContext(C.chippy_glXCreateContextAttribsARB(
		glXCreateContextAttribsARBPtr,
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXFBConfig(config),
		C.GLXContext(share),
		cDirect,
		(*C.int)(unsafe.Pointer(attribs)),
	))
}

var glXSwapIntervalEXTPtr unsafe.Pointer

func (d *Display) GLXSwapIntervalEXT(drawable GLXDrawable, interval int) {
	if glXSwapIntervalEXTPtr == nil {
		glXSwapIntervalEXTPtr = d.GLXGetProcAddressARB("glXSwapIntervalEXT")
	}

	C.chippy_glXSwapIntervalEXT(
		glXSwapIntervalEXTPtr,
		(*[0]byte)(unsafe.Pointer(d)),
		C.GLXDrawable(drawable),
		C.int(interval),
	)
}

var glXSwapIntervalMESAPtr unsafe.Pointer

func (d *Display) GLXSwapIntervalMESA(interval int) int {
	if glXSwapIntervalMESAPtr == nil {
		glXSwapIntervalMESAPtr = d.GLXGetProcAddressARB("glXSwapIntervalMESA")
	}

	// Should be OK to use glXSwapIntervalEXT but with the MESA pointer because
	// they have the same typedef
	return int(C.chippy_glXSwapIntervalMESA(
		glXSwapIntervalMESAPtr,
		C.int(interval),
	))
}

var glXSwapIntervalSGIPtr unsafe.Pointer

func (d *Display) GLXSwapIntervalSGI(interval int) int {
	if glXSwapIntervalSGIPtr == nil {
		glXSwapIntervalSGIPtr = d.GLXGetProcAddressARB("glXSwapIntervalSGI")
	}

	// Should be OK to use glXSwapIntervalEXT but with the SGI pointer because
	// they have the same typedef
	return int(C.chippy_glXSwapIntervalSGI(
		glXSwapIntervalSGIPtr,
		C.int(interval),
	))
}

func XFree(ptr unsafe.Pointer) {
	C.XFree(ptr)
}
