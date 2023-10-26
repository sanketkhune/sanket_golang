package main

import (
	"fmt"
)

func calculator(operator string, operand1 float64, operand2 float64) float64 {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		panic("invalid operator")
	}
}

func main() {
	// Get the operator and operands from the user.
	var operator string
	fmt.Print("Enter the operator (+, -, *, or /): ")
	fmt.Scanf("%s", &operator)

	var operand1, operand2 float64
	fmt.Print("Enter the first operand: ")
	fmt.Scanf("%f", &operand1)
	fmt.Print("Enter the second operand: ")
	fmt.Scanf("%f", &operand2)

	// Perform the calculation.
	result := calculator(operator, operand1, operand2)

	// Print the result to the console.
	fmt.Printf("The result is: %f\n", result)
}
