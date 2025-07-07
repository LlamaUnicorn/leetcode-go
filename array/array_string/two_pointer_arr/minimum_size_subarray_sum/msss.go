package minimum_size_subarray_sum

//https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1299/

//Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
//
//
//
//Example 1:
//
//Input: target = 7, nums = [2,3,1,2,4,3]
//Output: 2
//Explanation: The subarray [4,3] has the minimal length under the problem constraint.
//Example 2:
//
//Input: target = 4, nums = [1,4,4]
//Output: 1
//Example 3:
//
//Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//Output: 0
//
//
//Constraints:
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 104
//
//
//Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

func minSubArrayLen(target int, nums []int) int {
	left, sum, minLen := 0, 0, len(nums)+1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			currLen := right - left + 1
			if currLen < minLen {
				minLen = currLen
			}
			sum -= nums[left]
			left++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

//Prefix + Binary Search (O(n log n))
//import "sort"
//
//func minSubArrayLenBS(target int, nums []int) int {
//	n := len(nums)
//	P := make([]int, n+1)
//	for i := 1; i <= n; i++ {
//		P[i] = P[i-1] + nums[i-1]
//	}
//	minLen := n + 1
//	for i := 0; i < n; i++ {
//		needed := target + P[i]
//		j := sort.Search(n+1, func(j int) bool {
//			return P[j] >= needed
//		})
//		if j <= n {
//			if j-i < minLen {
//				minLen = j - i
//			}
//		}
//	}
//	if minLen == n+1 {
//		return 0
//	}
//	return minLen
//}123

//func main() {
//	// Test cases
//	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
//	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 //
//	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
//}
