package prime

func Factors(n int64) []int64 {
	primeFactors := []int64{}
	if n == 1 {
		return primeFactors
	}
	current := n 
	for i:=2; i<=int(n); i++ {
		if int(current)%i == 0 {
			primeFactors = append(primeFactors, int64(i))
			current = current/int64(i)
			i = i-1
			if current == 1 {
				break
			}
		}
	}
	return primeFactors
}


func IsPrime(num int) bool {
	for i:=2; i<num; i++ {
		if num%i == 0 {
			return false 
		}
	}
	return true 
}