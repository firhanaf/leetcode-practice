func convertTemperature(celsius float64) []float64 {
    var kelvin float64
    var fahrenheit float64
    var result []float64
    kelvin = celsius + 273.15
    fahrenheit = celsius * 1.80 +32.00
    result = append(result, kelvin, fahrenheit)
    return result
}