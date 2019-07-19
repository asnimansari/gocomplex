package gomath

import "math"

type Complex128 struct {
	real      float64
	imaginary float64
}
type Abc struct {
	i int
}

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
	return math.Sqrt(math.Hypot(c.real, c.imaginary))
}

func (c *Complex128) Conjugate() Complex128 {
	return Complex128{c.real, -c.imaginary}
}
func (c *Complex128) Polar() (float64, float64) {
	magnitude := c.Absolute()
	phase := math.Atan(c.imaginary / c.real)
	return magnitude, phase
}

func PolarToArgand(magnitude, phase float64) Complex128 {
	return Complex128{real: magnitude * math.Cos(phase), imaginary: magnitude * math.Sin(phase)}
}

func (c *Complex128) NthRoot(n float64) Complex128 {
	magnitude, phase := c.Polar()
	newPhase := phase / n
	return PolarToArgand(magnitude, newPhase)

}

// TODO (Square Root)
// TODO (tan)
// TODO (sin)
// TODO (log)
// TODO (rect)
// TODO (pow)
// TODO (polar)
// TODO (isnan.go)
// TODO (isinf.go)
// TODO (exp.go)
// TODO (conj.go)
// TODO (asin.go)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
