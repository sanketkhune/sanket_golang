package main

import (
	"fmt"
)

func isPalindrome(num int) bool {
	// Reverse the number.
	reversedNum := 0
	for num > 0 {
		remainder := num % 10
		reversedNum = reversedNum*10 + remainder
		num /= 10
	}

	// Check if the reversed number is equal to the original number.
	return reversedNum == num
}

func main() {
	// Get the number from the user.
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scanln(&num)

	// Check if the number is a palindrome.
	if isPalindrome(num) {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
