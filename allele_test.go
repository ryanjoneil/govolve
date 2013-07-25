package govolve

import (
	"math"
	"testing"
)

func TestOrder(t *testing.T) {
	a := NewValueAllele(3)
	b := NewValueAllele(4)

	// Tests that evaluation is correct on order-sensitive operations.
	var result float64
	d := []float64{}
	s := &Stack{}
	
	// b - a
	s.Push(a)
	s.Push(b)
	result = AlleleSubtract.Op(d, s)
	if result != (b.Op(d,s) - a.Op(d,s)) {
		t.Errorf("b-a == %v, want %v", result, b.Op(d,s) - a.Op(d,s))
	}

	// a / b
	s.Push(b)
	s.Push(a)
	result = AlleleDivide.Op(d, s)
	if result != (a.Op(d,s) / b.Op(d,s)) {
		t.Errorf("a/b == %v, want %v", result, a.Op(d,s) / b.Op(d,s))
	}

	// a ^ b
	s.Push(b)
	s.Push(a)
	result = AllelePower.Op(d, s)
	if result != math.Pow(a.Op(d,s), b.Op(d,s)) {
		t.Errorf("a^b == %v, want %v", result, math.Pow(a.Op(d,s), b.Op(d,s)))
	}
}

/* TODO: not sure if we need this...
func TestZeroDivision(t *testing.T) {
	c := NewValueAllele(-1)
	z := NewValueAllele(0)

	// Tests that divion by 0 returns Inf.
	var result float64
	d := []float64{}
	s := &Stack{}
	
	// a / 0
	s.Push(z)
	s.Push(a)
	result = AlleleDivide.Op(d, s)
	if result != math.MaxFloat64 {
		t.Errorf("a/b == %v, want %v", result, math.MaxFloat64)
	}
}
*/