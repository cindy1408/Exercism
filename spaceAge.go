package main

import (
	"fmt"
	"math"
)

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds/31557600
	age := 0.00
	switch planet {
	case "Mercury": 
		age = earthYears/0.2408467
	case "Venus":
		age = earthYears/0.61519726
	case "Earth":
		age = earthYears
	case "Mars": 
		age = earthYears/1.8808158
	case "Jupiter":
		age = earthYears/11.862615
	case "Saturn": 
		age = earthYears/29.447498
	case "Uranus":
		age = earthYears/84.016846
	case "Neptune":
		age = earthYears/164.79132
	}
	fmt.Println(math.Round(age*100)/100)
	age = math.Round(age*100)/100
	return age
}

func main() {
	Age(1000000000, "Earth")
}