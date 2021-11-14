package main

import (
	"fmt"
)

func nestedFlatten(arr interface{}, out *[]interface{}) {
	x := arr.([]interface{})
	for _, element := range x {
		fmt.Println("element: ", element)
		if element != nil {
			switch valueType := element.(type) {
			case int:
				*out = append(*out, valueType)
			case []interface{}:
				nestedFlatten(valueType, out)
			}
		}
	}
	fmt.Println(x)
}

func Flatten(nested interface{}) []interface{} {
	fmt.Println("incoming: ", nested)
		result := []interface{}{}
		nestedFlatten(nested, &result)
		return result
}

func main() {
	arr1 := []interface{}{0, 1, 2}
	Flatten(arr1)
	arr2 := []interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8}
	Flatten(arr2)
	arr3 := []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, 100, 4, []interface{}{[]interface{}{[]interface{}{50}}}}, -2}
	Flatten(arr3)
	arr4 := []interface{}{1, []interface{}{2, []interface{}{[]interface{}{3}}, []interface{}{4, []interface{}{[]interface{}{5}}}, 6, 7}, 8}
	Flatten(arr4)
	arr5 := []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, []interface{}{[]interface{}{100}}, interface{}(nil), []interface{}{[]interface{}{interface{}(nil)}}}, -2}
	Flatten(arr5)
}