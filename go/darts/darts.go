package darts

func Score(x, y float64) int {
	score := 0
	if x <= 1 && x >= -1 && y <= 1 && y >= -1 {
		score = score + 10
	} else if x <= 5 && x >= -5 && y <= 5 && y >= -5 {
		score = score + 5
	} else if x <= 10 && x >= -10 && y <= 10 && y >= -10 {
		score = score + 1
	}
	return score
}