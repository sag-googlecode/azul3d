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
	if w.glConfig != nil {
		return
	}
	w.glConfig = config

	if config == nil {
		panic("Invalid (nil) GLConfig; it must be an valid configuration!")
	}
	config.panicUnlessValid()

	dispatch(func() {
		if !win32.SetPixelFormat(w.dc, config.index, nil) {
			logger.Println("GLSetConfig failed; SetPixelFormat():", win32.GetLastErrorString())
		}
	})
}

func (w *W32Window) GLConfig() *GLConfig {
	return w.glConfig
}

func (w *W32Window) GLCreateContext(glVersionMinor, glVersionMajor, glVersionRevision uint) (GLContext, error) {
	if w.glConfig == nil {
		panic("Must call GLSetConfig() before GLCreateContext()!")
	}
	c := new(W32GLContext)
	c.valid = true

	var err error
	dispatch(func() {
		c.hglrc = win32.WglCreateContext(w.dc)
		if c.hglrc == nil {
			err = errors.New(fmt.Sprintf("Unable to create OpenGL Context; wglCreateContext(): %s", win32.GetLastErrorString()))
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
	wc := c.(*W32GLContext)
	wc.panicUnlessValid()
	wc.panicIfDestroyed()

	// Note: Avoid the temptation, never dispatch()!
	if !win32.WglMakeCurrent(w.dc, wc.hglrc) {
		logger.Println("Unable to make GL context current; wglMakeCurrent():", win32.GetLastErrorString())
	}
}

func (w *W32Window) GLSwapBuffers() {
	if w.Destroyed() {
		return
	}
	if !win32.SwapBuffers(w.dc) {
		logger.Println("Unable to swap GL buffers; SwapBuffers():", win32.GetLastErrorString())
	}
}
