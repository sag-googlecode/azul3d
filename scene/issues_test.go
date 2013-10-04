package scene

import (
	"testing"
)

func TestIssue4(t *testing.T) {
	root := New("parent")

	a := New("a")
	a.SetParent(root)

	b := New("b")
	root.AddChild(b)

	c := root.New("c")

	children := root.Children()
	if children[0] != a || children[1] != b || children[2] != c {
		t.Log("Children out of order.")
		t.Fail()
	}
}

func TestIssue3(t *testing.T) {
	root := New("parent")

	a := New("a")
	a.SetParent(root)

	a.Detatch()
	if a.Parent() != nil || len(a.Children()) != 0 {
		t.Log("Detatch() does not remove parent!")
		t.Fail()
	}
}
