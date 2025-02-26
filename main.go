package main

import "fmt"

var nums = []int{12, 345, 2, 6, 7896}

func findNumbers(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		for nums[i] > 0 {
			fmt.Println("num[i]", nums[i])
			nums[i] /= 10
			count++
		}
		if count%2 == 0 {
			total += 1
			fmt.Println("total", total)
		}
		fmt.Println("count:", count)
	}
	fmt.Println("final total:", total)
	return total
}
func main() {
	findNumbers(nums)
}
