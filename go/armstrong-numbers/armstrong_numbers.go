package armstrong

import (
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {
	stringNum := strconv.Itoa(n)
	power := strings.Split(stringNum, "")
	total := 0.0
	for _, value := range power {
		value, _ := strconv.Atoi(value)
		total = total + math.Pow(float64(value), float64(len(power)))
	}
	return int(total) == n
}
