package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	spacesRemoved := strings.ReplaceAll(id, " ", "")
	trimmedArr := strings.Split(spacesRemoved, "")
	reverseArr := []string{}
	for i:=len(trimmedArr)-1; i>=0 ; i-- {
		reverseArr = append(reverseArr, trimmedArr[i])
	}
	luhn := false 
	if len(reverseArr) < 2 {
		return false 
	}
	totalNumbers := 0
	for i, isNumber := range reverseArr {
		number, err := strconv.Atoi(isNumber)
		if err != nil {
			return false
		} 
		if i%2 != 0 {
			if int(number*2) > 9 {
				totalNumbers = totalNumbers + (int(number)*2-9)
			} else {
				totalNumbers = totalNumbers + int(number)*2
			}
		} else {
			totalNumbers = totalNumbers + int(number)
		}
	if totalNumbers%10 == 0 {
		luhn = true 
	} else {
		luhn = false
	}
	}
	return luhn
}
