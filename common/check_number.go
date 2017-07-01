package common

func CheckPrimeNumber(x int) bool {
	if x < 2 {
		return false
	}
	// try all possible factors of x
	// if if x has a factor, then it has one less than or equal to sqrt(x),
	// so for efficiency we only need to check factor <= sqrt(x) or
	// equivalently factor*factor <= x
	for factor := 2; factor*factor <= x; factor++ {
		// if factor divides evenly into x, x is not prime
		if x % factor == 0 {
			return false
		}
	}
	return true
}

func CheckPerfectNumber(x int) bool {
	aliquotSum := 0
	// check (1, 2, ..., x/2) for factors of x
	for i := 1; i <= x/2; i++ {
		if x % i == 0 {
			aliquotSum += i
		}
	}
	return x == aliquotSum
}
