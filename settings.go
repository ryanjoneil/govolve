package govolve

type Settings struct {
	UseArithmeticAlleles bool
}

func DefaultSettings() Settings {
	s := Settings{}
	s.UseArithmeticAlleles = true
	return s
}
