// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !no_opengl

package chippy

import (
	"code.google.com/p/azul3d/chippy/wrappers/win32"
	"errors"
	"fmt"
)

type W32GLContext struct {
	valid     bool
	destroyed bool
	hglrc     win32.HGLRC
}

func (w *W32GLContext) panicUnlessValid() {
	if !w.valid {
		panic("Invalid GLContext; did you attempt to create it yourself?")
	}
}

func (w *W32GLContext) panicIfDestroyed() {
	if w.destroyed {
		panic("GL Context is already destroyed!")
	}
}

func (w *W32GLContext) Share(c GLContext) {
	w.panicUnlessValid()
	w.panicIfDestroyed()

	oc := c.(*W32GLContext)
	oc.panicUnlessValid()
	oc.panicIfDestroyed()
}

type backend_GLConfig struct {
	pf    *win32.PIXELFORMATDESCRIPTOR
	index win32.Int
}

func (w *W32Window) GLConfigs() (configs []*GLConfig) {
	dispatch(func() {
		// Just to get started
		max := win32.Int(2)

		for index := win32.Int(1); index-1 < max; index++ {
			var pf *win32.PIXELFORMATDESCRIPTOR
			max, pf = win32.DescribePixelFormat(w.dc, index)
			if max == 0 {
				logger.Println("Unable to get GLBufferFormats; DescribePixelFormat():", win32.GetLastErrorString())
				return
			}

			// We only want ones who have an flag of PFD_SUPPORT_OPENGL
			if (pf.DwFlags() & win32.PFD_SUPPORT_OPENGL) > 0 {

				// We only want ones whose pixel type is PFD_TYPE_RGBA
				if pf.IPixelType() == win32.PFD_TYPE_RGBA {
					config := new(GLConfig)
					config.valid = true
					config.pf = pf
					config.index = index

					config.RedBits = uint8(pf.CRedBits())
					config.GreenBits = uint8(pf.CGreenBits())
					config.BlueBits = uint8(pf.CBlueBits())
					config.AlphaBits = uint8(pf.CAlphaBits())

					config.AccumRedBits = uint8(pf.CAccumRedBits())
					config.AccumGreenBits = uint8(pf.CAccumGreenBits())
					config.AccumBlueBits = uint8(pf.CAccumBlueBits())
					config.AccumAlphaBits = uint8(pf.CAccumAlphaBits())

					config.AuxBuffers = uint8(pf.CAuxBuffers())

					if (pf.DwFlags()&win32.PFD_GENERIC_ACCELERATED) == 0 && (pf.DwFlags()&win32.PFD_GENERIC_FORMAT) > 0 {
						config.Accelerated = false
					} else {
						config.Accelerated = true
					}

					//logger.Println("GENERIC_FORMAT", pf.DwFlags() & win32.PFD_GENERIC_FORMAT)
					//logger.Println("GENERIC_ACCELERATED", pf.DwFlags() & win32.PFD_GENERIC_ACCELERATED)
					//config.Accelerated = ((pf.DwFlags() & win32.PFD_GENERIC_FORMAT)  & (pf.DwFlags() & win32.PFD_GENERIC_ACCELERATED)) > 0
					//logger.Println(config.Accelerated)

					config.DoubleBuffered = (pf.DwFlags() & win32.PFD_DOUBLEBUFFER) > 0
					config.StereoScopic = (pf.DwFlags() & win32.PFD_STEREO) > 0
					config.DepthBits = uint8(pf.CDepthBits())
					config.StencilBits = uint8(pf.CStencilBits())

					configs = append(configs, config)
				}
			}
		}
	})
	return
}

func (w *W32Window) GLSetConfig(config *GLConfig) {
	if config == nil {
		panic("Invalid (nil) GLConfig; it must be an valid configuration!")
	}
	config.panicUnlessValid()

	if w.glConfig != nil {
		return
	}
	w.glConfig = config

	dispatch(func() {
		if !win32.SetPixelFormat(w.dc, config.index, nil) {
			logger.Println("GLSetConfig failed; SetPixelFormat():", win32.GetLastErrorString())
		}
	})
}

func (w *W32Window) GLConfig() *GLConfig {
	return w.glConfig
}

