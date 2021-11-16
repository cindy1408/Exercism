// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	stringArr := strings.Split(s, "-")
	joinedString := strings.Join(stringArr, " ")
	var abbreviation []string
	for _, word := range strings.Split(joinedString, " ") {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			if err != nil {
				log.Fatal(err)
			}
		cleanString := reg.ReplaceAllString(word, "")
		word := strings.ToUpper(cleanString)
		cleanStringArr := strings.Split(word, " ")
		for _, eachWord := range cleanStringArr {
			if eachWord != "" {
				letter := strings.Split(eachWord, "")
				abbreviation = append(abbreviation, letter[0])
			}
		}
	}
	return strings.Join(abbreviation, "")
}
