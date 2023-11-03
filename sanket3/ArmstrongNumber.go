package main

import (
	 "fmt"
    "math"
)

func isArmstrongNumber(n int) bool {
    // Find the number of digits in the number.
    numDigits := 0
    for n > 0 {
        n /= 10
        numDigits++
    }

    // Calculate the sum of the cubes of the digits.
    sumOfCubes := 0
    n = n
    for n > 0 {
        digit := n % 10
        sumOfCubes += int(math.Pow(float64(digit), float64(numDigits)))
        n /= 10
    }

    // Return true if the sum of the cubes is equal to the number, false otherwise.
    return sumOfCubes == n
}

func main() {
    // Get the number from the user.
    fmt.Print("Enter a number: ")
    var n int
    fmt.Scanln(&n)

    // Check if the number is an Armstrong number.
    if isArmstrongNumber(n) {
        fmt.Println("The number is an Armstrong number.")
    } else {
        fmt.Println("The number is not an Armstrong number.")
    }
