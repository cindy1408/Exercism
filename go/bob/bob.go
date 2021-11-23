package bob

import (
	"fmt"
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.ReplaceAll(remark, "\t", "")
	remark = strings.ReplaceAll(remark, "\n", "")
	remark = strings.ReplaceAll(remark, "\r", "")
	remark = strings.ReplaceAll(remark, " ", "")
	charArr := strings.Split(remark, "")
	compile := regexp.MustCompile(`[a-zA-Z\?]`)
	trimmedArr := []string{}
	for _, char := range charArr {
		if compile.MatchString(char) {
			trimmedArr = append(trimmedArr, char)
		}
	}
	remarkArr := strings.Split(remark, "")
	fmt.Println("remark ARRAY: ", remarkArr)
	
	if len(remarkArr) == 0 || remarkArr[len(remarkArr)-1] == " "{
		return "Fine. Be that way!"
	}
	// first check if it's an empty array first using the original string without any filtering
	for _, char := range remarkArr{
		count := 0 
		if char == "" {
			count ++ 
		}
		if count == len(remarkArr) || remark == "" {
			return "Fine. Be that way!"
		}
	}
	// empty trimmedArr means there are no alphabets in the array nor any question marks
	if len(trimmedArr) == 0 {
		return "Whatever."
	}
	// only contains ? and no alphabets, hence only a question without yelling 
	if len(trimmedArr) == 1 && trimmedArr[0] == "?" {
		return "Sure."
	} else if len(trimmedArr) == 1 && trimmedArr[0] != "?"{
		if strings.ToUpper(trimmedArr[0]) == trimmedArr[0] {
			return "Whoa, chill out!"
		}
	}

	joinedStrings := strings.Join(trimmedArr, "")
	if strings.ToUpper(joinedStrings) == joinedStrings {
		if trimmedArr[len(trimmedArr)-1] == "?"{
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else {
		if trimmedArr[len(trimmedArr)-1] == "?"{
			return "Sure."
		} else {
			return "Whatever."
		}
	}
}
