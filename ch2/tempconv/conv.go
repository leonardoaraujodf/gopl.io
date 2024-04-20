package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
// KtoC converts a Kelvin temperature to Celsius
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }
// KtoF converts a Kelvin temperature to Fahrenheit
func KtoF(k Kelvin) Fahrenheit{ return CToF(KtoC(k)) }
