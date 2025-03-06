package main

import "fmt"

func main() {
	merge(nums1, 3, nums2, 3)
}

// var nums1 = []int{1, 2, 3, 0, 0, 0}
// var m = 3
//
// var nums2 = []int{2, 5, 6}
// var n = 3
var nums1 = []int{}
var m = 0

var nums2 = []int{1}
var n = 1

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := 0
	if m == 0 {
		idx = len(nums2) - 1
	} else {
		idx = len(nums1) - 1
	}
	num1Index := m
	num2Index := n
	for i := idx; i > 0; i-- {
		fmt.Println(nums1[i])
		if nums1[num1Index-1] >= nums2[num2Index-1] {
			nums1[i] = nums1[num1Index-1]
			num1Index--
			fmt.Println("if", nums1, num1Index)
		} else {
			nums1[i] = nums2[num2Index-1]
			num2Index--
			fmt.Println("else", nums1, num2Index)
		}
	}
	fmt.Println(nums1, nums2)
}
