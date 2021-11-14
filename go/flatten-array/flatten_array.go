package flatten

func nestedFlatten(arr interface{}, out *[]interface{}) {
	x := arr.([]interface{})
	for _, element := range x {
		if element != nil {
			switch valueType := element.(type) {
			case int:
				*out = append(*out, valueType)
			case []interface{}:
				nestedFlatten(valueType, out)
			}
		}
	}
}

func Flatten(l interface{}) []interface{} {
		result := []interface{}{}
		nestedFlatten(l, &result)
		return result
}