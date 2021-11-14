package main

import (
	"fmt"
	"math"
)

func PreparationTime(ingredients []string, numLayer int) int {
	if numLayer == 0 {
		numLayer = 2
	}
	return len(ingredients) * numLayer
}

func Quantities(layer []string) (int, float64) {
	i := 0
	noodles := 0
	sauces := 0
	for i = 0; i < len(layer); i++ {
		if layer[i] == "noodles" {
			x := &noodles
			noodles = *x + 1
		}
		if layer[i] == "sauce" {
			x := &sauces
			sauces = *x + 1
		}
	}
	noodlesNeeded := noodles * 50
	saucesNeeded := float64(sauces) * 0.2
	return noodlesNeeded, saucesNeeded
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList = append(myList, friendsList[len(friendsList)-1])
	return myList
}

func ScaleRecipe(amount[]float64, portions int) []float64 {
	multiplier := float64(portions)/2
	newAmount := []float64{}
	j := 0
	for j=0; j< len(amount); j++ {
		updatedAmount := amount[j] * multiplier
		newAmount = append(newAmount, math.Round(updatedAmount*100)/100)
	}
	fmt.Println(newAmount)
	return newAmount
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	PreparationTime(layers, 3)
	PreparationTime(layers, 0)
	Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
	
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella"}
	AddSecretIngredient(friendsList, myList)

	quantities := []float64{ 1.2, 3.6, 10.5 }
	ScaleRecipe(quantities, 4)
	quantity := []float64{0.6, 300, 1, 0.5, 50, 0.1, 100}
	ScaleRecipe(quantity, 3)
}
