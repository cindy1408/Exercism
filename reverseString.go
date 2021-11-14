package main

import (
	"fmt"
	"strings"
)

func Reverse(input string) string {
	stringArr := strings.Split(input, "")
	reverseString := []string{}
	for i := len(stringArr)-1; i > -1; i -- {
		reverseString = append(reverseString, stringArr[i])
	}
	fmt.Println(strings.Join(reverseString, ""))
	return strings.Join(reverseString, "")
}

func main() {
	Reverse("cool")
}