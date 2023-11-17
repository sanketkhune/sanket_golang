package main

import (
	"fmt"
	"math"
)

func main() {
	var lower, upper int

	fmt.Print("Enter the lower limit: ")
	fmt.Scanln(&lower)

	fmt.Print("Enter the upper limit: ")
	fmt.Scanln(&upper)

	for num := lower; num <= upper; num++ {
		if num <= 1 {
			continue
		}

		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Print(num, " ")
		}
	}
}
