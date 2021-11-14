package main

import "fmt"


func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for i := 0; i < len(birdsPerDay); i++ {
		total = total+ birdsPerDay[i]
	}
	fmt.Println(total)
	return total
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	startDay := (week-1) * 7
	for i := startDay; i < startDay+7; i++ {
		total = total+ birdsPerDay[i]
	}
	fmt.Println(total)
	return total
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i % 2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		} 
	}
	fmt.Println(birdsPerDay)
	return birdsPerDay
}

func main() {
	// birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	// TotalBirdCount(birdsPerDay)
	// BirdsInWeek(birdsPerDay, 2)
	birdsPerDay := []int{2, 5, 0, 7, 4, 1}
	FixBirdCountLog(birdsPerDay)
}