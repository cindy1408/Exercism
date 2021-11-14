// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package main

import "fmt"

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 {
			fmt.Println(true)
			return true
		}
		if year% 400 == 0 {
			fmt.Println(true)
			return true 
		}
	}
	fmt.Println(false)
	return false
}

func main() {
	IsLeapYear(2015)
	IsLeapYear(1970)
	IsLeapYear(1996)
	IsLeapYear(2100)
	IsLeapYear(2000)
	IsLeapYear(1800)
}