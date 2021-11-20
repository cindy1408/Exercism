package purchase

import (
	"fmt"
	"sort"
)

func NeedsLicence(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func ChooseVehicle(option1, option2 string) string {
	stringArr := []string{option1, option2}
	sort.Strings(stringArr)
	return fmt.Sprintf("%v is clearly the better choice.", stringArr[0])
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	cost := 0.0
	if age < 3 {
		cost = 0.8*originalPrice
	} else if age < 10 {
		cost = 0.7*originalPrice
	} else {
		cost = 0.5*originalPrice
	}
	return cost
}
