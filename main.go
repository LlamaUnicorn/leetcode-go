package main

import "fmt"

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

func main() {
	fmt.Println(removeElement(nums, val))
}
