package converters

// Temperature represents a SI unit of temperature (in kelvin, K)
type Temperature Unit

const (
	Kelvin Temperature = 1e0
)

// FromCelsius converts temperature from °C to °K
func FromCelsius(t float64) Temperature {
	return Temperature(t + 273.15)
}

// FromDelisle converts temperature from °De to °K
func FromDelisle(t float64) Temperature {
	return Temperature(373.15 - (t * 2 / 3))
}

// FromFahrenheit converts temperature from °F to °K
func FromFahrenheit(t float64) Temperature {
	return Temperature((t + 459.67) * 5 / 9)
}

// FromKelvin converts temperature from °K to °K
func FromKelvin(t float64) Temperature {
	return Temperature(t)
}

// FromNewton converts temperature from °N to °K
func FromNewton(t float64) Temperature {
	return Temperature(t*100/33 + 273.15)
}

// FromRankine converts temperature from °Ra to °K
func FromRankine(t float64) Temperature {
	return Temperature((t-491.67)*5/9 + 273.15)
}

// FromReaumur converts temperature from °Re to °K
func FromReaumur(t float64) Temperature {
	return Temperature(t*5/4 + 273.15)
}

// FromRomer converts temperature from °Ro to °K
func FromRomer(t float64) Temperature {
	return Temperature((t-7.5)*40/21 + 273.15)
}

// Celsius returns the temperature in °C
func (t Temperature) Celsius() float64 {
	return float64(t - 273.15)
}

// Delisle returns the temperature in °De
func (t Temperature) Delisle() float64 {
	return float64((373.15 - t) * 3 / 2)
}

// Fahrenheit returns the temperature in °F
func (t Temperature) Fahrenheit() float64 {
	return float64((t * 9 / 5) - 459.67)
}

// Kelvin returns the temperature in °K
func (t Temperature) Kelvin() float64 {
	return float64(t)
}

// Newton returns the temperature in °N
func (t Temperature) Newton() float64 {
	return float64((t - 273.15) * 33 / 100)
}

// Rankine returns the temperature in °R
func (t Temperature) Rankine() float64 {
	return float64((t-273.15)*9/5 + 491.67)
}

// Reaumur returns the temperature in °Ré
func (t Temperature) Reaumur() float64 {
	return float64((t - 273.15) * 4 / 5)
}

// Romer returns the temperature in °Rø
func (t Temperature) Romer() float64 {
	return float64((t-273.15)*21/40 + 7.5)
}
