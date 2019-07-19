package gomath

import "testing"

func TestComplex128_Absolute(t *testing.T) {

	c := Complex128{3,4}.Absolute()
	if c!=5{
		t.Error("Expected 5 , got ",c)
	}
}