package sort_array_by_parity

//Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
//
//Return any array that satisfies this condition.
//
//
//
//Example 1:
//
//Input: nums = [3,1,2,4]
//Output: [2,4,3,1]
//Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
//Example 2:
//
//Input: nums = [0]
//Output: [0]
//
//
//Constraints:
//
//1 <= nums.length <= 5000
//0 <= nums[i] <= 5000

var nums = []int{3, 1, 2, 4} // [2,4,3,1]
//var nums = []int{0}  // [0]

func sortArrayByParity(nums []int) []int {
	right := len(nums) - 1
	for i := 0; i < right; i++ {
		if nums[i]%2 != 0 {
			for j := right; j > i; j-- {
				if nums[j]%2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					right--
					break
				}
			}
		}
	}
	return nums
}
