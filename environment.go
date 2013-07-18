package govolve

import "fmt"

type Environment struct {
	Operators []*Allele // All operators, including data indexes
	Indexes   []*Allele // Just the data indexes (for building chromosome tails)
}

func NewEnvironment(settings Settings, dataLength int) *Environment {
    total := dataLength
    if settings.UseArithmeticAlleles { total += len(ArithmeticAlleles) }

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
        e.Operators[i + j] = allele
        e.Indexes[j] = allele
    }

    return e
}