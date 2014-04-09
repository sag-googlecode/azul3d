// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"azul3d.org/v1/math"
	"image"
)

// Camera represents a camera object, it may be moved in 3D space using the
// objects transform and the viewing frustum controls how the camera views
// things. Since a camera is in itself also an object it may also have visible
// meshes attatched to it, etc.
type Camera struct {
	*Object

	// The projection matrix of the camera, which is responsible for projecting
	// world coordinates into device coordinates.
	Projection Mat4
}

// NewCamera returns a new *Camera with the default values.
func NewCamera() *Camera {
	return &Camera{
		NewObject(),
		ConvertMat4(math.Mat4Identity),
	}
}

// SetOrtho sets this camera's Projection matrix to an orthographic one.
//
// The view parameter is the viewing rectangle for the orthographic
// projection in window coordinates.
//
// The near and far parameters describe the minimum closest and maximum
// furthest clipping points of the view frustum.
//
// Clients who need advanced control over how the orthographic viewing frustum
// is set up may use this method's source as a reference (e.g. to change the
// center point, which this method sets at the bottom-left).
//
// Write access is required for this method to operate safely.
func (c *Camera) SetOrtho(view image.Rectangle, near, far float64) {
	w := float64(view.Dx())
	w = float64(int((w / 2.0)) * 2)
	h := float64(view.Dy())
	h = float64(int((h / 2.0)) * 2)
	m := math.Mat4Ortho(0, w, 0, h, near, far)
	c.Projection = ConvertMat4(m)
}

// SetPersp sets this camera's Projection matrix to an perspective one.
//
// The view parameter is the viewing rectangle for the orthographic
// projection in window coordinates.
//
// The fov parameter is the Y axis field of view (e.g. some games use 75) to
// use.
//
// The near and far parameters describe the minimum closest and maximum
// furthest clipping points of the view frustum.
//
// Clients who need advanced control over how the perspective viewing frustum
// is set up may use this method's source as a reference (e.g. to change the
// center point, which this method sets at the center).
//
// Write access is required for this method to operate safely.
func (c *Camera) SetPersp(view image.Rectangle, fov, near, far float64) {
	aspectRatio := float64(view.Dx()) / float64(view.Dy())
	m := math.Mat4Perspective(fov, aspectRatio, near, far)
	c.Projection = ConvertMat4(m)
}

var (
	// Get an matrix which will translate our matrix from ZUpRight to YUpRight
	zUpRightToYUpRight = math.CoordSysZUpRight.ConvertMat4(math.CoordSysYUpRight)
)

// Project returns a 2D point in normalized device space coordinates given a 3D
// point in the world.
//
// If ok=false is returned then the point is outside of the camera's view and
// the returned point may not be meaningful.
func (c *Camera) Project(p3 math.Vec3) (p2 math.Vec2, ok bool) {
	cameraInv, _ := c.Object.Transform.Mat4.Inverse()
	cameraInv = cameraInv.Mul(zUpRightToYUpRight)

	projection := c.Projection.Mat4()
	vp := cameraInv.Mul(projection)

	p4 := math.Vec4{p3.X, p3.Y, p3.Z, 1.0}
	p4 = p4.Transform(vp)
	if p4.W == 0 {
		p2 = math.Vec2Zero
		ok = false
		return
	}

	recipW := 1.0 / p4.W
	p2 = math.Vec2{p4.X * recipW, p4.Y * recipW}

	xValid := (p2.X >= -1) && (p2.X <= 1)
	yValid := (p2.Y >= -1) && (p2.Y <= 1)
	ok = (p4.W > 0) && xValid && yValid
	return
}

// Precision represents the precision in bits of the color, depth, and stencil
// buffers.
type Precision struct {
	// The precision in bits of each pixel in the color buffer, per color (e.g.
	// 8/8/8/8 would be 32bpp RGBA color, 8/8/8/0 would be 24bpp RGB color, and
	// so on).
	RedBits, GreenBits, BlueBits, AlphaBits uint8

	// The precision in bits of each pixel in the depth buffer (e.g. 8, 16, 24,
	// etc).
	DepthBits uint8

	// The precision in bits of each pixel in the stencil buffer (e.g. 8, 16,
	// 24, etc).
	StencilBits uint8

	// The number of samples available per pixel (e.g. the number of MSAA
	// samples).
	Samples int
}

