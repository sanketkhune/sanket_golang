// Bubblesort
package main

import (
	"fmt"
)

func Bubblesort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
func main() {
	array := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(Bubblesort(array))

}
