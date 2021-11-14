package main

import "fmt"

func Score(x, y float64) int {
	score := 0

	if x <= 1 && x >= -1 && y <= 1 && y >= -1 {
		score = score + 10
	} else if x <= 5 && x >= -5 && y <= 5 && y >= -5 {
		score = score + 5
	} else if x <= 10 && x >= -10 && y <= 10 && y >= -10 {
		score = score + 1
	}
	fmt.Println(score)
	return score
}

func main() {
	Score(-9, 9) // 1
	Score(0, 10) // 1
	Score(-5, 0) // 5
	Score(0, -1) // 10
	Score(0, 0) // 10
	Score(-9, 9) // 0 
	Score(0.8, -0.8) // 5
	Score(-3.6, -3.6) // 1
	Score(7.1, -7.1) // 0 
}

// If the dart lands outside the target, player earns no points (0 points).
// If the dart lands in the outer circle of the target, player earns 1 point.
// If the dart lands in the middle circle of the target, player earns 5 points.
// If the dart lands in the inner circle of the target, player earns 10 points.
