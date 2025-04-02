package third_max_num

//Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.
//
//
//Example 1:
//
//Input: nums = [3,2,1]
//Output: 1
//Explanation:
//The first distinct maximum is 3.
//The second distinct maximum is 2.
//The third distinct maximum is 1.
//Example 2:
//
//Input: nums = [1,2]
//Output: 2
//Explanation:
//The first distinct maximum is 2.
//The second distinct maximum is 1.
//The third distinct maximum does not exist, so the maximum (2) is returned instead.
//Example 3:
//
//Input: nums = [2,2,3,1]
//Output: 1
//Explanation:
//The first distinct maximum is 3.
//The second distinct maximum is 2 (both 2's are counted together since they have the same value).
//The third distinct maximum is 1.
//
//
//Constraints:
//
//1 <= nums.length <= 104
//-231 <= nums[i] <= 231 - 1
//
//
//Follow up: Can you find an O(n) solution?

// var nums = []int{3,2,1} // 1
var nums = []int{1, 2} // 2

//fmt.Println(thirdMax([]int{3, 2, 1}))    // Output: 1
//fmt.Println(thirdMax([]int{1, 2}))       // Output: 2
//fmt.Println(thirdMax([]int{2, 2, 3, 1})) // Output: 1
//fmt.Println(thirdMax([]int{5}))          // Output: 5
//fmt.Println(thirdMax([]int{2, 2, 2}))    // Output: 2
//fmt.Println(thirdMax([]int{-1, -2, -3})) // Output: -3

func thirdMax(nums []int) int {
	var max1, max2, max3 *int
	for i := range nums {
		num := nums[i]

		if (max1 != nil && num == *max1) || (max2 != nil && num == *max2) || (max3 != nil && num == *max3) {
			continue
		}

		if max1 == nil || num > *max1 {
			max3 = max2
			max2 = max1
			max1 = &nums[i]
		} else if max2 == nil || num > *max2 {
			max3 = max2
			max2 = &nums[i]
		} else if max3 == nil || num > *max3 {
			max3 = &nums[i]
		}
	}

	if max3 != nil {
		return *max3
	}
	return *max1
}
