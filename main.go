package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	var max1, max2, max3 *int
	for i := range nums {
		num := nums[i]

		if (max1 != nil && num == *max1) || (max2 != nil && num == *max2) || (max3 != nil && num == *max3) {
			continue
		}

		if max1 == nil || num > *max1 {
			max3 = max2
			max2 = max1
			max1 = &nums[i]
		} else if max2 == nil || num > *max2 {
			max3 = max2
			max2 = &nums[i]
		} else if max3 == nil || num > *max3 {
			max3 = &nums[i]
		}
	}

	if max3 != nil {
		return *max3
	}
	return *max1
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))    // Output: 1
	fmt.Println(thirdMax([]int{1, 2}))       // Output: 2
	fmt.Println(thirdMax([]int{2, 2, 3, 1})) // Output: 1
	fmt.Println(thirdMax([]int{5}))          // Output: 5
	fmt.Println(thirdMax([]int{2, 2, 2}))    // Output: 2
	fmt.Println(thirdMax([]int{-1, -2, -3})) // Output: -3
}
