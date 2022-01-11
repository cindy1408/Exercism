package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {

	stringMap := map[int][]string{}
	messageArr := strings.Split(message, "")
	directionUp := false 
	countj := 1
	countk := rails
	for i:=0; i<len(messageArr)-2; i++{
		if countj == 1 || countj == rails-1 || countk == rails || countk == 1 {
			directionUp = !directionUp
		}
		fmt.Println(directionUp)
		if !directionUp {
			for j:=countj; j<rails; j++ {
				fmt.Println("1: ", countj)
				stringMap[countj] = append(stringMap[countj], messageArr[i+j])
				// if countj == rails-1 {
				// 	countj = 1
				// } else {
				// 	countj +=1
				// }
			}
		} else {
			for k:=countk; k>1; k-- {
				fmt.Println("2: ", countk)
				stringMap[countk] = append(stringMap[countk], messageArr[i+k-1])
				// if countk == 2 {
				// 	countj = rails
				// } else {
				// 	countk -=1
				// }
			}
		}
	}
	result := []string{}
	for _, v := range stringMap {
		result = append(result, v...)
	}
	fmt.Println(stringMap)
	fmt.Println(strings.Join(result, ""))
	return strings.Join(result, "")
}

func Decode(message string, rails int) string {
	panic("Please implement the Decode function")
}
