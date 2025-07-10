package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	return []int{1, 2}
}
func main() {
	// Test cases
	fmt.Println(getRow(3)) // [1,3,3,1]
	fmt.Println(getRow(0)) // [1]
	fmt.Println(getRow(1)) // [1,1]
}

// Input: rowIndex = 3
// Output: [1,3,3,1]
// Example 2:
//
// Input: rowIndex = 0
// Output: [1]
// Example 3:
//
// Input: rowIndex = 1
// Output: [1,1]
