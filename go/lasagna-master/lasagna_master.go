package lasagna
import 	"math"
// TODO: define the 'PreparationTime()' function
func PreparationTime(ingredients []string, numLayer int) int {
	if numLayer == 0 {
		numLayer = 2
	}
	return len(ingredients)* numLayer
}

// TODO: define the 'Quantities()' function
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

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList = append(myList, friendsList[len(friendsList)-1])
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount[]float64, portions int) []float64 {
	multiplier := float64(portions)/2
	newAmount := []float64{}
	j := 0
	for j=0; j< len(amount); j++ {
		updatedAmount := amount[j] * multiplier
		newAmount = append(newAmount, math.Round(updatedAmount*100)/100)
	}
	return newAmount
}
