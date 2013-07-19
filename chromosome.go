package govolve

// A Chromosome is a series of Alleles that can be evaluated against a given data array.
// It contains a Head, which consists of both Operators and Indexes, and a Tail, which
// is merely Indexes. This is so, if the Head contains only Operators of maximum Arity,
// the Tail will always have enough data leftover to make the Chromosome syntactically
// valid. Note this does not imply it is semantically valid. Chromosomes are evaluated
// using prefix notation: +-a+bcd = (+ (- a (+ b c)) d) = (a - (b + c)) + d.
type Chromosome struct {
	//Id      int
	Alleles []*Allele
	//Head    []Allele
}

// Finds the last index that is part of a chromosome's phenotype. This is necessary
// to evaluate the chromosome, since all of its alleles may not be used.
func (c *Chromosome) LastAlleleIndex() int {
	arity := 1
	for i, a := range c.Alleles {
		arity += a.Arity - 1
		if arity < 1 {
			return i
		}
	}
	return len(c.Alleles) - 1
}

func (c *Chromosome) Eval(data []float64) float64 {
	var stack Stack
	for i := c.LastAlleleIndex(); i >= 0; i-- {
		allele := c.Alleles[i]
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

// TODO: func RandomChromosome(headLength int, tailLength int, )
