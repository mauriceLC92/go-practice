package calculator_test

import (
	"testing"

	"github.com/mauriceLC92/go-practice/calculator"
)

type testCase struct {
	inputs []float64
	want   float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{inputs: []float64{1, 2, 3}, want: 6},
		{inputs: []float64{-1, -2, -3}, want: -6},
		{inputs: []float64{}, want: 0},
		{inputs: []float64{10, 2, 33, 7, -5}, want: 47},
	}

	for _, test := range testCases {
		want := test.want
		got := calculator.Add(test.inputs...)
		if want != got {
			t.Errorf("Add(%v): want %f, got %f",
				test.inputs, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{inputs: []float64{1, 2, 3}, want: -4},
		{inputs: []float64{-1, -2, -3}, want: 4},
		{inputs: []float64{}, want: 0},
		{inputs: []float64{10, 2, 33, 7, -5}, want: -27},
	}

	for _, test := range testCases {
		want := test.want
		got := calculator.Subtract(test.inputs...)
		if want != got {
			t.Errorf("Subtract(%v): want %f, got %f",
				test.inputs, want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{inputs: []float64{1, 2, 3}, want: 6},
		{inputs: []float64{-1, -2, 0}, want: 0},
		{inputs: []float64{}, want: 0},
		{inputs: []float64{10, 2, 33, 7, -5}, want: -231000},
	}

	for _, test := range testCases {
		want := test.want
		got := calculator.Multiply(test.inputs...)
		if want != got {
			t.Errorf("Multiply(%v): want %f, got %f",
				test.inputs, want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	var numerator float64 = 6
	var denominator float64 = 0

	_, err := calculator.Divide(numerator, denominator)

	if err == nil {
		t.Errorf("wanted error for invalid input but got nil instead")
	}
}

func TestDivide(t *testing.T) {
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 4, b: 1, want: 4},
		{a: 2, b: 2, want: 1},
		{a: 10, b: 2, want: 5},
	}

	for _, test := range testCases {
		want := test.want
		got, _ := calculator.Divide(test.a, test.b)

		if want != got {
			t.Errorf("Divide(%f,%f): want %f, got %f", test.a, test.b, want, got)
		}
	}
}
