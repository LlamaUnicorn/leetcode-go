package main

import "fmt"

var nums = []int{3, 2, 2, 3}
var val = 3

func main() {
	removeElement(nums, val)
}

//Input: nums = [3,2,2,3], val = 3
//Output: 2, nums = [2,2,_,_]

func removeElement(nums []int, val int) int {
	fmt.Println(nums, val)
	return 1
}
