package main

import "fmt"

// CelsiusParaKelvin converte uma temperatura em graus Celsius para Kelvin.
func CelsiusParaKelvin(celsius float64) float64 {
	return celsius + 273
}

// KelvinParaCelsius converte uma temperatura em Kelvin para graus Celsius.
func KelvinParaCelsius(kelvin float64) float64 {
	return kelvin - 273
}

// CelsiusParaFahrenheit converte uma temperatura em graus Celsius para Fahrenheit.
func CelsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// FahrenheitParaCelsius converte uma temperatura em Fahrenheit para graus Celsius.
func FahrenheitParaCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// KelvinParaFahrenheit converte uma temperatura em Kelvin para Fahrenheit.
func KelvinParaFahrenheit(kelvin float64) float64 {
	return CelsiusParaFahrenheit(KelvinParaCelsius(kelvin))
}

// FahrenheitParaKelvin converte uma temperatura em Fahrenheit para Kelvin.
func FahrenheitParaKelvin(fahrenheit float64) float64 {
	return CelsiusParaKelvin(FahrenheitParaCelsius(fahrenheit))
}

func main() {
	// Valores equivalentes usados para demonstrar todas as conversões.
	celsius := 25.0
	kelvin := 298.0
	fahrenheit := 77.0

	fmt.Println("Conversões de temperatura")
	fmt.Println()

	// Conversões que partem de Celsius.
	fmt.Printf("%.2f °C = %.2f K\n", celsius, CelsiusParaKelvin(celsius))
	fmt.Printf("%.2f °C = %.2f °F\n", celsius, CelsiusParaFahrenheit(celsius))

	// Conversões que partem de Kelvin.
	fmt.Printf("%.2f K = %.2f °C\n", kelvin, KelvinParaCelsius(kelvin))
	fmt.Printf("%.2f K = %.2f °F\n", kelvin, KelvinParaFahrenheit(kelvin))

	// Conversões que partem de Fahrenheit.
	fmt.Printf("%.2f °F = %.2f °C\n", fahrenheit, FahrenheitParaCelsius(fahrenheit))
	fmt.Printf("%.2f °F = %.2f K\n", fahrenheit, FahrenheitParaKelvin(fahrenheit))
}
