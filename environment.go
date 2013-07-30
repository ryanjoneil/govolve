/*
Package govolve implements linear genetic programming for use against
floating point data sets. A simple example may look like:

	settings := govolve.DefaultSettings()
	env := govolve.Environment(settings, 5)
	// TODO FINISH
*/
package govolve

import "fmt"

type Environment struct {
	Operators []*Allele // All operators, including data indexes
	Indexes   []*Allele // Just the data indexes (for building chromosome tails)
}

// NOTE: tail length = max arity + (max arity - 1) * (head length - 1)
// TODO: change to take headLength instead of dataLength
func NewDefaultEnvironment(dataLength int) *Environment {
	return NewEnvironment(DefaultSettings(), dataLength)
}

func NewEnvironment(settings Settings, dataLength int) *Environment {
	total := dataLength
	if settings.UseArithmeticAlleles {
		total += len(ArithmeticAlleles)
	}

	e := &Environment{make([]*Allele, total), make([]*Allele, dataLength)}

	// Add all the various mathematical operators.
	i := 0
	if settings.UseArithmeticAlleles {
		copy(e.Operators[i:], ArithmeticAlleles)
		i += len(ArithmeticAlleles)
	}

	// Create data index alleles.
	sigils := "abcdefghijklmnopqrstuvwxzy"
	for j := 0; j < dataLength; j++ {
		// Use a letter for the index allele sigil, assuming there are enough letters.
		var sigil string
		if j < len(sigils) {
			sigil = string(sigils[j])
		} else {
			sigil = fmt.Sprintf("[%d]", j)
		}

		allele := NewIndexAllele(sigil, i)
		e.Operators[i+j] = allele
		e.Indexes[j] = allele
	}

	return e
}
