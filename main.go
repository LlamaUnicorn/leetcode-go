package main

import "fmt"

// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
//var nums = []int{1, 1, 2}

//Input: nums = [0,0,1,1,1,2,2,3,3,4]
//Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
//nums := []int{0,0,1,1,1,2,2,3,3,4}

//Input: [0,0,1,1,1,2,2,3,3,4]
//Output: [0,1,1,2,1,3,2,4]
//Expected: [0,1,2,3,4]

var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

func removeDuplicates(nums []int) int {
	ptr := 0
	uniqueCtr := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		} else if nums[i-1] == nums[i] {
			uniqueCtr++
			if uniqueCtr == 1 {
				ptr++
			}
			continue
		} else {
			nums[ptr] = nums[i]
			uniqueCtr = 0
		}
		fmt.Println(i, ptr, nums[i])
	}
	fmt.Println(nums, ptr)
	return ptr
}

func main() {
	removeDuplicates(nums)
}
