package gocomplex

import "math"

const float64EqualityThreshold = 0.0000001

// AlmostEqual Checks if the difference of two numbers are within a threshold
func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

// IsComplexNumberEqual Check if two complex numbers are almost equal by comparing real and imaginary part
func IsComplexNumberEqual(c1 Complex128, c2 Complex128) bool {
	return AlmostEqual(c1.real, c2.real) && AlmostEqual(c1.imaginary, c2.imaginary)
}
