// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package gl

import (
	opengl "azul3d.org/v1/native/gl"
	"azul3d.org/v1/scene"
	"azul3d.org/v1/scene/shader"
	"runtime"
	"strings"
)

func (r *Renderer) updateShaderInput(gls *GLShader, name string, value interface{}) {
	bts := []byte(name)
	bts = append(bts, 0)
	location := r.gl.GetUniformLocation(gls.Program, &bts[0])
	r.gl.Execute()
	if location < 0 {
		return
	}

	switch v := value.(type) {
	case bool:
		var intBool int32
		if v {
			intBool = 1
		}
		r.gl.Uniform1iv(location, 1, &intBool)

	case float32:
		r.gl.Uniform1fv(location, 1, &v)
	case []float32:
		if len(v) > 0 {
			r.gl.Uniform1fv(location, uint32(len(v)), &v[0])
		}

	case shader.Vec2:
		r.gl.Uniform2fv(location, 1, &v[0])
	case []shader.Vec2:
		if len(v) > 0 {
			r.gl.Uniform2fv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Vec3:
		r.gl.Uniform3fv(location, 1, &v[0])
	case []shader.Vec3:
		if len(v) > 0 {
			r.gl.Uniform3fv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Vec4:
		r.gl.Uniform4fv(location, 1, &v[0])
	case []shader.Vec4:
		if len(v) > 0 {
			r.gl.Uniform4fv(location, uint32(len(v)), &v[0][0])
		}

	case int32:
		r.gl.Uniform1iv(location, 1, &v)
	case []int32:
		if len(v) > 0 {
			r.gl.Uniform1iv(location, uint32(len(v)), &v[0])
		}

	case shader.Vec2i:
		r.gl.Uniform2iv(location, 1, &v[0])
	case []shader.Vec2i:
		if len(v) > 0 {
			r.gl.Uniform2iv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Vec3i:
		r.gl.Uniform3iv(location, 1, &v[0])
	case []shader.Vec3i:
		if len(v) > 0 {
			r.gl.Uniform3iv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Vec4i:
		r.gl.Uniform4iv(location, 1, &v[0])
	case []shader.Vec4i:
		if len(v) > 0 {
			r.gl.Uniform4iv(location, uint32(len(v)), &v[0][0])
		}

	case uint32:
		r.gl.Uniform1uiv(location, 1, &v)
	case []uint32:
		if len(v) > 0 {
			r.gl.Uniform1uiv(location, uint32(len(v)), &v[0])
		}

	case shader.Vec2ui:
		r.gl.Uniform2uiv(location, 1, &v[0])
	case []shader.Vec2ui:
		if len(v) > 0 {
			r.gl.Uniform2uiv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Vec3ui:
		r.gl.Uniform3uiv(location, 1, &v[0])
	case []shader.Vec3ui:
		if len(v) > 0 {
			r.gl.Uniform3uiv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Vec4ui:
		r.gl.Uniform4uiv(location, 1, &v[0])
	case []shader.Vec4ui:
		if len(v) > 0 {
			r.gl.Uniform4uiv(location, uint32(len(v)), &v[0][0])
		}

	case shader.Mat2:
		r.gl.UniformMatrix2fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat2:
		if len(v) > 0 {
			r.gl.UniformMatrix2fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat3:
		r.gl.UniformMatrix3fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat3:
		if len(v) > 0 {
			r.gl.UniformMatrix3fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat4:
		r.gl.UniformMatrix4fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat4:
		if len(v) > 0 {
			r.gl.UniformMatrix4fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat2x3:
		r.gl.UniformMatrix2x3fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat2x3:
		if len(v) > 0 {
			r.gl.UniformMatrix2x3fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat3x2:
		r.gl.UniformMatrix3x2fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat3x2:
		if len(v) > 0 {
			r.gl.UniformMatrix3x2fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat2x4:
		r.gl.UniformMatrix2x4fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat2x4:
		r.gl.UniformMatrix2x4fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])

	case shader.Mat4x2:
		r.gl.UniformMatrix4x2fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat4x2:
		if len(v) > 0 {
			r.gl.UniformMatrix4x2fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat3x4:
		r.gl.UniformMatrix3x4fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat3x4:
		if len(v) > 0 {
			r.gl.UniformMatrix3x4fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	case shader.Mat4x3:
		r.gl.UniformMatrix4x3fv(location, 1, opengl.GLBool(false), &v[0][0])
	case []shader.Mat4x3:
		if len(v) > 0 {
			r.gl.UniformMatrix4x3fv(location, uint32(len(v)), opengl.GLBool(false), &v[0][0][0])
		}

	default:
		panic("Invalid shader input type!")
	}
}

func (r *Renderer) updateShaderInputs(n *scene.Node, s *shader.Shader, gls *GLShader) {
	r.gl.UseProgram(gls.Program)

	for name, value := range shader.Inputs(n) {
		r.updateShaderInput(gls, name, value)
	}
}

func (r *Renderer) loadShader(s *shader.Shader, now bool) {
	if s.Loaded() {
		return
	}

	doLoadShader := func(ctx *opengl.Context) {
		if now {
			// Release our display context
			r.dcMakeCurrent(false)

			// Later on we will use it
			defer r.dcMakeCurrent(true)
		} else {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
		}

		// Lock the loading context to thread
		r.lcAccess.Lock()
		defer r.lcAccess.Unlock()

		// Make the loading context active for this OS thread.
		r.lcMakeCurrent(true)

		// Later on release the loading context.
		defer r.lcMakeCurrent(false)

		// Create the shader
		gls := new(GLShader)

		shaderCompilerLog := func(s uint32) []byte {
			var ok int32
			ctx.GetShaderiv(s, opengl.COMPILE_STATUS, &ok)
			ctx.Execute()
			if ok == 0 {
				// Shader compiler error
				var logSize int32
				ctx.GetShaderiv(s, opengl.INFO_LOG_LENGTH, &logSize)
				ctx.Execute()

				log := make([]byte, logSize)
				ctx.GetShaderInfoLog(s, uint32(logSize), nil, &log[0])
				ctx.Execute()
				return log
			}
			return nil
		}

		appendError := func(err []byte) {
			s.SetError(append(s.Error(), err...))
		}

		vertSource := s.Source(shader.Vertex)
		if vertSource != nil {
			sVertSource := string(vertSource)
			sVertSource = strings.Replace(sVertSource, " ", "", -1)
			sVertSource = strings.Replace(sVertSource, "\t", "", -1)
			sVertSource = strings.Replace(sVertSource, "\n", "", -1)
			sVertSource = strings.Replace(sVertSource, "\r", "", -1)
			sVertSource = strings.Replace(sVertSource, "\r\n", "", -1)
			if len(sVertSource) == 0 {
				// Behavior is undefined (normally driver crashes).
				appendError([]byte(s.Name() + " | Vertex shader with no source code.\n"))

			} else {
				// Build vertex shader
				gls.Vertex = ctx.CreateShader(opengl.VERTEX_SHADER)
				lengths := int32(len(vertSource))
				sources := &vertSource[0]
				ctx.ShaderSource(gls.Vertex, 1, &sources, &lengths)
				ctx.CompileShader(gls.Vertex)
				ctx.Execute()

				log := shaderCompilerLog(gls.Vertex)
				if log != nil {
					// Sanity
					gls.Vertex = 0

					appendError([]byte(s.Name() + " | Vertex shader errors:\n"))
					appendError(log)
				}
			}
		}

		fragSource := s.Source(shader.Fragment)
		if fragSource != nil {
			sFragSource := string(fragSource)
			sFragSource = strings.Replace(sFragSource, " ", "", -1)
			sFragSource = strings.Replace(sFragSource, "\t", "", -1)
			sFragSource = strings.Replace(sFragSource, "\n", "", -1)
			sFragSource = strings.Replace(sFragSource, "\r", "", -1)
			sFragSource = strings.Replace(sFragSource, "\r\n", "", -1)
			if len(sFragSource) == 0 {
				// Behavior is undefined (normally driver crashes).
				appendError([]byte(s.Name() + " | Fragment shader with no source code.\n"))

			} else {
				// Build fragment shader
				gls.Fragment = ctx.CreateShader(opengl.FRAGMENT_SHADER)
				lengths := int32(len(fragSource))
				sources := &fragSource[0]
				ctx.ShaderSource(gls.Fragment, 1, &sources, &lengths)
				ctx.CompileShader(gls.Fragment)
				ctx.Execute()

				log := shaderCompilerLog(gls.Fragment)
				if log != nil {
					// Sanity
					gls.Fragment = 0

					appendError([]byte(s.Name() + " | Fragment shader errors:\n"))
					appendError(log)
				}
			}
		}

		if gls.Vertex != 0 && gls.Fragment != 0 {
			gls.Program = ctx.CreateProgram()

			ctx.AttachShader(gls.Program, gls.Vertex)
			ctx.AttachShader(gls.Program, gls.Fragment)
			ctx.LinkProgram(gls.Program)

			// Link shader program
			var ok int32
			ctx.GetProgramiv(gls.Program, opengl.LINK_STATUS, &ok)
			ctx.Execute()
			if ok == 0 {
				// Program linker error
				var logSize int32
				ctx.GetProgramiv(gls.Program, opengl.INFO_LOG_LENGTH, &logSize)
				ctx.Execute()

				log := make([]byte, logSize)
				ctx.GetProgramInfoLog(gls.Program, uint32(logSize), nil, &log[0])
				ctx.Execute()

				// Sanity
				gls.Program = 0

				appendError([]byte(s.Name() + " | Linker errors:\n"))
				appendError(log)
			}
		}

		// Store the native identity
		s.SetNativeIdentity(gls)

		// Wait for shader to be compiled
		ctx.Finish()
		ctx.Execute()

		// Notify of completion
		s.MarkLoaded()
	}
	if now {
		doLoadShader(r.gl)
	} else {
		go doLoadShader(r.lcgl)
	}
}

// As with other renderer calls, this is made inside an single OS thread only.
//
// But we may push it to an different thread if we wish to (we want to, of
// course).
func (r *Renderer) LoadShader(s *shader.Shader) {
	r.loadShader(s, false)
}
