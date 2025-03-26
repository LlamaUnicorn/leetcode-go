package remove_element2

//var nums = []int{3, 2, 2, 3}
//var val = 3

var nums = []int{2}
var val = 3

//[2]

func removeElement(nums []int, val int) int {
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ptr] = nums[i]
			ptr++
		}
	}
	return ptr
}
