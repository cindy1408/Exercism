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

func DescribeAnything(i interface{}) string {
	// value is the actual value, isInt is bool whether the type is true/false
	int, isInt := i.(int)
	if isInt {
		return DescribeNumber(float64(int))
	}
	float, isFloat := i.(float64)
	if isFloat {
		return DescribeNumber(float)
	}
	numbox, isNumbox := i.(NumberBox)
	if isNumbox {
		return DescribeNumberBox(numbox)
	}
	fancyNumbox, isFancyNumbox := i.(FancyNumberBox)
	if isFancyNumbox {
		return DescribeFancyNumberBox(fancyNumbox)
	}
	return fmt.Sprintln("Return to sender")
}
