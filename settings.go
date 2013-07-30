package govolve

type Settings struct {
	UseArithmeticAlleles   bool
	UseTrigonometryAlleles bool
	UseLogarithmicAlleles  bool
	UseLogicalAlleles      bool
}

func DefaultSettings() Settings {
	s := Settings{}
	s.UseArithmeticAlleles = true
	return s
}
