package main

import (
	"fmt"
	"unicode/utf8"
)

func Application(log string) string {
	runeArr := []rune(log) 
	result := []string{}
	for _, runeChar := range(runeArr) {
		if runeChar == rune('❗') {
			result = append(result, "recommendation")
		} else if runeChar == rune('🔍') {
			result = append(result, "search")
		} else if  runeChar == rune('☀') {
			result = append(result, "weather")
		}
	}
	if len(result) == 0 {
		result = append(result, "default")
	}
	fmt.Println(result[0])
	return result[0]
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeArr := []rune(log)
	newRuneArr := []rune{}
	for _, runeChar := range(runeArr) {
		if runeChar == oldRune {
			newRuneArr = append(newRuneArr, newRune)
		} else {
			newRuneArr = append(newRuneArr, runeChar)
		}
	}
	return string(newRuneArr)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) < limit+1 {
		return true 
	}
	return false 
}


func main() {
	Application("❗ recommended product") // recommendation
	Application("executed search 🔍") // search
	Application("forecast: ☀ sunny") // weather
	Application("error: could not proceed") //default
	Application("❗ recommended search product 🔍") // recommendation 
}