package isbn

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	reg := regexp.MustCompile(`[0-9X']`)
	splitIsbn := strings.Split(isbn, "")
	cleanedIsbn := []int{}
	for i, char := range splitIsbn {
		if reg.MatchString(char) {
			if char == "X"  && i == len(splitIsbn)-1 {
				cleanedIsbn = append(cleanedIsbn, 10)
			} else {
				num, err := strconv.Atoi(char)
				cleanedIsbn = append(cleanedIsbn, num)
				if err != nil {
					return false
				}
			}
		}
	}
	total := 0
	if len(cleanedIsbn) == 10 {
		for i:= 10; i>=1; i -- {
			total = total + (cleanedIsbn[len(cleanedIsbn)-i] * i)
		}
		if total % 11 == 0 {
			return true
		} 
	}
	return false
}
