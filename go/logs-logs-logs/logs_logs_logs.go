package logs

import "unicode/utf8"

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
	return result[0]
}

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

func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) < limit+1 {
		return true 
	}
	return false 
}
