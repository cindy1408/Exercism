package romannumerals

import (
	"fmt"
	"strconv"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
inputString := strconv.Itoa(input)
romanNumerals := []string{}
for i, rune := range inputString {
	num, err := strconv.Atoi(string(rune))
	if err != nil {
		return "", err
	}
	if input < 0 || input == 0 || input > 3000 {
		return "", fmt.Errorf("the number %d is incorrect", num)
	}
	placeValue := len(inputString) - (i+1)
	fmt.Println("placeValue", placeValue)
	if placeValue == 3 {
		if num > 0 && num <= 3 {
			for i:=0; i<num; i++ {
			romanNumerals = append(romanNumerals, "M")
			}
		} else if num == 0 {
			continue
		} else {
			return "", err
		}
	} else if placeValue == 2 {
		if num > 5 && num < 10 {
			romanNumerals = aboveAroundFive(num, romanNumerals, placeValue)
		} else if num == 4 {
			romanNumerals = belowAroundFive(num, romanNumerals, placeValue)
		} else if num == 0 {
			continue
		} else if num == 5 {
			romanNumerals = append(romanNumerals, "D")
		} else {
			for i:=0; i<num; i++ {
				romanNumerals = append(romanNumerals, "C")
			}
		}
	} else if placeValue == 1 {
		if num > 5 && num < 10 {
			fmt.Println("HERERRRRRR!")
			romanNumerals = aboveAroundFive(num, romanNumerals, placeValue)
		} else if num == 4 {
			romanNumerals = belowAroundFive(num, romanNumerals, placeValue)
		} else if num == 0 {
			continue
		} else if num == 5 {
			romanNumerals = append(romanNumerals, "L")
		} else {
			for i :=0; i<num; i++ {
				romanNumerals = append(romanNumerals, "X")
			}
		}
	} else if placeValue == 0 {
		if num > 5 {
			romanNumerals = aboveAroundFive(num, romanNumerals, placeValue)
		} else if num < 5 && num > 3 {
			fmt.Println("It's HERE!")
			romanNumerals = belowAroundFive(num, romanNumerals, placeValue)
		} else if num == 5 {
			romanNumerals = append(romanNumerals, "V")
		} else if num == 0 {
			continue
		} else {
			for i :=0; i<num; i++ {
				romanNumerals = append(romanNumerals, "I")
			}
		}
	}
}
fmt.Println(romanNumerals)

return strings.Join(romanNumerals, ""), nil
}

func aboveAroundFive(num int, romanNumerals []string, placeValue int) []string {
	romanNumeral := []string{"I", "V", "X", "L", "C", "D", "M"}
	remainder := num - 5 
	if placeValue == 0 && num == 6 || placeValue == 0 && num == 7 || num == 8 {
		j := 1
		romanNumerals = append(romanNumerals, romanNumeral[j])
		for i:=0; i< remainder; i++ {
			romanNumerals = append(romanNumerals, romanNumeral[j-1])
		}
	} else if placeValue == 1 && num == 6 || num == 7 || num == 8 {
		j := 3
		romanNumerals = append(romanNumerals, romanNumeral[j])
		for i:=0; i< remainder; i++ {
			romanNumerals = append(romanNumerals, romanNumeral[j-1])
		}
	} else if placeValue == 2 && num == 6 || num == 7 || num == 8 {
		j := 5
		romanNumerals = append(romanNumerals, romanNumeral[j])
		for i:=0; i< remainder; i++ {
			romanNumerals = append(romanNumerals, romanNumeral[j-1])
		}
	} else if placeValue == 0 && num == 9 {
		j := 0 
		romanNumerals = append(romanNumerals, romanNumeral[j])
		romanNumerals = append(romanNumerals, romanNumeral[j+2])
	} else if placeValue == 1 && num == 9 {
		j := 2 
		romanNumerals = append(romanNumerals, romanNumeral[j])
		romanNumerals = append(romanNumerals, romanNumeral[j+2])
	} else if placeValue == 2 && num == 9 {
		j := 4
		romanNumerals = append(romanNumerals, romanNumeral[j])
		romanNumerals = append(romanNumerals, romanNumeral[j+2])
	}
	return romanNumerals
}

func belowAroundFive(num int, romanNumerals []string, placeValue int) []string {
	romanNumeral := []string{"I", "V", "X", "L", "C", "D", "M"}
	remainder := 5 - num 
	if placeValue == 0 {
		j := 1
		for i:=0; i< remainder; i++ {
			romanNumerals = append(romanNumerals, romanNumeral[j-1])
		}
		romanNumerals = append(romanNumerals, romanNumeral[j])
	} else if placeValue == 1 {
		j := 3
		for i:=0; i< remainder; i++ {
			romanNumerals = append(romanNumerals, romanNumeral[j-1])
		}
		romanNumerals = append(romanNumerals, romanNumeral[j])
	} else if placeValue == 2 {
		j := 5
		for i:=0; i< remainder; i++ {
			romanNumerals = append(romanNumerals, romanNumeral[j-1])
		}
		romanNumerals = append(romanNumerals, romanNumeral[j])
	}
	return romanNumerals
}

// I = 1  
//V = 5  
// X = 10
// L = 50
// C = 100
// D = 500
// M = 1000