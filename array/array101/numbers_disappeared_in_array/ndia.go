package numbers_disappeared_in_array

import "fmt"

//Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.
//
//
//
//Example 1:
//
//Input: nums = [4,3,2,7,8,2,3,1]
//Output: [5,6]
//Example 2:
//
//Input: nums = [1,1]
//Output: [2]
//
//
//Constraints:
//
//n == nums.length
//1 <= n <= 105
//1 <= nums[i] <= n
//
//
//Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
//
//Hide Hint #1
//This is a really easy problem if you decide to use additional memory. For those trying to write an initial solution using additional memory, think counters!
//Hide Hint #2
//However, the trick really is to not use any additional space than what is already available to use. Sometimes, multiple passes over the input array help find the solution. However, there's an interesting piece of information in this problem that makes it easy to re-use the input array itself for the solution.
//Hide Hint #3
//The problem specifies that the numbers in the array will be in the range [1, n] where n is the number of elements in the array. Can we use this information and modify the array in-place somehow to find what we need?

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
