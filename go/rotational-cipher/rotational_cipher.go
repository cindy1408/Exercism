package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	result := []string{}
	alphabet = append(alphabet, alphabet[:shiftKey]...)
	for _, char := range strings.Split(plain, "") {
		for i := 0; i < len(alphabet)-shiftKey; i++ {
			if strings.ToUpper(char) == char && strings.ToLower(char) == alphabet[i] {
				result = append(result, strings.ToUpper(alphabet[i+shiftKey]))
			} else if strings.ToUpper(char) != char && strings.ToLower(char) == alphabet[i] {
				result = append(result, alphabet[i+shiftKey])
			}
		}
		for _, rune := range char {
			if !unicode.IsLetter(rune) {
				result = append(result, char)
			}
		}
	}
	return strings.Join(result, "")
}
