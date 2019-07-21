package gomath

import (

	"math"
	"testing"
)

func TestComplex128_Absolute(t *testing.T) {

	c := Complex128{3, 4}.Magnitude()
	if c != 5 {
		t.Error("Expected 5 , got ", c)
	}
}

func TestComplex128_Phase(t *testing.T) {
	c := Complex128{0, 1}
	phase := c.Phase()
	if phase != math.Pi/2 {
		t.Error("Expected Pi/2 , got ", phase)
	}

	c = Complex128{1, 0}
	phase = c.Phase()
	if phase != 0 {
		t.Error("Expected 0 , got ", phase)
	}
}

func TestComplex128_GetImaginary(t *testing.T) {
	imaginary := Complex128{3, 4}.Imaginary()
	if imaginary != 4 {
		t.Error("Expected 10 , got ", imaginary)
	}
}

func TestComplex128_GetReal(t *testing.T) {
	real := Complex128{3, 4}.Real()
	if real != 3 {
		t.Error("Expected 3 , got ", real)
	}
}

func TestComplex128_Conjugate(t *testing.T) {
	conjugate := Complex128{3, 5}.Conjugate()
	if conjugate.real != 3 {
		t.Error("Expected Real part to be 3 , got ", conjugate.real)
	}
	if conjugate.imaginary != -5 {
		t.Error("Expected Real part to be -5 , got ", conjugate.imaginary)
	}

}

func TestComplex128_Polar(t *testing.T) {
	magnitude, phase := Complex128{1, 1}.Polar()

	if magnitude != 1.4142135623730951 {
		t.Error("Expected Magnitude part to be 1.41421356237 , got ", magnitude)
	}

	if phase != math.Pi/4 {
		t.Error("Expected Phase part to be math.Pi/4 , got ", phase)
	}
}

func TestPolarToArgand(t *testing.T) {
	referenceComplex := Complex128{10.00004, 20.0}
	refernceMagnitude, referencePhase := referenceComplex.Polar()

	complexCreatedFromMagnitudeAndPhase := PolarToArgand(refernceMagnitude, referencePhase)

	if !AlmostEqual(complexCreatedFromMagnitudeAndPhase.real, referenceComplex.real) || !AlmostEqual(complexCreatedFromMagnitudeAndPhase.imaginary, referenceComplex.imaginary) {
		t.Errorf("Expected (%v) , got (%v) ", referenceComplex, complexCreatedFromMagnitudeAndPhase)
	}
}

func TestComplex128_NthPower(t *testing.T) {

	originalComplex := Complex128{1, 1}
	expectedSquaredValue := Complex128{0, 2}
	squaredValue := originalComplex.NthPower(2)
	if !AlmostEqual(squaredValue.real, expectedSquaredValue.real) || !AlmostEqual(squaredValue.imaginary, expectedSquaredValue.imaginary) {
		t.Errorf("Expected (%v) , got (%v) ", expectedSquaredValue, squaredValue)

	}

	originalComplex = Complex128{3, 5}
	expectedSquaredValue = Complex128{-16, 30}
	squaredValue = originalComplex.NthPower(2)
	if !AlmostEqual(squaredValue.real, expectedSquaredValue.real) || !AlmostEqual(squaredValue.imaginary, expectedSquaredValue.imaginary) {
		t.Errorf("Expected (%v) , got (%v) ", expectedSquaredValue, squaredValue)

	}

}
func TestComplex128_NthRoot(t *testing.T) {

	originalComplex := Complex128{0, 2}
	expectedSquaredValue := Complex128{1, 1}
	squaredValue := originalComplex.NthRoot(2)
	if !AlmostEqual(squaredValue.real, expectedSquaredValue.real) || !AlmostEqual(squaredValue.imaginary, expectedSquaredValue.imaginary) {
		t.Errorf("Expected (%v) , got (%v) ", expectedSquaredValue, squaredValue)

	}

	originalComplex = Complex128{-16, 30}
	expectedSquaredValue = Complex128{5, -3}
	squaredValue = originalComplex.NthRoot(2)
	if !AlmostEqual(squaredValue.real, expectedSquaredValue.real) || !AlmostEqual(squaredValue.imaginary, expectedSquaredValue.imaginary) {
		t.Errorf("Expected (%v) , got (%v) ", expectedSquaredValue, squaredValue)

	}
}
