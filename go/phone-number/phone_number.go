package phonenumber
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
	//middle array is an int to avoid any messy conversions 
		if updatedArray[0] == 1 && len(updatedArray) == 11 {
			removedOne := updatedArray[1:]
			if removedOne[0] < 10 && removedOne[0]> 1 && removedOne[3] < 10 && removedOne[3] > 1 {
				finalArray := []string{}
				for _, value := range(removedOne) {
					finalArray = append(finalArray, strconv.Itoa(value))				
				}
				return strings.Join(finalArray, ""), nil 
			}
		} else if len(updatedArray) == 10 && updatedArray[0] < 10 && updatedArray[0]> 1 && updatedArray[3] < 10 && updatedArray[3] > 1{
			finalArray := []string{}
				for _, value := range(updatedArray) {
					finalArray = append(finalArray, strconv.Itoa(value))				
				}
				return strings.Join(finalArray, ""), nil 
		}
	
	return "", fmt.Errorf("Error")
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	numberArray := strings.Split(number, "")
	areaCode := []string{}
	for i:=0; i<3; i++ {
		areaCode = append(areaCode, numberArray[i])
	}
	return strings.Join(areaCode, ""), err
}


func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "",err
	}
	numberSplit := strings.Split(number, "")

	firstSection := numberSplit[:3]
	secondSection := numberSplit[3: 6]
	thirdSection := numberSplit[6:]
	lastSection := []string{}
	for _, value := range(thirdSection){
		lastSection = append(lastSection, value)
	}
	
	firstSection = append([]string{"("}, firstSection...)
	firstSection = append(firstSection, ") ")
	secondSection = append(secondSection, "-")
	
	joining := append(firstSection, secondSection...)
	finalJoin := append(joining, lastSection...)
	return strings.Join(finalJoin, ""), nil
}
