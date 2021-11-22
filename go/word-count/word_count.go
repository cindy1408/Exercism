package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var answers Frequency
	phrase = strings.ToLower(phrase)
	phrase = strings.ReplaceAll(phrase, ",", " ")
	// count is insensitive (convert to lower case)
	cleaned := regexp.MustCompile(`[a-z0-9' ]`)
		// all punctuation are ignored apart from the apostrophe 
	splitStrings := strings.Split(phrase, "")
	charArr := []string{}
	for _, char := range splitStrings {
		if cleaned.MatchString(char) {
			charArr = append(charArr, char)
		}
	}
	result := strings.Join(charArr, "")
	resultArr := strings.Split(result, " ")
	
	answers = make(map[string]int)
	for _, word := range resultArr {
		splitWord := strings.Split(word, "")
		fmt.Println(splitWord)
		if word == "" {
			continue
		} else {
			if splitWord[0] == "'" && splitWord[len(splitWord)-1] == "'" {
				updatedWordArr := splitWord[1:len(splitWord)-1]
				word = strings.Join(updatedWordArr, "")
			}
		}
		f := answers[word]
		answers[word] = f+1
	}
	return answers
}
