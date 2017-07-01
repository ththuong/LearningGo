package common

func CheckPrimeNumber(x int) bool {
	isPrime := true
	if x < 2 {
		isPrime = false
	}
	// try all possible factors of x
	// if if x has a factor, then it has one less than or equal to sqrt(x),
	// so for efficiency we only need to check factor <= sqrt(x) or
	// equivalently factor*factor <= x
	for factor := 2; factor*factor <= x; factor++ {
		// if factor divides evenly into x, x is not prime, so break out of loop
		if x % factor == 0 {
			isPrime = false
			break
		}
	}
	// return whether x is prime
	return isPrime
}

func CheckPerfectNumber(x int) bool {
	aliquotSum := 0
	// check (1, 2, ..., x/2) for factors of x
	for i := 1; i <= x/2; i++ {
		if x % i == 0 {
			aliquotSum += i
		}
	}
	// return whether x is perfect
	return x == aliquotSum
}
