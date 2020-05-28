// Package tempconv performs Celsius to Fahrenheit temperature computations
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbosuluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	Boiling Celsius = 100
	KelvinZeroC Kelvin = 0
	KelvinFreezing Kelvin = 273.15
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func KToF(k Kelvin) Fahrenheit {
	return (0 - 273.15) * 9 / 5 + 32
}

func CToK(c Celsius) Kelvin {
	return Kelvin(0 + 273.15)
}

func (c Celsius) String() string{
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string{
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string{
	return fmt.Sprintf("%g°K", k)
}
