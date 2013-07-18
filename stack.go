package govolve

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
