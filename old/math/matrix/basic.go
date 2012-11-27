//===========================================================================//
//============ Azul3D, open source 3D game engine written in Go! ============//
//===========================================================================//
// File: basic.go
// Created by: Stephen Gutekanst, 11/19/12
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

// Transpose returns a matrix(b) where m's columns are b's rows, and m's rows are b's columns
func (m *Matrix) Transpose() *Matrix {
    co := m.ActualCopy()

    if co.columns == co.rows {
        // Square matrix
        for i := 0; i < m.rows; i++ {
            for j := 0; j < m.columns; j++ {
                co.elements[co.Index(j, i)] = m.elements[co.Index(i, j)]
            }
        }
    }
    return co
}

// Add returns a matrix representing m + b, or returns nil if the matrix's are of different sizes
func (m *Matrix) Add(b *Matrix) *Matrix {
    m.access.RLock()
    defer m.access.RUnlock()

    b.access.RLock()
    defer b.access.RUnlock()

    if m.rows != b.rows || m.columns != b.columns {
        return nil
    }

    n := m.Copy()
    for e := 0; e < len(m.elements); e++ {
        n.elements[e] = m.elements[e] + b.elements[e]
    }
    return n
}

// Subtract returns a matrix representing m - b, or returns nil if the matrix's are of different sizes
func (m *Matrix) Subtract(b *Matrix) *Matrix {
    m.access.RLock()
    defer m.access.RUnlock()

    b.access.RLock()
    defer b.access.RUnlock()

    if m.rows != b.rows || m.columns != b.columns {
        return nil
    }

    n := m.Copy()
    for e := 0; e < len(m.elements); e++ {
        n.elements[e] = m.elements[e] - b.elements[e]
    }
    return n
}

