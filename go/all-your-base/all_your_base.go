package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 || outputBase < 2 {
		if inputBase < 2 {
			return nil, errors.New("input base must be >= 2")
		} else {
			return nil, errors.New("output base must be >= 2")
		}
	}
	results := []int{}
	for i, value := range inputDigits {
		if value >= inputBase || value < 0 {
			// fmt.Println("HERE")
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		inputBaseResult := math.Pow(float64(inputBase), float64(len(inputDigits)-i-1))
		results = append(results, value*int(inputBaseResult))
	}
	totalNum := 0
	for _, n := range results {
		totalNum += n
	}
	outputDigits := []int{}
	for totalNum > 0 {
		outputDigits = append([]int{totalNum % outputBase}, outputDigits...)
		totalNum /= outputBase
	}

	if len(outputDigits) == 0 {
		outputDigits = append(outputDigits, 0)
	}
	return outputDigits, nil
}

// all digits must satisfy 0 <= d < input base
