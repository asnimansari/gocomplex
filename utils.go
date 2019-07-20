package gomath

import "math"

const float64EqualityThreshold = 0.0000001

func AlmostEqual(a, b float64) bool {

	return math.Abs(a-b) <= float64EqualityThreshold
}
