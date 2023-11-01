package main

import ("fmt")

func main() {
	// Create a slice of slices to represent the matrix.
	var matrix [][]int = make([][]int, 3)
	for i := 0; i < 3; i++ {
		// Create a new slice for each row of the matrix.
		matrix[i] = make([]int, 3)
	}
	// Populate the matrix with values.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}
	// Print the matrix.
	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}
}
