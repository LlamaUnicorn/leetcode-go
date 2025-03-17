package remove_element

import "fmt"

//Input: nums = [3,2,2,3], val = 3
//Output: 2, nums = [2,2,_,_]

var nums = []int{3, 2, 2, 3}
var val = 3

//var nums = []int{2} //[2]
//var val = 3

func removeElement(nums []int, val int) int {
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ptr] = nums[i]
			ptr++
		}
	}
	fmt.Println(nums)
	return ptr
}
