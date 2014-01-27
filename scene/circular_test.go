// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build circular

package scene

import (
	"testing"
)

func TestCircularParent(t *testing.T) {
	// Consider this graph:
	//
	// - A
	//   - A
	//
	// The last one should panic, as A is an parent to A!
	defer func() {
		r := recover()
		if r != CircularErr {
			t.Log(r)
			t.Fail()
		}
	}()

	A := New("A")
	A.SetParent(A)
	t.Log("Circular parent references should cause an panic.")
	t.Fail()
}

func TestCircularChild(t *testing.T) {
	// Consider this graph:
	//
	// - A
	//   - B
	//     - A
	//
	// The last one should panic, as A is an parent to B!
	defer func() {
		r := recover()
		if r != CircularErr {
			t.Log(r)
			t.Fail()
		}
	}()

	A := New("A")
	B := New("B")
	A.AddChild(B)
	B.AddChild(A)
	t.Log("Circular child references should cause an panic.")
	t.Fail()
}
