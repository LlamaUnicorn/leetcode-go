package main

import "fmt"

func main() {
	merge(nums1, 3, nums2, 3)
}

var nums1 = []int{1, 2, 3, 0, 0, 0}
var m = 3

var nums2 = []int{2, 5, 6}
var n = 3

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := len(nums1) - 1
	//num1Index := m-1
	//num2Index := n-1
	//fmt.Println(idx)
	for i := idx; i >= 0; i-- {
		fmt.Println(nums1[i])
		if nums1[i] != 0 {
			fmt.Println(nums1[i])
		}
	}
	fmt.Println(nums1, nums2)
}
