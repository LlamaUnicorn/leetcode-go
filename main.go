package main

import "fmt"

func main() {
	merge(nums1, m, nums2, n)
}

var nums1 = []int{1, 2, 3, 0, 0, 0}
var m = 3

var nums2 = []int{2, 5, 6}
var n = 3

//var nums1 = []int{1}
//var m = 1
//
//var nums2 = []int{}
//var n = 0

//var nums1 = []int{0}
//var m = 0
//
//var nums2 = []int{1}
//var n = 1

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := 0
	if m == 0 {
		idx = len(nums2) - 1
	} else {
		idx = len(nums1) - 1
	}
	fmt.Println(idx, m)
	num1Index := m
	num2Index := n
	for i := idx; i >= 0; i-- {
		fmt.Println("for", nums1[i])
		if nums1[num1Index] >= nums2[num2Index] {
			nums1[i] = nums1[num1Index]
			num1Index--
			fmt.Println("if", nums1, num1Index)
		} else {
			nums1[i] = nums2[num2Index]
			num2Index--
			fmt.Println("else", nums1, num2Index)
		}
	}
	fmt.Println(nums1, nums2)
}
