package transpose

import (
	"fmt"
	"strings"
)

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}
	inputMap := map[int][]string{}
	for _, v := range input {
		for i, d := range strings.Split(v, "") {
			inputMap[i] = append(inputMap[i], d)
		}
	}
	fmt.Println(inputMap)
	result := []string{}
	for f, _ := range inputMap {
		result = append(result, inputMap[f]...)
	}
	return result
}
