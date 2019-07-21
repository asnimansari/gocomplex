package gocomplex

import "math"

type Complex128 struct {
	real      float64
	imaginary float64
}

func CreateComplex128(real float64, imaginary float64) Complex128 {
	c := Complex128{real: real, imaginary: imaginary}
	return c
}

// Returns Real Part of Complex number
func (c Complex128) Real() float64 {
	return c.real
}

// Returns Imaginary part of complex number
func (c Complex128) Imaginary() float64 {
	return c.imaginary
}

// Returns the conjugate of a complex number
func (c Complex128) Conjugate() Complex128 {
	return Complex128{c.real, -c.imaginary}
}

// Returns the phase of complex number
func (c Complex128) Phase() float64 {
	return math.Atan(c.imaginary / c.real)
}

// Returns the magnitude of complex numeber
func (c Complex128) Magnitude() float64 {
	return math.Hypot(c.real, c.imaginary)
}

// Returns the magnitude and phase of complex number
func (c Complex128) Polar() (float64, float64) {
	return c.Magnitude(), c.Phase()
}

// Converts Polar Representation of Complex number to a + ib form
func PolarToArgand(magnitude, phase float64) Complex128 {
	return Complex128{real: magnitude * math.Cos(phase), imaginary: magnitude * math.Sin(phase)}
}

// Returns Nth Power of complex number
func (c Complex128) NthPower(n float64) Complex128 {
	magnitude, phase := c.Polar()
	newPhase := phase * n
	newMagnitude := math.Pow(magnitude, n)
	return PolarToArgand(newMagnitude, newPhase)
}

// Returns Square of complex number
func (c Complex128) Square() Complex128 {
	return c.NthPower(2)
}

// Returns Nth root of complex number
func (c Complex128) NthRoot(n float64) Complex128 {
	magnitude, phase := c.Polar()
	newPhase := phase / n
	newMagnitude := math.Pow(magnitude, 1.0/n)
	return PolarToArgand(newMagnitude, newPhase)

}

// Returns Square root of complex number
func (c Complex128) SquareRoot() Complex128 {
	return c.NthRoot(2)
}

// Return Natural Logarithm of complex number
func (c Complex128) Log() Complex128 {
	return Complex128{real: math.Log(c.Magnitude()), imaginary: c.Phase()}
}

// Returns a complex number with same magnitude but opposite sign
func (c Complex128) Negate() Complex128 {
	return Complex128{real: -c.real, imaginary: -c.imaginary}
}

// Checks if both complex numbers are equal
func (c1 Complex128) Compare(c2 Complex128) bool {
	return c1.real == c2.real && c1.imaginary == c2.imaginary

}

func (c Complex128) Vector() []float64 {
	return []float64{c.real, c.imaginary}
}

// TODO (tan)
// TODO (sin)
// TODO (rect)
// TODO (isnan.go)
// TODO (isinf.go)
// TODO (exp.go)s
// TODO (asin.go)
