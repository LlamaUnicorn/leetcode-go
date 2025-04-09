package main

import (
	"fmt"
)

// var nums = []int{4, 3, 2, 7, 8, 2, 3, 1} //Output: [5,6]
var nums = []int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6} // Output: [1, 2, 8]
//var nums = []int{1,1} //Output: [2]

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	fmt.Println(nums)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	//fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))        // Output: [5,6]
	//fmt.Println(findDisappearedNumbers([]int{1, 1}))                          // Output: [2]
	fmt.Println(findDisappearedNumbers([]int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6})) // Output: [1, 2, 8]
}
