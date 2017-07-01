package main

import (
	"fmt"
	"github.com/ththuong/learning-go/common"
)
// TestDrive
func main() {
	x := 28
	// print out whether x is prime
	if common.CheckPrimeNumber(x) {
		fmt.Printf("%v is prime\n", x)
	} else {
		fmt.Printf("%v is not a prime\n", x)
	}
	// print out whether x is perfect
	if common.CheckPerfectNumber(x) {
		fmt.Printf("%v is perfect\n", x)
	} else {
		fmt.Printf("%v is not perfect\n", x)
	}
}