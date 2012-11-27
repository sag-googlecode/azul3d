//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: matrix.go
// Created by: Stephen Gutekanst, 11/18/12
//===========================================================================//
//===========================================================================//
// Copyright (c) 2012, Lightpoke
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//     * Redistributions of source code must retain the above copyright
//       notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, this list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//     * Neither the name of the <organization> nor the
//       names of its contributors may be used to endorse or promote products
//       derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package matrix

import "bytes"
import "sync"
import "fmt"

type Matrix struct {
    elements []float64
    rows, columns int
    access sync.RWMutex
}

func New(elements []float64, columns int) *Matrix {
    m := new(Matrix)
    m.elements = elements
    m.columns = columns
    m.rows = len(elements) / columns
    return m
}

// Copy returns a new copy of this matrix
func (m *Matrix) Copy() *Matrix {
    return New(m.elements, m.columns)
}

// ActualCopy returns a new copy of this matrix,
// unlike Copy, ActualCopy will copy the backend
// array of this matrix. This is an "actual" copy
func (m *Matrix) ActualCopy() *Matrix {
    n := []float64{}

    for e := 0; e < len(m.elements); e++ {
        n = append(n, m.elements[e])
    }
    return New(n, m.columns)
}

// Rows returns the number of rows
func (m *Matrix) Rows() int {
    return m.rows
}

// Columns returns the number of columns
func (m *Matrix) Columns() int {
    return m.columns
}

// Index returns the index of the element at row, column
func (m *Matrix) Index(row, column int) int {
    return row * m.columns + column
}

// Element returns the element at row, column
func (m *Matrix) Element(row, column int) float64 {
    return m.elements[m.Index(row, column)]
}

// Elements returns all of the elements
func (m *Matrix) Elements() []float64 {
    return m.elements
}

func (m *Matrix) String() string {
	b := bytes.Buffer{}
	fmt.Fprintf(&b, "Matrix(%dx%d,\n", m.rows, m.columns)

	for r := 0; r < m.rows; r++ {
		b.WriteString("    (")
		for c := 0; c < m.columns; c++ {
			if c > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(&b, "%f", m.elements[m.Index(r, c)])
		}
		b.WriteString(")\n")
	}
	b.WriteString(")")
	return b.String()
}
