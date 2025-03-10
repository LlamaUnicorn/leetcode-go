package merge_sorted_arrays

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
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
