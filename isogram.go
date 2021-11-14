package main

import (
	"fmt"
	"strings"
)

// func IsIsogram(word string) bool {
// 	//split the string
// 	subString := strings.ToLower(word)
// 	subStringLower := strings.Split(subString, "")
// 	for i:=0; i< len(subStringLower); i++ {
// 		for j:=0; j< len(subStringLower); j++ {
// 			if i != j &&  subStringLower[i] == subStringLower[j] {
// 				if subStringLower[i] != "-" || subStringLower[i] != " " {
// 					fmt.Println(false)
// 					return false 
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(true)
// 	return true
// }

func IsIsogram(word string) bool {
	//split the string
	subString := strings.ToLower(word)
	subStringLower := strings.Split(subString, "")
	cleanup := []string{}
	for i:=0; i< len(subStringLower); i++ {
		if subStringLower[i] == "-" || subStringLower[i] ==" "{

		} else {
			cleanup = append(cleanup, subStringLower[i])
		}
	}
	fmt.Println(cleanup)
	for j:=0; j< len(cleanup); j++ {
		for k:=0; k<len(cleanup); k++ {
			if j != k && cleanup[j] == cleanup[k]{
				fmt.Println(false)
				return false
			}
		}
	}
	fmt.Println(true)
	return true
}

func main() {
	IsIsogram("Hello")
	IsIsogram("abc")
	IsIsogram("hey-lo-ds")
	IsIsogram("isogram")
}