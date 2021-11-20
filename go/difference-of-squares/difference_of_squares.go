package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i:=0; i< n+1; i++ {
		sum = sum + i 
	}
	squareSum := sum*sum
	return squareSum
}

func SumOfSquares(n int) int {
	totalSquare := 0
	for i:=0; i< n+1; i++ {
		squared := i * i
		totalSquare = totalSquare + squared
	}
	return totalSquare
}
func Difference(n int) int {
	difference := SquareOfSum(n) - SumOfSquares(n)
	return difference
}
