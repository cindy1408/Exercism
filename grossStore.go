package main

import "fmt"

func Units() map[string]int {
	units:= make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if exists != false {
		bill[item] = units[unit]
		return true
	}
	return false
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	_, unitExists := units[unit]
   	numberItemsBill := bill[item]
	numberToRemove := units[unit]
    fmt.Println(bill)
	if itemExists == false {
		return false
	} 
	if unitExists == false {
		return false 
	}
	remainderNumberItem := numberItemsBill - numberToRemove
    fmt.Println("newQuantity" ,remainderNumberItem)
	if remainderNumberItem < 0 {
		return false
	}
	if remainderNumberItem == 0 {
		delete(bill, item)
		return true
	} 
	// amountPointer := &remainderNumberItem
	bill[item] = remainderNumberItem
	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	_, exists := bill[item]
	if exists == false {
		return 0, false
	}
	onBill := bill[item]
	return onBill, true
}

func main() {
	// units := Units()
	// fmt.Println(units)

	// bill := NewBill()
	// fmt.Println(bill)

	// bill := NewBill()
	// units := Units()
	// ok := AddItem(bill, units, "carrot", "dozen")
	// fmt.Println(ok)

	bill := NewBill()
	units := Units()
	ok := RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)


	// bill := map[string]int{"carrot", 12, "grapes", 3}
	// qty, ok := GetItem(bill, "carrot")
	// fmt.Println(qty)
	// // Output: 12
	// fmt.Println(ok)
	// // Output: true
}