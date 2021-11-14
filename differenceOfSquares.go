package main

import (
	"fmt"
)

func SquareOfSum(n int) int {
	sum := 0
	for i:=0; i< n+1; i++ {
		sum = sum + i 
	}
	squareSum := sum*sum
	fmt.Println(squareSum)
	return squareSum
}

func SumOfSquares(n int) int {
	totalSquare := 0
	for i:=0; i< n+1; i++ {
		squared := i * i
		totalSquare = totalSquare + squared
	}
	fmt.Println(totalSquare)
	return totalSquare
}

func Difference(n int) int {
	difference := SquareOfSum(n) - SumOfSquares(n)
	fmt.Println(difference)
	return difference
}

func main() {
	SquareOfSum(3)
	SumOfSquares(3)
	Difference(3)
}