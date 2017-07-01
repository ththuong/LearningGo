package main

import (
	"fmt"
	"github.com/ththuong/LearningGo/common"
)
// TestDrive
func main() {
	x := 6
	x := 8127
	isPrime := common.CheckPrimeNumber(x)
	isPerfect := common.CheckPerfectNumber(x)
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