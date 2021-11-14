package wordy

import (
	"strconv"
	"strings"
)

func Answer(word string) (int, error) {
	arr := strings.Split(strings.Trim(word, "?"), " ")
	number := 0 
	for _, element := range(arr) {
		num, err := strconv.Atoi(element)
		if err == nil {
			number = num 
		}
	}
	return number, nil 
}