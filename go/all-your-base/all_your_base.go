package allyourbase

import (
	"fmt"
	"math"
)
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 || outputBase < 2 {
		return nil, fmt.Errorf("there is an issue with %d or %d",inputBase, outputBase)
	}
	results := []int{}
	for i, value := range inputDigits {
		if value > inputBase && value < 0 {
			return nil, fmt.Errorf("value %d must be less than input base and more than or equal to 0", value)
		}
		result := math.Pow(float64(inputBase), float64(len(inputDigits)-i))
		results = append(results, value * int(result))
	}
	return results, nil
}



// all digits must satisfy 0 <= d < input base