package main

import "fmt"

var nums = []int{12, 345, 2, 6, 7896}

func findNumbers(nums []int) int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] / 10
		fmt.Println(nums[i])
	}
	return len(nums)
}
func main() {
	findNumbers(nums)
}
