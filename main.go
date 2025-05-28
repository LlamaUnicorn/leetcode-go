package main

import "fmt"

func dominantIndex(nums []int) int {
	maxValue, maxIndex := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxIndex = i
		}
	}
	for i := 0; i < len(nums); i++ {
		if maxValue < 2*nums[i] && maxValue != nums[i] {
			return -1
		}
	}
	return maxIndex
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0})) // Output: 1
	fmt.Println(dominantIndex([]int{1, 2, 3, 4})) // Output: -1
	fmt.Println(dominantIndex([]int{0, 0, 3, 2})) // Output: -1
}
