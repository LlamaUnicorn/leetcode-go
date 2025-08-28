package main

import (
	"fmt"
	"strings"

	"github.com/LlamaUnicorn/leetcode-go/patterns/two_pointer"
)

func mergeAlternately(word1 string, word2 string) string {
	// set two pointers equal to the len(word1) and len(word2)
	// loop over word1 and word2 incrementing pointers
	// append selected characters to the result (use buffer for write in the end)
	//
	len1, len2 := len(word1), len(word2)
	var sb strings.Builder

	for i := 0; i < len1; i++ {
		sb.WriteString(string(word1[i]))
		for j := 0; j < len2; j++ {
			sb.WriteString(string(word2[j]))
		}
	}
	fmt.Println(sb.String())
	result := sb.String()
	return result
}

//func main() {
//	// Test cases
//	fmt.Println(mergeAlternately("abc", "pqr")) // apbqcr
//	fmt.Println(mergeAlternately("ab", "pqrs")) // apbqrs
//}

//You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
//
//Return the merged string.
//
//Example 1:
//
//Input: word1 = "abc", word2 = "pqr"
//Output: "apbqcr"
//Explanation: The merged string will be merged as so:
//word1:  a   b   c
//word2:    p   q   r
//merged: a p b q c r
//Example 2:
//
//Input: word1 = "ab", word2 = "pqrs"
//Output: "apbqrs"
//Explanation: Notice that as word2 is longer, "rs" is appended to the end.
//word1:  a   b
//word2:    p   q   r   s
//merged: a p b q   r   s
//Example 3:
//
//Input: word1 = "abcd", word2 = "pq"
//Output: "apbqcd"
//Explanation: Notice that as word1 is longer, "cd" is appended to the end.
//word1:  a   b   c   d
//word2:    p   q
//merged: a p b q c   d

func twoSumSorted(nums []int, target int) (int, int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if target == sum {
			return left, right
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1
}

func main() {
	// Test Example 1: Two Sum
	nums1 := []int{2, 7, 11, 15}
	//idx1, idx2 := two_pointer.TwoSumSorted(nums1, 9)
	idx1, idx2 := twoSumSorted(nums1, 9)
	if idx1 != -1 {
		fmt.Printf("Two Sum: indices [%d, %d] -> values [%d, %d]\n",
			idx1, idx2, nums1[idx1], nums1[idx2])
	}

	// Test Example 2: Palindrome
	fmt.Printf("Is 'racecar' palindrome? %v\n", two_pointer.IsPalindrome("racecar"))
	fmt.Printf("Is 'hello' palindrome? %v\n", two_pointer.IsPalindrome("hello"))

	// Test Example 3: Remove Duplicates
	nums2 := []int{1, 1, 2, 2, 3, 4, 4, 5}
	length := two_pointer.RemoveDuplicates(nums2)
	fmt.Printf("Array after removing duplicates: %v (length: %d)\n",
		nums2[:length], length)

	// Test Example 4: Container With Most Water
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Max water area: %d\n", two_pointer.MaxArea(heights))

	// Test Example 5: Three Sum
	nums3 := []int{-1, 0, 1, 2, -1, -4}
	triplets := two_pointer.ThreeSum(nums3)
	fmt.Printf("Triplets that sum to 0: %v\n", triplets)
}