func (w *W32Window) GLCreateContext(glVersionMajor, glVersionMinor uint) (GLContext, error) {
	if w.glConfig == nil {
		panic("Must call GLSetConfig() before GLCreateContext()!")
	}
	c := new(W32GLContext)
	c.valid = true

	var err error
	dispatch(func() {
		// First, make an fake context to use for context creation
		fakeContext := win32.WglCreateContext(w.dc)
		if fakeContext == nil {
			err = errors.New(fmt.Sprintf("Unable to create OpenGL context; wglCreateContext(): %s", win32.GetLastErrorString()))
			return
		}
		if !win32.WglMakeCurrent(w.dc, fakeContext) {
			err = errors.New(fmt.Sprintf("Unable to create OpenGL context; wglMakeCurrent(): %s", win32.GetLastErrorString()))
			return
		}

		attribs := []int{
			win32.WGL_CONTEXT_MAJOR_VERSION_ARB, 3,
			win32.WGL_CONTEXT_MINOR_VERSION_ARB, 0,
			//win32.WGL_CONTEXT_FLAGS_ARB, 0,
			//win32.WGL_CONTEXT_PROFILE_MASK_ARB, win32.WGL_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB,
			0,
		}

		extensions, ok := win32.WglGetExtensionsStringARB(w.dc)

		if extSupported(extensions, "WGL_ARB_create_context") {
			c.hglrc, ok = win32.WglCreateContextAttribsARB(w.dc, nil, attribs)
			if !ok {
				// The wglCreateContextAttribsARB entry point is missing
				//
				// Fall back to old context.
				logger.Println("WGL_ARB_create_context supported -- but wglCreateContextAttribsARB is missing!")
				c.hglrc = fakeContext

			} else if c.hglrc == nil {
				// Context couldn't be created for some reason (likely the version is not supported).
				//
				// Fall back to old context.

				logger.Println("wglCreateContextAttribsARB() failed:", win32.GetLastErrorString())
				c.hglrc = fakeContext
			} else {
				// It worked! We got our context!
				//
				// Clean up the fake context
				win32.WglMakeCurrent(nil, nil)
				win32.WglDeleteContext(fakeContext)

				// So we can get the version
				win32.WglMakeCurrent(w.dc, c.hglrc)
			}
		} else {
			// They have no WGL_ARB_create_context support.
			//
			// Fall back to old context.
			logger.Println("WGL_ARB_create_context is unavailable.")
			c.hglrc = fakeContext
		}

		win32.WglMakeCurrent(w.dc, c.hglrc)
		defer win32.WglMakeCurrent(nil, nil)

		ver := win32.GlGetString(win32.GL_VERSION)
		//logger.Printf("OpenGL: Driver version string is %q\n", ver)
		if !versionSupported(ver, int(glVersionMajor), int(glVersionMinor)) {
			err = errors.New(fmt.Sprintf("No OpenGL %d.%d support.", glVersionMajor, glVersionMinor))
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return c, err
}

func (w *W32Window) GLDestroyContext(c GLContext) {
	wc := c.(*W32GLContext)
	if !wc.destroyed {
		wc.destroyed = true
		dispatch(func() {
			if !win32.WglDeleteContext(wc.hglrc) {
				logger.Println("Unable to destroy GL context; wglDeleteContext():", win32.GetLastErrorString())
			}
		})
	}
}

func (w *W32Window) GLMakeCurrent(c GLContext) {
	var hglrc win32.HGLRC

	if c != nil {
		wc := c.(*W32GLContext)
		wc.panicUnlessValid()
		wc.panicIfDestroyed()
		hglrc = wc.hglrc
	}

	// Note: Avoid the temptation, never dispatch()!
	if !win32.WglMakeCurrent(w.dc, hglrc) {
		logger.Println("Unable to make GL context current; wglMakeCurrent():", win32.GetLastErrorString())
	}
}

func (w *W32Window) GLSwapBuffers() {
	if w.Destroyed() {
		return
	}
	if w.glConfig.DoubleBuffered == false {
		return
	}
	if !win32.SwapBuffers(w.dc) {
		logger.Println("Unable to swap GL buffers; SwapBuffers():", win32.GetLastErrorString())
	}
}
