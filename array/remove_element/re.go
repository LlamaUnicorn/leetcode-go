package remove_element

//Input: nums = [3,2,2,3], val = 3
//Output: 2, nums = [2,2,_,_]

//var nums = []int{3, 2, 2, 3}
//var val = 3

var nums = []int{2}
var val = 3

func removeElement(nums []int, val int) int {
	ptr := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == val {
			continue
		} else {
			nums[ptr] = nums[i]
			ptr++
		}
	}
	nums = nums[0 : len(nums)-ptr]
	return ptr
}
