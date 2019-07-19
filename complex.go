package gomath

import "math"

type Complex128 struct {
	real      float64
	imaginary float64
}

func CreateComplex128(real float64, imaginary float64) Complex128 {
	c := Complex128{real: real, imaginary: imaginary}
	return c
}
func (c Complex128) Absolute() float64 {
	return math.Hypot(c.real, c.imaginary)
}