// Canvas defines a canvas that can be drawn to (i.e. a window that the user
// will visibly see, or a texture that will store the results for later use).
//
// All methods must be safe to call from multiple goroutines.
type Canvas interface {
	Downloadable

	// SetMSAA should request that this canvas use multi-sample anti-aliasing
	// during rendering. By default MSAA is enabled.
	//
	// Even if MSAA is requested to be enabled, there is no guarantee that it
	// will actually be used. For instance if the graphics hardware or
	// rendering API does not support it.
	SetMSAA(enabled bool)

	// MSAA returns the last value passed into SetMSAA on this renderer.
	MSAA() bool

	// Precision should return the precision of the canvas's color, depth, and
	// stencil buffers.
	Precision() Precision

	// Bounds should return the bounding rectangle of this canvas, any and all
	// methods of this canvas that take rectangles as parameters will be
	// clamped to these bounds.
	// The bounds returned by this method may change at any given time (e.g.
	// when a user resizes the window).
	Bounds() image.Rectangle

	// Clear submits a clear operation to the renderer. It will clear the given
	// rectangle of the canvas's color buffer to the specified background
	// color.
	//
	// If the rectangle is empty the entire canvas is cleared.
	Clear(r image.Rectangle, bg Color)

	// ClearDepth submits a depth-clear operation to the renderer. It will
	// clear the given rectangle of the canvas's depth buffer to the specified
	// depth value (in the range of 0.0 to 1.0, where 1.0 is furthest away).
	//
	// If the rectangle is empty the entire canvas is cleared.
	ClearDepth(r image.Rectangle, depth float64)

	// ClearStencil submits a stencil-clear operation to the renderer. It will
	// clear the given rectangle of the canvas's stencil buffer to the
	// specified stencil value.
	//
	// If the rectangle is empty the entire canvas is cleared.
	ClearStencil(r image.Rectangle, stencil int)

	// Draw submits a draw operation to the renderer. It will draw the given
	// graphics object onto the specified rectangle of the canvas.
	//
	// The canvas will lock the object and camera object and they may stay
	// locked until some point in the future when the draw operation completes.
	//
	// If not nil, then the object is drawn according to how it is seen by the
	// given camera object (taking into account the camera object's
	// transformation and projection matrices).
	//
	// If the GPU supports occlusion queries (see GPUInfo.OcclusionQuery) and
	// o.OcclusionTest is set to true then at some point in the future (or when
	// QueryWait() is called) the native object will record the number of
	// samples that passed depth and stencil testing phases such that when
	// SampleCount() is called it will return the number of samples last drawn
	// by the object.
	//
	// The object will not be drawn if any of the following cases are true:
	//  o.Shader == nil
	//  len(o.Shader.Error) > 0
	//  len(o.Meshes) == 0
	//  !o.Meshes[N].Loaded && len(o.Meshes[N].Vertices) == 0
	//
	// If the rectangle is empty the entire canvas is cleared.
	Draw(r image.Rectangle, o *Object, c *Camera)

	// QueryWait blocks until all pending draw object's occlusion queries
	// completely finish. Most clients should avoid this call as it can easilly
	// cause graphics pipeline stalls if not handled with care.
	//
	// Instead of calling QueryWait immediately for conditional rendering, it is
	// common practice to instead make use of the previous frame's occlusion
	// query results as this allows the CPU and GPU to work in parralel instead
	// of being directly synchronized with one another.
	//
	// If the GPU does not support occlusion queries (see
	// GPUInfo.OcclusionQuery) then this function is no-op.
	QueryWait()

	// Render should finalize all pending clear and draw operations as if they
	// where all submitted over a single channel like so:
	//  pending := len(ops)
	//  for i := 0; i < pending; i++ {
	//      op := <-ops
	//      finalize(op)
	//  }
	// and once complete the final frame should be sent to the graphics
	// hardware for rasterization.
	//
	// Additionally, a call to Render() means an implicit call to QueryWait().
	Render()
}

