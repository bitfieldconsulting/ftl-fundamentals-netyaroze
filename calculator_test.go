package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 16
	got := calculator.Multiply(4, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
