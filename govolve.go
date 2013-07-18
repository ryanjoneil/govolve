package main

import (
	"fmt"
	//"math"
)

// Stack container for evaluating prefix notation.
type Stack []*Allele

func (s *Stack) Push(e *Allele) {
	*s = append(*s, e)
}

func (s *Stack) Pop() *Allele {
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *Stack) Empty() bool {
	return len(*s) < 1
}

// An Allele represents a single operation that can occur against a stack.
type Allele struct {
	Sigil string
	Arity int
	Eval  func([]float64, *Stack) float64
}

func NewIndexAllele(sigil string, index int) *Allele {
	return &Allele{sigil, 0, func(d []float64, s *Stack) float64 { return d[index] }}
}

func NewValueAllele(value float64) *Allele {
	return &Allele{"", 0, func(d []float64, s *Stack) float64 { return value }}
}

// The default mathematical operators we start with.
var Operators = []*Allele{
	&Allele{"+", 2, func(data []float64, stack *Stack) float64 {
		return stack.Pop().Eval(data, stack) + stack.Pop().Eval(data, stack)
	}},
	&Allele{"-", 2, func(data []float64, stack *Stack) float64 {
		return stack.Pop().Eval(data, stack) - stack.Pop().Eval(data, stack)
	}},
	&Allele{"~", 1, func(data []float64, stack *Stack) float64 { return -stack.Pop().Eval(data, stack) }},
	&Allele{"*", 2, func(data []float64, stack *Stack) float64 {
		return stack.Pop().Eval(data, stack) * stack.Pop().Eval(data, stack)
	}},
	&Allele{"/", 2, func(data []float64, stack *Stack) float64 {
		return stack.Pop().Eval(data, stack) / stack.Pop().Eval(data, stack)
	}},
}

// These are operators that merely look up data in our input. We add one for every data index.
var Indexes []*Allele

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
	stack := Stack([]*Allele{})
	for i := c.LastAlleleIndex(); i >= 0; i-- {
		allele := c.Alleles[i]
		if allele.Arity > 0 {
			f := allele.Eval(data, &stack)
			a := NewValueAllele(f)
			stack.Push(a)
		} else {
			stack.Push(allele)
		}
	}
	return stack.Pop().Eval(data, &stack)
}

//func RandomChromosome(headLength int, tailLength int, )

func main() {
	//var stack Stack
	data := []float64{1, 2, 3, 5}
	sigils := []string{"a", "b", "c", "d"}

	// TODO: make this automatic based on data size
	// TODO: pre-size the slices with make to render this more efficient
	// TODO: use something other than Operators to store Operators + Indexes
	for i := 0; i < len(data); i++ {
		allele := NewIndexAllele(sigils[i], i)
		Operators = append(Operators, allele)
		Indexes = append(Indexes, allele)
	}

	//fmt.Println(Indexes[0])
	//fmt.Println(Operators[0].Eval(data, stack))
	//fmt.Println(Indexes[0].Eval(data, stack))

	c := &Chromosome{[]*Allele{ // +-a+bcd
		Operators[0], // +
		Operators[1], // -
		Operators[5], // a
		Operators[0], // +
		Operators[6], // b
		Operators[7], // c
		Operators[8], // d

		// Junk
		Operators[8], // d
		Operators[8], // d
		Operators[8], // d
	}}
	fmt.Println(c.Eval(data))

}
