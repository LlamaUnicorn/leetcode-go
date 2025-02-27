package main

import "fmt"

var nums = []int{-4, -1, 0, 3, 10}

func sortedSquares(nums []int) []int {
	var result []int
	for i, j := 0, len(nums); i != j; {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			result = append(result, nums[i])
			i++
		} else {
			result = append(result, nums[j])
			j--
		}
	}
	fmt.Println(result)
	return result
}
func main() {
	sortedSquares(nums)
}
