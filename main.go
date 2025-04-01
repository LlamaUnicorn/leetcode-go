package main

import "fmt"

// var nums = []int{3, 2, 1} // 1
var nums = []int{1, 2} // 2
//var nums = []int{2, 2, 3, 1} // 1

func thirdMax(nums []int) int {
	var max1, max2, max3 *int
	for i := 0; i < len(nums); i++ {
		if nums[i] == *max1 || nums[i] == *max2 || nums[i] == *max3 {
			continue
		}
		if max1 == nil || nums[i] > *max1 {
			max2 = max1
			max3 = max2
			max1 = &nums[i]
		} else if max2 == nil || nums[i] > *max2 {
			max3 = max2
			max2 = &nums[i]
		} else if max3 == nil || nums[i] > *max3 {
			max3 = &nums[i]
		}
	}
	fmt.Println(*max1, *max2, *max3)
	if max3 != nil {
		return *max3
	} else {
		return *max1
	}
}

func main() {
	fmt.Println(thirdMax(nums))
}
