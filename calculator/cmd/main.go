package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mauriceLC92/go-practice/calculator"
)

const (
	Add      string = "add"
	Subtract string = "subtract"
	Divide   string = "divide"
	Multply  string = "multiply"
)

func convertInputString(input []string) ([]float64, error) {
	fmt.Printf("%d", len(input))
	res := make([]float64, len(input))

	for i, in := range input {
		output, err := strconv.ParseFloat(in, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting %f to float64 value", output)
		}
		res[i] = output
	}
	return res, nil
}

func printAnswer(ans float64, inputs []string, operation string) {
	fmt.Printf("%s(%s) = %f", operation, inputs, ans)
}

func main() {
	operation := flag.String("operation", Add, "the calculator function to use")
	flag.Parse()

	values := flag.Args()
	valArr := strings.Split(values[0], ",")
	res, err := convertInputString(valArr)
	if err != nil {
		fmt.Printf("error converting input into float4 values: %s", err)
		os.Exit(1)
	}

	switch *operation {
	case Add:
		res := calculator.Add(res...)
		printAnswer(res, values, Add)
	}
}
