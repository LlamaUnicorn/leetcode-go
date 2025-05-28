package largest_number_twice_of_others

//Largest Number At Least Twice of Others
//
//You are given an integer array nums where the largest integer is unique.
//
//Determine whether the largest element in the array is at least twice as much as every other number in the array. If it is, return the index of the largest element, or return -1 otherwise.
//
//Example 1:
//
//Input: nums = [3,6,1,0]
//Output: 1
//Explanation: 6 is the largest integer.
//For every other number in the array x, 6 is at least twice as big as x.
//The index of value 6 is 1, so we return 1.
//Example 2:
//
//Input: nums = [1,2,3,4]
//Output: -1
//Explanation: 4 is less than twice the value of 3, so we return -1.
//
//
//Constraints:
//
//2 <= nums.length <= 50
//0 <= nums[i] <= 100
//The largest element in nums is unique.
//Hide Hint #1
//Scan through the array to find the unique largest element `m`, keeping track of it's index `maxIndex`. Scan through the array again. If we find some `x != m` with `m < 2*x`, we should return `-1`. Otherwise, we should return `maxIndex`.

//fmt.Println(dominantIndex([]int{3, 6, 1, 0})) // Output: 1
//fmt.Println(dominantIndex([]int{1, 2, 3, 4})) // Output: -1
//fmt.Println(dominantIndex([]int{0, 0, 3, 2})) // Output: -1

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
