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
