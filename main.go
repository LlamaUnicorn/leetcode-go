package main

import (
	"fmt"
)

func rotate(nums []int, k int) {

}

func main() {
	// Test cases
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	//rotate([]int{-1, -100, 3, 99}, 2)
	fmt.Println(arr) // [5,6,7,1,2,3,4]
	//fmt.Println(arr) // [3,99,-1,-100]
}