// GPUInfo describes general information and limitations of the graphics
// hardware, such as the maximum texture size and other features which may or
// may not be supported by the graphics hardware.
type GPUInfo struct {
	// MaxTextureSize is the maximum size of either X or Y dimension of texture
	// images for use with the renderer, or -1 if not available.
	MaxTextureSize int

	// Whether or not the AlphaToCoverage alpha mode is supported (if false
	// then BinaryAlpha will automatically be used as a fallback).
	AlphaToCoverage bool

	// Whether or not occlusion queries are supported or not.
	OcclusionQuery bool

	// The number of bits reserved for the sample count when performing
	// occlusion queries, if the number goes above what this many bits could
	// store then it is generally (but not always) clamped to that value.
	OcclusionQueryBits int

	// The name of the graphics hardware, or an empty string if not available.
	// For example it may look something like:
	//  Mesa DRI Intel(R) Sandybridge Mobile
	Name string

	// The vendor name of the graphics hardware, or an empty string if not
	// available. For example:
	//  Intel Open Source Technology Center
	Vendor string

	// Major and minor versions of the OpenGL version in use, or -1 if not
	// available. For example:
	//  3, 0 (for OpenGL 3.0)
	GLMajor, GLMinor int

	// A read-only slice of OpenGL extension strings, empty if not available.
	GLExtensions []string

	// Major and minor versions of the OpenGL Shading Language version in use,
	// or -1 if not available. For example:
	//  1, 30 (for GLSL 1.30)
	GLSLMajor, GLSLMinor int

	// The maximum number of floating-point variables available for varying
	// variables inside GLSL programs, or -1 if not available. Generally at
	// least 32.
	GLSLMaxVaryingFloats int

	// The maximum number of shader inputs (i.e. floating-point values, where a
	// 4x4 matrix is 16 floating-point values) that can be used inside a GLSL
	// vertex shader, or -1 if not available. Generally at least 512.
	GLSLMaxVertexInputs int

	// The maximum number of shader inputs (i.e. floating-point values, where a
	// 4x4 matrix is 16 floating-point values) that can be used inside a GLSL
	// fragment shader, or -1 if not available. Generally at least 64.
	GLSLMaxFragmentInputs int
}

// Renderer is capable of loading meshes, textures, and shaders. A renderer can
// be drawn to as it implements the Canvas interface, and can also be used to
// All methods must be safe to call from multiple goroutines.
type Renderer interface {
	Canvas

	// GPUInfo should return information about the graphics hardware.
	GPUInfo() GPUInfo

	// LoadMesh should begin loading the specified mesh asynchronously.
	//
	// Additionally, the renderer will set m.Loaded to true, and then invoke
	// m.ClearData(), thus allowing the data slices to be garbage collected).
	//
	// The renderer will lock the mesh and it may stay locked until sometime in
	// the future when the load operation completes. The mesh will be sent over
	// the done channel once the load operation has completed if the channel is
	// not nil and sending would not block.
	LoadMesh(m *Mesh, done chan *Mesh)

	// LoadTexture should begin loading the specified texture asynchronously.
	//
	// Additionally, the renderer will set t.Loaded to true, and then invoke
	// t.ClearData(), thus allowing the source image to be garbage collected.
	//
	// The renderer will lock the texture and it may stay locked until sometime
	// in the future when the load operation completes. The texture will be
	// sent over the done channel once the load operation has completed if the
	// channel is not nil and sending would not block.
	LoadTexture(t *Texture, done chan *Texture)

	// LoadShader should begin loading the specified shader asynchronously.
	//
	// Additionally, if the shader was successfully loaded (no error log was
	// written) then the renderer will set s.Loaded to true, and then invoke
	// s.ClearData(), thus allowing the source data slices to be garbage
	// collected.
	//
	// The renderer will lock the shader and it may stay locked until sometime
	// in the future when the load operation completes. The shader will be sent
	// over the done channel once the load operation has completed if the
	// channel is not nil and sending would not block.
	LoadShader(s *Shader, done chan *Shader)

	// RenderToTexture should return a canvas that when drawn to the results
	// are stored inside of the given texture that may then be used in other
	// drawing operations.
	//
	// If the texture's bounding rectangle is empty then it will be set to the
	// bounds of this renderer's canvas. The texture's bounding rectangle is
	// what effectively determines the resolution at which the returned canvas
	// and texture render at.
	RenderToTexture(t *Texture) Canvas
}
