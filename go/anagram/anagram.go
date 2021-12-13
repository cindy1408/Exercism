package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	listOfAnagram := []string{}
	for _, candidateString := range candidates {
		if !strings.EqualFold(candidateString, subject) {
			anagram := reflect.DeepEqual(convertToMap(subject), convertToMap(candidateString))
			if anagram {
				listOfAnagram = append(listOfAnagram, candidateString)
			}
		}
	}
	return listOfAnagram
}

func convertToMap(sentence string) map[string]int {
	resultMap := map[string]int{}
	sentence = strings.ToLower(sentence)
	splitSentence := strings.Split(sentence, "")
	for _, v := range splitSentence {
		resultMap[v] = resultMap[v] + 1 
	}
	return resultMap
}