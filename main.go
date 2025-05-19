package main

import "fmt"

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for i := 0; i < len(nums); i++ {
		leftSum = 0
		rightSum = 0
		for j := 0; j < len(nums); j++ {
			if j < i {
				leftSum += nums[j]
			}
			if j > i {
				rightSum += nums[j]
			}
		}
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) // Output: 3
	//fmt.Println(pivotIndex([]int{1,2,3}))        // Output: -1
	//fmt.Println(pivotIndex([]int{2,1,-1}))        // Output: 0
}
