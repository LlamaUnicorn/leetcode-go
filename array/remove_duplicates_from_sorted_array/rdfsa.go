package remove_duplicates_from_sorted_array

import "fmt"

// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
var nums = []int{1, 1, 2}

//Input: nums = [0,0,1,1,1,2,2,3,3,4]
//Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
//nums := []int{0,0,1,1,1,2,2,3,3,4}

func removeDuplicates(nums []int) int {
	ptr := 0
	uniquePtr := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, ptr, uniquePtr)
	}
	return 0
}
