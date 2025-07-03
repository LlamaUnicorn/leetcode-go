package main

import (
	"fmt"
)

//func minSubArrayLen(target int, nums []int) int {
//	return 0
//}

// Sliding Window (O(n))
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
//}

func main() {
	// Test cases
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 //
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}
