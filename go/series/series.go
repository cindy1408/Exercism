package series


func All(n int, s string) []string {
	result := []string{}
	for i:=0; i<=len(s)-n; i++ {
		element := UnsafeFirst(n, s[i:i+n])
		result = append(result, element)
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
