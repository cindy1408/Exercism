package lsproduct

import (
	"fmt"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	stringArr := strings.Split(digits, "")
	if span > len(stringArr) || span < 0 {
		return 0, fmt.Errorf("error")
	}
	if span == 0 {
		return 1, nil
	}
	numArr := []int{}
	for _, number := range stringArr {
		num, err := strconv.Atoi(number)
		if err != nil {
			return 0, err
		}
		numArr = append(numArr, num)
	}	
	clusteredArr := [][]int{}
	for i:= 0; i< len(numArr); i++ {
		if i+span <= len(numArr) {
			clusteredArr = append(clusteredArr, numArr[i:i+span])
		}
	}
	totalArr := []int{}
	for _, arr := range clusteredArr {
		total := arr[0]
		for i, value := range arr {
			if i != 0 {
				total = total * value
			}
		}
		totalArr = append(totalArr, total)
	}
	largest := 0
	for _, v := range totalArr {
		if v > largest {
			largest = v
		}
	}
	return int64(largest), nil
}