package sieve

func Sieve(limit int) []int {
	primeNums := []int{}
	factors := 0
	for i := 2; i <= limit; i++ {
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				factors = factors + 1
			}
		}
		if factors == 1 {
			primeNums = append(primeNums, i)
		}
		factors = 0
	}
	return primeNums
}
