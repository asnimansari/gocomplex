package converters

// Length represents a SI unit of length (in meters, m)
type Length Unit

const (
	Yoctometer              = Meter * 1e-24
	Zeptometer              = Meter * 1e-21
	Attometer               = Meter * 1e-18
	Femtometer              = Meter * 1e-15
	Picometer               = Meter * 1e-12
	Nanometer               = Meter * 1e-9
	Micrometer              = Meter * 1e-6
	Millimeter              = Meter * 1e-3
	Centimeter              = Meter * 1e-2
	Decimeter               = Meter * 1e-1
	Meter            Length = 1e0
	Decameter               = Meter * 1e1
	Hectometer              = Meter * 1e2
	Kilometer               = Meter * 1e3
	ScandinavianMile        = Meter * 1e4
	Megameter               = Meter * 1e6
	Gigameter               = Meter * 1e9
	Terameter               = Meter * 1e12
	Petameter               = Meter * 1e15
	Exameter                = Meter * 1e18
	Zettameter              = Meter * 1e21
	Yottameter              = Meter * 1e24

	Inch    = Meter * 0.0254
	Hand    = Inch * 4
	Foot    = Inch * 12
	Yard    = Foot * 3
	Link    = Chain / 100
	Rod     = Yard * 5.5
	Chain   = Rod * 4
	Furlong = Chain * 10
	Mile    = Meter * 1609.344

	Fathom       = Foot * 6
	Cable        = NauticalMile / 10
	NauticalMile = Meter * 1852

	LunarDistance    = Kilometer * 384400
	AstronomicalUnit = Meter * 149597870700
	LightYear        = Meter * 9460730472580800
)

func (l Length) Yoctometers() float64 {
	return float64(l / Yoctometer)
}

func (l Length) Zeptometers() float64 {
	return float64(l / Zeptometer)
}

func (l Length) Attometers() float64 {
	return float64(l / Attometer)
}

func (l Length) Femtometers() float64 {
	return float64(l / Femtometer)
}

func (l Length) Picometers() float64 {
	return float64(l / Picometer)
}

func (l Length) Nanometers() float64 {
	return float64(l / Nanometer)
}

func (l Length) Micrometers() float64 {
	return float64(l / Micrometer)
}

func (l Length) Millimeters() float64 {
	return float64(l / Millimeter)
}

func (l Length) Centimeters() float64 {
	return float64(l / Centimeter)
}

func (l Length) Decimeters() float64 {
	return float64(l / Decimeter)
}

func (l Length) Meters() float64 {
	return float64(l)
}

func (l Length) Decameters() float64 {
	return float64(l / Decameter)
}

func (l Length) Hectometers() float64 {
	return float64(l / Hectometer)
}

func (l Length) Kilometers() float64 {
	return float64(l / Kilometer)
}

func (l Length) ScandinavianMiles() float64 {
	return float64(l / ScandinavianMile)
}

func (l Length) Megameters() float64 {
	return float64(l / Megameter)
}

func (l Length) Gigameters() float64 {
	return float64(l / Gigameter)
}

func (l Length) Terameters() float64 {
	return float64(l / Terameter)
}

func (l Length) Petameters() float64 {
	return float64(l / Petameter)
}

func (l Length) Exameters() float64 {
	return float64(l / Exameter)
}

func (l Length) Zettameters() float64 {
	return float64(l / Zettameter)
}

func (l Length) Yottameters() float64 {
	return float64(l / Yottameter)
}

func (l Length) Inches() float64 {
	return float64(l / Inch)
}

func (l Length) Hands() float64 {
	return float64(l / Hand)
}

func (l Length) Feet() float64 {
	return float64(l / Foot)
}

func (l Length) Yards() float64 {
	return float64(l / Yard)
}

func (l Length) Rods() float64 {
	return float64(l / Rod)
}

func (l Length) Chains() float64 {
	return float64(l / Chain)
}

func (l Length) Links() float64 {
	return float64(l / Link)
}

func (l Length) Furlongs() float64 {
	return float64(l / Furlong)
}

func (l Length) Miles() float64 {
	return float64(l / Mile)
}

func (l Length) Fathoms() float64 {
	return float64(l / Fathom)
}

func (l Length) Cables() float64 {
	return float64(l / Cable)
}

func (l Length) NauticalMiles() float64 {
	return float64(l / NauticalMile)
}

func (l Length) LunarDistances() float64 {
	return float64(l / LunarDistance)
}

func (l Length) AstronomicalUnits() float64 {
	return float64(l / AstronomicalUnit)
}

func (l Length) LightYears() float64 {
	return float64(l / LightYear)
}
