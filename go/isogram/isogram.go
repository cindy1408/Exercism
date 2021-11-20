package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	subString := strings.ToLower(word)
	subStringLower := strings.Split(subString, "")
	cleanup := []string{}
	for i:=0; i< len(subStringLower); i++ {
		if subStringLower[i] == "-" || subStringLower[i] ==" "{

		} else {
			cleanup = append(cleanup, subStringLower[i])
		}
	}
	for j:=0; j< len(cleanup); j++ {
		for k:=0; k<len(cleanup); k++ {
			if j != k && cleanup[j] == cleanup[k]{
				return false
			}
		}
	}
	return true
}