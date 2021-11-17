package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0,fmt.Errorf("there is an error %v", "error")
	}
	count := 0 
	for n != 1 {
		if n%2 != 0 {
			n = 3*n + 1
			count++
		} else {
			n = n/2
			count++
		}
	}
	return count, nil 
}
