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

func (c Complex128) GetReal() float64 {
	return c.real
}

func (c Complex128) GetImaginary() float64 {
	return c.imaginary
}

func (c Complex128) Conjugate() Complex128 {
	return Complex128{c.real, -c.imaginary}
}

func (c Complex128) Phase() float64 {
	return math.Atan(c.imaginary / c.real)
}

func (c Complex128) Magnitude() float64 {
	return math.Hypot(c.real, c.imaginary)
}

func (c Complex128) Polar() (float64, float64) {
	return c.Magnitude(), c.Phase()
}

func PolarToArgand(magnitude, phase float64) Complex128 {
	return Complex128{real: magnitude * math.Cos(phase), imaginary: magnitude * math.Sin(phase)}
}

func (c Complex128) NthPower(n float64) Complex128 {
	magnitude, phase := c.Polar()
	newPhase := phase * n
	newMagnitude := math.Pow(magnitude, n)
	return PolarToArgand(newMagnitude, newPhase)
}

func (c Complex128) Square() Complex128 {
	return c.NthPower(2)
}

func (c Complex128) NthRoot(n float64) Complex128 {
	magnitude, phase := c.Polar()
	newPhase := phase / n
	newMagnitude := math.Pow(magnitude, 1.0/n)
	return PolarToArgand(newMagnitude, newPhase)

}
func (c Complex128) SquareRoot() Complex128 {
	return c.NthRoot(2)
}

func (c Complex128) Log() Complex128 {
	return Complex128{real: math.Log(c.Magnitude()), imaginary: c.Phase()}
}

// TODO (tan)
// TODO (sin)

// TODO (rect)
// TODO (pow)
// TODO (polar)
// TODO (isnan.go)
// TODO (isinf.go)
// TODO (exp.go)
// TODO (conj.go)
// TODO (asin.go)
