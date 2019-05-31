package google

type GoogleFunctions struct {
}

func (gf GoogleFunctions) Price(size float64) float64 {
	return size + 10
}

func (gf GoogleFunctions) Availability() bool {
	panic("implement me")
}
