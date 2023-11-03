package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printPrimeNumbers(lowerBound, upperBound int) {
	for i := lowerBound; i <= upperBound; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func main() {
	// Get the lower and upper bounds from the user.
	fmt.Print("Enter the lower bound: ")
	var lowerBound int
	fmt.Scanln(&lowerBound)

	fmt.Print("Enter the upper bound: ")
	var upperBound int
	fmt.Scanln(&upperBound)

	// Print the prime numbers in the given range.
	printPrimeNumbers(lowerBound, upperBound)
}
