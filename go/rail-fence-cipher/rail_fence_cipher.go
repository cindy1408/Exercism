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
	length := len(message)
	result := make([]rune, length)
	rail := 0
	initDist := rails*2 - 2
	dist := initDist
	idx := -initDist
	for _, r := range message {
		idx += dist
		if idx >= length {
			rail++
			idx = rail
			dist = initDist - 2*rail
			if dist == 0 {
				dist = initDist
			}
			result[idx] = r
		} else {
			result[idx] = r
			dist = initDist - dist
			if dist == 0 {
				dist = initDist
			}
		}
	}
	return string(result)
}
