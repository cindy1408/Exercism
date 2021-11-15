package darts

import (
	"math"
)

func Score(x, y float64) int {
	dx := math.Pow(x - 0, 2)
	dy := math.Pow(y - 0, 2)
	maxRadius := dx + dy
	score := 0
	if maxRadius <= 1 {
		score = score + 10
	} else if maxRadius <= 25 {
		score = score + 5
	} else if maxRadius <= 100 {
		score = score + 1
	}
	return score
}