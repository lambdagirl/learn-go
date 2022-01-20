package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Farenheit temperature to Celsius.
func FtoC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }
