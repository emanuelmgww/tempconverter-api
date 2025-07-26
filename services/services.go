package services

func ConversorCelsius(celsius float64) (float64, float64) {
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	kelvin := celsius + 273.15
	return fahrenheit, kelvin
}

func ConversorFahrenheit(fahrenheit float64) (float64, float64) {
	celsius := (fahrenheit - 32.0) / 1.8
	kelvin := (fahrenheit - 32.0) + 273
	return celsius, kelvin
}

func ConversorKelvin(kelvin float64) (float64, float64) {
	celsius := kelvin - 273.15
	fahrenheit := (kelvin-273.15)*1.8 + 32.0
	return celsius, fahrenheit
}
