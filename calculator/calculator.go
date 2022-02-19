package calculator

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
