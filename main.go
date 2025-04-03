package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	return []int{0}
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) // Output: [5,6]
	fmt.Println(findDisappearedNumbers([]int{1, 1}))                   // Output: [2]
}
