package main

import (
	"fmt"
)

var nums = []int{4, 3, 2, 7, 8, 2, 3, 1} //Output: [5,6]
//var nums = []int{1,1} //Output: [2]

func findDisappearedNumbers(nums []int) []int {
	notFound := false
	pointer := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				break
			}
		}
		if notFound {
			nums[i] = i
			pointer++
		}
	}
	return nums
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) // Output: [5,6]
	fmt.Println(findDisappearedNumbers([]int{1, 1}))                   // Output: [2]
}
