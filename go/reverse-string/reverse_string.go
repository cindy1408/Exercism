package reverse

import (
	"strings"
)

func Reverse(input string) string {
	stringArr := strings.Split(input, "")
	reverseString := []string{}
	for i := len(stringArr)-1; i > -1; i -- {
		reverseString = append(reverseString, stringArr[i])
	}
	return strings.Join(reverseString, "")
}
