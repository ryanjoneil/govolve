package govolve

import "testing"

func TestGeneEval(t *testing.T) {
	a := NewIndexAllele("a", 0)
	b := NewIndexAllele("b", 1)
	c := NewIndexAllele("c", 2)
	d := NewIndexAllele("d", 3)

	// +-c+d~a-bd == c - (d+(-a)) + (b-d)
	alleles := []*Allele{AlleleAdd, AlleleSubtract, c, AlleleAdd, d, AlleleNegate, a, AlleleSubtract, b, d}
	gene := NewGene(alleles)

	// Test that string representation of the gene is correct.
	s1 := gene.String()
	s2 := "+-c+d~a-bd"
	if s1 != s2 {
		t.Errorf("gene.String() == %v, want %v", s1, s2)
	}

	// Test evaluation of the gene.
	data := []float64{1, 2, 3, 4}
	v1 := gene.Eval(data)
	v2 := 3.0 - (4.0 + (-1.0)) + (2.0 - 4.0)
	if v1 != v2 {
		t.Errorf("gene.Eval(%v) == %v, want %v", data, v1, v2)
	}
}
