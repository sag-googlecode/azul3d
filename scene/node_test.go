// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package scene

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestParentRelationship(t *testing.T) {
	// Scene graph visualized:
	// - A
	//   - B
	//     - C
	//   - D
	//    - E
	//      - F
	A := New("A")
	B := A.New("B")
	C := B.New("C")
	D := A.New("D")
	E := D.New("E")
	F := E.New("F")

	fParents := F.Parents()
	if fParents[0] != E {
		t.Log("Parent relationship is not valid")
		t.Fail()

	} else if fParents[1] != D {
		t.Log("Parent relationship is not valid")
		t.Fail()

	} else if fParents[2] != A {
		t.Log("Parent relationship is not valid")
		t.Fail()
	}

	if F.Parent() != E {
		t.Log("Parent relationship is not valid")
		t.Fail()

	} else if E.Parent() != D {
		t.Log("Parent relationship is not valid")
		t.Fail()

	} else if D.Parent() != A {
		t.Log("Parent relationship is not valid")
		t.Fail()

	} else if C.Parent() != B {
		t.Log("Parent relationship is not valid")
		t.Fail()

	} else if B.Parent() != A {
		t.Log("Parent relationship is not valid")
		t.Fail()
	}
}

func TestChildRelationship(t *testing.T) {
	// Scene graph visualized:
	// - A
	//   - B
	//     - C
	//   - D
	//    - E
	//      - F
	A := New("A")
	B := A.New("B")
	C := B.New("C")
	D := A.New("D")
	E := D.New("E")
	F := E.New("F")

	if A.Children()[0] != B {
		t.Log("Child relationship is not valid")
		t.Fail()

	} else if B.Children()[0] != C {
		t.Log("Child relationship is not valid")
		t.Fail()

	} else if D.Children()[0] != E {
		t.Log("Child relationship is not valid")
		t.Fail()

	} else if E.Children()[0] != F {
		t.Log("Child relationship is not valid")
		t.Fail()
	}
}

func TestTop(t *testing.T) {
	// Scene graph visualized:
	// - A
	//   - B
	//     - C
	A := New("A")
	B := A.New("B")
	C := B.New("C")

	if C.Top() != A || B.Top() != A || A.Top() != A {
		t.Fail()
	}
}

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

func TestDoubleTraversal(t *testing.T) {
	// Scene graph visualized:
	// - A
	//   - B
	//     - C
	//   - D
	//    - E
	//      - F
	A := New("A")
	B := A.New("B")
	B.New("C")
	D := A.New("D")
	E := D.New("E")
	E.New("F")

	A.Traverse(func(i int, current *Node) bool {
		current.Traverse(func(i int, c *Node) bool {
			return true
		})
		return true
	})
}

func TestFindCommonParent(t *testing.T) {
	// Scene graph visualized:
	// - A
	//   - B
	//     - C
	//   - D
	//    - E
	//      - F
	A := New("A")
	B := A.New("B")
	C := B.New("C")
	D := A.New("D")
	E := D.New("E")
	F := E.New("F")

	if B.FindCommonParent(D) != A {
		t.Log("Failed to find common parent")
		t.Fail()
	}

	if D.FindCommonParent(B) != A {
		t.Log("Failed to find common parent")
		t.Fail()
	}

	if F.FindCommonParent(C) != A {
		t.Log("Failed to find common parent")
		t.Fail()
	}
}

func BenchmarkFindCommonParent(b *testing.B) {
	A := New("A")

	last := A
	for i := 0; i < 100; i++ {
		last = last.New(strconv.Itoa(i))
	}

	last = A
	for i := 0; i < 100; i++ {
		last = last.New(strconv.Itoa(i))
	}

	var possible []*Node
	A.Traverse(func(i int, current *Node) bool {
		possible = append(possible, current)
		return true
	})

	rand.Seed(time.Now().Unix())
	min := 0
	max := len(possible)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		choiceOne := rand.Intn(max-min) + min
		choiceTwo := rand.Intn(max-min) + min
		one := possible[choiceOne]
		two := possible[choiceTwo]

		one.FindCommonParent(two)
	}
}

func BenchmarkParents(b *testing.B) {
	A := New("A")

	last := A
	for i := 0; i < 100; i++ {
		last = last.New(strconv.Itoa(i))
	}

	last = A
	for i := 0; i < 100; i++ {
		last = last.New(strconv.Itoa(i))
	}

	var possible []*Node
	A.Traverse(func(i int, current *Node) bool {
		possible = append(possible, current)
		return true
	})

	rand.Seed(time.Now().Unix())
	min := 0
	max := len(possible)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		choiceOne := rand.Intn(max-min) + min
		one := possible[choiceOne]
		one.Parents()
	}
}
