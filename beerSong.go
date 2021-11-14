package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Song() string {
	verses, err := Verses(99, 0)
	finalVerse := verses
	if err != nil {}
	return finalVerse 
}

func Verses(start, stop int) (string, error) {
    if start < 0 || stop < 0 || start < stop || start > 99 {
		return "", fmt.Errorf("Error")
	}
	collectionVerses := []string{}
	for i:=start; i > stop-1; i-- {
		k := strconv.Itoa(i)
		j := strconv.Itoa(i-1)		
		verses :=fmt.Sprintf("%v bottles of beer on the wall, %v bottles of beer.\nTake one down and pass it around, %v bottles of beer on the wall.\n\n", k , k, j)
		collectionVerses = append(collectionVerses, verses)
	} 
	fmt.Println("this is the return! ",strings.Join(collectionVerses, ""))
	return strings.Join(collectionVerses, ","), nil
}

func Verse(n int) (string, error) {
	if n > 99 || n < 0 {
		return "", fmt.Errorf("Error")
	}
	if n == 1 {
		verses := fmt.Sprintf("%v bottle of beer on the wall, %v bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", n , n)
		return verses, nil 
	} else if n == 2 {
		k := strconv.Itoa(n-1)
		verses := fmt.Sprintf("%v bottles of beer on the wall, %v bottles of beer.\nTake one down and pass it around, %v bottle of beer on the wall.\n", n , n, k)
		return verses, nil 
	} else if n == 0 {
		verses := fmt.Sprintln("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.")
		return verses, nil 
	} else {
		k := strconv.Itoa(n-1)
		verses := fmt.Sprintf("%v bottles of beer on the wall, %v bottles of beer.\nTake one down and pass it around, %v bottles of beer on the wall.\n", n , n, k)
	return verses, nil 
	}
}

// Go to the store and buy some more, 99 bottles of beer on the wall.

func main() {
	Song()
}