package gomath

import "testing"

func TestComplex128_Absolute(t *testing.T) {

	c := Complex128{3, 4}.Absolute()
	if c != 5 {
		t.Error("Expected 5 , got ", c)
	}
}

func TestComplex128_GetImaginary(t *testing.T) {
	imaginary := Complex128{3, 4}.GetImaginary()
	if imaginary != 4 {
		t.Error("Expected 10 , got ", imaginary)
	}
}
