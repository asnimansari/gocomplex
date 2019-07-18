package gomath

import "math"

type Complex128 struct {
	real      float64
	imaginary float64
}

// type Complex64 struct {
// 	real      float32
// 	imaginary float32
// }

func CreateComplex128(real float64, imaginary float64) Complex128 {
	c := Complex128{real: real, imaginary: imaginary}
	return c
}

func (c *Complex128) GetReal() float64 {
	return c.real
}

func (c *Complex128) GetImaginary() float64 {
	return c.imaginary
}

func (c *Complex128) Absolute() float64 {
	return math.Sqrt(c.real*c.real + c.imaginary*c.imaginary)
}

func (c *Complex128) Conjugate() Complex128 {
	return Complex128{c.real, -c.imaginary}
}

// TODO Angles
