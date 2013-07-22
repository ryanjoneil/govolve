package govolve

// Stack container for evaluating prefix notation.
type Stack []*Allele

func (s *Stack) Push(a *Allele) {
	*s = append(*s, a)
}

func (s *Stack) Pop() *Allele {
	var a *Allele
	l := len(*s)

	if l > 0 {
		a = (*s)[l-1]
		*s = (*s)[:l-1]
	} else {
		a = nil
	}

	return a
}

func (s *Stack) Empty() bool {
	return len(*s) < 1
}
