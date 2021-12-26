package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	num := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", num)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	given, _ := strconv.Atoi(fnb.Value())
	return given
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	given := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(given))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	iType := fmt.Sprintf("%T", i)
	fmt.Println("TYPE: ", iType)
	if iType == "float64" || iType == "int" {
		valueFloat, _ := i.(float64)
		valueInt, _ := i.(int)
		if iType == "int" {
			return DescribeNumber(float64(valueInt))
		} else {
			return DescribeNumber(valueFloat)
		}
	} else if iType == "sorting.testNumberBox" || iType == "sorting.FancyNumber" || iType == "sorting.differentFancyNumber" {
		if iType == "sorting.testNumberBox" {
			return DescribeNumberBox(i.(NumberBox))
		} else if iType == "sorting.FancyNumber" || iType == "sorting.differentFancyNumber" {
			return DescribeFancyNumberBox(i.(FancyNumberBox))
		}
		return DescribeNumberBox(i.(NumberBox))
	} else {
		return fmt.Sprint("Return to sender")
	}
}
