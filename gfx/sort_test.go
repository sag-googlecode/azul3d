// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import (
	"azul3d.org/v1/math"
	"math/rand"
	"sort"
	"testing"
)

func TestSortByDist(t *testing.T) {
	a := NewObject()
	a.Transform.Pos = math.Vec3{10, 10, 10}

	b := NewObject()
	b.Transform.Pos = math.Vec3{-10, 2, 2}

	c := NewObject()
	c.Transform.Pos = math.Vec3{0, 6, 5}

	byDist := ByDist{
		Objects: []*Object{a, b, c, a, b, c, b, c, a},
		Target:  math.Vec3{0, 0, 0},
	}
	sort.Sort(byDist)

	for i := 0; i < 3; i++ {
		p := byDist.Objects[i].Transform.Pos
		if p != a.Pos {
			t.Fail()
		}
	}

	for i := 3; i < 6; i++ {
		p := byDist.Objects[i].Transform.Pos
		if p != b.Pos {
			t.Fail()
		}
	}

	for i := 6; i < 9; i++ {
		p := byDist.Objects[i].Transform.Pos
		if p != c.Pos {
			t.Fail()
		}
	}
}

func sortByDist(amount int, b *testing.B) {
	b.StopTimer()
	byDist := ByDist{
		Objects: make([]*Object, amount),
		Target: math.Vec3{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		},
	}
	for i := 0; i < amount; i++ {
		byDist.Objects[i] = NewObject()
	}

	for _, o := range byDist.Objects {
		o.Transform.Pos = math.Vec3{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		}
	}
	b.StartTimer()

	sort.Sort(byDist)
}

func BenchmarkSortByDist250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByDist(250, b)
	}
}

func BenchmarkSortByDist500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByDist(500, b)
	}
}

func BenchmarkSortByDist1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByDist(1000, b)
	}
}

func BenchmarkSortByDist5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByDist(5000, b)
	}
}

func TestSortByState(t *testing.T) {
	a := NewObject()
	a.State.Texturing = true
	a.State.Dithering = true
	a.State.DepthTest = true
	a.State.DepthWrite = true
	a.State.WriteRed = false
	a.State.WriteGreen = true
	a.State.WriteBlue = false
	a.State.WriteAlpha = true

	b := NewObject()
	b.State.Texturing = true
	b.State.Dithering = false
	b.State.DepthTest = true
	b.State.DepthWrite = false
	a.State.WriteRed = false
	a.State.WriteGreen = false
	a.State.WriteBlue = true
	a.State.WriteAlpha = true

	var l = []*Object{a, b, a, b, a, a, b, b, a, a, a, a, b, b, b, b}
	sort.Sort(ByState(l))

	for i := 0; i < 8; i++ {
		s := l[i].State
		if !s.Texturing || !s.Dithering || !s.DepthTest || !s.DepthWrite {
			t.Fail()
		}
	}

	for i := 8; i < 16; i++ {
		s := l[i].State
		if !s.Texturing || s.Dithering || !s.DepthTest || s.DepthWrite {
			t.Fail()
		}
	}
}

func sortByState(amount int, b *testing.B) {
	b.StopTimer()
	objs := make([]*Object, amount)
	for i := 0; i < amount; i++ {
		objs[i] = NewObject()
	}

	randBool := func() bool {
		return (rand.Int() % 2) == 0
	}

	for _, o := range objs {
		o.State = State{
			Texturing:   randBool(),
			WriteRed:    randBool(),
			WriteGreen:  randBool(),
			WriteBlue:   randBool(),
			WriteAlpha:  randBool(),
			Dithering:   randBool(),
			DepthTest:   randBool(),
			DepthWrite:  randBool(),
			StencilTest: randBool(),
		}
	}
	b.StartTimer()

	sort.Sort(ByState(objs))
}

func BenchmarkSortByState250(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByState(250, b)
	}
}

func BenchmarkSortByState500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByState(500, b)
	}
}

func BenchmarkSortByState1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByState(1000, b)
	}
}

func BenchmarkSortByState5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortByState(5000, b)
	}
}
