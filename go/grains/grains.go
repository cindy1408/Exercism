package grains

import "fmt"

func Square(number int) (uint64, error) {
	grains := 1
	if number >= 1 && number <=64 {
		if number == 1 {
			return 1, nil 
		}
		for i:=1; i< number; i++ {
			grains = grains*2
		}
		return uint64(grains), nil 
	}
	return 0, fmt.Errorf("no such number %d", number)
}

func Total() uint64 {
	totalGrains := 0 	
		for i:=1; i<=64; i++ {
			squareGrain, _ := Square(i)
			totalGrains = totalGrains + int(squareGrain)
		}
	return uint64(totalGrains)
}
