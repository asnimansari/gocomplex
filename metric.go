package gomath

const FEET_TO_METER_RATION = 3.28084
const METER_TO_FEET_RATION = 0.3048

type Meter float64

func (n Meter) Centimeter() Meter {
	return n * 100.0
}

func (n Meter) Feet() Meter {
	return n * FEET_TO_METER_RATION
}
