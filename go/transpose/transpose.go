package transpose

import (
	"fmt"
	"sort"
	"strings"
)

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}
	inputMap := map[int][]string{}
	updatedInput := [][]string{}
	lastElement := []string{}
	previousArr := strings.Split(input[0], "")
	for i, v := range input {
		currentArr := strings.Split(v, "")
		if i != 0 {
			previousArr = strings.Split(input[i-1], "")
		}
		if i == len(input)-1 {
			lastElement = append(lastElement, strings.Split(v, "")...)
		}
		fmt.Println("current array", currentArr)
		fmt.Println("previous array", previousArr)
		if len(previousArr) < len(currentArr) {
			diff := len(currentArr) - len(previousArr)
			for j := 0; j < diff; j++ {
				previousArr = append(previousArr, " ")
				fmt.Println("THIS ONE: ", previousArr[len(previousArr)-1])
			}
		}
		fmt.Println(len(previousArr))
		if i != 0 {
			updatedInput = append(updatedInput, previousArr)
		}
	}
	updatedInput = append(updatedInput, lastElement)
	fmt.Println("UPDATED INPUT: ", updatedInput)

	for _, currentRow := range updatedInput {
		fmt.Println("LENGTH: ", len(currentRow))
		fmt.Println("currentROW", currentRow)
		for k, d := range currentRow {
			inputMap[k] = append(inputMap[k], d)
		}
		fmt.Println(currentRow)
	}
	fmt.Println("input map: ", inputMap)
	result := []string{}
	index := []int{}
	for eachIndex := range inputMap {
		index = append(index, eachIndex)
	}
	sort.Ints(index)
	for _, order := range index {
		chars := inputMap[order]
		result = append(result, strings.Join(chars, ""))
	}
	return result
}
