package pangram

import (
	"fmt"
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	onlyChar := reg.ReplaceAllString(input, "")
	fmt.Println(onlyChar)
	charArray := strings.Split(onlyChar, "")
	fmt.Println(charArray)
	result := []string{}
	for _, i := range charArray {
		skip := false 
		for _, j := range result {
			if strings.EqualFold(i, j) {
				skip = true
			}
		}
		if !skip {
			result = append(result, i)
		}
	}
	fmt.Println("result: ", result)
	return len(result) == 26
}