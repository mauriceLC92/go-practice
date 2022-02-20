package calculator

import "errors"

func Add(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	res := inputs[0]
	for _, in := range inputs[1:] {
		res += in
	}
	return res
}

func Subtract(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	res := inputs[0]
	for _, in := range inputs[1:] {
		res -= in
	}
	return res
}

func Multiply(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}

	res := inputs[0]
	for _, in := range inputs {
		res *= in
	}
	return res
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no allowed to divide by zero")
	}

	return a / b, nil
}
