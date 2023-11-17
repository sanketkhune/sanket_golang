package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	result := factorial(number)
	fmt.Printf("The factorial of %d is %d\n", number, result)
}
