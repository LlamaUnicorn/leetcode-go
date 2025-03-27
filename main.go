package main

import "fmt"

// var heights = []int{1, 1, 4, 2, 1, 3} // 3
var heights = []int{5, 1, 2, 3, 4} // 5  index out of range
//var heights = []int{1,2,3,4,5} // 0

func heightChecker(heights []int) int {
	res := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		//fmt.Println(res)
		if i == 0 || heights[i] >= res[i-1] {
			res[i] = heights[i]
		} else {
			for j := i; heights[i] < res[j-1]; j-- {
				res[j-1], res[j] = heights[i], res[j-1]
			}
		}
	}
	fmt.Println(res)
	ctr := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != res[i] {
			//fmt.Println(heights[i], res[i])
			ctr++
		}
	}
	return ctr
}

func main() {
	fmt.Println(heightChecker(heights))
}
