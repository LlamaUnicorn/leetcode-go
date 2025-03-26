package mco

import "fmt"

// var Nums = []int{1, 1, 0, 1, 1, 1}
var Nums = []int{1, 0, 1, 1, 0, 1}

func MaxConsecutiveOnes(Nums []int) int {
	MaxCount := 0
	count := 0
	for i := 0; i < len(Nums); i++ {
		if Nums[i] == 1 {
			count++
			if count > MaxCount {
				MaxCount = count
			}
			continue
		} else {
			count = 0
		}
	}
	fmt.Println(MaxCount)
	return MaxCount
}
