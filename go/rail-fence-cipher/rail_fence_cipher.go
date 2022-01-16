package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {
	messageArr := strings.Split(message, "")
	downRail := true 
	startRail := 0 
	endRail := rails-1
	resultMap := map[int][]string{}
	for _, v := range messageArr {
		if downRail {
			fmt.Println(startRail, v)
			resultMap[startRail] = append(resultMap[startRail], v)
			startRail = startRail + 1
			if startRail == endRail {
				startRail = 0 
				downRail = false
			}
		} else {
			fmt.Println(endRail, v)
			resultMap[endRail] = append(resultMap[endRail], v)
			endRail = endRail - 1
			if endRail == 0 {
				endRail = rails-1 
				downRail = true
			}
		}	
	}
	resultArr := []string{}
	for _, v := range resultMap {
		resultArr = append(resultArr, strings.Join(v, ""))
	}
	fmt.Println(strings.Join(resultArr, ""))
	return strings.Join(resultArr, "")
}

func Decode(message string, rails int) string {
	messageArr := strings.Split(message, "")
	downRail := true 
	startRail := 0 
	endRail := rails - 1 
	resultMap := map[int][]string{}
	// for i:= 0; i < rails - 1; i ++ {
		
	// }
	for _, v := range messageArr {
		if downRail {
			if startRail == 0 {
				resultMap[startRail] = append(resultMap[startRail], v)
			}
			startRail = startRail + 1 
			if startRail == endRail {
				startRail = 0 
				downRail = false
			}
		} else {
			if endRail == 0 {
				resultMap[startRail] = append(resultMap[startRail], v)
			}
			endRail = endRail - 1 
			if endRail == startRail {
				endRail = rails -1 
				downRail = true
			}
		}
	}
	fmt.Println(messageArr, rails)
	fmt.Println(resultMap)
	panic("Please implement the Decode function")
}
