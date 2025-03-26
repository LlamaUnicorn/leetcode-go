package main

import "fmt"

var nums = []int{3, 1, 2, 4} // [2,4,3,1]
//var nums = []int{0}  // [0]

func sortArrayByParity(nums []int) []int {
	right := len(nums) - 1
	for i := 0; i < right; i++ {
		if nums[i]%2 != 0 {
			for j := right; j > i; j-- {
				if nums[j]%2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					right--
					break
				}
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParity(nums))
}
