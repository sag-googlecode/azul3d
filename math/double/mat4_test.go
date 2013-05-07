package double

import (
	"testing"
)

func TestEq(t *testing.T) {
	if !eq(1.0, 1.0) {
		t.Fail()
	}
	if eq(1.0, 1.1) {
		t.Fail()
	}
	if !eq(1.0/2.0, 1.0/2.0) {
		t.Fail()
	}
}

func TestMat4ScalarMultiply(t *testing.T) {
	a := Matrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	a.ScalarMultiply(2)

	b := Matrix4(
		2, 4, 6, 8,
		10, 12, 14, 16,
		18, 20, 22, 24,
		26, 28, 30, 32,
	)

	if !a.Equals(b) {
		t.Fail()
	}
}

func TestMat4Identity(t *testing.T) {
	a := new(Mat4)
	a.Identity()
	if !a.IsIdentity() {
		t.Fail()
	}
}

func TestMat4Equality(t *testing.T) {
	var a, b *Mat4

	// Should be equal
	a = Matrix4(
		0.1, 0.2, 0.3, 0.4,
		0.5, 0.6, 0.7, 0.8,
		0.9, 1.0, 1.1, 1.2,
		1.3, 1.4, 1.5, 1.6,
	)

	b = Matrix4(
		0.1, 0.2, 0.3, 0.4,
		0.5, 0.6, 0.7, 0.8,
		0.9, 1.0, 1.1, 1.2,
		1.3, 1.4, 1.5, 1.6,
	)
	if !a.Equals(b) {
		t.Fail()
	}

	// Should not be equal
	a = Matrix4(
		0.1, 0.2, 0.3, 0.4,
		0.5, 0.6, 0.7, 0.8,
		0.9, 1.0, 1.1, 1.2,
		1.3, 1.4, 1.5, 1.6,
	)

	b = Matrix4(
		0.2, 0.2, 0.3, 0.4,
		0.5, 0.6, 0.7, 0.8,
		0.9, 1.0, 1.1, 1.2,
		1.3, 1.4, 1.5, 1.6,
	)
	if a.Equals(b) {
		t.Fail()
	}
}

func TestMat4Transpose(t *testing.T) {
	var a, b *Mat4

	a = Matrix4(
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
	)

	b = Matrix4(
		0, 4, 8, 12,
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
	)

	a.Transpose()

	if !a.Equals(b) {
		t.Fail()
	}

	a.Transpose()
	b.Transpose()
	if !a.Equals(b) {
		t.Fail()
	}
}

func TestMat4Invert(t *testing.T) {
	a := new(Mat4)
	a.Identity()
	if !a.Invert() {
		t.Fail()
	}
}
