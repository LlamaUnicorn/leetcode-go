package main

import "fmt"

// Input:
var arr = []int{0, 3, 2, 1}

//Output: false

//Example 2:
//
//Input: arr = [3,5,5]
//Output: false
//Example 3:
//
//Input: arr = [0,3,2,1]
//Output: true

//var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func validMountainArray(arr []int) bool {
	ptr := 0
	for i := 1; i <= len(arr)-1; i++ {
		if (arr[i-1] < arr[i]) && (i-ptr == 1) {
			ptr++
			continue
		} else if arr[i-1] > arr[i] {
			if i == len(arr)-1 && ptr != 0 {
				return true
			}
			continue
		}
		return false
	}
	return false
}

func main() {
	fmt.Println(validMountainArray(arr))
}
