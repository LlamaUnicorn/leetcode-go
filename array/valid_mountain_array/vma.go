package valid_mountain_array

//Given an array of integers arr, return true if and only if it is a valid mountain array.
//
//Recall that arr is a mountain array if and only if:
//
//arr.length >= 3
//There exists some i with 0 < i < arr.length - 1 such that:
//arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
//
//Example 1:
//
//Input: arr = [2,1]
//Output: false
//Example 2:
//
//Input: arr = [3,5,5]
//Output: false
//Example 3:
//
//Input: arr = [0,3,2,1]
//Output: true
//
//
//Constraints:
//
//1 <= arr.length <= 104
//0 <= arr[i] <= 104
//Hide Hint #1
//It's very easy to keep track of a monotonically increasing or decreasing ordering of elements. You just need to be able to determine the start of the valley in the mountain and from that point onwards, it should be a valley i.e. no mini-hills after that. Use this information in regards to the values in the array and you will be able to come up with a straightforward solution.

// Input:
var arr = []int{0, 3, 2, 1}

//Output: false

//var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func validMountainArray(arr []int) bool {
	ptr := 0
	for i := 1; i <= len(arr)-1; i++ {
		if (arr[i-1] < arr[i]) && (i-ptr == 1) {
			ptr++
			continue
		} else if arr[i-1] > arr[i] {
			if i == len(arr)-1 && ptr != 0 {
				return true
			}
			continue
		}
		return false
	}
	return false
}

//func validMountainArray(arr []int) bool {
//	n := len(arr)
//	if n < 3 {
//		return false
//	}
//
//	peak := 0
//	// Find the peak index where the increasing phase ends
//	for peak < n-1 && arr[peak] < arr[peak+1] {
//		peak++
//	}
//
//	// Peak cannot be the first or last element
//	if peak == 0 || peak == n-1 {
//		return false
//	}
//
//	// Check the decreasing phase from the peak
//	for peak < n-1 && arr[peak] > arr[peak+1] {
//		peak++
//	}
//
//	// If we reached the end of the array, it's a valid mountain
//	return peak == n-1
//}
