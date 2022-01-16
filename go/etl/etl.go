package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	transformed := make(map[string]int)
	for num, char := range in {
		for _, eachChar := range char {
			transformed[strings.ToLower(eachChar)] = num
		}
	}
	return transformed
}
