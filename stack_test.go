package govolve

import "testing"

func TestEmpty(t *testing.T) {
	// Tests that Empty works as expected.
	var s Stack
	if !s.Empty() {
		t.Errorf("s.Empty() == %v, want %v", false, true)
	}

	s.Push(&Allele{})
	if s.Empty() {
		t.Errorf("s.Empty() == %v, want %v", true, false)
	}

	s.Pop()
	if !s.Empty() {
		t.Errorf("s.Empty() == %v, want %v", false, true)
	}
}

func TestPushPop(t *testing.T) {
	// Tests that Push/Pop return values in LIFO order.
	var s Stack
	a1 := &Allele{}
	a2 := &Allele{}
	a3 := &Allele{}

	s.Push(a1)
	s.Push(a2)
	s.Push(a3)

	if s.Pop() != a3 {
		t.Error("s.Pop() != a3")
	}

	s.Push(a1)

	if s.Pop() != a1 {
		t.Error("s.Pop() != a1")
	}
	if s.Pop() != a2 {
		t.Error("s.Pop() != a2")
	}
	if s.Pop() != a1 {
		t.Error("s.Pop() != a1")
	}
}

func TestOverPop(t *testing.T) {
	// Tests that Pop doesn't underflow the stack.
	var s Stack
	if s.Pop() != nil {
		t.Error("s.Pop() != nil")
	}
}
