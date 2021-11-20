package wordy

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

func Answer(q string) (int, bool) {

    trimString := []string{}
    q = strings.ReplaceAll(q, "?", "")
    numbers := []int{}
    operators := []string{}
    total := 0

    for _, value := range strings.Split(q, " ") {
        isNumber := true
        _, err := strconv.Atoi(value)
        if err != nil {
            isNumber = false 
        }
        if value == "plus" || value == "add" {
            trimString = append(trimString, "+")
        } else if value == "subtract" || value == "minus" {
            trimString = append(trimString, "-")
        } else if value == "multiplied" {
            trimString = append(trimString, "*")
        } else if value == "divided" {
            trimString = append(trimString, "/")
        } else if isNumber {
            trimString = append(trimString, value)
        } else if value == "cubed" {
            trimString = append(trimString, value)
        }
    }
    fmt.Println("FIRST: ",trimString)
    // if there is only one number it will return the number and true. 
    if len(trimString) != 0 {
        number, _ := strconv.Atoi(trimString[0])
    if len(trimString) == 1 && reflect.TypeOf(number).Kind() == reflect.Int {
        return number, true
    }
    }

    if len(trimString) == 0 {
        return 0, false
    }   
    // get rid of invalid arrays and empty arrays
    for i := 0; i < len(trimString)-1; i+=2 {
        _, err := strconv.Atoi(trimString[i])
        if err != nil || len(trimString) == 0 {
            return 0, false
        }
    }
    // return false if there are 2 numbers in a row 
    for _, value := range trimString {
        num, err := strconv.Atoi(value)
        if err != nil {
            operators = append(operators, value)
        } else {
            numbers = append(numbers, num)
        }
    }
    if trimString[len(trimString) -1] == "+" || trimString[len(trimString)-1] == "-" || trimString[len(trimString)-1] == "*" ||     trimString[len(trimString)-1] == "/" {
        return 0, false
    }
    fmt.Println("numbers: ", numbers)
    fmt.Println("operators: ", operators)


    if len(operators) != 0 && len(operators) == len(numbers) - 1 {
        for i, operator := range operators {
            if i == 0 {
                total = numbers[0]
            }
            if operator == "+" {
                total = total + numbers[i+1]
            } else if operator == "-" {
                total = total - numbers[i+1]
            } else if operator == "*" {
                total = total * numbers[i+1]
            } else if operator == "/" {
                total = total / numbers[i+1]
            } else {
                return 0, false
            }
        } 
    } else {
    return 0, false
    }

    fmt.Println("total", total)
    return total, true
}