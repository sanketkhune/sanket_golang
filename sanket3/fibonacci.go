package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n < 0 {
		panic("n must be a non-negative integer")
	} else if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	// Get the number from the user.
	fmt.Print("Enter the number of terms in the Fibonacci series: ")
	var n int
	fmt.Scanln(&n)

	// Print the Fibonacci series up to n terms.
	for i := 0; i < n; i++ {
		fmt.Println(fibonacci(i))
	}
}
