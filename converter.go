package converter

import "fmt"

func ConvertTemperature(value float64, fromUnit, toUnit string) (float64, error) {
	if fromUnit == "C" && toUnit == "F" {
		return (value * 9 / 5) + 32, nil
	} else if fromUnit == "F" && toUnit == "C" {
		return (value - 32) * 5 / 9, nil
	} else if fromUnit == "C" && toUnit == "K" {
		return value + 273.15, nil
	} else if fromUnit == "K" && toUnit == "C" {
		return value - 273.15, nil
	}
	return 0, fmt.Errorf("unsupported units: %s to %s", fromUnit, toUnit)
}
