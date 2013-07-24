package govolve

import "testing"

var (
	add = ArithmeticAlleles[0]
	subtract = ArithmeticAlleles[1]
	divide = ArithmeticAlleles[4]
	power = ArithmeticAlleles[5]
)

func TestEval(t *testing.T) {
	alleles := []*Allele{add, subtract, add, add, }
}