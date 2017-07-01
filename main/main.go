package main

import (
	"fmt"
)
// TestDrive
func main() {
	x := 6
	isPrime := CheckPrimeNumber(x)
	isPerfect := CheckPerfectNumber(x)
	// print out whether x is prime
	if isPrime {
		fmt.Printf("%v is prime\n", x)
	} else {
		fmt.Printf("%v is not a prime\n", x)
	}
	// print out whether x is perfect
	if isPerfect {
		fmt.Printf("%v is perfect\n", x)
	} else {
		fmt.Printf("%v is not perfect\n", x)
	}
}