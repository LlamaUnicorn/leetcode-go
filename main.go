package main

import "fmt"

var nums = []int{3, 1, 2, 4} // [2,4,3,1]
//var nums = []int{0}  // [0]

func sortArrayByParity(nums []int) []int {
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			for j := right; j > i; j-- {
				if nums[j]%2 == 0 {
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
					right--
					left++
				}
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParity(nums))
}
