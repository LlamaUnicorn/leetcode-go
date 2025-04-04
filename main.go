package main

import (
	"fmt"
)

var nums = []int{4, 3, 2, 7, 8, 2, 3, 1} //Output: [5,6]
//var nums = []int{1,1} //Output: [2]

func findDisappearedNumbers(nums []int) []int {
	pointer := 0
	for i := 1; i <= len(nums); i++ {
		found := false
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				found = true
				break
			}
		}
		if !found {
			nums[pointer] = i
			pointer++
		}
	}
	return nums[:pointer]
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))        // Output: [5,6]
	fmt.Println(findDisappearedNumbers([]int{1, 1}))                          // Output: [2]
	fmt.Println(findDisappearedNumbers([]int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6})) // Output: [1, 2, 8]
}
