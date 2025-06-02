package main

import (
	"fmt"
	"slices"
)

func plusOne(digits []int) []int {
	adder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			adder := (digits[i] + adder) / 10
			digits[i] = (digits[i] + adder) % 10
		} else {
			digits[i] += 1
			return digits
		}
		if i == 0 {
			digits = slices.Insert(digits, 0, adder)
			return digits
		}
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // Output: [1,2,4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // Output: [4,3,2,2]
	fmt.Println(plusOne([]int{1, 9}))       // Output: [1,0]
}
