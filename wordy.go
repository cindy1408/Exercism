package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Numbers(word string) int {
	arr := strings.Split(strings.Trim(word, "?"), " ")
	number := 0 
	for _, element := range(arr) {
		num, err := strconv.Atoi(element)
		if err == nil {
			number = num 
		}
	}
	fmt.Println(number)
	return number
}

func Operations(wordy string) (int, error){
	arr := strings.Split(strings.Trim(wordy, "?"), " ")
	numbers := []int{}
	result := 0
	operator := []string{}
	for _, element := range(arr) {
		num, err := strconv.Atoi(element)
		if err == nil {
			numbers = append(numbers, num)
		} else {
			operator = append(operator, element)
		}
	}

	for _, num := range(numbers) {
		for _, ifOperator := range(operator){
			if ifOperator == "add" || ifOperator == "plus"{
				current := num 
				result = result + current 
			} else if ifOperator  == "minus" {
				result = num - result
			} else if ifOperator  == "multiplied" {
				result = result * num
			} else if ifOperator == "divided" {
				result = result/num
			}
		}
	}
	fmt.Println(result)
	return result, nil 
}



func main() {
	Numbers("What is 5?") //5
	Operations("4 add 8?") //12
	Operations("5 minus 3?") // 2
	Operations("What is 6 multiplied by 4?") //24
	Operations("What is 25 divided by 5?") //5
	Operations("What is 5 plus 13 plus 6?") //24
	Operations("What is 3 plus 2 multiplied by 3?") //9
}