package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	phoneArray := strings.Split(phoneNumber, "")
	updatedArray := []int{}
	for _, value := range phoneArray {
		currentInt, err := strconv.Atoi(value)
		if err == nil {
		if currentInt < 10 && currentInt >= 0 {
			updatedArray = append(updatedArray, currentInt)
		}
		}
	}
		if updatedArray[0] == 1 && len(updatedArray) == 11 {
			removedOne := updatedArray[1:]
			if removedOne[0] < 10 && removedOne[0]> 1 && removedOne[3] < 10 && removedOne[3] > 1{
				finalArray := []string{}
				for _, value := range(removedOne) {
					finalArray = append(finalArray, strconv.Itoa(value))				
				}
				fmt.Println("PASS - ", strings.Join(finalArray, ""))
				return strings.Join( finalArray, ""), nil 
			
			}
		} else if len(updatedArray) == 10 && updatedArray[0] < 10 && updatedArray[0]> 1 && updatedArray[3] < 10 && updatedArray[3] > 1{
			finalArray := []string{}
				for _, value := range(updatedArray) {
					finalArray = append(finalArray, strconv.Itoa(value))				
				}
				fmt.Println("PASS - ", strings.Join(finalArray, ""))
				return strings.Join(finalArray, ""), nil 
		}
	
	fmt.Println("FAILED ", phoneArray)
	return "", fmt.Errorf("ERROR")
}


func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	numberArray := strings.Split(number, "")
	fmt.Println(numberArray)
	areaCode := []string{}
	for i:=0; i<3; i++ {
		areaCode = append(areaCode, numberArray[i])
	}
	fmt.Println(strings.Join(areaCode, ""))
	return strings.Join(areaCode, ""), err
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "",err
	}
	numberSplit := strings.Split(number, "")

	fmt.Println("this is the string number, ", numberSplit)
	firstSection := numberSplit[:3]
	fmt.Println(firstSection)
	secondSection := numberSplit[3: 6]
	fmt.Println(secondSection)
	thirdSection := numberSplit[6:]
	lastSection := []string{}
	for _, value := range(thirdSection){
		lastSection = append(lastSection, value)
	}
	fmt.Println("thirdSection ", thirdSection)
	
	firstSection = append([]string{"("}, firstSection...)
	firstSection = append(firstSection, ") ")
	secondSection = append(secondSection, "-")
	fmt.Println("last Section ", lastSection)
	
	joining := append(firstSection, secondSection...)
	fmt.Println("joining ",joining)
	fmt.Println("thirdSection ",thirdSection)
	finalJoin := append(joining, lastSection...)
	fmt.Println(strings.Join(finalJoin, ""))
	return strings.Join(finalJoin, ""), nil
}

func main() {
	Number("+1 (613)-995-0253")
	Number("(223) 456-7890")
	Number("223.456.7890")
	Number("223 456   7890   ")
	Number("123456789")
	Number("1 (223) 156-7890")
	AreaCode("+1 (613)-995-0253")
	AreaCode("(223) 456-7890")
	AreaCode("223.456.7890")
	AreaCode("223 456   7890   ")
	Format("+1 (613)-995-0253")
	Format("(223) 456-7890")
}