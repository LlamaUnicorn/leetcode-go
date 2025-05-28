package main

import "fmt"

func plusOne(digits []int) []int {
	return []int{1}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // Output: [1,2,4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // Output: [4,3,2,2]
	fmt.Println(plusOne([]int{9}))          // Output: [1,0]
}
