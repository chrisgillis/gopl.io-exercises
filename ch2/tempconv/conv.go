package tempconv

//CToF converts a Celsius temp to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//FToC converts a Fahrenheit temp to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius(f - 32*5/9)
}

//CToK converts a Celsius temp to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.16)
}

//KToC converts a Kelvin temp to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.16)
}
