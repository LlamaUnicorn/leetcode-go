package two_pointer

import "fmt"

// TwoSumSorted Example 1: Two Sum in Sorted Array
// Given a sorted array, find two numbers that add up to a target sum
// Time: O(n), Space: O(1)
func TwoSumSorted(nums []int, target int) (int, int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return left, right // Found the pair
		} else if sum < target {
			left++ // Need a larger sum, move left pointer right
		} else {
			right-- // Need a smaller sum, move right pointer left
		}
	}

	return -1, -1 // No valid pair found
}

// IsPalindrome Example 2: Check if String is Palindrome
// Use two pointers from both ends moving towards center
// Time: O(n), Space: O(1)
func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// RemoveDuplicates Example 3: Remove Duplicates from Sorted Array (in-place)
// Use slow and fast pointers
// Time: O(n), Space: O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// Slow pointer tracks position for next unique element
	slow := 0

	// Fast pointer explores the array
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	// Length of array with unique elements
	return slow + 1
}

// MaxArea Example 4: Container With Most Water
// Classic two-pointer problem from LeetCode
// Time: O(n), Space: O(1)
func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxWater := 0

	for left < right {
		// Calculate current area
		width := right - left
		minHeight := minFunc(height[left], height[right])
		currentArea := width * minHeight

		maxWater = maxFunc(maxWater, currentArea)

		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

// ThreeSum Example 5: Three Sum (finding triplets that sum to zero)
// Combines sorting with two-pointer technique
// Time: O(n²), Space: O(1) excluding the output
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	if n < 3 {
		return result
	}

	// First, sort the array (in practice, use sort.Ints(nums))
	// Simplified bubble sort for demonstration
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	for i := 0; i < n-2; i++ {
		// Skip duplicate values for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Two-pointer approach for the remaining array
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left and right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// Helper functions
func minFunc(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Run2Pointers() {
	// Test Example 1: Two Sum
	nums1 := []int{2, 7, 11, 15}
	idx1, idx2 := TwoSumSorted(nums1, 9)
	if idx1 != -1 {
		fmt.Printf("Two Sum: indices [%d, %d] -> values [%d, %d]\n",
			idx1, idx2, nums1[idx1], nums1[idx2])
	}

	// Test Example 2: Palindrome
	fmt.Printf("Is 'racecar' palindrome? %v\n", IsPalindrome("racecar"))
	fmt.Printf("Is 'hello' palindrome? %v\n", IsPalindrome("hello"))

	// Test Example 3: Remove Duplicates
	nums2 := []int{1, 1, 2, 2, 3, 4, 4, 5}
	length := RemoveDuplicates(nums2)
	fmt.Printf("Array after removing duplicates: %v (length: %d)\n",
		nums2[:length], length)

	// Test Example 4: Container With Most Water
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Max water area: %d\n", MaxArea(heights))

	// Test Example 5: Three Sum
	nums3 := []int{-1, 0, 1, 2, -1, -4}
	triplets := ThreeSum(nums3)
	fmt.Printf("Triplets that sum to 0: %v\n", triplets)
}
