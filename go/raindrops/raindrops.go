package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	raindrops := []string{}
	if number % 3 == 0 {
		raindrops = append(raindrops, "Pling")
	} 
	if number % 5 == 0 {
		raindrops = append(raindrops, "Plang")
	}  
	if number % 7 == 0 {
		raindrops = append(raindrops, "Plong")
	} 
	if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
	 	stringNumber := strconv.Itoa(number)
		 raindrops = append(raindrops, stringNumber)
	}
	return strings.Join(raindrops[:], "")
}