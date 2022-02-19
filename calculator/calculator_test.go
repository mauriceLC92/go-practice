package calculator_test

import (
	"testing"

	"github.com/mauriceLC92/go-practice/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 18
	got := calculator.Add(5, 6, 7)

	if want != got {
		t.Errorf("Add: want %f but got %f", want, got)
	}
}
