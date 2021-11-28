package prime

func Nth(n int) (int, bool) {
	nthPrime := 0
	num := 1
	if n == 0 {
		return 0, false 
	}
	for nthPrime != n {
		if CheckPrime(num) {
			nthPrime = nthPrime + 1
		}
		num++
	}
	return num - 1, true
}

func CheckPrime(num int) bool {
	if num == 1 || num == 0 {
		return false
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
}
