package Squares_of_a_Sorted_Array

// Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?
import (
	"fmt"
)

var nums = []int{-4, -1, 0, 3, 10}

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	pos := n - 1
	for left, right := 0, n-1; left <= right; {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			result[pos] = nums[left] * nums[left]
			left++
		} else {
			result[pos] = nums[right] * nums[right]
			right--
		}
		pos--
	}
	fmt.Println(result)
	return result
}
