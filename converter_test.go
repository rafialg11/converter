package converter

import (
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	result, err := ConvertTemperature(0, "C", "F")
	if err != nil || result != 32 {
		t.Errorf("Expected 32, got %f", result)
	}

	result, err = ConvertTemperature(32, "F", "C")
	if err != nil || result != 0 {
		t.Errorf("Expected 0, got %f", result)
	}

	result, err = ConvertTemperature(0, "C", "K")
	if err != nil || result != 273.15 {
		t.Errorf("Expected 273.15, got %f", result)
	}
}

func TestConvertWeight(t *testing.T) {
	result, err := ConvertWeight(0, "kg", "lb")
	if err != nil || result != 0 {
		t.Errorf("Expected 0, got %f", result)
	}

	result, err = ConvertWeight(0, "lb", "kg")
	if err != nil || result != 0 {
		t.Errorf("Expected 0, got %f", result)
	}
}
