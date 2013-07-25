package govolve

import "math"

// TODO: Other types of alleles

// Operators have access to the data they are operating on as
// well as the Allele Stack for evaluation of chromosomes.
type Operator func([]float64, *Stack) float64

// An Allele represents a single operation that can occur against a Stack.
type Allele struct {
	Sigil string
	Arity int
	Op    Operator
}

var (
	OperatorAdd = func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) + s.Pop().Op(d, s) }
	OperatorSubtract = func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) - s.Pop().Op(d, s) }
	OperatorNegate = func(d []float64, s *Stack) float64 { return -s.Pop().Op(d, s) }
	OperatorMultiply = func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) * s.Pop().Op(d, s) }
	OperatorDivide = func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) / s.Pop().Op(d, s) }
	OperatorPower = func(d []float64, s *Stack) float64 { return math.Pow(s.Pop().Op(d, s), s.Pop().Op(d, s)) }

	AlleleAdd = &Allele{"+", 2, OperatorAdd}
	AlleleSubtract = &Allele{"-", 2, OperatorSubtract}
	AlleleNegate = &Allele{"~", 1, OperatorNegate}
	AlleleMultiply = &Allele{"*", 2, OperatorMultiply}
	AlleleDivide = &Allele{"/", 2, OperatorDivide}
	AllelePower = &Allele{"^", 2, OperatorPower}

	ArithmeticAlleles = []*Allele{
		AlleleAdd,
		AlleleSubtract,
		AlleleNegate,
		AlleleMultiply,
		AlleleDivide,
		AllelePower,
	}
)

// Generates an Allele that simply returns the index of a data slice.
func NewIndexAllele(sigil string, index int) *Allele {
	return &Allele{sigil, 0, func(d []float64, s *Stack) float64 { return d[index] }}
}

// Generates an allele that simply returns a given value.
func NewValueAllele(value float64) *Allele {
	return &Allele{"", 0, func(d []float64, s *Stack) float64 { return value }}
}
