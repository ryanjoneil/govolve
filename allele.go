package govolve

// Operators have access to the data they are operating on as
// well as the Allele Stack for evaluation of chromosomes.
type Operator func([]float64, *Stack) float64

// An Allele represents a single operation that can occur against a stack.Stack.
type Allele struct {
	Sigil string
	Arity int
	Op    Operator
}

func NewIndexAllele(sigil string, index int) *Allele {
	return &Allele{sigil, 0, func(d []float64, s *Stack) float64 { return d[index] }}
}

func NewValueAllele(value float64) *Allele {
	return &Allele{"", 0, func(d []float64, s *Stack) float64 { return value }}
}

var ArithmeticAlleles = []*Allele{
	&Allele{"+", 2, func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) + s.Pop().Op(d, s) }},
	&Allele{"-", 2, func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) - s.Pop().Op(d, s) }},
	&Allele{"~", 1, func(d []float64, s *Stack) float64 { return -s.Pop().Op(d, s) }},
	&Allele{"*", 2, func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) * s.Pop().Op(d, s) }},
	&Allele{"/", 2, func(d []float64, s *Stack) float64 { return s.Pop().Op(d, s) / s.Pop().Op(d, s) }},
}

// TODO: 1/0, etc
// TODO: trig alleles, etc