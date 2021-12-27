package atbash

import (
	"regexp"
	"strconv"
	"strings"
)

func Atbash(s string) string {
	cipher := make(map[string]string)
	result := []string{}
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < len(alphabet); i++ {
		cipher[alphabet[i]] = alphabet[len(alphabet)-1-i]
	}
	r := regexp.MustCompile(`[a-zA-Z0-9]`)
	filtered := r.FindAllString(s, -1)
	for _, char := range filtered {
		_, err := strconv.Atoi(char)
		isNum := false
		if err == nil {
			isNum = true
		}
		for k, v := range cipher {
			if strings.ToLower(char) == k && !isNum {
				result = append(result, v)
			}
		}
		if isNum {
			result = append(result, char)
		}
	}
	i := 0
	resultBreak := []string{}
	for _, char := range result {
		if i == 5 {
			resultBreak = append(resultBreak, " ")
			i = 0
			resultBreak = append(resultBreak, char)
			i++
		} else {
			resultBreak = append(resultBreak, char)
			i++
		}
	}
	return strings.Join(resultBreak, "")
}
