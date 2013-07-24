package govolve

import "strings"

var nextId = 0

// A Gene is a series of Alleles that can be evaluated against a given data array.
// It contains a Head, which consists of both Operators and Indexes, and a Tail, which
// is merely Indexes. This is so, if the Head contains only Operators of maximum Arity,
// the Tail will always have enough data leftover to make the Gene syntactically
// valid. Note this does not imply it is semantically valid. Genes are evaluated
// using prefix notation: +-a+bcd = (+ (- a (+ b c)) d) = (a - (b + c)) + d.
type Gene struct {
	Id      int
	Alleles []*Allele
	Coding  int
}

// Creates a Gene type from a slice of Alleles.
func NewGene(a []*Allele) *Gene {
	g := &Gene{Id: nextId, Alleles: a}
	g.Coding = g.CodingIndex()
	nextId += 1
	return g
}

// Converts a Gene to its string representation.
func (g *Gene) String() string {
	// TODO: probably a more idiomatic way to do this...
	s := make([]string, len(g.Alleles))
	for i, a := range g.Alleles {
		s[i] = a.Sigil
	}
	return strings.Join(s, "")
}

// Finds the last index that is part of a Gene's phenotype. This is necessary
// to evaluate the Gene, since all of its alleles may not be used. Otherwise
// called the coding region of the Gene.
func (g *Gene) CodingIndex() int {
	arity := 1
	for i, a := range g.Alleles {
		arity += a.Arity - 1
		if arity < 1 {
			return i
		}
	}
	return len(g.Alleles) - 1
}

// Evaluates a given Gene against some set of data.
func (g *Gene) Eval(data []float64) float64 {
	// Genes are prefix notation, so we simply start at the last allele
	// in the coding region of the Gene and evaluate backward.
	var stack Stack
	for i := g.Coding; i >= 0; i-- {
		allele := g.Alleles[i]
		if allele.Arity > 0 {
			f := allele.Op(data, &stack)
			a := NewValueAllele(f)
			stack.Push(a)
		} else {
			stack.Push(allele)
		}
	}
	return stack.Pop().Op(data, &stack)
}
